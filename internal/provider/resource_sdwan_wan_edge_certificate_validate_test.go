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
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Chassis numbers dedicated to the WAN edge certificate acceptance tests. These devices must be
// onboarded in the WAN edge list of the test Manager and must not be shared with any other test.
const (
	wanEdgeCertificateChassis1 = "C8K-3D1A8960-6E76-532C-DA93-50626FC5797E"
	wanEdgeCertificateChassis2 = "C8K-D4CE7174-5261-7E6F-91EA-4926BCF4C2DD"
)

func TestAccSdwanWANEdgeCertificateValidate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanWANEdgeCertificateConfig(wanEdgeCertificateChassis1, "staging"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_wan_edge_certificate_validate.test", "chassis_number", wanEdgeCertificateChassis1),
					resource.TestCheckResourceAttr("sdwan_wan_edge_certificate_validate.test", "id", wanEdgeCertificateChassis1),
					resource.TestCheckResourceAttr("sdwan_wan_edge_certificate_validate.test", "validity", "staging"),
					resource.TestCheckResourceAttrSet("sdwan_wan_edge_certificate_validate.test", "serial_number"),
					resource.TestCheckResourceAttr("sdwan_wan_edge_certificate_validate.test", "version", "1"),
				),
			},
			{
				Config: testAccSdwanWANEdgeCertificateConfig(wanEdgeCertificateChassis1, "valid"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_wan_edge_certificate_validate.test", "validity", "valid"),
				),
			},
			{
				Config: testAccSdwanWANEdgeCertificateConfig(wanEdgeCertificateChassis1, "invalid"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_wan_edge_certificate_validate.test", "validity", "invalid"),
				),
			},
			// Restore to valid; Delete is a no-op, so this is the real cleanup for the device.
			{
				Config: testAccSdwanWANEdgeCertificateConfig(wanEdgeCertificateChassis1, "valid"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_wan_edge_certificate_validate.test", "validity", "valid"),
				),
			},
			// Reapplying the same validity must produce an empty plan (skip-redundant-save).
			{
				Config:   testAccSdwanWANEdgeCertificateConfig(wanEdgeCertificateChassis1, "valid"),
				PlanOnly: true,
			},
		},
	})
}

func TestAccSdwanWANEdgeCertificateValidateInvalidValidity(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccSdwanWANEdgeCertificateConfig(wanEdgeCertificateChassis1, "bogus"),
				ExpectError: regexp.MustCompile(`value must be one of`),
			},
		},
	})
}

func TestAccSdwanWANEdgeCertificateValidateUnknownChassis(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccSdwanWANEdgeCertificateConfig("C8K-00000000-0000-0000-0000-000000000000", "staging"),
				ExpectError: regexp.MustCompile(`chassis number`),
			},
		},
	})
}

func testAccSdwanWANEdgeCertificateConfig(chassis, validity string) string {
	config := `resource "sdwan_wan_edge_certificate_validate" "test" {` + "\n"
	config += fmt.Sprintf(`	chassis_number = "%s"`, chassis) + "\n"
	config += fmt.Sprintf(`	validity = "%s"`, validity) + "\n"
	config += `}` + "\n"
	return config
}
