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
	_ datasource.DataSource              = &CiscoVPNFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &CiscoVPNFeatureTemplateDataSource{}
)

func NewCiscoVPNFeatureTemplateDataSource() datasource.DataSource {
	return &CiscoVPNFeatureTemplateDataSource{}
}

type CiscoVPNFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *CiscoVPNFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_vpn_feature_template"
}

func (d *CiscoVPNFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Cisco VPN feature template.",

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
			"vpn_id": schema.Int64Attribute{
				MarkdownDescription: "List of VPN instances",
				Computed:            true,
			},
			"vpn_name": schema.StringAttribute{
				MarkdownDescription: "Name",
				Computed:            true,
			},
			"vpn_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tenant_vpn_id": schema.Int64Attribute{
				MarkdownDescription: "Tenant VPN",
				Computed:            true,
			},
			"organization_name": schema.StringAttribute{
				MarkdownDescription: "Org Name selected",
				Computed:            true,
			},
			"omp_admin_distance_ipv4": schema.Int64Attribute{
				MarkdownDescription: "omp-admin-distance-ipv4",
				Computed:            true,
			},
			"omp_admin_distance_ipv4_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"omp_admin_distance_ipv6": schema.Int64Attribute{
				MarkdownDescription: "omp-admin-distance-ipv6",
				Computed:            true,
			},
			"omp_admin_distance_ipv6_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"enhance_ecmp_keying": schema.BoolAttribute{
				MarkdownDescription: "Optional packet fields for ECMP keying",
				Computed:            true,
			},
			"enhance_ecmp_keying_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"dns_ipv4_servers": schema.ListNestedAttribute{
				MarkdownDescription: "DNS",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "DNS Address",
							Computed:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"role": schema.StringAttribute{
							MarkdownDescription: "Role",
							Computed:            true,
						},
						"role_variable": schema.StringAttribute{
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
			"dns_ipv6_servers": schema.ListNestedAttribute{
				MarkdownDescription: "DNS",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "DNS Address",
							Computed:            true,
						},
						"role": schema.StringAttribute{
							MarkdownDescription: "Role",
							Computed:            true,
						},
						"role_variable": schema.StringAttribute{
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
			"dns_hosts": schema.ListNestedAttribute{
				MarkdownDescription: "Static DNS mapping",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"hostname": schema.StringAttribute{
							MarkdownDescription: "Hostname",
							Computed:            true,
						},
						"hostname_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ip": schema.SetAttribute{
							MarkdownDescription: "List of IP",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"ip_variable": schema.StringAttribute{
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
			"services": schema.ListNestedAttribute{
				MarkdownDescription: "Configure services",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"service_types": schema.StringAttribute{
							MarkdownDescription: "Service Type",
							Computed:            true,
						},
						"address": schema.SetAttribute{
							MarkdownDescription: "List of IPv4 address",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"interface": schema.StringAttribute{
							MarkdownDescription: "Tracking Service",
							Computed:            true,
						},
						"interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"track_enable": schema.BoolAttribute{
							MarkdownDescription: "Tracking Service",
							Computed:            true,
						},
						"track_enable_variable": schema.StringAttribute{
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
			"ipv4_static_service_routes": schema.ListNestedAttribute{
				MarkdownDescription: "Configure IPv4 Static Service Routes",
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
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "Destination VPN to resolve the prefix",
							Computed:            true,
						},
						"service": schema.StringAttribute{
							MarkdownDescription: "Service",
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_static_routes": schema.ListNestedAttribute{
				MarkdownDescription: "Configure IPv4 Static Routes",
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
						"null0": schema.BoolAttribute{
							MarkdownDescription: "null0",
							Computed:            true,
						},
						"null0_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"distance": schema.Int64Attribute{
							MarkdownDescription: "Administrative distance",
							Computed:            true,
						},
						"distance_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "Destination VPN(!=0 or !=512) to resolve the prefix",
							Computed:            true,
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dhcp": schema.BoolAttribute{
							MarkdownDescription: "Default Gateway obtained from DHCP",
							Computed:            true,
						},
						"dhcp_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"next_hops": schema.ListNestedAttribute{
							MarkdownDescription: "IP gateway address",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: "IP Address",
										Computed:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"distance": schema.Int64Attribute{
										MarkdownDescription: "Administrative distance",
										Computed:            true,
									},
									"distance_variable": schema.StringAttribute{
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
						"track_next_hops": schema.ListNestedAttribute{
							MarkdownDescription: "IP gateway address",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: "IP Address",
										Computed:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"distance": schema.Int64Attribute{
										MarkdownDescription: "Administrative distance",
										Computed:            true,
									},
									"distance_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"tracker": schema.StringAttribute{
										MarkdownDescription: "Static route tracker",
										Computed:            true,
									},
									"tracker_variable": schema.StringAttribute{
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
			"ipv6_static_routes": schema.ListNestedAttribute{
				MarkdownDescription: "Configure IPv6 Static Routes",
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
						"null0": schema.BoolAttribute{
							MarkdownDescription: "null0",
							Computed:            true,
						},
						"null0_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "Destination VPN(!=0 or !=512) to resolve the prefix",
							Computed:            true,
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"nat": schema.StringAttribute{
							MarkdownDescription: "NAT",
							Computed:            true,
						},
						"nat_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"next_hops": schema.ListNestedAttribute{
							MarkdownDescription: "IP gateway address",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: "IP Address",
										Computed:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"distance": schema.Int64Attribute{
										MarkdownDescription: "Administrative distance",
										Computed:            true,
									},
									"distance_variable": schema.StringAttribute{
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
			"ipv4_static_gre_routes": schema.ListNestedAttribute{
				MarkdownDescription: "Configure routes pointing to a GRE tunnel",
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
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "Destination VPN to resolve the prefix",
							Computed:            true,
						},
						"interface": schema.SetAttribute{
							MarkdownDescription: "List of GRE Interfaces",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"interface_variable": schema.StringAttribute{
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
			"ipv4_static_ipsec_routes": schema.ListNestedAttribute{
				MarkdownDescription: "Configure routes pointing to a IPSEC tunnel",
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
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "Destination VPN to resolve the prefix",
							Computed:            true,
						},
						"interface": schema.SetAttribute{
							MarkdownDescription: "List of IPSEC Interfaces (Separated by commas)",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"interface_variable": schema.StringAttribute{
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
			"omp_advertise_ipv4_routes": schema.ListNestedAttribute{
				MarkdownDescription: "Advertise routes to OMP",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Advertised routes protocol",
							Computed:            true,
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"route_policy": schema.StringAttribute{
							MarkdownDescription: "Set Route Policy to OMP",
							Computed:            true,
						},
						"route_policy_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"protocol_sub_type": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"protocol_sub_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"prefixes": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix_entry": schema.StringAttribute{
										MarkdownDescription: "Prefix",
										Computed:            true,
									},
									"prefix_entry_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"aggregate_only": schema.BoolAttribute{
										MarkdownDescription: "Aggregate Only",
										Computed:            true,
									},
									"aggregate_only_variable": schema.StringAttribute{
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
			"omp_advertise_ipv6_routes": schema.ListNestedAttribute{
				MarkdownDescription: "Advertise routes to OMP",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Advertised routes protocol",
							Computed:            true,
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"route_policy": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"route_policy_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"protocol_sub_type": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"protocol_sub_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"prefixes": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix_entry": schema.StringAttribute{
										MarkdownDescription: "Prefix",
										Computed:            true,
									},
									"prefix_entry_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"aggregate_only": schema.BoolAttribute{
										MarkdownDescription: "Aggregate Only",
										Computed:            true,
									},
									"aggregate_only_variable": schema.StringAttribute{
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
			"nat64_pools": schema.ListNestedAttribute{
				MarkdownDescription: "Set NAT64 v4 pool range",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "NAT64 Pool name",
							Computed:            true,
						},
						"start_address": schema.StringAttribute{
							MarkdownDescription: "Starting IP address of NAT pool range",
							Computed:            true,
						},
						"start_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"end_address": schema.StringAttribute{
							MarkdownDescription: "Ending IP address of NAT pool range",
							Computed:            true,
						},
						"end_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"overload": schema.BoolAttribute{
							MarkdownDescription: "NAT 64 Overload Option",
							Computed:            true,
						},
						"overload_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"leak_from_global": schema.BoolAttribute{
							MarkdownDescription: "Enable Route Leaking from Global VPN to this Service VPN",
							Computed:            true,
						},
						"leak_from_global_protocol": schema.StringAttribute{
							MarkdownDescription: "Select protocol for route leaking",
							Computed:            true,
						},
						"leak_to_global": schema.BoolAttribute{
							MarkdownDescription: "Enable Route Leaking from this Service VPN to Global VPN",
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"nat_pools": schema.ListNestedAttribute{
				MarkdownDescription: "Configure NAT Pool entries",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.Int64Attribute{
							MarkdownDescription: "NAT Pool Name, natpool1..31",
							Computed:            true,
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"prefix_length": schema.Int64Attribute{
							MarkdownDescription: "Ending IP address of NAT Pool Prefix Length",
							Computed:            true,
						},
						"prefix_length_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"range_start": schema.StringAttribute{
							MarkdownDescription: "Starting IP address of NAT pool range",
							Computed:            true,
						},
						"range_start_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"range_end": schema.StringAttribute{
							MarkdownDescription: "Ending IP address of NAT pool range",
							Computed:            true,
						},
						"range_end_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"overload": schema.BoolAttribute{
							MarkdownDescription: "Enable port translation(PAT)",
							Computed:            true,
						},
						"overload_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"direction": schema.StringAttribute{
							MarkdownDescription: "Direction of NAT translation",
							Computed:            true,
						},
						"direction_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tracker_id": schema.Int64Attribute{
							MarkdownDescription: "Add Object/Object Group Tracker",
							Computed:            true,
						},
						"tracker_id_variable": schema.StringAttribute{
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
			"static_nat_rules": schema.ListNestedAttribute{
				MarkdownDescription: "Configure static NAT entries",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"pool_name": schema.Int64Attribute{
							MarkdownDescription: "NAT Pool Name, natpool1..31",
							Computed:            true,
						},
						"pool_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_ip": schema.StringAttribute{
							MarkdownDescription: "Source IP address to be translated",
							Computed:            true,
						},
						"source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"translate_ip": schema.StringAttribute{
							MarkdownDescription: "Statically translated source IP address",
							Computed:            true,
						},
						"translate_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"static_nat_direction": schema.StringAttribute{
							MarkdownDescription: "Direction of static NAT translation",
							Computed:            true,
						},
						"static_nat_direction_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tracker_id": schema.Int64Attribute{
							MarkdownDescription: "Add Object/Object Group Tracker",
							Computed:            true,
						},
						"tracker_id_variable": schema.StringAttribute{
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
			"static_nat_subnet_rules": schema.ListNestedAttribute{
				MarkdownDescription: "Configure static NAT Subnet entries",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_ip_subnet": schema.StringAttribute{
							MarkdownDescription: "Source IP Subnet to be translated",
							Computed:            true,
						},
						"source_ip_subnet_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"translate_ip_subnet": schema.StringAttribute{
							MarkdownDescription: "Statically translated source IP Subnet",
							Computed:            true,
						},
						"translate_ip_subnet_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"prefix_length": schema.Int64Attribute{
							MarkdownDescription: "Network Prefix Length",
							Computed:            true,
						},
						"prefix_length_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"static_nat_direction": schema.StringAttribute{
							MarkdownDescription: "Direction of static NAT translation",
							Computed:            true,
						},
						"static_nat_direction_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tracker_id": schema.Int64Attribute{
							MarkdownDescription: "Add Object/Object Group Tracker",
							Computed:            true,
						},
						"tracker_id_variable": schema.StringAttribute{
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
			"port_forward_rules": schema.ListNestedAttribute{
				MarkdownDescription: "Configure Port Forward entries",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"pool_name": schema.Int64Attribute{
							MarkdownDescription: "NAT Pool Name, natpool1..31",
							Computed:            true,
						},
						"pool_name_variable": schema.StringAttribute{
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
							MarkdownDescription: "Source IP address to be translated",
							Computed:            true,
						},
						"source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"translate_ip": schema.StringAttribute{
							MarkdownDescription: "Statically translated source IP address",
							Computed:            true,
						},
						"translate_ip_variable": schema.StringAttribute{
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
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"route_global_imports": schema.ListNestedAttribute{
				MarkdownDescription: "Enable route leaking from Global VPN to this Service VPN",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Select a Route Protocol to enable route leaking from Global VPN to this Service VPN",
							Computed:            true,
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"protocol_sub_type": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"protocol_sub_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"route_policy": schema.StringAttribute{
							MarkdownDescription: "Select a Route Policy to enable route leaking from Global VPN to this Service VPN",
							Computed:            true,
						},
						"redistributes": schema.ListNestedAttribute{
							MarkdownDescription: "Enable redistribution of replicated route protocol",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: "Select a Route Protocol to enable redistribution",
										Computed:            true,
									},
									"protocol_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"route_policy": schema.StringAttribute{
										MarkdownDescription: "Select a Route Policy to enable redistribution",
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
			"route_vpn_imports": schema.ListNestedAttribute{
				MarkdownDescription: "Enable route leak from Service VPN to current VPN",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_vpn_id": schema.Int64Attribute{
							MarkdownDescription: "Select a Source VPN where route leaks from",
							Computed:            true,
						},
						"source_vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Select a Route Protocol to enable route leaking to current VPN",
							Computed:            true,
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"protocol_sub_type": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"protocol_sub_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"route_policy": schema.StringAttribute{
							MarkdownDescription: "Select a Route Policy to enable route leaking to current VPN",
							Computed:            true,
						},
						"route_policy_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"redistributes": schema.ListNestedAttribute{
							MarkdownDescription: "Enable redistribution of replicated route protocol",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: "Select a Route Protocol to enable redistribution",
										Computed:            true,
									},
									"protocol_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"route_policy": schema.StringAttribute{
										MarkdownDescription: "Select a Route Policy to enable redistribution",
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
			"route_global_exports": schema.ListNestedAttribute{
				MarkdownDescription: "Enable route leaking to Global VPN from this Service VPN",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Select a Route Protocol to enable route leaking from this Service VPN to Global VPN",
							Computed:            true,
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"protocol_sub_type": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"protocol_sub_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"route_policy": schema.StringAttribute{
							MarkdownDescription: "Select a Route Policy to enable route leaking from this Service VPN to Global VPN",
							Computed:            true,
						},
						"redistributes": schema.ListNestedAttribute{
							MarkdownDescription: "Enable redistribution of replicated route protocol",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: "Select a Route Protocol to enable redistribution",
										Computed:            true,
									},
									"protocol_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"route_policy": schema.StringAttribute{
										MarkdownDescription: "Select a Route Policy to enable redistribution",
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
		},
	}
}

func (d *CiscoVPNFeatureTemplateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *CiscoVPNFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *CiscoVPNFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CiscoVPN

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
