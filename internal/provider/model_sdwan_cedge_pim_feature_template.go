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
type CEdgePIM struct {
	Id                            types.String               `tfsdk:"id"`
	Version                       types.Int64                `tfsdk:"version"`
	TemplateType                  types.String               `tfsdk:"template_type"`
	Name                          types.String               `tfsdk:"name"`
	Description                   types.String               `tfsdk:"description"`
	DeviceTypes                   types.Set                  `tfsdk:"device_types"`
	AutoRp                        types.Bool                 `tfsdk:"auto_rp"`
	AutoRpVariable                types.String               `tfsdk:"auto_rp_variable"`
	RpAnnounceFields              []CEdgePIMRpAnnounceFields `tfsdk:"rp_announce_fields"`
	InterfaceName                 types.String               `tfsdk:"interface_name"`
	InterfaceNameVariable         types.String               `tfsdk:"interface_name_variable"`
	RpCandidates                  []CEdgePIMRpCandidates     `tfsdk:"rp_candidates"`
	BsrCandidate                  types.String               `tfsdk:"bsr_candidate"`
	BsrCandidateVariable          types.String               `tfsdk:"bsr_candidate_variable"`
	HashMaskLength                types.String               `tfsdk:"hash_mask_length"`
	HashMaskLengthVariable        types.String               `tfsdk:"hash_mask_length_variable"`
	Priority                      types.Int64                `tfsdk:"priority"`
	PriorityVariable              types.String               `tfsdk:"priority_variable"`
	RpCandidateAccessList         types.String               `tfsdk:"rp_candidate_access_list"`
	RpCandidateAccessListVariable types.String               `tfsdk:"rp_candidate_access_list_variable"`
	Scope                         types.Int64                `tfsdk:"scope"`
	ScopeVariable                 types.String               `tfsdk:"scope_variable"`
	Range                         types.String               `tfsdk:"range"`
	RangeVariable                 types.String               `tfsdk:"range_variable"`
	Default                       types.Bool                 `tfsdk:"default"`
	DefaultVariable               types.String               `tfsdk:"default_variable"`
	RpAddresses                   []CEdgePIMRpAddresses      `tfsdk:"rp_addresses"`
	SptThreshold                  types.String               `tfsdk:"spt_threshold"`
	SptThresholdVariable          types.String               `tfsdk:"spt_threshold_variable"`
	Interfaces                    []CEdgePIMInterfaces       `tfsdk:"interfaces"`
}

type CEdgePIMRpAnnounceFields struct {
	Optional              types.Bool   `tfsdk:"optional"`
	InterfaceName         types.String `tfsdk:"interface_name"`
	InterfaceNameVariable types.String `tfsdk:"interface_name_variable"`
	Scope                 types.Int64  `tfsdk:"scope"`
	ScopeVariable         types.String `tfsdk:"scope_variable"`
}

type CEdgePIMRpCandidates struct {
	Optional           types.Bool   `tfsdk:"optional"`
	Interface          types.String `tfsdk:"interface"`
	InterfaceVariable  types.String `tfsdk:"interface_variable"`
	AccessList         types.String `tfsdk:"access_list"`
	AccessListVariable types.String `tfsdk:"access_list_variable"`
	Interval           types.Int64  `tfsdk:"interval"`
	IntervalVariable   types.String `tfsdk:"interval_variable"`
	Priority           types.Int64  `tfsdk:"priority"`
	PriorityVariable   types.String `tfsdk:"priority_variable"`
}

type CEdgePIMRpAddresses struct {
	Optional           types.Bool   `tfsdk:"optional"`
	AccessList         types.String `tfsdk:"access_list"`
	AccessListVariable types.String `tfsdk:"access_list_variable"`
	IpAddress          types.String `tfsdk:"ip_address"`
	IpAddressVariable  types.String `tfsdk:"ip_address_variable"`
	Override           types.Bool   `tfsdk:"override"`
	OverrideVariable   types.String `tfsdk:"override_variable"`
}

