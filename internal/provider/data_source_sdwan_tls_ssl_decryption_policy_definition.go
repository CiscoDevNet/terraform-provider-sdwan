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
	_ datasource.DataSource              = &TLSSSLDecryptionPolicyDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &TLSSSLDecryptionPolicyDefinitionDataSource{}
)

func NewTLSSSLDecryptionPolicyDefinitionDataSource() datasource.DataSource {
	return &TLSSSLDecryptionPolicyDefinitionDataSource{}
}

type TLSSSLDecryptionPolicyDefinitionDataSource struct {
	client *sdwan.Client
}

func (d *TLSSSLDecryptionPolicyDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tls_ssl_decryption_policy_definition"
}

func (d *TLSSSLDecryptionPolicyDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the TLS SSL Decryption Policy Definition .",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the policy definition.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the policy definition.",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "The policy mode",
				Computed:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: "Default action (applies when `mode` set to `security`)",
				Computed:            true,
			},
			"network_rules": schema.ListNestedAttribute{
				MarkdownDescription: "List of network rules (applies when `mode` set to `security`)",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"base_action": schema.StringAttribute{
							MarkdownDescription: "Rule base action",
							Computed:            true,
						},
						"rule_id": schema.Int64Attribute{
							MarkdownDescription: "Rule ID",
							Computed:            true,
						},
						"rule_name": schema.StringAttribute{
							MarkdownDescription: "Rule name",
							Computed:            true,
						},
						"rule_type": schema.StringAttribute{
							MarkdownDescription: "Rule type",
							Computed:            true,
						},
						"source_and_destination_configuration": schema.ListNestedAttribute{
							MarkdownDescription: "List of network source / destination configuration",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"option": schema.StringAttribute{
										MarkdownDescription: "source / destination option",
										Computed:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "source / destination option target",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"url_rules": schema.ListNestedAttribute{
				MarkdownDescription: "List of url rules (applies when `mode` set to `security`)",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"rule_name": schema.StringAttribute{
							MarkdownDescription: "Country",
							Computed:            true,
						},
						"target_vpns": schema.SetAttribute{
							MarkdownDescription: "List of VPN IDs",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"tls_ssl_profile_policy_id": schema.StringAttribute{
							MarkdownDescription: "TLS SSL Profile Policy ID",
							Computed:            true,
						},
						"tls_ssl_profile_policy_version": schema.Int64Attribute{
							MarkdownDescription: "TLS SSL Profile Policy version",
							Computed:            true,
						},
					},
				},
			},
			"ssl_decryption_enabled": schema.StringAttribute{
				MarkdownDescription: "SSL decryption enabled",
				Computed:            true,
			},
			"expired_certificate": schema.StringAttribute{
				MarkdownDescription: "Expired certificate action",
				Computed:            true,
			},
			"untrusted_certificate": schema.StringAttribute{
				MarkdownDescription: "Untrusted certificate action",
				Computed:            true,
			},
			"certificate_revocation_status": schema.StringAttribute{
				MarkdownDescription: "Certificate revocation status",
				Computed:            true,
			},
			"unknown_revocation_status": schema.StringAttribute{
				MarkdownDescription: "Unknown revocation status action",
				Computed:            true,
			},
			"unsupported_protocol_versions": schema.StringAttribute{
				MarkdownDescription: "Unsupported protocol versions action",
				Computed:            true,
			},
			"unsupported_cipher_suites": schema.StringAttribute{
				MarkdownDescription: "Unsupported cipher suites action",
				Computed:            true,
			},
			"failure_mode": schema.StringAttribute{
				MarkdownDescription: "Failure mode",
				Computed:            true,
			},
			"rsa_key_pair_modulus": schema.StringAttribute{
				MarkdownDescription: "RSA key pair modules",
				Computed:            true,
			},
			"ec_key_type": schema.StringAttribute{
				MarkdownDescription: "EC Key Type",
				Computed:            true,
			},
			"certificate_lifetime_in_days": schema.Int64Attribute{
				MarkdownDescription: "Certificate Lifetime(in Days)",
				Computed:            true,
			},
			"minimal_tls_version": schema.StringAttribute{
				MarkdownDescription: "Minimal TLS Version",
				Computed:            true,
			},
			"use_default_ca_cert_bundle": schema.BoolAttribute{
				MarkdownDescription: "Use default CA certificate bundle",
				Computed:            true,
			},
		},
	}
}

func (d *TLSSSLDecryptionPolicyDefinitionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *TLSSSLDecryptionPolicyDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config TLSSSLDecryptionPolicyDefinition

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
