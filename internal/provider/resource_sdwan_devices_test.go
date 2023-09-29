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

func TestAccSdwanDevices(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanDevicesConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_devices.test", "device_id", ""),
					resource.TestCheckResourceAttr("sdwan_devices.test", "uuid", ""),
					resource.TestCheckResourceAttr("sdwan_devices.test", "site_id", "400"),
					resource.TestCheckResourceAttr("sdwan_devices.test", "serial_number", ""),
					resource.TestCheckResourceAttr("sdwan_devices.test", "hostname", ""),
					resource.TestCheckResourceAttr("sdwan_devices.test", "reachability", "reachable"),
					resource.TestCheckResourceAttr("sdwan_devices.test", "status", "normal"),
					resource.TestCheckResourceAttr("sdwan_devices.test", "state", "green"),
				),
			},
		},
	})
}

const testAccSdwanDevicesConfig = `


resource "sdwan_devices" "test" {
	device_id = ""
	uuid = ""
	site_id = "400"
	serial_number = ""
	hostname = ""
	reachability = "reachable"
	status = "normal"
	state = "green"
}
`