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
var _ resource.Resource = &SystemNTPProfileParcelResource{}
var _ resource.ResourceWithImportState = &SystemNTPProfileParcelResource{}

func NewSystemNTPProfileParcelResource() resource.Resource {
	return &SystemNTPProfileParcelResource{}
}

type SystemNTPProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *SystemNTPProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_ntp_feature"
}

func (r *SystemNTPProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a System NTP Feature.").AddMinimumVersionDescription("20.12.0").String,

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
				Optional:            true,
			},
			"servers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure NTP servers").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"hostname_ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set hostname or IP address of server").String,
							Optional:            true,
						},
						"hostname_ip_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"authentication_key": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set authentication key for the server").AddIntegerRangeDescription(1, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
						"authentication_key_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vpn": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set VPN in which NTP server is located").AddIntegerRangeDescription(0, 65530).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(65530),
							},
						},
						"vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ntp_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set NTP version").AddIntegerRangeDescription(1, 4).AddDefaultValueDescription("4").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4),
							},
						},
						"ntp_version_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set interface to use to reach NTP server").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(3, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`(ATM|ATM-ACR|AppGigabitEthernet|AppNav-Compress|AppNav-UnCompress|Async|BD-VIF|BDI|CEM|CEM-ACR|Cellular|Dialer|Embedded-Service-Engine|Ethernet|Ethernet-Internal|FastEthernet|FiftyGigabitEthernet|FiveGigabitEthernet|FortyGigabitEthernet|FourHundredGigE|GMPLS|GigabitEthernet|Group-Async|HundredGigE|L2LISP|LISP|Loopback|MFR|Multilink|Port-channel|SM|Serial|Service-Engine|TenGigabitEthernet|Tunnel|TwentyFiveGigE|TwentyFiveGigabitEthernet|TwoGigabitEthernet|TwoHundredGigE|Vif|Virtual-PPP|Virtual-Template|VirtualPortGroup|Vlan|Wlan-GigabitEthernet|nat64|nat66|ntp|nve|ospfv3|overlay|pseudowire|ucse|vasileft|vasiright|vmi)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
							},
						},
						"source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"prefer_this_ntp_server": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prefer this NTP server").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"prefer_this_ntp_server_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"authentication_keys": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set MD5 authentication key").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("MD5 authentication key ID").AddIntegerRangeDescription(1, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
						"key_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"md5_value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enter cleartext or AES-encrypted MD5 authentication key").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
						},
						"md5_value_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"trusted_keys": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Designate authentication key as trustworthy").String,
				ElementType:         types.Int64Type,
				Optional:            true,
			},
			"trusted_keys_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"authoritative_ntp_server": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable device as NTP Leader").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"authoritative_ntp_server_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"stratum": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable device as NTP Leader").AddIntegerRangeDescription(1, 15).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 15),
				},
			},
			"stratum_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"source_interface": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable device as NTP Leader").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(3, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`(ATM|ATM-ACR|AppGigabitEthernet|AppNav-Compress|AppNav-UnCompress|Async|BD-VIF|BDI|CEM|CEM-ACR|Cellular|Dialer|Embedded-Service-Engine|Ethernet|Ethernet-Internal|FastEthernet|FiftyGigabitEthernet|FiveGigabitEthernet|FortyGigabitEthernet|FourHundredGigE|GMPLS|GigabitEthernet|Group-Async|HundredGigE|L2LISP|LISP|Loopback|MFR|Multilink|Port-channel|SM|Serial|Service-Engine|TenGigabitEthernet|Tunnel|TwentyFiveGigE|TwentyFiveGigabitEthernet|TwoGigabitEthernet|TwoHundredGigE|Vif|Virtual-PPP|Virtual-Template|VirtualPortGroup|Vlan|Wlan-GigabitEthernet|nat64|nat66|ntp|nve|ospfv3|overlay|pseudowire|ucse|vasileft|vasiright|vmi)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
				},
			},
			"source_interface_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
		},
	}
}

func (r *SystemNTPProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *SystemNTPProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemNTP

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
func (r *SystemNTPProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemNTP

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *SystemNTPProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemNTP

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
func (r *SystemNTPProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemNTP

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
func (r *SystemNTPProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
