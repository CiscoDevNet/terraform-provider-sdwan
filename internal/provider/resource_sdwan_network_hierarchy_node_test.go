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
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanNetworkHierarchyNode(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
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

// TestAccSdwanNetworkHierarchyNodeSiteType tests creating a site node.
// On SD-WAN Manager 20.18+ (SDWAN_2018 set) the site uses the geolocation model
// (location/latitude/longitude); on earlier versions it uses the address block.
func TestAccSdwanNetworkHierarchyNodeSiteType(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_node.test", "id"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "name", "EMEA-Site-001"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "type", "site"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "site_id", "1001"))
	if os.Getenv("SDWAN_2018") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "location", "London"))
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "latitude", "51.5074"))
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "longitude", "-0.1278"))
	} else {
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "address.street", "100 Technology Park"))
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "address.city", "London"))
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "address.state", "England"))
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "address.country", "United Kingdom"))
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "address.zipcode", "EC1A 1BB"))
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchyNodeConfig_siteType(),
				Check:  resource.ComposeTestCheckFunc(checks...),
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
	if os.Getenv("SDWAN_2018") != "" {
		config += `	location = "London"` + "\n"
		config += `	latitude = 51.5074` + "\n"
		config += `	longitude = -0.1278` + "\n"
	} else {
		config += `	address = {` + "\n"
		config += `	  street = "100 Technology Park"` + "\n"
		config += `	  city = "London"` + "\n"
		config += `	  state = "England"` + "\n"
		config += `	  country = "United Kingdom"` + "\n"
		config += `	  zipcode = "EC1A 1BB"` + "\n"
		config += `	}` + "\n"
	}
	config += `}` + "\n"
	return config
}

// TestAccSdwanNetworkHierarchyNodeRegionType tests creating a region node
func TestAccSdwanNetworkHierarchyNodeRegionType(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
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
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
	var checks []resource.TestCheckFunc
	// Check parent group
	checks = append(checks, resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_node.parent_group", "id"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.parent_group", "name", "Americas-Group"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.parent_group", "type", "group"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.parent_group", "parent_group", "Global"))
	// Check child region
	checks = append(checks, resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_node.child_region", "id"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.child_region", "name", "US-East-Region"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.child_region", "type", "region"))
	// Check grandchild site
	checks = append(checks, resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_node.child_site", "id"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.child_site", "name", "NYC-Site-001"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.child_site", "type", "site"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.child_site", "site_id", "2001"))
	if os.Getenv("SDWAN_2018") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.child_site", "location", "New York"))
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.child_site", "latitude", "40.7484"))
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.child_site", "longitude", "-73.9857"))
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchyNodeConfig_nestedHierarchy(),
				Check:  resource.ComposeTestCheckFunc(checks...),
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
	if os.Getenv("SDWAN_2018") != "" {
		config += `	location = "New York"` + "\n"
		config += `	latitude = 40.7484` + "\n"
		config += `	longitude = -73.9857` + "\n"
	} else {
		config += `	address = {` + "\n"
		config += `	  street = "350 Fifth Avenue"` + "\n"
		config += `	  city = "New York"` + "\n"
		config += `	  state = "NY"` + "\n"
		config += `	  country = "USA"` + "\n"
		config += `	  zipcode = "10118"` + "\n"
		config += `	}` + "\n"
	}
	config += `}` + "\n"
	return config
}

// TestAccSdwanNetworkHierarchyNodeGroupUnderRegion tests creating a group nested
// under a region as its parent_group. A region can parent a group or a site, but
// not another region - see TestAccSdwanNetworkHierarchyNodeRegionUnderRegionRejected.
func TestAccSdwanNetworkHierarchyNodeGroupUnderRegion(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchyNodeConfig_groupUnderRegion(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_node.child_group", "id"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.child_group", "name", "EMEA-Sub-Group"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.child_group", "type", "group"),
					resource.TestCheckResourceAttrPair("sdwan_network_hierarchy_node.child_group", "parent_group", "sdwan_network_hierarchy_node.parent_region", "name"),
				),
			},
		},
	})
}

