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
	"strings"
	"sync"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

var _ resource.Resource = &NetworkHierarchySecurityLoggingResource{}
var _ resource.ResourceWithImportState = &NetworkHierarchySecurityLoggingResource{}

func NewNetworkHierarchySecurityLoggingResource() resource.Resource {
	return &NetworkHierarchySecurityLoggingResource{}
}

type NetworkHierarchySecurityLoggingResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *NetworkHierarchySecurityLoggingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_hierarchy_security_logging"
}

func (r *NetworkHierarchySecurityLoggingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage Network Hierarchy Security Logging settings.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"node_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The UUID of the Global network hierarchy node. This is automatically fetched from the SD-WAN Manager.").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"high_speed_logging": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("High speed logging configuration (max 4 entries)").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vrf": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("VRF name or ID").String,
							Required:            true,
						},
						"server_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Server IPv4 or IPv6 address").String,
							Required:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Server port number").AddIntegerRangeDescription(1, 65535).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
					},
				},
			},
			"utd_syslog": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("UTD syslog configuration (max 1 entry)").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpn": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("VPN name or ID").String,
							Required:            true,
						},
						"server_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Server IPv4 address").String,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *NetworkHierarchySecurityLoggingResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

func (r *NetworkHierarchySecurityLoggingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NetworkHierarchySecurityLogging

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "NetworkHierarchySecurityLogging: Beginning Create")

	hierarchyRes, err := r.client.Get(plan.getNetworkHierarchyListPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve network hierarchy, got error: %s", err))
		return
	}

	globalNodeId, err := plan.getGlobalNodeId(hierarchyRes)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to get Global node ID: %s", err))
		return
	}
	plan.NodeId = types.StringValue(globalNodeId)

	tflog.Debug(ctx, fmt.Sprintf("%s: Using Global node ID", plan.NodeId.ValueString()))

	// Check if security-logging already exists for this node (only 1 allowed per node)
	// GET returns {"data":[{"id":"..."}]} - array format
	existingRes, _ := r.client.Get(plan.getPath())
	if existingRes.Get("data.0.id").Exists() {
		existingId := existingRes.Get("data.0.id").String()
		plan.Id = types.StringValue(existingId)
		tflog.Debug(ctx, fmt.Sprintf("%s: Security logging already exists with ID %s, updating instead", plan.NodeId.ValueString(), existingId))

		body := plan.toBody(ctx)
		res, err := r.client.Put(plan.getPathWithId(), body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	} else {
		body := plan.toBody(ctx)
		res, err := r.client.Post(plan.getPath(), body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
			return
		}

		if res.Get("id").Exists() {
			plan.Id = types.StringValue(res.Get("id").String())
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to get ID from response: %s", res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.NodeId.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworkHierarchySecurityLoggingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkHierarchySecurityLogging

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	if state.NodeId.IsNull() || state.NodeId.ValueString() == "" {
		hierarchyRes, err := r.client.Get(state.getNetworkHierarchyListPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve network hierarchy, got error: %s", err))
			return
		}

		globalNodeId, err := state.getGlobalNodeId(hierarchyRes)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to get Global node ID: %s", err))
			return
		}
		state.NodeId = types.StringValue(globalNodeId)
	}

	res, err := r.client.Get(state.getPathWithId())
	if strings.Contains(res.Get("error.message").String(), "Failed to find specified resource") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.NodeId.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworkHierarchySecurityLoggingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state NetworkHierarchySecurityLogging

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

	plan.NodeId = state.NodeId

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.NodeId.ValueString()))

	if plan.hasChanges(ctx, &state) {
		body := plan.toBody(ctx)
		r.updateMutex.Lock()
		res, err := r.client.Put(plan.getPathWithId(), body)
		r.updateMutex.Unlock()
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	} else {
		tflog.Debug(ctx, fmt.Sprintf("%s: No changes detected", plan.NodeId.ValueString()))
	}

	plan.Id = state.Id

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.NodeId.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworkHierarchySecurityLoggingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworkHierarchySecurityLogging

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.NodeId.ValueString()))

	body := state.toEmptyBody(ctx)
	res, err := r.client.Put(state.getPathWithId(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (PUT with empty data), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.NodeId.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *NetworkHierarchySecurityLoggingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
