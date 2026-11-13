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
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
var _ resource.Resource = &SystemRemoteAccessProfileParcelResource{}
var _ resource.ResourceWithImportState = &SystemRemoteAccessProfileParcelResource{}

func NewSystemRemoteAccessProfileParcelResource() resource.Resource {
	return &SystemRemoteAccessProfileParcelResource{}
}

type SystemRemoteAccessProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *SystemRemoteAccessProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_remote_access_feature"
}

func (r *SystemRemoteAccessProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a System Remote Access Feature.").AddMinimumVersionDescription("20.15.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Feature",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Feature",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Feature",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Feature",
				Optional:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Feature Profile ID").String,
				Required:            true,
			},
			"connection_type_ssl": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enabled SSL VPN").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"any_connect_eap_authentication_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription(", Attribute conditional on `connection_type_ssl` equal to `false`").AddStringEnumDescription("user", "device").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("user", "device"),
				},
			},
			"ipv4_pool_size": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 Pool Size").AddDefaultValueDescription("1000").String,
				Optional:            true,
			},
			"ipv4_pool_size_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv6_pool_size": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 Pool Size").AddDefaultValueDescription("1024").String,
				Optional:            true,
			},
			"ipv6_pool_size_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"enable_certificate_list_check": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"enable_certificate_list_check_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"psk_authentication_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("PSK Selection, Attribute conditional on `connection_type_ssl` equal to `false`").AddStringEnumDescription("aaa", "group").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("aaa", "group"),
				},
			},
			"psk_authentication_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `connection_type_ssl` equal to `false`").String,
				Optional:            true,
			},
			"psk_authentication_pre_shared_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("PSK Pre Shared Key, Attribute conditional on `psk_authentication_type` equal to `group`").AddMutualExclusivityDescription("Only one of `psk_authentication_pre_shared_key`, `psk_authentication_pre_shared_key_variable` and `psk_authentication_pre_shared_key_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `psk_authentication_pre_shared_key_wo` together with `psk_authentication_pre_shared_key_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 25),
				},
			},
			"psk_authentication_pre_shared_key_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `psk_authentication_type` equal to `group`").String,
				Optional:            true,
			},
			"psk_authentication_pre_shared_key_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("PSK Pre Shared Key, Attribute conditional on `psk_authentication_type` equal to `group`").AddMutualExclusivityDescription("Only one of `psk_authentication_pre_shared_key`, `psk_authentication_pre_shared_key_variable` and `psk_authentication_pre_shared_key_wo` can be set.").String,
				Sensitive:           true,
				WriteOnly:           true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 25),
				},
			},
			"psk_authentication_pre_shared_key_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `psk_authentication_pre_shared_key_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"radius_group_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"radius_group_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"aaa_specify_name_policy_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"aaa_specify_name_policy_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"aaa_specify_name_policy_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddMutualExclusivityDescription("Only one of `aaa_specify_name_policy_password`, `aaa_specify_name_policy_password_variable` and `aaa_specify_name_policy_password_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `aaa_specify_name_policy_password_wo` together with `aaa_specify_name_policy_password_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 25),
				},
			},
			"aaa_specify_name_policy_password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"aaa_specify_name_policy_password_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddMutualExclusivityDescription("Only one of `aaa_specify_name_policy_password`, `aaa_specify_name_policy_password_variable` and `aaa_specify_name_policy_password_wo` can be set.").String,
				Sensitive:           true,
				WriteOnly:           true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 25),
				},
			},
			"aaa_specify_name_policy_password_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `aaa_specify_name_policy_password_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"aaa_derive_name_from_peer_identity": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription(", Attribute conditional on `connection_type_ssl` equal to `false`").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 25),
				},
			},
			"aaa_derive_name_from_peer_identity_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `connection_type_ssl` equal to `false`").String,
				Optional:            true,
			},
			"aaa_derive_name_from_peer_domain": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription(", Attribute conditional on `connection_type_ssl` equal to `false`").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 25),
				},
			},
			"aaa_derive_name_from_peer_domain_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `connection_type_ssl` equal to `false`").String,
				Optional:            true,
			},
			"aaa_enable_accounting": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Accounting").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"aaa_enable_accounting_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ikev2_local_ike_identity_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription(", Attribute conditional on `connection_type_ssl` equal to `false`").AddStringEnumDescription("EMAIL", "FQDN", "KEYID", "IPv4 ADDRESS", "IPv6 ADDRESS").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("EMAIL", "FQDN", "KEYID", "IPv4 ADDRESS", "IPv6 ADDRESS"),
				},
			},
			"ikev2_local_ike_identity_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `connection_type_ssl` equal to `false`").String,
				Optional:            true,
			},
			"ikev2_local_ike_identity_value": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription(", Attribute conditional on `connection_type_ssl` equal to `false`").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"ikev2_local_ike_identity_value_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `connection_type_ssl` equal to `false`").String,
				Optional:            true,
			},
			"ikev2_security_association_lifetime": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Security Association Lifetime in Seconds, Attribute conditional on `connection_type_ssl` equal to `false`").AddIntegerRangeDescription(3600, 86400).AddDefaultValueDescription("86400").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(3600, 86400),
				},
			},
			"ikev2_security_association_lifetime_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `connection_type_ssl` equal to `false`").String,
				Optional:            true,
			},
			"ikev2_anti_dos_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Anti-DOS Threshold, Attribute conditional on `connection_type_ssl` equal to `false`").AddIntegerRangeDescription(10, 1000).AddDefaultValueDescription("100").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 1000),
				},
			},
			"ikev2_anti_dos_threshold_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `connection_type_ssl` equal to `false`").String,
				Optional:            true,
			},
			"ipsec_enable_anti_replay": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Anti-Replay, Attribute conditional on `connection_type_ssl` equal to `false`").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"ipsec_enable_anti_replay_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `connection_type_ssl` equal to `false`").String,
				Optional:            true,
			},
			"ipsec_anti_replay_window_size": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("security Association Lifetime, Attribute conditional on `ipsec_enable_anti_replay` equal to `true`").AddDefaultValueDescription("64").String,
				Optional:            true,
			},
			"ipsec_anti_replay_window_size_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `ipsec_enable_anti_replay` equal to `true`").String,
				Optional:            true,
			},
			"ipsec_security_association_lifetime": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Security Association Lifetime in Seconds, Attribute conditional on `connection_type_ssl` equal to `false`").AddIntegerRangeDescription(3600, 86400).AddDefaultValueDescription("3600").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(3600, 86400),
				},
			},
			"ipsec_security_association_lifetime_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `connection_type_ssl` equal to `false`").String,
				Optional:            true,
			},
			"ipsec_enable_perfect_foward_secrecy": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("security Association Lifetime, Attribute conditional on `connection_type_ssl` equal to `false`").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ipsec_enable_perfect_foward_secrecy_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `connection_type_ssl` equal to `false`").String,
				Optional:            true,
			},
		},
	}
}

