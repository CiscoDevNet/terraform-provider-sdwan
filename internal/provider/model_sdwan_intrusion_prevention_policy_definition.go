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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type IntrusionPreventionPolicyDefinition struct {
	Id                      types.String `tfsdk:"id"`
	Version                 types.Int64  `tfsdk:"version"`
	Name                    types.String `tfsdk:"name"`
	Description             types.String `tfsdk:"description"`
	Mode                    types.String `tfsdk:"mode"`
	InspectionMode          types.String `tfsdk:"inspection_mode"`
	LogLevel                types.String `tfsdk:"log_level"`
	SignatureSet            types.String `tfsdk:"signature_set"`
	IpsSignatureListId      types.String `tfsdk:"ips_signature_list_id"`
	IpsSignatureListVersion types.Int64  `tfsdk:"ips_signature_list_version"`
	TargetVpns              types.Set    `tfsdk:"target_vpns"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data IntrusionPreventionPolicyDefinition) getPath() string {
	return "/template/policy/definition/intrusionprevention/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data IntrusionPreventionPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "intrusionPrevention")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, "mode", data.Mode.ValueString())
	}
	if !data.InspectionMode.IsNull() {
		body, _ = sjson.Set(body, "definition.inspectionMode", data.InspectionMode.ValueString())
	}
	if !data.LogLevel.IsNull() {
		body, _ = sjson.Set(body, "definition.logLevel", data.LogLevel.ValueString())
	}
	if !data.SignatureSet.IsNull() {
		body, _ = sjson.Set(body, "definition.signatureSet", data.SignatureSet.ValueString())
	}
	if !data.IpsSignatureListId.IsNull() {
		body, _ = sjson.Set(body, "definition.signatureWhiteList.ref", data.IpsSignatureListId.ValueString())
	}
	if !data.TargetVpns.IsNull() {
		var values []string
		data.TargetVpns.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "definition.targetVpns", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *IntrusionPreventionPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("mode"); value.Exists() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("definition.inspectionMode"); value.Exists() {
		data.InspectionMode = types.StringValue(value.String())
	} else {
		data.InspectionMode = types.StringNull()
	}
	if value := res.Get("definition.logLevel"); value.Exists() {
		data.LogLevel = types.StringValue(value.String())
	} else {
		data.LogLevel = types.StringNull()
	}
	if value := res.Get("definition.signatureSet"); value.Exists() {
		data.SignatureSet = types.StringValue(value.String())
	} else {
		data.SignatureSet = types.StringNull()
	}
	if value := res.Get("definition.signatureWhiteList.ref"); value.Exists() {
		data.IpsSignatureListId = types.StringValue(value.String())
	} else {
		data.IpsSignatureListId = types.StringNull()
	}
	if value := res.Get("definition.targetVpns"); value.Exists() {
		data.TargetVpns = helpers.GetStringSet(value.Array())
	} else {
		data.TargetVpns = types.SetNull(types.StringType)
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *IntrusionPreventionPolicyDefinition) hasChanges(ctx context.Context, state *IntrusionPreventionPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.Mode.Equal(state.Mode) {
		hasChanges = true
	}
	if !data.InspectionMode.Equal(state.InspectionMode) {
		hasChanges = true
	}
	if !data.LogLevel.Equal(state.LogLevel) {
		hasChanges = true
	}
	if !data.SignatureSet.Equal(state.SignatureSet) {
		hasChanges = true
	}
	if !data.IpsSignatureListId.Equal(state.IpsSignatureListId) {
		hasChanges = true
	}
	if !data.TargetVpns.Equal(state.TargetVpns) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *IntrusionPreventionPolicyDefinition) updateVersions(ctx context.Context, state *IntrusionPreventionPolicyDefinition) {
	data.IpsSignatureListVersion = state.IpsSignatureListVersion
}

// End of section. //template:end updateVersions
