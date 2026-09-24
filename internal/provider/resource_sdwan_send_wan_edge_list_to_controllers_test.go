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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanSendWANEdgeListToControllers(t *testing.T) {
	if os.Getenv("SDWAN_TEST_CHASSIS_NUMBER") == "" {
		t.Skip("skipping test, set environment variable SDWAN_TEST_CHASSIS_NUMBER to enable certificate tests (this test pushes the WAN edge list to the controllers)")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `resource "sdwan_send_wan_edge_list_to_controllers" "test" {
	version = 1
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("sdwan_send_wan_edge_list_to_controllers.test", "id"),
					resource.TestCheckResourceAttr("sdwan_send_wan_edge_list_to_controllers.test", "version", "1"),
				),
			},
			{
				Config: `resource "sdwan_send_wan_edge_list_to_controllers" "test" {
	version = 2
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_send_wan_edge_list_to_controllers.test", "version", "2"),
				),
			},
		},
	})
}
