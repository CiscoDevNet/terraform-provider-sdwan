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
	if os.Getenv("SDWAN_2018") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "sequences.0.is_rule_set", "false"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "sequences.0.sequence_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "sequences.0.sequence_name", "security"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "sequences.0.base_action", "drop"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "sequences.0.sequence_type", "ngfirewall"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "sequences.0.disable_sequence", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "sequences.0.actions.0.type", "log"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "sequences.0.actions.0.parameter", "true"))
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
	config += `	sequences = [{` + "\n"
	if os.Getenv("SDWAN_2018") != "" {
		config += `	  is_rule_set = false` + "\n"
	}
	config += `	  sequence_id = "1"` + "\n"
	config += `	  sequence_name = "security"` + "\n"
	config += `	  base_action = "drop"` + "\n"
	config += `	  sequence_type = "ngfirewall"` + "\n"
	config += `	  disable_sequence = false` + "\n"
	config += `	}]` + "\n"
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
	config += `	sequences = [{` + "\n"
	if os.Getenv("SDWAN_2018") != "" {
		config += `	  is_rule_set = false` + "\n"
	}
	config += `	  sequence_id = "1"` + "\n"
	config += `	  sequence_name = "security"` + "\n"
	config += `	  base_action = "drop"` + "\n"
	config += `	  sequence_type = "ngfirewall"` + "\n"
	config += `	  disable_sequence = false` + "\n"
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

func TestAccSdwanEmbeddedSecurityNGFWProfileParcelRuleSet(t *testing.T) {
	if os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2018")
	}

	config := testAccSdwanEmbeddedSecurityNGFWRuleSetConfig
	checks := []resource.TestCheckFunc{
		resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test_ruleset", "sequences.0.is_rule_set", "true"),
		resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test_ruleset", "sequences.0.match_entries.0.rule_set_list_ids.#", "1"),
		resource.TestCheckResourceAttrPair("sdwan_embedded_security_ngfw_policy.test_ruleset", "sequences.0.match_entries.0.rule_set_list_ids.0", "sdwan_policy_object_security_rule_set.test", "id"),
		resource.TestCheckResourceAttr("data.sdwan_embedded_security_ngfw_policy.test_ruleset", "sequences.0.is_rule_set", "true"),
		resource.TestCheckResourceAttrPair("data.sdwan_embedded_security_ngfw_policy.test_ruleset", "sequences.0.match_entries.0.rule_set_list_ids.0", "sdwan_policy_object_security_rule_set.test", "id"),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config:             config,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

const testAccSdwanEmbeddedSecurityNGFWRuleSetConfig = `
resource "sdwan_embedded_security_feature_profile" "test_ruleset" {
	name        = "TF_TEST_NGFW_RULESET_FP"
	description = "Terraform test"
}

resource "sdwan_policy_object_feature_profile" "test_ruleset" {
	name        = "TF_TEST_NGFW_RULESET_PO"
	description = "Terraform test"
}

resource "sdwan_policy_object_security_rule_set" "test" {
	name               = "TF_TEST_NGFW_RULESET"
	description        = "Terraform test"
	feature_profile_id = sdwan_policy_object_feature_profile.test_ruleset.id
	sequences = [{
		name                = "rule1"
		sequence_ip_type    = "ipv4"
		sequence_id         = "1"
		action              = "drop"
		source_ipv4_prefixes = ["10.1.1.0/24"]
	}]
}

resource "sdwan_embedded_security_ngfw_policy" "test_ruleset" {
	name               = "TF_TEST_NGFW_RULESET_POLICY"
	description        = "Terraform test"
	feature_profile_id = sdwan_embedded_security_feature_profile.test_ruleset.id
	default_action     = "pass"
	sequences = [{
		is_rule_set      = true
		sequence_id      = "1"
		sequence_name    = "ruleset"
		base_action      = "drop"
		sequence_type    = "ngfirewall"
		disable_sequence = false
		match_entries = [{
			rule_set_list_ids = [sdwan_policy_object_security_rule_set.test.id]
		}]
		actions = [{
			type      = "log"
			parameter = "true"
		}]
	}]
}

data "sdwan_embedded_security_ngfw_policy" "test_ruleset" {
	id                 = sdwan_embedded_security_ngfw_policy.test_ruleset.id
	feature_profile_id = sdwan_embedded_security_feature_profile.test_ruleset.id
}
`
