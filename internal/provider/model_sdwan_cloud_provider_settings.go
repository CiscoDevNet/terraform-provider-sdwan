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
type CloudProviderSettings struct {
	Id                             types.String `tfsdk:"id"`
	UmbrellaOrgId                  types.String `tfsdk:"umbrella_org_id"`
	UmbrellaAuthKeyV2              types.String `tfsdk:"umbrella_auth_key_v2"`
	UmbrellaAuthKeyV2Wo            types.String `tfsdk:"umbrella_auth_key_v2_wo"`
	UmbrellaAuthKeyV2WoVersion     types.Int64  `tfsdk:"umbrella_auth_key_v2_wo_version"`
	UmbrellaAuthSecretV2           types.String `tfsdk:"umbrella_auth_secret_v2"`
	UmbrellaAuthSecretV2Wo         types.String `tfsdk:"umbrella_auth_secret_v2_wo"`
	UmbrellaAuthSecretV2WoVersion  types.Int64  `tfsdk:"umbrella_auth_secret_v2_wo_version"`
	UmbrellaSigAuthKey             types.String `tfsdk:"umbrella_sig_auth_key"`
	UmbrellaSigAuthKeyWo           types.String `tfsdk:"umbrella_sig_auth_key_wo"`
	UmbrellaSigAuthKeyWoVersion    types.Int64  `tfsdk:"umbrella_sig_auth_key_wo_version"`
	UmbrellaSigAuthSecret          types.String `tfsdk:"umbrella_sig_auth_secret"`
	UmbrellaSigAuthSecretWo        types.String `tfsdk:"umbrella_sig_auth_secret_wo"`
	UmbrellaSigAuthSecretWoVersion types.Int64  `tfsdk:"umbrella_sig_auth_secret_wo_version"`
	UmbrellaDnsAuthKey             types.String `tfsdk:"umbrella_dns_auth_key"`
	UmbrellaDnsAuthKeyWo           types.String `tfsdk:"umbrella_dns_auth_key_wo"`
	UmbrellaDnsAuthKeyWoVersion    types.Int64  `tfsdk:"umbrella_dns_auth_key_wo_version"`
	UmbrellaDnsAuthSecret          types.String `tfsdk:"umbrella_dns_auth_secret"`
	UmbrellaDnsAuthSecretWo        types.String `tfsdk:"umbrella_dns_auth_secret_wo"`
	UmbrellaDnsAuthSecretWoVersion types.Int64  `tfsdk:"umbrella_dns_auth_secret_wo_version"`
	ZscalerOrganization            types.String `tfsdk:"zscaler_organization"`
	ZscalerPartnerBaseUri          types.String `tfsdk:"zscaler_partner_base_uri"`
	ZscalerPartnerKey              types.String `tfsdk:"zscaler_partner_key"`
	ZscalerPartnerKeyWo            types.String `tfsdk:"zscaler_partner_key_wo"`
	ZscalerPartnerKeyWoVersion     types.Int64  `tfsdk:"zscaler_partner_key_wo_version"`
	ZscalerUsername                types.String `tfsdk:"zscaler_username"`
	ZscalerPassword                types.String `tfsdk:"zscaler_password"`
	ZscalerPasswordWo              types.String `tfsdk:"zscaler_password_wo"`
	ZscalerPasswordWoVersion       types.Int64  `tfsdk:"zscaler_password_wo_version"`
	CiscoSseOrgId                  types.String `tfsdk:"cisco_sse_org_id"`
	CiscoSseAuthKey                types.String `tfsdk:"cisco_sse_auth_key"`
	CiscoSseAuthKeyWo              types.String `tfsdk:"cisco_sse_auth_key_wo"`
	CiscoSseAuthKeyWoVersion       types.Int64  `tfsdk:"cisco_sse_auth_key_wo_version"`
	CiscoSseAuthSecret             types.String `tfsdk:"cisco_sse_auth_secret"`
	CiscoSseAuthSecretWo           types.String `tfsdk:"cisco_sse_auth_secret_wo"`
	CiscoSseAuthSecretWoVersion    types.Int64  `tfsdk:"cisco_sse_auth_secret_wo_version"`
	CiscoSseContextSharing         types.Bool   `tfsdk:"cisco_sse_context_sharing"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data CloudProviderSettings) getPath() string {
	return "/settings/configuration/cloudProviderSetting"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CloudProviderSettings) toBody(ctx context.Context) string {
	body := ""
	if !data.UmbrellaOrgId.IsNull() {
		body, _ = sjson.Set(body, "umbrellaOrgId", data.UmbrellaOrgId.ValueString())
	}
	if !data.UmbrellaAuthKeyV2.IsNull() {
		body, _ = sjson.Set(body, "umbrellaAuthKeyV2", data.UmbrellaAuthKeyV2.ValueString())
	}
	if !data.UmbrellaAuthKeyV2Wo.IsNull() {
		body, _ = sjson.Set(body, "umbrellaAuthKeyV2", data.UmbrellaAuthKeyV2Wo.ValueString())
	}
	if !data.UmbrellaAuthSecretV2.IsNull() {
		body, _ = sjson.Set(body, "umbrellaAuthSecretV2", data.UmbrellaAuthSecretV2.ValueString())
	}
	if !data.UmbrellaAuthSecretV2Wo.IsNull() {
		body, _ = sjson.Set(body, "umbrellaAuthSecretV2", data.UmbrellaAuthSecretV2Wo.ValueString())
	}
	if !data.UmbrellaSigAuthKey.IsNull() {
		body, _ = sjson.Set(body, "umbrellaSIGAuthKey", data.UmbrellaSigAuthKey.ValueString())
	}
	if !data.UmbrellaSigAuthKeyWo.IsNull() {
		body, _ = sjson.Set(body, "umbrellaSIGAuthKey", data.UmbrellaSigAuthKeyWo.ValueString())
	}
	if !data.UmbrellaSigAuthSecret.IsNull() {
		body, _ = sjson.Set(body, "umbrellaSIGAuthSecret", data.UmbrellaSigAuthSecret.ValueString())
	}
	if !data.UmbrellaSigAuthSecretWo.IsNull() {
		body, _ = sjson.Set(body, "umbrellaSIGAuthSecret", data.UmbrellaSigAuthSecretWo.ValueString())
	}
	if !data.UmbrellaDnsAuthKey.IsNull() {
		body, _ = sjson.Set(body, "umbrellaDNSAuthKey", data.UmbrellaDnsAuthKey.ValueString())
	}
	if !data.UmbrellaDnsAuthKeyWo.IsNull() {
		body, _ = sjson.Set(body, "umbrellaDNSAuthKey", data.UmbrellaDnsAuthKeyWo.ValueString())
	}
	if !data.UmbrellaDnsAuthSecret.IsNull() {
		body, _ = sjson.Set(body, "umbrellaDNSAuthSecret", data.UmbrellaDnsAuthSecret.ValueString())
	}
	if !data.UmbrellaDnsAuthSecretWo.IsNull() {
		body, _ = sjson.Set(body, "umbrellaDNSAuthSecret", data.UmbrellaDnsAuthSecretWo.ValueString())
	}
	if !data.ZscalerOrganization.IsNull() {
		body, _ = sjson.Set(body, "zscalerOrganization", data.ZscalerOrganization.ValueString())
	}
	if !data.ZscalerPartnerBaseUri.IsNull() {
		body, _ = sjson.Set(body, "zscalerPartnerBaseUri", data.ZscalerPartnerBaseUri.ValueString())
	}
	if !data.ZscalerPartnerKey.IsNull() {
		body, _ = sjson.Set(body, "zscalerPartnerKey", data.ZscalerPartnerKey.ValueString())
	}
	if !data.ZscalerPartnerKeyWo.IsNull() {
		body, _ = sjson.Set(body, "zscalerPartnerKey", data.ZscalerPartnerKeyWo.ValueString())
	}
	if !data.ZscalerUsername.IsNull() {
		body, _ = sjson.Set(body, "zscalerUsername", data.ZscalerUsername.ValueString())
	}
	if !data.ZscalerPassword.IsNull() {
		body, _ = sjson.Set(body, "zscalerPassword", data.ZscalerPassword.ValueString())
	}
	if !data.ZscalerPasswordWo.IsNull() {
		body, _ = sjson.Set(body, "zscalerPassword", data.ZscalerPasswordWo.ValueString())
	}
	if !data.CiscoSseOrgId.IsNull() {
		body, _ = sjson.Set(body, "ciscoSSEOrgId", data.CiscoSseOrgId.ValueString())
	}
	if !data.CiscoSseAuthKey.IsNull() {
		body, _ = sjson.Set(body, "ciscoSSEAuthKey", data.CiscoSseAuthKey.ValueString())
	}
	if !data.CiscoSseAuthKeyWo.IsNull() {
		body, _ = sjson.Set(body, "ciscoSSEAuthKey", data.CiscoSseAuthKeyWo.ValueString())
	}
	if !data.CiscoSseAuthSecret.IsNull() {
		body, _ = sjson.Set(body, "ciscoSSEAuthSecret", data.CiscoSseAuthSecret.ValueString())
	}
	if !data.CiscoSseAuthSecretWo.IsNull() {
		body, _ = sjson.Set(body, "ciscoSSEAuthSecret", data.CiscoSseAuthSecretWo.ValueString())
	}
	if !data.CiscoSseContextSharing.IsNull() {
		if false && data.CiscoSseContextSharing.ValueBool() {
			body, _ = sjson.Set(body, "ciscoSSEContextSharing", "")
		} else {
			body, _ = sjson.Set(body, "ciscoSSEContextSharing", data.CiscoSseContextSharing.ValueBool())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CloudProviderSettings) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("data.0.umbrellaOrgId"); value.Exists() {
		data.UmbrellaOrgId = types.StringValue(value.String())
	} else {
		data.UmbrellaOrgId = types.StringNull()
	}
	if value := res.Get("data.0.umbrellaAuthKeyV2"); value.Exists() {
		data.UmbrellaAuthKeyV2 = types.StringValue(value.String())
	} else {
		data.UmbrellaAuthKeyV2 = types.StringNull()
	}
	if value := res.Get("data.0.umbrellaAuthSecretV2"); value.Exists() {
		data.UmbrellaAuthSecretV2 = types.StringValue(value.String())
	} else {
		data.UmbrellaAuthSecretV2 = types.StringNull()
	}
	if value := res.Get("data.0.umbrellaSIGAuthKey"); value.Exists() {
		data.UmbrellaSigAuthKey = types.StringValue(value.String())
	} else {
		data.UmbrellaSigAuthKey = types.StringNull()
	}
	if value := res.Get("data.0.umbrellaDNSAuthKey"); value.Exists() {
		data.UmbrellaDnsAuthKey = types.StringValue(value.String())
	} else {
		data.UmbrellaDnsAuthKey = types.StringNull()
	}
	if value := res.Get("data.0.zscalerOrganization"); value.Exists() {
		data.ZscalerOrganization = types.StringValue(value.String())
	} else {
		data.ZscalerOrganization = types.StringNull()
	}
	if value := res.Get("data.0.zscalerPartnerBaseUri"); value.Exists() {
		data.ZscalerPartnerBaseUri = types.StringValue(value.String())
	} else {
		data.ZscalerPartnerBaseUri = types.StringNull()
	}
	if value := res.Get("data.0.zscalerPartnerKey"); value.Exists() {
		data.ZscalerPartnerKey = types.StringValue(value.String())
	} else {
		data.ZscalerPartnerKey = types.StringNull()
	}
	if value := res.Get("data.0.zscalerUsername"); value.Exists() {
		data.ZscalerUsername = types.StringValue(value.String())
	} else {
		data.ZscalerUsername = types.StringNull()
	}
	if value := res.Get("data.0.ciscoSSEOrgId"); value.Exists() {
		data.CiscoSseOrgId = types.StringValue(value.String())
	} else {
		data.CiscoSseOrgId = types.StringNull()
	}
	if value := res.Get("data.0.ciscoSSEAuthKey"); value.Exists() {
		data.CiscoSseAuthKey = types.StringValue(value.String())
	} else {
		data.CiscoSseAuthKey = types.StringNull()
	}
	if value := res.Get("data.0.ciscoSSEAuthSecret"); value.Exists() {
		data.CiscoSseAuthSecret = types.StringValue(value.String())
	} else {
		data.CiscoSseAuthSecret = types.StringNull()
	}
	if value := res.Get("data.0.ciscoSSEContextSharing"); value.Exists() {
		if false && value.String() == "" {
			data.CiscoSseContextSharing = types.BoolValue(true)
		} else {
			data.CiscoSseContextSharing = types.BoolValue(value.Bool())
		}
	} else {
		data.CiscoSseContextSharing = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CloudProviderSettings) hasChanges(ctx context.Context, state *CloudProviderSettings) bool {
	hasChanges := false
	if !data.UmbrellaOrgId.Equal(state.UmbrellaOrgId) {
		hasChanges = true
	}
	if !data.UmbrellaAuthKeyV2.Equal(state.UmbrellaAuthKeyV2) {
		hasChanges = true
	}
	if !data.UmbrellaAuthKeyV2Wo.Equal(state.UmbrellaAuthKeyV2Wo) {
		hasChanges = true
	}
	if !data.UmbrellaAuthSecretV2.Equal(state.UmbrellaAuthSecretV2) {
		hasChanges = true
	}
	if !data.UmbrellaAuthSecretV2Wo.Equal(state.UmbrellaAuthSecretV2Wo) {
		hasChanges = true
	}
	if !data.UmbrellaSigAuthKey.Equal(state.UmbrellaSigAuthKey) {
		hasChanges = true
	}
	if !data.UmbrellaSigAuthKeyWo.Equal(state.UmbrellaSigAuthKeyWo) {
		hasChanges = true
	}
	if !data.UmbrellaSigAuthSecret.Equal(state.UmbrellaSigAuthSecret) {
		hasChanges = true
	}
	if !data.UmbrellaSigAuthSecretWo.Equal(state.UmbrellaSigAuthSecretWo) {
		hasChanges = true
	}
	if !data.UmbrellaDnsAuthKey.Equal(state.UmbrellaDnsAuthKey) {
		hasChanges = true
	}
	if !data.UmbrellaDnsAuthKeyWo.Equal(state.UmbrellaDnsAuthKeyWo) {
		hasChanges = true
	}
	if !data.UmbrellaDnsAuthSecret.Equal(state.UmbrellaDnsAuthSecret) {
		hasChanges = true
	}
	if !data.UmbrellaDnsAuthSecretWo.Equal(state.UmbrellaDnsAuthSecretWo) {
		hasChanges = true
	}
	if !data.ZscalerOrganization.Equal(state.ZscalerOrganization) {
		hasChanges = true
	}
	if !data.ZscalerPartnerBaseUri.Equal(state.ZscalerPartnerBaseUri) {
		hasChanges = true
	}
	if !data.ZscalerPartnerKey.Equal(state.ZscalerPartnerKey) {
		hasChanges = true
	}
	if !data.ZscalerPartnerKeyWo.Equal(state.ZscalerPartnerKeyWo) {
		hasChanges = true
	}
	if !data.ZscalerUsername.Equal(state.ZscalerUsername) {
		hasChanges = true
	}
	if !data.ZscalerPassword.Equal(state.ZscalerPassword) {
		hasChanges = true
	}
	if !data.ZscalerPasswordWo.Equal(state.ZscalerPasswordWo) {
		hasChanges = true
	}
	if !data.CiscoSseOrgId.Equal(state.CiscoSseOrgId) {
		hasChanges = true
	}
	if !data.CiscoSseAuthKey.Equal(state.CiscoSseAuthKey) {
		hasChanges = true
	}
	if !data.CiscoSseAuthKeyWo.Equal(state.CiscoSseAuthKeyWo) {
		hasChanges = true
	}
	if !data.CiscoSseAuthSecret.Equal(state.CiscoSseAuthSecret) {
		hasChanges = true
	}
	if !data.CiscoSseAuthSecretWo.Equal(state.CiscoSseAuthSecretWo) {
		hasChanges = true
	}
	if !data.CiscoSseContextSharing.Equal(state.CiscoSseContextSharing) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *CloudProviderSettings) processImport(ctx context.Context) {
}

// End of section. //template:end processImport
