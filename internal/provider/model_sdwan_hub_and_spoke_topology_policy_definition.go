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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type HubAndSpokeTopologyPolicyDefinition struct {
	Id             types.String                                    `tfsdk:"id"`
	Version        types.Int64                                     `tfsdk:"version"`
	Type           types.String                                    `tfsdk:"type"`
	Name           types.String                                    `tfsdk:"name"`
	Description    types.String                                    `tfsdk:"description"`
	VpnListId      types.String                                    `tfsdk:"vpn_list_id"`
	VpnListVersion types.Int64                                     `tfsdk:"vpn_list_version"`
	Topologies     []HubAndSpokeTopologyPolicyDefinitionTopologies `tfsdk:"topologies"`
}

type HubAndSpokeTopologyPolicyDefinitionTopologies struct {
	Name              types.String                                          `tfsdk:"name"`
	AllHubsAreEqual   types.Bool                                            `tfsdk:"all_hubs_are_equal"`
	AdvertiseHubTlocs types.Bool                                            `tfsdk:"advertise_hub_tlocs"`
	TlocListId        types.String                                          `tfsdk:"tloc_list_id"`
	Spokes            []HubAndSpokeTopologyPolicyDefinitionTopologiesSpokes `tfsdk:"spokes"`
}

type HubAndSpokeTopologyPolicyDefinitionTopologiesSpokes struct {
	SiteListId      types.String                                              `tfsdk:"site_list_id"`
	SiteListVersion types.Int64                                               `tfsdk:"site_list_version"`
	Hubs            []HubAndSpokeTopologyPolicyDefinitionTopologiesSpokesHubs `tfsdk:"hubs"`
}

