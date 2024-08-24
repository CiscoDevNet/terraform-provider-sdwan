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
type TransportRoutingOSPFv3IPv6 struct {
	Id                                 types.String                              `tfsdk:"id"`
	Version                            types.Int64                               `tfsdk:"version"`
	Name                               types.String                              `tfsdk:"name"`
	Description                        types.String                              `tfsdk:"description"`
	FeatureProfileId                   types.String                              `tfsdk:"feature_profile_id"`
	RouterId                           types.String                              `tfsdk:"router_id"`
	RouterIdVariable                   types.String                              `tfsdk:"router_id_variable"`
	Distance                           types.Int64                               `tfsdk:"distance"`
	DistanceVariable                   types.String                              `tfsdk:"distance_variable"`
	DistanceForExternalRoutes          types.Int64                               `tfsdk:"distance_for_external_routes"`
	DistanceForExternalRoutesVariable  types.String                              `tfsdk:"distance_for_external_routes_variable"`
	DistanceForInterAreaRoutes         types.Int64                               `tfsdk:"distance_for_inter_area_routes"`
	DistanceForInterAreaRoutesVariable types.String                              `tfsdk:"distance_for_inter_area_routes_variable"`
	DistanceForIntraAreaRoutes         types.Int64                               `tfsdk:"distance_for_intra_area_routes"`
	DistanceForIntraAreaRoutesVariable types.String                              `tfsdk:"distance_for_intra_area_routes_variable"`
	ReferenceBandwidth                 types.Int64                               `tfsdk:"reference_bandwidth"`
	ReferenceBandwidthVariable         types.String                              `tfsdk:"reference_bandwidth_variable"`
	Rfc1583Compatible                  types.Bool                                `tfsdk:"rfc_1583_compatible"`
	Rfc1583CompatibleVariable          types.String                              `tfsdk:"rfc_1583_compatible_variable"`
	Originate                          types.Bool                                `tfsdk:"originate"`
	Always                             types.Bool                                `tfsdk:"always"`
	AlwaysVariable                     types.String                              `tfsdk:"always_variable"`
	Metric                             types.Int64                               `tfsdk:"metric"`
	MetricVariable                     types.String                              `tfsdk:"metric_variable"`
	MetricType                         types.String                              `tfsdk:"metric_type"`
	MetricTypeVariable                 types.String                              `tfsdk:"metric_type_variable"`
	SpfCalculationDeplay               types.Int64                               `tfsdk:"spf_calculation_deplay"`
	SpfCalculationDeplayVariable       types.String                              `tfsdk:"spf_calculation_deplay_variable"`
	InitialHoldTime                    types.Int64                               `tfsdk:"initial_hold_time"`
	InitialHoldTimeVariable            types.String                              `tfsdk:"initial_hold_time_variable"`
	MaximumHoldTime                    types.Int64                               `tfsdk:"maximum_hold_time"`
	MaximumHoldTimeVariable            types.String                              `tfsdk:"maximum_hold_time_variable"`
	RoutePolicyId                      types.String                              `tfsdk:"route_policy_id"`
	Filter                             types.Bool                                `tfsdk:"filter"`
	FilterVariable                     types.String                              `tfsdk:"filter_variable"`
	Redistributes                      []TransportRoutingOSPFv3IPv6Redistributes `tfsdk:"redistributes"`
	RouterLsaAction                    types.String                              `tfsdk:"router_lsa_action"`
	RouterLsaOnStartuPTime             types.Int64                               `tfsdk:"router_lsa_on_startu_p_time"`
	RouterLsaOnStartuPTimeVariable     types.String                              `tfsdk:"router_lsa_on_startu_p_time_variable"`
	Areas                              []TransportRoutingOSPFv3IPv6Areas         `tfsdk:"areas"`
}

type TransportRoutingOSPFv3IPv6Redistributes struct {
	Protocol         types.String `tfsdk:"protocol"`
	ProtocolVariable types.String `tfsdk:"protocol_variable"`
	RoutePolicyId    types.String `tfsdk:"route_policy_id"`
}

type TransportRoutingOSPFv3IPv6Areas struct {
	AreaNumber              types.Int64                                 `tfsdk:"area_number"`
	AreaNumberVariable      types.String                                `tfsdk:"area_number_variable"`
	AreaType                types.String                                `tfsdk:"area_type"`
	NoSummary               types.Bool                                  `tfsdk:"no_summary"`
	NoSummaryVariable       types.String                                `tfsdk:"no_summary_variable"`
	AlwaysTranslate         types.Bool                                  `tfsdk:"always_translate"`
	AlwaysTranslateVariable types.String                                `tfsdk:"always_translate_variable"`
	Interfaces              []TransportRoutingOSPFv3IPv6AreasInterfaces `tfsdk:"interfaces"`
	Ranges                  []TransportRoutingOSPFv3IPv6AreasRanges     `tfsdk:"ranges"`
}

