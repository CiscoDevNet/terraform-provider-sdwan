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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ConfigurationGroupDeviceVariables struct {
	Id                   types.String                               `tfsdk:"id"`
	ConfigurationGroupId types.String                               `tfsdk:"configuration_group_id"`
	Solution             types.String                               `tfsdk:"solution"`
	Devices              []ConfigurationGroupDeviceVariablesDevices `tfsdk:"devices"`
	Groups               []ConfigurationGroupDeviceVariablesGroups  `tfsdk:"groups"`
}

type ConfigurationGroupDeviceVariablesDevices struct {
	DeviceId  types.String                                        `tfsdk:"device_id"`
	Variables []ConfigurationGroupDeviceVariablesDevicesVariables `tfsdk:"variables"`
}

type ConfigurationGroupDeviceVariablesGroups struct {
	Name      types.String                                       `tfsdk:"name"`
	Variables []ConfigurationGroupDeviceVariablesGroupsVariables `tfsdk:"variables"`
}

type ConfigurationGroupDeviceVariablesDevicesVariables struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type ConfigurationGroupDeviceVariablesGroupsVariables struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ConfigurationGroupDeviceVariables) getPath() string {
	return fmt.Sprintf("/v1/config-group/%v/device/variables/", url.QueryEscape(data.ConfigurationGroupId.ValueString()))
}

// End of section. //template:end getPath

