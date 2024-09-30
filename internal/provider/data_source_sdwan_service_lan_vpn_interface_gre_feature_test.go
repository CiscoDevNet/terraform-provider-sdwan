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
func TestAccDataSourceSdwanServiceLANVPNInterfaceGREProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_gre_feature.test", "interface_name", "gre1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_gre_feature.test", "interface_description", "gre1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_gre_feature.test", "ipv4_address", "70.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_gre_feature.test", "ipv4_subnet_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_gre_feature.test", "shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_gre_feature.test", "tunnel_source_ipv4_address", "78.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_gre_feature.test", "tunnel_destination_ipv4_address", "79.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_gre_feature.test", "ip_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_gre_feature.test", "tcp_mss", "1460"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_gre_feature.test", "clear_dont_fragment", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_gre_feature.test", "application_tunnel_type", "none"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanServiceLANVPNInterfaceGREPrerequisitesProfileParcelConfig + testAccDataSourceSdwanServiceLANVPNInterfaceGREProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanServiceLANVPNInterfaceGREPrerequisitesProfileParcelConfig = `
resource "sdwan_service_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_service_lan_vpn_feature" "test" {
  name = "TF_TEST_SLAN"
  feature_profile_id = sdwan_service_feature_profile.test.id
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanServiceLANVPNInterfaceGREProfileParcelConfig() string {
	config := `resource "sdwan_service_lan_vpn_interface_gre_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id` + "\n"
	config += `	interface_name = "gre1"` + "\n"
	config += `	interface_description = "gre1"` + "\n"
	config += `	ipv4_address = "70.1.1.1"` + "\n"
	config += `	ipv4_subnet_mask = "255.255.255.0"` + "\n"
	config += `	shutdown = true` + "\n"
	config += `	tunnel_source_ipv4_address = "78.1.1.1"` + "\n"
	config += `	tunnel_destination_ipv4_address = "79.1.1.1"` + "\n"
	config += `	ip_mtu = 1500` + "\n"
	config += `	tcp_mss = 1460` + "\n"
	config += `	clear_dont_fragment = false` + "\n"
	config += `	application_tunnel_type = "none"` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_service_lan_vpn_interface_gre_feature" "test" {
			id = sdwan_service_lan_vpn_interface_gre_feature.test.id
			feature_profile_id = sdwan_service_feature_profile.test.id
			service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
