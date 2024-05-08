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
var _ resource.Resource = &CiscoVPNInterfaceIPSecFeatureTemplateResource{}
var _ resource.ResourceWithImportState = &CiscoVPNInterfaceIPSecFeatureTemplateResource{}

func NewCiscoVPNInterfaceIPSecFeatureTemplateResource() resource.Resource {
	return &CiscoVPNInterfaceIPSecFeatureTemplateResource{}
}

type CiscoVPNInterfaceIPSecFeatureTemplateResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *CiscoVPNInterfaceIPSecFeatureTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_vpn_interface_ipsec_feature_template"
}

func (r *CiscoVPNInterfaceIPSecFeatureTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Cisco VPN Interface IPSec feature template.").AddMinimumVersionDescription("15.0.0").String,

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
			"interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface name: IPsec when present").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(6, 8),
				},
			},
			"interface_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative state").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"shutdown_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"interface_description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface description").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 240),
				},
			},
			"interface_description_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ip_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Assign IPv4 address").String,
				Optional:            true,
			},
			"ip_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_source": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Tunnel source IP Address").String,
				Optional:            true,
			},
			"tunnel_source_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_source_interface": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("<1..32 characters> Interface name: ge0/<0-..> or ge0/<0-..>.vlanid").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"tunnel_source_interface_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_destination": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Tunnel destination IP address").String,
				Optional:            true,
			},
			"tunnel_destination_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"application": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Application Tunnel Type").AddStringEnumDescription("none", "sig").AddDefaultValueDescription("none").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("none", "sig"),
				},
			},
			"application_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tcp_mss_adjust": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("TCP MSS on SYN packets, in bytes").AddIntegerRangeDescription(500, 1460).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(500, 1460),
				},
			},
			"tcp_mss_adjust_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"clear_dont_fragment": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable clear dont fragment (Currently Only SDWAN Tunnel Interface)").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"clear_dont_fragment_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface MTU <68..9216>, in bytes").AddIntegerRangeDescription(68, 9216).AddDefaultValueDescription("1500").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(68, 9216),
				},
			},
			"mtu_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"dead_peer_detection_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IKE keepalive interval (seconds)").AddIntegerRangeDescription(10, 3600).AddDefaultValueDescription("10").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(10, 3600),
				},
			},
			"dead_peer_detection_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"dead_peer_detection_retries": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IKE keepalive retries").AddIntegerRangeDescription(2, 60).AddDefaultValueDescription("3").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 60),
				},
			},
			"dead_peer_detection_retries_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ike_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IKE Version <1..2>").AddIntegerRangeDescription(1, 2).AddDefaultValueDescription("1").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 2),
				},
			},
			"ike_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IKE integrity protocol").AddStringEnumDescription("main", "aggressive").AddDefaultValueDescription("main").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("main", "aggressive"),
				},
			},
			"ike_mode_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ike_rekey_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IKE rekey interval <60..86400> seconds").AddIntegerRangeDescription(60, 86400).AddDefaultValueDescription("14400").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(60, 86400),
				},
			},
			"ike_rekey_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ike_ciphersuite": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IKE identity the IKE preshared secret belongs to").AddStringEnumDescription("aes256-cbc-sha1", "aes256-cbc-sha2", "aes128-cbc-sha1", "aes128-cbc-sha2").AddDefaultValueDescription("aes256-cbc-sha1").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("aes256-cbc-sha1", "aes256-cbc-sha2", "aes128-cbc-sha1", "aes128-cbc-sha2"),
				},
			},
			"ike_ciphersuite_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ike_group": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IKE Diffie Hellman Groups").AddStringEnumDescription("2", "14", "15", "16", "19", "20", "21", "24").AddDefaultValueDescription("16").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("2", "14", "15", "16", "19", "20", "21", "24"),
				},
			},
			"ike_group_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ike_pre_shared_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use preshared key to authenticate IKE peer").String,
				Optional:            true,
			},
			"ike_pre_shared_key_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ike_pre_shared_key_local_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IKE ID for the local endpoint. Input IPv4 address, domain name, or email address").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 63),
				},
			},
			"ike_pre_shared_key_local_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ike_pre_shared_key_remote_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IKE ID for the remote endpoint. Input IPv4 address, domain name, or email address").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 63),
				},
			},
			"ike_pre_shared_key_remote_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipsec_rekey_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPsec rekey interval <300..1209600> seconds").AddIntegerRangeDescription(120, 2592000).AddDefaultValueDescription("3600").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(120, 2592000),
				},
			},
			"ipsec_rekey_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipsec_replay_window": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Replay window size 32..8192 (must be a power of 2)").AddIntegerRangeDescription(64, 4096).AddDefaultValueDescription("512").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(64, 4096),
				},
			},
			"ipsec_replay_window_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipsec_ciphersuite": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPsec(ESP) encryption and integrity protocol").AddStringEnumDescription("aes256-cbc-sha1", "aes256-cbc-sha384", "aes256-cbc-sha256", "aes256-cbc-sha512", "aes256-gcm", "null-sha1", "null-sha384", "null-sha256", "null-sha512").AddDefaultValueDescription("aes256-gcm").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("aes256-cbc-sha1", "aes256-cbc-sha384", "aes256-cbc-sha256", "aes256-cbc-sha512", "aes256-gcm", "null-sha1", "null-sha384", "null-sha256", "null-sha512"),
				},
			},
			"ipsec_ciphersuite_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipsec_perfect_forward_secrecy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPsec perfect forward secrecy settings").AddStringEnumDescription("group-1", "group-2", "group-5", "group-14", "group-15", "group-16", "group-19", "group-20", "group-21", "group-24", "none").AddDefaultValueDescription("group-16").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("group-1", "group-2", "group-5", "group-14", "group-15", "group-16", "group-19", "group-20", "group-21", "group-24", "none"),
				},
			},
			"ipsec_perfect_forward_secrecy_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tracker": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable tracker for this interface").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"tracker_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_route_via": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("<1..32 characters> Interface name: ge0/<0-..> or ge0/<0-..>.vlanid").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"tunnel_route_via_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
		},
	}
}

func (r *CiscoVPNInterfaceIPSecFeatureTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *CiscoVPNInterfaceIPSecFeatureTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CiscoVPNInterfaceIPSec

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
func (r *CiscoVPNInterfaceIPSecFeatureTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CiscoVPNInterfaceIPSec

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *CiscoVPNInterfaceIPSecFeatureTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CiscoVPNInterfaceIPSec

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
func (r *CiscoVPNInterfaceIPSecFeatureTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CiscoVPNInterfaceIPSec

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
func (r *CiscoVPNInterfaceIPSecFeatureTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
