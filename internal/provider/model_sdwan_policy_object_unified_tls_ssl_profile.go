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
type PolicyObjectUnifiedTLSSSLProfile struct {
	Id                    types.String `tfsdk:"id"`
	Version               types.Int64  `tfsdk:"version"`
	Name                  types.String `tfsdk:"name"`
	Description           types.String `tfsdk:"description"`
	FeatureProfileId      types.String `tfsdk:"feature_profile_id"`
	DecryptCategories     types.Set    `tfsdk:"decrypt_categories"`
	NoDecryptCategories   types.Set    `tfsdk:"no_decrypt_categories"`
	PassThroughCategories types.Set    `tfsdk:"pass_through_categories"`
	Reputation            types.Bool   `tfsdk:"reputation"`
	DecryptThreshold      types.String `tfsdk:"decrypt_threshold"`
	ThresholdCategories   types.String `tfsdk:"threshold_categories"`
	FailDecrypt           types.Bool   `tfsdk:"fail_decrypt"`
	UrlAllowListId        types.String `tfsdk:"url_allow_list_id"`
	UrlBlockListId        types.String `tfsdk:"url_block_list_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectUnifiedTLSSSLProfile) getModel() string {
	return "policy_object_unified_tls_ssl_profile"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectUnifiedTLSSSLProfile) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/unified/ssl-decryption-profile", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectUnifiedTLSSSLProfile) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.DecryptCategories.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"decryptCategories.optionType", "global")
			var values []string
			data.DecryptCategories.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"decryptCategories.value", values)
		}
	}
	if !data.NoDecryptCategories.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"neverDecryptCategories.optionType", "global")
			var values []string
			data.NoDecryptCategories.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"neverDecryptCategories.value", values)
		}
	}
	if !data.PassThroughCategories.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"skipDecryptCategories.optionType", "global")
			var values []string
			data.PassThroughCategories.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"skipDecryptCategories.value", values)
		}
	}
	if !data.Reputation.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"reputation.optionType", "global")
			body, _ = sjson.Set(body, path+"reputation.value", data.Reputation.ValueBool())
		}
	}
	if !data.DecryptThreshold.IsNull() {
		if true && data.Reputation.ValueBool() == true {
			body, _ = sjson.Set(body, path+"decryptThreshold.optionType", "global")
			body, _ = sjson.Set(body, path+"decryptThreshold.value", data.DecryptThreshold.ValueString())
		}
	}
	if !data.ThresholdCategories.IsNull() {
		if true && data.Reputation.ValueBool() == true {
			body, _ = sjson.Set(body, path+"skipDecryptThreshold.optionType", "global")
			body, _ = sjson.Set(body, path+"skipDecryptThreshold.value", data.ThresholdCategories.ValueString())
		}
	}
	if !data.FailDecrypt.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"failDecrypt.optionType", "global")
			body, _ = sjson.Set(body, path+"failDecrypt.value", data.FailDecrypt.ValueBool())
		}
	}
	if !data.UrlAllowListId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"urlAllowedList.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"urlAllowedList.refId.value", data.UrlAllowListId.ValueString())
		}
	}
	if !data.UrlBlockListId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"urlBlockedList.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"urlBlockedList.refId.value", data.UrlBlockListId.ValueString())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectUnifiedTLSSSLProfile) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.DecryptCategories = types.SetNull(types.StringType)

	if t := res.Get(path + "decryptCategories.optionType"); t.Exists() {
		va := res.Get(path + "decryptCategories.value")
		if t.String() == "global" {
			data.DecryptCategories = helpers.GetStringSet(va.Array())
		}
	}
	data.NoDecryptCategories = types.SetNull(types.StringType)

	if t := res.Get(path + "neverDecryptCategories.optionType"); t.Exists() {
		va := res.Get(path + "neverDecryptCategories.value")
		if t.String() == "global" {
			data.NoDecryptCategories = helpers.GetStringSet(va.Array())
		}
	}
	data.PassThroughCategories = types.SetNull(types.StringType)

	if t := res.Get(path + "skipDecryptCategories.optionType"); t.Exists() {
		va := res.Get(path + "skipDecryptCategories.value")
		if t.String() == "global" {
			data.PassThroughCategories = helpers.GetStringSet(va.Array())
		}
	}
	data.Reputation = types.BoolNull()

	if t := res.Get(path + "reputation.optionType"); t.Exists() {
		va := res.Get(path + "reputation.value")
		if t.String() == "global" {
			data.Reputation = types.BoolValue(va.Bool())
		}
	}
	data.DecryptThreshold = types.StringNull()

	if t := res.Get(path + "decryptThreshold.optionType"); t.Exists() {
		va := res.Get(path + "decryptThreshold.value")
		if t.String() == "global" {
			data.DecryptThreshold = types.StringValue(va.String())
		}
	}
	data.ThresholdCategories = types.StringNull()

	if t := res.Get(path + "skipDecryptThreshold.optionType"); t.Exists() {
		va := res.Get(path + "skipDecryptThreshold.value")
		if t.String() == "global" {
			data.ThresholdCategories = types.StringValue(va.String())
		}
	}
	data.FailDecrypt = types.BoolNull()

	if t := res.Get(path + "failDecrypt.optionType"); t.Exists() {
		va := res.Get(path + "failDecrypt.value")
		if t.String() == "global" {
			data.FailDecrypt = types.BoolValue(va.Bool())
		}
	}
	data.UrlAllowListId = types.StringNull()

	if t := res.Get(path + "urlAllowedList.refId.optionType"); t.Exists() {
		va := res.Get(path + "urlAllowedList.refId.value")
		if t.String() == "global" {
			data.UrlAllowListId = types.StringValue(va.String())
		}
	}
	data.UrlBlockListId = types.StringNull()

	if t := res.Get(path + "urlBlockedList.refId.optionType"); t.Exists() {
		va := res.Get(path + "urlBlockedList.refId.value")
		if t.String() == "global" {
			data.UrlBlockListId = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PolicyObjectUnifiedTLSSSLProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.DecryptCategories = types.SetNull(types.StringType)

	if t := res.Get(path + "decryptCategories.optionType"); t.Exists() {
		va := res.Get(path + "decryptCategories.value")
		if t.String() == "global" {
			data.DecryptCategories = helpers.GetStringSet(va.Array())
		}
	}
	data.NoDecryptCategories = types.SetNull(types.StringType)

	if t := res.Get(path + "neverDecryptCategories.optionType"); t.Exists() {
		va := res.Get(path + "neverDecryptCategories.value")
		if t.String() == "global" {
			data.NoDecryptCategories = helpers.GetStringSet(va.Array())
		}
	}
	data.PassThroughCategories = types.SetNull(types.StringType)

	if t := res.Get(path + "skipDecryptCategories.optionType"); t.Exists() {
		va := res.Get(path + "skipDecryptCategories.value")
		if t.String() == "global" {
			data.PassThroughCategories = helpers.GetStringSet(va.Array())
		}
	}
	data.Reputation = types.BoolNull()

	if t := res.Get(path + "reputation.optionType"); t.Exists() {
		va := res.Get(path + "reputation.value")
		if t.String() == "global" {
			data.Reputation = types.BoolValue(va.Bool())
		}
	}
	data.DecryptThreshold = types.StringNull()

	if t := res.Get(path + "decryptThreshold.optionType"); t.Exists() {
		va := res.Get(path + "decryptThreshold.value")
		if t.String() == "global" {
			data.DecryptThreshold = types.StringValue(va.String())
		}
	}
	data.ThresholdCategories = types.StringNull()

	if t := res.Get(path + "skipDecryptThreshold.optionType"); t.Exists() {
		va := res.Get(path + "skipDecryptThreshold.value")
		if t.String() == "global" {
			data.ThresholdCategories = types.StringValue(va.String())
		}
	}
	data.FailDecrypt = types.BoolNull()

	if t := res.Get(path + "failDecrypt.optionType"); t.Exists() {
		va := res.Get(path + "failDecrypt.value")
		if t.String() == "global" {
			data.FailDecrypt = types.BoolValue(va.Bool())
		}
	}
	data.UrlAllowListId = types.StringNull()

	if t := res.Get(path + "urlAllowedList.refId.optionType"); t.Exists() {
		va := res.Get(path + "urlAllowedList.refId.value")
		if t.String() == "global" {
			data.UrlAllowListId = types.StringValue(va.String())
		}
	}
	data.UrlBlockListId = types.StringNull()

	if t := res.Get(path + "urlBlockedList.refId.optionType"); t.Exists() {
		va := res.Get(path + "urlBlockedList.refId.value")
		if t.String() == "global" {
			data.UrlBlockListId = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *PolicyObjectUnifiedTLSSSLProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.DecryptCategories.IsNull() {
		return false
	}
	if !data.NoDecryptCategories.IsNull() {
		return false
	}
	if !data.PassThroughCategories.IsNull() {
		return false
	}
	if !data.Reputation.IsNull() {
		return false
	}
	if !data.DecryptThreshold.IsNull() {
		return false
	}
	if !data.ThresholdCategories.IsNull() {
		return false
	}
	if !data.FailDecrypt.IsNull() {
		return false
	}
	if !data.UrlAllowListId.IsNull() {
		return false
	}
	if !data.UrlBlockListId.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
