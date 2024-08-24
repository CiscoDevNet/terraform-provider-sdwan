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
func TestAccDataSourceSdwanServiceRoutingOSPFv3IPv4ProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "router_id", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "distance", "110"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "distance_for_external_routes", "110"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "distance_for_inter_area_routes", "110"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "distance_for_intra_area_routes", "110"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "reference_bandwidth", "101"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "rfc_1583_compatible", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "originate", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "always", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "metric", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "metric_type", "type1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "spf_calculation_deplay", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "initial_hold_time", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "maximum_hold_time", "10000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "filter", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "redistributes.0.protocol", "nat-route"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "redistributes.0.enable_nat_dia", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "router_lsa_action", "on-startup"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "router_lsa_on_startu_p_time", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "areas.0.area_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "areas.0.area_type", "stub"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "areas.0.interfaces.0.if_name", "GigabitEthernet2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "areas.0.interfaces.0.hello_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "areas.0.interfaces.0.dead_interval", "40"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "areas.0.interfaces.0.lsa_retransmit_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "areas.0.interfaces.0.interface_cost", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "areas.0.interfaces.0.ospf_network_type", "broadcast"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "areas.0.interfaces.0.passive_interface", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "areas.0.interfaces.0.auth_type", "no-auth"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "areas.0.ranges.0.address", "10.1.1.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "areas.0.ranges.0.mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "areas.0.ranges.0.cost", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_routing_ospfv3_ipv4_profile_parcel.test", "areas.0.ranges.0.no_advertise", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanServiceRoutingOSPFv3IPv4PrerequisitesProfileParcelConfig + testAccDataSourceSdwanServiceRoutingOSPFv3IPv4ProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanServiceRoutingOSPFv3IPv4PrerequisitesProfileParcelConfig = `
resource "sdwan_service_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanServiceRoutingOSPFv3IPv4ProfileParcelConfig() string {
	config := `resource "sdwan_service_routing_ospfv3_ipv4_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	router_id = "1.2.3.4"` + "\n"
	config += `	distance = 110` + "\n"
	config += `	distance_for_external_routes = 110` + "\n"
	config += `	distance_for_inter_area_routes = 110` + "\n"
	config += `	distance_for_intra_area_routes = 110` + "\n"
	config += `	reference_bandwidth = 101` + "\n"
	config += `	rfc_1583_compatible = true` + "\n"
	config += `	originate = false` + "\n"
	config += `	always = false` + "\n"
	config += `	metric = 1` + "\n"
	config += `	metric_type = "type1"` + "\n"
	config += `	spf_calculation_deplay = 200` + "\n"
	config += `	initial_hold_time = 1000` + "\n"
	config += `	maximum_hold_time = 10000` + "\n"
	config += `	filter = false` + "\n"
	config += `	redistributes = [{` + "\n"
	config += `	  protocol = "nat-route"` + "\n"
	config += `	  enable_nat_dia = true` + "\n"
	config += `	}]` + "\n"
	config += `	router_lsa_action = "on-startup"` + "\n"
	config += `	router_lsa_on_startu_p_time = 30` + "\n"
	config += `	areas = [{` + "\n"
	config += `	  area_number = 1` + "\n"
	config += `	  area_type = "stub"` + "\n"
	config += `	  interfaces = [{` + "\n"
	config += `		if_name = "GigabitEthernet2"` + "\n"
	config += `		hello_interval = 10` + "\n"
	config += `		dead_interval = 40` + "\n"
	config += `		lsa_retransmit_interval = 5` + "\n"
	config += `		interface_cost = 10` + "\n"
	config += `		ospf_network_type = "broadcast"` + "\n"
	config += `		passive_interface = false` + "\n"
	config += `		auth_type = "no-auth"` + "\n"
	config += `	}]` + "\n"
	config += `	  ranges = [{` + "\n"
	config += `		address = "10.1.1.0"` + "\n"
	config += `		mask = "255.255.255.0"` + "\n"
	config += `		cost = 1` + "\n"
	config += `		no_advertise = false` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_service_routing_ospfv3_ipv4_profile_parcel" "test" {
			id = sdwan_service_routing_ospfv3_ipv4_profile_parcel.test.id
			feature_profile_id = sdwan_service_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
