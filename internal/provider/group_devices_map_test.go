// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
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
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

// devices is a map in the Terraform shape and a list sorted by ID in the working model.
func TestConfigurationGroupDevicesConfigToModel(t *testing.T) {
	groups := types.ListValueMust(types.StringType, []attr.Value{types.StringValue("1"), types.StringValue("2")})
	vars := map[string]ConfigurationGroupDevicesVariableConfig{
		"system_ip":                {Value: types.StringValue("10.0.0.1"), ListValue: types.ListNull(types.StringType)},
		"host_name":                {Value: types.StringValue("r1"), ListValue: types.ListNull(types.StringType)},
		"system_controller_groups": {Value: types.StringNull(), ListValue: groups},
	}
	config := ConfigurationGroupDevicesConfig{
		ConfigurationGroupId: types.StringValue("cg1"),
		Devices: map[string]ConfigurationGroupDevicesDeviceConfig{
			"C8K-C": {Deploy: types.BoolValue(true)},
			"C8K-A": {Deploy: types.BoolValue(false), TopologyLabel: types.StringValue("EdgeDevice_1"), Variables: vars},
			"C8K-B": {Deploy: types.BoolValue(true)},
		},
	}
	data := config.toModel()
	var ids []string
	for _, d := range data.Devices {
		ids = append(ids, d.Id.ValueString())
	}
	if !reflect.DeepEqual(ids, []string{"C8K-A", "C8K-B", "C8K-C"}) {
		t.Fatalf("devices not sorted by ID: %v", ids)
	}
	got := data.Devices[0]
	if got.TopologyLabel.ValueString() != "EdgeDevice_1" || got.Deploy.ValueBool() {
		t.Errorf("C8K-A fields not carried over: %+v", got)
	}
	var names []string
	for _, v := range got.Variables {
		names = append(names, v.Name.ValueString())
	}
	if !reflect.DeepEqual(names, []string{"host_name", "system_controller_groups", "system_ip"}) {
		t.Errorf("variables not sorted by name: %v", names)
	}
	if !got.Variables[1].ListValue.Equal(groups) || !got.Variables[1].Value.IsNull() {
		t.Errorf("list variable not carried over: %+v", got.Variables[1])
	}
	if data.Devices[1].Variables != nil {
		t.Errorf("null variables of C8K-B became %v", data.Devices[1].Variables)
	}
	if back := data.toConfig(); !reflect.DeepEqual(back, config) {
		t.Errorf("round trip changed the config:\n got  %+v\n want %+v", back, config)
	}
}

// A null devices map stays null, an empty one stays empty: either way the state matches the config.
func TestGroupDevicesNullAndEmptyMap(t *testing.T) {
	if d := (ConfigurationGroupDevicesConfig{}).toModel(); d.Devices != nil || d.toConfig().Devices != nil {
		t.Errorf("configuration group: null devices not kept null")
	}
	empty := ConfigurationGroupDevicesConfig{Devices: map[string]ConfigurationGroupDevicesDeviceConfig{}}
	if d := empty.toModel(); d.Devices == nil || d.toConfig().Devices == nil {
		t.Errorf("configuration group: empty devices not kept empty")
	}
	if d := (PolicyGroupDevicesConfig{}).toModel(); d.Devices != nil || d.toConfig().Devices != nil {
		t.Errorf("policy group: null devices not kept null")
	}
	pgEmpty := PolicyGroupDevicesConfig{Devices: map[string]PolicyGroupDevicesDeviceConfig{}}
	if d := pgEmpty.toModel(); d.Devices == nil || d.toConfig().Devices == nil {
		t.Errorf("policy group: empty devices not kept empty")
	}
}

func TestPolicyGroupDevicesConfigToModel(t *testing.T) {
	config := PolicyGroupDevicesConfig{
		PolicyGroupId: types.StringValue("pg1"),
		Devices: map[string]PolicyGroupDevicesDeviceConfig{
			"C8K-B": {Deploy: types.BoolValue(true)},
			"C8K-A": {Deploy: types.BoolValue(true)},
		},
	}
	data := config.toModel()
	if len(data.Devices) != 2 || data.Devices[0].Id.ValueString() != "C8K-A" || data.Devices[1].Id.ValueString() != "C8K-B" {
		t.Fatalf("devices not sorted by ID: %+v", data.Devices)
	}
	if back := data.toConfig(); !reflect.DeepEqual(back, config) {
		t.Errorf("round trip changed the config:\n got  %+v\n want %+v", back, config)
	}
}

