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
type ZoneBasedFirewallPolicyDefinition struct {
	Id             types.String                                      `tfsdk:"id"`
	Version        types.Int64                                       `tfsdk:"version"`
	Name           types.String                                      `tfsdk:"name"`
	Description    types.String                                      `tfsdk:"description"`
	Mode           types.String                                      `tfsdk:"mode"`
	ApplyZonePairs []ZoneBasedFirewallPolicyDefinitionApplyZonePairs `tfsdk:"apply_zone_pairs"`
	DefaultAction  types.String                                      `tfsdk:"default_action"`
	Rules          []ZoneBasedFirewallPolicyDefinitionRules          `tfsdk:"rules"`
}

type ZoneBasedFirewallPolicyDefinitionApplyZonePairs struct {
	SourceZone      types.String `tfsdk:"source_zone"`
	DestinationZone types.String `tfsdk:"destination_zone"`
}

type ZoneBasedFirewallPolicyDefinitionRules struct {
	RuleOrder     types.Int64                                           `tfsdk:"rule_order"`
	RuleName      types.String                                          `tfsdk:"rule_name"`
	BaseAction    types.String                                          `tfsdk:"base_action"`
	IpType        types.String                                          `tfsdk:"ip_type"`
	MatchEntries  []ZoneBasedFirewallPolicyDefinitionRulesMatchEntries  `tfsdk:"match_entries"`
	ActionEntries []ZoneBasedFirewallPolicyDefinitionRulesActionEntries `tfsdk:"action_entries"`
}

