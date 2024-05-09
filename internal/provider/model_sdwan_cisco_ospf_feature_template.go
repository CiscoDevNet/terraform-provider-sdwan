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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type CiscoOSPF struct {
	Id                                            types.String                  `tfsdk:"id"`
	Version                                       types.Int64                   `tfsdk:"version"`
	TemplateType                                  types.String                  `tfsdk:"template_type"`
	Name                                          types.String                  `tfsdk:"name"`
	Description                                   types.String                  `tfsdk:"description"`
	DeviceTypes                                   types.Set                     `tfsdk:"device_types"`
	RouterId                                      types.String                  `tfsdk:"router_id"`
	RouterIdVariable                              types.String                  `tfsdk:"router_id_variable"`
	AutoCostReferenceBandwidth                    types.Int64                   `tfsdk:"auto_cost_reference_bandwidth"`
	AutoCostReferenceBandwidthVariable            types.String                  `tfsdk:"auto_cost_reference_bandwidth_variable"`
	CompatibleRfc1583                             types.Bool                    `tfsdk:"compatible_rfc1583"`
	CompatibleRfc1583Variable                     types.String                  `tfsdk:"compatible_rfc1583_variable"`
	DefaultInformationOriginate                   types.Bool                    `tfsdk:"default_information_originate"`
	DefaultInformationOriginateAlways             types.Bool                    `tfsdk:"default_information_originate_always"`
	DefaultInformationOriginateAlwaysVariable     types.String                  `tfsdk:"default_information_originate_always_variable"`
	DefaultInformationOriginateMetric             types.Int64                   `tfsdk:"default_information_originate_metric"`
	DefaultInformationOriginateMetricVariable     types.String                  `tfsdk:"default_information_originate_metric_variable"`
	DefaultInformationOriginateMetricType         types.String                  `tfsdk:"default_information_originate_metric_type"`
	DefaultInformationOriginateMetricTypeVariable types.String                  `tfsdk:"default_information_originate_metric_type_variable"`
	DistanceExternal                              types.Int64                   `tfsdk:"distance_external"`
	DistanceExternalVariable                      types.String                  `tfsdk:"distance_external_variable"`
	DistanceInterArea                             types.Int64                   `tfsdk:"distance_inter_area"`
	DistanceInterAreaVariable                     types.String                  `tfsdk:"distance_inter_area_variable"`
	DistanceIntraArea                             types.Int64                   `tfsdk:"distance_intra_area"`
	DistanceIntraAreaVariable                     types.String                  `tfsdk:"distance_intra_area_variable"`
	TimersSpfDelay                                types.Int64                   `tfsdk:"timers_spf_delay"`
	TimersSpfDelayVariable                        types.String                  `tfsdk:"timers_spf_delay_variable"`
	TimersSpfInitialHold                          types.Int64                   `tfsdk:"timers_spf_initial_hold"`
	TimersSpfInitialHoldVariable                  types.String                  `tfsdk:"timers_spf_initial_hold_variable"`
	TimersSpfMaxHold                              types.Int64                   `tfsdk:"timers_spf_max_hold"`
	TimersSpfMaxHoldVariable                      types.String                  `tfsdk:"timers_spf_max_hold_variable"`
	Redistribute                                  []CiscoOSPFRedistribute       `tfsdk:"redistribute"`
	MaxMetricRouterLsa                            []CiscoOSPFMaxMetricRouterLsa `tfsdk:"max_metric_router_lsa"`
	RoutePolicies                                 []CiscoOSPFRoutePolicies      `tfsdk:"route_policies"`
	Areas                                         []CiscoOSPFAreas              `tfsdk:"areas"`
}

type CiscoOSPFRedistribute struct {
	Optional            types.Bool   `tfsdk:"optional"`
	Protocol            types.String `tfsdk:"protocol"`
	ProtocolVariable    types.String `tfsdk:"protocol_variable"`
	RoutePolicy         types.String `tfsdk:"route_policy"`
	RoutePolicyVariable types.String `tfsdk:"route_policy_variable"`
	NatDia              types.Bool   `tfsdk:"nat_dia"`
	NatDiaVariable      types.String `tfsdk:"nat_dia_variable"`
}

type CiscoOSPFMaxMetricRouterLsa struct {
	Optional     types.Bool   `tfsdk:"optional"`
	AdType       types.String `tfsdk:"ad_type"`
	Time         types.Int64  `tfsdk:"time"`
	TimeVariable types.String `tfsdk:"time_variable"`
}

type CiscoOSPFRoutePolicies struct {
	Optional           types.Bool   `tfsdk:"optional"`
	Direction          types.String `tfsdk:"direction"`
	DirectionVariable  types.String `tfsdk:"direction_variable"`
	PolicyName         types.String `tfsdk:"policy_name"`
	PolicyNameVariable types.String `tfsdk:"policy_name_variable"`
}

type CiscoOSPFAreas struct {
	Optional              types.Bool                 `tfsdk:"optional"`
	AreaNumber            types.Int64                `tfsdk:"area_number"`
	AreaNumberVariable    types.String               `tfsdk:"area_number_variable"`
	Stub                  types.Bool                 `tfsdk:"stub"`
	StubNoSummary         types.Bool                 `tfsdk:"stub_no_summary"`
	StubNoSummaryVariable types.String               `tfsdk:"stub_no_summary_variable"`
	Nssa                  types.Bool                 `tfsdk:"nssa"`
	NssaNoSummary         types.Bool                 `tfsdk:"nssa_no_summary"`
	NssaNoSummaryVariable types.String               `tfsdk:"nssa_no_summary_variable"`
	Interfaces            []CiscoOSPFAreasInterfaces `tfsdk:"interfaces"`
	Ranges                []CiscoOSPFAreasRanges     `tfsdk:"ranges"`
}

