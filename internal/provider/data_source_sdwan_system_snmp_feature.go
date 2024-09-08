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
	_ datasource.DataSource              = &SystemSNMPProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &SystemSNMPProfileParcelDataSource{}
)

func NewSystemSNMPProfileParcelDataSource() datasource.DataSource {
	return &SystemSNMPProfileParcelDataSource{}
}

type SystemSNMPProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *SystemSNMPProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_snmp_feature"
}

func (d *SystemSNMPProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the System SNMP Feature.",

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
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable SNMP",
				Computed:            true,
			},
			"shutdown_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"contact_person": schema.StringAttribute{
				MarkdownDescription: "Set the contact for this managed node",
				Computed:            true,
			},
			"contact_person_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"location_of_device": schema.StringAttribute{
				MarkdownDescription: "Set the physical location of this managed node",
				Computed:            true,
			},
			"location_of_device_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"views": schema.ListNestedAttribute{
				MarkdownDescription: "Configure a view record",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Set the name of the SNMP view",
							Computed:            true,
						},
						"oids": schema.ListNestedAttribute{
							MarkdownDescription: "Configure SNMP object identifier",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Configure identifier of subtree of MIB objects",
										Computed:            true,
									},
									"id_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"exclude": schema.BoolAttribute{
										MarkdownDescription: "Exclude the OID",
										Computed:            true,
									},
									"exclude_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"communities": schema.ListNestedAttribute{
				MarkdownDescription: "Configure SNMP community",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Set name of the SNMP community",
							Computed:            true,
						},
						"user_label": schema.StringAttribute{
							MarkdownDescription: "Set user label of the SNMP community",
							Computed:            true,
						},
						"view": schema.StringAttribute{
							MarkdownDescription: "Set name of the SNMP view",
							Computed:            true,
						},
						"view_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"authorization": schema.StringAttribute{
							MarkdownDescription: "Configure access permissions",
							Computed:            true,
						},
						"authorization_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"groups": schema.ListNestedAttribute{
				MarkdownDescription: "Configure an SNMP group",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the SNMP group",
							Computed:            true,
						},
						"security_level": schema.StringAttribute{
							MarkdownDescription: "Configure security level",
							Computed:            true,
						},
						"view": schema.StringAttribute{
							MarkdownDescription: "Name of the SNMP view",
							Computed:            true,
						},
						"view_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"users": schema.ListNestedAttribute{
				MarkdownDescription: "Configure an SNMP user",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the SNMP user",
							Computed:            true,
						},
						"authentication_protocol": schema.StringAttribute{
							MarkdownDescription: "Configure authentication protocol",
							Computed:            true,
						},
						"authentication_protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"authentication_password": schema.StringAttribute{
							MarkdownDescription: "Specify authentication protocol password",
							Computed:            true,
						},
						"authentication_password_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"privacy_protocol": schema.StringAttribute{
							MarkdownDescription: "Configure privacy protocol",
							Computed:            true,
						},
						"privacy_protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"privacy_password": schema.StringAttribute{
							MarkdownDescription: "Specify privacy protocol password",
							Computed:            true,
						},
						"privacy_password_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"group": schema.StringAttribute{
							MarkdownDescription: "Name of the SNMP group",
							Computed:            true,
						},
						"group_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"trap_target_servers": schema.ListNestedAttribute{
				MarkdownDescription: "Configure SNMP server to receive SNMP traps",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "Set VPN in which SNMP server is located",
							Computed:            true,
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ip": schema.StringAttribute{
							MarkdownDescription: "Set IPv4/IPv6 address of SNMP server",
							Computed:            true,
						},
						"ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: "Set UDP port number to connect to SNMP server",
							Computed:            true,
						},
						"port_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"user_label": schema.StringAttribute{
							MarkdownDescription: "Set user label of the SNMP community",
							Computed:            true,
						},
						"user": schema.StringAttribute{
							MarkdownDescription: "Set name of the SNMP user",
							Computed:            true,
						},
						"user_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: "Source interface for outgoing SNMP traps",
							Computed:            true,
						},
						"source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SystemSNMPProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *SystemSNMPProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SystemSNMP

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
