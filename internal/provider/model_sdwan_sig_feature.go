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

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type Sig struct {
	Id                                  types.String        `tfsdk:"id"`
	Version                             types.Int64         `tfsdk:"version"`
	Name                                types.String        `tfsdk:"name"`
	Description                         types.String        `tfsdk:"description"`
	FeatureProfileId                    types.String        `tfsdk:"feature_profile_id"`
	SigProvider                         types.String        `tfsdk:"sig_provider"`
	SrcVpn                              types.Bool          `tfsdk:"src_vpn"`
	Interfaces                          []SigInterfaces     `tfsdk:"interfaces"`
	InterfacePairs                      []SigInterfacePairs `tfsdk:"interface_pairs"`
	AuthRequired                        types.Bool          `tfsdk:"auth_required"`
	XffForwardEnabled                   types.Bool          `tfsdk:"xff_forward_enabled"`
	OfwEnabled                          types.Bool          `tfsdk:"ofw_enabled"`
	IpsControl                          types.Bool          `tfsdk:"ips_control"`
	CautionEnabled                      types.Bool          `tfsdk:"caution_enabled"`
	ZscalerPrimaryDataCenter            types.String        `tfsdk:"zscaler_primary_data_center"`
	ZscalerPrimaryDataCenterVariable    types.String        `tfsdk:"zscaler_primary_data_center_variable"`
	ZscalerSecondaryDataCenter          types.String        `tfsdk:"zscaler_secondary_data_center"`
	ZscalerSecondaryDataCenterVariable  types.String        `tfsdk:"zscaler_secondary_data_center_variable"`
	SurrogateIp                         types.Bool          `tfsdk:"surrogate_ip"`
	IdleTime                            types.Int64         `tfsdk:"idle_time"`
	DisplayTimeUnit                     types.String        `tfsdk:"display_time_unit"`
	IpEnforcedForKnownBrowsers          types.Bool          `tfsdk:"ip_enforced_for_known_browsers"`
	RefreshTime                         types.Int64         `tfsdk:"refresh_time"`
	RefreshTimeUnit                     types.String        `tfsdk:"refresh_time_unit"`
	AupEnabled                          types.Bool          `tfsdk:"aup_enabled"`
	BlockInternetUntilAccepted          types.Bool          `tfsdk:"block_internet_until_accepted"`
	ForceSslInspection                  types.Bool          `tfsdk:"force_ssl_inspection"`
	AupTimeout                          types.Int64         `tfsdk:"aup_timeout"`
	LocationName                        types.String        `tfsdk:"location_name"`
	LocationNameVariable                types.String        `tfsdk:"location_name_variable"`
	UmbrellaPrimaryDataCenter           types.String        `tfsdk:"umbrella_primary_data_center"`
	UmbrellaPrimaryDataCenterVariable   types.String        `tfsdk:"umbrella_primary_data_center_variable"`
	UmbrellaSecondaryDataCenter         types.String        `tfsdk:"umbrella_secondary_data_center"`
	UmbrellaSecondaryDataCenterVariable types.String        `tfsdk:"umbrella_secondary_data_center_variable"`
	TrackerSourceIp                     types.String        `tfsdk:"tracker_source_ip"`
	TrackerSourceIpVariable             types.String        `tfsdk:"tracker_source_ip_variable"`
	Trackers                            []SigTrackers       `tfsdk:"trackers"`
}

type SigInterfaces struct {
	InterfaceName                 types.String `tfsdk:"interface_name"`
	Shutdown                      types.Bool   `tfsdk:"shutdown"`
	InterfaceDescription          types.String `tfsdk:"interface_description"`
	InterfaceDescriptionVariable  types.String `tfsdk:"interface_description_variable"`
	TunnelSourceInterface         types.String `tfsdk:"tunnel_source_interface"`
	TunnelSourceInterfaceVariable types.String `tfsdk:"tunnel_source_interface_variable"`
	TunnelRouteVia                types.String `tfsdk:"tunnel_route_via"`
	TunnelRouteViaVariable        types.String `tfsdk:"tunnel_route_via_variable"`
	TunnelDestination             types.String `tfsdk:"tunnel_destination"`
	TunnelDestinationVariable     types.String `tfsdk:"tunnel_destination_variable"`
	TunnelSet                     types.String `tfsdk:"tunnel_set"`
	TunnelDcPreference            types.String `tfsdk:"tunnel_dc_preference"`
	TcpMssAdjust                  types.Int64  `tfsdk:"tcp_mss_adjust"`
	TcpMssAdjustVariable          types.String `tfsdk:"tcp_mss_adjust_variable"`
	Mtu                           types.Int64  `tfsdk:"mtu"`
	MtuVariable                   types.String `tfsdk:"mtu_variable"`
	DpdInterval                   types.Int64  `tfsdk:"dpd_interval"`
	DpdIntervalVariable           types.String `tfsdk:"dpd_interval_variable"`
	DpdRetries                    types.Int64  `tfsdk:"dpd_retries"`
	DpdRetriesVariable            types.String `tfsdk:"dpd_retries_variable"`
	PreSharedSecret               types.String `tfsdk:"pre_shared_secret"`
	PreSharedSecretVariable       types.String `tfsdk:"pre_shared_secret_variable"`
	IkeRekeyInterval              types.Int64  `tfsdk:"ike_rekey_interval"`
	IkeRekeyIntervalVariable      types.String `tfsdk:"ike_rekey_interval_variable"`
	IkeCiphersuite                types.String `tfsdk:"ike_ciphersuite"`
	IkeCiphersuiteVariable        types.String `tfsdk:"ike_ciphersuite_variable"`
	IkeGroup                      types.String `tfsdk:"ike_group"`
	IkeGroupVariable              types.String `tfsdk:"ike_group_variable"`
	PreSharedKeyDynamic           types.Bool   `tfsdk:"pre_shared_key_dynamic"`
	IkeLocalId                    types.String `tfsdk:"ike_local_id"`
	IkeLocalIdVariable            types.String `tfsdk:"ike_local_id_variable"`
	IkeRemoteId                   types.String `tfsdk:"ike_remote_id"`
	IkeRemoteIdVariable           types.String `tfsdk:"ike_remote_id_variable"`
	IpsecRekeyInterval            types.Int64  `tfsdk:"ipsec_rekey_interval"`
	IpsecRekeyIntervalVariable    types.String `tfsdk:"ipsec_rekey_interval_variable"`
	IpsecReplayWindow             types.Int64  `tfsdk:"ipsec_replay_window"`
	IpsecReplayWindowVariable     types.String `tfsdk:"ipsec_replay_window_variable"`
	IpsecCiphersuite              types.String `tfsdk:"ipsec_ciphersuite"`
	IpsecCiphersuiteVariable      types.String `tfsdk:"ipsec_ciphersuite_variable"`
	PerfectForwardSecrecy         types.String `tfsdk:"perfect_forward_secrecy"`
	PerfectForwardSecrecyVariable types.String `tfsdk:"perfect_forward_secrecy_variable"`
	Tracker                       types.String `tfsdk:"tracker"`
	TrackEnable                   types.Bool   `tfsdk:"track_enable"`
	TunnelPublicIp                types.String `tfsdk:"tunnel_public_ip"`
	TunnelPublicIpVariable        types.String `tfsdk:"tunnel_public_ip_variable"`
}

