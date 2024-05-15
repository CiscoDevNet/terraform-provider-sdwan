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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type TransportWANVPNInterfaceEthernet struct {
	Id                                                 types.String                                                    `tfsdk:"id"`
	Version                                            types.Int64                                                     `tfsdk:"version"`
	Name                                               types.String                                                    `tfsdk:"name"`
	Description                                        types.String                                                    `tfsdk:"description"`
	FeatureProfileId                                   types.String                                                    `tfsdk:"feature_profile_id"`
	TransportWanVpnProfileParcelId                     types.String                                                    `tfsdk:"transport_wan_vpn_profile_parcel_id"`
	Shutdown                                           types.Bool                                                      `tfsdk:"shutdown"`
	ShutdownVariable                                   types.String                                                    `tfsdk:"shutdown_variable"`
	InterfaceName                                      types.String                                                    `tfsdk:"interface_name"`
	InterfaceNameVariable                              types.String                                                    `tfsdk:"interface_name_variable"`
	ConfigDescription                                  types.String                                                    `tfsdk:"config_description"`
	ConfigDescriptionVariable                          types.String                                                    `tfsdk:"config_description_variable"`
	Ipv4DhcpDistance                                   types.Int64                                                     `tfsdk:"ipv4_dhcp_distance"`
	Ipv4DhcpDistanceVariable                           types.String                                                    `tfsdk:"ipv4_dhcp_distance_variable"`
	Ipv4Address                                        types.String                                                    `tfsdk:"ipv4_address"`
	Ipv4AddressVariable                                types.String                                                    `tfsdk:"ipv4_address_variable"`
	Ipv4SubnetMask                                     types.String                                                    `tfsdk:"ipv4_subnet_mask"`
	Ipv4SubnetMaskVariable                             types.String                                                    `tfsdk:"ipv4_subnet_mask_variable"`
	Ipv4SecondaryAddresses                             []TransportWANVPNInterfaceEthernetIpv4SecondaryAddresses        `tfsdk:"ipv4_secondary_addresses"`
	Ipv4DhcpHelper                                     types.Set                                                       `tfsdk:"ipv4_dhcp_helper"`
	Ipv4DhcpHelperVariable                             types.String                                                    `tfsdk:"ipv4_dhcp_helper_variable"`
	EnableDhcpv6                                       types.Bool                                                      `tfsdk:"enable_dhcpv6"`
	Ipv6DhcpSecondaryAddress                           []TransportWANVPNInterfaceEthernetIpv6DhcpSecondaryAddress      `tfsdk:"ipv6__dhcp_secondary_address"`
	Ipv6Address                                        types.String                                                    `tfsdk:"ipv6_address"`
	Ipv6AddressVariable                                types.String                                                    `tfsdk:"ipv6_address_variable"`
	Ipv6SecondaryAddress                               []TransportWANVPNInterfaceEthernetIpv6SecondaryAddress          `tfsdk:"ipv6_secondary_address"`
	IperfServer                                        types.String                                                    `tfsdk:"iperf_server"`
	IperfServerVariable                                types.String                                                    `tfsdk:"iperf_server_variable"`
	BlockNonSourceIp                                   types.Bool                                                      `tfsdk:"block_non_source_ip"`
	BlockNonSourceIpVariable                           types.String                                                    `tfsdk:"block_non_source_ip_variable"`
	ServiceProvider                                    types.String                                                    `tfsdk:"service_provider"`
	ServiceProviderVariable                            types.String                                                    `tfsdk:"service_provider_variable"`
	BandwidthUpstream                                  types.Int64                                                     `tfsdk:"bandwidth_upstream"`
	BandwidthUpstreamVariable                          types.String                                                    `tfsdk:"bandwidth_upstream_variable"`
	BandwidthDownstream                                types.Int64                                                     `tfsdk:"bandwidth_downstream"`
	BandwidthDownstreamVariable                        types.String                                                    `tfsdk:"bandwidth_downstream_variable"`
	AutoDetectBandwidth                                types.Bool                                                      `tfsdk:"auto_detect_bandwidth"`
	AutoDetectBandwidthVariable                        types.String                                                    `tfsdk:"auto_detect_bandwidth_variable"`
	TunnelInterface                                    types.Bool                                                      `tfsdk:"tunnel_interface"`
	PerTunnelQos                                       types.Bool                                                      `tfsdk:"per_tunnel_qos"`
	PerTunnelQosVariable                               types.String                                                    `tfsdk:"per_tunnel_qos_variable"`
	TunnelQosMode                                      types.String                                                    `tfsdk:"tunnel_qos_mode"`
	TunnelQosModeVariable                              types.String                                                    `tfsdk:"tunnel_qos_mode_variable"`
	TunnelBandwidthPercent                             types.Int64                                                     `tfsdk:"tunnel_bandwidth_percent"`
	TunnelBandwidthPercentVariable                     types.String                                                    `tfsdk:"tunnel_bandwidth_percent_variable"`
	TunnelInterfaceBindLoopbackTunnel                  types.String                                                    `tfsdk:"tunnel_interface_bind_loopback_tunnel"`
	TunnelInterfaceBindLoopbackTunnelVariable          types.String                                                    `tfsdk:"tunnel_interface_bind_loopback_tunnel_variable"`
	TunnelInterfaceCarrier                             types.String                                                    `tfsdk:"tunnel_interface_carrier"`
	TunnelInterfaceCarrierVariable                     types.String                                                    `tfsdk:"tunnel_interface_carrier_variable"`
	TunnelInterfaceColor                               types.String                                                    `tfsdk:"tunnel_interface_color"`
	TunnelInterfaceColorVariable                       types.String                                                    `tfsdk:"tunnel_interface_color_variable"`
	TunnelInterfaceHelloInterval                       types.Int64                                                     `tfsdk:"tunnel_interface_hello_interval"`
	TunnelInterfaceHelloIntervalVariable               types.String                                                    `tfsdk:"tunnel_interface_hello_interval_variable"`
	TunnelInterfaceHelloTolerance                      types.Int64                                                     `tfsdk:"tunnel_interface_hello_tolerance"`
	TunnelInterfaceHelloToleranceVariable              types.String                                                    `tfsdk:"tunnel_interface_hello_tolerance_variable"`
	TunnelInterfaceLastResortCircuit                   types.Bool                                                      `tfsdk:"tunnel_interface_last_resort_circuit"`
	TunnelInterfaceLastResortCircuitVariable           types.String                                                    `tfsdk:"tunnel_interface_last_resort_circuit_variable"`
	TunnelInterfaceGreTunnelDestinationIp              types.String                                                    `tfsdk:"tunnel_interface_gre_tunnel_destination_ip"`
	TunnelInterfaceGreTunnelDestinationIpVariable      types.String                                                    `tfsdk:"tunnel_interface_gre_tunnel_destination_ip_variable"`
	TunnelInterfaceColorRestrict                       types.Bool                                                      `tfsdk:"tunnel_interface_color_restrict"`
	TunnelInterfaceColorRestrictVariable               types.String                                                    `tfsdk:"tunnel_interface_color_restrict_variable"`
	TunnelInterfaceGroups                              types.Int64                                                     `tfsdk:"tunnel_interface_groups"`
	TunnelInterfaceGroupsVariable                      types.String                                                    `tfsdk:"tunnel_interface_groups_variable"`
	TunnelInterfaceBorder                              types.Bool                                                      `tfsdk:"tunnel_interface_border"`
	TunnelInterfaceBorderVariable                      types.String                                                    `tfsdk:"tunnel_interface_border_variable"`
	TunnelInterfaceMaxControlConnections               types.Int64                                                     `tfsdk:"tunnel_interface_max_control_connections"`
	TunnelInterfaceMaxControlConnectionsVariable       types.String                                                    `tfsdk:"tunnel_interface_max_control_connections_variable"`
	TunnelInterfaceNatRefreshInterval                  types.Int64                                                     `tfsdk:"tunnel_interface_nat_refresh_interval"`
	TunnelInterfaceNatRefreshIntervalVariable          types.String                                                    `tfsdk:"tunnel_interface_nat_refresh_interval_variable"`
	TunnelInterfaceVbondAsStunServer                   types.Bool                                                      `tfsdk:"tunnel_interface_vbond_as_stun_server"`
	TunnelInterfaceVbondAsStunServerVariable           types.String                                                    `tfsdk:"tunnel_interface_vbond_as_stun_server_variable"`
	TunnelInterfaceExcludeControllerGroupList          types.Set                                                       `tfsdk:"tunnel_interface_exclude_controller_group_list"`
	TunnelInterfaceExcludeControllerGroupListVariable  types.String                                                    `tfsdk:"tunnel_interface_exclude_controller_group_list_variable"`
	TunnelInterfaceVmanageConnectionPreference         types.Int64                                                     `tfsdk:"tunnel_interface_vmanage_connection_preference"`
	TunnelInterfaceVmanageConnectionPreferenceVariable types.String                                                    `tfsdk:"tunnel_interface_vmanage_connection_preference_variable"`
	TunnelInterfacePortHop                             types.Bool                                                      `tfsdk:"tunnel_interface_port_hop"`
	TunnelInterfacePortHopVariable                     types.String                                                    `tfsdk:"tunnel_interface_port_hop_variable"`
	TunnelInterfaceLowBandwidthLink                    types.Bool                                                      `tfsdk:"tunnel_interface_low_bandwidth_link"`
	TunnelInterfaceLowBandwidthLinkVariable            types.String                                                    `tfsdk:"tunnel_interface_low_bandwidth_link_variable"`
	TunnelInterfaceTunnelTcpMss                        types.Int64                                                     `tfsdk:"tunnel_interface_tunnel_tcp_mss"`
	TunnelInterfaceTunnelTcpMssVariable                types.String                                                    `tfsdk:"tunnel_interface_tunnel_tcp_mss_variable"`
	TunnelInterfaceClearDontFragment                   types.Bool                                                      `tfsdk:"tunnel_interface_clear_dont_fragment"`
	TunnelInterfaceClearDontFragmentVariable           types.String                                                    `tfsdk:"tunnel_interface_clear_dont_fragment_variable"`
	TunnelInterfaceCtsSgtPropagation                   types.Bool                                                      `tfsdk:"tunnel_interface_cts_sgt_propagation"`
	TunnelInterfaceCtsSgtPropagationVariable           types.String                                                    `tfsdk:"tunnel_interface_cts_sgt_propagation_variable"`
	TunnelInterfaceNetworkBroadcast                    types.Bool                                                      `tfsdk:"tunnel_interface_network_broadcast"`
	TunnelInterfaceNetworkBroadcastVariable            types.String                                                    `tfsdk:"tunnel_interface_network_broadcast_variable"`
	TunnelInterfaceAllowAll                            types.Bool                                                      `tfsdk:"tunnel_interface_allow_all"`
	TunnelInterfaceAllowAllVariable                    types.String                                                    `tfsdk:"tunnel_interface_allow_all_variable"`
	TunnelInterfaceAllowBgp                            types.Bool                                                      `tfsdk:"tunnel_interface_allow_bgp"`
	TunnelInterfaceAllowBgpVariable                    types.String                                                    `tfsdk:"tunnel_interface_allow_bgp_variable"`
	TunnelInterfaceAllowDhcp                           types.Bool                                                      `tfsdk:"tunnel_interface_allow_dhcp"`
	TunnelInterfaceAllowDhcpVariable                   types.String                                                    `tfsdk:"tunnel_interface_allow_dhcp_variable"`
	TunnelInterfaceAllowNtp                            types.Bool                                                      `tfsdk:"tunnel_interface_allow_ntp"`
	TunnelInterfaceAllowNtpVariable                    types.String                                                    `tfsdk:"tunnel_interface_allow_ntp_variable"`
	TunnelInterfaceAllowSsh                            types.Bool                                                      `tfsdk:"tunnel_interface_allow_ssh"`
	TunnelInterfaceAllowSshVariable                    types.String                                                    `tfsdk:"tunnel_interface_allow_ssh_variable"`
	TunnelInterfaceAllowDbs                            types.Bool                                                      `tfsdk:"tunnel_interface_allow_dbs"`
	TunnelInterfaceAllowDbsVariable                    types.String                                                    `tfsdk:"tunnel_interface_allow_dbs_variable"`
	TunnelInterfaceAllowIcmp                           types.Bool                                                      `tfsdk:"tunnel_interface_allow_icmp"`
	TunnelInterfaceAllowIcmpVariable                   types.String                                                    `tfsdk:"tunnel_interface_allow_icmp_variable"`
	TunnelInterfaceAllowHttps                          types.Bool                                                      `tfsdk:"tunnel_interface_allow_https"`
	TunnelInterfaceAllowHttpsVariable                  types.String                                                    `tfsdk:"tunnel_interface_allow_https_variable"`
	TunnelInterfaceAllowOspf                           types.Bool                                                      `tfsdk:"tunnel_interface_allow_ospf"`
	TunnelInterfaceAllowOspfVariable                   types.String                                                    `tfsdk:"tunnel_interface_allow_ospf_variable"`
	TunnelInterfaceAllowStun                           types.Bool                                                      `tfsdk:"tunnel_interface_allow_stun"`
	TunnelInterfaceAllowStunVariable                   types.String                                                    `tfsdk:"tunnel_interface_allow_stun_variable"`
	TunnelInterfaceAllowSnmp                           types.Bool                                                      `tfsdk:"tunnel_interface_allow_snmp"`
	TunnelInterfaceAllowSnmpVariable                   types.String                                                    `tfsdk:"tunnel_interface_allow_snmp_variable"`
	TunnelInterfaceAllowNetconf                        types.Bool                                                      `tfsdk:"tunnel_interface_allow_netconf"`
	TunnelInterfaceAllowNetconfVariable                types.String                                                    `tfsdk:"tunnel_interface_allow_netconf_variable"`
	TunnelInterfaceAllowBfd                            types.Bool                                                      `tfsdk:"tunnel_interface_allow_bfd"`
	TunnelInterfaceAllowBfdVariable                    types.String                                                    `tfsdk:"tunnel_interface_allow_bfd_variable"`
	TunnelInterfaceEncapsulations                      []TransportWANVPNInterfaceEthernetTunnelInterfaceEncapsulations `tfsdk:"tunnel_interface_encapsulations"`
	NatIpv4                                            types.Bool                                                      `tfsdk:"nat_ipv4"`
	NatIpv4Variable                                    types.String                                                    `tfsdk:"nat_ipv4_variable"`
	NatType                                            types.String                                                    `tfsdk:"nat_type"`
	NatTypeVariable                                    types.String                                                    `tfsdk:"nat_type_variable"`
	NatRangeStart                                      types.String                                                    `tfsdk:"nat_range_start"`
	NatRangeStartVariable                              types.String                                                    `tfsdk:"nat_range_start_variable"`
	NatRangeEnd                                        types.String                                                    `tfsdk:"nat_range_end"`
	NatRangeEndVariable                                types.String                                                    `tfsdk:"nat_range_end_variable"`
	NatPrefixLength                                    types.Int64                                                     `tfsdk:"nat_prefix_length"`
	NatPrefixLengthVariable                            types.String                                                    `tfsdk:"nat_prefix_length_variable"`
	NatOverload                                        types.Bool                                                      `tfsdk:"nat_overload"`
	NatOverloadVariable                                types.String                                                    `tfsdk:"nat_overload_variable"`
	NatLoopback                                        types.String                                                    `tfsdk:"nat_loopback"`
	NatLoopbackVariable                                types.String                                                    `tfsdk:"nat_loopback_variable"`
	NatUdpTimeout                                      types.Int64                                                     `tfsdk:"nat_udp_timeout"`
	NatUdpTimeoutVariable                              types.String                                                    `tfsdk:"nat_udp_timeout_variable"`
	NatTcpTimeout                                      types.Int64                                                     `tfsdk:"nat_tcp_timeout"`
	NatTcpTimeoutVariable                              types.String                                                    `tfsdk:"nat_tcp_timeout_variable"`
	NewStaticNats                                      []TransportWANVPNInterfaceEthernetNewStaticNats                 `tfsdk:"new_static_nats"`
	NatIpv6                                            types.Bool                                                      `tfsdk:"nat_ipv6"`
	NatIpv6Variable                                    types.String                                                    `tfsdk:"nat_ipv6_variable"`
	Nat64                                              types.Bool                                                      `tfsdk:"nat64"`
	Nat66                                              types.Bool                                                      `tfsdk:"nat66"`
	StaticNat66                                        []TransportWANVPNInterfaceEthernetStaticNat66                   `tfsdk:"static_nat66"`
	AdaptiveQos                                        types.Bool                                                      `tfsdk:"adaptive_qos"`
	QosAdaptivePeriod                                  types.Int64                                                     `tfsdk:"qos_adaptive_period"`
	QosAdaptivePeriodVariable                          types.String                                                    `tfsdk:"qos_adaptive_period_variable"`
	QosAdaptiveBandwidthUpstream                       types.Bool                                                      `tfsdk:"qos_adaptive_bandwidth_upstream"`
	QosAdaptiveMinUpstream                             types.Int64                                                     `tfsdk:"qos_adaptive_min_upstream"`
	QosAdaptiveMinUpstreamVariable                     types.String                                                    `tfsdk:"qos_adaptive_min_upstream_variable"`
	QosAdaptiveMaxUpstream                             types.Int64                                                     `tfsdk:"qos_adaptive_max_upstream"`
	QosAdaptiveMaxUpstreamVariable                     types.String                                                    `tfsdk:"qos_adaptive_max_upstream_variable"`
	QosAdaptiveDefaultUpstream                         types.Int64                                                     `tfsdk:"qos_adaptive_default_upstream"`
	QosAdaptiveDefaultUpstreamVariable                 types.String                                                    `tfsdk:"qos_adaptive_default_upstream_variable"`
	QosAdaptiveBandwidthDownstream                     types.Bool                                                      `tfsdk:"qos_adaptive_bandwidth_downstream"`
	QosAdaptiveMinDownstream                           types.Int64                                                     `tfsdk:"qos_adaptive_min_downstream"`
	QosAdaptiveMinDownstreamVariable                   types.String                                                    `tfsdk:"qos_adaptive_min_downstream_variable"`
	QosAdaptiveMaxDownstream                           types.Int64                                                     `tfsdk:"qos_adaptive_max_downstream"`
	QosAdaptiveMaxDownstreamVariable                   types.String                                                    `tfsdk:"qos_adaptive_max_downstream_variable"`
	QosAdaptiveDefaultDownstream                       types.Int64                                                     `tfsdk:"qos_adaptive_default_downstream"`
	QosAdaptiveDefaultDownstreamVariable               types.String                                                    `tfsdk:"qos_adaptive_default_downstream_variable"`
	QosShapingRate                                     types.Int64                                                     `tfsdk:"qos_shaping_rate"`
	QosShapingRateVariable                             types.String                                                    `tfsdk:"qos_shaping_rate_variable"`
	Arps                                               []TransportWANVPNInterfaceEthernetArps                          `tfsdk:"arps"`
	AdvancedIcmpRedirectDisable                        types.Bool                                                      `tfsdk:"advanced_icmp_redirect_disable"`
	AdvancedIcmpRedirectDisableVariable                types.String                                                    `tfsdk:"advanced_icmp_redirect_disable_variable"`
	AdvancedDuplex                                     types.String                                                    `tfsdk:"advanced_duplex"`
	AdvancedDuplexVariable                             types.String                                                    `tfsdk:"advanced_duplex_variable"`
	AdvancedMacAddress                                 types.String                                                    `tfsdk:"advanced_mac_address"`
	AdvancedMacAddressVariable                         types.String                                                    `tfsdk:"advanced_mac_address_variable"`
	AdvancedIpMtu                                      types.Int64                                                     `tfsdk:"advanced_ip_mtu"`
	AdvancedIpMtuVariable                              types.String                                                    `tfsdk:"advanced_ip_mtu_variable"`
	AdvancedInterfaceMtu                               types.Int64                                                     `tfsdk:"advanced_interface_mtu"`
	AdvancedInterfaceMtuVariable                       types.String                                                    `tfsdk:"advanced_interface_mtu_variable"`
	AdvancedTcpMss                                     types.Int64                                                     `tfsdk:"advanced_tcp_mss"`
	AdvancedTcpMssVariable                             types.String                                                    `tfsdk:"advanced_tcp_mss_variable"`
	AdvancedSpeed                                      types.String                                                    `tfsdk:"advanced_speed"`
	AdvancedSpeedVariable                              types.String                                                    `tfsdk:"advanced_speed_variable"`
	AdvancedArpTimeout                                 types.Int64                                                     `tfsdk:"advanced_arp_timeout"`
	AdvancedArpTimeoutVariable                         types.String                                                    `tfsdk:"advanced_arp_timeout_variable"`
	AdvancedAutonegotiate                              types.Bool                                                      `tfsdk:"advanced_autonegotiate"`
	AdvancedAutonegotiateVariable                      types.String                                                    `tfsdk:"advanced_autonegotiate_variable"`
	AdvancedMediaType                                  types.String                                                    `tfsdk:"advanced_media_type"`
	AdvancedMediaTypeVariable                          types.String                                                    `tfsdk:"advanced_media_type_variable"`
	AdvancedTlocExtension                              types.String                                                    `tfsdk:"advanced_tloc_extension"`
	AdvancedTlocExtensionVariable                      types.String                                                    `tfsdk:"advanced_tloc_extension_variable"`
	AdvancedGreTunnelSourceIp                          types.String                                                    `tfsdk:"advanced_gre_tunnel_source_ip"`
	AdvancedGreTunnelSourceIpVariable                  types.String                                                    `tfsdk:"advanced_gre_tunnel_source_ip_variable"`
	AdvancedXconnect                                   types.String                                                    `tfsdk:"advanced_xconnect"`
	AdvancedXconnectVariable                           types.String                                                    `tfsdk:"advanced_xconnect_variable"`
	AdvancedLoadInterval                               types.Int64                                                     `tfsdk:"advanced_load_interval"`
	AdvancedLoadIntervalVariable                       types.String                                                    `tfsdk:"advanced_load_interval_variable"`
	AdvancedTracker                                    types.String                                                    `tfsdk:"advanced_tracker"`
	AdvancedTrackerVariable                            types.String                                                    `tfsdk:"advanced_tracker_variable"`
	AdvancedIpDirectedBroadcast                        types.Bool                                                      `tfsdk:"advanced_ip_directed_broadcast"`
	AdvancedIpDirectedBroadcastVariable                types.String                                                    `tfsdk:"advanced_ip_directed_broadcast_variable"`
}

