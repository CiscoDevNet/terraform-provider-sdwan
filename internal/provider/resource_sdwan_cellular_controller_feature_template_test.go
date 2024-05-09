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
func TestAccSdwanCellularControllerFeatureTemplate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cellular_controller_feature_template.test", "cellular_interface_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cellular_controller_feature_template.test", "data_profiles.0.slot_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cellular_controller_feature_template.test", "data_profiles.0.data_profile", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cellular_controller_feature_template.test", "data_profiles.0.attach_profile", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cellular_controller_feature_template.test", "primary_sim_slot", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cellular_controller_feature_template.test", "sim_failover_retries", "160"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cellular_controller_feature_template.test", "sim_failover_timeout", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_cellular_controller_feature_template.test", "firmware_auto_sim", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCellularControllerFeatureTemplateConfig_minimum(),
			},
			{
				Config: testAccSdwanCellularControllerFeatureTemplateConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanCellularControllerFeatureTemplateConfig_minimum() string {
	config := `resource "sdwan_cellular_controller_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanCellularControllerFeatureTemplateConfig_all() string {
	config := `resource "sdwan_cellular_controller_feature_template" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += ` device_types = ["vedge-C8000V"]` + "\n"
	config += `	cellular_interface_id = "1"` + "\n"
	config += `	data_profiles = [{` + "\n"
	config += `	  slot_number = 1` + "\n"
	config += `	  data_profile = 8` + "\n"
	config += `	  attach_profile = 8` + "\n"
	config += `	}]` + "\n"
	config += `	primary_sim_slot = 100` + "\n"
	config += `	sim_failover_retries = 160` + "\n"
	config += `	sim_failover_timeout = 3` + "\n"
	config += `	firmware_auto_sim = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
