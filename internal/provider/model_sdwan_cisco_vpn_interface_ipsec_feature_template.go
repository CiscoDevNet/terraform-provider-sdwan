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
type CiscoVPNInterfaceIPSec struct {
	Id                                 types.String `tfsdk:"id"`
	Version                            types.Int64  `tfsdk:"version"`
	TemplateType                       types.String `tfsdk:"template_type"`
	Name                               types.String `tfsdk:"name"`
	Description                        types.String `tfsdk:"description"`
	DeviceTypes                        types.Set    `tfsdk:"device_types"`
	InterfaceName                      types.String `tfsdk:"interface_name"`
	InterfaceNameVariable              types.String `tfsdk:"interface_name_variable"`
	Shutdown                           types.Bool   `tfsdk:"shutdown"`
	ShutdownVariable                   types.String `tfsdk:"shutdown_variable"`
	InterfaceDescription               types.String `tfsdk:"interface_description"`
	InterfaceDescriptionVariable       types.String `tfsdk:"interface_description_variable"`
	IpAddress                          types.String `tfsdk:"ip_address"`
	IpAddressVariable                  types.String `tfsdk:"ip_address_variable"`
	TunnelSource                       types.String `tfsdk:"tunnel_source"`
	TunnelSourceVariable               types.String `tfsdk:"tunnel_source_variable"`
	TunnelSourceInterface              types.String `tfsdk:"tunnel_source_interface"`
	TunnelSourceInterfaceVariable      types.String `tfsdk:"tunnel_source_interface_variable"`
	TunnelDestination                  types.String `tfsdk:"tunnel_destination"`
	TunnelDestinationVariable          types.String `tfsdk:"tunnel_destination_variable"`
	Application                        types.String `tfsdk:"application"`
	ApplicationVariable                types.String `tfsdk:"application_variable"`
	TcpMssAdjust                       types.Int64  `tfsdk:"tcp_mss_adjust"`
	TcpMssAdjustVariable               types.String `tfsdk:"tcp_mss_adjust_variable"`
	ClearDontFragment                  types.Bool   `tfsdk:"clear_dont_fragment"`
	ClearDontFragmentVariable          types.String `tfsdk:"clear_dont_fragment_variable"`
	Mtu                                types.Int64  `tfsdk:"mtu"`
	MtuVariable                        types.String `tfsdk:"mtu_variable"`
	DeadPeerDetectionInterval          types.Int64  `tfsdk:"dead_peer_detection_interval"`
	DeadPeerDetectionIntervalVariable  types.String `tfsdk:"dead_peer_detection_interval_variable"`
	DeadPeerDetectionRetries           types.Int64  `tfsdk:"dead_peer_detection_retries"`
	DeadPeerDetectionRetriesVariable   types.String `tfsdk:"dead_peer_detection_retries_variable"`
	IkeVersion                         types.Int64  `tfsdk:"ike_version"`
	IkeMode                            types.String `tfsdk:"ike_mode"`
	IkeModeVariable                    types.String `tfsdk:"ike_mode_variable"`
	IkeRekeyInterval                   types.Int64  `tfsdk:"ike_rekey_interval"`
	IkeRekeyIntervalVariable           types.String `tfsdk:"ike_rekey_interval_variable"`
	IkeCiphersuite                     types.String `tfsdk:"ike_ciphersuite"`
	IkeCiphersuiteVariable             types.String `tfsdk:"ike_ciphersuite_variable"`
	IkeGroup                           types.String `tfsdk:"ike_group"`
	IkeGroupVariable                   types.String `tfsdk:"ike_group_variable"`
	IkePreSharedKey                    types.String `tfsdk:"ike_pre_shared_key"`
	IkePreSharedKeyVariable            types.String `tfsdk:"ike_pre_shared_key_variable"`
	IkePreSharedKeyLocalId             types.String `tfsdk:"ike_pre_shared_key_local_id"`
	IkePreSharedKeyLocalIdVariable     types.String `tfsdk:"ike_pre_shared_key_local_id_variable"`
	IkePreSharedKeyRemoteId            types.String `tfsdk:"ike_pre_shared_key_remote_id"`
	IkePreSharedKeyRemoteIdVariable    types.String `tfsdk:"ike_pre_shared_key_remote_id_variable"`
	IpsecRekeyInterval                 types.Int64  `tfsdk:"ipsec_rekey_interval"`
	IpsecRekeyIntervalVariable         types.String `tfsdk:"ipsec_rekey_interval_variable"`
	IpsecReplayWindow                  types.Int64  `tfsdk:"ipsec_replay_window"`
	IpsecReplayWindowVariable          types.String `tfsdk:"ipsec_replay_window_variable"`
	IpsecCiphersuite                   types.String `tfsdk:"ipsec_ciphersuite"`
	IpsecCiphersuiteVariable           types.String `tfsdk:"ipsec_ciphersuite_variable"`
	IpsecPerfectForwardSecrecy         types.String `tfsdk:"ipsec_perfect_forward_secrecy"`
	IpsecPerfectForwardSecrecyVariable types.String `tfsdk:"ipsec_perfect_forward_secrecy_variable"`
	Tracker                            types.Set    `tfsdk:"tracker"`
	TrackerVariable                    types.String `tfsdk:"tracker_variable"`
	TunnelRouteVia                     types.String `tfsdk:"tunnel_route_via"`
	TunnelRouteViaVariable             types.String `tfsdk:"tunnel_route_via_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoVPNInterfaceIPSec) getModel() string {
	return "cisco_vpn_interface_ipsec"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CiscoVPNInterfaceIPSec) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_vpn_interface_ipsec")
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

	if !data.IpAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ip.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.address."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ip.address."+"vipVariableName", data.IpAddressVariable.ValueString())
	} else if data.IpAddress.IsNull() {
		body, _ = sjson.Set(body, path+"ip.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.address."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ip.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ip.address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ip.address."+"vipValue", data.IpAddress.ValueString())
	}

	if !data.TunnelSourceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-source."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-source."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-source."+"vipVariableName", data.TunnelSourceVariable.ValueString())
	} else if data.TunnelSource.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-source."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-source."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-source."+"vipValue", data.TunnelSource.ValueString())
	}

	if !data.TunnelSourceInterfaceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-source-interface."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-source-interface."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-source-interface."+"vipVariableName", data.TunnelSourceInterfaceVariable.ValueString())
	} else if data.TunnelSourceInterface.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-source-interface."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-source-interface."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-source-interface."+"vipValue", data.TunnelSourceInterface.ValueString())
	}

	if !data.TunnelDestinationVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-destination."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-destination."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-destination."+"vipVariableName", data.TunnelDestinationVariable.ValueString())
	} else if data.TunnelDestination.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-destination."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-destination."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-destination."+"vipValue", data.TunnelDestination.ValueString())
	}

	if !data.ApplicationVariable.IsNull() {
		body, _ = sjson.Set(body, path+"application."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"application."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"application."+"vipVariableName", data.ApplicationVariable.ValueString())
	} else if data.Application.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"application."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"application."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"application."+"vipValue", data.Application.ValueString())
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

	if !data.ClearDontFragmentVariable.IsNull() {
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipVariableName", data.ClearDontFragmentVariable.ValueString())
	} else if data.ClearDontFragment.IsNull() {
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"clear-dont-fragment."+"vipValue", strconv.FormatBool(data.ClearDontFragment.ValueBool()))
	}

	if !data.MtuVariable.IsNull() {
		body, _ = sjson.Set(body, path+"mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mtu."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"mtu."+"vipVariableName", data.MtuVariable.ValueString())
	} else if data.Mtu.IsNull() {
		body, _ = sjson.Set(body, path+"mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mtu."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mtu."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"mtu."+"vipValue", data.Mtu.ValueInt64())
	}

	if !data.DeadPeerDetectionIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"dead-peer-detection.dpd-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"dead-peer-detection.dpd-interval."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"dead-peer-detection.dpd-interval."+"vipVariableName", data.DeadPeerDetectionIntervalVariable.ValueString())
	} else if data.DeadPeerDetectionInterval.IsNull() {
		body, _ = sjson.Set(body, path+"dead-peer-detection.dpd-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"dead-peer-detection.dpd-interval."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"dead-peer-detection.dpd-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"dead-peer-detection.dpd-interval."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"dead-peer-detection.dpd-interval."+"vipValue", data.DeadPeerDetectionInterval.ValueInt64())
	}

	if !data.DeadPeerDetectionRetriesVariable.IsNull() {
		body, _ = sjson.Set(body, path+"dead-peer-detection.dpd-retries."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"dead-peer-detection.dpd-retries."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"dead-peer-detection.dpd-retries."+"vipVariableName", data.DeadPeerDetectionRetriesVariable.ValueString())
	} else if data.DeadPeerDetectionRetries.IsNull() {
		body, _ = sjson.Set(body, path+"dead-peer-detection.dpd-retries."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"dead-peer-detection.dpd-retries."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"dead-peer-detection.dpd-retries."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"dead-peer-detection.dpd-retries."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"dead-peer-detection.dpd-retries."+"vipValue", data.DeadPeerDetectionRetries.ValueInt64())
	}
	if data.IkeVersion.IsNull() {
		body, _ = sjson.Set(body, path+"ike.ike-version."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.ike-version."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ike.ike-version."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.ike-version."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ike.ike-version."+"vipValue", data.IkeVersion.ValueInt64())
	}

	if !data.IkeModeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ike.ike-mode."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.ike-mode."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ike.ike-mode."+"vipVariableName", data.IkeModeVariable.ValueString())
	} else if data.IkeMode.IsNull() {
		body, _ = sjson.Set(body, path+"ike.ike-mode."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.ike-mode."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ike.ike-mode."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.ike-mode."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ike.ike-mode."+"vipValue", data.IkeMode.ValueString())
	}

	if !data.IkeRekeyIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ike.ike-rekey-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.ike-rekey-interval."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ike.ike-rekey-interval."+"vipVariableName", data.IkeRekeyIntervalVariable.ValueString())
	} else if data.IkeRekeyInterval.IsNull() {
		body, _ = sjson.Set(body, path+"ike.ike-rekey-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.ike-rekey-interval."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ike.ike-rekey-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.ike-rekey-interval."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ike.ike-rekey-interval."+"vipValue", data.IkeRekeyInterval.ValueInt64())
	}

	if !data.IkeCiphersuiteVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ike.ike-ciphersuite."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.ike-ciphersuite."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ike.ike-ciphersuite."+"vipVariableName", data.IkeCiphersuiteVariable.ValueString())
	} else if data.IkeCiphersuite.IsNull() {
		body, _ = sjson.Set(body, path+"ike.ike-ciphersuite."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.ike-ciphersuite."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ike.ike-ciphersuite."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.ike-ciphersuite."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ike.ike-ciphersuite."+"vipValue", data.IkeCiphersuite.ValueString())
	}

	if !data.IkeGroupVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ike.ike-group."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.ike-group."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ike.ike-group."+"vipVariableName", data.IkeGroupVariable.ValueString())
	} else if data.IkeGroup.IsNull() {
		body, _ = sjson.Set(body, path+"ike.ike-group."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.ike-group."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ike.ike-group."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.ike-group."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ike.ike-group."+"vipValue", data.IkeGroup.ValueString())
	}

	if !data.IkePreSharedKeyVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.pre-shared-secret."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.pre-shared-secret."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.pre-shared-secret."+"vipVariableName", data.IkePreSharedKeyVariable.ValueString())
	} else if data.IkePreSharedKey.IsNull() {
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.pre-shared-secret."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.pre-shared-secret."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.pre-shared-secret."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.pre-shared-secret."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.pre-shared-secret."+"vipValue", data.IkePreSharedKey.ValueString())
	}

	if !data.IkePreSharedKeyLocalIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.ike-local-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.ike-local-id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.ike-local-id."+"vipVariableName", data.IkePreSharedKeyLocalIdVariable.ValueString())
	} else if data.IkePreSharedKeyLocalId.IsNull() {
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.ike-local-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.ike-local-id."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.ike-local-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.ike-local-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.ike-local-id."+"vipValue", data.IkePreSharedKeyLocalId.ValueString())
	}

	if !data.IkePreSharedKeyRemoteIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.ike-remote-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.ike-remote-id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.ike-remote-id."+"vipVariableName", data.IkePreSharedKeyRemoteIdVariable.ValueString())
	} else if data.IkePreSharedKeyRemoteId.IsNull() {
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.ike-remote-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.ike-remote-id."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.ike-remote-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.ike-remote-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ike.authentication-type.pre-shared-key.ike-remote-id."+"vipValue", data.IkePreSharedKeyRemoteId.ValueString())
	}

	if !data.IpsecRekeyIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.ipsec-rekey-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.ipsec-rekey-interval."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipsec.ipsec-rekey-interval."+"vipVariableName", data.IpsecRekeyIntervalVariable.ValueString())
	} else if data.IpsecRekeyInterval.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.ipsec-rekey-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.ipsec-rekey-interval."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipsec.ipsec-rekey-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.ipsec-rekey-interval."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipsec.ipsec-rekey-interval."+"vipValue", data.IpsecRekeyInterval.ValueInt64())
	}

	if !data.IpsecReplayWindowVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.ipsec-replay-window."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.ipsec-replay-window."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipsec.ipsec-replay-window."+"vipVariableName", data.IpsecReplayWindowVariable.ValueString())
	} else if data.IpsecReplayWindow.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.ipsec-replay-window."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.ipsec-replay-window."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipsec.ipsec-replay-window."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.ipsec-replay-window."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipsec.ipsec-replay-window."+"vipValue", data.IpsecReplayWindow.ValueInt64())
	}

	if !data.IpsecCiphersuiteVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.ipsec-ciphersuite."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.ipsec-ciphersuite."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipsec.ipsec-ciphersuite."+"vipVariableName", data.IpsecCiphersuiteVariable.ValueString())
	} else if data.IpsecCiphersuite.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.ipsec-ciphersuite."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.ipsec-ciphersuite."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipsec.ipsec-ciphersuite."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.ipsec-ciphersuite."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipsec.ipsec-ciphersuite."+"vipValue", data.IpsecCiphersuite.ValueString())
	}

	if !data.IpsecPerfectForwardSecrecyVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.perfect-forward-secrecy."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.perfect-forward-secrecy."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ipsec.perfect-forward-secrecy."+"vipVariableName", data.IpsecPerfectForwardSecrecyVariable.ValueString())
	} else if data.IpsecPerfectForwardSecrecy.IsNull() {
		body, _ = sjson.Set(body, path+"ipsec.perfect-forward-secrecy."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.perfect-forward-secrecy."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ipsec.perfect-forward-secrecy."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ipsec.perfect-forward-secrecy."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ipsec.perfect-forward-secrecy."+"vipValue", data.IpsecPerfectForwardSecrecy.ValueString())
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

	if !data.TunnelRouteViaVariable.IsNull() {
		body, _ = sjson.Set(body, path+"tunnel-route-via."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-route-via."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"tunnel-route-via."+"vipVariableName", data.TunnelRouteViaVariable.ValueString())
	} else if data.TunnelRouteVia.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"tunnel-route-via."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"tunnel-route-via."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tunnel-route-via."+"vipValue", data.TunnelRouteVia.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoVPNInterfaceIPSec) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "ip.address.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IpAddress = types.StringNull()

			v := res.Get(path + "ip.address.vipVariableName")
			data.IpAddressVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IpAddress = types.StringNull()
			data.IpAddressVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ip.address.vipValue")
			data.IpAddress = types.StringValue(v.String())
			data.IpAddressVariable = types.StringNull()
		}
	} else {
		data.IpAddress = types.StringNull()
		data.IpAddressVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-source.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelSource = types.StringNull()

			v := res.Get(path + "tunnel-source.vipVariableName")
			data.TunnelSourceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelSource = types.StringNull()
			data.TunnelSourceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-source.vipValue")
			data.TunnelSource = types.StringValue(v.String())
			data.TunnelSourceVariable = types.StringNull()
		}
	} else {
		data.TunnelSource = types.StringNull()
		data.TunnelSourceVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-source-interface.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelSourceInterface = types.StringNull()

			v := res.Get(path + "tunnel-source-interface.vipVariableName")
			data.TunnelSourceInterfaceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelSourceInterface = types.StringNull()
			data.TunnelSourceInterfaceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-source-interface.vipValue")
			data.TunnelSourceInterface = types.StringValue(v.String())
			data.TunnelSourceInterfaceVariable = types.StringNull()
		}
	} else {
		data.TunnelSourceInterface = types.StringNull()
		data.TunnelSourceInterfaceVariable = types.StringNull()
	}
	if value := res.Get(path + "tunnel-destination.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelDestination = types.StringNull()

			v := res.Get(path + "tunnel-destination.vipVariableName")
			data.TunnelDestinationVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelDestination = types.StringNull()
			data.TunnelDestinationVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-destination.vipValue")
			data.TunnelDestination = types.StringValue(v.String())
			data.TunnelDestinationVariable = types.StringNull()
		}
	} else {
		data.TunnelDestination = types.StringNull()
		data.TunnelDestinationVariable = types.StringNull()
	}
	if value := res.Get(path + "application.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Application = types.StringNull()

			v := res.Get(path + "application.vipVariableName")
			data.ApplicationVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Application = types.StringNull()
			data.ApplicationVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "application.vipValue")
			data.Application = types.StringValue(v.String())
			data.ApplicationVariable = types.StringNull()
		}
	} else {
		data.Application = types.StringNull()
		data.ApplicationVariable = types.StringNull()
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
	if value := res.Get(path + "clear-dont-fragment.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ClearDontFragment = types.BoolNull()

			v := res.Get(path + "clear-dont-fragment.vipVariableName")
			data.ClearDontFragmentVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ClearDontFragment = types.BoolNull()
			data.ClearDontFragmentVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "clear-dont-fragment.vipValue")
			data.ClearDontFragment = types.BoolValue(v.Bool())
			data.ClearDontFragmentVariable = types.StringNull()
		}
	} else {
		data.ClearDontFragment = types.BoolNull()
		data.ClearDontFragmentVariable = types.StringNull()
	}
	if value := res.Get(path + "mtu.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Mtu = types.Int64Null()

			v := res.Get(path + "mtu.vipVariableName")
			data.MtuVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Mtu = types.Int64Null()
			data.MtuVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "mtu.vipValue")
			data.Mtu = types.Int64Value(v.Int())
			data.MtuVariable = types.StringNull()
		}
	} else {
		data.Mtu = types.Int64Null()
		data.MtuVariable = types.StringNull()
	}
	if value := res.Get(path + "dead-peer-detection.dpd-interval.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DeadPeerDetectionInterval = types.Int64Null()

			v := res.Get(path + "dead-peer-detection.dpd-interval.vipVariableName")
			data.DeadPeerDetectionIntervalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DeadPeerDetectionInterval = types.Int64Null()
			data.DeadPeerDetectionIntervalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "dead-peer-detection.dpd-interval.vipValue")
			data.DeadPeerDetectionInterval = types.Int64Value(v.Int())
			data.DeadPeerDetectionIntervalVariable = types.StringNull()
		}
	} else {
		data.DeadPeerDetectionInterval = types.Int64Null()
		data.DeadPeerDetectionIntervalVariable = types.StringNull()
	}
	if value := res.Get(path + "dead-peer-detection.dpd-retries.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DeadPeerDetectionRetries = types.Int64Null()

			v := res.Get(path + "dead-peer-detection.dpd-retries.vipVariableName")
			data.DeadPeerDetectionRetriesVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DeadPeerDetectionRetries = types.Int64Null()
			data.DeadPeerDetectionRetriesVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "dead-peer-detection.dpd-retries.vipValue")
			data.DeadPeerDetectionRetries = types.Int64Value(v.Int())
			data.DeadPeerDetectionRetriesVariable = types.StringNull()
		}
	} else {
		data.DeadPeerDetectionRetries = types.Int64Null()
		data.DeadPeerDetectionRetriesVariable = types.StringNull()
	}
	if value := res.Get(path + "ike.ike-version.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IkeVersion = types.Int64Null()

		} else if value.String() == "ignore" {
			data.IkeVersion = types.Int64Null()

		} else if value.String() == "constant" {
			v := res.Get(path + "ike.ike-version.vipValue")
			data.IkeVersion = types.Int64Value(v.Int())

		}
	} else {
		data.IkeVersion = types.Int64Null()

	}
	if value := res.Get(path + "ike.ike-mode.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IkeMode = types.StringNull()

			v := res.Get(path + "ike.ike-mode.vipVariableName")
			data.IkeModeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IkeMode = types.StringNull()
			data.IkeModeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ike.ike-mode.vipValue")
			data.IkeMode = types.StringValue(v.String())
			data.IkeModeVariable = types.StringNull()
		}
	} else {
		data.IkeMode = types.StringNull()
		data.IkeModeVariable = types.StringNull()
	}
	if value := res.Get(path + "ike.ike-rekey-interval.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IkeRekeyInterval = types.Int64Null()

			v := res.Get(path + "ike.ike-rekey-interval.vipVariableName")
			data.IkeRekeyIntervalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IkeRekeyInterval = types.Int64Null()
			data.IkeRekeyIntervalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ike.ike-rekey-interval.vipValue")
			data.IkeRekeyInterval = types.Int64Value(v.Int())
			data.IkeRekeyIntervalVariable = types.StringNull()
		}
	} else {
		data.IkeRekeyInterval = types.Int64Null()
		data.IkeRekeyIntervalVariable = types.StringNull()
	}
	if value := res.Get(path + "ike.ike-ciphersuite.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IkeCiphersuite = types.StringNull()

			v := res.Get(path + "ike.ike-ciphersuite.vipVariableName")
			data.IkeCiphersuiteVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IkeCiphersuite = types.StringNull()
			data.IkeCiphersuiteVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ike.ike-ciphersuite.vipValue")
			data.IkeCiphersuite = types.StringValue(v.String())
			data.IkeCiphersuiteVariable = types.StringNull()
		}
	} else {
		data.IkeCiphersuite = types.StringNull()
		data.IkeCiphersuiteVariable = types.StringNull()
	}
	if value := res.Get(path + "ike.ike-group.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IkeGroup = types.StringNull()

			v := res.Get(path + "ike.ike-group.vipVariableName")
			data.IkeGroupVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IkeGroup = types.StringNull()
			data.IkeGroupVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ike.ike-group.vipValue")
			data.IkeGroup = types.StringValue(v.String())
			data.IkeGroupVariable = types.StringNull()
		}
	} else {
		data.IkeGroup = types.StringNull()
		data.IkeGroupVariable = types.StringNull()
	}
	if value := res.Get(path + "ike.authentication-type.pre-shared-key.pre-shared-secret.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IkePreSharedKey = types.StringNull()

			v := res.Get(path + "ike.authentication-type.pre-shared-key.pre-shared-secret.vipVariableName")
			data.IkePreSharedKeyVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IkePreSharedKey = types.StringNull()
			data.IkePreSharedKeyVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ike.authentication-type.pre-shared-key.pre-shared-secret.vipValue")
			data.IkePreSharedKey = types.StringValue(v.String())
			data.IkePreSharedKeyVariable = types.StringNull()
		}
	} else {
		data.IkePreSharedKey = types.StringNull()
		data.IkePreSharedKeyVariable = types.StringNull()
	}
	if value := res.Get(path + "ike.authentication-type.pre-shared-key.ike-local-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IkePreSharedKeyLocalId = types.StringNull()

			v := res.Get(path + "ike.authentication-type.pre-shared-key.ike-local-id.vipVariableName")
			data.IkePreSharedKeyLocalIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IkePreSharedKeyLocalId = types.StringNull()
			data.IkePreSharedKeyLocalIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ike.authentication-type.pre-shared-key.ike-local-id.vipValue")
			data.IkePreSharedKeyLocalId = types.StringValue(v.String())
			data.IkePreSharedKeyLocalIdVariable = types.StringNull()
		}
	} else {
		data.IkePreSharedKeyLocalId = types.StringNull()
		data.IkePreSharedKeyLocalIdVariable = types.StringNull()
	}
	if value := res.Get(path + "ike.authentication-type.pre-shared-key.ike-remote-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IkePreSharedKeyRemoteId = types.StringNull()

			v := res.Get(path + "ike.authentication-type.pre-shared-key.ike-remote-id.vipVariableName")
			data.IkePreSharedKeyRemoteIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IkePreSharedKeyRemoteId = types.StringNull()
			data.IkePreSharedKeyRemoteIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ike.authentication-type.pre-shared-key.ike-remote-id.vipValue")
			data.IkePreSharedKeyRemoteId = types.StringValue(v.String())
			data.IkePreSharedKeyRemoteIdVariable = types.StringNull()
		}
	} else {
		data.IkePreSharedKeyRemoteId = types.StringNull()
		data.IkePreSharedKeyRemoteIdVariable = types.StringNull()
	}
	if value := res.Get(path + "ipsec.ipsec-rekey-interval.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IpsecRekeyInterval = types.Int64Null()

			v := res.Get(path + "ipsec.ipsec-rekey-interval.vipVariableName")
			data.IpsecRekeyIntervalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IpsecRekeyInterval = types.Int64Null()
			data.IpsecRekeyIntervalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipsec.ipsec-rekey-interval.vipValue")
			data.IpsecRekeyInterval = types.Int64Value(v.Int())
			data.IpsecRekeyIntervalVariable = types.StringNull()
		}
	} else {
		data.IpsecRekeyInterval = types.Int64Null()
		data.IpsecRekeyIntervalVariable = types.StringNull()
	}
	if value := res.Get(path + "ipsec.ipsec-replay-window.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IpsecReplayWindow = types.Int64Null()

			v := res.Get(path + "ipsec.ipsec-replay-window.vipVariableName")
			data.IpsecReplayWindowVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IpsecReplayWindow = types.Int64Null()
			data.IpsecReplayWindowVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipsec.ipsec-replay-window.vipValue")
			data.IpsecReplayWindow = types.Int64Value(v.Int())
			data.IpsecReplayWindowVariable = types.StringNull()
		}
	} else {
		data.IpsecReplayWindow = types.Int64Null()
		data.IpsecReplayWindowVariable = types.StringNull()
	}
	if value := res.Get(path + "ipsec.ipsec-ciphersuite.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IpsecCiphersuite = types.StringNull()

			v := res.Get(path + "ipsec.ipsec-ciphersuite.vipVariableName")
			data.IpsecCiphersuiteVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IpsecCiphersuite = types.StringNull()
			data.IpsecCiphersuiteVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipsec.ipsec-ciphersuite.vipValue")
			data.IpsecCiphersuite = types.StringValue(v.String())
			data.IpsecCiphersuiteVariable = types.StringNull()
		}
	} else {
		data.IpsecCiphersuite = types.StringNull()
		data.IpsecCiphersuiteVariable = types.StringNull()
	}
	if value := res.Get(path + "ipsec.perfect-forward-secrecy.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IpsecPerfectForwardSecrecy = types.StringNull()

			v := res.Get(path + "ipsec.perfect-forward-secrecy.vipVariableName")
			data.IpsecPerfectForwardSecrecyVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IpsecPerfectForwardSecrecy = types.StringNull()
			data.IpsecPerfectForwardSecrecyVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ipsec.perfect-forward-secrecy.vipValue")
			data.IpsecPerfectForwardSecrecy = types.StringValue(v.String())
			data.IpsecPerfectForwardSecrecyVariable = types.StringNull()
		}
	} else {
		data.IpsecPerfectForwardSecrecy = types.StringNull()
		data.IpsecPerfectForwardSecrecyVariable = types.StringNull()
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
	if value := res.Get(path + "tunnel-route-via.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TunnelRouteVia = types.StringNull()

			v := res.Get(path + "tunnel-route-via.vipVariableName")
			data.TunnelRouteViaVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TunnelRouteVia = types.StringNull()
			data.TunnelRouteViaVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "tunnel-route-via.vipValue")
			data.TunnelRouteVia = types.StringValue(v.String())
			data.TunnelRouteViaVariable = types.StringNull()
		}
	} else {
		data.TunnelRouteVia = types.StringNull()
		data.TunnelRouteViaVariable = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoVPNInterfaceIPSec) hasChanges(ctx context.Context, state *CiscoVPNInterfaceIPSec) bool {
	hasChanges := false
	if !data.InterfaceName.Equal(state.InterfaceName) {
		hasChanges = true
	}
	if !data.Shutdown.Equal(state.Shutdown) {
		hasChanges = true
	}
	if !data.InterfaceDescription.Equal(state.InterfaceDescription) {
		hasChanges = true
	}
	if !data.IpAddress.Equal(state.IpAddress) {
		hasChanges = true
	}
	if !data.TunnelSource.Equal(state.TunnelSource) {
		hasChanges = true
	}
	if !data.TunnelSourceInterface.Equal(state.TunnelSourceInterface) {
		hasChanges = true
	}
	if !data.TunnelDestination.Equal(state.TunnelDestination) {
		hasChanges = true
	}
	if !data.Application.Equal(state.Application) {
		hasChanges = true
	}
	if !data.TcpMssAdjust.Equal(state.TcpMssAdjust) {
		hasChanges = true
	}
	if !data.ClearDontFragment.Equal(state.ClearDontFragment) {
		hasChanges = true
	}
	if !data.Mtu.Equal(state.Mtu) {
		hasChanges = true
	}
	if !data.DeadPeerDetectionInterval.Equal(state.DeadPeerDetectionInterval) {
		hasChanges = true
	}
	if !data.DeadPeerDetectionRetries.Equal(state.DeadPeerDetectionRetries) {
		hasChanges = true
	}
	if !data.IkeVersion.Equal(state.IkeVersion) {
		hasChanges = true
	}
	if !data.IkeMode.Equal(state.IkeMode) {
		hasChanges = true
	}
	if !data.IkeRekeyInterval.Equal(state.IkeRekeyInterval) {
		hasChanges = true
	}
	if !data.IkeCiphersuite.Equal(state.IkeCiphersuite) {
		hasChanges = true
	}
	if !data.IkeGroup.Equal(state.IkeGroup) {
		hasChanges = true
	}
	if !data.IkePreSharedKey.Equal(state.IkePreSharedKey) {
		hasChanges = true
	}
	if !data.IkePreSharedKeyLocalId.Equal(state.IkePreSharedKeyLocalId) {
		hasChanges = true
	}
	if !data.IkePreSharedKeyRemoteId.Equal(state.IkePreSharedKeyRemoteId) {
		hasChanges = true
	}
	if !data.IpsecRekeyInterval.Equal(state.IpsecRekeyInterval) {
		hasChanges = true
	}
	if !data.IpsecReplayWindow.Equal(state.IpsecReplayWindow) {
		hasChanges = true
	}
	if !data.IpsecCiphersuite.Equal(state.IpsecCiphersuite) {
		hasChanges = true
	}
	if !data.IpsecPerfectForwardSecrecy.Equal(state.IpsecPerfectForwardSecrecy) {
		hasChanges = true
	}
	if !data.Tracker.Equal(state.Tracker) {
		hasChanges = true
	}
	if !data.TunnelRouteVia.Equal(state.TunnelRouteVia) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
