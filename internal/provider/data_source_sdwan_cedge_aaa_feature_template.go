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
	_ datasource.DataSource              = &CEdgeAAAFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &CEdgeAAAFeatureTemplateDataSource{}
)

func NewCEdgeAAAFeatureTemplateDataSource() datasource.DataSource {
	return &CEdgeAAAFeatureTemplateDataSource{}
}

type CEdgeAAAFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *CEdgeAAAFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cedge_aaa_feature_template"
}

func (d *CEdgeAAAFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the cEdge AAA feature template.",

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
			"dot1x_authentication": schema.BoolAttribute{
				MarkdownDescription: "Authentication configurations parameters",
				Computed:            true,
			},
			"dot1x_authentication_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"dot1x_accounting": schema.BoolAttribute{
				MarkdownDescription: "Accounting configurations parameters",
				Computed:            true,
			},
			"dot1x_accounting_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"server_groups_priority_order": schema.StringAttribute{
				MarkdownDescription: "ServerGroups priority order",
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
						"secret": schema.StringAttribute{
							MarkdownDescription: "Set the user scrypt password/hash",
							Computed:            true,
						},
						"privilege_level": schema.StringAttribute{
							MarkdownDescription: "Set Privilege Level for this user",
							Computed:            true,
						},
						"privilege_level_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ssh_pubkeys": schema.ListNestedAttribute{
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
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Computed:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"radius_server_groups": schema.ListNestedAttribute{
				MarkdownDescription: "Configure the Radius serverGroup",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_name": schema.StringAttribute{
							MarkdownDescription: "Set Radius server Group Name",
							Computed:            true,
						},
						"vpn_id": schema.Int64Attribute{
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
									"authentication_port": schema.Int64Attribute{
										MarkdownDescription: "Set Authentication port to use to connect to Radius server",
										Computed:            true,
									},
									"authentication_port_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"accounting_port": schema.Int64Attribute{
										MarkdownDescription: "Set Accounting port to use to connect to Radius server",
										Computed:            true,
									},
									"accounting_port_variable": schema.StringAttribute{
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
									"encryption_type": schema.StringAttribute{
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
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Computed:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"radius_clients": schema.ListNestedAttribute{
				MarkdownDescription: "Specify a RADIUS client",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"client_ip": schema.StringAttribute{
							MarkdownDescription: "Client IP",
							Computed:            true,
						},
						"client_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"vpn_configurations": schema.ListNestedAttribute{
							MarkdownDescription: "VPN configuration",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"vpn_id": schema.Int64Attribute{
										MarkdownDescription: "VPN ID",
										Computed:            true,
									},
									"vpn_id_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"server_key": schema.StringAttribute{
										MarkdownDescription: "Specify a RADIUS client server-key",
										Computed:            true,
									},
									"server_key_variable": schema.StringAttribute{
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
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"radius_dynamic_author_server_key": schema.StringAttribute{
				MarkdownDescription: "Specify a radius dynamic author server-key",
				Computed:            true,
			},
			"radius_dynamic_author_server_key_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"radius_dynamic_author_domain_stripping": schema.StringAttribute{
				MarkdownDescription: "Domain Stripping",
				Computed:            true,
			},
			"radius_dynamic_author_domain_stripping_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"radius_dynamic_author_authentication_type": schema.StringAttribute{
				MarkdownDescription: "Authentication Type",
				Computed:            true,
			},
			"radius_dynamic_author_authentication_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"radius_dynamic_author_port": schema.Int64Attribute{
				MarkdownDescription: "Specify Radius Dynamic Author Port",
				Computed:            true,
			},
			"radius_dynamic_author_port_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"radius_trustsec_cts_authorization_list": schema.StringAttribute{
				MarkdownDescription: "CTS Authorization List",
				Computed:            true,
			},
			"radius_trustsec_cts_authorization_list_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"radius_trustsec_group": schema.StringAttribute{
				MarkdownDescription: "RADIUS trustsec group",
				Computed:            true,
			},
			"tacacs_server_groups": schema.ListNestedAttribute{
				MarkdownDescription: "Configure the TACACS serverGroup",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_name": schema.StringAttribute{
							MarkdownDescription: "Set TACACS server Group Name",
							Computed:            true,
						},
						"vpn_id": schema.Int64Attribute{
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
									"encryption_type": schema.StringAttribute{
										MarkdownDescription: "Type of encyption. To be used for type 6",
										Computed:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Computed:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"accounting_rules": schema.ListNestedAttribute{
				MarkdownDescription: "Configure the accounting rules",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Configure Accounting Rule ID",
							Computed:            true,
						},
						"method": schema.StringAttribute{
							MarkdownDescription: "Configure Accounting Method",
							Computed:            true,
						},
						"privilege_level": schema.StringAttribute{
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
						"groups": schema.StringAttribute{
							MarkdownDescription: "Comma separated list of groups",
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
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
						"name": schema.StringAttribute{
							MarkdownDescription: "Configure Authorization Rule ID",
							Computed:            true,
						},
						"method": schema.StringAttribute{
							MarkdownDescription: "Method",
							Computed:            true,
						},
						"privilege_level": schema.StringAttribute{
							MarkdownDescription: "Privilege level when method is commands",
							Computed:            true,
						},
						"groups": schema.StringAttribute{
							MarkdownDescription: "Comma separated list of groups",
							Computed:            true,
						},
						"authenticated": schema.BoolAttribute{
							MarkdownDescription: "Succeed if user has authenticated",
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

func (d *CEdgeAAAFeatureTemplateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *CEdgeAAAFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *CEdgeAAAFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CEdgeAAA

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
