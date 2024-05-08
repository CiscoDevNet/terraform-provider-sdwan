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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type CiscoVPNInterface struct {
	Id                                                 types.String                                     `tfsdk:"id"`
	Version                                            types.Int64                                      `tfsdk:"version"`
	TemplateType                                       types.String                                     `tfsdk:"template_type"`
	Name                                               types.String                                     `tfsdk:"name"`
	Description                                        types.String                                     `tfsdk:"description"`
	DeviceTypes                                        types.Set                                        `tfsdk:"device_types"`
	InterfaceName                                      types.String                                     `tfsdk:"interface_name"`
	InterfaceNameVariable                              types.String                                     `tfsdk:"interface_name_variable"`
	InterfaceDescription                               types.String                                     `tfsdk:"interface_description"`
	InterfaceDescriptionVariable                       types.String                                     `tfsdk:"interface_description_variable"`
	Poe                                                types.Bool                                       `tfsdk:"poe"`
	PoeVariable                                        types.String                                     `tfsdk:"poe_variable"`
	Address                                            types.String                                     `tfsdk:"address"`
	AddressVariable                                    types.String                                     `tfsdk:"address_variable"`
	Ipv4SecondaryAddresses                             []CiscoVPNInterfaceIpv4SecondaryAddresses        `tfsdk:"ipv4_secondary_addresses"`
	Dhcp                                               types.Bool                                       `tfsdk:"dhcp"`
	DhcpVariable                                       types.String                                     `tfsdk:"dhcp_variable"`
	DhcpDistance                                       types.Int64                                      `tfsdk:"dhcp_distance"`
	DhcpDistanceVariable                               types.String                                     `tfsdk:"dhcp_distance_variable"`
	Ipv6Address                                        types.String                                     `tfsdk:"ipv6_address"`
	Ipv6AddressVariable                                types.String                                     `tfsdk:"ipv6_address_variable"`
	Dhcpv6                                             types.Bool                                       `tfsdk:"dhcpv6"`
	Dhcpv6Variable                                     types.String                                     `tfsdk:"dhcpv6_variable"`
	Ipv6SecondaryAddresses                             []CiscoVPNInterfaceIpv6SecondaryAddresses        `tfsdk:"ipv6_secondary_addresses"`
	Ipv6AccessLists                                    []CiscoVPNInterfaceIpv6AccessLists               `tfsdk:"ipv6_access_lists"`
	Ipv4DhcpHelper                                     types.Set                                        `tfsdk:"ipv4_dhcp_helper"`
	Ipv4DhcpHelperVariable                             types.String                                     `tfsdk:"ipv4_dhcp_helper_variable"`
	Ipv6DhcpHelpers                                    []CiscoVPNInterfaceIpv6DhcpHelpers               `tfsdk:"ipv6_dhcp_helpers"`
	Tracker                                            types.Set                                        `tfsdk:"tracker"`
	TrackerVariable                                    types.String                                     `tfsdk:"tracker_variable"`
	AutoBandwidthDetect                                types.Bool                                       `tfsdk:"auto_bandwidth_detect"`
	AutoBandwidthDetectVariable                        types.String                                     `tfsdk:"auto_bandwidth_detect_variable"`
	IperfServer                                        types.String                                     `tfsdk:"iperf_server"`
	IperfServerVariable                                types.String                                     `tfsdk:"iperf_server_variable"`
	Nat                                                types.Bool                                       `tfsdk:"nat"`
	NatType                                            types.String                                     `tfsdk:"nat_type"`
	NatTypeVariable                                    types.String                                     `tfsdk:"nat_type_variable"`
	UdpTimeout                                         types.Int64                                      `tfsdk:"udp_timeout"`
	UdpTimeoutVariable                                 types.String                                     `tfsdk:"udp_timeout_variable"`
	TcpTimeout                                         types.Int64                                      `tfsdk:"tcp_timeout"`
	TcpTimeoutVariable                                 types.String                                     `tfsdk:"tcp_timeout_variable"`
	NatPoolRangeStart                                  types.String                                     `tfsdk:"nat_pool_range_start"`
	NatPoolRangeStartVariable                          types.String                                     `tfsdk:"nat_pool_range_start_variable"`
	NatPoolRangeEnd                                    types.String                                     `tfsdk:"nat_pool_range_end"`
	NatPoolRangeEndVariable                            types.String                                     `tfsdk:"nat_pool_range_end_variable"`
	NatOverload                                        types.Bool                                       `tfsdk:"nat_overload"`
	NatOverloadVariable                                types.String                                     `tfsdk:"nat_overload_variable"`
	NatInsideSourceLoopbackInterface                   types.String                                     `tfsdk:"nat_inside_source_loopback_interface"`
	NatInsideSourceLoopbackInterfaceVariable           types.String                                     `tfsdk:"nat_inside_source_loopback_interface_variable"`
	NatPoolPrefixLength                                types.Int64                                      `tfsdk:"nat_pool_prefix_length"`
	NatPoolPrefixLengthVariable                        types.String                                     `tfsdk:"nat_pool_prefix_length_variable"`
	Ipv6Nat                                            types.Bool                                       `tfsdk:"ipv6_nat"`
	Ipv6NatVariable                                    types.String                                     `tfsdk:"ipv6_nat_variable"`
	Nat64Interface                                     types.Bool                                       `tfsdk:"nat64_interface"`
	Nat66Interface                                     types.Bool                                       `tfsdk:"nat66_interface"`
	StaticNat66Entries                                 []CiscoVPNInterfaceStaticNat66Entries            `tfsdk:"static_nat66_entries"`
	StaticNatEntries                                   []CiscoVPNInterfaceStaticNatEntries              `tfsdk:"static_nat_entries"`
	StaticPortForwardEntries                           []CiscoVPNInterfaceStaticPortForwardEntries      `tfsdk:"static_port_forward_entries"`
	EnableCoreRegion                                   types.Bool                                       `tfsdk:"enable_core_region"`
	CoreRegion                                         types.String                                     `tfsdk:"core_region"`
	CoreRegionVariable                                 types.String                                     `tfsdk:"core_region_variable"`
	SecondaryRegion                                    types.String                                     `tfsdk:"secondary_region"`
	SecondaryRegionVariable                            types.String                                     `tfsdk:"secondary_region_variable"`
	TunnelInterfaceEncapsulations                      []CiscoVPNInterfaceTunnelInterfaceEncapsulations `tfsdk:"tunnel_interface_encapsulations"`
	TunnelInterfaceBorder                              types.Bool                                       `tfsdk:"tunnel_interface_border"`
	TunnelInterfaceBorderVariable                      types.String                                     `tfsdk:"tunnel_interface_border_variable"`
	TunnelQosMode                                      types.String                                     `tfsdk:"tunnel_qos_mode"`
	TunnelQosModeVariable                              types.String                                     `tfsdk:"tunnel_qos_mode_variable"`
	TunnelBandwidth                                    types.Int64                                      `tfsdk:"tunnel_bandwidth"`
	TunnelBandwidthVariable                            types.String                                     `tfsdk:"tunnel_bandwidth_variable"`
	TunnelInterfaceGroups                              types.Set                                        `tfsdk:"tunnel_interface_groups"`
	TunnelInterfaceGroupsVariable                      types.String                                     `tfsdk:"tunnel_interface_groups_variable"`
	TunnelInterfaceColor                               types.String                                     `tfsdk:"tunnel_interface_color"`
	TunnelInterfaceColorVariable                       types.String                                     `tfsdk:"tunnel_interface_color_variable"`
	TunnelInterfaceMaxControlConnections               types.Int64                                      `tfsdk:"tunnel_interface_max_control_connections"`
	TunnelInterfaceMaxControlConnectionsVariable       types.String                                     `tfsdk:"tunnel_interface_max_control_connections_variable"`
	TunnelInterfaceControlConnections                  types.Bool                                       `tfsdk:"tunnel_interface_control_connections"`
	TunnelInterfaceControlConnectionsVariable          types.String                                     `tfsdk:"tunnel_interface_control_connections_variable"`
	TunnelInterfaceVbondAsStunServer                   types.Bool                                       `tfsdk:"tunnel_interface_vbond_as_stun_server"`
	TunnelInterfaceVbondAsStunServerVariable           types.String                                     `tfsdk:"tunnel_interface_vbond_as_stun_server_variable"`
	TunnelInterfaceExcludeControllerGroupList          types.Set                                        `tfsdk:"tunnel_interface_exclude_controller_group_list"`
	TunnelInterfaceExcludeControllerGroupListVariable  types.String                                     `tfsdk:"tunnel_interface_exclude_controller_group_list_variable"`
	TunnelInterfaceVmanageConnectionPreference         types.Int64                                      `tfsdk:"tunnel_interface_vmanage_connection_preference"`
	TunnelInterfaceVmanageConnectionPreferenceVariable types.String                                     `tfsdk:"tunnel_interface_vmanage_connection_preference_variable"`
	TunnelInterfacePortHop                             types.Bool                                       `tfsdk:"tunnel_interface_port_hop"`
	TunnelInterfacePortHopVariable                     types.String                                     `tfsdk:"tunnel_interface_port_hop_variable"`
	TunnelInterfaceColorRestrict                       types.Bool                                       `tfsdk:"tunnel_interface_color_restrict"`
	TunnelInterfaceColorRestrictVariable               types.String                                     `tfsdk:"tunnel_interface_color_restrict_variable"`
	TunnelInterfaceGreTunnelDestinationIp              types.String                                     `tfsdk:"tunnel_interface_gre_tunnel_destination_ip"`
	TunnelInterfaceGreTunnelDestinationIpVariable      types.String                                     `tfsdk:"tunnel_interface_gre_tunnel_destination_ip_variable"`
	TunnelInterfaceCarrier                             types.String                                     `tfsdk:"tunnel_interface_carrier"`
	TunnelInterfaceCarrierVariable                     types.String                                     `tfsdk:"tunnel_interface_carrier_variable"`
	TunnelInterfaceNatRefreshInterval                  types.Int64                                      `tfsdk:"tunnel_interface_nat_refresh_interval"`
	TunnelInterfaceNatRefreshIntervalVariable          types.String                                     `tfsdk:"tunnel_interface_nat_refresh_interval_variable"`
	TunnelInterfaceHelloInterval                       types.Int64                                      `tfsdk:"tunnel_interface_hello_interval"`
	TunnelInterfaceHelloIntervalVariable               types.String                                     `tfsdk:"tunnel_interface_hello_interval_variable"`
	TunnelInterfaceHelloTolerance                      types.Int64                                      `tfsdk:"tunnel_interface_hello_tolerance"`
	TunnelInterfaceHelloToleranceVariable              types.String                                     `tfsdk:"tunnel_interface_hello_tolerance_variable"`
	TunnelInterfaceBindLoopbackTunnel                  types.String                                     `tfsdk:"tunnel_interface_bind_loopback_tunnel"`
	TunnelInterfaceBindLoopbackTunnelVariable          types.String                                     `tfsdk:"tunnel_interface_bind_loopback_tunnel_variable"`
	TunnelInterfaceLastResortCircuit                   types.Bool                                       `tfsdk:"tunnel_interface_last_resort_circuit"`
	TunnelInterfaceLastResortCircuitVariable           types.String                                     `tfsdk:"tunnel_interface_last_resort_circuit_variable"`
	TunnelInterfaceLowBandwidthLink                    types.Bool                                       `tfsdk:"tunnel_interface_low_bandwidth_link"`
	TunnelInterfaceLowBandwidthLinkVariable            types.String                                     `tfsdk:"tunnel_interface_low_bandwidth_link_variable"`
	TunnelInterfaceTunnelTcpMss                        types.Int64                                      `tfsdk:"tunnel_interface_tunnel_tcp_mss"`
	TunnelInterfaceTunnelTcpMssVariable                types.String                                     `tfsdk:"tunnel_interface_tunnel_tcp_mss_variable"`
	TunnelInterfaceClearDontFragment                   types.Bool                                       `tfsdk:"tunnel_interface_clear_dont_fragment"`
	TunnelInterfaceClearDontFragmentVariable           types.String                                     `tfsdk:"tunnel_interface_clear_dont_fragment_variable"`
	TunnelInterfacePropagateSgt                        types.Bool                                       `tfsdk:"tunnel_interface_propagate_sgt"`
	TunnelInterfacePropagateSgtVariable                types.String                                     `tfsdk:"tunnel_interface_propagate_sgt_variable"`
	TunnelInterfaceNetworkBroadcast                    types.Bool                                       `tfsdk:"tunnel_interface_network_broadcast"`
	TunnelInterfaceNetworkBroadcastVariable            types.String                                     `tfsdk:"tunnel_interface_network_broadcast_variable"`
	TunnelInterfaceAllowAll                            types.Bool                                       `tfsdk:"tunnel_interface_allow_all"`
	TunnelInterfaceAllowAllVariable                    types.String                                     `tfsdk:"tunnel_interface_allow_all_variable"`
	TunnelInterfaceAllowBgp                            types.Bool                                       `tfsdk:"tunnel_interface_allow_bgp"`
	TunnelInterfaceAllowBgpVariable                    types.String                                     `tfsdk:"tunnel_interface_allow_bgp_variable"`
	TunnelInterfaceAllowDhcp                           types.Bool                                       `tfsdk:"tunnel_interface_allow_dhcp"`
	TunnelInterfaceAllowDhcpVariable                   types.String                                     `tfsdk:"tunnel_interface_allow_dhcp_variable"`
	TunnelInterfaceAllowDns                            types.Bool                                       `tfsdk:"tunnel_interface_allow_dns"`
	TunnelInterfaceAllowDnsVariable                    types.String                                     `tfsdk:"tunnel_interface_allow_dns_variable"`
	TunnelInterfaceAllowIcmp                           types.Bool                                       `tfsdk:"tunnel_interface_allow_icmp"`
	TunnelInterfaceAllowIcmpVariable                   types.String                                     `tfsdk:"tunnel_interface_allow_icmp_variable"`
	TunnelInterfaceAllowSsh                            types.Bool                                       `tfsdk:"tunnel_interface_allow_ssh"`
	TunnelInterfaceAllowSshVariable                    types.String                                     `tfsdk:"tunnel_interface_allow_ssh_variable"`
	TunnelInterfaceAllowNetconf                        types.Bool                                       `tfsdk:"tunnel_interface_allow_netconf"`
	TunnelInterfaceAllowNetconfVariable                types.String                                     `tfsdk:"tunnel_interface_allow_netconf_variable"`
	TunnelInterfaceAllowNtp                            types.Bool                                       `tfsdk:"tunnel_interface_allow_ntp"`
	TunnelInterfaceAllowNtpVariable                    types.String                                     `tfsdk:"tunnel_interface_allow_ntp_variable"`
	TunnelInterfaceAllowOspf                           types.Bool                                       `tfsdk:"tunnel_interface_allow_ospf"`
	TunnelInterfaceAllowOspfVariable                   types.String                                     `tfsdk:"tunnel_interface_allow_ospf_variable"`
	TunnelInterfaceAllowStun                           types.Bool                                       `tfsdk:"tunnel_interface_allow_stun"`
	TunnelInterfaceAllowStunVariable                   types.String                                     `tfsdk:"tunnel_interface_allow_stun_variable"`
	TunnelInterfaceAllowSnmp                           types.Bool                                       `tfsdk:"tunnel_interface_allow_snmp"`
	TunnelInterfaceAllowSnmpVariable                   types.String                                     `tfsdk:"tunnel_interface_allow_snmp_variable"`
	TunnelInterfaceAllowHttps                          types.Bool                                       `tfsdk:"tunnel_interface_allow_https"`
	TunnelInterfaceAllowHttpsVariable                  types.String                                     `tfsdk:"tunnel_interface_allow_https_variable"`
	MediaType                                          types.String                                     `tfsdk:"media_type"`
	MediaTypeVariable                                  types.String                                     `tfsdk:"media_type_variable"`
	InterfaceMtu                                       types.Int64                                      `tfsdk:"interface_mtu"`
	InterfaceMtuVariable                               types.String                                     `tfsdk:"interface_mtu_variable"`
	IpMtu                                              types.Int64                                      `tfsdk:"ip_mtu"`
	IpMtuVariable                                      types.String                                     `tfsdk:"ip_mtu_variable"`
	TcpMssAdjust                                       types.Int64                                      `tfsdk:"tcp_mss_adjust"`
	TcpMssAdjustVariable                               types.String                                     `tfsdk:"tcp_mss_adjust_variable"`
	TlocExtension                                      types.String                                     `tfsdk:"tloc_extension"`
	TlocExtensionVariable                              types.String                                     `tfsdk:"tloc_extension_variable"`
	LoadInterval                                       types.Int64                                      `tfsdk:"load_interval"`
	LoadIntervalVariable                               types.String                                     `tfsdk:"load_interval_variable"`
	GreTunnelSourceIp                                  types.String                                     `tfsdk:"gre_tunnel_source_ip"`
	GreTunnelSourceIpVariable                          types.String                                     `tfsdk:"gre_tunnel_source_ip_variable"`
	GreTunnelXconnect                                  types.String                                     `tfsdk:"gre_tunnel_xconnect"`
	GreTunnelXconnectVariable                          types.String                                     `tfsdk:"gre_tunnel_xconnect_variable"`
	MacAddress                                         types.String                                     `tfsdk:"mac_address"`
	MacAddressVariable                                 types.String                                     `tfsdk:"mac_address_variable"`
	Speed                                              types.String                                     `tfsdk:"speed"`
	SpeedVariable                                      types.String                                     `tfsdk:"speed_variable"`
	Duplex                                             types.String                                     `tfsdk:"duplex"`
	DuplexVariable                                     types.String                                     `tfsdk:"duplex_variable"`
	Shutdown                                           types.Bool                                       `tfsdk:"shutdown"`
	ShutdownVariable                                   types.String                                     `tfsdk:"shutdown_variable"`
	ArpTimeout                                         types.Int64                                      `tfsdk:"arp_timeout"`
	ArpTimeoutVariable                                 types.String                                     `tfsdk:"arp_timeout_variable"`
	Autonegotiate                                      types.Bool                                       `tfsdk:"autonegotiate"`
	AutonegotiateVariable                              types.String                                     `tfsdk:"autonegotiate_variable"`
	IpDirectedBroadcast                                types.Bool                                       `tfsdk:"ip_directed_broadcast"`
	IpDirectedBroadcastVariable                        types.String                                     `tfsdk:"ip_directed_broadcast_variable"`
	IcmpRedirectDisable                                types.Bool                                       `tfsdk:"icmp_redirect_disable"`
	IcmpRedirectDisableVariable                        types.String                                     `tfsdk:"icmp_redirect_disable_variable"`
	QosAdaptivePeriod                                  types.Int64                                      `tfsdk:"qos_adaptive_period"`
	QosAdaptivePeriodVariable                          types.String                                     `tfsdk:"qos_adaptive_period_variable"`
	QosAdaptiveBandwidthDownstream                     types.Int64                                      `tfsdk:"qos_adaptive_bandwidth_downstream"`
	QosAdaptiveBandwidthDownstreamVariable             types.String                                     `tfsdk:"qos_adaptive_bandwidth_downstream_variable"`
	QosAdaptiveMinDownstream                           types.Int64                                      `tfsdk:"qos_adaptive_min_downstream"`
	QosAdaptiveMinDownstreamVariable                   types.String                                     `tfsdk:"qos_adaptive_min_downstream_variable"`
	QosAdaptiveMaxDownstream                           types.Int64                                      `tfsdk:"qos_adaptive_max_downstream"`
	QosAdaptiveMaxDownstreamVariable                   types.String                                     `tfsdk:"qos_adaptive_max_downstream_variable"`
	QosAdaptiveBandwidthUpstream                       types.Int64                                      `tfsdk:"qos_adaptive_bandwidth_upstream"`
	QosAdaptiveBandwidthUpstreamVariable               types.String                                     `tfsdk:"qos_adaptive_bandwidth_upstream_variable"`
	QosAdaptiveMinUpstream                             types.Int64                                      `tfsdk:"qos_adaptive_min_upstream"`
	QosAdaptiveMinUpstreamVariable                     types.String                                     `tfsdk:"qos_adaptive_min_upstream_variable"`
	QosAdaptiveMaxUpstream                             types.Int64                                      `tfsdk:"qos_adaptive_max_upstream"`
	QosAdaptiveMaxUpstreamVariable                     types.String                                     `tfsdk:"qos_adaptive_max_upstream_variable"`
	ShapingRate                                        types.Int64                                      `tfsdk:"shaping_rate"`
	ShapingRateVariable                                types.String                                     `tfsdk:"shaping_rate_variable"`
	QosMap                                             types.String                                     `tfsdk:"qos_map"`
	QosMapVariable                                     types.String                                     `tfsdk:"qos_map_variable"`
	QosMapVpn                                          types.String                                     `tfsdk:"qos_map_vpn"`
	QosMapVpnVariable                                  types.String                                     `tfsdk:"qos_map_vpn_variable"`
	BandwidthUpstream                                  types.Int64                                      `tfsdk:"bandwidth_upstream"`
	BandwidthUpstreamVariable                          types.String                                     `tfsdk:"bandwidth_upstream_variable"`
	BandwidthDownstream                                types.Int64                                      `tfsdk:"bandwidth_downstream"`
	BandwidthDownstreamVariable                        types.String                                     `tfsdk:"bandwidth_downstream_variable"`
	BlockNonSourceIp                                   types.Bool                                       `tfsdk:"block_non_source_ip"`
	BlockNonSourceIpVariable                           types.String                                     `tfsdk:"block_non_source_ip_variable"`
	RewriteRuleName                                    types.String                                     `tfsdk:"rewrite_rule_name"`
	RewriteRuleNameVariable                            types.String                                     `tfsdk:"rewrite_rule_name_variable"`
	AccessLists                                        []CiscoVPNInterfaceAccessLists                   `tfsdk:"access_lists"`
	StaticArps                                         []CiscoVPNInterfaceStaticArps                    `tfsdk:"static_arps"`
	Ipv4Vrrps                                          []CiscoVPNInterfaceIpv4Vrrps                     `tfsdk:"ipv4_vrrps"`
	Ipv6Vrrps                                          []CiscoVPNInterfaceIpv6Vrrps                     `tfsdk:"ipv6_vrrps"`
	PropagateSgt                                       types.Bool                                       `tfsdk:"propagate_sgt"`
	StaticSgt                                          types.Int64                                      `tfsdk:"static_sgt"`
	StaticSgtVariable                                  types.String                                     `tfsdk:"static_sgt_variable"`
	StaticSgtTrusted                                   types.Bool                                       `tfsdk:"static_sgt_trusted"`
	EnableSgt                                          types.Bool                                       `tfsdk:"enable_sgt"`
	SgtEnforcement                                     types.Bool                                       `tfsdk:"sgt_enforcement"`
	SgtEnforcementSgt                                  types.Int64                                      `tfsdk:"sgt_enforcement_sgt"`
	SgtEnforcementSgtVariable                          types.String                                     `tfsdk:"sgt_enforcement_sgt_variable"`
}

type CiscoVPNInterfaceIpv4SecondaryAddresses struct {
	Optional        types.Bool   `tfsdk:"optional"`
	Address         types.String `tfsdk:"address"`
	AddressVariable types.String `tfsdk:"address_variable"`
}

type CiscoVPNInterfaceIpv6SecondaryAddresses struct {
	Optional        types.Bool   `tfsdk:"optional"`
	Address         types.String `tfsdk:"address"`
	AddressVariable types.String `tfsdk:"address_variable"`
}

type CiscoVPNInterfaceIpv6AccessLists struct {
	Optional        types.Bool   `tfsdk:"optional"`
	Direction       types.String `tfsdk:"direction"`
	AclName         types.String `tfsdk:"acl_name"`
	AclNameVariable types.String `tfsdk:"acl_name_variable"`
}

