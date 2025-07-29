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
func TestAccDataSourceSdwanPolicyGroup(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_group.test", "name", "PG_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_group.test", "description", "My policy group 1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_group.test", "solution", "sdwan"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_group.test", "devices.0.id", "C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_group.test", "devices.0.variables.0.name", "host_name"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_policy_group.test", "devices.0.variables.0.value", "edge1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanPolicyGroupPrerequisitesConfig + testAccDataSourceSdwanPolicyGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanPolicyGroupPrerequisitesConfig = `
resource "sdwan_application_priority_feature_profile" "test" {
  name        = "APPLICATION_PRIORITY_TF"
  description = "Terraform test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanPolicyGroupConfig() string {
	config := ""
	config += `resource "sdwan_policy_group" "test" {` + "\n"
	config += `	name = "PG_1"` + "\n"
	config += `	description = "My policy group 1"` + "\n"
	config += `	solution = "sdwan"` + "\n"
	config += `	feature_profile_ids = [sdwan_application_priority_feature_profile.test.id]` + "\n"
	config += `	devices = [{` + "\n"
	config += `	  id = "C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B"` + "\n"
	config += `	  variables = [{` + "\n"
	config += `		name = "host_name"` + "\n"
	config += `		value = "edge1"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_policy_group" "test" {
			id = sdwan_policy_group.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
