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
type PolicyObjectSecurityLocalDomainList struct {
	Id               types.String                                 `tfsdk:"id"`
	Version          types.Int64                                  `tfsdk:"version"`
	Name             types.String                                 `tfsdk:"name"`
	Description      types.String                                 `tfsdk:"description"`
	FeatureProfileId types.String                                 `tfsdk:"feature_profile_id"`
	Entries          []PolicyObjectSecurityLocalDomainListEntries `tfsdk:"entries"`
}

type PolicyObjectSecurityLocalDomainListEntries struct {
	LocalDomain types.String `tfsdk:"local_domain"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectSecurityLocalDomainList) getModel() string {
	return "policy_object_security_local_domain_list"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectSecurityLocalDomainList) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/security-localdomain", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectSecurityLocalDomainList) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {

		for _, item := range data.Entries {
			itemBody := ""
			if !item.LocalDomain.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nameServer.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "nameServer.value", item.LocalDomain.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectSecurityLocalDomainList) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	if value := res.Get(path + "entries"); value.Exists() && len(value.Array()) > 0 {
		data.Entries = make([]PolicyObjectSecurityLocalDomainListEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyObjectSecurityLocalDomainListEntries{}
			item.LocalDomain = types.StringNull()

			if t := v.Get("nameServer.optionType"); t.Exists() {
				va := v.Get("nameServer.value")
				if t.String() == "global" {
					item.LocalDomain = types.StringValue(va.String())
				}
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PolicyObjectSecurityLocalDomainList) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	for i := range data.Entries {
		keys := [...]string{"nameServer"}
		keyValues := [...]string{data.Entries[i].LocalDomain.ValueString()}
		keyValuesVariables := [...]string{""}

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
						} else if tt.String() == "default" {
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
		data.Entries[i].LocalDomain = types.StringNull()

		if t := r.Get("nameServer.optionType"); t.Exists() {
			va := r.Get("nameServer.value")
			if t.String() == "global" {
				data.Entries[i].LocalDomain = types.StringValue(va.String())
			}
		}
	}
}

// End of section. //template:end updateFromBody
