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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
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
var _ resource.Resource = &LocalizedPolicyResource{}
var _ resource.ResourceWithImportState = &LocalizedPolicyResource{}

func NewLocalizedPolicyResource() resource.Resource {
	return &LocalizedPolicyResource{}
}

type LocalizedPolicyResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *LocalizedPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_localized_policy"
}

func (r *LocalizedPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Localized Policy .").String,

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
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the localized policy").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The description of the localized policy").String,
				Required:            true,
			},
			"flow_visibility_ipv4": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 flow visibility").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"flow_visibility_ipv6": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 flow visibility").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"application_visibility_ipv4": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 application visibility").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"application_visibility_ipv6": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 application visibility").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"cloud_qos": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cloud QoS").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"cloud_qos_service_side": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cloud QoS service side").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"implicit_acl_logging": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Implicit ACL logging").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"log_frequency": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Log frequency").AddIntegerRangeDescription(1, 2147483647).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 2147483647),
				},
			},
			"ipv4_visibility_cache_entries": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 visibility cache entries").AddIntegerRangeDescription(16, 2000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(16, 2000000),
				},
			},
			"ipv6_visibility_cache_entries": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 visibility cache entries").AddIntegerRangeDescription(16, 2000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(16, 2000000),
				},
			},
			"definitions": schema.SetNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of policy definitions").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Policy definition ID").String,
							Required:            true,
						},
						"version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Policy definition version").String,
							Optional:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Policy definition type").AddStringEnumDescription("qosMap", "rewriteRule", "vedgeRoute", "acl", "aclv6", "deviceAccessPolicy", "deviceAccessPolicyv6").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("qosMap", "rewriteRule", "vedgeRoute", "acl", "aclv6", "deviceAccessPolicy", "deviceAccessPolicyv6"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *LocalizedPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *LocalizedPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan LocalizedPolicy

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
	plan.Id = types.StringValue(res.Get("policyId").String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *LocalizedPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state LocalizedPolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get("/template/policy/vedge/definition/" + url.QueryEscape(state.Id.ValueString()))
	if strings.Contains(res.Get("error.message").String(), "Failed to find specified resource") || strings.Contains(res.Get("error.message").String(), "Invalid template type") || strings.Contains(res.Get("error.message").String(), "Template definition not found") || strings.Contains(res.Get("error.message").String(), "Invalid Profile Id") || strings.Contains(res.Get("error.message").String(), "Invalid feature Id") || strings.Contains(res.Get("error.message").String(), "Invalid config group passed") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)
	if state.Version.IsNull() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *LocalizedPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state LocalizedPolicy

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
func (r *LocalizedPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state LocalizedPolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	_, _ = r.client.Get(state.getPath())
	r.updateMutex.Lock()
	res, err := r.client.Delete(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	r.updateMutex.Unlock()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *LocalizedPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