type CEdgePIMInterfaces struct {
	Optional                  types.Bool   `tfsdk:"optional"`
	InterfaceName             types.String `tfsdk:"interface_name"`
	InterfaceNameVariable     types.String `tfsdk:"interface_name_variable"`
	QueryInterval             types.Int64  `tfsdk:"query_interval"`
	QueryIntervalVariable     types.String `tfsdk:"query_interval_variable"`
	JoinPruneInterval         types.Int64  `tfsdk:"join_prune_interval"`
	JoinPruneIntervalVariable types.String `tfsdk:"join_prune_interval_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CEdgePIM) getModel() string {
	return "cedge_pim"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CEdgePIM) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cedge_pim")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.AutoRpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pim.auto-rp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.auto-rp."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pim.auto-rp."+"vipVariableName", data.AutoRpVariable.ValueString())
	} else if data.AutoRp.IsNull() {
		body, _ = sjson.Set(body, path+"pim.auto-rp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.auto-rp."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"pim.auto-rp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.auto-rp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.auto-rp."+"vipValue", strconv.FormatBool(data.AutoRp.ValueBool()))
	}
	if len(data.RpAnnounceFields) > 0 {
		body, _ = sjson.Set(body, path+"pim.send-rp-announce.send-rp-announce-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"pim.send-rp-announce.send-rp-announce-list."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.send-rp-announce.send-rp-announce-list."+"vipPrimaryKey", []string{"if-name"})
		body, _ = sjson.Set(body, path+"pim.send-rp-announce.send-rp-announce-list."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"pim.send-rp-announce.send-rp-announce-list."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"pim.send-rp-announce.send-rp-announce-list."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"pim.send-rp-announce.send-rp-announce-list."+"vipPrimaryKey", []string{"if-name"})
		body, _ = sjson.Set(body, path+"pim.send-rp-announce.send-rp-announce-list."+"vipValue", []interface{}{})
	}
	for _, item := range data.RpAnnounceFields {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "if-name")

		if !item.InterfaceNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipVariableName", item.InterfaceNameVariable.ValueString())
		} else if item.InterfaceName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipValue", item.InterfaceName.ValueString())
		}
		itemAttributes = append(itemAttributes, "scope")

		if !item.ScopeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "scope."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "scope."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "scope."+"vipVariableName", item.ScopeVariable.ValueString())
		} else if item.Scope.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "scope."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "scope."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "scope."+"vipValue", item.Scope.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"pim.send-rp-announce.send-rp-announce-list."+"vipValue.-1", itemBody)
	}

	if !data.InterfaceNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pim.send-rp-discovery.if-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.send-rp-discovery.if-name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pim.send-rp-discovery.if-name."+"vipVariableName", data.InterfaceNameVariable.ValueString())
	} else if data.InterfaceName.IsNull() {
		if !gjson.Get(body, path+"pim.send-rp-discovery").Exists() {
			body, _ = sjson.Set(body, path+"pim.send-rp-discovery", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"pim.send-rp-discovery.if-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.send-rp-discovery.if-name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.send-rp-discovery.if-name."+"vipValue", data.InterfaceName.ValueString())
	}
	if len(data.RpCandidates) > 0 {
		body, _ = sjson.Set(body, path+"pim.rp-candidate."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"pim.rp-candidate."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.rp-candidate."+"vipPrimaryKey", []string{"pim-interface-name"})
		body, _ = sjson.Set(body, path+"pim.rp-candidate."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"pim.rp-candidate."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"pim.rp-candidate."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"pim.rp-candidate."+"vipPrimaryKey", []string{"pim-interface-name"})
		body, _ = sjson.Set(body, path+"pim.rp-candidate."+"vipValue", []interface{}{})
	}
	for _, item := range data.RpCandidates {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "pim-interface-name")

		if !item.InterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "pim-interface-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "pim-interface-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "pim-interface-name."+"vipVariableName", item.InterfaceVariable.ValueString())
		} else if item.Interface.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "pim-interface-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "pim-interface-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "pim-interface-name."+"vipValue", item.Interface.ValueString())
		}
		itemAttributes = append(itemAttributes, "group-list")

		if !item.AccessListVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "group-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "group-list."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "group-list."+"vipVariableName", item.AccessListVariable.ValueString())
		} else if item.AccessList.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "group-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "group-list."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "group-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "group-list."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "group-list."+"vipValue", item.AccessList.ValueString())
		}
		itemAttributes = append(itemAttributes, "interval")

		if !item.IntervalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipVariableName", item.IntervalVariable.ValueString())
		} else if item.Interval.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipValue", item.Interval.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "priority")

		if !item.PriorityVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipVariableName", item.PriorityVariable.ValueString())
		} else if item.Priority.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "priority."+"vipValue", item.Priority.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"pim.rp-candidate."+"vipValue.-1", itemBody)
	}

	if !data.BsrCandidateVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.bsr-interface-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.bsr-interface-name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.bsr-interface-name."+"vipVariableName", data.BsrCandidateVariable.ValueString())
	} else if data.BsrCandidate.IsNull() {
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.bsr-interface-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.bsr-interface-name."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.bsr-interface-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.bsr-interface-name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.bsr-interface-name."+"vipValue", data.BsrCandidate.ValueString())
	}

	if !data.HashMaskLengthVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.mask."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.mask."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.mask."+"vipVariableName", data.HashMaskLengthVariable.ValueString())
	} else if data.HashMaskLength.IsNull() {
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.mask."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.mask."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.mask."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.mask."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.mask."+"vipValue", data.HashMaskLength.ValueString())
	}

	if !data.PriorityVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.priority."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.priority."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.priority."+"vipVariableName", data.PriorityVariable.ValueString())
	} else if data.Priority.IsNull() {
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.priority."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.priority."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.priority."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.priority."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.priority."+"vipValue", data.Priority.ValueInt64())
	}

	if !data.RpCandidateAccessListVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.accept-rp-candidate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.accept-rp-candidate."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.accept-rp-candidate."+"vipVariableName", data.RpCandidateAccessListVariable.ValueString())
	} else if data.RpCandidateAccessList.IsNull() {
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.accept-rp-candidate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.accept-rp-candidate."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.accept-rp-candidate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.accept-rp-candidate."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.bsr-candidate.accept-rp-candidate."+"vipValue", data.RpCandidateAccessList.ValueString())
	}

	if !data.ScopeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pim.send-rp-discovery.scope."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.send-rp-discovery.scope."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pim.send-rp-discovery.scope."+"vipVariableName", data.ScopeVariable.ValueString())
	} else if data.Scope.IsNull() {
		if !gjson.Get(body, path+"pim.send-rp-discovery").Exists() {
			body, _ = sjson.Set(body, path+"pim.send-rp-discovery", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"pim.send-rp-discovery.scope."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.send-rp-discovery.scope."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.send-rp-discovery.scope."+"vipValue", data.Scope.ValueInt64())
	}

	if !data.RangeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pim.ssm.range."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.ssm.range."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pim.ssm.range."+"vipVariableName", data.RangeVariable.ValueString())
	} else if data.Range.IsNull() {
		body, _ = sjson.Set(body, path+"pim.ssm.range."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.ssm.range."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"pim.ssm.range."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.ssm.range."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.ssm.range."+"vipValue", data.Range.ValueString())
	}

	if !data.DefaultVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pim.ssm.default."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"pim.ssm.default."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pim.ssm.default."+"vipVariableName", data.DefaultVariable.ValueString())
	} else if data.Default.IsNull() {
		body, _ = sjson.Set(body, path+"pim.ssm.default."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"pim.ssm.default."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"pim.ssm.default."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"pim.ssm.default."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.ssm.default."+"vipValue", strconv.FormatBool(data.Default.ValueBool()))
	}
	if len(data.RpAddresses) > 0 {
		body, _ = sjson.Set(body, path+"pim.rp-addr."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"pim.rp-addr."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.rp-addr."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"pim.rp-addr."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"pim.rp-addr."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"pim.rp-addr."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"pim.rp-addr."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"pim.rp-addr."+"vipValue", []interface{}{})
	}
	for _, item := range data.RpAddresses {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "access-list")

		if !item.AccessListVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "access-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "access-list."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "access-list."+"vipVariableName", item.AccessListVariable.ValueString())
		} else if item.AccessList.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "access-list."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "access-list."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "access-list."+"vipValue", item.AccessList.ValueString())
		}
		itemAttributes = append(itemAttributes, "address")

		if !item.IpAddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipVariableName", item.IpAddressVariable.ValueString())
		} else if item.IpAddress.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipValue", item.IpAddress.ValueString())
		}
		itemAttributes = append(itemAttributes, "override")

		if !item.OverrideVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "override."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "override."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "override."+"vipVariableName", item.OverrideVariable.ValueString())
		} else if item.Override.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "override."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "override."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "override."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "override."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "override."+"vipValue", strconv.FormatBool(item.Override.ValueBool()))
		}
		itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		body, _ = sjson.SetRaw(body, path+"pim.rp-addr."+"vipValue.-1", itemBody)
	}

	if !data.SptThresholdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"pim.spt-threshold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.spt-threshold."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"pim.spt-threshold."+"vipVariableName", data.SptThresholdVariable.ValueString())
	} else if data.SptThreshold.IsNull() {
		if !gjson.Get(body, path+"pim").Exists() {
			body, _ = sjson.Set(body, path+"pim", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"pim.spt-threshold."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"pim.spt-threshold."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.spt-threshold."+"vipValue", data.SptThreshold.ValueString())
	}
	if len(data.Interfaces) > 0 {
		body, _ = sjson.Set(body, path+"pim.interface."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"pim.interface."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"pim.interface."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"pim.interface."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"pim.interface."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"pim.interface."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"pim.interface."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"pim.interface."+"vipValue", []interface{}{})
	}
	for _, item := range data.Interfaces {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")

		if !item.InterfaceNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipVariableName", item.InterfaceNameVariable.ValueString())
		} else if item.InterfaceName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.InterfaceName.ValueString())
		}
		itemAttributes = append(itemAttributes, "query-interval")

		if !item.QueryIntervalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "query-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "query-interval."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "query-interval."+"vipVariableName", item.QueryIntervalVariable.ValueString())
		} else if item.QueryInterval.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "query-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "query-interval."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "query-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "query-interval."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "query-interval."+"vipValue", item.QueryInterval.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "join-prune-interval")

		if !item.JoinPruneIntervalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "join-prune-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "join-prune-interval."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "join-prune-interval."+"vipVariableName", item.JoinPruneIntervalVariable.ValueString())
		} else if item.JoinPruneInterval.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "join-prune-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "join-prune-interval."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "join-prune-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "join-prune-interval."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "join-prune-interval."+"vipValue", item.JoinPruneInterval.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"pim.interface."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CEdgePIM) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "pim.auto-rp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AutoRp = types.BoolNull()

			v := res.Get(path + "pim.auto-rp.vipVariableName")
			data.AutoRpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AutoRp = types.BoolNull()
			data.AutoRpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pim.auto-rp.vipValue")
			data.AutoRp = types.BoolValue(v.Bool())
			data.AutoRpVariable = types.StringNull()
		}
	} else {
		data.AutoRp = types.BoolNull()
		data.AutoRpVariable = types.StringNull()
	}
	if value := res.Get(path + "pim.send-rp-announce.send-rp-announce-list.vipValue"); len(value.Array()) > 0 {
		data.RpAnnounceFields = make([]CEdgePIMRpAnnounceFields, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CEdgePIMRpAnnounceFields{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("if-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.InterfaceName = types.StringNull()

					cv := v.Get("if-name.vipVariableName")
					item.InterfaceNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.InterfaceName = types.StringNull()
					item.InterfaceNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("if-name.vipValue")
					item.InterfaceName = types.StringValue(cv.String())
					item.InterfaceNameVariable = types.StringNull()
				}
			} else {
				item.InterfaceName = types.StringNull()
				item.InterfaceNameVariable = types.StringNull()
			}
			if cValue := v.Get("scope.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Scope = types.Int64Null()

					cv := v.Get("scope.vipVariableName")
					item.ScopeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Scope = types.Int64Null()
					item.ScopeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("scope.vipValue")
					item.Scope = types.Int64Value(cv.Int())
					item.ScopeVariable = types.StringNull()
				}
			} else {
				item.Scope = types.Int64Null()
				item.ScopeVariable = types.StringNull()
			}
			data.RpAnnounceFields = append(data.RpAnnounceFields, item)
			return true
		})
	} else {
		if len(data.RpAnnounceFields) > 0 {
			data.RpAnnounceFields = []CEdgePIMRpAnnounceFields{}
		}
	}
	if value := res.Get(path + "pim.send-rp-discovery.if-name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.InterfaceName = types.StringNull()

			v := res.Get(path + "pim.send-rp-discovery.if-name.vipVariableName")
			data.InterfaceNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.InterfaceName = types.StringNull()
			data.InterfaceNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pim.send-rp-discovery.if-name.vipValue")
			data.InterfaceName = types.StringValue(v.String())
			data.InterfaceNameVariable = types.StringNull()
		}
	} else {
		data.InterfaceName = types.StringNull()
		data.InterfaceNameVariable = types.StringNull()
	}
	if value := res.Get(path + "pim.rp-candidate.vipValue"); len(value.Array()) > 0 {
		data.RpCandidates = make([]CEdgePIMRpCandidates, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CEdgePIMRpCandidates{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("pim-interface-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Interface = types.StringNull()

					cv := v.Get("pim-interface-name.vipVariableName")
					item.InterfaceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Interface = types.StringNull()
					item.InterfaceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("pim-interface-name.vipValue")
					item.Interface = types.StringValue(cv.String())
					item.InterfaceVariable = types.StringNull()
				}
			} else {
				item.Interface = types.StringNull()
				item.InterfaceVariable = types.StringNull()
			}
			if cValue := v.Get("group-list.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AccessList = types.StringNull()

					cv := v.Get("group-list.vipVariableName")
					item.AccessListVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AccessList = types.StringNull()
					item.AccessListVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("group-list.vipValue")
					item.AccessList = types.StringValue(cv.String())
					item.AccessListVariable = types.StringNull()
				}
			} else {
				item.AccessList = types.StringNull()
				item.AccessListVariable = types.StringNull()
			}
			if cValue := v.Get("interval.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Interval = types.Int64Null()

					cv := v.Get("interval.vipVariableName")
					item.IntervalVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Interval = types.Int64Null()
					item.IntervalVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("interval.vipValue")
					item.Interval = types.Int64Value(cv.Int())
					item.IntervalVariable = types.StringNull()
				}
			} else {
				item.Interval = types.Int64Null()
				item.IntervalVariable = types.StringNull()
			}
			if cValue := v.Get("priority.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Priority = types.Int64Null()

					cv := v.Get("priority.vipVariableName")
					item.PriorityVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Priority = types.Int64Null()
					item.PriorityVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("priority.vipValue")
					item.Priority = types.Int64Value(cv.Int())
					item.PriorityVariable = types.StringNull()
				}
			} else {
				item.Priority = types.Int64Null()
				item.PriorityVariable = types.StringNull()
			}
			data.RpCandidates = append(data.RpCandidates, item)
			return true
		})
	} else {
		if len(data.RpCandidates) > 0 {
			data.RpCandidates = []CEdgePIMRpCandidates{}
		}
	}
	if value := res.Get(path + "pim.bsr-candidate.bsr-interface-name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.BsrCandidate = types.StringNull()

			v := res.Get(path + "pim.bsr-candidate.bsr-interface-name.vipVariableName")
			data.BsrCandidateVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.BsrCandidate = types.StringNull()
			data.BsrCandidateVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pim.bsr-candidate.bsr-interface-name.vipValue")
			data.BsrCandidate = types.StringValue(v.String())
			data.BsrCandidateVariable = types.StringNull()
		}
	} else {
		data.BsrCandidate = types.StringNull()
		data.BsrCandidateVariable = types.StringNull()
	}
	if value := res.Get(path + "pim.bsr-candidate.mask.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.HashMaskLength = types.StringNull()

			v := res.Get(path + "pim.bsr-candidate.mask.vipVariableName")
			data.HashMaskLengthVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.HashMaskLength = types.StringNull()
			data.HashMaskLengthVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pim.bsr-candidate.mask.vipValue")
			data.HashMaskLength = types.StringValue(v.String())
			data.HashMaskLengthVariable = types.StringNull()
		}
	} else {
		data.HashMaskLength = types.StringNull()
		data.HashMaskLengthVariable = types.StringNull()
	}
	if value := res.Get(path + "pim.bsr-candidate.priority.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Priority = types.Int64Null()

			v := res.Get(path + "pim.bsr-candidate.priority.vipVariableName")
			data.PriorityVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Priority = types.Int64Null()
			data.PriorityVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pim.bsr-candidate.priority.vipValue")
			data.Priority = types.Int64Value(v.Int())
			data.PriorityVariable = types.StringNull()
		}
	} else {
		data.Priority = types.Int64Null()
		data.PriorityVariable = types.StringNull()
	}
	if value := res.Get(path + "pim.bsr-candidate.accept-rp-candidate.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.RpCandidateAccessList = types.StringNull()

			v := res.Get(path + "pim.bsr-candidate.accept-rp-candidate.vipVariableName")
			data.RpCandidateAccessListVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.RpCandidateAccessList = types.StringNull()
			data.RpCandidateAccessListVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pim.bsr-candidate.accept-rp-candidate.vipValue")
			data.RpCandidateAccessList = types.StringValue(v.String())
			data.RpCandidateAccessListVariable = types.StringNull()
		}
	} else {
		data.RpCandidateAccessList = types.StringNull()
		data.RpCandidateAccessListVariable = types.StringNull()
	}
	if value := res.Get(path + "pim.send-rp-discovery.scope.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Scope = types.Int64Null()

			v := res.Get(path + "pim.send-rp-discovery.scope.vipVariableName")
			data.ScopeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Scope = types.Int64Null()
			data.ScopeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pim.send-rp-discovery.scope.vipValue")
			data.Scope = types.Int64Value(v.Int())
			data.ScopeVariable = types.StringNull()
		}
	} else {
		data.Scope = types.Int64Null()
		data.ScopeVariable = types.StringNull()
	}
	if value := res.Get(path + "pim.ssm.range.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Range = types.StringNull()

			v := res.Get(path + "pim.ssm.range.vipVariableName")
			data.RangeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Range = types.StringNull()
			data.RangeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pim.ssm.range.vipValue")
			data.Range = types.StringValue(v.String())
			data.RangeVariable = types.StringNull()
		}
	} else {
		data.Range = types.StringNull()
		data.RangeVariable = types.StringNull()
	}
	if value := res.Get(path + "pim.ssm.default.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Default = types.BoolNull()

			v := res.Get(path + "pim.ssm.default.vipVariableName")
			data.DefaultVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Default = types.BoolNull()
			data.DefaultVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pim.ssm.default.vipValue")
			data.Default = types.BoolValue(v.Bool())
			data.DefaultVariable = types.StringNull()
		}
	} else {
		data.Default = types.BoolNull()
		data.DefaultVariable = types.StringNull()
	}
	if value := res.Get(path + "pim.rp-addr.vipValue"); len(value.Array()) > 0 {
		data.RpAddresses = make([]CEdgePIMRpAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CEdgePIMRpAddresses{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("access-list.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AccessList = types.StringNull()

					cv := v.Get("access-list.vipVariableName")
					item.AccessListVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AccessList = types.StringNull()
					item.AccessListVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("access-list.vipValue")
					item.AccessList = types.StringValue(cv.String())
					item.AccessListVariable = types.StringNull()
				}
			} else {
				item.AccessList = types.StringNull()
				item.AccessListVariable = types.StringNull()
			}
			if cValue := v.Get("address.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IpAddress = types.StringNull()

					cv := v.Get("address.vipVariableName")
					item.IpAddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IpAddress = types.StringNull()
					item.IpAddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("address.vipValue")
					item.IpAddress = types.StringValue(cv.String())
					item.IpAddressVariable = types.StringNull()
				}
			} else {
				item.IpAddress = types.StringNull()
				item.IpAddressVariable = types.StringNull()
			}
			if cValue := v.Get("override.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Override = types.BoolNull()

					cv := v.Get("override.vipVariableName")
					item.OverrideVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Override = types.BoolNull()
					item.OverrideVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("override.vipValue")
					item.Override = types.BoolValue(cv.Bool())
					item.OverrideVariable = types.StringNull()
				}
			} else {
				item.Override = types.BoolNull()
				item.OverrideVariable = types.StringNull()
			}
			data.RpAddresses = append(data.RpAddresses, item)
			return true
		})
	} else {
		if len(data.RpAddresses) > 0 {
			data.RpAddresses = []CEdgePIMRpAddresses{}
		}
	}
	if value := res.Get(path + "pim.spt-threshold.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SptThreshold = types.StringNull()

			v := res.Get(path + "pim.spt-threshold.vipVariableName")
			data.SptThresholdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SptThreshold = types.StringNull()
			data.SptThresholdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "pim.spt-threshold.vipValue")
			data.SptThreshold = types.StringValue(v.String())
			data.SptThresholdVariable = types.StringNull()
		}
	} else {
		data.SptThreshold = types.StringNull()
		data.SptThresholdVariable = types.StringNull()
	}
	if value := res.Get(path + "pim.interface.vipValue"); len(value.Array()) > 0 {
		data.Interfaces = make([]CEdgePIMInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CEdgePIMInterfaces{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.InterfaceName = types.StringNull()

					cv := v.Get("name.vipVariableName")
					item.InterfaceNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.InterfaceName = types.StringNull()
					item.InterfaceNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.InterfaceName = types.StringValue(cv.String())
					item.InterfaceNameVariable = types.StringNull()
				}
			} else {
				item.InterfaceName = types.StringNull()
				item.InterfaceNameVariable = types.StringNull()
			}
			if cValue := v.Get("query-interval.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.QueryInterval = types.Int64Null()

					cv := v.Get("query-interval.vipVariableName")
					item.QueryIntervalVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.QueryInterval = types.Int64Null()
					item.QueryIntervalVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("query-interval.vipValue")
					item.QueryInterval = types.Int64Value(cv.Int())
					item.QueryIntervalVariable = types.StringNull()
				}
			} else {
				item.QueryInterval = types.Int64Null()
				item.QueryIntervalVariable = types.StringNull()
			}
			if cValue := v.Get("join-prune-interval.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.JoinPruneInterval = types.Int64Null()

					cv := v.Get("join-prune-interval.vipVariableName")
					item.JoinPruneIntervalVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.JoinPruneInterval = types.Int64Null()
					item.JoinPruneIntervalVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("join-prune-interval.vipValue")
					item.JoinPruneInterval = types.Int64Value(cv.Int())
					item.JoinPruneIntervalVariable = types.StringNull()
				}
			} else {
				item.JoinPruneInterval = types.Int64Null()
				item.JoinPruneIntervalVariable = types.StringNull()
			}
			data.Interfaces = append(data.Interfaces, item)
			return true
		})
	} else {
		if len(data.Interfaces) > 0 {
			data.Interfaces = []CEdgePIMInterfaces{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CEdgePIM) hasChanges(ctx context.Context, state *CEdgePIM) bool {
	hasChanges := false
	if !data.AutoRp.Equal(state.AutoRp) {
		hasChanges = true
	}
	if len(data.RpAnnounceFields) != len(state.RpAnnounceFields) {
		hasChanges = true
	} else {
		for i := range data.RpAnnounceFields {
			if !data.RpAnnounceFields[i].InterfaceName.Equal(state.RpAnnounceFields[i].InterfaceName) {
				hasChanges = true
			}
			if !data.RpAnnounceFields[i].Scope.Equal(state.RpAnnounceFields[i].Scope) {
				hasChanges = true
			}
		}
	}
	if !data.InterfaceName.Equal(state.InterfaceName) {
		hasChanges = true
	}
	if len(data.RpCandidates) != len(state.RpCandidates) {
		hasChanges = true
	} else {
		for i := range data.RpCandidates {
			if !data.RpCandidates[i].Interface.Equal(state.RpCandidates[i].Interface) {
				hasChanges = true
			}
			if !data.RpCandidates[i].AccessList.Equal(state.RpCandidates[i].AccessList) {
				hasChanges = true
			}
			if !data.RpCandidates[i].Interval.Equal(state.RpCandidates[i].Interval) {
				hasChanges = true
			}
			if !data.RpCandidates[i].Priority.Equal(state.RpCandidates[i].Priority) {
				hasChanges = true
			}
		}
	}
	if !data.BsrCandidate.Equal(state.BsrCandidate) {
		hasChanges = true
	}
	if !data.HashMaskLength.Equal(state.HashMaskLength) {
		hasChanges = true
	}
	if !data.Priority.Equal(state.Priority) {
		hasChanges = true
	}
	if !data.RpCandidateAccessList.Equal(state.RpCandidateAccessList) {
		hasChanges = true
	}
	if !data.Scope.Equal(state.Scope) {
		hasChanges = true
	}
	if !data.Range.Equal(state.Range) {
		hasChanges = true
	}
	if !data.Default.Equal(state.Default) {
		hasChanges = true
	}
	if len(data.RpAddresses) != len(state.RpAddresses) {
		hasChanges = true
	} else {
		for i := range data.RpAddresses {
			if !data.RpAddresses[i].AccessList.Equal(state.RpAddresses[i].AccessList) {
				hasChanges = true
			}
			if !data.RpAddresses[i].IpAddress.Equal(state.RpAddresses[i].IpAddress) {
				hasChanges = true
			}
			if !data.RpAddresses[i].Override.Equal(state.RpAddresses[i].Override) {
				hasChanges = true
			}
		}
	}
	if !data.SptThreshold.Equal(state.SptThreshold) {
		hasChanges = true
	}
	if len(data.Interfaces) != len(state.Interfaces) {
		hasChanges = true
	} else {
		for i := range data.Interfaces {
			if !data.Interfaces[i].InterfaceName.Equal(state.Interfaces[i].InterfaceName) {
				hasChanges = true
			}
			if !data.Interfaces[i].QueryInterval.Equal(state.Interfaces[i].QueryInterval) {
				hasChanges = true
			}
			if !data.Interfaces[i].JoinPruneInterval.Equal(state.Interfaces[i].JoinPruneInterval) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
