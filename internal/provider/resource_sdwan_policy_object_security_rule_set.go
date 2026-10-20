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
var _ resource.Resource = &PolicyObjectSecurityRuleSetProfileParcelResource{}
var _ resource.ResourceWithImportState = &PolicyObjectSecurityRuleSetProfileParcelResource{}

func NewPolicyObjectSecurityRuleSetProfileParcelResource() resource.Resource {
	return &PolicyObjectSecurityRuleSetProfileParcelResource{}
}

type PolicyObjectSecurityRuleSetProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *PolicyObjectSecurityRuleSetProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_object_security_rule_set"
}

func (r *PolicyObjectSecurityRuleSetProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Policy Object Security Rule Set Policy_object.").AddMinimumVersionDescription("20.18.2").String,

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
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("select definition of rule").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 128),
							},
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"sequence_ip_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ipv4/ ipv6").AddStringEnumDescription("ipv4", "ipv6").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ipv4", "ipv6"),
							},
						},
						"sequence_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence Id").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^(0|[1-9][0-9]{0,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$`), ""),
							},
						},
						"sequence_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("action").AddStringEnumDescription("pass", "inspect", "drop").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("pass", "inspect", "drop"),
							},
						},
						"action_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_ipv4_prefixes": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"source_ipv4_prefixes_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_ipv6_prefixes": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"source_ipv6_prefixes_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"destination_ipv4_prefixes": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"destination_ipv4_prefixes_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"destination_ipv6_prefixes": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"destination_ipv6_prefixes_variable": schema.StringAttribute{
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
						"source_geo_locations": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source Geolocation").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"source_geo_locations_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"destination_geo_locations": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination Geolocation").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"destination_geo_locations_variable": schema.StringAttribute{
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
						"protocols": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("protocol").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"protocol_names": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"protocol_names_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_object_group_list_ids": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"destination_object_group_list_ids": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"source_data_ipv4_prefix_list_ids": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"source_data_ipv6_prefix_list_ids": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"destination_data_ipv4_prefix_list_ids": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"destination_data_ipv6_prefix_list_ids": schema.SetAttribute{
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
						"protocol_name_list_ids": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *PolicyObjectSecurityRuleSetProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *PolicyObjectSecurityRuleSetProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PolicyObjectSecurityRuleSet

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
func (r *PolicyObjectSecurityRuleSetProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PolicyObjectSecurityRuleSet

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
	state.fromBody(ctx, res, true)
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
func (r *PolicyObjectSecurityRuleSetProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PolicyObjectSecurityRuleSet

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
func (r *PolicyObjectSecurityRuleSetProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PolicyObjectSecurityRuleSet

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
func (r *PolicyObjectSecurityRuleSetProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "policy_object_security_rule_set_id" + ",feature_profile_id"
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
