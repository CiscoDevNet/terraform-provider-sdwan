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
func TestAccDataSourceSdwanPolicyObjectExtendedCommunityProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" && os.Getenv("POLICY_OBJECT_FEATURE_TEMPLATE_ID") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012 or POLICY_OBJECT_FEATURE_TEMPLATE_ID")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_object_extended_community_profile_parcel.test", "entries.0.extended_community", "soo 10.0.0.1:30"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanPolicyObjectExtendedCommunityProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanPolicyObjectExtendedCommunityProfileParcelConfig() string {
	config := `resource "sdwan_policy_object_extended_community_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = ` + "\"" + os.Getenv("POLICY_OBJECT_FEATURE_TEMPLATE_ID") + "\"" + `` + "\n"
	config += `	entries = [{` + "\n"
	config += `	  extended_community = "soo 10.0.0.1:30"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_policy_object_extended_community_profile_parcel" "test" {
			id = sdwan_policy_object_extended_community_profile_parcel.test.id
			feature_profile_id = ` + "\"" + os.Getenv("POLICY_OBJECT_FEATURE_TEMPLATE_ID") + "\"" + `
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig