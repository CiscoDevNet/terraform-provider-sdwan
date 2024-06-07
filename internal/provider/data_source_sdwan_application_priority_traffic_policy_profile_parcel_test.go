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
func TestAccDataSourceSdwanApplicationPriorityTrafficPolicyProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "default_action", "accept"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "simple_flow", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "target_direction", "all"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.sequence_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.name", "RULE_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.base_action", "accept"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.protocol", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.matches.0.dscp", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.matches.0.packet_length", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.matches.0.tcp", "gre"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.matches.0.traffic_to", "core"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.counter", "COUNTER_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.log", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sla_class.0.strict_drop", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sla_class.0.fallback_to_best_path", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.dscp", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.local_tloc_restrict", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.local_tloc_list_encapsulation", "gre"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.tloc_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.tloc_encapsulation", "gre"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.service_type", "FW"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.service_encapsulation", "ipsec"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.service_tloc_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.service_vpn", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.service_chain_type", "SC1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.service_chain_vpn", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.service_chain_local", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.service_chain_fallback_to_routing", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.service_chain_encapsulation", "ipsec"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.service_chain_id", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.next_hop", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.next_hop_ipv6", "2001:0:0:1::/64"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.sets.0.vpn", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.redirect_dns_field", "redirectDns"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.redirect_dns_value", "umbrella"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.tcp_optimization", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.dre_optimization", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.service_node_group", "SNG-APPQOE1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.loss_correction_type", "fecAdaptive"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.loss_correct_fec_threshold", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.cflowd", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.nat_pool", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.nat_vpn", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.nat_fallback", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.nat_bypass", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.secure_internet_gateway", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.fallback_to_routing", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_application_priority_traffic_policy_profile_parcel.test", "sequences.0.actions.0.secure_service_edge_instance", "zScaler"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanApplicationPriorityTrafficPolicyPrerequisitesProfileParcelConfig + testAccDataSourceSdwanApplicationPriorityTrafficPolicyProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanApplicationPriorityTrafficPolicyPrerequisitesProfileParcelConfig = `
resource "sdwan_application_priority_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanApplicationPriorityTrafficPolicyProfileParcelConfig() string {
	config := `resource "sdwan_application_priority_traffic_policy_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_application_priority_feature_profile.test.id` + "\n"
	config += `	default_action = "accept"` + "\n"
	config += `	simple_flow = false` + "\n"
	config += `	vpn = ["1"]` + "\n"
	config += `	target_direction = "all"` + "\n"
	config += `	sequences = [{` + "\n"
	config += `	  sequence_id = 1` + "\n"
	config += `	  name = "RULE_1"` + "\n"
	config += `	  base_action = "accept"` + "\n"
	config += `	  protocol = "ipv4"` + "\n"
	config += `	  matches = [{` + "\n"
	config += `		dscp = 1` + "\n"
	config += `		packet_length = "123"` + "\n"
	config += `		protocol = ["2"]` + "\n"
	config += `		tcp = "gre"` + "\n"
	config += `		traffic_to = "core"` + "\n"
	config += `	}]` + "\n"
	config += `	  actions = [{` + "\n"
	config += `		counter = "COUNTER_1"` + "\n"
	config += `		log = false` + "\n"
	config += `      sla_class = [{` + "\n"
	config += `			preferred_color = ["default"]` + "\n"
	config += `			strict_drop = true` + "\n"
	config += `			fallback_to_best_path = false` + "\n"
	config += `		}]` + "\n"
	config += `		backup_sla_preferred_color = ["default"]` + "\n"
	config += `      sets = [{` + "\n"
	config += `			dscp = 1` + "\n"
	config += `			local_tloc_list_color = ["default"]` + "\n"
	config += `			local_tloc_restrict = "false"` + "\n"
	config += `			local_tloc_list_encapsulation = "gre"` + "\n"
	config += `			tloc_ip = "1.2.3.4"` + "\n"
	config += `			tloc_color = ["default"]` + "\n"
	config += `			tloc_encapsulation = "gre"` + "\n"
	config += `			service_type = "FW"` + "\n"
	config += `			service_color = ["default"]` + "\n"
	config += `			service_encapsulation = "ipsec"` + "\n"
	config += `			service_tloc_ip = "1.2.3.4"` + "\n"
	config += `			service_vpn = "1"` + "\n"
	config += `			service_chain_type = "SC1"` + "\n"
	config += `			service_chain_vpn = 1` + "\n"
	config += `			service_chain_local = false` + "\n"
	config += `			service_chain_fallback_to_routing = false` + "\n"
	config += `			service_chain_tloc = ["default"]` + "\n"
	config += `			service_chain_encapsulation = "ipsec"` + "\n"
	config += `			service_chain_id = "1.2.3.4"` + "\n"
	config += `			next_hop = "1.2.3.4"` + "\n"
	config += `			next_hop_ipv6 = "2001:0:0:1::/64"` + "\n"
	config += `			vpn = "1"` + "\n"
	config += `		}]` + "\n"
	config += `		redirect_dns_field = "redirectDns"` + "\n"
	config += `		redirect_dns_value = "umbrella"` + "\n"
	config += `		tcp_optimization = true` + "\n"
	config += `		dre_optimization = true` + "\n"
	config += `		service_node_group = "SNG-APPQOE1"` + "\n"
	config += `		loss_correction_type = "fecAdaptive"` + "\n"
	config += `		loss_correct_fec_threshold = 1` + "\n"
	config += `		cflowd = true` + "\n"
	config += `		nat_pool = 2` + "\n"
	config += `		nat_vpn = 0` + "\n"
	config += `		nat_fallback = false` + "\n"
	config += `		nat_bypass = false` + "\n"
	config += `		nat_dia_pool = [1]` + "\n"
	config += `		nat_dia_interface = ["ethernet"]` + "\n"
	config += `		secure_internet_gateway = true` + "\n"
	config += `		fallback_to_routing = true` + "\n"
	config += `		secure_service_edge_instance = "zScaler"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_application_priority_traffic_policy_profile_parcel" "test" {
			id = sdwan_application_priority_traffic_policy_profile_parcel.test.id
			feature_profile_id = sdwan_application_priority_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
