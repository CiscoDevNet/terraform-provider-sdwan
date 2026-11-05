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
				MarkdownDescription: helpers.NewAttributeDescription("The version of the associated configuration group. Set this to `sdwan_configuration_group.<name>.version` so that a change to the configuration group is deployed in the same apply: from SD-WAN Manager 20.15 to the devices the Manager reports out of date, on older versions to all devices managed by this resource.").String,
				Optional:            true,
			},
			"retain_devices": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device IDs this resource stops managing without releasing them, for a hand-over to another state. A device removed from `devices` and listed here is dropped from the state but stays associated with the configuration group, so it is not disassociated and needs no deploy. A device can't be in `devices` and in `retain_devices` at the same time. Not needed with a single state.").String,
				ElementType:         types.StringType,
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

var _ resource.ResourceWithValidateConfig = &ConfigurationGroupDevicesResource{}

func (r *ConfigurationGroupDevicesResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	helpers.ValidateRetainDevices(ctx, req.Config, &resp.Diagnostics)
}

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

	// Associate devices. Devices already on this group are adopted: a previous Create that failed
	// after associating them (e.g. its deploy failed) left them there without a state, and a POST for
	// them would fail on CFGRP0018 against this very group.
	associatePath := fmt.Sprintf("/v1/config-group/%v/device/associate/", plan.ConfigurationGroupId.ValueString())
	newDevices := plan
	newDevices.Devices = nil
	var prior *ConfigurationGroupDevices
	if len(plan.Devices) > 0 {
		res, err := r.client.Get(associatePath)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
			return
		}
		currentLabels := make(map[string]string)
		res.Get("devices").ForEach(func(k, v gjson.Result) bool {
			currentLabels[v.Get("id").String()] = v.Get("groupTopologyLabel").String()
			return true
		})
		relabel := plan
		relabel.Devices = nil
		var adopted []string
		for _, d := range plan.Devices {
			label, exists := currentLabels[d.Id.ValueString()]
			if !exists {
				newDevices.Devices = append(newDevices.Devices, d)
				continue
			}
			adopted = append(adopted, d.Id.ValueString())
			if d.TopologyLabel.ValueString() != label {
				relabel.Devices = append(relabel.Devices, d)
			}
		}
		if len(newDevices.Devices) > 0 {
			res, err = r.associateDevices(ctx, "POST", associatePath, newDevices.toBodyConfigGroupDevices(ctx))
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure configuration group devices (POST), got error: %s, %s", err, res.String()))
				return
			}
		}
		if len(relabel.Devices) > 0 {
			res, err = r.associateDevices(ctx, "PUT", associatePath, relabel.toBodyConfigGroupDevices(ctx))
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure configuration group devices (PUT), got error: %s, %s", err, res.String()))
				return
			}
		}
		if len(adopted) > 0 {
			tflog.Info(ctx, fmt.Sprintf("%s: devices %v are already on this group, adopting them", plan.ConfigurationGroupId.ValueString(), adopted))
			var encrypted map[string][]string
			prior, encrypted, err = r.adoptedPriorState(ctx, plan, adopted)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
				return
			}
			if w := helpers.AdoptedEncryptedWarning(encrypted); w != "" {
				resp.Diagnostics.AddWarning("Encrypted variables not compared for adopted devices", w)
			}
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
			r.disassociateDevices(ctx, newDevices, &resp.Diagnostics) // not the adopted ones
			return
		}
	}

	// Deploy to devices
	if len(plan.Devices) > 0 {
		// Adopted devices deploy like in Update, against what the Manager has; the others always.
		r.Deploy(ctx, plan, prior, nil, nil, &resp.Diagnostics)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.ConfigurationGroupId.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// adoptedPriorState returns the adopted devices with the variables the Manager has for them now,
// read before Create or Update overwrites them. Deploy uses it as the prior state, so an adopted
// device is deployed only when it is out of date or out of sync, or its variables change.
// It also returns, per adopted device, the variables the Manager has encrypted.
func (r *ConfigurationGroupDevicesResource) adoptedPriorState(ctx context.Context, plan ConfigurationGroupDevices, adopted []string) (*ConfigurationGroupDevices, map[string][]string, error) {
	res, err := r.client.Get(fmt.Sprintf("/v1/config-group/%v/device/variables/", plan.ConfigurationGroupId.ValueString()))
	if err != nil {
		return nil, nil, fmt.Errorf("%s, %s", err, res.String())
	}
	prior := plan
	prior.Devices = nil
	for _, d := range plan.Devices {
		if helpers.Contains(adopted, d.Id.ValueString()) {
			d.Variables = nil
			prior.Devices = append(prior.Devices, d)
		}
	}
	// Encrypted values can't be compared and keep the planned value (as in Read).
	prior.fromBodyConfigGroupDeviceVariables(ctx, res, &plan)
	return &prior, helpers.EncryptedVariables(res, adopted), nil
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
	path := fmt.Sprintf("/v1/config-group/%v/device/associate/", data.ConfigurationGroupId.ValueString())
	ids := make([]string, 0, len(data.Devices))
	for _, d := range data.Devices {
		ids = append(ids, d.Id.ValueString())
	}
	// Release only devices still on this group (the Manager's DELETE ignores the group in the URL).
	ids, err := helpers.StillAssociated(ctx, r.client, path, ids)
	if err != nil {
		diag.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}
	if len(ids) == 0 {
		return
	}
	body, _ := sjson.Set("", "devices", []interface{}{})
	for _, id := range ids {
		itemBody, _ := sjson.Set("", "id", id)
		body, _ = sjson.SetRaw(body, "devices.-1", itemBody)
	}
	res, err := helpers.RetryWhileDeployInProgress(ctx, r.taskTimeout, func() (sdwan.Res, error) { return r.client.DeleteBody(path, body) })
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
		// A deploy running on the group refuses associate calls too (CFGRP0017): wait for it.
		res, err := helpers.RetryWhileDeployInProgress(ctx, r.taskTimeout, func() (sdwan.Res, error) {
			if method == "PUT" {
				return r.client.Put(path, body)
			}
			return r.client.Post(path, body)
		})
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

// associated, when set, is a GET .../device/associate/ response the caller has just made; it is used
// instead of reading the group again.
// importEncrypted lists, per device, variables that came back encrypted at import: a difference
// in only those doesn't make a device "updated".
func (r *ConfigurationGroupDevicesResource) Deploy(ctx context.Context, plan ConfigurationGroupDevices, state *ConfigurationGroupDevices, associated *sdwan.Res, importEncrypted map[string][]string, diag *diag.Diagnostics) {
	var updatedDevices []string
	if state != nil {
		updatedDevices = plan.getUpdatedDevices(ctx, state)
		if len(importEncrypted) > 0 {
			kept := updatedDevices[:0]
			var skipped []string
			for _, id := range updatedDevices {
				if names := importEncrypted[id]; len(names) > 0 && onlyVariablesDiffer(plan, state, id, names) {
					tflog.Info(ctx, fmt.Sprintf("%s: device %s only differs in variables encrypted at import %v, not redeploying for that", plan.ConfigurationGroupId.ValueString(), id, names))
					skipped = append(skipped, fmt.Sprintf("%s %v", id, names))
					continue
				}
				kept = append(kept, id)
			}
			updatedDevices = kept
			if len(skipped) > 0 {
				diag.AddWarning("Encrypted variables not compared after import",
					fmt.Sprintf("The Manager returns encrypted variables without their value, so right after an import they can't be compared with the configuration. These devices are not redeployed for them in this apply: %s. If you changed one of these values in this apply, it was written to the Manager but not deployed: run terraform apply -replace on this resource, or change it in a separate apply.", strings.Join(skipped, "; ")))
			}
		}
		// A device not in the prior state is new to this resource: deploy it whatever the Manager reports.
		inState := make(map[string]bool)
		for _, d := range state.Devices {
			inState[d.Id.ValueString()] = true
		}
		for _, d := range plan.Devices {
			if !inState[d.Id.ValueString()] {
				updatedDevices = append(updatedDevices, d.Id.ValueString())
			}
		}
	}

	// A change of the associated group version makes this resource run in the same apply as the
	// group change. From Manager 20.15 the Manager then reports exactly the affected devices out of
	// date, and only those are deployed; below 20.15 the version change deploys all devices
	// (helpers.GroupVersionForcesDeploy). A version that is null on either side is not a change:
	// null in state means never recorded (just imported, or wiring added later), null in the plan
	// means not wired (wiring removed).
	hasVersionChange := false
	if state != nil && !state.ConfigurationGroupVersion.IsNull() && !plan.ConfigurationGroupVersion.IsNull() && !plan.ConfigurationGroupVersion.Equal(state.ConfigurationGroupVersion) {
		hasVersionChange = helpers.GroupVersionForcesDeploy(r.client.ManagerVersion)
		if !hasVersionChange {
			tflog.Debug(ctx, fmt.Sprintf("%s: configuration group version changed, deploying the devices the Manager reports out of date", plan.ConfigurationGroupId.ValueString()))
		}
	}

	var res sdwan.Res
	var err error
	if associated != nil {
		res = *associated
	} else {
		path := fmt.Sprintf("/v1/config-group/%v/device/associate/", plan.ConfigurationGroupId.ValueString())
		res, err = r.client.Get(path)
	}
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
				if item.Id.ValueString() == id && item.Deploy.ValueBool() && (!v.Get("configGroupUpToDate").Bool() || helpers.IsOutOfSync(v) || updatedDevices == nil || helpers.Contains(updatedDevices, id) || hasVersionChange) {
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
		// Waits while another deploy runs on the group (one deploy per group at a time), and sends
		// the deploy again when its task collided with another deploy (helpers.DeployCollision).
		res, postErr, taskErr, warnings := helpers.DeployAndWait(ctx, r.client, path, body, r.taskTimeout)
		if postErr != nil {
			diag.AddError("Client Error", fmt.Sprintf("Failed to deploy to config group devices (POST), got error: %s, %s", postErr, res.String()))
			return
		}
		if taskErr != nil {
			diag.AddError("Client Error", fmt.Sprintf("Failed to deploy to config group devices, got error: %s", taskErr))
			return
		}
		if warnings != "" {
			// A device was skipped because another operation was in progress on it.
			diag.AddWarning("Client Warning", warnings)
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
	var variablesRes sdwan.Res
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
		variablesRes = res
	}

	state.updateTfAttributes(ctx, &oldState)

	if r.deployOnOutOfDate {
		if value := associateRes.Get("devices"); value.Exists() {
			value.ForEach(func(k, v gjson.Result) bool {
				// Out of date or "Out of Sync" (a deploy the router rejected); "Sync Pending" (offline) is
				// skipped, see helpers.NeedsRedeploy.
				if helpers.NeedsRedeploy(v, "configGroupUpToDate") {
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
	importEncrypted, diags := helpers.GetImportEncrypted(ctx, req.Private)
	resp.Diagnostics.Append(diags...)
	if imp {
		state.processImport(ctx)
		// Encrypted values come back as $CRYPT_CLUSTER and, with no earlier state to keep, are
		// stored empty, so the clear-text config value looks like a change. Remember which ones,
		// so the first Update doesn't redeploy for them alone.
		ids := make([]string, 0, len(state.Devices))
		for _, d := range state.Devices {
			ids = append(ids, d.Id.ValueString())
		}
		importEncrypted = helpers.EncryptedVariables(variablesRes, ids)
	}
	resp.Diagnostics.Append(helpers.SetImportEncrypted(ctx, resp.Private, importEncrypted)...)

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

	// Only retain_devices changed (a hand-over set or cleared): nothing to send to the Manager. The
	// variables write and deploy selection below would deploy this state's out-of-date devices at an
	// unexpected time.
	if only, err := helpers.OnlyAttributeChanged(req.Plan.Raw, req.State.Raw, "retain_devices"); err == nil && only {
		tflog.Debug(ctx, fmt.Sprintf("%s: only retain_devices changed, nothing to update", plan.ConfigurationGroupId.ValueString()))
		resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
		return
	}

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

	// Devices already on this group but not in this resource's state were handed over by another
	// state, or associated in the GUI. They are adopted as in Create (see the deploy below).
	inState := make(map[string]bool)
	for _, sd := range state.Devices {
		inState[sd.Id.ValueString()] = true
	}
	var adopted []string
	for _, d := range plan.Devices {
		if _, exists := currentLabels[d.Id.ValueString()]; exists && !inState[d.Id.ValueString()] {
			adopted = append(adopted, d.Id.ValueString())
		}
	}

	// Associate devices that are new in the plan or whose topology_label changed.
	associateBody, _ := sjson.Set("", "devices", []interface{}{})
	var added []string // not on this group yet; checked again after the associate call
	for _, d := range plan.Devices {
		label, exists := currentLabels[d.Id.ValueString()]
		needsAssociate := false
		if !exists {
			needsAssociate = true
			added = append(added, d.Id.ValueString())
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
	var removed []string
	for _, sd := range state.Devices {
		found := false
		for _, d := range plan.Devices {
			if d.Id.ValueString() == sd.Id.ValueString() {
				found = true
				break
			}
		}
		if !found {
			removed = append(removed, sd.Id.ValueString())
		}
	}
	// Devices in retain_devices are handed over to another state on this group: dropped from this
	// state, not released.
	associatePath := fmt.Sprintf("/v1/config-group/%v/device/associate/", plan.ConfigurationGroupId.ValueString())
	removed = helpers.ExceptRetained(ctx, plan.RetainDevices, removed, associatePath)
	// Release only devices still on this group, read right before the DELETE: the Manager's DELETE
	// removes a device from whatever group it is on, so a device another state or the GUI has moved
	// meanwhile would be pulled out of its new group.
	removed, err = helpers.StillAssociated(ctx, r.client, associatePath, removed)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}
	disassociateBody, _ := sjson.Set("", "devices", []interface{}{})
	for _, id := range removed {
		disassociateBody, _ = sjson.SetRaw(disassociateBody, "devices.-1", helpers.Must(sjson.Set("", "id", id)))
	}

	// Disassociate removed devices first, so that a device being moved out of this group is freed
	// before we try to associate a device being moved in. This avoids a deadlock when two resources
	// swap devices in the same apply: if associate ran first, both would block in their retry loop
	// waiting for the other's disassociate, which would never be reached.
	if len(gjson.Get(disassociateBody, "devices").Array()) > 0 {
		res, err = helpers.RetryWhileDeployInProgress(ctx, r.taskTimeout, func() (sdwan.Res, error) {
			return r.client.DeleteBody(fmt.Sprintf("/v1/config-group/%v/device/associate/", plan.ConfigurationGroupId.ValueString()), disassociateBody)
		})
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to disassociate config group devices (DELETE), got error: %s, %s", err, res.String()))
			return
		}
	}

	// Associate missing devices (retries on the "already associated to another group" conflict).
	if len(gjson.Get(associateBody, "devices").Array()) > 0 {
		res, err = r.associateDevices(ctx, "PUT", associatePath, associateBody)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure configuration group devices (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	// A device taken from another group can be removed again by that group's own release in the
	// same apply (see helpers.StillAssociated). Check the added devices are still here, associate
	// them again if not, and reuse this read for the deploy selection.
	var associated *sdwan.Res
	if len(added) > 0 {
		assoc, missing, err := helpers.EnsureAssociated(ctx, r.client, associatePath, added, func(missing []string) (sdwan.Res, error) {
			body, _ := sjson.Set("", "devices", []interface{}{})
			for _, d := range plan.Devices {
				if helpers.Contains(missing, d.Id.ValueString()) {
					itemBody := helpers.Must(sjson.Set("", "id", d.Id.ValueString()))
					if !d.TopologyLabel.IsNull() {
						itemBody = helpers.Must(sjson.Set(itemBody, "groupTopologyLabel", d.TopologyLabel.ValueString()))
					}
					body, _ = sjson.SetRaw(body, "devices.-1", itemBody)
				}
			}
			return r.associateDevices(ctx, "PUT", associatePath, body)
		})
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to verify configuration group devices, got error: %s, %s", err, assoc.String()))
			return
		}
		if len(missing) > 0 {
			resp.Diagnostics.AddWarning("Devices not associated", fmt.Sprintf("Devices %v were removed from the group again after being associated several times. Run terraform apply again.", missing))
		}
		associated = &assoc
	}

	// Adopted devices deploy against what the Manager has, as in Create: their variables are read
	// before the write below and used as their prior state. So a device taken over from another
	// state is deployed only when it is out of date or out of sync, or its variables change.
	deployState := state
	if len(adopted) > 0 {
		tflog.Info(ctx, fmt.Sprintf("%s: devices %v are already on this group, adopting them", plan.ConfigurationGroupId.ValueString(), adopted))
		prior, encrypted, err := r.adoptedPriorState(ctx, plan, adopted)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
			return
		}
		if w := helpers.AdoptedEncryptedWarning(encrypted); w != "" {
			resp.Diagnostics.AddWarning("Encrypted variables not compared for adopted devices", w)
		}
		deployState.Devices = append(append([]ConfigurationGroupDevicesDevices{}, state.Devices...), prior.Devices...)
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
	importEncrypted, diags := helpers.GetImportEncrypted(ctx, req.Private)
	resp.Diagnostics.Append(diags...)
	if len(plan.Devices) > 0 {
		r.Deploy(ctx, plan, &deployState, associated, importEncrypted, &resp.Diagnostics)
	}
	// From here the state holds the config values, so the import exception is used up.
	resp.Diagnostics.Append(helpers.SetImportEncrypted(ctx, resp.Private, nil)...)
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

// onlyVariablesDiffer reports whether device id differs between plan and state in the variables
// named in names only (everything else, including the set of other variables, is equal).
func onlyVariablesDiffer(plan ConfigurationGroupDevices, state *ConfigurationGroupDevices, id string, names []string) bool {
	collect := func(devices []ConfigurationGroupDevicesDevices) (map[string]string, bool) {
		for _, d := range devices {
			if d.Id.ValueString() == id {
				vars := map[string]string{}
				for _, v := range d.Variables {
					if !helpers.Contains(names, v.Name.ValueString()) {
						vars[v.Name.ValueString()] = v.Value.String() + "|" + v.ListValue.String()
					}
				}
				return vars, true
			}
		}
		return nil, false
	}
	p, okp := collect(plan.Devices)
	st, oks := collect(state.Devices)
	if !okp || !oks || len(p) != len(st) {
		return false
	}
	for k, v := range p {
		if st[k] != v {
			return false
		}
	}
	return true
}