type TransportWANVPNInterfaceEthernetIpv4SecondaryAddresses struct {
	Address            types.String `tfsdk:"address"`
	AddressVariable    types.String `tfsdk:"address_variable"`
	SubnetMask         types.String `tfsdk:"subnet_mask"`
	SubnetMaskVariable types.String `tfsdk:"subnet_mask_variable"`
}

type TransportWANVPNInterfaceEthernetIpv6DhcpSecondaryAddress struct {
	Address         types.String `tfsdk:"address"`
	AddressVariable types.String `tfsdk:"address_variable"`
}

type TransportWANVPNInterfaceEthernetIpv6SecondaryAddress struct {
	Address         types.String `tfsdk:"address"`
	AddressVariable types.String `tfsdk:"address_variable"`
}

type TransportWANVPNInterfaceEthernetTunnelInterfaceEncapsulations struct {
	Encapsulation      types.String `tfsdk:"encapsulation"`
	Preference         types.Int64  `tfsdk:"preference"`
	PreferenceVariable types.String `tfsdk:"preference_variable"`
	Weight             types.Int64  `tfsdk:"weight"`
	WeightVariable     types.String `tfsdk:"weight_variable"`
}

type TransportWANVPNInterfaceEthernetNewStaticNats struct {
	SourceIp             types.String `tfsdk:"source_ip"`
	SourceIpVariable     types.String `tfsdk:"source_ip_variable"`
	TranslatedIp         types.String `tfsdk:"translated_ip"`
	TranslatedIpVariable types.String `tfsdk:"translated_ip_variable"`
	Direction            types.String `tfsdk:"direction"`
	SourceVpn            types.Int64  `tfsdk:"source_vpn"`
	SourceVpnVariable    types.String `tfsdk:"source_vpn_variable"`
}

type TransportWANVPNInterfaceEthernetStaticNat66 struct {
	SourcePrefix                   types.String `tfsdk:"source_prefix"`
	SourcePrefixVariable           types.String `tfsdk:"source_prefix_variable"`
	TranslatedSourcePrefix         types.String `tfsdk:"translated_source_prefix"`
	TranslatedSourcePrefixVariable types.String `tfsdk:"translated_source_prefix_variable"`
	SourceVpnId                    types.Int64  `tfsdk:"source_vpn_id"`
	SourceVpnIdVariable            types.String `tfsdk:"source_vpn_id_variable"`
}

