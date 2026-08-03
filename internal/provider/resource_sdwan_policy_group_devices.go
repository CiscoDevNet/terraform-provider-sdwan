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
	"strings"
	"sync"
	"time"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &PolicyGroupDevicesResource{}
var _ resource.ResourceWithImportState = &PolicyGroupDevicesResource{}

func NewPolicyGroupDevicesResource() resource.Resource {
	return &PolicyGroupDevicesResource{}
}

type PolicyGroupDevicesResource struct {
	client            *sdwan.Client
	updateMutex       *sync.Mutex
	taskTimeout       *int64
	deployOnOutOfDate bool
}

func (r *PolicyGroupDevicesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_group_devices"
}

func (r *PolicyGroupDevicesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Policy Group Devices .").AddMinimumVersionDescription("20.15.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"policy_group_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the policy group to associate the devices with").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"solution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of solution").AddStringEnumDescription("sdwan").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sdwan"),
				},
			},
			"policy_group_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The version of the associated policy group. Set this to `sdwan_policy_group.<name>.version` so that any change to the policy group triggers a redeployment of the devices managed by this resource.").String,
				Optional:            true,
			},
			"devices": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of devices").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Device ID").String,
							Optional:            true,
						},
						"deploy": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Deploy to device if enabled.").AddDefaultValueDescription("false").String,
							Optional:            true,
							Computed:            true,
							Default:             booldefault.StaticBool(false),
						},
						"variables": schema.SetNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of variables").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Required:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable value").String,
										Optional:            true,
									},
									"list_value": schema.ListAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Use this instead of `value` in case value is of type `List`.").String,
										ElementType:         types.StringType,
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

func (r *PolicyGroupDevicesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
	r.taskTimeout = req.ProviderData.(*SdwanProviderData).TaskTimeout
	r.deployOnOutOfDate = req.ProviderData.(*SdwanProviderData).DeployOnOutOfDate
}

// End of section. //template:end model

func (r *PolicyGroupDevicesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PolicyGroupDevices

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.PolicyGroupId.ValueString()))

	plan.Id = plan.PolicyGroupId

	// Associate devices
	if len(plan.Devices) > 0 {
		body := plan.toBodyPolicyGroupDevices(ctx)
		path := fmt.Sprintf("/v1/policy-group/%v/device/associate/", plan.PolicyGroupId.ValueString())
		res, err := r.associateDevices(ctx, "POST", path, body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure policy group devices (POST), got error: %s, %s", err, res.String()))
			return
		}
	}

	// Need to get existing variables so SD-WAN Manager refreshes database correctly
	if len(plan.Devices) > 0 {
		r.client.Get(fmt.Sprintf("/v1/policy-group/%v/device/variables/", plan.PolicyGroupId.ValueString()))
	}

	// Set device variables
	if len(plan.Devices) > 0 && plan.hasPolicyGroupDeviceVariables(ctx) {
		varTypes := r.getDeviceVariableTypes(ctx, plan.PolicyGroupId.ValueString())
		body := plan.toBodyPolicyGroupDeviceVariables(ctx, varTypes)
		path := fmt.Sprintf("/v1/policy-group/%v/device/variables/", plan.PolicyGroupId.ValueString())
		res, err := r.client.Put(path, body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure policy group device variables (PUT), got error: %s, %s", err, res.String()))
			r.disassociateDevices(ctx, plan, &resp.Diagnostics)
			return
		}
	}

	// Deploy to devices
	if len(plan.Devices) > 0 {
		r.Deploy(ctx, plan, nil, &resp.Diagnostics)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.PolicyGroupId.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// getDeviceVariableTypes fetches the variable schema from the API and extracts type information
func (r *PolicyGroupDevicesResource) getDeviceVariableTypes(ctx context.Context, policyGroupId string) map[string]string {
	varTypes := make(map[string]string)

	path := fmt.Sprintf("/v1/policy-group/%v/device/variables/schema", policyGroupId)
	res, err := r.client.Get(path)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("Failed to fetch variable schema: %s", err))
		return varTypes
	}

	for _, item := range res.Array() {
		variablesArray := item.Get("variables")
		if !variablesArray.Exists() {
			continue
		}
		for _, variable := range variablesArray.Array() {
			schemaProps := variable.Get("schema.properties")
			if !schemaProps.Exists() {
				continue
			}
			schemaProps.ForEach(func(varName, varSchema gjson.Result) bool {
				if valueType := varSchema.Get("properties.value.type"); valueType.Exists() {
					typeStr := valueType.String()
					if typeStr == "array" {
						if itemsType := varSchema.Get("properties.value.items.type"); itemsType.Exists() {
							typeStr = itemsType.String()
						}
					}
					varTypes[varName.String()] = typeStr
				}
				return true
			})
		}
	}

	return varTypes
}

