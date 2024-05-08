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
func TestAccDataSourceSdwanRuleSetPolicyDefinition(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_rule_set_policy_definition.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_rule_set_policy_definition.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_rule_set_policy_definition.test", "rules.0.name", "Rule1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_rule_set_policy_definition.test", "rules.0.order", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_rule_set_policy_definition.test", "rules.0.source_ipv4_prefix", "10.1.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_rule_set_policy_definition.test", "rules.0.source_fqdn", "cisco.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_rule_set_policy_definition.test", "rules.0.source_port", "80-90"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_rule_set_policy_definition.test", "rules.0.source_geo_location", "AF"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_rule_set_policy_definition.test", "rules.0.destination_ipv4_prefix", "10.1.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_rule_set_policy_definition.test", "rules.0.destination_fqdn", "cisco.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_rule_set_policy_definition.test", "rules.0.destination_port", "80-90"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_rule_set_policy_definition.test", "rules.0.destination_geo_location", "AF"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_rule_set_policy_definition.test", "rules.0.protocol", "cifs"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanRuleSetPolicyDefinitionConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanRuleSetPolicyDefinitionConfig() string {
	config := ""
	config += `resource "sdwan_rule_set_policy_definition" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	rules = [{` + "\n"
	config += `	  name = "Rule1"` + "\n"
	config += `	  order = 1` + "\n"
	config += `	  source_ipv4_prefix = "10.1.1.0/24"` + "\n"
	config += `	  source_fqdn = "cisco.com"` + "\n"
	config += `	  source_port = "80-90"` + "\n"
	config += `	  source_geo_location = "AF"` + "\n"
	config += `	  destination_ipv4_prefix = "10.1.1.0/24"` + "\n"
	config += `	  destination_fqdn = "cisco.com"` + "\n"
	config += `	  destination_port = "80-90"` + "\n"
	config += `	  destination_geo_location = "AF"` + "\n"
	config += `	  protocol = "cifs"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_rule_set_policy_definition" "test" {
			id = sdwan_rule_set_policy_definition.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
