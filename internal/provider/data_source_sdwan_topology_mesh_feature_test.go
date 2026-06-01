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
func TestAccDataSourceSdwanTopologyMeshProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanTopologyMeshPrerequisitesProfileParcelConfig + testAccDataSourceSdwanTopologyMeshProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceSdwanTopologyMeshPrerequisitesProfileParcelConfig + testAccDataSourceSdwanTopologyMeshProfileParcelByNameConfig(),
				Check: resource.ComposeTestCheckFunc(
					append(checks,
						resource.TestCheckResourceAttr("data.sdwan_topology_mesh_feature.test", "name", "TF_TEST"),
						resource.TestCheckResourceAttrSet("data.sdwan_topology_mesh_feature.test", "id"),
					)...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanTopologyMeshPrerequisitesProfileParcelConfig = `
resource "sdwan_topology_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanTopologyMeshProfileParcelConfig() string {
	config := `resource "sdwan_topology_mesh_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_topology_feature_profile.test.id` + "\n"
	config += `	target_vpns = ["service_lan_vpn1"]` + "\n"
	config += `	sites = ["SITE_100"]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_topology_mesh_feature" "test" {
			id = sdwan_topology_mesh_feature.test.id
			feature_profile_id = sdwan_topology_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceByNameConfig
func testAccDataSourceSdwanTopologyMeshProfileParcelByNameConfig() string {
	config := `resource "sdwan_topology_mesh_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_topology_feature_profile.test.id` + "\n"
	config += `	target_vpns = ["service_lan_vpn1"]` + "\n"
	config += `	sites = ["SITE_100"]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_topology_mesh_feature" "test" {
			name = "TF_TEST"
			feature_profile_id = sdwan_topology_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceByNameConfig
