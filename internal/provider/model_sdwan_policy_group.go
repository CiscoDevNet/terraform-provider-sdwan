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
	"slices"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type PolicyGroup struct {
	Id                types.String         `tfsdk:"id"`
	Name              types.String         `tfsdk:"name"`
	Description       types.String         `tfsdk:"description"`
	Solution          types.String         `tfsdk:"solution"`
	FeatureProfileIds types.Set            `tfsdk:"feature_profile_ids"`
	Devices           []PolicyGroupDevices `tfsdk:"devices"`
	PolicyVersions    types.List           `tfsdk:"policy_versions"`
}

type PolicyGroupDevices struct {
	Id        types.String                  `tfsdk:"id"`
	Deploy    types.Bool                    `tfsdk:"deploy"`
	Variables []PolicyGroupDevicesVariables `tfsdk:"variables"`
}

type PolicyGroupDevicesVariables struct {
	Name      types.String `tfsdk:"name"`
	Value     types.String `tfsdk:"value"`
	ListValue types.List   `tfsdk:"list_value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyGroup) getPath() string {
	return "/v1/policy-group/"
}

// End of section. //template:end getPath

func (data PolicyGroup) toBodyPolicyGroup(ctx context.Context) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Solution.IsNull() {
		body, _ = sjson.Set(body, "solution", data.Solution.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "profiles", []interface{}{})
		for _, item := range data.FeatureProfileIds.Elements() {
			itemBody := ""
			if !item.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", strings.Trim(item.String(), "\""))
			}
			body, _ = sjson.SetRaw(body, "profiles.-1", itemBody)
		}
	}
	return body
}

func (data PolicyGroup) toBodyPolicyGroupDevices(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "devices", []interface{}{})
		for _, item := range data.Devices {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "devices.-1", itemBody)
		}
	}
	return body
}

func (data PolicyGroup) toBodyPolicyGroupDeviceVariables(ctx context.Context) string {
	body := ""
	if !data.Solution.IsNull() {
		body, _ = sjson.Set(body, "solution", data.Solution.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "devices", []interface{}{})
		for _, item := range data.Devices {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "device-id", item.Id.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "variables", []interface{}{})
				for _, childItem := range item.Variables {
					itemChildBody := ""
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.ListValue.IsNull() {
						var values []string
						childItem.ListValue.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "value", values)
					} else if !childItem.Value.IsNull() {
						if val, err := strconv.Atoi(childItem.Value.ValueString()); err == nil {
							itemChildBody, _ = sjson.Set(itemChildBody, "value", val)
						} else if val, err := strconv.ParseFloat(childItem.Value.ValueString(), 64); err == nil {
							itemChildBody, _ = sjson.Set(itemChildBody, "value", val)
						} else if val, err := strconv.ParseBool(childItem.Value.ValueString()); err == nil {
							itemChildBody, _ = sjson.Set(itemChildBody, "value", val)
						} else {
							itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "variables.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "devices.-1", itemBody)
		}
	}
	return body
}

func (data *PolicyGroup) fromBodyPolicyGroup(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("solution"); value.Exists() {
		data.Solution = types.StringValue(value.String())
	} else {
		data.Solution = types.StringNull()
	}
	if value := res.Get("profiles"); value.Exists() && len(value.Array()) > 0 {
		a := make([]attr.Value, len(value.Array()))
		c := 0
		value.ForEach(func(k, v gjson.Result) bool {
			if cValue := v.Get("id"); cValue.Exists() {
				a[c] = types.StringValue(cValue.String())
			} else {
				a[c] = types.StringNull()
			}
			c += 1
			return true
		})
		data.FeatureProfileIds = types.SetValueMust(types.StringType, a)
	} else {
		data.FeatureProfileIds = types.SetNull(types.StringType)
	}
}

func (data *PolicyGroup) fromBodyPolicyGroupDevices(ctx context.Context, res gjson.Result) {
	original := *data
	if value := res.Get("devices"); value.Exists() && len(value.Array()) > 0 {
		data.Devices = make([]PolicyGroupDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyGroupDevices{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			data.Devices = append(data.Devices, item)
			return true
		})
	} else {
		if len(data.Devices) > 0 {
			data.Devices = []PolicyGroupDevices{}
		}
	}
	// reorder
	slices.Reverse(original.Devices)
	for i := range original.Devices {
		keyValues := [...]string{original.Devices[i].Id.ValueString()}

		for y := range data.Devices {
			found := false
			for _, keyValue := range keyValues {
				if !data.Devices[y].Id.IsNull() {
					if data.Devices[y].Id.ValueString() == keyValue {
						found = true
						continue
					}
					found = false
					break
				}
				continue
			}
			if found {
				//insert at the beginning
				device := data.Devices[y]
				data.Devices = append(data.Devices[:y], data.Devices[y+1:]...)
				data.Devices = append([]PolicyGroupDevices{device}, data.Devices...)
			}
		}
	}
}

func (data *PolicyGroup) fromBodyPolicyGroupDeviceVariables(ctx context.Context, res gjson.Result) {
	original := *data
	if value := res.Get("family"); value.Exists() {
		data.Solution = types.StringValue(value.String())
	} else {
		data.Solution = types.StringNull()
	}
	if value := res.Get("devices"); value.Exists() && len(value.Array()) > 0 {
		data.Devices = make([]PolicyGroupDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyGroupDevices{}
			if cValue := v.Get("device-id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("variables"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Variables = make([]PolicyGroupDevicesVariables, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					// skip optional variables
					if !cv.Get("value").Exists() {
						return true
					}
					cItem := PolicyGroupDevicesVariables{}
					if ccValue := cv.Get("name"); ccValue.Exists() {
						cItem.Name = types.StringValue(ccValue.String())
					} else {
						cItem.Name = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() {
						if ccValue.IsArray() {
							cItem.ListValue = helpers.GetStringList(ccValue.Array())
							cItem.Value = types.StringNull()
						} else {
							cItem.ListValue = types.ListNull(types.StringType)
							if !strings.Contains(strings.ToLower(ccValue.String()), "$crypt_cluster") {
								cItem.Value = types.StringValue(ccValue.String())
							}
						}
					} else {
						cItem.ListValue = types.ListNull(types.StringType)
						cItem.Value = types.StringNull()
					}
					item.Variables = append(item.Variables, cItem)
					return true
				})
				if len(item.Variables) == 0 {
					item.Variables = nil
				}
			} else {
				if len(item.Variables) > 0 {
					item.Variables = []PolicyGroupDevicesVariables{}
				}
			}
			data.Devices = append(data.Devices, item)
			return true
		})
	} else {
		if len(data.Devices) > 0 {
			data.Devices = []PolicyGroupDevices{}
		}
	}

	// reorder
	slices.Reverse(original.Devices)
	for i := range original.Devices {
		keyValues := [...]string{original.Devices[i].Id.ValueString()}

		for y := range data.Devices {
			found := false
			for _, keyValue := range keyValues {
				if !data.Devices[y].Id.IsNull() {
					if data.Devices[y].Id.ValueString() == keyValue {
						found = true
						continue
					}
					found = false
					break
				}
				continue
			}
			if found {
				//insert at the beginning
				device := data.Devices[y]
				data.Devices = append(data.Devices[:y], data.Devices[y+1:]...)
				data.Devices = append([]PolicyGroupDevices{device}, data.Devices...)
			}
		}
	}
}

func (data *PolicyGroup) updateTfAttributes(ctx context.Context, state *PolicyGroup) {
	//data.FeatureVersions = state.FeatureVersions
	for i := range data.Devices {
		dataKeys := [...]string{fmt.Sprintf("%v", data.Devices[i].Id.ValueString())}
		stateIndex := -1
		for j := range state.Devices {
			stateKeys := [...]string{fmt.Sprintf("%v", state.Devices[j].Id.ValueString())}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		if stateIndex > -1 {
			data.Devices[i].Deploy = state.Devices[stateIndex].Deploy
		} else {
			data.Devices[i].Deploy = types.BoolNull()
		}
	}
}

func (data PolicyGroup) hasPolicyGroupDeviceVariables(ctx context.Context) bool {
	for _, device := range data.Devices {
		if len(device.Variables) > 0 {
			return true
		}
	}
	return false
}

func (data PolicyGroup) getUpdatedDevices(ctx context.Context, state *PolicyGroup) []string {
	updatedDevices := make([]string, 0)
	for _, device := range data.Devices {
		for _, stateDevice := range state.Devices {
			if device.Id.ValueString() == stateDevice.Id.ValueString() {
				for _, variable := range device.Variables {
					found := false
					for _, stateVariable := range stateDevice.Variables {
						if variable.Name.ValueString() == stateVariable.Name.ValueString() {
							found = true
							if variable.Value.ValueString() != stateVariable.Value.ValueString() {
								if !slices.Contains(updatedDevices, device.Id.ValueString()) {
									updatedDevices = append(updatedDevices, device.Id.ValueString())
								}
							}
							if variable.ListValue.String() != stateVariable.ListValue.String() {
								if !slices.Contains(updatedDevices, device.Id.ValueString()) {
									updatedDevices = append(updatedDevices, device.Id.ValueString())
								}
							}
						}
					}
					if !found {
						if !slices.Contains(updatedDevices, device.Id.ValueString()) {
							updatedDevices = append(updatedDevices, device.Id.ValueString())
						}
					}
				}
				for _, stateVariable := range stateDevice.Variables {
					found := false
					for _, variable := range device.Variables {
						if variable.Name.ValueString() == stateVariable.Name.ValueString() {
							found = true
						}
					}
					if !found {
						if !slices.Contains(updatedDevices, device.Id.ValueString()) {
							updatedDevices = append(updatedDevices, device.Id.ValueString())
						}
					}
				}
			}
		}
	}
	return updatedDevices
}

func (data PolicyGroup) hasPolicyVersionChanges(ctx context.Context, state *PolicyGroup) bool {
	var planValues, stateValues []string
	data.PolicyVersions.ElementsAs(ctx, &planValues, false)
	state.PolicyVersions.ElementsAs(ctx, &stateValues, false)
	if len(planValues) != len(stateValues) {
		return true
	}
	for i := range planValues {
		if i >= len(stateValues) || planValues[i] != stateValues[i] {
			return true
		}
	}
	return false
}
