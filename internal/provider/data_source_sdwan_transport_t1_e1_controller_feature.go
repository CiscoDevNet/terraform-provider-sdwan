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
	_ datasource.DataSource              = &TransportT1E1ControllerProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &TransportT1E1ControllerProfileParcelDataSource{}
)

func NewTransportT1E1ControllerProfileParcelDataSource() datasource.DataSource {
	return &TransportT1E1ControllerProfileParcelDataSource{}
}

type TransportT1E1ControllerProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *TransportT1E1ControllerProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transport_t1_e1_controller_feature"
}

func (d *TransportT1E1ControllerProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Transport T1 E1 Controller Feature.",

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
			"type": schema.StringAttribute{
				MarkdownDescription: "Card Type",
				Computed:            true,
			},
			"slot": schema.StringAttribute{
				MarkdownDescription: "Slot number",
				Computed:            true,
			},
			"slot_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"entries": schema.ListNestedAttribute{
				MarkdownDescription: "Controller tx-ex List",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"t1_description": schema.StringAttribute{
							MarkdownDescription: "Card Type",
							Computed:            true,
						},
						"t1_framing": schema.StringAttribute{
							MarkdownDescription: "Framing",
							Computed:            true,
						},
						"t1_framing_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"t1_linecode": schema.StringAttribute{
							MarkdownDescription: "LineCode",
							Computed:            true,
						},
						"t1_linecode_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"e1_description": schema.StringAttribute{
							MarkdownDescription: "Card Type",
							Computed:            true,
						},
						"e1_framing": schema.StringAttribute{
							MarkdownDescription: "Framing",
							Computed:            true,
						},
						"e1_framing_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"e1_linecode": schema.StringAttribute{
							MarkdownDescription: "LineCode",
							Computed:            true,
						},
						"e1_linecode_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"cable_length": schema.StringAttribute{
							MarkdownDescription: "Cable Config",
							Computed:            true,
						},
						"length_short": schema.StringAttribute{
							MarkdownDescription: "length",
							Computed:            true,
						},
						"length_short_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"length_long": schema.StringAttribute{
							MarkdownDescription: "length",
							Computed:            true,
						},
						"length_long_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"clock_source": schema.StringAttribute{
							MarkdownDescription: "Clock Source",
							Computed:            true,
						},
						"line_mode": schema.StringAttribute{
							MarkdownDescription: "Line Mode",
							Computed:            true,
						},
						"line_mode_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Description",
							Computed:            true,
						},
						"description_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"channel_group": schema.ListNestedAttribute{
							MarkdownDescription: "Channel Group List",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"number": schema.Int64Attribute{
										MarkdownDescription: "Number",
										Computed:            true,
									},
									"number_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"timeslots": schema.StringAttribute{
										MarkdownDescription: "Time slots",
										Computed:            true,
									},
									"timeslots_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
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

func (d *TransportT1E1ControllerProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *TransportT1E1ControllerProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config TransportT1E1Controller

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
