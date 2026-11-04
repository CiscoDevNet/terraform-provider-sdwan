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
type PolicyObjectSecurityObjectGroup struct {
	Id               types.String                             `tfsdk:"id"`
	Version          types.Int64                              `tfsdk:"version"`
	Name             types.String                             `tfsdk:"name"`
	Description      types.String                             `tfsdk:"description"`
	FeatureProfileId types.String                             `tfsdk:"feature_profile_id"`
	SequenceIpType   types.String                             `tfsdk:"sequence_ip_type"`
	Entries          []PolicyObjectSecurityObjectGroupEntries `tfsdk:"entries"`
}

type PolicyObjectSecurityObjectGroupEntries struct {
	DataIpv4Prefixes         types.Set    `tfsdk:"data_ipv4_prefixes"`
	DataIpv4PrefixesVariable types.String `tfsdk:"data_ipv4_prefixes_variable"`
	DataIpv6Prefixes         types.Set    `tfsdk:"data_ipv6_prefixes"`
	DataIpv6PrefixesVariable types.String `tfsdk:"data_ipv6_prefixes_variable"`
	Fqdns                    types.Set    `tfsdk:"fqdns"`
	FqdnsVariable            types.String `tfsdk:"fqdns_variable"`
	GeoLocations             types.Set    `tfsdk:"geo_locations"`
	GeoLocationsVariable     types.String `tfsdk:"geo_locations_variable"`
	Ports                    types.Set    `tfsdk:"ports"`
	PortsVariable            types.String `tfsdk:"ports_variable"`
	DataIpv4PrefixListIds    types.Set    `tfsdk:"data_ipv4_prefix_list_ids"`
	DataIpv6PrefixListIds    types.Set    `tfsdk:"data_ipv6_prefix_list_ids"`
	FqdnListIds              types.Set    `tfsdk:"fqdn_list_ids"`
	GeoLocationListIds       types.Set    `tfsdk:"geo_location_list_ids"`
	PortListIds              types.Set    `tfsdk:"port_list_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectSecurityObjectGroup) getModel() string {
	return "policy_object_security_object_group"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectSecurityObjectGroup) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/security-object-group", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectSecurityObjectGroup) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.SequenceIpType.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sequenceIpType.optionType", "global")
			body, _ = sjson.Set(body, path+"sequenceIpType.value", data.SequenceIpType.ValueString())
		}
	}
	if true {

		for _, item := range data.Entries {
			itemBody := ""

			if !item.DataIpv4PrefixesVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dataPrefix.ipv4Value.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "dataPrefix.ipv4Value.value", item.DataIpv4PrefixesVariable.ValueString())
				}
			} else if !item.DataIpv4Prefixes.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dataPrefix.ipv4Value.optionType", "global")
					var values []string
					item.DataIpv4Prefixes.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "dataPrefix.ipv4Value.value", values)
				}
			}

			if !item.DataIpv6PrefixesVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dataPrefixIpv6.ipv6Value.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "dataPrefixIpv6.ipv6Value.value", item.DataIpv6PrefixesVariable.ValueString())
				}
			} else if !item.DataIpv6Prefixes.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dataPrefixIpv6.ipv6Value.optionType", "global")
					var values []string
					item.DataIpv6Prefixes.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "dataPrefixIpv6.ipv6Value.value", values)
				}
			}

			if !item.FqdnsVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "fqdn.fqdnValue.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "fqdn.fqdnValue.value", item.FqdnsVariable.ValueString())
				}
			} else if !item.Fqdns.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "fqdn.fqdnValue.optionType", "global")
					var values []string
					item.Fqdns.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "fqdn.fqdnValue.value", values)
				}
			}

			if !item.GeoLocationsVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "geoLocation.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "geoLocation.value", item.GeoLocationsVariable.ValueString())
				}
			} else if !item.GeoLocations.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "geoLocation.optionType", "global")
					var values []string
					item.GeoLocations.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "geoLocation.value", values)
				}
			}

			if !item.PortsVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "port.portValue.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "port.portValue.value", item.PortsVariable.ValueString())
				}
			} else if !item.Ports.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "port.portValue.optionType", "global")
					var values []string
					item.Ports.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "port.portValue.value", values)
				}
			}
			if !item.DataIpv4PrefixListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dataPrefixList.refId.optionType", "global")
					var values []string
					item.DataIpv4PrefixListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "dataPrefixList.refId.value", values)
				}
			}
			if !item.DataIpv6PrefixListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "dataPrefixIpv6List.refId.optionType", "global")
					var values []string
					item.DataIpv6PrefixListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "dataPrefixIpv6List.refId.value", values)
				}
			}
			if !item.FqdnListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "fqdnList.refId.optionType", "global")
					var values []string
					item.FqdnListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "fqdnList.refId.value", values)
				}
			}
			if !item.GeoLocationListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "geoLocationList.refId.optionType", "global")
					var values []string
					item.GeoLocationListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "geoLocationList.refId.value", values)
				}
			}
			if !item.PortListIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "portList.refId.optionType", "global")
					var values []string
					item.PortListIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "portList.refId.value", values)
				}
			}
			body, _ = sjson.SetRaw(body, path+"entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectSecurityObjectGroup) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.SequenceIpType = types.StringNull()

	if t := res.Get(path + "sequenceIpType.optionType"); t.Exists() {
		va := res.Get(path + "sequenceIpType.value")
		if t.String() == "global" {
			data.SequenceIpType = types.StringValue(va.String())
		}
	}
	oldEntries := data.Entries
	if value := res.Get(path + "entries"); value.Exists() && len(value.Array()) > 0 {
		data.Entries = make([]PolicyObjectSecurityObjectGroupEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyObjectSecurityObjectGroupEntries{}
			item.DataIpv4Prefixes = types.SetNull(types.StringType)
			item.DataIpv4PrefixesVariable = types.StringNull()
			if t := v.Get("dataPrefix.ipv4Value.optionType"); t.Exists() {
				va := v.Get("dataPrefix.ipv4Value.value")
				if t.String() == "variable" {
					item.DataIpv4PrefixesVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.DataIpv4Prefixes = helpers.GetStringSet(va.Array())
				}
			}
			item.DataIpv6Prefixes = types.SetNull(types.StringType)
			item.DataIpv6PrefixesVariable = types.StringNull()
			if t := v.Get("dataPrefixIpv6.ipv6Value.optionType"); t.Exists() {
				va := v.Get("dataPrefixIpv6.ipv6Value.value")
				if t.String() == "variable" {
					item.DataIpv6PrefixesVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.DataIpv6Prefixes = helpers.GetStringSet(va.Array())
				}
			}
			item.Fqdns = types.SetNull(types.StringType)
			item.FqdnsVariable = types.StringNull()
			if t := v.Get("fqdn.fqdnValue.optionType"); t.Exists() {
				va := v.Get("fqdn.fqdnValue.value")
				if t.String() == "variable" {
					item.FqdnsVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Fqdns = helpers.GetStringSet(va.Array())
				}
			}
			item.GeoLocations = types.SetNull(types.StringType)
			item.GeoLocationsVariable = types.StringNull()
			if t := v.Get("geoLocation.optionType"); t.Exists() {
				va := v.Get("geoLocation.value")
				if t.String() == "variable" {
					item.GeoLocationsVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.GeoLocations = helpers.GetStringSet(va.Array())
				}
			}
			item.Ports = types.SetNull(types.StringType)
			item.PortsVariable = types.StringNull()
			if t := v.Get("port.portValue.optionType"); t.Exists() {
				va := v.Get("port.portValue.value")
				if t.String() == "variable" {
					item.PortsVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Ports = helpers.GetStringSet(va.Array())
				}
			}
			item.DataIpv4PrefixListIds = types.SetNull(types.StringType)

			if t := v.Get("dataPrefixList.refId.optionType"); t.Exists() {
				va := v.Get("dataPrefixList.refId.value")
				if t.String() == "global" {
					item.DataIpv4PrefixListIds = helpers.GetStringSet(va.Array())
				}
			}
			item.DataIpv6PrefixListIds = types.SetNull(types.StringType)

			if t := v.Get("dataPrefixIpv6List.refId.optionType"); t.Exists() {
				va := v.Get("dataPrefixIpv6List.refId.value")
				if t.String() == "global" {
					item.DataIpv6PrefixListIds = helpers.GetStringSet(va.Array())
				}
			}
			item.FqdnListIds = types.SetNull(types.StringType)

			if t := v.Get("fqdnList.refId.optionType"); t.Exists() {
				va := v.Get("fqdnList.refId.value")
				if t.String() == "global" {
					item.FqdnListIds = helpers.GetStringSet(va.Array())
				}
			}
			item.GeoLocationListIds = types.SetNull(types.StringType)

			if t := v.Get("geoLocationList.refId.optionType"); t.Exists() {
				va := v.Get("geoLocationList.refId.value")
				if t.String() == "global" {
					item.GeoLocationListIds = helpers.GetStringSet(va.Array())
				}
			}
			item.PortListIds = types.SetNull(types.StringType)

			if t := v.Get("portList.refId.optionType"); t.Exists() {
				va := v.Get("portList.refId.value")
				if t.String() == "global" {
					item.PortListIds = helpers.GetStringSet(va.Array())
				}
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	} else {
		data.Entries = nil
	}
	if !fullRead && data.Entries != nil {
		resultEntries := make([]PolicyObjectSecurityObjectGroupEntries, 0, len(data.Entries))
		matchedEntries := make([]bool, len(data.Entries))
		for _, oldItem := range oldEntries {
			for ni := range data.Entries {
				if matchedEntries[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.DataIpv4Prefixes).ValueString() != helpers.GetStringFromSet(data.Entries[ni].DataIpv4Prefixes).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.DataIpv6Prefixes).ValueString() != helpers.GetStringFromSet(data.Entries[ni].DataIpv6Prefixes).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.Fqdns).ValueString() != helpers.GetStringFromSet(data.Entries[ni].Fqdns).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.GeoLocations).ValueString() != helpers.GetStringFromSet(data.Entries[ni].GeoLocations).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.Ports).ValueString() != helpers.GetStringFromSet(data.Entries[ni].Ports).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.DataIpv4PrefixListIds).ValueString() != helpers.GetStringFromSet(data.Entries[ni].DataIpv4PrefixListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.DataIpv6PrefixListIds).ValueString() != helpers.GetStringFromSet(data.Entries[ni].DataIpv6PrefixListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.FqdnListIds).ValueString() != helpers.GetStringFromSet(data.Entries[ni].FqdnListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.GeoLocationListIds).ValueString() != helpers.GetStringFromSet(data.Entries[ni].GeoLocationListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					if helpers.GetStringFromSet(oldItem.PortListIds).ValueString() != helpers.GetStringFromSet(data.Entries[ni].PortListIds).ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedEntries[ni] = true
					resultEntries = append(resultEntries, data.Entries[ni])
					break
				}
			}
		}
		for ni := range data.Entries {
			if !matchedEntries[ni] {
				resultEntries = append(resultEntries, data.Entries[ni])
			}
		}
		data.Entries = resultEntries
	}
}

// End of section. //template:end fromBody
