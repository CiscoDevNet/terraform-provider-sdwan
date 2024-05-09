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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceSdwanCiscoVPNInterfaceFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "interface_name", "ge0/0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "interface_description", "My Interface Description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "poe", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "address", "1.1.1.1/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv4_secondary_addresses.0.address", "2.2.2.2/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "dhcp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "dhcp_distance", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv6_address", "2001:1::1/48"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "dhcpv6", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv6_secondary_addresses.0.address", "2.2.2.2/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv6_access_lists.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv6_access_lists.0.acl_name", "ACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv6_dhcp_helpers.0.address", "2001:7::7/48"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv6_dhcp_helpers.0.vpn_id", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "auto_bandwidth_detect", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "iperf_server", "8.8.8.8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "nat", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "nat_type", "interface"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "udp_timeout", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tcp_timeout", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "nat_pool_range_start", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "nat_pool_range_end", "10.1.1.255"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "nat_overload", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "nat_inside_source_loopback_interface", "lo1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "nat_pool_prefix_length", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv6_nat", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "nat64_interface", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "nat66_interface", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_nat66_entries.0.source_prefix", "2001:7::/48"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_nat66_entries.0.translated_source_prefix", "2001:8::/48"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_nat66_entries.0.source_vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_nat_entries.0.source_ip", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_nat_entries.0.translate_ip", "100.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_nat_entries.0.static_nat_direction", "inside"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_nat_entries.0.source_vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_port_forward_entries.0.source_ip", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_port_forward_entries.0.translate_ip", "100.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_port_forward_entries.0.static_nat_direction", "inside"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_port_forward_entries.0.source_port", "8000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_port_forward_entries.0.translate_port", "9000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_port_forward_entries.0.protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_port_forward_entries.0.source_vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "enable_core_region", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "core_region", "core"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "secondary_region", "off"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_encapsulations.0.encapsulation", "gre"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_encapsulations.0.preference", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_encapsulations.0.weight", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_border", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_qos_mode", "spoke"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_bandwidth", "50"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_color", "gold"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_max_control_connections", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_control_connections", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_vbond_as_stun_server", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_vmanage_connection_preference", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_port_hop", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_color_restrict", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_gre_tunnel_destination_ip", "5.5.5.5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_carrier", "carrier1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_nat_refresh_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_hello_interval", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_hello_tolerance", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_bind_loopback_tunnel", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_last_resort_circuit", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_low_bandwidth_link", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_tunnel_tcp_mss", "1460"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_clear_dont_fragment", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_propagate_sgt", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_network_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_all", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_bgp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_dhcp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_dns", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_icmp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_ssh", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_netconf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_ntp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_ospf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_stun", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_snmp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tunnel_interface_allow_https", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "media_type", "auto-select"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "interface_mtu", "9216"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ip_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tcp_mss_adjust", "1460"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "tloc_extension", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "load_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "gre_tunnel_source_ip", "3.3.3.3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "gre_tunnel_xconnect", "a123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "mac_address", "00-B0-D0-63-C2-26"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "speed", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "duplex", "full"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "arp_timeout", "1200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "autonegotiate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ip_directed_broadcast", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "icmp_redirect_disable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_period", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_bandwidth_downstream", "10000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_min_downstream", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_max_downstream", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_bandwidth_upstream", "10000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_min_upstream", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "qos_adaptive_max_upstream", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "shaping_rate", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "qos_map", "QOSMAP1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "qos_map_vpn", "QOSMAP2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "bandwidth_upstream", "10000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "bandwidth_downstream", "10000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "block_non_source_ip", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "rewrite_rule_name", "RULE1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "access_lists.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "access_lists.0.acl_name", "ACL1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_arps.0.ip_address", "8.8.8.8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_arps.0.mac", "00-B0-D0-63-C2-26"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.group_id", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.priority", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.timer", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.track_omp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.track_prefix_list", "PL1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.ip_address", "2.2.2.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.ipv4_secondary_addresses.0.ip_address", "2.2.2.3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.tloc_preference_change", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.tloc_preference_change_value", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.tracking_objects.0.tracker_id", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.tracking_objects.0.track_action", "decrement"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv4_vrrps.0.tracking_objects.0.decrement_value", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv6_vrrps.0.group_id", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv6_vrrps.0.priority", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv6_vrrps.0.timer", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv6_vrrps.0.track_omp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv6_vrrps.0.track_prefix_list", "PL1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv6_vrrps.0.ipv6_addresses.0.ipv6_link_local", "fe80::260:8ff:fe52:f9d8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "ipv6_vrrps.0.ipv6_addresses.0.prefix", "2001:9::/48"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "propagate_sgt", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_sgt", "1003"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "static_sgt_trusted", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "enable_sgt", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "sgt_enforcement", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_interface_feature_template.test", "sgt_enforcement_sgt", "1004"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCiscoVPNInterfaceFeatureTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanCiscoVPNInterfaceFeatureTemplateConfig() string {
	config := `resource "sdwan_cisco_vpn_interface_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	interface_name = "ge0/0"` + "\n"
	config += `	interface_description = "My Interface Description"` + "\n"
	config += `	poe = false` + "\n"
	config += `	address = "1.1.1.1/24"` + "\n"
	config += `	ipv4_secondary_addresses = [{` + "\n"
	config += `	  address = "2.2.2.2/24"` + "\n"
	config += `	}]` + "\n"
	config += `	dhcp = false` + "\n"
	config += `	dhcp_distance = 10` + "\n"
	config += `	ipv6_address = "2001:1::1/48"` + "\n"
	config += `	dhcpv6 = false` + "\n"
	config += `	ipv6_secondary_addresses = [{` + "\n"
	config += `	  address = "2.2.2.2/24"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_access_lists = [{` + "\n"
	config += `	  direction = "in"` + "\n"
	config += `	  acl_name = "ACL1"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_dhcp_helper = ["6.6.6.6"]` + "\n"
	config += `	ipv6_dhcp_helpers = [{` + "\n"
	config += `	  address = "2001:7::7/48"` + "\n"
	config += `	  vpn_id = 5` + "\n"
	config += `	}]` + "\n"
	config += `	tracker = ["tracker1"]` + "\n"
	config += `	auto_bandwidth_detect = false` + "\n"
	config += `	iperf_server = "8.8.8.8"` + "\n"
	config += `	nat = true` + "\n"
	config += `	nat_type = "interface"` + "\n"
	config += `	udp_timeout = 1` + "\n"
	config += `	tcp_timeout = 60` + "\n"
	config += `	nat_pool_range_start = "10.1.1.1"` + "\n"
	config += `	nat_pool_range_end = "10.1.1.255"` + "\n"
	config += `	nat_overload = false` + "\n"
	config += `	nat_inside_source_loopback_interface = "lo1"` + "\n"
	config += `	nat_pool_prefix_length = 24` + "\n"
	config += `	ipv6_nat = false` + "\n"
	config += `	nat64_interface = false` + "\n"
	config += `	nat66_interface = false` + "\n"
	config += `	static_nat66_entries = [{` + "\n"
	config += `	  source_prefix = "2001:7::/48"` + "\n"
	config += `	  translated_source_prefix = "2001:8::/48"` + "\n"
	config += `	  source_vpn_id = 1` + "\n"
	config += `	}]` + "\n"
	config += `	static_nat_entries = [{` + "\n"
	config += `	  source_ip = "10.1.1.1"` + "\n"
	config += `	  translate_ip = "100.1.1.1"` + "\n"
	config += `	  static_nat_direction = "inside"` + "\n"
	config += `	  source_vpn_id = 1` + "\n"
	config += `	}]` + "\n"
	config += `	static_port_forward_entries = [{` + "\n"
	config += `	  source_ip = "10.1.1.1"` + "\n"
	config += `	  translate_ip = "100.1.1.1"` + "\n"
	config += `	  static_nat_direction = "inside"` + "\n"
	config += `	  source_port = 8000` + "\n"
	config += `	  translate_port = 9000` + "\n"
	config += `	  protocol = "tcp"` + "\n"
	config += `	  source_vpn_id = 1` + "\n"
	config += `	}]` + "\n"
	config += `	enable_core_region = false` + "\n"
	config += `	core_region = "core"` + "\n"
	config += `	secondary_region = "off"` + "\n"
	config += `	tunnel_interface_encapsulations = [{` + "\n"
	config += `	  encapsulation = "gre"` + "\n"
	config += `	  preference = 10` + "\n"
	config += `	  weight = 100` + "\n"
	config += `	}]` + "\n"
	config += `	tunnel_interface_border = false` + "\n"
	config += `	tunnel_qos_mode = "spoke"` + "\n"
	config += `	tunnel_bandwidth = 50` + "\n"
	config += `	tunnel_interface_groups = [5]` + "\n"
	config += `	tunnel_interface_color = "gold"` + "\n"
	config += `	tunnel_interface_max_control_connections = 10` + "\n"
	config += `	tunnel_interface_control_connections = false` + "\n"
	config += `	tunnel_interface_vbond_as_stun_server = false` + "\n"
	config += `	tunnel_interface_exclude_controller_group_list = [10]` + "\n"
	config += `	tunnel_interface_vmanage_connection_preference = 5` + "\n"
	config += `	tunnel_interface_port_hop = false` + "\n"
	config += `	tunnel_interface_color_restrict = false` + "\n"
	config += `	tunnel_interface_gre_tunnel_destination_ip = "5.5.5.5"` + "\n"
	config += `	tunnel_interface_carrier = "carrier1"` + "\n"
	config += `	tunnel_interface_nat_refresh_interval = 5` + "\n"
	config += `	tunnel_interface_hello_interval = 1000` + "\n"
	config += `	tunnel_interface_hello_tolerance = 12` + "\n"
	config += `	tunnel_interface_bind_loopback_tunnel = "1"` + "\n"
	config += `	tunnel_interface_last_resort_circuit = false` + "\n"
	config += `	tunnel_interface_low_bandwidth_link = false` + "\n"
	config += `	tunnel_interface_tunnel_tcp_mss = 1460` + "\n"
	config += `	tunnel_interface_clear_dont_fragment = false` + "\n"
	config += `	tunnel_interface_propagate_sgt = false` + "\n"
	config += `	tunnel_interface_network_broadcast = false` + "\n"
	config += `	tunnel_interface_allow_all = false` + "\n"
	config += `	tunnel_interface_allow_bgp = false` + "\n"
	config += `	tunnel_interface_allow_dhcp = false` + "\n"
	config += `	tunnel_interface_allow_dns = false` + "\n"
	config += `	tunnel_interface_allow_icmp = false` + "\n"
	config += `	tunnel_interface_allow_ssh = false` + "\n"
	config += `	tunnel_interface_allow_netconf = false` + "\n"
	config += `	tunnel_interface_allow_ntp = false` + "\n"
	config += `	tunnel_interface_allow_ospf = false` + "\n"
	config += `	tunnel_interface_allow_stun = false` + "\n"
	config += `	tunnel_interface_allow_snmp = false` + "\n"
	config += `	tunnel_interface_allow_https = false` + "\n"
	config += `	media_type = "auto-select"` + "\n"
	config += `	interface_mtu = 9216` + "\n"
	config += `	ip_mtu = 1500` + "\n"
	config += `	tcp_mss_adjust = 1460` + "\n"
	config += `	tloc_extension = "123"` + "\n"
	config += `	load_interval = 30` + "\n"
	config += `	gre_tunnel_source_ip = "3.3.3.3"` + "\n"
	config += `	gre_tunnel_xconnect = "a123"` + "\n"
	config += `	mac_address = "00-B0-D0-63-C2-26"` + "\n"
	config += `	speed = "1000"` + "\n"
	config += `	duplex = "full"` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	arp_timeout = 1200` + "\n"
	config += `	autonegotiate = true` + "\n"
	config += `	ip_directed_broadcast = false` + "\n"
	config += `	icmp_redirect_disable = false` + "\n"
	config += `	qos_adaptive_period = 15` + "\n"
	config += `	qos_adaptive_bandwidth_downstream = 10000` + "\n"
	config += `	qos_adaptive_min_downstream = 100` + "\n"
	config += `	qos_adaptive_max_downstream = 100000` + "\n"
	config += `	qos_adaptive_bandwidth_upstream = 10000` + "\n"
	config += `	qos_adaptive_min_upstream = 100` + "\n"
	config += `	qos_adaptive_max_upstream = 100000` + "\n"
	config += `	shaping_rate = 1000` + "\n"
	config += `	qos_map = "QOSMAP1"` + "\n"
	config += `	qos_map_vpn = "QOSMAP2"` + "\n"
	config += `	bandwidth_upstream = 10000` + "\n"
	config += `	bandwidth_downstream = 10000` + "\n"
	config += `	block_non_source_ip = false` + "\n"
	config += `	rewrite_rule_name = "RULE1"` + "\n"
	config += `	access_lists = [{` + "\n"
	config += `	  direction = "in"` + "\n"
	config += `	  acl_name = "ACL1"` + "\n"
	config += `	}]` + "\n"
	config += `	static_arps = [{` + "\n"
	config += `	  ip_address = "8.8.8.8"` + "\n"
	config += `	  mac = "00-B0-D0-63-C2-26"` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_vrrps = [{` + "\n"
	config += `	  group_id = 100` + "\n"
	config += `	  priority = 100` + "\n"
	config += `	  timer = 100` + "\n"
	config += `	  track_omp = false` + "\n"
	config += `	  track_prefix_list = "PL1"` + "\n"
	config += `	  ip_address = "2.2.2.2"` + "\n"
	config += `	  ipv4_secondary_addresses = [{` + "\n"
	config += `		ip_address = "2.2.2.3"` + "\n"
	config += `	}]` + "\n"
	config += `	  tloc_preference_change = false` + "\n"
	config += `	  tloc_preference_change_value = 10` + "\n"
	config += `	  tracking_objects = [{` + "\n"
	config += `		tracker_id = 10` + "\n"
	config += `		track_action = "decrement"` + "\n"
	config += `		decrement_value = 100` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_vrrps = [{` + "\n"
	config += `	  group_id = 100` + "\n"
	config += `	  priority = 100` + "\n"
	config += `	  timer = 100` + "\n"
	config += `	  track_omp = false` + "\n"
	config += `	  track_prefix_list = "PL1"` + "\n"
	config += `	  ipv6_addresses = [{` + "\n"
	config += `		ipv6_link_local = "fe80::260:8ff:fe52:f9d8"` + "\n"
	config += `		prefix = "2001:9::/48"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	propagate_sgt = false` + "\n"
	config += `	static_sgt = 1003` + "\n"
	config += `	static_sgt_trusted = false` + "\n"
	config += `	enable_sgt = true` + "\n"
	config += `	sgt_enforcement = true` + "\n"
	config += `	sgt_enforcement_sgt = 1004` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_cisco_vpn_interface_feature_template" "test" {
			id = sdwan_cisco_vpn_interface_feature_template.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
