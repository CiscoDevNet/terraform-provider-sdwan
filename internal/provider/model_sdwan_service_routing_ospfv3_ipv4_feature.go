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
type ServiceRoutingOSPFv3IPv4 struct {
	Id                                            types.String                            `tfsdk:"id"`
	Version                                       types.Int64                             `tfsdk:"version"`
	Name                                          types.String                            `tfsdk:"name"`
	Description                                   types.String                            `tfsdk:"description"`
	FeatureProfileId                              types.String                            `tfsdk:"feature_profile_id"`
	RouterId                                      types.String                            `tfsdk:"router_id"`
	RouterIdVariable                              types.String                            `tfsdk:"router_id_variable"`
	Distance                                      types.Int64                             `tfsdk:"distance"`
	DistanceVariable                              types.String                            `tfsdk:"distance_variable"`
	DistanceExternal                              types.Int64                             `tfsdk:"distance_external"`
	DistanceExternalVariable                      types.String                            `tfsdk:"distance_external_variable"`
	DistanceInterArea                             types.Int64                             `tfsdk:"distance_inter_area"`
	DistanceInterAreaVariable                     types.String                            `tfsdk:"distance_inter_area_variable"`
	DistanceIntraArea                             types.Int64                             `tfsdk:"distance_intra_area"`
	DistanceIntraAreaVariable                     types.String                            `tfsdk:"distance_intra_area_variable"`
	ReferenceBandwidth                            types.Int64                             `tfsdk:"reference_bandwidth"`
	ReferenceBandwidthVariable                    types.String                            `tfsdk:"reference_bandwidth_variable"`
	Rfc1583Compatible                             types.Bool                              `tfsdk:"rfc_1583_compatible"`
	Rfc1583CompatibleVariable                     types.String                            `tfsdk:"rfc_1583_compatible_variable"`
	DefaultInformationOriginate                   types.Bool                              `tfsdk:"default_information_originate"`
	DefaultInformationOriginateAlways             types.Bool                              `tfsdk:"default_information_originate_always"`
	DefaultInformationOriginateAlwaysVariable     types.String                            `tfsdk:"default_information_originate_always_variable"`
	DefaultInformationOriginateMetric             types.Int64                             `tfsdk:"default_information_originate_metric"`
	DefaultInformationOriginateMetricVariable     types.String                            `tfsdk:"default_information_originate_metric_variable"`
	DefaultInformationOriginateMetricType         types.String                            `tfsdk:"default_information_originate_metric_type"`
	DefaultInformationOriginateMetricTypeVariable types.String                            `tfsdk:"default_information_originate_metric_type_variable"`
	SpfCalculationDelay                           types.Int64                             `tfsdk:"spf_calculation_delay"`
	SpfCalculationDelayVariable                   types.String                            `tfsdk:"spf_calculation_delay_variable"`
	SpfInitialHoldTime                            types.Int64                             `tfsdk:"spf_initial_hold_time"`
	SpfInitialHoldTimeVariable                    types.String                            `tfsdk:"spf_initial_hold_time_variable"`
	SpfMaximumHoldTime                            types.Int64                             `tfsdk:"spf_maximum_hold_time"`
	SpfMaximumHoldTimeVariable                    types.String                            `tfsdk:"spf_maximum_hold_time_variable"`
	RoutePolicyId                                 types.String                            `tfsdk:"route_policy_id"`
	Filter                                        types.Bool                              `tfsdk:"filter"`
	FilterVariable                                types.String                            `tfsdk:"filter_variable"`
	Redistributes                                 []ServiceRoutingOSPFv3IPv4Redistributes `tfsdk:"redistributes"`
	RouterLsaAction                               types.String                            `tfsdk:"router_lsa_action"`
	RouterLsaOnStartupTime                        types.Int64                             `tfsdk:"router_lsa_on_startup_time"`
	RouterLsaOnStartupTimeVariable                types.String                            `tfsdk:"router_lsa_on_startup_time_variable"`
	Areas                                         []ServiceRoutingOSPFv3IPv4Areas         `tfsdk:"areas"`
}

type ServiceRoutingOSPFv3IPv4Redistributes struct {
	Protocol         types.String `tfsdk:"protocol"`
	ProtocolVariable types.String `tfsdk:"protocol_variable"`
	NatDia           types.Bool   `tfsdk:"nat_dia"`
	NatDiaVariable   types.String `tfsdk:"nat_dia_variable"`
	RoutePolicyId    types.String `tfsdk:"route_policy_id"`
}

