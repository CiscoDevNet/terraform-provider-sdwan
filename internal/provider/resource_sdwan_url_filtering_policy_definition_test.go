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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanURLFilteringPolicyDefinition(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_url_filtering_policy_definition.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_url_filtering_policy_definition.test", "description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_url_filtering_policy_definition.test", "mode", "security"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_url_filtering_policy_definition.test", "web_categories_action", "allow"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_url_filtering_policy_definition.test", "web_reputation", "moderate-risk"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_url_filtering_policy_definition.test", "block_page_action", "text"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_url_filtering_policy_definition.test", "block_page_contents", "Access to the requested page has been denied. Please contact your Network Administrator"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanURLFilteringPolicyDefinitionConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanURLFilteringPolicyDefinitionConfig_all() string {
	config := `resource "sdwan_url_filtering_policy_definition" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "My description"` + "\n"
	config += `	mode = "security"` + "\n"
	config += `	alerts = ["blacklist"]` + "\n"
	config += `	web_categories = ["alcohol-and-tobacco"]` + "\n"
	config += `	web_categories_action = "allow"` + "\n"
	config += `	web_reputation = "moderate-risk"` + "\n"
	config += `	target_vpns = ["1"]` + "\n"
	config += `	block_page_action = "text"` + "\n"
	config += `	block_page_contents = "Access to the requested page has been denied. Please contact your Network Administrator"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
