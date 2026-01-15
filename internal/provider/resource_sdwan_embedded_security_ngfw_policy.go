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
var _ resource.Resource = &EmbeddedSecurityNGFWProfileParcelResource{}
var _ resource.ResourceWithImportState = &EmbeddedSecurityNGFWProfileParcelResource{}

func NewEmbeddedSecurityNGFWProfileParcelResource() resource.Resource {
	return &EmbeddedSecurityNGFWProfileParcelResource{}
}

type EmbeddedSecurityNGFWProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *EmbeddedSecurityNGFWProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_embedded_security_ngfw_policy"
}

func (r *EmbeddedSecurityNGFWProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Embedded Security NGFW Policy.").AddMinimumVersionDescription("20.15.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Policy",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Policy",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Policy",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Policy",
				Optional:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Feature Profile ID").String,
				Required:            true,
			},
			"default_actione": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("pass", "drop").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("pass", "drop"),
				},
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"sequence_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^(0|[1-9][0-9]{0,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$`), ""),
							},
						},
						"rule_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 128),
							},
						},
						"base_action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("pass", "inspect", "drop").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("pass", "inspect", "drop"),
							},
						},
						"rule_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`ngfirewall`), ""),
							},
						},
						"disable_rule": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
						},
						"matches": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"source_data_prefix_list_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"destination_data_prefix_list_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"destination_fqdn_list_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"source_geo_location_list_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"destination_geo_location_list_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"source_port_list_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"destination_port_list_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"source_scalable_group_tag_list_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"destination_scalable_group_tag_list_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"source_indentity_list_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"protocol_name_list_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"app_list_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"flat_app_list_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"source_security_group_list_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"destination_security_group_list_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"source_data_prefixs": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"source_data_prefixs_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"destination_data_prefixs": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"destination_data_prefixs_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"destination_fqdns": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"destination_fqdns_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"source_ports": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"source_ports_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"destination_ports": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"destination_ports_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"source_geo_locations": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"source_geo_locations_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"destination_geo_locations": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"destination_geo_locations_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"source_identity_users": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"source_identity_usergroups": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"applications": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"application_families": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"protocols": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"protocol_names": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
								},
							},
						},
						"actions": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("can be empty array or with type or parameter").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("log", "connectionEvents").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("log", "connectionEvents"),
										},
									},
									"parameter": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
									},
									"parameter_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
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

func (r *EmbeddedSecurityNGFWProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *EmbeddedSecurityNGFWProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan EmbeddedSecurityNGFW

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
func (r *EmbeddedSecurityNGFWProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state EmbeddedSecurityNGFW

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
func (r *EmbeddedSecurityNGFWProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state EmbeddedSecurityNGFW

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
func (r *EmbeddedSecurityNGFWProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state EmbeddedSecurityNGFW

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
func (r *EmbeddedSecurityNGFWProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "embedded_security_ngfw_policy_id" + ",feature_profile_id"
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
