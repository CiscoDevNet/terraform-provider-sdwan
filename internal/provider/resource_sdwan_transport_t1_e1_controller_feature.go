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
var _ resource.Resource = &TransportT1E1ControllerProfileParcelResource{}
var _ resource.ResourceWithImportState = &TransportT1E1ControllerProfileParcelResource{}

func NewTransportT1E1ControllerProfileParcelResource() resource.Resource {
	return &TransportT1E1ControllerProfileParcelResource{}
}

type TransportT1E1ControllerProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *TransportT1E1ControllerProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transport_t1_e1_controller_feature"
}

func (r *TransportT1E1ControllerProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Transport T1 E1 Controller Feature.").AddMinimumVersionDescription("20.12.0").String,

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
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Card Type").AddStringEnumDescription("e1", "t1").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("e1", "t1"),
				},
			},
			"slot": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Slot number").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"slot_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"entries": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Controller tx-ex List").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"t1_description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Card Type").AddStringEnumDescription("T1").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("T1"),
							},
						},
						"t1_framing": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Framing").AddStringEnumDescription("esf", "sf").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("esf", "sf"),
							},
						},
						"t1_framing_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"t1_linecode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("LineCode").AddStringEnumDescription("ami", "b8zs").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ami", "b8zs"),
							},
						},
						"t1_linecode_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"e1_description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Card Type").AddStringEnumDescription("E1").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("E1"),
							},
						},
						"e1_framing": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Framing").AddStringEnumDescription("crc4", "no-crc4").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("crc4", "no-crc4"),
							},
						},
						"e1_framing_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"e1_linecode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("LineCode").AddStringEnumDescription("ami", "hdb3").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ami", "hdb3"),
							},
						},
						"e1_linecode_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"cable_length": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Cable Config").AddStringEnumDescription("short", "long").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("short", "long"),
							},
						},
						"length_short": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("length, Attribute conditional on `cable_length` being equal to `short`").AddStringEnumDescription("110ft", "220ft", "330ft", "440ft", "550ft", "660ft").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("110ft", "220ft", "330ft", "440ft", "550ft", "660ft"),
							},
						},
						"length_short_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"length_long": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("length, Attribute conditional on `cable_length` being equal to `long`").AddStringEnumDescription("-15db", "-22.5db", "-7.5db", "0db").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("-15db", "-22.5db", "-7.5db", "0db"),
							},
						},
						"length_long_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"clock_source": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Clock Source").AddStringEnumDescription("line", "internal", "loop-timed", "network").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("line", "internal", "loop-timed", "network"),
							},
						},
						"line_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Line Mode").AddStringEnumDescription("secondary", "primary").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("secondary", "primary"),
							},
						},
						"line_mode_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Description").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 240),
							},
						},
						"description_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"channel_groups": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Channel Group List").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"channel_group": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Number").AddIntegerRangeDescription(0, 23).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.AtMost(23),
										},
									},
									"channel_group_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"time_slot": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Time slots").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(1, 24),
										},
									},
									"time_slot_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *TransportT1E1ControllerProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *TransportT1E1ControllerProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TransportT1E1Controller

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
func (r *TransportT1E1ControllerProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state TransportT1E1Controller

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
func (r *TransportT1E1ControllerProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state TransportT1E1Controller

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
func (r *TransportT1E1ControllerProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TransportT1E1Controller

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
func (r *TransportT1E1ControllerProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "transport_t1_e1_controller_feature_id" + ",feature_profile_id"
	if len(parts) != (count + 1) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: %s. Got: %q, %q", pattern, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("feature_profile_id"), parts[1])...)
}

// End of section. //template:end import
