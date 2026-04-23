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
type ServiceRoutePolicy struct {
	Id               types.String                  `tfsdk:"id"`
	Version          types.Int64                   `tfsdk:"version"`
	Name             types.String                  `tfsdk:"name"`
	Description      types.String                  `tfsdk:"description"`
	FeatureProfileId types.String                  `tfsdk:"feature_profile_id"`
	DefaultAction    types.String                  `tfsdk:"default_action"`
	Sequences        []ServiceRoutePolicySequences `tfsdk:"sequences"`
}

type ServiceRoutePolicySequences struct {
	Id           types.Int64                               `tfsdk:"id"`
	Name         types.String                              `tfsdk:"name"`
	BaseAction   types.String                              `tfsdk:"base_action"`
	Protocol     types.String                              `tfsdk:"protocol"`
	MatchEntries []ServiceRoutePolicySequencesMatchEntries `tfsdk:"match_entries"`
	Actions      []ServiceRoutePolicySequencesActions      `tfsdk:"actions"`
}

type ServiceRoutePolicySequencesMatchEntries struct {
	AsPathListId                  types.String                                                    `tfsdk:"as_path_list_id"`
	StandardCommunityListCriteria types.String                                                    `tfsdk:"standard_community_list_criteria"`
	StandardCommunityLists        []ServiceRoutePolicySequencesMatchEntriesStandardCommunityLists `tfsdk:"standard_community_lists"`
	ExpandedCommunityListId       types.String                                                    `tfsdk:"expanded_community_list_id"`
	ExtendedCommunityListId       types.String                                                    `tfsdk:"extended_community_list_id"`
	BgpLocalPreference            types.Int64                                                     `tfsdk:"bgp_local_preference"`
	Metric                        types.Int64                                                     `tfsdk:"metric"`
	OmpTag                        types.Int64                                                     `tfsdk:"omp_tag"`
	OspfTag                       types.Int64                                                     `tfsdk:"ospf_tag"`
	Ipv4AddressPrefixListId       types.String                                                    `tfsdk:"ipv4_address_prefix_list_id"`
	Ipv4NextHopPrefixListId       types.String                                                    `tfsdk:"ipv4_next_hop_prefix_list_id"`
	Ipv6AddressPrefixListId       types.String                                                    `tfsdk:"ipv6_address_prefix_list_id"`
	Ipv6NextHopPrefixListId       types.String                                                    `tfsdk:"ipv6_next_hop_prefix_list_id"`
}
type ServiceRoutePolicySequencesActions struct {
	AsPathPrepend     types.List   `tfsdk:"as_path_prepend"`
	CommunityAdditive types.Bool   `tfsdk:"community_additive"`
	Community         types.Set    `tfsdk:"community"`
	CommunityVariable types.String `tfsdk:"community_variable"`
	LocalPreference   types.Int64  `tfsdk:"local_preference"`
	Metric            types.Int64  `tfsdk:"metric"`
	MetricType        types.String `tfsdk:"metric_type"`
	OmpTag            types.Int64  `tfsdk:"omp_tag"`
	Origin            types.String `tfsdk:"origin"`
	OspfTag           types.Int64  `tfsdk:"ospf_tag"`
	Weight            types.Int64  `tfsdk:"weight"`
	Ipv4NextHop       types.String `tfsdk:"ipv4_next_hop"`
	Ipv6NextHop       types.String `tfsdk:"ipv6_next_hop"`
}

