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
	"github.com/hashicorp/go-version"
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
var _ resource.Resource = &TransportCellularProfileProfileParcelResource{}
var _ resource.ResourceWithImportState = &TransportCellularProfileProfileParcelResource{}

func NewTransportCellularProfileProfileParcelResource() resource.Resource {
	return &TransportCellularProfileProfileParcelResource{}
}

type TransportCellularProfileProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *TransportCellularProfileProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transport_cellular_profile_feature"
}

func (r *TransportCellularProfileProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Transport Cellular Profile Feature.").AddMinimumVersionDescription("20.15.0").String,

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
			"profile_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set Profile ID").AddIntegerRangeDescription(1, 16).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 16),
				},
			},
			"profile_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"access_point_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set access point name").String,
				Optional:            true,
			},
			"access_point_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"requires_authentication": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Require authentication type").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"authentication_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set authentication type, Attribute conditional on `requires_authentication` equal to `true`").AddStringEnumDescription("pap", "chap", "pap_chap").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("pap", "chap", "pap_chap"),
				},
			},
			"authentication_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `requires_authentication` equal to `true`").String,
				Optional:            true,
			},
			"profile_username": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the profile username, Attribute conditional on `requires_authentication` equal to `true`").String,
				Optional:            true,
			},
			"profile_username_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `requires_authentication` equal to `true`").String,
				Optional:            true,
			},
			"profile_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the profile password [Note: Catalyst SD-WAN Manager will encrypt this field before saving. Cleartext strings will not be returned back to the user in GET responses for sensitive fields.], Attribute conditional on `requires_authentication` equal to `true`").AddMutualExclusivityDescription("Only one of `profile_password`, `profile_password_variable` and `profile_password_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `profile_password_wo` together with `profile_password_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"profile_password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `requires_authentication` equal to `true`").String,
				Optional:            true,
			},
			"profile_password_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the profile password [Note: Catalyst SD-WAN Manager will encrypt this field before saving. Cleartext strings will not be returned back to the user in GET responses for sensitive fields.], Attribute conditional on `requires_authentication` equal to `true`").AddMutualExclusivityDescription("Only one of `profile_password`, `profile_password_variable` and `profile_password_wo` can be set.").String,
				Sensitive:           true,
				WriteOnly:           true,
				Optional:            true,
			},
			"profile_password_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `profile_password_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"packet_data_network_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set packet data network type").AddStringEnumDescription("ipv4", "ipv4v6", "ipv6").AddDefaultValueDescription("ipv4").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ipv4", "ipv4v6", "ipv6"),
				},
			},
			"packet_data_network_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"no_overwrite": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("No Overwrite").String,
				Optional:            true,
			},
			"no_overwrite_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"slice_type": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("S-NSSAI slice type number: 1(eMBB), 2(URLLC), 3(MioT), Attribute conditional on SD-WAN Manager version `20.18.1` or higher").String,
				Optional:            true,
			},
			"slice_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on SD-WAN Manager version `20.18.1` or higher").String,
				Optional:            true,
			},
			"slice_differentiator": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("S-NSSAI slice differentiator, Attribute conditional on `slice_type` being set and SD-WAN Manager version `20.18.1` or higher").AddIntegerRangeDescription(0, 16777214).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtMost(16777214),
				},
			},
			"slice_differentiator_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `slice_type` being set and SD-WAN Manager version `20.18.1` or higher").String,
				Optional:            true,
			},
		},
	}
}

func (r *TransportCellularProfileProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

var _ resource.ResourceWithValidateConfig = &TransportCellularProfileProfileParcelResource{}

// ValidateConfig enforces the rules between a secret attribute, its "_variable"
// sibling and its write-only twin.
//
// All diagnostics here are resource-scoped (AddError/AddWarning, never
// AddAttributeError) and the schema deliberately carries no
// ConflictsWith/ExactlyOneOf/AlsoRequires validators. Terraform renders an
// attribute-scoped diagnostic together with the offending raw configuration line,
// which would print the secret into plan output and CI logs. Sensitive: true does
// not suppress that.
func (r *TransportCellularProfileProfileParcelResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {

	var vProfilePassword types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("profile_password"), &vProfilePassword)...)
	var vProfilePasswordVariable types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("profile_password_variable"), &vProfilePasswordVariable)...)
	var vProfilePasswordWo types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("profile_password_wo"), &vProfilePasswordWo)...)
	var vProfilePasswordWoVersion types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("profile_password_wo_version"), &vProfilePasswordWoVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An unknown value is one that is not resolved yet at validate time, for example an
	// ephemeral resource attribute or a random_password reference. Treat it as "not
	// decided" rather than "set", otherwise every such configuration trips a false error.
	if !vProfilePassword.IsUnknown() && !vProfilePasswordWo.IsUnknown() && !vProfilePassword.IsNull() && !vProfilePasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `profile_password` and `profile_password_wo` can be set.",
		)
	}
	if !vProfilePasswordVariable.IsUnknown() && !vProfilePasswordWo.IsUnknown() && !vProfilePasswordVariable.IsNull() && !vProfilePasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"`profile_password_variable` and `profile_password_wo` cannot both be set. `profile_password_variable` defers the value to a device variable, so the write-only value would never be sent.",
		)
	}
	if !vProfilePassword.IsUnknown() && !vProfilePasswordVariable.IsUnknown() && !vProfilePassword.IsNull() && !vProfilePasswordVariable.IsNull() {
		resp.Diagnostics.AddWarning(
			"Conflicting Attribute Combination",
			"Both `profile_password` and `profile_password_variable` are set. `profile_password_variable` takes precedence, so the value of `profile_password` is not sent to the device - but it is still stored in Terraform state. Set only one of them.",
		)
	}
	if !vProfilePasswordWo.IsUnknown() && !vProfilePasswordWoVersion.IsUnknown() && !vProfilePasswordWo.IsNull() && vProfilePasswordWoVersion.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`profile_password_wo_version` must be set when `profile_password_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !vProfilePasswordWoVersion.IsUnknown() && !vProfilePasswordWo.IsUnknown() && !vProfilePasswordWoVersion.IsNull() && vProfilePasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`profile_password_wo` must be set when `profile_password_wo_version` is used.",
		)
	}
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *TransportCellularProfileProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TransportCellularProfile

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "profile_password_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because a parcel PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("profile_password_wo"), &plan.ProfilePasswordWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object

	ver := version.Must(version.NewVersion(r.client.ManagerVersion))
	body := plan.toBody(ctx, ver)

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
func (r *TransportCellularProfileProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state TransportCellularProfile

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
func (r *TransportCellularProfileProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state TransportCellularProfile

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "profile_password_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because a parcel PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("profile_password_wo"), &plan.ProfilePasswordWo)...)
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

	ver := version.Must(version.NewVersion(r.client.ManagerVersion))
	body := plan.toBody(ctx, ver)
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
func (r *TransportCellularProfileProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TransportCellularProfile

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
func (r *TransportCellularProfileProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "transport_cellular_profile_feature_id" + ",feature_profile_id"
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
