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
var _ resource.Resource = &CiscoWirelessLANFeatureTemplateResource{}
var _ resource.ResourceWithImportState = &CiscoWirelessLANFeatureTemplateResource{}

func NewCiscoWirelessLANFeatureTemplateResource() resource.Resource {
	return &CiscoWirelessLANFeatureTemplateResource{}
}

type CiscoWirelessLANFeatureTemplateResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *CiscoWirelessLANFeatureTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_wireless_lan_feature_template"
}

func (r *CiscoWirelessLANFeatureTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Cisco Wireless LAN feature template.").AddMinimumVersionDescription("15.0.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the feature template",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the feature template",
				Computed:            true,
			},
			"template_type": schema.StringAttribute{
				MarkdownDescription: "The template type",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the feature template",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the feature template",
				Required:            true,
			},
			"device_types": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of supported device types").AddStringEnumDescription("vedge-C8000V", "vedge-C8300-1N1S-4T2X", "vedge-C8300-1N1S-6T", "vedge-C8300-2N2S-6T", "vedge-C8300-2N2S-4T2X", "vedge-C8500-12X4QC", "vedge-C8500-12X", "vedge-C8500-20X6C", "vedge-C8500L-8S4X", "vedge-C8200-1N-4T", "vedge-C8200L-1N-4T").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"shutdown_2_4ghz": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("2.4GHz Shutdown").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"shutdown_2_4ghz_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"shutdown_5ghz": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("5GHz Shutdown").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"shutdown_5ghz_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ssids": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Wi-Fi SSID").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"wireless_network_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure wlan SSID").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 64),
							},
						},
						"admin_state": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set admin state").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"admin_state_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"broadcast_ssid": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable broadcast SSID").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"vlan_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set VLAN ID").AddIntegerRangeDescription(1, 4094).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4094),
							},
						},
						"vlan_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"radio_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select radio type").AddStringEnumDescription("24ghz", "5ghz", "all").AddDefaultValueDescription("all").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("24ghz", "5ghz", "all"),
							},
						},
						"radio_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"security_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select security type").AddStringEnumDescription("enterprise", "personal", "open").AddDefaultValueDescription("personal").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("enterprise", "personal", "open"),
							},
						},
						"security_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"radius_server_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set RADIUS server IP").String,
							Optional:            true,
						},
						"radius_server_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"radius_server_port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set RADIUS server authentication port").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("1812").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
						"radius_server_port_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"radius_server_secret": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set RADIUS server shared secret").String,
							Optional:            true,
						},
						"radius_server_secret_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"passphrase": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set passphrase").String,
							Optional:            true,
						},
						"passphrase_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"qos_profile": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select QoS profile").AddStringEnumDescription("platinum", "gold", "silver", "bronze").AddDefaultValueDescription("silver").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("platinum", "gold", "silver", "bronze"),
							},
						},
						"qos_profile_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"country": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Select country").AddStringEnumDescription("AE", "AR", "AT", "AU", "BA", "BB", "BE", "BG", "BH", "BN", "BO", "BR", "BY", "CA", "CA2", "CH", "CL", "CM", "CN", "CO", "CR", "CY", "CZ", "DE", "DK", "DO", "DZ", "EC", "EE", "EG", "ES", "FI", "FJ", "FR", "GB", "GH", "GI", "GR", "HK", "HR", "HU", "ID", "IE", "IL", "IO", "IN", "IQ", "IS", "IT", "J2", "J4", "JM", "JO", "KE", "KN", "KW", "KZ", "LB", "LI", "LK", "LT", "LU", "LV", "LY", "MA", "MC", "ME", "MK", "MN", "MO", "MT", "MX", "MY", "NL", "NO", "NZ", "OM", "PA", "PE", "PH", "PH2", "PK", "PL", "PR", "PT", "PY", "QA", "RO", "RS", "RU", "SA", "SE", "SG", "SI", "SK", "TH", "TN", "TR", "TW", "UA", "US", "UY", "VE", "VN", "ZA").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("AE", "AR", "AT", "AU", "BA", "BB", "BE", "BG", "BH", "BN", "BO", "BR", "BY", "CA", "CA2", "CH", "CL", "CM", "CN", "CO", "CR", "CY", "CZ", "DE", "DK", "DO", "DZ", "EC", "EE", "EG", "ES", "FI", "FJ", "FR", "GB", "GH", "GI", "GR", "HK", "HR", "HU", "ID", "IE", "IL", "IO", "IN", "IQ", "IS", "IT", "J2", "J4", "JM", "JO", "KE", "KN", "KW", "KZ", "LB", "LI", "LK", "LT", "LU", "LV", "LY", "MA", "MC", "ME", "MK", "MN", "MO", "MT", "MX", "MY", "NL", "NO", "NZ", "OM", "PA", "PE", "PH", "PH2", "PK", "PL", "PR", "PT", "PY", "QA", "RO", "RS", "RU", "SA", "SE", "SG", "SI", "SK", "TH", "TN", "TR", "TW", "UA", "US", "UY", "VE", "VN", "ZA"),
				},
			},
			"country_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set management username").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"username_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set management password").String,
				Optional:            true,
			},
			"password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"controller_ip_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set mobile express controller address").AddDefaultValueDescription("0.0.0.0").String,
				Optional:            true,
			},
			"controller_ip_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"controller_subnet_mask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set mobile express controller subnet mask").AddDefaultValueDescription("0.0.0.0").String,
				Optional:            true,
			},
			"controller_subnet_mask_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"controller_default_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set mobile express default gateway").AddDefaultValueDescription("0.0.0.0").String,
				Optional:            true,
			},
			"controller_default_gateway_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
		},
	}
}

func (r *CiscoWirelessLANFeatureTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *CiscoWirelessLANFeatureTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CiscoWirelessLAN

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post("/template/feature", body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("templateId").String())
	plan.Version = types.Int64Value(0)
	plan.TemplateType = types.StringValue(plan.getModel())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *CiscoWirelessLANFeatureTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CiscoWirelessLAN

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get("/template/feature/object/" + url.QueryEscape(state.Id.ValueString()))
	if res.Get("error.message").String() == "Invalid Template Id" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)
	if state.Version == types.Int64Null() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *CiscoWirelessLANFeatureTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CiscoWirelessLAN

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

	body := plan.toBody(ctx)
	r.updateMutex.Lock()
	res, err := r.client.Put("/template/feature/"+url.QueryEscape(plan.Id.ValueString()), body)
	r.updateMutex.Unlock()
	if err != nil {
		if res.Get("error.message").String() == "Template locked in edit mode." {
			resp.Diagnostics.AddWarning("Client Warning", fmt.Sprintf("Failed to modify template due to template being locked by another change. Template changes will not be applied. Re-run 'terraform apply' to try again."))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	if plan.hasChanges(ctx, &state) {
		plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)
	} else {
		plan.Version = state.Version
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *CiscoWirelessLANFeatureTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CiscoWirelessLAN

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete("/template/feature/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && res.Get("error.message").String() != "Invalid Template Id" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *CiscoWirelessLANFeatureTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
