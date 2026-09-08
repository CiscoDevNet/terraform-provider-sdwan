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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ServiceAppQoE struct {
	Id                           types.String                                `tfsdk:"id"`
	Version                      types.Int64                                 `tfsdk:"version"`
	Name                         types.String                                `tfsdk:"name"`
	Description                  types.String                                `tfsdk:"description"`
	FeatureProfileId             types.String                                `tfsdk:"feature_profile_id"`
	AppqoeDeviceRole             types.String                                `tfsdk:"appqoe_device_role"`
	VirtualApplications          []ServiceAppQoEVirtualApplications          `tfsdk:"virtual_applications"`
	ForwarderControllerGroups    []ServiceAppQoEForwarderControllerGroups    `tfsdk:"forwarder_controller_groups"`
	ForwarderServiceNodeGroups   []ServiceAppQoEForwarderServiceNodeGroups   `tfsdk:"forwarder_service_node_groups"`
	ForwarderServiceContexts     []ServiceAppQoEForwarderServiceContexts     `tfsdk:"forwarder_service_contexts"`
	CombinedControllerGroups     []ServiceAppQoECombinedControllerGroups     `tfsdk:"combined_controller_groups"`
	CombinedServiceNodeGroups    []ServiceAppQoECombinedServiceNodeGroups    `tfsdk:"combined_service_node_groups"`
	CombinedServiceContexts      []ServiceAppQoECombinedServiceContexts      `tfsdk:"combined_service_contexts"`
	ServiceNodeServiceNodeGroups []ServiceAppQoEServiceNodeServiceNodeGroups `tfsdk:"service_node_service_node_groups"`
}

type ServiceAppQoEVirtualApplications struct {
	ResourceProfile         types.String `tfsdk:"resource_profile"`
	ResourceProfileVariable types.String `tfsdk:"resource_profile_variable"`
}

type ServiceAppQoEForwarderControllerGroups struct {
	AppnavControllers []ServiceAppQoEForwarderControllerGroupsAppnavControllers `tfsdk:"appnav_controllers"`
}

type ServiceAppQoEForwarderServiceNodeGroups struct {
	Name         types.String                                          `tfsdk:"name"`
	ServiceNodes []ServiceAppQoEForwarderServiceNodeGroupsServiceNodes `tfsdk:"service_nodes"`
}

type ServiceAppQoEForwarderServiceContexts struct {
	AppnavControllerGroup types.String `tfsdk:"appnav_controller_group"`
	ServiceNodeGroup      types.String `tfsdk:"service_node_group"`
	Enable                types.Bool   `tfsdk:"enable"`
	Vpn                   types.Int64  `tfsdk:"vpn"`
	VpnVariable           types.String `tfsdk:"vpn_variable"`
}

type ServiceAppQoECombinedControllerGroups struct {
	GroupName         types.String                                             `tfsdk:"group_name"`
	AppnavControllers []ServiceAppQoECombinedControllerGroupsAppnavControllers `tfsdk:"appnav_controllers"`
}

type ServiceAppQoECombinedServiceNodeGroups struct {
	Name         types.String                                         `tfsdk:"name"`
	ServiceNodes []ServiceAppQoECombinedServiceNodeGroupsServiceNodes `tfsdk:"service_nodes"`
}

type ServiceAppQoECombinedServiceContexts struct {
	AppnavControllerGroup types.String `tfsdk:"appnav_controller_group"`
	ServiceNodeGroup      types.String `tfsdk:"service_node_group"`
	Enable                types.Bool   `tfsdk:"enable"`
	Vpn                   types.Int64  `tfsdk:"vpn"`
	VpnVariable           types.String `tfsdk:"vpn_variable"`
}

type ServiceAppQoEServiceNodeServiceNodeGroups struct {
	Name         types.String                                            `tfsdk:"name"`
	ServiceNodes []ServiceAppQoEServiceNodeServiceNodeGroupsServiceNodes `tfsdk:"service_nodes"`
}

type ServiceAppQoEForwarderControllerGroupsAppnavControllers struct {
	Address         types.String `tfsdk:"address"`
	AddressVariable types.String `tfsdk:"address_variable"`
	Vpn             types.Int64  `tfsdk:"vpn"`
}

type ServiceAppQoEForwarderServiceNodeGroupsServiceNodes struct {
	Address types.String `tfsdk:"address"`
}

type ServiceAppQoECombinedControllerGroupsAppnavControllers struct {
	Address types.String `tfsdk:"address"`
}

type ServiceAppQoECombinedServiceNodeGroupsServiceNodes struct {
	Address types.String `tfsdk:"address"`
}

