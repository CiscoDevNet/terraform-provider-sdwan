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
type ServiceIPv4ACL struct {
	Id               types.String              `tfsdk:"id"`
	Version          types.Int64               `tfsdk:"version"`
	Name             types.String              `tfsdk:"name"`
	Description      types.String              `tfsdk:"description"`
	FeatureProfileId types.String              `tfsdk:"feature_profile_id"`
	DefaultAction    types.String              `tfsdk:"default_action"`
	Sequences        []ServiceIPv4ACLSequences `tfsdk:"sequences"`
}

type ServiceIPv4ACLSequences struct {
	SequenceId   types.Int64                           `tfsdk:"sequence_id"`
	SequenceName types.String                          `tfsdk:"sequence_name"`
	BaseAction   types.String                          `tfsdk:"base_action"`
	MatchEntries []ServiceIPv4ACLSequencesMatchEntries `tfsdk:"match_entries"`
	Actions      []ServiceIPv4ACLSequencesActions      `tfsdk:"actions"`
}

type ServiceIPv4ACLSequencesMatchEntries struct {
	Dscps                         types.Set                                             `tfsdk:"dscps"`
	PacketLength                  types.String                                          `tfsdk:"packet_length"`
	Protocols                     types.Set                                             `tfsdk:"protocols"`
	IcmpMessages                  types.Set                                             `tfsdk:"icmp_messages"`
	SourceDataPrefixListId        types.String                                          `tfsdk:"source_data_prefix_list_id"`
	SourceDataPrefix              types.String                                          `tfsdk:"source_data_prefix"`
	SourceDataPrefixVariable      types.String                                          `tfsdk:"source_data_prefix_variable"`
	SourcePorts                   []ServiceIPv4ACLSequencesMatchEntriesSourcePorts      `tfsdk:"source_ports"`
	DestinationDataPrefixListId   types.String                                          `tfsdk:"destination_data_prefix_list_id"`
	DestinationDataPrefix         types.String                                          `tfsdk:"destination_data_prefix"`
	DestinationDataPrefixVariable types.String                                          `tfsdk:"destination_data_prefix_variable"`
	DestinationPorts              []ServiceIPv4ACLSequencesMatchEntriesDestinationPorts `tfsdk:"destination_ports"`
	TcpState                      types.String                                          `tfsdk:"tcp_state"`
}
type ServiceIPv4ACLSequencesActions struct {
	AcceptSetDscp      types.Int64  `tfsdk:"accept_set_dscp"`
	AcceptCounterName  types.String `tfsdk:"accept_counter_name"`
	AcceptLog          types.Bool   `tfsdk:"accept_log"`
	AcceptSetNextHop   types.String `tfsdk:"accept_set_next_hop"`
	AcceptMirrorListId types.String `tfsdk:"accept_mirror_list_id"`
	AcceptPolicerId    types.String `tfsdk:"accept_policer_id"`
	DropCounterName    types.String `tfsdk:"drop_counter_name"`
	DropLog            types.Bool   `tfsdk:"drop_log"`
}

