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
var _ resource.Resource = &ServiceIPv4ACLProfileParcelResource{}
var _ resource.ResourceWithImportState = &ServiceIPv4ACLProfileParcelResource{}

func NewServiceIPv4ACLProfileParcelResource() resource.Resource {
	return &ServiceIPv4ACLProfileParcelResource{}
}

type ServiceIPv4ACLProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *ServiceIPv4ACLProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_ipv4_acl_profile_parcel"
}

func (r *ServiceIPv4ACLProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Service IPv4 ACL profile parcel.").AddMinimumVersionDescription("20.12.0").String,

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
				Optional:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default Action").AddStringEnumDescription("drop", "accept").AddDefaultValueDescription("drop").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("drop", "accept"),
				},
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Access Control List").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"sequence_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence Id").AddIntegerRangeDescription(1, 65536).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65536),
							},
						},
						"sequence_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence Name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 19),
							},
						},
						"base_action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Base Action").AddStringEnumDescription("drop", "accept").AddDefaultValueDescription("accept").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("drop", "accept"),
							},
						},
						"conditions": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Define match conditions").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"dscps": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("DSCP number").String,
										ElementType:         types.Int64Type,
										Optional:            true,
									},
									"packet_length": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Packet Length").String,
										Optional:            true,
									},
									"protocols": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("protocol number list with at least one item").String,
										ElementType:         types.Int64Type,
										Optional:            true,
									},
									"icmp_messages": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("ICMP Message").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"source_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"source_data_prefix": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source Data IP Prefix").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(/)([0-2]?[0-9]$|[3]?[0-2])`), ""),
										},
									},
									"source_data_prefix_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"source_ports": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source Port List").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"port": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("source port range or individual port number").String,
													Optional:            true,
												},
											},
										},
									},
									"destination_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"destination_data_prefix": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination Data IP Prefix").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(/)([0-2]?[0-9]$|[3]?[0-2])`), ""),
										},
									},
									"destination_data_prefix_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"destination_ports": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination Port List").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"port": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("destination port range or individual port number").String,
													Optional:            true,
												},
											},
										},
									},
									"tcp_state": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("TCP States").AddStringEnumDescription("syn").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("syn"),
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
									"accept_dscp": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("DSCP number").AddIntegerRangeDescription(0, 63).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.AtMost(63),
										},
									},
									"accept_counter_name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Counter Name").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(1, 20),
										},
									},
									"accept_log": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Enable Log").AddDefaultValueDescription("false").String,
										Optional:            true,
									},
									"accept_next_hop": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set Next Hop (IPV4 address)").String,
										Optional:            true,
									},
									"accept_mirror_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"accept_policer_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"drop_counter_name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Counter Name").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(1, 20),
										},
									},
									"drop_log": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Enable Log").AddDefaultValueDescription("false").String,
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

func (r *ServiceIPv4ACLProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *ServiceIPv4ACLProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ServiceIPv4ACL

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
func (r *ServiceIPv4ACLProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ServiceIPv4ACL

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
func (r *ServiceIPv4ACLProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ServiceIPv4ACL

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
func (r *ServiceIPv4ACLProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ServiceIPv4ACL

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
func (r *ServiceIPv4ACLProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
