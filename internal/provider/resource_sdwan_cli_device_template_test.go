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

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanCLIDeviceTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCLIDeviceTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cli_device_template.test", "name", "TF_TEST_ALL"),
					resource.TestCheckResourceAttr("sdwan_cli_device_template.test", "description", "Terraform integration test"),
					resource.TestCheckResourceAttr("sdwan_cli_device_template.test", "device_type", "vedge-ISR-4331"),
					resource.TestCheckResourceAttr("sdwan_cli_device_template.test", "cli_type", "device"),
					resource.TestCheckResourceAttr("sdwan_cli_device_template.test", "cli_configuration", " system\n host-name             {{test}}-ISR4331-1200-1"),
				),
			},
		},
	})
}

func testAccSdwanCLIDeviceTemplateConfig_all() string {
	return `
	resource "sdwan_cli_device_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_type = "vedge-ISR-4331"
		cli_type = "device"
		cli_configuration = " system\n host-name             {{test}}-ISR4331-1200-1"
	}
	`
}