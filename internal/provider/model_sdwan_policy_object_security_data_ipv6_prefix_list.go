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
type PolicyObjectSecurityDataIPv6PrefixList struct {
	Id               types.String                                    `tfsdk:"id"`
	Version          types.Int64                                     `tfsdk:"version"`
	Name             types.String                                    `tfsdk:"name"`
	Description      types.String                                    `tfsdk:"description"`
	FeatureProfileId types.String                                    `tfsdk:"feature_profile_id"`
	Entries          []PolicyObjectSecurityDataIPv6PrefixListEntries `tfsdk:"entries"`
}

type PolicyObjectSecurityDataIPv6PrefixListEntries struct {
	Ipv6Prefix         types.String `tfsdk:"ipv6_prefix"`
	Ipv6PrefixVariable types.String `tfsdk:"ipv6_prefix_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectSecurityDataIPv6PrefixList) getModel() string {
	return "policy_object_security_data_ipv6_prefix_list"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectSecurityDataIPv6PrefixList) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/security-data-ipv6-prefix", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectSecurityDataIPv6PrefixList) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {

		for _, item := range data.Entries {
			itemBody := ""

			if !item.Ipv6PrefixVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipv6Prefix.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ipv6Prefix.value", item.Ipv6PrefixVariable.ValueString())
				}
			} else if !item.Ipv6Prefix.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipv6Prefix.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ipv6Prefix.value", item.Ipv6Prefix.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectSecurityDataIPv6PrefixList) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	oldEntries := data.Entries
	if value := res.Get(path + "entries"); value.Exists() && len(value.Array()) > 0 {
		data.Entries = make([]PolicyObjectSecurityDataIPv6PrefixListEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyObjectSecurityDataIPv6PrefixListEntries{}
			item.Ipv6Prefix = types.StringNull()
			item.Ipv6PrefixVariable = types.StringNull()
			if t := v.Get("ipv6Prefix.optionType"); t.Exists() {
				va := v.Get("ipv6Prefix.value")
				if t.String() == "variable" {
					item.Ipv6PrefixVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Ipv6Prefix = types.StringValue(va.String())
				}
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	} else {
		data.Entries = nil
	}
	if !fullRead && data.Entries != nil {
		resultEntries := make([]PolicyObjectSecurityDataIPv6PrefixListEntries, 0, len(data.Entries))
		matchedEntries := make([]bool, len(data.Entries))
		for _, oldItem := range oldEntries {
			for ni := range data.Entries {
				if matchedEntries[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.Ipv6PrefixVariable.ValueString() != "" || data.Entries[ni].Ipv6PrefixVariable.ValueString() != "") {
					if oldItem.Ipv6PrefixVariable.ValueString() != data.Entries[ni].Ipv6PrefixVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.Ipv6Prefix.ValueString() != data.Entries[ni].Ipv6Prefix.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedEntries[ni] = true
					resultEntries = append(resultEntries, data.Entries[ni])
					break
				}
			}
		}
		for ni := range data.Entries {
			if !matchedEntries[ni] {
				resultEntries = append(resultEntries, data.Entries[ni])
			}
		}
		data.Entries = resultEntries
	}
}

// End of section. //template:end fromBody
