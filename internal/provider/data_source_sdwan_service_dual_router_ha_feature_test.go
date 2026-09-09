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
func TestAccDataSourceSdwanServiceDualRouterHAProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_dual_router_ha_feature.test", "redundancy_groups.0.group_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_dual_router_ha_feature.test", "redundancy_groups.0.tag_name", "example"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_dual_router_ha_feature.test", "enable_optimize_paths", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanServiceDualRouterHAPrerequisitesProfileParcelConfig + testAccDataSourceSdwanServiceDualRouterHAProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceSdwanServiceDualRouterHAPrerequisitesProfileParcelConfig + testAccDataSourceSdwanServiceDualRouterHAProfileParcelByNameConfig(),
				Check: resource.ComposeTestCheckFunc(
					append(checks,
						resource.TestCheckResourceAttr("data.sdwan_service_dual_router_ha_feature.test", "name", "TF_TEST"),
						resource.TestCheckResourceAttrSet("data.sdwan_service_dual_router_ha_feature.test", "id"),
					)...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanServiceDualRouterHAPrerequisitesProfileParcelConfig = `
resource "sdwan_service_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_service_lan_vpn_feature" "test" {
  name               = "TF_TEST_SLAN"
  vpn                = 1
  feature_profile_id = sdwan_service_feature_profile.test.id
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanServiceDualRouterHAProfileParcelConfig() string {
	config := `resource "sdwan_service_dual_router_ha_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	redundancy_groups = [{` + "\n"
	config += `	  group_id = 1` + "\n"
	config += `	  vpn_ids = [{` + "\n"
	config += `		vpn_id = sdwan_service_lan_vpn_feature.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	  tag_name = "example"` + "\n"
	config += `	}]` + "\n"
	config += `	enable_optimize_paths = false` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_service_dual_router_ha_feature" "test" {
			id = sdwan_service_dual_router_ha_feature.test.id
			feature_profile_id = sdwan_service_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceByNameConfig
func testAccDataSourceSdwanServiceDualRouterHAProfileParcelByNameConfig() string {
	config := `resource "sdwan_service_dual_router_ha_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	redundancy_groups = [{` + "\n"
	config += `	  group_id = 1` + "\n"
	config += `	  vpn_ids = [{` + "\n"
	config += `		vpn_id = sdwan_service_lan_vpn_feature.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	  tag_name = "example"` + "\n"
	config += `	}]` + "\n"
	config += `	enable_optimize_paths = false` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_service_dual_router_ha_feature" "test" {
			name = "TF_TEST"
			feature_profile_id = sdwan_service_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceByNameConfig
