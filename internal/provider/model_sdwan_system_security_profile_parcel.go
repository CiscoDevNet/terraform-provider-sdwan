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
type SystemSecurity struct {
	Id                               types.String              `tfsdk:"id"`
	Version                          types.Int64               `tfsdk:"version"`
	Name                             types.String              `tfsdk:"name"`
	Description                      types.String              `tfsdk:"description"`
	FeatureProfileId                 types.String              `tfsdk:"feature_profile_id"`
	Rekey                            types.Int64               `tfsdk:"rekey"`
	RekeyVariable                    types.String              `tfsdk:"rekey_variable"`
	AntiReplayWindow                 types.String              `tfsdk:"anti_replay_window"`
	AntiReplayWindowVariable         types.String              `tfsdk:"anti_replay_window_variable"`
	ExtendedAntiReplayWindow         types.Int64               `tfsdk:"extended_anti_replay_window"`
	ExtendedAntiReplayWindowVariable types.String              `tfsdk:"extended_anti_replay_window_variable"`
	IpsecPairwiseKeying              types.Bool                `tfsdk:"ipsec_pairwise_keying"`
	IpsecPairwiseKeyingVariable      types.String              `tfsdk:"ipsec_pairwise_keying_variable"`
	IntegrityType                    types.Set                 `tfsdk:"integrity_type"`
	IntegrityTypeVariable            types.String              `tfsdk:"integrity_type_variable"`
	Keychains                        []SystemSecurityKeychains `tfsdk:"keychains"`
	Keys                             []SystemSecurityKeys      `tfsdk:"keys"`
}

type SystemSecurityKeychains struct {
	KeyChainName types.String `tfsdk:"key_chain_name"`
	KeyId        types.Int64  `tfsdk:"key_id"`
}

