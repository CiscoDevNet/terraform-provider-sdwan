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
	_ datasource.DataSource              = &DNSSecurityProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &DNSSecurityProfileParcelDataSource{}
)

func NewDNSSecurityProfileParcelDataSource() datasource.DataSource {
	return &DNSSecurityProfileParcelDataSource{}
}

type DNSSecurityProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *DNSSecurityProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_security_policy"
}

func (d *DNSSecurityProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the DNS Security Policy.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Policy",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Policy",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Policy",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Policy",
				Computed:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: "Feature Profile ID",
				Required:            true,
			},
			"local_domain_bypass_list_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"match_all_vpn": schema.BoolAttribute{
				MarkdownDescription: "If this is true, there shouldn't be a targetVpns field, if false then targetVpns field should be there",
				Computed:            true,
			},
			"umbrella_default": schema.BoolAttribute{
				MarkdownDescription: "Field will only be under data field if matchAllVpn is true, otherwise field will be under targetVpns and set per entry",
				Computed:            true,
			},
			"dns_server_ip": schema.StringAttribute{
				MarkdownDescription: "Field will only be under data field if matchAllVpn is true, otherwise field will be under targetVpns and set per entry",
				Computed:            true,
			},
			"local_domain_bypass_enabled": schema.BoolAttribute{
				MarkdownDescription: "Field will only be under data field if matchAllVpn is true, otherwise field will be under targetVpns and set per entry",
				Computed:            true,
			},
			"dns_crypt": schema.BoolAttribute{
				MarkdownDescription: "If matchAllVpn is false, this field is only true if at least one of the targetVpns entires contains an umbrellaDefault true",
				Computed:            true,
			},
			"child_org_id": schema.StringAttribute{
				MarkdownDescription: "String that is a number that corresponds to Umbrella Multi Org, can be empty if not using Umbrella Multi Org",
				Computed:            true,
			},
			"target_vpns": schema.ListNestedAttribute{
				MarkdownDescription: "Will be under data field only if matchAllVpn is false, if matchAllVpn is true field should not be in payload",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpns": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"umbrella_default": schema.BoolAttribute{
							MarkdownDescription: "Field will only be under data field if matchAllVpn is true, otherwise field will be under targetVpns and set per entry",
							Computed:            true,
						},
						"dns_server_ip": schema.StringAttribute{
							MarkdownDescription: "Field will only be under data field if matchAllVpn is true, otherwise field will be under targetVpns and set per entry",
							Computed:            true,
						},
						"local_domain_bypass_enabled": schema.BoolAttribute{
							MarkdownDescription: "Field will only be under data field if matchAllVpn is true, otherwise field will be under targetVpns and set per entry",
							Computed:            true,
						},
						"uid": schema.StringAttribute{
							MarkdownDescription: "non empty interger string",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *DNSSecurityProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *DNSSecurityProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DNSSecurity

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
