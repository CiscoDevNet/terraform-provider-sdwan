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
type TLSSSLProfilePolicyDefinition struct {
	Id                     types.String `tfsdk:"id"`
	Version                types.Int64  `tfsdk:"version"`
	Name                   types.String `tfsdk:"name"`
	Description            types.String `tfsdk:"description"`
	Mode                   types.String `tfsdk:"mode"`
	DecryptCategories      types.Set    `tfsdk:"decrypt_categories"`
	NeverDecryptCategories types.Set    `tfsdk:"never_decrypt_categories"`
	SkipDecryptCategories  types.Set    `tfsdk:"skip_decrypt_categories"`
	DecryptThreshold       types.String `tfsdk:"decrypt_threshold"`
	Reputation             types.Bool   `tfsdk:"reputation"`
	AllowUrlListId         types.String `tfsdk:"allow_url_list_id"`
	AllowUrlListVersion    types.Int64  `tfsdk:"allow_url_list_version"`
	BlockUrlListId         types.String `tfsdk:"block_url_list_id"`
	BlockUrlListVersion    types.Int64  `tfsdk:"block_url_list_version"`
	FailDecrypt            types.Bool   `tfsdk:"fail_decrypt"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TLSSSLProfilePolicyDefinition) getPath() string {
	return "/template/policy/definition/sslutdprofile/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TLSSSLProfilePolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "sslUtdDecryptProfile")
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
	if !data.DecryptCategories.IsNull() {
		var values []string
		data.DecryptCategories.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "definition.decryptCategories", values)
	}
	if !data.NeverDecryptCategories.IsNull() {
		var values []string
		data.NeverDecryptCategories.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "definition.neverDecryptCategories", values)
	}
	if !data.SkipDecryptCategories.IsNull() {
		var values []string
		data.SkipDecryptCategories.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "definition.skipDecryptCategories", values)
	}
	if !data.DecryptThreshold.IsNull() {
		body, _ = sjson.Set(body, "definition.decryptThreshold", data.DecryptThreshold.ValueString())
	}
	if !data.Reputation.IsNull() {
		if false && data.Reputation.ValueBool() {
			body, _ = sjson.Set(body, "definition.reputation", "")
		} else {
			body, _ = sjson.Set(body, "definition.reputation", data.Reputation.ValueBool())
		}
	}
	if !data.AllowUrlListId.IsNull() {
		body, _ = sjson.Set(body, "definition.urlWhiteList.ref", data.AllowUrlListId.ValueString())
	}
	if !data.BlockUrlListId.IsNull() {
		body, _ = sjson.Set(body, "definition.urlBlackList.ref", data.BlockUrlListId.ValueString())
	}
	if !data.FailDecrypt.IsNull() {
		if false && data.FailDecrypt.ValueBool() {
			body, _ = sjson.Set(body, "definition.failDecrypt", "")
		} else {
			body, _ = sjson.Set(body, "definition.failDecrypt", data.FailDecrypt.ValueBool())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TLSSSLProfilePolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("definition.decryptCategories"); value.Exists() {
		data.DecryptCategories = helpers.GetStringSet(value.Array())
	} else {
		data.DecryptCategories = types.SetNull(types.StringType)
	}
	if value := res.Get("definition.neverDecryptCategories"); value.Exists() {
		data.NeverDecryptCategories = helpers.GetStringSet(value.Array())
	} else {
		data.NeverDecryptCategories = types.SetNull(types.StringType)
	}
	if value := res.Get("definition.skipDecryptCategories"); value.Exists() {
		data.SkipDecryptCategories = helpers.GetStringSet(value.Array())
	} else {
		data.SkipDecryptCategories = types.SetNull(types.StringType)
	}
	if value := res.Get("definition.decryptThreshold"); value.Exists() {
		data.DecryptThreshold = types.StringValue(value.String())
	} else {
		data.DecryptThreshold = types.StringNull()
	}
	if value := res.Get("definition.reputation"); value.Exists() {
		if false && value.String() == "" {
			data.Reputation = types.BoolValue(true)
		} else {
			data.Reputation = types.BoolValue(value.Bool())
		}
	} else {
		data.Reputation = types.BoolNull()
	}
	if value := res.Get("definition.urlWhiteList.ref"); value.Exists() {
		data.AllowUrlListId = types.StringValue(value.String())
	} else {
		data.AllowUrlListId = types.StringNull()
	}
	if value := res.Get("definition.urlBlackList.ref"); value.Exists() {
		data.BlockUrlListId = types.StringValue(value.String())
	} else {
		data.BlockUrlListId = types.StringNull()
	}
	if value := res.Get("definition.failDecrypt"); value.Exists() {
		if false && value.String() == "" {
			data.FailDecrypt = types.BoolValue(true)
		} else {
			data.FailDecrypt = types.BoolValue(value.Bool())
		}
	} else {
		data.FailDecrypt = types.BoolNull()
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *TLSSSLProfilePolicyDefinition) hasChanges(ctx context.Context, state *TLSSSLProfilePolicyDefinition) bool {
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
	if !data.DecryptCategories.Equal(state.DecryptCategories) {
		hasChanges = true
	}
	if !data.NeverDecryptCategories.Equal(state.NeverDecryptCategories) {
		hasChanges = true
	}
	if !data.SkipDecryptCategories.Equal(state.SkipDecryptCategories) {
		hasChanges = true
	}
	if !data.DecryptThreshold.Equal(state.DecryptThreshold) {
		hasChanges = true
	}
	if !data.Reputation.Equal(state.Reputation) {
		hasChanges = true
	}
	if !data.AllowUrlListId.Equal(state.AllowUrlListId) {
		hasChanges = true
	}
	if !data.BlockUrlListId.Equal(state.BlockUrlListId) {
		hasChanges = true
	}
	if !data.FailDecrypt.Equal(state.FailDecrypt) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *TLSSSLProfilePolicyDefinition) updateVersions(ctx context.Context, state *TLSSSLProfilePolicyDefinition) {
	data.AllowUrlListVersion = state.AllowUrlListVersion
	data.BlockUrlListVersion = state.BlockUrlListVersion
}

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *TLSSSLProfilePolicyDefinition) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
	if data.AllowUrlListId != types.StringNull() {
		data.AllowUrlListVersion = types.Int64Value(0)
	}
	if data.BlockUrlListId != types.StringNull() {
		data.BlockUrlListVersion = types.Int64Value(0)
	}
}

// End of section. //template:end processImport