// disassociateDevices disassociates all of this resource's own devices from the policy group. In normal
// operation the devices passed here are genuinely associated with this group (Read is state-scoped), so
// the DELETE simply removes them and returns success; any real API error is surfaced.
func (r *PolicyGroupDevicesResource) disassociateDevices(ctx context.Context, data PolicyGroupDevices, diag *diag.Diagnostics) {
	if len(data.Devices) == 0 {
		return
	}
	body, _ := sjson.Set("", "devices", []interface{}{})
	for _, d := range data.Devices {
		itemBody, _ := sjson.Set("", "id", d.Id.ValueString())
		body, _ = sjson.SetRaw(body, "devices.-1", itemBody)
	}
	path := fmt.Sprintf("/v1/policy-group/%v/device/associate/", data.PolicyGroupId.ValueString())
	res, err := r.client.DeleteBody(path, body)
	if err != nil {
		diag.AddError("Client Error", fmt.Sprintf("Failed to disassociate policy group devices (DELETE), got error: %s, %s", err, res.String()))
	}
}

// associateDevices performs an associate (POST) or update-associate (PUT) call, retrying on the
// "device already associated to another group" conflict. When a device is being moved between groups,
// Terraform may run the disassociation (from the old group) and this association (to the new group)
// concurrently; retrying gives the concurrent disassociation time to complete without this resource
// ever touching the other group itself.
func (r *PolicyGroupDevicesResource) associateDevices(ctx context.Context, method, path, body string) (sdwan.Res, error) {
	const maxWait = 30 * time.Second
	const interval = 3 * time.Second
	deadline := time.Now().Add(maxWait)

	for {
		var res sdwan.Res
		var err error
		if method == "PUT" {
			res, err = r.client.Put(path, body)
		} else {
			res, err = r.client.Post(path, body)
		}
		if err == nil {
			return res, nil
		}
		// The conflict detail carries the "already associated to group(s)" text under error.details.
		conflict := strings.Contains(res.Get("error.details").String(), "already associated") || res.Get("error.code").String() == "CFGRP0018"
		if !conflict || time.Now().After(deadline) {
			return res, err
		}
		tflog.Debug(ctx, fmt.Sprintf("Device(s) still associated to another group, retrying association in %s", interval))
		select {
		case <-ctx.Done():
			return res, ctx.Err()
		case <-time.After(interval):
		}
	}
}

