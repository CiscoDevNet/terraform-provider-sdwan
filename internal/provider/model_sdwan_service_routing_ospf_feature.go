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
type ServiceRoutingOSPF struct {
	Id                                            types.String                      `tfsdk:"id"`
	Version                                       types.Int64                       `tfsdk:"version"`
	Name                                          types.String                      `tfsdk:"name"`
	Description                                   types.String                      `tfsdk:"description"`
	FeatureProfileId                              types.String                      `tfsdk:"feature_profile_id"`
	RouterId                                      types.String                      `tfsdk:"router_id"`
	RouterIdVariable                              types.String                      `tfsdk:"router_id_variable"`
	ReferenceBandwidth                            types.Int64                       `tfsdk:"reference_bandwidth"`
	ReferenceBandwidthVariable                    types.String                      `tfsdk:"reference_bandwidth_variable"`
	Rfc1583Compatible                             types.Bool                        `tfsdk:"rfc_1583_compatible"`
	Rfc1583CompatibleVariable                     types.String                      `tfsdk:"rfc_1583_compatible_variable"`
	DefaultInformationOriginate                   types.Bool                        `tfsdk:"default_information_originate"`
	DefaultInformationOriginateAlways             types.Bool                        `tfsdk:"default_information_originate_always"`
	DefaultInformationOriginateAlwaysVariable     types.String                      `tfsdk:"default_information_originate_always_variable"`
	DefaultInformationOriginateMetric             types.Int64                       `tfsdk:"default_information_originate_metric"`
	DefaultInformationOriginateMetricVariable     types.String                      `tfsdk:"default_information_originate_metric_variable"`
	DefaultInformationOriginateMetricType         types.String                      `tfsdk:"default_information_originate_metric_type"`
	DefaultInformationOriginateMetricTypeVariable types.String                      `tfsdk:"default_information_originate_metric_type_variable"`
	DistanceExternal                              types.Int64                       `tfsdk:"distance_external"`
	DistanceExternalVariable                      types.String                      `tfsdk:"distance_external_variable"`
	DistanceInterArea                             types.Int64                       `tfsdk:"distance_inter_area"`
	DistanceInterAreaVariable                     types.String                      `tfsdk:"distance_inter_area_variable"`
	DistanceIntraArea                             types.Int64                       `tfsdk:"distance_intra_area"`
	DistanceIntraAreaVariable                     types.String                      `tfsdk:"distance_intra_area_variable"`
	SpfCalculationDelay                           types.Int64                       `tfsdk:"spf_calculation_delay"`
	SpfCalculationDelayVariable                   types.String                      `tfsdk:"spf_calculation_delay_variable"`
	SpfInitialHoldTime                            types.Int64                       `tfsdk:"spf_initial_hold_time"`
	SpfInitialHoldTimeVariable                    types.String                      `tfsdk:"spf_initial_hold_time_variable"`
	SpfMaximumHoldTime                            types.Int64                       `tfsdk:"spf_maximum_hold_time"`
	SpfMaximumHoldTimeVariable                    types.String                      `tfsdk:"spf_maximum_hold_time_variable"`
	Redistributes                                 []ServiceRoutingOSPFRedistributes `tfsdk:"redistributes"`
	RouterLsas                                    []ServiceRoutingOSPFRouterLsas    `tfsdk:"router_lsas"`
	RoutePolicyId                                 types.String                      `tfsdk:"route_policy_id"`
	Areas                                         []ServiceRoutingOSPFAreas         `tfsdk:"areas"`
}

type ServiceRoutingOSPFRedistributes struct {
	Protocol         types.String `tfsdk:"protocol"`
	ProtocolVariable types.String `tfsdk:"protocol_variable"`
	NatDia           types.Bool   `tfsdk:"nat_dia"`
	NatDiaVariable   types.String `tfsdk:"nat_dia_variable"`
	RoutePolicyId    types.String `tfsdk:"route_policy_id"`
}

type ServiceRoutingOSPFRouterLsas struct {
	Type         types.String `tfsdk:"type"`
	Time         types.Int64  `tfsdk:"time"`
	TimeVariable types.String `tfsdk:"time_variable"`
}

type ServiceRoutingOSPFAreas struct {
	AreaNumber         types.Int64                         `tfsdk:"area_number"`
	AreaNumberVariable types.String                        `tfsdk:"area_number_variable"`
	AreaType           types.String                        `tfsdk:"area_type"`
	NoSummary          types.Bool                          `tfsdk:"no_summary"`
	NoSummaryVariable  types.String                        `tfsdk:"no_summary_variable"`
	Interfaces         []ServiceRoutingOSPFAreasInterfaces `tfsdk:"interfaces"`
	Ranges             []ServiceRoutingOSPFAreasRanges     `tfsdk:"ranges"`
}