type CiscoOSPFAreasInterfaces struct {
	Optional                                 types.Bool   `tfsdk:"optional"`
	Name                                     types.String `tfsdk:"name"`
	NameVariable                             types.String `tfsdk:"name_variable"`
	HelloInterval                            types.Int64  `tfsdk:"hello_interval"`
	HelloIntervalVariable                    types.String `tfsdk:"hello_interval_variable"`
	DeadInterval                             types.Int64  `tfsdk:"dead_interval"`
	DeadIntervalVariable                     types.String `tfsdk:"dead_interval_variable"`
	RetransmitInterval                       types.Int64  `tfsdk:"retransmit_interval"`
	RetransmitIntervalVariable               types.String `tfsdk:"retransmit_interval_variable"`
	Cost                                     types.Int64  `tfsdk:"cost"`
	CostVariable                             types.String `tfsdk:"cost_variable"`
	Priority                                 types.Int64  `tfsdk:"priority"`
	PriorityVariable                         types.String `tfsdk:"priority_variable"`
	Network                                  types.String `tfsdk:"network"`
	NetworkVariable                          types.String `tfsdk:"network_variable"`
	PassiveInterface                         types.Bool   `tfsdk:"passive_interface"`
	PassiveInterfaceVariable                 types.String `tfsdk:"passive_interface_variable"`
	AuthenticationType                       types.String `tfsdk:"authentication_type"`
	AuthenticationTypeVariable               types.String `tfsdk:"authentication_type_variable"`
	AuthenticationMessageDigestKeyId         types.Int64  `tfsdk:"authentication_message_digest_key_id"`
	AuthenticationMessageDigestKeyIdVariable types.String `tfsdk:"authentication_message_digest_key_id_variable"`
	AuthenticationMessageDigestKey           types.String `tfsdk:"authentication_message_digest_key"`
	AuthenticationMessageDigestKeyVariable   types.String `tfsdk:"authentication_message_digest_key_variable"`
}
type CiscoOSPFAreasRanges struct {
	Optional            types.Bool   `tfsdk:"optional"`
	Address             types.String `tfsdk:"address"`
	AddressVariable     types.String `tfsdk:"address_variable"`
	Cost                types.Int64  `tfsdk:"cost"`
	CostVariable        types.String `tfsdk:"cost_variable"`
	NoAdvertise         types.Bool   `tfsdk:"no_advertise"`
	NoAdvertiseVariable types.String `tfsdk:"no_advertise_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoOSPF) getModel() string {
	return "cisco_ospf"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CiscoOSPF) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_ospf")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.RouterIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.router-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.router-id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospf.router-id."+"vipVariableName", data.RouterIdVariable.ValueString())
	} else if data.RouterId.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.router-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.router-id."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospf.router-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.router-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospf.router-id."+"vipValue", data.RouterId.ValueString())
	}

	if !data.AutoCostReferenceBandwidthVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.auto-cost.reference-bandwidth."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.auto-cost.reference-bandwidth."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospf.auto-cost.reference-bandwidth."+"vipVariableName", data.AutoCostReferenceBandwidthVariable.ValueString())
	} else if data.AutoCostReferenceBandwidth.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.auto-cost.reference-bandwidth."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.auto-cost.reference-bandwidth."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospf.auto-cost.reference-bandwidth."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.auto-cost.reference-bandwidth."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospf.auto-cost.reference-bandwidth."+"vipValue", data.AutoCostReferenceBandwidth.ValueInt64())
	}

	if !data.CompatibleRfc1583Variable.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.compatible.rfc1583."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.compatible.rfc1583."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospf.compatible.rfc1583."+"vipVariableName", data.CompatibleRfc1583Variable.ValueString())
	} else if data.CompatibleRfc1583.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.compatible.rfc1583."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.compatible.rfc1583."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospf.compatible.rfc1583."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.compatible.rfc1583."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospf.compatible.rfc1583."+"vipValue", strconv.FormatBool(data.CompatibleRfc1583.ValueBool()))
	}
	if !data.DefaultInformationOriginate.IsNull() {
		if data.DefaultInformationOriginate.ValueBool() {
			if !gjson.Get(body, path+"ospf.default-information.originate").Exists() {
				body, _ = sjson.Set(body, path+"ospf.default-information.originate", map[string]interface{}{})
			}
		} else {
			body, _ = sjson.Set(body, path+"ospf.default-information.originate."+"vipObjectType", "node-only")
			body, _ = sjson.Set(body, path+"ospf.default-information.originate."+"vipType", "ignore")
		}
	}

	if !data.DefaultInformationOriginateAlwaysVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.always."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.always."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.always."+"vipVariableName", data.DefaultInformationOriginateAlwaysVariable.ValueString())
	} else if data.DefaultInformationOriginateAlways.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.always."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.always."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.always."+"vipValue", strconv.FormatBool(data.DefaultInformationOriginateAlways.ValueBool()))
	}

	if !data.DefaultInformationOriginateMetricVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.metric."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.metric."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.metric."+"vipVariableName", data.DefaultInformationOriginateMetricVariable.ValueString())
	} else if data.DefaultInformationOriginateMetric.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.metric."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.metric."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.metric."+"vipValue", data.DefaultInformationOriginateMetric.ValueInt64())
	}

	if !data.DefaultInformationOriginateMetricTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.metric-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.metric-type."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.metric-type."+"vipVariableName", data.DefaultInformationOriginateMetricTypeVariable.ValueString())
	} else if data.DefaultInformationOriginateMetricType.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.metric-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.metric-type."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospf.default-information.originate.metric-type."+"vipValue", data.DefaultInformationOriginateMetricType.ValueString())
	}

	if !data.DistanceExternalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.distance.external."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.distance.external."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospf.distance.external."+"vipVariableName", data.DistanceExternalVariable.ValueString())
	} else if data.DistanceExternal.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.distance.external."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.distance.external."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospf.distance.external."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.distance.external."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospf.distance.external."+"vipValue", data.DistanceExternal.ValueInt64())
	}

	if !data.DistanceInterAreaVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.distance.inter-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.distance.inter-area."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospf.distance.inter-area."+"vipVariableName", data.DistanceInterAreaVariable.ValueString())
	} else if data.DistanceInterArea.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.distance.inter-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.distance.inter-area."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospf.distance.inter-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.distance.inter-area."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospf.distance.inter-area."+"vipValue", data.DistanceInterArea.ValueInt64())
	}

	if !data.DistanceIntraAreaVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.distance.intra-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.distance.intra-area."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospf.distance.intra-area."+"vipVariableName", data.DistanceIntraAreaVariable.ValueString())
	} else if data.DistanceIntraArea.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.distance.intra-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.distance.intra-area."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospf.distance.intra-area."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.distance.intra-area."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospf.distance.intra-area."+"vipValue", data.DistanceIntraArea.ValueInt64())
	}

	if !data.TimersSpfDelayVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.timers.spf.delay."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.timers.spf.delay."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospf.timers.spf.delay."+"vipVariableName", data.TimersSpfDelayVariable.ValueString())
	} else if data.TimersSpfDelay.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.timers.spf.delay."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.timers.spf.delay."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospf.timers.spf.delay."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.timers.spf.delay."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospf.timers.spf.delay."+"vipValue", data.TimersSpfDelay.ValueInt64())
	}

	if !data.TimersSpfInitialHoldVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.timers.spf.initial-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.timers.spf.initial-hold."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospf.timers.spf.initial-hold."+"vipVariableName", data.TimersSpfInitialHoldVariable.ValueString())
	} else if data.TimersSpfInitialHold.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.timers.spf.initial-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.timers.spf.initial-hold."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospf.timers.spf.initial-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.timers.spf.initial-hold."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospf.timers.spf.initial-hold."+"vipValue", data.TimersSpfInitialHold.ValueInt64())
	}

	if !data.TimersSpfMaxHoldVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.timers.spf.max-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.timers.spf.max-hold."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ospf.timers.spf.max-hold."+"vipVariableName", data.TimersSpfMaxHoldVariable.ValueString())
	} else if data.TimersSpfMaxHold.IsNull() {
		body, _ = sjson.Set(body, path+"ospf.timers.spf.max-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.timers.spf.max-hold."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ospf.timers.spf.max-hold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ospf.timers.spf.max-hold."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospf.timers.spf.max-hold."+"vipValue", data.TimersSpfMaxHold.ValueInt64())
	}
	if len(data.Redistribute) > 0 {
		body, _ = sjson.Set(body, path+"ospf.redistribute."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospf.redistribute."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospf.redistribute."+"vipPrimaryKey", []string{"protocol"})
		body, _ = sjson.Set(body, path+"ospf.redistribute."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ospf.redistribute."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospf.redistribute."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ospf.redistribute."+"vipPrimaryKey", []string{"protocol"})
		body, _ = sjson.Set(body, path+"ospf.redistribute."+"vipValue", []interface{}{})
	}
	for _, item := range data.Redistribute {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "protocol")

		if !item.ProtocolVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipVariableName", item.ProtocolVariable.ValueString())
		} else if item.Protocol.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipValue", item.Protocol.ValueString())
		}
		itemAttributes = append(itemAttributes, "route-policy")

		if !item.RoutePolicyVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipVariableName", item.RoutePolicyVariable.ValueString())
		} else if item.RoutePolicy.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "route-policy."+"vipValue", item.RoutePolicy.ValueString())
		}
		itemAttributes = append(itemAttributes, "dia")

		if !item.NatDiaVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dia."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dia."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dia."+"vipVariableName", item.NatDiaVariable.ValueString())
		} else if item.NatDia.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dia."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dia."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "dia."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dia."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dia."+"vipValue", strconv.FormatBool(item.NatDia.ValueBool()))
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ospf.redistribute."+"vipValue.-1", itemBody)
	}
	if len(data.MaxMetricRouterLsa) > 0 {
		body, _ = sjson.Set(body, path+"ospf.max-metric.router-lsa."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospf.max-metric.router-lsa."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospf.max-metric.router-lsa."+"vipPrimaryKey", []string{"ad-type"})
		body, _ = sjson.Set(body, path+"ospf.max-metric.router-lsa."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ospf.max-metric.router-lsa."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospf.max-metric.router-lsa."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ospf.max-metric.router-lsa."+"vipPrimaryKey", []string{"ad-type"})
		body, _ = sjson.Set(body, path+"ospf.max-metric.router-lsa."+"vipValue", []interface{}{})
	}
	for _, item := range data.MaxMetricRouterLsa {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "ad-type")
		if item.AdType.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ad-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ad-type."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ad-type."+"vipValue", item.AdType.ValueString())
		}
		itemAttributes = append(itemAttributes, "time")

		if !item.TimeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "time."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "time."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "time."+"vipVariableName", item.TimeVariable.ValueString())
		} else if item.Time.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "time."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "time."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "time."+"vipValue", item.Time.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ospf.max-metric.router-lsa."+"vipValue.-1", itemBody)
	}
	if len(data.RoutePolicies) > 0 {
		body, _ = sjson.Set(body, path+"ospf.route-policy."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospf.route-policy."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospf.route-policy."+"vipPrimaryKey", []string{"direction"})
		body, _ = sjson.Set(body, path+"ospf.route-policy."+"vipValue", []interface{}{})
	} else {
	}
	for _, item := range data.RoutePolicies {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "direction")

		if !item.DirectionVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipVariableName", item.DirectionVariable.ValueString())
		} else if item.Direction.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "direction."+"vipValue", item.Direction.ValueString())
		}
		itemAttributes = append(itemAttributes, "pol-name")

		if !item.PolicyNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "pol-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "pol-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "pol-name."+"vipVariableName", item.PolicyNameVariable.ValueString())
		} else if item.PolicyName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "pol-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "pol-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "pol-name."+"vipValue", item.PolicyName.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ospf.route-policy."+"vipValue.-1", itemBody)
	}
	if len(data.Areas) > 0 {
		body, _ = sjson.Set(body, path+"ospf.area."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospf.area."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ospf.area."+"vipPrimaryKey", []string{"a-num"})
		body, _ = sjson.Set(body, path+"ospf.area."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ospf.area."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ospf.area."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ospf.area."+"vipPrimaryKey", []string{"a-num"})
		body, _ = sjson.Set(body, path+"ospf.area."+"vipValue", []interface{}{})
	}
	for _, item := range data.Areas {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "a-num")

		if !item.AreaNumberVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipVariableName", item.AreaNumberVariable.ValueString())
		} else if item.AreaNumber.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "a-num."+"vipValue", item.AreaNumber.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "stub")
		if !item.Stub.IsNull() {
			if item.Stub.ValueBool() {
				if !gjson.Get(itemBody, "stub").Exists() {
					itemBody, _ = sjson.Set(itemBody, "stub", map[string]interface{}{})
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "stub."+"vipObjectType", "")
				itemBody, _ = sjson.Set(itemBody, "stub."+"vipType", "ignore")
			}
		}
		itemAttributes = append(itemAttributes, "stub")

		if !item.StubNoSummaryVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipVariableName", item.StubNoSummaryVariable.ValueString())
		} else if item.StubNoSummary.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "stub.no-summary."+"vipValue", strconv.FormatBool(item.StubNoSummary.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "nssa")
		if !item.Nssa.IsNull() {
			if item.Nssa.ValueBool() {
				if !gjson.Get(itemBody, "nssa").Exists() {
					itemBody, _ = sjson.Set(itemBody, "nssa", map[string]interface{}{})
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "nssa."+"vipObjectType", "")
				itemBody, _ = sjson.Set(itemBody, "nssa."+"vipType", "ignore")
			}
		}
		itemAttributes = append(itemAttributes, "nssa")

		if !item.NssaNoSummaryVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipVariableName", item.NssaNoSummaryVariable.ValueString())
		} else if item.NssaNoSummary.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "nssa.no-summary."+"vipValue", strconv.FormatBool(item.NssaNoSummary.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "interface")
		if len(item.Interfaces) > 0 {
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipPrimaryKey", []string{"name"})
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipPrimaryKey", []string{"name"})
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Interfaces {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "name")

			if !childItem.NameVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipVariableName", childItem.NameVariable.ValueString())
			} else if childItem.Name.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipValue", childItem.Name.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "hello-interval")

			if !childItem.HelloIntervalVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipVariableName", childItem.HelloIntervalVariable.ValueString())
			} else if childItem.HelloInterval.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "hello-interval."+"vipValue", childItem.HelloInterval.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "dead-interval")

			if !childItem.DeadIntervalVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipVariableName", childItem.DeadIntervalVariable.ValueString())
			} else if childItem.DeadInterval.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "dead-interval."+"vipValue", childItem.DeadInterval.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "retransmit-interval")

			if !childItem.RetransmitIntervalVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipVariableName", childItem.RetransmitIntervalVariable.ValueString())
			} else if childItem.RetransmitInterval.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit-interval."+"vipValue", childItem.RetransmitInterval.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "cost")

			if !childItem.CostVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipVariableName", childItem.CostVariable.ValueString())
			} else if childItem.Cost.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipValue", childItem.Cost.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "priority")

			if !childItem.PriorityVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "priority."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "priority."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "priority."+"vipVariableName", childItem.PriorityVariable.ValueString())
			} else if childItem.Priority.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "priority."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "priority."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "priority."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "priority."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "priority."+"vipValue", childItem.Priority.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "network")

			if !childItem.NetworkVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipVariableName", childItem.NetworkVariable.ValueString())
			} else if childItem.Network.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "network."+"vipValue", childItem.Network.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "passive-interface")

			if !childItem.PassiveInterfaceVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipVariableName", childItem.PassiveInterfaceVariable.ValueString())
			} else if childItem.PassiveInterface.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "passive-interface."+"vipValue", strconv.FormatBool(childItem.PassiveInterface.ValueBool()))
			}
			itemChildAttributes = append(itemChildAttributes, "authentication")

			if !childItem.AuthenticationTypeVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipVariableName", childItem.AuthenticationTypeVariable.ValueString())
			} else if childItem.AuthenticationType.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type."+"vipValue", childItem.AuthenticationType.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "authentication")

			if !childItem.AuthenticationMessageDigestKeyIdVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.message-digest.message-digest-key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.message-digest.message-digest-key."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.message-digest.message-digest-key."+"vipVariableName", childItem.AuthenticationMessageDigestKeyIdVariable.ValueString())
			} else if childItem.AuthenticationMessageDigestKeyId.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.message-digest.message-digest-key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.message-digest.message-digest-key."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.message-digest.message-digest-key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.message-digest.message-digest-key."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.message-digest.message-digest-key."+"vipValue", childItem.AuthenticationMessageDigestKeyId.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "authentication")

			if !childItem.AuthenticationMessageDigestKeyVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.message-digest.md5."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.message-digest.md5."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.message-digest.md5."+"vipVariableName", childItem.AuthenticationMessageDigestKeyVariable.ValueString())
			} else if childItem.AuthenticationMessageDigestKey.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.message-digest.md5."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.message-digest.md5."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.message-digest.md5."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.message-digest.md5."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "authentication.message-digest.md5."+"vipValue", childItem.AuthenticationMessageDigestKey.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "interface."+"vipValue.-1", itemChildBody)
		}
		itemAttributes = append(itemAttributes, "range")
		if len(item.Ranges) > 0 {
			itemBody, _ = sjson.Set(itemBody, "range."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "range."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "range."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "range."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "range."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "range."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "range."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "range."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Ranges {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "address")

			if !childItem.AddressVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipVariableName", childItem.AddressVariable.ValueString())
			} else if childItem.Address.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipValue", childItem.Address.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "cost")

			if !childItem.CostVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipVariableName", childItem.CostVariable.ValueString())
			} else if childItem.Cost.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "cost."+"vipValue", childItem.Cost.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "no-advertise")

			if !childItem.NoAdvertiseVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipVariableName", childItem.NoAdvertiseVariable.ValueString())
			} else if childItem.NoAdvertise.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "no-advertise."+"vipValue", strconv.FormatBool(childItem.NoAdvertise.ValueBool()))
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "range."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ospf.area."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoOSPF) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceType"); value.Exists() {
		data.DeviceTypes = helpers.GetStringSet(value.Array())
	} else {
		data.DeviceTypes = types.SetNull(types.StringType)
	}
	if value := res.Get("templateDescription"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("templateName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("templateType"); value.Exists() {
		data.TemplateType = types.StringValue(value.String())
	} else {
		data.TemplateType = types.StringNull()
	}

	path := "templateDefinition."
	if value := res.Get(path + "ospf.router-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.RouterId = types.StringNull()

			v := res.Get(path + "ospf.router-id.vipVariableName")
			data.RouterIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.RouterId = types.StringNull()
			data.RouterIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospf.router-id.vipValue")
			data.RouterId = types.StringValue(v.String())
			data.RouterIdVariable = types.StringNull()
		}
	} else {
		data.RouterId = types.StringNull()
		data.RouterIdVariable = types.StringNull()
	}
	if value := res.Get(path + "ospf.auto-cost.reference-bandwidth.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AutoCostReferenceBandwidth = types.Int64Null()

			v := res.Get(path + "ospf.auto-cost.reference-bandwidth.vipVariableName")
			data.AutoCostReferenceBandwidthVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AutoCostReferenceBandwidth = types.Int64Null()
			data.AutoCostReferenceBandwidthVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospf.auto-cost.reference-bandwidth.vipValue")
			data.AutoCostReferenceBandwidth = types.Int64Value(v.Int())
			data.AutoCostReferenceBandwidthVariable = types.StringNull()
		}
	} else {
		data.AutoCostReferenceBandwidth = types.Int64Null()
		data.AutoCostReferenceBandwidthVariable = types.StringNull()
	}
	if value := res.Get(path + "ospf.compatible.rfc1583.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.CompatibleRfc1583 = types.BoolNull()

			v := res.Get(path + "ospf.compatible.rfc1583.vipVariableName")
			data.CompatibleRfc1583Variable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.CompatibleRfc1583 = types.BoolNull()
			data.CompatibleRfc1583Variable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospf.compatible.rfc1583.vipValue")
			data.CompatibleRfc1583 = types.BoolValue(v.Bool())
			data.CompatibleRfc1583Variable = types.StringNull()
		}
	} else {
		data.CompatibleRfc1583 = types.BoolNull()
		data.CompatibleRfc1583Variable = types.StringNull()
	}
	if value := res.Get(path + "ospf.default-information.originate.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DefaultInformationOriginate = types.BoolNull()

		} else if value.String() == "ignore" {
			data.DefaultInformationOriginate = types.BoolValue(false)

		} else if value.String() == "constant" {
			v := res.Get(path + "ospf.default-information.originate.vipValue")
			data.DefaultInformationOriginate = types.BoolValue(v.Bool())

		}
	} else if value := res.Get(path + "ospf.default-information.originate"); value.Exists() {
		data.DefaultInformationOriginate = types.BoolValue(true)

	} else {
		data.DefaultInformationOriginate = types.BoolNull()

	}
	if value := res.Get(path + "ospf.default-information.originate.always.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DefaultInformationOriginateAlways = types.BoolNull()

			v := res.Get(path + "ospf.default-information.originate.always.vipVariableName")
			data.DefaultInformationOriginateAlwaysVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DefaultInformationOriginateAlways = types.BoolNull()
			data.DefaultInformationOriginateAlwaysVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospf.default-information.originate.always.vipValue")
			data.DefaultInformationOriginateAlways = types.BoolValue(v.Bool())
			data.DefaultInformationOriginateAlwaysVariable = types.StringNull()
		}
	} else {
		data.DefaultInformationOriginateAlways = types.BoolNull()
		data.DefaultInformationOriginateAlwaysVariable = types.StringNull()
	}
	if value := res.Get(path + "ospf.default-information.originate.metric.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DefaultInformationOriginateMetric = types.Int64Null()

			v := res.Get(path + "ospf.default-information.originate.metric.vipVariableName")
			data.DefaultInformationOriginateMetricVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DefaultInformationOriginateMetric = types.Int64Null()
			data.DefaultInformationOriginateMetricVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospf.default-information.originate.metric.vipValue")
			data.DefaultInformationOriginateMetric = types.Int64Value(v.Int())
			data.DefaultInformationOriginateMetricVariable = types.StringNull()
		}
	} else {
		data.DefaultInformationOriginateMetric = types.Int64Null()
		data.DefaultInformationOriginateMetricVariable = types.StringNull()
	}
	if value := res.Get(path + "ospf.default-information.originate.metric-type.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DefaultInformationOriginateMetricType = types.StringNull()

			v := res.Get(path + "ospf.default-information.originate.metric-type.vipVariableName")
			data.DefaultInformationOriginateMetricTypeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DefaultInformationOriginateMetricType = types.StringNull()
			data.DefaultInformationOriginateMetricTypeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospf.default-information.originate.metric-type.vipValue")
			data.DefaultInformationOriginateMetricType = types.StringValue(v.String())
			data.DefaultInformationOriginateMetricTypeVariable = types.StringNull()
		}
	} else {
		data.DefaultInformationOriginateMetricType = types.StringNull()
		data.DefaultInformationOriginateMetricTypeVariable = types.StringNull()
	}
	if value := res.Get(path + "ospf.distance.external.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DistanceExternal = types.Int64Null()

			v := res.Get(path + "ospf.distance.external.vipVariableName")
			data.DistanceExternalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DistanceExternal = types.Int64Null()
			data.DistanceExternalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospf.distance.external.vipValue")
			data.DistanceExternal = types.Int64Value(v.Int())
			data.DistanceExternalVariable = types.StringNull()
		}
	} else {
		data.DistanceExternal = types.Int64Null()
		data.DistanceExternalVariable = types.StringNull()
	}
	if value := res.Get(path + "ospf.distance.inter-area.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DistanceInterArea = types.Int64Null()

			v := res.Get(path + "ospf.distance.inter-area.vipVariableName")
			data.DistanceInterAreaVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DistanceInterArea = types.Int64Null()
			data.DistanceInterAreaVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospf.distance.inter-area.vipValue")
			data.DistanceInterArea = types.Int64Value(v.Int())
			data.DistanceInterAreaVariable = types.StringNull()
		}
	} else {
		data.DistanceInterArea = types.Int64Null()
		data.DistanceInterAreaVariable = types.StringNull()
	}
	if value := res.Get(path + "ospf.distance.intra-area.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DistanceIntraArea = types.Int64Null()

			v := res.Get(path + "ospf.distance.intra-area.vipVariableName")
			data.DistanceIntraAreaVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DistanceIntraArea = types.Int64Null()
			data.DistanceIntraAreaVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospf.distance.intra-area.vipValue")
			data.DistanceIntraArea = types.Int64Value(v.Int())
			data.DistanceIntraAreaVariable = types.StringNull()
		}
	} else {
		data.DistanceIntraArea = types.Int64Null()
		data.DistanceIntraAreaVariable = types.StringNull()
	}
	if value := res.Get(path + "ospf.timers.spf.delay.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TimersSpfDelay = types.Int64Null()

			v := res.Get(path + "ospf.timers.spf.delay.vipVariableName")
			data.TimersSpfDelayVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TimersSpfDelay = types.Int64Null()
			data.TimersSpfDelayVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospf.timers.spf.delay.vipValue")
			data.TimersSpfDelay = types.Int64Value(v.Int())
			data.TimersSpfDelayVariable = types.StringNull()
		}
	} else {
		data.TimersSpfDelay = types.Int64Null()
		data.TimersSpfDelayVariable = types.StringNull()
	}
	if value := res.Get(path + "ospf.timers.spf.initial-hold.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TimersSpfInitialHold = types.Int64Null()

			v := res.Get(path + "ospf.timers.spf.initial-hold.vipVariableName")
			data.TimersSpfInitialHoldVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TimersSpfInitialHold = types.Int64Null()
			data.TimersSpfInitialHoldVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospf.timers.spf.initial-hold.vipValue")
			data.TimersSpfInitialHold = types.Int64Value(v.Int())
			data.TimersSpfInitialHoldVariable = types.StringNull()
		}
	} else {
		data.TimersSpfInitialHold = types.Int64Null()
		data.TimersSpfInitialHoldVariable = types.StringNull()
	}
	if value := res.Get(path + "ospf.timers.spf.max-hold.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TimersSpfMaxHold = types.Int64Null()

			v := res.Get(path + "ospf.timers.spf.max-hold.vipVariableName")
			data.TimersSpfMaxHoldVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TimersSpfMaxHold = types.Int64Null()
			data.TimersSpfMaxHoldVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ospf.timers.spf.max-hold.vipValue")
			data.TimersSpfMaxHold = types.Int64Value(v.Int())
			data.TimersSpfMaxHoldVariable = types.StringNull()
		}
	} else {
		data.TimersSpfMaxHold = types.Int64Null()
		data.TimersSpfMaxHoldVariable = types.StringNull()
	}
	if value := res.Get(path + "ospf.redistribute.vipValue"); len(value.Array()) > 0 {
		data.Redistribute = make([]CiscoOSPFRedistribute, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoOSPFRedistribute{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("protocol.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Protocol = types.StringNull()

					cv := v.Get("protocol.vipVariableName")
					item.ProtocolVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Protocol = types.StringNull()
					item.ProtocolVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("protocol.vipValue")
					item.Protocol = types.StringValue(cv.String())
					item.ProtocolVariable = types.StringNull()
				}
			} else {
				item.Protocol = types.StringNull()
				item.ProtocolVariable = types.StringNull()
			}
			if cValue := v.Get("route-policy.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.RoutePolicy = types.StringNull()

					cv := v.Get("route-policy.vipVariableName")
					item.RoutePolicyVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.RoutePolicy = types.StringNull()
					item.RoutePolicyVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("route-policy.vipValue")
					item.RoutePolicy = types.StringValue(cv.String())
					item.RoutePolicyVariable = types.StringNull()
				}
			} else {
				item.RoutePolicy = types.StringNull()
				item.RoutePolicyVariable = types.StringNull()
			}
			if cValue := v.Get("dia.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.NatDia = types.BoolNull()

					cv := v.Get("dia.vipVariableName")
					item.NatDiaVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.NatDia = types.BoolNull()
					item.NatDiaVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dia.vipValue")
					item.NatDia = types.BoolValue(cv.Bool())
					item.NatDiaVariable = types.StringNull()
				}
			} else {
				item.NatDia = types.BoolNull()
				item.NatDiaVariable = types.StringNull()
			}
			data.Redistribute = append(data.Redistribute, item)
			return true
		})
	} else {
		if len(data.Redistribute) > 0 {
			data.Redistribute = []CiscoOSPFRedistribute{}
		}
	}
	if value := res.Get(path + "ospf.max-metric.router-lsa.vipValue"); len(value.Array()) > 0 {
		data.MaxMetricRouterLsa = make([]CiscoOSPFMaxMetricRouterLsa, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoOSPFMaxMetricRouterLsa{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("ad-type.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AdType = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.AdType = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("ad-type.vipValue")
					item.AdType = types.StringValue(cv.String())

				}
			} else {
				item.AdType = types.StringNull()

			}
			if cValue := v.Get("time.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Time = types.Int64Null()

					cv := v.Get("time.vipVariableName")
					item.TimeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Time = types.Int64Null()
					item.TimeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("time.vipValue")
					item.Time = types.Int64Value(cv.Int())
					item.TimeVariable = types.StringNull()
				}
			} else {
				item.Time = types.Int64Null()
				item.TimeVariable = types.StringNull()
			}
			data.MaxMetricRouterLsa = append(data.MaxMetricRouterLsa, item)
			return true
		})
	} else {
		if len(data.MaxMetricRouterLsa) > 0 {
			data.MaxMetricRouterLsa = []CiscoOSPFMaxMetricRouterLsa{}
		}
	}
	if value := res.Get(path + "ospf.route-policy.vipValue"); len(value.Array()) > 0 {
		data.RoutePolicies = make([]CiscoOSPFRoutePolicies, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoOSPFRoutePolicies{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("direction.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Direction = types.StringNull()

					cv := v.Get("direction.vipVariableName")
					item.DirectionVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Direction = types.StringNull()
					item.DirectionVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("direction.vipValue")
					item.Direction = types.StringValue(cv.String())
					item.DirectionVariable = types.StringNull()
				}
			} else {
				item.Direction = types.StringNull()
				item.DirectionVariable = types.StringNull()
			}
			if cValue := v.Get("pol-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PolicyName = types.StringNull()

					cv := v.Get("pol-name.vipVariableName")
					item.PolicyNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.PolicyName = types.StringNull()
					item.PolicyNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("pol-name.vipValue")
					item.PolicyName = types.StringValue(cv.String())
					item.PolicyNameVariable = types.StringNull()
				}
			} else {
				item.PolicyName = types.StringNull()
				item.PolicyNameVariable = types.StringNull()
			}
			data.RoutePolicies = append(data.RoutePolicies, item)
			return true
		})
	} else {
		if len(data.RoutePolicies) > 0 {
			data.RoutePolicies = []CiscoOSPFRoutePolicies{}
		}
	}
	if value := res.Get(path + "ospf.area.vipValue"); len(value.Array()) > 0 {
		data.Areas = make([]CiscoOSPFAreas, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoOSPFAreas{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("a-num.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AreaNumber = types.Int64Null()

					cv := v.Get("a-num.vipVariableName")
					item.AreaNumberVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AreaNumber = types.Int64Null()
					item.AreaNumberVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("a-num.vipValue")
					item.AreaNumber = types.Int64Value(cv.Int())
					item.AreaNumberVariable = types.StringNull()
				}
			} else {
				item.AreaNumber = types.Int64Null()
				item.AreaNumberVariable = types.StringNull()
			}
			if cValue := v.Get("stub.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Stub = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.Stub = types.BoolValue(false)

				} else if cValue.String() == "constant" {
					cv := v.Get("stub.vipValue")
					item.Stub = types.BoolValue(cv.Bool())

				}
			} else if cValue := v.Get("stub"); cValue.Exists() {
				item.Stub = types.BoolValue(true)

			} else {
				item.Stub = types.BoolNull()

			}
			if cValue := v.Get("stub.no-summary.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.StubNoSummary = types.BoolNull()

					cv := v.Get("stub.no-summary.vipVariableName")
					item.StubNoSummaryVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.StubNoSummary = types.BoolNull()
					item.StubNoSummaryVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("stub.no-summary.vipValue")
					item.StubNoSummary = types.BoolValue(cv.Bool())
					item.StubNoSummaryVariable = types.StringNull()
				}
			} else {
				item.StubNoSummary = types.BoolNull()
				item.StubNoSummaryVariable = types.StringNull()
			}
			if cValue := v.Get("nssa.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Nssa = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.Nssa = types.BoolValue(false)

				} else if cValue.String() == "constant" {
					cv := v.Get("nssa.vipValue")
					item.Nssa = types.BoolValue(cv.Bool())

				}
			} else if cValue := v.Get("nssa"); cValue.Exists() {
				item.Nssa = types.BoolValue(true)

			} else {
				item.Nssa = types.BoolNull()

			}
			if cValue := v.Get("nssa.no-summary.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.NssaNoSummary = types.BoolNull()

					cv := v.Get("nssa.no-summary.vipVariableName")
					item.NssaNoSummaryVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.NssaNoSummary = types.BoolNull()
					item.NssaNoSummaryVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("nssa.no-summary.vipValue")
					item.NssaNoSummary = types.BoolValue(cv.Bool())
					item.NssaNoSummaryVariable = types.StringNull()
				}
			} else {
				item.NssaNoSummary = types.BoolNull()
				item.NssaNoSummaryVariable = types.StringNull()
			}
			if cValue := v.Get("interface.vipValue"); len(cValue.Array()) > 0 {
				item.Interfaces = make([]CiscoOSPFAreasInterfaces, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoOSPFAreasInterfaces{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("name.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Name = types.StringNull()

							ccv := cv.Get("name.vipVariableName")
							cItem.NameVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Name = types.StringNull()
							cItem.NameVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("name.vipValue")
							cItem.Name = types.StringValue(ccv.String())
							cItem.NameVariable = types.StringNull()
						}
					} else {
						cItem.Name = types.StringNull()
						cItem.NameVariable = types.StringNull()
					}
					if ccValue := cv.Get("hello-interval.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.HelloInterval = types.Int64Null()

							ccv := cv.Get("hello-interval.vipVariableName")
							cItem.HelloIntervalVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.HelloInterval = types.Int64Null()
							cItem.HelloIntervalVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("hello-interval.vipValue")
							cItem.HelloInterval = types.Int64Value(ccv.Int())
							cItem.HelloIntervalVariable = types.StringNull()
						}
					} else {
						cItem.HelloInterval = types.Int64Null()
						cItem.HelloIntervalVariable = types.StringNull()
					}
					if ccValue := cv.Get("dead-interval.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.DeadInterval = types.Int64Null()

							ccv := cv.Get("dead-interval.vipVariableName")
							cItem.DeadIntervalVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.DeadInterval = types.Int64Null()
							cItem.DeadIntervalVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("dead-interval.vipValue")
							cItem.DeadInterval = types.Int64Value(ccv.Int())
							cItem.DeadIntervalVariable = types.StringNull()
						}
					} else {
						cItem.DeadInterval = types.Int64Null()
						cItem.DeadIntervalVariable = types.StringNull()
					}
					if ccValue := cv.Get("retransmit-interval.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.RetransmitInterval = types.Int64Null()

							ccv := cv.Get("retransmit-interval.vipVariableName")
							cItem.RetransmitIntervalVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.RetransmitInterval = types.Int64Null()
							cItem.RetransmitIntervalVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("retransmit-interval.vipValue")
							cItem.RetransmitInterval = types.Int64Value(ccv.Int())
							cItem.RetransmitIntervalVariable = types.StringNull()
						}
					} else {
						cItem.RetransmitInterval = types.Int64Null()
						cItem.RetransmitIntervalVariable = types.StringNull()
					}
					if ccValue := cv.Get("cost.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Cost = types.Int64Null()

							ccv := cv.Get("cost.vipVariableName")
							cItem.CostVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Cost = types.Int64Null()
							cItem.CostVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("cost.vipValue")
							cItem.Cost = types.Int64Value(ccv.Int())
							cItem.CostVariable = types.StringNull()
						}
					} else {
						cItem.Cost = types.Int64Null()
						cItem.CostVariable = types.StringNull()
					}
					if ccValue := cv.Get("priority.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Priority = types.Int64Null()

							ccv := cv.Get("priority.vipVariableName")
							cItem.PriorityVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Priority = types.Int64Null()
							cItem.PriorityVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("priority.vipValue")
							cItem.Priority = types.Int64Value(ccv.Int())
							cItem.PriorityVariable = types.StringNull()
						}
					} else {
						cItem.Priority = types.Int64Null()
						cItem.PriorityVariable = types.StringNull()
					}
					if ccValue := cv.Get("network.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Network = types.StringNull()

							ccv := cv.Get("network.vipVariableName")
							cItem.NetworkVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Network = types.StringNull()
							cItem.NetworkVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("network.vipValue")
							cItem.Network = types.StringValue(ccv.String())
							cItem.NetworkVariable = types.StringNull()
						}
					} else {
						cItem.Network = types.StringNull()
						cItem.NetworkVariable = types.StringNull()
					}
					if ccValue := cv.Get("passive-interface.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.PassiveInterface = types.BoolNull()

							ccv := cv.Get("passive-interface.vipVariableName")
							cItem.PassiveInterfaceVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.PassiveInterface = types.BoolNull()
							cItem.PassiveInterfaceVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("passive-interface.vipValue")
							cItem.PassiveInterface = types.BoolValue(ccv.Bool())
							cItem.PassiveInterfaceVariable = types.StringNull()
						}
					} else {
						cItem.PassiveInterface = types.BoolNull()
						cItem.PassiveInterfaceVariable = types.StringNull()
					}
					if ccValue := cv.Get("authentication.type.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AuthenticationType = types.StringNull()

							ccv := cv.Get("authentication.type.vipVariableName")
							cItem.AuthenticationTypeVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AuthenticationType = types.StringNull()
							cItem.AuthenticationTypeVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("authentication.type.vipValue")
							cItem.AuthenticationType = types.StringValue(ccv.String())
							cItem.AuthenticationTypeVariable = types.StringNull()
						}
					} else {
						cItem.AuthenticationType = types.StringNull()
						cItem.AuthenticationTypeVariable = types.StringNull()
					}
					if ccValue := cv.Get("authentication.message-digest.message-digest-key.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AuthenticationMessageDigestKeyId = types.Int64Null()

							ccv := cv.Get("authentication.message-digest.message-digest-key.vipVariableName")
							cItem.AuthenticationMessageDigestKeyIdVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AuthenticationMessageDigestKeyId = types.Int64Null()
							cItem.AuthenticationMessageDigestKeyIdVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("authentication.message-digest.message-digest-key.vipValue")
							cItem.AuthenticationMessageDigestKeyId = types.Int64Value(ccv.Int())
							cItem.AuthenticationMessageDigestKeyIdVariable = types.StringNull()
						}
					} else {
						cItem.AuthenticationMessageDigestKeyId = types.Int64Null()
						cItem.AuthenticationMessageDigestKeyIdVariable = types.StringNull()
					}
					if ccValue := cv.Get("authentication.message-digest.md5.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AuthenticationMessageDigestKey = types.StringNull()

							ccv := cv.Get("authentication.message-digest.md5.vipVariableName")
							cItem.AuthenticationMessageDigestKeyVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AuthenticationMessageDigestKey = types.StringNull()
							cItem.AuthenticationMessageDigestKeyVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("authentication.message-digest.md5.vipValue")
							cItem.AuthenticationMessageDigestKey = types.StringValue(ccv.String())
							cItem.AuthenticationMessageDigestKeyVariable = types.StringNull()
						}
					} else {
						cItem.AuthenticationMessageDigestKey = types.StringNull()
						cItem.AuthenticationMessageDigestKeyVariable = types.StringNull()
					}
					item.Interfaces = append(item.Interfaces, cItem)
					return true
				})
			} else {
				if len(item.Interfaces) > 0 {
					item.Interfaces = []CiscoOSPFAreasInterfaces{}
				}
			}
			if cValue := v.Get("range.vipValue"); len(cValue.Array()) > 0 {
				item.Ranges = make([]CiscoOSPFAreasRanges, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoOSPFAreasRanges{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("address.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Address = types.StringNull()

							ccv := cv.Get("address.vipVariableName")
							cItem.AddressVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Address = types.StringNull()
							cItem.AddressVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("address.vipValue")
							cItem.Address = types.StringValue(ccv.String())
							cItem.AddressVariable = types.StringNull()
						}
					} else {
						cItem.Address = types.StringNull()
						cItem.AddressVariable = types.StringNull()
					}
					if ccValue := cv.Get("cost.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Cost = types.Int64Null()

							ccv := cv.Get("cost.vipVariableName")
							cItem.CostVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Cost = types.Int64Null()
							cItem.CostVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("cost.vipValue")
							cItem.Cost = types.Int64Value(ccv.Int())
							cItem.CostVariable = types.StringNull()
						}
					} else {
						cItem.Cost = types.Int64Null()
						cItem.CostVariable = types.StringNull()
					}
					if ccValue := cv.Get("no-advertise.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.NoAdvertise = types.BoolNull()

							ccv := cv.Get("no-advertise.vipVariableName")
							cItem.NoAdvertiseVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.NoAdvertise = types.BoolNull()
							cItem.NoAdvertiseVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("no-advertise.vipValue")
							cItem.NoAdvertise = types.BoolValue(ccv.Bool())
							cItem.NoAdvertiseVariable = types.StringNull()
						}
					} else {
						cItem.NoAdvertise = types.BoolNull()
						cItem.NoAdvertiseVariable = types.StringNull()
					}
					item.Ranges = append(item.Ranges, cItem)
					return true
				})
			} else {
				if len(item.Ranges) > 0 {
					item.Ranges = []CiscoOSPFAreasRanges{}
				}
			}
			data.Areas = append(data.Areas, item)
			return true
		})
	} else {
		if len(data.Areas) > 0 {
			data.Areas = []CiscoOSPFAreas{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoOSPF) hasChanges(ctx context.Context, state *CiscoOSPF) bool {
	hasChanges := false
	if !data.RouterId.Equal(state.RouterId) {
		hasChanges = true
	}
	if !data.AutoCostReferenceBandwidth.Equal(state.AutoCostReferenceBandwidth) {
		hasChanges = true
	}
	if !data.CompatibleRfc1583.Equal(state.CompatibleRfc1583) {
		hasChanges = true
	}
	if !data.DefaultInformationOriginate.Equal(state.DefaultInformationOriginate) {
		hasChanges = true
	}
	if !data.DefaultInformationOriginateAlways.Equal(state.DefaultInformationOriginateAlways) {
		hasChanges = true
	}
	if !data.DefaultInformationOriginateMetric.Equal(state.DefaultInformationOriginateMetric) {
		hasChanges = true
	}
	if !data.DefaultInformationOriginateMetricType.Equal(state.DefaultInformationOriginateMetricType) {
		hasChanges = true
	}
	if !data.DistanceExternal.Equal(state.DistanceExternal) {
		hasChanges = true
	}
	if !data.DistanceInterArea.Equal(state.DistanceInterArea) {
		hasChanges = true
	}
	if !data.DistanceIntraArea.Equal(state.DistanceIntraArea) {
		hasChanges = true
	}
	if !data.TimersSpfDelay.Equal(state.TimersSpfDelay) {
		hasChanges = true
	}
	if !data.TimersSpfInitialHold.Equal(state.TimersSpfInitialHold) {
		hasChanges = true
	}
	if !data.TimersSpfMaxHold.Equal(state.TimersSpfMaxHold) {
		hasChanges = true
	}
	if len(data.Redistribute) != len(state.Redistribute) {
		hasChanges = true
	} else {
		for i := range data.Redistribute {
			if !data.Redistribute[i].Protocol.Equal(state.Redistribute[i].Protocol) {
				hasChanges = true
			}
			if !data.Redistribute[i].RoutePolicy.Equal(state.Redistribute[i].RoutePolicy) {
				hasChanges = true
			}
			if !data.Redistribute[i].NatDia.Equal(state.Redistribute[i].NatDia) {
				hasChanges = true
			}
		}
	}
	if len(data.MaxMetricRouterLsa) != len(state.MaxMetricRouterLsa) {
		hasChanges = true
	} else {
		for i := range data.MaxMetricRouterLsa {
			if !data.MaxMetricRouterLsa[i].AdType.Equal(state.MaxMetricRouterLsa[i].AdType) {
				hasChanges = true
			}
			if !data.MaxMetricRouterLsa[i].Time.Equal(state.MaxMetricRouterLsa[i].Time) {
				hasChanges = true
			}
		}
	}
	if len(data.RoutePolicies) != len(state.RoutePolicies) {
		hasChanges = true
	} else {
		for i := range data.RoutePolicies {
			if !data.RoutePolicies[i].Direction.Equal(state.RoutePolicies[i].Direction) {
				hasChanges = true
			}
			if !data.RoutePolicies[i].PolicyName.Equal(state.RoutePolicies[i].PolicyName) {
				hasChanges = true
			}
		}
	}
	if len(data.Areas) != len(state.Areas) {
		hasChanges = true
	} else {
		for i := range data.Areas {
			if !data.Areas[i].AreaNumber.Equal(state.Areas[i].AreaNumber) {
				hasChanges = true
			}
			if !data.Areas[i].Stub.Equal(state.Areas[i].Stub) {
				hasChanges = true
			}
			if !data.Areas[i].StubNoSummary.Equal(state.Areas[i].StubNoSummary) {
				hasChanges = true
			}
			if !data.Areas[i].Nssa.Equal(state.Areas[i].Nssa) {
				hasChanges = true
			}
			if !data.Areas[i].NssaNoSummary.Equal(state.Areas[i].NssaNoSummary) {
				hasChanges = true
			}
			if len(data.Areas[i].Interfaces) != len(state.Areas[i].Interfaces) {
				hasChanges = true
			} else {
				for ii := range data.Areas[i].Interfaces {
					if !data.Areas[i].Interfaces[ii].Name.Equal(state.Areas[i].Interfaces[ii].Name) {
						hasChanges = true
					}
					if !data.Areas[i].Interfaces[ii].HelloInterval.Equal(state.Areas[i].Interfaces[ii].HelloInterval) {
						hasChanges = true
					}
					if !data.Areas[i].Interfaces[ii].DeadInterval.Equal(state.Areas[i].Interfaces[ii].DeadInterval) {
						hasChanges = true
					}
					if !data.Areas[i].Interfaces[ii].RetransmitInterval.Equal(state.Areas[i].Interfaces[ii].RetransmitInterval) {
						hasChanges = true
					}
					if !data.Areas[i].Interfaces[ii].Cost.Equal(state.Areas[i].Interfaces[ii].Cost) {
						hasChanges = true
					}
					if !data.Areas[i].Interfaces[ii].Priority.Equal(state.Areas[i].Interfaces[ii].Priority) {
						hasChanges = true
					}
					if !data.Areas[i].Interfaces[ii].Network.Equal(state.Areas[i].Interfaces[ii].Network) {
						hasChanges = true
					}
					if !data.Areas[i].Interfaces[ii].PassiveInterface.Equal(state.Areas[i].Interfaces[ii].PassiveInterface) {
						hasChanges = true
					}
					if !data.Areas[i].Interfaces[ii].AuthenticationType.Equal(state.Areas[i].Interfaces[ii].AuthenticationType) {
						hasChanges = true
					}
					if !data.Areas[i].Interfaces[ii].AuthenticationMessageDigestKeyId.Equal(state.Areas[i].Interfaces[ii].AuthenticationMessageDigestKeyId) {
						hasChanges = true
					}
					if !data.Areas[i].Interfaces[ii].AuthenticationMessageDigestKey.Equal(state.Areas[i].Interfaces[ii].AuthenticationMessageDigestKey) {
						hasChanges = true
					}
				}
			}
			if len(data.Areas[i].Ranges) != len(state.Areas[i].Ranges) {
				hasChanges = true
			} else {
				for ii := range data.Areas[i].Ranges {
					if !data.Areas[i].Ranges[ii].Address.Equal(state.Areas[i].Ranges[ii].Address) {
						hasChanges = true
					}
					if !data.Areas[i].Ranges[ii].Cost.Equal(state.Areas[i].Ranges[ii].Cost) {
						hasChanges = true
					}
					if !data.Areas[i].Ranges[ii].NoAdvertise.Equal(state.Areas[i].Ranges[ii].NoAdvertise) {
						hasChanges = true
					}
				}
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
