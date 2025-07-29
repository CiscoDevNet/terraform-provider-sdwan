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
type AdvancedInspectionProfilePolicyDefinition struct {
	Id                               types.String `tfsdk:"id"`
	Version                          types.Int64  `tfsdk:"version"`
	Type                             types.String `tfsdk:"type"`
	Name                             types.String `tfsdk:"name"`
	Description                      types.String `tfsdk:"description"`
	TlsAction                        types.String `tfsdk:"tls_action"`
	IntrusionPreventionId            types.String `tfsdk:"intrusion_prevention_id"`
	IntrusionPreventionVersion       types.Int64  `tfsdk:"intrusion_prevention_version"`
	UrlFilteringId                   types.String `tfsdk:"url_filtering_id"`
	UrlFilteringVersion              types.Int64  `tfsdk:"url_filtering_version"`
	AdvancedMalwareProtectionId      types.String `tfsdk:"advanced_malware_protection_id"`
	AdvancedMalwareProtectionVersion types.Int64  `tfsdk:"advanced_malware_protection_version"`
	TlsSslDecryptionId               types.String `tfsdk:"tls_ssl_decryption_id"`
	TlsSslDecryptionVersion          types.Int64  `tfsdk:"tls_ssl_decryption_version"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data AdvancedInspectionProfilePolicyDefinition) getPath() string {
	return "/template/policy/definition/advancedinspectionprofile/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data AdvancedInspectionProfilePolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "advancedInspectionProfile")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.TlsAction.IsNull() {
		body, _ = sjson.Set(body, "definition.tlsDecryptionAction", data.TlsAction.ValueString())
	}
	if !data.IntrusionPreventionId.IsNull() {
		body, _ = sjson.Set(body, "definition.intrusionPrevention.ref", data.IntrusionPreventionId.ValueString())
	}
	if !data.UrlFilteringId.IsNull() {
		body, _ = sjson.Set(body, "definition.urlFiltering.ref", data.UrlFilteringId.ValueString())
	}
	if !data.AdvancedMalwareProtectionId.IsNull() {
		body, _ = sjson.Set(body, "definition.advancedMalwareProtection.ref", data.AdvancedMalwareProtectionId.ValueString())
	}
	if !data.TlsSslDecryptionId.IsNull() {
		body, _ = sjson.Set(body, "definition.sslDecryption.ref", data.TlsSslDecryptionId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *AdvancedInspectionProfilePolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("definition.tlsDecryptionAction"); value.Exists() {
		data.TlsAction = types.StringValue(value.String())
	} else {
		data.TlsAction = types.StringNull()
	}
	if value := res.Get("definition.intrusionPrevention.ref"); value.Exists() {
		data.IntrusionPreventionId = types.StringValue(value.String())
	} else {
		data.IntrusionPreventionId = types.StringNull()
	}
	if value := res.Get("definition.urlFiltering.ref"); value.Exists() {
		data.UrlFilteringId = types.StringValue(value.String())
	} else {
		data.UrlFilteringId = types.StringNull()
	}
	if value := res.Get("definition.advancedMalwareProtection.ref"); value.Exists() {
		data.AdvancedMalwareProtectionId = types.StringValue(value.String())
	} else {
		data.AdvancedMalwareProtectionId = types.StringNull()
	}
	if value := res.Get("definition.sslDecryption.ref"); value.Exists() {
		data.TlsSslDecryptionId = types.StringValue(value.String())
	} else {
		data.TlsSslDecryptionId = types.StringNull()
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *AdvancedInspectionProfilePolicyDefinition) hasChanges(ctx context.Context, state *AdvancedInspectionProfilePolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.TlsAction.Equal(state.TlsAction) {
		hasChanges = true
	}
	if !data.IntrusionPreventionId.Equal(state.IntrusionPreventionId) {
		hasChanges = true
	}
	if !data.UrlFilteringId.Equal(state.UrlFilteringId) {
		hasChanges = true
	}
	if !data.AdvancedMalwareProtectionId.Equal(state.AdvancedMalwareProtectionId) {
		hasChanges = true
	}
	if !data.TlsSslDecryptionId.Equal(state.TlsSslDecryptionId) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *AdvancedInspectionProfilePolicyDefinition) updateVersions(ctx context.Context, state *AdvancedInspectionProfilePolicyDefinition) {
	data.IntrusionPreventionVersion = state.IntrusionPreventionVersion
	data.UrlFilteringVersion = state.UrlFilteringVersion
	data.AdvancedMalwareProtectionVersion = state.AdvancedMalwareProtectionVersion
	data.TlsSslDecryptionVersion = state.TlsSslDecryptionVersion
}

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *AdvancedInspectionProfilePolicyDefinition) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
	data.Type = types.StringValue("advancedInspectionProfile")
	if data.IntrusionPreventionId != types.StringNull() {
		data.IntrusionPreventionVersion = types.Int64Value(0)
	}
	if data.UrlFilteringId != types.StringNull() {
		data.UrlFilteringVersion = types.Int64Value(0)
	}
	if data.AdvancedMalwareProtectionId != types.StringNull() {
		data.AdvancedMalwareProtectionVersion = types.Int64Value(0)
	}
	if data.TlsSslDecryptionId != types.StringNull() {
		data.TlsSslDecryptionVersion = types.Int64Value(0)
	}
}

// End of section. //template:end processImport
