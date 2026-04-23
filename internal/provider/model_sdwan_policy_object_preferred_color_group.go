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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type PolicyObjectPreferredColorGroup struct {
	Id               types.String                             `tfsdk:"id"`
	Version          types.Int64                              `tfsdk:"version"`
	Name             types.String                             `tfsdk:"name"`
	Description      types.String                             `tfsdk:"description"`
	FeatureProfileId types.String                             `tfsdk:"feature_profile_id"`
	Entries          []PolicyObjectPreferredColorGroupEntries `tfsdk:"entries"`
}

type PolicyObjectPreferredColorGroupEntries struct {
	PrimaryColorPreference   types.Set    `tfsdk:"primary_color_preference"`
	PrimaryPathPreference    types.String `tfsdk:"primary_path_preference"`
	SecondaryColorPreference types.Set    `tfsdk:"secondary_color_preference"`
	SecondaryPathPreference  types.String `tfsdk:"secondary_path_preference"`
	TertiaryColorPreference  types.Set    `tfsdk:"tertiary_color_preference"`
	TertiaryPathPreference   types.String `tfsdk:"tertiary_path_preference"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectPreferredColorGroup) getModel() string {
	return "policy_object_preferred_color_group"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectPreferredColorGroup) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/preferred-color-group", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

func (data PolicyObjectPreferredColorGroup) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {

		for _, item := range data.Entries {
			itemBody := ""
			if item.PrimaryColorPreference.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "primaryPreference", map[string]interface{}{})
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "primaryPreference.colorPreference.optionType", "global")
					var values []string
					item.PrimaryColorPreference.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "primaryPreference.colorPreference.value", values)
				}
			}
			if !item.PrimaryPathPreference.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "primaryPreference.pathPreference.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "primaryPreference.pathPreference.value", item.PrimaryPathPreference.ValueString())
				}
			}
			if item.SecondaryColorPreference.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secondaryPreference", map[string]interface{}{})
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "secondaryPreference.colorPreference.optionType", "global")
					var values []string
					item.SecondaryColorPreference.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "secondaryPreference.colorPreference.value", values)
				}
			}
			if !item.SecondaryPathPreference.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "secondaryPreference.pathPreference.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "secondaryPreference.pathPreference.value", item.SecondaryPathPreference.ValueString())
				}
			}
			if item.TertiaryColorPreference.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "tertiaryPreference", map[string]interface{}{})
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tertiaryPreference.colorPreference.optionType", "global")
					var values []string
					item.TertiaryColorPreference.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "tertiaryPreference.colorPreference.value", values)
				}
			}
			if !item.TertiaryPathPreference.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tertiaryPreference.pathPreference.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tertiaryPreference.pathPreference.value", item.TertiaryPathPreference.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"entries.-1", itemBody)
		}
	}
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectPreferredColorGroup) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	oldEntries := data.Entries
	if value := res.Get(path + "entries"); value.Exists() && len(value.Array()) > 0 {
		data.Entries = make([]PolicyObjectPreferredColorGroupEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyObjectPreferredColorGroupEntries{}
			item.PrimaryColorPreference = types.SetNull(types.StringType)

			if t := v.Get("primaryPreference.colorPreference.optionType"); t.Exists() {
				va := v.Get("primaryPreference.colorPreference.value")
				if t.String() == "global" {
					item.PrimaryColorPreference = helpers.GetStringSet(va.Array())
				}
			}
			item.PrimaryPathPreference = types.StringNull()

			if t := v.Get("primaryPreference.pathPreference.optionType"); t.Exists() {
				va := v.Get("primaryPreference.pathPreference.value")
				if t.String() == "global" {
					item.PrimaryPathPreference = types.StringValue(va.String())
				}
			}
			item.SecondaryColorPreference = types.SetNull(types.StringType)

			if t := v.Get("secondaryPreference.colorPreference.optionType"); t.Exists() {
				va := v.Get("secondaryPreference.colorPreference.value")
				if t.String() == "global" {
					item.SecondaryColorPreference = helpers.GetStringSet(va.Array())
				}
			}
			item.SecondaryPathPreference = types.StringNull()

			if t := v.Get("secondaryPreference.pathPreference.optionType"); t.Exists() {
				va := v.Get("secondaryPreference.pathPreference.value")
				if t.String() == "global" {
					item.SecondaryPathPreference = types.StringValue(va.String())
				}
			}
			item.TertiaryColorPreference = types.SetNull(types.StringType)

			if t := v.Get("tertiaryPreference.colorPreference.optionType"); t.Exists() {
				va := v.Get("tertiaryPreference.colorPreference.value")
				if t.String() == "global" {
					item.TertiaryColorPreference = helpers.GetStringSet(va.Array())
				}
			}
			item.TertiaryPathPreference = types.StringNull()

			if t := v.Get("tertiaryPreference.pathPreference.optionType"); t.Exists() {
				va := v.Get("tertiaryPreference.pathPreference.value")
				if t.String() == "global" {
					item.TertiaryPathPreference = types.StringValue(va.String())
				}
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	} else {
		data.Entries = nil
	}
	if !fullRead {
		resultEntries := make([]PolicyObjectPreferredColorGroupEntries, 0, len(data.Entries))
		matchedEntries := make([]bool, len(data.Entries))
		for _, oldItem := range oldEntries {
			for ni := range data.Entries {
				if matchedEntries[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.PrimaryColorPreference).ValueString() != helpers.GetStringFromSet(data.Entries[ni].PrimaryColorPreference).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if oldItem.PrimaryPathPreference.ValueString() != data.Entries[ni].PrimaryPathPreference.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.SecondaryColorPreference).ValueString() != helpers.GetStringFromSet(data.Entries[ni].SecondaryColorPreference).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if oldItem.SecondaryPathPreference.ValueString() != data.Entries[ni].SecondaryPathPreference.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.TertiaryColorPreference).ValueString() != helpers.GetStringFromSet(data.Entries[ni].TertiaryColorPreference).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if oldItem.TertiaryPathPreference.ValueString() != data.Entries[ni].TertiaryPathPreference.ValueString() {
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
