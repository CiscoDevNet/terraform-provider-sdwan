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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ServiceWirelessLANProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &ServiceWirelessLANProfileParcelDataSource{}
)

func NewServiceWirelessLANProfileParcelDataSource() datasource.DataSource {
	return &ServiceWirelessLANProfileParcelDataSource{}
}

type ServiceWirelessLANProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *ServiceWirelessLANProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_wireless_lan_profile_parcel"
}

func (d *ServiceWirelessLANProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Service Wireless LAN profile parcel.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the profile parcel",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the profile parcel",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the profile parcel",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the profile parcel",
				Computed:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: "Feature Profile ID",
				Required:            true,
			},
			"enable_24g": schema.BoolAttribute{
				MarkdownDescription: "2.4GHz Enabled",
				Computed:            true,
			},
			"enable_24g_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"enable_5g": schema.BoolAttribute{
				MarkdownDescription: "5GHz Enabled",
				Computed:            true,
			},
			"enable_5g_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ssids": schema.ListNestedAttribute{
				MarkdownDescription: "Configure Wi-Fi SSID profile",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ssid_name": schema.StringAttribute{
							MarkdownDescription: "Configure wlan SSID",
							Computed:            true,
						},
						"admin_state": schema.BoolAttribute{
							MarkdownDescription: "Set admin state",
							Computed:            true,
						},
						"admin_state_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"broadcast_ssid": schema.BoolAttribute{
							MarkdownDescription: "Enable broadcast SSID",
							Computed:            true,
						},
						"broadcast_ssid_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"vlan_id": schema.Int64Attribute{
							MarkdownDescription: "Set VLAN ID",
							Computed:            true,
						},
						"vlan_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"radio_type": schema.StringAttribute{
							MarkdownDescription: "Select radio type",
							Computed:            true,
						},
						"radio_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"security_type": schema.StringAttribute{
							MarkdownDescription: "Select security type",
							Computed:            true,
						},
						"radius_server_ip": schema.StringAttribute{
							MarkdownDescription: "Set RADIUS server IP",
							Computed:            true,
						},
						"radius_server_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"radius_server_port": schema.Int64Attribute{
							MarkdownDescription: "Set RADIUS server authentication port",
							Computed:            true,
						},
						"radius_server_port_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"radius_server_secret": schema.StringAttribute{
							MarkdownDescription: "Set RADIUS server shared secret",
							Computed:            true,
						},
						"radius_server_secret_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"passphrase": schema.StringAttribute{
							MarkdownDescription: "Set passphrase",
							Computed:            true,
						},
						"passphrase_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"qos_profile": schema.StringAttribute{
							MarkdownDescription: "Select QoS profile",
							Computed:            true,
						},
						"qos_profile_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"country": schema.StringAttribute{
				MarkdownDescription: "Select country",
				Computed:            true,
			},
			"country_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "Set management username",
				Computed:            true,
			},
			"username_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Set management password,the password must contains characters from all of the following classes,lowercase letters,uppercase letters,digits,and special characters. No character in the password can be repeated more than three times consecutively. The password must not be the same as the associated username or the username reversed. The password must not be cisco,ocsic,or any variant obtained by changing the capitalization of the letters in word cisco. In addition,you can't substitute 1,l,or ! for i,0 for o,$ for s.",
				Computed:            true,
			},
			"password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"me_dynamic_ip_enabled": schema.BoolAttribute{
				MarkdownDescription: "ME management IP dynamic allocated by DHCP",
				Computed:            true,
			},
			"me_ipv4_address": schema.StringAttribute{
				MarkdownDescription: "Set mobile express controller address",
				Computed:            true,
			},
			"me_ipv4_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"subnet_mask": schema.StringAttribute{
				MarkdownDescription: "Set mobile express controller subnet mask",
				Computed:            true,
			},
			"subnet_mask_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"default_gateway": schema.StringAttribute{
				MarkdownDescription: "Set mobile express default gateway",
				Computed:            true,
			},
			"default_gateway_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
		},
	}
}

func (d *ServiceWirelessLANProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *ServiceWirelessLANProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ServiceWirelessLAN

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
