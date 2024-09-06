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
func TestAccSdwanSystemGlobalProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "http_server", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "https_server", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "ftp_passive", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "domain_lookup", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "arp_proxy", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "rsh_rcp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "line_vty", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "cdp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "lldp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "source_interface", "GigabitEthernet0/0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "tcp_keepalives_in", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "tcp_keepalives_out", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "tcp_small_servers", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "udp_small_servers", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "console_logging", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "ip_source_routing", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "vty_line_logging", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "snmp_ifindex_persist", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "ignore_bootp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "nat64_udp_timeout", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "nat64_tcp_timeout", "3600"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "http_authentication", "aaa"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_global_feature.test", "ssh_version", "2"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanSystemGlobalPrerequisitesProfileParcelConfig + testAccSdwanSystemGlobalProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanSystemGlobalPrerequisitesProfileParcelConfig + testAccSdwanSystemGlobalProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanSystemGlobalPrerequisitesProfileParcelConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanSystemGlobalProfileParcelConfig_minimum() string {
	config := `resource "sdwan_system_global_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanSystemGlobalProfileParcelConfig_all() string {
	config := `resource "sdwan_system_global_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `	http_server = false` + "\n"
	config += `	https_server = false` + "\n"
	config += `	ftp_passive = false` + "\n"
	config += `	domain_lookup = false` + "\n"
	config += `	arp_proxy = false` + "\n"
	config += `	rsh_rcp = false` + "\n"
	config += `	line_vty = false` + "\n"
	config += `	cdp = true` + "\n"
	config += `	lldp = true` + "\n"
	config += `	source_interface = "GigabitEthernet0/0/1"` + "\n"
	config += `	tcp_keepalives_in = true` + "\n"
	config += `	tcp_keepalives_out = true` + "\n"
	config += `	tcp_small_servers = false` + "\n"
	config += `	udp_small_servers = false` + "\n"
	config += `	console_logging = true` + "\n"
	config += `	ip_source_routing = false` + "\n"
	config += `	vty_line_logging = false` + "\n"
	config += `	snmp_ifindex_persist = true` + "\n"
	config += `	ignore_bootp = true` + "\n"
	config += `	nat64_udp_timeout = 300` + "\n"
	config += `	nat64_tcp_timeout = 3600` + "\n"
	config += `	http_authentication = "aaa"` + "\n"
	config += `	ssh_version = "2"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
