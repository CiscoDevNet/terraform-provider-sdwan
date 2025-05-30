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
var _ resource.Resource = &VPNInterfaceSVIFeatureTemplateResource{}
var _ resource.ResourceWithImportState = &VPNInterfaceSVIFeatureTemplateResource{}

func NewVPNInterfaceSVIFeatureTemplateResource() resource.Resource {
	return &VPNInterfaceSVIFeatureTemplateResource{}
}

type VPNInterfaceSVIFeatureTemplateResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *VPNInterfaceSVIFeatureTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_interface_svi_feature_template"
}

func (r *VPNInterfaceSVIFeatureTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a VPN Interface SVI feature template.").AddMinimumVersionDescription("15.0.0").String,

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
			"if_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface name: VLAN 1 - VLAN 4094 when present").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 8),
				},
			},
			"if_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"interface_description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface description").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
				},
			},
			"interface_description_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Assign IPv4 address").String,
				Optional:            true,
			},
			"ipv4_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv4_secondary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Assign secondary IP addresses").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ipv4_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP Address").String,
							Optional:            true,
						},
						"ipv4_address_variable": schema.StringAttribute{
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
			"ipv6_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Assign IPv6 address").String,
				Optional:            true,
			},
			"ipv6_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv6_dhcp_client": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable DHCPv6").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ipv6_dhcp_client_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv6_dhcp_distance": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set administrative distance for DHCP default route").AddIntegerRangeDescription(1, 65536).AddDefaultValueDescription("1").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65536),
				},
			},
			"ipv6_dhcp_distance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv6_dhcp_rapid_commit": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable DHCPv6 rapid commit").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ipv6_dhcp_rapid_commit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv6_secondary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Assign secondary IPv6 addresses").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ipv6_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 Address").String,
							Optional:            true,
						},
						"ipv6_address_variable": schema.StringAttribute{
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
			"ipv4_dhcp_helper": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of DHCP helper addresses").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ipv4_dhcp_helper_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv6_dhcp_helpers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DHCPv6 Helper").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("DHCPv6 Helper address").String,
							Optional:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("DHCPv6 Helper VPN").AddIntegerRangeDescription(1, 65536).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65536),
							},
						},
						"vpn_id_variable": schema.StringAttribute{
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
			"ip_directed_broadcast": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP Directed-Broadcast").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ip_directed_broadcast_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface MTU <1500..9216> in bytes").AddIntegerRangeDescription(1500, 9216).AddDefaultValueDescription("1500").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1500, 9216),
				},
			},
			"mtu_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ip_mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP MTU <576..Interface MTU>, in bytes").AddIntegerRangeDescription(576, 9216).AddDefaultValueDescription("1500").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(576, 9216),
				},
			},
			"ip_mtu_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tcp_mss_adjust": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("TCP MSS on SYN packets, in bytes").AddIntegerRangeDescription(552, 1960).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(552, 1960),
				},
			},
			"tcp_mss_adjust_variable": schema.StringAttribute{
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
			"arp_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timeout value for dynamically learned ARP entries, <0..2678400> seconds").AddIntegerRangeDescription(0, 2678400).AddDefaultValueDescription("1200").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2678400),
				},
			},
			"arp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv4_access_lists": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Apply ACL").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Direction").AddStringEnumDescription("in", "out").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("in", "out"),
							},
						},
						"acl_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of access list").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 127),
							},
						},
						"acl_name_variable": schema.StringAttribute{
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
			"ipv6_access_lists": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Apply ACL").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Direction").AddStringEnumDescription("in", "out").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("in", "out"),
							},
						},
						"acl_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of access list").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 127),
							},
						},
						"acl_name_variable": schema.StringAttribute{
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
			"policers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable policer").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Direction").AddStringEnumDescription("in", "out").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("in", "out"),
							},
						},
						"policer_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of policer").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"static_arp_entries": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure static ARP entries").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ipv4_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP Address").String,
							Optional:            true,
						},
						"ipv4_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"mac_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("MAC address").String,
							Optional:            true,
						},
						"mac_address_variable": schema.StringAttribute{
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
			"ipv4_vrrps": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable VRRP").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Group ID").AddIntegerRangeDescription(1, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
						"group_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set priority").AddIntegerRangeDescription(1, 254).AddDefaultValueDescription("100").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 254),
							},
						},
						"priority_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"timer": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Timer interval for successive advertisements, in milliseconds").AddIntegerRangeDescription(100, 40950).AddDefaultValueDescription("100").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(100, 40950),
							},
						},
						"timer_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"track_omp": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Track OMP status").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"track_omp_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"track_prefix_list": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Track Prefix List").String,
							Optional:            true,
						},
						"track_prefix_list_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ipv4_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Assign IP Address").String,
							Optional:            true,
						},
						"ipv4_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ipv4_secondary_addresses": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("VRRP Secondary IP address").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ipv4_address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("VRRP Secondary IP address").String,
										Optional:            true,
									},
									"ipv4_address_variable": schema.StringAttribute{
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
						"tloc_preference_change": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("change TLOC preference").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"tloc_preference_change_value": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set tloc preference change value").AddIntegerRangeDescription(1, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"tloc_preference_change_value_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"tracking_objects": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("tracking object for VRRP configuration").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Tracker ID").AddIntegerRangeDescription(1, 1000).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 1000),
										},
									},
									"name_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"track_action": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Track Action").AddStringEnumDescription("decrement", "shutdown").AddDefaultValueDescription("decrement").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("decrement", "shutdown"),
										},
									},
									"track_action_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"decrement_value": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Decrement Value for VRRP priority").AddIntegerRangeDescription(1, 255).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 255),
										},
									},
									"decrement_value_variable": schema.StringAttribute{
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
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"ipv6_vrrps": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable VRRP").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Group ID").AddIntegerRangeDescription(1, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
						"group_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set priority").AddIntegerRangeDescription(1, 254).AddDefaultValueDescription("100").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 254),
							},
						},
						"priority_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"timer": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Timer interval for successive advertisements, in milliseconds").AddIntegerRangeDescription(100, 40950).AddDefaultValueDescription("100").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(100, 40950),
							},
						},
						"timer_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"track_omp": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Track OMP status").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"track_omp_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"track_prefix_list": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Track Prefix List").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"track_prefix_list_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ipv6_addresses": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 VRRP").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"link_local_address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Use link-local IPv6 Address").String,
										Optional:            true,
									},
									"link_local_address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"prefix": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Assign Global IPv6 Prefix").String,
										Optional:            true,
									},
									"prefix_variable": schema.StringAttribute{
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
						"ipv6_secondary_addresses": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 Secondary IP address").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IPv6 Secondary IP address").String,
										Optional:            true,
									},
									"prefix_variable": schema.StringAttribute{
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
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *VPNInterfaceSVIFeatureTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *VPNInterfaceSVIFeatureTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VPNInterfaceSVI

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
func (r *VPNInterfaceSVIFeatureTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VPNInterfaceSVI

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
	if state.Version.IsNull() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *VPNInterfaceSVIFeatureTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state VPNInterfaceSVI

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
func (r *VPNInterfaceSVIFeatureTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VPNInterfaceSVI

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
func (r *VPNInterfaceSVIFeatureTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