type ServiceRoutePolicySequencesMatchEntriesStandardCommunityLists struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceRoutePolicy) getModel() string {
	return "service_route_policy"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceRoutePolicy) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/route-policy", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceRoutePolicy) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if data.DefaultAction.IsNull() || data.DefaultAction.ValueString() == "reject" {
		if true {
			body, _ = sjson.Set(body, path+"defaultAction.optionType", "default")
			body, _ = sjson.Set(body, path+"defaultAction.value", "reject")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"defaultAction.optionType", "global")
			body, _ = sjson.Set(body, path+"defaultAction.value", data.DefaultAction.ValueString())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"sequences", []interface{}{})
		for _, item := range data.Sequences {
			itemBody := ""
			if !item.Id.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceId.value", item.Id.ValueInt64())
				}
			}
			if !item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceName.value", item.Name.ValueString())
				}
			}
			if item.BaseAction.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "baseAction.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "baseAction.value", "reject")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "baseAction.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "baseAction.value", item.BaseAction.ValueString())
				}
			}
			if item.Protocol.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "protocol.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "protocol.value", "IPV4")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "protocol.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "protocol.value", item.Protocol.ValueString())
				}
			}
			if true {

				for _, childItem := range item.MatchEntries {
					itemChildBody := ""
					if !childItem.AsPathListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "asPathList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "asPathList.refId.value", childItem.AsPathListId.ValueString())
						}
					}
					if !childItem.StandardCommunityListCriteria.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "communityList.criteria.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "communityList.criteria.value", childItem.StandardCommunityListCriteria.ValueString())
						}
					}
					if true {

						for _, childChildItem := range childItem.StandardCommunityLists {
							itemChildChildBody := ""
							if !childChildItem.Id.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "refId.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "refId.value", childChildItem.Id.ValueString())
								}
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "communityList.standardCommunityList.-1", itemChildChildBody)
						}
					}
					if !childItem.ExpandedCommunityListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "communityList.expandedCommunityList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "communityList.expandedCommunityList.refId.value", childItem.ExpandedCommunityListId.ValueString())
						}
					}
					if !childItem.ExtendedCommunityListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "extCommunityList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "extCommunityList.refId.value", childItem.ExtendedCommunityListId.ValueString())
						}
					}
					if !childItem.BgpLocalPreference.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "bgpLocalPreference.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "bgpLocalPreference.value", childItem.BgpLocalPreference.ValueInt64())
						}
					}
					if !childItem.Metric.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "metric.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "metric.value", childItem.Metric.ValueInt64())
						}
					}
					if !childItem.OmpTag.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "ompTag.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "ompTag.value", childItem.OmpTag.ValueInt64())
						}
					}
					if !childItem.OspfTag.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "ospfTag.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "ospfTag.value", childItem.OspfTag.ValueInt64())
						}
					}
					if !childItem.Ipv4AddressPrefixListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "ipv4Address.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "ipv4Address.refId.value", childItem.Ipv4AddressPrefixListId.ValueString())
						}
					}
					if !childItem.Ipv4NextHopPrefixListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "ipv4NextHop.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "ipv4NextHop.refId.value", childItem.Ipv4NextHopPrefixListId.ValueString())
						}
					}
					if !childItem.Ipv6AddressPrefixListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "ipv6Address.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "ipv6Address.refId.value", childItem.Ipv6AddressPrefixListId.ValueString())
						}
					}
					if !childItem.Ipv6NextHopPrefixListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "ipv6NextHop.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "ipv6NextHop.refId.value", childItem.Ipv6NextHopPrefixListId.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "matchEntries.-1", itemChildBody)
				}
			}
			if true {

				for _, childItem := range item.Actions {
					itemChildBody := ""
					if true {
						itemChildBody, _ = sjson.Set(itemChildBody, "accept.enableAcceptAction.optionType", "default")
						itemChildBody, _ = sjson.Set(itemChildBody, "accept.enableAcceptAction.value", true)
					}
					if !childItem.AsPathPrepend.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setAsPath.prepend.optionType", "global")
							var values []string
							childItem.AsPathPrepend.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setAsPath.prepend.value", values)
						}
					}
					if childItem.CommunityAdditive.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setCommunity.additive.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setCommunity.additive.value", false)
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setCommunity.additive.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setCommunity.additive.value", childItem.CommunityAdditive.ValueBool())
						}
					}

					if !childItem.CommunityVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setCommunity.community.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setCommunity.community.value", childItem.CommunityVariable.ValueString())
						}
					} else if !childItem.Community.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setCommunity.community.optionType", "global")
							var values []string
							childItem.Community.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setCommunity.community.value", values)
						}
					}
					if !childItem.LocalPreference.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setLocalPreference.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setLocalPreference.value", childItem.LocalPreference.ValueInt64())
						}
					}
					if !childItem.Metric.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setMetric.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setMetric.value", childItem.Metric.ValueInt64())
						}
					}
					if !childItem.MetricType.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setMetricType.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setMetricType.value", childItem.MetricType.ValueString())
						}
					}
					if !childItem.OmpTag.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setOmpTag.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setOmpTag.value", childItem.OmpTag.ValueInt64())
						}
					}
					if !childItem.Origin.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setOrigin.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setOrigin.value", childItem.Origin.ValueString())
						}
					}
					if !childItem.OspfTag.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setOspfTag.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setOspfTag.value", childItem.OspfTag.ValueInt64())
						}
					}
					if !childItem.Weight.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setWeight.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setWeight.value", childItem.Weight.ValueInt64())
						}
					}
					if !childItem.Ipv4NextHop.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setIpv4NextHop.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setIpv4NextHop.value", childItem.Ipv4NextHop.ValueString())
						}
					}
					if !childItem.Ipv6NextHop.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setIpv6NextHop.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setIpv6NextHop.value", childItem.Ipv6NextHop.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "actions.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"sequences.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceRoutePolicy) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	tempDefaultAction := data.DefaultAction
	data.DefaultAction = types.StringNull()

	if t := res.Get(path + "defaultAction.optionType"); t.Exists() {
		va := res.Get(path + "defaultAction.value")
		if t.String() == "global" || (t.String() == "default" && (tempDefaultAction.ValueString() == "reject" || (fullRead && tempDefaultAction.ValueString() == ""))) {
			data.DefaultAction = types.StringValue(va.String())
		}
	}
	oldSequences := data.Sequences
	if value := res.Get(path + "sequences"); value.Exists() && len(value.Array()) > 0 {
		data.Sequences = make([]ServiceRoutePolicySequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutePolicySequences{}
			item.Id = types.Int64Null()

			if t := v.Get("sequenceId.optionType"); t.Exists() {
				va := v.Get("sequenceId.value")
				if t.String() == "global" {
					item.Id = types.Int64Value(va.Int())
				}
			}
			item.Name = types.StringNull()

			if t := v.Get("sequenceName.optionType"); t.Exists() {
				va := v.Get("sequenceName.value")
				if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			item.BaseAction = types.StringNull()

			if t := v.Get("baseAction.optionType"); t.Exists() {
				va := v.Get("baseAction.value")
				if t.String() == "global" {
					item.BaseAction = types.StringValue(va.String())
				}
			}
			item.Protocol = types.StringNull()

			if t := v.Get("protocol.optionType"); t.Exists() {
				va := v.Get("protocol.value")
				if t.String() == "global" {
					item.Protocol = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("matchEntries"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.MatchEntries = make([]ServiceRoutePolicySequencesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceRoutePolicySequencesMatchEntries{}
					cItem.AsPathListId = types.StringNull()

					if t := cv.Get("asPathList.refId.optionType"); t.Exists() {
						va := cv.Get("asPathList.refId.value")
						if t.String() == "global" {
							cItem.AsPathListId = types.StringValue(va.String())
						}
					}
					cItem.StandardCommunityListCriteria = types.StringNull()

					if t := cv.Get("communityList.criteria.optionType"); t.Exists() {
						va := cv.Get("communityList.criteria.value")
						if t.String() == "global" {
							cItem.StandardCommunityListCriteria = types.StringValue(va.String())
						}
					}
					if ccValue := cv.Get("communityList.standardCommunityList"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.StandardCommunityLists = make([]ServiceRoutePolicySequencesMatchEntriesStandardCommunityLists, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := ServiceRoutePolicySequencesMatchEntriesStandardCommunityLists{}
							ccItem.Id = types.StringNull()

							if t := ccv.Get("refId.optionType"); t.Exists() {
								va := ccv.Get("refId.value")
								if t.String() == "global" {
									ccItem.Id = types.StringValue(va.String())
								}
							}
							cItem.StandardCommunityLists = append(cItem.StandardCommunityLists, ccItem)
							return true
						})
					}
					cItem.ExpandedCommunityListId = types.StringNull()

					if t := cv.Get("communityList.expandedCommunityList.refId.optionType"); t.Exists() {
						va := cv.Get("communityList.expandedCommunityList.refId.value")
						if t.String() == "global" {
							cItem.ExpandedCommunityListId = types.StringValue(va.String())
						}
					}
					cItem.ExtendedCommunityListId = types.StringNull()

					if t := cv.Get("extCommunityList.refId.optionType"); t.Exists() {
						va := cv.Get("extCommunityList.refId.value")
						if t.String() == "global" {
							cItem.ExtendedCommunityListId = types.StringValue(va.String())
						}
					}
					cItem.BgpLocalPreference = types.Int64Null()

					if t := cv.Get("bgpLocalPreference.optionType"); t.Exists() {
						va := cv.Get("bgpLocalPreference.value")
						if t.String() == "global" {
							cItem.BgpLocalPreference = types.Int64Value(va.Int())
						}
					}
					cItem.Metric = types.Int64Null()

					if t := cv.Get("metric.optionType"); t.Exists() {
						va := cv.Get("metric.value")
						if t.String() == "global" {
							cItem.Metric = types.Int64Value(va.Int())
						}
					}
					cItem.OmpTag = types.Int64Null()

					if t := cv.Get("ompTag.optionType"); t.Exists() {
						va := cv.Get("ompTag.value")
						if t.String() == "global" {
							cItem.OmpTag = types.Int64Value(va.Int())
						}
					}
					cItem.OspfTag = types.Int64Null()

					if t := cv.Get("ospfTag.optionType"); t.Exists() {
						va := cv.Get("ospfTag.value")
						if t.String() == "global" {
							cItem.OspfTag = types.Int64Value(va.Int())
						}
					}
					cItem.Ipv4AddressPrefixListId = types.StringNull()

					if t := cv.Get("ipv4Address.refId.optionType"); t.Exists() {
						va := cv.Get("ipv4Address.refId.value")
						if t.String() == "global" {
							cItem.Ipv4AddressPrefixListId = types.StringValue(va.String())
						}
					}
					cItem.Ipv4NextHopPrefixListId = types.StringNull()

					if t := cv.Get("ipv4NextHop.refId.optionType"); t.Exists() {
						va := cv.Get("ipv4NextHop.refId.value")
						if t.String() == "global" {
							cItem.Ipv4NextHopPrefixListId = types.StringValue(va.String())
						}
					}
					cItem.Ipv6AddressPrefixListId = types.StringNull()

					if t := cv.Get("ipv6Address.refId.optionType"); t.Exists() {
						va := cv.Get("ipv6Address.refId.value")
						if t.String() == "global" {
							cItem.Ipv6AddressPrefixListId = types.StringValue(va.String())
						}
					}
					cItem.Ipv6NextHopPrefixListId = types.StringNull()

					if t := cv.Get("ipv6NextHop.refId.optionType"); t.Exists() {
						va := cv.Get("ipv6NextHop.refId.value")
						if t.String() == "global" {
							cItem.Ipv6NextHopPrefixListId = types.StringValue(va.String())
						}
					}
					item.MatchEntries = append(item.MatchEntries, cItem)
					return true
				})
			}
			if cValue := v.Get("actions"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Actions = make([]ServiceRoutePolicySequencesActions, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceRoutePolicySequencesActions{}
					cItem.AsPathPrepend = types.ListNull(types.StringType)

					if t := cv.Get("accept.setAsPath.prepend.optionType"); t.Exists() {
						va := cv.Get("accept.setAsPath.prepend.value")
						if t.String() == "global" {
							cItem.AsPathPrepend = helpers.GetStringList(va.Array())
						}
					}
					cItem.CommunityAdditive = types.BoolNull()

					if t := cv.Get("accept.setCommunity.additive.optionType"); t.Exists() {
						va := cv.Get("accept.setCommunity.additive.value")
						if t.String() == "global" {
							cItem.CommunityAdditive = types.BoolValue(va.Bool())
						}
					}
					cItem.Community = types.SetNull(types.StringType)
					cItem.CommunityVariable = types.StringNull()
					if t := cv.Get("accept.setCommunity.community.optionType"); t.Exists() {
						va := cv.Get("accept.setCommunity.community.value")
						if t.String() == "variable" {
							cItem.CommunityVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Community = helpers.GetStringSet(va.Array())
						}
					}
					cItem.LocalPreference = types.Int64Null()

					if t := cv.Get("accept.setLocalPreference.optionType"); t.Exists() {
						va := cv.Get("accept.setLocalPreference.value")
						if t.String() == "global" {
							cItem.LocalPreference = types.Int64Value(va.Int())
						}
					}
					cItem.Metric = types.Int64Null()

					if t := cv.Get("accept.setMetric.optionType"); t.Exists() {
						va := cv.Get("accept.setMetric.value")
						if t.String() == "global" {
							cItem.Metric = types.Int64Value(va.Int())
						}
					}
					cItem.MetricType = types.StringNull()

					if t := cv.Get("accept.setMetricType.optionType"); t.Exists() {
						va := cv.Get("accept.setMetricType.value")
						if t.String() == "global" {
							cItem.MetricType = types.StringValue(va.String())
						}
					}
					cItem.OmpTag = types.Int64Null()

					if t := cv.Get("accept.setOmpTag.optionType"); t.Exists() {
						va := cv.Get("accept.setOmpTag.value")
						if t.String() == "global" {
							cItem.OmpTag = types.Int64Value(va.Int())
						}
					}
					cItem.Origin = types.StringNull()

					if t := cv.Get("accept.setOrigin.optionType"); t.Exists() {
						va := cv.Get("accept.setOrigin.value")
						if t.String() == "global" {
							cItem.Origin = types.StringValue(va.String())
						}
					}
					cItem.OspfTag = types.Int64Null()

					if t := cv.Get("accept.setOspfTag.optionType"); t.Exists() {
						va := cv.Get("accept.setOspfTag.value")
						if t.String() == "global" {
							cItem.OspfTag = types.Int64Value(va.Int())
						}
					}
					cItem.Weight = types.Int64Null()

					if t := cv.Get("accept.setWeight.optionType"); t.Exists() {
						va := cv.Get("accept.setWeight.value")
						if t.String() == "global" {
							cItem.Weight = types.Int64Value(va.Int())
						}
					}
					cItem.Ipv4NextHop = types.StringNull()

					if t := cv.Get("accept.setIpv4NextHop.optionType"); t.Exists() {
						va := cv.Get("accept.setIpv4NextHop.value")
						if t.String() == "global" {
							cItem.Ipv4NextHop = types.StringValue(va.String())
						}
					}
					cItem.Ipv6NextHop = types.StringNull()

					if t := cv.Get("accept.setIpv6NextHop.optionType"); t.Exists() {
						va := cv.Get("accept.setIpv6NextHop.value")
						if t.String() == "global" {
							cItem.Ipv6NextHop = types.StringValue(va.String())
						}
					}
					item.Actions = append(item.Actions, cItem)
					return true
				})
			}
			data.Sequences = append(data.Sequences, item)
			return true
		})
	} else {
		data.Sequences = nil
	}
	if !fullRead {
		resultSequences := make([]ServiceRoutePolicySequences, 0, len(data.Sequences))
		matchedSequences := make([]bool, len(data.Sequences))
		for _, oldItem := range oldSequences {
			for ni := range data.Sequences {
				if matchedSequences[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.Id.ValueInt64() != data.Sequences[ni].Id.ValueInt64() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedSequences[ni] = true
					{
						resultC := make([]ServiceRoutePolicySequencesMatchEntries, 0, len(data.Sequences[ni].MatchEntries))
						matchedC := make([]bool, len(data.Sequences[ni].MatchEntries))
						for _, oldCItem := range oldItem.MatchEntries {
							for nci := range data.Sequences[ni].MatchEntries {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if oldCItem.AsPathListId.ValueString() != data.Sequences[ni].MatchEntries[nci].AsPathListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.StandardCommunityListCriteria.ValueString() != data.Sequences[ni].MatchEntries[nci].StandardCommunityListCriteria.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.ExpandedCommunityListId.ValueString() != data.Sequences[ni].MatchEntries[nci].ExpandedCommunityListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.ExtendedCommunityListId.ValueString() != data.Sequences[ni].MatchEntries[nci].ExtendedCommunityListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.BgpLocalPreference.ValueInt64() != data.Sequences[ni].MatchEntries[nci].BgpLocalPreference.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Metric.ValueInt64() != data.Sequences[ni].MatchEntries[nci].Metric.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.OmpTag.ValueInt64() != data.Sequences[ni].MatchEntries[nci].OmpTag.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.OspfTag.ValueInt64() != data.Sequences[ni].MatchEntries[nci].OspfTag.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Ipv4AddressPrefixListId.ValueString() != data.Sequences[ni].MatchEntries[nci].Ipv4AddressPrefixListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Ipv4NextHopPrefixListId.ValueString() != data.Sequences[ni].MatchEntries[nci].Ipv4NextHopPrefixListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Ipv6AddressPrefixListId.ValueString() != data.Sequences[ni].MatchEntries[nci].Ipv6AddressPrefixListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Ipv6NextHopPrefixListId.ValueString() != data.Sequences[ni].MatchEntries[nci].Ipv6NextHopPrefixListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									{
										resultCC := make([]ServiceRoutePolicySequencesMatchEntriesStandardCommunityLists, 0, len(data.Sequences[ni].MatchEntries[nci].StandardCommunityLists))
										matchedCC := make([]bool, len(data.Sequences[ni].MatchEntries[nci].StandardCommunityLists))
										for _, oldCCItem := range oldCItem.StandardCommunityLists {
											for ncci := range data.Sequences[ni].MatchEntries[nci].StandardCommunityLists {
												if matchedCC[ncci] {
													continue
												}
												keyMatchCC := true
												if keyMatchCC {
													if oldCCItem.Id.ValueString() != data.Sequences[ni].MatchEntries[nci].StandardCommunityLists[ncci].Id.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													matchedCC[ncci] = true
													resultCC = append(resultCC, data.Sequences[ni].MatchEntries[nci].StandardCommunityLists[ncci])
													break
												}
											}
										}
										for ncci := range data.Sequences[ni].MatchEntries[nci].StandardCommunityLists {
											if !matchedCC[ncci] {
												resultCC = append(resultCC, data.Sequences[ni].MatchEntries[nci].StandardCommunityLists[ncci])
											}
										}
										data.Sequences[ni].MatchEntries[nci].StandardCommunityLists = resultCC
									}
									resultC = append(resultC, data.Sequences[ni].MatchEntries[nci])
									break
								}
							}
						}
						for nci := range data.Sequences[ni].MatchEntries {
							if !matchedC[nci] {
								resultC = append(resultC, data.Sequences[ni].MatchEntries[nci])
							}
						}
						data.Sequences[ni].MatchEntries = resultC
					}
					{
						resultC := make([]ServiceRoutePolicySequencesActions, 0, len(data.Sequences[ni].Actions))
						matchedC := make([]bool, len(data.Sequences[ni].Actions))
						for _, oldCItem := range oldItem.Actions {
							for nci := range data.Sequences[ni].Actions {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if helpers.GetStringFromList(oldCItem.AsPathPrepend).ValueString() != helpers.GetStringFromList(data.Sequences[ni].Actions[nci].AsPathPrepend).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.CommunityAdditive.ValueBool() != data.Sequences[ni].Actions[nci].CommunityAdditive.ValueBool() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.Community).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].Actions[nci].Community).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.LocalPreference.ValueInt64() != data.Sequences[ni].Actions[nci].LocalPreference.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Metric.ValueInt64() != data.Sequences[ni].Actions[nci].Metric.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.MetricType.ValueString() != data.Sequences[ni].Actions[nci].MetricType.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.OmpTag.ValueInt64() != data.Sequences[ni].Actions[nci].OmpTag.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Origin.ValueString() != data.Sequences[ni].Actions[nci].Origin.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.OspfTag.ValueInt64() != data.Sequences[ni].Actions[nci].OspfTag.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Weight.ValueInt64() != data.Sequences[ni].Actions[nci].Weight.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Ipv4NextHop.ValueString() != data.Sequences[ni].Actions[nci].Ipv4NextHop.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Ipv6NextHop.ValueString() != data.Sequences[ni].Actions[nci].Ipv6NextHop.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.Sequences[ni].Actions[nci])
									break
								}
							}
						}
						for nci := range data.Sequences[ni].Actions {
							if !matchedC[nci] {
								resultC = append(resultC, data.Sequences[ni].Actions[nci])
							}
						}
						data.Sequences[ni].Actions = resultC
					}
					resultSequences = append(resultSequences, data.Sequences[ni])
					break
				}
			}
		}
		for ni := range data.Sequences {
			if !matchedSequences[ni] {
				resultSequences = append(resultSequences, data.Sequences[ni])
			}
		}
		data.Sequences = resultSequences
	}
}

// End of section. //template:end fromBody
