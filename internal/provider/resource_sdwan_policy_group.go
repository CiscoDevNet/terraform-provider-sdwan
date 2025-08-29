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
	"github.com/hashicorp/go-version"
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

var MinPolicyGroupUpdateVersion = version.Must(version.NewVersion("20.15.0"))

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &PolicyGroupResource{}
var _ resource.ResourceWithImportState = &PolicyGroupResource{}

func NewPolicyGroupResource() resource.Resource {
	return &PolicyGroupResource{}
}

type PolicyGroupResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *PolicyGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_group"
}

func (r *PolicyGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Policy Group .").AddMinimumVersionDescription("20.12.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the policy group").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description").String,
				Required:            true,
			},
			"solution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of solution").AddStringEnumDescription("sdwan").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("sdwan"),
				},
			},
			"feature_profile_ids": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of feature profile IDs").String,
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
			"policy_versions": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of all associated policy versions").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func (r *PolicyGroupResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

func (r *PolicyGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PolicyGroup

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create policy group
	body := plan.toBodyPolicyGroup(ctx)
	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())

	// Create policy group devices
	if len(plan.Devices) > 0 {
		body = plan.toBodyPolicyGroupDevices(ctx)

		path := fmt.Sprintf("/v1/policy-group/%v/device/associate/", plan.Id.ValueString())
		res, err = r.client.Post(path, body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure policy group devices (POST), got error: %s, %s", err, res.String()))
			r.DeletePolicyGroup(ctx, plan, &resp.Diagnostics)
			return
		}
	}

	// Create policy group device variables
	if len(plan.Devices) > 0 && plan.hasPolicyGroupDeviceVariables(ctx) {
		body = plan.toBodyPolicyGroupDeviceVariables(ctx)

		path := fmt.Sprintf("/v1/policy-group/%v/device/variables/", plan.Id.ValueString())
		res, err = r.client.Put(path, body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure policy group device variables (PUT), got error: %s, %s", err, res.String()))
			r.DeletePolicyGroup(ctx, plan, &resp.Diagnostics)
			return
		}
	}

	// Deploy policy group to devices
	if len(plan.Devices) > 0 {
		r.Deploy(ctx, plan, nil, &resp.Diagnostics, true)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *PolicyGroupResource) Deploy(ctx context.Context, plan PolicyGroup, state *PolicyGroup, diag *diag.Diagnostics, deleteOnError bool) {
	var updatedDevices []string
	if state != nil {
		updatedDevices = plan.getUpdatedDevices(ctx, state)
	}

	hasPolicyVersionChanges := false
	currentVersion := version.Must(version.NewVersion(r.client.ManagerVersion))
	if state != nil && currentVersion.LessThan(MinPolicyGroupUpdateVersion) {
		hasPolicyVersionChanges = plan.hasPolicyVersionChanges(ctx, state)
	}

	path := fmt.Sprintf("/v1/policy-group/%v/device/associate/", plan.Id.ValueString())
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
				if item.Id.ValueString() == id && item.Deploy.ValueBool() && (!v.Get("policyGroupUpToDate").Bool() || updatedDevices == nil || helpers.Contains(updatedDevices, id) || hasPolicyVersionChanges) {
					itemBody, _ := sjson.Set("", "id", id)
					body, _ = sjson.SetRaw(body, "devices.-1", itemBody)
					tflog.Debug(ctx, fmt.Sprintf("%s: Deploying to device %s", plan.Name.ValueString(), id))
				}
			}
			return true
		})
	}
	if len(gjson.Get(body, "devices").Array()) > 0 {
		path := fmt.Sprintf("/v1/policy-group/%v/device/deploy/", plan.Id.ValueString())
		res, err = r.client.Post(path, body)
		if err != nil {
			diag.AddError("Client Error", fmt.Sprintf("Failed to deploy to policy group devices (POST), got error: %s, %s", err, res.String()))
			if deleteOnError {
				r.DeletePolicyGroup(ctx, plan, diag)
			}
			return
		}

		// Wait for deploy action to complete
		actionId := res.Get("parentTaskId").String()
		err = helpers.WaitForActionToComplete(ctx, r.client, actionId)
		if err != nil {
			diag.AddError("Client Error", fmt.Sprintf("Failed to deploy to config group devices, got error: %s", err))
			if deleteOnError {
				r.DeletePolicyGroup(ctx, plan, diag)
			}
			return
		}
	}
}