type TransportRoutingOSPFv3IPv6AreasInterfaces struct {
	IfName                        types.String `tfsdk:"if_name"`
	IfNameVariable                types.String `tfsdk:"if_name_variable"`
	HelloInterval                 types.Int64  `tfsdk:"hello_interval"`
	HelloIntervalVariable         types.String `tfsdk:"hello_interval_variable"`
	DeadInterval                  types.Int64  `tfsdk:"dead_interval"`
	DeadIntervalVariable          types.String `tfsdk:"dead_interval_variable"`
	LsaRetransmitInterval         types.Int64  `tfsdk:"lsa_retransmit_interval"`
	LsaRetransmitIntervalVariable types.String `tfsdk:"lsa_retransmit_interval_variable"`
	InterfaceCost                 types.Int64  `tfsdk:"interface_cost"`
	InterfaceCostVariable         types.String `tfsdk:"interface_cost_variable"`
	OspfNetworkType               types.String `tfsdk:"ospf_network_type"`
	OspfNetworkTypeVariable       types.String `tfsdk:"ospf_network_type_variable"`
	PassiveInterface              types.Bool   `tfsdk:"passive_interface"`
	PassiveInterfaceVariable      types.String `tfsdk:"passive_interface_variable"`
	AuthType                      types.String `tfsdk:"auth_type"`
	Spi                           types.Int64  `tfsdk:"spi"`
	SpiVariable                   types.String `tfsdk:"spi_variable"`
	AuthKey                       types.String `tfsdk:"auth_key"`
	AuthKeyVariable               types.String `tfsdk:"auth_key_variable"`
}
type TransportRoutingOSPFv3IPv6AreasRanges struct {
	Network             types.String `tfsdk:"network"`
	NetworkVariable     types.String `tfsdk:"network_variable"`
	Cost                types.Int64  `tfsdk:"cost"`
	CostVariable        types.String `tfsdk:"cost_variable"`
	NoAdvertise         types.Bool   `tfsdk:"no_advertise"`
	NoAdvertiseVariable types.String `tfsdk:"no_advertise_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TransportRoutingOSPFv3IPv6) getModel() string {
	return "transport_routing_ospfv3_ipv6"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportRoutingOSPFv3IPv6) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/routing/ospfv3/ipv6", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportRoutingOSPFv3IPv6) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.RouterIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.routerId.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.routerId.value", data.RouterIdVariable.ValueString())
	} else if data.RouterId.IsNull() {
		body, _ = sjson.Set(body, path+"basic.routerId.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"basic.routerId.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.routerId.value", data.RouterId.ValueString())
	}

	if !data.DistanceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.distance.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.distance.value", data.DistanceVariable.ValueString())
	} else if data.Distance.IsNull() {
		body, _ = sjson.Set(body, path+"basic.distance.optionType", "default")
		body, _ = sjson.Set(body, path+"basic.distance.value", 110)
	} else {
		body, _ = sjson.Set(body, path+"basic.distance.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.distance.value", data.Distance.ValueInt64())
	}

	if !data.DistanceForExternalRoutesVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.externalDistance.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.externalDistance.value", data.DistanceForExternalRoutesVariable.ValueString())
	} else if data.DistanceForExternalRoutes.IsNull() {
		body, _ = sjson.Set(body, path+"basic.externalDistance.optionType", "default")
		body, _ = sjson.Set(body, path+"basic.externalDistance.value", 110)
	} else {
		body, _ = sjson.Set(body, path+"basic.externalDistance.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.externalDistance.value", data.DistanceForExternalRoutes.ValueInt64())
	}

	if !data.DistanceForInterAreaRoutesVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.interAreaDistance.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.interAreaDistance.value", data.DistanceForInterAreaRoutesVariable.ValueString())
	} else if data.DistanceForInterAreaRoutes.IsNull() {
		body, _ = sjson.Set(body, path+"basic.interAreaDistance.optionType", "default")
		body, _ = sjson.Set(body, path+"basic.interAreaDistance.value", 110)
	} else {
		body, _ = sjson.Set(body, path+"basic.interAreaDistance.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.interAreaDistance.value", data.DistanceForInterAreaRoutes.ValueInt64())
	}

	if !data.DistanceForIntraAreaRoutesVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.intraAreaDistance.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.intraAreaDistance.value", data.DistanceForIntraAreaRoutesVariable.ValueString())
	} else if data.DistanceForIntraAreaRoutes.IsNull() {
		body, _ = sjson.Set(body, path+"basic.intraAreaDistance.optionType", "default")
		body, _ = sjson.Set(body, path+"basic.intraAreaDistance.value", 110)
	} else {
		body, _ = sjson.Set(body, path+"basic.intraAreaDistance.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.intraAreaDistance.value", data.DistanceForIntraAreaRoutes.ValueInt64())
	}

	if !data.ReferenceBandwidthVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.referenceBandwidth.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.referenceBandwidth.value", data.ReferenceBandwidthVariable.ValueString())
	} else if data.ReferenceBandwidth.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.referenceBandwidth.optionType", "default")
		body, _ = sjson.Set(body, path+"advanced.referenceBandwidth.value", 100)
	} else {
		body, _ = sjson.Set(body, path+"advanced.referenceBandwidth.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.referenceBandwidth.value", data.ReferenceBandwidth.ValueInt64())
	}

	if !data.Rfc1583CompatibleVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.compatibleRfc1583.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.compatibleRfc1583.value", data.Rfc1583CompatibleVariable.ValueString())
	} else if data.Rfc1583Compatible.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.compatibleRfc1583.optionType", "default")
		body, _ = sjson.Set(body, path+"advanced.compatibleRfc1583.value", true)
	} else {
		body, _ = sjson.Set(body, path+"advanced.compatibleRfc1583.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.compatibleRfc1583.value", data.Rfc1583Compatible.ValueBool())
	}
	if !data.Originate.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.defaultOriginate.originate.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.defaultOriginate.originate.value", data.Originate.ValueBool())
	}

	if !data.AlwaysVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.defaultOriginate.always.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.defaultOriginate.always.value", data.AlwaysVariable.ValueString())
	} else if !data.Always.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.defaultOriginate.always.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.defaultOriginate.always.value", data.Always.ValueBool())
	}

	if !data.MetricVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.defaultOriginate.metric.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.defaultOriginate.metric.value", data.MetricVariable.ValueString())
	} else if !data.Metric.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.defaultOriginate.metric.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.defaultOriginate.metric.value", data.Metric.ValueInt64())
	}

	if !data.MetricTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.defaultOriginate.metricType.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.defaultOriginate.metricType.value", data.MetricTypeVariable.ValueString())
	} else if !data.MetricType.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.defaultOriginate.metricType.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.defaultOriginate.metricType.value", data.MetricType.ValueString())
	}

	if !data.SpfCalculationDeplayVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.spfTimers.delay.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.spfTimers.delay.value", data.SpfCalculationDeplayVariable.ValueString())
	} else if data.SpfCalculationDeplay.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.spfTimers.delay.optionType", "default")
		body, _ = sjson.Set(body, path+"advanced.spfTimers.delay.value", 200)
	} else {
		body, _ = sjson.Set(body, path+"advanced.spfTimers.delay.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.spfTimers.delay.value", data.SpfCalculationDeplay.ValueInt64())
	}

	if !data.InitialHoldTimeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.spfTimers.initialHold.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.spfTimers.initialHold.value", data.InitialHoldTimeVariable.ValueString())
	} else if data.InitialHoldTime.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.spfTimers.initialHold.optionType", "default")
		body, _ = sjson.Set(body, path+"advanced.spfTimers.initialHold.value", 1000)
	} else {
		body, _ = sjson.Set(body, path+"advanced.spfTimers.initialHold.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.spfTimers.initialHold.value", data.InitialHoldTime.ValueInt64())
	}

	if !data.MaximumHoldTimeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.spfTimers.maxHold.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.spfTimers.maxHold.value", data.MaximumHoldTimeVariable.ValueString())
	} else if data.MaximumHoldTime.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.spfTimers.maxHold.optionType", "default")
		body, _ = sjson.Set(body, path+"advanced.spfTimers.maxHold.value", 10000)
	} else {
		body, _ = sjson.Set(body, path+"advanced.spfTimers.maxHold.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.spfTimers.maxHold.value", data.MaximumHoldTime.ValueInt64())
	}
	if !data.RoutePolicyId.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.policyName.refId.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.policyName.refId.value", data.RoutePolicyId.ValueString())
	}

	if !data.FilterVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.filter.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.filter.value", data.FilterVariable.ValueString())
	} else if data.Filter.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.filter.optionType", "default")
		body, _ = sjson.Set(body, path+"advanced.filter.value", false)
	} else {
		body, _ = sjson.Set(body, path+"advanced.filter.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.filter.value", data.Filter.ValueBool())
	}
	body, _ = sjson.Set(body, path+"redistribute", []interface{}{})
	for _, item := range data.Redistributes {
		itemBody := ""

		if !item.ProtocolVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "protocol.value", item.ProtocolVariable.ValueString())
		} else if !item.Protocol.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "protocol.value", item.Protocol.ValueString())
		}
		if !item.RoutePolicyId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.value", item.RoutePolicyId.ValueString())
		}
		body, _ = sjson.SetRaw(body, path+"redistribute.-1", itemBody)
	}
	if !data.RouterLsaAction.IsNull() {
		body, _ = sjson.Set(body, path+"maxMetricRouterLsa.action.optionType", "global")
		body, _ = sjson.Set(body, path+"maxMetricRouterLsa.action.value", data.RouterLsaAction.ValueString())
	}

	if !data.RouterLsaOnStartuPTimeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"maxMetricRouterLsa.onStartUpTime.optionType", "variable")
		body, _ = sjson.Set(body, path+"maxMetricRouterLsa.onStartUpTime.value", data.RouterLsaOnStartuPTimeVariable.ValueString())
	} else if !data.RouterLsaOnStartuPTime.IsNull() {
		body, _ = sjson.Set(body, path+"maxMetricRouterLsa.onStartUpTime.optionType", "global")
		body, _ = sjson.Set(body, path+"maxMetricRouterLsa.onStartUpTime.value", data.RouterLsaOnStartuPTime.ValueInt64())
	}

	for _, item := range data.Areas {
		itemBody := ""

		if !item.AreaNumberVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "areaNum.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "areaNum.value", item.AreaNumberVariable.ValueString())
		} else if !item.AreaNumber.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "areaNum.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "areaNum.value", item.AreaNumber.ValueInt64())
		}
		if !item.AreaType.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.areaType.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.areaType.value", item.AreaType.ValueString())
		}

		if !item.NoSummaryVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.noSummary.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.noSummary.value", item.NoSummaryVariable.ValueString())
		} else if !item.NoSummary.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.noSummary.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.noSummary.value", item.NoSummary.ValueBool())
		}

		if !item.AlwaysTranslateVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.alwaysTranslate.optionType", "variable")
			itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.alwaysTranslate.value", item.AlwaysTranslateVariable.ValueString())
		} else if !item.AlwaysTranslate.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.alwaysTranslate.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.alwaysTranslate.value", item.AlwaysTranslate.ValueBool())
		}

		for _, childItem := range item.Interfaces {
			itemChildBody := ""

			if !childItem.IfNameVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "ifName.optionType", "variable")
				itemChildBody, _ = sjson.Set(itemChildBody, "ifName.value", childItem.IfNameVariable.ValueString())
			} else if !childItem.IfName.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "ifName.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "ifName.value", childItem.IfName.ValueString())
			}

			if !childItem.HelloIntervalVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "helloInterval.optionType", "variable")
				itemChildBody, _ = sjson.Set(itemChildBody, "helloInterval.value", childItem.HelloIntervalVariable.ValueString())
			} else if childItem.HelloInterval.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "helloInterval.optionType", "default")
				itemChildBody, _ = sjson.Set(itemChildBody, "helloInterval.value", 10)
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "helloInterval.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "helloInterval.value", childItem.HelloInterval.ValueInt64())
			}

			if !childItem.DeadIntervalVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "deadInterval.optionType", "variable")
				itemChildBody, _ = sjson.Set(itemChildBody, "deadInterval.value", childItem.DeadIntervalVariable.ValueString())
			} else if childItem.DeadInterval.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "deadInterval.optionType", "default")
				itemChildBody, _ = sjson.Set(itemChildBody, "deadInterval.value", 40)
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "deadInterval.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "deadInterval.value", childItem.DeadInterval.ValueInt64())
			}

			if !childItem.LsaRetransmitIntervalVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmitInterval.optionType", "variable")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmitInterval.value", childItem.LsaRetransmitIntervalVariable.ValueString())
			} else if childItem.LsaRetransmitInterval.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmitInterval.optionType", "default")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmitInterval.value", 5)
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmitInterval.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmitInterval.value", childItem.LsaRetransmitInterval.ValueInt64())
			}

			if !childItem.InterfaceCostVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost.optionType", "variable")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost.value", childItem.InterfaceCostVariable.ValueString())
			} else if childItem.InterfaceCost.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost.optionType", "default")

			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost.value", childItem.InterfaceCost.ValueInt64())
			}

			if !childItem.OspfNetworkTypeVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "networkType.optionType", "variable")
				itemChildBody, _ = sjson.Set(itemChildBody, "networkType.value", childItem.OspfNetworkTypeVariable.ValueString())
			} else if childItem.OspfNetworkType.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "networkType.optionType", "default")

			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "networkType.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "networkType.value", childItem.OspfNetworkType.ValueString())
			}

			if !childItem.PassiveInterfaceVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "passiveInterface.optionType", "variable")
				itemChildBody, _ = sjson.Set(itemChildBody, "passiveInterface.value", childItem.PassiveInterfaceVariable.ValueString())
			} else if childItem.PassiveInterface.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "passiveInterface.optionType", "default")
				itemChildBody, _ = sjson.Set(itemChildBody, "passiveInterface.value", false)
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "passiveInterface.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "passiveInterface.value", childItem.PassiveInterface.ValueBool())
			}
			if !childItem.AuthType.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.authType.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.authType.value", childItem.AuthType.ValueString())
			}

			if !childItem.SpiVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.spi.optionType", "variable")
				itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.spi.value", childItem.SpiVariable.ValueString())
			} else if !childItem.Spi.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.spi.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.spi.value", childItem.Spi.ValueInt64())
			}

			if !childItem.AuthKeyVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.authKey.optionType", "variable")
				itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.authKey.value", childItem.AuthKeyVariable.ValueString())
			} else if !childItem.AuthKey.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.authKey.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.authKey.value", childItem.AuthKey.ValueString())
			}
			itemBody, _ = sjson.SetRaw(itemBody, "interfaces.-1", itemChildBody)
		}
		itemBody, _ = sjson.Set(itemBody, "ranges", []interface{}{})
		for _, childItem := range item.Ranges {
			itemChildBody := ""

			if !childItem.NetworkVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "network.optionType", "variable")
				itemChildBody, _ = sjson.Set(itemChildBody, "network.value", childItem.NetworkVariable.ValueString())
			} else if !childItem.Network.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "network.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "network.value", childItem.Network.ValueString())
			}

			if !childItem.CostVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost.optionType", "variable")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost.value", childItem.CostVariable.ValueString())
			} else if childItem.Cost.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost.optionType", "default")

			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost.value", childItem.Cost.ValueInt64())
			}

			if !childItem.NoAdvertiseVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "noAdvertise.optionType", "variable")
				itemChildBody, _ = sjson.Set(itemChildBody, "noAdvertise.value", childItem.NoAdvertiseVariable.ValueString())
			} else if childItem.NoAdvertise.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "noAdvertise.optionType", "default")
				itemChildBody, _ = sjson.Set(itemChildBody, "noAdvertise.value", false)
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "noAdvertise.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "noAdvertise.value", childItem.NoAdvertise.ValueBool())
			}
			itemBody, _ = sjson.SetRaw(itemBody, "ranges.-1", itemChildBody)
		}
		body, _ = sjson.SetRaw(body, path+"area.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportRoutingOSPFv3IPv6) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.RouterId = types.StringNull()
	data.RouterIdVariable = types.StringNull()
	if t := res.Get(path + "basic.routerId.optionType"); t.Exists() {
		va := res.Get(path + "basic.routerId.value")
		if t.String() == "variable" {
			data.RouterIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RouterId = types.StringValue(va.String())
		}
	}
	data.Distance = types.Int64Null()
	data.DistanceVariable = types.StringNull()
	if t := res.Get(path + "basic.distance.optionType"); t.Exists() {
		va := res.Get(path + "basic.distance.value")
		if t.String() == "variable" {
			data.DistanceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Distance = types.Int64Value(va.Int())
		}
	}
	data.DistanceForExternalRoutes = types.Int64Null()
	data.DistanceForExternalRoutesVariable = types.StringNull()
	if t := res.Get(path + "basic.externalDistance.optionType"); t.Exists() {
		va := res.Get(path + "basic.externalDistance.value")
		if t.String() == "variable" {
			data.DistanceForExternalRoutesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceForExternalRoutes = types.Int64Value(va.Int())
		}
	}
	data.DistanceForInterAreaRoutes = types.Int64Null()
	data.DistanceForInterAreaRoutesVariable = types.StringNull()
	if t := res.Get(path + "basic.interAreaDistance.optionType"); t.Exists() {
		va := res.Get(path + "basic.interAreaDistance.value")
		if t.String() == "variable" {
			data.DistanceForInterAreaRoutesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceForInterAreaRoutes = types.Int64Value(va.Int())
		}
	}
	data.DistanceForIntraAreaRoutes = types.Int64Null()
	data.DistanceForIntraAreaRoutesVariable = types.StringNull()
	if t := res.Get(path + "basic.intraAreaDistance.optionType"); t.Exists() {
		va := res.Get(path + "basic.intraAreaDistance.value")
		if t.String() == "variable" {
			data.DistanceForIntraAreaRoutesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceForIntraAreaRoutes = types.Int64Value(va.Int())
		}
	}
	data.ReferenceBandwidth = types.Int64Null()
	data.ReferenceBandwidthVariable = types.StringNull()
	if t := res.Get(path + "advanced.referenceBandwidth.optionType"); t.Exists() {
		va := res.Get(path + "advanced.referenceBandwidth.value")
		if t.String() == "variable" {
			data.ReferenceBandwidthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ReferenceBandwidth = types.Int64Value(va.Int())
		}
	}
	data.Rfc1583Compatible = types.BoolNull()
	data.Rfc1583CompatibleVariable = types.StringNull()
	if t := res.Get(path + "advanced.compatibleRfc1583.optionType"); t.Exists() {
		va := res.Get(path + "advanced.compatibleRfc1583.value")
		if t.String() == "variable" {
			data.Rfc1583CompatibleVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Rfc1583Compatible = types.BoolValue(va.Bool())
		}
	}
	data.Originate = types.BoolNull()

	if t := res.Get(path + "advanced.defaultOriginate.originate.optionType"); t.Exists() {
		va := res.Get(path + "advanced.defaultOriginate.originate.value")
		if t.String() == "global" {
			data.Originate = types.BoolValue(va.Bool())
		}
	}
	data.Always = types.BoolNull()
	data.AlwaysVariable = types.StringNull()
	if t := res.Get(path + "advanced.defaultOriginate.always.optionType"); t.Exists() {
		va := res.Get(path + "advanced.defaultOriginate.always.value")
		if t.String() == "variable" {
			data.AlwaysVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Always = types.BoolValue(va.Bool())
		}
	}
	data.Metric = types.Int64Null()
	data.MetricVariable = types.StringNull()
	if t := res.Get(path + "advanced.defaultOriginate.metric.optionType"); t.Exists() {
		va := res.Get(path + "advanced.defaultOriginate.metric.value")
		if t.String() == "variable" {
			data.MetricVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Metric = types.Int64Value(va.Int())
		}
	}
	data.MetricType = types.StringNull()
	data.MetricTypeVariable = types.StringNull()
	if t := res.Get(path + "advanced.defaultOriginate.metricType.optionType"); t.Exists() {
		va := res.Get(path + "advanced.defaultOriginate.metricType.value")
		if t.String() == "variable" {
			data.MetricTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MetricType = types.StringValue(va.String())
		}
	}
	data.SpfCalculationDeplay = types.Int64Null()
	data.SpfCalculationDeplayVariable = types.StringNull()
	if t := res.Get(path + "advanced.spfTimers.delay.optionType"); t.Exists() {
		va := res.Get(path + "advanced.spfTimers.delay.value")
		if t.String() == "variable" {
			data.SpfCalculationDeplayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SpfCalculationDeplay = types.Int64Value(va.Int())
		}
	}
	data.InitialHoldTime = types.Int64Null()
	data.InitialHoldTimeVariable = types.StringNull()
	if t := res.Get(path + "advanced.spfTimers.initialHold.optionType"); t.Exists() {
		va := res.Get(path + "advanced.spfTimers.initialHold.value")
		if t.String() == "variable" {
			data.InitialHoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InitialHoldTime = types.Int64Value(va.Int())
		}
	}
	data.MaximumHoldTime = types.Int64Null()
	data.MaximumHoldTimeVariable = types.StringNull()
	if t := res.Get(path + "advanced.spfTimers.maxHold.optionType"); t.Exists() {
		va := res.Get(path + "advanced.spfTimers.maxHold.value")
		if t.String() == "variable" {
			data.MaximumHoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MaximumHoldTime = types.Int64Value(va.Int())
		}
	}
	data.RoutePolicyId = types.StringNull()

	if t := res.Get(path + "advanced.policyName.refId.optionType"); t.Exists() {
		va := res.Get(path + "advanced.policyName.refId.value")
		if t.String() == "global" {
			data.RoutePolicyId = types.StringValue(va.String())
		}
	}
	data.Filter = types.BoolNull()
	data.FilterVariable = types.StringNull()
	if t := res.Get(path + "advanced.filter.optionType"); t.Exists() {
		va := res.Get(path + "advanced.filter.value")
		if t.String() == "variable" {
			data.FilterVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Filter = types.BoolValue(va.Bool())
		}
	}
	if value := res.Get(path + "redistribute"); value.Exists() {
		data.Redistributes = make([]TransportRoutingOSPFv3IPv6Redistributes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportRoutingOSPFv3IPv6Redistributes{}
			item.Protocol = types.StringNull()
			item.ProtocolVariable = types.StringNull()
			if t := v.Get("protocol.optionType"); t.Exists() {
				va := v.Get("protocol.value")
				if t.String() == "variable" {
					item.ProtocolVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Protocol = types.StringValue(va.String())
				}
			}
			item.RoutePolicyId = types.StringNull()

			if t := v.Get("routePolicy.refId.optionType"); t.Exists() {
				va := v.Get("routePolicy.refId.value")
				if t.String() == "global" {
					item.RoutePolicyId = types.StringValue(va.String())
				}
			}
			data.Redistributes = append(data.Redistributes, item)
			return true
		})
	}
	data.RouterLsaAction = types.StringNull()

	if t := res.Get(path + "maxMetricRouterLsa.action.optionType"); t.Exists() {
		va := res.Get(path + "maxMetricRouterLsa.action.value")
		if t.String() == "global" {
			data.RouterLsaAction = types.StringValue(va.String())
		}
	}
	data.RouterLsaOnStartuPTime = types.Int64Null()
	data.RouterLsaOnStartuPTimeVariable = types.StringNull()
	if t := res.Get(path + "maxMetricRouterLsa.onStartUpTime.optionType"); t.Exists() {
		va := res.Get(path + "maxMetricRouterLsa.onStartUpTime.value")
		if t.String() == "variable" {
			data.RouterLsaOnStartuPTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RouterLsaOnStartuPTime = types.Int64Value(va.Int())
		}
	}
	if value := res.Get(path + "area"); value.Exists() {
		data.Areas = make([]TransportRoutingOSPFv3IPv6Areas, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportRoutingOSPFv3IPv6Areas{}
			item.AreaNumber = types.Int64Null()
			item.AreaNumberVariable = types.StringNull()
			if t := v.Get("areaNum.optionType"); t.Exists() {
				va := v.Get("areaNum.value")
				if t.String() == "variable" {
					item.AreaNumberVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AreaNumber = types.Int64Value(va.Int())
				}
			}
			item.AreaType = types.StringNull()

			if t := v.Get("areaTypeConfig.areaType.optionType"); t.Exists() {
				va := v.Get("areaTypeConfig.areaType.value")
				if t.String() == "global" {
					item.AreaType = types.StringValue(va.String())
				}
			}
			item.NoSummary = types.BoolNull()
			item.NoSummaryVariable = types.StringNull()
			if t := v.Get("areaTypeConfig.noSummary.optionType"); t.Exists() {
				va := v.Get("areaTypeConfig.noSummary.value")
				if t.String() == "variable" {
					item.NoSummaryVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NoSummary = types.BoolValue(va.Bool())
				}
			}
			item.AlwaysTranslate = types.BoolNull()
			item.AlwaysTranslateVariable = types.StringNull()
			if t := v.Get("areaTypeConfig.alwaysTranslate.optionType"); t.Exists() {
				va := v.Get("areaTypeConfig.alwaysTranslate.value")
				if t.String() == "variable" {
					item.AlwaysTranslateVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AlwaysTranslate = types.BoolValue(va.Bool())
				}
			}
			if cValue := v.Get("interfaces"); cValue.Exists() {
				item.Interfaces = make([]TransportRoutingOSPFv3IPv6AreasInterfaces, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TransportRoutingOSPFv3IPv6AreasInterfaces{}
					cItem.IfName = types.StringNull()
					cItem.IfNameVariable = types.StringNull()
					if t := cv.Get("ifName.optionType"); t.Exists() {
						va := cv.Get("ifName.value")
						if t.String() == "variable" {
							cItem.IfNameVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.IfName = types.StringValue(va.String())
						}
					}
					cItem.HelloInterval = types.Int64Null()
					cItem.HelloIntervalVariable = types.StringNull()
					if t := cv.Get("helloInterval.optionType"); t.Exists() {
						va := cv.Get("helloInterval.value")
						if t.String() == "variable" {
							cItem.HelloIntervalVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.HelloInterval = types.Int64Value(va.Int())
						}
					}
					cItem.DeadInterval = types.Int64Null()
					cItem.DeadIntervalVariable = types.StringNull()
					if t := cv.Get("deadInterval.optionType"); t.Exists() {
						va := cv.Get("deadInterval.value")
						if t.String() == "variable" {
							cItem.DeadIntervalVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.DeadInterval = types.Int64Value(va.Int())
						}
					}
					cItem.LsaRetransmitInterval = types.Int64Null()
					cItem.LsaRetransmitIntervalVariable = types.StringNull()
					if t := cv.Get("retransmitInterval.optionType"); t.Exists() {
						va := cv.Get("retransmitInterval.value")
						if t.String() == "variable" {
							cItem.LsaRetransmitIntervalVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.LsaRetransmitInterval = types.Int64Value(va.Int())
						}
					}
					cItem.InterfaceCost = types.Int64Null()
					cItem.InterfaceCostVariable = types.StringNull()
					if t := cv.Get("cost.optionType"); t.Exists() {
						va := cv.Get("cost.value")
						if t.String() == "variable" {
							cItem.InterfaceCostVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.InterfaceCost = types.Int64Value(va.Int())
						}
					}
					cItem.OspfNetworkType = types.StringNull()
					cItem.OspfNetworkTypeVariable = types.StringNull()
					if t := cv.Get("networkType.optionType"); t.Exists() {
						va := cv.Get("networkType.value")
						if t.String() == "variable" {
							cItem.OspfNetworkTypeVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.OspfNetworkType = types.StringValue(va.String())
						}
					}
					cItem.PassiveInterface = types.BoolNull()
					cItem.PassiveInterfaceVariable = types.StringNull()
					if t := cv.Get("passiveInterface.optionType"); t.Exists() {
						va := cv.Get("passiveInterface.value")
						if t.String() == "variable" {
							cItem.PassiveInterfaceVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.PassiveInterface = types.BoolValue(va.Bool())
						}
					}
					cItem.AuthType = types.StringNull()

					if t := cv.Get("authenticationConfig.authType.optionType"); t.Exists() {
						va := cv.Get("authenticationConfig.authType.value")
						if t.String() == "global" {
							cItem.AuthType = types.StringValue(va.String())
						}
					}
					cItem.Spi = types.Int64Null()
					cItem.SpiVariable = types.StringNull()
					if t := cv.Get("authenticationConfig.spi.optionType"); t.Exists() {
						va := cv.Get("authenticationConfig.spi.value")
						if t.String() == "variable" {
							cItem.SpiVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Spi = types.Int64Value(va.Int())
						}
					}
					cItem.AuthKey = types.StringNull()
					cItem.AuthKeyVariable = types.StringNull()
					if t := cv.Get("authenticationConfig.authKey.optionType"); t.Exists() {
						va := cv.Get("authenticationConfig.authKey.value")
						if t.String() == "variable" {
							cItem.AuthKeyVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.AuthKey = types.StringValue(va.String())
						}
					}
					item.Interfaces = append(item.Interfaces, cItem)
					return true
				})
			}
			if cValue := v.Get("ranges"); cValue.Exists() {
				item.Ranges = make([]TransportRoutingOSPFv3IPv6AreasRanges, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TransportRoutingOSPFv3IPv6AreasRanges{}
					cItem.Network = types.StringNull()
					cItem.NetworkVariable = types.StringNull()
					if t := cv.Get("network.optionType"); t.Exists() {
						va := cv.Get("network.value")
						if t.String() == "variable" {
							cItem.NetworkVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Network = types.StringValue(va.String())
						}
					}
					cItem.Cost = types.Int64Null()
					cItem.CostVariable = types.StringNull()
					if t := cv.Get("cost.optionType"); t.Exists() {
						va := cv.Get("cost.value")
						if t.String() == "variable" {
							cItem.CostVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Cost = types.Int64Value(va.Int())
						}
					}
					cItem.NoAdvertise = types.BoolNull()
					cItem.NoAdvertiseVariable = types.StringNull()
					if t := cv.Get("noAdvertise.optionType"); t.Exists() {
						va := cv.Get("noAdvertise.value")
						if t.String() == "variable" {
							cItem.NoAdvertiseVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.NoAdvertise = types.BoolValue(va.Bool())
						}
					}
					item.Ranges = append(item.Ranges, cItem)
					return true
				})
			}
			data.Areas = append(data.Areas, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TransportRoutingOSPFv3IPv6) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.RouterId = types.StringNull()
	data.RouterIdVariable = types.StringNull()
	if t := res.Get(path + "basic.routerId.optionType"); t.Exists() {
		va := res.Get(path + "basic.routerId.value")
		if t.String() == "variable" {
			data.RouterIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RouterId = types.StringValue(va.String())
		}
	}
	data.Distance = types.Int64Null()
	data.DistanceVariable = types.StringNull()
	if t := res.Get(path + "basic.distance.optionType"); t.Exists() {
		va := res.Get(path + "basic.distance.value")
		if t.String() == "variable" {
			data.DistanceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Distance = types.Int64Value(va.Int())
		}
	}
	data.DistanceForExternalRoutes = types.Int64Null()
	data.DistanceForExternalRoutesVariable = types.StringNull()
	if t := res.Get(path + "basic.externalDistance.optionType"); t.Exists() {
		va := res.Get(path + "basic.externalDistance.value")
		if t.String() == "variable" {
			data.DistanceForExternalRoutesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceForExternalRoutes = types.Int64Value(va.Int())
		}
	}
	data.DistanceForInterAreaRoutes = types.Int64Null()
	data.DistanceForInterAreaRoutesVariable = types.StringNull()
	if t := res.Get(path + "basic.interAreaDistance.optionType"); t.Exists() {
		va := res.Get(path + "basic.interAreaDistance.value")
		if t.String() == "variable" {
			data.DistanceForInterAreaRoutesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceForInterAreaRoutes = types.Int64Value(va.Int())
		}
	}
	data.DistanceForIntraAreaRoutes = types.Int64Null()
	data.DistanceForIntraAreaRoutesVariable = types.StringNull()
	if t := res.Get(path + "basic.intraAreaDistance.optionType"); t.Exists() {
		va := res.Get(path + "basic.intraAreaDistance.value")
		if t.String() == "variable" {
			data.DistanceForIntraAreaRoutesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceForIntraAreaRoutes = types.Int64Value(va.Int())
		}
	}
	data.ReferenceBandwidth = types.Int64Null()
	data.ReferenceBandwidthVariable = types.StringNull()
	if t := res.Get(path + "advanced.referenceBandwidth.optionType"); t.Exists() {
		va := res.Get(path + "advanced.referenceBandwidth.value")
		if t.String() == "variable" {
			data.ReferenceBandwidthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ReferenceBandwidth = types.Int64Value(va.Int())
		}
	}
	data.Rfc1583Compatible = types.BoolNull()
	data.Rfc1583CompatibleVariable = types.StringNull()
	if t := res.Get(path + "advanced.compatibleRfc1583.optionType"); t.Exists() {
		va := res.Get(path + "advanced.compatibleRfc1583.value")
		if t.String() == "variable" {
			data.Rfc1583CompatibleVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Rfc1583Compatible = types.BoolValue(va.Bool())
		}
	}
	data.Originate = types.BoolNull()

	if t := res.Get(path + "advanced.defaultOriginate.originate.optionType"); t.Exists() {
		va := res.Get(path + "advanced.defaultOriginate.originate.value")
		if t.String() == "global" {
			data.Originate = types.BoolValue(va.Bool())
		}
	}
	data.Always = types.BoolNull()
	data.AlwaysVariable = types.StringNull()
	if t := res.Get(path + "advanced.defaultOriginate.always.optionType"); t.Exists() {
		va := res.Get(path + "advanced.defaultOriginate.always.value")
		if t.String() == "variable" {
			data.AlwaysVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Always = types.BoolValue(va.Bool())
		}
	}
	data.Metric = types.Int64Null()
	data.MetricVariable = types.StringNull()
	if t := res.Get(path + "advanced.defaultOriginate.metric.optionType"); t.Exists() {
		va := res.Get(path + "advanced.defaultOriginate.metric.value")
		if t.String() == "variable" {
			data.MetricVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Metric = types.Int64Value(va.Int())
		}
	}
	data.MetricType = types.StringNull()
	data.MetricTypeVariable = types.StringNull()
	if t := res.Get(path + "advanced.defaultOriginate.metricType.optionType"); t.Exists() {
		va := res.Get(path + "advanced.defaultOriginate.metricType.value")
		if t.String() == "variable" {
			data.MetricTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MetricType = types.StringValue(va.String())
		}
	}
	data.SpfCalculationDeplay = types.Int64Null()
	data.SpfCalculationDeplayVariable = types.StringNull()
	if t := res.Get(path + "advanced.spfTimers.delay.optionType"); t.Exists() {
		va := res.Get(path + "advanced.spfTimers.delay.value")
		if t.String() == "variable" {
			data.SpfCalculationDeplayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SpfCalculationDeplay = types.Int64Value(va.Int())
		}
	}
	data.InitialHoldTime = types.Int64Null()
	data.InitialHoldTimeVariable = types.StringNull()
	if t := res.Get(path + "advanced.spfTimers.initialHold.optionType"); t.Exists() {
		va := res.Get(path + "advanced.spfTimers.initialHold.value")
		if t.String() == "variable" {
			data.InitialHoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InitialHoldTime = types.Int64Value(va.Int())
		}
	}
	data.MaximumHoldTime = types.Int64Null()
	data.MaximumHoldTimeVariable = types.StringNull()
	if t := res.Get(path + "advanced.spfTimers.maxHold.optionType"); t.Exists() {
		va := res.Get(path + "advanced.spfTimers.maxHold.value")
		if t.String() == "variable" {
			data.MaximumHoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MaximumHoldTime = types.Int64Value(va.Int())
		}
	}
	data.RoutePolicyId = types.StringNull()

	if t := res.Get(path + "advanced.policyName.refId.optionType"); t.Exists() {
		va := res.Get(path + "advanced.policyName.refId.value")
		if t.String() == "global" {
			data.RoutePolicyId = types.StringValue(va.String())
		}
	}
	data.Filter = types.BoolNull()
	data.FilterVariable = types.StringNull()
	if t := res.Get(path + "advanced.filter.optionType"); t.Exists() {
		va := res.Get(path + "advanced.filter.value")
		if t.String() == "variable" {
			data.FilterVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Filter = types.BoolValue(va.Bool())
		}
	}
	for i := range data.Redistributes {
		keys := [...]string{"protocol"}
		keyValues := [...]string{data.Redistributes[i].Protocol.ValueString()}
		keyValuesVariables := [...]string{data.Redistributes[i].ProtocolVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "redistribute").ForEach(
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
		data.Redistributes[i].Protocol = types.StringNull()
		data.Redistributes[i].ProtocolVariable = types.StringNull()
		if t := r.Get("protocol.optionType"); t.Exists() {
			va := r.Get("protocol.value")
			if t.String() == "variable" {
				data.Redistributes[i].ProtocolVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Redistributes[i].Protocol = types.StringValue(va.String())
			}
		}
		data.Redistributes[i].RoutePolicyId = types.StringNull()

		if t := r.Get("routePolicy.refId.optionType"); t.Exists() {
			va := r.Get("routePolicy.refId.value")
			if t.String() == "global" {
				data.Redistributes[i].RoutePolicyId = types.StringValue(va.String())
			}
		}
	}
	data.RouterLsaAction = types.StringNull()

	if t := res.Get(path + "maxMetricRouterLsa.action.optionType"); t.Exists() {
		va := res.Get(path + "maxMetricRouterLsa.action.value")
		if t.String() == "global" {
			data.RouterLsaAction = types.StringValue(va.String())
		}
	}
	data.RouterLsaOnStartuPTime = types.Int64Null()
	data.RouterLsaOnStartuPTimeVariable = types.StringNull()
	if t := res.Get(path + "maxMetricRouterLsa.onStartUpTime.optionType"); t.Exists() {
		va := res.Get(path + "maxMetricRouterLsa.onStartUpTime.value")
		if t.String() == "variable" {
			data.RouterLsaOnStartuPTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RouterLsaOnStartuPTime = types.Int64Value(va.Int())
		}
	}
	for i := range data.Areas {
		keys := [...]string{"areaNum"}
		keyValues := [...]string{strconv.FormatInt(data.Areas[i].AreaNumber.ValueInt64(), 10)}
		keyValuesVariables := [...]string{data.Areas[i].AreaNumberVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "area").ForEach(
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
		data.Areas[i].AreaNumber = types.Int64Null()
		data.Areas[i].AreaNumberVariable = types.StringNull()
		if t := r.Get("areaNum.optionType"); t.Exists() {
			va := r.Get("areaNum.value")
			if t.String() == "variable" {
				data.Areas[i].AreaNumberVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Areas[i].AreaNumber = types.Int64Value(va.Int())
			}
		}
		data.Areas[i].AreaType = types.StringNull()

		if t := r.Get("areaTypeConfig.areaType.optionType"); t.Exists() {
			va := r.Get("areaTypeConfig.areaType.value")
			if t.String() == "global" {
				data.Areas[i].AreaType = types.StringValue(va.String())
			}
		}
		data.Areas[i].NoSummary = types.BoolNull()
		data.Areas[i].NoSummaryVariable = types.StringNull()
		if t := r.Get("areaTypeConfig.noSummary.optionType"); t.Exists() {
			va := r.Get("areaTypeConfig.noSummary.value")
			if t.String() == "variable" {
				data.Areas[i].NoSummaryVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Areas[i].NoSummary = types.BoolValue(va.Bool())
			}
		}
		data.Areas[i].AlwaysTranslate = types.BoolNull()
		data.Areas[i].AlwaysTranslateVariable = types.StringNull()
		if t := r.Get("areaTypeConfig.alwaysTranslate.optionType"); t.Exists() {
			va := r.Get("areaTypeConfig.alwaysTranslate.value")
			if t.String() == "variable" {
				data.Areas[i].AlwaysTranslateVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Areas[i].AlwaysTranslate = types.BoolValue(va.Bool())
			}
		}
		for ci := range data.Areas[i].Interfaces {
			keys := [...]string{"ifName"}
			keyValues := [...]string{data.Areas[i].Interfaces[ci].IfName.ValueString()}
			keyValuesVariables := [...]string{data.Areas[i].Interfaces[ci].IfNameVariable.ValueString()}

			var cr gjson.Result
			r.Get("interfaces").ForEach(
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
			data.Areas[i].Interfaces[ci].IfName = types.StringNull()
			data.Areas[i].Interfaces[ci].IfNameVariable = types.StringNull()
			if t := cr.Get("ifName.optionType"); t.Exists() {
				va := cr.Get("ifName.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].IfNameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].IfName = types.StringValue(va.String())
				}
			}
			data.Areas[i].Interfaces[ci].HelloInterval = types.Int64Null()
			data.Areas[i].Interfaces[ci].HelloIntervalVariable = types.StringNull()
			if t := cr.Get("helloInterval.optionType"); t.Exists() {
				va := cr.Get("helloInterval.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].HelloIntervalVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].HelloInterval = types.Int64Value(va.Int())
				}
			}
			data.Areas[i].Interfaces[ci].DeadInterval = types.Int64Null()
			data.Areas[i].Interfaces[ci].DeadIntervalVariable = types.StringNull()
			if t := cr.Get("deadInterval.optionType"); t.Exists() {
				va := cr.Get("deadInterval.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].DeadIntervalVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].DeadInterval = types.Int64Value(va.Int())
				}
			}
			data.Areas[i].Interfaces[ci].LsaRetransmitInterval = types.Int64Null()
			data.Areas[i].Interfaces[ci].LsaRetransmitIntervalVariable = types.StringNull()
			if t := cr.Get("retransmitInterval.optionType"); t.Exists() {
				va := cr.Get("retransmitInterval.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].LsaRetransmitIntervalVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].LsaRetransmitInterval = types.Int64Value(va.Int())
				}
			}
			data.Areas[i].Interfaces[ci].InterfaceCost = types.Int64Null()
			data.Areas[i].Interfaces[ci].InterfaceCostVariable = types.StringNull()
			if t := cr.Get("cost.optionType"); t.Exists() {
				va := cr.Get("cost.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].InterfaceCostVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].InterfaceCost = types.Int64Value(va.Int())
				}
			}
			data.Areas[i].Interfaces[ci].OspfNetworkType = types.StringNull()
			data.Areas[i].Interfaces[ci].OspfNetworkTypeVariable = types.StringNull()
			if t := cr.Get("networkType.optionType"); t.Exists() {
				va := cr.Get("networkType.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].OspfNetworkTypeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].OspfNetworkType = types.StringValue(va.String())
				}
			}
			data.Areas[i].Interfaces[ci].PassiveInterface = types.BoolNull()
			data.Areas[i].Interfaces[ci].PassiveInterfaceVariable = types.StringNull()
			if t := cr.Get("passiveInterface.optionType"); t.Exists() {
				va := cr.Get("passiveInterface.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].PassiveInterfaceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].PassiveInterface = types.BoolValue(va.Bool())
				}
			}
			data.Areas[i].Interfaces[ci].AuthType = types.StringNull()

			if t := cr.Get("authenticationConfig.authType.optionType"); t.Exists() {
				va := cr.Get("authenticationConfig.authType.value")
				if t.String() == "global" {
					data.Areas[i].Interfaces[ci].AuthType = types.StringValue(va.String())
				}
			}
			data.Areas[i].Interfaces[ci].Spi = types.Int64Null()
			data.Areas[i].Interfaces[ci].SpiVariable = types.StringNull()
			if t := cr.Get("authenticationConfig.spi.optionType"); t.Exists() {
				va := cr.Get("authenticationConfig.spi.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].SpiVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].Spi = types.Int64Value(va.Int())
				}
			}
			data.Areas[i].Interfaces[ci].AuthKey = types.StringNull()
			data.Areas[i].Interfaces[ci].AuthKeyVariable = types.StringNull()
			if t := cr.Get("authenticationConfig.authKey.optionType"); t.Exists() {
				va := cr.Get("authenticationConfig.authKey.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].AuthKeyVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].AuthKey = types.StringValue(va.String())
				}
			}
		}
		for ci := range data.Areas[i].Ranges {
			keys := [...]string{"network"}
			keyValues := [...]string{data.Areas[i].Ranges[ci].Network.ValueString()}
			keyValuesVariables := [...]string{data.Areas[i].Ranges[ci].NetworkVariable.ValueString()}

			var cr gjson.Result
			r.Get("ranges").ForEach(
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
			data.Areas[i].Ranges[ci].Network = types.StringNull()
			data.Areas[i].Ranges[ci].NetworkVariable = types.StringNull()
			if t := cr.Get("network.optionType"); t.Exists() {
				va := cr.Get("network.value")
				if t.String() == "variable" {
					data.Areas[i].Ranges[ci].NetworkVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Ranges[ci].Network = types.StringValue(va.String())
				}
			}
			data.Areas[i].Ranges[ci].Cost = types.Int64Null()
			data.Areas[i].Ranges[ci].CostVariable = types.StringNull()
			if t := cr.Get("cost.optionType"); t.Exists() {
				va := cr.Get("cost.value")
				if t.String() == "variable" {
					data.Areas[i].Ranges[ci].CostVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Ranges[ci].Cost = types.Int64Value(va.Int())
				}
			}
			data.Areas[i].Ranges[ci].NoAdvertise = types.BoolNull()
			data.Areas[i].Ranges[ci].NoAdvertiseVariable = types.StringNull()
			if t := cr.Get("noAdvertise.optionType"); t.Exists() {
				va := cr.Get("noAdvertise.value")
				if t.String() == "variable" {
					data.Areas[i].Ranges[ci].NoAdvertiseVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Ranges[ci].NoAdvertise = types.BoolValue(va.Bool())
				}
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *TransportRoutingOSPFv3IPv6) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.RouterId.IsNull() {
		return false
	}
	if !data.RouterIdVariable.IsNull() {
		return false
	}
	if !data.Distance.IsNull() {
		return false
	}
	if !data.DistanceVariable.IsNull() {
		return false
	}
	if !data.DistanceForExternalRoutes.IsNull() {
		return false
	}
	if !data.DistanceForExternalRoutesVariable.IsNull() {
		return false
	}
	if !data.DistanceForInterAreaRoutes.IsNull() {
		return false
	}
	if !data.DistanceForInterAreaRoutesVariable.IsNull() {
		return false
	}
	if !data.DistanceForIntraAreaRoutes.IsNull() {
		return false
	}
	if !data.DistanceForIntraAreaRoutesVariable.IsNull() {
		return false
	}
	if !data.ReferenceBandwidth.IsNull() {
		return false
	}
	if !data.ReferenceBandwidthVariable.IsNull() {
		return false
	}
	if !data.Rfc1583Compatible.IsNull() {
		return false
	}
	if !data.Rfc1583CompatibleVariable.IsNull() {
		return false
	}
	if !data.Originate.IsNull() {
		return false
	}
	if !data.Always.IsNull() {
		return false
	}
	if !data.AlwaysVariable.IsNull() {
		return false
	}
	if !data.Metric.IsNull() {
		return false
	}
	if !data.MetricVariable.IsNull() {
		return false
	}
	if !data.MetricType.IsNull() {
		return false
	}
	if !data.MetricTypeVariable.IsNull() {
		return false
	}
	if !data.SpfCalculationDeplay.IsNull() {
		return false
	}
	if !data.SpfCalculationDeplayVariable.IsNull() {
		return false
	}
	if !data.InitialHoldTime.IsNull() {
		return false
	}
	if !data.InitialHoldTimeVariable.IsNull() {
		return false
	}
	if !data.MaximumHoldTime.IsNull() {
		return false
	}
	if !data.MaximumHoldTimeVariable.IsNull() {
		return false
	}
	if !data.RoutePolicyId.IsNull() {
		return false
	}
	if !data.Filter.IsNull() {
		return false
	}
	if !data.FilterVariable.IsNull() {
		return false
	}
	if len(data.Redistributes) > 0 {
		return false
	}
	if !data.RouterLsaAction.IsNull() {
		return false
	}
	if !data.RouterLsaOnStartuPTime.IsNull() {
		return false
	}
	if !data.RouterLsaOnStartuPTimeVariable.IsNull() {
		return false
	}
	if len(data.Areas) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
