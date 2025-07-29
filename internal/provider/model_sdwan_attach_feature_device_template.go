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
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type AttachFeatureDeviceTemplate struct {
	Id      types.String                        `tfsdk:"id"`
	Version types.Int64                         `tfsdk:"version"`
	Devices []AttachFeatureDeviceTemplateDevice `tfsdk:"devices"`
}

type AttachFeatureDeviceTemplateDevice struct {
	Id        types.String `tfsdk:"id"`
	Variables types.Map    `tfsdk:"variables"`
}

func (data AttachFeatureDeviceTemplate) getVariables(ctx context.Context, client *sdwan.Client, edited bool, state *AttachFeatureDeviceTemplate) (map[string]map[string]string, error) {
	deviceVariables := make(map[string]map[string]string)
	var updatedDevices []string
	var attachedDevices []string

	// Get Updated Devices
	if state != nil {
		updatedDevices = data.getUpdatedDevices(ctx, state)
	} else {
		for _, item := range data.Devices {
			updatedDevices = append(updatedDevices, item.Id.ValueString())
		}
	}

	// Get Attached Devices
	if state != nil {
		for _, device := range data.Devices {
			found := false
			for _, stateDevice := range state.Devices {
				if device.Id.ValueString() == stateDevice.Id.ValueString() {
					found = true
					break
				}
			}
			if !found {
				attachedDevices = append(attachedDevices, device.Id.ValueString())
			}
		}
	}

	for _, item := range data.Devices {
		if helpers.Contains(updatedDevices, item.Id.ValueString()) || helpers.Contains(attachedDevices, item.Id.ValueString()) || edited {
			// Retrieve device variables
			body, _ := sjson.Set("", "templateId", data.Id.ValueString())
			body, _ = sjson.Set(body, "deviceIds", []string{item.Id.ValueString()})
			body, _ = sjson.Set(body, "isEdited", edited)
			body, _ = sjson.Set(body, "isMasterEdited", false)
			res, err := client.Post("/template/device/config/input", body)
			if err != nil {
				return nil, err
			}
			// Get variable mappings
			mappings := make(map[string]string)
			if value := res.Get("header.columns"); value.Exists() {
				value.ForEach(func(k, v gjson.Result) bool {
					if v.Get("editable").Bool() {
						title := v.Get("title").String()
						if strings.Contains(title, "(") {
							matches := regexp.MustCompile(`\((.*?)\)`).FindAllString(title, -1)
							varName := matches[len(matches)-1]
							if len(varName) > 0 {
								varName = varName[1 : len(varName)-1]
							}
							mappings[v.Get("property").String()] = varName
						} else {
							// handle factory default feature template variables
							property := v.Get("property").String()
							if property == "//system/host-name" {
								mappings[property] = "system_host_name"
							} else if property == "//system/system-ip" {
								mappings[property] = "system_system_ip"
							} else if property == "//system/site-id" {
								mappings[property] = "system_site_id"
							}
						}
					}
					return true
				})
			}
			// Get current device variables
			variables := make(map[string]string)
			if value := res.Get("data.0"); value.Exists() {
				value.ForEach(func(k, v gjson.Result) bool {
					variables[k.String()] = v.String()
					return true
				})
			}
			// Resolve variable names and insert template variable values
			var templateVariables map[string]string
			item.Variables.ElementsAs(ctx, &templateVariables, false)
			for k := range variables {
				if _, ok := templateVariables[k]; ok {
					variables[k] = templateVariables[k]
					continue
				}
				if templateVariableName, ok := mappings[k]; ok {
					if _, ok := templateVariables[templateVariableName]; ok {
						variables[k] = templateVariables[templateVariableName]
					}
				}
			}
			deviceVariables[item.Id.ValueString()] = variables
		}
	}

	return deviceVariables, nil
}

func (data AttachFeatureDeviceTemplate) getAttachedDevices(ctx context.Context, client *sdwan.Client) (map[string]string, error) {
	res, err := client.Get("/template/device/config/attached/" + data.Id.ValueString())
	if err != nil {
		return nil, err
	}
	devices := make(map[string]string)
	for _, device := range res.Get("data").Array() {
		devices[device.Get("uuid").String()] = device.Get("deviceIP").String()
	}
	return devices, nil
}

