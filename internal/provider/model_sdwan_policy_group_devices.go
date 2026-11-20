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
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// PolicyGroupDevicesConfig is the Terraform shape of the resource (plan, state, config). devices
// is a map keyed by device ID, so Terraform matches a device by its ID: adding, removing or
// changing one device shows only that device in the plan. The code works on PolicyGroupDevices,
// the same data with the devices as a list sorted by ID; use getPolicyGroupDevices and toConfig to
// convert.
type PolicyGroupDevicesConfig struct {
	Id                 types.String                              `tfsdk:"id"`
	PolicyGroupId      types.String                              `tfsdk:"policy_group_id"`
	Solution           types.String                              `tfsdk:"solution"`
	PolicyGroupVersion types.Int64                               `tfsdk:"policy_group_version"`
	RetainDevices      types.Set                                 `tfsdk:"retain_devices"`
	Devices            map[string]PolicyGroupDevicesDeviceConfig `tfsdk:"devices"`
}

// PolicyGroupDevicesDeviceConfig is one device in the map; its ID is the map key.
type PolicyGroupDevicesDeviceConfig struct {
	Deploy    types.Bool                                  `tfsdk:"deploy"`
	Variables map[string]PolicyGroupDevicesVariableConfig `tfsdk:"variables"`
}

// PolicyGroupDevicesVariableConfig is one variable in the map; its name is the map key.
type PolicyGroupDevicesVariableConfig struct {
	Value     types.String `tfsdk:"value"`
	ListValue types.List   `tfsdk:"list_value"`
}

// PolicyGroupDevices is the working model: devices as a list sorted by ID.
type PolicyGroupDevices struct {
	Id                 types.String
	PolicyGroupId      types.String
	Solution           types.String
	PolicyGroupVersion types.Int64
	RetainDevices      types.Set
	Devices            []PolicyGroupDevicesDevices
}

type PolicyGroupDevicesDevices struct {
	Id        types.String
	Deploy    types.Bool
	Variables []PolicyGroupDevicesDevicesVariables
}

type PolicyGroupDevicesDevicesVariables struct {
	Name      types.String
	Value     types.String
	ListValue types.List
}

func (data PolicyGroupDevices) getPath() string {
	return "/v1/policy-group/"
}

// getPolicyGroupDevices reads a plan or state into the working model.
func getPolicyGroupDevices(ctx context.Context, src stateGetter, data *PolicyGroupDevices) diag.Diagnostics {
	var config PolicyGroupDevicesConfig
	diags := src.Get(ctx, &config)
	*data = config.toModel()
	return diags
}

func (config PolicyGroupDevicesConfig) toModel() PolicyGroupDevices {
	data := PolicyGroupDevices{
		Id:                 config.Id,
		PolicyGroupId:      config.PolicyGroupId,
		Solution:           config.Solution,
		PolicyGroupVersion: config.PolicyGroupVersion,
		RetainDevices:      config.RetainDevices,
	}
	if config.Devices != nil {
		data.Devices = make([]PolicyGroupDevicesDevices, 0, len(config.Devices))
		for _, id := range slices.Sorted(maps.Keys(config.Devices)) {
			d := config.Devices[id]
			data.Devices = append(data.Devices, PolicyGroupDevicesDevices{
				Id:        types.StringValue(id),
				Deploy:    d.Deploy,
				Variables: variablesToModelPolicyGroupDevices(d.Variables),
			})
		}
	}
	return data
}

// toConfig returns the Terraform shape, for resp.State.Set.
func (data PolicyGroupDevices) toConfig() PolicyGroupDevicesConfig {
	config := PolicyGroupDevicesConfig{
		Id:                 data.Id,
		PolicyGroupId:      data.PolicyGroupId,
		Solution:           data.Solution,
		PolicyGroupVersion: data.PolicyGroupVersion,
		RetainDevices:      data.RetainDevices,
	}
	if data.Devices != nil {
		config.Devices = make(map[string]PolicyGroupDevicesDeviceConfig, len(data.Devices))
		for _, d := range data.Devices {
			config.Devices[d.Id.ValueString()] = PolicyGroupDevicesDeviceConfig{
				Deploy:    d.Deploy,
				Variables: variablesToConfigPolicyGroupDevices(d.Variables),
			}
		}
	}
	return config
}

// variablesToModelPolicyGroupDevices returns the variables sorted by name; null stays null.
func variablesToModelPolicyGroupDevices(vars map[string]PolicyGroupDevicesVariableConfig) []PolicyGroupDevicesDevicesVariables {
	if vars == nil {
		return nil
	}
	out := make([]PolicyGroupDevicesDevicesVariables, 0, len(vars))
	for _, name := range slices.Sorted(maps.Keys(vars)) {
		out = append(out, PolicyGroupDevicesDevicesVariables{Name: types.StringValue(name), Value: vars[name].Value, ListValue: vars[name].ListValue})
	}
	return out
}

func variablesToConfigPolicyGroupDevices(vars []PolicyGroupDevicesDevicesVariables) map[string]PolicyGroupDevicesVariableConfig {
	if vars == nil {
		return nil
	}
	out := make(map[string]PolicyGroupDevicesVariableConfig, len(vars))
	for _, v := range vars {
		listValue := v.ListValue
		if listValue.IsNull() {
			listValue = types.ListNull(types.StringType)
		}
		out[v.Name.ValueString()] = PolicyGroupDevicesVariableConfig{Value: v.Value, ListValue: listValue}
	}
	return out
}

