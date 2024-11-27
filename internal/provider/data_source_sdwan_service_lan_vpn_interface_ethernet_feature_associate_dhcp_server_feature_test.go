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
func TestAccDataSourceSdwanServiceLANVPNInterfaceEthernetFeatureAssociateDHCPServerFeature(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanServiceLANVPNInterfaceEthernetFeatureAssociateDHCPServerFeaturePrerequisitesConfig + testAccDataSourceSdwanServiceLANVPNInterfaceEthernetFeatureAssociateDHCPServerFeatureConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanServiceLANVPNInterfaceEthernetFeatureAssociateDHCPServerFeaturePrerequisitesConfig = `
resource "sdwan_service_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_service_lan_vpn_feature" "test" {
  name                       = "TF_TEST_SERVICE_LAN"
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

resource "sdwan_service_lan_vpn_interface_ethernet_feature" "test" {
  name                       = "TF_TEST_SERVICE_LAN_INTERFACE_ETHERNET"
  description                = "Terraform test"
  feature_profile_id         = sdwan_service_feature_profile.test.id
  service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id
  shutdown                   = false
  interface_name             = "GigabitEthernet3"
  interface_description      = "LAN"
  ipv4_address               = "1.2.3.4"
  ipv4_subnet_mask           = "0.0.0.0"
  ipv4_secondary_addresses = [
    {
      address     = "1.2.3.5"
      subnet_mask = "0.0.0.0"
    }
  ]
  ipv4_dhcp_helper = ["1.2.3.4"]
  ipv6_dhcp_helpers = [
    {
      address           = "2001:0:0:1::0"
      dhcpv6_helper_vpn = 1
    }
  ]
  ipv4_nat               = false
  ipv4_nat_type          = "pool"
  ipv4_nat_range_start   = "1.2.3.4"
  ipv4_nat_range_end     = "4.5.6.7"
  ipv4_nat_prefix_length = 1
  ipv4_nat_overload      = true
  ipv4_nat_loopback      = "123"
  ipv4_nat_udp_timeout   = 123
  ipv4_nat_tcp_timeout   = 123
  static_nats = [
    {
      source_ip    = "1.2.3.4"
      translate_ip = "2.3.4.5"
      direction    = "inside"
      source_vpn   = 0
    }
  ]
  ipv6_nat         = true
  nat64            = false
  acl_shaping_rate = 12
  ipv6_vrrps = [
    {
      group_id  = 1
      priority  = 100
      timer     = 1000
      track_omp = false
      ipv6_addresses = [
        {
          link_local_address = "1::1"
          global_address     = "1::1/24"
        }
      ]
    }
  ]
  ipv4_vrrps = [
    {
      group_id  = 1
      priority  = 100
      timer     = 1000
      track_omp = false
      address   = "1.2.3.4"
      secondary_addresses = [
        {
          address     = "2.3.4.5"
          subnet_mask = "0.0.0.0"
        }
      ]
      tloc_prefix_change     = true
      tloc_pref_change_value = 100
    }
  ]
  arps = [
    {
      ip_address  = "1.2.3.4"
      mac_address = "00-B0-D0-63-C2-26"
    }
  ]
  trustsec_enable_sgt_propogation      = false
  trustsec_propogate                   = true
  trustsec_security_group_tag          = 123
  trustsec_enable_enforced_propogation = false
  trustsec_enforced_security_group_tag = 1234
  duplex                               = "full"
  mac_address                          = "00-B0-D0-63-C2-26"
  ip_mtu                               = 1500
  interface_mtu                        = 1500
  tcp_mss                              = 500
  speed                                = "1000"
  arp_timeout                          = 1200
  autonegotiate                        = false
  media_type                           = "auto-select"
  load_interval                        = 30
  tracker                              = "TRACKER1"
  icmp_redirect_disable                = true
  xconnect                             = "1"
  ip_directed_broadcast                = false
}

resource "sdwan_service_dhcp_server_feature" "test" {
  name               = "TF_TEST_DHCP_SERVER"
  description        = "Terraform Test"
  feature_profile_id = sdwan_service_feature_profile.test.id
  network_address    = "1.2.3.4"
  subnet_mask        = "255.255.255.0"
  exclude            = ["192.168.1.1"]
  lease_time         = 86400
  interface_mtu      = 65535
  domain_name        = "example.com"
  default_gateway    = "1.2.3.4"
  dns_servers        = ["8.8.8.8"]
  tftp_servers       = ["1.1.1.1"]
  static_leases = [
    {
      mac_address = "01:02:03:04:05:06"
      ip_address  = "1.2.3.4"
    }
  ]
  option_codes = [
    {
      code  = 250
      ascii = "example"
    }
  ]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanServiceLANVPNInterfaceEthernetFeatureAssociateDHCPServerFeatureConfig() string {
	config := ""
	config += `resource "sdwan_service_lan_vpn_interface_ethernet_feature_associate_dhcp_server_feature" "test" {` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id` + "\n"
	config += `	service_lan_vpn_interface_ethernet_feature_id = sdwan_service_lan_vpn_interface_ethernet_feature.test.id` + "\n"
	config += `	service_dhcp_server_feature_id = sdwan_service_dhcp_server_feature.test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_service_lan_vpn_interface_ethernet_feature_associate_dhcp_server_feature" "test" {
			feature_profile_id = sdwan_service_feature_profile.test.id
			service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id
			service_lan_vpn_interface_ethernet_feature_id = sdwan_service_lan_vpn_interface_ethernet_feature.test.id
			id = sdwan_service_lan_vpn_interface_ethernet_feature_associate_dhcp_server_feature.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
