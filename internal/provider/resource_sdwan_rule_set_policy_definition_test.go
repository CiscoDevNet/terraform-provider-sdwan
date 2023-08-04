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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanRuleSetPolicyDefinition(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanRuleSetPolicyDefinitionConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_rule_set_policy_definition.test", "name", "Example"),
					resource.TestCheckResourceAttr("sdwan_rule_set_policy_definition.test", "description", "My description"),
					resource.TestCheckResourceAttr("sdwan_rule_set_policy_definition.test", "rules.0.name", "Rule1"),
					resource.TestCheckResourceAttr("sdwan_rule_set_policy_definition.test", "rules.0.order", "1"),
					resource.TestCheckResourceAttr("sdwan_rule_set_policy_definition.test", "rules.0.source_ipv4_prefix", "10.1.1.0/24"),
					resource.TestCheckResourceAttr("sdwan_rule_set_policy_definition.test", "rules.0.source_fqdn", "cisco.com"),
					resource.TestCheckResourceAttr("sdwan_rule_set_policy_definition.test", "rules.0.source_port", "80-90"),
					resource.TestCheckResourceAttr("sdwan_rule_set_policy_definition.test", "rules.0.source_geo_location", "AF"),
					resource.TestCheckResourceAttr("sdwan_rule_set_policy_definition.test", "rules.0.destination_ipv4_prefix", "10.1.1.0/24"),
					resource.TestCheckResourceAttr("sdwan_rule_set_policy_definition.test", "rules.0.destination_fqdn", "cisco.com"),
					resource.TestCheckResourceAttr("sdwan_rule_set_policy_definition.test", "rules.0.destination_port", "80-90"),
					resource.TestCheckResourceAttr("sdwan_rule_set_policy_definition.test", "rules.0.destination_geo_location", "AF"),
					resource.TestCheckResourceAttr("sdwan_rule_set_policy_definition.test", "rules.0.protocol", "cifs"),
				),
			},
		},
	})
}

const testAccSdwanRuleSetPolicyDefinitionConfig = `


resource "sdwan_rule_set_policy_definition" "test" {
	name = "Example"
	description = "My description"
	rules = [{
		name = "Rule1"
		order = 1
		source_ipv4_prefix = "10.1.1.0/24"
		source_fqdn = "cisco.com"
		source_port = "80-90"
		source_geo_location = "AF"
		destination_ipv4_prefix = "10.1.1.0/24"
		destination_fqdn = "cisco.com"
		destination_port = "80-90"
		destination_geo_location = "AF"
		protocol = "cifs"
	}]
}
`