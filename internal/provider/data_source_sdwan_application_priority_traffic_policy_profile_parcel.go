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
	_ datasource.DataSource              = &ApplicationPriorityTrafficPolicyProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplicationPriorityTrafficPolicyProfileParcelDataSource{}
)

func NewApplicationPriorityTrafficPolicyProfileParcelDataSource() datasource.DataSource {
	return &ApplicationPriorityTrafficPolicyProfileParcelDataSource{}
}

type ApplicationPriorityTrafficPolicyProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *ApplicationPriorityTrafficPolicyProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_priority_traffic_policy_profile_parcel"
}

func (d *ApplicationPriorityTrafficPolicyProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Application Priority Traffic Policy profile parcel.",

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
			"default_action": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"simple_flow": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"vpn": schema.SetAttribute{
				MarkdownDescription: "",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"target_direction": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: "Traffic policy sequence list",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"sequence_id": schema.Int64Attribute{
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
							MarkdownDescription: "Sequence IP Type",
							Computed:            true,
						},
						"matches": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"application_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"dns_application_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"dns": schema.StringAttribute{
										MarkdownDescription: "Dns",
										Computed:            true,
									},
									"dscp": schema.Int64Attribute{
										MarkdownDescription: "DSCP number",
										Computed:            true,
									},
									"packet_length": schema.StringAttribute{
										MarkdownDescription: "Packet Length",
										Computed:            true,
									},
									"protocol": schema.SetAttribute{
										MarkdownDescription: "protocol (0-255) range or individual number separated by space",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"source_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"source_data_ipv6_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"destination_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"destination_data_ipv6_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"destination_region": schema.StringAttribute{
										MarkdownDescription: "Destination Region",
										Computed:            true,
									},
									"tcp": schema.StringAttribute{
										MarkdownDescription: "TCP States",
										Computed:            true,
									},
									"traffic_to": schema.StringAttribute{
										MarkdownDescription: "Traffic to",
										Computed:            true,
									},
									"traffic_class": schema.StringAttribute{
										MarkdownDescription: "Traffic Class",
										Computed:            true,
									},
								},
							},
						},
						"actions": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"counter": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"log": schema.BoolAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"sla_class": schema.ListNestedAttribute{
										MarkdownDescription: "slaClass",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"sla_class_list_id": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"preferred_color": schema.SetAttribute{
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Computed:            true,
												},
												"preferred_color_group_id": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"strict_drop": schema.BoolAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"fallback_to_best_path": schema.BoolAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"preferred_remote_color": schema.SetAttribute{
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Computed:            true,
												},
												"remote_color_restrict": schema.BoolAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
											},
										},
									},
									"backup_sla_preferred_color": schema.SetAttribute{
										MarkdownDescription: "Backup SLA perferred color",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"sets": schema.ListNestedAttribute{
										MarkdownDescription: "",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"dscp": schema.Int64Attribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"policer_id": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"preferred_color_group_id": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"forwarding_class_id": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"local_tloc_list_color": schema.SetAttribute{
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Computed:            true,
												},
												"local_tloc_restrict": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"local_tloc_list_encapsulation": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"preferred_remote_color_id": schema.SetAttribute{
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Computed:            true,
												},
												"preferred_remote_color_restrict": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"tloc_ip": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"tloc_color": schema.SetAttribute{
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Computed:            true,
												},
												"tloc_encapsulation": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"tloc_list_id": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"service_type": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"service_color": schema.SetAttribute{
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Computed:            true,
												},
												"service_encapsulation": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"service_tloc_ip": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"service_vpn": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"service_tloc_list_id": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"service_chain_type": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"service_chain_vpn": schema.Int64Attribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"service_chain_local": schema.BoolAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"service_chain_fallback_to_routing": schema.BoolAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"service_chain_tloc": schema.SetAttribute{
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Computed:            true,
												},
												"service_chain_encapsulation": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"service_chain_id": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"service_chain_tloc_list_id": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"next_hop": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"next_hop_ipv6": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
												"vpn": schema.StringAttribute{
													MarkdownDescription: "",
													Computed:            true,
												},
											},
										},
									},
									"redirect_dns_field": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"redirect_dns_value": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"tcp_optimization": schema.BoolAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"dre_optimization": schema.BoolAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"service_node_group": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"loss_correction_type": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"loss_correct_fec_threshold": schema.Int64Attribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"cflowd": schema.BoolAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"nat_pool": schema.Int64Attribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"nat_vpn": schema.BoolAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"nat_fallback": schema.BoolAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"nat_bypass": schema.BoolAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"nat_dia_pool": schema.SetAttribute{
										MarkdownDescription: "",
										ElementType:         types.Int64Type,
										Computed:            true,
									},
									"nat_dia_interface": schema.SetAttribute{
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"secure_internet_gateway": schema.BoolAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"fallback_to_routing": schema.BoolAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"secure_service_edge": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"secure_service_edge_instance": schema.StringAttribute{
										MarkdownDescription: "",
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

func (d *ApplicationPriorityTrafficPolicyProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *ApplicationPriorityTrafficPolicyProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ApplicationPriorityTrafficPolicy

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
