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
	_ datasource.DataSource              = &ServiceRoutePolicyProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &ServiceRoutePolicyProfileParcelDataSource{}
)

func NewServiceRoutePolicyProfileParcelDataSource() datasource.DataSource {
	return &ServiceRoutePolicyProfileParcelDataSource{}
}

type ServiceRoutePolicyProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *ServiceRoutePolicyProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_route_policy_feature"
}

func (d *ServiceRoutePolicyProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Service Route Policy Feature.",

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
			"default_action": schema.StringAttribute{
				MarkdownDescription: "Default Action",
				Computed:            true,
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: "Route Policy List",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: "Sequence Id",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Sequence Name",
							Computed:            true,
						},
						"base_action": schema.StringAttribute{
							MarkdownDescription: "Base Action",
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "protocol such as IPV4, IPV6, or BOTH",
							Computed:            true,
						},
						"match_entries": schema.ListNestedAttribute{
							MarkdownDescription: "Define match conditions",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"as_path_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"standard_community_list_criteria": schema.StringAttribute{
										MarkdownDescription: "Select a condition such as OR, AND or EXACT",
										Computed:            true,
									},
									"standard_community_lists": schema.ListNestedAttribute{
										MarkdownDescription: "Select a standard community list",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
											},
										},
									},
									"expanded_community_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"extended_community_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"bgp_local_preference": schema.Int64Attribute{
										MarkdownDescription: "BGP Local Preference",
										Computed:            true,
									},
									"metric": schema.Int64Attribute{
										MarkdownDescription: "Select Metric",
										Computed:            true,
									},
									"omp_tag": schema.Int64Attribute{
										MarkdownDescription: "Select OMP Tag",
										Computed:            true,
									},
									"ospf_tag": schema.Int64Attribute{
										MarkdownDescription: "Select OSPF Tag",
										Computed:            true,
									},
									"ipv4_address_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"ipv4_next_hop_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"ipv6_address_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"ipv6_next_hop_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
						"actions": schema.ListNestedAttribute{
							MarkdownDescription: "Define list of actions",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"as_path_prepend": schema.ListAttribute{
										MarkdownDescription: "",
										ElementType:         types.Int64Type,
										Computed:            true,
									},
									"community_additive": schema.BoolAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"community": schema.SetAttribute{
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"community_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"local_preference": schema.Int64Attribute{
										MarkdownDescription: "Set Local Preference",
										Computed:            true,
									},
									"metric": schema.Int64Attribute{
										MarkdownDescription: "Set Metric",
										Computed:            true,
									},
									"metric_type": schema.StringAttribute{
										MarkdownDescription: "Set Metric Type",
										Computed:            true,
									},
									"omp_tag": schema.Int64Attribute{
										MarkdownDescription: "Set OMP Tag",
										Computed:            true,
									},
									"origin": schema.StringAttribute{
										MarkdownDescription: "Set Origin",
										Computed:            true,
									},
									"ospf_tag": schema.Int64Attribute{
										MarkdownDescription: "Set OSPF Tag",
										Computed:            true,
									},
									"weight": schema.Int64Attribute{
										MarkdownDescription: "Set Weight",
										Computed:            true,
									},
									"ipv4_next_hop": schema.StringAttribute{
										MarkdownDescription: "Set Ipv4 Next Hop",
										Computed:            true,
									},
									"ipv6_next_hop": schema.StringAttribute{
										MarkdownDescription: "Set Ipv6 Next Hop",
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

func (d *ServiceRoutePolicyProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *ServiceRoutePolicyProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ServiceRoutePolicy

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
