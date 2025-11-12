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
	_ datasource.DataSource              = &ZoneBasedFirewallPolicyDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &ZoneBasedFirewallPolicyDefinitionDataSource{}
)

func NewZoneBasedFirewallPolicyDefinitionDataSource() datasource.DataSource {
	return &ZoneBasedFirewallPolicyDefinitionDataSource{}
}

type ZoneBasedFirewallPolicyDefinitionDataSource struct {
	client *sdwan.Client
}

func (d *ZoneBasedFirewallPolicyDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zone_based_firewall_policy_definition"
}

func (d *ZoneBasedFirewallPolicyDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Zone Based Firewall Policy Definition .",

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
			"apply_zone_pairs": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_zone": schema.StringAttribute{
							MarkdownDescription: "Source Zone",
							Computed:            true,
						},
						"destination_zone": schema.StringAttribute{
							MarkdownDescription: "Destination Zone",
							Computed:            true,
						},
					},
				},
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: "Default Action",
				Computed:            true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"rule_order": schema.Int64Attribute{
							MarkdownDescription: "Rule",
							Computed:            true,
						},
						"rule_name": schema.StringAttribute{
							MarkdownDescription: "Rule name",
							Computed:            true,
						},
						"base_action": schema.StringAttribute{
							MarkdownDescription: "Base action",
							Computed:            true,
						},
						"ip_type": schema.StringAttribute{
							MarkdownDescription: "Rule Type",
							Computed:            true,
						},
						"match_entries": schema.SetNestedAttribute{
							MarkdownDescription: "List of match entries",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of match entry",
										Computed:            true,
									},
									"policy_id": schema.StringAttribute{
										MarkdownDescription: "policy id for selected match entry",
										Computed:            true,
									},
									"policy_version": schema.Int64Attribute{
										MarkdownDescription: "Policy version",
										Computed:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "value for selected match entry",
										Computed:            true,
									},
									"protocol_type": schema.StringAttribute{
										MarkdownDescription: "Should be included with additionally entries for `destinationPort` and `protocol` whenever the type `protocolName` is used.",
										Computed:            true,
									},
									"value_variable": schema.StringAttribute{
										MarkdownDescription: "variable value for selected match entry if it has variable option (sourceIp & destinationIp)",
										Computed:            true,
									},
								},
							},
						},
						"action_entries": schema.SetNestedAttribute{
							MarkdownDescription: "List of actions entries",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of action entry",
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

func (d *ZoneBasedFirewallPolicyDefinitionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *ZoneBasedFirewallPolicyDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ZoneBasedFirewallPolicyDefinition

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
