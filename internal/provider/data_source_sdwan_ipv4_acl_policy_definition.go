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
	_ datasource.DataSource              = &IPv4ACLPolicyDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &IPv4ACLPolicyDefinitionDataSource{}
)

func NewIPv4ACLPolicyDefinitionDataSource() datasource.DataSource {
	return &IPv4ACLPolicyDefinitionDataSource{}
}

type IPv4ACLPolicyDefinitionDataSource struct {
	client *sdwan.Client
}

func (d *IPv4ACLPolicyDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv4_acl_policy_definition"
}

func (d *IPv4ACLPolicyDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the IPv4 ACL Policy Definition .",

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
				MarkdownDescription: "Default action, either `accept` or `drop`",
				Computed:            true,
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: "List of ACL sequences",
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
						"base_action": schema.StringAttribute{
							MarkdownDescription: "Base action, either `accept` or `drop`",
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
									"dscp": schema.StringAttribute{
										MarkdownDescription: "DSCP value",
										Computed:            true,
									},
									"source_ip": schema.StringAttribute{
										MarkdownDescription: "Source IP prefix",
										Computed:            true,
									},
									"source_ip_variable": schema.StringAttribute{
										MarkdownDescription: "Source IP prefix variable",
										Computed:            true,
									},
									"icmp_message": schema.StringAttribute{
										MarkdownDescription: "ICMP Message",
										Computed:            true,
									},
									"destination_ip": schema.StringAttribute{
										MarkdownDescription: "Destination IP prefix",
										Computed:            true,
									},
									"destination_ip_variable": schema.StringAttribute{
										MarkdownDescription: "Destination IP prefix variable",
										Computed:            true,
									},
									"class_map_id": schema.StringAttribute{
										MarkdownDescription: "Class map ID",
										Computed:            true,
									},
									"class_map_version": schema.Int64Attribute{
										MarkdownDescription: "Class map version",
										Computed:            true,
									},
									"packet_length": schema.Int64Attribute{
										MarkdownDescription: "Packet length",
										Computed:            true,
									},
									"priority": schema.StringAttribute{
										MarkdownDescription: "PLP - priority",
										Computed:            true,
									},
									"source_ports": schema.StringAttribute{
										MarkdownDescription: "Source ports. Single value (0-65535) or ranges separated by spaces.",
										Computed:            true,
									},
									"destination_ports": schema.StringAttribute{
										MarkdownDescription: "Destination ports. Single value (0-65535) or ranges separated by spaces.",
										Computed:            true,
									},
									"source_data_ipv4_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "Source data IPv4 prefix list ID",
										Computed:            true,
									},
									"source_data_ipv4_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: "Source data IPv4 prefix list version",
										Computed:            true,
									},
									"destination_data_ipv4_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "Destination data IPv4 prefix list ID",
										Computed:            true,
									},
									"destination_data_ipv4_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: "Destination data IPv4 prefix list version",
										Computed:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: "Single value (0-255) or multiple values separated by spaces",
										Computed:            true,
									},
									"tcp": schema.StringAttribute{
										MarkdownDescription: "TCP parameters",
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
									"class_map_id": schema.StringAttribute{
										MarkdownDescription: "Class map ID",
										Computed:            true,
									},
									"class_map_version": schema.Int64Attribute{
										MarkdownDescription: "Class map version",
										Computed:            true,
									},
									"counter_name": schema.StringAttribute{
										MarkdownDescription: "Counter name",
										Computed:            true,
									},
									"log": schema.BoolAttribute{
										MarkdownDescription: "Enable logging",
										Computed:            true,
									},
									"mirror_id": schema.StringAttribute{
										MarkdownDescription: "Mirror ID",
										Computed:            true,
									},
									"mirror_version": schema.Int64Attribute{
										MarkdownDescription: "Mirror version",
										Computed:            true,
									},
									"policer_id": schema.StringAttribute{
										MarkdownDescription: "Policer ID",
										Computed:            true,
									},
									"policer_version": schema.Int64Attribute{
										MarkdownDescription: "Policer version",
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
												"dscp": schema.Int64Attribute{
													MarkdownDescription: "DSCP value",
													Computed:            true,
												},
												"next_hop": schema.StringAttribute{
													MarkdownDescription: "Next hop IP",
													Computed:            true,
												},
											},
										},
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

func (d *IPv4ACLPolicyDefinitionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *IPv4ACLPolicyDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config IPv4ACLPolicyDefinition

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
	config.Type = types.StringValue("acl")

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
