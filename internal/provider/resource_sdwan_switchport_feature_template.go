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
var _ resource.Resource = &SwitchportFeatureTemplateResource{}
var _ resource.ResourceWithImportState = &SwitchportFeatureTemplateResource{}

func NewSwitchportFeatureTemplateResource() resource.Resource {
	return &SwitchportFeatureTemplateResource{}
}

type SwitchportFeatureTemplateResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *SwitchportFeatureTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switchport_feature_template"
}

func (r *SwitchportFeatureTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Switchport feature template.").AddMinimumVersionDescription("15.0.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the feature template",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the feature template",
				Computed:            true,
			},
			"template_type": schema.StringAttribute{
				MarkdownDescription: "The template type",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the feature template",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the feature template",
				Required:            true,
			},
			"device_types": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of supported device types").AddStringEnumDescription("vedge-C8000V", "vedge-C8300-1N1S-4T2X", "vedge-C8300-1N1S-6T", "vedge-C8300-2N2S-6T", "vedge-C8300-2N2S-4T2X", "vedge-C8500-12X4QC", "vedge-C8500-12X", "vedge-C8500-20X6C", "vedge-C8500L-8S4X", "vedge-C8200-1N-4T", "vedge-C8200L-1N-4T").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"slot": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of Slots").AddIntegerRangeDescription(0, 31).AddDefaultValueDescription("0").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 31),
				},
			},
			"sub_slot": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of Sub-Slots").AddIntegerRangeDescription(0, 31).AddDefaultValueDescription("0").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 31),
				},
			},
			"module_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Module type").AddStringEnumDescription("4", "8", "22", "50").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("4", "8", "22", "50"),
				},
			},
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface name: GigabitEthernet0/<>/<> when present").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set Interface name").String,
							Optional:            true,
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"switchport_mode": schema.StringAttribute{
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
						"dot1x_enable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set 802.1x on off").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"dot1x_enable_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dot1x_port_control": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set Port-Control Mode").AddStringEnumDescription("auto", "force-unauthorized", "force-authorized").AddDefaultValueDescription("auto").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("auto", "force-unauthorized", "force-authorized"),
							},
						},
						"dot1x_port_control_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dot1x_authentication_order": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify authentication methods in the order of preference").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"dot1x_authentication_order_variable": schema.StringAttribute{
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
						"dot1x_pae_enable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set 802.1x Interface Pae Type").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"dot1x_pae_enable_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dot1x_mac_authentication_bypass": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("MAC Authentication Bypass").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"dot1x_mac_authentication_bypass_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dot1x_host_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set host mode").AddStringEnumDescription("single-host", "multi-auth", "multi-host", "multi-domain").AddDefaultValueDescription("single-host").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("single-host", "multi-auth", "multi-host", "multi-domain"),
							},
						},
						"dot1x_host_mode_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dot1x_enable_periodic_reauth": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable Periodic Reauthentication").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"dot1x_enable_periodic_reauth_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dot1x_periodic_reauth_inactivity_timeout": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Periodic Reauthentication Inactivity Timeout (in seconds)").AddIntegerRangeDescription(1, 1440).AddDefaultValueDescription("60").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 1440),
							},
						},
						"dot1x_periodic_reauth_inactivity_timeout_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dot1x_periodic_reauth_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Periodic Reauthentication Interval (in seconds)").AddIntegerRangeDescription(0, 1440).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 1440),
							},
						},
						"dot1x_periodic_reauth_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dot1x_control_direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set uni or bi directional authorization mode").AddStringEnumDescription("both", "in").AddDefaultValueDescription("both").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("both", "in"),
							},
						},
						"dot1x_control_direction_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dot1x_restricted_vlan": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set Restricted VLAN ID").AddIntegerRangeDescription(1, 4094).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4094),
							},
						},
						"dot1x_restricted_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dot1x_guest_vlan": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set vlan to drop non-802.1x enabled clients into if client is not in MAB list").AddIntegerRangeDescription(1, 4094).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4094),
							},
						},
						"dot1x_guest_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dot1x_critical_vlan": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set Critical VLAN").AddIntegerRangeDescription(1, 4094).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4094),
							},
						},
						"dot1x_critical_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dot1x_enable_criticial_voice_vlan": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable Critical Voice VLAN").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"dot1x_enable_criticial_voice_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"age_out_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set when a MAC table entry ages out (0 to disable, 10-1000000 otherwise)").AddIntegerRangeDescription(0, 1000000).AddDefaultValueDescription("300").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1000000),
				},
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
						"if_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface name: GigabitEthernet0/<>/<>").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"if_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vlan": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure VLAN ID used with the mac and interface").AddIntegerRangeDescription(1, 4094).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4094),
							},
						},
						"vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *SwitchportFeatureTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *SwitchportFeatureTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Switchport

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post("/template/feature", body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("templateId").String())
	plan.Version = types.Int64Value(0)
	plan.TemplateType = types.StringValue(plan.getModel())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *SwitchportFeatureTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Switchport

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get("/template/feature/object/" + url.QueryEscape(state.Id.ValueString()))
	if res.Get("error.message").String() == "Invalid Template Id" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)
	if state.Version.IsNull() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *SwitchportFeatureTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Switchport

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
	r.updateMutex.Lock()
	res, err := r.client.Put("/template/feature/"+url.QueryEscape(plan.Id.ValueString()), body)
	r.updateMutex.Unlock()
	if err != nil {
		if res.Get("error.message").String() == "Template locked in edit mode." {
			resp.Diagnostics.AddWarning("Client Warning", fmt.Sprintf("Failed to modify template due to template being locked by another change. Template changes will not be applied. Re-run 'terraform apply' to try again."))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	if plan.hasChanges(ctx, &state) {
		plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)
	} else {
		plan.Version = state.Version
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *SwitchportFeatureTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Switchport

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete("/template/feature/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && res.Get("error.message").String() != "Invalid Template Id" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *SwitchportFeatureTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
