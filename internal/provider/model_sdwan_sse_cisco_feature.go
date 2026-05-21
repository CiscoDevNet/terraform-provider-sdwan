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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type SSECisco struct {
	Id                      types.String             `tfsdk:"id"`
	Version                 types.Int64              `tfsdk:"version"`
	Name                    types.String             `tfsdk:"name"`
	Description             types.String             `tfsdk:"description"`
	FeatureProfileId        types.String             `tfsdk:"feature_profile_id"`
	ContextSharingForVpn    types.Bool               `tfsdk:"context_sharing_for_vpn"`
	ContextSharingForSgt    types.Bool               `tfsdk:"context_sharing_for_sgt"`
	Interfaces              []SSECiscoInterfaces     `tfsdk:"interfaces"`
	InterfacePairs          []SSECiscoInterfacePairs `tfsdk:"interface_pairs"`
	Region                  types.String             `tfsdk:"region"`
	RegionVariable          types.String             `tfsdk:"region_variable"`
	TrackerSourceIp         types.String             `tfsdk:"tracker_source_ip"`
	TrackerSourceIpVariable types.String             `tfsdk:"tracker_source_ip_variable"`
	Trackers                []SSECiscoTrackers       `tfsdk:"trackers"`
}

type SSECiscoInterfaces struct {
	InterfaceName                 types.String `tfsdk:"interface_name"`
	Shutdown                      types.Bool   `tfsdk:"shutdown"`
	ShutdownVariable              types.String `tfsdk:"shutdown_variable"`
	TunnelSourceInterface         types.String `tfsdk:"tunnel_source_interface"`
	TunnelSourceInterfaceVariable types.String `tfsdk:"tunnel_source_interface_variable"`
	TunnelRouteVia                types.String `tfsdk:"tunnel_route_via"`
	TunnelRouteViaVariable        types.String `tfsdk:"tunnel_route_via_variable"`
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
	IkeRekeyInterval              types.Int64  `tfsdk:"ike_rekey_interval"`
	IkeRekeyIntervalVariable      types.String `tfsdk:"ike_rekey_interval_variable"`
	IkeCiphersuite                types.String `tfsdk:"ike_ciphersuite"`
	IkeCiphersuiteVariable        types.String `tfsdk:"ike_ciphersuite_variable"`
	IkeGroup                      types.String `tfsdk:"ike_group"`
	IkeGroupVariable              types.String `tfsdk:"ike_group_variable"`
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
	TrackEnableVariable           types.String `tfsdk:"track_enable_variable"`
}

type SSECiscoInterfacePairs struct {
	ActiveInterface       types.String `tfsdk:"active_interface"`
	ActiveInterfaceWeight types.Int64  `tfsdk:"active_interface_weight"`
	BackupInterface       types.String `tfsdk:"backup_interface"`
	BackupInterfaceWeight types.Int64  `tfsdk:"backup_interface_weight"`
}

