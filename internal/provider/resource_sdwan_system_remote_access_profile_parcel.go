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
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a System Remote Access Feature.").AddMinimumVersionDescription("20.12.0").String,

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
				Optional:            true,
			},
			"connection_type_ssl": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enabled SSL VPN").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"any_connect_eap_authentication_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription(", Attribute conditional on `connection_type_ssl` being equal to `false`").AddStringEnumDescription("user", "device").String,
				Required:            true,
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
				MarkdownDescription: helpers.NewAttributeDescription("PSK Selection, Attribute conditional on `connection_type_ssl` being equal to `false`").AddStringEnumDescription("aaa", "group").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("aaa", "group"),
				},
			},
			"psk_authentication_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"psk_authentication_pre_shared_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("PSK Pre Shared Key, Attribute conditional on `psk_authentication_type` being equal to `group`").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 25),
				},
			},
			"psk_authentication_pre_shared_key_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
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
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 25),
				},
			},
			"aaa_specify_name_policy_password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"aaa_derive_name_from_peer_identity": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription(", Attribute conditional on `connection_type_ssl` being equal to `false`").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 25),
				},
			},
			"aaa_derive_name_from_peer_identity_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"aaa_derive_name_from_peer_domain": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription(", Attribute conditional on `connection_type_ssl` being equal to `false`").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 25),
				},
			},
			"aaa_derive_name_from_peer_domain_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
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
				MarkdownDescription: helpers.NewAttributeDescription(", Attribute conditional on `connection_type_ssl` being equal to `false`").AddStringEnumDescription("EMAIL", "FQDN", "KEYID", "IPv4 ADDRESS", "IPv6 ADDRESS").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("EMAIL", "FQDN", "KEYID", "IPv4 ADDRESS", "IPv6 ADDRESS"),
				},
			},
			"ikev2_local_ike_identity_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ikev2_local_ike_identity_value": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription(", Attribute conditional on `connection_type_ssl` being equal to `false`").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"ikev2_local_ike_identity_value_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ikev2_security_association_lifetime": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Security Association Lifetime in Seconds, Attribute conditional on `connection_type_ssl` being equal to `false`").AddIntegerRangeDescription(3600, 86400).AddDefaultValueDescription("86400").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(3600, 86400),
				},
			},
			"ikev2_security_association_lifetime_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ikev2_anti_dos_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Anti-DOS Threshold, Attribute conditional on `connection_type_ssl` being equal to `false`").AddIntegerRangeDescription(10, 1000).AddDefaultValueDescription("100").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 1000),
				},
			},
			"ikev2_anti_dos_threshold_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipsec_enable_anti_replay": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Anti-Replay, Attribute conditional on `connection_type_ssl` being equal to `false`").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"ipsec_enable_anti_replay_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipsec_anti_replay_window_size": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("security Association Lifetime, Attribute conditional on `ipsec_enable_anti_replay` being equal to `true`").AddDefaultValueDescription("64").String,
				Optional:            true,
			},
			"ipsec_anti_replay_window_size_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipsec_security_association_lifetime": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Security Association Lifetime in Seconds, Attribute conditional on `connection_type_ssl` being equal to `false`").AddIntegerRangeDescription(3600, 86400).AddDefaultValueDescription("3600").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(3600, 86400),
				},
			},
			"ipsec_security_association_lifetime_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipsec_enable_perfect_foward_secrecy": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("security Association Lifetime, Attribute conditional on `connection_type_ssl` being equal to `false`").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ipsec_enable_perfect_foward_secrecy_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
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
	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
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
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
