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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

var MinServiceLANVPNInterfaceEthernetUpdateVersion = version.Must(version.NewVersion("20.14.0"))

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ServiceLANVPNInterfaceEthernet struct {
	Id                                       types.String                                               `tfsdk:"id"`
	Version                                  types.Int64                                                `tfsdk:"version"`
	Name                                     types.String                                               `tfsdk:"name"`
	Description                              types.String                                               `tfsdk:"description"`
	FeatureProfileId                         types.String                                               `tfsdk:"feature_profile_id"`
	ServiceLanVpnFeatureId                   types.String                                               `tfsdk:"service_lan_vpn_feature_id"`
	Shutdown                                 types.Bool                                                 `tfsdk:"shutdown"`
	ShutdownVariable                         types.String                                               `tfsdk:"shutdown_variable"`
	InterfaceName                            types.String                                               `tfsdk:"interface_name"`
	InterfaceNameVariable                    types.String                                               `tfsdk:"interface_name_variable"`
	InterfaceDescription                     types.String                                               `tfsdk:"interface_description"`
	InterfaceDescriptionVariable             types.String                                               `tfsdk:"interface_description_variable"`
	Ipv4ConfigurationType                    types.String                                               `tfsdk:"ipv4_configuration_type"`
	Ipv4DhcpDistance                         types.Int64                                                `tfsdk:"ipv4_dhcp_distance"`
	Ipv4DhcpDistanceVariable                 types.String                                               `tfsdk:"ipv4_dhcp_distance_variable"`
	Ipv4Address                              types.String                                               `tfsdk:"ipv4_address"`
	Ipv4AddressVariable                      types.String                                               `tfsdk:"ipv4_address_variable"`
	Ipv4SubnetMask                           types.String                                               `tfsdk:"ipv4_subnet_mask"`
	Ipv4SubnetMaskVariable                   types.String                                               `tfsdk:"ipv4_subnet_mask_variable"`
	Ipv4SecondaryAddresses                   []ServiceLANVPNInterfaceEthernetIpv4SecondaryAddresses     `tfsdk:"ipv4_secondary_addresses"`
	Ipv4DhcpHelper                           types.Set                                                  `tfsdk:"ipv4_dhcp_helper"`
	Ipv4DhcpHelperVariable                   types.String                                               `tfsdk:"ipv4_dhcp_helper_variable"`
	Ipv6ConfigurationType                    types.String                                               `tfsdk:"ipv6_configuration_type"`
	EnableDhcpv6                             types.Bool                                                 `tfsdk:"enable_dhcpv6"`
	Ipv6DhcpSecondaryAddresses               []ServiceLANVPNInterfaceEthernetIpv6DhcpSecondaryAddresses `tfsdk:"ipv6_dhcp_secondary_addresses"`
	Ipv6Address                              types.String                                               `tfsdk:"ipv6_address"`
	Ipv6AddressVariable                      types.String                                               `tfsdk:"ipv6_address_variable"`
	Ipv6SecondaryAddresses                   []ServiceLANVPNInterfaceEthernetIpv6SecondaryAddresses     `tfsdk:"ipv6_secondary_addresses"`
	Ipv6DhcpHelpers                          []ServiceLANVPNInterfaceEthernetIpv6DhcpHelpers            `tfsdk:"ipv6_dhcp_helpers"`
	Ipv4Nat                                  types.Bool                                                 `tfsdk:"ipv4_nat"`
	Ipv4NatType                              types.String                                               `tfsdk:"ipv4_nat_type"`
	Ipv4NatTypeVariable                      types.String                                               `tfsdk:"ipv4_nat_type_variable"`
	Ipv4NatRangeStart                        types.String                                               `tfsdk:"ipv4_nat_range_start"`
	Ipv4NatRangeStartVariable                types.String                                               `tfsdk:"ipv4_nat_range_start_variable"`
	Ipv4NatRangeEnd                          types.String                                               `tfsdk:"ipv4_nat_range_end"`
	Ipv4NatRangeEndVariable                  types.String                                               `tfsdk:"ipv4_nat_range_end_variable"`
	Ipv4NatPrefixLength                      types.Int64                                                `tfsdk:"ipv4_nat_prefix_length"`
	Ipv4NatPrefixLengthVariable              types.String                                               `tfsdk:"ipv4_nat_prefix_length_variable"`
	Ipv4NatOverload                          types.Bool                                                 `tfsdk:"ipv4_nat_overload"`
	Ipv4NatOverloadVariable                  types.String                                               `tfsdk:"ipv4_nat_overload_variable"`
	Ipv4NatLoopback                          types.String                                               `tfsdk:"ipv4_nat_loopback"`
	Ipv4NatLoopbackVariable                  types.String                                               `tfsdk:"ipv4_nat_loopback_variable"`
	Ipv4NatUdpTimeout                        types.Int64                                                `tfsdk:"ipv4_nat_udp_timeout"`
	Ipv4NatUdpTimeoutVariable                types.String                                               `tfsdk:"ipv4_nat_udp_timeout_variable"`
	Ipv4NatTcpTimeout                        types.Int64                                                `tfsdk:"ipv4_nat_tcp_timeout"`
	Ipv4NatTcpTimeoutVariable                types.String                                               `tfsdk:"ipv4_nat_tcp_timeout_variable"`
	StaticNats                               []ServiceLANVPNInterfaceEthernetStaticNats                 `tfsdk:"static_nats"`
	Ipv6Nat                                  types.Bool                                                 `tfsdk:"ipv6_nat"`
	Nat64                                    types.Bool                                                 `tfsdk:"nat64"`
	AclShapingRate                           types.Int64                                                `tfsdk:"acl_shaping_rate"`
	AclShapingRateVariable                   types.String                                               `tfsdk:"acl_shaping_rate_variable"`
	AclIpv4EgressPolicyId                    types.String                                               `tfsdk:"acl_ipv4_egress_policy_id"`
	AclIpv4IngressPolicyId                   types.String                                               `tfsdk:"acl_ipv4_ingress_policy_id"`
	AclIpv6EgressPolicyId                    types.String                                               `tfsdk:"acl_ipv6_egress_policy_id"`
	AclIpv6IngressPolicyId                   types.String                                               `tfsdk:"acl_ipv6_ingress_policy_id"`
	Ipv6Vrrps                                []ServiceLANVPNInterfaceEthernetIpv6Vrrps                  `tfsdk:"ipv6_vrrps"`
	Ipv4Vrrps                                []ServiceLANVPNInterfaceEthernetIpv4Vrrps                  `tfsdk:"ipv4_vrrps"`
	Arps                                     []ServiceLANVPNInterfaceEthernetArps                       `tfsdk:"arps"`
	TrustsecEnableSgtPropogation             types.Bool                                                 `tfsdk:"trustsec_enable_sgt_propogation"`
	TrustsecPropogate                        types.Bool                                                 `tfsdk:"trustsec_propogate"`
	TrustsecSecurityGroupTag                 types.Int64                                                `tfsdk:"trustsec_security_group_tag"`
	TrustsecSecurityGroupTagVariable         types.String                                               `tfsdk:"trustsec_security_group_tag_variable"`
	TrustsecEnableEnforcedPropogation        types.Bool                                                 `tfsdk:"trustsec_enable_enforced_propogation"`
	TrustsecEnforcedSecurityGroupTag         types.Int64                                                `tfsdk:"trustsec_enforced_security_group_tag"`
	TrustsecEnforcedSecurityGroupTagVariable types.String                                               `tfsdk:"trustsec_enforced_security_group_tag_variable"`
	Duplex                                   types.String                                               `tfsdk:"duplex"`
	DuplexVariable                           types.String                                               `tfsdk:"duplex_variable"`
	MacAddress                               types.String                                               `tfsdk:"mac_address"`
	MacAddressVariable                       types.String                                               `tfsdk:"mac_address_variable"`
	IpMtu                                    types.Int64                                                `tfsdk:"ip_mtu"`
	IpMtuVariable                            types.String                                               `tfsdk:"ip_mtu_variable"`
	InterfaceMtu                             types.Int64                                                `tfsdk:"interface_mtu"`
	InterfaceMtuVariable                     types.String                                               `tfsdk:"interface_mtu_variable"`
	TcpMss                                   types.Int64                                                `tfsdk:"tcp_mss"`
	TcpMssVariable                           types.String                                               `tfsdk:"tcp_mss_variable"`
	Speed                                    types.String                                               `tfsdk:"speed"`
	SpeedVariable                            types.String                                               `tfsdk:"speed_variable"`
	ArpTimeout                               types.Int64                                                `tfsdk:"arp_timeout"`
	ArpTimeoutVariable                       types.String                                               `tfsdk:"arp_timeout_variable"`
	Autonegotiate                            types.Bool                                                 `tfsdk:"autonegotiate"`
	AutonegotiateVariable                    types.String                                               `tfsdk:"autonegotiate_variable"`
	MediaType                                types.String                                               `tfsdk:"media_type"`
	MediaTypeVariable                        types.String                                               `tfsdk:"media_type_variable"`
	LoadInterval                             types.Int64                                                `tfsdk:"load_interval"`
	LoadIntervalVariable                     types.String                                               `tfsdk:"load_interval_variable"`
	Tracker                                  types.String                                               `tfsdk:"tracker"`
	TrackerVariable                          types.String                                               `tfsdk:"tracker_variable"`
	IcmpRedirectDisable                      types.Bool                                                 `tfsdk:"icmp_redirect_disable"`
	IcmpRedirectDisableVariable              types.String                                               `tfsdk:"icmp_redirect_disable_variable"`
	Xconnect                                 types.String                                               `tfsdk:"xconnect"`
	XconnectVariable                         types.String                                               `tfsdk:"xconnect_variable"`
	IpDirectedBroadcast                      types.Bool                                                 `tfsdk:"ip_directed_broadcast"`
	IpDirectedBroadcastVariable              types.String                                               `tfsdk:"ip_directed_broadcast_variable"`
}

