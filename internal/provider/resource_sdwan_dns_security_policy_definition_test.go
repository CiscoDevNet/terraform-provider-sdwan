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
func TestAccSdwanDNSSecurityPolicyDefinition(t *testing.T) {
	if os.Getenv("SDWAN_209") == "" {
		t.Skip("skipping test, set environment variable SDWAN_209")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_dns_security_policy_definition.test", "name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_dns_security_policy_definition.test", "description", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_dns_security_policy_definition.test", "local_domain_bypass_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_dns_security_policy_definition.test", "match_all_vpn", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_dns_security_policy_definition.test", "dnscrypt", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_dns_security_policy_definition.test", "umbrella_dns_default", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_dns_security_policy_definition.test", "cisco_sig_credentials_feature_template_id", "d77f1f7f-b755-47f2-bc9d-8af572f3bf4b"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanDNSSecurityPolicyDefinitionPrerequisitesConfig + testAccSdwanDNSSecurityPolicyDefinitionConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanDNSSecurityPolicyDefinitionPrerequisitesConfig = `
resource "sdwan_domain_list_policy_object" "test" {
  name = "TEST_TF"
  entries = [
    {
      domain = ".*.cisco.com"
    }
  ]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanDNSSecurityPolicyDefinitionConfig_all() string {
	config := `resource "sdwan_dns_security_policy_definition" "test" {` + "\n"
	config += `	name = "Example"` + "\n"
	config += `	description = "Example"` + "\n"
	config += `	domain_list_id = sdwan_domain_list_policy_object.test.id` + "\n"
	config += `	local_domain_bypass_enabled = false` + "\n"
	config += `	match_all_vpn = true` + "\n"
	config += `	dnscrypt = true` + "\n"
	config += `	umbrella_dns_default = true` + "\n"
	config += `	cisco_sig_credentials_feature_template_id = "d77f1f7f-b755-47f2-bc9d-8af572f3bf4b"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
