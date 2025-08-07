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
	"strings"

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
	TransportWanVpnFeatureId                           types.String                                                    `tfsdk:"transport_wan_vpn_feature_id"`
	Shutdown                                           types.Bool                                                      `tfsdk:"shutdown"`
	ShutdownVariable                                   types.String                                                    `tfsdk:"shutdown_variable"`
	InterfaceName                                      types.String                                                    `tfsdk:"interface_name"`
	InterfaceNameVariable                              types.String                                                    `tfsdk:"interface_name_variable"`
	InterfaceDescription                               types.String                                                    `tfsdk:"interface_description"`
	InterfaceDescriptionVariable                       types.String                                                    `tfsdk:"interface_description_variable"`
	Ipv4ConfigurationType                              types.String                                                    `tfsdk:"ipv4_configuration_type"`
	Ipv4DhcpDistance                                   types.Int64                                                     `tfsdk:"ipv4_dhcp_distance"`
	Ipv4DhcpDistanceVariable                           types.String                                                    `tfsdk:"ipv4_dhcp_distance_variable"`
	Ipv4Address                                        types.String                                                    `tfsdk:"ipv4_address"`
	Ipv4AddressVariable                                types.String                                                    `tfsdk:"ipv4_address_variable"`
	Ipv4SubnetMask                                     types.String                                                    `tfsdk:"ipv4_subnet_mask"`
	Ipv4SubnetMaskVariable                             types.String                                                    `tfsdk:"ipv4_subnet_mask_variable"`
	Ipv4SecondaryAddresses                             []TransportWANVPNInterfaceEthernetIpv4SecondaryAddresses        `tfsdk:"ipv4_secondary_addresses"`
	Ipv4DhcpHelper                                     types.Set                                                       `tfsdk:"ipv4_dhcp_helper"`
	Ipv4DhcpHelperVariable                             types.String                                                    `tfsdk:"ipv4_dhcp_helper_variable"`
	Ipv6ConfigurationType                              types.String                                                    `tfsdk:"ipv6_configuration_type"`
	EnableDhcpv6                                       types.Bool                                                      `tfsdk:"enable_dhcpv6"`
	Ipv6DhcpSecondaryAddress                           []TransportWANVPNInterfaceEthernetIpv6DhcpSecondaryAddress      `tfsdk:"ipv6_dhcp_secondary_address"`
	Ipv6Address                                        types.String                                                    `tfsdk:"ipv6_address"`
	Ipv6AddressVariable                                types.String                                                    `tfsdk:"ipv6_address_variable"`
	Ipv6SecondaryAddresses                             []TransportWANVPNInterfaceEthernetIpv6SecondaryAddresses        `tfsdk:"ipv6_secondary_addresses"`
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
	TunnelInterfaceAllowDns                            types.Bool                                                      `tfsdk:"tunnel_interface_allow_dns"`
	TunnelInterfaceAllowDnsVariable                    types.String                                                    `tfsdk:"tunnel_interface_allow_dns_variable"`
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
	QosAdaptive                                        types.Bool                                                      `tfsdk:"qos_adaptive"`
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
	AclIpv4EgressFeatureId                             types.String                                                    `tfsdk:"acl_ipv4_egress_feature_id"`
	AclIpv4IngressFeatureId                            types.String                                                    `tfsdk:"acl_ipv4_ingress_feature_id"`
	AclIpv6EgressFeatureId                             types.String                                                    `tfsdk:"acl_ipv6_egress_feature_id"`
	AclIpv6IngressFeatureId                            types.String                                                    `tfsdk:"acl_ipv6_ingress_feature_id"`
	Arps                                               []TransportWANVPNInterfaceEthernetArps                          `tfsdk:"arps"`
	IcmpRedirectDisable                                types.Bool                                                      `tfsdk:"icmp_redirect_disable"`
	IcmpRedirectDisableVariable                        types.String                                                    `tfsdk:"icmp_redirect_disable_variable"`
	Duplex                                             types.String                                                    `tfsdk:"duplex"`
	DuplexVariable                                     types.String                                                    `tfsdk:"duplex_variable"`
	MacAddress                                         types.String                                                    `tfsdk:"mac_address"`
	MacAddressVariable                                 types.String                                                    `tfsdk:"mac_address_variable"`
	IpMtu                                              types.Int64                                                     `tfsdk:"ip_mtu"`
	IpMtuVariable                                      types.String                                                    `tfsdk:"ip_mtu_variable"`
	InterfaceMtu                                       types.Int64                                                     `tfsdk:"interface_mtu"`
	InterfaceMtuVariable                               types.String                                                    `tfsdk:"interface_mtu_variable"`
	TcpMss                                             types.Int64                                                     `tfsdk:"tcp_mss"`
	TcpMssVariable                                     types.String                                                    `tfsdk:"tcp_mss_variable"`
	Speed                                              types.String                                                    `tfsdk:"speed"`
	SpeedVariable                                      types.String                                                    `tfsdk:"speed_variable"`
	ArpTimeout                                         types.Int64                                                     `tfsdk:"arp_timeout"`
	ArpTimeoutVariable                                 types.String                                                    `tfsdk:"arp_timeout_variable"`
	Autonegotiate                                      types.Bool                                                      `tfsdk:"autonegotiate"`
	AutonegotiateVariable                              types.String                                                    `tfsdk:"autonegotiate_variable"`
	MediaType                                          types.String                                                    `tfsdk:"media_type"`
	MediaTypeVariable                                  types.String                                                    `tfsdk:"media_type_variable"`
	TlocExtension                                      types.String                                                    `tfsdk:"tloc_extension"`
	TlocExtensionVariable                              types.String                                                    `tfsdk:"tloc_extension_variable"`
	GreTunnelSourceIp                                  types.String                                                    `tfsdk:"gre_tunnel_source_ip"`
	GreTunnelSourceIpVariable                          types.String                                                    `tfsdk:"gre_tunnel_source_ip_variable"`
	Xconnect                                           types.String                                                    `tfsdk:"xconnect"`
	XconnectVariable                                   types.String                                                    `tfsdk:"xconnect_variable"`
	LoadInterval                                       types.Int64                                                     `tfsdk:"load_interval"`
	LoadIntervalVariable                               types.String                                                    `tfsdk:"load_interval_variable"`
	Tracker                                            types.String                                                    `tfsdk:"tracker"`
	TrackerVariable                                    types.String                                                    `tfsdk:"tracker_variable"`
	IpDirectedBroadcast                                types.Bool                                                      `tfsdk:"ip_directed_broadcast"`
	IpDirectedBroadcastVariable                        types.String                                                    `tfsdk:"ip_directed_broadcast_variable"`
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

type TransportWANVPNInterfaceEthernetIpv6SecondaryAddresses struct {
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
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/wan/vpn/%s/interface/ethernet", url.QueryEscape(data.FeatureProfileId.ValueString()), url.QueryEscape(data.TransportWanVpnFeatureId.ValueString()))
}

// End of section. //template:end getPath

