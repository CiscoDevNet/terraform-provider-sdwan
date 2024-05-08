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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

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
	_ datasource.DataSource              = &CustomControlTopologyPolicyDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &CustomControlTopologyPolicyDefinitionDataSource{}
)

func NewCustomControlTopologyPolicyDefinitionDataSource() datasource.DataSource {
	return &CustomControlTopologyPolicyDefinitionDataSource{}
}

type CustomControlTopologyPolicyDefinitionDataSource struct {
	client *sdwan.Client
}

func (d *CustomControlTopologyPolicyDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_control_topology_policy_definition"
}

func (d *CustomControlTopologyPolicyDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Custom Control Topology Policy Definition .",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the policy definition",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the policy definition",
				Computed:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: "Default action, either `accept` or `reject`",
				Computed:            true,
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: "List of sequences",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: "Sequence ID",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Sequence name",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Sequence type, either `route` or `tloc`",
							Computed:            true,
						},
						"ip_type": schema.StringAttribute{
							MarkdownDescription: "Sequence IP type, either `ipv4`, `ipv6` or `all`",
							Computed:            true,
						},
						"base_action": schema.StringAttribute{
							MarkdownDescription: "Base action, either `accept` or `reject`",
							Computed:            true,
						},
						"match_entries": schema.ListNestedAttribute{
							MarkdownDescription: "List of match entries",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of match entry",
										Computed:            true,
									},
									"color_list_id": schema.StringAttribute{
										MarkdownDescription: "Color list ID",
										Computed:            true,
									},
									"color_list_version": schema.Int64Attribute{
										MarkdownDescription: "Color list version",
										Computed:            true,
									},
									"community_list_id": schema.StringAttribute{
										MarkdownDescription: "Community list ID",
										Computed:            true,
									},
									"community_list_version": schema.Int64Attribute{
										MarkdownDescription: "Community list version",
										Computed:            true,
									},
									"expanded_community_list_id": schema.StringAttribute{
										MarkdownDescription: "Expanded community list ID",
										Computed:            true,
									},
									"expanded_community_list_version": schema.Int64Attribute{
										MarkdownDescription: "Expanded community list version",
										Computed:            true,
									},
									"omp_tag": schema.Int64Attribute{
										MarkdownDescription: "OMP tag",
										Computed:            true,
									},
									"origin": schema.StringAttribute{
										MarkdownDescription: "Origin",
										Computed:            true,
									},
									"originator": schema.StringAttribute{
										MarkdownDescription: "Originator IP",
										Computed:            true,
									},
									"preference": schema.Int64Attribute{
										MarkdownDescription: "Preference",
										Computed:            true,
									},
									"site_list_id": schema.StringAttribute{
										MarkdownDescription: "Site list ID",
										Computed:            true,
									},
									"site_list_version": schema.Int64Attribute{
										MarkdownDescription: "Site list version",
										Computed:            true,
									},
									"path_type": schema.StringAttribute{
										MarkdownDescription: "Path type",
										Computed:            true,
									},
									"tloc_list_id": schema.StringAttribute{
										MarkdownDescription: "TLOC list ID",
										Computed:            true,
									},
									"tloc_list_version": schema.Int64Attribute{
										MarkdownDescription: "TLOC list version",
										Computed:            true,
									},
									"vpn_list_id": schema.StringAttribute{
										MarkdownDescription: "VPN list ID",
										Computed:            true,
									},
									"vpn_list_version": schema.Int64Attribute{
										MarkdownDescription: "VPN list version",
										Computed:            true,
									},
									"prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "Prefix list ID",
										Computed:            true,
									},
									"prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: "Prefix list version",
										Computed:            true,
									},
									"vpn_id": schema.Int64Attribute{
										MarkdownDescription: "VPN ID",
										Computed:            true,
									},
									"tloc_ip": schema.StringAttribute{
										MarkdownDescription: "TLOC IP address",
										Computed:            true,
									},
									"tloc_color": schema.StringAttribute{
										MarkdownDescription: "TLOC color",
										Computed:            true,
									},
									"tloc_encapsulation": schema.StringAttribute{
										MarkdownDescription: "TLOC encapsulation",
										Computed:            true,
									},
									"site_id": schema.Int64Attribute{
										MarkdownDescription: "Site ID",
										Computed:            true,
									},
									"carrier": schema.StringAttribute{
										MarkdownDescription: "Carrier",
										Computed:            true,
									},
									"domain_id": schema.Int64Attribute{
										MarkdownDescription: "Domain ID",
										Computed:            true,
									},
									"group_id": schema.Int64Attribute{
										MarkdownDescription: "Group ID",
										Computed:            true,
									},
								},
							},
						},
						"action_entries": schema.ListNestedAttribute{
							MarkdownDescription: "List of action entries",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of action entry",
										Computed:            true,
									},
									"set_parameters": schema.ListNestedAttribute{
										MarkdownDescription: "List of set parameters",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: "Type of set parameter",
													Computed:            true,
												},
												"tloc_list_id": schema.StringAttribute{
													MarkdownDescription: "TLOC list ID",
													Computed:            true,
												},
												"tloc_list_version": schema.Int64Attribute{
													MarkdownDescription: "TLOC list version",
													Computed:            true,
												},
												"tloc_ip": schema.StringAttribute{
													MarkdownDescription: "TLOC IP address",
													Computed:            true,
												},
												"tloc_color": schema.StringAttribute{
													MarkdownDescription: "TLOC color",
													Computed:            true,
												},
												"tloc_encapsulation": schema.StringAttribute{
													MarkdownDescription: "TLOC encapsulation",
													Computed:            true,
												},
												"tloc_action": schema.StringAttribute{
													MarkdownDescription: "TLOC action",
													Computed:            true,
												},
												"preference": schema.Int64Attribute{
													MarkdownDescription: "Preference",
													Computed:            true,
												},
												"omp_tag": schema.Int64Attribute{
													MarkdownDescription: "OMP tag",
													Computed:            true,
												},
												"community": schema.StringAttribute{
													MarkdownDescription: "Community value, e.g. `1000:10000` or `internet` or `local-AS`",
													Computed:            true,
												},
												"community_additive": schema.BoolAttribute{
													MarkdownDescription: "Community additive",
													Computed:            true,
												},
												"service_type": schema.StringAttribute{
													MarkdownDescription: "Service type",
													Computed:            true,
												},
												"service_vpn_id": schema.Int64Attribute{
													MarkdownDescription: "Service VPN ID",
													Computed:            true,
												},
												"service_tloc_list_id": schema.StringAttribute{
													MarkdownDescription: "Service TLOC list ID",
													Computed:            true,
												},
												"service_tloc_list_version": schema.Int64Attribute{
													MarkdownDescription: "Service TLOC list version",
													Computed:            true,
												},
												"service_tloc_ip": schema.StringAttribute{
													MarkdownDescription: "Service TLOC IP address",
													Computed:            true,
												},
												"service_tloc_color": schema.StringAttribute{
													MarkdownDescription: "Service TLOC color",
													Computed:            true,
												},
												"service_tloc_encapsulation": schema.StringAttribute{
													MarkdownDescription: "Service TLOC encapsulation",
													Computed:            true,
												},
											},
										},
									},
									"export_to_vpn_list_id": schema.StringAttribute{
										MarkdownDescription: "Export to VPN list ID",
										Computed:            true,
									},
									"export_to_vpn_list_version": schema.Int64Attribute{
										MarkdownDescription: "Export to VPN list version",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *CustomControlTopologyPolicyDefinitionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *CustomControlTopologyPolicyDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CustomControlTopologyPolicyDefinition

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get(config.getPath() + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)
	config.Type = types.StringValue("control")

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