type SSECiscoTrackers struct {
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
func (data SSECisco) getModel() string {
	return "sse_cisco"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SSECisco) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/sse/%v/cisco", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SSECisco) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {
		body, _ = sjson.Set(body, path+"sse_instance.optionType", "global")
		body, _ = sjson.Set(body, path+"sse_instance.value", "Cisco-Secure-Access")
	}
	if data.ContextSharingForVpn.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"contextSharingForVpn.optionType", "default")
			body, _ = sjson.Set(body, path+"contextSharingForVpn.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"contextSharingForVpn.optionType", "global")
			body, _ = sjson.Set(body, path+"contextSharingForVpn.value", data.ContextSharingForVpn.ValueBool())
		}
	}
	if data.ContextSharingForSgt.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"contextSharingForSgt.optionType", "default")
			body, _ = sjson.Set(body, path+"contextSharingForSgt.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"contextSharingForSgt.optionType", "global")
			body, _ = sjson.Set(body, path+"contextSharingForSgt.value", data.ContextSharingForSgt.ValueBool())
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

			if !item.ShutdownVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "shutdown.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "shutdown.value", item.ShutdownVariable.ValueString())
				}
			} else if item.Shutdown.IsNull() {
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
					itemBody, _ = sjson.Set(itemBody, "ikeGroup.value", "2")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ikeGroup.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ikeGroup.value", item.IkeGroup.ValueString())
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
					itemBody, _ = sjson.Set(itemBody, "ipsecCiphersuite.value", "aes256-cbc-sha1")
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
					itemBody, _ = sjson.Set(itemBody, "perfectForwardSecrecy.value", "group-2")
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

			if !item.TrackEnableVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackEnable.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "trackEnable.value", item.TrackEnableVariable.ValueString())
				}
			} else if item.TrackEnable.IsNull() {
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
			body, _ = sjson.SetRaw(body, path+"interface.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"interfacePair", []interface{}{})
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
			body, _ = sjson.SetRaw(body, path+"interfacePair.-1", itemBody)
		}
	}

	if !data.RegionVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"region.optionType", "variable")
			body, _ = sjson.Set(body, path+"region.value", data.RegionVariable.ValueString())
		}
	} else if data.Region.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"region.optionType", "default")
			body, _ = sjson.Set(body, path+"region.value", "auto")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"region.optionType", "global")
			body, _ = sjson.Set(body, path+"region.value", data.Region.ValueString())
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
			if true {
				itemBody, _ = sjson.Set(itemBody, "trackerType.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "trackerType.value", "cisco-sse")
			}
			body, _ = sjson.SetRaw(body, path+"tracker.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SSECisco) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.ContextSharingForVpn = types.BoolNull()

	if t := res.Get(path + "contextSharingForVpn.optionType"); t.Exists() {
		va := res.Get(path + "contextSharingForVpn.value")
		if t.String() == "global" {
			data.ContextSharingForVpn = types.BoolValue(va.Bool())
		}
	}
	data.ContextSharingForSgt = types.BoolNull()

	if t := res.Get(path + "contextSharingForSgt.optionType"); t.Exists() {
		va := res.Get(path + "contextSharingForSgt.value")
		if t.String() == "global" {
			data.ContextSharingForSgt = types.BoolValue(va.Bool())
		}
	}
	if value := res.Get(path + "interface"); value.Exists() && len(value.Array()) > 0 {
		data.Interfaces = make([]SSECiscoInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SSECiscoInterfaces{}
			item.InterfaceName = types.StringNull()

			if t := v.Get("ifName.optionType"); t.Exists() {
				va := v.Get("ifName.value")
				if t.String() == "global" {
					item.InterfaceName = types.StringValue(va.String())
				}
			}
			item.Shutdown = types.BoolNull()
			item.ShutdownVariable = types.StringNull()
			if t := v.Get("shutdown.optionType"); t.Exists() {
				va := v.Get("shutdown.value")
				if t.String() == "variable" {
					item.ShutdownVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Shutdown = types.BoolValue(va.Bool())
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
			item.TrackEnableVariable = types.StringNull()
			if t := v.Get("trackEnable.optionType"); t.Exists() {
				va := v.Get("trackEnable.value")
				if t.String() == "variable" {
					item.TrackEnableVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TrackEnable = types.BoolValue(va.Bool())
				}
			}
			data.Interfaces = append(data.Interfaces, item)
			return true
		})
	}
	if value := res.Get(path + "interfacePair"); value.Exists() && len(value.Array()) > 0 {
		data.InterfacePairs = make([]SSECiscoInterfacePairs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SSECiscoInterfacePairs{}
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
	data.Region = types.StringNull()
	data.RegionVariable = types.StringNull()
	if t := res.Get(path + "region.optionType"); t.Exists() {
		va := res.Get(path + "region.value")
		if t.String() == "variable" {
			data.RegionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Region = types.StringValue(va.String())
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
	if value := res.Get(path + "tracker"); value.Exists() && len(value.Array()) > 0 {
		data.Trackers = make([]SSECiscoTrackers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SSECiscoTrackers{}
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
func (data *SSECisco) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.ContextSharingForVpn = types.BoolNull()

	if t := res.Get(path + "contextSharingForVpn.optionType"); t.Exists() {
		va := res.Get(path + "contextSharingForVpn.value")
		if t.String() == "global" {
			data.ContextSharingForVpn = types.BoolValue(va.Bool())
		}
	}
	data.ContextSharingForSgt = types.BoolNull()

	if t := res.Get(path + "contextSharingForSgt.optionType"); t.Exists() {
		va := res.Get(path + "contextSharingForSgt.value")
		if t.String() == "global" {
			data.ContextSharingForSgt = types.BoolValue(va.Bool())
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
		data.Interfaces[i].Shutdown = types.BoolNull()
		data.Interfaces[i].ShutdownVariable = types.StringNull()
		if t := r.Get("shutdown.optionType"); t.Exists() {
			va := r.Get("shutdown.value")
			if t.String() == "variable" {
				data.Interfaces[i].ShutdownVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].Shutdown = types.BoolValue(va.Bool())
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
		data.Interfaces[i].TrackEnableVariable = types.StringNull()
		if t := r.Get("trackEnable.optionType"); t.Exists() {
			va := r.Get("trackEnable.value")
			if t.String() == "variable" {
				data.Interfaces[i].TrackEnableVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].TrackEnable = types.BoolValue(va.Bool())
			}
		}
	}
	for i := range data.InterfacePairs {
		keys := [...]string{"activeInterface", "activeInterfaceWeight", "backupInterface", "backupInterfaceWeight"}
		keyValues := [...]string{data.InterfacePairs[i].ActiveInterface.ValueString(), strconv.FormatInt(data.InterfacePairs[i].ActiveInterfaceWeight.ValueInt64(), 10), data.InterfacePairs[i].BackupInterface.ValueString(), strconv.FormatInt(data.InterfacePairs[i].BackupInterfaceWeight.ValueInt64(), 10)}
		keyValuesVariables := [...]string{"", "", "", ""}

		var r gjson.Result
		res.Get(path + "interfacePair").ForEach(
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
	data.Region = types.StringNull()
	data.RegionVariable = types.StringNull()
	if t := res.Get(path + "region.optionType"); t.Exists() {
		va := res.Get(path + "region.value")
		if t.String() == "variable" {
			data.RegionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Region = types.StringValue(va.String())
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
