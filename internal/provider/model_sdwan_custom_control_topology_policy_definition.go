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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type CustomControlTopologyPolicyDefinition struct {
	Id            types.String                                     `tfsdk:"id"`
	Version       types.Int64                                      `tfsdk:"version"`
	Type          types.String                                     `tfsdk:"type"`
	Name          types.String                                     `tfsdk:"name"`
	Description   types.String                                     `tfsdk:"description"`
	DefaultAction types.String                                     `tfsdk:"default_action"`
	Sequences     []CustomControlTopologyPolicyDefinitionSequences `tfsdk:"sequences"`
}

type CustomControlTopologyPolicyDefinitionSequences struct {
	Id            types.Int64                                                   `tfsdk:"id"`
	Name          types.String                                                  `tfsdk:"name"`
	Type          types.String                                                  `tfsdk:"type"`
	IpType        types.String                                                  `tfsdk:"ip_type"`
	BaseAction    types.String                                                  `tfsdk:"base_action"`
	MatchEntries  []CustomControlTopologyPolicyDefinitionSequencesMatchEntries  `tfsdk:"match_entries"`
	ActionEntries []CustomControlTopologyPolicyDefinitionSequencesActionEntries `tfsdk:"action_entries"`
}

type CustomControlTopologyPolicyDefinitionSequencesMatchEntries struct {
	Type                         types.String `tfsdk:"type"`
	ColorListId                  types.String `tfsdk:"color_list_id"`
	ColorListVersion             types.Int64  `tfsdk:"color_list_version"`
	CommunityListId              types.String `tfsdk:"community_list_id"`
	CommunityListVersion         types.Int64  `tfsdk:"community_list_version"`
	ExpandedCommunityListId      types.String `tfsdk:"expanded_community_list_id"`
	ExpandedCommunityListVersion types.Int64  `tfsdk:"expanded_community_list_version"`
	OmpTag                       types.Int64  `tfsdk:"omp_tag"`
	Origin                       types.String `tfsdk:"origin"`
	Originator                   types.String `tfsdk:"originator"`
	Preference                   types.Int64  `tfsdk:"preference"`
	SiteListId                   types.String `tfsdk:"site_list_id"`
	SiteListVersion              types.Int64  `tfsdk:"site_list_version"`
	PathType                     types.String `tfsdk:"path_type"`
	TlocListId                   types.String `tfsdk:"tloc_list_id"`
	TlocListVersion              types.Int64  `tfsdk:"tloc_list_version"`
	VpnListId                    types.String `tfsdk:"vpn_list_id"`
	VpnListVersion               types.Int64  `tfsdk:"vpn_list_version"`
	PrefixListId                 types.String `tfsdk:"prefix_list_id"`
	PrefixListVersion            types.Int64  `tfsdk:"prefix_list_version"`
	VpnId                        types.Int64  `tfsdk:"vpn_id"`
	TlocIp                       types.String `tfsdk:"tloc_ip"`
	TlocColor                    types.String `tfsdk:"tloc_color"`
	TlocEncapsulation            types.String `tfsdk:"tloc_encapsulation"`
	SiteId                       types.Int64  `tfsdk:"site_id"`
	Carrier                      types.String `tfsdk:"carrier"`
	DomainId                     types.Int64  `tfsdk:"domain_id"`
	GroupId                      types.Int64  `tfsdk:"group_id"`
	RegionId                     types.Int64  `tfsdk:"region_id"`
	Role                         types.String `tfsdk:"role"`
	RegionListId                 types.String `tfsdk:"region_list_id"`
}
type CustomControlTopologyPolicyDefinitionSequencesActionEntries struct {
	Type                   types.String                                                               `tfsdk:"type"`
	SetParameters          []CustomControlTopologyPolicyDefinitionSequencesActionEntriesSetParameters `tfsdk:"set_parameters"`
	ExportToVpnListId      types.String                                                               `tfsdk:"export_to_vpn_list_id"`
	ExportToVpnListVersion types.Int64                                                                `tfsdk:"export_to_vpn_list_version"`
}

