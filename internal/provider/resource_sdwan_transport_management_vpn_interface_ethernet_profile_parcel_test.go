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
func TestAccSdwanTransportManagementVPNInterfaceEthernetProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "interface_name", "GigabitEthernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "interface_description", "Transport Management VPN Interface Ethernet"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "ipv4_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "ipv4_subnet_mask", "0.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "ipv4_secondary_addresses.0.address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "ipv4_secondary_addresses.0.subnet_mask", "0.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "ipv4_iperf_server", "example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "ipv4_auto_detect_bandwidth", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "ipv6_address", "2001:0:0:1::/64"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "arp_entries.0.ip_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "arp_entries.0.mac_address", "00-B0-D0-63-C2-26"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "duplex", "full"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "mac_address", "00-B0-D0-63-C2-26"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "ip_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "interface_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "tcp_mss", "505"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "speed", "2500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "arp_timeout", "1200"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "autonegotiate", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "media_type", "rj45"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "load_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "icmp_redirect_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_management_vpn_interface_ethernet_feature.test", "ip_directed_broadcast", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTransportManagementVPNInterfaceEthernetPrerequisitesProfileParcelConfig + testAccSdwanTransportManagementVPNInterfaceEthernetProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanTransportManagementVPNInterfaceEthernetPrerequisitesProfileParcelConfig + testAccSdwanTransportManagementVPNInterfaceEthernetProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanTransportManagementVPNInterfaceEthernetPrerequisitesProfileParcelConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_transport_management_vpn_feature" "test" {
  name = "TF_TEST_VPN"
  feature_profile_id = sdwan_transport_feature_profile.test.id
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanTransportManagementVPNInterfaceEthernetProfileParcelConfig_minimum() string {
	config := `resource "sdwan_transport_management_vpn_interface_ethernet_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	transport_management_vpn_profile_parcel_id = sdwan_transport_management_vpn_feature.test.id` + "\n"
	config += `	interface_name = "GigabitEthernet1"` + "\n"
	config += `	ipv4_address = "1.2.3.4"` + "\n"
	config += `	ipv4_subnet_mask = "0.0.0.0"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTransportManagementVPNInterfaceEthernetProfileParcelConfig_all() string {
	config := `resource "sdwan_transport_management_vpn_interface_ethernet_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	transport_management_vpn_profile_parcel_id = sdwan_transport_management_vpn_feature.test.id` + "\n"
	config += `	shutdown = true` + "\n"
	config += `	interface_name = "GigabitEthernet1"` + "\n"
	config += `	interface_description = "Transport Management VPN Interface Ethernet"` + "\n"
	config += `	ipv4_address = "1.2.3.4"` + "\n"
	config += `	ipv4_subnet_mask = "0.0.0.0"` + "\n"
	config += `	ipv4_secondary_addresses = [{` + "\n"
	config += `	  address = "1.2.3.4"` + "\n"
	config += `	  subnet_mask = "0.0.0.0"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_dhcp_helper = ["1.2.3.4"]` + "\n"
	config += `	ipv4_iperf_server = "example"` + "\n"
	config += `	ipv4_auto_detect_bandwidth = false` + "\n"
	config += `	ipv6_address = "2001:0:0:1::/64"` + "\n"
	config += `	arp_entries = [{` + "\n"
	config += `	  ip_address = "1.2.3.4"` + "\n"
	config += `	  mac_address = "00-B0-D0-63-C2-26"` + "\n"
	config += `	}]` + "\n"
	config += `	duplex = "full"` + "\n"
	config += `	mac_address = "00-B0-D0-63-C2-26"` + "\n"
	config += `	ip_mtu = 1500` + "\n"
	config += `	interface_mtu = 1500` + "\n"
	config += `	tcp_mss = 505` + "\n"
	config += `	speed = "2500"` + "\n"
	config += `	arp_timeout = 1200` + "\n"
	config += `	autonegotiate = false` + "\n"
	config += `	media_type = "rj45"` + "\n"
	config += `	load_interval = 30` + "\n"
	config += `	icmp_redirect_disable = true` + "\n"
	config += `	ip_directed_broadcast = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
