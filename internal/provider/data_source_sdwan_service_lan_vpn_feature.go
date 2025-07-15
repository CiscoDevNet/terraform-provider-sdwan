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
	"github.com/hashicorp/go-version"
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
	_ datasource.DataSource              = &ServiceLANVPNProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &ServiceLANVPNProfileParcelDataSource{}
)

func NewServiceLANVPNProfileParcelDataSource() datasource.DataSource {
	return &ServiceLANVPNProfileParcelDataSource{}
}

type ServiceLANVPNProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *ServiceLANVPNProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_lan_vpn_feature"
}

func (d *ServiceLANVPNProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Service LAN VPN Feature.",

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
			"vpn": schema.Int64Attribute{
				MarkdownDescription: "VPN",
				Computed:            true,
			},
			"vpn_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"config_description": schema.StringAttribute{
				MarkdownDescription: "Name",
				Computed:            true,
			},
			"config_description_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"omp_admin_distance_ipv4": schema.Int64Attribute{
				MarkdownDescription: "OMP Admin Distance IPv4",
				Computed:            true,
			},
			"omp_admin_distance_ipv4_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"omp_admin_distance_ipv6": schema.Int64Attribute{
				MarkdownDescription: "OMP Admin Distance IPv6",
				Computed:            true,
			},
			"omp_admin_distance_ipv6_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"enable_sdwan_remote_access": schema.BoolAttribute{
				MarkdownDescription: "Enable SDWAN Remote Access",
				Computed:            true,
			},
			"primary_dns_address_ipv4": schema.StringAttribute{
				MarkdownDescription: "Primary DNS Address (IPv4)",
				Computed:            true,
			},
			"primary_dns_address_ipv4_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"secondary_dns_address_ipv4": schema.StringAttribute{
				MarkdownDescription: "Secondary DNS Address (IPv4)",
				Computed:            true,
			},
			"secondary_dns_address_ipv4_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"primary_dns_address_ipv6": schema.StringAttribute{
				MarkdownDescription: "Primary DNS Address (IPv6)",
				Computed:            true,
			},
			"primary_dns_address_ipv6_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"secondary_dns_address_ipv6": schema.StringAttribute{
				MarkdownDescription: "Secondary DNS Address (IPv6)",
				Computed:            true,
			},
			"secondary_dns_address_ipv6_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"host_mappings": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host_name": schema.StringAttribute{
							MarkdownDescription: "Hostname",
							Computed:            true,
						},
						"host_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"list_of_ips": schema.SetAttribute{
							MarkdownDescription: "List of IP",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"list_of_ips_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"advertise_omp_ipv4s": schema.ListNestedAttribute{
				MarkdownDescription: "OMP Advertise IPv4",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Protocol",
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
						"prefixes": schema.ListNestedAttribute{
							MarkdownDescription: "IPv4 Prefix List",
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
									"aggregate_only": schema.BoolAttribute{
										MarkdownDescription: "Aggregate Only",
										Computed:            true,
									},
									"region": schema.StringAttribute{
										MarkdownDescription: "Applied to Region",
										Computed:            true,
									},
									"region_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"advertise_omp_ipv6s": schema.ListNestedAttribute{
				MarkdownDescription: "OMP Advertise IPv6",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Protocol",
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
						"protocol_sub_type": schema.StringAttribute{
							MarkdownDescription: "Protocol Sub Type",
							Computed:            true,
						},
						"protocol_sub_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"prefixes": schema.ListNestedAttribute{
							MarkdownDescription: "IPv6 Prefix List",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix": schema.StringAttribute{
										MarkdownDescription: "IPv6 Prefix",
										Computed:            true,
									},
									"prefix_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"aggregate_only": schema.BoolAttribute{
										MarkdownDescription: "Aggregate Only",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"ipv4_static_routes": schema.ListNestedAttribute{
				MarkdownDescription: "IPv4 Static Route",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_address": schema.StringAttribute{
							MarkdownDescription: "IP Address",
							Computed:            true,
						},
						"network_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: "Subnet Mask",
							Computed:            true,
						},
						"subnet_mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"next_hops": schema.ListNestedAttribute{
							MarkdownDescription: "IPv4 Route Gateway Next Hop",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: "Address",
										Computed:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"administrative_distance": schema.Int64Attribute{
										MarkdownDescription: "Administrative distance",
										Computed:            true,
									},
									"administrative_distance_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
								},
							},
						},
						"next_hop_with_trackers": schema.ListNestedAttribute{
							MarkdownDescription: "IPv4 Route Gateway Next Hop with Tracker",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: "Address",
										Computed:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"administrative_distance": schema.Int64Attribute{
										MarkdownDescription: "Administrative distance",
										Computed:            true,
									},
									"administrative_distance_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"tracker_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
						"null0": schema.BoolAttribute{
							MarkdownDescription: "IPv4 Route Gateway Next Hop",
							Computed:            true,
						},
						"gateway_dhcp": schema.BoolAttribute{
							MarkdownDescription: "IPv4 Route Gateway DHCP",
							Computed:            true,
						},
						"vpn": schema.BoolAttribute{
							MarkdownDescription: "IPv4 Route Gateway VPN",
							Computed:            true,
						},
					},
				},
			},
			"ipv6_static_routes": schema.ListNestedAttribute{
				MarkdownDescription: "IPv6 Static Route",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"prefix": schema.StringAttribute{
							MarkdownDescription: "Prefix",
							Computed:            true,
						},
						"prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"next_hops": schema.ListNestedAttribute{
							MarkdownDescription: "IPv6 Route Gateway Next Hop",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: "Address",
										Computed:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"administrative_distance": schema.Int64Attribute{
										MarkdownDescription: "Administrative distance",
										Computed:            true,
									},
									"administrative_distance_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
								},
							},
						},
						"null0": schema.BoolAttribute{
							MarkdownDescription: "IPv6 Route Gateway Next Hop",
							Computed:            true,
						},
						"nat": schema.StringAttribute{
							MarkdownDescription: "IPv6 Nat",
							Computed:            true,
						},
						"nat_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"services": schema.ListNestedAttribute{
				MarkdownDescription: "Service",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"service_type": schema.StringAttribute{
							MarkdownDescription: "Service Type",
							Computed:            true,
						},
						"service_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ipv4_addresses": schema.SetAttribute{
							MarkdownDescription: "IPv4 Addresses (Maximum: 4)",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"ipv4_addresses_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tracking": schema.BoolAttribute{
							MarkdownDescription: "Tracking",
							Computed:            true,
						},
						"tracking_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"service_routes": schema.ListNestedAttribute{
				MarkdownDescription: "Service",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_address": schema.StringAttribute{
							MarkdownDescription: "IP Address",
							Computed:            true,
						},
						"network_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: "Subnet Mask",
							Computed:            true,
						},
						"subnet_mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"service": schema.StringAttribute{
							MarkdownDescription: "Service",
							Computed:            true,
						},
						"service_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"vpn": schema.Int64Attribute{
							MarkdownDescription: "Service",
							Computed:            true,
						},
					},
				},
			},
			"gre_routes": schema.ListNestedAttribute{
				MarkdownDescription: "IPv4 Static GRE Route",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_address": schema.StringAttribute{
							MarkdownDescription: "IP Address",
							Computed:            true,
						},
						"network_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: "Subnet Mask",
							Computed:            true,
						},
						"subnet_mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"interface": schema.SetAttribute{
							MarkdownDescription: "Interface",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"vpn": schema.Int64Attribute{
							MarkdownDescription: "Service",
							Computed:            true,
						},
					},
				},
			},
			"ipsec_routes": schema.ListNestedAttribute{
				MarkdownDescription: "IPv4 Static IPSEC Route",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_address": schema.StringAttribute{
							MarkdownDescription: "IP Address",
							Computed:            true,
						},
						"network_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: "Subnet Mask",
							Computed:            true,
						},
						"subnet_mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"interface": schema.SetAttribute{
							MarkdownDescription: "Interface",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"nat_pools": schema.ListNestedAttribute{
				MarkdownDescription: "NAT Pool",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"nat_pool_name": schema.Int64Attribute{
							MarkdownDescription: "NAT Pool Name",
							Computed:            true,
						},
						"nat_pool_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"prefix_length": schema.Int64Attribute{
							MarkdownDescription: "NAT Pool Prefix Length",
							Computed:            true,
						},
						"prefix_length_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"range_start": schema.StringAttribute{
							MarkdownDescription: "NAT Pool Range Start",
							Computed:            true,
						},
						"range_start_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"range_end": schema.StringAttribute{
							MarkdownDescription: "NAT Pool Range End",
							Computed:            true,
						},
						"range_end_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"overload": schema.BoolAttribute{
							MarkdownDescription: "NAT Overload",
							Computed:            true,
						},
						"overload_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"direction": schema.StringAttribute{
							MarkdownDescription: "NAT Direction",
							Computed:            true,
						},
						"direction_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tracker_object_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
			"nat_port_forwards": schema.ListNestedAttribute{
				MarkdownDescription: "NAT Port Forward",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"nat_pool_name": schema.Int64Attribute{
							MarkdownDescription: "NAT Pool Name",
							Computed:            true,
						},
						"nat_pool_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_port": schema.Int64Attribute{
							MarkdownDescription: "Source Port",
							Computed:            true,
						},
						"source_port_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"translate_port": schema.Int64Attribute{
							MarkdownDescription: "Translate Port",
							Computed:            true,
						},
						"translate_port_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_ip": schema.StringAttribute{
							MarkdownDescription: "Source IP Address",
							Computed:            true,
						},
						"source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"translated_source_ip": schema.StringAttribute{
							MarkdownDescription: "Translated Source IP Address",
							Computed:            true,
						},
						"translated_source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Protocol",
							Computed:            true,
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"static_nats": schema.ListNestedAttribute{
				MarkdownDescription: "Static NAT Rules",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"nat_pool_name": schema.Int64Attribute{
							MarkdownDescription: "NAT Pool Name",
							Computed:            true,
						},
						"nat_pool_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_ip": schema.StringAttribute{
							MarkdownDescription: "Source IP Address",
							Computed:            true,
						},
						"source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"translated_source_ip": schema.StringAttribute{
							MarkdownDescription: "Translated Source IP Address",
							Computed:            true,
						},
						"translated_source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"static_nat_direction": schema.StringAttribute{
							MarkdownDescription: "Static NAT Direction",
							Computed:            true,
						},
						"static_nat_direction_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tracker_object_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
			"nat_64_v4_pools": schema.ListNestedAttribute{
				MarkdownDescription: "NAT64 V4 Pool",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "NAT64 v4 Pool Name",
							Computed:            true,
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"range_start": schema.StringAttribute{
							MarkdownDescription: "NAT64 Pool Range Start",
							Computed:            true,
						},
						"range_start_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"range_end": schema.StringAttribute{
							MarkdownDescription: "NAT64 Pool Range End",
							Computed:            true,
						},
						"range_end_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"overload": schema.BoolAttribute{
							MarkdownDescription: "NAT64 Overload",
							Computed:            true,
						},
						"overload_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"route_leak_from_global_vpns": schema.ListNestedAttribute{
				MarkdownDescription: "Enable route leaking from Global to Service VPN",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"route_protocol": schema.StringAttribute{
							MarkdownDescription: "Leak Routes of particular protocol from Global to Service VPN",
							Computed:            true,
						},
						"route_protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"route_policy_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"redistributions": schema.ListNestedAttribute{
							MarkdownDescription: "Redistribute Routes to specific Protocol on Service VPN",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: "Protocol to restributed leaked routes",
										Computed:            true,
									},
									"protocol_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"redistribution_policy_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"route_leak_to_global_vpns": schema.ListNestedAttribute{
				MarkdownDescription: "Enable route leaking from Service to Global VPN",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"route_protocol": schema.StringAttribute{
							MarkdownDescription: "Leak Routes of particular protocol from Service to Global VPN",
							Computed:            true,
						},
						"route_protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"route_policy_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"redistributions": schema.ListNestedAttribute{
							MarkdownDescription: "Redistribute Routes to specific Protocol on Global VPN",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: "Protocol to restributed leaked routes",
										Computed:            true,
									},
									"protocol_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"redistribution_policy_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"route_leak_from_other_services": schema.ListNestedAttribute{
				MarkdownDescription: "Enable route leak from another Service VPN to current Service VPN",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_vpn": schema.Int64Attribute{
							MarkdownDescription: "Source Service VPN from where route are to be leaked",
							Computed:            true,
						},
						"source_vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"route_protocol": schema.StringAttribute{
							MarkdownDescription: "Leak Route of particular protocol from Source Service VPN",
							Computed:            true,
						},
						"route_protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"route_policy_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"redistributions": schema.ListNestedAttribute{
							MarkdownDescription: "Redistribute Route to specific Protocol on Current Service VPN",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: "Protocol to restributed leaked routes",
										Computed:            true,
									},
									"protocol_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"redistribution_policy_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"ipv4_import_route_targets": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"route_target": schema.StringAttribute{
							MarkdownDescription: "Route target",
							Computed:            true,
						},
						"route_target_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"ipv4_export_route_targets": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"route_target": schema.StringAttribute{
							MarkdownDescription: "Route target",
							Computed:            true,
						},
						"route_target_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"ipv6_import_route_targets": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"route_target": schema.StringAttribute{
							MarkdownDescription: "Route target",
							Computed:            true,
						},
						"route_target_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"ipv6_export_route_targets": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"route_target": schema.StringAttribute{
							MarkdownDescription: "Route target",
							Computed:            true,
						},
						"route_target_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *ServiceLANVPNProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

func (d *ServiceLANVPNProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ServiceLANVPN

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	// Get Manager Version
	d.client.Authenticate()
	currentVersion := version.Must(version.NewVersion(d.client.ManagerVersion))

	res, err := d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res, currentVersion)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
