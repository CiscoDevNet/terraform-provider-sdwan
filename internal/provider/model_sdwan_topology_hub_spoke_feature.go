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
func (data *TopologyHubSpoke) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
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
	oldSpokes := data.Spokes
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
	} else {
		data.Spokes = nil
	}
	if !fullRead && data.Spokes != nil {
		resultSpokes := make([]TopologyHubSpokeSpokes, 0, len(data.Spokes))
		matchedSpokes := make([]bool, len(data.Spokes))
		for _, oldItem := range oldSpokes {
			for ni := range data.Spokes {
				if matchedSpokes[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.Name.ValueString() != data.Spokes[ni].Name.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedSpokes[ni] = true
					if data.Spokes[ni].HubSites != nil {
						resultC := make([]TopologyHubSpokeSpokesHubSites, 0, len(data.Spokes[ni].HubSites))
						matchedC := make([]bool, len(data.Spokes[ni].HubSites))
						for _, oldCItem := range oldItem.HubSites {
							for nci := range data.Spokes[ni].HubSites {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.Sites).ValueString() != helpers.GetStringFromSet(data.Spokes[ni].HubSites[nci].Sites).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Preference.ValueInt64() != data.Spokes[ni].HubSites[nci].Preference.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.Spokes[ni].HubSites[nci])
									break
								}
							}
						}
						for nci := range data.Spokes[ni].HubSites {
							if !matchedC[nci] {
								resultC = append(resultC, data.Spokes[ni].HubSites[nci])
							}
						}
						data.Spokes[ni].HubSites = resultC
					}
					resultSpokes = append(resultSpokes, data.Spokes[ni])
					break
				}
			}
		}
		for ni := range data.Spokes {
			if !matchedSpokes[ni] {
				resultSpokes = append(resultSpokes, data.Spokes[ni])
			}
		}
		data.Spokes = resultSpokes
	}
}

// End of section. //template:end fromBody
