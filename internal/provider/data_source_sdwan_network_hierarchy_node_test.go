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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceSdwanNetworkHierarchyNode(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test", "parent_group", "Global"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test", "name", "EMEA-Group"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test", "description", "EMEA group"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test", "type", "group"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test", "is_secondary", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test", "site_id", "101"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanNetworkHierarchyNodeConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanNetworkHierarchyNodeConfig() string {
	config := ""

	config += `
		data "sdwan_network_hierarchy_node" "test" {
			id = sdwan_network_hierarchy_node.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig

// Custom data source test - creates resource first, then reads via data source

func TestAccDataSourceSdwanNetworkHierarchyNodeWithResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanNetworkHierarchyNodeConfig_withResource(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.sdwan_network_hierarchy_node.test", "id"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test", "parent_group", "Global"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test", "name", "EMEA-Group"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test", "description", "EMEA group"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test", "type", "group"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test", "is_secondary", "false"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test", "site_id", "101"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test", "address.0.street", "123 Main St"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test", "address.0.city", "New York"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test", "controllers.#", "2"),
				),
			},
		},
	})
}

func testAccDataSourceSdwanNetworkHierarchyNodeConfig_withResource() string {
	config := `resource "sdwan_network_hierarchy_node" "test" {` + "\n"
	config += `	parent_group = "Global"` + "\n"
	config += `	name = "EMEA-Group"` + "\n"
	config += `	description = "EMEA group"` + "\n"
	config += `	type = "group"` + "\n"
	config += `	is_secondary = false` + "\n"
	config += `	site_id = 101` + "\n"
	config += `	address = [{` + "\n"
	config += `	  street = "123 Main St"` + "\n"
	config += `	  city = "New York"` + "\n"
	config += `	  state = "NY"` + "\n"
	config += `	  country = "USA"` + "\n"
	config += `	  zipcode = "10001"` + "\n"
	config += `	}]` + "\n"
	config += `	controllers = ["controller-uuid-1", "controller-uuid-2"]` + "\n"
	config += `}` + "\n"
	config += "\n"
	config += `data "sdwan_network_hierarchy_node" "test" {` + "\n"
	config += `	id = sdwan_network_hierarchy_node.test.id` + "\n"
	config += `}` + "\n"
	return config
}
