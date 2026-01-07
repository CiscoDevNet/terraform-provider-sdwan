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
func TestAccDataSourceSdwanServiceLANVPNInterfaceIPSecFeatureAssociateDHCPServerFeature(t *testing.T) {
	if os.Getenv("SDWAN_2015_IN_PROGRESS") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015_IN_PROGRESS")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanServiceLANVPNInterfaceIPSecFeatureAssociateDHCPServerFeaturePrerequisitesConfig + testAccDataSourceSdwanServiceLANVPNInterfaceIPSecFeatureAssociateDHCPServerFeatureConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanServiceLANVPNInterfaceIPSecFeatureAssociateDHCPServerFeaturePrerequisitesConfig = `
resource "sdwan_service_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_service_lan_vpn_feature" "test" {
  name                       = "TF_TEST_SERVICE_LAN"
  description                = "Terraform test"
  feature_profile_id         = sdwan_service_feature_profile.test.id
}

resource "sdwan_service_lan_vpn_interface_ipsec_feature" "test" {
  name                       = "TF_TEST_SERVICE_LAN_INTERFACE_IPSEC"
  description                = "Terraform test"
  feature_profile_id         = sdwan_service_feature_profile.test.id
  service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id
  interface_name                      = "ipsec987"
  ipv4_address                        = "9.7.5.4"
  ipv4_subnet_mask                    = "255.255.255.0"
  tunnel_source_ipv4_address          = "1.3.5.88"
  tunnel_source_ipv4_subnet_mask      = "255.255.255.0"
  tunnel_source_interface             = "GigabitEthernet8"
  tunnel_destination_ipv4_address     = "2.55.67.99"
  tunnel_destination_ipv4_subnet_mask = "255.255.255.0"
  application_tunnel_type             = "none"
  ike_preshared_key                   = "123"
}

resource "sdwan_service_dhcp_server_feature" "test" {
  name               = "TF_TEST_DHCP_SERVER"
  description        = "Terraform Test"
  feature_profile_id = sdwan_service_feature_profile.test.id
  network_address    = "1.2.3.4"
  subnet_mask        = "255.255.255.0"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanServiceLANVPNInterfaceIPSecFeatureAssociateDHCPServerFeatureConfig() string {
	config := ""
	config += `resource "sdwan_service_lan_vpn_interface_ipsec_feature_associate_dhcp_server_feature" "test" {` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id` + "\n"
	config += `	service_lan_vpn_interface_ipsec_feature_id = sdwan_service_lan_vpn_interface_ipsec_feature.test.id` + "\n"
	config += `	service_dhcp_server_feature_id = sdwan_service_dhcp_server_feature.test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_service_lan_vpn_interface_ipsec_feature_associate_dhcp_server_feature" "test" {
			feature_profile_id = sdwan_service_feature_profile.test.id
			service_lan_vpn_feature_id = sdwan_service_lan_vpn_feature.test.id
			service_lan_vpn_interface_ipsec_feature_id = sdwan_service_lan_vpn_interface_ipsec_feature.test.id
			id = sdwan_service_lan_vpn_interface_ipsec_feature_associate_dhcp_server_feature.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
