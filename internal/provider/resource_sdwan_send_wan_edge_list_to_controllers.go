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
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &WANEdgeCertificateSendResource{}

func NewWANEdgeCertificateSendResource() resource.Resource {
	return &WANEdgeCertificateSendResource{}
}

type WANEdgeCertificateSendResource struct {
	client      *sdwan.Client
	taskTimeout *int64
}

func (r *WANEdgeCertificateSendResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_send_wan_edge_list_to_controllers"
}

func (r *WANEdgeCertificateSendResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource sends the current WAN edge certificate list to the controllers. It does not manage the certificate validity of any device, use the `sdwan_wan_edge_certificate_validate` resource for that. The initial create or any change to `version` issues a new push; use `version` to connect this resource to a certificate validation resource.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Placeholder identifier attribute",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "A version value that, when changed, triggers a new push of the WAN edge certificate list to the controllers",
				Optional:            true,
			},
		},
	}
}

func (r *WANEdgeCertificateSendResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.taskTimeout = req.ProviderData.(*SdwanProviderData).TaskTimeout
}

func (r *WANEdgeCertificateSendResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WANEdgeCertificatePush

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Beginning Create of WAN edge certificate push")

	if err := plan.push(ctx, r.client, r.taskTimeout); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("failed to send the WAN edge list to the controllers (POST), got error: %s", err))
		return
	}
	plan.Id = types.StringValue("push")

	tflog.Debug(ctx, "Create finished successfully for WAN edge certificate push")

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *WANEdgeCertificateSendResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// The push is a one-time action against the Manager with no queryable state, so state is
	// left untouched here.
	var state WANEdgeCertificatePush

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *WANEdgeCertificateSendResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan WANEdgeCertificatePush

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Beginning Update of WAN edge certificate push")

	if err := plan.push(ctx, r.client, r.taskTimeout); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("failed to send the WAN edge list to the controllers (POST), got error: %s", err))
		return
	}
	plan.Id = types.StringValue("push")

	tflog.Debug(ctx, "Update finished successfully for WAN edge certificate push")

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *WANEdgeCertificateSendResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Removing this resource does not un-push or otherwise change controller state.
	resp.State.RemoveResource(ctx)
}
