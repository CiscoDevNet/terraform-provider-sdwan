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
	_ datasource.DataSource              = &SystemIPv4DeviceAccessProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &SystemIPv4DeviceAccessProfileParcelDataSource{}
)

func NewSystemIPv4DeviceAccessProfileParcelDataSource() datasource.DataSource {
	return &SystemIPv4DeviceAccessProfileParcelDataSource{}
}

type SystemIPv4DeviceAccessProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *SystemIPv4DeviceAccessProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_ipv4_device_access_feature"
}

func (d *SystemIPv4DeviceAccessProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the System IPv4 Device Access Feature.",

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
				MarkdownDescription: "Device Access Control List",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: "Sequence Id",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Sequence Name",
							Computed:            true,
						},
						"base_action": schema.StringAttribute{
							MarkdownDescription: "Base Action",
							Computed:            true,
						},
						"device_access_port": schema.Int64Attribute{
							MarkdownDescription: "device access protocol",
							Computed:            true,
						},
						"destination_data_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"destination_ip_prefix_list": schema.SetAttribute{
							MarkdownDescription: "Destination Data IP Prefix List",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"destination_ip_prefix_list_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_ports": schema.SetAttribute{
							MarkdownDescription: "Source Port List",
							ElementType:         types.Int64Type,
							Computed:            true,
						},
						"source_data_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"source_ip_prefix_list": schema.SetAttribute{
							MarkdownDescription: "Source Data IP Prefix List",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"source_ip_prefix_list_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SystemIPv4DeviceAccessProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *SystemIPv4DeviceAccessProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SystemIPv4DeviceAccess

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