type ServiceIPv4ACLSequencesMatchEntriesSourcePorts struct {
	Port types.String `tfsdk:"port"`
}
type ServiceIPv4ACLSequencesMatchEntriesDestinationPorts struct {
	Port types.String `tfsdk:"port"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceIPv4ACL) getModel() string {
	return "service_ipv4_acl"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceIPv4ACL) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/ipv4-acl", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceIPv4ACL) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if data.DefaultAction.IsNull() {
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
func (data *ServiceIPv4ACL) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "sequences"); value.Exists() && len(value.Array()) > 0 {
		data.Sequences = make([]ServiceIPv4ACLSequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceIPv4ACLSequences{}
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
				item.MatchEntries = make([]ServiceIPv4ACLSequencesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceIPv4ACLSequencesMatchEntries{}
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
						cItem.SourcePorts = make([]ServiceIPv4ACLSequencesMatchEntriesSourcePorts, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := ServiceIPv4ACLSequencesMatchEntriesSourcePorts{}
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
						cItem.DestinationPorts = make([]ServiceIPv4ACLSequencesMatchEntriesDestinationPorts, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := ServiceIPv4ACLSequencesMatchEntriesDestinationPorts{}
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
				item.Actions = make([]ServiceIPv4ACLSequencesActions, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceIPv4ACLSequencesActions{}
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
	}
}

// End of section. //template:end fromBody

func (data *ServiceIPv4ACL) updateFromBody(ctx context.Context, res gjson.Result) {
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
		keyValues := [...]string{strconv.FormatInt(data.Sequences[i].SequenceId.ValueInt64(), 10)}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "sequences").ForEach(
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
		data.Sequences[i].SequenceId = types.Int64Null()

		if t := r.Get("sequenceId.optionType"); t.Exists() {
			va := r.Get("sequenceId.value")
			if t.String() == "global" {
				data.Sequences[i].SequenceId = types.Int64Value(va.Int())
			}
		}
		data.Sequences[i].SequenceName = types.StringNull()

		if t := r.Get("sequenceName.optionType"); t.Exists() {
			va := r.Get("sequenceName.value")
			if t.String() == "global" {
				data.Sequences[i].SequenceName = types.StringValue(va.String())
			}
		}
		data.Sequences[i].BaseAction = types.StringNull()

		if t := r.Get("baseAction.optionType"); t.Exists() {
			va := r.Get("baseAction.value")
			if t.String() == "global" {
				data.Sequences[i].BaseAction = types.StringValue(va.String())
			}
		}
		for ci := range data.Sequences[i].MatchEntries {
			if len(r.Get("matchEntries").Array()) <= 0 {
				continue
			}
			cr := r.Get("matchEntries").Array()[0]

			data.Sequences[i].MatchEntries[ci].Dscps = types.SetNull(types.Int64Type)

			if t := cr.Get("dscp.optionType"); t.Exists() {
				va := cr.Get("dscp.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].Dscps = helpers.GetInt64Set(va.Array())
				}
			}
			data.Sequences[i].MatchEntries[ci].PacketLength = types.StringNull()

			if t := cr.Get("packetLength.optionType"); t.Exists() {
				va := cr.Get("packetLength.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].PacketLength = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].Protocols = types.SetNull(types.Int64Type)

			if t := cr.Get("protocol.optionType"); t.Exists() {
				va := cr.Get("protocol.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].Protocols = helpers.GetInt64Set(va.Array())
				}
			}
			data.Sequences[i].MatchEntries[ci].IcmpMessages = types.SetNull(types.StringType)

			if t := cr.Get("icmpMsg.optionType"); t.Exists() {
				va := cr.Get("icmpMsg.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].IcmpMessages = helpers.GetStringSet(va.Array())
				}
			}
			data.Sequences[i].MatchEntries[ci].SourceDataPrefixListId = types.StringNull()

			if t := cr.Get("sourceDataPrefix.sourceDataPrefixList.refId.optionType"); t.Exists() {
				va := cr.Get("sourceDataPrefix.sourceDataPrefixList.refId.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].SourceDataPrefixListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].SourceDataPrefix = types.StringNull()
			data.Sequences[i].MatchEntries[ci].SourceDataPrefixVariable = types.StringNull()
			if t := cr.Get("sourceDataPrefix.sourceIpPrefix.optionType"); t.Exists() {
				va := cr.Get("sourceDataPrefix.sourceIpPrefix.value")
				if t.String() == "variable" {
					data.Sequences[i].MatchEntries[ci].SourceDataPrefixVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].SourceDataPrefix = types.StringValue(va.String())
				}
			}
			for cci := range data.Sequences[i].MatchEntries[ci].SourcePorts {
				keys := [...]string{"sourcePort"}
				keyValues := [...]string{data.Sequences[i].MatchEntries[ci].SourcePorts[cci].Port.ValueString()}
				keyValuesVariables := [...]string{""}

				var ccr gjson.Result
				cr.Get("sourcePorts").ForEach(
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
							ccr = v
							return false
						}
						return true
					},
				)
				data.Sequences[i].MatchEntries[ci].SourcePorts[cci].Port = types.StringNull()

				if t := ccr.Get("sourcePort.optionType"); t.Exists() {
					va := ccr.Get("sourcePort.value")
					if t.String() == "global" {
						data.Sequences[i].MatchEntries[ci].SourcePorts[cci].Port = types.StringValue(va.String())
					}
				}
			}
			data.Sequences[i].MatchEntries[ci].DestinationDataPrefixListId = types.StringNull()

			if t := cr.Get("destinationDataPrefix.destinationDataPrefixList.refId.optionType"); t.Exists() {
				va := cr.Get("destinationDataPrefix.destinationDataPrefixList.refId.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].DestinationDataPrefixListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].MatchEntries[ci].DestinationDataPrefix = types.StringNull()
			data.Sequences[i].MatchEntries[ci].DestinationDataPrefixVariable = types.StringNull()
			if t := cr.Get("destinationDataPrefix.destinationIpPrefix.optionType"); t.Exists() {
				va := cr.Get("destinationDataPrefix.destinationIpPrefix.value")
				if t.String() == "variable" {
					data.Sequences[i].MatchEntries[ci].DestinationDataPrefixVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].DestinationDataPrefix = types.StringValue(va.String())
				}
			}
			for cci := range data.Sequences[i].MatchEntries[ci].DestinationPorts {
				keys := [...]string{"destinationPort"}
				keyValues := [...]string{data.Sequences[i].MatchEntries[ci].DestinationPorts[cci].Port.ValueString()}
				keyValuesVariables := [...]string{""}

				var ccr gjson.Result
				cr.Get("destinationPorts").ForEach(
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
							ccr = v
							return false
						}
						return true
					},
				)
				data.Sequences[i].MatchEntries[ci].DestinationPorts[cci].Port = types.StringNull()

				if t := ccr.Get("destinationPort.optionType"); t.Exists() {
					va := ccr.Get("destinationPort.value")
					if t.String() == "global" {
						data.Sequences[i].MatchEntries[ci].DestinationPorts[cci].Port = types.StringValue(va.String())
					}
				}
			}
			data.Sequences[i].MatchEntries[ci].TcpState = types.StringNull()

			if t := cr.Get("tcp.optionType"); t.Exists() {
				va := cr.Get("tcp.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].TcpState = types.StringValue(va.String())
				}
			}
		}
		for ci := range data.Sequences[i].Actions {
			keys := [...]string{"accept.setDscp", "accept.counterName", "accept.log", "accept.setNextHop", "accept.mirror.refId", "accept.policer.refId", "drop.counterName", "drop.log"}
			keyValues := [...]string{strconv.FormatInt(data.Sequences[i].Actions[ci].AcceptSetDscp.ValueInt64(), 10), data.Sequences[i].Actions[ci].AcceptCounterName.ValueString(), strconv.FormatBool(data.Sequences[i].Actions[ci].AcceptLog.ValueBool()), data.Sequences[i].Actions[ci].AcceptSetNextHop.ValueString(), data.Sequences[i].Actions[ci].AcceptMirrorListId.ValueString(), data.Sequences[i].Actions[ci].AcceptPolicerId.ValueString(), data.Sequences[i].Actions[ci].DropCounterName.ValueString(), strconv.FormatBool(data.Sequences[i].Actions[ci].DropLog.ValueBool())}
			keyValuesVariables := [...]string{"", "", "", "", "", "", "", ""}

			var cr gjson.Result
			r.Get("actions").ForEach(
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
			data.Sequences[i].Actions[ci].AcceptSetDscp = types.Int64Null()

			if t := cr.Get("accept.setDscp.optionType"); t.Exists() {
				va := cr.Get("accept.setDscp.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].AcceptSetDscp = types.Int64Value(va.Int())
				}
			}
			data.Sequences[i].Actions[ci].AcceptCounterName = types.StringNull()

			if t := cr.Get("accept.counterName.optionType"); t.Exists() {
				va := cr.Get("accept.counterName.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].AcceptCounterName = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Actions[ci].AcceptLog = types.BoolNull()

			if t := cr.Get("accept.log.optionType"); t.Exists() {
				va := cr.Get("accept.log.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].AcceptLog = types.BoolValue(va.Bool())
				}
			}
			data.Sequences[i].Actions[ci].AcceptSetNextHop = types.StringNull()

			if t := cr.Get("accept.setNextHop.optionType"); t.Exists() {
				va := cr.Get("accept.setNextHop.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].AcceptSetNextHop = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Actions[ci].AcceptMirrorListId = types.StringNull()

			if t := cr.Get("accept.mirror.refId.optionType"); t.Exists() {
				va := cr.Get("accept.mirror.refId.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].AcceptMirrorListId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Actions[ci].AcceptPolicerId = types.StringNull()

			if t := cr.Get("accept.policer.refId.optionType"); t.Exists() {
				va := cr.Get("accept.policer.refId.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].AcceptPolicerId = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Actions[ci].DropCounterName = types.StringNull()

			if t := cr.Get("drop.counterName.optionType"); t.Exists() {
				va := cr.Get("drop.counterName.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].DropCounterName = types.StringValue(va.String())
				}
			}
			data.Sequences[i].Actions[ci].DropLog = types.BoolNull()

			if t := cr.Get("drop.log.optionType"); t.Exists() {
				va := cr.Get("drop.log.value")
				if t.String() == "global" {
					data.Sequences[i].Actions[ci].DropLog = types.BoolValue(va.Bool())
				}
			}
		}
	}
}
