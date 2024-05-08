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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SystemSecurityProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &SystemSecurityProfileParcelDataSource{}
)

func NewSystemSecurityProfileParcelDataSource() datasource.DataSource {
	return &SystemSecurityProfileParcelDataSource{}
}

type SystemSecurityProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *SystemSecurityProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_security_profile_parcel"
}

func (d *SystemSecurityProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the System Security profile parcel.",

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
			"rekey": schema.Int64Attribute{
				MarkdownDescription: "Set how often to change the AES key for DTLS connections",
				Computed:            true,
			},
			"rekey_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"anti_replay_window": schema.StringAttribute{
				MarkdownDescription: "Set the sliding replay window size",
				Computed:            true,
			},
			"anti_replay_window_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"extended_anti_replay_window": schema.Int64Attribute{
				MarkdownDescription: "Extended Anti-Replay Window",
				Computed:            true,
			},
			"extended_anti_replay_window_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipsec_pairwise_keying": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable IPsec pairwise-keying",
				Computed:            true,
			},
			"ipsec_pairwise_keying_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"integrity_type": schema.SetAttribute{
				MarkdownDescription: "Set the authentication type for DTLS connections",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"integrity_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"keychains": schema.ListNestedAttribute{
				MarkdownDescription: "Configure a Keychain",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key_chain_name": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the Keychain",
							Computed:            true,
						},
						"key_id": schema.Int64Attribute{
							MarkdownDescription: "Specify the Key ID",
							Computed:            true,
						},
					},
				},
			},
			"keys": schema.ListNestedAttribute{
				MarkdownDescription: "Configure a Key",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: "Select the Key ID",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Select the chain name",
							Computed:            true,
						},
						"send_id": schema.Int64Attribute{
							MarkdownDescription: "Specify the Send ID",
							Computed:            true,
						},
						"send_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"receiver_id": schema.Int64Attribute{
							MarkdownDescription: "Specify the Receiver ID",
							Computed:            true,
						},
						"receiver_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"include_tcp_options": schema.BoolAttribute{
							MarkdownDescription: "Configure Include TCP Options",
							Computed:            true,
						},
						"include_tcp_options_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"accept_ao_mismatch": schema.BoolAttribute{
							MarkdownDescription: "Configure Accept AO Mismatch",
							Computed:            true,
						},
						"accept_ao_mismatch_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"crypto_algorithm": schema.StringAttribute{
							MarkdownDescription: "Crypto Algorithm",
							Computed:            true,
						},
						"key_string": schema.StringAttribute{
							MarkdownDescription: "Specify the Key String",
							Computed:            true,
						},
						"key_string_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"send_life_time_local": schema.BoolAttribute{
							MarkdownDescription: "Configure Send lifetime Local",
							Computed:            true,
						},
						"send_life_time_local_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"send_life_time_start_epoch": schema.Int64Attribute{
							MarkdownDescription: "Configure Key lifetime start time",
							Computed:            true,
						},
						"send_life_time_infinite": schema.BoolAttribute{
							MarkdownDescription: "Infinite lifetime",
							Computed:            true,
						},
						"send_life_time_infinite_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"send_life_time_duration": schema.Int64Attribute{
							MarkdownDescription: "Send lifetime Duration (seconds)",
							Computed:            true,
						},
						"send_life_time_duration_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"send_life_time_exact": schema.Int64Attribute{
							MarkdownDescription: "Configure Key lifetime end time",
							Computed:            true,
						},
						"accept_life_time_local": schema.BoolAttribute{
							MarkdownDescription: "Configure Send lifetime Local",
							Computed:            true,
						},
						"accept_life_time_local_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"accept_life_time_start_epoch": schema.Int64Attribute{
							MarkdownDescription: "Configure Key lifetime start time",
							Computed:            true,
						},
						"accept_life_time_infinite": schema.BoolAttribute{
							MarkdownDescription: "Infinite lifetime",
							Computed:            true,
						},
						"accept_life_time_infinite_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"accept_life_time_duration": schema.Int64Attribute{
							MarkdownDescription: "Send lifetime Duration (seconds)",
							Computed:            true,
						},
						"accept_life_time_duration_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"accept_life_time_exact": schema.Int64Attribute{
							MarkdownDescription: "Configure Key lifetime end time",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SystemSecurityProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *SystemSecurityProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SystemSecurity

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
