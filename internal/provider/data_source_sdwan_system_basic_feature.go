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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SystemBasicProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &SystemBasicProfileParcelDataSource{}
)

func NewSystemBasicProfileParcelDataSource() datasource.DataSource {
	return &SystemBasicProfileParcelDataSource{}
}

type SystemBasicProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *SystemBasicProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_basic_feature"
}

func (d *SystemBasicProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the System Basic Feature.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Feature",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Feature",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Feature",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Feature",
				Computed:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: "Feature Profile ID",
				Required:            true,
			},
			"timezone": schema.StringAttribute{
				MarkdownDescription: "Set the timezone",
				Computed:            true,
			},
			"timezone_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"config_description": schema.StringAttribute{
				MarkdownDescription: "Set a text description of the device",
				Computed:            true,
			},
			"config_description_variable": schema.StringAttribute{
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
			"gps_longitude": schema.Float64Attribute{
				MarkdownDescription: "Set the device physical longitude",
				Computed:            true,
			},
			"gps_longitude_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"gps_latitude": schema.Float64Attribute{
				MarkdownDescription: "Set the device physical latitude",
				Computed:            true,
			},
			"gps_latitude_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"gps_geo_fencing_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable Geo fencing",
				Computed:            true,
			},
			"gps_geo_fencing_range": schema.Int64Attribute{
				MarkdownDescription: "Set the device’s geo fencing range",
				Computed:            true,
			},
			"gps_geo_fencing_range_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"gps_sms_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable device’s geo fencing SMS",
				Computed:            true,
			},
			"gps_sms_mobile_numbers": schema.ListNestedAttribute{
				MarkdownDescription: "Set device’s geo fencing SMS phone number",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"number": schema.StringAttribute{
							MarkdownDescription: "Mobile number, ex: 1231234414",
							Computed:            true,
						},
						"number_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"device_groups": schema.SetAttribute{
				MarkdownDescription: "Device groups",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"device_groups_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"controller_groups": schema.SetAttribute{
				MarkdownDescription: "Configure a list of comma-separated controller groups",
				ElementType:         types.Int64Type,
				Computed:            true,
			},
			"controller_groups_variable": schema.StringAttribute{
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
			"on_demand_enable": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable On-demand Tunnel",
				Computed:            true,
			},
			"on_demand_enable_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"on_demand_idle_timeout": schema.Int64Attribute{
				MarkdownDescription: "Set the idle timeout for on-demand tunnels",
				Computed:            true,
			},
			"on_demand_idle_timeout_variable": schema.StringAttribute{
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
			"enhanced_app_aware_routing": schema.StringAttribute{
				MarkdownDescription: "Enable SLA Dampening and Enhanced App Routing.",
				Computed:            true,
			},
			"enhanced_app_aware_routing_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"site_types": schema.SetAttribute{
				MarkdownDescription: "Site Type",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"site_types_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"affinity_group_number": schema.Int64Attribute{
				MarkdownDescription: "Affinity Group Number",
				Computed:            true,
			},
			"affinity_group_number_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"affinity_group_preferences": schema.SetAttribute{
				MarkdownDescription: "Affinity Group Preference",
				ElementType:         types.Int64Type,
				Computed:            true,
			},
			"affinity_group_preferences_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"affinity_preference_auto": schema.BoolAttribute{
				MarkdownDescription: "Affinity Group Preference Auto",
				Computed:            true,
			},
			"affinity_preference_auto_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"affinity_per_vrfs": schema.ListNestedAttribute{
				MarkdownDescription: "Affinity Group Number for VRFs",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"affinity_group_number": schema.Int64Attribute{
							MarkdownDescription: "Affinity Group Number",
							Computed:            true,
						},
						"affinity_group_number_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"vrf_range": schema.StringAttribute{
							MarkdownDescription: "Range of VRFs",
							Computed:            true,
						},
						"vrf_range_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SystemBasicProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *SystemBasicProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SystemBasic

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
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
