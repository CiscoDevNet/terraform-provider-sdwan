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

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &WANEdgeCertificateValidateResource{}
var _ resource.ResourceWithImportState = &WANEdgeCertificateValidateResource{}

func NewWANEdgeCertificateValidateResource() resource.Resource {
	return &WANEdgeCertificateValidateResource{}
}

type WANEdgeCertificateValidateResource struct {
	client *sdwan.Client
}

func (r *WANEdgeCertificateValidateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wan_edge_certificate_validate"
}

func (r *WANEdgeCertificateValidateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the certificate validity of a WAN edge device (e.g. cEdge). The device must already be present in the WAN edge list of the Manager, this resource only changes its certificate state. This resource does not send the updated WAN edge list to the controllers, use the `sdwan_send_wan_edge_list_to_controllers` resource for that.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The chassis number of the WAN edge device",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"chassis_number": schema.StringAttribute{
				MarkdownDescription: "The chassis number (UUID) of the WAN edge device",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"serial_number": schema.StringAttribute{
				MarkdownDescription: "The certificate serial number of the WAN edge device, by default the serial number known by the Manager is used",
				Optional:            true,
				Computed:            true,
			},
			"validity": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The certificate validity of the WAN edge device").AddStringEnumDescription("invalid", "staging", "valid").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("invalid", "staging", "valid"),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "A state version that changes after a successful certificate validity update and can trigger the controller push resource",
				Computed:            true,
			},
		},
	}
}

func (r *WANEdgeCertificateValidateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
}

// apply resolves the serial number and saves the validity. It does not push the updated WAN
// edge list to the controllers, use the sdwan_send_wan_edge_list_to_controllers resource for that.
func (r *WANEdgeCertificateValidateResource) apply(ctx context.Context, data *WANEdgeCertificate, validity string) error {
	entry, found, err := data.getCertificate(ctx, r.client)
	if err != nil {
		return fmt.Errorf("failed to retrieve WAN edge list (GET), got error: %s", err)
	}
	if !found {
		return fmt.Errorf("chassis number %s not found in the WAN edge list of the Manager", data.ChassisNumber.ValueString())
	}

	if data.SerialNumber.IsNull() || data.SerialNumber.IsUnknown() || data.SerialNumber.ValueString() == "" {
		data.SerialNumber = types.StringValue(entry.Get("serialNumber").String())
	}
	if data.SerialNumber.ValueString() == "" {
		return fmt.Errorf("chassis number %s has no serial number in the WAN edge list of the Manager", data.ChassisNumber.ValueString())
	}

	if err := data.setValidity(ctx, r.client, validity); err != nil {
		return fmt.Errorf("failed to set certificate validity (POST), got error: %s", err)
	}
	return nil
}

func (r *WANEdgeCertificateValidateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WANEdgeCertificate

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.ChassisNumber.ValueString()))

	if err := r.apply(ctx, &plan, plan.Validity.ValueString()); err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}
	plan.Id = plan.ChassisNumber
	plan.Version = types.Int64Value(1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.ChassisNumber.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *WANEdgeCertificateValidateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WANEdgeCertificate

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	entry, found, err := state.getCertificate(ctx, r.client)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	state.fromBody(ctx, entry)

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	if imp {
		state.processImport(ctx)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *WANEdgeCertificateValidateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state WANEdgeCertificate

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.ChassisNumber.ValueString()))

	if err := r.apply(ctx, &plan, plan.Validity.ValueString()); err != nil {
		resp.Diagnostics.AddError("Client Error", err.Error())
		return
	}
	plan.Id = plan.ChassisNumber
	if state.Version.IsNull() || state.Version.IsUnknown() {
		plan.Version = types.Int64Value(1)
	} else {
		plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.ChassisNumber.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *WANEdgeCertificateValidateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

func (r *WANEdgeCertificateValidateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("chassis_number"), req, resp)
	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
