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
	_ datasource.DataSource              = &RoutePolicyDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &RoutePolicyDefinitionDataSource{}
)

func NewRoutePolicyDefinitionDataSource() datasource.DataSource {
	return &RoutePolicyDefinitionDataSource{}
}

type RoutePolicyDefinitionDataSource struct {
	client *sdwan.Client
}

func (d *RoutePolicyDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route_policy_definition"
}

func (d *RoutePolicyDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Route Policy Definition .",

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
				MarkdownDescription: "List of ACL sequences",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: "Sequence ID",
							Computed:            true,
						},
						"ip_type": schema.StringAttribute{
							MarkdownDescription: "IP version, either `ipv4` or `ipv6`",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Sequence name",
							Computed:            true,
						},
						"base_action": schema.StringAttribute{
							MarkdownDescription: "Base action, either `accept` or `reject`",
							Computed:            true,
						},
						"match_entries": schema.SetNestedAttribute{
							MarkdownDescription: "List of match entries",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of match entry",
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
									"as_path_list_id": schema.StringAttribute{
										MarkdownDescription: "AS path list ID",
										Computed:            true,
									},
									"as_path_list_version": schema.Int64Attribute{
										MarkdownDescription: "AS path list version",
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
									"community_list_match_flag": schema.StringAttribute{
										MarkdownDescription: "Community list match flag",
										Computed:            true,
									},
									"community_list_ids": schema.SetAttribute{
										MarkdownDescription: "Community list IDs",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"community_list_versions": schema.ListAttribute{
										MarkdownDescription: "Community list versions",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"expanded_community_list_id": schema.StringAttribute{
										MarkdownDescription: "Expanded community list ID",
										Computed:            true,
									},
									"expanded_community_list_variable": schema.StringAttribute{
										MarkdownDescription: "Expanded community list variable",
										Computed:            true,
									},
									"expanded_community_list_version": schema.Int64Attribute{
										MarkdownDescription: "Expanded community list version",
										Computed:            true,
									},
									"extended_community_list_id": schema.StringAttribute{
										MarkdownDescription: "Extended community list ID",
										Computed:            true,
									},
									"extended_community_list_version": schema.Int64Attribute{
										MarkdownDescription: "Extended community list version",
										Computed:            true,
									},
									"local_preference": schema.Int64Attribute{
										MarkdownDescription: "Local preference",
										Computed:            true,
									},
									"metric": schema.Int64Attribute{
										MarkdownDescription: "Metric",
										Computed:            true,
									},
									"next_hop_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "Next hop prefix list ID",
										Computed:            true,
									},
									"next_hop_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: "Next hop prefix list version",
										Computed:            true,
									},
									"origin": schema.StringAttribute{
										MarkdownDescription: "Origin",
										Computed:            true,
									},
									"peer": schema.StringAttribute{
										MarkdownDescription: "Peer IP",
										Computed:            true,
									},
									"omp_tag": schema.Int64Attribute{
										MarkdownDescription: "OMP tag",
										Computed:            true,
									},
									"ospf_tag": schema.Int64Attribute{
										MarkdownDescription: "OSPF tag",
										Computed:            true,
									},
								},
							},
						},
						"action_entries": schema.SetNestedAttribute{
							MarkdownDescription: "List of action entries",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of action entry",
										Computed:            true,
									},
									"aggregator": schema.Int64Attribute{
										MarkdownDescription: "Aggregator",
										Computed:            true,
									},
									"aggregator_ip_address": schema.StringAttribute{
										MarkdownDescription: "IP address",
										Computed:            true,
									},
									"as_path_prepend": schema.StringAttribute{
										MarkdownDescription: "Space separated list of ASN to prepend",
										Computed:            true,
									},
									"as_path_exclude": schema.StringAttribute{
										MarkdownDescription: "Space separated list of ASN to exclude",
										Computed:            true,
									},
									"atomic_aggregate": schema.BoolAttribute{
										MarkdownDescription: "Atomic aggregate",
										Computed:            true,
									},
									"community": schema.StringAttribute{
										MarkdownDescription: "Community value, e.g. `1000:10000` or `internet` or `local-AS`",
										Computed:            true,
									},
									"community_variable": schema.StringAttribute{
										MarkdownDescription: "Community variable",
										Computed:            true,
									},
									"community_additive": schema.BoolAttribute{
										MarkdownDescription: "Community additive",
										Computed:            true,
									},
									"local_preference": schema.Int64Attribute{
										MarkdownDescription: "Local preference",
										Computed:            true,
									},
									"metric": schema.Int64Attribute{
										MarkdownDescription: "Metric",
										Computed:            true,
									},
									"weight": schema.Int64Attribute{
										MarkdownDescription: "Weight",
										Computed:            true,
									},
									"metric_type": schema.StringAttribute{
										MarkdownDescription: "Metric type",
										Computed:            true,
									},
									"next_hop": schema.StringAttribute{
										MarkdownDescription: "Next hop IP",
										Computed:            true,
									},
									"omp_tag": schema.Int64Attribute{
										MarkdownDescription: "OMP tag",
										Computed:            true,
									},
									"ospf_tag": schema.Int64Attribute{
										MarkdownDescription: "OSPF tag",
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *RoutePolicyDefinitionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *RoutePolicyDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RoutePolicyDefinition

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
	config.Type = types.StringValue("vedgeRoute")

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
