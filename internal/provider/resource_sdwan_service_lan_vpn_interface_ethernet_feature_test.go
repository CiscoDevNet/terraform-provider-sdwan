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
func TestAccSdwanServiceLANVPNInterfaceEthernetProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "interface_name", "GigabitEthernet3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "interface_description", "LAN"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_subnet_mask", "0.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_secondary_addresses.0.address", "1.2.3.5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_secondary_addresses.0.subnet_mask", "0.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv6_dhcp_helpers.0.address", "2001:0:0:1::0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv6_dhcp_helpers.0.dhcpv6_helper_vpn", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_nat", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_nat_type", "pool"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_nat_range_start", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_nat_range_end", "4.5.6.7"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_nat_prefix_length", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_nat_overload", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_nat_loopback", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_nat_udp_timeout", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_nat_tcp_timeout", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "static_nats.0.source_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "static_nats.0.translate_ip", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "static_nats.0.direction", "inside"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "static_nats.0.source_vpn", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv6_nat", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "nat64", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "acl_shaping_rate", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv6_vrrps.0.group_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv6_vrrps.0.priority", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv6_vrrps.0.timer", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv6_vrrps.0.track_omp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv6_vrrps.0.ipv6_addresses.0.link_local_address", "1::1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv6_vrrps.0.ipv6_addresses.0.global_address", "1::1/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_vrrps.0.group_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_vrrps.0.priority", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_vrrps.0.timer", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_vrrps.0.track_omp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_vrrps.0.address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_vrrps.0.secondary_addresses.0.address", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_vrrps.0.secondary_addresses.0.subnet_mask", "0.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_vrrps.0.tloc_prefix_change", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_vrrps.0.tloc_pref_change_value", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_vrrps.0.tracking_objects.0.tracker_action", "Decrement"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ipv4_vrrps.0.tracking_objects.0.decrement_value", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "arps.0.ip_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "arps.0.mac_address", "00-B0-D0-63-C2-26"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "trustsec_enable_sgt_propogation", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "trustsec_propogate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "trustsec_security_group_tag", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "trustsec_enable_enforced_propogation", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "trustsec_enforced_security_group_tag", "1234"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "duplex", "full"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "mac_address", "00-B0-D0-63-C2-26"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ip_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "interface_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "tcp_mss", "500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "speed", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "arp_timeout", "1200"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "autonegotiate", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "media_type", "auto-select"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "load_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "tracker", "TRACKER1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "icmp_redirect_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "xconnect", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_feature.test", "ip_directed_broadcast", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanServiceLANVPNInterfaceEthernetPrerequisitesProfileParcelConfig + testAccSdwanServiceLANVPNInterfaceEthernetProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanServiceLANVPNInterfaceEthernetPrerequisitesProfileParcelConfig + testAccSdwanServiceLANVPNInterfaceEthernetProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanServiceLANVPNInterfaceEthernetPrerequisitesProfileParcelConfig = `
resource "sdwan_service_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_service_lan_vpn_feature" "test" {
  name = "TF_TEST_SLAN"
  feature_profile_id = sdwan_service_feature_profile.test.id
}

resource "sdwan_service_tracker_feature" "test" {
  name                  = "TF_TEST_TRACKER"
  description           = "Terraform test"
  feature_profile_id    = sdwan_service_feature_profile.test.id
  tracker_name          = "TRACKER_1"
  endpoint_api_url      = "google.com"
  endpoint_dns_name     = "google.com"
  endpoint_ip           = "1.2.3.4"
  interval              = 30
  multiplier            = 3
  threshold             = 300
  endpoint_tracker_type = "static-route"
  tracker_type          = "endpoint"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanServiceLANVPNInterfaceEthernetProfileParcelConfig_minimum() string {
	config := `resource "sdwan_service_lan_vpn_interface_ethernet_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id` + "\n"
	config += `	interface_name = "GigabitEthernet3"` + "\n"
	config += `	ipv4_address = "1.2.3.4"` + "\n"
	config += `	ipv4_subnet_mask = "0.0.0.0"` + "\n"
	config += `	ipv4_nat_type = "pool"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanServiceLANVPNInterfaceEthernetProfileParcelConfig_all() string {
	config := `resource "sdwan_service_lan_vpn_interface_ethernet_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	interface_name = "GigabitEthernet3"` + "\n"
	config += `	interface_description = "LAN"` + "\n"
	config += `	ipv4_address = "1.2.3.4"` + "\n"
	config += `	ipv4_subnet_mask = "0.0.0.0"` + "\n"
	config += `	ipv4_secondary_addresses = [{` + "\n"
	config += `	  address = "1.2.3.5"` + "\n"
	config += `	  subnet_mask = "0.0.0.0"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_dhcp_helper = ["1.2.3.4"]` + "\n"
	config += `	ipv6_dhcp_helpers = [{` + "\n"
	config += `	  address = "2001:0:0:1::0"` + "\n"
	config += `	  dhcpv6_helper_vpn = 1` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_nat = false` + "\n"
	config += `	ipv4_nat_type = "pool"` + "\n"
	config += `	ipv4_nat_range_start = "1.2.3.4"` + "\n"
	config += `	ipv4_nat_range_end = "4.5.6.7"` + "\n"
	config += `	ipv4_nat_prefix_length = 1` + "\n"
	config += `	ipv4_nat_overload = true` + "\n"
	config += `	ipv4_nat_loopback = "123"` + "\n"
	config += `	ipv4_nat_udp_timeout = 123` + "\n"
	config += `	ipv4_nat_tcp_timeout = 123` + "\n"
	config += `	static_nats = [{` + "\n"
	config += `	  source_ip = "1.2.3.4"` + "\n"
	config += `	  translate_ip = "2.3.4.5"` + "\n"
	config += `	  direction = "inside"` + "\n"
	config += `	  source_vpn = 0` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_nat = true` + "\n"
	config += `	nat64 = false` + "\n"
	config += `	acl_shaping_rate = 12` + "\n"
	config += `	ipv6_vrrps = [{` + "\n"
	config += `	  group_id = 1` + "\n"
	config += `	  priority = 100` + "\n"
	config += `	  timer = 1000` + "\n"
	config += `	  track_omp = false` + "\n"
	config += `	  ipv6_addresses = [{` + "\n"
	config += `		link_local_address = "1::1"` + "\n"
	config += `		global_address = "1::1/24"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_vrrps = [{` + "\n"
	config += `	  group_id = 1` + "\n"
	config += `	  priority = 100` + "\n"
	config += `	  timer = 1000` + "\n"
	config += `	  track_omp = false` + "\n"
	config += `	  address = "1.2.3.4"` + "\n"
	config += `	  secondary_addresses = [{` + "\n"
	config += `		address = "2.3.4.5"` + "\n"
	config += `		subnet_mask = "0.0.0.0"` + "\n"
	config += `	}]` + "\n"
	config += `	  tloc_prefix_change = true` + "\n"
	config += `	  tloc_pref_change_value = 100` + "\n"
	config += `	  tracking_objects = [{` + "\n"
	config += `		tracker_id = sdwan_service_tracker_feature.test.id` + "\n"
	config += `		tracker_action = "Decrement"` + "\n"
	config += `		decrement_value = 100` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	arps = [{` + "\n"
	config += `	  ip_address = "1.2.3.4"` + "\n"
	config += `	  mac_address = "00-B0-D0-63-C2-26"` + "\n"
	config += `	}]` + "\n"
	config += `	trustsec_enable_sgt_propogation = false` + "\n"
	config += `	trustsec_propogate = true` + "\n"
	config += `	trustsec_security_group_tag = 123` + "\n"
	config += `	trustsec_enable_enforced_propogation = false` + "\n"
	config += `	trustsec_enforced_security_group_tag = 1234` + "\n"
	config += `	duplex = "full"` + "\n"
	config += `	mac_address = "00-B0-D0-63-C2-26"` + "\n"
	config += `	ip_mtu = 1500` + "\n"
	config += `	interface_mtu = 1500` + "\n"
	config += `	tcp_mss = 500` + "\n"
	config += `	speed = "1000"` + "\n"
	config += `	arp_timeout = 1200` + "\n"
	config += `	autonegotiate = false` + "\n"
	config += `	media_type = "auto-select"` + "\n"
	config += `	load_interval = 30` + "\n"
	config += `	tracker = "TRACKER1"` + "\n"
	config += `	icmp_redirect_disable = true` + "\n"
	config += `	xconnect = "1"` + "\n"
	config += `	ip_directed_broadcast = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
