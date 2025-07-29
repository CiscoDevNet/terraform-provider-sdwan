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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type AppProbeClassPolicyObject struct {
	Id              types.String                        `tfsdk:"id"`
	Version         types.Int64                         `tfsdk:"version"`
	Type            types.String                        `tfsdk:"type"`
	Name            types.String                        `tfsdk:"name"`
	ForwardingClass types.String                        `tfsdk:"forwarding_class"`
	Mappings        []AppProbeClassPolicyObjectMappings `tfsdk:"mappings"`
}

type AppProbeClassPolicyObjectMappings struct {
	Color types.String `tfsdk:"color"`
	Dscp  types.Int64  `tfsdk:"dscp"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data AppProbeClassPolicyObject) getPath() string {
	return "/template/policy/list/appprobe/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data AppProbeClassPolicyObject) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "appProbe")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.ForwardingClass.IsNull() {
		body, _ = sjson.Set(body, "entries.0.forwardingClass", data.ForwardingClass.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "entries.0.map", []interface{}{})
		for _, item := range data.Mappings {
			itemBody := ""
			if !item.Color.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "color", item.Color.ValueString())
			}
			if !item.Dscp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dscp", item.Dscp.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "entries.0.map.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *AppProbeClassPolicyObject) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("entries.0.forwardingClass"); value.Exists() {
		data.ForwardingClass = types.StringValue(value.String())
	} else {
		data.ForwardingClass = types.StringNull()
	}
	if value := res.Get("entries.0.map"); value.Exists() && len(value.Array()) > 0 {
		data.Mappings = make([]AppProbeClassPolicyObjectMappings, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AppProbeClassPolicyObjectMappings{}
			if cValue := v.Get("color"); cValue.Exists() {
				item.Color = types.StringValue(cValue.String())
			} else {
				item.Color = types.StringNull()
			}
			if cValue := v.Get("dscp"); cValue.Exists() {
				item.Dscp = types.Int64Value(cValue.Int())
			} else {
				item.Dscp = types.Int64Null()
			}
			data.Mappings = append(data.Mappings, item)
			return true
		})
	} else {
		if len(data.Mappings) > 0 {
			data.Mappings = []AppProbeClassPolicyObjectMappings{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *AppProbeClassPolicyObject) hasChanges(ctx context.Context, state *AppProbeClassPolicyObject) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.ForwardingClass.Equal(state.ForwardingClass) {
		hasChanges = true
	}
	if len(data.Mappings) != len(state.Mappings) {
		hasChanges = true
	} else {
		for i := range data.Mappings {
			if !data.Mappings[i].Color.Equal(state.Mappings[i].Color) {
				hasChanges = true
			}
			if !data.Mappings[i].Dscp.Equal(state.Mappings[i].Dscp) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *AppProbeClassPolicyObject) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
	data.Type = types.StringValue("appProbe")
}

// End of section. //template:end processImport
