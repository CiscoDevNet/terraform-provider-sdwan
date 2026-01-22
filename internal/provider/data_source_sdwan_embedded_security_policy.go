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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &EmbeddedSecurityProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &EmbeddedSecurityProfileParcelDataSource{}
)

func NewEmbeddedSecurityProfileParcelDataSource() datasource.DataSource {
	return &EmbeddedSecurityProfileParcelDataSource{}
}

type EmbeddedSecurityProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *EmbeddedSecurityProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_embedded_security_policy"
}

func (d *EmbeddedSecurityProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Embedded Security Policy.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Policy",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Policy",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Policy",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Policy",
				Computed:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: "Feature Profile ID",
				Required:            true,
			},
			"assembly": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"advanced_inspection_profile_policy_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ssl_decryption_profile_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"ngfw_policy_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"entries": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"source_zone_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"destination_zone_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"tcp_syn_flood_limit": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"max_incomplete_tcp_limit": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"max_incomplete_udp_limit": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"max_incomplete_icmp_limit": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"audit_trail": schema.StringAttribute{
				MarkdownDescription: "Setting can be string 'on' or missing for off",
				Computed:            true,
			},
			"unified_logging": schema.StringAttribute{
				MarkdownDescription: "Setting can be string 'on' or missing for off",
				Computed:            true,
			},
			"session_reclassify_allow": schema.StringAttribute{
				MarkdownDescription: "Setting can be string 'on' or missing for off",
				Computed:            true,
			},
			"imcp_unreachable_allow": schema.StringAttribute{
				MarkdownDescription: "Setting can be string 'on' or missing for off",
				Computed:            true,
			},
			"failure_mode": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"security_logging": schema.StringAttribute{
				MarkdownDescription: "HSL and UTD syslog, pulled from network settings page",
				Computed:            true,
			},
			"nat": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"nat_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"download_url_database_on_device": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"download_url_database_on_device_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"resource_profile": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"resource_profile_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
		},
	}
}

func (d *EmbeddedSecurityProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *EmbeddedSecurityProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config EmbeddedSecurity

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