type ZoneBasedFirewallPolicyDefinitionRulesMatchEntries struct {
	Type          types.String `tfsdk:"type"`
	PolicyId      types.String `tfsdk:"policy_id"`
	PolicyVersion types.Int64  `tfsdk:"policy_version"`
	Value         types.String `tfsdk:"value"`
	ProtocolType  types.String `tfsdk:"protocol_type"`
	ValueVariable types.String `tfsdk:"value_variable"`
}
type ZoneBasedFirewallPolicyDefinitionRulesActionEntries struct {
	Type types.String `tfsdk:"type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ZoneBasedFirewallPolicyDefinition) getPath() string {
	return "/template/policy/definition/zonebasedfw/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ZoneBasedFirewallPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "zoneBasedFW")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, "mode", data.Mode.ValueString())
	}
	if true && data.Mode.ValueString() == "security" {
		body, _ = sjson.Set(body, "definition.entries", []interface{}{})
		for _, item := range data.ApplyZonePairs {
			itemBody := ""
			if !item.SourceZone.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourceZone", item.SourceZone.ValueString())
			}
			if !item.DestinationZone.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destinationZone", item.DestinationZone.ValueString())
			}
			body, _ = sjson.SetRaw(body, "definition.entries.-1", itemBody)
		}
	}
	if !data.DefaultAction.IsNull() {
		body, _ = sjson.Set(body, "definition.defaultAction.type", data.DefaultAction.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "definition.sequences", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.RuleOrder.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceId", item.RuleOrder.ValueInt64())
			}
			if !item.RuleName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceName", item.RuleName.ValueString())
			}
			if !item.BaseAction.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "baseAction", item.BaseAction.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "sequenceType", "zoneBasedFW")
			}
			if !item.IpType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceIpType", item.IpType.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "match.entries", []interface{}{})
				for _, childItem := range item.MatchEntries {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "field", childItem.Type.ValueString())
					}
					if !childItem.PolicyId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.PolicyId.ValueString())
					}
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					if !childItem.ProtocolType.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "app", childItem.ProtocolType.ValueString())
					}
					if !childItem.ValueVariable.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "vipVariableName", childItem.ValueVariable.ValueString())
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
					itemBody, _ = sjson.SetRaw(itemBody, "actions.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "definition.sequences.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ZoneBasedFirewallPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("mode"); value.Exists() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("definition.entries"); value.Exists() && len(value.Array()) > 0 && data.Mode.ValueString() == "security" {
		data.ApplyZonePairs = make([]ZoneBasedFirewallPolicyDefinitionApplyZonePairs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ZoneBasedFirewallPolicyDefinitionApplyZonePairs{}
			if cValue := v.Get("sourceZone"); cValue.Exists() {
				item.SourceZone = types.StringValue(cValue.String())
			} else {
				item.SourceZone = types.StringNull()
			}
			if cValue := v.Get("destinationZone"); cValue.Exists() {
				item.DestinationZone = types.StringValue(cValue.String())
			} else {
				item.DestinationZone = types.StringNull()
			}
			data.ApplyZonePairs = append(data.ApplyZonePairs, item)
			return true
		})
	} else {
		if len(data.ApplyZonePairs) > 0 {
			data.ApplyZonePairs = []ZoneBasedFirewallPolicyDefinitionApplyZonePairs{}
		}
	}
	if value := res.Get("definition.defaultAction.type"); value.Exists() {
		data.DefaultAction = types.StringValue(value.String())
	} else {
		data.DefaultAction = types.StringNull()
	}
	if value := res.Get("definition.sequences"); value.Exists() && len(value.Array()) > 0 {
		data.Rules = make([]ZoneBasedFirewallPolicyDefinitionRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ZoneBasedFirewallPolicyDefinitionRules{}
			if cValue := v.Get("sequenceId"); cValue.Exists() {
				item.RuleOrder = types.Int64Value(cValue.Int())
			} else {
				item.RuleOrder = types.Int64Null()
			}
			if cValue := v.Get("sequenceName"); cValue.Exists() {
				item.RuleName = types.StringValue(cValue.String())
			} else {
				item.RuleName = types.StringNull()
			}
			if cValue := v.Get("baseAction"); cValue.Exists() {
				item.BaseAction = types.StringValue(cValue.String())
			} else {
				item.BaseAction = types.StringNull()
			}
			if cValue := v.Get("sequenceIpType"); cValue.Exists() {
				item.IpType = types.StringValue(cValue.String())
			} else {
				item.IpType = types.StringNull()
			}
			if cValue := v.Get("match.entries"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.MatchEntries = make([]ZoneBasedFirewallPolicyDefinitionRulesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ZoneBasedFirewallPolicyDefinitionRulesMatchEntries{}
					if ccValue := cv.Get("field"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() {
						cItem.PolicyId = types.StringValue(ccValue.String())
					} else {
						cItem.PolicyId = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() {
						cItem.Value = types.StringValue(ccValue.String())
					} else {
						cItem.Value = types.StringNull()
					}
					if ccValue := cv.Get("app"); ccValue.Exists() {
						cItem.ProtocolType = types.StringValue(ccValue.String())
					} else {
						cItem.ProtocolType = types.StringNull()
					}
					if ccValue := cv.Get("vipVariableName"); ccValue.Exists() {
						cItem.ValueVariable = types.StringValue(ccValue.String())
					} else {
						cItem.ValueVariable = types.StringNull()
					}
					item.MatchEntries = append(item.MatchEntries, cItem)
					return true
				})
			} else {
				if len(item.MatchEntries) > 0 {
					item.MatchEntries = []ZoneBasedFirewallPolicyDefinitionRulesMatchEntries{}
				}
			}
			if cValue := v.Get("actions"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ActionEntries = make([]ZoneBasedFirewallPolicyDefinitionRulesActionEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ZoneBasedFirewallPolicyDefinitionRulesActionEntries{}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					item.ActionEntries = append(item.ActionEntries, cItem)
					return true
				})
			} else {
				if len(item.ActionEntries) > 0 {
					item.ActionEntries = []ZoneBasedFirewallPolicyDefinitionRulesActionEntries{}
				}
			}
			data.Rules = append(data.Rules, item)
			return true
		})
	} else {
		if len(data.Rules) > 0 {
			data.Rules = []ZoneBasedFirewallPolicyDefinitionRules{}
		}
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *ZoneBasedFirewallPolicyDefinition) hasChanges(ctx context.Context, state *ZoneBasedFirewallPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.Mode.Equal(state.Mode) {
		hasChanges = true
	}
	if len(data.ApplyZonePairs) != len(state.ApplyZonePairs) {
		hasChanges = true
	} else {
		for i := range data.ApplyZonePairs {
			if !data.ApplyZonePairs[i].SourceZone.Equal(state.ApplyZonePairs[i].SourceZone) {
				hasChanges = true
			}
			if !data.ApplyZonePairs[i].DestinationZone.Equal(state.ApplyZonePairs[i].DestinationZone) {
				hasChanges = true
			}
		}
	}
	if !data.DefaultAction.Equal(state.DefaultAction) {
		hasChanges = true
	}
	if len(data.Rules) != len(state.Rules) {
		hasChanges = true
	} else {
		for i := range data.Rules {
			if !data.Rules[i].RuleOrder.Equal(state.Rules[i].RuleOrder) {
				hasChanges = true
			}
			if !data.Rules[i].RuleName.Equal(state.Rules[i].RuleName) {
				hasChanges = true
			}
			if !data.Rules[i].BaseAction.Equal(state.Rules[i].BaseAction) {
				hasChanges = true
			}
			if !data.Rules[i].IpType.Equal(state.Rules[i].IpType) {
				hasChanges = true
			}
			if len(data.Rules[i].MatchEntries) != len(state.Rules[i].MatchEntries) {
				hasChanges = true
			} else {
				for ii := range data.Rules[i].MatchEntries {
					if !data.Rules[i].MatchEntries[ii].Type.Equal(state.Rules[i].MatchEntries[ii].Type) {
						hasChanges = true
					}
					if !data.Rules[i].MatchEntries[ii].PolicyId.Equal(state.Rules[i].MatchEntries[ii].PolicyId) {
						hasChanges = true
					}
					if !data.Rules[i].MatchEntries[ii].Value.Equal(state.Rules[i].MatchEntries[ii].Value) {
						hasChanges = true
					}
					if !data.Rules[i].MatchEntries[ii].ProtocolType.Equal(state.Rules[i].MatchEntries[ii].ProtocolType) {
						hasChanges = true
					}
					if !data.Rules[i].MatchEntries[ii].ValueVariable.Equal(state.Rules[i].MatchEntries[ii].ValueVariable) {
						hasChanges = true
					}
				}
			}
			if len(data.Rules[i].ActionEntries) != len(state.Rules[i].ActionEntries) {
				hasChanges = true
			} else {
				for ii := range data.Rules[i].ActionEntries {
					if !data.Rules[i].ActionEntries[ii].Type.Equal(state.Rules[i].ActionEntries[ii].Type) {
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

func (data *ZoneBasedFirewallPolicyDefinition) updateVersions(ctx context.Context, state *ZoneBasedFirewallPolicyDefinition) {
	for i := range data.Rules {
		dataKeys := [...]string{fmt.Sprintf("%v", data.Rules[i].RuleOrder.ValueInt64())}
		stateIndex := -1
		for j := range state.Rules {
			stateKeys := [...]string{fmt.Sprintf("%v", state.Rules[j].RuleOrder.ValueInt64())}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		for ii := range data.Rules[i].MatchEntries {
			cDataKeys := [...]string{fmt.Sprintf("%v", data.Rules[i].MatchEntries[ii].Type.ValueString())}
			cStateIndex := -1
			if stateIndex > -1 {
				for jj := range state.Rules[stateIndex].MatchEntries {
					cStateKeys := [...]string{fmt.Sprintf("%v", state.Rules[stateIndex].MatchEntries[jj].Type.ValueString())}
					if cDataKeys == cStateKeys {
						cStateIndex = jj
						break
					}
				}
			}
			if cStateIndex > -1 {
				data.Rules[i].MatchEntries[ii].PolicyVersion = state.Rules[stateIndex].MatchEntries[cStateIndex].PolicyVersion
			} else {
				data.Rules[i].MatchEntries[ii].PolicyVersion = types.Int64Null()
			}
		}
	}
}

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *ZoneBasedFirewallPolicyDefinition) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
	for i := range data.Rules {
		for ii := range data.Rules[i].MatchEntries {
			if data.Rules[i].MatchEntries[ii].PolicyId != types.StringNull() {
				data.Rules[i].MatchEntries[ii].PolicyVersion = types.Int64Value(0)
			}
		}
	}
}

// End of section. //template:end processImport
