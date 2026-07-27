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
	"regexp"
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
var _ resource.Resource = &OtherTrustSecProfileParcelResource{}
var _ resource.ResourceWithImportState = &OtherTrustSecProfileParcelResource{}

func NewOtherTrustSecProfileParcelResource() resource.Resource {
	return &OtherTrustSecProfileParcelResource{}
}

type OtherTrustSecProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *OtherTrustSecProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_other_trustsec_feature"
}

func (r *OtherTrustSecProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Other TrustSec Feature.").AddMinimumVersionDescription("20.18.0").String,

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
			"device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify the TrustSec Network Access Device ID, should be same as mentioned in the Identity Services Engine (upto 32 char)").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"device_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"device_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the password for the device").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 24),
				},
			},
			"device_password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"device_sgt": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Local device security group <2..65519>").AddIntegerRangeDescription(2, 65519).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 65519),
				},
			},
			"device_sgt_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"enable_enforcement": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Role-based Access Control enforcement").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"enable_enforcement_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"enable_sxp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable CTS SXP support").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"listener_hold_time_min": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Minimum allowed hold-time for listener in seconds <1..65534>").AddIntegerRangeDescription(1, 65534).AddDefaultValueDescription("90").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65534),
				},
			},
			"listener_hold_time_min_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"listener_hold_time_max": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Maximum allowed hold-time for listener in seconds <1..65534>").AddIntegerRangeDescription(1, 65534).AddDefaultValueDescription("180").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65534),
				},
			},
			"listener_hold_time_max_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"speaker_hold_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Speaker hold-time in seconds <1..65534>").AddIntegerRangeDescription(1, 65534).AddDefaultValueDescription("120").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65534),
				},
			},
			"speaker_hold_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_default_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure SXP default password").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 24),
				},
			},
			"sxp_default_password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_key_chain": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure SXP key-chain").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(256),
				},
			},
			"sxp_key_chain_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_log_binding_changes": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enables logging for IP-to-SGT binding changes").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"sxp_log_binding_changes_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_reconciliation_period": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the SXP reconciliation period in seconds <0..64000>").AddIntegerRangeDescription(0, 64000).AddDefaultValueDescription("120").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtMost(64000),
				},
			},
			"sxp_reconciliation_period_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_retry_period": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Retry period for SXP connection in seconds <0..64000>").AddIntegerRangeDescription(0, 64000).AddDefaultValueDescription("120").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtMost(64000),
				},
			},
			"sxp_retry_period_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_source_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SXP Source IP").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)\.){3}(25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)$`), ""),
				},
			},
			"sxp_source_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_connections": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure SXP Connections").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"peer_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure SXP Peer IP address (IPv4)").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)\.){3}(25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)$`), ""),
							},
						},
						"peer_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure SXP Source IP address (IPv4)").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)\.){3}(25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)$`), ""),
							},
						},
						"source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"preshared_key": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Define Preshared Key type").AddStringEnumDescription("password", "none", "key chain").AddDefaultValueDescription("none").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("password", "none", "key chain"),
							},
						},
						"mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Define Mode of connection").AddStringEnumDescription("local", "peer").AddDefaultValueDescription("local").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("local", "peer"),
							},
						},
						"mode_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Define Role of a device <speaker/listener/both>").AddStringEnumDescription("speaker", "listener", "both").AddDefaultValueDescription("speaker").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("speaker", "listener", "both"),
							},
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Connection VPN (VRF) ID").AddIntegerRangeDescription(0, 65527).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(65527),
							},
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"min_hold_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Connection Minimum hold time <0..65535>, Attribute conditional on `mode_type` not equal to `both`").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(65535),
							},
						},
						"min_hold_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `mode_type` not equal to `both`").String,
							Optional:            true,
						},
						"max_hold_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Connection Maximum hold time <0..65535>, Attribute conditional on (`mode` equal to `peer` and `mode_type` equal to `speaker`) or (`mode` equal to `local` and `mode_type` equal to `listener`)").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(65535),
							},
						},
						"max_hold_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on (`mode` equal to `peer` and `mode_type` equal to `speaker`) or (`mode` equal to `local` and `mode_type` equal to `listener`)").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *OtherTrustSecProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *OtherTrustSecProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan OtherTrustSec

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
func (r *OtherTrustSecProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state OtherTrustSec

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
func (r *OtherTrustSecProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state OtherTrustSec

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
func (r *OtherTrustSecProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OtherTrustSec

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
func (r *OtherTrustSecProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "other_trustsec_feature_id" + ",feature_profile_id"
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
