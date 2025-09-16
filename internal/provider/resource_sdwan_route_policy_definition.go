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
var _ resource.Resource = &RoutePolicyDefinitionResource{}
var _ resource.ResourceWithImportState = &RoutePolicyDefinitionResource{}

func NewRoutePolicyDefinitionResource() resource.Resource {
	return &RoutePolicyDefinitionResource{}
}

type RoutePolicyDefinitionResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
	taskTimeout *int64
}

func (r *RoutePolicyDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route_policy_definition"
}

func (r *RoutePolicyDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Route Policy Definition .").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Type",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the policy definition").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The description of the policy definition").String,
				Required:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default action, either `accept` or `reject`").AddStringEnumDescription("accept", "reject").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("accept", "reject"),
				},
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of ACL sequences").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence ID").AddIntegerRangeDescription(1, 65534).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65534),
							},
						},
						"ip_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP version, either `ipv4` or `ipv6`").AddStringEnumDescription("ipv4", "ipv6").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ipv4", "ipv6"),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence name").String,
							Required:            true,
						},
						"base_action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Base action, either `accept` or `reject`").AddStringEnumDescription("accept", "reject").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("accept", "reject"),
							},
						},
						"match_entries": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of match entries").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of match entry").AddStringEnumDescription("address", "asPath", "community", "advancedCommunity", "expandedCommunity", "expandedCommunityInline", "extCommunity", "localPreference", "metric", "nextHop", "origin", "peer", "ompTag", "ospfTag").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("address", "asPath", "community", "advancedCommunity", "expandedCommunity", "expandedCommunityInline", "extCommunity", "localPreference", "metric", "nextHop", "origin", "peer", "ompTag", "ospfTag"),
										},
									},
									"prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Prefix list ID, Attribute conditional on `type` being equal to `address`").String,
										Optional:            true,
									},
									"prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Prefix list version").String,
										Optional:            true,
									},
									"as_path_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("AS path list ID, Attribute conditional on `type` being equal to `asPath`").String,
										Optional:            true,
									},
									"as_path_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("AS path list version").String,
										Optional:            true,
									},
									"community_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Community list ID, Attribute conditional on `type` being equal to `community`").String,
										Optional:            true,
									},
									"community_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Community list version").String,
										Optional:            true,
									},
									"community_list_match_flag": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Community list match flag").AddStringEnumDescription("and", "or", "exact").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("and", "or", "exact"),
										},
									},
									"community_list_ids": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Community list IDs, Attribute conditional on `type` being equal to `advancedCommunity`").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"community_list_versions": schema.ListAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Community list versions").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"expanded_community_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Expanded community list ID, Attribute conditional on `type` being equal to `expandedCommunity`").String,
										Optional:            true,
									},
									"expanded_community_list_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Expanded community list variable, Attribute conditional on `type` being equal to `expandedCommunityInline`").String,
										Optional:            true,
									},
									"expanded_community_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Expanded community list version").String,
										Optional:            true,
									},
									"extended_community_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Extended community list ID, Attribute conditional on `type` being equal to `extCommunity`").String,
										Optional:            true,
									},
									"extended_community_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Extended community list version").String,
										Optional:            true,
									},
									"local_preference": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Local preference, Attribute conditional on `type` being equal to `localPreference`").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
									"metric": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Metric, Attribute conditional on `type` being equal to `metric`").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
									"next_hop_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Next hop prefix list ID, Attribute conditional on `type` being equal to `nextHop`").String,
										Optional:            true,
									},
									"next_hop_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Next hop prefix list version").String,
										Optional:            true,
									},
									"origin": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Origin, Attribute conditional on `type` being equal to `origin`").AddStringEnumDescription("igp", "egp", "incomplete").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("igp", "egp", "incomplete"),
										},
									},
									"peer": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Peer IP, Attribute conditional on `type` being equal to `peer`").String,
										Optional:            true,
									},
									"omp_tag": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("OMP tag, Attribute conditional on `type` being equal to `ompTag`").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
									"ospf_tag": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("OSPF tag, Attribute conditional on `type` being equal to `ospfTag`").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
								},
							},
						},
						"action_entries": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of action entries").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of action entry").AddStringEnumDescription("aggregator", "asPath", "atomicAggregate", "community", "communityAdditive", "localPreference", "metric", "weight", "metricType", "nextHop", "ompTag", "ospfTag", "origin", "originator").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("aggregator", "asPath", "atomicAggregate", "community", "communityAdditive", "localPreference", "metric", "weight", "metricType", "nextHop", "ompTag", "ospfTag", "origin", "originator"),
										},
									},
									"aggregator": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Aggregator, Attribute conditional on `type` being equal to `aggregator`").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
									"aggregator_ip_address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP address, Attribute conditional on `type` being equal to `aggregator`").String,
										Optional:            true,
									},
									"as_path_prepend": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Space separated list of ASN to prepend, Attribute conditional on `type` being equal to `asPath`").String,
										Optional:            true,
									},
									"as_path_exclude": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Space separated list of ASN to exclude, Attribute conditional on `type` being equal to `asPath`").String,
										Optional:            true,
									},
									"atomic_aggregate": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Atomic aggregate, Attribute conditional on `type` being equal to `atomicAggregate`").String,
										Optional:            true,
									},
									"community": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Community value, e.g. `1000:10000` or `internet` or `local-AS`, Attribute conditional on `type` being equal to `community`").String,
										Optional:            true,
									},
									"community_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Community variable, Attribute conditional on `type` being equal to `community`").String,
										Optional:            true,
									},
									"community_additive": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Community additive, Attribute conditional on `type` being equal to `communityAdditive`").String,
										Optional:            true,
									},
									"local_preference": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Local preference, Attribute conditional on `type` being equal to `localPreference`").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
									"metric": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Metric, Attribute conditional on `type` being equal to `metric`").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
									"weight": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Weight, Attribute conditional on `type` being equal to `weight`").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
									"metric_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Metric type, Attribute conditional on `type` being equal to `metricType`").AddStringEnumDescription("type1", "type2").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("type1", "type2"),
										},
									},
									"next_hop": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Next hop IP, Attribute conditional on `type` being equal to `nextHop`").String,
										Optional:            true,
									},
									"omp_tag": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("OMP tag, Attribute conditional on `type` being equal to `ompTag`").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
									"ospf_tag": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("OSPF tag, Attribute conditional on `type` being equal to `ospfTag`").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 4294967295),
										},
									},
									"origin": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Origin, Attribute conditional on `type` being equal to `origin`").AddStringEnumDescription("igp", "egp", "incomplete").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("igp", "egp", "incomplete"),
										},
									},
									"originator": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Originator IP, Attribute conditional on `type` being equal to `originator`").String,
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

