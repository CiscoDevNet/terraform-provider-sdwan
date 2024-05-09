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
	_ datasource.DataSource              = &CiscoSystemFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &CiscoSystemFeatureTemplateDataSource{}
)

func NewCiscoSystemFeatureTemplateDataSource() datasource.DataSource {
	return &CiscoSystemFeatureTemplateDataSource{}
}

type CiscoSystemFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *CiscoSystemFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_system_feature_template"
}

func (d *CiscoSystemFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Cisco System feature template.",

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
			"timezone": schema.StringAttribute{
				MarkdownDescription: "Set the timezone",
				Computed:            true,
			},
			"timezone_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"hostname": schema.StringAttribute{
				MarkdownDescription: "Set the hostname",
				Computed:            true,
			},
			"hostname_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"system_description": schema.StringAttribute{
				MarkdownDescription: "Set a text description of the device",
				Computed:            true,
			},
			"system_description_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"location": schema.StringAttribute{
				MarkdownDescription: "Set the location of the device",
				Computed:            true,
			},
			"location_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"latitude": schema.Float64Attribute{
				MarkdownDescription: "Set the device’s physical latitude",
				Computed:            true,
			},
			"latitude_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"longitude": schema.Float64Attribute{
				MarkdownDescription: "Set the device’s physical longitude",
				Computed:            true,
			},
			"longitude_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"geo_fencing": schema.BoolAttribute{
				MarkdownDescription: "Enable Geo fencing",
				Computed:            true,
			},
			"geo_fencing_range": schema.Int64Attribute{
				MarkdownDescription: "Set the device’s geo fencing range",
				Computed:            true,
			},
			"geo_fencing_range_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"geo_fencing_sms": schema.BoolAttribute{
				MarkdownDescription: "Enable Geo fencing",
				Computed:            true,
			},
			"geo_fencing_sms_phone_numbers": schema.ListNestedAttribute{
				MarkdownDescription: "Set device’s geo fencing SMS phone number",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"number": schema.StringAttribute{
							MarkdownDescription: "Mobile number, ex: +1231234414",
							Computed:            true,
						},
						"number_variable": schema.StringAttribute{
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
			"device_groups": schema.SetAttribute{
				MarkdownDescription: "Device groups (Use comma(,) for multiple groups)",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"device_groups_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"controller_group_list": schema.SetAttribute{
				MarkdownDescription: "Configure a list of comma-separated device groups",
				ElementType:         types.Int64Type,
				Computed:            true,
			},
			"controller_group_list_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"system_ip": schema.StringAttribute{
				MarkdownDescription: "Set the system IP address",
				Computed:            true,
			},
			"system_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"overlay_id": schema.Int64Attribute{
				MarkdownDescription: "Set the Overlay ID",
				Computed:            true,
			},
			"overlay_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"site_id": schema.Int64Attribute{
				MarkdownDescription: "Set the site identifier",
				Computed:            true,
			},
			"site_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"port_offset": schema.Int64Attribute{
				MarkdownDescription: "Set the TLOC port offset when multiple devices are behind a NAT",
				Computed:            true,
			},
			"port_offset_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"port_hopping": schema.BoolAttribute{
				MarkdownDescription: "Enable port hopping",
				Computed:            true,
			},
			"port_hopping_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"control_session_pps": schema.Int64Attribute{
				MarkdownDescription: "Set the policer rate for control sessions",
				Computed:            true,
			},
			"control_session_pps_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"track_transport": schema.BoolAttribute{
				MarkdownDescription: "Configure tracking of transport",
				Computed:            true,
			},
			"track_transport_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"track_interface_tag": schema.Int64Attribute{
				MarkdownDescription: "OMP Tag attached to routes based on interface tracking",
				Computed:            true,
			},
			"track_interface_tag_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"console_baud_rate": schema.StringAttribute{
				MarkdownDescription: "Set the console baud rate",
				Computed:            true,
			},
			"console_baud_rate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"max_omp_sessions": schema.Int64Attribute{
				MarkdownDescription: "Set the maximum number of OMP sessions <1..100> the device can have",
				Computed:            true,
			},
			"max_omp_sessions_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"multi_tenant": schema.BoolAttribute{
				MarkdownDescription: "Device is multi-tenant",
				Computed:            true,
			},
			"multi_tenant_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"track_default_gateway": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable default gateway tracking",
				Computed:            true,
			},
			"track_default_gateway_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"admin_tech_on_failure": schema.BoolAttribute{
				MarkdownDescription: "Collect admin-tech before reboot due to daemon failure",
				Computed:            true,
			},
			"admin_tech_on_failure_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"idle_timeout": schema.Int64Attribute{
				MarkdownDescription: "Idle CLI timeout in minutes",
				Computed:            true,
			},
			"idle_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"trackers": schema.ListNestedAttribute{
				MarkdownDescription: "Tracker configuration",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Tracker name",
							Computed:            true,
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"endpoint_ip": schema.StringAttribute{
							MarkdownDescription: "IP address of endpoint",
							Computed:            true,
						},
						"endpoint_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"endpoint_dns_name": schema.StringAttribute{
							MarkdownDescription: "DNS name of endpoint",
							Computed:            true,
						},
						"endpoint_dns_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"endpoint_api_url": schema.StringAttribute{
							MarkdownDescription: "API url of endpoint",
							Computed:            true,
						},
						"endpoint_api_url_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"elements": schema.SetAttribute{
							MarkdownDescription: "Tracker member names separated by space",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"elements_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"boolean": schema.StringAttribute{
							MarkdownDescription: "Type of grouping to be performed for tracker group",
							Computed:            true,
						},
						"boolean_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"threshold": schema.Int64Attribute{
							MarkdownDescription: "Probe Timeout threshold <100..1000> milliseconds",
							Computed:            true,
						},
						"threshold_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"interval": schema.Int64Attribute{
							MarkdownDescription: "Probe interval <10..600> seconds",
							Computed:            true,
						},
						"interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"multiplier": schema.Int64Attribute{
							MarkdownDescription: "Probe failure multiplier <1..10> failed attempts",
							Computed:            true,
						},
						"multiplier_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Default(Interface)",
							Computed:            true,
						},
						"type_variable": schema.StringAttribute{
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
			"object_trackers": schema.ListNestedAttribute{
				MarkdownDescription: "Object Track configuration",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"object_number": schema.Int64Attribute{
							MarkdownDescription: "Object tracker ID",
							Computed:            true,
						},
						"object_number_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"interface": schema.StringAttribute{
							MarkdownDescription: "interface name",
							Computed:            true,
						},
						"interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"sig": schema.StringAttribute{
							MarkdownDescription: "service sig",
							Computed:            true,
						},
						"sig_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ip": schema.StringAttribute{
							MarkdownDescription: "IP address of route",
							Computed:            true,
						},
						"ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"mask": schema.StringAttribute{
							MarkdownDescription: "Route Ip Mask",
							Computed:            true,
						},
						"mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "VPN",
							Computed:            true,
						},
						"group_tracks_ids": schema.ListNestedAttribute{
							MarkdownDescription: "Tracks id in group configuration",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"track_id": schema.Int64Attribute{
										MarkdownDescription: "Track id",
										Computed:            true,
									},
									"track_id_variable": schema.StringAttribute{
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
						"boolean": schema.StringAttribute{
							MarkdownDescription: "Type of grouping to be performed for tracker group",
							Computed:            true,
						},
						"boolean_variable": schema.StringAttribute{
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
			"on_demand_tunnel": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable On-demand Tunnel",
				Computed:            true,
			},
			"on_demand_tunnel_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"on_demand_tunnel_idle_timeout": schema.Int64Attribute{
				MarkdownDescription: "Idle CLI timeout in minutes",
				Computed:            true,
			},
			"on_demand_tunnel_idle_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"region_id": schema.Int64Attribute{
				MarkdownDescription: "Set region ID",
				Computed:            true,
			},
			"region_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"secondary_region_id": schema.Int64Attribute{
				MarkdownDescription: "Set secondary region ID",
				Computed:            true,
			},
			"secondary_region_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"role": schema.StringAttribute{
				MarkdownDescription: "Set the role for router",
				Computed:            true,
			},
			"role_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"affinity_group_number": schema.Int64Attribute{
				MarkdownDescription: "Set the affinity group number for router",
				Computed:            true,
			},
			"affinity_group_number_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"affinity_group_preference": schema.SetAttribute{
				MarkdownDescription: "Set the affinity group preference",
				ElementType:         types.Int64Type,
				Computed:            true,
			},
			"affinity_group_preference_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"transport_gateway": schema.BoolAttribute{
				MarkdownDescription: "Enable transport gateway",
				Computed:            true,
			},
			"transport_gateway_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"enable_mrf_migration": schema.StringAttribute{
				MarkdownDescription: "Enable migration mode to Multi-Region Fabric",
				Computed:            true,
			},
			"migration_bgp_community": schema.Int64Attribute{
				MarkdownDescription: "Set BGP community during migration from BGP-core based network",
				Computed:            true,
			},
		},
	}
}

func (d *CiscoSystemFeatureTemplateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *CiscoSystemFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *CiscoSystemFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CiscoSystem

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
