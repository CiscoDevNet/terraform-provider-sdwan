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
var _ resource.Resource = &ServiceLANVPNProfileParcelResource{}
var _ resource.ResourceWithImportState = &ServiceLANVPNProfileParcelResource{}

func NewServiceLANVPNProfileParcelResource() resource.Resource {
	return &ServiceLANVPNProfileParcelResource{}
}

type ServiceLANVPNProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *ServiceLANVPNProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_lan_vpn_feature"
}

func (r *ServiceLANVPNProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Service LAN VPN Feature.").AddMinimumVersionDescription("20.12.0").String,

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
				Required:            true,
			},
			"vpn": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("VPN").AddIntegerRangeDescription(1, 65527).AddDefaultValueDescription("0").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65527),
				},
			},
			"vpn_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"config_description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(244),
				},
			},
			"config_description_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"omp_admin_distance_ipv4": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("OMP Admin Distance IPv4").AddIntegerRangeDescription(1, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"omp_admin_distance_ipv4_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"omp_admin_distance_ipv6": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("OMP Admin Distance IPv6").AddIntegerRangeDescription(1, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"omp_admin_distance_ipv6_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"enable_sdwan_remote_access": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable SDWAN Remote Access").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"primary_dns_address_ipv4": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Primary DNS Address (IPv4)").String,
				Optional:            true,
			},
			"primary_dns_address_ipv4_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"secondary_dns_address_ipv4": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secondary DNS Address (IPv4)").String,
				Optional:            true,
			},
			"secondary_dns_address_ipv4_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"primary_dns_address_ipv6": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Primary DNS Address (IPv6)").String,
				Optional:            true,
			},
			"primary_dns_address_ipv6_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"secondary_dns_address_ipv6": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secondary DNS Address (IPv6)").String,
				Optional:            true,
			},
			"secondary_dns_address_ipv6_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"host_mappings": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Hostname").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"host_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"list_of_ips": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of IP").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"list_of_ips_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"advertise_omp_ipv4s": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("OMP Advertise IPv4").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol").AddStringEnumDescription("bgp", "ospf", "ospfv3", "connected", "static", "network", "aggregate", "eigrp", "lisp", "isis").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("bgp", "ospf", "ospfv3", "connected", "static", "network", "aggregate", "eigrp", "lisp", "isis"),
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
						"prefixes": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv4 Prefix List").String,
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
									"aggregate_only": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Aggregate Only").AddDefaultValueDescription("false").String,
										Optional:            true,
									},
									"region": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Applied to Region").AddStringEnumDescription("core-and-access", "core", "access").AddDefaultValueDescription("core-and-access").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("core-and-access", "core", "access"),
										},
									},
									"region_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
			"advertise_omp_ipv6s": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("OMP Advertise IPv6").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol").AddStringEnumDescription("BGP", "OSPF", "Connected", "Static", "Network", "Aggregate").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("BGP", "OSPF", "Connected", "Static", "Network", "Aggregate"),
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
						"protocol_sub_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol Sub Type").AddStringEnumDescription("External").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("External"),
							},
						},
						"protocol_sub_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"prefixes": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 Prefix List").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IPv6 Prefix").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`((^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*(/)(\b([0-9]{1,2}|1[01][0-9]|12[0-8])\b)$))`), ""),
										},
									},
									"prefix_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"aggregate_only": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Aggregate Only").AddDefaultValueDescription("false").String,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
			"ipv4_static_routes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 Static Route").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP Address").String,
							Optional:            true,
						},
						"network_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Subnet Mask").AddStringEnumDescription("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0"),
							},
						},
						"subnet_mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"gateway": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Gateway type").AddStringEnumDescription("nextHop", "null0", "vpn", "dhcp").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("nextHop", "null0", "vpn", "dhcp"),
							},
						},
						"next_hops": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv4 Route Gateway Next Hop, Attribute conditional on `gateway` being equal to `nextHop`").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Address").String,
										Optional:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"administrative_distance": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Administrative distance").AddIntegerRangeDescription(1, 255).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 255),
										},
									},
									"administrative_distance_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
								},
							},
						},
						"next_hop_with_trackers": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv4 Route Gateway Next Hop with Tracker, Attribute conditional on `gateway` being equal to `nextHop`").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Address").String,
										Optional:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"administrative_distance": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Administrative distance").AddIntegerRangeDescription(1, 255).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 255),
										},
									},
									"administrative_distance_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"tracker_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
								},
							},
						},
						"null0": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv4 Route Gateway Next Hop, Attribute conditional on `gateway` being equal to `null0`").String,
							Optional:            true,
						},
						"gateway_dhcp": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv4 Route Gateway DHCP, Attribute conditional on `gateway` being equal to `dhcp`").String,
							Optional:            true,
						},
						"vpn": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv4 Route Gateway VPN, Attribute conditional on `gateway` being equal to `vpn`").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv6_static_routes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 Static Route").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prefix").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`((^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*(/)(\b([0-9]{1,2}|1[01][0-9]|12[0-8])\b)$))`), ""),
							},
						},
						"prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"gateway": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Gateway type").AddStringEnumDescription("nextHop", "null0", "nat").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("nextHop", "null0", "nat"),
							},
						},
						"next_hops": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 Route Gateway Next Hop, Attribute conditional on `gateway` being equal to `nextHop`").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Address").String,
										Optional:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"administrative_distance": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Administrative distance").AddIntegerRangeDescription(1, 254).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 254),
										},
									},
									"administrative_distance_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
								},
							},
						},
						"null0": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 Route Gateway Next Hop, Attribute conditional on `gateway` being equal to `null0`").String,
							Optional:            true,
						},
						"nat": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 Nat, Attribute conditional on `gateway` being equal to `nat`").AddStringEnumDescription("NAT64", "NAT66").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("NAT64", "NAT66"),
							},
						},
						"nat_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `gateway` being equal to `nat`").String,
							Optional:            true,
						},
					},
				},
			},
			"services": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Service").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"service_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Service Type").AddStringEnumDescription("FW", "IDS", "IDP", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "TE", "appqoe").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("FW", "IDS", "IDP", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "TE", "appqoe"),
							},
						},
						"service_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ipv4_addresses": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv4 Addresses (Maximum: 4)").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"ipv4_addresses_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"tracking": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tracking").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"tracking_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"service_routes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Service").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP Address").String,
							Optional:            true,
						},
						"network_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Subnet Mask").AddStringEnumDescription("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0"),
							},
						},
						"subnet_mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"service": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Service").AddStringEnumDescription("SIG").AddDefaultValueDescription("SIG").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("SIG"),
							},
						},
						"service_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vpn": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Service").String,
							Optional:            true,
						},
					},
				},
			},
			"gre_routes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 Static GRE Route").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP Address").String,
							Optional:            true,
						},
						"network_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Subnet Mask").AddStringEnumDescription("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0"),
							},
						},
						"subnet_mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"interface": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vpn": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Service").String,
							Optional:            true,
						},
					},
				},
			},
			"ipsec_routes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 Static IPSEC Route").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP Address").String,
							Optional:            true,
						},
						"network_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Subnet Mask").AddStringEnumDescription("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0"),
							},
						},
						"subnet_mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"interface": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"nat_pools": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("NAT Pool").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"nat_pool_name": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT Pool Name").AddIntegerRangeDescription(1, 32).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 32),
							},
						},
						"nat_pool_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"prefix_length": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT Pool Prefix Length").AddIntegerRangeDescription(1, 32).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 32),
							},
						},
						"prefix_length_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"range_start": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT Pool Range Start").String,
							Optional:            true,
						},
						"range_start_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"range_end": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT Pool Range End").String,
							Optional:            true,
						},
						"range_end_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"overload": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT Overload").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"overload_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT Direction").AddStringEnumDescription("inside", "outside").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("inside", "outside"),
							},
						},
						"direction_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"tracker_object_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
							},
						},
					},
				},
			},
			"nat_port_forwards": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("NAT Port Forward").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"nat_pool_name": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT Pool Name").AddIntegerRangeDescription(1, 32).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 32),
							},
						},
						"nat_pool_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source Port").String,
							Optional:            true,
						},
						"source_port_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"translate_port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Translate Port").String,
							Optional:            true,
						},
						"translate_port_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source IP Address").String,
							Optional:            true,
						},
						"source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"translated_source_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Translated Source IP Address").String,
							Optional:            true,
						},
						"translated_source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol").AddStringEnumDescription("TCP", "UDP").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("TCP", "UDP"),
							},
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"static_nats": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static NAT Rules").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"nat_pool_name": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT Pool Name").AddIntegerRangeDescription(1, 32).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 32),
							},
						},
						"nat_pool_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source IP Address").String,
							Optional:            true,
						},
						"source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"translated_source_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Translated Source IP Address").String,
							Optional:            true,
						},
						"translated_source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"static_nat_direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Static NAT Direction").AddStringEnumDescription("inside", "outside").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("inside", "outside"),
							},
						},
						"static_nat_direction_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"tracker_object_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
							},
						},
					},
				},
			},
			"nat_64_v4_pools": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("NAT64 V4 Pool").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT64 v4 Pool Name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"range_start": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT64 Pool Range Start").String,
							Optional:            true,
						},
						"range_start_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"range_end": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT64 Pool Range End").String,
							Optional:            true,
						},
						"range_end_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"overload": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT64 Overload").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"overload_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"route_leak_from_global_vpns": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable route leaking from Global to Service VPN").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"route_protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Leak Routes of particular protocol from Global to Service VPN").AddStringEnumDescription("static", "connected", "bgp", "ospf").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("static", "connected", "bgp", "ospf"),
							},
						},
						"route_protocol_variable": schema.StringAttribute{
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
						"redistributions": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Redistribute Routes to specific Protocol on Service VPN").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Protocol to restributed leaked routes").AddStringEnumDescription("bgp", "ospf").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("bgp", "ospf"),
										},
									},
									"protocol_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"redistribution_policy_id": schema.StringAttribute{
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
			"route_leak_to_global_vpns": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable route leaking from Service to Global VPN").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"route_protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Leak Routes of particular protocol from Service to Global VPN").AddStringEnumDescription("static", "connected", "bgp", "ospf").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("static", "connected", "bgp", "ospf"),
							},
						},
						"route_protocol_variable": schema.StringAttribute{
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
						"redistributions": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Redistribute Routes to specific Protocol on Global VPN").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Protocol to restributed leaked routes").AddStringEnumDescription("bgp", "ospf").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("bgp", "ospf"),
										},
									},
									"protocol_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"redistribution_policy_id": schema.StringAttribute{
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
			"route_leak_from_other_services": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable route leak from another Service VPN to current Service VPN").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_vpn": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source Service VPN from where route are to be leaked").AddIntegerRangeDescription(1, 65530).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65530),
							},
						},
						"source_vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"route_protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Leak Route of particular protocol from Source Service VPN").AddStringEnumDescription("static", "connected", "bgp", "ospf").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("static", "connected", "bgp", "ospf"),
							},
						},
						"route_protocol_variable": schema.StringAttribute{
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
						"redistributions": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Redistribute Route to specific Protocol on Current Service VPN").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Protocol to restributed leaked routes").AddStringEnumDescription("bgp", "ospf").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("bgp", "ospf"),
										},
									},
									"protocol_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"redistribution_policy_id": schema.StringAttribute{
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
			"ipv4_import_route_targets": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"route_target": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Route target").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+\.[0-9]+)|([0-9]+)|((([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))):[0-9]+$`), ""),
							},
						},
						"route_target_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv4_export_route_targets": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"route_target": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Route target").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+\.[0-9]+)|([0-9]+)|((([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))):[0-9]+$`), ""),
							},
						},
						"route_target_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv6_import_route_targets": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"route_target": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Route target").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+\.[0-9]+)|([0-9]+)|((([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))):[0-9]+$`), ""),
							},
						},
						"route_target_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv6_export_route_targets": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"route_target": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Route target").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+\.[0-9]+)|([0-9]+)|((([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))):[0-9]+$`), ""),
							},
						},
						"route_target_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *ServiceLANVPNProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *ServiceLANVPNProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ServiceLANVPN

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
func (r *ServiceLANVPNProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ServiceLANVPN

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
func (r *ServiceLANVPNProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ServiceLANVPN

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
func (r *ServiceLANVPNProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ServiceLANVPN

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
func (r *ServiceLANVPNProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "service_lan_vpn_feature_id" + ",feature_profile_id"
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