type CustomControlTopologyPolicyDefinitionSequencesActionEntriesSetParameters struct {
	Type                     types.String `tfsdk:"type"`
	TlocListId               types.String `tfsdk:"tloc_list_id"`
	TlocListVersion          types.Int64  `tfsdk:"tloc_list_version"`
	TlocIp                   types.String `tfsdk:"tloc_ip"`
	TlocColor                types.String `tfsdk:"tloc_color"`
	TlocEncapsulation        types.String `tfsdk:"tloc_encapsulation"`
	TlocAction               types.String `tfsdk:"tloc_action"`
	Preference               types.Int64  `tfsdk:"preference"`
	OmpTag                   types.Int64  `tfsdk:"omp_tag"`
	Community                types.String `tfsdk:"community"`
	CommunityAdditive        types.Bool   `tfsdk:"community_additive"`
	ServiceType              types.String `tfsdk:"service_type"`
	ServiceVpnId             types.Int64  `tfsdk:"service_vpn_id"`
	ServiceTlocListId        types.String `tfsdk:"service_tloc_list_id"`
	ServiceTlocListVersion   types.Int64  `tfsdk:"service_tloc_list_version"`
	ServiceTlocIp            types.String `tfsdk:"service_tloc_ip"`
	ServiceTlocColor         types.String `tfsdk:"service_tloc_color"`
	ServiceTlocEncapsulation types.String `tfsdk:"service_tloc_encapsulation"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data CustomControlTopologyPolicyDefinition) getPath() string {
	return "/template/policy/definition/control/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CustomControlTopologyPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "control")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.DefaultAction.IsNull() {
		body, _ = sjson.Set(body, "defaultAction.type", data.DefaultAction.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "sequences", []interface{}{})
		for _, item := range data.Sequences {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceId", item.Id.ValueInt64())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceName", item.Name.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceType", item.Type.ValueString())
			}
			if !item.IpType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceIpType", item.IpType.ValueString())
			}
			if !item.BaseAction.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "baseAction", item.BaseAction.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "match.entries", []interface{}{})
				for _, childItem := range item.MatchEntries {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "field", childItem.Type.ValueString())
					}
					if !childItem.ColorListId.IsNull() && childItem.Type.ValueString() == "colorList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.ColorListId.ValueString())
					}
					if !childItem.CommunityListId.IsNull() && childItem.Type.ValueString() == "community" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.CommunityListId.ValueString())
					}
					if !childItem.ExpandedCommunityListId.IsNull() && childItem.Type.ValueString() == "expandedCommunity" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.ExpandedCommunityListId.ValueString())
					}
					if !childItem.OmpTag.IsNull() && childItem.Type.ValueString() == "ompTag" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.OmpTag.ValueInt64()))
					}
					if !childItem.Origin.IsNull() && childItem.Type.ValueString() == "origin" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Origin.ValueString())
					}
					if !childItem.Originator.IsNull() && childItem.Type.ValueString() == "originator" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Originator.ValueString())
					}
					if !childItem.Preference.IsNull() && childItem.Type.ValueString() == "preference" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.Preference.ValueInt64()))
					}
					if !childItem.SiteListId.IsNull() && childItem.Type.ValueString() == "siteList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.SiteListId.ValueString())
					}
					if !childItem.PathType.IsNull() && childItem.Type.ValueString() == "pathType" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.PathType.ValueString())
					}
					if !childItem.TlocListId.IsNull() && childItem.Type.ValueString() == "tlocList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.TlocListId.ValueString())
					}
					if !childItem.VpnListId.IsNull() && childItem.Type.ValueString() == "vpnList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.VpnListId.ValueString())
					}
					if !childItem.PrefixListId.IsNull() && childItem.Type.ValueString() == "prefixList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.PrefixListId.ValueString())
					}
					if !childItem.VpnId.IsNull() && childItem.Type.ValueString() == "vpn" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.VpnId.ValueInt64()))
					}
					if !childItem.TlocIp.IsNull() && childItem.Type.ValueString() == "tloc" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.ip", childItem.TlocIp.ValueString())
					}
					if !childItem.TlocColor.IsNull() && childItem.Type.ValueString() == "tloc" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.color", childItem.TlocColor.ValueString())
					}
					if !childItem.TlocEncapsulation.IsNull() && childItem.Type.ValueString() == "tloc" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.encap", childItem.TlocEncapsulation.ValueString())
					}
					if !childItem.SiteId.IsNull() && childItem.Type.ValueString() == "siteId" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.SiteId.ValueInt64()))
					}
					if !childItem.Carrier.IsNull() && childItem.Type.ValueString() == "carrier" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Carrier.ValueString())
					}
					if !childItem.DomainId.IsNull() && childItem.Type.ValueString() == "domainId" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.DomainId.ValueInt64()))
					}
					if !childItem.GroupId.IsNull() && childItem.Type.ValueString() == "groupId" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.GroupId.ValueInt64()))
					}
					if !childItem.RegionId.IsNull() && childItem.Type.ValueString() == "regionId" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.RegionId.ValueInt64()))
					}
					if !childItem.Role.IsNull() && childItem.Type.ValueString() == "regionId" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Role.ValueString())
					}
					if !childItem.RegionListId.IsNull() && childItem.Type.ValueString() == "regionList" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.RegionListId.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "match.entries.-1", itemChildBody)
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "actions", []interface{}{})
				for _, childItem := range item.ActionEntries {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if true && childItem.Type.ValueString() == "set" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter", []interface{}{})
						for _, childChildItem := range childItem.SetParameters {
							itemChildChildBody := ""
							if !childChildItem.Type.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "field", childChildItem.Type.ValueString())
							}
							if !childChildItem.TlocListId.IsNull() && childChildItem.Type.ValueString() == "tlocList" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "ref", childChildItem.TlocListId.ValueString())
							}
							if !childChildItem.TlocIp.IsNull() && childChildItem.Type.ValueString() == "tloc" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.ip", childChildItem.TlocIp.ValueString())
							}
							if !childChildItem.TlocColor.IsNull() && childChildItem.Type.ValueString() == "tloc" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.color", childChildItem.TlocColor.ValueString())
							}
							if !childChildItem.TlocEncapsulation.IsNull() && childChildItem.Type.ValueString() == "tloc" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.encap", childChildItem.TlocEncapsulation.ValueString())
							}
							if !childChildItem.TlocAction.IsNull() && childChildItem.Type.ValueString() == "tlocAction" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", childChildItem.TlocAction.ValueString())
							}
							if !childChildItem.Preference.IsNull() && childChildItem.Type.ValueString() == "preference" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", fmt.Sprint(childChildItem.Preference.ValueInt64()))
							}
							if !childChildItem.OmpTag.IsNull() && childChildItem.Type.ValueString() == "ompTag" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", fmt.Sprint(childChildItem.OmpTag.ValueInt64()))
							}
							if !childChildItem.Community.IsNull() && childChildItem.Type.ValueString() == "community" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", childChildItem.Community.ValueString())
							}
							if !childChildItem.CommunityAdditive.IsNull() && childChildItem.Type.ValueString() == "communityAdditive" {
								if false && childChildItem.CommunityAdditive.ValueBool() {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", "")
								} else {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value", fmt.Sprint(childChildItem.CommunityAdditive.ValueBool()))
								}
							}
							if !childChildItem.ServiceType.IsNull() && childChildItem.Type.ValueString() == "service" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.type", childChildItem.ServiceType.ValueString())
							}
							if !childChildItem.ServiceVpnId.IsNull() && childChildItem.Type.ValueString() == "service" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.vpn", fmt.Sprint(childChildItem.ServiceVpnId.ValueInt64()))
							}
							if !childChildItem.ServiceTlocListId.IsNull() && childChildItem.Type.ValueString() == "service" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.tlocList", childChildItem.ServiceTlocListId.ValueString())
							}
							if !childChildItem.ServiceTlocIp.IsNull() && childChildItem.Type.ValueString() == "service" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.tloc.ip", childChildItem.ServiceTlocIp.ValueString())
							}
							if !childChildItem.ServiceTlocColor.IsNull() && childChildItem.Type.ValueString() == "service" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.tloc.color", childChildItem.ServiceTlocColor.ValueString())
							}
							if !childChildItem.ServiceTlocEncapsulation.IsNull() && childChildItem.Type.ValueString() == "service" {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "value.tloc.encap", childChildItem.ServiceTlocEncapsulation.ValueString())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "parameter.-1", itemChildChildBody)
						}
					}
					if true && childItem.Type.ValueString() == "exportTo" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter.field", "vpnList")
					}
					if !childItem.ExportToVpnListId.IsNull() && childItem.Type.ValueString() == "exportTo" {
						itemChildBody, _ = sjson.Set(itemChildBody, "parameter.ref", childItem.ExportToVpnListId.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "actions.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "sequences.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CustomControlTopologyPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("defaultAction.type"); value.Exists() {
		data.DefaultAction = types.StringValue(value.String())
	} else {
		data.DefaultAction = types.StringNull()
	}
	if value := res.Get("sequences"); value.Exists() && len(value.Array()) > 0 {
		data.Sequences = make([]CustomControlTopologyPolicyDefinitionSequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CustomControlTopologyPolicyDefinitionSequences{}
			if cValue := v.Get("sequenceId"); cValue.Exists() {
				item.Id = types.Int64Value(cValue.Int())
			} else {
				item.Id = types.Int64Null()
			}
			if cValue := v.Get("sequenceName"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("sequenceType"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			} else {
				item.Type = types.StringNull()
			}
			if cValue := v.Get("sequenceIpType"); cValue.Exists() {
				item.IpType = types.StringValue(cValue.String())
			} else {
				item.IpType = types.StringNull()
			}
			if cValue := v.Get("baseAction"); cValue.Exists() {
				item.BaseAction = types.StringValue(cValue.String())
			} else {
				item.BaseAction = types.StringNull()
			}
			if cValue := v.Get("match.entries"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.MatchEntries = make([]CustomControlTopologyPolicyDefinitionSequencesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CustomControlTopologyPolicyDefinitionSequencesMatchEntries{}
					if ccValue := cv.Get("field"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "colorList" {
						cItem.ColorListId = types.StringValue(ccValue.String())
					} else {
						cItem.ColorListId = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "community" {
						cItem.CommunityListId = types.StringValue(ccValue.String())
					} else {
						cItem.CommunityListId = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "expandedCommunity" {
						cItem.ExpandedCommunityListId = types.StringValue(ccValue.String())
					} else {
						cItem.ExpandedCommunityListId = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "ompTag" {
						cItem.OmpTag = types.Int64Value(ccValue.Int())
					} else {
						cItem.OmpTag = types.Int64Null()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "origin" {
						cItem.Origin = types.StringValue(ccValue.String())
					} else {
						cItem.Origin = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "originator" {
						cItem.Originator = types.StringValue(ccValue.String())
					} else {
						cItem.Originator = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "preference" {
						cItem.Preference = types.Int64Value(ccValue.Int())
					} else {
						cItem.Preference = types.Int64Null()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "siteList" {
						cItem.SiteListId = types.StringValue(ccValue.String())
					} else {
						cItem.SiteListId = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "pathType" {
						cItem.PathType = types.StringValue(ccValue.String())
					} else {
						cItem.PathType = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "tlocList" {
						cItem.TlocListId = types.StringValue(ccValue.String())
					} else {
						cItem.TlocListId = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "vpnList" {
						cItem.VpnListId = types.StringValue(ccValue.String())
					} else {
						cItem.VpnListId = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "prefixList" {
						cItem.PrefixListId = types.StringValue(ccValue.String())
					} else {
						cItem.PrefixListId = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "vpn" {
						cItem.VpnId = types.Int64Value(ccValue.Int())
					} else {
						cItem.VpnId = types.Int64Null()
					}
					if ccValue := cv.Get("value.ip"); ccValue.Exists() && cItem.Type.ValueString() == "tloc" {
						cItem.TlocIp = types.StringValue(ccValue.String())
					} else {
						cItem.TlocIp = types.StringNull()
					}
					if ccValue := cv.Get("value.color"); ccValue.Exists() && cItem.Type.ValueString() == "tloc" {
						cItem.TlocColor = types.StringValue(ccValue.String())
					} else {
						cItem.TlocColor = types.StringNull()
					}
					if ccValue := cv.Get("value.encap"); ccValue.Exists() && cItem.Type.ValueString() == "tloc" {
						cItem.TlocEncapsulation = types.StringValue(ccValue.String())
					} else {
						cItem.TlocEncapsulation = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "siteId" {
						cItem.SiteId = types.Int64Value(ccValue.Int())
					} else {
						cItem.SiteId = types.Int64Null()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "carrier" {
						cItem.Carrier = types.StringValue(ccValue.String())
					} else {
						cItem.Carrier = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "domainId" {
						cItem.DomainId = types.Int64Value(ccValue.Int())
					} else {
						cItem.DomainId = types.Int64Null()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "groupId" {
						cItem.GroupId = types.Int64Value(ccValue.Int())
					} else {
						cItem.GroupId = types.Int64Null()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "regionId" {
						cItem.RegionId = types.Int64Value(ccValue.Int())
					} else {
						cItem.RegionId = types.Int64Null()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "regionId" {
						cItem.Role = types.StringValue(ccValue.String())
					} else {
						cItem.Role = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "regionList" {
						cItem.RegionListId = types.StringValue(ccValue.String())
					} else {
						cItem.RegionListId = types.StringNull()
					}
					item.MatchEntries = append(item.MatchEntries, cItem)
					return true
				})
			} else {
				if len(item.MatchEntries) > 0 {
					item.MatchEntries = []CustomControlTopologyPolicyDefinitionSequencesMatchEntries{}
				}
			}
			if cValue := v.Get("actions"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ActionEntries = make([]CustomControlTopologyPolicyDefinitionSequencesActionEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CustomControlTopologyPolicyDefinitionSequencesActionEntries{}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("parameter"); ccValue.Exists() && len(ccValue.Array()) > 0 && cItem.Type.ValueString() == "set" {
						cItem.SetParameters = make([]CustomControlTopologyPolicyDefinitionSequencesActionEntriesSetParameters, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := CustomControlTopologyPolicyDefinitionSequencesActionEntriesSetParameters{}
							if cccValue := ccv.Get("field"); cccValue.Exists() {
								ccItem.Type = types.StringValue(cccValue.String())
							} else {
								ccItem.Type = types.StringNull()
							}
							if cccValue := ccv.Get("ref"); cccValue.Exists() && ccItem.Type.ValueString() == "tlocList" {
								ccItem.TlocListId = types.StringValue(cccValue.String())
							} else {
								ccItem.TlocListId = types.StringNull()
							}
							if cccValue := ccv.Get("value.ip"); cccValue.Exists() && ccItem.Type.ValueString() == "tloc" {
								ccItem.TlocIp = types.StringValue(cccValue.String())
							} else {
								ccItem.TlocIp = types.StringNull()
							}
							if cccValue := ccv.Get("value.color"); cccValue.Exists() && ccItem.Type.ValueString() == "tloc" {
								ccItem.TlocColor = types.StringValue(cccValue.String())
							} else {
								ccItem.TlocColor = types.StringNull()
							}
							if cccValue := ccv.Get("value.encap"); cccValue.Exists() && ccItem.Type.ValueString() == "tloc" {
								ccItem.TlocEncapsulation = types.StringValue(cccValue.String())
							} else {
								ccItem.TlocEncapsulation = types.StringNull()
							}
							if cccValue := ccv.Get("value"); cccValue.Exists() && ccItem.Type.ValueString() == "tlocAction" {
								ccItem.TlocAction = types.StringValue(cccValue.String())
							} else {
								ccItem.TlocAction = types.StringNull()
							}
							if cccValue := ccv.Get("value"); cccValue.Exists() && ccItem.Type.ValueString() == "preference" {
								ccItem.Preference = types.Int64Value(cccValue.Int())
							} else {
								ccItem.Preference = types.Int64Null()
							}
							if cccValue := ccv.Get("value"); cccValue.Exists() && ccItem.Type.ValueString() == "ompTag" {
								ccItem.OmpTag = types.Int64Value(cccValue.Int())
							} else {
								ccItem.OmpTag = types.Int64Null()
							}
							if cccValue := ccv.Get("value"); cccValue.Exists() && ccItem.Type.ValueString() == "community" {
								ccItem.Community = types.StringValue(cccValue.String())
							} else {
								ccItem.Community = types.StringNull()
							}
							if cccValue := ccv.Get("value"); cccValue.Exists() && ccItem.Type.ValueString() == "communityAdditive" {
								if false && cccValue.String() == "" {
									ccItem.CommunityAdditive = types.BoolValue(true)
								} else {
									ccItem.CommunityAdditive = types.BoolValue(cccValue.Bool())
								}
							} else {
								ccItem.CommunityAdditive = types.BoolNull()
							}
							if cccValue := ccv.Get("value.type"); cccValue.Exists() && ccItem.Type.ValueString() == "service" {
								ccItem.ServiceType = types.StringValue(cccValue.String())
							} else {
								ccItem.ServiceType = types.StringNull()
							}
							if cccValue := ccv.Get("value.vpn"); cccValue.Exists() && ccItem.Type.ValueString() == "service" {
								ccItem.ServiceVpnId = types.Int64Value(cccValue.Int())
							} else {
								ccItem.ServiceVpnId = types.Int64Null()
							}
							if cccValue := ccv.Get("value.tlocList"); cccValue.Exists() && ccItem.Type.ValueString() == "service" {
								ccItem.ServiceTlocListId = types.StringValue(cccValue.String())
							} else {
								ccItem.ServiceTlocListId = types.StringNull()
							}
							if cccValue := ccv.Get("value.tloc.ip"); cccValue.Exists() && ccItem.Type.ValueString() == "service" {
								ccItem.ServiceTlocIp = types.StringValue(cccValue.String())
							} else {
								ccItem.ServiceTlocIp = types.StringNull()
							}
							if cccValue := ccv.Get("value.tloc.color"); cccValue.Exists() && ccItem.Type.ValueString() == "service" {
								ccItem.ServiceTlocColor = types.StringValue(cccValue.String())
							} else {
								ccItem.ServiceTlocColor = types.StringNull()
							}
							if cccValue := ccv.Get("value.tloc.encap"); cccValue.Exists() && ccItem.Type.ValueString() == "service" {
								ccItem.ServiceTlocEncapsulation = types.StringValue(cccValue.String())
							} else {
								ccItem.ServiceTlocEncapsulation = types.StringNull()
							}
							cItem.SetParameters = append(cItem.SetParameters, ccItem)
							return true
						})
					} else {
						if len(cItem.SetParameters) > 0 {
							cItem.SetParameters = []CustomControlTopologyPolicyDefinitionSequencesActionEntriesSetParameters{}
						}
					}
					if ccValue := cv.Get("parameter.ref"); ccValue.Exists() && cItem.Type.ValueString() == "exportTo" {
						cItem.ExportToVpnListId = types.StringValue(ccValue.String())
					} else {
						cItem.ExportToVpnListId = types.StringNull()
					}
					item.ActionEntries = append(item.ActionEntries, cItem)
					return true
				})
			} else {
				if len(item.ActionEntries) > 0 {
					item.ActionEntries = []CustomControlTopologyPolicyDefinitionSequencesActionEntries{}
				}
			}
			data.Sequences = append(data.Sequences, item)
			return true
		})
	} else {
		if len(data.Sequences) > 0 {
			data.Sequences = []CustomControlTopologyPolicyDefinitionSequences{}
		}
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CustomControlTopologyPolicyDefinition) hasChanges(ctx context.Context, state *CustomControlTopologyPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.DefaultAction.Equal(state.DefaultAction) {
		hasChanges = true
	}
	if len(data.Sequences) != len(state.Sequences) {
		hasChanges = true
	} else {
		for i := range data.Sequences {
			if !data.Sequences[i].Id.Equal(state.Sequences[i].Id) {
				hasChanges = true
			}
			if !data.Sequences[i].Name.Equal(state.Sequences[i].Name) {
				hasChanges = true
			}
			if !data.Sequences[i].Type.Equal(state.Sequences[i].Type) {
				hasChanges = true
			}
			if !data.Sequences[i].IpType.Equal(state.Sequences[i].IpType) {
				hasChanges = true
			}
			if !data.Sequences[i].BaseAction.Equal(state.Sequences[i].BaseAction) {
				hasChanges = true
			}
			if len(data.Sequences[i].MatchEntries) != len(state.Sequences[i].MatchEntries) {
				hasChanges = true
			} else {
				for ii := range data.Sequences[i].MatchEntries {
					if !data.Sequences[i].MatchEntries[ii].Type.Equal(state.Sequences[i].MatchEntries[ii].Type) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].ColorListId.Equal(state.Sequences[i].MatchEntries[ii].ColorListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].CommunityListId.Equal(state.Sequences[i].MatchEntries[ii].CommunityListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].ExpandedCommunityListId.Equal(state.Sequences[i].MatchEntries[ii].ExpandedCommunityListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].OmpTag.Equal(state.Sequences[i].MatchEntries[ii].OmpTag) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Origin.Equal(state.Sequences[i].MatchEntries[ii].Origin) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Originator.Equal(state.Sequences[i].MatchEntries[ii].Originator) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Preference.Equal(state.Sequences[i].MatchEntries[ii].Preference) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].SiteListId.Equal(state.Sequences[i].MatchEntries[ii].SiteListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].PathType.Equal(state.Sequences[i].MatchEntries[ii].PathType) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].TlocListId.Equal(state.Sequences[i].MatchEntries[ii].TlocListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].VpnListId.Equal(state.Sequences[i].MatchEntries[ii].VpnListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].PrefixListId.Equal(state.Sequences[i].MatchEntries[ii].PrefixListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].VpnId.Equal(state.Sequences[i].MatchEntries[ii].VpnId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].TlocIp.Equal(state.Sequences[i].MatchEntries[ii].TlocIp) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].TlocColor.Equal(state.Sequences[i].MatchEntries[ii].TlocColor) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].TlocEncapsulation.Equal(state.Sequences[i].MatchEntries[ii].TlocEncapsulation) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].SiteId.Equal(state.Sequences[i].MatchEntries[ii].SiteId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Carrier.Equal(state.Sequences[i].MatchEntries[ii].Carrier) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].DomainId.Equal(state.Sequences[i].MatchEntries[ii].DomainId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].GroupId.Equal(state.Sequences[i].MatchEntries[ii].GroupId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].RegionId.Equal(state.Sequences[i].MatchEntries[ii].RegionId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Role.Equal(state.Sequences[i].MatchEntries[ii].Role) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].RegionListId.Equal(state.Sequences[i].MatchEntries[ii].RegionListId) {
						hasChanges = true
					}
				}
			}
			if len(data.Sequences[i].ActionEntries) != len(state.Sequences[i].ActionEntries) {
				hasChanges = true
			} else {
				for ii := range data.Sequences[i].ActionEntries {
					if !data.Sequences[i].ActionEntries[ii].Type.Equal(state.Sequences[i].ActionEntries[ii].Type) {
						hasChanges = true
					}
					if len(data.Sequences[i].ActionEntries[ii].SetParameters) != len(state.Sequences[i].ActionEntries[ii].SetParameters) {
						hasChanges = true
					} else {
						for iii := range data.Sequences[i].ActionEntries[ii].SetParameters {
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].Type.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].Type) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocListId.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocListId) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocIp.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocIp) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocColor.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocColor) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocEncapsulation.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocEncapsulation) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocAction.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocAction) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].Preference.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].Preference) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].OmpTag.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].OmpTag) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].Community.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].Community) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].CommunityAdditive.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].CommunityAdditive) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceType.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceType) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceVpnId.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceVpnId) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocListId.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocListId) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocIp.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocIp) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocColor.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocColor) {
								hasChanges = true
							}
							if !data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocEncapsulation.Equal(state.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocEncapsulation) {
								hasChanges = true
							}
						}
					}
					if !data.Sequences[i].ActionEntries[ii].ExportToVpnListId.Equal(state.Sequences[i].ActionEntries[ii].ExportToVpnListId) {
						hasChanges = true
					}
				}
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *CustomControlTopologyPolicyDefinition) updateVersions(ctx context.Context, state *CustomControlTopologyPolicyDefinition) {
	for i := range data.Sequences {
		dataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].Id.ValueInt64()), fmt.Sprintf("%v", data.Sequences[i].Name.ValueString())}
		stateIndex := -1
		for j := range state.Sequences {
			stateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[j].Id.ValueInt64()), fmt.Sprintf("%v", state.Sequences[j].Name.ValueString())}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		for ii := range data.Sequences[i].MatchEntries {
			cDataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].MatchEntries[ii].Type.ValueString())}
			cStateIndex := -1
			if stateIndex > -1 {
				for jj := range state.Sequences[stateIndex].MatchEntries {
					cStateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[stateIndex].MatchEntries[jj].Type.ValueString())}
					if cDataKeys == cStateKeys {
						cStateIndex = jj
						break
					}
				}
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].ColorListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].ColorListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].ColorListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].CommunityListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].CommunityListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].CommunityListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].ExpandedCommunityListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].ExpandedCommunityListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].ExpandedCommunityListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].SiteListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].SiteListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].SiteListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].TlocListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].TlocListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].TlocListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].VpnListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].VpnListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].VpnListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].PrefixListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].PrefixListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].PrefixListVersion = types.Int64Null()
			}
		}
		for ii := range data.Sequences[i].ActionEntries {
			cDataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].ActionEntries[ii].Type.ValueString())}
			cStateIndex := -1
			if stateIndex > -1 {
				for jj := range state.Sequences[stateIndex].ActionEntries {
					cStateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[stateIndex].ActionEntries[jj].Type.ValueString())}
					if cDataKeys == cStateKeys {
						cStateIndex = jj
						break
					}
				}
			}
			for iii := range data.Sequences[i].ActionEntries[ii].SetParameters {
				ccDataKeys := [...]string{fmt.Sprintf("%v", data.Sequences[i].ActionEntries[ii].SetParameters[iii].Type.ValueString())}
				ccStateIndex := -1
				if cStateIndex > -1 {
					for jjj := range state.Sequences[stateIndex].ActionEntries[cStateIndex].SetParameters {
						ccStateKeys := [...]string{fmt.Sprintf("%v", state.Sequences[stateIndex].ActionEntries[cStateIndex].SetParameters[jjj].Type.ValueString())}
						if ccDataKeys == ccStateKeys {
							ccStateIndex = jjj
							break
						}
					}
				}
				if ccStateIndex > -1 {
					data.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocListVersion = state.Sequences[stateIndex].ActionEntries[cStateIndex].SetParameters[ccStateIndex].TlocListVersion
				} else {
					data.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocListVersion = types.Int64Null()
				}
				if ccStateIndex > -1 {
					data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocListVersion = state.Sequences[stateIndex].ActionEntries[cStateIndex].SetParameters[ccStateIndex].ServiceTlocListVersion
				} else {
					data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocListVersion = types.Int64Null()
				}
			}
			if cStateIndex > -1 {
				data.Sequences[i].ActionEntries[ii].ExportToVpnListVersion = state.Sequences[stateIndex].ActionEntries[cStateIndex].ExportToVpnListVersion
			} else {
				data.Sequences[i].ActionEntries[ii].ExportToVpnListVersion = types.Int64Null()
			}
		}
	}
}

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *CustomControlTopologyPolicyDefinition) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
	data.Type = types.StringValue("control")
	for i := range data.Sequences {
		for ii := range data.Sequences[i].MatchEntries {
			if data.Sequences[i].MatchEntries[ii].ColorListId != types.StringNull() {
				data.Sequences[i].MatchEntries[ii].ColorListVersion = types.Int64Value(0)
			}
			if data.Sequences[i].MatchEntries[ii].CommunityListId != types.StringNull() {
				data.Sequences[i].MatchEntries[ii].CommunityListVersion = types.Int64Value(0)
			}
			if data.Sequences[i].MatchEntries[ii].ExpandedCommunityListId != types.StringNull() {
				data.Sequences[i].MatchEntries[ii].ExpandedCommunityListVersion = types.Int64Value(0)
			}
			if data.Sequences[i].MatchEntries[ii].SiteListId != types.StringNull() {
				data.Sequences[i].MatchEntries[ii].SiteListVersion = types.Int64Value(0)
			}
			if data.Sequences[i].MatchEntries[ii].TlocListId != types.StringNull() {
				data.Sequences[i].MatchEntries[ii].TlocListVersion = types.Int64Value(0)
			}
			if data.Sequences[i].MatchEntries[ii].VpnListId != types.StringNull() {
				data.Sequences[i].MatchEntries[ii].VpnListVersion = types.Int64Value(0)
			}
			if data.Sequences[i].MatchEntries[ii].PrefixListId != types.StringNull() {
				data.Sequences[i].MatchEntries[ii].PrefixListVersion = types.Int64Value(0)
			}
		}
		for ii := range data.Sequences[i].ActionEntries {
			for iii := range data.Sequences[i].ActionEntries[ii].SetParameters {
				if data.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocListId != types.StringNull() {
					data.Sequences[i].ActionEntries[ii].SetParameters[iii].TlocListVersion = types.Int64Value(0)
				}
				if data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocListId != types.StringNull() {
					data.Sequences[i].ActionEntries[ii].SetParameters[iii].ServiceTlocListVersion = types.Int64Value(0)
				}
			}
			if data.Sequences[i].ActionEntries[ii].ExportToVpnListId != types.StringNull() {
				data.Sequences[i].ActionEntries[ii].ExportToVpnListVersion = types.Int64Value(0)
			}
		}
	}
}

// End of section. //template:end processImport
