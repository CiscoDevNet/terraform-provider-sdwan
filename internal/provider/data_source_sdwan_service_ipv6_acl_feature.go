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
	_ datasource.DataSource              = &ServiceIPv6ACLProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &ServiceIPv6ACLProfileParcelDataSource{}
)

func NewServiceIPv6ACLProfileParcelDataSource() datasource.DataSource {
	return &ServiceIPv6ACLProfileParcelDataSource{}
}

type ServiceIPv6ACLProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *ServiceIPv6ACLProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_ipv6_acl_feature"
}

func (d *ServiceIPv6ACLProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Service IPv6 ACL Feature.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Feature",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Feature",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Feature",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Feature",
				Computed:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: "Feature Profile ID",
				Required:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: "Default Action",
				Computed:            true,
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: "Access Control List",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"sequence_id": schema.Int64Attribute{
							MarkdownDescription: "Sequence Id",
							Computed:            true,
						},
						"sequence_name": schema.StringAttribute{
							MarkdownDescription: "Sequence Name",
							Computed:            true,
						},
						"base_action": schema.StringAttribute{
							MarkdownDescription: "Base Action",
							Computed:            true,
						},
						"conditions": schema.ListNestedAttribute{
							MarkdownDescription: "Define match conditions",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"next_header": schema.Int64Attribute{
										MarkdownDescription: "next header number",
										Computed:            true,
									},
									"packet_length": schema.Int64Attribute{
										MarkdownDescription: "Packet Length",
										Computed:            true,
									},
									"source_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"source_data_prefix": schema.StringAttribute{
										MarkdownDescription: "Source Data IP Prefix",
										Computed:            true,
									},
									"source_ports": schema.ListNestedAttribute{
										MarkdownDescription: "Source Port List",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"port": schema.Int64Attribute{
													MarkdownDescription: "source port range or individual port number",
													Computed:            true,
												},
											},
										},
									},
									"destination_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"destination_data_prefix": schema.StringAttribute{
										MarkdownDescription: "Destination Data IP Prefix",
										Computed:            true,
									},
									"destination_ports": schema.ListNestedAttribute{
										MarkdownDescription: "Destination Port List",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"port": schema.Int64Attribute{
													MarkdownDescription: "destination port range or individual port number",
													Computed:            true,
												},
											},
										},
									},
									"tcp_state": schema.StringAttribute{
										MarkdownDescription: "TCP States",
										Computed:            true,
									},
									"traffic_class": schema.SetAttribute{
										MarkdownDescription: "Select Traffic Class",
										ElementType:         types.Int64Type,
										Computed:            true,
									},
									"icmp_messages": schema.SetAttribute{
										MarkdownDescription: "ICMP6 Message",
										ElementType:         types.StringType,
										Computed:            true,
									},
								},
							},
						},
						"actions": schema.ListNestedAttribute{
							MarkdownDescription: "Define list of actions",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"accept_counter_name": schema.StringAttribute{
										MarkdownDescription: "Counter Name",
										Computed:            true,
									},
									"accept_log": schema.BoolAttribute{
										MarkdownDescription: "Enable Log",
										Computed:            true,
									},
									"accept_next_hop": schema.StringAttribute{
										MarkdownDescription: "Set Next Hop (IPV6 address)",
										Computed:            true,
									},
									"accept_traffic_class": schema.Int64Attribute{
										MarkdownDescription: "set traffic class number",
										Computed:            true,
									},
									"accept_mirror_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"accept_policer_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"drop_counter_name": schema.StringAttribute{
										MarkdownDescription: "Counter Name",
										Computed:            true,
									},
									"drop_log": schema.BoolAttribute{
										MarkdownDescription: "Enable Log",
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

func (d *ServiceIPv6ACLProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *ServiceIPv6ACLProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ServiceIPv6ACL

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
