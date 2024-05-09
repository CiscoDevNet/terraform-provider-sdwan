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
func TestAccSdwanVPNInterfaceDSLIPoEFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ethernet_interface_name", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ethernet_ipv4_address", "1.2.3.4/24"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ethernet_enable_dhcp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ethernet_dhcp_distance", "1234"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "internal_controller_type", "ipoe"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ethernet_description", "My Description"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vdsl_configurations.0.controller_vdsl_slot", "Example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vdsl_configurations.0.sra", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vdsl_configurations.0.mode_adsl1", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vdsl_configurations.0.mode_adsl2", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vdsl_configurations.0.mode_adsl2plus", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vdsl_configurations.0.mode_vdsl2", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vdsl_configurations.0.mode_ansi", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "vdsl_configurations.0.vdsl_modem_configuration", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "encap", "4094"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "dialer_pool_number", "255"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ppp_maximum_payload", "1790"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "dialer_address_negotiated", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "unnumbered_loopback_interface", "example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ppp_authentication_protocol", "chap"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ppp_authentication_protocol_pap", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "chap_hostname", "chap-example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "chap_ppp_auth_password", "myPassword"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "pap_username", "pap-username"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "pap_password", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "pap_ppp_auth_password", "myPassword"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_encapsulations.0.encapsulation", "gre"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_encapsulations.0.preference", "4294967"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_encapsulations.0.weight", "250"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_border", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "per_tunnel_qos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "per_tunnel_qos_aggregator", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_qos_mode", "spoke"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_color", "custom1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_last_resort_circuit", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_low_bandwidth_link", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_tunnel_tcp_mss", "1460"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_clear_dont_fragment", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_network_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_max_control_connections", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_control_connections", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_vbond_as_stun_server", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_vmanage_connection_preference", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_port_hop", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_color_restrict", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_carrier", "carrier1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_nat_refresh_interval", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_hello_interval", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_hello_tolerance", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_bind_loopback_tunnel", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_allow_all", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_allow_bgp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_allow_dhcp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_allow_dns", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_allow_icmp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_allow_ssh", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_allow_ntp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_allow_netconf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_allow_ospf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_allow_stun", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_allow_snmp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tunnel_interface_allow_https", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "nat", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "nat_refresh_mode", "outbound"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "nat_udp_timeout", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "nat_tcp_timeout", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "nat_block_icmp_error", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "nat_response_to_ping", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "nat_port_forwards.0.port_start_range", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "nat_port_forwards.0.port_end_range", "65530"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "nat_port_forwards.0.protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "nat_port_forwards.0.private_vpn", "65530"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "nat_port_forwards.0.private_ip_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "qos_adaptive_period", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "qos_adaptive_bandwidth_downstream", "10000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "qos_adaptive_min_downstream", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "qos_adaptive_max_downstream", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "qos_adaptive_bandwidth_upstream", "10000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "qos_adaptive_min_upstream", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "qos_adaptive_max_upstream", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "shaping_rate", "10000000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "qos_map", "test"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "qos_map_vpn", "test"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "bandwidth_upstream", "214748300"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "bandwidth_downstream", "214748300"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "write_rule", "RULE1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "access_lists.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "access_lists.0.acl_name", "ACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "policers.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "policers.0.policer_name", "example"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ip_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tcp_mss", "720"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "tloc_extension", "tloc"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_vpn_interface_dsl_ipoe_feature_template.test", "ip_directed_broadcast", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanVPNInterfaceDSLIPoEFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanVPNInterfaceDSLIPoEFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanVPNInterfaceDSLIPoEFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_vpn_interface_dsl_ipoe_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanVPNInterfaceDSLIPoEFeatureTemplateConfig_all() string {
	config := `resource "sdwan_vpn_interface_dsl_ipoe_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	ethernet_interface_name = "Example"` + "\n"
	config += `	ethernet_ipv4_address = "1.2.3.4/24"` + "\n"
	config += `	ethernet_enable_dhcp = false` + "\n"
	config += `	ethernet_dhcp_distance = 1234` + "\n"
	config += `	ethernet_dhcp_helper = ["3"]` + "\n"
	config += `	internal_controller_type = "ipoe"` + "\n"
	config += `	shutdown = true` + "\n"
	config += `	ethernet_description = "My Description"` + "\n"
	config += `	vdsl_configurations = [{` + "\n"
	config += `	  controller_vdsl_slot = "Example"` + "\n"
	config += `	  sra = true` + "\n"
	config += `	  mode_adsl1 = false` + "\n"
	config += `	  mode_adsl2 = false` + "\n"
	config += `	  mode_adsl2plus = false` + "\n"
	config += `	  mode_vdsl2 = false` + "\n"
	config += `	  mode_ansi = false` + "\n"
	config += `	  vdsl_modem_configuration = "100"` + "\n"
	config += `	}]` + "\n"
	config += `	encap = 4094` + "\n"
	config += `	dialer_pool_number = 255` + "\n"
	config += `	ppp_maximum_payload = 1790` + "\n"
	config += `	dialer_address_negotiated = false` + "\n"
	config += `	unnumbered_loopback_interface = "example"` + "\n"
	config += `	ppp_authentication_protocol = "chap"` + "\n"
	config += `	ppp_authentication_protocol_pap = false` + "\n"
	config += `	chap_hostname = "chap-example"` + "\n"
	config += `	chap_ppp_auth_password = "myPassword"` + "\n"
	config += `	pap_username = "pap-username"` + "\n"
	config += `	pap_password = true` + "\n"
	config += `	pap_ppp_auth_password = "myPassword"` + "\n"
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
	config += `	tunnel_interface_max_control_connections = 8` + "\n"
	config += `	tunnel_interface_control_connections = true` + "\n"
	config += `	tunnel_interface_vbond_as_stun_server = false` + "\n"
	config += `	tunnel_interface_exclude_controller_group_list = [100]` + "\n"
	config += `	tunnel_interface_vmanage_connection_preference = 5` + "\n"
	config += `	tunnel_interface_port_hop = false` + "\n"
	config += `	tunnel_interface_color_restrict = false` + "\n"
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
	config += `	nat = true` + "\n"
	config += `	nat_refresh_mode = "outbound"` + "\n"
	config += `	nat_udp_timeout = 1` + "\n"
	config += `	nat_tcp_timeout = 60` + "\n"
	config += `	nat_block_icmp_error = true` + "\n"
	config += `	nat_response_to_ping = false` + "\n"
	config += `	nat_port_forwards = [{` + "\n"
	config += `	  port_start_range = 0` + "\n"
	config += `	  port_end_range = 65530` + "\n"
	config += `	  protocol = "tcp"` + "\n"
	config += `	  private_vpn = 65530` + "\n"
	config += `	  private_ip_address = "1.2.3.4"` + "\n"
	config += `	}]` + "\n"
	config += `	qos_adaptive_period = 15` + "\n"
	config += `	qos_adaptive_bandwidth_downstream = 10000` + "\n"
	config += `	qos_adaptive_min_downstream = 100` + "\n"
	config += `	qos_adaptive_max_downstream = 100000` + "\n"
	config += `	qos_adaptive_bandwidth_upstream = 10000` + "\n"
	config += `	qos_adaptive_min_upstream = 100` + "\n"
	config += `	qos_adaptive_max_upstream = 100000` + "\n"
	config += `	shaping_rate = 10000000` + "\n"
	config += `	qos_map = "test"` + "\n"
	config += `	qos_map_vpn = "test"` + "\n"
	config += `	bandwidth_upstream = 214748300` + "\n"
	config += `	bandwidth_downstream = 214748300` + "\n"
	config += `	write_rule = "RULE1"` + "\n"
	config += `	access_lists = [{` + "\n"
	config += `	  direction = "in"` + "\n"
	config += `	  acl_name = "ACL1"` + "\n"
	config += `	}]` + "\n"
	config += `	policers = [{` + "\n"
	config += `	  direction = "in"` + "\n"
	config += `	  policer_name = "example"` + "\n"
	config += `	}]` + "\n"
	config += `	ip_mtu = 1500` + "\n"
	config += `	tcp_mss = 720` + "\n"
	config += `	tloc_extension = "tloc"` + "\n"
	config += `	tracker = ["tracker1"]` + "\n"
	config += `	ip_directed_broadcast = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