func (r *PolicyGroupDevicesResource) Deploy(ctx context.Context, plan PolicyGroupDevices, state *PolicyGroupDevices, diag *diag.Diagnostics) {
	var updatedDevices []string
	if state != nil {
		updatedDevices = plan.getUpdatedDevices(ctx, state)
	}

	// Force a redeploy when the associated policy group version changed.
	hasVersionChange := false
	if state != nil {
		hasVersionChange = !plan.PolicyGroupVersion.Equal(state.PolicyGroupVersion)
	}

	path := fmt.Sprintf("/v1/policy-group/%v/device/associate/", plan.PolicyGroupId.ValueString())
	res, err := r.client.Get(path)
	if err != nil {
		diag.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	body, _ := sjson.Set("", "devices", []interface{}{})
	if value := res.Get("devices"); value.Exists() && len(value.Array()) > 0 {
		value.ForEach(func(k, v gjson.Result) bool {
			id := v.Get("id").String()
			for _, item := range plan.Devices {
				if item.Id.ValueString() == id && item.Deploy.ValueBool() && (!v.Get("policyGroupUpToDate").Bool() || updatedDevices == nil || helpers.Contains(updatedDevices, id) || hasVersionChange) {
					itemBody, _ := sjson.Set("", "id", id)
					body, _ = sjson.SetRaw(body, "devices.-1", itemBody)
					tflog.Debug(ctx, fmt.Sprintf("%s: Deploying to device %s", plan.PolicyGroupId.ValueString(), id))
				}
			}
			return true
		})
	}
	if len(gjson.Get(body, "devices").Array()) > 0 {
		path := fmt.Sprintf("/v1/policy-group/%v/device/deploy/", plan.PolicyGroupId.ValueString())
		res, err = r.client.Post(path, body)
		if err != nil {
			diag.AddError("Client Error", fmt.Sprintf("Failed to deploy to policy group devices (POST), got error: %s, %s", err, res.String()))
			return
		}
		actionId := res.Get("parentTaskId").String()
		err, _ = helpers.WaitForActionToComplete(ctx, r.client, actionId, r.taskTimeout)
		if err != nil {
			diag.AddError("Client Error", fmt.Sprintf("Failed to deploy to policy group devices, got error: %s", err))
			return
		}
	}
}

func (r *PolicyGroupDevicesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PolicyGroupDevices

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.PolicyGroupId.ValueString()))

	oldState := state

	// Read policy group device associations (filtered to this resource's own devices). A deleted policy
	// group surfaces here as "Invalid policy group passed" and removes the resource.
	path := fmt.Sprintf("/v1/policy-group/%v/device/associate/", state.PolicyGroupId.ValueString())
	res, err := r.client.Get(path)
	if strings.Contains(res.Get("error.message").String(), "Invalid policy group passed") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	associateRes := res
	state.fromBodyPolicyGroupDevices(ctx, res)

	// Read policy group device variables (fromBody only updates devices we own)
	if len(state.Devices) > 0 {
		path = fmt.Sprintf("/v1/policy-group/%v/device/variables/", state.PolicyGroupId.ValueString())
		res, err = r.client.Get(path)
		if strings.Contains(res.Get("error.message").String(), "Invalid policy group passed") {
			resp.State.RemoveResource(ctx)
			return
		} else if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}
		state.fromBodyPolicyGroupDeviceVariables(ctx, res)
	}

	state.updateTfAttributes(ctx, &oldState)

	if r.deployOnOutOfDate {
		if value := associateRes.Get("devices"); value.Exists() {
			value.ForEach(func(k, v gjson.Result) bool {
				// An offline / not-yet-synced device always reports policyGroupUpToDate=false with a
				// "Sync Pending" status; that is not a genuine out-of-date condition, so skip it to
				// avoid a perpetual redeploy diff.
				if !v.Get("policyGroupUpToDate").Bool() && v.Get("configStatusMessage").String() != "Sync Pending" {
					id := v.Get("id").String()
					for i := range state.Devices {
						if state.Devices[i].Id.ValueString() == id {
							state.Devices[i].Deploy = types.BoolValue(false)
							break
						}
					}
				}
				return true
			})
		}
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}
	if imp {
		state.processImport(ctx)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.PolicyGroupId.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *PolicyGroupDevicesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PolicyGroupDevices

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.PolicyGroupId.ValueString()))

	res, err := r.client.Get(fmt.Sprintf("/v1/policy-group/%v/device/associate/", plan.PolicyGroupId.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	var currentDeviceIds []string
	if value := res.Get("devices"); value.Exists() {
		value.ForEach(func(k, v gjson.Result) bool {
			currentDeviceIds = append(currentDeviceIds, v.Get("id").String())
			return true
		})
	}

	// Associate devices new in the plan.
	associateBody, _ := sjson.Set("", "devices", []interface{}{})
	for _, d := range plan.Devices {
		found := false
		for _, cdid := range currentDeviceIds {
			if d.Id.ValueString() == cdid {
				found = true
				break
			}
		}
		if !found {
			associateBody, _ = sjson.SetRaw(associateBody, "devices.-1", helpers.Must(sjson.Set("", "id", d.Id.ValueString())))
		}
	}

	// Disassociate only devices that were in this resource's state but are no longer in the plan.
	disassociateBody, _ := sjson.Set("", "devices", []interface{}{})
	for _, sd := range state.Devices {
		found := false
		for _, d := range plan.Devices {
			if d.Id.ValueString() == sd.Id.ValueString() {
				found = true
				break
			}
		}
		if !found {
			disassociateBody, _ = sjson.SetRaw(disassociateBody, "devices.-1", helpers.Must(sjson.Set("", "id", sd.Id.ValueString())))
		}
	}

	// Disassociate removed devices first, so a device being moved out of this group is freed before we
	// associate a device being moved in (avoids a deadlock when two resources swap devices concurrently).
	if len(gjson.Get(disassociateBody, "devices").Array()) > 0 {
		res, err = r.client.DeleteBody(fmt.Sprintf("/v1/policy-group/%v/device/associate/", plan.PolicyGroupId.ValueString()), disassociateBody)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to disassociate policy group devices (DELETE), got error: %s, %s", err, res.String()))
			return
		}
	}

	// Associate missing devices (retries on the "already associated to another group" conflict).
	if len(gjson.Get(associateBody, "devices").Array()) > 0 {
		res, err = r.associateDevices(ctx, "PUT", fmt.Sprintf("/v1/policy-group/%v/device/associate/", plan.PolicyGroupId.ValueString()), associateBody)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure policy group devices (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	// Need to get existing variables so SD-WAN Manager refreshes database correctly
	if len(plan.Devices) > 0 {
		r.client.Get(fmt.Sprintf("/v1/policy-group/%v/device/variables/", plan.PolicyGroupId.ValueString()))
	}

	// Update device variables
	if len(plan.Devices) > 0 && plan.hasPolicyGroupDeviceVariables(ctx) {
		varTypes := r.getDeviceVariableTypes(ctx, plan.PolicyGroupId.ValueString())
		body := plan.toBodyPolicyGroupDeviceVariables(ctx, varTypes)
		res, err = r.client.Put(fmt.Sprintf("/v1/policy-group/%v/device/variables/", plan.PolicyGroupId.ValueString()), body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure policy group device variables (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	if len(plan.Devices) > 0 {
		r.Deploy(ctx, plan, &state, &resp.Diagnostics)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.PolicyGroupId.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PolicyGroupDevicesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PolicyGroupDevices

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.PolicyGroupId.ValueString()))

	// Disassociate only this resource's own devices. The policy group itself is managed by the
	// sdwan_policy_group resource and is intentionally left untouched.
	r.disassociateDevices(ctx, state, &resp.Diagnostics)

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.PolicyGroupId.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *PolicyGroupDevicesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.SplitN(req.ID, ",[", 2)
	if len(parts) != 2 || !strings.HasSuffix(parts[1], "]") {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: policy_group_id,[deviceId1,deviceId2,...]. Got: %q", req.ID),
		)
		return
	}

	policyGroupId := parts[0]
	deviceIds := strings.Split(strings.TrimSuffix(parts[1], "]"), ",")
	if len(policyGroupId) == 0 || len(deviceIds) == 0 {
		resp.Diagnostics.AddError(
			"Invalid Import Data", "Policy group ID and at least one device ID must be provided",
		)
		return
	}

	var devices []PolicyGroupDevicesDevices
	for _, deviceId := range deviceIds {
		devices = append(devices, PolicyGroupDevicesDevices{
			Id: types.StringValue(deviceId),
		})
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), policyGroupId)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("policy_group_id"), policyGroupId)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("devices"), devices)...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
