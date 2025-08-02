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
type FeatureDeviceTemplate struct {
	Id                    types.String                            `tfsdk:"id"`
	Version               types.Int64                             `tfsdk:"version"`
	Name                  types.String                            `tfsdk:"name"`
	Description           types.String                            `tfsdk:"description"`
	DeviceType            types.String                            `tfsdk:"device_type"`
	DeviceRole            types.String                            `tfsdk:"device_role"`
	PolicyId              types.String                            `tfsdk:"policy_id"`
	PolicyVersion         types.Int64                             `tfsdk:"policy_version"`
	SecurityPolicyId      types.String                            `tfsdk:"security_policy_id"`
	SecurityPolicyVersion types.Int64                             `tfsdk:"security_policy_version"`
	GeneralTemplates      []FeatureDeviceTemplateGeneralTemplates `tfsdk:"general_templates"`
}

type FeatureDeviceTemplateGeneralTemplates struct {
	Id           types.String                                        `tfsdk:"id"`
	Version      types.Int64                                         `tfsdk:"version"`
	Type         types.String                                        `tfsdk:"type"`
	SubTemplates []FeatureDeviceTemplateGeneralTemplatesSubTemplates `tfsdk:"sub_templates"`
}

type FeatureDeviceTemplateGeneralTemplatesSubTemplates struct {
	Id           types.String                                                    `tfsdk:"id"`
	Version      types.Int64                                                     `tfsdk:"version"`
	Type         types.String                                                    `tfsdk:"type"`
	SubTemplates []FeatureDeviceTemplateGeneralTemplatesSubTemplatesSubTemplates `tfsdk:"sub_templates"`
}