func (data PolicyGroupDevices) toBodyPolicyGroupDevices(ctx context.Context) string {
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

func (data PolicyGroupDevices) toBodyPolicyGroupDeviceVariables(ctx context.Context, varTypes map[string]string) string {
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

						varName := childItem.Name.ValueString()
						// Convert each element in the list based on schema type
						convertedValues := make([]interface{}, len(values))
						schemaType := varTypes[varName] // empty string if not found
						for i, valueStr := range values {
							convertedValues[i] = convertValueByType(valueStr, schemaType)
						}
						itemChildBody, _ = sjson.Set(itemChildBody, "value", convertedValues)
					} else if !childItem.Value.IsNull() {
						valueStr := childItem.Value.ValueString()
						varName := childItem.Name.ValueString()
						schemaType := varTypes[varName] // empty string if not found
						convertedValue := convertValueByType(valueStr, schemaType)
						itemChildBody, _ = sjson.Set(itemChildBody, "value", convertedValue)
					}
					itemBody, _ = sjson.SetRaw(itemBody, "variables.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "devices.-1", itemBody)
		}
	}
	return body
}

func (data *PolicyGroupDevices) fromBodyPolicyGroupDevices(ctx context.Context, res gjson.Result) {
	original := *data
	// Only manage devices already tracked in this resource's state; devices associated with the same
	// policy group by other resources / state files are ignored (state-scoped isolation).
	stateDeviceIds := make(map[string]bool)
	for _, d := range original.Devices {
		stateDeviceIds[d.Id.ValueString()] = true
	}
	if value := res.Get("devices"); value.Exists() && len(value.Array()) > 0 {
		data.Devices = make([]PolicyGroupDevicesDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			if id := v.Get("id").String(); !stateDeviceIds[id] {
				return true
			}
			item := PolicyGroupDevicesDevices{}
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
			data.Devices = []PolicyGroupDevicesDevices{}
		}
	}
	// No reordering: devices are a map in the state, keyed by ID.
}

// priorVariableValue returns the value variable name had for device id in data, if it had one.
func (data *PolicyGroupDevices) priorVariableValue(id, name string) (types.String, bool) {
	for _, d := range data.Devices {
		if d.Id.ValueString() != id {
			continue
		}
		for _, v := range d.Variables {
			if v.Name.ValueString() == name && !v.Value.IsNull() {
				return v.Value, true
			}
		}
		break
	}
	return types.StringNull(), false
}

// fromBodyPolicyGroupDeviceVariables reads the variables of the devices in data from the Manager.
// The Manager returns a secret variable as a placeholder ($cRYPT_CLUSTER1), never its value, so for
// those the value in original (the earlier state or plan) is kept. With none, e.g. after an import,
// the value stays null.
func (data *PolicyGroupDevices) fromBodyPolicyGroupDeviceVariables(ctx context.Context, res gjson.Result, original *PolicyGroupDevices) {
	if value := res.Get("family"); value.Exists() {
		data.Solution = types.StringValue(value.String())
	} else {
		data.Solution = types.StringNull()
	}

	// Update variables on existing (state-owned) devices in place, so devices managed by other
	// resources are not pulled into this resource's state.
	if value := res.Get("devices"); value.Exists() && len(value.Array()) > 0 {
		value.ForEach(func(k, v gjson.Result) bool {
			deviceId := v.Get("device-id").String()
			for i := range data.Devices {
				if data.Devices[i].Id.ValueString() != deviceId {
					continue
				}
				if cValue := v.Get("variables"); cValue.Exists() && len(cValue.Array()) > 0 {
					data.Devices[i].Variables = make([]PolicyGroupDevicesDevicesVariables, 0)
					cValue.ForEach(func(ck, cv gjson.Result) bool {
						// skip optional variables
						if !cv.Get("value").Exists() {
							return true
						}
						cItem := PolicyGroupDevicesDevicesVariables{}
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
								} else {
									cItem.Value = types.StringNull()
									if original != nil {
										if prior, ok := original.priorVariableValue(deviceId, cItem.Name.ValueString()); ok {
											cItem.Value = prior
										}
									}
								}
							}
						} else {
							cItem.ListValue = types.ListNull(types.StringType)
							cItem.Value = types.StringNull()
						}
						data.Devices[i].Variables = append(data.Devices[i].Variables, cItem)
						return true
					})
					if len(data.Devices[i].Variables) == 0 {
						data.Devices[i].Variables = nil
					}
				} else {
					if len(data.Devices[i].Variables) > 0 {
						data.Devices[i].Variables = []PolicyGroupDevicesDevicesVariables{}
					}
				}
				break
			}
			return true
		})
	}
}

func (data *PolicyGroupDevices) updateTfAttributes(ctx context.Context, state *PolicyGroupDevices) {
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

func (data PolicyGroupDevices) hasPolicyGroupDeviceVariables(ctx context.Context) bool {
	for _, device := range data.Devices {
		if len(device.Variables) > 0 {
			return true
		}
	}
	return false
}

func (data PolicyGroupDevices) getUpdatedDevices(ctx context.Context, state *PolicyGroupDevices) []string {
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

func (data *PolicyGroupDevices) processImport(ctx context.Context) {
	for i := range data.Devices {
		data.Devices[i].Deploy = types.BoolValue(true)
	}
}
