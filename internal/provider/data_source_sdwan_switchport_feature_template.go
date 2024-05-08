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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SwitchportFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &SwitchportFeatureTemplateDataSource{}
)

func NewSwitchportFeatureTemplateDataSource() datasource.DataSource {
	return &SwitchportFeatureTemplateDataSource{}
}

type SwitchportFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *SwitchportFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switchport_feature_template"
}

func (d *SwitchportFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Switchport feature template.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the feature template",
				Optional:            true,
				Computed:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the feature template",
				Computed:            true,
			},
			"template_type": schema.StringAttribute{
				MarkdownDescription: "The template type",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the feature template",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the feature template",
				Computed:            true,
			},
			"device_types": schema.SetAttribute{
				MarkdownDescription: "List of supported device types",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"slot": schema.Int64Attribute{
				MarkdownDescription: "Number of Slots",
				Computed:            true,
			},
			"sub_slot": schema.Int64Attribute{
				MarkdownDescription: "Number of Sub-Slots",
				Computed:            true,
			},
			"module_type": schema.StringAttribute{
				MarkdownDescription: "Module type",
				Computed:            true,
			},
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "Interface name: GigabitEthernet0/<>/<> when present",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Set Interface name",
							Computed:            true,
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"switchport_mode": schema.StringAttribute{
							MarkdownDescription: "Set type of switch port: access/trunk",
							Computed:            true,
						},
						"shutdown": schema.BoolAttribute{
							MarkdownDescription: "Administrative state",
							Computed:            true,
						},
						"shutdown_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"speed": schema.StringAttribute{
							MarkdownDescription: "Set interface speed",
							Computed:            true,
						},
						"speed_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"duplex": schema.StringAttribute{
							MarkdownDescription: "Duplex mode",
							Computed:            true,
						},
						"duplex_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"switchport_access_vlan": schema.Int64Attribute{
							MarkdownDescription: "Set VLAN identifier associated with bridging domain",
							Computed:            true,
						},
						"switchport_access_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"switchport_trunk_allowed_vlans": schema.StringAttribute{
							MarkdownDescription: "Configure VLAN IDs used with the trunk",
							Computed:            true,
						},
						"switchport_trunk_allowed_vlans_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"switchport_trunk_native_vlan": schema.Int64Attribute{
							MarkdownDescription: "Configure VLAN ID used for native VLAN",
							Computed:            true,
						},
						"switchport_trunk_native_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dot1x_enable": schema.BoolAttribute{
							MarkdownDescription: "Set 802.1x on off",
							Computed:            true,
						},
						"dot1x_enable_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dot1x_port_control": schema.StringAttribute{
							MarkdownDescription: "Set Port-Control Mode",
							Computed:            true,
						},
						"dot1x_port_control_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dot1x_authentication_order": schema.SetAttribute{
							MarkdownDescription: "Specify authentication methods in the order of preference",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"dot1x_authentication_order_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"voice_vlan": schema.Int64Attribute{
							MarkdownDescription: "Configure Voice Vlan",
							Computed:            true,
						},
						"voice_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dot1x_pae_enable": schema.BoolAttribute{
							MarkdownDescription: "Set 802.1x Interface Pae Type",
							Computed:            true,
						},
						"dot1x_pae_enable_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dot1x_mac_authentication_bypass": schema.BoolAttribute{
							MarkdownDescription: "MAC Authentication Bypass",
							Computed:            true,
						},
						"dot1x_mac_authentication_bypass_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dot1x_host_mode": schema.StringAttribute{
							MarkdownDescription: "Set host mode",
							Computed:            true,
						},
						"dot1x_host_mode_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dot1x_enable_periodic_reauth": schema.BoolAttribute{
							MarkdownDescription: "Enable Periodic Reauthentication",
							Computed:            true,
						},
						"dot1x_enable_periodic_reauth_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dot1x_periodic_reauth_inactivity_timeout": schema.Int64Attribute{
							MarkdownDescription: "Periodic Reauthentication Inactivity Timeout (in seconds)",
							Computed:            true,
						},
						"dot1x_periodic_reauth_inactivity_timeout_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dot1x_periodic_reauth_interval": schema.Int64Attribute{
							MarkdownDescription: "Periodic Reauthentication Interval (in seconds)",
							Computed:            true,
						},
						"dot1x_periodic_reauth_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dot1x_control_direction": schema.StringAttribute{
							MarkdownDescription: "Set uni or bi directional authorization mode",
							Computed:            true,
						},
						"dot1x_control_direction_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dot1x_restricted_vlan": schema.Int64Attribute{
							MarkdownDescription: "Set Restricted VLAN ID",
							Computed:            true,
						},
						"dot1x_restricted_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dot1x_guest_vlan": schema.Int64Attribute{
							MarkdownDescription: "Set vlan to drop non-802.1x enabled clients into if client is not in MAB list",
							Computed:            true,
						},
						"dot1x_guest_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dot1x_critical_vlan": schema.Int64Attribute{
							MarkdownDescription: "Set Critical VLAN",
							Computed:            true,
						},
						"dot1x_critical_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dot1x_enable_criticial_voice_vlan": schema.BoolAttribute{
							MarkdownDescription: "Enable Critical Voice VLAN",
							Computed:            true,
						},
						"dot1x_enable_criticial_voice_vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"age_out_time": schema.Int64Attribute{
				MarkdownDescription: "Set when a MAC table entry ages out (0 to disable, 10-1000000 otherwise)",
				Computed:            true,
			},
			"age_out_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"static_mac_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Add static MAC address entries for interface",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"mac_address": schema.StringAttribute{
							MarkdownDescription: "Set MAC address in xxxx.xxxx.xxxx format",
							Computed:            true,
						},
						"mac_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"if_name": schema.StringAttribute{
							MarkdownDescription: "Interface name: GigabitEthernet0/<>/<>",
							Computed:            true,
						},
						"if_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"vlan": schema.Int64Attribute{
							MarkdownDescription: "Configure VLAN ID used with the mac and interface",
							Computed:            true,
						},
						"vlan_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SwitchportFeatureTemplateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *SwitchportFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *SwitchportFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Switchport

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	if config.Id.IsNull() && !config.Name.IsNull() {
		res, err := d.client.Get("/template/feature")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve existing templates, got error: %s", err))
			return
		}
		if value := res.Get("data"); len(value.Array()) > 0 {
			value.ForEach(func(k, v gjson.Result) bool {
				if config.Name.ValueString() == v.Get("templateName").String() {
					config.Id = types.StringValue(v.Get("templateId").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found feature template with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
					return false
				}
				return true
			})
		}
		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find feature template with name: %s", config.Name.ValueString()))
			return
		}
	}

	res, err := d.client.Get("/template/feature/object/" + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
