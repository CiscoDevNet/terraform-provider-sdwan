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
func TestAccSdwanServiceLANVPNInterfaceGREProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "interface_name", "gre1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "interface_description", "gre1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "ipv4_address", "70.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "ipv4_subnet_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "tunnel_protection", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "tunnel_mode", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "tunnel_source_ipv4_address", "78.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "ipv4_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "ipv4_tcp_mss", "1460"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "clear_dont_fragment", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "dpd_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "dpd_retries", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "ike_version", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "ike_mode", "main"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "ike_rekey_interval", "14400"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "ike_ciphersuite", "aes256-cbc-sha1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "ike_group", "16"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "pre_shared_secret", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "ike_local_id", "xxx"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "ike_remote_id", "xxx"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "ipsec_rekey_interval", "3600"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "ipsec_replay_window", "512"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "ipsec_ciphersuite", "aes256-gcm"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "perfect_forward_secrecy", "group-16"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_gre_feature.test", "application_tunnel_type", "none"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanServiceLANVPNInterfaceGREPrerequisitesProfileParcelConfig + testAccSdwanServiceLANVPNInterfaceGREProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanServiceLANVPNInterfaceGREPrerequisitesProfileParcelConfig + testAccSdwanServiceLANVPNInterfaceGREProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanServiceLANVPNInterfaceGREPrerequisitesProfileParcelConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanServiceLANVPNInterfaceGREProfileParcelConfig_minimum() string {
	config := `resource "sdwan_service_lan_vpn_interface_gre_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id` + "\n"
	config += `	interface_name = "gre1"` + "\n"
	config += `	ipv4_address = "70.1.1.1"` + "\n"
	config += `	ipv4_subnet_mask = "255.255.255.0"` + "\n"
	config += `	tunnel_destination_ipv4_address = "79.1.1.1"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanServiceLANVPNInterfaceGREProfileParcelConfig_all() string {
	config := `resource "sdwan_service_lan_vpn_interface_gre_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id` + "\n"
	config += `	interface_name = "gre1"` + "\n"
	config += `	interface_description = "gre1"` + "\n"
	config += `	ipv4_address = "70.1.1.1"` + "\n"
	config += `	ipv4_subnet_mask = "255.255.255.0"` + "\n"
	config += `	shutdown = true` + "\n"
	config += `	tunnel_protection = false` + "\n"
	config += `	tunnel_mode = "ipv4"` + "\n"
	config += `	tunnel_source_ipv4_address = "78.1.1.1"` + "\n"
	config += `	tunnel_destination_ipv4_address = "79.1.1.1"` + "\n"
	config += `	ipv4_mtu = 1500` + "\n"
	config += `	ipv4_tcp_mss = 1460` + "\n"
	config += `	clear_dont_fragment = false` + "\n"
	config += `	dpd_interval = 10` + "\n"
	config += `	dpd_retries = 3` + "\n"
	config += `	ike_version = 1` + "\n"
	config += `	ike_mode = "main"` + "\n"
	config += `	ike_rekey_interval = 14400` + "\n"
	config += `	ike_ciphersuite = "aes256-cbc-sha1"` + "\n"
	config += `	ike_group = "16"` + "\n"
	config += `	pre_shared_secret = "123"` + "\n"
	config += `	ike_local_id = "xxx"` + "\n"
	config += `	ike_remote_id = "xxx"` + "\n"
	config += `	ipsec_rekey_interval = 3600` + "\n"
	config += `	ipsec_replay_window = 512` + "\n"
	config += `	ipsec_ciphersuite = "aes256-gcm"` + "\n"
	config += `	perfect_forward_secrecy = "group-16"` + "\n"
	config += `	application_tunnel_type = "none"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
