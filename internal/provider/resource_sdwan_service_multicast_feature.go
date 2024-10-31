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
var _ resource.Resource = &ServiceMulticastProfileParcelResource{}
var _ resource.ResourceWithImportState = &ServiceMulticastProfileParcelResource{}

func NewServiceMulticastProfileParcelResource() resource.Resource {
	return &ServiceMulticastProfileParcelResource{}
}

type ServiceMulticastProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *ServiceMulticastProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_multicast_feature"
}

func (r *ServiceMulticastProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Service Multicast Feature.").AddMinimumVersionDescription("20.12.0").String,

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
			"spt_only": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Shortest Path Tree (SPT) Only Mode").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"spt_only_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"local_replicator": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Replicator is local to this device").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"local_replicator_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"local_replicator_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set number of joins per group the router supports").AddIntegerRangeDescription(0, 131072).String,
				Optional:            true,
			},
			"local_replicator_threshold_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"igmp_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set IGMP interface parameters").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set interface name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(3, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`(GigabitEthernet|TwoGigabitEthernet|TenGigabitEthernet|FortyGigabitEthernet|HundredGigE|Vlan|Tunnel|Loopback)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
							},
						},
						"interface_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("igmp Version <1..3>").AddIntegerRangeDescription(1, 3).AddDefaultValueDescription("2").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 3),
							},
						},
						"join_groups": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure static joins").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"group_address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set group address").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^2(?:2[4-9]|3\d)(?:\.(?:25[0-5]|2[0-4]\d|1\d\d|[1-9]\d?|0)){3}$`), ""),
										},
									},
									"group_address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"source_address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set source address").String,
										Optional:            true,
									},
									"source_address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
			"pim_source_specific_multicast_enable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("turn SSM on/off").String,
				Required:            true,
			},
			"pim_source_specific_multicast_access_list": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set Access List for PIM SSM").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"pim_source_specific_multicast_access_list_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"pim_spt_threshold": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set when PIM router joins the SPT (kbps)").AddStringEnumDescription("0", "infinity").AddDefaultValueDescription("0").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("0", "infinity"),
				},
			},
			"pim_spt_threshold_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"pim_interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set PIM interface parameters").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set interface name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(3, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`(GigabitEthernet|TwoGigabitEthernet|TenGigabitEthernet|FortyGigabitEthernet|HundredGigE|Vlan|Tunnel|Loopback)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
							},
						},
						"interface_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"query_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set PIM query interval").AddIntegerRangeDescription(1, 18725).AddDefaultValueDescription("30").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 18725),
							},
						},
						"query_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"join_prune_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set interval at which PIM multicast traffic can join or be removed from RPT or SPT").AddIntegerRangeDescription(10, 600).AddDefaultValueDescription("60").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(10, 600),
							},
						},
						"join_prune_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"static_rp_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set Static RP Address(es)").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set Static RP IP Address").String,
							Optional:            true,
						},
						"ip_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"access_list": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set Static RP Access List").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"access_list_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"override": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set override flag").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"override_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"enable_auto_rp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable auto-RP").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"enable_auto_rp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"auto_rp_announces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable RP Announce").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set RP Announce Interface Name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(3, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`(GigabitEthernet|TwoGigabitEthernet|TenGigabitEthernet|FortyGigabitEthernet|HundredGigE|Vlan|Tunnel|Loopback)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
							},
						},
						"interface_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"scope": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set RP Announce Scope").AddIntegerRangeDescription(1, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
						"scope_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"auto_rp_discoveries": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable RP Discovery").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set RP Discovery Interface Name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(3, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`(GigabitEthernet|TwoGigabitEthernet|TenGigabitEthernet|FortyGigabitEthernet|HundredGigE|Vlan|Tunnel|Loopback)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
							},
						},
						"interface_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"scope": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set RP Discovery Scope").AddIntegerRangeDescription(1, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
						"scope_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"pim_bsr_rp_candidates": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set RP Discovery Scope").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set Autonomic-Networking virtual interface").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(3, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`(GigabitEthernet|TwoGigabitEthernet|TenGigabitEthernet|FortyGigabitEthernet|HundredGigE|Vlan|Tunnel|Loopback)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
							},
						},
						"interface_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"access_list_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set IP Access List for PIM RP Candidate").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"access_list_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set RP candidate advertisement interval").AddIntegerRangeDescription(1, 16383).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 16383),
							},
						},
						"interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set RP candidate priority").AddIntegerRangeDescription(0, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(255),
							},
						},
						"priority_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"pim_bsr_candidates": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("bsr candidate Attributes").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set Autonomic-Networking virtual interface").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(3, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`(GigabitEthernet|TwoGigabitEthernet|TenGigabitEthernet|FortyGigabitEthernet|HundredGigE|Vlan|Tunnel|Loopback)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
							},
						},
						"interface_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"hash_mask_length": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Hash Mask length for RP selection").AddIntegerRangeDescription(0, 32).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(32),
							},
						},
						"hash_mask_length_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set RP candidate priority").AddIntegerRangeDescription(0, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(255),
							},
						},
						"priority_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"accept_candidate_access_list": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set BSR RP candidate filter").String,
							Optional:            true,
						},
						"accept_candidate_access_list_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"msdp_groups": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("multicast MSDP peer").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"mesh_group_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set MSDP mesh group").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
								stringvalidator.RegexMatches(regexp.MustCompile(`^[^<! ]+$`), ""),
							},
						},
						"mesh_group_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"peers": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure peer").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"peer_ip": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set MSDP peer ip").String,
										Optional:            true,
									},
									"peer_ip_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"connection_source_interface": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set MSDP peer ip connect-source interface").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(3, 32),
											stringvalidator.RegexMatches(regexp.MustCompile(`(ATM|ATM-ACR|AppGigabitEthernet|AppNav-Compress|AppNav-UnCompress|Async|BD-VIF|BDI|CEM|CEM-ACR|Cellular|Dialer|Embedded-Service-Engine|Ethernet|Ethernet-Internal|FastEthernet|FiftyGigabitEthernet|FiveGigabitEthernet|FortyGigabitEthernet|FourHundredGigE|GMPLS|GigabitEthernet|Group-Async|HundredGigE|L2LISP|LISP|Loopback|MFR|Multilink|Port-channel|SM|Serial|Service-Engine|TenGigabitEthernet|Tunnel|TwentyFiveGigE|TwentyFiveGigabitEthernet|TwoGigabitEthernet|TwoHundredGigE|Vif|Virtual-PPP|Virtual-Template|VirtualPortGroup|Vlan|Wlan-GigabitEthernet|nat64|nat66|ntp|nve|ospfv3|overlay|pseudowire|ucse|vasileft|vasiright|vmi)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
										},
									},
									"connection_source_interface_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"remote_as": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set MSDP peer ip remote autonomous system number").AddIntegerRangeDescription(1, 65535).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
									},
									"remote_as_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"peer_authentication_password": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set MSDP peer ip password").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(1, 25),
										},
									},
									"peer_authentication_password_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"keepalive_interval": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set MSDP peer ip keepalive interval").AddIntegerRangeDescription(1, 60).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 60),
										},
									},
									"keepalive_interval_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"keepalive_hold_time": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set MSDP peer ip keepalive hold time").AddIntegerRangeDescription(1, 75).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 75),
										},
									},
									"keepalive_hold_time_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"sa_limit": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set MSDP peer ip SA limit message number").AddIntegerRangeDescription(1, 2147483646).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 2147483646),
										},
									},
									"sa_limit_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"default_peer": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set MSDP default peer").String,
										Optional:            true,
									},
									"prefix_list_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
										},
									},
								},
							},
						},
					},
				},
			},
			"msdp_originator_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set MSDP originator ID").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(3, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`(ATM|ATM-ACR|AppGigabitEthernet|AppNav-Compress|AppNav-UnCompress|Async|BD-VIF|BDI|CEM|CEM-ACR|Cellular|Dialer|Embedded-Service-Engine|Ethernet|Ethernet-Internal|FastEthernet|FiftyGigabitEthernet|FiveGigabitEthernet|FortyGigabitEthernet|FourHundredGigE|GMPLS|GigabitEthernet|Group-Async|HundredGigE|L2LISP|LISP|Loopback|MFR|Multilink|Port-channel|SM|Serial|Service-Engine|TenGigabitEthernet|Tunnel|TwentyFiveGigE|TwentyFiveGigabitEthernet|TwoGigabitEthernet|TwoHundredGigE|Vif|Virtual-PPP|Virtual-Template|VirtualPortGroup|Vlan|Wlan-GigabitEthernet|nat64|nat66|ntp|nve|ospfv3|overlay|pseudowire|ucse|vasileft|vasiright|vmi)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
				},
			},
			"msdp_originator_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"msdp_connection_retry_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set MSDP refresh timer").AddIntegerRangeDescription(1, 60).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 60),
				},
			},
			"msdp_connection_retry_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
		},
	}
}

func (r *ServiceMulticastProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *ServiceMulticastProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ServiceMulticast

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
func (r *ServiceMulticastProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ServiceMulticast

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
	if state.Version.IsNull() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *ServiceMulticastProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ServiceMulticast

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
func (r *ServiceMulticastProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ServiceMulticast

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
func (r *ServiceMulticastProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "service_multicast_feature_id" + ",feature_profile_id"
	if len(parts) != (count + 1) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: %s. Got: %q", pattern, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("feature_profile_id"), parts[1])...)
}

// End of section. //template:end import
