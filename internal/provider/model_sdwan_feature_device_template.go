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

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type FeatureDeviceTemplate struct {
	Id               types.String                           `tfsdk:"id"`
	Version          types.Int64                            `tfsdk:"version"`
	Name             types.String                           `tfsdk:"name"`
	Description      types.String                           `tfsdk:"description"`
	DeviceType       types.String                           `tfsdk:"device_type"`
	DeviceRole       types.String                           `tfsdk:"device_role"`
	PolicyId         types.String                           `tfsdk:"policy_id"`
	PolicyVersion    types.Int64                            `tfsdk:"policy_version"`
	SecurityPolicyId types.String                           `tfsdk:"security_policy_id"`
	GeneralTemplates []FeatureDeviceTemplateGeneralTemplate `tfsdk:"general_templates"`
}

type FeatureDeviceTemplateGeneralTemplate struct {
	Id           types.String                       `tfsdk:"id"`
	Version      types.Int64                        `tfsdk:"version"`
	Type         types.String                       `tfsdk:"type"`
	SubTemplates []FeatureDeviceTemplateSubTemplate `tfsdk:"sub_templates"`
}

type FeatureDeviceTemplateSubTemplate struct {
	Id           types.String                          `tfsdk:"id"`
	Version      types.Int64                           `tfsdk:"version"`
	Type         types.String                          `tfsdk:"type"`
	SubTemplates []FeatureDeviceTemplateSubSubTemplate `tfsdk:"sub_templates"`
}

type FeatureDeviceTemplateSubSubTemplate struct {
	Id      types.String `tfsdk:"id"`
	Version types.Int64  `tfsdk:"version"`
	Type    types.String `tfsdk:"type"`
}

func (data FeatureDeviceTemplate) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "deviceType", data.DeviceType.ValueString())
	body, _ = sjson.Set(body, "deviceRole", data.DeviceRole.ValueString())
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "configType", "template")
	body, _ = sjson.Set(body, "draftMode", false)
	body, _ = sjson.Set(body, "featureTemplateUidRange", []interface{}{})
	body, _ = sjson.Set(body, "policyId", data.PolicyId.ValueString())
	body, _ = sjson.Set(body, "securityPolicyId", data.SecurityPolicyId.ValueString())
	for _, item := range data.GeneralTemplates {
		itemBody := ""
		itemBody, _ = sjson.Set(itemBody, "templateId", item.Id.ValueString())
		itemBody, _ = sjson.Set(itemBody, "templateType", item.Type.ValueString())
		for _, sub := range item.SubTemplates {
			subBody := ""
			subBody, _ = sjson.Set(subBody, "templateId", sub.Id.ValueString())
			subBody, _ = sjson.Set(subBody, "templateType", sub.Type.ValueString())
			for _, subSub := range sub.SubTemplates {
				subSubBody := ""
				subSubBody, _ = sjson.Set(subSubBody, "templateId", subSub.Id.ValueString())
				subSubBody, _ = sjson.Set(subSubBody, "templateType", subSub.Type.ValueString())
				subBody, _ = sjson.SetRaw(subBody, "subTemplates.-1", subSubBody)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "subTemplates.-1", subBody)
		}
		body, _ = sjson.SetRaw(body, "generalTemplates.-1", itemBody)
	}
	return body
}