// A variable read from the Manager has a null list_value; toConfig gives it the string element type
// the schema needs, and an empty variables map stays empty.
func TestGroupDevicesVariablesToConfig(t *testing.T) {
	data := ConfigurationGroupDevices{Devices: []ConfigurationGroupDevicesDevices{
		{Id: types.StringValue("C8K-A"), Variables: []ConfigurationGroupDevicesDevicesVariables{{Name: types.StringValue("host_name"), Value: types.StringValue("r1")}}},
		{Id: types.StringValue("C8K-B"), Variables: []ConfigurationGroupDevicesDevicesVariables{}},
	}}
	config := data.toConfig()
	v := config.Devices["C8K-A"].Variables["host_name"]
	if v.Value.ValueString() != "r1" || !v.ListValue.Equal(types.ListNull(types.StringType)) {
		t.Errorf("host_name: %+v", v)
	}
	if vars := config.Devices["C8K-B"].Variables; vars == nil || len(vars) != 0 {
		t.Errorf("empty variables not kept empty: %v", vars)
	}
}

// Policy and CoR SaaS variables, scalar and list, reach the variables body typed by Manager's schema,
// as the list shape sent them.
func TestPolicyGroupDevicesVariablesBody(t *testing.T) {
	strs := func(v ...string) types.List {
		var e []attr.Value
		for _, s := range v {
			e = append(e, types.StringValue(s))
		}
		return types.ListValueMust(types.StringType, e)
	}
	scalar := func(s string) PolicyGroupDevicesVariableConfig {
		return PolicyGroupDevicesVariableConfig{Value: types.StringValue(s), ListValue: types.ListNull(types.StringType)}
	}
	list := func(l types.List) PolicyGroupDevicesVariableConfig {
		return PolicyGroupDevicesVariableConfig{Value: types.StringNull(), ListValue: l}
	}
	config := PolicyGroupDevicesConfig{
		Solution: types.StringValue("sdwan"),
		Devices: map[string]PolicyGroupDevicesDeviceConfig{
			"C8K-A": {Deploy: types.BoolValue(true), Variables: map[string]PolicyGroupDevicesVariableConfig{
				"qos_if":                          scalar("GigabitEthernet1"),
				"qos_ifs":                         list(strs("GigabitEthernet2", "GigabitEthernet3")),
				"vlan":                            scalar("10"),
				"cloudSaasLatency_variable":       scalar("100"),
				"cloudSaasLBEnabled_variable":     scalar("true"),
				"cloudSaaSDeviceRole_variable":    scalar("client"),
				"cloudSaasInterfaceList_variable": list(strs("GigabitEthernet1", "GigabitEthernet2")),
				"cloudSaasTlocList_variable":      list(strs()),
				"ids":                             list(strs("1", "2")),
			}},
		},
	}
	varTypes := map[string]string{
		"qos_if": "string", "qos_ifs": "string", "vlan": "integer",
		"cloudSaasLatency_variable": "integer", "cloudSaasLBEnabled_variable": "boolean",
		"cloudSaaSDeviceRole_variable": "string", "cloudSaasInterfaceList_variable": "string",
		"cloudSaasTlocList_variable": "string", "ids": "integer",
	}
	body := gjson.Parse(config.toModel().toBodyPolicyGroupDeviceVariables(context.Background(), varTypes))
	if got := body.Get("devices.0.device-id").String(); got != "C8K-A" {
		t.Fatalf("device-id: %q in %s", got, body.Raw)
	}
	got := map[string]string{}
	for _, v := range body.Get("devices.0.variables").Array() {
		got[v.Get("name").String()] = v.Get("value").Raw
	}
	want := map[string]string{
		"qos_if":                          `"GigabitEthernet1"`,
		"qos_ifs":                         `["GigabitEthernet2","GigabitEthernet3"]`,
		"vlan":                            `10`,
		"cloudSaasLatency_variable":       `100`,
		"cloudSaasLBEnabled_variable":     `true`,
		"cloudSaaSDeviceRole_variable":    `"client"`,
		"cloudSaasInterfaceList_variable": `["GigabitEthernet1","GigabitEthernet2"]`,
		"cloudSaasTlocList_variable":      `[]`,
		"ids":                             `[1,2]`,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("variables body:\n got  %v\n want %v\n body %s", got, want, body.Raw)
	}
}
