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
type GeoLocationListPolicyObject struct {
	Id      types.String                         `tfsdk:"id"`
	Version types.Int64                          `tfsdk:"version"`
	Type    types.String                         `tfsdk:"type"`
	Name    types.String                         `tfsdk:"name"`
	Entries []GeoLocationListPolicyObjectEntries `tfsdk:"entries"`
}

type GeoLocationListPolicyObjectEntries struct {
	Country   types.String `tfsdk:"country"`
	Continent types.String `tfsdk:"continent"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data GeoLocationListPolicyObject) getPath() string {
	return "/template/policy/list/geolocation/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data GeoLocationListPolicyObject) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "geolocation")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "entries", []interface{}{})
		for _, item := range data.Entries {
			itemBody := ""
			if !item.Country.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "country", item.Country.ValueString())
			}
			if !item.Continent.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "continent", item.Continent.ValueString())
			}
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *GeoLocationListPolicyObject) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("entries"); value.Exists() && len(value.Array()) > 0 {
		data.Entries = make([]GeoLocationListPolicyObjectEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := GeoLocationListPolicyObjectEntries{}
			if cValue := v.Get("country"); cValue.Exists() {
				item.Country = types.StringValue(cValue.String())
			} else {
				item.Country = types.StringNull()
			}
			if cValue := v.Get("continent"); cValue.Exists() {
				item.Continent = types.StringValue(cValue.String())
			} else {
				item.Continent = types.StringNull()
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	} else {
		if len(data.Entries) > 0 {
			data.Entries = []GeoLocationListPolicyObjectEntries{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *GeoLocationListPolicyObject) hasChanges(ctx context.Context, state *GeoLocationListPolicyObject) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if len(data.Entries) != len(state.Entries) {
		hasChanges = true
	} else {
		for i := range data.Entries {
			if !data.Entries[i].Country.Equal(state.Entries[i].Country) {
				hasChanges = true
			}
			if !data.Entries[i].Continent.Equal(state.Entries[i].Continent) {
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
func (data *GeoLocationListPolicyObject) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
	data.Type = types.StringValue("geolocation")
}

// End of section. //template:end processImport
