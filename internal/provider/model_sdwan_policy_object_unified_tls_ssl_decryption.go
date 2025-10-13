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
type PolicyObjectUnifiedTLSSSLDecryption struct {
	Id                          types.String `tfsdk:"id"`
	Version                     types.Int64  `tfsdk:"version"`
	Name                        types.String `tfsdk:"name"`
	Description                 types.String `tfsdk:"description"`
	FeatureProfileId            types.String `tfsdk:"feature_profile_id"`
	ExpiredCertificate          types.String `tfsdk:"expired_certificate"`
	UntrustedCertificate        types.String `tfsdk:"untrusted_certificate"`
	CertificateRevocationStatus types.String `tfsdk:"certificate_revocation_status"`
	UnknownRevocationStatus     types.String `tfsdk:"unknown_revocation_status"`
	UnsupportedProtocolVersions types.String `tfsdk:"unsupported_protocol_versions"`
	UnsupportedCipherSuites     types.String `tfsdk:"unsupported_cipher_suites"`
	FailureMode                 types.String `tfsdk:"failure_mode"`
	DefaultCaCertificateBundle  types.Bool   `tfsdk:"default_ca_certificate_bundle"`
	FileName                    types.String `tfsdk:"file_name"`
	BundleString                types.String `tfsdk:"bundle_string"`
	RsaKeypairModules           types.String `tfsdk:"rsa_keypair_modules"`
	EcKeyType                   types.String `tfsdk:"ec_key_type"`
	CertificateLifetime         types.String `tfsdk:"certificate_lifetime"`
	MinimalTlsVer               types.String `tfsdk:"minimal_tls_ver"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectUnifiedTLSSSLDecryption) getModel() string {
	return "policy_object_unified_tls_ssl_decryption"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectUnifiedTLSSSLDecryption) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/unified/unified/ssl-decryption", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectUnifiedTLSSSLDecryption) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {
		body, _ = sjson.Set(body, path+"sslEnable.optionType", "default")
		body, _ = sjson.Set(body, path+"sslEnable.value", true)
	}
	if !data.ExpiredCertificate.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"expiredCertificate.optionType", "global")
			body, _ = sjson.Set(body, path+"expiredCertificate.value", data.ExpiredCertificate.ValueString())
		}
	}
	if !data.UntrustedCertificate.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"untrustedCertificate.optionType", "global")
			body, _ = sjson.Set(body, path+"untrustedCertificate.value", data.UntrustedCertificate.ValueString())
		}
	}
	if !data.CertificateRevocationStatus.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"certificateRevocationStatus.optionType", "global")
			body, _ = sjson.Set(body, path+"certificateRevocationStatus.value", data.CertificateRevocationStatus.ValueString())
		}
	}
	if !data.UnknownRevocationStatus.IsNull() {
		if true && data.CertificateRevocationStatus.ValueString() == "ocsp" {
			body, _ = sjson.Set(body, path+"unknownStatus.optionType", "global")
			body, _ = sjson.Set(body, path+"unknownStatus.value", data.UnknownRevocationStatus.ValueString())
		}
	}
	if !data.UnsupportedProtocolVersions.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"unsupportedProtocolVersions.optionType", "global")
			body, _ = sjson.Set(body, path+"unsupportedProtocolVersions.value", data.UnsupportedProtocolVersions.ValueString())
		}
	}
	if !data.UnsupportedCipherSuites.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"unsupportedCipherSuites.optionType", "global")
			body, _ = sjson.Set(body, path+"unsupportedCipherSuites.value", data.UnsupportedCipherSuites.ValueString())
		}
	}
	if !data.FailureMode.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"failureMode.optionType", "global")
			body, _ = sjson.Set(body, path+"failureMode.value", data.FailureMode.ValueString())
		}
	}
	if !data.DefaultCaCertificateBundle.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"caCertBundle.default.optionType", "global")
			body, _ = sjson.Set(body, path+"caCertBundle.default.value", data.DefaultCaCertificateBundle.ValueBool())
		}
	}
	if !data.FileName.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"caCertBundle.fileName.optionType", "global")
			body, _ = sjson.Set(body, path+"caCertBundle.fileName.value", data.FileName.ValueString())
		}
	}
	if !data.BundleString.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"caCertBundle.bundleString.optionType", "global")
			body, _ = sjson.Set(body, path+"caCertBundle.bundleString.value", data.BundleString.ValueString())
		}
	}
	if !data.RsaKeypairModules.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"keyModulus.optionType", "global")
			body, _ = sjson.Set(body, path+"keyModulus.value", data.RsaKeypairModules.ValueString())
		}
	}
	if !data.EcKeyType.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"eckeyType.optionType", "global")
			body, _ = sjson.Set(body, path+"eckeyType.value", data.EcKeyType.ValueString())
		}
	}
	if !data.CertificateLifetime.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"certificateLifetime.optionType", "global")
			body, _ = sjson.Set(body, path+"certificateLifetime.value", data.CertificateLifetime.ValueString())
		}
	}
	if !data.MinimalTlsVer.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"minTlsVer.optionType", "global")
			body, _ = sjson.Set(body, path+"minTlsVer.value", data.MinimalTlsVer.ValueString())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"caTpLabel.optionType", "default")
		body, _ = sjson.Set(body, path+"caTpLabel.value", "PROXY-SIGNING-CA")
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectUnifiedTLSSSLDecryption) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.ExpiredCertificate = types.StringNull()

	if t := res.Get(path + "expiredCertificate.optionType"); t.Exists() {
		va := res.Get(path + "expiredCertificate.value")
		if t.String() == "global" {
			data.ExpiredCertificate = types.StringValue(va.String())
		}
	}
	data.UntrustedCertificate = types.StringNull()

	if t := res.Get(path + "untrustedCertificate.optionType"); t.Exists() {
		va := res.Get(path + "untrustedCertificate.value")
		if t.String() == "global" {
			data.UntrustedCertificate = types.StringValue(va.String())
		}
	}
	data.CertificateRevocationStatus = types.StringNull()

	if t := res.Get(path + "certificateRevocationStatus.optionType"); t.Exists() {
		va := res.Get(path + "certificateRevocationStatus.value")
		if t.String() == "global" {
			data.CertificateRevocationStatus = types.StringValue(va.String())
		}
	}
	data.UnknownRevocationStatus = types.StringNull()

	if t := res.Get(path + "unknownStatus.optionType"); t.Exists() {
		va := res.Get(path + "unknownStatus.value")
		if t.String() == "global" {
			data.UnknownRevocationStatus = types.StringValue(va.String())
		}
		data.CertificateRevocationStatus = types.StringValue("ocsp")
	}
	data.UnsupportedProtocolVersions = types.StringNull()

	if t := res.Get(path + "unsupportedProtocolVersions.optionType"); t.Exists() {
		va := res.Get(path + "unsupportedProtocolVersions.value")
		if t.String() == "global" {
			data.UnsupportedProtocolVersions = types.StringValue(va.String())
		}
	}
	data.UnsupportedCipherSuites = types.StringNull()

	if t := res.Get(path + "unsupportedCipherSuites.optionType"); t.Exists() {
		va := res.Get(path + "unsupportedCipherSuites.value")
		if t.String() == "global" {
			data.UnsupportedCipherSuites = types.StringValue(va.String())
		}
	}
	data.FailureMode = types.StringNull()

	if t := res.Get(path + "failureMode.optionType"); t.Exists() {
		va := res.Get(path + "failureMode.value")
		if t.String() == "global" {
			data.FailureMode = types.StringValue(va.String())
		}
	}
	data.DefaultCaCertificateBundle = types.BoolNull()

	if t := res.Get(path + "caCertBundle.default.optionType"); t.Exists() {
		va := res.Get(path + "caCertBundle.default.value")
		if t.String() == "global" {
			data.DefaultCaCertificateBundle = types.BoolValue(va.Bool())
		}
	}
	data.FileName = types.StringNull()

	if t := res.Get(path + "caCertBundle.fileName.optionType"); t.Exists() {
		va := res.Get(path + "caCertBundle.fileName.value")
		if t.String() == "global" {
			data.FileName = types.StringValue(va.String())
		}
	}
	data.BundleString = types.StringNull()

	if t := res.Get(path + "caCertBundle.bundleString.optionType"); t.Exists() {
		va := res.Get(path + "caCertBundle.bundleString.value")
		if t.String() == "global" {
			data.BundleString = types.StringValue(va.String())
		}
	}
	data.RsaKeypairModules = types.StringNull()

	if t := res.Get(path + "keyModulus.optionType"); t.Exists() {
		va := res.Get(path + "keyModulus.value")
		if t.String() == "global" {
			data.RsaKeypairModules = types.StringValue(va.String())
		}
	}
	data.EcKeyType = types.StringNull()

	if t := res.Get(path + "eckeyType.optionType"); t.Exists() {
		va := res.Get(path + "eckeyType.value")
		if t.String() == "global" {
			data.EcKeyType = types.StringValue(va.String())
		}
	}
	data.CertificateLifetime = types.StringNull()

	if t := res.Get(path + "certificateLifetime.optionType"); t.Exists() {
		va := res.Get(path + "certificateLifetime.value")
		if t.String() == "global" {
			data.CertificateLifetime = types.StringValue(va.String())
		}
	}
	data.MinimalTlsVer = types.StringNull()

	if t := res.Get(path + "minTlsVer.optionType"); t.Exists() {
		va := res.Get(path + "minTlsVer.value")
		if t.String() == "global" {
			data.MinimalTlsVer = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PolicyObjectUnifiedTLSSSLDecryption) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.ExpiredCertificate = types.StringNull()

	if t := res.Get(path + "expiredCertificate.optionType"); t.Exists() {
		va := res.Get(path + "expiredCertificate.value")
		if t.String() == "global" {
			data.ExpiredCertificate = types.StringValue(va.String())
		}
	}
	data.UntrustedCertificate = types.StringNull()

	if t := res.Get(path + "untrustedCertificate.optionType"); t.Exists() {
		va := res.Get(path + "untrustedCertificate.value")
		if t.String() == "global" {
			data.UntrustedCertificate = types.StringValue(va.String())
		}
	}
	data.CertificateRevocationStatus = types.StringNull()

	if t := res.Get(path + "certificateRevocationStatus.optionType"); t.Exists() {
		va := res.Get(path + "certificateRevocationStatus.value")
		if t.String() == "global" {
			data.CertificateRevocationStatus = types.StringValue(va.String())
		}
	}
	data.UnknownRevocationStatus = types.StringNull()

	if t := res.Get(path + "unknownStatus.optionType"); t.Exists() {
		va := res.Get(path + "unknownStatus.value")
		if t.String() == "global" {
			data.UnknownRevocationStatus = types.StringValue(va.String())
		}
	}
	data.UnsupportedProtocolVersions = types.StringNull()

	if t := res.Get(path + "unsupportedProtocolVersions.optionType"); t.Exists() {
		va := res.Get(path + "unsupportedProtocolVersions.value")
		if t.String() == "global" {
			data.UnsupportedProtocolVersions = types.StringValue(va.String())
		}
	}
	data.UnsupportedCipherSuites = types.StringNull()

	if t := res.Get(path + "unsupportedCipherSuites.optionType"); t.Exists() {
		va := res.Get(path + "unsupportedCipherSuites.value")
		if t.String() == "global" {
			data.UnsupportedCipherSuites = types.StringValue(va.String())
		}
	}
	data.FailureMode = types.StringNull()

	if t := res.Get(path + "failureMode.optionType"); t.Exists() {
		va := res.Get(path + "failureMode.value")
		if t.String() == "global" {
			data.FailureMode = types.StringValue(va.String())
		}
	}
	data.DefaultCaCertificateBundle = types.BoolNull()

	if t := res.Get(path + "caCertBundle.default.optionType"); t.Exists() {
		va := res.Get(path + "caCertBundle.default.value")
		if t.String() == "global" {
			data.DefaultCaCertificateBundle = types.BoolValue(va.Bool())
		}
	}
	data.FileName = types.StringNull()

	if t := res.Get(path + "caCertBundle.fileName.optionType"); t.Exists() {
		va := res.Get(path + "caCertBundle.fileName.value")
		if t.String() == "global" {
			data.FileName = types.StringValue(va.String())
		}
	}
	data.BundleString = types.StringNull()

	if t := res.Get(path + "caCertBundle.bundleString.optionType"); t.Exists() {
		va := res.Get(path + "caCertBundle.bundleString.value")
		if t.String() == "global" {
			data.BundleString = types.StringValue(va.String())
		}
	}
	data.RsaKeypairModules = types.StringNull()

	if t := res.Get(path + "keyModulus.optionType"); t.Exists() {
		va := res.Get(path + "keyModulus.value")
		if t.String() == "global" {
			data.RsaKeypairModules = types.StringValue(va.String())
		}
	}
	data.EcKeyType = types.StringNull()

	if t := res.Get(path + "eckeyType.optionType"); t.Exists() {
		va := res.Get(path + "eckeyType.value")
		if t.String() == "global" {
			data.EcKeyType = types.StringValue(va.String())
		}
	}
	data.CertificateLifetime = types.StringNull()

	if t := res.Get(path + "certificateLifetime.optionType"); t.Exists() {
		va := res.Get(path + "certificateLifetime.value")
		if t.String() == "global" {
			data.CertificateLifetime = types.StringValue(va.String())
		}
	}
	data.MinimalTlsVer = types.StringNull()

	if t := res.Get(path + "minTlsVer.optionType"); t.Exists() {
		va := res.Get(path + "minTlsVer.value")
		if t.String() == "global" {
			data.MinimalTlsVer = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end updateFromBody
