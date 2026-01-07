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
func TestAccSdwanServiceLANVPNFeatureAssociateRoutingOSPFv3IPv4Feature(t *testing.T) {
	if os.Getenv("SDWAN_2015_IN_PROGRESS") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015_IN_PROGRESS")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanServiceLANVPNFeatureAssociateRoutingOSPFv3IPv4FeaturePrerequisitesConfig + testAccSdwanServiceLANVPNFeatureAssociateRoutingOSPFv3IPv4FeatureConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanServiceLANVPNFeatureAssociateRoutingOSPFv3IPv4FeaturePrerequisitesConfig = `
resource "sdwan_service_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_service_lan_vpn_feature" "test" {
  name                       = "TF_TEST_SERVICE_LAN"
  description                = "Terraform test"
  feature_profile_id         = sdwan_service_feature_profile.test.id
}

resource "sdwan_service_routing_ospfv3_ipv4_feature" "test" {
  name                                      = "TF_TEST_ROUTING_OSPFV3_IPV4"
  description                               = "Terraform test"
  feature_profile_id                        = sdwan_service_feature_profile.test.id
  router_id                                 = "1.2.3.4"
  distance                                  = 110
  distance_external                         = 110
  distance_inter_area                       = 110
  distance_intra_area                       = 110
  reference_bandwidth                       = 101
  rfc_1583_compatible                       = true
  default_information_originate             = false
  default_information_originate_always      = false
  default_information_originate_metric      = 1
  default_information_originate_metric_type = "type1"
  spf_calculation_delay                     = 200
  spf_initial_hold_time                     = 1000
  spf_maximum_hold_time                     = 10000
  filter                                    = false
  redistributes = [
    {
      protocol = "nat-route"
      nat_dia  = true
    }
  ]
  router_lsa_action          = "on-startup"
  router_lsa_on_startup_time = 30
  areas = [
    {
      area_number = 1
      area_type   = "stub"
      interfaces = [
        {
          name                    = "GigabitEthernet2"
          hello_interval          = 10
          dead_interval           = 40
          lsa_retransmit_interval = 5
          cost                    = 10
          network_type            = "broadcast"
          passive_interface       = false
          authentication_type     = "no-auth"
        }
      ]
      ranges = [
        {
          ip_address   = "10.1.1.0"
          subnet_mask  = "255.255.255.0"
          cost         = 1
          no_advertise = false
        }
      ]
    }
  ]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanServiceLANVPNFeatureAssociateRoutingOSPFv3IPv4FeatureConfig_all() string {
	config := `resource "sdwan_service_lan_vpn_feature_associate_routing_ospfv3_ipv4_feature" "test" {` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id` + "\n"
	config += `	service_routing_ospfv3_ipv4_feature_id = sdwan_service_routing_ospfv3_ipv4_feature.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