type ServiceAppQoEServiceNodeServiceNodeGroupsServiceNodes struct {
	Address types.String `tfsdk:"address"`
	VpgIp   types.String `tfsdk:"vpg_ip"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceAppQoE) getModel() string {
	return "service_appqoe"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceAppQoE) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/appqoe", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceAppQoE) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.AppqoeDeviceRole.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"appqoeDeviceRole.optionType", "global")
			body, _ = sjson.Set(body, path+"appqoeDeviceRole.value", data.AppqoeDeviceRole.ValueString())
		}
	}
	if true && strings.Contains(data.AppqoeDeviceRole.ValueString(), "Dre") {
		body, _ = sjson.Set(body, path+"dreopt.optionType", "global")
		body, _ = sjson.Set(body, path+"dreopt.value", true)
	}
	if true && !(strings.Contains(data.AppqoeDeviceRole.ValueString(), "Dre")) {
		body, _ = sjson.Set(body, path+"dreopt.optionType", "global")
		body, _ = sjson.Set(body, path+"dreopt.value", false)
	}
	if true && strings.Contains(data.AppqoeDeviceRole.ValueString(), "Dre") {
		body, _ = sjson.Set(body, path+"virtualApplication", []interface{}{})
		for _, item := range data.VirtualApplications {
			itemBody := ""
			if true {
				itemBody, _ = sjson.Set(itemBody, "instanceId.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "instanceId.value", 1)
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "applicationType.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "applicationType.value", "dreopt")
			}

			if !item.ResourceProfileVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "resourceProfile.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "resourceProfile.value", item.ResourceProfileVariable.ValueString())
				}
			} else if item.ResourceProfile.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "resourceProfile.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "resourceProfile.value", "default")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "resourceProfile.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "resourceProfile.value", item.ResourceProfile.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"virtualApplication.-1", itemBody)
		}
	}
	if true && data.AppqoeDeviceRole.ValueString() == "forwarder" {
		body, _ = sjson.Set(body, path+"forwarder.appnavControllerGroup", []interface{}{})
		for _, item := range data.ForwarderControllerGroups {
			itemBody := ""
			if true {
				itemBody, _ = sjson.Set(itemBody, "groupName.optionType", "default")
				itemBody, _ = sjson.Set(itemBody, "groupName.value", "ACG-APPQOE")
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "appnavControllers", []interface{}{})
				for _, childItem := range item.AppnavControllers {
					itemChildBody := ""

					if !childItem.AddressVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.AddressVariable.ValueString())
						}
					} else if !childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
						}
					}
					if !childItem.Vpn.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "vpn.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "vpn.value", childItem.Vpn.ValueInt64())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "appnavControllers.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"forwarder.appnavControllerGroup.-1", itemBody)
		}
	}
	if true && data.AppqoeDeviceRole.ValueString() == "forwarder" {
		body, _ = sjson.Set(body, path+"forwarder.serviceNodeGroup", []interface{}{})
		for _, item := range data.ForwarderServiceNodeGroups {
			itemBody := ""
			if item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "name.value", "SNG-APPQOE")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "internal.optionType", "default")
				itemBody, _ = sjson.Set(itemBody, "internal.value", false)
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "serviceNode", []interface{}{})
				for _, childItem := range item.ServiceNodes {
					itemChildBody := ""
					if !childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "serviceNode.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"forwarder.serviceNodeGroup.-1", itemBody)
		}
	}
	if true && data.AppqoeDeviceRole.ValueString() == "forwarder" {
		body, _ = sjson.Set(body, path+"forwarder.serviceContext.appqoe", []interface{}{})
		for _, item := range data.ForwarderServiceContexts {
			itemBody := ""
			if true {
				itemBody, _ = sjson.Set(itemBody, "name.optionType", "default")
				itemBody, _ = sjson.Set(itemBody, "name.value", "/1")
			}
			if !item.AppnavControllerGroup.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "appnavControllerGroup.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "appnavControllerGroup.value", item.AppnavControllerGroup.ValueString())
				}
			}
			if !item.ServiceNodeGroup.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "serviceNodeGroup.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "serviceNodeGroup.value", item.ServiceNodeGroup.ValueString())
				}
			}
			if !item.Enable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "enable.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "enable.value", item.Enable.ValueBool())
				}
			}

			if !item.VpnVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.VpnVariable.ValueString())
				}
			} else if item.Vpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"forwarder.serviceContext.appqoe.-1", itemBody)
		}
	}
	if true && strings.Contains(data.AppqoeDeviceRole.ValueString(), "forwarderAndServiceNode") {
		body, _ = sjson.Set(body, path+"forwarderAndServiceNode.appnavControllerGroup", []interface{}{})
		for _, item := range data.CombinedControllerGroups {
			itemBody := ""
			if item.GroupName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "groupName.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "groupName.value", "ACG-APPQOE")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "groupName.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "groupName.value", item.GroupName.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "appnavControllers", []interface{}{})
				for _, childItem := range item.AppnavControllers {
					itemChildBody := ""
					if childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", "192.168.2.1")
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "appnavControllers.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"forwarderAndServiceNode.appnavControllerGroup.-1", itemBody)
		}
	}
	if true && strings.Contains(data.AppqoeDeviceRole.ValueString(), "forwarderAndServiceNode") {
		body, _ = sjson.Set(body, path+"forwarderAndServiceNode.serviceNodeGroup", []interface{}{})
		for _, item := range data.CombinedServiceNodeGroups {
			itemBody := ""
			if item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "name.value", "SNG-APPQOE")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "internal.optionType", "default")
				itemBody, _ = sjson.Set(itemBody, "internal.value", true)
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "serviceNode", []interface{}{})
				for _, childItem := range item.ServiceNodes {
					itemChildBody := ""
					if childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", "192.168.2.2")
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "serviceNode.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"forwarderAndServiceNode.serviceNodeGroup.-1", itemBody)
		}
	}
	if true && strings.Contains(data.AppqoeDeviceRole.ValueString(), "forwarderAndServiceNode") {
		body, _ = sjson.Set(body, path+"forwarderAndServiceNode.serviceContext.appqoe", []interface{}{})
		for _, item := range data.CombinedServiceContexts {
			itemBody := ""
			if true {
				itemBody, _ = sjson.Set(itemBody, "name.optionType", "default")
				itemBody, _ = sjson.Set(itemBody, "name.value", "/1")
			}
			if !item.AppnavControllerGroup.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "appnavControllerGroup.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "appnavControllerGroup.value", item.AppnavControllerGroup.ValueString())
				}
			}
			if !item.ServiceNodeGroup.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "serviceNodeGroup.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "serviceNodeGroup.value", item.ServiceNodeGroup.ValueString())
				}
			}
			if !item.Enable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "enable.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "enable.value", item.Enable.ValueBool())
				}
			}

			if !item.VpnVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.VpnVariable.ValueString())
				}
			} else if item.Vpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"forwarderAndServiceNode.serviceContext.appqoe.-1", itemBody)
		}
	}
	if true && strings.Contains(data.AppqoeDeviceRole.ValueString(), "serviceNode") {
		body, _ = sjson.Set(body, path+"serviceNode.serviceNodeGroup", []interface{}{})
		for _, item := range data.ServiceNodeServiceNodeGroups {
			itemBody := ""
			if item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "name.value", "SNG-APPQOE")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "externalNode.optionType", "default")
				itemBody, _ = sjson.Set(itemBody, "externalNode.value", true)
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "serviceNode", []interface{}{})
				for _, childItem := range item.ServiceNodes {
					itemChildBody := ""
					if childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", "192.168.2.2")
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
						}
					}
					if childItem.VpgIp.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "vpgIp.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "vpgIp.value", "192.168.2.1/24")
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "vpgIp.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "vpgIp.value", childItem.VpgIp.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "serviceNode.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"serviceNode.serviceNodeGroup.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceAppQoE) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.AppqoeDeviceRole = types.StringNull()

	if t := res.Get(path + "appqoeDeviceRole.optionType"); t.Exists() {
		va := res.Get(path + "appqoeDeviceRole.value")
		if t.String() == "global" {
			data.AppqoeDeviceRole = types.StringValue(va.String())
		}
	}
	oldVirtualApplications := data.VirtualApplications
	if value := res.Get(path + "virtualApplication"); value.Exists() && len(value.Array()) > 0 {
		data.VirtualApplications = make([]ServiceAppQoEVirtualApplications, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceAppQoEVirtualApplications{}
			item.ResourceProfile = types.StringNull()
			item.ResourceProfileVariable = types.StringNull()
			if t := v.Get("resourceProfile.optionType"); t.Exists() {
				va := v.Get("resourceProfile.value")
				if t.String() == "variable" {
					item.ResourceProfileVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.ResourceProfile = types.StringValue(va.String())
				}
			}
			data.VirtualApplications = append(data.VirtualApplications, item)
			return true
		})
	} else {
		data.VirtualApplications = nil
	}
	if !fullRead && data.VirtualApplications != nil {
		resultVirtualApplications := make([]ServiceAppQoEVirtualApplications, 0, len(data.VirtualApplications))
		matchedVirtualApplications := make([]bool, len(data.VirtualApplications))
		for _, oldItem := range oldVirtualApplications {
			for ni := range data.VirtualApplications {
				if matchedVirtualApplications[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.ResourceProfileVariable.ValueString() != "" || data.VirtualApplications[ni].ResourceProfileVariable.ValueString() != "") {
					if oldItem.ResourceProfileVariable.ValueString() != data.VirtualApplications[ni].ResourceProfileVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.ResourceProfile.ValueString() != data.VirtualApplications[ni].ResourceProfile.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedVirtualApplications[ni] = true
					resultVirtualApplications = append(resultVirtualApplications, data.VirtualApplications[ni])
					break
				}
			}
		}
		for ni := range data.VirtualApplications {
			if !matchedVirtualApplications[ni] {
				resultVirtualApplications = append(resultVirtualApplications, data.VirtualApplications[ni])
			}
		}
		data.VirtualApplications = resultVirtualApplications
	}
	oldForwarderControllerGroups := data.ForwarderControllerGroups
	if value := res.Get(path + "forwarder.appnavControllerGroup"); value.Exists() && len(value.Array()) > 0 {
		data.ForwarderControllerGroups = make([]ServiceAppQoEForwarderControllerGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceAppQoEForwarderControllerGroups{}
			if cValue := v.Get("appnavControllers"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.AppnavControllers = make([]ServiceAppQoEForwarderControllerGroupsAppnavControllers, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceAppQoEForwarderControllerGroupsAppnavControllers{}
					cItem.Address = types.StringNull()
					cItem.AddressVariable = types.StringNull()
					if t := cv.Get("address.optionType"); t.Exists() {
						va := cv.Get("address.value")
						if t.String() == "variable" {
							cItem.AddressVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					cItem.Vpn = types.Int64Null()

					if t := cv.Get("vpn.optionType"); t.Exists() {
						va := cv.Get("vpn.value")
						if t.String() == "global" {
							cItem.Vpn = types.Int64Value(va.Int())
						}
					}
					item.AppnavControllers = append(item.AppnavControllers, cItem)
					return true
				})
			}
			data.ForwarderControllerGroups = append(data.ForwarderControllerGroups, item)
			return true
		})
	} else {
		data.ForwarderControllerGroups = nil
	}
	if !fullRead && data.ForwarderControllerGroups != nil {
		resultForwarderControllerGroups := make([]ServiceAppQoEForwarderControllerGroups, 0, len(data.ForwarderControllerGroups))
		matchedForwarderControllerGroups := make([]bool, len(data.ForwarderControllerGroups))
		for _, oldItem := range oldForwarderControllerGroups {
			for ni := range data.ForwarderControllerGroups {
				if matchedForwarderControllerGroups[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					matchedForwarderControllerGroups[ni] = true
					if data.ForwarderControllerGroups[ni].AppnavControllers != nil {
						resultC := make([]ServiceAppQoEForwarderControllerGroupsAppnavControllers, 0, len(data.ForwarderControllerGroups[ni].AppnavControllers))
						matchedC := make([]bool, len(data.ForwarderControllerGroups[ni].AppnavControllers))
						for _, oldCItem := range oldItem.AppnavControllers {
							for nci := range data.ForwarderControllerGroups[ni].AppnavControllers {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC && (oldCItem.AddressVariable.ValueString() != "" || data.ForwarderControllerGroups[ni].AppnavControllers[nci].AddressVariable.ValueString() != "") {
									if oldCItem.AddressVariable.ValueString() != data.ForwarderControllerGroups[ni].AppnavControllers[nci].AddressVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.Address.ValueString() != data.ForwarderControllerGroups[ni].AppnavControllers[nci].Address.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.ForwarderControllerGroups[ni].AppnavControllers[nci])
									break
								}
							}
						}
						for nci := range data.ForwarderControllerGroups[ni].AppnavControllers {
							if !matchedC[nci] {
								resultC = append(resultC, data.ForwarderControllerGroups[ni].AppnavControllers[nci])
							}
						}
						data.ForwarderControllerGroups[ni].AppnavControllers = resultC
					}
					resultForwarderControllerGroups = append(resultForwarderControllerGroups, data.ForwarderControllerGroups[ni])
					break
				}
			}
		}
		for ni := range data.ForwarderControllerGroups {
			if !matchedForwarderControllerGroups[ni] {
				resultForwarderControllerGroups = append(resultForwarderControllerGroups, data.ForwarderControllerGroups[ni])
			}
		}
		data.ForwarderControllerGroups = resultForwarderControllerGroups
	}
	oldForwarderServiceNodeGroups := data.ForwarderServiceNodeGroups
	if value := res.Get(path + "forwarder.serviceNodeGroup"); value.Exists() && len(value.Array()) > 0 {
		data.ForwarderServiceNodeGroups = make([]ServiceAppQoEForwarderServiceNodeGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceAppQoEForwarderServiceNodeGroups{}
			item.Name = types.StringNull()

			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("serviceNode"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ServiceNodes = make([]ServiceAppQoEForwarderServiceNodeGroupsServiceNodes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceAppQoEForwarderServiceNodeGroupsServiceNodes{}
					cItem.Address = types.StringNull()

					if t := cv.Get("address.optionType"); t.Exists() {
						va := cv.Get("address.value")
						if t.String() == "global" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					item.ServiceNodes = append(item.ServiceNodes, cItem)
					return true
				})
			}
			data.ForwarderServiceNodeGroups = append(data.ForwarderServiceNodeGroups, item)
			return true
		})
	} else {
		data.ForwarderServiceNodeGroups = nil
	}
	if !fullRead && data.ForwarderServiceNodeGroups != nil {
		resultForwarderServiceNodeGroups := make([]ServiceAppQoEForwarderServiceNodeGroups, 0, len(data.ForwarderServiceNodeGroups))
		matchedForwarderServiceNodeGroups := make([]bool, len(data.ForwarderServiceNodeGroups))
		for _, oldItem := range oldForwarderServiceNodeGroups {
			for ni := range data.ForwarderServiceNodeGroups {
				if matchedForwarderServiceNodeGroups[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.Name.ValueString() != data.ForwarderServiceNodeGroups[ni].Name.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedForwarderServiceNodeGroups[ni] = true
					if data.ForwarderServiceNodeGroups[ni].ServiceNodes != nil {
						resultC := make([]ServiceAppQoEForwarderServiceNodeGroupsServiceNodes, 0, len(data.ForwarderServiceNodeGroups[ni].ServiceNodes))
						matchedC := make([]bool, len(data.ForwarderServiceNodeGroups[ni].ServiceNodes))
						for _, oldCItem := range oldItem.ServiceNodes {
							for nci := range data.ForwarderServiceNodeGroups[ni].ServiceNodes {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if oldCItem.Address.ValueString() != data.ForwarderServiceNodeGroups[ni].ServiceNodes[nci].Address.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.ForwarderServiceNodeGroups[ni].ServiceNodes[nci])
									break
								}
							}
						}
						for nci := range data.ForwarderServiceNodeGroups[ni].ServiceNodes {
							if !matchedC[nci] {
								resultC = append(resultC, data.ForwarderServiceNodeGroups[ni].ServiceNodes[nci])
							}
						}
						data.ForwarderServiceNodeGroups[ni].ServiceNodes = resultC
					}
					resultForwarderServiceNodeGroups = append(resultForwarderServiceNodeGroups, data.ForwarderServiceNodeGroups[ni])
					break
				}
			}
		}
		for ni := range data.ForwarderServiceNodeGroups {
			if !matchedForwarderServiceNodeGroups[ni] {
				resultForwarderServiceNodeGroups = append(resultForwarderServiceNodeGroups, data.ForwarderServiceNodeGroups[ni])
			}
		}
		data.ForwarderServiceNodeGroups = resultForwarderServiceNodeGroups
	}
	oldForwarderServiceContexts := data.ForwarderServiceContexts
	if value := res.Get(path + "forwarder.serviceContext.appqoe"); value.Exists() && len(value.Array()) > 0 {
		data.ForwarderServiceContexts = make([]ServiceAppQoEForwarderServiceContexts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceAppQoEForwarderServiceContexts{}
			item.AppnavControllerGroup = types.StringNull()

			if t := v.Get("appnavControllerGroup.optionType"); t.Exists() {
				va := v.Get("appnavControllerGroup.value")
				if t.String() == "global" {
					item.AppnavControllerGroup = types.StringValue(va.String())
				}
			}
			item.ServiceNodeGroup = types.StringNull()

			if t := v.Get("serviceNodeGroup.optionType"); t.Exists() {
				va := v.Get("serviceNodeGroup.value")
				if t.String() == "global" {
					item.ServiceNodeGroup = types.StringValue(va.String())
				}
			}
			item.Enable = types.BoolNull()

			if t := v.Get("enable.optionType"); t.Exists() {
				va := v.Get("enable.value")
				if t.String() == "global" {
					item.Enable = types.BoolValue(va.Bool())
				}
			}
			item.Vpn = types.Int64Null()
			item.VpnVariable = types.StringNull()
			if t := v.Get("vpn.optionType"); t.Exists() {
				va := v.Get("vpn.value")
				if t.String() == "variable" {
					item.VpnVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Vpn = types.Int64Value(va.Int())
				}
			}
			data.ForwarderServiceContexts = append(data.ForwarderServiceContexts, item)
			return true
		})
	} else {
		data.ForwarderServiceContexts = nil
	}
	if !fullRead && data.ForwarderServiceContexts != nil {
		resultForwarderServiceContexts := make([]ServiceAppQoEForwarderServiceContexts, 0, len(data.ForwarderServiceContexts))
		matchedForwarderServiceContexts := make([]bool, len(data.ForwarderServiceContexts))
		for _, oldItem := range oldForwarderServiceContexts {
			for ni := range data.ForwarderServiceContexts {
				if matchedForwarderServiceContexts[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.AppnavControllerGroup.ValueString() != data.ForwarderServiceContexts[ni].AppnavControllerGroup.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedForwarderServiceContexts[ni] = true
					resultForwarderServiceContexts = append(resultForwarderServiceContexts, data.ForwarderServiceContexts[ni])
					break
				}
			}
		}
		for ni := range data.ForwarderServiceContexts {
			if !matchedForwarderServiceContexts[ni] {
				resultForwarderServiceContexts = append(resultForwarderServiceContexts, data.ForwarderServiceContexts[ni])
			}
		}
		data.ForwarderServiceContexts = resultForwarderServiceContexts
	}
	oldCombinedControllerGroups := data.CombinedControllerGroups
	if value := res.Get(path + "forwarderAndServiceNode.appnavControllerGroup"); value.Exists() && len(value.Array()) > 0 {
		data.CombinedControllerGroups = make([]ServiceAppQoECombinedControllerGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceAppQoECombinedControllerGroups{}
			item.GroupName = types.StringNull()

			if t := v.Get("groupName.optionType"); t.Exists() {
				va := v.Get("groupName.value")
				if t.String() == "global" || t.String() == "default" {
					item.GroupName = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("appnavControllers"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.AppnavControllers = make([]ServiceAppQoECombinedControllerGroupsAppnavControllers, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceAppQoECombinedControllerGroupsAppnavControllers{}
					cItem.Address = types.StringNull()

					if t := cv.Get("address.optionType"); t.Exists() {
						va := cv.Get("address.value")
						if t.String() == "global" || t.String() == "default" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					item.AppnavControllers = append(item.AppnavControllers, cItem)
					return true
				})
			}
			data.CombinedControllerGroups = append(data.CombinedControllerGroups, item)
			return true
		})
	} else {
		data.CombinedControllerGroups = nil
	}
	if !fullRead && data.CombinedControllerGroups != nil {
		resultCombinedControllerGroups := make([]ServiceAppQoECombinedControllerGroups, 0, len(data.CombinedControllerGroups))
		matchedCombinedControllerGroups := make([]bool, len(data.CombinedControllerGroups))
		for _, oldItem := range oldCombinedControllerGroups {
			for ni := range data.CombinedControllerGroups {
				if matchedCombinedControllerGroups[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.GroupName.ValueString() != data.CombinedControllerGroups[ni].GroupName.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedCombinedControllerGroups[ni] = true
					if data.CombinedControllerGroups[ni].AppnavControllers != nil {
						resultC := make([]ServiceAppQoECombinedControllerGroupsAppnavControllers, 0, len(data.CombinedControllerGroups[ni].AppnavControllers))
						matchedC := make([]bool, len(data.CombinedControllerGroups[ni].AppnavControllers))
						for _, oldCItem := range oldItem.AppnavControllers {
							for nci := range data.CombinedControllerGroups[ni].AppnavControllers {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if oldCItem.Address.ValueString() != data.CombinedControllerGroups[ni].AppnavControllers[nci].Address.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.CombinedControllerGroups[ni].AppnavControllers[nci])
									break
								}
							}
						}
						for nci := range data.CombinedControllerGroups[ni].AppnavControllers {
							if !matchedC[nci] {
								resultC = append(resultC, data.CombinedControllerGroups[ni].AppnavControllers[nci])
							}
						}
						data.CombinedControllerGroups[ni].AppnavControllers = resultC
					}
					resultCombinedControllerGroups = append(resultCombinedControllerGroups, data.CombinedControllerGroups[ni])
					break
				}
			}
		}
		for ni := range data.CombinedControllerGroups {
			if !matchedCombinedControllerGroups[ni] {
				resultCombinedControllerGroups = append(resultCombinedControllerGroups, data.CombinedControllerGroups[ni])
			}
		}
		data.CombinedControllerGroups = resultCombinedControllerGroups
	}
	oldCombinedServiceNodeGroups := data.CombinedServiceNodeGroups
	if value := res.Get(path + "forwarderAndServiceNode.serviceNodeGroup"); value.Exists() && len(value.Array()) > 0 {
		data.CombinedServiceNodeGroups = make([]ServiceAppQoECombinedServiceNodeGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceAppQoECombinedServiceNodeGroups{}
			item.Name = types.StringNull()

			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "global" || t.String() == "default" {
					item.Name = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("serviceNode"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ServiceNodes = make([]ServiceAppQoECombinedServiceNodeGroupsServiceNodes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceAppQoECombinedServiceNodeGroupsServiceNodes{}
					cItem.Address = types.StringNull()

					if t := cv.Get("address.optionType"); t.Exists() {
						va := cv.Get("address.value")
						if t.String() == "global" || t.String() == "default" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					item.ServiceNodes = append(item.ServiceNodes, cItem)
					return true
				})
			}
			data.CombinedServiceNodeGroups = append(data.CombinedServiceNodeGroups, item)
			return true
		})
	} else {
		data.CombinedServiceNodeGroups = nil
	}
	if !fullRead && data.CombinedServiceNodeGroups != nil {
		resultCombinedServiceNodeGroups := make([]ServiceAppQoECombinedServiceNodeGroups, 0, len(data.CombinedServiceNodeGroups))
		matchedCombinedServiceNodeGroups := make([]bool, len(data.CombinedServiceNodeGroups))
		for _, oldItem := range oldCombinedServiceNodeGroups {
			for ni := range data.CombinedServiceNodeGroups {
				if matchedCombinedServiceNodeGroups[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.Name.ValueString() != data.CombinedServiceNodeGroups[ni].Name.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedCombinedServiceNodeGroups[ni] = true
					if data.CombinedServiceNodeGroups[ni].ServiceNodes != nil {
						resultC := make([]ServiceAppQoECombinedServiceNodeGroupsServiceNodes, 0, len(data.CombinedServiceNodeGroups[ni].ServiceNodes))
						matchedC := make([]bool, len(data.CombinedServiceNodeGroups[ni].ServiceNodes))
						for _, oldCItem := range oldItem.ServiceNodes {
							for nci := range data.CombinedServiceNodeGroups[ni].ServiceNodes {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if oldCItem.Address.ValueString() != data.CombinedServiceNodeGroups[ni].ServiceNodes[nci].Address.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.CombinedServiceNodeGroups[ni].ServiceNodes[nci])
									break
								}
							}
						}
						for nci := range data.CombinedServiceNodeGroups[ni].ServiceNodes {
							if !matchedC[nci] {
								resultC = append(resultC, data.CombinedServiceNodeGroups[ni].ServiceNodes[nci])
							}
						}
						data.CombinedServiceNodeGroups[ni].ServiceNodes = resultC
					}
					resultCombinedServiceNodeGroups = append(resultCombinedServiceNodeGroups, data.CombinedServiceNodeGroups[ni])
					break
				}
			}
		}
		for ni := range data.CombinedServiceNodeGroups {
			if !matchedCombinedServiceNodeGroups[ni] {
				resultCombinedServiceNodeGroups = append(resultCombinedServiceNodeGroups, data.CombinedServiceNodeGroups[ni])
			}
		}
		data.CombinedServiceNodeGroups = resultCombinedServiceNodeGroups
	}
	oldCombinedServiceContexts := data.CombinedServiceContexts
	if value := res.Get(path + "forwarderAndServiceNode.serviceContext.appqoe"); value.Exists() && len(value.Array()) > 0 {
		data.CombinedServiceContexts = make([]ServiceAppQoECombinedServiceContexts, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceAppQoECombinedServiceContexts{}
			item.AppnavControllerGroup = types.StringNull()

			if t := v.Get("appnavControllerGroup.optionType"); t.Exists() {
				va := v.Get("appnavControllerGroup.value")
				if t.String() == "global" {
					item.AppnavControllerGroup = types.StringValue(va.String())
				}
			}
			item.ServiceNodeGroup = types.StringNull()

			if t := v.Get("serviceNodeGroup.optionType"); t.Exists() {
				va := v.Get("serviceNodeGroup.value")
				if t.String() == "global" {
					item.ServiceNodeGroup = types.StringValue(va.String())
				}
			}
			item.Enable = types.BoolNull()

			if t := v.Get("enable.optionType"); t.Exists() {
				va := v.Get("enable.value")
				if t.String() == "global" {
					item.Enable = types.BoolValue(va.Bool())
				}
			}
			item.Vpn = types.Int64Null()
			item.VpnVariable = types.StringNull()
			if t := v.Get("vpn.optionType"); t.Exists() {
				va := v.Get("vpn.value")
				if t.String() == "variable" {
					item.VpnVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Vpn = types.Int64Value(va.Int())
				}
			}
			data.CombinedServiceContexts = append(data.CombinedServiceContexts, item)
			return true
		})
	} else {
		data.CombinedServiceContexts = nil
	}
	if !fullRead && data.CombinedServiceContexts != nil {
		resultCombinedServiceContexts := make([]ServiceAppQoECombinedServiceContexts, 0, len(data.CombinedServiceContexts))
		matchedCombinedServiceContexts := make([]bool, len(data.CombinedServiceContexts))
		for _, oldItem := range oldCombinedServiceContexts {
			for ni := range data.CombinedServiceContexts {
				if matchedCombinedServiceContexts[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.AppnavControllerGroup.ValueString() != data.CombinedServiceContexts[ni].AppnavControllerGroup.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedCombinedServiceContexts[ni] = true
					resultCombinedServiceContexts = append(resultCombinedServiceContexts, data.CombinedServiceContexts[ni])
					break
				}
			}
		}
		for ni := range data.CombinedServiceContexts {
			if !matchedCombinedServiceContexts[ni] {
				resultCombinedServiceContexts = append(resultCombinedServiceContexts, data.CombinedServiceContexts[ni])
			}
		}
		data.CombinedServiceContexts = resultCombinedServiceContexts
	}
	oldServiceNodeServiceNodeGroups := data.ServiceNodeServiceNodeGroups
	if value := res.Get(path + "serviceNode.serviceNodeGroup"); value.Exists() && len(value.Array()) > 0 {
		data.ServiceNodeServiceNodeGroups = make([]ServiceAppQoEServiceNodeServiceNodeGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceAppQoEServiceNodeServiceNodeGroups{}
			item.Name = types.StringNull()

			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "global" || t.String() == "default" {
					item.Name = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("serviceNode"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ServiceNodes = make([]ServiceAppQoEServiceNodeServiceNodeGroupsServiceNodes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceAppQoEServiceNodeServiceNodeGroupsServiceNodes{}
					cItem.Address = types.StringNull()

					if t := cv.Get("address.optionType"); t.Exists() {
						va := cv.Get("address.value")
						if t.String() == "global" || t.String() == "default" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					cItem.VpgIp = types.StringNull()

					if t := cv.Get("vpgIp.optionType"); t.Exists() {
						va := cv.Get("vpgIp.value")
						if t.String() == "global" || t.String() == "default" {
							cItem.VpgIp = types.StringValue(va.String())
						}
					}
					item.ServiceNodes = append(item.ServiceNodes, cItem)
					return true
				})
			}
			data.ServiceNodeServiceNodeGroups = append(data.ServiceNodeServiceNodeGroups, item)
			return true
		})
	} else {
		data.ServiceNodeServiceNodeGroups = nil
	}
	if !fullRead && data.ServiceNodeServiceNodeGroups != nil {
		resultServiceNodeServiceNodeGroups := make([]ServiceAppQoEServiceNodeServiceNodeGroups, 0, len(data.ServiceNodeServiceNodeGroups))
		matchedServiceNodeServiceNodeGroups := make([]bool, len(data.ServiceNodeServiceNodeGroups))
		for _, oldItem := range oldServiceNodeServiceNodeGroups {
			for ni := range data.ServiceNodeServiceNodeGroups {
				if matchedServiceNodeServiceNodeGroups[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.Name.ValueString() != data.ServiceNodeServiceNodeGroups[ni].Name.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedServiceNodeServiceNodeGroups[ni] = true
					if data.ServiceNodeServiceNodeGroups[ni].ServiceNodes != nil {
						resultC := make([]ServiceAppQoEServiceNodeServiceNodeGroupsServiceNodes, 0, len(data.ServiceNodeServiceNodeGroups[ni].ServiceNodes))
						matchedC := make([]bool, len(data.ServiceNodeServiceNodeGroups[ni].ServiceNodes))
						for _, oldCItem := range oldItem.ServiceNodes {
							for nci := range data.ServiceNodeServiceNodeGroups[ni].ServiceNodes {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if oldCItem.Address.ValueString() != data.ServiceNodeServiceNodeGroups[ni].ServiceNodes[nci].Address.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.ServiceNodeServiceNodeGroups[ni].ServiceNodes[nci])
									break
								}
							}
						}
						for nci := range data.ServiceNodeServiceNodeGroups[ni].ServiceNodes {
							if !matchedC[nci] {
								resultC = append(resultC, data.ServiceNodeServiceNodeGroups[ni].ServiceNodes[nci])
							}
						}
						data.ServiceNodeServiceNodeGroups[ni].ServiceNodes = resultC
					}
					resultServiceNodeServiceNodeGroups = append(resultServiceNodeServiceNodeGroups, data.ServiceNodeServiceNodeGroups[ni])
					break
				}
			}
		}
		for ni := range data.ServiceNodeServiceNodeGroups {
			if !matchedServiceNodeServiceNodeGroups[ni] {
				resultServiceNodeServiceNodeGroups = append(resultServiceNodeServiceNodeGroups, data.ServiceNodeServiceNodeGroups[ni])
			}
		}
		data.ServiceNodeServiceNodeGroups = resultServiceNodeServiceNodeGroups
	}
}

// End of section. //template:end fromBody
