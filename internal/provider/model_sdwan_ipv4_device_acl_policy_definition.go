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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type IPv4DeviceACLPolicyDefinition struct {
	Id            types.String                             `tfsdk:"id"`
	Version       types.Int64                              `tfsdk:"version"`
	Type          types.String                             `tfsdk:"type"`
	Name          types.String                             `tfsdk:"name"`
	Description   types.String                             `tfsdk:"description"`
	DefaultAction types.String                             `tfsdk:"default_action"`
	Sequences     []IPv4DeviceACLPolicyDefinitionSequences `tfsdk:"sequences"`
}

type IPv4DeviceACLPolicyDefinitionSequences struct {
	Id            types.Int64                                           `tfsdk:"id"`
	Name          types.String                                          `tfsdk:"name"`
	BaseAction    types.String                                          `tfsdk:"base_action"`
	MatchEntries  []IPv4DeviceACLPolicyDefinitionSequencesMatchEntries  `tfsdk:"match_entries"`
	ActionEntries []IPv4DeviceACLPolicyDefinitionSequencesActionEntries `tfsdk:"action_entries"`
}

type IPv4DeviceACLPolicyDefinitionSequencesMatchEntries struct {
	Type                                 types.String `tfsdk:"type"`
	SourceIp                             types.String `tfsdk:"source_ip"`
	DestinationIp                        types.String `tfsdk:"destination_ip"`
	SourcePorts                          types.String `tfsdk:"source_ports"`
	DestinationPort                      types.Int64  `tfsdk:"destination_port"`
	SourceDataIpv4PrefixListId           types.String `tfsdk:"source_data_ipv4_prefix_list_id"`
	SourceDataIpv4PrefixListVersion      types.Int64  `tfsdk:"source_data_ipv4_prefix_list_version"`
	DestinationDataIpv4PrefixListId      types.String `tfsdk:"destination_data_ipv4_prefix_list_id"`
	DestinationDataIpv4PrefixListVersion types.Int64  `tfsdk:"destination_data_ipv4_prefix_list_version"`
}
type IPv4DeviceACLPolicyDefinitionSequencesActionEntries struct {
	Type        types.String `tfsdk:"type"`
	CounterName types.String `tfsdk:"counter_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data IPv4DeviceACLPolicyDefinition) getPath() string {
	return "/template/policy/definition/deviceaccesspolicy/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data IPv4DeviceACLPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "deviceAccessPolicy")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.DefaultAction.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.type", data.DefaultAction.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "sequences", []interface{}{})
		for _, item := range data.Sequences {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceId", item.Id.ValueInt64())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceName", item.Name.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "sequenceType", "deviceaccesspolicy")
			}
			if !item.BaseAction.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "baseAction", item.BaseAction.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "match.entries", []interface{}{})
				for _, childItem := range item.MatchEntries {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "field", childItem.Type.ValueString())
					}
					if !childItem.SourceIp.IsNull() && childItem.Type.ValueString() == "sourceIp" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.SourceIp.ValueString())
					}
					if !childItem.DestinationIp.IsNull() && childItem.Type.ValueString() == "destinationIp" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.DestinationIp.ValueString())
					}
					if !childItem.SourcePorts.IsNull() && childItem.Type.ValueString() == "sourcePort" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.SourcePorts.ValueString())
					}
					if !childItem.DestinationPort.IsNull() && childItem.Type.ValueString() == "destinationPort" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.DestinationPort.ValueInt64()))
					}
					if !childItem.SourceDataIpv4PrefixListId.IsNull() && childItem.Type.ValueString() == "sourceDataPrefixList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.SourceDataIpv4PrefixListId.ValueString())
					}
					if !childItem.DestinationDataIpv4PrefixListId.IsNull() && childItem.Type.ValueString() == "destinationDataPrefixList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.DestinationDataIpv4PrefixListId.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "match.entries.-1", itemChildBody)
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "actions", []interface{}{})
				for _, childItem := range item.ActionEntries {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.CounterName.IsNull() && childItem.Type.ValueString() == "count" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter", childItem.CounterName.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "actions.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "sequences.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *IPv4DeviceACLPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
	state := *data
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("defaultAction.type"); value.Exists() {
		data.DefaultAction = types.StringValue(value.String())
	} else {
		data.DefaultAction = types.StringNull()
	}
	if value := res.Get("sequences"); value.Exists() && len(value.Array()) > 0 {
		data.Sequences = make([]IPv4DeviceACLPolicyDefinitionSequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := IPv4DeviceACLPolicyDefinitionSequences{}
			if cValue := v.Get("sequenceId"); cValue.Exists() {
				item.Id = types.Int64Value(cValue.Int())
			} else {
				item.Id = types.Int64Null()
			}
			if cValue := v.Get("sequenceName"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("baseAction"); cValue.Exists() {
				item.BaseAction = types.StringValue(cValue.String())
			} else {
				item.BaseAction = types.StringNull()
			}
			if cValue := v.Get("match.entries"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.MatchEntries = make([]IPv4DeviceACLPolicyDefinitionSequencesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := IPv4DeviceACLPolicyDefinitionSequencesMatchEntries{}
					if ccValue := cv.Get("field"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "sourceIp" {
						cItem.SourceIp = types.StringValue(ccValue.String())
					} else {
						cItem.SourceIp = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "destinationIp" {
						cItem.DestinationIp = types.StringValue(ccValue.String())
					} else {
						cItem.DestinationIp = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "sourcePort" {
						cItem.SourcePorts = types.StringValue(ccValue.String())
					} else {
						cItem.SourcePorts = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "destinationPort" {
						cItem.DestinationPort = types.Int64Value(ccValue.Int())
					} else {
						cItem.DestinationPort = types.Int64Null()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "sourceDataPrefixList" {
						cItem.SourceDataIpv4PrefixListId = types.StringValue(ccValue.String())
					} else {
						cItem.SourceDataIpv4PrefixListId = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "destinationDataPrefixList" {
						cItem.DestinationDataIpv4PrefixListId = types.StringValue(ccValue.String())
					} else {
						cItem.DestinationDataIpv4PrefixListId = types.StringNull()
					}
					item.MatchEntries = append(item.MatchEntries, cItem)
					return true
				})
			} else {
				if len(item.MatchEntries) > 0 {
					item.MatchEntries = []IPv4DeviceACLPolicyDefinitionSequencesMatchEntries{}
				}
			}
			if cValue := v.Get("actions"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ActionEntries = make([]IPv4DeviceACLPolicyDefinitionSequencesActionEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := IPv4DeviceACLPolicyDefinitionSequencesActionEntries{}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && cItem.Type.ValueString() == "count" {
						cItem.CounterName = types.StringValue(ccValue.String())
					} else {
						cItem.CounterName = types.StringNull()
					}
					item.ActionEntries = append(item.ActionEntries, cItem)
					return true
				})
			} else {
				if len(item.ActionEntries) > 0 {
					item.ActionEntries = []IPv4DeviceACLPolicyDefinitionSequencesActionEntries{}
				}
			}
			data.Sequences = append(data.Sequences, item)
			return true
		})
	} else {
		if len(data.Sequences) > 0 {
			data.Sequences = []IPv4DeviceACLPolicyDefinitionSequences{}
		}
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *IPv4DeviceACLPolicyDefinition) hasChanges(ctx context.Context, state *IPv4DeviceACLPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.DefaultAction.Equal(state.DefaultAction) {
		hasChanges = true
	}
	if len(data.Sequences) != len(state.Sequences) {
		hasChanges = true
	} else {
		for i := range data.Sequences {
			if !data.Sequences[i].Id.Equal(state.Sequences[i].Id) {
				hasChanges = true
			}
			if !data.Sequences[i].Name.Equal(state.Sequences[i].Name) {
				hasChanges = true
			}
			if !data.Sequences[i].BaseAction.Equal(state.Sequences[i].BaseAction) {
				hasChanges = true
			}
			if len(data.Sequences[i].MatchEntries) != len(state.Sequences[i].MatchEntries) {
				hasChanges = true
			} else {
				for ii := range data.Sequences[i].MatchEntries {
					if !data.Sequences[i].MatchEntries[ii].Type.Equal(state.Sequences[i].MatchEntries[ii].Type) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].SourceIp.Equal(state.Sequences[i].MatchEntries[ii].SourceIp) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DestinationIp.Equal(state.Sequences[i].MatchEntries[ii].DestinationIp) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].SourcePorts.Equal(state.Sequences[i].MatchEntries[ii].SourcePorts) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DestinationPort.Equal(state.Sequences[i].MatchEntries[ii].DestinationPort) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].SourceDataIpv4PrefixListId.Equal(state.Sequences[i].MatchEntries[ii].SourceDataIpv4PrefixListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DestinationDataIpv4PrefixListId.Equal(state.Sequences[i].MatchEntries[ii].DestinationDataIpv4PrefixListId) {
						hasChanges = true
					}
				}
			}
			if len(data.Sequences[i].ActionEntries) != len(state.Sequences[i].ActionEntries) {
				hasChanges = true
			} else {
				for ii := range data.Sequences[i].ActionEntries {
					if !data.Sequences[i].ActionEntries[ii].Type.Equal(state.Sequences[i].ActionEntries[ii].Type) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].CounterName.Equal(state.Sequences[i].ActionEntries[ii].CounterName) {
						hasChanges = true
					}
				}
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *IPv4DeviceACLPolicyDefinition) updateVersions(ctx context.Context, state *IPv4DeviceACLPolicyDefinition) {
	for i := range data.Sequences {
		dataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].Id.ValueInt64()), fmt.Sprintf("%v", data.Sequences[i].Name.ValueString())}
		stateIndex := -1
		for j := range state.Sequences {
			stateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[j].Id.ValueInt64()), fmt.Sprintf("%v", state.Sequences[j].Name.ValueString())}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		for ii := range data.Sequences[i].MatchEntries {
			cDataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].MatchEntries[ii].Type.ValueString())}
			cStateIndex := -1
			if stateIndex > -1 {
				for jj := range state.Sequences[stateIndex].MatchEntries {
					cStateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[stateIndex].MatchEntries[jj].Type.ValueString())}
					if cDataKeys == cStateKeys {
						cStateIndex = jj
						break
					}
				}
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].SourceDataIpv4PrefixListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].SourceDataIpv4PrefixListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].SourceDataIpv4PrefixListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].DestinationDataIpv4PrefixListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].DestinationDataIpv4PrefixListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].DestinationDataIpv4PrefixListVersion = types.Int64Null()
			}
		}
	}
}

// End of section. //template:end updateVersions
