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
	"regexp"
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
var _ resource.Resource = &ServiceSwitchportProfileParcelResource{}
var _ resource.ResourceWithImportState = &ServiceSwitchportProfileParcelResource{}

func NewServiceSwitchportProfileParcelResource() resource.Resource {
	return &ServiceSwitchportProfileParcelResource{}
}

type ServiceSwitchportProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *ServiceSwitchportProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_switchport_feature"
}

func (r *ServiceSwitchportProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Service Switchport Feature.").AddMinimumVersionDescription("20.12.0").String,

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
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface name: GigabitEthernet0/<>/<> when present").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set Interface name").String,
							Optional:            true,
						},
						"interface_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set type of switch port: access/trunk").AddStringEnumDescription("access", "trunk").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("access", "trunk"),
							},
						},
						"shutdown": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Administrative state").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"shutdown_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"speed": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set interface speed").AddStringEnumDescription("10", "100", "1000", "2500", "10000").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("10", "100", "1000", "2500", "10000"),
							},
						},
						"speed_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"duplex": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Duplex mode").AddStringEnumDescription("full", "half").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("full", "half"),
							},
						},
						"duplex_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"switchport_access_vlan": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set VLAN identifier associated with bridging domain").AddIntegerRangeDescription(1, 4094).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4094),
							},
						},
						"switchport_access_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"switchport_trunk_allowed_vlans": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure VLAN IDs used with the trunk").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^(?:[1-9]\d{0,2}|[1-3]\d{3}|40(?:[0-8]\d|9[0-4]))(?:[,-](?:[1-9]\d{0,2}|[1-3]\d{3}|40(?:[0-8]\d|9[0-4])))*$`), ""),
							},
						},
						"switchport_trunk_allowed_vlans_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"switchport_trunk_native_vlan": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure VLAN ID used for native VLAN").AddIntegerRangeDescription(1, 4094).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4094),
							},
						},
						"switchport_trunk_native_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"port_control": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set Port-Control Mode").AddStringEnumDescription("auto", "force-unauthorized", "force-authorized").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("auto", "force-unauthorized", "force-authorized"),
							},
						},
						"port_control_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"voice_vlan": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Voice Vlan").AddIntegerRangeDescription(1, 4094).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4094),
							},
						},
						"voice_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"pae_enable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set 802.1x Interface Pae Type").String,
							Optional:            true,
						},
						"pae_enable_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"mac_authentication_bypass": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("MAC Authentication Bypass").String,
							Optional:            true,
						},
						"mac_authentication_bypass_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"host_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set host mode").AddStringEnumDescription("single-host", "multi-auth", "multi-host", "multi-domain").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("single-host", "multi-auth", "multi-host", "multi-domain"),
							},
						},
						"host_mode_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"enable_periodic_reauth": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable Periodic Reauthentication").String,
							Optional:            true,
						},
						"enable_periodic_reauth_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"inactivity": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Periodic Reauthentication Inactivity Timeout (in seconds)").AddIntegerRangeDescription(1, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
						"inactivity_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"reauthentication": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Periodic Reauthentication Interval (in seconds)").AddIntegerRangeDescription(1, 1073741823).AddDefaultValueDescription("3600").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 1073741823),
							},
						},
						"reauthentication_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"control_direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set uni or bi directional authorization mode").AddStringEnumDescription("both", "in").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("both", "in"),
							},
						},
						"control_direction_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"restricted_vlan": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set Restricted VLAN ID").AddIntegerRangeDescription(1, 4094).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4094),
							},
						},
						"restricted_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"guest_vlan": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set vlan to drop non-802.1x enabled clients into if client is not in MAB list").AddIntegerRangeDescription(1, 4094).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4094),
							},
						},
						"guest_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"critical_vlan": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set Critical VLAN").AddIntegerRangeDescription(1, 4094).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4094),
							},
						},
						"critical_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"enable_voice": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable Critical Voice VLAN").String,
							Optional:            true,
						},
						"enable_voice_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"age_out_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set when a MAC table entry ages out (0 to disable, 10-1000000 otherwise)").AddIntegerRangeDescription(0, 1000000).AddDefaultValueDescription("300").String,
				Optional:            true,
			},
			"age_out_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"static_mac_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Add static MAC address entries for interface").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"mac_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set MAC address in xxxx.xxxx.xxxx format").String,
							Optional:            true,
						},
						"mac_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vlan_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure VLAN ID used with the mac and interface").AddIntegerRangeDescription(1, 4094).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4094),
							},
						},
						"vlan_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface name: GigabitEthernet0/<>/<>").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"interface_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *ServiceSwitchportProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *ServiceSwitchportProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ServiceSwitchport

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
func (r *ServiceSwitchportProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ServiceSwitchport

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
	stateCopy := state
	stateCopy.FeatureProfileId = types.StringNull()

	if stateCopy.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}
	if state.Version.IsNull() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *ServiceSwitchportProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ServiceSwitchport

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
func (r *ServiceSwitchportProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ServiceSwitchport

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
func (r *ServiceSwitchportProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "service_switchport_feature_id" + ",feature_profile_id"
	if len(parts) != (count + 1) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: %s. Got: %q", pattern, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("feature_profile_id"), parts[1])...)
}

// End of section. //template:end import
