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
type PolicyObjectSecurityGeolocationList struct {
	Id               types.String                                 `tfsdk:"id"`
	Version          types.Int64                                  `tfsdk:"version"`
	Name             types.String                                 `tfsdk:"name"`
	Description      types.String                                 `tfsdk:"description"`
	FeatureProfileId types.String                                 `tfsdk:"feature_profile_id"`
	Entries          []PolicyObjectSecurityGeolocationListEntries `tfsdk:"entries"`
}

type PolicyObjectSecurityGeolocationListEntries struct {
	Country   types.String `tfsdk:"country"`
	Continent types.String `tfsdk:"continent"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectSecurityGeolocationList) getModel() string {
	return "policy_object_security_geolocation_list"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectSecurityGeolocationList) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/security-geolocation", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectSecurityGeolocationList) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {

		for _, item := range data.Entries {
			itemBody := ""
			if !item.Country.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "country.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "country.value", item.Country.ValueString())
				}
			}
			if !item.Continent.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "continent.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "continent.value", item.Continent.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectSecurityGeolocationList) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	if value := res.Get(path + "entries"); value.Exists() {
		data.Entries = make([]PolicyObjectSecurityGeolocationListEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyObjectSecurityGeolocationListEntries{}
			item.Country = types.StringNull()

			if t := v.Get("country.optionType"); t.Exists() {
				va := v.Get("country.value")
				if t.String() == "global" {
					item.Country = types.StringValue(va.String())
				}
			}
			item.Continent = types.StringNull()

			if t := v.Get("continent.optionType"); t.Exists() {
				va := v.Get("continent.value")
				if t.String() == "global" {
					item.Continent = types.StringValue(va.String())
				}
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PolicyObjectSecurityGeolocationList) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	for i := range data.Entries {
		keys := [...]string{"country", "continent"}
		keyValues := [...]string{data.Entries[i].Country.ValueString(), data.Entries[i].Continent.ValueString()}
		keyValuesVariables := [...]string{"", ""}

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
		data.Entries[i].Country = types.StringNull()

		if t := r.Get("country.optionType"); t.Exists() {
			va := r.Get("country.value")
			if t.String() == "global" {
				data.Entries[i].Country = types.StringValue(va.String())
			}
		}
		data.Entries[i].Continent = types.StringNull()

		if t := r.Get("continent.optionType"); t.Exists() {
			va := r.Get("continent.value")
			if t.String() == "global" {
				data.Entries[i].Continent = types.StringValue(va.String())
			}
		}
	}
}

// End of section. //template:end updateFromBody