type ServiceLANVPNInterfaceEthernetIpv4SecondaryAddresses struct {
	Address            types.String `tfsdk:"address"`
	AddressVariable    types.String `tfsdk:"address_variable"`
	SubnetMask         types.String `tfsdk:"subnet_mask"`
	SubnetMaskVariable types.String `tfsdk:"subnet_mask_variable"`
}

type ServiceLANVPNInterfaceEthernetIpv6DhcpSecondaryAddresses struct {
	Address         types.String `tfsdk:"address"`
	AddressVariable types.String `tfsdk:"address_variable"`
}

type ServiceLANVPNInterfaceEthernetIpv6SecondaryAddresses struct {
	Address         types.String `tfsdk:"address"`
	AddressVariable types.String `tfsdk:"address_variable"`
}

type ServiceLANVPNInterfaceEthernetIpv6DhcpHelpers struct {
	Address                 types.String `tfsdk:"address"`
	AddressVariable         types.String `tfsdk:"address_variable"`
	Dhcpv6HelperVpn         types.Int64  `tfsdk:"dhcpv6_helper_vpn"`
	Dhcpv6HelperVpnVariable types.String `tfsdk:"dhcpv6_helper_vpn_variable"`
}

type ServiceLANVPNInterfaceEthernetStaticNats struct {
	SourceIp            types.String `tfsdk:"source_ip"`
	SourceIpVariable    types.String `tfsdk:"source_ip_variable"`
	TranslateIp         types.String `tfsdk:"translate_ip"`
	TranslateIpVariable types.String `tfsdk:"translate_ip_variable"`
	Direction           types.String `tfsdk:"direction"`
	SourceVpn           types.Int64  `tfsdk:"source_vpn"`
	SourceVpnVariable   types.String `tfsdk:"source_vpn_variable"`
}

type ServiceLANVPNInterfaceEthernetIpv6Vrrps struct {
	GroupId          types.Int64                                            `tfsdk:"group_id"`
	GroupIdVariable  types.String                                           `tfsdk:"group_id_variable"`
	Priority         types.Int64                                            `tfsdk:"priority"`
	PriorityVariable types.String                                           `tfsdk:"priority_variable"`
	Timer            types.Int64                                            `tfsdk:"timer"`
	TimerVariable    types.String                                           `tfsdk:"timer_variable"`
	TrackOmp         types.Bool                                             `tfsdk:"track_omp"`
	Ipv6Addresses    []ServiceLANVPNInterfaceEthernetIpv6VrrpsIpv6Addresses `tfsdk:"ipv6_addresses"`
}

type ServiceLANVPNInterfaceEthernetIpv4Vrrps struct {
	GroupId             types.Int64                                                 `tfsdk:"group_id"`
	GroupIdVariable     types.String                                                `tfsdk:"group_id_variable"`
	Priority            types.Int64                                                 `tfsdk:"priority"`
	PriorityVariable    types.String                                                `tfsdk:"priority_variable"`
	Timer               types.Int64                                                 `tfsdk:"timer"`
	TimerVariable       types.String                                                `tfsdk:"timer_variable"`
	TrackOmp            types.Bool                                                  `tfsdk:"track_omp"`
	Address             types.String                                                `tfsdk:"address"`
	AddressVariable     types.String                                                `tfsdk:"address_variable"`
	SecondaryAddresses  []ServiceLANVPNInterfaceEthernetIpv4VrrpsSecondaryAddresses `tfsdk:"secondary_addresses"`
	TlocPrefixChange    types.Bool                                                  `tfsdk:"tloc_prefix_change"`
	TlocPrefChangeValue types.Int64                                                 `tfsdk:"tloc_pref_change_value"`
	TrackingObjects     []ServiceLANVPNInterfaceEthernetIpv4VrrpsTrackingObjects    `tfsdk:"tracking_objects"`
}

type ServiceLANVPNInterfaceEthernetArps struct {
	IpAddress          types.String `tfsdk:"ip_address"`
	IpAddressVariable  types.String `tfsdk:"ip_address_variable"`
	MacAddress         types.String `tfsdk:"mac_address"`
	MacAddressVariable types.String `tfsdk:"mac_address_variable"`
}

type ServiceLANVPNInterfaceEthernetIpv6VrrpsIpv6Addresses struct {
	LinkLocalAddress         types.String `tfsdk:"link_local_address"`
	LinkLocalAddressVariable types.String `tfsdk:"link_local_address_variable"`
	GlobalAddress            types.String `tfsdk:"global_address"`
	GlobalAddressVariable    types.String `tfsdk:"global_address_variable"`
}