type TransportWANVPNInterfaceEthernetArps struct {
	IpAddress          types.String `tfsdk:"ip_address"`
	IpAddressVariable  types.String `tfsdk:"ip_address_variable"`
	MacAddress         types.String `tfsdk:"mac_address"`
	MacAddressVariable types.String `tfsdk:"mac_address_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TransportWANVPNInterfaceEthernet) getModel() string {
	return "transport_wan_vpn_interface_ethernet"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportWANVPNInterfaceEthernet) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/wan/vpn/%s/interface/ethernet", url.QueryEscape(data.FeatureProfileId.ValueString()), url.QueryEscape(data.TransportWanVpnProfileParcelId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportWANVPNInterfaceEthernet) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.ShutdownVariable.IsNull() {
		body, _ = sjson.Set(body, path+"shutdown.optionType", "variable")
		body, _ = sjson.Set(body, path+"shutdown.value", data.ShutdownVariable.ValueString())
	} else if data.Shutdown.IsNull() {
		body, _ = sjson.Set(body, path+"shutdown.optionType", "default")
		body, _ = sjson.Set(body, path+"shutdown.value", true)
	} else {
		body, _ = sjson.Set(body, path+"shutdown.optionType", "global")
		body, _ = sjson.Set(body, path+"shutdown.value", data.Shutdown.ValueBool())
	}

	if !data.InterfaceNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"interfaceName.optionType", "variable")
		body, _ = sjson.Set(body, path+"interfaceName.value", data.InterfaceNameVariable.ValueString())
	} else if !data.InterfaceName.IsNull() {
		body, _ = sjson.Set(body, path+"interfaceName.optionType", "global")
		body, _ = sjson.Set(body, path+"interfaceName.value", data.InterfaceName.ValueString())
	}

	if !data.ConfigDescriptionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"description.optionType", "variable")
		body, _ = sjson.Set(body, path+"description.value", data.ConfigDescriptionVariable.ValueString())
	} else if data.ConfigDescription.IsNull() {
		body, _ = sjson.Set(body, path+"description.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"description.optionType", "global")
		body, _ = sjson.Set(body, path+"description.value", data.ConfigDescription.ValueString())
	}

	if !data.Ipv4DhcpDistanceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"intfIpAddress.dynamic.dynamicDhcpDistance.optionType", "variable")
		body, _ = sjson.Set(body, path+"intfIpAddress.dynamic.dynamicDhcpDistance.value", data.Ipv4DhcpDistanceVariable.ValueString())
	} else if !data.Ipv4DhcpDistance.IsNull() {
		body, _ = sjson.Set(body, path+"intfIpAddress.dynamic.dynamicDhcpDistance.optionType", "global")
		body, _ = sjson.Set(body, path+"intfIpAddress.dynamic.dynamicDhcpDistance.value", data.Ipv4DhcpDistance.ValueInt64())
	}

	if !data.Ipv4AddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.optionType", "variable")
		body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.value", data.Ipv4AddressVariable.ValueString())
	} else if !data.Ipv4Address.IsNull() {
		body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.optionType", "global")
		body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.value", data.Ipv4Address.ValueString())
	}

	if !data.Ipv4SubnetMaskVariable.IsNull() {
		body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.optionType", "variable")
		body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.value", data.Ipv4SubnetMaskVariable.ValueString())
	} else if !data.Ipv4SubnetMask.IsNull() {
		body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.optionType", "global")
		body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.value", data.Ipv4SubnetMask.ValueString())
	}
	body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressSecondary", []interface{}{})
	for _, item := range data.Ipv4SecondaryAddresses {
		itemBody := ""

		if !item.AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ipAddress.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "ipAddress.value", item.AddressVariable.ValueString())
		} else if !item.Address.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ipAddress.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "ipAddress.value", item.Address.ValueString())
		}

		if !item.SubnetMaskVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "subnetMask.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "subnetMask.value", item.SubnetMaskVariable.ValueString())
		} else if !item.SubnetMask.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "subnetMask.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "subnetMask.value", item.SubnetMask.ValueString())
		}
		body, _ = sjson.SetRaw(body, path+"intfIpAddress.static.staticIpV4AddressSecondary.-1", itemBody)
	}

	if !data.Ipv4DhcpHelperVariable.IsNull() {
		body, _ = sjson.Set(body, path+"dhcpHelper.optionType", "variable")
		body, _ = sjson.Set(body, path+"dhcpHelper.value", data.Ipv4DhcpHelperVariable.ValueString())
	} else if data.Ipv4DhcpHelper.IsNull() {
		body, _ = sjson.Set(body, path+"dhcpHelper.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"dhcpHelper.optionType", "global")
		var values []string
		data.Ipv4DhcpHelper.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"dhcpHelper.value", values)
	}
	if !data.EnableDhcpv6.IsNull() {
		body, _ = sjson.Set(body, path+"intfIpV6Address.dynamic.dhcpClient.optionType", "global")
		body, _ = sjson.Set(body, path+"intfIpV6Address.dynamic.dhcpClient.value", data.EnableDhcpv6.ValueBool())
	}

	for _, item := range data.Ipv6DhcpSecondaryAddress {
		itemBody := ""

		if !item.AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "address.value", item.AddressVariable.ValueString())
		} else if !item.Address.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "address.value", item.Address.ValueString())
		}
		body, _ = sjson.SetRaw(body, path+"intfIpV6Address.dynamic.secondaryIpV6Address.-1", itemBody)
	}

	if !data.Ipv6AddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"intfIpV6Address.static.primaryIpV6Address.address.optionType", "variable")
		body, _ = sjson.Set(body, path+"intfIpV6Address.static.primaryIpV6Address.address.value", data.Ipv6AddressVariable.ValueString())
	} else if !data.Ipv6Address.IsNull() {
		body, _ = sjson.Set(body, path+"intfIpV6Address.static.primaryIpV6Address.address.optionType", "global")
		body, _ = sjson.Set(body, path+"intfIpV6Address.static.primaryIpV6Address.address.value", data.Ipv6Address.ValueString())
	}

	for _, item := range data.Ipv6SecondaryAddress {
		itemBody := ""

		if !item.AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "address.value", item.AddressVariable.ValueString())
		} else if !item.Address.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "address.value", item.Address.ValueString())
		}
		body, _ = sjson.SetRaw(body, path+"intfIpV6Address.static.secondaryIpV6Address.-1", itemBody)
	}

	if !data.IperfServerVariable.IsNull() {
		body, _ = sjson.Set(body, path+"iperfServer.optionType", "variable")
		body, _ = sjson.Set(body, path+"iperfServer.value", data.IperfServerVariable.ValueString())
	} else if data.IperfServer.IsNull() {
		body, _ = sjson.Set(body, path+"iperfServer.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"iperfServer.optionType", "global")
		body, _ = sjson.Set(body, path+"iperfServer.value", data.IperfServer.ValueString())
	}

	if !data.BlockNonSourceIpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"blockNonSourceIp.optionType", "variable")
		body, _ = sjson.Set(body, path+"blockNonSourceIp.value", data.BlockNonSourceIpVariable.ValueString())
	} else if data.BlockNonSourceIp.IsNull() {
		body, _ = sjson.Set(body, path+"blockNonSourceIp.optionType", "default")
		body, _ = sjson.Set(body, path+"blockNonSourceIp.value", false)
	} else {
		body, _ = sjson.Set(body, path+"blockNonSourceIp.optionType", "global")
		body, _ = sjson.Set(body, path+"blockNonSourceIp.value", data.BlockNonSourceIp.ValueBool())
	}

	if !data.ServiceProviderVariable.IsNull() {
		body, _ = sjson.Set(body, path+"serviceProvider.optionType", "variable")
		body, _ = sjson.Set(body, path+"serviceProvider.value", data.ServiceProviderVariable.ValueString())
	} else if data.ServiceProvider.IsNull() {
		body, _ = sjson.Set(body, path+"serviceProvider.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"serviceProvider.optionType", "global")
		body, _ = sjson.Set(body, path+"serviceProvider.value", data.ServiceProvider.ValueString())
	}

	if !data.BandwidthUpstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bandwidthUpstream.optionType", "variable")
		body, _ = sjson.Set(body, path+"bandwidthUpstream.value", data.BandwidthUpstreamVariable.ValueString())
	} else if data.BandwidthUpstream.IsNull() {
		body, _ = sjson.Set(body, path+"bandwidthUpstream.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"bandwidthUpstream.optionType", "global")
		body, _ = sjson.Set(body, path+"bandwidthUpstream.value", data.BandwidthUpstream.ValueInt64())
	}

	if !data.BandwidthDownstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bandwidthDownstream.optionType", "variable")
		body, _ = sjson.Set(body, path+"bandwidthDownstream.value", data.BandwidthDownstreamVariable.ValueString())
	} else if data.BandwidthDownstream.IsNull() {
		body, _ = sjson.Set(body, path+"bandwidthDownstream.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"bandwidthDownstream.optionType", "global")
		body, _ = sjson.Set(body, path+"bandwidthDownstream.value", data.BandwidthDownstream.ValueInt64())
	}

	if !data.AutoDetectBandwidthVariable.IsNull() {
		body, _ = sjson.Set(body, path+"autoDetectBandwidth.optionType", "variable")
		body, _ = sjson.Set(body, path+"autoDetectBandwidth.value", data.AutoDetectBandwidthVariable.ValueString())
	} else if data.AutoDetectBandwidth.IsNull() {
		body, _ = sjson.Set(body, path+"autoDetectBandwidth.optionType", "default")
		body, _ = sjson.Set(body, path+"autoDetectBandwidth.value", false)
	} else {
		body, _ = sjson.Set(body, path+"autoDetectBandwidth.optionType", "global")
		body, _ = sjson.Set(body, path+"autoDetectBandwidth.value", data.AutoDetectBandwidth.ValueBool())
	}
	if data.TunnelInterface.IsNull() {
		body, _ = sjson.Set(body, path+"tunnelInterface.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnelInterface.value", false)
	} else {
		body, _ = sjson.Set(body, path+"tunnelInterface.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnelInterface.value", data.TunnelInterface.ValueBool())
	}

	if !data.PerTunnelQosVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.value", data.PerTunnelQosVariable.ValueString())
	} else if data.PerTunnelQos.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.value", false)
	} else {
		body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.value", data.PerTunnelQos.ValueBool())
	}

	if !data.TunnelQosModeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.mode.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.mode.value", data.TunnelQosModeVariable.ValueString())
	} else if !data.TunnelQosMode.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.mode.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.mode.value", data.TunnelQosMode.ValueString())
	}

	if !data.TunnelBandwidthPercentVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.value", data.TunnelBandwidthPercentVariable.ValueString())
	} else if data.TunnelBandwidthPercent.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.value", 50)
	} else {
		body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.value", data.TunnelBandwidthPercent.ValueInt64())
	}

	if !data.TunnelInterfaceBindLoopbackTunnelVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.bind.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.bind.value", data.TunnelInterfaceBindLoopbackTunnelVariable.ValueString())
	} else if data.TunnelInterfaceBindLoopbackTunnel.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.bind.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"tunnel.bind.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.bind.value", data.TunnelInterfaceBindLoopbackTunnel.ValueString())
	}

	if !data.TunnelInterfaceCarrierVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.carrier.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.carrier.value", data.TunnelInterfaceCarrierVariable.ValueString())
	} else if data.TunnelInterfaceCarrier.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.carrier.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.carrier.value", "default")
	} else {
		body, _ = sjson.Set(body, path+"tunnel.carrier.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.carrier.value", data.TunnelInterfaceCarrier.ValueString())
	}

	if !data.TunnelInterfaceColorVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.color.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.color.value", data.TunnelInterfaceColorVariable.ValueString())
	} else if data.TunnelInterfaceColor.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.color.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.color.value", "mpls")
	} else {
		body, _ = sjson.Set(body, path+"tunnel.color.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.color.value", data.TunnelInterfaceColor.ValueString())
	}

	if !data.TunnelInterfaceHelloIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.helloInterval.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.helloInterval.value", data.TunnelInterfaceHelloIntervalVariable.ValueString())
	} else if data.TunnelInterfaceHelloInterval.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.helloInterval.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.helloInterval.value", 1000)
	} else {
		body, _ = sjson.Set(body, path+"tunnel.helloInterval.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.helloInterval.value", data.TunnelInterfaceHelloInterval.ValueInt64())
	}

	if !data.TunnelInterfaceHelloToleranceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.helloTolerance.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.helloTolerance.value", data.TunnelInterfaceHelloToleranceVariable.ValueString())
	} else if data.TunnelInterfaceHelloTolerance.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.helloTolerance.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.helloTolerance.value", 12)
	} else {
		body, _ = sjson.Set(body, path+"tunnel.helloTolerance.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.helloTolerance.value", data.TunnelInterfaceHelloTolerance.ValueInt64())
	}

	if !data.TunnelInterfaceLastResortCircuitVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.value", data.TunnelInterfaceLastResortCircuitVariable.ValueString())
	} else if data.TunnelInterfaceLastResortCircuit.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.value", false)
	} else {
		body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.value", data.TunnelInterfaceLastResortCircuit.ValueBool())
	}

	if !data.TunnelInterfaceGreTunnelDestinationIpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.tlocExtensionGreTo.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.tlocExtensionGreTo.value", data.TunnelInterfaceGreTunnelDestinationIpVariable.ValueString())
	} else if data.TunnelInterfaceGreTunnelDestinationIp.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.tlocExtensionGreTo.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"tunnel.tlocExtensionGreTo.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.tlocExtensionGreTo.value", data.TunnelInterfaceGreTunnelDestinationIp.ValueString())
	}

	if !data.TunnelInterfaceColorRestrictVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.restrict.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.restrict.value", data.TunnelInterfaceColorRestrictVariable.ValueString())
	} else if data.TunnelInterfaceColorRestrict.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.restrict.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.restrict.value", false)
	} else {
		body, _ = sjson.Set(body, path+"tunnel.restrict.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.restrict.value", data.TunnelInterfaceColorRestrict.ValueBool())
	}

	if !data.TunnelInterfaceGroupsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.group.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.group.value", data.TunnelInterfaceGroupsVariable.ValueString())
	} else if data.TunnelInterfaceGroups.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.group.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"tunnel.group.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.group.value", data.TunnelInterfaceGroups.ValueInt64())
	}

	if !data.TunnelInterfaceBorderVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.border.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.border.value", data.TunnelInterfaceBorderVariable.ValueString())
	} else if data.TunnelInterfaceBorder.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.border.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.border.value", false)
	} else {
		body, _ = sjson.Set(body, path+"tunnel.border.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.border.value", data.TunnelInterfaceBorder.ValueBool())
	}

	if !data.TunnelInterfaceMaxControlConnectionsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.maxControlConnections.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.maxControlConnections.value", data.TunnelInterfaceMaxControlConnectionsVariable.ValueString())
	} else if data.TunnelInterfaceMaxControlConnections.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.maxControlConnections.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"tunnel.maxControlConnections.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.maxControlConnections.value", data.TunnelInterfaceMaxControlConnections.ValueInt64())
	}

	if !data.TunnelInterfaceNatRefreshIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.value", data.TunnelInterfaceNatRefreshIntervalVariable.ValueString())
	} else if data.TunnelInterfaceNatRefreshInterval.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.value", 5)
	} else {
		body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.value", data.TunnelInterfaceNatRefreshInterval.ValueInt64())
	}

	if !data.TunnelInterfaceVbondAsStunServerVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.vBondAsStunServer.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.vBondAsStunServer.value", data.TunnelInterfaceVbondAsStunServerVariable.ValueString())
	} else if data.TunnelInterfaceVbondAsStunServer.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.vBondAsStunServer.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.vBondAsStunServer.value", false)
	} else {
		body, _ = sjson.Set(body, path+"tunnel.vBondAsStunServer.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.vBondAsStunServer.value", data.TunnelInterfaceVbondAsStunServer.ValueBool())
	}

	if !data.TunnelInterfaceExcludeControllerGroupListVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.excludeControllerGroupList.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.excludeControllerGroupList.value", data.TunnelInterfaceExcludeControllerGroupListVariable.ValueString())
	} else if data.TunnelInterfaceExcludeControllerGroupList.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.excludeControllerGroupList.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"tunnel.excludeControllerGroupList.optionType", "global")
		var values []int64
		data.TunnelInterfaceExcludeControllerGroupList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"tunnel.excludeControllerGroupList.value", values)
	}

	if !data.TunnelInterfaceVmanageConnectionPreferenceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.vManageConnectionPreference.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.vManageConnectionPreference.value", data.TunnelInterfaceVmanageConnectionPreferenceVariable.ValueString())
	} else if data.TunnelInterfaceVmanageConnectionPreference.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.vManageConnectionPreference.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.vManageConnectionPreference.value", 5)
	} else {
		body, _ = sjson.Set(body, path+"tunnel.vManageConnectionPreference.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.vManageConnectionPreference.value", data.TunnelInterfaceVmanageConnectionPreference.ValueInt64())
	}

	if !data.TunnelInterfacePortHopVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.portHop.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.portHop.value", data.TunnelInterfacePortHopVariable.ValueString())
	} else if data.TunnelInterfacePortHop.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.portHop.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.portHop.value", true)
	} else {
		body, _ = sjson.Set(body, path+"tunnel.portHop.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.portHop.value", data.TunnelInterfacePortHop.ValueBool())
	}

	if !data.TunnelInterfaceLowBandwidthLinkVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.value", data.TunnelInterfaceLowBandwidthLinkVariable.ValueString())
	} else if data.TunnelInterfaceLowBandwidthLink.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.value", false)
	} else {
		body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.value", data.TunnelInterfaceLowBandwidthLink.ValueBool())
	}

	if !data.TunnelInterfaceTunnelTcpMssVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.tunnelTcpMss.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.tunnelTcpMss.value", data.TunnelInterfaceTunnelTcpMssVariable.ValueString())
	} else if data.TunnelInterfaceTunnelTcpMss.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.tunnelTcpMss.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"tunnel.tunnelTcpMss.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.tunnelTcpMss.value", data.TunnelInterfaceTunnelTcpMss.ValueInt64())
	}

	if !data.TunnelInterfaceClearDontFragmentVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.value", data.TunnelInterfaceClearDontFragmentVariable.ValueString())
	} else if data.TunnelInterfaceClearDontFragment.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.value", false)
	} else {
		body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.value", data.TunnelInterfaceClearDontFragment.ValueBool())
	}

	if !data.TunnelInterfaceCtsSgtPropagationVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.ctsSgtPropagation.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.ctsSgtPropagation.value", data.TunnelInterfaceCtsSgtPropagationVariable.ValueString())
	} else if data.TunnelInterfaceCtsSgtPropagation.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.ctsSgtPropagation.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.ctsSgtPropagation.value", false)
	} else {
		body, _ = sjson.Set(body, path+"tunnel.ctsSgtPropagation.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.ctsSgtPropagation.value", data.TunnelInterfaceCtsSgtPropagation.ValueBool())
	}

	if !data.TunnelInterfaceNetworkBroadcastVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.value", data.TunnelInterfaceNetworkBroadcastVariable.ValueString())
	} else if data.TunnelInterfaceNetworkBroadcast.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.optionType", "default")
		body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.value", false)
	} else {
		body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.value", data.TunnelInterfaceNetworkBroadcast.ValueBool())
	}

	if !data.TunnelInterfaceAllowAllVariable.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.all.optionType", "variable")
		body, _ = sjson.Set(body, path+"allowService.all.value", data.TunnelInterfaceAllowAllVariable.ValueString())
	} else if data.TunnelInterfaceAllowAll.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.all.optionType", "default")
		body, _ = sjson.Set(body, path+"allowService.all.value", false)
	} else {
		body, _ = sjson.Set(body, path+"allowService.all.optionType", "global")
		body, _ = sjson.Set(body, path+"allowService.all.value", data.TunnelInterfaceAllowAll.ValueBool())
	}

	if !data.TunnelInterfaceAllowBgpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.bgp.optionType", "variable")
		body, _ = sjson.Set(body, path+"allowService.bgp.value", data.TunnelInterfaceAllowBgpVariable.ValueString())
	} else if data.TunnelInterfaceAllowBgp.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.bgp.optionType", "default")
		body, _ = sjson.Set(body, path+"allowService.bgp.value", false)
	} else {
		body, _ = sjson.Set(body, path+"allowService.bgp.optionType", "global")
		body, _ = sjson.Set(body, path+"allowService.bgp.value", data.TunnelInterfaceAllowBgp.ValueBool())
	}

	if !data.TunnelInterfaceAllowDhcpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.dhcp.optionType", "variable")
		body, _ = sjson.Set(body, path+"allowService.dhcp.value", data.TunnelInterfaceAllowDhcpVariable.ValueString())
	} else if data.TunnelInterfaceAllowDhcp.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.dhcp.optionType", "default")
		body, _ = sjson.Set(body, path+"allowService.dhcp.value", true)
	} else {
		body, _ = sjson.Set(body, path+"allowService.dhcp.optionType", "global")
		body, _ = sjson.Set(body, path+"allowService.dhcp.value", data.TunnelInterfaceAllowDhcp.ValueBool())
	}

	if !data.TunnelInterfaceAllowNtpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.ntp.optionType", "variable")
		body, _ = sjson.Set(body, path+"allowService.ntp.value", data.TunnelInterfaceAllowNtpVariable.ValueString())
	} else if !data.TunnelInterfaceAllowNtp.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.ntp.optionType", "global")
		body, _ = sjson.Set(body, path+"allowService.ntp.value", data.TunnelInterfaceAllowNtp.ValueBool())
	}

	if !data.TunnelInterfaceAllowSshVariable.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.ssh.optionType", "variable")
		body, _ = sjson.Set(body, path+"allowService.ssh.value", data.TunnelInterfaceAllowSshVariable.ValueString())
	} else if !data.TunnelInterfaceAllowSsh.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.ssh.optionType", "global")
		body, _ = sjson.Set(body, path+"allowService.ssh.value", data.TunnelInterfaceAllowSsh.ValueBool())
	}

	if !data.TunnelInterfaceAllowDbsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.dns.optionType", "variable")
		body, _ = sjson.Set(body, path+"allowService.dns.value", data.TunnelInterfaceAllowDbsVariable.ValueString())
	} else if data.TunnelInterfaceAllowDbs.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.dns.optionType", "default")
		body, _ = sjson.Set(body, path+"allowService.dns.value", true)
	} else {
		body, _ = sjson.Set(body, path+"allowService.dns.optionType", "global")
		body, _ = sjson.Set(body, path+"allowService.dns.value", data.TunnelInterfaceAllowDbs.ValueBool())
	}

	if !data.TunnelInterfaceAllowIcmpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.icmp.optionType", "variable")
		body, _ = sjson.Set(body, path+"allowService.icmp.value", data.TunnelInterfaceAllowIcmpVariable.ValueString())
	} else if data.TunnelInterfaceAllowIcmp.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.icmp.optionType", "default")
		body, _ = sjson.Set(body, path+"allowService.icmp.value", true)
	} else {
		body, _ = sjson.Set(body, path+"allowService.icmp.optionType", "global")
		body, _ = sjson.Set(body, path+"allowService.icmp.value", data.TunnelInterfaceAllowIcmp.ValueBool())
	}

	if !data.TunnelInterfaceAllowHttpsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.https.optionType", "variable")
		body, _ = sjson.Set(body, path+"allowService.https.value", data.TunnelInterfaceAllowHttpsVariable.ValueString())
	} else if data.TunnelInterfaceAllowHttps.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.https.optionType", "default")
		body, _ = sjson.Set(body, path+"allowService.https.value", true)
	} else {
		body, _ = sjson.Set(body, path+"allowService.https.optionType", "global")
		body, _ = sjson.Set(body, path+"allowService.https.value", data.TunnelInterfaceAllowHttps.ValueBool())
	}

	if !data.TunnelInterfaceAllowOspfVariable.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.ospf.optionType", "variable")
		body, _ = sjson.Set(body, path+"allowService.ospf.value", data.TunnelInterfaceAllowOspfVariable.ValueString())
	} else if data.TunnelInterfaceAllowOspf.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.ospf.optionType", "default")
		body, _ = sjson.Set(body, path+"allowService.ospf.value", false)
	} else {
		body, _ = sjson.Set(body, path+"allowService.ospf.optionType", "global")
		body, _ = sjson.Set(body, path+"allowService.ospf.value", data.TunnelInterfaceAllowOspf.ValueBool())
	}

	if !data.TunnelInterfaceAllowStunVariable.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.stun.optionType", "variable")
		body, _ = sjson.Set(body, path+"allowService.stun.value", data.TunnelInterfaceAllowStunVariable.ValueString())
	} else if data.TunnelInterfaceAllowStun.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.stun.optionType", "default")
		body, _ = sjson.Set(body, path+"allowService.stun.value", false)
	} else {
		body, _ = sjson.Set(body, path+"allowService.stun.optionType", "global")
		body, _ = sjson.Set(body, path+"allowService.stun.value", data.TunnelInterfaceAllowStun.ValueBool())
	}

	if !data.TunnelInterfaceAllowSnmpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.snmp.optionType", "variable")
		body, _ = sjson.Set(body, path+"allowService.snmp.value", data.TunnelInterfaceAllowSnmpVariable.ValueString())
	} else if data.TunnelInterfaceAllowSnmp.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.snmp.optionType", "default")
		body, _ = sjson.Set(body, path+"allowService.snmp.value", false)
	} else {
		body, _ = sjson.Set(body, path+"allowService.snmp.optionType", "global")
		body, _ = sjson.Set(body, path+"allowService.snmp.value", data.TunnelInterfaceAllowSnmp.ValueBool())
	}

	if !data.TunnelInterfaceAllowNetconfVariable.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.netconf.optionType", "variable")
		body, _ = sjson.Set(body, path+"allowService.netconf.value", data.TunnelInterfaceAllowNetconfVariable.ValueString())
	} else if data.TunnelInterfaceAllowNetconf.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.netconf.optionType", "default")
		body, _ = sjson.Set(body, path+"allowService.netconf.value", false)
	} else {
		body, _ = sjson.Set(body, path+"allowService.netconf.optionType", "global")
		body, _ = sjson.Set(body, path+"allowService.netconf.value", data.TunnelInterfaceAllowNetconf.ValueBool())
	}

	if !data.TunnelInterfaceAllowBfdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.bfd.optionType", "variable")
		body, _ = sjson.Set(body, path+"allowService.bfd.value", data.TunnelInterfaceAllowBfdVariable.ValueString())
	} else if data.TunnelInterfaceAllowBfd.IsNull() {
		body, _ = sjson.Set(body, path+"allowService.bfd.optionType", "default")
		body, _ = sjson.Set(body, path+"allowService.bfd.value", false)
	} else {
		body, _ = sjson.Set(body, path+"allowService.bfd.optionType", "global")
		body, _ = sjson.Set(body, path+"allowService.bfd.value", data.TunnelInterfaceAllowBfd.ValueBool())
	}
	body, _ = sjson.Set(body, path+"encapsulation", []interface{}{})
	for _, item := range data.TunnelInterfaceEncapsulations {
		itemBody := ""
		if !item.Encapsulation.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "encap.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "encap.value", item.Encapsulation.ValueString())
		}

		if !item.PreferenceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "preference.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "preference.value", item.PreferenceVariable.ValueString())
		} else if item.Preference.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "preference.optionType", "default")

		} else {
			itemBody, _ = sjson.Set(itemBody, "preference.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "preference.value", item.Preference.ValueInt64())
		}

		if !item.WeightVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "weight.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "weight.value", item.WeightVariable.ValueString())
		} else if item.Weight.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "weight.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "weight.value", 1)
		} else {
			itemBody, _ = sjson.Set(itemBody, "weight.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "weight.value", item.Weight.ValueInt64())
		}
		body, _ = sjson.SetRaw(body, path+"encapsulation.-1", itemBody)
	}

	if !data.NatIpv4Variable.IsNull() {
		body, _ = sjson.Set(body, path+"nat.optionType", "variable")
		body, _ = sjson.Set(body, path+"nat.value", data.NatIpv4Variable.ValueString())
	} else if data.NatIpv4.IsNull() {
		body, _ = sjson.Set(body, path+"nat.optionType", "default")
		body, _ = sjson.Set(body, path+"nat.value", false)
	} else {
		body, _ = sjson.Set(body, path+"nat.optionType", "global")
		body, _ = sjson.Set(body, path+"nat.value", data.NatIpv4.ValueBool())
	}

	if !data.NatTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natType.optionType", "variable")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natType.value", data.NatTypeVariable.ValueString())
	} else if data.NatType.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natType.optionType", "default")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natType.value", "interface")
	} else {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natType.optionType", "global")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natType.value", data.NatType.ValueString())
	}

	if !data.NatRangeStartVariable.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeStart.optionType", "variable")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeStart.value", data.NatRangeStartVariable.ValueString())
	} else if data.NatRangeStart.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeStart.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeStart.optionType", "global")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeStart.value", data.NatRangeStart.ValueString())
	}

	if !data.NatRangeEndVariable.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeEnd.optionType", "variable")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeEnd.value", data.NatRangeEndVariable.ValueString())
	} else if data.NatRangeEnd.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeEnd.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeEnd.optionType", "global")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeEnd.value", data.NatRangeEnd.ValueString())
	}

	if !data.NatPrefixLengthVariable.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.prefixLength.optionType", "variable")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.prefixLength.value", data.NatPrefixLengthVariable.ValueString())
	} else if data.NatPrefixLength.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.prefixLength.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.prefixLength.optionType", "global")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.prefixLength.value", data.NatPrefixLength.ValueInt64())
	}

	if !data.NatOverloadVariable.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.overload.optionType", "variable")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.overload.value", data.NatOverloadVariable.ValueString())
	} else if data.NatOverload.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.overload.optionType", "default")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.overload.value", true)
	} else {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.overload.optionType", "global")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.overload.value", data.NatOverload.ValueBool())
	}

	if !data.NatLoopbackVariable.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natLookback.optionType", "variable")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natLookback.value", data.NatLoopbackVariable.ValueString())
	} else if data.NatLoopback.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natLookback.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natLookback.optionType", "global")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.natLookback.value", data.NatLoopback.ValueString())
	}

	if !data.NatUdpTimeoutVariable.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.optionType", "variable")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.value", data.NatUdpTimeoutVariable.ValueString())
	} else if data.NatUdpTimeout.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.optionType", "default")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.value", 1)
	} else {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.optionType", "global")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.value", data.NatUdpTimeout.ValueInt64())
	}

	if !data.NatTcpTimeoutVariable.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.optionType", "variable")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.value", data.NatTcpTimeoutVariable.ValueString())
	} else if data.NatTcpTimeout.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.optionType", "default")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.value", 60)
	} else {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.optionType", "global")
		body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.value", data.NatTcpTimeout.ValueInt64())
	}
	body, _ = sjson.Set(body, path+"natAttributesIpv4.newStaticNat", []interface{}{})
	for _, item := range data.NewStaticNats {
		itemBody := ""

		if !item.SourceIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sourceIp.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "sourceIp.value", item.SourceIpVariable.ValueString())
		} else if !item.SourceIp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sourceIp.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sourceIp.value", item.SourceIp.ValueString())
		}

		if !item.TranslatedIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "translateIp.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "translateIp.value", item.TranslatedIpVariable.ValueString())
		} else if !item.TranslatedIp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "translateIp.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "translateIp.value", item.TranslatedIp.ValueString())
		}
		if item.Direction.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "staticNatDirection.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "staticNatDirection.value", "inside")
		} else {
			itemBody, _ = sjson.Set(itemBody, "staticNatDirection.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "staticNatDirection.value", item.Direction.ValueString())
		}

		if !item.SourceVpnVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sourceVpn.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "sourceVpn.value", item.SourceVpnVariable.ValueString())
		} else if item.SourceVpn.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sourceVpn.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "sourceVpn.value", 0)
		} else {
			itemBody, _ = sjson.Set(itemBody, "sourceVpn.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sourceVpn.value", item.SourceVpn.ValueInt64())
		}
		body, _ = sjson.SetRaw(body, path+"natAttributesIpv4.newStaticNat.-1", itemBody)
	}

	if !data.NatIpv6Variable.IsNull() {
		body, _ = sjson.Set(body, path+"natIpv6.optionType", "variable")
		body, _ = sjson.Set(body, path+"natIpv6.value", data.NatIpv6Variable.ValueString())
	} else if data.NatIpv6.IsNull() {
		body, _ = sjson.Set(body, path+"natIpv6.optionType", "default")
		body, _ = sjson.Set(body, path+"natIpv6.value", false)
	} else {
		body, _ = sjson.Set(body, path+"natIpv6.optionType", "global")
		body, _ = sjson.Set(body, path+"natIpv6.value", data.NatIpv6.ValueBool())
	}
	if data.Nat64.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv6.nat64.optionType", "default")
		body, _ = sjson.Set(body, path+"natAttributesIpv6.nat64.value", false)
	} else {
		body, _ = sjson.Set(body, path+"natAttributesIpv6.nat64.optionType", "global")
		body, _ = sjson.Set(body, path+"natAttributesIpv6.nat64.value", data.Nat64.ValueBool())
	}
	if data.Nat66.IsNull() {
		body, _ = sjson.Set(body, path+"natAttributesIpv6.nat66.optionType", "default")
		body, _ = sjson.Set(body, path+"natAttributesIpv6.nat66.value", false)
	} else {
		body, _ = sjson.Set(body, path+"natAttributesIpv6.nat66.optionType", "global")
		body, _ = sjson.Set(body, path+"natAttributesIpv6.nat66.value", data.Nat66.ValueBool())
	}

	for _, item := range data.StaticNat66 {
		itemBody := ""

		if !item.SourcePrefixVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sourcePrefix.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "sourcePrefix.value", item.SourcePrefixVariable.ValueString())
		} else if !item.SourcePrefix.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sourcePrefix.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sourcePrefix.value", item.SourcePrefix.ValueString())
		}

		if !item.TranslatedSourcePrefixVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "translatedSourcePrefix.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "translatedSourcePrefix.value", item.TranslatedSourcePrefixVariable.ValueString())
		} else if !item.TranslatedSourcePrefix.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "translatedSourcePrefix.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "translatedSourcePrefix.value", item.TranslatedSourcePrefix.ValueString())
		}

		if !item.SourceVpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sourceVpnId.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "sourceVpnId.value", item.SourceVpnIdVariable.ValueString())
		} else if item.SourceVpnId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sourceVpnId.optionType", "default")

		} else {
			itemBody, _ = sjson.Set(itemBody, "sourceVpnId.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sourceVpnId.value", item.SourceVpnId.ValueInt64())
		}
		body, _ = sjson.SetRaw(body, path+"natAttributesIpv6.staticNat66.-1", itemBody)
	}
	if data.AdaptiveQos.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.adaptiveQoS.optionType", "default")
		body, _ = sjson.Set(body, path+"aclQos.adaptiveQoS.value", false)
	} else {
		body, _ = sjson.Set(body, path+"aclQos.adaptiveQoS.optionType", "global")
		body, _ = sjson.Set(body, path+"aclQos.adaptiveQoS.value", data.AdaptiveQos.ValueBool())
	}

	if !data.QosAdaptivePeriodVariable.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.adaptPeriod.optionType", "variable")
		body, _ = sjson.Set(body, path+"aclQos.adaptPeriod.value", data.QosAdaptivePeriodVariable.ValueString())
	} else if data.QosAdaptivePeriod.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.adaptPeriod.optionType", "default")
		body, _ = sjson.Set(body, path+"aclQos.adaptPeriod.value", 15)
	} else {
		body, _ = sjson.Set(body, path+"aclQos.adaptPeriod.optionType", "global")
		body, _ = sjson.Set(body, path+"aclQos.adaptPeriod.value", data.QosAdaptivePeriod.ValueInt64())
	}
	if data.QosAdaptiveBandwidthUpstream.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstream.optionType", "default")
		body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstream.value", false)
	} else {
		body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstream.optionType", "global")
		body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstream.value", data.QosAdaptiveBandwidthUpstream.ValueBool())
	}

	if !data.QosAdaptiveMinUpstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.optionType", "variable")
		body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.value", data.QosAdaptiveMinUpstreamVariable.ValueString())
	} else if !data.QosAdaptiveMinUpstream.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.optionType", "global")
		body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.value", data.QosAdaptiveMinUpstream.ValueInt64())
	}

	if !data.QosAdaptiveMaxUpstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.optionType", "variable")
		body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.value", data.QosAdaptiveMaxUpstreamVariable.ValueString())
	} else if !data.QosAdaptiveMaxUpstream.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.optionType", "global")
		body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.value", data.QosAdaptiveMaxUpstream.ValueInt64())
	}

	if !data.QosAdaptiveDefaultUpstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.optionType", "variable")
		body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.value", data.QosAdaptiveDefaultUpstreamVariable.ValueString())
	} else if !data.QosAdaptiveDefaultUpstream.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.optionType", "global")
		body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.value", data.QosAdaptiveDefaultUpstream.ValueInt64())
	}
	if data.QosAdaptiveBandwidthDownstream.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstream.optionType", "default")
		body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstream.value", false)
	} else {
		body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstream.optionType", "global")
		body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstream.value", data.QosAdaptiveBandwidthDownstream.ValueBool())
	}

	if !data.QosAdaptiveMinDownstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.optionType", "variable")
		body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.value", data.QosAdaptiveMinDownstreamVariable.ValueString())
	} else if !data.QosAdaptiveMinDownstream.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.optionType", "global")
		body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.value", data.QosAdaptiveMinDownstream.ValueInt64())
	}

	if !data.QosAdaptiveMaxDownstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.optionType", "variable")
		body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.value", data.QosAdaptiveMaxDownstreamVariable.ValueString())
	} else if !data.QosAdaptiveMaxDownstream.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.optionType", "global")
		body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.value", data.QosAdaptiveMaxDownstream.ValueInt64())
	}

	if !data.QosAdaptiveDefaultDownstreamVariable.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.optionType", "variable")
		body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.value", data.QosAdaptiveDefaultDownstreamVariable.ValueString())
	} else if !data.QosAdaptiveDefaultDownstream.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.optionType", "global")
		body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.value", data.QosAdaptiveDefaultDownstream.ValueInt64())
	}

	if !data.QosShapingRateVariable.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.shapingRate.optionType", "variable")
		body, _ = sjson.Set(body, path+"aclQos.shapingRate.value", data.QosShapingRateVariable.ValueString())
	} else if !data.QosShapingRate.IsNull() {
		body, _ = sjson.Set(body, path+"aclQos.shapingRate.optionType", "global")
		body, _ = sjson.Set(body, path+"aclQos.shapingRate.value", data.QosShapingRate.ValueInt64())
	}
	body, _ = sjson.Set(body, path+"arp", []interface{}{})
	for _, item := range data.Arps {
		itemBody := ""

		if !item.IpAddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ipAddress.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "ipAddress.value", item.IpAddressVariable.ValueString())
		} else if item.IpAddress.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ipAddress.optionType", "default")

		} else {
			itemBody, _ = sjson.Set(itemBody, "ipAddress.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "ipAddress.value", item.IpAddress.ValueString())
		}

		if !item.MacAddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "macAddress.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "macAddress.value", item.MacAddressVariable.ValueString())
		} else if item.MacAddress.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "macAddress.optionType", "default")

		} else {
			itemBody, _ = sjson.Set(itemBody, "macAddress.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "macAddress.value", item.MacAddress.ValueString())
		}
		body, _ = sjson.SetRaw(body, path+"arp.-1", itemBody)
	}

	if !data.AdvancedIcmpRedirectDisableVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.icmpRedirectDisable.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.icmpRedirectDisable.value", data.AdvancedIcmpRedirectDisableVariable.ValueString())
	} else if data.AdvancedIcmpRedirectDisable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.icmpRedirectDisable.optionType", "default")
		body, _ = sjson.Set(body, path+"advanced.icmpRedirectDisable.value", true)
	} else {
		body, _ = sjson.Set(body, path+"advanced.icmpRedirectDisable.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.icmpRedirectDisable.value", data.AdvancedIcmpRedirectDisable.ValueBool())
	}

	if !data.AdvancedDuplexVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.duplex.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.duplex.value", data.AdvancedDuplexVariable.ValueString())
	} else if data.AdvancedDuplex.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.duplex.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"advanced.duplex.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.duplex.value", data.AdvancedDuplex.ValueString())
	}

	if !data.AdvancedMacAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.macAddress.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.macAddress.value", data.AdvancedMacAddressVariable.ValueString())
	} else if data.AdvancedMacAddress.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.macAddress.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"advanced.macAddress.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.macAddress.value", data.AdvancedMacAddress.ValueString())
	}

	if !data.AdvancedIpMtuVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.ipMtu.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.ipMtu.value", data.AdvancedIpMtuVariable.ValueString())
	} else if data.AdvancedIpMtu.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.ipMtu.optionType", "default")
		body, _ = sjson.Set(body, path+"advanced.ipMtu.value", 1500)
	} else {
		body, _ = sjson.Set(body, path+"advanced.ipMtu.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.ipMtu.value", data.AdvancedIpMtu.ValueInt64())
	}

	if !data.AdvancedInterfaceMtuVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.intrfMtu.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.intrfMtu.value", data.AdvancedInterfaceMtuVariable.ValueString())
	} else if data.AdvancedInterfaceMtu.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.intrfMtu.optionType", "default")
		body, _ = sjson.Set(body, path+"advanced.intrfMtu.value", 1500)
	} else {
		body, _ = sjson.Set(body, path+"advanced.intrfMtu.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.intrfMtu.value", data.AdvancedInterfaceMtu.ValueInt64())
	}

	if !data.AdvancedTcpMssVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.tcpMss.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.tcpMss.value", data.AdvancedTcpMssVariable.ValueString())
	} else if data.AdvancedTcpMss.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.tcpMss.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"advanced.tcpMss.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.tcpMss.value", data.AdvancedTcpMss.ValueInt64())
	}

	if !data.AdvancedSpeedVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.speed.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.speed.value", data.AdvancedSpeedVariable.ValueString())
	} else if data.AdvancedSpeed.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.speed.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"advanced.speed.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.speed.value", data.AdvancedSpeed.ValueString())
	}

	if !data.AdvancedArpTimeoutVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.arpTimeout.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.arpTimeout.value", data.AdvancedArpTimeoutVariable.ValueString())
	} else if data.AdvancedArpTimeout.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.arpTimeout.optionType", "default")
		body, _ = sjson.Set(body, path+"advanced.arpTimeout.value", 1200)
	} else {
		body, _ = sjson.Set(body, path+"advanced.arpTimeout.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.arpTimeout.value", data.AdvancedArpTimeout.ValueInt64())
	}

	if !data.AdvancedAutonegotiateVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.autonegotiate.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.autonegotiate.value", data.AdvancedAutonegotiateVariable.ValueString())
	} else if data.AdvancedAutonegotiate.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.autonegotiate.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"advanced.autonegotiate.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.autonegotiate.value", data.AdvancedAutonegotiate.ValueBool())
	}

	if !data.AdvancedMediaTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.mediaType.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.mediaType.value", data.AdvancedMediaTypeVariable.ValueString())
	} else if data.AdvancedMediaType.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.mediaType.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"advanced.mediaType.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.mediaType.value", data.AdvancedMediaType.ValueString())
	}

	if !data.AdvancedTlocExtensionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.tlocExtension.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.tlocExtension.value", data.AdvancedTlocExtensionVariable.ValueString())
	} else if data.AdvancedTlocExtension.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.tlocExtension.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"advanced.tlocExtension.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.tlocExtension.value", data.AdvancedTlocExtension.ValueString())
	}

	if !data.AdvancedGreTunnelSourceIpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.sourceIp.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.sourceIp.value", data.AdvancedGreTunnelSourceIpVariable.ValueString())
	} else if data.AdvancedGreTunnelSourceIp.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.sourceIp.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.sourceIp.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.sourceIp.value", data.AdvancedGreTunnelSourceIp.ValueString())
	}

	if !data.AdvancedXconnectVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.xconnect.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.xconnect.value", data.AdvancedXconnectVariable.ValueString())
	} else if data.AdvancedXconnect.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.xconnect.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.xconnect.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.xconnect.value", data.AdvancedXconnect.ValueString())
	}

	if !data.AdvancedLoadIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.loadInterval.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.loadInterval.value", data.AdvancedLoadIntervalVariable.ValueString())
	} else if data.AdvancedLoadInterval.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.loadInterval.optionType", "default")
		body, _ = sjson.Set(body, path+"advanced.loadInterval.value", 30)
	} else {
		body, _ = sjson.Set(body, path+"advanced.loadInterval.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.loadInterval.value", data.AdvancedLoadInterval.ValueInt64())
	}

	if !data.AdvancedTrackerVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.tracker.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.tracker.value", data.AdvancedTrackerVariable.ValueString())
	} else if data.AdvancedTracker.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.tracker.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"advanced.tracker.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.tracker.value", data.AdvancedTracker.ValueString())
	}

	if !data.AdvancedIpDirectedBroadcastVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.ipDirectedBroadcast.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.ipDirectedBroadcast.value", data.AdvancedIpDirectedBroadcastVariable.ValueString())
	} else if data.AdvancedIpDirectedBroadcast.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.ipDirectedBroadcast.optionType", "default")
		body, _ = sjson.Set(body, path+"advanced.ipDirectedBroadcast.value", false)
	} else {
		body, _ = sjson.Set(body, path+"advanced.ipDirectedBroadcast.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.ipDirectedBroadcast.value", data.AdvancedIpDirectedBroadcast.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportWANVPNInterfaceEthernet) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Shutdown = types.BoolNull()
	data.ShutdownVariable = types.StringNull()
	if t := res.Get(path + "shutdown.optionType"); t.Exists() {
		va := res.Get(path + "shutdown.value")
		if t.String() == "variable" {
			data.ShutdownVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Shutdown = types.BoolValue(va.Bool())
		}
	}
	data.InterfaceName = types.StringNull()
	data.InterfaceNameVariable = types.StringNull()
	if t := res.Get(path + "interfaceName.optionType"); t.Exists() {
		va := res.Get(path + "interfaceName.value")
		if t.String() == "variable" {
			data.InterfaceNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceName = types.StringValue(va.String())
		}
	}
	data.ConfigDescription = types.StringNull()
	data.ConfigDescriptionVariable = types.StringNull()
	if t := res.Get(path + "description.optionType"); t.Exists() {
		va := res.Get(path + "description.value")
		if t.String() == "variable" {
			data.ConfigDescriptionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ConfigDescription = types.StringValue(va.String())
		}
	}
	data.Ipv4DhcpDistance = types.Int64Null()
	data.Ipv4DhcpDistanceVariable = types.StringNull()
	if t := res.Get(path + "intfIpAddress.dynamic.dynamicDhcpDistance.optionType"); t.Exists() {
		va := res.Get(path + "intfIpAddress.dynamic.dynamicDhcpDistance.value")
		if t.String() == "variable" {
			data.Ipv4DhcpDistanceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4DhcpDistance = types.Int64Value(va.Int())
		}
	}
	data.Ipv4Address = types.StringNull()
	data.Ipv4AddressVariable = types.StringNull()
	if t := res.Get(path + "intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.optionType"); t.Exists() {
		va := res.Get(path + "intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.value")
		if t.String() == "variable" {
			data.Ipv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Address = types.StringValue(va.String())
		}
	}
	data.Ipv4SubnetMask = types.StringNull()
	data.Ipv4SubnetMaskVariable = types.StringNull()
	if t := res.Get(path + "intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.optionType"); t.Exists() {
		va := res.Get(path + "intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.value")
		if t.String() == "variable" {
			data.Ipv4SubnetMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4SubnetMask = types.StringValue(va.String())
		}
	}
	if value := res.Get(path + "intfIpAddress.static.staticIpV4AddressSecondary"); value.Exists() {
		data.Ipv4SecondaryAddresses = make([]TransportWANVPNInterfaceEthernetIpv4SecondaryAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportWANVPNInterfaceEthernetIpv4SecondaryAddresses{}
			item.Address = types.StringNull()
			item.AddressVariable = types.StringNull()
			if t := v.Get("ipAddress.optionType"); t.Exists() {
				va := v.Get("ipAddress.value")
				if t.String() == "variable" {
					item.AddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Address = types.StringValue(va.String())
				}
			}
			item.SubnetMask = types.StringNull()
			item.SubnetMaskVariable = types.StringNull()
			if t := v.Get("subnetMask.optionType"); t.Exists() {
				va := v.Get("subnetMask.value")
				if t.String() == "variable" {
					item.SubnetMaskVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SubnetMask = types.StringValue(va.String())
				}
			}
			data.Ipv4SecondaryAddresses = append(data.Ipv4SecondaryAddresses, item)
			return true
		})
	}
	data.Ipv4DhcpHelper = types.SetNull(types.StringType)
	data.Ipv4DhcpHelperVariable = types.StringNull()
	if t := res.Get(path + "dhcpHelper.optionType"); t.Exists() {
		va := res.Get(path + "dhcpHelper.value")
		if t.String() == "variable" {
			data.Ipv4DhcpHelperVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4DhcpHelper = helpers.GetStringSet(va.Array())
		}
	}
	data.EnableDhcpv6 = types.BoolNull()

	if t := res.Get(path + "intfIpV6Address.dynamic.dhcpClient.optionType"); t.Exists() {
		va := res.Get(path + "intfIpV6Address.dynamic.dhcpClient.value")
		if t.String() == "global" {
			data.EnableDhcpv6 = types.BoolValue(va.Bool())
		}
	}
	if value := res.Get(path + "intfIpV6Address.dynamic.secondaryIpV6Address"); value.Exists() {
		data.Ipv6DhcpSecondaryAddress = make([]TransportWANVPNInterfaceEthernetIpv6DhcpSecondaryAddress, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportWANVPNInterfaceEthernetIpv6DhcpSecondaryAddress{}
			item.Address = types.StringNull()
			item.AddressVariable = types.StringNull()
			if t := v.Get("address.optionType"); t.Exists() {
				va := v.Get("address.value")
				if t.String() == "variable" {
					item.AddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Address = types.StringValue(va.String())
				}
			}
			data.Ipv6DhcpSecondaryAddress = append(data.Ipv6DhcpSecondaryAddress, item)
			return true
		})
	}
	data.Ipv6Address = types.StringNull()
	data.Ipv6AddressVariable = types.StringNull()
	if t := res.Get(path + "intfIpV6Address.static.primaryIpV6Address.address.optionType"); t.Exists() {
		va := res.Get(path + "intfIpV6Address.static.primaryIpV6Address.address.value")
		if t.String() == "variable" {
			data.Ipv6AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6Address = types.StringValue(va.String())
		}
	}
	if value := res.Get(path + "intfIpV6Address.static.secondaryIpV6Address"); value.Exists() {
		data.Ipv6SecondaryAddress = make([]TransportWANVPNInterfaceEthernetIpv6SecondaryAddress, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportWANVPNInterfaceEthernetIpv6SecondaryAddress{}
			item.Address = types.StringNull()
			item.AddressVariable = types.StringNull()
			if t := v.Get("address.optionType"); t.Exists() {
				va := v.Get("address.value")
				if t.String() == "variable" {
					item.AddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Address = types.StringValue(va.String())
				}
			}
			data.Ipv6SecondaryAddress = append(data.Ipv6SecondaryAddress, item)
			return true
		})
	}
	data.IperfServer = types.StringNull()
	data.IperfServerVariable = types.StringNull()
	if t := res.Get(path + "iperfServer.optionType"); t.Exists() {
		va := res.Get(path + "iperfServer.value")
		if t.String() == "variable" {
			data.IperfServerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IperfServer = types.StringValue(va.String())
		}
	}
	data.BlockNonSourceIp = types.BoolNull()
	data.BlockNonSourceIpVariable = types.StringNull()
	if t := res.Get(path + "blockNonSourceIp.optionType"); t.Exists() {
		va := res.Get(path + "blockNonSourceIp.value")
		if t.String() == "variable" {
			data.BlockNonSourceIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.BlockNonSourceIp = types.BoolValue(va.Bool())
		}
	}
	data.ServiceProvider = types.StringNull()
	data.ServiceProviderVariable = types.StringNull()
	if t := res.Get(path + "serviceProvider.optionType"); t.Exists() {
		va := res.Get(path + "serviceProvider.value")
		if t.String() == "variable" {
			data.ServiceProviderVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ServiceProvider = types.StringValue(va.String())
		}
	}
	data.BandwidthUpstream = types.Int64Null()
	data.BandwidthUpstreamVariable = types.StringNull()
	if t := res.Get(path + "bandwidthUpstream.optionType"); t.Exists() {
		va := res.Get(path + "bandwidthUpstream.value")
		if t.String() == "variable" {
			data.BandwidthUpstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.BandwidthUpstream = types.Int64Value(va.Int())
		}
	}
	data.BandwidthDownstream = types.Int64Null()
	data.BandwidthDownstreamVariable = types.StringNull()
	if t := res.Get(path + "bandwidthDownstream.optionType"); t.Exists() {
		va := res.Get(path + "bandwidthDownstream.value")
		if t.String() == "variable" {
			data.BandwidthDownstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.BandwidthDownstream = types.Int64Value(va.Int())
		}
	}
	data.AutoDetectBandwidth = types.BoolNull()
	data.AutoDetectBandwidthVariable = types.StringNull()
	if t := res.Get(path + "autoDetectBandwidth.optionType"); t.Exists() {
		va := res.Get(path + "autoDetectBandwidth.value")
		if t.String() == "variable" {
			data.AutoDetectBandwidthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AutoDetectBandwidth = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterface = types.BoolNull()

	if t := res.Get(path + "tunnelInterface.optionType"); t.Exists() {
		va := res.Get(path + "tunnelInterface.value")
		if t.String() == "global" {
			data.TunnelInterface = types.BoolValue(va.Bool())
		}
	}
	data.PerTunnelQos = types.BoolNull()
	data.PerTunnelQosVariable = types.StringNull()
	if t := res.Get(path + "tunnel.perTunnelQos.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.perTunnelQos.value")
		if t.String() == "variable" {
			data.PerTunnelQosVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PerTunnelQos = types.BoolValue(va.Bool())
		}
	}
	data.TunnelQosMode = types.StringNull()
	data.TunnelQosModeVariable = types.StringNull()
	if t := res.Get(path + "tunnel.mode.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.mode.value")
		if t.String() == "variable" {
			data.TunnelQosModeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelQosMode = types.StringValue(va.String())
		}
	}
	data.TunnelBandwidthPercent = types.Int64Null()
	data.TunnelBandwidthPercentVariable = types.StringNull()
	if t := res.Get(path + "tunnel.bandwidthPercent.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.bandwidthPercent.value")
		if t.String() == "variable" {
			data.TunnelBandwidthPercentVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelBandwidthPercent = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceBindLoopbackTunnel = types.StringNull()
	data.TunnelInterfaceBindLoopbackTunnelVariable = types.StringNull()
	if t := res.Get(path + "tunnel.bind.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.bind.value")
		if t.String() == "variable" {
			data.TunnelInterfaceBindLoopbackTunnelVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceBindLoopbackTunnel = types.StringValue(va.String())
		}
	}
	data.TunnelInterfaceCarrier = types.StringNull()
	data.TunnelInterfaceCarrierVariable = types.StringNull()
	if t := res.Get(path + "tunnel.carrier.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.carrier.value")
		if t.String() == "variable" {
			data.TunnelInterfaceCarrierVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceCarrier = types.StringValue(va.String())
		}
	}
	data.TunnelInterfaceColor = types.StringNull()
	data.TunnelInterfaceColorVariable = types.StringNull()
	if t := res.Get(path + "tunnel.color.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.color.value")
		if t.String() == "variable" {
			data.TunnelInterfaceColorVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceColor = types.StringValue(va.String())
		}
	}
	data.TunnelInterfaceHelloInterval = types.Int64Null()
	data.TunnelInterfaceHelloIntervalVariable = types.StringNull()
	if t := res.Get(path + "tunnel.helloInterval.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.helloInterval.value")
		if t.String() == "variable" {
			data.TunnelInterfaceHelloIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceHelloInterval = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceHelloTolerance = types.Int64Null()
	data.TunnelInterfaceHelloToleranceVariable = types.StringNull()
	if t := res.Get(path + "tunnel.helloTolerance.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.helloTolerance.value")
		if t.String() == "variable" {
			data.TunnelInterfaceHelloToleranceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceHelloTolerance = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceLastResortCircuit = types.BoolNull()
	data.TunnelInterfaceLastResortCircuitVariable = types.StringNull()
	if t := res.Get(path + "tunnel.lastResortCircuit.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.lastResortCircuit.value")
		if t.String() == "variable" {
			data.TunnelInterfaceLastResortCircuitVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceLastResortCircuit = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceGreTunnelDestinationIp = types.StringNull()
	data.TunnelInterfaceGreTunnelDestinationIpVariable = types.StringNull()
	if t := res.Get(path + "tunnel.tlocExtensionGreTo.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.tlocExtensionGreTo.value")
		if t.String() == "variable" {
			data.TunnelInterfaceGreTunnelDestinationIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceGreTunnelDestinationIp = types.StringValue(va.String())
		}
	}
	data.TunnelInterfaceColorRestrict = types.BoolNull()
	data.TunnelInterfaceColorRestrictVariable = types.StringNull()
	if t := res.Get(path + "tunnel.restrict.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.restrict.value")
		if t.String() == "variable" {
			data.TunnelInterfaceColorRestrictVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceColorRestrict = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceGroups = types.Int64Null()
	data.TunnelInterfaceGroupsVariable = types.StringNull()
	if t := res.Get(path + "tunnel.group.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.group.value")
		if t.String() == "variable" {
			data.TunnelInterfaceGroupsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceGroups = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceBorder = types.BoolNull()
	data.TunnelInterfaceBorderVariable = types.StringNull()
	if t := res.Get(path + "tunnel.border.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.border.value")
		if t.String() == "variable" {
			data.TunnelInterfaceBorderVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceBorder = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceMaxControlConnections = types.Int64Null()
	data.TunnelInterfaceMaxControlConnectionsVariable = types.StringNull()
	if t := res.Get(path + "tunnel.maxControlConnections.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.maxControlConnections.value")
		if t.String() == "variable" {
			data.TunnelInterfaceMaxControlConnectionsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceMaxControlConnections = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceNatRefreshInterval = types.Int64Null()
	data.TunnelInterfaceNatRefreshIntervalVariable = types.StringNull()
	if t := res.Get(path + "tunnel.natRefreshInterval.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.natRefreshInterval.value")
		if t.String() == "variable" {
			data.TunnelInterfaceNatRefreshIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceNatRefreshInterval = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceVbondAsStunServer = types.BoolNull()
	data.TunnelInterfaceVbondAsStunServerVariable = types.StringNull()
	if t := res.Get(path + "tunnel.vBondAsStunServer.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.vBondAsStunServer.value")
		if t.String() == "variable" {
			data.TunnelInterfaceVbondAsStunServerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceVbondAsStunServer = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceExcludeControllerGroupList = types.SetNull(types.Int64Type)
	data.TunnelInterfaceExcludeControllerGroupListVariable = types.StringNull()
	if t := res.Get(path + "tunnel.excludeControllerGroupList.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.excludeControllerGroupList.value")
		if t.String() == "variable" {
			data.TunnelInterfaceExcludeControllerGroupListVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceExcludeControllerGroupList = helpers.GetInt64Set(va.Array())
		}
	}
	data.TunnelInterfaceVmanageConnectionPreference = types.Int64Null()
	data.TunnelInterfaceVmanageConnectionPreferenceVariable = types.StringNull()
	if t := res.Get(path + "tunnel.vManageConnectionPreference.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.vManageConnectionPreference.value")
		if t.String() == "variable" {
			data.TunnelInterfaceVmanageConnectionPreferenceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceVmanageConnectionPreference = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfacePortHop = types.BoolNull()
	data.TunnelInterfacePortHopVariable = types.StringNull()
	if t := res.Get(path + "tunnel.portHop.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.portHop.value")
		if t.String() == "variable" {
			data.TunnelInterfacePortHopVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfacePortHop = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceLowBandwidthLink = types.BoolNull()
	data.TunnelInterfaceLowBandwidthLinkVariable = types.StringNull()
	if t := res.Get(path + "tunnel.lowBandwidthLink.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.lowBandwidthLink.value")
		if t.String() == "variable" {
			data.TunnelInterfaceLowBandwidthLinkVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceLowBandwidthLink = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceTunnelTcpMss = types.Int64Null()
	data.TunnelInterfaceTunnelTcpMssVariable = types.StringNull()
	if t := res.Get(path + "tunnel.tunnelTcpMss.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.tunnelTcpMss.value")
		if t.String() == "variable" {
			data.TunnelInterfaceTunnelTcpMssVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceTunnelTcpMss = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceClearDontFragment = types.BoolNull()
	data.TunnelInterfaceClearDontFragmentVariable = types.StringNull()
	if t := res.Get(path + "tunnel.clearDontFragment.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.clearDontFragment.value")
		if t.String() == "variable" {
			data.TunnelInterfaceClearDontFragmentVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceClearDontFragment = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceCtsSgtPropagation = types.BoolNull()
	data.TunnelInterfaceCtsSgtPropagationVariable = types.StringNull()
	if t := res.Get(path + "tunnel.ctsSgtPropagation.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.ctsSgtPropagation.value")
		if t.String() == "variable" {
			data.TunnelInterfaceCtsSgtPropagationVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceCtsSgtPropagation = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceNetworkBroadcast = types.BoolNull()
	data.TunnelInterfaceNetworkBroadcastVariable = types.StringNull()
	if t := res.Get(path + "tunnel.networkBroadcast.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.networkBroadcast.value")
		if t.String() == "variable" {
			data.TunnelInterfaceNetworkBroadcastVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceNetworkBroadcast = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowAll = types.BoolNull()
	data.TunnelInterfaceAllowAllVariable = types.StringNull()
	if t := res.Get(path + "allowService.all.optionType"); t.Exists() {
		va := res.Get(path + "allowService.all.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowAllVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowAll = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowBgp = types.BoolNull()
	data.TunnelInterfaceAllowBgpVariable = types.StringNull()
	if t := res.Get(path + "allowService.bgp.optionType"); t.Exists() {
		va := res.Get(path + "allowService.bgp.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowBgpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowBgp = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowDhcp = types.BoolNull()
	data.TunnelInterfaceAllowDhcpVariable = types.StringNull()
	if t := res.Get(path + "allowService.dhcp.optionType"); t.Exists() {
		va := res.Get(path + "allowService.dhcp.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowDhcpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowDhcp = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowNtp = types.BoolNull()
	data.TunnelInterfaceAllowNtpVariable = types.StringNull()
	if t := res.Get(path + "allowService.ntp.optionType"); t.Exists() {
		va := res.Get(path + "allowService.ntp.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowNtpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowNtp = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowSsh = types.BoolNull()
	data.TunnelInterfaceAllowSshVariable = types.StringNull()
	if t := res.Get(path + "allowService.ssh.optionType"); t.Exists() {
		va := res.Get(path + "allowService.ssh.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowSshVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowSsh = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowDbs = types.BoolNull()
	data.TunnelInterfaceAllowDbsVariable = types.StringNull()
	if t := res.Get(path + "allowService.dns.optionType"); t.Exists() {
		va := res.Get(path + "allowService.dns.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowDbsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowDbs = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowIcmp = types.BoolNull()
	data.TunnelInterfaceAllowIcmpVariable = types.StringNull()
	if t := res.Get(path + "allowService.icmp.optionType"); t.Exists() {
		va := res.Get(path + "allowService.icmp.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowIcmpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowIcmp = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowHttps = types.BoolNull()
	data.TunnelInterfaceAllowHttpsVariable = types.StringNull()
	if t := res.Get(path + "allowService.https.optionType"); t.Exists() {
		va := res.Get(path + "allowService.https.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowHttpsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowHttps = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowOspf = types.BoolNull()
	data.TunnelInterfaceAllowOspfVariable = types.StringNull()
	if t := res.Get(path + "allowService.ospf.optionType"); t.Exists() {
		va := res.Get(path + "allowService.ospf.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowOspfVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowOspf = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowStun = types.BoolNull()
	data.TunnelInterfaceAllowStunVariable = types.StringNull()
	if t := res.Get(path + "allowService.stun.optionType"); t.Exists() {
		va := res.Get(path + "allowService.stun.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowStunVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowStun = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowSnmp = types.BoolNull()
	data.TunnelInterfaceAllowSnmpVariable = types.StringNull()
	if t := res.Get(path + "allowService.snmp.optionType"); t.Exists() {
		va := res.Get(path + "allowService.snmp.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowSnmpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowSnmp = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowNetconf = types.BoolNull()
	data.TunnelInterfaceAllowNetconfVariable = types.StringNull()
	if t := res.Get(path + "allowService.netconf.optionType"); t.Exists() {
		va := res.Get(path + "allowService.netconf.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowNetconfVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowNetconf = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowBfd = types.BoolNull()
	data.TunnelInterfaceAllowBfdVariable = types.StringNull()
	if t := res.Get(path + "allowService.bfd.optionType"); t.Exists() {
		va := res.Get(path + "allowService.bfd.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowBfdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowBfd = types.BoolValue(va.Bool())
		}
	}
	if value := res.Get(path + "encapsulation"); value.Exists() {
		data.TunnelInterfaceEncapsulations = make([]TransportWANVPNInterfaceEthernetTunnelInterfaceEncapsulations, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportWANVPNInterfaceEthernetTunnelInterfaceEncapsulations{}
			item.Encapsulation = types.StringNull()

			if t := v.Get("encap.optionType"); t.Exists() {
				va := v.Get("encap.value")
				if t.String() == "global" {
					item.Encapsulation = types.StringValue(va.String())
				}
			}
			item.Preference = types.Int64Null()
			item.PreferenceVariable = types.StringNull()
			if t := v.Get("preference.optionType"); t.Exists() {
				va := v.Get("preference.value")
				if t.String() == "variable" {
					item.PreferenceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Preference = types.Int64Value(va.Int())
				}
			}
			item.Weight = types.Int64Null()
			item.WeightVariable = types.StringNull()
			if t := v.Get("weight.optionType"); t.Exists() {
				va := v.Get("weight.value")
				if t.String() == "variable" {
					item.WeightVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Weight = types.Int64Value(va.Int())
				}
			}
			data.TunnelInterfaceEncapsulations = append(data.TunnelInterfaceEncapsulations, item)
			return true
		})
	}
	data.NatIpv4 = types.BoolNull()
	data.NatIpv4Variable = types.StringNull()
	if t := res.Get(path + "nat.optionType"); t.Exists() {
		va := res.Get(path + "nat.value")
		if t.String() == "variable" {
			data.NatIpv4Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatIpv4 = types.BoolValue(va.Bool())
		}
	}
	data.NatType = types.StringNull()
	data.NatTypeVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natType.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natType.value")
		if t.String() == "variable" {
			data.NatTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatType = types.StringValue(va.String())
		}
	}
	data.NatRangeStart = types.StringNull()
	data.NatRangeStartVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natPool.rangeStart.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natPool.rangeStart.value")
		if t.String() == "variable" {
			data.NatRangeStartVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatRangeStart = types.StringValue(va.String())
		}
	}
	data.NatRangeEnd = types.StringNull()
	data.NatRangeEndVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natPool.rangeEnd.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natPool.rangeEnd.value")
		if t.String() == "variable" {
			data.NatRangeEndVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatRangeEnd = types.StringValue(va.String())
		}
	}
	data.NatPrefixLength = types.Int64Null()
	data.NatPrefixLengthVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natPool.prefixLength.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natPool.prefixLength.value")
		if t.String() == "variable" {
			data.NatPrefixLengthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatPrefixLength = types.Int64Value(va.Int())
		}
	}
	data.NatOverload = types.BoolNull()
	data.NatOverloadVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natPool.overload.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natPool.overload.value")
		if t.String() == "variable" {
			data.NatOverloadVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatOverload = types.BoolValue(va.Bool())
		}
	}
	data.NatLoopback = types.StringNull()
	data.NatLoopbackVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natLookback.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natLookback.value")
		if t.String() == "variable" {
			data.NatLoopbackVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatLoopback = types.StringValue(va.String())
		}
	}
	data.NatUdpTimeout = types.Int64Null()
	data.NatUdpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.udpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.udpTimeout.value")
		if t.String() == "variable" {
			data.NatUdpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatUdpTimeout = types.Int64Value(va.Int())
		}
	}
	data.NatTcpTimeout = types.Int64Null()
	data.NatTcpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.tcpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.tcpTimeout.value")
		if t.String() == "variable" {
			data.NatTcpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatTcpTimeout = types.Int64Value(va.Int())
		}
	}
	if value := res.Get(path + "natAttributesIpv4.newStaticNat"); value.Exists() {
		data.NewStaticNats = make([]TransportWANVPNInterfaceEthernetNewStaticNats, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportWANVPNInterfaceEthernetNewStaticNats{}
			item.SourceIp = types.StringNull()
			item.SourceIpVariable = types.StringNull()
			if t := v.Get("sourceIp.optionType"); t.Exists() {
				va := v.Get("sourceIp.value")
				if t.String() == "variable" {
					item.SourceIpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourceIp = types.StringValue(va.String())
				}
			}
			item.TranslatedIp = types.StringNull()
			item.TranslatedIpVariable = types.StringNull()
			if t := v.Get("translateIp.optionType"); t.Exists() {
				va := v.Get("translateIp.value")
				if t.String() == "variable" {
					item.TranslatedIpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TranslatedIp = types.StringValue(va.String())
				}
			}
			item.Direction = types.StringNull()

			if t := v.Get("staticNatDirection.optionType"); t.Exists() {
				va := v.Get("staticNatDirection.value")
				if t.String() == "global" {
					item.Direction = types.StringValue(va.String())
				}
			}
			item.SourceVpn = types.Int64Null()
			item.SourceVpnVariable = types.StringNull()
			if t := v.Get("sourceVpn.optionType"); t.Exists() {
				va := v.Get("sourceVpn.value")
				if t.String() == "variable" {
					item.SourceVpnVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourceVpn = types.Int64Value(va.Int())
				}
			}
			data.NewStaticNats = append(data.NewStaticNats, item)
			return true
		})
	}
	data.NatIpv6 = types.BoolNull()
	data.NatIpv6Variable = types.StringNull()
	if t := res.Get(path + "natIpv6.optionType"); t.Exists() {
		va := res.Get(path + "natIpv6.value")
		if t.String() == "variable" {
			data.NatIpv6Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatIpv6 = types.BoolValue(va.Bool())
		}
	}
	data.Nat64 = types.BoolNull()

	if t := res.Get(path + "natAttributesIpv6.nat64.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv6.nat64.value")
		if t.String() == "global" {
			data.Nat64 = types.BoolValue(va.Bool())
		}
	}
	data.Nat66 = types.BoolNull()

	if t := res.Get(path + "natAttributesIpv6.nat66.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv6.nat66.value")
		if t.String() == "global" {
			data.Nat66 = types.BoolValue(va.Bool())
		}
	}
	if value := res.Get(path + "natAttributesIpv6.staticNat66"); value.Exists() {
		data.StaticNat66 = make([]TransportWANVPNInterfaceEthernetStaticNat66, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportWANVPNInterfaceEthernetStaticNat66{}
			item.SourcePrefix = types.StringNull()
			item.SourcePrefixVariable = types.StringNull()
			if t := v.Get("sourcePrefix.optionType"); t.Exists() {
				va := v.Get("sourcePrefix.value")
				if t.String() == "variable" {
					item.SourcePrefixVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourcePrefix = types.StringValue(va.String())
				}
			}
			item.TranslatedSourcePrefix = types.StringNull()
			item.TranslatedSourcePrefixVariable = types.StringNull()
			if t := v.Get("translatedSourcePrefix.optionType"); t.Exists() {
				va := v.Get("translatedSourcePrefix.value")
				if t.String() == "variable" {
					item.TranslatedSourcePrefixVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TranslatedSourcePrefix = types.StringValue(va.String())
				}
			}
			item.SourceVpnId = types.Int64Null()
			item.SourceVpnIdVariable = types.StringNull()
			if t := v.Get("sourceVpnId.optionType"); t.Exists() {
				va := v.Get("sourceVpnId.value")
				if t.String() == "variable" {
					item.SourceVpnIdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourceVpnId = types.Int64Value(va.Int())
				}
			}
			data.StaticNat66 = append(data.StaticNat66, item)
			return true
		})
	}
	data.AdaptiveQos = types.BoolNull()

	if t := res.Get(path + "aclQos.adaptiveQoS.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.adaptiveQoS.value")
		if t.String() == "global" {
			data.AdaptiveQos = types.BoolValue(va.Bool())
		}
	}
	data.QosAdaptivePeriod = types.Int64Null()
	data.QosAdaptivePeriodVariable = types.StringNull()
	if t := res.Get(path + "aclQos.adaptPeriod.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.adaptPeriod.value")
		if t.String() == "variable" {
			data.QosAdaptivePeriodVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.QosAdaptivePeriod = types.Int64Value(va.Int())
		}
	}
	data.QosAdaptiveBandwidthUpstream = types.BoolNull()

	if t := res.Get(path + "aclQos.shapingRateUpstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateUpstream.value")
		if t.String() == "global" {
			data.QosAdaptiveBandwidthUpstream = types.BoolValue(va.Bool())
		}
	}
	data.QosAdaptiveMinUpstream = types.Int64Null()
	data.QosAdaptiveMinUpstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.value")
		if t.String() == "variable" {
			data.QosAdaptiveMinUpstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.QosAdaptiveMinUpstream = types.Int64Value(va.Int())
		}
	}
	data.QosAdaptiveMaxUpstream = types.Int64Null()
	data.QosAdaptiveMaxUpstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.value")
		if t.String() == "variable" {
			data.QosAdaptiveMaxUpstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.QosAdaptiveMaxUpstream = types.Int64Value(va.Int())
		}
	}
	data.QosAdaptiveDefaultUpstream = types.Int64Null()
	data.QosAdaptiveDefaultUpstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.value")
		if t.String() == "variable" {
			data.QosAdaptiveDefaultUpstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.QosAdaptiveDefaultUpstream = types.Int64Value(va.Int())
		}
	}
	data.QosAdaptiveBandwidthDownstream = types.BoolNull()

	if t := res.Get(path + "aclQos.shapingRateDownstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateDownstream.value")
		if t.String() == "global" {
			data.QosAdaptiveBandwidthDownstream = types.BoolValue(va.Bool())
		}
	}
	data.QosAdaptiveMinDownstream = types.Int64Null()
	data.QosAdaptiveMinDownstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.value")
		if t.String() == "variable" {
			data.QosAdaptiveMinDownstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.QosAdaptiveMinDownstream = types.Int64Value(va.Int())
		}
	}
	data.QosAdaptiveMaxDownstream = types.Int64Null()
	data.QosAdaptiveMaxDownstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.value")
		if t.String() == "variable" {
			data.QosAdaptiveMaxDownstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.QosAdaptiveMaxDownstream = types.Int64Value(va.Int())
		}
	}
	data.QosAdaptiveDefaultDownstream = types.Int64Null()
	data.QosAdaptiveDefaultDownstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.value")
		if t.String() == "variable" {
			data.QosAdaptiveDefaultDownstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.QosAdaptiveDefaultDownstream = types.Int64Value(va.Int())
		}
	}
	data.QosShapingRate = types.Int64Null()
	data.QosShapingRateVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRate.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRate.value")
		if t.String() == "variable" {
			data.QosShapingRateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.QosShapingRate = types.Int64Value(va.Int())
		}
	}
	if value := res.Get(path + "arp"); value.Exists() {
		data.Arps = make([]TransportWANVPNInterfaceEthernetArps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportWANVPNInterfaceEthernetArps{}
			item.IpAddress = types.StringNull()
			item.IpAddressVariable = types.StringNull()
			if t := v.Get("ipAddress.optionType"); t.Exists() {
				va := v.Get("ipAddress.value")
				if t.String() == "variable" {
					item.IpAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IpAddress = types.StringValue(va.String())
				}
			}
			item.MacAddress = types.StringNull()
			item.MacAddressVariable = types.StringNull()
			if t := v.Get("macAddress.optionType"); t.Exists() {
				va := v.Get("macAddress.value")
				if t.String() == "variable" {
					item.MacAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.MacAddress = types.StringValue(va.String())
				}
			}
			data.Arps = append(data.Arps, item)
			return true
		})
	}
	data.AdvancedIcmpRedirectDisable = types.BoolNull()
	data.AdvancedIcmpRedirectDisableVariable = types.StringNull()
	if t := res.Get(path + "advanced.icmpRedirectDisable.optionType"); t.Exists() {
		va := res.Get(path + "advanced.icmpRedirectDisable.value")
		if t.String() == "variable" {
			data.AdvancedIcmpRedirectDisableVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedIcmpRedirectDisable = types.BoolValue(va.Bool())
		}
	}
	data.AdvancedDuplex = types.StringNull()
	data.AdvancedDuplexVariable = types.StringNull()
	if t := res.Get(path + "advanced.duplex.optionType"); t.Exists() {
		va := res.Get(path + "advanced.duplex.value")
		if t.String() == "variable" {
			data.AdvancedDuplexVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedDuplex = types.StringValue(va.String())
		}
	}
	data.AdvancedMacAddress = types.StringNull()
	data.AdvancedMacAddressVariable = types.StringNull()
	if t := res.Get(path + "advanced.macAddress.optionType"); t.Exists() {
		va := res.Get(path + "advanced.macAddress.value")
		if t.String() == "variable" {
			data.AdvancedMacAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedMacAddress = types.StringValue(va.String())
		}
	}
	data.AdvancedIpMtu = types.Int64Null()
	data.AdvancedIpMtuVariable = types.StringNull()
	if t := res.Get(path + "advanced.ipMtu.optionType"); t.Exists() {
		va := res.Get(path + "advanced.ipMtu.value")
		if t.String() == "variable" {
			data.AdvancedIpMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedIpMtu = types.Int64Value(va.Int())
		}
	}
	data.AdvancedInterfaceMtu = types.Int64Null()
	data.AdvancedInterfaceMtuVariable = types.StringNull()
	if t := res.Get(path + "advanced.intrfMtu.optionType"); t.Exists() {
		va := res.Get(path + "advanced.intrfMtu.value")
		if t.String() == "variable" {
			data.AdvancedInterfaceMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedInterfaceMtu = types.Int64Value(va.Int())
		}
	}
	data.AdvancedTcpMss = types.Int64Null()
	data.AdvancedTcpMssVariable = types.StringNull()
	if t := res.Get(path + "advanced.tcpMss.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tcpMss.value")
		if t.String() == "variable" {
			data.AdvancedTcpMssVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedTcpMss = types.Int64Value(va.Int())
		}
	}
	data.AdvancedSpeed = types.StringNull()
	data.AdvancedSpeedVariable = types.StringNull()
	if t := res.Get(path + "advanced.speed.optionType"); t.Exists() {
		va := res.Get(path + "advanced.speed.value")
		if t.String() == "variable" {
			data.AdvancedSpeedVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedSpeed = types.StringValue(va.String())
		}
	}
	data.AdvancedArpTimeout = types.Int64Null()
	data.AdvancedArpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "advanced.arpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "advanced.arpTimeout.value")
		if t.String() == "variable" {
			data.AdvancedArpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedArpTimeout = types.Int64Value(va.Int())
		}
	}
	data.AdvancedAutonegotiate = types.BoolNull()
	data.AdvancedAutonegotiateVariable = types.StringNull()
	if t := res.Get(path + "advanced.autonegotiate.optionType"); t.Exists() {
		va := res.Get(path + "advanced.autonegotiate.value")
		if t.String() == "variable" {
			data.AdvancedAutonegotiateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedAutonegotiate = types.BoolValue(va.Bool())
		}
	}
	data.AdvancedMediaType = types.StringNull()
	data.AdvancedMediaTypeVariable = types.StringNull()
	if t := res.Get(path + "advanced.mediaType.optionType"); t.Exists() {
		va := res.Get(path + "advanced.mediaType.value")
		if t.String() == "variable" {
			data.AdvancedMediaTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedMediaType = types.StringValue(va.String())
		}
	}
	data.AdvancedTlocExtension = types.StringNull()
	data.AdvancedTlocExtensionVariable = types.StringNull()
	if t := res.Get(path + "advanced.tlocExtension.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tlocExtension.value")
		if t.String() == "variable" {
			data.AdvancedTlocExtensionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedTlocExtension = types.StringValue(va.String())
		}
	}
	data.AdvancedGreTunnelSourceIp = types.StringNull()
	data.AdvancedGreTunnelSourceIpVariable = types.StringNull()
	if t := res.Get(path + "advanced.tlocExtensionGreFrom.sourceIp.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tlocExtensionGreFrom.sourceIp.value")
		if t.String() == "variable" {
			data.AdvancedGreTunnelSourceIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedGreTunnelSourceIp = types.StringValue(va.String())
		}
	}
	data.AdvancedXconnect = types.StringNull()
	data.AdvancedXconnectVariable = types.StringNull()
	if t := res.Get(path + "advanced.tlocExtensionGreFrom.xconnect.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tlocExtensionGreFrom.xconnect.value")
		if t.String() == "variable" {
			data.AdvancedXconnectVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedXconnect = types.StringValue(va.String())
		}
	}
	data.AdvancedLoadInterval = types.Int64Null()
	data.AdvancedLoadIntervalVariable = types.StringNull()
	if t := res.Get(path + "advanced.loadInterval.optionType"); t.Exists() {
		va := res.Get(path + "advanced.loadInterval.value")
		if t.String() == "variable" {
			data.AdvancedLoadIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedLoadInterval = types.Int64Value(va.Int())
		}
	}
	data.AdvancedTracker = types.StringNull()
	data.AdvancedTrackerVariable = types.StringNull()
	if t := res.Get(path + "advanced.tracker.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tracker.value")
		if t.String() == "variable" {
			data.AdvancedTrackerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedTracker = types.StringValue(va.String())
		}
	}
	data.AdvancedIpDirectedBroadcast = types.BoolNull()
	data.AdvancedIpDirectedBroadcastVariable = types.StringNull()
	if t := res.Get(path + "advanced.ipDirectedBroadcast.optionType"); t.Exists() {
		va := res.Get(path + "advanced.ipDirectedBroadcast.value")
		if t.String() == "variable" {
			data.AdvancedIpDirectedBroadcastVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedIpDirectedBroadcast = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TransportWANVPNInterfaceEthernet) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Shutdown = types.BoolNull()
	data.ShutdownVariable = types.StringNull()
	if t := res.Get(path + "shutdown.optionType"); t.Exists() {
		va := res.Get(path + "shutdown.value")
		if t.String() == "variable" {
			data.ShutdownVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Shutdown = types.BoolValue(va.Bool())
		}
	}
	data.InterfaceName = types.StringNull()
	data.InterfaceNameVariable = types.StringNull()
	if t := res.Get(path + "interfaceName.optionType"); t.Exists() {
		va := res.Get(path + "interfaceName.value")
		if t.String() == "variable" {
			data.InterfaceNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceName = types.StringValue(va.String())
		}
	}
	data.ConfigDescription = types.StringNull()
	data.ConfigDescriptionVariable = types.StringNull()
	if t := res.Get(path + "description.optionType"); t.Exists() {
		va := res.Get(path + "description.value")
		if t.String() == "variable" {
			data.ConfigDescriptionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ConfigDescription = types.StringValue(va.String())
		}
	}
	data.Ipv4DhcpDistance = types.Int64Null()
	data.Ipv4DhcpDistanceVariable = types.StringNull()
	if t := res.Get(path + "intfIpAddress.dynamic.dynamicDhcpDistance.optionType"); t.Exists() {
		va := res.Get(path + "intfIpAddress.dynamic.dynamicDhcpDistance.value")
		if t.String() == "variable" {
			data.Ipv4DhcpDistanceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4DhcpDistance = types.Int64Value(va.Int())
		}
	}
	data.Ipv4Address = types.StringNull()
	data.Ipv4AddressVariable = types.StringNull()
	if t := res.Get(path + "intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.optionType"); t.Exists() {
		va := res.Get(path + "intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.value")
		if t.String() == "variable" {
			data.Ipv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Address = types.StringValue(va.String())
		}
	}
	data.Ipv4SubnetMask = types.StringNull()
	data.Ipv4SubnetMaskVariable = types.StringNull()
	if t := res.Get(path + "intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.optionType"); t.Exists() {
		va := res.Get(path + "intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.value")
		if t.String() == "variable" {
			data.Ipv4SubnetMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4SubnetMask = types.StringValue(va.String())
		}
	}
	for i := range data.Ipv4SecondaryAddresses {
		keys := [...]string{"ipAddress"}
		keyValues := [...]string{data.Ipv4SecondaryAddresses[i].Address.ValueString()}
		keyValuesVariables := [...]string{data.Ipv4SecondaryAddresses[i].AddressVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "intfIpAddress.static.staticIpV4AddressSecondary").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Ipv4SecondaryAddresses[i].Address = types.StringNull()
		data.Ipv4SecondaryAddresses[i].AddressVariable = types.StringNull()
		if t := r.Get("ipAddress.optionType"); t.Exists() {
			va := r.Get("ipAddress.value")
			if t.String() == "variable" {
				data.Ipv4SecondaryAddresses[i].AddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4SecondaryAddresses[i].Address = types.StringValue(va.String())
			}
		}
		data.Ipv4SecondaryAddresses[i].SubnetMask = types.StringNull()
		data.Ipv4SecondaryAddresses[i].SubnetMaskVariable = types.StringNull()
		if t := r.Get("subnetMask.optionType"); t.Exists() {
			va := r.Get("subnetMask.value")
			if t.String() == "variable" {
				data.Ipv4SecondaryAddresses[i].SubnetMaskVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4SecondaryAddresses[i].SubnetMask = types.StringValue(va.String())
			}
		}
	}
	data.Ipv4DhcpHelper = types.SetNull(types.StringType)
	data.Ipv4DhcpHelperVariable = types.StringNull()
	if t := res.Get(path + "dhcpHelper.optionType"); t.Exists() {
		va := res.Get(path + "dhcpHelper.value")
		if t.String() == "variable" {
			data.Ipv4DhcpHelperVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4DhcpHelper = helpers.GetStringSet(va.Array())
		}
	}
	data.EnableDhcpv6 = types.BoolNull()

	if t := res.Get(path + "intfIpV6Address.dynamic.dhcpClient.optionType"); t.Exists() {
		va := res.Get(path + "intfIpV6Address.dynamic.dhcpClient.value")
		if t.String() == "global" {
			data.EnableDhcpv6 = types.BoolValue(va.Bool())
		}
	}
	for i := range data.Ipv6DhcpSecondaryAddress {
		keys := [...]string{"address"}
		keyValues := [...]string{data.Ipv6DhcpSecondaryAddress[i].Address.ValueString()}
		keyValuesVariables := [...]string{data.Ipv6DhcpSecondaryAddress[i].AddressVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "intfIpV6Address.dynamic.secondaryIpV6Address").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Ipv6DhcpSecondaryAddress[i].Address = types.StringNull()
		data.Ipv6DhcpSecondaryAddress[i].AddressVariable = types.StringNull()
		if t := r.Get("address.optionType"); t.Exists() {
			va := r.Get("address.value")
			if t.String() == "variable" {
				data.Ipv6DhcpSecondaryAddress[i].AddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6DhcpSecondaryAddress[i].Address = types.StringValue(va.String())
			}
		}
	}
	data.Ipv6Address = types.StringNull()
	data.Ipv6AddressVariable = types.StringNull()
	if t := res.Get(path + "intfIpV6Address.static.primaryIpV6Address.address.optionType"); t.Exists() {
		va := res.Get(path + "intfIpV6Address.static.primaryIpV6Address.address.value")
		if t.String() == "variable" {
			data.Ipv6AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6Address = types.StringValue(va.String())
		}
	}
	for i := range data.Ipv6SecondaryAddress {
		keys := [...]string{"address"}
		keyValues := [...]string{data.Ipv6SecondaryAddress[i].Address.ValueString()}
		keyValuesVariables := [...]string{data.Ipv6SecondaryAddress[i].AddressVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "intfIpV6Address.static.secondaryIpV6Address").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Ipv6SecondaryAddress[i].Address = types.StringNull()
		data.Ipv6SecondaryAddress[i].AddressVariable = types.StringNull()
		if t := r.Get("address.optionType"); t.Exists() {
			va := r.Get("address.value")
			if t.String() == "variable" {
				data.Ipv6SecondaryAddress[i].AddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6SecondaryAddress[i].Address = types.StringValue(va.String())
			}
		}
	}
	data.IperfServer = types.StringNull()
	data.IperfServerVariable = types.StringNull()
	if t := res.Get(path + "iperfServer.optionType"); t.Exists() {
		va := res.Get(path + "iperfServer.value")
		if t.String() == "variable" {
			data.IperfServerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IperfServer = types.StringValue(va.String())
		}
	}
	data.BlockNonSourceIp = types.BoolNull()
	data.BlockNonSourceIpVariable = types.StringNull()
	if t := res.Get(path + "blockNonSourceIp.optionType"); t.Exists() {
		va := res.Get(path + "blockNonSourceIp.value")
		if t.String() == "variable" {
			data.BlockNonSourceIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.BlockNonSourceIp = types.BoolValue(va.Bool())
		}
	}
	data.ServiceProvider = types.StringNull()
	data.ServiceProviderVariable = types.StringNull()
	if t := res.Get(path + "serviceProvider.optionType"); t.Exists() {
		va := res.Get(path + "serviceProvider.value")
		if t.String() == "variable" {
			data.ServiceProviderVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ServiceProvider = types.StringValue(va.String())
		}
	}
	data.BandwidthUpstream = types.Int64Null()
	data.BandwidthUpstreamVariable = types.StringNull()
	if t := res.Get(path + "bandwidthUpstream.optionType"); t.Exists() {
		va := res.Get(path + "bandwidthUpstream.value")
		if t.String() == "variable" {
			data.BandwidthUpstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.BandwidthUpstream = types.Int64Value(va.Int())
		}
	}
	data.BandwidthDownstream = types.Int64Null()
	data.BandwidthDownstreamVariable = types.StringNull()
	if t := res.Get(path + "bandwidthDownstream.optionType"); t.Exists() {
		va := res.Get(path + "bandwidthDownstream.value")
		if t.String() == "variable" {
			data.BandwidthDownstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.BandwidthDownstream = types.Int64Value(va.Int())
		}
	}
	data.AutoDetectBandwidth = types.BoolNull()
	data.AutoDetectBandwidthVariable = types.StringNull()
	if t := res.Get(path + "autoDetectBandwidth.optionType"); t.Exists() {
		va := res.Get(path + "autoDetectBandwidth.value")
		if t.String() == "variable" {
			data.AutoDetectBandwidthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AutoDetectBandwidth = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterface = types.BoolNull()

	if t := res.Get(path + "tunnelInterface.optionType"); t.Exists() {
		va := res.Get(path + "tunnelInterface.value")
		if t.String() == "global" {
			data.TunnelInterface = types.BoolValue(va.Bool())
		}
	}
	data.PerTunnelQos = types.BoolNull()
	data.PerTunnelQosVariable = types.StringNull()
	if t := res.Get(path + "tunnel.perTunnelQos.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.perTunnelQos.value")
		if t.String() == "variable" {
			data.PerTunnelQosVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PerTunnelQos = types.BoolValue(va.Bool())
		}
	}
	data.TunnelQosMode = types.StringNull()
	data.TunnelQosModeVariable = types.StringNull()
	if t := res.Get(path + "tunnel.mode.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.mode.value")
		if t.String() == "variable" {
			data.TunnelQosModeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelQosMode = types.StringValue(va.String())
		}
	}
	data.TunnelBandwidthPercent = types.Int64Null()
	data.TunnelBandwidthPercentVariable = types.StringNull()
	if t := res.Get(path + "tunnel.bandwidthPercent.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.bandwidthPercent.value")
		if t.String() == "variable" {
			data.TunnelBandwidthPercentVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelBandwidthPercent = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceBindLoopbackTunnel = types.StringNull()
	data.TunnelInterfaceBindLoopbackTunnelVariable = types.StringNull()
	if t := res.Get(path + "tunnel.bind.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.bind.value")
		if t.String() == "variable" {
			data.TunnelInterfaceBindLoopbackTunnelVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceBindLoopbackTunnel = types.StringValue(va.String())
		}
	}
	data.TunnelInterfaceCarrier = types.StringNull()
	data.TunnelInterfaceCarrierVariable = types.StringNull()
	if t := res.Get(path + "tunnel.carrier.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.carrier.value")
		if t.String() == "variable" {
			data.TunnelInterfaceCarrierVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceCarrier = types.StringValue(va.String())
		}
	}
	data.TunnelInterfaceColor = types.StringNull()
	data.TunnelInterfaceColorVariable = types.StringNull()
	if t := res.Get(path + "tunnel.color.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.color.value")
		if t.String() == "variable" {
			data.TunnelInterfaceColorVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceColor = types.StringValue(va.String())
		}
	}
	data.TunnelInterfaceHelloInterval = types.Int64Null()
	data.TunnelInterfaceHelloIntervalVariable = types.StringNull()
	if t := res.Get(path + "tunnel.helloInterval.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.helloInterval.value")
		if t.String() == "variable" {
			data.TunnelInterfaceHelloIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceHelloInterval = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceHelloTolerance = types.Int64Null()
	data.TunnelInterfaceHelloToleranceVariable = types.StringNull()
	if t := res.Get(path + "tunnel.helloTolerance.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.helloTolerance.value")
		if t.String() == "variable" {
			data.TunnelInterfaceHelloToleranceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceHelloTolerance = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceLastResortCircuit = types.BoolNull()
	data.TunnelInterfaceLastResortCircuitVariable = types.StringNull()
	if t := res.Get(path + "tunnel.lastResortCircuit.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.lastResortCircuit.value")
		if t.String() == "variable" {
			data.TunnelInterfaceLastResortCircuitVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceLastResortCircuit = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceGreTunnelDestinationIp = types.StringNull()
	data.TunnelInterfaceGreTunnelDestinationIpVariable = types.StringNull()
	if t := res.Get(path + "tunnel.tlocExtensionGreTo.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.tlocExtensionGreTo.value")
		if t.String() == "variable" {
			data.TunnelInterfaceGreTunnelDestinationIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceGreTunnelDestinationIp = types.StringValue(va.String())
		}
	}
	data.TunnelInterfaceColorRestrict = types.BoolNull()
	data.TunnelInterfaceColorRestrictVariable = types.StringNull()
	if t := res.Get(path + "tunnel.restrict.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.restrict.value")
		if t.String() == "variable" {
			data.TunnelInterfaceColorRestrictVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceColorRestrict = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceGroups = types.Int64Null()
	data.TunnelInterfaceGroupsVariable = types.StringNull()
	if t := res.Get(path + "tunnel.group.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.group.value")
		if t.String() == "variable" {
			data.TunnelInterfaceGroupsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceGroups = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceBorder = types.BoolNull()
	data.TunnelInterfaceBorderVariable = types.StringNull()
	if t := res.Get(path + "tunnel.border.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.border.value")
		if t.String() == "variable" {
			data.TunnelInterfaceBorderVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceBorder = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceMaxControlConnections = types.Int64Null()
	data.TunnelInterfaceMaxControlConnectionsVariable = types.StringNull()
	if t := res.Get(path + "tunnel.maxControlConnections.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.maxControlConnections.value")
		if t.String() == "variable" {
			data.TunnelInterfaceMaxControlConnectionsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceMaxControlConnections = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceNatRefreshInterval = types.Int64Null()
	data.TunnelInterfaceNatRefreshIntervalVariable = types.StringNull()
	if t := res.Get(path + "tunnel.natRefreshInterval.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.natRefreshInterval.value")
		if t.String() == "variable" {
			data.TunnelInterfaceNatRefreshIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceNatRefreshInterval = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceVbondAsStunServer = types.BoolNull()
	data.TunnelInterfaceVbondAsStunServerVariable = types.StringNull()
	if t := res.Get(path + "tunnel.vBondAsStunServer.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.vBondAsStunServer.value")
		if t.String() == "variable" {
			data.TunnelInterfaceVbondAsStunServerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceVbondAsStunServer = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceExcludeControllerGroupList = types.SetNull(types.Int64Type)
	data.TunnelInterfaceExcludeControllerGroupListVariable = types.StringNull()
	if t := res.Get(path + "tunnel.excludeControllerGroupList.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.excludeControllerGroupList.value")
		if t.String() == "variable" {
			data.TunnelInterfaceExcludeControllerGroupListVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceExcludeControllerGroupList = helpers.GetInt64Set(va.Array())
		}
	}
	data.TunnelInterfaceVmanageConnectionPreference = types.Int64Null()
	data.TunnelInterfaceVmanageConnectionPreferenceVariable = types.StringNull()
	if t := res.Get(path + "tunnel.vManageConnectionPreference.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.vManageConnectionPreference.value")
		if t.String() == "variable" {
			data.TunnelInterfaceVmanageConnectionPreferenceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceVmanageConnectionPreference = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfacePortHop = types.BoolNull()
	data.TunnelInterfacePortHopVariable = types.StringNull()
	if t := res.Get(path + "tunnel.portHop.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.portHop.value")
		if t.String() == "variable" {
			data.TunnelInterfacePortHopVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfacePortHop = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceLowBandwidthLink = types.BoolNull()
	data.TunnelInterfaceLowBandwidthLinkVariable = types.StringNull()
	if t := res.Get(path + "tunnel.lowBandwidthLink.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.lowBandwidthLink.value")
		if t.String() == "variable" {
			data.TunnelInterfaceLowBandwidthLinkVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceLowBandwidthLink = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceTunnelTcpMss = types.Int64Null()
	data.TunnelInterfaceTunnelTcpMssVariable = types.StringNull()
	if t := res.Get(path + "tunnel.tunnelTcpMss.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.tunnelTcpMss.value")
		if t.String() == "variable" {
			data.TunnelInterfaceTunnelTcpMssVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceTunnelTcpMss = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceClearDontFragment = types.BoolNull()
	data.TunnelInterfaceClearDontFragmentVariable = types.StringNull()
	if t := res.Get(path + "tunnel.clearDontFragment.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.clearDontFragment.value")
		if t.String() == "variable" {
			data.TunnelInterfaceClearDontFragmentVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceClearDontFragment = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceCtsSgtPropagation = types.BoolNull()
	data.TunnelInterfaceCtsSgtPropagationVariable = types.StringNull()
	if t := res.Get(path + "tunnel.ctsSgtPropagation.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.ctsSgtPropagation.value")
		if t.String() == "variable" {
			data.TunnelInterfaceCtsSgtPropagationVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceCtsSgtPropagation = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceNetworkBroadcast = types.BoolNull()
	data.TunnelInterfaceNetworkBroadcastVariable = types.StringNull()
	if t := res.Get(path + "tunnel.networkBroadcast.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.networkBroadcast.value")
		if t.String() == "variable" {
			data.TunnelInterfaceNetworkBroadcastVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceNetworkBroadcast = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowAll = types.BoolNull()
	data.TunnelInterfaceAllowAllVariable = types.StringNull()
	if t := res.Get(path + "allowService.all.optionType"); t.Exists() {
		va := res.Get(path + "allowService.all.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowAllVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowAll = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowBgp = types.BoolNull()
	data.TunnelInterfaceAllowBgpVariable = types.StringNull()
	if t := res.Get(path + "allowService.bgp.optionType"); t.Exists() {
		va := res.Get(path + "allowService.bgp.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowBgpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowBgp = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowDhcp = types.BoolNull()
	data.TunnelInterfaceAllowDhcpVariable = types.StringNull()
	if t := res.Get(path + "allowService.dhcp.optionType"); t.Exists() {
		va := res.Get(path + "allowService.dhcp.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowDhcpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowDhcp = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowNtp = types.BoolNull()
	data.TunnelInterfaceAllowNtpVariable = types.StringNull()
	if t := res.Get(path + "allowService.ntp.optionType"); t.Exists() {
		va := res.Get(path + "allowService.ntp.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowNtpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowNtp = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowSsh = types.BoolNull()
	data.TunnelInterfaceAllowSshVariable = types.StringNull()
	if t := res.Get(path + "allowService.ssh.optionType"); t.Exists() {
		va := res.Get(path + "allowService.ssh.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowSshVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowSsh = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowDbs = types.BoolNull()
	data.TunnelInterfaceAllowDbsVariable = types.StringNull()
	if t := res.Get(path + "allowService.dns.optionType"); t.Exists() {
		va := res.Get(path + "allowService.dns.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowDbsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowDbs = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowIcmp = types.BoolNull()
	data.TunnelInterfaceAllowIcmpVariable = types.StringNull()
	if t := res.Get(path + "allowService.icmp.optionType"); t.Exists() {
		va := res.Get(path + "allowService.icmp.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowIcmpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowIcmp = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowHttps = types.BoolNull()
	data.TunnelInterfaceAllowHttpsVariable = types.StringNull()
	if t := res.Get(path + "allowService.https.optionType"); t.Exists() {
		va := res.Get(path + "allowService.https.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowHttpsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowHttps = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowOspf = types.BoolNull()
	data.TunnelInterfaceAllowOspfVariable = types.StringNull()
	if t := res.Get(path + "allowService.ospf.optionType"); t.Exists() {
		va := res.Get(path + "allowService.ospf.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowOspfVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowOspf = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowStun = types.BoolNull()
	data.TunnelInterfaceAllowStunVariable = types.StringNull()
	if t := res.Get(path + "allowService.stun.optionType"); t.Exists() {
		va := res.Get(path + "allowService.stun.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowStunVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowStun = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowSnmp = types.BoolNull()
	data.TunnelInterfaceAllowSnmpVariable = types.StringNull()
	if t := res.Get(path + "allowService.snmp.optionType"); t.Exists() {
		va := res.Get(path + "allowService.snmp.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowSnmpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowSnmp = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowNetconf = types.BoolNull()
	data.TunnelInterfaceAllowNetconfVariable = types.StringNull()
	if t := res.Get(path + "allowService.netconf.optionType"); t.Exists() {
		va := res.Get(path + "allowService.netconf.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowNetconfVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowNetconf = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowBfd = types.BoolNull()
	data.TunnelInterfaceAllowBfdVariable = types.StringNull()
	if t := res.Get(path + "allowService.bfd.optionType"); t.Exists() {
		va := res.Get(path + "allowService.bfd.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowBfdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowBfd = types.BoolValue(va.Bool())
		}
	}
	for i := range data.TunnelInterfaceEncapsulations {
		keys := [...]string{"encap", "preference"}
		keyValues := [...]string{data.TunnelInterfaceEncapsulations[i].Encapsulation.ValueString(), strconv.FormatInt(data.TunnelInterfaceEncapsulations[i].Preference.ValueInt64(), 10)}
		keyValuesVariables := [...]string{"", data.TunnelInterfaceEncapsulations[i].PreferenceVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "encapsulation").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.TunnelInterfaceEncapsulations[i].Encapsulation = types.StringNull()

		if t := r.Get("encap.optionType"); t.Exists() {
			va := r.Get("encap.value")
			if t.String() == "global" {
				data.TunnelInterfaceEncapsulations[i].Encapsulation = types.StringValue(va.String())
			}
		}
		data.TunnelInterfaceEncapsulations[i].Preference = types.Int64Null()
		data.TunnelInterfaceEncapsulations[i].PreferenceVariable = types.StringNull()
		if t := r.Get("preference.optionType"); t.Exists() {
			va := r.Get("preference.value")
			if t.String() == "variable" {
				data.TunnelInterfaceEncapsulations[i].PreferenceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.TunnelInterfaceEncapsulations[i].Preference = types.Int64Value(va.Int())
			}
		}
		data.TunnelInterfaceEncapsulations[i].Weight = types.Int64Null()
		data.TunnelInterfaceEncapsulations[i].WeightVariable = types.StringNull()
		if t := r.Get("weight.optionType"); t.Exists() {
			va := r.Get("weight.value")
			if t.String() == "variable" {
				data.TunnelInterfaceEncapsulations[i].WeightVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.TunnelInterfaceEncapsulations[i].Weight = types.Int64Value(va.Int())
			}
		}
	}
	data.NatIpv4 = types.BoolNull()
	data.NatIpv4Variable = types.StringNull()
	if t := res.Get(path + "nat.optionType"); t.Exists() {
		va := res.Get(path + "nat.value")
		if t.String() == "variable" {
			data.NatIpv4Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatIpv4 = types.BoolValue(va.Bool())
		}
	}
	data.NatType = types.StringNull()
	data.NatTypeVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natType.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natType.value")
		if t.String() == "variable" {
			data.NatTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatType = types.StringValue(va.String())
		}
	}
	data.NatRangeStart = types.StringNull()
	data.NatRangeStartVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natPool.rangeStart.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natPool.rangeStart.value")
		if t.String() == "variable" {
			data.NatRangeStartVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatRangeStart = types.StringValue(va.String())
		}
	}
	data.NatRangeEnd = types.StringNull()
	data.NatRangeEndVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natPool.rangeEnd.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natPool.rangeEnd.value")
		if t.String() == "variable" {
			data.NatRangeEndVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatRangeEnd = types.StringValue(va.String())
		}
	}
	data.NatPrefixLength = types.Int64Null()
	data.NatPrefixLengthVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natPool.prefixLength.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natPool.prefixLength.value")
		if t.String() == "variable" {
			data.NatPrefixLengthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatPrefixLength = types.Int64Value(va.Int())
		}
	}
	data.NatOverload = types.BoolNull()
	data.NatOverloadVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natPool.overload.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natPool.overload.value")
		if t.String() == "variable" {
			data.NatOverloadVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatOverload = types.BoolValue(va.Bool())
		}
	}
	data.NatLoopback = types.StringNull()
	data.NatLoopbackVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natLookback.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natLookback.value")
		if t.String() == "variable" {
			data.NatLoopbackVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatLoopback = types.StringValue(va.String())
		}
	}
	data.NatUdpTimeout = types.Int64Null()
	data.NatUdpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.udpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.udpTimeout.value")
		if t.String() == "variable" {
			data.NatUdpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatUdpTimeout = types.Int64Value(va.Int())
		}
	}
	data.NatTcpTimeout = types.Int64Null()
	data.NatTcpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.tcpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.tcpTimeout.value")
		if t.String() == "variable" {
			data.NatTcpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatTcpTimeout = types.Int64Value(va.Int())
		}
	}
	for i := range data.NewStaticNats {
		keys := [...]string{"sourceIp", "translateIp"}
		keyValues := [...]string{data.NewStaticNats[i].SourceIp.ValueString(), data.NewStaticNats[i].TranslatedIp.ValueString()}
		keyValuesVariables := [...]string{data.NewStaticNats[i].SourceIpVariable.ValueString(), data.NewStaticNats[i].TranslatedIpVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "natAttributesIpv4.newStaticNat").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.NewStaticNats[i].SourceIp = types.StringNull()
		data.NewStaticNats[i].SourceIpVariable = types.StringNull()
		if t := r.Get("sourceIp.optionType"); t.Exists() {
			va := r.Get("sourceIp.value")
			if t.String() == "variable" {
				data.NewStaticNats[i].SourceIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NewStaticNats[i].SourceIp = types.StringValue(va.String())
			}
		}
		data.NewStaticNats[i].TranslatedIp = types.StringNull()
		data.NewStaticNats[i].TranslatedIpVariable = types.StringNull()
		if t := r.Get("translateIp.optionType"); t.Exists() {
			va := r.Get("translateIp.value")
			if t.String() == "variable" {
				data.NewStaticNats[i].TranslatedIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NewStaticNats[i].TranslatedIp = types.StringValue(va.String())
			}
		}
		data.NewStaticNats[i].Direction = types.StringNull()

		if t := r.Get("staticNatDirection.optionType"); t.Exists() {
			va := r.Get("staticNatDirection.value")
			if t.String() == "global" {
				data.NewStaticNats[i].Direction = types.StringValue(va.String())
			}
		}
		data.NewStaticNats[i].SourceVpn = types.Int64Null()
		data.NewStaticNats[i].SourceVpnVariable = types.StringNull()
		if t := r.Get("sourceVpn.optionType"); t.Exists() {
			va := r.Get("sourceVpn.value")
			if t.String() == "variable" {
				data.NewStaticNats[i].SourceVpnVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NewStaticNats[i].SourceVpn = types.Int64Value(va.Int())
			}
		}
	}
	data.NatIpv6 = types.BoolNull()
	data.NatIpv6Variable = types.StringNull()
	if t := res.Get(path + "natIpv6.optionType"); t.Exists() {
		va := res.Get(path + "natIpv6.value")
		if t.String() == "variable" {
			data.NatIpv6Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatIpv6 = types.BoolValue(va.Bool())
		}
	}
	data.Nat64 = types.BoolNull()

	if t := res.Get(path + "natAttributesIpv6.nat64.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv6.nat64.value")
		if t.String() == "global" {
			data.Nat64 = types.BoolValue(va.Bool())
		}
	}
	data.Nat66 = types.BoolNull()

	if t := res.Get(path + "natAttributesIpv6.nat66.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv6.nat66.value")
		if t.String() == "global" {
			data.Nat66 = types.BoolValue(va.Bool())
		}
	}
	for i := range data.StaticNat66 {
		keys := [...]string{"sourcePrefix", "translatedSourcePrefix"}
		keyValues := [...]string{data.StaticNat66[i].SourcePrefix.ValueString(), data.StaticNat66[i].TranslatedSourcePrefix.ValueString()}
		keyValuesVariables := [...]string{data.StaticNat66[i].SourcePrefixVariable.ValueString(), data.StaticNat66[i].TranslatedSourcePrefixVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "natAttributesIpv6.staticNat66").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.StaticNat66[i].SourcePrefix = types.StringNull()
		data.StaticNat66[i].SourcePrefixVariable = types.StringNull()
		if t := r.Get("sourcePrefix.optionType"); t.Exists() {
			va := r.Get("sourcePrefix.value")
			if t.String() == "variable" {
				data.StaticNat66[i].SourcePrefixVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.StaticNat66[i].SourcePrefix = types.StringValue(va.String())
			}
		}
		data.StaticNat66[i].TranslatedSourcePrefix = types.StringNull()
		data.StaticNat66[i].TranslatedSourcePrefixVariable = types.StringNull()
		if t := r.Get("translatedSourcePrefix.optionType"); t.Exists() {
			va := r.Get("translatedSourcePrefix.value")
			if t.String() == "variable" {
				data.StaticNat66[i].TranslatedSourcePrefixVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.StaticNat66[i].TranslatedSourcePrefix = types.StringValue(va.String())
			}
		}
		data.StaticNat66[i].SourceVpnId = types.Int64Null()
		data.StaticNat66[i].SourceVpnIdVariable = types.StringNull()
		if t := r.Get("sourceVpnId.optionType"); t.Exists() {
			va := r.Get("sourceVpnId.value")
			if t.String() == "variable" {
				data.StaticNat66[i].SourceVpnIdVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.StaticNat66[i].SourceVpnId = types.Int64Value(va.Int())
			}
		}
	}
	data.AdaptiveQos = types.BoolNull()

	if t := res.Get(path + "aclQos.adaptiveQoS.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.adaptiveQoS.value")
		if t.String() == "global" {
			data.AdaptiveQos = types.BoolValue(va.Bool())
		}
	}
	data.QosAdaptivePeriod = types.Int64Null()
	data.QosAdaptivePeriodVariable = types.StringNull()
	if t := res.Get(path + "aclQos.adaptPeriod.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.adaptPeriod.value")
		if t.String() == "variable" {
			data.QosAdaptivePeriodVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.QosAdaptivePeriod = types.Int64Value(va.Int())
		}
	}
	data.QosAdaptiveBandwidthUpstream = types.BoolNull()

	if t := res.Get(path + "aclQos.shapingRateUpstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateUpstream.value")
		if t.String() == "global" {
			data.QosAdaptiveBandwidthUpstream = types.BoolValue(va.Bool())
		}
	}
	data.QosAdaptiveMinUpstream = types.Int64Null()
	data.QosAdaptiveMinUpstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.value")
		if t.String() == "variable" {
			data.QosAdaptiveMinUpstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.QosAdaptiveMinUpstream = types.Int64Value(va.Int())
		}
	}
	data.QosAdaptiveMaxUpstream = types.Int64Null()
	data.QosAdaptiveMaxUpstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.value")
		if t.String() == "variable" {
			data.QosAdaptiveMaxUpstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.QosAdaptiveMaxUpstream = types.Int64Value(va.Int())
		}
	}
	data.QosAdaptiveDefaultUpstream = types.Int64Null()
	data.QosAdaptiveDefaultUpstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.value")
		if t.String() == "variable" {
			data.QosAdaptiveDefaultUpstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.QosAdaptiveDefaultUpstream = types.Int64Value(va.Int())
		}
	}
	data.QosAdaptiveBandwidthDownstream = types.BoolNull()

	if t := res.Get(path + "aclQos.shapingRateDownstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateDownstream.value")
		if t.String() == "global" {
			data.QosAdaptiveBandwidthDownstream = types.BoolValue(va.Bool())
		}
	}
	data.QosAdaptiveMinDownstream = types.Int64Null()
	data.QosAdaptiveMinDownstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.value")
		if t.String() == "variable" {
			data.QosAdaptiveMinDownstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.QosAdaptiveMinDownstream = types.Int64Value(va.Int())
		}
	}
	data.QosAdaptiveMaxDownstream = types.Int64Null()
	data.QosAdaptiveMaxDownstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.value")
		if t.String() == "variable" {
			data.QosAdaptiveMaxDownstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.QosAdaptiveMaxDownstream = types.Int64Value(va.Int())
		}
	}
	data.QosAdaptiveDefaultDownstream = types.Int64Null()
	data.QosAdaptiveDefaultDownstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.value")
		if t.String() == "variable" {
			data.QosAdaptiveDefaultDownstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.QosAdaptiveDefaultDownstream = types.Int64Value(va.Int())
		}
	}
	data.QosShapingRate = types.Int64Null()
	data.QosShapingRateVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRate.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRate.value")
		if t.String() == "variable" {
			data.QosShapingRateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.QosShapingRate = types.Int64Value(va.Int())
		}
	}
	for i := range data.Arps {
		keys := [...]string{"ipAddress", "macAddress"}
		keyValues := [...]string{data.Arps[i].IpAddress.ValueString(), data.Arps[i].MacAddress.ValueString()}
		keyValuesVariables := [...]string{data.Arps[i].IpAddressVariable.ValueString(), data.Arps[i].MacAddressVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "arp").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Arps[i].IpAddress = types.StringNull()
		data.Arps[i].IpAddressVariable = types.StringNull()
		if t := r.Get("ipAddress.optionType"); t.Exists() {
			va := r.Get("ipAddress.value")
			if t.String() == "variable" {
				data.Arps[i].IpAddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Arps[i].IpAddress = types.StringValue(va.String())
			}
		}
		data.Arps[i].MacAddress = types.StringNull()
		data.Arps[i].MacAddressVariable = types.StringNull()
		if t := r.Get("macAddress.optionType"); t.Exists() {
			va := r.Get("macAddress.value")
			if t.String() == "variable" {
				data.Arps[i].MacAddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Arps[i].MacAddress = types.StringValue(va.String())
			}
		}
	}
	data.AdvancedIcmpRedirectDisable = types.BoolNull()
	data.AdvancedIcmpRedirectDisableVariable = types.StringNull()
	if t := res.Get(path + "advanced.icmpRedirectDisable.optionType"); t.Exists() {
		va := res.Get(path + "advanced.icmpRedirectDisable.value")
		if t.String() == "variable" {
			data.AdvancedIcmpRedirectDisableVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedIcmpRedirectDisable = types.BoolValue(va.Bool())
		}
	}
	data.AdvancedDuplex = types.StringNull()
	data.AdvancedDuplexVariable = types.StringNull()
	if t := res.Get(path + "advanced.duplex.optionType"); t.Exists() {
		va := res.Get(path + "advanced.duplex.value")
		if t.String() == "variable" {
			data.AdvancedDuplexVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedDuplex = types.StringValue(va.String())
		}
	}
	data.AdvancedMacAddress = types.StringNull()
	data.AdvancedMacAddressVariable = types.StringNull()
	if t := res.Get(path + "advanced.macAddress.optionType"); t.Exists() {
		va := res.Get(path + "advanced.macAddress.value")
		if t.String() == "variable" {
			data.AdvancedMacAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedMacAddress = types.StringValue(va.String())
		}
	}
	data.AdvancedIpMtu = types.Int64Null()
	data.AdvancedIpMtuVariable = types.StringNull()
	if t := res.Get(path + "advanced.ipMtu.optionType"); t.Exists() {
		va := res.Get(path + "advanced.ipMtu.value")
		if t.String() == "variable" {
			data.AdvancedIpMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedIpMtu = types.Int64Value(va.Int())
		}
	}
	data.AdvancedInterfaceMtu = types.Int64Null()
	data.AdvancedInterfaceMtuVariable = types.StringNull()
	if t := res.Get(path + "advanced.intrfMtu.optionType"); t.Exists() {
		va := res.Get(path + "advanced.intrfMtu.value")
		if t.String() == "variable" {
			data.AdvancedInterfaceMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedInterfaceMtu = types.Int64Value(va.Int())
		}
	}
	data.AdvancedTcpMss = types.Int64Null()
	data.AdvancedTcpMssVariable = types.StringNull()
	if t := res.Get(path + "advanced.tcpMss.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tcpMss.value")
		if t.String() == "variable" {
			data.AdvancedTcpMssVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedTcpMss = types.Int64Value(va.Int())
		}
	}
	data.AdvancedSpeed = types.StringNull()
	data.AdvancedSpeedVariable = types.StringNull()
	if t := res.Get(path + "advanced.speed.optionType"); t.Exists() {
		va := res.Get(path + "advanced.speed.value")
		if t.String() == "variable" {
			data.AdvancedSpeedVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedSpeed = types.StringValue(va.String())
		}
	}
	data.AdvancedArpTimeout = types.Int64Null()
	data.AdvancedArpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "advanced.arpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "advanced.arpTimeout.value")
		if t.String() == "variable" {
			data.AdvancedArpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedArpTimeout = types.Int64Value(va.Int())
		}
	}
	data.AdvancedAutonegotiate = types.BoolNull()
	data.AdvancedAutonegotiateVariable = types.StringNull()
	if t := res.Get(path + "advanced.autonegotiate.optionType"); t.Exists() {
		va := res.Get(path + "advanced.autonegotiate.value")
		if t.String() == "variable" {
			data.AdvancedAutonegotiateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedAutonegotiate = types.BoolValue(va.Bool())
		}
	}
	data.AdvancedMediaType = types.StringNull()
	data.AdvancedMediaTypeVariable = types.StringNull()
	if t := res.Get(path + "advanced.mediaType.optionType"); t.Exists() {
		va := res.Get(path + "advanced.mediaType.value")
		if t.String() == "variable" {
			data.AdvancedMediaTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedMediaType = types.StringValue(va.String())
		}
	}
	data.AdvancedTlocExtension = types.StringNull()
	data.AdvancedTlocExtensionVariable = types.StringNull()
	if t := res.Get(path + "advanced.tlocExtension.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tlocExtension.value")
		if t.String() == "variable" {
			data.AdvancedTlocExtensionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedTlocExtension = types.StringValue(va.String())
		}
	}
	data.AdvancedGreTunnelSourceIp = types.StringNull()
	data.AdvancedGreTunnelSourceIpVariable = types.StringNull()
	if t := res.Get(path + "advanced.tlocExtensionGreFrom.sourceIp.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tlocExtensionGreFrom.sourceIp.value")
		if t.String() == "variable" {
			data.AdvancedGreTunnelSourceIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedGreTunnelSourceIp = types.StringValue(va.String())
		}
	}
	data.AdvancedXconnect = types.StringNull()
	data.AdvancedXconnectVariable = types.StringNull()
	if t := res.Get(path + "advanced.tlocExtensionGreFrom.xconnect.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tlocExtensionGreFrom.xconnect.value")
		if t.String() == "variable" {
			data.AdvancedXconnectVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedXconnect = types.StringValue(va.String())
		}
	}
	data.AdvancedLoadInterval = types.Int64Null()
	data.AdvancedLoadIntervalVariable = types.StringNull()
	if t := res.Get(path + "advanced.loadInterval.optionType"); t.Exists() {
		va := res.Get(path + "advanced.loadInterval.value")
		if t.String() == "variable" {
			data.AdvancedLoadIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedLoadInterval = types.Int64Value(va.Int())
		}
	}
	data.AdvancedTracker = types.StringNull()
	data.AdvancedTrackerVariable = types.StringNull()
	if t := res.Get(path + "advanced.tracker.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tracker.value")
		if t.String() == "variable" {
			data.AdvancedTrackerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedTracker = types.StringValue(va.String())
		}
	}
	data.AdvancedIpDirectedBroadcast = types.BoolNull()
	data.AdvancedIpDirectedBroadcastVariable = types.StringNull()
	if t := res.Get(path + "advanced.ipDirectedBroadcast.optionType"); t.Exists() {
		va := res.Get(path + "advanced.ipDirectedBroadcast.value")
		if t.String() == "variable" {
			data.AdvancedIpDirectedBroadcastVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvancedIpDirectedBroadcast = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *TransportWANVPNInterfaceEthernet) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.TransportWanVpnProfileParcelId.IsNull() {
		return false
	}
	if !data.Shutdown.IsNull() {
		return false
	}
	if !data.ShutdownVariable.IsNull() {
		return false
	}
	if !data.InterfaceName.IsNull() {
		return false
	}
	if !data.InterfaceNameVariable.IsNull() {
		return false
	}
	if !data.ConfigDescription.IsNull() {
		return false
	}
	if !data.ConfigDescriptionVariable.IsNull() {
		return false
	}
	if !data.Ipv4DhcpDistance.IsNull() {
		return false
	}
	if !data.Ipv4DhcpDistanceVariable.IsNull() {
		return false
	}
	if !data.Ipv4Address.IsNull() {
		return false
	}
	if !data.Ipv4AddressVariable.IsNull() {
		return false
	}
	if !data.Ipv4SubnetMask.IsNull() {
		return false
	}
	if !data.Ipv4SubnetMaskVariable.IsNull() {
		return false
	}
	if len(data.Ipv4SecondaryAddresses) > 0 {
		return false
	}
	if !data.Ipv4DhcpHelper.IsNull() {
		return false
	}
	if !data.Ipv4DhcpHelperVariable.IsNull() {
		return false
	}
	if !data.EnableDhcpv6.IsNull() {
		return false
	}
	if len(data.Ipv6DhcpSecondaryAddress) > 0 {
		return false
	}
	if !data.Ipv6Address.IsNull() {
		return false
	}
	if !data.Ipv6AddressVariable.IsNull() {
		return false
	}
	if len(data.Ipv6SecondaryAddress) > 0 {
		return false
	}
	if !data.IperfServer.IsNull() {
		return false
	}
	if !data.IperfServerVariable.IsNull() {
		return false
	}
	if !data.BlockNonSourceIp.IsNull() {
		return false
	}
	if !data.BlockNonSourceIpVariable.IsNull() {
		return false
	}
	if !data.ServiceProvider.IsNull() {
		return false
	}
	if !data.ServiceProviderVariable.IsNull() {
		return false
	}
	if !data.BandwidthUpstream.IsNull() {
		return false
	}
	if !data.BandwidthUpstreamVariable.IsNull() {
		return false
	}
	if !data.BandwidthDownstream.IsNull() {
		return false
	}
	if !data.BandwidthDownstreamVariable.IsNull() {
		return false
	}
	if !data.AutoDetectBandwidth.IsNull() {
		return false
	}
	if !data.AutoDetectBandwidthVariable.IsNull() {
		return false
	}
	if !data.TunnelInterface.IsNull() {
		return false
	}
	if !data.PerTunnelQos.IsNull() {
		return false
	}
	if !data.PerTunnelQosVariable.IsNull() {
		return false
	}
	if !data.TunnelQosMode.IsNull() {
		return false
	}
	if !data.TunnelQosModeVariable.IsNull() {
		return false
	}
	if !data.TunnelBandwidthPercent.IsNull() {
		return false
	}
	if !data.TunnelBandwidthPercentVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceBindLoopbackTunnel.IsNull() {
		return false
	}
	if !data.TunnelInterfaceBindLoopbackTunnelVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceCarrier.IsNull() {
		return false
	}
	if !data.TunnelInterfaceCarrierVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceColor.IsNull() {
		return false
	}
	if !data.TunnelInterfaceColorVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceHelloInterval.IsNull() {
		return false
	}
	if !data.TunnelInterfaceHelloIntervalVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceHelloTolerance.IsNull() {
		return false
	}
	if !data.TunnelInterfaceHelloToleranceVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceLastResortCircuit.IsNull() {
		return false
	}
	if !data.TunnelInterfaceLastResortCircuitVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceGreTunnelDestinationIp.IsNull() {
		return false
	}
	if !data.TunnelInterfaceGreTunnelDestinationIpVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceColorRestrict.IsNull() {
		return false
	}
	if !data.TunnelInterfaceColorRestrictVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceGroups.IsNull() {
		return false
	}
	if !data.TunnelInterfaceGroupsVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceBorder.IsNull() {
		return false
	}
	if !data.TunnelInterfaceBorderVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceMaxControlConnections.IsNull() {
		return false
	}
	if !data.TunnelInterfaceMaxControlConnectionsVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceNatRefreshInterval.IsNull() {
		return false
	}
	if !data.TunnelInterfaceNatRefreshIntervalVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceVbondAsStunServer.IsNull() {
		return false
	}
	if !data.TunnelInterfaceVbondAsStunServerVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceExcludeControllerGroupList.IsNull() {
		return false
	}
	if !data.TunnelInterfaceExcludeControllerGroupListVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceVmanageConnectionPreference.IsNull() {
		return false
	}
	if !data.TunnelInterfaceVmanageConnectionPreferenceVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfacePortHop.IsNull() {
		return false
	}
	if !data.TunnelInterfacePortHopVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceLowBandwidthLink.IsNull() {
		return false
	}
	if !data.TunnelInterfaceLowBandwidthLinkVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceTunnelTcpMss.IsNull() {
		return false
	}
	if !data.TunnelInterfaceTunnelTcpMssVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceClearDontFragment.IsNull() {
		return false
	}
	if !data.TunnelInterfaceClearDontFragmentVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceCtsSgtPropagation.IsNull() {
		return false
	}
	if !data.TunnelInterfaceCtsSgtPropagationVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceNetworkBroadcast.IsNull() {
		return false
	}
	if !data.TunnelInterfaceNetworkBroadcastVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowAll.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowAllVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowBgp.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowBgpVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowDhcp.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowDhcpVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowNtp.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowNtpVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowSsh.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowSshVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowDbs.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowDbsVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowIcmp.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowIcmpVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowHttps.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowHttpsVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowOspf.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowOspfVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowStun.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowStunVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowSnmp.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowSnmpVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowNetconf.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowNetconfVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowBfd.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowBfdVariable.IsNull() {
		return false
	}
	if len(data.TunnelInterfaceEncapsulations) > 0 {
		return false
	}
	if !data.NatIpv4.IsNull() {
		return false
	}
	if !data.NatIpv4Variable.IsNull() {
		return false
	}
	if !data.NatType.IsNull() {
		return false
	}
	if !data.NatTypeVariable.IsNull() {
		return false
	}
	if !data.NatRangeStart.IsNull() {
		return false
	}
	if !data.NatRangeStartVariable.IsNull() {
		return false
	}
	if !data.NatRangeEnd.IsNull() {
		return false
	}
	if !data.NatRangeEndVariable.IsNull() {
		return false
	}
	if !data.NatPrefixLength.IsNull() {
		return false
	}
	if !data.NatPrefixLengthVariable.IsNull() {
		return false
	}
	if !data.NatOverload.IsNull() {
		return false
	}
	if !data.NatOverloadVariable.IsNull() {
		return false
	}
	if !data.NatLoopback.IsNull() {
		return false
	}
	if !data.NatLoopbackVariable.IsNull() {
		return false
	}
	if !data.NatUdpTimeout.IsNull() {
		return false
	}
	if !data.NatUdpTimeoutVariable.IsNull() {
		return false
	}
	if !data.NatTcpTimeout.IsNull() {
		return false
	}
	if !data.NatTcpTimeoutVariable.IsNull() {
		return false
	}
	if len(data.NewStaticNats) > 0 {
		return false
	}
	if !data.NatIpv6.IsNull() {
		return false
	}
	if !data.NatIpv6Variable.IsNull() {
		return false
	}
	if !data.Nat64.IsNull() {
		return false
	}
	if !data.Nat66.IsNull() {
		return false
	}
	if len(data.StaticNat66) > 0 {
		return false
	}
	if !data.AdaptiveQos.IsNull() {
		return false
	}
	if !data.QosAdaptivePeriod.IsNull() {
		return false
	}
	if !data.QosAdaptivePeriodVariable.IsNull() {
		return false
	}
	if !data.QosAdaptiveBandwidthUpstream.IsNull() {
		return false
	}
	if !data.QosAdaptiveMinUpstream.IsNull() {
		return false
	}
	if !data.QosAdaptiveMinUpstreamVariable.IsNull() {
		return false
	}
	if !data.QosAdaptiveMaxUpstream.IsNull() {
		return false
	}
	if !data.QosAdaptiveMaxUpstreamVariable.IsNull() {
		return false
	}
	if !data.QosAdaptiveDefaultUpstream.IsNull() {
		return false
	}
	if !data.QosAdaptiveDefaultUpstreamVariable.IsNull() {
		return false
	}
	if !data.QosAdaptiveBandwidthDownstream.IsNull() {
		return false
	}
	if !data.QosAdaptiveMinDownstream.IsNull() {
		return false
	}
	if !data.QosAdaptiveMinDownstreamVariable.IsNull() {
		return false
	}
	if !data.QosAdaptiveMaxDownstream.IsNull() {
		return false
	}
	if !data.QosAdaptiveMaxDownstreamVariable.IsNull() {
		return false
	}
	if !data.QosAdaptiveDefaultDownstream.IsNull() {
		return false
	}
	if !data.QosAdaptiveDefaultDownstreamVariable.IsNull() {
		return false
	}
	if !data.QosShapingRate.IsNull() {
		return false
	}
	if !data.QosShapingRateVariable.IsNull() {
		return false
	}
	if len(data.Arps) > 0 {
		return false
	}
	if !data.AdvancedIcmpRedirectDisable.IsNull() {
		return false
	}
	if !data.AdvancedIcmpRedirectDisableVariable.IsNull() {
		return false
	}
	if !data.AdvancedDuplex.IsNull() {
		return false
	}
	if !data.AdvancedDuplexVariable.IsNull() {
		return false
	}
	if !data.AdvancedMacAddress.IsNull() {
		return false
	}
	if !data.AdvancedMacAddressVariable.IsNull() {
		return false
	}
	if !data.AdvancedIpMtu.IsNull() {
		return false
	}
	if !data.AdvancedIpMtuVariable.IsNull() {
		return false
	}
	if !data.AdvancedInterfaceMtu.IsNull() {
		return false
	}
	if !data.AdvancedInterfaceMtuVariable.IsNull() {
		return false
	}
	if !data.AdvancedTcpMss.IsNull() {
		return false
	}
	if !data.AdvancedTcpMssVariable.IsNull() {
		return false
	}
	if !data.AdvancedSpeed.IsNull() {
		return false
	}
	if !data.AdvancedSpeedVariable.IsNull() {
		return false
	}
	if !data.AdvancedArpTimeout.IsNull() {
		return false
	}
	if !data.AdvancedArpTimeoutVariable.IsNull() {
		return false
	}
	if !data.AdvancedAutonegotiate.IsNull() {
		return false
	}
	if !data.AdvancedAutonegotiateVariable.IsNull() {
		return false
	}
	if !data.AdvancedMediaType.IsNull() {
		return false
	}
	if !data.AdvancedMediaTypeVariable.IsNull() {
		return false
	}
	if !data.AdvancedTlocExtension.IsNull() {
		return false
	}
	if !data.AdvancedTlocExtensionVariable.IsNull() {
		return false
	}
	if !data.AdvancedGreTunnelSourceIp.IsNull() {
		return false
	}
	if !data.AdvancedGreTunnelSourceIpVariable.IsNull() {
		return false
	}
	if !data.AdvancedXconnect.IsNull() {
		return false
	}
	if !data.AdvancedXconnectVariable.IsNull() {
		return false
	}
	if !data.AdvancedLoadInterval.IsNull() {
		return false
	}
	if !data.AdvancedLoadIntervalVariable.IsNull() {
		return false
	}
	if !data.AdvancedTracker.IsNull() {
		return false
	}
	if !data.AdvancedTrackerVariable.IsNull() {
		return false
	}
	if !data.AdvancedIpDirectedBroadcast.IsNull() {
		return false
	}
	if !data.AdvancedIpDirectedBroadcastVariable.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
