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
func TestAccDataSourceSdwanServiceLANVPNInterfaceEthernetFeatureAssociateTrackerGroupFeature(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanServiceLANVPNInterfaceEthernetFeatureAssociateTrackerGroupFeaturePrerequisitesConfig + testAccDataSourceSdwanServiceLANVPNInterfaceEthernetFeatureAssociateTrackerGroupFeatureConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanServiceLANVPNInterfaceEthernetFeatureAssociateTrackerGroupFeaturePrerequisitesConfig = `
resource "sdwan_service_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_service_lan_vpn_feature" "test" {
  name                       = "TF_TEST_SERVICE_LAN"
  description                = "Terraform test"
  feature_profile_id         = sdwan_service_feature_profile.test.id
}

resource "sdwan_service_lan_vpn_interface_ethernet_feature" "test" {
  name                       = "TF_TEST_SERVICE_LAN_INTERFACE_ETHERNET"
  description                = "Terraform test"
  feature_profile_id         = sdwan_service_feature_profile.test.id
  service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id
  interface_name             = "GigabitEthernet3"
  interface_description      = "LAN"
  ipv4_configuration_type    = "static"
  ipv4_address               = "1.2.3.4"
  ipv4_subnet_mask           = "0.0.0.0"
  ipv4_nat_type              = "pool"
}

resource "sdwan_service_tracker_feature" "test-1" {
  name                  = "TF_TEST_1"
  description           = "Terraform test"
  feature_profile_id    = sdwan_service_feature_profile.test.id
  tracker_name          = "TRACKER_2"
}

resource "sdwan_service_tracker_feature" "test-2" {
  name                  = "TF_TEST_2"
  description           = "Terraform test"
  feature_profile_id    = sdwan_service_feature_profile.test.id
  tracker_name          = "TRACKER_2"
}

resource "sdwan_service_tracker_group_feature" "test" {
  name               = "TF_TEST_TRACKER_GROUP"
  description        = "Terraform test"
  feature_profile_id = sdwan_service_feature_profile.test.id
  tracker_elements = [
    {
      tracker_id = sdwan_service_tracker_feature.test-1.id
    },
    {
      tracker_id = sdwan_service_tracker_feature.test-2.id
    }
  ]
  tracker_boolean = "or"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanServiceLANVPNInterfaceEthernetFeatureAssociateTrackerGroupFeatureConfig() string {
	config := ""
	config += `resource "sdwan_service_lan_vpn_interface_ethernet_feature_associate_tracker_group_feature" "test" {` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id` + "\n"
	config += `	service_lan_vpn_interface_ethernet_feature_id = sdwan_service_lan_vpn_interface_ethernet_feature.test.id` + "\n"
	config += `	service_tracker_group_feature_id = sdwan_service_tracker_group_feature.test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_service_lan_vpn_interface_ethernet_feature_associate_tracker_group_feature" "test" {
			feature_profile_id = sdwan_service_feature_profile.test.id
			service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id
			service_lan_vpn_interface_ethernet_feature_id = sdwan_service_lan_vpn_interface_ethernet_feature.test.id
			id = sdwan_service_lan_vpn_interface_ethernet_feature_associate_tracker_group_feature.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
