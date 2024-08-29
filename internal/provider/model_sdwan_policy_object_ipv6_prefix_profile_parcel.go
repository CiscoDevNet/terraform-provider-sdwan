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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type PolicyObjectIPv6Prefix struct {
	Id               types.String                    `tfsdk:"id"`
	Version          types.Int64                     `tfsdk:"version"`
	Name             types.String                    `tfsdk:"name"`
	Description      types.String                    `tfsdk:"description"`
	FeatureProfileId types.String                    `tfsdk:"feature_profile_id"`
	Entries          []PolicyObjectIPv6PrefixEntries `tfsdk:"entries"`
}

type PolicyObjectIPv6PrefixEntries struct {
	Ipv6Address         types.String `tfsdk:"ipv6_address"`
	Ipv6PrefixLength    types.Int64  `tfsdk:"ipv6_prefix_length"`
	LeRangePrefixLength types.Int64  `tfsdk:"le_range_prefix_length"`
	GeRangePrefixLength types.Int64  `tfsdk:"ge_range_prefix_length"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectIPv6Prefix) getModel() string {
	return "policy_object_ipv6_prefix"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectIPv6Prefix) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/ipv6-prefix", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectIPv6Prefix) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	for _, item := range data.Entries {
		itemBody := ""
		if !item.Ipv6Address.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ipv6Address.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "ipv6Address.value", item.Ipv6Address.ValueString())
		}
		if !item.Ipv6PrefixLength.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ipv6PrefixLength.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "ipv6PrefixLength.value", item.Ipv6PrefixLength.ValueInt64())
		}
		if !item.LeRangePrefixLength.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "leRangePrefixLength.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "leRangePrefixLength.value", item.LeRangePrefixLength.ValueInt64())
		}
		if !item.GeRangePrefixLength.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "geRangePrefixLength.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "geRangePrefixLength.value", item.GeRangePrefixLength.ValueInt64())
		}
		body, _ = sjson.SetRaw(body, path+"entries.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectIPv6Prefix) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	if value := res.Get(path + "entries"); value.Exists() {
		data.Entries = make([]PolicyObjectIPv6PrefixEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyObjectIPv6PrefixEntries{}
			item.Ipv6Address = types.StringNull()

			if t := v.Get("ipv6Address.optionType"); t.Exists() {
				va := v.Get("ipv6Address.value")
				if t.String() == "global" {
					item.Ipv6Address = types.StringValue(va.String())
				}
			}
			item.Ipv6PrefixLength = types.Int64Null()

			if t := v.Get("ipv6PrefixLength.optionType"); t.Exists() {
				va := v.Get("ipv6PrefixLength.value")
				if t.String() == "global" {
					item.Ipv6PrefixLength = types.Int64Value(va.Int())
				}
			}
			item.LeRangePrefixLength = types.Int64Null()

			if t := v.Get("leRangePrefixLength.optionType"); t.Exists() {
				va := v.Get("leRangePrefixLength.value")
				if t.String() == "global" {
					item.LeRangePrefixLength = types.Int64Value(va.Int())
				}
			}
			item.GeRangePrefixLength = types.Int64Null()

			if t := v.Get("geRangePrefixLength.optionType"); t.Exists() {
				va := v.Get("geRangePrefixLength.value")
				if t.String() == "global" {
					item.GeRangePrefixLength = types.Int64Value(va.Int())
				}
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PolicyObjectIPv6Prefix) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	for i := range data.Entries {
		keys := [...]string{"ipv6Address", "ipv6PrefixLength"}
		keyValues := [...]string{data.Entries[i].Ipv6Address.ValueString(), strconv.FormatInt(data.Entries[i].Ipv6PrefixLength.ValueInt64(), 10)}
		keyValuesVariables := [...]string{"", ""}

		var r gjson.Result
		res.Get(path + "entries").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Entries[i].Ipv6Address = types.StringNull()

		if t := r.Get("ipv6Address.optionType"); t.Exists() {
			va := r.Get("ipv6Address.value")
			if t.String() == "global" {
				data.Entries[i].Ipv6Address = types.StringValue(va.String())
			}
		}
		data.Entries[i].Ipv6PrefixLength = types.Int64Null()

		if t := r.Get("ipv6PrefixLength.optionType"); t.Exists() {
			va := r.Get("ipv6PrefixLength.value")
			if t.String() == "global" {
				data.Entries[i].Ipv6PrefixLength = types.Int64Value(va.Int())
			}
		}
		data.Entries[i].LeRangePrefixLength = types.Int64Null()

		if t := r.Get("leRangePrefixLength.optionType"); t.Exists() {
			va := r.Get("leRangePrefixLength.value")
			if t.String() == "global" {
				data.Entries[i].LeRangePrefixLength = types.Int64Value(va.Int())
			}
		}
		data.Entries[i].GeRangePrefixLength = types.Int64Null()

		if t := r.Get("geRangePrefixLength.optionType"); t.Exists() {
			va := r.Get("geRangePrefixLength.value")
			if t.String() == "global" {
				data.Entries[i].GeRangePrefixLength = types.Int64Value(va.Int())
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *PolicyObjectIPv6Prefix) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if len(data.Entries) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
