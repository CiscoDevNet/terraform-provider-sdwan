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

func TestAccDataSourceSdwanVPNMembershipPolicyDefinition(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_membership_policy_definition.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_membership_policy_definition.test", "description", "My description"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanVPNMembershipPolicyDefinitionPrerequisitesConfig + testAccDataSourceSdwanVPNMembershipPolicyDefinitionConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccDataSourceSdwanVPNMembershipPolicyDefinitionPrerequisitesConfig = `
resource "sdwan_site_list_policy_object" "sites1" {
  name = "TF_TEST"
  entries = [
    {
      site_id = "100-200"
    }
  ]
}

resource "sdwan_vpn_list_policy_object" "vpns1" {
  name = "TF_TEST"
  entries = [
    {
      vpn_id = "100-200"
    }
  ]
}

`

func testAccDataSourceSdwanVPNMembershipPolicyDefinitionConfig() string {
	config := `resource "sdwan_vpn_membership_policy_definition" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	sites = [{` + "\n"
	config += `	  site_list_id = sdwan_site_list_policy_object.sites1.id` + "\n"
	config += `	  vpn_list_ids = [sdwan_vpn_list_policy_object.vpns1.id]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_vpn_membership_policy_definition" "test" {
			id = sdwan_vpn_membership_policy_definition.test.id
		}
	`
	return config
}
