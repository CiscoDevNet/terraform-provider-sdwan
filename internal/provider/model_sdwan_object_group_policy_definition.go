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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ObjectGroupPolicyDefinition struct {
	Id                        types.String `tfsdk:"id"`
	Version                   types.Int64  `tfsdk:"version"`
	Name                      types.String `tfsdk:"name"`
	Description               types.String `tfsdk:"description"`
	DataIpv4PrefixListId      types.String `tfsdk:"data_ipv4_prefix_list_id"`
	DataIpv4PrefixListVersion types.Int64  `tfsdk:"data_ipv4_prefix_list_version"`
	Ipv4PrefixVariable        types.String `tfsdk:"ipv4_prefix_variable"`
	Ipv4Prefix                types.String `tfsdk:"ipv4_prefix"`
	DataFqdnPrefixListId      types.String `tfsdk:"data_fqdn_prefix_list_id"`
	DataFqdnPrefixListVersion types.Int64  `tfsdk:"data_fqdn_prefix_list_version"`
	Fqdn                      types.String `tfsdk:"fqdn"`
	PortListId                types.String `tfsdk:"port_list_id"`
	PortListVersion           types.Int64  `tfsdk:"port_list_version"`
	Port                      types.String `tfsdk:"port"`
	GeoLocationListId         types.String `tfsdk:"geo_location_list_id"`
	GeoLocationListVersion    types.Int64  `tfsdk:"geo_location_list_version"`
	GeoLocation               types.String `tfsdk:"geo_location"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ObjectGroupPolicyDefinition) getPath() string {
	return "/template/policy/definition/securitygroup/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ObjectGroupPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "securityGroup")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.DataIpv4PrefixListId.IsNull() {
		body, _ = sjson.Set(body, "definition.dataPrefixList.ref", data.DataIpv4PrefixListId.ValueString())
	}
	if !data.Ipv4PrefixVariable.IsNull() {
		body, _ = sjson.Set(body, "definition.dataPrefix.vipVariableName", data.Ipv4PrefixVariable.ValueString())
	}
	if !data.Ipv4Prefix.IsNull() {
		body, _ = sjson.Set(body, "definition.dataPrefix", data.Ipv4Prefix.ValueString())
	}
	if !data.DataFqdnPrefixListId.IsNull() {
		body, _ = sjson.Set(body, "definition.fqdnList.ref", data.DataFqdnPrefixListId.ValueString())
	}
	if !data.Fqdn.IsNull() {
		body, _ = sjson.Set(body, "definition.fqdn", data.Fqdn.ValueString())
	}
	if !data.PortListId.IsNull() {
		body, _ = sjson.Set(body, "definition.portList.ref", data.PortListId.ValueString())
	}
	if !data.Port.IsNull() {
		body, _ = sjson.Set(body, "definition.port", data.Port.ValueString())
	}
	if !data.GeoLocationListId.IsNull() {
		body, _ = sjson.Set(body, "definition.geoLocationList.ref", data.GeoLocationListId.ValueString())
	}
	if !data.GeoLocation.IsNull() {
		body, _ = sjson.Set(body, "definition.geoLocation", data.GeoLocation.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ObjectGroupPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("definition.dataPrefixList.ref"); value.Exists() {
		data.DataIpv4PrefixListId = types.StringValue(value.String())
	} else {
		data.DataIpv4PrefixListId = types.StringNull()
	}
	if value := res.Get("definition.dataPrefix.vipVariableName"); value.Exists() {
		data.Ipv4PrefixVariable = types.StringValue(value.String())
	} else {
		data.Ipv4PrefixVariable = types.StringNull()
	}
	if value := res.Get("definition.dataPrefix"); value.Exists() {
		data.Ipv4Prefix = types.StringValue(value.String())
	} else {
		data.Ipv4Prefix = types.StringNull()
	}
	if value := res.Get("definition.fqdnList.ref"); value.Exists() {
		data.DataFqdnPrefixListId = types.StringValue(value.String())
	} else {
		data.DataFqdnPrefixListId = types.StringNull()
	}
	if value := res.Get("definition.fqdn"); value.Exists() {
		data.Fqdn = types.StringValue(value.String())
	} else {
		data.Fqdn = types.StringNull()
	}
	if value := res.Get("definition.portList.ref"); value.Exists() {
		data.PortListId = types.StringValue(value.String())
	} else {
		data.PortListId = types.StringNull()
	}
	if value := res.Get("definition.port"); value.Exists() {
		data.Port = types.StringValue(value.String())
	} else {
		data.Port = types.StringNull()
	}
	if value := res.Get("definition.geoLocationList.ref"); value.Exists() {
		data.GeoLocationListId = types.StringValue(value.String())
	} else {
		data.GeoLocationListId = types.StringNull()
	}
	if value := res.Get("definition.geoLocation"); value.Exists() {
		data.GeoLocation = types.StringValue(value.String())
	} else {
		data.GeoLocation = types.StringNull()
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *ObjectGroupPolicyDefinition) hasChanges(ctx context.Context, state *ObjectGroupPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.DataIpv4PrefixListId.Equal(state.DataIpv4PrefixListId) {
		hasChanges = true
	}
	if !data.Ipv4PrefixVariable.Equal(state.Ipv4PrefixVariable) {
		hasChanges = true
	}
	if !data.Ipv4Prefix.Equal(state.Ipv4Prefix) {
		hasChanges = true
	}
	if !data.DataFqdnPrefixListId.Equal(state.DataFqdnPrefixListId) {
		hasChanges = true
	}
	if !data.Fqdn.Equal(state.Fqdn) {
		hasChanges = true
	}
	if !data.PortListId.Equal(state.PortListId) {
		hasChanges = true
	}
	if !data.Port.Equal(state.Port) {
		hasChanges = true
	}
	if !data.GeoLocationListId.Equal(state.GeoLocationListId) {
		hasChanges = true
	}
	if !data.GeoLocation.Equal(state.GeoLocation) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *ObjectGroupPolicyDefinition) updateVersions(ctx context.Context, state *ObjectGroupPolicyDefinition) {
	data.DataIpv4PrefixListVersion = state.DataIpv4PrefixListVersion
	data.DataFqdnPrefixListVersion = state.DataFqdnPrefixListVersion
	data.PortListVersion = state.PortListVersion
	data.GeoLocationListVersion = state.GeoLocationListVersion
}

// End of section. //template:end updateVersions
