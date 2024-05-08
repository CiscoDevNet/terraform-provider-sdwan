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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanVPNInterfaceT1E1SerialFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "serial_interface_name", "SERIAL1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "interface_description", "My description"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "ipv4_address", "1.2.3.4/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "ipv6_address", "2001:0:0:1::/64"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "ipv6_access_lists.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "ipv6_access_lists.0.acl_name", "ACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "enable_core_region", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "core_region", "core"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "secondary_region", "off"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_encapsulations.0.encapsulation", "gre"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_encapsulations.0.preference", "4294967"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_encapsulations.0.weight", "250"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_border", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "per_tunnel_qos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "per_tunnel_qos_aggregator", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_qos_mode", "spoke"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_color", "custom1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_last_resort_circuit", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_low_bandwidth_link", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_tunnel_tcp_mss", "1460"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_clear_dont_fragment", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_network_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_control_connections", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_vbond_as_stun_server", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_vmanage_connection_preference", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_port_hop", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_restrict", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_carrier", "carrier1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_nat_refresh_interval", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_hello_interval", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_hello_tolerance", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_bind_loopback_tunnel", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_allow_all", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_allow_bgp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_allow_dhcp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_allow_dns", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_allow_icmp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_allow_ssh", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_allow_ntp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_allow_netconf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_allow_ospf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_allow_stun", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_allow_snmp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tunnel_interface_allow_https", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "clear_dont_fragment_bit", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "pmtu_discovery", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "ip_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "static_ingress_qos", "6"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tcp_mss", "720"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "tloc_extension", "tloc"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "autonegotiate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "shaping_rate", "10000000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "qos_map", "test"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "qos_map_vpn", "test"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "interface_bandwidth_capacity", "128"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "clock_rate", "5300000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "encapsulation", "hdlc"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "interface_downstream_bandwidth_capacity", "10000000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "write_rule", "RULE1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "ipv4_access_lists.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_t1_e1_serial_feature_template.test", "ipv4_access_lists.0.acl_name", "ACL2"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanVPNInterfaceT1E1SerialFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanVPNInterfaceT1E1SerialFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanVPNInterfaceT1E1SerialFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_vpn_interface_t1_e1_serial_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanVPNInterfaceT1E1SerialFeatureTemplateConfig_all() string {
	config := `resource "sdwan_vpn_interface_t1_e1_serial_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	serial_interface_name = "SERIAL1"` + "\n"
	config += `	interface_description = "My description"` + "\n"
	config += `	ipv4_address = "1.2.3.4/24"` + "\n"
	config += `	ipv6_address = "2001:0:0:1::/64"` + "\n"
	config += `	ipv6_access_lists = [{` + "\n"
	config += `	  direction = "in"` + "\n"
	config += `	  acl_name = "ACL1"` + "\n"
	config += `	}]` + "\n"
	config += `	enable_core_region = true` + "\n"
	config += `	core_region = "core"` + "\n"
	config += `	secondary_region = "off"` + "\n"
	config += `	tunnel_interface_encapsulations = [{` + "\n"
	config += `	  encapsulation = "gre"` + "\n"
	config += `	  preference = 4294967` + "\n"
	config += `	  weight = 250` + "\n"
	config += `	}]` + "\n"
	config += `	tunnel_interface_groups = [42949672]` + "\n"
	config += `	tunnel_interface_border = true` + "\n"
	config += `	per_tunnel_qos = true` + "\n"
	config += `	per_tunnel_qos_aggregator = false` + "\n"
	config += `	tunnel_qos_mode = "spoke"` + "\n"
	config += `	tunnel_interface_color = "custom1"` + "\n"
	config += `	tunnel_interface_last_resort_circuit = false` + "\n"
	config += `	tunnel_interface_low_bandwidth_link = false` + "\n"
	config += `	tunnel_interface_tunnel_tcp_mss = 1460` + "\n"
	config += `	tunnel_interface_clear_dont_fragment = false` + "\n"
	config += `	tunnel_interface_network_broadcast = false` + "\n"
	config += `	tunnel_interface_control_connections = 8` + "\n"
	config += `	tunnel_interface_vbond_as_stun_server = false` + "\n"
	config += `	tunnel_interface_exclude_controller_group_list = [100]` + "\n"
	config += `	tunnel_interface_vmanage_connection_preference = 5` + "\n"
	config += `	tunnel_interface_port_hop = false` + "\n"
	config += `	tunnel_interface_restrict = false` + "\n"
	config += `	tunnel_interface_carrier = "carrier1"` + "\n"
	config += `	tunnel_interface_nat_refresh_interval = 15` + "\n"
	config += `	tunnel_interface_hello_interval = 1000` + "\n"
	config += `	tunnel_interface_hello_tolerance = 12` + "\n"
	config += `	tunnel_interface_bind_loopback_tunnel = "12"` + "\n"
	config += `	tunnel_interface_allow_all = false` + "\n"
	config += `	tunnel_interface_allow_bgp = false` + "\n"
	config += `	tunnel_interface_allow_dhcp = true` + "\n"
	config += `	tunnel_interface_allow_dns = true` + "\n"
	config += `	tunnel_interface_allow_icmp = true` + "\n"
	config += `	tunnel_interface_allow_ssh = false` + "\n"
	config += `	tunnel_interface_allow_ntp = false` + "\n"
	config += `	tunnel_interface_allow_netconf = false` + "\n"
	config += `	tunnel_interface_allow_ospf = false` + "\n"
	config += `	tunnel_interface_allow_stun = false` + "\n"
	config += `	tunnel_interface_allow_snmp = false` + "\n"
	config += `	tunnel_interface_allow_https = true` + "\n"
	config += `	clear_dont_fragment_bit = false` + "\n"
	config += `	pmtu_discovery = false` + "\n"
	config += `	ip_mtu = 1500` + "\n"
	config += `	static_ingress_qos = 6` + "\n"
	config += `	tcp_mss = 720` + "\n"
	config += `	tloc_extension = "tloc"` + "\n"
	config += `	shutdown = true` + "\n"
	config += `	autonegotiate = true` + "\n"
	config += `	shaping_rate = 10000000` + "\n"
	config += `	qos_map = "test"` + "\n"
	config += `	qos_map_vpn = "test"` + "\n"
	config += `	interface_bandwidth_capacity = 128` + "\n"
	config += `	clock_rate = "5300000"` + "\n"
	config += `	encapsulation = "hdlc"` + "\n"
	config += `	interface_downstream_bandwidth_capacity = 10000000` + "\n"
	config += `	write_rule = "RULE1"` + "\n"
	config += `	ipv4_access_lists = [{` + "\n"
	config += `	  direction = "in"` + "\n"
	config += `	  acl_name = "ACL2"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