type SigInterfacePairs struct {
	ActiveInterface       types.String `tfsdk:"active_interface"`
	ActiveInterfaceWeight types.Int64  `tfsdk:"active_interface_weight"`
	BackupInterface       types.String `tfsdk:"backup_interface"`
	BackupInterfaceWeight types.Int64  `tfsdk:"backup_interface_weight"`
}

type SigTrackers struct {
	Name                   types.String `tfsdk:"name"`
	EndpointApiUrl         types.String `tfsdk:"endpoint_api_url"`
	EndpointApiUrlVariable types.String `tfsdk:"endpoint_api_url_variable"`
	Threshold              types.Int64  `tfsdk:"threshold"`
	ThresholdVariable      types.String `tfsdk:"threshold_variable"`
	Interval               types.Int64  `tfsdk:"interval"`
	IntervalVariable       types.String `tfsdk:"interval_variable"`
	Multiplier             types.Int64  `tfsdk:"multiplier"`
	MultiplierVariable     types.String `tfsdk:"multiplier_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data Sig) getModel() string {
	return "sig_security_sig"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Sig) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/sig-security/%v/sig", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Sig) toBody(ctx context.Context, ver *version.Version) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.SigProvider.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sigProvider.optionType", "global")
			body, _ = sjson.Set(body, path+"sigProvider.value", data.SigProvider.ValueString())
		}
	}
	if data.SrcVpn.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"interfaceMetadataSharing.srcVpn.optionType", "default")
			body, _ = sjson.Set(body, path+"interfaceMetadataSharing.srcVpn.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"interfaceMetadataSharing.srcVpn.optionType", "global")
			body, _ = sjson.Set(body, path+"interfaceMetadataSharing.srcVpn.value", data.SrcVpn.ValueBool())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"interface", []interface{}{})
		for _, item := range data.Interfaces {
			itemBody := ""
			if !item.InterfaceName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ifName.value", item.InterfaceName.ValueString())
				}
			}
			if item.Shutdown.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "shutdown.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "shutdown.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "shutdown.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "shutdown.value", item.Shutdown.ValueBool())
				}
			}

			if !item.InterfaceDescriptionVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "description.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "description.value", item.InterfaceDescriptionVariable.ValueString())
				}
			} else if item.InterfaceDescription.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "description.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "description.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "description.value", item.InterfaceDescription.ValueString())
				}
			}

			if !item.TunnelSourceInterfaceVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tunnelSourceInterface.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "tunnelSourceInterface.value", item.TunnelSourceInterfaceVariable.ValueString())
				}
			} else if !item.TunnelSourceInterface.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tunnelSourceInterface.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tunnelSourceInterface.value", item.TunnelSourceInterface.ValueString())
				}
			}

			if !item.TunnelRouteViaVariable.IsNull() {
				if true && ver.GreaterThanOrEqual(version.Must(version.NewVersion("20.18.1"))) {
					itemBody, _ = sjson.Set(itemBody, "tunnelRouteVia.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "tunnelRouteVia.value", item.TunnelRouteViaVariable.ValueString())
				}
			} else if item.TunnelRouteVia.IsNull() {
				if true && ver.GreaterThanOrEqual(version.Must(version.NewVersion("20.18.1"))) {
					itemBody, _ = sjson.Set(itemBody, "tunnelRouteVia.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "tunnelRouteVia.value", "")
				}
			} else {
				if true && ver.GreaterThanOrEqual(version.Must(version.NewVersion("20.18.1"))) {
					itemBody, _ = sjson.Set(itemBody, "tunnelRouteVia.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tunnelRouteVia.value", item.TunnelRouteVia.ValueString())
				}
			}

			if !item.TunnelDestinationVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tunnelDestination.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "tunnelDestination.value", item.TunnelDestinationVariable.ValueString())
				}
			} else if !item.TunnelDestination.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tunnelDestination.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tunnelDestination.value", item.TunnelDestination.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "application.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "application.value", "sig")
			}
			if !item.TunnelSet.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tunnelSet.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tunnelSet.value", item.TunnelSet.ValueString())
				}
			}
			if !item.TunnelDcPreference.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tunnelDcPreference.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tunnelDcPreference.value", item.TunnelDcPreference.ValueString())
				}
			}

			if !item.TcpMssAdjustVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tcpMssAdjust.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "tcpMssAdjust.value", item.TcpMssAdjustVariable.ValueString())
				}
			} else if item.TcpMssAdjust.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tcpMssAdjust.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tcpMssAdjust.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tcpMssAdjust.value", item.TcpMssAdjust.ValueInt64())
				}
			}

			if !item.MtuVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "mtu.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "mtu.value", item.MtuVariable.ValueString())
				}
			} else if !item.Mtu.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "mtu.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "mtu.value", item.Mtu.ValueInt64())
				}
			}

			if !item.DpdIntervalVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dpdInterval.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "dpdInterval.value", item.DpdIntervalVariable.ValueString())
				}
			} else if item.DpdInterval.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dpdInterval.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "dpdInterval.value", 10)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dpdInterval.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "dpdInterval.value", item.DpdInterval.ValueInt64())
				}
			}

			if !item.DpdRetriesVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dpdRetries.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "dpdRetries.value", item.DpdRetriesVariable.ValueString())
				}
			} else if item.DpdRetries.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dpdRetries.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "dpdRetries.value", 3)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dpdRetries.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "dpdRetries.value", item.DpdRetries.ValueInt64())
				}
			}

			if !item.PreSharedSecretVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "preSharedSecret.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "preSharedSecret.value", item.PreSharedSecretVariable.ValueString())
				}
			} else if item.PreSharedSecret.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "preSharedSecret.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "preSharedSecret.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "preSharedSecret.value", item.PreSharedSecret.ValueString())
				}
			}

			if !item.IkeRekeyIntervalVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeRekeyInterval.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ikeRekeyInterval.value", item.IkeRekeyIntervalVariable.ValueString())
				}
			} else if item.IkeRekeyInterval.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeRekeyInterval.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "ikeRekeyInterval.value", 14400)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeRekeyInterval.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ikeRekeyInterval.value", item.IkeRekeyInterval.ValueInt64())
				}
			}

			if !item.IkeCiphersuiteVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeCiphersuite.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ikeCiphersuite.value", item.IkeCiphersuiteVariable.ValueString())
				}
			} else if item.IkeCiphersuite.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeCiphersuite.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "ikeCiphersuite.value", "aes256-cbc-sha1")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeCiphersuite.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ikeCiphersuite.value", item.IkeCiphersuite.ValueString())
				}
			}

			if !item.IkeGroupVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeGroup.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ikeGroup.value", item.IkeGroupVariable.ValueString())
				}
			} else if item.IkeGroup.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeGroup.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "ikeGroup.value", "16")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeGroup.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ikeGroup.value", item.IkeGroup.ValueString())
				}
			}
			if !item.PreSharedKeyDynamic.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "preSharedKeyDynamic.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "preSharedKeyDynamic.value", item.PreSharedKeyDynamic.ValueBool())
				}
			}

			if !item.IkeLocalIdVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeLocalId.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ikeLocalId.value", item.IkeLocalIdVariable.ValueString())
				}
			} else if item.IkeLocalId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeLocalId.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeLocalId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ikeLocalId.value", item.IkeLocalId.ValueString())
				}
			}

			if !item.IkeRemoteIdVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeRemoteId.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ikeRemoteId.value", item.IkeRemoteIdVariable.ValueString())
				}
			} else if item.IkeRemoteId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeRemoteId.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeRemoteId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ikeRemoteId.value", item.IkeRemoteId.ValueString())
				}
			}

			if !item.IpsecRekeyIntervalVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipsecRekeyInterval.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ipsecRekeyInterval.value", item.IpsecRekeyIntervalVariable.ValueString())
				}
			} else if item.IpsecRekeyInterval.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipsecRekeyInterval.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "ipsecRekeyInterval.value", 3600)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipsecRekeyInterval.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ipsecRekeyInterval.value", item.IpsecRekeyInterval.ValueInt64())
				}
			}

			if !item.IpsecReplayWindowVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipsecReplayWindow.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ipsecReplayWindow.value", item.IpsecReplayWindowVariable.ValueString())
				}
			} else if item.IpsecReplayWindow.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipsecReplayWindow.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "ipsecReplayWindow.value", 512)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipsecReplayWindow.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ipsecReplayWindow.value", item.IpsecReplayWindow.ValueInt64())
				}
			}

			if !item.IpsecCiphersuiteVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipsecCiphersuite.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ipsecCiphersuite.value", item.IpsecCiphersuiteVariable.ValueString())
				}
			} else if item.IpsecCiphersuite.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipsecCiphersuite.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "ipsecCiphersuite.value", "aes256-cbc-sha512")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipsecCiphersuite.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ipsecCiphersuite.value", item.IpsecCiphersuite.ValueString())
				}
			}

			if !item.PerfectForwardSecrecyVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "perfectForwardSecrecy.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "perfectForwardSecrecy.value", item.PerfectForwardSecrecyVariable.ValueString())
				}
			} else if item.PerfectForwardSecrecy.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "perfectForwardSecrecy.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "perfectForwardSecrecy.value", "none")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "perfectForwardSecrecy.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "perfectForwardSecrecy.value", item.PerfectForwardSecrecy.ValueString())
				}
			}
			if item.Tracker.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tracker.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tracker.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tracker.value", item.Tracker.ValueString())
				}
			}
			if item.TrackEnable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackEnable.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "trackEnable.value", true)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackEnable.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "trackEnable.value", item.TrackEnable.ValueBool())
				}
			}

			if !item.TunnelPublicIpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tunnelPublicIp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "tunnelPublicIp.value", item.TunnelPublicIpVariable.ValueString())
				}
			} else if item.TunnelPublicIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tunnelPublicIp.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "tunnelPublicIp.value", "Auto")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tunnelPublicIp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tunnelPublicIp.value", item.TunnelPublicIp.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"interface.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"service.interfacePair", []interface{}{})
		for _, item := range data.InterfacePairs {
			itemBody := ""
			if !item.ActiveInterface.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "activeInterface.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "activeInterface.value", item.ActiveInterface.ValueString())
				}
			}
			if !item.ActiveInterfaceWeight.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "activeInterfaceWeight.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "activeInterfaceWeight.value", item.ActiveInterfaceWeight.ValueInt64())
				}
			}
			if !item.BackupInterface.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "backupInterface.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "backupInterface.value", item.BackupInterface.ValueString())
				}
			}
			if !item.BackupInterfaceWeight.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "backupInterfaceWeight.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "backupInterfaceWeight.value", item.BackupInterfaceWeight.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"service.interfacePair.-1", itemBody)
		}
	}
	if data.AuthRequired.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.authRequired.optionType", "default")
			body, _ = sjson.Set(body, path+"service.authRequired.value", false)
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.authRequired.optionType", "global")
			body, _ = sjson.Set(body, path+"service.authRequired.value", data.AuthRequired.ValueBool())
		}
	}
	if data.XffForwardEnabled.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.xffForwardEnabled.optionType", "default")
			body, _ = sjson.Set(body, path+"service.xffForwardEnabled.value", false)
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.xffForwardEnabled.optionType", "global")
			body, _ = sjson.Set(body, path+"service.xffForwardEnabled.value", data.XffForwardEnabled.ValueBool())
		}
	}
	if data.OfwEnabled.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.ofwEnabled.optionType", "default")
			body, _ = sjson.Set(body, path+"service.ofwEnabled.value", false)
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.ofwEnabled.optionType", "global")
			body, _ = sjson.Set(body, path+"service.ofwEnabled.value", data.OfwEnabled.ValueBool())
		}
	}
	if data.IpsControl.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.ipsControl.optionType", "default")
			body, _ = sjson.Set(body, path+"service.ipsControl.value", false)
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.ipsControl.optionType", "global")
			body, _ = sjson.Set(body, path+"service.ipsControl.value", data.IpsControl.ValueBool())
		}
	}
	if data.CautionEnabled.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.AuthRequired.ValueBool() == false {
			body, _ = sjson.Set(body, path+"service.cautionEnabled.optionType", "default")
			body, _ = sjson.Set(body, path+"service.cautionEnabled.value", false)
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.AuthRequired.ValueBool() == false {
			body, _ = sjson.Set(body, path+"service.cautionEnabled.optionType", "global")
			body, _ = sjson.Set(body, path+"service.cautionEnabled.value", data.CautionEnabled.ValueBool())
		}
	}

	if !data.ZscalerPrimaryDataCenterVariable.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.primaryDataCenter.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.primaryDataCenter.value", data.ZscalerPrimaryDataCenterVariable.ValueString())
		}
	} else if data.ZscalerPrimaryDataCenter.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.primaryDataCenter.optionType", "default")
			body, _ = sjson.Set(body, path+"service.primaryDataCenter.value", "Auto")
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.primaryDataCenter.optionType", "global")
			body, _ = sjson.Set(body, path+"service.primaryDataCenter.value", data.ZscalerPrimaryDataCenter.ValueString())
		}
	}

	if !data.ZscalerSecondaryDataCenterVariable.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.secondaryDataCenter.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.secondaryDataCenter.value", data.ZscalerSecondaryDataCenterVariable.ValueString())
		}
	} else if data.ZscalerSecondaryDataCenter.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.secondaryDataCenter.optionType", "default")
			body, _ = sjson.Set(body, path+"service.secondaryDataCenter.value", "Auto")
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.secondaryDataCenter.optionType", "global")
			body, _ = sjson.Set(body, path+"service.secondaryDataCenter.value", data.ZscalerSecondaryDataCenter.ValueString())
		}
	}
	if data.SurrogateIp.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.AuthRequired.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.ip.optionType", "default")
			body, _ = sjson.Set(body, path+"service.ip.value", false)
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.AuthRequired.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.ip.optionType", "global")
			body, _ = sjson.Set(body, path+"service.ip.value", data.SurrogateIp.ValueBool())
		}
	}
	if data.IdleTime.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.SurrogateIp.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.idleTime.optionType", "default")
			body, _ = sjson.Set(body, path+"service.idleTime.value", 0)
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.SurrogateIp.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.idleTime.optionType", "global")
			body, _ = sjson.Set(body, path+"service.idleTime.value", data.IdleTime.ValueInt64())
		}
	}
	if data.DisplayTimeUnit.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.SurrogateIp.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.displayTimeUnit.optionType", "default")
			body, _ = sjson.Set(body, path+"service.displayTimeUnit.value", "MINUTE")
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.SurrogateIp.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.displayTimeUnit.optionType", "global")
			body, _ = sjson.Set(body, path+"service.displayTimeUnit.value", data.DisplayTimeUnit.ValueString())
		}
	}
	if data.IpEnforcedForKnownBrowsers.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.SurrogateIp.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.ipEnforcedForKnownBrowsers.optionType", "default")
			body, _ = sjson.Set(body, path+"service.ipEnforcedForKnownBrowsers.value", false)
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.SurrogateIp.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.ipEnforcedForKnownBrowsers.optionType", "global")
			body, _ = sjson.Set(body, path+"service.ipEnforcedForKnownBrowsers.value", data.IpEnforcedForKnownBrowsers.ValueBool())
		}
	}
	if data.RefreshTime.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.IpEnforcedForKnownBrowsers.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.refreshTime.optionType", "default")
			body, _ = sjson.Set(body, path+"service.refreshTime.value", 0)
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.IpEnforcedForKnownBrowsers.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.refreshTime.optionType", "global")
			body, _ = sjson.Set(body, path+"service.refreshTime.value", data.RefreshTime.ValueInt64())
		}
	}
	if data.RefreshTimeUnit.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.IpEnforcedForKnownBrowsers.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.refreshTimeUnit.optionType", "default")
			body, _ = sjson.Set(body, path+"service.refreshTimeUnit.value", "MINUTE")
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.IpEnforcedForKnownBrowsers.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.refreshTimeUnit.optionType", "global")
			body, _ = sjson.Set(body, path+"service.refreshTimeUnit.value", data.RefreshTimeUnit.ValueString())
		}
	}
	if data.AupEnabled.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.AuthRequired.ValueBool() == false {
			body, _ = sjson.Set(body, path+"service.enabled.optionType", "default")
			body, _ = sjson.Set(body, path+"service.enabled.value", false)
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.AuthRequired.ValueBool() == false {
			body, _ = sjson.Set(body, path+"service.enabled.optionType", "global")
			body, _ = sjson.Set(body, path+"service.enabled.value", data.AupEnabled.ValueBool())
		}
	}
	if data.BlockInternetUntilAccepted.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.AupEnabled.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.blockInternetUntilAccepted.optionType", "default")
			body, _ = sjson.Set(body, path+"service.blockInternetUntilAccepted.value", false)
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.AupEnabled.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.blockInternetUntilAccepted.optionType", "global")
			body, _ = sjson.Set(body, path+"service.blockInternetUntilAccepted.value", data.BlockInternetUntilAccepted.ValueBool())
		}
	}
	if data.ForceSslInspection.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.AupEnabled.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.forceSslInspection.optionType", "default")
			body, _ = sjson.Set(body, path+"service.forceSslInspection.value", false)
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.AupEnabled.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.forceSslInspection.optionType", "global")
			body, _ = sjson.Set(body, path+"service.forceSslInspection.value", data.ForceSslInspection.ValueBool())
		}
	}
	if data.AupTimeout.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.AupEnabled.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.timeout.optionType", "default")
			body, _ = sjson.Set(body, path+"service.timeout.value", 0)
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" && data.AupEnabled.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.timeout.optionType", "global")
			body, _ = sjson.Set(body, path+"service.timeout.value", data.AupTimeout.ValueInt64())
		}
	}

	if !data.LocationNameVariable.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.locationName.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.locationName.value", data.LocationNameVariable.ValueString())
		}
	} else if data.LocationName.IsNull() {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.locationName.optionType", "default")
			body, _ = sjson.Set(body, path+"service.locationName.value", "Auto")
		}
	} else {
		if true && data.SigProvider.ValueString() == "Zscaler" {
			body, _ = sjson.Set(body, path+"service.locationName.optionType", "default")
			body, _ = sjson.Set(body, path+"service.locationName.value", data.LocationName.ValueString())
		}
	}

	if !data.UmbrellaPrimaryDataCenterVariable.IsNull() {
		if true && data.SigProvider.ValueString() == "Umbrella" {
			body, _ = sjson.Set(body, path+"service.dataCenterPrimary.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.dataCenterPrimary.value", data.UmbrellaPrimaryDataCenterVariable.ValueString())
		}
	} else if data.UmbrellaPrimaryDataCenter.IsNull() {
		if true && data.SigProvider.ValueString() == "Umbrella" {
			body, _ = sjson.Set(body, path+"service.dataCenterPrimary.optionType", "default")
			body, _ = sjson.Set(body, path+"service.dataCenterPrimary.value", "Auto")
		}
	} else {
		if true && data.SigProvider.ValueString() == "Umbrella" {
			body, _ = sjson.Set(body, path+"service.dataCenterPrimary.optionType", "global")
			body, _ = sjson.Set(body, path+"service.dataCenterPrimary.value", data.UmbrellaPrimaryDataCenter.ValueString())
		}
	}

	if !data.UmbrellaSecondaryDataCenterVariable.IsNull() {
		if true && data.SigProvider.ValueString() == "Umbrella" {
			body, _ = sjson.Set(body, path+"service.dataCenterSecondary.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.dataCenterSecondary.value", data.UmbrellaSecondaryDataCenterVariable.ValueString())
		}
	} else if data.UmbrellaSecondaryDataCenter.IsNull() {
		if true && data.SigProvider.ValueString() == "Umbrella" {
			body, _ = sjson.Set(body, path+"service.dataCenterSecondary.optionType", "default")
			body, _ = sjson.Set(body, path+"service.dataCenterSecondary.value", "Auto")
		}
	} else {
		if true && data.SigProvider.ValueString() == "Umbrella" {
			body, _ = sjson.Set(body, path+"service.dataCenterSecondary.optionType", "global")
			body, _ = sjson.Set(body, path+"service.dataCenterSecondary.value", data.UmbrellaSecondaryDataCenter.ValueString())
		}
	}

	if !data.TrackerSourceIpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trackerSrcIp.optionType", "variable")
			body, _ = sjson.Set(body, path+"trackerSrcIp.value", data.TrackerSourceIpVariable.ValueString())
		}
	} else if !data.TrackerSourceIp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trackerSrcIp.optionType", "global")
			body, _ = sjson.Set(body, path+"trackerSrcIp.value", data.TrackerSourceIp.ValueString())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"tracker", []interface{}{})
		for _, item := range data.Trackers {
			itemBody := ""
			if !item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
				}
			}

			if !item.EndpointApiUrlVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "endpointApiUrl.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "endpointApiUrl.value", item.EndpointApiUrlVariable.ValueString())
				}
			} else if !item.EndpointApiUrl.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "endpointApiUrl.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "endpointApiUrl.value", item.EndpointApiUrl.ValueString())
				}
			}

			if !item.ThresholdVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "threshold.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "threshold.value", item.ThresholdVariable.ValueString())
				}
			} else if item.Threshold.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "threshold.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "threshold.value", 300)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "threshold.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "threshold.value", item.Threshold.ValueInt64())
				}
			}

			if !item.IntervalVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interval.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "interval.value", item.IntervalVariable.ValueString())
				}
			} else if item.Interval.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interval.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "interval.value", 60)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interval.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "interval.value", item.Interval.ValueInt64())
				}
			}

			if !item.MultiplierVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "multiplier.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "multiplier.value", item.MultiplierVariable.ValueString())
				}
			} else if item.Multiplier.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "multiplier.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "multiplier.value", 3)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "multiplier.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "multiplier.value", item.Multiplier.ValueInt64())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "trackerType.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "trackerType.value", "SIG")
			}
			body, _ = sjson.SetRaw(body, path+"tracker.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Sig) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.SigProvider = types.StringNull()

	if t := res.Get(path + "sigProvider.optionType"); t.Exists() {
		va := res.Get(path + "sigProvider.value")
		if t.String() == "global" {
			data.SigProvider = types.StringValue(va.String())
		}
	}
	data.SrcVpn = types.BoolNull()

	if t := res.Get(path + "interfaceMetadataSharing.srcVpn.optionType"); t.Exists() {
		va := res.Get(path + "interfaceMetadataSharing.srcVpn.value")
		if t.String() == "global" {
			data.SrcVpn = types.BoolValue(va.Bool())
		}
	}
	oldInterfaces := data.Interfaces
	if value := res.Get(path + "interface"); value.Exists() && len(value.Array()) > 0 {
		data.Interfaces = make([]SigInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SigInterfaces{}
			item.InterfaceName = types.StringNull()

			if t := v.Get("ifName.optionType"); t.Exists() {
				va := v.Get("ifName.value")
				if t.String() == "global" {
					item.InterfaceName = types.StringValue(va.String())
				}
			}
			item.Shutdown = types.BoolNull()

			if t := v.Get("shutdown.optionType"); t.Exists() {
				va := v.Get("shutdown.value")
				if t.String() == "global" {
					item.Shutdown = types.BoolValue(va.Bool())
				}
			}
			item.InterfaceDescription = types.StringNull()
			item.InterfaceDescriptionVariable = types.StringNull()
			if t := v.Get("description.optionType"); t.Exists() {
				va := v.Get("description.value")
				if t.String() == "variable" {
					item.InterfaceDescriptionVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.InterfaceDescription = types.StringValue(va.String())
				}
			}
			item.TunnelSourceInterface = types.StringNull()
			item.TunnelSourceInterfaceVariable = types.StringNull()
			if t := v.Get("tunnelSourceInterface.optionType"); t.Exists() {
				va := v.Get("tunnelSourceInterface.value")
				if t.String() == "variable" {
					item.TunnelSourceInterfaceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TunnelSourceInterface = types.StringValue(va.String())
				}
			}
			item.TunnelRouteVia = types.StringNull()
			item.TunnelRouteViaVariable = types.StringNull()
			if t := v.Get("tunnelRouteVia.optionType"); t.Exists() {
				va := v.Get("tunnelRouteVia.value")
				if t.String() == "variable" {
					item.TunnelRouteViaVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TunnelRouteVia = types.StringValue(va.String())
				}
			}
			item.TunnelDestination = types.StringNull()
			item.TunnelDestinationVariable = types.StringNull()
			if t := v.Get("tunnelDestination.optionType"); t.Exists() {
				va := v.Get("tunnelDestination.value")
				if t.String() == "variable" {
					item.TunnelDestinationVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TunnelDestination = types.StringValue(va.String())
				}
			}
			item.TunnelSet = types.StringNull()

			if t := v.Get("tunnelSet.optionType"); t.Exists() {
				va := v.Get("tunnelSet.value")
				if t.String() == "global" {
					item.TunnelSet = types.StringValue(va.String())
				}
			}
			item.TunnelDcPreference = types.StringNull()

			if t := v.Get("tunnelDcPreference.optionType"); t.Exists() {
				va := v.Get("tunnelDcPreference.value")
				if t.String() == "global" {
					item.TunnelDcPreference = types.StringValue(va.String())
				}
			}
			item.TcpMssAdjust = types.Int64Null()
			item.TcpMssAdjustVariable = types.StringNull()
			if t := v.Get("tcpMssAdjust.optionType"); t.Exists() {
				va := v.Get("tcpMssAdjust.value")
				if t.String() == "variable" {
					item.TcpMssAdjustVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TcpMssAdjust = types.Int64Value(va.Int())
				}
			}
			item.Mtu = types.Int64Null()
			item.MtuVariable = types.StringNull()
			if t := v.Get("mtu.optionType"); t.Exists() {
				va := v.Get("mtu.value")
				if t.String() == "variable" {
					item.MtuVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Mtu = types.Int64Value(va.Int())
				}
			}
			item.DpdInterval = types.Int64Null()
			item.DpdIntervalVariable = types.StringNull()
			if t := v.Get("dpdInterval.optionType"); t.Exists() {
				va := v.Get("dpdInterval.value")
				if t.String() == "variable" {
					item.DpdIntervalVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.DpdInterval = types.Int64Value(va.Int())
				}
			}
			item.DpdRetries = types.Int64Null()
			item.DpdRetriesVariable = types.StringNull()
			if t := v.Get("dpdRetries.optionType"); t.Exists() {
				va := v.Get("dpdRetries.value")
				if t.String() == "variable" {
					item.DpdRetriesVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.DpdRetries = types.Int64Value(va.Int())
				}
			}
			item.PreSharedSecret = types.StringNull()
			item.PreSharedSecretVariable = types.StringNull()
			if t := v.Get("preSharedSecret.optionType"); t.Exists() {
				va := v.Get("preSharedSecret.value")
				if t.String() == "variable" {
					item.PreSharedSecretVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.PreSharedSecret = types.StringValue(va.String())
				}
			}
			item.IkeRekeyInterval = types.Int64Null()
			item.IkeRekeyIntervalVariable = types.StringNull()
			if t := v.Get("ikeRekeyInterval.optionType"); t.Exists() {
				va := v.Get("ikeRekeyInterval.value")
				if t.String() == "variable" {
					item.IkeRekeyIntervalVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IkeRekeyInterval = types.Int64Value(va.Int())
				}
			}
			item.IkeCiphersuite = types.StringNull()
			item.IkeCiphersuiteVariable = types.StringNull()
			if t := v.Get("ikeCiphersuite.optionType"); t.Exists() {
				va := v.Get("ikeCiphersuite.value")
				if t.String() == "variable" {
					item.IkeCiphersuiteVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IkeCiphersuite = types.StringValue(va.String())
				}
			}
			item.IkeGroup = types.StringNull()
			item.IkeGroupVariable = types.StringNull()
			if t := v.Get("ikeGroup.optionType"); t.Exists() {
				va := v.Get("ikeGroup.value")
				if t.String() == "variable" {
					item.IkeGroupVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IkeGroup = types.StringValue(va.String())
				}
			}
			item.PreSharedKeyDynamic = types.BoolNull()

			if t := v.Get("preSharedKeyDynamic.optionType"); t.Exists() {
				va := v.Get("preSharedKeyDynamic.value")
				if t.String() == "global" {
					item.PreSharedKeyDynamic = types.BoolValue(va.Bool())
				}
			}
			item.IkeLocalId = types.StringNull()
			item.IkeLocalIdVariable = types.StringNull()
			if t := v.Get("ikeLocalId.optionType"); t.Exists() {
				va := v.Get("ikeLocalId.value")
				if t.String() == "variable" {
					item.IkeLocalIdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IkeLocalId = types.StringValue(va.String())
				}
			}
			item.IkeRemoteId = types.StringNull()
			item.IkeRemoteIdVariable = types.StringNull()
			if t := v.Get("ikeRemoteId.optionType"); t.Exists() {
				va := v.Get("ikeRemoteId.value")
				if t.String() == "variable" {
					item.IkeRemoteIdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IkeRemoteId = types.StringValue(va.String())
				}
			}
			item.IpsecRekeyInterval = types.Int64Null()
			item.IpsecRekeyIntervalVariable = types.StringNull()
			if t := v.Get("ipsecRekeyInterval.optionType"); t.Exists() {
				va := v.Get("ipsecRekeyInterval.value")
				if t.String() == "variable" {
					item.IpsecRekeyIntervalVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IpsecRekeyInterval = types.Int64Value(va.Int())
				}
			}
			item.IpsecReplayWindow = types.Int64Null()
			item.IpsecReplayWindowVariable = types.StringNull()
			if t := v.Get("ipsecReplayWindow.optionType"); t.Exists() {
				va := v.Get("ipsecReplayWindow.value")
				if t.String() == "variable" {
					item.IpsecReplayWindowVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IpsecReplayWindow = types.Int64Value(va.Int())
				}
			}
			item.IpsecCiphersuite = types.StringNull()
			item.IpsecCiphersuiteVariable = types.StringNull()
			if t := v.Get("ipsecCiphersuite.optionType"); t.Exists() {
				va := v.Get("ipsecCiphersuite.value")
				if t.String() == "variable" {
					item.IpsecCiphersuiteVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IpsecCiphersuite = types.StringValue(va.String())
				}
			}
			item.PerfectForwardSecrecy = types.StringNull()
			item.PerfectForwardSecrecyVariable = types.StringNull()
			if t := v.Get("perfectForwardSecrecy.optionType"); t.Exists() {
				va := v.Get("perfectForwardSecrecy.value")
				if t.String() == "variable" {
					item.PerfectForwardSecrecyVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.PerfectForwardSecrecy = types.StringValue(va.String())
				}
			}
			item.Tracker = types.StringNull()

			if t := v.Get("tracker.optionType"); t.Exists() {
				va := v.Get("tracker.value")
				if t.String() == "global" {
					item.Tracker = types.StringValue(va.String())
				}
			}
			item.TrackEnable = types.BoolNull()

			if t := v.Get("trackEnable.optionType"); t.Exists() {
				va := v.Get("trackEnable.value")
				if t.String() == "global" {
					item.TrackEnable = types.BoolValue(va.Bool())
				}
			}
			item.TunnelPublicIp = types.StringNull()
			item.TunnelPublicIpVariable = types.StringNull()
			if t := v.Get("tunnelPublicIp.optionType"); t.Exists() {
				va := v.Get("tunnelPublicIp.value")
				if t.String() == "variable" {
					item.TunnelPublicIpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TunnelPublicIp = types.StringValue(va.String())
				}
			}
			data.Interfaces = append(data.Interfaces, item)
			return true
		})
	} else {
		data.Interfaces = nil
	}
	if !fullRead && data.Interfaces != nil {
		resultInterfaces := make([]SigInterfaces, 0, len(data.Interfaces))
		matchedInterfaces := make([]bool, len(data.Interfaces))
		for _, oldItem := range oldInterfaces {
			for ni := range data.Interfaces {
				if matchedInterfaces[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.InterfaceName.ValueString() != data.Interfaces[ni].InterfaceName.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedInterfaces[ni] = true
					resultInterfaces = append(resultInterfaces, data.Interfaces[ni])
					break
				}
			}
		}
		for ni := range data.Interfaces {
			if !matchedInterfaces[ni] {
				resultInterfaces = append(resultInterfaces, data.Interfaces[ni])
			}
		}
		data.Interfaces = resultInterfaces
	}
	oldInterfacePairs := data.InterfacePairs
	if value := res.Get(path + "service.interfacePair"); value.Exists() && len(value.Array()) > 0 {
		data.InterfacePairs = make([]SigInterfacePairs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SigInterfacePairs{}
			item.ActiveInterface = types.StringNull()

			if t := v.Get("activeInterface.optionType"); t.Exists() {
				va := v.Get("activeInterface.value")
				if t.String() == "global" {
					item.ActiveInterface = types.StringValue(va.String())
				}
			}
			item.ActiveInterfaceWeight = types.Int64Null()

			if t := v.Get("activeInterfaceWeight.optionType"); t.Exists() {
				va := v.Get("activeInterfaceWeight.value")
				if t.String() == "global" {
					item.ActiveInterfaceWeight = types.Int64Value(va.Int())
				}
			}
			item.BackupInterface = types.StringNull()

			if t := v.Get("backupInterface.optionType"); t.Exists() {
				va := v.Get("backupInterface.value")
				if t.String() == "global" {
					item.BackupInterface = types.StringValue(va.String())
				}
			}
			item.BackupInterfaceWeight = types.Int64Null()

			if t := v.Get("backupInterfaceWeight.optionType"); t.Exists() {
				va := v.Get("backupInterfaceWeight.value")
				if t.String() == "global" {
					item.BackupInterfaceWeight = types.Int64Value(va.Int())
				}
			}
			data.InterfacePairs = append(data.InterfacePairs, item)
			return true
		})
	} else {
		data.InterfacePairs = nil
	}
	if !fullRead && data.InterfacePairs != nil {
		resultInterfacePairs := make([]SigInterfacePairs, 0, len(data.InterfacePairs))
		matchedInterfacePairs := make([]bool, len(data.InterfacePairs))
		for _, oldItem := range oldInterfacePairs {
			for ni := range data.InterfacePairs {
				if matchedInterfacePairs[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.ActiveInterface.ValueString() != data.InterfacePairs[ni].ActiveInterface.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if oldItem.ActiveInterfaceWeight.ValueInt64() != data.InterfacePairs[ni].ActiveInterfaceWeight.ValueInt64() {
						keyMatch = false
					}
				}
				if keyMatch {
					if oldItem.BackupInterface.ValueString() != data.InterfacePairs[ni].BackupInterface.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if oldItem.BackupInterfaceWeight.ValueInt64() != data.InterfacePairs[ni].BackupInterfaceWeight.ValueInt64() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedInterfacePairs[ni] = true
					resultInterfacePairs = append(resultInterfacePairs, data.InterfacePairs[ni])
					break
				}
			}
		}
		for ni := range data.InterfacePairs {
			if !matchedInterfacePairs[ni] {
				resultInterfacePairs = append(resultInterfacePairs, data.InterfacePairs[ni])
			}
		}
		data.InterfacePairs = resultInterfacePairs
	}
	data.AuthRequired = types.BoolNull()

	if t := res.Get(path + "service.authRequired.optionType"); t.Exists() {
		va := res.Get(path + "service.authRequired.value")
		if t.String() == "global" {
			data.AuthRequired = types.BoolValue(va.Bool())
		}
	}
	data.XffForwardEnabled = types.BoolNull()

	if t := res.Get(path + "service.xffForwardEnabled.optionType"); t.Exists() {
		va := res.Get(path + "service.xffForwardEnabled.value")
		if t.String() == "global" {
			data.XffForwardEnabled = types.BoolValue(va.Bool())
		}
	}
	data.OfwEnabled = types.BoolNull()

	if t := res.Get(path + "service.ofwEnabled.optionType"); t.Exists() {
		va := res.Get(path + "service.ofwEnabled.value")
		if t.String() == "global" {
			data.OfwEnabled = types.BoolValue(va.Bool())
		}
	}
	data.IpsControl = types.BoolNull()

	if t := res.Get(path + "service.ipsControl.optionType"); t.Exists() {
		va := res.Get(path + "service.ipsControl.value")
		if t.String() == "global" {
			data.IpsControl = types.BoolValue(va.Bool())
		}
	}
	data.CautionEnabled = types.BoolNull()

	if t := res.Get(path + "service.cautionEnabled.optionType"); t.Exists() {
		va := res.Get(path + "service.cautionEnabled.value")
		if t.String() == "global" {
			data.CautionEnabled = types.BoolValue(va.Bool())
		}
	}
	data.ZscalerPrimaryDataCenter = types.StringNull()
	data.ZscalerPrimaryDataCenterVariable = types.StringNull()
	if t := res.Get(path + "service.primaryDataCenter.optionType"); t.Exists() {
		va := res.Get(path + "service.primaryDataCenter.value")
		if t.String() == "variable" {
			data.ZscalerPrimaryDataCenterVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ZscalerPrimaryDataCenter = types.StringValue(va.String())
		}
	}
	data.ZscalerSecondaryDataCenter = types.StringNull()
	data.ZscalerSecondaryDataCenterVariable = types.StringNull()
	if t := res.Get(path + "service.secondaryDataCenter.optionType"); t.Exists() {
		va := res.Get(path + "service.secondaryDataCenter.value")
		if t.String() == "variable" {
			data.ZscalerSecondaryDataCenterVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ZscalerSecondaryDataCenter = types.StringValue(va.String())
		}
	}
	data.SurrogateIp = types.BoolNull()

	if t := res.Get(path + "service.ip.optionType"); t.Exists() {
		va := res.Get(path + "service.ip.value")
		if t.String() == "global" {
			data.SurrogateIp = types.BoolValue(va.Bool())
		}
	}
	data.IdleTime = types.Int64Null()

	if t := res.Get(path + "service.idleTime.optionType"); t.Exists() {
		va := res.Get(path + "service.idleTime.value")
		if t.String() == "global" {
			data.IdleTime = types.Int64Value(va.Int())
		}
	}
	data.DisplayTimeUnit = types.StringNull()

	if t := res.Get(path + "service.displayTimeUnit.optionType"); t.Exists() {
		va := res.Get(path + "service.displayTimeUnit.value")
		if t.String() == "global" {
			data.DisplayTimeUnit = types.StringValue(va.String())
		}
	}
	data.IpEnforcedForKnownBrowsers = types.BoolNull()

	if t := res.Get(path + "service.ipEnforcedForKnownBrowsers.optionType"); t.Exists() {
		va := res.Get(path + "service.ipEnforcedForKnownBrowsers.value")
		if t.String() == "global" {
			data.IpEnforcedForKnownBrowsers = types.BoolValue(va.Bool())
		}
	}
	data.RefreshTime = types.Int64Null()

	if t := res.Get(path + "service.refreshTime.optionType"); t.Exists() {
		va := res.Get(path + "service.refreshTime.value")
		if t.String() == "global" {
			data.RefreshTime = types.Int64Value(va.Int())
		}
	}
	data.RefreshTimeUnit = types.StringNull()

	if t := res.Get(path + "service.refreshTimeUnit.optionType"); t.Exists() {
		va := res.Get(path + "service.refreshTimeUnit.value")
		if t.String() == "global" {
			data.RefreshTimeUnit = types.StringValue(va.String())
		}
	}
	data.AupEnabled = types.BoolNull()

	if t := res.Get(path + "service.enabled.optionType"); t.Exists() {
		va := res.Get(path + "service.enabled.value")
		if t.String() == "global" {
			data.AupEnabled = types.BoolValue(va.Bool())
		}
	}
	data.BlockInternetUntilAccepted = types.BoolNull()

	if t := res.Get(path + "service.blockInternetUntilAccepted.optionType"); t.Exists() {
		va := res.Get(path + "service.blockInternetUntilAccepted.value")
		if t.String() == "global" {
			data.BlockInternetUntilAccepted = types.BoolValue(va.Bool())
		}
	}
	data.ForceSslInspection = types.BoolNull()

	if t := res.Get(path + "service.forceSslInspection.optionType"); t.Exists() {
		va := res.Get(path + "service.forceSslInspection.value")
		if t.String() == "global" {
			data.ForceSslInspection = types.BoolValue(va.Bool())
		}
	}
	data.AupTimeout = types.Int64Null()

	if t := res.Get(path + "service.timeout.optionType"); t.Exists() {
		va := res.Get(path + "service.timeout.value")
		if t.String() == "global" {
			data.AupTimeout = types.Int64Value(va.Int())
		}
	}
	tempLocationName := data.LocationName
	data.LocationName = types.StringNull()
	data.LocationNameVariable = types.StringNull()
	if t := res.Get(path + "service.locationName.optionType"); t.Exists() {
		va := res.Get(path + "service.locationName.value")
		if t.String() == "variable" {
			data.LocationNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" || (t.String() == "default" && (!tempLocationName.IsNull() || fullRead)) {
			data.LocationName = types.StringValue(va.String())
		}
	}
	data.UmbrellaPrimaryDataCenter = types.StringNull()
	data.UmbrellaPrimaryDataCenterVariable = types.StringNull()
	if t := res.Get(path + "service.dataCenterPrimary.optionType"); t.Exists() {
		va := res.Get(path + "service.dataCenterPrimary.value")
		if t.String() == "variable" {
			data.UmbrellaPrimaryDataCenterVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.UmbrellaPrimaryDataCenter = types.StringValue(va.String())
		}
	}
	data.UmbrellaSecondaryDataCenter = types.StringNull()
	data.UmbrellaSecondaryDataCenterVariable = types.StringNull()
	if t := res.Get(path + "service.dataCenterSecondary.optionType"); t.Exists() {
		va := res.Get(path + "service.dataCenterSecondary.value")
		if t.String() == "variable" {
			data.UmbrellaSecondaryDataCenterVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.UmbrellaSecondaryDataCenter = types.StringValue(va.String())
		}
	}
	data.TrackerSourceIp = types.StringNull()
	data.TrackerSourceIpVariable = types.StringNull()
	if t := res.Get(path + "trackerSrcIp.optionType"); t.Exists() {
		va := res.Get(path + "trackerSrcIp.value")
		if t.String() == "variable" {
			data.TrackerSourceIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackerSourceIp = types.StringValue(va.String())
		}
	}
	oldTrackers := data.Trackers
	if value := res.Get(path + "tracker"); value.Exists() && len(value.Array()) > 0 {
		data.Trackers = make([]SigTrackers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SigTrackers{}
			item.Name = types.StringNull()

			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			item.EndpointApiUrl = types.StringNull()
			item.EndpointApiUrlVariable = types.StringNull()
			if t := v.Get("endpointApiUrl.optionType"); t.Exists() {
				va := v.Get("endpointApiUrl.value")
				if t.String() == "variable" {
					item.EndpointApiUrlVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.EndpointApiUrl = types.StringValue(va.String())
				}
			}
			item.Threshold = types.Int64Null()
			item.ThresholdVariable = types.StringNull()
			if t := v.Get("threshold.optionType"); t.Exists() {
				va := v.Get("threshold.value")
				if t.String() == "variable" {
					item.ThresholdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Threshold = types.Int64Value(va.Int())
				}
			}
			item.Interval = types.Int64Null()
			item.IntervalVariable = types.StringNull()
			if t := v.Get("interval.optionType"); t.Exists() {
				va := v.Get("interval.value")
				if t.String() == "variable" {
					item.IntervalVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Interval = types.Int64Value(va.Int())
				}
			}
			item.Multiplier = types.Int64Null()
			item.MultiplierVariable = types.StringNull()
			if t := v.Get("multiplier.optionType"); t.Exists() {
				va := v.Get("multiplier.value")
				if t.String() == "variable" {
					item.MultiplierVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Multiplier = types.Int64Value(va.Int())
				}
			}
			data.Trackers = append(data.Trackers, item)
			return true
		})
	} else {
		data.Trackers = nil
	}
	if !fullRead && data.Trackers != nil {
		resultTrackers := make([]SigTrackers, 0, len(data.Trackers))
		matchedTrackers := make([]bool, len(data.Trackers))
		for _, oldItem := range oldTrackers {
			for ni := range data.Trackers {
				if matchedTrackers[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.Name.ValueString() != data.Trackers[ni].Name.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedTrackers[ni] = true
					resultTrackers = append(resultTrackers, data.Trackers[ni])
					break
				}
			}
		}
		for ni := range data.Trackers {
			if !matchedTrackers[ni] {
				resultTrackers = append(resultTrackers, data.Trackers[ni])
			}
		}
		data.Trackers = resultTrackers
	}
}

// End of section. //template:end fromBody
