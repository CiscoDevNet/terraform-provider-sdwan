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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ConfigurationGroupDeviceVariablesDataSource{}
	_ datasource.DataSourceWithConfigure = &ConfigurationGroupDeviceVariablesDataSource{}
)

func NewConfigurationGroupDeviceVariablesDataSource() datasource.DataSource {
	return &ConfigurationGroupDeviceVariablesDataSource{}
}

type ConfigurationGroupDeviceVariablesDataSource struct {
	client *sdwan.Client
}

func (d *ConfigurationGroupDeviceVariablesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_configuration_group_device_variables"
}

func (d *ConfigurationGroupDeviceVariablesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Configuration Group Device Variables .",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"configuration_group_id": schema.StringAttribute{
				MarkdownDescription: "Configuration Group ID",
				Required:            true,
			},
			"solution": schema.StringAttribute{
				MarkdownDescription: "Type of solution",
				Computed:            true,
			},
			"devices": schema.ListNestedAttribute{
				MarkdownDescription: "List of devices",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"device_id": schema.StringAttribute{
							MarkdownDescription: "Device ID",
							Computed:            true,
						},
						"variables": schema.ListNestedAttribute{
							MarkdownDescription: "List of variables",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Variable name",
										Computed:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Variable value",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"groups": schema.ListNestedAttribute{
				MarkdownDescription: "List of device groups",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Group name",
							Computed:            true,
						},
						"variables": schema.ListNestedAttribute{
							MarkdownDescription: "List of variables",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Variable name",
										Computed:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Variable value",
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

func (d *ConfigurationGroupDeviceVariablesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

func (d *ConfigurationGroupDeviceVariablesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ConfigurationGroupDeviceVariables

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.ConfigurationGroupId.String()))

	deviceIds := make([]string, 0)
	for d := range config.Devices {
		deviceIds = append(deviceIds, config.Devices[d].DeviceId.ValueString())
	}
	params := "?device-id=" + url.QueryEscape(strings.Join(deviceIds, ",")) + "&suggestions=false"

	res, err := d.client.Get(config.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)
	config.Id = config.ConfigurationGroupId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.ConfigurationGroupId.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
