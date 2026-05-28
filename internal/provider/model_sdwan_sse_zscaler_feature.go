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
type SSEZscaler struct {
	Id                                 types.String               `tfsdk:"id"`
	Version                            types.Int64                `tfsdk:"version"`
	Name                               types.String               `tfsdk:"name"`
	Description                        types.String               `tfsdk:"description"`
	FeatureProfileId                   types.String               `tfsdk:"feature_profile_id"`
	SrcVpn                             types.Bool                 `tfsdk:"src_vpn"`
	Interfaces                         []SSEZscalerInterfaces     `tfsdk:"interfaces"`
	InterfacePairs                     []SSEZscalerInterfacePairs `tfsdk:"interface_pairs"`
	AuthRequired                       types.Bool                 `tfsdk:"auth_required"`
	AuthRequiredVariable               types.String               `tfsdk:"auth_required_variable"`
	XffForwardEnabled                  types.Bool                 `tfsdk:"xff_forward_enabled"`
	XffForwardEnabledVariable          types.String               `tfsdk:"xff_forward_enabled_variable"`
	OfwEnabled                         types.Bool                 `tfsdk:"ofw_enabled"`
	OfwEnabledVariable                 types.String               `tfsdk:"ofw_enabled_variable"`
	IpsControl                         types.Bool                 `tfsdk:"ips_control"`
	IpsControlVariable                 types.String               `tfsdk:"ips_control_variable"`
	CautionEnabled                     types.Bool                 `tfsdk:"caution_enabled"`
	CautionEnabledVariable             types.String               `tfsdk:"caution_enabled_variable"`
	PrimaryDataCenter                  types.String               `tfsdk:"primary_data_center"`
	PrimaryDataCenterVariable          types.String               `tfsdk:"primary_data_center_variable"`
	SecondaryDataCenter                types.String               `tfsdk:"secondary_data_center"`
	SecondaryDataCenterVariable        types.String               `tfsdk:"secondary_data_center_variable"`
	SurrogateIp                        types.Bool                 `tfsdk:"surrogate_ip"`
	SurrogateIpVariable                types.String               `tfsdk:"surrogate_ip_variable"`
	IdleTime                           types.Int64                `tfsdk:"idle_time"`
	IdleTimeVariable                   types.String               `tfsdk:"idle_time_variable"`
	DisplayTimeUnit                    types.String               `tfsdk:"display_time_unit"`
	DisplayTimeUnitVariable            types.String               `tfsdk:"display_time_unit_variable"`
	IpEnforcedForKnownBrowsers         types.Bool                 `tfsdk:"ip_enforced_for_known_browsers"`
	IpEnforcedForKnownBrowsersVariable types.String               `tfsdk:"ip_enforced_for_known_browsers_variable"`
	RefreshTime                        types.Int64                `tfsdk:"refresh_time"`
	RefreshTimeVariable                types.String               `tfsdk:"refresh_time_variable"`
	RefreshTimeUnit                    types.String               `tfsdk:"refresh_time_unit"`
	RefreshTimeUnitVariable            types.String               `tfsdk:"refresh_time_unit_variable"`
	AupEnabled                         types.Bool                 `tfsdk:"aup_enabled"`
	AupEnabledVariable                 types.String               `tfsdk:"aup_enabled_variable"`
	BlockInternetUntilAccepted         types.Bool                 `tfsdk:"block_internet_until_accepted"`
	BlockInternetUntilAcceptedVariable types.String               `tfsdk:"block_internet_until_accepted_variable"`
	ForceSslInspection                 types.Bool                 `tfsdk:"force_ssl_inspection"`
	ForceSslInspectionVariable         types.String               `tfsdk:"force_ssl_inspection_variable"`
	AupTimeout                         types.Int64                `tfsdk:"aup_timeout"`
	AupTimeoutVariable                 types.String               `tfsdk:"aup_timeout_variable"`
	LocationName                       types.String               `tfsdk:"location_name"`
	LocationNameVariable               types.String               `tfsdk:"location_name_variable"`
	Country                            types.Bool                 `tfsdk:"country"`
	CountryVariable                    types.String               `tfsdk:"country_variable"`
	EnforceBandwidthControl            types.Bool                 `tfsdk:"enforce_bandwidth_control"`
	EnforceBandwidthControlVariable    types.String               `tfsdk:"enforce_bandwidth_control_variable"`
	DnBandwidth                        types.Float64              `tfsdk:"dn_bandwidth"`
	DnBandwidthVariable                types.String               `tfsdk:"dn_bandwidth_variable"`
	UpBandwidth                        types.Float64              `tfsdk:"up_bandwidth"`
	UpBandwidthVariable                types.String               `tfsdk:"up_bandwidth_variable"`
	SubLocations                       []SSEZscalerSubLocations   `tfsdk:"sub_locations"`
	TrackerSourceIp                    types.String               `tfsdk:"tracker_source_ip"`
	TrackerSourceIpVariable            types.String               `tfsdk:"tracker_source_ip_variable"`
	Trackers                           []SSEZscalerTrackers       `tfsdk:"trackers"`
}

