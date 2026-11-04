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
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/tidwall/gjson"
)

// Creates its own nhds-* nodes, so no assertion depends on what the Manager already holds.
func TestAccDataSourceSdwanNetworkHierarchyNodes(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanNetworkHierarchyNodesConfig_resolve(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_nodes.test", "resolved_sites.#", "3"),
					resource.TestCheckTypeSetElemNestedAttrs("data.sdwan_network_hierarchy_nodes.test", "resolved_sites.*", map[string]string{
						"name":    "nhds-site-a",
						"site_id": "9101",
					}),
					resource.TestCheckTypeSetElemNestedAttrs("data.sdwan_network_hierarchy_nodes.test", "resolved_sites.*", map[string]string{
						"name":    "nhds-site-b",
						"site_id": "9102",
					}),
					resource.TestCheckTypeSetElemNestedAttrs("data.sdwan_network_hierarchy_nodes.test", "resolved_sites.*", map[string]string{
						"name":    "nhds-site-c",
						"site_id": "9103",
					}),
					// only requested containers are keys
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_nodes.test", "container_sites.%", "2"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_nodes.test", "container_sites.nhds-group.#", "2"),
					resource.TestCheckTypeSetElemAttr("data.sdwan_network_hierarchy_nodes.test", "container_sites.nhds-group.*", "nhds-site-b"),
					resource.TestCheckTypeSetElemAttr("data.sdwan_network_hierarchy_nodes.test", "container_sites.nhds-group.*", "nhds-site-c"),
					// an empty group is present as [], not absent
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_nodes.test", "container_sites.nhds-empty-group.#", "0"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_nodes.test", "nodes.%", "3"),
					resource.TestCheckResourceAttrSet("data.sdwan_network_hierarchy_nodes.test", "nodes.nhds-group.id"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_nodes.test", "nodes.nhds-group.type", "group"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_nodes.test", "nodes.nhds-site-a.type", "site"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_nodes.test", "nodes.nhds-empty-group.type", "group"),
				),
			},
			{
				// scope case: node ids only; expanding Global would pull every site into state
				Config: testAccDataSourceSdwanNetworkHierarchyNodesConfig_noExpand(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_nodes.test", "nodes.%", "3"),
					resource.TestCheckResourceAttrSet("data.sdwan_network_hierarchy_nodes.test", "nodes.Global.id"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_nodes.test", "nodes.Global.type", "global"),
					resource.TestCheckResourceAttrSet("data.sdwan_network_hierarchy_nodes.test", "nodes.nhds-group.id"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_nodes.test", "nodes.nhds-group.type", "group"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_nodes.test", "nodes.nhds-site-a.type", "site"),
					// no site view at all, even for the site named directly (its id is in nodes)
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_nodes.test", "container_sites.%", "0"),
					resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_nodes.test", "resolved_sites.#", "0"),
				),
			},
			{
				Config:      testAccDataSourceSdwanNetworkHierarchyNodesConfig_missing(),
				ExpectError: regexp.MustCompile(`(?s)Unresolved Network Hierarchy References.*nhds-does-not-exist`),
			},
		},
	})
}

// Shared by every step: a standalone site, a group with a site and a child group with a site, and an empty group.
func networkHierarchyNodesTestFixture() string {
	config := `resource "sdwan_network_hierarchy_node" "site_a" {` + "\n"
	config += `	parent_group = "Global"` + "\n"
	config += `	name = "nhds-site-a"` + "\n"
	config += `	type = "site"` + "\n"
	config += `	site_id = 9101` + "\n"
	config += networkHierarchyNodeSiteBody()
	config += `}` + "\n\n"

	config += `resource "sdwan_network_hierarchy_node" "group" {` + "\n"
	config += `	parent_group = "Global"` + "\n"
	config += `	name = "nhds-group"` + "\n"
	config += `	type = "group"` + "\n"
	config += `}` + "\n\n"

	config += `resource "sdwan_network_hierarchy_node" "site_b" {` + "\n"
	config += `	parent_group = sdwan_network_hierarchy_node.group.name` + "\n"
	config += `	name = "nhds-site-b"` + "\n"
	config += `	type = "site"` + "\n"
	config += `	site_id = 9102` + "\n"
	config += networkHierarchyNodeSiteBody()
	config += `}` + "\n\n"

	config += `resource "sdwan_network_hierarchy_node" "child_group" {` + "\n"
	config += `	parent_group = sdwan_network_hierarchy_node.group.name` + "\n"
	config += `	name = "nhds-child-group"` + "\n"
	config += `	type = "group"` + "\n"
	config += `}` + "\n\n"

	config += `resource "sdwan_network_hierarchy_node" "site_c" {` + "\n"
	config += `	parent_group = sdwan_network_hierarchy_node.child_group.name` + "\n"
	config += `	name = "nhds-site-c"` + "\n"
	config += `	type = "site"` + "\n"
	config += `	site_id = 9103` + "\n"
	config += networkHierarchyNodeSiteBody()
	config += `}` + "\n\n"

	config += `resource "sdwan_network_hierarchy_node" "empty_group" {` + "\n"
	config += `	parent_group = "Global"` + "\n"
	config += `	name = "nhds-empty-group"` + "\n"
	config += `	type = "group"` + "\n"
	config += `}` + "\n\n"
	return config
}

