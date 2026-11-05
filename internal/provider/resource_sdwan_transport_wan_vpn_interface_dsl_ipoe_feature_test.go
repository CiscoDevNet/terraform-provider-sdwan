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
func TestAccSdwanTransportWANVPNInterfaceDSLIPoEProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "vdsl_slot", "0/0/0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "interface_name", "GigabitEthernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "ipv4_address", "192.0.2.10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "ipv4_subnet_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "tunnel_interface_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "tunnel_interface_color", "default"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTransportWANVPNInterfaceDSLIPoEPrerequisitesProfileParcelConfig + testAccSdwanTransportWANVPNInterfaceDSLIPoEProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanTransportWANVPNInterfaceDSLIPoEPrerequisitesProfileParcelConfig + testAccSdwanTransportWANVPNInterfaceDSLIPoEProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanTransportWANVPNInterfaceDSLIPoEPrerequisitesProfileParcelConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_transport_wan_vpn_feature" "test" {
  name               = "TF_TEST_WAN"
  description        = "Terraform test"
  feature_profile_id = sdwan_transport_feature_profile.test.id
  vpn                = 0
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanTransportWANVPNInterfaceDSLIPoEProfileParcelConfig_minimum() string {
	config := `resource "sdwan_transport_wan_vpn_interface_dsl_ipoe_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id` + "\n"
	config += `	vdsl_slot = "0/0/0"` + "\n"
	config += `	interface_name = "GigabitEthernet1"` + "\n"
	config += `	ipv4_address_type = "dynamic"` + "\n"
	config += `	tunnel_interface_enabled = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTransportWANVPNInterfaceDSLIPoEProfileParcelConfig_all() string {
	config := `resource "sdwan_transport_wan_vpn_interface_dsl_ipoe_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id` + "\n"
	config += `	shutdown = true` + "\n"
	config += `	vdsl_slot = "0/0/0"` + "\n"
	config += `	interface_name = "GigabitEthernet1"` + "\n"
	config += `	ipv4_address_type = "static"` + "\n"
	config += `	ipv4_address = "192.0.2.10"` + "\n"
	config += `	ipv4_subnet_mask = "255.255.255.0"` + "\n"
	config += `	tunnel_interface_enabled = true` + "\n"
	config += `	tunnel_interface_color = "default"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