type SSEZscalerInterfaces struct {
	InterfaceName                 types.String `tfsdk:"interface_name"`
	Auto                          types.Bool   `tfsdk:"auto"`
	Shutdown                      types.Bool   `tfsdk:"shutdown"`
	InterfaceDescription          types.String `tfsdk:"interface_description"`
	InterfaceDescriptionVariable  types.String `tfsdk:"interface_description_variable"`
	Unnumbered                    types.Bool   `tfsdk:"unnumbered"`
	Ipv4Address                   types.String `tfsdk:"ipv4_address"`
	Ipv4AddressVariable           types.String `tfsdk:"ipv4_address_variable"`
	TunnelSource                  types.String `tfsdk:"tunnel_source"`
	TunnelSourceVariable          types.String `tfsdk:"tunnel_source_variable"`
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
	IkeVersion                    types.Int64  `tfsdk:"ike_version"`
	IkeVersionVariable            types.String `tfsdk:"ike_version_variable"`
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

type SSEZscalerInterfacePairs struct {
	ActiveInterface       types.String `tfsdk:"active_interface"`
	ActiveInterfaceWeight types.Int64  `tfsdk:"active_interface_weight"`
	BackupInterface       types.String `tfsdk:"backup_interface"`
	BackupInterfaceWeight types.Int64  `tfsdk:"backup_interface_weight"`
}

type SSEZscalerSubLocations struct {
	Name                               types.String                       `tfsdk:"name"`
	NameVariable                       types.String                       `tfsdk:"name_variable"`
	AuthRequired                       types.Bool                         `tfsdk:"auth_required"`
	AuthRequiredVariable               types.String                       `tfsdk:"auth_required_variable"`
	SurrogateIp                        types.Bool                         `tfsdk:"surrogate_ip"`
	SurrogateIpVariable                types.String                       `tfsdk:"surrogate_ip_variable"`
	IdleTime                           types.Int64                        `tfsdk:"idle_time"`
	IdleTimeVariable                   types.String                       `tfsdk:"idle_time_variable"`
	DisplayTimeUnit                    types.String                       `tfsdk:"display_time_unit"`
	DisplayTimeUnitVariable            types.String                       `tfsdk:"display_time_unit_variable"`
	IpEnforcedForKnownBrowsers         types.Bool                         `tfsdk:"ip_enforced_for_known_browsers"`
	IpEnforcedForKnownBrowsersVariable types.String                       `tfsdk:"ip_enforced_for_known_browsers_variable"`
	RefreshTime                        types.Int64                        `tfsdk:"refresh_time"`
	RefreshTimeVariable                types.String                       `tfsdk:"refresh_time_variable"`
	RefreshTimeUnit                    types.String                       `tfsdk:"refresh_time_unit"`
	RefreshTimeUnitVariable            types.String                       `tfsdk:"refresh_time_unit_variable"`
	OfwEnabled                         types.Bool                         `tfsdk:"ofw_enabled"`
	OfwEnabledVariable                 types.String                       `tfsdk:"ofw_enabled_variable"`
	CautionEnabled                     types.Bool                         `tfsdk:"caution_enabled"`
	CautionEnabledVariable             types.String                       `tfsdk:"caution_enabled_variable"`
	ServiceVpn                         types.Set                          `tfsdk:"service_vpn"`
	ServiceVpnVariable                 types.String                       `tfsdk:"service_vpn_variable"`
	InternalIp                         []SSEZscalerSubLocationsInternalIp `tfsdk:"internal_ip"`
	AupEnabled                         types.Bool                         `tfsdk:"aup_enabled"`
	AupEnabledVariable                 types.String                       `tfsdk:"aup_enabled_variable"`
	BlockInternetUntilAccepted         types.Bool                         `tfsdk:"block_internet_until_accepted"`
	BlockInternetUntilAcceptedVariable types.String                       `tfsdk:"block_internet_until_accepted_variable"`
	ForceSslInspection                 types.Bool                         `tfsdk:"force_ssl_inspection"`
	ForceSslInspectionVariable         types.String                       `tfsdk:"force_ssl_inspection_variable"`
	AupTimeout                         types.Int64                        `tfsdk:"aup_timeout"`
	AupTimeoutVariable                 types.String                       `tfsdk:"aup_timeout_variable"`
	EnforceBandwidthControl            types.String                       `tfsdk:"enforce_bandwidth_control"`
	EnforceBandwidthControlVariable    types.String                       `tfsdk:"enforce_bandwidth_control_variable"`
	UpBandwidth                        types.Float64                      `tfsdk:"up_bandwidth"`
	UpBandwidthVariable                types.String                       `tfsdk:"up_bandwidth_variable"`
	DnBandwidth                        types.Float64                      `tfsdk:"dn_bandwidth"`
	DnBandwidthVariable                types.String                       `tfsdk:"dn_bandwidth_variable"`
}

type SSEZscalerTrackers struct {
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

type SSEZscalerSubLocationsInternalIp struct {
	InternalIpValue         types.String `tfsdk:"internal_ip_value"`
	InternalIpValueVariable types.String `tfsdk:"internal_ip_value_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SSEZscaler) getModel() string {
	return "sse_zscaler"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SSEZscaler) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/sse/%v/zscaler", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SSEZscaler) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {
		body, _ = sjson.Set(body, path+"sse_instance.optionType", "default")
		body, _ = sjson.Set(body, path+"sse_instance.value", "zScaler")
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
			if !item.Auto.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "auto.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "auto.value", item.Auto.ValueBool())
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
			if !item.Unnumbered.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "unnumbered.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "unnumbered.value", item.Unnumbered.ValueBool())
				}
			}

			if !item.Ipv4AddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.Ipv4AddressVariable.ValueString())
				}
			} else if item.Ipv4Address.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.Ipv4Address.ValueString())
				}
			}

			if !item.TunnelSourceVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tunnelSource.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "tunnelSource.value", item.TunnelSourceVariable.ValueString())
				}
			} else if !item.TunnelSource.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tunnelSource.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tunnelSource.value", item.TunnelSource.ValueString())
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
				if true {
					itemBody, _ = sjson.Set(itemBody, "tunnelRouteVia.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "tunnelRouteVia.value", item.TunnelRouteViaVariable.ValueString())
				}
			} else if !item.TunnelRouteVia.IsNull() {
				if true {
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
				itemBody, _ = sjson.Set(itemBody, "application.value", "sse")
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

			if !item.IkeVersionVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeVersion.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ikeVersion.value", item.IkeVersionVariable.ValueString())
				}
			} else if item.IkeVersion.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeVersion.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "ikeVersion.value", 2)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeVersion.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ikeVersion.value", item.IkeVersion.ValueInt64())
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
					itemBody, _ = sjson.Set(itemBody, "tracker.value", "DefaultTracker")
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

	if !data.AuthRequiredVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.authRequired.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.authRequired.value", data.AuthRequiredVariable.ValueString())
		}
	} else if data.AuthRequired.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.authRequired.optionType", "default")
			body, _ = sjson.Set(body, path+"service.authRequired.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"service.authRequired.optionType", "global")
			body, _ = sjson.Set(body, path+"service.authRequired.value", data.AuthRequired.ValueBool())
		}
	}

	if !data.XffForwardEnabledVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.xffForwardEnabled.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.xffForwardEnabled.value", data.XffForwardEnabledVariable.ValueString())
		}
	} else if data.XffForwardEnabled.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.xffForwardEnabled.optionType", "default")
			body, _ = sjson.Set(body, path+"service.xffForwardEnabled.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"service.xffForwardEnabled.optionType", "global")
			body, _ = sjson.Set(body, path+"service.xffForwardEnabled.value", data.XffForwardEnabled.ValueBool())
		}
	}

	if !data.OfwEnabledVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.ofwEnabled.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.ofwEnabled.value", data.OfwEnabledVariable.ValueString())
		}
	} else if data.OfwEnabled.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.ofwEnabled.optionType", "default")
			body, _ = sjson.Set(body, path+"service.ofwEnabled.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"service.ofwEnabled.optionType", "global")
			body, _ = sjson.Set(body, path+"service.ofwEnabled.value", data.OfwEnabled.ValueBool())
		}
	}

	if !data.IpsControlVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.ipsControl.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.ipsControl.value", data.IpsControlVariable.ValueString())
		}
	} else if data.IpsControl.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.ipsControl.optionType", "default")
			body, _ = sjson.Set(body, path+"service.ipsControl.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"service.ipsControl.optionType", "global")
			body, _ = sjson.Set(body, path+"service.ipsControl.value", data.IpsControl.ValueBool())
		}
	}

	if !data.CautionEnabledVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.cautionEnabled.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.cautionEnabled.value", data.CautionEnabledVariable.ValueString())
		}
	} else if data.CautionEnabled.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.cautionEnabled.optionType", "default")
			body, _ = sjson.Set(body, path+"service.cautionEnabled.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"service.cautionEnabled.optionType", "global")
			body, _ = sjson.Set(body, path+"service.cautionEnabled.value", data.CautionEnabled.ValueBool())
		}
	}

	if !data.PrimaryDataCenterVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.primaryDataCenter.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.primaryDataCenter.value", data.PrimaryDataCenterVariable.ValueString())
		}
	} else if data.PrimaryDataCenter.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.primaryDataCenter.optionType", "default")
			body, _ = sjson.Set(body, path+"service.primaryDataCenter.value", "Auto")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"service.primaryDataCenter.optionType", "global")
			body, _ = sjson.Set(body, path+"service.primaryDataCenter.value", data.PrimaryDataCenter.ValueString())
		}
	}

	if !data.SecondaryDataCenterVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.secondaryDataCenter.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.secondaryDataCenter.value", data.SecondaryDataCenterVariable.ValueString())
		}
	} else if data.SecondaryDataCenter.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.secondaryDataCenter.optionType", "default")
			body, _ = sjson.Set(body, path+"service.secondaryDataCenter.value", "Auto")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"service.secondaryDataCenter.optionType", "global")
			body, _ = sjson.Set(body, path+"service.secondaryDataCenter.value", data.SecondaryDataCenter.ValueString())
		}
	}

	if !data.SurrogateIpVariable.IsNull() {
		if true && data.AuthRequired.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.ip.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.ip.value", data.SurrogateIpVariable.ValueString())
		}
	} else if data.SurrogateIp.IsNull() {
		if true && data.AuthRequired.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.ip.optionType", "default")
			body, _ = sjson.Set(body, path+"service.ip.value", false)
		}
	} else {
		if true && data.AuthRequired.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.ip.optionType", "global")
			body, _ = sjson.Set(body, path+"service.ip.value", data.SurrogateIp.ValueBool())
		}
	}

	if !data.IdleTimeVariable.IsNull() {
		if true && data.SurrogateIp.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.idleTime.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.idleTime.value", data.IdleTimeVariable.ValueString())
		}
	} else if data.IdleTime.IsNull() {
		if true && data.SurrogateIp.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.idleTime.optionType", "default")
			body, _ = sjson.Set(body, path+"service.idleTime.value", 1)
		}
	} else {
		if true && data.SurrogateIp.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.idleTime.optionType", "global")
			body, _ = sjson.Set(body, path+"service.idleTime.value", data.IdleTime.ValueInt64())
		}
	}

	if !data.DisplayTimeUnitVariable.IsNull() {
		if true && data.SurrogateIp.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.displayTimeUnit.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.displayTimeUnit.value", data.DisplayTimeUnitVariable.ValueString())
		}
	} else if data.DisplayTimeUnit.IsNull() {
		if true && data.SurrogateIp.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.displayTimeUnit.optionType", "default")
			body, _ = sjson.Set(body, path+"service.displayTimeUnit.value", "MINUTE")
		}
	} else {
		if true && data.SurrogateIp.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.displayTimeUnit.optionType", "global")
			body, _ = sjson.Set(body, path+"service.displayTimeUnit.value", data.DisplayTimeUnit.ValueString())
		}
	}

	if !data.IpEnforcedForKnownBrowsersVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.ipEnforcedForKnownBrowsers.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.ipEnforcedForKnownBrowsers.value", data.IpEnforcedForKnownBrowsersVariable.ValueString())
		}
	} else if data.IpEnforcedForKnownBrowsers.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.ipEnforcedForKnownBrowsers.optionType", "default")
			body, _ = sjson.Set(body, path+"service.ipEnforcedForKnownBrowsers.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"service.ipEnforcedForKnownBrowsers.optionType", "global")
			body, _ = sjson.Set(body, path+"service.ipEnforcedForKnownBrowsers.value", data.IpEnforcedForKnownBrowsers.ValueBool())
		}
	}

	if !data.RefreshTimeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.refreshTime.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.refreshTime.value", data.RefreshTimeVariable.ValueString())
		}
	} else if data.RefreshTime.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.refreshTime.optionType", "default")
			body, _ = sjson.Set(body, path+"service.refreshTime.value", 1)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"service.refreshTime.optionType", "global")
			body, _ = sjson.Set(body, path+"service.refreshTime.value", data.RefreshTime.ValueInt64())
		}
	}

	if !data.RefreshTimeUnitVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.refreshTimeUnit.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.refreshTimeUnit.value", data.RefreshTimeUnitVariable.ValueString())
		}
	} else if data.RefreshTimeUnit.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.refreshTimeUnit.optionType", "default")
			body, _ = sjson.Set(body, path+"service.refreshTimeUnit.value", "MINUTE")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"service.refreshTimeUnit.optionType", "global")
			body, _ = sjson.Set(body, path+"service.refreshTimeUnit.value", data.RefreshTimeUnit.ValueString())
		}
	}

	if !data.AupEnabledVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.enabled.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.enabled.value", data.AupEnabledVariable.ValueString())
		}
	} else if data.AupEnabled.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.enabled.optionType", "default")
			body, _ = sjson.Set(body, path+"service.enabled.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"service.enabled.optionType", "global")
			body, _ = sjson.Set(body, path+"service.enabled.value", data.AupEnabled.ValueBool())
		}
	}

	if !data.BlockInternetUntilAcceptedVariable.IsNull() {
		if true && data.AupEnabled.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.blockInternetUntilAccepted.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.blockInternetUntilAccepted.value", data.BlockInternetUntilAcceptedVariable.ValueString())
		}
	} else if data.BlockInternetUntilAccepted.IsNull() {
		if true && data.AupEnabled.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.blockInternetUntilAccepted.optionType", "default")
			body, _ = sjson.Set(body, path+"service.blockInternetUntilAccepted.value", false)
		}
	} else {
		if true && data.AupEnabled.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.blockInternetUntilAccepted.optionType", "global")
			body, _ = sjson.Set(body, path+"service.blockInternetUntilAccepted.value", data.BlockInternetUntilAccepted.ValueBool())
		}
	}

	if !data.ForceSslInspectionVariable.IsNull() {
		if true && data.AupEnabled.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.forceSslInspection.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.forceSslInspection.value", data.ForceSslInspectionVariable.ValueString())
		}
	} else if data.ForceSslInspection.IsNull() {
		if true && data.AupEnabled.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.forceSslInspection.optionType", "default")
			body, _ = sjson.Set(body, path+"service.forceSslInspection.value", false)
		}
	} else {
		if true && data.AupEnabled.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.forceSslInspection.optionType", "global")
			body, _ = sjson.Set(body, path+"service.forceSslInspection.value", data.ForceSslInspection.ValueBool())
		}
	}

	if !data.AupTimeoutVariable.IsNull() {
		if true && data.AupEnabled.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.timeout.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.timeout.value", data.AupTimeoutVariable.ValueString())
		}
	} else if data.AupTimeout.IsNull() {
		if true && data.AupEnabled.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.timeout.optionType", "default")
			body, _ = sjson.Set(body, path+"service.timeout.value", 1)
		}
	} else {
		if true && data.AupEnabled.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.timeout.optionType", "global")
			body, _ = sjson.Set(body, path+"service.timeout.value", data.AupTimeout.ValueInt64())
		}
	}

	if !data.LocationNameVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.locationName.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.locationName.value", data.LocationNameVariable.ValueString())
		}
	} else if data.LocationName.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.locationName.optionType", "default")
			body, _ = sjson.Set(body, path+"service.locationName.value", "Auto")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"service.locationName.optionType", "default")
			body, _ = sjson.Set(body, path+"service.locationName.value", data.LocationName.ValueString())
		}
	}

	if !data.CountryVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.country.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.country.value", data.CountryVariable.ValueString())
		}
	} else if data.Country.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.country.optionType", "default")
			body, _ = sjson.Set(body, path+"service.country.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"service.country.optionType", "global")
			body, _ = sjson.Set(body, path+"service.country.value", data.Country.ValueBool())
		}
	}

	if !data.EnforceBandwidthControlVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.enforceBandwidthControl.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.enforceBandwidthControl.value", data.EnforceBandwidthControlVariable.ValueString())
		}
	} else if data.EnforceBandwidthControl.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"service.enforceBandwidthControl.optionType", "default")
			body, _ = sjson.Set(body, path+"service.enforceBandwidthControl.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"service.enforceBandwidthControl.optionType", "global")
			body, _ = sjson.Set(body, path+"service.enforceBandwidthControl.value", data.EnforceBandwidthControl.ValueBool())
		}
	}

	if !data.DnBandwidthVariable.IsNull() {
		if true && data.EnforceBandwidthControl.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.dnBandwidth.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.dnBandwidth.value", data.DnBandwidthVariable.ValueString())
		}
	} else if !data.DnBandwidth.IsNull() {
		if true && data.EnforceBandwidthControl.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.dnBandwidth.optionType", "global")
			body, _ = sjson.Set(body, path+"service.dnBandwidth.value", data.DnBandwidth.ValueFloat64())
		}
	}

	if !data.UpBandwidthVariable.IsNull() {
		if true && data.EnforceBandwidthControl.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.upBandwidth.optionType", "variable")
			body, _ = sjson.Set(body, path+"service.upBandwidth.value", data.UpBandwidthVariable.ValueString())
		}
	} else if !data.UpBandwidth.IsNull() {
		if true && data.EnforceBandwidthControl.ValueBool() == true {
			body, _ = sjson.Set(body, path+"service.upBandwidth.optionType", "global")
			body, _ = sjson.Set(body, path+"service.upBandwidth.value", data.UpBandwidth.ValueFloat64())
		}
	}
	if true {

		for _, item := range data.SubLocations {
			itemBody := ""

			if !item.NameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.NameVariable.ValueString())
				}
			} else if !item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
				}
			}

			if !item.AuthRequiredVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "authRequired.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "authRequired.value", item.AuthRequiredVariable.ValueString())
				}
			} else if item.AuthRequired.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "authRequired.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "authRequired.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "authRequired.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "authRequired.value", item.AuthRequired.ValueBool())
				}
			}

			if !item.SurrogateIpVariable.IsNull() {
				if true && item.AuthRequired.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "ip.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ip.value", item.SurrogateIpVariable.ValueString())
				}
			} else if item.SurrogateIp.IsNull() {
				if true && item.AuthRequired.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "ip.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "ip.value", false)
				}
			} else {
				if true && item.AuthRequired.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "ip.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ip.value", item.SurrogateIp.ValueBool())
				}
			}

			if !item.IdleTimeVariable.IsNull() {
				if true && item.SurrogateIp.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "idleTime.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "idleTime.value", item.IdleTimeVariable.ValueString())
				}
			} else if item.IdleTime.IsNull() {
				if true && item.SurrogateIp.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "idleTime.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "idleTime.value", 1)
				}
			} else {
				if true && item.SurrogateIp.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "idleTime.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "idleTime.value", item.IdleTime.ValueInt64())
				}
			}

			if !item.DisplayTimeUnitVariable.IsNull() {
				if true && item.SurrogateIp.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "displayTimeUnit.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "displayTimeUnit.value", item.DisplayTimeUnitVariable.ValueString())
				}
			} else if item.DisplayTimeUnit.IsNull() {
				if true && item.SurrogateIp.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "displayTimeUnit.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "displayTimeUnit.value", "MINUTE")
				}
			} else {
				if true && item.SurrogateIp.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "displayTimeUnit.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "displayTimeUnit.value", item.DisplayTimeUnit.ValueString())
				}
			}

			if !item.IpEnforcedForKnownBrowsersVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipEnforcedForKnownBrowsers.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ipEnforcedForKnownBrowsers.value", item.IpEnforcedForKnownBrowsersVariable.ValueString())
				}
			} else if item.IpEnforcedForKnownBrowsers.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipEnforcedForKnownBrowsers.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "ipEnforcedForKnownBrowsers.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipEnforcedForKnownBrowsers.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ipEnforcedForKnownBrowsers.value", item.IpEnforcedForKnownBrowsers.ValueBool())
				}
			}

			if !item.RefreshTimeVariable.IsNull() {
				if true && item.IpEnforcedForKnownBrowsers.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "refreshTime.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "refreshTime.value", item.RefreshTimeVariable.ValueString())
				}
			} else if item.RefreshTime.IsNull() {
				if true && item.IpEnforcedForKnownBrowsers.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "refreshTime.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "refreshTime.value", 1)
				}
			} else {
				if true && item.IpEnforcedForKnownBrowsers.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "refreshTime.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "refreshTime.value", item.RefreshTime.ValueInt64())
				}
			}

			if !item.RefreshTimeUnitVariable.IsNull() {
				if true && item.IpEnforcedForKnownBrowsers.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "refreshTimeUnit.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "refreshTimeUnit.value", item.RefreshTimeUnitVariable.ValueString())
				}
			} else if item.RefreshTimeUnit.IsNull() {
				if true && item.IpEnforcedForKnownBrowsers.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "refreshTimeUnit.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "refreshTimeUnit.value", "MINUTE")
				}
			} else {
				if true && item.IpEnforcedForKnownBrowsers.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "refreshTimeUnit.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "refreshTimeUnit.value", item.RefreshTimeUnit.ValueString())
				}
			}

			if !item.OfwEnabledVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ofwEnabled.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ofwEnabled.value", item.OfwEnabledVariable.ValueString())
				}
			} else if item.OfwEnabled.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ofwEnabled.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "ofwEnabled.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ofwEnabled.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ofwEnabled.value", item.OfwEnabled.ValueBool())
				}
			}

			if !item.CautionEnabledVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "cautionEnabled.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "cautionEnabled.value", item.CautionEnabledVariable.ValueString())
				}
			} else if item.CautionEnabled.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "cautionEnabled.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "cautionEnabled.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "cautionEnabled.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "cautionEnabled.value", item.CautionEnabled.ValueBool())
				}
			}

			if !item.ServiceVpnVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "serviceVPN.serviceVpnValue.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "serviceVPN.serviceVpnValue.value", item.ServiceVpnVariable.ValueString())
				}
			} else if !item.ServiceVpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "serviceVPN.serviceVpnValue.optionType", "global")
					var values []string
					item.ServiceVpn.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "serviceVPN.serviceVpnValue.value", values)
				}
			}
			if true {

				for _, childItem := range item.InternalIp {
					itemChildBody := ""

					if !childItem.InternalIpValueVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "internalIpValue.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "internalIpValue.value", childItem.InternalIpValueVariable.ValueString())
						}
					} else if !childItem.InternalIpValue.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "internalIpValue.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "internalIpValue.value", childItem.InternalIpValue.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "internalIp.-1", itemChildBody)
				}
			}

			if !item.AupEnabledVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "aupEnabled.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "aupEnabled.value", item.AupEnabledVariable.ValueString())
				}
			} else if item.AupEnabled.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "aupEnabled.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "aupEnabled.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "aupEnabled.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "aupEnabled.value", item.AupEnabled.ValueBool())
				}
			}

			if !item.BlockInternetUntilAcceptedVariable.IsNull() {
				if true && item.AupEnabled.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "blockInternetUntilAccepted.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "blockInternetUntilAccepted.value", item.BlockInternetUntilAcceptedVariable.ValueString())
				}
			} else if !item.BlockInternetUntilAccepted.IsNull() {
				if true && item.AupEnabled.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "blockInternetUntilAccepted.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "blockInternetUntilAccepted.value", item.BlockInternetUntilAccepted.ValueBool())
				}
			}

			if !item.ForceSslInspectionVariable.IsNull() {
				if true && item.AupEnabled.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "forceSslInspection.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "forceSslInspection.value", item.ForceSslInspectionVariable.ValueString())
				}
			} else if !item.ForceSslInspection.IsNull() {
				if true && item.AupEnabled.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "forceSslInspection.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "forceSslInspection.value", item.ForceSslInspection.ValueBool())
				}
			}

			if !item.AupTimeoutVariable.IsNull() {
				if true && item.AupEnabled.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "timeout.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "timeout.value", item.AupTimeoutVariable.ValueString())
				}
			} else if item.AupTimeout.IsNull() {
				if true && item.AupEnabled.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "timeout.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "timeout.value", 1)
				}
			} else {
				if true && item.AupEnabled.ValueBool() == true {
					itemBody, _ = sjson.Set(itemBody, "timeout.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "timeout.value", item.AupTimeout.ValueInt64())
				}
			}

			if !item.EnforceBandwidthControlVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "enforceBandwidthControl.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "enforceBandwidthControl.value", item.EnforceBandwidthControlVariable.ValueString())
				}
			} else if item.EnforceBandwidthControl.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "enforceBandwidthControl.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "enforceBandwidthControl.value", "location-bandwidth")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "enforceBandwidthControl.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "enforceBandwidthControl.value", item.EnforceBandwidthControl.ValueString())
				}
			}

			if !item.UpBandwidthVariable.IsNull() {
				if true && item.EnforceBandwidthControl.ValueString() == "override" {
					itemBody, _ = sjson.Set(itemBody, "upBandwidth.bandwidthValue.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "upBandwidth.bandwidthValue.value", item.UpBandwidthVariable.ValueString())
				}
			} else if !item.UpBandwidth.IsNull() {
				if true && item.EnforceBandwidthControl.ValueString() == "override" {
					itemBody, _ = sjson.Set(itemBody, "upBandwidth.bandwidthValue.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "upBandwidth.bandwidthValue.value", item.UpBandwidth.ValueFloat64())
				}
			}

			if !item.DnBandwidthVariable.IsNull() {
				if true && item.EnforceBandwidthControl.ValueString() == "override" {
					itemBody, _ = sjson.Set(itemBody, "dnBandwidth.bandwidthValue.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "dnBandwidth.bandwidthValue.value", item.DnBandwidthVariable.ValueString())
				}
			} else if !item.DnBandwidth.IsNull() {
				if true && item.EnforceBandwidthControl.ValueString() == "override" {
					itemBody, _ = sjson.Set(itemBody, "dnBandwidth.bandwidthValue.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "dnBandwidth.bandwidthValue.value", item.DnBandwidth.ValueFloat64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"service.subLocations.-1", itemBody)
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
					itemBody, _ = sjson.Set(itemBody, "threshold.value", 1000)
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
					itemBody, _ = sjson.Set(itemBody, "interval.value", 30)
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
					itemBody, _ = sjson.Set(itemBody, "multiplier.value", 2)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "multiplier.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "multiplier.value", item.Multiplier.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"tracker.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SSEZscaler) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.SrcVpn = types.BoolNull()

	if t := res.Get(path + "interfaceMetadataSharing.srcVpn.optionType"); t.Exists() {
		va := res.Get(path + "interfaceMetadataSharing.srcVpn.value")
		if t.String() == "global" {
			data.SrcVpn = types.BoolValue(va.Bool())
		}
	}
	if value := res.Get(path + "interface"); value.Exists() && len(value.Array()) > 0 {
		data.Interfaces = make([]SSEZscalerInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SSEZscalerInterfaces{}
			item.InterfaceName = types.StringNull()

			if t := v.Get("ifName.optionType"); t.Exists() {
				va := v.Get("ifName.value")
				if t.String() == "global" {
					item.InterfaceName = types.StringValue(va.String())
				}
			}
			item.Auto = types.BoolNull()

			if t := v.Get("auto.optionType"); t.Exists() {
				va := v.Get("auto.value")
				if t.String() == "global" {
					item.Auto = types.BoolValue(va.Bool())
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
			item.Unnumbered = types.BoolNull()

			if t := v.Get("unnumbered.optionType"); t.Exists() {
				va := v.Get("unnumbered.value")
				if t.String() == "global" {
					item.Unnumbered = types.BoolValue(va.Bool())
				}
			}
			item.Ipv4Address = types.StringNull()
			item.Ipv4AddressVariable = types.StringNull()
			if t := v.Get("address.optionType"); t.Exists() {
				va := v.Get("address.value")
				if t.String() == "variable" {
					item.Ipv4AddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Ipv4Address = types.StringValue(va.String())
				}
			}
			item.TunnelSource = types.StringNull()
			item.TunnelSourceVariable = types.StringNull()
			if t := v.Get("tunnelSource.optionType"); t.Exists() {
				va := v.Get("tunnelSource.value")
				if t.String() == "variable" {
					item.TunnelSourceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TunnelSource = types.StringValue(va.String())
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
			item.IkeVersion = types.Int64Null()
			item.IkeVersionVariable = types.StringNull()
			if t := v.Get("ikeVersion.optionType"); t.Exists() {
				va := v.Get("ikeVersion.value")
				if t.String() == "variable" {
					item.IkeVersionVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IkeVersion = types.Int64Value(va.Int())
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
	}
	if value := res.Get(path + "service.interfacePair"); value.Exists() && len(value.Array()) > 0 {
		data.InterfacePairs = make([]SSEZscalerInterfacePairs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SSEZscalerInterfacePairs{}
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
	}
	data.AuthRequired = types.BoolNull()
	data.AuthRequiredVariable = types.StringNull()
	if t := res.Get(path + "service.authRequired.optionType"); t.Exists() {
		va := res.Get(path + "service.authRequired.value")
		if t.String() == "variable" {
			data.AuthRequiredVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AuthRequired = types.BoolValue(va.Bool())
		}
	}
	data.XffForwardEnabled = types.BoolNull()
	data.XffForwardEnabledVariable = types.StringNull()
	if t := res.Get(path + "service.xffForwardEnabled.optionType"); t.Exists() {
		va := res.Get(path + "service.xffForwardEnabled.value")
		if t.String() == "variable" {
			data.XffForwardEnabledVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.XffForwardEnabled = types.BoolValue(va.Bool())
		}
	}
	data.OfwEnabled = types.BoolNull()
	data.OfwEnabledVariable = types.StringNull()
	if t := res.Get(path + "service.ofwEnabled.optionType"); t.Exists() {
		va := res.Get(path + "service.ofwEnabled.value")
		if t.String() == "variable" {
			data.OfwEnabledVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.OfwEnabled = types.BoolValue(va.Bool())
		}
	}
	data.IpsControl = types.BoolNull()
	data.IpsControlVariable = types.StringNull()
	if t := res.Get(path + "service.ipsControl.optionType"); t.Exists() {
		va := res.Get(path + "service.ipsControl.value")
		if t.String() == "variable" {
			data.IpsControlVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsControl = types.BoolValue(va.Bool())
		}
	}
	data.CautionEnabled = types.BoolNull()
	data.CautionEnabledVariable = types.StringNull()
	if t := res.Get(path + "service.cautionEnabled.optionType"); t.Exists() {
		va := res.Get(path + "service.cautionEnabled.value")
		if t.String() == "variable" {
			data.CautionEnabledVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.CautionEnabled = types.BoolValue(va.Bool())
		}
	}
	data.PrimaryDataCenter = types.StringNull()
	data.PrimaryDataCenterVariable = types.StringNull()
	if t := res.Get(path + "service.primaryDataCenter.optionType"); t.Exists() {
		va := res.Get(path + "service.primaryDataCenter.value")
		if t.String() == "variable" {
			data.PrimaryDataCenterVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PrimaryDataCenter = types.StringValue(va.String())
		}
	}
	data.SecondaryDataCenter = types.StringNull()
	data.SecondaryDataCenterVariable = types.StringNull()
	if t := res.Get(path + "service.secondaryDataCenter.optionType"); t.Exists() {
		va := res.Get(path + "service.secondaryDataCenter.value")
		if t.String() == "variable" {
			data.SecondaryDataCenterVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SecondaryDataCenter = types.StringValue(va.String())
		}
	}
	data.SurrogateIp = types.BoolNull()
	data.SurrogateIpVariable = types.StringNull()
	if t := res.Get(path + "service.ip.optionType"); t.Exists() {
		va := res.Get(path + "service.ip.value")
		if t.String() == "variable" {
			data.SurrogateIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SurrogateIp = types.BoolValue(va.Bool())
		}
	}
	data.IdleTime = types.Int64Null()
	data.IdleTimeVariable = types.StringNull()
	if t := res.Get(path + "service.idleTime.optionType"); t.Exists() {
		va := res.Get(path + "service.idleTime.value")
		if t.String() == "variable" {
			data.IdleTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IdleTime = types.Int64Value(va.Int())
		}
	}
	data.DisplayTimeUnit = types.StringNull()
	data.DisplayTimeUnitVariable = types.StringNull()
	if t := res.Get(path + "service.displayTimeUnit.optionType"); t.Exists() {
		va := res.Get(path + "service.displayTimeUnit.value")
		if t.String() == "variable" {
			data.DisplayTimeUnitVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DisplayTimeUnit = types.StringValue(va.String())
		}
	}
	data.IpEnforcedForKnownBrowsers = types.BoolNull()
	data.IpEnforcedForKnownBrowsersVariable = types.StringNull()
	if t := res.Get(path + "service.ipEnforcedForKnownBrowsers.optionType"); t.Exists() {
		va := res.Get(path + "service.ipEnforcedForKnownBrowsers.value")
		if t.String() == "variable" {
			data.IpEnforcedForKnownBrowsersVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpEnforcedForKnownBrowsers = types.BoolValue(va.Bool())
		}
	}
	data.RefreshTime = types.Int64Null()
	data.RefreshTimeVariable = types.StringNull()
	if t := res.Get(path + "service.refreshTime.optionType"); t.Exists() {
		va := res.Get(path + "service.refreshTime.value")
		if t.String() == "variable" {
			data.RefreshTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RefreshTime = types.Int64Value(va.Int())
		}
	}
	data.RefreshTimeUnit = types.StringNull()
	data.RefreshTimeUnitVariable = types.StringNull()
	if t := res.Get(path + "service.refreshTimeUnit.optionType"); t.Exists() {
		va := res.Get(path + "service.refreshTimeUnit.value")
		if t.String() == "variable" {
			data.RefreshTimeUnitVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RefreshTimeUnit = types.StringValue(va.String())
		}
	}
	data.AupEnabled = types.BoolNull()
	data.AupEnabledVariable = types.StringNull()
	if t := res.Get(path + "service.enabled.optionType"); t.Exists() {
		va := res.Get(path + "service.enabled.value")
		if t.String() == "variable" {
			data.AupEnabledVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AupEnabled = types.BoolValue(va.Bool())
		}
	}
	data.BlockInternetUntilAccepted = types.BoolNull()
	data.BlockInternetUntilAcceptedVariable = types.StringNull()
	if t := res.Get(path + "service.blockInternetUntilAccepted.optionType"); t.Exists() {
		va := res.Get(path + "service.blockInternetUntilAccepted.value")
		if t.String() == "variable" {
			data.BlockInternetUntilAcceptedVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.BlockInternetUntilAccepted = types.BoolValue(va.Bool())
		}
	}
	data.ForceSslInspection = types.BoolNull()
	data.ForceSslInspectionVariable = types.StringNull()
	if t := res.Get(path + "service.forceSslInspection.optionType"); t.Exists() {
		va := res.Get(path + "service.forceSslInspection.value")
		if t.String() == "variable" {
			data.ForceSslInspectionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ForceSslInspection = types.BoolValue(va.Bool())
		}
	}
	data.AupTimeout = types.Int64Null()
	data.AupTimeoutVariable = types.StringNull()
	if t := res.Get(path + "service.timeout.optionType"); t.Exists() {
		va := res.Get(path + "service.timeout.value")
		if t.String() == "variable" {
			data.AupTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AupTimeout = types.Int64Value(va.Int())
		}
	}
	data.LocationName = types.StringNull()
	data.LocationNameVariable = types.StringNull()
	if t := res.Get(path + "service.locationName.optionType"); t.Exists() {
		va := res.Get(path + "service.locationName.value")
		if t.String() == "variable" {
			data.LocationNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" || t.String() == "default" {
			data.LocationName = types.StringValue(va.String())
		}
	}
	data.Country = types.BoolNull()
	data.CountryVariable = types.StringNull()
	if t := res.Get(path + "service.country.optionType"); t.Exists() {
		va := res.Get(path + "service.country.value")
		if t.String() == "variable" {
			data.CountryVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Country = types.BoolValue(va.Bool())
		}
	}
	data.EnforceBandwidthControl = types.BoolNull()
	data.EnforceBandwidthControlVariable = types.StringNull()
	if t := res.Get(path + "service.enforceBandwidthControl.optionType"); t.Exists() {
		va := res.Get(path + "service.enforceBandwidthControl.value")
		if t.String() == "variable" {
			data.EnforceBandwidthControlVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EnforceBandwidthControl = types.BoolValue(va.Bool())
		}
	}
	data.DnBandwidth = types.Float64Null()
	data.DnBandwidthVariable = types.StringNull()
	if t := res.Get(path + "service.dnBandwidth.optionType"); t.Exists() {
		va := res.Get(path + "service.dnBandwidth.value")
		if t.String() == "variable" {
			data.DnBandwidthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DnBandwidth = types.Float64Value(va.Float())
		}
	}
	data.UpBandwidth = types.Float64Null()
	data.UpBandwidthVariable = types.StringNull()
	if t := res.Get(path + "service.upBandwidth.optionType"); t.Exists() {
		va := res.Get(path + "service.upBandwidth.value")
		if t.String() == "variable" {
			data.UpBandwidthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.UpBandwidth = types.Float64Value(va.Float())
		}
	}
	if value := res.Get(path + "service.subLocations"); value.Exists() && len(value.Array()) > 0 {
		data.SubLocations = make([]SSEZscalerSubLocations, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SSEZscalerSubLocations{}
			item.Name = types.StringNull()
			item.NameVariable = types.StringNull()
			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "variable" {
					item.NameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			item.AuthRequired = types.BoolNull()
			item.AuthRequiredVariable = types.StringNull()
			if t := v.Get("authRequired.optionType"); t.Exists() {
				va := v.Get("authRequired.value")
				if t.String() == "variable" {
					item.AuthRequiredVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AuthRequired = types.BoolValue(va.Bool())
				}
			}
			item.SurrogateIp = types.BoolNull()
			item.SurrogateIpVariable = types.StringNull()
			if t := v.Get("ip.optionType"); t.Exists() {
				va := v.Get("ip.value")
				if t.String() == "variable" {
					item.SurrogateIpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SurrogateIp = types.BoolValue(va.Bool())
				}
			}
			item.IdleTime = types.Int64Null()
			item.IdleTimeVariable = types.StringNull()
			if t := v.Get("idleTime.optionType"); t.Exists() {
				va := v.Get("idleTime.value")
				if t.String() == "variable" {
					item.IdleTimeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IdleTime = types.Int64Value(va.Int())
				}
			}
			item.DisplayTimeUnit = types.StringNull()
			item.DisplayTimeUnitVariable = types.StringNull()
			if t := v.Get("displayTimeUnit.optionType"); t.Exists() {
				va := v.Get("displayTimeUnit.value")
				if t.String() == "variable" {
					item.DisplayTimeUnitVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.DisplayTimeUnit = types.StringValue(va.String())
				}
			}
			item.IpEnforcedForKnownBrowsers = types.BoolNull()
			item.IpEnforcedForKnownBrowsersVariable = types.StringNull()
			if t := v.Get("ipEnforcedForKnownBrowsers.optionType"); t.Exists() {
				va := v.Get("ipEnforcedForKnownBrowsers.value")
				if t.String() == "variable" {
					item.IpEnforcedForKnownBrowsersVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IpEnforcedForKnownBrowsers = types.BoolValue(va.Bool())
				}
			}
			item.RefreshTime = types.Int64Null()
			item.RefreshTimeVariable = types.StringNull()
			if t := v.Get("refreshTime.optionType"); t.Exists() {
				va := v.Get("refreshTime.value")
				if t.String() == "variable" {
					item.RefreshTimeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RefreshTime = types.Int64Value(va.Int())
				}
			}
			item.RefreshTimeUnit = types.StringNull()
			item.RefreshTimeUnitVariable = types.StringNull()
			if t := v.Get("refreshTimeUnit.optionType"); t.Exists() {
				va := v.Get("refreshTimeUnit.value")
				if t.String() == "variable" {
					item.RefreshTimeUnitVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RefreshTimeUnit = types.StringValue(va.String())
				}
			}
			item.OfwEnabled = types.BoolNull()
			item.OfwEnabledVariable = types.StringNull()
			if t := v.Get("ofwEnabled.optionType"); t.Exists() {
				va := v.Get("ofwEnabled.value")
				if t.String() == "variable" {
					item.OfwEnabledVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.OfwEnabled = types.BoolValue(va.Bool())
				}
			}
			item.CautionEnabled = types.BoolNull()
			item.CautionEnabledVariable = types.StringNull()
			if t := v.Get("cautionEnabled.optionType"); t.Exists() {
				va := v.Get("cautionEnabled.value")
				if t.String() == "variable" {
					item.CautionEnabledVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.CautionEnabled = types.BoolValue(va.Bool())
				}
			}
			item.ServiceVpn = types.SetNull(types.StringType)
			item.ServiceVpnVariable = types.StringNull()
			if t := v.Get("serviceVPN.serviceVpnValue.optionType"); t.Exists() {
				va := v.Get("serviceVPN.serviceVpnValue.value")
				if t.String() == "variable" {
					item.ServiceVpnVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.ServiceVpn = helpers.GetStringSet(va.Array())
				}
			}
			if cValue := v.Get("internalIp"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.InternalIp = make([]SSEZscalerSubLocationsInternalIp, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := SSEZscalerSubLocationsInternalIp{}
					cItem.InternalIpValue = types.StringNull()
					cItem.InternalIpValueVariable = types.StringNull()
					if t := cv.Get("internalIpValue.optionType"); t.Exists() {
						va := cv.Get("internalIpValue.value")
						if t.String() == "variable" {
							cItem.InternalIpValueVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.InternalIpValue = types.StringValue(va.String())
						}
					}
					item.InternalIp = append(item.InternalIp, cItem)
					return true
				})
			}
			item.AupEnabled = types.BoolNull()
			item.AupEnabledVariable = types.StringNull()
			if t := v.Get("aupEnabled.optionType"); t.Exists() {
				va := v.Get("aupEnabled.value")
				if t.String() == "variable" {
					item.AupEnabledVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AupEnabled = types.BoolValue(va.Bool())
				}
			}
			item.BlockInternetUntilAccepted = types.BoolNull()
			item.BlockInternetUntilAcceptedVariable = types.StringNull()
			if t := v.Get("blockInternetUntilAccepted.optionType"); t.Exists() {
				va := v.Get("blockInternetUntilAccepted.value")
				if t.String() == "variable" {
					item.BlockInternetUntilAcceptedVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.BlockInternetUntilAccepted = types.BoolValue(va.Bool())
				}
			}
			item.ForceSslInspection = types.BoolNull()
			item.ForceSslInspectionVariable = types.StringNull()
			if t := v.Get("forceSslInspection.optionType"); t.Exists() {
				va := v.Get("forceSslInspection.value")
				if t.String() == "variable" {
					item.ForceSslInspectionVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.ForceSslInspection = types.BoolValue(va.Bool())
				}
			}
			item.AupTimeout = types.Int64Null()
			item.AupTimeoutVariable = types.StringNull()
			if t := v.Get("timeout.optionType"); t.Exists() {
				va := v.Get("timeout.value")
				if t.String() == "variable" {
					item.AupTimeoutVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AupTimeout = types.Int64Value(va.Int())
				}
			}
			item.EnforceBandwidthControl = types.StringNull()
			item.EnforceBandwidthControlVariable = types.StringNull()
			if t := v.Get("enforceBandwidthControl.optionType"); t.Exists() {
				va := v.Get("enforceBandwidthControl.value")
				if t.String() == "variable" {
					item.EnforceBandwidthControlVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.EnforceBandwidthControl = types.StringValue(va.String())
				}
			}
			item.UpBandwidth = types.Float64Null()
			item.UpBandwidthVariable = types.StringNull()
			if t := v.Get("upBandwidth.bandwidthValue.optionType"); t.Exists() {
				va := v.Get("upBandwidth.bandwidthValue.value")
				if t.String() == "variable" {
					item.UpBandwidthVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.UpBandwidth = types.Float64Value(va.Float())
				}
			}
			item.DnBandwidth = types.Float64Null()
			item.DnBandwidthVariable = types.StringNull()
			if t := v.Get("dnBandwidth.bandwidthValue.optionType"); t.Exists() {
				va := v.Get("dnBandwidth.bandwidthValue.value")
				if t.String() == "variable" {
					item.DnBandwidthVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.DnBandwidth = types.Float64Value(va.Float())
				}
			}
			data.SubLocations = append(data.SubLocations, item)
			return true
		})
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
	if value := res.Get(path + "tracker"); value.Exists() && len(value.Array()) > 0 {
		data.Trackers = make([]SSEZscalerTrackers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SSEZscalerTrackers{}
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
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SSEZscaler) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.SrcVpn = types.BoolNull()

	if t := res.Get(path + "interfaceMetadataSharing.srcVpn.optionType"); t.Exists() {
		va := res.Get(path + "interfaceMetadataSharing.srcVpn.value")
		if t.String() == "global" {
			data.SrcVpn = types.BoolValue(va.Bool())
		}
	}
	for i := range data.Interfaces {
		keys := [...]string{"ifName"}
		keyValues := [...]string{data.Interfaces[i].InterfaceName.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "interface").ForEach(
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
		data.Interfaces[i].InterfaceName = types.StringNull()

		if t := r.Get("ifName.optionType"); t.Exists() {
			va := r.Get("ifName.value")
			if t.String() == "global" {
				data.Interfaces[i].InterfaceName = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].Auto = types.BoolNull()

		if t := r.Get("auto.optionType"); t.Exists() {
			va := r.Get("auto.value")
			if t.String() == "global" {
				data.Interfaces[i].Auto = types.BoolValue(va.Bool())
			}
		}
		data.Interfaces[i].Shutdown = types.BoolNull()

		if t := r.Get("shutdown.optionType"); t.Exists() {
			va := r.Get("shutdown.value")
			if t.String() == "global" {
				data.Interfaces[i].Shutdown = types.BoolValue(va.Bool())
			}
		}
		data.Interfaces[i].InterfaceDescription = types.StringNull()
		data.Interfaces[i].InterfaceDescriptionVariable = types.StringNull()
		if t := r.Get("description.optionType"); t.Exists() {
			va := r.Get("description.value")
			if t.String() == "variable" {
				data.Interfaces[i].InterfaceDescriptionVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].InterfaceDescription = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].Unnumbered = types.BoolNull()

		if t := r.Get("unnumbered.optionType"); t.Exists() {
			va := r.Get("unnumbered.value")
			if t.String() == "global" {
				data.Interfaces[i].Unnumbered = types.BoolValue(va.Bool())
			}
		}
		data.Interfaces[i].Ipv4Address = types.StringNull()
		data.Interfaces[i].Ipv4AddressVariable = types.StringNull()
		if t := r.Get("address.optionType"); t.Exists() {
			va := r.Get("address.value")
			if t.String() == "variable" {
				data.Interfaces[i].Ipv4AddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].Ipv4Address = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].TunnelSource = types.StringNull()
		data.Interfaces[i].TunnelSourceVariable = types.StringNull()
		if t := r.Get("tunnelSource.optionType"); t.Exists() {
			va := r.Get("tunnelSource.value")
			if t.String() == "variable" {
				data.Interfaces[i].TunnelSourceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].TunnelSource = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].TunnelSourceInterface = types.StringNull()
		data.Interfaces[i].TunnelSourceInterfaceVariable = types.StringNull()
		if t := r.Get("tunnelSourceInterface.optionType"); t.Exists() {
			va := r.Get("tunnelSourceInterface.value")
			if t.String() == "variable" {
				data.Interfaces[i].TunnelSourceInterfaceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].TunnelSourceInterface = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].TunnelRouteVia = types.StringNull()
		data.Interfaces[i].TunnelRouteViaVariable = types.StringNull()
		if t := r.Get("tunnelRouteVia.optionType"); t.Exists() {
			va := r.Get("tunnelRouteVia.value")
			if t.String() == "variable" {
				data.Interfaces[i].TunnelRouteViaVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].TunnelRouteVia = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].TunnelDestination = types.StringNull()
		data.Interfaces[i].TunnelDestinationVariable = types.StringNull()
		if t := r.Get("tunnelDestination.optionType"); t.Exists() {
			va := r.Get("tunnelDestination.value")
			if t.String() == "variable" {
				data.Interfaces[i].TunnelDestinationVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].TunnelDestination = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].TunnelSet = types.StringNull()

		if t := r.Get("tunnelSet.optionType"); t.Exists() {
			va := r.Get("tunnelSet.value")
			if t.String() == "global" {
				data.Interfaces[i].TunnelSet = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].TunnelDcPreference = types.StringNull()

		if t := r.Get("tunnelDcPreference.optionType"); t.Exists() {
			va := r.Get("tunnelDcPreference.value")
			if t.String() == "global" {
				data.Interfaces[i].TunnelDcPreference = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].TcpMssAdjust = types.Int64Null()
		data.Interfaces[i].TcpMssAdjustVariable = types.StringNull()
		if t := r.Get("tcpMssAdjust.optionType"); t.Exists() {
			va := r.Get("tcpMssAdjust.value")
			if t.String() == "variable" {
				data.Interfaces[i].TcpMssAdjustVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].TcpMssAdjust = types.Int64Value(va.Int())
			}
		}
		data.Interfaces[i].Mtu = types.Int64Null()
		data.Interfaces[i].MtuVariable = types.StringNull()
		if t := r.Get("mtu.optionType"); t.Exists() {
			va := r.Get("mtu.value")
			if t.String() == "variable" {
				data.Interfaces[i].MtuVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].Mtu = types.Int64Value(va.Int())
			}
		}
		data.Interfaces[i].DpdInterval = types.Int64Null()
		data.Interfaces[i].DpdIntervalVariable = types.StringNull()
		if t := r.Get("dpdInterval.optionType"); t.Exists() {
			va := r.Get("dpdInterval.value")
			if t.String() == "variable" {
				data.Interfaces[i].DpdIntervalVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].DpdInterval = types.Int64Value(va.Int())
			}
		}
		data.Interfaces[i].DpdRetries = types.Int64Null()
		data.Interfaces[i].DpdRetriesVariable = types.StringNull()
		if t := r.Get("dpdRetries.optionType"); t.Exists() {
			va := r.Get("dpdRetries.value")
			if t.String() == "variable" {
				data.Interfaces[i].DpdRetriesVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].DpdRetries = types.Int64Value(va.Int())
			}
		}
		data.Interfaces[i].IkeVersion = types.Int64Null()
		data.Interfaces[i].IkeVersionVariable = types.StringNull()
		if t := r.Get("ikeVersion.optionType"); t.Exists() {
			va := r.Get("ikeVersion.value")
			if t.String() == "variable" {
				data.Interfaces[i].IkeVersionVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].IkeVersion = types.Int64Value(va.Int())
			}
		}
		data.Interfaces[i].PreSharedSecret = types.StringNull()
		data.Interfaces[i].PreSharedSecretVariable = types.StringNull()
		if t := r.Get("preSharedSecret.optionType"); t.Exists() {
			va := r.Get("preSharedSecret.value")
			if t.String() == "variable" {
				data.Interfaces[i].PreSharedSecretVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].PreSharedSecret = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].IkeRekeyInterval = types.Int64Null()
		data.Interfaces[i].IkeRekeyIntervalVariable = types.StringNull()
		if t := r.Get("ikeRekeyInterval.optionType"); t.Exists() {
			va := r.Get("ikeRekeyInterval.value")
			if t.String() == "variable" {
				data.Interfaces[i].IkeRekeyIntervalVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].IkeRekeyInterval = types.Int64Value(va.Int())
			}
		}
		data.Interfaces[i].IkeCiphersuite = types.StringNull()
		data.Interfaces[i].IkeCiphersuiteVariable = types.StringNull()
		if t := r.Get("ikeCiphersuite.optionType"); t.Exists() {
			va := r.Get("ikeCiphersuite.value")
			if t.String() == "variable" {
				data.Interfaces[i].IkeCiphersuiteVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].IkeCiphersuite = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].IkeGroup = types.StringNull()
		data.Interfaces[i].IkeGroupVariable = types.StringNull()
		if t := r.Get("ikeGroup.optionType"); t.Exists() {
			va := r.Get("ikeGroup.value")
			if t.String() == "variable" {
				data.Interfaces[i].IkeGroupVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].IkeGroup = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].PreSharedKeyDynamic = types.BoolNull()

		if t := r.Get("preSharedKeyDynamic.optionType"); t.Exists() {
			va := r.Get("preSharedKeyDynamic.value")
			if t.String() == "global" {
				data.Interfaces[i].PreSharedKeyDynamic = types.BoolValue(va.Bool())
			}
		}
		data.Interfaces[i].IkeLocalId = types.StringNull()
		data.Interfaces[i].IkeLocalIdVariable = types.StringNull()
		if t := r.Get("ikeLocalId.optionType"); t.Exists() {
			va := r.Get("ikeLocalId.value")
			if t.String() == "variable" {
				data.Interfaces[i].IkeLocalIdVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].IkeLocalId = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].IkeRemoteId = types.StringNull()
		data.Interfaces[i].IkeRemoteIdVariable = types.StringNull()
		if t := r.Get("ikeRemoteId.optionType"); t.Exists() {
			va := r.Get("ikeRemoteId.value")
			if t.String() == "variable" {
				data.Interfaces[i].IkeRemoteIdVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].IkeRemoteId = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].IpsecRekeyInterval = types.Int64Null()
		data.Interfaces[i].IpsecRekeyIntervalVariable = types.StringNull()
		if t := r.Get("ipsecRekeyInterval.optionType"); t.Exists() {
			va := r.Get("ipsecRekeyInterval.value")
			if t.String() == "variable" {
				data.Interfaces[i].IpsecRekeyIntervalVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].IpsecRekeyInterval = types.Int64Value(va.Int())
			}
		}
		data.Interfaces[i].IpsecReplayWindow = types.Int64Null()
		data.Interfaces[i].IpsecReplayWindowVariable = types.StringNull()
		if t := r.Get("ipsecReplayWindow.optionType"); t.Exists() {
			va := r.Get("ipsecReplayWindow.value")
			if t.String() == "variable" {
				data.Interfaces[i].IpsecReplayWindowVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].IpsecReplayWindow = types.Int64Value(va.Int())
			}
		}
		data.Interfaces[i].IpsecCiphersuite = types.StringNull()
		data.Interfaces[i].IpsecCiphersuiteVariable = types.StringNull()
		if t := r.Get("ipsecCiphersuite.optionType"); t.Exists() {
			va := r.Get("ipsecCiphersuite.value")
			if t.String() == "variable" {
				data.Interfaces[i].IpsecCiphersuiteVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].IpsecCiphersuite = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].PerfectForwardSecrecy = types.StringNull()
		data.Interfaces[i].PerfectForwardSecrecyVariable = types.StringNull()
		if t := r.Get("perfectForwardSecrecy.optionType"); t.Exists() {
			va := r.Get("perfectForwardSecrecy.value")
			if t.String() == "variable" {
				data.Interfaces[i].PerfectForwardSecrecyVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].PerfectForwardSecrecy = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].Tracker = types.StringNull()

		if t := r.Get("tracker.optionType"); t.Exists() {
			va := r.Get("tracker.value")
			if t.String() == "global" {
				data.Interfaces[i].Tracker = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].TrackEnable = types.BoolNull()

		if t := r.Get("trackEnable.optionType"); t.Exists() {
			va := r.Get("trackEnable.value")
			if t.String() == "global" {
				data.Interfaces[i].TrackEnable = types.BoolValue(va.Bool())
			}
		}
		data.Interfaces[i].TunnelPublicIp = types.StringNull()
		data.Interfaces[i].TunnelPublicIpVariable = types.StringNull()
		if t := r.Get("tunnelPublicIp.optionType"); t.Exists() {
			va := r.Get("tunnelPublicIp.value")
			if t.String() == "variable" {
				data.Interfaces[i].TunnelPublicIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].TunnelPublicIp = types.StringValue(va.String())
			}
		}
	}
	for i := range data.InterfacePairs {
		keys := [...]string{"activeInterface", "activeInterfaceWeight", "backupInterface", "backupInterfaceWeight"}
		keyValues := [...]string{data.InterfacePairs[i].ActiveInterface.ValueString(), strconv.FormatInt(data.InterfacePairs[i].ActiveInterfaceWeight.ValueInt64(), 10), data.InterfacePairs[i].BackupInterface.ValueString(), strconv.FormatInt(data.InterfacePairs[i].BackupInterfaceWeight.ValueInt64(), 10)}
		keyValuesVariables := [...]string{"", "", "", ""}

		var r gjson.Result
		res.Get(path + "service.interfacePair").ForEach(
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
		data.InterfacePairs[i].ActiveInterface = types.StringNull()

		if t := r.Get("activeInterface.optionType"); t.Exists() {
			va := r.Get("activeInterface.value")
			if t.String() == "global" {
				data.InterfacePairs[i].ActiveInterface = types.StringValue(va.String())
			}
		}
		data.InterfacePairs[i].ActiveInterfaceWeight = types.Int64Null()

		if t := r.Get("activeInterfaceWeight.optionType"); t.Exists() {
			va := r.Get("activeInterfaceWeight.value")
			if t.String() == "global" {
				data.InterfacePairs[i].ActiveInterfaceWeight = types.Int64Value(va.Int())
			}
		}
		data.InterfacePairs[i].BackupInterface = types.StringNull()

		if t := r.Get("backupInterface.optionType"); t.Exists() {
			va := r.Get("backupInterface.value")
			if t.String() == "global" {
				data.InterfacePairs[i].BackupInterface = types.StringValue(va.String())
			}
		}
		data.InterfacePairs[i].BackupInterfaceWeight = types.Int64Null()

		if t := r.Get("backupInterfaceWeight.optionType"); t.Exists() {
			va := r.Get("backupInterfaceWeight.value")
			if t.String() == "global" {
				data.InterfacePairs[i].BackupInterfaceWeight = types.Int64Value(va.Int())
			}
		}
	}
	data.AuthRequired = types.BoolNull()
	data.AuthRequiredVariable = types.StringNull()
	if t := res.Get(path + "service.authRequired.optionType"); t.Exists() {
		va := res.Get(path + "service.authRequired.value")
		if t.String() == "variable" {
			data.AuthRequiredVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AuthRequired = types.BoolValue(va.Bool())
		}
	}
	data.XffForwardEnabled = types.BoolNull()
	data.XffForwardEnabledVariable = types.StringNull()
	if t := res.Get(path + "service.xffForwardEnabled.optionType"); t.Exists() {
		va := res.Get(path + "service.xffForwardEnabled.value")
		if t.String() == "variable" {
			data.XffForwardEnabledVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.XffForwardEnabled = types.BoolValue(va.Bool())
		}
	}
	data.OfwEnabled = types.BoolNull()
	data.OfwEnabledVariable = types.StringNull()
	if t := res.Get(path + "service.ofwEnabled.optionType"); t.Exists() {
		va := res.Get(path + "service.ofwEnabled.value")
		if t.String() == "variable" {
			data.OfwEnabledVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.OfwEnabled = types.BoolValue(va.Bool())
		}
	}
	data.IpsControl = types.BoolNull()
	data.IpsControlVariable = types.StringNull()
	if t := res.Get(path + "service.ipsControl.optionType"); t.Exists() {
		va := res.Get(path + "service.ipsControl.value")
		if t.String() == "variable" {
			data.IpsControlVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsControl = types.BoolValue(va.Bool())
		}
	}
	data.CautionEnabled = types.BoolNull()
	data.CautionEnabledVariable = types.StringNull()
	if t := res.Get(path + "service.cautionEnabled.optionType"); t.Exists() {
		va := res.Get(path + "service.cautionEnabled.value")
		if t.String() == "variable" {
			data.CautionEnabledVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.CautionEnabled = types.BoolValue(va.Bool())
		}
	}
	data.PrimaryDataCenter = types.StringNull()
	data.PrimaryDataCenterVariable = types.StringNull()
	if t := res.Get(path + "service.primaryDataCenter.optionType"); t.Exists() {
		va := res.Get(path + "service.primaryDataCenter.value")
		if t.String() == "variable" {
			data.PrimaryDataCenterVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PrimaryDataCenter = types.StringValue(va.String())
		}
	}
	data.SecondaryDataCenter = types.StringNull()
	data.SecondaryDataCenterVariable = types.StringNull()
	if t := res.Get(path + "service.secondaryDataCenter.optionType"); t.Exists() {
		va := res.Get(path + "service.secondaryDataCenter.value")
		if t.String() == "variable" {
			data.SecondaryDataCenterVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SecondaryDataCenter = types.StringValue(va.String())
		}
	}
	data.SurrogateIp = types.BoolNull()
	data.SurrogateIpVariable = types.StringNull()
	if t := res.Get(path + "service.ip.optionType"); t.Exists() {
		va := res.Get(path + "service.ip.value")
		if t.String() == "variable" {
			data.SurrogateIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SurrogateIp = types.BoolValue(va.Bool())
		}
	}
	data.IdleTime = types.Int64Null()
	data.IdleTimeVariable = types.StringNull()
	if t := res.Get(path + "service.idleTime.optionType"); t.Exists() {
		va := res.Get(path + "service.idleTime.value")
		if t.String() == "variable" {
			data.IdleTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IdleTime = types.Int64Value(va.Int())
		}
	}
	data.DisplayTimeUnit = types.StringNull()
	data.DisplayTimeUnitVariable = types.StringNull()
	if t := res.Get(path + "service.displayTimeUnit.optionType"); t.Exists() {
		va := res.Get(path + "service.displayTimeUnit.value")
		if t.String() == "variable" {
			data.DisplayTimeUnitVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DisplayTimeUnit = types.StringValue(va.String())
		}
	}
	data.IpEnforcedForKnownBrowsers = types.BoolNull()
	data.IpEnforcedForKnownBrowsersVariable = types.StringNull()
	if t := res.Get(path + "service.ipEnforcedForKnownBrowsers.optionType"); t.Exists() {
		va := res.Get(path + "service.ipEnforcedForKnownBrowsers.value")
		if t.String() == "variable" {
			data.IpEnforcedForKnownBrowsersVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpEnforcedForKnownBrowsers = types.BoolValue(va.Bool())
		}
	}
	data.RefreshTime = types.Int64Null()
	data.RefreshTimeVariable = types.StringNull()
	if t := res.Get(path + "service.refreshTime.optionType"); t.Exists() {
		va := res.Get(path + "service.refreshTime.value")
		if t.String() == "variable" {
			data.RefreshTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RefreshTime = types.Int64Value(va.Int())
		}
	}
	data.RefreshTimeUnit = types.StringNull()
	data.RefreshTimeUnitVariable = types.StringNull()
	if t := res.Get(path + "service.refreshTimeUnit.optionType"); t.Exists() {
		va := res.Get(path + "service.refreshTimeUnit.value")
		if t.String() == "variable" {
			data.RefreshTimeUnitVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RefreshTimeUnit = types.StringValue(va.String())
		}
	}
	data.AupEnabled = types.BoolNull()
	data.AupEnabledVariable = types.StringNull()
	if t := res.Get(path + "service.enabled.optionType"); t.Exists() {
		va := res.Get(path + "service.enabled.value")
		if t.String() == "variable" {
			data.AupEnabledVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AupEnabled = types.BoolValue(va.Bool())
		}
	}
	data.BlockInternetUntilAccepted = types.BoolNull()
	data.BlockInternetUntilAcceptedVariable = types.StringNull()
	if t := res.Get(path + "service.blockInternetUntilAccepted.optionType"); t.Exists() {
		va := res.Get(path + "service.blockInternetUntilAccepted.value")
		if t.String() == "variable" {
			data.BlockInternetUntilAcceptedVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.BlockInternetUntilAccepted = types.BoolValue(va.Bool())
		}
	}
	data.ForceSslInspection = types.BoolNull()
	data.ForceSslInspectionVariable = types.StringNull()
	if t := res.Get(path + "service.forceSslInspection.optionType"); t.Exists() {
		va := res.Get(path + "service.forceSslInspection.value")
		if t.String() == "variable" {
			data.ForceSslInspectionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ForceSslInspection = types.BoolValue(va.Bool())
		}
	}
	data.AupTimeout = types.Int64Null()
	data.AupTimeoutVariable = types.StringNull()
	if t := res.Get(path + "service.timeout.optionType"); t.Exists() {
		va := res.Get(path + "service.timeout.value")
		if t.String() == "variable" {
			data.AupTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
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
		} else if t.String() == "global" || (t.String() == "default" && !tempLocationName.IsNull()) {
			data.LocationName = types.StringValue(va.String())
		}
	}
	data.Country = types.BoolNull()
	data.CountryVariable = types.StringNull()
	if t := res.Get(path + "service.country.optionType"); t.Exists() {
		va := res.Get(path + "service.country.value")
		if t.String() == "variable" {
			data.CountryVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Country = types.BoolValue(va.Bool())
		}
	}
	data.EnforceBandwidthControl = types.BoolNull()
	data.EnforceBandwidthControlVariable = types.StringNull()
	if t := res.Get(path + "service.enforceBandwidthControl.optionType"); t.Exists() {
		va := res.Get(path + "service.enforceBandwidthControl.value")
		if t.String() == "variable" {
			data.EnforceBandwidthControlVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EnforceBandwidthControl = types.BoolValue(va.Bool())
		}
	}
	data.DnBandwidth = types.Float64Null()
	data.DnBandwidthVariable = types.StringNull()
	if t := res.Get(path + "service.dnBandwidth.optionType"); t.Exists() {
		va := res.Get(path + "service.dnBandwidth.value")
		if t.String() == "variable" {
			data.DnBandwidthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DnBandwidth = types.Float64Value(va.Float())
		}
	}
	data.UpBandwidth = types.Float64Null()
	data.UpBandwidthVariable = types.StringNull()
	if t := res.Get(path + "service.upBandwidth.optionType"); t.Exists() {
		va := res.Get(path + "service.upBandwidth.value")
		if t.String() == "variable" {
			data.UpBandwidthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.UpBandwidth = types.Float64Value(va.Float())
		}
	}
	for i := range data.SubLocations {
		keys := [...]string{"name"}
		keyValues := [...]string{data.SubLocations[i].Name.ValueString()}
		keyValuesVariables := [...]string{data.SubLocations[i].NameVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "service.subLocations").ForEach(
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
		data.SubLocations[i].Name = types.StringNull()
		data.SubLocations[i].NameVariable = types.StringNull()
		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "variable" {
				data.SubLocations[i].NameVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].Name = types.StringValue(va.String())
			}
		}
		data.SubLocations[i].AuthRequired = types.BoolNull()
		data.SubLocations[i].AuthRequiredVariable = types.StringNull()
		if t := r.Get("authRequired.optionType"); t.Exists() {
			va := r.Get("authRequired.value")
			if t.String() == "variable" {
				data.SubLocations[i].AuthRequiredVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].AuthRequired = types.BoolValue(va.Bool())
			}
		}
		data.SubLocations[i].SurrogateIp = types.BoolNull()
		data.SubLocations[i].SurrogateIpVariable = types.StringNull()
		if t := r.Get("ip.optionType"); t.Exists() {
			va := r.Get("ip.value")
			if t.String() == "variable" {
				data.SubLocations[i].SurrogateIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].SurrogateIp = types.BoolValue(va.Bool())
			}
		}
		data.SubLocations[i].IdleTime = types.Int64Null()
		data.SubLocations[i].IdleTimeVariable = types.StringNull()
		if t := r.Get("idleTime.optionType"); t.Exists() {
			va := r.Get("idleTime.value")
			if t.String() == "variable" {
				data.SubLocations[i].IdleTimeVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].IdleTime = types.Int64Value(va.Int())
			}
		}
		data.SubLocations[i].DisplayTimeUnit = types.StringNull()
		data.SubLocations[i].DisplayTimeUnitVariable = types.StringNull()
		if t := r.Get("displayTimeUnit.optionType"); t.Exists() {
			va := r.Get("displayTimeUnit.value")
			if t.String() == "variable" {
				data.SubLocations[i].DisplayTimeUnitVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].DisplayTimeUnit = types.StringValue(va.String())
			}
		}
		data.SubLocations[i].IpEnforcedForKnownBrowsers = types.BoolNull()
		data.SubLocations[i].IpEnforcedForKnownBrowsersVariable = types.StringNull()
		if t := r.Get("ipEnforcedForKnownBrowsers.optionType"); t.Exists() {
			va := r.Get("ipEnforcedForKnownBrowsers.value")
			if t.String() == "variable" {
				data.SubLocations[i].IpEnforcedForKnownBrowsersVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].IpEnforcedForKnownBrowsers = types.BoolValue(va.Bool())
			}
		}
		data.SubLocations[i].RefreshTime = types.Int64Null()
		data.SubLocations[i].RefreshTimeVariable = types.StringNull()
		if t := r.Get("refreshTime.optionType"); t.Exists() {
			va := r.Get("refreshTime.value")
			if t.String() == "variable" {
				data.SubLocations[i].RefreshTimeVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].RefreshTime = types.Int64Value(va.Int())
			}
		}
		data.SubLocations[i].RefreshTimeUnit = types.StringNull()
		data.SubLocations[i].RefreshTimeUnitVariable = types.StringNull()
		if t := r.Get("refreshTimeUnit.optionType"); t.Exists() {
			va := r.Get("refreshTimeUnit.value")
			if t.String() == "variable" {
				data.SubLocations[i].RefreshTimeUnitVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].RefreshTimeUnit = types.StringValue(va.String())
			}
		}
		data.SubLocations[i].OfwEnabled = types.BoolNull()
		data.SubLocations[i].OfwEnabledVariable = types.StringNull()
		if t := r.Get("ofwEnabled.optionType"); t.Exists() {
			va := r.Get("ofwEnabled.value")
			if t.String() == "variable" {
				data.SubLocations[i].OfwEnabledVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].OfwEnabled = types.BoolValue(va.Bool())
			}
		}
		data.SubLocations[i].CautionEnabled = types.BoolNull()
		data.SubLocations[i].CautionEnabledVariable = types.StringNull()
		if t := r.Get("cautionEnabled.optionType"); t.Exists() {
			va := r.Get("cautionEnabled.value")
			if t.String() == "variable" {
				data.SubLocations[i].CautionEnabledVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].CautionEnabled = types.BoolValue(va.Bool())
			}
		}
		data.SubLocations[i].ServiceVpn = types.SetNull(types.StringType)
		data.SubLocations[i].ServiceVpnVariable = types.StringNull()
		if t := r.Get("serviceVPN.serviceVpnValue.optionType"); t.Exists() {
			va := r.Get("serviceVPN.serviceVpnValue.value")
			if t.String() == "variable" {
				data.SubLocations[i].ServiceVpnVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].ServiceVpn = helpers.GetStringSet(va.Array())
			}
		}
		for ci := range data.SubLocations[i].InternalIp {
			keys := [...]string{"internalIpValue"}
			keyValues := [...]string{data.SubLocations[i].InternalIp[ci].InternalIpValue.ValueString()}
			keyValuesVariables := [...]string{data.SubLocations[i].InternalIp[ci].InternalIpValueVariable.ValueString()}

			var cr gjson.Result
			r.Get("internalIp").ForEach(
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
			data.SubLocations[i].InternalIp[ci].InternalIpValue = types.StringNull()
			data.SubLocations[i].InternalIp[ci].InternalIpValueVariable = types.StringNull()
			if t := cr.Get("internalIpValue.optionType"); t.Exists() {
				va := cr.Get("internalIpValue.value")
				if t.String() == "variable" {
					data.SubLocations[i].InternalIp[ci].InternalIpValueVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.SubLocations[i].InternalIp[ci].InternalIpValue = types.StringValue(va.String())
				}
			}
		}
		data.SubLocations[i].AupEnabled = types.BoolNull()
		data.SubLocations[i].AupEnabledVariable = types.StringNull()
		if t := r.Get("aupEnabled.optionType"); t.Exists() {
			va := r.Get("aupEnabled.value")
			if t.String() == "variable" {
				data.SubLocations[i].AupEnabledVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].AupEnabled = types.BoolValue(va.Bool())
			}
		}
		data.SubLocations[i].BlockInternetUntilAccepted = types.BoolNull()
		data.SubLocations[i].BlockInternetUntilAcceptedVariable = types.StringNull()
		if t := r.Get("blockInternetUntilAccepted.optionType"); t.Exists() {
			va := r.Get("blockInternetUntilAccepted.value")
			if t.String() == "variable" {
				data.SubLocations[i].BlockInternetUntilAcceptedVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].BlockInternetUntilAccepted = types.BoolValue(va.Bool())
			}
		}
		data.SubLocations[i].ForceSslInspection = types.BoolNull()
		data.SubLocations[i].ForceSslInspectionVariable = types.StringNull()
		if t := r.Get("forceSslInspection.optionType"); t.Exists() {
			va := r.Get("forceSslInspection.value")
			if t.String() == "variable" {
				data.SubLocations[i].ForceSslInspectionVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].ForceSslInspection = types.BoolValue(va.Bool())
			}
		}
		data.SubLocations[i].AupTimeout = types.Int64Null()
		data.SubLocations[i].AupTimeoutVariable = types.StringNull()
		if t := r.Get("timeout.optionType"); t.Exists() {
			va := r.Get("timeout.value")
			if t.String() == "variable" {
				data.SubLocations[i].AupTimeoutVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].AupTimeout = types.Int64Value(va.Int())
			}
		}
		data.SubLocations[i].EnforceBandwidthControl = types.StringNull()
		data.SubLocations[i].EnforceBandwidthControlVariable = types.StringNull()
		if t := r.Get("enforceBandwidthControl.optionType"); t.Exists() {
			va := r.Get("enforceBandwidthControl.value")
			if t.String() == "variable" {
				data.SubLocations[i].EnforceBandwidthControlVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].EnforceBandwidthControl = types.StringValue(va.String())
			}
		}
		data.SubLocations[i].UpBandwidth = types.Float64Null()
		data.SubLocations[i].UpBandwidthVariable = types.StringNull()
		if t := r.Get("upBandwidth.bandwidthValue.optionType"); t.Exists() {
			va := r.Get("upBandwidth.bandwidthValue.value")
			if t.String() == "variable" {
				data.SubLocations[i].UpBandwidthVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].UpBandwidth = types.Float64Value(va.Float())
			}
		}
		data.SubLocations[i].DnBandwidth = types.Float64Null()
		data.SubLocations[i].DnBandwidthVariable = types.StringNull()
		if t := r.Get("dnBandwidth.bandwidthValue.optionType"); t.Exists() {
			va := r.Get("dnBandwidth.bandwidthValue.value")
			if t.String() == "variable" {
				data.SubLocations[i].DnBandwidthVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SubLocations[i].DnBandwidth = types.Float64Value(va.Float())
			}
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
	for i := range data.Trackers {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Trackers[i].Name.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "tracker").ForEach(
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
		data.Trackers[i].Name = types.StringNull()

		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "global" {
				data.Trackers[i].Name = types.StringValue(va.String())
			}
		}
		data.Trackers[i].EndpointApiUrl = types.StringNull()
		data.Trackers[i].EndpointApiUrlVariable = types.StringNull()
		if t := r.Get("endpointApiUrl.optionType"); t.Exists() {
			va := r.Get("endpointApiUrl.value")
			if t.String() == "variable" {
				data.Trackers[i].EndpointApiUrlVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Trackers[i].EndpointApiUrl = types.StringValue(va.String())
			}
		}
		data.Trackers[i].Threshold = types.Int64Null()
		data.Trackers[i].ThresholdVariable = types.StringNull()
		if t := r.Get("threshold.optionType"); t.Exists() {
			va := r.Get("threshold.value")
			if t.String() == "variable" {
				data.Trackers[i].ThresholdVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Trackers[i].Threshold = types.Int64Value(va.Int())
			}
		}
		data.Trackers[i].Interval = types.Int64Null()
		data.Trackers[i].IntervalVariable = types.StringNull()
		if t := r.Get("interval.optionType"); t.Exists() {
			va := r.Get("interval.value")
			if t.String() == "variable" {
				data.Trackers[i].IntervalVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Trackers[i].Interval = types.Int64Value(va.Int())
			}
		}
		data.Trackers[i].Multiplier = types.Int64Null()
		data.Trackers[i].MultiplierVariable = types.StringNull()
		if t := r.Get("multiplier.optionType"); t.Exists() {
			va := r.Get("multiplier.value")
			if t.String() == "variable" {
				data.Trackers[i].MultiplierVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Trackers[i].Multiplier = types.Int64Value(va.Int())
			}
		}
	}
}

// End of section. //template:end updateFromBody
