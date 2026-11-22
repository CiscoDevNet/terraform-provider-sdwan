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
	"slices"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

func pgVars(vars ...PolicyGroupDevicesDevicesVariables) []PolicyGroupDevicesDevicesVariables {
	return vars
}

func pgVar(name, value string) PolicyGroupDevicesDevicesVariables {
	return PolicyGroupDevicesDevicesVariables{
		Name:      types.StringValue(name),
		Value:     types.StringValue(value),
		ListValue: types.ListNull(types.StringType),
	}
}

func pgVarOf(t *testing.T, d PolicyGroupDevicesDevices, name string) PolicyGroupDevicesDevicesVariables {
	t.Helper()
	for _, v := range d.Variables {
		if v.Name.ValueString() == name {
			return v
		}
	}
	t.Fatalf("device %s has no variable %s: %v", d.Id.ValueString(), name, d.Variables)
	return PolicyGroupDevicesDevicesVariables{}
}

// The Manager returns a secret variable as the placeholder $cRYPT_CLUSTER1, never its value. Read
// keeps the value already in state, so the plan has no difference and the device is not redeployed.
func TestPolicyGroupDeviceVariablesKeepEncrypted(t *testing.T) {
	ctx := context.Background()
	res := gjson.Parse(`{"family":"sdwan","devices":[
		{"device-id":"C8K-A","variables":[
			{"name":"secret","value":"$cRYPT_CLUSTER1"},
			{"name":"host_name","value":"edge1"},
			{"name":"qos_if","value":["GigabitEthernet4"]}]},
		{"device-id":"C8K-B","variables":[
			{"name":"secret","value":"$cRYPT_CLUSTER1"}]}]}`)

	// The state has a value for C8K-A's secret. C8K-B's secret has none, as after an import. C8K-B
	// has a different value for the same name: it must not be taken for C8K-A, or the reverse.
	state := PolicyGroupDevices{Devices: []PolicyGroupDevicesDevices{
		{Id: types.StringValue("C8K-A"), Variables: pgVars(pgVar("secret", "Secret-A!"), pgVar("host_name", "edge1"))},
		{Id: types.StringValue("C8K-B"), Variables: pgVars(PolicyGroupDevicesDevicesVariables{
			Name: types.StringValue("secret"), Value: types.StringNull(), ListValue: types.ListNull(types.StringType),
		})},
	}}
	oldState := state
	oldState.Devices = slices.Clone(state.Devices)

	state.fromBodyPolicyGroupDeviceVariables(ctx, res, &oldState)

	a, b := state.Devices[0], state.Devices[1]
	if got := pgVarOf(t, a, "secret").Value; got.ValueString() != "Secret-A!" {
		t.Errorf("C8K-A secret = %v, want the value from state", got)
	}
	if got := pgVarOf(t, a, "host_name").Value; got.ValueString() != "edge1" {
		t.Errorf("C8K-A host_name = %v, want the Manager's value", got)
	}
	if got := pgVarOf(t, a, "qos_if").ListValue; got.Equal(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("GigabitEthernet4")})) == false {
		t.Errorf("C8K-A qos_if = %v, want the Manager's list", got)
	}
	if got := pgVarOf(t, b, "secret").Value; !got.IsNull() {
		t.Errorf("C8K-B secret = %v, want null with no earlier value", got)
	}

	// Nothing differs from what the configuration has: no redeploy.
	plan := PolicyGroupDevices{Devices: []PolicyGroupDevicesDevices{
		{Id: types.StringValue("C8K-A"), Variables: pgVars(
			pgVar("secret", "Secret-A!"), pgVar("host_name", "edge1"),
			PolicyGroupDevicesDevicesVariables{Name: types.StringValue("qos_if"), Value: types.StringNull(),
				ListValue: types.ListValueMust(types.StringType, []attr.Value{types.StringValue("GigabitEthernet4")})},
		)},
	}}
	if updated := plan.getUpdatedDevices(ctx, &state); slices.Contains(updated, "C8K-A") {
		t.Errorf("device with an unchanged encrypted variable selected for deploy: %v", updated)
	}

	// A new value in the configuration still differs from the kept one and does select the device.
	plan.Devices[0].Variables[0] = pgVar("secret", "Rotated-A!")
	if updated := plan.getUpdatedDevices(ctx, &state); !slices.Contains(updated, "C8K-A") {
		t.Errorf("changed secret not selected for deploy: %v", updated)
	}
}

// Adopted devices have no earlier state: the planned value is used, so an unchanged secret is not
// taken for a change.
func TestPolicyGroupAdoptedPriorStateKeepsEncrypted(t *testing.T) {
	ctx := context.Background()
	res := gjson.Parse(`{"family":"sdwan","devices":[
		{"device-id":"C8K-A","variables":[{"name":"secret","value":"$cRYPT_CLUSTER1"}]}]}`)

	plan := PolicyGroupDevices{Devices: []PolicyGroupDevicesDevices{
		{Id: types.StringValue("C8K-A"), Variables: pgVars(pgVar("secret", "Secret-A!"))},
	}}
	prior := adoptedPriorState(ctx, plan, []string{"C8K-A"}, res)

	if got := pgVarOf(t, prior.Devices[0], "secret").Value; got.ValueString() != "Secret-A!" {
		t.Fatalf("adopted secret = %v, want the planned value", got)
	}
	if updated := plan.getUpdatedDevices(ctx, prior); slices.Contains(updated, "C8K-A") {
		t.Errorf("adopted device with an unchanged secret selected for deploy: %v", updated)
	}
}

// Right after an import the state has no value for an encrypted variable, so the plan differs in
// it. A device differing in those variables only is not redeployed; any other difference still is.
func TestOnlyPolicyGroupVariablesDiffer(t *testing.T) {
	state := &PolicyGroupDevices{Devices: []PolicyGroupDevicesDevices{
		{Id: types.StringValue("C8K-A"), Variables: pgVars(
			PolicyGroupDevicesDevicesVariables{Name: types.StringValue("secret"), Value: types.StringNull(), ListValue: types.ListNull(types.StringType)},
			pgVar("host_name", "edge1"),
		)},
	}}
	plan := func(vars ...PolicyGroupDevicesDevicesVariables) PolicyGroupDevices {
		return PolicyGroupDevices{Devices: []PolicyGroupDevicesDevices{{Id: types.StringValue("C8K-A"), Variables: vars}}}
	}
	encrypted := []string{"secret"}

	if !onlyPolicyGroupVariablesDiffer(plan(pgVar("secret", "Secret-A!"), pgVar("host_name", "edge1")), state, "C8K-A", encrypted) {
		t.Error("only the encrypted variable differs: want true")
	}
	if onlyPolicyGroupVariablesDiffer(plan(pgVar("secret", "Secret-A!"), pgVar("host_name", "edge2")), state, "C8K-A", encrypted) {
		t.Error("another variable changed too: want false")
	}
	if onlyPolicyGroupVariablesDiffer(plan(pgVar("secret", "Secret-A!"), pgVar("host_name", "edge1"), pgVar("qos_if", "Gi4")), state, "C8K-A", encrypted) {
		t.Error("a variable was added: want false")
	}
	if onlyPolicyGroupVariablesDiffer(plan(pgVar("secret", "Secret-A!"), pgVar("host_name", "edge1")), state, "C8K-B", encrypted) {
		t.Error("device not in plan and state: want false")
	}
}