func (r *PolicyGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PolicyGroup

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.ValueString()))

	oldState := state

	// Read policy group
	res, err := r.client.Get(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if res.Raw == "" && err == nil {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBodyPolicyGroup(ctx, res)

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	if imp {
		state.processImport(ctx)
	}

	// Read policy group device associations
	path := fmt.Sprintf("/v1/policy-group/%v/device/associate/", state.Id.ValueString())
	res, err = r.client.Get(path)
	if strings.Contains(res.Get("error.message").String(), "Invalid policy group passed") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBodyPolicyGroupDevices(ctx, res)

	// Read policy group device variables
	if value := res.Get("devices"); value.Exists() && len(value.Array()) > 0 {
		path = fmt.Sprintf("/v1/policy-group/%v/device/variables/", state.Id.ValueString())
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

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *PolicyGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PolicyGroup

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

	// Update policy group
	body := plan.toBodyPolicyGroup(ctx)

	res, err := r.client.Put(plan.getPath()+url.QueryEscape(plan.Id.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	res, err = r.client.Get(fmt.Sprintf("/v1/policy-group/%v/device/associate/", plan.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	var currentDeviceIds []string
	if value := res.Get("devices"); value.Exists() && len(value.Array()) > 0 {
		value.ForEach(func(k, v gjson.Result) bool {
			currentDeviceIds = append(currentDeviceIds, v.Get("id").String())
			return true
		})
	}

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

	disassociateBody, _ := sjson.Set("", "devices", []interface{}{})
	for _, cdid := range currentDeviceIds {
		found := false
		for _, d := range plan.Devices {
			if d.Id.ValueString() == cdid {
				found = true
				break
			}
		}
		if !found {
			disassociateBody, _ = sjson.SetRaw(disassociateBody, "devices.-1", helpers.Must(sjson.Set("", "id", cdid)))
		}
	}

	// associate missing devices
	if len(gjson.Get(associateBody, "devices").Array()) > 0 {
		res, err = r.client.Put(fmt.Sprintf("/v1/policy-group/%v/device/associate/", plan.Id.ValueString()), associateBody)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure policy group devices (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	// disassociate extra devices
	if len(gjson.Get(disassociateBody, "devices").Array()) > 0 {
		res, err = r.client.DeleteBody(fmt.Sprintf("/v1/policy-group/%v/device/associate/", plan.Id.ValueString()), disassociateBody)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete policy group devices (DELETE), got error: %s, %s", err, res.String()))
			return
		}
	}

	// Update policy group device variables
	if len(plan.Devices) > 0 && plan.hasPolicyGroupDeviceVariables(ctx) {
		body = plan.toBodyPolicyGroupDeviceVariables(ctx)

		path := fmt.Sprintf("/v1/policy-group/%v/device/variables/", plan.Id.ValueString())
		res, err = r.client.Put(path, body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure policy group device variables (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	// Deploy policy group to devices
	if len(plan.Devices) > 0 {
		r.Deploy(ctx, plan, &state, &resp.Diagnostics, false)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PolicyGroupResource) DeletePolicyGroup(ctx context.Context, state PolicyGroup, diag *diag.Diagnostics) {
	path := fmt.Sprintf("/v1/policy-group/%v/device/associate/", state.Id.ValueString())
	res, err := r.client.Get(path)
	if err == nil {
		body, _ := sjson.Set("", "devices", []interface{}{})
		if value := res.Get("devices"); value.Exists() && len(value.Array()) > 0 {
			value.ForEach(func(k, v gjson.Result) bool {
				id := v.Get("id").String()
				itemBody, _ := sjson.Set("", "id", id)
				body, _ = sjson.SetRaw(body, "devices.-1", itemBody)
				return true
			})
		}
		if len(gjson.Get(body, "devices").Array()) > 0 {
			res, err := r.client.DeleteBody(path, body)
			if err != nil {
				diag.AddError("Client Error", fmt.Sprintf("Failed to delete policy group devices (DELETE), got error: %s, %s", err, res.String()))
				return
			}
		}
	}

	res, err = r.client.Delete(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if err != nil {
		diag.AddError("Client Error", fmt.Sprintf("Failed to delete policy group (DELETE), got error: %s, %s", err, res.String()))
		return
	}
}

func (r *PolicyGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PolicyGroup

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	r.DeletePolicyGroup(ctx, state, &resp.Diagnostics)

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *PolicyGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