func testAccSdwanNetworkHierarchyNodeConfig_groupUnderRegion() string {
	config := `resource "sdwan_network_hierarchy_node" "parent_region" {` + "\n"
	config += `	parent_group = "Global"` + "\n"
	config += `	name = "EMEA-Parent-Region"` + "\n"
	config += `	description = "Region used as the parent of a nested group"` + "\n"
	config += `	type = "region"` + "\n"
	config += `	is_secondary = false` + "\n"
	config += `}` + "\n"
	config += "\n"
	config += `resource "sdwan_network_hierarchy_node" "child_group" {` + "\n"
	config += `	parent_group = sdwan_network_hierarchy_node.parent_region.name` + "\n"
	config += `	name = "EMEA-Sub-Group"` + "\n"
	config += `	description = "Group nested under a region"` + "\n"
	config += `	type = "group"` + "\n"
	config += `}` + "\n"
	return config
}

// TestAccSdwanNetworkHierarchyNodeRegionUnderRegionRejected verifies that a
// region cannot be nested under another region - only under a group or Global.
// See TestAccSdwanNetworkHierarchyNodeGroupUnderRegion for the group case, which
// a region can parent.
func TestAccSdwanNetworkHierarchyNodeRegionUnderRegionRejected(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccSdwanNetworkHierarchyNodeConfig_regionUnderRegionRejected(),
				ExpectError: regexp.MustCompile(`Parent group 'EMEA-Outer-Region' not found in network hierarchy`),
			},
		},
	})
}

func testAccSdwanNetworkHierarchyNodeConfig_regionUnderRegionRejected() string {
	config := `resource "sdwan_network_hierarchy_node" "outer_region" {` + "\n"
	config += `	parent_group = "Global"` + "\n"
	config += `	name = "EMEA-Outer-Region"` + "\n"
	config += `	description = "Region that another region incorrectly nests under"` + "\n"
	config += `	type = "region"` + "\n"
	config += `	is_secondary = false` + "\n"
	config += `}` + "\n"
	config += "\n"
	config += `resource "sdwan_network_hierarchy_node" "inner_region" {` + "\n"
	config += `	parent_group = sdwan_network_hierarchy_node.outer_region.name` + "\n"
	config += `	name = "EMEA-Inner-Region"` + "\n"
	config += `	description = "Region incorrectly nested under another region"` + "\n"
	config += `	type = "region"` + "\n"
	config += `	is_secondary = false` + "\n"
	config += `}` + "\n"
	return config
}

// TestAccSdwanNetworkHierarchyNodeGeoDefault verifies that on SD-WAN Manager
// 20.18+ a site created without a `location` defaults to "Undisclosed" (the
// provider always sends `data.location`) and that a second apply of the same
// config is a no-op (idempotent). Only relevant on 20.18+, so it is skipped
// unless SDWAN_2018 is set.
func TestAccSdwanNetworkHierarchyNodeGeoDefault(t *testing.T) {
	if os.Getenv("SDWAN_2018") == "" {
		t.Skip("Skipping: geolocation default only applies on SD-WAN Manager 20.18+ (set SDWAN_2018)")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchyNodeConfig_geoDefault(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("sdwan_network_hierarchy_node.test", "id"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "name", "EMEA-Site-Undisclosed"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "type", "site"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "site_id", "1002"),
					resource.TestCheckResourceAttr("sdwan_network_hierarchy_node.test", "location", "Undisclosed"),
				),
			},
			{
				// Re-apply the identical config: no drift expected.
				Config:   testAccSdwanNetworkHierarchyNodeConfig_geoDefault(),
				PlanOnly: true,
			},
		},
	})
}

func testAccSdwanNetworkHierarchyNodeConfig_geoDefault() string {
	config := `resource "sdwan_network_hierarchy_node" "test" {` + "\n"
	config += `	parent_group = "Global"` + "\n"
	config += `	name = "EMEA-Site-Undisclosed"` + "\n"
	config += `	description = "Site without an explicit location"` + "\n"
	config += `	type = "site"` + "\n"
	config += `	site_id = 1002` + "\n"
	config += `}` + "\n"
	return config
}
