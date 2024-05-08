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
	_ datasource.DataSource              = &EigrpFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &EigrpFeatureTemplateDataSource{}
)

func NewEigrpFeatureTemplateDataSource() datasource.DataSource {
	return &EigrpFeatureTemplateDataSource{}
}

type EigrpFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *EigrpFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_eigrp_feature_template"
}

func (d *EigrpFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the eigrp feature template.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the feature template",
				Optional:            true,
				Computed:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the feature template",
				Computed:            true,
			},
			"template_type": schema.StringAttribute{
				MarkdownDescription: "The template type",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the feature template",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the feature template",
				Computed:            true,
			},
			"device_types": schema.SetAttribute{
				MarkdownDescription: "List of supported device types",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"as_number": schema.Int64Attribute{
				MarkdownDescription: "Set autonomous system ID <1..65535>",
				Computed:            true,
			},
			"as_number_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"address_families": schema.ListNestedAttribute{
				MarkdownDescription: "Set EIGRP address family",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Set EIGRP address family",
							Computed:            true,
						},
						"redistributes": schema.ListNestedAttribute{
							MarkdownDescription: "Redistribute routes into EIGRP",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: "Set the protocol to redistribute routes from",
										Computed:            true,
									},
									"protocol_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"route_policy": schema.StringAttribute{
										MarkdownDescription: "Configure policy to apply to prefixes received from EIGRP neighbor",
										Computed:            true,
									},
									"route_policy_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Computed:            true,
									},
								},
							},
						},
						"networks": schema.ListNestedAttribute{
							MarkdownDescription: "Configure the networks for EIGRP to advertise",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix": schema.StringAttribute{
										MarkdownDescription: "Configure the prefixes for EIGRP to announce",
										Computed:            true,
									},
									"prefix_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Computed:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"hello_interval": schema.Int64Attribute{
				MarkdownDescription: "Set EIGRP hello interval",
				Computed:            true,
			},
			"hello_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"hold_time": schema.Int64Attribute{
				MarkdownDescription: "Set EIGRP hold time",
				Computed:            true,
			},
			"hold_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"route_policy_name": schema.StringAttribute{
				MarkdownDescription: "Configure policy to apply to prefixes received from EIGRP neighbor",
				Computed:            true,
			},
			"route_policy_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"filter": schema.BoolAttribute{
				MarkdownDescription: "Selective route download",
				Computed:            true,
			},
			"filter_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"authentication_type": schema.StringAttribute{
				MarkdownDescription: "Set EIGRP router authentication type",
				Computed:            true,
			},
			"authentication_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"hmac_authentication_key": schema.StringAttribute{
				MarkdownDescription: "Set hmac-sha-256 authentication key",
				Computed:            true,
			},
			"hmac_authentication_key_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"keys": schema.ListNestedAttribute{
				MarkdownDescription: "Set keychain name",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"md5_key_id": schema.Int64Attribute{
							MarkdownDescription: "Set MD5 key ID",
							Computed:            true,
						},
						"md5_key_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"md5_authentication_key": schema.StringAttribute{
							MarkdownDescription: "Set MD5 key",
							Computed:            true,
						},
						"md5_authentication_key_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "Configure IPv4 Static Routes",
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
						"shutdown": schema.BoolAttribute{
							MarkdownDescription: "Enable/disable EIGRP",
							Computed:            true,
						},
						"shutdown_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"summary_addresses": schema.ListNestedAttribute{
							MarkdownDescription: "Set summary addresses",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix": schema.StringAttribute{
										MarkdownDescription: "Set summary address prefix",
										Computed:            true,
									},
									"prefix_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Computed:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *EigrpFeatureTemplateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *EigrpFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *EigrpFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Eigrp

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	if config.Id.IsNull() && !config.Name.IsNull() {
		res, err := d.client.Get("/template/feature")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve existing templates, got error: %s", err))
			return
		}
		if value := res.Get("data"); len(value.Array()) > 0 {
			value.ForEach(func(k, v gjson.Result) bool {
				if config.Name.ValueString() == v.Get("templateName").String() {
					config.Id = types.StringValue(v.Get("templateId").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found feature template with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
					return false
				}
				return true
			})
		}
		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find feature template with name: %s", config.Name.ValueString()))
			return
		}
	}

	res, err := d.client.Get("/template/feature/object/" + url.QueryEscape(config.Id.ValueString()))
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
