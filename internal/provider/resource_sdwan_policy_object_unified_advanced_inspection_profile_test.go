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
func TestAccSdwanPolicyObjectUnifiedAdvancedInspectionProfileProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_advanced_inspection_profile.test", "tls_decryption_action", "decrypt"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			{
				Config: testAccSdwanPolicyObjectUnifiedAdvancedInspectionProfilePrerequisitesProfileParcelConfig + testAccSdwanPolicyObjectUnifiedAdvancedInspectionProfileProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanPolicyObjectUnifiedAdvancedInspectionProfilePrerequisitesProfileParcelConfig = `
resource "sdwan_policy_object_feature_profile" "test" {
  name = "POLICY_OBJECT_FP_1"
  description = "My policy object feature profile 1"
}

resource "sdwan_policy_object_security_url_allow_list" "test" {
  name               = "TF_TEST_ALLOW"
  description        = "My Example"
  feature_profile_id = sdwan_policy_object_feature_profile.test.id
  entries = [
    {
      pattern = "www.cisco.com"
    }
  ]
}

resource "sdwan_policy_object_security_url_block_list" "test" {
  name               = "TF_TEST_BLOCK"
  description        = "My Example"
  feature_profile_id = sdwan_policy_object_feature_profile.test.id
  entries = [
    {
      pattern = "www.cisco.com"
    }
  ]
}

resource "sdwan_policy_object_unified_url_filtering" "test" {
  name                  = "TF_TEST_URL_FILTERING"
  description           = "My Example"
  feature_profile_id    = sdwan_policy_object_feature_profile.test.id
  web_categories_action = "block"
  web_categories        = ["confirmed-spam-sources"]
  web_reputation        = "suspicious"
  url_allow_list_id     = sdwan_policy_object_security_url_allow_list.test.id
  url_block_list_id     = sdwan_policy_object_security_url_block_list.test.id
  block_page_action     = "text"
  block_page_contents   = "Access to the requested page has been denied. Please contact your Network Administrator"
  redirect_url          = "www.example.com"
  enable_alerts         = true
  alerts                = ["blacklist"]
}

resource "sdwan_policy_object_security_ips_signature" "test" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = sdwan_policy_object_feature_profile.test.id
  entries = [
    {
      generator_id = "1234"
      signature_id = "5678"
    }
  ]
}

resource "sdwan_policy_object_unified_intrusion_prevention" "test" {
  name                  = "TF_TEST_INTRUSION"
  description           = "My Example"
  feature_profile_id    = sdwan_policy_object_feature_profile.test.id
  signature_set         = "balanced"
  inspection_mode       = "detection"
  ips_signature_list_id = sdwan_policy_object_security_ips_signature.test.id
  log_level             = "error"
  custom_signature      = false
}

resource "sdwan_policy_object_unified_advanced_malware_protection" "test" {
  name                          = "TF_TEST_ADVANCED_MALWARE"
  description                   = "My Example"
  feature_profile_id            = sdwan_policy_object_feature_profile.test.id
  amp_cloud_region              = "nam"
  amp_cloud_region_est_server   = "nam"
  alert_log_level               = "critical"
  file_analysis                 = true
  file_analysis_cloud_region    = "nam"
  file_analysis_file_types      = ["pdf"]
  file_analysis_alert_log_level = "critical"
}

resource "sdwan_policy_object_unified_tls_ssl_profile" "test" {
  name                    = "TF_TEST_TLS_SSL_PROFILE"
  description             = "My Example"
  feature_profile_id      = sdwan_policy_object_feature_profile.test.id
  decrypt_categories      = ["alcohol-and-tobacco"]
  no_decrypt_categories   = ["abortion"]
  pass_through_categories = ["auctions"]
  reputation              = true
  decrypt_threshold       = "moderate-risk"
  threshold_categories    = "moderate-risk"
  fail_decrypt            = true
  url_allow_list_id       = sdwan_policy_object_security_url_allow_list.test.id
  url_block_list_id       = sdwan_policy_object_security_url_block_list.test.id
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanPolicyObjectUnifiedAdvancedInspectionProfileProfileParcelConfig_all() string {
	config := `resource "sdwan_policy_object_unified_advanced_inspection_profile" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_policy_object_feature_profile.test.id` + "\n"
	config += `	tls_decryption_action = "decrypt"` + "\n"
	config += `	intrusion_prevention_list_id = sdwan_policy_object_unified_intrusion_prevention.test.id` + "\n"
	config += `	url_filtering_list_id = sdwan_policy_object_security_url_allow_list.test.id` + "\n"
	config += `	advanced_malware_protection_list_id = sdwan_policy_object_unified_advanced_malware_protection.test.id` + "\n"
	config += `	tls_ssl_profile_list_id = sdwan_policy_object_unified_tls_ssl_profile.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
