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
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

var _ datasource.DataSource = &NetworkHierarchyNodeDataSource{}

func NewNetworkHierarchyNodeDataSource() datasource.DataSource {
	return &NetworkHierarchyNodeDataSource{}
}

type NetworkHierarchyNodeDataSource struct {
	client *sdwan.Client
}

func (d *NetworkHierarchyNodeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_hierarchy_node"
}

func (d *NetworkHierarchyNodeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read a Network Hierarchy Node.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the node",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the node",
				Computed:            true,
			},
			"parent_group": schema.StringAttribute{
				MarkdownDescription: "The name of the parent group",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "The type of node (group, region, or site)",
				Computed:            true,
			},
			"site_id": schema.Int64Attribute{
				MarkdownDescription: "The site ID (only for site type nodes)",
				Computed:            true,
			},
			"is_secondary": schema.BoolAttribute{
				MarkdownDescription: "Whether this is a secondary region (only for region type nodes)",
				Computed:            true,
			},
			"address": schema.SingleNestedAttribute{
				MarkdownDescription: "The address of the site (only for site type nodes)",
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"street": schema.StringAttribute{
						MarkdownDescription: "Street address",
						Computed:            true,
					},
					"city": schema.StringAttribute{
						MarkdownDescription: "City",
						Computed:            true,
					},
					"state": schema.StringAttribute{
						MarkdownDescription: "State or province",
						Computed:            true,
					},
					"country": schema.StringAttribute{
						MarkdownDescription: "Country",
						Computed:            true,
					},
					"zipcode": schema.StringAttribute{
						MarkdownDescription: "Zip or postal code",
						Computed:            true,
					},
				},
			},
			"controllers": schema.SetAttribute{
				MarkdownDescription: "List of controller UUIDs assigned to this region (only applicable for region type nodes)",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}

func (d *NetworkHierarchyNodeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

func (d *NetworkHierarchyNodeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkHierarchyNode

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get(config.getPath() + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	parentUuid := config.fromBody(ctx, res)

	hierarchyRes, err := d.client.Get(config.getHierarchyListPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve network hierarchy (GET), got error: %s, %s", err, hierarchyRes.String()))
		return
	}

	parentGroupName := config.resolveParentIdToGroup(hierarchyRes, parentUuid)
	if parentGroupName != "" {
		config.ParentGroup = types.StringValue(parentGroupName)
	} else {
		config.ParentGroup = types.StringNull()
	}

	if config.Type.ValueString() == "region" {
		regionId := res.Get("data.hierarchyId.regionId").Int()
		if regionId > 0 {
			controllersRes, err := d.client.Get(config.getControllersPath())
			if err != nil {
				resp.Diagnostics.AddWarning("Client Warning", fmt.Sprintf("Failed to retrieve controllers (GET), got error: %s", err))
			} else {
				config.readControllersFromResponse(ctx, controllersRes, regionId)
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.String()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
