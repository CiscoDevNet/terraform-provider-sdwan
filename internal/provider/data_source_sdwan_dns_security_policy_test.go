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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceSdwanDNSSecurityProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_dns_security_policy.test", "match_all_vpn", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_dns_security_policy.test", "umbrella_default", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_dns_security_policy.test", "dns_server_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_dns_security_policy.test", "local_domain_bypass_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_dns_security_policy.test", "dns_crypt", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_dns_security_policy.test", "child_org_id", "12334"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanDNSSecurityPrerequisitesProfileParcelConfig + testAccDataSourceSdwanDNSSecurityProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanDNSSecurityPrerequisitesProfileParcelConfig = `
resource "sdwan_dns_security_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_policy_object_feature_profile" "test" {
  name = "TF_TEST_POLICY_OBJECT"
  description = "My policy object feature profile 1"
}

resource "sdwan_policy_object_security_local_domain_list" "test" {
  name               = "TF_TEST_LOCAL_DOMAIN"
  description        = "My Example"
  feature_profile_id = sdwan_policy_object_feature_profile.test.id
  entries = [
    {
      local_domain = "hello.com"
    }
  ]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanDNSSecurityProfileParcelConfig() string {
	config := `resource "sdwan_dns_security_policy" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_dns_security_feature_profile.test.id` + "\n"
	config += `	local_domain_bypass_list_id = sdwan_policy_object_security_local_domain_list.test.id` + "\n"
	config += `	match_all_vpn = true` + "\n"
	config += `	umbrella_default = false` + "\n"
	config += `	dns_server_ip = "1.2.3.4"` + "\n"
	config += `	local_domain_bypass_enabled = true` + "\n"
	config += `	dns_crypt = false` + "\n"
	config += `	child_org_id = "12334"` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_dns_security_policy" "test" {
			id = sdwan_dns_security_policy.test.id
			feature_profile_id = sdwan_dns_security_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
