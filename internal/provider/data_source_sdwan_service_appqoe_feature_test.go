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
func TestAccDataSourceSdwanServiceAppQoEProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_appqoe_feature.test", "appqoe_device_role", "forwarder"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_appqoe_feature.test", "forwarder_controller_groups.0.appnav_controllers.0.address", "192.168.2.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_appqoe_feature.test", "forwarder_controller_groups.0.appnav_controllers.0.vpn", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_appqoe_feature.test", "forwarder_service_node_groups.0.service_nodes.0.address", "192.168.2.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_appqoe_feature.test", "forwarder_service_contexts.0.appnav_controller_group", "ACG-APPQOE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_appqoe_feature.test", "forwarder_service_contexts.0.service_node_group", "SNG-APPQOE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_appqoe_feature.test", "forwarder_service_contexts.0.enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_service_appqoe_feature.test", "forwarder_service_contexts.0.vpn", "0"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanServiceAppQoEPrerequisitesProfileParcelConfig + testAccDataSourceSdwanServiceAppQoEProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceSdwanServiceAppQoEPrerequisitesProfileParcelConfig + testAccDataSourceSdwanServiceAppQoEProfileParcelByNameConfig(),
				Check: resource.ComposeTestCheckFunc(
					append(checks,
						resource.TestCheckResourceAttr("data.sdwan_service_appqoe_feature.test", "name", "TF_TEST"),
						resource.TestCheckResourceAttrSet("data.sdwan_service_appqoe_feature.test", "id"),
					)...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanServiceAppQoEPrerequisitesProfileParcelConfig = `
resource "sdwan_service_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanServiceAppQoEProfileParcelConfig() string {
	config := `resource "sdwan_service_appqoe_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	appqoe_device_role = "forwarder"` + "\n"
	config += `	forwarder_controller_groups = [{` + "\n"
	config += `	  appnav_controllers = [{` + "\n"
	config += `		address = "192.168.2.1"` + "\n"
	config += `		vpn = 1` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	forwarder_service_node_groups = [{` + "\n"
	config += `	  service_nodes = [{` + "\n"
	config += `		address = "192.168.2.2"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	forwarder_service_contexts = [{` + "\n"
	config += `	  appnav_controller_group = "ACG-APPQOE"` + "\n"
	config += `	  service_node_group = "SNG-APPQOE"` + "\n"
	config += `	  enable = true` + "\n"
	config += `	  vpn = 0` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_service_appqoe_feature" "test" {
			id = sdwan_service_appqoe_feature.test.id
			feature_profile_id = sdwan_service_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceByNameConfig
func testAccDataSourceSdwanServiceAppQoEProfileParcelByNameConfig() string {
	config := `resource "sdwan_service_appqoe_feature" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_service_feature_profile.test.id` + "\n"
	config += `	appqoe_device_role = "forwarder"` + "\n"
	config += `	forwarder_controller_groups = [{` + "\n"
	config += `	  appnav_controllers = [{` + "\n"
	config += `		address = "192.168.2.1"` + "\n"
	config += `		vpn = 1` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	forwarder_service_node_groups = [{` + "\n"
	config += `	  service_nodes = [{` + "\n"
	config += `		address = "192.168.2.2"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `	forwarder_service_contexts = [{` + "\n"
	config += `	  appnav_controller_group = "ACG-APPQOE"` + "\n"
	config += `	  service_node_group = "SNG-APPQOE"` + "\n"
	config += `	  enable = true` + "\n"
	config += `	  vpn = 0` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_service_appqoe_feature" "test" {
			name = "TF_TEST"
			feature_profile_id = sdwan_service_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceByNameConfig
