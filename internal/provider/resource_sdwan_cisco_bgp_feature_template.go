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
var _ resource.Resource = &CiscoBGPFeatureTemplateResource{}
var _ resource.ResourceWithImportState = &CiscoBGPFeatureTemplateResource{}

func NewCiscoBGPFeatureTemplateResource() resource.Resource {
	return &CiscoBGPFeatureTemplateResource{}
}

type CiscoBGPFeatureTemplateResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *CiscoBGPFeatureTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_bgp_feature_template"
}

func (r *CiscoBGPFeatureTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Cisco BGP feature template.").AddMinimumVersionDescription("15.0.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the feature template",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the feature template",
				Computed:            true,
			},
			"template_type": schema.StringAttribute{
				MarkdownDescription: "The template type",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the feature template",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the feature template",
				Required:            true,
			},
			"device_types": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of supported device types").AddStringEnumDescription("vedge-C8000V", "vedge-C8300-1N1S-4T2X", "vedge-C8300-1N1S-6T", "vedge-C8300-2N2S-6T", "vedge-C8300-2N2S-4T2X", "vedge-C8500-12X4QC", "vedge-C8500-12X", "vedge-C8500-20X6C", "vedge-C8500L-8S4X", "vedge-C8200-1N-4T", "vedge-C8200L-1N-4T").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"as_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set autonomous system number <1..4294967295> or <XX.YY>").String,
				Optional:            true,
			},
			"as_number_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable BGP").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"shutdown_variable": schema.StringAttribute{
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
			"propagate_aspath": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Propagate AS Path ").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"propagate_aspath_variable": schema.StringAttribute{
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
			"ipv4_route_targets": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Router Target for IPV4").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("VPN ID for IPv4").AddIntegerRangeDescription(1, 65527).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65527),
							},
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"export": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Export Target-VPN community for IPV4").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"asn_ip": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("asn-ip").String,
										Optional:            true,
									},
									"asn_ip_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Optional:            true,
									},
								},
							},
						},
						"import": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Import Target-VPN community for IPV4").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"asn_ip": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("asn-ip").String,
										Optional:            true,
									},
									"asn_ip_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Optional:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"ipv6_route_targets": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Router Target for IPV6").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("VPN ID for IPv6").AddIntegerRangeDescription(1, 65527).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65527),
							},
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"export": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Export Target-VPN community for IPV6").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"asn_ip": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("asn-ip").String,
										Optional:            true,
									},
									"asn_ip_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Optional:            true,
									},
								},
							},
						},
						"import": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Import Target-VPN community for IPV6").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"asn_ip": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("asn-ip").String,
										Optional:            true,
									},
									"asn_ip_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Optional:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"mpls_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MPLS BGP Interface").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface Name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"interface_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"distance_external": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set administrative distance for external BGP routes").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("20").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"distance_external_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"distance_internal": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set administrative distance for internal BGP routes").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("200").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"distance_internal_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"distance_local": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set administrative distance for local BGP routes").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("20").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"distance_local_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"keepalive": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set how often keepalive messages are sent to BGP peer").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("60").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"keepalive_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"holdtime": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the interval when BGP considers a neighbor to be down").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("180").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65535),
				},
			},
			"holdtime_variable": schema.StringAttribute{
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
			"missing_med_worst": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If path has no MED, consider it to be worst path when selecting active BGP paths").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"missing_med_worst_variable": schema.StringAttribute{
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
			"address_families": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set BGP address family").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"family_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set BGP address family").AddStringEnumDescription("ipv4-unicast", "ipv6-unicast").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ipv4-unicast", "ipv6-unicast"),
							},
						},
						"ipv4_aggregate_addresses": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Aggregate prefixes in specific range").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Configure the prefixes to aggregate").String,
										Optional:            true,
									},
									"prefix_variable": schema.StringAttribute{
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
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Optional:            true,
									},
								},
							},
						},
						"ipv6_aggregate_addresses": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 Aggregate prefixes in specific range").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Configure the IPv6 prefixes to aggregate").String,
										Optional:            true,
									},
									"prefix_variable": schema.StringAttribute{
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
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
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
									"prefix": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Configure the prefixes for BGP to announce").String,
										Optional:            true,
									},
									"prefix_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
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
									"prefix": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Configure the prefixes for BGP to announce").String,
										Optional:            true,
									},
									"prefix_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Optional:            true,
									},
								},
							},
						},
						"maximum_paths": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set maximum number of parallel IBGP paths for multipath load sharing").AddIntegerRangeDescription(0, 32).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 32),
							},
						},
						"maximum_paths_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"default_information_originate": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("BGP Default Information Originate").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"default_information_originate_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"table_map_policy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Map external entry attributes into routing table").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"table_map_policy_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"table_map_filter": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Filter").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"table_map_filter_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"redistribute_routes": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Redistribute routes into BGP").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set the protocol to redistribute routes from").AddStringEnumDescription("static", "connected", "ospf", "ospfv3", "omp", "eigrp", "nat").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("static", "connected", "ospf", "ospfv3", "omp", "eigrp", "nat"),
										},
									},
									"protocol_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"route_policy": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Configure policy to apply to prefixes received from BGP neighbor").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(1, 32),
										},
									},
									"route_policy_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Optional:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"ipv4_neighbors": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set BGP neighbors").String,
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
						"remote_as": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set remote autonomous system number").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"remote_as_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"keepalive": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set how often to advertise keepalive messages to BGP peer").AddIntegerRangeDescription(0, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"keepalive_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"holdtime": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set how long to wait since receiving a keepalive message to consider BGP peer unavailable").AddIntegerRangeDescription(0, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"holdtime_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set IP address of interface for TCP connection to BGP neighbor").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"source_interface_variable": schema.StringAttribute{
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
						"send_ext_community": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Send extended community attribute").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"send_ext_community_variable": schema.StringAttribute{
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
						"send_label_explicit": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Send label").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"send_label_explicit_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"as_override": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("As Override").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"as_override_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"allow_as_in": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("As Number").AddIntegerRangeDescription(1, 10).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 10),
							},
						},
						"allow_as_in_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"address_families": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set BGP address family").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"family_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set BGP address family").AddStringEnumDescription("ipv4-unicast", "vpnv4-unicast", "vpnv6-unicast").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("ipv4-unicast", "vpnv4-unicast", "vpnv6-unicast"),
										},
									},
									"maximum_prefixes": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set maximum number of prefixes accepted from BGP peer").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
									"maximum_prefixes_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"maximum_prefixes_threshold": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set threshold at which to generate a warning message").AddIntegerRangeDescription(0, 100).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 100),
										},
									},
									"maximum_prefixes_threshold_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"maximum_prefixes_restart": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set when to restart BGP connection if threshold is exceeded").AddIntegerRangeDescription(0, 65535).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 65535),
										},
									},
									"maximum_prefixes_restart_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"maximum_prefixes_warning_only": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Display only a warning message when threshold is exceeded").AddDefaultValueDescription("false").String,
										Optional:            true,
									},
									"maximum_prefixes_warning_only_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"route_policies": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Select route policy to apply to prefixes received from BGP neighbor").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"direction": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Set direction for applying route policy").AddStringEnumDescription("in", "out").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("in", "out"),
													},
												},
												"policy_name": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Configure name of route policy").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.LengthBetween(1, 32),
													},
												},
												"policy_name_variable": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
													Optional:            true,
												},
												"optional": schema.BoolAttribute{
													MarkdownDescription: "Indicates if list item is considered optional.",
													Optional:            true,
												},
											},
										},
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Optional:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
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
						"remote_as": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set remote autonomous system number").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"remote_as_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"keepalive": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set how often to advertise keepalive messages to BGP peer").AddIntegerRangeDescription(0, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"keepalive_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"holdtime": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set how long to wait since receiving a keepalive message to consider BGP peer unavailable").AddIntegerRangeDescription(0, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"holdtime_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set IP address of interface for TCP connection to BGP neighbor").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"source_interface_variable": schema.StringAttribute{
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
						"send_ext_community": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Send extended community attribute").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"send_ext_community_variable": schema.StringAttribute{
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
						"send_label_explicit": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Send label Explicit").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"send_label_explicit_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"as_override": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("As Override").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"as_override_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"allow_as_in": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription(" As Number").AddIntegerRangeDescription(1, 10).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 10),
							},
						},
						"allow_as_in_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"address_families": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set BGP address family").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"family_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set BGP address family").AddStringEnumDescription("ipv6-unicast").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("ipv6-unicast"),
										},
									},
									"maximum_prefixes": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set maximum number of prefixes accepted from BGP peer").AddIntegerRangeDescription(0, 4294967295).AddDefaultValueDescription("0").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
									"maximum_prefixes_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"maximum_prefixes_threshold": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set threshold at which to generate a warning message").AddIntegerRangeDescription(0, 100).AddDefaultValueDescription("0").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 100),
										},
									},
									"maximum_prefixes_threshold_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"maximum_prefixes_restart": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set when to restart BGP connection if threshold is exceeded").AddIntegerRangeDescription(0, 65535).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 65535),
										},
									},
									"maximum_prefixes_restart_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"maximum_prefixes_warning_only": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Display only a warning message when threshold is exceeded").AddDefaultValueDescription("false").String,
										Optional:            true,
									},
									"maximum_prefixes_warning_only_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"route_policies": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Select route policy to apply to prefixes received from BGP neighbor").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"direction": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Set direction for applying route policy").AddStringEnumDescription("in", "out").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("in", "out"),
													},
												},
												"policy_name": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Configure name of route policy").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.LengthBetween(1, 32),
													},
												},
												"policy_name_variable": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
													Optional:            true,
												},
												"optional": schema.BoolAttribute{
													MarkdownDescription: "Indicates if list item is considered optional.",
													Optional:            true,
												},
											},
										},
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Optional:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *CiscoBGPFeatureTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *CiscoBGPFeatureTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CiscoBGP

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post("/template/feature", body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("templateId").String())
	plan.Version = types.Int64Value(0)
	plan.TemplateType = types.StringValue(plan.getModel())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *CiscoBGPFeatureTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CiscoBGP

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get("/template/feature/object/" + url.QueryEscape(state.Id.ValueString()))
	if res.Get("error.message").String() == "Invalid Template Id" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *CiscoBGPFeatureTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CiscoBGP

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
	r.updateMutex.Lock()
	res, err := r.client.Put("/template/feature/"+url.QueryEscape(plan.Id.ValueString()), body)
	r.updateMutex.Unlock()
	if err != nil {
		if res.Get("error.message").String() == "Template locked in edit mode." {
			resp.Diagnostics.AddWarning("Client Warning", fmt.Sprintf("Failed to modify template due to template being locked by another change. Template changes will not be applied. Re-run 'terraform apply' to try again."))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	if plan.hasChanges(ctx, &state) {
		plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)
	} else {
		plan.Version = state.Version
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *CiscoBGPFeatureTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CiscoBGP

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete("/template/feature/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && res.Get("error.message").String() != "Invalid Template Id" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *CiscoBGPFeatureTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
