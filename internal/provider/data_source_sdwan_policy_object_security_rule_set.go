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
	_ datasource.DataSource                     = &PolicyObjectSecurityRuleSetProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure        = &PolicyObjectSecurityRuleSetProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigValidators = &PolicyObjectSecurityRuleSetProfileParcelDataSource{}
)

func NewPolicyObjectSecurityRuleSetProfileParcelDataSource() datasource.DataSource {
	return &PolicyObjectSecurityRuleSetProfileParcelDataSource{}
}

type PolicyObjectSecurityRuleSetProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *PolicyObjectSecurityRuleSetProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_object_security_rule_set"
}

func (d *PolicyObjectSecurityRuleSetProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Policy Object Security Rule Set Policy_object.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Policy_object",
				Optional:            true,
				Computed:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Policy_object",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Policy_object",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Policy_object",
				Computed:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: "Feature Profile ID",
				Required:            true,
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: "select definition of rule",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Sequence name",
							Computed:            true,
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"sequence_ip_type": schema.StringAttribute{
							MarkdownDescription: "ipv4/ ipv6",
							Computed:            true,
						},
						"sequence_id": schema.StringAttribute{
							MarkdownDescription: "Sequence Id",
							Computed:            true,
						},
						"sequence_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"action": schema.StringAttribute{
							MarkdownDescription: "action",
							Computed:            true,
						},
						"action_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_ipv4_prefixes": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"source_ipv4_prefixes_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_ipv6_prefixes": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"source_ipv6_prefixes_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"destination_ipv4_prefixes": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"destination_ipv4_prefixes_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"destination_ipv6_prefixes": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"destination_ipv6_prefixes_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"destination_fqdns": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"destination_fqdns_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_geo_locations": schema.SetAttribute{
							MarkdownDescription: "Source Geolocation",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"source_geo_locations_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"destination_geo_locations": schema.SetAttribute{
							MarkdownDescription: "Destination Geolocation",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"destination_geo_locations_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_ports": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"source_ports_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"destination_ports": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"destination_ports_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"protocols": schema.SetAttribute{
							MarkdownDescription: "protocol",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"protocol_names": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"protocol_names_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_object_group_list_ids": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"destination_object_group_list_ids": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"source_data_ipv4_prefix_list_ids": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"source_data_ipv6_prefix_list_ids": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"destination_data_ipv4_prefix_list_ids": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"destination_data_ipv6_prefix_list_ids": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"destination_fqdn_list_ids": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"source_geo_location_list_ids": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"destination_geo_location_list_ids": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"source_port_list_ids": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"destination_port_list_ids": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"protocol_name_list_ids": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *PolicyObjectSecurityRuleSetProfileParcelDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *PolicyObjectSecurityRuleSetProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *PolicyObjectSecurityRuleSetProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config PolicyObjectSecurityRuleSet

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

	config.fromBody(ctx, res, true)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
