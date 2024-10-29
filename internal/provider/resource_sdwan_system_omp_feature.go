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

// Section below is generated&owned by "gen/generator.go". //template:begin imports
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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &SystemOMPProfileParcelResource{}
var _ resource.ResourceWithImportState = &SystemOMPProfileParcelResource{}

func NewSystemOMPProfileParcelResource() resource.Resource {
	return &SystemOMPProfileParcelResource{}
}

type SystemOMPProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *SystemOMPProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_omp_feature"
}

func (r *SystemOMPProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a System OMP Feature.").AddMinimumVersionDescription("20.12.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Feature",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Feature",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Feature",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Feature",
				Optional:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Feature Profile ID").String,
				Required:            true,
			},
			"graceful_restart": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Graceful Restart for OMP").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"graceful_restart_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"overlay_as": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Overlay AS Number").AddIntegerRangeDescription(1, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
			},
			"overlay_as_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"paths_advertised_per_prefix": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of Paths Advertised per Prefix").AddIntegerRangeDescription(1, 16).AddDefaultValueDescription("4").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 16),
				},
			},
			"paths_advertised_per_prefix_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ecmp_limit": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set maximum number of OMP paths to install in cEdge route table").AddIntegerRangeDescription(1, 0).AddDefaultValueDescription("4").String,
				Optional:            true,
			},
			"ecmp_limit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Shutdown").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"shutdown_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"omp_admin_distance_ipv4": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("OMP Admin Distance IPv4").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("251").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"omp_admin_distance_ipv4_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"omp_admin_distance_ipv6": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("OMP Admin Distance IPv6").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("251").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"omp_admin_distance_ipv6_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"advertisement_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Advertisement Interval (seconds)").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("1").String,
				Optional:            true,
			},
			"advertisement_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"graceful_restart_timer": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Graceful Restart Timer (seconds)").AddIntegerRangeDescription(1, 604800).AddDefaultValueDescription("43200").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 604800),
				},
			},
			"graceful_restart_timer_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"eor_timer": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("EOR Timer").AddIntegerRangeDescription(1, 3600).AddDefaultValueDescription("300").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 3600),
				},
			},
			"eor_timer_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"holdtime": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hold Time (seconds)").AddDefaultValueDescription("60").String,
				Optional:            true,
			},
			"holdtime_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"advertise_ipv4_bgp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("BGP").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"advertise_ipv4_bgp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"advertise_ipv4_ospf": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("OSPF").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"advertise_ipv4_ospf_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"advertise_ipv4_ospf_v3": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("OSPFV3").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"advertise_ipv4_ospf_v3_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"advertise_ipv4_connected": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Connected").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"advertise_ipv4_connected_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"advertise_ipv4_static": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"advertise_ipv4_static_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"advertise_ipv4_eigrp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("EIGRP").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"advertise_ipv4_eigrp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"advertise_ipv4_lisp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("LISP").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"advertise_ipv4_lisp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"advertise_ipv4_isis": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ISIS").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"advertise_ipv4_isis_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"advertise_ipv6_bgp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("BGP").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"advertise_ipv6_bgp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"advertise_ipv6_ospf": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("OSPF").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"advertise_ipv6_ospf_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"advertise_ipv6_connected": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Connected").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"advertise_ipv6_connected_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"advertise_ipv6_static": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"advertise_ipv6_static_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"advertise_ipv6_eigrp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("EIGRP").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"advertise_ipv6_eigrp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"advertise_ipv6_lisp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("LISP").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"advertise_ipv6_lisp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"advertise_ipv6_isis": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ISIS").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"advertise_ipv6_isis_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ignore_region_path_length": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Treat hierarchical and direct (secondary region) paths equally").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ignore_region_path_length_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"transport_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Transport Gateway Path Behavior").AddStringEnumDescription("prefer", "ecmp-with-direct-path").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("prefer", "ecmp-with-direct-path"),
				},
			},
			"transport_gateway_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"site_types": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Site Types").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"site_types_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
		},
	}
}

func (r *SystemOMPProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *SystemOMPProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemOMP

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("parcelId").String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *SystemOMPProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemOMP

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if res.Get("error.message").String() == "Invalid feature Id" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}
	if state.Version == types.Int64Null() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *SystemOMPProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemOMP

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Name.ValueString()))

	body := plan.toBody(ctx)
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *SystemOMPProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemOMP

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && res.Get("error.message").String() != "Invalid Template Id" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *SystemOMPProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "system_omp_feature_id" + ",feature_profile_id"
	if len(parts) != (count + 1) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: %s. Got: %q, %q", pattern, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("feature_profile_id"), parts[1])...)
}

// End of section. //template:end import
