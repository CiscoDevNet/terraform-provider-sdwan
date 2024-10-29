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
var _ resource.Resource = &CiscoVPNFeatureTemplateResource{}
var _ resource.ResourceWithImportState = &CiscoVPNFeatureTemplateResource{}

func NewCiscoVPNFeatureTemplateResource() resource.Resource {
	return &CiscoVPNFeatureTemplateResource{}
}

type CiscoVPNFeatureTemplateResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *CiscoVPNFeatureTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_vpn_feature_template"
}

func (r *CiscoVPNFeatureTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Cisco VPN feature template.").AddMinimumVersionDescription("15.0.0").String,

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
			"vpn_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of VPN instances").AddIntegerRangeDescription(0, 65527).AddDefaultValueDescription("0").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65527),
				},
			},
			"vpn_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"vpn_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tenant_vpn_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Tenant VPN").AddIntegerRangeDescription(0, 65527).AddDefaultValueDescription("0").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 65527),
				},
			},
			"organization_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Org Name selected").String,
				Optional:            true,
			},
			"omp_admin_distance_ipv4": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("omp-admin-distance-ipv4").AddIntegerRangeDescription(1, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"omp_admin_distance_ipv4_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"omp_admin_distance_ipv6": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("omp-admin-distance-ipv6").AddIntegerRangeDescription(1, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"omp_admin_distance_ipv6_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"enhance_ecmp_keying": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Optional packet fields for ECMP keying").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"enhance_ecmp_keying_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"dns_ipv4_servers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DNS").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("DNS Address").String,
							Optional:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"role": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Role").AddStringEnumDescription("primary", "secondary").AddDefaultValueDescription("primary").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("primary", "secondary"),
							},
						},
						"role_variable": schema.StringAttribute{
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
			"dns_ipv6_servers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DNS").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("DNS Address").String,
							Optional:            true,
						},
						"role": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Role").AddStringEnumDescription("primary", "secondary").AddDefaultValueDescription("primary").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("primary", "secondary"),
							},
						},
						"role_variable": schema.StringAttribute{
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
			"dns_hosts": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static DNS mapping").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"hostname": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Hostname").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 128),
							},
						},
						"hostname_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ip": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of IP").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"ip_variable": schema.StringAttribute{
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
			"services": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure services").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"service_types": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Service Type").AddStringEnumDescription("FW", "IDS", "IDP", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "TE", "appqoe").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("FW", "IDS", "IDP", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "TE", "appqoe"),
							},
						},
						"address": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of IPv4 address").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tracking Service").String,
							Optional:            true,
						},
						"interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"track_enable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tracking Service").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"track_enable_variable": schema.StringAttribute{
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
			"ipv4_static_service_routes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure IPv4 Static Service Routes").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prefix").String,
							Optional:            true,
						},
						"prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination VPN to resolve the prefix").AddDefaultValueDescription("0").String,
							Optional:            true,
						},
						"service": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Service").AddStringEnumDescription("sig").AddDefaultValueDescription("sig").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("sig"),
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"ipv4_static_routes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure IPv4 Static Routes").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prefix").String,
							Optional:            true,
						},
						"prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"null0": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("null0").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"null0_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"distance": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Administrative distance").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("1").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
						"distance_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination VPN(!=0 or !=512) to resolve the prefix").AddDefaultValueDescription("0").String,
							Optional:            true,
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dhcp": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Default Gateway obtained from DHCP").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"dhcp_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"next_hops": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP gateway address").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP Address").String,
										Optional:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"distance": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Administrative distance").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("1").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 255),
										},
									},
									"distance_variable": schema.StringAttribute{
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
						"track_next_hops": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP gateway address").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP Address").String,
										Optional:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"distance": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Administrative distance").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("1").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 255),
										},
									},
									"distance_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"tracker": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Static route tracker").String,
										Optional:            true,
									},
									"tracker_variable": schema.StringAttribute{
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
			"ipv6_static_routes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure IPv6 Static Routes").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prefix").String,
							Optional:            true,
						},
						"prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"null0": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("null0").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"null0_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination VPN(!=0 or !=512) to resolve the prefix").AddDefaultValueDescription("0").String,
							Optional:            true,
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"nat": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT").AddStringEnumDescription("NAT64", "NAT66").AddDefaultValueDescription("NAT64").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("NAT64", "NAT66"),
							},
						},
						"nat_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"next_hops": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP gateway address").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP Address").String,
										Optional:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"distance": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Administrative distance").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("1").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 255),
										},
									},
									"distance_variable": schema.StringAttribute{
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
			"ipv4_static_gre_routes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure routes pointing to a GRE tunnel").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prefix").String,
							Optional:            true,
						},
						"prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination VPN to resolve the prefix").AddDefaultValueDescription("0").String,
							Optional:            true,
						},
						"interface": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of GRE Interfaces").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"interface_variable": schema.StringAttribute{
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
			"ipv4_static_ipsec_routes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure routes pointing to a IPSEC tunnel").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prefix").String,
							Optional:            true,
						},
						"prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Destination VPN to resolve the prefix").AddDefaultValueDescription("0").String,
							Optional:            true,
						},
						"interface": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of IPSEC Interfaces (Separated by commas)").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"interface_variable": schema.StringAttribute{
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
			"omp_advertise_ipv4_routes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Advertise routes to OMP").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Advertised routes protocol").AddStringEnumDescription("bgp", "ospf", "ospfv3", "connected", "static", "network", "aggregate", "eigrp", "lisp", "isis").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("bgp", "ospf", "ospfv3", "connected", "static", "network", "aggregate", "eigrp", "lisp", "isis"),
							},
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"route_policy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set Route Policy to OMP").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 127),
							},
						},
						"route_policy_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"protocol_sub_type": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"protocol_sub_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"prefixes": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix_entry": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Prefix").String,
										Optional:            true,
									},
									"prefix_entry_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"aggregate_only": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Aggregate Only").AddDefaultValueDescription("false").String,
										Optional:            true,
									},
									"aggregate_only_variable": schema.StringAttribute{
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
			"omp_advertise_ipv6_routes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Advertise routes to OMP").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Advertised routes protocol").AddStringEnumDescription("bgp", "ospf", "connected", "static", "network", "aggregate").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("bgp", "ospf", "connected", "static", "network", "aggregate"),
							},
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"route_policy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 127),
							},
						},
						"route_policy_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"protocol_sub_type": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"protocol_sub_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"prefixes": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix_entry": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Prefix").String,
										Optional:            true,
									},
									"prefix_entry_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"aggregate_only": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Aggregate Only").AddDefaultValueDescription("false").String,
										Optional:            true,
									},
									"aggregate_only_variable": schema.StringAttribute{
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
			"nat64_pools": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set NAT64 v4 pool range").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT64 Pool name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"start_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Starting IP address of NAT pool range").String,
							Optional:            true,
						},
						"start_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"end_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Ending IP address of NAT pool range").String,
							Optional:            true,
						},
						"end_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"overload": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT 64 Overload Option").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"overload_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"leak_from_global": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable Route Leaking from Global VPN to this Service VPN").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"leak_from_global_protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select protocol for route leaking").AddStringEnumDescription("all", "static", "mobile", "connected", "rip", "odr").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("all", "static", "mobile", "connected", "rip", "odr"),
							},
						},
						"leak_to_global": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable Route Leaking from this Service VPN to Global VPN").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"nat_pools": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure NAT Pool entries").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT Pool Name, natpool1..31").AddIntegerRangeDescription(1, 31).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 31),
							},
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"prefix_length": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Ending IP address of NAT Pool Prefix Length").AddIntegerRangeDescription(1, 32).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 32),
							},
						},
						"prefix_length_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"range_start": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Starting IP address of NAT pool range").String,
							Optional:            true,
						},
						"range_start_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"range_end": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Ending IP address of NAT pool range").String,
							Optional:            true,
						},
						"range_end_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"overload": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable port translation(PAT)").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"overload_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Direction of NAT translation").AddStringEnumDescription("inside", "outside").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("inside", "outside"),
							},
						},
						"direction_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"tracker_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Add Object/Object Group Tracker").AddIntegerRangeDescription(1, 1000).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 1000),
							},
						},
						"tracker_id_variable": schema.StringAttribute{
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
			"static_nat_rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure static NAT entries").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"pool_name": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT Pool Name, natpool1..31").String,
							Optional:            true,
						},
						"pool_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source IP address to be translated").String,
							Optional:            true,
						},
						"source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"translate_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Statically translated source IP address").String,
							Optional:            true,
						},
						"translate_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"static_nat_direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Direction of static NAT translation").AddStringEnumDescription("inside", "outside").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("inside", "outside"),
							},
						},
						"static_nat_direction_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"tracker_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Add Object/Object Group Tracker").AddIntegerRangeDescription(1, 1000).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 1000),
							},
						},
						"tracker_id_variable": schema.StringAttribute{
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
			"static_nat_subnet_rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure static NAT Subnet entries").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_ip_subnet": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source IP Subnet to be translated").String,
							Optional:            true,
						},
						"source_ip_subnet_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"translate_ip_subnet": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Statically translated source IP Subnet").String,
							Optional:            true,
						},
						"translate_ip_subnet_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"prefix_length": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Network Prefix Length").AddIntegerRangeDescription(1, 32).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 32),
							},
						},
						"prefix_length_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"static_nat_direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Direction of static NAT translation").AddStringEnumDescription("inside", "outside").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("inside", "outside"),
							},
						},
						"static_nat_direction_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"tracker_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Add Object/Object Group Tracker").AddIntegerRangeDescription(1, 1000).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 1000),
							},
						},
						"tracker_id_variable": schema.StringAttribute{
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
			"port_forward_rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Port Forward entries").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"pool_name": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("NAT Pool Name, natpool1..31").String,
							Optional:            true,
						},
						"pool_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source Port").AddDefaultValueDescription("0").String,
							Optional:            true,
						},
						"source_port_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"translate_port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Translate Port").AddDefaultValueDescription("0").String,
							Optional:            true,
						},
						"translate_port_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source IP address to be translated").String,
							Optional:            true,
						},
						"source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"translate_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Statically translated source IP address").String,
							Optional:            true,
						},
						"translate_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol").AddStringEnumDescription("tcp", "udp").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("tcp", "udp"),
							},
						},
						"protocol_variable": schema.StringAttribute{
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
			"route_global_imports": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable route leaking from Global VPN to this Service VPN").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select a Route Protocol to enable route leaking from Global VPN to this Service VPN").AddStringEnumDescription("static", "connected", "bgp", "ospf").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("static", "connected", "bgp", "ospf"),
							},
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"protocol_sub_type": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").AddDefaultValueDescription("external").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"protocol_sub_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"route_policy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select a Route Policy to enable route leaking from Global VPN to this Service VPN").String,
							Optional:            true,
						},
						"redistributes": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable redistribution of replicated route protocol").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Select a Route Protocol to enable redistribution").AddStringEnumDescription("bgp", "eigrp", "ospf").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("bgp", "eigrp", "ospf"),
										},
									},
									"protocol_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"route_policy": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Select a Route Policy to enable redistribution").String,
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
			"route_vpn_imports": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable route leak from Service VPN to current VPN").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select a Source VPN where route leaks from").AddIntegerRangeDescription(1, 65530).AddDefaultValueDescription("1").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65530),
							},
						},
						"source_vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select a Route Protocol to enable route leaking to current VPN").AddStringEnumDescription("static", "connected", "bgp", "ospf", "eigrp").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("static", "connected", "bgp", "ospf", "eigrp"),
							},
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"protocol_sub_type": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").AddDefaultValueDescription("external").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"protocol_sub_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"route_policy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select a Route Policy to enable route leaking to current VPN").String,
							Optional:            true,
						},
						"route_policy_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"redistributes": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable redistribution of replicated route protocol").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Select a Route Protocol to enable redistribution").AddStringEnumDescription("bgp", "eigrp", "ospf").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("bgp", "eigrp", "ospf"),
										},
									},
									"protocol_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"route_policy": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Select a Route Policy to enable redistribution").String,
										Optional:            true,
									},
									"route_policy_variable": schema.StringAttribute{
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
			"route_global_exports": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable route leaking to Global VPN from this Service VPN").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select a Route Protocol to enable route leaking from this Service VPN to Global VPN").AddStringEnumDescription("static", "connected", "bgp", "eigrp", "ospf").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("static", "connected", "bgp", "eigrp", "ospf"),
							},
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"protocol_sub_type": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").AddDefaultValueDescription("external").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"protocol_sub_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"route_policy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select a Route Policy to enable route leaking from this Service VPN to Global VPN").String,
							Optional:            true,
						},
						"redistributes": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable redistribution of replicated route protocol").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Select a Route Protocol to enable redistribution").AddStringEnumDescription("bgp", "ospf").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("bgp", "ospf"),
										},
									},
									"protocol_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"route_policy": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Select a Route Policy to enable redistribution").String,
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

func (r *CiscoVPNFeatureTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *CiscoVPNFeatureTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CiscoVPN

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
func (r *CiscoVPNFeatureTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CiscoVPN

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
func (r *CiscoVPNFeatureTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CiscoVPN

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
func (r *CiscoVPNFeatureTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CiscoVPN

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
func (r *CiscoVPNFeatureTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
