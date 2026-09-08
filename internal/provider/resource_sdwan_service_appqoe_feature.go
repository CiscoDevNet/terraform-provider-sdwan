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
var _ resource.Resource = &ServiceAppQoEProfileParcelResource{}
var _ resource.ResourceWithImportState = &ServiceAppQoEProfileParcelResource{}

func NewServiceAppQoEProfileParcelResource() resource.Resource {
	return &ServiceAppQoEProfileParcelResource{}
}

type ServiceAppQoEProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *ServiceAppQoEProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_appqoe_feature"
}

func (r *ServiceAppQoEProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("`appqoe_device_role` selects which subtree is sent: `forwarder_*` attributes apply only to the `forwarder` role, `combined_*` only to `forwarderAndServiceNode` and `forwarderAndServiceNodeWithDre`, and `service_node_*` only to `serviceNode` and `serviceNodeWithDre`. Attributes belonging to a role other than the configured one are silently omitted from the request and will show as a recurring difference in `terraform plan`. DRE optimization is enabled by selecting a `...WithDre` role, which also unlocks `virtual_applications` for the DRE resource profile. Only one service node group can be bound per service context.").AddMinimumVersionDescription("20.15.0").String,

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
			"appqoe_device_role": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Appqoe Device Role").AddStringEnumDescription("forwarder", "forwarderAndServiceNode", "serviceNode", "serviceNodeWithDre", "forwarderAndServiceNodeWithDre").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("forwarder", "forwarderAndServiceNode", "serviceNode", "serviceNodeWithDre", "forwarderAndServiceNodeWithDre"),
				},
			},
			"virtual_applications": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Virtual application Instance, Attribute conditional on `appqoe_device_role` containing `Dre`").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"resource_profile": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Resource Profile").AddStringEnumDescription("small", "medium", "large", "extra-large", "default").AddDefaultValueDescription("default").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("small", "medium", "large", "extra-large", "default"),
							},
						},
						"resource_profile_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"forwarder_controller_groups": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Appnav controller group name, Attribute conditional on `appqoe_device_role` equal to `forwarder`").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"appnav_controllers": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of controllers").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Controller IP Address").String,
										Optional:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"vpn": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("vpn id").AddIntegerRangeDescription(1, 65530).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 65530),
										},
									},
								},
							},
						},
					},
				},
			},
			"forwarder_service_node_groups": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name, Attribute conditional on `appqoe_device_role` equal to `forwarder`").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of service node group").AddDefaultValueDescription("SNG-APPQOE").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`\b(SNG-APPQOE([1-9]|[12][0-9]|3[0-1])|SNG-APPQOE)\b`), ""),
							},
						},
						"service_nodes": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Service Node Information").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP Address").String,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
			"forwarder_service_contexts": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Appqoe, Attribute conditional on `appqoe_device_role` equal to `forwarder`").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"appnav_controller_group": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Appnav controller group").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 128),
								stringvalidator.RegexMatches(regexp.MustCompile(`^[^&<>! "]+$`), ""),
							},
						},
						"service_node_group": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Service node group").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 128),
								stringvalidator.RegexMatches(regexp.MustCompile(`^[^&<>! "]+$`), ""),
							},
						},
						"enable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("enable service context").String,
							Optional:            true,
						},
						"vpn": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Vpn").String,
							Optional:            true,
						},
						"vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"combined_controller_groups": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Appnav controller group name, Attribute conditional on `appqoe_device_role` containing `forwarderAndServiceNode`").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of controller group").AddDefaultValueDescription("ACG-APPQOE").String,
							Optional:            true,
						},
						"appnav_controllers": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of controllers").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Controller IP Address").AddDefaultValueDescription("192.168.2.1").String,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
			"combined_service_node_groups": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name, Attribute conditional on `appqoe_device_role` containing `forwarderAndServiceNode`").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of service node group").AddDefaultValueDescription("SNG-APPQOE").String,
							Optional:            true,
						},
						"service_nodes": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Service Node Information").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP Address").AddDefaultValueDescription("192.168.2.2").String,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
			"combined_service_contexts": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Appqoe, Attribute conditional on `appqoe_device_role` containing `forwarderAndServiceNode`").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"appnav_controller_group": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Appnav controller group").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 128),
								stringvalidator.RegexMatches(regexp.MustCompile(`^[^&<>! "]+$`), ""),
							},
						},
						"service_node_group": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Service node group").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 128),
								stringvalidator.RegexMatches(regexp.MustCompile(`^[^&<>! "]+$`), ""),
							},
						},
						"enable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("enable service context").String,
							Optional:            true,
						},
						"vpn": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Vpn").String,
							Optional:            true,
						},
						"vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"service_node_service_node_groups": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name, Attribute conditional on `appqoe_device_role` containing `serviceNode`").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of service node group").AddDefaultValueDescription("SNG-APPQOE").String,
							Optional:            true,
						},
						"service_nodes": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Service Node Information").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP Address").AddDefaultValueDescription("192.168.2.2").String,
										Optional:            true,
									},
									"vpg_ip": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("ip and prefix").AddDefaultValueDescription("192.168.2.1/24").String,
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

func (r *ServiceAppQoEProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *ServiceAppQoEProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ServiceAppQoE

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
func (r *ServiceAppQoEProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ServiceAppQoE

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
func (r *ServiceAppQoEProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ServiceAppQoE

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
func (r *ServiceAppQoEProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ServiceAppQoE

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
func (r *ServiceAppQoEProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "service_appqoe_feature_id" + ",feature_profile_id"
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
