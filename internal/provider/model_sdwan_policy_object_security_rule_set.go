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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type PolicyObjectSecurityRuleSet struct {
	Id               types.String                           `tfsdk:"id"`
	Version          types.Int64                            `tfsdk:"version"`
	Name             types.String                           `tfsdk:"name"`
	Description      types.String                           `tfsdk:"description"`
	FeatureProfileId types.String                           `tfsdk:"feature_profile_id"`
	Sequences        []PolicyObjectSecurityRuleSetSequences `tfsdk:"sequences"`
}

type PolicyObjectSecurityRuleSetSequences struct {
	Name                             types.String `tfsdk:"name"`
	NameVariable                     types.String `tfsdk:"name_variable"`
	SequenceIpType                   types.String `tfsdk:"sequence_ip_type"`
	SequenceId                       types.String `tfsdk:"sequence_id"`
	SequenceIdVariable               types.String `tfsdk:"sequence_id_variable"`
	Action                           types.String `tfsdk:"action"`
	ActionVariable                   types.String `tfsdk:"action_variable"`
	SourceIpv4Prefixes               types.Set    `tfsdk:"source_ipv4_prefixes"`
	SourceIpv4PrefixesVariable       types.String `tfsdk:"source_ipv4_prefixes_variable"`
	SourceIpv6Prefixes               types.Set    `tfsdk:"source_ipv6_prefixes"`
	SourceIpv6PrefixesVariable       types.String `tfsdk:"source_ipv6_prefixes_variable"`
	DestinationIpv4Prefixes          types.Set    `tfsdk:"destination_ipv4_prefixes"`
	DestinationIpv4PrefixesVariable  types.String `tfsdk:"destination_ipv4_prefixes_variable"`
	DestinationIpv6Prefixes          types.Set    `tfsdk:"destination_ipv6_prefixes"`
	DestinationIpv6PrefixesVariable  types.String `tfsdk:"destination_ipv6_prefixes_variable"`
	DestinationFqdns                 types.Set    `tfsdk:"destination_fqdns"`
	DestinationFqdnsVariable         types.String `tfsdk:"destination_fqdns_variable"`
	SourceGeoLocations               types.Set    `tfsdk:"source_geo_locations"`
	SourceGeoLocationsVariable       types.String `tfsdk:"source_geo_locations_variable"`
	DestinationGeoLocations          types.Set    `tfsdk:"destination_geo_locations"`
	DestinationGeoLocationsVariable  types.String `tfsdk:"destination_geo_locations_variable"`
	SourcePorts                      types.Set    `tfsdk:"source_ports"`
	SourcePortsVariable              types.String `tfsdk:"source_ports_variable"`
	DestinationPorts                 types.Set    `tfsdk:"destination_ports"`
	DestinationPortsVariable         types.String `tfsdk:"destination_ports_variable"`
	Protocols                        types.Set    `tfsdk:"protocols"`
	ProtocolNames                    types.Set    `tfsdk:"protocol_names"`
	ProtocolNamesVariable            types.String `tfsdk:"protocol_names_variable"`
	SourceObjectGroupListIds         types.Set    `tfsdk:"source_object_group_list_ids"`
	DestinationObjectGroupListIds    types.Set    `tfsdk:"destination_object_group_list_ids"`
	SourceDataIpv4PrefixListIds      types.Set    `tfsdk:"source_data_ipv4_prefix_list_ids"`
	SourceDataIpv6PrefixListIds      types.Set    `tfsdk:"source_data_ipv6_prefix_list_ids"`
	DestinationDataIpv4PrefixListIds types.Set    `tfsdk:"destination_data_ipv4_prefix_list_ids"`
	DestinationDataIpv6PrefixListIds types.Set    `tfsdk:"destination_data_ipv6_prefix_list_ids"`
	DestinationFqdnListIds           types.Set    `tfsdk:"destination_fqdn_list_ids"`
	SourceGeoLocationListIds         types.Set    `tfsdk:"source_geo_location_list_ids"`
	DestinationGeoLocationListIds    types.Set    `tfsdk:"destination_geo_location_list_ids"`
	SourcePortListIds                types.Set    `tfsdk:"source_port_list_ids"`
	DestinationPortListIds           types.Set    `tfsdk:"destination_port_list_ids"`
	ProtocolNameListIds              types.Set    `tfsdk:"protocol_name_list_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectSecurityRuleSet) getModel() string {
	return "policy_object_security_rule_set"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectSecurityRuleSet) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/security-rule-set", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectSecurityRuleSet) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {

		for _, item := range data.Sequences {
			itemBody := ""

			if !item.NameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sequenceName.value", item.NameVariable.ValueString())
				}
			} else if !item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceName.value", item.Name.ValueString())
				}
			}
			if !item.SequenceIpType.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceIpType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceIpType.value", item.SequenceIpType.ValueString())
				}
			}

			if !item.SequenceIdVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceId.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sequenceId.value", item.SequenceIdVariable.ValueString())
				}
			} else if !item.SequenceId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceId.value", item.SequenceId.ValueString())
				}
			}

			if !item.ActionVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "action.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "action.value", item.ActionVariable.ValueString())
				}
			} else if !item.Action.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "action.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "action.value", item.Action.ValueString())
				}
			}

			if !item.SourceIpv4PrefixesVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceIp.ipv4Value.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourceIp.ipv4Value.value", item.SourceIpv4PrefixesVariable.ValueString())
				}
			} else if !item.SourceIpv4Prefixes.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceIp.ipv4Value.optionType", "global")
					var values []string
					item.SourceIpv4Prefixes.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "sourceIp.ipv4Value.value", values)
				}
			}

			if !item.SourceIpv6PrefixesVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceIpv6.ipv6Value.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourceIpv6.ipv6Value.value", item.SourceIpv6PrefixesVariable.ValueString())
				}
			} else if !item.SourceIpv6Prefixes.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceIpv6.ipv6Value.optionType", "global")
					var values []string
					item.SourceIpv6Prefixes.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "sourceIpv6.ipv6Value.value", values)
				}
			}

			if !item.DestinationIpv4PrefixesVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "destinationIp.ipv4Value.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "destinationIp.ipv4Value.value", item.DestinationIpv4PrefixesVariable.ValueString())
				}
			} else if !item.DestinationIpv4Prefixes.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "destinationIp.ipv4Value.optionType", "global")
					var values []string
					item.DestinationIpv4Prefixes.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "destinationIp.ipv4Value.value", values)
				}
			}

			if !item.DestinationIpv6PrefixesVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "destinationIpv6.ipv6Value.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "destinationIpv6.ipv6Value.value", item.DestinationIpv6PrefixesVariable.ValueString())
				}
			} else if !item.DestinationIpv6Prefixes.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "destinationIpv6.ipv6Value.optionType", "global")
					var values []string
					item.DestinationIpv6Prefixes.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "destinationIpv6.ipv6Value.value", values)
				}
			}

			if !item.DestinationFqdnsVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "destinationFqdn.fqdnValue.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "destinationFqdn.fqdnValue.value", item.DestinationFqdnsVariable.ValueString())
				}
			} else if !item.DestinationFqdns.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "destinationFqdn.fqdnValue.optionType", "global")
					var values []string
					item.DestinationFqdns.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "destinationFqdn.fqdnValue.value", values)
				}
			}

			if !item.SourceGeoLocationsVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceGeoLocation.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourceGeoLocation.value", item.SourceGeoLocationsVariable.ValueString())
				}
			} else if !item.SourceGeoLocations.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceGeoLocation.optionType", "global")
					var values []string
					item.SourceGeoLocations.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "sourceGeoLocation.value", values)
				}
			}

			if !item.DestinationGeoLocationsVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "destinationGeoLocation.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "destinationGeoLocation.value", item.DestinationGeoLocationsVariable.ValueString())
				}
			} else if !item.DestinationGeoLocations.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "destinationGeoLocation.optionType", "global")
					var values []string
					item.DestinationGeoLocations.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "destinationGeoLocation.value", values)
				}
			}

			if !item.SourcePortsVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourcePort.portValue.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourcePort.portValue.value", item.SourcePortsVariable.ValueString())
				}
			} else if !item.SourcePorts.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourcePort.portValue.optionType", "global")
					var values []string
					item.SourcePorts.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "sourcePort.portValue.value", values)
				}
			}

			if !item.DestinationPortsVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "destinationPort.portValue.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "destinationPort.portValue.value", item.DestinationPortsVariable.ValueString())
				}
			} else if !item.DestinationPorts.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "destinationPort.portValue.optionType", "global")
					var values []string
					item.DestinationPorts.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "destinationPort.portValue.value", values)
				}
			}
			if !item.Protocols.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "protocol.optionType", "global")
					var values []string
					item.Protocols.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "protocol.value", values)
				}
			}

			if !item.ProtocolNamesVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "protocolName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "protocolName.value", item.ProtocolNamesVariable.ValueString())
				}
			} else if !item.ProtocolNames.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "protocolName.optionType", "global")
					var values []string
					item.ProtocolNames.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "protocolName.value", values)
				}
			}
			if !item.SourceObjectGroupListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceObjectGroup.refId.optionType", "global")
					var values []string
					item.SourceObjectGroupListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "sourceObjectGroup.refId.value", values)
				}
			}
			if !item.DestinationObjectGroupListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "destinationObjectGroup.refId.optionType", "global")
					var values []string
					item.DestinationObjectGroupListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "destinationObjectGroup.refId.value", values)
				}
			}
			if !item.SourceDataIpv4PrefixListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceDataPrefixList.refId.optionType", "global")
					var values []string
					item.SourceDataIpv4PrefixListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "sourceDataPrefixList.refId.value", values)
				}
			}
			if !item.SourceDataIpv6PrefixListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceDataIpv6PrefixList.refId.optionType", "global")
					var values []string
					item.SourceDataIpv6PrefixListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "sourceDataIpv6PrefixList.refId.value", values)
				}
			}
			if !item.DestinationDataIpv4PrefixListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "destinationDataPrefixList.refId.optionType", "global")
					var values []string
					item.DestinationDataIpv4PrefixListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "destinationDataPrefixList.refId.value", values)
				}
			}
			if !item.DestinationDataIpv6PrefixListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "destinationDataIpv6PrefixList.refId.optionType", "global")
					var values []string
					item.DestinationDataIpv6PrefixListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "destinationDataIpv6PrefixList.refId.value", values)
				}
			}
			if !item.DestinationFqdnListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "destinationFqdnList.refId.optionType", "global")
					var values []string
					item.DestinationFqdnListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "destinationFqdnList.refId.value", values)
				}
			}
			if !item.SourceGeoLocationListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceGeoLocationList.refId.optionType", "global")
					var values []string
					item.SourceGeoLocationListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "sourceGeoLocationList.refId.value", values)
				}
			}
			if !item.DestinationGeoLocationListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "destinationGeoLocationList.refId.optionType", "global")
					var values []string
					item.DestinationGeoLocationListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "destinationGeoLocationList.refId.value", values)
				}
			}
			if !item.SourcePortListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourcePortList.refId.optionType", "global")
					var values []string
					item.SourcePortListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "sourcePortList.refId.value", values)
				}
			}
			if !item.DestinationPortListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "destinationPortList.refId.optionType", "global")
					var values []string
					item.DestinationPortListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "destinationPortList.refId.value", values)
				}
			}
			if !item.ProtocolNameListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "protocolNameList.refId.optionType", "global")
					var values []string
					item.ProtocolNameListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "protocolNameList.refId.value", values)
				}
			}
			body, _ = sjson.SetRaw(body, path+"sequences.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectSecurityRuleSet) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	oldSequences := data.Sequences
	if value := res.Get(path + "sequences"); value.Exists() && len(value.Array()) > 0 {
		data.Sequences = make([]PolicyObjectSecurityRuleSetSequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyObjectSecurityRuleSetSequences{}
			item.Name = types.StringNull()
			item.NameVariable = types.StringNull()
			if t := v.Get("sequenceName.optionType"); t.Exists() {
				va := v.Get("sequenceName.value")
				if t.String() == "variable" {
					item.NameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			item.SequenceIpType = types.StringNull()

			if t := v.Get("sequenceIpType.optionType"); t.Exists() {
				va := v.Get("sequenceIpType.value")
				if t.String() == "global" {
					item.SequenceIpType = types.StringValue(va.String())
				}
			}
			item.SequenceId = types.StringNull()
			item.SequenceIdVariable = types.StringNull()
			if t := v.Get("sequenceId.optionType"); t.Exists() {
				va := v.Get("sequenceId.value")
				if t.String() == "variable" {
					item.SequenceIdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SequenceId = types.StringValue(va.String())
				}
			}
			item.Action = types.StringNull()
			item.ActionVariable = types.StringNull()
			if t := v.Get("action.optionType"); t.Exists() {
				va := v.Get("action.value")
				if t.String() == "variable" {
					item.ActionVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Action = types.StringValue(va.String())
				}
			}
			item.SourceIpv4Prefixes = types.SetNull(types.StringType)
			item.SourceIpv4PrefixesVariable = types.StringNull()
			if t := v.Get("sourceIp.ipv4Value.optionType"); t.Exists() {
				va := v.Get("sourceIp.ipv4Value.value")
				if t.String() == "variable" {
					item.SourceIpv4PrefixesVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourceIpv4Prefixes = helpers.GetStringSet(va.Array())
				}
			}
			item.SourceIpv6Prefixes = types.SetNull(types.StringType)
			item.SourceIpv6PrefixesVariable = types.StringNull()
			if t := v.Get("sourceIpv6.ipv6Value.optionType"); t.Exists() {
				va := v.Get("sourceIpv6.ipv6Value.value")
				if t.String() == "variable" {
					item.SourceIpv6PrefixesVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourceIpv6Prefixes = helpers.GetStringSet(va.Array())
				}
			}
			item.DestinationIpv4Prefixes = types.SetNull(types.StringType)
			item.DestinationIpv4PrefixesVariable = types.StringNull()
			if t := v.Get("destinationIp.ipv4Value.optionType"); t.Exists() {
				va := v.Get("destinationIp.ipv4Value.value")
				if t.String() == "variable" {
					item.DestinationIpv4PrefixesVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.DestinationIpv4Prefixes = helpers.GetStringSet(va.Array())
				}
			}
			item.DestinationIpv6Prefixes = types.SetNull(types.StringType)
			item.DestinationIpv6PrefixesVariable = types.StringNull()
			if t := v.Get("destinationIpv6.ipv6Value.optionType"); t.Exists() {
				va := v.Get("destinationIpv6.ipv6Value.value")
				if t.String() == "variable" {
					item.DestinationIpv6PrefixesVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.DestinationIpv6Prefixes = helpers.GetStringSet(va.Array())
				}
			}
			item.DestinationFqdns = types.SetNull(types.StringType)
			item.DestinationFqdnsVariable = types.StringNull()
			if t := v.Get("destinationFqdn.fqdnValue.optionType"); t.Exists() {
				va := v.Get("destinationFqdn.fqdnValue.value")
				if t.String() == "variable" {
					item.DestinationFqdnsVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.DestinationFqdns = helpers.GetStringSet(va.Array())
				}
			}
			item.SourceGeoLocations = types.SetNull(types.StringType)
			item.SourceGeoLocationsVariable = types.StringNull()
			if t := v.Get("sourceGeoLocation.optionType"); t.Exists() {
				va := v.Get("sourceGeoLocation.value")
				if t.String() == "variable" {
					item.SourceGeoLocationsVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourceGeoLocations = helpers.GetStringSet(va.Array())
				}
			}
			item.DestinationGeoLocations = types.SetNull(types.StringType)
			item.DestinationGeoLocationsVariable = types.StringNull()
			if t := v.Get("destinationGeoLocation.optionType"); t.Exists() {
				va := v.Get("destinationGeoLocation.value")
				if t.String() == "variable" {
					item.DestinationGeoLocationsVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.DestinationGeoLocations = helpers.GetStringSet(va.Array())
				}
			}
			item.SourcePorts = types.SetNull(types.StringType)
			item.SourcePortsVariable = types.StringNull()
			if t := v.Get("sourcePort.portValue.optionType"); t.Exists() {
				va := v.Get("sourcePort.portValue.value")
				if t.String() == "variable" {
					item.SourcePortsVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourcePorts = helpers.GetStringSet(va.Array())
				}
			}
			item.DestinationPorts = types.SetNull(types.StringType)
			item.DestinationPortsVariable = types.StringNull()
			if t := v.Get("destinationPort.portValue.optionType"); t.Exists() {
				va := v.Get("destinationPort.portValue.value")
				if t.String() == "variable" {
					item.DestinationPortsVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.DestinationPorts = helpers.GetStringSet(va.Array())
				}
			}
			item.Protocols = types.SetNull(types.StringType)

			if t := v.Get("protocol.optionType"); t.Exists() {
				va := v.Get("protocol.value")
				if t.String() == "global" {
					item.Protocols = helpers.GetStringSet(va.Array())
				}
			}
			item.ProtocolNames = types.SetNull(types.StringType)
			item.ProtocolNamesVariable = types.StringNull()
			if t := v.Get("protocolName.optionType"); t.Exists() {
				va := v.Get("protocolName.value")
				if t.String() == "variable" {
					item.ProtocolNamesVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.ProtocolNames = helpers.GetStringSet(va.Array())
				}
			}
			item.SourceObjectGroupListIds = types.SetNull(types.StringType)

			if t := v.Get("sourceObjectGroup.refId.optionType"); t.Exists() {
				va := v.Get("sourceObjectGroup.refId.value")
				if t.String() == "global" {
					item.SourceObjectGroupListIds = helpers.GetStringSet(va.Array())
				}
			}
			item.DestinationObjectGroupListIds = types.SetNull(types.StringType)

			if t := v.Get("destinationObjectGroup.refId.optionType"); t.Exists() {
				va := v.Get("destinationObjectGroup.refId.value")
				if t.String() == "global" {
					item.DestinationObjectGroupListIds = helpers.GetStringSet(va.Array())
				}
			}
			item.SourceDataIpv4PrefixListIds = types.SetNull(types.StringType)

			if t := v.Get("sourceDataPrefixList.refId.optionType"); t.Exists() {
				va := v.Get("sourceDataPrefixList.refId.value")
				if t.String() == "global" {
					item.SourceDataIpv4PrefixListIds = helpers.GetStringSet(va.Array())
				}
			}
			item.SourceDataIpv6PrefixListIds = types.SetNull(types.StringType)

			if t := v.Get("sourceDataIpv6PrefixList.refId.optionType"); t.Exists() {
				va := v.Get("sourceDataIpv6PrefixList.refId.value")
				if t.String() == "global" {
					item.SourceDataIpv6PrefixListIds = helpers.GetStringSet(va.Array())
				}
			}
			item.DestinationDataIpv4PrefixListIds = types.SetNull(types.StringType)

			if t := v.Get("destinationDataPrefixList.refId.optionType"); t.Exists() {
				va := v.Get("destinationDataPrefixList.refId.value")
				if t.String() == "global" {
					item.DestinationDataIpv4PrefixListIds = helpers.GetStringSet(va.Array())
				}
			}
			item.DestinationDataIpv6PrefixListIds = types.SetNull(types.StringType)

			if t := v.Get("destinationDataIpv6PrefixList.refId.optionType"); t.Exists() {
				va := v.Get("destinationDataIpv6PrefixList.refId.value")
				if t.String() == "global" {
					item.DestinationDataIpv6PrefixListIds = helpers.GetStringSet(va.Array())
				}
			}
			item.DestinationFqdnListIds = types.SetNull(types.StringType)

			if t := v.Get("destinationFqdnList.refId.optionType"); t.Exists() {
				va := v.Get("destinationFqdnList.refId.value")
				if t.String() == "global" {
					item.DestinationFqdnListIds = helpers.GetStringSet(va.Array())
				}
			}
			item.SourceGeoLocationListIds = types.SetNull(types.StringType)

			if t := v.Get("sourceGeoLocationList.refId.optionType"); t.Exists() {
				va := v.Get("sourceGeoLocationList.refId.value")
				if t.String() == "global" {
					item.SourceGeoLocationListIds = helpers.GetStringSet(va.Array())
				}
			}
			item.DestinationGeoLocationListIds = types.SetNull(types.StringType)

			if t := v.Get("destinationGeoLocationList.refId.optionType"); t.Exists() {
				va := v.Get("destinationGeoLocationList.refId.value")
				if t.String() == "global" {
					item.DestinationGeoLocationListIds = helpers.GetStringSet(va.Array())
				}
			}
			item.SourcePortListIds = types.SetNull(types.StringType)

			if t := v.Get("sourcePortList.refId.optionType"); t.Exists() {
				va := v.Get("sourcePortList.refId.value")
				if t.String() == "global" {
					item.SourcePortListIds = helpers.GetStringSet(va.Array())
				}
			}
			item.DestinationPortListIds = types.SetNull(types.StringType)

			if t := v.Get("destinationPortList.refId.optionType"); t.Exists() {
				va := v.Get("destinationPortList.refId.value")
				if t.String() == "global" {
					item.DestinationPortListIds = helpers.GetStringSet(va.Array())
				}
			}
			item.ProtocolNameListIds = types.SetNull(types.StringType)

			if t := v.Get("protocolNameList.refId.optionType"); t.Exists() {
				va := v.Get("protocolNameList.refId.value")
				if t.String() == "global" {
					item.ProtocolNameListIds = helpers.GetStringSet(va.Array())
				}
			}
			data.Sequences = append(data.Sequences, item)
			return true
		})
	} else {
		data.Sequences = nil
	}
	if !fullRead && data.Sequences != nil {
		resultSequences := make([]PolicyObjectSecurityRuleSetSequences, 0, len(data.Sequences))
		matchedSequences := make([]bool, len(data.Sequences))
		for _, oldItem := range oldSequences {
			for ni := range data.Sequences {
				if matchedSequences[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.NameVariable.ValueString() != "" || data.Sequences[ni].NameVariable.ValueString() != "") {
					if oldItem.NameVariable.ValueString() != data.Sequences[ni].NameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.Name.ValueString() != data.Sequences[ni].Name.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if oldItem.SequenceIpType.ValueString() != data.Sequences[ni].SequenceIpType.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch && (oldItem.SequenceIdVariable.ValueString() != "" || data.Sequences[ni].SequenceIdVariable.ValueString() != "") {
					if oldItem.SequenceIdVariable.ValueString() != data.Sequences[ni].SequenceIdVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.SequenceId.ValueString() != data.Sequences[ni].SequenceId.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch && (oldItem.ActionVariable.ValueString() != "" || data.Sequences[ni].ActionVariable.ValueString() != "") {
					if oldItem.ActionVariable.ValueString() != data.Sequences[ni].ActionVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.Action.ValueString() != data.Sequences[ni].Action.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.SourceIpv4Prefixes).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].SourceIpv4Prefixes).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.SourceIpv6Prefixes).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].SourceIpv6Prefixes).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.DestinationIpv4Prefixes).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].DestinationIpv4Prefixes).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.DestinationIpv6Prefixes).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].DestinationIpv6Prefixes).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.DestinationFqdns).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].DestinationFqdns).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.SourceGeoLocations).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].SourceGeoLocations).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.DestinationGeoLocations).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].DestinationGeoLocations).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.SourcePorts).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].SourcePorts).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.DestinationPorts).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].DestinationPorts).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.Protocols).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].Protocols).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.ProtocolNames).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].ProtocolNames).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.SourceObjectGroupListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].SourceObjectGroupListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.DestinationObjectGroupListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].DestinationObjectGroupListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.SourceDataIpv4PrefixListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].SourceDataIpv4PrefixListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.SourceDataIpv6PrefixListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].SourceDataIpv6PrefixListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.DestinationDataIpv4PrefixListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].DestinationDataIpv4PrefixListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.DestinationDataIpv6PrefixListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].DestinationDataIpv6PrefixListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.DestinationFqdnListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].DestinationFqdnListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.SourceGeoLocationListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].SourceGeoLocationListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.DestinationGeoLocationListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].DestinationGeoLocationListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.SourcePortListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].SourcePortListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.DestinationPortListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].DestinationPortListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.ProtocolNameListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].ProtocolNameListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedSequences[ni] = true
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
