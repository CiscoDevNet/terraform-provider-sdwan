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
func TestAccSdwanPolicyObjectPolicerProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" && os.Getenv("POLICY_OBJECT_FEATURE_TEMPLATE_ID") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012 or POLICY_OBJECT_FEATURE_TEMPLATE_ID")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_policer_profile_parcel.test", "entries.0.burst_bytes", "56500"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_policer_profile_parcel.test", "entries.0.select_value", "remark"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_policy_object_policer_profile_parcel.test", "entries.0.rate_bps", "60000"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanPolicyObjectPolicerProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanPolicyObjectPolicerProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanPolicyObjectPolicerProfileParcelConfig_minimum() string {
	config := `resource "sdwan_policy_object_policer_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = ` + os.Getenv("POLICY_OBJECT_FEATURE_TEMPLATE_ID") + `` + "\n"
	config += `	entries = [{` + "\n"
	config += `	  burst_bytes = 56500` + "\n"
	config += `	  select_value = "remark"` + "\n"
	config += `	  rate_bps = 60000` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanPolicyObjectPolicerProfileParcelConfig_all() string {
	config := `resource "sdwan_policy_object_policer_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = ` + os.Getenv("POLICY_OBJECT_FEATURE_TEMPLATE_ID") + `` + "\n"
	config += `	entries = [{` + "\n"
	config += `	  burst_bytes = 56500` + "\n"
	config += `	  select_value = "remark"` + "\n"
	config += `	  rate_bps = 60000` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
