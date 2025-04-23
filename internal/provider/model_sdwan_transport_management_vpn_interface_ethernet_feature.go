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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type TransportManagementVPNInterfaceEthernet struct {
	Id                              types.String                                                    `tfsdk:"id"`
	Version                         types.Int64                                                     `tfsdk:"version"`
	Name                            types.String                                                    `tfsdk:"name"`
	Description                     types.String                                                    `tfsdk:"description"`
	FeatureProfileId                types.String                                                    `tfsdk:"feature_profile_id"`
	TransportManagementVpnFeatureId types.String                                                    `tfsdk:"transport_management_vpn_feature_id"`
	Shutdown                        types.Bool                                                      `tfsdk:"shutdown"`
	ShutdownVariable                types.String                                                    `tfsdk:"shutdown_variable"`
	InterfaceName                   types.String                                                    `tfsdk:"interface_name"`
	InterfaceNameVariable           types.String                                                    `tfsdk:"interface_name_variable"`
	InterfaceDescription            types.String                                                    `tfsdk:"interface_description"`
	InterfaceDescriptionVariable    types.String                                                    `tfsdk:"interface_description_variable"`
	Ipv4ConfigurationType           types.String                                                    `tfsdk:"ipv4_configuration_type"`
	Ipv4DhcpDistance                types.Int64                                                     `tfsdk:"ipv4_dhcp_distance"`
	Ipv4DhcpDistanceVariable        types.String                                                    `tfsdk:"ipv4_dhcp_distance_variable"`
	Ipv4Address                     types.String                                                    `tfsdk:"ipv4_address"`
	Ipv4AddressVariable             types.String                                                    `tfsdk:"ipv4_address_variable"`
	Ipv4SubnetMask                  types.String                                                    `tfsdk:"ipv4_subnet_mask"`
	Ipv4SubnetMaskVariable          types.String                                                    `tfsdk:"ipv4_subnet_mask_variable"`
	Ipv4SecondaryAddresses          []TransportManagementVPNInterfaceEthernetIpv4SecondaryAddresses `tfsdk:"ipv4_secondary_addresses"`
	Ipv4DhcpHelper                  types.Set                                                       `tfsdk:"ipv4_dhcp_helper"`
	Ipv4DhcpHelperVariable          types.String                                                    `tfsdk:"ipv4_dhcp_helper_variable"`
	Ipv4IperfServer                 types.String                                                    `tfsdk:"ipv4_iperf_server"`
	Ipv4IperfServerVariable         types.String                                                    `tfsdk:"ipv4_iperf_server_variable"`
	Ipv4AutoDetectBandwidth         types.Bool                                                      `tfsdk:"ipv4_auto_detect_bandwidth"`
	Ipv4AutoDetectBandwidthVariable types.String                                                    `tfsdk:"ipv4_auto_detect_bandwidth_variable"`
	Ipv6ConfigurationType           types.String                                                    `tfsdk:"ipv6_configuration_type"`
	EnableDhcpv6                    types.Bool                                                      `tfsdk:"enable_dhcpv6"`
	Ipv6Address                     types.String                                                    `tfsdk:"ipv6_address"`
	Ipv6AddressVariable             types.String                                                    `tfsdk:"ipv6_address_variable"`
	ArpEntries                      []TransportManagementVPNInterfaceEthernetArpEntries             `tfsdk:"arp_entries"`
	Duplex                          types.String                                                    `tfsdk:"duplex"`
	DuplexVariable                  types.String                                                    `tfsdk:"duplex_variable"`
	MacAddress                      types.String                                                    `tfsdk:"mac_address"`
	MacAddressVariable              types.String                                                    `tfsdk:"mac_address_variable"`
	IpMtu                           types.Int64                                                     `tfsdk:"ip_mtu"`
	IpMtuVariable                   types.String                                                    `tfsdk:"ip_mtu_variable"`
	InterfaceMtu                    types.Int64                                                     `tfsdk:"interface_mtu"`
	InterfaceMtuVariable            types.String                                                    `tfsdk:"interface_mtu_variable"`
	TcpMss                          types.Int64                                                     `tfsdk:"tcp_mss"`
	TcpMssVariable                  types.String                                                    `tfsdk:"tcp_mss_variable"`
	Speed                           types.String                                                    `tfsdk:"speed"`
	SpeedVariable                   types.String                                                    `tfsdk:"speed_variable"`
	ArpTimeout                      types.Int64                                                     `tfsdk:"arp_timeout"`
	ArpTimeoutVariable              types.String                                                    `tfsdk:"arp_timeout_variable"`
	Autonegotiate                   types.Bool                                                      `tfsdk:"autonegotiate"`
	AutonegotiateVariable           types.String                                                    `tfsdk:"autonegotiate_variable"`
	MediaType                       types.String                                                    `tfsdk:"media_type"`
	MediaTypeVariable               types.String                                                    `tfsdk:"media_type_variable"`
	LoadInterval                    types.Int64                                                     `tfsdk:"load_interval"`
	LoadIntervalVariable            types.String                                                    `tfsdk:"load_interval_variable"`
	IcmpRedirectDisable             types.Bool                                                      `tfsdk:"icmp_redirect_disable"`
	IcmpRedirectDisableVariable     types.String                                                    `tfsdk:"icmp_redirect_disable_variable"`
	IpDirectedBroadcast             types.Bool                                                      `tfsdk:"ip_directed_broadcast"`
	IpDirectedBroadcastVariable     types.String                                                    `tfsdk:"ip_directed_broadcast_variable"`
}

