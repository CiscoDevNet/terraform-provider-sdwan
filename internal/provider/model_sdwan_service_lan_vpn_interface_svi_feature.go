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
type ServiceLANVPNInterfaceSVI struct {
	Id                           types.String                                      `tfsdk:"id"`
	Version                      types.Int64                                       `tfsdk:"version"`
	Name                         types.String                                      `tfsdk:"name"`
	Description                  types.String                                      `tfsdk:"description"`
	FeatureProfileId             types.String                                      `tfsdk:"feature_profile_id"`
	ServiceLanVpnFeatureId       types.String                                      `tfsdk:"service_lan_vpn_feature_id"`
	Shutdown                     types.Bool                                        `tfsdk:"shutdown"`
	ShutdownVariable             types.String                                      `tfsdk:"shutdown_variable"`
	InterfaceName                types.String                                      `tfsdk:"interface_name"`
	InterfaceNameVariable        types.String                                      `tfsdk:"interface_name_variable"`
	InterfaceDescription         types.String                                      `tfsdk:"interface_description"`
	InterfaceDescriptionVariable types.String                                      `tfsdk:"interface_description_variable"`
	InterfaceMtu                 types.Int64                                       `tfsdk:"interface_mtu"`
	InterfaceMtuVariable         types.String                                      `tfsdk:"interface_mtu_variable"`
	IpMtu                        types.Int64                                       `tfsdk:"ip_mtu"`
	IpMtuVariable                types.String                                      `tfsdk:"ip_mtu_variable"`
	Ipv4Address                  types.String                                      `tfsdk:"ipv4_address"`
	Ipv4AddressVariable          types.String                                      `tfsdk:"ipv4_address_variable"`
	Ipv4SubnetMask               types.String                                      `tfsdk:"ipv4_subnet_mask"`
	Ipv4SubnetMaskVariable       types.String                                      `tfsdk:"ipv4_subnet_mask_variable"`
	Ipv4SecondaryAddresses       []ServiceLANVPNInterfaceSVIIpv4SecondaryAddresses `tfsdk:"ipv4_secondary_addresses"`
	Ipv4DhcpHelpers              types.Set                                         `tfsdk:"ipv4_dhcp_helpers"`
	Ipv4DhcpHelpersVariable      types.String                                      `tfsdk:"ipv4_dhcp_helpers_variable"`
	Ipv6Address                  types.String                                      `tfsdk:"ipv6_address"`
	Ipv6AddressVariable          types.String                                      `tfsdk:"ipv6_address_variable"`
	Ipv6SecondaryAddresses       []ServiceLANVPNInterfaceSVIIpv6SecondaryAddresses `tfsdk:"ipv6_secondary_addresses"`
	Ipv6DhcpHelpers              []ServiceLANVPNInterfaceSVIIpv6DhcpHelpers        `tfsdk:"ipv6_dhcp_helpers"`
	Arps                         []ServiceLANVPNInterfaceSVIArps                   `tfsdk:"arps"`
	Ipv4Vrrps                    []ServiceLANVPNInterfaceSVIIpv4Vrrps              `tfsdk:"ipv4_vrrps"`
	Ipv6Vrrps                    []ServiceLANVPNInterfaceSVIIpv6Vrrps              `tfsdk:"ipv6_vrrps"`
	EnableDhcpv6                 types.Bool                                        `tfsdk:"enable_dhcpv6"`
	EnableDhcpv6Variable         types.String                                      `tfsdk:"enable_dhcpv6_variable"`
	TcpMss                       types.Int64                                       `tfsdk:"tcp_mss"`
	TcpMssVariable               types.String                                      `tfsdk:"tcp_mss_variable"`
	ArpTimeout                   types.Int64                                       `tfsdk:"arp_timeout"`
	ArpTimeoutVariable           types.String                                      `tfsdk:"arp_timeout_variable"`
	IpDirectedBroadcast          types.Bool                                        `tfsdk:"ip_directed_broadcast"`
	IpDirectedBroadcastVariable  types.String                                      `tfsdk:"ip_directed_broadcast_variable"`
	IcmpRedirectDisable          types.Bool                                        `tfsdk:"icmp_redirect_disable"`
	IcmpRedirectDisableVariable  types.String                                      `tfsdk:"icmp_redirect_disable_variable"`
}

type ServiceLANVPNInterfaceSVIIpv4SecondaryAddresses struct {
	Address                types.String `tfsdk:"address"`
	AddressVariable        types.String `tfsdk:"address_variable"`
	Ipv4SubnetMask         types.String `tfsdk:"ipv4_subnet_mask"`
	Ipv4SubnetMaskVariable types.String `tfsdk:"ipv4_subnet_mask_variable"`
}

