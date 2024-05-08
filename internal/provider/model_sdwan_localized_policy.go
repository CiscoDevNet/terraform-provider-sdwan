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
type LocalizedPolicy struct {
	Id                         types.String                 `tfsdk:"id"`
	Version                    types.Int64                  `tfsdk:"version"`
	Name                       types.String                 `tfsdk:"name"`
	Description                types.String                 `tfsdk:"description"`
	FlowVisibilityIpv4         types.Bool                   `tfsdk:"flow_visibility_ipv4"`
	FlowVisibilityIpv6         types.Bool                   `tfsdk:"flow_visibility_ipv6"`
	ApplicationVisibilityIpv4  types.Bool                   `tfsdk:"application_visibility_ipv4"`
	ApplicationVisibilityIpv6  types.Bool                   `tfsdk:"application_visibility_ipv6"`
	CloudQos                   types.Bool                   `tfsdk:"cloud_qos"`
	CloudQosServiceSide        types.Bool                   `tfsdk:"cloud_qos_service_side"`
	ImplicitAclLogging         types.Bool                   `tfsdk:"implicit_acl_logging"`
	LogFrequency               types.Int64                  `tfsdk:"log_frequency"`
	Ipv4VisibilityCacheEntries types.Int64                  `tfsdk:"ipv4_visibility_cache_entries"`
	Ipv6VisibilityCacheEntries types.Int64                  `tfsdk:"ipv6_visibility_cache_entries"`
	Definitions                []LocalizedPolicyDefinitions `tfsdk:"definitions"`
}

