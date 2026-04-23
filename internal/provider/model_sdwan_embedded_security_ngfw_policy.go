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
type EmbeddedSecurityNGFW struct {
	Id               types.String                    `tfsdk:"id"`
	Version          types.Int64                     `tfsdk:"version"`
	Name             types.String                    `tfsdk:"name"`
	Description      types.String                    `tfsdk:"description"`
	FeatureProfileId types.String                    `tfsdk:"feature_profile_id"`
	DefaultAction    types.String                    `tfsdk:"default_action"`
	Sequences        []EmbeddedSecurityNGFWSequences `tfsdk:"sequences"`
}

type EmbeddedSecurityNGFWSequences struct {
	SequenceId      types.String                                `tfsdk:"sequence_id"`
	SequenceName    types.String                                `tfsdk:"sequence_name"`
	BaseAction      types.String                                `tfsdk:"base_action"`
	SequenceType    types.String                                `tfsdk:"sequence_type"`
	DisableSequence types.Bool                                  `tfsdk:"disable_sequence"`
	MatchEntries    []EmbeddedSecurityNGFWSequencesMatchEntries `tfsdk:"match_entries"`
	Actions         []EmbeddedSecurityNGFWSequencesActions      `tfsdk:"actions"`
}

type EmbeddedSecurityNGFWSequencesMatchEntries struct {
	SourceDataPrefixListIds            types.Set    `tfsdk:"source_data_prefix_list_ids"`
	DestinationDataPrefixListIds       types.Set    `tfsdk:"destination_data_prefix_list_ids"`
	DestinationFqdnListIds             types.Set    `tfsdk:"destination_fqdn_list_ids"`
	SourceGeoLocationListIds           types.Set    `tfsdk:"source_geo_location_list_ids"`
	DestinationGeoLocationListIds      types.Set    `tfsdk:"destination_geo_location_list_ids"`
	SourcePortListIds                  types.Set    `tfsdk:"source_port_list_ids"`
	DestinationPortListIds             types.Set    `tfsdk:"destination_port_list_ids"`
	SourceScalableGroupTagListIds      types.Set    `tfsdk:"source_scalable_group_tag_list_ids"`
	DestinationScalableGroupTagListIds types.Set    `tfsdk:"destination_scalable_group_tag_list_ids"`
	SourceIdentityListIds              types.Set    `tfsdk:"source_identity_list_ids"`
	ProtocolNameListIds                types.Set    `tfsdk:"protocol_name_list_ids"`
	AppListIds                         types.Set    `tfsdk:"app_list_ids"`
	FlatAppListIds                     types.Set    `tfsdk:"flat_app_list_ids"`
	SourceSecurityGroupListIds         types.Set    `tfsdk:"source_security_group_list_ids"`
	DestinationSecurityGroupListIds    types.Set    `tfsdk:"destination_security_group_list_ids"`
	SourceDataPrefixes                 types.Set    `tfsdk:"source_data_prefixes"`
	SourceDataPrefixesVariable         types.String `tfsdk:"source_data_prefixes_variable"`
	DestinationDataPrefixes            types.Set    `tfsdk:"destination_data_prefixes"`
	DestinationDataPrefixesVariable    types.String `tfsdk:"destination_data_prefixes_variable"`
	DestinationFqdns                   types.Set    `tfsdk:"destination_fqdns"`
	DestinationFqdnsVariable           types.String `tfsdk:"destination_fqdns_variable"`
	SourcePorts                        types.Set    `tfsdk:"source_ports"`
	SourcePortsVariable                types.String `tfsdk:"source_ports_variable"`
	DestinationPorts                   types.Set    `tfsdk:"destination_ports"`
	DestinationPortsVariable           types.String `tfsdk:"destination_ports_variable"`
	SourceGeoLocations                 types.Set    `tfsdk:"source_geo_locations"`
	SourceGeoLocationsVariable         types.String `tfsdk:"source_geo_locations_variable"`
	DestinationGeoLocations            types.Set    `tfsdk:"destination_geo_locations"`
	DestinationGeoLocationsVariable    types.String `tfsdk:"destination_geo_locations_variable"`
	SourceIdentityUsers                types.Set    `tfsdk:"source_identity_users"`
	SourceIdentityUsergroups           types.Set    `tfsdk:"source_identity_usergroups"`
	Applications                       types.Set    `tfsdk:"applications"`
	ApplicationFamilies                types.Set    `tfsdk:"application_families"`
	Protocols                          types.Set    `tfsdk:"protocols"`
	ProtocolNames                      types.Set    `tfsdk:"protocol_names"`
}
type EmbeddedSecurityNGFWSequencesActions struct {
	Type        types.String `tfsdk:"type"`
	Parameter   types.String `tfsdk:"parameter"`
	ParameterId types.String `tfsdk:"parameter_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data EmbeddedSecurityNGFW) getModel() string {
	return "embedded_security_ngfw"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data EmbeddedSecurityNGFW) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/embedded-security/%v/unified/ngfirewall", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data EmbeddedSecurityNGFW) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.DefaultAction.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"defaultActionType.optionType", "global")
			body, _ = sjson.Set(body, path+"defaultActionType.value", data.DefaultAction.ValueString())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"sequences", []interface{}{})
		for _, item := range data.Sequences {
			itemBody := ""
			if !item.SequenceId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceId.value", item.SequenceId.ValueString())
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
			if !item.SequenceType.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceType.value", item.SequenceType.ValueString())
				}
			}
			if !item.DisableSequence.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "disableSequence.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "disableSequence.value", item.DisableSequence.ValueBool())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "match.entries", []interface{}{})
				for _, childItem := range item.MatchEntries {
					itemChildBody := ""
					if !childItem.SourceDataPrefixListIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataPrefixList.refId.optionType", "global")
							var values []string
							childItem.SourceDataPrefixListIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceDataPrefixList.refId.value", values)
						}
					}
					if !childItem.DestinationDataPrefixListIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationDataPrefixList.refId.optionType", "global")
							var values []string
							childItem.DestinationDataPrefixListIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationDataPrefixList.refId.value", values)
						}
					}
					if !childItem.DestinationFqdnListIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationFqdnList.refId.optionType", "global")
							var values []string
							childItem.DestinationFqdnListIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationFqdnList.refId.value", values)
						}
					}
					if !childItem.SourceGeoLocationListIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceGeoLocationList.refId.optionType", "global")
							var values []string
							childItem.SourceGeoLocationListIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceGeoLocationList.refId.value", values)
						}
					}
					if !childItem.DestinationGeoLocationListIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationGeoLocationList.refId.optionType", "global")
							var values []string
							childItem.DestinationGeoLocationListIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationGeoLocationList.refId.value", values)
						}
					}
					if !childItem.SourcePortListIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourcePortList.refId.optionType", "global")
							var values []string
							childItem.SourcePortListIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "sourcePortList.refId.value", values)
						}
					}
					if !childItem.DestinationPortListIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationPortList.refId.optionType", "global")
							var values []string
							childItem.DestinationPortListIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationPortList.refId.value", values)
						}
					}
					if !childItem.SourceScalableGroupTagListIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceScalableGroupTagList.refId.optionType", "global")
							var values []string
							childItem.SourceScalableGroupTagListIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceScalableGroupTagList.refId.value", values)
						}
					}
					if !childItem.DestinationScalableGroupTagListIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationScalableGroupTagList.refId.optionType", "global")
							var values []string
							childItem.DestinationScalableGroupTagListIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationScalableGroupTagList.refId.value", values)
						}
					}
					if !childItem.SourceIdentityListIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceIdentityList.refId.optionType", "global")
							var values []string
							childItem.SourceIdentityListIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceIdentityList.refId.value", values)
						}
					}
					if !childItem.ProtocolNameListIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "protocolNameList.refId.optionType", "global")
							var values []string
							childItem.ProtocolNameListIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "protocolNameList.refId.value", values)
						}
					}
					if !childItem.AppListIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "appList.refId.optionType", "global")
							var values []string
							childItem.AppListIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "appList.refId.value", values)
						}
					}
					if !childItem.FlatAppListIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "appListFlat.refId.optionType", "global")
							var values []string
							childItem.FlatAppListIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "appListFlat.refId.value", values)
						}
					}
					if !childItem.SourceSecurityGroupListIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceSecurityGroup.refId.optionType", "global")
							var values []string
							childItem.SourceSecurityGroupListIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceSecurityGroup.refId.value", values)
						}
					}
					if !childItem.DestinationSecurityGroupListIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationSecurityGroup.refId.optionType", "global")
							var values []string
							childItem.DestinationSecurityGroupListIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationSecurityGroup.refId.value", values)
						}
					}

					if !childItem.SourceDataPrefixesVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceIp.ipv4Value.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceIp.ipv4Value.value", childItem.SourceDataPrefixesVariable.ValueString())
						}
					} else if !childItem.SourceDataPrefixes.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceIp.ipv4Value.optionType", "global")
							var values []string
							childItem.SourceDataPrefixes.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceIp.ipv4Value.value", values)
						}
					}

					if !childItem.DestinationDataPrefixesVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationIp.ipv4Value.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationIp.ipv4Value.value", childItem.DestinationDataPrefixesVariable.ValueString())
						}
					} else if !childItem.DestinationDataPrefixes.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationIp.ipv4Value.optionType", "global")
							var values []string
							childItem.DestinationDataPrefixes.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationIp.ipv4Value.value", values)
						}
					}

					if !childItem.DestinationFqdnsVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationFqdn.fqdnValue.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationFqdn.fqdnValue.value", childItem.DestinationFqdnsVariable.ValueString())
						}
					} else if !childItem.DestinationFqdns.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationFqdn.fqdnValue.optionType", "global")
							var values []string
							childItem.DestinationFqdns.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationFqdn.fqdnValue.value", values)
						}
					}

					if !childItem.SourcePortsVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourcePort.portValue.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "sourcePort.portValue.value", childItem.SourcePortsVariable.ValueString())
						}
					} else if !childItem.SourcePorts.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourcePort.portValue.optionType", "global")
							var values []string
							childItem.SourcePorts.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "sourcePort.portValue.value", values)
						}
					}

					if !childItem.DestinationPortsVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationPort.portValue.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationPort.portValue.value", childItem.DestinationPortsVariable.ValueString())
						}
					} else if !childItem.DestinationPorts.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationPort.portValue.optionType", "global")
							var values []string
							childItem.DestinationPorts.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationPort.portValue.value", values)
						}
					}

					if !childItem.SourceGeoLocationsVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceGeoLocation.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceGeoLocation.value", childItem.SourceGeoLocationsVariable.ValueString())
						}
					} else if !childItem.SourceGeoLocations.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceGeoLocation.optionType", "global")
							var values []string
							childItem.SourceGeoLocations.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceGeoLocation.value", values)
						}
					}

					if !childItem.DestinationGeoLocationsVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationGeoLocation.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationGeoLocation.value", childItem.DestinationGeoLocationsVariable.ValueString())
						}
					} else if !childItem.DestinationGeoLocations.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationGeoLocation.optionType", "global")
							var values []string
							childItem.DestinationGeoLocations.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "destinationGeoLocation.value", values)
						}
					}
					if !childItem.SourceIdentityUsers.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceIdentityUser.optionType", "global")
							var values []string
							childItem.SourceIdentityUsers.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceIdentityUser.value", values)
						}
					}
					if !childItem.SourceIdentityUsergroups.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceIdentityUsergroup.optionType", "global")
							var values []string
							childItem.SourceIdentityUsergroups.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "sourceIdentityUsergroup.value", values)
						}
					}
					if !childItem.Applications.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "app.optionType", "global")
							var values []string
							childItem.Applications.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "app.value", values)
						}
					}
					if !childItem.ApplicationFamilies.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "appFamily.optionType", "global")
							var values []string
							childItem.ApplicationFamilies.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "appFamily.value", values)
						}
					}
					if !childItem.Protocols.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.optionType", "global")
							var values []string
							childItem.Protocols.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.value", values)
						}
					}
					if !childItem.ProtocolNames.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "protocolName.optionType", "global")
							var values []string
							childItem.ProtocolNames.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "protocolName.value", values)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "match.entries.-1", itemChildBody)
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "actions", []interface{}{})
				for _, childItem := range item.Actions {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "type.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "type.value", childItem.Type.ValueString())
						}
					}
					if !childItem.Parameter.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "parameter.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "parameter.value", childItem.Parameter.ValueString())
						}
					}
					if !childItem.ParameterId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "parameter.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "parameter.refId.value", childItem.ParameterId.ValueString())
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
func (data *EmbeddedSecurityNGFW) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.DefaultAction = types.StringNull()

	if t := res.Get(path + "defaultActionType.optionType"); t.Exists() {
		va := res.Get(path + "defaultActionType.value")
		if t.String() == "global" {
			data.DefaultAction = types.StringValue(va.String())
		}
	}
	oldSequences := data.Sequences
	if value := res.Get(path + "sequences"); value.Exists() && len(value.Array()) > 0 {
		data.Sequences = make([]EmbeddedSecurityNGFWSequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := EmbeddedSecurityNGFWSequences{}
			item.SequenceId = types.StringNull()

			if t := v.Get("sequenceId.optionType"); t.Exists() {
				va := v.Get("sequenceId.value")
				if t.String() == "global" {
					item.SequenceId = types.StringValue(va.String())
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
			item.SequenceType = types.StringNull()

			if t := v.Get("sequenceType.optionType"); t.Exists() {
				va := v.Get("sequenceType.value")
				if t.String() == "global" {
					item.SequenceType = types.StringValue(va.String())
				}
			}
			item.DisableSequence = types.BoolNull()

			if t := v.Get("disableSequence.optionType"); t.Exists() {
				va := v.Get("disableSequence.value")
				if t.String() == "global" {
					item.DisableSequence = types.BoolValue(va.Bool())
				}
			}
			if cValue := v.Get("match.entries"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.MatchEntries = make([]EmbeddedSecurityNGFWSequencesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := EmbeddedSecurityNGFWSequencesMatchEntries{}
					cItem.SourceDataPrefixListIds = types.SetNull(types.StringType)

					if t := cv.Get("sourceDataPrefixList.refId.optionType"); t.Exists() {
						va := cv.Get("sourceDataPrefixList.refId.value")
						if t.String() == "global" {
							cItem.SourceDataPrefixListIds = helpers.GetStringSet(va.Array())
						}
					}
					cItem.DestinationDataPrefixListIds = types.SetNull(types.StringType)

					if t := cv.Get("destinationDataPrefixList.refId.optionType"); t.Exists() {
						va := cv.Get("destinationDataPrefixList.refId.value")
						if t.String() == "global" {
							cItem.DestinationDataPrefixListIds = helpers.GetStringSet(va.Array())
						}
					}
					cItem.DestinationFqdnListIds = types.SetNull(types.StringType)

					if t := cv.Get("destinationFqdnList.refId.optionType"); t.Exists() {
						va := cv.Get("destinationFqdnList.refId.value")
						if t.String() == "global" {
							cItem.DestinationFqdnListIds = helpers.GetStringSet(va.Array())
						}
					}
					cItem.SourceGeoLocationListIds = types.SetNull(types.StringType)

					if t := cv.Get("sourceGeoLocationList.refId.optionType"); t.Exists() {
						va := cv.Get("sourceGeoLocationList.refId.value")
						if t.String() == "global" {
							cItem.SourceGeoLocationListIds = helpers.GetStringSet(va.Array())
						}
					}
					cItem.DestinationGeoLocationListIds = types.SetNull(types.StringType)

					if t := cv.Get("destinationGeoLocationList.refId.optionType"); t.Exists() {
						va := cv.Get("destinationGeoLocationList.refId.value")
						if t.String() == "global" {
							cItem.DestinationGeoLocationListIds = helpers.GetStringSet(va.Array())
						}
					}
					cItem.SourcePortListIds = types.SetNull(types.StringType)

					if t := cv.Get("sourcePortList.refId.optionType"); t.Exists() {
						va := cv.Get("sourcePortList.refId.value")
						if t.String() == "global" {
							cItem.SourcePortListIds = helpers.GetStringSet(va.Array())
						}
					}
					cItem.DestinationPortListIds = types.SetNull(types.StringType)

					if t := cv.Get("destinationPortList.refId.optionType"); t.Exists() {
						va := cv.Get("destinationPortList.refId.value")
						if t.String() == "global" {
							cItem.DestinationPortListIds = helpers.GetStringSet(va.Array())
						}
					}
					cItem.SourceScalableGroupTagListIds = types.SetNull(types.StringType)

					if t := cv.Get("sourceScalableGroupTagList.refId.optionType"); t.Exists() {
						va := cv.Get("sourceScalableGroupTagList.refId.value")
						if t.String() == "global" {
							cItem.SourceScalableGroupTagListIds = helpers.GetStringSet(va.Array())
						}
					}
					cItem.DestinationScalableGroupTagListIds = types.SetNull(types.StringType)

					if t := cv.Get("destinationScalableGroupTagList.refId.optionType"); t.Exists() {
						va := cv.Get("destinationScalableGroupTagList.refId.value")
						if t.String() == "global" {
							cItem.DestinationScalableGroupTagListIds = helpers.GetStringSet(va.Array())
						}
					}
					cItem.SourceIdentityListIds = types.SetNull(types.StringType)

					if t := cv.Get("sourceIdentityList.refId.optionType"); t.Exists() {
						va := cv.Get("sourceIdentityList.refId.value")
						if t.String() == "global" {
							cItem.SourceIdentityListIds = helpers.GetStringSet(va.Array())
						}
					}
					cItem.ProtocolNameListIds = types.SetNull(types.StringType)

					if t := cv.Get("protocolNameList.refId.optionType"); t.Exists() {
						va := cv.Get("protocolNameList.refId.value")
						if t.String() == "global" {
							cItem.ProtocolNameListIds = helpers.GetStringSet(va.Array())
						}
					}
					cItem.AppListIds = types.SetNull(types.StringType)

					if t := cv.Get("appList.refId.optionType"); t.Exists() {
						va := cv.Get("appList.refId.value")
						if t.String() == "global" {
							cItem.AppListIds = helpers.GetStringSet(va.Array())
						}
					}
					cItem.FlatAppListIds = types.SetNull(types.StringType)

					if t := cv.Get("appListFlat.refId.optionType"); t.Exists() {
						va := cv.Get("appListFlat.refId.value")
						if t.String() == "global" {
							cItem.FlatAppListIds = helpers.GetStringSet(va.Array())
						}
					}
					cItem.SourceSecurityGroupListIds = types.SetNull(types.StringType)

					if t := cv.Get("sourceSecurityGroup.refId.optionType"); t.Exists() {
						va := cv.Get("sourceSecurityGroup.refId.value")
						if t.String() == "global" {
							cItem.SourceSecurityGroupListIds = helpers.GetStringSet(va.Array())
						}
					}
					cItem.DestinationSecurityGroupListIds = types.SetNull(types.StringType)

					if t := cv.Get("destinationSecurityGroup.refId.optionType"); t.Exists() {
						va := cv.Get("destinationSecurityGroup.refId.value")
						if t.String() == "global" {
							cItem.DestinationSecurityGroupListIds = helpers.GetStringSet(va.Array())
						}
					}
					cItem.SourceDataPrefixes = types.SetNull(types.StringType)
					cItem.SourceDataPrefixesVariable = types.StringNull()
					if t := cv.Get("sourceIp.ipv4Value.optionType"); t.Exists() {
						va := cv.Get("sourceIp.ipv4Value.value")
						if t.String() == "variable" {
							cItem.SourceDataPrefixesVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.SourceDataPrefixes = helpers.GetStringSet(va.Array())
						}
					}
					cItem.DestinationDataPrefixes = types.SetNull(types.StringType)
					cItem.DestinationDataPrefixesVariable = types.StringNull()
					if t := cv.Get("destinationIp.ipv4Value.optionType"); t.Exists() {
						va := cv.Get("destinationIp.ipv4Value.value")
						if t.String() == "variable" {
							cItem.DestinationDataPrefixesVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.DestinationDataPrefixes = helpers.GetStringSet(va.Array())
						}
					}
					cItem.DestinationFqdns = types.SetNull(types.StringType)
					cItem.DestinationFqdnsVariable = types.StringNull()
					if t := cv.Get("destinationFqdn.fqdnValue.optionType"); t.Exists() {
						va := cv.Get("destinationFqdn.fqdnValue.value")
						if t.String() == "variable" {
							cItem.DestinationFqdnsVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.DestinationFqdns = helpers.GetStringSet(va.Array())
						}
					}
					cItem.SourcePorts = types.SetNull(types.StringType)
					cItem.SourcePortsVariable = types.StringNull()
					if t := cv.Get("sourcePort.portValue.optionType"); t.Exists() {
						va := cv.Get("sourcePort.portValue.value")
						if t.String() == "variable" {
							cItem.SourcePortsVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.SourcePorts = helpers.GetStringSet(va.Array())
						}
					}
					cItem.DestinationPorts = types.SetNull(types.StringType)
					cItem.DestinationPortsVariable = types.StringNull()
					if t := cv.Get("destinationPort.portValue.optionType"); t.Exists() {
						va := cv.Get("destinationPort.portValue.value")
						if t.String() == "variable" {
							cItem.DestinationPortsVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.DestinationPorts = helpers.GetStringSet(va.Array())
						}
					}
					cItem.SourceGeoLocations = types.SetNull(types.StringType)
					cItem.SourceGeoLocationsVariable = types.StringNull()
					if t := cv.Get("sourceGeoLocation.optionType"); t.Exists() {
						va := cv.Get("sourceGeoLocation.value")
						if t.String() == "variable" {
							cItem.SourceGeoLocationsVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.SourceGeoLocations = helpers.GetStringSet(va.Array())
						}
					}
					cItem.DestinationGeoLocations = types.SetNull(types.StringType)
					cItem.DestinationGeoLocationsVariable = types.StringNull()
					if t := cv.Get("destinationGeoLocation.optionType"); t.Exists() {
						va := cv.Get("destinationGeoLocation.value")
						if t.String() == "variable" {
							cItem.DestinationGeoLocationsVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.DestinationGeoLocations = helpers.GetStringSet(va.Array())
						}
					}
					cItem.SourceIdentityUsers = types.SetNull(types.StringType)

					if t := cv.Get("sourceIdentityUser.optionType"); t.Exists() {
						va := cv.Get("sourceIdentityUser.value")
						if t.String() == "global" {
							cItem.SourceIdentityUsers = helpers.GetStringSet(va.Array())
						}
					}
					cItem.SourceIdentityUsergroups = types.SetNull(types.StringType)

					if t := cv.Get("sourceIdentityUsergroup.optionType"); t.Exists() {
						va := cv.Get("sourceIdentityUsergroup.value")
						if t.String() == "global" {
							cItem.SourceIdentityUsergroups = helpers.GetStringSet(va.Array())
						}
					}
					cItem.Applications = types.SetNull(types.StringType)

					if t := cv.Get("app.optionType"); t.Exists() {
						va := cv.Get("app.value")
						if t.String() == "global" {
							cItem.Applications = helpers.GetStringSet(va.Array())
						}
					}
					cItem.ApplicationFamilies = types.SetNull(types.StringType)

					if t := cv.Get("appFamily.optionType"); t.Exists() {
						va := cv.Get("appFamily.value")
						if t.String() == "global" {
							cItem.ApplicationFamilies = helpers.GetStringSet(va.Array())
						}
					}
					cItem.Protocols = types.SetNull(types.StringType)

					if t := cv.Get("protocol.optionType"); t.Exists() {
						va := cv.Get("protocol.value")
						if t.String() == "global" {
							cItem.Protocols = helpers.GetStringSet(va.Array())
						}
					}
					cItem.ProtocolNames = types.SetNull(types.StringType)

					if t := cv.Get("protocolName.optionType"); t.Exists() {
						va := cv.Get("protocolName.value")
						if t.String() == "global" {
							cItem.ProtocolNames = helpers.GetStringSet(va.Array())
						}
					}
					item.MatchEntries = append(item.MatchEntries, cItem)
					return true
				})
			}
			if cValue := v.Get("actions"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Actions = make([]EmbeddedSecurityNGFWSequencesActions, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := EmbeddedSecurityNGFWSequencesActions{}
					cItem.Type = types.StringNull()

					if t := cv.Get("type.optionType"); t.Exists() {
						va := cv.Get("type.value")
						if t.String() == "global" {
							cItem.Type = types.StringValue(va.String())
						}
					}
					cItem.Parameter = types.StringNull()

					if t := cv.Get("parameter.optionType"); t.Exists() {
						va := cv.Get("parameter.value")
						if t.String() == "global" {
							cItem.Parameter = types.StringValue(va.String())
						}
					}
					cItem.ParameterId = types.StringNull()

					if t := cv.Get("parameter.refId.optionType"); t.Exists() {
						va := cv.Get("parameter.refId.value")
						if t.String() == "global" {
							cItem.ParameterId = types.StringValue(va.String())
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
		resultSequences := make([]EmbeddedSecurityNGFWSequences, 0, len(data.Sequences))
		matchedSequences := make([]bool, len(data.Sequences))
		for _, oldItem := range oldSequences {
			for ni := range data.Sequences {
				if matchedSequences[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.SequenceId.ValueString() != data.Sequences[ni].SequenceId.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if oldItem.SequenceName.ValueString() != data.Sequences[ni].SequenceName.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if oldItem.BaseAction.ValueString() != data.Sequences[ni].BaseAction.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if oldItem.SequenceType.ValueString() != data.Sequences[ni].SequenceType.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if oldItem.DisableSequence.ValueBool() != data.Sequences[ni].DisableSequence.ValueBool() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedSequences[ni] = true
					{
						resultC := make([]EmbeddedSecurityNGFWSequencesMatchEntries, 0, len(data.Sequences[ni].MatchEntries))
						matchedC := make([]bool, len(data.Sequences[ni].MatchEntries))
						for _, oldCItem := range oldItem.MatchEntries {
							for nci := range data.Sequences[ni].MatchEntries {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.SourceDataPrefixListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].SourceDataPrefixListIds).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.DestinationDataPrefixListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].DestinationDataPrefixListIds).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.DestinationFqdnListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].DestinationFqdnListIds).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.SourceGeoLocationListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].SourceGeoLocationListIds).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.DestinationGeoLocationListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].DestinationGeoLocationListIds).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.SourcePortListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].SourcePortListIds).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.DestinationPortListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].DestinationPortListIds).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.SourceScalableGroupTagListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].SourceScalableGroupTagListIds).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.DestinationScalableGroupTagListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].DestinationScalableGroupTagListIds).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.SourceIdentityListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].SourceIdentityListIds).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.ProtocolNameListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].ProtocolNameListIds).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.AppListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].AppListIds).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.FlatAppListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].FlatAppListIds).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.SourceSecurityGroupListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].SourceSecurityGroupListIds).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.DestinationSecurityGroupListIds).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].DestinationSecurityGroupListIds).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.SourceDataPrefixes).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].SourceDataPrefixes).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.DestinationDataPrefixes).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].DestinationDataPrefixes).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.DestinationFqdns).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].DestinationFqdns).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.SourcePorts).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].SourcePorts).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.DestinationPorts).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].DestinationPorts).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.SourceGeoLocations).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].SourceGeoLocations).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.DestinationGeoLocations).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].DestinationGeoLocations).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.SourceIdentityUsers).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].SourceIdentityUsers).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.SourceIdentityUsergroups).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].SourceIdentityUsergroups).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.Applications).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].Applications).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.ApplicationFamilies).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].ApplicationFamilies).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.Protocols).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].Protocols).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if helpers.GetStringFromSet(oldCItem.ProtocolNames).ValueString() != helpers.GetStringFromSet(data.Sequences[ni].MatchEntries[nci].ProtocolNames).ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
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
						resultC := make([]EmbeddedSecurityNGFWSequencesActions, 0, len(data.Sequences[ni].Actions))
						matchedC := make([]bool, len(data.Sequences[ni].Actions))
						for _, oldCItem := range oldItem.Actions {
							for nci := range data.Sequences[ni].Actions {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if oldCItem.Type.ValueString() != data.Sequences[ni].Actions[nci].Type.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.Parameter.ValueString() != data.Sequences[ni].Actions[nci].Parameter.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.ParameterId.ValueString() != data.Sequences[ni].Actions[nci].ParameterId.ValueString() {
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
