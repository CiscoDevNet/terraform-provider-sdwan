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
func TestAccSdwanEmbeddedSecurityNGFWProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "default_action", "pass"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "rules.0.rule_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "rules.0.rule_name", "Rule1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "rules.0.base_action", "drop"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "rules.0.rule_type", "ngfirewall"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "rules.0.disable_rule", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "rules.0.actions.0.type", "log"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "rules.0.actions.0.parameter", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanEmbeddedSecurityNGFWPrerequisitesProfileParcelConfig + testAccSdwanEmbeddedSecurityNGFWProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanEmbeddedSecurityNGFWPrerequisitesProfileParcelConfig + testAccSdwanEmbeddedSecurityNGFWProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanEmbeddedSecurityNGFWPrerequisitesProfileParcelConfig = `
resource "sdwan_embedded_security_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanEmbeddedSecurityNGFWProfileParcelConfig_minimum() string {
	config := `resource "sdwan_embedded_security_ngfw_policy" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_embedded_security_feature_profile.test.id` + "\n"
	config += `	default_action = "pass"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanEmbeddedSecurityNGFWProfileParcelConfig_all() string {
	config := `resource "sdwan_embedded_security_ngfw_policy" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_embedded_security_feature_profile.test.id` + "\n"
	config += `	default_action = "pass"` + "\n"
	config += `	rules = [{` + "\n"
	config += `	  rule_id = "1"` + "\n"
	config += `	  rule_name = "Rule1"` + "\n"
	config += `	  base_action = "drop"` + "\n"
	config += `	  rule_type = "ngfirewall"` + "\n"
	config += `	  disable_rule = false` + "\n"
	config += `	  match_entries = [{` + "\n"
	config += `		source_ports = ["123"]` + "\n"
	config += `	}]` + "\n"
	config += `	  actions = [{` + "\n"
	config += `		type = "log"` + "\n"
	config += `		parameter = "true"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