func (data TransportWANVPNInterfaceEthernet) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.ShutdownVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"shutdown.optionType", "variable")
			body, _ = sjson.Set(body, path+"shutdown.value", data.ShutdownVariable.ValueString())
		}
	} else if data.Shutdown.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"shutdown.optionType", "default")
			body, _ = sjson.Set(body, path+"shutdown.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"shutdown.optionType", "global")
			body, _ = sjson.Set(body, path+"shutdown.value", data.Shutdown.ValueBool())
		}
	}

	if !data.InterfaceNameVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"interfaceName.optionType", "variable")
			body, _ = sjson.Set(body, path+"interfaceName.value", data.InterfaceNameVariable.ValueString())
		}
	} else if !data.InterfaceName.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"interfaceName.optionType", "global")
			body, _ = sjson.Set(body, path+"interfaceName.value", data.InterfaceName.ValueString())
		}
	}

	if !data.InterfaceDescriptionVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"description.optionType", "variable")
			body, _ = sjson.Set(body, path+"description.value", data.InterfaceDescriptionVariable.ValueString())
		}
	} else if data.InterfaceDescription.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"description.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"description.optionType", "global")
			body, _ = sjson.Set(body, path+"description.value", data.InterfaceDescription.ValueString())
		}
	}

	if !data.Ipv4DhcpDistanceVariable.IsNull() {
		if true && data.Ipv4ConfigurationType.ValueString() == "dynamic" {
			body, _ = sjson.Set(body, path+"intfIpAddress.dynamic.dynamicDhcpDistance.optionType", "variable")
			body, _ = sjson.Set(body, path+"intfIpAddress.dynamic.dynamicDhcpDistance.value", data.Ipv4DhcpDistanceVariable.ValueString())
		}
	} else if data.Ipv4DhcpDistance.IsNull() {
		if true && data.Ipv4ConfigurationType.ValueString() == "dynamic" {
			body, _ = sjson.Set(body, path+"intfIpAddress.dynamic.dynamicDhcpDistance.optionType", "default")
			body, _ = sjson.Set(body, path+"intfIpAddress.dynamic.dynamicDhcpDistance.value", 1)
		}
	} else {
		if true && data.Ipv4ConfigurationType.ValueString() == "dynamic" {
			body, _ = sjson.Set(body, path+"intfIpAddress.dynamic.dynamicDhcpDistance.optionType", "global")
			body, _ = sjson.Set(body, path+"intfIpAddress.dynamic.dynamicDhcpDistance.value", data.Ipv4DhcpDistance.ValueInt64())
		}
	}

	if !data.Ipv4AddressVariable.IsNull() {
		if true && data.Ipv4ConfigurationType.ValueString() == "static" {
			body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.optionType", "variable")
			body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.value", data.Ipv4AddressVariable.ValueString())
		}
	} else if !data.Ipv4Address.IsNull() {
		if true && data.Ipv4ConfigurationType.ValueString() == "static" {
			body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.optionType", "global")
			body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.value", data.Ipv4Address.ValueString())
		}
	}

	if !data.Ipv4SubnetMaskVariable.IsNull() {
		if true && data.Ipv4ConfigurationType.ValueString() == "static" {
			body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.optionType", "variable")
			body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.value", data.Ipv4SubnetMaskVariable.ValueString())
		}
	} else if !data.Ipv4SubnetMask.IsNull() {
		if true && data.Ipv4ConfigurationType.ValueString() == "static" {
			body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.optionType", "global")
			body, _ = sjson.Set(body, path+"intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.value", data.Ipv4SubnetMask.ValueString())
		}
	}
	if true && data.Ipv4ConfigurationType.ValueString() == "static" {

		for _, item := range data.Ipv4SecondaryAddresses {
			itemBody := ""

			if !item.AddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipAddress.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ipAddress.value", item.AddressVariable.ValueString())
				}
			} else if !item.Address.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipAddress.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ipAddress.value", item.Address.ValueString())
				}
			}

			if !item.SubnetMaskVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "subnetMask.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "subnetMask.value", item.SubnetMaskVariable.ValueString())
				}
			} else if !item.SubnetMask.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "subnetMask.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "subnetMask.value", item.SubnetMask.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"intfIpAddress.static.staticIpV4AddressSecondary.-1", itemBody)
		}
	}

	if !data.Ipv4DhcpHelperVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dhcpHelper.optionType", "variable")
			body, _ = sjson.Set(body, path+"dhcpHelper.value", data.Ipv4DhcpHelperVariable.ValueString())
		}
	} else if data.Ipv4DhcpHelper.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dhcpHelper.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"dhcpHelper.optionType", "global")
			var values []string
			data.Ipv4DhcpHelper.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"dhcpHelper.value", values)
		}
	}
	if !data.EnableDhcpv6.IsNull() {
		if true && data.Ipv6ConfigurationType.ValueString() == "dynamic" {
			body, _ = sjson.Set(body, path+"intfIpV6Address.dynamic.dhcpClient.optionType", "global")
			body, _ = sjson.Set(body, path+"intfIpV6Address.dynamic.dhcpClient.value", data.EnableDhcpv6.ValueBool())
		}
	}
	if true && data.Ipv6ConfigurationType.ValueString() == "dynamic" {

		for _, item := range data.Ipv6DhcpSecondaryAddress {
			itemBody := ""

			if !item.AddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.AddressVariable.ValueString())
				}
			} else if !item.Address.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.Address.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"intfIpV6Address.dynamic.secondaryIpV6Address.-1", itemBody)
		}
	}

	if !data.Ipv6AddressVariable.IsNull() {
		if true && data.Ipv6ConfigurationType.ValueString() == "static" {
			body, _ = sjson.Set(body, path+"intfIpV6Address.static.primaryIpV6Address.address.optionType", "variable")
			body, _ = sjson.Set(body, path+"intfIpV6Address.static.primaryIpV6Address.address.value", data.Ipv6AddressVariable.ValueString())
		}
	} else if !data.Ipv6Address.IsNull() {
		if true && data.Ipv6ConfigurationType.ValueString() == "static" {
			body, _ = sjson.Set(body, path+"intfIpV6Address.static.primaryIpV6Address.address.optionType", "global")
			body, _ = sjson.Set(body, path+"intfIpV6Address.static.primaryIpV6Address.address.value", data.Ipv6Address.ValueString())
		}
	}
	if true && data.Ipv6ConfigurationType.ValueString() == "static" {

		for _, item := range data.Ipv6SecondaryAddresses {
			itemBody := ""

			if !item.AddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.AddressVariable.ValueString())
				}
			} else if !item.Address.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.Address.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"intfIpV6Address.static.secondaryIpV6Address.-1", itemBody)
		}
	}

	if !data.IperfServerVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"iperfServer.optionType", "variable")
			body, _ = sjson.Set(body, path+"iperfServer.value", data.IperfServerVariable.ValueString())
		}
	} else if data.IperfServer.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"iperfServer.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"iperfServer.optionType", "global")
			body, _ = sjson.Set(body, path+"iperfServer.value", data.IperfServer.ValueString())
		}
	}

	if !data.BlockNonSourceIpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"blockNonSourceIp.optionType", "variable")
			body, _ = sjson.Set(body, path+"blockNonSourceIp.value", data.BlockNonSourceIpVariable.ValueString())
		}
	} else if data.BlockNonSourceIp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"blockNonSourceIp.optionType", "default")
			body, _ = sjson.Set(body, path+"blockNonSourceIp.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"blockNonSourceIp.optionType", "global")
			body, _ = sjson.Set(body, path+"blockNonSourceIp.value", data.BlockNonSourceIp.ValueBool())
		}
	}

	if !data.ServiceProviderVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"serviceProvider.optionType", "variable")
			body, _ = sjson.Set(body, path+"serviceProvider.value", data.ServiceProviderVariable.ValueString())
		}
	} else if data.ServiceProvider.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"serviceProvider.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"serviceProvider.optionType", "global")
			body, _ = sjson.Set(body, path+"serviceProvider.value", data.ServiceProvider.ValueString())
		}
	}

	if !data.BandwidthUpstreamVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"bandwidthUpstream.optionType", "variable")
			body, _ = sjson.Set(body, path+"bandwidthUpstream.value", data.BandwidthUpstreamVariable.ValueString())
		}
	} else if data.BandwidthUpstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"bandwidthUpstream.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"bandwidthUpstream.optionType", "global")
			body, _ = sjson.Set(body, path+"bandwidthUpstream.value", data.BandwidthUpstream.ValueInt64())
		}
	}

	if !data.BandwidthDownstreamVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"bandwidthDownstream.optionType", "variable")
			body, _ = sjson.Set(body, path+"bandwidthDownstream.value", data.BandwidthDownstreamVariable.ValueString())
		}
	} else if data.BandwidthDownstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"bandwidthDownstream.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"bandwidthDownstream.optionType", "global")
			body, _ = sjson.Set(body, path+"bandwidthDownstream.value", data.BandwidthDownstream.ValueInt64())
		}
	}

	if !data.AutoDetectBandwidthVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"autoDetectBandwidth.optionType", "variable")
			body, _ = sjson.Set(body, path+"autoDetectBandwidth.value", data.AutoDetectBandwidthVariable.ValueString())
		}
	} else if data.AutoDetectBandwidth.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"autoDetectBandwidth.optionType", "default")
			body, _ = sjson.Set(body, path+"autoDetectBandwidth.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"autoDetectBandwidth.optionType", "global")
			body, _ = sjson.Set(body, path+"autoDetectBandwidth.value", data.AutoDetectBandwidth.ValueBool())
		}
	}
	if data.TunnelInterface.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelInterface.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelInterface.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelInterface.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelInterface.value", data.TunnelInterface.ValueBool())
		}
	}

	if !data.PerTunnelQosVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.value", data.PerTunnelQosVariable.ValueString())
		}
	} else if data.PerTunnelQos.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.value", false)
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.value", data.PerTunnelQos.ValueBool())
		}
	}

	if !data.TunnelQosModeVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.mode.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.mode.value", data.TunnelQosModeVariable.ValueString())
		}
	} else if !data.TunnelQosMode.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.mode.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.mode.value", data.TunnelQosMode.ValueString())
		}
	}

	if !data.TunnelBandwidthPercentVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.value", data.TunnelBandwidthPercentVariable.ValueString())
		}
	} else if data.TunnelBandwidthPercent.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.value", 50)
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.value", data.TunnelBandwidthPercent.ValueInt64())
		}
	}

	if !data.TunnelInterfaceBindLoopbackTunnelVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.bind.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.bind.value", data.TunnelInterfaceBindLoopbackTunnelVariable.ValueString())
		}
	} else if data.TunnelInterfaceBindLoopbackTunnel.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.bind.optionType", "default")

		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.bind.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.bind.value", data.TunnelInterfaceBindLoopbackTunnel.ValueString())
		}
	}

	if !data.TunnelInterfaceCarrierVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.carrier.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.carrier.value", data.TunnelInterfaceCarrierVariable.ValueString())
		}
	} else if data.TunnelInterfaceCarrier.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.carrier.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.carrier.value", "default")
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.carrier.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.carrier.value", data.TunnelInterfaceCarrier.ValueString())
		}
	}

	if !data.TunnelInterfaceColorVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.color.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.color.value", data.TunnelInterfaceColorVariable.ValueString())
		}
	} else if data.TunnelInterfaceColor.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.color.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.color.value", "mpls")
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.color.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.color.value", data.TunnelInterfaceColor.ValueString())
		}
	}

	if !data.TunnelInterfaceHelloIntervalVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.helloInterval.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.helloInterval.value", data.TunnelInterfaceHelloIntervalVariable.ValueString())
		}
	} else if data.TunnelInterfaceHelloInterval.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.helloInterval.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.helloInterval.value", 1000)
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.helloInterval.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.helloInterval.value", data.TunnelInterfaceHelloInterval.ValueInt64())
		}
	}

	if !data.TunnelInterfaceHelloToleranceVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.helloTolerance.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.helloTolerance.value", data.TunnelInterfaceHelloToleranceVariable.ValueString())
		}
	} else if data.TunnelInterfaceHelloTolerance.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.helloTolerance.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.helloTolerance.value", 12)
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.helloTolerance.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.helloTolerance.value", data.TunnelInterfaceHelloTolerance.ValueInt64())
		}
	}

	if !data.TunnelInterfaceLastResortCircuitVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.value", data.TunnelInterfaceLastResortCircuitVariable.ValueString())
		}
	} else if data.TunnelInterfaceLastResortCircuit.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.value", false)
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.value", data.TunnelInterfaceLastResortCircuit.ValueBool())
		}
	}

	if !data.TunnelInterfaceGreTunnelDestinationIpVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.tlocExtensionGreTo.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.tlocExtensionGreTo.value", data.TunnelInterfaceGreTunnelDestinationIpVariable.ValueString())
		}
	} else if data.TunnelInterfaceGreTunnelDestinationIp.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.tlocExtensionGreTo.optionType", "default")

		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.tlocExtensionGreTo.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.tlocExtensionGreTo.value", data.TunnelInterfaceGreTunnelDestinationIp.ValueString())
		}
	}

	if !data.TunnelInterfaceColorRestrictVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.restrict.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.restrict.value", data.TunnelInterfaceColorRestrictVariable.ValueString())
		}
	} else if data.TunnelInterfaceColorRestrict.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.restrict.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.restrict.value", false)
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.restrict.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.restrict.value", data.TunnelInterfaceColorRestrict.ValueBool())
		}
	}

	if !data.TunnelInterfaceGroupsVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.group.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.group.value", data.TunnelInterfaceGroupsVariable.ValueString())
		}
	} else if data.TunnelInterfaceGroups.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.group.optionType", "default")

		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.group.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.group.value", data.TunnelInterfaceGroups.ValueInt64())
		}
	}

	if !data.TunnelInterfaceBorderVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.border.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.border.value", data.TunnelInterfaceBorderVariable.ValueString())
		}
	} else if data.TunnelInterfaceBorder.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.border.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.border.value", false)
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.border.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.border.value", data.TunnelInterfaceBorder.ValueBool())
		}
	}

	if !data.TunnelInterfaceMaxControlConnectionsVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.maxControlConnections.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.maxControlConnections.value", data.TunnelInterfaceMaxControlConnectionsVariable.ValueString())
		}
	} else if data.TunnelInterfaceMaxControlConnections.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.maxControlConnections.optionType", "default")

		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.maxControlConnections.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.maxControlConnections.value", data.TunnelInterfaceMaxControlConnections.ValueInt64())
		}
	}

	if !data.TunnelInterfaceNatRefreshIntervalVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.value", data.TunnelInterfaceNatRefreshIntervalVariable.ValueString())
		}
	} else if data.TunnelInterfaceNatRefreshInterval.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.value", 5)
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.value", data.TunnelInterfaceNatRefreshInterval.ValueInt64())
		}
	}

	if !data.TunnelInterfaceVbondAsStunServerVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.vBondAsStunServer.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.vBondAsStunServer.value", data.TunnelInterfaceVbondAsStunServerVariable.ValueString())
		}
	} else if data.TunnelInterfaceVbondAsStunServer.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.vBondAsStunServer.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.vBondAsStunServer.value", false)
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.vBondAsStunServer.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.vBondAsStunServer.value", data.TunnelInterfaceVbondAsStunServer.ValueBool())
		}
	}

	if !data.TunnelInterfaceExcludeControllerGroupListVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.excludeControllerGroupList.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.excludeControllerGroupList.value", data.TunnelInterfaceExcludeControllerGroupListVariable.ValueString())
		}
	} else if data.TunnelInterfaceExcludeControllerGroupList.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.excludeControllerGroupList.optionType", "default")

		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.excludeControllerGroupList.optionType", "global")
			var values []int64
			data.TunnelInterfaceExcludeControllerGroupList.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"tunnel.excludeControllerGroupList.value", values)
		}
	}

	if !data.TunnelInterfaceVmanageConnectionPreferenceVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.vManageConnectionPreference.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.vManageConnectionPreference.value", data.TunnelInterfaceVmanageConnectionPreferenceVariable.ValueString())
		}
	} else if data.TunnelInterfaceVmanageConnectionPreference.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.vManageConnectionPreference.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.vManageConnectionPreference.value", 5)
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.vManageConnectionPreference.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.vManageConnectionPreference.value", data.TunnelInterfaceVmanageConnectionPreference.ValueInt64())
		}
	}

	if !data.TunnelInterfacePortHopVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.portHop.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.portHop.value", data.TunnelInterfacePortHopVariable.ValueString())
		}
	} else if data.TunnelInterfacePortHop.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.portHop.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.portHop.value", true)
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.portHop.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.portHop.value", data.TunnelInterfacePortHop.ValueBool())
		}
	}

	if !data.TunnelInterfaceLowBandwidthLinkVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.value", data.TunnelInterfaceLowBandwidthLinkVariable.ValueString())
		}
	} else if data.TunnelInterfaceLowBandwidthLink.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.value", false)
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.value", data.TunnelInterfaceLowBandwidthLink.ValueBool())
		}
	}

	if !data.TunnelInterfaceTunnelTcpMssVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.tunnelTcpMss.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.tunnelTcpMss.value", data.TunnelInterfaceTunnelTcpMssVariable.ValueString())
		}
	} else if data.TunnelInterfaceTunnelTcpMss.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.tunnelTcpMss.optionType", "default")

		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.tunnelTcpMss.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.tunnelTcpMss.value", data.TunnelInterfaceTunnelTcpMss.ValueInt64())
		}
	}

	if !data.TunnelInterfaceClearDontFragmentVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.value", data.TunnelInterfaceClearDontFragmentVariable.ValueString())
		}
	} else if data.TunnelInterfaceClearDontFragment.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.value", false)
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.value", data.TunnelInterfaceClearDontFragment.ValueBool())
		}
	}

	if !data.TunnelInterfaceCtsSgtPropagationVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.ctsSgtPropagation.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.ctsSgtPropagation.value", data.TunnelInterfaceCtsSgtPropagationVariable.ValueString())
		}
	} else if data.TunnelInterfaceCtsSgtPropagation.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.ctsSgtPropagation.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.ctsSgtPropagation.value", false)
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.ctsSgtPropagation.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.ctsSgtPropagation.value", data.TunnelInterfaceCtsSgtPropagation.ValueBool())
		}
	}

	if !data.TunnelInterfaceNetworkBroadcastVariable.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.value", data.TunnelInterfaceNetworkBroadcastVariable.ValueString())
		}
	} else if data.TunnelInterfaceNetworkBroadcast.IsNull() {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.value", false)
		}
	} else {
		if true && data.TunnelInterface.ValueBool() == true {
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.value", data.TunnelInterfaceNetworkBroadcast.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowAllVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.all.optionType", "variable")
			body, _ = sjson.Set(body, path+"allowService.all.value", data.TunnelInterfaceAllowAllVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowAll.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.all.optionType", "default")
			body, _ = sjson.Set(body, path+"allowService.all.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"allowService.all.optionType", "global")
			body, _ = sjson.Set(body, path+"allowService.all.value", data.TunnelInterfaceAllowAll.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowBgpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.bgp.optionType", "variable")
			body, _ = sjson.Set(body, path+"allowService.bgp.value", data.TunnelInterfaceAllowBgpVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowBgp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.bgp.optionType", "default")
			body, _ = sjson.Set(body, path+"allowService.bgp.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"allowService.bgp.optionType", "global")
			body, _ = sjson.Set(body, path+"allowService.bgp.value", data.TunnelInterfaceAllowBgp.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowDhcpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.dhcp.optionType", "variable")
			body, _ = sjson.Set(body, path+"allowService.dhcp.value", data.TunnelInterfaceAllowDhcpVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowDhcp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.dhcp.optionType", "default")
			body, _ = sjson.Set(body, path+"allowService.dhcp.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"allowService.dhcp.optionType", "global")
			body, _ = sjson.Set(body, path+"allowService.dhcp.value", data.TunnelInterfaceAllowDhcp.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowNtpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.ntp.optionType", "variable")
			body, _ = sjson.Set(body, path+"allowService.ntp.value", data.TunnelInterfaceAllowNtpVariable.ValueString())
		}
	} else if !data.TunnelInterfaceAllowNtp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.ntp.optionType", "global")
			body, _ = sjson.Set(body, path+"allowService.ntp.value", data.TunnelInterfaceAllowNtp.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowSshVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.ssh.optionType", "variable")
			body, _ = sjson.Set(body, path+"allowService.ssh.value", data.TunnelInterfaceAllowSshVariable.ValueString())
		}
	} else if !data.TunnelInterfaceAllowSsh.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.ssh.optionType", "global")
			body, _ = sjson.Set(body, path+"allowService.ssh.value", data.TunnelInterfaceAllowSsh.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowDnsVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.dns.optionType", "variable")
			body, _ = sjson.Set(body, path+"allowService.dns.value", data.TunnelInterfaceAllowDnsVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowDns.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.dns.optionType", "default")
			body, _ = sjson.Set(body, path+"allowService.dns.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"allowService.dns.optionType", "global")
			body, _ = sjson.Set(body, path+"allowService.dns.value", data.TunnelInterfaceAllowDns.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowIcmpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.icmp.optionType", "variable")
			body, _ = sjson.Set(body, path+"allowService.icmp.value", data.TunnelInterfaceAllowIcmpVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowIcmp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.icmp.optionType", "default")
			body, _ = sjson.Set(body, path+"allowService.icmp.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"allowService.icmp.optionType", "global")
			body, _ = sjson.Set(body, path+"allowService.icmp.value", data.TunnelInterfaceAllowIcmp.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowHttpsVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.https.optionType", "variable")
			body, _ = sjson.Set(body, path+"allowService.https.value", data.TunnelInterfaceAllowHttpsVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowHttps.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.https.optionType", "default")
			body, _ = sjson.Set(body, path+"allowService.https.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"allowService.https.optionType", "global")
			body, _ = sjson.Set(body, path+"allowService.https.value", data.TunnelInterfaceAllowHttps.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowOspfVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.ospf.optionType", "variable")
			body, _ = sjson.Set(body, path+"allowService.ospf.value", data.TunnelInterfaceAllowOspfVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowOspf.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.ospf.optionType", "default")
			body, _ = sjson.Set(body, path+"allowService.ospf.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"allowService.ospf.optionType", "global")
			body, _ = sjson.Set(body, path+"allowService.ospf.value", data.TunnelInterfaceAllowOspf.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowStunVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.stun.optionType", "variable")
			body, _ = sjson.Set(body, path+"allowService.stun.value", data.TunnelInterfaceAllowStunVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowStun.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.stun.optionType", "default")
			body, _ = sjson.Set(body, path+"allowService.stun.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"allowService.stun.optionType", "global")
			body, _ = sjson.Set(body, path+"allowService.stun.value", data.TunnelInterfaceAllowStun.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowSnmpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.snmp.optionType", "variable")
			body, _ = sjson.Set(body, path+"allowService.snmp.value", data.TunnelInterfaceAllowSnmpVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowSnmp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.snmp.optionType", "default")
			body, _ = sjson.Set(body, path+"allowService.snmp.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"allowService.snmp.optionType", "global")
			body, _ = sjson.Set(body, path+"allowService.snmp.value", data.TunnelInterfaceAllowSnmp.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowNetconfVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.netconf.optionType", "variable")
			body, _ = sjson.Set(body, path+"allowService.netconf.value", data.TunnelInterfaceAllowNetconfVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowNetconf.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.netconf.optionType", "default")
			body, _ = sjson.Set(body, path+"allowService.netconf.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"allowService.netconf.optionType", "global")
			body, _ = sjson.Set(body, path+"allowService.netconf.value", data.TunnelInterfaceAllowNetconf.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowBfdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.bfd.optionType", "variable")
			body, _ = sjson.Set(body, path+"allowService.bfd.value", data.TunnelInterfaceAllowBfdVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowBfd.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.bfd.optionType", "default")
			body, _ = sjson.Set(body, path+"allowService.bfd.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"allowService.bfd.optionType", "global")
			body, _ = sjson.Set(body, path+"allowService.bfd.value", data.TunnelInterfaceAllowBfd.ValueBool())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"encapsulation", []interface{}{})
		for _, item := range data.TunnelInterfaceEncapsulations {
			itemBody := ""
			if !item.Encapsulation.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "encap.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "encap.value", item.Encapsulation.ValueString())
				}
			}

			if !item.PreferenceVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "preference.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "preference.value", item.PreferenceVariable.ValueString())
				}
			} else if item.Preference.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "preference.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "preference.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "preference.value", item.Preference.ValueInt64())
				}
			}

			if !item.WeightVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "weight.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "weight.value", item.WeightVariable.ValueString())
				}
			} else if item.Weight.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "weight.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "weight.value", 1)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "weight.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "weight.value", item.Weight.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"encapsulation.-1", itemBody)
		}
	}

	if !data.NatIpv4Variable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"nat.optionType", "variable")
			body, _ = sjson.Set(body, path+"nat.value", data.NatIpv4Variable.ValueString())
		}
	} else if data.NatIpv4.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"nat.optionType", "default")
			body, _ = sjson.Set(body, path+"nat.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"nat.optionType", "global")
			body, _ = sjson.Set(body, path+"nat.value", data.NatIpv4.ValueBool())
		}
	}

	if !data.NatTypeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natType.optionType", "variable")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natType.value", data.NatTypeVariable.ValueString())
		}
	} else if data.NatType.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natType.optionType", "default")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natType.value", "interface")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natType.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natType.value", data.NatType.ValueString())
		}
	}

	if !data.NatRangeStartVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeStart.optionType", "variable")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeStart.value", data.NatRangeStartVariable.ValueString())
		}
	} else if !data.NatRangeStart.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeStart.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeStart.value", data.NatRangeStart.ValueString())
		}
	}

	if !data.NatRangeEndVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeEnd.optionType", "variable")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeEnd.value", data.NatRangeEndVariable.ValueString())
		}
	} else if !data.NatRangeEnd.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeEnd.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeEnd.value", data.NatRangeEnd.ValueString())
		}
	}

	if !data.NatPrefixLengthVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.prefixLength.optionType", "variable")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.prefixLength.value", data.NatPrefixLengthVariable.ValueString())
		}
	} else if !data.NatPrefixLength.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.prefixLength.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.prefixLength.value", data.NatPrefixLength.ValueInt64())
		}
	}

	if !data.NatOverloadVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.overload.optionType", "variable")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.overload.value", data.NatOverloadVariable.ValueString())
		}
	} else if !data.NatOverload.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.overload.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.overload.value", data.NatOverload.ValueBool())
		}
	}

	if !data.NatLoopbackVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natLookback.optionType", "variable")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natLookback.value", data.NatLoopbackVariable.ValueString())
		}
	} else if !data.NatLoopback.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natLookback.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natLookback.value", data.NatLoopback.ValueString())
		}
	}

	if !data.NatUdpTimeoutVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.optionType", "variable")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.value", data.NatUdpTimeoutVariable.ValueString())
		}
	} else if data.NatUdpTimeout.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.optionType", "default")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.value", 1)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.value", data.NatUdpTimeout.ValueInt64())
		}
	}

	if !data.NatTcpTimeoutVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.optionType", "variable")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.value", data.NatTcpTimeoutVariable.ValueString())
		}
	} else if data.NatTcpTimeout.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.optionType", "default")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.value", 60)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.value", data.NatTcpTimeout.ValueInt64())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.newStaticNat", []interface{}{})
		for _, item := range data.NewStaticNats {
			itemBody := ""

			if !item.SourceIpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceIp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourceIp.value", item.SourceIpVariable.ValueString())
				}
			} else if !item.SourceIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceIp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sourceIp.value", item.SourceIp.ValueString())
				}
			}

			if !item.TranslatedIpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "translateIp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "translateIp.value", item.TranslatedIpVariable.ValueString())
				}
			} else if !item.TranslatedIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "translateIp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "translateIp.value", item.TranslatedIp.ValueString())
				}
			}
			if item.Direction.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "staticNatDirection.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "staticNatDirection.value", "inside")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "staticNatDirection.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "staticNatDirection.value", item.Direction.ValueString())
				}
			}

			if !item.SourceVpnVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceVpn.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourceVpn.value", item.SourceVpnVariable.ValueString())
				}
			} else if item.SourceVpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceVpn.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "sourceVpn.value", 0)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceVpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sourceVpn.value", item.SourceVpn.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"natAttributesIpv4.newStaticNat.-1", itemBody)
		}
	}

	if !data.NatIpv6Variable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natIpv6.optionType", "variable")
			body, _ = sjson.Set(body, path+"natIpv6.value", data.NatIpv6Variable.ValueString())
		}
	} else if data.NatIpv6.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natIpv6.optionType", "default")
			body, _ = sjson.Set(body, path+"natIpv6.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"natIpv6.optionType", "global")
			body, _ = sjson.Set(body, path+"natIpv6.value", data.NatIpv6.ValueBool())
		}
	}
	if !data.Nat64.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv6.nat64.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv6.nat64.value", data.Nat64.ValueBool())
		}
	}
	if !data.Nat66.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv6.nat66.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv6.nat66.value", data.Nat66.ValueBool())
		}
	}
	if true {

		for _, item := range data.StaticNat66 {
			itemBody := ""

			if !item.SourcePrefixVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourcePrefix.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourcePrefix.value", item.SourcePrefixVariable.ValueString())
				}
			} else if !item.SourcePrefix.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourcePrefix.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sourcePrefix.value", item.SourcePrefix.ValueString())
				}
			}

			if !item.TranslatedSourcePrefixVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "translatedSourcePrefix.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "translatedSourcePrefix.value", item.TranslatedSourcePrefixVariable.ValueString())
				}
			} else if !item.TranslatedSourcePrefix.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "translatedSourcePrefix.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "translatedSourcePrefix.value", item.TranslatedSourcePrefix.ValueString())
				}
			}

			if !item.SourceVpnIdVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceVpnId.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourceVpnId.value", item.SourceVpnIdVariable.ValueString())
				}
			} else if item.SourceVpnId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceVpnId.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceVpnId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sourceVpnId.value", item.SourceVpnId.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"natAttributesIpv6.staticNat66.-1", itemBody)
		}
	}
	if !data.QosAdaptive.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.adaptiveQoS.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.adaptiveQoS.value", data.QosAdaptive.ValueBool())
		}
	}

	if !data.QosAdaptivePeriodVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.adaptPeriod.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.adaptPeriod.value", data.QosAdaptivePeriodVariable.ValueString())
		}
	} else if !data.QosAdaptivePeriod.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.adaptPeriod.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.adaptPeriod.value", data.QosAdaptivePeriod.ValueInt64())
		}
	}
	if !data.QosAdaptiveBandwidthUpstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstream.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstream.value", data.QosAdaptiveBandwidthUpstream.ValueBool())
		}
	}

	if !data.QosAdaptiveMinUpstreamVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.value", data.QosAdaptiveMinUpstreamVariable.ValueString())
		}
	} else if !data.QosAdaptiveMinUpstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.value", data.QosAdaptiveMinUpstream.ValueInt64())
		}
	}

	if !data.QosAdaptiveMaxUpstreamVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.value", data.QosAdaptiveMaxUpstreamVariable.ValueString())
		}
	} else if !data.QosAdaptiveMaxUpstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.value", data.QosAdaptiveMaxUpstream.ValueInt64())
		}
	}

	if !data.QosAdaptiveDefaultUpstreamVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.value", data.QosAdaptiveDefaultUpstreamVariable.ValueString())
		}
	} else if !data.QosAdaptiveDefaultUpstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.value", data.QosAdaptiveDefaultUpstream.ValueInt64())
		}
	}
	if !data.QosAdaptiveBandwidthDownstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstream.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstream.value", data.QosAdaptiveBandwidthDownstream.ValueBool())
		}
	}

	if !data.QosAdaptiveMinDownstreamVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.value", data.QosAdaptiveMinDownstreamVariable.ValueString())
		}
	} else if !data.QosAdaptiveMinDownstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.value", data.QosAdaptiveMinDownstream.ValueInt64())
		}
	}

	if !data.QosAdaptiveMaxDownstreamVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.value", data.QosAdaptiveMaxDownstreamVariable.ValueString())
		}
	} else if !data.QosAdaptiveMaxDownstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.value", data.QosAdaptiveMaxDownstream.ValueInt64())
		}
	}

	if !data.QosAdaptiveDefaultDownstreamVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.value", data.QosAdaptiveDefaultDownstreamVariable.ValueString())
		}
	} else if !data.QosAdaptiveDefaultDownstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.value", data.QosAdaptiveDefaultDownstream.ValueInt64())
		}
	}

	if !data.QosShapingRateVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRate.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.shapingRate.value", data.QosShapingRateVariable.ValueString())
		}
	} else if !data.QosShapingRate.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRate.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRate.value", data.QosShapingRate.ValueInt64())
		}
	}
	if !data.AclIpv4EgressFeatureId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.ipv4AclEgress.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.ipv4AclEgress.refId.value", data.AclIpv4EgressFeatureId.ValueString())
		}
	}
	if !data.AclIpv4IngressFeatureId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.ipv4AclIngress.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.ipv4AclIngress.refId.value", data.AclIpv4IngressFeatureId.ValueString())
		}
	}
	if !data.AclIpv6EgressFeatureId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.ipv6AclEgress.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.ipv6AclEgress.refId.value", data.AclIpv6EgressFeatureId.ValueString())
		}
	}
	if !data.AclIpv6IngressFeatureId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.ipv6AclIngress.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.ipv6AclIngress.refId.value", data.AclIpv6IngressFeatureId.ValueString())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"arp", []interface{}{})
		for _, item := range data.Arps {
			itemBody := ""

			if !item.IpAddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipAddress.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ipAddress.value", item.IpAddressVariable.ValueString())
				}
			} else if item.IpAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipAddress.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipAddress.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ipAddress.value", item.IpAddress.ValueString())
				}
			}

			if !item.MacAddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "macAddress.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "macAddress.value", item.MacAddressVariable.ValueString())
				}
			} else if item.MacAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "macAddress.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "macAddress.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "macAddress.value", item.MacAddress.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"arp.-1", itemBody)
		}
	}

	if !data.IcmpRedirectDisableVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.icmpRedirectDisable.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.icmpRedirectDisable.value", data.IcmpRedirectDisableVariable.ValueString())
		}
	} else if data.IcmpRedirectDisable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.icmpRedirectDisable.optionType", "default")
			body, _ = sjson.Set(body, path+"advanced.icmpRedirectDisable.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.icmpRedirectDisable.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.icmpRedirectDisable.value", data.IcmpRedirectDisable.ValueBool())
		}
	}

	if !data.DuplexVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.duplex.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.duplex.value", data.DuplexVariable.ValueString())
		}
	} else if data.Duplex.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.duplex.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.duplex.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.duplex.value", data.Duplex.ValueString())
		}
	}

	if !data.MacAddressVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.macAddress.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.macAddress.value", data.MacAddressVariable.ValueString())
		}
	} else if data.MacAddress.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.macAddress.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.macAddress.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.macAddress.value", data.MacAddress.ValueString())
		}
	}

	if !data.IpMtuVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.ipMtu.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.ipMtu.value", data.IpMtuVariable.ValueString())
		}
	} else if data.IpMtu.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.ipMtu.optionType", "default")
			body, _ = sjson.Set(body, path+"advanced.ipMtu.value", 1500)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.ipMtu.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.ipMtu.value", data.IpMtu.ValueInt64())
		}
	}

	if !data.InterfaceMtuVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.intrfMtu.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.intrfMtu.value", data.InterfaceMtuVariable.ValueString())
		}
	} else if data.InterfaceMtu.IsNull() {
		if !strings.Contains(data.InterfaceName.ValueString(), ".") {
			body, _ = sjson.Set(body, path+"advanced.intrfMtu.optionType", "default")
			body, _ = sjson.Set(body, path+"advanced.intrfMtu.value", 1500)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.intrfMtu.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.intrfMtu.value", data.InterfaceMtu.ValueInt64())
		}
	}

	if !data.TcpMssVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tcpMss.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.tcpMss.value", data.TcpMssVariable.ValueString())
		}
	} else if data.TcpMss.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tcpMss.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tcpMss.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.tcpMss.value", data.TcpMss.ValueInt64())
		}
	}

	if !data.SpeedVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.speed.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.speed.value", data.SpeedVariable.ValueString())
		}
	} else if data.Speed.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.speed.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.speed.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.speed.value", data.Speed.ValueString())
		}
	}

	if !data.ArpTimeoutVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.arpTimeout.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.arpTimeout.value", data.ArpTimeoutVariable.ValueString())
		}
	} else if data.ArpTimeout.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.arpTimeout.optionType", "default")
			body, _ = sjson.Set(body, path+"advanced.arpTimeout.value", 1200)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.arpTimeout.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.arpTimeout.value", data.ArpTimeout.ValueInt64())
		}
	}

	if !data.AutonegotiateVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.autonegotiate.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.autonegotiate.value", data.AutonegotiateVariable.ValueString())
		}
	} else if data.Autonegotiate.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.autonegotiate.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.autonegotiate.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.autonegotiate.value", data.Autonegotiate.ValueBool())
		}
	}

	if !data.MediaTypeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.mediaType.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.mediaType.value", data.MediaTypeVariable.ValueString())
		}
	} else if data.MediaType.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.mediaType.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.mediaType.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.mediaType.value", data.MediaType.ValueString())
		}
	}

	if !data.TlocExtensionVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tlocExtension.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.tlocExtension.value", data.TlocExtensionVariable.ValueString())
		}
	} else if data.TlocExtension.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tlocExtension.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tlocExtension.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.tlocExtension.value", data.TlocExtension.ValueString())
		}
	}

	if !data.GreTunnelSourceIpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.sourceIp.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.sourceIp.value", data.GreTunnelSourceIpVariable.ValueString())
		}
	} else if data.GreTunnelSourceIp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.sourceIp.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.sourceIp.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.sourceIp.value", data.GreTunnelSourceIp.ValueString())
		}
	}

	if !data.XconnectVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.xconnect.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.xconnect.value", data.XconnectVariable.ValueString())
		}
	} else if data.Xconnect.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.xconnect.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.xconnect.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.tlocExtensionGreFrom.xconnect.value", data.Xconnect.ValueString())
		}
	}

	if !data.LoadIntervalVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.loadInterval.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.loadInterval.value", data.LoadIntervalVariable.ValueString())
		}
	} else if data.LoadInterval.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.loadInterval.optionType", "default")
			body, _ = sjson.Set(body, path+"advanced.loadInterval.value", 30)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.loadInterval.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.loadInterval.value", data.LoadInterval.ValueInt64())
		}
	}

	if !data.TrackerVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tracker.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.tracker.value", data.TrackerVariable.ValueString())
		}
	} else if data.Tracker.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tracker.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tracker.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.tracker.value", data.Tracker.ValueString())
		}
	}

	if !data.IpDirectedBroadcastVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.ipDirectedBroadcast.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.ipDirectedBroadcast.value", data.IpDirectedBroadcastVariable.ValueString())
		}
	} else if data.IpDirectedBroadcast.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.ipDirectedBroadcast.optionType", "default")
			body, _ = sjson.Set(body, path+"advanced.ipDirectedBroadcast.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.ipDirectedBroadcast.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.ipDirectedBroadcast.value", data.IpDirectedBroadcast.ValueBool())
		}
	}
	return body
}

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
	data.InterfaceDescription = types.StringNull()
	data.InterfaceDescriptionVariable = types.StringNull()
	if t := res.Get(path + "description.optionType"); t.Exists() {
		va := res.Get(path + "description.value")
		if t.String() == "variable" {
			data.InterfaceDescriptionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceDescription = types.StringValue(va.String())
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
		data.Ipv6SecondaryAddresses = make([]TransportWANVPNInterfaceEthernetIpv6SecondaryAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportWANVPNInterfaceEthernetIpv6SecondaryAddresses{}
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
			data.Ipv6SecondaryAddresses = append(data.Ipv6SecondaryAddresses, item)
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
	data.TunnelInterfaceAllowDns = types.BoolNull()
	data.TunnelInterfaceAllowDnsVariable = types.StringNull()
	if t := res.Get(path + "allowService.dns.optionType"); t.Exists() {
		va := res.Get(path + "allowService.dns.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowDnsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowDns = types.BoolValue(va.Bool())
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
	data.QosAdaptive = types.BoolNull()

	if t := res.Get(path + "aclQos.adaptiveQoS.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.adaptiveQoS.value")
		if t.String() == "global" {
			data.QosAdaptive = types.BoolValue(va.Bool())
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
	data.AclIpv4EgressFeatureId = types.StringNull()

	if t := res.Get(path + "aclQos.ipv4AclEgress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv4AclEgress.refId.value")
		if t.String() == "global" {
			data.AclIpv4EgressFeatureId = types.StringValue(va.String())
		}
	}
	data.AclIpv4IngressFeatureId = types.StringNull()

	if t := res.Get(path + "aclQos.ipv4AclIngress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv4AclIngress.refId.value")
		if t.String() == "global" {
			data.AclIpv4IngressFeatureId = types.StringValue(va.String())
		}
	}
	data.AclIpv6EgressFeatureId = types.StringNull()

	if t := res.Get(path + "aclQos.ipv6AclEgress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv6AclEgress.refId.value")
		if t.String() == "global" {
			data.AclIpv6EgressFeatureId = types.StringValue(va.String())
		}
	}
	data.AclIpv6IngressFeatureId = types.StringNull()

	if t := res.Get(path + "aclQos.ipv6AclIngress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv6AclIngress.refId.value")
		if t.String() == "global" {
			data.AclIpv6IngressFeatureId = types.StringValue(va.String())
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
	data.IcmpRedirectDisable = types.BoolNull()
	data.IcmpRedirectDisableVariable = types.StringNull()
	if t := res.Get(path + "advanced.icmpRedirectDisable.optionType"); t.Exists() {
		va := res.Get(path + "advanced.icmpRedirectDisable.value")
		if t.String() == "variable" {
			data.IcmpRedirectDisableVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IcmpRedirectDisable = types.BoolValue(va.Bool())
		}
	}
	data.Duplex = types.StringNull()
	data.DuplexVariable = types.StringNull()
	if t := res.Get(path + "advanced.duplex.optionType"); t.Exists() {
		va := res.Get(path + "advanced.duplex.value")
		if t.String() == "variable" {
			data.DuplexVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Duplex = types.StringValue(va.String())
		}
	}
	data.MacAddress = types.StringNull()
	data.MacAddressVariable = types.StringNull()
	if t := res.Get(path + "advanced.macAddress.optionType"); t.Exists() {
		va := res.Get(path + "advanced.macAddress.value")
		if t.String() == "variable" {
			data.MacAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MacAddress = types.StringValue(va.String())
		}
	}
	data.IpMtu = types.Int64Null()
	data.IpMtuVariable = types.StringNull()
	if t := res.Get(path + "advanced.ipMtu.optionType"); t.Exists() {
		va := res.Get(path + "advanced.ipMtu.value")
		if t.String() == "variable" {
			data.IpMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpMtu = types.Int64Value(va.Int())
		}
	}
	data.InterfaceMtu = types.Int64Null()
	data.InterfaceMtuVariable = types.StringNull()
	if t := res.Get(path + "advanced.intrfMtu.optionType"); t.Exists() {
		va := res.Get(path + "advanced.intrfMtu.value")
		if t.String() == "variable" {
			data.InterfaceMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceMtu = types.Int64Value(va.Int())
		}
	}
	data.TcpMss = types.Int64Null()
	data.TcpMssVariable = types.StringNull()
	if t := res.Get(path + "advanced.tcpMss.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tcpMss.value")
		if t.String() == "variable" {
			data.TcpMssVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TcpMss = types.Int64Value(va.Int())
		}
	}
	data.Speed = types.StringNull()
	data.SpeedVariable = types.StringNull()
	if t := res.Get(path + "advanced.speed.optionType"); t.Exists() {
		va := res.Get(path + "advanced.speed.value")
		if t.String() == "variable" {
			data.SpeedVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Speed = types.StringValue(va.String())
		}
	}
	data.ArpTimeout = types.Int64Null()
	data.ArpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "advanced.arpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "advanced.arpTimeout.value")
		if t.String() == "variable" {
			data.ArpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ArpTimeout = types.Int64Value(va.Int())
		}
	}
	data.Autonegotiate = types.BoolNull()
	data.AutonegotiateVariable = types.StringNull()
	if t := res.Get(path + "advanced.autonegotiate.optionType"); t.Exists() {
		va := res.Get(path + "advanced.autonegotiate.value")
		if t.String() == "variable" {
			data.AutonegotiateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Autonegotiate = types.BoolValue(va.Bool())
		}
	}
	data.MediaType = types.StringNull()
	data.MediaTypeVariable = types.StringNull()
	if t := res.Get(path + "advanced.mediaType.optionType"); t.Exists() {
		va := res.Get(path + "advanced.mediaType.value")
		if t.String() == "variable" {
			data.MediaTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MediaType = types.StringValue(va.String())
		}
	}
	data.TlocExtension = types.StringNull()
	data.TlocExtensionVariable = types.StringNull()
	if t := res.Get(path + "advanced.tlocExtension.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tlocExtension.value")
		if t.String() == "variable" {
			data.TlocExtensionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TlocExtension = types.StringValue(va.String())
		}
	}
	data.GreTunnelSourceIp = types.StringNull()
	data.GreTunnelSourceIpVariable = types.StringNull()
	if t := res.Get(path + "advanced.tlocExtensionGreFrom.sourceIp.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tlocExtensionGreFrom.sourceIp.value")
		if t.String() == "variable" {
			data.GreTunnelSourceIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GreTunnelSourceIp = types.StringValue(va.String())
		}
	}
	data.Xconnect = types.StringNull()
	data.XconnectVariable = types.StringNull()
	if t := res.Get(path + "advanced.tlocExtensionGreFrom.xconnect.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tlocExtensionGreFrom.xconnect.value")
		if t.String() == "variable" {
			data.XconnectVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Xconnect = types.StringValue(va.String())
		}
	}
	data.LoadInterval = types.Int64Null()
	data.LoadIntervalVariable = types.StringNull()
	if t := res.Get(path + "advanced.loadInterval.optionType"); t.Exists() {
		va := res.Get(path + "advanced.loadInterval.value")
		if t.String() == "variable" {
			data.LoadIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.LoadInterval = types.Int64Value(va.Int())
		}
	}
	data.Tracker = types.StringNull()
	data.TrackerVariable = types.StringNull()
	if t := res.Get(path + "advanced.tracker.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tracker.value")
		if t.String() == "variable" {
			data.TrackerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Tracker = types.StringValue(va.String())
		}
	}
	data.IpDirectedBroadcast = types.BoolNull()
	data.IpDirectedBroadcastVariable = types.StringNull()
	if t := res.Get(path + "advanced.ipDirectedBroadcast.optionType"); t.Exists() {
		va := res.Get(path + "advanced.ipDirectedBroadcast.value")
		if t.String() == "variable" {
			data.IpDirectedBroadcastVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpDirectedBroadcast = types.BoolValue(va.Bool())
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
	data.InterfaceDescription = types.StringNull()
	data.InterfaceDescriptionVariable = types.StringNull()
	if t := res.Get(path + "description.optionType"); t.Exists() {
		va := res.Get(path + "description.value")
		if t.String() == "variable" {
			data.InterfaceDescriptionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceDescription = types.StringValue(va.String())
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
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
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
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
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
	for i := range data.Ipv6SecondaryAddresses {
		keys := [...]string{"address"}
		keyValues := [...]string{data.Ipv6SecondaryAddresses[i].Address.ValueString()}
		keyValuesVariables := [...]string{data.Ipv6SecondaryAddresses[i].AddressVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "intfIpV6Address.static.secondaryIpV6Address").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Ipv6SecondaryAddresses[i].Address = types.StringNull()
		data.Ipv6SecondaryAddresses[i].AddressVariable = types.StringNull()
		if t := r.Get("address.optionType"); t.Exists() {
			va := r.Get("address.value")
			if t.String() == "variable" {
				data.Ipv6SecondaryAddresses[i].AddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6SecondaryAddresses[i].Address = types.StringValue(va.String())
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
	data.TunnelInterfaceAllowDns = types.BoolNull()
	data.TunnelInterfaceAllowDnsVariable = types.StringNull()
	if t := res.Get(path + "allowService.dns.optionType"); t.Exists() {
		va := res.Get(path + "allowService.dns.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowDnsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowDns = types.BoolValue(va.Bool())
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
		keys := [...]string{"encap"}
		keyValues := [...]string{data.TunnelInterfaceEncapsulations[i].Encapsulation.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "encapsulation").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
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
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
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
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
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
	data.QosAdaptive = types.BoolNull()

	if t := res.Get(path + "aclQos.adaptiveQoS.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.adaptiveQoS.value")
		if t.String() == "global" {
			data.QosAdaptive = types.BoolValue(va.Bool())
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
	data.AclIpv4EgressFeatureId = types.StringNull()

	if t := res.Get(path + "aclQos.ipv4AclEgress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv4AclEgress.refId.value")
		if t.String() == "global" {
			data.AclIpv4EgressFeatureId = types.StringValue(va.String())
		}
	}
	data.AclIpv4IngressFeatureId = types.StringNull()

	if t := res.Get(path + "aclQos.ipv4AclIngress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv4AclIngress.refId.value")
		if t.String() == "global" {
			data.AclIpv4IngressFeatureId = types.StringValue(va.String())
		}
	}
	data.AclIpv6EgressFeatureId = types.StringNull()

	if t := res.Get(path + "aclQos.ipv6AclEgress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv6AclEgress.refId.value")
		if t.String() == "global" {
			data.AclIpv6EgressFeatureId = types.StringValue(va.String())
		}
	}
	data.AclIpv6IngressFeatureId = types.StringNull()

	if t := res.Get(path + "aclQos.ipv6AclIngress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv6AclIngress.refId.value")
		if t.String() == "global" {
			data.AclIpv6IngressFeatureId = types.StringValue(va.String())
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
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
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
	data.IcmpRedirectDisable = types.BoolNull()
	data.IcmpRedirectDisableVariable = types.StringNull()
	if t := res.Get(path + "advanced.icmpRedirectDisable.optionType"); t.Exists() {
		va := res.Get(path + "advanced.icmpRedirectDisable.value")
		if t.String() == "variable" {
			data.IcmpRedirectDisableVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IcmpRedirectDisable = types.BoolValue(va.Bool())
		}
	}
	data.Duplex = types.StringNull()
	data.DuplexVariable = types.StringNull()
	if t := res.Get(path + "advanced.duplex.optionType"); t.Exists() {
		va := res.Get(path + "advanced.duplex.value")
		if t.String() == "variable" {
			data.DuplexVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Duplex = types.StringValue(va.String())
		}
	}
	data.MacAddress = types.StringNull()
	data.MacAddressVariable = types.StringNull()
	if t := res.Get(path + "advanced.macAddress.optionType"); t.Exists() {
		va := res.Get(path + "advanced.macAddress.value")
		if t.String() == "variable" {
			data.MacAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MacAddress = types.StringValue(va.String())
		}
	}
	data.IpMtu = types.Int64Null()
	data.IpMtuVariable = types.StringNull()
	if t := res.Get(path + "advanced.ipMtu.optionType"); t.Exists() {
		va := res.Get(path + "advanced.ipMtu.value")
		if t.String() == "variable" {
			data.IpMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpMtu = types.Int64Value(va.Int())
		}
	}
	data.InterfaceMtu = types.Int64Null()
	data.InterfaceMtuVariable = types.StringNull()
	if t := res.Get(path + "advanced.intrfMtu.optionType"); t.Exists() {
		va := res.Get(path + "advanced.intrfMtu.value")
		if t.String() == "variable" {
			data.InterfaceMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceMtu = types.Int64Value(va.Int())
		}
	}
	data.TcpMss = types.Int64Null()
	data.TcpMssVariable = types.StringNull()
	if t := res.Get(path + "advanced.tcpMss.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tcpMss.value")
		if t.String() == "variable" {
			data.TcpMssVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TcpMss = types.Int64Value(va.Int())
		}
	}
	data.Speed = types.StringNull()
	data.SpeedVariable = types.StringNull()
	if t := res.Get(path + "advanced.speed.optionType"); t.Exists() {
		va := res.Get(path + "advanced.speed.value")
		if t.String() == "variable" {
			data.SpeedVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Speed = types.StringValue(va.String())
		}
	}
	data.ArpTimeout = types.Int64Null()
	data.ArpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "advanced.arpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "advanced.arpTimeout.value")
		if t.String() == "variable" {
			data.ArpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ArpTimeout = types.Int64Value(va.Int())
		}
	}
	data.Autonegotiate = types.BoolNull()
	data.AutonegotiateVariable = types.StringNull()
	if t := res.Get(path + "advanced.autonegotiate.optionType"); t.Exists() {
		va := res.Get(path + "advanced.autonegotiate.value")
		if t.String() == "variable" {
			data.AutonegotiateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Autonegotiate = types.BoolValue(va.Bool())
		}
	}
	data.MediaType = types.StringNull()
	data.MediaTypeVariable = types.StringNull()
	if t := res.Get(path + "advanced.mediaType.optionType"); t.Exists() {
		va := res.Get(path + "advanced.mediaType.value")
		if t.String() == "variable" {
			data.MediaTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MediaType = types.StringValue(va.String())
		}
	}
	data.TlocExtension = types.StringNull()
	data.TlocExtensionVariable = types.StringNull()
	if t := res.Get(path + "advanced.tlocExtension.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tlocExtension.value")
		if t.String() == "variable" {
			data.TlocExtensionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TlocExtension = types.StringValue(va.String())
		}
	}
	data.GreTunnelSourceIp = types.StringNull()
	data.GreTunnelSourceIpVariable = types.StringNull()
	if t := res.Get(path + "advanced.tlocExtensionGreFrom.sourceIp.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tlocExtensionGreFrom.sourceIp.value")
		if t.String() == "variable" {
			data.GreTunnelSourceIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GreTunnelSourceIp = types.StringValue(va.String())
		}
	}
	data.Xconnect = types.StringNull()
	data.XconnectVariable = types.StringNull()
	if t := res.Get(path + "advanced.tlocExtensionGreFrom.xconnect.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tlocExtensionGreFrom.xconnect.value")
		if t.String() == "variable" {
			data.XconnectVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Xconnect = types.StringValue(va.String())
		}
	}
	data.LoadInterval = types.Int64Null()
	data.LoadIntervalVariable = types.StringNull()
	if t := res.Get(path + "advanced.loadInterval.optionType"); t.Exists() {
		va := res.Get(path + "advanced.loadInterval.value")
		if t.String() == "variable" {
			data.LoadIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.LoadInterval = types.Int64Value(va.Int())
		}
	}
	data.Tracker = types.StringNull()
	data.TrackerVariable = types.StringNull()
	if t := res.Get(path + "advanced.tracker.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tracker.value")
		if t.String() == "variable" {
			data.TrackerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Tracker = types.StringValue(va.String())
		}
	}
	data.IpDirectedBroadcast = types.BoolNull()
	data.IpDirectedBroadcastVariable = types.StringNull()
	if t := res.Get(path + "advanced.ipDirectedBroadcast.optionType"); t.Exists() {
		va := res.Get(path + "advanced.ipDirectedBroadcast.value")
		if t.String() == "variable" {
			data.IpDirectedBroadcastVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpDirectedBroadcast = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end updateFromBody