type ServiceLANVPNInterfaceSVIIpv6SecondaryAddresses struct {
	Address         types.String `tfsdk:"address"`
	AddressVariable types.String `tfsdk:"address_variable"`
}

type ServiceLANVPNInterfaceSVIIpv6DhcpHelpers struct {
	Address         types.String `tfsdk:"address"`
	AddressVariable types.String `tfsdk:"address_variable"`
	Vpn             types.Int64  `tfsdk:"vpn"`
	VpnVariable     types.String `tfsdk:"vpn_variable"`
}

type ServiceLANVPNInterfaceSVIArps struct {
	IpAddress          types.String `tfsdk:"ip_address"`
	IpAddressVariable  types.String `tfsdk:"ip_address_variable"`
	MacAddress         types.String `tfsdk:"mac_address"`
	MacAddressVariable types.String `tfsdk:"mac_address_variable"`
}

type ServiceLANVPNInterfaceSVIIpv4Vrrps struct {
	GroupId                       types.Int64                                            `tfsdk:"group_id"`
	GroupIdVariable               types.String                                           `tfsdk:"group_id_variable"`
	Priority                      types.Int64                                            `tfsdk:"priority"`
	PriorityVariable              types.String                                           `tfsdk:"priority_variable"`
	Timer                         types.Int64                                            `tfsdk:"timer"`
	TimerVariable                 types.String                                           `tfsdk:"timer_variable"`
	TrackOmp                      types.Bool                                             `tfsdk:"track_omp"`
	TrackOmpVariable              types.String                                           `tfsdk:"track_omp_variable"`
	PrefixList                    types.String                                           `tfsdk:"prefix_list"`
	PrefixListVariable            types.String                                           `tfsdk:"prefix_list_variable"`
	Address                       types.String                                           `tfsdk:"address"`
	AddressVariable               types.String                                           `tfsdk:"address_variable"`
	SecondaryAddresses            []ServiceLANVPNInterfaceSVIIpv4VrrpsSecondaryAddresses `tfsdk:"secondary_addresses"`
	TlocPrefixChange              types.Bool                                             `tfsdk:"tloc_prefix_change"`
	TlocPrefixChangeValue         types.Int64                                            `tfsdk:"tloc_prefix_change_value"`
	TlocPrefixChangeValueVariable types.String                                           `tfsdk:"tloc_prefix_change_value_variable"`
}

type ServiceLANVPNInterfaceSVIIpv6Vrrps struct {
	GroupId                 types.Int64                                            `tfsdk:"group_id"`
	GroupIdVariable         types.String                                           `tfsdk:"group_id_variable"`
	Priority                types.Int64                                            `tfsdk:"priority"`
	PriorityVariable        types.String                                           `tfsdk:"priority_variable"`
	Timer                   types.Int64                                            `tfsdk:"timer"`
	TimerVariable           types.String                                           `tfsdk:"timer_variable"`
	TrackOmp                types.Bool                                             `tfsdk:"track_omp"`
	TrackOmpVariable        types.String                                           `tfsdk:"track_omp_variable"`
	TrackPrefixList         types.String                                           `tfsdk:"track_prefix_list"`
	TrackPrefixListVariable types.String                                           `tfsdk:"track_prefix_list_variable"`
	Addresses               []ServiceLANVPNInterfaceSVIIpv6VrrpsAddresses          `tfsdk:"addresses"`
	SecondaryAddresses      []ServiceLANVPNInterfaceSVIIpv6VrrpsSecondaryAddresses `tfsdk:"secondary_addresses"`
}

type ServiceLANVPNInterfaceSVIIpv4VrrpsSecondaryAddresses struct {
	Address         types.String `tfsdk:"address"`
	AddressVariable types.String `tfsdk:"address_variable"`
}