func testAccDataSourceSdwanNetworkHierarchyNodesConfig_resolve() string {
	config := networkHierarchyNodesTestFixture()
	config += `data "sdwan_network_hierarchy_nodes" "test" {` + "\n"
	config += `	names = [sdwan_network_hierarchy_node.group.name, sdwan_network_hierarchy_node.site_a.name, sdwan_network_hierarchy_node.empty_group.name]` + "\n"
	config += `	depends_on = [sdwan_network_hierarchy_node.site_b, sdwan_network_hierarchy_node.site_c]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccDataSourceSdwanNetworkHierarchyNodesConfig_noExpand() string {
	config := networkHierarchyNodesTestFixture()
	config += `data "sdwan_network_hierarchy_nodes" "test" {` + "\n"
	config += `	names = [sdwan_network_hierarchy_node.group.name, sdwan_network_hierarchy_node.site_a.name, "Global"]` + "\n"
	config += `	expand_containers = false` + "\n"
	config += `}` + "\n"
	return config
}

func testAccDataSourceSdwanNetworkHierarchyNodesConfig_missing() string {
	config := networkHierarchyNodesTestFixture()
	config += `data "sdwan_network_hierarchy_nodes" "test" {` + "\n"
	config += `	names = [sdwan_network_hierarchy_node.site_a.name, "nhds-does-not-exist"]` + "\n"
	config += `}` + "\n"
	return config
}

// Groups carry no data.label, sites carry label and siteId.
func nhNodeJSON(name, path, label string, siteID int) string {
	var data string
	switch label {
	case "SITE":
		data = fmt.Sprintf(`"data":{"label":"SITE","hierarchyId":{"siteId":%d}},`, siteID)
	case "":
		data = `"data":{"parentUuid":"x"},`
	default:
		data = fmt.Sprintf(`"data":{"label":%q},`, label)
	}
	return fmt.Sprintf(`{"id":"%s-id",%s"name":%q,"hierarchyPath":%q}`, name, data, name, path)
}

// Covers a region under a group (a site four levels deep), a name that prefixes another ("EMEA" vs "EMEA1"), and an empty region.
func nhFixture() []networkHierarchyRawNode {
	nodes := []string{
		nhNodeJSON("Global", "Global", "GLOBAL", 0),
		nhNodeJSON("EMEA1", "Global/EMEA1", "", 0),
		nhNodeJSON("test1", "Global/EMEA1/test1", "REGION", 0),
		nhNodeJSON("SITE_1", "Global/EMEA1/test1/SITE_1", "SITE", 1),
		nhNodeJSON("SITE_300", "Global/EMEA1/SITE_300", "SITE", 300),
		nhNodeJSON("Region_EMEA", "Global/Region_EMEA", "REGION", 0),
		nhNodeJSON("SITE_2", "Global/Region_EMEA/SITE_2", "SITE", 2),
		nhNodeJSON("EMEA", "Global/EMEA", "", 0),
		nhNodeJSON("SITE_9", "Global/EMEA/SITE_9", "SITE", 9),
		nhNodeJSON("REGION1", "Global/REGION1", "REGION", 0),
		nhNodeJSON("SITE_100", "Global/SITE_100", "SITE", 100),
	}
	return parseNetworkHierarchyNodes(gjson.Parse("[" + strings.Join(nodes, ",") + "]"))
}

func nhNames(nodes []networkHierarchyRawNode) []string {
	names := []string{}
	for _, n := range nodes {
		names = append(names, n.name)
	}
	return names
}

// Pure resolution logic, so no Manager is needed; covers cases that would take many extra nodes to create.
func TestResolveNetworkHierarchyNames(t *testing.T) {
	all := nhFixture()
	tests := []struct {
		name           string
		names          []string
		noExpand       bool
		wantResolved   []string
		wantContainers map[string][]string
		wantNodes      map[string]string // name -> type; nil means not checked
		wantMissing    []string
	}{
		{
			name:           "site resolves directly and is never a container key",
			names:          []string{"SITE_100"},
			wantResolved:   []string{"SITE_100"},
			wantContainers: map[string][]string{},
		},
		{
			name:           "group expands through a nested region",
			names:          []string{"EMEA1"},
			wantResolved:   []string{"SITE_1", "SITE_300"},
			wantContainers: map[string][]string{"EMEA1": {"SITE_1", "SITE_300"}},
		},
		{
			name:         "two groups keep per-group attribution that the flat list loses",
			names:        []string{"EMEA1", "Region_EMEA"},
			wantResolved: []string{"SITE_1", "SITE_300", "SITE_2"},
			wantContainers: map[string][]string{
				"EMEA1":       {"SITE_1", "SITE_300"},
				"Region_EMEA": {"SITE_2"},
			},
		},
		{
			name:           "path prefix does not leak across sibling names",
			names:          []string{"EMEA"},
			wantResolved:   []string{"SITE_9"},
			wantContainers: map[string][]string{"EMEA": {"SITE_9"}},
		},
		{
			name:           "empty container is present with an empty list, not an error",
			names:          []string{"REGION1"},
			wantResolved:   []string{},
			wantContainers: map[string][]string{"REGION1": {}},
		},
		{
			name:           "group plus its own member dedups the flat list, member is not a key",
			names:          []string{"EMEA1", "SITE_1"},
			wantResolved:   []string{"SITE_1", "SITE_300"},
			wantContainers: map[string][]string{"EMEA1": {"SITE_1", "SITE_300"}},
		},
		{
			name:           "same container requested twice is expanded once",
			names:          []string{"EMEA1", "EMEA1"},
			wantResolved:   []string{"SITE_1", "SITE_300"},
			wantContainers: map[string][]string{"EMEA1": {"SITE_1", "SITE_300"}},
		},
		{
			name:           "unknown names are reported while the rest still resolve",
			names:          []string{"SITE_100", "nope", "also-nope"},
			wantResolved:   []string{"SITE_100"},
			wantContainers: map[string][]string{},
			wantMissing:    []string{"nope", "also-nope"},
		},
		{
			name:         "nodes reports every requested name with its type, whatever the kind",
			names:        []string{"EMEA1", "SITE_100", "Region_EMEA"},
			wantResolved: []string{"SITE_1", "SITE_300", "SITE_100", "SITE_2"},
			wantContainers: map[string][]string{
				"EMEA1":       {"SITE_1", "SITE_300"},
				"Region_EMEA": {"SITE_2"},
			},
			wantNodes: map[string]string{"EMEA1": "group", "SITE_100": "site", "Region_EMEA": "region"},
		},
		{
			name:           "expand off: only nodes is populated, no site view at all",
			names:          []string{"EMEA1", "SITE_100", "Region_EMEA", "Global"},
			noExpand:       true,
			wantResolved:   []string{},
			wantContainers: map[string][]string{},
			wantNodes:      map[string]string{"EMEA1": "group", "SITE_100": "site", "Region_EMEA": "region", "Global": "global"},
		},
		{
			name:           "expand off still reports unknown names",
			names:          []string{"EMEA1", "nope"},
			noExpand:       true,
			wantResolved:   []string{},
			wantContainers: map[string][]string{},
			wantNodes:      map[string]string{"EMEA1": "group"},
			wantMissing:    []string{"nope"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := resolveNetworkHierarchyNames(tc.names, all, !tc.noExpand)
			if got := nhNames(res.resolved); !reflect.DeepEqual(got, tc.wantResolved) {
				t.Errorf("resolved = %v, want %v", got, tc.wantResolved)
			}
			if !reflect.DeepEqual(res.containers, tc.wantContainers) {
				t.Errorf("containers = %v, want %v", res.containers, tc.wantContainers)
			}
			if !reflect.DeepEqual(res.missing, tc.wantMissing) {
				t.Errorf("missing = %v, want %v", res.missing, tc.wantMissing)
			}
			if tc.wantNodes != nil {
				got := map[string]string{}
				for name, n := range res.nodes {
					got[name] = n.kind
				}
				if !reflect.DeepEqual(got, tc.wantNodes) {
					t.Errorf("nodes = %v, want %v", got, tc.wantNodes)
				}
			}
			for name, n := range res.nodes {
				if n.id != name+"-id" {
					t.Errorf("node %q carries id %q, want it to point at its own node", name, n.id)
				}
			}
			for name, members := range res.containers {
				if members == nil {
					t.Errorf("container %q has nil members; an empty container must be a non-nil empty slice", name)
				}
			}
		})
	}
}
