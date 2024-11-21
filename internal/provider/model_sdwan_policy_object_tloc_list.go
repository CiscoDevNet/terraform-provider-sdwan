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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type PolicyObjectTLOCList struct {
	Id               types.String                  `tfsdk:"id"`
	Version          types.Int64                   `tfsdk:"version"`
	Name             types.String                  `tfsdk:"name"`
	Description      types.String                  `tfsdk:"description"`
	FeatureProfileId types.String                  `tfsdk:"feature_profile_id"`
	Entries          []PolicyObjectTLOCListEntries `tfsdk:"entries"`
}

type PolicyObjectTLOCListEntries struct {
	TlocIp        types.String `tfsdk:"tloc_ip"`
	Color         types.String `tfsdk:"color"`
	Encapsulation types.String `tfsdk:"encapsulation"`
	Preference    types.String `tfsdk:"preference"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectTLOCList) getModel() string {
	return "policy_object_tloc_list"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectTLOCList) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/tloc", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectTLOCList) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {

		for _, item := range data.Entries {
			itemBody := ""
			if !item.TlocIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tloc.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tloc.value", item.TlocIp.ValueString())
				}
			}
			if !item.Color.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "color.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "color.value", item.Color.ValueString())
				}
			}
			if !item.Encapsulation.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "encap.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "encap.value", item.Encapsulation.ValueString())
				}
			}
			if !item.Preference.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "preference.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "preference.value", item.Preference.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectTLOCList) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	if value := res.Get(path + "entries"); value.Exists() {
		data.Entries = make([]PolicyObjectTLOCListEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyObjectTLOCListEntries{}
			item.TlocIp = types.StringNull()

			if t := v.Get("tloc.optionType"); t.Exists() {
				va := v.Get("tloc.value")
				if t.String() == "global" {
					item.TlocIp = types.StringValue(va.String())
				}
			}
			item.Color = types.StringNull()

			if t := v.Get("color.optionType"); t.Exists() {
				va := v.Get("color.value")
				if t.String() == "global" {
					item.Color = types.StringValue(va.String())
				}
			}
			item.Encapsulation = types.StringNull()

			if t := v.Get("encap.optionType"); t.Exists() {
				va := v.Get("encap.value")
				if t.String() == "global" {
					item.Encapsulation = types.StringValue(va.String())
				}
			}
			item.Preference = types.StringNull()

			if t := v.Get("preference.optionType"); t.Exists() {
				va := v.Get("preference.value")
				if t.String() == "global" {
					item.Preference = types.StringValue(va.String())
				}
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PolicyObjectTLOCList) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	for i := range data.Entries {
		keys := [...]string{"tloc", "color", "encap", "preference"}
		keyValues := [...]string{data.Entries[i].TlocIp.ValueString(), data.Entries[i].Color.ValueString(), data.Entries[i].Encapsulation.ValueString(), data.Entries[i].Preference.ValueString()}
		keyValuesVariables := [...]string{"", "", "", ""}

		var r gjson.Result
		res.Get(path + "entries").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Entries[i].TlocIp = types.StringNull()

		if t := r.Get("tloc.optionType"); t.Exists() {
			va := r.Get("tloc.value")
			if t.String() == "global" {
				data.Entries[i].TlocIp = types.StringValue(va.String())
			}
		}
		data.Entries[i].Color = types.StringNull()

		if t := r.Get("color.optionType"); t.Exists() {
			va := r.Get("color.value")
			if t.String() == "global" {
				data.Entries[i].Color = types.StringValue(va.String())
			}
		}
		data.Entries[i].Encapsulation = types.StringNull()

		if t := r.Get("encap.optionType"); t.Exists() {
			va := r.Get("encap.value")
			if t.String() == "global" {
				data.Entries[i].Encapsulation = types.StringValue(va.String())
			}
		}
		data.Entries[i].Preference = types.StringNull()

		if t := r.Get("preference.optionType"); t.Exists() {
			va := r.Get("preference.value")
			if t.String() == "global" {
				data.Entries[i].Preference = types.StringValue(va.String())
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *PolicyObjectTLOCList) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if len(data.Entries) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
