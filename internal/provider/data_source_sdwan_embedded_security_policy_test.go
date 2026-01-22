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
func TestAccDataSourceSdwanEmbeddedSecurityProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_embedded_security_policy.test", "tcp_syn_flood_limit", "432"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_embedded_security_policy.test", "max_incomplete_tcp_limit", "12345"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_embedded_security_policy.test", "max_incomplete_udp_limit", "12345"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_embedded_security_policy.test", "max_incomplete_icmp_limit", "12345"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_embedded_security_policy.test", "audit_trail", "on"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_embedded_security_policy.test", "unified_logging", "off"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_embedded_security_policy.test", "session_reclassify_allow", "off"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_embedded_security_policy.test", "imcp_unreachable_allow", "off"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_embedded_security_policy.test", "failure_mode", "close"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_embedded_security_policy.test", "security_logging", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_embedded_security_policy.test", "nat", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_embedded_security_policy.test", "download_url_database_on_device", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_embedded_security_policy.test", "resource_profile", "low"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanEmbeddedSecurityPrerequisitesProfileParcelConfig + testAccDataSourceSdwanEmbeddedSecurityProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanEmbeddedSecurityPrerequisitesProfileParcelConfig = `
resource "sdwan_embedded_security_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanEmbeddedSecurityProfileParcelConfig() string {
	config := `resource "sdwan_embedded_security_policy" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_embedded_security_feature_profile.test.id` + "\n"
	config += `	assembly = [{` + "\n"
	config += `	  entries = [{` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	tcp_syn_flood_limit = "432"` + "\n"
	config += `	max_incomplete_tcp_limit = "12345"` + "\n"
	config += `	max_incomplete_udp_limit = "12345"` + "\n"
	config += `	max_incomplete_icmp_limit = "12345"` + "\n"
	config += `	audit_trail = "on"` + "\n"
	config += `	unified_logging = "off"` + "\n"
	config += `	session_reclassify_allow = "off"` + "\n"
	config += `	imcp_unreachable_allow = "off"` + "\n"
	config += `	failure_mode = "close"` + "\n"
	config += `	security_logging = "true"` + "\n"
	config += `	nat = true` + "\n"
	config += `	download_url_database_on_device = false` + "\n"
	config += `	resource_profile = "low"` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_embedded_security_policy" "test" {
			id = sdwan_embedded_security_policy.test.id
			feature_profile_id = sdwan_embedded_security_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
