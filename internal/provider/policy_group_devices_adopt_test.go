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

func pgDevice(id, qosIf string) PolicyGroupDevicesDevices {
	return PolicyGroupDevicesDevices{
		Id:     types.StringValue(id),
		Deploy: types.BoolValue(true),
		Variables: []PolicyGroupDevicesDevicesVariables{{
			Name:      types.StringValue("qos_if"),
			Value:     types.StringNull(),
			ListValue: types.ListValueMust(types.StringType, []attr.Value{types.StringValue(qosIf)}),
		}},
	}
}

// A device taken over from another state is compared against the Manager's variables, not
// treated as new: unchanged variables select nothing, a changed one selects the device.
func TestPolicyGroupAdoptedPriorState(t *testing.T) {
	ctx := context.Background()
	// GET .../device/variables/ as the Manager returns it: the adopted device and one this
	// resource doesn't manage.
	res := gjson.Parse(`{"family":"sdwan","devices":[
		{"device-id":"C8K-A","variables":[{"name":"qos_if","value":["GigabitEthernet4"]}]},
		{"device-id":"C8K-OTHER","variables":[{"name":"qos_if","value":["GigabitEthernet9"]}]}]}`)

	plan := PolicyGroupDevices{Devices: []PolicyGroupDevicesDevices{pgDevice("C8K-A", "GigabitEthernet4"), pgDevice("C8K-B", "GigabitEthernet4")}}
	prior := adoptedPriorState(ctx, plan, []string{"C8K-A"}, res)

	if len(prior.Devices) != 1 || prior.Devices[0].Id.ValueString() != "C8K-A" {
		t.Fatalf("prior devices = %v, want only C8K-A", prior.Devices)
	}
	if got := prior.Devices[0].Variables; len(got) != 1 || got[0].ListValue.String() != `["GigabitEthernet4"]` {
		t.Fatalf("prior variables = %v, want the Manager's qos_if", got)
	}
	if updated := plan.getUpdatedDevices(ctx, prior); slices.Contains(updated, "C8K-A") {
		t.Errorf("unchanged adopted device selected for deploy: %v", updated)
	}

	changed := PolicyGroupDevices{Devices: []PolicyGroupDevicesDevices{pgDevice("C8K-A", "GigabitEthernet5")}}
	if updated := changed.getUpdatedDevices(ctx, prior); !slices.Contains(updated, "C8K-A") {
		t.Errorf("adopted device with a changed variable not selected: %v", updated)
	}
}
