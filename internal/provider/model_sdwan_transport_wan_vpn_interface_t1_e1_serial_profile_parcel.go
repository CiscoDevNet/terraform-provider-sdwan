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
type TransportWANVPNInterfaceT1E1Serial struct {
	Id                                                 types.String                                                      `tfsdk:"id"`
	Version                                            types.Int64                                                       `tfsdk:"version"`
	Name                                               types.String                                                      `tfsdk:"name"`
	Description                                        types.String                                                      `tfsdk:"description"`
	FeatureProfileId                                   types.String                                                      `tfsdk:"feature_profile_id"`
	TransportWanVpnFeatureId                           types.String                                                      `tfsdk:"transport_wan_vpn_feature_id"`
	Shutdown                                           types.Bool                                                        `tfsdk:"shutdown"`
	ShutdownVariable                                   types.String                                                      `tfsdk:"shutdown_variable"`
	InterfaceName                                      types.String                                                      `tfsdk:"interface_name"`
	InterfaceNameVariable                              types.String                                                      `tfsdk:"interface_name_variable"`
	Ipv4Address                                        types.String                                                      `tfsdk:"ipv4_address"`
	Ipv4AddressVariable                                types.String                                                      `tfsdk:"ipv4_address_variable"`
	Ipv4SubnetMask                                     types.String                                                      `tfsdk:"ipv4_subnet_mask"`
	Ipv4SubnetMaskVariable                             types.String                                                      `tfsdk:"ipv4_subnet_mask_variable"`
	Ipv6Address                                        types.String                                                      `tfsdk:"ipv6_address"`
	Ipv6AddressVariable                                types.String                                                      `tfsdk:"ipv6_address_variable"`
	Bandwidth                                          types.Int64                                                       `tfsdk:"bandwidth"`
	BandwidthVariable                                  types.String                                                      `tfsdk:"bandwidth_variable"`
	BandwidthDownstream                                types.Int64                                                       `tfsdk:"bandwidth_downstream"`
	BandwidthDownstreamVariable                        types.String                                                      `tfsdk:"bandwidth_downstream_variable"`
	ClockRate                                          types.String                                                      `tfsdk:"clock_rate"`
	ClockRateVariable                                  types.String                                                      `tfsdk:"clock_rate_variable"`
	Encapsulation                                      types.String                                                      `tfsdk:"encapsulation"`
	EncapsulationVariable                              types.String                                                      `tfsdk:"encapsulation_variable"`
	TunnelInterface                                    types.Bool                                                        `tfsdk:"tunnel_interface"`
	PerTunnelQos                                       types.Bool                                                        `tfsdk:"per_tunnel_qos"`
	PerTunnelQosVariable                               types.String                                                      `tfsdk:"per_tunnel_qos_variable"`
	PerTunnelQosAggregator                             types.Bool                                                        `tfsdk:"per_tunnel_qos_aggregator"`
	PerTunnelQosAggregatorVariable                     types.String                                                      `tfsdk:"per_tunnel_qos_aggregator_variable"`
	TunnelQosMode                                      types.String                                                      `tfsdk:"tunnel_qos_mode"`
	TunnelQosModeVariable                              types.String                                                      `tfsdk:"tunnel_qos_mode_variable"`
	TunnelInterfaceColor                               types.String                                                      `tfsdk:"tunnel_interface_color"`
	TunnelInterfaceColorVariable                       types.String                                                      `tfsdk:"tunnel_interface_color_variable"`
	TunnelInterfaceRestrict                            types.Bool                                                        `tfsdk:"tunnel_interface_restrict"`
	TunnelInterfaceRestrictVariable                    types.String                                                      `tfsdk:"tunnel_interface_restrict_variable"`
	TunnelInterfaceGroups                              types.Int64                                                       `tfsdk:"tunnel_interface_groups"`
	TunnelInterfaceGroupsVariable                      types.String                                                      `tfsdk:"tunnel_interface_groups_variable"`
	TunnelInterfaceBorder                              types.Bool                                                        `tfsdk:"tunnel_interface_border"`
	TunnelInterfaceBorderVariable                      types.String                                                      `tfsdk:"tunnel_interface_border_variable"`
	TunnelInterfaceMaxControlConnections               types.Int64                                                       `tfsdk:"tunnel_interface_max_control_connections"`
	TunnelInterfaceMaxControlConnectionsVariable       types.String                                                      `tfsdk:"tunnel_interface_max_control_connections_variable"`
	TunnelInterfaceVbondAsStunServer                   types.Bool                                                        `tfsdk:"tunnel_interface_vbond_as_stun_server"`
	TunnelInterfaceVbondAsStunServerVariable           types.String                                                      `tfsdk:"tunnel_interface_vbond_as_stun_server_variable"`
	TunnelInterfaceExcludeControllerGroupList          types.Set                                                         `tfsdk:"tunnel_interface_exclude_controller_group_list"`
	TunnelInterfaceExcludeControllerGroupListVariable  types.String                                                      `tfsdk:"tunnel_interface_exclude_controller_group_list_variable"`
	TunnelInterfaceVmanageConnectionPreference         types.Int64                                                       `tfsdk:"tunnel_interface_vmanage_connection_preference"`
	TunnelInterfaceVmanageConnectionPreferenceVariable types.String                                                      `tfsdk:"tunnel_interface_vmanage_connection_preference_variable"`
	TunnelInterfacePortHop                             types.Bool                                                        `tfsdk:"tunnel_interface_port_hop"`
	TunnelInterfacePortHopVariable                     types.String                                                      `tfsdk:"tunnel_interface_port_hop_variable"`
	TunnelInterfaceLowBandwidthLink                    types.Bool                                                        `tfsdk:"tunnel_interface_low_bandwidth_link"`
	TunnelInterfaceLowBandwidthLinkVariable            types.String                                                      `tfsdk:"tunnel_interface_low_bandwidth_link_variable"`
	TunnelInterfaceTunnelTcpMss                        types.Int64                                                       `tfsdk:"tunnel_interface_tunnel_tcp_mss"`
	TunnelInterfaceTunnelTcpMssVariable                types.String                                                      `tfsdk:"tunnel_interface_tunnel_tcp_mss_variable"`
	TunnelInterfaceClearDontFragment                   types.Bool                                                        `tfsdk:"tunnel_interface_clear_dont_fragment"`
	TunnelInterfaceClearDontFragmentVariable           types.String                                                      `tfsdk:"tunnel_interface_clear_dont_fragment_variable"`
	TunnelInterfaceClearNetworkBroadcast               types.Bool                                                        `tfsdk:"tunnel_interface_clear_network_broadcast"`
	TunnelInterfaceClearNetworkBroadcastVariable       types.String                                                      `tfsdk:"tunnel_interface_clear_network_broadcast_variable"`
	TunnelInterfaceCarrier                             types.String                                                      `tfsdk:"tunnel_interface_carrier"`
	TunnelInterfaceCarrierVariable                     types.String                                                      `tfsdk:"tunnel_interface_carrier_variable"`
	TunnelInterfaceBindLoopbackTunnel                  types.String                                                      `tfsdk:"tunnel_interface_bind_loopback_tunnel"`
	TunnelInterfaceBindLoopbackTunnelVariable          types.String                                                      `tfsdk:"tunnel_interface_bind_loopback_tunnel_variable"`
	TunnelInterfaceLastResortCircuit                   types.Bool                                                        `tfsdk:"tunnel_interface_last_resort_circuit"`
	TunnelInterfaceLastResortCircuitVariable           types.String                                                      `tfsdk:"tunnel_interface_last_resort_circuit_variable"`
	TunnelInterfaceNatRefreshInterval                  types.Int64                                                       `tfsdk:"tunnel_interface_nat_refresh_interval"`
	TunnelInterfaceNatRefreshIntervalVariable          types.String                                                      `tfsdk:"tunnel_interface_nat_refresh_interval_variable"`
	TunnelInterfaceHelloInterval                       types.Int64                                                       `tfsdk:"tunnel_interface_hello_interval"`
	TunnelInterfaceHelloIntervalVariable               types.String                                                      `tfsdk:"tunnel_interface_hello_interval_variable"`
	TunnelInterfaceHelloTolerance                      types.Int64                                                       `tfsdk:"tunnel_interface_hello_tolerance"`
	TunnelInterfaceHelloToleranceVariable              types.String                                                      `tfsdk:"tunnel_interface_hello_tolerance_variable"`
	TunnelInterfaceAllowAll                            types.Bool                                                        `tfsdk:"tunnel_interface_allow_all"`
	TunnelInterfaceAllowAllVariable                    types.String                                                      `tfsdk:"tunnel_interface_allow_all_variable"`
	TunnelInterfaceAllowBgp                            types.Bool                                                        `tfsdk:"tunnel_interface_allow_bgp"`
	TunnelInterfaceAllowBgpVariable                    types.String                                                      `tfsdk:"tunnel_interface_allow_bgp_variable"`
	TunnelInterfaceAllowDhcp                           types.Bool                                                        `tfsdk:"tunnel_interface_allow_dhcp"`
	TunnelInterfaceAllowDhcpVariable                   types.String                                                      `tfsdk:"tunnel_interface_allow_dhcp_variable"`
	TunnelInterfaceAllowDns                            types.Bool                                                        `tfsdk:"tunnel_interface_allow_dns"`
	TunnelInterfaceAllowDnsVariable                    types.String                                                      `tfsdk:"tunnel_interface_allow_dns_variable"`
	TunnelInterfaceAllowIcmp                           types.Bool                                                        `tfsdk:"tunnel_interface_allow_icmp"`
	TunnelInterfaceAllowIcmpVariable                   types.String                                                      `tfsdk:"tunnel_interface_allow_icmp_variable"`
	TunnelInterfaceAllowNetconf                        types.Bool                                                        `tfsdk:"tunnel_interface_allow_netconf"`
	TunnelInterfaceAllowNetconfVariable                types.String                                                      `tfsdk:"tunnel_interface_allow_netconf_variable"`
	TunnelInterfaceAllowNtp                            types.Bool                                                        `tfsdk:"tunnel_interface_allow_ntp"`
	TunnelInterfaceAllowNtpVariable                    types.String                                                      `tfsdk:"tunnel_interface_allow_ntp_variable"`
	TunnelInterfaceAllowOspf                           types.Bool                                                        `tfsdk:"tunnel_interface_allow_ospf"`
	TunnelInterfaceAllowOspfVariable                   types.String                                                      `tfsdk:"tunnel_interface_allow_ospf_variable"`
	TunnelInterfaceAllowSsh                            types.Bool                                                        `tfsdk:"tunnel_interface_allow_ssh"`
	TunnelInterfaceAllowSshVariable                    types.String                                                      `tfsdk:"tunnel_interface_allow_ssh_variable"`
	TunnelInterfaceAllowStun                           types.Bool                                                        `tfsdk:"tunnel_interface_allow_stun"`
	TunnelInterfaceAllowStunVariable                   types.String                                                      `tfsdk:"tunnel_interface_allow_stun_variable"`
	TunnelInterfaceAllowHttps                          types.Bool                                                        `tfsdk:"tunnel_interface_allow_https"`
	TunnelInterfaceAllowHttpsVariable                  types.String                                                      `tfsdk:"tunnel_interface_allow_https_variable"`
	TunnelInterfaceAllowSnmp                           types.Bool                                                        `tfsdk:"tunnel_interface_allow_snmp"`
	TunnelInterfaceAllowSnmpVariable                   types.String                                                      `tfsdk:"tunnel_interface_allow_snmp_variable"`
	TunnelInterfaceAllowBfd                            types.Bool                                                        `tfsdk:"tunnel_interface_allow_bfd"`
	TunnelInterfaceAllowBfdVariable                    types.String                                                      `tfsdk:"tunnel_interface_allow_bfd_variable"`
	TunnelInterfaceEncapsulations                      []TransportWANVPNInterfaceT1E1SerialTunnelInterfaceEncapsulations `tfsdk:"tunnel_interface_encapsulations"`
	QosShapingRate                                     types.Int64                                                       `tfsdk:"qos_shaping_rate"`
	QosShapingRateVariable                             types.String                                                      `tfsdk:"qos_shaping_rate_variable"`
	TcpMss                                             types.Int64                                                       `tfsdk:"tcp_mss"`
	TcpMssVariable                                     types.String                                                      `tfsdk:"tcp_mss_variable"`
	Mtu                                                types.Int64                                                       `tfsdk:"mtu"`
	MtuVariable                                        types.String                                                      `tfsdk:"mtu_variable"`
	IpMtu                                              types.Int64                                                       `tfsdk:"ip_mtu"`
	IpMtuVariable                                      types.String                                                      `tfsdk:"ip_mtu_variable"`
	TlocExtension                                      types.String                                                      `tfsdk:"tloc_extension"`
	TlocExtensionVariable                              types.String                                                      `tfsdk:"tloc_extension_variable"`
}

