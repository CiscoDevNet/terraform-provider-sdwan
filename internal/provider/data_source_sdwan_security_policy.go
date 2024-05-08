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
	_ datasource.DataSource              = &SecurityPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &SecurityPolicyDataSource{}
)

func NewSecurityPolicyDataSource() datasource.DataSource {
	return &SecurityPolicyDataSource{}
}

type SecurityPolicyDataSource struct {
	client *sdwan.Client
}

func (d *SecurityPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_policy"
}

func (d *SecurityPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Security Policy .",

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
				MarkdownDescription: "The name of the security policy",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the security policy",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "The policy mode",
				Computed:            true,
			},
			"use_case": schema.StringAttribute{
				MarkdownDescription: "The use case of the security policy",
				Computed:            true,
			},
			"definitions": schema.ListNestedAttribute{
				MarkdownDescription: "List of policy definitions",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Policy definition ID",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Policy definition type",
							Computed:            true,
						},
					},
				},
			},
			"direct_internet_applications": schema.StringAttribute{
				MarkdownDescription: "Bypass firewall policy and allow all Internet traffic to/from VPN 0",
				Computed:            true,
			},
			"tcp_syn_flood_limit": schema.StringAttribute{
				MarkdownDescription: "TCP SYN Flood Limit, value from 1 to 4294967295",
				Computed:            true,
			},
			"audit_trail": schema.StringAttribute{
				MarkdownDescription: "Audit trail",
				Computed:            true,
			},
			"match_statistics_per_filter": schema.StringAttribute{
				MarkdownDescription: "Match Statistics per-filter",
				Computed:            true,
			},
			"failure_mode": schema.StringAttribute{
				MarkdownDescription: "Failure mode",
				Computed:            true,
			},
			"high_speed_logging_server_ip": schema.StringAttribute{
				MarkdownDescription: "High Speed Logging Server IP",
				Computed:            true,
			},
			"high_speed_logging_vpn": schema.StringAttribute{
				MarkdownDescription: "High Speed Logging VPN",
				Computed:            true,
			},
			"high_speed_logging_server_port": schema.StringAttribute{
				MarkdownDescription: "High Speed Logging Port",
				Computed:            true,
			},
			"logging": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"external_syslog_server_ip": schema.StringAttribute{
							MarkdownDescription: "External Syslog Server IP",
							Computed:            true,
						},
						"external_syslog_server_vpn": schema.StringAttribute{
							MarkdownDescription: "External Syslog Server VPN",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SecurityPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *SecurityPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SecurityPolicy

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get("/template/policy/security/definition/" + url.QueryEscape(config.Id.ValueString()))
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
