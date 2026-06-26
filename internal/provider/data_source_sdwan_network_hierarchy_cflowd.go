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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

var _ datasource.DataSource = &NetworkHierarchyCflowdDataSource{}

func NewNetworkHierarchyCflowdDataSource() datasource.DataSource {
	return &NetworkHierarchyCflowdDataSource{}
}

type NetworkHierarchyCflowdDataSource struct {
	client *sdwan.Client
}

func (d *NetworkHierarchyCflowdDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_hierarchy_cflowd"
}

func (d *NetworkHierarchyCflowdDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read Network Hierarchy Cflowd settings.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"node_id": schema.StringAttribute{
				MarkdownDescription: "The UUID of the Global network hierarchy node. This is automatically fetched from the SD-WAN Manager.",
				Computed:            true,
			},
			"flow_active_timeout": schema.Int64Attribute{
				MarkdownDescription: "Active flow timeout in seconds",
				Computed:            true,
			},
			"flow_inactive_timeout": schema.Int64Attribute{
				MarkdownDescription: "Inactive flow timeout in seconds",
				Computed:            true,
			},
			"flow_refresh_time": schema.Int64Attribute{
				MarkdownDescription: "Flow refresh time in seconds",
				Computed:            true,
			},
			"flow_sampling_interval": schema.Int64Attribute{
				MarkdownDescription: "Flow sampling interval",
				Computed:            true,
			},
			"collect_tloc_loopback": schema.BoolAttribute{
				MarkdownDescription: "Collect SDWAN TLOC loopback interface name instead of physical",
				Computed:            true,
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: "FNF Protocol",
				Computed:            true,
			},
			"collect_tos": schema.BoolAttribute{
				MarkdownDescription: "Collect TOS record field",
				Computed:            true,
			},
			"collect_dscp_output": schema.BoolAttribute{
				MarkdownDescription: "Collect remarked DSCP",
				Computed:            true,
			},
			"collectors": schema.SetNestedAttribute{
				MarkdownDescription: "List of collectors",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "VPN ID",
							Computed:            true,
						},
						"address": schema.StringAttribute{
							MarkdownDescription: "Collector IPv4 or IPv6 address",
							Computed:            true,
						},
						"udp_port": schema.Int64Attribute{
							MarkdownDescription: "Collector UDP port number",
							Computed:            true,
						},
						"export_spread": schema.BoolAttribute{
							MarkdownDescription: "Enable export spreading",
							Computed:            true,
						},
						"bfd_metrics_export": schema.BoolAttribute{
							MarkdownDescription: "Enable BFD metrics exporting",
							Computed:            true,
						},
						"export_interval": schema.Int64Attribute{
							MarkdownDescription: "BFD export interval in seconds",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworkHierarchyCflowdDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

func (d *NetworkHierarchyCflowdDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkHierarchyCflowd

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	hierarchyRes, err := d.client.Get(config.getNetworkHierarchyListPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve network hierarchy, got error: %s", err))
		return
	}

	globalNodeId, err := config.getGlobalNodeId(hierarchyRes)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to get Global node ID: %s", err))
		return
	}
	config.NodeId = types.StringValue(globalNodeId)

	res, err := d.client.Get(config.getPathWithId())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.String()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
