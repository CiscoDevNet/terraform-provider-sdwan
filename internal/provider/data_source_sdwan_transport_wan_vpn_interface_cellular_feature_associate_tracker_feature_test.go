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
func TestAccDataSourceSdwanTransportWANVPNInterfaceCellularFeatureAssociateTrackerFeature(t *testing.T) {
	if os.Getenv("SDWAN_2015_IN_PROGRESS") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015_IN_PROGRESS")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanTransportWANVPNInterfaceCellularFeatureAssociateTrackerFeaturePrerequisitesConfig + testAccDataSourceSdwanTransportWANVPNInterfaceCellularFeatureAssociateTrackerFeatureConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanTransportWANVPNInterfaceCellularFeatureAssociateTrackerFeaturePrerequisitesConfig = `
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

resource "sdwan_transport_tracker_feature" "test" {
  name                  = "TF_TEST_TRACER"
  description           = "My Example"
  feature_profile_id    = sdwan_transport_feature_profile.test.id
  tracker_name          = "TRACKER_1"
}

resource "sdwan_transport_wan_vpn_interface_cellular_feature" "test" {
  name                         = "TF_TEST_INTERFACE_CELLULAR"
  description                  = "My Example"
  feature_profile_id           = sdwan_transport_feature_profile.test.id
  transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id
  shutdown                                       = true
  interface_name                                 = "GigabitEthernet1"
  interface_description                          = "WAN"
  ipv4_dhcp_helper                               = ["1.2.3.4"]
  service_provider                               = "example"
  bandwidth_upstream                             = 21474836
  bandwidth_downstream                           = 21474836
  tunnel_interface                               = true
  per_tunnel_qos                                 = true
  tunnel_qos_mode                                = "hub"
  tunnel_bandwidth_percent                       = 82
  tunnel_interface_bind_loopback_tunnel          = "example"
  tunnel_interface_carrier                       = "default"
  tunnel_interface_color                         = "default"
  tunnel_interface_hello_interval                = 1000
  tunnel_interface_hello_tolerance               = 12
  tunnel_interface_last_resort_circuit           = false
  tunnel_interface_color_restrict                = true
  tunnel_interface_groups                        = 42949672
  tunnel_interface_border                        = false
  tunnel_interface_max_control_connections       = 62
  tunnel_interface_nat_refresh_interval          = 5
  tunnel_interface_vbond_as_stun_server          = false
  tunnel_interface_exclude_controller_group_list = [2]
  tunnel_interface_vmanage_connection_preference = 8
  tunnel_interface_port_hop                      = true
  tunnel_interface_low_bandwidth_link            = false
  tunnel_interface_tunnel_tcp_mss                = 1460
  tunnel_interface_clear_dont_fragment           = false
  tunnel_interface_network_broadcast             = false
  tunnel_interface_allow_all                     = false
  tunnel_interface_allow_bgp                     = false
  tunnel_interface_allow_dhcp                    = true
  tunnel_interface_allow_ntp                     = false
  tunnel_interface_allow_ssh                     = false
  tunnel_interface_allow_dns                     = true
  tunnel_interface_allow_icmp                    = true
  tunnel_interface_allow_https                   = true
  tunnel_interface_allow_ospf                    = false
  tunnel_interface_allow_stun                    = false
  tunnel_interface_allow_snmp                    = false
  tunnel_interface_allow_netconf                 = false
  tunnel_interface_allow_bfd                     = false
  tunnel_interface_encapsulations = [
    {
      encapsulation = "gre"
      preference    = 4294967
      weight        = 250
    }
  ]
  nat_ipv4                    = true
  nat_udp_timeout             = 1
  nat_tcp_timeout             = 60
  qos_adaptive                = false
  qos_shaping_rate            = 16
  arps = [
    {
      ip_address  = "1.2.3.4"
      mac_address = "00-B0-D0-63-C2-26"
    }
  ]
  ip_mtu                = 1500
  interface_mtu         = 1500
  tcp_mss               = 505
  tloc_extension        = "tloc"
  tracker               = "example"
  ip_directed_broadcast = false
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanTransportWANVPNInterfaceCellularFeatureAssociateTrackerFeatureConfig() string {
	config := ""
	config += `resource "sdwan_transport_wan_vpn_interface_cellular_feature_associate_tracker_feature" "test" {` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id` + "\n"
	config += `	transport_wan_vpn_interface_cellular_feature_id = sdwan_transport_wan_vpn_interface_cellular_feature.test.id` + "\n"
	config += `	transport_tracker_feature_id = sdwan_transport_tracker_feature.test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_transport_wan_vpn_interface_cellular_feature_associate_tracker_feature" "test" {
			feature_profile_id = sdwan_transport_feature_profile.test.id
			transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id
			transport_wan_vpn_interface_cellular_feature_id = sdwan_transport_wan_vpn_interface_cellular_feature.test.id
			id = sdwan_transport_wan_vpn_interface_cellular_feature_associate_tracker_feature.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
