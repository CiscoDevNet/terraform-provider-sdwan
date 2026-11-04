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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

// The list endpoint has no total count, so pages are read until one comes back short.
const networkHierarchyNodesPageSize = 1500

type NetworkHierarchyNodes struct {
	Id               types.String                         `tfsdk:"id"`
	Names            []types.String                       `tfsdk:"names"`
	ExpandContainers types.Bool                           `tfsdk:"expand_containers"`
	ResolvedSites    []NetworkHierarchyNodesResolvedSites `tfsdk:"resolved_sites"`
	ContainerSites   map[string][]types.String            `tfsdk:"container_sites"`
	Nodes            map[string]NetworkHierarchyNodesNode `tfsdk:"nodes"`
}

type NetworkHierarchyNodesResolvedSites struct {
	Name   types.String `tfsdk:"name"`
	Id     types.String `tfsdk:"id"`
	SiteId types.Int64  `tfsdk:"site_id"`
}

type NetworkHierarchyNodesNode struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

func (data NetworkHierarchyNodes) getHierarchyListPath(offset int) string {
	return fmt.Sprintf("/v1/network-hierarchy?offset=%d&limit=%d", offset, networkHierarchyNodesPageSize)
}

type networkHierarchyRawNode struct {
	id            string
	name          string
	hierarchyPath string
	kind          string
	siteId        int64
}

// Group nodes have no data.label at all, so an absent label means group.
func networkHierarchyNodeKind(label string) string {
	switch label {
	case "SITE":
		return "site"
	case "REGION":
		return "region"
	case "GLOBAL":
		return "global"
	default:
		return "group"
	}
}

func parseNetworkHierarchyNodes(page gjson.Result) []networkHierarchyRawNode {
	var nodes []networkHierarchyRawNode
	page.ForEach(func(_, v gjson.Result) bool {
		nodes = append(nodes, networkHierarchyRawNode{
			id:            v.Get("id").String(),
			name:          v.Get("name").String(),
			hierarchyPath: v.Get("hierarchyPath").String(),
			kind:          networkHierarchyNodeKind(v.Get("data.label").String()),
			siteId:        v.Get("data.hierarchyId.siteId").Int(),
		})
		return true
	})
	return nodes
}

type networkHierarchyResolution struct {
	// every requested name that exists, any kind
	nodes map[string]networkHierarchyRawNode
	// flat, deduplicated sites including group members, which are not requested and so not in nodes
	resolved []networkHierarchyRawNode
	// requested non-site name -> member site names; an empty container stays a non-nil empty slice
	containers map[string][]string
	// requested names with no matching node
	missing []string
}

// With expand false only nodes is filled: no site view is built.
func resolveNetworkHierarchyNames(names []string, all []networkHierarchyRawNode, expand bool) networkHierarchyResolution {
	byName := make(map[string]networkHierarchyRawNode, len(all))
	for _, n := range all {
		byName[n.name] = n
	}

	res := networkHierarchyResolution{
		nodes:      make(map[string]networkHierarchyRawNode),
		containers: make(map[string][]string),
	}
	seen := make(map[string]bool)
	add := func(n networkHierarchyRawNode) {
		if !seen[n.name] {
			seen[n.name] = true
			res.resolved = append(res.resolved, n)
		}
	}

	for _, want := range names {
		n, ok := byName[want]
		if !ok {
			res.missing = append(res.missing, want)
			continue
		}
		res.nodes[want] = n
		if !expand {
			continue
		}
		if n.kind == "site" {
			add(n)
			continue
		}
		if _, done := res.containers[want]; done {
			continue
		}
		members := []string{}
		prefix := n.hierarchyPath + "/" // the slash keeps "EMEA" from matching "EMEA1/..."
		for _, s := range all {
			if s.kind == "site" && strings.HasPrefix(s.hierarchyPath, prefix) {
				members = append(members, s.name)
				add(s)
			}
		}
		res.containers[want] = members
	}
	return res
}
