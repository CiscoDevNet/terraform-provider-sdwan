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
func TestAccDataSourceSdwanTopologyCustomControlProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_topology_custom_control_feature.test", "default_action", "reject"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_topology_custom_control_feature.test", "sequences.0.id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_topology_custom_control_feature.test", "sequences.0.name", "Rule1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_topology_custom_control_feature.test", "sequences.0.base_action", "accept"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_topology_custom_control_feature.test", "sequences.0.type", "route"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_topology_custom_control_feature.test", "sequences.0.ip_type", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_topology_custom_control_feature.test", "sequences.0.match_entries.0.omp_tag", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_topology_custom_control_feature.test", "sequences.0.action_entries.0.set_parameters.0.preference", "100"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanTopologyCustomControlPrerequisitesProfileParcelConfig + testAccDataSourceSdwanTopologyCustomControlProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanTopologyCustomControlPrerequisitesProfileParcelConfig = `
resource "sdwan_topology_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanTopologyCustomControlProfileParcelConfig() string {
	config := `resource "sdwan_topology_custom_control_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_topology_feature_profile.test.id` + "\n"
	config += `	default_action = "reject"` + "\n"
	config += `	sequences = [{` + "\n"
	config += `	  id = 1` + "\n"
	config += `	  name = "Rule1"` + "\n"
	config += `	  base_action = "accept"` + "\n"
	config += `	  type = "route"` + "\n"
	config += `	  ip_type = "ipv4"` + "\n"
	config += `	  match_entries = [{` + "\n"
	config += `		omp_tag = 100` + "\n"
	config += `	}]` + "\n"
	config += `	  action_entries = [{` + "\n"
	config += `      set_parameters = [{` + "\n"
	config += `			preference = 100` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_topology_custom_control_feature" "test" {
			id = sdwan_topology_custom_control_feature.test.id
			feature_profile_id = sdwan_topology_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
