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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ConfigurationGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &ConfigurationGroupDataSource{}
)

func NewConfigurationGroupDataSource() datasource.DataSource {
	return &ConfigurationGroupDataSource{}
}

type ConfigurationGroupDataSource struct {
	client *sdwan.Client
}

func (d *ConfigurationGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_configuration_group"
}

func (d *ConfigurationGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Configuration Group .",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the configuration group",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description",
				Computed:            true,
			},
			"solution": schema.StringAttribute{
				MarkdownDescription: "Type of solution",
				Computed:            true,
			},
			"feature_profiles": schema.SetNestedAttribute{
				MarkdownDescription: "List of feature profiles",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Feature profile ID",
							Computed:            true,
						},
					},
				},
			},
			"topology_devices": schema.ListNestedAttribute{
				MarkdownDescription: "List of topology device types",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"criteria_attribute": schema.StringAttribute{
							MarkdownDescription: "Criteria attribute",
							Computed:            true,
						},
						"criteria_value": schema.StringAttribute{
							MarkdownDescription: "Criteria value",
							Computed:            true,
						},
						"unsupported_features": schema.ListNestedAttribute{
							MarkdownDescription: "List of unsupported features",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"parcel_type": schema.StringAttribute{
										MarkdownDescription: "Parcel type",
										Computed:            true,
									},
									"parcel_id": schema.StringAttribute{
										MarkdownDescription: "Parcel ID",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"topology_site_devices": schema.Int64Attribute{
				MarkdownDescription: "Number of devices per site",
				Computed:            true,
			},
			"devices": schema.ListNestedAttribute{
				MarkdownDescription: "List of devices",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Device ID",
							Computed:            true,
						},
						"deploy": schema.BoolAttribute{
							MarkdownDescription: "Deploy to device if enabled.",
							Computed:            true,
						},
						"variables": schema.SetNestedAttribute{
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
									"list_value": schema.ListAttribute{
										MarkdownDescription: "Use this instead of `value` in case value is of type `List`.",
										ElementType:         types.StringType,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"feature_versions": schema.ListAttribute{
				MarkdownDescription: "List of all associated feature versions",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}

func (d *ConfigurationGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

func (d *ConfigurationGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ConfigurationGroup

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	// Read config group
	res, err := d.client.Get(config.getPath() + url.QueryEscape(config.Id.ValueString()))
	if strings.Contains(res.Get("error.message").String(), "Invalid config group passed") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	config.fromBodyConfigGroup(ctx, res)

	// Read config group devices
	path := fmt.Sprintf("/v1/config-group/%v/device/associate/", config.Id.ValueString())
	res, err = d.client.Get(path)
	if strings.Contains(res.Get("error.message").String(), "Invalid config group passed") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	config.fromBodyConfigGroupDevices(ctx, res)

	// Read config group devices
	path = fmt.Sprintf("/v1/config-group/%v/device/variables/", config.Id.ValueString())
	res, err = d.client.Get(path)
	if strings.Contains(res.Get("error.message").String(), "Invalid config group passed") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	config.fromBodyConfigGroupDeviceVariables(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
