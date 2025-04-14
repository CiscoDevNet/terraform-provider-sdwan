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
	_ datasource.DataSource              = &SystemAAAProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &SystemAAAProfileParcelDataSource{}
)

func NewSystemAAAProfileParcelDataSource() datasource.DataSource {
	return &SystemAAAProfileParcelDataSource{}
}

type SystemAAAProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *SystemAAAProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_aaa_feature"
}

func (d *SystemAAAProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the System AAA Feature.",

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
			"authentication_group": schema.BoolAttribute{
				MarkdownDescription: "Authentication configurations parameters",
				Computed:            true,
			},
			"authentication_group_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"accounting_group": schema.BoolAttribute{
				MarkdownDescription: "Accounting configurations parameters",
				Computed:            true,
			},
			"accounting_group_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"server_auth_order": schema.ListAttribute{
				MarkdownDescription: "ServerGroups priority order",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"users": schema.ListNestedAttribute{
				MarkdownDescription: "Create local login account",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Set the username",
							Computed:            true,
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: "Set the user password",
							Computed:            true,
						},
						"password_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"privilege": schema.StringAttribute{
							MarkdownDescription: "Set Privilege Level for this user",
							Computed:            true,
						},
						"privilege_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"public_keys": schema.ListNestedAttribute{
							MarkdownDescription: "List of RSA public-keys per user",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key_string": schema.StringAttribute{
										MarkdownDescription: "Set the RSA key string",
										Computed:            true,
									},
									"key_type": schema.StringAttribute{
										MarkdownDescription: "Only RSA is supported",
										Computed:            true,
									},
									"key_type_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"radius_groups": schema.ListNestedAttribute{
				MarkdownDescription: "Configure the Radius serverGroup",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_name": schema.StringAttribute{
							MarkdownDescription: "Set Radius server Group Name",
							Computed:            true,
						},
						"vpn": schema.Int64Attribute{
							MarkdownDescription: "Set VPN in which Radius server is located",
							Computed:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: "Set interface to use to reach Radius server",
							Computed:            true,
						},
						"source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"servers": schema.ListNestedAttribute{
							MarkdownDescription: "Configure the Radius server",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: "Set IP address of Radius server",
										Computed:            true,
									},
									"auth_port": schema.Int64Attribute{
										MarkdownDescription: "Set Authentication port to use to connect to Radius server",
										Computed:            true,
									},
									"auth_port_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"acct_port": schema.Int64Attribute{
										MarkdownDescription: "Set Accounting port to use to connect to Radius server",
										Computed:            true,
									},
									"acct_port_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"timeout": schema.Int64Attribute{
										MarkdownDescription: "Configure how long to wait for replies from the Radius server",
										Computed:            true,
									},
									"timeout_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"retransmit": schema.Int64Attribute{
										MarkdownDescription: "Configure how many times to contact this Radius server",
										Computed:            true,
									},
									"retransmit_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"key": schema.StringAttribute{
										MarkdownDescription: "Set the Radius server shared key",
										Computed:            true,
									},
									"secret_key": schema.StringAttribute{
										MarkdownDescription: "Set the Radius server shared type 7 encrypted key",
										Computed:            true,
									},
									"secret_key_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"key_enum": schema.StringAttribute{
										MarkdownDescription: "Type of encyption. To be used for type 6",
										Computed:            true,
									},
									"key_type": schema.StringAttribute{
										MarkdownDescription: "key type",
										Computed:            true,
									},
									"key_type_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"tacacs_groups": schema.ListNestedAttribute{
				MarkdownDescription: "Configure the TACACS serverGroup",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_name": schema.StringAttribute{
							MarkdownDescription: "Set TACACS server Group Name",
							Computed:            true,
						},
						"vpn": schema.Int64Attribute{
							MarkdownDescription: "Set VPN in which TACACS server is located",
							Computed:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: "Set interface to use to reach TACACS server",
							Computed:            true,
						},
						"source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"servers": schema.ListNestedAttribute{
							MarkdownDescription: "Configure the TACACS server",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: "Set IP address of TACACS server",
										Computed:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: "TACACS Port",
										Computed:            true,
									},
									"port_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"timeout": schema.Int64Attribute{
										MarkdownDescription: "Configure how long to wait for replies from the TACACS server",
										Computed:            true,
									},
									"timeout_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"key": schema.StringAttribute{
										MarkdownDescription: "Set the TACACS server shared key",
										Computed:            true,
									},
									"secret_key": schema.StringAttribute{
										MarkdownDescription: "Set the TACACS server shared type 7 encrypted key",
										Computed:            true,
									},
									"secret_key_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"key_enum": schema.StringAttribute{
										MarkdownDescription: "Type of encyption. To be used for type 6",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"accounting_rules": schema.ListNestedAttribute{
				MarkdownDescription: "Configure the accounting rules",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"rule_id": schema.StringAttribute{
							MarkdownDescription: "Configure Accounting Rule ID",
							Computed:            true,
						},
						"method": schema.StringAttribute{
							MarkdownDescription: "Configure Accounting Method",
							Computed:            true,
						},
						"level": schema.StringAttribute{
							MarkdownDescription: "Privilege level when method is commands",
							Computed:            true,
						},
						"start_stop": schema.BoolAttribute{
							MarkdownDescription: "Record start and stop without waiting",
							Computed:            true,
						},
						"start_stop_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"group": schema.SetAttribute{
							MarkdownDescription: "Use Server-group",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
			"authorization_console": schema.BoolAttribute{
				MarkdownDescription: "For enabling console authorization",
				Computed:            true,
			},
			"authorization_console_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"authorization_config_commands": schema.BoolAttribute{
				MarkdownDescription: "For configuration mode commands.",
				Computed:            true,
			},
			"authorization_config_commands_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"authorization_rules": schema.ListNestedAttribute{
				MarkdownDescription: "Configure the Authorization Rules",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"rule_id": schema.StringAttribute{
							MarkdownDescription: "Configure Authorization Rule ID",
							Computed:            true,
						},
						"method": schema.StringAttribute{
							MarkdownDescription: "Method",
							Computed:            true,
						},
						"level": schema.StringAttribute{
							MarkdownDescription: "Privilege level when method is commands",
							Computed:            true,
						},
						"group": schema.SetAttribute{
							MarkdownDescription: "Use Server-group",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"if_authenticated": schema.BoolAttribute{
							MarkdownDescription: "Succeed if user has authenticated",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SystemAAAProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *SystemAAAProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SystemAAA

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
