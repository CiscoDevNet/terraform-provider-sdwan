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
	_ datasource.DataSource              = &CiscoVPNInterfaceFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &CiscoVPNInterfaceFeatureTemplateDataSource{}
)

func NewCiscoVPNInterfaceFeatureTemplateDataSource() datasource.DataSource {
	return &CiscoVPNInterfaceFeatureTemplateDataSource{}
}

type CiscoVPNInterfaceFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *CiscoVPNInterfaceFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_vpn_interface_feature_template"
}

func (d *CiscoVPNInterfaceFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Cisco VPN Interface feature template.",

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
			"interface_name": schema.StringAttribute{
				MarkdownDescription: "Interface name: ge0/<0-..> or ge0/<0-..>.vlanid or irb<bridgeid:1-63> or loopback<string> or natpool-<1..31> when present",
				Computed:            true,
			},
			"interface_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"interface_description": schema.StringAttribute{
				MarkdownDescription: "Interface description",
				Computed:            true,
			},
			"interface_description_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"poe": schema.BoolAttribute{
				MarkdownDescription: "Configure interface as Power-over-Ethernet source",
				Computed:            true,
			},
			"poe_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"address": schema.StringAttribute{
				MarkdownDescription: "Assign IPv4 address",
				Computed:            true,
			},
			"address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_secondary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Assign secondary IP addresses",
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
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"dhcp": schema.BoolAttribute{
				MarkdownDescription: "Enable DHCP",
				Computed:            true,
			},
			"dhcp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"dhcp_distance": schema.Int64Attribute{
				MarkdownDescription: "Set administrative distance for DHCP default route",
				Computed:            true,
			},
			"dhcp_distance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_address": schema.StringAttribute{
				MarkdownDescription: "Assign IPv6 address",
				Computed:            true,
			},
			"ipv6_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"dhcpv6": schema.BoolAttribute{
				MarkdownDescription: "Enable DHCPv6",
				Computed:            true,
			},
			"dhcpv6_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_secondary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Assign secondary IPv6 addresses",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IPv6 Address",
							Computed:            true,
						},
						"address_variable": schema.StringAttribute{
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
			"ipv6_access_lists": schema.ListNestedAttribute{
				MarkdownDescription: "Apply IPv6 access list",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"direction": schema.StringAttribute{
							MarkdownDescription: "Direction",
							Computed:            true,
						},
						"acl_name": schema.StringAttribute{
							MarkdownDescription: "Name of access list",
							Computed:            true,
						},
						"acl_name_variable": schema.StringAttribute{
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
			"ipv4_dhcp_helper": schema.SetAttribute{
				MarkdownDescription: "List of DHCP IPv4 helper addresses",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ipv4_dhcp_helper_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_dhcp_helpers": schema.ListNestedAttribute{
				MarkdownDescription: "DHCPv6 Helper",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "DHCPv6 Helper address",
							Computed:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "DHCPv6 Helper VPN",
							Computed:            true,
						},
						"vpn_id_variable": schema.StringAttribute{
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
			"tracker": schema.SetAttribute{
				MarkdownDescription: "Enable tracker for this interface",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"tracker_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"auto_bandwidth_detect": schema.BoolAttribute{
				MarkdownDescription: "Interface auto detect bandwidth",
				Computed:            true,
			},
			"auto_bandwidth_detect_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"iperf_server": schema.StringAttribute{
				MarkdownDescription: "Iperf server for auto bandwidth detect",
				Computed:            true,
			},
			"iperf_server_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat": schema.BoolAttribute{
				MarkdownDescription: "Network Address Translation on this interface",
				Computed:            true,
			},
			"nat_type": schema.StringAttribute{
				MarkdownDescription: "NAT type",
				Computed:            true,
			},
			"nat_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"udp_timeout": schema.Int64Attribute{
				MarkdownDescription: "Set NAT UDP session timeout, in minutes",
				Computed:            true,
			},
			"udp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tcp_timeout": schema.Int64Attribute{
				MarkdownDescription: "Set NAT TCP session timeout, in minutes",
				Computed:            true,
			},
			"tcp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat_pool_range_start": schema.StringAttribute{
				MarkdownDescription: "Starting IP address of NAT pool range",
				Computed:            true,
			},
			"nat_pool_range_start_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat_pool_range_end": schema.StringAttribute{
				MarkdownDescription: "Ending IP address of NAT pool range",
				Computed:            true,
			},
			"nat_pool_range_end_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat_overload": schema.BoolAttribute{
				MarkdownDescription: "Enable port translation(PAT)",
				Computed:            true,
			},
			"nat_overload_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat_inside_source_loopback_interface": schema.StringAttribute{
				MarkdownDescription: "Configure NAT Inside Loopback Interface",
				Computed:            true,
			},
			"nat_inside_source_loopback_interface_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat_pool_prefix_length": schema.Int64Attribute{
				MarkdownDescription: "Ending IP address of NAT Pool Prefix Length",
				Computed:            true,
			},
			"nat_pool_prefix_length_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_nat": schema.BoolAttribute{
				MarkdownDescription: "NAT64 on this interface",
				Computed:            true,
			},
			"ipv6_nat_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat64_interface": schema.BoolAttribute{
				MarkdownDescription: "NAT64 on this interface",
				Computed:            true,
			},
			"nat66_interface": schema.BoolAttribute{
				MarkdownDescription: "NAT66 on this interface",
				Computed:            true,
			},
			"static_nat66_entries": schema.ListNestedAttribute{
				MarkdownDescription: "static NAT",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_prefix": schema.StringAttribute{
							MarkdownDescription: "Source Prefix",
							Computed:            true,
						},
						"source_prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"translated_source_prefix": schema.StringAttribute{
							MarkdownDescription: "Translated Source Prefix",
							Computed:            true,
						},
						"translated_source_prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_vpn_id": schema.Int64Attribute{
							MarkdownDescription: "Source VPN ID",
							Computed:            true,
						},
						"source_vpn_id_variable": schema.StringAttribute{
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
			"static_nat_entries": schema.ListNestedAttribute{
				MarkdownDescription: "Configure static NAT entries",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
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
						"source_vpn_id": schema.Int64Attribute{
							MarkdownDescription: "Configure VPN ID",
							Computed:            true,
						},
						"source_vpn_id_variable": schema.StringAttribute{
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
			"static_port_forward_entries": schema.ListNestedAttribute{
				MarkdownDescription: "Configure Port Forward entries",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
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
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Protocol",
							Computed:            true,
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_vpn_id": schema.Int64Attribute{
							MarkdownDescription: "Configure VPN ID",
							Computed:            true,
						},
						"source_vpn_id_variable": schema.StringAttribute{
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
			"enable_core_region": schema.BoolAttribute{
				MarkdownDescription: "Enable core region",
				Computed:            true,
			},
			"core_region": schema.StringAttribute{
				MarkdownDescription: "Enable core region",
				Computed:            true,
			},
			"core_region_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"secondary_region": schema.StringAttribute{
				MarkdownDescription: "Enable secondary region",
				Computed:            true,
			},
			"secondary_region_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_encapsulations": schema.ListNestedAttribute{
				MarkdownDescription: "Encapsulation for TLOC",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"encapsulation": schema.StringAttribute{
							MarkdownDescription: "Encapsulation",
							Computed:            true,
						},
						"preference": schema.Int64Attribute{
							MarkdownDescription: "Set preference for TLOC",
							Computed:            true,
						},
						"preference_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"weight": schema.Int64Attribute{
							MarkdownDescription: "Set weight for TLOC",
							Computed:            true,
						},
						"weight_variable": schema.StringAttribute{
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
			"tunnel_interface_border": schema.BoolAttribute{
				MarkdownDescription: "Set TLOC as border TLOC",
				Computed:            true,
			},
			"tunnel_interface_border_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_qos_mode": schema.StringAttribute{
				MarkdownDescription: "Set tunnel QoS mode",
				Computed:            true,
			},
			"tunnel_qos_mode_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_bandwidth": schema.Int64Attribute{
				MarkdownDescription: "Tunnels Bandwidth Percent",
				Computed:            true,
			},
			"tunnel_bandwidth_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_groups": schema.SetAttribute{
				MarkdownDescription: "List of groups",
				ElementType:         types.Int64Type,
				Computed:            true,
			},
			"tunnel_interface_groups_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_color": schema.StringAttribute{
				MarkdownDescription: "Set color for TLOC",
				Computed:            true,
			},
			"tunnel_interface_color_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_max_control_connections": schema.Int64Attribute{
				MarkdownDescription: "Set the maximum number of control connections for this TLOC",
				Computed:            true,
			},
			"tunnel_interface_max_control_connections_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_control_connections": schema.BoolAttribute{
				MarkdownDescription: "Allow Control Connection",
				Computed:            true,
			},
			"tunnel_interface_control_connections_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_vbond_as_stun_server": schema.BoolAttribute{
				MarkdownDescription: "Put this wan interface in STUN mode only",
				Computed:            true,
			},
			"tunnel_interface_vbond_as_stun_server_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_exclude_controller_group_list": schema.SetAttribute{
				MarkdownDescription: "Exclude the following controller groups defined in this list",
				ElementType:         types.Int64Type,
				Computed:            true,
			},
			"tunnel_interface_exclude_controller_group_list_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_vmanage_connection_preference": schema.Int64Attribute{
				MarkdownDescription: "Set interface preference for control connection to vManage <0..8>",
				Computed:            true,
			},
			"tunnel_interface_vmanage_connection_preference_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_port_hop": schema.BoolAttribute{
				MarkdownDescription: "Disallow port hopping on the tunnel interface",
				Computed:            true,
			},
			"tunnel_interface_port_hop_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_color_restrict": schema.BoolAttribute{
				MarkdownDescription: "Restrict this TLOC behavior",
				Computed:            true,
			},
			"tunnel_interface_color_restrict_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_gre_tunnel_destination_ip": schema.StringAttribute{
				MarkdownDescription: "Extend the TLOC to a remote node over GRE tunnel",
				Computed:            true,
			},
			"tunnel_interface_gre_tunnel_destination_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_carrier": schema.StringAttribute{
				MarkdownDescription: "Set carrier for TLOC",
				Computed:            true,
			},
			"tunnel_interface_carrier_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_nat_refresh_interval": schema.Int64Attribute{
				MarkdownDescription: "Set time period of nat refresh packets <1...60> seconds",
				Computed:            true,
			},
			"tunnel_interface_nat_refresh_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_hello_interval": schema.Int64Attribute{
				MarkdownDescription: "Set time period of control hello packets <100..600000> milli seconds",
				Computed:            true,
			},
			"tunnel_interface_hello_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_hello_tolerance": schema.Int64Attribute{
				MarkdownDescription: "Set tolerance of control hello packets <12..6000> seconds",
				Computed:            true,
			},
			"tunnel_interface_hello_tolerance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_bind_loopback_tunnel": schema.StringAttribute{
				MarkdownDescription: "Bind loopback tunnel interface to a physical interface",
				Computed:            true,
			},
			"tunnel_interface_bind_loopback_tunnel_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_last_resort_circuit": schema.BoolAttribute{
				MarkdownDescription: "Set TLOC as last resort",
				Computed:            true,
			},
			"tunnel_interface_last_resort_circuit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_low_bandwidth_link": schema.BoolAttribute{
				MarkdownDescription: "Set the interface as a low-bandwidth circuit",
				Computed:            true,
			},
			"tunnel_interface_low_bandwidth_link_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_tunnel_tcp_mss": schema.Int64Attribute{
				MarkdownDescription: "Tunnel TCP MSS on SYN packets, in bytes",
				Computed:            true,
			},
			"tunnel_interface_tunnel_tcp_mss_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_clear_dont_fragment": schema.BoolAttribute{
				MarkdownDescription: "Enable clear dont fragment (Currently Only SDWAN Tunnel Interface)",
				Computed:            true,
			},
			"tunnel_interface_clear_dont_fragment_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_propagate_sgt": schema.BoolAttribute{
				MarkdownDescription: "CTS SGT Propagation configuration",
				Computed:            true,
			},
			"tunnel_interface_propagate_sgt_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_network_broadcast": schema.BoolAttribute{
				MarkdownDescription: "Accept and respond to network-prefix-directed broadcasts)",
				Computed:            true,
			},
			"tunnel_interface_network_broadcast_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_all": schema.BoolAttribute{
				MarkdownDescription: "Allow all traffic. Overrides all other allow-service options if allow-service all is set",
				Computed:            true,
			},
			"tunnel_interface_allow_all_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_bgp": schema.BoolAttribute{
				MarkdownDescription: "Allow/deny BGP",
				Computed:            true,
			},
			"tunnel_interface_allow_bgp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_dhcp": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny DHCP",
				Computed:            true,
			},
			"tunnel_interface_allow_dhcp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_dns": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny DNS",
				Computed:            true,
			},
			"tunnel_interface_allow_dns_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_icmp": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny ICMP",
				Computed:            true,
			},
			"tunnel_interface_allow_icmp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_ssh": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny SSH",
				Computed:            true,
			},
			"tunnel_interface_allow_ssh_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_netconf": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny NETCONF",
				Computed:            true,
			},
			"tunnel_interface_allow_netconf_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_ntp": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny NTP",
				Computed:            true,
			},
			"tunnel_interface_allow_ntp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_ospf": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny OSPF",
				Computed:            true,
			},
			"tunnel_interface_allow_ospf_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_stun": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny STUN",
				Computed:            true,
			},
			"tunnel_interface_allow_stun_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_snmp": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny SNMP",
				Computed:            true,
			},
			"tunnel_interface_allow_snmp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_https": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny Https",
				Computed:            true,
			},
			"tunnel_interface_allow_https_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"media_type": schema.StringAttribute{
				MarkdownDescription: "Media type",
				Computed:            true,
			},
			"media_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"interface_mtu": schema.Int64Attribute{
				MarkdownDescription: "Interface MTU GigabitEthernet0 <1500..1518>, Other GigabitEthernet <1500..9216> in bytes",
				Computed:            true,
			},
			"interface_mtu_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ip_mtu": schema.Int64Attribute{
				MarkdownDescription: "IP MTU for GigabitEthernet main <576..Interface MTU>, GigabitEthernet subinterface <576..9216>, Other Interfaces <576..2000> in bytes",
				Computed:            true,
			},
			"ip_mtu_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tcp_mss_adjust": schema.Int64Attribute{
				MarkdownDescription: "TCP MSS on SYN packets, in bytes",
				Computed:            true,
			},
			"tcp_mss_adjust_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tloc_extension": schema.StringAttribute{
				MarkdownDescription: "Extends a local TLOC to a remote node only for vpn 0",
				Computed:            true,
			},
			"tloc_extension_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"load_interval": schema.Int64Attribute{
				MarkdownDescription: "Interval for interface load calculation",
				Computed:            true,
			},
			"load_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"gre_tunnel_source_ip": schema.StringAttribute{
				MarkdownDescription: "Extend remote TLOC over a GRE tunnel to a local WAN interface",
				Computed:            true,
			},
			"gre_tunnel_source_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"gre_tunnel_xconnect": schema.StringAttribute{
				MarkdownDescription: "Extend remote TLOC over a GRE tunnel to a local WAN interface",
				Computed:            true,
			},
			"gre_tunnel_xconnect_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"mac_address": schema.StringAttribute{
				MarkdownDescription: "Set MAC-layer address",
				Computed:            true,
			},
			"mac_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"speed": schema.StringAttribute{
				MarkdownDescription: "Set interface speed",
				Computed:            true,
			},
			"speed_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"duplex": schema.StringAttribute{
				MarkdownDescription: "Duplex mode",
				Computed:            true,
			},
			"duplex_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"shutdown_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"arp_timeout": schema.Int64Attribute{
				MarkdownDescription: "Timeout value for dynamically learned ARP entries, <0..2678400> seconds",
				Computed:            true,
			},
			"arp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"autonegotiate": schema.BoolAttribute{
				MarkdownDescription: "Link autonegotiation",
				Computed:            true,
			},
			"autonegotiate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ip_directed_broadcast": schema.BoolAttribute{
				MarkdownDescription: "IP Directed-Broadcast",
				Computed:            true,
			},
			"ip_directed_broadcast_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"icmp_redirect_disable": schema.BoolAttribute{
				MarkdownDescription: "Set this option to disable the icmp/icmpv6 redirect packets",
				Computed:            true,
			},
			"icmp_redirect_disable_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_adaptive_period": schema.Int64Attribute{
				MarkdownDescription: "Periodic timer for adaptive QoS in minutes",
				Computed:            true,
			},
			"qos_adaptive_period_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_adaptive_bandwidth_downstream": schema.Int64Attribute{
				MarkdownDescription: "Adaptive QoS default downstream bandwidth",
				Computed:            true,
			},
			"qos_adaptive_bandwidth_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_adaptive_min_downstream": schema.Int64Attribute{
				MarkdownDescription: "Downstream min bandwidth limit",
				Computed:            true,
			},
			"qos_adaptive_min_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_adaptive_max_downstream": schema.Int64Attribute{
				MarkdownDescription: "Downstream max bandwidth limit",
				Computed:            true,
			},
			"qos_adaptive_max_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_adaptive_bandwidth_upstream": schema.Int64Attribute{
				MarkdownDescription: "Adaptive QoS default upstream bandwidth",
				Computed:            true,
			},
			"qos_adaptive_bandwidth_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_adaptive_min_upstream": schema.Int64Attribute{
				MarkdownDescription: "Upstream min bandwidth limit",
				Computed:            true,
			},
			"qos_adaptive_min_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_adaptive_max_upstream": schema.Int64Attribute{
				MarkdownDescription: "Upstream max bandwidth limit",
				Computed:            true,
			},
			"qos_adaptive_max_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"shaping_rate": schema.Int64Attribute{
				MarkdownDescription: "1ge  interfaces: [0..1000000]kbps; 10ge interfaces: [0..10000000]kbps",
				Computed:            true,
			},
			"shaping_rate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_map": schema.StringAttribute{
				MarkdownDescription: "Name of QoS map",
				Computed:            true,
			},
			"qos_map_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_map_vpn": schema.StringAttribute{
				MarkdownDescription: "Name of VPN QoS map",
				Computed:            true,
			},
			"qos_map_vpn_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"bandwidth_upstream": schema.Int64Attribute{
				MarkdownDescription: "Interface upstream bandwidth capacity, in kbps",
				Computed:            true,
			},
			"bandwidth_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"bandwidth_downstream": schema.Int64Attribute{
				MarkdownDescription: "Interface downstream bandwidth capacity, in kbps",
				Computed:            true,
			},
			"bandwidth_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"block_non_source_ip": schema.BoolAttribute{
				MarkdownDescription: "Block packets originating from IP address that is not from this source",
				Computed:            true,
			},
			"block_non_source_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"rewrite_rule_name": schema.StringAttribute{
				MarkdownDescription: "Name of rewrite rule",
				Computed:            true,
			},
			"rewrite_rule_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"access_lists": schema.ListNestedAttribute{
				MarkdownDescription: "Apply ACL",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"direction": schema.StringAttribute{
							MarkdownDescription: "Direction",
							Computed:            true,
						},
						"acl_name": schema.StringAttribute{
							MarkdownDescription: "Name of access list",
							Computed:            true,
						},
						"acl_name_variable": schema.StringAttribute{
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
			"static_arps": schema.ListNestedAttribute{
				MarkdownDescription: "Configure static ARP entries",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "IP Address",
							Computed:            true,
						},
						"ip_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: "MAC address",
							Computed:            true,
						},
						"mac_variable": schema.StringAttribute{
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
			"ipv4_vrrps": schema.ListNestedAttribute{
				MarkdownDescription: "Enable VRRP",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.Int64Attribute{
							MarkdownDescription: "Group ID",
							Computed:            true,
						},
						"group_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: "Set priority",
							Computed:            true,
						},
						"priority_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"timer": schema.Int64Attribute{
							MarkdownDescription: "Timer interval for successive advertisements, in milliseconds",
							Computed:            true,
						},
						"timer_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"track_omp": schema.BoolAttribute{
							MarkdownDescription: "Track OMP status",
							Computed:            true,
						},
						"track_prefix_list": schema.StringAttribute{
							MarkdownDescription: "Track Prefix List",
							Computed:            true,
						},
						"track_prefix_list_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "Assign IP Address",
							Computed:            true,
						},
						"ip_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ipv4_secondary_addresses": schema.ListNestedAttribute{
							MarkdownDescription: "VRRP Secondary IP address",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										MarkdownDescription: "VRRP Secondary IP address",
										Computed:            true,
									},
									"ip_address_variable": schema.StringAttribute{
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
						"tloc_preference_change": schema.BoolAttribute{
							MarkdownDescription: "change TLOC preference",
							Computed:            true,
						},
						"tloc_preference_change_value": schema.Int64Attribute{
							MarkdownDescription: "Set tloc preference change value",
							Computed:            true,
						},
						"tloc_preference_change_value_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tracking_objects": schema.ListNestedAttribute{
							MarkdownDescription: "tracking object for VRRP configuration",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"tracker_id": schema.Int64Attribute{
										MarkdownDescription: "Tracker ID",
										Computed:            true,
									},
									"tracker_id_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"track_action": schema.StringAttribute{
										MarkdownDescription: "Track Action",
										Computed:            true,
									},
									"track_action_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"decrement_value": schema.Int64Attribute{
										MarkdownDescription: "Decrement Value for VRRP priority",
										Computed:            true,
									},
									"decrement_value_variable": schema.StringAttribute{
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
			"ipv6_vrrps": schema.ListNestedAttribute{
				MarkdownDescription: "Enable VRRP",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.Int64Attribute{
							MarkdownDescription: "Group ID",
							Computed:            true,
						},
						"group_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: "Set priority",
							Computed:            true,
						},
						"priority_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"timer": schema.Int64Attribute{
							MarkdownDescription: "Timer interval for successive advertisements, in milliseconds",
							Computed:            true,
						},
						"timer_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"track_omp": schema.BoolAttribute{
							MarkdownDescription: "Track OMP status",
							Computed:            true,
						},
						"track_omp_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"track_prefix_list": schema.StringAttribute{
							MarkdownDescription: "Track Prefix List",
							Computed:            true,
						},
						"track_prefix_list_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ipv6_addresses": schema.ListNestedAttribute{
							MarkdownDescription: "IPv6 VRRP",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ipv6_link_local": schema.StringAttribute{
										MarkdownDescription: "Use link-local IPv6 Address",
										Computed:            true,
									},
									"ipv6_link_local_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"prefix": schema.StringAttribute{
										MarkdownDescription: "Assign Global IPv6 Prefix",
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
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"propagate_sgt": schema.BoolAttribute{
				MarkdownDescription: "Enable/Disable CTS SGT propagation on an interface.",
				Computed:            true,
			},
			"static_sgt": schema.Int64Attribute{
				MarkdownDescription: "SGT value between 2 and 65519.",
				Computed:            true,
			},
			"static_sgt_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"static_sgt_trusted": schema.BoolAttribute{
				MarkdownDescription: "Indicates that the interface is trustworthy for CTS.",
				Computed:            true,
			},
			"enable_sgt": schema.BoolAttribute{
				MarkdownDescription: "Enables the interface for CTS SGT authorization and forwarding.",
				Computed:            true,
			},
			"sgt_enforcement": schema.BoolAttribute{
				MarkdownDescription: "Enables the interface for CTS SGT authorization and forwarding.",
				Computed:            true,
			},
			"sgt_enforcement_sgt": schema.Int64Attribute{
				MarkdownDescription: "SGT value between 2 and 65519.",
				Computed:            true,
			},
			"sgt_enforcement_sgt_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
		},
	}
}

func (d *CiscoVPNInterfaceFeatureTemplateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *CiscoVPNInterfaceFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *CiscoVPNInterfaceFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CiscoVPNInterface

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
