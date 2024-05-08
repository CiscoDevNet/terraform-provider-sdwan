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
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &QoSMapPolicyDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &QoSMapPolicyDefinitionDataSource{}
)

func NewQoSMapPolicyDefinitionDataSource() datasource.DataSource {
	return &QoSMapPolicyDefinitionDataSource{}
}

type QoSMapPolicyDefinitionDataSource struct {
	client *sdwan.Client
}

func (d *QoSMapPolicyDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_qos_map_policy_definition"
}

func (d *QoSMapPolicyDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the QoS Map Policy Definition .",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the policy definition",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the policy definition",
				Computed:            true,
			},
			"qos_schedulers": schema.ListNestedAttribute{
				MarkdownDescription: "List of QoS schedulers",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"queue": schema.Int64Attribute{
							MarkdownDescription: "Queue number",
							Computed:            true,
						},
						"class_map_id": schema.StringAttribute{
							MarkdownDescription: "Class map ID",
							Computed:            true,
						},
						"class_map_version": schema.Int64Attribute{
							MarkdownDescription: "Class map version",
							Computed:            true,
						},
						"bandwidth_percent": schema.Int64Attribute{
							MarkdownDescription: "Bandwidth percent",
							Computed:            true,
						},
						"buffer_percent": schema.Int64Attribute{
							MarkdownDescription: "Buffer percent",
							Computed:            true,
						},
						"burst": schema.Int64Attribute{
							MarkdownDescription: "Burst size",
							Computed:            true,
						},
						"drop_type": schema.StringAttribute{
							MarkdownDescription: "Drop type",
							Computed:            true,
						},
						"scheduling_type": schema.StringAttribute{
							MarkdownDescription: "Scheduling type",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *QoSMapPolicyDefinitionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *QoSMapPolicyDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config QoSMapPolicyDefinition

	// Read config
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

	config.fromBody(ctx, res)
	config.Type = types.StringValue("qosMap")

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