type TransportWANVPNInterfaceT1E1SerialTunnelInterfaceEncapsulations struct {
	Encapsulation      types.String `tfsdk:"encapsulation"`
	Preference         types.Int64  `tfsdk:"preference"`
	PreferenceVariable types.String `tfsdk:"preference_variable"`
	Weight             types.Int64  `tfsdk:"weight"`
	WeightVariable     types.String `tfsdk:"weight_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TransportWANVPNInterfaceT1E1Serial) getModel() string {
	return "transport_wan_vpn_interface_t1_e1_serial"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportWANVPNInterfaceT1E1Serial) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/wan/vpn/%s/interface/serial", url.QueryEscape(data.FeatureProfileId.ValueString()), url.QueryEscape(data.TransportWanVpnFeatureId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportWANVPNInterfaceT1E1Serial) toBody(ctx context.Context) string {
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

	if !data.Ipv4AddressVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"addressV4.address.optionType", "variable")
			body, _ = sjson.Set(body, path+"addressV4.address.value", data.Ipv4AddressVariable.ValueString())
		}
	} else if data.Ipv4Address.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"addressV4.address.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"addressV4.address.optionType", "global")
			body, _ = sjson.Set(body, path+"addressV4.address.value", data.Ipv4Address.ValueString())
		}
	}

	if !data.Ipv4SubnetMaskVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"addressV4.mask.optionType", "variable")
			body, _ = sjson.Set(body, path+"addressV4.mask.value", data.Ipv4SubnetMaskVariable.ValueString())
		}
	} else if !data.Ipv4SubnetMask.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"addressV4.mask.optionType", "global")
			body, _ = sjson.Set(body, path+"addressV4.mask.value", data.Ipv4SubnetMask.ValueString())
		}
	}

	if !data.Ipv6AddressVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"addressV6.optionType", "variable")
			body, _ = sjson.Set(body, path+"addressV6.value", data.Ipv6AddressVariable.ValueString())
		}
	} else if data.Ipv6Address.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"addressV6.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"addressV6.optionType", "global")
			body, _ = sjson.Set(body, path+"addressV6.value", data.Ipv6Address.ValueString())
		}
	}

	if !data.BandwidthVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"bandwidth.optionType", "variable")
			body, _ = sjson.Set(body, path+"bandwidth.value", data.BandwidthVariable.ValueString())
		}
	} else if data.Bandwidth.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"bandwidth.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"bandwidth.optionType", "global")
			body, _ = sjson.Set(body, path+"bandwidth.value", data.Bandwidth.ValueInt64())
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

	if !data.ClockRateVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"clockRate.optionType", "variable")
			body, _ = sjson.Set(body, path+"clockRate.value", data.ClockRateVariable.ValueString())
		}
	} else if data.ClockRate.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"clockRate.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"clockRate.optionType", "global")
			body, _ = sjson.Set(body, path+"clockRate.value", data.ClockRate.ValueString())
		}
	}

	if !data.EncapsulationVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"encapsulationSerial.optionType", "variable")
			body, _ = sjson.Set(body, path+"encapsulationSerial.value", data.EncapsulationVariable.ValueString())
		}
	} else if data.Encapsulation.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"encapsulationSerial.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"encapsulationSerial.optionType", "global")
			body, _ = sjson.Set(body, path+"encapsulationSerial.value", data.Encapsulation.ValueString())
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
		if true {
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.value", data.PerTunnelQosVariable.ValueString())
		}
	} else if data.PerTunnelQos.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQos.value", data.PerTunnelQos.ValueBool())
		}
	}

	if !data.PerTunnelQosAggregatorVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQosAggregator.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQosAggregator.value", data.PerTunnelQosAggregatorVariable.ValueString())
		}
	} else if data.PerTunnelQosAggregator.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQosAggregator.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQosAggregator.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQosAggregator.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.perTunnelQosAggregator.value", data.PerTunnelQosAggregator.ValueBool())
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

	if !data.TunnelInterfaceRestrictVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.restrict.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.restrict.value", data.TunnelInterfaceRestrictVariable.ValueString())
		}
	} else if data.TunnelInterfaceRestrict.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.restrict.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.restrict.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.restrict.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.restrict.value", data.TunnelInterfaceRestrict.ValueBool())
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
			body, _ = sjson.Set(body, path+"tunnel.group.value", data.TunnelInterfaceGroups.ValueInt64())
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
			var values []int64
			data.TunnelInterfaceExcludeControllerGroupList.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"tunnel.excludeControllerGroupList.value", values)
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

	if !data.TunnelInterfaceClearNetworkBroadcastVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.value", data.TunnelInterfaceClearNetworkBroadcastVariable.ValueString())
		}
	} else if data.TunnelInterfaceClearNetworkBroadcast.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.networkBroadcast.value", data.TunnelInterfaceClearNetworkBroadcast.ValueBool())
		}
	}

	if !data.TunnelInterfaceCarrierVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.carrier.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.carrier.value", data.TunnelInterfaceCarrierVariable.ValueString())
		}
	} else if data.TunnelInterfaceCarrier.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.carrier.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.carrier.value", "default")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.carrier.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.carrier.value", data.TunnelInterfaceCarrier.ValueString())
		}
	}

	if !data.TunnelInterfaceBindLoopbackTunnelVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.bind.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.bind.value", data.TunnelInterfaceBindLoopbackTunnelVariable.ValueString())
		}
	} else if data.TunnelInterfaceBindLoopbackTunnel.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.bind.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.bind.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.bind.value", data.TunnelInterfaceBindLoopbackTunnel.ValueString())
		}
	}

	if !data.TunnelInterfaceLastResortCircuitVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.value", data.TunnelInterfaceLastResortCircuitVariable.ValueString())
		}
	} else if data.TunnelInterfaceLastResortCircuit.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.lastResortCircuit.value", data.TunnelInterfaceLastResortCircuit.ValueBool())
		}
	}

	if !data.TunnelInterfaceNatRefreshIntervalVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.value", data.TunnelInterfaceNatRefreshIntervalVariable.ValueString())
		}
	} else if data.TunnelInterfaceNatRefreshInterval.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.value", 5)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.natRefreshInterval.value", data.TunnelInterfaceNatRefreshInterval.ValueInt64())
		}
	}

	if !data.TunnelInterfaceHelloIntervalVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.helloInterval.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.helloInterval.value", data.TunnelInterfaceHelloIntervalVariable.ValueString())
		}
	} else if data.TunnelInterfaceHelloInterval.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.helloInterval.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.helloInterval.value", 1000)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.helloInterval.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.helloInterval.value", data.TunnelInterfaceHelloInterval.ValueInt64())
		}
	}

	if !data.TunnelInterfaceHelloToleranceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.helloTolerance.optionType", "variable")
			body, _ = sjson.Set(body, path+"tunnel.helloTolerance.value", data.TunnelInterfaceHelloToleranceVariable.ValueString())
		}
	} else if data.TunnelInterfaceHelloTolerance.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.helloTolerance.optionType", "default")
			body, _ = sjson.Set(body, path+"tunnel.helloTolerance.value", 12)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tunnel.helloTolerance.optionType", "global")
			body, _ = sjson.Set(body, path+"tunnel.helloTolerance.value", data.TunnelInterfaceHelloTolerance.ValueInt64())
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

	if !data.TunnelInterfaceAllowSshVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.sshd.optionType", "variable")
			body, _ = sjson.Set(body, path+"allowService.sshd.value", data.TunnelInterfaceAllowSshVariable.ValueString())
		}
	} else if !data.TunnelInterfaceAllowSsh.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"allowService.sshd.optionType", "global")
			body, _ = sjson.Set(body, path+"allowService.sshd.value", data.TunnelInterfaceAllowSsh.ValueBool())
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

	if !data.TcpMssVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tcpMssAdjust.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.tcpMssAdjust.value", data.TcpMssVariable.ValueString())
		}
	} else if data.TcpMss.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tcpMssAdjust.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tcpMssAdjust.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.tcpMssAdjust.value", data.TcpMss.ValueInt64())
		}
	}

	if !data.MtuVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.mtu.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.mtu.value", data.MtuVariable.ValueString())
		}
	} else if data.Mtu.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.mtu.optionType", "default")
			body, _ = sjson.Set(body, path+"advanced.mtu.value", 1500)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.mtu.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.mtu.value", data.Mtu.ValueInt64())
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
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportWANVPNInterfaceT1E1Serial) fromBody(ctx context.Context, res gjson.Result) {
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
	data.Ipv4Address = types.StringNull()
	data.Ipv4AddressVariable = types.StringNull()
	if t := res.Get(path + "addressV4.address.optionType"); t.Exists() {
		va := res.Get(path + "addressV4.address.value")
		if t.String() == "variable" {
			data.Ipv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Address = types.StringValue(va.String())
		}
	}
	data.Ipv4SubnetMask = types.StringNull()
	data.Ipv4SubnetMaskVariable = types.StringNull()
	if t := res.Get(path + "addressV4.mask.optionType"); t.Exists() {
		va := res.Get(path + "addressV4.mask.value")
		if t.String() == "variable" {
			data.Ipv4SubnetMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4SubnetMask = types.StringValue(va.String())
		}
	}
	data.Ipv6Address = types.StringNull()
	data.Ipv6AddressVariable = types.StringNull()
	if t := res.Get(path + "addressV6.optionType"); t.Exists() {
		va := res.Get(path + "addressV6.value")
		if t.String() == "variable" {
			data.Ipv6AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6Address = types.StringValue(va.String())
		}
	}
	data.Bandwidth = types.Int64Null()
	data.BandwidthVariable = types.StringNull()
	if t := res.Get(path + "bandwidth.optionType"); t.Exists() {
		va := res.Get(path + "bandwidth.value")
		if t.String() == "variable" {
			data.BandwidthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Bandwidth = types.Int64Value(va.Int())
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
	data.ClockRate = types.StringNull()
	data.ClockRateVariable = types.StringNull()
	if t := res.Get(path + "clockRate.optionType"); t.Exists() {
		va := res.Get(path + "clockRate.value")
		if t.String() == "variable" {
			data.ClockRateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ClockRate = types.StringValue(va.String())
		}
	}
	data.Encapsulation = types.StringNull()
	data.EncapsulationVariable = types.StringNull()
	if t := res.Get(path + "encapsulationSerial.optionType"); t.Exists() {
		va := res.Get(path + "encapsulationSerial.value")
		if t.String() == "variable" {
			data.EncapsulationVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Encapsulation = types.StringValue(va.String())
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
	data.PerTunnelQosAggregator = types.BoolNull()
	data.PerTunnelQosAggregatorVariable = types.StringNull()
	if t := res.Get(path + "tunnel.perTunnelQosAggregator.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.perTunnelQosAggregator.value")
		if t.String() == "variable" {
			data.PerTunnelQosAggregatorVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PerTunnelQosAggregator = types.BoolValue(va.Bool())
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
	data.TunnelInterfaceRestrict = types.BoolNull()
	data.TunnelInterfaceRestrictVariable = types.StringNull()
	if t := res.Get(path + "tunnel.restrict.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.restrict.value")
		if t.String() == "variable" {
			data.TunnelInterfaceRestrictVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceRestrict = types.BoolValue(va.Bool())
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
	if t := res.Get(path + "tunnel.vmanageConnectionPreference.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.vmanageConnectionPreference.value")
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
	data.TunnelInterfaceClearNetworkBroadcast = types.BoolNull()
	data.TunnelInterfaceClearNetworkBroadcastVariable = types.StringNull()
	if t := res.Get(path + "tunnel.networkBroadcast.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.networkBroadcast.value")
		if t.String() == "variable" {
			data.TunnelInterfaceClearNetworkBroadcastVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceClearNetworkBroadcast = types.BoolValue(va.Bool())
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
	data.TunnelInterfaceAllowSsh = types.BoolNull()
	data.TunnelInterfaceAllowSshVariable = types.StringNull()
	if t := res.Get(path + "allowService.sshd.optionType"); t.Exists() {
		va := res.Get(path + "allowService.sshd.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowSshVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowSsh = types.BoolValue(va.Bool())
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
		data.TunnelInterfaceEncapsulations = make([]TransportWANVPNInterfaceT1E1SerialTunnelInterfaceEncapsulations, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportWANVPNInterfaceT1E1SerialTunnelInterfaceEncapsulations{}
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
	data.TcpMss = types.Int64Null()
	data.TcpMssVariable = types.StringNull()
	if t := res.Get(path + "advanced.tcpMssAdjust.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tcpMssAdjust.value")
		if t.String() == "variable" {
			data.TcpMssVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TcpMss = types.Int64Value(va.Int())
		}
	}
	data.Mtu = types.Int64Null()
	data.MtuVariable = types.StringNull()
	if t := res.Get(path + "advanced.mtu.optionType"); t.Exists() {
		va := res.Get(path + "advanced.mtu.value")
		if t.String() == "variable" {
			data.MtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Mtu = types.Int64Value(va.Int())
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TransportWANVPNInterfaceT1E1Serial) updateFromBody(ctx context.Context, res gjson.Result) {
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
	data.Ipv4Address = types.StringNull()
	data.Ipv4AddressVariable = types.StringNull()
	if t := res.Get(path + "addressV4.address.optionType"); t.Exists() {
		va := res.Get(path + "addressV4.address.value")
		if t.String() == "variable" {
			data.Ipv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Address = types.StringValue(va.String())
		}
	}
	data.Ipv4SubnetMask = types.StringNull()
	data.Ipv4SubnetMaskVariable = types.StringNull()
	if t := res.Get(path + "addressV4.mask.optionType"); t.Exists() {
		va := res.Get(path + "addressV4.mask.value")
		if t.String() == "variable" {
			data.Ipv4SubnetMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4SubnetMask = types.StringValue(va.String())
		}
	}
	data.Ipv6Address = types.StringNull()
	data.Ipv6AddressVariable = types.StringNull()
	if t := res.Get(path + "addressV6.optionType"); t.Exists() {
		va := res.Get(path + "addressV6.value")
		if t.String() == "variable" {
			data.Ipv6AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6Address = types.StringValue(va.String())
		}
	}
	data.Bandwidth = types.Int64Null()
	data.BandwidthVariable = types.StringNull()
	if t := res.Get(path + "bandwidth.optionType"); t.Exists() {
		va := res.Get(path + "bandwidth.value")
		if t.String() == "variable" {
			data.BandwidthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Bandwidth = types.Int64Value(va.Int())
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
	data.ClockRate = types.StringNull()
	data.ClockRateVariable = types.StringNull()
	if t := res.Get(path + "clockRate.optionType"); t.Exists() {
		va := res.Get(path + "clockRate.value")
		if t.String() == "variable" {
			data.ClockRateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ClockRate = types.StringValue(va.String())
		}
	}
	data.Encapsulation = types.StringNull()
	data.EncapsulationVariable = types.StringNull()
	if t := res.Get(path + "encapsulationSerial.optionType"); t.Exists() {
		va := res.Get(path + "encapsulationSerial.value")
		if t.String() == "variable" {
			data.EncapsulationVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Encapsulation = types.StringValue(va.String())
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
	data.PerTunnelQosAggregator = types.BoolNull()
	data.PerTunnelQosAggregatorVariable = types.StringNull()
	if t := res.Get(path + "tunnel.perTunnelQosAggregator.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.perTunnelQosAggregator.value")
		if t.String() == "variable" {
			data.PerTunnelQosAggregatorVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PerTunnelQosAggregator = types.BoolValue(va.Bool())
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
	data.TunnelInterfaceRestrict = types.BoolNull()
	data.TunnelInterfaceRestrictVariable = types.StringNull()
	if t := res.Get(path + "tunnel.restrict.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.restrict.value")
		if t.String() == "variable" {
			data.TunnelInterfaceRestrictVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceRestrict = types.BoolValue(va.Bool())
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
	if t := res.Get(path + "tunnel.vmanageConnectionPreference.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.vmanageConnectionPreference.value")
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
	data.TunnelInterfaceClearNetworkBroadcast = types.BoolNull()
	data.TunnelInterfaceClearNetworkBroadcastVariable = types.StringNull()
	if t := res.Get(path + "tunnel.networkBroadcast.optionType"); t.Exists() {
		va := res.Get(path + "tunnel.networkBroadcast.value")
		if t.String() == "variable" {
			data.TunnelInterfaceClearNetworkBroadcastVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceClearNetworkBroadcast = types.BoolValue(va.Bool())
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
	data.TunnelInterfaceAllowSsh = types.BoolNull()
	data.TunnelInterfaceAllowSshVariable = types.StringNull()
	if t := res.Get(path + "allowService.sshd.optionType"); t.Exists() {
		va := res.Get(path + "allowService.sshd.value")
		if t.String() == "variable" {
			data.TunnelInterfaceAllowSshVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelInterfaceAllowSsh = types.BoolValue(va.Bool())
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
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
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
	data.TcpMss = types.Int64Null()
	data.TcpMssVariable = types.StringNull()
	if t := res.Get(path + "advanced.tcpMssAdjust.optionType"); t.Exists() {
		va := res.Get(path + "advanced.tcpMssAdjust.value")
		if t.String() == "variable" {
			data.TcpMssVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TcpMss = types.Int64Value(va.Int())
		}
	}
	data.Mtu = types.Int64Null()
	data.MtuVariable = types.StringNull()
	if t := res.Get(path + "advanced.mtu.optionType"); t.Exists() {
		va := res.Get(path + "advanced.mtu.value")
		if t.String() == "variable" {
			data.MtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Mtu = types.Int64Value(va.Int())
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
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *TransportWANVPNInterfaceT1E1Serial) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.TransportWanVpnFeatureId.IsNull() {
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
	if !data.Ipv6Address.IsNull() {
		return false
	}
	if !data.Ipv6AddressVariable.IsNull() {
		return false
	}
	if !data.Bandwidth.IsNull() {
		return false
	}
	if !data.BandwidthVariable.IsNull() {
		return false
	}
	if !data.BandwidthDownstream.IsNull() {
		return false
	}
	if !data.BandwidthDownstreamVariable.IsNull() {
		return false
	}
	if !data.ClockRate.IsNull() {
		return false
	}
	if !data.ClockRateVariable.IsNull() {
		return false
	}
	if !data.Encapsulation.IsNull() {
		return false
	}
	if !data.EncapsulationVariable.IsNull() {
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
	if !data.PerTunnelQosAggregator.IsNull() {
		return false
	}
	if !data.PerTunnelQosAggregatorVariable.IsNull() {
		return false
	}
	if !data.TunnelQosMode.IsNull() {
		return false
	}
	if !data.TunnelQosModeVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceColor.IsNull() {
		return false
	}
	if !data.TunnelInterfaceColorVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceRestrict.IsNull() {
		return false
	}
	if !data.TunnelInterfaceRestrictVariable.IsNull() {
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
	if !data.TunnelInterfaceClearNetworkBroadcast.IsNull() {
		return false
	}
	if !data.TunnelInterfaceClearNetworkBroadcastVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceCarrier.IsNull() {
		return false
	}
	if !data.TunnelInterfaceCarrierVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceBindLoopbackTunnel.IsNull() {
		return false
	}
	if !data.TunnelInterfaceBindLoopbackTunnelVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceLastResortCircuit.IsNull() {
		return false
	}
	if !data.TunnelInterfaceLastResortCircuitVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceNatRefreshInterval.IsNull() {
		return false
	}
	if !data.TunnelInterfaceNatRefreshIntervalVariable.IsNull() {
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
	if !data.TunnelInterfaceAllowDns.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowDnsVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowIcmp.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowIcmpVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowNetconf.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowNetconfVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowNtp.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowNtpVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowOspf.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowOspfVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowSsh.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowSshVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowStun.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowStunVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowHttps.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowHttpsVariable.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowSnmp.IsNull() {
		return false
	}
	if !data.TunnelInterfaceAllowSnmpVariable.IsNull() {
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
	if !data.QosShapingRate.IsNull() {
		return false
	}
	if !data.QosShapingRateVariable.IsNull() {
		return false
	}
	if !data.TcpMss.IsNull() {
		return false
	}
	if !data.TcpMssVariable.IsNull() {
		return false
	}
	if !data.Mtu.IsNull() {
		return false
	}
	if !data.MtuVariable.IsNull() {
		return false
	}
	if !data.IpMtu.IsNull() {
		return false
	}
	if !data.IpMtuVariable.IsNull() {
		return false
	}
	if !data.TlocExtension.IsNull() {
		return false
	}
	if !data.TlocExtensionVariable.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
