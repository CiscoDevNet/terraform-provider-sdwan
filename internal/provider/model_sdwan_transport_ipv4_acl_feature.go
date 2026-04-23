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
type TransportIPv4ACL struct {
	Id               types.String                `tfsdk:"id"`
	Version          types.Int64                 `tfsdk:"version"`
	Name             types.String                `tfsdk:"name"`
	Description      types.String                `tfsdk:"description"`
	FeatureProfileId types.String                `tfsdk:"feature_profile_id"`
	DefaultAction    types.String                `tfsdk:"default_action"`
	Sequences        []TransportIPv4ACLSequences `tfsdk:"sequences"`
}

type TransportIPv4ACLSequences struct {
	SequenceId   types.Int64                             `tfsdk:"sequence_id"`
	SequenceName types.String                            `tfsdk:"sequence_name"`
	BaseAction   types.String                            `tfsdk:"base_action"`
	MatchEntries []TransportIPv4ACLSequencesMatchEntries `tfsdk:"match_entries"`
	Actions      []TransportIPv4ACLSequencesActions      `tfsdk:"actions"`
}

type TransportIPv4ACLSequencesMatchEntries struct {
	Dscps                         types.Set                                               `tfsdk:"dscps"`
	PacketLength                  types.String                                            `tfsdk:"packet_length"`
	Protocols                     types.Set                                               `tfsdk:"protocols"`
	IcmpMessages                  types.Set                                               `tfsdk:"icmp_messages"`
	SourceDataPrefixListId        types.String                                            `tfsdk:"source_data_prefix_list_id"`
	SourceDataPrefix              types.String                                            `tfsdk:"source_data_prefix"`
	SourceDataPrefixVariable      types.String                                            `tfsdk:"source_data_prefix_variable"`
	SourcePorts                   []TransportIPv4ACLSequencesMatchEntriesSourcePorts      `tfsdk:"source_ports"`
	DestinationDataPrefixListId   types.String                                            `tfsdk:"destination_data_prefix_list_id"`
	DestinationDataPrefix         types.String                                            `tfsdk:"destination_data_prefix"`
	DestinationDataPrefixVariable types.String                                            `tfsdk:"destination_data_prefix_variable"`
	DestinationPorts              []TransportIPv4ACLSequencesMatchEntriesDestinationPorts `tfsdk:"destination_ports"`
	TcpState                      types.String                                            `tfsdk:"tcp_state"`
}
type TransportIPv4ACLSequencesActions struct {
	AcceptSetDscp                         types.Int64  `tfsdk:"accept_set_dscp"`
	AcceptCounterName                     types.String `tfsdk:"accept_counter_name"`
	AcceptLog                             types.Bool   `tfsdk:"accept_log"`
	AcceptSetNextHop                      types.String `tfsdk:"accept_set_next_hop"`
	AcceptSetServiceChainName             types.String `tfsdk:"accept_set_service_chain_name"`
	AcceptSetServiceChainNameVariable     types.String `tfsdk:"accept_set_service_chain_name_variable"`
	AcceptSetServiceChainVpn              types.Int64  `tfsdk:"accept_set_service_chain_vpn"`
	AcceptSetServiceChainVpnVariable      types.String `tfsdk:"accept_set_service_chain_vpn_variable"`
	AcceptSetServiceChainFallback         types.Bool   `tfsdk:"accept_set_service_chain_fallback"`
	AcceptSetServiceChainFallbackVariable types.String `tfsdk:"accept_set_service_chain_fallback_variable"`
	AcceptMirrorListId                    types.String `tfsdk:"accept_mirror_list_id"`
	AcceptPolicerId                       types.String `tfsdk:"accept_policer_id"`
	DropCounterName                       types.String `tfsdk:"drop_counter_name"`
	DropLog                               types.Bool   `tfsdk:"drop_log"`
}

