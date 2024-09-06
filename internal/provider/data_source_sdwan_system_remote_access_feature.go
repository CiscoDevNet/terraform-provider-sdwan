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
	_ datasource.DataSource              = &SystemRemoteAccessProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &SystemRemoteAccessProfileParcelDataSource{}
)

func NewSystemRemoteAccessProfileParcelDataSource() datasource.DataSource {
	return &SystemRemoteAccessProfileParcelDataSource{}
}

type SystemRemoteAccessProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *SystemRemoteAccessProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_remote_access_feature"
}

func (d *SystemRemoteAccessProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the System Remote Access Feature.",

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
			"connection_type_ssl": schema.BoolAttribute{
				MarkdownDescription: "Enabled SSL VPN",
				Computed:            true,
			},
			"any_connect_eap_authentication_type": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ipv4_pool_size": schema.Int64Attribute{
				MarkdownDescription: "IPv4 Pool Size",
				Computed:            true,
			},
			"ipv4_pool_size_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_pool_size": schema.Int64Attribute{
				MarkdownDescription: "IPv6 Pool Size",
				Computed:            true,
			},
			"ipv6_pool_size_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"enable_certificate_list_check": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enable_certificate_list_check_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"psk_authentication_type": schema.StringAttribute{
				MarkdownDescription: "PSK Selection",
				Computed:            true,
			},
			"psk_authentication_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"psk_authentication_pre_shared_key": schema.StringAttribute{
				MarkdownDescription: "PSK Pre Shared Key",
				Computed:            true,
			},
			"psk_authentication_pre_shared_key_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"radius_group_name": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"radius_group_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"aaa_specify_name_policy_name": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"aaa_specify_name_policy_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"aaa_specify_name_policy_password": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"aaa_specify_name_policy_password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"aaa_derive_name_from_peer_identity": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"aaa_derive_name_from_peer_identity_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"aaa_derive_name_from_peer_domain": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"aaa_derive_name_from_peer_domain_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"aaa_enable_accounting": schema.BoolAttribute{
				MarkdownDescription: "Enable Accounting",
				Computed:            true,
			},
			"aaa_enable_accounting_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ikev2_local_ike_identity_type": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ikev2_local_ike_identity_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ikev2_local_ike_identity_value": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ikev2_local_ike_identity_value_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ikev2_security_association_lifetime": schema.Int64Attribute{
				MarkdownDescription: "Security Association Lifetime in Seconds",
				Computed:            true,
			},
			"ikev2_security_association_lifetime_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ikev2_anti_dos_threshold": schema.Int64Attribute{
				MarkdownDescription: "Anti-DOS Threshold",
				Computed:            true,
			},
			"ikev2_anti_dos_threshold_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipsec_enable_anti_replay": schema.BoolAttribute{
				MarkdownDescription: "Enable Anti-Replay",
				Computed:            true,
			},
			"ipsec_enable_anti_replay_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipsec_anti_replay_window_size": schema.Int64Attribute{
				MarkdownDescription: "security Association Lifetime",
				Computed:            true,
			},
			"ipsec_anti_replay_window_size_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipsec_security_association_lifetime": schema.Int64Attribute{
				MarkdownDescription: "Security Association Lifetime in Seconds",
				Computed:            true,
			},
			"ipsec_security_association_lifetime_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipsec_enable_perfect_foward_secrecy": schema.BoolAttribute{
				MarkdownDescription: "security Association Lifetime",
				Computed:            true,
			},
			"ipsec_enable_perfect_foward_secrecy_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
		},
	}
}

func (d *SystemRemoteAccessProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *SystemRemoteAccessProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SystemRemoteAccess

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
