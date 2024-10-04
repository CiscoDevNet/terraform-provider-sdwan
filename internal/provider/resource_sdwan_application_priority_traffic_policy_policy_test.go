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
func TestAccSdwanApplicationPriorityTrafficPolicyProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" && os.Getenv("TF_VAR_policy_object_feature_template_id") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012 or TF_VAR_policy_object_feature_template_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_application_priority_traffic_policy_policy.test", "default_action", "accept"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_application_priority_traffic_policy_policy.test", "direction", "all"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_application_priority_traffic_policy_policy.test", "sequences.0.sequence_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_application_priority_traffic_policy_policy.test", "sequences.0.sequence_name", "traffic"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_application_priority_traffic_policy_policy.test", "sequences.0.base_action", "accept"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_application_priority_traffic_policy_policy.test", "sequences.0.protocol", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_application_priority_traffic_policy_policy.test", "sequences.0.matches.0.dscp", "1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			{
				Config: testAccSdwanApplicationPriorityTrafficPolicyPrerequisitesProfileParcelConfig + testAccSdwanApplicationPriorityTrafficPolicyProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanApplicationPriorityTrafficPolicyPrerequisitesProfileParcelConfig = `
resource "sdwan_policy_object_feature_profile" "test" {
  name = "TF_TEST_POLICY_OBJECT"
  description = "My policy object feature profile 1"
}

resource "sdwan_application_priority_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_policy_object_policer" "test" {
  name               = "TF_TEST_POLICER"
  description        = "My Example"
  feature_profile_id = sdwan_policy_object_feature_profile.test.id
  entries = [
    {
      burst_bytes   = 56500
      exceed_action = "remark"
      rate_bps      = 60000
    }
  ]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanApplicationPriorityTrafficPolicyProfileParcelConfig_all() string {
	config := `resource "sdwan_application_priority_traffic_policy_policy" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_application_priority_feature_profile.test.id` + "\n"
	config += `	default_action = "accept"` + "\n"
	config += `	vpns = ["Local_Internet_for_Guests"]` + "\n"
	config += `	direction = "all"` + "\n"
	config += `	sequences = [{` + "\n"
	config += `	  sequence_id = 1` + "\n"
	config += `	  sequence_name = "traffic"` + "\n"
	config += `	  base_action = "accept"` + "\n"
	config += `	  protocol = "ipv4"` + "\n"
	config += `	  matches = [{` + "\n"
	config += `		dscp = 1` + "\n"
	config += `	}]` + "\n"
	config += `	  actions = [{` + "\n"
	config += `      set_parameters = [{` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
