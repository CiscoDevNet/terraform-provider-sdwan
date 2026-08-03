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
var _ resource.Resource = &ConfigurationGroupDevicesResource{}
var _ resource.ResourceWithImportState = &ConfigurationGroupDevicesResource{}

func NewConfigurationGroupDevicesResource() resource.Resource {
	return &ConfigurationGroupDevicesResource{}
}

type ConfigurationGroupDevicesResource struct {
	client            *sdwan.Client
	updateMutex       *sync.Mutex
	taskTimeout       *int64
	deployOnOutOfDate bool
}

func (r *ConfigurationGroupDevicesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_configuration_group_devices"
}

func (r *ConfigurationGroupDevicesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Configuration Group Devices .").AddMinimumVersionDescription("20.15.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"configuration_group_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the configuration group to associate the devices with").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"solution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of solution").AddStringEnumDescription("mobility", "sdwan", "nfvirtual").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("mobility", "sdwan", "nfvirtual"),
				},
			},
			"configuration_group_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The version of the associated configuration group. Set this to `sdwan_configuration_group.<name>.version` so that any change to the configuration group triggers a redeployment of the devices managed by this resource.").String,
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
						"topology_label": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Topology label for dual device configuration group (supported from version 20.18.1 onwards)").String,
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

func (r *ConfigurationGroupDevicesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
	r.taskTimeout = req.ProviderData.(*SdwanProviderData).TaskTimeout
	r.deployOnOutOfDate = req.ProviderData.(*SdwanProviderData).DeployOnOutOfDate
}

// End of section. //template:end model

func (r *ConfigurationGroupDevicesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ConfigurationGroupDevices

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.ConfigurationGroupId.ValueString()))

	plan.Id = plan.ConfigurationGroupId

	// Associate devices
	if len(plan.Devices) > 0 {
		body := plan.toBodyConfigGroupDevices(ctx)
		path := fmt.Sprintf("/v1/config-group/%v/device/associate/", plan.ConfigurationGroupId.ValueString())
		res, err := r.associateDevices(ctx, "POST", path, body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure configuration group devices (POST), got error: %s, %s", err, res.String()))
			return
		}
	}

	// Set device variables
	if len(plan.Devices) > 0 && plan.hasConfigGroupDeviceVariables(ctx) {
		varTypes := r.getDeviceVariableTypes(ctx, plan.ConfigurationGroupId.ValueString())
		body := plan.toBodyConfigGroupDeviceVariables(ctx, varTypes)
		path := fmt.Sprintf("/v1/config-group/%v/device/variables/", plan.ConfigurationGroupId.ValueString())
		res, err := r.client.Put(path, body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure configuration group device variables (PUT), got error: %s, %s", err, res.String()))
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

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.ConfigurationGroupId.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// getDeviceVariableTypes fetches the variable schema from the API and extracts type information
func (r *ConfigurationGroupDevicesResource) getDeviceVariableTypes(ctx context.Context, configGroupId string) map[string]string {
	varTypes := make(map[string]string)

	path := fmt.Sprintf("/v1/config-group/%v/device/variables/schema", configGroupId)
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

// disassociateDevices disassociates all of this resource's own devices from the configuration group.
// In normal operation the devices passed here are genuinely associated with this group (Read is
// state-scoped), so the DELETE simply removes them and returns success; any real API error is surfaced.
func (r *ConfigurationGroupDevicesResource) disassociateDevices(ctx context.Context, data ConfigurationGroupDevices, diag *diag.Diagnostics) {
	if len(data.Devices) == 0 {
		return
	}
	body, _ := sjson.Set("", "devices", []interface{}{})
	for _, d := range data.Devices {
		itemBody, _ := sjson.Set("", "id", d.Id.ValueString())
		body, _ = sjson.SetRaw(body, "devices.-1", itemBody)
	}
	path := fmt.Sprintf("/v1/config-group/%v/device/associate/", data.ConfigurationGroupId.ValueString())
	res, err := r.client.DeleteBody(path, body)
	if err != nil {
		diag.AddError("Client Error", fmt.Sprintf("Failed to disassociate config group devices (DELETE), got error: %s, %s", err, res.String()))
	}
}

// associateDevices performs an associate (POST) or update-associate (PUT) call, retrying on the
// "device already associated to another group" conflict (CFGRP0018). When a device is being moved
// between configuration groups, Terraform may run the disassociation (from the old group) and this
// association (to the new group) concurrently as independent resources; retrying gives the concurrent
// disassociation time to complete without this resource ever touching the other group itself.
func (r *ConfigurationGroupDevicesResource) associateDevices(ctx context.Context, method, path, body string) (sdwan.Res, error) {
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
		// Only retry the association conflict; surface every other error immediately. The conflict
		// detail carries the "already associated to group(s)" text under error.details (code CFGRP0018),
		// while error.message is just a generic "Config Group Association error".
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

func (r *ConfigurationGroupDevicesResource) Deploy(ctx context.Context, plan ConfigurationGroupDevices, state *ConfigurationGroupDevices, diag *diag.Diagnostics) {
	var updatedDevices []string
	if state != nil {
		updatedDevices = plan.getUpdatedDevices(ctx, state)
	}

	// Force a redeploy when the associated configuration group version changed.
	hasVersionChange := false
	if state != nil {
		hasVersionChange = !plan.ConfigurationGroupVersion.Equal(state.ConfigurationGroupVersion)
	}

	path := fmt.Sprintf("/v1/config-group/%v/device/associate/", plan.ConfigurationGroupId.ValueString())
	res, err := r.client.Get(path)
	if err != nil {
		diag.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// Build deploy body
	body, _ := sjson.Set("", "devices", []interface{}{})
	if value := res.Get("devices"); value.Exists() && len(value.Array()) > 0 {
		value.ForEach(func(k, v gjson.Result) bool {
			id := v.Get("id").String()
			for _, item := range plan.Devices {
				if item.Id.ValueString() == id && item.Deploy.ValueBool() && (!v.Get("configGroupUpToDate").Bool() || updatedDevices == nil || helpers.Contains(updatedDevices, id) || hasVersionChange) {
					itemBody, _ := sjson.Set("", "id", id)
					body, _ = sjson.SetRaw(body, "devices.-1", itemBody)
					tflog.Debug(ctx, fmt.Sprintf("%s: Deploying to device %s", plan.ConfigurationGroupId.ValueString(), id))
				}
			}
			return true
		})
	}
	if len(gjson.Get(body, "devices").Array()) > 0 {
		path := fmt.Sprintf("/v1/config-group/%v/device/deploy/", plan.ConfigurationGroupId.ValueString())
		res, err = r.client.Post(path, body)
		if err != nil {
			diag.AddError("Client Error", fmt.Sprintf("Failed to deploy to config group devices (POST), got error: %s, %s", err, res.String()))
			return
		}

		// Wait for deploy action to complete
		actionId := res.Get("parentTaskId").String()
		err, _ = helpers.WaitForActionToComplete(ctx, r.client, actionId, r.taskTimeout)
		if err != nil {
			diag.AddError("Client Error", fmt.Sprintf("Failed to deploy to config group devices, got error: %s", err))
			return
		}
	}
}

func (r *ConfigurationGroupDevicesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ConfigurationGroupDevices

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.ConfigurationGroupId.ValueString()))

	oldState := state

	// Read config group device associations (filtered to this resource's own devices). A deleted
	// configuration group surfaces here as "Invalid config group passed" and removes the resource,
	// so a separate group-existence GET is unnecessary.
	path := fmt.Sprintf("/v1/config-group/%v/device/associate/", state.ConfigurationGroupId.ValueString())
	res, err := r.client.Get(path)
	if strings.Contains(res.Get("error.message").String(), "Invalid config group passed") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	associateRes := res
	state.fromBodyConfigGroupDevices(ctx, res)

	// Read config group device variables (fromBody only updates devices we own)
	if len(state.Devices) > 0 {
		path = fmt.Sprintf("/v1/config-group/%v/device/variables/", state.ConfigurationGroupId.ValueString())
		res, err = r.client.Get(path)
		if strings.Contains(res.Get("error.message").String(), "Invalid config group passed") {
			resp.State.RemoveResource(ctx)
			return
		} else if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}
		state.fromBodyConfigGroupDeviceVariables(ctx, res, &oldState)
	}

	state.updateTfAttributes(ctx, &oldState)

	if r.deployOnOutOfDate {
		if value := associateRes.Get("devices"); value.Exists() {
			value.ForEach(func(k, v gjson.Result) bool {
				// An offline / not-yet-synced device always reports configGroupUpToDate=false with a
				// "Sync Pending" status. That is not a genuine out-of-date condition, so skip it —
				// otherwise the device would be flagged for redeploy on every refresh, producing a
				// perpetual diff (e.g. on SD-WAN Manager 20.15, where offline devices never report
				// configGroupUpToDate=true).
				if !v.Get("configGroupUpToDate").Bool() && v.Get("configStatusMessage").String() != "Sync Pending" {
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

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.ConfigurationGroupId.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *ConfigurationGroupDevicesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ConfigurationGroupDevices

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.ConfigurationGroupId.ValueString()))

	// Retrieve currently associated devices with their topology labels.
	res, err := r.client.Get(fmt.Sprintf("/v1/config-group/%v/device/associate/", plan.ConfigurationGroupId.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	currentLabels := make(map[string]string)
	if value := res.Get("devices"); value.Exists() {
		value.ForEach(func(k, v gjson.Result) bool {
			currentLabels[v.Get("id").String()] = v.Get("groupTopologyLabel").String()
			return true
		})
	}

	// Associate devices that are new in the plan or whose topology_label changed.
	associateBody, _ := sjson.Set("", "devices", []interface{}{})
	for _, d := range plan.Devices {
		label, exists := currentLabels[d.Id.ValueString()]
		needsAssociate := false
		if !exists {
			needsAssociate = true
		} else {
			planLabel := ""
			if !d.TopologyLabel.IsNull() {
				planLabel = d.TopologyLabel.ValueString()
			}
			if planLabel != label {
				needsAssociate = true
			}
		}
		if needsAssociate {
			itemBody := helpers.Must(sjson.Set("", "id", d.Id.ValueString()))
			if !d.TopologyLabel.IsNull() {
				itemBody = helpers.Must(sjson.Set(itemBody, "groupTopologyLabel", d.TopologyLabel.ValueString()))
			}
			associateBody, _ = sjson.SetRaw(associateBody, "devices.-1", itemBody)
		}
	}

	// Disassociate only devices that were in this resource's state but are no longer in the plan.
	// Devices owned by other resources / state files are never touched.
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

	// Disassociate removed devices first, so that a device being moved out of this group is freed
	// before we try to associate a device being moved in. This avoids a deadlock when two resources
	// swap devices in the same apply: if associate ran first, both would block in their retry loop
	// waiting for the other's disassociate, which would never be reached.
	if len(gjson.Get(disassociateBody, "devices").Array()) > 0 {
		res, err = r.client.DeleteBody(fmt.Sprintf("/v1/config-group/%v/device/associate/", plan.ConfigurationGroupId.ValueString()), disassociateBody)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to disassociate config group devices (DELETE), got error: %s, %s", err, res.String()))
			return
		}
	}

	// Associate missing devices (retries on the "already associated to another group" conflict).
	if len(gjson.Get(associateBody, "devices").Array()) > 0 {
		res, err = r.associateDevices(ctx, "PUT", fmt.Sprintf("/v1/config-group/%v/device/associate/", plan.ConfigurationGroupId.ValueString()), associateBody)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure configuration group devices (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	// Update device variables
	if len(plan.Devices) > 0 && plan.hasConfigGroupDeviceVariables(ctx) {
		varTypes := r.getDeviceVariableTypes(ctx, plan.ConfigurationGroupId.ValueString())
		body := plan.toBodyConfigGroupDeviceVariables(ctx, varTypes)
		path := fmt.Sprintf("/v1/config-group/%v/device/variables/", plan.ConfigurationGroupId.ValueString())
		res, err = r.client.Put(path, body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure configuration group device variables (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	// Deploy to devices
	if len(plan.Devices) > 0 {
		r.Deploy(ctx, plan, &state, &resp.Diagnostics)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.ConfigurationGroupId.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ConfigurationGroupDevicesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ConfigurationGroupDevices

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.ConfigurationGroupId.ValueString()))

	// Disassociate only this resource's own devices. The configuration group itself is managed by the
	// sdwan_configuration_group resource and is intentionally left untouched.
	r.disassociateDevices(ctx, state, &resp.Diagnostics)

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.ConfigurationGroupId.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *ConfigurationGroupDevicesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.SplitN(req.ID, ",[", 2)
	if len(parts) != 2 || !strings.HasSuffix(parts[1], "]") {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: configuration_group_id,[deviceId1,deviceId2,...]. Got: %q", req.ID),
		)
		return
	}

	configGroupId := parts[0]
	deviceIds := strings.Split(strings.TrimSuffix(parts[1], "]"), ",")
	if len(configGroupId) == 0 || len(deviceIds) == 0 {
		resp.Diagnostics.AddError(
			"Invalid Import Data", "Configuration group ID and at least one device ID must be provided",
		)
		return
	}

	var devices []ConfigurationGroupDevicesDevices
	for _, deviceId := range deviceIds {
		devices = append(devices, ConfigurationGroupDevicesDevices{
			Id: types.StringValue(deviceId),
		})
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), configGroupId)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("configuration_group_id"), configGroupId)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("devices"), devices)...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
