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
	"slices"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestExceptRetained(t *testing.T) {
	ctx := context.Background()
	retain := types.SetValueMust(types.StringType, []attr.Value{types.StringValue("C8K-B")})
	cases := []struct {
		name   string
		retain types.Set
		ids    []string
		want   []string
	}{
		{"retained device dropped", retain, []string{"C8K-A", "C8K-B"}, []string{"C8K-A"}},
		{"null set keeps all", types.SetNull(types.StringType), []string{"C8K-A", "C8K-B"}, []string{"C8K-A", "C8K-B"}},
		{"unknown set keeps all", types.SetUnknown(types.StringType), []string{"C8K-B"}, []string{"C8K-B"}},
		{"only retained", retain, []string{"C8K-B"}, nil},
	}
	for _, c := range cases {
		if got := ExceptRetained(ctx, c.retain, c.ids, "/v1/test/"); !slices.Equal(got, c.want) {
			t.Errorf("%s: got %v, want %v", c.name, got, c.want)
		}
	}
}

func TestValidateRetainDevices(t *testing.T) {
	ctx := context.Background()
	s := schema.Schema{Attributes: map[string]schema.Attribute{
		"retain_devices": schema.SetAttribute{ElementType: types.StringType, Optional: true},
		"devices": schema.ListNestedAttribute{Optional: true, NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{"id": schema.StringAttribute{Optional: true}},
		}},
	}}
	devType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"id": tftypes.String}}
	cfg := func(retain []string, devices []string) tfsdk.Config {
		r := []tftypes.Value{}
		for _, s := range retain {
			r = append(r, tftypes.NewValue(tftypes.String, s))
		}
		d := []tftypes.Value{}
		for _, s := range devices {
			d = append(d, tftypes.NewValue(devType, map[string]tftypes.Value{"id": tftypes.NewValue(tftypes.String, s)}))
		}
		return tfsdk.Config{Schema: s, Raw: tftypes.NewValue(tftypes.Object{AttributeTypes: map[string]tftypes.Type{
			"retain_devices": tftypes.Set{ElementType: tftypes.String},
			"devices":        tftypes.List{ElementType: devType},
		}}, map[string]tftypes.Value{
			"retain_devices": tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, r),
			"devices":        tftypes.NewValue(tftypes.List{ElementType: devType}, d),
		})}
	}
	cases := []struct {
		name            string
		retain, devices []string
		wantErrors      int
	}{
		{"disjoint", []string{"C8K-B"}, []string{"C8K-A"}, 0},
		{"overlap", []string{"C8K-A", "C8K-B"}, []string{"C8K-A"}, 1},
		{"empty retain", nil, []string{"C8K-A"}, 0},
	}
	for _, c := range cases {
		var diags diag.Diagnostics
		ValidateRetainDevices(ctx, cfg(c.retain, c.devices), &diags)
		if diags.ErrorsCount() != c.wantErrors {
			t.Errorf("%s: %d errors, want %d: %v", c.name, diags.ErrorsCount(), c.wantErrors, diags)
		}
	}
}

func TestOnlyAttributeChanged(t *testing.T) {
	objType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{
		"id":             tftypes.String,
		"retain_devices": tftypes.Set{ElementType: tftypes.String},
		"devices":        tftypes.List{ElementType: tftypes.String},
	}}
	obj := func(retain []string, devices []string) tftypes.Value {
		r := []tftypes.Value{}
		for _, s := range retain {
			r = append(r, tftypes.NewValue(tftypes.String, s))
		}
		d := []tftypes.Value{}
		for _, s := range devices {
			d = append(d, tftypes.NewValue(tftypes.String, s))
		}
		rv := tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, nil)
		if retain != nil {
			rv = tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, r)
		}
		return tftypes.NewValue(objType, map[string]tftypes.Value{
			"id":             tftypes.NewValue(tftypes.String, "g1"),
			"retain_devices": rv,
			"devices":        tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, d),
		})
	}
	cases := []struct {
		name        string
		plan, state tftypes.Value
		want        bool
	}{
		{"only retain cleared", obj(nil, []string{"A"}), obj([]string{"B"}, []string{"A"}), true},
		{"only retain set", obj([]string{"B"}, []string{"A"}), obj(nil, []string{"A"}), true},
		{"devices changed too", obj([]string{"B"}, []string{"A"}), obj(nil, []string{"A", "B"}), false},
		{"devices only", obj(nil, []string{"A"}), obj(nil, []string{"A", "B"}), false},
		{"nothing changed", obj(nil, []string{"A"}), obj(nil, []string{"A"}), false},
	}
	for _, c := range cases {
		got, err := OnlyAttributeChanged(c.plan, c.state, "retain_devices")
		if err != nil || got != c.want {
			t.Errorf("%s: got %v (err %v), want %v", c.name, got, err, c.want)
		}
	}
}
