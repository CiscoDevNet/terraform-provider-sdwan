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
type PolicyObjectUnifiedAdvancedInspectionProfile struct {
	Id                              types.String `tfsdk:"id"`
	Version                         types.Int64  `tfsdk:"version"`
	Name                            types.String `tfsdk:"name"`
	Description                     types.String `tfsdk:"description"`
	FeatureProfileId                types.String `tfsdk:"feature_profile_id"`
	TlsDecryptionAction             types.String `tfsdk:"tls_decryption_action"`
	IntrusionPreventionListId       types.String `tfsdk:"intrusion_prevention_list_id"`
	UrlFilteringListId              types.String `tfsdk:"url_filtering_list_id"`
	AdvancedMalwareProtectionListId types.String `tfsdk:"advanced_malware_protection_list_id"`
	TlsSslProfileListId             types.String `tfsdk:"tls_ssl_profile_list_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectUnifiedAdvancedInspectionProfile) getModel() string {
	return "policy_object_unified_advanced_inspection_profile"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectUnifiedAdvancedInspectionProfile) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/unified/advanced-inspection-profile", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectUnifiedAdvancedInspectionProfile) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.TlsDecryptionAction.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tlsDecryptionAction.optionType", "global")
			body, _ = sjson.Set(body, path+"tlsDecryptionAction.value", data.TlsDecryptionAction.ValueString())
		}
	}
	if !data.IntrusionPreventionListId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"intrusionPrevention.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"intrusionPrevention.refId.value", data.IntrusionPreventionListId.ValueString())
		}
	}
	if !data.UrlFilteringListId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"urlFiltering.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"urlFiltering.refId.value", data.UrlFilteringListId.ValueString())
		}
	}
	if !data.AdvancedMalwareProtectionListId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advancedMalwareProtection.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"advancedMalwareProtection.refId.value", data.AdvancedMalwareProtectionListId.ValueString())
		}
	}
	if !data.TlsSslProfileListId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sslDecryptionProfile.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"sslDecryptionProfile.refId.value", data.TlsSslProfileListId.ValueString())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectUnifiedAdvancedInspectionProfile) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.TlsDecryptionAction = types.StringNull()

	if t := res.Get(path + "tlsDecryptionAction.optionType"); t.Exists() {
		va := res.Get(path + "tlsDecryptionAction.value")
		if t.String() == "global" {
			data.TlsDecryptionAction = types.StringValue(va.String())
		}
	}
	data.IntrusionPreventionListId = types.StringNull()

	if t := res.Get(path + "intrusionPrevention.refId.optionType"); t.Exists() {
		va := res.Get(path + "intrusionPrevention.refId.value")
		if t.String() == "global" {
			data.IntrusionPreventionListId = types.StringValue(va.String())
		}
	}
	data.UrlFilteringListId = types.StringNull()

	if t := res.Get(path + "urlFiltering.refId.optionType"); t.Exists() {
		va := res.Get(path + "urlFiltering.refId.value")
		if t.String() == "global" {
			data.UrlFilteringListId = types.StringValue(va.String())
		}
	}
	data.AdvancedMalwareProtectionListId = types.StringNull()

	if t := res.Get(path + "advancedMalwareProtection.refId.optionType"); t.Exists() {
		va := res.Get(path + "advancedMalwareProtection.refId.value")
		if t.String() == "global" {
			data.AdvancedMalwareProtectionListId = types.StringValue(va.String())
		}
	}
	data.TlsSslProfileListId = types.StringNull()

	if t := res.Get(path + "sslDecryptionProfile.refId.optionType"); t.Exists() {
		va := res.Get(path + "sslDecryptionProfile.refId.value")
		if t.String() == "global" {
			data.TlsSslProfileListId = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PolicyObjectUnifiedAdvancedInspectionProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.TlsDecryptionAction = types.StringNull()

	if t := res.Get(path + "tlsDecryptionAction.optionType"); t.Exists() {
		va := res.Get(path + "tlsDecryptionAction.value")
		if t.String() == "global" {
			data.TlsDecryptionAction = types.StringValue(va.String())
		}
	}
	data.IntrusionPreventionListId = types.StringNull()

	if t := res.Get(path + "intrusionPrevention.refId.optionType"); t.Exists() {
		va := res.Get(path + "intrusionPrevention.refId.value")
		if t.String() == "global" {
			data.IntrusionPreventionListId = types.StringValue(va.String())
		}
	}
	data.UrlFilteringListId = types.StringNull()

	if t := res.Get(path + "urlFiltering.refId.optionType"); t.Exists() {
		va := res.Get(path + "urlFiltering.refId.value")
		if t.String() == "global" {
			data.UrlFilteringListId = types.StringValue(va.String())
		}
	}
	data.AdvancedMalwareProtectionListId = types.StringNull()

	if t := res.Get(path + "advancedMalwareProtection.refId.optionType"); t.Exists() {
		va := res.Get(path + "advancedMalwareProtection.refId.value")
		if t.String() == "global" {
			data.AdvancedMalwareProtectionListId = types.StringValue(va.String())
		}
	}
	data.TlsSslProfileListId = types.StringNull()

	if t := res.Get(path + "sslDecryptionProfile.refId.optionType"); t.Exists() {
		va := res.Get(path + "sslDecryptionProfile.refId.value")
		if t.String() == "global" {
			data.TlsSslProfileListId = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *PolicyObjectUnifiedAdvancedInspectionProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.TlsDecryptionAction.IsNull() {
		return false
	}
	if !data.IntrusionPreventionListId.IsNull() {
		return false
	}
	if !data.UrlFilteringListId.IsNull() {
		return false
	}
	if !data.AdvancedMalwareProtectionListId.IsNull() {
		return false
	}
	if !data.TlsSslProfileListId.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
