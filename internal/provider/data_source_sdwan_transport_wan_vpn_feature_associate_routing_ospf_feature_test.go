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
func TestAccDataSourceSdwanTransportWANVPNFeatureAssociateRoutingOSPFFeature(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanTransportWANVPNFeatureAssociateRoutingOSPFFeaturePrerequisitesConfig + testAccDataSourceSdwanTransportWANVPNFeatureAssociateRoutingOSPFFeatureConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanTransportWANVPNFeatureAssociateRoutingOSPFFeaturePrerequisitesConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_transport_wan_vpn_feature" "test" {
  name               = "TF_TEST_WAN_VPN"
  description        = "Terraform test"
  feature_profile_id = sdwan_transport_feature_profile.test.id
  vpn                = 0
}

resource "sdwan_transport_routing_ospf_feature" "test" {
  name                                      = "TF_TEST_OSPF"
  description                               = "Terraform test"
  feature_profile_id                        = sdwan_transport_feature_profile.test.id
  router_id                                 = "1.2.3.4"
  reference_bandwidth                       = 101
  rfc_1583_compatible                       = true
  default_information_originate             = false
  default_information_originate_always      = false
  default_information_originate_metric      = 1
  default_information_originate_metric_type = "type1"
  distance_external                         = 110
  distance_inter_area                       = 110
  distance_intra_area                       = 110
  spf_calculation_delay                     = 200
  spf_initial_hold_time                     = 1000
  spf_maximum_hold_time                     = 10000
  redistributes = [
    {
      protocol = "static"
      nat_dia  = true
    }
  ]
  router_lsas = [
    {
      type = "on-startup"
      time = 5
    }
  ]
  areas = [
    {
      area_number = 1
      area_type   = "stub"
      no_summary  = false
      interfaces = [
        {
          name                       = "GigabitEthernet2"
          hello_interval             = 10
          dead_interval              = 40
          lsa_retransmit_interval    = 5
          cost                       = 10
          designated_router_priority = 1
          network_type               = "broadcast"
          passive_interface          = false
          authentication_type        = "message-digest"
          message_digest_key_id      = 7
          message_digest_key         = "sdjfhsghbjdjr"
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanTransportWANVPNFeatureAssociateRoutingOSPFFeatureConfig() string {
	config := ""
	config += `resource "sdwan_transport_wan_vpn_feature_associate_routing_ospf_feature" "test" {` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id` + "\n"
	config += `	transport_routing_ospf_feature_id = sdwan_transport_routing_ospf_feature.test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_transport_wan_vpn_feature_associate_routing_ospf_feature" "test" {
			feature_profile_id = sdwan_transport_feature_profile.test.id
			transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id
			id = sdwan_transport_wan_vpn_feature_associate_routing_ospf_feature.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig