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
var _ resource.Resource = &SystemSecurityProfileParcelResource{}
var _ resource.ResourceWithImportState = &SystemSecurityProfileParcelResource{}

func NewSystemSecurityProfileParcelResource() resource.Resource {
	return &SystemSecurityProfileParcelResource{}
}

type SystemSecurityProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *SystemSecurityProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_security_feature"
}

func (r *SystemSecurityProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a System Security Feature.").AddMinimumVersionDescription("20.12.0").String,

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
			"rekey": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set how often to change the AES key for DTLS connections").AddIntegerRangeDescription(10, 1209600).AddDefaultValueDescription("86400").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 1209600),
				},
			},
			"rekey_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"anti_replay_window": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the sliding replay window size").AddStringEnumDescription("64", "128", "256", "512", "1024", "2048", "4096", "8192").AddDefaultValueDescription("512").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("64", "128", "256", "512", "1024", "2048", "4096", "8192"),
				},
			},
			"anti_replay_window_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"extended_anti_replay_window": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Extended Anti-Replay Window").AddIntegerRangeDescription(10, 2048).AddDefaultValueDescription("256").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 2048),
				},
			},
			"extended_anti_replay_window_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipsec_pairwise_keying": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable IPsec pairwise-keying").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ipsec_pairwise_keying_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"integrity_type": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the authentication type for DTLS connections").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"integrity_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"keychains": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure a Keychain").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key_chain_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify the name of the Keychain").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthAtMost(236),
							},
						},
						"key_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify the Key ID").AddIntegerRangeDescription(0, 2147483647).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(2147483647),
							},
						},
					},
				},
			},
			"keys": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure a Key").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select the Key ID").String,
							Optional:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select the chain name").String,
							Optional:            true,
						},
						"send_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify the Send ID").AddIntegerRangeDescription(0, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(255),
							},
						},
						"send_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"receiver_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify the Receiver ID").AddIntegerRangeDescription(0, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(255),
							},
						},
						"receiver_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"include_tcp_options": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Include TCP Options").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"include_tcp_options_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"accept_ao_mismatch": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Accept AO Mismatch").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"accept_ao_mismatch_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"crypto_algorithm": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Crypto Algorithm").AddStringEnumDescription("aes-128-cmac", "hmac-sha-1", "hmac-sha-256").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("aes-128-cmac", "hmac-sha-1", "hmac-sha-256"),
							},
						},
						"key_string": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify the Key String").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
						},
						"key_string_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"send_life_time_local": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Send lifetime Local").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"send_life_time_local_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"send_life_time_start_epoch": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Key lifetime start time").String,
							Optional:            true,
						},
						"send_life_time_infinite": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Infinite lifetime").String,
							Optional:            true,
						},
						"send_life_time_infinite_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"send_life_time_duration": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Send lifetime Duration (seconds)").AddIntegerRangeDescription(1, 2147483646).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 2147483646),
							},
						},
						"send_life_time_duration_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"send_life_time_exact": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Key lifetime end time").String,
							Optional:            true,
						},
						"accept_life_time_local": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Send lifetime Local").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"accept_life_time_local_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"accept_life_time_start_epoch": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Key lifetime start time").String,
							Optional:            true,
						},
						"accept_life_time_infinite": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Infinite lifetime").String,
							Optional:            true,
						},
						"accept_life_time_infinite_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"accept_life_time_duration": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Send lifetime Duration (seconds)").AddIntegerRangeDescription(1, 2147483646).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 2147483646),
							},
						},
						"accept_life_time_duration_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"accept_life_time_exact": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Key lifetime end time").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *SystemSecurityProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *SystemSecurityProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemSecurity

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
func (r *SystemSecurityProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemSecurity

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

	if state.Version.IsNull() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *SystemSecurityProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemSecurity

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
func (r *SystemSecurityProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemSecurity

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
func (r *SystemSecurityProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "system_security_feature_id" + ",feature_profile_id"
	if len(parts) != (count + 1) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: %s. Got: %q", pattern, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("feature_profile_id"), parts[1])...)
}

// End of section. //template:end import
