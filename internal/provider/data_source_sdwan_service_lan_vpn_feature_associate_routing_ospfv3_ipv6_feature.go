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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ServiceLANVPNFeatureAssociateRoutingOSPFv3IPv6FeatureDataSource{}
	_ datasource.DataSourceWithConfigure = &ServiceLANVPNFeatureAssociateRoutingOSPFv3IPv6FeatureDataSource{}
)

func NewServiceLANVPNFeatureAssociateRoutingOSPFv3IPv6FeatureDataSource() datasource.DataSource {
	return &ServiceLANVPNFeatureAssociateRoutingOSPFv3IPv6FeatureDataSource{}
}

type ServiceLANVPNFeatureAssociateRoutingOSPFv3IPv6FeatureDataSource struct {
	client *sdwan.Client
}

func (d *ServiceLANVPNFeatureAssociateRoutingOSPFv3IPv6FeatureDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_lan_vpn_feature_associate_routing_ospfv3_ipv6_feature"
}

func (d *ServiceLANVPNFeatureAssociateRoutingOSPFv3IPv6FeatureDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Service LAN VPN Feature Associate Routing OSPFv3 IPv6 Feature .",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: "Feature Profile ID",
				Required:            true,
			},
			"service_lan_vpn_feature_id": schema.StringAttribute{
				MarkdownDescription: "Service LAN VPN Feature ID",
				Required:            true,
			},
			"service_routing_ospfv3_ipv6_feature_id": schema.StringAttribute{
				MarkdownDescription: "Service Routing OSPFv3 IPv6 Feature ID",
				Computed:            true,
			},
		},
	}
}

func (d *ServiceLANVPNFeatureAssociateRoutingOSPFv3IPv6FeatureDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *ServiceLANVPNFeatureAssociateRoutingOSPFv3IPv6FeatureDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ServiceLANVPNFeatureAssociateRoutingOSPFv3IPv6Feature

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get(config.getPath() + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