type TransportManagementVPNInterfaceEthernetIpv4SecondaryAddresses struct {
	Address            types.String `tfsdk:"address"`
	AddressVariable    types.String `tfsdk:"address_variable"`
	SubnetMask         types.String `tfsdk:"subnet_mask"`
	SubnetMaskVariable types.String `tfsdk:"subnet_mask_variable"`
}

type TransportManagementVPNInterfaceEthernetArpEntries struct {
	IpAddress          types.String `tfsdk:"ip_address"`
	IpAddressVariable  types.String `tfsdk:"ip_address_variable"`
	MacAddress         types.String `tfsdk:"mac_address"`
	MacAddressVariable types.String `tfsdk:"mac_address_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TransportManagementVPNInterfaceEthernet) getModel() string {
	return "transport_management_vpn_interface_ethernet"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportManagementVPNInterfaceEthernet) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/management/vpn/%s/interface/ethernet", url.QueryEscape(data.FeatureProfileId.ValueString()), url.QueryEscape(data.TransportManagementVpnFeatureId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportManagementVPNInterfaceEthernet) toBody(ctx context.Context) string {
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

	if !data.Ipv4IperfServerVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"iperfServer.optionType", "variable")
			body, _ = sjson.Set(body, path+"iperfServer.value", data.Ipv4IperfServerVariable.ValueString())
		}
	} else if data.Ipv4IperfServer.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"iperfServer.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"iperfServer.optionType", "global")
			body, _ = sjson.Set(body, path+"iperfServer.value", data.Ipv4IperfServer.ValueString())
		}
	}

	if !data.Ipv4AutoDetectBandwidthVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"autoDetectBandwidth.optionType", "variable")
			body, _ = sjson.Set(body, path+"autoDetectBandwidth.value", data.Ipv4AutoDetectBandwidthVariable.ValueString())
		}
	} else if data.Ipv4AutoDetectBandwidth.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"autoDetectBandwidth.optionType", "default")
			body, _ = sjson.Set(body, path+"autoDetectBandwidth.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"autoDetectBandwidth.optionType", "global")
			body, _ = sjson.Set(body, path+"autoDetectBandwidth.value", data.Ipv4AutoDetectBandwidth.ValueBool())
		}
	}
	if !data.EnableDhcpv6.IsNull() {
		if true && data.Ipv6ConfigurationType.ValueString() == "dynamic" {
			body, _ = sjson.Set(body, path+"intfIpV6Address.dynamic.dhcpClient.optionType", "global")
			body, _ = sjson.Set(body, path+"intfIpV6Address.dynamic.dhcpClient.value", data.EnableDhcpv6.ValueBool())
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
	if true {
		body, _ = sjson.Set(body, path+"arp", []interface{}{})
		for _, item := range data.ArpEntries {
			itemBody := ""

			if !item.IpAddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipAddress.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ipAddress.value", item.IpAddressVariable.ValueString())
				}
			} else if !item.IpAddress.IsNull() {
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
			} else if !item.MacAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "macAddress.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "macAddress.value", item.MacAddress.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"arp.-1", itemBody)
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
		if true {
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

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportManagementVPNInterfaceEthernet) fromBody(ctx context.Context, res gjson.Result) {
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
		data.Ipv4SecondaryAddresses = make([]TransportManagementVPNInterfaceEthernetIpv4SecondaryAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportManagementVPNInterfaceEthernetIpv4SecondaryAddresses{}
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
	data.Ipv4IperfServer = types.StringNull()
	data.Ipv4IperfServerVariable = types.StringNull()
	if t := res.Get(path + "iperfServer.optionType"); t.Exists() {
		va := res.Get(path + "iperfServer.value")
		if t.String() == "variable" {
			data.Ipv4IperfServerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4IperfServer = types.StringValue(va.String())
		}
	}
	data.Ipv4AutoDetectBandwidth = types.BoolNull()
	data.Ipv4AutoDetectBandwidthVariable = types.StringNull()
	if t := res.Get(path + "autoDetectBandwidth.optionType"); t.Exists() {
		va := res.Get(path + "autoDetectBandwidth.value")
		if t.String() == "variable" {
			data.Ipv4AutoDetectBandwidthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4AutoDetectBandwidth = types.BoolValue(va.Bool())
		}
	}
	data.EnableDhcpv6 = types.BoolNull()

	if t := res.Get(path + "intfIpV6Address.dynamic.dhcpClient.optionType"); t.Exists() {
		va := res.Get(path + "intfIpV6Address.dynamic.dhcpClient.value")
		if t.String() == "global" {
			data.EnableDhcpv6 = types.BoolValue(va.Bool())
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
	if value := res.Get(path + "arp"); value.Exists() {
		data.ArpEntries = make([]TransportManagementVPNInterfaceEthernetArpEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportManagementVPNInterfaceEthernetArpEntries{}
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
			data.ArpEntries = append(data.ArpEntries, item)
			return true
		})
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
func (data *TransportManagementVPNInterfaceEthernet) updateFromBody(ctx context.Context, res gjson.Result) {
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
	data.Ipv4IperfServer = types.StringNull()
	data.Ipv4IperfServerVariable = types.StringNull()
	if t := res.Get(path + "iperfServer.optionType"); t.Exists() {
		va := res.Get(path + "iperfServer.value")
		if t.String() == "variable" {
			data.Ipv4IperfServerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4IperfServer = types.StringValue(va.String())
		}
	}
	data.Ipv4AutoDetectBandwidth = types.BoolNull()
	data.Ipv4AutoDetectBandwidthVariable = types.StringNull()
	if t := res.Get(path + "autoDetectBandwidth.optionType"); t.Exists() {
		va := res.Get(path + "autoDetectBandwidth.value")
		if t.String() == "variable" {
			data.Ipv4AutoDetectBandwidthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4AutoDetectBandwidth = types.BoolValue(va.Bool())
		}
	}
	data.EnableDhcpv6 = types.BoolNull()

	if t := res.Get(path + "intfIpV6Address.dynamic.dhcpClient.optionType"); t.Exists() {
		va := res.Get(path + "intfIpV6Address.dynamic.dhcpClient.value")
		if t.String() == "global" {
			data.EnableDhcpv6 = types.BoolValue(va.Bool())
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
	for i := range data.ArpEntries {
		keys := [...]string{"ipAddress", "macAddress"}
		keyValues := [...]string{data.ArpEntries[i].IpAddress.ValueString(), data.ArpEntries[i].MacAddress.ValueString()}
		keyValuesVariables := [...]string{data.ArpEntries[i].IpAddressVariable.ValueString(), data.ArpEntries[i].MacAddressVariable.ValueString()}

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
		data.ArpEntries[i].IpAddress = types.StringNull()
		data.ArpEntries[i].IpAddressVariable = types.StringNull()
		if t := r.Get("ipAddress.optionType"); t.Exists() {
			va := r.Get("ipAddress.value")
			if t.String() == "variable" {
				data.ArpEntries[i].IpAddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.ArpEntries[i].IpAddress = types.StringValue(va.String())
			}
		}
		data.ArpEntries[i].MacAddress = types.StringNull()
		data.ArpEntries[i].MacAddressVariable = types.StringNull()
		if t := r.Get("macAddress.optionType"); t.Exists() {
			va := r.Get("macAddress.value")
			if t.String() == "variable" {
				data.ArpEntries[i].MacAddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.ArpEntries[i].MacAddress = types.StringValue(va.String())
			}
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