type ServiceLANVPNInterfaceSVIIpv6VrrpsAddresses struct {
	LinkLocalAddress         types.String `tfsdk:"link_local_address"`
	LinkLocalAddressVariable types.String `tfsdk:"link_local_address_variable"`
	GlobalAddress            types.String `tfsdk:"global_address"`
	GlobalAddressVariable    types.String `tfsdk:"global_address_variable"`
}
type ServiceLANVPNInterfaceSVIIpv6VrrpsSecondaryAddresses struct {
	Prefix         types.String `tfsdk:"prefix"`
	PrefixVariable types.String `tfsdk:"prefix_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceLANVPNInterfaceSVI) getModel() string {
	return "service_lan_vpn_interface_svi"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceLANVPNInterfaceSVI) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/lan/vpn/%s/interface/svi", url.QueryEscape(data.FeatureProfileId.ValueString()), url.QueryEscape(data.ServiceLanVpnFeatureId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceLANVPNInterfaceSVI) toBody(ctx context.Context) string {
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

	if !data.InterfaceMtuVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ifMtu.optionType", "variable")
			body, _ = sjson.Set(body, path+"ifMtu.value", data.InterfaceMtuVariable.ValueString())
		}
	} else if data.InterfaceMtu.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ifMtu.optionType", "default")
			body, _ = sjson.Set(body, path+"ifMtu.value", 1500)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ifMtu.optionType", "global")
			body, _ = sjson.Set(body, path+"ifMtu.value", data.InterfaceMtu.ValueInt64())
		}
	}

	if !data.IpMtuVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ipMtu.optionType", "variable")
			body, _ = sjson.Set(body, path+"ipMtu.value", data.IpMtuVariable.ValueString())
		}
	} else if data.IpMtu.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ipMtu.optionType", "default")
			body, _ = sjson.Set(body, path+"ipMtu.value", 1500)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ipMtu.optionType", "global")
			body, _ = sjson.Set(body, path+"ipMtu.value", data.IpMtu.ValueInt64())
		}
	}

	if !data.Ipv4AddressVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ipv4.addressV4.ipAddress.optionType", "variable")
			body, _ = sjson.Set(body, path+"ipv4.addressV4.ipAddress.value", data.Ipv4AddressVariable.ValueString())
		}
	} else if !data.Ipv4Address.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ipv4.addressV4.ipAddress.optionType", "global")
			body, _ = sjson.Set(body, path+"ipv4.addressV4.ipAddress.value", data.Ipv4Address.ValueString())
		}
	}

	if !data.Ipv4SubnetMaskVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ipv4.addressV4.subnetMask.optionType", "variable")
			body, _ = sjson.Set(body, path+"ipv4.addressV4.subnetMask.value", data.Ipv4SubnetMaskVariable.ValueString())
		}
	} else if !data.Ipv4SubnetMask.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ipv4.addressV4.subnetMask.optionType", "global")
			body, _ = sjson.Set(body, path+"ipv4.addressV4.subnetMask.value", data.Ipv4SubnetMask.ValueString())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"ipv4.secondaryAddressV4", []interface{}{})
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

			if !item.Ipv4SubnetMaskVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "subnetMask.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "subnetMask.value", item.Ipv4SubnetMaskVariable.ValueString())
				}
			} else if !item.Ipv4SubnetMask.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "subnetMask.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "subnetMask.value", item.Ipv4SubnetMask.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"ipv4.secondaryAddressV4.-1", itemBody)
		}
	}

	if !data.Ipv4DhcpHelpersVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ipv4.dhcpHelperV4.optionType", "variable")
			body, _ = sjson.Set(body, path+"ipv4.dhcpHelperV4.value", data.Ipv4DhcpHelpersVariable.ValueString())
		}
	} else if data.Ipv4DhcpHelpers.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ipv4.dhcpHelperV4.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ipv4.dhcpHelperV4.optionType", "global")
			var values []string
			data.Ipv4DhcpHelpers.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"ipv4.dhcpHelperV4.value", values)
		}
	}

	if !data.Ipv6AddressVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ipv6.addressV6.optionType", "variable")
			body, _ = sjson.Set(body, path+"ipv6.addressV6.value", data.Ipv6AddressVariable.ValueString())
		}
	} else if data.Ipv6Address.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ipv6.addressV6.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ipv6.addressV6.optionType", "global")
			body, _ = sjson.Set(body, path+"ipv6.addressV6.value", data.Ipv6Address.ValueString())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"ipv6.secondaryAddressV6", []interface{}{})
		for _, item := range data.Ipv6SecondaryAddresses {
			itemBody := ""

			if !item.AddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.AddressVariable.ValueString())
				}
			} else if item.Address.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.Address.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"ipv6.secondaryAddressV6.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"ipv6.dhcpHelperV6", []interface{}{})
		for _, item := range data.Ipv6DhcpHelpers {
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

			if !item.VpnVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.VpnVariable.ValueString())
				}
			} else if item.Vpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"ipv6.dhcpHelperV6.-1", itemBody)
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

			if !item.TrackOmpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackOmp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "trackOmp.value", item.TrackOmpVariable.ValueString())
				}
			} else if item.TrackOmp.IsNull() {
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

			if !item.PrefixListVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefixList.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefixList.value", item.PrefixListVariable.ValueString())
				}
			} else if item.PrefixList.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefixList.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefixList.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefixList.value", item.PrefixList.ValueString())
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
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.AddressVariable.ValueString())
						}
					} else if !childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
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

			if !item.TlocPrefixChangeValueVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlocPrefChangeValue.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "tlocPrefChangeValue.value", item.TlocPrefixChangeValueVariable.ValueString())
				}
			} else if !item.TlocPrefixChangeValue.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tlocPrefChangeValue.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tlocPrefChangeValue.value", item.TlocPrefixChangeValue.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"vrrp.-1", itemBody)
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

			if !item.TrackOmpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackOmp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "trackOmp.value", item.TrackOmpVariable.ValueString())
				}
			} else if item.TrackOmp.IsNull() {
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

			if !item.TrackPrefixListVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackPrefixList.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "trackPrefixList.value", item.TrackPrefixListVariable.ValueString())
				}
			} else if item.TrackPrefixList.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackPrefixList.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackPrefixList.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "trackPrefixList.value", item.TrackPrefixList.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "ipv6", []interface{}{})
				for _, childItem := range item.Addresses {
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
			if true {
				itemBody, _ = sjson.Set(itemBody, "ipv6Secondary", []interface{}{})
				for _, childItem := range item.SecondaryAddresses {
					itemChildBody := ""

					if !childItem.PrefixVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.value", childItem.PrefixVariable.ValueString())
						}
					} else if !childItem.Prefix.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.value", childItem.Prefix.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipv6Secondary.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"vrrpIpv6.-1", itemBody)
		}
	}

	if !data.EnableDhcpv6Variable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dhcpClientV6.optionType", "variable")
			body, _ = sjson.Set(body, path+"dhcpClientV6.value", data.EnableDhcpv6Variable.ValueString())
		}
	} else if data.EnableDhcpv6.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dhcpClientV6.optionType", "default")
			body, _ = sjson.Set(body, path+"dhcpClientV6.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"dhcpClientV6.optionType", "global")
			body, _ = sjson.Set(body, path+"dhcpClientV6.value", data.EnableDhcpv6.ValueBool())
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
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceLANVPNInterfaceSVI) fromBody(ctx context.Context, res gjson.Result) {
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
	data.InterfaceMtu = types.Int64Null()
	data.InterfaceMtuVariable = types.StringNull()
	if t := res.Get(path + "ifMtu.optionType"); t.Exists() {
		va := res.Get(path + "ifMtu.value")
		if t.String() == "variable" {
			data.InterfaceMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceMtu = types.Int64Value(va.Int())
		}
	}
	data.IpMtu = types.Int64Null()
	data.IpMtuVariable = types.StringNull()
	if t := res.Get(path + "ipMtu.optionType"); t.Exists() {
		va := res.Get(path + "ipMtu.value")
		if t.String() == "variable" {
			data.IpMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpMtu = types.Int64Value(va.Int())
		}
	}
	data.Ipv4Address = types.StringNull()
	data.Ipv4AddressVariable = types.StringNull()
	if t := res.Get(path + "ipv4.addressV4.ipAddress.optionType"); t.Exists() {
		va := res.Get(path + "ipv4.addressV4.ipAddress.value")
		if t.String() == "variable" {
			data.Ipv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Address = types.StringValue(va.String())
		}
	}
	data.Ipv4SubnetMask = types.StringNull()
	data.Ipv4SubnetMaskVariable = types.StringNull()
	if t := res.Get(path + "ipv4.addressV4.subnetMask.optionType"); t.Exists() {
		va := res.Get(path + "ipv4.addressV4.subnetMask.value")
		if t.String() == "variable" {
			data.Ipv4SubnetMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4SubnetMask = types.StringValue(va.String())
		}
	}
	if value := res.Get(path + "ipv4.secondaryAddressV4"); value.Exists() {
		data.Ipv4SecondaryAddresses = make([]ServiceLANVPNInterfaceSVIIpv4SecondaryAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNInterfaceSVIIpv4SecondaryAddresses{}
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
			item.Ipv4SubnetMask = types.StringNull()
			item.Ipv4SubnetMaskVariable = types.StringNull()
			if t := v.Get("subnetMask.optionType"); t.Exists() {
				va := v.Get("subnetMask.value")
				if t.String() == "variable" {
					item.Ipv4SubnetMaskVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Ipv4SubnetMask = types.StringValue(va.String())
				}
			}
			data.Ipv4SecondaryAddresses = append(data.Ipv4SecondaryAddresses, item)
			return true
		})
	}
	data.Ipv4DhcpHelpers = types.SetNull(types.StringType)
	data.Ipv4DhcpHelpersVariable = types.StringNull()
	if t := res.Get(path + "ipv4.dhcpHelperV4.optionType"); t.Exists() {
		va := res.Get(path + "ipv4.dhcpHelperV4.value")
		if t.String() == "variable" {
			data.Ipv4DhcpHelpersVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4DhcpHelpers = helpers.GetStringSet(va.Array())
		}
	}
	data.Ipv6Address = types.StringNull()
	data.Ipv6AddressVariable = types.StringNull()
	if t := res.Get(path + "ipv6.addressV6.optionType"); t.Exists() {
		va := res.Get(path + "ipv6.addressV6.value")
		if t.String() == "variable" {
			data.Ipv6AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6Address = types.StringValue(va.String())
		}
	}
	if value := res.Get(path + "ipv6.secondaryAddressV6"); value.Exists() {
		data.Ipv6SecondaryAddresses = make([]ServiceLANVPNInterfaceSVIIpv6SecondaryAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNInterfaceSVIIpv6SecondaryAddresses{}
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
	if value := res.Get(path + "ipv6.dhcpHelperV6"); value.Exists() {
		data.Ipv6DhcpHelpers = make([]ServiceLANVPNInterfaceSVIIpv6DhcpHelpers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNInterfaceSVIIpv6DhcpHelpers{}
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
			item.Vpn = types.Int64Null()
			item.VpnVariable = types.StringNull()
			if t := v.Get("vpn.optionType"); t.Exists() {
				va := v.Get("vpn.value")
				if t.String() == "variable" {
					item.VpnVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Vpn = types.Int64Value(va.Int())
				}
			}
			data.Ipv6DhcpHelpers = append(data.Ipv6DhcpHelpers, item)
			return true
		})
	}
	if value := res.Get(path + "arp"); value.Exists() {
		data.Arps = make([]ServiceLANVPNInterfaceSVIArps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNInterfaceSVIArps{}
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
	if value := res.Get(path + "vrrp"); value.Exists() {
		data.Ipv4Vrrps = make([]ServiceLANVPNInterfaceSVIIpv4Vrrps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNInterfaceSVIIpv4Vrrps{}
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
			item.TrackOmpVariable = types.StringNull()
			if t := v.Get("trackOmp.optionType"); t.Exists() {
				va := v.Get("trackOmp.value")
				if t.String() == "variable" {
					item.TrackOmpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TrackOmp = types.BoolValue(va.Bool())
				}
			}
			item.PrefixList = types.StringNull()
			item.PrefixListVariable = types.StringNull()
			if t := v.Get("prefixList.optionType"); t.Exists() {
				va := v.Get("prefixList.value")
				if t.String() == "variable" {
					item.PrefixListVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.PrefixList = types.StringValue(va.String())
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
			if cValue := v.Get("ipAddressSecondary"); cValue.Exists() {
				item.SecondaryAddresses = make([]ServiceLANVPNInterfaceSVIIpv4VrrpsSecondaryAddresses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceLANVPNInterfaceSVIIpv4VrrpsSecondaryAddresses{}
					cItem.Address = types.StringNull()
					cItem.AddressVariable = types.StringNull()
					if t := cv.Get("address.optionType"); t.Exists() {
						va := cv.Get("address.value")
						if t.String() == "variable" {
							cItem.AddressVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Address = types.StringValue(va.String())
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
			item.TlocPrefixChangeValue = types.Int64Null()
			item.TlocPrefixChangeValueVariable = types.StringNull()
			if t := v.Get("tlocPrefChangeValue.optionType"); t.Exists() {
				va := v.Get("tlocPrefChangeValue.value")
				if t.String() == "variable" {
					item.TlocPrefixChangeValueVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TlocPrefixChangeValue = types.Int64Value(va.Int())
				}
			}
			data.Ipv4Vrrps = append(data.Ipv4Vrrps, item)
			return true
		})
	}
	if value := res.Get(path + "vrrpIpv6"); value.Exists() {
		data.Ipv6Vrrps = make([]ServiceLANVPNInterfaceSVIIpv6Vrrps, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNInterfaceSVIIpv6Vrrps{}
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
			item.TrackOmpVariable = types.StringNull()
			if t := v.Get("trackOmp.optionType"); t.Exists() {
				va := v.Get("trackOmp.value")
				if t.String() == "variable" {
					item.TrackOmpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TrackOmp = types.BoolValue(va.Bool())
				}
			}
			item.TrackPrefixList = types.StringNull()
			item.TrackPrefixListVariable = types.StringNull()
			if t := v.Get("trackPrefixList.optionType"); t.Exists() {
				va := v.Get("trackPrefixList.value")
				if t.String() == "variable" {
					item.TrackPrefixListVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TrackPrefixList = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("ipv6"); cValue.Exists() {
				item.Addresses = make([]ServiceLANVPNInterfaceSVIIpv6VrrpsAddresses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceLANVPNInterfaceSVIIpv6VrrpsAddresses{}
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
					item.Addresses = append(item.Addresses, cItem)
					return true
				})
			}
			if cValue := v.Get("ipv6Secondary"); cValue.Exists() {
				item.SecondaryAddresses = make([]ServiceLANVPNInterfaceSVIIpv6VrrpsSecondaryAddresses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceLANVPNInterfaceSVIIpv6VrrpsSecondaryAddresses{}
					cItem.Prefix = types.StringNull()
					cItem.PrefixVariable = types.StringNull()
					if t := cv.Get("prefix.optionType"); t.Exists() {
						va := cv.Get("prefix.value")
						if t.String() == "variable" {
							cItem.PrefixVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Prefix = types.StringValue(va.String())
						}
					}
					item.SecondaryAddresses = append(item.SecondaryAddresses, cItem)
					return true
				})
			}
			data.Ipv6Vrrps = append(data.Ipv6Vrrps, item)
			return true
		})
	}
	data.EnableDhcpv6 = types.BoolNull()
	data.EnableDhcpv6Variable = types.StringNull()
	if t := res.Get(path + "dhcpClientV6.optionType"); t.Exists() {
		va := res.Get(path + "dhcpClientV6.value")
		if t.String() == "variable" {
			data.EnableDhcpv6Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EnableDhcpv6 = types.BoolValue(va.Bool())
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ServiceLANVPNInterfaceSVI) updateFromBody(ctx context.Context, res gjson.Result) {
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
	data.InterfaceMtu = types.Int64Null()
	data.InterfaceMtuVariable = types.StringNull()
	if t := res.Get(path + "ifMtu.optionType"); t.Exists() {
		va := res.Get(path + "ifMtu.value")
		if t.String() == "variable" {
			data.InterfaceMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceMtu = types.Int64Value(va.Int())
		}
	}
	data.IpMtu = types.Int64Null()
	data.IpMtuVariable = types.StringNull()
	if t := res.Get(path + "ipMtu.optionType"); t.Exists() {
		va := res.Get(path + "ipMtu.value")
		if t.String() == "variable" {
			data.IpMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpMtu = types.Int64Value(va.Int())
		}
	}
	data.Ipv4Address = types.StringNull()
	data.Ipv4AddressVariable = types.StringNull()
	if t := res.Get(path + "ipv4.addressV4.ipAddress.optionType"); t.Exists() {
		va := res.Get(path + "ipv4.addressV4.ipAddress.value")
		if t.String() == "variable" {
			data.Ipv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Address = types.StringValue(va.String())
		}
	}
	data.Ipv4SubnetMask = types.StringNull()
	data.Ipv4SubnetMaskVariable = types.StringNull()
	if t := res.Get(path + "ipv4.addressV4.subnetMask.optionType"); t.Exists() {
		va := res.Get(path + "ipv4.addressV4.subnetMask.value")
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
		res.Get(path + "ipv4.secondaryAddressV4").ForEach(
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
		data.Ipv4SecondaryAddresses[i].Ipv4SubnetMask = types.StringNull()
		data.Ipv4SecondaryAddresses[i].Ipv4SubnetMaskVariable = types.StringNull()
		if t := r.Get("subnetMask.optionType"); t.Exists() {
			va := r.Get("subnetMask.value")
			if t.String() == "variable" {
				data.Ipv4SecondaryAddresses[i].Ipv4SubnetMaskVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4SecondaryAddresses[i].Ipv4SubnetMask = types.StringValue(va.String())
			}
		}
	}
	data.Ipv4DhcpHelpers = types.SetNull(types.StringType)
	data.Ipv4DhcpHelpersVariable = types.StringNull()
	if t := res.Get(path + "ipv4.dhcpHelperV4.optionType"); t.Exists() {
		va := res.Get(path + "ipv4.dhcpHelperV4.value")
		if t.String() == "variable" {
			data.Ipv4DhcpHelpersVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4DhcpHelpers = helpers.GetStringSet(va.Array())
		}
	}
	data.Ipv6Address = types.StringNull()
	data.Ipv6AddressVariable = types.StringNull()
	if t := res.Get(path + "ipv6.addressV6.optionType"); t.Exists() {
		va := res.Get(path + "ipv6.addressV6.value")
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
		res.Get(path + "ipv6.secondaryAddressV6").ForEach(
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
		keys := [...]string{"address"}
		keyValues := [...]string{data.Ipv6DhcpHelpers[i].Address.ValueString()}
		keyValuesVariables := [...]string{data.Ipv6DhcpHelpers[i].AddressVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "ipv6.dhcpHelperV6").ForEach(
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
		data.Ipv6DhcpHelpers[i].Address = types.StringNull()
		data.Ipv6DhcpHelpers[i].AddressVariable = types.StringNull()
		if t := r.Get("address.optionType"); t.Exists() {
			va := r.Get("address.value")
			if t.String() == "variable" {
				data.Ipv6DhcpHelpers[i].AddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6DhcpHelpers[i].Address = types.StringValue(va.String())
			}
		}
		data.Ipv6DhcpHelpers[i].Vpn = types.Int64Null()
		data.Ipv6DhcpHelpers[i].VpnVariable = types.StringNull()
		if t := r.Get("vpn.optionType"); t.Exists() {
			va := r.Get("vpn.value")
			if t.String() == "variable" {
				data.Ipv6DhcpHelpers[i].VpnVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6DhcpHelpers[i].Vpn = types.Int64Value(va.Int())
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
		data.Ipv4Vrrps[i].TrackOmpVariable = types.StringNull()
		if t := r.Get("trackOmp.optionType"); t.Exists() {
			va := r.Get("trackOmp.value")
			if t.String() == "variable" {
				data.Ipv4Vrrps[i].TrackOmpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Vrrps[i].TrackOmp = types.BoolValue(va.Bool())
			}
		}
		data.Ipv4Vrrps[i].PrefixList = types.StringNull()
		data.Ipv4Vrrps[i].PrefixListVariable = types.StringNull()
		if t := r.Get("prefixList.optionType"); t.Exists() {
			va := r.Get("prefixList.value")
			if t.String() == "variable" {
				data.Ipv4Vrrps[i].PrefixListVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Vrrps[i].PrefixList = types.StringValue(va.String())
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
			keys := [...]string{"address"}
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
			if t := cr.Get("address.optionType"); t.Exists() {
				va := cr.Get("address.value")
				if t.String() == "variable" {
					data.Ipv4Vrrps[i].SecondaryAddresses[ci].AddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv4Vrrps[i].SecondaryAddresses[ci].Address = types.StringValue(va.String())
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
		data.Ipv4Vrrps[i].TlocPrefixChangeValue = types.Int64Null()
		data.Ipv4Vrrps[i].TlocPrefixChangeValueVariable = types.StringNull()
		if t := r.Get("tlocPrefChangeValue.optionType"); t.Exists() {
			va := r.Get("tlocPrefChangeValue.value")
			if t.String() == "variable" {
				data.Ipv4Vrrps[i].TlocPrefixChangeValueVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Vrrps[i].TlocPrefixChangeValue = types.Int64Value(va.Int())
			}
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
		data.Ipv6Vrrps[i].TrackOmpVariable = types.StringNull()
		if t := r.Get("trackOmp.optionType"); t.Exists() {
			va := r.Get("trackOmp.value")
			if t.String() == "variable" {
				data.Ipv6Vrrps[i].TrackOmpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Vrrps[i].TrackOmp = types.BoolValue(va.Bool())
			}
		}
		data.Ipv6Vrrps[i].TrackPrefixList = types.StringNull()
		data.Ipv6Vrrps[i].TrackPrefixListVariable = types.StringNull()
		if t := r.Get("trackPrefixList.optionType"); t.Exists() {
			va := r.Get("trackPrefixList.value")
			if t.String() == "variable" {
				data.Ipv6Vrrps[i].TrackPrefixListVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Vrrps[i].TrackPrefixList = types.StringValue(va.String())
			}
		}
		for ci := range data.Ipv6Vrrps[i].Addresses {
			keys := [...]string{"ipv6LinkLocal"}
			keyValues := [...]string{data.Ipv6Vrrps[i].Addresses[ci].LinkLocalAddress.ValueString()}
			keyValuesVariables := [...]string{data.Ipv6Vrrps[i].Addresses[ci].LinkLocalAddressVariable.ValueString()}

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
			data.Ipv6Vrrps[i].Addresses[ci].LinkLocalAddress = types.StringNull()
			data.Ipv6Vrrps[i].Addresses[ci].LinkLocalAddressVariable = types.StringNull()
			if t := cr.Get("ipv6LinkLocal.optionType"); t.Exists() {
				va := cr.Get("ipv6LinkLocal.value")
				if t.String() == "variable" {
					data.Ipv6Vrrps[i].Addresses[ci].LinkLocalAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv6Vrrps[i].Addresses[ci].LinkLocalAddress = types.StringValue(va.String())
				}
			}
			data.Ipv6Vrrps[i].Addresses[ci].GlobalAddress = types.StringNull()
			data.Ipv6Vrrps[i].Addresses[ci].GlobalAddressVariable = types.StringNull()
			if t := cr.Get("prefix.optionType"); t.Exists() {
				va := cr.Get("prefix.value")
				if t.String() == "variable" {
					data.Ipv6Vrrps[i].Addresses[ci].GlobalAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv6Vrrps[i].Addresses[ci].GlobalAddress = types.StringValue(va.String())
				}
			}
		}
		for ci := range data.Ipv6Vrrps[i].SecondaryAddresses {
			keys := [...]string{"prefix"}
			keyValues := [...]string{data.Ipv6Vrrps[i].SecondaryAddresses[ci].Prefix.ValueString()}
			keyValuesVariables := [...]string{data.Ipv6Vrrps[i].SecondaryAddresses[ci].PrefixVariable.ValueString()}

			var cr gjson.Result
			r.Get("ipv6Secondary").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			data.Ipv6Vrrps[i].SecondaryAddresses[ci].Prefix = types.StringNull()
			data.Ipv6Vrrps[i].SecondaryAddresses[ci].PrefixVariable = types.StringNull()
			if t := cr.Get("prefix.optionType"); t.Exists() {
				va := cr.Get("prefix.value")
				if t.String() == "variable" {
					data.Ipv6Vrrps[i].SecondaryAddresses[ci].PrefixVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv6Vrrps[i].SecondaryAddresses[ci].Prefix = types.StringValue(va.String())
				}
			}
		}
	}
	data.EnableDhcpv6 = types.BoolNull()
	data.EnableDhcpv6Variable = types.StringNull()
	if t := res.Get(path + "dhcpClientV6.optionType"); t.Exists() {
		va := res.Get(path + "dhcpClientV6.value")
		if t.String() == "variable" {
			data.EnableDhcpv6Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EnableDhcpv6 = types.BoolValue(va.Bool())
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
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *ServiceLANVPNInterfaceSVI) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.ServiceLanVpnFeatureId.IsNull() {
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
	if !data.InterfaceDescription.IsNull() {
		return false
	}
	if !data.InterfaceDescriptionVariable.IsNull() {
		return false
	}
	if !data.InterfaceMtu.IsNull() {
		return false
	}
	if !data.InterfaceMtuVariable.IsNull() {
		return false
	}
	if !data.IpMtu.IsNull() {
		return false
	}
	if !data.IpMtuVariable.IsNull() {
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
	if !data.Ipv4DhcpHelpers.IsNull() {
		return false
	}
	if !data.Ipv4DhcpHelpersVariable.IsNull() {
		return false
	}
	if !data.Ipv6Address.IsNull() {
		return false
	}
	if !data.Ipv6AddressVariable.IsNull() {
		return false
	}
	if len(data.Ipv6SecondaryAddresses) > 0 {
		return false
	}
	if len(data.Ipv6DhcpHelpers) > 0 {
		return false
	}
	if len(data.Arps) > 0 {
		return false
	}
	if len(data.Ipv4Vrrps) > 0 {
		return false
	}
	if len(data.Ipv6Vrrps) > 0 {
		return false
	}
	if !data.EnableDhcpv6.IsNull() {
		return false
	}
	if !data.EnableDhcpv6Variable.IsNull() {
		return false
	}
	if !data.TcpMss.IsNull() {
		return false
	}
	if !data.TcpMssVariable.IsNull() {
		return false
	}
	if !data.ArpTimeout.IsNull() {
		return false
	}
	if !data.ArpTimeoutVariable.IsNull() {
		return false
	}
	if !data.IpDirectedBroadcast.IsNull() {
		return false
	}
	if !data.IpDirectedBroadcastVariable.IsNull() {
		return false
	}
	if !data.IcmpRedirectDisable.IsNull() {
		return false
	}
	if !data.IcmpRedirectDisableVariable.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
