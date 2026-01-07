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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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

func NewAttachFeatureDeviceTemplateResource() resource.Resource {
	return &AttachFeatureDeviceTemplateResource{}
}

type AttachFeatureDeviceTemplateResource struct {
	client      *sdwan.Client
	taskTimeout *int64
}

func (r *AttachFeatureDeviceTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_attach_feature_device_template"
}

func (r *AttachFeatureDeviceTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can attach a feature device template. Due to limitations of the API, once a device template is attached to a device, only one change can be applied per `terraform apply` operation. More information is available [here](https://registry.terraform.io/providers/CiscoDevNet/sdwan/latest/docs/guides/updating_templates).").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the device template",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the device template",
				Optional:            true,
			},
			"devices": schema.ListNestedAttribute{
				MarkdownDescription: "Devices",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Device ID",
							Required:            true,
						},
						"variables": schema.MapAttribute{
							MarkdownDescription: "Device variables",
							ElementType:         types.StringType,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *AttachFeatureDeviceTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.taskTimeout = req.ProviderData.(*SdwanProviderData).TaskTimeout
}

func (r *AttachFeatureDeviceTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AttachFeatureDeviceTemplate

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body, err := plan.toBody(ctx, r.client, false, nil)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to render body, got error: %s", err))
		return
	}
	res, err := r.client.Post("/template/device/config/attachfeature", body)
	if strings.Contains(res.Get("error.message").String(), "Template edit request has expired") {
		resp.Diagnostics.AddWarning("Client Warning", fmt.Sprintf("No changes detected to trigger an attachment."))
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	if resp.Diagnostics.WarningsCount() == 0 {
		actionId := res.Get("id").String()
		err, warnings := helpers.WaitForActionToComplete(ctx, r.client, actionId, r.taskTimeout)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to attach device template, got error: %s", err))
			return
		} else if warnings != "" {
			resp.Diagnostics.AddWarning("Client Warning", warnings)
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *AttachFeatureDeviceTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AttachFeatureDeviceTemplate

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	err := state.readVariables(ctx, r.client)
	if err != nil && strings.Contains(err.Error(), "StatusCode 400") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *AttachFeatureDeviceTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state AttachFeatureDeviceTemplate

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

	edited := false
	if !plan.Version.Equal(state.Version) {
		edited = true
	}

	body, err := plan.toBody(ctx, r.client, edited, &state)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to render body, got error: %s", err))
		return
	}

	if body != "" {
		res, err := r.client.Post("/template/device/config/attachfeature", body)
		if strings.Contains(res.Get("error.message").String(), "Template edit request has expired") {
			resp.Diagnostics.AddWarning("Client Warning", fmt.Sprintf("No changes detected to trigger an attachment."))
		} else if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
			return
		}

		if resp.Diagnostics.WarningsCount() == 0 {
			actionId := res.Get("id").String()
			err, warnings := helpers.WaitForActionToComplete(ctx, r.client, actionId, r.taskTimeout)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to attach device template, got error: %s", err))
				return
			} else if warnings != "" {
				resp.Diagnostics.AddWarning("Client Warning", warnings)
			}
		}
	}

	// Removes Detached Devices
	var detachedDevices []string
	for _, stateDevice := range state.Devices {
		found := false
		for _, device := range plan.Devices {
			if stateDevice.Id.ValueString() == device.Id.ValueString() {
				found = true
				break
			}
		}
		if !found {
			detachedDevices = append(detachedDevices, stateDevice.Id.ValueString())
		}
	}

	if len(detachedDevices) > 0 {
		res, err := state.detachDevices(ctx, r.client, &plan, []string{})
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve currently attached devices, got error: %s", err))
			return
		}
		if res.Get("id").Exists() {
			err, warnings := helpers.WaitForActionToComplete(ctx, r.client, res.Get("id").String(), r.taskTimeout)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to detach device template, got error: %s", err))
				return
			} else if warnings != "" {
				resp.Diagnostics.AddWarning("Client Warning", warnings)
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *AttachFeatureDeviceTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var plan, state AttachFeatureDeviceTemplate

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	res, err := state.detachDevices(ctx, r.client, &plan, []string{})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve currently attached devices, got error: %s", err))
		return
	}
	if res.Get("id").Exists() {
		err, warnings := helpers.WaitForActionToComplete(ctx, r.client, res.Get("id").String(), r.taskTimeout)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to detach device template, got error: %s", err))
			return
		} else if warnings != "" {
			resp.Diagnostics.AddWarning("Client Warning", warnings)
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *AttachFeatureDeviceTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var _ AttachFeatureDeviceTemplateDevice

	parts := strings.SplitN(req.ID, ",[", 2)
	if len(parts) != 2 || !strings.HasSuffix(parts[1], "]") {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: templateId,[deviceId1,deviceId2,...]. Got: %q", req.ID),
		)
		return
	}

	templateId := parts[0]
	deviceIds := strings.Split(strings.TrimSuffix(parts[1], "]"), ",")
	if len(templateId) == 0 || len(deviceIds) == 0 {
		resp.Diagnostics.AddError(
			"Invalid Import Data", fmt.Sprintf("Template ID and at least one device ID must be provided"),
		)
		return
	}

	variables, diags := types.MapValue(types.StringType, map[string]attr.Value{})
	if diags != nil {
		resp.Diagnostics.Append(diags...)
		return
	}

	var devices []AttachFeatureDeviceTemplateDevice
	for _, deviceId := range deviceIds {
		devices = append(devices, AttachFeatureDeviceTemplateDevice{
			Id:        types.StringValue(deviceId),
			Variables: variables,
		})
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), templateId)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("devices"), devices)...)
}