type TransportIPv4ACLSequencesMatchEntriesSourcePorts struct {
	Port types.String `tfsdk:"port"`
}
type TransportIPv4ACLSequencesMatchEntriesDestinationPorts struct {
	Port types.String `tfsdk:"port"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TransportIPv4ACL) getModel() string {
	return "transport_ipv4_acl"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportIPv4ACL) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/ipv4-acl", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportIPv4ACL) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if data.DefaultAction.IsNull() || data.DefaultAction.ValueString() == "drop" {
		if true {
			body, _ = sjson.Set(body, path+"defaultAction.optionType", "default")
			body, _ = sjson.Set(body, path+"defaultAction.value", "drop")
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
			if !item.SequenceId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceId.value", item.SequenceId.ValueInt64())
				}
			}
			if !item.SequenceName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceName.value", item.SequenceName.ValueString())
				}
			}
			if !item.BaseAction.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "baseAction.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "baseAction.value", item.BaseAction.ValueString())
				}
			}
			if true {

				for _, childItem := range item.MatchEntries {
					itemChildBody := ""
					if !childItem.Dscps.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "dscp.optionType", "global")
							var values []int64
							childItem.Dscps.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "dscp.value", values)
						}
					}
					if !childItem.PacketLength.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "packetLength.optionType", "global")
							if numValue, err := strconv.Atoi(childItem.PacketLength.ValueString()); err != nil {
								itemChildBody, _ = sjson.Set(itemChildBody, "packetLength.value", childItem.PacketLength.ValueString())
							} else {
								itemChildBody, _ = sjson.Set(itemChildBody, "packetLength.value", numValue)
							}
						}
					}
					if !childItem.Protocols.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.optionType", "global")
							var values []int64
							childItem.Protocols.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.value", values)
						}
					}
					if !childItem.IcmpMessages.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "icmpMsg.optionType", "global")
							var values []string
							childItem.IcmpMessages.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "icmpMsg.value", values)
						}
					}
					if !childItem.SourceDataPrefixListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataPrefix.sourceDataPrefixList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataPrefix.sourceDataPrefixList.refId.value", childItem.SourceDataPrefixListId.ValueString())
						}
					}

					if !childItem.SourceDataPrefixVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataPrefix.sourceIpPrefix.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataPrefix.sourceIpPrefix.value", childItem.SourceDataPrefixVariable.ValueString())
						}
					} else if !childItem.SourceDataPrefix.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataPrefix.sourceIpPrefix.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataPrefix.sourceIpPrefix.value", childItem.SourceDataPrefix.ValueString())
						}
					}
					if true {

						for _, childChildItem := range childItem.SourcePorts {
							itemChildChildBody := ""
							if !childChildItem.Port.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "sourcePort.optionType", "global")
									if numValue, err := strconv.Atoi(childChildItem.Port.ValueString()); err != nil {
										itemChildChildBody, _ = sjson.Set(itemChildChildBody, "sourcePort.value", childChildItem.Port.ValueString())
									} else {
										itemChildChildBody, _ = sjson.Set(itemChildChildBody, "sourcePort.value", numValue)
									}
								}
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "sourcePorts.-1", itemChildChildBody)
						}
					}
					if !childItem.DestinationDataPrefixListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationDataPrefix.destinationDataPrefixList.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationDataPrefix.destinationDataPrefixList.refId.value", childItem.DestinationDataPrefixListId.ValueString())
						}
					}

					if !childItem.DestinationDataPrefixVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationDataPrefix.destinationIpPrefix.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationDataPrefix.destinationIpPrefix.value", childItem.DestinationDataPrefixVariable.ValueString())
						}
					} else if !childItem.DestinationDataPrefix.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationDataPrefix.destinationIpPrefix.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationDataPrefix.destinationIpPrefix.value", childItem.DestinationDataPrefix.ValueString())
						}
					}
					if true {

						for _, childChildItem := range childItem.DestinationPorts {
							itemChildChildBody := ""
							if !childChildItem.Port.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "destinationPort.optionType", "global")
									if numValue, err := strconv.Atoi(childChildItem.Port.ValueString()); err != nil {
										itemChildChildBody, _ = sjson.Set(itemChildChildBody, "destinationPort.value", childChildItem.Port.ValueString())
									} else {
										itemChildChildBody, _ = sjson.Set(itemChildChildBody, "destinationPort.value", numValue)
									}
								}
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "destinationPorts.-1", itemChildChildBody)
						}
					}
					if !childItem.TcpState.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "tcp.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "tcp.value", childItem.TcpState.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "matchEntries.-1", itemChildBody)
				}
			}
			if true {

				for _, childItem := range item.Actions {
					itemChildBody := ""
					if !childItem.AcceptSetDscp.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setDscp.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setDscp.value", childItem.AcceptSetDscp.ValueInt64())
						}
					}
					if !childItem.AcceptCounterName.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.counterName.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.counterName.value", childItem.AcceptCounterName.ValueString())
						}
					}
					if !childItem.AcceptLog.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.log.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.log.value", childItem.AcceptLog.ValueBool())
						}
					}
					if !childItem.AcceptSetNextHop.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setNextHop.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setNextHop.value", childItem.AcceptSetNextHop.ValueString())
						}
					}

					if !childItem.AcceptSetServiceChainNameVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setServiceChain.serviceChainNumber.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setServiceChain.serviceChainNumber.value", childItem.AcceptSetServiceChainNameVariable.ValueString())
						}
					} else if !childItem.AcceptSetServiceChainName.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setServiceChain.serviceChainNumber.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setServiceChain.serviceChainNumber.value", childItem.AcceptSetServiceChainName.ValueString())
						}
					}

					if !childItem.AcceptSetServiceChainVpnVariable.IsNull() {
						if true && (!(childItem.AcceptSetServiceChainName.IsNull()) || !(childItem.AcceptSetServiceChainNameVariable.IsNull())) {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setServiceChain.vpn.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setServiceChain.vpn.value", childItem.AcceptSetServiceChainVpnVariable.ValueString())
						}
					} else if !childItem.AcceptSetServiceChainVpn.IsNull() {
						if true && (!(childItem.AcceptSetServiceChainName.IsNull()) || !(childItem.AcceptSetServiceChainNameVariable.IsNull())) {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setServiceChain.vpn.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setServiceChain.vpn.value", childItem.AcceptSetServiceChainVpn.ValueInt64())
						}
					}

					if !childItem.AcceptSetServiceChainFallbackVariable.IsNull() {
						if true && (!(childItem.AcceptSetServiceChainName.IsNull()) || !(childItem.AcceptSetServiceChainNameVariable.IsNull())) {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setServiceChain.fallback.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setServiceChain.fallback.value", childItem.AcceptSetServiceChainFallbackVariable.ValueString())
						}
					} else if childItem.AcceptSetServiceChainFallback.IsNull() {
						if true && (!(childItem.AcceptSetServiceChainName.IsNull()) || !(childItem.AcceptSetServiceChainNameVariable.IsNull())) {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setServiceChain.fallback.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setServiceChain.fallback.value", false)
						}
					} else {
						if true && (!(childItem.AcceptSetServiceChainName.IsNull()) || !(childItem.AcceptSetServiceChainNameVariable.IsNull())) {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setServiceChain.fallback.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.setServiceChain.fallback.value", childItem.AcceptSetServiceChainFallback.ValueBool())
						}
					}
					if !childItem.AcceptMirrorListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.mirror.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.mirror.refId.value", childItem.AcceptMirrorListId.ValueString())
						}
					}
					if !childItem.AcceptPolicerId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.policer.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "accept.policer.refId.value", childItem.AcceptPolicerId.ValueString())
						}
					}
					if !childItem.DropCounterName.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "drop.counterName.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "drop.counterName.value", childItem.DropCounterName.ValueString())
						}
					}
					if !childItem.DropLog.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "drop.log.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "drop.log.value", childItem.DropLog.ValueBool())
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
func (data *TransportIPv4ACL) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
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
		if t.String() == "global" || (t.String() == "default" && (tempDefaultAction.ValueString() == "drop" || (fullRead && tempDefaultAction.ValueString() == ""))) {
			data.DefaultAction = types.StringValue(va.String())
		}
	}
	oldSequences := data.Sequences
	if value := res.Get(path + "sequences"); value.Exists() && len(value.Array()) > 0 {
		data.Sequences = make([]TransportIPv4ACLSequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportIPv4ACLSequences{}
			item.SequenceId = types.Int64Null()

			if t := v.Get("sequenceId.optionType"); t.Exists() {
				va := v.Get("sequenceId.value")
				if t.String() == "global" {
					item.SequenceId = types.Int64Value(va.Int())
				}
			}
			item.SequenceName = types.StringNull()

			if t := v.Get("sequenceName.optionType"); t.Exists() {
				va := v.Get("sequenceName.value")
				if t.String() == "global" {
					item.SequenceName = types.StringValue(va.String())
				}
			}
			item.BaseAction = types.StringNull()

			if t := v.Get("baseAction.optionType"); t.Exists() {
				va := v.Get("baseAction.value")
				if t.String() == "global" {
					item.BaseAction = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("matchEntries"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.MatchEntries = make([]TransportIPv4ACLSequencesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TransportIPv4ACLSequencesMatchEntries{}
					cItem.Dscps = types.SetNull(types.Int64Type)

					if t := cv.Get("dscp.optionType"); t.Exists() {
						va := cv.Get("dscp.value")
						if t.String() == "global" {
							cItem.Dscps = helpers.GetInt64Set(va.Array())
						}
					}
					cItem.PacketLength = types.StringNull()

					if t := cv.Get("packetLength.optionType"); t.Exists() {
						va := cv.Get("packetLength.value")
						if t.String() == "global" {
							cItem.PacketLength = types.StringValue(va.String())
						}
					}
					cItem.Protocols = types.SetNull(types.Int64Type)

					if t := cv.Get("protocol.optionType"); t.Exists() {
						va := cv.Get("protocol.value")
						if t.String() == "global" {
							cItem.Protocols = helpers.GetInt64Set(va.Array())
						}
					}
					cItem.IcmpMessages = types.SetNull(types.StringType)

					if t := cv.Get("icmpMsg.optionType"); t.Exists() {
						va := cv.Get("icmpMsg.value")
						if t.String() == "global" {
							cItem.IcmpMessages = helpers.GetStringSet(va.Array())
						}
					}
					cItem.SourceDataPrefixListId = types.StringNull()

					if t := cv.Get("sourceDataPrefix.sourceDataPrefixList.refId.optionType"); t.Exists() {
						va := cv.Get("sourceDataPrefix.sourceDataPrefixList.refId.value")
						if t.String() == "global" {
							cItem.SourceDataPrefixListId = types.StringValue(va.String())
						}
					}
					cItem.SourceDataPrefix = types.StringNull()
					cItem.SourceDataPrefixVariable = types.StringNull()
					if t := cv.Get("sourceDataPrefix.sourceIpPrefix.optionType"); t.Exists() {
						va := cv.Get("sourceDataPrefix.sourceIpPrefix.value")
						if t.String() == "variable" {
							cItem.SourceDataPrefixVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.SourceDataPrefix = types.StringValue(va.String())
						}
					}
					if ccValue := cv.Get("sourcePorts"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.SourcePorts = make([]TransportIPv4ACLSequencesMatchEntriesSourcePorts, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := TransportIPv4ACLSequencesMatchEntriesSourcePorts{}
							ccItem.Port = types.StringNull()

							if t := ccv.Get("sourcePort.optionType"); t.Exists() {
								va := ccv.Get("sourcePort.value")
								if t.String() == "global" {
									ccItem.Port = types.StringValue(va.String())
								}
							}
							cItem.SourcePorts = append(cItem.SourcePorts, ccItem)
							return true
						})
					}
					cItem.DestinationDataPrefixListId = types.StringNull()

					if t := cv.Get("destinationDataPrefix.destinationDataPrefixList.refId.optionType"); t.Exists() {
						va := cv.Get("destinationDataPrefix.destinationDataPrefixList.refId.value")
						if t.String() == "global" {
							cItem.DestinationDataPrefixListId = types.StringValue(va.String())
						}
					}
					cItem.DestinationDataPrefix = types.StringNull()
					cItem.DestinationDataPrefixVariable = types.StringNull()
					if t := cv.Get("destinationDataPrefix.destinationIpPrefix.optionType"); t.Exists() {
						va := cv.Get("destinationDataPrefix.destinationIpPrefix.value")
						if t.String() == "variable" {
							cItem.DestinationDataPrefixVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.DestinationDataPrefix = types.StringValue(va.String())
						}
					}
					if ccValue := cv.Get("destinationPorts"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.DestinationPorts = make([]TransportIPv4ACLSequencesMatchEntriesDestinationPorts, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := TransportIPv4ACLSequencesMatchEntriesDestinationPorts{}
							ccItem.Port = types.StringNull()

							if t := ccv.Get("destinationPort.optionType"); t.Exists() {
								va := ccv.Get("destinationPort.value")
								if t.String() == "global" {
									ccItem.Port = types.StringValue(va.String())
								}
							}
							cItem.DestinationPorts = append(cItem.DestinationPorts, ccItem)
							return true
						})
					}
					cItem.TcpState = types.StringNull()

					if t := cv.Get("tcp.optionType"); t.Exists() {
						va := cv.Get("tcp.value")
						if t.String() == "global" {
							cItem.TcpState = types.StringValue(va.String())
						}
					}
					item.MatchEntries = append(item.MatchEntries, cItem)
					return true
				})
			}
			if cValue := v.Get("actions"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Actions = make([]TransportIPv4ACLSequencesActions, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TransportIPv4ACLSequencesActions{}
					cItem.AcceptSetDscp = types.Int64Null()

					if t := cv.Get("accept.setDscp.optionType"); t.Exists() {
						va := cv.Get("accept.setDscp.value")
						if t.String() == "global" {
							cItem.AcceptSetDscp = types.Int64Value(va.Int())
						}
					}
					cItem.AcceptCounterName = types.StringNull()

					if t := cv.Get("accept.counterName.optionType"); t.Exists() {
						va := cv.Get("accept.counterName.value")
						if t.String() == "global" {
							cItem.AcceptCounterName = types.StringValue(va.String())
						}
					}
					cItem.AcceptLog = types.BoolNull()

					if t := cv.Get("accept.log.optionType"); t.Exists() {
						va := cv.Get("accept.log.value")
						if t.String() == "global" {
							cItem.AcceptLog = types.BoolValue(va.Bool())
						}
					}
					cItem.AcceptSetNextHop = types.StringNull()

					if t := cv.Get("accept.setNextHop.optionType"); t.Exists() {
						va := cv.Get("accept.setNextHop.value")
						if t.String() == "global" {
							cItem.AcceptSetNextHop = types.StringValue(va.String())
						}
					}
					cItem.AcceptSetServiceChainName = types.StringNull()
					cItem.AcceptSetServiceChainNameVariable = types.StringNull()
					if t := cv.Get("accept.setServiceChain.serviceChainNumber.optionType"); t.Exists() {
						va := cv.Get("accept.setServiceChain.serviceChainNumber.value")
						if t.String() == "variable" {
							cItem.AcceptSetServiceChainNameVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.AcceptSetServiceChainName = types.StringValue(va.String())
						}
					}
					cItem.AcceptSetServiceChainVpn = types.Int64Null()
					cItem.AcceptSetServiceChainVpnVariable = types.StringNull()
					if t := cv.Get("accept.setServiceChain.vpn.optionType"); t.Exists() {
						va := cv.Get("accept.setServiceChain.vpn.value")
						if t.String() == "variable" {
							cItem.AcceptSetServiceChainVpnVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.AcceptSetServiceChainVpn = types.Int64Value(va.Int())
						}
					}
					cItem.AcceptSetServiceChainFallback = types.BoolNull()
					cItem.AcceptSetServiceChainFallbackVariable = types.StringNull()
					if t := cv.Get("accept.setServiceChain.fallback.optionType"); t.Exists() {
						va := cv.Get("accept.setServiceChain.fallback.value")
						if t.String() == "variable" {
							cItem.AcceptSetServiceChainFallbackVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.AcceptSetServiceChainFallback = types.BoolValue(va.Bool())
						}
					}
					cItem.AcceptMirrorListId = types.StringNull()

					if t := cv.Get("accept.mirror.refId.optionType"); t.Exists() {
						va := cv.Get("accept.mirror.refId.value")
						if t.String() == "global" {
							cItem.AcceptMirrorListId = types.StringValue(va.String())
						}
					}
					cItem.AcceptPolicerId = types.StringNull()

					if t := cv.Get("accept.policer.refId.optionType"); t.Exists() {
						va := cv.Get("accept.policer.refId.value")
						if t.String() == "global" {
							cItem.AcceptPolicerId = types.StringValue(va.String())
						}
					}
					cItem.DropCounterName = types.StringNull()

					if t := cv.Get("drop.counterName.optionType"); t.Exists() {
						va := cv.Get("drop.counterName.value")
						if t.String() == "global" {
							cItem.DropCounterName = types.StringValue(va.String())
						}
					}
					cItem.DropLog = types.BoolNull()

					if t := cv.Get("drop.log.optionType"); t.Exists() {
						va := cv.Get("drop.log.value")
						if t.String() == "global" {
							cItem.DropLog = types.BoolValue(va.Bool())
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
		resultSequences := make([]TransportIPv4ACLSequences, 0, len(data.Sequences))
		matchedSequences := make([]bool, len(data.Sequences))
		for _, oldItem := range oldSequences {
			for ni := range data.Sequences {
				if matchedSequences[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.SequenceId.ValueInt64() != data.Sequences[ni].SequenceId.ValueInt64() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedSequences[ni] = true
					{
						resultC := make([]TransportIPv4ACLSequencesMatchEntries, 0, len(data.Sequences[ni].MatchEntries))
						matchedC := make([]bool, len(data.Sequences[ni].MatchEntries))
						for _, oldCItem := range oldItem.MatchEntries {
							for nci := range data.Sequences[ni].MatchEntries {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.Dscps).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].Dscps).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.PacketLength.ValueString() != data.Sequences[ni].MatchEntries[nci].PacketLength.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.Protocols).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].Protocols).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.IcmpMessages).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].IcmpMessages).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.SourceDataPrefixListId.ValueString() != data.Sequences[ni].MatchEntries[nci].SourceDataPrefixListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC && (oldCItem.SourceDataPrefixVariable.ValueString() != "" || data.Sequences[ni].MatchEntries[nci].SourceDataPrefixVariable.ValueString() != "") {
									if oldCItem.SourceDataPrefixVariable.ValueString() != data.Sequences[ni].MatchEntries[nci].SourceDataPrefixVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.SourceDataPrefix.ValueString() != data.Sequences[ni].MatchEntries[nci].SourceDataPrefix.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.DestinationDataPrefixListId.ValueString() != data.Sequences[ni].MatchEntries[nci].DestinationDataPrefixListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC && (oldCItem.DestinationDataPrefixVariable.ValueString() != "" || data.Sequences[ni].MatchEntries[nci].DestinationDataPrefixVariable.ValueString() != "") {
									if oldCItem.DestinationDataPrefixVariable.ValueString() != data.Sequences[ni].MatchEntries[nci].DestinationDataPrefixVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.DestinationDataPrefix.ValueString() != data.Sequences[ni].MatchEntries[nci].DestinationDataPrefix.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.TcpState.ValueString() != data.Sequences[ni].MatchEntries[nci].TcpState.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									{
										resultCC := make([]TransportIPv4ACLSequencesMatchEntriesSourcePorts, 0, len(data.Sequences[ni].MatchEntries[nci].SourcePorts))
										matchedCC := make([]bool, len(data.Sequences[ni].MatchEntries[nci].SourcePorts))
										for _, oldCCItem := range oldCItem.SourcePorts {
											for ncci := range data.Sequences[ni].MatchEntries[nci].SourcePorts {
												if matchedCC[ncci] {
													continue
												}
												keyMatchCC := true
												if keyMatchCC {
													if oldCCItem.Port.ValueString() != data.Sequences[ni].MatchEntries[nci].SourcePorts[ncci].Port.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													matchedCC[ncci] = true
													resultCC = append(resultCC, data.Sequences[ni].MatchEntries[nci].SourcePorts[ncci])
													break
												}
											}
										}
										for ncci := range data.Sequences[ni].MatchEntries[nci].SourcePorts {
											if !matchedCC[ncci] {
												resultCC = append(resultCC, data.Sequences[ni].MatchEntries[nci].SourcePorts[ncci])
											}
										}
										data.Sequences[ni].MatchEntries[nci].SourcePorts = resultCC
									}
									{
										resultCC := make([]TransportIPv4ACLSequencesMatchEntriesDestinationPorts, 0, len(data.Sequences[ni].MatchEntries[nci].DestinationPorts))
										matchedCC := make([]bool, len(data.Sequences[ni].MatchEntries[nci].DestinationPorts))
										for _, oldCCItem := range oldCItem.DestinationPorts {
											for ncci := range data.Sequences[ni].MatchEntries[nci].DestinationPorts {
												if matchedCC[ncci] {
													continue
												}
												keyMatchCC := true
												if keyMatchCC {
													if oldCCItem.Port.ValueString() != data.Sequences[ni].MatchEntries[nci].DestinationPorts[ncci].Port.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													matchedCC[ncci] = true
													resultCC = append(resultCC, data.Sequences[ni].MatchEntries[nci].DestinationPorts[ncci])
													break
												}
											}
										}
										for ncci := range data.Sequences[ni].MatchEntries[nci].DestinationPorts {
											if !matchedCC[ncci] {
												resultCC = append(resultCC, data.Sequences[ni].MatchEntries[nci].DestinationPorts[ncci])
											}
										}
										data.Sequences[ni].MatchEntries[nci].DestinationPorts = resultCC
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
						resultC := make([]TransportIPv4ACLSequencesActions, 0, len(data.Sequences[ni].Actions))
						matchedC := make([]bool, len(data.Sequences[ni].Actions))
						for _, oldCItem := range oldItem.Actions {
							for nci := range data.Sequences[ni].Actions {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if oldCItem.AcceptSetDscp.ValueInt64() != data.Sequences[ni].Actions[nci].AcceptSetDscp.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.AcceptCounterName.ValueString() != data.Sequences[ni].Actions[nci].AcceptCounterName.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.AcceptLog.ValueBool() != data.Sequences[ni].Actions[nci].AcceptLog.ValueBool() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.AcceptSetNextHop.ValueString() != data.Sequences[ni].Actions[nci].AcceptSetNextHop.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC && (oldCItem.AcceptSetServiceChainNameVariable.ValueString() != "" || data.Sequences[ni].Actions[nci].AcceptSetServiceChainNameVariable.ValueString() != "") {
									if oldCItem.AcceptSetServiceChainNameVariable.ValueString() != data.Sequences[ni].Actions[nci].AcceptSetServiceChainNameVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.AcceptSetServiceChainName.ValueString() != data.Sequences[ni].Actions[nci].AcceptSetServiceChainName.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC && (oldCItem.AcceptSetServiceChainVpnVariable.ValueString() != "" || data.Sequences[ni].Actions[nci].AcceptSetServiceChainVpnVariable.ValueString() != "") {
									if oldCItem.AcceptSetServiceChainVpnVariable.ValueString() != data.Sequences[ni].Actions[nci].AcceptSetServiceChainVpnVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.AcceptSetServiceChainVpn.ValueInt64() != data.Sequences[ni].Actions[nci].AcceptSetServiceChainVpn.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC && (oldCItem.AcceptSetServiceChainFallbackVariable.ValueString() != "" || data.Sequences[ni].Actions[nci].AcceptSetServiceChainFallbackVariable.ValueString() != "") {
									if oldCItem.AcceptSetServiceChainFallbackVariable.ValueString() != data.Sequences[ni].Actions[nci].AcceptSetServiceChainFallbackVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.AcceptSetServiceChainFallback.ValueBool() != data.Sequences[ni].Actions[nci].AcceptSetServiceChainFallback.ValueBool() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.AcceptMirrorListId.ValueString() != data.Sequences[ni].Actions[nci].AcceptMirrorListId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.AcceptPolicerId.ValueString() != data.Sequences[ni].Actions[nci].AcceptPolicerId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.DropCounterName.ValueString() != data.Sequences[ni].Actions[nci].DropCounterName.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.DropLog.ValueBool() != data.Sequences[ni].Actions[nci].DropLog.ValueBool() {
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
