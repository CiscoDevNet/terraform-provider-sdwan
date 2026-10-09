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
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "default_action", "pass"))
	if os.Getenv("SDWAN_2018") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "sequences.0.is_rule_set", "false"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "sequences.0.sequence_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "sequences.0.sequence_name", "security"))
	if os.Getenv("SDWAN_2018") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "sequences.0.sequence_comment", "comment1"))
	}
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "sequences.0.base_action", "drop"))
	if os.Getenv("SDWAN_2018") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test", "sequences.0.sequence_ip_type", "ipv4"))
	}
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
	if os.Getenv("SDWAN_2018") != "" {
		config += `	  sequence_comment = "comment1"` + "\n"
	}
	config += `	  base_action = "drop"` + "\n"
	if os.Getenv("SDWAN_2018") != "" {
		config += `	  sequence_ip_type = "ipv4"` + "\n"
	}
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

const testAccSdwanEmbeddedSecurityNGFWIPv6PrerequisitesProfileParcelConfig = `
resource "sdwan_embedded_security_feature_profile" "test_ipv6" {
  name        = "TF_TEST_NGFW_IPV6"
  description = "Terraform test"
}

resource "sdwan_policy_object_feature_profile" "test_ipv6" {
  name        = "TF_TEST_NGFW_IPV6_PO"
  description = "Terraform test"
}

resource "sdwan_policy_object_data_ipv6_prefix_list" "test_ipv6_source" {
  name               = "TF_TEST_NGFW_IPV6_SRC"
  description        = "Terraform test"
  feature_profile_id = sdwan_policy_object_feature_profile.test_ipv6.id
  entries = [
    {
      ipv6_address       = "2001:db8:1::"
      ipv6_prefix_length = 48
    }
  ]
}

resource "sdwan_policy_object_data_ipv6_prefix_list" "test_ipv6_destination" {
  name               = "TF_TEST_NGFW_IPV6_DST"
  description        = "Terraform test"
  feature_profile_id = sdwan_policy_object_feature_profile.test_ipv6.id
  entries = [
    {
      ipv6_address       = "2001:db8:2::"
      ipv6_prefix_length = 48
    }
  ]
}
`

// Each match entry carries exactly one criterion (`maxProperties: 1` in the parcel schema).
func testAccSdwanEmbeddedSecurityNGFWIPv6ProfileParcelConfig() string {
	config := `resource "sdwan_embedded_security_ngfw_policy" "test_ipv6" {` + "\n"
	config += `	name = "TF_TEST_NGFW_IPV6"` + "\n"
	config += `	description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_embedded_security_feature_profile.test_ipv6.id` + "\n"
	config += `	default_action = "pass"` + "\n"
	config += `	sequences = [{` + "\n"
	config += `	  sequence_id = "1"` + "\n"
	config += `	  sequence_name = "security-ipv6"` + "\n"
	config += `	  base_action = "drop"` + "\n"
	config += `	  sequence_ip_type = "ipv6"` + "\n"
	config += `	  sequence_type = "ngfirewall"` + "\n"
	config += `	  disable_sequence = false` + "\n"
	config += `	  match_entries = [{` + "\n"
	config += `		source_ipv6_data_prefixes = ["2001:db8:3::/48"]` + "\n"
	config += `	  }, {` + "\n"
	config += `		destination_ipv6_data_prefixes = ["2001:db8:4::/48"]` + "\n"
	config += `	  }, {` + "\n"
	config += `		source_data_ipv6_prefix_list_ids = [sdwan_policy_object_data_ipv6_prefix_list.test_ipv6_source.id]` + "\n"
	config += `	  }, {` + "\n"
	config += `		destination_data_ipv6_prefix_list_ids = [sdwan_policy_object_data_ipv6_prefix_list.test_ipv6_destination.id]` + "\n"
	config += `	  }]` + "\n"
	config += `	  actions = [{` + "\n"
	config += `		type = "log"` + "\n"
	config += `		parameter = "true"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

func TestAccSdwanEmbeddedSecurityNGFWProfileParcelIPv6Sequence(t *testing.T) {
	if os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2018")
	}

	config := testAccSdwanEmbeddedSecurityNGFWIPv6PrerequisitesProfileParcelConfig + testAccSdwanEmbeddedSecurityNGFWIPv6ProfileParcelConfig()

	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test_ipv6", "sequences.0.sequence_ip_type", "ipv6"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test_ipv6", "sequences.0.match_entries.#", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test_ipv6", "sequences.0.match_entries.0.source_ipv6_data_prefixes.0", "2001:db8:3::/48"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_embedded_security_ngfw_policy.test_ipv6", "sequences.0.match_entries.1.destination_ipv6_data_prefixes.0", "2001:db8:4::/48"))
	checks = append(checks, resource.TestCheckResourceAttrPair("sdwan_embedded_security_ngfw_policy.test_ipv6", "sequences.0.match_entries.2.source_data_ipv6_prefix_list_ids.0", "sdwan_policy_object_data_ipv6_prefix_list.test_ipv6_source", "id"))
	checks = append(checks, resource.TestCheckResourceAttrPair("sdwan_embedded_security_ngfw_policy.test_ipv6", "sequences.0.match_entries.3.destination_data_ipv6_prefix_list_ids.0", "sdwan_policy_object_data_ipv6_prefix_list.test_ipv6_destination", "id"))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			// Drift guard: if a match criterion were silently dropped on write, the
			// refreshed state would differ from the config and the plan would not be empty.
			{
				Config:             config,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}
