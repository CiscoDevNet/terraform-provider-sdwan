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
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func attachTestVariables(kv ...string) types.Map {
	m := map[string]attr.Value{}
	for i := 0; i < len(kv); i += 2 {
		m[kv[i]] = types.StringValue(kv[i+1])
	}
	return types.MapValueMust(types.StringType, m)
}

// devices is a map in the Terraform shape and a list sorted by device ID in the working model.
func TestAttachFeatureDeviceTemplateConfigToModel(t *testing.T) {
	config := AttachFeatureDeviceTemplateConfig{
		Id:      types.StringValue("tmpl1"),
		Version: types.Int64Value(3),
		Devices: map[string]AttachFeatureDeviceTemplateDeviceConfig{
			"C8K-C": {Variables: attachTestVariables("system_host_name", "r3")},
			"C8K-A": {Variables: attachTestVariables("system_host_name", "r1", "system_system_ip", "10.0.0.1")},
			"C8K-B": {Variables: attachTestVariables()},
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
	if !data.Devices[0].Variables.Equal(config.Devices["C8K-A"].Variables) || data.Id != config.Id || data.Version != config.Version {
		t.Errorf("fields not carried over: %+v", data)
	}
	if back := data.toConfig(); !reflect.DeepEqual(back, config) {
		t.Errorf("round trip changed the config:\n got  %+v\n want %+v", back, config)
	}
}

// A null devices map stays null, an empty one stays empty, and a device read back with null
// variables gets an empty map (variables is required).
func TestAttachFeatureDeviceTemplateNullAndEmpty(t *testing.T) {
	if d := (AttachFeatureDeviceTemplateConfig{}).toModel(); d.Devices != nil || d.toConfig().Devices != nil {
		t.Errorf("null devices not kept null")
	}
	empty := AttachFeatureDeviceTemplateConfig{Devices: map[string]AttachFeatureDeviceTemplateDeviceConfig{}}
	if d := empty.toModel(); d.Devices == nil || d.toConfig().Devices == nil {
		t.Errorf("empty devices not kept empty")
	}
	data := AttachFeatureDeviceTemplate{Devices: []AttachFeatureDeviceTemplateDevice{
		{Id: types.StringValue("C8K-A"), Variables: types.MapNull(types.StringType)},
	}}
	if v := data.toConfig().Devices["C8K-A"].Variables; v.IsNull() || len(v.Elements()) != 0 {
		t.Errorf("null variables became %v", v)
	}
}

// A version 0 (list) or version 1 (set) state upgrades to devices keyed by device ID.
func TestAttachFeatureDeviceTemplateUpgradeState(t *testing.T) {
	ctx := context.Background()
	r := &AttachFeatureDeviceTemplateResource{}
	var current resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &current)
	prior := AttachFeatureDeviceTemplate{
		Id:      types.StringValue("tmpl1"),
		Version: types.Int64Value(2),
		Devices: []AttachFeatureDeviceTemplateDevice{
			{Id: types.StringValue("C8K-B"), Variables: attachTestVariables("system_host_name", "r2")},
			{Id: types.StringValue("C8K-A"), Variables: attachTestVariables("system_host_name", "r1", "system_site_id", "100")},
		},
	}
	want := AttachFeatureDeviceTemplateConfig{
		Id:      types.StringValue("tmpl1"),
		Version: types.Int64Value(2),
		Devices: map[string]AttachFeatureDeviceTemplateDeviceConfig{
			"C8K-A": {Variables: attachTestVariables("system_host_name", "r1", "system_site_id", "100")},
			"C8K-B": {Variables: attachTestVariables("system_host_name", "r2")},
		},
	}
	upgraders := r.UpgradeState(ctx)
	for _, version := range []int64{0, 1} {
		u, ok := upgraders[version]
		if !ok {
			t.Fatalf("no upgrader for version %d", version)
		}
		priorState := tfsdk.State{Schema: *u.PriorSchema, Raw: tftypes.NewValue(u.PriorSchema.Type().TerraformType(ctx), nil)}
		if d := priorState.Set(ctx, prior); d.HasError() {
			t.Fatalf("version %d: prior state: %v", version, d)
		}
		resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: current.Schema, Raw: tftypes.NewValue(current.Schema.Type().TerraformType(ctx), nil)}}
		u.StateUpgrader(ctx, resource.UpgradeStateRequest{State: &priorState}, &resp)
		if resp.Diagnostics.HasError() {
			t.Fatalf("version %d: upgrade: %v", version, resp.Diagnostics)
		}
		var got AttachFeatureDeviceTemplateConfig
		if d := resp.State.Get(ctx, &got); d.HasError() {
			t.Fatalf("version %d: upgraded state: %v", version, d)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("version %d:\n got  %+v\n want %+v", version, got, want)
		}
	}
}
