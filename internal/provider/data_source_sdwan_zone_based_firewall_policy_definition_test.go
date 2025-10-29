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
func TestAccDataSourceSdwanZoneBasedFirewallPolicyDefinition(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_zone_based_firewall_policy_definition.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_zone_based_firewall_policy_definition.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_zone_based_firewall_policy_definition.test", "mode", "security"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_zone_based_firewall_policy_definition.test", "apply_zone_pairs.0.source_zone", "self"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_zone_based_firewall_policy_definition.test", "default_action", "pass"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_zone_based_firewall_policy_definition.test", "rules.0.rule_order", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_zone_based_firewall_policy_definition.test", "rules.0.rule_name", "RULE_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_zone_based_firewall_policy_definition.test", "rules.0.base_action", "inspect"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_zone_based_firewall_policy_definition.test", "rules.0.ip_type", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_zone_based_firewall_policy_definition.test", "rules.0.match_entries.0.type", "sourceGeoLocationList"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_zone_based_firewall_policy_definition.test", "rules.0.action_entries.0.type", "log"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanZoneBasedFirewallPolicyDefinitionPrerequisitesConfig + testAccDataSourceSdwanZoneBasedFirewallPolicyDefinitionConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanZoneBasedFirewallPolicyDefinitionPrerequisitesConfig = `
resource "sdwan_zone_list_policy_object" "test" {
  name = "TF_TEST"
  entries = [
    {
      vpn = "1"
    }
  ]
}

resource "sdwan_geo_location_list_policy_object" "test" {
  name = "Example"
  entries = [
    {
      country   = "USA"
      continent = "AS"
    }
  ]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanZoneBasedFirewallPolicyDefinitionConfig() string {
	config := ""
	config += `resource "sdwan_zone_based_firewall_policy_definition" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	mode = "security"` + "\n"
	config += `	apply_zone_pairs = [{` + "\n"
	config += `	  source_zone = "self"` + "\n"
	config += `	  destination_zone = sdwan_zone_list_policy_object.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	default_action = "pass"` + "\n"
	config += `	rules = [{` + "\n"
	config += `	  rule_order = 1` + "\n"
	config += `	  rule_name = "RULE_1"` + "\n"
	config += `	  base_action = "inspect"` + "\n"
	config += `	  ip_type = "ipv4"` + "\n"
	config += `	  match_entries = [{` + "\n"
	config += `		type = "sourceGeoLocationList"` + "\n"
	config += `		policy_id = sdwan_geo_location_list_policy_object.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	  action_entries = [{` + "\n"
	config += `		type = "log"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_zone_based_firewall_policy_definition" "test" {
			id = sdwan_zone_based_firewall_policy_definition.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
