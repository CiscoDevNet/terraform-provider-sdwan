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
func TestAccDataSourceSdwanSystemMRFProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_mrf_feature.test", "region_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_mrf_feature.test", "secondary_region_id", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_mrf_feature.test", "role", "edge-router"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_mrf_feature.test", "enable_migration_to_mrf", "enabled"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_system_mrf_feature.test", "migration_bgp_community", "100"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanSystemMRFPrerequisitesProfileParcelConfig + testAccDataSourceSdwanSystemMRFProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanSystemMRFPrerequisitesProfileParcelConfig = `
resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanSystemMRFProfileParcelConfig() string {
	config := `resource "sdwan_system_mrf_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_system_feature_profile.test.id` + "\n"
	config += `	region_id = 1` + "\n"
	config += `	secondary_region_id = 2` + "\n"
	config += `	role = "edge-router"` + "\n"
	config += `	enable_migration_to_mrf = "enabled"` + "\n"
	config += `	migration_bgp_community = 100` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_system_mrf_feature" "test" {
			id = sdwan_system_mrf_feature.test.id
			feature_profile_id = sdwan_system_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
