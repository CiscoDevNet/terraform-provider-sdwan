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
	"regexp"
	"strings"
	"sync"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
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
var _ resource.Resource = &CloudProviderSettingsResource{}

func NewCloudProviderSettingsResource() resource.Resource {
	return &CloudProviderSettingsResource{}
}

type CloudProviderSettingsResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
	taskTimeout *int64
}

func (r *CloudProviderSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_provider_settings"
}

func (r *CloudProviderSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Cloud Provider Settings .").AddMinimumVersionDescription("20.15.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"umbrella_org_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Umbrella Organization ID (numeric string)").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+$`), ""),
				},
			},
			"umbrella_auth_key_v2": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Umbrella Authentication Key (v2)").AddMutualExclusivityDescription("Only one of `umbrella_auth_key_v2` and `umbrella_auth_key_v2_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `umbrella_auth_key_v2_wo` together with `umbrella_auth_key_v2_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"umbrella_auth_key_v2_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Umbrella Authentication Key (v2)").AddMutualExclusivityDescription("Only one of `umbrella_auth_key_v2` and `umbrella_auth_key_v2_wo` can be set.").String,
				Sensitive:           true,
				WriteOnly:           true,
				Optional:            true,
			},
			"umbrella_auth_key_v2_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `umbrella_auth_key_v2_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"umbrella_auth_secret_v2": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Umbrella Authentication Secret (v2)").AddMutualExclusivityDescription("Only one of `umbrella_auth_secret_v2` and `umbrella_auth_secret_v2_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `umbrella_auth_secret_v2_wo` together with `umbrella_auth_secret_v2_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"umbrella_auth_secret_v2_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Umbrella Authentication Secret (v2)").AddMutualExclusivityDescription("Only one of `umbrella_auth_secret_v2` and `umbrella_auth_secret_v2_wo` can be set.").String,
				Sensitive:           true,
				WriteOnly:           true,
				Optional:            true,
			},
			"umbrella_auth_secret_v2_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `umbrella_auth_secret_v2_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"umbrella_sig_auth_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Umbrella SIG Authentication Key").AddMutualExclusivityDescription("Only one of `umbrella_sig_auth_key` and `umbrella_sig_auth_key_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `umbrella_sig_auth_key_wo` together with `umbrella_sig_auth_key_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"umbrella_sig_auth_key_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Umbrella SIG Authentication Key").AddMutualExclusivityDescription("Only one of `umbrella_sig_auth_key` and `umbrella_sig_auth_key_wo` can be set.").String,
				Sensitive:           true,
				WriteOnly:           true,
				Optional:            true,
			},
			"umbrella_sig_auth_key_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `umbrella_sig_auth_key_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"umbrella_sig_auth_secret": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Umbrella SIG Authentication Secret").AddMutualExclusivityDescription("Only one of `umbrella_sig_auth_secret` and `umbrella_sig_auth_secret_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `umbrella_sig_auth_secret_wo` together with `umbrella_sig_auth_secret_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"umbrella_sig_auth_secret_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Umbrella SIG Authentication Secret").AddMutualExclusivityDescription("Only one of `umbrella_sig_auth_secret` and `umbrella_sig_auth_secret_wo` can be set.").String,
				Sensitive:           true,
				WriteOnly:           true,
				Optional:            true,
			},
			"umbrella_sig_auth_secret_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `umbrella_sig_auth_secret_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"umbrella_dns_auth_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Umbrella DNS Authentication Key").AddMutualExclusivityDescription("Only one of `umbrella_dns_auth_key` and `umbrella_dns_auth_key_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `umbrella_dns_auth_key_wo` together with `umbrella_dns_auth_key_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"umbrella_dns_auth_key_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Umbrella DNS Authentication Key").AddMutualExclusivityDescription("Only one of `umbrella_dns_auth_key` and `umbrella_dns_auth_key_wo` can be set.").String,
				Sensitive:           true,
				WriteOnly:           true,
				Optional:            true,
			},
			"umbrella_dns_auth_key_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `umbrella_dns_auth_key_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"umbrella_dns_auth_secret": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Umbrella DNS Authentication Secret").AddMutualExclusivityDescription("Only one of `umbrella_dns_auth_secret` and `umbrella_dns_auth_secret_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `umbrella_dns_auth_secret_wo` together with `umbrella_dns_auth_secret_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"umbrella_dns_auth_secret_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Umbrella DNS Authentication Secret").AddMutualExclusivityDescription("Only one of `umbrella_dns_auth_secret` and `umbrella_dns_auth_secret_wo` can be set.").String,
				Sensitive:           true,
				WriteOnly:           true,
				Optional:            true,
			},
			"umbrella_dns_auth_secret_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `umbrella_dns_auth_secret_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"zscaler_organization": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Zscaler Organization").String,
				Optional:            true,
			},
			"zscaler_partner_base_uri": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Zscaler Partner Base URI").String,
				Optional:            true,
			},
			"zscaler_partner_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Zscaler Partner Key").AddMutualExclusivityDescription("Only one of `zscaler_partner_key` and `zscaler_partner_key_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `zscaler_partner_key_wo` together with `zscaler_partner_key_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"zscaler_partner_key_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Zscaler Partner Key").AddMutualExclusivityDescription("Only one of `zscaler_partner_key` and `zscaler_partner_key_wo` can be set.").String,
				Sensitive:           true,
				WriteOnly:           true,
				Optional:            true,
			},
			"zscaler_partner_key_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `zscaler_partner_key_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"zscaler_username": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Zscaler Username").String,
				Optional:            true,
			},
			"zscaler_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Zscaler Password").AddMutualExclusivityDescription("Only one of `zscaler_password` and `zscaler_password_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `zscaler_password_wo` together with `zscaler_password_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"zscaler_password_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Zscaler Password").AddMutualExclusivityDescription("Only one of `zscaler_password` and `zscaler_password_wo` can be set.").String,
				Sensitive:           true,
				WriteOnly:           true,
				Optional:            true,
			},
			"zscaler_password_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `zscaler_password_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"cisco_sse_org_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cisco SSE Organization ID (numeric string)").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+$`), ""),
				},
			},
			"cisco_sse_auth_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cisco SSE Authentication Key").AddMutualExclusivityDescription("Only one of `cisco_sse_auth_key` and `cisco_sse_auth_key_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `cisco_sse_auth_key_wo` together with `cisco_sse_auth_key_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"cisco_sse_auth_key_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cisco SSE Authentication Key").AddMutualExclusivityDescription("Only one of `cisco_sse_auth_key` and `cisco_sse_auth_key_wo` can be set.").String,
				Sensitive:           true,
				WriteOnly:           true,
				Optional:            true,
			},
			"cisco_sse_auth_key_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `cisco_sse_auth_key_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"cisco_sse_auth_secret": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cisco SSE Authentication Secret").AddMutualExclusivityDescription("Only one of `cisco_sse_auth_secret` and `cisco_sse_auth_secret_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `cisco_sse_auth_secret_wo` together with `cisco_sse_auth_secret_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"cisco_sse_auth_secret_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cisco SSE Authentication Secret").AddMutualExclusivityDescription("Only one of `cisco_sse_auth_secret` and `cisco_sse_auth_secret_wo` can be set.").String,
				Sensitive:           true,
				WriteOnly:           true,
				Optional:            true,
			},
			"cisco_sse_auth_secret_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `cisco_sse_auth_secret_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"cisco_sse_context_sharing": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Cisco SSE Context Sharing").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
}

