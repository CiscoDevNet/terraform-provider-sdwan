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
var _ resource.Resource = &PolicyObjectUnifiedTLSSSLDecryptionProfileParcelResource{}
var _ resource.ResourceWithImportState = &PolicyObjectUnifiedTLSSSLDecryptionProfileParcelResource{}

func NewPolicyObjectUnifiedTLSSSLDecryptionProfileParcelResource() resource.Resource {
	return &PolicyObjectUnifiedTLSSSLDecryptionProfileParcelResource{}
}

type PolicyObjectUnifiedTLSSSLDecryptionProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *PolicyObjectUnifiedTLSSSLDecryptionProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_object_unified_tls_ssl_decryption"
}

func (r *PolicyObjectUnifiedTLSSSLDecryptionProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Policy Object Unified TLS SSL Decryption Policy_object.").AddMinimumVersionDescription("20.12.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Policy_object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Policy_object",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Policy_object",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Policy_object",
				Optional:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Feature Profile ID").String,
				Required:            true,
			},
			"expired_certificate": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("decrypt", "drop").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("decrypt", "drop"),
				},
			},
			"untrusted_certificate": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("decrypt", "drop").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("decrypt", "drop"),
				},
			},
			"certificate_revocation_status": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If value is none unknown status not required, if value is ocsp then unknown status is required").AddStringEnumDescription("ocsp", "none").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ocsp", "none"),
				},
			},
			"unknown_revocation_status": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Only required if certificateRevocationStatus is oscp, if value is none then field shouldn't be here, Attribute conditional on `certificate_revocation_status` equal to `ocsp`").AddStringEnumDescription("decrypt", "drop").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("decrypt", "drop"),
				},
			},
			"unsupported_protocol_versions": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("no-decrypt", "drop").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("no-decrypt", "drop"),
				},
			},
			"unsupported_cipher_suites": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("no-decrypt", "drop").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("no-decrypt", "drop"),
				},
			},
			"failure_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("close", "open").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("close", "open"),
				},
			},
			"default_ca_certificate_bundle": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"file_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"bundle_string": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"rsa_keypair_modules": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("1024", "2048", "4096").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1024", "2048", "4096"),
				},
			},
			"ec_key_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("P256", "P384", "P521").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("P256", "P384", "P521"),
				},
			},
			"certificate_lifetime": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If you have vManage as CA or vManage as intermediate CA, this value should be 1").String,
				Required:            true,
			},
			"minimal_tls_ver": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("TLSv1", "TLSv1.1", "TLSv1.2").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("TLSv1", "TLSv1.1", "TLSv1.2"),
				},
			},
		},
	}
}

func (r *PolicyObjectUnifiedTLSSSLDecryptionProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *PolicyObjectUnifiedTLSSSLDecryptionProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PolicyObjectUnifiedTLSSSLDecryption

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

	plan.Id = types.StringValue(res.Get("parcelId").String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *PolicyObjectUnifiedTLSSSLDecryptionProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PolicyObjectUnifiedTLSSSLDecryption

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if res.Get("error.message").String() == "Invalid feature Id" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	if imp {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}
	if state.Version.IsNull() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *PolicyObjectUnifiedTLSSSLDecryptionProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PolicyObjectUnifiedTLSSSLDecryption

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

	body := plan.toBody(ctx)
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *PolicyObjectUnifiedTLSSSLDecryptionProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PolicyObjectUnifiedTLSSSLDecryption

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && res.Get("error.message").String() != "Invalid Template Id" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *PolicyObjectUnifiedTLSSSLDecryptionProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "policy_object_unified_tls_ssl_decryption_id" + ",feature_profile_id"
	if len(parts) != (count + 1) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: %s. Got: %q", pattern, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("feature_profile_id"), parts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
