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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &PolicyObjectFeatureProfileParcelsDataSource{}
	_ datasource.DataSourceWithConfigure = &PolicyObjectFeatureProfileParcelsDataSource{}
)

func NewPolicyObjectFeatureProfileParcelsDataSource() datasource.DataSource {
	return &PolicyObjectFeatureProfileParcelsDataSource{}
}

type PolicyObjectFeatureProfileParcelsDataSource struct {
	client *sdwan.Client
}

func (d *PolicyObjectFeatureProfileParcelsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_object_feature_profile_parcels"
}

func (d *PolicyObjectFeatureProfileParcelsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the parcels of a Policy Object Feature Profile.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: "Policy Object Feature Profile ID",
				Required:            true,
			},
			"parcels": schema.ListNestedAttribute{
				MarkdownDescription: "List of parcels associated with this profile",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"parcel_id": schema.StringAttribute{
							MarkdownDescription: "Parcel identifier",
							Computed:            true,
						},
						"parcel_type": schema.StringAttribute{
							MarkdownDescription: "Parcel type (e.g. app-list, data-prefix, etc.)",
							Computed:            true,
						},
						"created_by": schema.StringAttribute{
							MarkdownDescription: "User who created the parcel (system for system-created parcels)",
							Computed:            true,
						},
						"last_updated_by": schema.StringAttribute{
							MarkdownDescription: "User who last updated the parcel",
							Computed:            true,
						},
						"parcel_name": schema.StringAttribute{
							MarkdownDescription: "Name of the parcel",
							Computed:            true,
						},
						"parcel_description": schema.StringAttribute{
							MarkdownDescription: "Description of the parcel",
							Computed:            true,
						},
						"reference_count": schema.Int64Attribute{
							MarkdownDescription: "Number of references to this parcel",
							Computed:            true,
						},
					},
				},
			},
			"created_by": schema.StringAttribute{
				MarkdownDescription: "Filter by creator (e.g. 'system' for system-created parcels)",
				Optional:            true,
			},
			"parcel_type": schema.StringAttribute{
				MarkdownDescription: "Filter by parcel type (e.g. 'app-list', 'data-prefix')",
				Optional:            true,
			},
		},
	}
}

func (d *PolicyObjectFeatureProfileParcelsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *PolicyObjectFeatureProfileParcelsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config PolicyObjectFeatureProfileParcels

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var params = "?"

	res, err := d.client.Get(config.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	// Apply client-side filters to list results
	config.applyFilters(ctx)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
