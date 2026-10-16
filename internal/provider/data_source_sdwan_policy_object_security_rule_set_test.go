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
func TestAccDataSourceSdwanPolicyObjectSecurityRuleSetProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2018")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_security_rule_set.test", "sequences.0.name", "rule1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_security_rule_set.test", "sequences.0.sequence_ip_type", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_security_rule_set.test", "sequences.0.sequence_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_security_rule_set.test", "sequences.0.action", "drop"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanPolicyObjectSecurityRuleSetPrerequisitesProfileParcelConfig + testAccDataSourceSdwanPolicyObjectSecurityRuleSetProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceSdwanPolicyObjectSecurityRuleSetPrerequisitesProfileParcelConfig + testAccDataSourceSdwanPolicyObjectSecurityRuleSetProfileParcelByNameConfig(),
				Check: resource.ComposeTestCheckFunc(
					append(checks,
						resource.TestCheckResourceAttr("data.sdwan_policy_object_security_rule_set.test", "name", "TF_TEST"),
						resource.TestCheckResourceAttrSet("data.sdwan_policy_object_security_rule_set.test", "id"),
					)...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanPolicyObjectSecurityRuleSetPrerequisitesProfileParcelConfig = `
resource "sdwan_policy_object_feature_profile" "test" {
  name = "POLICY_OBJECT_FP_1"
  description = "My policy object feature profile 1"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanPolicyObjectSecurityRuleSetProfileParcelConfig() string {
	config := `resource "sdwan_policy_object_security_rule_set" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_policy_object_feature_profile.test.id` + "\n"
	config += `	sequences = [{` + "\n"
	config += `	  name = "rule1"` + "\n"
	config += `	  sequence_ip_type = "ipv4"` + "\n"
	config += `	  sequence_id = "1"` + "\n"
	config += `	  action = "drop"` + "\n"
	config += `	  source_ipv4_prefixes = ["10.1.1.0/24"]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_policy_object_security_rule_set" "test" {
			id = sdwan_policy_object_security_rule_set.test.id
			feature_profile_id = sdwan_policy_object_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceByNameConfig
func testAccDataSourceSdwanPolicyObjectSecurityRuleSetProfileParcelByNameConfig() string {
	config := `resource "sdwan_policy_object_security_rule_set" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_policy_object_feature_profile.test.id` + "\n"
	config += `	sequences = [{` + "\n"
	config += `	  name = "rule1"` + "\n"
	config += `	  sequence_ip_type = "ipv4"` + "\n"
	config += `	  sequence_id = "1"` + "\n"
	config += `	  action = "drop"` + "\n"
	config += `	  source_ipv4_prefixes = ["10.1.1.0/24"]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_policy_object_security_rule_set" "test" {
			name = "TF_TEST"
			feature_profile_id = sdwan_policy_object_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceByNameConfig
