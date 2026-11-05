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

package helpers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// retain_devices on the *_group_devices resources lists devices handed over to another state that
// manages them on the same group. Releasing them would pull them off the group (and, for a policy
// group, withdraw what the controllers derive from it) until the new owner deploys them again.

// RetainedIds returns the known IDs in a retain_devices set.
func RetainedIds(retain types.Set) map[string]bool {
	ids := map[string]bool{}
	if retain.IsNull() || retain.IsUnknown() {
		return ids
	}
	for _, v := range retain.Elements() {
		if s, ok := v.(types.String); ok && !s.IsNull() && !s.IsUnknown() {
			ids[s.ValueString()] = true
		}
	}
	return ids
}

// ExceptRetained returns ids without the devices listed in retain: those are dropped from the
// state without a release. associatePath only labels the log line.
func ExceptRetained(ctx context.Context, retain types.Set, ids []string, associatePath string) []string {
	keep := RetainedIds(retain)
	var out []string
	for _, id := range ids {
		if keep[id] {
			tflog.Info(ctx, fmt.Sprintf("%s: device %s is in retain_devices, dropping it from the state without releasing it", associatePath, id))
			continue
		}
		out = append(out, id)
	}
	return out
}

// ValidateRetainDevices adds an error for each device listed in both devices and retain_devices.
// Unknown values are skipped; they are checked again once known.
func ValidateRetainDevices(ctx context.Context, config tfsdk.Config, diags *diag.Diagnostics) {
	var retain types.Set
	if d := config.GetAttribute(ctx, path.Root("retain_devices"), &retain); d.HasError() {
		return
	}
	keep := RetainedIds(retain)
	if len(keep) == 0 {
		return
	}
	var devices types.List
	if d := config.GetAttribute(ctx, path.Root("devices"), &devices); d.HasError() || devices.IsNull() || devices.IsUnknown() {
		return
	}
	for _, e := range devices.Elements() {
		obj, ok := e.(types.Object)
		if !ok || obj.IsNull() || obj.IsUnknown() {
			continue
		}
		if id, ok := obj.Attributes()["id"].(types.String); ok && !id.IsUnknown() && keep[id.ValueString()] {
			diags.AddAttributeError(path.Root("retain_devices"), "Device in devices and retain_devices",
				fmt.Sprintf("Device %s is listed in both devices and retain_devices. A retained device is handed over to another state: list it in only one of them.", id.ValueString()))
		}
	}
}

// OnlyAttributeChanged reports whether plan and state differ in the top-level attribute name only.
func OnlyAttributeChanged(plan, state tftypes.Value, name string) (bool, error) {
	attrPath := tftypes.NewAttributePath().WithAttributeName(name)
	stateValue, _, err := tftypes.WalkAttributePath(state, attrPath)
	if err != nil {
		return false, err
	}
	sv, ok := stateValue.(tftypes.Value)
	if !ok {
		return false, fmt.Errorf("attribute %s: unexpected state value %T", name, stateValue)
	}
	// The plan with the state's value for name: equal to the state when nothing else changed.
	merged, err := tftypes.Transform(plan, func(p *tftypes.AttributePath, v tftypes.Value) (tftypes.Value, error) {
		if p.Equal(attrPath) {
			return sv, nil
		}
		return v, nil
	})
	if err != nil {
		return false, err
	}
	return !plan.Equal(state) && merged.Equal(state), nil
}
