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
	_ datasource.DataSource              = &HubAndSpokeTopologyPolicyDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &HubAndSpokeTopologyPolicyDefinitionDataSource{}
)

func NewHubAndSpokeTopologyPolicyDefinitionDataSource() datasource.DataSource {
	return &HubAndSpokeTopologyPolicyDefinitionDataSource{}
}

type HubAndSpokeTopologyPolicyDefinitionDataSource struct {
	client *sdwan.Client
}

func (d *HubAndSpokeTopologyPolicyDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hub_and_spoke_topology_policy_definition"
}

func (d *HubAndSpokeTopologyPolicyDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Hub and Spoke Topology Policy Definition .",

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
			"vpn_list_id": schema.StringAttribute{
				MarkdownDescription: "VPN list ID",
				Computed:            true,
			},
			"vpn_list_version": schema.Int64Attribute{
				MarkdownDescription: "VPN list version",
				Computed:            true,
			},
			"topologies": schema.ListNestedAttribute{
				MarkdownDescription: "List of topologies",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Topology name",
							Computed:            true,
						},
						"all_hubs_are_equal": schema.BoolAttribute{
							MarkdownDescription: "All hubs are equal (All Spokes Sites connect to all Hubs)",
							Computed:            true,
						},
						"advertise_hub_tlocs": schema.BoolAttribute{
							MarkdownDescription: "Advertise Hub TLOCs",
							Computed:            true,
						},
						"tloc_list_id": schema.StringAttribute{
							MarkdownDescription: "TLOC list ID (required when `advertise_hub_tlocs` is 'true')",
							Computed:            true,
						},
						"spokes": schema.ListNestedAttribute{
							MarkdownDescription: "List of spokes",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"site_list_id": schema.StringAttribute{
										MarkdownDescription: "Site list ID",
										Computed:            true,
									},
									"site_list_version": schema.Int64Attribute{
										MarkdownDescription: "Site list version",
										Computed:            true,
									},
									"hubs": schema.ListNestedAttribute{
										MarkdownDescription: "List of hubs",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"site_list_id": schema.StringAttribute{
													MarkdownDescription: "Site list ID",
													Computed:            true,
												},
												"site_list_version": schema.Int64Attribute{
													MarkdownDescription: "Site list version",
													Computed:            true,
												},
												"preference": schema.StringAttribute{
													MarkdownDescription: "Preference, multiple of 10 (for example 70, 80, 90, 100). The higher the value the higher the priority of the associated hub (required when `all_hubs_are_equal` is 'false')",
													Computed:            true,
												},
												"ipv4_prefix_list_ids": schema.SetAttribute{
													MarkdownDescription: "List of IPv4 prefix list IDs",
													ElementType:         types.StringType,
													Computed:            true,
												},
												"ipv6_prefix_list_ids": schema.SetAttribute{
													MarkdownDescription: "List of IPv6 prefix list IDs",
													ElementType:         types.StringType,
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

func (d *HubAndSpokeTopologyPolicyDefinitionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *HubAndSpokeTopologyPolicyDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config HubAndSpokeTopologyPolicyDefinition

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
	config.Type = types.StringValue("hubAndSpoke")

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
