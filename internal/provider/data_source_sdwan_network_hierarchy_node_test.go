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

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceSdwanNetworkHierarchyNode(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanNetworkHierarchyNodeConfig_withResource(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.sdwan_network_hierarchy_node.test_site", "id"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test_site", "parent_group", "Global"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test_site", "name", "EMEA-Site"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test_site", "description", "EMEA Site"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test_site", "type", "site"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test_site", "site_id", "101"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test_site", "address.state", "NY"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test_site", "address.street", "350 Fifth Avenue"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test_site", "address.city", "New York"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test_site", "address.country", "USA"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test_site", "address.zipcode", "10118"),

					resource.TestCheckResourceAttrSet("data.sdwan_network_hierarchy_node.test_region", "id"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test_region", "parent_group", "Global"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_node.test_region", "name", "EMEA-Region"),
				),
			},
		},
	})
}

func testAccDataSourceSdwanNetworkHierarchyNodeConfig_withResource() string {
	config := `resource "sdwan_network_hierarchy_node" "test_site" {` + "\n"
	config += `	parent_group = "Global"` + "\n"
	config += `	name = "EMEA-Site"` + "\n"
	config += `	description = "EMEA Site"` + "\n"
	config += `	type = "site"` + "\n"
	config += `	site_id = 101` + "\n"
	config += `	address = {` + "\n"
	config += `	  street = "350 Fifth Avenue"` + "\n"
	config += `	  city = "New York"` + "\n"
	config += `	  state = "NY"` + "\n"
	config += `	  country = "USA"` + "\n"
	config += `	  zipcode = "10118"` + "\n"
	config += `	}` + "\n"
	config += `}` + "\n"
	config += "\n"
	config += `resource "sdwan_network_hierarchy_node" "test_region" {` + "\n"
	config += `	parent_group = "Global"` + "\n"
	config += `	name = "EMEA-Region"` + "\n"
	config += `	description = "EMEA Region"` + "\n"
	config += `	type = "region"` + "\n"
	config += `	is_secondary = false` + "\n"
	config += `}` + "\n"
	config += "\n"
	config += `data "sdwan_network_hierarchy_node" "test_site" {` + "\n"
	config += `	id = sdwan_network_hierarchy_node.test_site.id` + "\n"
	config += `}` + "\n"
	config += `data "sdwan_network_hierarchy_node" "test_region" {` + "\n"
	config += `	id = sdwan_network_hierarchy_node.test_region.id` + "\n"
	config += `}` + "\n"
	return config
}
