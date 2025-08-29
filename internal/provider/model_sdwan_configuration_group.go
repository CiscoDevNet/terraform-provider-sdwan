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
type ConfigurationGroup struct {
	Id                  types.String                        `tfsdk:"id"`
	Name                types.String                        `tfsdk:"name"`
	Description         types.String                        `tfsdk:"description"`
	Solution            types.String                        `tfsdk:"solution"`
	FeatureProfileIds   types.Set                           `tfsdk:"feature_profile_ids"`
	TopologyDevices     []ConfigurationGroupTopologyDevices `tfsdk:"topology_devices"`
	TopologySiteDevices types.Int64                         `tfsdk:"topology_site_devices"`
	Devices             []ConfigurationGroupDevices         `tfsdk:"devices"`
	FeatureVersions     types.List                          `tfsdk:"feature_versions"`
}

type ConfigurationGroupTopologyDevices struct {
	CriteriaAttribute   types.String                                           `tfsdk:"criteria_attribute"`
	CriteriaValue       types.String                                           `tfsdk:"criteria_value"`
	UnsupportedFeatures []ConfigurationGroupTopologyDevicesUnsupportedFeatures `tfsdk:"unsupported_features"`
}

type ConfigurationGroupDevices struct {
	Id        types.String                         `tfsdk:"id"`
	Deploy    types.Bool                           `tfsdk:"deploy"`
	Variables []ConfigurationGroupDevicesVariables `tfsdk:"variables"`
}

type ConfigurationGroupTopologyDevicesUnsupportedFeatures struct {
	ParcelType types.String `tfsdk:"parcel_type"`
	ParcelId   types.String `tfsdk:"parcel_id"`
}