type SystemSecurityKeys struct {
	Id                             types.Int64  `tfsdk:"id"`
	Name                           types.String `tfsdk:"name"`
	SendId                         types.Int64  `tfsdk:"send_id"`
	SendIdVariable                 types.String `tfsdk:"send_id_variable"`
	ReceiverId                     types.Int64  `tfsdk:"receiver_id"`
	ReceiverIdVariable             types.String `tfsdk:"receiver_id_variable"`
	IncludeTcpOptions              types.Bool   `tfsdk:"include_tcp_options"`
	IncludeTcpOptionsVariable      types.String `tfsdk:"include_tcp_options_variable"`
	AcceptAoMismatch               types.Bool   `tfsdk:"accept_ao_mismatch"`
	AcceptAoMismatchVariable       types.String `tfsdk:"accept_ao_mismatch_variable"`
	CryptoAlgorithm                types.String `tfsdk:"crypto_algorithm"`
	KeyString                      types.String `tfsdk:"key_string"`
	KeyStringVariable              types.String `tfsdk:"key_string_variable"`
	SendLifeTimeLocal              types.Bool   `tfsdk:"send_life_time_local"`
	SendLifeTimeLocalVariable      types.String `tfsdk:"send_life_time_local_variable"`
	SendLifeTimeStartEpoch         types.Int64  `tfsdk:"send_life_time_start_epoch"`
	SendLifeTimeInfinite           types.Bool   `tfsdk:"send_life_time_infinite"`
	SendLifeTimeInfiniteVariable   types.String `tfsdk:"send_life_time_infinite_variable"`
	SendLifeTimeDuration           types.Int64  `tfsdk:"send_life_time_duration"`
	SendLifeTimeDurationVariable   types.String `tfsdk:"send_life_time_duration_variable"`
	SendLifeTimeExact              types.Int64  `tfsdk:"send_life_time_exact"`
	AcceptLifeTimeLocal            types.Bool   `tfsdk:"accept_life_time_local"`
	AcceptLifeTimeLocalVariable    types.String `tfsdk:"accept_life_time_local_variable"`
	AcceptLifeTimeStartEpoch       types.Int64  `tfsdk:"accept_life_time_start_epoch"`
	AcceptLifeTimeInfinite         types.Bool   `tfsdk:"accept_life_time_infinite"`
	AcceptLifeTimeInfiniteVariable types.String `tfsdk:"accept_life_time_infinite_variable"`
	AcceptLifeTimeDuration         types.Int64  `tfsdk:"accept_life_time_duration"`
	AcceptLifeTimeDurationVariable types.String `tfsdk:"accept_life_time_duration_variable"`
	AcceptLifeTimeExact            types.Int64  `tfsdk:"accept_life_time_exact"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SystemSecurity) getModel() string {
	return "system_security"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SystemSecurity) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/security", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SystemSecurity) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.RekeyVariable.IsNull() {
		body, _ = sjson.Set(body, path+"rekey.optionType", "variable")
		body, _ = sjson.Set(body, path+"rekey.value", data.RekeyVariable.ValueString())
	} else if data.Rekey.IsNull() {
		body, _ = sjson.Set(body, path+"rekey.optionType", "default")
		body, _ = sjson.Set(body, path+"rekey.value", 86400)
	} else {
		body, _ = sjson.Set(body, path+"rekey.optionType", "global")
		body, _ = sjson.Set(body, path+"rekey.value", data.Rekey.ValueInt64())
	}

	if !data.AntiReplayWindowVariable.IsNull() {
		body, _ = sjson.Set(body, path+"replayWindow.optionType", "variable")
		body, _ = sjson.Set(body, path+"replayWindow.value", data.AntiReplayWindowVariable.ValueString())
	} else if data.AntiReplayWindow.IsNull() {
		body, _ = sjson.Set(body, path+"replayWindow.optionType", "default")
		body, _ = sjson.Set(body, path+"replayWindow.value", "512")
	} else {
		body, _ = sjson.Set(body, path+"replayWindow.optionType", "global")
		body, _ = sjson.Set(body, path+"replayWindow.value", data.AntiReplayWindow.ValueString())
	}

	if !data.ExtendedAntiReplayWindowVariable.IsNull() {
		body, _ = sjson.Set(body, path+"extendedArWindow.optionType", "variable")
		body, _ = sjson.Set(body, path+"extendedArWindow.value", data.ExtendedAntiReplayWindowVariable.ValueString())
	} else if data.ExtendedAntiReplayWindow.IsNull() {
		body, _ = sjson.Set(body, path+"extendedArWindow.optionType", "default")
		body, _ = sjson.Set(body, path+"extendedArWindow.value", 256)
	} else {
		body, _ = sjson.Set(body, path+"extendedArWindow.optionType", "global")
		body, _ = sjson.Set(body, path+"extendedArWindow.value", data.ExtendedAntiReplayWindow.ValueInt64())
	}

	if !data.IpsecPairwiseKeyingVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pairwiseKeying.optionType", "variable")
		body, _ = sjson.Set(body, path+"pairwiseKeying.value", data.IpsecPairwiseKeyingVariable.ValueString())
	} else if data.IpsecPairwiseKeying.IsNull() {
		body, _ = sjson.Set(body, path+"pairwiseKeying.optionType", "default")
		body, _ = sjson.Set(body, path+"pairwiseKeying.value", false)
	} else {
		body, _ = sjson.Set(body, path+"pairwiseKeying.optionType", "global")
		body, _ = sjson.Set(body, path+"pairwiseKeying.value", data.IpsecPairwiseKeying.ValueBool())
	}

	if !data.IntegrityTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"integrityType.optionType", "variable")
		body, _ = sjson.Set(body, path+"integrityType.value", data.IntegrityTypeVariable.ValueString())
	} else if !data.IntegrityType.IsNull() {
		body, _ = sjson.Set(body, path+"integrityType.optionType", "global")
		var values []string
		data.IntegrityType.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"integrityType.value", values)
	}
	body, _ = sjson.Set(body, path+"keychain", []interface{}{})
	for _, item := range data.Keychains {
		itemBody := ""
		if !item.KeyChainName.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "name.value", item.KeyChainName.ValueString())
		}
		if !item.KeyId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "id.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "id.value", item.KeyId.ValueInt64())
		}
		body, _ = sjson.SetRaw(body, path+"keychain.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"key", []interface{}{})
	for _, item := range data.Keys {
		itemBody := ""
		if !item.Id.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "id.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "id.value", item.Id.ValueInt64())
		}
		if !item.Name.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
		}

		if !item.SendIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sendId.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "sendId.value", item.SendIdVariable.ValueString())
		} else if !item.SendId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sendId.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sendId.value", item.SendId.ValueInt64())
		}

		if !item.ReceiverIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "recvId.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "recvId.value", item.ReceiverIdVariable.ValueString())
		} else if !item.ReceiverId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "recvId.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "recvId.value", item.ReceiverId.ValueInt64())
		}

		if !item.IncludeTcpOptionsVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "includeTcpOptions.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "includeTcpOptions.value", item.IncludeTcpOptionsVariable.ValueString())
		} else if item.IncludeTcpOptions.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "includeTcpOptions.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "includeTcpOptions.value", false)
		} else {
			itemBody, _ = sjson.Set(itemBody, "includeTcpOptions.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "includeTcpOptions.value", item.IncludeTcpOptions.ValueBool())
		}

		if !item.AcceptAoMismatchVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acceptAoMismatch.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "acceptAoMismatch.value", item.AcceptAoMismatchVariable.ValueString())
		} else if item.AcceptAoMismatch.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acceptAoMismatch.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "acceptAoMismatch.value", false)
		} else {
			itemBody, _ = sjson.Set(itemBody, "acceptAoMismatch.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "acceptAoMismatch.value", item.AcceptAoMismatch.ValueBool())
		}
		if !item.CryptoAlgorithm.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "tcp.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "tcp.value", item.CryptoAlgorithm.ValueString())
		}

		if !item.KeyStringVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "keyString.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "keyString.value", item.KeyStringVariable.ValueString())
		} else if !item.KeyString.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "keyString.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "keyString.value", item.KeyString.ValueString())
		}

		if !item.SendLifeTimeLocalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.local.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.local.value", item.SendLifeTimeLocalVariable.ValueString())
		} else if item.SendLifeTimeLocal.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.local.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.local.value", false)
		} else {
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.local.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.local.value", item.SendLifeTimeLocal.ValueBool())
		}
		if !item.SendLifeTimeStartEpoch.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.startEpoch.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.startEpoch.value", item.SendLifeTimeStartEpoch.ValueInt64())
		}

		if !item.SendLifeTimeInfiniteVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.oneOfendChoice.infinite.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.oneOfendChoice.infinite.value", item.SendLifeTimeInfiniteVariable.ValueString())
		} else if !item.SendLifeTimeInfinite.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.oneOfendChoice.infinite.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.oneOfendChoice.infinite.value", item.SendLifeTimeInfinite.ValueBool())
		}

		if !item.SendLifeTimeDurationVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.oneOfendChoice.duration.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.oneOfendChoice.duration.value", item.SendLifeTimeDurationVariable.ValueString())
		} else if !item.SendLifeTimeDuration.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.oneOfendChoice.duration.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.oneOfendChoice.duration.value", item.SendLifeTimeDuration.ValueInt64())
		}
		if !item.SendLifeTimeExact.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.oneOfendChoice.exact.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sendLifetime.oneOfendChoice.exact.value", item.SendLifeTimeExact.ValueInt64())
		}

		if !item.AcceptLifeTimeLocalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.local.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.local.value", item.AcceptLifeTimeLocalVariable.ValueString())
		} else if item.AcceptLifeTimeLocal.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.local.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.local.value", false)
		} else {
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.local.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.local.value", item.AcceptLifeTimeLocal.ValueBool())
		}
		if !item.AcceptLifeTimeStartEpoch.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.startEpoch.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.startEpoch.value", item.AcceptLifeTimeStartEpoch.ValueInt64())
		}

		if !item.AcceptLifeTimeInfiniteVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.oneOfendChoice.infinite.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.oneOfendChoice.infinite.value", item.AcceptLifeTimeInfiniteVariable.ValueString())
		} else if !item.AcceptLifeTimeInfinite.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.oneOfendChoice.infinite.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.oneOfendChoice.infinite.value", item.AcceptLifeTimeInfinite.ValueBool())
		}

		if !item.AcceptLifeTimeDurationVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.oneOfendChoice.duration.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.oneOfendChoice.duration.value", item.AcceptLifeTimeDurationVariable.ValueString())
		} else if !item.AcceptLifeTimeDuration.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.oneOfendChoice.duration.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.oneOfendChoice.duration.value", item.AcceptLifeTimeDuration.ValueInt64())
		}
		if !item.AcceptLifeTimeExact.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.oneOfendChoice.exact.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "acceptLifetime.oneOfendChoice.exact.value", item.AcceptLifeTimeExact.ValueInt64())
		}
		body, _ = sjson.SetRaw(body, path+"key.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemSecurity) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Rekey = types.Int64Null()
	data.RekeyVariable = types.StringNull()
	if t := res.Get(path + "rekey.optionType"); t.Exists() {
		va := res.Get(path + "rekey.value")
		if t.String() == "variable" {
			data.RekeyVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Rekey = types.Int64Value(va.Int())
		}
	}
	data.AntiReplayWindow = types.StringNull()
	data.AntiReplayWindowVariable = types.StringNull()
	if t := res.Get(path + "replayWindow.optionType"); t.Exists() {
		va := res.Get(path + "replayWindow.value")
		if t.String() == "variable" {
			data.AntiReplayWindowVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AntiReplayWindow = types.StringValue(va.String())
		}
	}
	data.ExtendedAntiReplayWindow = types.Int64Null()
	data.ExtendedAntiReplayWindowVariable = types.StringNull()
	if t := res.Get(path + "extendedArWindow.optionType"); t.Exists() {
		va := res.Get(path + "extendedArWindow.value")
		if t.String() == "variable" {
			data.ExtendedAntiReplayWindowVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ExtendedAntiReplayWindow = types.Int64Value(va.Int())
		}
	}
	data.IpsecPairwiseKeying = types.BoolNull()
	data.IpsecPairwiseKeyingVariable = types.StringNull()
	if t := res.Get(path + "pairwiseKeying.optionType"); t.Exists() {
		va := res.Get(path + "pairwiseKeying.value")
		if t.String() == "variable" {
			data.IpsecPairwiseKeyingVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecPairwiseKeying = types.BoolValue(va.Bool())
		}
	}
	data.IntegrityType = types.SetNull(types.StringType)
	data.IntegrityTypeVariable = types.StringNull()
	if t := res.Get(path + "integrityType.optionType"); t.Exists() {
		va := res.Get(path + "integrityType.value")
		if t.String() == "variable" {
			data.IntegrityTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IntegrityType = helpers.GetStringSet(va.Array())
		}
	}
	if value := res.Get(path + "keychain"); value.Exists() {
		data.Keychains = make([]SystemSecurityKeychains, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemSecurityKeychains{}
			item.KeyChainName = types.StringNull()

			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "global" {
					item.KeyChainName = types.StringValue(va.String())
				}
			}
			item.KeyId = types.Int64Null()

			if t := v.Get("id.optionType"); t.Exists() {
				va := v.Get("id.value")
				if t.String() == "global" {
					item.KeyId = types.Int64Value(va.Int())
				}
			}
			data.Keychains = append(data.Keychains, item)
			return true
		})
	}
	if value := res.Get(path + "key"); value.Exists() {
		data.Keys = make([]SystemSecurityKeys, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemSecurityKeys{}
			item.Id = types.Int64Null()

			if t := v.Get("id.optionType"); t.Exists() {
				va := v.Get("id.value")
				if t.String() == "global" {
					item.Id = types.Int64Value(va.Int())
				}
			}
			item.Name = types.StringNull()

			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			item.SendId = types.Int64Null()
			item.SendIdVariable = types.StringNull()
			if t := v.Get("sendId.optionType"); t.Exists() {
				va := v.Get("sendId.value")
				if t.String() == "variable" {
					item.SendIdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SendId = types.Int64Value(va.Int())
				}
			}
			item.ReceiverId = types.Int64Null()
			item.ReceiverIdVariable = types.StringNull()
			if t := v.Get("recvId.optionType"); t.Exists() {
				va := v.Get("recvId.value")
				if t.String() == "variable" {
					item.ReceiverIdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.ReceiverId = types.Int64Value(va.Int())
				}
			}
			item.IncludeTcpOptions = types.BoolNull()
			item.IncludeTcpOptionsVariable = types.StringNull()
			if t := v.Get("includeTcpOptions.optionType"); t.Exists() {
				va := v.Get("includeTcpOptions.value")
				if t.String() == "variable" {
					item.IncludeTcpOptionsVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IncludeTcpOptions = types.BoolValue(va.Bool())
				}
			}
			item.AcceptAoMismatch = types.BoolNull()
			item.AcceptAoMismatchVariable = types.StringNull()
			if t := v.Get("acceptAoMismatch.optionType"); t.Exists() {
				va := v.Get("acceptAoMismatch.value")
				if t.String() == "variable" {
					item.AcceptAoMismatchVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AcceptAoMismatch = types.BoolValue(va.Bool())
				}
			}
			item.CryptoAlgorithm = types.StringNull()

			if t := v.Get("tcp.optionType"); t.Exists() {
				va := v.Get("tcp.value")
				if t.String() == "global" {
					item.CryptoAlgorithm = types.StringValue(va.String())
				}
			}
			item.SendLifeTimeLocal = types.BoolNull()
			item.SendLifeTimeLocalVariable = types.StringNull()
			if t := v.Get("sendLifetime.local.optionType"); t.Exists() {
				va := v.Get("sendLifetime.local.value")
				if t.String() == "variable" {
					item.SendLifeTimeLocalVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SendLifeTimeLocal = types.BoolValue(va.Bool())
				}
			}
			item.SendLifeTimeStartEpoch = types.Int64Null()

			if t := v.Get("sendLifetime.startEpoch.optionType"); t.Exists() {
				va := v.Get("sendLifetime.startEpoch.value")
				if t.String() == "global" {
					item.SendLifeTimeStartEpoch = types.Int64Value(va.Int())
				}
			}
			item.SendLifeTimeInfinite = types.BoolNull()
			item.SendLifeTimeInfiniteVariable = types.StringNull()
			if t := v.Get("sendLifetime.oneOfendChoice.infinite.optionType"); t.Exists() {
				va := v.Get("sendLifetime.oneOfendChoice.infinite.value")
				if t.String() == "variable" {
					item.SendLifeTimeInfiniteVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SendLifeTimeInfinite = types.BoolValue(va.Bool())
				}
			}
			item.SendLifeTimeDuration = types.Int64Null()
			item.SendLifeTimeDurationVariable = types.StringNull()
			if t := v.Get("sendLifetime.oneOfendChoice.duration.optionType"); t.Exists() {
				va := v.Get("sendLifetime.oneOfendChoice.duration.value")
				if t.String() == "variable" {
					item.SendLifeTimeDurationVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SendLifeTimeDuration = types.Int64Value(va.Int())
				}
			}
			item.SendLifeTimeExact = types.Int64Null()

			if t := v.Get("sendLifetime.oneOfendChoice.exact.optionType"); t.Exists() {
				va := v.Get("sendLifetime.oneOfendChoice.exact.value")
				if t.String() == "global" {
					item.SendLifeTimeExact = types.Int64Value(va.Int())
				}
			}
			item.AcceptLifeTimeLocal = types.BoolNull()
			item.AcceptLifeTimeLocalVariable = types.StringNull()
			if t := v.Get("acceptLifetime.local.optionType"); t.Exists() {
				va := v.Get("acceptLifetime.local.value")
				if t.String() == "variable" {
					item.AcceptLifeTimeLocalVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AcceptLifeTimeLocal = types.BoolValue(va.Bool())
				}
			}
			item.AcceptLifeTimeStartEpoch = types.Int64Null()

			if t := v.Get("acceptLifetime.startEpoch.optionType"); t.Exists() {
				va := v.Get("acceptLifetime.startEpoch.value")
				if t.String() == "global" {
					item.AcceptLifeTimeStartEpoch = types.Int64Value(va.Int())
				}
			}
			item.AcceptLifeTimeInfinite = types.BoolNull()
			item.AcceptLifeTimeInfiniteVariable = types.StringNull()
			if t := v.Get("acceptLifetime.oneOfendChoice.infinite.optionType"); t.Exists() {
				va := v.Get("acceptLifetime.oneOfendChoice.infinite.value")
				if t.String() == "variable" {
					item.AcceptLifeTimeInfiniteVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AcceptLifeTimeInfinite = types.BoolValue(va.Bool())
				}
			}
			item.AcceptLifeTimeDuration = types.Int64Null()
			item.AcceptLifeTimeDurationVariable = types.StringNull()
			if t := v.Get("acceptLifetime.oneOfendChoice.duration.optionType"); t.Exists() {
				va := v.Get("acceptLifetime.oneOfendChoice.duration.value")
				if t.String() == "variable" {
					item.AcceptLifeTimeDurationVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AcceptLifeTimeDuration = types.Int64Value(va.Int())
				}
			}
			item.AcceptLifeTimeExact = types.Int64Null()

			if t := v.Get("acceptLifetime.oneOfendChoice.exact.optionType"); t.Exists() {
				va := v.Get("acceptLifetime.oneOfendChoice.exact.value")
				if t.String() == "global" {
					item.AcceptLifeTimeExact = types.Int64Value(va.Int())
				}
			}
			data.Keys = append(data.Keys, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SystemSecurity) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Rekey = types.Int64Null()
	data.RekeyVariable = types.StringNull()
	if t := res.Get(path + "rekey.optionType"); t.Exists() {
		va := res.Get(path + "rekey.value")
		if t.String() == "variable" {
			data.RekeyVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Rekey = types.Int64Value(va.Int())
		}
	}
	data.AntiReplayWindow = types.StringNull()
	data.AntiReplayWindowVariable = types.StringNull()
	if t := res.Get(path + "replayWindow.optionType"); t.Exists() {
		va := res.Get(path + "replayWindow.value")
		if t.String() == "variable" {
			data.AntiReplayWindowVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AntiReplayWindow = types.StringValue(va.String())
		}
	}
	data.ExtendedAntiReplayWindow = types.Int64Null()
	data.ExtendedAntiReplayWindowVariable = types.StringNull()
	if t := res.Get(path + "extendedArWindow.optionType"); t.Exists() {
		va := res.Get(path + "extendedArWindow.value")
		if t.String() == "variable" {
			data.ExtendedAntiReplayWindowVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ExtendedAntiReplayWindow = types.Int64Value(va.Int())
		}
	}
	data.IpsecPairwiseKeying = types.BoolNull()
	data.IpsecPairwiseKeyingVariable = types.StringNull()
	if t := res.Get(path + "pairwiseKeying.optionType"); t.Exists() {
		va := res.Get(path + "pairwiseKeying.value")
		if t.String() == "variable" {
			data.IpsecPairwiseKeyingVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpsecPairwiseKeying = types.BoolValue(va.Bool())
		}
	}
	data.IntegrityType = types.SetNull(types.StringType)
	data.IntegrityTypeVariable = types.StringNull()
	if t := res.Get(path + "integrityType.optionType"); t.Exists() {
		va := res.Get(path + "integrityType.value")
		if t.String() == "variable" {
			data.IntegrityTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IntegrityType = helpers.GetStringSet(va.Array())
		}
	}
	for i := range data.Keychains {
		keys := [...]string{"id"}
		keyValues := [...]string{strconv.FormatInt(data.Keychains[i].KeyId.ValueInt64(), 10)}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "keychain").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Keychains[i].KeyChainName = types.StringNull()

		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "global" {
				data.Keychains[i].KeyChainName = types.StringValue(va.String())
			}
		}
		data.Keychains[i].KeyId = types.Int64Null()

		if t := r.Get("id.optionType"); t.Exists() {
			va := r.Get("id.value")
			if t.String() == "global" {
				data.Keychains[i].KeyId = types.Int64Value(va.Int())
			}
		}
	}
	for i := range data.Keys {
		keys := [...]string{"id"}
		keyValues := [...]string{strconv.FormatInt(data.Keys[i].Id.ValueInt64(), 10)}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "key").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Keys[i].Id = types.Int64Null()

		if t := r.Get("id.optionType"); t.Exists() {
			va := r.Get("id.value")
			if t.String() == "global" {
				data.Keys[i].Id = types.Int64Value(va.Int())
			}
		}
		data.Keys[i].Name = types.StringNull()

		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "global" {
				data.Keys[i].Name = types.StringValue(va.String())
			}
		}
		data.Keys[i].SendId = types.Int64Null()
		data.Keys[i].SendIdVariable = types.StringNull()
		if t := r.Get("sendId.optionType"); t.Exists() {
			va := r.Get("sendId.value")
			if t.String() == "variable" {
				data.Keys[i].SendIdVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Keys[i].SendId = types.Int64Value(va.Int())
			}
		}
		data.Keys[i].ReceiverId = types.Int64Null()
		data.Keys[i].ReceiverIdVariable = types.StringNull()
		if t := r.Get("recvId.optionType"); t.Exists() {
			va := r.Get("recvId.value")
			if t.String() == "variable" {
				data.Keys[i].ReceiverIdVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Keys[i].ReceiverId = types.Int64Value(va.Int())
			}
		}
		data.Keys[i].IncludeTcpOptions = types.BoolNull()
		data.Keys[i].IncludeTcpOptionsVariable = types.StringNull()
		if t := r.Get("includeTcpOptions.optionType"); t.Exists() {
			va := r.Get("includeTcpOptions.value")
			if t.String() == "variable" {
				data.Keys[i].IncludeTcpOptionsVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Keys[i].IncludeTcpOptions = types.BoolValue(va.Bool())
			}
		}
		data.Keys[i].AcceptAoMismatch = types.BoolNull()
		data.Keys[i].AcceptAoMismatchVariable = types.StringNull()
		if t := r.Get("acceptAoMismatch.optionType"); t.Exists() {
			va := r.Get("acceptAoMismatch.value")
			if t.String() == "variable" {
				data.Keys[i].AcceptAoMismatchVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Keys[i].AcceptAoMismatch = types.BoolValue(va.Bool())
			}
		}
		data.Keys[i].CryptoAlgorithm = types.StringNull()

		if t := r.Get("tcp.optionType"); t.Exists() {
			va := r.Get("tcp.value")
			if t.String() == "global" {
				data.Keys[i].CryptoAlgorithm = types.StringValue(va.String())
			}
		}
		data.Keys[i].SendLifeTimeLocal = types.BoolNull()
		data.Keys[i].SendLifeTimeLocalVariable = types.StringNull()
		if t := r.Get("sendLifetime.local.optionType"); t.Exists() {
			va := r.Get("sendLifetime.local.value")
			if t.String() == "variable" {
				data.Keys[i].SendLifeTimeLocalVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Keys[i].SendLifeTimeLocal = types.BoolValue(va.Bool())
			}
		}
		data.Keys[i].SendLifeTimeStartEpoch = types.Int64Null()

		if t := r.Get("sendLifetime.startEpoch.optionType"); t.Exists() {
			va := r.Get("sendLifetime.startEpoch.value")
			if t.String() == "global" {
				data.Keys[i].SendLifeTimeStartEpoch = types.Int64Value(va.Int())
			}
		}
		data.Keys[i].SendLifeTimeInfinite = types.BoolNull()
		data.Keys[i].SendLifeTimeInfiniteVariable = types.StringNull()
		if t := r.Get("sendLifetime.oneOfendChoice.infinite.optionType"); t.Exists() {
			va := r.Get("sendLifetime.oneOfendChoice.infinite.value")
			if t.String() == "variable" {
				data.Keys[i].SendLifeTimeInfiniteVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Keys[i].SendLifeTimeInfinite = types.BoolValue(va.Bool())
			}
		}
		data.Keys[i].SendLifeTimeDuration = types.Int64Null()
		data.Keys[i].SendLifeTimeDurationVariable = types.StringNull()
		if t := r.Get("sendLifetime.oneOfendChoice.duration.optionType"); t.Exists() {
			va := r.Get("sendLifetime.oneOfendChoice.duration.value")
			if t.String() == "variable" {
				data.Keys[i].SendLifeTimeDurationVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Keys[i].SendLifeTimeDuration = types.Int64Value(va.Int())
			}
		}
		data.Keys[i].SendLifeTimeExact = types.Int64Null()

		if t := r.Get("sendLifetime.oneOfendChoice.exact.optionType"); t.Exists() {
			va := r.Get("sendLifetime.oneOfendChoice.exact.value")
			if t.String() == "global" {
				data.Keys[i].SendLifeTimeExact = types.Int64Value(va.Int())
			}
		}
		data.Keys[i].AcceptLifeTimeLocal = types.BoolNull()
		data.Keys[i].AcceptLifeTimeLocalVariable = types.StringNull()
		if t := r.Get("acceptLifetime.local.optionType"); t.Exists() {
			va := r.Get("acceptLifetime.local.value")
			if t.String() == "variable" {
				data.Keys[i].AcceptLifeTimeLocalVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Keys[i].AcceptLifeTimeLocal = types.BoolValue(va.Bool())
			}
		}
		data.Keys[i].AcceptLifeTimeStartEpoch = types.Int64Null()

		if t := r.Get("acceptLifetime.startEpoch.optionType"); t.Exists() {
			va := r.Get("acceptLifetime.startEpoch.value")
			if t.String() == "global" {
				data.Keys[i].AcceptLifeTimeStartEpoch = types.Int64Value(va.Int())
			}
		}
		data.Keys[i].AcceptLifeTimeInfinite = types.BoolNull()
		data.Keys[i].AcceptLifeTimeInfiniteVariable = types.StringNull()
		if t := r.Get("acceptLifetime.oneOfendChoice.infinite.optionType"); t.Exists() {
			va := r.Get("acceptLifetime.oneOfendChoice.infinite.value")
			if t.String() == "variable" {
				data.Keys[i].AcceptLifeTimeInfiniteVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Keys[i].AcceptLifeTimeInfinite = types.BoolValue(va.Bool())
			}
		}
		data.Keys[i].AcceptLifeTimeDuration = types.Int64Null()
		data.Keys[i].AcceptLifeTimeDurationVariable = types.StringNull()
		if t := r.Get("acceptLifetime.oneOfendChoice.duration.optionType"); t.Exists() {
			va := r.Get("acceptLifetime.oneOfendChoice.duration.value")
			if t.String() == "variable" {
				data.Keys[i].AcceptLifeTimeDurationVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Keys[i].AcceptLifeTimeDuration = types.Int64Value(va.Int())
			}
		}
		data.Keys[i].AcceptLifeTimeExact = types.Int64Null()

		if t := r.Get("acceptLifetime.oneOfendChoice.exact.optionType"); t.Exists() {
			va := r.Get("acceptLifetime.oneOfendChoice.exact.value")
			if t.String() == "global" {
				data.Keys[i].AcceptLifeTimeExact = types.Int64Value(va.Int())
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *SystemSecurity) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.Rekey.IsNull() {
		return false
	}
	if !data.RekeyVariable.IsNull() {
		return false
	}
	if !data.AntiReplayWindow.IsNull() {
		return false
	}
	if !data.AntiReplayWindowVariable.IsNull() {
		return false
	}
	if !data.ExtendedAntiReplayWindow.IsNull() {
		return false
	}
	if !data.ExtendedAntiReplayWindowVariable.IsNull() {
		return false
	}
	if !data.IpsecPairwiseKeying.IsNull() {
		return false
	}
	if !data.IpsecPairwiseKeyingVariable.IsNull() {
		return false
	}
	if !data.IntegrityType.IsNull() {
		return false
	}
	if !data.IntegrityTypeVariable.IsNull() {
		return false
	}
	if len(data.Keychains) > 0 {
		return false
	}
	if len(data.Keys) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
