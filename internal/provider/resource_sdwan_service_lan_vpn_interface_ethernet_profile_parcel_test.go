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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanServiceLANVPNInterfaceEthernetProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "interface_name", "GigabitEthernet3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "config_description", "LAN"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv6_settings_dhcp_helper.0.dhcpv6_helper", "2001:0:0:1::0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv6_settings_dhcp_helper.0.dhcpv6_helper_vpn", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_nat", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_nat_type", "pool"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_nat_range_start", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_nat_range_end", "4.5.6.7"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_nat_prefix_length", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_nat_overload", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_nat_loopback", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_nat_udp_timeout", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_nat_tcp_timeout", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "static_nats.0.source_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "static_nats.0.translate_ip", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "static_nats.0.direction", "inside"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "static_nats.0.source_vpn", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv6_nat", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "enable_nat64", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "acl_shaping_rate", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv6_vrrps.0.group_id", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv6_vrrps.0.priority", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv6_vrrps.0.timer", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv6_vrrps.0.track_omp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv6_vrrps.0.addresses.0.ipv6_link_local_address", "1::1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv6_vrrps.0.addresses.0.global_ipv6_prefix", "1::1/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_vrrps.0.group_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_vrrps.0.priority", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_vrrps.0.timer", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_vrrps.0.track_omp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_vrrps.0.ip_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_vrrps.0.secondary_addresses.0.ip_address", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_vrrps.0.secondary_addresses.0.subnet_mask", "0.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_vrrps.0.tloc_prefix_change", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ipv4_vrrps.0.tloc_pref_change_value", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "arps.0.ip_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "arps.0.mac_address", "00-B0-D0-63-C2-26"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "trustsec_enable_sgt_propogation", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "trustsec_propogate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "trustsec_security_group_tags", "example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "trustsec_enable_enforced_propogation", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "trustsec_enforced_security_group_tags", "example-2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "duplex", "full"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "mac_address", "00-B0-D0-63-C2-26"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ip_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "interface_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "tcp_mss", "500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "speed", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "arp_timeout", "1200"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "autonegotiate", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "media_type", "auto-select"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "load_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "tracker", "TRACKER1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "icmp_redirect_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "xconnect", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_service_lan_vpn_interface_ethernet_profile_parcel.test", "ip_directed_broadcast", "false"))
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

const testAccSdwanServiceLANVPNInterfaceEthernetPrerequisitesProfileParcelConfig = `
resource "sdwan_service_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_service_lan_vpn_profile_parcel" "test" {
  name                       = "TF_TEST"
  description                = "Terraform test"
  feature_profile_id         = sdwan_service_feature_profile.test.id
  vpn                        = 1
  config_description         = "VPN1"
  omp_admin_distance_ipv4    = 1
  omp_admin_distance_ipv6    = 1
  enable_sdwan_remote_access = false
  primary_dns_address_ipv4   = "1.2.3.4"
  secondary_dns_address_ipv4 = "2.3.4.5"
  primary_dns_address_ipv6   = "2001:0:0:1::0"
  secondary_dns_address_ipv6 = "2001:0:0:2::0"
  host_mappings = [
    {
      host_name   = "HOSTMAPPING1"
      list_of_ips = ["1.2.3.4"]
    }
  ]
  ipv4_static_routes = [
    {
      network_address = "1.2.3.4"
      subnet_mask     = "0.0.0.0"
      next_hops = [
        {
          address                 = "1.2.3.4"
          administrative_distance = 1
        }
      ]
    }
  ]
  ipv6_static_routes = [
    {
      prefix = "2001:0:0:1::0/12"
      next_hops = [
        {
          address                 = "2001:0:0:1::0"
          administrative_distance = 1
        }
      ]
    }
  ]
  services = [
    {
      service_type   = "FW"
      ipv4_addresses = ["1.2.3.4"]
      tracking       = true
    }
  ]
  service_routes = [
    {
      network_address = "1.2.3.4"
      subnet_mask     = "0.0.0.0"
      service         = "SIG"
      vpn             = 0
    }
  ]
  gre_routes = [
    {
      network_address = "1.2.3.4"
      subnet_mask     = "0.0.0.0"
      interface       = ["gre01"]
      vpn             = 0
    }
  ]
  ipsec_routes = [
    {
      network_address = "1.2.3.4"
      subnet_mask     = "0.0.0.0"
      interface       = ["ipsec01"]
    }
  ]
  nat_pools = [
    {
      nat_pool_name = 1
      prefix_length = 3
      range_start   = "1.2.3.4"
      range_end     = "2.3.4.5"
      overload      = true
      direction     = "inside"
    }
  ]
  nat_port_forwards = [
    {
      nat_pool_name        = 2
      source_port          = 122
      translate_port       = 330
      source_ip            = "1.2.3.4"
      translated_source_ip = "2.3.4.5"
      protocol             = "TCP"
    }
  ]
  static_nats = [
    {
      nat_pool_name        = 3
      source_ip            = "1.2.3.4"
      translated_source_ip = "2.3.4.5"
      static_nat_direction = "inside"
    }
  ]
  nat_64_v4_pools = [
    {
      name        = "NATPOOL1"
      range_start = "1.2.3.4"
      range_end   = "2.3.4.5"
      overload    = false
    }
  ]
  ipv4_import_route_targets = [
    {
      route_target = "1.1.1.3:200"
    }
  ]
  ipv4_export_route_targets = [
    {
      route_target = "1.1.1.3:200"
    }
  ]
  ipv6_import_route_targets = [
    {
      route_target = "1.1.1.3:200"
    }
  ]
  ipv6_export_route_targets = [
    {
      route_target = "1.1.1.3:200"
    }
  ]
}
`

