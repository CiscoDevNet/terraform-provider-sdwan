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
type RoutePolicyDefinition struct {
	Id            types.String                     `tfsdk:"id"`
	Version       types.Int64                      `tfsdk:"version"`
	Type          types.String                     `tfsdk:"type"`
	Name          types.String                     `tfsdk:"name"`
	Description   types.String                     `tfsdk:"description"`
	DefaultAction types.String                     `tfsdk:"default_action"`
	Sequences     []RoutePolicyDefinitionSequences `tfsdk:"sequences"`
}

type RoutePolicyDefinitionSequences struct {
	Id            types.Int64                                   `tfsdk:"id"`
	IpType        types.String                                  `tfsdk:"ip_type"`
	Name          types.String                                  `tfsdk:"name"`
	BaseAction    types.String                                  `tfsdk:"base_action"`
	MatchEntries  []RoutePolicyDefinitionSequencesMatchEntries  `tfsdk:"match_entries"`
	ActionEntries []RoutePolicyDefinitionSequencesActionEntries `tfsdk:"action_entries"`
}

type RoutePolicyDefinitionSequencesMatchEntries struct {
	Type                          types.String `tfsdk:"type"`
	PrefixListId                  types.String `tfsdk:"prefix_list_id"`
	PrefixListVersion             types.Int64  `tfsdk:"prefix_list_version"`
	AsPathListId                  types.String `tfsdk:"as_path_list_id"`
	AsPathListVersion             types.Int64  `tfsdk:"as_path_list_version"`
	CommunityListId               types.String `tfsdk:"community_list_id"`
	CommunityListVersion          types.Int64  `tfsdk:"community_list_version"`
	CommunityListMatchFlag        types.String `tfsdk:"community_list_match_flag"`
	CommunityListIds              types.Set    `tfsdk:"community_list_ids"`
	CommunityListVersions         types.List   `tfsdk:"community_list_versions"`
	ExpandedCommunityListId       types.String `tfsdk:"expanded_community_list_id"`
	ExpandedCommunityListVariable types.String `tfsdk:"expanded_community_list_variable"`
	ExpandedCommunityListVersion  types.Int64  `tfsdk:"expanded_community_list_version"`
	ExtendedCommunityListId       types.String `tfsdk:"extended_community_list_id"`
	ExtendedCommunityListVersion  types.Int64  `tfsdk:"extended_community_list_version"`
	LocalPreference               types.Int64  `tfsdk:"local_preference"`
	Metric                        types.Int64  `tfsdk:"metric"`
	NextHopPrefixListId           types.String `tfsdk:"next_hop_prefix_list_id"`
	NextHopPrefixListVersion      types.Int64  `tfsdk:"next_hop_prefix_list_version"`
	Origin                        types.String `tfsdk:"origin"`
	Peer                          types.String `tfsdk:"peer"`
	OmpTag                        types.Int64  `tfsdk:"omp_tag"`
	OspfTag                       types.Int64  `tfsdk:"ospf_tag"`
}
type RoutePolicyDefinitionSequencesActionEntries struct {
	Type                types.String `tfsdk:"type"`
	Aggregator          types.Int64  `tfsdk:"aggregator"`
	AggregatorIpAddress types.String `tfsdk:"aggregator_ip_address"`
	AsPathPrepend       types.String `tfsdk:"as_path_prepend"`
	AsPathExclude       types.String `tfsdk:"as_path_exclude"`
	AtomicAggregate     types.Bool   `tfsdk:"atomic_aggregate"`
	Community           types.String `tfsdk:"community"`
	CommunityVariable   types.String `tfsdk:"community_variable"`
	CommunityAdditive   types.Bool   `tfsdk:"community_additive"`
	LocalPreference     types.Int64  `tfsdk:"local_preference"`
	Metric              types.Int64  `tfsdk:"metric"`
	Weight              types.Int64  `tfsdk:"weight"`
	MetricType          types.String `tfsdk:"metric_type"`
	NextHop             types.String `tfsdk:"next_hop"`
	OmpTag              types.Int64  `tfsdk:"omp_tag"`
	OspfTag             types.Int64  `tfsdk:"ospf_tag"`
	Origin              types.String `tfsdk:"origin"`
	Originator          types.String `tfsdk:"originator"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data RoutePolicyDefinition) getPath() string {
	return "/template/policy/definition/vedgeroute/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data RoutePolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "vedgeRoute")
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
			if !item.IpType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceIpType", item.IpType.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceName", item.Name.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "sequenceType", "vedgeRoute")
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
					if !childItem.PrefixListId.IsNull() && childItem.Type.ValueString() == "address" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.PrefixListId.ValueString())
					}
					if !childItem.AsPathListId.IsNull() && childItem.Type.ValueString() == "asPath" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.AsPathListId.ValueString())
					}
					if !childItem.CommunityListId.IsNull() && childItem.Type.ValueString() == "community" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.CommunityListId.ValueString())
					}
					if !childItem.CommunityListMatchFlag.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "matchFlag", childItem.CommunityListMatchFlag.ValueString())
					}
					if !childItem.CommunityListIds.IsNull() && childItem.Type.ValueString() == "advancedCommunity" {
						var values []string
						childItem.CommunityListIds.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "refs", values)
					}
					if !childItem.ExpandedCommunityListId.IsNull() && childItem.Type.ValueString() == "expandedCommunity" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.ExpandedCommunityListId.ValueString())
					}
					if !childItem.ExpandedCommunityListVariable.IsNull() && childItem.Type.ValueString() == "expandedCommunityInline" {
						itemChildBody, _ = sjson.Set(itemChildBody, "vipVariableName", childItem.ExpandedCommunityListVariable.ValueString())
					}
					if !childItem.ExtendedCommunityListId.IsNull() && childItem.Type.ValueString() == "extCommunity" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.ExtendedCommunityListId.ValueString())
					}
					if !childItem.LocalPreference.IsNull() && childItem.Type.ValueString() == "localPreference" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.LocalPreference.ValueInt64()))
					}
					if !childItem.Metric.IsNull() && childItem.Type.ValueString() == "metric" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.Metric.ValueInt64()))
					}
					if !childItem.NextHopPrefixListId.IsNull() && childItem.Type.ValueString() == "nextHop" {
						itemChildBody, _ = sjson.Set(itemChildBody, "ref", childItem.NextHopPrefixListId.ValueString())
					}
					if !childItem.Origin.IsNull() && childItem.Type.ValueString() == "origin" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Origin.ValueString())
					}
					if !childItem.Peer.IsNull() && childItem.Type.ValueString() == "peer" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Peer.ValueString())
					}
					if !childItem.OmpTag.IsNull() && childItem.Type.ValueString() == "ompTag" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.OmpTag.ValueInt64()))
					}
					if !childItem.OspfTag.IsNull() && childItem.Type.ValueString() == "ospfTag" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.OspfTag.ValueInt64()))
					}
					itemBody, _ = sjson.SetRaw(itemBody, "match.entries.-1", itemChildBody)
				}
			}
			if true && len(item.ActionEntries) > 0 {
				itemBody, _ = sjson.Set(itemBody, "actions.0.type", "set")
			}
			if true && len(item.ActionEntries) > 0 {
				itemBody, _ = sjson.Set(itemBody, "actions.0.parameter", []interface{}{})
				for _, childItem := range item.ActionEntries {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "field", childItem.Type.ValueString())
					}
					if !childItem.Aggregator.IsNull() && childItem.Type.ValueString() == "aggregator" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.aggregator", fmt.Sprint(childItem.Aggregator.ValueInt64()))
					}
					if !childItem.AggregatorIpAddress.IsNull() && childItem.Type.ValueString() == "aggregator" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.ipAddress", childItem.AggregatorIpAddress.ValueString())
					}
					if !childItem.AsPathPrepend.IsNull() && childItem.Type.ValueString() == "asPath" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.prepend", childItem.AsPathPrepend.ValueString())
					}
					if !childItem.AsPathExclude.IsNull() && childItem.Type.ValueString() == "asPath" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.exclude", childItem.AsPathExclude.ValueString())
					}
					if !childItem.AtomicAggregate.IsNull() && childItem.Type.ValueString() == "atomicAggregate" {
						if false && childItem.AtomicAggregate.ValueBool() {
							itemChildBody, _ = sjson.Set(itemChildBody, "value", "")
						} else {
							itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.AtomicAggregate.ValueBool()))
						}
					}
					if !childItem.Community.IsNull() && childItem.Type.ValueString() == "community" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Community.ValueString())
					}
					if !childItem.CommunityVariable.IsNull() && childItem.Type.ValueString() == "community" {
						itemChildBody, _ = sjson.Set(itemChildBody, "vipVariableName", childItem.CommunityVariable.ValueString())
					}
					if !childItem.CommunityAdditive.IsNull() && childItem.Type.ValueString() == "communityAdditive" {
						if false && childItem.CommunityAdditive.ValueBool() {
							itemChildBody, _ = sjson.Set(itemChildBody, "value", "")
						} else {
							itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.CommunityAdditive.ValueBool()))
						}
					}
					if !childItem.LocalPreference.IsNull() && childItem.Type.ValueString() == "localPreference" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.LocalPreference.ValueInt64()))
					}
					if !childItem.Metric.IsNull() && childItem.Type.ValueString() == "metric" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.Metric.ValueInt64()))
					}
					if !childItem.Weight.IsNull() && childItem.Type.ValueString() == "weight" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.Weight.ValueInt64()))
					}
					if !childItem.MetricType.IsNull() && childItem.Type.ValueString() == "metricType" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.MetricType.ValueString())
					}
					if !childItem.NextHop.IsNull() && childItem.Type.ValueString() == "nextHop" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.NextHop.ValueString())
					}
					if !childItem.OmpTag.IsNull() && childItem.Type.ValueString() == "ompTag" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.OmpTag.ValueInt64()))
					}
					if !childItem.OspfTag.IsNull() && childItem.Type.ValueString() == "ospfTag" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", fmt.Sprint(childItem.OspfTag.ValueInt64()))
					}
					if !childItem.Origin.IsNull() && childItem.Type.ValueString() == "origin" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Origin.ValueString())
					}
					if !childItem.Originator.IsNull() && childItem.Type.ValueString() == "originator" {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Originator.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "actions.0.parameter.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "sequences.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *RoutePolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
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
		data.Sequences = make([]RoutePolicyDefinitionSequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RoutePolicyDefinitionSequences{}
			if cValue := v.Get("sequenceId"); cValue.Exists() {
				item.Id = types.Int64Value(cValue.Int())
			} else {
				item.Id = types.Int64Null()
			}
			if cValue := v.Get("sequenceIpType"); cValue.Exists() {
				item.IpType = types.StringValue(cValue.String())
			} else {
				item.IpType = types.StringNull()
			}
			if cValue := v.Get("sequenceName"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("baseAction"); cValue.Exists() {
				item.BaseAction = types.StringValue(cValue.String())
			} else {
				item.BaseAction = types.StringNull()
			}
			if cValue := v.Get("match.entries"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.MatchEntries = make([]RoutePolicyDefinitionSequencesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := RoutePolicyDefinitionSequencesMatchEntries{}
					if ccValue := cv.Get("field"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "address" {
						cItem.PrefixListId = types.StringValue(ccValue.String())
					} else {
						cItem.PrefixListId = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "asPath" {
						cItem.AsPathListId = types.StringValue(ccValue.String())
					} else {
						cItem.AsPathListId = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "community" {
						cItem.CommunityListId = types.StringValue(ccValue.String())
					} else {
						cItem.CommunityListId = types.StringNull()
					}
					if ccValue := cv.Get("matchFlag"); ccValue.Exists() {
						cItem.CommunityListMatchFlag = types.StringValue(ccValue.String())
					} else {
						cItem.CommunityListMatchFlag = types.StringNull()
					}
					if ccValue := cv.Get("refs"); ccValue.Exists() && cItem.Type.ValueString() == "advancedCommunity" {
						cItem.CommunityListIds = helpers.GetStringSet(ccValue.Array())
					} else {
						cItem.CommunityListIds = types.SetNull(types.StringType)
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "expandedCommunity" {
						cItem.ExpandedCommunityListId = types.StringValue(ccValue.String())
					} else {
						cItem.ExpandedCommunityListId = types.StringNull()
					}
					if ccValue := cv.Get("vipVariableName"); ccValue.Exists() && cItem.Type.ValueString() == "expandedCommunityInline" {
						cItem.ExpandedCommunityListVariable = types.StringValue(ccValue.String())
					} else {
						cItem.ExpandedCommunityListVariable = types.StringNull()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "extCommunity" {
						cItem.ExtendedCommunityListId = types.StringValue(ccValue.String())
					} else {
						cItem.ExtendedCommunityListId = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "localPreference" {
						cItem.LocalPreference = types.Int64Value(ccValue.Int())
					} else {
						cItem.LocalPreference = types.Int64Null()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "metric" {
						cItem.Metric = types.Int64Value(ccValue.Int())
					} else {
						cItem.Metric = types.Int64Null()
					}
					if ccValue := cv.Get("ref"); ccValue.Exists() && cItem.Type.ValueString() == "nextHop" {
						cItem.NextHopPrefixListId = types.StringValue(ccValue.String())
					} else {
						cItem.NextHopPrefixListId = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "origin" {
						cItem.Origin = types.StringValue(ccValue.String())
					} else {
						cItem.Origin = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "peer" {
						cItem.Peer = types.StringValue(ccValue.String())
					} else {
						cItem.Peer = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "ompTag" {
						cItem.OmpTag = types.Int64Value(ccValue.Int())
					} else {
						cItem.OmpTag = types.Int64Null()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "ospfTag" {
						cItem.OspfTag = types.Int64Value(ccValue.Int())
					} else {
						cItem.OspfTag = types.Int64Null()
					}
					item.MatchEntries = append(item.MatchEntries, cItem)
					return true
				})
			} else {
				if len(item.MatchEntries) > 0 {
					item.MatchEntries = []RoutePolicyDefinitionSequencesMatchEntries{}
				}
			}
			if cValue := v.Get("actions.0.parameter"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.ActionEntries = make([]RoutePolicyDefinitionSequencesActionEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := RoutePolicyDefinitionSequencesActionEntries{}
					if ccValue := cv.Get("field"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					if ccValue := cv.Get("value.aggregator"); ccValue.Exists() && cItem.Type.ValueString() == "aggregator" {
						cItem.Aggregator = types.Int64Value(ccValue.Int())
					} else {
						cItem.Aggregator = types.Int64Null()
					}
					if ccValue := cv.Get("value.ipAddress"); ccValue.Exists() && cItem.Type.ValueString() == "aggregator" {
						cItem.AggregatorIpAddress = types.StringValue(ccValue.String())
					} else {
						cItem.AggregatorIpAddress = types.StringNull()
					}
					if ccValue := cv.Get("value.prepend"); ccValue.Exists() && cItem.Type.ValueString() == "asPath" {
						cItem.AsPathPrepend = types.StringValue(ccValue.String())
					} else {
						cItem.AsPathPrepend = types.StringNull()
					}
					if ccValue := cv.Get("value.exclude"); ccValue.Exists() && cItem.Type.ValueString() == "asPath" {
						cItem.AsPathExclude = types.StringValue(ccValue.String())
					} else {
						cItem.AsPathExclude = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "atomicAggregate" {
						if false && ccValue.String() == "" {
							cItem.AtomicAggregate = types.BoolValue(true)
						} else {
							cItem.AtomicAggregate = types.BoolValue(ccValue.Bool())
						}
					} else {
						cItem.AtomicAggregate = types.BoolNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "community" {
						cItem.Community = types.StringValue(ccValue.String())
					} else {
						cItem.Community = types.StringNull()
					}
					if ccValue := cv.Get("vipVariableName"); ccValue.Exists() && cItem.Type.ValueString() == "community" {
						cItem.CommunityVariable = types.StringValue(ccValue.String())
					} else {
						cItem.CommunityVariable = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "communityAdditive" {
						if false && ccValue.String() == "" {
							cItem.CommunityAdditive = types.BoolValue(true)
						} else {
							cItem.CommunityAdditive = types.BoolValue(ccValue.Bool())
						}
					} else {
						cItem.CommunityAdditive = types.BoolNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "localPreference" {
						cItem.LocalPreference = types.Int64Value(ccValue.Int())
					} else {
						cItem.LocalPreference = types.Int64Null()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "metric" {
						cItem.Metric = types.Int64Value(ccValue.Int())
					} else {
						cItem.Metric = types.Int64Null()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "weight" {
						cItem.Weight = types.Int64Value(ccValue.Int())
					} else {
						cItem.Weight = types.Int64Null()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "metricType" {
						cItem.MetricType = types.StringValue(ccValue.String())
					} else {
						cItem.MetricType = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "nextHop" {
						cItem.NextHop = types.StringValue(ccValue.String())
					} else {
						cItem.NextHop = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "ompTag" {
						cItem.OmpTag = types.Int64Value(ccValue.Int())
					} else {
						cItem.OmpTag = types.Int64Null()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() && cItem.Type.ValueString() == "ospfTag" {
						cItem.OspfTag = types.Int64Value(ccValue.Int())
					} else {
						cItem.OspfTag = types.Int64Null()
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
					item.ActionEntries = append(item.ActionEntries, cItem)
					return true
				})
			} else {
				if len(item.ActionEntries) > 0 {
					item.ActionEntries = []RoutePolicyDefinitionSequencesActionEntries{}
				}
			}
			data.Sequences = append(data.Sequences, item)
			return true
		})
	} else {
		if len(data.Sequences) > 0 {
			data.Sequences = []RoutePolicyDefinitionSequences{}
		}
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *RoutePolicyDefinition) hasChanges(ctx context.Context, state *RoutePolicyDefinition) bool {
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
			if !data.Sequences[i].IpType.Equal(state.Sequences[i].IpType) {
				hasChanges = true
			}
			if !data.Sequences[i].Name.Equal(state.Sequences[i].Name) {
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
					if !data.Sequences[i].MatchEntries[ii].PrefixListId.Equal(state.Sequences[i].MatchEntries[ii].PrefixListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].AsPathListId.Equal(state.Sequences[i].MatchEntries[ii].AsPathListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].CommunityListId.Equal(state.Sequences[i].MatchEntries[ii].CommunityListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].CommunityListMatchFlag.Equal(state.Sequences[i].MatchEntries[ii].CommunityListMatchFlag) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].CommunityListIds.Equal(state.Sequences[i].MatchEntries[ii].CommunityListIds) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].ExpandedCommunityListId.Equal(state.Sequences[i].MatchEntries[ii].ExpandedCommunityListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].ExpandedCommunityListVariable.Equal(state.Sequences[i].MatchEntries[ii].ExpandedCommunityListVariable) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].ExtendedCommunityListId.Equal(state.Sequences[i].MatchEntries[ii].ExtendedCommunityListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].LocalPreference.Equal(state.Sequences[i].MatchEntries[ii].LocalPreference) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Metric.Equal(state.Sequences[i].MatchEntries[ii].Metric) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].NextHopPrefixListId.Equal(state.Sequences[i].MatchEntries[ii].NextHopPrefixListId) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Origin.Equal(state.Sequences[i].MatchEntries[ii].Origin) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].Peer.Equal(state.Sequences[i].MatchEntries[ii].Peer) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].OmpTag.Equal(state.Sequences[i].MatchEntries[ii].OmpTag) {
						hasChanges = true
					}
					if !data.Sequences[i].MatchEntries[ii].OspfTag.Equal(state.Sequences[i].MatchEntries[ii].OspfTag) {
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
					if !data.Sequences[i].ActionEntries[ii].Aggregator.Equal(state.Sequences[i].ActionEntries[ii].Aggregator) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].AggregatorIpAddress.Equal(state.Sequences[i].ActionEntries[ii].AggregatorIpAddress) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].AsPathPrepend.Equal(state.Sequences[i].ActionEntries[ii].AsPathPrepend) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].AsPathExclude.Equal(state.Sequences[i].ActionEntries[ii].AsPathExclude) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].AtomicAggregate.Equal(state.Sequences[i].ActionEntries[ii].AtomicAggregate) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].Community.Equal(state.Sequences[i].ActionEntries[ii].Community) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].CommunityVariable.Equal(state.Sequences[i].ActionEntries[ii].CommunityVariable) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].CommunityAdditive.Equal(state.Sequences[i].ActionEntries[ii].CommunityAdditive) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].LocalPreference.Equal(state.Sequences[i].ActionEntries[ii].LocalPreference) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].Metric.Equal(state.Sequences[i].ActionEntries[ii].Metric) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].Weight.Equal(state.Sequences[i].ActionEntries[ii].Weight) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].MetricType.Equal(state.Sequences[i].ActionEntries[ii].MetricType) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].NextHop.Equal(state.Sequences[i].ActionEntries[ii].NextHop) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].OmpTag.Equal(state.Sequences[i].ActionEntries[ii].OmpTag) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].OspfTag.Equal(state.Sequences[i].ActionEntries[ii].OspfTag) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].Origin.Equal(state.Sequences[i].ActionEntries[ii].Origin) {
						hasChanges = true
					}
					if !data.Sequences[i].ActionEntries[ii].Originator.Equal(state.Sequences[i].ActionEntries[ii].Originator) {
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

func (data *RoutePolicyDefinition) updateVersions(ctx context.Context, state *RoutePolicyDefinition) {
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
				data.Sequences[i].MatchEntries[ii].PrefixListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].PrefixListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].PrefixListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].AsPathListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].AsPathListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].AsPathListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].CommunityListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].CommunityListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].CommunityListVersion = types.Int64Null()
			}
			if cStateIndex > -1 && !state.Sequences[stateIndex].MatchEntries[cStateIndex].CommunityListVersions.IsNull() {
				data.Sequences[i].MatchEntries[ii].CommunityListVersions = state.Sequences[stateIndex].MatchEntries[cStateIndex].CommunityListVersions
			} else {
				data.Sequences[i].MatchEntries[ii].CommunityListVersions = types.ListNull(types.StringType)
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].ExpandedCommunityListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].ExpandedCommunityListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].ExpandedCommunityListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].ExtendedCommunityListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].ExtendedCommunityListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].ExtendedCommunityListVersion = types.Int64Null()
			}
			if cStateIndex > -1 {
				data.Sequences[i].MatchEntries[ii].NextHopPrefixListVersion = state.Sequences[stateIndex].MatchEntries[cStateIndex].NextHopPrefixListVersion
			} else {
				data.Sequences[i].MatchEntries[ii].NextHopPrefixListVersion = types.Int64Null()
			}
		}
	}
}

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *RoutePolicyDefinition) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
	data.Type = types.StringValue("vedgeRoute")
	for i := range data.Sequences {
		for ii := range data.Sequences[i].MatchEntries {
			if data.Sequences[i].MatchEntries[ii].PrefixListId != types.StringNull() {
				data.Sequences[i].MatchEntries[ii].PrefixListVersion = types.Int64Value(0)
			}
			if data.Sequences[i].MatchEntries[ii].AsPathListId != types.StringNull() {
				data.Sequences[i].MatchEntries[ii].AsPathListVersion = types.Int64Value(0)
			}
			if data.Sequences[i].MatchEntries[ii].CommunityListId != types.StringNull() {
				data.Sequences[i].MatchEntries[ii].CommunityListVersion = types.Int64Value(0)
			}
			if !data.Sequences[i].MatchEntries[ii].CommunityListIds.IsNull() {
				data.Sequences[i].MatchEntries[ii].CommunityListVersions = types.ListNull(types.StringType)
			}
			if data.Sequences[i].MatchEntries[ii].ExpandedCommunityListId != types.StringNull() {
				data.Sequences[i].MatchEntries[ii].ExpandedCommunityListVersion = types.Int64Value(0)
			}
			if data.Sequences[i].MatchEntries[ii].ExtendedCommunityListId != types.StringNull() {
				data.Sequences[i].MatchEntries[ii].ExtendedCommunityListVersion = types.Int64Value(0)
			}
			if data.Sequences[i].MatchEntries[ii].NextHopPrefixListId != types.StringNull() {
				data.Sequences[i].MatchEntries[ii].NextHopPrefixListVersion = types.Int64Value(0)
			}
		}
	}
}

// End of section. //template:end processImport