type ServiceLANVPNInterfaceEthernetIpv4VrrpsSecondaryAddresses struct {
	Address            types.String `tfsdk:"address"`
	AddressVariable    types.String `tfsdk:"address_variable"`
	SubnetMask         types.String `tfsdk:"subnet_mask"`
	SubnetMaskVariable types.String `tfsdk:"subnet_mask_variable"`
}
type ServiceLANVPNInterfaceEthernetIpv4VrrpsTrackingObjects struct {
	TrackerId              types.String `tfsdk:"tracker_id"`
	TrackerAction          types.String `tfsdk:"tracker_action"`
	TrackerActionVariable  types.String `tfsdk:"tracker_action_variable"`
	DecrementValue         types.Int64  `tfsdk:"decrement_value"`
	DecrementValueVariable types.String `tfsdk:"decrement_value_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceLANVPNInterfaceEthernet) getModel() string {
	return "service_lan_vpn_interface_ethernet"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceLANVPNInterfaceEthernet) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/lan/vpn/%s/interface/ethernet", url.QueryEscape(data.FeatureProfileId.ValueString()), url.QueryEscape(data.ServiceLanVpnFeatureId.ValueString()))
}

// End of section. //template:end getPath

func (data ServiceLANVPNInterfaceEthernet) toBody(ctx context.Context, currentVersion *version.Version) string {
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
	} else if !data.Ipv4DhcpDistance.IsNull() {
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

		for _, item := range data.Ipv6DhcpSecondaryAddresses {
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
	if true && data.Ipv6ConfigurationType.ValueString() == "static" {

		for _, item := range data.Ipv6DhcpHelpers {
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

			if !item.Dhcpv6HelperVpnVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Dhcpv6HelperVpnVariable.ValueString())
				}
			} else if !item.Dhcpv6HelperVpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Dhcpv6HelperVpn.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"intfIpV6Address.static.dhcpHelperV6.-1", itemBody)
		}
	}
	if data.Ipv4Nat.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"nat.optionType", "default")
			body, _ = sjson.Set(body, path+"nat.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"nat.optionType", "global")
			body, _ = sjson.Set(body, path+"nat.value", data.Ipv4Nat.ValueBool())
		}
	}

	if !data.Ipv4NatTypeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natType.optionType", "variable")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natType.value", data.Ipv4NatTypeVariable.ValueString())
		}
	} else if !data.Ipv4NatType.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natType.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natType.value", data.Ipv4NatType.ValueString())
		}
	}

	if !data.Ipv4NatRangeStartVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeStart.optionType", "variable")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeStart.value", data.Ipv4NatRangeStartVariable.ValueString())
		}
	} else if !data.Ipv4NatRangeStart.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeStart.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeStart.value", data.Ipv4NatRangeStart.ValueString())
		}
	}

	if !data.Ipv4NatRangeEndVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeEnd.optionType", "variable")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeEnd.value", data.Ipv4NatRangeEndVariable.ValueString())
		}
	} else if !data.Ipv4NatRangeEnd.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeEnd.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.rangeEnd.value", data.Ipv4NatRangeEnd.ValueString())
		}
	}

	if !data.Ipv4NatPrefixLengthVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.prefixLength.optionType", "variable")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.prefixLength.value", data.Ipv4NatPrefixLengthVariable.ValueString())
		}
	} else if !data.Ipv4NatPrefixLength.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.prefixLength.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.prefixLength.value", data.Ipv4NatPrefixLength.ValueInt64())
		}
	}

	if !data.Ipv4NatOverloadVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.overload.optionType", "variable")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.overload.value", data.Ipv4NatOverloadVariable.ValueString())
		}
	} else if !data.Ipv4NatOverload.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.overload.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.natPool.overload.value", data.Ipv4NatOverload.ValueBool())
		}
	}

	natLookBackRef := "natLookback"
	if !currentVersion.LessThan(MinServiceLANVPNInterfaceEthernetUpdateVersion) {
		natLookBackRef = "natLoopback"
	}
	if !data.Ipv4NatLoopbackVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4."+natLookBackRef+".optionType", "variable")
			body, _ = sjson.Set(body, path+"natAttributesIpv4."+natLookBackRef+".value", data.Ipv4NatLoopbackVariable.ValueString())
		}
	} else if !data.Ipv4NatLoopback.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4."+natLookBackRef+".optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv4."+natLookBackRef+".value", data.Ipv4NatLoopback.ValueString())
		}
	}

	if !data.Ipv4NatUdpTimeoutVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.optionType", "variable")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.value", data.Ipv4NatUdpTimeoutVariable.ValueString())
		}
	} else if data.Ipv4NatUdpTimeout.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.optionType", "default")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.value", 1)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.udpTimeout.value", data.Ipv4NatUdpTimeout.ValueInt64())
		}
	}

	if !data.Ipv4NatTcpTimeoutVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.optionType", "variable")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.value", data.Ipv4NatTcpTimeoutVariable.ValueString())
		}
	} else if data.Ipv4NatTcpTimeout.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.optionType", "default")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.value", 60)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv4.tcpTimeout.value", data.Ipv4NatTcpTimeout.ValueInt64())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"natAttributesIpv4.newStaticNat", []interface{}{})
		for _, item := range data.StaticNats {
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

			if !item.TranslateIpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "translateIp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "translateIp.value", item.TranslateIpVariable.ValueString())
				}
			} else if !item.TranslateIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "translateIp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "translateIp.value", item.TranslateIp.ValueString())
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
	if data.Ipv6Nat.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natIpv6.optionType", "default")
			body, _ = sjson.Set(body, path+"natIpv6.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"natIpv6.optionType", "global")
			body, _ = sjson.Set(body, path+"natIpv6.value", data.Ipv6Nat.ValueBool())
		}
	}
	if !data.Nat64.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"natAttributesIpv6.nat64.optionType", "global")
			body, _ = sjson.Set(body, path+"natAttributesIpv6.nat64.value", data.Nat64.ValueBool())
		}
	}

	if !data.AclShapingRateVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRate.optionType", "variable")
			body, _ = sjson.Set(body, path+"aclQos.shapingRate.value", data.AclShapingRateVariable.ValueString())
		}
	} else if data.AclShapingRate.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRate.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.shapingRate.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.shapingRate.value", data.AclShapingRate.ValueInt64())
		}
	}
	if !data.AclIpv4EgressPolicyId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.ipv4AclEgress.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.ipv4AclEgress.refId.value", data.AclIpv4EgressPolicyId.ValueString())
		}
	}
	if !data.AclIpv4IngressPolicyId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.ipv4AclIngress.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.ipv4AclIngress.refId.value", data.AclIpv4IngressPolicyId.ValueString())
		}
	}
	if !data.AclIpv6EgressPolicyId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.ipv6AclEgress.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.ipv6AclEgress.refId.value", data.AclIpv6EgressPolicyId.ValueString())
		}
	}
	if !data.AclIpv6IngressPolicyId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"aclQos.ipv6AclIngress.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"aclQos.ipv6AclIngress.refId.value", data.AclIpv6IngressPolicyId.ValueString())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"vrrpIpv6", []interface{}{})
		for _, item := range data.Ipv6Vrrps {
			itemBody := ""

			if !item.GroupIdVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "groupId.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "groupId.value", item.GroupIdVariable.ValueString())
				}
			} else if !item.GroupId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "groupId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "groupId.value", item.GroupId.ValueInt64())
				}
			}

			if !item.PriorityVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "priority.value", item.PriorityVariable.ValueString())
				}
			} else if item.Priority.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "priority.value", 100)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "priority.value", item.Priority.ValueInt64())
				}
			}

			if !item.TimerVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "timer.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "timer.value", item.TimerVariable.ValueString())
				}
			} else if item.Timer.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "timer.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "timer.value", 1000)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "timer.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "timer.value", item.Timer.ValueInt64())
				}
			}
			if item.TrackOmp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackOmp.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "trackOmp.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackOmp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "trackOmp.value", item.TrackOmp.ValueBool())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "ipv6", []interface{}{})
				for _, childItem := range item.Ipv6Addresses {
					itemChildBody := ""

					if !childItem.LinkLocalAddressVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "ipv6LinkLocal.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "ipv6LinkLocal.value", childItem.LinkLocalAddressVariable.ValueString())
						}
					} else if !childItem.LinkLocalAddress.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "ipv6LinkLocal.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "ipv6LinkLocal.value", childItem.LinkLocalAddress.ValueString())
						}
					}

					if !childItem.GlobalAddressVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.value", childItem.GlobalAddressVariable.ValueString())
						}
					} else if childItem.GlobalAddress.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.optionType", "default")

						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.value", childItem.GlobalAddress.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipv6.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"vrrpIpv6.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"vrrp", []interface{}{})
		for _, item := range data.Ipv4Vrrps {
			itemBody := ""

			if !item.GroupIdVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "group_id.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "group_id.value", item.GroupIdVariable.ValueString())
				}
			} else if !item.GroupId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "group_id.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "group_id.value", item.GroupId.ValueInt64())
				}
			}

			if !item.PriorityVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "priority.value", item.PriorityVariable.ValueString())
				}
			} else if item.Priority.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "priority.value", 100)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priority.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "priority.value", item.Priority.ValueInt64())
				}
			}

			if !item.TimerVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "timer.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "timer.value", item.TimerVariable.ValueString())
				}
			} else if item.Timer.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "timer.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "timer.value", 1000)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "timer.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "timer.value", item.Timer.ValueInt64())
				}
			}
			if item.TrackOmp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackOmp.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "trackOmp.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackOmp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "trackOmp.value", item.TrackOmp.ValueBool())
				}
			}

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
			if true {
				itemBody, _ = sjson.Set(itemBody, "ipAddressSecondary", []interface{}{})
				for _, childItem := range item.SecondaryAddresses {
					itemChildBody := ""

					if !childItem.AddressVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "ipAddress.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "ipAddress.value", childItem.AddressVariable.ValueString())
						}
					} else if !childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "ipAddress.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "ipAddress.value", childItem.Address.ValueString())
						}
					}

					if !childItem.SubnetMaskVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "subnetMask.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "subnetMask.value", childItem.SubnetMaskVariable.ValueString())
						}
					} else if !childItem.SubnetMask.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "subnetMask.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "subnetMask.value", childItem.SubnetMask.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipAddressSecondary.-1", itemChildBody)
				}
			}
			if item.TlocPrefixChange.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlocPrefChange.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "tlocPrefChange.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlocPrefChange.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tlocPrefChange.value", item.TlocPrefixChange.ValueBool())
				}
			}
			if item.TlocPrefChangeValue.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlocPrefChangeValue.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlocPrefChangeValue.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tlocPrefChangeValue.value", item.TlocPrefChangeValue.ValueInt64())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "trackingObject", []interface{}{})
				for _, childItem := range item.TrackingObjects {
					itemChildBody := ""
					if !childItem.TrackerId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "trackerId.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "trackerId.refId.value", childItem.TrackerId.ValueString())
						}
					}

					if !childItem.TrackerActionVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "trackerAction.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "trackerAction.value", childItem.TrackerActionVariable.ValueString())
						}
					} else if !childItem.TrackerAction.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "trackerAction.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "trackerAction.value", childItem.TrackerAction.ValueString())
						}
					}

					if !childItem.DecrementValueVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "decrementValue.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "decrementValue.value", childItem.DecrementValueVariable.ValueString())
						}
					} else if !childItem.DecrementValue.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "decrementValue.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "decrementValue.value", childItem.DecrementValue.ValueInt64())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "trackingObject.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"vrrp.-1", itemBody)
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
			} else if !item.MacAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "macAddress.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "macAddress.value", item.MacAddress.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"arp.-1", itemBody)
		}
	}
	if data.TrustsecEnableSgtPropogation.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trustsec.enableSGTPropogation.optionType", "default")
			body, _ = sjson.Set(body, path+"trustsec.enableSGTPropogation.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"trustsec.enableSGTPropogation.optionType", "global")
			body, _ = sjson.Set(body, path+"trustsec.enableSGTPropogation.value", data.TrustsecEnableSgtPropogation.ValueBool())
		}
	}
	if data.TrustsecPropogate.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trustsec.propogate.optionType", "default")
			body, _ = sjson.Set(body, path+"trustsec.propogate.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"trustsec.propogate.optionType", "global")
			body, _ = sjson.Set(body, path+"trustsec.propogate.value", data.TrustsecPropogate.ValueBool())
		}
	}

	if !data.TrustsecSecurityGroupTagVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trustsec.securityGroupTag.optionType", "variable")
			body, _ = sjson.Set(body, path+"trustsec.securityGroupTag.value", data.TrustsecSecurityGroupTagVariable.ValueString())
		}
	} else if data.TrustsecSecurityGroupTag.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trustsec.securityGroupTag.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"trustsec.securityGroupTag.optionType", "global")
			body, _ = sjson.Set(body, path+"trustsec.securityGroupTag.value", data.TrustsecSecurityGroupTag.ValueInt64())
		}
	}
	if data.TrustsecEnableEnforcedPropogation.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trustsec.enableEnforcedPropogation.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"trustsec.enableEnforcedPropogation.optionType", "global")
			body, _ = sjson.Set(body, path+"trustsec.enableEnforcedPropogation.value", data.TrustsecEnableEnforcedPropogation.ValueBool())
		}
	}

	if !data.TrustsecEnforcedSecurityGroupTagVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trustsec.enforcedSecurityGroupTag.optionType", "variable")
			body, _ = sjson.Set(body, path+"trustsec.enforcedSecurityGroupTag.value", data.TrustsecEnforcedSecurityGroupTagVariable.ValueString())
		}
	} else if data.TrustsecEnforcedSecurityGroupTag.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trustsec.enforcedSecurityGroupTag.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"trustsec.enforcedSecurityGroupTag.optionType", "global")
			body, _ = sjson.Set(body, path+"trustsec.enforcedSecurityGroupTag.value", data.TrustsecEnforcedSecurityGroupTag.ValueInt64())
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
		if currentVersion.LessThan(MinServiceLANVPNUpdateVersion) {
			if true {
				body, _ = sjson.Set(body, path+"advanced.tracker.optionType", "default")
			}
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.tracker.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.tracker.value", data.Tracker.ValueString())
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

	if !data.XconnectVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.xconnect.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.xconnect.value", data.XconnectVariable.ValueString())
		}
	} else if data.Xconnect.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.xconnect.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.xconnect.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.xconnect.value", data.Xconnect.ValueString())
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

func (data *ServiceLANVPNInterfaceEthernet) fromBody(ctx context.Context, res gjson.Result, currentVersion *version.Version) {
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
		data.Ipv4ConfigurationType = types.StringValue("dynamic")
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
		data.Ipv4ConfigurationType = types.StringValue("static")
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
		data.Ipv4ConfigurationType = types.StringValue("static")
	}
	if value := res.Get(path + "intfIpAddress.static.staticIpV4AddressSecondary"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv4SecondaryAddresses = make([]ServiceLANVPNInterfaceEthernetIpv4SecondaryAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNInterfaceEthernetIpv4SecondaryAddresses{}
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
		data.Ipv4ConfigurationType = types.StringValue("static")
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
		data.Ipv6ConfigurationType = types.StringValue("dynamic")
	}
	if value := res.Get(path + "intfIpV6Address.dynamic.secondaryIpV6Address"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv6DhcpSecondaryAddresses = make([]ServiceLANVPNInterfaceEthernetIpv6DhcpSecondaryAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNInterfaceEthernetIpv6DhcpSecondaryAddresses{}
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
			data.Ipv6DhcpSecondaryAddresses = append(data.Ipv6DhcpSecondaryAddresses, item)
			return true
		})
		data.Ipv6ConfigurationType = types.StringValue("dynamic")
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
		data.Ipv6ConfigurationType = types.StringValue("static")
	}
	if value := res.Get(path + "intfIpV6Address.static.secondaryIpV6Address"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv6SecondaryAddresses = make([]ServiceLANVPNInterfaceEthernetIpv6SecondaryAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNInterfaceEthernetIpv6SecondaryAddresses{}
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
		data.Ipv6ConfigurationType = types.StringValue("static")
	}
	if value := res.Get(path + "intfIpV6Address.static.dhcpHelperV6"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv6DhcpHelpers = make([]ServiceLANVPNInterfaceEthernetIpv6DhcpHelpers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNInterfaceEthernetIpv6DhcpHelpers{}
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
			item.Dhcpv6HelperVpn = types.Int64Null()
			item.Dhcpv6HelperVpnVariable = types.StringNull()
			if t := v.Get("vpn.optionType"); t.Exists() {
				va := v.Get("vpn.value")
				if t.String() == "variable" {
					item.Dhcpv6HelperVpnVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Dhcpv6HelperVpn = types.Int64Value(va.Int())
				}
			}
			data.Ipv6DhcpHelpers = append(data.Ipv6DhcpHelpers, item)
			return true
		})
		data.Ipv6ConfigurationType = types.StringValue("static")
	}
	data.Ipv4Nat = types.BoolNull()

	if t := res.Get(path + "nat.optionType"); t.Exists() {
		va := res.Get(path + "nat.value")
		if t.String() == "global" {
			data.Ipv4Nat = types.BoolValue(va.Bool())
		}
	}
	data.Ipv4NatType = types.StringNull()
	data.Ipv4NatTypeVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natType.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natType.value")
		if t.String() == "variable" {
			data.Ipv4NatTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4NatType = types.StringValue(va.String())
		}
	}
	data.Ipv4NatRangeStart = types.StringNull()
	data.Ipv4NatRangeStartVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natPool.rangeStart.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natPool.rangeStart.value")
		if t.String() == "variable" {
			data.Ipv4NatRangeStartVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4NatRangeStart = types.StringValue(va.String())
		}
	}
	data.Ipv4NatRangeEnd = types.StringNull()
	data.Ipv4NatRangeEndVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natPool.rangeEnd.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natPool.rangeEnd.value")
		if t.String() == "variable" {
			data.Ipv4NatRangeEndVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4NatRangeEnd = types.StringValue(va.String())
		}
	}
	data.Ipv4NatPrefixLength = types.Int64Null()
	data.Ipv4NatPrefixLengthVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natPool.prefixLength.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natPool.prefixLength.value")
		if t.String() == "variable" {
			data.Ipv4NatPrefixLengthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4NatPrefixLength = types.Int64Value(va.Int())
		}
	}
	data.Ipv4NatOverload = types.BoolNull()
	data.Ipv4NatOverloadVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natPool.overload.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natPool.overload.value")
		if t.String() == "variable" {
			data.Ipv4NatOverloadVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4NatOverload = types.BoolValue(va.Bool())
		}
	}
	data.Ipv4NatLoopback = types.StringNull()
	data.Ipv4NatLoopbackVariable = types.StringNull()
	natLookBackRef := "natLookback"
	if !currentVersion.LessThan(MinServiceLANVPNInterfaceEthernetUpdateVersion) {
		natLookBackRef = "natLoopback"
	}
	if t := res.Get(path + "natAttributesIpv4." + natLookBackRef + ".optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4." + natLookBackRef + ".value")
		if t.String() == "variable" {
			data.Ipv4NatLoopbackVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4NatLoopback = types.StringValue(va.String())
		}
	}
	data.Ipv4NatUdpTimeout = types.Int64Null()
	data.Ipv4NatUdpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.udpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.udpTimeout.value")
		if t.String() == "variable" {
			data.Ipv4NatUdpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4NatUdpTimeout = types.Int64Value(va.Int())
		}
	}
	data.Ipv4NatTcpTimeout = types.Int64Null()
	data.Ipv4NatTcpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.tcpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.tcpTimeout.value")
		if t.String() == "variable" {
			data.Ipv4NatTcpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4NatTcpTimeout = types.Int64Value(va.Int())
		}
	}
	if value := res.Get(path + "natAttributesIpv4.newStaticNat"); value.Exists() && len(value.Array()) > 0 {
		data.StaticNats = make([]ServiceLANVPNInterfaceEthernetStaticNats, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNInterfaceEthernetStaticNats{}
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
			item.TranslateIp = types.StringNull()
			item.TranslateIpVariable = types.StringNull()
			if t := v.Get("translateIp.optionType"); t.Exists() {
				va := v.Get("translateIp.value")
				if t.String() == "variable" {
					item.TranslateIpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TranslateIp = types.StringValue(va.String())
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
			data.StaticNats = append(data.StaticNats, item)
			return true
		})
	}
	data.Ipv6Nat = types.BoolNull()

	if t := res.Get(path + "natIpv6.optionType"); t.Exists() {
		va := res.Get(path + "natIpv6.value")
		if t.String() == "global" {
			data.Ipv6Nat = types.BoolValue(va.Bool())
		}
	}
	data.Nat64 = types.BoolNull()

	if t := res.Get(path + "natAttributesIpv6.nat64.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv6.nat64.value")
		if t.String() == "global" {
			data.Nat64 = types.BoolValue(va.Bool())
		}
	}
	data.AclShapingRate = types.Int64Null()
	data.AclShapingRateVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRate.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRate.value")
		if t.String() == "variable" {
			data.AclShapingRateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AclShapingRate = types.Int64Value(va.Int())
		}
	}
	data.AclIpv4EgressPolicyId = types.StringNull()

	if t := res.Get(path + "aclQos.ipv4AclEgress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv4AclEgress.refId.value")
		if t.String() == "global" {
			data.AclIpv4EgressPolicyId = types.StringValue(va.String())
		}
	}
	data.AclIpv4IngressPolicyId = types.StringNull()

	if t := res.Get(path + "aclQos.ipv4AclIngress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv4AclIngress.refId.value")
		if t.String() == "global" {
			data.AclIpv4IngressPolicyId = types.StringValue(va.String())
		}
	}
	data.AclIpv6EgressPolicyId = types.StringNull()

	if t := res.Get(path + "aclQos.ipv6AclEgress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv6AclEgress.refId.value")
		if t.String() == "global" {
			data.AclIpv6EgressPolicyId = types.StringValue(va.String())
		}
	}
	data.AclIpv6IngressPolicyId = types.StringNull()

	if t := res.Get(path + "aclQos.ipv6AclIngress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv6AclIngress.refId.value")
		if t.String() == "global" {
			data.AclIpv6IngressPolicyId = types.StringValue(va.String())
		}
	}
	if value := res.Get(path + "vrrpIpv6"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv6Vrrps = make([]ServiceLANVPNInterfaceEthernetIpv6Vrrps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNInterfaceEthernetIpv6Vrrps{}
			item.GroupId = types.Int64Null()
			item.GroupIdVariable = types.StringNull()
			if t := v.Get("groupId.optionType"); t.Exists() {
				va := v.Get("groupId.value")
				if t.String() == "variable" {
					item.GroupIdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.GroupId = types.Int64Value(va.Int())
				}
			}
			item.Priority = types.Int64Null()
			item.PriorityVariable = types.StringNull()
			if t := v.Get("priority.optionType"); t.Exists() {
				va := v.Get("priority.value")
				if t.String() == "variable" {
					item.PriorityVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Priority = types.Int64Value(va.Int())
				}
			}
			item.Timer = types.Int64Null()
			item.TimerVariable = types.StringNull()
			if t := v.Get("timer.optionType"); t.Exists() {
				va := v.Get("timer.value")
				if t.String() == "variable" {
					item.TimerVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Timer = types.Int64Value(va.Int())
				}
			}
			item.TrackOmp = types.BoolNull()

			if t := v.Get("trackOmp.optionType"); t.Exists() {
				va := v.Get("trackOmp.value")
				if t.String() == "global" {
					item.TrackOmp = types.BoolValue(va.Bool())
				}
			}
			if cValue := v.Get("ipv6"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Ipv6Addresses = make([]ServiceLANVPNInterfaceEthernetIpv6VrrpsIpv6Addresses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceLANVPNInterfaceEthernetIpv6VrrpsIpv6Addresses{}
					cItem.LinkLocalAddress = types.StringNull()
					cItem.LinkLocalAddressVariable = types.StringNull()
					if t := cv.Get("ipv6LinkLocal.optionType"); t.Exists() {
						va := cv.Get("ipv6LinkLocal.value")
						if t.String() == "variable" {
							cItem.LinkLocalAddressVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.LinkLocalAddress = types.StringValue(va.String())
						}
					}
					cItem.GlobalAddress = types.StringNull()
					cItem.GlobalAddressVariable = types.StringNull()
					if t := cv.Get("prefix.optionType"); t.Exists() {
						va := cv.Get("prefix.value")
						if t.String() == "variable" {
							cItem.GlobalAddressVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.GlobalAddress = types.StringValue(va.String())
						}
					}
					item.Ipv6Addresses = append(item.Ipv6Addresses, cItem)
					return true
				})
			}
			data.Ipv6Vrrps = append(data.Ipv6Vrrps, item)
			return true
		})
	}
	if value := res.Get(path + "vrrp"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv4Vrrps = make([]ServiceLANVPNInterfaceEthernetIpv4Vrrps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNInterfaceEthernetIpv4Vrrps{}
			item.GroupId = types.Int64Null()
			item.GroupIdVariable = types.StringNull()
			if t := v.Get("group_id.optionType"); t.Exists() {
				va := v.Get("group_id.value")
				if t.String() == "variable" {
					item.GroupIdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.GroupId = types.Int64Value(va.Int())
				}
			}
			item.Priority = types.Int64Null()
			item.PriorityVariable = types.StringNull()
			if t := v.Get("priority.optionType"); t.Exists() {
				va := v.Get("priority.value")
				if t.String() == "variable" {
					item.PriorityVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Priority = types.Int64Value(va.Int())
				}
			}
			item.Timer = types.Int64Null()
			item.TimerVariable = types.StringNull()
			if t := v.Get("timer.optionType"); t.Exists() {
				va := v.Get("timer.value")
				if t.String() == "variable" {
					item.TimerVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Timer = types.Int64Value(va.Int())
				}
			}
			item.TrackOmp = types.BoolNull()

			if t := v.Get("trackOmp.optionType"); t.Exists() {
				va := v.Get("trackOmp.value")
				if t.String() == "global" {
					item.TrackOmp = types.BoolValue(va.Bool())
				}
			}
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
			if cValue := v.Get("ipAddressSecondary"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.SecondaryAddresses = make([]ServiceLANVPNInterfaceEthernetIpv4VrrpsSecondaryAddresses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceLANVPNInterfaceEthernetIpv4VrrpsSecondaryAddresses{}
					cItem.Address = types.StringNull()
					cItem.AddressVariable = types.StringNull()
					if t := cv.Get("ipAddress.optionType"); t.Exists() {
						va := cv.Get("ipAddress.value")
						if t.String() == "variable" {
							cItem.AddressVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					cItem.SubnetMask = types.StringNull()
					cItem.SubnetMaskVariable = types.StringNull()
					if t := cv.Get("subnetMask.optionType"); t.Exists() {
						va := cv.Get("subnetMask.value")
						if t.String() == "variable" {
							cItem.SubnetMaskVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.SubnetMask = types.StringValue(va.String())
						}
					}
					item.SecondaryAddresses = append(item.SecondaryAddresses, cItem)
					return true
				})
			}
			item.TlocPrefixChange = types.BoolNull()

			if t := v.Get("tlocPrefChange.optionType"); t.Exists() {
				va := v.Get("tlocPrefChange.value")
				if t.String() == "global" {
					item.TlocPrefixChange = types.BoolValue(va.Bool())
				}
			}
			item.TlocPrefChangeValue = types.Int64Null()

			if t := v.Get("tlocPrefChangeValue.optionType"); t.Exists() {
				va := v.Get("tlocPrefChangeValue.value")
				if t.String() == "global" {
					item.TlocPrefChangeValue = types.Int64Value(va.Int())
				}
			}
			if cValue := v.Get("trackingObject"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.TrackingObjects = make([]ServiceLANVPNInterfaceEthernetIpv4VrrpsTrackingObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceLANVPNInterfaceEthernetIpv4VrrpsTrackingObjects{}
					cItem.TrackerId = types.StringNull()

					if t := cv.Get("trackerId.refId.optionType"); t.Exists() {
						va := cv.Get("trackerId.refId.value")
						if t.String() == "global" {
							cItem.TrackerId = types.StringValue(va.String())
						}
					}
					cItem.TrackerAction = types.StringNull()
					cItem.TrackerActionVariable = types.StringNull()
					if t := cv.Get("trackerAction.optionType"); t.Exists() {
						va := cv.Get("trackerAction.value")
						if t.String() == "variable" {
							cItem.TrackerActionVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.TrackerAction = types.StringValue(va.String())
						}
					}
					cItem.DecrementValue = types.Int64Null()
					cItem.DecrementValueVariable = types.StringNull()
					if t := cv.Get("decrementValue.optionType"); t.Exists() {
						va := cv.Get("decrementValue.value")
						if t.String() == "variable" {
							cItem.DecrementValueVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.DecrementValue = types.Int64Value(va.Int())
						}
					}
					item.TrackingObjects = append(item.TrackingObjects, cItem)
					return true
				})
			}
			data.Ipv4Vrrps = append(data.Ipv4Vrrps, item)
			return true
		})
	}
	if value := res.Get(path + "arp"); value.Exists() && len(value.Array()) > 0 {
		data.Arps = make([]ServiceLANVPNInterfaceEthernetArps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNInterfaceEthernetArps{}
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
	data.TrustsecEnableSgtPropogation = types.BoolNull()

	if t := res.Get(path + "trustsec.enableSGTPropogation.optionType"); t.Exists() {
		va := res.Get(path + "trustsec.enableSGTPropogation.value")
		if t.String() == "global" {
			data.TrustsecEnableSgtPropogation = types.BoolValue(va.Bool())
		}
	}
	data.TrustsecPropogate = types.BoolNull()

	if t := res.Get(path + "trustsec.propogate.optionType"); t.Exists() {
		va := res.Get(path + "trustsec.propogate.value")
		if t.String() == "global" {
			data.TrustsecPropogate = types.BoolValue(va.Bool())
		}
	}
	data.TrustsecSecurityGroupTag = types.Int64Null()
	data.TrustsecSecurityGroupTagVariable = types.StringNull()
	if t := res.Get(path + "trustsec.securityGroupTag.optionType"); t.Exists() {
		va := res.Get(path + "trustsec.securityGroupTag.value")
		if t.String() == "variable" {
			data.TrustsecSecurityGroupTagVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrustsecSecurityGroupTag = types.Int64Value(va.Int())
		}
	}
	data.TrustsecEnableEnforcedPropogation = types.BoolNull()

	if t := res.Get(path + "trustsec.enableEnforcedPropogation.optionType"); t.Exists() {
		va := res.Get(path + "trustsec.enableEnforcedPropogation.value")
		if t.String() == "global" {
			data.TrustsecEnableEnforcedPropogation = types.BoolValue(va.Bool())
		}
	}
	data.TrustsecEnforcedSecurityGroupTag = types.Int64Null()
	data.TrustsecEnforcedSecurityGroupTagVariable = types.StringNull()
	if t := res.Get(path + "trustsec.enforcedSecurityGroupTag.optionType"); t.Exists() {
		va := res.Get(path + "trustsec.enforcedSecurityGroupTag.value")
		if t.String() == "variable" {
			data.TrustsecEnforcedSecurityGroupTagVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrustsecEnforcedSecurityGroupTag = types.Int64Value(va.Int())
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
	data.Xconnect = types.StringNull()
	data.XconnectVariable = types.StringNull()
	if t := res.Get(path + "advanced.xconnect.optionType"); t.Exists() {
		va := res.Get(path + "advanced.xconnect.value")
		if t.String() == "variable" {
			data.XconnectVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Xconnect = types.StringValue(va.String())
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

func (data *ServiceLANVPNInterfaceEthernet) updateFromBody(ctx context.Context, res gjson.Result, currentVersion *version.Version) {
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
	for i := range data.Ipv6DhcpSecondaryAddresses {
		keys := [...]string{"address"}
		keyValues := [...]string{data.Ipv6DhcpSecondaryAddresses[i].Address.ValueString()}
		keyValuesVariables := [...]string{data.Ipv6DhcpSecondaryAddresses[i].AddressVariable.ValueString()}

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
		data.Ipv6DhcpSecondaryAddresses[i].Address = types.StringNull()
		data.Ipv6DhcpSecondaryAddresses[i].AddressVariable = types.StringNull()
		if t := r.Get("address.optionType"); t.Exists() {
			va := r.Get("address.value")
			if t.String() == "variable" {
				data.Ipv6DhcpSecondaryAddresses[i].AddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6DhcpSecondaryAddresses[i].Address = types.StringValue(va.String())
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
	for i := range data.Ipv6DhcpHelpers {
		keys := [...]string{"ipAddress", "vpn"}
		keyValues := [...]string{data.Ipv6DhcpHelpers[i].Address.ValueString(), strconv.FormatInt(data.Ipv6DhcpHelpers[i].Dhcpv6HelperVpn.ValueInt64(), 10)}
		keyValuesVariables := [...]string{data.Ipv6DhcpHelpers[i].AddressVariable.ValueString(), data.Ipv6DhcpHelpers[i].Dhcpv6HelperVpnVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "intfIpV6Address.static.dhcpHelperV6").ForEach(
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
		data.Ipv6DhcpHelpers[i].Address = types.StringNull()
		data.Ipv6DhcpHelpers[i].AddressVariable = types.StringNull()
		if t := r.Get("ipAddress.optionType"); t.Exists() {
			va := r.Get("ipAddress.value")
			if t.String() == "variable" {
				data.Ipv6DhcpHelpers[i].AddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6DhcpHelpers[i].Address = types.StringValue(va.String())
			}
		}
		data.Ipv6DhcpHelpers[i].Dhcpv6HelperVpn = types.Int64Null()
		data.Ipv6DhcpHelpers[i].Dhcpv6HelperVpnVariable = types.StringNull()
		if t := r.Get("vpn.optionType"); t.Exists() {
			va := r.Get("vpn.value")
			if t.String() == "variable" {
				data.Ipv6DhcpHelpers[i].Dhcpv6HelperVpnVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6DhcpHelpers[i].Dhcpv6HelperVpn = types.Int64Value(va.Int())
			}
		}
	}
	data.Ipv4Nat = types.BoolNull()

	if t := res.Get(path + "nat.optionType"); t.Exists() {
		va := res.Get(path + "nat.value")
		if t.String() == "global" {
			data.Ipv4Nat = types.BoolValue(va.Bool())
		}
	}
	data.Ipv4NatType = types.StringNull()
	data.Ipv4NatTypeVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natType.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natType.value")
		if t.String() == "variable" {
			data.Ipv4NatTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4NatType = types.StringValue(va.String())
		}
	}
	data.Ipv4NatRangeStart = types.StringNull()
	data.Ipv4NatRangeStartVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natPool.rangeStart.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natPool.rangeStart.value")
		if t.String() == "variable" {
			data.Ipv4NatRangeStartVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4NatRangeStart = types.StringValue(va.String())
		}
	}
	data.Ipv4NatRangeEnd = types.StringNull()
	data.Ipv4NatRangeEndVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natPool.rangeEnd.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natPool.rangeEnd.value")
		if t.String() == "variable" {
			data.Ipv4NatRangeEndVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4NatRangeEnd = types.StringValue(va.String())
		}
	}
	data.Ipv4NatPrefixLength = types.Int64Null()
	data.Ipv4NatPrefixLengthVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natPool.prefixLength.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natPool.prefixLength.value")
		if t.String() == "variable" {
			data.Ipv4NatPrefixLengthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4NatPrefixLength = types.Int64Value(va.Int())
		}
	}
	data.Ipv4NatOverload = types.BoolNull()
	data.Ipv4NatOverloadVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.natPool.overload.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.natPool.overload.value")
		if t.String() == "variable" {
			data.Ipv4NatOverloadVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4NatOverload = types.BoolValue(va.Bool())
		}
	}
	data.Ipv4NatLoopback = types.StringNull()
	data.Ipv4NatLoopbackVariable = types.StringNull()
	natLookBackRef := "natLookback"
	if !currentVersion.LessThan(MinServiceLANVPNInterfaceEthernetUpdateVersion) {
		natLookBackRef = "natLoopback"
	}
	if t := res.Get(path + "natAttributesIpv4." + natLookBackRef + ".optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4." + natLookBackRef + ".value")
		if t.String() == "variable" {
			data.Ipv4NatLoopbackVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4NatLoopback = types.StringValue(va.String())
		}
	}
	data.Ipv4NatUdpTimeout = types.Int64Null()
	data.Ipv4NatUdpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.udpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.udpTimeout.value")
		if t.String() == "variable" {
			data.Ipv4NatUdpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4NatUdpTimeout = types.Int64Value(va.Int())
		}
	}
	data.Ipv4NatTcpTimeout = types.Int64Null()
	data.Ipv4NatTcpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "natAttributesIpv4.tcpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv4.tcpTimeout.value")
		if t.String() == "variable" {
			data.Ipv4NatTcpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4NatTcpTimeout = types.Int64Value(va.Int())
		}
	}
	for i := range data.StaticNats {
		keys := [...]string{"sourceIp", "translateIp", "staticNatDirection", "sourceVpn"}
		keyValues := [...]string{data.StaticNats[i].SourceIp.ValueString(), data.StaticNats[i].TranslateIp.ValueString(), data.StaticNats[i].Direction.ValueString(), strconv.FormatInt(data.StaticNats[i].SourceVpn.ValueInt64(), 10)}
		keyValuesVariables := [...]string{data.StaticNats[i].SourceIpVariable.ValueString(), data.StaticNats[i].TranslateIpVariable.ValueString(), "", data.StaticNats[i].SourceVpnVariable.ValueString()}

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
		data.StaticNats[i].SourceIp = types.StringNull()
		data.StaticNats[i].SourceIpVariable = types.StringNull()
		if t := r.Get("sourceIp.optionType"); t.Exists() {
			va := r.Get("sourceIp.value")
			if t.String() == "variable" {
				data.StaticNats[i].SourceIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.StaticNats[i].SourceIp = types.StringValue(va.String())
			}
		}
		data.StaticNats[i].TranslateIp = types.StringNull()
		data.StaticNats[i].TranslateIpVariable = types.StringNull()
		if t := r.Get("translateIp.optionType"); t.Exists() {
			va := r.Get("translateIp.value")
			if t.String() == "variable" {
				data.StaticNats[i].TranslateIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.StaticNats[i].TranslateIp = types.StringValue(va.String())
			}
		}
		data.StaticNats[i].Direction = types.StringNull()

		if t := r.Get("staticNatDirection.optionType"); t.Exists() {
			va := r.Get("staticNatDirection.value")
			if t.String() == "global" {
				data.StaticNats[i].Direction = types.StringValue(va.String())
			}
		}
		data.StaticNats[i].SourceVpn = types.Int64Null()
		data.StaticNats[i].SourceVpnVariable = types.StringNull()
		if t := r.Get("sourceVpn.optionType"); t.Exists() {
			va := r.Get("sourceVpn.value")
			if t.String() == "variable" {
				data.StaticNats[i].SourceVpnVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.StaticNats[i].SourceVpn = types.Int64Value(va.Int())
			}
		}
	}
	data.Ipv6Nat = types.BoolNull()

	if t := res.Get(path + "natIpv6.optionType"); t.Exists() {
		va := res.Get(path + "natIpv6.value")
		if t.String() == "global" {
			data.Ipv6Nat = types.BoolValue(va.Bool())
		}
	}
	data.Nat64 = types.BoolNull()

	if t := res.Get(path + "natAttributesIpv6.nat64.optionType"); t.Exists() {
		va := res.Get(path + "natAttributesIpv6.nat64.value")
		if t.String() == "global" {
			data.Nat64 = types.BoolValue(va.Bool())
		}
	}
	data.AclShapingRate = types.Int64Null()
	data.AclShapingRateVariable = types.StringNull()
	if t := res.Get(path + "aclQos.shapingRate.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.shapingRate.value")
		if t.String() == "variable" {
			data.AclShapingRateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AclShapingRate = types.Int64Value(va.Int())
		}
	}
	data.AclIpv4EgressPolicyId = types.StringNull()

	if t := res.Get(path + "aclQos.ipv4AclEgress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv4AclEgress.refId.value")
		if t.String() == "global" {
			data.AclIpv4EgressPolicyId = types.StringValue(va.String())
		}
	}
	data.AclIpv4IngressPolicyId = types.StringNull()

	if t := res.Get(path + "aclQos.ipv4AclIngress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv4AclIngress.refId.value")
		if t.String() == "global" {
			data.AclIpv4IngressPolicyId = types.StringValue(va.String())
		}
	}
	data.AclIpv6EgressPolicyId = types.StringNull()

	if t := res.Get(path + "aclQos.ipv6AclEgress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv6AclEgress.refId.value")
		if t.String() == "global" {
			data.AclIpv6EgressPolicyId = types.StringValue(va.String())
		}
	}
	data.AclIpv6IngressPolicyId = types.StringNull()

	if t := res.Get(path + "aclQos.ipv6AclIngress.refId.optionType"); t.Exists() {
		va := res.Get(path + "aclQos.ipv6AclIngress.refId.value")
		if t.String() == "global" {
			data.AclIpv6IngressPolicyId = types.StringValue(va.String())
		}
	}
	for i := range data.Ipv6Vrrps {
		keys := [...]string{"groupId"}
		keyValues := [...]string{strconv.FormatInt(data.Ipv6Vrrps[i].GroupId.ValueInt64(), 10)}
		keyValuesVariables := [...]string{data.Ipv6Vrrps[i].GroupIdVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "vrrpIpv6").ForEach(
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
		data.Ipv6Vrrps[i].GroupId = types.Int64Null()
		data.Ipv6Vrrps[i].GroupIdVariable = types.StringNull()
		if t := r.Get("groupId.optionType"); t.Exists() {
			va := r.Get("groupId.value")
			if t.String() == "variable" {
				data.Ipv6Vrrps[i].GroupIdVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Vrrps[i].GroupId = types.Int64Value(va.Int())
			}
		}
		data.Ipv6Vrrps[i].Priority = types.Int64Null()
		data.Ipv6Vrrps[i].PriorityVariable = types.StringNull()
		if t := r.Get("priority.optionType"); t.Exists() {
			va := r.Get("priority.value")
			if t.String() == "variable" {
				data.Ipv6Vrrps[i].PriorityVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Vrrps[i].Priority = types.Int64Value(va.Int())
			}
		}
		data.Ipv6Vrrps[i].Timer = types.Int64Null()
		data.Ipv6Vrrps[i].TimerVariable = types.StringNull()
		if t := r.Get("timer.optionType"); t.Exists() {
			va := r.Get("timer.value")
			if t.String() == "variable" {
				data.Ipv6Vrrps[i].TimerVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Vrrps[i].Timer = types.Int64Value(va.Int())
			}
		}
		data.Ipv6Vrrps[i].TrackOmp = types.BoolNull()

		if t := r.Get("trackOmp.optionType"); t.Exists() {
			va := r.Get("trackOmp.value")
			if t.String() == "global" {
				data.Ipv6Vrrps[i].TrackOmp = types.BoolValue(va.Bool())
			}
		}
		for ci := range data.Ipv6Vrrps[i].Ipv6Addresses {
			keys := [...]string{"ipv6LinkLocal"}
			keyValues := [...]string{data.Ipv6Vrrps[i].Ipv6Addresses[ci].LinkLocalAddress.ValueString()}
			keyValuesVariables := [...]string{data.Ipv6Vrrps[i].Ipv6Addresses[ci].LinkLocalAddressVariable.ValueString()}

			var cr gjson.Result
			r.Get("ipv6").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			data.Ipv6Vrrps[i].Ipv6Addresses[ci].LinkLocalAddress = types.StringNull()
			data.Ipv6Vrrps[i].Ipv6Addresses[ci].LinkLocalAddressVariable = types.StringNull()
			if t := cr.Get("ipv6LinkLocal.optionType"); t.Exists() {
				va := cr.Get("ipv6LinkLocal.value")
				if t.String() == "variable" {
					data.Ipv6Vrrps[i].Ipv6Addresses[ci].LinkLocalAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv6Vrrps[i].Ipv6Addresses[ci].LinkLocalAddress = types.StringValue(va.String())
				}
			}
			data.Ipv6Vrrps[i].Ipv6Addresses[ci].GlobalAddress = types.StringNull()
			data.Ipv6Vrrps[i].Ipv6Addresses[ci].GlobalAddressVariable = types.StringNull()
			if t := cr.Get("prefix.optionType"); t.Exists() {
				va := cr.Get("prefix.value")
				if t.String() == "variable" {
					data.Ipv6Vrrps[i].Ipv6Addresses[ci].GlobalAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv6Vrrps[i].Ipv6Addresses[ci].GlobalAddress = types.StringValue(va.String())
				}
			}
		}
	}
	for i := range data.Ipv4Vrrps {
		keys := [...]string{"group_id"}
		keyValues := [...]string{strconv.FormatInt(data.Ipv4Vrrps[i].GroupId.ValueInt64(), 10)}
		keyValuesVariables := [...]string{data.Ipv4Vrrps[i].GroupIdVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "vrrp").ForEach(
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
		data.Ipv4Vrrps[i].GroupId = types.Int64Null()
		data.Ipv4Vrrps[i].GroupIdVariable = types.StringNull()
		if t := r.Get("group_id.optionType"); t.Exists() {
			va := r.Get("group_id.value")
			if t.String() == "variable" {
				data.Ipv4Vrrps[i].GroupIdVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Vrrps[i].GroupId = types.Int64Value(va.Int())
			}
		}
		data.Ipv4Vrrps[i].Priority = types.Int64Null()
		data.Ipv4Vrrps[i].PriorityVariable = types.StringNull()
		if t := r.Get("priority.optionType"); t.Exists() {
			va := r.Get("priority.value")
			if t.String() == "variable" {
				data.Ipv4Vrrps[i].PriorityVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Vrrps[i].Priority = types.Int64Value(va.Int())
			}
		}
		data.Ipv4Vrrps[i].Timer = types.Int64Null()
		data.Ipv4Vrrps[i].TimerVariable = types.StringNull()
		if t := r.Get("timer.optionType"); t.Exists() {
			va := r.Get("timer.value")
			if t.String() == "variable" {
				data.Ipv4Vrrps[i].TimerVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Vrrps[i].Timer = types.Int64Value(va.Int())
			}
		}
		data.Ipv4Vrrps[i].TrackOmp = types.BoolNull()

		if t := r.Get("trackOmp.optionType"); t.Exists() {
			va := r.Get("trackOmp.value")
			if t.String() == "global" {
				data.Ipv4Vrrps[i].TrackOmp = types.BoolValue(va.Bool())
			}
		}
		data.Ipv4Vrrps[i].Address = types.StringNull()
		data.Ipv4Vrrps[i].AddressVariable = types.StringNull()
		if t := r.Get("ipAddress.optionType"); t.Exists() {
			va := r.Get("ipAddress.value")
			if t.String() == "variable" {
				data.Ipv4Vrrps[i].AddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Vrrps[i].Address = types.StringValue(va.String())
			}
		}
		for ci := range data.Ipv4Vrrps[i].SecondaryAddresses {
			keys := [...]string{"ipAddress"}
			keyValues := [...]string{data.Ipv4Vrrps[i].SecondaryAddresses[ci].Address.ValueString()}
			keyValuesVariables := [...]string{data.Ipv4Vrrps[i].SecondaryAddresses[ci].AddressVariable.ValueString()}

			var cr gjson.Result
			r.Get("ipAddressSecondary").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			data.Ipv4Vrrps[i].SecondaryAddresses[ci].Address = types.StringNull()
			data.Ipv4Vrrps[i].SecondaryAddresses[ci].AddressVariable = types.StringNull()
			if t := cr.Get("ipAddress.optionType"); t.Exists() {
				va := cr.Get("ipAddress.value")
				if t.String() == "variable" {
					data.Ipv4Vrrps[i].SecondaryAddresses[ci].AddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv4Vrrps[i].SecondaryAddresses[ci].Address = types.StringValue(va.String())
				}
			}
			data.Ipv4Vrrps[i].SecondaryAddresses[ci].SubnetMask = types.StringNull()
			data.Ipv4Vrrps[i].SecondaryAddresses[ci].SubnetMaskVariable = types.StringNull()
			if t := cr.Get("subnetMask.optionType"); t.Exists() {
				va := cr.Get("subnetMask.value")
				if t.String() == "variable" {
					data.Ipv4Vrrps[i].SecondaryAddresses[ci].SubnetMaskVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv4Vrrps[i].SecondaryAddresses[ci].SubnetMask = types.StringValue(va.String())
				}
			}
		}
		data.Ipv4Vrrps[i].TlocPrefixChange = types.BoolNull()

		if t := r.Get("tlocPrefChange.optionType"); t.Exists() {
			va := r.Get("tlocPrefChange.value")
			if t.String() == "global" {
				data.Ipv4Vrrps[i].TlocPrefixChange = types.BoolValue(va.Bool())
			}
		}
		data.Ipv4Vrrps[i].TlocPrefChangeValue = types.Int64Null()

		if t := r.Get("tlocPrefChangeValue.optionType"); t.Exists() {
			va := r.Get("tlocPrefChangeValue.value")
			if t.String() == "global" {
				data.Ipv4Vrrps[i].TlocPrefChangeValue = types.Int64Value(va.Int())
			}
		}
		for ci := range data.Ipv4Vrrps[i].TrackingObjects {
			keys := [...]string{"trackerId.refId"}
			keyValues := [...]string{data.Ipv4Vrrps[i].TrackingObjects[ci].TrackerId.ValueString()}
			keyValuesVariables := [...]string{""}

			var cr gjson.Result
			r.Get("trackingObject").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			data.Ipv4Vrrps[i].TrackingObjects[ci].TrackerId = types.StringNull()

			if t := cr.Get("trackerId.refId.optionType"); t.Exists() {
				va := cr.Get("trackerId.refId.value")
				if t.String() == "global" {
					data.Ipv4Vrrps[i].TrackingObjects[ci].TrackerId = types.StringValue(va.String())
				}
			}
			data.Ipv4Vrrps[i].TrackingObjects[ci].TrackerAction = types.StringNull()
			data.Ipv4Vrrps[i].TrackingObjects[ci].TrackerActionVariable = types.StringNull()
			if t := cr.Get("trackerAction.optionType"); t.Exists() {
				va := cr.Get("trackerAction.value")
				if t.String() == "variable" {
					data.Ipv4Vrrps[i].TrackingObjects[ci].TrackerActionVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv4Vrrps[i].TrackingObjects[ci].TrackerAction = types.StringValue(va.String())
				}
			}
			data.Ipv4Vrrps[i].TrackingObjects[ci].DecrementValue = types.Int64Null()
			data.Ipv4Vrrps[i].TrackingObjects[ci].DecrementValueVariable = types.StringNull()
			if t := cr.Get("decrementValue.optionType"); t.Exists() {
				va := cr.Get("decrementValue.value")
				if t.String() == "variable" {
					data.Ipv4Vrrps[i].TrackingObjects[ci].DecrementValueVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv4Vrrps[i].TrackingObjects[ci].DecrementValue = types.Int64Value(va.Int())
				}
			}
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
	data.TrustsecEnableSgtPropogation = types.BoolNull()

	if t := res.Get(path + "trustsec.enableSGTPropogation.optionType"); t.Exists() {
		va := res.Get(path + "trustsec.enableSGTPropogation.value")
		if t.String() == "global" {
			data.TrustsecEnableSgtPropogation = types.BoolValue(va.Bool())
		}
	}
	data.TrustsecPropogate = types.BoolNull()

	if t := res.Get(path + "trustsec.propogate.optionType"); t.Exists() {
		va := res.Get(path + "trustsec.propogate.value")
		if t.String() == "global" {
			data.TrustsecPropogate = types.BoolValue(va.Bool())
		}
	}
	data.TrustsecSecurityGroupTag = types.Int64Null()
	data.TrustsecSecurityGroupTagVariable = types.StringNull()
	if t := res.Get(path + "trustsec.securityGroupTag.optionType"); t.Exists() {
		va := res.Get(path + "trustsec.securityGroupTag.value")
		if t.String() == "variable" {
			data.TrustsecSecurityGroupTagVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrustsecSecurityGroupTag = types.Int64Value(va.Int())
		}
	}
	data.TrustsecEnableEnforcedPropogation = types.BoolNull()

	if t := res.Get(path + "trustsec.enableEnforcedPropogation.optionType"); t.Exists() {
		va := res.Get(path + "trustsec.enableEnforcedPropogation.value")
		if t.String() == "global" {
			data.TrustsecEnableEnforcedPropogation = types.BoolValue(va.Bool())
		}
	}
	data.TrustsecEnforcedSecurityGroupTag = types.Int64Null()
	data.TrustsecEnforcedSecurityGroupTagVariable = types.StringNull()
	if t := res.Get(path + "trustsec.enforcedSecurityGroupTag.optionType"); t.Exists() {
		va := res.Get(path + "trustsec.enforcedSecurityGroupTag.value")
		if t.String() == "variable" {
			data.TrustsecEnforcedSecurityGroupTagVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrustsecEnforcedSecurityGroupTag = types.Int64Value(va.Int())
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
	data.Xconnect = types.StringNull()
	data.XconnectVariable = types.StringNull()
	if t := res.Get(path + "advanced.xconnect.optionType"); t.Exists() {
		va := res.Get(path + "advanced.xconnect.value")
		if t.String() == "variable" {
			data.XconnectVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Xconnect = types.StringValue(va.String())
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
