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
func TestAccDataSourceSdwanCiscoBGPFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "as_number", "65000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "router_id", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "propagate_aspath", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "propagate_community", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_route_targets.0.vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_route_targets.0.export.0.asn_ip", "10:100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_route_targets.0.import.0.asn_ip", "10:100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_route_targets.0.vpn_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_route_targets.0.export.0.asn_ip", "10:100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_route_targets.0.import.0.asn_ip", "10:100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "mpls_interfaces.0.interface_name", "GigabitEthernet0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "distance_external", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "distance_internal", "210"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "distance_local", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "keepalive", "90"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "holdtime", "220"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "always_compare_med", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "deterministic_med", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "missing_med_worst", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "compare_router_id", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "multipath_relax", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.family_type", "ipv4-unicast"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.ipv4_aggregate_addresses.0.prefix", "10.0.0.0/8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.ipv4_aggregate_addresses.0.as_set_path", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.ipv4_aggregate_addresses.0.summary_only", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.ipv4_networks.0.prefix", "10.2.2.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.maximum_paths", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.default_information_originate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.table_map_policy", "MAP1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.table_map_filter", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.redistribute_routes.0.protocol", "ospf"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.redistribute_routes.0.route_policy", "POLICY1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.address", "10.2.2.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.description", "My neighbor"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.remote_as", "65001"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.keepalive", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.holdtime", "90"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.source_interface", "GigabitEthernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.next_hop_self", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.send_community", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.send_ext_community", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.ebgp_multihop", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.password", "cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.send_label", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.send_label_explicit", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.as_override", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.allow_as_in", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.address_families.0.family_type", "ipv4-unicast"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.address_families.0.maximum_prefixes", "10000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.address_families.0.maximum_prefixes_threshold", "80"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.address_families.0.maximum_prefixes_restart", "180"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.address_families.0.maximum_prefixes_warning_only", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.address_families.0.route_policies.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.address_families.0.route_policies.0.policy_name", "POLICY1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.address", "2001:1::1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.description", "My neighbor"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.remote_as", "65001"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.keepalive", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.holdtime", "90"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.source_interface", "GigabitEthernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.next_hop_self", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.send_community", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.send_ext_community", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.ebgp_multihop", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.password", "cisco123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.send_label", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.send_label_explicit", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.as_override", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.allow_as_in", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.address_families.0.family_type", "ipv6-unicast"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.address_families.0.maximum_prefixes", "10000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.address_families.0.maximum_prefixes_threshold", "80"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.address_families.0.maximum_prefixes_restart", "180"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.address_families.0.maximum_prefixes_warning_only", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.address_families.0.route_policies.0.direction", "in"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.address_families.0.route_policies.0.policy_name", "POLICY1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCiscoBGPFeatureTemplateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanCiscoBGPFeatureTemplateConfig() string {
	config := `resource "sdwan_cisco_bgp_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	as_number = "65000"` + "\n"
	config += `	shutdown = true` + "\n"
	config += `	router_id = "1.2.3.4"` + "\n"
	config += `	propagate_aspath = true` + "\n"
	config += `	propagate_community = true` + "\n"
	config += `	ipv4_route_targets = [{` + "\n"
	config += `	  vpn_id = 1` + "\n"
	config += `	  export = [{` + "\n"
	config += `		asn_ip = "10:100"` + "\n"
	config += `	}]` + "\n"
	config += `	  import = [{` + "\n"
	config += `		asn_ip = "10:100"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_route_targets = [{` + "\n"
	config += `	  vpn_id = 1` + "\n"
	config += `	  export = [{` + "\n"
	config += `		asn_ip = "10:100"` + "\n"
	config += `	}]` + "\n"
	config += `	  import = [{` + "\n"
	config += `		asn_ip = "10:100"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	mpls_interfaces = [{` + "\n"
	config += `	  interface_name = "GigabitEthernet0"` + "\n"
	config += `	}]` + "\n"
	config += `	distance_external = 30` + "\n"
	config += `	distance_internal = 210` + "\n"
	config += `	distance_local = 30` + "\n"
	config += `	keepalive = 90` + "\n"
	config += `	holdtime = 220` + "\n"
	config += `	always_compare_med = true` + "\n"
	config += `	deterministic_med = true` + "\n"
	config += `	missing_med_worst = true` + "\n"
	config += `	compare_router_id = true` + "\n"
	config += `	multipath_relax = true` + "\n"
	config += `	address_families = [{` + "\n"
	config += `	  family_type = "ipv4-unicast"` + "\n"
	config += `	  ipv4_aggregate_addresses = [{` + "\n"
	config += `		prefix = "10.0.0.0/8"` + "\n"
	config += `		as_set_path = true` + "\n"
	config += `		summary_only = true` + "\n"
	config += `	}]` + "\n"
	config += `	  ipv4_networks = [{` + "\n"
	config += `		prefix = "10.2.2.0/24"` + "\n"
	config += `	}]` + "\n"
	config += `	  maximum_paths = 8` + "\n"
	config += `	  default_information_originate = true` + "\n"
	config += `	  table_map_policy = "MAP1"` + "\n"
	config += `	  table_map_filter = true` + "\n"
	config += `	  redistribute_routes = [{` + "\n"
	config += `		protocol = "ospf"` + "\n"
	config += `		route_policy = "POLICY1"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_neighbors = [{` + "\n"
	config += `	  address = "10.2.2.2"` + "\n"
	config += `	  description = "My neighbor"` + "\n"
	config += `	  shutdown = true` + "\n"
	config += `	  remote_as = "65001"` + "\n"
	config += `	  keepalive = 30` + "\n"
	config += `	  holdtime = 90` + "\n"
	config += `	  source_interface = "GigabitEthernet1"` + "\n"
	config += `	  next_hop_self = true` + "\n"
	config += `	  send_community = false` + "\n"
	config += `	  send_ext_community = false` + "\n"
	config += `	  ebgp_multihop = 10` + "\n"
	config += `	  password = "cisco123"` + "\n"
	config += `	  send_label = true` + "\n"
	config += `	  send_label_explicit = true` + "\n"
	config += `	  as_override = true` + "\n"
	config += `	  allow_as_in = 5` + "\n"
	config += `	  address_families = [{` + "\n"
	config += `		family_type = "ipv4-unicast"` + "\n"
	config += `		maximum_prefixes = 10000` + "\n"
	config += `		maximum_prefixes_threshold = 80` + "\n"
	config += `		maximum_prefixes_restart = 180` + "\n"
	config += `		maximum_prefixes_warning_only = true` + "\n"
	config += `      route_policies = [{` + "\n"
	config += `			direction = "in"` + "\n"
	config += `			policy_name = "POLICY1"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_neighbors = [{` + "\n"
	config += `	  address = "2001:1::1"` + "\n"
	config += `	  description = "My neighbor"` + "\n"
	config += `	  shutdown = true` + "\n"
	config += `	  remote_as = "65001"` + "\n"
	config += `	  keepalive = 30` + "\n"
	config += `	  holdtime = 90` + "\n"
	config += `	  source_interface = "GigabitEthernet1"` + "\n"
	config += `	  next_hop_self = true` + "\n"
	config += `	  send_community = false` + "\n"
	config += `	  send_ext_community = false` + "\n"
	config += `	  ebgp_multihop = 10` + "\n"
	config += `	  password = "cisco123"` + "\n"
	config += `	  send_label = true` + "\n"
	config += `	  send_label_explicit = true` + "\n"
	config += `	  as_override = true` + "\n"
	config += `	  allow_as_in = 5` + "\n"
	config += `	  address_families = [{` + "\n"
	config += `		family_type = "ipv6-unicast"` + "\n"
	config += `		maximum_prefixes = 10000` + "\n"
	config += `		maximum_prefixes_threshold = 80` + "\n"
	config += `		maximum_prefixes_restart = 180` + "\n"
	config += `		maximum_prefixes_warning_only = true` + "\n"
	config += `      route_policies = [{` + "\n"
	config += `			direction = "in"` + "\n"
	config += `			policy_name = "POLICY1"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_cisco_bgp_feature_template" "test" {
			id = sdwan_cisco_bgp_feature_template.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