type ServiceRoutingOSPFAreasInterfaces struct {
	Name                             types.String `tfsdk:"name"`
	NameVariable                     types.String `tfsdk:"name_variable"`
	HelloInterval                    types.Int64  `tfsdk:"hello_interval"`
	HelloIntervalVariable            types.String `tfsdk:"hello_interval_variable"`
	DeadInterval                     types.Int64  `tfsdk:"dead_interval"`
	DeadIntervalVariable             types.String `tfsdk:"dead_interval_variable"`
	LsaRetransmitInterval            types.Int64  `tfsdk:"lsa_retransmit_interval"`
	LsaRetransmitIntervalVariable    types.String `tfsdk:"lsa_retransmit_interval_variable"`
	Cost                             types.Int64  `tfsdk:"cost"`
	CostVariable                     types.String `tfsdk:"cost_variable"`
	DesignatedRouterPriority         types.Int64  `tfsdk:"designated_router_priority"`
	DesignatedRouterPriorityVariable types.String `tfsdk:"designated_router_priority_variable"`
	NetworkType                      types.String `tfsdk:"network_type"`
	NetworkTypeVariable              types.String `tfsdk:"network_type_variable"`
	PassiveInterface                 types.Bool   `tfsdk:"passive_interface"`
	PassiveInterfaceVariable         types.String `tfsdk:"passive_interface_variable"`
	AuthenticationType               types.String `tfsdk:"authentication_type"`
	AuthenticationTypeVariable       types.String `tfsdk:"authentication_type_variable"`
	MessageDigestKeyId               types.Int64  `tfsdk:"message_digest_key_id"`
	MessageDigestKeyIdVariable       types.String `tfsdk:"message_digest_key_id_variable"`
	MessageDigestKey                 types.String `tfsdk:"message_digest_key"`
	MessageDigestKeyVariable         types.String `tfsdk:"message_digest_key_variable"`
}
type ServiceRoutingOSPFAreasRanges struct {
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
func (data ServiceRoutingOSPF) getModel() string {
	return "service_routing_ospf"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceRoutingOSPF) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/routing/ospf", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceRoutingOSPF) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.RouterIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"routerId.optionType", "variable")
			body, _ = sjson.Set(body, path+"routerId.value", data.RouterIdVariable.ValueString())
		}
	} else if data.RouterId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"routerId.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"routerId.optionType", "global")
			body, _ = sjson.Set(body, path+"routerId.value", data.RouterId.ValueString())
		}
	}

	if !data.ReferenceBandwidthVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"referenceBandwidth.optionType", "variable")
			body, _ = sjson.Set(body, path+"referenceBandwidth.value", data.ReferenceBandwidthVariable.ValueString())
		}
	} else if data.ReferenceBandwidth.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"referenceBandwidth.optionType", "default")
			body, _ = sjson.Set(body, path+"referenceBandwidth.value", 100)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"referenceBandwidth.optionType", "global")
			body, _ = sjson.Set(body, path+"referenceBandwidth.value", data.ReferenceBandwidth.ValueInt64())
		}
	}

	if !data.Rfc1583CompatibleVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"rfc1583.optionType", "variable")
			body, _ = sjson.Set(body, path+"rfc1583.value", data.Rfc1583CompatibleVariable.ValueString())
		}
	} else if data.Rfc1583Compatible.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"rfc1583.optionType", "default")
			body, _ = sjson.Set(body, path+"rfc1583.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"rfc1583.optionType", "global")
			body, _ = sjson.Set(body, path+"rfc1583.value", data.Rfc1583Compatible.ValueBool())
		}
	}
	if data.DefaultInformationOriginate.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"originate.optionType", "default")
			body, _ = sjson.Set(body, path+"originate.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"originate.optionType", "global")
			body, _ = sjson.Set(body, path+"originate.value", data.DefaultInformationOriginate.ValueBool())
		}
	}

	if !data.DefaultInformationOriginateAlwaysVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"always.optionType", "variable")
			body, _ = sjson.Set(body, path+"always.value", data.DefaultInformationOriginateAlwaysVariable.ValueString())
		}
	} else if data.DefaultInformationOriginateAlways.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"always.optionType", "default")
			body, _ = sjson.Set(body, path+"always.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"always.optionType", "global")
			body, _ = sjson.Set(body, path+"always.value", data.DefaultInformationOriginateAlways.ValueBool())
		}
	}

	if !data.DefaultInformationOriginateMetricVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"metric.optionType", "variable")
			body, _ = sjson.Set(body, path+"metric.value", data.DefaultInformationOriginateMetricVariable.ValueString())
		}
	} else if data.DefaultInformationOriginateMetric.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"metric.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"metric.optionType", "global")
			body, _ = sjson.Set(body, path+"metric.value", data.DefaultInformationOriginateMetric.ValueInt64())
		}
	}

	if !data.DefaultInformationOriginateMetricTypeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"metricType.optionType", "variable")
			body, _ = sjson.Set(body, path+"metricType.value", data.DefaultInformationOriginateMetricTypeVariable.ValueString())
		}
	} else if data.DefaultInformationOriginateMetricType.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"metricType.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"metricType.optionType", "global")
			body, _ = sjson.Set(body, path+"metricType.value", data.DefaultInformationOriginateMetricType.ValueString())
		}
	}

	if !data.DistanceExternalVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"external.optionType", "variable")
			body, _ = sjson.Set(body, path+"external.value", data.DistanceExternalVariable.ValueString())
		}
	} else if data.DistanceExternal.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"external.optionType", "default")
			body, _ = sjson.Set(body, path+"external.value", 110)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"external.optionType", "global")
			body, _ = sjson.Set(body, path+"external.value", data.DistanceExternal.ValueInt64())
		}
	}

	if !data.DistanceInterAreaVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"interArea.optionType", "variable")
			body, _ = sjson.Set(body, path+"interArea.value", data.DistanceInterAreaVariable.ValueString())
		}
	} else if data.DistanceInterArea.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"interArea.optionType", "default")
			body, _ = sjson.Set(body, path+"interArea.value", 110)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"interArea.optionType", "global")
			body, _ = sjson.Set(body, path+"interArea.value", data.DistanceInterArea.ValueInt64())
		}
	}

	if !data.DistanceIntraAreaVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"intraArea.optionType", "variable")
			body, _ = sjson.Set(body, path+"intraArea.value", data.DistanceIntraAreaVariable.ValueString())
		}
	} else if data.DistanceIntraArea.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"intraArea.optionType", "default")
			body, _ = sjson.Set(body, path+"intraArea.value", 110)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"intraArea.optionType", "global")
			body, _ = sjson.Set(body, path+"intraArea.value", data.DistanceIntraArea.ValueInt64())
		}
	}

	if !data.SpfCalculationDelayVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"delay.optionType", "variable")
			body, _ = sjson.Set(body, path+"delay.value", data.SpfCalculationDelayVariable.ValueString())
		}
	} else if data.SpfCalculationDelay.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"delay.optionType", "default")
			body, _ = sjson.Set(body, path+"delay.value", 200)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"delay.optionType", "global")
			body, _ = sjson.Set(body, path+"delay.value", data.SpfCalculationDelay.ValueInt64())
		}
	}

	if !data.SpfInitialHoldTimeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"initialHold.optionType", "variable")
			body, _ = sjson.Set(body, path+"initialHold.value", data.SpfInitialHoldTimeVariable.ValueString())
		}
	} else if data.SpfInitialHoldTime.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"initialHold.optionType", "default")
			body, _ = sjson.Set(body, path+"initialHold.value", 1000)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"initialHold.optionType", "global")
			body, _ = sjson.Set(body, path+"initialHold.value", data.SpfInitialHoldTime.ValueInt64())
		}
	}

	if !data.SpfMaximumHoldTimeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"maxHold.optionType", "variable")
			body, _ = sjson.Set(body, path+"maxHold.value", data.SpfMaximumHoldTimeVariable.ValueString())
		}
	} else if data.SpfMaximumHoldTime.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"maxHold.optionType", "default")
			body, _ = sjson.Set(body, path+"maxHold.value", 10000)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"maxHold.optionType", "global")
			body, _ = sjson.Set(body, path+"maxHold.value", data.SpfMaximumHoldTime.ValueInt64())
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
					itemBody, _ = sjson.Set(itemBody, "dia.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "dia.value", item.NatDiaVariable.ValueString())
				}
			} else if item.NatDia.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dia.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "dia.value", true)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dia.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "dia.value", item.NatDia.ValueBool())
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
	if true {
		body, _ = sjson.Set(body, path+"routerLsa", []interface{}{})
		for _, item := range data.RouterLsas {
			itemBody := ""
			if !item.Type.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "adType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "adType.value", item.Type.ValueString())
				}
			}

			if !item.TimeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "time.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "time.value", item.TimeVariable.ValueString())
				}
			} else if !item.Time.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "time.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "time.value", item.Time.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"routerLsa.-1", itemBody)
		}
	}
	if !data.RoutePolicyId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"routePolicy.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"routePolicy.refId.value", data.RoutePolicyId.ValueString())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"area", []interface{}{})
		for _, item := range data.Areas {
			itemBody := ""

			if !item.AreaNumberVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "aNum.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "aNum.value", item.AreaNumberVariable.ValueString())
				}
			} else if !item.AreaNumber.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "aNum.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "aNum.value", item.AreaNumber.ValueInt64())
				}
			}
			if item.AreaType.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "aType.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "aType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "aType.value", item.AreaType.ValueString())
				}
			}

			if !item.NoSummaryVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "noSummary.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "noSummary.value", item.NoSummaryVariable.ValueString())
				}
			} else if item.NoSummary.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "noSummary.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "noSummary.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "noSummary.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "noSummary.value", item.NoSummary.ValueBool())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "interface", []interface{}{})
				for _, childItem := range item.Interfaces {
					itemChildBody := ""

					if !childItem.NameVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "name.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "name.value", childItem.NameVariable.ValueString())
						}
					} else if !childItem.Name.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "name.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "name.value", childItem.Name.ValueString())
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

					if !childItem.DesignatedRouterPriorityVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "priority.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "priority.value", childItem.DesignatedRouterPriorityVariable.ValueString())
						}
					} else if childItem.DesignatedRouterPriority.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "priority.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "priority.value", 1)
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "priority.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "priority.value", childItem.DesignatedRouterPriority.ValueInt64())
						}
					}

					if !childItem.NetworkTypeVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "network.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "network.value", childItem.NetworkTypeVariable.ValueString())
						}
					} else if childItem.NetworkType.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "network.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "network.value", "broadcast")
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "network.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "network.value", childItem.NetworkType.ValueString())
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

					if !childItem.AuthenticationTypeVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "type.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "type.value", childItem.AuthenticationTypeVariable.ValueString())
						}
					} else if childItem.AuthenticationType.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "type.optionType", "default")

						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "type.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "type.value", childItem.AuthenticationType.ValueString())
						}
					}

					if !childItem.MessageDigestKeyIdVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "messageDigestKey.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "messageDigestKey.value", childItem.MessageDigestKeyIdVariable.ValueString())
						}
					} else if childItem.MessageDigestKeyId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "messageDigestKey.optionType", "default")

						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "messageDigestKey.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "messageDigestKey.value", childItem.MessageDigestKeyId.ValueInt64())
						}
					}

					if !childItem.MessageDigestKeyVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "md5.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "md5.value", childItem.MessageDigestKeyVariable.ValueString())
						}
					} else if childItem.MessageDigestKey.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "md5.optionType", "default")

						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "md5.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "md5.value", childItem.MessageDigestKey.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "interface.-1", itemChildBody)
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "range", []interface{}{})
				for _, childItem := range item.Ranges {
					itemChildBody := ""

					if !childItem.IpAddressVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.ipAddress.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.ipAddress.value", childItem.IpAddressVariable.ValueString())
						}
					} else if !childItem.IpAddress.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.ipAddress.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.ipAddress.value", childItem.IpAddress.ValueString())
						}
					}

					if !childItem.SubnetMaskVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.subnetMask.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.subnetMask.value", childItem.SubnetMaskVariable.ValueString())
						}
					} else if !childItem.SubnetMask.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.subnetMask.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.subnetMask.value", childItem.SubnetMask.ValueString())
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
					itemBody, _ = sjson.SetRaw(itemBody, "range.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"area.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceRoutingOSPF) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.RouterId = types.StringNull()
	data.RouterIdVariable = types.StringNull()
	if t := res.Get(path + "routerId.optionType"); t.Exists() {
		va := res.Get(path + "routerId.value")
		if t.String() == "variable" {
			data.RouterIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RouterId = types.StringValue(va.String())
		}
	}
	data.ReferenceBandwidth = types.Int64Null()
	data.ReferenceBandwidthVariable = types.StringNull()
	if t := res.Get(path + "referenceBandwidth.optionType"); t.Exists() {
		va := res.Get(path + "referenceBandwidth.value")
		if t.String() == "variable" {
			data.ReferenceBandwidthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ReferenceBandwidth = types.Int64Value(va.Int())
		}
	}
	data.Rfc1583Compatible = types.BoolNull()
	data.Rfc1583CompatibleVariable = types.StringNull()
	if t := res.Get(path + "rfc1583.optionType"); t.Exists() {
		va := res.Get(path + "rfc1583.value")
		if t.String() == "variable" {
			data.Rfc1583CompatibleVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Rfc1583Compatible = types.BoolValue(va.Bool())
		}
	}
	data.DefaultInformationOriginate = types.BoolNull()

	if t := res.Get(path + "originate.optionType"); t.Exists() {
		va := res.Get(path + "originate.value")
		if t.String() == "global" {
			data.DefaultInformationOriginate = types.BoolValue(va.Bool())
		}
	}
	data.DefaultInformationOriginateAlways = types.BoolNull()
	data.DefaultInformationOriginateAlwaysVariable = types.StringNull()
	if t := res.Get(path + "always.optionType"); t.Exists() {
		va := res.Get(path + "always.value")
		if t.String() == "variable" {
			data.DefaultInformationOriginateAlwaysVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultInformationOriginateAlways = types.BoolValue(va.Bool())
		}
	}
	data.DefaultInformationOriginateMetric = types.Int64Null()
	data.DefaultInformationOriginateMetricVariable = types.StringNull()
	if t := res.Get(path + "metric.optionType"); t.Exists() {
		va := res.Get(path + "metric.value")
		if t.String() == "variable" {
			data.DefaultInformationOriginateMetricVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultInformationOriginateMetric = types.Int64Value(va.Int())
		}
	}
	data.DefaultInformationOriginateMetricType = types.StringNull()
	data.DefaultInformationOriginateMetricTypeVariable = types.StringNull()
	if t := res.Get(path + "metricType.optionType"); t.Exists() {
		va := res.Get(path + "metricType.value")
		if t.String() == "variable" {
			data.DefaultInformationOriginateMetricTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultInformationOriginateMetricType = types.StringValue(va.String())
		}
	}
	data.DistanceExternal = types.Int64Null()
	data.DistanceExternalVariable = types.StringNull()
	if t := res.Get(path + "external.optionType"); t.Exists() {
		va := res.Get(path + "external.value")
		if t.String() == "variable" {
			data.DistanceExternalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceExternal = types.Int64Value(va.Int())
		}
	}
	data.DistanceInterArea = types.Int64Null()
	data.DistanceInterAreaVariable = types.StringNull()
	if t := res.Get(path + "interArea.optionType"); t.Exists() {
		va := res.Get(path + "interArea.value")
		if t.String() == "variable" {
			data.DistanceInterAreaVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceInterArea = types.Int64Value(va.Int())
		}
	}
	data.DistanceIntraArea = types.Int64Null()
	data.DistanceIntraAreaVariable = types.StringNull()
	if t := res.Get(path + "intraArea.optionType"); t.Exists() {
		va := res.Get(path + "intraArea.value")
		if t.String() == "variable" {
			data.DistanceIntraAreaVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceIntraArea = types.Int64Value(va.Int())
		}
	}
	data.SpfCalculationDelay = types.Int64Null()
	data.SpfCalculationDelayVariable = types.StringNull()
	if t := res.Get(path + "delay.optionType"); t.Exists() {
		va := res.Get(path + "delay.value")
		if t.String() == "variable" {
			data.SpfCalculationDelayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SpfCalculationDelay = types.Int64Value(va.Int())
		}
	}
	data.SpfInitialHoldTime = types.Int64Null()
	data.SpfInitialHoldTimeVariable = types.StringNull()
	if t := res.Get(path + "initialHold.optionType"); t.Exists() {
		va := res.Get(path + "initialHold.value")
		if t.String() == "variable" {
			data.SpfInitialHoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SpfInitialHoldTime = types.Int64Value(va.Int())
		}
	}
	data.SpfMaximumHoldTime = types.Int64Null()
	data.SpfMaximumHoldTimeVariable = types.StringNull()
	if t := res.Get(path + "maxHold.optionType"); t.Exists() {
		va := res.Get(path + "maxHold.value")
		if t.String() == "variable" {
			data.SpfMaximumHoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SpfMaximumHoldTime = types.Int64Value(va.Int())
		}
	}
	if value := res.Get(path + "redistribute"); value.Exists() {
		data.Redistributes = make([]ServiceRoutingOSPFRedistributes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingOSPFRedistributes{}
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
			if t := v.Get("dia.optionType"); t.Exists() {
				va := v.Get("dia.value")
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
	if value := res.Get(path + "routerLsa"); value.Exists() {
		data.RouterLsas = make([]ServiceRoutingOSPFRouterLsas, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingOSPFRouterLsas{}
			item.Type = types.StringNull()

			if t := v.Get("adType.optionType"); t.Exists() {
				va := v.Get("adType.value")
				if t.String() == "global" {
					item.Type = types.StringValue(va.String())
				}
			}
			item.Time = types.Int64Null()
			item.TimeVariable = types.StringNull()
			if t := v.Get("time.optionType"); t.Exists() {
				va := v.Get("time.value")
				if t.String() == "variable" {
					item.TimeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Time = types.Int64Value(va.Int())
				}
			}
			data.RouterLsas = append(data.RouterLsas, item)
			return true
		})
	}
	data.RoutePolicyId = types.StringNull()

	if t := res.Get(path + "routePolicy.refId.optionType"); t.Exists() {
		va := res.Get(path + "routePolicy.refId.value")
		if t.String() == "global" {
			data.RoutePolicyId = types.StringValue(va.String())
		}
	}
	if value := res.Get(path + "area"); value.Exists() {
		data.Areas = make([]ServiceRoutingOSPFAreas, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingOSPFAreas{}
			item.AreaNumber = types.Int64Null()
			item.AreaNumberVariable = types.StringNull()
			if t := v.Get("aNum.optionType"); t.Exists() {
				va := v.Get("aNum.value")
				if t.String() == "variable" {
					item.AreaNumberVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AreaNumber = types.Int64Value(va.Int())
				}
			}
			item.AreaType = types.StringNull()

			if t := v.Get("aType.optionType"); t.Exists() {
				va := v.Get("aType.value")
				if t.String() == "global" {
					item.AreaType = types.StringValue(va.String())
				}
			}
			item.NoSummary = types.BoolNull()
			item.NoSummaryVariable = types.StringNull()
			if t := v.Get("noSummary.optionType"); t.Exists() {
				va := v.Get("noSummary.value")
				if t.String() == "variable" {
					item.NoSummaryVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NoSummary = types.BoolValue(va.Bool())
				}
			}
			if cValue := v.Get("interface"); cValue.Exists() {
				item.Interfaces = make([]ServiceRoutingOSPFAreasInterfaces, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceRoutingOSPFAreasInterfaces{}
					cItem.Name = types.StringNull()
					cItem.NameVariable = types.StringNull()
					if t := cv.Get("name.optionType"); t.Exists() {
						va := cv.Get("name.value")
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
					cItem.DesignatedRouterPriority = types.Int64Null()
					cItem.DesignatedRouterPriorityVariable = types.StringNull()
					if t := cv.Get("priority.optionType"); t.Exists() {
						va := cv.Get("priority.value")
						if t.String() == "variable" {
							cItem.DesignatedRouterPriorityVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.DesignatedRouterPriority = types.Int64Value(va.Int())
						}
					}
					cItem.NetworkType = types.StringNull()
					cItem.NetworkTypeVariable = types.StringNull()
					if t := cv.Get("network.optionType"); t.Exists() {
						va := cv.Get("network.value")
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
					cItem.AuthenticationTypeVariable = types.StringNull()
					if t := cv.Get("type.optionType"); t.Exists() {
						va := cv.Get("type.value")
						if t.String() == "variable" {
							cItem.AuthenticationTypeVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.AuthenticationType = types.StringValue(va.String())
						}
					}
					cItem.MessageDigestKeyId = types.Int64Null()
					cItem.MessageDigestKeyIdVariable = types.StringNull()
					if t := cv.Get("messageDigestKey.optionType"); t.Exists() {
						va := cv.Get("messageDigestKey.value")
						if t.String() == "variable" {
							cItem.MessageDigestKeyIdVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.MessageDigestKeyId = types.Int64Value(va.Int())
						}
					}
					item.Interfaces = append(item.Interfaces, cItem)
					return true
				})
			}
			if cValue := v.Get("range"); cValue.Exists() {
				item.Ranges = make([]ServiceRoutingOSPFAreasRanges, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceRoutingOSPFAreasRanges{}
					cItem.IpAddress = types.StringNull()
					cItem.IpAddressVariable = types.StringNull()
					if t := cv.Get("address.ipAddress.optionType"); t.Exists() {
						va := cv.Get("address.ipAddress.value")
						if t.String() == "variable" {
							cItem.IpAddressVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.IpAddress = types.StringValue(va.String())
						}
					}
					cItem.SubnetMask = types.StringNull()
					cItem.SubnetMaskVariable = types.StringNull()
					if t := cv.Get("address.subnetMask.optionType"); t.Exists() {
						va := cv.Get("address.subnetMask.value")
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
func (data *ServiceRoutingOSPF) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.RouterId = types.StringNull()
	data.RouterIdVariable = types.StringNull()
	if t := res.Get(path + "routerId.optionType"); t.Exists() {
		va := res.Get(path + "routerId.value")
		if t.String() == "variable" {
			data.RouterIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RouterId = types.StringValue(va.String())
		}
	}
	data.ReferenceBandwidth = types.Int64Null()
	data.ReferenceBandwidthVariable = types.StringNull()
	if t := res.Get(path + "referenceBandwidth.optionType"); t.Exists() {
		va := res.Get(path + "referenceBandwidth.value")
		if t.String() == "variable" {
			data.ReferenceBandwidthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ReferenceBandwidth = types.Int64Value(va.Int())
		}
	}
	data.Rfc1583Compatible = types.BoolNull()
	data.Rfc1583CompatibleVariable = types.StringNull()
	if t := res.Get(path + "rfc1583.optionType"); t.Exists() {
		va := res.Get(path + "rfc1583.value")
		if t.String() == "variable" {
			data.Rfc1583CompatibleVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Rfc1583Compatible = types.BoolValue(va.Bool())
		}
	}
	data.DefaultInformationOriginate = types.BoolNull()

	if t := res.Get(path + "originate.optionType"); t.Exists() {
		va := res.Get(path + "originate.value")
		if t.String() == "global" {
			data.DefaultInformationOriginate = types.BoolValue(va.Bool())
		}
	}
	data.DefaultInformationOriginateAlways = types.BoolNull()
	data.DefaultInformationOriginateAlwaysVariable = types.StringNull()
	if t := res.Get(path + "always.optionType"); t.Exists() {
		va := res.Get(path + "always.value")
		if t.String() == "variable" {
			data.DefaultInformationOriginateAlwaysVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultInformationOriginateAlways = types.BoolValue(va.Bool())
		}
	}
	data.DefaultInformationOriginateMetric = types.Int64Null()
	data.DefaultInformationOriginateMetricVariable = types.StringNull()
	if t := res.Get(path + "metric.optionType"); t.Exists() {
		va := res.Get(path + "metric.value")
		if t.String() == "variable" {
			data.DefaultInformationOriginateMetricVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultInformationOriginateMetric = types.Int64Value(va.Int())
		}
	}
	data.DefaultInformationOriginateMetricType = types.StringNull()
	data.DefaultInformationOriginateMetricTypeVariable = types.StringNull()
	if t := res.Get(path + "metricType.optionType"); t.Exists() {
		va := res.Get(path + "metricType.value")
		if t.String() == "variable" {
			data.DefaultInformationOriginateMetricTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultInformationOriginateMetricType = types.StringValue(va.String())
		}
	}
	data.DistanceExternal = types.Int64Null()
	data.DistanceExternalVariable = types.StringNull()
	if t := res.Get(path + "external.optionType"); t.Exists() {
		va := res.Get(path + "external.value")
		if t.String() == "variable" {
			data.DistanceExternalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceExternal = types.Int64Value(va.Int())
		}
	}
	data.DistanceInterArea = types.Int64Null()
	data.DistanceInterAreaVariable = types.StringNull()
	if t := res.Get(path + "interArea.optionType"); t.Exists() {
		va := res.Get(path + "interArea.value")
		if t.String() == "variable" {
			data.DistanceInterAreaVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceInterArea = types.Int64Value(va.Int())
		}
	}
	data.DistanceIntraArea = types.Int64Null()
	data.DistanceIntraAreaVariable = types.StringNull()
	if t := res.Get(path + "intraArea.optionType"); t.Exists() {
		va := res.Get(path + "intraArea.value")
		if t.String() == "variable" {
			data.DistanceIntraAreaVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DistanceIntraArea = types.Int64Value(va.Int())
		}
	}
	data.SpfCalculationDelay = types.Int64Null()
	data.SpfCalculationDelayVariable = types.StringNull()
	if t := res.Get(path + "delay.optionType"); t.Exists() {
		va := res.Get(path + "delay.value")
		if t.String() == "variable" {
			data.SpfCalculationDelayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SpfCalculationDelay = types.Int64Value(va.Int())
		}
	}
	data.SpfInitialHoldTime = types.Int64Null()
	data.SpfInitialHoldTimeVariable = types.StringNull()
	if t := res.Get(path + "initialHold.optionType"); t.Exists() {
		va := res.Get(path + "initialHold.value")
		if t.String() == "variable" {
			data.SpfInitialHoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SpfInitialHoldTime = types.Int64Value(va.Int())
		}
	}
	data.SpfMaximumHoldTime = types.Int64Null()
	data.SpfMaximumHoldTimeVariable = types.StringNull()
	if t := res.Get(path + "maxHold.optionType"); t.Exists() {
		va := res.Get(path + "maxHold.value")
		if t.String() == "variable" {
			data.SpfMaximumHoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SpfMaximumHoldTime = types.Int64Value(va.Int())
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
		if t := r.Get("dia.optionType"); t.Exists() {
			va := r.Get("dia.value")
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
	for i := range data.RouterLsas {
		keys := [...]string{"adType"}
		keyValues := [...]string{data.RouterLsas[i].Type.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "routerLsa").ForEach(
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
		data.RouterLsas[i].Type = types.StringNull()

		if t := r.Get("adType.optionType"); t.Exists() {
			va := r.Get("adType.value")
			if t.String() == "global" {
				data.RouterLsas[i].Type = types.StringValue(va.String())
			}
		}
		data.RouterLsas[i].Time = types.Int64Null()
		data.RouterLsas[i].TimeVariable = types.StringNull()
		if t := r.Get("time.optionType"); t.Exists() {
			va := r.Get("time.value")
			if t.String() == "variable" {
				data.RouterLsas[i].TimeVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.RouterLsas[i].Time = types.Int64Value(va.Int())
			}
		}
	}
	data.RoutePolicyId = types.StringNull()

	if t := res.Get(path + "routePolicy.refId.optionType"); t.Exists() {
		va := res.Get(path + "routePolicy.refId.value")
		if t.String() == "global" {
			data.RoutePolicyId = types.StringValue(va.String())
		}
	}
	for i := range data.Areas {
		keys := [...]string{"aNum"}
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
		if t := r.Get("aNum.optionType"); t.Exists() {
			va := r.Get("aNum.value")
			if t.String() == "variable" {
				data.Areas[i].AreaNumberVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Areas[i].AreaNumber = types.Int64Value(va.Int())
			}
		}
		data.Areas[i].AreaType = types.StringNull()

		if t := r.Get("aType.optionType"); t.Exists() {
			va := r.Get("aType.value")
			if t.String() == "global" {
				data.Areas[i].AreaType = types.StringValue(va.String())
			}
		}
		data.Areas[i].NoSummary = types.BoolNull()
		data.Areas[i].NoSummaryVariable = types.StringNull()
		if t := r.Get("noSummary.optionType"); t.Exists() {
			va := r.Get("noSummary.value")
			if t.String() == "variable" {
				data.Areas[i].NoSummaryVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Areas[i].NoSummary = types.BoolValue(va.Bool())
			}
		}
		for ci := range data.Areas[i].Interfaces {
			keys := [...]string{"name"}
			keyValues := [...]string{data.Areas[i].Interfaces[ci].Name.ValueString()}
			keyValuesVariables := [...]string{data.Areas[i].Interfaces[ci].NameVariable.ValueString()}

			var cr gjson.Result
			r.Get("interface").ForEach(
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
			if t := cr.Get("name.optionType"); t.Exists() {
				va := cr.Get("name.value")
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
			data.Areas[i].Interfaces[ci].DesignatedRouterPriority = types.Int64Null()
			data.Areas[i].Interfaces[ci].DesignatedRouterPriorityVariable = types.StringNull()
			if t := cr.Get("priority.optionType"); t.Exists() {
				va := cr.Get("priority.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].DesignatedRouterPriorityVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].DesignatedRouterPriority = types.Int64Value(va.Int())
				}
			}
			data.Areas[i].Interfaces[ci].NetworkType = types.StringNull()
			data.Areas[i].Interfaces[ci].NetworkTypeVariable = types.StringNull()
			if t := cr.Get("network.optionType"); t.Exists() {
				va := cr.Get("network.value")
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
			data.Areas[i].Interfaces[ci].AuthenticationTypeVariable = types.StringNull()
			if t := cr.Get("type.optionType"); t.Exists() {
				va := cr.Get("type.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].AuthenticationTypeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].AuthenticationType = types.StringValue(va.String())
				}
			}
			data.Areas[i].Interfaces[ci].MessageDigestKeyId = types.Int64Null()
			data.Areas[i].Interfaces[ci].MessageDigestKeyIdVariable = types.StringNull()
			if t := cr.Get("messageDigestKey.optionType"); t.Exists() {
				va := cr.Get("messageDigestKey.value")
				if t.String() == "variable" {
					data.Areas[i].Interfaces[ci].MessageDigestKeyIdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Interfaces[ci].MessageDigestKeyId = types.Int64Value(va.Int())
				}
			}
		}
		for ci := range data.Areas[i].Ranges {
			keys := [...]string{"address.ipAddress"}
			keyValues := [...]string{data.Areas[i].Ranges[ci].IpAddress.ValueString()}
			keyValuesVariables := [...]string{data.Areas[i].Ranges[ci].IpAddressVariable.ValueString()}

			var cr gjson.Result
			r.Get("range").ForEach(
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
			if t := cr.Get("address.ipAddress.optionType"); t.Exists() {
				va := cr.Get("address.ipAddress.value")
				if t.String() == "variable" {
					data.Areas[i].Ranges[ci].IpAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Areas[i].Ranges[ci].IpAddress = types.StringValue(va.String())
				}
			}
			data.Areas[i].Ranges[ci].SubnetMask = types.StringNull()
			data.Areas[i].Ranges[ci].SubnetMaskVariable = types.StringNull()
			if t := cr.Get("address.subnetMask.optionType"); t.Exists() {
				va := cr.Get("address.subnetMask.value")
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
