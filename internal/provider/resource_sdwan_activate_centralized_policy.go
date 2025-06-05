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

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &AttachFeatureDeviceTemplateResource{}
var _ resource.ResourceWithImportState = &AttachFeatureDeviceTemplateResource{}

func NewActivateCentralizedPolicyResource() resource.Resource {
	return &ActivateCentralizedPolicyResource{}
}

type ActivateCentralizedPolicyResource struct {
	client *sdwan.Client
}

func (r *ActivateCentralizedPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_activate_centralized_policy"
}

func (r *ActivateCentralizedPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can activate a centralized policy.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the centralized policy",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the centralized policy",
				Optional:            true,
			},
		},
	}
}

func (r *ActivateCentralizedPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
}

func (r *ActivateCentralizedPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ActivateCentralizedPolicy

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Activate policy
	res, err := r.client.Post("/template/policy/vsmart/activate/"+plan.Id.ValueString(), "{}")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	actionId := res.Get("id").String()
	err = helpers.WaitForActionToComplete(ctx, r.client, actionId)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to activate centralized policy, got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ActivateCentralizedPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ActivateCentralizedPolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	if state.Id.ValueString() != "" {
		res, err := r.client.Get("/template/policy/vsmart/definition/" + state.Id.ValueString())
		if res.Get("error.message").String() == "Failed to find specified resource" {
			resp.State.RemoveResource(ctx)
			return
		} else if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}

		if !res.Get("isPolicyActivated").Bool() {
			state.Id = types.StringValue("")
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *ActivateCentralizedPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ActivateCentralizedPolicy

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	if !plan.Id.Equal(state.Id) {
		tflog.Debug(ctx, "Policy ID is changing, running activate")
		res, err := r.client.Post("/template/policy/vsmart/activate/"+plan.Id.ValueString(), "{}")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
			return
		}
		actionId := res.Get("id").String()
		err = helpers.WaitForActionToComplete(ctx, r.client, actionId)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to activate centralized policy, got error: %s", err))
			return
		}
	} else {
		tflog.Debug(ctx, "Policy ID is not changing, running repush")

		attachPayload, attachPayloadErr := GetPushBody(r.client)
		if attachPayloadErr != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to build attach payload for vsmart templates, got error: %s", attachPayloadErr))
			return
		}
		res, err := r.client.Post("/template/device/config/attachfeature", attachPayload)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to POST device config attachfeature (POST), got error: %s", err))
			return
		}
		tflog.Debug(ctx, fmt.Sprintf("AttachFeature response for vsmart templates: %s", res.String()))
		if resp.Diagnostics.WarningsCount() == 0 {
			actionId := res.Get("id").String()
			err = helpers.WaitForActionToComplete(ctx, r.client, actionId)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update activated policy, got error: %s", err))
				return
			}
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ActivateCentralizedPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ActivateCentralizedPolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	// Deactivate policy
	if state.Id.ValueString() != "" {
		res, err := r.client.Post("/template/policy/vsmart/deactivate/"+state.Id.ValueString(), "{}")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
			return
		}
		actionId := res.Get("id").String()
		err = helpers.WaitForActionToComplete(ctx, r.client, actionId)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to deactivate centralized policy, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *ActivateCentralizedPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
