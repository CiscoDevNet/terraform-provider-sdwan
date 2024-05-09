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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceSdwanLocalizedPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "flow_visibility_ipv4", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "flow_visibility_ipv6", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "application_visibility_ipv4", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "application_visibility_ipv6", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "cloud_qos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "cloud_qos_service_side", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "implicit_acl_logging", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "log_frequency", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "ipv4_visibility_cache_entries", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "ipv6_visibility_cache_entries", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "definitions.0.type", "acl"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanLocalizedPolicyPrerequisitesConfig + testAccDataSourceSdwanLocalizedPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanLocalizedPolicyPrerequisitesConfig = `
resource "sdwan_ipv4_acl_policy_definition" "test" {
  name           = "TF_TEST"
  description    = "Terraform test"
  default_action = "drop"
  sequences = [
    {
      id          = 10
      ip_type     = "ipv4"
      name        = "Sequence 10"
      base_action = "accept"
      match_entries = [
        {
          type = "dscp"
          dscp = 16
        }
      ]
      action_entries = [
        {
          type = "set"
          set_parameters = [
            {
              type = "dscp"
              dscp = 16
            }
          ]
        }
      ]
    }
  ]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanLocalizedPolicyConfig() string {
	config := ""
	config += `resource "sdwan_localized_policy" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	flow_visibility_ipv4 = true` + "\n"
	config += `	flow_visibility_ipv6 = true` + "\n"
	config += `	application_visibility_ipv4 = true` + "\n"
	config += `	application_visibility_ipv6 = true` + "\n"
	config += `	cloud_qos = true` + "\n"
	config += `	cloud_qos_service_side = true` + "\n"
	config += `	implicit_acl_logging = true` + "\n"
	config += `	log_frequency = 1000` + "\n"
	config += `	ipv4_visibility_cache_entries = 1000` + "\n"
	config += `	ipv6_visibility_cache_entries = 1000` + "\n"
	config += `	definitions = [{` + "\n"
	config += `	  id = sdwan_ipv4_acl_policy_definition.test.id` + "\n"
	config += `	  type = "acl"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_localized_policy" "test" {
			id = sdwan_localized_policy.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
