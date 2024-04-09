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

func TestAccDataSourceSdwanVPNInterfaceMultilinkFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "interface_name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "multilink_group_number", "2147483"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "interface_description", "My Description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "ipv4_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "ipv6_address", "2001:0:0:1::/64"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "ipv6_access_lists.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "ipv6_access_lists.0.acl_name", "Egress ACL - IPv4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "ppp_authentication_protocol", "chap"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "ppp_authentication_protocol_pap", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "authentication_type", "callin"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "chap_hostname", "chap-example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "chap_ppp_auth_password", "myPassword"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "pap_username", "pap-username"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "pap_password", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "pap_ppp_auth_password", "myPassword"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "enable_core_region", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "core_region", "core"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "secondary_region", "off"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "encapsulation.0.encapsulation_type", "gre"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "encapsulation.0.preference", "4294967"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "encapsulation.0.weight", "250"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "border", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "per_tunnel_qos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "per_tunnel_qos_aggregator", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "color", "custom1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "last_resort_circuit", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "low_bandwidth_link", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "tunnel_tcp_mss", "1460"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "enable_clear_dont_fragment", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "network_broadcast_1", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "max_control_connections", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "control_connections", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "vbond_as_stun_server", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "vmanage_connection_preference", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "port_hop", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "restrict", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "carrier", "carrier1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "nat_refresh_interval", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "hello_interval", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "hello_tolerance", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "bind_loopback_tunnel", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "all", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "network_broadcast_2", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "bgp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "dhcp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "dns", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "icmp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "ssh", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "netconf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "ospf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "stun", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "snmp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "https", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "disable_fragmentation", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "fragment_max_delay", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "interleaving_fragment", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "clear_dont_fragment_bit", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "pmtu_discovery", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "ip_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "static_ingress_qos", "6"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "tcp_mss", "720"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "ip_directed_broadcast", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "tloc_extension", "tloc"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "administrative_shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "link_autonegotiate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "shaping_rate", "10000000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "qos_map", "test"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "vpn_qos_map", "test"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "bandwidth_upstream", "214748300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "bandwidth_downstream", "214748300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "write_rule", "test_rule"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "access_list.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "access_list.0.acl_name", "Egress ACL - IPv4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "controller_tx_ex_list.0.card_type", "E1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "controller_tx_ex_list.0.slot", "interface-t1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "controller_tx_ex_list.0.framing", "example-framing"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "controller_tx_ex_list.0.line_mode", "primary"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "controller_tx_ex_list.0.internal", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "controller_tx_ex_list.0.description", "example-interface"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "controller_tx_ex_list.0.linecode", "ami"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "controller_tx_ex_list.0.set_length_for_long", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "controller_tx_ex_list.0.set_length_for_short", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "controller_tx_ex_list.0.channel_group_list.0.channel_group", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "nim_interface_list.0.nim_serial_interface_type", "2t"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "nim_interface_list.0.interface_name", "nim-interface"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "nim_interface_list.0.interface_description", "interface description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "nim_interface_list.0.bandwidth", "21474836"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "nim_interface_list.0.clock_rate", "120000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_vpn_interface_multilink_feature_template.test", "nim_interface_list.0.encapsulation_serial", "hdlc"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanVPNInterfaceMultilinkFeatureTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceSdwanVPNInterfaceMultilinkFeatureTemplateConfig() string {
	config := `resource "sdwan_vpn_interface_multilink_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	interface_name = "Example"` + "\n"
	config += `	multilink_group_number = 2147483` + "\n"
	config += `	interface_description = "My Description"` + "\n"
	config += `	ipv4_address = "1.2.3.4"` + "\n"
	config += `	ipv6_address = "2001:0:0:1::/64"` + "\n"
	config += `	ipv6_access_lists = [{` + "\n"
	config += `	  direction = "in"` + "\n"
	config += `	  acl_name = "Egress ACL - IPv4"` + "\n"
	config += `	}]` + "\n"
	config += `	ppp_authentication_protocol = "chap"` + "\n"
	config += `	ppp_authentication_protocol_pap = false` + "\n"
	config += `	authentication_type = "callin"` + "\n"
	config += `	chap_hostname = "chap-example"` + "\n"
	config += `	chap_ppp_auth_password = "myPassword"` + "\n"
	config += `	pap_username = "pap-username"` + "\n"
	config += `	pap_password = true` + "\n"
	config += `	pap_ppp_auth_password = "myPassword"` + "\n"
	config += `	enable_core_region = true` + "\n"
	config += `	core_region = "core"` + "\n"
	config += `	secondary_region = "off"` + "\n"
	config += `	encapsulation = [{` + "\n"
	config += `	  encapsulation_type = "gre"` + "\n"
	config += `	  preference = 4294967` + "\n"
	config += `	  weight = 250` + "\n"
	config += `	}]` + "\n"
	config += `	groups = [42949672]` + "\n"
	config += `	border = true` + "\n"
	config += `	per_tunnel_qos = true` + "\n"
	config += `	per_tunnel_qos_aggregator = false` + "\n"
	config += `	color = "custom1"` + "\n"
	config += `	last_resort_circuit = false` + "\n"
	config += `	low_bandwidth_link = false` + "\n"
	config += `	tunnel_tcp_mss = 1460` + "\n"
	config += `	enable_clear_dont_fragment = false` + "\n"
	config += `	network_broadcast_1 = false` + "\n"
	config += `	max_control_connections = 8` + "\n"
	config += `	control_connections = true` + "\n"
	config += `	vbond_as_stun_server = false` + "\n"
	config += `	exclude_controller_group_list = [100]` + "\n"
	config += `	vmanage_connection_preference = 5` + "\n"
	config += `	port_hop = false` + "\n"
	config += `	restrict = false` + "\n"
	config += `	carrier = "carrier1"` + "\n"
	config += `	nat_refresh_interval = 15` + "\n"
	config += `	hello_interval = 1000` + "\n"
	config += `	hello_tolerance = 12` + "\n"
	config += `	bind_loopback_tunnel = "12"` + "\n"
	config += `	all = false` + "\n"
	config += `	network_broadcast_2 = false` + "\n"
	config += `	bgp = false` + "\n"
	config += `	dhcp = true` + "\n"
	config += `	dns = true` + "\n"
	config += `	icmp = true` + "\n"
	config += `	ssh = false` + "\n"
	config += `	netconf = false` + "\n"
	config += `	ospf = false` + "\n"
	config += `	stun = false` + "\n"
	config += `	snmp = false` + "\n"
	config += `	https = true` + "\n"
	config += `	disable_fragmentation = true` + "\n"
	config += `	fragment_max_delay = 1` + "\n"
	config += `	interleaving_fragment = false` + "\n"
	config += `	clear_dont_fragment_bit = false` + "\n"
	config += `	pmtu_discovery = false` + "\n"
	config += `	ip_mtu = 1500` + "\n"
	config += `	static_ingress_qos = 6` + "\n"
	config += `	tcp_mss = 720` + "\n"
	config += `	ip_directed_broadcast = true` + "\n"
	config += `	tloc_extension = "tloc"` + "\n"
	config += `	administrative_shutdown = true` + "\n"
	config += `	link_autonegotiate = true` + "\n"
	config += `	shaping_rate = 10000000` + "\n"
	config += `	qos_map = "test"` + "\n"
	config += `	vpn_qos_map = "test"` + "\n"
	config += `	bandwidth_upstream = 214748300` + "\n"
	config += `	bandwidth_downstream = 214748300` + "\n"
	config += `	write_rule = "test_rule"` + "\n"
	config += `	access_list = [{` + "\n"
	config += `	  direction = "in"` + "\n"
	config += `	  acl_name = "Egress ACL - IPv4"` + "\n"
	config += `	}]` + "\n"
	config += `	controller_tx_ex_list = [{` + "\n"
	config += `	  card_type = "E1"` + "\n"
	config += `	  slot = "interface-t1"` + "\n"
	config += `	  framing = "example-framing"` + "\n"
	config += `	  line_mode = "primary"` + "\n"
	config += `	  internal = false` + "\n"
	config += `	  description = "example-interface"` + "\n"
	config += `	  linecode = "ami"` + "\n"
	config += `	  set_length_for_long = "100"` + "\n"
	config += `	  set_length_for_short = "100"` + "\n"
	config += `	  channel_group_list = [{` + "\n"
	config += `		channel_group = 30` + "\n"
	config += `		timeslots = ["example"]` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	nim_interface_list = [{` + "\n"
	config += `	  nim_serial_interface_type = "2t"` + "\n"
	config += `	  interface_name = "nim-interface"` + "\n"
	config += `	  interface_description = "interface description"` + "\n"
	config += `	  bandwidth = 21474836` + "\n"
	config += `	  clock_rate = 120000` + "\n"
	config += `	  encapsulation_serial = "hdlc"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_vpn_interface_multilink_feature_template" "test" {
			id = sdwan_vpn_interface_multilink_feature_template.test.id
		}
	`
	return config
}
