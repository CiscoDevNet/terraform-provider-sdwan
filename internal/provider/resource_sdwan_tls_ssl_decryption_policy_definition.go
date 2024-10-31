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
	"sync"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &TLSSSLDecryptionPolicyDefinitionResource{}
var _ resource.ResourceWithImportState = &TLSSSLDecryptionPolicyDefinitionResource{}

func NewTLSSSLDecryptionPolicyDefinitionResource() resource.Resource {
	return &TLSSSLDecryptionPolicyDefinitionResource{}
}

type TLSSSLDecryptionPolicyDefinitionResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *TLSSSLDecryptionPolicyDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tls_ssl_decryption_policy_definition"
}

func (r *TLSSSLDecryptionPolicyDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a TLS SSL Decryption Policy Definition .").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the policy definition.").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The description of the policy definition.").String,
				Required:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The policy mode").AddStringEnumDescription("security", "unified").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("security", "unified"),
				},
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default action (applies when `mode` set to `security`)").AddStringEnumDescription("noIntent", "doNotDecrypt", "decrypt").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("noIntent", "doNotDecrypt", "decrypt"),
				},
			},
			"network_rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of network rules (applies when `mode` set to `security`)").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"base_action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Rule base action").AddStringEnumDescription("noIntent", "doNotDecrypt", "decrypt").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("noIntent", "doNotDecrypt", "decrypt"),
							},
						},
						"rule_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Rule ID").String,
							Optional:            true,
						},
						"rule_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Rule name").String,
							Optional:            true,
						},
						"rule_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Rule type").String,
							Optional:            true,
						},
						"source_and_destination_configuration": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of network source / destination configuration").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"option": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("source / destination option").AddStringEnumDescription("sourceIp", "sourcePort", "destinationVpn", "destinationIp", "destinationPort").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("sourceIp", "sourcePort", "destinationVpn", "destinationIp", "destinationPort"),
										},
									},
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("source / destination option target").String,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
			"url_rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of url rules (applies when `mode` set to `security`)").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"rule_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Country").String,
							Optional:            true,
						},
						"target_vpns": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of VPN IDs").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"tls_ssl_profile_policy_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("TLS SSL Profile Policy ID").String,
							Optional:            true,
						},
						"tls_ssl_profile_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("TLS SSL Profile Policy version").String,
							Optional:            true,
						},
					},
				},
			},
			"ssl_decryption_enabled": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SSL decryption enabled").String,
				Optional:            true,
			},
			"expired_certificate": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Expired certificate action").AddStringEnumDescription("drop", "decrypt").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("drop", "decrypt"),
				},
			},
			"untrusted_certificate": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Untrusted certificate action").AddStringEnumDescription("drop", "decrypt").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("drop", "decrypt"),
				},
			},
			"certificate_revocation_status": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Certificate revocation status").AddStringEnumDescription("ocsp", "none").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ocsp", "none"),
				},
			},
			"unknown_revocation_status": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Unknown revocation status action").AddStringEnumDescription("drop", "decrypt").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("drop", "decrypt"),
				},
			},
			"unsupported_protocol_versions": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Unsupported protocol versions action").AddStringEnumDescription("drop", "no-decrypt").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("drop", "no-decrypt"),
				},
			},
			"unsupported_cipher_suites": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Unsupported cipher suites action").AddStringEnumDescription("drop", "no-decrypt").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("drop", "no-decrypt"),
				},
			},
			"failure_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Failure mode").AddStringEnumDescription("open", "close").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("open", "close"),
				},
			},
			"rsa_key_pair_modulus": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("RSA key pair modules").AddStringEnumDescription("1024", "2048", "4096").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1024", "2048", "4096"),
				},
			},
			"ec_key_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("EC Key Type").AddStringEnumDescription("P256", "P384", "P521").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("P256", "P384", "P521"),
				},
			},
			"certificate_lifetime_in_days": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Certificate Lifetime(in Days)").String,
				Optional:            true,
			},
			"minimal_tls_version": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimal TLS Version").AddStringEnumDescription("TLSv1.0", "TLSv1.1", "TLSv1.2").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("TLSv1.0", "TLSv1.1", "TLSv1.2"),
				},
			},
			"use_default_ca_cert_bundle": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use default CA certificate bundle").String,
				Optional:            true,
			},
		},
	}
}

func (r *TLSSSLDecryptionPolicyDefinitionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *TLSSSLDecryptionPolicyDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TLSSSLDecryptionPolicyDefinition

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("definitionId").String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *TLSSSLDecryptionPolicyDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state TLSSSLDecryptionPolicyDefinition

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if strings.Contains(res.Get("error.message").String(), "Failed to find specified resource") || strings.Contains(res.Get("error.message").String(), "Invalid template type") || strings.Contains(res.Get("error.message").String(), "Template definition not found") || strings.Contains(res.Get("error.message").String(), "Invalid Profile Id") || strings.Contains(res.Get("error.message").String(), "Invalid feature Id") || strings.Contains(res.Get("error.message").String(), "Invalid config group passed") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)
	if state.Version.IsNull() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *TLSSSLDecryptionPolicyDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state TLSSSLDecryptionPolicyDefinition

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Name.ValueString()))

	if plan.hasChanges(ctx, &state) {
		body := plan.toBody(ctx)
		r.updateMutex.Lock()
		res, err := r.client.Put(plan.getPath()+url.QueryEscape(plan.Id.ValueString()), body)
		r.updateMutex.Unlock()
		if err != nil {
			if strings.Contains(res.Get("error.message").String(), "Failed to acquire lock") {
				resp.Diagnostics.AddWarning("Client Warning", "Failed to modify policy due to policy being locked by another change. Policy changes will not be applied. Re-run 'terraform apply' to try again.")
			} else if strings.Contains(res.Get("error.message").String(), "Template locked in edit mode") {
				resp.Diagnostics.AddWarning("Client Warning", "Failed to modify template due to template being locked by another change. Template changes will not be applied. Re-run 'terraform apply' to try again.")
			} else {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
				return
			}
		}
	} else {
		tflog.Debug(ctx, fmt.Sprintf("%s: No changes detected", plan.Name.ValueString()))
	}
	plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *TLSSSLDecryptionPolicyDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TLSSSLDecryptionPolicyDefinition

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *TLSSSLDecryptionPolicyDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
