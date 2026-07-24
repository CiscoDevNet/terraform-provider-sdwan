// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"net/url"
	"sync"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

var _ resource.Resource = &ActivateTopologyGroupResource{}
var _ resource.ResourceWithImportState = &ActivateTopologyGroupResource{}

func NewActivateTopologyGroupResource() resource.Resource {
	return &ActivateTopologyGroupResource{}
}

type ActivateTopologyGroupResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
	taskTimeout *int64
}

func (r *ActivateTopologyGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_activate_topology_group"
}

func (r *ActivateTopologyGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: helpers.NewAttributeDescription("This resource can activate a topology group. Only one topology group can be active at a time. To switch the active group, change the `id` attribute on the existing resource in place instead of replacing the resource (destroy + create); an in-place change issues a single activation that supersedes the previous group, whereas a replacement deactivates the previous group first and withdraws all overlay control policy from the vSmarts during the switch.").AddMinimumVersionDescription("20.15.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the topology group to activate",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"feature_versions": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of all associated feature versions. Any change to this list will trigger a re-deployment of the topology group.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"deployed_version": schema.Int64Attribute{
				MarkdownDescription: "Server-side version of the topology group captured at last activation. Used internally for drift detection.",
				Computed:            true,
			},
		},
	}
}

func (r *ActivateTopologyGroupResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
	r.taskTimeout = req.ProviderData.(*SdwanProviderData).TaskTimeout
}

func (r *ActivateTopologyGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ActivateTopologyGroup

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	r.updateMutex.Lock()
	err := plan.activateTopologyGroup(ctx, r.client, r.taskTimeout)
	r.updateMutex.Unlock()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to activate topology group, got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *ActivateTopologyGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ActivateTopologyGroup

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	serverVersion := types.Int64Null()
	if state.Id.ValueString() != "" {
		res, err := r.client.Get(state.getPath() + url.QueryEscape(state.Id.ValueString()))
		if res.Raw == "" && err == nil {
			resp.State.RemoveResource(ctx)
			return
		} else if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}

		serverVersion = types.Int64Value(res.Get("version").Int())

		if !res.Get("activeStatus").Bool() {
			state.Id = types.StringValue("")
		} else if !state.DeployedVersion.IsNull() && res.Get("version").Int() != state.DeployedVersion.ValueInt64() {
			// Server version advanced since last activation (e.g. a parcel was removed).
			// Null both markers to force a re-activation: deployed_version is Computed and
			// nulling it alone yields no plan diff, so we also null the config-backed
			// feature_versions to trigger Update. The gate checks deployed_version==null
			// before feature_versions==null, keeping removal drift distinct from import.
			serverV := res.Get("version").Int()
			deployedV := state.DeployedVersion.ValueInt64()
			tflog.Debug(ctx, fmt.Sprintf("%s: Removal drift detected (server version %d != deployed version %d); re-activating on next apply", state.Id.ValueString(), serverV, deployedV))
			resp.Diagnostics.AddWarning(
				"Topology group drift detected",
				fmt.Sprintf("Server-side version (%d) of topology group %s differs from the last deployed version (%d). "+
					"A parcel or referenced object was likely removed, so the group will be re-activated on apply to converge. "+
					"The 'feature_versions' and 'deployed_version' changes shown in the plan are an internal drift marker, "+
					"not a configuration change.", serverV, state.Id.ValueString(), deployedV),
			)
			state.DeployedVersion = types.Int64Null()
			state.FeatureVersions = types.ListNull(types.StringType)
		}
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	if imp {
		state.processImport(ctx, serverVersion)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *ActivateTopologyGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ActivateTopologyGroup

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	// Decide whether a re-activation is required, using two null markers evaluated in
	// priority order:
	//   - deployed_version == null -> removal drift (set by Read) -> deploy
	//   - feature_versions == null -> import reconcile -> skip deploy, resync
	//     deployed_version from a fresh GET (the topology group's own PUT earlier in
	//     this apply may have bumped the version).
	needsDeploy := false
	importReconcile := false
	switch {
	case !plan.Id.Equal(state.Id):
		tflog.Debug(ctx, fmt.Sprintf("%s: Topology group ID changed, activating new topology group", plan.Id.ValueString()))
		needsDeploy = true
	case state.DeployedVersion.IsNull():
		tflog.Debug(ctx, fmt.Sprintf("%s: Deployed version invalidated (removal drift), re-deploying topology group", plan.Id.ValueString()))
		needsDeploy = true
	case state.FeatureVersions.IsNull():
		tflog.Debug(ctx, fmt.Sprintf("%s: Feature versions null (import reconcile), skipping re-deployment", plan.Id.ValueString()))
		importReconcile = true
	case plan.hasFeatureVersionChanges(ctx, &state):
		tflog.Debug(ctx, fmt.Sprintf("%s: Feature versions changed, re-deploying topology group", plan.Id.ValueString()))
		needsDeploy = true
	}

	if needsDeploy {
		r.updateMutex.Lock()
		err := plan.activateTopologyGroup(ctx, r.client, r.taskTimeout)
		r.updateMutex.Unlock()
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to activate topology group, got error: %s", err))
			return
		}
	} else if importReconcile {
		// No re-activation; just resync deployed_version from a fresh GET so a later
		// PUT by the topology group resource does not read back as drift next apply.
		res, err := r.client.Get(plan.getPath() + url.QueryEscape(plan.Id.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}
		plan.DeployedVersion = types.Int64Value(res.Get("version").Int())
	} else {
		// No change: carry the baseline forward so the Computed deployed_version
		// resolves to a known value.
		plan.DeployedVersion = state.DeployedVersion
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ActivateTopologyGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ActivateTopologyGroup

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	if state.Id.ValueString() != "" {
		r.updateMutex.Lock()
		err := state.deactivateTopologyGroup(ctx, r.client, r.taskTimeout)
		r.updateMutex.Unlock()
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to deactivate topology group, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *ActivateTopologyGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
