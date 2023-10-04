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

func TestAccSdwanDevice(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanDeviceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_device.test", "serial_number", "2AJC9DJ"),
					resource.TestCheckResourceAttr("sdwan_device.test", "devices.0.device_id", "2081c2f4-3f9f-4fee-8078-dcc8904e368d"),
					resource.TestCheckResourceAttr("sdwan_device.test", "devices.0.uuid", "2081c2f4-3f9f-4fee-8078-dcc8904e368d"),
					resource.TestCheckResourceAttr("sdwan_device.test", "devices.0.site_id", "400"),
					resource.TestCheckResourceAttr("sdwan_device.test", "devices.0.serial_number", "2AJC9DJ"),
					resource.TestCheckResourceAttr("sdwan_device.test", "devices.0.hostname", "vEdge-123"),
					resource.TestCheckResourceAttr("sdwan_device.test", "devices.0.reachability", "reachable"),
					resource.TestCheckResourceAttr("sdwan_device.test", "devices.0.status", "normal"),
					resource.TestCheckResourceAttr("sdwan_device.test", "devices.0.state", "green"),
				),
			},
		},
	})
}

const testAccSdwanDeviceConfig = `


resource "sdwan_device" "test" {
	serial_number = "2AJC9DJ"
	devices = [{
		device_id = "2081c2f4-3f9f-4fee-8078-dcc8904e368d"
		uuid = "2081c2f4-3f9f-4fee-8078-dcc8904e368d"
		site_id = "400"
		serial_number = "2AJC9DJ"
		hostname = "vEdge-123"
		reachability = "reachable"
		status = "normal"
		state = "green"
	}]
}
`