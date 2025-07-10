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
	_ datasource.DataSource              = &ApplicationAwareRoutingPolicyDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplicationAwareRoutingPolicyDefinitionDataSource{}
)

func NewApplicationAwareRoutingPolicyDefinitionDataSource() datasource.DataSource {
	return &ApplicationAwareRoutingPolicyDefinitionDataSource{}
}

type ApplicationAwareRoutingPolicyDefinitionDataSource struct {
	client *sdwan.Client
}

func (d *ApplicationAwareRoutingPolicyDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_aware_routing_policy_definition"
}

func (d *ApplicationAwareRoutingPolicyDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Application Aware Routing Policy Definition .",

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
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: "List of sequences",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: "Sequence ID",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Sequence name",
							Computed:            true,
						},
						"ip_type": schema.StringAttribute{
							MarkdownDescription: "Sequence IP type, either `ipv4`, `ipv6` or `all`",
							Computed:            true,
						},
						"match_entries": schema.ListNestedAttribute{
							MarkdownDescription: "List of match entries",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of match entry",
										Computed:            true,
									},
									"application_list_id": schema.StringAttribute{
										MarkdownDescription: "Application list ID",
										Computed:            true,
									},
									"application_list_version": schema.Int64Attribute{
										MarkdownDescription: "Application list version",
										Computed:            true,
									},
									"dns_application_list_id": schema.StringAttribute{
										MarkdownDescription: "DNS Application list ID",
										Computed:            true,
									},
									"dns_application_list_version": schema.Int64Attribute{
										MarkdownDescription: "DNS Application list version",
										Computed:            true,
									},
									"icmp_message": schema.StringAttribute{
										MarkdownDescription: "ICMP Message",
										Computed:            true,
									},
									"dns": schema.StringAttribute{
										MarkdownDescription: "DNS request or response",
										Computed:            true,
									},
									"dscp": schema.StringAttribute{
										MarkdownDescription: "DSCP value",
										Computed:            true,
									},
									"plp": schema.StringAttribute{
										MarkdownDescription: "PLP",
										Computed:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: "IP Protocol, 0-255 (Single value or multiple values separated by spaces)",
										Computed:            true,
									},
									"source_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "Source Data Prefix list ID",
										Computed:            true,
									},
									"source_data_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: "Source Data Prefix list version",
										Computed:            true,
									},
									"source_ip": schema.StringAttribute{
										MarkdownDescription: "Source IP",
										Computed:            true,
									},
									"source_port": schema.StringAttribute{
										MarkdownDescription: "Source port, 0-65535 (Single value, range or multiple values separated by spaces)",
										Computed:            true,
									},
									"destination_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "Destination Data Prefix list ID",
										Computed:            true,
									},
									"destination_data_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: "Destination Data Prefix list version",
										Computed:            true,
									},
									"destination_ip": schema.StringAttribute{
										MarkdownDescription: "Destination IP",
										Computed:            true,
									},
									"destination_port": schema.StringAttribute{
										MarkdownDescription: "Destination port, 0-65535 (Single value, range or multiple values separated by spaces)",
										Computed:            true,
									},
									"destination_region": schema.StringAttribute{
										MarkdownDescription: "Destination region",
										Computed:            true,
									},
									"traffic_to": schema.StringAttribute{
										MarkdownDescription: "Traffic to",
										Computed:            true,
									},
								},
							},
						},
						"action_entries": schema.ListNestedAttribute{
							MarkdownDescription: "List of action entries",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of action entry",
										Computed:            true,
									},
									"backup_sla_preferred_color": schema.StringAttribute{
										MarkdownDescription: "Backup SLA preferred color (Single value or multiple values separated by spaces)",
										Computed:            true,
									},
									"counter": schema.StringAttribute{
										MarkdownDescription: "Counter name",
										Computed:            true,
									},
									"log": schema.BoolAttribute{
										MarkdownDescription: "Enable logging",
										Computed:            true,
									},
									"cloud_sla": schema.BoolAttribute{
										MarkdownDescription: "Cloud SLA",
										Computed:            true,
									},
									"sla_class_parameters": schema.ListNestedAttribute{
										MarkdownDescription: "List of SLA class parameters",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: "Type of SLA class parameter",
													Computed:            true,
												},
												"sla_class_list": schema.StringAttribute{
													MarkdownDescription: "SLA class list ID",
													Computed:            true,
												},
												"sla_class_list_version": schema.Int64Attribute{
													MarkdownDescription: "SLA class list version",
													Computed:            true,
												},
												"preferred_color_group_list": schema.StringAttribute{
													MarkdownDescription: "Preferred color group list ID",
													Computed:            true,
												},
												"preferred_color_group_list_version": schema.Int64Attribute{
													MarkdownDescription: "Preferred color group list version",
													Computed:            true,
												},
												"preferred_color": schema.StringAttribute{
													MarkdownDescription: "preferred color (Single value or multiple values separated by spaces)",
													Computed:            true,
												},
											},
										},
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

func (d *ApplicationAwareRoutingPolicyDefinitionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *ApplicationAwareRoutingPolicyDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ApplicationAwareRoutingPolicyDefinition

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
	config.Type = types.StringValue("appRoute")

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