type CiscoVPNInterfaceIpv6DhcpHelpers struct {
	Optional        types.Bool   `tfsdk:"optional"`
	Address         types.String `tfsdk:"address"`
	AddressVariable types.String `tfsdk:"address_variable"`
	VpnId           types.Int64  `tfsdk:"vpn_id"`
	VpnIdVariable   types.String `tfsdk:"vpn_id_variable"`
}

type CiscoVPNInterfaceStaticNat66Entries struct {
	Optional                       types.Bool   `tfsdk:"optional"`
	SourcePrefix                   types.String `tfsdk:"source_prefix"`
	SourcePrefixVariable           types.String `tfsdk:"source_prefix_variable"`
	TranslatedSourcePrefix         types.String `tfsdk:"translated_source_prefix"`
	TranslatedSourcePrefixVariable types.String `tfsdk:"translated_source_prefix_variable"`
	SourceVpnId                    types.Int64  `tfsdk:"source_vpn_id"`
	SourceVpnIdVariable            types.String `tfsdk:"source_vpn_id_variable"`
}

type CiscoVPNInterfaceStaticNatEntries struct {
	Optional                   types.Bool   `tfsdk:"optional"`
	SourceIp                   types.String `tfsdk:"source_ip"`
	SourceIpVariable           types.String `tfsdk:"source_ip_variable"`
	TranslateIp                types.String `tfsdk:"translate_ip"`
	TranslateIpVariable        types.String `tfsdk:"translate_ip_variable"`
	StaticNatDirection         types.String `tfsdk:"static_nat_direction"`
	StaticNatDirectionVariable types.String `tfsdk:"static_nat_direction_variable"`
	SourceVpnId                types.Int64  `tfsdk:"source_vpn_id"`
	SourceVpnIdVariable        types.String `tfsdk:"source_vpn_id_variable"`
}

type CiscoVPNInterfaceStaticPortForwardEntries struct {
	Optional                   types.Bool   `tfsdk:"optional"`
	SourceIp                   types.String `tfsdk:"source_ip"`
	SourceIpVariable           types.String `tfsdk:"source_ip_variable"`
	TranslateIp                types.String `tfsdk:"translate_ip"`
	TranslateIpVariable        types.String `tfsdk:"translate_ip_variable"`
	StaticNatDirection         types.String `tfsdk:"static_nat_direction"`
	StaticNatDirectionVariable types.String `tfsdk:"static_nat_direction_variable"`
	SourcePort                 types.Int64  `tfsdk:"source_port"`
	SourcePortVariable         types.String `tfsdk:"source_port_variable"`
	TranslatePort              types.Int64  `tfsdk:"translate_port"`
	TranslatePortVariable      types.String `tfsdk:"translate_port_variable"`
	Protocol                   types.String `tfsdk:"protocol"`
	ProtocolVariable           types.String `tfsdk:"protocol_variable"`
	SourceVpnId                types.Int64  `tfsdk:"source_vpn_id"`
	SourceVpnIdVariable        types.String `tfsdk:"source_vpn_id_variable"`
}

type CiscoVPNInterfaceTunnelInterfaceEncapsulations struct {
	Optional           types.Bool   `tfsdk:"optional"`
	Encapsulation      types.String `tfsdk:"encapsulation"`
	Preference         types.Int64  `tfsdk:"preference"`
	PreferenceVariable types.String `tfsdk:"preference_variable"`
	Weight             types.Int64  `tfsdk:"weight"`
	WeightVariable     types.String `tfsdk:"weight_variable"`
}

type CiscoVPNInterfaceAccessLists struct {
	Optional        types.Bool   `tfsdk:"optional"`
	Direction       types.String `tfsdk:"direction"`
	AclName         types.String `tfsdk:"acl_name"`
	AclNameVariable types.String `tfsdk:"acl_name_variable"`
}

type CiscoVPNInterfaceStaticArps struct {
	Optional          types.Bool   `tfsdk:"optional"`
	IpAddress         types.String `tfsdk:"ip_address"`
	IpAddressVariable types.String `tfsdk:"ip_address_variable"`
	Mac               types.String `tfsdk:"mac"`
	MacVariable       types.String `tfsdk:"mac_variable"`
}

type CiscoVPNInterfaceIpv4Vrrps struct {
	Optional                          types.Bool                                         `tfsdk:"optional"`
	GroupId                           types.Int64                                        `tfsdk:"group_id"`
	GroupIdVariable                   types.String                                       `tfsdk:"group_id_variable"`
	Priority                          types.Int64                                        `tfsdk:"priority"`
	PriorityVariable                  types.String                                       `tfsdk:"priority_variable"`
	Timer                             types.Int64                                        `tfsdk:"timer"`
	TimerVariable                     types.String                                       `tfsdk:"timer_variable"`
	TrackOmp                          types.Bool                                         `tfsdk:"track_omp"`
	TrackPrefixList                   types.String                                       `tfsdk:"track_prefix_list"`
	TrackPrefixListVariable           types.String                                       `tfsdk:"track_prefix_list_variable"`
	IpAddress                         types.String                                       `tfsdk:"ip_address"`
	IpAddressVariable                 types.String                                       `tfsdk:"ip_address_variable"`
	Ipv4SecondaryAddresses            []CiscoVPNInterfaceIpv4VrrpsIpv4SecondaryAddresses `tfsdk:"ipv4_secondary_addresses"`
	TlocPreferenceChange              types.Bool                                         `tfsdk:"tloc_preference_change"`
	TlocPreferenceChangeValue         types.Int64                                        `tfsdk:"tloc_preference_change_value"`
	TlocPreferenceChangeValueVariable types.String                                       `tfsdk:"tloc_preference_change_value_variable"`
	TrackingObjects                   []CiscoVPNInterfaceIpv4VrrpsTrackingObjects        `tfsdk:"tracking_objects"`
}

type CiscoVPNInterfaceIpv6Vrrps struct {
	Optional                types.Bool                                `tfsdk:"optional"`
	GroupId                 types.Int64                               `tfsdk:"group_id"`
	GroupIdVariable         types.String                              `tfsdk:"group_id_variable"`
	Priority                types.Int64                               `tfsdk:"priority"`
	PriorityVariable        types.String                              `tfsdk:"priority_variable"`
	Timer                   types.Int64                               `tfsdk:"timer"`
	TimerVariable           types.String                              `tfsdk:"timer_variable"`
	TrackOmp                types.Bool                                `tfsdk:"track_omp"`
	TrackOmpVariable        types.String                              `tfsdk:"track_omp_variable"`
	TrackPrefixList         types.String                              `tfsdk:"track_prefix_list"`
	TrackPrefixListVariable types.String                              `tfsdk:"track_prefix_list_variable"`
	Ipv6Addresses           []CiscoVPNInterfaceIpv6VrrpsIpv6Addresses `tfsdk:"ipv6_addresses"`
}

type CiscoVPNInterfaceIpv4VrrpsIpv4SecondaryAddresses struct {
	Optional          types.Bool   `tfsdk:"optional"`
	IpAddress         types.String `tfsdk:"ip_address"`
	IpAddressVariable types.String `tfsdk:"ip_address_variable"`
}
type CiscoVPNInterfaceIpv4VrrpsTrackingObjects struct {
	Optional               types.Bool   `tfsdk:"optional"`
	TrackerId              types.Int64  `tfsdk:"tracker_id"`
	TrackerIdVariable      types.String `tfsdk:"tracker_id_variable"`
	TrackAction            types.String `tfsdk:"track_action"`
	TrackActionVariable    types.String `tfsdk:"track_action_variable"`
	DecrementValue         types.Int64  `tfsdk:"decrement_value"`
	DecrementValueVariable types.String `tfsdk:"decrement_value_variable"`
}

