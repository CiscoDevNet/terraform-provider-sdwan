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

func TestAccSdwanCiscoBFDFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoBFDFeatureTemplateConfig_minimum(),
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccSdwanCiscoBFDFeatureTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cisco_bfd_feature_template.test", "multiplier", "3"),
					resource.TestCheckResourceAttr("sdwan_cisco_bfd_feature_template.test", "poll_interval", "800000"),
					resource.TestCheckResourceAttr("sdwan_cisco_bfd_feature_template.test", "default_dscp", "48"),
					resource.TestCheckResourceAttr("sdwan_cisco_bfd_feature_template.test", "colors.0.color", "private5"),
					resource.TestCheckResourceAttr("sdwan_cisco_bfd_feature_template.test", "colors.0.hello_interval", "1000"),
					resource.TestCheckResourceAttr("sdwan_cisco_bfd_feature_template.test", "colors.0.multiplier", "7"),
					resource.TestCheckResourceAttr("sdwan_cisco_bfd_feature_template.test", "colors.0.pmtu_discovery", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_bfd_feature_template.test", "colors.0.dscp", "46"),
				),
			},
		},
	})
}

func testAccSdwanCiscoBFDFeatureTemplateConfig_minimum() string {
	return `
	resource "sdwan_cisco_bfd_feature_template" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
	}
	`
}

func testAccSdwanCiscoBFDFeatureTemplateConfig_all() string {
	return `
	resource "sdwan_cisco_bfd_feature_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		multiplier = 3
		poll_interval = 800000
		default_dscp = 48
		colors = [{
			color = "private5"
			hello_interval = 1000
			multiplier = 7
			pmtu_discovery = true
			dscp = 46
		}]
	}
	`
}
