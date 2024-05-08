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
func TestAccDataSourceSdwanConfigurationGroupDevices(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_configuration_group_devices.test", "devices.0.id", "C8K-15411CCC-D476-0B3B-21F2-5D6AC387EE7B"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanConfigurationGroupDevicesPrerequisitesConfig + testAccDataSourceSdwanConfigurationGroupDevicesConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanConfigurationGroupDevicesPrerequisitesConfig = `
resource "sdwan_configuration_group" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
  solution    = "sdwan"
}
`

// End of section. //template:end testPrerequisites

func testAccDataSourceSdwanConfigurationGroupDevicesConfig() string {
	config := ""
	config += `resource "sdwan_configuration_group_devices" "test" {` + "\n"
	config += `	configuration_group_id = sdwan_configuration_group.test.id` + "\n"
	config += `	devices = [{` + "\n"
	config += `	  id = "C8K-15411CCC-D476-0B3B-21F2-5D6AC387EE7B"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_configuration_group_devices" "test" {
			configuration_group_id = sdwan_configuration_group.test.id
			depends_on = [sdwan_configuration_group_devices.test]
		}
	`
	return config
}
