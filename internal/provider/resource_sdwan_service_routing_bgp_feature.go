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
	"regexp"
	"strings"
	"sync"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &ServiceRoutingBGPProfileParcelResource{}
var _ resource.ResourceWithImportState = &ServiceRoutingBGPProfileParcelResource{}

func NewServiceRoutingBGPProfileParcelResource() resource.Resource {
	return &ServiceRoutingBGPProfileParcelResource{}
}

type ServiceRoutingBGPProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *ServiceRoutingBGPProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_routing_bgp_feature"
}

func (r *ServiceRoutingBGPProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Service Routing BGP Feature.").AddMinimumVersionDescription("20.12.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Feature",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Feature",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Feature",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Feature",
				Optional:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Feature Profile ID").String,
				Optional:            true,
			},
			"as_number": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set autonomous system number <1..4294967295> or <XX.YY>").String,
				Optional:            true,
			},
			"as_number_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"router_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure BGP router identifier").String,
				Optional:            true,
			},
			"router_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"propagate_as_path": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Propagate AS Path").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"propagate_as_path_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"propagate_community": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Propagate Community").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"propagate_community_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"external_routes_distance": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set administrative distance for external BGP routes").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("20").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"external_routes_distance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"internal_routes_distance": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set administrative distance for internal BGP routes").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("200").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"internal_routes_distance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"local_routes_distance": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set administrative distance for local BGP routes").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("20").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"local_routes_distance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"keepalive_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval (seconds) of keepalive messages sent to its BGP peer").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("60").String,
				Optional:            true,
			},
			"keepalive_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"hold_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval (seconds) not receiving a keepalive message declares a BGP peer down").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("180").String,
				Optional:            true,
			},
			"hold_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"always_compare_med": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Compare MEDs from all ASs when selecting active BGP paths").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"always_compare_med_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"deterministic_med": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Compare MEDs from all routes from same AS when selecting active BGP paths").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"deterministic_med_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"missing_med_as_worst": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If path has no MED, consider it to be worst path when selecting active BGP paths").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"missing_med_as_worst_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"compare_router_id": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Compare router IDs when selecting active BGP paths").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"compare_router_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"multipath_relax": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ignore AS for multipath selection").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"multipath_relax_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv4_neighbors": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set BGP IPv4 neighbors").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set neighbor address").String,
							Optional:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set description").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"description_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"shutdown": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable or disable a BGP neighbor").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"shutdown_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"remote_as": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set remote autonomous system number").String,
							Optional:            true,
						},
						"remote_as_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"local_as": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set local autonomous number,Local-AS cannot have the local BGP protocol AS number or the AS number of the remote peer.The local-as is valid only if the peer is a true eBGP peer. It does not work for two peers in different sub-ASs in a confederation.").String,
							Optional:            true,
						},
						"local_as_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"keepalive_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interval (seconds) of keepalive messages sent to its BGP peer").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("60").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(65535),
							},
						},
						"keepalive_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"hold_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interval (seconds) not receiving a keepalive message declares a BGP peer down").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("180").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(65535),
							},
						},
						"hold_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"update_source_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source interface name for BGP neighbor").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(3, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`(ATM|ATM-ACR|AppGigabitEthernet|AppNav-Compress|AppNav-UnCompress|Async|BD-VIF|BDI|CEM|CEM-ACR|Cellular|Dialer|Embedded-Service-Engine|Ethernet|Ethernet-Internal|FastEthernet|FiftyGigabitEthernet|FiveGigabitEthernet|FortyGigabitEthernet|FourHundredGigE|GMPLS|GigabitEthernet|Group-Async|HundredGigE|L2LISP|LISP|Loopback|MFR|Multilink|Port-channel|SM|Serial|Service-Engine|TenGigabitEthernet|Tunnel|TwentyFiveGigE|TwentyFiveGigabitEthernet|TwoGigabitEthernet|TwoHundredGigE|Vif|Virtual-PPP|Virtual-Template|VirtualPortGroup|Vlan|Wlan-GigabitEthernet|nat64|nat66|ntp|nve|ospfv3|overlay|pseudowire|ucse|vasileft|vasiright|vmi)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
							},
						},
						"update_source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"next_hop_self": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set router to be next hop for routes advertised to neighbor").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"next_hop_self_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"send_community": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Send community attribute").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"send_community_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"send_extended_community": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Send extended community attribute").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"send_extended_community_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ebgp_multihop": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set TTL value for peers that are not directly connected").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("1").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
						"ebgp_multihop_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set MD5 password on TCP connection with BGP peer").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 25),
							},
						},
						"password_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"send_label": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Send label").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"send_label_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"as_override": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Override matching AS-number while sending update").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"as_override_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"allowas_in_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The number of accept as-path with my AS present in it").AddIntegerRangeDescription(1, 10).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 10),
							},
						},
						"allowas_in_number_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"address_families": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set BGP address family").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"family_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set IPv4 unicast address family").String,
										Optional:            true,
									},
									"max_number_of_prefixes": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set maximum number of prefixes accepted from BGP peer").AddIntegerRangeDescription(1, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 4294967295),
										},
									},
									"max_number_of_prefixes_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"threshold": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set threshold(1 to 100) at which to generate a warning message").AddIntegerRangeDescription(1, 100).AddDefaultValueDescription("75").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 100),
										},
									},
									"threshold_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"policy_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Neighbor received maximum prefix policy is disabled.").String,
										Optional:            true,
									},
									"restart_interval": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set the restart interval(minutes) when to restart BGP connection if threshold is exceeded").AddIntegerRangeDescription(1, 65535).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
									},
									"restart_interval_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"in_route_policy_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"out_route_policy_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
								},
							},
						},
					},
				},
			},
			"ipv6_neighbors": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set BGP IPv6 neighbors").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set IPv6 neighbor address").String,
							Optional:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set description").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"description_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"shutdown": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable or disable a BGP neighbor").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"shutdown_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"remote_as": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set remote autonomous system number").String,
							Optional:            true,
						},
						"remote_as_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"local_as": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set local autonomous number,Local-AS cannot have the local BGP protocol AS number or the AS number of the remote peer.The local-as is valid only if the peer is a true eBGP peer. It does not work for two peers in different sub-ASs in a confederation.").String,
							Optional:            true,
						},
						"local_as_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"keepalive_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set how often to advertise keepalive messages to BGP peer").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("60").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(65535),
							},
						},
						"keepalive_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"hold_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set how long to wait since receiving a keepalive message to consider BGP peer unavailable").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("180").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(65535),
							},
						},
						"hold_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"update_source_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source interface name for BGP neighbor").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(3, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`(ATM|ATM-ACR|AppGigabitEthernet|AppNav-Compress|AppNav-UnCompress|Async|BD-VIF|BDI|CEM|CEM-ACR|Cellular|Dialer|Embedded-Service-Engine|Ethernet|Ethernet-Internal|FastEthernet|FiftyGigabitEthernet|FiveGigabitEthernet|FortyGigabitEthernet|FourHundredGigE|GMPLS|GigabitEthernet|Group-Async|HundredGigE|L2LISP|LISP|Loopback|MFR|Multilink|Port-channel|SM|Serial|Service-Engine|TenGigabitEthernet|Tunnel|TwentyFiveGigE|TwentyFiveGigabitEthernet|TwoGigabitEthernet|TwoHundredGigE|Vif|Virtual-PPP|Virtual-Template|VirtualPortGroup|Vlan|Wlan-GigabitEthernet|nat64|nat66|ntp|nve|ospfv3|overlay|pseudowire|ucse|vasileft|vasiright|vmi)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
							},
						},
						"update_source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"next_hop_self": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set router to be next hop for routes advertised to neighbor").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"next_hop_self_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"send_community": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Send community attribute").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"send_community_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"send_extended_community": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Send extended community attribute").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"send_extended_community_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ebgp_multihop": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set TTL value for peers that are not directly connected").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("1").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
						"ebgp_multihop_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"password": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set MD5 password on TCP connection with BGP peer").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 25),
							},
						},
						"password_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"as_override": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Override matching AS-number while sending update").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"as_override_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"allowas_in_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The number of accept as-path with my AS present in it").AddIntegerRangeDescription(1, 10).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 10),
							},
						},
						"allowas_in_number_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"address_families": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set IPv6 BGP address family").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"family_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set IPv6 unicast address family").String,
										Optional:            true,
									},
									"max_number_of_prefixes": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set maximum number of prefixes accepted from BGP peer").AddIntegerRangeDescription(1, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 4294967295),
										},
									},
									"max_number_of_prefixes_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"threshold": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set threshold(1 to 100) at which to generate a warning message").AddIntegerRangeDescription(1, 100).AddDefaultValueDescription("75").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 100),
										},
									},
									"threshold_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"policy_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Neighbor received maximum prefix policy is disabled.").String,
										Optional:            true,
									},
									"restart_interval": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set the restart interval(minutes) when to restart BGP connection if threshold is exceeded").AddIntegerRangeDescription(1, 65535).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
									},
									"restart_interval_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"in_route_policy_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"out_route_policy_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
								},
							},
						},
					},
				},
			},
			"ipv4_aggregate_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Aggregate prefixes in specific range").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
						},
						"network_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0"),
							},
						},
						"subnet_mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"as_set_path": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set AS set path information").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"as_set_path_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"summary_only": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Filter out more specific routes from updates").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"summary_only_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv4_networks": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the networks for BGP to advertise").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
						},
						"network_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0"),
							},
						},
						"subnet_mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv4_eibgp_maximum_paths": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set maximum number of parallel IBGP paths for multipath load sharing").AddIntegerRangeDescription(1, 32).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 32),
				},
			},
			"ipv4_eibgp_maximum_paths_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv4_originate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("BGP Default Information Originate").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ipv4_originate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv4_table_map_route_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
				},
			},
			"ipv4_table_map_filter": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Table map filtered or not").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ipv4_table_map_filter_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv4_redistributes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Redistribute routes into BGP").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the protocol to redistribute routes from").AddStringEnumDescription("static", "connected", "omp", "nat", "ospf", "ospfv3", "eigrp").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("static", "connected", "omp", "nat", "ospf", "ospfv3", "eigrp"),
							},
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"route_policy_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
							},
						},
					},
				},
			},
			"ipv6_aggregate_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 Aggregate prefixes in specific range").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"aggregate_prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure the IPv6 prefixes to aggregate").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`((^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*(/)(\b([0-9]{1,2}|1[01][0-9]|12[0-8])\b)$))`), ""),
							},
						},
						"aggregate_prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"as_set_path": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set AS set path information").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"as_set_path_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"summary_only": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Filter out more specific routes from updates").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"summary_only_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv6_networks": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the networks for BGP to advertise").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure the prefixes for BGP to announce").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`((^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*(/)(\b([0-9]{1,2}|1[01][0-9]|12[0-8])\b)$))`), ""),
							},
						},
						"network_prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv6_eibgp_maximum_paths": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set maximum number of parallel IBGP paths for multipath load sharing").AddIntegerRangeDescription(1, 32).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 32),
				},
			},
			"ipv6_eibgp_maximum_paths_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv6_originate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("BGP Default Information Originate").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ipv6_originate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv6_table_map_route_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
				},
			},
			"ipv6_table_map_filter": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Table map filtered or not").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ipv6_table_map_filter_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv6_redistributes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Redistribute routes into BGP").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the protocol to redistribute routes from").AddStringEnumDescription("static", "connected", "ospf", "omp").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("static", "connected", "ospf", "omp"),
							},
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"route_policy_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
							},
						},
					},
				},
			},
		},
	}
}

func (r *ServiceRoutingBGPProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *ServiceRoutingBGPProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ServiceRoutingBGP

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("parcelId").String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *ServiceRoutingBGPProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ServiceRoutingBGP

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if res.Get("error.message").String() == "Invalid feature Id" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	if imp {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}
	if state.Version.IsNull() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *ServiceRoutingBGPProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ServiceRoutingBGP

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Name.ValueString()))

	body := plan.toBody(ctx)
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *ServiceRoutingBGPProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ServiceRoutingBGP

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && res.Get("error.message").String() != "Invalid Template Id" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *ServiceRoutingBGPProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "service_routing_bgp_feature_id" + ",feature_profile_id"
	if len(parts) != (count + 1) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: %s. Got: %q", pattern, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("feature_profile_id"), parts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
