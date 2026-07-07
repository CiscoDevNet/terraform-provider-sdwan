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
var _ resource.Resource = &SSECiscoProfileParcelResource{}
var _ resource.ResourceWithImportState = &SSECiscoProfileParcelResource{}

func NewSSECiscoProfileParcelResource() resource.Resource {
	return &SSECiscoProfileParcelResource{}
}

type SSECiscoProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *SSECiscoProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sse_cisco_feature"
}

func (r *SSECiscoProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a SSE Cisco Feature.").AddMinimumVersionDescription("20.15.0").String,

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
			"context_sharing_for_vpn": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable/disable Cisco SSE context sharing for vpn").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"context_sharing_for_sgt": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable/disable Cisco SSE context sharing for sgt").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface name: IPsec when present").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface name: ipsec(1..255)").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(5, 8),
								stringvalidator.RegexMatches(regexp.MustCompile(`^ipsec(?:1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)$`), ""),
							},
						},
						"shutdown": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Administrative state").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"shutdown_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"tunnel_source_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("<1..32 characters> Interface name: ge0/<0-..> or ge0/<0-..>.vlanid").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(3, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`(ATM|ATM-ACR|AppGigabitEthernet|AppNav-Compress|AppNav-UnCompress|Async|BD-VIF|BDI|CEM|CEM-ACR|Cellular|Dialer|Embedded-Service-Engine|Ethernet|Ethernet-Internal|FastEthernet|FiftyGigabitEthernet|FiveGigabitEthernet|FortyGigabitEthernet|FourHundredGigE|GMPLS|GigabitEthernet|Group-Async|HundredGigE|L2LISP|LISP|Loopback|MFR|Multilink|Port-channel|SM|Serial|Service-Engine|TenGigabitEthernet|Tunnel|TwentyFiveGigE|TwentyFiveGigabitEthernet|TwoGigabitEthernet|TwoHundredGigE|Vif|Virtual-PPP|Virtual-Template|VirtualPortGroup|Vlan|Wlan-GigabitEthernet|nat64|nat66|ntp|nve|ospfv3|overlay|pseudowire|ucse|vasileft|vasiright|vmi)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
							},
						},
						"tunnel_source_interface_variable": schema.StringAttribute{
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
						"tunnel_dc_preference": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("SSE Tunnel Data Center").AddStringEnumDescription("primary-dc", "secondary-dc").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("primary-dc", "secondary-dc"),
							},
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
						"mtu": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface MTU <576..2000>, in bytes").AddIntegerRangeDescription(576, 2000).AddDefaultValueDescription("1400").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(576, 2000),
							},
						},
						"mtu_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dpd_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE keepalive interval (seconds)").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("10").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(65535),
							},
						},
						"dpd_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dpd_retries": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE keepalive retries").AddIntegerRangeDescription(0, 255).AddDefaultValueDescription("3").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(255),
							},
						},
						"dpd_retries_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ike_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE Version <1..2>").AddIntegerRangeDescription(1, 2).AddDefaultValueDescription("2").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 2),
							},
						},
						"ike_version_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ike_rekey_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE rekey interval <300..1209600> seconds").AddIntegerRangeDescription(300, 1209600).AddDefaultValueDescription("14400").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(300, 1209600),
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
							MarkdownDescription: helpers.NewAttributeDescription("IKE Diffie Hellman Groups").AddStringEnumDescription("2", "5", "14", "15", "16", "19", "20", "21").AddDefaultValueDescription("16").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2", "5", "14", "15", "16", "19", "20", "21"),
							},
						},
						"ike_group_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ipsec_rekey_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPsec rekey interval <300..1209600> seconds").AddIntegerRangeDescription(300, 1209600).AddDefaultValueDescription("3600").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(300, 1209600),
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
							MarkdownDescription: helpers.NewAttributeDescription("IPsec(ESP) encryption and integrity protocol").AddStringEnumDescription("aes256-cbc-sha1", "aes256-cbc-sha384", "aes256-cbc-sha256", "aes256-cbc-sha512", "aes256-gcm").AddDefaultValueDescription("aes256-cbc-sha512").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("aes256-cbc-sha1", "aes256-cbc-sha384", "aes256-cbc-sha256", "aes256-cbc-sha512", "aes256-gcm"),
							},
						},
						"ipsec_ciphersuite_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"perfect_forward_secrecy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPsec perfect forward secrecy settings").AddStringEnumDescription("group-2", "group-5", "group-14", "group-15", "group-16", "group-19", "group-20", "group-21", "none").AddDefaultValueDescription("group-16").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("group-2", "group-5", "group-14", "group-15", "group-16", "group-19", "group-20", "group-21", "none"),
							},
						},
						"perfect_forward_secrecy_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"tracker": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable tracker for this interface").AddDefaultValueDescription("DefaultTracker").String,
							Optional:            true,
						},
						"track_enable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable/disable Cisco SSE tracking").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"track_enable_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"interface_pairs": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface Pair for active and backup").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"active_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Active Tunnel Interface for SSE").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(5, 8),
								stringvalidator.RegexMatches(regexp.MustCompile(`^ipsec(?:1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)$`), ""),
							},
						},
						"active_interface_weight": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Active Tunnel Interface Weight").AddIntegerRangeDescription(1, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
						"backup_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Backup Tunnel Interface for Cisco SSE").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(4, 8),
								stringvalidator.RegexMatches(regexp.MustCompile(`^(None|ipsec(?:1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d))$`), ""),
							},
						},
						"backup_interface_weight": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Backup Tunnel Interface Weight").AddIntegerRangeDescription(1, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
					},
				},
			},
			"region": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Region for Primary and Secondary Datacenter").AddDefaultValueDescription("auto").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"region_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tracker_source_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Source IP address for Tracker").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)\.){3}(25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)$`), ""),
				},
			},
			"tracker_source_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"trackers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Tracker configuration").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tracker name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 128),
							},
						},
						"endpoint_api_url": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("API url of endpoint").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthAtMost(512),
								stringvalidator.RegexMatches(regexp.MustCompile(`^http\:\/\/[0-9a-zA-Z]([-.\w]*[0-9a-zA-Z])*(:(0-9)*)*(\/?)([a-zA-Z0-9\-\.\?\,\'\/\\\+&%\$#_]*)?$`), ""),
							},
						},
						"endpoint_api_url_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"threshold": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Probe Timeout threshold <100..1000> milliseconds").AddIntegerRangeDescription(100, 1000).AddDefaultValueDescription("1000").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(100, 1000),
							},
						},
						"threshold_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Probe interval <10..600> seconds").AddIntegerRangeDescription(20, 600).AddDefaultValueDescription("30").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(20, 600),
							},
						},
						"interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"multiplier": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Probe failure multiplier <1..10> failed attempts").AddIntegerRangeDescription(1, 10).AddDefaultValueDescription("2").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 10),
							},
						},
						"multiplier_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *SSECiscoProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *SSECiscoProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SSECisco

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

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *SSECiscoProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SSECisco

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
	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	state.fromBody(ctx, res, imp)
	if state.Version.IsNull() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *SSECiscoProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SSECisco

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
func (r *SSECiscoProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SSECisco

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
func (r *SSECiscoProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "sse_cisco_feature_id" + ",feature_profile_id"
	if len(parts) != (count + 1) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: %s. Got: %q", pattern, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("feature_profile_id"), parts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
