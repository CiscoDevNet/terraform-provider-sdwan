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
var _ resource.Resource = &TransportRoutePolicyProfileParcelResource{}
var _ resource.ResourceWithImportState = &TransportRoutePolicyProfileParcelResource{}

func NewTransportRoutePolicyProfileParcelResource() resource.Resource {
	return &TransportRoutePolicyProfileParcelResource{}
}

type TransportRoutePolicyProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *TransportRoutePolicyProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transport_route_policy_profile_parcel"
}

func (r *TransportRoutePolicyProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Transport Route Policy profile parcel.").AddMinimumVersionDescription("20.12.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the profile parcel",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the profile parcel",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the profile parcel",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the profile parcel",
				Optional:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Feature Profile ID").String,
				Required:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default Action").AddStringEnumDescription("reject", "accept").AddDefaultValueDescription("reject").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("reject", "accept"),
				},
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Route Policy List").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence Id").AddIntegerRangeDescription(1, 65536).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65536),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence Name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 19),
							},
						},
						"base_action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Base Action").AddStringEnumDescription("reject", "accept").AddDefaultValueDescription("reject").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("reject", "accept"),
							},
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("protocol such as IPV4, IPV6, or BOTH").AddStringEnumDescription("IPV4", "IPV6", "BOTH").AddDefaultValueDescription("IPV4").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("IPV4", "IPV6", "BOTH"),
							},
						},
						"match_entries": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Define match conditions").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"as_path_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"standard_community_list_criteria": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Select a condition such as OR, AND or EXACT").AddStringEnumDescription("OR", "AND", "EXACT").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("OR", "AND", "EXACT"),
										},
									},
									"standard_community_lists": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Select a standard community list").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
													},
												},
											},
										},
									},
									"expanded_community_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"extended_community_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"bgp_local_preference": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("BGP Local Preference").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.AtMost(4294967295),
										},
									},
									"metric": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Select Metric").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.AtMost(4294967295),
										},
									},
									"omp_tag": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Select OMP Tag").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.AtMost(4294967295),
										},
									},
									"ospf_tag": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Select OSPF Tag").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.AtMost(4294967295),
										},
									},
									"ipv4_address_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"ipv4_next_hop_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"ipv6_address_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"ipv6_next_hop_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
								},
							},
						},
						"actions": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Define list of actions").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"as_path_prend": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.Int64Type,
										Optional:            true,
									},
									"community_additive": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").AddDefaultValueDescription("false").String,
										Optional:            true,
									},
									"community": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"community_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"local_preference": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set Local Preference").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.AtMost(4294967295),
										},
									},
									"metric": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set Metric").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.AtMost(4294967295),
										},
									},
									"metric_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set Metric Type").AddStringEnumDescription("type1", "type2").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("type1", "type2"),
										},
									},
									"omp_tag": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set OMP Tag").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.AtMost(4294967295),
										},
									},
									"origin": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set Origin").AddStringEnumDescription("EGP", "IGP", "Incomplete").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("EGP", "IGP", "Incomplete"),
										},
									},
									"ospf_tag": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set OSPF Tag").AddIntegerRangeDescription(0, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.AtMost(4294967295),
										},
									},
									"weight": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set Weight").AddIntegerRangeDescription(0, 65535).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.AtMost(65535),
										},
									},
									"ipv4_next_hop": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set Ipv4 Next Hop").String,
										Optional:            true,
									},
									"ipv6_next_hop": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set Ipv6 Next Hop").String,
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

func (r *TransportRoutePolicyProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *TransportRoutePolicyProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TransportRoutePolicy

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
func (r *TransportRoutePolicyProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state TransportRoutePolicy

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
func (r *TransportRoutePolicyProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state TransportRoutePolicy

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
func (r *TransportRoutePolicyProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TransportRoutePolicy

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
func (r *TransportRoutePolicyProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
