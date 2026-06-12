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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type TopologyHubSpoke struct {
	Id               types.String             `tfsdk:"id"`
	Version          types.Int64              `tfsdk:"version"`
	Name             types.String             `tfsdk:"name"`
	Description      types.String             `tfsdk:"description"`
	FeatureProfileId types.String             `tfsdk:"feature_profile_id"`
	TargetVpns       types.Set                `tfsdk:"target_vpns"`
	SelectedHubs     types.Set                `tfsdk:"selected_hubs"`
	Spokes           []TopologyHubSpokeSpokes `tfsdk:"spokes"`
}

type TopologyHubSpokeSpokes struct {
	Name       types.String                     `tfsdk:"name"`
	SpokeSites types.Set                        `tfsdk:"spoke_sites"`
	HubSites   []TopologyHubSpokeSpokesHubSites `tfsdk:"hub_sites"`
}

type TopologyHubSpokeSpokesHubSites struct {
	Sites      types.Set   `tfsdk:"sites"`
	Preference types.Int64 `tfsdk:"preference"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TopologyHubSpoke) getModel() string {
	return "topology_hubspoke"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TopologyHubSpoke) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/topology/%v/hubspoke", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TopologyHubSpoke) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.TargetVpns.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"target.vpn.optionType", "global")
			var values []string
			data.TargetVpns.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"target.vpn.value", values)
		}
	}
	if !data.SelectedHubs.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"selectedHubs.optionType", "global")
			var values []string
			data.SelectedHubs.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"selectedHubs.value", values)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"spokes", []interface{}{})
		for _, item := range data.Spokes {
			itemBody := ""
			if !item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
				}
			}
			if !item.SpokeSites.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "spokeSites.optionType", "global")
					var values []string
					item.SpokeSites.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "spokeSites.value", values)
				}
			}
			if true {

				for _, childItem := range item.HubSites {
					itemChildBody := ""
					if !childItem.Sites.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sites.optionType", "global")
							var values []string
							childItem.Sites.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "sites.value", values)
						}
					}
					if !childItem.Preference.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "preference.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "preference.value", childItem.Preference.ValueInt64())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "hubSites.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"spokes.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TopologyHubSpoke) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.TargetVpns = types.SetNull(types.StringType)

	if t := res.Get(path + "target.vpn.optionType"); t.Exists() {
		va := res.Get(path + "target.vpn.value")
		if t.String() == "global" {
			data.TargetVpns = helpers.GetStringSet(va.Array())
		}
	}
	data.SelectedHubs = types.SetNull(types.StringType)

	if t := res.Get(path + "selectedHubs.optionType"); t.Exists() {
		va := res.Get(path + "selectedHubs.value")
		if t.String() == "global" {
			data.SelectedHubs = helpers.GetStringSet(va.Array())
		}
	}
	if value := res.Get(path + "spokes"); value.Exists() && len(value.Array()) > 0 {
		data.Spokes = make([]TopologyHubSpokeSpokes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TopologyHubSpokeSpokes{}
			item.Name = types.StringNull()

			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			item.SpokeSites = types.SetNull(types.StringType)

			if t := v.Get("spokeSites.optionType"); t.Exists() {
				va := v.Get("spokeSites.value")
				if t.String() == "global" {
					item.SpokeSites = helpers.GetStringSet(va.Array())
				}
			}
			if cValue := v.Get("hubSites"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.HubSites = make([]TopologyHubSpokeSpokesHubSites, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TopologyHubSpokeSpokesHubSites{}
					cItem.Sites = types.SetNull(types.StringType)

					if t := cv.Get("sites.optionType"); t.Exists() {
						va := cv.Get("sites.value")
						if t.String() == "global" {
							cItem.Sites = helpers.GetStringSet(va.Array())
						}
					}
					cItem.Preference = types.Int64Null()

					if t := cv.Get("preference.optionType"); t.Exists() {
						va := cv.Get("preference.value")
						if t.String() == "global" {
							cItem.Preference = types.Int64Value(va.Int())
						}
					}
					item.HubSites = append(item.HubSites, cItem)
					return true
				})
			}
			data.Spokes = append(data.Spokes, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TopologyHubSpoke) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.TargetVpns = types.SetNull(types.StringType)

	if t := res.Get(path + "target.vpn.optionType"); t.Exists() {
		va := res.Get(path + "target.vpn.value")
		if t.String() == "global" {
			data.TargetVpns = helpers.GetStringSet(va.Array())
		}
	}
	data.SelectedHubs = types.SetNull(types.StringType)

	if t := res.Get(path + "selectedHubs.optionType"); t.Exists() {
		va := res.Get(path + "selectedHubs.value")
		if t.String() == "global" {
			data.SelectedHubs = helpers.GetStringSet(va.Array())
		}
	}
	for i := range data.Spokes {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Spokes[i].Name.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "spokes").ForEach(
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
		data.Spokes[i].Name = types.StringNull()

		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "global" {
				data.Spokes[i].Name = types.StringValue(va.String())
			}
		}
		data.Spokes[i].SpokeSites = types.SetNull(types.StringType)

		if t := r.Get("spokeSites.optionType"); t.Exists() {
			va := r.Get("spokeSites.value")
			if t.String() == "global" {
				data.Spokes[i].SpokeSites = helpers.GetStringSet(va.Array())
			}
		}
		for ci := range data.Spokes[i].HubSites {
			keys := [...]string{"sites", "preference"}
			keyValues := [...]string{helpers.GetStringFromSet(data.Spokes[i].HubSites[ci].Sites).ValueString(), strconv.FormatInt(data.Spokes[i].HubSites[ci].Preference.ValueInt64(), 10)}
			keyValuesVariables := [...]string{"", ""}

			var cr gjson.Result
			r.Get("hubSites").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			data.Spokes[i].HubSites[ci].Sites = types.SetNull(types.StringType)

			if t := cr.Get("sites.optionType"); t.Exists() {
				va := cr.Get("sites.value")
				if t.String() == "global" {
					data.Spokes[i].HubSites[ci].Sites = helpers.GetStringSet(va.Array())
				}
			}
			data.Spokes[i].HubSites[ci].Preference = types.Int64Null()

			if t := cr.Get("preference.optionType"); t.Exists() {
				va := cr.Get("preference.value")
				if t.String() == "global" {
					data.Spokes[i].HubSites[ci].Preference = types.Int64Value(va.Int())
				}
			}
		}
	}
}

// End of section. //template:end updateFromBody
