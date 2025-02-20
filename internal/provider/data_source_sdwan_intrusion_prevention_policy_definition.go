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
	_ datasource.DataSource              = &IntrusionPreventionPolicyDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &IntrusionPreventionPolicyDefinitionDataSource{}
)

func NewIntrusionPreventionPolicyDefinitionDataSource() datasource.DataSource {
	return &IntrusionPreventionPolicyDefinitionDataSource{}
}

type IntrusionPreventionPolicyDefinitionDataSource struct {
	client *sdwan.Client
}

func (d *IntrusionPreventionPolicyDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_intrusion_prevention_policy_definition"
}

func (d *IntrusionPreventionPolicyDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Intrusion Prevention Policy Definition .",

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
				MarkdownDescription: "The name of the policy definition",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the policy definition",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "The policy mode",
				Computed:            true,
			},
			"inspection_mode": schema.StringAttribute{
				MarkdownDescription: "The inspection mode",
				Computed:            true,
			},
			"log_level": schema.StringAttribute{
				MarkdownDescription: "Log level",
				Computed:            true,
			},
			"custom_signature": schema.BoolAttribute{
				MarkdownDescription: "Custom signature",
				Computed:            true,
			},
			"signature_set": schema.StringAttribute{
				MarkdownDescription: "Signature set",
				Computed:            true,
			},
			"ips_signature_list_id": schema.StringAttribute{
				MarkdownDescription: "IPS signature list ID",
				Computed:            true,
			},
			"ips_signature_list_version": schema.Int64Attribute{
				MarkdownDescription: "IPS signature list version",
				Computed:            true,
			},
			"target_vpns": schema.SetAttribute{
				MarkdownDescription: "List of VPN IDs",
				ElementType:         types.StringType,
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

func (d *IntrusionPreventionPolicyDefinitionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *IntrusionPreventionPolicyDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config IntrusionPreventionPolicyDefinition

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
