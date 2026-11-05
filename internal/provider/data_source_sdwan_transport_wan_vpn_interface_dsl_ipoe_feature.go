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
	_ datasource.DataSource                     = &TransportWANVPNInterfaceDSLIPoEProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure        = &TransportWANVPNInterfaceDSLIPoEProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigValidators = &TransportWANVPNInterfaceDSLIPoEProfileParcelDataSource{}
)

func NewTransportWANVPNInterfaceDSLIPoEProfileParcelDataSource() datasource.DataSource {
	return &TransportWANVPNInterfaceDSLIPoEProfileParcelDataSource{}
}

type TransportWANVPNInterfaceDSLIPoEProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *TransportWANVPNInterfaceDSLIPoEProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transport_wan_vpn_interface_dsl_ipoe_feature"
}

func (d *TransportWANVPNInterfaceDSLIPoEProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Transport WAN VPN Interface DSL IPoE Feature.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Feature",
				Optional:            true,
				Computed:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Feature",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Feature",
				Optional:            true,
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
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"shutdown_variable": schema.StringAttribute{
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
			"vdsl_slot": schema.StringAttribute{
				MarkdownDescription: "Set module slot/subslot/port number",
				Computed:            true,
			},
			"vdsl_slot_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"vdsl_mode": schema.StringAttribute{
				MarkdownDescription: "Set VDSL operating mode",
				Computed:            true,
			},
			"vdsl_sra": schema.BoolAttribute{
				MarkdownDescription: "Seamless rate adaption",
				Computed:            true,
			},
			"vdsl_sra_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat_enabled": schema.BoolAttribute{
				MarkdownDescription: "Network Address Translation on this interface",
				Computed:            true,
			},
			"nat_enabled_variable": schema.StringAttribute{
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
			"nat_udp_timeout": schema.Int64Attribute{
				MarkdownDescription: "Set NAT UDP session timeout, in minutes",
				Computed:            true,
			},
			"nat_udp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: "Ethernet Interface/Sub Interface Name including sub interface number",
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
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: "Encapsulation VLAN id",
				Computed:            true,
			},
			"vlan_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_address_type": schema.StringAttribute{
				MarkdownDescription: "IPv4 address configuration type",
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
			"ipv4_dhcp_helper": schema.StringAttribute{
				MarkdownDescription: "List of DHCP IPv4 helper addresses",
				Computed:            true,
			},
			"ipv4_dhcp_helper_variable": schema.StringAttribute{
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
			"tunnel_interface_enabled": schema.BoolAttribute{
				MarkdownDescription: "Tunnel Interface on/off",
				Computed:            true,
			},
			"tunnel_per_tunnel_qos": schema.BoolAttribute{
				MarkdownDescription: "Per-tunnel Qos",
				Computed:            true,
			},
			"tunnel_per_tunnel_qos_variable": schema.StringAttribute{
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
			"tunnel_interface_color_description": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"tunnel_interface_color_description_variable": schema.StringAttribute{
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
			"tunnel_interface_color_restrict": schema.BoolAttribute{
				MarkdownDescription: "Restrict this TLOC behavior",
				Computed:            true,
			},
			"tunnel_interface_color_restrict_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_groups": schema.StringAttribute{
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
				MarkdownDescription: "Set the maximum number of control connections for this TLOC",
				Computed:            true,
			},
			"tunnel_interface_max_control_connections_variable": schema.StringAttribute{
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
			"tunnel_interface_exclude_controller_group_list": schema.StringAttribute{
				MarkdownDescription: "Exclude the following controller groups defined in this list",
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
			"tunnel_interface_full_port_hop": schema.BoolAttribute{
				MarkdownDescription: "Enable port hopping on the tunnel interface",
				Computed:            true,
			},
			"tunnel_interface_full_port_hop_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_port_hop": schema.BoolAttribute{
				MarkdownDescription: "The port hop functionality is deprecated for devices 17.18 and higher. Use the full-port-hop field instead",
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
			"tunnel_interface_network_broadcast": schema.BoolAttribute{
				MarkdownDescription: "Accept and respond to network-prefix-directed broadcasts)",
				Computed:            true,
			},
			"tunnel_interface_network_broadcast_variable": schema.StringAttribute{
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
			"tunnel_interface_allow_ntp": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny NTP",
				Computed:            true,
			},
			"tunnel_interface_allow_ntp_variable": schema.StringAttribute{
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
			"tunnel_interface_allow_netconf": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny NETCONF",
				Computed:            true,
			},
			"tunnel_interface_allow_netconf_variable": schema.StringAttribute{
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
			"tunnel_interface_allow_snmp": schema.BoolAttribute{
				MarkdownDescription: "Allow/Deny SNMP",
				Computed:            true,
			},
			"tunnel_interface_allow_snmp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_ipsec_encapsulation": schema.BoolAttribute{
				MarkdownDescription: "Encapsulation",
				Computed:            true,
			},
			"tunnel_interface_ipsec_preference": schema.Int64Attribute{
				MarkdownDescription: "Set preference for TLOC",
				Computed:            true,
			},
			"tunnel_interface_ipsec_preference_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_ipsec_weight": schema.Int64Attribute{
				MarkdownDescription: "Set weight for TLOC",
				Computed:            true,
			},
			"tunnel_interface_ipsec_weight_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_gre_encapsulation": schema.BoolAttribute{
				MarkdownDescription: "Encapsulation",
				Computed:            true,
			},
			"tunnel_interface_gre_preference": schema.Int64Attribute{
				MarkdownDescription: "Set preference for TLOC",
				Computed:            true,
			},
			"tunnel_interface_gre_preference_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tunnel_interface_gre_weight": schema.Int64Attribute{
				MarkdownDescription: "Set weight for TLOC",
				Computed:            true,
			},
			"tunnel_interface_gre_weight_variable": schema.StringAttribute{
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
			"multi_region_fabric_enable_core_region": schema.BoolAttribute{
				MarkdownDescription: "Enable Core Region",
				Computed:            true,
			},
			"multi_region_fabric_core_region": schema.StringAttribute{
				MarkdownDescription: "Core Region",
				Computed:            true,
			},
			"multi_region_fabric_enable_secondary_region": schema.BoolAttribute{
				MarkdownDescription: "Enable Secondary Region",
				Computed:            true,
			},
			"multi_region_fabric_secondary_region": schema.StringAttribute{
				MarkdownDescription: "Enable secondary region",
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
			"ip_mtu": schema.Int64Attribute{
				MarkdownDescription: "IP MTU for GigabitEthernet main <576..Interface MTU>, GigabitEthernet subinterface <576..9216>, Other Interfaces <576..2000> in bytes",
				Computed:            true,
			},
			"ip_mtu_variable": schema.StringAttribute{
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
			"tloc_extension": schema.StringAttribute{
				MarkdownDescription: "Extends a local TLOC to a remote node only for vpn 0",
				Computed:            true,
			},
			"tloc_extension_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"adaptive_qos": schema.BoolAttribute{
				MarkdownDescription: "Adaptive QoS",
				Computed:            true,
			},
			"adaptive_qos_period": schema.Int64Attribute{
				MarkdownDescription: "Adapt Period(Minutes)",
				Computed:            true,
			},
			"adaptive_qos_period_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"shaping_rate_upstream": schema.BoolAttribute{
				MarkdownDescription: "Shaping Rate Upstream",
				Computed:            true,
			},
			"min_shaping_rate_upstream": schema.Int64Attribute{
				MarkdownDescription: "Upstream min bandwidth limit (kbps)",
				Computed:            true,
			},
			"min_shaping_rate_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"max_shaping_rate_upstream": schema.Int64Attribute{
				MarkdownDescription: "Upstream max bandwidth limit (kbps)",
				Computed:            true,
			},
			"max_shaping_rate_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"default_shaping_rate_upstream": schema.Int64Attribute{
				MarkdownDescription: "Adaptive QoS default upstream bandwidth (kbps)",
				Computed:            true,
			},
			"default_shaping_rate_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"shaping_rate_downstream": schema.BoolAttribute{
				MarkdownDescription: "Shaping Rate Downstream",
				Computed:            true,
			},
			"min_shaping_rate_downstream": schema.Int64Attribute{
				MarkdownDescription: "Downstream min bandwidth limit (kbps)",
				Computed:            true,
			},
			"min_shaping_rate_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"max_shaping_rate_downstream": schema.Int64Attribute{
				MarkdownDescription: "Downstream max bandwidth limit (kbps)",
				Computed:            true,
			},
			"max_shaping_rate_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"default_shaping_rate_downstream": schema.Int64Attribute{
				MarkdownDescription: "Adaptive QoS default downstream bandwidth (kbps)",
				Computed:            true,
			},
			"default_shaping_rate_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"shaping_rate": schema.Int64Attribute{
				MarkdownDescription: "Shaping Rate (Kbps)",
				Computed:            true,
			},
			"shaping_rate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"acl_ipv4_egress_feature_id": schema.StringAttribute{
				MarkdownDescription: "IPv4 egress ACL feature ID. Variable references require Manager 20.18.1 or later.",
				Computed:            true,
			},
			"acl_ipv4_egress_feature_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"acl_ipv4_ingress_feature_id": schema.StringAttribute{
				MarkdownDescription: "IPv4 ingress ACL feature ID. Variable references require Manager 20.18.1 or later.",
				Computed:            true,
			},
			"acl_ipv4_ingress_feature_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"acl_ipv6_egress_feature_id": schema.StringAttribute{
				MarkdownDescription: "IPv6 egress ACL feature ID. Variable references require Manager 20.18.1 or later.",
				Computed:            true,
			},
			"acl_ipv6_egress_feature_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"acl_ipv6_ingress_feature_id": schema.StringAttribute{
				MarkdownDescription: "IPv6 ingress ACL feature ID. Variable references require Manager 20.18.1 or later.",
				Computed:            true,
			},
			"acl_ipv6_ingress_feature_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
		},
	}
}

func (d *TransportWANVPNInterfaceDSLIPoEProfileParcelDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *TransportWANVPNInterfaceDSLIPoEProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *TransportWANVPNInterfaceDSLIPoEProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config TransportWANVPNInterfaceDSLIPoE

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))
	if config.Id.IsNull() && !config.Name.IsNull() {
		// Look up parcel ID by name
		res, err := d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve parcels, got error: %s", err))
			return
		}
		found := false
		res.Get("data").ForEach(func(_, v gjson.Result) bool {
			if v.Get("payload.name").String() == config.Name.ValueString() {
				config.Id = types.StringValue(v.Get("parcelId").String())
				found = true
				return false
			}
			return true
		})
		if !found {
			resp.Diagnostics.AddError("Not Found", fmt.Sprintf("No parcel found with name: %s", config.Name.ValueString()))
			return
		}
	}

	res, err := d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res, true)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
