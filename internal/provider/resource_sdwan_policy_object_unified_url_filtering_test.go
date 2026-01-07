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
func TestAccSdwanPolicyObjectUnifiedURLFilteringProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015_IN_PROGRESS") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015_IN_PROGRESS")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_url_filtering.test", "web_categories_action", "block"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_url_filtering.test", "web_reputation", "suspicious"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_url_filtering.test", "block_page_action", "text"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_url_filtering.test", "block_page_contents", "Access to the requested page has been denied. Please contact your Network Administrator"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_unified_url_filtering.test", "enable_alerts", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			{
				Config: testAccSdwanPolicyObjectUnifiedURLFilteringPrerequisitesProfileParcelConfig + testAccSdwanPolicyObjectUnifiedURLFilteringProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanPolicyObjectUnifiedURLFilteringPrerequisitesProfileParcelConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanPolicyObjectUnifiedURLFilteringProfileParcelConfig_all() string {
	config := `resource "sdwan_policy_object_unified_url_filtering" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_policy_object_feature_profile.test.id` + "\n"
	config += `	web_categories_action = "block"` + "\n"
	config += `	web_categories = ["confirmed-spam-sources"]` + "\n"
	config += `	web_reputation = "suspicious"` + "\n"
	config += `	url_allow_list_id = sdwan_policy_object_security_url_allow_list.test.id` + "\n"
	config += `	url_block_list_id = sdwan_policy_object_security_url_block_list.test.id` + "\n"
	config += `	block_page_action = "text"` + "\n"
	config += `	block_page_contents = "Access to the requested page has been denied. Please contact your Network Administrator"` + "\n"
	config += `	enable_alerts = true` + "\n"
	config += `	alerts = ["blacklist"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
