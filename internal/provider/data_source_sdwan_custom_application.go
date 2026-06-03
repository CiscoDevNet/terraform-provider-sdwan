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
	_ datasource.DataSource              = &CustomApplicationDataSource{}
	_ datasource.DataSourceWithConfigure = &CustomApplicationDataSource{}
)

func NewCustomApplicationDataSource() datasource.DataSource {
	return &CustomApplicationDataSource{}
}

type CustomApplicationDataSource struct {
	client *sdwan.Client
}

func (d *CustomApplicationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_application"
}

func (d *CustomApplicationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Custom Application .",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"app_name": schema.StringAttribute{
				MarkdownDescription: "Application Name",
				Computed:            true,
			},
			"server_names": schema.SetAttribute{
				MarkdownDescription: "Server Names (Fully Qualified Domain names or Regex starting with `*` but not ending with `*` or both separated by commas.)",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"l3l4": schema.SetNestedAttribute{
				MarkdownDescription: "L3/L4 Attributes",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_addresses": schema.SetAttribute{
							MarkdownDescription: "IPv4 Address (10.X.X.X, 20.0.0.0/24 separated by commas, subnet prefix length 24 to 32.)",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"l4_protocol": schema.StringAttribute{
							MarkdownDescription: "L4 Protocol",
							Computed:            true,
						},
						"ports": schema.StringAttribute{
							MarkdownDescription: "Ports (Space separated ports or range or both.)",
							Computed:            true,
						},
					},
				},
			},
			"application_family": schema.StringAttribute{
				MarkdownDescription: "Application Family",
				Computed:            true,
			},
			"application_group": schema.StringAttribute{
				MarkdownDescription: "Application Group",
				Computed:            true,
			},
			"traffic_class": schema.StringAttribute{
				MarkdownDescription: "Traffic Class",
				Computed:            true,
			},
			"business_relevance": schema.StringAttribute{
				MarkdownDescription: "Business Relevance",
				Computed:            true,
			},
		},
	}
}

func (d *CustomApplicationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *CustomApplicationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CustomApplication

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
