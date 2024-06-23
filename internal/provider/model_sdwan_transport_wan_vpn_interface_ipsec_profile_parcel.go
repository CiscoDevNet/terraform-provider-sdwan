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
type TransportWANVPNInterfaceIPSEC struct {
	Id                               types.String `tfsdk:"id"`
	Version                          types.Int64  `tfsdk:"version"`
	Name                             types.String `tfsdk:"name"`
	Description                      types.String `tfsdk:"description"`
	FeatureProfileId                 types.String `tfsdk:"feature_profile_id"`
	TransportWanVpnProfileParcelId   types.String `tfsdk:"transport_wan_vpn_profile_parcel_id"`
	InterfaceName                    types.String `tfsdk:"interface_name"`
	InterfaceNameVariable            types.String `tfsdk:"interface_name_variable"`
	Shutdown                         types.Bool   `tfsdk:"shutdown"`
	ShutdownVariable                 types.String `tfsdk:"shutdown_variable"`
	InterfaceDescription             types.String `tfsdk:"interface_description"`
	InterfaceDescriptionVariable     types.String `tfsdk:"interface_description_variable"`
	InterfaceAddress                 types.String `tfsdk:"interface_address"`
	InterfaceAddressVariable         types.String `tfsdk:"interface_address_variable"`
	Mask                             types.String `tfsdk:"mask"`
	MaskVariable                     types.String `tfsdk:"mask_variable"`
	TunnelSourceAddress              types.String `tfsdk:"tunnel_source_address"`
	TunnelSourceAddressVariable      types.String `tfsdk:"tunnel_source_address_variable"`
	TunnelSourceMask                 types.String `tfsdk:"tunnel_source_mask"`
	TunnelSourceMaskVariable         types.String `tfsdk:"tunnel_source_mask_variable"`
	TunnelSourceInterface            types.String `tfsdk:"tunnel_source_interface"`
	TunnelSourceInterfaceVariable    types.String `tfsdk:"tunnel_source_interface_variable"`
	TunnelDestinationAddress         types.String `tfsdk:"tunnel_destination_address"`
	TunnelDestinationAddressVariable types.String `tfsdk:"tunnel_destination_address_variable"`
	TunnelDestinationMask            types.String `tfsdk:"tunnel_destination_mask"`
	TunnelDestinationMaskVariable    types.String `tfsdk:"tunnel_destination_mask_variable"`
	ApplicationTunnelType            types.String `tfsdk:"application_tunnel_type"`
	ApplicationTunnelTypeVariable    types.String `tfsdk:"application_tunnel_type_variable"`
	TcpMss                           types.Int64  `tfsdk:"tcp_mss"`
	TcpMssVariable                   types.String `tfsdk:"tcp_mss_variable"`
	ClearDontFragment                types.Bool   `tfsdk:"clear_dont_fragment"`
	ClearDontFragmentVariable        types.String `tfsdk:"clear_dont_fragment_variable"`
	IpMtu                            types.Int64  `tfsdk:"ip_mtu"`
	IpMtuVariable                    types.String `tfsdk:"ip_mtu_variable"`
	DpdInterval                      types.Int64  `tfsdk:"dpd_interval"`
	DpdIntervalVariable              types.String `tfsdk:"dpd_interval_variable"`
	DpdRetries                       types.Int64  `tfsdk:"dpd_retries"`
	DpdRetriesVariable               types.String `tfsdk:"dpd_retries_variable"`
	PresharedKeyForIke               types.String `tfsdk:"preshared_key_for_ike"`
	PresharedKeyForIkeVariable       types.String `tfsdk:"preshared_key_for_ike_variable"`
	IkeVersion                       types.Int64  `tfsdk:"ike_version"`
	IkeIntegrityProtocol             types.String `tfsdk:"ike_integrity_protocol"`
	IkeIntegrityProtocolVariable     types.String `tfsdk:"ike_integrity_protocol_variable"`
	IkeRekeyInterval                 types.Int64  `tfsdk:"ike_rekey_interval"`
	IkeRekeyIntervalVariable         types.String `tfsdk:"ike_rekey_interval_variable"`
	IkeCiphersuite                   types.String `tfsdk:"ike_ciphersuite"`
	IkeCiphersuiteVariable           types.String `tfsdk:"ike_ciphersuite_variable"`
	IkeDiffieHellmanGroup            types.String `tfsdk:"ike_diffie_hellman_group"`
	IkeDiffieHellmanGroupVariable    types.String `tfsdk:"ike_diffie_hellman_group_variable"`
	IkeIdForLocalEndPoint            types.String `tfsdk:"ike_id_for_local_end_point"`
	IkeIdForLocalEndPointVariable    types.String `tfsdk:"ike_id_for_local_end_point_variable"`
	IkeIdForRemoteEndPoint           types.String `tfsdk:"ike_id_for_remote_end_point"`
	IkeIdForRemoteEndPointVariable   types.String `tfsdk:"ike_id_for_remote_end_point_variable"`
	IpsecRekeyInterval               types.Int64  `tfsdk:"ipsec_rekey_interval"`
	IpsecRekeyIntervalVariable       types.String `tfsdk:"ipsec_rekey_interval_variable"`
	IpsecReplayWindow                types.Int64  `tfsdk:"ipsec_replay_window"`
	IpsecReplayWindowVariable        types.String `tfsdk:"ipsec_replay_window_variable"`
	IpsecCiphersuite                 types.String `tfsdk:"ipsec_ciphersuite"`
	IpsecCiphersuiteVariable         types.String `tfsdk:"ipsec_ciphersuite_variable"`
	PerfectForwardSecrecy            types.String `tfsdk:"perfect_forward_secrecy"`
	PerfectForwardSecrecyVariable    types.String `tfsdk:"perfect_forward_secrecy_variable"`
	TrackerId                        types.String `tfsdk:"tracker_id"`
	TrackerIdVariable                types.String `tfsdk:"tracker_id_variable"`
	TunnelRouteVia                   types.String `tfsdk:"tunnel_route_via"`
	TunnelRouteViaVariable           types.String `tfsdk:"tunnel_route_via_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TransportWANVPNInterfaceIPSEC) getModel() string {
	return "transport_wan_vpn_interface_ipsec"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportWANVPNInterfaceIPSEC) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/wan/vpn/%s/interface/ipsec", url.QueryEscape(data.FeatureProfileId.ValueString()), url.QueryEscape(data.TransportWanVpnProfileParcelId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportWANVPNInterfaceIPSEC) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.InterfaceNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ifName.optionType", "variable")
		body, _ = sjson.Set(body, path+"ifName.value", data.InterfaceNameVariable.ValueString())
	} else if !data.InterfaceName.IsNull() {
		body, _ = sjson.Set(body, path+"ifName.optionType", "global")
		body, _ = sjson.Set(body, path+"ifName.value", data.InterfaceName.ValueString())
	}

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

	if !data.InterfaceDescriptionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"description.optionType", "variable")
		body, _ = sjson.Set(body, path+"description.value", data.InterfaceDescriptionVariable.ValueString())
	} else if data.InterfaceDescription.IsNull() {
		body, _ = sjson.Set(body, path+"description.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"description.optionType", "global")
		body, _ = sjson.Set(body, path+"description.value", data.InterfaceDescription.ValueString())
	}

	if !data.InterfaceAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"address.address.optionType", "variable")
		body, _ = sjson.Set(body, path+"address.address.value", data.InterfaceAddressVariable.ValueString())
	} else if !data.InterfaceAddress.IsNull() {
		body, _ = sjson.Set(body, path+"address.address.optionType", "global")
		body, _ = sjson.Set(body, path+"address.address.value", data.InterfaceAddress.ValueString())
	}

	if !data.MaskVariable.IsNull() {
		body, _ = sjson.Set(body, path+"address.mask.optionType", "variable")
		body, _ = sjson.Set(body, path+"address.mask.value", data.MaskVariable.ValueString())
	} else if !data.Mask.IsNull() {
		body, _ = sjson.Set(body, path+"address.mask.optionType", "global")
		body, _ = sjson.Set(body, path+"address.mask.value", data.Mask.ValueString())
	}

	if !data.TunnelSourceAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnelSource.address.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnelSource.address.value", data.TunnelSourceAddressVariable.ValueString())
	} else if !data.TunnelSourceAddress.IsNull() {
		body, _ = sjson.Set(body, path+"tunnelSource.address.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnelSource.address.value", data.TunnelSourceAddress.ValueString())
	}

	if !data.TunnelSourceMaskVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnelSource.mask.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnelSource.mask.value", data.TunnelSourceMaskVariable.ValueString())
	} else if !data.TunnelSourceMask.IsNull() {
		body, _ = sjson.Set(body, path+"tunnelSource.mask.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnelSource.mask.value", data.TunnelSourceMask.ValueString())
	}

	if !data.TunnelSourceInterfaceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnelSourceInterface.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnelSourceInterface.value", data.TunnelSourceInterfaceVariable.ValueString())
	} else if !data.TunnelSourceInterface.IsNull() {
		body, _ = sjson.Set(body, path+"tunnelSourceInterface.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnelSourceInterface.value", data.TunnelSourceInterface.ValueString())
	}

	if !data.TunnelDestinationAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnelDestination.address.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnelDestination.address.value", data.TunnelDestinationAddressVariable.ValueString())
	} else if !data.TunnelDestinationAddress.IsNull() {
		body, _ = sjson.Set(body, path+"tunnelDestination.address.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnelDestination.address.value", data.TunnelDestinationAddress.ValueString())
	}

	if !data.TunnelDestinationMaskVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnelDestination.mask.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnelDestination.mask.value", data.TunnelDestinationMaskVariable.ValueString())
	} else if !data.TunnelDestinationMask.IsNull() {
		body, _ = sjson.Set(body, path+"tunnelDestination.mask.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnelDestination.mask.value", data.TunnelDestinationMask.ValueString())
	}

	if !data.ApplicationTunnelTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"application.optionType", "variable")
		body, _ = sjson.Set(body, path+"application.value", data.ApplicationTunnelTypeVariable.ValueString())
	} else if !data.ApplicationTunnelType.IsNull() {
		body, _ = sjson.Set(body, path+"application.optionType", "global")
		body, _ = sjson.Set(body, path+"application.value", data.ApplicationTunnelType.ValueString())
	}

	if !data.TcpMssVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tcpMssAdjust.optionType", "variable")
		body, _ = sjson.Set(body, path+"tcpMssAdjust.value", data.TcpMssVariable.ValueString())
	} else if data.TcpMss.IsNull() {
		body, _ = sjson.Set(body, path+"tcpMssAdjust.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"tcpMssAdjust.optionType", "global")
		body, _ = sjson.Set(body, path+"tcpMssAdjust.value", data.TcpMss.ValueInt64())
	}

	if !data.ClearDontFragmentVariable.IsNull() {
		body, _ = sjson.Set(body, path+"clearDontFragment.optionType", "variable")
		body, _ = sjson.Set(body, path+"clearDontFragment.value", data.ClearDontFragmentVariable.ValueString())
	} else if data.ClearDontFragment.IsNull() {
		body, _ = sjson.Set(body, path+"clearDontFragment.optionType", "default")
		body, _ = sjson.Set(body, path+"clearDontFragment.value", false)
	} else {
		body, _ = sjson.Set(body, path+"clearDontFragment.optionType", "global")
		body, _ = sjson.Set(body, path+"clearDontFragment.value", data.ClearDontFragment.ValueBool())
	}

	if !data.IpMtuVariable.IsNull() {
		body, _ = sjson.Set(body, path+"mtu.optionType", "variable")
		body, _ = sjson.Set(body, path+"mtu.value", data.IpMtuVariable.ValueString())
	} else if data.IpMtu.IsNull() {
		body, _ = sjson.Set(body, path+"mtu.optionType", "default")
		body, _ = sjson.Set(body, path+"mtu.value", 1500)
	} else {
		body, _ = sjson.Set(body, path+"mtu.optionType", "global")
		body, _ = sjson.Set(body, path+"mtu.value", data.IpMtu.ValueInt64())
	}

	if !data.DpdIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"dpdInterval.optionType", "variable")
		body, _ = sjson.Set(body, path+"dpdInterval.value", data.DpdIntervalVariable.ValueString())
	} else if data.DpdInterval.IsNull() {
		body, _ = sjson.Set(body, path+"dpdInterval.optionType", "default")
		body, _ = sjson.Set(body, path+"dpdInterval.value", 10)
	} else {
		body, _ = sjson.Set(body, path+"dpdInterval.optionType", "global")
		body, _ = sjson.Set(body, path+"dpdInterval.value", data.DpdInterval.ValueInt64())
	}

	if !data.DpdRetriesVariable.IsNull() {
		body, _ = sjson.Set(body, path+"dpdRetries.optionType", "variable")
		body, _ = sjson.Set(body, path+"dpdRetries.value", data.DpdRetriesVariable.ValueString())
	} else if data.DpdRetries.IsNull() {
		body, _ = sjson.Set(body, path+"dpdRetries.optionType", "default")
		body, _ = sjson.Set(body, path+"dpdRetries.value", 3)
	} else {
		body, _ = sjson.Set(body, path+"dpdRetries.optionType", "global")
		body, _ = sjson.Set(body, path+"dpdRetries.value", data.DpdRetries.ValueInt64())
	}

	if !data.PresharedKeyForIkeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"preSharedSecret.optionType", "variable")
		body, _ = sjson.Set(body, path+"preSharedSecret.value", data.PresharedKeyForIkeVariable.ValueString())
	} else if !data.PresharedKeyForIke.IsNull() {
		body, _ = sjson.Set(body, path+"preSharedSecret.optionType", "global")
		body, _ = sjson.Set(body, path+"preSharedSecret.value", data.PresharedKeyForIke.ValueString())
	}
	if data.IkeVersion.IsNull() {
		body, _ = sjson.Set(body, path+"ikeVersion.optionType", "default")
		body, _ = sjson.Set(body, path+"ikeVersion.value", 1)
	} else {
		body, _ = sjson.Set(body, path+"ikeVersion.optionType", "global")
		body, _ = sjson.Set(body, path+"ikeVersion.value", data.IkeVersion.ValueInt64())
	}

	if !data.IkeIntegrityProtocolVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ikeMode.optionType", "variable")
		body, _ = sjson.Set(body, path+"ikeMode.value", data.IkeIntegrityProtocolVariable.ValueString())
	} else if data.IkeIntegrityProtocol.IsNull() {
		body, _ = sjson.Set(body, path+"ikeMode.optionType", "default")
		body, _ = sjson.Set(body, path+"ikeMode.value", "main")
	} else {
		body, _ = sjson.Set(body, path+"ikeMode.optionType", "global")
		body, _ = sjson.Set(body, path+"ikeMode.value", data.IkeIntegrityProtocol.ValueString())
	}

	if !data.IkeRekeyIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ikeRekeyInterval.optionType", "variable")
		body, _ = sjson.Set(body, path+"ikeRekeyInterval.value", data.IkeRekeyIntervalVariable.ValueString())
	} else if data.IkeRekeyInterval.IsNull() {
		body, _ = sjson.Set(body, path+"ikeRekeyInterval.optionType", "default")
		body, _ = sjson.Set(body, path+"ikeRekeyInterval.value", 14400)
	} else {
		body, _ = sjson.Set(body, path+"ikeRekeyInterval.optionType", "global")
		body, _ = sjson.Set(body, path+"ikeRekeyInterval.value", data.IkeRekeyInterval.ValueInt64())
	}

	if !data.IkeCiphersuiteVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ikeCiphersuite.optionType", "variable")
		body, _ = sjson.Set(body, path+"ikeCiphersuite.value", data.IkeCiphersuiteVariable.ValueString())
	} else if data.IkeCiphersuite.IsNull() {
		body, _ = sjson.Set(body, path+"ikeCiphersuite.optionType", "default")
		body, _ = sjson.Set(body, path+"ikeCiphersuite.value", "aes256-cbc-sha1")
	} else {
		body, _ = sjson.Set(body, path+"ikeCiphersuite.optionType", "global")
		body, _ = sjson.Set(body, path+"ikeCiphersuite.value", data.IkeCiphersuite.ValueString())
	}

	if !data.IkeDiffieHellmanGroupVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ikeGroup.optionType", "variable")
		body, _ = sjson.Set(body, path+"ikeGroup.value", data.IkeDiffieHellmanGroupVariable.ValueString())
	} else if data.IkeDiffieHellmanGroup.IsNull() {
		body, _ = sjson.Set(body, path+"ikeGroup.optionType", "default")
		body, _ = sjson.Set(body, path+"ikeGroup.value", "16")
	} else {
		body, _ = sjson.Set(body, path+"ikeGroup.optionType", "global")
		body, _ = sjson.Set(body, path+"ikeGroup.value", data.IkeDiffieHellmanGroup.ValueString())
	}

	if !data.IkeIdForLocalEndPointVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ikeLocalId.optionType", "variable")
		body, _ = sjson.Set(body, path+"ikeLocalId.value", data.IkeIdForLocalEndPointVariable.ValueString())
	} else if data.IkeIdForLocalEndPoint.IsNull() {
		body, _ = sjson.Set(body, path+"ikeLocalId.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"ikeLocalId.optionType", "global")
		body, _ = sjson.Set(body, path+"ikeLocalId.value", data.IkeIdForLocalEndPoint.ValueString())
	}

	if !data.IkeIdForRemoteEndPointVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ikeRemoteId.optionType", "variable")
		body, _ = sjson.Set(body, path+"ikeRemoteId.value", data.IkeIdForRemoteEndPointVariable.ValueString())
	} else if data.IkeIdForRemoteEndPoint.IsNull() {
		body, _ = sjson.Set(body, path+"ikeRemoteId.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"ikeRemoteId.optionType", "global")
		body, _ = sjson.Set(body, path+"ikeRemoteId.value", data.IkeIdForRemoteEndPoint.ValueString())
	}

	if !data.IpsecRekeyIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipsecRekeyInterval.optionType", "variable")
		body, _ = sjson.Set(body, path+"ipsecRekeyInterval.value", data.IpsecRekeyIntervalVariable.ValueString())
	} else if data.IpsecRekeyInterval.IsNull() {
		body, _ = sjson.Set(body, path+"ipsecRekeyInterval.optionType", "default")
		body, _ = sjson.Set(body, path+"ipsecRekeyInterval.value", 3600)
	} else {
		body, _ = sjson.Set(body, path+"ipsecRekeyInterval.optionType", "global")
		body, _ = sjson.Set(body, path+"ipsecRekeyInterval.value", data.IpsecRekeyInterval.ValueInt64())
	}

	if !data.IpsecReplayWindowVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipsecReplayWindow.optionType", "variable")
		body, _ = sjson.Set(body, path+"ipsecReplayWindow.value", data.IpsecReplayWindowVariable.ValueString())
	} else if data.IpsecReplayWindow.IsNull() {
		body, _ = sjson.Set(body, path+"ipsecReplayWindow.optionType", "default")
		body, _ = sjson.Set(body, path+"ipsecReplayWindow.value", 512)
	} else {
		body, _ = sjson.Set(body, path+"ipsecReplayWindow.optionType", "global")
		body, _ = sjson.Set(body, path+"ipsecReplayWindow.value", data.IpsecReplayWindow.ValueInt64())
	}

	if !data.IpsecCiphersuiteVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipsecCiphersuite.optionType", "variable")
		body, _ = sjson.Set(body, path+"ipsecCiphersuite.value", data.IpsecCiphersuiteVariable.ValueString())
	} else if data.IpsecCiphersuite.IsNull() {
		body, _ = sjson.Set(body, path+"ipsecCiphersuite.optionType", "default")
		body, _ = sjson.Set(body, path+"ipsecCiphersuite.value", "aes256-gcm")
	} else {
		body, _ = sjson.Set(body, path+"ipsecCiphersuite.optionType", "global")
		body, _ = sjson.Set(body, path+"ipsecCiphersuite.value", data.IpsecCiphersuite.ValueString())
	}

	if !data.PerfectForwardSecrecyVariable.IsNull() {
		body, _ = sjson.Set(body, path+"perfectForwardSecrecy.optionType", "variable")
		body, _ = sjson.Set(body, path+"perfectForwardSecrecy.value", data.PerfectForwardSecrecyVariable.ValueString())
	} else if data.PerfectForwardSecrecy.IsNull() {
		body, _ = sjson.Set(body, path+"perfectForwardSecrecy.optionType", "default")
		body, _ = sjson.Set(body, path+"perfectForwardSecrecy.value", "group-16")
	} else {
		body, _ = sjson.Set(body, path+"perfectForwardSecrecy.optionType", "global")
		body, _ = sjson.Set(body, path+"perfectForwardSecrecy.value", data.PerfectForwardSecrecy.ValueString())
	}

	if !data.TrackerIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tracker.optionType", "variable")
		body, _ = sjson.Set(body, path+"tracker.value", data.TrackerIdVariable.ValueString())
	} else if data.TrackerId.IsNull() {
		body, _ = sjson.Set(body, path+"tracker.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"tracker.optionType", "global")
		body, _ = sjson.Set(body, path+"tracker.value", data.TrackerId.ValueString())
	}

	if !data.TunnelRouteViaVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnelRouteVia.optionType", "variable")
		body, _ = sjson.Set(body, path+"tunnelRouteVia.value", data.TunnelRouteViaVariable.ValueString())
	} else if data.TunnelRouteVia.IsNull() {
		body, _ = sjson.Set(body, path+"tunnelRouteVia.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"tunnelRouteVia.optionType", "global")
		body, _ = sjson.Set(body, path+"tunnelRouteVia.value", data.TunnelRouteVia.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportWANVPNInterfaceIPSEC) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.InterfaceName = types.StringNull()
	data.InterfaceNameVariable = types.StringNull()
	if t := res.Get(path + "ifName.optionType"); t.Exists() {
		va := res.Get(path + "ifName.value")
		if t.String() == "variable" {
			data.InterfaceNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceName = types.StringValue(va.String())
		}
	}
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
	data.InterfaceAddress = types.StringNull()
	data.InterfaceAddressVariable = types.StringNull()
	if t := res.Get(path + "address.address.optionType"); t.Exists() {
		va := res.Get(path + "address.address.value")
		if t.String() == "variable" {
			data.InterfaceAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceAddress = types.StringValue(va.String())
		}
	}
	data.Mask = types.StringNull()
	data.MaskVariable = types.StringNull()
	if t := res.Get(path + "address.mask.optionType"); t.Exists() {
		va := res.Get(path + "address.mask.value")
		if t.String() == "variable" {
			data.MaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Mask = types.StringValue(va.String())
		}
	}
	data.TunnelSourceAddress = types.StringNull()
	data.TunnelSourceAddressVariable = types.StringNull()
	if t := res.Get(path + "tunnelSource.address.optionType"); t.Exists() {
		va := res.Get(path + "tunnelSource.address.value")
		if t.String() == "variable" {
			data.TunnelSourceAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceAddress = types.StringValue(va.String())
		}
	}
	data.TunnelSourceMask = types.StringNull()
	data.TunnelSourceMaskVariable = types.StringNull()
	if t := res.Get(path + "tunnelSource.mask.optionType"); t.Exists() {
		va := res.Get(path + "tunnelSource.mask.value")
		if t.String() == "variable" {
			data.TunnelSourceMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceMask = types.StringValue(va.String())
		}
	}
	data.TunnelSourceInterface = types.StringNull()
	data.TunnelSourceInterfaceVariable = types.StringNull()
	if t := res.Get(path + "tunnelSourceInterface.optionType"); t.Exists() {
		va := res.Get(path + "tunnelSourceInterface.value")
		if t.String() == "variable" {
			data.TunnelSourceInterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceInterface = types.StringValue(va.String())
		}
	}
	data.TunnelDestinationAddress = types.StringNull()
	data.TunnelDestinationAddressVariable = types.StringNull()
	if t := res.Get(path + "tunnelDestination.address.optionType"); t.Exists() {
		va := res.Get(path + "tunnelDestination.address.value")
		if t.String() == "variable" {
			data.TunnelDestinationAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelDestinationAddress = types.StringValue(va.String())
		}
	}
	data.TunnelDestinationMask = types.StringNull()
	data.TunnelDestinationMaskVariable = types.StringNull()
	if t := res.Get(path + "tunnelDestination.mask.optionType"); t.Exists() {
		va := res.Get(path + "tunnelDestination.mask.value")
		if t.String() == "variable" {
			data.TunnelDestinationMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelDestinationMask = types.StringValue(va.String())
		}
	}
	data.ApplicationTunnelType = types.StringNull()
	data.ApplicationTunnelTypeVariable = types.StringNull()
	if t := res.Get(path + "application.optionType"); t.Exists() {
		va := res.Get(path + "application.value")
		if t.String() == "variable" {
			data.ApplicationTunnelTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ApplicationTunnelType = types.StringValue(va.String())
		}
	}
	data.TcpMss = types.Int64Null()
	data.TcpMssVariable = types.StringNull()
	if t := res.Get(path + "tcpMssAdjust.optionType"); t.Exists() {
		va := res.Get(path + "tcpMssAdjust.value")
		if t.String() == "variable" {
			data.TcpMssVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TcpMss = types.Int64Value(va.Int())
		}
	}
	data.ClearDontFragment = types.BoolNull()
	data.ClearDontFragmentVariable = types.StringNull()
	if t := res.Get(path + "clearDontFragment.optionType"); t.Exists() {
		va := res.Get(path + "clearDontFragment.value")
		if t.String() == "variable" {
			data.ClearDontFragmentVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ClearDontFragment = types.BoolValue(va.Bool())
		}
	}
	data.IpMtu = types.Int64Null()
	data.IpMtuVariable = types.StringNull()
	if t := res.Get(path + "mtu.optionType"); t.Exists() {
		va := res.Get(path + "mtu.value")
		if t.String() == "variable" {
			data.IpMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpMtu = types.Int64Value(va.Int())
		}
	}
	data.DpdInterval = types.Int64Null()
	data.DpdIntervalVariable = types.StringNull()
	if t := res.Get(path + "dpdInterval.optionType"); t.Exists() {
		va := res.Get(path + "dpdInterval.value")
		if t.String() == "variable" {
			data.DpdIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DpdInterval = types.Int64Value(va.Int())
		}
	}
	data.DpdRetries = types.Int64Null()
	data.DpdRetriesVariable = types.StringNull()
	if t := res.Get(path + "dpdRetries.optionType"); t.Exists() {
		va := res.Get(path + "dpdRetries.value")
		if t.String() == "variable" {
			data.DpdRetriesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DpdRetries = types.Int64Value(va.Int())
		}
	}
	data.PresharedKeyForIke = types.StringNull()
	data.PresharedKeyForIkeVariable = types.StringNull()
	if t := res.Get(path + "preSharedSecret.optionType"); t.Exists() {
		va := res.Get(path + "preSharedSecret.value")
		if t.String() == "variable" {
			data.PresharedKeyForIkeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PresharedKeyForIke = types.StringValue(va.String())
		}
	}
	data.IkeVersion = types.Int64Null()

	if t := res.Get(path + "ikeVersion.optionType"); t.Exists() {
		va := res.Get(path + "ikeVersion.value")
		if t.String() == "global" {
			data.IkeVersion = types.Int64Value(va.Int())
		}
	}
	data.IkeIntegrityProtocol = types.StringNull()
	data.IkeIntegrityProtocolVariable = types.StringNull()
	if t := res.Get(path + "ikeMode.optionType"); t.Exists() {
		va := res.Get(path + "ikeMode.value")
		if t.String() == "variable" {
			data.IkeIntegrityProtocolVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeIntegrityProtocol = types.StringValue(va.String())
		}
	}
	data.IkeRekeyInterval = types.Int64Null()
	data.IkeRekeyIntervalVariable = types.StringNull()
	if t := res.Get(path + "ikeRekeyInterval.optionType"); t.Exists() {
		va := res.Get(path + "ikeRekeyInterval.value")
		if t.String() == "variable" {
			data.IkeRekeyIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeRekeyInterval = types.Int64Value(va.Int())
		}
	}
	data.IkeCiphersuite = types.StringNull()
	data.IkeCiphersuiteVariable = types.StringNull()
	if t := res.Get(path + "ikeCiphersuite.optionType"); t.Exists() {
		va := res.Get(path + "ikeCiphersuite.value")
		if t.String() == "variable" {
			data.IkeCiphersuiteVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeCiphersuite = types.StringValue(va.String())
		}
	}
	data.IkeDiffieHellmanGroup = types.StringNull()
	data.IkeDiffieHellmanGroupVariable = types.StringNull()
	if t := res.Get(path + "ikeGroup.optionType"); t.Exists() {
		va := res.Get(path + "ikeGroup.value")
		if t.String() == "variable" {
			data.IkeDiffieHellmanGroupVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeDiffieHellmanGroup = types.StringValue(va.String())
		}
	}
	data.IkeIdForLocalEndPoint = types.StringNull()
	data.IkeIdForLocalEndPointVariable = types.StringNull()
	if t := res.Get(path + "ikeLocalId.optionType"); t.Exists() {
		va := res.Get(path + "ikeLocalId.value")
		if t.String() == "variable" {
			data.IkeIdForLocalEndPointVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeIdForLocalEndPoint = types.StringValue(va.String())
		}
	}
	data.IkeIdForRemoteEndPoint = types.StringNull()
	data.IkeIdForRemoteEndPointVariable = types.StringNull()
	if t := res.Get(path + "ikeRemoteId.optionType"); t.Exists() {
		va := res.Get(path + "ikeRemoteId.value")
		if t.String() == "variable" {
			data.IkeIdForRemoteEndPointVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeIdForRemoteEndPoint = types.StringValue(va.String())
		}
	}
	data.IpsecRekeyInterval = types.Int64Null()
	data.IpsecRekeyIntervalVariable = types.StringNull()
	if t := res.Get(path + "ipsecRekeyInterval.optionType"); t.Exists() {
		va := res.Get(path + "ipsecRekeyInterval.value")
		if t.String() == "variable" {
			data.IpsecRekeyIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecRekeyInterval = types.Int64Value(va.Int())
		}
	}
	data.IpsecReplayWindow = types.Int64Null()
	data.IpsecReplayWindowVariable = types.StringNull()
	if t := res.Get(path + "ipsecReplayWindow.optionType"); t.Exists() {
		va := res.Get(path + "ipsecReplayWindow.value")
		if t.String() == "variable" {
			data.IpsecReplayWindowVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecReplayWindow = types.Int64Value(va.Int())
		}
	}
	data.IpsecCiphersuite = types.StringNull()
	data.IpsecCiphersuiteVariable = types.StringNull()
	if t := res.Get(path + "ipsecCiphersuite.optionType"); t.Exists() {
		va := res.Get(path + "ipsecCiphersuite.value")
		if t.String() == "variable" {
			data.IpsecCiphersuiteVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecCiphersuite = types.StringValue(va.String())
		}
	}
	data.PerfectForwardSecrecy = types.StringNull()
	data.PerfectForwardSecrecyVariable = types.StringNull()
	if t := res.Get(path + "perfectForwardSecrecy.optionType"); t.Exists() {
		va := res.Get(path + "perfectForwardSecrecy.value")
		if t.String() == "variable" {
			data.PerfectForwardSecrecyVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PerfectForwardSecrecy = types.StringValue(va.String())
		}
	}
	data.TrackerId = types.StringNull()
	data.TrackerIdVariable = types.StringNull()
	if t := res.Get(path + "tracker.optionType"); t.Exists() {
		va := res.Get(path + "tracker.value")
		if t.String() == "variable" {
			data.TrackerIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackerId = types.StringValue(va.String())
		}
	}
	data.TunnelRouteVia = types.StringNull()
	data.TunnelRouteViaVariable = types.StringNull()
	if t := res.Get(path + "tunnelRouteVia.optionType"); t.Exists() {
		va := res.Get(path + "tunnelRouteVia.value")
		if t.String() == "variable" {
			data.TunnelRouteViaVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelRouteVia = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TransportWANVPNInterfaceIPSEC) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.InterfaceName = types.StringNull()
	data.InterfaceNameVariable = types.StringNull()
	if t := res.Get(path + "ifName.optionType"); t.Exists() {
		va := res.Get(path + "ifName.value")
		if t.String() == "variable" {
			data.InterfaceNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceName = types.StringValue(va.String())
		}
	}
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
	data.InterfaceAddress = types.StringNull()
	data.InterfaceAddressVariable = types.StringNull()
	if t := res.Get(path + "address.address.optionType"); t.Exists() {
		va := res.Get(path + "address.address.value")
		if t.String() == "variable" {
			data.InterfaceAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceAddress = types.StringValue(va.String())
		}
	}
	data.Mask = types.StringNull()
	data.MaskVariable = types.StringNull()
	if t := res.Get(path + "address.mask.optionType"); t.Exists() {
		va := res.Get(path + "address.mask.value")
		if t.String() == "variable" {
			data.MaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Mask = types.StringValue(va.String())
		}
	}
	data.TunnelSourceAddress = types.StringNull()
	data.TunnelSourceAddressVariable = types.StringNull()
	if t := res.Get(path + "tunnelSource.address.optionType"); t.Exists() {
		va := res.Get(path + "tunnelSource.address.value")
		if t.String() == "variable" {
			data.TunnelSourceAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceAddress = types.StringValue(va.String())
		}
	}
	data.TunnelSourceMask = types.StringNull()
	data.TunnelSourceMaskVariable = types.StringNull()
	if t := res.Get(path + "tunnelSource.mask.optionType"); t.Exists() {
		va := res.Get(path + "tunnelSource.mask.value")
		if t.String() == "variable" {
			data.TunnelSourceMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceMask = types.StringValue(va.String())
		}
	}
	data.TunnelSourceInterface = types.StringNull()
	data.TunnelSourceInterfaceVariable = types.StringNull()
	if t := res.Get(path + "tunnelSourceInterface.optionType"); t.Exists() {
		va := res.Get(path + "tunnelSourceInterface.value")
		if t.String() == "variable" {
			data.TunnelSourceInterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceInterface = types.StringValue(va.String())
		}
	}
	data.TunnelDestinationAddress = types.StringNull()
	data.TunnelDestinationAddressVariable = types.StringNull()
	if t := res.Get(path + "tunnelDestination.address.optionType"); t.Exists() {
		va := res.Get(path + "tunnelDestination.address.value")
		if t.String() == "variable" {
			data.TunnelDestinationAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelDestinationAddress = types.StringValue(va.String())
		}
	}
	data.TunnelDestinationMask = types.StringNull()
	data.TunnelDestinationMaskVariable = types.StringNull()
	if t := res.Get(path + "tunnelDestination.mask.optionType"); t.Exists() {
		va := res.Get(path + "tunnelDestination.mask.value")
		if t.String() == "variable" {
			data.TunnelDestinationMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelDestinationMask = types.StringValue(va.String())
		}
	}
	data.ApplicationTunnelType = types.StringNull()
	data.ApplicationTunnelTypeVariable = types.StringNull()
	if t := res.Get(path + "application.optionType"); t.Exists() {
		va := res.Get(path + "application.value")
		if t.String() == "variable" {
			data.ApplicationTunnelTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ApplicationTunnelType = types.StringValue(va.String())
		}
	}
	data.TcpMss = types.Int64Null()
	data.TcpMssVariable = types.StringNull()
	if t := res.Get(path + "tcpMssAdjust.optionType"); t.Exists() {
		va := res.Get(path + "tcpMssAdjust.value")
		if t.String() == "variable" {
			data.TcpMssVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TcpMss = types.Int64Value(va.Int())
		}
	}
	data.ClearDontFragment = types.BoolNull()
	data.ClearDontFragmentVariable = types.StringNull()
	if t := res.Get(path + "clearDontFragment.optionType"); t.Exists() {
		va := res.Get(path + "clearDontFragment.value")
		if t.String() == "variable" {
			data.ClearDontFragmentVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ClearDontFragment = types.BoolValue(va.Bool())
		}
	}
	data.IpMtu = types.Int64Null()
	data.IpMtuVariable = types.StringNull()
	if t := res.Get(path + "mtu.optionType"); t.Exists() {
		va := res.Get(path + "mtu.value")
		if t.String() == "variable" {
			data.IpMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpMtu = types.Int64Value(va.Int())
		}
	}
	data.DpdInterval = types.Int64Null()
	data.DpdIntervalVariable = types.StringNull()
	if t := res.Get(path + "dpdInterval.optionType"); t.Exists() {
		va := res.Get(path + "dpdInterval.value")
		if t.String() == "variable" {
			data.DpdIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DpdInterval = types.Int64Value(va.Int())
		}
	}
	data.DpdRetries = types.Int64Null()
	data.DpdRetriesVariable = types.StringNull()
	if t := res.Get(path + "dpdRetries.optionType"); t.Exists() {
		va := res.Get(path + "dpdRetries.value")
		if t.String() == "variable" {
			data.DpdRetriesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DpdRetries = types.Int64Value(va.Int())
		}
	}
	data.PresharedKeyForIke = types.StringNull()
	data.PresharedKeyForIkeVariable = types.StringNull()
	if t := res.Get(path + "preSharedSecret.optionType"); t.Exists() {
		va := res.Get(path + "preSharedSecret.value")
		if t.String() == "variable" {
			data.PresharedKeyForIkeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PresharedKeyForIke = types.StringValue(va.String())
		}
	}
	data.IkeVersion = types.Int64Null()

	if t := res.Get(path + "ikeVersion.optionType"); t.Exists() {
		va := res.Get(path + "ikeVersion.value")
		if t.String() == "global" {
			data.IkeVersion = types.Int64Value(va.Int())
		}
	}
	data.IkeIntegrityProtocol = types.StringNull()
	data.IkeIntegrityProtocolVariable = types.StringNull()
	if t := res.Get(path + "ikeMode.optionType"); t.Exists() {
		va := res.Get(path + "ikeMode.value")
		if t.String() == "variable" {
			data.IkeIntegrityProtocolVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeIntegrityProtocol = types.StringValue(va.String())
		}
	}
	data.IkeRekeyInterval = types.Int64Null()
	data.IkeRekeyIntervalVariable = types.StringNull()
	if t := res.Get(path + "ikeRekeyInterval.optionType"); t.Exists() {
		va := res.Get(path + "ikeRekeyInterval.value")
		if t.String() == "variable" {
			data.IkeRekeyIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeRekeyInterval = types.Int64Value(va.Int())
		}
	}
	data.IkeCiphersuite = types.StringNull()
	data.IkeCiphersuiteVariable = types.StringNull()
	if t := res.Get(path + "ikeCiphersuite.optionType"); t.Exists() {
		va := res.Get(path + "ikeCiphersuite.value")
		if t.String() == "variable" {
			data.IkeCiphersuiteVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeCiphersuite = types.StringValue(va.String())
		}
	}
	data.IkeDiffieHellmanGroup = types.StringNull()
	data.IkeDiffieHellmanGroupVariable = types.StringNull()
	if t := res.Get(path + "ikeGroup.optionType"); t.Exists() {
		va := res.Get(path + "ikeGroup.value")
		if t.String() == "variable" {
			data.IkeDiffieHellmanGroupVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeDiffieHellmanGroup = types.StringValue(va.String())
		}
	}
	data.IkeIdForLocalEndPoint = types.StringNull()
	data.IkeIdForLocalEndPointVariable = types.StringNull()
	if t := res.Get(path + "ikeLocalId.optionType"); t.Exists() {
		va := res.Get(path + "ikeLocalId.value")
		if t.String() == "variable" {
			data.IkeIdForLocalEndPointVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeIdForLocalEndPoint = types.StringValue(va.String())
		}
	}
	data.IkeIdForRemoteEndPoint = types.StringNull()
	data.IkeIdForRemoteEndPointVariable = types.StringNull()
	if t := res.Get(path + "ikeRemoteId.optionType"); t.Exists() {
		va := res.Get(path + "ikeRemoteId.value")
		if t.String() == "variable" {
			data.IkeIdForRemoteEndPointVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IkeIdForRemoteEndPoint = types.StringValue(va.String())
		}
	}
	data.IpsecRekeyInterval = types.Int64Null()
	data.IpsecRekeyIntervalVariable = types.StringNull()
	if t := res.Get(path + "ipsecRekeyInterval.optionType"); t.Exists() {
		va := res.Get(path + "ipsecRekeyInterval.value")
		if t.String() == "variable" {
			data.IpsecRekeyIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecRekeyInterval = types.Int64Value(va.Int())
		}
	}
	data.IpsecReplayWindow = types.Int64Null()
	data.IpsecReplayWindowVariable = types.StringNull()
	if t := res.Get(path + "ipsecReplayWindow.optionType"); t.Exists() {
		va := res.Get(path + "ipsecReplayWindow.value")
		if t.String() == "variable" {
			data.IpsecReplayWindowVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecReplayWindow = types.Int64Value(va.Int())
		}
	}
	data.IpsecCiphersuite = types.StringNull()
	data.IpsecCiphersuiteVariable = types.StringNull()
	if t := res.Get(path + "ipsecCiphersuite.optionType"); t.Exists() {
		va := res.Get(path + "ipsecCiphersuite.value")
		if t.String() == "variable" {
			data.IpsecCiphersuiteVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecCiphersuite = types.StringValue(va.String())
		}
	}
	data.PerfectForwardSecrecy = types.StringNull()
	data.PerfectForwardSecrecyVariable = types.StringNull()
	if t := res.Get(path + "perfectForwardSecrecy.optionType"); t.Exists() {
		va := res.Get(path + "perfectForwardSecrecy.value")
		if t.String() == "variable" {
			data.PerfectForwardSecrecyVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PerfectForwardSecrecy = types.StringValue(va.String())
		}
	}
	data.TrackerId = types.StringNull()
	data.TrackerIdVariable = types.StringNull()
	if t := res.Get(path + "tracker.optionType"); t.Exists() {
		va := res.Get(path + "tracker.value")
		if t.String() == "variable" {
			data.TrackerIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackerId = types.StringValue(va.String())
		}
	}
	data.TunnelRouteVia = types.StringNull()
	data.TunnelRouteViaVariable = types.StringNull()
	if t := res.Get(path + "tunnelRouteVia.optionType"); t.Exists() {
		va := res.Get(path + "tunnelRouteVia.value")
		if t.String() == "variable" {
			data.TunnelRouteViaVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelRouteVia = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *TransportWANVPNInterfaceIPSEC) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.TransportWanVpnProfileParcelId.IsNull() {
		return false
	}
	if !data.InterfaceName.IsNull() {
		return false
	}
	if !data.InterfaceNameVariable.IsNull() {
		return false
	}
	if !data.Shutdown.IsNull() {
		return false
	}
	if !data.ShutdownVariable.IsNull() {
		return false
	}
	if !data.InterfaceDescription.IsNull() {
		return false
	}
	if !data.InterfaceDescriptionVariable.IsNull() {
		return false
	}
	if !data.InterfaceAddress.IsNull() {
		return false
	}
	if !data.InterfaceAddressVariable.IsNull() {
		return false
	}
	if !data.Mask.IsNull() {
		return false
	}
	if !data.MaskVariable.IsNull() {
		return false
	}
	if !data.TunnelSourceAddress.IsNull() {
		return false
	}
	if !data.TunnelSourceAddressVariable.IsNull() {
		return false
	}
	if !data.TunnelSourceMask.IsNull() {
		return false
	}
	if !data.TunnelSourceMaskVariable.IsNull() {
		return false
	}
	if !data.TunnelSourceInterface.IsNull() {
		return false
	}
	if !data.TunnelSourceInterfaceVariable.IsNull() {
		return false
	}
	if !data.TunnelDestinationAddress.IsNull() {
		return false
	}
	if !data.TunnelDestinationAddressVariable.IsNull() {
		return false
	}
	if !data.TunnelDestinationMask.IsNull() {
		return false
	}
	if !data.TunnelDestinationMaskVariable.IsNull() {
		return false
	}
	if !data.ApplicationTunnelType.IsNull() {
		return false
	}
	if !data.ApplicationTunnelTypeVariable.IsNull() {
		return false
	}
	if !data.TcpMss.IsNull() {
		return false
	}
	if !data.TcpMssVariable.IsNull() {
		return false
	}
	if !data.ClearDontFragment.IsNull() {
		return false
	}
	if !data.ClearDontFragmentVariable.IsNull() {
		return false
	}
	if !data.IpMtu.IsNull() {
		return false
	}
	if !data.IpMtuVariable.IsNull() {
		return false
	}
	if !data.DpdInterval.IsNull() {
		return false
	}
	if !data.DpdIntervalVariable.IsNull() {
		return false
	}
	if !data.DpdRetries.IsNull() {
		return false
	}
	if !data.DpdRetriesVariable.IsNull() {
		return false
	}
	if !data.PresharedKeyForIke.IsNull() {
		return false
	}
	if !data.PresharedKeyForIkeVariable.IsNull() {
		return false
	}
	if !data.IkeVersion.IsNull() {
		return false
	}
	if !data.IkeIntegrityProtocol.IsNull() {
		return false
	}
	if !data.IkeIntegrityProtocolVariable.IsNull() {
		return false
	}
	if !data.IkeRekeyInterval.IsNull() {
		return false
	}
	if !data.IkeRekeyIntervalVariable.IsNull() {
		return false
	}
	if !data.IkeCiphersuite.IsNull() {
		return false
	}
	if !data.IkeCiphersuiteVariable.IsNull() {
		return false
	}
	if !data.IkeDiffieHellmanGroup.IsNull() {
		return false
	}
	if !data.IkeDiffieHellmanGroupVariable.IsNull() {
		return false
	}
	if !data.IkeIdForLocalEndPoint.IsNull() {
		return false
	}
	if !data.IkeIdForLocalEndPointVariable.IsNull() {
		return false
	}
	if !data.IkeIdForRemoteEndPoint.IsNull() {
		return false
	}
	if !data.IkeIdForRemoteEndPointVariable.IsNull() {
		return false
	}
	if !data.IpsecRekeyInterval.IsNull() {
		return false
	}
	if !data.IpsecRekeyIntervalVariable.IsNull() {
		return false
	}
	if !data.IpsecReplayWindow.IsNull() {
		return false
	}
	if !data.IpsecReplayWindowVariable.IsNull() {
		return false
	}
	if !data.IpsecCiphersuite.IsNull() {
		return false
	}
	if !data.IpsecCiphersuiteVariable.IsNull() {
		return false
	}
	if !data.PerfectForwardSecrecy.IsNull() {
		return false
	}
	if !data.PerfectForwardSecrecyVariable.IsNull() {
		return false
	}
	if !data.TrackerId.IsNull() {
		return false
	}
	if !data.TrackerIdVariable.IsNull() {
		return false
	}
	if !data.TunnelRouteVia.IsNull() {
		return false
	}
	if !data.TunnelRouteViaVariable.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