type LocalizedPolicyDefinitions struct {
	Id      types.String `tfsdk:"id"`
	Version types.Int64  `tfsdk:"version"`
	Type    types.String `tfsdk:"type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data LocalizedPolicy) getPath() string {
	return "/template/policy/vedge/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data LocalizedPolicy) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "policyType", "feature")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "policyName", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "policyDescription", data.Description.ValueString())
	}
	if !data.FlowVisibilityIpv4.IsNull() {
		if false && data.FlowVisibilityIpv4.ValueBool() {
			body, _ = sjson.Set(body, "policyDefinition.settings.flowVisibility", "")
		} else {
			body, _ = sjson.Set(body, "policyDefinition.settings.flowVisibility", data.FlowVisibilityIpv4.ValueBool())
		}
	}
	if !data.FlowVisibilityIpv6.IsNull() {
		if false && data.FlowVisibilityIpv6.ValueBool() {
			body, _ = sjson.Set(body, "policyDefinition.settings.flowVisibilityIPv6", "")
		} else {
			body, _ = sjson.Set(body, "policyDefinition.settings.flowVisibilityIPv6", data.FlowVisibilityIpv6.ValueBool())
		}
	}
	if !data.ApplicationVisibilityIpv4.IsNull() {
		if false && data.ApplicationVisibilityIpv4.ValueBool() {
			body, _ = sjson.Set(body, "policyDefinition.settings.appVisibility", "")
		} else {
			body, _ = sjson.Set(body, "policyDefinition.settings.appVisibility", data.ApplicationVisibilityIpv4.ValueBool())
		}
	}
	if !data.ApplicationVisibilityIpv6.IsNull() {
		if false && data.ApplicationVisibilityIpv6.ValueBool() {
			body, _ = sjson.Set(body, "policyDefinition.settings.appVisibilityIPv6", "")
		} else {
			body, _ = sjson.Set(body, "policyDefinition.settings.appVisibilityIPv6", data.ApplicationVisibilityIpv6.ValueBool())
		}
	}
	if !data.CloudQos.IsNull() {
		if false && data.CloudQos.ValueBool() {
			body, _ = sjson.Set(body, "policyDefinition.settings.cloudQos", "")
		} else {
			body, _ = sjson.Set(body, "policyDefinition.settings.cloudQos", data.CloudQos.ValueBool())
		}
	}
	if !data.CloudQosServiceSide.IsNull() {
		if false && data.CloudQosServiceSide.ValueBool() {
			body, _ = sjson.Set(body, "policyDefinition.settings.cloudQosServiceSide", "")
		} else {
			body, _ = sjson.Set(body, "policyDefinition.settings.cloudQosServiceSide", data.CloudQosServiceSide.ValueBool())
		}
	}
	if !data.ImplicitAclLogging.IsNull() {
		if false && data.ImplicitAclLogging.ValueBool() {
			body, _ = sjson.Set(body, "policyDefinition.settings.implicitAclLogging", "")
		} else {
			body, _ = sjson.Set(body, "policyDefinition.settings.implicitAclLogging", data.ImplicitAclLogging.ValueBool())
		}
	}
	if !data.LogFrequency.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.logFrequency", data.LogFrequency.ValueInt64())
	}
	if !data.Ipv4VisibilityCacheEntries.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.ipVisibilityCacheEntries", data.Ipv4VisibilityCacheEntries.ValueInt64())
	}
	if !data.Ipv6VisibilityCacheEntries.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.ipV6VisibilityCacheEntries", data.Ipv6VisibilityCacheEntries.ValueInt64())
	}
	if true {
		body, _ = sjson.Set(body, "policyDefinition.assembly", []interface{}{})
		for _, item := range data.Definitions {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "definitionId", item.Id.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			body, _ = sjson.SetRaw(body, "policyDefinition.assembly.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *LocalizedPolicy) fromBody(ctx context.Context, res gjson.Result) {
	state := *data
	if value := res.Get("policyName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("policyDescription"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.flowVisibility"); value.Exists() {
		if false && value.String() == "" {
			data.FlowVisibilityIpv4 = types.BoolValue(true)
		} else {
			data.FlowVisibilityIpv4 = types.BoolValue(value.Bool())
		}
	} else {
		data.FlowVisibilityIpv4 = types.BoolNull()
	}
	if value := res.Get("policyDefinition.settings.flowVisibilityIPv6"); value.Exists() {
		if false && value.String() == "" {
			data.FlowVisibilityIpv6 = types.BoolValue(true)
		} else {
			data.FlowVisibilityIpv6 = types.BoolValue(value.Bool())
		}
	} else {
		data.FlowVisibilityIpv6 = types.BoolNull()
	}
	if value := res.Get("policyDefinition.settings.appVisibility"); value.Exists() {
		if false && value.String() == "" {
			data.ApplicationVisibilityIpv4 = types.BoolValue(true)
		} else {
			data.ApplicationVisibilityIpv4 = types.BoolValue(value.Bool())
		}
	} else {
		data.ApplicationVisibilityIpv4 = types.BoolNull()
	}
	if value := res.Get("policyDefinition.settings.appVisibilityIPv6"); value.Exists() {
		if false && value.String() == "" {
			data.ApplicationVisibilityIpv6 = types.BoolValue(true)
		} else {
			data.ApplicationVisibilityIpv6 = types.BoolValue(value.Bool())
		}
	} else {
		data.ApplicationVisibilityIpv6 = types.BoolNull()
	}
	if value := res.Get("policyDefinition.settings.cloudQos"); value.Exists() {
		if false && value.String() == "" {
			data.CloudQos = types.BoolValue(true)
		} else {
			data.CloudQos = types.BoolValue(value.Bool())
		}
	} else {
		data.CloudQos = types.BoolNull()
	}
	if value := res.Get("policyDefinition.settings.cloudQosServiceSide"); value.Exists() {
		if false && value.String() == "" {
			data.CloudQosServiceSide = types.BoolValue(true)
		} else {
			data.CloudQosServiceSide = types.BoolValue(value.Bool())
		}
	} else {
		data.CloudQosServiceSide = types.BoolNull()
	}
	if value := res.Get("policyDefinition.settings.implicitAclLogging"); value.Exists() {
		if false && value.String() == "" {
			data.ImplicitAclLogging = types.BoolValue(true)
		} else {
			data.ImplicitAclLogging = types.BoolValue(value.Bool())
		}
	} else {
		data.ImplicitAclLogging = types.BoolNull()
	}
	if value := res.Get("policyDefinition.settings.logFrequency"); value.Exists() {
		data.LogFrequency = types.Int64Value(value.Int())
	} else {
		data.LogFrequency = types.Int64Null()
	}
	if value := res.Get("policyDefinition.settings.ipVisibilityCacheEntries"); value.Exists() {
		data.Ipv4VisibilityCacheEntries = types.Int64Value(value.Int())
	} else {
		data.Ipv4VisibilityCacheEntries = types.Int64Null()
	}
	if value := res.Get("policyDefinition.settings.ipV6VisibilityCacheEntries"); value.Exists() {
		data.Ipv6VisibilityCacheEntries = types.Int64Value(value.Int())
	} else {
		data.Ipv6VisibilityCacheEntries = types.Int64Null()
	}
	if value := res.Get("policyDefinition.assembly"); value.Exists() && len(value.Array()) > 0 {
		data.Definitions = make([]LocalizedPolicyDefinitions, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := LocalizedPolicyDefinitions{}
			if cValue := v.Get("definitionId"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("type"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			} else {
				item.Type = types.StringNull()
			}
			data.Definitions = append(data.Definitions, item)
			return true
		})
	} else {
		if len(data.Definitions) > 0 {
			data.Definitions = []LocalizedPolicyDefinitions{}
		}
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *LocalizedPolicy) hasChanges(ctx context.Context, state *LocalizedPolicy) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.FlowVisibilityIpv4.Equal(state.FlowVisibilityIpv4) {
		hasChanges = true
	}
	if !data.FlowVisibilityIpv6.Equal(state.FlowVisibilityIpv6) {
		hasChanges = true
	}
	if !data.ApplicationVisibilityIpv4.Equal(state.ApplicationVisibilityIpv4) {
		hasChanges = true
	}
	if !data.ApplicationVisibilityIpv6.Equal(state.ApplicationVisibilityIpv6) {
		hasChanges = true
	}
	if !data.CloudQos.Equal(state.CloudQos) {
		hasChanges = true
	}
	if !data.CloudQosServiceSide.Equal(state.CloudQosServiceSide) {
		hasChanges = true
	}
	if !data.ImplicitAclLogging.Equal(state.ImplicitAclLogging) {
		hasChanges = true
	}
	if !data.LogFrequency.Equal(state.LogFrequency) {
		hasChanges = true
	}
	if !data.Ipv4VisibilityCacheEntries.Equal(state.Ipv4VisibilityCacheEntries) {
		hasChanges = true
	}
	if !data.Ipv6VisibilityCacheEntries.Equal(state.Ipv6VisibilityCacheEntries) {
		hasChanges = true
	}
	if len(data.Definitions) != len(state.Definitions) {
		hasChanges = true
	} else {
		for i := range data.Definitions {
			if !data.Definitions[i].Id.Equal(state.Definitions[i].Id) {
				hasChanges = true
			}
			if !data.Definitions[i].Type.Equal(state.Definitions[i].Type) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *LocalizedPolicy) updateVersions(ctx context.Context, state *LocalizedPolicy) {
	for i := range data.Definitions {
		dataKeys := [...]string{fmt.Sprintf("%v", data.Definitions[i].Id.ValueString())}
		stateIndex := -1
		for j := range state.Definitions {
			stateKeys := [...]string{fmt.Sprintf("%v", state.Definitions[j].Id.ValueString())}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		if stateIndex > -1 {
			data.Definitions[i].Version = state.Definitions[stateIndex].Version
		} else {
			data.Definitions[i].Version = types.Int64Null()
		}
	}
}

// End of section. //template:end updateVersions
