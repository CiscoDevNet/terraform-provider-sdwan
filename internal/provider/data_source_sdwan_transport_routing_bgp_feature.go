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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &TransportRoutingBGPProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &TransportRoutingBGPProfileParcelDataSource{}
)

func NewTransportRoutingBGPProfileParcelDataSource() datasource.DataSource {
	return &TransportRoutingBGPProfileParcelDataSource{}
}

type TransportRoutingBGPProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *TransportRoutingBGPProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transport_routing_bgp_feature"
}

func (d *TransportRoutingBGPProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Transport Routing BGP Feature.",

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
			"as_number": schema.Int64Attribute{
				MarkdownDescription: "Set autonomous system number <1..4294967295> or <XX.YY>",
				Computed:            true,
			},
			"as_number_variable": schema.StringAttribute{
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
			"propagate_as_path": schema.BoolAttribute{
				MarkdownDescription: "Propagate AS Path",
				Computed:            true,
			},
			"propagate_as_path_variable": schema.StringAttribute{
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
			"external_routes_distance": schema.Int64Attribute{
				MarkdownDescription: "Set administrative distance for external BGP routes",
				Computed:            true,
			},
			"external_routes_distance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"internal_routes_distance": schema.Int64Attribute{
				MarkdownDescription: "Set administrative distance for internal BGP routes",
				Computed:            true,
			},
			"internal_routes_distance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"local_routes_distance": schema.Int64Attribute{
				MarkdownDescription: "Set administrative distance for local BGP routes",
				Computed:            true,
			},
			"local_routes_distance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"keepalive_time": schema.Int64Attribute{
				MarkdownDescription: "Interval (seconds) of keepalive messages sent to its BGP peer",
				Computed:            true,
			},
			"keepalive_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"hold_time": schema.Int64Attribute{
				MarkdownDescription: "Interval (seconds) not receiving a keepalive message declares a BGP peer down",
				Computed:            true,
			},
			"hold_time_variable": schema.StringAttribute{
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
			"missing_med_as_worst": schema.BoolAttribute{
				MarkdownDescription: "If path has no MED, consider it to be worst path when selecting active BGP paths",
				Computed:            true,
			},
			"missing_med_as_worst_variable": schema.StringAttribute{
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
			"ipv4_neighbors": schema.ListNestedAttribute{
				MarkdownDescription: "Set BGP IPv4 neighbors",
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
						"remote_as": schema.Int64Attribute{
							MarkdownDescription: "Set remote autonomous system number",
							Computed:            true,
						},
						"remote_as_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"local_as": schema.Int64Attribute{
							MarkdownDescription: "Set local autonomous number,Local-AS cannot have the local BGP protocol AS number or the AS number of the remote peer.The local-as is valid only if the peer is a true eBGP peer. It does not work for two peers in different sub-ASs in a confederation.",
							Computed:            true,
						},
						"local_as_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"keepalive_time": schema.Int64Attribute{
							MarkdownDescription: "Set how often to advertise keepalive messages to BGP peer",
							Computed:            true,
						},
						"keepalive_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"hold_time": schema.Int64Attribute{
							MarkdownDescription: "Set how long to wait since receiving a keepalive message to consider BGP peer unavailable",
							Computed:            true,
						},
						"hold_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"update_source_interface": schema.StringAttribute{
							MarkdownDescription: "Source interface name for BGP neighbor",
							Computed:            true,
						},
						"update_source_interface_variable": schema.StringAttribute{
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
						"send_extended_community": schema.BoolAttribute{
							MarkdownDescription: "Send extended community attribute",
							Computed:            true,
						},
						"send_extended_community_variable": schema.StringAttribute{
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
						"explicit_null": schema.BoolAttribute{
							MarkdownDescription: "Send explicit null label",
							Computed:            true,
						},
						"explicit_null_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"as_override": schema.BoolAttribute{
							MarkdownDescription: "Override matching AS-number while sending update",
							Computed:            true,
						},
						"as_override_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"allowas_in_number": schema.Int64Attribute{
							MarkdownDescription: "The number of accept as-path with my AS present in it",
							Computed:            true,
						},
						"allowas_in_number_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"address_families": schema.ListNestedAttribute{
							MarkdownDescription: "Set BGP address family",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"family_type": schema.StringAttribute{
										MarkdownDescription: "Set IPv4 unicast address family",
										Computed:            true,
									},
									"policy_type": schema.StringAttribute{
										MarkdownDescription: "Neighbor received maximum prefix policy is disabled.",
										Computed:            true,
									},
									"restart_max_number_of_prefixes": schema.Int64Attribute{
										MarkdownDescription: "Set maximum number of prefixes accepted from BGP peer",
										Computed:            true,
									},
									"restart_max_number_of_prefixes_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"restart_threshold": schema.Int64Attribute{
										MarkdownDescription: "Set threshold(1 to 100) at which to generate a warning message",
										Computed:            true,
									},
									"restart_threshold_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"restart_interval": schema.Int64Attribute{
										MarkdownDescription: "Set the restart interval(minutes) when to restart BGP connection if threshold is exceeded",
										Computed:            true,
									},
									"restart_interval_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"warning_message_max_number_of_prefixes": schema.Int64Attribute{
										MarkdownDescription: "Set maximum number of prefixes accepted from BGP peer",
										Computed:            true,
									},
									"warning_message_max_number_of_prefixes_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"warning_message_threshold": schema.Int64Attribute{
										MarkdownDescription: "Set threshold(1 to 100) at which to generate a warning message",
										Computed:            true,
									},
									"warning_message_threshold_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"disable_peer_max_number_of_prefixes": schema.Int64Attribute{
										MarkdownDescription: "Set maximum number of prefixes accepted from BGP peer",
										Computed:            true,
									},
									"disable_peer_max_number_of_prefixes_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"disable_peer_threshold": schema.Int64Attribute{
										MarkdownDescription: "Set threshold(1 to 100) at which to generate a warning message",
										Computed:            true,
									},
									"disable_peer_threshold_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"in_route_policy_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"out_route_policy_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
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
							MarkdownDescription: "Set IPv6 neighbor address",
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
						"remote_as": schema.Int64Attribute{
							MarkdownDescription: "Set remote autonomous system number",
							Computed:            true,
						},
						"remote_as_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"local_as": schema.Int64Attribute{
							MarkdownDescription: "Set local autonomous system number,Local-AS cannot have the local BGP protocol AS number or the AS number of the remote peer.The local-as is valid only if the peer is a true eBGP peer. It does not work for two peers in different sub-ASs in a confederation.",
							Computed:            true,
						},
						"local_as_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"keepalive_time": schema.Int64Attribute{
							MarkdownDescription: "Interval (seconds) of keepalive messages sent to its BGP peer",
							Computed:            true,
						},
						"keepalive_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"hold_time": schema.Int64Attribute{
							MarkdownDescription: "Interval (seconds) not receiving a keepalive message declares a BGP peer down",
							Computed:            true,
						},
						"hold_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"update_source_interface": schema.StringAttribute{
							MarkdownDescription: "Source interface name for BGP neighbor",
							Computed:            true,
						},
						"update_source_interface_variable": schema.StringAttribute{
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
						"send_extended_community": schema.BoolAttribute{
							MarkdownDescription: "Send extended community attribute",
							Computed:            true,
						},
						"send_extended_community_variable": schema.StringAttribute{
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
						"as_override": schema.BoolAttribute{
							MarkdownDescription: "Override matching AS-number while sending update",
							Computed:            true,
						},
						"as_override_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"allowas_in_number": schema.Int64Attribute{
							MarkdownDescription: "The number of accept as-path with my AS present in it",
							Computed:            true,
						},
						"allowas_in_number_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"address_families": schema.ListNestedAttribute{
							MarkdownDescription: "Set IPv6 BGP address family",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"family_type": schema.StringAttribute{
										MarkdownDescription: "Set IPv6 unicast address family",
										Computed:            true,
									},
									"max_number_of_prefixes": schema.Int64Attribute{
										MarkdownDescription: "Set maximum number of prefixes accepted from BGP peer",
										Computed:            true,
									},
									"max_number_of_prefixes_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"threshold": schema.Int64Attribute{
										MarkdownDescription: "Set threshold(1 to 100) at which to generate a warning message",
										Computed:            true,
									},
									"threshold_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"policy_type": schema.StringAttribute{
										MarkdownDescription: "Neighbor received maximum prefix policy is disabled.",
										Computed:            true,
									},
									"restart_interval": schema.Int64Attribute{
										MarkdownDescription: "Set the restart interval(minutes) when to restart BGP connection if threshold is exceeded",
										Computed:            true,
									},
									"restart_interval_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"in_route_policy_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"out_route_policy_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"ipv4_aggregate_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Aggregate prefixes in specific range",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_address": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"network_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"subnet_mask_variable": schema.StringAttribute{
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
					},
				},
			},
			"ipv4_networks": schema.ListNestedAttribute{
				MarkdownDescription: "Configure the networks for BGP to advertise",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_address": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"network_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"subnet_mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"ipv4_eibgp_maximum_paths": schema.Int64Attribute{
				MarkdownDescription: "Set maximum number of parallel IBGP paths for multipath load sharing",
				Computed:            true,
			},
			"ipv4_eibgp_maximum_paths_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_originate": schema.BoolAttribute{
				MarkdownDescription: "BGP Default Information Originate",
				Computed:            true,
			},
			"ipv4_originate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_table_map_route_policy_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ipv4_table_map_filter": schema.BoolAttribute{
				MarkdownDescription: "Table map filtered or not",
				Computed:            true,
			},
			"ipv4_table_map_filter_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_redistributes": schema.ListNestedAttribute{
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
						"route_policy_id": schema.StringAttribute{
							MarkdownDescription: "",
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
						"aggregate_prefix": schema.StringAttribute{
							MarkdownDescription: "Configure the IPv6 prefixes to aggregate",
							Computed:            true,
						},
						"aggregate_prefix_variable": schema.StringAttribute{
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
					},
				},
			},
			"ipv6_networks": schema.ListNestedAttribute{
				MarkdownDescription: "Configure the networks for BGP to advertise",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_prefix": schema.StringAttribute{
							MarkdownDescription: "Configure the prefixes for BGP to announce",
							Computed:            true,
						},
						"network_prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"ipv6_eibgp_maximum_paths": schema.Int64Attribute{
				MarkdownDescription: "Set maximum number of parallel IBGP paths for multipath load sharing",
				Computed:            true,
			},
			"ipv6_eibgp_maximum_paths_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_originate": schema.BoolAttribute{
				MarkdownDescription: "BGP Default Information Originate",
				Computed:            true,
			},
			"ipv6_originate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_table_map_route_policy_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ipv6_table_map_filter": schema.BoolAttribute{
				MarkdownDescription: "Table map filtered or not",
				Computed:            true,
			},
			"ipv6_table_map_filter_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_redistributes": schema.ListNestedAttribute{
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
						"route_policy_id": schema.StringAttribute{
							MarkdownDescription: "",
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
					},
				},
			},
		},
	}
}

func (d *TransportRoutingBGPProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *TransportRoutingBGPProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config TransportRoutingBGP

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
