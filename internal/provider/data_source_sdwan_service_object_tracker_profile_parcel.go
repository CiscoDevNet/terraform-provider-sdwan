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
	_ datasource.DataSource              = &ServiceObjectTrackerProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &ServiceObjectTrackerProfileParcelDataSource{}
)

func NewServiceObjectTrackerProfileParcelDataSource() datasource.DataSource {
	return &ServiceObjectTrackerProfileParcelDataSource{}
}

type ServiceObjectTrackerProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *ServiceObjectTrackerProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_object_tracker_profile_parcel"
}

func (d *ServiceObjectTrackerProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Service Object Tracker profile parcel.",

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
			"object_tracker_id": schema.Int64Attribute{
				MarkdownDescription: "Object tracker ID",
				Computed:            true,
			},
			"object_tracker_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"object_tracker_type": schema.StringAttribute{
				MarkdownDescription: "objectTrackerType:Interface SIG Route",
				Computed:            true,
			},
			"interface": schema.StringAttribute{
				MarkdownDescription: "interface name",
				Computed:            true,
			},
			"interface_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"route_ip": schema.StringAttribute{
				MarkdownDescription: "IP address",
				Computed:            true,
			},
			"route_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"route_mask": schema.StringAttribute{
				MarkdownDescription: "IP mask",
				Computed:            true,
			},
			"route_mask_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"vpn": schema.Int64Attribute{
				MarkdownDescription: "VPN",
				Computed:            true,
			},
			"vpn_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
		},
	}
}

func (d *ServiceObjectTrackerProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *ServiceObjectTrackerProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ServiceObjectTracker

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
