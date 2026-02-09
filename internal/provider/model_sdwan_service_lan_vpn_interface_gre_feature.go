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
type ServiceLANVPNInterfaceGRE struct {
	Id                                    types.String `tfsdk:"id"`
	Version                               types.Int64  `tfsdk:"version"`
	Name                                  types.String `tfsdk:"name"`
	Description                           types.String `tfsdk:"description"`
	FeatureProfileId                      types.String `tfsdk:"feature_profile_id"`
	ServiceLanVpnFeatureId                types.String `tfsdk:"service_lan_vpn_feature_id"`
	InterfaceName                         types.String `tfsdk:"interface_name"`
	InterfaceNameVariable                 types.String `tfsdk:"interface_name_variable"`
	InterfaceDescription                  types.String `tfsdk:"interface_description"`
	InterfaceDescriptionVariable          types.String `tfsdk:"interface_description_variable"`
	Ipv4Address                           types.String `tfsdk:"ipv4_address"`
	Ipv4AddressVariable                   types.String `tfsdk:"ipv4_address_variable"`
	Ipv4SubnetMask                        types.String `tfsdk:"ipv4_subnet_mask"`
	Ipv4SubnetMaskVariable                types.String `tfsdk:"ipv4_subnet_mask_variable"`
	Ipv6Address                           types.String `tfsdk:"ipv6_address"`
	Ipv6AddressVariable                   types.String `tfsdk:"ipv6_address_variable"`
	Shutdown                              types.Bool   `tfsdk:"shutdown"`
	ShutdownVariable                      types.String `tfsdk:"shutdown_variable"`
	TunnelProtection                      types.Bool   `tfsdk:"tunnel_protection"`
	TunnelMode                            types.String `tfsdk:"tunnel_mode"`
	TunnelSourceIpv4Address               types.String `tfsdk:"tunnel_source_ipv4_address"`
	TunnelSourceIpv4AddressVariable       types.String `tfsdk:"tunnel_source_ipv4_address_variable"`
	TunnelRouteViaIpv4Address             types.String `tfsdk:"tunnel_route_via_ipv4_address"`
	TunnelRouteViaIpv4AddressVariable     types.String `tfsdk:"tunnel_route_via_ipv4_address_variable"`
	TunnelSourceInterface                 types.String `tfsdk:"tunnel_source_interface"`
	TunnelSourceInterfaceVariable         types.String `tfsdk:"tunnel_source_interface_variable"`
	TunnelRouteViaInterface               types.String `tfsdk:"tunnel_route_via_interface"`
	TunnelRouteViaInterfaceVariable       types.String `tfsdk:"tunnel_route_via_interface_variable"`
	TunnelSourceInterfaceLoopback         types.String `tfsdk:"tunnel_source_interface_loopback"`
	TunnelSourceInterfaceLoopbackVariable types.String `tfsdk:"tunnel_source_interface_loopback_variable"`
	TunnelRouteViaLoopback                types.String `tfsdk:"tunnel_route_via_loopback"`
	TunnelRouteViaLoopbackVariable        types.String `tfsdk:"tunnel_route_via_loopback_variable"`
	TunnelSourceIpv6Address               types.String `tfsdk:"tunnel_source_ipv6_address"`
	TunnelSourceIpv6AddressVariable       types.String `tfsdk:"tunnel_source_ipv6_address_variable"`
	TunnelRouteViaIpv6Address             types.String `tfsdk:"tunnel_route_via_ipv6_address"`
	TunnelRouteViaIpv6AddressVariable     types.String `tfsdk:"tunnel_route_via_ipv6_address_variable"`
	TunnelDestinationIpv4Address          types.String `tfsdk:"tunnel_destination_ipv4_address"`
	TunnelDestinationIpv4AddressVariable  types.String `tfsdk:"tunnel_destination_ipv4_address_variable"`
	TunnelDestinationIpv6Address          types.String `tfsdk:"tunnel_destination_ipv6_address"`
	TunnelDestinationIpv6AddressVariable  types.String `tfsdk:"tunnel_destination_ipv6_address_variable"`
	Ipv4Mtu                               types.Int64  `tfsdk:"ipv4_mtu"`
	Ipv4MtuVariable                       types.String `tfsdk:"ipv4_mtu_variable"`
	Ipv6Mtu                               types.Int64  `tfsdk:"ipv6_mtu"`
	Ipv6MtuVariable                       types.String `tfsdk:"ipv6_mtu_variable"`
	Ipv4TcpMss                            types.Int64  `tfsdk:"ipv4_tcp_mss"`
	Ipv4TcpMssVariable                    types.String `tfsdk:"ipv4_tcp_mss_variable"`
	Ipv6TcpMss                            types.Int64  `tfsdk:"ipv6_tcp_mss"`
	Ipv6TcpMssVariable                    types.String `tfsdk:"ipv6_tcp_mss_variable"`
	ClearDontFragment                     types.Bool   `tfsdk:"clear_dont_fragment"`
	ClearDontFragmentVariable             types.String `tfsdk:"clear_dont_fragment_variable"`
	DpdInterval                           types.Int64  `tfsdk:"dpd_interval"`
	DpdIntervalVariable                   types.String `tfsdk:"dpd_interval_variable"`
	DpdRetries                            types.Int64  `tfsdk:"dpd_retries"`
	DpdRetriesVariable                    types.String `tfsdk:"dpd_retries_variable"`
	IkeVersion                            types.Int64  `tfsdk:"ike_version"`
	IkeMode                               types.String `tfsdk:"ike_mode"`
	IkeModeVariable                       types.String `tfsdk:"ike_mode_variable"`
	IkeRekeyInterval                      types.Int64  `tfsdk:"ike_rekey_interval"`
	IkeRekeyIntervalVariable              types.String `tfsdk:"ike_rekey_interval_variable"`
	IkeCiphersuite                        types.String `tfsdk:"ike_ciphersuite"`
	IkeCiphersuiteVariable                types.String `tfsdk:"ike_ciphersuite_variable"`
	IkeGroup                              types.String `tfsdk:"ike_group"`
	IkeGroupVariable                      types.String `tfsdk:"ike_group_variable"`
	PreSharedSecret                       types.String `tfsdk:"pre_shared_secret"`
	PreSharedSecretVariable               types.String `tfsdk:"pre_shared_secret_variable"`
	IkeLocalId                            types.String `tfsdk:"ike_local_id"`
	IkeLocalIdVariable                    types.String `tfsdk:"ike_local_id_variable"`
	IkeRemoteId                           types.String `tfsdk:"ike_remote_id"`
	IkeRemoteIdVariable                   types.String `tfsdk:"ike_remote_id_variable"`
	IpsecRekeyInterval                    types.Int64  `tfsdk:"ipsec_rekey_interval"`
	IpsecRekeyIntervalVariable            types.String `tfsdk:"ipsec_rekey_interval_variable"`
	IpsecReplayWindow                     types.Int64  `tfsdk:"ipsec_replay_window"`
	IpsecReplayWindowVariable             types.String `tfsdk:"ipsec_replay_window_variable"`
	IpsecCiphersuite                      types.String `tfsdk:"ipsec_ciphersuite"`
	IpsecCiphersuiteVariable              types.String `tfsdk:"ipsec_ciphersuite_variable"`
	PerfectForwardSecrecy                 types.String `tfsdk:"perfect_forward_secrecy"`
	PerfectForwardSecrecyVariable         types.String `tfsdk:"perfect_forward_secrecy_variable"`
	ApplicationTunnelType                 types.String `tfsdk:"application_tunnel_type"`
	ApplicationTunnelTypeVariable         types.String `tfsdk:"application_tunnel_type_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceLANVPNInterfaceGRE) getModel() string {
	return "service_lan_vpn_interface_gre"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceLANVPNInterfaceGRE) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/lan/vpn/%s/interface/gre", url.QueryEscape(data.FeatureProfileId.ValueString()), url.QueryEscape(data.ServiceLanVpnFeatureId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceLANVPNInterfaceGRE) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.InterfaceNameVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ifName.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.ifName.value", data.InterfaceNameVariable.ValueString())
		}
	} else if !data.InterfaceName.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ifName.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.ifName.value", data.InterfaceName.ValueString())
		}
	}

	if !data.InterfaceDescriptionVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.description.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.description.value", data.InterfaceDescriptionVariable.ValueString())
		}
	} else if data.InterfaceDescription.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.description.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.description.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.description.value", data.InterfaceDescription.ValueString())
		}
	}

	if !data.Ipv4AddressVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.address.address.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.address.address.value", data.Ipv4AddressVariable.ValueString())
		}
	} else if data.Ipv4Address.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.address.address.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.address.address.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.address.address.value", data.Ipv4Address.ValueString())
		}
	}

	if !data.Ipv4SubnetMaskVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.address.mask.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.address.mask.value", data.Ipv4SubnetMaskVariable.ValueString())
		}
	} else if data.Ipv4SubnetMask.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.address.mask.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.address.mask.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.address.mask.value", data.Ipv4SubnetMask.ValueString())
		}
	}

	if !data.Ipv6AddressVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ipv6Address.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.ipv6Address.value", data.Ipv6AddressVariable.ValueString())
		}
	} else if data.Ipv6Address.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ipv6Address.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.ipv6Address.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.ipv6Address.value", data.Ipv6Address.ValueString())
		}
	}

	if !data.ShutdownVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.shutdown.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.shutdown.value", data.ShutdownVariable.ValueString())
		}
	} else if data.Shutdown.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.shutdown.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.shutdown.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.shutdown.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.shutdown.value", data.Shutdown.ValueBool())
		}
	}
	if data.TunnelProtection.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelProtection.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.tunnelProtection.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelProtection.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.tunnelProtection.value", data.TunnelProtection.ValueBool())
		}
	}
	if data.TunnelMode.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelMode.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.tunnelMode.value", "ipv4")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelMode.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.tunnelMode.value", data.TunnelMode.ValueString())
		}
	}

	if !data.TunnelSourceIpv4AddressVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIp.tunnelSource.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIp.tunnelSource.value", data.TunnelSourceIpv4AddressVariable.ValueString())
		}
	} else if !data.TunnelSourceIpv4Address.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIp.tunnelSource.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIp.tunnelSource.value", data.TunnelSourceIpv4Address.ValueString())
		}
	}

	if !data.TunnelRouteViaIpv4AddressVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIp.tunnelRouteVia.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIp.tunnelRouteVia.value", data.TunnelRouteViaIpv4AddressVariable.ValueString())
		}
	} else if !data.TunnelRouteViaIpv4Address.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIp.tunnelRouteVia.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIp.tunnelRouteVia.value", data.TunnelRouteViaIpv4Address.ValueString())
		}
	}

	if !data.TunnelSourceInterfaceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceNotLoopback.tunnelSourceInterface.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceNotLoopback.tunnelSourceInterface.value", data.TunnelSourceInterfaceVariable.ValueString())
		}
	} else if !data.TunnelSourceInterface.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceNotLoopback.tunnelSourceInterface.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceNotLoopback.tunnelSourceInterface.value", data.TunnelSourceInterface.ValueString())
		}
	}

	if !data.TunnelRouteViaInterfaceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceNotLoopback.tunnelRouteVia.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceNotLoopback.tunnelRouteVia.value", data.TunnelRouteViaInterfaceVariable.ValueString())
		}
	} else if !data.TunnelRouteViaInterface.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceNotLoopback.tunnelRouteVia.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceNotLoopback.tunnelRouteVia.value", data.TunnelRouteViaInterface.ValueString())
		}
	}

	if !data.TunnelSourceInterfaceLoopbackVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceLoopback.tunnelSourceInterface.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceLoopback.tunnelSourceInterface.value", data.TunnelSourceInterfaceLoopbackVariable.ValueString())
		}
	} else if !data.TunnelSourceInterfaceLoopback.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceLoopback.tunnelSourceInterface.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceLoopback.tunnelSourceInterface.value", data.TunnelSourceInterfaceLoopback.ValueString())
		}
	}

	if !data.TunnelRouteViaLoopbackVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceLoopback.tunnelRouteVia.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceLoopback.tunnelRouteVia.value", data.TunnelRouteViaLoopbackVariable.ValueString())
		}
	} else if !data.TunnelRouteViaLoopback.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceLoopback.tunnelRouteVia.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceLoopback.tunnelRouteVia.value", data.TunnelRouteViaLoopback.ValueString())
		}
	}

	if !data.TunnelSourceIpv6AddressVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIpv6.tunnelSourceV6.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIpv6.tunnelSourceV6.value", data.TunnelSourceIpv6AddressVariable.ValueString())
		}
	} else if !data.TunnelSourceIpv6Address.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIpv6.tunnelSourceV6.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIpv6.tunnelSourceV6.value", data.TunnelSourceIpv6Address.ValueString())
		}
	}

	if !data.TunnelRouteViaIpv6AddressVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIpv6.tunnelRouteVia.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIpv6.tunnelRouteVia.value", data.TunnelRouteViaIpv6AddressVariable.ValueString())
		}
	} else if !data.TunnelRouteViaIpv6Address.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIpv6.tunnelRouteVia.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIpv6.tunnelRouteVia.value", data.TunnelRouteViaIpv6Address.ValueString())
		}
	}

	if !data.TunnelDestinationIpv4AddressVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelDestination.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.tunnelDestination.value", data.TunnelDestinationIpv4AddressVariable.ValueString())
		}
	} else if !data.TunnelDestinationIpv4Address.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelDestination.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.tunnelDestination.value", data.TunnelDestinationIpv4Address.ValueString())
		}
	}

	if !data.TunnelDestinationIpv6AddressVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelDestinationV6.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.tunnelDestinationV6.value", data.TunnelDestinationIpv6AddressVariable.ValueString())
		}
	} else if !data.TunnelDestinationIpv6Address.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tunnelDestinationV6.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.tunnelDestinationV6.value", data.TunnelDestinationIpv6Address.ValueString())
		}
	}

	if !data.Ipv4MtuVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.mtu.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.mtu.value", data.Ipv4MtuVariable.ValueString())
		}
	} else if data.Ipv4Mtu.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.mtu.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.mtu.value", 1500)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.mtu.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.mtu.value", data.Ipv4Mtu.ValueInt64())
		}
	}

	if !data.Ipv6MtuVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.mtuV6.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.mtuV6.value", data.Ipv6MtuVariable.ValueString())
		}
	} else if data.Ipv6Mtu.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.mtuV6.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.mtuV6.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.mtuV6.value", data.Ipv6Mtu.ValueInt64())
		}
	}

	if !data.Ipv4TcpMssVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tcpMssAdjust.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.tcpMssAdjust.value", data.Ipv4TcpMssVariable.ValueString())
		}
	} else if data.Ipv4TcpMss.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tcpMssAdjust.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.tcpMssAdjust.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.tcpMssAdjust.value", data.Ipv4TcpMss.ValueInt64())
		}
	}

	if !data.Ipv6TcpMssVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tcpMssAdjustV6.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.tcpMssAdjustV6.value", data.Ipv6TcpMssVariable.ValueString())
		}
	} else if data.Ipv6TcpMss.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.tcpMssAdjustV6.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.tcpMssAdjustV6.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.tcpMssAdjustV6.value", data.Ipv6TcpMss.ValueInt64())
		}
	}

	if !data.ClearDontFragmentVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.clearDontFragment.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.clearDontFragment.value", data.ClearDontFragmentVariable.ValueString())
		}
	} else if data.ClearDontFragment.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.clearDontFragment.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.clearDontFragment.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.clearDontFragment.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.clearDontFragment.value", data.ClearDontFragment.ValueBool())
		}
	}

	if !data.DpdIntervalVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.dpdInterval.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.dpdInterval.value", data.DpdIntervalVariable.ValueString())
		}
	} else if data.DpdInterval.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.dpdInterval.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.dpdInterval.value", 10)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.dpdInterval.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.dpdInterval.value", data.DpdInterval.ValueInt64())
		}
	}

	if !data.DpdRetriesVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.dpdRetries.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.dpdRetries.value", data.DpdRetriesVariable.ValueString())
		}
	} else if data.DpdRetries.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.dpdRetries.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.dpdRetries.value", 3)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.dpdRetries.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.dpdRetries.value", data.DpdRetries.ValueInt64())
		}
	}
	if data.IkeVersion.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeVersion.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.ikeVersion.value", 1)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeVersion.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.ikeVersion.value", data.IkeVersion.ValueInt64())
		}
	}

	if !data.IkeModeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeMode.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.ikeMode.value", data.IkeModeVariable.ValueString())
		}
	} else if data.IkeMode.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeMode.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.ikeMode.value", "main")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeMode.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.ikeMode.value", data.IkeMode.ValueString())
		}
	}

	if !data.IkeRekeyIntervalVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeRekeyInterval.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.ikeRekeyInterval.value", data.IkeRekeyIntervalVariable.ValueString())
		}
	} else if data.IkeRekeyInterval.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeRekeyInterval.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.ikeRekeyInterval.value", 14400)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeRekeyInterval.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.ikeRekeyInterval.value", data.IkeRekeyInterval.ValueInt64())
		}
	}

	if !data.IkeCiphersuiteVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeCiphersuite.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.ikeCiphersuite.value", data.IkeCiphersuiteVariable.ValueString())
		}
	} else if data.IkeCiphersuite.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeCiphersuite.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.ikeCiphersuite.value", "aes256-cbc-sha1")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeCiphersuite.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.ikeCiphersuite.value", data.IkeCiphersuite.ValueString())
		}
	}

	if !data.IkeGroupVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeGroup.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.ikeGroup.value", data.IkeGroupVariable.ValueString())
		}
	} else if data.IkeGroup.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeGroup.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.ikeGroup.value", "16")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeGroup.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.ikeGroup.value", data.IkeGroup.ValueString())
		}
	}

	if !data.PreSharedSecretVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.preSharedSecret.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.preSharedSecret.value", data.PreSharedSecretVariable.ValueString())
		}
	} else if data.PreSharedSecret.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.preSharedSecret.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.preSharedSecret.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.preSharedSecret.value", data.PreSharedSecret.ValueString())
		}
	}

	if !data.IkeLocalIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeLocalId.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.ikeLocalId.value", data.IkeLocalIdVariable.ValueString())
		}
	} else if data.IkeLocalId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeLocalId.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeLocalId.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.ikeLocalId.value", data.IkeLocalId.ValueString())
		}
	}

	if !data.IkeRemoteIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeRemoteId.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.ikeRemoteId.value", data.IkeRemoteIdVariable.ValueString())
		}
	} else if data.IkeRemoteId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeRemoteId.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.ikeRemoteId.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.ikeRemoteId.value", data.IkeRemoteId.ValueString())
		}
	}

	if !data.IpsecRekeyIntervalVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ipsecRekeyInterval.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.ipsecRekeyInterval.value", data.IpsecRekeyIntervalVariable.ValueString())
		}
	} else if data.IpsecRekeyInterval.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ipsecRekeyInterval.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.ipsecRekeyInterval.value", 3600)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.ipsecRekeyInterval.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.ipsecRekeyInterval.value", data.IpsecRekeyInterval.ValueInt64())
		}
	}

	if !data.IpsecReplayWindowVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ipsecReplayWindow.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.ipsecReplayWindow.value", data.IpsecReplayWindowVariable.ValueString())
		}
	} else if data.IpsecReplayWindow.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ipsecReplayWindow.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.ipsecReplayWindow.value", 512)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.ipsecReplayWindow.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.ipsecReplayWindow.value", data.IpsecReplayWindow.ValueInt64())
		}
	}

	if !data.IpsecCiphersuiteVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ipsecCiphersuite.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.ipsecCiphersuite.value", data.IpsecCiphersuiteVariable.ValueString())
		}
	} else if data.IpsecCiphersuite.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.ipsecCiphersuite.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.ipsecCiphersuite.value", "aes256-gcm")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.ipsecCiphersuite.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.ipsecCiphersuite.value", data.IpsecCiphersuite.ValueString())
		}
	}

	if !data.PerfectForwardSecrecyVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.perfectForwardSecrecy.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.perfectForwardSecrecy.value", data.PerfectForwardSecrecyVariable.ValueString())
		}
	} else if data.PerfectForwardSecrecy.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.perfectForwardSecrecy.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.perfectForwardSecrecy.value", "group-16")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.perfectForwardSecrecy.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.perfectForwardSecrecy.value", data.PerfectForwardSecrecy.ValueString())
		}
	}

	if !data.ApplicationTunnelTypeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.application.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.application.value", data.ApplicationTunnelTypeVariable.ValueString())
		}
	} else if !data.ApplicationTunnelType.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.application.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.application.value", data.ApplicationTunnelType.ValueString())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceLANVPNInterfaceGRE) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.InterfaceName = types.StringNull()
	data.InterfaceNameVariable = types.StringNull()
	if t := res.Get(path + "basic.ifName.optionType"); t.Exists() {
		va := res.Get(path + "basic.ifName.value")
		if t.String() == "variable" {
			data.InterfaceNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceName = types.StringValue(va.String())
		}
	}
	data.InterfaceDescription = types.StringNull()
	data.InterfaceDescriptionVariable = types.StringNull()
	if t := res.Get(path + "basic.description.optionType"); t.Exists() {
		va := res.Get(path + "basic.description.value")
		if t.String() == "variable" {
			data.InterfaceDescriptionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceDescription = types.StringValue(va.String())
		}
	}
	data.Ipv4Address = types.StringNull()
	data.Ipv4AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.address.address.optionType"); t.Exists() {
		va := res.Get(path + "basic.address.address.value")
		if t.String() == "variable" {
			data.Ipv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Address = types.StringValue(va.String())
		}
	}
	data.Ipv4SubnetMask = types.StringNull()
	data.Ipv4SubnetMaskVariable = types.StringNull()
	if t := res.Get(path + "basic.address.mask.optionType"); t.Exists() {
		va := res.Get(path + "basic.address.mask.value")
		if t.String() == "variable" {
			data.Ipv4SubnetMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4SubnetMask = types.StringValue(va.String())
		}
	}
	data.Ipv6Address = types.StringNull()
	data.Ipv6AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.ipv6Address.optionType"); t.Exists() {
		va := res.Get(path + "basic.ipv6Address.value")
		if t.String() == "variable" {
			data.Ipv6AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6Address = types.StringValue(va.String())
		}
	}
	data.Shutdown = types.BoolNull()
	data.ShutdownVariable = types.StringNull()
	if t := res.Get(path + "basic.shutdown.optionType"); t.Exists() {
		va := res.Get(path + "basic.shutdown.value")
		if t.String() == "variable" {
			data.ShutdownVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Shutdown = types.BoolValue(va.Bool())
		}
	}
	data.TunnelProtection = types.BoolNull()

	if t := res.Get(path + "basic.tunnelProtection.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelProtection.value")
		if t.String() == "global" {
			data.TunnelProtection = types.BoolValue(va.Bool())
		}
	}
	data.TunnelMode = types.StringNull()

	if t := res.Get(path + "basic.tunnelMode.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelMode.value")
		if t.String() == "global" {
			data.TunnelMode = types.StringValue(va.String())
		}
	}
	data.TunnelSourceIpv4Address = types.StringNull()
	data.TunnelSourceIpv4AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceIp.tunnelSource.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceIp.tunnelSource.value")
		if t.String() == "variable" {
			data.TunnelSourceIpv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceIpv4Address = types.StringValue(va.String())
		}
	}
	data.TunnelRouteViaIpv4Address = types.StringNull()
	data.TunnelRouteViaIpv4AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceIp.tunnelRouteVia.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceIp.tunnelRouteVia.value")
		if t.String() == "variable" {
			data.TunnelRouteViaIpv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelRouteViaIpv4Address = types.StringValue(va.String())
		}
	}
	data.TunnelSourceInterface = types.StringNull()
	data.TunnelSourceInterfaceVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceNotLoopback.tunnelSourceInterface.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceNotLoopback.tunnelSourceInterface.value")
		if t.String() == "variable" {
			data.TunnelSourceInterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceInterface = types.StringValue(va.String())
		}
	}
	data.TunnelRouteViaInterface = types.StringNull()
	data.TunnelRouteViaInterfaceVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceNotLoopback.tunnelRouteVia.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceNotLoopback.tunnelRouteVia.value")
		if t.String() == "variable" {
			data.TunnelRouteViaInterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelRouteViaInterface = types.StringValue(va.String())
		}
	}
	data.TunnelSourceInterfaceLoopback = types.StringNull()
	data.TunnelSourceInterfaceLoopbackVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceLoopback.tunnelSourceInterface.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceLoopback.tunnelSourceInterface.value")
		if t.String() == "variable" {
			data.TunnelSourceInterfaceLoopbackVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceInterfaceLoopback = types.StringValue(va.String())
		}
	}
	data.TunnelRouteViaLoopback = types.StringNull()
	data.TunnelRouteViaLoopbackVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceLoopback.tunnelRouteVia.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceLoopback.tunnelRouteVia.value")
		if t.String() == "variable" {
			data.TunnelRouteViaLoopbackVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelRouteViaLoopback = types.StringValue(va.String())
		}
	}
	data.TunnelSourceIpv6Address = types.StringNull()
	data.TunnelSourceIpv6AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceIpv6.tunnelSourceV6.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceIpv6.tunnelSourceV6.value")
		if t.String() == "variable" {
			data.TunnelSourceIpv6AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceIpv6Address = types.StringValue(va.String())
		}
	}
	data.TunnelRouteViaIpv6Address = types.StringNull()
	data.TunnelRouteViaIpv6AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceIpv6.tunnelRouteVia.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceIpv6.tunnelRouteVia.value")
		if t.String() == "variable" {
			data.TunnelRouteViaIpv6AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelRouteViaIpv6Address = types.StringValue(va.String())
		}
	}
	data.TunnelDestinationIpv4Address = types.StringNull()
	data.TunnelDestinationIpv4AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelDestination.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelDestination.value")
		if t.String() == "variable" {
			data.TunnelDestinationIpv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelDestinationIpv4Address = types.StringValue(va.String())
		}
	}
	data.TunnelDestinationIpv6Address = types.StringNull()
	data.TunnelDestinationIpv6AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelDestinationV6.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelDestinationV6.value")
		if t.String() == "variable" {
			data.TunnelDestinationIpv6AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelDestinationIpv6Address = types.StringValue(va.String())
		}
	}
	data.Ipv4Mtu = types.Int64Null()
	data.Ipv4MtuVariable = types.StringNull()
	if t := res.Get(path + "basic.mtu.optionType"); t.Exists() {
		va := res.Get(path + "basic.mtu.value")
		if t.String() == "variable" {
			data.Ipv4MtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Mtu = types.Int64Value(va.Int())
		}
	}
	data.Ipv6Mtu = types.Int64Null()
	data.Ipv6MtuVariable = types.StringNull()
	if t := res.Get(path + "basic.mtuV6.optionType"); t.Exists() {
		va := res.Get(path + "basic.mtuV6.value")
		if t.String() == "variable" {
			data.Ipv6MtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6Mtu = types.Int64Value(va.Int())
		}
	}
	data.Ipv4TcpMss = types.Int64Null()
	data.Ipv4TcpMssVariable = types.StringNull()
	if t := res.Get(path + "basic.tcpMssAdjust.optionType"); t.Exists() {
		va := res.Get(path + "basic.tcpMssAdjust.value")
		if t.String() == "variable" {
			data.Ipv4TcpMssVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4TcpMss = types.Int64Value(va.Int())
		}
	}
	data.Ipv6TcpMss = types.Int64Null()
	data.Ipv6TcpMssVariable = types.StringNull()
	if t := res.Get(path + "basic.tcpMssAdjustV6.optionType"); t.Exists() {
		va := res.Get(path + "basic.tcpMssAdjustV6.value")
		if t.String() == "variable" {
			data.Ipv6TcpMssVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6TcpMss = types.Int64Value(va.Int())
		}
	}
	data.ClearDontFragment = types.BoolNull()
	data.ClearDontFragmentVariable = types.StringNull()
	if t := res.Get(path + "basic.clearDontFragment.optionType"); t.Exists() {
		va := res.Get(path + "basic.clearDontFragment.value")
		if t.String() == "variable" {
			data.ClearDontFragmentVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ClearDontFragment = types.BoolValue(va.Bool())
		}
	}
	data.DpdInterval = types.Int64Null()
	data.DpdIntervalVariable = types.StringNull()
	if t := res.Get(path + "basic.dpdInterval.optionType"); t.Exists() {
		va := res.Get(path + "basic.dpdInterval.value")
		if t.String() == "variable" {
			data.DpdIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DpdInterval = types.Int64Value(va.Int())
		}
	}
	data.DpdRetries = types.Int64Null()
	data.DpdRetriesVariable = types.StringNull()
	if t := res.Get(path + "basic.dpdRetries.optionType"); t.Exists() {
		va := res.Get(path + "basic.dpdRetries.value")
		if t.String() == "variable" {
			data.DpdRetriesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DpdRetries = types.Int64Value(va.Int())
		}
	}
	data.IkeVersion = types.Int64Null()

	if t := res.Get(path + "basic.ikeVersion.optionType"); t.Exists() {
		va := res.Get(path + "basic.ikeVersion.value")
		if t.String() == "global" {
			data.IkeVersion = types.Int64Value(va.Int())
		}
	}
	data.IkeMode = types.StringNull()
	data.IkeModeVariable = types.StringNull()
	if t := res.Get(path + "basic.ikeMode.optionType"); t.Exists() {
		va := res.Get(path + "basic.ikeMode.value")
		if t.String() == "variable" {
			data.IkeModeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeMode = types.StringValue(va.String())
		}
	}
	data.IkeRekeyInterval = types.Int64Null()
	data.IkeRekeyIntervalVariable = types.StringNull()
	if t := res.Get(path + "basic.ikeRekeyInterval.optionType"); t.Exists() {
		va := res.Get(path + "basic.ikeRekeyInterval.value")
		if t.String() == "variable" {
			data.IkeRekeyIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeRekeyInterval = types.Int64Value(va.Int())
		}
	}
	data.IkeCiphersuite = types.StringNull()
	data.IkeCiphersuiteVariable = types.StringNull()
	if t := res.Get(path + "basic.ikeCiphersuite.optionType"); t.Exists() {
		va := res.Get(path + "basic.ikeCiphersuite.value")
		if t.String() == "variable" {
			data.IkeCiphersuiteVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeCiphersuite = types.StringValue(va.String())
		}
	}
	data.IkeGroup = types.StringNull()
	data.IkeGroupVariable = types.StringNull()
	if t := res.Get(path + "basic.ikeGroup.optionType"); t.Exists() {
		va := res.Get(path + "basic.ikeGroup.value")
		if t.String() == "variable" {
			data.IkeGroupVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeGroup = types.StringValue(va.String())
		}
	}
	data.PreSharedSecret = types.StringNull()
	data.PreSharedSecretVariable = types.StringNull()
	if t := res.Get(path + "basic.preSharedSecret.optionType"); t.Exists() {
		va := res.Get(path + "basic.preSharedSecret.value")
		if t.String() == "variable" {
			data.PreSharedSecretVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PreSharedSecret = types.StringValue(va.String())
		}
	}
	data.IkeLocalId = types.StringNull()
	data.IkeLocalIdVariable = types.StringNull()
	if t := res.Get(path + "basic.ikeLocalId.optionType"); t.Exists() {
		va := res.Get(path + "basic.ikeLocalId.value")
		if t.String() == "variable" {
			data.IkeLocalIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeLocalId = types.StringValue(va.String())
		}
	}
	data.IkeRemoteId = types.StringNull()
	data.IkeRemoteIdVariable = types.StringNull()
	if t := res.Get(path + "basic.ikeRemoteId.optionType"); t.Exists() {
		va := res.Get(path + "basic.ikeRemoteId.value")
		if t.String() == "variable" {
			data.IkeRemoteIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeRemoteId = types.StringValue(va.String())
		}
	}
	data.IpsecRekeyInterval = types.Int64Null()
	data.IpsecRekeyIntervalVariable = types.StringNull()
	if t := res.Get(path + "basic.ipsecRekeyInterval.optionType"); t.Exists() {
		va := res.Get(path + "basic.ipsecRekeyInterval.value")
		if t.String() == "variable" {
			data.IpsecRekeyIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecRekeyInterval = types.Int64Value(va.Int())
		}
	}
	data.IpsecReplayWindow = types.Int64Null()
	data.IpsecReplayWindowVariable = types.StringNull()
	if t := res.Get(path + "basic.ipsecReplayWindow.optionType"); t.Exists() {
		va := res.Get(path + "basic.ipsecReplayWindow.value")
		if t.String() == "variable" {
			data.IpsecReplayWindowVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecReplayWindow = types.Int64Value(va.Int())
		}
	}
	data.IpsecCiphersuite = types.StringNull()
	data.IpsecCiphersuiteVariable = types.StringNull()
	if t := res.Get(path + "basic.ipsecCiphersuite.optionType"); t.Exists() {
		va := res.Get(path + "basic.ipsecCiphersuite.value")
		if t.String() == "variable" {
			data.IpsecCiphersuiteVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecCiphersuite = types.StringValue(va.String())
		}
	}
	data.PerfectForwardSecrecy = types.StringNull()
	data.PerfectForwardSecrecyVariable = types.StringNull()
	if t := res.Get(path + "basic.perfectForwardSecrecy.optionType"); t.Exists() {
		va := res.Get(path + "basic.perfectForwardSecrecy.value")
		if t.String() == "variable" {
			data.PerfectForwardSecrecyVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PerfectForwardSecrecy = types.StringValue(va.String())
		}
	}
	data.ApplicationTunnelType = types.StringNull()
	data.ApplicationTunnelTypeVariable = types.StringNull()
	if t := res.Get(path + "advanced.application.optionType"); t.Exists() {
		va := res.Get(path + "advanced.application.value")
		if t.String() == "variable" {
			data.ApplicationTunnelTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ApplicationTunnelType = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ServiceLANVPNInterfaceGRE) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.InterfaceName = types.StringNull()
	data.InterfaceNameVariable = types.StringNull()
	if t := res.Get(path + "basic.ifName.optionType"); t.Exists() {
		va := res.Get(path + "basic.ifName.value")
		if t.String() == "variable" {
			data.InterfaceNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceName = types.StringValue(va.String())
		}
	}
	data.InterfaceDescription = types.StringNull()
	data.InterfaceDescriptionVariable = types.StringNull()
	if t := res.Get(path + "basic.description.optionType"); t.Exists() {
		va := res.Get(path + "basic.description.value")
		if t.String() == "variable" {
			data.InterfaceDescriptionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceDescription = types.StringValue(va.String())
		}
	}
	data.Ipv4Address = types.StringNull()
	data.Ipv4AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.address.address.optionType"); t.Exists() {
		va := res.Get(path + "basic.address.address.value")
		if t.String() == "variable" {
			data.Ipv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Address = types.StringValue(va.String())
		}
	}
	data.Ipv4SubnetMask = types.StringNull()
	data.Ipv4SubnetMaskVariable = types.StringNull()
	if t := res.Get(path + "basic.address.mask.optionType"); t.Exists() {
		va := res.Get(path + "basic.address.mask.value")
		if t.String() == "variable" {
			data.Ipv4SubnetMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4SubnetMask = types.StringValue(va.String())
		}
	}
	data.Ipv6Address = types.StringNull()
	data.Ipv6AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.ipv6Address.optionType"); t.Exists() {
		va := res.Get(path + "basic.ipv6Address.value")
		if t.String() == "variable" {
			data.Ipv6AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6Address = types.StringValue(va.String())
		}
	}
	data.Shutdown = types.BoolNull()
	data.ShutdownVariable = types.StringNull()
	if t := res.Get(path + "basic.shutdown.optionType"); t.Exists() {
		va := res.Get(path + "basic.shutdown.value")
		if t.String() == "variable" {
			data.ShutdownVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Shutdown = types.BoolValue(va.Bool())
		}
	}
	data.TunnelProtection = types.BoolNull()

	if t := res.Get(path + "basic.tunnelProtection.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelProtection.value")
		if t.String() == "global" {
			data.TunnelProtection = types.BoolValue(va.Bool())
		}
	}
	data.TunnelMode = types.StringNull()

	if t := res.Get(path + "basic.tunnelMode.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelMode.value")
		if t.String() == "global" {
			data.TunnelMode = types.StringValue(va.String())
		}
	}
	data.TunnelSourceIpv4Address = types.StringNull()
	data.TunnelSourceIpv4AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceIp.tunnelSource.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceIp.tunnelSource.value")
		if t.String() == "variable" {
			data.TunnelSourceIpv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceIpv4Address = types.StringValue(va.String())
		}
	}
	data.TunnelRouteViaIpv4Address = types.StringNull()
	data.TunnelRouteViaIpv4AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceIp.tunnelRouteVia.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceIp.tunnelRouteVia.value")
		if t.String() == "variable" {
			data.TunnelRouteViaIpv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelRouteViaIpv4Address = types.StringValue(va.String())
		}
	}
	data.TunnelSourceInterface = types.StringNull()
	data.TunnelSourceInterfaceVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceNotLoopback.tunnelSourceInterface.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceNotLoopback.tunnelSourceInterface.value")
		if t.String() == "variable" {
			data.TunnelSourceInterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceInterface = types.StringValue(va.String())
		}
	}
	data.TunnelRouteViaInterface = types.StringNull()
	data.TunnelRouteViaInterfaceVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceNotLoopback.tunnelRouteVia.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceNotLoopback.tunnelRouteVia.value")
		if t.String() == "variable" {
			data.TunnelRouteViaInterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelRouteViaInterface = types.StringValue(va.String())
		}
	}
	data.TunnelSourceInterfaceLoopback = types.StringNull()
	data.TunnelSourceInterfaceLoopbackVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceLoopback.tunnelSourceInterface.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceLoopback.tunnelSourceInterface.value")
		if t.String() == "variable" {
			data.TunnelSourceInterfaceLoopbackVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceInterfaceLoopback = types.StringValue(va.String())
		}
	}
	data.TunnelRouteViaLoopback = types.StringNull()
	data.TunnelRouteViaLoopbackVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceLoopback.tunnelRouteVia.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceLoopback.tunnelRouteVia.value")
		if t.String() == "variable" {
			data.TunnelRouteViaLoopbackVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelRouteViaLoopback = types.StringValue(va.String())
		}
	}
	data.TunnelSourceIpv6Address = types.StringNull()
	data.TunnelSourceIpv6AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceIpv6.tunnelSourceV6.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceIpv6.tunnelSourceV6.value")
		if t.String() == "variable" {
			data.TunnelSourceIpv6AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceIpv6Address = types.StringValue(va.String())
		}
	}
	data.TunnelRouteViaIpv6Address = types.StringNull()
	data.TunnelRouteViaIpv6AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceIpv6.tunnelRouteVia.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceIpv6.tunnelRouteVia.value")
		if t.String() == "variable" {
			data.TunnelRouteViaIpv6AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelRouteViaIpv6Address = types.StringValue(va.String())
		}
	}
	data.TunnelDestinationIpv4Address = types.StringNull()
	data.TunnelDestinationIpv4AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelDestination.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelDestination.value")
		if t.String() == "variable" {
			data.TunnelDestinationIpv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelDestinationIpv4Address = types.StringValue(va.String())
		}
	}
	data.TunnelDestinationIpv6Address = types.StringNull()
	data.TunnelDestinationIpv6AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelDestinationV6.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelDestinationV6.value")
		if t.String() == "variable" {
			data.TunnelDestinationIpv6AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelDestinationIpv6Address = types.StringValue(va.String())
		}
	}
	data.Ipv4Mtu = types.Int64Null()
	data.Ipv4MtuVariable = types.StringNull()
	if t := res.Get(path + "basic.mtu.optionType"); t.Exists() {
		va := res.Get(path + "basic.mtu.value")
		if t.String() == "variable" {
			data.Ipv4MtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Mtu = types.Int64Value(va.Int())
		}
	}
	data.Ipv6Mtu = types.Int64Null()
	data.Ipv6MtuVariable = types.StringNull()
	if t := res.Get(path + "basic.mtuV6.optionType"); t.Exists() {
		va := res.Get(path + "basic.mtuV6.value")
		if t.String() == "variable" {
			data.Ipv6MtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6Mtu = types.Int64Value(va.Int())
		}
	}
	data.Ipv4TcpMss = types.Int64Null()
	data.Ipv4TcpMssVariable = types.StringNull()
	if t := res.Get(path + "basic.tcpMssAdjust.optionType"); t.Exists() {
		va := res.Get(path + "basic.tcpMssAdjust.value")
		if t.String() == "variable" {
			data.Ipv4TcpMssVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4TcpMss = types.Int64Value(va.Int())
		}
	}
	data.Ipv6TcpMss = types.Int64Null()
	data.Ipv6TcpMssVariable = types.StringNull()
	if t := res.Get(path + "basic.tcpMssAdjustV6.optionType"); t.Exists() {
		va := res.Get(path + "basic.tcpMssAdjustV6.value")
		if t.String() == "variable" {
			data.Ipv6TcpMssVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6TcpMss = types.Int64Value(va.Int())
		}
	}
	data.ClearDontFragment = types.BoolNull()
	data.ClearDontFragmentVariable = types.StringNull()
	if t := res.Get(path + "basic.clearDontFragment.optionType"); t.Exists() {
		va := res.Get(path + "basic.clearDontFragment.value")
		if t.String() == "variable" {
			data.ClearDontFragmentVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ClearDontFragment = types.BoolValue(va.Bool())
		}
	}
	data.DpdInterval = types.Int64Null()
	data.DpdIntervalVariable = types.StringNull()
	if t := res.Get(path + "basic.dpdInterval.optionType"); t.Exists() {
		va := res.Get(path + "basic.dpdInterval.value")
		if t.String() == "variable" {
			data.DpdIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DpdInterval = types.Int64Value(va.Int())
		}
	}
	data.DpdRetries = types.Int64Null()
	data.DpdRetriesVariable = types.StringNull()
	if t := res.Get(path + "basic.dpdRetries.optionType"); t.Exists() {
		va := res.Get(path + "basic.dpdRetries.value")
		if t.String() == "variable" {
			data.DpdRetriesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DpdRetries = types.Int64Value(va.Int())
		}
	}
	data.IkeVersion = types.Int64Null()

	if t := res.Get(path + "basic.ikeVersion.optionType"); t.Exists() {
		va := res.Get(path + "basic.ikeVersion.value")
		if t.String() == "global" {
			data.IkeVersion = types.Int64Value(va.Int())
		}
	}
	data.IkeMode = types.StringNull()
	data.IkeModeVariable = types.StringNull()
	if t := res.Get(path + "basic.ikeMode.optionType"); t.Exists() {
		va := res.Get(path + "basic.ikeMode.value")
		if t.String() == "variable" {
			data.IkeModeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeMode = types.StringValue(va.String())
		}
	}
	data.IkeRekeyInterval = types.Int64Null()
	data.IkeRekeyIntervalVariable = types.StringNull()
	if t := res.Get(path + "basic.ikeRekeyInterval.optionType"); t.Exists() {
		va := res.Get(path + "basic.ikeRekeyInterval.value")
		if t.String() == "variable" {
			data.IkeRekeyIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeRekeyInterval = types.Int64Value(va.Int())
		}
	}
	data.IkeCiphersuite = types.StringNull()
	data.IkeCiphersuiteVariable = types.StringNull()
	if t := res.Get(path + "basic.ikeCiphersuite.optionType"); t.Exists() {
		va := res.Get(path + "basic.ikeCiphersuite.value")
		if t.String() == "variable" {
			data.IkeCiphersuiteVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeCiphersuite = types.StringValue(va.String())
		}
	}
	data.IkeGroup = types.StringNull()
	data.IkeGroupVariable = types.StringNull()
	if t := res.Get(path + "basic.ikeGroup.optionType"); t.Exists() {
		va := res.Get(path + "basic.ikeGroup.value")
		if t.String() == "variable" {
			data.IkeGroupVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeGroup = types.StringValue(va.String())
		}
	}
	data.PreSharedSecret = types.StringNull()
	data.PreSharedSecretVariable = types.StringNull()
	if t := res.Get(path + "basic.preSharedSecret.optionType"); t.Exists() {
		va := res.Get(path + "basic.preSharedSecret.value")
		if t.String() == "variable" {
			data.PreSharedSecretVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PreSharedSecret = types.StringValue(va.String())
		}
	}
	data.IkeLocalId = types.StringNull()
	data.IkeLocalIdVariable = types.StringNull()
	if t := res.Get(path + "basic.ikeLocalId.optionType"); t.Exists() {
		va := res.Get(path + "basic.ikeLocalId.value")
		if t.String() == "variable" {
			data.IkeLocalIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeLocalId = types.StringValue(va.String())
		}
	}
	data.IkeRemoteId = types.StringNull()
	data.IkeRemoteIdVariable = types.StringNull()
	if t := res.Get(path + "basic.ikeRemoteId.optionType"); t.Exists() {
		va := res.Get(path + "basic.ikeRemoteId.value")
		if t.String() == "variable" {
			data.IkeRemoteIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeRemoteId = types.StringValue(va.String())
		}
	}
	data.IpsecRekeyInterval = types.Int64Null()
	data.IpsecRekeyIntervalVariable = types.StringNull()
	if t := res.Get(path + "basic.ipsecRekeyInterval.optionType"); t.Exists() {
		va := res.Get(path + "basic.ipsecRekeyInterval.value")
		if t.String() == "variable" {
			data.IpsecRekeyIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecRekeyInterval = types.Int64Value(va.Int())
		}
	}
	data.IpsecReplayWindow = types.Int64Null()
	data.IpsecReplayWindowVariable = types.StringNull()
	if t := res.Get(path + "basic.ipsecReplayWindow.optionType"); t.Exists() {
		va := res.Get(path + "basic.ipsecReplayWindow.value")
		if t.String() == "variable" {
			data.IpsecReplayWindowVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecReplayWindow = types.Int64Value(va.Int())
		}
	}
	data.IpsecCiphersuite = types.StringNull()
	data.IpsecCiphersuiteVariable = types.StringNull()
	if t := res.Get(path + "basic.ipsecCiphersuite.optionType"); t.Exists() {
		va := res.Get(path + "basic.ipsecCiphersuite.value")
		if t.String() == "variable" {
			data.IpsecCiphersuiteVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecCiphersuite = types.StringValue(va.String())
		}
	}
	data.PerfectForwardSecrecy = types.StringNull()
	data.PerfectForwardSecrecyVariable = types.StringNull()
	if t := res.Get(path + "basic.perfectForwardSecrecy.optionType"); t.Exists() {
		va := res.Get(path + "basic.perfectForwardSecrecy.value")
		if t.String() == "variable" {
			data.PerfectForwardSecrecyVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PerfectForwardSecrecy = types.StringValue(va.String())
		}
	}
	data.ApplicationTunnelType = types.StringNull()
	data.ApplicationTunnelTypeVariable = types.StringNull()
	if t := res.Get(path + "advanced.application.optionType"); t.Exists() {
		va := res.Get(path + "advanced.application.value")
		if t.String() == "variable" {
			data.ApplicationTunnelTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ApplicationTunnelType = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end updateFromBody
