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
type PolicyObjectSecurityLocalApplicationList struct {
	Id               types.String                                      `tfsdk:"id"`
	Version          types.Int64                                       `tfsdk:"version"`
	Name             types.String                                      `tfsdk:"name"`
	Description      types.String                                      `tfsdk:"description"`
	FeatureProfileId types.String                                      `tfsdk:"feature_profile_id"`
	Entries          []PolicyObjectSecurityLocalApplicationListEntries `tfsdk:"entries"`
}

type PolicyObjectSecurityLocalApplicationListEntries struct {
	App       types.String `tfsdk:"app"`
	AppFamily types.String `tfsdk:"app_family"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectSecurityLocalApplicationList) getModel() string {
	return "policy_object_security_local_application_list"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectSecurityLocalApplicationList) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/security-localapp", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectSecurityLocalApplicationList) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {

		for _, item := range data.Entries {
			itemBody := ""
			if !item.App.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "app.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "app.value", item.App.ValueString())
				}
			}
			if !item.AppFamily.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "appFamily.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "appFamily.value", item.AppFamily.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectSecurityLocalApplicationList) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	if value := res.Get(path + "entries"); value.Exists() {
		data.Entries = make([]PolicyObjectSecurityLocalApplicationListEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyObjectSecurityLocalApplicationListEntries{}
			item.App = types.StringNull()

			if t := v.Get("app.optionType"); t.Exists() {
				va := v.Get("app.value")
				if t.String() == "global" {
					item.App = types.StringValue(va.String())
				}
			}
			item.AppFamily = types.StringNull()

			if t := v.Get("appFamily.optionType"); t.Exists() {
				va := v.Get("appFamily.value")
				if t.String() == "global" {
					item.AppFamily = types.StringValue(va.String())
				}
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PolicyObjectSecurityLocalApplicationList) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	for i := range data.Entries {
		keys := [...]string{"app", "appFamily"}
		keyValues := [...]string{data.Entries[i].App.ValueString(), data.Entries[i].AppFamily.ValueString()}
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
		data.Entries[i].App = types.StringNull()

		if t := r.Get("app.optionType"); t.Exists() {
			va := r.Get("app.value")
			if t.String() == "global" {
				data.Entries[i].App = types.StringValue(va.String())
			}
		}
		data.Entries[i].AppFamily = types.StringNull()

		if t := r.Get("appFamily.optionType"); t.Exists() {
			va := r.Get("appFamily.value")
			if t.String() == "global" {
				data.Entries[i].AppFamily = types.StringValue(va.String())
			}
		}
	}
}

// End of section. //template:end updateFromBody
