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
type SystemIPv4DeviceAccess struct {
	Id               types.String                      `tfsdk:"id"`
	Version          types.Int64                       `tfsdk:"version"`
	Name             types.String                      `tfsdk:"name"`
	Description      types.String                      `tfsdk:"description"`
	FeatureProfileId types.String                      `tfsdk:"feature_profile_id"`
	DefaultAction    types.String                      `tfsdk:"default_action"`
	Sequences        []SystemIPv4DeviceAccessSequences `tfsdk:"sequences"`
}

type SystemIPv4DeviceAccessSequences struct {
	Id                              types.Int64  `tfsdk:"id"`
	Name                            types.String `tfsdk:"name"`
	BaseAction                      types.String `tfsdk:"base_action"`
	DeviceAccessPort                types.Int64  `tfsdk:"device_access_port"`
	DestinationDataPrefixListId     types.String `tfsdk:"destination_data_prefix_list_id"`
	DestinationIpPrefixList         types.Set    `tfsdk:"destination_ip_prefix_list"`
	DestinationIpPrefixListVariable types.String `tfsdk:"destination_ip_prefix_list_variable"`
	SourcePorts                     types.Set    `tfsdk:"source_ports"`
	SourceDataPrefixListId          types.String `tfsdk:"source_data_prefix_list_id"`
	SourceIpPrefixList              types.Set    `tfsdk:"source_ip_prefix_list"`
	SourceIpPrefixListVariable      types.String `tfsdk:"source_ip_prefix_list_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SystemIPv4DeviceAccess) getModel() string {
	return "system_ipv4_device_access"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SystemIPv4DeviceAccess) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/ipv4-device-access-policy", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SystemIPv4DeviceAccess) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if data.DefaultAction.IsNull() || data.DefaultAction.ValueString() == "drop" {
		if true {
			body, _ = sjson.Set(body, path+"defaultAction.optionType", "default")
			body, _ = sjson.Set(body, path+"defaultAction.value", "drop")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"defaultAction.optionType", "global")
			body, _ = sjson.Set(body, path+"defaultAction.value", data.DefaultAction.ValueString())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"sequences", []interface{}{})
		for _, item := range data.Sequences {
			itemBody := ""
			if !item.Id.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceId.value", item.Id.ValueInt64())
				}
			}
			if !item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceName.value", item.Name.ValueString())
				}
			}
			if !item.BaseAction.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "baseAction.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "baseAction.value", item.BaseAction.ValueString())
				}
			}
			if !item.DeviceAccessPort.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "matchEntries.destinationPort.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "matchEntries.destinationPort.value", item.DeviceAccessPort.ValueInt64())
				}
			}
			if !item.DestinationDataPrefixListId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "matchEntries.destinationDataPrefix.destinationDataPrefixList.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "matchEntries.destinationDataPrefix.destinationDataPrefixList.refId.value", item.DestinationDataPrefixListId.ValueString())
				}
			}

			if !item.DestinationIpPrefixListVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "matchEntries.destinationDataPrefix.destinationIpPrefixList.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "matchEntries.destinationDataPrefix.destinationIpPrefixList.value", item.DestinationIpPrefixListVariable.ValueString())
				}
			} else if !item.DestinationIpPrefixList.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "matchEntries.destinationDataPrefix.destinationIpPrefixList.optionType", "global")
					var values []string
					item.DestinationIpPrefixList.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "matchEntries.destinationDataPrefix.destinationIpPrefixList.value", values)
				}
			}
			if !item.SourcePorts.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "matchEntries.sourcePorts.optionType", "global")
					var values []int64
					item.SourcePorts.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "matchEntries.sourcePorts.value", values)
				}
			}
			if !item.SourceDataPrefixListId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "matchEntries.sourceDataPrefix.sourceDataPrefixList.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "matchEntries.sourceDataPrefix.sourceDataPrefixList.refId.value", item.SourceDataPrefixListId.ValueString())
				}
			}

			if !item.SourceIpPrefixListVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "matchEntries.sourceDataPrefix.sourceIpPrefixList.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "matchEntries.sourceDataPrefix.sourceIpPrefixList.value", item.SourceIpPrefixListVariable.ValueString())
				}
			} else if !item.SourceIpPrefixList.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "matchEntries.sourceDataPrefix.sourceIpPrefixList.optionType", "global")
					var values []string
					item.SourceIpPrefixList.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "matchEntries.sourceDataPrefix.sourceIpPrefixList.value", values)
				}
			}
			body, _ = sjson.SetRaw(body, path+"sequences.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemIPv4DeviceAccess) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	tempDefaultAction := data.DefaultAction
	data.DefaultAction = types.StringNull()

	if t := res.Get(path + "defaultAction.optionType"); t.Exists() {
		va := res.Get(path + "defaultAction.value")
		if t.String() == "global" || (t.String() == "default" && tempDefaultAction.ValueString() == "drop") {
			data.DefaultAction = types.StringValue(va.String())
		}
	}
	if value := res.Get(path + "sequences"); value.Exists() {
		data.Sequences = make([]SystemIPv4DeviceAccessSequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemIPv4DeviceAccessSequences{}
			item.Id = types.Int64Null()

			if t := v.Get("sequenceId.optionType"); t.Exists() {
				va := v.Get("sequenceId.value")
				if t.String() == "global" {
					item.Id = types.Int64Value(va.Int())
				}
			}
			item.Name = types.StringNull()

			if t := v.Get("sequenceName.optionType"); t.Exists() {
				va := v.Get("sequenceName.value")
				if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			item.BaseAction = types.StringNull()

			if t := v.Get("baseAction.optionType"); t.Exists() {
				va := v.Get("baseAction.value")
				if t.String() == "global" {
					item.BaseAction = types.StringValue(va.String())
				}
			}
			item.DeviceAccessPort = types.Int64Null()

			if t := v.Get("matchEntries.destinationPort.optionType"); t.Exists() {
				va := v.Get("matchEntries.destinationPort.value")
				if t.String() == "global" {
					item.DeviceAccessPort = types.Int64Value(va.Int())
				}
			}
			item.DestinationDataPrefixListId = types.StringNull()

			if t := v.Get("matchEntries.destinationDataPrefix.destinationDataPrefixList.refId.optionType"); t.Exists() {
				va := v.Get("matchEntries.destinationDataPrefix.destinationDataPrefixList.refId.value")
				if t.String() == "global" {
					item.DestinationDataPrefixListId = types.StringValue(va.String())
				}
			}
			item.DestinationIpPrefixList = types.SetNull(types.StringType)
			item.DestinationIpPrefixListVariable = types.StringNull()
			if t := v.Get("matchEntries.destinationDataPrefix.destinationIpPrefixList.optionType"); t.Exists() {
				va := v.Get("matchEntries.destinationDataPrefix.destinationIpPrefixList.value")
				if t.String() == "variable" {
					item.DestinationIpPrefixListVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.DestinationIpPrefixList = helpers.GetStringSet(va.Array())
				}
			}
			item.SourcePorts = types.SetNull(types.Int64Type)

			if t := v.Get("matchEntries.sourcePorts.optionType"); t.Exists() {
				va := v.Get("matchEntries.sourcePorts.value")
				if t.String() == "global" {
					item.SourcePorts = helpers.GetInt64Set(va.Array())
				}
			}
			item.SourceDataPrefixListId = types.StringNull()

			if t := v.Get("matchEntries.sourceDataPrefix.sourceDataPrefixList.refId.optionType"); t.Exists() {
				va := v.Get("matchEntries.sourceDataPrefix.sourceDataPrefixList.refId.value")
				if t.String() == "global" {
					item.SourceDataPrefixListId = types.StringValue(va.String())
				}
			}
			item.SourceIpPrefixList = types.SetNull(types.StringType)
			item.SourceIpPrefixListVariable = types.StringNull()
			if t := v.Get("matchEntries.sourceDataPrefix.sourceIpPrefixList.optionType"); t.Exists() {
				va := v.Get("matchEntries.sourceDataPrefix.sourceIpPrefixList.value")
				if t.String() == "variable" {
					item.SourceIpPrefixListVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourceIpPrefixList = helpers.GetStringSet(va.Array())
				}
			}
			data.Sequences = append(data.Sequences, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SystemIPv4DeviceAccess) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	tempDefaultAction := data.DefaultAction
	data.DefaultAction = types.StringNull()

	if t := res.Get(path + "defaultAction.optionType"); t.Exists() {
		va := res.Get(path + "defaultAction.value")
		if t.String() == "global" || (t.String() == "default" && tempDefaultAction.ValueString() == "drop") {
			data.DefaultAction = types.StringValue(va.String())
		}
	}
	for i := range data.Sequences {
		keys := [...]string{"sequenceId"}
		keyValues := [...]string{strconv.FormatInt(data.Sequences[i].Id.ValueInt64(), 10)}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "sequences").ForEach(
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
		data.Sequences[i].Id = types.Int64Null()

		if t := r.Get("sequenceId.optionType"); t.Exists() {
			va := r.Get("sequenceId.value")
			if t.String() == "global" {
				data.Sequences[i].Id = types.Int64Value(va.Int())
			}
		}
		data.Sequences[i].Name = types.StringNull()

		if t := r.Get("sequenceName.optionType"); t.Exists() {
			va := r.Get("sequenceName.value")
			if t.String() == "global" {
				data.Sequences[i].Name = types.StringValue(va.String())
			}
		}
		data.Sequences[i].BaseAction = types.StringNull()

		if t := r.Get("baseAction.optionType"); t.Exists() {
			va := r.Get("baseAction.value")
			if t.String() == "global" {
				data.Sequences[i].BaseAction = types.StringValue(va.String())
			}
		}
		data.Sequences[i].DeviceAccessPort = types.Int64Null()

		if t := r.Get("matchEntries.destinationPort.optionType"); t.Exists() {
			va := r.Get("matchEntries.destinationPort.value")
			if t.String() == "global" {
				data.Sequences[i].DeviceAccessPort = types.Int64Value(va.Int())
			}
		}
		data.Sequences[i].DestinationDataPrefixListId = types.StringNull()

		if t := r.Get("matchEntries.destinationDataPrefix.destinationDataPrefixList.refId.optionType"); t.Exists() {
			va := r.Get("matchEntries.destinationDataPrefix.destinationDataPrefixList.refId.value")
			if t.String() == "global" {
				data.Sequences[i].DestinationDataPrefixListId = types.StringValue(va.String())
			}
		}
		data.Sequences[i].DestinationIpPrefixList = types.SetNull(types.StringType)
		data.Sequences[i].DestinationIpPrefixListVariable = types.StringNull()
		if t := r.Get("matchEntries.destinationDataPrefix.destinationIpPrefixList.optionType"); t.Exists() {
			va := r.Get("matchEntries.destinationDataPrefix.destinationIpPrefixList.value")
			if t.String() == "variable" {
				data.Sequences[i].DestinationIpPrefixListVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Sequences[i].DestinationIpPrefixList = helpers.GetStringSet(va.Array())
			}
		}
		data.Sequences[i].SourcePorts = types.SetNull(types.Int64Type)

		if t := r.Get("matchEntries.sourcePorts.optionType"); t.Exists() {
			va := r.Get("matchEntries.sourcePorts.value")
			if t.String() == "global" {
				data.Sequences[i].SourcePorts = helpers.GetInt64Set(va.Array())
			}
		}
		data.Sequences[i].SourceDataPrefixListId = types.StringNull()

		if t := r.Get("matchEntries.sourceDataPrefix.sourceDataPrefixList.refId.optionType"); t.Exists() {
			va := r.Get("matchEntries.sourceDataPrefix.sourceDataPrefixList.refId.value")
			if t.String() == "global" {
				data.Sequences[i].SourceDataPrefixListId = types.StringValue(va.String())
			}
		}
		data.Sequences[i].SourceIpPrefixList = types.SetNull(types.StringType)
		data.Sequences[i].SourceIpPrefixListVariable = types.StringNull()
		if t := r.Get("matchEntries.sourceDataPrefix.sourceIpPrefixList.optionType"); t.Exists() {
			va := r.Get("matchEntries.sourceDataPrefix.sourceIpPrefixList.value")
			if t.String() == "variable" {
				data.Sequences[i].SourceIpPrefixListVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Sequences[i].SourceIpPrefixList = helpers.GetStringSet(va.Array())
			}
		}
	}
}

// End of section. //template:end updateFromBody
