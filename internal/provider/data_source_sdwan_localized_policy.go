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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &LocalizedPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &LocalizedPolicyDataSource{}
)

func NewLocalizedPolicyDataSource() datasource.DataSource {
	return &LocalizedPolicyDataSource{}
}

type LocalizedPolicyDataSource struct {
	client *sdwan.Client
}

func (d *LocalizedPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_localized_policy"
}

func (d *LocalizedPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Localized Policy .",

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
				MarkdownDescription: "The name of the localized policy",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the localized policy",
				Computed:            true,
			},
			"flow_visibility_ipv4": schema.BoolAttribute{
				MarkdownDescription: "IPv4 flow visibility",
				Computed:            true,
			},
			"flow_visibility_ipv6": schema.BoolAttribute{
				MarkdownDescription: "IPv6 flow visibility",
				Computed:            true,
			},
			"application_visibility_ipv4": schema.BoolAttribute{
				MarkdownDescription: "IPv4 application visibility",
				Computed:            true,
			},
			"application_visibility_ipv6": schema.BoolAttribute{
				MarkdownDescription: "IPv6 application visibility",
				Computed:            true,
			},
			"cloud_qos": schema.BoolAttribute{
				MarkdownDescription: "Cloud QoS",
				Computed:            true,
			},
			"cloud_qos_service_side": schema.BoolAttribute{
				MarkdownDescription: "Cloud QoS service side",
				Computed:            true,
			},
			"implicit_acl_logging": schema.BoolAttribute{
				MarkdownDescription: "Implicit ACL logging",
				Computed:            true,
			},
			"log_frequency": schema.Int64Attribute{
				MarkdownDescription: "Log frequency",
				Computed:            true,
			},
			"ipv4_visibility_cache_entries": schema.Int64Attribute{
				MarkdownDescription: "IPv4 visibility cache entries",
				Computed:            true,
			},
			"ipv6_visibility_cache_entries": schema.Int64Attribute{
				MarkdownDescription: "IPv6 visibility cache entries",
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
					},
				},
			},
		},
	}
}

func (d *LocalizedPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *LocalizedPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config LocalizedPolicy

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get("/template/policy/vedge/definition/" + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
