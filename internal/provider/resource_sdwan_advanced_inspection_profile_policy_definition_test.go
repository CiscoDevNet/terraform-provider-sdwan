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

func TestAccSdwanAdvancedInspectionProfilePolicyDefinition(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanAdvancedInspectionProfilePolicyDefinitionConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_advanced_inspection_profile_policy_definition.test", "name", "Example"),
					resource.TestCheckResourceAttr("sdwan_advanced_inspection_profile_policy_definition.test", "description", "My description"),
					resource.TestCheckResourceAttr("sdwan_advanced_inspection_profile_policy_definition.test", "tls_action", "decrypt"),
				),
			},
		},
	})
}

const testAccSdwanAdvancedInspectionProfilePolicyDefinitionConfig = `
resource "sdwan_url_filtering_policy_definition" "test" {
  name                  = "TF_TEST"
  description           = "Terraform test"
  mode                  = "security"
  alerts                = ["blacklist"]
  web_categories        = ["alcohol-and-tobacco"]
  web_categories_action = "allow"
  web_reputation        = "moderate-risk"
  target_vpns           = ["1"]
  block_page_action     = "text"
  block_page_contents   = "Access to the requested page has been denied. Please contact your Network Administrator"
}


resource "sdwan_advanced_inspection_profile_policy_definition" "test" {
	name = "Example"
	description = "My description"
	tls_action = "decrypt"
	url_filtering_id = sdwan_url_filtering_policy_definition.test.id
}
`