func (data AttachFeatureDeviceTemplate) detachDevices(ctx context.Context, client *sdwan.Client, plan *AttachFeatureDeviceTemplate, detachList []string) (gjson.Result, error) {
	devices, err := data.getAttachedDevices(ctx, client)
	if err != nil {
		return gjson.Result{}, err
	} else if len(devices) == 0 {
		return gjson.Result{}, nil
	}

	// Filter devices to only device attached with this resource
	attached_devices := make(map[string]string)
	for attached_id, attached_ip := range devices {
		for _, state_device := range data.Devices {
			if state_device.Id.ValueString() == attached_id {
				attached_devices[attached_id] = attached_ip
			}
		}
	}

	body, _ := sjson.Set("", "deviceType", "vedge")
	index := 0
	for k, v := range attached_devices {
		found := false
		for _, d := range plan.Devices {
			if d.Id.ValueString() == k {
				found = true
				break
			}
		}
		if len(detachList) != 0 && !slices.Contains(detachList, k) {
			found = true
		}
		if !found {
			body, _ = sjson.Set(body, "devices."+strconv.Itoa(index)+".deviceId", k)
			body, _ = sjson.Set(body, "devices."+strconv.Itoa(index)+".deviceIP", v)
			index++
		}
	}

	res, err := client.Post("/template/config/device/mode/cli", body)
	if err != nil {
		return gjson.Result{}, err
	}
	return res, nil
}

func (data AttachFeatureDeviceTemplate) toBody(ctx context.Context, client *sdwan.Client, edited bool, state *AttachFeatureDeviceTemplate) (string, error) {
	deviceVariables, err := data.getVariables(ctx, client, edited, state)
	if err != nil || len(deviceVariables) == 0 {
		return "", err
	}
	body, _ := sjson.Set("", "deviceTemplateList.0.templateId", data.Id.ValueString())
	body, _ = sjson.Set(body, "deviceTemplateList.0.isEdited", edited)
	body, _ = sjson.Set(body, "deviceTemplateList.0.isMasterEdited", false)
	body, _ = sjson.Set(body, "deviceTemplateList.0.isDraftDisabled", false)

	for _, v := range deviceVariables {
		itemBody := ""
		for name, value := range v {
			name = strings.Replace(name, ".", `\.`, -1)
			itemBody, _ = sjson.Set(itemBody, name, value)
		}
		body, _ = sjson.SetRaw(body, "deviceTemplateList.0.device.-1", itemBody)
	}

	return body, nil
}

func (data *AttachFeatureDeviceTemplate) readVariables(ctx context.Context, client *sdwan.Client) error {
	for i := range data.Devices {
		// Retrieve device variables
		body, _ := sjson.Set("", "templateId", data.Id.ValueString())
		body, _ = sjson.Set(body, "deviceIds", []string{data.Devices[i].Id.ValueString()})
		body, _ = sjson.Set(body, "isEdited", false)
		body, _ = sjson.Set(body, "isMasterEdited", false)
		res, err := client.Post("/template/device/config/input", body)
		if err != nil {
			return err
		}
		// Get variable mappings
		mappings := make(map[string]string)
		if value := res.Get("header.columns"); value.Exists() {
			value.ForEach(func(k, v gjson.Result) bool {
				if v.Get("editable").Bool() {
					title := v.Get("title").String()
					if strings.Contains(title, "(") {
						matches := regexp.MustCompile(`\((.*?)\)`).FindAllString(title, -1)
						varName := matches[len(matches)-1]
						if len(varName) > 0 {
							varName = varName[1 : len(varName)-1]
						}
						mappings[varName] = v.Get("property").String()
					} else if !v.Get("templateType").Exists() {
						// handle CLI template variables
						mappings[title] = title
					} else {
						// handle factory default feature template variables
						property := v.Get("property").String()
						if property == "//system/host-name" {
							mappings["system_host_name"] = property
						} else if property == "//system/system-ip" {
							mappings["system_system_ip"] = property
						} else if property == "//system/site-id" {
							mappings["system_site_id"] = property
						}
					}
				}
				return true
			})
		}
		// Get current device variables
		variables := make(map[string]string)
		if value := res.Get("data.0"); value.Exists() {
			value.ForEach(func(k, v gjson.Result) bool {
				variables[k.String()] = v.String()
				return true
			})
		}

		// Resolve variable names and insert template variable values
		newTemplateVariables := make(map[string]attr.Value)
		for k := range mappings {
			if _, ok := variables[mappings[k]]; ok {
				newTemplateVariables[k] = types.StringValue(variables[mappings[k]])
			}
		}
		data.Devices[i].Variables = types.MapValueMust(types.StringType, newTemplateVariables)
	}

	return nil
}

func (data *AttachFeatureDeviceTemplate) getUpdatedDevices(ctx context.Context, state *AttachFeatureDeviceTemplate) []string {
	updatedDevices := make([]string, 0)
	for _, device := range data.Devices {
		for _, stateDevice := range state.Devices {
			if device.Id.ValueString() == stateDevice.Id.ValueString() {
				for plan_k, plan_v := range device.Variables.Elements() {
					found := false
					for state_k, state_v := range stateDevice.Variables.Elements() {
						if plan_k == state_k {
							found = true
							if plan_v != state_v {
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

				for state_k, _ := range stateDevice.Variables.Elements() {
					found := false
					for plan_k, _ := range device.Variables.Elements() {
						if plan_k == state_k {
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

// Section below is generated&owned by "gen/generator.go". //template:begin processImport

// End of section. //template:end processImport
