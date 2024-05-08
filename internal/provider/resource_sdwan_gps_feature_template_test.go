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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanGpsFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_gps_feature_template.test", "enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_gps_feature_template.test", "gps_mode", "ms-based"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_gps_feature_template.test", "nmea", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_gps_feature_template.test", "source_address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_gps_feature_template.test", "destination_address", "2.3.4.5"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_gps_feature_template.test", "destination_port", "1234"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanGpsFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanGpsFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanGpsFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_gps_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanGpsFeatureTemplateConfig_all() string {
	config := `resource "sdwan_gps_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	enable = true` + "\n"
	config += `	gps_mode = "ms-based"` + "\n"
	config += `	nmea = true` + "\n"
	config += `	source_address = "1.2.3.4"` + "\n"
	config += `	destination_address = "2.3.4.5"` + "\n"
	config += `	destination_port = 1234` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
