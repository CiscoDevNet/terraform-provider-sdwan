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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceSdwanPolicyObjectUnifiedIntrusionPreventionProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015_IN_PROGRESS") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015_IN_PROGRESS")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_unified_intrusion_prevention.test", "signature_set", "balanced"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_unified_intrusion_prevention.test", "inspection_mode", "detection"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_unified_intrusion_prevention.test", "log_level", "error"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_unified_intrusion_prevention.test", "custom_signature", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanPolicyObjectUnifiedIntrusionPreventionPrerequisitesProfileParcelConfig + testAccDataSourceSdwanPolicyObjectUnifiedIntrusionPreventionProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanPolicyObjectUnifiedIntrusionPreventionPrerequisitesProfileParcelConfig = `
resource "sdwan_policy_object_feature_profile" "test" {
  name = "POLICY_OBJECT_FP_1"
  description = "My policy object feature profile 1"
}

resource "sdwan_policy_object_security_ips_signature" "test" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = sdwan_policy_object_feature_profile.test.id
  entries = [
    {
      generator_id = "1234"
      signature_id = "5678"
    }
  ]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanPolicyObjectUnifiedIntrusionPreventionProfileParcelConfig() string {
	config := `resource "sdwan_policy_object_unified_intrusion_prevention" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_policy_object_feature_profile.test.id` + "\n"
	config += `	signature_set = "balanced"` + "\n"
	config += `	inspection_mode = "detection"` + "\n"
	config += `	ips_signature_allow_list_id = sdwan_policy_object_security_ips_signature.test.id` + "\n"
	config += `	log_level = "error"` + "\n"
	config += `	custom_signature = false` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_policy_object_unified_intrusion_prevention" "test" {
			id = sdwan_policy_object_unified_intrusion_prevention.test.id
			feature_profile_id = sdwan_policy_object_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
