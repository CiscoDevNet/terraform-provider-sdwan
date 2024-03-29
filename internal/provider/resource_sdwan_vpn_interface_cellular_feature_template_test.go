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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanVPNInterfaceCellularFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanVPNInterfaceCellularFeatureTemplateConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "cellular_interface_name", "Cellular1"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "per_tunnel_qos", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "per_tunnel_qos_aggregator", "false"),
				),
			},
			{
				Config: testAccSdwanVPNInterfaceCellularFeatureTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "cellular_interface_name", "Cellular1"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "interface_description", "My Description"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "ipv6_access_lists.0.direction", "in"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "ipv6_access_lists.0.acl_name", "ACL1"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "nat", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "nat_refresh_mode", "outbound"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "nat_udp_timeout", "1"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "nat_tcp_timeout", "60"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "nat_block_icmp_error", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "nat_response_to_ping", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "nat_port_forwards.0.port_start_range", "0"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "nat_port_forwards.0.port_end_range", "65530"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "nat_port_forwards.0.protocol", "tcp"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "nat_port_forwards.0.private_vpn", "65530"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "nat_port_forwards.0.private_ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "enable_core_region", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "core_region", "core"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "secondary_region", "off"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_encapsulations.0.encapsulation", "gre"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_encapsulations.0.preference", "4294967"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_encapsulations.0.weight", "250"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_border", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "per_tunnel_qos", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "per_tunnel_qos_aggregator", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_qos_mode", "spoke"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_color", "custom1"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_last_resort_circuit", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_low_bandwidth_link", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_tunnel_tcp_mss", "1460"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_clear_dont_fragment", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_network_broadcast", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_max_control_connections", "8"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_control_connections", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_vbond_as_stun_server", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_vmanage_connection_preference", "5"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_port_hop", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_color_restrict", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_carrier", "carrier1"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_nat_refresh_interval", "15"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_hello_interval", "1000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_hello_tolerance", "12"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_bind_loopback_tunnel", "12"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_allow_all", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_allow_bgp", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_allow_dhcp", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_allow_dns", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_allow_icmp", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_allow_ssh", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_allow_ntp", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_allow_netconf", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_allow_ospf", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_allow_stun", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_allow_snmp", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tunnel_interface_allow_https", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "clear_dont_fragment_bit", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "pmtu_discovery", "false"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "ip_mtu", "1500"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "static_ingress_qos", "6"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tcp_mss", "720"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "tloc_extension", "tloc"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "ip_directed_broadcast", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "shutdown", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "autonegotiate", "true"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "qos_adaptive_period", "15"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "qos_adaptive_bandwidth_downstream", "10000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "qos_adaptive_min_downstream", "100"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "qos_adaptive_max_downstream", "100000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "qos_adaptive_bandwidth_upstream", "10000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "qos_adaptive_min_upstream", "100"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "qos_adaptive_max_upstream", "100000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "shaping_rate", "10000000"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "qos_map", "test"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "qos_map_vpn", "test"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "bandwidth_upstream", "214748300"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "bandwidth_downstream", "214748300"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "write_rule", "RULE1"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "ipv4_access_lists.0.direction", "in"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "ipv4_access_lists.0.acl_name", "ACL2"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "policers.0.direction", "in"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "policers.0.policer_name", "example"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "static_arps.0.ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr("sdwan_vpn_interface_cellular_feature_template.test", "static_arps.0.mac", "00-B0-D0-63-C2-26"),
				),
			},
		},
	})
}

func testAccSdwanVPNInterfaceCellularFeatureTemplateConfig_minimum() string {
	return `
	resource "sdwan_vpn_interface_cellular_feature_template" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		cellular_interface_name = "Cellular1"
		per_tunnel_qos = true
		per_tunnel_qos_aggregator = false
	}
	`
}

func testAccSdwanVPNInterfaceCellularFeatureTemplateConfig_all() string {
	return `
	resource "sdwan_vpn_interface_cellular_feature_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		cellular_interface_name = "Cellular1"
		interface_description = "My Description"
		ipv6_access_lists = [{
			direction = "in"
			acl_name = "ACL1"
		}]
		ipv4_dhcp_helper = ["6.6.6.6"]
		tracker = ["tracker1"]
		nat = true
		nat_refresh_mode = "outbound"
		nat_udp_timeout = 1
		nat_tcp_timeout = 60
		nat_block_icmp_error = true
		nat_response_to_ping = false
		nat_port_forwards = [{
			port_start_range = 0
			port_end_range = 65530
			protocol = "tcp"
			private_vpn = 65530
			private_ip_address = "1.2.3.4"
		}]
		enable_core_region = true
		core_region = "core"
		secondary_region = "off"
		tunnel_interface_encapsulations = [{
			encapsulation = "gre"
			preference = 4294967
			weight = 250
		}]
		tunnel_interface_groups = [42949672]
		tunnel_interface_border = true
		per_tunnel_qos = true
		per_tunnel_qos_aggregator = false
		tunnel_qos_mode = "spoke"
		tunnel_interface_color = "custom1"
		tunnel_interface_last_resort_circuit = false
		tunnel_interface_low_bandwidth_link = false
		tunnel_interface_tunnel_tcp_mss = 1460
		tunnel_interface_clear_dont_fragment = false
		tunnel_interface_network_broadcast = false
		tunnel_interface_max_control_connections = 8
		tunnel_interface_control_connections = true
		tunnel_interface_vbond_as_stun_server = false
		tunnel_interface_exclude_controller_group_list = [100]
		tunnel_interface_vmanage_connection_preference = 5
		tunnel_interface_port_hop = false
		tunnel_interface_color_restrict = false
		tunnel_interface_carrier = "carrier1"
		tunnel_interface_nat_refresh_interval = 15
		tunnel_interface_hello_interval = 1000
		tunnel_interface_hello_tolerance = 12
		tunnel_interface_bind_loopback_tunnel = "12"
		tunnel_interface_allow_all = false
		tunnel_interface_allow_bgp = false
		tunnel_interface_allow_dhcp = true
		tunnel_interface_allow_dns = true
		tunnel_interface_allow_icmp = true
		tunnel_interface_allow_ssh = false
		tunnel_interface_allow_ntp = false
		tunnel_interface_allow_netconf = false
		tunnel_interface_allow_ospf = false
		tunnel_interface_allow_stun = false
		tunnel_interface_allow_snmp = false
		tunnel_interface_allow_https = true
		clear_dont_fragment_bit = false
		pmtu_discovery = false
		ip_mtu = 1500
		static_ingress_qos = 6
		tcp_mss = 720
		tloc_extension = "tloc"
		ip_directed_broadcast = true
		shutdown = true
		autonegotiate = true
		qos_adaptive_period = 15
		qos_adaptive_bandwidth_downstream = 10000
		qos_adaptive_min_downstream = 100
		qos_adaptive_max_downstream = 100000
		qos_adaptive_bandwidth_upstream = 10000
		qos_adaptive_min_upstream = 100
		qos_adaptive_max_upstream = 100000
		shaping_rate = 10000000
		qos_map = "test"
		qos_map_vpn = "test"
		bandwidth_upstream = 214748300
		bandwidth_downstream = 214748300
		write_rule = "RULE1"
		ipv4_access_lists = [{
			direction = "in"
			acl_name = "ACL2"
		}]
		policers = [{
			direction = "in"
			policer_name = "example"
		}]
		static_arps = [{
			ip_address = "1.2.3.4"
			mac = "00-B0-D0-63-C2-26"
		}]
	}
	`
}
