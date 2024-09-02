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
var _ resource.Resource = &IPv4ACLPolicyDefinitionResource{}
var _ resource.ResourceWithImportState = &IPv4ACLPolicyDefinitionResource{}

func NewIPv4ACLPolicyDefinitionResource() resource.Resource {
	return &IPv4ACLPolicyDefinitionResource{}
}

type IPv4ACLPolicyDefinitionResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *IPv4ACLPolicyDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv4_acl_policy_definition"
}

func (r *IPv4ACLPolicyDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a IPv4 ACL Policy Definition .").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Default action, either `accept` or `drop`").AddStringEnumDescription("accept", "drop").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("accept", "drop"),
				},
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of ACL sequences").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence ID").AddIntegerRangeDescription(1, 65534).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65534),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence name").String,
							Required:            true,
						},
						"base_action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Base action, either `accept` or `drop`").AddStringEnumDescription("accept", "drop").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("accept", "drop"),
							},
						},
						"match_entries": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of match entries").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of match entry").AddStringEnumDescription("dscp", "sourceIp", "destinationIp", "class", "packetLength", "plp", "sourcePort", "destinationPort", "sourceDataPrefixList", "destinationDataPrefixList", "protocol", "tcp", "icmpMessage").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("dscp", "sourceIp", "destinationIp", "class", "packetLength", "plp", "sourcePort", "destinationPort", "sourceDataPrefixList", "destinationDataPrefixList", "protocol", "tcp", "icmpMessage"),
										},
									},
									"dscp": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("DSCP value").AddIntegerRangeDescription(0, 63).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 63),
										},
									},
									"source_ip": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source IP prefix").String,
										Optional:            true,
									},
									"icmp_message": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("ICMP Message").String,
										Optional:            true,
									},
									"destination_ip": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination IP prefix").String,
										Optional:            true,
									},
									"class_map_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Class map ID").String,
										Optional:            true,
									},
									"class_map_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Class map version").String,
										Optional:            true,
									},
									"packet_length": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Packet length").AddIntegerRangeDescription(0, 65535).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 65535),
										},
									},
									"priority": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("PLP - priority").AddStringEnumDescription("high", "low").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("high", "low"),
										},
									},
									"source_ports": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source ports. Single value (0-65535) or ranges separated by spaces.").String,
										Optional:            true,
									},
									"destination_ports": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination ports. Single value (0-65535) or ranges separated by spaces.").String,
										Optional:            true,
									},
									"source_data_ipv4_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source data IPv4 prefix list ID").String,
										Optional:            true,
									},
									"source_data_ipv4_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source data IPv4 prefix list version").String,
										Optional:            true,
									},
									"destination_data_ipv4_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination data IPv4 prefix list ID").String,
										Optional:            true,
									},
									"destination_data_ipv4_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination data IPv4 prefix list version").String,
										Optional:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Single value (0-255) or multiple values separated by spaces").String,
										Optional:            true,
									},
									"tcp": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("TCP parameters").AddStringEnumDescription("syn").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("syn"),
										},
									},
								},
							},
						},
						"action_entries": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of action entries").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of action entry").AddStringEnumDescription("class", "count", "set", "log", "mirror", "policer").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("class", "count", "set", "log", "mirror", "policer"),
										},
									},
									"class_map_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Class map ID").String,
										Optional:            true,
									},
									"class_map_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Class map version").String,
										Optional:            true,
									},
									"counter_name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Counter name").String,
										Optional:            true,
									},
									"log": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Enable logging").String,
										Optional:            true,
									},
									"mirror_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Mirror ID").String,
										Optional:            true,
									},
									"mirror_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Mirror version").String,
										Optional:            true,
									},
									"policer_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Policer ID").String,
										Optional:            true,
									},
									"policer_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Policer version").String,
										Optional:            true,
									},
									"set_parameters": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("List of set parameters").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Type of set parameter").AddStringEnumDescription("dscp", "nextHop").String,
													Required:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("dscp", "nextHop"),
													},
												},
												"dscp": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("DSCP value").AddIntegerRangeDescription(0, 63).String,
													Optional:            true,
													Validators: []validator.Int64{
														int64validator.Between(0, 63),
													},
												},
												"next_hop": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Next hop IP").String,
													Optional:            true,
												},
											},
										},
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

func (r *IPv4ACLPolicyDefinitionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *IPv4ACLPolicyDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPv4ACLPolicyDefinition

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
	plan.Type = types.StringValue("acl")

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *IPv4ACLPolicyDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPv4ACLPolicyDefinition

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if strings.Contains(res.Get("error.message").String(), "Failed to find specified resource") || strings.Contains(res.Get("error.message").String(), "Invalid template type") || strings.Contains(res.Get("error.message").String(), "Template definition not found") || strings.Contains(res.Get("error.message").String(), "Invalid Profile Id") || strings.Contains(res.Get("error.message").String(), "Invalid feature Id") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *IPv4ACLPolicyDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPv4ACLPolicyDefinition

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
func (r *IPv4ACLPolicyDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPv4ACLPolicyDefinition

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
func (r *IPv4ACLPolicyDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
