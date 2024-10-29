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
var _ resource.Resource = &CEdgeGlobalFeatureTemplateResource{}
var _ resource.ResourceWithImportState = &CEdgeGlobalFeatureTemplateResource{}

func NewCEdgeGlobalFeatureTemplateResource() resource.Resource {
	return &CEdgeGlobalFeatureTemplateResource{}
}

type CEdgeGlobalFeatureTemplateResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *CEdgeGlobalFeatureTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cedge_global_feature_template"
}

func (r *CEdgeGlobalFeatureTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a cEdge Global feature template.").AddMinimumVersionDescription("15.0.0").String,

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
			"nat64_udp_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set NAT64 UDP session timeout, in seconds").AddIntegerRangeDescription(1, 536870).AddDefaultValueDescription("300").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 536870),
				},
			},
			"nat64_udp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"nat64_tcp_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set NAT64 TCP session timeout, in seconds").AddIntegerRangeDescription(1, 536870).AddDefaultValueDescription("3600").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 536870),
				},
			},
			"nat64_tcp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"http_authentication": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set preference for HTTP Authentication").AddStringEnumDescription("local", "aaa").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("local", "aaa"),
				},
			},
			"http_authentication_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ssh_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set SSH version").AddIntegerRangeDescription(1, 2).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 2),
				},
			},
			"ssh_version_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"http_server": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set HTTP Server").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"http_server_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"https_server": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set HTTPS Server").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"https_server_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"source_interface": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify interface for source address in all HTTP(S) client connections").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 256),
				},
			},
			"source_interface_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ip_source_routing": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set Source Route").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ip_source_routing_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"arp_proxy": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set ARP Proxy").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"arp_proxy_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ftp_passive": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set Passive FTP").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ftp_passive_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"rsh_rcp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set RSH/RCP").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"rsh_rcp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"bootp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Ignore BOOTP").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"bootp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"domain_lookup": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Domain-Lookup").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"domain_lookup_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tcp_keepalives_out": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure tcp-keepalives-out").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"tcp_keepalives_out_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tcp_keepalives_in": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure tcp-keepalives-in").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"tcp_keepalives_in_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tcp_small_servers": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure tcp-small-servers").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tcp_small_servers_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"udp_small_servers": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure udp-small-servers").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"udp_small_servers_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"lldp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure LLDP").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"lldp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"cdp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure CDP").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"cdp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"snmp_ifindex_persist": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure SNMP Ifindex Persist").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"snmp_ifindex_persist_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"console_logging": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Console Logging").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"console_logging_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"vty_logging": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure VTY Line Logging").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"vty_logging_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"line_vty": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Telnet (Outbound)").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"line_vty_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
		},
	}
}

func (r *CEdgeGlobalFeatureTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *CEdgeGlobalFeatureTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CEdgeGlobal

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
func (r *CEdgeGlobalFeatureTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CEdgeGlobal

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
func (r *CEdgeGlobalFeatureTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CEdgeGlobal

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
func (r *CEdgeGlobalFeatureTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CEdgeGlobal

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
func (r *CEdgeGlobalFeatureTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