type ConfigurationGroupDevicesVariables struct {
	Name      types.String `tfsdk:"name"`
	Value     types.String `tfsdk:"value"`
	ListValue types.List   `tfsdk:"list_value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ConfigurationGroup) getPath() string {
	return "/v1/config-group/"
}

// End of section. //template:end getPath

func (data ConfigurationGroup) toBodyConfigGroup(ctx context.Context) string {
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
	if true && len(data.TopologyDevices) > 0 {
		body, _ = sjson.Set(body, "topology.devices", []interface{}{})
		for _, item := range data.TopologyDevices {
			itemBody := ""
			if !item.CriteriaAttribute.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "criteria.attribute", item.CriteriaAttribute.ValueString())
			}
			if !item.CriteriaValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "criteria.value", item.CriteriaValue.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "unsupportedFeatures", []interface{}{})
				for _, childItem := range item.UnsupportedFeatures {
					itemChildBody := ""
					if !childItem.ParcelType.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "parcelType", childItem.ParcelType.ValueString())
					}
					if !childItem.ParcelId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "parcelId", childItem.ParcelId.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "unsupportedFeatures.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "topology.devices.-1", itemBody)
		}
	}
	if !data.TopologySiteDevices.IsNull() {
		body, _ = sjson.Set(body, "topology.siteDevices", data.TopologySiteDevices.ValueInt64())
	}
	return body
}

func (data ConfigurationGroup) toBodyConfigGroupDevices(ctx context.Context) string {
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

func (data ConfigurationGroup) toBodyConfigGroupDeviceVariables(ctx context.Context) string {
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
	// if true {
	// 	//body, _ = sjson.Set(body, "groups", []interface{}{})
	// 	for _, item := range data.DeviceGroups {
	// 		itemBody := ""
	// 		if !item.Name.IsNull() {
	// 			itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
	// 		}
	// 		if true {
	// 			itemBody, _ = sjson.Set(itemBody, "group-variables", []interface{}{})
	// 			for _, childItem := range item.Variables {
	// 				itemChildBody := ""
	// 				if !childItem.Name.IsNull() {
	// 					itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
	// 				}
	// 				if !childItem.Value.IsNull() {
	// 					if val, err := strconv.Atoi(childItem.Value.ValueString()); err == nil {
	// 						itemChildBody, _ = sjson.Set(itemChildBody, "value", val)
	// 					} else if val, err := strconv.ParseFloat(childItem.Value.ValueString(), 64); err == nil {
	// 						itemChildBody, _ = sjson.Set(itemChildBody, "value", val)
	// 					} else if val, err := strconv.ParseBool(childItem.Value.ValueString()); err == nil {
	// 						itemChildBody, _ = sjson.Set(itemChildBody, "value", val)
	// 					} else {
	// 						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
	// 					}
	// 				}
	// 				itemBody, _ = sjson.SetRaw(itemBody, "group-variables.-1", itemChildBody)
	// 			}
	// 		}
	// 		body, _ = sjson.SetRaw(body, "groups.-1", itemBody)
	// 	}
	// }
	return body
}

func (data *ConfigurationGroup) fromBodyConfigGroup(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("topology.devices"); value.Exists() && len(value.Array()) > 0 {
		data.TopologyDevices = make([]ConfigurationGroupTopologyDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ConfigurationGroupTopologyDevices{}
			if cValue := v.Get("criteria.attribute"); cValue.Exists() {
				item.CriteriaAttribute = types.StringValue(cValue.String())
			} else {
				item.CriteriaAttribute = types.StringNull()
			}
			if cValue := v.Get("criteria.value"); cValue.Exists() {
				item.CriteriaValue = types.StringValue(cValue.String())
			} else {
				item.CriteriaValue = types.StringNull()
			}
			if cValue := v.Get("unsupportedFeatures"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.UnsupportedFeatures = make([]ConfigurationGroupTopologyDevicesUnsupportedFeatures, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ConfigurationGroupTopologyDevicesUnsupportedFeatures{}
					if ccValue := cv.Get("parcelType"); ccValue.Exists() {
						cItem.ParcelType = types.StringValue(ccValue.String())
					} else {
						cItem.ParcelType = types.StringNull()
					}
					if ccValue := cv.Get("parcelId"); ccValue.Exists() {
						cItem.ParcelId = types.StringValue(ccValue.String())
					} else {
						cItem.ParcelId = types.StringNull()
					}
					item.UnsupportedFeatures = append(item.UnsupportedFeatures, cItem)
					return true
				})
			} else {
				if len(item.UnsupportedFeatures) > 0 {
					item.UnsupportedFeatures = []ConfigurationGroupTopologyDevicesUnsupportedFeatures{}
				}
			}
			data.TopologyDevices = append(data.TopologyDevices, item)
			return true
		})
	} else {
		if len(data.TopologyDevices) > 0 {
			data.TopologyDevices = []ConfigurationGroupTopologyDevices{}
		}
	}
	if value := res.Get("topology.siteDevices"); value.Exists() {
		data.TopologySiteDevices = types.Int64Value(value.Int())
	} else {
		data.TopologySiteDevices = types.Int64Null()
	}
}

func (data *ConfigurationGroup) fromBodyConfigGroupDevices(ctx context.Context, res gjson.Result) {
	original := *data
	if value := res.Get("devices"); value.Exists() && len(value.Array()) > 0 {
		data.Devices = make([]ConfigurationGroupDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ConfigurationGroupDevices{}
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
			data.Devices = []ConfigurationGroupDevices{}
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
				data.Devices = append([]ConfigurationGroupDevices{device}, data.Devices...)
			}
		}
	}
}

func (data *ConfigurationGroup) fromBodyConfigGroupDeviceVariables(ctx context.Context, res gjson.Result) {
	original := *data
	if value := res.Get("family"); value.Exists() {
		data.Solution = types.StringValue(value.String())
	} else {
		data.Solution = types.StringNull()
	}
	if value := res.Get("devices"); value.Exists() && len(value.Array()) > 0 {
		data.Devices = make([]ConfigurationGroupDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ConfigurationGroupDevices{}
			if cValue := v.Get("device-id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("variables"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Variables = make([]ConfigurationGroupDevicesVariables, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					// skip optional variables
					if !cv.Get("value").Exists() {
						return true
					}
					cItem := ConfigurationGroupDevicesVariables{}
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
			} else {
				if len(item.Variables) > 0 {
					item.Variables = []ConfigurationGroupDevicesVariables{}
				}
			}
			data.Devices = append(data.Devices, item)
			return true
		})
	} else {
		if len(data.Devices) > 0 {
			data.Devices = []ConfigurationGroupDevices{}
		}
	}
	// if value := res.Get("groups"); value.Exists() && len(value.Array()) > 0 {
	// 	data.DeviceGroups = make([]ConfigurationGroupDeviceGroups, 0)
	// 	value.ForEach(func(k, v gjson.Result) bool {
	// 		item := ConfigurationGroupDeviceGroups{}
	// 		if cValue := v.Get("name"); cValue.Exists() {
	// 			item.Name = types.StringValue(cValue.String())
	// 		} else {
	// 			item.Name = types.StringNull()
	// 		}
	// 		if cValue := v.Get("group-variables"); cValue.Exists() && len(cValue.Array()) > 0 {
	// 			item.Variables = make([]ConfigurationGroupDeviceGroupsVariables, 0)
	// 			cValue.ForEach(func(ck, cv gjson.Result) bool {
	// 				cItem := ConfigurationGroupDeviceGroupsVariables{}
	// 				if ccValue := cv.Get("name"); ccValue.Exists() {
	// 					cItem.Name = types.StringValue(ccValue.String())
	// 				} else {
	// 					cItem.Name = types.StringNull()
	// 				}
	// 				if ccValue := cv.Get("value"); ccValue.Exists() {
	// 					cItem.Value = types.StringValue(ccValue.String())
	// 				} else {
	// 					cItem.Value = types.StringNull()
	// 				}
	// 				item.Variables = append(item.Variables, cItem)
	// 				return true
	// 			})
	// 		} else {
	// 			if len(item.Variables) > 0 {
	// 				item.Variables = []ConfigurationGroupDeviceGroupsVariables{}
	// 			}
	// 		}
	// 		data.DeviceGroups = append(data.DeviceGroups, item)
	// 		return true
	// 	})
	// } else {
	// 	if len(data.DeviceGroups) > 0 {
	// 		data.DeviceGroups = []ConfigurationGroupDeviceGroups{}
	// 	}
	// }

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
				data.Devices = append([]ConfigurationGroupDevices{device}, data.Devices...)
			}
		}
	}
}

func (data *ConfigurationGroup) updateTfAttributes(ctx context.Context, state *ConfigurationGroup) {
	data.FeatureVersions = state.FeatureVersions
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

func (data ConfigurationGroup) hasConfigGroupDeviceVariables(ctx context.Context) bool {
	for _, device := range data.Devices {
		if len(device.Variables) > 0 {
			return true
		}
	}
	return false
}

func (data ConfigurationGroup) getUpdatedDevices(ctx context.Context, state *ConfigurationGroup) []string {
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

func (data ConfigurationGroup) hasFeatureVersionChanges(ctx context.Context, state *ConfigurationGroup) bool {
	var planValues, stateValues []string
	data.FeatureVersions.ElementsAs(ctx, &planValues, false)
	state.FeatureVersions.ElementsAs(ctx, &stateValues, false)
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

func (data ConfigurationGroup) processImport(ctx context.Context) {
	var elementsFeatureList []attr.Value
	var countFeatureProfile int = 0

	if !(data.FeatureProfileIds.IsNull() || data.FeatureProfileIds.IsUnknown()) {
		countFeatureProfile = len(data.FeatureProfileIds.Elements())
	}

	if countFeatureProfile > 0 {
		elementsFeatureList = make([]attr.Value, countFeatureProfile)
		for i := 0; i < countFeatureProfile; i++ {
			elementsFeatureList[i] = types.Int64Value(0)
		}
		data.FeatureVersions = types.ListValueMust(types.Int64Type, elementsFeatureList)
	} else {
		data.FeatureVersions = types.ListNull(types.StringType)
	}
}
