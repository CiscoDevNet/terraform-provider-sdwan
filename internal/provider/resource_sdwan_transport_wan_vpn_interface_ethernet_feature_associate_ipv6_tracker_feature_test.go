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
func TestAccSdwanTransportWANVPNInterfaceEthernetFeatureAssociateIPv6TrackerFeature(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTransportWANVPNInterfaceEthernetFeatureAssociateIPv6TrackerFeaturePrerequisitesConfig + testAccSdwanTransportWANVPNInterfaceEthernetFeatureAssociateIPv6TrackerFeatureConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanTransportWANVPNInterfaceEthernetFeatureAssociateIPv6TrackerFeaturePrerequisitesConfig = `
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

resource "sdwan_transport_ipv6_tracker_feature" "test" {
  name                  = "TF_TEST_TRACER"
  description           = "My Example"
  feature_profile_id    = sdwan_transport_feature_profile.test.id
  tracker_name          = "TRACKER_1"
  endpoint_api_url      = "google.com"
  endpoint_dns_name     = "google.com"
  endpoint_ip           = "2001:0:0:1::0"
  interval              = 30
  multiplier            = 3
  threshold             = 300
  endpoint_tracker_type = "ipv6-interface"
  tracker_type          = "endpoint"
}

resource "sdwan_transport_wan_vpn_interface_ethernet_feature" "test" {
  name                         = "TF_TEST_INTERFACE_ETHERNET"
  description                  = "My Example"
  feature_profile_id           = sdwan_transport_feature_profile.test.id
  transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id
  shutdown                     = true
  interface_name               = "GigabitEthernet1"
  interface_description        = "WAN"
  ipv4_configuration_type      = "static"
  ipv4_address                 = "1.2.3.4"
  ipv4_subnet_mask             = "0.0.0.0"
  ipv4_secondary_addresses = [
    {
      address     = "1.2.3.4"
      subnet_mask = "0.0.0.0"
    }
  ]
  ipv4_dhcp_helper                               = ["1.2.3.4"]
  ipv6_configuration_type                        = "static"
  ipv6_address                                   = "2001:0:0:1::/64"
  iperf_server                                   = "example"
  block_non_source_ip                            = false
  service_provider                               = "example"
  bandwidth_upstream                             = 21474836
  bandwidth_downstream                           = 21474836
  auto_detect_bandwidth                          = false
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
  tunnel_interface_gre_tunnel_destination_ip     = "1.2.3.4"
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
  tunnel_interface_cts_sgt_propagation           = false
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
  nat_ipv4        = true
  nat_type        = "interface"
  nat_udp_timeout = 1
  nat_tcp_timeout = 60
  new_static_nats = [
    {
      source_ip     = "1.2.3.4"
      translated_ip = "2.3.4.5"
      direction     = "inside"
      source_vpn    = 3
    }
  ]
  nat_ipv6 = true
  nat64    = false
  nat66    = true
  static_nat66 = [
    {
      source_prefix            = "2001:0db8:85a3::/48"
      translated_source_prefix = "abcd:1234:5678::/48"
      source_vpn_id            = 4
    }
  ]
  arps = [
    {
      ip_address  = "1.2.3.4"
      mac_address = "00-B0-D0-63-C2-26"
    }
  ]
  icmp_redirect_disable = true
  duplex                = "full"
  mac_address           = "00-B0-D0-63-C2-26"
  ip_mtu                = 1500
  interface_mtu         = 1500
  tcp_mss               = 505
  speed                 = "2500"
  arp_timeout           = 1200
  autonegotiate         = false
  media_type            = "rj45"
  tloc_extension        = "tloc"
  gre_tunnel_source_ip  = "1.2.3.4"
  xconnect              = "example"
  load_interval         = 30
  tracker               = "example"
  ip_directed_broadcast = false
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTransportWANVPNInterfaceEthernetFeatureAssociateIPv6TrackerFeatureConfig_all() string {
	config := `resource "sdwan_transport_wan_vpn_interface_ethernet_feature_associate_ipv6_tracker_feature" "test" {` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id` + "\n"
	config += `	transport_wan_vpn_interface_ethernet_feature_id = sdwan_transport_wan_vpn_interface_ethernet_feature.test.id` + "\n"
	config += `	transport_ipv6_tracker_feature_id = sdwan_transport_ipv6_tracker_feature.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
