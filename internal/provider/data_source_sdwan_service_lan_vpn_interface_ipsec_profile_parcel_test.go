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
func TestAccDataSourceSdwanServiceLANVPNInterfaceIPSecProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "interface_name", "ipsec987"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "interface_description", "ipsec987"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "ipv4_address", "9.7.5.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "ipv4_subnet_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "tunnel_source_ipv4_address", "1.3.5.88"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "tunnel_source_ipv4_subnet_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "tunnel_source_interface", "GigabitEthernet8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "tunnel_destination_ipv4_address", "2.55.67.99"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "tunnel_destination_ipv4_subnet_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "application_tunnel_type", "none"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "tcp_mss", "1460"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "clear_dont_fragment", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "ip_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "dpd_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "dpd_retries", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "ike_preshared_key", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "ike_version", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "ike_integrity_protocol", "main"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "ike_rekey_interval", "14400"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "ike_ciphersuite", "aes256-cbc-sha1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "ike_diffie_hellman_group", "16"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "ike_id_local_end_point", "xxx"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "ike_id_remote_end_point", "xxx"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "ipsec_rekey_interval", "3600"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "ipsec_replay_window", "512"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "ipsec_ciphersuite", "aes256-gcm"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "perfect_forward_secrecy", "group-16"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_lan_vpn_interface_ipsec_feature.test", "tunnel_route_via", "2222"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanServiceLANVPNInterfaceIPSecPrerequisitesProfileParcelConfig + testAccDataSourceSdwanServiceLANVPNInterfaceIPSecProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanServiceLANVPNInterfaceIPSecPrerequisitesProfileParcelConfig = `
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
func testAccDataSourceSdwanServiceLANVPNInterfaceIPSecProfileParcelConfig() string {
	config := `resource "sdwan_service_lan_vpn_interface_ipsec_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	service_lan_vpn_profile_parcel_id = sdwan_service_lan_vpn_feature.test.id` + "\n"
	config += `	interface_name = "ipsec987"` + "\n"
	config += `	shutdown = true` + "\n"
	config += `	interface_description = "ipsec987"` + "\n"
	config += `	ipv4_address = "9.7.5.4"` + "\n"
	config += `	ipv4_subnet_mask = "255.255.255.0"` + "\n"
	config += `	tunnel_source_ipv4_address = "1.3.5.88"` + "\n"
	config += `	tunnel_source_ipv4_subnet_mask = "255.255.255.0"` + "\n"
	config += `	tunnel_source_interface = "GigabitEthernet8"` + "\n"
	config += `	tunnel_destination_ipv4_address = "2.55.67.99"` + "\n"
	config += `	tunnel_destination_ipv4_subnet_mask = "255.255.255.0"` + "\n"
	config += `	application_tunnel_type = "none"` + "\n"
	config += `	tcp_mss = 1460` + "\n"
	config += `	clear_dont_fragment = false` + "\n"
	config += `	ip_mtu = 1500` + "\n"
	config += `	dpd_interval = 10` + "\n"
	config += `	dpd_retries = 3` + "\n"
	config += `	ike_preshared_key = "123"` + "\n"
	config += `	ike_version = 1` + "\n"
	config += `	ike_integrity_protocol = "main"` + "\n"
	config += `	ike_rekey_interval = 14400` + "\n"
	config += `	ike_ciphersuite = "aes256-cbc-sha1"` + "\n"
	config += `	ike_diffie_hellman_group = "16"` + "\n"
	config += `	ike_id_local_end_point = "xxx"` + "\n"
	config += `	ike_id_remote_end_point = "xxx"` + "\n"
	config += `	ipsec_rekey_interval = 3600` + "\n"
	config += `	ipsec_replay_window = 512` + "\n"
	config += `	ipsec_ciphersuite = "aes256-gcm"` + "\n"
	config += `	perfect_forward_secrecy = "group-16"` + "\n"
	config += `	tunnel_route_via = "2222"` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_service_lan_vpn_interface_ipsec_feature" "test" {
			id = sdwan_service_lan_vpn_interface_ipsec_feature.test.id
			feature_profile_id = sdwan_service_feature_profile.test.id
			service_lan_vpn_profile_parcel_id = sdwan_service_lan_vpn_feature.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
