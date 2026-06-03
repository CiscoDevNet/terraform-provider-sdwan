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
	_ datasource.DataSource                     = &TopologyCustomControlProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure        = &TopologyCustomControlProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigValidators = &TopologyCustomControlProfileParcelDataSource{}
)

func NewTopologyCustomControlProfileParcelDataSource() datasource.DataSource {
	return &TopologyCustomControlProfileParcelDataSource{}
}

type TopologyCustomControlProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *TopologyCustomControlProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_topology_custom_control_feature"
}

func (d *TopologyCustomControlProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Topology Custom Control Feature.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Feature",
				Optional:            true,
				Computed:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Feature",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Feature",
				Optional:            true,
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
			"target_vpn": schema.SetAttribute{
				MarkdownDescription: "",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"target_role": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"target_level": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"target_inbound_sites": schema.SetAttribute{
				MarkdownDescription: "",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"target_outbound_sites": schema.SetAttribute{
				MarkdownDescription: "",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"target_inbound_regions": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"region": schema.StringAttribute{
							MarkdownDescription: "Region name",
							Computed:            true,
						},
						"sub_regions": schema.SetAttribute{
							MarkdownDescription: "Sub-region list",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
			"target_outbound_regions": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"region": schema.StringAttribute{
							MarkdownDescription: "Region name",
							Computed:            true,
						},
						"sub_regions": schema.SetAttribute{
							MarkdownDescription: "Sub-region list",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: "Sequence list",
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
						"type": schema.StringAttribute{
							MarkdownDescription: "Sequence Type",
							Computed:            true,
						},
						"ip_type": schema.StringAttribute{
							MarkdownDescription: "Sequence IP Type",
							Computed:            true,
						},
						"match_entries": schema.SetNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"color_list_id": schema.StringAttribute{
										MarkdownDescription: "Color list ID",
										Computed:            true,
									},
									"community_list_id": schema.StringAttribute{
										MarkdownDescription: "Community list ID",
										Computed:            true,
									},
									"expanded_community_list_id": schema.StringAttribute{
										MarkdownDescription: "Expanded community list ID",
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
									"site": schema.SetAttribute{
										MarkdownDescription: "Site list",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"match_regions": schema.ListNestedAttribute{
										MarkdownDescription: "Match regions list",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"region": schema.StringAttribute{
													MarkdownDescription: "Region name",
													Computed:            true,
												},
												"sub_regions": schema.SetAttribute{
													MarkdownDescription: "Sub-region list",
													ElementType:         types.StringType,
													Computed:            true,
												},
											},
										},
									},
									"role": schema.StringAttribute{
										MarkdownDescription: "Role",
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
									"vpn": schema.SetAttribute{
										MarkdownDescription: "VPN list",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "Prefix list ID",
										Computed:            true,
									},
									"ipv6_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "IPv6 prefix list ID",
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
								},
							},
						},
						"action_entries": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"export_to_vpn": schema.SetAttribute{
										MarkdownDescription: "Export to VPN list",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"set_parameters": schema.SetNestedAttribute{
										MarkdownDescription: "",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"preference": schema.Int64Attribute{
													MarkdownDescription: "Set preference",
													Computed:            true,
												},
												"omp_tag": schema.Int64Attribute{
													MarkdownDescription: "Set OMP tag",
													Computed:            true,
												},
												"community": schema.StringAttribute{
													MarkdownDescription: "Set community value, e.g. `1000:10000` or `internet` or `local-AS`",
													Computed:            true,
												},
												"community_additive": schema.BoolAttribute{
													MarkdownDescription: "Set community additive",
													Computed:            true,
												},
												"affinity": schema.Int64Attribute{
													MarkdownDescription: "Set affinity",
													Computed:            true,
												},
												"service_type": schema.StringAttribute{
													MarkdownDescription: "Set service type",
													Computed:            true,
												},
												"service_vpn": schema.Int64Attribute{
													MarkdownDescription: "Set service VPN ID",
													Computed:            true,
												},
												"service_tloc_ip": schema.StringAttribute{
													MarkdownDescription: "Set service TLOC IP address",
													Computed:            true,
												},
												"service_tloc_color": schema.StringAttribute{
													MarkdownDescription: "Set service TLOC color",
													Computed:            true,
												},
												"service_tloc_encapsulation": schema.StringAttribute{
													MarkdownDescription: "Set service TLOC encapsulation",
													Computed:            true,
												},
												"service_tloc_list_id": schema.StringAttribute{
													MarkdownDescription: "Set service TLOC list ID",
													Computed:            true,
												},
												"service_chain_type": schema.StringAttribute{
													MarkdownDescription: "Set service chain type",
													Computed:            true,
												},
												"service_chain_vpn": schema.Int64Attribute{
													MarkdownDescription: "Set service chain VPN ID",
													Computed:            true,
												},
												"service_chain_tloc_ip": schema.StringAttribute{
													MarkdownDescription: "Set service chain TLOC IP address",
													Computed:            true,
												},
												"service_chain_tloc_color": schema.StringAttribute{
													MarkdownDescription: "Set service chain TLOC color",
													Computed:            true,
												},
												"service_chain_tloc_encapsulation": schema.StringAttribute{
													MarkdownDescription: "Set service chain TLOC encapsulation",
													Computed:            true,
												},
												"service_chain_tloc_list_id": schema.StringAttribute{
													MarkdownDescription: "Set service chain TLOC list ID",
													Computed:            true,
												},
												"tloc_ip": schema.StringAttribute{
													MarkdownDescription: "Set TLOC IP address",
													Computed:            true,
												},
												"tloc_color": schema.StringAttribute{
													MarkdownDescription: "Set TLOC color",
													Computed:            true,
												},
												"tloc_encapsulation": schema.StringAttribute{
													MarkdownDescription: "Set TLOC encapsulation",
													Computed:            true,
												},
												"tloc_list_id": schema.StringAttribute{
													MarkdownDescription: "Set TLOC list ID",
													Computed:            true,
												},
												"tloc_action": schema.StringAttribute{
													MarkdownDescription: "Set TLOC action",
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

func (d *TopologyCustomControlProfileParcelDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *TopologyCustomControlProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *TopologyCustomControlProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config TopologyCustomControl

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))
	if config.Id.IsNull() && !config.Name.IsNull() {
		// Look up parcel ID by name
		res, err := d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve parcels, got error: %s", err))
			return
		}
		found := false
		res.Get("data").ForEach(func(_, v gjson.Result) bool {
			if v.Get("payload.name").String() == config.Name.ValueString() {
				config.Id = types.StringValue(v.Get("parcelId").String())
				found = true
				return false
			}
			return true
		})
		if !found {
			resp.Diagnostics.AddError("Not Found", fmt.Sprintf("No parcel found with name: %s", config.Name.ValueString()))
			return
		}
	}

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