type HubAndSpokeTopologyPolicyDefinitionTopologiesSpokesHubs struct {
	SiteListId        types.String `tfsdk:"site_list_id"`
	SiteListVersion   types.Int64  `tfsdk:"site_list_version"`
	Preference        types.String `tfsdk:"preference"`
	Ipv4PrefixListIds types.Set    `tfsdk:"ipv4_prefix_list_ids"`
	Ipv6PrefixListIds types.Set    `tfsdk:"ipv6_prefix_list_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data HubAndSpokeTopologyPolicyDefinition) getPath() string {
	return "/template/policy/definition/hubandspoke/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data HubAndSpokeTopologyPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "hubAndSpoke")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.VpnListId.IsNull() {
		body, _ = sjson.Set(body, "definition.vpnList", data.VpnListId.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "definition.subDefinitions", []interface{}{})
		for _, item := range data.Topologies {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.AllHubsAreEqual.IsNull() {
				if false && item.AllHubsAreEqual.ValueBool() {
					itemBody, _ = sjson.Set(itemBody, "equalPreference", "")
				} else {
					itemBody, _ = sjson.Set(itemBody, "equalPreference", item.AllHubsAreEqual.ValueBool())
				}
			}
			if !item.AdvertiseHubTlocs.IsNull() {
				if false && item.AdvertiseHubTlocs.ValueBool() {
					itemBody, _ = sjson.Set(itemBody, "advertiseTloc", "")
				} else {
					itemBody, _ = sjson.Set(itemBody, "advertiseTloc", item.AdvertiseHubTlocs.ValueBool())
				}
			}
			if !item.TlocListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "tlocList", item.TlocListId.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "spokes", []interface{}{})
				for _, childItem := range item.Spokes {
					itemChildBody := ""
					if !childItem.SiteListId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "siteList", childItem.SiteListId.ValueString())
					}
					if true {
						itemChildBody, _ = sjson.Set(itemChildBody, "hubs", []interface{}{})
						for _, childChildItem := range childItem.Hubs {
							itemChildChildBody := ""
							if !childChildItem.SiteListId.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "siteList", childChildItem.SiteListId.ValueString())
							}
							if !childChildItem.Preference.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preference", childChildItem.Preference.ValueString())
							}
							if !childChildItem.Ipv4PrefixListIds.IsNull() {
								var values []string
								childChildItem.Ipv4PrefixListIds.ElementsAs(ctx, &values, false)
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "prefixLists", values)
							}
							if !childChildItem.Ipv6PrefixListIds.IsNull() {
								var values []string
								childChildItem.Ipv6PrefixListIds.ElementsAs(ctx, &values, false)
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "ipv6PrefixLists", values)
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "hubs.-1", itemChildChildBody)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "spokes.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "definition.subDefinitions.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *HubAndSpokeTopologyPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
	state := *data
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("definition.vpnList"); value.Exists() {
		data.VpnListId = types.StringValue(value.String())
	} else {
		data.VpnListId = types.StringNull()
	}
	if value := res.Get("definition.subDefinitions"); value.Exists() && len(value.Array()) > 0 {
		data.Topologies = make([]HubAndSpokeTopologyPolicyDefinitionTopologies, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := HubAndSpokeTopologyPolicyDefinitionTopologies{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("equalPreference"); cValue.Exists() {
				if false && cValue.String() == "" {
					item.AllHubsAreEqual = types.BoolValue(true)
				} else {
					item.AllHubsAreEqual = types.BoolValue(cValue.Bool())
				}
			} else {
				item.AllHubsAreEqual = types.BoolNull()
			}
			if cValue := v.Get("advertiseTloc"); cValue.Exists() {
				if false && cValue.String() == "" {
					item.AdvertiseHubTlocs = types.BoolValue(true)
				} else {
					item.AdvertiseHubTlocs = types.BoolValue(cValue.Bool())
				}
			} else {
				item.AdvertiseHubTlocs = types.BoolNull()
			}
			if cValue := v.Get("tlocList"); cValue.Exists() {
				item.TlocListId = types.StringValue(cValue.String())
			} else {
				item.TlocListId = types.StringNull()
			}
			if cValue := v.Get("spokes"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Spokes = make([]HubAndSpokeTopologyPolicyDefinitionTopologiesSpokes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := HubAndSpokeTopologyPolicyDefinitionTopologiesSpokes{}
					if ccValue := cv.Get("siteList"); ccValue.Exists() {
						cItem.SiteListId = types.StringValue(ccValue.String())
					} else {
						cItem.SiteListId = types.StringNull()
					}
					if ccValue := cv.Get("hubs"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.Hubs = make([]HubAndSpokeTopologyPolicyDefinitionTopologiesSpokesHubs, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := HubAndSpokeTopologyPolicyDefinitionTopologiesSpokesHubs{}
							if cccValue := ccv.Get("siteList"); cccValue.Exists() {
								ccItem.SiteListId = types.StringValue(cccValue.String())
							} else {
								ccItem.SiteListId = types.StringNull()
							}
							if cccValue := ccv.Get("preference"); cccValue.Exists() {
								ccItem.Preference = types.StringValue(cccValue.String())
							} else {
								ccItem.Preference = types.StringNull()
							}
							if cccValue := ccv.Get("prefixLists"); cccValue.Exists() {
								ccItem.Ipv4PrefixListIds = helpers.GetStringSet(cccValue.Array())
							} else {
								ccItem.Ipv4PrefixListIds = types.SetNull(types.StringType)
							}
							if cccValue := ccv.Get("ipv6PrefixLists"); cccValue.Exists() {
								ccItem.Ipv6PrefixListIds = helpers.GetStringSet(cccValue.Array())
							} else {
								ccItem.Ipv6PrefixListIds = types.SetNull(types.StringType)
							}
							cItem.Hubs = append(cItem.Hubs, ccItem)
							return true
						})
					} else {
						if len(cItem.Hubs) > 0 {
							cItem.Hubs = []HubAndSpokeTopologyPolicyDefinitionTopologiesSpokesHubs{}
						}
					}
					item.Spokes = append(item.Spokes, cItem)
					return true
				})
			} else {
				if len(item.Spokes) > 0 {
					item.Spokes = []HubAndSpokeTopologyPolicyDefinitionTopologiesSpokes{}
				}
			}
			data.Topologies = append(data.Topologies, item)
			return true
		})
	} else {
		if len(data.Topologies) > 0 {
			data.Topologies = []HubAndSpokeTopologyPolicyDefinitionTopologies{}
		}
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *HubAndSpokeTopologyPolicyDefinition) hasChanges(ctx context.Context, state *HubAndSpokeTopologyPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.VpnListId.Equal(state.VpnListId) {
		hasChanges = true
	}
	if len(data.Topologies) != len(state.Topologies) {
		hasChanges = true
	} else {
		for i := range data.Topologies {
			if !data.Topologies[i].Name.Equal(state.Topologies[i].Name) {
				hasChanges = true
			}
			if !data.Topologies[i].AllHubsAreEqual.Equal(state.Topologies[i].AllHubsAreEqual) {
				hasChanges = true
			}
			if !data.Topologies[i].AdvertiseHubTlocs.Equal(state.Topologies[i].AdvertiseHubTlocs) {
				hasChanges = true
			}
			if !data.Topologies[i].TlocListId.Equal(state.Topologies[i].TlocListId) {
				hasChanges = true
			}
			if len(data.Topologies[i].Spokes) != len(state.Topologies[i].Spokes) {
				hasChanges = true
			} else {
				for ii := range data.Topologies[i].Spokes {
					if !data.Topologies[i].Spokes[ii].SiteListId.Equal(state.Topologies[i].Spokes[ii].SiteListId) {
						hasChanges = true
					}
					if len(data.Topologies[i].Spokes[ii].Hubs) != len(state.Topologies[i].Spokes[ii].Hubs) {
						hasChanges = true
					} else {
						for iii := range data.Topologies[i].Spokes[ii].Hubs {
							if !data.Topologies[i].Spokes[ii].Hubs[iii].SiteListId.Equal(state.Topologies[i].Spokes[ii].Hubs[iii].SiteListId) {
								hasChanges = true
							}
							if !data.Topologies[i].Spokes[ii].Hubs[iii].Preference.Equal(state.Topologies[i].Spokes[ii].Hubs[iii].Preference) {
								hasChanges = true
							}
							if !data.Topologies[i].Spokes[ii].Hubs[iii].Ipv4PrefixListIds.Equal(state.Topologies[i].Spokes[ii].Hubs[iii].Ipv4PrefixListIds) {
								hasChanges = true
							}
							if !data.Topologies[i].Spokes[ii].Hubs[iii].Ipv6PrefixListIds.Equal(state.Topologies[i].Spokes[ii].Hubs[iii].Ipv6PrefixListIds) {
								hasChanges = true
							}
						}
					}
				}
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *HubAndSpokeTopologyPolicyDefinition) updateVersions(ctx context.Context, state *HubAndSpokeTopologyPolicyDefinition) {
	data.VpnListVersion = state.VpnListVersion
	for i := range data.Topologies {
		dataKeys := [...]string{fmt.Sprintf("%v", data.Topologies[i].Name.ValueString())}
		stateIndex := -1
		for j := range state.Topologies {
			stateKeys := [...]string{fmt.Sprintf("%v", state.Topologies[j].Name.ValueString())}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		for ii := range data.Topologies[i].Spokes {
			cDataKeys := [...]string{fmt.Sprintf("%v", data.Topologies[i].Spokes[ii].SiteListId.ValueString())}
			cStateIndex := -1
			if stateIndex > -1 {
				for jj := range state.Topologies[stateIndex].Spokes {
					cStateKeys := [...]string{fmt.Sprintf("%v", state.Topologies[stateIndex].Spokes[jj].SiteListId.ValueString())}
					if cDataKeys == cStateKeys {
						cStateIndex = jj
						break
					}
				}
			}
			if cStateIndex > -1 {
				data.Topologies[i].Spokes[ii].SiteListVersion = state.Topologies[stateIndex].Spokes[cStateIndex].SiteListVersion
			} else {
				data.Topologies[i].Spokes[ii].SiteListVersion = types.Int64Null()
			}
			for iii := range data.Topologies[i].Spokes[ii].Hubs {
				ccDataKeys := [...]string{fmt.Sprintf("%v", data.Topologies[i].Spokes[ii].Hubs[iii].SiteListId.ValueString())}
				ccStateIndex := -1
				if cStateIndex > -1 {
					for jjj := range state.Topologies[stateIndex].Spokes[cStateIndex].Hubs {
						ccStateKeys := [...]string{fmt.Sprintf("%v", state.Topologies[stateIndex].Spokes[cStateIndex].Hubs[jjj].SiteListId.ValueString())}
						if ccDataKeys == ccStateKeys {
							ccStateIndex = jjj
							break
						}
					}
				}
				if ccStateIndex > -1 {
					data.Topologies[i].Spokes[ii].Hubs[iii].SiteListVersion = state.Topologies[stateIndex].Spokes[cStateIndex].Hubs[ccStateIndex].SiteListVersion
				} else {
					data.Topologies[i].Spokes[ii].Hubs[iii].SiteListVersion = types.Int64Null()
				}
			}
		}
	}
}

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *HubAndSpokeTopologyPolicyDefinition) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
	data.Type = types.StringValue("hubAndSpoke")
	if data.VpnListId != types.StringNull() {
		data.VpnListVersion = types.Int64Value(0)
	}
	for i := range data.Topologies {
		for ii := range data.Topologies[i].Spokes {
			if data.Topologies[i].Spokes[ii].SiteListId != types.StringNull() {
				data.Topologies[i].Spokes[ii].SiteListVersion = types.Int64Value(0)
			}
			for iii := range data.Topologies[i].Spokes[ii].Hubs {
				if data.Topologies[i].Spokes[ii].Hubs[iii].SiteListId != types.StringNull() {
					data.Topologies[i].Spokes[ii].Hubs[iii].SiteListVersion = types.Int64Value(0)
				}
			}
		}
	}
}

// End of section. //template:end processImport
