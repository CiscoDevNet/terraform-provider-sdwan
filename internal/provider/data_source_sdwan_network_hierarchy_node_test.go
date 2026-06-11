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
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test", "parent_id", "9cdc05d1-5306-41ef-8487-85829a4cfbe6"))
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
			parent_id = "9cdc05d1-5306-41ef-8487-85829a4cfbe6"
			id = sdwan_network_hierarchy_node.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
