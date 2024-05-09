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
var _ resource.Resource = &TrafficDataPolicyDefinitionResource{}
var _ resource.ResourceWithImportState = &TrafficDataPolicyDefinitionResource{}

func NewTrafficDataPolicyDefinitionResource() resource.Resource {
	return &TrafficDataPolicyDefinitionResource{}
}

type TrafficDataPolicyDefinitionResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *TrafficDataPolicyDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_traffic_data_policy_definition"
}

func (r *TrafficDataPolicyDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Traffic Data Policy Definition .").String,

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
			"type": schema.StringAttribute{
				MarkdownDescription: "Type",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the policy definition").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The description of the policy definition").String,
				Required:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Default action, either `accept` or `drop`").AddStringEnumDescription("accept", "drop").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("accept", "drop"),
				},
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of sequences").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence ID").String,
							Required:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence name").String,
							Required:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence type").AddStringEnumDescription("applicationFirewall", "qos", "serviceChaining", "trafficEngineering", "data").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("applicationFirewall", "qos", "serviceChaining", "trafficEngineering", "data"),
							},
						},
						"ip_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence IP type, either `ipv4`, `ipv6` or `all`").AddStringEnumDescription("ipv4", "ipv6", "all").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ipv4", "ipv6", "all"),
							},
						},
						"base_action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Base action, either `accept` or `drop`").AddStringEnumDescription("accept", "drop").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("accept", "drop"),
							},
						},
						"match_entries": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of match entries").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of match entry").AddStringEnumDescription("appList", "dnsAppList", "dns", "dscp", "packetLength", "plp", "protocol", "sourceDataPrefixList", "sourceIp", "sourcePort", "destinationDataPrefixList", "destinationIp", "destinationRegion", "destinationPort", "tcp", "trafficTo").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("appList", "dnsAppList", "dns", "dscp", "packetLength", "plp", "protocol", "sourceDataPrefixList", "sourceIp", "sourcePort", "destinationDataPrefixList", "destinationIp", "destinationRegion", "destinationPort", "tcp", "trafficTo"),
										},
									},
									"application_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Application list ID").String,
										Optional:            true,
									},
									"application_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Application list version").String,
										Optional:            true,
									},
									"dns_application_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("DNS Application list ID").String,
										Optional:            true,
									},
									"dns_application_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("DNS Application list version").String,
										Optional:            true,
									},
									"dns": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("DNS request or response").AddStringEnumDescription("request", "response").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("request", "response"),
										},
									},
									"dscp": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("DSCP value").AddIntegerRangeDescription(0, 63).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 63),
										},
									},
									"packet_length": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Packet length").AddIntegerRangeDescription(0, 65535).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 65535),
										},
									},
									"plp": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("PLP").AddStringEnumDescription("low", "high").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("low", "high"),
										},
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP Protocol, 0-255 (Single value or multiple values separated by spaces)").String,
										Optional:            true,
									},
									"source_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source Data Prefix list ID").String,
										Optional:            true,
									},
									"source_data_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source Data Prefix list version").String,
										Optional:            true,
									},
									"source_ip": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source IP").String,
										Optional:            true,
									},
									"source_port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source port, 0-65535 (Single value, range or multiple values separated by spaces)").String,
										Optional:            true,
									},
									"destination_data_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination Data Prefix list ID").String,
										Optional:            true,
									},
									"destination_data_prefix_list_version": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination Data Prefix list version").String,
										Optional:            true,
									},
									"destination_ip": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination IP").String,
										Optional:            true,
									},
									"destination_port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination port, 0-65535 (Single value, range or multiple values separated by spaces)").String,
										Optional:            true,
									},
									"destination_region": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination region").AddStringEnumDescription("primary-region", "secondary-region", "other-region").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("primary-region", "secondary-region", "other-region"),
										},
									},
									"tcp": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("TCP flags").AddStringEnumDescription("syn").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("syn"),
										},
									},
									"traffic_to": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Traffic to").AddStringEnumDescription("access", "core", "service").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("access", "core", "service"),
										},
									},
								},
							},
						},
						"action_entries": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of action entries").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of action entry").AddStringEnumDescription("cflowd", "count", "dreOptimization", "fallbackToRouting", "log", "lossProtect", "lossProtectFec", "nat", "redirectDns", "serviceNodeGroup", "set", "sig", "tcpOptimization").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("cflowd", "count", "dreOptimization", "fallbackToRouting", "log", "lossProtect", "lossProtectFec", "nat", "redirectDns", "serviceNodeGroup", "set", "sig", "tcpOptimization"),
										},
									},
									"cflowd": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Enable cflowd").String,
										Optional:            true,
									},
									"counter": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Counter name").String,
										Optional:            true,
									},
									"dre_optimization": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Enable DRE optimization").String,
										Optional:            true,
									},
									"fallback_to_routing": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Enable fallback to routing").String,
										Optional:            true,
									},
									"log": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Enable logging").String,
										Optional:            true,
									},
									"loss_correction": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Loss correction").AddStringEnumDescription("fecAdaptive", "fecAlways", "packetDuplication").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("fecAdaptive", "fecAlways", "packetDuplication"),
										},
									},
									"loss_correction_fec": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Loss correction FEC").AddStringEnumDescription("fecAdaptive", "fecAlways", "packetDuplication").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("fecAdaptive", "fecAlways", "packetDuplication"),
										},
									},
									"loss_correction_fec_threshold": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Loss correction FEC threshold").AddIntegerRangeDescription(1, 5).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 5),
										},
									},
									"nat_pool": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("NAT pool").AddStringEnumDescription("pool").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("pool"),
										},
									},
									"nat_pool_id": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("NAT pool ID").AddIntegerRangeDescription(1, 31).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 31),
										},
									},
									"redirect_dns": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Redirect DNS").AddStringEnumDescription("dnsType", "ipAddress").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("dnsType", "ipAddress"),
										},
									},
									"redirect_dns_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Redirect DNS type").AddStringEnumDescription("host", "umbrella").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("host", "umbrella"),
										},
									},
									"redirect_dns_address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Redirect DNS IP address").String,
										Optional:            true,
									},
									"service_node_group": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Service node group").String,
										Optional:            true,
									},
									"secure_internet_gateway": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Enable secure internet gateway").String,
										Optional:            true,
									},
									"tcp_optimization": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Enable TCP optimization").String,
										Optional:            true,
									},
									"set_parameters": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("List of set parameters").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Type of set parameter").AddStringEnumDescription("dscp", "forwardingClass", "localTlocList", "nextHop", "nextHopLoose", "policer", "preferredColorGroup", "tlocList", "tloc", "service", "vpn").String,
													Required:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("dscp", "forwardingClass", "localTlocList", "nextHop", "nextHopLoose", "policer", "preferredColorGroup", "tlocList", "tloc", "service", "vpn"),
													},
												},
												"dscp": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("DSCP").AddIntegerRangeDescription(0, 63).String,
													Optional:            true,
													Validators: []validator.Int64{
														int64validator.Between(0, 63),
													},
												},
												"forwarding_class": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Forwarding class").String,
													Optional:            true,
												},
												"next_hop": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Next hop IP").String,
													Optional:            true,
												},
												"local_tloc_list_color": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Local TLOC list color. Space separated list of colors.").String,
													Optional:            true,
												},
												"local_tloc_list_encap": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Local TLOC list encapsulation.").AddStringEnumDescription("ipsec", "gre", "ipsec gre").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("ipsec", "gre", "ipsec gre"),
													},
												},
												"local_tloc_list_restrict": schema.BoolAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Local TLOC list restrict").String,
													Optional:            true,
												},
												"next_hop_loose": schema.BoolAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Use routing table entry to forward the packet in case Next-hop is not available").String,
													Optional:            true,
												},
												"policer_list_id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Policer list ID").String,
													Optional:            true,
												},
												"policer_list_version": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("Policer list version").String,
													Optional:            true,
												},
												"preferred_color_group_list": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Preferred color group list ID").String,
													Optional:            true,
												},
												"preferred_color_group_list_version": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("Preferred color group list version").String,
													Optional:            true,
												},
												"tloc_list_id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("TLOC list ID").String,
													Optional:            true,
												},
												"tloc_list_version": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("TLOC list version").String,
													Optional:            true,
												},
												"tloc_ip": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("TLOC IP address").String,
													Optional:            true,
												},
												"tloc_color": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("TLOC color").String,
													Optional:            true,
												},
												"tloc_encapsulation": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("TLOC encapsulation").AddStringEnumDescription("ipsec", "gre", "ipsec gre").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("ipsec", "gre", "ipsec gre"),
													},
												},
												"service_type": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Service type").AddStringEnumDescription("FW", "IDP", "IDS", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "netsvc5").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("FW", "IDP", "IDS", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "netsvc5"),
													},
												},
												"service_vpn_id": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("Service VPN ID").AddIntegerRangeDescription(0, 65536).String,
													Optional:            true,
													Validators: []validator.Int64{
														int64validator.Between(0, 65536),
													},
												},
												"service_tloc_list_id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Service TLOC list ID").String,
													Optional:            true,
												},
												"service_tloc_list_version": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("Service TLOC list version").String,
													Optional:            true,
												},
												"service_tloc_ip": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Service TLOC IP address").String,
													Optional:            true,
												},
												"service_tloc_local": schema.BoolAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Service TLOC Local").String,
													Optional:            true,
												},
												"service_tloc_restrict": schema.BoolAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Service TLOC Restrict").String,
													Optional:            true,
												},
												"service_tloc_color": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Service TLOC color").String,
													Optional:            true,
												},
												"service_tloc_encapsulation": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Service TLOC encapsulation").AddStringEnumDescription("ipsec", "gre", "ipsec gre").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("ipsec", "gre", "ipsec gre"),
													},
												},
												"vpn_id": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("DSCP").AddIntegerRangeDescription(0, 65530).String,
													Optional:            true,
													Validators: []validator.Int64{
														int64validator.Between(0, 65530),
													},
												},
											},
										},
									},
									"nat_parameters": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("List of NAT parameters").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Type of NAT parameter").AddStringEnumDescription("useVpn", "fallback").String,
													Required:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("useVpn", "fallback"),
													},
												},
												"vpn_id": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("DSCP").String,
													Optional:            true,
												},
												"fallback": schema.BoolAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Fallback").String,
													Optional:            true,
												},
											},
										},
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

func (r *TrafficDataPolicyDefinitionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *TrafficDataPolicyDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TrafficDataPolicyDefinition

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
	plan.Id = types.StringValue(res.Get("definitionId").String())
	plan.Version = types.Int64Value(0)
	plan.Type = types.StringValue("data")

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *TrafficDataPolicyDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state TrafficDataPolicyDefinition

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if strings.Contains(res.Get("error.message").String(), "Failed to find specified resource") || strings.Contains(res.Get("error.message").String(), "Invalid template type") || strings.Contains(res.Get("error.message").String(), "Template definition not found") || strings.Contains(res.Get("error.message").String(), "Invalid Profile Id") {
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
func (r *TrafficDataPolicyDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state TrafficDataPolicyDefinition

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

	if plan.hasChanges(ctx, &state) {
		body := plan.toBody(ctx)
		r.updateMutex.Lock()
		res, err := r.client.Put(plan.getPath()+url.QueryEscape(plan.Id.ValueString()), body)
		r.updateMutex.Unlock()
		if err != nil {
			if strings.Contains(res.Get("error.message").String(), "Failed to acquire lock") {
				resp.Diagnostics.AddWarning("Client Warning", "Failed to modify policy due to policy being locked by another change. Policy changes will not be applied. Re-run 'terraform apply' to try again.")
			} else if strings.Contains(res.Get("error.message").String(), "Template locked in edit mode") {
				resp.Diagnostics.AddWarning("Client Warning", "Failed to modify template due to template being locked by another change. Template changes will not be applied. Re-run 'terraform apply' to try again.")
			} else {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
				return
			}
		}
	} else {
		tflog.Debug(ctx, fmt.Sprintf("%s: No changes detected", plan.Name.ValueString()))
	}
	plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *TrafficDataPolicyDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TrafficDataPolicyDefinition

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *TrafficDataPolicyDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
