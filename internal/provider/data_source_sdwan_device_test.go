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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceSdwanDevice(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanDeviceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_device.test", "devices.0.device_id", "100.0.0.101"),
					resource.TestCheckResourceAttr("data.sdwan_device.test", "devices.0.uuid", "a6e71b79-1cf1-4387-b90d-1835d4501c7f"),
					resource.TestCheckResourceAttr("data.sdwan_device.test", "devices.0.site_id", "100"),
					resource.TestCheckResourceAttr("data.sdwan_device.test", "devices.0.serial_number", "E6486A92D4404888BC795DB045AF86D1"),
					resource.TestCheckResourceAttr("data.sdwan_device.test", "devices.0.hostname", "Controller01"),
					resource.TestCheckResourceAttr("data.sdwan_device.test", "devices.0.reachability", "reachable"),
					resource.TestCheckResourceAttr("data.sdwan_device.test", "devices.0.status", "normal"),
					resource.TestCheckResourceAttr("data.sdwan_device.test", "devices.0.state", "green"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanDeviceConfig = `


data "sdwan_device" "test" {
  serial_number = "E6486A92D4404888BC795DB045AF86D1"
  name = "Controller01"
}
`
