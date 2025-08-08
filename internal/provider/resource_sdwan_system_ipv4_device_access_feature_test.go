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
func TestAccSdwanSystemIPv4DeviceAccessProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_ipv4_device_access_feature.test", "sequences.0.id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_ipv4_device_access_feature.test", "sequences.0.name", "SEQ_1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_ipv4_device_access_feature.test", "sequences.0.base_action", "accept"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_system_ipv4_device_access_feature.test", "sequences.0.device_access_port", "161"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanSystemIPv4DeviceAccessPrerequisitesProfileParcelConfig + testAccSdwanSystemIPv4DeviceAccessProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanSystemIPv4DeviceAccessPrerequisitesProfileParcelConfig + testAccSdwanSystemIPv4DeviceAccessProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanSystemIPv4DeviceAccessPrerequisitesProfileParcelConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanSystemIPv4DeviceAccessProfileParcelConfig_minimum() string {
	config := `resource "sdwan_system_ipv4_device_access_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanSystemIPv4DeviceAccessProfileParcelConfig_all() string {
	config := `resource "sdwan_system_ipv4_device_access_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `	sequences = [{` + "\n"
	config += `	  id = 1` + "\n"
	config += `	  name = "SEQ_1"` + "\n"
	config += `	  base_action = "accept"` + "\n"
	config += `	  device_access_port = 161` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
