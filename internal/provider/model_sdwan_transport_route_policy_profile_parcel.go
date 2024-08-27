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
type TransportRoutePolicy struct {
	Id               types.String                    `tfsdk:"id"`
	Version          types.Int64                     `tfsdk:"version"`
	Name             types.String                    `tfsdk:"name"`
	Description      types.String                    `tfsdk:"description"`
	FeatureProfileId types.String                    `tfsdk:"feature_profile_id"`
	DefaultAction    types.String                    `tfsdk:"default_action"`
	Sequences        []TransportRoutePolicySequences `tfsdk:"sequences"`
}

type TransportRoutePolicySequences struct {
	Id           types.Int64                                 `tfsdk:"id"`
	Name         types.String                                `tfsdk:"name"`
	BaseAction   types.String                                `tfsdk:"base_action"`
	Protocol     types.String                                `tfsdk:"protocol"`
	MatchEntries []TransportRoutePolicySequencesMatchEntries `tfsdk:"match_entries"`
	Actions      []TransportRoutePolicySequencesActions      `tfsdk:"actions"`
}

type TransportRoutePolicySequencesMatchEntries struct {
	AsPathListId                  types.String                                                      `tfsdk:"as_path_list_id"`
	StandardCommunityListCriteria types.String                                                      `tfsdk:"standard_community_list_criteria"`
	StandardCommunityLists        []TransportRoutePolicySequencesMatchEntriesStandardCommunityLists `tfsdk:"standard_community_lists"`
	ExpandedCommunityListId       types.String                                                      `tfsdk:"expanded_community_list_id"`
	ExtendedCommunityListId       types.String                                                      `tfsdk:"extended_community_list_id"`
	BgpLocalPreference            types.Int64                                                       `tfsdk:"bgp_local_preference"`
	Metric                        types.Int64                                                       `tfsdk:"metric"`
	OmpTag                        types.Int64                                                       `tfsdk:"omp_tag"`
	OspfTag                       types.Int64                                                       `tfsdk:"ospf_tag"`
	Ipv4AddressId                 types.String                                                      `tfsdk:"ipv4_address_id"`
	Ipv4NextHopId                 types.String                                                      `tfsdk:"ipv4_next_hop_id"`
	Ipv6AddressId                 types.String                                                      `tfsdk:"ipv6_address_id"`
	Ipv6NextHopId                 types.String                                                      `tfsdk:"ipv6_next_hop_id"`
}
type TransportRoutePolicySequencesActions struct {
	AsPathPrend       types.Set    `tfsdk:"as_path_prend"`
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

type TransportRoutePolicySequencesMatchEntriesStandardCommunityLists struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TransportRoutePolicy) getModel() string {
	return "transport_route_policy"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportRoutePolicy) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/route-policy", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportRoutePolicy) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if data.DefaultAction.IsNull() {
		body, _ = sjson.Set(body, path+"defaultAction.optionType", "default")
		body, _ = sjson.Set(body, path+"defaultAction.value", "reject")
	} else {
		body, _ = sjson.Set(body, path+"defaultAction.optionType", "global")
		body, _ = sjson.Set(body, path+"defaultAction.value", data.DefaultAction.ValueString())
	}
	body, _ = sjson.Set(body, path+"sequences", []interface{}{})
	for _, item := range data.Sequences {
		itemBody := ""
		if !item.Id.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sequenceId.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sequenceId.value", item.Id.ValueInt64())
		}
		if !item.Name.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sequenceName.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sequenceName.value", item.Name.ValueString())
		}
		if item.BaseAction.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "baseAction.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "baseAction.value", "reject")
		} else {
			itemBody, _ = sjson.Set(itemBody, "baseAction.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "baseAction.value", item.BaseAction.ValueString())
		}
		if item.Protocol.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "protocol.value", "IPV4")
		} else {
			itemBody, _ = sjson.Set(itemBody, "protocol.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "protocol.value", item.Protocol.ValueString())
		}

		for _, childItem := range item.MatchEntries {
			itemChildBody := ""
			if !childItem.AsPathListId.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "asPathList.refId.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "asPathList.refId.value", childItem.AsPathListId.ValueString())
			}
			if !childItem.StandardCommunityListCriteria.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "communityList.criteria.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "communityList.criteria.value", childItem.StandardCommunityListCriteria.ValueString())
			}

			for _, childChildItem := range childItem.StandardCommunityLists {
				itemChildChildBody := ""
				if !childChildItem.Id.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "refId.optionType", "global")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "refId.value", childChildItem.Id.ValueString())
				}
				itemChildBody, _ = sjson.SetRaw(itemChildBody, "communityList.standardCommunityList.-1", itemChildChildBody)
			}
			if !childItem.ExpandedCommunityListId.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "communityList.expandedCommunityList.refId.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "communityList.expandedCommunityList.refId.value", childItem.ExpandedCommunityListId.ValueString())
			}
			if !childItem.ExtendedCommunityListId.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "extCommunityList.refId.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "extCommunityList.refId.value", childItem.ExtendedCommunityListId.ValueString())
			}
			if !childItem.BgpLocalPreference.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "bgpLocalPreference.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "bgpLocalPreference.value", childItem.BgpLocalPreference.ValueInt64())
			}
			if !childItem.Metric.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "metric.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "metric.value", childItem.Metric.ValueInt64())
			}
			if !childItem.OmpTag.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "ompTag.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "ompTag.value", childItem.OmpTag.ValueInt64())
			}
			if !childItem.OspfTag.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "ospfTag.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "ospfTag.value", childItem.OspfTag.ValueInt64())
			}
			if !childItem.Ipv4AddressId.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv4Address.refId.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv4Address.refId.value", childItem.Ipv4AddressId.ValueString())
			}
			if !childItem.Ipv4NextHopId.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv4NextHop.refId.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv4NextHop.refId.value", childItem.Ipv4NextHopId.ValueString())
			}
			if !childItem.Ipv6AddressId.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv6Address.refId.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv6Address.refId.value", childItem.Ipv6AddressId.ValueString())
			}
			if !childItem.Ipv6NextHopId.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv6NextHop.refId.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "ipv6NextHop.refId.value", childItem.Ipv6NextHopId.ValueString())
			}
			itemBody, _ = sjson.SetRaw(itemBody, "matchEntries.-1", itemChildBody)
		}

		for _, childItem := range item.Actions {
			itemChildBody := ""
			itemChildBody, _ = sjson.Set(itemChildBody, "accept.enableAcceptAction.optionType", "default")
			itemChildBody, _ = sjson.Set(itemChildBody, "accept.enableAcceptAction.value", true)
			if !childItem.AsPathPrend.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setAsPath.prepend.optionType", "global")
				var values []int64
				childItem.AsPathPrend.ElementsAs(ctx, &values, false)
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setAsPath.prepend.value", values)
			}
			if childItem.CommunityAdditive.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setCommunity.additive.optionType", "default")
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setCommunity.additive.value", false)
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setCommunity.additive.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setCommunity.additive.value", childItem.CommunityAdditive.ValueBool())
			}

			if !childItem.CommunityVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setCommunity.community.optionType", "variable")
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setCommunity.community.value", childItem.CommunityVariable.ValueString())
			} else if !childItem.Community.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setCommunity.community.optionType", "global")
				var values []string
				childItem.Community.ElementsAs(ctx, &values, false)
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setCommunity.community.value", values)
			}
			if !childItem.LocalPreference.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setLocalPreference.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setLocalPreference.value", childItem.LocalPreference.ValueInt64())
			}
			if !childItem.Metric.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setMetric.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setMetric.value", childItem.Metric.ValueInt64())
			}
			if !childItem.MetricType.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setMetricType.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setMetricType.value", childItem.MetricType.ValueString())
			}
			if !childItem.OmpTag.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setOmpTag.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setOmpTag.value", childItem.OmpTag.ValueInt64())
			}
			if !childItem.Origin.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setOrigin.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setOrigin.value", childItem.Origin.ValueString())
			}
			if !childItem.OspfTag.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setOspfTag.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setOspfTag.value", childItem.OspfTag.ValueInt64())
			}
			if !childItem.Weight.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setWeight.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setWeight.value", childItem.Weight.ValueInt64())
			}
			if !childItem.Ipv4NextHop.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setIpv4NextHop.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setIpv4NextHop.value", childItem.Ipv4NextHop.ValueString())
			}
			if !childItem.Ipv6NextHop.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setIpv6NextHop.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "accept.setIpv6NextHop.value", childItem.Ipv6NextHop.ValueString())
			}
			itemBody, _ = sjson.SetRaw(itemBody, "actions.-1", itemChildBody)
		}
		body, _ = sjson.SetRaw(body, path+"sequences.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportRoutePolicy) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.DefaultAction = types.StringNull()

	if t := res.Get(path + "defaultAction.optionType"); t.Exists() {
		va := res.Get(path + "defaultAction.value")
		if t.String() == "global" {
			data.DefaultAction = types.StringValue(va.String())
		}
	}
	if value := res.Get(path + "sequences"); value.Exists() {
		data.Sequences = make([]TransportRoutePolicySequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportRoutePolicySequences{}
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
			if cValue := v.Get("matchEntries"); cValue.Exists() {
				item.MatchEntries = make([]TransportRoutePolicySequencesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TransportRoutePolicySequencesMatchEntries{}
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
					if ccValue := cv.Get("communityList.standardCommunityList"); ccValue.Exists() {
						cItem.StandardCommunityLists = make([]TransportRoutePolicySequencesMatchEntriesStandardCommunityLists, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := TransportRoutePolicySequencesMatchEntriesStandardCommunityLists{}
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
					cItem.Ipv4AddressId = types.StringNull()

					if t := cv.Get("ipv4Address.refId.optionType"); t.Exists() {
						va := cv.Get("ipv4Address.refId.value")
						if t.String() == "global" {
							cItem.Ipv4AddressId = types.StringValue(va.String())
						}
					}
					cItem.Ipv4NextHopId = types.StringNull()

					if t := cv.Get("ipv4NextHop.refId.optionType"); t.Exists() {
						va := cv.Get("ipv4NextHop.refId.value")
						if t.String() == "global" {
							cItem.Ipv4NextHopId = types.StringValue(va.String())
						}
					}
					cItem.Ipv6AddressId = types.StringNull()

					if t := cv.Get("ipv6Address.refId.optionType"); t.Exists() {
						va := cv.Get("ipv6Address.refId.value")
						if t.String() == "global" {
							cItem.Ipv6AddressId = types.StringValue(va.String())
						}
					}
					cItem.Ipv6NextHopId = types.StringNull()

					if t := cv.Get("ipv6NextHop.refId.optionType"); t.Exists() {
						va := cv.Get("ipv6NextHop.refId.value")
						if t.String() == "global" {
							cItem.Ipv6NextHopId = types.StringValue(va.String())
						}
					}
					item.MatchEntries = append(item.MatchEntries, cItem)
					return true
				})
			}
			if cValue := v.Get("actions"); cValue.Exists() {
				item.Actions = make([]TransportRoutePolicySequencesActions, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TransportRoutePolicySequencesActions{}
					cItem.AsPathPrend = types.SetNull(types.Int64Type)

					if t := cv.Get("accept.setAsPath.prepend.optionType"); t.Exists() {
						va := cv.Get("accept.setAsPath.prepend.value")
						if t.String() == "global" {
							cItem.AsPathPrend = helpers.GetInt64Set(va.Array())
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
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TransportRoutePolicy) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.DefaultAction = types.StringNull()

	if t := res.Get(path + "defaultAction.optionType"); t.Exists() {
		va := res.Get(path + "defaultAction.value")
		if t.String() == "global" {
			data.DefaultAction = types.StringValue(va.String())
		}
	}
	for i := range data.Sequences {
		keys := [...]string{"sequenceId"}
		keyValues := [...]string{strconv.FormatInt(data.Sequences[i].Id.ValueInt64(), 10)}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "sequences").ForEach(
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
		data.Sequences[i].Id = types.Int64Null()

		if t := r.Get("sequenceId.optionType"); t.Exists() {
			va := r.Get("sequenceId.value")
			if t.String() == "global" {
				data.Sequences[i].Id = types.Int64Value(va.Int())
			}
		}
		data.Sequences[i].Name = types.StringNull()

		if t := r.Get("sequenceName.optionType"); t.Exists() {
			va := r.Get("sequenceName.value")
			if t.String() == "global" {
				data.Sequences[i].Name = types.StringValue(va.String())
			}
		}
		data.Sequences[i].BaseAction = types.StringNull()

		if t := r.Get("baseAction.optionType"); t.Exists() {
			va := r.Get("baseAction.value")
			if t.String() == "global" {
				data.Sequences[i].BaseAction = types.StringValue(va.String())
			}
		}
		data.Sequences[i].Protocol = types.StringNull()

		if t := r.Get("protocol.optionType"); t.Exists() {
			va := r.Get("protocol.value")
			if t.String() == "global" {
				data.Sequences[i].Protocol = types.StringValue(va.String())
			}
		}
		for ci := range data.Sequences[i].MatchEntries {
			keys := [...]string{"asPathList.refId", "communityList.criteria", "communityList.expandedCommunityList.refId", "extCommunityList.refId", "bgpLocalPreference", "metric", "ompTag", "ospfTag", "ipv4Address.refId", "ipv4NextHop.refId", "ipv6Address.refId", "ipv6NextHop.refId"}
			keyValues := [...]string{data.Sequences[i].MatchEntries[ci].AsPathListId.ValueString(), data.Sequences[i].MatchEntries[ci].StandardCommunityListCriteria.ValueString(), data.Sequences[i].MatchEntries[ci].ExpandedCommunityListId.ValueString(), data.Sequences[i].MatchEntries[ci].ExtendedCommunityListId.ValueString(), strconv.FormatInt(data.Sequences[i].MatchEntries[ci].BgpLocalPreference.ValueInt64(), 10), strconv.FormatInt(data.Sequences[i].MatchEntries[ci].Metric.ValueInt64(), 10), strconv.FormatInt(data.Sequences[i].MatchEntries[ci].OmpTag.ValueInt64(), 10), strconv.FormatInt(data.Sequences[i].MatchEntries[ci].OspfTag.ValueInt64(), 10), data.Sequences[i].MatchEntries[ci].Ipv4AddressId.ValueString(), data.Sequences[i].MatchEntries[ci].Ipv4NextHopId.ValueString(), data.Sequences[i].MatchEntries[ci].Ipv6AddressId.ValueString(), data.Sequences[i].MatchEntries[ci].Ipv6NextHopId.ValueString()}
			keyValuesVariables := [...]string{"", "", "", "", "", "", "", "", "", "", "", "", ""}

			var cr gjson.Result
			r.Get("matchEntries").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			data.Sequences[i].MatchEntries[ci].AsPathListId = types.StringNull()

			if t := cr.Get("asPathList.refId.optionType"); t.Exists() {
				va := cr.Get("asPathList.refId.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].AsPathListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].StandardCommunityListCriteria = types.StringNull()

			if t := cr.Get("communityList.criteria.optionType"); t.Exists() {
				va := cr.Get("communityList.criteria.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].StandardCommunityListCriteria = types.StringValue(va.String())
				}
			}
			for cci := range data.Sequences[i].MatchEntries[ci].StandardCommunityLists {
				keys := [...]string{"refId"}
				keyValues := [...]string{data.Sequences[i].MatchEntries[ci].StandardCommunityLists[cci].Id.ValueString()}
				keyValuesVariables := [...]string{""}

				var ccr gjson.Result
				cr.Get("communityList.standardCommunityList").ForEach(
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
							ccr = v
							return false
						}
						return true
					},
				)
				data.Sequences[i].MatchEntries[ci].StandardCommunityLists[cci].Id = types.StringNull()

				if t := ccr.Get("refId.optionType"); t.Exists() {
					va := ccr.Get("refId.value")
					if t.String() == "global" {
						data.Sequences[i].MatchEntries[ci].StandardCommunityLists[cci].Id = types.StringValue(va.String())
					}
				}
			}
			data.Sequences[i].MatchEntries[ci].ExpandedCommunityListId = types.StringNull()

			if t := cr.Get("communityList.expandedCommunityList.refId.optionType"); t.Exists() {
				va := cr.Get("communityList.expandedCommunityList.refId.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].ExpandedCommunityListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].ExtendedCommunityListId = types.StringNull()

			if t := cr.Get("extCommunityList.refId.optionType"); t.Exists() {
				va := cr.Get("extCommunityList.refId.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].ExtendedCommunityListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].BgpLocalPreference = types.Int64Null()

			if t := cr.Get("bgpLocalPreference.optionType"); t.Exists() {
				va := cr.Get("bgpLocalPreference.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].BgpLocalPreference = types.Int64Value(va.Int())
				}
			}
			data.Sequences[i].MatchEntries[ci].Metric = types.Int64Null()

			if t := cr.Get("metric.optionType"); t.Exists() {
				va := cr.Get("metric.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].Metric = types.Int64Value(va.Int())
				}
			}
			data.Sequences[i].MatchEntries[ci].OmpTag = types.Int64Null()

			if t := cr.Get("ompTag.optionType"); t.Exists() {
				va := cr.Get("ompTag.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].OmpTag = types.Int64Value(va.Int())
				}
			}
			data.Sequences[i].MatchEntries[ci].OspfTag = types.Int64Null()

			if t := cr.Get("ospfTag.optionType"); t.Exists() {
				va := cr.Get("ospfTag.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].OspfTag = types.Int64Value(va.Int())
				}
			}
			data.Sequences[i].MatchEntries[ci].Ipv4AddressId = types.StringNull()

			if t := cr.Get("ipv4Address.refId.optionType"); t.Exists() {
				va := cr.Get("ipv4Address.refId.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].Ipv4AddressId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].Ipv4NextHopId = types.StringNull()

			if t := cr.Get("ipv4NextHop.refId.optionType"); t.Exists() {
				va := cr.Get("ipv4NextHop.refId.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].Ipv4NextHopId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].Ipv6AddressId = types.StringNull()

			if t := cr.Get("ipv6Address.refId.optionType"); t.Exists() {
				va := cr.Get("ipv6Address.refId.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].Ipv6AddressId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].Ipv6NextHopId = types.StringNull()

			if t := cr.Get("ipv6NextHop.refId.optionType"); t.Exists() {
				va := cr.Get("ipv6NextHop.refId.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].Ipv6NextHopId = types.StringValue(va.String())
				}
			}
		}
		for ci := range data.Sequences[i].Actions {
			keys := [...]string{"accept.setCommunity.additive", "accept.setLocalPreference", "accept.setMetric", "accept.setMetricType", "accept.setOmpTag", "accept.setOrigin", "accept.setOspfTag", "accept.setWeight", "accept.setIpv4NextHop", "accept.setIpv6NextHop"}
			keyValues := [...]string{strconv.FormatBool(data.Sequences[i].Actions[ci].CommunityAdditive.ValueBool()), strconv.FormatInt(data.Sequences[i].Actions[ci].LocalPreference.ValueInt64(), 10), strconv.FormatInt(data.Sequences[i].Actions[ci].Metric.ValueInt64(), 10), data.Sequences[i].Actions[ci].MetricType.ValueString(), strconv.FormatInt(data.Sequences[i].Actions[ci].OmpTag.ValueInt64(), 10), data.Sequences[i].Actions[ci].Origin.ValueString(), strconv.FormatInt(data.Sequences[i].Actions[ci].OspfTag.ValueInt64(), 10), strconv.FormatInt(data.Sequences[i].Actions[ci].Weight.ValueInt64(), 10), data.Sequences[i].Actions[ci].Ipv4NextHop.ValueString(), data.Sequences[i].Actions[ci].Ipv6NextHop.ValueString()}
			keyValuesVariables := [...]string{"", "", "", data.Sequences[i].Actions[ci].CommunityVariable.ValueString(), "", "", "", "", "", "", "", "", ""}

			var cr gjson.Result
			r.Get("actions").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			data.Sequences[i].Actions[ci].AsPathPrend = types.SetNull(types.Int64Type)

			if t := cr.Get("accept.setAsPath.prepend.optionType"); t.Exists() {
				va := cr.Get("accept.setAsPath.prepend.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].AsPathPrend = helpers.GetInt64Set(va.Array())
				}
			}
			data.Sequences[i].Actions[ci].CommunityAdditive = types.BoolNull()

			if t := cr.Get("accept.setCommunity.additive.optionType"); t.Exists() {
				va := cr.Get("accept.setCommunity.additive.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].CommunityAdditive = types.BoolValue(va.Bool())
				}
			}
			data.Sequences[i].Actions[ci].Community = types.SetNull(types.StringType)
			data.Sequences[i].Actions[ci].CommunityVariable = types.StringNull()
			if t := cr.Get("accept.setCommunity.community.optionType"); t.Exists() {
				va := cr.Get("accept.setCommunity.community.value")
				if t.String() == "variable" {
					data.Sequences[i].Actions[ci].CommunityVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Sequences[i].Actions[ci].Community = helpers.GetStringSet(va.Array())
				}
			}
			data.Sequences[i].Actions[ci].LocalPreference = types.Int64Null()

			if t := cr.Get("accept.setLocalPreference.optionType"); t.Exists() {
				va := cr.Get("accept.setLocalPreference.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].LocalPreference = types.Int64Value(va.Int())
				}
			}
			data.Sequences[i].Actions[ci].Metric = types.Int64Null()

			if t := cr.Get("accept.setMetric.optionType"); t.Exists() {
				va := cr.Get("accept.setMetric.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].Metric = types.Int64Value(va.Int())
				}
			}
			data.Sequences[i].Actions[ci].MetricType = types.StringNull()

			if t := cr.Get("accept.setMetricType.optionType"); t.Exists() {
				va := cr.Get("accept.setMetricType.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].MetricType = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Actions[ci].OmpTag = types.Int64Null()

			if t := cr.Get("accept.setOmpTag.optionType"); t.Exists() {
				va := cr.Get("accept.setOmpTag.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].OmpTag = types.Int64Value(va.Int())
				}
			}
			data.Sequences[i].Actions[ci].Origin = types.StringNull()

			if t := cr.Get("accept.setOrigin.optionType"); t.Exists() {
				va := cr.Get("accept.setOrigin.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].Origin = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Actions[ci].OspfTag = types.Int64Null()

			if t := cr.Get("accept.setOspfTag.optionType"); t.Exists() {
				va := cr.Get("accept.setOspfTag.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].OspfTag = types.Int64Value(va.Int())
				}
			}
			data.Sequences[i].Actions[ci].Weight = types.Int64Null()

			if t := cr.Get("accept.setWeight.optionType"); t.Exists() {
				va := cr.Get("accept.setWeight.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].Weight = types.Int64Value(va.Int())
				}
			}
			data.Sequences[i].Actions[ci].Ipv4NextHop = types.StringNull()

			if t := cr.Get("accept.setIpv4NextHop.optionType"); t.Exists() {
				va := cr.Get("accept.setIpv4NextHop.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].Ipv4NextHop = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Actions[ci].Ipv6NextHop = types.StringNull()

			if t := cr.Get("accept.setIpv6NextHop.optionType"); t.Exists() {
				va := cr.Get("accept.setIpv6NextHop.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].Ipv6NextHop = types.StringValue(va.String())
				}
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *TransportRoutePolicy) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.DefaultAction.IsNull() {
		return false
	}
	if len(data.Sequences) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
