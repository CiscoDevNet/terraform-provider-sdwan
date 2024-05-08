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
type TLSSSLDecryptionPolicyDefinition struct {
	Id                          types.String                                   `tfsdk:"id"`
	Version                     types.Int64                                    `tfsdk:"version"`
	Name                        types.String                                   `tfsdk:"name"`
	Description                 types.String                                   `tfsdk:"description"`
	Mode                        types.String                                   `tfsdk:"mode"`
	DefaultAction               types.String                                   `tfsdk:"default_action"`
	NetworkRules                []TLSSSLDecryptionPolicyDefinitionNetworkRules `tfsdk:"network_rules"`
	UrlRules                    []TLSSSLDecryptionPolicyDefinitionUrlRules     `tfsdk:"url_rules"`
	SslDecryptionEnabled        types.String                                   `tfsdk:"ssl_decryption_enabled"`
	ExpiredCertificate          types.String                                   `tfsdk:"expired_certificate"`
	UntrustedCertificate        types.String                                   `tfsdk:"untrusted_certificate"`
	CertificateRevocationStatus types.String                                   `tfsdk:"certificate_revocation_status"`
	UnknownRevocationStatus     types.String                                   `tfsdk:"unknown_revocation_status"`
	UnsupportedProtocolVersions types.String                                   `tfsdk:"unsupported_protocol_versions"`
	UnsupportedCipherSuites     types.String                                   `tfsdk:"unsupported_cipher_suites"`
	FailureMode                 types.String                                   `tfsdk:"failure_mode"`
	RsaKeyPairModulus           types.String                                   `tfsdk:"rsa_key_pair_modulus"`
	EcKeyType                   types.String                                   `tfsdk:"ec_key_type"`
	CertificateLifetimeInDays   types.Int64                                    `tfsdk:"certificate_lifetime_in_days"`
	MinimalTlsVersion           types.String                                   `tfsdk:"minimal_tls_version"`
	UseDefaultCaCertBundle      types.Bool                                     `tfsdk:"use_default_ca_cert_bundle"`
}

type TLSSSLDecryptionPolicyDefinitionNetworkRules struct {
	BaseAction                        types.String                                                                    `tfsdk:"base_action"`
	RuleId                            types.Int64                                                                     `tfsdk:"rule_id"`
	RuleName                          types.String                                                                    `tfsdk:"rule_name"`
	RuleType                          types.String                                                                    `tfsdk:"rule_type"`
	SourceAndDestinationConfiguration []TLSSSLDecryptionPolicyDefinitionNetworkRulesSourceAndDestinationConfiguration `tfsdk:"source_and_destination_configuration"`
}

type TLSSSLDecryptionPolicyDefinitionUrlRules struct {
	RuleName              types.String `tfsdk:"rule_name"`
	TargetVpns            types.Set    `tfsdk:"target_vpns"`
	TlsSslProfilePolicyId types.String `tfsdk:"tls_ssl_profile_policy_id"`
	TlsSslProfileVersion  types.Int64  `tfsdk:"tls_ssl_profile_version"`
}

