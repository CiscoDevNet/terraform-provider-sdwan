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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &CloudProviderSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &CloudProviderSettingsDataSource{}
)

func NewCloudProviderSettingsDataSource() datasource.DataSource {
	return &CloudProviderSettingsDataSource{}
}

type CloudProviderSettingsDataSource struct {
	client *sdwan.Client
}

func (d *CloudProviderSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_provider_settings"
}

func (d *CloudProviderSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Cloud Provider Settings .",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"umbrella_org_id": schema.StringAttribute{
				MarkdownDescription: "Umbrella Organization ID (numeric string)",
				Computed:            true,
			},
			"umbrella_auth_key_v2": schema.StringAttribute{
				MarkdownDescription: "Umbrella Authentication Key (v2)",
				Computed:            true,
			},
			"umbrella_auth_secret_v2": schema.StringAttribute{
				MarkdownDescription: "Umbrella Authentication Secret (v2)",
				Computed:            true,
			},
			"umbrella_sig_auth_key": schema.StringAttribute{
				MarkdownDescription: "Umbrella SIG Authentication Key",
				Computed:            true,
			},
			"umbrella_sig_auth_secret": schema.StringAttribute{
				MarkdownDescription: "Umbrella SIG Authentication Secret",
				Computed:            true,
			},
			"umbrella_dns_auth_key": schema.StringAttribute{
				MarkdownDescription: "Umbrella DNS Authentication Key",
				Computed:            true,
			},
			"umbrella_dns_auth_secret": schema.StringAttribute{
				MarkdownDescription: "Umbrella DNS Authentication Secret",
				Computed:            true,
			},
			"zscaler_organization": schema.StringAttribute{
				MarkdownDescription: "Zscaler Organization",
				Computed:            true,
			},
			"zscaler_partner_base_uri": schema.StringAttribute{
				MarkdownDescription: "Zscaler Partner Base URI",
				Computed:            true,
			},
			"zscaler_partner_key": schema.StringAttribute{
				MarkdownDescription: "Zscaler Partner Key",
				Computed:            true,
			},
			"zscaler_username": schema.StringAttribute{
				MarkdownDescription: "Zscaler Username",
				Computed:            true,
			},
			"zscaler_password": schema.StringAttribute{
				MarkdownDescription: "Zscaler Password",
				Computed:            true,
			},
			"cisco_sse_org_id": schema.StringAttribute{
				MarkdownDescription: "Cisco SSE Organization ID (numeric string)",
				Computed:            true,
			},
			"cisco_sse_auth_key": schema.StringAttribute{
				MarkdownDescription: "Cisco SSE Authentication Key",
				Computed:            true,
			},
			"cisco_sse_auth_secret": schema.StringAttribute{
				MarkdownDescription: "Cisco SSE Authentication Secret",
				Computed:            true,
			},
			"cisco_sse_context_sharing": schema.BoolAttribute{
				MarkdownDescription: "Enable Cisco SSE Context Sharing",
				Computed:            true,
			},
		},
	}
}

func (d *CloudProviderSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *CloudProviderSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CloudProviderSettings

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var params = "?"

	res, err := d.client.Get(config.getPath() + params)
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
