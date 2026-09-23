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
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// The device has to be onboarded in the WAN edge list of the Manager already.
func testAccWANEdgeCertificateChassis(t *testing.T) string {
	chassis := os.Getenv("SDWAN_TEST_CHASSIS_NUMBER")
	if chassis == "" {
		t.Skip("skipping test, set environment variable SDWAN_TEST_CHASSIS_NUMBER to the chassis number of a WAN edge device")
	}
	return chassis
}

func TestAccSdwanWANEdgeCertificateValidate(t *testing.T) {
	chassis := testAccWANEdgeCertificateChassis(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanWANEdgeCertificateConfig(chassis, "staging"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_wan_edge_certificate_validate.test", "chassis_number", chassis),
					resource.TestCheckResourceAttr("sdwan_wan_edge_certificate_validate.test", "id", chassis),
					resource.TestCheckResourceAttr("sdwan_wan_edge_certificate_validate.test", "validity", "staging"),
					resource.TestCheckResourceAttrSet("sdwan_wan_edge_certificate_validate.test", "serial_number"),
					resource.TestCheckResourceAttr("sdwan_wan_edge_certificate_validate.test", "version", "1"),
				),
			},
			{
				Config: testAccSdwanWANEdgeCertificateConfig(chassis, "valid"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_wan_edge_certificate_validate.test", "validity", "valid"),
				),
			},
			{
				Config: testAccSdwanWANEdgeCertificateConfig(chassis, "invalid"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_wan_edge_certificate_validate.test", "validity", "invalid"),
				),
			},
			{
				ResourceName:      "sdwan_wan_edge_certificate_validate.test",
				ImportState:       true,
				ImportStateId:     chassis,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccSdwanWANEdgeCertificateValidateInvalidValidity(t *testing.T) {
	chassis := testAccWANEdgeCertificateChassis(t)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccSdwanWANEdgeCertificateConfig(chassis, "bogus"),
				ExpectError: regexp.MustCompile(`value must be one of`),
			},
		},
	})
}

func TestAccSdwanWANEdgeCertificateValidateUnknownChassis(t *testing.T) {
	// The helper is used for its environment guard; this test intentionally uses a fixed unknown chassis.
	testAccWANEdgeCertificateChassis(t)

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
