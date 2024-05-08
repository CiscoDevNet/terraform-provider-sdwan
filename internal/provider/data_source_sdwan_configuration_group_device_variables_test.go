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

func TestAccDataSourceSdwanConfigurationGroupDeviceVariables(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_configuration_group_device_variables.test", "solution", "sdwan"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_configuration_group_device_variables.test", "devices.0.variables.0.name", "system_ip"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanConfigurationGroupDeviceVariablesPrerequisitesConfig + testAccDataSourceSdwanConfigurationGroupDeviceVariablesConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanConfigurationGroupDeviceVariablesPrerequisitesConfig = `
resource "sdwan_system_feature_profile" "test" {
  name        = "SYSTEM_TF"
  description = "My system feature profile 1"
}

resource "sdwan_configuration_group" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
  solution    = "sdwan"
  feature_profiles = [
    {
      id = sdwan_system_feature_profile.test.id
    }
  ]
}

resource "sdwan_configuration_group_devices" "test" {
  configuration_group_id = sdwan_configuration_group.test.id
  devices = [{
    id = "C8K-15411CCC-D476-0B3B-21F2-5D6AC387EE7B"
  }]
}
`

// End of section. //template:end testPrerequisites

func testAccDataSourceSdwanConfigurationGroupDeviceVariablesConfig() string {
	config := ""
	config += `resource "sdwan_configuration_group_device_variables" "test" {` + "\n"
	config += `	 configuration_group_id = sdwan_configuration_group.test.id` + "\n"
	config += `	 solution = "sdwan"` + "\n"
	config += `	 devices = [{` + "\n"
	config += `	   device_id = sdwan_configuration_group_devices.test.devices.0.id` + "\n"
	config += `	   variables = [` + "\n"
	config += `	     {` + "\n"
	config += `	       name = "host_name"` + "\n"
	config += `	       value = "edge1"` + "\n"
	config += `	   	 },` + "\n"
	config += `	     {` + "\n"
	config += `	       name = "pseudo_commit_timer"` + "\n"
	config += `	       value = "0"` + "\n"
	config += `		 },` + "\n"
	config += `	     {` + "\n"
	config += `	       name = "site_id"` + "\n"
	config += `	       value = "1"` + "\n"
	config += `		 },` + "\n"
	config += `	     {` + "\n"
	config += `	       name = "system_ip"` + "\n"
	config += `	       value = "10.1.1.1"` + "\n"
	config += `	     },` + "\n"
	config += `	   ]` + "\n"
	config += `	 }]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_configuration_group_device_variables" "test" {
			configuration_group_id = sdwan_configuration_group.test.id
		}
	`
	return config
}