type CiscoVPNInterfaceIpv6VrrpsIpv6Addresses struct {
	Optional              types.Bool   `tfsdk:"optional"`
	Ipv6LinkLocal         types.String `tfsdk:"ipv6_link_local"`
	Ipv6LinkLocalVariable types.String `tfsdk:"ipv6_link_local_variable"`
	Prefix                types.String `tfsdk:"prefix"`
	PrefixVariable        types.String `tfsdk:"prefix_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoVPNInterface) getModel() string {
	return "cisco_vpn_interface"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CiscoVPNInterface) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_vpn_interface")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.InterfaceNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"if-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"if-name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"if-name."+"vipVariableName", data.InterfaceNameVariable.ValueString())
	} else if data.InterfaceName.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"if-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"if-name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"if-name."+"vipValue", data.InterfaceName.ValueString())
	}

	if !data.InterfaceDescriptionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"description."+"vipVariableName", data.InterfaceDescriptionVariable.ValueString())
	} else if data.InterfaceDescription.IsNull() {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"description."+"vipValue", data.InterfaceDescription.ValueString())
	}

	if !data.PoeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"poe."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"poe."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"poe."+"vipVariableName", data.PoeVariable.ValueString())
	} else if data.Poe.IsNull() {
		body, _ = sjson.Set(body, path+"poe."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"poe."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"poe."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"poe."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"poe."+"vipValue", strconv.FormatBool(data.Poe.ValueBool()))
	}

	if !data.AddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ip.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.address."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ip.address."+"vipVariableName", data.AddressVariable.ValueString())
	} else if data.Address.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"ip.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip.address."+"vipValue", data.Address.ValueString())
	}
	if len(data.Ipv4SecondaryAddresses) > 0 {
		body, _ = sjson.Set(body, path+"ip.secondary-address."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ip.secondary-address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip.secondary-address."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"ip.secondary-address."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ip.secondary-address."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ip.secondary-address."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ip.secondary-address."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"ip.secondary-address."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv4SecondaryAddresses {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "address")

		if !item.AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipVariableName", item.AddressVariable.ValueString())
		} else if item.Address.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipValue", item.Address.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ip.secondary-address."+"vipValue.-1", itemBody)
	}

	if !data.DhcpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ip.dhcp-client."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.dhcp-client."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ip.dhcp-client."+"vipVariableName", data.DhcpVariable.ValueString())
	} else if data.Dhcp.IsNull() {
		body, _ = sjson.Set(body, path+"ip.dhcp-client."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.dhcp-client."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ip.dhcp-client."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.dhcp-client."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip.dhcp-client."+"vipValue", strconv.FormatBool(data.Dhcp.ValueBool()))
	}

	if !data.DhcpDistanceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ip.dhcp-distance."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.dhcp-distance."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ip.dhcp-distance."+"vipVariableName", data.DhcpDistanceVariable.ValueString())
	} else if data.DhcpDistance.IsNull() {
		body, _ = sjson.Set(body, path+"ip.dhcp-distance."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.dhcp-distance."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ip.dhcp-distance."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.dhcp-distance."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip.dhcp-distance."+"vipValue", data.DhcpDistance.ValueInt64())
	}

	if !data.Ipv6AddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipVariableName", data.Ipv6AddressVariable.ValueString())
	} else if data.Ipv6Address.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6.address."+"vipValue", data.Ipv6Address.ValueString())
	}

	if !data.Dhcpv6Variable.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6.dhcp-client."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-client."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-client."+"vipVariableName", data.Dhcpv6Variable.ValueString())
	} else if data.Dhcpv6.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6.dhcp-client."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-client."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipv6.dhcp-client."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-client."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-client."+"vipValue", strconv.FormatBool(data.Dhcpv6.ValueBool()))
	}
	if len(data.Ipv6SecondaryAddresses) > 0 {
		body, _ = sjson.Set(body, path+"ipv6.secondary-address."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6.secondary-address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6.secondary-address."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"ipv6.secondary-address."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ipv6.secondary-address."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6.secondary-address."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ipv6.secondary-address."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"ipv6.secondary-address."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv6SecondaryAddresses {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "address")

		if !item.AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipVariableName", item.AddressVariable.ValueString())
		} else if item.Address.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipValue", item.Address.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ipv6.secondary-address."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv6AccessLists) > 0 {
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipPrimaryKey", []string{"direction"})
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipPrimaryKey", []string{"direction"})
		body, _ = sjson.Set(body, path+"ipv6.access-list."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv6AccessLists {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "direction")
		if item.Direction.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipValue", item.Direction.ValueString())
		}
		itemAttributes = append(itemAttributes, "acl-name")

		if !item.AclNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipVariableName", item.AclNameVariable.ValueString())
		} else if item.AclName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipValue", item.AclName.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ipv6.access-list."+"vipValue.-1", itemBody)
	}

	if !data.Ipv4DhcpHelperVariable.IsNull() {
		body, _ = sjson.Set(body, path+"dhcp-helper."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"dhcp-helper."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"dhcp-helper."+"vipVariableName", data.Ipv4DhcpHelperVariable.ValueString())
	} else if data.Ipv4DhcpHelper.IsNull() {
		body, _ = sjson.Set(body, path+"dhcp-helper."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"dhcp-helper."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"dhcp-helper."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"dhcp-helper."+"vipType", "constant")
		var values []string
		data.Ipv4DhcpHelper.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"dhcp-helper."+"vipValue", values)
	}
	if len(data.Ipv6DhcpHelpers) > 0 {
		body, _ = sjson.Set(body, path+"ipv6.dhcp-helper-v6."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-helper-v6."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-helper-v6."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"ipv6.dhcp-helper-v6."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ipv6.dhcp-helper-v6."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-helper-v6."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ipv6.dhcp-helper-v6."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"ipv6.dhcp-helper-v6."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv6DhcpHelpers {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "address")

		if !item.AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipVariableName", item.AddressVariable.ValueString())
		} else if item.Address.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipValue", item.Address.ValueString())
		}
		itemAttributes = append(itemAttributes, "vpn")

		if !item.VpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipVariableName", item.VpnIdVariable.ValueString())
		} else if item.VpnId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", item.VpnId.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ipv6.dhcp-helper-v6."+"vipValue.-1", itemBody)
	}

	if !data.TrackerVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tracker."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tracker."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tracker."+"vipVariableName", data.TrackerVariable.ValueString())
	} else if data.Tracker.IsNull() {
		body, _ = sjson.Set(body, path+"tracker."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tracker."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"tracker."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tracker."+"vipType", "constant")
		var values []string
		data.Tracker.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"tracker."+"vipValue", values)
	}

	if !data.AutoBandwidthDetectVariable.IsNull() {
		body, _ = sjson.Set(body, path+"auto-bandwidth-detect."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"auto-bandwidth-detect."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"auto-bandwidth-detect."+"vipVariableName", data.AutoBandwidthDetectVariable.ValueString())
	} else if data.AutoBandwidthDetect.IsNull() {
		body, _ = sjson.Set(body, path+"auto-bandwidth-detect."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"auto-bandwidth-detect."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"auto-bandwidth-detect."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"auto-bandwidth-detect."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"auto-bandwidth-detect."+"vipValue", strconv.FormatBool(data.AutoBandwidthDetect.ValueBool()))
	}

	if !data.IperfServerVariable.IsNull() {
		body, _ = sjson.Set(body, path+"iperf-server."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"iperf-server."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"iperf-server."+"vipVariableName", data.IperfServerVariable.ValueString())
	} else if data.IperfServer.IsNull() {
		body, _ = sjson.Set(body, path+"iperf-server."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"iperf-server."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"iperf-server."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"iperf-server."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"iperf-server."+"vipValue", data.IperfServer.ValueString())
	}
	if !data.Nat.IsNull() {
		if data.Nat.ValueBool() {
			if !gjson.Get(body, path+"nat").Exists() {
				body, _ = sjson.Set(body, path+"nat", map[string]interface{}{})
			}
		} else {
			body, _ = sjson.Set(body, path+"nat."+"vipObjectType", "node-only")
			body, _ = sjson.Set(body, path+"nat."+"vipType", "ignore")
		}
	}

	if !data.NatTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nat.nat-choice."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.nat-choice."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"nat.nat-choice."+"vipVariableName", data.NatTypeVariable.ValueString())
	} else if data.NatType.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"nat.nat-choice."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.nat-choice."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.nat-choice."+"vipValue", data.NatType.ValueString())
	}

	if !data.UdpTimeoutVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nat.udp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.udp-timeout."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"nat.udp-timeout."+"vipVariableName", data.UdpTimeoutVariable.ValueString())
	} else if data.UdpTimeout.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"nat.udp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.udp-timeout."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.udp-timeout."+"vipValue", data.UdpTimeout.ValueInt64())
	}

	if !data.TcpTimeoutVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nat.tcp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.tcp-timeout."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"nat.tcp-timeout."+"vipVariableName", data.TcpTimeoutVariable.ValueString())
	} else if data.TcpTimeout.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"nat.tcp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.tcp-timeout."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.tcp-timeout."+"vipValue", data.TcpTimeout.ValueInt64())
	}

	if !data.NatPoolRangeStartVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nat.natpool.range-start."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.natpool.range-start."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"nat.natpool.range-start."+"vipVariableName", data.NatPoolRangeStartVariable.ValueString())
	} else if data.NatPoolRangeStart.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"nat.natpool.range-start."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.natpool.range-start."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.natpool.range-start."+"vipValue", data.NatPoolRangeStart.ValueString())
	}

	if !data.NatPoolRangeEndVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nat.natpool.range-end."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.natpool.range-end."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"nat.natpool.range-end."+"vipVariableName", data.NatPoolRangeEndVariable.ValueString())
	} else if data.NatPoolRangeEnd.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"nat.natpool.range-end."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.natpool.range-end."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.natpool.range-end."+"vipValue", data.NatPoolRangeEnd.ValueString())
	}

	if !data.NatOverloadVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nat.overload."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.overload."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"nat.overload."+"vipVariableName", data.NatOverloadVariable.ValueString())
	} else if data.NatOverload.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"nat.overload."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.overload."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.overload."+"vipValue", strconv.FormatBool(data.NatOverload.ValueBool()))
	}

	if !data.NatInsideSourceLoopbackInterfaceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nat.interface.loopback-interface."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.interface.loopback-interface."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"nat.interface.loopback-interface."+"vipVariableName", data.NatInsideSourceLoopbackInterfaceVariable.ValueString())
	} else if data.NatInsideSourceLoopbackInterface.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"nat.interface.loopback-interface."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.interface.loopback-interface."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.interface.loopback-interface."+"vipValue", data.NatInsideSourceLoopbackInterface.ValueString())
	}

	if !data.NatPoolPrefixLengthVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nat.natpool.prefix-length."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.natpool.prefix-length."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"nat.natpool.prefix-length."+"vipVariableName", data.NatPoolPrefixLengthVariable.ValueString())
	} else if data.NatPoolPrefixLength.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"nat.natpool.prefix-length."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat.natpool.prefix-length."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.natpool.prefix-length."+"vipValue", data.NatPoolPrefixLength.ValueInt64())
	}

	if !data.Ipv6NatVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nat64.enable."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"nat64.enable."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"nat64.enable."+"vipVariableName", data.Ipv6NatVariable.ValueString())
	} else if data.Ipv6Nat.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"nat64.enable."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"nat64.enable."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat64.enable."+"vipValue", strconv.FormatBool(data.Ipv6Nat.ValueBool()))
	}
	if !data.Nat64Interface.IsNull() {
		if data.Nat64Interface.ValueBool() {
			if !gjson.Get(body, path+"nat64").Exists() {
				body, _ = sjson.Set(body, path+"nat64", map[string]interface{}{})
			}
		} else {
			body, _ = sjson.Set(body, path+"nat64."+"vipObjectType", "node-only")
			body, _ = sjson.Set(body, path+"nat64."+"vipType", "ignore")
		}
	}
	if !data.Nat66Interface.IsNull() {
		if data.Nat66Interface.ValueBool() {
			if !gjson.Get(body, path+"nat66").Exists() {
				body, _ = sjson.Set(body, path+"nat66", map[string]interface{}{})
			}
		} else {
			body, _ = sjson.Set(body, path+"nat66."+"vipObjectType", "node-only")
			body, _ = sjson.Set(body, path+"nat66."+"vipType", "ignore")
		}
	}
	if len(data.StaticNat66Entries) > 0 {
		body, _ = sjson.Set(body, path+"nat66.static-nat66."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"nat66.static-nat66."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat66.static-nat66."+"vipPrimaryKey", []string{"source-prefix", "translated-source-prefix", "source-vpn-id"})
		body, _ = sjson.Set(body, path+"nat66.static-nat66."+"vipValue", []interface{}{})
	} else {
	}
	for _, item := range data.StaticNat66Entries {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "source-prefix")

		if !item.SourcePrefixVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-prefix."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-prefix."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-prefix."+"vipVariableName", item.SourcePrefixVariable.ValueString())
		} else if item.SourcePrefix.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-prefix."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-prefix."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-prefix."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-prefix."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-prefix."+"vipValue", item.SourcePrefix.ValueString())
		}
		itemAttributes = append(itemAttributes, "translated-source-prefix")

		if !item.TranslatedSourcePrefixVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "translated-source-prefix."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translated-source-prefix."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "translated-source-prefix."+"vipVariableName", item.TranslatedSourcePrefixVariable.ValueString())
		} else if item.TranslatedSourcePrefix.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "translated-source-prefix."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translated-source-prefix."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "translated-source-prefix."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translated-source-prefix."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "translated-source-prefix."+"vipValue", item.TranslatedSourcePrefix.ValueString())
		}
		itemAttributes = append(itemAttributes, "source-vpn-id")

		if !item.SourceVpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-vpn-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-vpn-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-vpn-id."+"vipVariableName", item.SourceVpnIdVariable.ValueString())
		} else if item.SourceVpnId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-vpn-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-vpn-id."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-vpn-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-vpn-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-vpn-id."+"vipValue", item.SourceVpnId.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"nat66.static-nat66."+"vipValue.-1", itemBody)
	}
	if len(data.StaticNatEntries) > 0 {
		body, _ = sjson.Set(body, path+"nat.static."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"nat.static."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.static."+"vipPrimaryKey", []string{"source-ip", "translate-ip"})
		body, _ = sjson.Set(body, path+"nat.static."+"vipValue", []interface{}{})
	} else {
	}
	for _, item := range data.StaticNatEntries {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "source-ip")

		if !item.SourceIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipVariableName", item.SourceIpVariable.ValueString())
		} else if item.SourceIp.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipValue", item.SourceIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "translate-ip")

		if !item.TranslateIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipVariableName", item.TranslateIpVariable.ValueString())
		} else if item.TranslateIp.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipValue", item.TranslateIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "static-nat-direction")

		if !item.StaticNatDirectionVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipVariableName", item.StaticNatDirectionVariable.ValueString())
		} else if item.StaticNatDirection.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipValue", item.StaticNatDirection.ValueString())
		}
		itemAttributes = append(itemAttributes, "source-vpn")

		if !item.SourceVpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipVariableName", item.SourceVpnIdVariable.ValueString())
		} else if item.SourceVpnId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipValue", item.SourceVpnId.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"nat.static."+"vipValue.-1", itemBody)
	}
	if len(data.StaticPortForwardEntries) > 0 {
		body, _ = sjson.Set(body, path+"nat.static-port-forward."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"nat.static-port-forward."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat.static-port-forward."+"vipPrimaryKey", []string{"source-ip", "translate-ip", "proto", "source-port", "translate-port"})
		body, _ = sjson.Set(body, path+"nat.static-port-forward."+"vipValue", []interface{}{})
	} else {
	}
	for _, item := range data.StaticPortForwardEntries {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "source-ip")

		if !item.SourceIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipVariableName", item.SourceIpVariable.ValueString())
		} else if item.SourceIp.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-ip."+"vipValue", item.SourceIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "translate-ip")

		if !item.TranslateIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipVariableName", item.TranslateIpVariable.ValueString())
		} else if item.TranslateIp.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "translate-ip."+"vipValue", item.TranslateIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "static-nat-direction")

		if !item.StaticNatDirectionVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipVariableName", item.StaticNatDirectionVariable.ValueString())
		} else if item.StaticNatDirection.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "static-nat-direction."+"vipValue", item.StaticNatDirection.ValueString())
		}
		itemAttributes = append(itemAttributes, "source-port")

		if !item.SourcePortVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-port."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-port."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-port."+"vipVariableName", item.SourcePortVariable.ValueString())
		} else if item.SourcePort.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-port."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-port."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-port."+"vipValue", item.SourcePort.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "translate-port")

		if !item.TranslatePortVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "translate-port."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translate-port."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "translate-port."+"vipVariableName", item.TranslatePortVariable.ValueString())
		} else if item.TranslatePort.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "translate-port."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "translate-port."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "translate-port."+"vipValue", item.TranslatePort.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "proto")

		if !item.ProtocolVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "proto."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "proto."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "proto."+"vipVariableName", item.ProtocolVariable.ValueString())
		} else if item.Protocol.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "proto."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "proto."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "proto."+"vipValue", item.Protocol.ValueString())
		}
		itemAttributes = append(itemAttributes, "source-vpn")

		if !item.SourceVpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipVariableName", item.SourceVpnIdVariable.ValueString())
		} else if item.SourceVpnId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-vpn."+"vipValue", item.SourceVpnId.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"nat.static-port-forward."+"vipValue.-1", itemBody)
	}
	if data.EnableCoreRegion.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.enable-core-region."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.enable-core-region."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.enable-core-region."+"vipValue", strconv.FormatBool(data.EnableCoreRegion.ValueBool()))
	}

	if !data.CoreRegionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.core-region."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.core-region."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.core-region."+"vipVariableName", data.CoreRegionVariable.ValueString())
	} else if data.CoreRegion.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.core-region."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.core-region."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.core-region."+"vipValue", data.CoreRegion.ValueString())
	}

	if !data.SecondaryRegionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.secondary-region."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.secondary-region."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.secondary-region."+"vipVariableName", data.SecondaryRegionVariable.ValueString())
	} else if data.SecondaryRegion.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.secondary-region."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.secondary-region."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.secondary-region."+"vipValue", data.SecondaryRegion.ValueString())
	}
	if len(data.TunnelInterfaceEncapsulations) > 0 {
		body, _ = sjson.Set(body, path+"tunnel-interface.encapsulation."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"tunnel-interface.encapsulation."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.encapsulation."+"vipPrimaryKey", []string{"encap"})
		body, _ = sjson.Set(body, path+"tunnel-interface.encapsulation."+"vipValue", []interface{}{})
	} else {
	}
	for _, item := range data.TunnelInterfaceEncapsulations {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "encap")
		if item.Encapsulation.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "encap."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "encap."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "encap."+"vipValue", item.Encapsulation.ValueString())
		}
		itemAttributes = append(itemAttributes, "preference")

		if !item.PreferenceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "preference."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "preference."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "preference."+"vipVariableName", item.PreferenceVariable.ValueString())
		} else if item.Preference.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "preference."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "preference."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "preference."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "preference."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "preference."+"vipValue", item.Preference.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "weight")

		if !item.WeightVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "weight."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "weight."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "weight."+"vipVariableName", item.WeightVariable.ValueString())
		} else if item.Weight.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "weight."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "weight."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "weight."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "weight."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "weight."+"vipValue", item.Weight.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"tunnel-interface.encapsulation."+"vipValue.-1", itemBody)
	}

	if !data.TunnelInterfaceBorderVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.border."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.border."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.border."+"vipVariableName", data.TunnelInterfaceBorderVariable.ValueString())
	} else if data.TunnelInterfaceBorder.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.border."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.border."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.border."+"vipValue", strconv.FormatBool(data.TunnelInterfaceBorder.ValueBool()))
	}

	if !data.TunnelQosModeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-qos.mode."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-qos.mode."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-qos.mode."+"vipVariableName", data.TunnelQosModeVariable.ValueString())
	} else if data.TunnelQosMode.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-qos.mode."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-qos.mode."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-qos.mode."+"vipValue", data.TunnelQosMode.ValueString())
	}

	if !data.TunnelBandwidthVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnels-bandwidth."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnels-bandwidth."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnels-bandwidth."+"vipVariableName", data.TunnelBandwidthVariable.ValueString())
	} else if data.TunnelBandwidth.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnels-bandwidth."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnels-bandwidth."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnels-bandwidth."+"vipValue", data.TunnelBandwidth.ValueInt64())
	}

	if !data.TunnelInterfaceGroupsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.group."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tunnel-interface.group."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.group."+"vipVariableName", data.TunnelInterfaceGroupsVariable.ValueString())
	} else if data.TunnelInterfaceGroups.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.group."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tunnel-interface.group."+"vipType", "constant")
		var values []int64
		data.TunnelInterfaceGroups.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"tunnel-interface.group."+"vipValue", values)
	}

	if !data.TunnelInterfaceColorVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.color.value."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.color.value."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.color.value."+"vipVariableName", data.TunnelInterfaceColorVariable.ValueString())
	} else if data.TunnelInterfaceColor.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.color.value."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.color.value."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.color.value."+"vipValue", data.TunnelInterfaceColor.ValueString())
	}

	if !data.TunnelInterfaceMaxControlConnectionsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.max-control-connections."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.max-control-connections."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.max-control-connections."+"vipVariableName", data.TunnelInterfaceMaxControlConnectionsVariable.ValueString())
	} else if data.TunnelInterfaceMaxControlConnections.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.max-control-connections."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.max-control-connections."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.max-control-connections."+"vipValue", data.TunnelInterfaceMaxControlConnections.ValueInt64())
	}

	if !data.TunnelInterfaceControlConnectionsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.control-connections."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.control-connections."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.control-connections."+"vipVariableName", data.TunnelInterfaceControlConnectionsVariable.ValueString())
	} else if data.TunnelInterfaceControlConnections.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.control-connections."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.control-connections."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.control-connections."+"vipValue", strconv.FormatBool(data.TunnelInterfaceControlConnections.ValueBool()))
	}

	if !data.TunnelInterfaceVbondAsStunServerVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.vbond-as-stun-server."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.vbond-as-stun-server."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.vbond-as-stun-server."+"vipVariableName", data.TunnelInterfaceVbondAsStunServerVariable.ValueString())
	} else if data.TunnelInterfaceVbondAsStunServer.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.vbond-as-stun-server."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.vbond-as-stun-server."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.vbond-as-stun-server."+"vipValue", strconv.FormatBool(data.TunnelInterfaceVbondAsStunServer.ValueBool()))
	}

	if !data.TunnelInterfaceExcludeControllerGroupListVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.exclude-controller-group-list."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tunnel-interface.exclude-controller-group-list."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.exclude-controller-group-list."+"vipVariableName", data.TunnelInterfaceExcludeControllerGroupListVariable.ValueString())
	} else if data.TunnelInterfaceExcludeControllerGroupList.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.exclude-controller-group-list."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"tunnel-interface.exclude-controller-group-list."+"vipType", "constant")
		var values []int64
		data.TunnelInterfaceExcludeControllerGroupList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"tunnel-interface.exclude-controller-group-list."+"vipValue", values)
	}

	if !data.TunnelInterfaceVmanageConnectionPreferenceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.vmanage-connection-preference."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.vmanage-connection-preference."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.vmanage-connection-preference."+"vipVariableName", data.TunnelInterfaceVmanageConnectionPreferenceVariable.ValueString())
	} else if data.TunnelInterfaceVmanageConnectionPreference.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.vmanage-connection-preference."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.vmanage-connection-preference."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.vmanage-connection-preference."+"vipValue", data.TunnelInterfaceVmanageConnectionPreference.ValueInt64())
	}

	if !data.TunnelInterfacePortHopVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.port-hop."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.port-hop."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.port-hop."+"vipVariableName", data.TunnelInterfacePortHopVariable.ValueString())
	} else if data.TunnelInterfacePortHop.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.port-hop."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.port-hop."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.port-hop."+"vipValue", strconv.FormatBool(data.TunnelInterfacePortHop.ValueBool()))
	}

	if !data.TunnelInterfaceColorRestrictVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.color.restrict."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"tunnel-interface.color.restrict."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.color.restrict."+"vipVariableName", data.TunnelInterfaceColorRestrictVariable.ValueString())
	} else if data.TunnelInterfaceColorRestrict.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.color.restrict."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"tunnel-interface.color.restrict."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.color.restrict."+"vipValue", strconv.FormatBool(data.TunnelInterfaceColorRestrict.ValueBool()))
	}

	if !data.TunnelInterfaceGreTunnelDestinationIpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.tloc-extension-gre-to.dst-ip."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.tloc-extension-gre-to.dst-ip."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.tloc-extension-gre-to.dst-ip."+"vipVariableName", data.TunnelInterfaceGreTunnelDestinationIpVariable.ValueString())
	} else if data.TunnelInterfaceGreTunnelDestinationIp.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.tloc-extension-gre-to.dst-ip."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.tloc-extension-gre-to.dst-ip."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.tloc-extension-gre-to.dst-ip."+"vipValue", data.TunnelInterfaceGreTunnelDestinationIp.ValueString())
	}

	if !data.TunnelInterfaceCarrierVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.carrier."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.carrier."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.carrier."+"vipVariableName", data.TunnelInterfaceCarrierVariable.ValueString())
	} else if data.TunnelInterfaceCarrier.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.carrier."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.carrier."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.carrier."+"vipValue", data.TunnelInterfaceCarrier.ValueString())
	}

	if !data.TunnelInterfaceNatRefreshIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.nat-refresh-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.nat-refresh-interval."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.nat-refresh-interval."+"vipVariableName", data.TunnelInterfaceNatRefreshIntervalVariable.ValueString())
	} else if data.TunnelInterfaceNatRefreshInterval.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.nat-refresh-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.nat-refresh-interval."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.nat-refresh-interval."+"vipValue", data.TunnelInterfaceNatRefreshInterval.ValueInt64())
	}

	if !data.TunnelInterfaceHelloIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-interval."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-interval."+"vipVariableName", data.TunnelInterfaceHelloIntervalVariable.ValueString())
	} else if data.TunnelInterfaceHelloInterval.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-interval."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-interval."+"vipValue", data.TunnelInterfaceHelloInterval.ValueInt64())
	}

	if !data.TunnelInterfaceHelloToleranceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-tolerance."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-tolerance."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-tolerance."+"vipVariableName", data.TunnelInterfaceHelloToleranceVariable.ValueString())
	} else if data.TunnelInterfaceHelloTolerance.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-tolerance."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-tolerance."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.hello-tolerance."+"vipValue", data.TunnelInterfaceHelloTolerance.ValueInt64())
	}

	if !data.TunnelInterfaceBindLoopbackTunnelVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.bind."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.bind."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.bind."+"vipVariableName", data.TunnelInterfaceBindLoopbackTunnelVariable.ValueString())
	} else if data.TunnelInterfaceBindLoopbackTunnel.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.bind."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.bind."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.bind."+"vipValue", data.TunnelInterfaceBindLoopbackTunnel.ValueString())
	}

	if !data.TunnelInterfaceLastResortCircuitVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.last-resort-circuit."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.last-resort-circuit."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.last-resort-circuit."+"vipVariableName", data.TunnelInterfaceLastResortCircuitVariable.ValueString())
	} else if data.TunnelInterfaceLastResortCircuit.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.last-resort-circuit."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.last-resort-circuit."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.last-resort-circuit."+"vipValue", strconv.FormatBool(data.TunnelInterfaceLastResortCircuit.ValueBool()))
	}

	if !data.TunnelInterfaceLowBandwidthLinkVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.low-bandwidth-link."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.low-bandwidth-link."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.low-bandwidth-link."+"vipVariableName", data.TunnelInterfaceLowBandwidthLinkVariable.ValueString())
	} else if data.TunnelInterfaceLowBandwidthLink.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.low-bandwidth-link."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.low-bandwidth-link."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.low-bandwidth-link."+"vipValue", strconv.FormatBool(data.TunnelInterfaceLowBandwidthLink.ValueBool()))
	}

	if !data.TunnelInterfaceTunnelTcpMssVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-tcp-mss-adjust."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-tcp-mss-adjust."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-tcp-mss-adjust."+"vipVariableName", data.TunnelInterfaceTunnelTcpMssVariable.ValueString())
	} else if data.TunnelInterfaceTunnelTcpMss.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-tcp-mss-adjust."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-tcp-mss-adjust."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.tunnel-tcp-mss-adjust."+"vipValue", data.TunnelInterfaceTunnelTcpMss.ValueInt64())
	}

	if !data.TunnelInterfaceClearDontFragmentVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.clear-dont-fragment."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.clear-dont-fragment."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.clear-dont-fragment."+"vipVariableName", data.TunnelInterfaceClearDontFragmentVariable.ValueString())
	} else if data.TunnelInterfaceClearDontFragment.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.clear-dont-fragment."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.clear-dont-fragment."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.clear-dont-fragment."+"vipValue", strconv.FormatBool(data.TunnelInterfaceClearDontFragment.ValueBool()))
	}

	if !data.TunnelInterfacePropagateSgtVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.propagate-sgt."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.propagate-sgt."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.propagate-sgt."+"vipVariableName", data.TunnelInterfacePropagateSgtVariable.ValueString())
	} else if data.TunnelInterfacePropagateSgt.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.propagate-sgt."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.propagate-sgt."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.propagate-sgt."+"vipValue", strconv.FormatBool(data.TunnelInterfacePropagateSgt.ValueBool()))
	}

	if !data.TunnelInterfaceNetworkBroadcastVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.network-broadcast."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.network-broadcast."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.network-broadcast."+"vipVariableName", data.TunnelInterfaceNetworkBroadcastVariable.ValueString())
	} else if data.TunnelInterfaceNetworkBroadcast.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.network-broadcast."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.network-broadcast."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.network-broadcast."+"vipValue", strconv.FormatBool(data.TunnelInterfaceNetworkBroadcast.ValueBool()))
	}

	if !data.TunnelInterfaceAllowAllVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.all."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.all."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.all."+"vipVariableName", data.TunnelInterfaceAllowAllVariable.ValueString())
	} else if data.TunnelInterfaceAllowAll.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.all."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.all."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.all."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowAll.ValueBool()))
	}

	if !data.TunnelInterfaceAllowBgpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.bgp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.bgp."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.bgp."+"vipVariableName", data.TunnelInterfaceAllowBgpVariable.ValueString())
	} else if data.TunnelInterfaceAllowBgp.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.bgp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.bgp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.bgp."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowBgp.ValueBool()))
	}

	if !data.TunnelInterfaceAllowDhcpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dhcp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dhcp."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dhcp."+"vipVariableName", data.TunnelInterfaceAllowDhcpVariable.ValueString())
	} else if data.TunnelInterfaceAllowDhcp.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dhcp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dhcp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dhcp."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowDhcp.ValueBool()))
	}

	if !data.TunnelInterfaceAllowDnsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dns."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dns."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dns."+"vipVariableName", data.TunnelInterfaceAllowDnsVariable.ValueString())
	} else if data.TunnelInterfaceAllowDns.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dns."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dns."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.dns."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowDns.ValueBool()))
	}

	if !data.TunnelInterfaceAllowIcmpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.icmp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.icmp."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.icmp."+"vipVariableName", data.TunnelInterfaceAllowIcmpVariable.ValueString())
	} else if data.TunnelInterfaceAllowIcmp.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.icmp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.icmp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.icmp."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowIcmp.ValueBool()))
	}

	if !data.TunnelInterfaceAllowSshVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.sshd."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.sshd."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.sshd."+"vipVariableName", data.TunnelInterfaceAllowSshVariable.ValueString())
	} else if data.TunnelInterfaceAllowSsh.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.sshd."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.sshd."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.sshd."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowSsh.ValueBool()))
	}

	if !data.TunnelInterfaceAllowNetconfVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.netconf."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.netconf."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.netconf."+"vipVariableName", data.TunnelInterfaceAllowNetconfVariable.ValueString())
	} else if data.TunnelInterfaceAllowNetconf.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.netconf."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.netconf."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.netconf."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowNetconf.ValueBool()))
	}

	if !data.TunnelInterfaceAllowNtpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ntp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ntp."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ntp."+"vipVariableName", data.TunnelInterfaceAllowNtpVariable.ValueString())
	} else if data.TunnelInterfaceAllowNtp.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ntp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ntp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ntp."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowNtp.ValueBool()))
	}

	if !data.TunnelInterfaceAllowOspfVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ospf."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ospf."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ospf."+"vipVariableName", data.TunnelInterfaceAllowOspfVariable.ValueString())
	} else if data.TunnelInterfaceAllowOspf.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ospf."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ospf."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.ospf."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowOspf.ValueBool()))
	}

	if !data.TunnelInterfaceAllowStunVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.stun."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.stun."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.stun."+"vipVariableName", data.TunnelInterfaceAllowStunVariable.ValueString())
	} else if data.TunnelInterfaceAllowStun.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.stun."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.stun."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.stun."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowStun.ValueBool()))
	}

	if !data.TunnelInterfaceAllowSnmpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.snmp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.snmp."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.snmp."+"vipVariableName", data.TunnelInterfaceAllowSnmpVariable.ValueString())
	} else if data.TunnelInterfaceAllowSnmp.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.snmp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.snmp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.snmp."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowSnmp.ValueBool()))
	}

	if !data.TunnelInterfaceAllowHttpsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.https."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.https."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.https."+"vipVariableName", data.TunnelInterfaceAllowHttpsVariable.ValueString())
	} else if data.TunnelInterfaceAllowHttps.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.https."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.https."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-interface.allow-service.https."+"vipValue", strconv.FormatBool(data.TunnelInterfaceAllowHttps.ValueBool()))
	}

	if !data.MediaTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"media-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"media-type."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"media-type."+"vipVariableName", data.MediaTypeVariable.ValueString())
	} else if data.MediaType.IsNull() {
		body, _ = sjson.Set(body, path+"media-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"media-type."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"media-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"media-type."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"media-type."+"vipValue", data.MediaType.ValueString())
	}

	if !data.InterfaceMtuVariable.IsNull() {
		body, _ = sjson.Set(body, path+"intrf-mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"intrf-mtu."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"intrf-mtu."+"vipVariableName", data.InterfaceMtuVariable.ValueString())
	} else if data.InterfaceMtu.IsNull() {
		body, _ = sjson.Set(body, path+"intrf-mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"intrf-mtu."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"intrf-mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"intrf-mtu."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"intrf-mtu."+"vipValue", data.InterfaceMtu.ValueInt64())
	}

	if !data.IpMtuVariable.IsNull() {
		body, _ = sjson.Set(body, path+"mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mtu."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"mtu."+"vipVariableName", data.IpMtuVariable.ValueString())
	} else if data.IpMtu.IsNull() {
		body, _ = sjson.Set(body, path+"mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mtu."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mtu."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"mtu."+"vipValue", data.IpMtu.ValueInt64())
	}

	if !data.TcpMssAdjustVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipVariableName", data.TcpMssAdjustVariable.ValueString())
	} else if data.TcpMssAdjust.IsNull() {
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tcp-mss-adjust."+"vipValue", data.TcpMssAdjust.ValueInt64())
	}

	if !data.TlocExtensionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tloc-extension."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tloc-extension."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tloc-extension."+"vipVariableName", data.TlocExtensionVariable.ValueString())
	} else if data.TlocExtension.IsNull() {
		body, _ = sjson.Set(body, path+"tloc-extension."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tloc-extension."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"tloc-extension."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tloc-extension."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tloc-extension."+"vipValue", data.TlocExtension.ValueString())
	}

	if !data.LoadIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"load-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"load-interval."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"load-interval."+"vipVariableName", data.LoadIntervalVariable.ValueString())
	} else if data.LoadInterval.IsNull() {
		body, _ = sjson.Set(body, path+"load-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"load-interval."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"load-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"load-interval."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"load-interval."+"vipValue", data.LoadInterval.ValueInt64())
	}

	if !data.GreTunnelSourceIpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tloc-extension-gre-from.src-ip."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tloc-extension-gre-from.src-ip."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tloc-extension-gre-from.src-ip."+"vipVariableName", data.GreTunnelSourceIpVariable.ValueString())
	} else if data.GreTunnelSourceIp.IsNull() {
		body, _ = sjson.Set(body, path+"tloc-extension-gre-from.src-ip."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tloc-extension-gre-from.src-ip."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"tloc-extension-gre-from.src-ip."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tloc-extension-gre-from.src-ip."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tloc-extension-gre-from.src-ip."+"vipValue", data.GreTunnelSourceIp.ValueString())
	}

	if !data.GreTunnelXconnectVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tloc-extension-gre-from.xconnect."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tloc-extension-gre-from.xconnect."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tloc-extension-gre-from.xconnect."+"vipVariableName", data.GreTunnelXconnectVariable.ValueString())
	} else if data.GreTunnelXconnect.IsNull() {
		body, _ = sjson.Set(body, path+"tloc-extension-gre-from.xconnect."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tloc-extension-gre-from.xconnect."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"tloc-extension-gre-from.xconnect."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tloc-extension-gre-from.xconnect."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tloc-extension-gre-from.xconnect."+"vipValue", data.GreTunnelXconnect.ValueString())
	}

	if !data.MacAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"mac-address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mac-address."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"mac-address."+"vipVariableName", data.MacAddressVariable.ValueString())
	} else if data.MacAddress.IsNull() {
		body, _ = sjson.Set(body, path+"mac-address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mac-address."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"mac-address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mac-address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"mac-address."+"vipValue", data.MacAddress.ValueString())
	}

	if !data.SpeedVariable.IsNull() {
		body, _ = sjson.Set(body, path+"speed."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"speed."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"speed."+"vipVariableName", data.SpeedVariable.ValueString())
	} else if data.Speed.IsNull() {
		body, _ = sjson.Set(body, path+"speed."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"speed."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"speed."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"speed."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"speed."+"vipValue", data.Speed.ValueString())
	}

	if !data.DuplexVariable.IsNull() {
		body, _ = sjson.Set(body, path+"duplex."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"duplex."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"duplex."+"vipVariableName", data.DuplexVariable.ValueString())
	} else if data.Duplex.IsNull() {
		body, _ = sjson.Set(body, path+"duplex."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"duplex."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"duplex."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"duplex."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"duplex."+"vipValue", data.Duplex.ValueString())
	}

	if !data.ShutdownVariable.IsNull() {
		body, _ = sjson.Set(body, path+"shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"shutdown."+"vipVariableName", data.ShutdownVariable.ValueString())
	} else if data.Shutdown.IsNull() {
		body, _ = sjson.Set(body, path+"shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"shutdown."+"vipValue", strconv.FormatBool(data.Shutdown.ValueBool()))
	}

	if !data.ArpTimeoutVariable.IsNull() {
		body, _ = sjson.Set(body, path+"arp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"arp-timeout."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"arp-timeout."+"vipVariableName", data.ArpTimeoutVariable.ValueString())
	} else if data.ArpTimeout.IsNull() {
		body, _ = sjson.Set(body, path+"arp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"arp-timeout."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"arp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"arp-timeout."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"arp-timeout."+"vipValue", data.ArpTimeout.ValueInt64())
	}

	if !data.AutonegotiateVariable.IsNull() {
		body, _ = sjson.Set(body, path+"autonegotiate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"autonegotiate."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"autonegotiate."+"vipVariableName", data.AutonegotiateVariable.ValueString())
	} else if data.Autonegotiate.IsNull() {
		body, _ = sjson.Set(body, path+"autonegotiate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"autonegotiate."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"autonegotiate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"autonegotiate."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"autonegotiate."+"vipValue", strconv.FormatBool(data.Autonegotiate.ValueBool()))
	}

	if !data.IpDirectedBroadcastVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipVariableName", data.IpDirectedBroadcastVariable.ValueString())
	} else if data.IpDirectedBroadcast.IsNull() {
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip-directed-broadcast."+"vipValue", strconv.FormatBool(data.IpDirectedBroadcast.ValueBool()))
	}

	if !data.IcmpRedirectDisableVariable.IsNull() {
		body, _ = sjson.Set(body, path+"icmp-redirect-disable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"icmp-redirect-disable."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"icmp-redirect-disable."+"vipVariableName", data.IcmpRedirectDisableVariable.ValueString())
	} else if data.IcmpRedirectDisable.IsNull() {
		body, _ = sjson.Set(body, path+"icmp-redirect-disable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"icmp-redirect-disable."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"icmp-redirect-disable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"icmp-redirect-disable."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"icmp-redirect-disable."+"vipValue", strconv.FormatBool(data.IcmpRedirectDisable.ValueBool()))
	}

	if !data.QosAdaptivePeriodVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-adaptive.period."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.period."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-adaptive.period."+"vipVariableName", data.QosAdaptivePeriodVariable.ValueString())
	} else if data.QosAdaptivePeriod.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"qos-adaptive.period."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.period."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-adaptive.period."+"vipValue", data.QosAdaptivePeriod.ValueInt64())
	}

	if !data.QosAdaptiveBandwidthDownstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.bandwidth-down."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.bandwidth-down."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.bandwidth-down."+"vipVariableName", data.QosAdaptiveBandwidthDownstreamVariable.ValueString())
	} else if data.QosAdaptiveBandwidthDownstream.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.bandwidth-down."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.bandwidth-down."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.bandwidth-down."+"vipValue", data.QosAdaptiveBandwidthDownstream.ValueInt64())
	}

	if !data.QosAdaptiveMinDownstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmin."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmin."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmin."+"vipVariableName", data.QosAdaptiveMinDownstreamVariable.ValueString())
	} else if data.QosAdaptiveMinDownstream.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmin."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmin."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmin."+"vipValue", data.QosAdaptiveMinDownstream.ValueInt64())
	}

	if !data.QosAdaptiveMaxDownstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmax."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmax."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmax."+"vipVariableName", data.QosAdaptiveMaxDownstreamVariable.ValueString())
	} else if data.QosAdaptiveMaxDownstream.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmax."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmax."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-adaptive.downstream.range.dmax."+"vipValue", data.QosAdaptiveMaxDownstream.ValueInt64())
	}

	if !data.QosAdaptiveBandwidthUpstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.bandwidth-up."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.bandwidth-up."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.bandwidth-up."+"vipVariableName", data.QosAdaptiveBandwidthUpstreamVariable.ValueString())
	} else if data.QosAdaptiveBandwidthUpstream.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.bandwidth-up."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.bandwidth-up."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.bandwidth-up."+"vipValue", data.QosAdaptiveBandwidthUpstream.ValueInt64())
	}

	if !data.QosAdaptiveMinUpstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umin."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umin."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umin."+"vipVariableName", data.QosAdaptiveMinUpstreamVariable.ValueString())
	} else if data.QosAdaptiveMinUpstream.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umin."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umin."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umin."+"vipValue", data.QosAdaptiveMinUpstream.ValueInt64())
	}

	if !data.QosAdaptiveMaxUpstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umax."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umax."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umax."+"vipVariableName", data.QosAdaptiveMaxUpstreamVariable.ValueString())
	} else if data.QosAdaptiveMaxUpstream.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umax."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umax."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-adaptive.upstream.range.umax."+"vipValue", data.QosAdaptiveMaxUpstream.ValueInt64())
	}

	if !data.ShapingRateVariable.IsNull() {
		body, _ = sjson.Set(body, path+"shaping-rate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shaping-rate."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"shaping-rate."+"vipVariableName", data.ShapingRateVariable.ValueString())
	} else if data.ShapingRate.IsNull() {
		body, _ = sjson.Set(body, path+"shaping-rate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shaping-rate."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"shaping-rate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shaping-rate."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"shaping-rate."+"vipValue", data.ShapingRate.ValueInt64())
	}

	if !data.QosMapVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-map."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-map."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-map."+"vipVariableName", data.QosMapVariable.ValueString())
	} else if data.QosMap.IsNull() {
		body, _ = sjson.Set(body, path+"qos-map."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-map."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"qos-map."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-map."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-map."+"vipValue", data.QosMap.ValueString())
	}

	if !data.QosMapVpnVariable.IsNull() {
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipVariableName", data.QosMapVpnVariable.ValueString())
	} else if data.QosMapVpn.IsNull() {
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"qos-map-vpn."+"vipValue", data.QosMapVpn.ValueString())
	}

	if !data.BandwidthUpstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bandwidth-upstream."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bandwidth-upstream."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bandwidth-upstream."+"vipVariableName", data.BandwidthUpstreamVariable.ValueString())
	} else if data.BandwidthUpstream.IsNull() {
		body, _ = sjson.Set(body, path+"bandwidth-upstream."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bandwidth-upstream."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bandwidth-upstream."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bandwidth-upstream."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bandwidth-upstream."+"vipValue", data.BandwidthUpstream.ValueInt64())
	}

	if !data.BandwidthDownstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bandwidth-downstream."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bandwidth-downstream."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bandwidth-downstream."+"vipVariableName", data.BandwidthDownstreamVariable.ValueString())
	} else if data.BandwidthDownstream.IsNull() {
		body, _ = sjson.Set(body, path+"bandwidth-downstream."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bandwidth-downstream."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bandwidth-downstream."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bandwidth-downstream."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bandwidth-downstream."+"vipValue", data.BandwidthDownstream.ValueInt64())
	}

	if !data.BlockNonSourceIpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"block-non-source-ip."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"block-non-source-ip."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"block-non-source-ip."+"vipVariableName", data.BlockNonSourceIpVariable.ValueString())
	} else if data.BlockNonSourceIp.IsNull() {
		body, _ = sjson.Set(body, path+"block-non-source-ip."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"block-non-source-ip."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"block-non-source-ip."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"block-non-source-ip."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"block-non-source-ip."+"vipValue", strconv.FormatBool(data.BlockNonSourceIp.ValueBool()))
	}

	if !data.RewriteRuleNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipVariableName", data.RewriteRuleNameVariable.ValueString())
	} else if data.RewriteRuleName.IsNull() {
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"rewrite-rule.rule-name."+"vipValue", data.RewriteRuleName.ValueString())
	}
	if len(data.AccessLists) > 0 {
		body, _ = sjson.Set(body, path+"access-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"access-list."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"access-list."+"vipPrimaryKey", []string{"direction"})
		body, _ = sjson.Set(body, path+"access-list."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"access-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"access-list."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"access-list."+"vipPrimaryKey", []string{"direction"})
		body, _ = sjson.Set(body, path+"access-list."+"vipValue", []interface{}{})
	}
	for _, item := range data.AccessLists {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "direction")
		if item.Direction.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipValue", item.Direction.ValueString())
		}
		itemAttributes = append(itemAttributes, "acl-name")

		if !item.AclNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipVariableName", item.AclNameVariable.ValueString())
		} else if item.AclName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "acl-name."+"vipValue", item.AclName.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"access-list."+"vipValue.-1", itemBody)
	}
	if len(data.StaticArps) > 0 {
		body, _ = sjson.Set(body, path+"arp.ip."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"arp.ip."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"arp.ip."+"vipPrimaryKey", []string{"addr"})
		body, _ = sjson.Set(body, path+"arp.ip."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"arp.ip."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"arp.ip."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"arp.ip."+"vipPrimaryKey", []string{"addr"})
		body, _ = sjson.Set(body, path+"arp.ip."+"vipValue", []interface{}{})
	}
	for _, item := range data.StaticArps {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "addr")

		if !item.IpAddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "addr."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "addr."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "addr."+"vipVariableName", item.IpAddressVariable.ValueString())
		} else if item.IpAddress.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "addr."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "addr."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "addr."+"vipValue", item.IpAddress.ValueString())
		}
		itemAttributes = append(itemAttributes, "mac")

		if !item.MacVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "mac."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "mac."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "mac."+"vipVariableName", item.MacVariable.ValueString())
		} else if item.Mac.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "mac."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "mac."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "mac."+"vipValue", item.Mac.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"arp.ip."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv4Vrrps) > 0 {
		body, _ = sjson.Set(body, path+"vrrp."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"vrrp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"vrrp."+"vipPrimaryKey", []string{"grp-id"})
		body, _ = sjson.Set(body, path+"vrrp."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"vrrp."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"vrrp."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"vrrp."+"vipPrimaryKey", []string{"grp-id"})
		body, _ = sjson.Set(body, path+"vrrp."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv4Vrrps {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "grp-id")

		if !item.GroupIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipVariableName", item.GroupIdVariable.ValueString())
		} else if item.GroupId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipValue", item.GroupId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "priority")

		if !item.PriorityVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipVariableName", item.PriorityVariable.ValueString())
		} else if item.Priority.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipValue", item.Priority.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "timer")

		if !item.TimerVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipVariableName", item.TimerVariable.ValueString())
		} else if item.Timer.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipValue", item.Timer.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "track-omp")
		if item.TrackOmp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipValue", strconv.FormatBool(item.TrackOmp.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "track-prefix-list")

		if !item.TrackPrefixListVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipVariableName", item.TrackPrefixListVariable.ValueString())
		} else if item.TrackPrefixList.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipValue", item.TrackPrefixList.ValueString())
		}
		itemAttributes = append(itemAttributes, "ipv4")

		if !item.IpAddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ipv4.address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ipv4.address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ipv4.address."+"vipVariableName", item.IpAddressVariable.ValueString())
		} else if item.IpAddress.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ipv4.address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ipv4.address."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "ipv4.address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ipv4.address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ipv4.address."+"vipValue", item.IpAddress.ValueString())
		}
		itemAttributes = append(itemAttributes, "ipv4-secondary")
		if len(item.Ipv4SecondaryAddresses) > 0 {
			itemBody, _ = sjson.Set(itemBody, "ipv4-secondary."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "ipv4-secondary."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ipv4-secondary."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "ipv4-secondary."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "ipv4-secondary."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "ipv4-secondary."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "ipv4-secondary."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "ipv4-secondary."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Ipv4SecondaryAddresses {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "address")

			if !childItem.IpAddressVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipVariableName", childItem.IpAddressVariable.ValueString())
			} else if childItem.IpAddress.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipValue", childItem.IpAddress.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "ipv4-secondary."+"vipValue.-1", itemChildBody)
		}
		itemAttributes = append(itemAttributes, "tloc-change-pref")
		if item.TlocPreferenceChange.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tloc-change-pref."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tloc-change-pref."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "tloc-change-pref."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "tloc-change-pref."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tloc-change-pref."+"vipValue", strconv.FormatBool(item.TlocPreferenceChange.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "value")

		if !item.TlocPreferenceChangeValueVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "value."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "value."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "value."+"vipVariableName", item.TlocPreferenceChangeValueVariable.ValueString())
		} else if item.TlocPreferenceChangeValue.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "value."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "value."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "value."+"vipValue", item.TlocPreferenceChangeValue.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "tracking-object")
		if len(item.TrackingObjects) > 0 {
			itemBody, _ = sjson.Set(itemBody, "tracking-object."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "tracking-object."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "tracking-object."+"vipPrimaryKey", []string{"name"})
			itemBody, _ = sjson.Set(itemBody, "tracking-object."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "tracking-object."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "tracking-object."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "tracking-object."+"vipPrimaryKey", []string{"name"})
			itemBody, _ = sjson.Set(itemBody, "tracking-object."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.TrackingObjects {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "name")

			if !childItem.TrackerIdVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipVariableName", childItem.TrackerIdVariable.ValueString())
			} else if childItem.TrackerId.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipValue", childItem.TrackerId.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "track-action")

			if !childItem.TrackActionVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "track-action."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "track-action."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "track-action."+"vipVariableName", childItem.TrackActionVariable.ValueString())
			} else if childItem.TrackAction.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "track-action."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "track-action."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "track-action."+"vipValue", childItem.TrackAction.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "decrement")

			if !childItem.DecrementValueVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "decrement."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "decrement."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "decrement."+"vipVariableName", childItem.DecrementValueVariable.ValueString())
			} else if childItem.DecrementValue.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "decrement."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "decrement."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "decrement."+"vipValue", childItem.DecrementValue.ValueInt64())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "tracking-object."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"vrrp."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv6Vrrps) > 0 {
		body, _ = sjson.Set(body, path+"ipv6-vrrp."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6-vrrp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipv6-vrrp."+"vipPrimaryKey", []string{"grp-id"})
		body, _ = sjson.Set(body, path+"ipv6-vrrp."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ipv6-vrrp."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ipv6-vrrp."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ipv6-vrrp."+"vipPrimaryKey", []string{"grp-id"})
		body, _ = sjson.Set(body, path+"ipv6-vrrp."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv6Vrrps {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "grp-id")

		if !item.GroupIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipVariableName", item.GroupIdVariable.ValueString())
		} else if item.GroupId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "grp-id."+"vipValue", item.GroupId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "priority")

		if !item.PriorityVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipVariableName", item.PriorityVariable.ValueString())
		} else if item.Priority.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipValue", item.Priority.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "timer")

		if !item.TimerVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipVariableName", item.TimerVariable.ValueString())
		} else if item.Timer.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "timer."+"vipValue", item.Timer.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "track-omp")

		if !item.TrackOmpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipVariableName", item.TrackOmpVariable.ValueString())
		} else if item.TrackOmp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "track-omp."+"vipValue", strconv.FormatBool(item.TrackOmp.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "track-prefix-list")

		if !item.TrackPrefixListVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipVariableName", item.TrackPrefixListVariable.ValueString())
		} else if item.TrackPrefixList.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "track-prefix-list."+"vipValue", item.TrackPrefixList.ValueString())
		}
		itemAttributes = append(itemAttributes, "ipv6")
		if len(item.Ipv6Addresses) > 0 {
			itemBody, _ = sjson.Set(itemBody, "ipv6."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "ipv6."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ipv6."+"vipPrimaryKey", []string{"ipv6-link-local"})
			itemBody, _ = sjson.Set(itemBody, "ipv6."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "ipv6."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "ipv6."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "ipv6."+"vipPrimaryKey", []string{"ipv6-link-local"})
			itemBody, _ = sjson.Set(itemBody, "ipv6."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Ipv6Addresses {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "ipv6-link-local")

			if !childItem.Ipv6LinkLocalVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv6-link-local."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv6-link-local."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv6-link-local."+"vipVariableName", childItem.Ipv6LinkLocalVariable.ValueString())
			} else if childItem.Ipv6LinkLocal.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv6-link-local."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv6-link-local."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv6-link-local."+"vipValue", childItem.Ipv6LinkLocal.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "prefix")

			if !childItem.PrefixVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipVariableName", childItem.PrefixVariable.ValueString())
			} else if childItem.Prefix.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipValue", childItem.Prefix.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "ipv6."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ipv6-vrrp."+"vipValue.-1", itemBody)
	}
	if data.PropagateSgt.IsNull() {
		body, _ = sjson.Set(body, path+"trustsec.propagate.sgt."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"trustsec.propagate.sgt."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"trustsec.propagate.sgt."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"trustsec.propagate.sgt."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"trustsec.propagate.sgt."+"vipValue", strconv.FormatBool(data.PropagateSgt.ValueBool()))
	}

	if !data.StaticSgtVariable.IsNull() {
		body, _ = sjson.Set(body, path+"trustsec.static.sgt."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"trustsec.static.sgt."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"trustsec.static.sgt."+"vipVariableName", data.StaticSgtVariable.ValueString())
	} else if data.StaticSgt.IsNull() {
		body, _ = sjson.Set(body, path+"trustsec.static.sgt."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"trustsec.static.sgt."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"trustsec.static.sgt."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"trustsec.static.sgt."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"trustsec.static.sgt."+"vipValue", data.StaticSgt.ValueInt64())
	}
	if data.StaticSgtTrusted.IsNull() {
		body, _ = sjson.Set(body, path+"trustsec.static.trusted."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"trustsec.static.trusted."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"trustsec.static.trusted."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"trustsec.static.trusted."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"trustsec.static.trusted."+"vipValue", strconv.FormatBool(data.StaticSgtTrusted.ValueBool()))
	}
	if data.EnableSgt.IsNull() {
		body, _ = sjson.Set(body, path+"trustsec.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"trustsec.enable."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"trustsec.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"trustsec.enable."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"trustsec.enable."+"vipValue", strconv.FormatBool(data.EnableSgt.ValueBool()))
	}
	if data.SgtEnforcement.IsNull() {
		body, _ = sjson.Set(body, path+"trustsec.enforcement.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"trustsec.enforcement.enable."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"trustsec.enforcement.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"trustsec.enforcement.enable."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"trustsec.enforcement.enable."+"vipValue", strconv.FormatBool(data.SgtEnforcement.ValueBool()))
	}

	if !data.SgtEnforcementSgtVariable.IsNull() {
		body, _ = sjson.Set(body, path+"trustsec.enforcement.sgt."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"trustsec.enforcement.sgt."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"trustsec.enforcement.sgt."+"vipVariableName", data.SgtEnforcementSgtVariable.ValueString())
	} else if data.SgtEnforcementSgt.IsNull() {
		body, _ = sjson.Set(body, path+"trustsec.enforcement.sgt."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"trustsec.enforcement.sgt."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"trustsec.enforcement.sgt."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"trustsec.enforcement.sgt."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"trustsec.enforcement.sgt."+"vipValue", data.SgtEnforcementSgt.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoVPNInterface) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceType"); value.Exists() {
		data.DeviceTypes = helpers.GetStringSet(value.Array())
	} else {
		data.DeviceTypes = types.SetNull(types.StringType)
	}
	if value := res.Get("templateDescription"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("templateName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("templateType"); value.Exists() {
		data.TemplateType = types.StringValue(value.String())
	} else {
		data.TemplateType = types.StringNull()
	}

	path := "templateDefinition."
	if value := res.Get(path + "if-name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.InterfaceName = types.StringNull()

			v := res.Get(path + "if-name.vipVariableName")
			data.InterfaceNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.InterfaceName = types.StringNull()
			data.InterfaceNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "if-name.vipValue")
			data.InterfaceName = types.StringValue(v.String())
			data.InterfaceNameVariable = types.StringNull()
		}
	} else {
		data.InterfaceName = types.StringNull()
		data.InterfaceNameVariable = types.StringNull()
	}
	if value := res.Get(path + "description.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.InterfaceDescription = types.StringNull()

			v := res.Get(path + "description.vipVariableName")
			data.InterfaceDescriptionVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.InterfaceDescription = types.StringNull()
			data.InterfaceDescriptionVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "description.vipValue")
			data.InterfaceDescription = types.StringValue(v.String())
			data.InterfaceDescriptionVariable = types.StringNull()
		}
	} else {
		data.InterfaceDescription = types.StringNull()
		data.InterfaceDescriptionVariable = types.StringNull()
	}
	if value := res.Get(path + "poe.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Poe = types.BoolNull()

			v := res.Get(path + "poe.vipVariableName")
			data.PoeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Poe = types.BoolNull()
			data.PoeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "poe.vipValue")
			data.Poe = types.BoolValue(v.Bool())
			data.PoeVariable = types.StringNull()
		}
	} else {
		data.Poe = types.BoolNull()
		data.PoeVariable = types.StringNull()
	}
	if value := res.Get(path + "ip.address.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Address = types.StringNull()

			v := res.Get(path + "ip.address.vipVariableName")
			data.AddressVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Address = types.StringNull()
			data.AddressVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ip.address.vipValue")
			data.Address = types.StringValue(v.String())
			data.AddressVariable = types.StringNull()
		}
	} else {
		data.Address = types.StringNull()
		data.AddressVariable = types.StringNull()
	}
	if value := res.Get(path + "ip.secondary-address.vipValue"); len(value.Array()) > 0 {
		data.Ipv4SecondaryAddresses = make([]CiscoVPNInterfaceIpv4SecondaryAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNInterfaceIpv4SecondaryAddresses{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("address.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Address = types.StringNull()

					cv := v.Get("address.vipVariableName")
					item.AddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Address = types.StringNull()
					item.AddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("address.vipValue")
					item.Address = types.StringValue(cv.String())
					item.AddressVariable = types.StringNull()
				}
			} else {
				item.Address = types.StringNull()
				item.AddressVariable = types.StringNull()
			}
			data.Ipv4SecondaryAddresses = append(data.Ipv4SecondaryAddresses, item)
			return true
		})
	} else {
		if len(data.Ipv4SecondaryAddresses) > 0 {
			data.Ipv4SecondaryAddresses = []CiscoVPNInterfaceIpv4SecondaryAddresses{}
		}
	}
	if value := res.Get(path + "ip.dhcp-client.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Dhcp = types.BoolNull()

			v := res.Get(path + "ip.dhcp-client.vipVariableName")
			data.DhcpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Dhcp = types.BoolNull()
			data.DhcpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ip.dhcp-client.vipValue")
			data.Dhcp = types.BoolValue(v.Bool())
			data.DhcpVariable = types.StringNull()
		}
	} else {
		data.Dhcp = types.BoolNull()
		data.DhcpVariable = types.StringNull()
	}
	if value := res.Get(path + "ip.dhcp-distance.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DhcpDistance = types.Int64Null()

			v := res.Get(path + "ip.dhcp-distance.vipVariableName")
			data.DhcpDistanceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DhcpDistance = types.Int64Null()
			data.DhcpDistanceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ip.dhcp-distance.vipValue")
			data.DhcpDistance = types.Int64Value(v.Int())
			data.DhcpDistanceVariable = types.StringNull()
		}
	} else {
		data.DhcpDistance = types.Int64Null()
		data.DhcpDistanceVariable = types.StringNull()
	}
	if value := res.Get(path + "ipv6.address.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6Address = types.StringNull()

			v := res.Get(path + "ipv6.address.vipVariableName")
			data.Ipv6AddressVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6Address = types.StringNull()
			data.Ipv6AddressVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipv6.address.vipValue")
			data.Ipv6Address = types.StringValue(v.String())
			data.Ipv6AddressVariable = types.StringNull()
		}
	} else {
		data.Ipv6Address = types.StringNull()
		data.Ipv6AddressVariable = types.StringNull()
	}
	if value := res.Get(path + "ipv6.dhcp-client.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Dhcpv6 = types.BoolNull()

			v := res.Get(path + "ipv6.dhcp-client.vipVariableName")
			data.Dhcpv6Variable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Dhcpv6 = types.BoolNull()
			data.Dhcpv6Variable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipv6.dhcp-client.vipValue")
			data.Dhcpv6 = types.BoolValue(v.Bool())
			data.Dhcpv6Variable = types.StringNull()
		}
	} else {
		data.Dhcpv6 = types.BoolNull()
		data.Dhcpv6Variable = types.StringNull()
	}
	if value := res.Get(path + "ipv6.secondary-address.vipValue"); len(value.Array()) > 0 {
		data.Ipv6SecondaryAddresses = make([]CiscoVPNInterfaceIpv6SecondaryAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNInterfaceIpv6SecondaryAddresses{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("address.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Address = types.StringNull()

					cv := v.Get("address.vipVariableName")
					item.AddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Address = types.StringNull()
					item.AddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("address.vipValue")
					item.Address = types.StringValue(cv.String())
					item.AddressVariable = types.StringNull()
				}
			} else {
				item.Address = types.StringNull()
				item.AddressVariable = types.StringNull()
			}
			data.Ipv6SecondaryAddresses = append(data.Ipv6SecondaryAddresses, item)
			return true
		})
	} else {
		if len(data.Ipv6SecondaryAddresses) > 0 {
			data.Ipv6SecondaryAddresses = []CiscoVPNInterfaceIpv6SecondaryAddresses{}
		}
	}
	if value := res.Get(path + "ipv6.access-list.vipValue"); len(value.Array()) > 0 {
		data.Ipv6AccessLists = make([]CiscoVPNInterfaceIpv6AccessLists, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNInterfaceIpv6AccessLists{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("direction.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Direction = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Direction = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("direction.vipValue")
					item.Direction = types.StringValue(cv.String())

				}
			} else {
				item.Direction = types.StringNull()

			}
			if cValue := v.Get("acl-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AclName = types.StringNull()

					cv := v.Get("acl-name.vipVariableName")
					item.AclNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AclName = types.StringNull()
					item.AclNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("acl-name.vipValue")
					item.AclName = types.StringValue(cv.String())
					item.AclNameVariable = types.StringNull()
				}
			} else {
				item.AclName = types.StringNull()
				item.AclNameVariable = types.StringNull()
			}
			data.Ipv6AccessLists = append(data.Ipv6AccessLists, item)
			return true
		})
	} else {
		if len(data.Ipv6AccessLists) > 0 {
			data.Ipv6AccessLists = []CiscoVPNInterfaceIpv6AccessLists{}
		}
	}
	if value := res.Get(path + "dhcp-helper.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.Ipv4DhcpHelper = types.SetNull(types.StringType)

			v := res.Get(path + "dhcp-helper.vipVariableName")
			data.Ipv4DhcpHelperVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv4DhcpHelper = types.SetNull(types.StringType)
			data.Ipv4DhcpHelperVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "dhcp-helper.vipValue")
			data.Ipv4DhcpHelper = helpers.GetStringSet(v.Array())
			data.Ipv4DhcpHelperVariable = types.StringNull()
		}
	} else {
		data.Ipv4DhcpHelper = types.SetNull(types.StringType)
		data.Ipv4DhcpHelperVariable = types.StringNull()
	}
	if value := res.Get(path + "ipv6.dhcp-helper-v6.vipValue"); len(value.Array()) > 0 {
		data.Ipv6DhcpHelpers = make([]CiscoVPNInterfaceIpv6DhcpHelpers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNInterfaceIpv6DhcpHelpers{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("address.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Address = types.StringNull()

					cv := v.Get("address.vipVariableName")
					item.AddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Address = types.StringNull()
					item.AddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("address.vipValue")
					item.Address = types.StringValue(cv.String())
					item.AddressVariable = types.StringNull()
				}
			} else {
				item.Address = types.StringNull()
				item.AddressVariable = types.StringNull()
			}
			if cValue := v.Get("vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

					cv := v.Get("vpn.vipVariableName")
					item.VpnIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()
					item.VpnIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("vpn.vipValue")
					item.VpnId = types.Int64Value(cv.Int())
					item.VpnIdVariable = types.StringNull()
				}
			} else {
				item.VpnId = types.Int64Null()
				item.VpnIdVariable = types.StringNull()
			}
			data.Ipv6DhcpHelpers = append(data.Ipv6DhcpHelpers, item)
			return true
		})
	} else {
		if len(data.Ipv6DhcpHelpers) > 0 {
			data.Ipv6DhcpHelpers = []CiscoVPNInterfaceIpv6DhcpHelpers{}
		}
	}
	if value := res.Get(path + "tracker.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.Tracker = types.SetNull(types.StringType)

			v := res.Get(path + "tracker.vipVariableName")
			data.TrackerVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Tracker = types.SetNull(types.StringType)
			data.TrackerVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tracker.vipValue")
			data.Tracker = helpers.GetStringSet(v.Array())
			data.TrackerVariable = types.StringNull()
		}
	} else {
		data.Tracker = types.SetNull(types.StringType)
		data.TrackerVariable = types.StringNull()
	}
	if value := res.Get(path + "auto-bandwidth-detect.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AutoBandwidthDetect = types.BoolNull()

			v := res.Get(path + "auto-bandwidth-detect.vipVariableName")
			data.AutoBandwidthDetectVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AutoBandwidthDetect = types.BoolNull()
			data.AutoBandwidthDetectVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "auto-bandwidth-detect.vipValue")
			data.AutoBandwidthDetect = types.BoolValue(v.Bool())
			data.AutoBandwidthDetectVariable = types.StringNull()
		}
	} else {
		data.AutoBandwidthDetect = types.BoolNull()
		data.AutoBandwidthDetectVariable = types.StringNull()
	}
	if value := res.Get(path + "iperf-server.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IperfServer = types.StringNull()

			v := res.Get(path + "iperf-server.vipVariableName")
			data.IperfServerVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IperfServer = types.StringNull()
			data.IperfServerVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "iperf-server.vipValue")
			data.IperfServer = types.StringValue(v.String())
			data.IperfServerVariable = types.StringNull()
		}
	} else {
		data.IperfServer = types.StringNull()
		data.IperfServerVariable = types.StringNull()
	}
	if value := res.Get(path + "nat.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Nat = types.BoolNull()

		} else if value.String() == "ignore" {
			data.Nat = types.BoolValue(false)

		} else if value.String() == "constant" {
			v := res.Get(path + "nat.vipValue")
			data.Nat = types.BoolValue(v.Bool())

		}
	} else if value := res.Get(path + "nat"); value.Exists() {
		data.Nat = types.BoolValue(true)

	} else {
		data.Nat = types.BoolNull()

	}
	if value := res.Get(path + "nat.nat-choice.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.NatType = types.StringNull()

			v := res.Get(path + "nat.nat-choice.vipVariableName")
			data.NatTypeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.NatType = types.StringNull()
			data.NatTypeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "nat.nat-choice.vipValue")
			data.NatType = types.StringValue(v.String())
			data.NatTypeVariable = types.StringNull()
		}
	} else {
		data.NatType = types.StringNull()
		data.NatTypeVariable = types.StringNull()
	}
	if value := res.Get(path + "nat.udp-timeout.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.UdpTimeout = types.Int64Null()

			v := res.Get(path + "nat.udp-timeout.vipVariableName")
			data.UdpTimeoutVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.UdpTimeout = types.Int64Null()
			data.UdpTimeoutVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "nat.udp-timeout.vipValue")
			data.UdpTimeout = types.Int64Value(v.Int())
			data.UdpTimeoutVariable = types.StringNull()
		}
	} else {
		data.UdpTimeout = types.Int64Null()
		data.UdpTimeoutVariable = types.StringNull()
	}
	if value := res.Get(path + "nat.tcp-timeout.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TcpTimeout = types.Int64Null()

			v := res.Get(path + "nat.tcp-timeout.vipVariableName")
			data.TcpTimeoutVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TcpTimeout = types.Int64Null()
			data.TcpTimeoutVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "nat.tcp-timeout.vipValue")
			data.TcpTimeout = types.Int64Value(v.Int())
			data.TcpTimeoutVariable = types.StringNull()
		}
	} else {
		data.TcpTimeout = types.Int64Null()
		data.TcpTimeoutVariable = types.StringNull()
	}
	if value := res.Get(path + "nat.natpool.range-start.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.NatPoolRangeStart = types.StringNull()

			v := res.Get(path + "nat.natpool.range-start.vipVariableName")
			data.NatPoolRangeStartVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.NatPoolRangeStart = types.StringNull()
			data.NatPoolRangeStartVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "nat.natpool.range-start.vipValue")
			data.NatPoolRangeStart = types.StringValue(v.String())
			data.NatPoolRangeStartVariable = types.StringNull()
		}
	} else {
		data.NatPoolRangeStart = types.StringNull()
		data.NatPoolRangeStartVariable = types.StringNull()
	}
	if value := res.Get(path + "nat.natpool.range-end.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.NatPoolRangeEnd = types.StringNull()

			v := res.Get(path + "nat.natpool.range-end.vipVariableName")
			data.NatPoolRangeEndVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.NatPoolRangeEnd = types.StringNull()
			data.NatPoolRangeEndVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "nat.natpool.range-end.vipValue")
			data.NatPoolRangeEnd = types.StringValue(v.String())
			data.NatPoolRangeEndVariable = types.StringNull()
		}
	} else {
		data.NatPoolRangeEnd = types.StringNull()
		data.NatPoolRangeEndVariable = types.StringNull()
	}
	if value := res.Get(path + "nat.overload.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.NatOverload = types.BoolNull()

			v := res.Get(path + "nat.overload.vipVariableName")
			data.NatOverloadVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.NatOverload = types.BoolNull()
			data.NatOverloadVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "nat.overload.vipValue")
			data.NatOverload = types.BoolValue(v.Bool())
			data.NatOverloadVariable = types.StringNull()
		}
	} else {
		data.NatOverload = types.BoolNull()
		data.NatOverloadVariable = types.StringNull()
	}
	if value := res.Get(path + "nat.interface.loopback-interface.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.NatInsideSourceLoopbackInterface = types.StringNull()

			v := res.Get(path + "nat.interface.loopback-interface.vipVariableName")
			data.NatInsideSourceLoopbackInterfaceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.NatInsideSourceLoopbackInterface = types.StringNull()
			data.NatInsideSourceLoopbackInterfaceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "nat.interface.loopback-interface.vipValue")
			data.NatInsideSourceLoopbackInterface = types.StringValue(v.String())
			data.NatInsideSourceLoopbackInterfaceVariable = types.StringNull()
		}
	} else {
		data.NatInsideSourceLoopbackInterface = types.StringNull()
		data.NatInsideSourceLoopbackInterfaceVariable = types.StringNull()
	}
	if value := res.Get(path + "nat.natpool.prefix-length.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.NatPoolPrefixLength = types.Int64Null()

			v := res.Get(path + "nat.natpool.prefix-length.vipVariableName")
			data.NatPoolPrefixLengthVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.NatPoolPrefixLength = types.Int64Null()
			data.NatPoolPrefixLengthVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "nat.natpool.prefix-length.vipValue")
			data.NatPoolPrefixLength = types.Int64Value(v.Int())
			data.NatPoolPrefixLengthVariable = types.StringNull()
		}
	} else {
		data.NatPoolPrefixLength = types.Int64Null()
		data.NatPoolPrefixLengthVariable = types.StringNull()
	}
	if value := res.Get(path + "nat64.enable.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Ipv6Nat = types.BoolNull()

			v := res.Get(path + "nat64.enable.vipVariableName")
			data.Ipv6NatVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Ipv6Nat = types.BoolNull()
			data.Ipv6NatVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "nat64.enable.vipValue")
			data.Ipv6Nat = types.BoolValue(v.Bool())
			data.Ipv6NatVariable = types.StringNull()
		}
	} else {
		data.Ipv6Nat = types.BoolNull()
		data.Ipv6NatVariable = types.StringNull()
	}
	if value := res.Get(path + "nat64.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Nat64Interface = types.BoolNull()

		} else if value.String() == "ignore" {
			data.Nat64Interface = types.BoolValue(false)

		} else if value.String() == "constant" {
			v := res.Get(path + "nat64.vipValue")
			data.Nat64Interface = types.BoolValue(v.Bool())

		}
	} else if value := res.Get(path + "nat64"); value.Exists() {
		data.Nat64Interface = types.BoolValue(true)

	} else {
		data.Nat64Interface = types.BoolNull()

	}
	if value := res.Get(path + "nat66.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Nat66Interface = types.BoolNull()

		} else if value.String() == "ignore" {
			data.Nat66Interface = types.BoolValue(false)

		} else if value.String() == "constant" {
			v := res.Get(path + "nat66.vipValue")
			data.Nat66Interface = types.BoolValue(v.Bool())

		}
	} else if value := res.Get(path + "nat66"); value.Exists() {
		data.Nat66Interface = types.BoolValue(true)

	} else {
		data.Nat66Interface = types.BoolNull()

	}
	if value := res.Get(path + "nat66.static-nat66.vipValue"); len(value.Array()) > 0 {
		data.StaticNat66Entries = make([]CiscoVPNInterfaceStaticNat66Entries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNInterfaceStaticNat66Entries{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("source-prefix.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourcePrefix = types.StringNull()

					cv := v.Get("source-prefix.vipVariableName")
					item.SourcePrefixVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourcePrefix = types.StringNull()
					item.SourcePrefixVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-prefix.vipValue")
					item.SourcePrefix = types.StringValue(cv.String())
					item.SourcePrefixVariable = types.StringNull()
				}
			} else {
				item.SourcePrefix = types.StringNull()
				item.SourcePrefixVariable = types.StringNull()
			}
			if cValue := v.Get("translated-source-prefix.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TranslatedSourcePrefix = types.StringNull()

					cv := v.Get("translated-source-prefix.vipVariableName")
					item.TranslatedSourcePrefixVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TranslatedSourcePrefix = types.StringNull()
					item.TranslatedSourcePrefixVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("translated-source-prefix.vipValue")
					item.TranslatedSourcePrefix = types.StringValue(cv.String())
					item.TranslatedSourcePrefixVariable = types.StringNull()
				}
			} else {
				item.TranslatedSourcePrefix = types.StringNull()
				item.TranslatedSourcePrefixVariable = types.StringNull()
			}
			if cValue := v.Get("source-vpn-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceVpnId = types.Int64Null()

					cv := v.Get("source-vpn-id.vipVariableName")
					item.SourceVpnIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceVpnId = types.Int64Null()
					item.SourceVpnIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-vpn-id.vipValue")
					item.SourceVpnId = types.Int64Value(cv.Int())
					item.SourceVpnIdVariable = types.StringNull()
				}
			} else {
				item.SourceVpnId = types.Int64Null()
				item.SourceVpnIdVariable = types.StringNull()
			}
			data.StaticNat66Entries = append(data.StaticNat66Entries, item)
			return true
		})
	} else {
		if len(data.StaticNat66Entries) > 0 {
			data.StaticNat66Entries = []CiscoVPNInterfaceStaticNat66Entries{}
		}
	}
	if value := res.Get(path + "nat.static.vipValue"); len(value.Array()) > 0 {
		data.StaticNatEntries = make([]CiscoVPNInterfaceStaticNatEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNInterfaceStaticNatEntries{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("source-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceIp = types.StringNull()

					cv := v.Get("source-ip.vipVariableName")
					item.SourceIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceIp = types.StringNull()
					item.SourceIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-ip.vipValue")
					item.SourceIp = types.StringValue(cv.String())
					item.SourceIpVariable = types.StringNull()
				}
			} else {
				item.SourceIp = types.StringNull()
				item.SourceIpVariable = types.StringNull()
			}
			if cValue := v.Get("translate-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TranslateIp = types.StringNull()

					cv := v.Get("translate-ip.vipVariableName")
					item.TranslateIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TranslateIp = types.StringNull()
					item.TranslateIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("translate-ip.vipValue")
					item.TranslateIp = types.StringValue(cv.String())
					item.TranslateIpVariable = types.StringNull()
				}
			} else {
				item.TranslateIp = types.StringNull()
				item.TranslateIpVariable = types.StringNull()
			}
			if cValue := v.Get("static-nat-direction.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.StaticNatDirection = types.StringNull()

					cv := v.Get("static-nat-direction.vipVariableName")
					item.StaticNatDirectionVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.StaticNatDirection = types.StringNull()
					item.StaticNatDirectionVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("static-nat-direction.vipValue")
					item.StaticNatDirection = types.StringValue(cv.String())
					item.StaticNatDirectionVariable = types.StringNull()
				}
			} else {
				item.StaticNatDirection = types.StringNull()
				item.StaticNatDirectionVariable = types.StringNull()
			}
			if cValue := v.Get("source-vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceVpnId = types.Int64Null()

					cv := v.Get("source-vpn.vipVariableName")
					item.SourceVpnIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceVpnId = types.Int64Null()
					item.SourceVpnIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-vpn.vipValue")
					item.SourceVpnId = types.Int64Value(cv.Int())
					item.SourceVpnIdVariable = types.StringNull()
				}
			} else {
				item.SourceVpnId = types.Int64Null()
				item.SourceVpnIdVariable = types.StringNull()
			}
			data.StaticNatEntries = append(data.StaticNatEntries, item)
			return true
		})
	} else {
		if len(data.StaticNatEntries) > 0 {
			data.StaticNatEntries = []CiscoVPNInterfaceStaticNatEntries{}
		}
	}
	if value := res.Get(path + "nat.static-port-forward.vipValue"); len(value.Array()) > 0 {
		data.StaticPortForwardEntries = make([]CiscoVPNInterfaceStaticPortForwardEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNInterfaceStaticPortForwardEntries{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("source-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceIp = types.StringNull()

					cv := v.Get("source-ip.vipVariableName")
					item.SourceIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceIp = types.StringNull()
					item.SourceIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-ip.vipValue")
					item.SourceIp = types.StringValue(cv.String())
					item.SourceIpVariable = types.StringNull()
				}
			} else {
				item.SourceIp = types.StringNull()
				item.SourceIpVariable = types.StringNull()
			}
			if cValue := v.Get("translate-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TranslateIp = types.StringNull()

					cv := v.Get("translate-ip.vipVariableName")
					item.TranslateIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TranslateIp = types.StringNull()
					item.TranslateIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("translate-ip.vipValue")
					item.TranslateIp = types.StringValue(cv.String())
					item.TranslateIpVariable = types.StringNull()
				}
			} else {
				item.TranslateIp = types.StringNull()
				item.TranslateIpVariable = types.StringNull()
			}
			if cValue := v.Get("static-nat-direction.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.StaticNatDirection = types.StringNull()

					cv := v.Get("static-nat-direction.vipVariableName")
					item.StaticNatDirectionVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.StaticNatDirection = types.StringNull()
					item.StaticNatDirectionVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("static-nat-direction.vipValue")
					item.StaticNatDirection = types.StringValue(cv.String())
					item.StaticNatDirectionVariable = types.StringNull()
				}
			} else {
				item.StaticNatDirection = types.StringNull()
				item.StaticNatDirectionVariable = types.StringNull()
			}
			if cValue := v.Get("source-port.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourcePort = types.Int64Null()

					cv := v.Get("source-port.vipVariableName")
					item.SourcePortVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourcePort = types.Int64Null()
					item.SourcePortVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-port.vipValue")
					item.SourcePort = types.Int64Value(cv.Int())
					item.SourcePortVariable = types.StringNull()
				}
			} else {
				item.SourcePort = types.Int64Null()
				item.SourcePortVariable = types.StringNull()
			}
			if cValue := v.Get("translate-port.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TranslatePort = types.Int64Null()

					cv := v.Get("translate-port.vipVariableName")
					item.TranslatePortVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TranslatePort = types.Int64Null()
					item.TranslatePortVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("translate-port.vipValue")
					item.TranslatePort = types.Int64Value(cv.Int())
					item.TranslatePortVariable = types.StringNull()
				}
			} else {
				item.TranslatePort = types.Int64Null()
				item.TranslatePortVariable = types.StringNull()
			}
			if cValue := v.Get("proto.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Protocol = types.StringNull()

					cv := v.Get("proto.vipVariableName")
					item.ProtocolVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Protocol = types.StringNull()
					item.ProtocolVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("proto.vipValue")
					item.Protocol = types.StringValue(cv.String())
					item.ProtocolVariable = types.StringNull()
				}
			} else {
				item.Protocol = types.StringNull()
				item.ProtocolVariable = types.StringNull()
			}
			if cValue := v.Get("source-vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceVpnId = types.Int64Null()

					cv := v.Get("source-vpn.vipVariableName")
					item.SourceVpnIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceVpnId = types.Int64Null()
					item.SourceVpnIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-vpn.vipValue")
					item.SourceVpnId = types.Int64Value(cv.Int())
					item.SourceVpnIdVariable = types.StringNull()
				}
			} else {
				item.SourceVpnId = types.Int64Null()
				item.SourceVpnIdVariable = types.StringNull()
			}
			data.StaticPortForwardEntries = append(data.StaticPortForwardEntries, item)
			return true
		})
	} else {
		if len(data.StaticPortForwardEntries) > 0 {
			data.StaticPortForwardEntries = []CiscoVPNInterfaceStaticPortForwardEntries{}
		}
	}
	if value := res.Get(path + "tunnel-interface.enable-core-region.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.EnableCoreRegion = types.BoolNull()

		} else if value.String() == "ignore" {
			data.EnableCoreRegion = types.BoolNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.enable-core-region.vipValue")
			data.EnableCoreRegion = types.BoolValue(v.Bool())

		}
	} else {
		data.EnableCoreRegion = types.BoolNull()

	}
	if value := res.Get(path + "tunnel-interface.core-region.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.CoreRegion = types.StringNull()

			v := res.Get(path + "tunnel-interface.core-region.vipVariableName")
			data.CoreRegionVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.CoreRegion = types.StringNull()
			data.CoreRegionVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.core-region.vipValue")
			data.CoreRegion = types.StringValue(v.String())
			data.CoreRegionVariable = types.StringNull()
		}
	} else {
		data.CoreRegion = types.StringNull()
		data.CoreRegionVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.secondary-region.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SecondaryRegion = types.StringNull()

			v := res.Get(path + "tunnel-interface.secondary-region.vipVariableName")
			data.SecondaryRegionVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SecondaryRegion = types.StringNull()
			data.SecondaryRegionVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.secondary-region.vipValue")
			data.SecondaryRegion = types.StringValue(v.String())
			data.SecondaryRegionVariable = types.StringNull()
		}
	} else {
		data.SecondaryRegion = types.StringNull()
		data.SecondaryRegionVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.encapsulation.vipValue"); len(value.Array()) > 0 {
		data.TunnelInterfaceEncapsulations = make([]CiscoVPNInterfaceTunnelInterfaceEncapsulations, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNInterfaceTunnelInterfaceEncapsulations{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("encap.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Encapsulation = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Encapsulation = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("encap.vipValue")
					item.Encapsulation = types.StringValue(cv.String())

				}
			} else {
				item.Encapsulation = types.StringNull()

			}
			if cValue := v.Get("preference.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Preference = types.Int64Null()

					cv := v.Get("preference.vipVariableName")
					item.PreferenceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Preference = types.Int64Null()
					item.PreferenceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("preference.vipValue")
					item.Preference = types.Int64Value(cv.Int())
					item.PreferenceVariable = types.StringNull()
				}
			} else {
				item.Preference = types.Int64Null()
				item.PreferenceVariable = types.StringNull()
			}
			if cValue := v.Get("weight.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Weight = types.Int64Null()

					cv := v.Get("weight.vipVariableName")
					item.WeightVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Weight = types.Int64Null()
					item.WeightVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("weight.vipValue")
					item.Weight = types.Int64Value(cv.Int())
					item.WeightVariable = types.StringNull()
				}
			} else {
				item.Weight = types.Int64Null()
				item.WeightVariable = types.StringNull()
			}
			data.TunnelInterfaceEncapsulations = append(data.TunnelInterfaceEncapsulations, item)
			return true
		})
	} else {
		if len(data.TunnelInterfaceEncapsulations) > 0 {
			data.TunnelInterfaceEncapsulations = []CiscoVPNInterfaceTunnelInterfaceEncapsulations{}
		}
	}
	if value := res.Get(path + "tunnel-interface.border.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceBorder = types.BoolNull()

			v := res.Get(path + "tunnel-interface.border.vipVariableName")
			data.TunnelInterfaceBorderVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceBorder = types.BoolNull()
			data.TunnelInterfaceBorderVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.border.vipValue")
			data.TunnelInterfaceBorder = types.BoolValue(v.Bool())
			data.TunnelInterfaceBorderVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceBorder = types.BoolNull()
		data.TunnelInterfaceBorderVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.tunnel-qos.mode.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelQosMode = types.StringNull()

			v := res.Get(path + "tunnel-interface.tunnel-qos.mode.vipVariableName")
			data.TunnelQosModeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelQosMode = types.StringNull()
			data.TunnelQosModeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.tunnel-qos.mode.vipValue")
			data.TunnelQosMode = types.StringValue(v.String())
			data.TunnelQosModeVariable = types.StringNull()
		}
	} else {
		data.TunnelQosMode = types.StringNull()
		data.TunnelQosModeVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.tunnels-bandwidth.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelBandwidth = types.Int64Null()

			v := res.Get(path + "tunnel-interface.tunnels-bandwidth.vipVariableName")
			data.TunnelBandwidthVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelBandwidth = types.Int64Null()
			data.TunnelBandwidthVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.tunnels-bandwidth.vipValue")
			data.TunnelBandwidth = types.Int64Value(v.Int())
			data.TunnelBandwidthVariable = types.StringNull()
		}
	} else {
		data.TunnelBandwidth = types.Int64Null()
		data.TunnelBandwidthVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.group.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.TunnelInterfaceGroups = types.SetNull(types.Int64Type)

			v := res.Get(path + "tunnel-interface.group.vipVariableName")
			data.TunnelInterfaceGroupsVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceGroups = types.SetNull(types.Int64Type)
			data.TunnelInterfaceGroupsVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.group.vipValue")
			data.TunnelInterfaceGroups = helpers.GetInt64Set(v.Array())
			data.TunnelInterfaceGroupsVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceGroups = types.SetNull(types.Int64Type)
		data.TunnelInterfaceGroupsVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.color.value.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceColor = types.StringNull()

			v := res.Get(path + "tunnel-interface.color.value.vipVariableName")
			data.TunnelInterfaceColorVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceColor = types.StringNull()
			data.TunnelInterfaceColorVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.color.value.vipValue")
			data.TunnelInterfaceColor = types.StringValue(v.String())
			data.TunnelInterfaceColorVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceColor = types.StringNull()
		data.TunnelInterfaceColorVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.max-control-connections.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceMaxControlConnections = types.Int64Null()

			v := res.Get(path + "tunnel-interface.max-control-connections.vipVariableName")
			data.TunnelInterfaceMaxControlConnectionsVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceMaxControlConnections = types.Int64Null()
			data.TunnelInterfaceMaxControlConnectionsVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.max-control-connections.vipValue")
			data.TunnelInterfaceMaxControlConnections = types.Int64Value(v.Int())
			data.TunnelInterfaceMaxControlConnectionsVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceMaxControlConnections = types.Int64Null()
		data.TunnelInterfaceMaxControlConnectionsVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.control-connections.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceControlConnections = types.BoolNull()

			v := res.Get(path + "tunnel-interface.control-connections.vipVariableName")
			data.TunnelInterfaceControlConnectionsVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceControlConnections = types.BoolNull()
			data.TunnelInterfaceControlConnectionsVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.control-connections.vipValue")
			data.TunnelInterfaceControlConnections = types.BoolValue(v.Bool())
			data.TunnelInterfaceControlConnectionsVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceControlConnections = types.BoolNull()
		data.TunnelInterfaceControlConnectionsVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.vbond-as-stun-server.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceVbondAsStunServer = types.BoolNull()

			v := res.Get(path + "tunnel-interface.vbond-as-stun-server.vipVariableName")
			data.TunnelInterfaceVbondAsStunServerVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceVbondAsStunServer = types.BoolNull()
			data.TunnelInterfaceVbondAsStunServerVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.vbond-as-stun-server.vipValue")
			data.TunnelInterfaceVbondAsStunServer = types.BoolValue(v.Bool())
			data.TunnelInterfaceVbondAsStunServerVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceVbondAsStunServer = types.BoolNull()
		data.TunnelInterfaceVbondAsStunServerVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.exclude-controller-group-list.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.TunnelInterfaceExcludeControllerGroupList = types.SetNull(types.Int64Type)

			v := res.Get(path + "tunnel-interface.exclude-controller-group-list.vipVariableName")
			data.TunnelInterfaceExcludeControllerGroupListVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceExcludeControllerGroupList = types.SetNull(types.Int64Type)
			data.TunnelInterfaceExcludeControllerGroupListVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.exclude-controller-group-list.vipValue")
			data.TunnelInterfaceExcludeControllerGroupList = helpers.GetInt64Set(v.Array())
			data.TunnelInterfaceExcludeControllerGroupListVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceExcludeControllerGroupList = types.SetNull(types.Int64Type)
		data.TunnelInterfaceExcludeControllerGroupListVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.vmanage-connection-preference.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceVmanageConnectionPreference = types.Int64Null()

			v := res.Get(path + "tunnel-interface.vmanage-connection-preference.vipVariableName")
			data.TunnelInterfaceVmanageConnectionPreferenceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceVmanageConnectionPreference = types.Int64Null()
			data.TunnelInterfaceVmanageConnectionPreferenceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.vmanage-connection-preference.vipValue")
			data.TunnelInterfaceVmanageConnectionPreference = types.Int64Value(v.Int())
			data.TunnelInterfaceVmanageConnectionPreferenceVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceVmanageConnectionPreference = types.Int64Null()
		data.TunnelInterfaceVmanageConnectionPreferenceVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.port-hop.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfacePortHop = types.BoolNull()

			v := res.Get(path + "tunnel-interface.port-hop.vipVariableName")
			data.TunnelInterfacePortHopVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfacePortHop = types.BoolNull()
			data.TunnelInterfacePortHopVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.port-hop.vipValue")
			data.TunnelInterfacePortHop = types.BoolValue(v.Bool())
			data.TunnelInterfacePortHopVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfacePortHop = types.BoolNull()
		data.TunnelInterfacePortHopVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.color.restrict.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceColorRestrict = types.BoolNull()

			v := res.Get(path + "tunnel-interface.color.restrict.vipVariableName")
			data.TunnelInterfaceColorRestrictVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceColorRestrict = types.BoolNull()
			data.TunnelInterfaceColorRestrictVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.color.restrict.vipValue")
			data.TunnelInterfaceColorRestrict = types.BoolValue(v.Bool())
			data.TunnelInterfaceColorRestrictVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceColorRestrict = types.BoolNull()
		data.TunnelInterfaceColorRestrictVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.tloc-extension-gre-to.dst-ip.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceGreTunnelDestinationIp = types.StringNull()

			v := res.Get(path + "tunnel-interface.tloc-extension-gre-to.dst-ip.vipVariableName")
			data.TunnelInterfaceGreTunnelDestinationIpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceGreTunnelDestinationIp = types.StringNull()
			data.TunnelInterfaceGreTunnelDestinationIpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.tloc-extension-gre-to.dst-ip.vipValue")
			data.TunnelInterfaceGreTunnelDestinationIp = types.StringValue(v.String())
			data.TunnelInterfaceGreTunnelDestinationIpVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceGreTunnelDestinationIp = types.StringNull()
		data.TunnelInterfaceGreTunnelDestinationIpVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.carrier.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceCarrier = types.StringNull()

			v := res.Get(path + "tunnel-interface.carrier.vipVariableName")
			data.TunnelInterfaceCarrierVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceCarrier = types.StringNull()
			data.TunnelInterfaceCarrierVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.carrier.vipValue")
			data.TunnelInterfaceCarrier = types.StringValue(v.String())
			data.TunnelInterfaceCarrierVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceCarrier = types.StringNull()
		data.TunnelInterfaceCarrierVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.nat-refresh-interval.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceNatRefreshInterval = types.Int64Null()

			v := res.Get(path + "tunnel-interface.nat-refresh-interval.vipVariableName")
			data.TunnelInterfaceNatRefreshIntervalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceNatRefreshInterval = types.Int64Null()
			data.TunnelInterfaceNatRefreshIntervalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.nat-refresh-interval.vipValue")
			data.TunnelInterfaceNatRefreshInterval = types.Int64Value(v.Int())
			data.TunnelInterfaceNatRefreshIntervalVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceNatRefreshInterval = types.Int64Null()
		data.TunnelInterfaceNatRefreshIntervalVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.hello-interval.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceHelloInterval = types.Int64Null()

			v := res.Get(path + "tunnel-interface.hello-interval.vipVariableName")
			data.TunnelInterfaceHelloIntervalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceHelloInterval = types.Int64Null()
			data.TunnelInterfaceHelloIntervalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.hello-interval.vipValue")
			data.TunnelInterfaceHelloInterval = types.Int64Value(v.Int())
			data.TunnelInterfaceHelloIntervalVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceHelloInterval = types.Int64Null()
		data.TunnelInterfaceHelloIntervalVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.hello-tolerance.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceHelloTolerance = types.Int64Null()

			v := res.Get(path + "tunnel-interface.hello-tolerance.vipVariableName")
			data.TunnelInterfaceHelloToleranceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceHelloTolerance = types.Int64Null()
			data.TunnelInterfaceHelloToleranceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.hello-tolerance.vipValue")
			data.TunnelInterfaceHelloTolerance = types.Int64Value(v.Int())
			data.TunnelInterfaceHelloToleranceVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceHelloTolerance = types.Int64Null()
		data.TunnelInterfaceHelloToleranceVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.bind.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceBindLoopbackTunnel = types.StringNull()

			v := res.Get(path + "tunnel-interface.bind.vipVariableName")
			data.TunnelInterfaceBindLoopbackTunnelVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceBindLoopbackTunnel = types.StringNull()
			data.TunnelInterfaceBindLoopbackTunnelVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.bind.vipValue")
			data.TunnelInterfaceBindLoopbackTunnel = types.StringValue(v.String())
			data.TunnelInterfaceBindLoopbackTunnelVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceBindLoopbackTunnel = types.StringNull()
		data.TunnelInterfaceBindLoopbackTunnelVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.last-resort-circuit.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceLastResortCircuit = types.BoolNull()

			v := res.Get(path + "tunnel-interface.last-resort-circuit.vipVariableName")
			data.TunnelInterfaceLastResortCircuitVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceLastResortCircuit = types.BoolNull()
			data.TunnelInterfaceLastResortCircuitVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.last-resort-circuit.vipValue")
			data.TunnelInterfaceLastResortCircuit = types.BoolValue(v.Bool())
			data.TunnelInterfaceLastResortCircuitVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceLastResortCircuit = types.BoolNull()
		data.TunnelInterfaceLastResortCircuitVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.low-bandwidth-link.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceLowBandwidthLink = types.BoolNull()

			v := res.Get(path + "tunnel-interface.low-bandwidth-link.vipVariableName")
			data.TunnelInterfaceLowBandwidthLinkVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceLowBandwidthLink = types.BoolNull()
			data.TunnelInterfaceLowBandwidthLinkVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.low-bandwidth-link.vipValue")
			data.TunnelInterfaceLowBandwidthLink = types.BoolValue(v.Bool())
			data.TunnelInterfaceLowBandwidthLinkVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceLowBandwidthLink = types.BoolNull()
		data.TunnelInterfaceLowBandwidthLinkVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.tunnel-tcp-mss-adjust.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceTunnelTcpMss = types.Int64Null()

			v := res.Get(path + "tunnel-interface.tunnel-tcp-mss-adjust.vipVariableName")
			data.TunnelInterfaceTunnelTcpMssVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceTunnelTcpMss = types.Int64Null()
			data.TunnelInterfaceTunnelTcpMssVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.tunnel-tcp-mss-adjust.vipValue")
			data.TunnelInterfaceTunnelTcpMss = types.Int64Value(v.Int())
			data.TunnelInterfaceTunnelTcpMssVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceTunnelTcpMss = types.Int64Null()
		data.TunnelInterfaceTunnelTcpMssVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.clear-dont-fragment.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceClearDontFragment = types.BoolNull()

			v := res.Get(path + "tunnel-interface.clear-dont-fragment.vipVariableName")
			data.TunnelInterfaceClearDontFragmentVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceClearDontFragment = types.BoolNull()
			data.TunnelInterfaceClearDontFragmentVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.clear-dont-fragment.vipValue")
			data.TunnelInterfaceClearDontFragment = types.BoolValue(v.Bool())
			data.TunnelInterfaceClearDontFragmentVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceClearDontFragment = types.BoolNull()
		data.TunnelInterfaceClearDontFragmentVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.propagate-sgt.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfacePropagateSgt = types.BoolNull()

			v := res.Get(path + "tunnel-interface.propagate-sgt.vipVariableName")
			data.TunnelInterfacePropagateSgtVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfacePropagateSgt = types.BoolNull()
			data.TunnelInterfacePropagateSgtVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.propagate-sgt.vipValue")
			data.TunnelInterfacePropagateSgt = types.BoolValue(v.Bool())
			data.TunnelInterfacePropagateSgtVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfacePropagateSgt = types.BoolNull()
		data.TunnelInterfacePropagateSgtVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.network-broadcast.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceNetworkBroadcast = types.BoolNull()

			v := res.Get(path + "tunnel-interface.network-broadcast.vipVariableName")
			data.TunnelInterfaceNetworkBroadcastVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceNetworkBroadcast = types.BoolNull()
			data.TunnelInterfaceNetworkBroadcastVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.network-broadcast.vipValue")
			data.TunnelInterfaceNetworkBroadcast = types.BoolValue(v.Bool())
			data.TunnelInterfaceNetworkBroadcastVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceNetworkBroadcast = types.BoolNull()
		data.TunnelInterfaceNetworkBroadcastVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.all.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowAll = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.all.vipVariableName")
			data.TunnelInterfaceAllowAllVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowAll = types.BoolNull()
			data.TunnelInterfaceAllowAllVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.all.vipValue")
			data.TunnelInterfaceAllowAll = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowAllVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowAll = types.BoolNull()
		data.TunnelInterfaceAllowAllVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.bgp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowBgp = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.bgp.vipVariableName")
			data.TunnelInterfaceAllowBgpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowBgp = types.BoolNull()
			data.TunnelInterfaceAllowBgpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.bgp.vipValue")
			data.TunnelInterfaceAllowBgp = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowBgpVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowBgp = types.BoolNull()
		data.TunnelInterfaceAllowBgpVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.dhcp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowDhcp = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.dhcp.vipVariableName")
			data.TunnelInterfaceAllowDhcpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowDhcp = types.BoolNull()
			data.TunnelInterfaceAllowDhcpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.dhcp.vipValue")
			data.TunnelInterfaceAllowDhcp = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowDhcpVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowDhcp = types.BoolNull()
		data.TunnelInterfaceAllowDhcpVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.dns.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowDns = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.dns.vipVariableName")
			data.TunnelInterfaceAllowDnsVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowDns = types.BoolNull()
			data.TunnelInterfaceAllowDnsVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.dns.vipValue")
			data.TunnelInterfaceAllowDns = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowDnsVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowDns = types.BoolNull()
		data.TunnelInterfaceAllowDnsVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.icmp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowIcmp = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.icmp.vipVariableName")
			data.TunnelInterfaceAllowIcmpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowIcmp = types.BoolNull()
			data.TunnelInterfaceAllowIcmpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.icmp.vipValue")
			data.TunnelInterfaceAllowIcmp = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowIcmpVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowIcmp = types.BoolNull()
		data.TunnelInterfaceAllowIcmpVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.sshd.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowSsh = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.sshd.vipVariableName")
			data.TunnelInterfaceAllowSshVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowSsh = types.BoolNull()
			data.TunnelInterfaceAllowSshVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.sshd.vipValue")
			data.TunnelInterfaceAllowSsh = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowSshVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowSsh = types.BoolNull()
		data.TunnelInterfaceAllowSshVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.netconf.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowNetconf = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.netconf.vipVariableName")
			data.TunnelInterfaceAllowNetconfVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowNetconf = types.BoolNull()
			data.TunnelInterfaceAllowNetconfVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.netconf.vipValue")
			data.TunnelInterfaceAllowNetconf = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowNetconfVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowNetconf = types.BoolNull()
		data.TunnelInterfaceAllowNetconfVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.ntp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowNtp = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.ntp.vipVariableName")
			data.TunnelInterfaceAllowNtpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowNtp = types.BoolNull()
			data.TunnelInterfaceAllowNtpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.ntp.vipValue")
			data.TunnelInterfaceAllowNtp = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowNtpVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowNtp = types.BoolNull()
		data.TunnelInterfaceAllowNtpVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.ospf.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowOspf = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.ospf.vipVariableName")
			data.TunnelInterfaceAllowOspfVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowOspf = types.BoolNull()
			data.TunnelInterfaceAllowOspfVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.ospf.vipValue")
			data.TunnelInterfaceAllowOspf = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowOspfVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowOspf = types.BoolNull()
		data.TunnelInterfaceAllowOspfVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.stun.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowStun = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.stun.vipVariableName")
			data.TunnelInterfaceAllowStunVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowStun = types.BoolNull()
			data.TunnelInterfaceAllowStunVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.stun.vipValue")
			data.TunnelInterfaceAllowStun = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowStunVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowStun = types.BoolNull()
		data.TunnelInterfaceAllowStunVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.snmp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowSnmp = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.snmp.vipVariableName")
			data.TunnelInterfaceAllowSnmpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowSnmp = types.BoolNull()
			data.TunnelInterfaceAllowSnmpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.snmp.vipValue")
			data.TunnelInterfaceAllowSnmp = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowSnmpVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowSnmp = types.BoolNull()
		data.TunnelInterfaceAllowSnmpVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-interface.allow-service.https.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelInterfaceAllowHttps = types.BoolNull()

			v := res.Get(path + "tunnel-interface.allow-service.https.vipVariableName")
			data.TunnelInterfaceAllowHttpsVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelInterfaceAllowHttps = types.BoolNull()
			data.TunnelInterfaceAllowHttpsVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-interface.allow-service.https.vipValue")
			data.TunnelInterfaceAllowHttps = types.BoolValue(v.Bool())
			data.TunnelInterfaceAllowHttpsVariable = types.StringNull()
		}
	} else {
		data.TunnelInterfaceAllowHttps = types.BoolNull()
		data.TunnelInterfaceAllowHttpsVariable = types.StringNull()
	}
	if value := res.Get(path + "media-type.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.MediaType = types.StringNull()

			v := res.Get(path + "media-type.vipVariableName")
			data.MediaTypeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.MediaType = types.StringNull()
			data.MediaTypeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "media-type.vipValue")
			data.MediaType = types.StringValue(v.String())
			data.MediaTypeVariable = types.StringNull()
		}
	} else {
		data.MediaType = types.StringNull()
		data.MediaTypeVariable = types.StringNull()
	}
	if value := res.Get(path + "intrf-mtu.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.InterfaceMtu = types.Int64Null()

			v := res.Get(path + "intrf-mtu.vipVariableName")
			data.InterfaceMtuVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.InterfaceMtu = types.Int64Null()
			data.InterfaceMtuVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "intrf-mtu.vipValue")
			data.InterfaceMtu = types.Int64Value(v.Int())
			data.InterfaceMtuVariable = types.StringNull()
		}
	} else {
		data.InterfaceMtu = types.Int64Null()
		data.InterfaceMtuVariable = types.StringNull()
	}
	if value := res.Get(path + "mtu.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IpMtu = types.Int64Null()

			v := res.Get(path + "mtu.vipVariableName")
			data.IpMtuVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IpMtu = types.Int64Null()
			data.IpMtuVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "mtu.vipValue")
			data.IpMtu = types.Int64Value(v.Int())
			data.IpMtuVariable = types.StringNull()
		}
	} else {
		data.IpMtu = types.Int64Null()
		data.IpMtuVariable = types.StringNull()
	}
	if value := res.Get(path + "tcp-mss-adjust.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TcpMssAdjust = types.Int64Null()

			v := res.Get(path + "tcp-mss-adjust.vipVariableName")
			data.TcpMssAdjustVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TcpMssAdjust = types.Int64Null()
			data.TcpMssAdjustVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tcp-mss-adjust.vipValue")
			data.TcpMssAdjust = types.Int64Value(v.Int())
			data.TcpMssAdjustVariable = types.StringNull()
		}
	} else {
		data.TcpMssAdjust = types.Int64Null()
		data.TcpMssAdjustVariable = types.StringNull()
	}
	if value := res.Get(path + "tloc-extension.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TlocExtension = types.StringNull()

			v := res.Get(path + "tloc-extension.vipVariableName")
			data.TlocExtensionVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TlocExtension = types.StringNull()
			data.TlocExtensionVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tloc-extension.vipValue")
			data.TlocExtension = types.StringValue(v.String())
			data.TlocExtensionVariable = types.StringNull()
		}
	} else {
		data.TlocExtension = types.StringNull()
		data.TlocExtensionVariable = types.StringNull()
	}
	if value := res.Get(path + "load-interval.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.LoadInterval = types.Int64Null()

			v := res.Get(path + "load-interval.vipVariableName")
			data.LoadIntervalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.LoadInterval = types.Int64Null()
			data.LoadIntervalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "load-interval.vipValue")
			data.LoadInterval = types.Int64Value(v.Int())
			data.LoadIntervalVariable = types.StringNull()
		}
	} else {
		data.LoadInterval = types.Int64Null()
		data.LoadIntervalVariable = types.StringNull()
	}
	if value := res.Get(path + "tloc-extension-gre-from.src-ip.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.GreTunnelSourceIp = types.StringNull()

			v := res.Get(path + "tloc-extension-gre-from.src-ip.vipVariableName")
			data.GreTunnelSourceIpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.GreTunnelSourceIp = types.StringNull()
			data.GreTunnelSourceIpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tloc-extension-gre-from.src-ip.vipValue")
			data.GreTunnelSourceIp = types.StringValue(v.String())
			data.GreTunnelSourceIpVariable = types.StringNull()
		}
	} else {
		data.GreTunnelSourceIp = types.StringNull()
		data.GreTunnelSourceIpVariable = types.StringNull()
	}
	if value := res.Get(path + "tloc-extension-gre-from.xconnect.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.GreTunnelXconnect = types.StringNull()

			v := res.Get(path + "tloc-extension-gre-from.xconnect.vipVariableName")
			data.GreTunnelXconnectVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.GreTunnelXconnect = types.StringNull()
			data.GreTunnelXconnectVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tloc-extension-gre-from.xconnect.vipValue")
			data.GreTunnelXconnect = types.StringValue(v.String())
			data.GreTunnelXconnectVariable = types.StringNull()
		}
	} else {
		data.GreTunnelXconnect = types.StringNull()
		data.GreTunnelXconnectVariable = types.StringNull()
	}
	if value := res.Get(path + "mac-address.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.MacAddress = types.StringNull()

			v := res.Get(path + "mac-address.vipVariableName")
			data.MacAddressVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.MacAddress = types.StringNull()
			data.MacAddressVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "mac-address.vipValue")
			data.MacAddress = types.StringValue(v.String())
			data.MacAddressVariable = types.StringNull()
		}
	} else {
		data.MacAddress = types.StringNull()
		data.MacAddressVariable = types.StringNull()
	}
	if value := res.Get(path + "speed.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Speed = types.StringNull()

			v := res.Get(path + "speed.vipVariableName")
			data.SpeedVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Speed = types.StringNull()
			data.SpeedVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "speed.vipValue")
			data.Speed = types.StringValue(v.String())
			data.SpeedVariable = types.StringNull()
		}
	} else {
		data.Speed = types.StringNull()
		data.SpeedVariable = types.StringNull()
	}
	if value := res.Get(path + "duplex.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Duplex = types.StringNull()

			v := res.Get(path + "duplex.vipVariableName")
			data.DuplexVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Duplex = types.StringNull()
			data.DuplexVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "duplex.vipValue")
			data.Duplex = types.StringValue(v.String())
			data.DuplexVariable = types.StringNull()
		}
	} else {
		data.Duplex = types.StringNull()
		data.DuplexVariable = types.StringNull()
	}
	if value := res.Get(path + "shutdown.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Shutdown = types.BoolNull()

			v := res.Get(path + "shutdown.vipVariableName")
			data.ShutdownVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Shutdown = types.BoolNull()
			data.ShutdownVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "shutdown.vipValue")
			data.Shutdown = types.BoolValue(v.Bool())
			data.ShutdownVariable = types.StringNull()
		}
	} else {
		data.Shutdown = types.BoolNull()
		data.ShutdownVariable = types.StringNull()
	}
	if value := res.Get(path + "arp-timeout.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ArpTimeout = types.Int64Null()

			v := res.Get(path + "arp-timeout.vipVariableName")
			data.ArpTimeoutVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ArpTimeout = types.Int64Null()
			data.ArpTimeoutVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "arp-timeout.vipValue")
			data.ArpTimeout = types.Int64Value(v.Int())
			data.ArpTimeoutVariable = types.StringNull()
		}
	} else {
		data.ArpTimeout = types.Int64Null()
		data.ArpTimeoutVariable = types.StringNull()
	}
	if value := res.Get(path + "autonegotiate.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Autonegotiate = types.BoolNull()

			v := res.Get(path + "autonegotiate.vipVariableName")
			data.AutonegotiateVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Autonegotiate = types.BoolNull()
			data.AutonegotiateVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "autonegotiate.vipValue")
			data.Autonegotiate = types.BoolValue(v.Bool())
			data.AutonegotiateVariable = types.StringNull()
		}
	} else {
		data.Autonegotiate = types.BoolNull()
		data.AutonegotiateVariable = types.StringNull()
	}
	if value := res.Get(path + "ip-directed-broadcast.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IpDirectedBroadcast = types.BoolNull()

			v := res.Get(path + "ip-directed-broadcast.vipVariableName")
			data.IpDirectedBroadcastVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IpDirectedBroadcast = types.BoolNull()
			data.IpDirectedBroadcastVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ip-directed-broadcast.vipValue")
			data.IpDirectedBroadcast = types.BoolValue(v.Bool())
			data.IpDirectedBroadcastVariable = types.StringNull()
		}
	} else {
		data.IpDirectedBroadcast = types.BoolNull()
		data.IpDirectedBroadcastVariable = types.StringNull()
	}
	if value := res.Get(path + "icmp-redirect-disable.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IcmpRedirectDisable = types.BoolNull()

			v := res.Get(path + "icmp-redirect-disable.vipVariableName")
			data.IcmpRedirectDisableVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IcmpRedirectDisable = types.BoolNull()
			data.IcmpRedirectDisableVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "icmp-redirect-disable.vipValue")
			data.IcmpRedirectDisable = types.BoolValue(v.Bool())
			data.IcmpRedirectDisableVariable = types.StringNull()
		}
	} else {
		data.IcmpRedirectDisable = types.BoolNull()
		data.IcmpRedirectDisableVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-adaptive.period.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosAdaptivePeriod = types.Int64Null()

			v := res.Get(path + "qos-adaptive.period.vipVariableName")
			data.QosAdaptivePeriodVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosAdaptivePeriod = types.Int64Null()
			data.QosAdaptivePeriodVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-adaptive.period.vipValue")
			data.QosAdaptivePeriod = types.Int64Value(v.Int())
			data.QosAdaptivePeriodVariable = types.StringNull()
		}
	} else {
		data.QosAdaptivePeriod = types.Int64Null()
		data.QosAdaptivePeriodVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-adaptive.downstream.bandwidth-down.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosAdaptiveBandwidthDownstream = types.Int64Null()

			v := res.Get(path + "qos-adaptive.downstream.bandwidth-down.vipVariableName")
			data.QosAdaptiveBandwidthDownstreamVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosAdaptiveBandwidthDownstream = types.Int64Null()
			data.QosAdaptiveBandwidthDownstreamVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-adaptive.downstream.bandwidth-down.vipValue")
			data.QosAdaptiveBandwidthDownstream = types.Int64Value(v.Int())
			data.QosAdaptiveBandwidthDownstreamVariable = types.StringNull()
		}
	} else {
		data.QosAdaptiveBandwidthDownstream = types.Int64Null()
		data.QosAdaptiveBandwidthDownstreamVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-adaptive.downstream.range.dmin.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosAdaptiveMinDownstream = types.Int64Null()

			v := res.Get(path + "qos-adaptive.downstream.range.dmin.vipVariableName")
			data.QosAdaptiveMinDownstreamVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosAdaptiveMinDownstream = types.Int64Null()
			data.QosAdaptiveMinDownstreamVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-adaptive.downstream.range.dmin.vipValue")
			data.QosAdaptiveMinDownstream = types.Int64Value(v.Int())
			data.QosAdaptiveMinDownstreamVariable = types.StringNull()
		}
	} else {
		data.QosAdaptiveMinDownstream = types.Int64Null()
		data.QosAdaptiveMinDownstreamVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-adaptive.downstream.range.dmax.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosAdaptiveMaxDownstream = types.Int64Null()

			v := res.Get(path + "qos-adaptive.downstream.range.dmax.vipVariableName")
			data.QosAdaptiveMaxDownstreamVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosAdaptiveMaxDownstream = types.Int64Null()
			data.QosAdaptiveMaxDownstreamVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-adaptive.downstream.range.dmax.vipValue")
			data.QosAdaptiveMaxDownstream = types.Int64Value(v.Int())
			data.QosAdaptiveMaxDownstreamVariable = types.StringNull()
		}
	} else {
		data.QosAdaptiveMaxDownstream = types.Int64Null()
		data.QosAdaptiveMaxDownstreamVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-adaptive.upstream.bandwidth-up.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosAdaptiveBandwidthUpstream = types.Int64Null()

			v := res.Get(path + "qos-adaptive.upstream.bandwidth-up.vipVariableName")
			data.QosAdaptiveBandwidthUpstreamVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosAdaptiveBandwidthUpstream = types.Int64Null()
			data.QosAdaptiveBandwidthUpstreamVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-adaptive.upstream.bandwidth-up.vipValue")
			data.QosAdaptiveBandwidthUpstream = types.Int64Value(v.Int())
			data.QosAdaptiveBandwidthUpstreamVariable = types.StringNull()
		}
	} else {
		data.QosAdaptiveBandwidthUpstream = types.Int64Null()
		data.QosAdaptiveBandwidthUpstreamVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-adaptive.upstream.range.umin.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosAdaptiveMinUpstream = types.Int64Null()

			v := res.Get(path + "qos-adaptive.upstream.range.umin.vipVariableName")
			data.QosAdaptiveMinUpstreamVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosAdaptiveMinUpstream = types.Int64Null()
			data.QosAdaptiveMinUpstreamVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-adaptive.upstream.range.umin.vipValue")
			data.QosAdaptiveMinUpstream = types.Int64Value(v.Int())
			data.QosAdaptiveMinUpstreamVariable = types.StringNull()
		}
	} else {
		data.QosAdaptiveMinUpstream = types.Int64Null()
		data.QosAdaptiveMinUpstreamVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-adaptive.upstream.range.umax.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosAdaptiveMaxUpstream = types.Int64Null()

			v := res.Get(path + "qos-adaptive.upstream.range.umax.vipVariableName")
			data.QosAdaptiveMaxUpstreamVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosAdaptiveMaxUpstream = types.Int64Null()
			data.QosAdaptiveMaxUpstreamVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-adaptive.upstream.range.umax.vipValue")
			data.QosAdaptiveMaxUpstream = types.Int64Value(v.Int())
			data.QosAdaptiveMaxUpstreamVariable = types.StringNull()
		}
	} else {
		data.QosAdaptiveMaxUpstream = types.Int64Null()
		data.QosAdaptiveMaxUpstreamVariable = types.StringNull()
	}
	if value := res.Get(path + "shaping-rate.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ShapingRate = types.Int64Null()

			v := res.Get(path + "shaping-rate.vipVariableName")
			data.ShapingRateVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ShapingRate = types.Int64Null()
			data.ShapingRateVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "shaping-rate.vipValue")
			data.ShapingRate = types.Int64Value(v.Int())
			data.ShapingRateVariable = types.StringNull()
		}
	} else {
		data.ShapingRate = types.Int64Null()
		data.ShapingRateVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-map.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosMap = types.StringNull()

			v := res.Get(path + "qos-map.vipVariableName")
			data.QosMapVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosMap = types.StringNull()
			data.QosMapVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-map.vipValue")
			data.QosMap = types.StringValue(v.String())
			data.QosMapVariable = types.StringNull()
		}
	} else {
		data.QosMap = types.StringNull()
		data.QosMapVariable = types.StringNull()
	}
	if value := res.Get(path + "qos-map-vpn.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.QosMapVpn = types.StringNull()

			v := res.Get(path + "qos-map-vpn.vipVariableName")
			data.QosMapVpnVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.QosMapVpn = types.StringNull()
			data.QosMapVpnVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "qos-map-vpn.vipValue")
			data.QosMapVpn = types.StringValue(v.String())
			data.QosMapVpnVariable = types.StringNull()
		}
	} else {
		data.QosMapVpn = types.StringNull()
		data.QosMapVpnVariable = types.StringNull()
	}
	if value := res.Get(path + "bandwidth-upstream.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.BandwidthUpstream = types.Int64Null()

			v := res.Get(path + "bandwidth-upstream.vipVariableName")
			data.BandwidthUpstreamVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.BandwidthUpstream = types.Int64Null()
			data.BandwidthUpstreamVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bandwidth-upstream.vipValue")
			data.BandwidthUpstream = types.Int64Value(v.Int())
			data.BandwidthUpstreamVariable = types.StringNull()
		}
	} else {
		data.BandwidthUpstream = types.Int64Null()
		data.BandwidthUpstreamVariable = types.StringNull()
	}
	if value := res.Get(path + "bandwidth-downstream.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.BandwidthDownstream = types.Int64Null()

			v := res.Get(path + "bandwidth-downstream.vipVariableName")
			data.BandwidthDownstreamVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.BandwidthDownstream = types.Int64Null()
			data.BandwidthDownstreamVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bandwidth-downstream.vipValue")
			data.BandwidthDownstream = types.Int64Value(v.Int())
			data.BandwidthDownstreamVariable = types.StringNull()
		}
	} else {
		data.BandwidthDownstream = types.Int64Null()
		data.BandwidthDownstreamVariable = types.StringNull()
	}
	if value := res.Get(path + "block-non-source-ip.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.BlockNonSourceIp = types.BoolNull()

			v := res.Get(path + "block-non-source-ip.vipVariableName")
			data.BlockNonSourceIpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.BlockNonSourceIp = types.BoolNull()
			data.BlockNonSourceIpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "block-non-source-ip.vipValue")
			data.BlockNonSourceIp = types.BoolValue(v.Bool())
			data.BlockNonSourceIpVariable = types.StringNull()
		}
	} else {
		data.BlockNonSourceIp = types.BoolNull()
		data.BlockNonSourceIpVariable = types.StringNull()
	}
	if value := res.Get(path + "rewrite-rule.rule-name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.RewriteRuleName = types.StringNull()

			v := res.Get(path + "rewrite-rule.rule-name.vipVariableName")
			data.RewriteRuleNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.RewriteRuleName = types.StringNull()
			data.RewriteRuleNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "rewrite-rule.rule-name.vipValue")
			data.RewriteRuleName = types.StringValue(v.String())
			data.RewriteRuleNameVariable = types.StringNull()
		}
	} else {
		data.RewriteRuleName = types.StringNull()
		data.RewriteRuleNameVariable = types.StringNull()
	}
	if value := res.Get(path + "access-list.vipValue"); len(value.Array()) > 0 {
		data.AccessLists = make([]CiscoVPNInterfaceAccessLists, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNInterfaceAccessLists{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("direction.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Direction = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Direction = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("direction.vipValue")
					item.Direction = types.StringValue(cv.String())

				}
			} else {
				item.Direction = types.StringNull()

			}
			if cValue := v.Get("acl-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AclName = types.StringNull()

					cv := v.Get("acl-name.vipVariableName")
					item.AclNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AclName = types.StringNull()
					item.AclNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("acl-name.vipValue")
					item.AclName = types.StringValue(cv.String())
					item.AclNameVariable = types.StringNull()
				}
			} else {
				item.AclName = types.StringNull()
				item.AclNameVariable = types.StringNull()
			}
			data.AccessLists = append(data.AccessLists, item)
			return true
		})
	} else {
		if len(data.AccessLists) > 0 {
			data.AccessLists = []CiscoVPNInterfaceAccessLists{}
		}
	}
	if value := res.Get(path + "arp.ip.vipValue"); len(value.Array()) > 0 {
		data.StaticArps = make([]CiscoVPNInterfaceStaticArps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNInterfaceStaticArps{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("addr.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IpAddress = types.StringNull()

					cv := v.Get("addr.vipVariableName")
					item.IpAddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IpAddress = types.StringNull()
					item.IpAddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("addr.vipValue")
					item.IpAddress = types.StringValue(cv.String())
					item.IpAddressVariable = types.StringNull()
				}
			} else {
				item.IpAddress = types.StringNull()
				item.IpAddressVariable = types.StringNull()
			}
			if cValue := v.Get("mac.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Mac = types.StringNull()

					cv := v.Get("mac.vipVariableName")
					item.MacVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Mac = types.StringNull()
					item.MacVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("mac.vipValue")
					item.Mac = types.StringValue(cv.String())
					item.MacVariable = types.StringNull()
				}
			} else {
				item.Mac = types.StringNull()
				item.MacVariable = types.StringNull()
			}
			data.StaticArps = append(data.StaticArps, item)
			return true
		})
	} else {
		if len(data.StaticArps) > 0 {
			data.StaticArps = []CiscoVPNInterfaceStaticArps{}
		}
	}
	if value := res.Get(path + "vrrp.vipValue"); len(value.Array()) > 0 {
		data.Ipv4Vrrps = make([]CiscoVPNInterfaceIpv4Vrrps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNInterfaceIpv4Vrrps{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("grp-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.GroupId = types.Int64Null()

					cv := v.Get("grp-id.vipVariableName")
					item.GroupIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.GroupId = types.Int64Null()
					item.GroupIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("grp-id.vipValue")
					item.GroupId = types.Int64Value(cv.Int())
					item.GroupIdVariable = types.StringNull()
				}
			} else {
				item.GroupId = types.Int64Null()
				item.GroupIdVariable = types.StringNull()
			}
			if cValue := v.Get("priority.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Priority = types.Int64Null()

					cv := v.Get("priority.vipVariableName")
					item.PriorityVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Priority = types.Int64Null()
					item.PriorityVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("priority.vipValue")
					item.Priority = types.Int64Value(cv.Int())
					item.PriorityVariable = types.StringNull()
				}
			} else {
				item.Priority = types.Int64Null()
				item.PriorityVariable = types.StringNull()
			}
			if cValue := v.Get("timer.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Timer = types.Int64Null()

					cv := v.Get("timer.vipVariableName")
					item.TimerVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Timer = types.Int64Null()
					item.TimerVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("timer.vipValue")
					item.Timer = types.Int64Value(cv.Int())
					item.TimerVariable = types.StringNull()
				}
			} else {
				item.Timer = types.Int64Null()
				item.TimerVariable = types.StringNull()
			}
			if cValue := v.Get("track-omp.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TrackOmp = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.TrackOmp = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("track-omp.vipValue")
					item.TrackOmp = types.BoolValue(cv.Bool())

				}
			} else {
				item.TrackOmp = types.BoolNull()

			}
			if cValue := v.Get("track-prefix-list.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TrackPrefixList = types.StringNull()

					cv := v.Get("track-prefix-list.vipVariableName")
					item.TrackPrefixListVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TrackPrefixList = types.StringNull()
					item.TrackPrefixListVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("track-prefix-list.vipValue")
					item.TrackPrefixList = types.StringValue(cv.String())
					item.TrackPrefixListVariable = types.StringNull()
				}
			} else {
				item.TrackPrefixList = types.StringNull()
				item.TrackPrefixListVariable = types.StringNull()
			}
			if cValue := v.Get("ipv4.address.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IpAddress = types.StringNull()

					cv := v.Get("ipv4.address.vipVariableName")
					item.IpAddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IpAddress = types.StringNull()
					item.IpAddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ipv4.address.vipValue")
					item.IpAddress = types.StringValue(cv.String())
					item.IpAddressVariable = types.StringNull()
				}
			} else {
				item.IpAddress = types.StringNull()
				item.IpAddressVariable = types.StringNull()
			}
			if cValue := v.Get("ipv4-secondary.vipValue"); len(cValue.Array()) > 0 {
				item.Ipv4SecondaryAddresses = make([]CiscoVPNInterfaceIpv4VrrpsIpv4SecondaryAddresses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoVPNInterfaceIpv4VrrpsIpv4SecondaryAddresses{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("address.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.IpAddress = types.StringNull()

							ccv := cv.Get("address.vipVariableName")
							cItem.IpAddressVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.IpAddress = types.StringNull()
							cItem.IpAddressVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("address.vipValue")
							cItem.IpAddress = types.StringValue(ccv.String())
							cItem.IpAddressVariable = types.StringNull()
						}
					} else {
						cItem.IpAddress = types.StringNull()
						cItem.IpAddressVariable = types.StringNull()
					}
					item.Ipv4SecondaryAddresses = append(item.Ipv4SecondaryAddresses, cItem)
					return true
				})
			} else {
				if len(item.Ipv4SecondaryAddresses) > 0 {
					item.Ipv4SecondaryAddresses = []CiscoVPNInterfaceIpv4VrrpsIpv4SecondaryAddresses{}
				}
			}
			if cValue := v.Get("tloc-change-pref.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TlocPreferenceChange = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.TlocPreferenceChange = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("tloc-change-pref.vipValue")
					item.TlocPreferenceChange = types.BoolValue(cv.Bool())

				}
			} else {
				item.TlocPreferenceChange = types.BoolNull()

			}
			if cValue := v.Get("value.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TlocPreferenceChangeValue = types.Int64Null()

					cv := v.Get("value.vipVariableName")
					item.TlocPreferenceChangeValueVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TlocPreferenceChangeValue = types.Int64Null()
					item.TlocPreferenceChangeValueVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("value.vipValue")
					item.TlocPreferenceChangeValue = types.Int64Value(cv.Int())
					item.TlocPreferenceChangeValueVariable = types.StringNull()
				}
			} else {
				item.TlocPreferenceChangeValue = types.Int64Null()
				item.TlocPreferenceChangeValueVariable = types.StringNull()
			}
			if cValue := v.Get("tracking-object.vipValue"); len(cValue.Array()) > 0 {
				item.TrackingObjects = make([]CiscoVPNInterfaceIpv4VrrpsTrackingObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoVPNInterfaceIpv4VrrpsTrackingObjects{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("name.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.TrackerId = types.Int64Null()

							ccv := cv.Get("name.vipVariableName")
							cItem.TrackerIdVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.TrackerId = types.Int64Null()
							cItem.TrackerIdVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("name.vipValue")
							cItem.TrackerId = types.Int64Value(ccv.Int())
							cItem.TrackerIdVariable = types.StringNull()
						}
					} else {
						cItem.TrackerId = types.Int64Null()
						cItem.TrackerIdVariable = types.StringNull()
					}
					if ccValue := cv.Get("track-action.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.TrackAction = types.StringNull()

							ccv := cv.Get("track-action.vipVariableName")
							cItem.TrackActionVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.TrackAction = types.StringNull()
							cItem.TrackActionVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("track-action.vipValue")
							cItem.TrackAction = types.StringValue(ccv.String())
							cItem.TrackActionVariable = types.StringNull()
						}
					} else {
						cItem.TrackAction = types.StringNull()
						cItem.TrackActionVariable = types.StringNull()
					}
					if ccValue := cv.Get("decrement.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.DecrementValue = types.Int64Null()

							ccv := cv.Get("decrement.vipVariableName")
							cItem.DecrementValueVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.DecrementValue = types.Int64Null()
							cItem.DecrementValueVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("decrement.vipValue")
							cItem.DecrementValue = types.Int64Value(ccv.Int())
							cItem.DecrementValueVariable = types.StringNull()
						}
					} else {
						cItem.DecrementValue = types.Int64Null()
						cItem.DecrementValueVariable = types.StringNull()
					}
					item.TrackingObjects = append(item.TrackingObjects, cItem)
					return true
				})
			} else {
				if len(item.TrackingObjects) > 0 {
					item.TrackingObjects = []CiscoVPNInterfaceIpv4VrrpsTrackingObjects{}
				}
			}
			data.Ipv4Vrrps = append(data.Ipv4Vrrps, item)
			return true
		})
	} else {
		if len(data.Ipv4Vrrps) > 0 {
			data.Ipv4Vrrps = []CiscoVPNInterfaceIpv4Vrrps{}
		}
	}
	if value := res.Get(path + "ipv6-vrrp.vipValue"); len(value.Array()) > 0 {
		data.Ipv6Vrrps = make([]CiscoVPNInterfaceIpv6Vrrps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoVPNInterfaceIpv6Vrrps{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("grp-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.GroupId = types.Int64Null()

					cv := v.Get("grp-id.vipVariableName")
					item.GroupIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.GroupId = types.Int64Null()
					item.GroupIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("grp-id.vipValue")
					item.GroupId = types.Int64Value(cv.Int())
					item.GroupIdVariable = types.StringNull()
				}
			} else {
				item.GroupId = types.Int64Null()
				item.GroupIdVariable = types.StringNull()
			}
			if cValue := v.Get("priority.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Priority = types.Int64Null()

					cv := v.Get("priority.vipVariableName")
					item.PriorityVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Priority = types.Int64Null()
					item.PriorityVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("priority.vipValue")
					item.Priority = types.Int64Value(cv.Int())
					item.PriorityVariable = types.StringNull()
				}
			} else {
				item.Priority = types.Int64Null()
				item.PriorityVariable = types.StringNull()
			}
			if cValue := v.Get("timer.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Timer = types.Int64Null()

					cv := v.Get("timer.vipVariableName")
					item.TimerVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Timer = types.Int64Null()
					item.TimerVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("timer.vipValue")
					item.Timer = types.Int64Value(cv.Int())
					item.TimerVariable = types.StringNull()
				}
			} else {
				item.Timer = types.Int64Null()
				item.TimerVariable = types.StringNull()
			}
			if cValue := v.Get("track-omp.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TrackOmp = types.BoolNull()

					cv := v.Get("track-omp.vipVariableName")
					item.TrackOmpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TrackOmp = types.BoolNull()
					item.TrackOmpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("track-omp.vipValue")
					item.TrackOmp = types.BoolValue(cv.Bool())
					item.TrackOmpVariable = types.StringNull()
				}
			} else {
				item.TrackOmp = types.BoolNull()
				item.TrackOmpVariable = types.StringNull()
			}
			if cValue := v.Get("track-prefix-list.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TrackPrefixList = types.StringNull()

					cv := v.Get("track-prefix-list.vipVariableName")
					item.TrackPrefixListVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TrackPrefixList = types.StringNull()
					item.TrackPrefixListVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("track-prefix-list.vipValue")
					item.TrackPrefixList = types.StringValue(cv.String())
					item.TrackPrefixListVariable = types.StringNull()
				}
			} else {
				item.TrackPrefixList = types.StringNull()
				item.TrackPrefixListVariable = types.StringNull()
			}
			if cValue := v.Get("ipv6.vipValue"); len(cValue.Array()) > 0 {
				item.Ipv6Addresses = make([]CiscoVPNInterfaceIpv6VrrpsIpv6Addresses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoVPNInterfaceIpv6VrrpsIpv6Addresses{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("ipv6-link-local.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Ipv6LinkLocal = types.StringNull()

							ccv := cv.Get("ipv6-link-local.vipVariableName")
							cItem.Ipv6LinkLocalVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Ipv6LinkLocal = types.StringNull()
							cItem.Ipv6LinkLocalVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("ipv6-link-local.vipValue")
							cItem.Ipv6LinkLocal = types.StringValue(ccv.String())
							cItem.Ipv6LinkLocalVariable = types.StringNull()
						}
					} else {
						cItem.Ipv6LinkLocal = types.StringNull()
						cItem.Ipv6LinkLocalVariable = types.StringNull()
					}
					if ccValue := cv.Get("prefix.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Prefix = types.StringNull()

							ccv := cv.Get("prefix.vipVariableName")
							cItem.PrefixVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Prefix = types.StringNull()
							cItem.PrefixVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("prefix.vipValue")
							cItem.Prefix = types.StringValue(ccv.String())
							cItem.PrefixVariable = types.StringNull()
						}
					} else {
						cItem.Prefix = types.StringNull()
						cItem.PrefixVariable = types.StringNull()
					}
					item.Ipv6Addresses = append(item.Ipv6Addresses, cItem)
					return true
				})
			} else {
				if len(item.Ipv6Addresses) > 0 {
					item.Ipv6Addresses = []CiscoVPNInterfaceIpv6VrrpsIpv6Addresses{}
				}
			}
			data.Ipv6Vrrps = append(data.Ipv6Vrrps, item)
			return true
		})
	} else {
		if len(data.Ipv6Vrrps) > 0 {
			data.Ipv6Vrrps = []CiscoVPNInterfaceIpv6Vrrps{}
		}
	}
	if value := res.Get(path + "trustsec.propagate.sgt.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PropagateSgt = types.BoolNull()

		} else if value.String() == "ignore" {
			data.PropagateSgt = types.BoolNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "trustsec.propagate.sgt.vipValue")
			data.PropagateSgt = types.BoolValue(v.Bool())

		}
	} else {
		data.PropagateSgt = types.BoolNull()

	}
	if value := res.Get(path + "trustsec.static.sgt.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.StaticSgt = types.Int64Null()

			v := res.Get(path + "trustsec.static.sgt.vipVariableName")
			data.StaticSgtVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.StaticSgt = types.Int64Null()
			data.StaticSgtVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "trustsec.static.sgt.vipValue")
			data.StaticSgt = types.Int64Value(v.Int())
			data.StaticSgtVariable = types.StringNull()
		}
	} else {
		data.StaticSgt = types.Int64Null()
		data.StaticSgtVariable = types.StringNull()
	}
	if value := res.Get(path + "trustsec.static.trusted.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.StaticSgtTrusted = types.BoolNull()

		} else if value.String() == "ignore" {
			data.StaticSgtTrusted = types.BoolNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "trustsec.static.trusted.vipValue")
			data.StaticSgtTrusted = types.BoolValue(v.Bool())

		}
	} else {
		data.StaticSgtTrusted = types.BoolNull()

	}
	if value := res.Get(path + "trustsec.enable.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.EnableSgt = types.BoolNull()

		} else if value.String() == "ignore" {
			data.EnableSgt = types.BoolNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "trustsec.enable.vipValue")
			data.EnableSgt = types.BoolValue(v.Bool())

		}
	} else {
		data.EnableSgt = types.BoolNull()

	}
	if value := res.Get(path + "trustsec.enforcement.enable.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SgtEnforcement = types.BoolNull()

		} else if value.String() == "ignore" {
			data.SgtEnforcement = types.BoolNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "trustsec.enforcement.enable.vipValue")
			data.SgtEnforcement = types.BoolValue(v.Bool())

		}
	} else {
		data.SgtEnforcement = types.BoolNull()

	}
	if value := res.Get(path + "trustsec.enforcement.sgt.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SgtEnforcementSgt = types.Int64Null()

			v := res.Get(path + "trustsec.enforcement.sgt.vipVariableName")
			data.SgtEnforcementSgtVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SgtEnforcementSgt = types.Int64Null()
			data.SgtEnforcementSgtVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "trustsec.enforcement.sgt.vipValue")
			data.SgtEnforcementSgt = types.Int64Value(v.Int())
			data.SgtEnforcementSgtVariable = types.StringNull()
		}
	} else {
		data.SgtEnforcementSgt = types.Int64Null()
		data.SgtEnforcementSgtVariable = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoVPNInterface) hasChanges(ctx context.Context, state *CiscoVPNInterface) bool {
	hasChanges := false
	if !data.InterfaceName.Equal(state.InterfaceName) {
		hasChanges = true
	}
	if !data.InterfaceDescription.Equal(state.InterfaceDescription) {
		hasChanges = true
	}
	if !data.Poe.Equal(state.Poe) {
		hasChanges = true
	}
	if !data.Address.Equal(state.Address) {
		hasChanges = true
	}
	if len(data.Ipv4SecondaryAddresses) != len(state.Ipv4SecondaryAddresses) {
		hasChanges = true
	} else {
		for i := range data.Ipv4SecondaryAddresses {
			if !data.Ipv4SecondaryAddresses[i].Address.Equal(state.Ipv4SecondaryAddresses[i].Address) {
				hasChanges = true
			}
		}
	}
	if !data.Dhcp.Equal(state.Dhcp) {
		hasChanges = true
	}
	if !data.DhcpDistance.Equal(state.DhcpDistance) {
		hasChanges = true
	}
	if !data.Ipv6Address.Equal(state.Ipv6Address) {
		hasChanges = true
	}
	if !data.Dhcpv6.Equal(state.Dhcpv6) {
		hasChanges = true
	}
	if len(data.Ipv6SecondaryAddresses) != len(state.Ipv6SecondaryAddresses) {
		hasChanges = true
	} else {
		for i := range data.Ipv6SecondaryAddresses {
			if !data.Ipv6SecondaryAddresses[i].Address.Equal(state.Ipv6SecondaryAddresses[i].Address) {
				hasChanges = true
			}
		}
	}
	if len(data.Ipv6AccessLists) != len(state.Ipv6AccessLists) {
		hasChanges = true
	} else {
		for i := range data.Ipv6AccessLists {
			if !data.Ipv6AccessLists[i].Direction.Equal(state.Ipv6AccessLists[i].Direction) {
				hasChanges = true
			}
			if !data.Ipv6AccessLists[i].AclName.Equal(state.Ipv6AccessLists[i].AclName) {
				hasChanges = true
			}
		}
	}
	if !data.Ipv4DhcpHelper.Equal(state.Ipv4DhcpHelper) {
		hasChanges = true
	}
	if len(data.Ipv6DhcpHelpers) != len(state.Ipv6DhcpHelpers) {
		hasChanges = true
	} else {
		for i := range data.Ipv6DhcpHelpers {
			if !data.Ipv6DhcpHelpers[i].Address.Equal(state.Ipv6DhcpHelpers[i].Address) {
				hasChanges = true
			}
			if !data.Ipv6DhcpHelpers[i].VpnId.Equal(state.Ipv6DhcpHelpers[i].VpnId) {
				hasChanges = true
			}
		}
	}
	if !data.Tracker.Equal(state.Tracker) {
		hasChanges = true
	}
	if !data.AutoBandwidthDetect.Equal(state.AutoBandwidthDetect) {
		hasChanges = true
	}
	if !data.IperfServer.Equal(state.IperfServer) {
		hasChanges = true
	}
	if !data.Nat.Equal(state.Nat) {
		hasChanges = true
	}
	if !data.NatType.Equal(state.NatType) {
		hasChanges = true
	}
	if !data.UdpTimeout.Equal(state.UdpTimeout) {
		hasChanges = true
	}
	if !data.TcpTimeout.Equal(state.TcpTimeout) {
		hasChanges = true
	}
	if !data.NatPoolRangeStart.Equal(state.NatPoolRangeStart) {
		hasChanges = true
	}
	if !data.NatPoolRangeEnd.Equal(state.NatPoolRangeEnd) {
		hasChanges = true
	}
	if !data.NatOverload.Equal(state.NatOverload) {
		hasChanges = true
	}
	if !data.NatInsideSourceLoopbackInterface.Equal(state.NatInsideSourceLoopbackInterface) {
		hasChanges = true
	}
	if !data.NatPoolPrefixLength.Equal(state.NatPoolPrefixLength) {
		hasChanges = true
	}
	if !data.Ipv6Nat.Equal(state.Ipv6Nat) {
		hasChanges = true
	}
	if !data.Nat64Interface.Equal(state.Nat64Interface) {
		hasChanges = true
	}
	if !data.Nat66Interface.Equal(state.Nat66Interface) {
		hasChanges = true
	}
	if len(data.StaticNat66Entries) != len(state.StaticNat66Entries) {
		hasChanges = true
	} else {
		for i := range data.StaticNat66Entries {
			if !data.StaticNat66Entries[i].SourcePrefix.Equal(state.StaticNat66Entries[i].SourcePrefix) {
				hasChanges = true
			}
			if !data.StaticNat66Entries[i].TranslatedSourcePrefix.Equal(state.StaticNat66Entries[i].TranslatedSourcePrefix) {
				hasChanges = true
			}
			if !data.StaticNat66Entries[i].SourceVpnId.Equal(state.StaticNat66Entries[i].SourceVpnId) {
				hasChanges = true
			}
		}
	}
	if len(data.StaticNatEntries) != len(state.StaticNatEntries) {
		hasChanges = true
	} else {
		for i := range data.StaticNatEntries {
			if !data.StaticNatEntries[i].SourceIp.Equal(state.StaticNatEntries[i].SourceIp) {
				hasChanges = true
			}
			if !data.StaticNatEntries[i].TranslateIp.Equal(state.StaticNatEntries[i].TranslateIp) {
				hasChanges = true
			}
			if !data.StaticNatEntries[i].StaticNatDirection.Equal(state.StaticNatEntries[i].StaticNatDirection) {
				hasChanges = true
			}
			if !data.StaticNatEntries[i].SourceVpnId.Equal(state.StaticNatEntries[i].SourceVpnId) {
				hasChanges = true
			}
		}
	}
	if len(data.StaticPortForwardEntries) != len(state.StaticPortForwardEntries) {
		hasChanges = true
	} else {
		for i := range data.StaticPortForwardEntries {
			if !data.StaticPortForwardEntries[i].SourceIp.Equal(state.StaticPortForwardEntries[i].SourceIp) {
				hasChanges = true
			}
			if !data.StaticPortForwardEntries[i].TranslateIp.Equal(state.StaticPortForwardEntries[i].TranslateIp) {
				hasChanges = true
			}
			if !data.StaticPortForwardEntries[i].StaticNatDirection.Equal(state.StaticPortForwardEntries[i].StaticNatDirection) {
				hasChanges = true
			}
			if !data.StaticPortForwardEntries[i].SourcePort.Equal(state.StaticPortForwardEntries[i].SourcePort) {
				hasChanges = true
			}
			if !data.StaticPortForwardEntries[i].TranslatePort.Equal(state.StaticPortForwardEntries[i].TranslatePort) {
				hasChanges = true
			}
			if !data.StaticPortForwardEntries[i].Protocol.Equal(state.StaticPortForwardEntries[i].Protocol) {
				hasChanges = true
			}
			if !data.StaticPortForwardEntries[i].SourceVpnId.Equal(state.StaticPortForwardEntries[i].SourceVpnId) {
				hasChanges = true
			}
		}
	}
	if !data.EnableCoreRegion.Equal(state.EnableCoreRegion) {
		hasChanges = true
	}
	if !data.CoreRegion.Equal(state.CoreRegion) {
		hasChanges = true
	}
	if !data.SecondaryRegion.Equal(state.SecondaryRegion) {
		hasChanges = true
	}
	if len(data.TunnelInterfaceEncapsulations) != len(state.TunnelInterfaceEncapsulations) {
		hasChanges = true
	} else {
		for i := range data.TunnelInterfaceEncapsulations {
			if !data.TunnelInterfaceEncapsulations[i].Encapsulation.Equal(state.TunnelInterfaceEncapsulations[i].Encapsulation) {
				hasChanges = true
			}
			if !data.TunnelInterfaceEncapsulations[i].Preference.Equal(state.TunnelInterfaceEncapsulations[i].Preference) {
				hasChanges = true
			}
			if !data.TunnelInterfaceEncapsulations[i].Weight.Equal(state.TunnelInterfaceEncapsulations[i].Weight) {
				hasChanges = true
			}
		}
	}
	if !data.TunnelInterfaceBorder.Equal(state.TunnelInterfaceBorder) {
		hasChanges = true
	}
	if !data.TunnelQosMode.Equal(state.TunnelQosMode) {
		hasChanges = true
	}
	if !data.TunnelBandwidth.Equal(state.TunnelBandwidth) {
		hasChanges = true
	}
	if !data.TunnelInterfaceGroups.Equal(state.TunnelInterfaceGroups) {
		hasChanges = true
	}
	if !data.TunnelInterfaceColor.Equal(state.TunnelInterfaceColor) {
		hasChanges = true
	}
	if !data.TunnelInterfaceMaxControlConnections.Equal(state.TunnelInterfaceMaxControlConnections) {
		hasChanges = true
	}
	if !data.TunnelInterfaceControlConnections.Equal(state.TunnelInterfaceControlConnections) {
		hasChanges = true
	}
	if !data.TunnelInterfaceVbondAsStunServer.Equal(state.TunnelInterfaceVbondAsStunServer) {
		hasChanges = true
	}
	if !data.TunnelInterfaceExcludeControllerGroupList.Equal(state.TunnelInterfaceExcludeControllerGroupList) {
		hasChanges = true
	}
	if !data.TunnelInterfaceVmanageConnectionPreference.Equal(state.TunnelInterfaceVmanageConnectionPreference) {
		hasChanges = true
	}
	if !data.TunnelInterfacePortHop.Equal(state.TunnelInterfacePortHop) {
		hasChanges = true
	}
	if !data.TunnelInterfaceColorRestrict.Equal(state.TunnelInterfaceColorRestrict) {
		hasChanges = true
	}
	if !data.TunnelInterfaceGreTunnelDestinationIp.Equal(state.TunnelInterfaceGreTunnelDestinationIp) {
		hasChanges = true
	}
	if !data.TunnelInterfaceCarrier.Equal(state.TunnelInterfaceCarrier) {
		hasChanges = true
	}
	if !data.TunnelInterfaceNatRefreshInterval.Equal(state.TunnelInterfaceNatRefreshInterval) {
		hasChanges = true
	}
	if !data.TunnelInterfaceHelloInterval.Equal(state.TunnelInterfaceHelloInterval) {
		hasChanges = true
	}
	if !data.TunnelInterfaceHelloTolerance.Equal(state.TunnelInterfaceHelloTolerance) {
		hasChanges = true
	}
	if !data.TunnelInterfaceBindLoopbackTunnel.Equal(state.TunnelInterfaceBindLoopbackTunnel) {
		hasChanges = true
	}
	if !data.TunnelInterfaceLastResortCircuit.Equal(state.TunnelInterfaceLastResortCircuit) {
		hasChanges = true
	}
	if !data.TunnelInterfaceLowBandwidthLink.Equal(state.TunnelInterfaceLowBandwidthLink) {
		hasChanges = true
	}
	if !data.TunnelInterfaceTunnelTcpMss.Equal(state.TunnelInterfaceTunnelTcpMss) {
		hasChanges = true
	}
	if !data.TunnelInterfaceClearDontFragment.Equal(state.TunnelInterfaceClearDontFragment) {
		hasChanges = true
	}
	if !data.TunnelInterfacePropagateSgt.Equal(state.TunnelInterfacePropagateSgt) {
		hasChanges = true
	}
	if !data.TunnelInterfaceNetworkBroadcast.Equal(state.TunnelInterfaceNetworkBroadcast) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowAll.Equal(state.TunnelInterfaceAllowAll) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowBgp.Equal(state.TunnelInterfaceAllowBgp) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowDhcp.Equal(state.TunnelInterfaceAllowDhcp) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowDns.Equal(state.TunnelInterfaceAllowDns) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowIcmp.Equal(state.TunnelInterfaceAllowIcmp) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowSsh.Equal(state.TunnelInterfaceAllowSsh) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowNetconf.Equal(state.TunnelInterfaceAllowNetconf) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowNtp.Equal(state.TunnelInterfaceAllowNtp) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowOspf.Equal(state.TunnelInterfaceAllowOspf) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowStun.Equal(state.TunnelInterfaceAllowStun) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowSnmp.Equal(state.TunnelInterfaceAllowSnmp) {
		hasChanges = true
	}
	if !data.TunnelInterfaceAllowHttps.Equal(state.TunnelInterfaceAllowHttps) {
		hasChanges = true
	}
	if !data.MediaType.Equal(state.MediaType) {
		hasChanges = true
	}
	if !data.InterfaceMtu.Equal(state.InterfaceMtu) {
		hasChanges = true
	}
	if !data.IpMtu.Equal(state.IpMtu) {
		hasChanges = true
	}
	if !data.TcpMssAdjust.Equal(state.TcpMssAdjust) {
		hasChanges = true
	}
	if !data.TlocExtension.Equal(state.TlocExtension) {
		hasChanges = true
	}
	if !data.LoadInterval.Equal(state.LoadInterval) {
		hasChanges = true
	}
	if !data.GreTunnelSourceIp.Equal(state.GreTunnelSourceIp) {
		hasChanges = true
	}
	if !data.GreTunnelXconnect.Equal(state.GreTunnelXconnect) {
		hasChanges = true
	}
	if !data.MacAddress.Equal(state.MacAddress) {
		hasChanges = true
	}
	if !data.Speed.Equal(state.Speed) {
		hasChanges = true
	}
	if !data.Duplex.Equal(state.Duplex) {
		hasChanges = true
	}
	if !data.Shutdown.Equal(state.Shutdown) {
		hasChanges = true
	}
	if !data.ArpTimeout.Equal(state.ArpTimeout) {
		hasChanges = true
	}
	if !data.Autonegotiate.Equal(state.Autonegotiate) {
		hasChanges = true
	}
	if !data.IpDirectedBroadcast.Equal(state.IpDirectedBroadcast) {
		hasChanges = true
	}
	if !data.IcmpRedirectDisable.Equal(state.IcmpRedirectDisable) {
		hasChanges = true
	}
	if !data.QosAdaptivePeriod.Equal(state.QosAdaptivePeriod) {
		hasChanges = true
	}
	if !data.QosAdaptiveBandwidthDownstream.Equal(state.QosAdaptiveBandwidthDownstream) {
		hasChanges = true
	}
	if !data.QosAdaptiveMinDownstream.Equal(state.QosAdaptiveMinDownstream) {
		hasChanges = true
	}
	if !data.QosAdaptiveMaxDownstream.Equal(state.QosAdaptiveMaxDownstream) {
		hasChanges = true
	}
	if !data.QosAdaptiveBandwidthUpstream.Equal(state.QosAdaptiveBandwidthUpstream) {
		hasChanges = true
	}
	if !data.QosAdaptiveMinUpstream.Equal(state.QosAdaptiveMinUpstream) {
		hasChanges = true
	}
	if !data.QosAdaptiveMaxUpstream.Equal(state.QosAdaptiveMaxUpstream) {
		hasChanges = true
	}
	if !data.ShapingRate.Equal(state.ShapingRate) {
		hasChanges = true
	}
	if !data.QosMap.Equal(state.QosMap) {
		hasChanges = true
	}
	if !data.QosMapVpn.Equal(state.QosMapVpn) {
		hasChanges = true
	}
	if !data.BandwidthUpstream.Equal(state.BandwidthUpstream) {
		hasChanges = true
	}
	if !data.BandwidthDownstream.Equal(state.BandwidthDownstream) {
		hasChanges = true
	}
	if !data.BlockNonSourceIp.Equal(state.BlockNonSourceIp) {
		hasChanges = true
	}
	if !data.RewriteRuleName.Equal(state.RewriteRuleName) {
		hasChanges = true
	}
	if len(data.AccessLists) != len(state.AccessLists) {
		hasChanges = true
	} else {
		for i := range data.AccessLists {
			if !data.AccessLists[i].Direction.Equal(state.AccessLists[i].Direction) {
				hasChanges = true
			}
			if !data.AccessLists[i].AclName.Equal(state.AccessLists[i].AclName) {
				hasChanges = true
			}
		}
	}
	if len(data.StaticArps) != len(state.StaticArps) {
		hasChanges = true
	} else {
		for i := range data.StaticArps {
			if !data.StaticArps[i].IpAddress.Equal(state.StaticArps[i].IpAddress) {
				hasChanges = true
			}
			if !data.StaticArps[i].Mac.Equal(state.StaticArps[i].Mac) {
				hasChanges = true
			}
		}
	}
	if len(data.Ipv4Vrrps) != len(state.Ipv4Vrrps) {
		hasChanges = true
	} else {
		for i := range data.Ipv4Vrrps {
			if !data.Ipv4Vrrps[i].GroupId.Equal(state.Ipv4Vrrps[i].GroupId) {
				hasChanges = true
			}
			if !data.Ipv4Vrrps[i].Priority.Equal(state.Ipv4Vrrps[i].Priority) {
				hasChanges = true
			}
			if !data.Ipv4Vrrps[i].Timer.Equal(state.Ipv4Vrrps[i].Timer) {
				hasChanges = true
			}
			if !data.Ipv4Vrrps[i].TrackOmp.Equal(state.Ipv4Vrrps[i].TrackOmp) {
				hasChanges = true
			}
			if !data.Ipv4Vrrps[i].TrackPrefixList.Equal(state.Ipv4Vrrps[i].TrackPrefixList) {
				hasChanges = true
			}
			if !data.Ipv4Vrrps[i].IpAddress.Equal(state.Ipv4Vrrps[i].IpAddress) {
				hasChanges = true
			}
			if len(data.Ipv4Vrrps[i].Ipv4SecondaryAddresses) != len(state.Ipv4Vrrps[i].Ipv4SecondaryAddresses) {
				hasChanges = true
			} else {
				for ii := range data.Ipv4Vrrps[i].Ipv4SecondaryAddresses {
					if !data.Ipv4Vrrps[i].Ipv4SecondaryAddresses[ii].IpAddress.Equal(state.Ipv4Vrrps[i].Ipv4SecondaryAddresses[ii].IpAddress) {
						hasChanges = true
					}
				}
			}
			if !data.Ipv4Vrrps[i].TlocPreferenceChange.Equal(state.Ipv4Vrrps[i].TlocPreferenceChange) {
				hasChanges = true
			}
			if !data.Ipv4Vrrps[i].TlocPreferenceChangeValue.Equal(state.Ipv4Vrrps[i].TlocPreferenceChangeValue) {
				hasChanges = true
			}
			if len(data.Ipv4Vrrps[i].TrackingObjects) != len(state.Ipv4Vrrps[i].TrackingObjects) {
				hasChanges = true
			} else {
				for ii := range data.Ipv4Vrrps[i].TrackingObjects {
					if !data.Ipv4Vrrps[i].TrackingObjects[ii].TrackerId.Equal(state.Ipv4Vrrps[i].TrackingObjects[ii].TrackerId) {
						hasChanges = true
					}
					if !data.Ipv4Vrrps[i].TrackingObjects[ii].TrackAction.Equal(state.Ipv4Vrrps[i].TrackingObjects[ii].TrackAction) {
						hasChanges = true
					}
					if !data.Ipv4Vrrps[i].TrackingObjects[ii].DecrementValue.Equal(state.Ipv4Vrrps[i].TrackingObjects[ii].DecrementValue) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.Ipv6Vrrps) != len(state.Ipv6Vrrps) {
		hasChanges = true
	} else {
		for i := range data.Ipv6Vrrps {
			if !data.Ipv6Vrrps[i].GroupId.Equal(state.Ipv6Vrrps[i].GroupId) {
				hasChanges = true
			}
			if !data.Ipv6Vrrps[i].Priority.Equal(state.Ipv6Vrrps[i].Priority) {
				hasChanges = true
			}
			if !data.Ipv6Vrrps[i].Timer.Equal(state.Ipv6Vrrps[i].Timer) {
				hasChanges = true
			}
			if !data.Ipv6Vrrps[i].TrackOmp.Equal(state.Ipv6Vrrps[i].TrackOmp) {
				hasChanges = true
			}
			if !data.Ipv6Vrrps[i].TrackPrefixList.Equal(state.Ipv6Vrrps[i].TrackPrefixList) {
				hasChanges = true
			}
			if len(data.Ipv6Vrrps[i].Ipv6Addresses) != len(state.Ipv6Vrrps[i].Ipv6Addresses) {
				hasChanges = true
			} else {
				for ii := range data.Ipv6Vrrps[i].Ipv6Addresses {
					if !data.Ipv6Vrrps[i].Ipv6Addresses[ii].Ipv6LinkLocal.Equal(state.Ipv6Vrrps[i].Ipv6Addresses[ii].Ipv6LinkLocal) {
						hasChanges = true
					}
					if !data.Ipv6Vrrps[i].Ipv6Addresses[ii].Prefix.Equal(state.Ipv6Vrrps[i].Ipv6Addresses[ii].Prefix) {
						hasChanges = true
					}
				}
			}
		}
	}
	if !data.PropagateSgt.Equal(state.PropagateSgt) {
		hasChanges = true
	}
	if !data.StaticSgt.Equal(state.StaticSgt) {
		hasChanges = true
	}
	if !data.StaticSgtTrusted.Equal(state.StaticSgtTrusted) {
		hasChanges = true
	}
	if !data.EnableSgt.Equal(state.EnableSgt) {
		hasChanges = true
	}
	if !data.SgtEnforcement.Equal(state.SgtEnforcement) {
		hasChanges = true
	}
	if !data.SgtEnforcementSgt.Equal(state.SgtEnforcementSgt) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
