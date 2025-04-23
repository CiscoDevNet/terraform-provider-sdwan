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
type PolicyObjectPolicer struct {
	Id               types.String                 `tfsdk:"id"`
	Version          types.Int64                  `tfsdk:"version"`
	Name             types.String                 `tfsdk:"name"`
	Description      types.String                 `tfsdk:"description"`
	FeatureProfileId types.String                 `tfsdk:"feature_profile_id"`
	Entries          []PolicyObjectPolicerEntries `tfsdk:"entries"`
}

type PolicyObjectPolicerEntries struct {
	BurstBytes   types.Int64  `tfsdk:"burst_bytes"`
	ExceedAction types.String `tfsdk:"exceed_action"`
	RateBps      types.Int64  `tfsdk:"rate_bps"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectPolicer) getModel() string {
	return "policy_object_policer"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectPolicer) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/policer", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectPolicer) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {

		for _, item := range data.Entries {
			itemBody := ""
			if !item.BurstBytes.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "burst.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "burst.value", item.BurstBytes.ValueInt64())
				}
			}
			if !item.ExceedAction.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "exceed.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "exceed.value", item.ExceedAction.ValueString())
				}
			}
			if !item.RateBps.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "rate.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "rate.value", item.RateBps.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectPolicer) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	if value := res.Get(path + "entries"); value.Exists() {
		data.Entries = make([]PolicyObjectPolicerEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyObjectPolicerEntries{}
			item.BurstBytes = types.Int64Null()

			if t := v.Get("burst.optionType"); t.Exists() {
				va := v.Get("burst.value")
				if t.String() == "global" {
					item.BurstBytes = types.Int64Value(va.Int())
				}
			}
			item.ExceedAction = types.StringNull()

			if t := v.Get("exceed.optionType"); t.Exists() {
				va := v.Get("exceed.value")
				if t.String() == "global" {
					item.ExceedAction = types.StringValue(va.String())
				}
			}
			item.RateBps = types.Int64Null()

			if t := v.Get("rate.optionType"); t.Exists() {
				va := v.Get("rate.value")
				if t.String() == "global" {
					item.RateBps = types.Int64Value(va.Int())
				}
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PolicyObjectPolicer) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	for i := range data.Entries {
		keys := [...]string{"burst", "exceed", "rate"}
		keyValues := [...]string{strconv.FormatInt(data.Entries[i].BurstBytes.ValueInt64(), 10), data.Entries[i].ExceedAction.ValueString(), strconv.FormatInt(data.Entries[i].RateBps.ValueInt64(), 10)}
		keyValuesVariables := [...]string{"", "", ""}

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
		data.Entries[i].BurstBytes = types.Int64Null()

		if t := r.Get("burst.optionType"); t.Exists() {
			va := r.Get("burst.value")
			if t.String() == "global" {
				data.Entries[i].BurstBytes = types.Int64Value(va.Int())
			}
		}
		data.Entries[i].ExceedAction = types.StringNull()

		if t := r.Get("exceed.optionType"); t.Exists() {
			va := r.Get("exceed.value")
			if t.String() == "global" {
				data.Entries[i].ExceedAction = types.StringValue(va.String())
			}
		}
		data.Entries[i].RateBps = types.Int64Null()

		if t := r.Get("rate.optionType"); t.Exists() {
			va := r.Get("rate.value")
			if t.String() == "global" {
				data.Entries[i].RateBps = types.Int64Value(va.Int())
			}
		}
	}
}

// End of section. //template:end updateFromBody