func testAccSdwanServiceLANVPNInterfaceEthernetProfileParcelConfig_minimum() string {
	config := `resource "sdwan_service_lan_vpn_interface_ethernet_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	profile_parcel_id = sdwan_service_lan_vpn_profile_parcel.test.id` + "\n"
	config += `	acl_ipv4_egress_policy_id = "1b270f6d-479b-47e3-ab0b-51bc6811a303"` + "\n"
	config += `	acl_ipv4_ingress_policy_id = "1b270f6d-479b-47e3-ab0b-51bc6811a303"` + "\n"
	config += `	acl_ipv6_egress_policy_id = "1b270f6d-479b-47e3-ab0b-51bc6811a303"` + "\n"
	config += `	acl_ipv6_ingress_policy_id = "1b270f6d-479b-47e3-ab0b-51bc6811a303"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccSdwanServiceLANVPNInterfaceEthernetProfileParcelConfig_all() string {
	config := `resource "sdwan_service_lan_vpn_interface_ethernet_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	profile_parcel_id = sdwan_service_lan_vpn_profile_parcel.test.id` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	interface_name = "GigabitEthernet3"` + "\n"
	config += `	config_description = "LAN"` + "\n"
	config += `	ipv4_settings_dhcp_helper = ["false"]` + "\n"
	config += `	ipv6_settings_dhcp_helper = [{` + "\n"
	config += `	  dhcpv6_helper = "2001:0:0:1::0"` + "\n"
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
	config += `	enable_nat64 = false` + "\n"
	config += `	acl_shaping_rate = 12` + "\n"
	config += `	ipv6_vrrps = [{` + "\n"
	config += `	  group_id = 0` + "\n"
	config += `	  priority = 100` + "\n"
	config += `	  timer = 1000` + "\n"
	config += `	  track_omp = false` + "\n"
	config += `	  addresses = [{` + "\n"
	config += `		ipv6_link_local_address = "1::1"` + "\n"
	config += `		global_ipv6_prefix = "1::1/24"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_vrrps = [{` + "\n"
	config += `	  group_id = 1` + "\n"
	config += `	  priority = 100` + "\n"
	config += `	  timer = 1000` + "\n"
	config += `	  track_omp = false` + "\n"
	config += `	  ip_address = "1.2.3.4"` + "\n"
	config += `	  secondary_addresses = [{` + "\n"
	config += `		ip_address = "2.3.4.5"` + "\n"
	config += `		subnet_mask = "0.0.0.0"` + "\n"
	config += `	}]` + "\n"
	config += `	  tloc_prefix_change = true` + "\n"
	config += `	  tloc_pref_change_value = 100` + "\n"
	config += `	}]` + "\n"
	config += `	arps = [{` + "\n"
	config += `	  ip_address = "1.2.3.4"` + "\n"
	config += `	  mac_address = "00-B0-D0-63-C2-26"` + "\n"
	config += `	}]` + "\n"
	config += `	trustsec_enable_sgt_propogation = false` + "\n"
	config += `	trustsec_propogate = true` + "\n"
	config += `	trustsec_security_group_tags = example` + "\n"
	config += `	trustsec_enable_enforced_propogation = false` + "\n"
	config += `	trustsec_enforced_security_group_tags = example-2` + "\n"
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