func (r *SystemRemoteAccessProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

var _ resource.ResourceWithValidateConfig = &SystemRemoteAccessProfileParcelResource{}

// ValidateConfig enforces the rules between a secret attribute, its "_variable"
// sibling and its write-only twin.
//
// All diagnostics here are resource-scoped (AddError/AddWarning, never
// AddAttributeError) and the schema deliberately carries no
// ConflictsWith/ExactlyOneOf/AlsoRequires validators. Terraform renders an
// attribute-scoped diagnostic together with the offending raw configuration line,
// which would print the secret into plan output and CI logs. Sensitive: true does
// not suppress that.
func (r *SystemRemoteAccessProfileParcelResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {

	var vPskAuthenticationPreSharedKey types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("psk_authentication_pre_shared_key"), &vPskAuthenticationPreSharedKey)...)
	var vPskAuthenticationPreSharedKeyVariable types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("psk_authentication_pre_shared_key_variable"), &vPskAuthenticationPreSharedKeyVariable)...)
	var vPskAuthenticationPreSharedKeyWo types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("psk_authentication_pre_shared_key_wo"), &vPskAuthenticationPreSharedKeyWo)...)
	var vPskAuthenticationPreSharedKeyWoVersion types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("psk_authentication_pre_shared_key_wo_version"), &vPskAuthenticationPreSharedKeyWoVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An unknown value is one that is not resolved yet at validate time, for example an
	// ephemeral resource attribute or a random_password reference. Treat it as "not
	// decided" rather than "set", otherwise every such configuration trips a false error.
	if !vPskAuthenticationPreSharedKey.IsUnknown() && !vPskAuthenticationPreSharedKeyWo.IsUnknown() && !vPskAuthenticationPreSharedKey.IsNull() && !vPskAuthenticationPreSharedKeyWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `psk_authentication_pre_shared_key` and `psk_authentication_pre_shared_key_wo` can be set.",
		)
	}
	if !vPskAuthenticationPreSharedKeyVariable.IsUnknown() && !vPskAuthenticationPreSharedKeyWo.IsUnknown() && !vPskAuthenticationPreSharedKeyVariable.IsNull() && !vPskAuthenticationPreSharedKeyWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"`psk_authentication_pre_shared_key_variable` and `psk_authentication_pre_shared_key_wo` cannot both be set. `psk_authentication_pre_shared_key_variable` defers the value to a device variable, so the write-only value would never be sent.",
		)
	}
	if !vPskAuthenticationPreSharedKey.IsUnknown() && !vPskAuthenticationPreSharedKeyVariable.IsUnknown() && !vPskAuthenticationPreSharedKey.IsNull() && !vPskAuthenticationPreSharedKeyVariable.IsNull() {
		resp.Diagnostics.AddWarning(
			"Conflicting Attribute Combination",
			"Both `psk_authentication_pre_shared_key` and `psk_authentication_pre_shared_key_variable` are set. `psk_authentication_pre_shared_key_variable` takes precedence, so the value of `psk_authentication_pre_shared_key` is not sent to the device - but it is still stored in Terraform state. Set only one of them.",
		)
	}
	if !vPskAuthenticationPreSharedKeyWo.IsUnknown() && !vPskAuthenticationPreSharedKeyWoVersion.IsUnknown() && !vPskAuthenticationPreSharedKeyWo.IsNull() && vPskAuthenticationPreSharedKeyWoVersion.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`psk_authentication_pre_shared_key_wo_version` must be set when `psk_authentication_pre_shared_key_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !vPskAuthenticationPreSharedKeyWoVersion.IsUnknown() && !vPskAuthenticationPreSharedKeyWo.IsUnknown() && !vPskAuthenticationPreSharedKeyWoVersion.IsNull() && vPskAuthenticationPreSharedKeyWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`psk_authentication_pre_shared_key_wo` must be set when `psk_authentication_pre_shared_key_wo_version` is used.",
		)
	}

	var vAaaSpecifyNamePolicyPassword types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("aaa_specify_name_policy_password"), &vAaaSpecifyNamePolicyPassword)...)
	var vAaaSpecifyNamePolicyPasswordVariable types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("aaa_specify_name_policy_password_variable"), &vAaaSpecifyNamePolicyPasswordVariable)...)
	var vAaaSpecifyNamePolicyPasswordWo types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("aaa_specify_name_policy_password_wo"), &vAaaSpecifyNamePolicyPasswordWo)...)
	var vAaaSpecifyNamePolicyPasswordWoVersion types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("aaa_specify_name_policy_password_wo_version"), &vAaaSpecifyNamePolicyPasswordWoVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An unknown value is one that is not resolved yet at validate time, for example an
	// ephemeral resource attribute or a random_password reference. Treat it as "not
	// decided" rather than "set", otherwise every such configuration trips a false error.
	if !vAaaSpecifyNamePolicyPassword.IsUnknown() && !vAaaSpecifyNamePolicyPasswordWo.IsUnknown() && !vAaaSpecifyNamePolicyPassword.IsNull() && !vAaaSpecifyNamePolicyPasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `aaa_specify_name_policy_password` and `aaa_specify_name_policy_password_wo` can be set.",
		)
	}
	if !vAaaSpecifyNamePolicyPasswordVariable.IsUnknown() && !vAaaSpecifyNamePolicyPasswordWo.IsUnknown() && !vAaaSpecifyNamePolicyPasswordVariable.IsNull() && !vAaaSpecifyNamePolicyPasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"`aaa_specify_name_policy_password_variable` and `aaa_specify_name_policy_password_wo` cannot both be set. `aaa_specify_name_policy_password_variable` defers the value to a device variable, so the write-only value would never be sent.",
		)
	}
	if !vAaaSpecifyNamePolicyPassword.IsUnknown() && !vAaaSpecifyNamePolicyPasswordVariable.IsUnknown() && !vAaaSpecifyNamePolicyPassword.IsNull() && !vAaaSpecifyNamePolicyPasswordVariable.IsNull() {
		resp.Diagnostics.AddWarning(
			"Conflicting Attribute Combination",
			"Both `aaa_specify_name_policy_password` and `aaa_specify_name_policy_password_variable` are set. `aaa_specify_name_policy_password_variable` takes precedence, so the value of `aaa_specify_name_policy_password` is not sent to the device - but it is still stored in Terraform state. Set only one of them.",
		)
	}
	if !vAaaSpecifyNamePolicyPasswordWo.IsUnknown() && !vAaaSpecifyNamePolicyPasswordWoVersion.IsUnknown() && !vAaaSpecifyNamePolicyPasswordWo.IsNull() && vAaaSpecifyNamePolicyPasswordWoVersion.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`aaa_specify_name_policy_password_wo_version` must be set when `aaa_specify_name_policy_password_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !vAaaSpecifyNamePolicyPasswordWoVersion.IsUnknown() && !vAaaSpecifyNamePolicyPasswordWo.IsUnknown() && !vAaaSpecifyNamePolicyPasswordWoVersion.IsNull() && vAaaSpecifyNamePolicyPasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`aaa_specify_name_policy_password_wo` must be set when `aaa_specify_name_policy_password_wo_version` is used.",
		)
	}
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *SystemRemoteAccessProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemRemoteAccess

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "psk_authentication_pre_shared_key_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because a parcel PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("psk_authentication_pre_shared_key_wo"), &plan.PskAuthenticationPreSharedKeyWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "aaa_specify_name_policy_password_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because a parcel PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("aaa_specify_name_policy_password_wo"), &plan.AaaSpecifyNamePolicyPasswordWo)...)
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
func (r *SystemRemoteAccessProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemRemoteAccess

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

	state.fromBody(ctx, res, imp)
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
func (r *SystemRemoteAccessProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemRemoteAccess

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "psk_authentication_pre_shared_key_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because a parcel PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("psk_authentication_pre_shared_key_wo"), &plan.PskAuthenticationPreSharedKeyWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "aaa_specify_name_policy_password_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because a parcel PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("aaa_specify_name_policy_password_wo"), &plan.AaaSpecifyNamePolicyPasswordWo)...)
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
func (r *SystemRemoteAccessProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemRemoteAccess

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
func (r *SystemRemoteAccessProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "system_remote_access_feature_id" + ",feature_profile_id"
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
