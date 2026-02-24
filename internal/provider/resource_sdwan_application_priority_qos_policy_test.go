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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanApplicationPriorityQoSProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_application_priority_qos_policy.test", "qos_schedulers.0.drops", "tail-drop"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_application_priority_qos_policy.test", "qos_schedulers.0.queue", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_application_priority_qos_policy.test", "qos_schedulers.0.bandwidth", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_application_priority_qos_policy.test", "qos_schedulers.0.scheduling_type", "llq"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanApplicationPriorityQoSPrerequisitesProfileParcelConfig + testAccSdwanApplicationPriorityQoSProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanApplicationPriorityQoSPrerequisitesProfileParcelConfig + testAccSdwanApplicationPriorityQoSProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanApplicationPriorityQoSPrerequisitesProfileParcelConfig = `
resource "sdwan_application_priority_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanApplicationPriorityQoSProfileParcelConfig_minimum() string {
	config := `resource "sdwan_application_priority_qos_policy" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_application_priority_feature_profile.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanApplicationPriorityQoSProfileParcelConfig_all() string {
	config := `resource "sdwan_application_priority_qos_policy" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_application_priority_feature_profile.test.id` + "\n"
	config += `	target_interfaces = ["{{interface_var_1}}"]` + "\n"
	config += `	qos_schedulers = [{` + "\n"
	config += `	  drops = "tail-drop"` + "\n"
	config += `	  queue = "0"` + "\n"
	config += `	  bandwidth = "10"` + "\n"
	config += `	  scheduling_type = "llq"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