func (r *CloudProviderSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
	r.taskTimeout = req.ProviderData.(*SdwanProviderData).TaskTimeout
}

var _ resource.ResourceWithValidateConfig = &CloudProviderSettingsResource{}

// ValidateConfig enforces the rules between a secret attribute and its write-only twin.
//
// All diagnostics here are resource-scoped (AddError/AddWarning, never
// AddAttributeError) and the schema deliberately carries no
// ConflictsWith/ExactlyOneOf/AlsoRequires validators. Terraform renders an
// attribute-scoped diagnostic together with the offending raw configuration line,
// which would print the secret into plan output and CI logs. Sensitive: true does
// not suppress that.
//
// Generic resources have no "_variable" siblings - that concept only exists in the
// profile_parcels family - so the variable/secret precedence checks are omitted here.
func (r *CloudProviderSettingsResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {

	var vUmbrellaAuthKeyV2 types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_auth_key_v2"), &vUmbrellaAuthKeyV2)...)
	var vUmbrellaAuthKeyV2Wo types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_auth_key_v2_wo"), &vUmbrellaAuthKeyV2Wo)...)
	var vUmbrellaAuthKeyV2WoVersion types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_auth_key_v2_wo_version"), &vUmbrellaAuthKeyV2WoVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An unknown value is one that is not resolved yet at validate time, for example an
	// ephemeral resource attribute or a random_password reference. Treat it as "not
	// decided" rather than "set", otherwise every such configuration trips a false error.
	if !vUmbrellaAuthKeyV2.IsUnknown() && !vUmbrellaAuthKeyV2Wo.IsUnknown() && !vUmbrellaAuthKeyV2.IsNull() && !vUmbrellaAuthKeyV2Wo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `umbrella_auth_key_v2` and `umbrella_auth_key_v2_wo` can be set.",
		)
	}
	if !vUmbrellaAuthKeyV2Wo.IsUnknown() && !vUmbrellaAuthKeyV2WoVersion.IsUnknown() && !vUmbrellaAuthKeyV2Wo.IsNull() && vUmbrellaAuthKeyV2WoVersion.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`umbrella_auth_key_v2_wo_version` must be set when `umbrella_auth_key_v2_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !vUmbrellaAuthKeyV2WoVersion.IsUnknown() && !vUmbrellaAuthKeyV2Wo.IsUnknown() && !vUmbrellaAuthKeyV2WoVersion.IsNull() && vUmbrellaAuthKeyV2Wo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`umbrella_auth_key_v2_wo` must be set when `umbrella_auth_key_v2_wo_version` is used.",
		)
	}

	var vUmbrellaAuthSecretV2 types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_auth_secret_v2"), &vUmbrellaAuthSecretV2)...)
	var vUmbrellaAuthSecretV2Wo types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_auth_secret_v2_wo"), &vUmbrellaAuthSecretV2Wo)...)
	var vUmbrellaAuthSecretV2WoVersion types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_auth_secret_v2_wo_version"), &vUmbrellaAuthSecretV2WoVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An unknown value is one that is not resolved yet at validate time, for example an
	// ephemeral resource attribute or a random_password reference. Treat it as "not
	// decided" rather than "set", otherwise every such configuration trips a false error.
	if !vUmbrellaAuthSecretV2.IsUnknown() && !vUmbrellaAuthSecretV2Wo.IsUnknown() && !vUmbrellaAuthSecretV2.IsNull() && !vUmbrellaAuthSecretV2Wo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `umbrella_auth_secret_v2` and `umbrella_auth_secret_v2_wo` can be set.",
		)
	}
	if !vUmbrellaAuthSecretV2Wo.IsUnknown() && !vUmbrellaAuthSecretV2WoVersion.IsUnknown() && !vUmbrellaAuthSecretV2Wo.IsNull() && vUmbrellaAuthSecretV2WoVersion.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`umbrella_auth_secret_v2_wo_version` must be set when `umbrella_auth_secret_v2_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !vUmbrellaAuthSecretV2WoVersion.IsUnknown() && !vUmbrellaAuthSecretV2Wo.IsUnknown() && !vUmbrellaAuthSecretV2WoVersion.IsNull() && vUmbrellaAuthSecretV2Wo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`umbrella_auth_secret_v2_wo` must be set when `umbrella_auth_secret_v2_wo_version` is used.",
		)
	}

	var vUmbrellaSigAuthKey types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_sig_auth_key"), &vUmbrellaSigAuthKey)...)
	var vUmbrellaSigAuthKeyWo types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_sig_auth_key_wo"), &vUmbrellaSigAuthKeyWo)...)
	var vUmbrellaSigAuthKeyWoVersion types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_sig_auth_key_wo_version"), &vUmbrellaSigAuthKeyWoVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An unknown value is one that is not resolved yet at validate time, for example an
	// ephemeral resource attribute or a random_password reference. Treat it as "not
	// decided" rather than "set", otherwise every such configuration trips a false error.
	if !vUmbrellaSigAuthKey.IsUnknown() && !vUmbrellaSigAuthKeyWo.IsUnknown() && !vUmbrellaSigAuthKey.IsNull() && !vUmbrellaSigAuthKeyWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `umbrella_sig_auth_key` and `umbrella_sig_auth_key_wo` can be set.",
		)
	}
	if !vUmbrellaSigAuthKeyWo.IsUnknown() && !vUmbrellaSigAuthKeyWoVersion.IsUnknown() && !vUmbrellaSigAuthKeyWo.IsNull() && vUmbrellaSigAuthKeyWoVersion.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`umbrella_sig_auth_key_wo_version` must be set when `umbrella_sig_auth_key_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !vUmbrellaSigAuthKeyWoVersion.IsUnknown() && !vUmbrellaSigAuthKeyWo.IsUnknown() && !vUmbrellaSigAuthKeyWoVersion.IsNull() && vUmbrellaSigAuthKeyWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`umbrella_sig_auth_key_wo` must be set when `umbrella_sig_auth_key_wo_version` is used.",
		)
	}

	var vUmbrellaSigAuthSecret types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_sig_auth_secret"), &vUmbrellaSigAuthSecret)...)
	var vUmbrellaSigAuthSecretWo types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_sig_auth_secret_wo"), &vUmbrellaSigAuthSecretWo)...)
	var vUmbrellaSigAuthSecretWoVersion types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_sig_auth_secret_wo_version"), &vUmbrellaSigAuthSecretWoVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An unknown value is one that is not resolved yet at validate time, for example an
	// ephemeral resource attribute or a random_password reference. Treat it as "not
	// decided" rather than "set", otherwise every such configuration trips a false error.
	if !vUmbrellaSigAuthSecret.IsUnknown() && !vUmbrellaSigAuthSecretWo.IsUnknown() && !vUmbrellaSigAuthSecret.IsNull() && !vUmbrellaSigAuthSecretWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `umbrella_sig_auth_secret` and `umbrella_sig_auth_secret_wo` can be set.",
		)
	}
	if !vUmbrellaSigAuthSecretWo.IsUnknown() && !vUmbrellaSigAuthSecretWoVersion.IsUnknown() && !vUmbrellaSigAuthSecretWo.IsNull() && vUmbrellaSigAuthSecretWoVersion.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`umbrella_sig_auth_secret_wo_version` must be set when `umbrella_sig_auth_secret_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !vUmbrellaSigAuthSecretWoVersion.IsUnknown() && !vUmbrellaSigAuthSecretWo.IsUnknown() && !vUmbrellaSigAuthSecretWoVersion.IsNull() && vUmbrellaSigAuthSecretWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`umbrella_sig_auth_secret_wo` must be set when `umbrella_sig_auth_secret_wo_version` is used.",
		)
	}

	var vUmbrellaDnsAuthKey types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_dns_auth_key"), &vUmbrellaDnsAuthKey)...)
	var vUmbrellaDnsAuthKeyWo types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_dns_auth_key_wo"), &vUmbrellaDnsAuthKeyWo)...)
	var vUmbrellaDnsAuthKeyWoVersion types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_dns_auth_key_wo_version"), &vUmbrellaDnsAuthKeyWoVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An unknown value is one that is not resolved yet at validate time, for example an
	// ephemeral resource attribute or a random_password reference. Treat it as "not
	// decided" rather than "set", otherwise every such configuration trips a false error.
	if !vUmbrellaDnsAuthKey.IsUnknown() && !vUmbrellaDnsAuthKeyWo.IsUnknown() && !vUmbrellaDnsAuthKey.IsNull() && !vUmbrellaDnsAuthKeyWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `umbrella_dns_auth_key` and `umbrella_dns_auth_key_wo` can be set.",
		)
	}
	if !vUmbrellaDnsAuthKeyWo.IsUnknown() && !vUmbrellaDnsAuthKeyWoVersion.IsUnknown() && !vUmbrellaDnsAuthKeyWo.IsNull() && vUmbrellaDnsAuthKeyWoVersion.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`umbrella_dns_auth_key_wo_version` must be set when `umbrella_dns_auth_key_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !vUmbrellaDnsAuthKeyWoVersion.IsUnknown() && !vUmbrellaDnsAuthKeyWo.IsUnknown() && !vUmbrellaDnsAuthKeyWoVersion.IsNull() && vUmbrellaDnsAuthKeyWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`umbrella_dns_auth_key_wo` must be set when `umbrella_dns_auth_key_wo_version` is used.",
		)
	}

	var vUmbrellaDnsAuthSecret types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_dns_auth_secret"), &vUmbrellaDnsAuthSecret)...)
	var vUmbrellaDnsAuthSecretWo types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_dns_auth_secret_wo"), &vUmbrellaDnsAuthSecretWo)...)
	var vUmbrellaDnsAuthSecretWoVersion types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_dns_auth_secret_wo_version"), &vUmbrellaDnsAuthSecretWoVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An unknown value is one that is not resolved yet at validate time, for example an
	// ephemeral resource attribute or a random_password reference. Treat it as "not
	// decided" rather than "set", otherwise every such configuration trips a false error.
	if !vUmbrellaDnsAuthSecret.IsUnknown() && !vUmbrellaDnsAuthSecretWo.IsUnknown() && !vUmbrellaDnsAuthSecret.IsNull() && !vUmbrellaDnsAuthSecretWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `umbrella_dns_auth_secret` and `umbrella_dns_auth_secret_wo` can be set.",
		)
	}
	if !vUmbrellaDnsAuthSecretWo.IsUnknown() && !vUmbrellaDnsAuthSecretWoVersion.IsUnknown() && !vUmbrellaDnsAuthSecretWo.IsNull() && vUmbrellaDnsAuthSecretWoVersion.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`umbrella_dns_auth_secret_wo_version` must be set when `umbrella_dns_auth_secret_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !vUmbrellaDnsAuthSecretWoVersion.IsUnknown() && !vUmbrellaDnsAuthSecretWo.IsUnknown() && !vUmbrellaDnsAuthSecretWoVersion.IsNull() && vUmbrellaDnsAuthSecretWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`umbrella_dns_auth_secret_wo` must be set when `umbrella_dns_auth_secret_wo_version` is used.",
		)
	}

	var vZscalerPartnerKey types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("zscaler_partner_key"), &vZscalerPartnerKey)...)
	var vZscalerPartnerKeyWo types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("zscaler_partner_key_wo"), &vZscalerPartnerKeyWo)...)
	var vZscalerPartnerKeyWoVersion types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("zscaler_partner_key_wo_version"), &vZscalerPartnerKeyWoVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An unknown value is one that is not resolved yet at validate time, for example an
	// ephemeral resource attribute or a random_password reference. Treat it as "not
	// decided" rather than "set", otherwise every such configuration trips a false error.
	if !vZscalerPartnerKey.IsUnknown() && !vZscalerPartnerKeyWo.IsUnknown() && !vZscalerPartnerKey.IsNull() && !vZscalerPartnerKeyWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `zscaler_partner_key` and `zscaler_partner_key_wo` can be set.",
		)
	}
	if !vZscalerPartnerKeyWo.IsUnknown() && !vZscalerPartnerKeyWoVersion.IsUnknown() && !vZscalerPartnerKeyWo.IsNull() && vZscalerPartnerKeyWoVersion.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`zscaler_partner_key_wo_version` must be set when `zscaler_partner_key_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !vZscalerPartnerKeyWoVersion.IsUnknown() && !vZscalerPartnerKeyWo.IsUnknown() && !vZscalerPartnerKeyWoVersion.IsNull() && vZscalerPartnerKeyWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`zscaler_partner_key_wo` must be set when `zscaler_partner_key_wo_version` is used.",
		)
	}

	var vZscalerPassword types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("zscaler_password"), &vZscalerPassword)...)
	var vZscalerPasswordWo types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("zscaler_password_wo"), &vZscalerPasswordWo)...)
	var vZscalerPasswordWoVersion types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("zscaler_password_wo_version"), &vZscalerPasswordWoVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An unknown value is one that is not resolved yet at validate time, for example an
	// ephemeral resource attribute or a random_password reference. Treat it as "not
	// decided" rather than "set", otherwise every such configuration trips a false error.
	if !vZscalerPassword.IsUnknown() && !vZscalerPasswordWo.IsUnknown() && !vZscalerPassword.IsNull() && !vZscalerPasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `zscaler_password` and `zscaler_password_wo` can be set.",
		)
	}
	if !vZscalerPasswordWo.IsUnknown() && !vZscalerPasswordWoVersion.IsUnknown() && !vZscalerPasswordWo.IsNull() && vZscalerPasswordWoVersion.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`zscaler_password_wo_version` must be set when `zscaler_password_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !vZscalerPasswordWoVersion.IsUnknown() && !vZscalerPasswordWo.IsUnknown() && !vZscalerPasswordWoVersion.IsNull() && vZscalerPasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`zscaler_password_wo` must be set when `zscaler_password_wo_version` is used.",
		)
	}

	var vCiscoSseAuthKey types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("cisco_sse_auth_key"), &vCiscoSseAuthKey)...)
	var vCiscoSseAuthKeyWo types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("cisco_sse_auth_key_wo"), &vCiscoSseAuthKeyWo)...)
	var vCiscoSseAuthKeyWoVersion types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("cisco_sse_auth_key_wo_version"), &vCiscoSseAuthKeyWoVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An unknown value is one that is not resolved yet at validate time, for example an
	// ephemeral resource attribute or a random_password reference. Treat it as "not
	// decided" rather than "set", otherwise every such configuration trips a false error.
	if !vCiscoSseAuthKey.IsUnknown() && !vCiscoSseAuthKeyWo.IsUnknown() && !vCiscoSseAuthKey.IsNull() && !vCiscoSseAuthKeyWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `cisco_sse_auth_key` and `cisco_sse_auth_key_wo` can be set.",
		)
	}
	if !vCiscoSseAuthKeyWo.IsUnknown() && !vCiscoSseAuthKeyWoVersion.IsUnknown() && !vCiscoSseAuthKeyWo.IsNull() && vCiscoSseAuthKeyWoVersion.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`cisco_sse_auth_key_wo_version` must be set when `cisco_sse_auth_key_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !vCiscoSseAuthKeyWoVersion.IsUnknown() && !vCiscoSseAuthKeyWo.IsUnknown() && !vCiscoSseAuthKeyWoVersion.IsNull() && vCiscoSseAuthKeyWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`cisco_sse_auth_key_wo` must be set when `cisco_sse_auth_key_wo_version` is used.",
		)
	}

	var vCiscoSseAuthSecret types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("cisco_sse_auth_secret"), &vCiscoSseAuthSecret)...)
	var vCiscoSseAuthSecretWo types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("cisco_sse_auth_secret_wo"), &vCiscoSseAuthSecretWo)...)
	var vCiscoSseAuthSecretWoVersion types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("cisco_sse_auth_secret_wo_version"), &vCiscoSseAuthSecretWoVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An unknown value is one that is not resolved yet at validate time, for example an
	// ephemeral resource attribute or a random_password reference. Treat it as "not
	// decided" rather than "set", otherwise every such configuration trips a false error.
	if !vCiscoSseAuthSecret.IsUnknown() && !vCiscoSseAuthSecretWo.IsUnknown() && !vCiscoSseAuthSecret.IsNull() && !vCiscoSseAuthSecretWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `cisco_sse_auth_secret` and `cisco_sse_auth_secret_wo` can be set.",
		)
	}
	if !vCiscoSseAuthSecretWo.IsUnknown() && !vCiscoSseAuthSecretWoVersion.IsUnknown() && !vCiscoSseAuthSecretWo.IsNull() && vCiscoSseAuthSecretWoVersion.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`cisco_sse_auth_secret_wo_version` must be set when `cisco_sse_auth_secret_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !vCiscoSseAuthSecretWoVersion.IsUnknown() && !vCiscoSseAuthSecretWo.IsUnknown() && !vCiscoSseAuthSecretWoVersion.IsNull() && vCiscoSseAuthSecretWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`cisco_sse_auth_secret_wo` must be set when `cisco_sse_auth_secret_wo_version` is used.",
		)
	}
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *CloudProviderSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CloudProviderSettings

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "umbrella_auth_key_v2_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_auth_key_v2_wo"), &plan.UmbrellaAuthKeyV2Wo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "umbrella_auth_secret_v2_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_auth_secret_v2_wo"), &plan.UmbrellaAuthSecretV2Wo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "umbrella_sig_auth_key_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_sig_auth_key_wo"), &plan.UmbrellaSigAuthKeyWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "umbrella_sig_auth_secret_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_sig_auth_secret_wo"), &plan.UmbrellaSigAuthSecretWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "umbrella_dns_auth_key_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_dns_auth_key_wo"), &plan.UmbrellaDnsAuthKeyWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "umbrella_dns_auth_secret_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_dns_auth_secret_wo"), &plan.UmbrellaDnsAuthSecretWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "zscaler_partner_key_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("zscaler_partner_key_wo"), &plan.ZscalerPartnerKeyWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "zscaler_password_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("zscaler_password_wo"), &plan.ZscalerPasswordWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "cisco_sse_auth_key_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("cisco_sse_auth_key_wo"), &plan.CiscoSseAuthKeyWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "cisco_sse_auth_secret_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("cisco_sse_auth_secret_wo"), &plan.CiscoSseAuthSecretWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue("cloud_provider_settings")

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *CloudProviderSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CloudProviderSettings

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	res, err := r.client.Get(state.getPath())
	if strings.Contains(res.Get("error.message").String(), "Failed to find specified resource") || strings.Contains(res.Get("error.message").String(), "Invalid template type") || strings.Contains(res.Get("error.message").String(), "Template definition not found") || strings.Contains(res.Get("error.message").String(), "Invalid Profile Id") || strings.Contains(res.Get("error.message").String(), "Invalid feature Id") || strings.Contains(res.Get("error.message").String(), "Invalid config group passed") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *CloudProviderSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CloudProviderSettings

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "umbrella_auth_key_v2_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_auth_key_v2_wo"), &plan.UmbrellaAuthKeyV2Wo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "umbrella_auth_secret_v2_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_auth_secret_v2_wo"), &plan.UmbrellaAuthSecretV2Wo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "umbrella_sig_auth_key_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_sig_auth_key_wo"), &plan.UmbrellaSigAuthKeyWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "umbrella_sig_auth_secret_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_sig_auth_secret_wo"), &plan.UmbrellaSigAuthSecretWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "umbrella_dns_auth_key_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_dns_auth_key_wo"), &plan.UmbrellaDnsAuthKeyWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "umbrella_dns_auth_secret_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("umbrella_dns_auth_secret_wo"), &plan.UmbrellaDnsAuthSecretWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "zscaler_partner_key_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("zscaler_partner_key_wo"), &plan.ZscalerPartnerKeyWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "zscaler_password_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("zscaler_password_wo"), &plan.ZscalerPasswordWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "cisco_sse_auth_key_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("cisco_sse_auth_key_wo"), &plan.CiscoSseAuthKeyWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "cisco_sse_auth_secret_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because the PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("cisco_sse_auth_secret_wo"), &plan.CiscoSseAuthSecretWo)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	if plan.hasChanges(ctx, &state) {
		body := plan.toBody(ctx)
		r.updateMutex.Lock()
		res, err := r.client.Put(plan.getPath(), body)
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
		tflog.Debug(ctx, fmt.Sprintf("%s: No changes detected", plan.Id.ValueString()))
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *CloudProviderSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CloudProviderSettings

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	// No delete operation supported by the API for this resource; removing from state only.

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import
