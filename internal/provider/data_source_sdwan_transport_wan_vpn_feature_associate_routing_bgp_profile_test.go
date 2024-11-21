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
func TestAccDataSourceSdwanTransportWANVPNFeatureAssociateRoutingBGPProfile(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanTransportWANVPNFeatureAssociateRoutingBGPProfilePrerequisitesConfig + testAccDataSourceSdwanTransportWANVPNFeatureAssociateRoutingBGPProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanTransportWANVPNFeatureAssociateRoutingBGPProfilePrerequisitesConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_transport_wan_vpn_feature" "test" {
  name               = "TF_TEST_WAN_VPN"
  description        = "Terraform test"
  feature_profile_id = sdwan_transport_feature_profile.test.id
  vpn                = 0
}

resource "sdwan_transport_routing_bgp_feature" "test" {
  name                     = "TF_TEST_ROUTING_BGP"
  description              = "Terraform test"
  feature_profile_id       = sdwan_transport_feature_profile.test.id
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
      explicit_null           = false
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
  mpls_interfaces = [
    {
      interface_name = "GigabitEthernet1"
    }
  ]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanTransportWANVPNFeatureAssociateRoutingBGPProfileConfig() string {
	config := ""
	config += `resource "sdwan_transport_wan_vpn_feature_associate_routing_bgp_profile" "test" {` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id` + "\n"
	config += `	transport_routing_bgp_feature_id = sdwan_transport_routing_bgp_feature.test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_transport_wan_vpn_feature_associate_routing_bgp_profile" "test" {
			feature_profile_id = sdwan_transport_feature_profile.test.id
			transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id
			id = sdwan_transport_wan_vpn_feature_associate_routing_bgp_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