func (r *RoutePolicyDefinitionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
	r.taskTimeout = req.ProviderData.(*SdwanProviderData).TaskTimeout
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *RoutePolicyDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutePolicyDefinition

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
	plan.Id = types.StringValue(res.Get("definitionId").String())
	plan.Version = types.Int64Value(0)
	plan.Type = types.StringValue("vedgeRoute")

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *RoutePolicyDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutePolicyDefinition

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.ValueString()))

	res, err := r.client.Get(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if strings.Contains(res.Get("error.message").String(), "Failed to find specified resource") || strings.Contains(res.Get("error.message").String(), "Invalid template type") || strings.Contains(res.Get("error.message").String(), "Template definition not found") || strings.Contains(res.Get("error.message").String(), "Invalid Profile Id") || strings.Contains(res.Get("error.message").String(), "Invalid feature Id") || strings.Contains(res.Get("error.message").String(), "Invalid config group passed") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)
	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	if imp {
		state.processImport(ctx)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *RoutePolicyDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutePolicyDefinition

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

	if plan.hasChanges(ctx, &state) {
		body := plan.toBody(ctx)
		r.updateMutex.Lock()
		res, err := r.client.Put(plan.getPath()+url.QueryEscape(plan.Id.ValueString()), body)
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
		tflog.Debug(ctx, fmt.Sprintf("%s: No changes detected", plan.Name.ValueString()))
	}
	plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *RoutePolicyDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutePolicyDefinition

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	_, _ = r.client.Get(state.getPath())
	res, err := r.client.Delete(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *RoutePolicyDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