type ServiceRoutingOSPFv3IPv4Areas struct {
	AreaNumber              types.Int64                               `tfsdk:"area_number"`
	AreaNumberVariable      types.String                              `tfsdk:"area_number_variable"`
	AreaType                types.String                              `tfsdk:"area_type"`
	NoSummary               types.Bool                                `tfsdk:"no_summary"`
	NoSummaryVariable       types.String                              `tfsdk:"no_summary_variable"`
	AlwaysTranslate         types.Bool                                `tfsdk:"always_translate"`
	AlwaysTranslateVariable types.String                              `tfsdk:"always_translate_variable"`
	Interfaces              []ServiceRoutingOSPFv3IPv4AreasInterfaces `tfsdk:"interfaces"`
	Ranges                  []ServiceRoutingOSPFv3IPv4AreasRanges     `tfsdk:"ranges"`
}

type ServiceRoutingOSPFv3IPv4AreasInterfaces struct {
	Name                          types.String `tfsdk:"name"`
	NameVariable                  types.String `tfsdk:"name_variable"`
	HelloInterval                 types.Int64  `tfsdk:"hello_interval"`
	HelloIntervalVariable         types.String `tfsdk:"hello_interval_variable"`
	DeadInterval                  types.Int64  `tfsdk:"dead_interval"`
	DeadIntervalVariable          types.String `tfsdk:"dead_interval_variable"`
	LsaRetransmitInterval         types.Int64  `tfsdk:"lsa_retransmit_interval"`
	LsaRetransmitIntervalVariable types.String `tfsdk:"lsa_retransmit_interval_variable"`
	Cost                          types.Int64  `tfsdk:"cost"`
	CostVariable                  types.String `tfsdk:"cost_variable"`
	NetworkType                   types.String `tfsdk:"network_type"`
	NetworkTypeVariable           types.String `tfsdk:"network_type_variable"`
	PassiveInterface              types.Bool   `tfsdk:"passive_interface"`
	PassiveInterfaceVariable      types.String `tfsdk:"passive_interface_variable"`
	AuthenticationType            types.String `tfsdk:"authentication_type"`
	AuthenticationSpi             types.Int64  `tfsdk:"authentication_spi"`
	AuthenticationSpiVariable     types.String `tfsdk:"authentication_spi_variable"`
	AuthenticationKey             types.String `tfsdk:"authentication_key"`
	AuthenticationKeyVariable     types.String `tfsdk:"authentication_key_variable"`
}
type ServiceRoutingOSPFv3IPv4AreasRanges struct {
	IpAddress           types.String `tfsdk:"ip_address"`
	IpAddressVariable   types.String `tfsdk:"ip_address_variable"`
	SubnetMask          types.String `tfsdk:"subnet_mask"`
	SubnetMaskVariable  types.String `tfsdk:"subnet_mask_variable"`
	Cost                types.Int64  `tfsdk:"cost"`
	CostVariable        types.String `tfsdk:"cost_variable"`
	NoAdvertise         types.Bool   `tfsdk:"no_advertise"`
	NoAdvertiseVariable types.String `tfsdk:"no_advertise_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceRoutingOSPFv3IPv4) getModel() string {
	return "service_routing_ospfv3_ipv4"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceRoutingOSPFv3IPv4) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/routing/ospfv3/ipv4", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceRoutingOSPFv3IPv4) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.RouterIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.routerId.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.routerId.value", data.RouterIdVariable.ValueString())
		}
	} else if data.RouterId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.routerId.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.routerId.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.routerId.value", data.RouterId.ValueString())
		}
	}

	if !data.DistanceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.distance.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.distance.value", data.DistanceVariable.ValueString())
		}
	} else if data.Distance.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.distance.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.distance.value", 110)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.distance.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.distance.value", data.Distance.ValueInt64())
		}
	}

	if !data.DistanceExternalVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.externalDistance.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.externalDistance.value", data.DistanceExternalVariable.ValueString())
		}
	} else if data.DistanceExternal.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.externalDistance.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.externalDistance.value", 110)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.externalDistance.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.externalDistance.value", data.DistanceExternal.ValueInt64())
		}
	}

	if !data.DistanceInterAreaVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.interAreaDistance.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.interAreaDistance.value", data.DistanceInterAreaVariable.ValueString())
		}
	} else if data.DistanceInterArea.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.interAreaDistance.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.interAreaDistance.value", 110)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.interAreaDistance.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.interAreaDistance.value", data.DistanceInterArea.ValueInt64())
		}
	}

	if !data.DistanceIntraAreaVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.intraAreaDistance.optionType", "variable")
			body, _ = sjson.Set(body, path+"basic.intraAreaDistance.value", data.DistanceIntraAreaVariable.ValueString())
		}
	} else if data.DistanceIntraArea.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"basic.intraAreaDistance.optionType", "default")
			body, _ = sjson.Set(body, path+"basic.intraAreaDistance.value", 110)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"basic.intraAreaDistance.optionType", "global")
			body, _ = sjson.Set(body, path+"basic.intraAreaDistance.value", data.DistanceIntraArea.ValueInt64())
		}
	}

	if !data.ReferenceBandwidthVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.referenceBandwidth.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.referenceBandwidth.value", data.ReferenceBandwidthVariable.ValueString())
		}
	} else if data.ReferenceBandwidth.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.referenceBandwidth.optionType", "default")
			body, _ = sjson.Set(body, path+"advanced.referenceBandwidth.value", 100)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.referenceBandwidth.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.referenceBandwidth.value", data.ReferenceBandwidth.ValueInt64())
		}
	}

	if !data.Rfc1583CompatibleVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.compatibleRfc1583.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.compatibleRfc1583.value", data.Rfc1583CompatibleVariable.ValueString())
		}
	} else if data.Rfc1583Compatible.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.compatibleRfc1583.optionType", "default")
			body, _ = sjson.Set(body, path+"advanced.compatibleRfc1583.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.compatibleRfc1583.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.compatibleRfc1583.value", data.Rfc1583Compatible.ValueBool())
		}
	}
	if !data.DefaultInformationOriginate.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.defaultOriginate.originate.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.defaultOriginate.originate.value", data.DefaultInformationOriginate.ValueBool())
		}
	}

	if !data.DefaultInformationOriginateAlwaysVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.defaultOriginate.always.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.defaultOriginate.always.value", data.DefaultInformationOriginateAlwaysVariable.ValueString())
		}
	} else if !data.DefaultInformationOriginateAlways.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.defaultOriginate.always.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.defaultOriginate.always.value", data.DefaultInformationOriginateAlways.ValueBool())
		}
	}

	if !data.DefaultInformationOriginateMetricVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.defaultOriginate.metric.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.defaultOriginate.metric.value", data.DefaultInformationOriginateMetricVariable.ValueString())
		}
	} else if !data.DefaultInformationOriginateMetric.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.defaultOriginate.metric.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.defaultOriginate.metric.value", data.DefaultInformationOriginateMetric.ValueInt64())
		}
	}

	if !data.DefaultInformationOriginateMetricTypeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.defaultOriginate.metricType.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.defaultOriginate.metricType.value", data.DefaultInformationOriginateMetricTypeVariable.ValueString())
		}
	} else if !data.DefaultInformationOriginateMetricType.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.defaultOriginate.metricType.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.defaultOriginate.metricType.value", data.DefaultInformationOriginateMetricType.ValueString())
		}
	}

	if !data.SpfCalculationDelayVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.spfTimers.delay.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.spfTimers.delay.value", data.SpfCalculationDelayVariable.ValueString())
		}
	} else if data.SpfCalculationDelay.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.spfTimers.delay.optionType", "default")
			body, _ = sjson.Set(body, path+"advanced.spfTimers.delay.value", 200)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.spfTimers.delay.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.spfTimers.delay.value", data.SpfCalculationDelay.ValueInt64())
		}
	}

	if !data.SpfInitialHoldTimeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.spfTimers.initialHold.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.spfTimers.initialHold.value", data.SpfInitialHoldTimeVariable.ValueString())
		}
	} else if data.SpfInitialHoldTime.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.spfTimers.initialHold.optionType", "default")
			body, _ = sjson.Set(body, path+"advanced.spfTimers.initialHold.value", 1000)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.spfTimers.initialHold.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.spfTimers.initialHold.value", data.SpfInitialHoldTime.ValueInt64())
		}
	}

	if !data.SpfMaximumHoldTimeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.spfTimers.maxHold.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.spfTimers.maxHold.value", data.SpfMaximumHoldTimeVariable.ValueString())
		}
	} else if data.SpfMaximumHoldTime.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.spfTimers.maxHold.optionType", "default")
			body, _ = sjson.Set(body, path+"advanced.spfTimers.maxHold.value", 10000)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.spfTimers.maxHold.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.spfTimers.maxHold.value", data.SpfMaximumHoldTime.ValueInt64())
		}
	}
	if !data.RoutePolicyId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.policyName.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.policyName.refId.value", data.RoutePolicyId.ValueString())
		}
	}

	if !data.FilterVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.filter.optionType", "variable")
			body, _ = sjson.Set(body, path+"advanced.filter.value", data.FilterVariable.ValueString())
		}
	} else if data.Filter.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advanced.filter.optionType", "default")
			body, _ = sjson.Set(body, path+"advanced.filter.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advanced.filter.optionType", "global")
			body, _ = sjson.Set(body, path+"advanced.filter.value", data.Filter.ValueBool())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"redistribute", []interface{}{})
		for _, item := range data.Redistributes {
			itemBody := ""

			if !item.ProtocolVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "protocol.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "protocol.value", item.ProtocolVariable.ValueString())
				}
			} else if !item.Protocol.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "protocol.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "protocol.value", item.Protocol.ValueString())
				}
			}

			if !item.NatDiaVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "natDia.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "natDia.value", item.NatDiaVariable.ValueString())
				}
			} else if item.NatDia.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "natDia.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "natDia.value", true)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "natDia.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "natDia.value", item.NatDia.ValueBool())
				}
			}
			if !item.RoutePolicyId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.value", item.RoutePolicyId.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"redistribute.-1", itemBody)
		}
	}
	if !data.RouterLsaAction.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"maxMetricRouterLsa.action.optionType", "global")
			body, _ = sjson.Set(body, path+"maxMetricRouterLsa.action.value", data.RouterLsaAction.ValueString())
		}
	}

	if !data.RouterLsaOnStartupTimeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"maxMetricRouterLsa.onStartUpTime.optionType", "variable")
			body, _ = sjson.Set(body, path+"maxMetricRouterLsa.onStartUpTime.value", data.RouterLsaOnStartupTimeVariable.ValueString())
		}
	} else if !data.RouterLsaOnStartupTime.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"maxMetricRouterLsa.onStartUpTime.optionType", "global")
			body, _ = sjson.Set(body, path+"maxMetricRouterLsa.onStartUpTime.value", data.RouterLsaOnStartupTime.ValueInt64())
		}
	}
	if true {

		for _, item := range data.Areas {
			itemBody := ""

			if !item.AreaNumberVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "areaNum.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "areaNum.value", item.AreaNumberVariable.ValueString())
				}
			} else if !item.AreaNumber.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "areaNum.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "areaNum.value", item.AreaNumber.ValueInt64())
				}
			}
			if !item.AreaType.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.areaType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.areaType.value", item.AreaType.ValueString())
				}
			}

			if !item.NoSummaryVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.noSummary.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.noSummary.value", item.NoSummaryVariable.ValueString())
				}
			} else if !item.NoSummary.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.noSummary.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.noSummary.value", item.NoSummary.ValueBool())
				}
			}

			if !item.AlwaysTranslateVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.alwaysTranslate.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.alwaysTranslate.value", item.AlwaysTranslateVariable.ValueString())
				}
			} else if !item.AlwaysTranslate.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.alwaysTranslate.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "areaTypeConfig.alwaysTranslate.value", item.AlwaysTranslate.ValueBool())
				}
			}
			if true {

				for _, childItem := range item.Interfaces {
					itemChildBody := ""

					if !childItem.NameVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "ifName.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "ifName.value", childItem.NameVariable.ValueString())
						}
					} else if !childItem.Name.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "ifName.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "ifName.value", childItem.Name.ValueString())
						}
					}

					if !childItem.HelloIntervalVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "helloInterval.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "helloInterval.value", childItem.HelloIntervalVariable.ValueString())
						}
					} else if childItem.HelloInterval.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "helloInterval.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "helloInterval.value", 10)
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "helloInterval.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "helloInterval.value", childItem.HelloInterval.ValueInt64())
						}
					}

					if !childItem.DeadIntervalVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "deadInterval.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "deadInterval.value", childItem.DeadIntervalVariable.ValueString())
						}
					} else if childItem.DeadInterval.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "deadInterval.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "deadInterval.value", 40)
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "deadInterval.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "deadInterval.value", childItem.DeadInterval.ValueInt64())
						}
					}

					if !childItem.LsaRetransmitIntervalVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "retransmitInterval.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "retransmitInterval.value", childItem.LsaRetransmitIntervalVariable.ValueString())
						}
					} else if childItem.LsaRetransmitInterval.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "retransmitInterval.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "retransmitInterval.value", 5)
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "retransmitInterval.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "retransmitInterval.value", childItem.LsaRetransmitInterval.ValueInt64())
						}
					}

					if !childItem.CostVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "cost.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "cost.value", childItem.CostVariable.ValueString())
						}
					} else if childItem.Cost.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "cost.optionType", "default")

						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "cost.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "cost.value", childItem.Cost.ValueInt64())
						}
					}

					if !childItem.NetworkTypeVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "networkType.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "networkType.value", childItem.NetworkTypeVariable.ValueString())
						}
					} else if childItem.NetworkType.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "networkType.optionType", "default")

						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "networkType.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "networkType.value", childItem.NetworkType.ValueString())
						}
					}

					if !childItem.PassiveInterfaceVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "passiveInterface.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "passiveInterface.value", childItem.PassiveInterfaceVariable.ValueString())
						}
					} else if childItem.PassiveInterface.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "passiveInterface.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "passiveInterface.value", false)
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "passiveInterface.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "passiveInterface.value", childItem.PassiveInterface.ValueBool())
						}
					}
					if !childItem.AuthenticationType.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.authType.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.authType.value", childItem.AuthenticationType.ValueString())
						}
					}

					if !childItem.AuthenticationSpiVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.spi.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.spi.value", childItem.AuthenticationSpiVariable.ValueString())
						}
					} else if !childItem.AuthenticationSpi.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.spi.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.spi.value", childItem.AuthenticationSpi.ValueInt64())
						}
					}

					if !childItem.AuthenticationKeyVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.authKey.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.authKey.value", childItem.AuthenticationKeyVariable.ValueString())
						}
					} else if !childItem.AuthenticationKey.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.authKey.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "authenticationConfig.authKey.value", childItem.AuthenticationKey.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "interfaces.-1", itemChildBody)
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "ranges", []interface{}{})
				for _, childItem := range item.Ranges {
					itemChildBody := ""

					if !childItem.IpAddressVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "network.address.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "network.address.value", childItem.IpAddressVariable.ValueString())
						}
					} else if !childItem.IpAddress.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "network.address.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "network.address.value", childItem.IpAddress.ValueString())
						}
					}

					if !childItem.SubnetMaskVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "network.mask.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "network.mask.value", childItem.SubnetMaskVariable.ValueString())
						}
					} else if !childItem.SubnetMask.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "network.mask.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "network.mask.value", childItem.SubnetMask.ValueString())
						}
					}

					if !childItem.CostVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "cost.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "cost.value", childItem.CostVariable.ValueString())
						}
					} else if childItem.Cost.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "cost.optionType", "default")

						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "cost.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "cost.value", childItem.Cost.ValueInt64())
						}
					}

					if !childItem.NoAdvertiseVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "noAdvertise.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "noAdvertise.value", childItem.NoAdvertiseVariable.ValueString())
						}
					} else if childItem.NoAdvertise.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "noAdvertise.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "noAdvertise.value", false)
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "noAdvertise.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "noAdvertise.value", childItem.NoAdvertise.ValueBool())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ranges.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"area.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceRoutingOSPFv3IPv4) fromBody(ctx context.Context, res gjson.Result) {
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
	data.DistanceExternal = types.Int64Null()
	data.DistanceExternalVariable = types.StringNull()
	if t := res.Get(path + "basic.externalDistance.optionType"); t.Exists() {
		va := res.Get(path + "basic.externalDistance.value")
		if t.String() == "variable" {
			data.DistanceExternalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceExternal = types.Int64Value(va.Int())
		}
	}
	data.DistanceInterArea = types.Int64Null()
	data.DistanceInterAreaVariable = types.StringNull()
	if t := res.Get(path + "basic.interAreaDistance.optionType"); t.Exists() {
		va := res.Get(path + "basic.interAreaDistance.value")
		if t.String() == "variable" {
			data.DistanceInterAreaVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceInterArea = types.Int64Value(va.Int())
		}
	}
	data.DistanceIntraArea = types.Int64Null()
	data.DistanceIntraAreaVariable = types.StringNull()
	if t := res.Get(path + "basic.intraAreaDistance.optionType"); t.Exists() {
		va := res.Get(path + "basic.intraAreaDistance.value")
		if t.String() == "variable" {
			data.DistanceIntraAreaVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceIntraArea = types.Int64Value(va.Int())
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
	data.DefaultInformationOriginate = types.BoolNull()

	if t := res.Get(path + "advanced.defaultOriginate.originate.optionType"); t.Exists() {
		va := res.Get(path + "advanced.defaultOriginate.originate.value")
		if t.String() == "global" {
			data.DefaultInformationOriginate = types.BoolValue(va.Bool())
		}
	}
	data.DefaultInformationOriginateAlways = types.BoolNull()
	data.DefaultInformationOriginateAlwaysVariable = types.StringNull()
	if t := res.Get(path + "advanced.defaultOriginate.always.optionType"); t.Exists() {
		va := res.Get(path + "advanced.defaultOriginate.always.value")
		if t.String() == "variable" {
			data.DefaultInformationOriginateAlwaysVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultInformationOriginateAlways = types.BoolValue(va.Bool())
		}
	}
	data.DefaultInformationOriginateMetric = types.Int64Null()
	data.DefaultInformationOriginateMetricVariable = types.StringNull()
	if t := res.Get(path + "advanced.defaultOriginate.metric.optionType"); t.Exists() {
		va := res.Get(path + "advanced.defaultOriginate.metric.value")
		if t.String() == "variable" {
			data.DefaultInformationOriginateMetricVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultInformationOriginateMetric = types.Int64Value(va.Int())
		}
	}
	data.DefaultInformationOriginateMetricType = types.StringNull()
	data.DefaultInformationOriginateMetricTypeVariable = types.StringNull()
	if t := res.Get(path + "advanced.defaultOriginate.metricType.optionType"); t.Exists() {
		va := res.Get(path + "advanced.defaultOriginate.metricType.value")
		if t.String() == "variable" {
			data.DefaultInformationOriginateMetricTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultInformationOriginateMetricType = types.StringValue(va.String())
		}
	}
	data.SpfCalculationDelay = types.Int64Null()
	data.SpfCalculationDelayVariable = types.StringNull()
	if t := res.Get(path + "advanced.spfTimers.delay.optionType"); t.Exists() {
		va := res.Get(path + "advanced.spfTimers.delay.value")
		if t.String() == "variable" {
			data.SpfCalculationDelayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SpfCalculationDelay = types.Int64Value(va.Int())
		}
	}
	data.SpfInitialHoldTime = types.Int64Null()
	data.SpfInitialHoldTimeVariable = types.StringNull()
	if t := res.Get(path + "advanced.spfTimers.initialHold.optionType"); t.Exists() {
		va := res.Get(path + "advanced.spfTimers.initialHold.value")
		if t.String() == "variable" {
			data.SpfInitialHoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SpfInitialHoldTime = types.Int64Value(va.Int())
		}
	}
	data.SpfMaximumHoldTime = types.Int64Null()
	data.SpfMaximumHoldTimeVariable = types.StringNull()
	if t := res.Get(path + "advanced.spfTimers.maxHold.optionType"); t.Exists() {
		va := res.Get(path + "advanced.spfTimers.maxHold.value")
		if t.String() == "variable" {
			data.SpfMaximumHoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SpfMaximumHoldTime = types.Int64Value(va.Int())
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
	if value := res.Get(path + "redistribute"); value.Exists() && len(value.Array()) > 0 {
		data.Redistributes = make([]ServiceRoutingOSPFv3IPv4Redistributes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingOSPFv3IPv4Redistributes{}
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
			item.NatDia = types.BoolNull()
			item.NatDiaVariable = types.StringNull()
			if t := v.Get("natDia.optionType"); t.Exists() {
				va := v.Get("natDia.value")
				if t.String() == "variable" {
					item.NatDiaVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NatDia = types.BoolValue(va.Bool())
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
	data.RouterLsaOnStartupTime = types.Int64Null()
	data.RouterLsaOnStartupTimeVariable = types.StringNull()
	if t := res.Get(path + "maxMetricRouterLsa.onStartUpTime.optionType"); t.Exists() {
		va := res.Get(path + "maxMetricRouterLsa.onStartUpTime.value")
		if t.String() == "variable" {
			data.RouterLsaOnStartupTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RouterLsaOnStartupTime = types.Int64Value(va.Int())
		}
	}
	if value := res.Get(path + "area"); value.Exists() && len(value.Array()) > 0 {
		data.Areas = make([]ServiceRoutingOSPFv3IPv4Areas, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingOSPFv3IPv4Areas{}
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
			if cValue := v.Get("interfaces"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Interfaces = make([]ServiceRoutingOSPFv3IPv4AreasInterfaces, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceRoutingOSPFv3IPv4AreasInterfaces{}
					cItem.Name = types.StringNull()
					cItem.NameVariable = types.StringNull()
					if t := cv.Get("ifName.optionType"); t.Exists() {
						va := cv.Get("ifName.value")
						if t.String() == "variable" {
							cItem.NameVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Name = types.StringValue(va.String())
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
					cItem.NetworkType = types.StringNull()
					cItem.NetworkTypeVariable = types.StringNull()
					if t := cv.Get("networkType.optionType"); t.Exists() {
						va := cv.Get("networkType.value")
						if t.String() == "variable" {
							cItem.NetworkTypeVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.NetworkType = types.StringValue(va.String())
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
					cItem.AuthenticationType = types.StringNull()

					if t := cv.Get("authenticationConfig.authType.optionType"); t.Exists() {
						va := cv.Get("authenticationConfig.authType.value")
						if t.String() == "global" {
							cItem.AuthenticationType = types.StringValue(va.String())
						}
					}
					cItem.AuthenticationSpi = types.Int64Null()
					cItem.AuthenticationSpiVariable = types.StringNull()
					if t := cv.Get("authenticationConfig.spi.optionType"); t.Exists() {
						va := cv.Get("authenticationConfig.spi.value")
						if t.String() == "variable" {
							cItem.AuthenticationSpiVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.AuthenticationSpi = types.Int64Value(va.Int())
						}
					}
					cItem.AuthenticationKey = types.StringNull()
					cItem.AuthenticationKeyVariable = types.StringNull()
					if t := cv.Get("authenticationConfig.authKey.optionType"); t.Exists() {
						va := cv.Get("authenticationConfig.authKey.value")
						if t.String() == "variable" {
							cItem.AuthenticationKeyVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.AuthenticationKey = types.StringValue(va.String())
						}
					}
					item.Interfaces = append(item.Interfaces, cItem)
					return true
				})
			}
			if cValue := v.Get("ranges"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Ranges = make([]ServiceRoutingOSPFv3IPv4AreasRanges, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceRoutingOSPFv3IPv4AreasRanges{}
					cItem.IpAddress = types.StringNull()
					cItem.IpAddressVariable = types.StringNull()
					if t := cv.Get("network.address.optionType"); t.Exists() {
						va := cv.Get("network.address.value")
						if t.String() == "variable" {
							cItem.IpAddressVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.IpAddress = types.StringValue(va.String())
						}
					}
					cItem.SubnetMask = types.StringNull()
					cItem.SubnetMaskVariable = types.StringNull()
					if t := cv.Get("network.mask.optionType"); t.Exists() {
						va := cv.Get("network.mask.value")
						if t.String() == "variable" {
							cItem.SubnetMaskVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.SubnetMask = types.StringValue(va.String())
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
func (data *ServiceRoutingOSPFv3IPv4) updateFromBody(ctx context.Context, res gjson.Result) {
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
	data.DistanceExternal = types.Int64Null()
	data.DistanceExternalVariable = types.StringNull()
	if t := res.Get(path + "basic.externalDistance.optionType"); t.Exists() {
		va := res.Get(path + "basic.externalDistance.value")
		if t.String() == "variable" {
			data.DistanceExternalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceExternal = types.Int64Value(va.Int())
		}
	}
	data.DistanceInterArea = types.Int64Null()
	data.DistanceInterAreaVariable = types.StringNull()
	if t := res.Get(path + "basic.interAreaDistance.optionType"); t.Exists() {
		va := res.Get(path + "basic.interAreaDistance.value")
		if t.String() == "variable" {
			data.DistanceInterAreaVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceInterArea = types.Int64Value(va.Int())
		}
	}
	data.DistanceIntraArea = types.Int64Null()
	data.DistanceIntraAreaVariable = types.StringNull()
	if t := res.Get(path + "basic.intraAreaDistance.optionType"); t.Exists() {
		va := res.Get(path + "basic.intraAreaDistance.value")
		if t.String() == "variable" {
			data.DistanceIntraAreaVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceIntraArea = types.Int64Value(va.Int())
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
	data.DefaultInformationOriginate = types.BoolNull()

	if t := res.Get(path + "advanced.defaultOriginate.originate.optionType"); t.Exists() {
		va := res.Get(path + "advanced.defaultOriginate.originate.value")
		if t.String() == "global" {
			data.DefaultInformationOriginate = types.BoolValue(va.Bool())
		}
	}
	data.DefaultInformationOriginateAlways = types.BoolNull()
	data.DefaultInformationOriginateAlwaysVariable = types.StringNull()
	if t := res.Get(path + "advanced.defaultOriginate.always.optionType"); t.Exists() {
		va := res.Get(path + "advanced.defaultOriginate.always.value")
		if t.String() == "variable" {
			data.DefaultInformationOriginateAlwaysVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultInformationOriginateAlways = types.BoolValue(va.Bool())
		}
	}
	data.DefaultInformationOriginateMetric = types.Int64Null()
	data.DefaultInformationOriginateMetricVariable = types.StringNull()
	if t := res.Get(path + "advanced.defaultOriginate.metric.optionType"); t.Exists() {
		va := res.Get(path + "advanced.defaultOriginate.metric.value")
		if t.String() == "variable" {
			data.DefaultInformationOriginateMetricVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultInformationOriginateMetric = types.Int64Value(va.Int())
		}
	}
	data.DefaultInformationOriginateMetricType = types.StringNull()
	data.DefaultInformationOriginateMetricTypeVariable = types.StringNull()
	if t := res.Get(path + "advanced.defaultOriginate.metricType.optionType"); t.Exists() {
		va := res.Get(path + "advanced.defaultOriginate.metricType.value")
		if t.String() == "variable" {
			data.DefaultInformationOriginateMetricTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultInformationOriginateMetricType = types.StringValue(va.String())
		}
	}
	data.SpfCalculationDelay = types.Int64Null()
	data.SpfCalculationDelayVariable = types.StringNull()
	if t := res.Get(path + "advanced.spfTimers.delay.optionType"); t.Exists() {
		va := res.Get(path + "advanced.spfTimers.delay.value")
		if t.String() == "variable" {
			data.SpfCalculationDelayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SpfCalculationDelay = types.Int64Value(va.Int())
		}
	}
	data.SpfInitialHoldTime = types.Int64Null()
	data.SpfInitialHoldTimeVariable = types.StringNull()
	if t := res.Get(path + "advanced.spfTimers.initialHold.optionType"); t.Exists() {
		va := res.Get(path + "advanced.spfTimers.initialHold.value")
		if t.String() == "variable" {
			data.SpfInitialHoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SpfInitialHoldTime = types.Int64Value(va.Int())
		}
	}
	data.SpfMaximumHoldTime = types.Int64Null()
	data.SpfMaximumHoldTimeVariable = types.StringNull()
	if t := res.Get(path + "advanced.spfTimers.maxHold.optionType"); t.Exists() {
		va := res.Get(path + "advanced.spfTimers.maxHold.value")
		if t.String() == "variable" {
			data.SpfMaximumHoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SpfMaximumHoldTime = types.Int64Value(va.Int())
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
		data.Redistributes[i].NatDia = types.BoolNull()
		data.Redistributes[i].NatDiaVariable = types.StringNull()
		if t := r.Get("natDia.optionType"); t.Exists() {
			va := r.Get("natDia.value")
			if t.String() == "variable" {
				data.Redistributes[i].NatDiaVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Redistributes[i].NatDia = types.BoolValue(va.Bool())
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
	data.RouterLsaOnStartupTime = types.Int64Null()
	data.RouterLsaOnStartupTimeVariable = types.StringNull()
	if t := res.Get(path + "maxMetricRouterLsa.onStartUpTime.optionType"); t.Exists() {
		va := res.Get(path + "maxMetricRouterLsa.onStartUpTime.value")
		if t.String() == "variable" {
			data.RouterLsaOnStartupTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RouterLsaOnStartupTime = types.Int64Value(va.Int())
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
			keyValues := [...]string{data.Areas[i].Interfaces[ci].Name.ValueString()}
			keyValuesVariables := [...]string{data.Areas[i].Interfaces[ci].NameVariable.ValueString()}

			var cr gjson.Result
			r.Get("interfaces").ForEach(
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
			data.Areas[i].Interfaces[ci].Name = types.StringNull()
			data.Areas[i].Interfaces[ci].NameVariable = types.StringNull()
			if t := cr.Get("ifName.optionType"); t.Exists() {
				va := cr.Get("ifName.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].NameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].Name = types.StringValue(va.String())
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
			data.Areas[i].Interfaces[ci].Cost = types.Int64Null()
			data.Areas[i].Interfaces[ci].CostVariable = types.StringNull()
			if t := cr.Get("cost.optionType"); t.Exists() {
				va := cr.Get("cost.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].CostVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].Cost = types.Int64Value(va.Int())
				}
			}
			data.Areas[i].Interfaces[ci].NetworkType = types.StringNull()
			data.Areas[i].Interfaces[ci].NetworkTypeVariable = types.StringNull()
			if t := cr.Get("networkType.optionType"); t.Exists() {
				va := cr.Get("networkType.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].NetworkTypeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].NetworkType = types.StringValue(va.String())
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
			data.Areas[i].Interfaces[ci].AuthenticationType = types.StringNull()

			if t := cr.Get("authenticationConfig.authType.optionType"); t.Exists() {
				va := cr.Get("authenticationConfig.authType.value")
				if t.String() == "global" {
					data.Areas[i].Interfaces[ci].AuthenticationType = types.StringValue(va.String())
				}
			}
			data.Areas[i].Interfaces[ci].AuthenticationSpi = types.Int64Null()
			data.Areas[i].Interfaces[ci].AuthenticationSpiVariable = types.StringNull()
			if t := cr.Get("authenticationConfig.spi.optionType"); t.Exists() {
				va := cr.Get("authenticationConfig.spi.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].AuthenticationSpiVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].AuthenticationSpi = types.Int64Value(va.Int())
				}
			}
			data.Areas[i].Interfaces[ci].AuthenticationKey = types.StringNull()
			data.Areas[i].Interfaces[ci].AuthenticationKeyVariable = types.StringNull()
			if t := cr.Get("authenticationConfig.authKey.optionType"); t.Exists() {
				va := cr.Get("authenticationConfig.authKey.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].AuthenticationKeyVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].AuthenticationKey = types.StringValue(va.String())
				}
			}
		}
		for ci := range data.Areas[i].Ranges {
			keys := [...]string{"network.address"}
			keyValues := [...]string{data.Areas[i].Ranges[ci].IpAddress.ValueString()}
			keyValuesVariables := [...]string{data.Areas[i].Ranges[ci].IpAddressVariable.ValueString()}

			var cr gjson.Result
			r.Get("ranges").ForEach(
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
			data.Areas[i].Ranges[ci].IpAddress = types.StringNull()
			data.Areas[i].Ranges[ci].IpAddressVariable = types.StringNull()
			if t := cr.Get("network.address.optionType"); t.Exists() {
				va := cr.Get("network.address.value")
				if t.String() == "variable" {
					data.Areas[i].Ranges[ci].IpAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Ranges[ci].IpAddress = types.StringValue(va.String())
				}
			}
			data.Areas[i].Ranges[ci].SubnetMask = types.StringNull()
			data.Areas[i].Ranges[ci].SubnetMaskVariable = types.StringNull()
			if t := cr.Get("network.mask.optionType"); t.Exists() {
				va := cr.Get("network.mask.value")
				if t.String() == "variable" {
					data.Areas[i].Ranges[ci].SubnetMaskVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Ranges[ci].SubnetMask = types.StringValue(va.String())
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
