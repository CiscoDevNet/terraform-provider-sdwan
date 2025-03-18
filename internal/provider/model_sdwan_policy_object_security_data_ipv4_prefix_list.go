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
type PolicyObjectSecurityDataIPv4PrefixList struct {
	Id               types.String                                    `tfsdk:"id"`
	Version          types.Int64                                     `tfsdk:"version"`
	Name             types.String                                    `tfsdk:"name"`
	Description      types.String                                    `tfsdk:"description"`
	FeatureProfileId types.String                                    `tfsdk:"feature_profile_id"`
	Entries          []PolicyObjectSecurityDataIPv4PrefixListEntries `tfsdk:"entries"`
}

type PolicyObjectSecurityDataIPv4PrefixListEntries struct {
	IpPrefix         types.String `tfsdk:"ip_prefix"`
	IpPrefixVariable types.String `tfsdk:"ip_prefix_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectSecurityDataIPv4PrefixList) getModel() string {
	return "policy_object_security_data_ipv4_prefix_list"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectSecurityDataIPv4PrefixList) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/security-data-ip-prefix", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectSecurityDataIPv4PrefixList) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {

		for _, item := range data.Entries {
			itemBody := ""

			if !item.IpPrefixVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipPrefix.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ipPrefix.value", item.IpPrefixVariable.ValueString())
				}
			} else if !item.IpPrefix.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipPrefix.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ipPrefix.value", item.IpPrefix.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectSecurityDataIPv4PrefixList) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	if value := res.Get(path + "entries"); value.Exists() {
		data.Entries = make([]PolicyObjectSecurityDataIPv4PrefixListEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyObjectSecurityDataIPv4PrefixListEntries{}
			item.IpPrefix = types.StringNull()
			item.IpPrefixVariable = types.StringNull()
			if t := v.Get("ipPrefix.optionType"); t.Exists() {
				va := v.Get("ipPrefix.value")
				if t.String() == "variable" {
					item.IpPrefixVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IpPrefix = types.StringValue(va.String())
				}
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PolicyObjectSecurityDataIPv4PrefixList) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	for i := range data.Entries {
		keys := [...]string{"ipPrefix"}
		keyValues := [...]string{data.Entries[i].IpPrefix.ValueString()}
		keyValuesVariables := [...]string{data.Entries[i].IpPrefixVariable.ValueString()}

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
		data.Entries[i].IpPrefix = types.StringNull()
		data.Entries[i].IpPrefixVariable = types.StringNull()
		if t := r.Get("ipPrefix.optionType"); t.Exists() {
			va := r.Get("ipPrefix.value")
			if t.String() == "variable" {
				data.Entries[i].IpPrefixVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Entries[i].IpPrefix = types.StringValue(va.String())
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *PolicyObjectSecurityDataIPv4PrefixList) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if len(data.Entries) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