type FeatureDeviceTemplateGeneralTemplatesSubTemplatesSubTemplates struct {
	Id      types.String `tfsdk:"id"`
	Version types.Int64  `tfsdk:"version"`
	Type    types.String `tfsdk:"type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data FeatureDeviceTemplate) getPath() string {
	return "/template/device/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data FeatureDeviceTemplate) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "configType", "template")
	}
	if true {
		body, _ = sjson.Set(body, "factoryDefault", false)
	}
	if true {
		body, _ = sjson.Set(body, "featureTemplateUidRange", []interface{}{})
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	}
	if !data.DeviceType.IsNull() {
		body, _ = sjson.Set(body, "deviceType", data.DeviceType.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "deviceRole", data.DeviceRole.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "policyId", data.PolicyId.ValueString())
	}
	if !data.SecurityPolicyId.IsNull() {
		body, _ = sjson.Set(body, "securityPolicyId", data.SecurityPolicyId.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "generalTemplates", []interface{}{})
		for _, item := range data.GeneralTemplates {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "templateId", item.Id.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "templateType", item.Type.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "subTemplates", []interface{}{})
				for _, childItem := range item.SubTemplates {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "templateId", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "templateType", childItem.Type.ValueString())
					}
					if true {
						itemChildBody, _ = sjson.Set(itemChildBody, "subTemplates", []interface{}{})
						for _, childChildItem := range childItem.SubTemplates {
							itemChildChildBody := ""
							if !childChildItem.Id.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "templateId", childChildItem.Id.ValueString())
							}
							if !childChildItem.Type.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "templateType", childChildItem.Type.ValueString())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "subTemplates.-1", itemChildChildBody)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "subTemplates.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "generalTemplates.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *FeatureDeviceTemplate) fromBody(ctx context.Context, res gjson.Result) {
	state := *data
	if value := res.Get("templateName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("templateDescription"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("deviceType"); value.Exists() {
		data.DeviceType = types.StringValue(value.String())
	} else {
		data.DeviceType = types.StringNull()
	}
	if value := res.Get("deviceRole"); value.Exists() && value.String() != "" {
		data.DeviceRole = types.StringValue(value.String())
	} else {
		data.DeviceRole = types.StringNull()
	}
	if value := res.Get("policyId"); value.Exists() && value.String() != "" {
		data.PolicyId = types.StringValue(value.String())
	} else {
		data.PolicyId = types.StringNull()
	}
	if value := res.Get("securityPolicyId"); value.Exists() {
		data.SecurityPolicyId = types.StringValue(value.String())
	} else {
		data.SecurityPolicyId = types.StringNull()
	}
	if value := res.Get("generalTemplates"); value.Exists() && len(value.Array()) > 0 {
		data.GeneralTemplates = make([]FeatureDeviceTemplateGeneralTemplates, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := FeatureDeviceTemplateGeneralTemplates{}
			if cValue := v.Get("templateId"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("templateType"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			} else {
				item.Type = types.StringNull()
			}
			if cValue := v.Get("subTemplates"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.SubTemplates = make([]FeatureDeviceTemplateGeneralTemplatesSubTemplates, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := FeatureDeviceTemplateGeneralTemplatesSubTemplates{}
					if ccValue := cv.Get("templateId"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					if ccValue := cv.Get("templateType"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("subTemplates"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.SubTemplates = make([]FeatureDeviceTemplateGeneralTemplatesSubTemplatesSubTemplates, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := FeatureDeviceTemplateGeneralTemplatesSubTemplatesSubTemplates{}
							if cccValue := ccv.Get("templateId"); cccValue.Exists() {
								ccItem.Id = types.StringValue(cccValue.String())
							} else {
								ccItem.Id = types.StringNull()
							}
							if cccValue := ccv.Get("templateType"); cccValue.Exists() {
								ccItem.Type = types.StringValue(cccValue.String())
							} else {
								ccItem.Type = types.StringNull()
							}
							cItem.SubTemplates = append(cItem.SubTemplates, ccItem)
							return true
						})
					} else {
						if len(cItem.SubTemplates) > 0 {
							cItem.SubTemplates = []FeatureDeviceTemplateGeneralTemplatesSubTemplatesSubTemplates{}
						}
					}
					item.SubTemplates = append(item.SubTemplates, cItem)
					return true
				})
			} else {
				if len(item.SubTemplates) > 0 {
					item.SubTemplates = []FeatureDeviceTemplateGeneralTemplatesSubTemplates{}
				}
			}
			data.GeneralTemplates = append(data.GeneralTemplates, item)
			return true
		})
	} else {
		if len(data.GeneralTemplates) > 0 {
			data.GeneralTemplates = []FeatureDeviceTemplateGeneralTemplates{}
		}
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *FeatureDeviceTemplate) hasChanges(ctx context.Context, state *FeatureDeviceTemplate) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.DeviceType.Equal(state.DeviceType) {
		hasChanges = true
	}
	if !data.DeviceRole.Equal(state.DeviceRole) {
		hasChanges = true
	}
	if !data.PolicyId.Equal(state.PolicyId) {
		hasChanges = true
	}
	if !data.SecurityPolicyId.Equal(state.SecurityPolicyId) {
		hasChanges = true
	}
	if len(data.GeneralTemplates) != len(state.GeneralTemplates) {
		hasChanges = true
	} else {
		for i := range data.GeneralTemplates {
			if !data.GeneralTemplates[i].Id.Equal(state.GeneralTemplates[i].Id) {
				hasChanges = true
			}
			if !data.GeneralTemplates[i].Type.Equal(state.GeneralTemplates[i].Type) {
				hasChanges = true
			}
			if len(data.GeneralTemplates[i].SubTemplates) != len(state.GeneralTemplates[i].SubTemplates) {
				hasChanges = true
			} else {
				for ii := range data.GeneralTemplates[i].SubTemplates {
					if !data.GeneralTemplates[i].SubTemplates[ii].Id.Equal(state.GeneralTemplates[i].SubTemplates[ii].Id) {
						hasChanges = true
					}
					if !data.GeneralTemplates[i].SubTemplates[ii].Type.Equal(state.GeneralTemplates[i].SubTemplates[ii].Type) {
						hasChanges = true
					}
					if len(data.GeneralTemplates[i].SubTemplates[ii].SubTemplates) != len(state.GeneralTemplates[i].SubTemplates[ii].SubTemplates) {
						hasChanges = true
					} else {
						for iii := range data.GeneralTemplates[i].SubTemplates[ii].SubTemplates {
							if !data.GeneralTemplates[i].SubTemplates[ii].SubTemplates[iii].Id.Equal(state.GeneralTemplates[i].SubTemplates[ii].SubTemplates[iii].Id) {
								hasChanges = true
							}
							if !data.GeneralTemplates[i].SubTemplates[ii].SubTemplates[iii].Type.Equal(state.GeneralTemplates[i].SubTemplates[ii].SubTemplates[iii].Type) {
								hasChanges = true
							}
						}
					}
				}
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *FeatureDeviceTemplate) updateVersions(ctx context.Context, state *FeatureDeviceTemplate) {
	data.PolicyVersion = state.PolicyVersion
	data.SecurityPolicyVersion = state.SecurityPolicyVersion
	for i := range data.GeneralTemplates {
		dataKeys := [...]string{fmt.Sprintf("%v", data.GeneralTemplates[i].Id.ValueString())}
		stateIndex := -1
		for j := range state.GeneralTemplates {
			stateKeys := [...]string{fmt.Sprintf("%v", state.GeneralTemplates[j].Id.ValueString())}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		if stateIndex > -1 {
			data.GeneralTemplates[i].Version = state.GeneralTemplates[stateIndex].Version
		} else {
			data.GeneralTemplates[i].Version = types.Int64Null()
		}
		for ii := range data.GeneralTemplates[i].SubTemplates {
			cDataKeys := [...]string{fmt.Sprintf("%v", data.GeneralTemplates[i].SubTemplates[ii].Id.ValueString())}
			cStateIndex := -1
			if stateIndex > -1 {
				for jj := range state.GeneralTemplates[stateIndex].SubTemplates {
					cStateKeys := [...]string{fmt.Sprintf("%v", state.GeneralTemplates[stateIndex].SubTemplates[jj].Id.ValueString())}
					if cDataKeys == cStateKeys {
						cStateIndex = jj
						break
					}
				}
			}
			if cStateIndex > -1 {
				data.GeneralTemplates[i].SubTemplates[ii].Version = state.GeneralTemplates[stateIndex].SubTemplates[cStateIndex].Version
			} else {
				data.GeneralTemplates[i].SubTemplates[ii].Version = types.Int64Null()
			}
			for iii := range data.GeneralTemplates[i].SubTemplates[ii].SubTemplates {
				ccDataKeys := [...]string{fmt.Sprintf("%v", data.GeneralTemplates[i].SubTemplates[ii].SubTemplates[iii].Id.ValueString())}
				ccStateIndex := -1
				if cStateIndex > -1 {
					for jjj := range state.GeneralTemplates[stateIndex].SubTemplates[cStateIndex].SubTemplates {
						ccStateKeys := [...]string{fmt.Sprintf("%v", state.GeneralTemplates[stateIndex].SubTemplates[cStateIndex].SubTemplates[jjj].Id.ValueString())}
						if ccDataKeys == ccStateKeys {
							ccStateIndex = jjj
							break
						}
					}
				}
				if ccStateIndex > -1 {
					data.GeneralTemplates[i].SubTemplates[ii].SubTemplates[iii].Version = state.GeneralTemplates[stateIndex].SubTemplates[cStateIndex].SubTemplates[ccStateIndex].Version
				} else {
					data.GeneralTemplates[i].SubTemplates[ii].SubTemplates[iii].Version = types.Int64Null()
				}
			}
		}
	}
}

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *FeatureDeviceTemplate) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
	if data.PolicyId != types.StringNull() {
		data.PolicyVersion = types.Int64Value(0)
	}
	if data.SecurityPolicyId != types.StringNull() {
		data.SecurityPolicyVersion = types.Int64Value(0)
	}
	for i := range data.GeneralTemplates {
		if data.GeneralTemplates[i].Id != types.StringNull() {
			data.GeneralTemplates[i].Version = types.Int64Value(0)
		}
		for ii := range data.GeneralTemplates[i].SubTemplates {
			if data.GeneralTemplates[i].SubTemplates[ii].Id != types.StringNull() {
				data.GeneralTemplates[i].SubTemplates[ii].Version = types.Int64Value(0)
			}
			for iii := range data.GeneralTemplates[i].SubTemplates[ii].SubTemplates {
				if data.GeneralTemplates[i].SubTemplates[ii].SubTemplates[iii].Id != types.StringNull() {
					data.GeneralTemplates[i].SubTemplates[ii].SubTemplates[iii].Version = types.Int64Value(0)
				}
			}
		}
	}
}

// End of section. //template:end processImport
