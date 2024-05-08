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
func TestAccDataSourceSdwanSecurityPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_security_policy.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_security_policy.test", "description", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_security_policy.test", "mode", "security"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_security_policy.test", "use_case", "custom"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_security_policy.test", "definitions.0.type", "urlFiltering"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_security_policy.test", "failure_mode", "close"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_security_policy.test", "logging.0.external_syslog_server_ip", "10.0.0.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_security_policy.test", "logging.0.external_syslog_server_vpn", "123"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanSecurityPolicyPrerequisitesConfig + testAccDataSourceSdwanSecurityPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanSecurityPolicyPrerequisitesConfig = `
resource "sdwan_url_filtering_policy_definition" "test" {
  name                  = "TEST_TF"
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

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanSecurityPolicyConfig() string {
	config := ""
	config += `resource "sdwan_security_policy" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "Example"` + "\n"
	config += `	mode = "security"` + "\n"
	config += `	use_case = "custom"` + "\n"
	config += `	definitions = [{` + "\n"
	config += `	  id = sdwan_url_filtering_policy_definition.test.id` + "\n"
	config += `	  type = "urlFiltering"` + "\n"
	config += `	}]` + "\n"
	config += `	failure_mode = "close"` + "\n"
	config += `	logging = [{` + "\n"
	config += `	  external_syslog_server_ip = "10.0.0.1"` + "\n"
	config += `	  external_syslog_server_vpn = "123"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_security_policy" "test" {
			id = sdwan_security_policy.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