func (data *FeatureDeviceTemplate) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("securityPolicyId"); value.Exists() && value.String() != "" {
		data.SecurityPolicyId = types.StringValue(value.String())
	} else {
		data.SecurityPolicyId = types.StringNull()
	}
	if value := res.Get("generalTemplates"); value.Exists() {
		data.GeneralTemplates = make([]FeatureDeviceTemplateGeneralTemplate, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := FeatureDeviceTemplateGeneralTemplate{}
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
			if sValue := v.Get("subTemplates"); sValue.Exists() {
				item.SubTemplates = make([]FeatureDeviceTemplateSubTemplate, 0)
				sValue.ForEach(func(k, v gjson.Result) bool {
					sItem := FeatureDeviceTemplateSubTemplate{}
					if csValue := v.Get("templateId"); csValue.Exists() {
						sItem.Id = types.StringValue(csValue.String())
					} else {
						sItem.Id = types.StringNull()
					}
					if csValue := v.Get("templateType"); csValue.Exists() {
						sItem.Type = types.StringValue(csValue.String())
					} else {
						sItem.Type = types.StringNull()
					}
					if ssValue := v.Get("subTemplates"); ssValue.Exists() {
						sItem.SubTemplates = make([]FeatureDeviceTemplateSubSubTemplate, 0)
						ssValue.ForEach(func(k, v gjson.Result) bool {
							ssItem := FeatureDeviceTemplateSubSubTemplate{}
							if cssValue := v.Get("templateId"); cssValue.Exists() {
								ssItem.Id = types.StringValue(cssValue.String())
							} else {
								ssItem.Id = types.StringNull()
							}
							if cssValue := v.Get("templateType"); cssValue.Exists() {
								ssItem.Type = types.StringValue(cssValue.String())
							} else {
								ssItem.Type = types.StringNull()
							}
							sItem.SubTemplates = append(sItem.SubTemplates, ssItem)
							return true
						})
					}
					item.SubTemplates = append(item.SubTemplates, sItem)
					return true
				})
			}
			data.GeneralTemplates = append(data.GeneralTemplates, item)
			return true
		})
	}
}

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
		for gt := range data.GeneralTemplates {
			if !data.GeneralTemplates[gt].Id.Equal(state.GeneralTemplates[gt].Id) {
				hasChanges = true
			}
			if !data.GeneralTemplates[gt].Type.Equal(state.GeneralTemplates[gt].Type) {
				hasChanges = true
			}
			if len(data.GeneralTemplates[gt].SubTemplates) != len(state.GeneralTemplates[gt].SubTemplates) {
				hasChanges = true
			} else {
				for st := range data.GeneralTemplates[gt].SubTemplates {
					if !data.GeneralTemplates[gt].SubTemplates[st].Id.Equal(state.GeneralTemplates[gt].SubTemplates[st].Id) {
						hasChanges = true
					}
					if !data.GeneralTemplates[gt].SubTemplates[st].Type.Equal(state.GeneralTemplates[gt].SubTemplates[st].Type) {
						hasChanges = true
					}
					if len(data.GeneralTemplates[gt].SubTemplates[st].SubTemplates) != len(state.GeneralTemplates[gt].SubTemplates[st].SubTemplates) {
						hasChanges = true
					} else {
						for sst := range data.GeneralTemplates[gt].SubTemplates[st].SubTemplates {
							if !data.GeneralTemplates[gt].SubTemplates[st].SubTemplates[sst].Id.Equal(state.GeneralTemplates[gt].SubTemplates[st].SubTemplates[sst].Id) {
								hasChanges = true
							}
							if !data.GeneralTemplates[gt].SubTemplates[st].SubTemplates[sst].Type.Equal(state.GeneralTemplates[gt].SubTemplates[st].SubTemplates[sst].Type) {
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

func (data *FeatureDeviceTemplate) getGeneralTemplateVersion(ctx context.Context, id string) types.Int64 {
	for _, item := range data.GeneralTemplates {
		if item.Id.ValueString() == id {
			return item.Version
		}
	}
	return types.Int64Null()
}

func (data *FeatureDeviceTemplate) getSubTemplateVersion(ctx context.Context, id, subId string) types.Int64 {
	for _, item := range data.GeneralTemplates {
		if item.Id.ValueString() == id {
			for _, subItem := range item.SubTemplates {
				if subItem.Id.ValueString() == subId {
					return subItem.Version
				}
			}
		}
	}
	return types.Int64Null()
}

func (data *FeatureDeviceTemplate) getSubSubTemplateVersion(ctx context.Context, id, subId, subSubId string) types.Int64 {
	for _, item := range data.GeneralTemplates {
		if item.Id.ValueString() == id {
			for _, subItem := range item.SubTemplates {
				if subItem.Id.ValueString() == subId {
					for _, subSubItem := range subItem.SubTemplates {
						if subSubItem.Id.ValueString() == subSubId {
							return subSubItem.Version
						}
					}
				}
			}
		}
	}
	return types.Int64Null()
}

func (data *FeatureDeviceTemplate) updateTemplateVersions(ctx context.Context, state FeatureDeviceTemplate) {
	data.PolicyVersion = state.PolicyVersion
	for gt := range data.GeneralTemplates {
		id := data.GeneralTemplates[gt].Id.ValueString()
		data.GeneralTemplates[gt].Version = state.getGeneralTemplateVersion(ctx, id)
		for st := range data.GeneralTemplates[gt].SubTemplates {
			subId := data.GeneralTemplates[gt].SubTemplates[st].Id.ValueString()
			data.GeneralTemplates[gt].SubTemplates[st].Version = state.getSubTemplateVersion(ctx, id, subId)
			for sst := range data.GeneralTemplates[gt].SubTemplates[st].SubTemplates {
				subSubId := data.GeneralTemplates[gt].SubTemplates[st].SubTemplates[sst].Id.ValueString()
				data.GeneralTemplates[gt].SubTemplates[st].SubTemplates[sst].Version = state.getSubSubTemplateVersion(ctx, id, subId, subSubId)
			}
		}
	}
}
