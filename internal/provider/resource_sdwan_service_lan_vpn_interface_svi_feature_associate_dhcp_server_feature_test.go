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
func TestAccSdwanServiceLANVPNInterfaceSVIFeatureAssociateDHCPServerFeature(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanServiceLANVPNInterfaceSVIFeatureAssociateDHCPServerFeaturePrerequisitesConfig + testAccSdwanServiceLANVPNInterfaceSVIFeatureAssociateDHCPServerFeatureConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanServiceLANVPNInterfaceSVIFeatureAssociateDHCPServerFeaturePrerequisitesConfig = `
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

resource "sdwan_service_lan_vpn_interface_svi_feature" "test" {
  name                       = "TF_TEST_SERVICE_LAN_INTERFACE_SVI"
  description                = "Terraform test"
  feature_profile_id         = sdwan_service_feature_profile.test.id
  service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id
  shutdown                   = false
  interface_name             = "Vlan1"
  interface_description      = "SVI"
  interface_mtu              = 1500
  ip_mtu                     = 1500
  ipv4_address               = "1.2.3.4"
  ipv4_subnet_mask           = "0.0.0.0"
  ipv4_secondary_addresses = [
    {
      address          = "2.3.4.5"
      ipv4_subnet_mask = "0.0.0.0"
    }
  ]
  ipv4_dhcp_helpers = ["4.5.6.7"]
  ipv6_address      = "2001:0:0:1::0/32"
  ipv6_secondary_addresses = [
    {
      address = "::2/32"
    }
  ]
  ipv6_dhcp_helpers = [
    {
      address = "2001:0:0:1::0"
      vpn     = 1
    }
  ]
  arps = [
    {
      ip_address  = "1.2.3.4"
      mac_address = "00-B0-D0-63-C2-26"
    }
  ]
  ipv4_vrrps = [
    {
      group_id    = 1
      priority    = 100
      timer       = 1000
      track_omp   = false
      prefix_list = "prefix"
      address     = "1.2.3.4"
      secondary_addresses = [
        {
          address = "2.3.4.5"
        }
      ]
      tloc_prefix_change       = true
      tloc_prefix_change_value = 100
    }
  ]
  ipv6_vrrps = [
    {
      group_id          = 1
      priority          = 100
      timer             = 1000
      track_omp         = false
      track_prefix_list = "1"
      addresses = [
        {
          link_local_address = "1::1"
          global_address     = "1::1/24"
        }
      ]
      secondary_addresses = [
        {
          prefix = "::20/32"
        }
      ]
    }
  ]
  enable_dhcpv6         = false
  tcp_mss               = 1024
  arp_timeout           = 1200
  ip_directed_broadcast = false
  icmp_redirect_disable = true
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanServiceLANVPNInterfaceSVIFeatureAssociateDHCPServerFeatureConfig_all() string {
	config := `resource "sdwan_service_lan_vpn_interface_svi_feature_associate_dhcp_server_feature" "test" {` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id` + "\n"
	config += `	service_lan_vpn_interface_svi_feature_id = sdwan_service_lan_vpn_interface_svi_feature.test.id` + "\n"
	config += `	service_dhcp_server_feature_id = sdwan_service_dhcp_server_feature.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