func (data ConfigurationGroupDeviceVariables) toBody(ctx context.Context) string {
	body := ""
	if !data.Solution.IsNull() {
		body, _ = sjson.Set(body, "solution", data.Solution.ValueString())
	}
	if len(data.Devices) > 0 {
		body, _ = sjson.Set(body, "devices", []interface{}{})
		for _, item := range data.Devices {
			itemBody := ""
			if !item.DeviceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "device-id", item.DeviceId.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "variables", []interface{}{})
				for _, childItem := range item.Variables {
					itemChildBody := ""
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.Value.IsNull() {
						value := childItem.Value.ValueString()
						if i, err := strconv.Atoi(value); err == nil {
							itemChildBody, _ = sjson.Set(itemChildBody, "value", i)
						} else {
							itemChildBody, _ = sjson.Set(itemChildBody, "value", value)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "variables.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "devices.-1", itemBody)
		}
	}
	if len(data.Groups) > 0 {
		body, _ = sjson.Set(body, "groups", []interface{}{})
		for _, item := range data.Groups {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "group-variables", []interface{}{})
				for _, childItem := range item.Variables {
					itemChildBody := ""
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.Value.IsNull() {
						value := childItem.Value.ValueString()
						if i, err := strconv.Atoi(value); err == nil {
							itemChildBody, _ = sjson.Set(itemChildBody, "value", i)
						} else {
							itemChildBody, _ = sjson.Set(itemChildBody, "value", value)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "group-variables.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "groups.-1", itemBody)
		}
	}
	return body
}

func (data *ConfigurationGroupDeviceVariables) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("family"); value.Exists() {
		data.Solution = types.StringValue(value.String())
	} else {
		data.Solution = types.StringNull()
	}
	if value := res.Get("devices"); value.Exists() && len(value.Array()) > 0 {
		data.Devices = make([]ConfigurationGroupDeviceVariablesDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ConfigurationGroupDeviceVariablesDevices{}
			if cValue := v.Get("device-id"); cValue.Exists() {
				item.DeviceId = types.StringValue(cValue.String())
			} else {
				item.DeviceId = types.StringNull()
			}
			if cValue := v.Get("variables"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Variables = make([]ConfigurationGroupDeviceVariablesDevicesVariables, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ConfigurationGroupDeviceVariablesDevicesVariables{}
					if ccValue := cv.Get("name"); ccValue.Exists() {
						cItem.Name = types.StringValue(ccValue.String())
					} else {
						cItem.Name = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() {
						cItem.Value = types.StringValue(ccValue.String())
					} else {
						cItem.Value = types.StringNull()
						return true
					}
					item.Variables = append(item.Variables, cItem)
					return true
				})
			} else {
				if len(item.Variables) > 0 {
					item.Variables = []ConfigurationGroupDeviceVariablesDevicesVariables{}
				}
			}
			data.Devices = append(data.Devices, item)
			return true
		})
	} else {
		if len(data.Devices) > 0 {
			data.Devices = []ConfigurationGroupDeviceVariablesDevices{}
		}
	}
	if value := res.Get("groups"); value.Exists() && len(value.Array()) > 0 {
		data.Groups = make([]ConfigurationGroupDeviceVariablesGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ConfigurationGroupDeviceVariablesGroups{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("group-variables"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Variables = make([]ConfigurationGroupDeviceVariablesGroupsVariables, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ConfigurationGroupDeviceVariablesGroupsVariables{}
					if ccValue := cv.Get("name"); ccValue.Exists() {
						cItem.Name = types.StringValue(ccValue.String())
					} else {
						cItem.Name = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() {
						cItem.Value = types.StringValue(ccValue.String())
					} else {
						cItem.Value = types.StringNull()
						return true
					}
					item.Variables = append(item.Variables, cItem)
					return true
				})
			} else {
				if len(item.Variables) > 0 {
					item.Variables = []ConfigurationGroupDeviceVariablesGroupsVariables{}
				}
			}
			data.Groups = append(data.Groups, item)
			return true
		})
	} else {
		if len(data.Groups) > 0 {
			data.Groups = []ConfigurationGroupDeviceVariablesGroups{}
		}
	}
}

func (data *ConfigurationGroupDeviceVariables) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("family"); value.Exists() {
		data.Solution = types.StringValue(value.String())
	} else {
		data.Solution = types.StringNull()
	}
	for i := range data.Devices {
		keys := [...]string{"device-id"}
		keyValues := [...]string{data.Devices[i].DeviceId.ValueString()}

		var r gjson.Result
		res.Get("devices").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
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

		if cValue := r.Get("device-id"); cValue.Exists() {
			data.Devices[i].DeviceId = types.StringValue(cValue.String())
		} else {
			data.Devices[i].DeviceId = types.StringNull()
		}
		for ci := range data.Devices[i].Variables {
			keys := [...]string{"name"}
			keyValues := [...]string{data.Devices[i].Variables[ci].Name.ValueString()}

			var cr gjson.Result
			r.Get("variables").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() == keyValues[ik] {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)

			if ccValue := cr.Get("name"); ccValue.Exists() {
				data.Devices[i].Variables[ci].Name = types.StringValue(ccValue.String())
			} else {
				data.Devices[i].Variables[ci].Name = types.StringNull()
			}
			if ccValue := cr.Get("value"); ccValue.Exists() {
				data.Devices[i].Variables[ci].Value = types.StringValue(ccValue.String())
			} else {
				data.Devices[i].Variables[ci].Value = types.StringNull()
			}
		}
	}
	for i := range data.Groups {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Groups[i].Name.ValueString()}

		var r gjson.Result
		res.Get("groups").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
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

		if cValue := r.Get("name"); cValue.Exists() {
			data.Groups[i].Name = types.StringValue(cValue.String())
		} else {
			data.Groups[i].Name = types.StringNull()
		}
		for ci := range data.Groups[i].Variables {
			keys := [...]string{"name"}
			keyValues := [...]string{data.Groups[i].Variables[ci].Name.ValueString()}

			var cr gjson.Result
			r.Get("group-variables").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() == keyValues[ik] {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)

			if ccValue := cr.Get("name"); ccValue.Exists() {
				data.Groups[i].Variables[ci].Name = types.StringValue(ccValue.String())
			} else {
				data.Groups[i].Variables[ci].Name = types.StringNull()
			}
			if ccValue := cr.Get("value"); ccValue.Exists() {
				data.Groups[i].Variables[ci].Value = types.StringValue(ccValue.String())
			} else {
				data.Groups[i].Variables[ci].Value = types.StringNull()
			}
		}
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *ConfigurationGroupDeviceVariables) hasChanges(ctx context.Context, state *ConfigurationGroupDeviceVariables) bool {
	hasChanges := false
	if !data.ConfigurationGroupId.Equal(state.ConfigurationGroupId) {
		hasChanges = true
	}
	if !data.Solution.Equal(state.Solution) {
		hasChanges = true
	}
	if len(data.Devices) != len(state.Devices) {
		hasChanges = true
	} else {
		for i := range data.Devices {
			if !data.Devices[i].DeviceId.Equal(state.Devices[i].DeviceId) {
				hasChanges = true
			}
			if len(data.Devices[i].Variables) != len(state.Devices[i].Variables) {
				hasChanges = true
			} else {
				for ii := range data.Devices[i].Variables {
					if !data.Devices[i].Variables[ii].Name.Equal(state.Devices[i].Variables[ii].Name) {
						hasChanges = true
					}
					if !data.Devices[i].Variables[ii].Value.Equal(state.Devices[i].Variables[ii].Value) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.Groups) != len(state.Groups) {
		hasChanges = true
	} else {
		for i := range data.Groups {
			if !data.Groups[i].Name.Equal(state.Groups[i].Name) {
				hasChanges = true
			}
			if len(data.Groups[i].Variables) != len(state.Groups[i].Variables) {
				hasChanges = true
			} else {
				for ii := range data.Groups[i].Variables {
					if !data.Groups[i].Variables[ii].Name.Equal(state.Groups[i].Variables[ii].Name) {
						hasChanges = true
					}
					if !data.Groups[i].Variables[ii].Value.Equal(state.Groups[i].Variables[ii].Value) {
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

// End of section. //template:end updateVersions
