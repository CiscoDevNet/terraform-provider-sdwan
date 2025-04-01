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
var _ resource.Resource = &ApplicationPriorityTrafficPolicyProfileParcelResource{}
var _ resource.ResourceWithImportState = &ApplicationPriorityTrafficPolicyProfileParcelResource{}

func NewApplicationPriorityTrafficPolicyProfileParcelResource() resource.Resource {
	return &ApplicationPriorityTrafficPolicyProfileParcelResource{}
}

type ApplicationPriorityTrafficPolicyProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *ApplicationPriorityTrafficPolicyProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_priority_traffic_policy_policy"
}

func (r *ApplicationPriorityTrafficPolicyProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Application Priority Traffic Policy Policy.").AddMinimumVersionDescription("20.12.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Policy",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Policy",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Policy",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Policy",
				Optional:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Feature Profile ID").String,
				Required:            true,
			},
			"default_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("drop", "accept").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("drop", "accept"),
				},
			},
			"vpns": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"direction": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("service", "tunnel", "all").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("service", "tunnel", "all"),
				},
			},
			"sequences": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Traffic policy sequence list").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"sequence_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence Id").AddIntegerRangeDescription(1, 65536).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65536),
							},
						},
						"sequence_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence Name").String,
							Optional:            true,
						},
						"base_action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Base Action").AddStringEnumDescription("drop", "accept").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("drop", "accept"),
							},
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sequence IP Type").AddStringEnumDescription("ipv4", "ipv6", "all").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ipv4", "ipv6", "all"),
							},
						},
						"match_entries": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"application_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"saas_application_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"service_area": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("M365 Service Area").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"traffic_category": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("M365 Traffic Category").AddStringEnumDescription("optimize-allow", "optimize", "all").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("optimize-allow", "optimize", "all"),
										},
									},
									"dns_application_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"traffic_class": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Traffic Class").AddStringEnumDescription("gold-voip-telephony", "gold-broadcast-video", "gold-real-time-interactive", "gold-multimedia-conferencing", "gold-multimedia-streaming", "gold-network-control", "gold-signaling", "gold-ops-admin-mgmt", "gold-transactional-data", "gold-bulk-data", "silver", "bronze").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("gold-voip-telephony", "gold-broadcast-video", "gold-real-time-interactive", "gold-multimedia-conferencing", "gold-multimedia-streaming", "gold-network-control", "gold-signaling", "gold-ops-admin-mgmt", "gold-transactional-data", "gold-bulk-data", "silver", "bronze"),
										},
									},
									"dscp": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("DSCP number").AddIntegerRangeDescription(0, 63).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.AtMost(63),
										},
									},
									"packet_length": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Packet Length").String,
										Optional:            true,
									},
									"protocols": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("protocol (0-255) range or individual number separated by space").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"icmp_message": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("ICMP Message").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"icmp6_message": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("ICMP6 Message").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"source_data_ipv4_prefx_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"source_data_ipv6_prefx_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"source_ipv4_prefix": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source Data IP Prefix").String,
										Optional:            true,
									},
									"source_ipv6_prefix": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source Data IP Prefix").String,
										Optional:            true,
									},
									"source_ports": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Source Port (0-65535) range or individual number separated by space").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"destination_data_ipv4_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"destination_data_ipv6_prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
									"destination_ipv4_prefix": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination Data IP Prefix").String,
										Optional:            true,
									},
									"destination_ipv6_prefix": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination Data IP Prefix").String,
										Optional:            true,
									},
									"destination_ports": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination Port (0-65535) range or individual number separated by space").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"tcp": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("TCP States").AddStringEnumDescription("syn").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("syn"),
										},
									},
									"destination_region": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination Region").AddStringEnumDescription("primary-region", "secondary-region", "other-region").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("primary-region", "secondary-region", "other-region"),
										},
									},
									"traffic_to": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Traffic to").AddStringEnumDescription("core", "service", "access").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("core", "service", "access"),
										},
									},
									"dns": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Dns").AddStringEnumDescription("request", "response").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("request", "response"),
										},
									},
								},
							},
						},
						"actions": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"sla_classes": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("slaClass").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"sla_class_list_id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
													},
												},
												"preferred_color": schema.SetAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													ElementType:         types.StringType,
													Optional:            true,
												},
												"preferred_color_group_list_id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
													},
												},
												"strict": schema.BoolAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
												},
												"fallback_to_best_path": schema.BoolAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
												},
												"preferred_remote_color": schema.SetAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													ElementType:         types.StringType,
													Optional:            true,
												},
												"remote_color_restrict": schema.BoolAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
												},
											},
										},
									},
									"backup_sla_preferred_color": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Backup SLA perferred color").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"set_parameters": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"dscp": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(0, 63).String,
													Optional:            true,
													Validators: []validator.Int64{
														int64validator.AtMost(63),
													},
												},
												"policer_id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
													},
												},
												"preferred_color_group_id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
													},
												},
												"forwarding_class_list_id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
													},
												},
												"local_tloc_list_color": schema.SetAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													ElementType:         types.StringType,
													Optional:            true,
												},
												"local_tloc_list_restrict": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
												},
												"local_tloc_list_encapsulation": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("ipsec", "gre").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("ipsec", "gre"),
													},
												},
												"preferred_remote_color_id": schema.SetAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													ElementType:         types.StringType,
													Optional:            true,
												},
												"preferred_remote_color_restrict": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
												},
												"tloc_color": schema.SetAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													ElementType:         types.StringType,
													Optional:            true,
												},
												"tloc_encapsulation": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("ipsec", "gre").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("ipsec", "gre"),
													},
												},
												"tloc_ip": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
												},
												"tloc_list_id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
													},
												},
												"service_tloc_color": schema.SetAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													ElementType:         types.StringType,
													Optional:            true,
												},
												"service_tloc_encapsulation": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("ipsec", "gre").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("ipsec", "gre"),
													},
												},
												"service_tloc_ip": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
												},
												"service_vpn": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.LengthBetween(1, 64),
														stringvalidator.RegexMatches(regexp.MustCompile(`^\{\{[./\[\]a-zA-Z0-9_-]+\}\}$`), ""),
													},
												},
												"service_type": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("FW", "IDS", "IDP", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "appqoe").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("FW", "IDS", "IDP", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "appqoe"),
													},
												},
												"service_tloc_list_id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
													},
												},
												"service_chain_type": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("SC1", "SC2", "SC4", "SC5", "SC6", "SC7", "SC8", "SC9", "SC10", "SC11", "SC12", "SC13", "SC14", "SC15", "SC16").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("SC1", "SC2", "SC4", "SC5", "SC6", "SC7", "SC8", "SC9", "SC10", "SC11", "SC12", "SC13", "SC14", "SC15", "SC16"),
													},
												},
												"service_chain_vpn": schema.Int64Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(0, 65530).String,
													Optional:            true,
													Validators: []validator.Int64{
														int64validator.AtMost(65530),
													},
												},
												"service_chain_local": schema.BoolAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
												},
												"service_chain_fallback_to_routing": schema.BoolAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
												},
												"service_chain_tloc_color": schema.SetAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													ElementType:         types.StringType,
													Optional:            true,
												},
												"service_chain_tloc_encapsulation": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("ipsec", "gre").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.OneOf("ipsec", "gre"),
													},
												},
												"service_chain_tloc_ip": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
												},
												"service_chain_tloc_list_id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
													},
												},
												"next_hop_ipv4": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
												},
												"next_hop_ipv6": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
												},
												"next_hop_loose": schema.BoolAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
												},
												"vpn": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("").String,
													Optional:            true,
													Validators: []validator.String{
														stringvalidator.LengthBetween(1, 64),
														stringvalidator.RegexMatches(regexp.MustCompile(`^\{\{[./\[\]a-zA-Z0-9_-]+\}\}$`), ""),
													},
												},
											},
										},
									},
									"redirect_dns_field": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("ipAddress", "redirectDns").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("ipAddress", "redirectDns"),
										},
									},
									"redirect_dns_value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
									},
									"loss_correct_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("fecAdaptive", "fecAlways", "packetDuplication").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("fecAdaptive", "fecAlways", "packetDuplication"),
										},
									},
									"loss_correct_fec_threshold": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(1, 5).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 5),
										},
									},
									"count": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(1, 20),
										},
									},
									"log": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
									},
									"cloud_saas": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
									},
									"cloud_probe": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
									},
									"nat_pool": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("").AddIntegerRangeDescription(1, 31).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 31),
										},
									},
									"nat_vpn": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
									},
									"nat_fallback": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
									},
									"nat_bypass": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
									},
									"nat_dia_pool": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.Int64Type,
										Optional:            true,
									},
									"nat_dia_interface": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"secure_internet_gateway": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
									},
									"fallback_to_routing": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
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

func (r *ApplicationPriorityTrafficPolicyProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *ApplicationPriorityTrafficPolicyProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplicationPriorityTrafficPolicy

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

	plan.Id = types.StringValue(res.Get("id").String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *ApplicationPriorityTrafficPolicyProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplicationPriorityTrafficPolicy

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

	if imp {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}
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
func (r *ApplicationPriorityTrafficPolicyProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ApplicationPriorityTrafficPolicy

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
func (r *ApplicationPriorityTrafficPolicyProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplicationPriorityTrafficPolicy

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
func (r *ApplicationPriorityTrafficPolicyProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "application_priority_traffic_policy_policy_id" + ",feature_profile_id"
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
