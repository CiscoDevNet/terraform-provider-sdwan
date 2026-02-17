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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanPolicyObjectUnifiedTLSSSLDecryptionProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_tls_ssl_decryption.test", "expired_certificate", "drop"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_tls_ssl_decryption.test", "untrusted_certificate", "drop"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_tls_ssl_decryption.test", "certificate_revocation_status", "none"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_tls_ssl_decryption.test", "unsupported_protocol_versions", "drop"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_tls_ssl_decryption.test", "unsupported_cipher_suites", "drop"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_tls_ssl_decryption.test", "failure_mode", "close"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_tls_ssl_decryption.test", "default_ca_certificate_bundle", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_tls_ssl_decryption.test", "rsa_keypair_modules", "2048"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_tls_ssl_decryption.test", "ec_key_type", "P256"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_tls_ssl_decryption.test", "certificate_lifetime", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_tls_ssl_decryption.test", "minimal_tls_ver", "TLSv1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			{
				Config: testAccSdwanPolicyObjectUnifiedTLSSSLDecryptionPrerequisitesProfileParcelConfig + testAccSdwanPolicyObjectUnifiedTLSSSLDecryptionProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanPolicyObjectUnifiedTLSSSLDecryptionPrerequisitesProfileParcelConfig = `
resource "sdwan_policy_object_feature_profile" "test" {
  name = "POLICY_OBJECT_FP_1"
  description = "My policy object feature profile 1"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanPolicyObjectUnifiedTLSSSLDecryptionProfileParcelConfig_all() string {
	config := `resource "sdwan_policy_object_unified_tls_ssl_decryption" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_policy_object_feature_profile.test.id` + "\n"
	config += `	expired_certificate = "drop"` + "\n"
	config += `	untrusted_certificate = "drop"` + "\n"
	config += `	certificate_revocation_status = "none"` + "\n"
	config += `	unsupported_protocol_versions = "drop"` + "\n"
	config += `	unsupported_cipher_suites = "drop"` + "\n"
	config += `	failure_mode = "close"` + "\n"
	config += `	default_ca_certificate_bundle = true` + "\n"
	config += `	rsa_keypair_modules = "2048"` + "\n"
	config += `	ec_key_type = "P256"` + "\n"
	config += `	certificate_lifetime = "1"` + "\n"
	config += `	minimal_tls_ver = "TLSv1"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
