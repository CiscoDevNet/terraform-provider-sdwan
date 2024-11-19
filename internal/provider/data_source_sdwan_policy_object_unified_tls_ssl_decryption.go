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
	_ datasource.DataSource              = &PolicyObjectUnifiedTLSSSLDecryptionProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &PolicyObjectUnifiedTLSSSLDecryptionProfileParcelDataSource{}
)

func NewPolicyObjectUnifiedTLSSSLDecryptionProfileParcelDataSource() datasource.DataSource {
	return &PolicyObjectUnifiedTLSSSLDecryptionProfileParcelDataSource{}
}

type PolicyObjectUnifiedTLSSSLDecryptionProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *PolicyObjectUnifiedTLSSSLDecryptionProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_object_unified_tls_ssl_decryption"
}

func (d *PolicyObjectUnifiedTLSSSLDecryptionProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Policy Object Unified TLS SSL Decryption Policy_object.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Policy_object",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Policy_object",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Policy_object",
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
			"expired_certificate": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"untrusted_certificate": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"certificate_revocation_status": schema.StringAttribute{
				MarkdownDescription: "If value is none unknown status not required, if value is ocsp then unknown status is required",
				Computed:            true,
			},
			"unknown_revocation_status": schema.StringAttribute{
				MarkdownDescription: "Only required if certificateRevocationStatus is oscp, if value is none then field shouldn't be here",
				Computed:            true,
			},
			"unsupported_protocol_versions": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"unsupported_cipher_suites": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"failure_mode": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"default_ca_certificate_bundle": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"file_name": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"bundle_string": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"rsa_keypair_modules": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ec_key_type": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"certificate_lifetime": schema.StringAttribute{
				MarkdownDescription: "If you have vManage as CA or vManage as intermediate CA, this value should be 1",
				Computed:            true,
			},
			"minimal_tls_ver": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
		},
	}
}

func (d *PolicyObjectUnifiedTLSSSLDecryptionProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *PolicyObjectUnifiedTLSSSLDecryptionProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config PolicyObjectUnifiedTLSSSLDecryption

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
