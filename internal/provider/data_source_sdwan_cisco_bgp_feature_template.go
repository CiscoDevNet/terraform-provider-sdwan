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
	_ datasource.DataSource              = &CiscoBGPFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &CiscoBGPFeatureTemplateDataSource{}
)

func NewCiscoBGPFeatureTemplateDataSource() datasource.DataSource {
	return &CiscoBGPFeatureTemplateDataSource{}
}

type CiscoBGPFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *CiscoBGPFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_bgp_feature_template"
}

func (d *CiscoBGPFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Cisco BGP feature template.",

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
			"as_number": schema.StringAttribute{
				MarkdownDescription: "Set autonomous system number <1..4294967295> or <XX.YY>",
				Computed:            true,
			},
			"as_number_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable BGP",
				Computed:            true,
			},
			"shutdown_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"router_id": schema.StringAttribute{
				MarkdownDescription: "Configure BGP router identifier",
				Computed:            true,
			},
			"router_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"propagate_aspath": schema.BoolAttribute{
				MarkdownDescription: "Propagate AS Path ",
				Computed:            true,
			},
			"propagate_aspath_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"propagate_community": schema.BoolAttribute{
				MarkdownDescription: "Propagate Community",
				Computed:            true,
			},
			"propagate_community_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_route_targets": schema.ListNestedAttribute{
				MarkdownDescription: "Router Target for IPV4",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "VPN ID for IPv4",
							Computed:            true,
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"export": schema.ListNestedAttribute{
							MarkdownDescription: "Export Target-VPN community for IPV4",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"asn_ip": schema.StringAttribute{
										MarkdownDescription: "asn-ip",
										Computed:            true,
									},
									"asn_ip_variable": schema.StringAttribute{
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
						"import": schema.ListNestedAttribute{
							MarkdownDescription: "Import Target-VPN community for IPV4",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"asn_ip": schema.StringAttribute{
										MarkdownDescription: "asn-ip",
										Computed:            true,
									},
									"asn_ip_variable": schema.StringAttribute{
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
			"ipv6_route_targets": schema.ListNestedAttribute{
				MarkdownDescription: "Router Target for IPV6",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "VPN ID for IPv6",
							Computed:            true,
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"export": schema.ListNestedAttribute{
							MarkdownDescription: "Export Target-VPN community for IPV6",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"asn_ip": schema.StringAttribute{
										MarkdownDescription: "asn-ip",
										Computed:            true,
									},
									"asn_ip_variable": schema.StringAttribute{
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
						"import": schema.ListNestedAttribute{
							MarkdownDescription: "Import Target-VPN community for IPV6",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"asn_ip": schema.StringAttribute{
										MarkdownDescription: "asn-ip",
										Computed:            true,
									},
									"asn_ip_variable": schema.StringAttribute{
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
			"mpls_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "MPLS BGP Interface",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Interface Name",
							Computed:            true,
						},
						"interface_name_variable": schema.StringAttribute{
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
			"distance_external": schema.Int64Attribute{
				MarkdownDescription: "Set administrative distance for external BGP routes",
				Computed:            true,
			},
			"distance_external_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"distance_internal": schema.Int64Attribute{
				MarkdownDescription: "Set administrative distance for internal BGP routes",
				Computed:            true,
			},
			"distance_internal_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"distance_local": schema.Int64Attribute{
				MarkdownDescription: "Set administrative distance for local BGP routes",
				Computed:            true,
			},
			"distance_local_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"keepalive": schema.Int64Attribute{
				MarkdownDescription: "Set how often keepalive messages are sent to BGP peer",
				Computed:            true,
			},
			"keepalive_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"holdtime": schema.Int64Attribute{
				MarkdownDescription: "Set the interval when BGP considers a neighbor to be down",
				Computed:            true,
			},
			"holdtime_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"always_compare_med": schema.BoolAttribute{
				MarkdownDescription: "Compare MEDs from all ASs when selecting active BGP paths",
				Computed:            true,
			},
			"always_compare_med_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"deterministic_med": schema.BoolAttribute{
				MarkdownDescription: "Compare MEDs from all routes from same AS when selecting active BGP paths",
				Computed:            true,
			},
			"deterministic_med_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"missing_med_worst": schema.BoolAttribute{
				MarkdownDescription: "If path has no MED, consider it to be worst path when selecting active BGP paths",
				Computed:            true,
			},
			"missing_med_worst_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"compare_router_id": schema.BoolAttribute{
				MarkdownDescription: "Compare router IDs when selecting active BGP paths",
				Computed:            true,
			},
			"compare_router_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"multipath_relax": schema.BoolAttribute{
				MarkdownDescription: "Ignore AS for multipath selection",
				Computed:            true,
			},
			"multipath_relax_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"address_families": schema.ListNestedAttribute{
				MarkdownDescription: "Set BGP address family",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"family_type": schema.StringAttribute{
							MarkdownDescription: "Set BGP address family",
							Computed:            true,
						},
						"ipv4_aggregate_addresses": schema.ListNestedAttribute{
							MarkdownDescription: "Aggregate prefixes in specific range",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix": schema.StringAttribute{
										MarkdownDescription: "Configure the prefixes to aggregate",
										Computed:            true,
									},
									"prefix_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"as_set_path": schema.BoolAttribute{
										MarkdownDescription: "Set AS set path information",
										Computed:            true,
									},
									"as_set_path_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"summary_only": schema.BoolAttribute{
										MarkdownDescription: "Filter out more specific routes from updates",
										Computed:            true,
									},
									"summary_only_variable": schema.StringAttribute{
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
						"ipv6_aggregate_addresses": schema.ListNestedAttribute{
							MarkdownDescription: "IPv6 Aggregate prefixes in specific range",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix": schema.StringAttribute{
										MarkdownDescription: "Configure the IPv6 prefixes to aggregate",
										Computed:            true,
									},
									"prefix_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"as_set_path": schema.BoolAttribute{
										MarkdownDescription: "Set AS set path information",
										Computed:            true,
									},
									"as_set_path_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"summary_only": schema.BoolAttribute{
										MarkdownDescription: "Filter out more specific routes from updates",
										Computed:            true,
									},
									"summary_only_variable": schema.StringAttribute{
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
						"ipv4_networks": schema.ListNestedAttribute{
							MarkdownDescription: "Configure the networks for BGP to advertise",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix": schema.StringAttribute{
										MarkdownDescription: "Configure the prefixes for BGP to announce",
										Computed:            true,
									},
									"prefix_variable": schema.StringAttribute{
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
						"ipv6_networks": schema.ListNestedAttribute{
							MarkdownDescription: "Configure the networks for BGP to advertise",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix": schema.StringAttribute{
										MarkdownDescription: "Configure the prefixes for BGP to announce",
										Computed:            true,
									},
									"prefix_variable": schema.StringAttribute{
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
						"maximum_paths": schema.Int64Attribute{
							MarkdownDescription: "Set maximum number of parallel IBGP paths for multipath load sharing",
							Computed:            true,
						},
						"maximum_paths_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"default_information_originate": schema.BoolAttribute{
							MarkdownDescription: "BGP Default Information Originate",
							Computed:            true,
						},
						"default_information_originate_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"table_map_policy": schema.StringAttribute{
							MarkdownDescription: "Map external entry attributes into routing table",
							Computed:            true,
						},
						"table_map_policy_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"table_map_filter": schema.BoolAttribute{
							MarkdownDescription: "Filter",
							Computed:            true,
						},
						"table_map_filter_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"redistribute_routes": schema.ListNestedAttribute{
							MarkdownDescription: "Redistribute routes into BGP",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: "Set the protocol to redistribute routes from",
										Computed:            true,
									},
									"protocol_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"route_policy": schema.StringAttribute{
										MarkdownDescription: "Configure policy to apply to prefixes received from BGP neighbor",
										Computed:            true,
									},
									"route_policy_variable": schema.StringAttribute{
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
			"ipv4_neighbors": schema.ListNestedAttribute{
				MarkdownDescription: "Set BGP neighbors",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "Set neighbor address",
							Computed:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Set description",
							Computed:            true,
						},
						"description_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"shutdown": schema.BoolAttribute{
							MarkdownDescription: "Enable or disable a BGP neighbor",
							Computed:            true,
						},
						"shutdown_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"remote_as": schema.StringAttribute{
							MarkdownDescription: "Set remote autonomous system number",
							Computed:            true,
						},
						"remote_as_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"keepalive": schema.Int64Attribute{
							MarkdownDescription: "Set how often to advertise keepalive messages to BGP peer",
							Computed:            true,
						},
						"keepalive_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"holdtime": schema.Int64Attribute{
							MarkdownDescription: "Set how long to wait since receiving a keepalive message to consider BGP peer unavailable",
							Computed:            true,
						},
						"holdtime_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: "Set IP address of interface for TCP connection to BGP neighbor",
							Computed:            true,
						},
						"source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"next_hop_self": schema.BoolAttribute{
							MarkdownDescription: "Set router to be next hop for routes advertised to neighbor",
							Computed:            true,
						},
						"next_hop_self_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"send_community": schema.BoolAttribute{
							MarkdownDescription: "Send community attribute",
							Computed:            true,
						},
						"send_community_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"send_ext_community": schema.BoolAttribute{
							MarkdownDescription: "Send extended community attribute",
							Computed:            true,
						},
						"send_ext_community_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ebgp_multihop": schema.Int64Attribute{
							MarkdownDescription: "Set TTL value for peers that are not directly connected",
							Computed:            true,
						},
						"ebgp_multihop_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: "Set MD5 password on TCP connection with BGP peer",
							Computed:            true,
						},
						"password_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"send_label": schema.BoolAttribute{
							MarkdownDescription: "Send label",
							Computed:            true,
						},
						"send_label_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"send_label_explicit": schema.BoolAttribute{
							MarkdownDescription: "Send label",
							Computed:            true,
						},
						"send_label_explicit_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"as_override": schema.BoolAttribute{
							MarkdownDescription: "As Override",
							Computed:            true,
						},
						"as_override_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"allow_as_in": schema.Int64Attribute{
							MarkdownDescription: "As Number",
							Computed:            true,
						},
						"allow_as_in_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"address_families": schema.ListNestedAttribute{
							MarkdownDescription: "Set BGP address family",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"family_type": schema.StringAttribute{
										MarkdownDescription: "Set BGP address family",
										Computed:            true,
									},
									"maximum_prefixes": schema.Int64Attribute{
										MarkdownDescription: "Set maximum number of prefixes accepted from BGP peer",
										Computed:            true,
									},
									"maximum_prefixes_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"maximum_prefixes_threshold": schema.Int64Attribute{
										MarkdownDescription: "Set threshold at which to generate a warning message",
										Computed:            true,
									},
									"maximum_prefixes_threshold_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"maximum_prefixes_restart": schema.Int64Attribute{
										MarkdownDescription: "Set when to restart BGP connection if threshold is exceeded",
										Computed:            true,
									},
									"maximum_prefixes_restart_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"maximum_prefixes_warning_only": schema.BoolAttribute{
										MarkdownDescription: "Display only a warning message when threshold is exceeded",
										Computed:            true,
									},
									"maximum_prefixes_warning_only_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"route_policies": schema.ListNestedAttribute{
										MarkdownDescription: "Select route policy to apply to prefixes received from BGP neighbor",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"direction": schema.StringAttribute{
													MarkdownDescription: "Set direction for applying route policy",
													Computed:            true,
												},
												"policy_name": schema.StringAttribute{
													MarkdownDescription: "Configure name of route policy",
													Computed:            true,
												},
												"policy_name_variable": schema.StringAttribute{
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
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"ipv6_neighbors": schema.ListNestedAttribute{
				MarkdownDescription: "Set BGP IPv6 neighbors",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "Set neighbor address",
							Computed:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Set description",
							Computed:            true,
						},
						"description_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"shutdown": schema.BoolAttribute{
							MarkdownDescription: "Enable or disable a BGP neighbor",
							Computed:            true,
						},
						"shutdown_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"remote_as": schema.StringAttribute{
							MarkdownDescription: "Set remote autonomous system number",
							Computed:            true,
						},
						"remote_as_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"keepalive": schema.Int64Attribute{
							MarkdownDescription: "Set how often to advertise keepalive messages to BGP peer",
							Computed:            true,
						},
						"keepalive_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"holdtime": schema.Int64Attribute{
							MarkdownDescription: "Set how long to wait since receiving a keepalive message to consider BGP peer unavailable",
							Computed:            true,
						},
						"holdtime_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: "Set IP address of interface for TCP connection to BGP neighbor",
							Computed:            true,
						},
						"source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"next_hop_self": schema.BoolAttribute{
							MarkdownDescription: "Set router to be next hop for routes advertised to neighbor",
							Computed:            true,
						},
						"next_hop_self_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"send_community": schema.BoolAttribute{
							MarkdownDescription: "Send community attribute",
							Computed:            true,
						},
						"send_community_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"send_ext_community": schema.BoolAttribute{
							MarkdownDescription: "Send extended community attribute",
							Computed:            true,
						},
						"send_ext_community_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ebgp_multihop": schema.Int64Attribute{
							MarkdownDescription: "Set TTL value for peers that are not directly connected",
							Computed:            true,
						},
						"ebgp_multihop_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: "Set MD5 password on TCP connection with BGP peer",
							Computed:            true,
						},
						"password_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"send_label": schema.BoolAttribute{
							MarkdownDescription: "Send label",
							Computed:            true,
						},
						"send_label_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"send_label_explicit": schema.BoolAttribute{
							MarkdownDescription: "Send label Explicit",
							Computed:            true,
						},
						"send_label_explicit_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"as_override": schema.BoolAttribute{
							MarkdownDescription: "As Override",
							Computed:            true,
						},
						"as_override_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"allow_as_in": schema.Int64Attribute{
							MarkdownDescription: " As Number",
							Computed:            true,
						},
						"allow_as_in_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"address_families": schema.ListNestedAttribute{
							MarkdownDescription: "Set BGP address family",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"family_type": schema.StringAttribute{
										MarkdownDescription: "Set BGP address family",
										Computed:            true,
									},
									"maximum_prefixes": schema.Int64Attribute{
										MarkdownDescription: "Set maximum number of prefixes accepted from BGP peer",
										Computed:            true,
									},
									"maximum_prefixes_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"maximum_prefixes_threshold": schema.Int64Attribute{
										MarkdownDescription: "Set threshold at which to generate a warning message",
										Computed:            true,
									},
									"maximum_prefixes_threshold_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"maximum_prefixes_restart": schema.Int64Attribute{
										MarkdownDescription: "Set when to restart BGP connection if threshold is exceeded",
										Computed:            true,
									},
									"maximum_prefixes_restart_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"maximum_prefixes_warning_only": schema.BoolAttribute{
										MarkdownDescription: "Display only a warning message when threshold is exceeded",
										Computed:            true,
									},
									"maximum_prefixes_warning_only_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"route_policies": schema.ListNestedAttribute{
										MarkdownDescription: "Select route policy to apply to prefixes received from BGP neighbor",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"direction": schema.StringAttribute{
													MarkdownDescription: "Set direction for applying route policy",
													Computed:            true,
												},
												"policy_name": schema.StringAttribute{
													MarkdownDescription: "Configure name of route policy",
													Computed:            true,
												},
												"policy_name_variable": schema.StringAttribute{
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

func (d *CiscoBGPFeatureTemplateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *CiscoBGPFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *CiscoBGPFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CiscoBGP

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
