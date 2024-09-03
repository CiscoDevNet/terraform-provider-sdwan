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
	_ datasource.DataSource              = &ServiceMulticastProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &ServiceMulticastProfileParcelDataSource{}
)

func NewServiceMulticastProfileParcelDataSource() datasource.DataSource {
	return &ServiceMulticastProfileParcelDataSource{}
}

type ServiceMulticastProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *ServiceMulticastProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_multicast_profile_parcel"
}

func (d *ServiceMulticastProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Service Multicast profile parcel.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the profile parcel",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the profile parcel",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the profile parcel",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the profile parcel",
				Computed:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: "Feature Profile ID",
				Required:            true,
			},
			"spt_only": schema.BoolAttribute{
				MarkdownDescription: "Shortest Path Tree (SPT) Only Mode",
				Computed:            true,
			},
			"spt_only_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"local_replicator": schema.BoolAttribute{
				MarkdownDescription: "Replicator is local to this device",
				Computed:            true,
			},
			"local_replicator_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"threshold": schema.Int64Attribute{
				MarkdownDescription: "Set number of joins per group the router supports",
				Computed:            true,
			},
			"threshold_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"igmp_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "Set IGMP interface parameters",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Set interface name",
							Computed:            true,
						},
						"interface_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"version": schema.Int64Attribute{
							MarkdownDescription: "igmp Version <1..3>",
							Computed:            true,
						},
						"join_groups": schema.ListNestedAttribute{
							MarkdownDescription: "Configure static joins",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"group_address": schema.StringAttribute{
										MarkdownDescription: "Set group address",
										Computed:            true,
									},
									"group_address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"source_address": schema.StringAttribute{
										MarkdownDescription: "Set source address",
										Computed:            true,
									},
									"source_address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"pim_enable_source_specific_multicast": schema.BoolAttribute{
				MarkdownDescription: "turn SSM on/off",
				Computed:            true,
			},
			"pim_spt_threshold": schema.StringAttribute{
				MarkdownDescription: "Set Access List for PIM SSM",
				Computed:            true,
			},
			"pim_spt_threshold_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"pim_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "Set PIM interface parameters",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Set interface name",
							Computed:            true,
						},
						"interface_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"query_interval": schema.Int64Attribute{
							MarkdownDescription: "Set PIM query interval",
							Computed:            true,
						},
						"query_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"join_prune_interval": schema.Int64Attribute{
							MarkdownDescription: "Set interval at which PIM multicast traffic can join or be removed from RPT or SPT",
							Computed:            true,
						},
						"join_prune_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"static_rp_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Set Static RP Address(es)",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "Set Static RP IP Address",
							Computed:            true,
						},
						"ip_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"access_list": schema.StringAttribute{
							MarkdownDescription: "Set Static RP Access List",
							Computed:            true,
						},
						"access_list_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"override": schema.BoolAttribute{
							MarkdownDescription: "Set override flag",
							Computed:            true,
						},
						"override_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"enable_auto_rp": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable auto-RP",
				Computed:            true,
			},
			"enable_auto_rp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"auto_rp_announces": schema.ListNestedAttribute{
				MarkdownDescription: "Enable or disable RP Announce",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Set RP Announce Interface Name",
							Computed:            true,
						},
						"interface_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"scope": schema.Int64Attribute{
							MarkdownDescription: "Set RP Announce Scope",
							Computed:            true,
						},
						"scope_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"auto_rp_discoveries": schema.ListNestedAttribute{
				MarkdownDescription: "Enable or disable RP Discovery",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Set RP Discovery Interface Name",
							Computed:            true,
						},
						"interface_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"scope": schema.Int64Attribute{
							MarkdownDescription: "Set RP Discovery Scope",
							Computed:            true,
						},
						"scope_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"pim_bsr_rp_candidates": schema.ListNestedAttribute{
				MarkdownDescription: "Set RP Discovery Scope",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Set Autonomic-Networking virtual interface",
							Computed:            true,
						},
						"interface_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"access_list_id": schema.StringAttribute{
							MarkdownDescription: "Set IP Access List for PIM RP Candidate",
							Computed:            true,
						},
						"access_list_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"interval": schema.Int64Attribute{
							MarkdownDescription: "Set RP candidate advertisement interval",
							Computed:            true,
						},
						"interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: "Set RP candidate priority",
							Computed:            true,
						},
						"priority_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"pim_bsr_candidates": schema.ListNestedAttribute{
				MarkdownDescription: "bsr candidate Attributes",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Set Autonomic-Networking virtual interface",
							Computed:            true,
						},
						"interface_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"hash_mask_length": schema.Int64Attribute{
							MarkdownDescription: "Hash Mask length for RP selection",
							Computed:            true,
						},
						"hash_mask_length_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: "Set RP candidate priority",
							Computed:            true,
						},
						"priority_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"accept_candidate_access_list": schema.StringAttribute{
							MarkdownDescription: "Set BSR RP candidate filter",
							Computed:            true,
						},
						"accept_candidate_access_list_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"msdp_groups": schema.ListNestedAttribute{
				MarkdownDescription: "multicast MSDP peer",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"mesh_group_name": schema.StringAttribute{
							MarkdownDescription: "Set MSDP mesh group",
							Computed:            true,
						},
						"mesh_group_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"peers": schema.ListNestedAttribute{
							MarkdownDescription: "Configure peer",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"peer_ip": schema.StringAttribute{
										MarkdownDescription: "Set MSDP peer ip",
										Computed:            true,
									},
									"peer_ip_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"connection_source_interface": schema.StringAttribute{
										MarkdownDescription: "Set MSDP peer ip connect-source interface",
										Computed:            true,
									},
									"connection_source_interface_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"remote_as": schema.Int64Attribute{
										MarkdownDescription: "Set MSDP peer ip remote autonomous system number",
										Computed:            true,
									},
									"remote_as_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"peer_authentication_password": schema.StringAttribute{
										MarkdownDescription: "Set MSDP peer ip password",
										Computed:            true,
									},
									"peer_authentication_password_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"keepalive_interval": schema.Int64Attribute{
										MarkdownDescription: "Set MSDP peer ip keepalive interval",
										Computed:            true,
									},
									"keepalive_interval_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"keepalive_hold_time": schema.Int64Attribute{
										MarkdownDescription: "Set MSDP peer ip keepalive hold time",
										Computed:            true,
									},
									"keepalive_hold_time_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"sa_limit": schema.Int64Attribute{
										MarkdownDescription: "Set MSDP peer ip SA limit message number",
										Computed:            true,
									},
									"sa_limit_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"default_peer": schema.BoolAttribute{
										MarkdownDescription: "Set MSDP default peer",
										Computed:            true,
									},
									"prefix_list_id": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"msdp_originator_id": schema.StringAttribute{
				MarkdownDescription: "Set MSDP originator ID",
				Computed:            true,
			},
			"msdp_originator_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"msdp_connection_retry_interval": schema.Int64Attribute{
				MarkdownDescription: "Set MSDP refresh timer",
				Computed:            true,
			},
			"msdp_connection_retry_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
		},
	}
}

func (d *ServiceMulticastProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *ServiceMulticastProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ServiceMulticast

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
