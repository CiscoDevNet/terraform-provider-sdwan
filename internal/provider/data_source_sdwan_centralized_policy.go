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

// Code generated by "gen/generator.go"; DO NOT EDIT.

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

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &CentralizedPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &CentralizedPolicyDataSource{}
)

func NewCentralizedPolicyDataSource() datasource.DataSource {
	return &CentralizedPolicyDataSource{}
}

type CentralizedPolicyDataSource struct {
	client *sdwan.Client
}

func (d *CentralizedPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_centralized_policy"
}

func (d *CentralizedPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Centralized Policy .",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the centralized policy",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the centralized policy",
				Computed:            true,
			},
			"definitions": schema.SetNestedAttribute{
				MarkdownDescription: "List of policy definitions",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Policy definition ID",
							Computed:            true,
						},
						"version": schema.Int64Attribute{
							MarkdownDescription: "Policy definition version",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Policy definition type",
							Computed:            true,
						},
						"entries": schema.SetNestedAttribute{
							MarkdownDescription: "List of entries",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"site_list_ids": schema.ListAttribute{
										MarkdownDescription: "List of site list IDs",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"site_list_versions": schema.ListAttribute{
										MarkdownDescription: "List of site list versions",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"vpn_list_ids": schema.ListAttribute{
										MarkdownDescription: "List of VPN list IDs",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"vpn_list_versions": schema.ListAttribute{
										MarkdownDescription: "List of VPN list versions",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"direction": schema.StringAttribute{
										MarkdownDescription: "Direction",
										Computed:            true,
									},
									"region_list_ids": schema.ListAttribute{
										MarkdownDescription: "List of region list IDs",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"region_list_versions": schema.ListAttribute{
										MarkdownDescription: "List of region list versions",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"region_ids": schema.ListAttribute{
										MarkdownDescription: "List of region IDs",
										ElementType:         types.StringType,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *CentralizedPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

func (d *CentralizedPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CentralizedPolicy

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get("/template/policy/vsmart/definition/" + config.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
