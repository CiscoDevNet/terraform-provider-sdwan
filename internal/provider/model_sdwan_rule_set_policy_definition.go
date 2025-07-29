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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type RuleSetPolicyDefinition struct {
	Id          types.String                   `tfsdk:"id"`
	Version     types.Int64                    `tfsdk:"version"`
	Type        types.String                   `tfsdk:"type"`
	Name        types.String                   `tfsdk:"name"`
	Description types.String                   `tfsdk:"description"`
	Rules       []RuleSetPolicyDefinitionRules `tfsdk:"rules"`
}

type RuleSetPolicyDefinitionRules struct {
	Name                                 types.String `tfsdk:"name"`
	Order                                types.Int64  `tfsdk:"order"`
	SourceObjectGroupId                  types.String `tfsdk:"source_object_group_id"`
	SourceObjectGroupVersion             types.Int64  `tfsdk:"source_object_group_version"`
	SourceDataIpv4PrefixListId           types.String `tfsdk:"source_data_ipv4_prefix_list_id"`
	SourceDataIpv4PrefixListVersion      types.Int64  `tfsdk:"source_data_ipv4_prefix_list_version"`
	SourceIpv4Prefix                     types.String `tfsdk:"source_ipv4_prefix"`
	SourceIpv4PrefixVariable             types.String `tfsdk:"source_ipv4_prefix_variable"`
	SourceDataFqdnPrefixListId           types.String `tfsdk:"source_data_fqdn_prefix_list_id"`
	SourceDataFqdnPrefixListVersion      types.Int64  `tfsdk:"source_data_fqdn_prefix_list_version"`
	SourceFqdn                           types.String `tfsdk:"source_fqdn"`
	SourcePortListId                     types.String `tfsdk:"source_port_list_id"`
	SourcePortListVersion                types.Int64  `tfsdk:"source_port_list_version"`
	SourcePort                           types.String `tfsdk:"source_port"`
	SourceGeoLocationListId              types.String `tfsdk:"source_geo_location_list_id"`
	SourceGeoLocationListVersion         types.Int64  `tfsdk:"source_geo_location_list_version"`
	SourceGeoLocation                    types.String `tfsdk:"source_geo_location"`
	DestinationObjectGroupId             types.String `tfsdk:"destination_object_group_id"`
	DestinationObjectGroupVersion        types.Int64  `tfsdk:"destination_object_group_version"`
	DestinationDataIpv4PrefixListId      types.String `tfsdk:"destination_data_ipv4_prefix_list_id"`
	DestinationDataIpv4PrefixListVersion types.Int64  `tfsdk:"destination_data_ipv4_prefix_list_version"`
	DestinationIpv4Prefix                types.String `tfsdk:"destination_ipv4_prefix"`
	DestinationIpv4PrefixVariable        types.String `tfsdk:"destination_ipv4_prefix_variable"`
	DestinationDataFqdnPrefixListId      types.String `tfsdk:"destination_data_fqdn_prefix_list_id"`
	DestinationDataFqdnPrefixListVersion types.Int64  `tfsdk:"destination_data_fqdn_prefix_list_version"`
	DestinationFqdn                      types.String `tfsdk:"destination_fqdn"`
	DestinationPortListId                types.String `tfsdk:"destination_port_list_id"`
	DestinationPortListVersion           types.Int64  `tfsdk:"destination_port_list_version"`
	DestinationPort                      types.String `tfsdk:"destination_port"`
	DestinationGeoLocationListId         types.String `tfsdk:"destination_geo_location_list_id"`
	DestinationGeoLocationListVersion    types.Int64  `tfsdk:"destination_geo_location_list_version"`
	DestinationGeoLocation               types.String `tfsdk:"destination_geo_location"`
	ProtocolListId                       types.String `tfsdk:"protocol_list_id"`
	ProtocolListVersion                  types.Int64  `tfsdk:"protocol_list_version"`
	Protocol                             types.String `tfsdk:"protocol"`
	ProtocolNumber                       types.Int64  `tfsdk:"protocol_number"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data RuleSetPolicyDefinition) getPath() string {
	return "/template/policy/definition/ruleset/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data RuleSetPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "ruleSet")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "definition.rules", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if true {
				itemBody, _ = sjson.Set(itemBody, "action", "permit")
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "rule", item.Name.ValueString())
			}
			if !item.Order.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "order", fmt.Sprint(item.Order.ValueInt64()))
			}
			if !item.SourceObjectGroupId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourceSecurityGroup.ref", item.SourceObjectGroupId.ValueString())
			}
			if !item.SourceDataIpv4PrefixListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourceDataPrefixList.ref", item.SourceDataIpv4PrefixListId.ValueString())
			}
			if !item.SourceIpv4Prefix.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourceIP", item.SourceIpv4Prefix.ValueString())
			}
			if !item.SourceIpv4PrefixVariable.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourceIP.vipVariableName", item.SourceIpv4PrefixVariable.ValueString())
			}
			if !item.SourceDataFqdnPrefixListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourceFqdnList.ref", item.SourceDataFqdnPrefixListId.ValueString())
			}
			if !item.SourceFqdn.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourceFqdn", item.SourceFqdn.ValueString())
			}
			if !item.SourcePortListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourcePortList.ref", item.SourcePortListId.ValueString())
			}
			if !item.SourcePort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourcePort", item.SourcePort.ValueString())
			}
			if !item.SourceGeoLocationListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourceGeoLocationList.ref", item.SourceGeoLocationListId.ValueString())
			}
			if !item.SourceGeoLocation.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourceGeoLocation", item.SourceGeoLocation.ValueString())
			}
			if !item.DestinationObjectGroupId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destinationSecurityGroup.ref", item.DestinationObjectGroupId.ValueString())
			}
			if !item.DestinationDataIpv4PrefixListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destinationDataPrefixList.ref", item.DestinationDataIpv4PrefixListId.ValueString())
			}
			if !item.DestinationIpv4Prefix.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destinationIP", item.DestinationIpv4Prefix.ValueString())
			}
			if !item.DestinationIpv4PrefixVariable.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destinationIP.vipVariableName", item.DestinationIpv4PrefixVariable.ValueString())
			}
			if !item.DestinationDataFqdnPrefixListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destinationFqdnList.ref", item.DestinationDataFqdnPrefixListId.ValueString())
			}
			if !item.DestinationFqdn.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destinationFqdn", item.DestinationFqdn.ValueString())
			}
			if !item.DestinationPortListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destinationPortList.ref", item.DestinationPortListId.ValueString())
			}
			if !item.DestinationPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destinationPort", item.DestinationPort.ValueString())
			}
			if !item.DestinationGeoLocationListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destinationGeoLocationList.ref", item.DestinationGeoLocationListId.ValueString())
			}
			if !item.DestinationGeoLocation.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destinationGeoLocation", item.DestinationGeoLocation.ValueString())
			}
			if !item.ProtocolListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocolNameList.ref", item.ProtocolListId.ValueString())
			}
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocolName", item.Protocol.ValueString())
			}
			if !item.ProtocolNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", fmt.Sprint(item.ProtocolNumber.ValueInt64()))
			}
			body, _ = sjson.SetRaw(body, "definition.rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *RuleSetPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("definition.rules"); value.Exists() && len(value.Array()) > 0 {
		data.Rules = make([]RuleSetPolicyDefinitionRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := RuleSetPolicyDefinitionRules{}
			if cValue := v.Get("rule"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("order"); cValue.Exists() {
				item.Order = types.Int64Value(cValue.Int())
			} else {
				item.Order = types.Int64Null()
			}
			if cValue := v.Get("sourceSecurityGroup.ref"); cValue.Exists() {
				item.SourceObjectGroupId = types.StringValue(cValue.String())
			} else {
				item.SourceObjectGroupId = types.StringNull()
			}
			if cValue := v.Get("sourceDataPrefixList.ref"); cValue.Exists() {
				item.SourceDataIpv4PrefixListId = types.StringValue(cValue.String())
			} else {
				item.SourceDataIpv4PrefixListId = types.StringNull()
			}
			if cValue := v.Get("sourceIP"); cValue.Exists() {
				item.SourceIpv4Prefix = types.StringValue(cValue.String())
			} else {
				item.SourceIpv4Prefix = types.StringNull()
			}
			if cValue := v.Get("sourceIP.vipVariableName"); cValue.Exists() {
				item.SourceIpv4PrefixVariable = types.StringValue(cValue.String())
			} else {
				item.SourceIpv4PrefixVariable = types.StringNull()
			}
			if cValue := v.Get("sourceFqdnList.ref"); cValue.Exists() {
				item.SourceDataFqdnPrefixListId = types.StringValue(cValue.String())
			} else {
				item.SourceDataFqdnPrefixListId = types.StringNull()
			}
			if cValue := v.Get("sourceFqdn"); cValue.Exists() {
				item.SourceFqdn = types.StringValue(cValue.String())
			} else {
				item.SourceFqdn = types.StringNull()
			}
			if cValue := v.Get("sourcePortList.ref"); cValue.Exists() {
				item.SourcePortListId = types.StringValue(cValue.String())
			} else {
				item.SourcePortListId = types.StringNull()
			}
			if cValue := v.Get("sourcePort"); cValue.Exists() {
				item.SourcePort = types.StringValue(cValue.String())
			} else {
				item.SourcePort = types.StringNull()
			}
			if cValue := v.Get("sourceGeoLocationList.ref"); cValue.Exists() {
				item.SourceGeoLocationListId = types.StringValue(cValue.String())
			} else {
				item.SourceGeoLocationListId = types.StringNull()
			}
			if cValue := v.Get("sourceGeoLocation"); cValue.Exists() {
				item.SourceGeoLocation = types.StringValue(cValue.String())
			} else {
				item.SourceGeoLocation = types.StringNull()
			}
			if cValue := v.Get("destinationSecurityGroup.ref"); cValue.Exists() {
				item.DestinationObjectGroupId = types.StringValue(cValue.String())
			} else {
				item.DestinationObjectGroupId = types.StringNull()
			}
			if cValue := v.Get("destinationDataPrefixList.ref"); cValue.Exists() {
				item.DestinationDataIpv4PrefixListId = types.StringValue(cValue.String())
			} else {
				item.DestinationDataIpv4PrefixListId = types.StringNull()
			}
			if cValue := v.Get("destinationIP"); cValue.Exists() {
				item.DestinationIpv4Prefix = types.StringValue(cValue.String())
			} else {
				item.DestinationIpv4Prefix = types.StringNull()
			}
			if cValue := v.Get("destinationIP.vipVariableName"); cValue.Exists() {
				item.DestinationIpv4PrefixVariable = types.StringValue(cValue.String())
			} else {
				item.DestinationIpv4PrefixVariable = types.StringNull()
			}
			if cValue := v.Get("destinationFqdnList.ref"); cValue.Exists() {
				item.DestinationDataFqdnPrefixListId = types.StringValue(cValue.String())
			} else {
				item.DestinationDataFqdnPrefixListId = types.StringNull()
			}
			if cValue := v.Get("destinationFqdn"); cValue.Exists() {
				item.DestinationFqdn = types.StringValue(cValue.String())
			} else {
				item.DestinationFqdn = types.StringNull()
			}
			if cValue := v.Get("destinationPortList.ref"); cValue.Exists() {
				item.DestinationPortListId = types.StringValue(cValue.String())
			} else {
				item.DestinationPortListId = types.StringNull()
			}
			if cValue := v.Get("destinationPort"); cValue.Exists() {
				item.DestinationPort = types.StringValue(cValue.String())
			} else {
				item.DestinationPort = types.StringNull()
			}
			if cValue := v.Get("destinationGeoLocationList.ref"); cValue.Exists() {
				item.DestinationGeoLocationListId = types.StringValue(cValue.String())
			} else {
				item.DestinationGeoLocationListId = types.StringNull()
			}
			if cValue := v.Get("destinationGeoLocation"); cValue.Exists() {
				item.DestinationGeoLocation = types.StringValue(cValue.String())
			} else {
				item.DestinationGeoLocation = types.StringNull()
			}
			if cValue := v.Get("protocolNameList.ref"); cValue.Exists() {
				item.ProtocolListId = types.StringValue(cValue.String())
			} else {
				item.ProtocolListId = types.StringNull()
			}
			if cValue := v.Get("protocolName"); cValue.Exists() {
				item.Protocol = types.StringValue(cValue.String())
			} else {
				item.Protocol = types.StringNull()
			}
			if cValue := v.Get("protocol"); cValue.Exists() {
				item.ProtocolNumber = types.Int64Value(cValue.Int())
			} else {
				item.ProtocolNumber = types.Int64Null()
			}
			data.Rules = append(data.Rules, item)
			return true
		})
	} else {
		if len(data.Rules) > 0 {
			data.Rules = []RuleSetPolicyDefinitionRules{}
		}
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *RuleSetPolicyDefinition) hasChanges(ctx context.Context, state *RuleSetPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if len(data.Rules) != len(state.Rules) {
		hasChanges = true
	} else {
		for i := range data.Rules {
			if !data.Rules[i].Name.Equal(state.Rules[i].Name) {
				hasChanges = true
			}
			if !data.Rules[i].Order.Equal(state.Rules[i].Order) {
				hasChanges = true
			}
			if !data.Rules[i].SourceObjectGroupId.Equal(state.Rules[i].SourceObjectGroupId) {
				hasChanges = true
			}
			if !data.Rules[i].SourceDataIpv4PrefixListId.Equal(state.Rules[i].SourceDataIpv4PrefixListId) {
				hasChanges = true
			}
			if !data.Rules[i].SourceIpv4Prefix.Equal(state.Rules[i].SourceIpv4Prefix) {
				hasChanges = true
			}
			if !data.Rules[i].SourceIpv4PrefixVariable.Equal(state.Rules[i].SourceIpv4PrefixVariable) {
				hasChanges = true
			}
			if !data.Rules[i].SourceDataFqdnPrefixListId.Equal(state.Rules[i].SourceDataFqdnPrefixListId) {
				hasChanges = true
			}
			if !data.Rules[i].SourceFqdn.Equal(state.Rules[i].SourceFqdn) {
				hasChanges = true
			}
			if !data.Rules[i].SourcePortListId.Equal(state.Rules[i].SourcePortListId) {
				hasChanges = true
			}
			if !data.Rules[i].SourcePort.Equal(state.Rules[i].SourcePort) {
				hasChanges = true
			}
			if !data.Rules[i].SourceGeoLocationListId.Equal(state.Rules[i].SourceGeoLocationListId) {
				hasChanges = true
			}
			if !data.Rules[i].SourceGeoLocation.Equal(state.Rules[i].SourceGeoLocation) {
				hasChanges = true
			}
			if !data.Rules[i].DestinationObjectGroupId.Equal(state.Rules[i].DestinationObjectGroupId) {
				hasChanges = true
			}
			if !data.Rules[i].DestinationDataIpv4PrefixListId.Equal(state.Rules[i].DestinationDataIpv4PrefixListId) {
				hasChanges = true
			}
			if !data.Rules[i].DestinationIpv4Prefix.Equal(state.Rules[i].DestinationIpv4Prefix) {
				hasChanges = true
			}
			if !data.Rules[i].DestinationIpv4PrefixVariable.Equal(state.Rules[i].DestinationIpv4PrefixVariable) {
				hasChanges = true
			}
			if !data.Rules[i].DestinationDataFqdnPrefixListId.Equal(state.Rules[i].DestinationDataFqdnPrefixListId) {
				hasChanges = true
			}
			if !data.Rules[i].DestinationFqdn.Equal(state.Rules[i].DestinationFqdn) {
				hasChanges = true
			}
			if !data.Rules[i].DestinationPortListId.Equal(state.Rules[i].DestinationPortListId) {
				hasChanges = true
			}
			if !data.Rules[i].DestinationPort.Equal(state.Rules[i].DestinationPort) {
				hasChanges = true
			}
			if !data.Rules[i].DestinationGeoLocationListId.Equal(state.Rules[i].DestinationGeoLocationListId) {
				hasChanges = true
			}
			if !data.Rules[i].DestinationGeoLocation.Equal(state.Rules[i].DestinationGeoLocation) {
				hasChanges = true
			}
			if !data.Rules[i].ProtocolListId.Equal(state.Rules[i].ProtocolListId) {
				hasChanges = true
			}
			if !data.Rules[i].Protocol.Equal(state.Rules[i].Protocol) {
				hasChanges = true
			}
			if !data.Rules[i].ProtocolNumber.Equal(state.Rules[i].ProtocolNumber) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *RuleSetPolicyDefinition) updateVersions(ctx context.Context, state *RuleSetPolicyDefinition) {
	for i := range data.Rules {
		dataKeys := [...]string{}
		stateIndex := -1
		for j := range state.Rules {
			stateKeys := [...]string{}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		if stateIndex > -1 {
			data.Rules[i].SourceObjectGroupVersion = state.Rules[stateIndex].SourceObjectGroupVersion
		} else {
			data.Rules[i].SourceObjectGroupVersion = types.Int64Null()
		}
		if stateIndex > -1 {
			data.Rules[i].SourceDataIpv4PrefixListVersion = state.Rules[stateIndex].SourceDataIpv4PrefixListVersion
		} else {
			data.Rules[i].SourceDataIpv4PrefixListVersion = types.Int64Null()
		}
		if stateIndex > -1 {
			data.Rules[i].SourceDataFqdnPrefixListVersion = state.Rules[stateIndex].SourceDataFqdnPrefixListVersion
		} else {
			data.Rules[i].SourceDataFqdnPrefixListVersion = types.Int64Null()
		}
		if stateIndex > -1 {
			data.Rules[i].SourcePortListVersion = state.Rules[stateIndex].SourcePortListVersion
		} else {
			data.Rules[i].SourcePortListVersion = types.Int64Null()
		}
		if stateIndex > -1 {
			data.Rules[i].SourceGeoLocationListVersion = state.Rules[stateIndex].SourceGeoLocationListVersion
		} else {
			data.Rules[i].SourceGeoLocationListVersion = types.Int64Null()
		}
		if stateIndex > -1 {
			data.Rules[i].DestinationObjectGroupVersion = state.Rules[stateIndex].DestinationObjectGroupVersion
		} else {
			data.Rules[i].DestinationObjectGroupVersion = types.Int64Null()
		}
		if stateIndex > -1 {
			data.Rules[i].DestinationDataIpv4PrefixListVersion = state.Rules[stateIndex].DestinationDataIpv4PrefixListVersion
		} else {
			data.Rules[i].DestinationDataIpv4PrefixListVersion = types.Int64Null()
		}
		if stateIndex > -1 {
			data.Rules[i].DestinationDataFqdnPrefixListVersion = state.Rules[stateIndex].DestinationDataFqdnPrefixListVersion
		} else {
			data.Rules[i].DestinationDataFqdnPrefixListVersion = types.Int64Null()
		}
		if stateIndex > -1 {
			data.Rules[i].DestinationPortListVersion = state.Rules[stateIndex].DestinationPortListVersion
		} else {
			data.Rules[i].DestinationPortListVersion = types.Int64Null()
		}
		if stateIndex > -1 {
			data.Rules[i].DestinationGeoLocationListVersion = state.Rules[stateIndex].DestinationGeoLocationListVersion
		} else {
			data.Rules[i].DestinationGeoLocationListVersion = types.Int64Null()
		}
		if stateIndex > -1 {
			data.Rules[i].ProtocolListVersion = state.Rules[stateIndex].ProtocolListVersion
		} else {
			data.Rules[i].ProtocolListVersion = types.Int64Null()
		}
	}
}

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *RuleSetPolicyDefinition) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
	data.Type = types.StringValue("ruleSet")
	for i := range data.Rules {
		if data.Rules[i].SourceObjectGroupId != types.StringNull() {
			data.Rules[i].SourceObjectGroupVersion = types.Int64Value(0)
		}
		if data.Rules[i].SourceDataIpv4PrefixListId != types.StringNull() {
			data.Rules[i].SourceDataIpv4PrefixListVersion = types.Int64Value(0)
		}
		if data.Rules[i].SourceDataFqdnPrefixListId != types.StringNull() {
			data.Rules[i].SourceDataFqdnPrefixListVersion = types.Int64Value(0)
		}
		if data.Rules[i].SourcePortListId != types.StringNull() {
			data.Rules[i].SourcePortListVersion = types.Int64Value(0)
		}
		if data.Rules[i].SourceGeoLocationListId != types.StringNull() {
			data.Rules[i].SourceGeoLocationListVersion = types.Int64Value(0)
		}
		if data.Rules[i].DestinationObjectGroupId != types.StringNull() {
			data.Rules[i].DestinationObjectGroupVersion = types.Int64Value(0)
		}
		if data.Rules[i].DestinationDataIpv4PrefixListId != types.StringNull() {
			data.Rules[i].DestinationDataIpv4PrefixListVersion = types.Int64Value(0)
		}
		if data.Rules[i].DestinationDataFqdnPrefixListId != types.StringNull() {
			data.Rules[i].DestinationDataFqdnPrefixListVersion = types.Int64Value(0)
		}
		if data.Rules[i].DestinationPortListId != types.StringNull() {
			data.Rules[i].DestinationPortListVersion = types.Int64Value(0)
		}
		if data.Rules[i].DestinationGeoLocationListId != types.StringNull() {
			data.Rules[i].DestinationGeoLocationListVersion = types.Int64Value(0)
		}
		if data.Rules[i].ProtocolListId != types.StringNull() {
			data.Rules[i].ProtocolListVersion = types.Int64Value(0)
		}
	}
}

// End of section. //template:end processImport
