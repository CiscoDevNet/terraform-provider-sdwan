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
type OtherTrustSec struct {
	Id                              types.String                  `tfsdk:"id"`
	Version                         types.Int64                   `tfsdk:"version"`
	Name                            types.String                  `tfsdk:"name"`
	Description                     types.String                  `tfsdk:"description"`
	FeatureProfileId                types.String                  `tfsdk:"feature_profile_id"`
	DeviceId                        types.String                  `tfsdk:"device_id"`
	DeviceIdVariable                types.String                  `tfsdk:"device_id_variable"`
	DevicePassword                  types.String                  `tfsdk:"device_password"`
	DevicePasswordVariable          types.String                  `tfsdk:"device_password_variable"`
	DeviceSgt                       types.Int64                   `tfsdk:"device_sgt"`
	DeviceSgtVariable               types.String                  `tfsdk:"device_sgt_variable"`
	EnableEnforcement               types.Bool                    `tfsdk:"enable_enforcement"`
	EnableEnforcementVariable       types.String                  `tfsdk:"enable_enforcement_variable"`
	EnableSxp                       types.Bool                    `tfsdk:"enable_sxp"`
	ListenerHoldTimeMin             types.Int64                   `tfsdk:"listener_hold_time_min"`
	ListenerHoldTimeMinVariable     types.String                  `tfsdk:"listener_hold_time_min_variable"`
	ListenerHoldTimeMax             types.Int64                   `tfsdk:"listener_hold_time_max"`
	ListenerHoldTimeMaxVariable     types.String                  `tfsdk:"listener_hold_time_max_variable"`
	SpeakerHoldTime                 types.Int64                   `tfsdk:"speaker_hold_time"`
	SpeakerHoldTimeVariable         types.String                  `tfsdk:"speaker_hold_time_variable"`
	SxpDefaultPassword              types.String                  `tfsdk:"sxp_default_password"`
	SxpDefaultPasswordVariable      types.String                  `tfsdk:"sxp_default_password_variable"`
	SxpKeyChain                     types.String                  `tfsdk:"sxp_key_chain"`
	SxpKeyChainVariable             types.String                  `tfsdk:"sxp_key_chain_variable"`
	SxpLogBindingChanges            types.Bool                    `tfsdk:"sxp_log_binding_changes"`
	SxpLogBindingChangesVariable    types.String                  `tfsdk:"sxp_log_binding_changes_variable"`
	SxpReconciliationPeriod         types.Int64                   `tfsdk:"sxp_reconciliation_period"`
	SxpReconciliationPeriodVariable types.String                  `tfsdk:"sxp_reconciliation_period_variable"`
	SxpRetryPeriod                  types.Int64                   `tfsdk:"sxp_retry_period"`
	SxpRetryPeriodVariable          types.String                  `tfsdk:"sxp_retry_period_variable"`
	SxpSourceIp                     types.String                  `tfsdk:"sxp_source_ip"`
	SxpSourceIpVariable             types.String                  `tfsdk:"sxp_source_ip_variable"`
	SxpConnections                  []OtherTrustSecSxpConnections `tfsdk:"sxp_connections"`
}