type TLSSSLDecryptionPolicyDefinitionNetworkRulesSourceAndDestinationConfiguration struct {
	Option types.String `tfsdk:"option"`
	Value  types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TLSSSLDecryptionPolicyDefinition) getPath() string {
	return "/template/policy/definition/ssldecryption/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TLSSSLDecryptionPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "sslDecryption")
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
	if !data.DefaultAction.IsNull() {
		body, _ = sjson.Set(body, "definition.defaultAction.type", data.DefaultAction.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "definition.sequences", []interface{}{})
		for _, item := range data.NetworkRules {
			itemBody := ""
			if !item.BaseAction.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "baseAction", item.BaseAction.ValueString())
			}
			if !item.RuleId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceId", item.RuleId.ValueInt64())
			}
			if !item.RuleName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceName", item.RuleName.ValueString())
			}
			if !item.RuleType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequenceType", item.RuleType.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "match.entries", []interface{}{})
				for _, childItem := range item.SourceAndDestinationConfiguration {
					itemChildBody := ""
					if !childItem.Option.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "field", childItem.Option.ValueString())
					}
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "match.entries.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "definition.sequences.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, "definition.profiles", []interface{}{})
		for _, item := range data.UrlRules {
			itemBody := ""
			if !item.RuleName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.RuleName.ValueString())
			}
			if !item.TargetVpns.IsNull() {
				var values []string
				item.TargetVpns.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "vpn", values)
			}
			if !item.TlsSslProfilePolicyId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ref", item.TlsSslProfilePolicyId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "definition.profiles.-1", itemBody)
		}
	}
	if !data.SslDecryptionEnabled.IsNull() {
		body, _ = sjson.Set(body, "definition.settings.sslEnable", data.SslDecryptionEnabled.ValueString())
	}
	if !data.ExpiredCertificate.IsNull() {
		body, _ = sjson.Set(body, "definition.settings.expiredCertificate", data.ExpiredCertificate.ValueString())
	}
	if !data.UntrustedCertificate.IsNull() {
		body, _ = sjson.Set(body, "definition.settings.untrustedCertificate", data.UntrustedCertificate.ValueString())
	}
	if !data.CertificateRevocationStatus.IsNull() {
		body, _ = sjson.Set(body, "definition.settings.certificateRevocationStatus", data.CertificateRevocationStatus.ValueString())
	}
	if !data.UnknownRevocationStatus.IsNull() {
		body, _ = sjson.Set(body, "definition.settings.unknownStatus", data.UnknownRevocationStatus.ValueString())
	}
	if !data.UnsupportedProtocolVersions.IsNull() {
		body, _ = sjson.Set(body, "definition.settings.unsupportedProtocolVersions", data.UnsupportedProtocolVersions.ValueString())
	}
	if !data.UnsupportedCipherSuites.IsNull() {
		body, _ = sjson.Set(body, "definition.settings.unsupportedCipherSuites", data.UnsupportedCipherSuites.ValueString())
	}
	if !data.FailureMode.IsNull() {
		body, _ = sjson.Set(body, "definition.settings.failureMode", data.FailureMode.ValueString())
	}
	if !data.RsaKeyPairModulus.IsNull() {
		body, _ = sjson.Set(body, "definition.settings.keyModulus", data.RsaKeyPairModulus.ValueString())
	}
	if !data.EcKeyType.IsNull() {
		body, _ = sjson.Set(body, "definition.settings.eckeyType", data.EcKeyType.ValueString())
	}
	if !data.CertificateLifetimeInDays.IsNull() {
		body, _ = sjson.Set(body, "definition.settings.certificateLifetime", data.CertificateLifetimeInDays.ValueInt64())
	}
	if !data.MinimalTlsVersion.IsNull() {
		body, _ = sjson.Set(body, "definition.settings.minTlsVer", data.MinimalTlsVersion.ValueString())
	}
	if !data.UseDefaultCaCertBundle.IsNull() {
		if false && data.UseDefaultCaCertBundle.ValueBool() {
			body, _ = sjson.Set(body, "definition.settings.caCertBundle.default", "")
		} else {
			body, _ = sjson.Set(body, "definition.settings.caCertBundle.default", data.UseDefaultCaCertBundle.ValueBool())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TLSSSLDecryptionPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("definition.defaultAction.type"); value.Exists() {
		data.DefaultAction = types.StringValue(value.String())
	} else {
		data.DefaultAction = types.StringNull()
	}
	if value := res.Get("definition.sequences"); value.Exists() && len(value.Array()) > 0 {
		data.NetworkRules = make([]TLSSSLDecryptionPolicyDefinitionNetworkRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TLSSSLDecryptionPolicyDefinitionNetworkRules{}
			if cValue := v.Get("baseAction"); cValue.Exists() {
				item.BaseAction = types.StringValue(cValue.String())
			} else {
				item.BaseAction = types.StringNull()
			}
			if cValue := v.Get("sequenceId"); cValue.Exists() {
				item.RuleId = types.Int64Value(cValue.Int())
			} else {
				item.RuleId = types.Int64Null()
			}
			if cValue := v.Get("sequenceName"); cValue.Exists() {
				item.RuleName = types.StringValue(cValue.String())
			} else {
				item.RuleName = types.StringNull()
			}
			if cValue := v.Get("sequenceType"); cValue.Exists() {
				item.RuleType = types.StringValue(cValue.String())
			} else {
				item.RuleType = types.StringNull()
			}
			if cValue := v.Get("match.entries"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.SourceAndDestinationConfiguration = make([]TLSSSLDecryptionPolicyDefinitionNetworkRulesSourceAndDestinationConfiguration, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TLSSSLDecryptionPolicyDefinitionNetworkRulesSourceAndDestinationConfiguration{}
					if ccValue := cv.Get("field"); ccValue.Exists() {
						cItem.Option = types.StringValue(ccValue.String())
					} else {
						cItem.Option = types.StringNull()
					}
					if ccValue := cv.Get("value"); ccValue.Exists() {
						cItem.Value = types.StringValue(ccValue.String())
					} else {
						cItem.Value = types.StringNull()
					}
					item.SourceAndDestinationConfiguration = append(item.SourceAndDestinationConfiguration, cItem)
					return true
				})
			} else {
				if len(item.SourceAndDestinationConfiguration) > 0 {
					item.SourceAndDestinationConfiguration = []TLSSSLDecryptionPolicyDefinitionNetworkRulesSourceAndDestinationConfiguration{}
				}
			}
			data.NetworkRules = append(data.NetworkRules, item)
			return true
		})
	} else {
		if len(data.NetworkRules) > 0 {
			data.NetworkRules = []TLSSSLDecryptionPolicyDefinitionNetworkRules{}
		}
	}
	if value := res.Get("definition.profiles"); value.Exists() && len(value.Array()) > 0 {
		data.UrlRules = make([]TLSSSLDecryptionPolicyDefinitionUrlRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TLSSSLDecryptionPolicyDefinitionUrlRules{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.RuleName = types.StringValue(cValue.String())
			} else {
				item.RuleName = types.StringNull()
			}
			if cValue := v.Get("vpn"); cValue.Exists() {
				item.TargetVpns = helpers.GetStringSet(cValue.Array())
			} else {
				item.TargetVpns = types.SetNull(types.StringType)
			}
			if cValue := v.Get("ref"); cValue.Exists() {
				item.TlsSslProfilePolicyId = types.StringValue(cValue.String())
			} else {
				item.TlsSslProfilePolicyId = types.StringNull()
			}
			data.UrlRules = append(data.UrlRules, item)
			return true
		})
	} else {
		if len(data.UrlRules) > 0 {
			data.UrlRules = []TLSSSLDecryptionPolicyDefinitionUrlRules{}
		}
	}
	if value := res.Get("definition.settings.sslEnable"); value.Exists() {
		data.SslDecryptionEnabled = types.StringValue(value.String())
	} else {
		data.SslDecryptionEnabled = types.StringNull()
	}
	if value := res.Get("definition.settings.expiredCertificate"); value.Exists() {
		data.ExpiredCertificate = types.StringValue(value.String())
	} else {
		data.ExpiredCertificate = types.StringNull()
	}
	if value := res.Get("definition.settings.untrustedCertificate"); value.Exists() {
		data.UntrustedCertificate = types.StringValue(value.String())
	} else {
		data.UntrustedCertificate = types.StringNull()
	}
	if value := res.Get("definition.settings.certificateRevocationStatus"); value.Exists() {
		data.CertificateRevocationStatus = types.StringValue(value.String())
	} else {
		data.CertificateRevocationStatus = types.StringNull()
	}
	if value := res.Get("definition.settings.unknownStatus"); value.Exists() {
		data.UnknownRevocationStatus = types.StringValue(value.String())
	} else {
		data.UnknownRevocationStatus = types.StringNull()
	}
	if value := res.Get("definition.settings.unsupportedProtocolVersions"); value.Exists() {
		data.UnsupportedProtocolVersions = types.StringValue(value.String())
	} else {
		data.UnsupportedProtocolVersions = types.StringNull()
	}
	if value := res.Get("definition.settings.unsupportedCipherSuites"); value.Exists() {
		data.UnsupportedCipherSuites = types.StringValue(value.String())
	} else {
		data.UnsupportedCipherSuites = types.StringNull()
	}
	if value := res.Get("definition.settings.failureMode"); value.Exists() {
		data.FailureMode = types.StringValue(value.String())
	} else {
		data.FailureMode = types.StringNull()
	}
	if value := res.Get("definition.settings.keyModulus"); value.Exists() {
		data.RsaKeyPairModulus = types.StringValue(value.String())
	} else {
		data.RsaKeyPairModulus = types.StringNull()
	}
	if value := res.Get("definition.settings.eckeyType"); value.Exists() {
		data.EcKeyType = types.StringValue(value.String())
	} else {
		data.EcKeyType = types.StringNull()
	}
	if value := res.Get("definition.settings.certificateLifetime"); value.Exists() {
		data.CertificateLifetimeInDays = types.Int64Value(value.Int())
	} else {
		data.CertificateLifetimeInDays = types.Int64Null()
	}
	if value := res.Get("definition.settings.minTlsVer"); value.Exists() {
		data.MinimalTlsVersion = types.StringValue(value.String())
	} else {
		data.MinimalTlsVersion = types.StringNull()
	}
	if value := res.Get("definition.settings.caCertBundle.default"); value.Exists() {
		if false && value.String() == "" {
			data.UseDefaultCaCertBundle = types.BoolValue(true)
		} else {
			data.UseDefaultCaCertBundle = types.BoolValue(value.Bool())
		}
	} else {
		data.UseDefaultCaCertBundle = types.BoolNull()
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *TLSSSLDecryptionPolicyDefinition) hasChanges(ctx context.Context, state *TLSSSLDecryptionPolicyDefinition) bool {
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
	if !data.DefaultAction.Equal(state.DefaultAction) {
		hasChanges = true
	}
	if len(data.NetworkRules) != len(state.NetworkRules) {
		hasChanges = true
	} else {
		for i := range data.NetworkRules {
			if !data.NetworkRules[i].BaseAction.Equal(state.NetworkRules[i].BaseAction) {
				hasChanges = true
			}
			if !data.NetworkRules[i].RuleId.Equal(state.NetworkRules[i].RuleId) {
				hasChanges = true
			}
			if !data.NetworkRules[i].RuleName.Equal(state.NetworkRules[i].RuleName) {
				hasChanges = true
			}
			if !data.NetworkRules[i].RuleType.Equal(state.NetworkRules[i].RuleType) {
				hasChanges = true
			}
			if len(data.NetworkRules[i].SourceAndDestinationConfiguration) != len(state.NetworkRules[i].SourceAndDestinationConfiguration) {
				hasChanges = true
			} else {
				for ii := range data.NetworkRules[i].SourceAndDestinationConfiguration {
					if !data.NetworkRules[i].SourceAndDestinationConfiguration[ii].Option.Equal(state.NetworkRules[i].SourceAndDestinationConfiguration[ii].Option) {
						hasChanges = true
					}
					if !data.NetworkRules[i].SourceAndDestinationConfiguration[ii].Value.Equal(state.NetworkRules[i].SourceAndDestinationConfiguration[ii].Value) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.UrlRules) != len(state.UrlRules) {
		hasChanges = true
	} else {
		for i := range data.UrlRules {
			if !data.UrlRules[i].RuleName.Equal(state.UrlRules[i].RuleName) {
				hasChanges = true
			}
			if !data.UrlRules[i].TargetVpns.Equal(state.UrlRules[i].TargetVpns) {
				hasChanges = true
			}
			if !data.UrlRules[i].TlsSslProfilePolicyId.Equal(state.UrlRules[i].TlsSslProfilePolicyId) {
				hasChanges = true
			}
		}
	}
	if !data.SslDecryptionEnabled.Equal(state.SslDecryptionEnabled) {
		hasChanges = true
	}
	if !data.ExpiredCertificate.Equal(state.ExpiredCertificate) {
		hasChanges = true
	}
	if !data.UntrustedCertificate.Equal(state.UntrustedCertificate) {
		hasChanges = true
	}
	if !data.CertificateRevocationStatus.Equal(state.CertificateRevocationStatus) {
		hasChanges = true
	}
	if !data.UnknownRevocationStatus.Equal(state.UnknownRevocationStatus) {
		hasChanges = true
	}
	if !data.UnsupportedProtocolVersions.Equal(state.UnsupportedProtocolVersions) {
		hasChanges = true
	}
	if !data.UnsupportedCipherSuites.Equal(state.UnsupportedCipherSuites) {
		hasChanges = true
	}
	if !data.FailureMode.Equal(state.FailureMode) {
		hasChanges = true
	}
	if !data.RsaKeyPairModulus.Equal(state.RsaKeyPairModulus) {
		hasChanges = true
	}
	if !data.EcKeyType.Equal(state.EcKeyType) {
		hasChanges = true
	}
	if !data.CertificateLifetimeInDays.Equal(state.CertificateLifetimeInDays) {
		hasChanges = true
	}
	if !data.MinimalTlsVersion.Equal(state.MinimalTlsVersion) {
		hasChanges = true
	}
	if !data.UseDefaultCaCertBundle.Equal(state.UseDefaultCaCertBundle) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *TLSSSLDecryptionPolicyDefinition) updateVersions(ctx context.Context, state *TLSSSLDecryptionPolicyDefinition) {
	for i := range data.UrlRules {
		dataKeys := [...]string{}
		stateIndex := -1
		for j := range state.UrlRules {
			stateKeys := [...]string{}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		if stateIndex > -1 {
			data.UrlRules[i].TlsSslProfileVersion = state.UrlRules[stateIndex].TlsSslProfileVersion
		} else {
			data.UrlRules[i].TlsSslProfileVersion = types.Int64Null()
		}
	}
}

// End of section. //template:end updateVersions
