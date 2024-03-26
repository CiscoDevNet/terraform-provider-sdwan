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

func TestAccSdwanCiscoOSPFv3FeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoOSPFv3FeatureTemplateConfig_minimum(),
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccSdwanCiscoOSPFv3FeatureTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_router_id", "1.2.3.4"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_auto_cost_reference_bandwidth", "100000"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_compatible_rfc1583", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_default_information_originate", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_default_information_originate_always", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_default_information_originate_metric", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_default_information_originate_metric_type", "type1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_distance_external", "111"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_distance_inter_area", "111"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_distance_intra_area", "112"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_timers_spf_delay", "300"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_timers_spf_initial_hold", "2000"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_timers_spf_max_hold", "20000"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_distance", "110"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_policy_name", "POLICY1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_filter", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_redistributes.0.protocol", "static"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_redistributes.0.route_policy", "RP1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_redistributes.0.nat_dia", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_max_metric_router_lsas.0.ad_type", "on-startup"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_max_metric_router_lsas.0.time", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.area_number", "1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.stub", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.stub_no_summary", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.nssa", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.nssa_no_summary", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.translate", "always"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.normal", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.interfaces.0.name", "e1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.interfaces.0.hello_interval", "20"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.interfaces.0.dead_interval", "60"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.interfaces.0.retransmit_interval", "10"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.interfaces.0.cost", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.interfaces.0.network", "point-to-point"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.interfaces.0.passive_interface", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.interfaces.0.authentication_type", "message-digest"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.interfaces.0.authentication_key", "authenticationKey"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.interfaces.0.ipsec_spi", "256"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.ranges.0.address", "1.1.1.0/24"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.ranges.0.cost", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv4_areas.0.ranges.0.no_advertise", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_router_id", "1.2.3.4"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_auto_cost_reference_bandwidth", "100000"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_compatible_rfc1583", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_default_information_originate", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_default_information_originate_always", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_default_information_originate_metric", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_default_information_originate_metric_type", "type1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_distance_external", "111"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_distance_inter_area", "111"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_distance_intra_area", "112"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_timers_spf_delay", "300"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_timers_spf_initial_hold", "2000"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_timers_spf_max_hold", "20000"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_distance", "110"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_policy_name", "POLICY2"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_filter", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_redistributes.0.protocol", "static"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_redistributes.0.route_policy", "RP1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_max_metric_router_lsas.0.ad_type", "on-startup"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_max_metric_router_lsas.0.time", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.area_number", "1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.stub", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.stub_no_summary", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.nssa", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.nssa_no_summary", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.translate", "always"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.normal", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.interfaces.0.name", "e1"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.interfaces.0.hello_interval", "20"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.interfaces.0.dead_interval", "60"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.interfaces.0.retransmit_interval", "10"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.interfaces.0.cost", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.interfaces.0.network", "point-to-point"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.interfaces.0.passive_interface", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.interfaces.0.authentication_type", "message-digest"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.interfaces.0.authentication_key", "authenticationKey"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.interfaces.0.ipsec_spi", "256"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.ranges.0.address", "2001::/48"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.ranges.0.cost", "100"),
					resource.TestCheckResourceAttr("sdwan_cisco_ospfv3_feature_template.test", "ipv6_areas.0.ranges.0.no_advertise", "true"),
				),
			},
		},
	})
}

func testAccSdwanCiscoOSPFv3FeatureTemplateConfig_minimum() string {
	return `
	resource "sdwan_cisco_ospfv3_feature_template" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
	}
	`
}

func testAccSdwanCiscoOSPFv3FeatureTemplateConfig_all() string {
	return `
	resource "sdwan_cisco_ospfv3_feature_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		ipv4_router_id = "1.2.3.4"
		ipv4_auto_cost_reference_bandwidth = 100000
		ipv4_compatible_rfc1583 = true
		ipv4_default_information_originate = true
		ipv4_default_information_originate_always = true
		ipv4_default_information_originate_metric = 100
		ipv4_default_information_originate_metric_type = "type1"
		ipv4_distance_external = 111
		ipv4_distance_inter_area = 111
		ipv4_distance_intra_area = 112
		ipv4_timers_spf_delay = 300
		ipv4_timers_spf_initial_hold = 2000
		ipv4_timers_spf_max_hold = 20000
		ipv4_distance = 110
		ipv4_policy_name = "POLICY1"
		ipv4_filter = false
		ipv4_redistributes = [{
			protocol = "static"
			route_policy = "RP1"
			nat_dia = true
		}]
		ipv4_max_metric_router_lsas = [{
			ad_type = "on-startup"
			time = 100
		}]
		ipv4_areas = [{
			area_number = 1
			stub = false
			stub_no_summary = false
			nssa = false
			nssa_no_summary = true
			translate = "always"
			normal = false
			interfaces = [{
				name = "e1"
				hello_interval = 20
				dead_interval = 60
				retransmit_interval = 10
				cost = 100
				network = "point-to-point"
				passive_interface = true
				authentication_type = "message-digest"
				authentication_key = "authenticationKey"
				ipsec_spi = 256
			}]
			ranges = [{
				address = "1.1.1.0/24"
				cost = 100
				no_advertise = true
			}]
		}]
		ipv6_router_id = "1.2.3.4"
		ipv6_auto_cost_reference_bandwidth = 100000
		ipv6_compatible_rfc1583 = true
		ipv6_default_information_originate = true
		ipv6_default_information_originate_always = true
		ipv6_default_information_originate_metric = 100
		ipv6_default_information_originate_metric_type = "type1"
		ipv6_distance_external = 111
		ipv6_distance_inter_area = 111
		ipv6_distance_intra_area = 112
		ipv6_timers_spf_delay = 300
		ipv6_timers_spf_initial_hold = 2000
		ipv6_timers_spf_max_hold = 20000
		ipv6_distance = 110
		ipv6_policy_name = "POLICY2"
		ipv6_filter = false
		ipv6_redistributes = [{
			protocol = "static"
			route_policy = "RP1"
		}]
		ipv6_max_metric_router_lsas = [{
			ad_type = "on-startup"
			time = 100
		}]
		ipv6_areas = [{
			area_number = 1
			stub = false
			stub_no_summary = false
			nssa = false
			nssa_no_summary = true
			translate = "always"
			normal = false
			interfaces = [{
				name = "e1"
				hello_interval = 20
				dead_interval = 60
				retransmit_interval = 10
				cost = 100
				network = "point-to-point"
				passive_interface = true
				authentication_type = "message-digest"
				authentication_key = "authenticationKey"
				ipsec_spi = 256
			}]
			ranges = [{
				address = "2001::/48"
				cost = 100
				no_advertise = true
			}]
		}]
	}
	`
}
