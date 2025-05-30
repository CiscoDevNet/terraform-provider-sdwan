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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &TransportWANVPNInterfaceEthernetProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &TransportWANVPNInterfaceEthernetProfileParcelDataSource{}
)

func NewTransportWANVPNInterfaceEthernetProfileParcelDataSource() datasource.DataSource {
	return &TransportWANVPNInterfaceEthernetProfileParcelDataSource{}
}

type TransportWANVPNInterfaceEthernetProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *TransportWANVPNInterfaceEthernetProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transport_wan_vpn_interface_ethernet_feature"
}

func (d *TransportWANVPNInterfaceEthernetProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Transport WAN VPN Interface Ethernet Feature.",

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
			"transport_wan_vpn_feature_id": schema.StringAttribute{
				MarkdownDescription: "Transport WAN VPN Feature ID",
				Required:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"shutdown_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"interface_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"interface_description": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"interface_description_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_configuration_type": schema.StringAttribute{
				MarkdownDescription: "IPv4 Configuration Type",
				Computed:            true,
			},
			"ipv4_dhcp_distance": schema.Int64Attribute{
				MarkdownDescription: "DHCP Distance",
				Computed:            true,
			},
			"ipv4_dhcp_distance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: "IP Address",
				Computed:            true,
			},
			"ipv4_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_subnet_mask": schema.StringAttribute{
				MarkdownDescription: "Subnet Mask",
				Computed:            true,
			},
			"ipv4_subnet_mask_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_secondary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Secondary IpV4 Addresses",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IpV4 Address",
							Computed:            true,
						},
						"address_variable": schema.StringAttribute{
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
					},
				},
			},
			"ipv4_dhcp_helper": schema.SetAttribute{
				MarkdownDescription: "List of DHCP IPv4 helper addresses (min 1, max 8)",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ipv4_dhcp_helper_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_configuration_type": schema.StringAttribute{
				MarkdownDescription: "IPv6 Configuration Type",
				Computed:            true,
			},
			"enable_dhcpv6": schema.BoolAttribute{
				MarkdownDescription: "Enable DHCPv6",
				Computed:            true,
			},
			"ipv6_dhcp_secondary_address": schema.ListNestedAttribute{
				MarkdownDescription: "secondary IPv6 addresses",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IPv6 Address Secondary",
							Computed:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"ipv6_address": schema.StringAttribute{
				MarkdownDescription: "IPv6 Address Secondary",
				Computed:            true,
			},
			"ipv6_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_secondary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Static secondary IPv6 addresses",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IPv6 Address Secondary",
							Computed:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"iperf_server": schema.StringAttribute{
				MarkdownDescription: "Iperf server for auto bandwidth detect",
				Computed:            true,
			},
			"iperf_server_variable": schema.StringAttribute{
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
			"service_provider": schema.StringAttribute{
				MarkdownDescription: "Service Provider Name",
				Computed:            true,
			},
			"service_provider_variable": schema.StringAttribute{
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
			"auto_detect_bandwidth": schema.BoolAttribute{
				MarkdownDescription: "Interface auto detect bandwidth",
				Computed:            true,
			},
			"auto_detect_bandwidth_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface": schema.BoolAttribute{
				MarkdownDescription: "Tunnel Interface on/off",
				Computed:            true,
			},
			"per_tunnel_qos": schema.BoolAttribute{
				MarkdownDescription: "Per-tunnel Qos",
				Computed:            true,
			},
			"per_tunnel_qos_variable": schema.StringAttribute{
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
			"tunnel_bandwidth_percent": schema.Int64Attribute{
				MarkdownDescription: "Tunnels Bandwidth Percent",
				Computed:            true,
			},
			"tunnel_bandwidth_percent_variable": schema.StringAttribute{
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
			"tunnel_interface_carrier": schema.StringAttribute{
				MarkdownDescription: "Set carrier for TLOC",
				Computed:            true,
			},
			"tunnel_interface_carrier_variable": schema.StringAttribute{
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
			"tunnel_interface_last_resort_circuit": schema.BoolAttribute{
				MarkdownDescription: "Set TLOC as last resort",
				Computed:            true,
			},
			"tunnel_interface_last_resort_circuit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_gre_tunnel_destination_ip": schema.StringAttribute{
				MarkdownDescription: "GRE tunnel destination IP",
				Computed:            true,
			},
			"tunnel_interface_gre_tunnel_destination_ip_variable": schema.StringAttribute{
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
			"tunnel_interface_groups": schema.Int64Attribute{
				MarkdownDescription: "List of groups",
				Computed:            true,
			},
			"tunnel_interface_groups_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_border": schema.BoolAttribute{
				MarkdownDescription: "Set TLOC as border TLOC",
				Computed:            true,
			},
			"tunnel_interface_border_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_max_control_connections": schema.Int64Attribute{
				MarkdownDescription: "Maximum Control Connections",
				Computed:            true,
			},
			"tunnel_interface_max_control_connections_variable": schema.StringAttribute{
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
			"tunnel_interface_vbond_as_stun_server": schema.BoolAttribute{
				MarkdownDescription: "Put this wan interface in STUN mode only",
				Computed:            true,
			},
			"tunnel_interface_vbond_as_stun_server_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_exclude_controller_group_list": schema.SetAttribute{
				MarkdownDescription: "Exclude the following controller groups defined in this list.",
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
			"tunnel_interface_cts_sgt_propagation": schema.BoolAttribute{
				MarkdownDescription: "CTS SGT Propagation configuration",
				Computed:            true,
			},
			"tunnel_interface_cts_sgt_propagation_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_network_broadcast": schema.BoolAttribute{
				MarkdownDescription: "Accept and respond to network-prefix-directed broadcasts",
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
			"tunnel_interface_allow_ntp": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny NTP",
				Computed:            true,
			},
			"tunnel_interface_allow_ntp_variable": schema.StringAttribute{
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
			"tunnel_interface_allow_https": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny HTTPS",
				Computed:            true,
			},
			"tunnel_interface_allow_https_variable": schema.StringAttribute{
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
			"tunnel_interface_allow_netconf": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny NETCONF",
				Computed:            true,
			},
			"tunnel_interface_allow_netconf_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_allow_bfd": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny BFD",
				Computed:            true,
			},
			"tunnel_interface_allow_bfd_variable": schema.StringAttribute{
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
					},
				},
			},
			"nat_ipv4": schema.BoolAttribute{
				MarkdownDescription: "enable Network Address Translation on this interface",
				Computed:            true,
			},
			"nat_ipv4_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat_type": schema.StringAttribute{
				MarkdownDescription: "NAT Type",
				Computed:            true,
			},
			"nat_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat_range_start": schema.StringAttribute{
				MarkdownDescription: "NAT Pool Range Start",
				Computed:            true,
			},
			"nat_range_start_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat_range_end": schema.StringAttribute{
				MarkdownDescription: "NAT Pool Range End",
				Computed:            true,
			},
			"nat_range_end_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat_prefix_length": schema.Int64Attribute{
				MarkdownDescription: "NAT Pool Prefix Length",
				Computed:            true,
			},
			"nat_prefix_length_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat_overload": schema.BoolAttribute{
				MarkdownDescription: "NAT Overload",
				Computed:            true,
			},
			"nat_overload_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat_loopback": schema.StringAttribute{
				MarkdownDescription: "NAT Inside Source Loopback Interface",
				Computed:            true,
			},
			"nat_loopback_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat_udp_timeout": schema.Int64Attribute{
				MarkdownDescription: "Set NAT UDP session timeout, in minutes",
				Computed:            true,
			},
			"nat_udp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat_tcp_timeout": schema.Int64Attribute{
				MarkdownDescription: "Set NAT TCP session timeout, in minutes",
				Computed:            true,
			},
			"nat_tcp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"new_static_nats": schema.ListNestedAttribute{
				MarkdownDescription: "static NAT",
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
						"translated_ip": schema.StringAttribute{
							MarkdownDescription: "Statically translated source IP address",
							Computed:            true,
						},
						"translated_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"direction": schema.StringAttribute{
							MarkdownDescription: "Direction of static NAT translation",
							Computed:            true,
						},
						"source_vpn": schema.Int64Attribute{
							MarkdownDescription: "Source VPN ID",
							Computed:            true,
						},
						"source_vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"nat_ipv6": schema.BoolAttribute{
				MarkdownDescription: "enable Network Address Translation ipv6 on this interface",
				Computed:            true,
			},
			"nat_ipv6_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat64": schema.BoolAttribute{
				MarkdownDescription: "NAT64 on this interface",
				Computed:            true,
			},
			"nat66": schema.BoolAttribute{
				MarkdownDescription: "NAT66 on this interface",
				Computed:            true,
			},
			"static_nat66": schema.ListNestedAttribute{
				MarkdownDescription: "static NAT66",
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
					},
				},
			},
			"qos_adaptive": schema.BoolAttribute{
				MarkdownDescription: "Adaptive QoS",
				Computed:            true,
			},
			"qos_adaptive_period": schema.Int64Attribute{
				MarkdownDescription: "Adapt Period(Minutes)",
				Computed:            true,
			},
			"qos_adaptive_period_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_adaptive_bandwidth_upstream": schema.BoolAttribute{
				MarkdownDescription: "Shaping Rate Upstream",
				Computed:            true,
			},
			"qos_adaptive_min_upstream": schema.Int64Attribute{
				MarkdownDescription: "Upstream min bandwidth limit (kbps)",
				Computed:            true,
			},
			"qos_adaptive_min_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_adaptive_max_upstream": schema.Int64Attribute{
				MarkdownDescription: "Upstream max bandwidth limit (kbps)",
				Computed:            true,
			},
			"qos_adaptive_max_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_adaptive_default_upstream": schema.Int64Attribute{
				MarkdownDescription: "Adaptive QoS default upstream bandwidth (kbps)",
				Computed:            true,
			},
			"qos_adaptive_default_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_adaptive_bandwidth_downstream": schema.BoolAttribute{
				MarkdownDescription: "Shaping Rate Downstream",
				Computed:            true,
			},
			"qos_adaptive_min_downstream": schema.Int64Attribute{
				MarkdownDescription: "Downstream min bandwidth limit (kbps)",
				Computed:            true,
			},
			"qos_adaptive_min_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_adaptive_max_downstream": schema.Int64Attribute{
				MarkdownDescription: "Downstream max bandwidth limit (kbps)",
				Computed:            true,
			},
			"qos_adaptive_max_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_adaptive_default_downstream": schema.Int64Attribute{
				MarkdownDescription: "Adaptive QoS default downstream bandwidth (kbps)",
				Computed:            true,
			},
			"qos_adaptive_default_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"qos_shaping_rate": schema.Int64Attribute{
				MarkdownDescription: "Shaping Rate (Kbps)",
				Computed:            true,
			},
			"qos_shaping_rate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"acl_ipv4_egress_feature_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"acl_ipv4_ingress_feature_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"acl_ipv6_egress_feature_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"acl_ipv6_ingress_feature_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"arps": schema.ListNestedAttribute{
				MarkdownDescription: "Configure ARP entries",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "IP V4 Address",
							Computed:            true,
						},
						"ip_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"mac_address": schema.StringAttribute{
							MarkdownDescription: "MAC Address",
							Computed:            true,
						},
						"mac_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"icmp_redirect_disable": schema.BoolAttribute{
				MarkdownDescription: "ICMP/ICMPv6 Redirect Disable",
				Computed:            true,
			},
			"icmp_redirect_disable_variable": schema.StringAttribute{
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
			"mac_address": schema.StringAttribute{
				MarkdownDescription: "MAC Address",
				Computed:            true,
			},
			"mac_address_variable": schema.StringAttribute{
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
			"interface_mtu": schema.Int64Attribute{
				MarkdownDescription: "Interface MTU GigabitEthernet0 <1500..1518>, Other GigabitEthernet <1500..9216> in bytes",
				Computed:            true,
			},
			"interface_mtu_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tcp_mss": schema.Int64Attribute{
				MarkdownDescription: "TCP MSS on SYN packets, in bytes",
				Computed:            true,
			},
			"tcp_mss_variable": schema.StringAttribute{
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
			"media_type": schema.StringAttribute{
				MarkdownDescription: "Media type",
				Computed:            true,
			},
			"media_type_variable": schema.StringAttribute{
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
			"gre_tunnel_source_ip": schema.StringAttribute{
				MarkdownDescription: "GRE tunnel source IP",
				Computed:            true,
			},
			"gre_tunnel_source_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"xconnect": schema.StringAttribute{
				MarkdownDescription: "Extend remote TLOC over a GRE tunnel to a local WAN interface",
				Computed:            true,
			},
			"xconnect_variable": schema.StringAttribute{
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
			"tracker": schema.StringAttribute{
				MarkdownDescription: "Enable tracker for this interface",
				Computed:            true,
			},
			"tracker_variable": schema.StringAttribute{
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
		},
	}
}

func (d *TransportWANVPNInterfaceEthernetProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *TransportWANVPNInterfaceEthernetProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config TransportWANVPNInterfaceEthernet

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
