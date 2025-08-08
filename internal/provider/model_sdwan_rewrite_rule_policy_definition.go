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
type RewriteRulePolicyDefinition struct {
	Id          types.String                       `tfsdk:"id"`
	Version     types.Int64                        `tfsdk:"version"`
	Type        types.String                       `tfsdk:"type"`
	Name        types.String                       `tfsdk:"name"`
	Description types.String                       `tfsdk:"description"`
	Rules       []RewriteRulePolicyDefinitionRules `tfsdk:"rules"`
}

type RewriteRulePolicyDefinitionRules struct {
	ClassMapId      types.String `tfsdk:"class_map_id"`
	ClassMapVersion types.Int64  `tfsdk:"class_map_version"`
	Priority        types.String `tfsdk:"priority"`
	Dscp            types.Int64  `tfsdk:"dscp"`
	Layer2Cos       types.Int64  `tfsdk:"layer2_cos"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data RewriteRulePolicyDefinition) getPath() string {
	return "/template/policy/definition/rewriterule/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data RewriteRulePolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "rewriteRule")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "definition.rules", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.ClassMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "class", item.ClassMapId.ValueString())
			}
			if !item.Priority.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "plp", item.Priority.ValueString())
			}
			if !item.Dscp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dscp", fmt.Sprint(item.Dscp.ValueInt64()))
			}
			if !item.Layer2Cos.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "layer2Cos", fmt.Sprint(item.Layer2Cos.ValueInt64()))
			}
			body, _ = sjson.SetRaw(body, "definition.rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *RewriteRulePolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("definition.rules"); value.Exists() && len(value.Array()) > 0 {
		data.Rules = make([]RewriteRulePolicyDefinitionRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RewriteRulePolicyDefinitionRules{}
			if cValue := v.Get("class"); cValue.Exists() {
				item.ClassMapId = types.StringValue(cValue.String())
			} else {
				item.ClassMapId = types.StringNull()
			}
			if cValue := v.Get("plp"); cValue.Exists() {
				item.Priority = types.StringValue(cValue.String())
			} else {
				item.Priority = types.StringNull()
			}
			if cValue := v.Get("dscp"); cValue.Exists() {
				item.Dscp = types.Int64Value(cValue.Int())
			} else {
				item.Dscp = types.Int64Null()
			}
			if cValue := v.Get("layer2Cos"); cValue.Exists() {
				item.Layer2Cos = types.Int64Value(cValue.Int())
			} else {
				item.Layer2Cos = types.Int64Null()
			}
			data.Rules = append(data.Rules, item)
			return true
		})
	} else {
		if len(data.Rules) > 0 {
			data.Rules = []RewriteRulePolicyDefinitionRules{}
		}
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *RewriteRulePolicyDefinition) hasChanges(ctx context.Context, state *RewriteRulePolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if len(data.Rules) != len(state.Rules) {
		hasChanges = true
	} else {
		for i := range data.Rules {
			if !data.Rules[i].ClassMapId.Equal(state.Rules[i].ClassMapId) {
				hasChanges = true
			}
			if !data.Rules[i].Priority.Equal(state.Rules[i].Priority) {
				hasChanges = true
			}
			if !data.Rules[i].Dscp.Equal(state.Rules[i].Dscp) {
				hasChanges = true
			}
			if !data.Rules[i].Layer2Cos.Equal(state.Rules[i].Layer2Cos) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *RewriteRulePolicyDefinition) updateVersions(ctx context.Context, state *RewriteRulePolicyDefinition) {
	for i := range data.Rules {
		dataKeys := [...]string{fmt.Sprintf("%v", data.Rules[i].ClassMapId.ValueString())}
		stateIndex := -1
		for j := range state.Rules {
			stateKeys := [...]string{fmt.Sprintf("%v", state.Rules[j].ClassMapId.ValueString())}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		if stateIndex > -1 {
			data.Rules[i].ClassMapVersion = state.Rules[stateIndex].ClassMapVersion
		} else {
			data.Rules[i].ClassMapVersion = types.Int64Null()
		}
	}
}

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *RewriteRulePolicyDefinition) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
	data.Type = types.StringValue("rewriteRule")
	for i := range data.Rules {
		if data.Rules[i].ClassMapId != types.StringNull() {
			data.Rules[i].ClassMapVersion = types.Int64Value(0)
		}
	}
}

// End of section. //template:end processImport
