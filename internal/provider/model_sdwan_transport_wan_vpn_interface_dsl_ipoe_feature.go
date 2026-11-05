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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type TransportWANVPNInterfaceDSLIPoE struct {
	Id                                                 types.String `tfsdk:"id"`
	Version                                            types.Int64  `tfsdk:"version"`
	Name                                               types.String `tfsdk:"name"`
	Description                                        types.String `tfsdk:"description"`
	FeatureProfileId                                   types.String `tfsdk:"feature_profile_id"`
	TransportWanVpnFeatureId                           types.String `tfsdk:"transport_wan_vpn_feature_id"`
	Shutdown                                           types.Bool   `tfsdk:"shutdown"`
	ShutdownVariable                                   types.String `tfsdk:"shutdown_variable"`
	ServiceProvider                                    types.String `tfsdk:"service_provider"`
	ServiceProviderVariable                            types.String `tfsdk:"service_provider_variable"`
	BandwidthUpstream                                  types.Int64  `tfsdk:"bandwidth_upstream"`
	BandwidthUpstreamVariable                          types.String `tfsdk:"bandwidth_upstream_variable"`
	BandwidthDownstream                                types.Int64  `tfsdk:"bandwidth_downstream"`
	BandwidthDownstreamVariable                        types.String `tfsdk:"bandwidth_downstream_variable"`
	VdslSlot                                           types.String `tfsdk:"vdsl_slot"`
	VdslSlotVariable                                   types.String `tfsdk:"vdsl_slot_variable"`
	VdslMode                                           types.String `tfsdk:"vdsl_mode"`
	VdslSra                                            types.Bool   `tfsdk:"vdsl_sra"`
	VdslSraVariable                                    types.String `tfsdk:"vdsl_sra_variable"`
	NatEnabled                                         types.Bool   `tfsdk:"nat_enabled"`
	NatEnabledVariable                                 types.String `tfsdk:"nat_enabled_variable"`
	NatTcpTimeout                                      types.Int64  `tfsdk:"nat_tcp_timeout"`
	NatTcpTimeoutVariable                              types.String `tfsdk:"nat_tcp_timeout_variable"`
	NatUdpTimeout                                      types.Int64  `tfsdk:"nat_udp_timeout"`
	NatUdpTimeoutVariable                              types.String `tfsdk:"nat_udp_timeout_variable"`
	InterfaceName                                      types.String `tfsdk:"interface_name"`
	InterfaceNameVariable                              types.String `tfsdk:"interface_name_variable"`
	InterfaceDescription                               types.String `tfsdk:"interface_description"`
	InterfaceDescriptionVariable                       types.String `tfsdk:"interface_description_variable"`
	VlanId                                             types.Int64  `tfsdk:"vlan_id"`
	VlanIdVariable                                     types.String `tfsdk:"vlan_id_variable"`
	Ipv4AddressType                                    types.String `tfsdk:"ipv4_address_type"`
	Ipv4DhcpDistance                                   types.Int64  `tfsdk:"ipv4_dhcp_distance"`
	Ipv4DhcpDistanceVariable                           types.String `tfsdk:"ipv4_dhcp_distance_variable"`
	Ipv4DhcpHelper                                     types.String `tfsdk:"ipv4_dhcp_helper"`
	Ipv4DhcpHelperVariable                             types.String `tfsdk:"ipv4_dhcp_helper_variable"`
	Ipv4Address                                        types.String `tfsdk:"ipv4_address"`
	Ipv4AddressVariable                                types.String `tfsdk:"ipv4_address_variable"`
	Ipv4SubnetMask                                     types.String `tfsdk:"ipv4_subnet_mask"`
	Ipv4SubnetMaskVariable                             types.String `tfsdk:"ipv4_subnet_mask_variable"`
	TunnelInterfaceEnabled                             types.Bool   `tfsdk:"tunnel_interface_enabled"`
	TunnelPerTunnelQos                                 types.Bool   `tfsdk:"tunnel_per_tunnel_qos"`
	TunnelPerTunnelQosVariable                         types.String `tfsdk:"tunnel_per_tunnel_qos_variable"`
	TunnelQosMode                                      types.String `tfsdk:"tunnel_qos_mode"`
	TunnelQosModeVariable                              types.String `tfsdk:"tunnel_qos_mode_variable"`
	TunnelBandwidthPercent                             types.Int64  `tfsdk:"tunnel_bandwidth_percent"`
	TunnelBandwidthPercentVariable                     types.String `tfsdk:"tunnel_bandwidth_percent_variable"`
	TunnelInterfaceBindLoopbackTunnel                  types.String `tfsdk:"tunnel_interface_bind_loopback_tunnel"`
	TunnelInterfaceBindLoopbackTunnelVariable          types.String `tfsdk:"tunnel_interface_bind_loopback_tunnel_variable"`
	TunnelInterfaceCarrier                             types.String `tfsdk:"tunnel_interface_carrier"`
	TunnelInterfaceCarrierVariable                     types.String `tfsdk:"tunnel_interface_carrier_variable"`
	TunnelInterfaceColor                               types.String `tfsdk:"tunnel_interface_color"`
	TunnelInterfaceColorVariable                       types.String `tfsdk:"tunnel_interface_color_variable"`
	TunnelInterfaceColorDescription                    types.String `tfsdk:"tunnel_interface_color_description"`
	TunnelInterfaceColorDescriptionVariable            types.String `tfsdk:"tunnel_interface_color_description_variable"`
	TunnelInterfaceHelloInterval                       types.Int64  `tfsdk:"tunnel_interface_hello_interval"`
	TunnelInterfaceHelloIntervalVariable               types.String `tfsdk:"tunnel_interface_hello_interval_variable"`
	TunnelInterfaceHelloTolerance                      types.Int64  `tfsdk:"tunnel_interface_hello_tolerance"`
	TunnelInterfaceHelloToleranceVariable              types.String `tfsdk:"tunnel_interface_hello_tolerance_variable"`
	TunnelInterfaceLastResortCircuit                   types.Bool   `tfsdk:"tunnel_interface_last_resort_circuit"`
	TunnelInterfaceLastResortCircuitVariable           types.String `tfsdk:"tunnel_interface_last_resort_circuit_variable"`
	TunnelInterfaceColorRestrict                       types.Bool   `tfsdk:"tunnel_interface_color_restrict"`
	TunnelInterfaceColorRestrictVariable               types.String `tfsdk:"tunnel_interface_color_restrict_variable"`
	TunnelInterfaceGroups                              types.String `tfsdk:"tunnel_interface_groups"`
	TunnelInterfaceGroupsVariable                      types.String `tfsdk:"tunnel_interface_groups_variable"`
	TunnelInterfaceBorder                              types.Bool   `tfsdk:"tunnel_interface_border"`
	TunnelInterfaceBorderVariable                      types.String `tfsdk:"tunnel_interface_border_variable"`
	TunnelInterfaceMaxControlConnections               types.Int64  `tfsdk:"tunnel_interface_max_control_connections"`
	TunnelInterfaceMaxControlConnectionsVariable       types.String `tfsdk:"tunnel_interface_max_control_connections_variable"`
	TunnelInterfaceVbondAsStunServer                   types.Bool   `tfsdk:"tunnel_interface_vbond_as_stun_server"`
	TunnelInterfaceVbondAsStunServerVariable           types.String `tfsdk:"tunnel_interface_vbond_as_stun_server_variable"`
	TunnelInterfaceExcludeControllerGroupList          types.String `tfsdk:"tunnel_interface_exclude_controller_group_list"`
	TunnelInterfaceExcludeControllerGroupListVariable  types.String `tfsdk:"tunnel_interface_exclude_controller_group_list_variable"`
	TunnelInterfaceVmanageConnectionPreference         types.Int64  `tfsdk:"tunnel_interface_vmanage_connection_preference"`
	TunnelInterfaceVmanageConnectionPreferenceVariable types.String `tfsdk:"tunnel_interface_vmanage_connection_preference_variable"`
	TunnelInterfaceFullPortHop                         types.Bool   `tfsdk:"tunnel_interface_full_port_hop"`
	TunnelInterfaceFullPortHopVariable                 types.String `tfsdk:"tunnel_interface_full_port_hop_variable"`
	TunnelInterfacePortHop                             types.Bool   `tfsdk:"tunnel_interface_port_hop"`
	TunnelInterfacePortHopVariable                     types.String `tfsdk:"tunnel_interface_port_hop_variable"`
	TunnelInterfaceLowBandwidthLink                    types.Bool   `tfsdk:"tunnel_interface_low_bandwidth_link"`
	TunnelInterfaceLowBandwidthLinkVariable            types.String `tfsdk:"tunnel_interface_low_bandwidth_link_variable"`
	TunnelInterfaceNetworkBroadcast                    types.Bool   `tfsdk:"tunnel_interface_network_broadcast"`
	TunnelInterfaceNetworkBroadcastVariable            types.String `tfsdk:"tunnel_interface_network_broadcast_variable"`
	TunnelInterfaceTunnelTcpMss                        types.Int64  `tfsdk:"tunnel_interface_tunnel_tcp_mss"`
	TunnelInterfaceTunnelTcpMssVariable                types.String `tfsdk:"tunnel_interface_tunnel_tcp_mss_variable"`
	TunnelInterfaceClearDontFragment                   types.Bool   `tfsdk:"tunnel_interface_clear_dont_fragment"`
	TunnelInterfaceClearDontFragmentVariable           types.String `tfsdk:"tunnel_interface_clear_dont_fragment_variable"`
	TunnelInterfaceAllowAll                            types.Bool   `tfsdk:"tunnel_interface_allow_all"`
	TunnelInterfaceAllowAllVariable                    types.String `tfsdk:"tunnel_interface_allow_all_variable"`
	TunnelInterfaceAllowBgp                            types.Bool   `tfsdk:"tunnel_interface_allow_bgp"`
	TunnelInterfaceAllowBgpVariable                    types.String `tfsdk:"tunnel_interface_allow_bgp_variable"`
	TunnelInterfaceAllowDhcp                           types.Bool   `tfsdk:"tunnel_interface_allow_dhcp"`
	TunnelInterfaceAllowDhcpVariable                   types.String `tfsdk:"tunnel_interface_allow_dhcp_variable"`
	TunnelInterfaceAllowDns                            types.Bool   `tfsdk:"tunnel_interface_allow_dns"`
	TunnelInterfaceAllowDnsVariable                    types.String `tfsdk:"tunnel_interface_allow_dns_variable"`
	TunnelInterfaceAllowIcmp                           types.Bool   `tfsdk:"tunnel_interface_allow_icmp"`
	TunnelInterfaceAllowIcmpVariable                   types.String `tfsdk:"tunnel_interface_allow_icmp_variable"`
	TunnelInterfaceAllowSsh                            types.Bool   `tfsdk:"tunnel_interface_allow_ssh"`
	TunnelInterfaceAllowSshVariable                    types.String `tfsdk:"tunnel_interface_allow_ssh_variable"`
	TunnelInterfaceAllowNtp                            types.Bool   `tfsdk:"tunnel_interface_allow_ntp"`
	TunnelInterfaceAllowNtpVariable                    types.String `tfsdk:"tunnel_interface_allow_ntp_variable"`
	TunnelInterfaceAllowStun                           types.Bool   `tfsdk:"tunnel_interface_allow_stun"`
	TunnelInterfaceAllowStunVariable                   types.String `tfsdk:"tunnel_interface_allow_stun_variable"`
	TunnelInterfaceAllowNetconf                        types.Bool   `tfsdk:"tunnel_interface_allow_netconf"`
	TunnelInterfaceAllowNetconfVariable                types.String `tfsdk:"tunnel_interface_allow_netconf_variable"`
	TunnelInterfaceAllowOspf                           types.Bool   `tfsdk:"tunnel_interface_allow_ospf"`
	TunnelInterfaceAllowOspfVariable                   types.String `tfsdk:"tunnel_interface_allow_ospf_variable"`
	TunnelInterfaceAllowSnmp                           types.Bool   `tfsdk:"tunnel_interface_allow_snmp"`
	TunnelInterfaceAllowSnmpVariable                   types.String `tfsdk:"tunnel_interface_allow_snmp_variable"`
	TunnelInterfaceIpsecEncapsulation                  types.Bool   `tfsdk:"tunnel_interface_ipsec_encapsulation"`
	TunnelInterfaceIpsecPreference                     types.Int64  `tfsdk:"tunnel_interface_ipsec_preference"`
	TunnelInterfaceIpsecPreferenceVariable             types.String `tfsdk:"tunnel_interface_ipsec_preference_variable"`
	TunnelInterfaceIpsecWeight                         types.Int64  `tfsdk:"tunnel_interface_ipsec_weight"`
	TunnelInterfaceIpsecWeightVariable                 types.String `tfsdk:"tunnel_interface_ipsec_weight_variable"`
	TunnelInterfaceGreEncapsulation                    types.Bool   `tfsdk:"tunnel_interface_gre_encapsulation"`
	TunnelInterfaceGrePreference                       types.Int64  `tfsdk:"tunnel_interface_gre_preference"`
	TunnelInterfaceGrePreferenceVariable               types.String `tfsdk:"tunnel_interface_gre_preference_variable"`
	TunnelInterfaceGreWeight                           types.Int64  `tfsdk:"tunnel_interface_gre_weight"`
	TunnelInterfaceGreWeightVariable                   types.String `tfsdk:"tunnel_interface_gre_weight_variable"`
	TunnelInterfaceNatRefreshInterval                  types.Int64  `tfsdk:"tunnel_interface_nat_refresh_interval"`
	TunnelInterfaceNatRefreshIntervalVariable          types.String `tfsdk:"tunnel_interface_nat_refresh_interval_variable"`
	MultiRegionFabricEnableCoreRegion                  types.Bool   `tfsdk:"multi_region_fabric_enable_core_region"`
	MultiRegionFabricCoreRegion                        types.String `tfsdk:"multi_region_fabric_core_region"`
	MultiRegionFabricEnableSecondaryRegion             types.Bool   `tfsdk:"multi_region_fabric_enable_secondary_region"`
	MultiRegionFabricSecondaryRegion                   types.String `tfsdk:"multi_region_fabric_secondary_region"`
	IpDirectedBroadcast                                types.Bool   `tfsdk:"ip_directed_broadcast"`
	IpDirectedBroadcastVariable                        types.String `tfsdk:"ip_directed_broadcast_variable"`
	IpMtu                                              types.Int64  `tfsdk:"ip_mtu"`
	IpMtuVariable                                      types.String `tfsdk:"ip_mtu_variable"`
	TcpMss                                             types.Int64  `tfsdk:"tcp_mss"`
	TcpMssVariable                                     types.String `tfsdk:"tcp_mss_variable"`
	TlocExtension                                      types.String `tfsdk:"tloc_extension"`
	TlocExtensionVariable                              types.String `tfsdk:"tloc_extension_variable"`
	AdaptiveQos                                        types.Bool   `tfsdk:"adaptive_qos"`
	AdaptiveQosPeriod                                  types.Int64  `tfsdk:"adaptive_qos_period"`
	AdaptiveQosPeriodVariable                          types.String `tfsdk:"adaptive_qos_period_variable"`
	ShapingRateUpstream                                types.Bool   `tfsdk:"shaping_rate_upstream"`
	MinShapingRateUpstream                             types.Int64  `tfsdk:"min_shaping_rate_upstream"`
	MinShapingRateUpstreamVariable                     types.String `tfsdk:"min_shaping_rate_upstream_variable"`
	MaxShapingRateUpstream                             types.Int64  `tfsdk:"max_shaping_rate_upstream"`
	MaxShapingRateUpstreamVariable                     types.String `tfsdk:"max_shaping_rate_upstream_variable"`
	DefaultShapingRateUpstream                         types.Int64  `tfsdk:"default_shaping_rate_upstream"`
	DefaultShapingRateUpstreamVariable                 types.String `tfsdk:"default_shaping_rate_upstream_variable"`
	ShapingRateDownstream                              types.Bool   `tfsdk:"shaping_rate_downstream"`
	MinShapingRateDownstream                           types.Int64  `tfsdk:"min_shaping_rate_downstream"`
	MinShapingRateDownstreamVariable                   types.String `tfsdk:"min_shaping_rate_downstream_variable"`
	MaxShapingRateDownstream                           types.Int64  `tfsdk:"max_shaping_rate_downstream"`
	MaxShapingRateDownstreamVariable                   types.String `tfsdk:"max_shaping_rate_downstream_variable"`
	DefaultShapingRateDownstream                       types.Int64  `tfsdk:"default_shaping_rate_downstream"`
	DefaultShapingRateDownstreamVariable               types.String `tfsdk:"default_shaping_rate_downstream_variable"`
	ShapingRate                                        types.Int64  `tfsdk:"shaping_rate"`
	ShapingRateVariable                                types.String `tfsdk:"shaping_rate_variable"`
	AclIpv4EgressFeatureId                             types.String `tfsdk:"acl_ipv4_egress_feature_id"`
	AclIpv4EgressFeatureIdVariable                     types.String `tfsdk:"acl_ipv4_egress_feature_id_variable"`
	AclIpv4IngressFeatureId                            types.String `tfsdk:"acl_ipv4_ingress_feature_id"`
	AclIpv4IngressFeatureIdVariable                    types.String `tfsdk:"acl_ipv4_ingress_feature_id_variable"`
	AclIpv6EgressFeatureId                             types.String `tfsdk:"acl_ipv6_egress_feature_id"`
	AclIpv6EgressFeatureIdVariable                     types.String `tfsdk:"acl_ipv6_egress_feature_id_variable"`
	AclIpv6IngressFeatureId                            types.String `tfsdk:"acl_ipv6_ingress_feature_id"`
	AclIpv6IngressFeatureIdVariable                    types.String `tfsdk:"acl_ipv6_ingress_feature_id_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TransportWANVPNInterfaceDSLIPoE) getModel() string {
	return "transport_wan_vpn_interface_dsl_ipoe"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportWANVPNInterfaceDSLIPoE) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/wan/vpn/%s/interface/dsl-ipoe", url.QueryEscape(data.FeatureProfileId.ValueString()), url.QueryEscape(data.TransportWanVpnFeatureId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportWANVPNInterfaceDSLIPoE) toBody(ctx context.Context) string {
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
			body, _ = sjson.Set(body, path+"shutdown.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"shutdown.optionType", "global")
			body, _ = sjson.Set(body, path+"shutdown.value", data.Shutdown.ValueBool())
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

	if !data.VdslSlotVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"vdsl.slot.optionType", "variable")
			body, _ = sjson.Set(body, path+"vdsl.slot.value", data.VdslSlotVariable.ValueString())
		}
	} else if !data.VdslSlot.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"vdsl.slot.optionType", "global")
			body, _ = sjson.Set(body, path+"vdsl.slot.value", data.VdslSlot.ValueString())
		}
	}
	if data.VdslMode.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"vdsl.mode.optionType", "default")
			body, _ = sjson.Set(body, path+"vdsl.mode.value", "auto")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"vdsl.mode.optionType", "global")
			body, _ = sjson.Set(body, path+"vdsl.mode.value", data.VdslMode.ValueString())
		}
	}

	if !data.VdslSraVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"vdsl.sra.optionType", "variable")
			body, _ = sjson.Set(body, path+"vdsl.sra.value", data.VdslSraVariable.ValueString())
		}
	} else if data.VdslSra.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"vdsl.sra.optionType", "default")
			body, _ = sjson.Set(body, path+"vdsl.sra.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"vdsl.sra.optionType", "global")
			body, _ = sjson.Set(body, path+"vdsl.sra.value", data.VdslSra.ValueBool())
		}
	}

	if !data.NatEnabledVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natProp.nat.optionType", "variable")
			body, _ = sjson.Set(body, path+"natProp.nat.value", data.NatEnabledVariable.ValueString())
		}
	} else if data.NatEnabled.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natProp.nat.optionType", "default")
			body, _ = sjson.Set(body, path+"natProp.nat.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"natProp.nat.optionType", "global")
			body, _ = sjson.Set(body, path+"natProp.nat.value", data.NatEnabled.ValueBool())
		}
	}

	if !data.NatTcpTimeoutVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natProp.tcpTimeout.optionType", "variable")
			body, _ = sjson.Set(body, path+"natProp.tcpTimeout.value", data.NatTcpTimeoutVariable.ValueString())
		}
	} else if data.NatTcpTimeout.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natProp.tcpTimeout.optionType", "default")
			body, _ = sjson.Set(body, path+"natProp.tcpTimeout.value", 60)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"natProp.tcpTimeout.optionType", "global")
			body, _ = sjson.Set(body, path+"natProp.tcpTimeout.value", data.NatTcpTimeout.ValueInt64())
		}
	}

	if !data.NatUdpTimeoutVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natProp.udpTimeout.optionType", "variable")
			body, _ = sjson.Set(body, path+"natProp.udpTimeout.value", data.NatUdpTimeoutVariable.ValueString())
		}
	} else if data.NatUdpTimeout.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natProp.udpTimeout.optionType", "default")
			body, _ = sjson.Set(body, path+"natProp.udpTimeout.value", 1)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"natProp.udpTimeout.optionType", "global")
			body, _ = sjson.Set(body, path+"natProp.udpTimeout.value", data.NatUdpTimeout.ValueInt64())
		}
	}

	if !data.InterfaceNameVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ethernet.ifName.optionType", "variable")
			body, _ = sjson.Set(body, path+"ethernet.ifName.value", data.InterfaceNameVariable.ValueString())
		}
	} else if !data.InterfaceName.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ethernet.ifName.optionType", "global")
			body, _ = sjson.Set(body, path+"ethernet.ifName.value", data.InterfaceName.ValueString())
		}
	}

	if !data.InterfaceDescriptionVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ethernet.description.optionType", "variable")
			body, _ = sjson.Set(body, path+"ethernet.description.value", data.InterfaceDescriptionVariable.ValueString())
		}
	} else if data.InterfaceDescription.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ethernet.description.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ethernet.description.optionType", "global")
			body, _ = sjson.Set(body, path+"ethernet.description.value", data.InterfaceDescription.ValueString())
		}
	}

	if !data.VlanIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ethernet.vlanId.optionType", "variable")
			body, _ = sjson.Set(body, path+"ethernet.vlanId.value", data.VlanIdVariable.ValueString())
		}
	} else if data.VlanId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ethernet.vlanId.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ethernet.vlanId.optionType", "global")
			body, _ = sjson.Set(body, path+"ethernet.vlanId.value", data.VlanId.ValueInt64())
		}
	}

	if !data.Ipv4DhcpDistanceVariable.IsNull() {
		if true && data.Ipv4AddressType.ValueString() == "dynamic" {
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.dynamic.dynamicDhcpDistance.optionType", "variable")
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.dynamic.dynamicDhcpDistance.value", data.Ipv4DhcpDistanceVariable.ValueString())
		}
	} else if data.Ipv4DhcpDistance.IsNull() {
		if true && data.Ipv4AddressType.ValueString() == "dynamic" {
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.dynamic.dynamicDhcpDistance.optionType", "default")
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.dynamic.dynamicDhcpDistance.value", 1)
		}
	} else {
		if true && data.Ipv4AddressType.ValueString() == "dynamic" {
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.dynamic.dynamicDhcpDistance.optionType", "global")
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.dynamic.dynamicDhcpDistance.value", data.Ipv4DhcpDistance.ValueInt64())
		}
	}

	if !data.Ipv4DhcpHelperVariable.IsNull() {
		if true && data.Ipv4AddressType.ValueString() == "dynamic" {
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.dynamic.dhcpHelper.optionType", "variable")
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.dynamic.dhcpHelper.value", data.Ipv4DhcpHelperVariable.ValueString())
		}
	} else if !data.Ipv4DhcpHelper.IsNull() {
		if true && data.Ipv4AddressType.ValueString() == "dynamic" {
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.dynamic.dhcpHelper.optionType", "global")
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.dynamic.dhcpHelper.value", data.Ipv4DhcpHelper.ValueString())
		}
	}

	if !data.Ipv4AddressVariable.IsNull() {
		if true && data.Ipv4AddressType.ValueString() == "static" {
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.optionType", "variable")
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.value", data.Ipv4AddressVariable.ValueString())
		}
	} else if !data.Ipv4Address.IsNull() {
		if true && data.Ipv4AddressType.ValueString() == "static" {
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.optionType", "global")
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.value", data.Ipv4Address.ValueString())
		}
	}

	if !data.Ipv4SubnetMaskVariable.IsNull() {
		if true && data.Ipv4AddressType.ValueString() == "static" {
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.optionType", "variable")
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.value", data.Ipv4SubnetMaskVariable.ValueString())
		}
	} else if !data.Ipv4SubnetMask.IsNull() {
		if true && data.Ipv4AddressType.ValueString() == "static" {
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.optionType", "global")
			body, _ = sjson.Set(body, path+"ethernet.intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.value", data.Ipv4SubnetMask.ValueString())
		}
	}
	if data.TunnelInterfaceEnabled.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.tunnelInterface.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.tunnelInterface.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.tunnelInterface.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.tunnelInterface.value", data.TunnelInterfaceEnabled.ValueBool())
		}
	}

	if !data.TunnelPerTunnelQosVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.value", data.TunnelPerTunnelQosVariable.ValueString())
		}
	} else if data.TunnelPerTunnelQos.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.value", data.TunnelPerTunnelQos.ValueBool())
		}
	}

	if !data.TunnelQosModeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.mode.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.mode.value", data.TunnelQosModeVariable.ValueString())
		}
	} else if !data.TunnelQosMode.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.mode.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.mode.value", data.TunnelQosMode.ValueString())
		}
	}

	if !data.TunnelBandwidthPercentVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.value", data.TunnelBandwidthPercentVariable.ValueString())
		}
	} else if data.TunnelBandwidthPercent.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.value", 50)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.bandwidthPercent.value", data.TunnelBandwidthPercent.ValueInt64())
		}
	}

	if !data.TunnelInterfaceBindLoopbackTunnelVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.bind.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.bind.value", data.TunnelInterfaceBindLoopbackTunnelVariable.ValueString())
		}
	} else if data.TunnelInterfaceBindLoopbackTunnel.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.bind.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.bind.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.bind.value", data.TunnelInterfaceBindLoopbackTunnel.ValueString())
		}
	}

	if !data.TunnelInterfaceCarrierVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.carrier.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.carrier.value", data.TunnelInterfaceCarrierVariable.ValueString())
		}
	} else if data.TunnelInterfaceCarrier.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.carrier.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.carrier.value", "default")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.carrier.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.carrier.value", data.TunnelInterfaceCarrier.ValueString())
		}
	}

	if !data.TunnelInterfaceColorVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.color.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.color.value", data.TunnelInterfaceColorVariable.ValueString())
		}
	} else if data.TunnelInterfaceColor.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.color.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.color.value", "default")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.color.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.color.value", data.TunnelInterfaceColor.ValueString())
		}
	}

	if !data.TunnelInterfaceColorDescriptionVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.colorDescription.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.colorDescription.value", data.TunnelInterfaceColorDescriptionVariable.ValueString())
		}
	} else if data.TunnelInterfaceColorDescription.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.colorDescription.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.colorDescription.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.colorDescription.value", data.TunnelInterfaceColorDescription.ValueString())
		}
	}

	if !data.TunnelInterfaceHelloIntervalVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.helloInterval.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.helloInterval.value", data.TunnelInterfaceHelloIntervalVariable.ValueString())
		}
	} else if data.TunnelInterfaceHelloInterval.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.helloInterval.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.helloInterval.value", 1000)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.helloInterval.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.helloInterval.value", data.TunnelInterfaceHelloInterval.ValueInt64())
		}
	}

	if !data.TunnelInterfaceHelloToleranceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.helloTolerance.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.helloTolerance.value", data.TunnelInterfaceHelloToleranceVariable.ValueString())
		}
	} else if data.TunnelInterfaceHelloTolerance.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.helloTolerance.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.helloTolerance.value", 12)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.helloTolerance.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.helloTolerance.value", data.TunnelInterfaceHelloTolerance.ValueInt64())
		}
	}

	if !data.TunnelInterfaceLastResortCircuitVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.lastResortCircuit.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.lastResortCircuit.value", data.TunnelInterfaceLastResortCircuitVariable.ValueString())
		}
	} else if data.TunnelInterfaceLastResortCircuit.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.lastResortCircuit.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.lastResortCircuit.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.lastResortCircuit.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.lastResortCircuit.value", data.TunnelInterfaceLastResortCircuit.ValueBool())
		}
	}

	if !data.TunnelInterfaceColorRestrictVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.restrict.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.restrict.value", data.TunnelInterfaceColorRestrictVariable.ValueString())
		}
	} else if data.TunnelInterfaceColorRestrict.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.restrict.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.restrict.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.restrict.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.restrict.value", data.TunnelInterfaceColorRestrict.ValueBool())
		}
	}

	if !data.TunnelInterfaceGroupsVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.group.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.group.value", data.TunnelInterfaceGroupsVariable.ValueString())
		}
	} else if data.TunnelInterfaceGroups.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.group.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.group.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.group.value", data.TunnelInterfaceGroups.ValueString())
		}
	}

	if !data.TunnelInterfaceBorderVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.border.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.border.value", data.TunnelInterfaceBorderVariable.ValueString())
		}
	} else if data.TunnelInterfaceBorder.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.border.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.border.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.border.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.border.value", data.TunnelInterfaceBorder.ValueBool())
		}
	}

	if !data.TunnelInterfaceMaxControlConnectionsVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.maxControlConnections.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.maxControlConnections.value", data.TunnelInterfaceMaxControlConnectionsVariable.ValueString())
		}
	} else if data.TunnelInterfaceMaxControlConnections.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.maxControlConnections.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.maxControlConnections.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.maxControlConnections.value", data.TunnelInterfaceMaxControlConnections.ValueInt64())
		}
	}

	if !data.TunnelInterfaceVbondAsStunServerVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.vbondAsStunServer.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.vbondAsStunServer.value", data.TunnelInterfaceVbondAsStunServerVariable.ValueString())
		}
	} else if data.TunnelInterfaceVbondAsStunServer.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.vbondAsStunServer.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.vbondAsStunServer.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.vbondAsStunServer.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.vbondAsStunServer.value", data.TunnelInterfaceVbondAsStunServer.ValueBool())
		}
	}

	if !data.TunnelInterfaceExcludeControllerGroupListVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.excludeControllerGroupList.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.excludeControllerGroupList.value", data.TunnelInterfaceExcludeControllerGroupListVariable.ValueString())
		}
	} else if data.TunnelInterfaceExcludeControllerGroupList.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.excludeControllerGroupList.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.excludeControllerGroupList.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.excludeControllerGroupList.value", data.TunnelInterfaceExcludeControllerGroupList.ValueString())
		}
	}

	if !data.TunnelInterfaceVmanageConnectionPreferenceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.vmanageConnectionPreference.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.vmanageConnectionPreference.value", data.TunnelInterfaceVmanageConnectionPreferenceVariable.ValueString())
		}
	} else if data.TunnelInterfaceVmanageConnectionPreference.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.vmanageConnectionPreference.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.vmanageConnectionPreference.value", 5)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.vmanageConnectionPreference.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.vmanageConnectionPreference.value", data.TunnelInterfaceVmanageConnectionPreference.ValueInt64())
		}
	}

	if !data.TunnelInterfaceFullPortHopVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.fullPortHop.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.fullPortHop.value", data.TunnelInterfaceFullPortHopVariable.ValueString())
		}
	} else if data.TunnelInterfaceFullPortHop.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.fullPortHop.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.fullPortHop.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.fullPortHop.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.fullPortHop.value", data.TunnelInterfaceFullPortHop.ValueBool())
		}
	}

	if !data.TunnelInterfacePortHopVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.portHop.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.portHop.value", data.TunnelInterfacePortHopVariable.ValueString())
		}
	} else if data.TunnelInterfacePortHop.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.portHop.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.portHop.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.portHop.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.portHop.value", data.TunnelInterfacePortHop.ValueBool())
		}
	}

	if !data.TunnelInterfaceLowBandwidthLinkVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.value", data.TunnelInterfaceLowBandwidthLinkVariable.ValueString())
		}
	} else if data.TunnelInterfaceLowBandwidthLink.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.lowBandwidthLink.value", data.TunnelInterfaceLowBandwidthLink.ValueBool())
		}
	}

	if !data.TunnelInterfaceNetworkBroadcastVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.value", data.TunnelInterfaceNetworkBroadcastVariable.ValueString())
		}
	} else if data.TunnelInterfaceNetworkBroadcast.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.value", data.TunnelInterfaceNetworkBroadcast.ValueBool())
		}
	}

	if !data.TunnelInterfaceTunnelTcpMssVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.tunnelTcpMssAdjust.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.tunnelTcpMssAdjust.value", data.TunnelInterfaceTunnelTcpMssVariable.ValueString())
		}
	} else if data.TunnelInterfaceTunnelTcpMss.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.tunnelTcpMssAdjust.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.tunnelTcpMssAdjust.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.tunnelTcpMssAdjust.value", data.TunnelInterfaceTunnelTcpMss.ValueInt64())
		}
	}

	if !data.TunnelInterfaceClearDontFragmentVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.value", data.TunnelInterfaceClearDontFragmentVariable.ValueString())
		}
	} else if data.TunnelInterfaceClearDontFragment.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.clearDontFragment.value", data.TunnelInterfaceClearDontFragment.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowAllVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.all.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAllowService.all.value", data.TunnelInterfaceAllowAllVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowAll.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.all.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAllowService.all.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.all.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAllowService.all.value", data.TunnelInterfaceAllowAll.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowBgpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.bgp.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAllowService.bgp.value", data.TunnelInterfaceAllowBgpVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowBgp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.bgp.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAllowService.bgp.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.bgp.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAllowService.bgp.value", data.TunnelInterfaceAllowBgp.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowDhcpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.dhcp.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAllowService.dhcp.value", data.TunnelInterfaceAllowDhcpVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowDhcp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.dhcp.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAllowService.dhcp.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.dhcp.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAllowService.dhcp.value", data.TunnelInterfaceAllowDhcp.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowDnsVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.dns.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAllowService.dns.value", data.TunnelInterfaceAllowDnsVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowDns.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.dns.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAllowService.dns.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.dns.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAllowService.dns.value", data.TunnelInterfaceAllowDns.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowIcmpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.icmp.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAllowService.icmp.value", data.TunnelInterfaceAllowIcmpVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowIcmp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.icmp.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAllowService.icmp.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.icmp.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAllowService.icmp.value", data.TunnelInterfaceAllowIcmp.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowSshVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.sshd.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAllowService.sshd.value", data.TunnelInterfaceAllowSshVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowSsh.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.sshd.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAllowService.sshd.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.sshd.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAllowService.sshd.value", data.TunnelInterfaceAllowSsh.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowNtpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.ntp.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAllowService.ntp.value", data.TunnelInterfaceAllowNtpVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowNtp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.ntp.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAllowService.ntp.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.ntp.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAllowService.ntp.value", data.TunnelInterfaceAllowNtp.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowStunVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.stun.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAllowService.stun.value", data.TunnelInterfaceAllowStunVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowStun.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.stun.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAllowService.stun.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.stun.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAllowService.stun.value", data.TunnelInterfaceAllowStun.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowNetconfVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.netconf.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAllowService.netconf.value", data.TunnelInterfaceAllowNetconfVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowNetconf.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.netconf.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAllowService.netconf.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.netconf.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAllowService.netconf.value", data.TunnelInterfaceAllowNetconf.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowOspfVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.ospf.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAllowService.ospf.value", data.TunnelInterfaceAllowOspfVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowOspf.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.ospf.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAllowService.ospf.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.ospf.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAllowService.ospf.value", data.TunnelInterfaceAllowOspf.ValueBool())
		}
	}

	if !data.TunnelInterfaceAllowSnmpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.snmp.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAllowService.snmp.value", data.TunnelInterfaceAllowSnmpVariable.ValueString())
		}
	} else if data.TunnelInterfaceAllowSnmp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.snmp.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAllowService.snmp.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAllowService.snmp.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAllowService.snmp.value", data.TunnelInterfaceAllowSnmp.ValueBool())
		}
	}
	if data.TunnelInterfaceIpsecEncapsulation.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.ipsecEncap.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.ipsecEncap.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.ipsecEncap.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.ipsecEncap.value", data.TunnelInterfaceIpsecEncapsulation.ValueBool())
		}
	}

	if !data.TunnelInterfaceIpsecPreferenceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.ipsecPreference.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.ipsecPreference.value", data.TunnelInterfaceIpsecPreferenceVariable.ValueString())
		}
	} else if data.TunnelInterfaceIpsecPreference.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.ipsecPreference.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.ipsecPreference.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.ipsecPreference.value", data.TunnelInterfaceIpsecPreference.ValueInt64())
		}
	}

	if !data.TunnelInterfaceIpsecWeightVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.ipsecWeight.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.ipsecWeight.value", data.TunnelInterfaceIpsecWeightVariable.ValueString())
		}
	} else if data.TunnelInterfaceIpsecWeight.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.ipsecWeight.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.ipsecWeight.value", 1)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.ipsecWeight.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.ipsecWeight.value", data.TunnelInterfaceIpsecWeight.ValueInt64())
		}
	}
	if data.TunnelInterfaceGreEncapsulation.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.greEncap.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.greEncap.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.greEncap.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.greEncap.value", data.TunnelInterfaceGreEncapsulation.ValueBool())
		}
	}

	if !data.TunnelInterfaceGrePreferenceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.grePreference.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.grePreference.value", data.TunnelInterfaceGrePreferenceVariable.ValueString())
		}
	} else if data.TunnelInterfaceGrePreference.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.grePreference.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.grePreference.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.grePreference.value", data.TunnelInterfaceGrePreference.ValueInt64())
		}
	}

	if !data.TunnelInterfaceGreWeightVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.greWeight.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.greWeight.value", data.TunnelInterfaceGreWeightVariable.ValueString())
		}
	} else if data.TunnelInterfaceGreWeight.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.greWeight.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.greWeight.value", 1)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.greWeight.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.greWeight.value", data.TunnelInterfaceGreWeight.ValueInt64())
		}
	}

	if !data.TunnelInterfaceNatRefreshIntervalVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.natRefreshInterval.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.natRefreshInterval.value", data.TunnelInterfaceNatRefreshIntervalVariable.ValueString())
		}
	} else if data.TunnelInterfaceNatRefreshInterval.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.natRefreshInterval.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.natRefreshInterval.value", 5)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.natRefreshInterval.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnelAdvancedOption.natRefreshInterval.value", data.TunnelInterfaceNatRefreshInterval.ValueInt64())
		}
	}
	if data.MultiRegionFabricEnableCoreRegion.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"multiRegionFabric.enableCoreRegion.optionType", "default")
			body, _ = sjson.Set(body, path+"multiRegionFabric.enableCoreRegion.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"multiRegionFabric.enableCoreRegion.optionType", "global")
			body, _ = sjson.Set(body, path+"multiRegionFabric.enableCoreRegion.value", data.MultiRegionFabricEnableCoreRegion.ValueBool())
		}
	}
	if data.MultiRegionFabricCoreRegion.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"multiRegionFabric.coreRegion.optionType", "default")
			body, _ = sjson.Set(body, path+"multiRegionFabric.coreRegion.value", "core-shared")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"multiRegionFabric.coreRegion.optionType", "global")
			body, _ = sjson.Set(body, path+"multiRegionFabric.coreRegion.value", data.MultiRegionFabricCoreRegion.ValueString())
		}
	}
	if data.MultiRegionFabricEnableSecondaryRegion.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"multiRegionFabric.enableSecondaryRegion.optionType", "default")
			body, _ = sjson.Set(body, path+"multiRegionFabric.enableSecondaryRegion.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"multiRegionFabric.enableSecondaryRegion.optionType", "global")
			body, _ = sjson.Set(body, path+"multiRegionFabric.enableSecondaryRegion.value", data.MultiRegionFabricEnableSecondaryRegion.ValueBool())
		}
	}
	if data.MultiRegionFabricSecondaryRegion.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"multiRegionFabric.secondaryRegion.optionType", "default")
			body, _ = sjson.Set(body, path+"multiRegionFabric.secondaryRegion.value", "secondary-shared")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"multiRegionFabric.secondaryRegion.optionType", "global")
			body, _ = sjson.Set(body, path+"multiRegionFabric.secondaryRegion.value", data.MultiRegionFabricSecondaryRegion.ValueString())
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
	if data.AdaptiveQos.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.adaptiveQoS.optionType", "default")
			body, _ = sjson.Set(body, path+"aclQos.adaptiveQoS.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.adaptiveQoS.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.adaptiveQoS.value", data.AdaptiveQos.ValueBool())
		}
	}

	if !data.AdaptiveQosPeriodVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.adaptPeriod.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.adaptPeriod.value", data.AdaptiveQosPeriodVariable.ValueString())
		}
	} else if data.AdaptiveQosPeriod.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.adaptPeriod.optionType", "default")
			body, _ = sjson.Set(body, path+"aclQos.adaptPeriod.value", 15)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.adaptPeriod.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.adaptPeriod.value", data.AdaptiveQosPeriod.ValueInt64())
		}
	}
	if data.ShapingRateUpstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstream.optionType", "default")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstream.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstream.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstream.value", data.ShapingRateUpstream.ValueBool())
		}
	}

	if !data.MinShapingRateUpstreamVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.value", data.MinShapingRateUpstreamVariable.ValueString())
		}
	} else if !data.MinShapingRateUpstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.value", data.MinShapingRateUpstream.ValueInt64())
		}
	}

	if !data.MaxShapingRateUpstreamVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.value", data.MaxShapingRateUpstreamVariable.ValueString())
		}
	} else if !data.MaxShapingRateUpstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.value", data.MaxShapingRateUpstream.ValueInt64())
		}
	}

	if !data.DefaultShapingRateUpstreamVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.value", data.DefaultShapingRateUpstreamVariable.ValueString())
		}
	} else if !data.DefaultShapingRateUpstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.value", data.DefaultShapingRateUpstream.ValueInt64())
		}
	}
	if data.ShapingRateDownstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstream.optionType", "default")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstream.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstream.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstream.value", data.ShapingRateDownstream.ValueBool())
		}
	}

	if !data.MinShapingRateDownstreamVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.value", data.MinShapingRateDownstreamVariable.ValueString())
		}
	} else if !data.MinShapingRateDownstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.value", data.MinShapingRateDownstream.ValueInt64())
		}
	}

	if !data.MaxShapingRateDownstreamVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.value", data.MaxShapingRateDownstreamVariable.ValueString())
		}
	} else if !data.MaxShapingRateDownstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.value", data.MaxShapingRateDownstream.ValueInt64())
		}
	}

	if !data.DefaultShapingRateDownstreamVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.value", data.DefaultShapingRateDownstreamVariable.ValueString())
		}
	} else if !data.DefaultShapingRateDownstream.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.value", data.DefaultShapingRateDownstream.ValueInt64())
		}
	}

	if !data.ShapingRateVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRate.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.shapingRate.value", data.ShapingRateVariable.ValueString())
		}
	} else if data.ShapingRate.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRate.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRate.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRate.value", data.ShapingRate.ValueInt64())
		}
	}

	if !data.AclIpv4EgressFeatureIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.ipv4AclEgress.refId.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.ipv4AclEgress.refId.value", data.AclIpv4EgressFeatureIdVariable.ValueString())
		}
	} else if !data.AclIpv4EgressFeatureId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.ipv4AclEgress.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.ipv4AclEgress.refId.value", data.AclIpv4EgressFeatureId.ValueString())
		}
	}

	if !data.AclIpv4IngressFeatureIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.ipv4AclIngress.refId.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.ipv4AclIngress.refId.value", data.AclIpv4IngressFeatureIdVariable.ValueString())
		}
	} else if !data.AclIpv4IngressFeatureId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.ipv4AclIngress.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.ipv4AclIngress.refId.value", data.AclIpv4IngressFeatureId.ValueString())
		}
	}

	if !data.AclIpv6EgressFeatureIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.ipv6AclEgress.refId.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.ipv6AclEgress.refId.value", data.AclIpv6EgressFeatureIdVariable.ValueString())
		}
	} else if !data.AclIpv6EgressFeatureId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.ipv6AclEgress.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.ipv6AclEgress.refId.value", data.AclIpv6EgressFeatureId.ValueString())
		}
	}

	if !data.AclIpv6IngressFeatureIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.ipv6AclIngress.refId.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.ipv6AclIngress.refId.value", data.AclIpv6IngressFeatureIdVariable.ValueString())
		}
	} else if !data.AclIpv6IngressFeatureId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.ipv6AclIngress.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.ipv6AclIngress.refId.value", data.AclIpv6IngressFeatureId.ValueString())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportWANVPNInterfaceDSLIPoE) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
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
	data.VdslSlot = types.StringNull()
	data.VdslSlotVariable = types.StringNull()
	if t := res.Get(path + "vdsl.slot.optionType"); t.Exists() {
		va := res.Get(path + "vdsl.slot.value")
		if t.String() == "variable" {
			data.VdslSlotVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.VdslSlot = types.StringValue(va.String())
		}
	}
	data.VdslMode = types.StringNull()

	if t := res.Get(path + "vdsl.mode.optionType"); t.Exists() {
		va := res.Get(path + "vdsl.mode.value")
		if t.String() == "global" {
			data.VdslMode = types.StringValue(va.String())
		}
	}
	data.VdslSra = types.BoolNull()
	data.VdslSraVariable = types.StringNull()
	if t := res.Get(path + "vdsl.sra.optionType"); t.Exists() {
		va := res.Get(path + "vdsl.sra.value")
		if t.String() == "variable" {
			data.VdslSraVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.VdslSra = types.BoolValue(va.Bool())
		}
	}
	data.NatEnabled = types.BoolNull()
	data.NatEnabledVariable = types.StringNull()
	if t := res.Get(path + "natProp.nat.optionType"); t.Exists() {
		va := res.Get(path + "natProp.nat.value")
		if t.String() == "variable" {
			data.NatEnabledVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatEnabled = types.BoolValue(va.Bool())
		}
	}
	data.NatTcpTimeout = types.Int64Null()
	data.NatTcpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "natProp.tcpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "natProp.tcpTimeout.value")
		if t.String() == "variable" {
			data.NatTcpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatTcpTimeout = types.Int64Value(va.Int())
		}
	}
	data.NatUdpTimeout = types.Int64Null()
	data.NatUdpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "natProp.udpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "natProp.udpTimeout.value")
		if t.String() == "variable" {
			data.NatUdpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NatUdpTimeout = types.Int64Value(va.Int())
		}
	}
	data.InterfaceName = types.StringNull()
	data.InterfaceNameVariable = types.StringNull()
	if t := res.Get(path + "ethernet.ifName.optionType"); t.Exists() {
		va := res.Get(path + "ethernet.ifName.value")
		if t.String() == "variable" {
			data.InterfaceNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceName = types.StringValue(va.String())
		}
	}
	data.InterfaceDescription = types.StringNull()
	data.InterfaceDescriptionVariable = types.StringNull()
	if t := res.Get(path + "ethernet.description.optionType"); t.Exists() {
		va := res.Get(path + "ethernet.description.value")
		if t.String() == "variable" {
			data.InterfaceDescriptionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceDescription = types.StringValue(va.String())
		}
	}
	data.VlanId = types.Int64Null()
	data.VlanIdVariable = types.StringNull()
	if t := res.Get(path + "ethernet.vlanId.optionType"); t.Exists() {
		va := res.Get(path + "ethernet.vlanId.value")
		if t.String() == "variable" {
			data.VlanIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.VlanId = types.Int64Value(va.Int())
		}
	}
	data.Ipv4DhcpDistance = types.Int64Null()
	data.Ipv4DhcpDistanceVariable = types.StringNull()
	if t := res.Get(path + "ethernet.intfIpAddress.dynamic.dynamicDhcpDistance.optionType"); t.Exists() {
		va := res.Get(path + "ethernet.intfIpAddress.dynamic.dynamicDhcpDistance.value")
		if t.String() == "variable" {
			data.Ipv4DhcpDistanceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4DhcpDistance = types.Int64Value(va.Int())
		}
		data.Ipv4AddressType = types.StringValue("dynamic")
	}
	data.Ipv4DhcpHelper = types.StringNull()
	data.Ipv4DhcpHelperVariable = types.StringNull()
	if t := res.Get(path + "ethernet.intfIpAddress.dynamic.dhcpHelper.optionType"); t.Exists() {
		va := res.Get(path + "ethernet.intfIpAddress.dynamic.dhcpHelper.value")
		if t.String() == "variable" {
			data.Ipv4DhcpHelperVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4DhcpHelper = types.StringValue(va.String())
		}
		data.Ipv4AddressType = types.StringValue("dynamic")
	}
	data.Ipv4Address = types.StringNull()
	data.Ipv4AddressVariable = types.StringNull()
	if t := res.Get(path + "ethernet.intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.optionType"); t.Exists() {
		va := res.Get(path + "ethernet.intfIpAddress.static.staticIpV4AddressPrimary.ipAddress.value")
		if t.String() == "variable" {
			data.Ipv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Address = types.StringValue(va.String())
		}
		data.Ipv4AddressType = types.StringValue("static")
	}
	data.Ipv4SubnetMask = types.StringNull()
	data.Ipv4SubnetMaskVariable = types.StringNull()
	if t := res.Get(path + "ethernet.intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.optionType"); t.Exists() {
		va := res.Get(path + "ethernet.intfIpAddress.static.staticIpV4AddressPrimary.subnetMask.value")
		if t.String() == "variable" {
			data.Ipv4SubnetMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4SubnetMask = types.StringValue(va.String())
		}
		data.Ipv4AddressType = types.StringValue("static")
	}
	data.TunnelInterfaceEnabled = types.BoolNull()

	if t := res.Get(path + "tunnel.tunnelInterface.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.tunnelInterface.value")
		if t.String() == "global" {
			data.TunnelInterfaceEnabled = types.BoolValue(va.Bool())
		}
	}
	data.TunnelPerTunnelQos = types.BoolNull()
	data.TunnelPerTunnelQosVariable = types.StringNull()
	if t := res.Get(path + "tunnel.perTunnelQos.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.perTunnelQos.value")
		if t.String() == "variable" {
			data.TunnelPerTunnelQosVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelPerTunnelQos = types.BoolValue(va.Bool())
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
	if t := res.Get(path + "tunnelAdvancedOption.bind.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAdvancedOption.bind.value")
		if t.String() == "variable" {
			data.TunnelInterfaceBindLoopbackTunnelVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceBindLoopbackTunnel = types.StringValue(va.String())
		}
	}
	data.TunnelInterfaceCarrier = types.StringNull()
	data.TunnelInterfaceCarrierVariable = types.StringNull()
	if t := res.Get(path + "tunnelAdvancedOption.carrier.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAdvancedOption.carrier.value")
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
	data.TunnelInterfaceColorDescription = types.StringNull()
	data.TunnelInterfaceColorDescriptionVariable = types.StringNull()
	if t := res.Get(path + "tunnel.colorDescription.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.colorDescription.value")
		if t.String() == "variable" {
			data.TunnelInterfaceColorDescriptionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceColorDescription = types.StringValue(va.String())
		}
	}
	data.TunnelInterfaceHelloInterval = types.Int64Null()
	data.TunnelInterfaceHelloIntervalVariable = types.StringNull()
	if t := res.Get(path + "tunnelAdvancedOption.helloInterval.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAdvancedOption.helloInterval.value")
		if t.String() == "variable" {
			data.TunnelInterfaceHelloIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceHelloInterval = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceHelloTolerance = types.Int64Null()
	data.TunnelInterfaceHelloToleranceVariable = types.StringNull()
	if t := res.Get(path + "tunnelAdvancedOption.helloTolerance.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAdvancedOption.helloTolerance.value")
		if t.String() == "variable" {
			data.TunnelInterfaceHelloToleranceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceHelloTolerance = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceLastResortCircuit = types.BoolNull()
	data.TunnelInterfaceLastResortCircuitVariable = types.StringNull()
	if t := res.Get(path + "tunnelAdvancedOption.lastResortCircuit.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAdvancedOption.lastResortCircuit.value")
		if t.String() == "variable" {
			data.TunnelInterfaceLastResortCircuitVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceLastResortCircuit = types.BoolValue(va.Bool())
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
	data.TunnelInterfaceGroups = types.StringNull()
	data.TunnelInterfaceGroupsVariable = types.StringNull()
	if t := res.Get(path + "tunnel.group.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.group.value")
		if t.String() == "variable" {
			data.TunnelInterfaceGroupsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceGroups = types.StringValue(va.String())
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
	data.TunnelInterfaceVbondAsStunServer = types.BoolNull()
	data.TunnelInterfaceVbondAsStunServerVariable = types.StringNull()
	if t := res.Get(path + "tunnel.vbondAsStunServer.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.vbondAsStunServer.value")
		if t.String() == "variable" {
			data.TunnelInterfaceVbondAsStunServerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceVbondAsStunServer = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceExcludeControllerGroupList = types.StringNull()
	data.TunnelInterfaceExcludeControllerGroupListVariable = types.StringNull()
	if t := res.Get(path + "tunnel.excludeControllerGroupList.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.excludeControllerGroupList.value")
		if t.String() == "variable" {
			data.TunnelInterfaceExcludeControllerGroupListVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceExcludeControllerGroupList = types.StringValue(va.String())
		}
	}
	data.TunnelInterfaceVmanageConnectionPreference = types.Int64Null()
	data.TunnelInterfaceVmanageConnectionPreferenceVariable = types.StringNull()
	if t := res.Get(path + "tunnel.vmanageConnectionPreference.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.vmanageConnectionPreference.value")
		if t.String() == "variable" {
			data.TunnelInterfaceVmanageConnectionPreferenceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceVmanageConnectionPreference = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceFullPortHop = types.BoolNull()
	data.TunnelInterfaceFullPortHopVariable = types.StringNull()
	if t := res.Get(path + "tunnel.fullPortHop.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.fullPortHop.value")
		if t.String() == "variable" {
			data.TunnelInterfaceFullPortHopVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceFullPortHop = types.BoolValue(va.Bool())
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
	data.TunnelInterfaceTunnelTcpMss = types.Int64Null()
	data.TunnelInterfaceTunnelTcpMssVariable = types.StringNull()
	if t := res.Get(path + "tunnel.tunnelTcpMssAdjust.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.tunnelTcpMssAdjust.value")
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
	data.TunnelInterfaceAllowAll = types.BoolNull()
	data.TunnelInterfaceAllowAllVariable = types.StringNull()
	if t := res.Get(path + "tunnelAllowService.all.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAllowService.all.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowAllVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowAll = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowBgp = types.BoolNull()
	data.TunnelInterfaceAllowBgpVariable = types.StringNull()
	if t := res.Get(path + "tunnelAllowService.bgp.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAllowService.bgp.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowBgpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowBgp = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowDhcp = types.BoolNull()
	data.TunnelInterfaceAllowDhcpVariable = types.StringNull()
	if t := res.Get(path + "tunnelAllowService.dhcp.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAllowService.dhcp.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowDhcpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowDhcp = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowDns = types.BoolNull()
	data.TunnelInterfaceAllowDnsVariable = types.StringNull()
	if t := res.Get(path + "tunnelAllowService.dns.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAllowService.dns.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowDnsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowDns = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowIcmp = types.BoolNull()
	data.TunnelInterfaceAllowIcmpVariable = types.StringNull()
	if t := res.Get(path + "tunnelAllowService.icmp.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAllowService.icmp.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowIcmpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowIcmp = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowSsh = types.BoolNull()
	data.TunnelInterfaceAllowSshVariable = types.StringNull()
	if t := res.Get(path + "tunnelAllowService.sshd.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAllowService.sshd.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowSshVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowSsh = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowNtp = types.BoolNull()
	data.TunnelInterfaceAllowNtpVariable = types.StringNull()
	if t := res.Get(path + "tunnelAllowService.ntp.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAllowService.ntp.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowNtpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowNtp = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowStun = types.BoolNull()
	data.TunnelInterfaceAllowStunVariable = types.StringNull()
	if t := res.Get(path + "tunnelAllowService.stun.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAllowService.stun.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowStunVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowStun = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowNetconf = types.BoolNull()
	data.TunnelInterfaceAllowNetconfVariable = types.StringNull()
	if t := res.Get(path + "tunnelAllowService.netconf.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAllowService.netconf.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowNetconfVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowNetconf = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowOspf = types.BoolNull()
	data.TunnelInterfaceAllowOspfVariable = types.StringNull()
	if t := res.Get(path + "tunnelAllowService.ospf.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAllowService.ospf.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowOspfVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowOspf = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceAllowSnmp = types.BoolNull()
	data.TunnelInterfaceAllowSnmpVariable = types.StringNull()
	if t := res.Get(path + "tunnelAllowService.snmp.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAllowService.snmp.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowSnmpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowSnmp = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceIpsecEncapsulation = types.BoolNull()

	if t := res.Get(path + "tunnelAdvancedOption.ipsecEncap.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAdvancedOption.ipsecEncap.value")
		if t.String() == "global" {
			data.TunnelInterfaceIpsecEncapsulation = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceIpsecPreference = types.Int64Null()
	data.TunnelInterfaceIpsecPreferenceVariable = types.StringNull()
	if t := res.Get(path + "tunnelAdvancedOption.ipsecPreference.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAdvancedOption.ipsecPreference.value")
		if t.String() == "variable" {
			data.TunnelInterfaceIpsecPreferenceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceIpsecPreference = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceIpsecWeight = types.Int64Null()
	data.TunnelInterfaceIpsecWeightVariable = types.StringNull()
	if t := res.Get(path + "tunnelAdvancedOption.ipsecWeight.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAdvancedOption.ipsecWeight.value")
		if t.String() == "variable" {
			data.TunnelInterfaceIpsecWeightVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceIpsecWeight = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceGreEncapsulation = types.BoolNull()

	if t := res.Get(path + "tunnelAdvancedOption.greEncap.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAdvancedOption.greEncap.value")
		if t.String() == "global" {
			data.TunnelInterfaceGreEncapsulation = types.BoolValue(va.Bool())
		}
	}
	data.TunnelInterfaceGrePreference = types.Int64Null()
	data.TunnelInterfaceGrePreferenceVariable = types.StringNull()
	if t := res.Get(path + "tunnelAdvancedOption.grePreference.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAdvancedOption.grePreference.value")
		if t.String() == "variable" {
			data.TunnelInterfaceGrePreferenceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceGrePreference = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceGreWeight = types.Int64Null()
	data.TunnelInterfaceGreWeightVariable = types.StringNull()
	if t := res.Get(path + "tunnelAdvancedOption.greWeight.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAdvancedOption.greWeight.value")
		if t.String() == "variable" {
			data.TunnelInterfaceGreWeightVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceGreWeight = types.Int64Value(va.Int())
		}
	}
	data.TunnelInterfaceNatRefreshInterval = types.Int64Null()
	data.TunnelInterfaceNatRefreshIntervalVariable = types.StringNull()
	if t := res.Get(path + "tunnelAdvancedOption.natRefreshInterval.optionType"); t.Exists() {
		va := res.Get(path + "tunnelAdvancedOption.natRefreshInterval.value")
		if t.String() == "variable" {
			data.TunnelInterfaceNatRefreshIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceNatRefreshInterval = types.Int64Value(va.Int())
		}
	}
	data.MultiRegionFabricEnableCoreRegion = types.BoolNull()

	if t := res.Get(path + "multiRegionFabric.enableCoreRegion.optionType"); t.Exists() {
		va := res.Get(path + "multiRegionFabric.enableCoreRegion.value")
		if t.String() == "global" {
			data.MultiRegionFabricEnableCoreRegion = types.BoolValue(va.Bool())
		}
	}
	data.MultiRegionFabricCoreRegion = types.StringNull()

	if t := res.Get(path + "multiRegionFabric.coreRegion.optionType"); t.Exists() {
		va := res.Get(path + "multiRegionFabric.coreRegion.value")
		if t.String() == "global" {
			data.MultiRegionFabricCoreRegion = types.StringValue(va.String())
		}
	}
	data.MultiRegionFabricEnableSecondaryRegion = types.BoolNull()

	if t := res.Get(path + "multiRegionFabric.enableSecondaryRegion.optionType"); t.Exists() {
		va := res.Get(path + "multiRegionFabric.enableSecondaryRegion.value")
		if t.String() == "global" {
			data.MultiRegionFabricEnableSecondaryRegion = types.BoolValue(va.Bool())
		}
	}
	data.MultiRegionFabricSecondaryRegion = types.StringNull()

	if t := res.Get(path + "multiRegionFabric.secondaryRegion.optionType"); t.Exists() {
		va := res.Get(path + "multiRegionFabric.secondaryRegion.value")
		if t.String() == "global" {
			data.MultiRegionFabricSecondaryRegion = types.StringValue(va.String())
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
	data.AdaptiveQos = types.BoolNull()

	if t := res.Get(path + "aclQos.adaptiveQoS.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.adaptiveQoS.value")
		if t.String() == "global" {
			data.AdaptiveQos = types.BoolValue(va.Bool())
		}
	}
	data.AdaptiveQosPeriod = types.Int64Null()
	data.AdaptiveQosPeriodVariable = types.StringNull()
	if t := res.Get(path + "aclQos.adaptPeriod.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.adaptPeriod.value")
		if t.String() == "variable" {
			data.AdaptiveQosPeriodVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdaptiveQosPeriod = types.Int64Value(va.Int())
		}
	}
	data.ShapingRateUpstream = types.BoolNull()

	if t := res.Get(path + "aclQos.shapingRateUpstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateUpstream.value")
		if t.String() == "global" {
			data.ShapingRateUpstream = types.BoolValue(va.Bool())
		}
	}
	data.MinShapingRateUpstream = types.Int64Null()
	data.MinShapingRateUpstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateUpstreamConfig.minShapingRateUpstream.value")
		if t.String() == "variable" {
			data.MinShapingRateUpstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MinShapingRateUpstream = types.Int64Value(va.Int())
		}
	}
	data.MaxShapingRateUpstream = types.Int64Null()
	data.MaxShapingRateUpstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateUpstreamConfig.maxShapingRateUpstream.value")
		if t.String() == "variable" {
			data.MaxShapingRateUpstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MaxShapingRateUpstream = types.Int64Value(va.Int())
		}
	}
	data.DefaultShapingRateUpstream = types.Int64Null()
	data.DefaultShapingRateUpstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateUpstreamConfig.defaultShapingRateUpstream.value")
		if t.String() == "variable" {
			data.DefaultShapingRateUpstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultShapingRateUpstream = types.Int64Value(va.Int())
		}
	}
	data.ShapingRateDownstream = types.BoolNull()

	if t := res.Get(path + "aclQos.shapingRateDownstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateDownstream.value")
		if t.String() == "global" {
			data.ShapingRateDownstream = types.BoolValue(va.Bool())
		}
	}
	data.MinShapingRateDownstream = types.Int64Null()
	data.MinShapingRateDownstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateDownstreamConfig.minShapingRateDownstream.value")
		if t.String() == "variable" {
			data.MinShapingRateDownstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MinShapingRateDownstream = types.Int64Value(va.Int())
		}
	}
	data.MaxShapingRateDownstream = types.Int64Null()
	data.MaxShapingRateDownstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateDownstreamConfig.maxShapingRateDownstream.value")
		if t.String() == "variable" {
			data.MaxShapingRateDownstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MaxShapingRateDownstream = types.Int64Value(va.Int())
		}
	}
	data.DefaultShapingRateDownstream = types.Int64Null()
	data.DefaultShapingRateDownstreamVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRateDownstreamConfig.defaultShapingRateDownstream.value")
		if t.String() == "variable" {
			data.DefaultShapingRateDownstreamVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultShapingRateDownstream = types.Int64Value(va.Int())
		}
	}
	data.ShapingRate = types.Int64Null()
	data.ShapingRateVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRate.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRate.value")
		if t.String() == "variable" {
			data.ShapingRateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ShapingRate = types.Int64Value(va.Int())
		}
	}
	data.AclIpv4EgressFeatureId = types.StringNull()
	data.AclIpv4EgressFeatureIdVariable = types.StringNull()
	if t := res.Get(path + "aclQos.ipv4AclEgress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv4AclEgress.refId.value")
		if t.String() == "variable" {
			data.AclIpv4EgressFeatureIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AclIpv4EgressFeatureId = types.StringValue(va.String())
		}
	}
	data.AclIpv4IngressFeatureId = types.StringNull()
	data.AclIpv4IngressFeatureIdVariable = types.StringNull()
	if t := res.Get(path + "aclQos.ipv4AclIngress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv4AclIngress.refId.value")
		if t.String() == "variable" {
			data.AclIpv4IngressFeatureIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AclIpv4IngressFeatureId = types.StringValue(va.String())
		}
	}
	data.AclIpv6EgressFeatureId = types.StringNull()
	data.AclIpv6EgressFeatureIdVariable = types.StringNull()
	if t := res.Get(path + "aclQos.ipv6AclEgress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv6AclEgress.refId.value")
		if t.String() == "variable" {
			data.AclIpv6EgressFeatureIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AclIpv6EgressFeatureId = types.StringValue(va.String())
		}
	}
	data.AclIpv6IngressFeatureId = types.StringNull()
	data.AclIpv6IngressFeatureIdVariable = types.StringNull()
	if t := res.Get(path + "aclQos.ipv6AclIngress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv6AclIngress.refId.value")
		if t.String() == "variable" {
			data.AclIpv6IngressFeatureIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AclIpv6IngressFeatureId = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end fromBody
