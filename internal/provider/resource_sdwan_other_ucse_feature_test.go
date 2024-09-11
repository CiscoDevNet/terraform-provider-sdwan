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
func TestAccSdwanOtherUCSEProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_other_ucse_feature.test", "bay", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_other_ucse_feature.test", "slot", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_other_ucse_feature.test", "access_port_dedicated", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_other_ucse_feature.test", "access_port_shared_type", "ge1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_other_ucse_feature.test", "access_port_shared_failover_type", "ge2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_other_ucse_feature.test", "vlan_id", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_other_ucse_feature.test", "assign_priority", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_other_ucse_feature.test", "interfaces.0.interface_name", "ucse2/0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_other_ucse_feature.test", "interfaces.0.ucse_interface_vpn", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_other_ucse_feature.test", "interfaces.0.ipv4_address", "10.1.15.15/24"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanOtherUCSEPrerequisitesProfileParcelConfig + testAccSdwanOtherUCSEProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanOtherUCSEPrerequisitesProfileParcelConfig + testAccSdwanOtherUCSEProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanOtherUCSEPrerequisitesProfileParcelConfig = `
resource "sdwan_other_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanOtherUCSEProfileParcelConfig_minimum() string {
	config := `resource "sdwan_other_ucse_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_other_feature_profile.test.id` + "\n"
	config += `	bay = 2` + "\n"
	config += `	slot = 0` + "\n"
	config += `	access_port_shared_type = "ge1"` + "\n"
	config += `	access_port_shared_failover_type = "ge2"` + "\n"
	config += `	ipv4_address = "2.2.2.2/24"` + "\n"
	config += `	default_gateway = "2.2.2.2"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanOtherUCSEProfileParcelConfig_all() string {
	config := `resource "sdwan_other_ucse_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_other_feature_profile.test.id` + "\n"
	config += `	bay = 2` + "\n"
	config += `	slot = 0` + "\n"
	config += `	access_port_dedicated = false` + "\n"
	config += `	access_port_shared_type = "ge1"` + "\n"
	config += `	access_port_shared_failover_type = "ge2"` + "\n"
	config += `	ipv4_address = "2.2.2.2/24"` + "\n"
	config += `	default_gateway = "2.2.2.2"` + "\n"
	config += `	vlan_id = 3` + "\n"
	config += `	assign_priority = 3` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `	  interface_name = "ucse2/0"` + "\n"
	config += `	  ucse_interface_vpn = 2` + "\n"
	config += `	  ipv4_address = "10.1.15.15/24"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
