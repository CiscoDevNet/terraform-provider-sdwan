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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

var _ datasource.DataSource = &NetworkHierarchySecurityLoggingDataSource{}

func NewNetworkHierarchySecurityLoggingDataSource() datasource.DataSource {
	return &NetworkHierarchySecurityLoggingDataSource{}
}

type NetworkHierarchySecurityLoggingDataSource struct {
	client *sdwan.Client
}

func (d *NetworkHierarchySecurityLoggingDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_hierarchy_security_logging"
}

func (d *NetworkHierarchySecurityLoggingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read Network Hierarchy Security Logging settings.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"node_id": schema.StringAttribute{
				MarkdownDescription: "The UUID of the network hierarchy node",
				Required:            true,
			},
			"high_speed_logging": schema.SetNestedAttribute{
				MarkdownDescription: "High speed logging configuration",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vrf": schema.StringAttribute{
							MarkdownDescription: "VRF name or ID",
							Computed:            true,
						},
						"server_ip": schema.StringAttribute{
							MarkdownDescription: "Server IPv4 or IPv6 address",
							Computed:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: "Server port number",
							Computed:            true,
						},
					},
				},
			},
			"utd_syslog": schema.ListNestedAttribute{
				MarkdownDescription: "UTD syslog configuration",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpn": schema.StringAttribute{
							MarkdownDescription: "VPN name or ID",
							Computed:            true,
						},
						"server_ip": schema.StringAttribute{
							MarkdownDescription: "Server IPv4 address",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworkHierarchySecurityLoggingDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

func (d *NetworkHierarchySecurityLoggingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkHierarchySecurityLogging

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

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
