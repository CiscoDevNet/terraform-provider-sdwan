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
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

var (
	_ datasource.DataSource              = &NetworkHierarchyNodesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkHierarchyNodesDataSource{}
)

func NewNetworkHierarchyNodesDataSource() datasource.DataSource {
	return &NetworkHierarchyNodesDataSource{}
}

type NetworkHierarchyNodesDataSource struct {
	client *sdwan.Client
}

func (d *NetworkHierarchyNodesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_hierarchy_nodes"
}

func (d *NetworkHierarchyNodesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "This data source resolves a list of Network Hierarchy site, group and region names " +
			"without requiring the hierarchy to be managed by this provider, in a single paginated read. " +
			"`nodes` returns the UUID and type of every requested node, whatever its type, for consumers that " +
			"reference a node itself (for example a scope). `resolved_sites` and `container_sites` expand groups " +
			"and regions to the sites nested underneath them, at any depth, for consumers that need sites " +
			"(for example topology targeting, where the Manager GUI sends the member sites and never the group's own UUID). " +
			"Set `expand_containers` to `false` to return only `nodes`.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"names": schema.ListAttribute{
				MarkdownDescription: "Site, group, or region names to resolve. A name that does not exist in " +
					"the Network Hierarchy fails the read.",
				Required:    true,
				ElementType: types.StringType,
			},
			"expand_containers": schema.BoolAttribute{
				MarkdownDescription: "Whether to build the site view: expand each group, region or `Global` in `names` to the sites " +
					"nested underneath it. Defaults to `true`. When `false`, `resolved_sites` and `container_sites` are both empty " +
					"and only `nodes` is returned. Use `false` when only the nodes' own ids are needed, so that a name like " +
					"`Global` does not write every site in the hierarchy to state.",
				Optional: true,
			},
			"nodes": schema.MapNestedAttribute{
				MarkdownDescription: "One entry per name in `names`, keyed by name, whatever the node's type. " +
					"Always populated, regardless of `expand_containers`.",
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The node's Network Hierarchy UUID",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "The type of node: `site`, `region`, `group` or `global`",
							Computed:            true,
						},
					},
				},
			},
			"resolved_sites": schema.ListNestedAttribute{
				MarkdownDescription: "The resolved, deduplicated set of sites: the sites named in `names` plus every site " +
					"nested under a named group or region, so the UUID of a group member is available even though the member " +
					"was not itself requested. Never includes the groups or regions themselves; use `nodes` for their ids. " +
					"Empty when `expand_containers` is `false`.",
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "The site's name",
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: "The site's Network Hierarchy node UUID",
							Computed:            true,
						},
						"site_id": schema.Int64Attribute{
							MarkdownDescription: "The site's numeric site ID",
							Computed:            true,
						},
					},
				},
			},
			"container_sites": schema.MapAttribute{
				MarkdownDescription: "For each name in `names` that is a group, region, or `Global` (anything that is not a site), " +
					"the names of the sites nested underneath it, at any depth. Use this to expand a group per consumer: " +
					"`resolved_sites` is one flat, deduplicated set and cannot say which group a site came from. " +
					"Site names in `names` never appear as keys. A container with no sites underneath it is present " +
					"with an empty list - it is not an error here, so the caller decides whether an empty selection is acceptable.",
				Computed:    true,
				ElementType: types.ListType{ElemType: types.StringType},
			},
		},
	}
}

func (d *NetworkHierarchyNodesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

func (d *NetworkHierarchyNodesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkHierarchyNodes

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "NetworkHierarchyNodes: Beginning Read")

	var all []networkHierarchyRawNode
	for offset := 0; ; offset += networkHierarchyNodesPageSize {
		res, err := d.client.Get(config.getHierarchyListPath(offset))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve network hierarchy list, got error: %s", err))
			return
		}
		page := parseNetworkHierarchyNodes(res)
		all = append(all, page...)
		if len(page) < networkHierarchyNodesPageSize {
			break
		}
	}

	names := make([]string, len(config.Names))
	for i, n := range config.Names {
		names[i] = n.ValueString()
	}

	// Data source schemas have no Default, so an unset expand_containers means true.
	expand := config.ExpandContainers.IsNull() || config.ExpandContainers.ValueBool()

	res := resolveNetworkHierarchyNames(names, all, expand)
	if len(res.missing) > 0 {
		resp.Diagnostics.AddError(
			"Unresolved Network Hierarchy References",
			fmt.Sprintf("No site, group, or region found in the Network Hierarchy for: %v", res.missing),
		)
		return
	}

	config.ResolvedSites = make([]NetworkHierarchyNodesResolvedSites, len(res.resolved))
	for i, n := range res.resolved {
		config.ResolvedSites[i] = NetworkHierarchyNodesResolvedSites{
			Name:   types.StringValue(n.name),
			Id:     types.StringValue(n.id),
			SiteId: types.Int64Value(n.siteId),
		}
	}

	config.Nodes = make(map[string]NetworkHierarchyNodesNode, len(res.nodes))
	for name, n := range res.nodes {
		config.Nodes[name] = NetworkHierarchyNodesNode{
			Id:   types.StringValue(n.id),
			Type: types.StringValue(n.kind),
		}
	}

	config.ContainerSites = make(map[string][]types.String, len(res.containers))
	for container, members := range res.containers {
		values := make([]types.String, len(members))
		for i, m := range members {
			values[i] = types.StringValue(m)
		}
		config.ContainerSites[container] = values
	}

	tflog.Debug(ctx, "NetworkHierarchyNodes: Read finished successfully")

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
