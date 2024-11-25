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
func TestAccDataSourceSdwanServiceLANVPNFeatureAssociateRoutingBGPFeature(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanServiceLANVPNFeatureAssociateRoutingBGPFeaturePrerequisitesConfig + testAccDataSourceSdwanServiceLANVPNFeatureAssociateRoutingBGPFeatureConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanServiceLANVPNFeatureAssociateRoutingBGPFeaturePrerequisitesConfig = `
resource "sdwan_service_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_service_lan_vpn_feature" "test" {
  name                       = "TF_TEST_SERVICE_LAN"
  description                = "TF_TEST"
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

resource "sdwan_service_routing_bgp_feature" "test" {
  name                     = "TF_TEST_ROUTING_BGP"
  description              = "Terraform test"
  feature_profile_id         = sdwan_service_feature_profile.test.id
  as_number                = 429
  router_id                = "1.2.3.4"
  propagate_as_path        = false
  propagate_community      = false
  external_routes_distance = 20
  internal_routes_distance = 200
  local_routes_distance    = 20
  keepalive_time           = 60
  hold_time                = 180
  always_compare_med       = false
  deterministic_med        = false
  missing_med_as_worst     = false
  compare_router_id        = false
  multipath_relax          = false
  ipv4_neighbors = [
    {
      address                 = "1.2.3.4"
      description             = "neighbor1"
      shutdown                = false
      remote_as               = 200
      local_as                = 200
      keepalive_time          = 40
      hold_time               = 200
      update_source_interface = "GigabitEthernet0"
      next_hop_self           = false
      send_community          = true
      send_extended_community = true
      ebgp_multihop           = 1
      password                = "myPassword"
      send_label              = true
      as_override             = false
      allowas_in_number       = 1
      address_families = [
        {
          family_type            = "ipv4-unicast"
          max_number_of_prefixes = 2000
          threshold              = 75
          policy_type            = "restart"
          restart_interval       = 30
        }
      ]
    }
  ]
  ipv6_neighbors = [
    {
      address                 = "2001::1"
      description             = "neighbor2"
      shutdown                = false
      remote_as               = 200
      local_as                = 200
      keepalive_time          = 180
      hold_time               = 60
      update_source_interface = "Loopback1"
      next_hop_self           = true
      send_community          = true
      send_extended_community = true
      ebgp_multihop           = 3
      password                = "myPassword"
      as_override             = true
      allowas_in_number       = 3
      address_families = [
        {
          family_type            = "ipv6-unicast"
          max_number_of_prefixes = 2000
          threshold              = 75
          policy_type            = "restart"
          restart_interval       = 30
        }
      ]
    }
  ]
  ipv4_aggregate_addresses = [
    {
      network_address = "10.10.0.0"
      subnet_mask     = "255.255.0.0"
      as_set_path     = false
      summary_only    = false
    }
  ]
  ipv4_networks = [
    {
      network_address = "10.10.0.0"
      subnet_mask     = "255.255.0.0"
    }
  ]
  ipv4_eibgp_maximum_paths = 1
  ipv4_originate           = false
  ipv4_table_map_filter    = false
  ipv6_aggregate_addresses = [
    {
      aggregate_prefix = "3001::1/128"
      as_set_path      = false
      summary_only     = false
    }
  ]
  ipv6_networks = [
    {
      network_prefix = "2001:0DB8:0000:000b::/64"
    }
  ]
  ipv6_eibgp_maximum_paths = 2
  ipv6_originate           = true
  ipv6_table_map_filter    = false
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanServiceLANVPNFeatureAssociateRoutingBGPFeatureConfig() string {
	config := ""
	config += `resource "sdwan_service_lan_vpn_feature_associate_routing_bgp_feature" "test" {` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id` + "\n"
	config += `	service_routing_bgp_feature_id = sdwan_service_routing_bgp_feature.test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_service_lan_vpn_feature_associate_routing_bgp_feature" "test" {
			feature_profile_id = sdwan_service_feature_profile.test.id
			service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id
			id = sdwan_service_lan_vpn_feature_associate_routing_bgp_feature.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
