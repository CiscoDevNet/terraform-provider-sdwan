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
	"strings"
	"sync"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/gjson"
)

var _ resource.Resource = &NetworkHierarchyNodeResource{}
var _ resource.ResourceWithImportState = &NetworkHierarchyNodeResource{}

func NewNetworkHierarchyNodeResource() resource.Resource {
	return &NetworkHierarchyNodeResource{}
}

type NetworkHierarchyNodeResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *NetworkHierarchyNodeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_hierarchy_node"
}

func (r *NetworkHierarchyNodeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Network Hierarchy Node (group, region, or site).").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the node").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The description of the node").String,
				Optional:            true,
			},
			"parent_group": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the parent group. Use 'Global' for top-level nodes.").String,
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The type of node").AddStringEnumDescription("group", "region", "site").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("group", "region", "site"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"site_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The site ID (only for site type nodes)").AddIntegerRangeDescription(1, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
			},
			"is_secondary": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether this is a secondary region (only for region type nodes)").String,
				Optional:            true,
			},
			"address": schema.SingleNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The address of the site (only for site type nodes)").String,
				Optional:            true,
				Attributes: map[string]schema.Attribute{
					"street": schema.StringAttribute{
						MarkdownDescription: helpers.NewAttributeDescription("Street address").String,
						Required:            true,
					},
					"city": schema.StringAttribute{
						MarkdownDescription: helpers.NewAttributeDescription("City").String,
						Required:            true,
					},
					"state": schema.StringAttribute{
						MarkdownDescription: helpers.NewAttributeDescription("State or province").String,
						Required:            true,
					},
					"country": schema.StringAttribute{
						MarkdownDescription: helpers.NewAttributeDescription("Country").String,
						Required:            true,
					},
					"zipcode": schema.StringAttribute{
						MarkdownDescription: helpers.NewAttributeDescription("Zip or postal code").String,
						Required:            true,
					},
				},
			},
			"controllers": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of controller UUIDs to assign to this region (only applicable for region type nodes)").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func (r *NetworkHierarchyNodeResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

func (r *NetworkHierarchyNodeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NetworkHierarchyNode

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	hierarchyRes, err := r.client.Get(plan.getHierarchyListPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve network hierarchy (GET), got error: %s, %s", err, hierarchyRes.String()))
		return
	}

	parentId, errMsg := plan.resolveParentGroupToId(hierarchyRes)
	if errMsg != "" {
		resp.Diagnostics.AddError("Configuration Error", errMsg)
		return
	}

	body := plan.toBody(ctx, parentId)

	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	if res.Get("created.0.uuid").Exists() {
		plan.Id = types.StringValue(res.Get("created.0.uuid").String())
	} else {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to get UUID from response: %s", res.String()))
		return
	}

	// Add controller assignment for regions
	if plan.Type.ValueString() == "region" && !plan.Controllers.IsNull() && !plan.Controllers.IsUnknown() {
		nodeRes, err := r.client.Get(plan.getPath() + url.QueryEscape(plan.Id.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve node for regionId (GET), got error: %s, %s", err, nodeRes.String()))
			return
		}

		regionId := nodeRes.Get("data.hierarchyId.regionId").Int()
		if regionId == 0 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to get regionId from node response: %s", nodeRes.String()))
			return
		}

		wfBody := plan.toWorkflowBody(ctx, regionId)
		wfRes, err := r.client.Post("/workflow/execute", wfBody)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to assign controllers (POST workflow/execute), got error: %s, %s", err, wfRes.String()))
			return
		}
		resp.Diagnostics.AddWarning("Controller Assignment", "vSmart controller assignment to region submitted successfully. Changes may take a few seconds to reflect in the system.")
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworkHierarchyNodeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkHierarchyNode

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.ValueString()))

	res, err := r.client.Get(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if strings.Contains(res.Get("error.message").String(), "Failed to find specified resource") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	parentUuid := state.fromBody(ctx, res)

	hierarchyRes, err := r.client.Get(state.getHierarchyListPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve network hierarchy (GET), got error: %s, %s", err, hierarchyRes.String()))
		return
	}

	parentGroupName := state.resolveParentIdToGroup(hierarchyRes, parentUuid)
	if parentGroupName != "" {
		state.ParentGroup = types.StringValue(parentGroupName)
	} else {
		state.ParentGroup = types.StringNull()
	}

	if state.Type.ValueString() == "region" {
		regionId := res.Get("data.hierarchyId.regionId").Int()
		if regionId > 0 {
			controllersRes, err := r.client.Get(state.getControllersPath())
			if err != nil {
				resp.Diagnostics.AddWarning("Client Warning", fmt.Sprintf("Failed to retrieve controllers (GET), got error: %s", err))
			} else {
				state.readControllersFromResponse(ctx, controllersRes, regionId)
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworkHierarchyNodeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state NetworkHierarchyNode

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Name.ValueString()))

	hierarchyRes, err := r.client.Get(plan.getHierarchyListPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve network hierarchy (GET), got error: %s, %s", err, hierarchyRes.String()))
		return
	}

	parentId, errMsg := plan.resolveParentGroupToId(hierarchyRes)
	if errMsg != "" {
		resp.Diagnostics.AddError("Configuration Error", errMsg)
		return
	}

	var currentNodeRes gjson.Result
	nodeNeedsUpdate := plan.hasChanges(ctx, &state)
	controllerNeedsUpdate := plan.Type.ValueString() == "region" && !plan.Controllers.Equal(state.Controllers)

	if nodeNeedsUpdate || controllerNeedsUpdate {
		currentNodeRes, err = r.client.Get(plan.getPath() + url.QueryEscape(plan.Id.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve current node (GET), got error: %s, %s", err, currentNodeRes.String()))
			return
		}
	}

	if nodeNeedsUpdate {
		body := plan.toUpdateBody(ctx, currentNodeRes, parentId)
		r.updateMutex.Lock()
		res, err := r.client.Put(plan.getPath()+url.QueryEscape(plan.Id.ValueString()), body)
		r.updateMutex.Unlock()
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	} else {
		tflog.Debug(ctx, fmt.Sprintf("%s: No changes detected", plan.Name.ValueString()))
	}

	if controllerNeedsUpdate {
		regionId := currentNodeRes.Get("data.hierarchyId.regionId").Int()
		if regionId == 0 {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to get regionId from node response: %s", currentNodeRes.String()))
			return
		}

		wfBody := plan.toWorkflowBody(ctx, regionId)
		r.updateMutex.Lock()
		wfRes, err := r.client.Post("/workflow/execute", wfBody)
		r.updateMutex.Unlock()
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update controllers (POST workflow/execute), got error: %s, %s", err, wfRes.String()))
			return
		}
		resp.Diagnostics.AddWarning("Controller Assignment", "vSmart controller assignment to region submitted successfully. Changes may take a few seconds to reflect in the system.")
	}

	plan.Id = state.Id

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworkHierarchyNodeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworkHierarchyNode

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *NetworkHierarchyNodeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
