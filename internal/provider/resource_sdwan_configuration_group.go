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
var _ resource.Resource = &ConfigurationGroupResource{}
var _ resource.ResourceWithImportState = &ConfigurationGroupResource{}

func NewConfigurationGroupResource() resource.Resource {
	return &ConfigurationGroupResource{}
}

type ConfigurationGroupResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
	taskTimeout *int64
}

func (r *ConfigurationGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_configuration_group"
}

func (r *ConfigurationGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Configuration Group .").AddMinimumVersionDescription("20.15.0").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("The name of the configuration group").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description").String,
				Required:            true,
			},
			"solution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of solution").AddStringEnumDescription("mobility", "sdwan", "nfvirtual").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("mobility", "sdwan", "nfvirtual"),
				},
			},
			"feature_profile_ids": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of feature profile IDs").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"topology_devices": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of topology device types").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"criteria_attribute": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Criteria attribute").AddStringEnumDescription("tag").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("tag"),
							},
						},
						"criteria_value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Criteria value").String,
							Optional:            true,
						},
						"unsupported_features": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of unsupported features").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"parcel_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Parcel type").AddStringEnumDescription("wan/vpn/interface/gre", "wan/vpn/interface/ethernet", "wan/vpn/interface/cellular", "wan/vpn/interface/ipsec", "wan/vpn/interface/serial", "route-policy", "routing/bgp", "routing/ospf", "lan/vpn/interface/ethernet", "lan/vpn/interface/svi", "lan/vpn/interface/ipsec", "lan/vpn").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("wan/vpn/interface/gre", "wan/vpn/interface/ethernet", "wan/vpn/interface/cellular", "wan/vpn/interface/ipsec", "wan/vpn/interface/serial", "route-policy", "routing/bgp", "routing/ospf", "lan/vpn/interface/ethernet", "lan/vpn/interface/svi", "lan/vpn/interface/ipsec", "lan/vpn"),
										},
									},
									"parcel_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Parcel ID").String,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
			"topology_site_devices": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of devices per site").AddIntegerRangeDescription(1, 20).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 20),
				},
			},
			"feature_versions": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of all associated feature versions").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func (r *ConfigurationGroupResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
	r.taskTimeout = req.ProviderData.(*SdwanProviderData).TaskTimeout
}

// End of section. //template:end model

func (r *ConfigurationGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ConfigurationGroup

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create config group
	body := plan.toBodyConfigGroup(ctx)

	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure configuration group (POST), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *ConfigurationGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ConfigurationGroup

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	// Read config group
	res, err := r.client.Get(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if strings.Contains(res.Get("error.message").String(), "Invalid config group passed") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBodyConfigGroup(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *ConfigurationGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ConfigurationGroup

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

	// Update config group only if the group body actually changed.
	planBody := plan.toBodyConfigGroup(ctx)
	stateBody := state.toBodyConfigGroup(ctx)
	if planBody != stateBody {
		res, err := r.client.Put(plan.getPath()+url.QueryEscape(plan.Id.ValueString()), planBody)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure configuration group (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	// Bump the version whenever the group body or any associated feature version changed, so that
	// dependent sdwan_configuration_group_devices resources can detect the change and redeploy.
	if planBody != stateBody || plan.hasFeatureVersionChanges(ctx, &state) {
		plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)
	} else {
		plan.Version = state.Version
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ConfigurationGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ConfigurationGroup

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	// Delete the config group. This intentionally does not disassociate devices: those are managed
	// by sdwan_configuration_group_devices resources (possibly in other state files). If devices are
	// still associated, the API returns an error, which correctly forces the devices to be destroyed first.
	res, err := r.client.Delete(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete config group (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *ConfigurationGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
