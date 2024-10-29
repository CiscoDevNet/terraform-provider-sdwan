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
	"regexp"
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
var _ resource.Resource = &ServiceWirelessLANProfileParcelResource{}
var _ resource.ResourceWithImportState = &ServiceWirelessLANProfileParcelResource{}

func NewServiceWirelessLANProfileParcelResource() resource.Resource {
	return &ServiceWirelessLANProfileParcelResource{}
}

type ServiceWirelessLANProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *ServiceWirelessLANProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_wireless_lan_feature"
}

func (r *ServiceWirelessLANProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Service Wireless LAN Feature.").AddMinimumVersionDescription("20.12.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Feature",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Feature",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Feature",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Feature",
				Optional:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Feature Profile ID").String,
				Required:            true,
			},
			"enable_24g": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("2.4GHz Enabled").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"enable_24g_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"enable_5g": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("5GHz Enabled").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"enable_5g_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ssids": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Wi-Fi SSID profile").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ssid_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure wlan SSID").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 64),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-zA-Z!->@^_]*`), ""),
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
						"broadcast_ssid_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
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
							MarkdownDescription: helpers.NewAttributeDescription("Select security type").AddStringEnumDescription("enterprise", "personal", "open").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("enterprise", "personal", "open"),
							},
						},
						"radius_server_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set RADIUS server IP, Attribute conditional on `security_type` being equal to `enterprise`").String,
							Optional:            true,
						},
						"radius_server_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"radius_server_port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set RADIUS server authentication port, Attribute conditional on `security_type` being equal to `enterprise`").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("1812").String,
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
							MarkdownDescription: helpers.NewAttributeDescription("Set RADIUS server shared secret, Attribute conditional on `security_type` being equal to `enterprise`").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(8, 32),
							},
						},
						"radius_server_secret_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"passphrase": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set passphrase, Attribute conditional on `security_type` being equal to `personal`").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(8, 63),
							},
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
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-zA-Z!->@^_]*`), ""),
				},
			},
			"username_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set management password,the password must contains characters from all of the following classes,lowercase letters,uppercase letters,digits,and special characters. No character in the password can be repeated more than three times consecutively. The password must not be the same as the associated username or the username reversed. The password must not be cisco,ocsic,or any variant obtained by changing the capitalization of the letters in word cisco. In addition,you can't substitute 1,l,or ! for i,0 for o,$ for s.").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9~!@#%^&()_+|<>,.?/:;'\[\]{}"]{8,64}$`), ""),
				},
			},
			"password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"me_dynamic_ip_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ME management IP dynamic allocated by DHCP").String,
				Required:            true,
			},
			"me_ipv4_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set mobile express controller address").String,
				Optional:            true,
			},
			"me_ipv4_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"me_subnet_mask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set mobile express controller subnet mask").AddStringEnumDescription("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0"),
				},
			},
			"me_subnet_mask_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"me_default_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set mobile express default gateway").String,
				Optional:            true,
			},
			"me_default_gateway_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
		},
	}
}

func (r *ServiceWirelessLANProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *ServiceWirelessLANProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ServiceWirelessLAN

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("parcelId").String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *ServiceWirelessLANProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ServiceWirelessLAN

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if res.Get("error.message").String() == "Invalid feature Id" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}
	if state.Version == types.Int64Null() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *ServiceWirelessLANProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ServiceWirelessLAN

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
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *ServiceWirelessLANProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ServiceWirelessLAN

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && res.Get("error.message").String() != "Invalid Template Id" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *ServiceWirelessLANProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "service_wireless_lan_feature_id" + ",feature_profile_id"
	if len(parts) != (count + 1) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: %s. Got: %q, %q", pattern, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("feature_profile_id"), parts[1])...)
}

// End of section. //template:end import
