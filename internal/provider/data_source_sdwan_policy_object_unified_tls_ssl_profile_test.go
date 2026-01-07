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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceSdwanPolicyObjectUnifiedTLSSSLProfileProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015_IN_PROGRESS") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015_IN_PROGRESS")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_unified_tls_ssl_profile.test", "reputation", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_unified_tls_ssl_profile.test", "decrypt_threshold", "moderate-risk"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_unified_tls_ssl_profile.test", "threshold_categories", "moderate-risk"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_unified_tls_ssl_profile.test", "fail_decrypt", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanPolicyObjectUnifiedTLSSSLProfilePrerequisitesProfileParcelConfig + testAccDataSourceSdwanPolicyObjectUnifiedTLSSSLProfileProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanPolicyObjectUnifiedTLSSSLProfilePrerequisitesProfileParcelConfig = `
resource "sdwan_policy_object_feature_profile" "test" {
  name = "POLICY_OBJECT_FP_1"
  description = "My policy object feature profile 1"
}

resource "sdwan_policy_object_security_url_block_list" "test" {
  name               = "TF_TEST_ALLOW"
  description        = "My Example"
  feature_profile_id = sdwan_policy_object_feature_profile.test.id
  entries = [
    {
      pattern = "www.cisco.com"
    }
  ]
}

resource "sdwan_policy_object_security_url_allow_list" "test" {
  name               = "TF_TEST_BLOCK"
  description        = "My Example"
  feature_profile_id = sdwan_policy_object_feature_profile.test.id
  entries = [
    {
      pattern = "www.cisco.com"
    }
  ]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanPolicyObjectUnifiedTLSSSLProfileProfileParcelConfig() string {
	config := `resource "sdwan_policy_object_unified_tls_ssl_profile" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_policy_object_feature_profile.test.id` + "\n"
	config += `	decrypt_categories = ["alcohol-and-tobacco"]` + "\n"
	config += `	no_decrypt_categories = ["abortion"]` + "\n"
	config += `	pass_through_categories = ["auctions"]` + "\n"
	config += `	reputation = true` + "\n"
	config += `	decrypt_threshold = "moderate-risk"` + "\n"
	config += `	threshold_categories = "moderate-risk"` + "\n"
	config += `	fail_decrypt = true` + "\n"
	config += `	url_allow_list_id = sdwan_policy_object_security_url_allow_list.test.id` + "\n"
	config += `	url_block_list_id = sdwan_policy_object_security_url_block_list.test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_policy_object_unified_tls_ssl_profile" "test" {
			id = sdwan_policy_object_unified_tls_ssl_profile.test.id
			feature_profile_id = sdwan_policy_object_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
