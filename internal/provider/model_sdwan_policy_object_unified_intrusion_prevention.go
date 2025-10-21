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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type PolicyObjectUnifiedIntrusionPrevention struct {
	Id                      types.String `tfsdk:"id"`
	Version                 types.Int64  `tfsdk:"version"`
	Name                    types.String `tfsdk:"name"`
	Description             types.String `tfsdk:"description"`
	FeatureProfileId        types.String `tfsdk:"feature_profile_id"`
	SignatureSet            types.String `tfsdk:"signature_set"`
	InspectionMode          types.String `tfsdk:"inspection_mode"`
	IpsSignatureAllowListId types.String `tfsdk:"ips_signature_allow_list_id"`
	LogLevel                types.String `tfsdk:"log_level"`
	CustomSignature         types.Bool   `tfsdk:"custom_signature"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectUnifiedIntrusionPrevention) getModel() string {
	return "policy_object_unified_intrusion_prevention"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectUnifiedIntrusionPrevention) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/unified/intrusion-prevention", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectUnifiedIntrusionPrevention) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.SignatureSet.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"signatureSet.optionType", "global")
			body, _ = sjson.Set(body, path+"signatureSet.value", data.SignatureSet.ValueString())
		}
	}
	if !data.InspectionMode.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"inspectionMode.optionType", "global")
			body, _ = sjson.Set(body, path+"inspectionMode.value", data.InspectionMode.ValueString())
		}
	}
	if !data.IpsSignatureAllowListId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"signatureAllowedList.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"signatureAllowedList.refId.value", data.IpsSignatureAllowListId.ValueString())
		}
	}
	if !data.LogLevel.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"logLevel.optionType", "global")
			body, _ = sjson.Set(body, path+"logLevel.value", data.LogLevel.ValueString())
		}
	}
	if !data.CustomSignature.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"customSignature.optionType", "global")
			body, _ = sjson.Set(body, path+"customSignature.value", data.CustomSignature.ValueBool())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectUnifiedIntrusionPrevention) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.SignatureSet = types.StringNull()

	if t := res.Get(path + "signatureSet.optionType"); t.Exists() {
		va := res.Get(path + "signatureSet.value")
		if t.String() == "global" {
			data.SignatureSet = types.StringValue(va.String())
		}
	}
	data.InspectionMode = types.StringNull()

	if t := res.Get(path + "inspectionMode.optionType"); t.Exists() {
		va := res.Get(path + "inspectionMode.value")
		if t.String() == "global" {
			data.InspectionMode = types.StringValue(va.String())
		}
	}
	data.IpsSignatureAllowListId = types.StringNull()

	if t := res.Get(path + "signatureAllowedList.refId.optionType"); t.Exists() {
		va := res.Get(path + "signatureAllowedList.refId.value")
		if t.String() == "global" {
			data.IpsSignatureAllowListId = types.StringValue(va.String())
		}
	}
	data.LogLevel = types.StringNull()

	if t := res.Get(path + "logLevel.optionType"); t.Exists() {
		va := res.Get(path + "logLevel.value")
		if t.String() == "global" {
			data.LogLevel = types.StringValue(va.String())
		}
	}
	data.CustomSignature = types.BoolNull()

	if t := res.Get(path + "customSignature.optionType"); t.Exists() {
		va := res.Get(path + "customSignature.value")
		if t.String() == "global" {
			data.CustomSignature = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PolicyObjectUnifiedIntrusionPrevention) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.SignatureSet = types.StringNull()

	if t := res.Get(path + "signatureSet.optionType"); t.Exists() {
		va := res.Get(path + "signatureSet.value")
		if t.String() == "global" {
			data.SignatureSet = types.StringValue(va.String())
		}
	}
	data.InspectionMode = types.StringNull()

	if t := res.Get(path + "inspectionMode.optionType"); t.Exists() {
		va := res.Get(path + "inspectionMode.value")
		if t.String() == "global" {
			data.InspectionMode = types.StringValue(va.String())
		}
	}
	data.IpsSignatureAllowListId = types.StringNull()

	if t := res.Get(path + "signatureAllowedList.refId.optionType"); t.Exists() {
		va := res.Get(path + "signatureAllowedList.refId.value")
		if t.String() == "global" {
			data.IpsSignatureAllowListId = types.StringValue(va.String())
		}
	}
	data.LogLevel = types.StringNull()

	if t := res.Get(path + "logLevel.optionType"); t.Exists() {
		va := res.Get(path + "logLevel.value")
		if t.String() == "global" {
			data.LogLevel = types.StringValue(va.String())
		}
	}
	data.CustomSignature = types.BoolNull()

	if t := res.Get(path + "customSignature.optionType"); t.Exists() {
		va := res.Get(path + "customSignature.value")
		if t.String() == "global" {
			data.CustomSignature = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end updateFromBody
