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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanNetworkHierarchyNode(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "parent_group", "Global"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "name", "EMEA-Group"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "description", "EMEA group"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "type", "group"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchyNodeConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanNetworkHierarchyNodeConfig_all() string {
	config := `resource "sdwan_network_hierarchy_node" "test" {` + "\n"
	config += `	parent_group = "Global"` + "\n"
	config += `	name = "EMEA-Group"` + "\n"
	config += `	description = "EMEA group"` + "\n"
	config += `	type = "group"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Custom tests for comprehensive coverage - these persist through code generation

// TestAccSdwanNetworkHierarchyNodeSiteType tests creating a site node with full address details
func TestAccSdwanNetworkHierarchyNodeSiteType(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchyNodeConfig_siteType(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_node.test", "id"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "name", "EMEA-Site-001"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "type", "site"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "site_id", "1001"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "is_secondary", "false"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "address.street", "100 Technology Park"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "address.city", "London"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "address.state", "England"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "address.country", "United Kingdom"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "address.zipcode", "EC1A 1BB"),
				),
			},
		},
	})
}

func testAccSdwanNetworkHierarchyNodeConfig_siteType() string {
	config := `resource "sdwan_network_hierarchy_node" "test" {` + "\n"
	config += `	parent_group = "Global"` + "\n"
	config += `	name = "EMEA-Site-001"` + "\n"
	config += `	description = "Primary EMEA datacenter site"` + "\n"
	config += `	type = "site"` + "\n"
	config += `	site_id = 1001` + "\n"
	config += `	is_secondary = false` + "\n"
	config += `	address = {` + "\n"
	config += `	  street = "100 Technology Park"` + "\n"
	config += `	  city = "London"` + "\n"
	config += `	  state = "England"` + "\n"
	config += `	  country = "United Kingdom"` + "\n"
	config += `	  zipcode = "EC1A 1BB"` + "\n"
	config += `	}` + "\n"
	config += `}` + "\n"
	return config
}

// TestAccSdwanNetworkHierarchyNodeRegionType tests creating a region node
func TestAccSdwanNetworkHierarchyNodeRegionType(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchyNodeConfig_regionType(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_node.test", "id"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "name", "APAC-Region"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "type", "region"),
				),
			},
		},
	})
}

func testAccSdwanNetworkHierarchyNodeConfig_regionType() string {
	config := `resource "sdwan_network_hierarchy_node" "test" {` + "\n"
	config += `	parent_group = "Global"` + "\n"
	config += `	name = "APAC-Region"` + "\n"
	config += `	description = "Asia-Pacific region"` + "\n"
	config += `	type = "region"` + "\n"
	config += `	is_secondary = false` + "\n"
	config += `}` + "\n"
	return config
}

// TestAccSdwanNetworkHierarchyNodeNestedHierarchy tests creating nested parent-child nodes
func TestAccSdwanNetworkHierarchyNodeNestedHierarchy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchyNodeConfig_nestedHierarchy(),
				Check: resource.ComposeTestCheckFunc(
					// Check parent group
					resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_node.parent_group", "id"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.parent_group", "name", "Americas-Group"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.parent_group", "type", "group"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.parent_group", "parent_group", "Global"),
					// Check child region
					resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_node.child_region", "id"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.child_region", "name", "US-East-Region"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.child_region", "type", "region"),
					// Check grandchild site
					resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_node.child_site", "id"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.child_site", "name", "NYC-Site-001"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.child_site", "type", "site"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.child_site", "site_id", "2001"),
				),
			},
		},
	})
}

func testAccSdwanNetworkHierarchyNodeConfig_nestedHierarchy() string {
	config := `resource "sdwan_network_hierarchy_node" "parent_group" {` + "\n"
	config += `	parent_group = "Global"` + "\n"
	config += `	name = "Americas-Group"` + "\n"
	config += `	description = "Americas organizational group"` + "\n"
	config += `	type = "group"` + "\n"
	config += `}` + "\n"
	config += "\n"
	config += `resource "sdwan_network_hierarchy_node" "child_region" {` + "\n"
	config += `	parent_group = sdwan_network_hierarchy_node.parent_group.name` + "\n"
	config += `	name = "US-East-Region"` + "\n"
	config += `	description = "US East Coast region"` + "\n"
	config += `	type = "region"` + "\n"
	config += `	is_secondary = false` + "\n"
	config += `}` + "\n"
	config += "\n"
	config += `resource "sdwan_network_hierarchy_node" "child_site" {` + "\n"
	config += `	parent_group = sdwan_network_hierarchy_node.parent_group.name` + "\n"
	config += `	name = "NYC-Site-001"` + "\n"
	config += `	description = "New York City primary site"` + "\n"
	config += `	type = "site"` + "\n"
	config += `	site_id = 2001` + "\n"
	config += `	address = {` + "\n"
	config += `	  street = "350 Fifth Avenue"` + "\n"
	config += `	  city = "New York"` + "\n"
	config += `	  state = "NY"` + "\n"
	config += `	  country = "USA"` + "\n"
	config += `	  zipcode = "10118"` + "\n"
	config += `	}` + "\n"
	config += `}` + "\n"
	return config
}