type OtherTrustSecSxpConnections struct {
	PeerIp              types.String `tfsdk:"peer_ip"`
	PeerIpVariable      types.String `tfsdk:"peer_ip_variable"`
	SourceIp            types.String `tfsdk:"source_ip"`
	SourceIpVariable    types.String `tfsdk:"source_ip_variable"`
	PresharedKey        types.String `tfsdk:"preshared_key"`
	Mode                types.String `tfsdk:"mode"`
	ModeType            types.String `tfsdk:"mode_type"`
	VpnId               types.Int64  `tfsdk:"vpn_id"`
	VpnIdVariable       types.String `tfsdk:"vpn_id_variable"`
	MinHoldTime         types.Int64  `tfsdk:"min_hold_time"`
	MinHoldTimeVariable types.String `tfsdk:"min_hold_time_variable"`
	MaxHoldTime         types.Int64  `tfsdk:"max_hold_time"`
	MaxHoldTimeVariable types.String `tfsdk:"max_hold_time_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data OtherTrustSec) getModel() string {
	return "other_trustsec"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data OtherTrustSec) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/other/%v/trustsec", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data OtherTrustSec) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.DeviceIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"deviceId.optionType", "variable")
			body, _ = sjson.Set(body, path+"deviceId.value", data.DeviceIdVariable.ValueString())
		}
	} else if data.DeviceId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"deviceId.optionType", "default")
			body, _ = sjson.Set(body, path+"deviceId.value", "")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"deviceId.optionType", "global")
			body, _ = sjson.Set(body, path+"deviceId.value", data.DeviceId.ValueString())
		}
	}

	if !data.DevicePasswordVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"devicePassword.optionType", "variable")
			body, _ = sjson.Set(body, path+"devicePassword.value", data.DevicePasswordVariable.ValueString())
		}
	} else if data.DevicePassword.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"devicePassword.optionType", "default")
			body, _ = sjson.Set(body, path+"devicePassword.value", "")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"devicePassword.optionType", "global")
			body, _ = sjson.Set(body, path+"devicePassword.value", data.DevicePassword.ValueString())
		}
	}

	if !data.DeviceSgtVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"deviceSgt.optionType", "variable")
			body, _ = sjson.Set(body, path+"deviceSgt.value", data.DeviceSgtVariable.ValueString())
		}
	} else if !data.DeviceSgt.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"deviceSgt.optionType", "global")
			body, _ = sjson.Set(body, path+"deviceSgt.value", data.DeviceSgt.ValueInt64())
		}
	}

	if !data.EnableEnforcementVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"enableEnforcement.optionType", "variable")
			body, _ = sjson.Set(body, path+"enableEnforcement.value", data.EnableEnforcementVariable.ValueString())
		}
	} else if data.EnableEnforcement.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"enableEnforcement.optionType", "default")
			body, _ = sjson.Set(body, path+"enableEnforcement.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"enableEnforcement.optionType", "global")
			body, _ = sjson.Set(body, path+"enableEnforcement.value", data.EnableEnforcement.ValueBool())
		}
	}
	if data.EnableSxp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"enableSxp.optionType", "default")
			body, _ = sjson.Set(body, path+"enableSxp.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"enableSxp.optionType", "global")
			body, _ = sjson.Set(body, path+"enableSxp.value", data.EnableSxp.ValueBool())
		}
	}

	if !data.ListenerHoldTimeMinVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"listenerHoldTimeMin.optionType", "variable")
			body, _ = sjson.Set(body, path+"listenerHoldTimeMin.value", data.ListenerHoldTimeMinVariable.ValueString())
		}
	} else if data.ListenerHoldTimeMin.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"listenerHoldTimeMin.optionType", "default")
			body, _ = sjson.Set(body, path+"listenerHoldTimeMin.value", 90)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"listenerHoldTimeMin.optionType", "global")
			body, _ = sjson.Set(body, path+"listenerHoldTimeMin.value", data.ListenerHoldTimeMin.ValueInt64())
		}
	}

	if !data.ListenerHoldTimeMaxVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"listenerHoldTimeMax.optionType", "variable")
			body, _ = sjson.Set(body, path+"listenerHoldTimeMax.value", data.ListenerHoldTimeMaxVariable.ValueString())
		}
	} else if data.ListenerHoldTimeMax.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"listenerHoldTimeMax.optionType", "default")
			body, _ = sjson.Set(body, path+"listenerHoldTimeMax.value", 180)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"listenerHoldTimeMax.optionType", "global")
			body, _ = sjson.Set(body, path+"listenerHoldTimeMax.value", data.ListenerHoldTimeMax.ValueInt64())
		}
	}

	if !data.SpeakerHoldTimeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"speakerHoldTime.optionType", "variable")
			body, _ = sjson.Set(body, path+"speakerHoldTime.value", data.SpeakerHoldTimeVariable.ValueString())
		}
	} else if data.SpeakerHoldTime.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"speakerHoldTime.optionType", "default")
			body, _ = sjson.Set(body, path+"speakerHoldTime.value", 120)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"speakerHoldTime.optionType", "global")
			body, _ = sjson.Set(body, path+"speakerHoldTime.value", data.SpeakerHoldTime.ValueInt64())
		}
	}

	if !data.SxpDefaultPasswordVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sxpDefaultPassword.optionType", "variable")
			body, _ = sjson.Set(body, path+"sxpDefaultPassword.value", data.SxpDefaultPasswordVariable.ValueString())
		}
	} else if data.SxpDefaultPassword.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sxpDefaultPassword.optionType", "default")
			body, _ = sjson.Set(body, path+"sxpDefaultPassword.value", "")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"sxpDefaultPassword.optionType", "global")
			body, _ = sjson.Set(body, path+"sxpDefaultPassword.value", data.SxpDefaultPassword.ValueString())
		}
	}

	if !data.SxpKeyChainVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sxpKeyChain.optionType", "variable")
			body, _ = sjson.Set(body, path+"sxpKeyChain.value", data.SxpKeyChainVariable.ValueString())
		}
	} else if data.SxpKeyChain.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sxpKeyChain.optionType", "default")
			body, _ = sjson.Set(body, path+"sxpKeyChain.value", "")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"sxpKeyChain.optionType", "global")
			body, _ = sjson.Set(body, path+"sxpKeyChain.value", data.SxpKeyChain.ValueString())
		}
	}

	if !data.SxpLogBindingChangesVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sxpLogBindingChanges.optionType", "variable")
			body, _ = sjson.Set(body, path+"sxpLogBindingChanges.value", data.SxpLogBindingChangesVariable.ValueString())
		}
	} else if data.SxpLogBindingChanges.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sxpLogBindingChanges.optionType", "default")
			body, _ = sjson.Set(body, path+"sxpLogBindingChanges.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"sxpLogBindingChanges.optionType", "global")
			body, _ = sjson.Set(body, path+"sxpLogBindingChanges.value", data.SxpLogBindingChanges.ValueBool())
		}
	}

	if !data.SxpReconciliationPeriodVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sxpReconciliationPeriod.optionType", "variable")
			body, _ = sjson.Set(body, path+"sxpReconciliationPeriod.value", data.SxpReconciliationPeriodVariable.ValueString())
		}
	} else if data.SxpReconciliationPeriod.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sxpReconciliationPeriod.optionType", "default")
			body, _ = sjson.Set(body, path+"sxpReconciliationPeriod.value", 120)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"sxpReconciliationPeriod.optionType", "global")
			body, _ = sjson.Set(body, path+"sxpReconciliationPeriod.value", data.SxpReconciliationPeriod.ValueInt64())
		}
	}

	if !data.SxpRetryPeriodVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sxpRetryPeriod.optionType", "variable")
			body, _ = sjson.Set(body, path+"sxpRetryPeriod.value", data.SxpRetryPeriodVariable.ValueString())
		}
	} else if data.SxpRetryPeriod.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sxpRetryPeriod.optionType", "default")
			body, _ = sjson.Set(body, path+"sxpRetryPeriod.value", 120)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"sxpRetryPeriod.optionType", "global")
			body, _ = sjson.Set(body, path+"sxpRetryPeriod.value", data.SxpRetryPeriod.ValueInt64())
		}
	}

	if !data.SxpSourceIpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sxpSourceIp.optionType", "variable")
			body, _ = sjson.Set(body, path+"sxpSourceIp.value", data.SxpSourceIpVariable.ValueString())
		}
	} else if data.SxpSourceIp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sxpSourceIp.optionType", "default")
			body, _ = sjson.Set(body, path+"sxpSourceIp.value", "")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"sxpSourceIp.optionType", "global")
			body, _ = sjson.Set(body, path+"sxpSourceIp.value", data.SxpSourceIp.ValueString())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"sxpNodeId.optionType", "default")
		body, _ = sjson.Set(body, path+"sxpNodeId.value", "System IP")
	}
	if true {
		body, _ = sjson.Set(body, path+"sxpConnectionList", []interface{}{})
		for _, item := range data.SxpConnections {
			itemBody := ""

			if !item.PeerIpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionPeerIp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "connectionPeerIp.value", item.PeerIpVariable.ValueString())
				}
			} else if !item.PeerIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionPeerIp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "connectionPeerIp.value", item.PeerIp.ValueString())
				}
			}

			if !item.SourceIpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionSourceIp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "connectionSourceIp.value", item.SourceIpVariable.ValueString())
				}
			} else if !item.SourceIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionSourceIp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "connectionSourceIp.value", item.SourceIp.ValueString())
				}
			}
			if item.PresharedKey.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionPresharedKey.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "connectionPresharedKey.value", "none")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionPresharedKey.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "connectionPresharedKey.value", item.PresharedKey.ValueString())
				}
			}
			if item.Mode.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionMode.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "connectionMode.value", "local")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionMode.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "connectionMode.value", item.Mode.ValueString())
				}
			}
			if item.ModeType.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionModeType.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "connectionModeType.value", "speaker")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionModeType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "connectionModeType.value", item.ModeType.ValueString())
				}
			}

			if !item.VpnIdVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionVpnId.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "connectionVpnId.value", item.VpnIdVariable.ValueString())
				}
			} else if item.VpnId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionVpnId.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "connectionVpnId.value", 0)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionVpnId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "connectionVpnId.value", item.VpnId.ValueInt64())
				}
			}

			if !item.MinHoldTimeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionMinHoldTime.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "connectionMinHoldTime.value", item.MinHoldTimeVariable.ValueString())
				}
			} else if !item.MinHoldTime.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionMinHoldTime.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "connectionMinHoldTime.value", item.MinHoldTime.ValueInt64())
				}
			}

			if !item.MaxHoldTimeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionMaxHoldTime.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "connectionMaxHoldTime.value", item.MaxHoldTimeVariable.ValueString())
				}
			} else if !item.MaxHoldTime.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "connectionMaxHoldTime.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "connectionMaxHoldTime.value", item.MaxHoldTime.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"sxpConnectionList.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *OtherTrustSec) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.DeviceId = types.StringNull()
	data.DeviceIdVariable = types.StringNull()
	if t := res.Get(path + "deviceId.optionType"); t.Exists() {
		va := res.Get(path + "deviceId.value")
		if t.String() == "variable" {
			data.DeviceIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DeviceId = types.StringValue(va.String())
		}
	}
	data.DeviceSgt = types.Int64Null()
	data.DeviceSgtVariable = types.StringNull()
	if t := res.Get(path + "deviceSgt.optionType"); t.Exists() {
		va := res.Get(path + "deviceSgt.value")
		if t.String() == "variable" {
			data.DeviceSgtVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DeviceSgt = types.Int64Value(va.Int())
		}
	}
	data.EnableEnforcement = types.BoolNull()
	data.EnableEnforcementVariable = types.StringNull()
	if t := res.Get(path + "enableEnforcement.optionType"); t.Exists() {
		va := res.Get(path + "enableEnforcement.value")
		if t.String() == "variable" {
			data.EnableEnforcementVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EnableEnforcement = types.BoolValue(va.Bool())
		}
	}
	data.EnableSxp = types.BoolNull()

	if t := res.Get(path + "enableSxp.optionType"); t.Exists() {
		va := res.Get(path + "enableSxp.value")
		if t.String() == "global" {
			data.EnableSxp = types.BoolValue(va.Bool())
		}
	}
	data.ListenerHoldTimeMin = types.Int64Null()
	data.ListenerHoldTimeMinVariable = types.StringNull()
	if t := res.Get(path + "listenerHoldTimeMin.optionType"); t.Exists() {
		va := res.Get(path + "listenerHoldTimeMin.value")
		if t.String() == "variable" {
			data.ListenerHoldTimeMinVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ListenerHoldTimeMin = types.Int64Value(va.Int())
		}
	}
	data.ListenerHoldTimeMax = types.Int64Null()
	data.ListenerHoldTimeMaxVariable = types.StringNull()
	if t := res.Get(path + "listenerHoldTimeMax.optionType"); t.Exists() {
		va := res.Get(path + "listenerHoldTimeMax.value")
		if t.String() == "variable" {
			data.ListenerHoldTimeMaxVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ListenerHoldTimeMax = types.Int64Value(va.Int())
		}
	}
	data.SpeakerHoldTime = types.Int64Null()
	data.SpeakerHoldTimeVariable = types.StringNull()
	if t := res.Get(path + "speakerHoldTime.optionType"); t.Exists() {
		va := res.Get(path + "speakerHoldTime.value")
		if t.String() == "variable" {
			data.SpeakerHoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SpeakerHoldTime = types.Int64Value(va.Int())
		}
	}
	data.SxpKeyChain = types.StringNull()
	data.SxpKeyChainVariable = types.StringNull()
	if t := res.Get(path + "sxpKeyChain.optionType"); t.Exists() {
		va := res.Get(path + "sxpKeyChain.value")
		if t.String() == "variable" {
			data.SxpKeyChainVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SxpKeyChain = types.StringValue(va.String())
		}
	}
	data.SxpLogBindingChanges = types.BoolNull()
	data.SxpLogBindingChangesVariable = types.StringNull()
	if t := res.Get(path + "sxpLogBindingChanges.optionType"); t.Exists() {
		va := res.Get(path + "sxpLogBindingChanges.value")
		if t.String() == "variable" {
			data.SxpLogBindingChangesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SxpLogBindingChanges = types.BoolValue(va.Bool())
		}
	}
	data.SxpReconciliationPeriod = types.Int64Null()
	data.SxpReconciliationPeriodVariable = types.StringNull()
	if t := res.Get(path + "sxpReconciliationPeriod.optionType"); t.Exists() {
		va := res.Get(path + "sxpReconciliationPeriod.value")
		if t.String() == "variable" {
			data.SxpReconciliationPeriodVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SxpReconciliationPeriod = types.Int64Value(va.Int())
		}
	}
	data.SxpRetryPeriod = types.Int64Null()
	data.SxpRetryPeriodVariable = types.StringNull()
	if t := res.Get(path + "sxpRetryPeriod.optionType"); t.Exists() {
		va := res.Get(path + "sxpRetryPeriod.value")
		if t.String() == "variable" {
			data.SxpRetryPeriodVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SxpRetryPeriod = types.Int64Value(va.Int())
		}
	}
	data.SxpSourceIp = types.StringNull()
	data.SxpSourceIpVariable = types.StringNull()
	if t := res.Get(path + "sxpSourceIp.optionType"); t.Exists() {
		va := res.Get(path + "sxpSourceIp.value")
		if t.String() == "variable" {
			data.SxpSourceIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SxpSourceIp = types.StringValue(va.String())
		}
	}
	if value := res.Get(path + "sxpConnectionList"); value.Exists() && len(value.Array()) > 0 {
		data.SxpConnections = make([]OtherTrustSecSxpConnections, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := OtherTrustSecSxpConnections{}
			item.PeerIp = types.StringNull()
			item.PeerIpVariable = types.StringNull()
			if t := v.Get("connectionPeerIp.optionType"); t.Exists() {
				va := v.Get("connectionPeerIp.value")
				if t.String() == "variable" {
					item.PeerIpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.PeerIp = types.StringValue(va.String())
				}
			}
			item.SourceIp = types.StringNull()
			item.SourceIpVariable = types.StringNull()
			if t := v.Get("connectionSourceIp.optionType"); t.Exists() {
				va := v.Get("connectionSourceIp.value")
				if t.String() == "variable" {
					item.SourceIpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourceIp = types.StringValue(va.String())
				}
			}
			item.PresharedKey = types.StringNull()

			if t := v.Get("connectionPresharedKey.optionType"); t.Exists() {
				va := v.Get("connectionPresharedKey.value")
				if t.String() == "global" {
					item.PresharedKey = types.StringValue(va.String())
				}
			}
			item.Mode = types.StringNull()

			if t := v.Get("connectionMode.optionType"); t.Exists() {
				va := v.Get("connectionMode.value")
				if t.String() == "global" {
					item.Mode = types.StringValue(va.String())
				}
			}
			item.ModeType = types.StringNull()

			if t := v.Get("connectionModeType.optionType"); t.Exists() {
				va := v.Get("connectionModeType.value")
				if t.String() == "global" {
					item.ModeType = types.StringValue(va.String())
				}
			}
			item.VpnId = types.Int64Null()
			item.VpnIdVariable = types.StringNull()
			if t := v.Get("connectionVpnId.optionType"); t.Exists() {
				va := v.Get("connectionVpnId.value")
				if t.String() == "variable" {
					item.VpnIdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.VpnId = types.Int64Value(va.Int())
				}
			}
			item.MinHoldTime = types.Int64Null()
			item.MinHoldTimeVariable = types.StringNull()
			if t := v.Get("connectionMinHoldTime.optionType"); t.Exists() {
				va := v.Get("connectionMinHoldTime.value")
				if t.String() == "variable" {
					item.MinHoldTimeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.MinHoldTime = types.Int64Value(va.Int())
				}
			}
			item.MaxHoldTime = types.Int64Null()
			item.MaxHoldTimeVariable = types.StringNull()
			if t := v.Get("connectionMaxHoldTime.optionType"); t.Exists() {
				va := v.Get("connectionMaxHoldTime.value")
				if t.String() == "variable" {
					item.MaxHoldTimeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.MaxHoldTime = types.Int64Value(va.Int())
				}
			}
			data.SxpConnections = append(data.SxpConnections, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *OtherTrustSec) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.DeviceId = types.StringNull()
	data.DeviceIdVariable = types.StringNull()
	if t := res.Get(path + "deviceId.optionType"); t.Exists() {
		va := res.Get(path + "deviceId.value")
		if t.String() == "variable" {
			data.DeviceIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DeviceId = types.StringValue(va.String())
		}
	}
	data.DeviceSgt = types.Int64Null()
	data.DeviceSgtVariable = types.StringNull()
	if t := res.Get(path + "deviceSgt.optionType"); t.Exists() {
		va := res.Get(path + "deviceSgt.value")
		if t.String() == "variable" {
			data.DeviceSgtVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DeviceSgt = types.Int64Value(va.Int())
		}
	}
	data.EnableEnforcement = types.BoolNull()
	data.EnableEnforcementVariable = types.StringNull()
	if t := res.Get(path + "enableEnforcement.optionType"); t.Exists() {
		va := res.Get(path + "enableEnforcement.value")
		if t.String() == "variable" {
			data.EnableEnforcementVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EnableEnforcement = types.BoolValue(va.Bool())
		}
	}
	data.EnableSxp = types.BoolNull()

	if t := res.Get(path + "enableSxp.optionType"); t.Exists() {
		va := res.Get(path + "enableSxp.value")
		if t.String() == "global" {
			data.EnableSxp = types.BoolValue(va.Bool())
		}
	}
	data.ListenerHoldTimeMin = types.Int64Null()
	data.ListenerHoldTimeMinVariable = types.StringNull()
	if t := res.Get(path + "listenerHoldTimeMin.optionType"); t.Exists() {
		va := res.Get(path + "listenerHoldTimeMin.value")
		if t.String() == "variable" {
			data.ListenerHoldTimeMinVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ListenerHoldTimeMin = types.Int64Value(va.Int())
		}
	}
	data.ListenerHoldTimeMax = types.Int64Null()
	data.ListenerHoldTimeMaxVariable = types.StringNull()
	if t := res.Get(path + "listenerHoldTimeMax.optionType"); t.Exists() {
		va := res.Get(path + "listenerHoldTimeMax.value")
		if t.String() == "variable" {
			data.ListenerHoldTimeMaxVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ListenerHoldTimeMax = types.Int64Value(va.Int())
		}
	}
	data.SpeakerHoldTime = types.Int64Null()
	data.SpeakerHoldTimeVariable = types.StringNull()
	if t := res.Get(path + "speakerHoldTime.optionType"); t.Exists() {
		va := res.Get(path + "speakerHoldTime.value")
		if t.String() == "variable" {
			data.SpeakerHoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SpeakerHoldTime = types.Int64Value(va.Int())
		}
	}
	data.SxpKeyChain = types.StringNull()
	data.SxpKeyChainVariable = types.StringNull()
	if t := res.Get(path + "sxpKeyChain.optionType"); t.Exists() {
		va := res.Get(path + "sxpKeyChain.value")
		if t.String() == "variable" {
			data.SxpKeyChainVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SxpKeyChain = types.StringValue(va.String())
		}
	}
	data.SxpLogBindingChanges = types.BoolNull()
	data.SxpLogBindingChangesVariable = types.StringNull()
	if t := res.Get(path + "sxpLogBindingChanges.optionType"); t.Exists() {
		va := res.Get(path + "sxpLogBindingChanges.value")
		if t.String() == "variable" {
			data.SxpLogBindingChangesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SxpLogBindingChanges = types.BoolValue(va.Bool())
		}
	}
	data.SxpReconciliationPeriod = types.Int64Null()
	data.SxpReconciliationPeriodVariable = types.StringNull()
	if t := res.Get(path + "sxpReconciliationPeriod.optionType"); t.Exists() {
		va := res.Get(path + "sxpReconciliationPeriod.value")
		if t.String() == "variable" {
			data.SxpReconciliationPeriodVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SxpReconciliationPeriod = types.Int64Value(va.Int())
		}
	}
	data.SxpRetryPeriod = types.Int64Null()
	data.SxpRetryPeriodVariable = types.StringNull()
	if t := res.Get(path + "sxpRetryPeriod.optionType"); t.Exists() {
		va := res.Get(path + "sxpRetryPeriod.value")
		if t.String() == "variable" {
			data.SxpRetryPeriodVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SxpRetryPeriod = types.Int64Value(va.Int())
		}
	}
	data.SxpSourceIp = types.StringNull()
	data.SxpSourceIpVariable = types.StringNull()
	if t := res.Get(path + "sxpSourceIp.optionType"); t.Exists() {
		va := res.Get(path + "sxpSourceIp.value")
		if t.String() == "variable" {
			data.SxpSourceIpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SxpSourceIp = types.StringValue(va.String())
		}
	}
	for i := range data.SxpConnections {
		keys := [...]string{"connectionPeerIp"}
		keyValues := [...]string{data.SxpConnections[i].PeerIp.ValueString()}
		keyValuesVariables := [...]string{data.SxpConnections[i].PeerIpVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "sxpConnectionList").ForEach(
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
		data.SxpConnections[i].PeerIp = types.StringNull()
		data.SxpConnections[i].PeerIpVariable = types.StringNull()
		if t := r.Get("connectionPeerIp.optionType"); t.Exists() {
			va := r.Get("connectionPeerIp.value")
			if t.String() == "variable" {
				data.SxpConnections[i].PeerIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SxpConnections[i].PeerIp = types.StringValue(va.String())
			}
		}
		data.SxpConnections[i].SourceIp = types.StringNull()
		data.SxpConnections[i].SourceIpVariable = types.StringNull()
		if t := r.Get("connectionSourceIp.optionType"); t.Exists() {
			va := r.Get("connectionSourceIp.value")
			if t.String() == "variable" {
				data.SxpConnections[i].SourceIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SxpConnections[i].SourceIp = types.StringValue(va.String())
			}
		}
		data.SxpConnections[i].PresharedKey = types.StringNull()

		if t := r.Get("connectionPresharedKey.optionType"); t.Exists() {
			va := r.Get("connectionPresharedKey.value")
			if t.String() == "global" {
				data.SxpConnections[i].PresharedKey = types.StringValue(va.String())
			}
		}
		data.SxpConnections[i].Mode = types.StringNull()

		if t := r.Get("connectionMode.optionType"); t.Exists() {
			va := r.Get("connectionMode.value")
			if t.String() == "global" {
				data.SxpConnections[i].Mode = types.StringValue(va.String())
			}
		}
		data.SxpConnections[i].ModeType = types.StringNull()

		if t := r.Get("connectionModeType.optionType"); t.Exists() {
			va := r.Get("connectionModeType.value")
			if t.String() == "global" {
				data.SxpConnections[i].ModeType = types.StringValue(va.String())
			}
		}
		data.SxpConnections[i].VpnId = types.Int64Null()
		data.SxpConnections[i].VpnIdVariable = types.StringNull()
		if t := r.Get("connectionVpnId.optionType"); t.Exists() {
			va := r.Get("connectionVpnId.value")
			if t.String() == "variable" {
				data.SxpConnections[i].VpnIdVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SxpConnections[i].VpnId = types.Int64Value(va.Int())
			}
		}
		data.SxpConnections[i].MinHoldTime = types.Int64Null()
		data.SxpConnections[i].MinHoldTimeVariable = types.StringNull()
		if t := r.Get("connectionMinHoldTime.optionType"); t.Exists() {
			va := r.Get("connectionMinHoldTime.value")
			if t.String() == "variable" {
				data.SxpConnections[i].MinHoldTimeVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SxpConnections[i].MinHoldTime = types.Int64Value(va.Int())
			}
		}
		data.SxpConnections[i].MaxHoldTime = types.Int64Null()
		data.SxpConnections[i].MaxHoldTimeVariable = types.StringNull()
		if t := r.Get("connectionMaxHoldTime.optionType"); t.Exists() {
			va := r.Get("connectionMaxHoldTime.value")
			if t.String() == "variable" {
				data.SxpConnections[i].MaxHoldTimeVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.SxpConnections[i].MaxHoldTime = types.Int64Value(va.Int())
			}
		}
	}
}

// End of section. //template:end updateFromBody
