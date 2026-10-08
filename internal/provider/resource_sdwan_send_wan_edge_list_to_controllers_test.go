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
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanSendWANEdgeListToControllers(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanSendWANEdgeListConfig("staging", "invalid"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("sdwan_send_wan_edge_list_to_controllers.test", "id"),
					resource.TestCheckResourceAttr("sdwan_send_wan_edge_list_to_controllers.test", "synced", "true"),
				),
			},
			{
				Config: testAccSdwanSendWANEdgeListConfig("invalid", "staging"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_send_wan_edge_list_to_controllers.test", "synced", "true"),
				),
			},
			// Restore both devices to valid; this is the real cleanup, not terraform destroy.
			{
				Config: testAccSdwanSendWANEdgeListConfig("valid", "valid"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_send_wan_edge_list_to_controllers.test", "synced", "true"),
				),
			},
			// Reapplying the same state must produce an empty plan (no dirty list, no push).
			{
				Config:   testAccSdwanSendWANEdgeListConfig("valid", "valid"),
				PlanOnly: true,
			},
		},
	})
}

func testAccSdwanSendWANEdgeListConfig(validity1, validity2 string) string {
	config := `resource "sdwan_wan_edge_certificate_validate" "test1" {` + "\n"
	config += fmt.Sprintf(`	chassis_number = "%s"`, wanEdgeCertificateChassis1) + "\n"
	config += fmt.Sprintf(`	validity = "%s"`, validity1) + "\n"
	config += `}` + "\n"
	config += `resource "sdwan_wan_edge_certificate_validate" "test2" {` + "\n"
	config += fmt.Sprintf(`	chassis_number = "%s"`, wanEdgeCertificateChassis2) + "\n"
	config += fmt.Sprintf(`	validity = "%s"`, validity2) + "\n"
	config += `}` + "\n"
	config += `resource "sdwan_send_wan_edge_list_to_controllers" "test" {` + "\n"
	config += `	version = sdwan_wan_edge_certificate_validate.test1.version + sdwan_wan_edge_certificate_validate.test2.version` + "\n"
	config += `}` + "\n"
	return config
}
