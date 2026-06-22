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
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource                     = &OtherTrustSecProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure        = &OtherTrustSecProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigValidators = &OtherTrustSecProfileParcelDataSource{}
)

func NewOtherTrustSecProfileParcelDataSource() datasource.DataSource {
	return &OtherTrustSecProfileParcelDataSource{}
}

type OtherTrustSecProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *OtherTrustSecProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_other_trustsec_feature"
}

func (d *OtherTrustSecProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Other TrustSec Feature.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Feature",
				Optional:            true,
				Computed:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Feature",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Feature",
				Optional:            true,
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
			"device_id": schema.StringAttribute{
				MarkdownDescription: "Specify the TrustSec Network Access Device ID, should be same as mentioned in the Identity Services Engine (upto 32 char)",
				Computed:            true,
			},
			"device_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"device_password": schema.StringAttribute{
				MarkdownDescription: "Set the password for the device",
				Computed:            true,
			},
			"device_password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"device_sgt": schema.Int64Attribute{
				MarkdownDescription: "Configure Local device security group <2..65519>",
				Computed:            true,
			},
			"device_sgt_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"enable_enforcement": schema.BoolAttribute{
				MarkdownDescription: "Enable Role-based Access Control enforcement",
				Computed:            true,
			},
			"enable_enforcement_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"enable_sxp": schema.BoolAttribute{
				MarkdownDescription: "Enable CTS SXP support",
				Computed:            true,
			},
			"listener_hold_time_min": schema.Int64Attribute{
				MarkdownDescription: "Configure Minimum allowed hold-time for listener in seconds <1..65534>",
				Computed:            true,
			},
			"listener_hold_time_min_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"listener_hold_time_max": schema.Int64Attribute{
				MarkdownDescription: "Configure Maximum allowed hold-time for listener in seconds <1..65534>",
				Computed:            true,
			},
			"listener_hold_time_max_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"speaker_hold_time": schema.Int64Attribute{
				MarkdownDescription: "Configure Speaker hold-time in seconds <1..65534>",
				Computed:            true,
			},
			"speaker_hold_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"sxp_default_password": schema.StringAttribute{
				MarkdownDescription: "Configure SXP default password",
				Computed:            true,
			},
			"sxp_default_password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"sxp_key_chain": schema.StringAttribute{
				MarkdownDescription: "Configure SXP key-chain",
				Computed:            true,
			},
			"sxp_key_chain_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"sxp_log_binding_changes": schema.BoolAttribute{
				MarkdownDescription: "Enables logging for IP-to-SGT binding changes",
				Computed:            true,
			},
			"sxp_log_binding_changes_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"sxp_reconciliation_period": schema.Int64Attribute{
				MarkdownDescription: "Configure the SXP reconciliation period in seconds <0..64000>",
				Computed:            true,
			},
			"sxp_reconciliation_period_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"sxp_retry_period": schema.Int64Attribute{
				MarkdownDescription: "Configure Retry period for SXP connection in seconds <0..64000>",
				Computed:            true,
			},
			"sxp_retry_period_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"sxp_source_ip": schema.StringAttribute{
				MarkdownDescription: "SXP Source IP",
				Computed:            true,
			},
			"sxp_source_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"sxp_connections": schema.ListNestedAttribute{
				MarkdownDescription: "Configure SXP Connections",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"peer_ip": schema.StringAttribute{
							MarkdownDescription: "Configure SXP Peer IP address (IPv4)",
							Computed:            true,
						},
						"peer_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_ip": schema.StringAttribute{
							MarkdownDescription: "Configure SXP Source IP address (IPv4)",
							Computed:            true,
						},
						"source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"preshared_key": schema.StringAttribute{
							MarkdownDescription: "Define Preshared Key type",
							Computed:            true,
						},
						"mode": schema.StringAttribute{
							MarkdownDescription: "Define Mode of connection",
							Computed:            true,
						},
						"mode_type": schema.StringAttribute{
							MarkdownDescription: "Define Role of a device <speaker/listener/both>",
							Computed:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "Configure Connection VPN (VRF) ID",
							Computed:            true,
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"min_hold_time": schema.Int64Attribute{
							MarkdownDescription: "Configure Connection Minimum hold time <0..65535>",
							Computed:            true,
						},
						"min_hold_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"max_hold_time": schema.Int64Attribute{
							MarkdownDescription: "Configure Connection Maximum hold time <0..65535>",
							Computed:            true,
						},
						"max_hold_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OtherTrustSecProfileParcelDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *OtherTrustSecProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *OtherTrustSecProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config OtherTrustSec

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))
	if config.Id.IsNull() && !config.Name.IsNull() {
		// Look up parcel ID by name
		res, err := d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve parcels, got error: %s", err))
			return
		}
		found := false
		res.Get("data").ForEach(func(_, v gjson.Result) bool {
			if v.Get("payload.name").String() == config.Name.ValueString() {
				config.Id = types.StringValue(v.Get("parcelId").String())
				found = true
				return false
			}
			return true
		})
		if !found {
			resp.Diagnostics.AddError("Not Found", fmt.Sprintf("No parcel found with name: %s", config.Name.ValueString()))
			return
		}
	}

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
