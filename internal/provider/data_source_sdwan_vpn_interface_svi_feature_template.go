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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &VPNInterfaceSVIFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &VPNInterfaceSVIFeatureTemplateDataSource{}
)

func NewVPNInterfaceSVIFeatureTemplateDataSource() datasource.DataSource {
	return &VPNInterfaceSVIFeatureTemplateDataSource{}
}

type VPNInterfaceSVIFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *VPNInterfaceSVIFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_interface_svi_feature_template"
}

func (d *VPNInterfaceSVIFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the VPN Interface SVI feature template.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the feature template",
				Optional:            true,
				Computed:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the feature template",
				Computed:            true,
			},
			"template_type": schema.StringAttribute{
				MarkdownDescription: "The template type",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the feature template",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the feature template",
				Computed:            true,
			},
			"device_types": schema.SetAttribute{
				MarkdownDescription: "List of supported device types",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"if_name": schema.StringAttribute{
				MarkdownDescription: "Interface name: VLAN 1 - VLAN 4094 when present",
				Computed:            true,
			},
			"if_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"interface_description": schema.StringAttribute{
				MarkdownDescription: "Interface description",
				Computed:            true,
			},
			"interface_description_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: "Assign IPv4 address",
				Computed:            true,
			},
			"ipv4_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_secondary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Assign secondary IP addresses",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ipv4_address": schema.StringAttribute{
							MarkdownDescription: "IP Address",
							Computed:            true,
						},
						"ipv4_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"ipv6_address": schema.StringAttribute{
				MarkdownDescription: "Assign IPv6 address",
				Computed:            true,
			},
			"ipv6_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_dhcp_client": schema.BoolAttribute{
				MarkdownDescription: "Enable DHCPv6",
				Computed:            true,
			},
			"ipv6_dhcp_client_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_dhcp_distance": schema.Int64Attribute{
				MarkdownDescription: "Set administrative distance for DHCP default route",
				Computed:            true,
			},
			"ipv6_dhcp_distance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_dhcp_rapid_commit": schema.BoolAttribute{
				MarkdownDescription: "Enable DHCPv6 rapid commit",
				Computed:            true,
			},
			"ipv6_dhcp_rapid_commit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_secondary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Assign secondary IPv6 addresses",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ipv6_address": schema.StringAttribute{
							MarkdownDescription: "IPv6 Address",
							Computed:            true,
						},
						"ipv6_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_dhcp_helper": schema.SetAttribute{
				MarkdownDescription: "List of DHCP helper addresses",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ipv4_dhcp_helper_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_dhcp_helpers": schema.ListNestedAttribute{
				MarkdownDescription: "DHCPv6 Helper",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "DHCPv6 Helper address",
							Computed:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "DHCPv6 Helper VPN",
							Computed:            true,
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"ip_directed_broadcast": schema.BoolAttribute{
				MarkdownDescription: "IP Directed-Broadcast",
				Computed:            true,
			},
			"ip_directed_broadcast_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: "Interface MTU <1500..9216> in bytes",
				Computed:            true,
			},
			"mtu_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ip_mtu": schema.Int64Attribute{
				MarkdownDescription: "IP MTU <576..Interface MTU>, in bytes",
				Computed:            true,
			},
			"ip_mtu_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tcp_mss_adjust": schema.Int64Attribute{
				MarkdownDescription: "TCP MSS on SYN packets, in bytes",
				Computed:            true,
			},
			"tcp_mss_adjust_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "Administrative state",
				Computed:            true,
			},
			"shutdown_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"arp_timeout": schema.Int64Attribute{
				MarkdownDescription: "Timeout value for dynamically learned ARP entries, <0..2678400> seconds",
				Computed:            true,
			},
			"arp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_access_lists": schema.ListNestedAttribute{
				MarkdownDescription: "Apply ACL",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"direction": schema.StringAttribute{
							MarkdownDescription: "Direction",
							Computed:            true,
						},
						"acl_name": schema.StringAttribute{
							MarkdownDescription: "Name of access list",
							Computed:            true,
						},
						"acl_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"ipv6_access_lists": schema.ListNestedAttribute{
				MarkdownDescription: "Apply ACL",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"direction": schema.StringAttribute{
							MarkdownDescription: "Direction",
							Computed:            true,
						},
						"acl_name": schema.StringAttribute{
							MarkdownDescription: "Name of access list",
							Computed:            true,
						},
						"acl_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"policers": schema.ListNestedAttribute{
				MarkdownDescription: "Enable policer",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"direction": schema.StringAttribute{
							MarkdownDescription: "Direction",
							Computed:            true,
						},
						"policer_name": schema.StringAttribute{
							MarkdownDescription: "Name of policer",
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"static_arp_entries": schema.ListNestedAttribute{
				MarkdownDescription: "Configure static ARP entries",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ipv4_address": schema.StringAttribute{
							MarkdownDescription: "IP Address",
							Computed:            true,
						},
						"ipv4_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"mac_address": schema.StringAttribute{
							MarkdownDescription: "MAC address",
							Computed:            true,
						},
						"mac_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"ipv4_vrrps": schema.ListNestedAttribute{
				MarkdownDescription: "Enable VRRP",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.Int64Attribute{
							MarkdownDescription: "Group ID",
							Computed:            true,
						},
						"group_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: "Set priority",
							Computed:            true,
						},
						"priority_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"timer": schema.Int64Attribute{
							MarkdownDescription: "Timer interval for successive advertisements, in milliseconds",
							Computed:            true,
						},
						"timer_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"track_omp": schema.BoolAttribute{
							MarkdownDescription: "Track OMP status",
							Computed:            true,
						},
						"track_omp_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"track_prefix_list": schema.StringAttribute{
							MarkdownDescription: "Track Prefix List",
							Computed:            true,
						},
						"track_prefix_list_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ipv4_address": schema.StringAttribute{
							MarkdownDescription: "Assign IP Address",
							Computed:            true,
						},
						"ipv4_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ipv4_secondary_addresses": schema.ListNestedAttribute{
							MarkdownDescription: "VRRP Secondary IP address",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ipv4_address": schema.StringAttribute{
										MarkdownDescription: "VRRP Secondary IP address",
										Computed:            true,
									},
									"ipv4_address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Computed:            true,
									},
								},
							},
						},
						"tloc_preference_change": schema.BoolAttribute{
							MarkdownDescription: "change TLOC preference",
							Computed:            true,
						},
						"tloc_preference_change_value": schema.Int64Attribute{
							MarkdownDescription: "Set tloc preference change value",
							Computed:            true,
						},
						"tloc_preference_change_value_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tracking_objects": schema.ListNestedAttribute{
							MarkdownDescription: "tracking object for VRRP configuration",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.Int64Attribute{
										MarkdownDescription: "Tracker ID",
										Computed:            true,
									},
									"name_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"track_action": schema.StringAttribute{
										MarkdownDescription: "Track Action",
										Computed:            true,
									},
									"track_action_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"decrement_value": schema.Int64Attribute{
										MarkdownDescription: "Decrement Value for VRRP priority",
										Computed:            true,
									},
									"decrement_value_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Computed:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"ipv6_vrrps": schema.ListNestedAttribute{
				MarkdownDescription: "Enable VRRP",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"group_id": schema.Int64Attribute{
							MarkdownDescription: "Group ID",
							Computed:            true,
						},
						"group_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: "Set priority",
							Computed:            true,
						},
						"priority_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"timer": schema.Int64Attribute{
							MarkdownDescription: "Timer interval for successive advertisements, in milliseconds",
							Computed:            true,
						},
						"timer_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"track_omp": schema.BoolAttribute{
							MarkdownDescription: "Track OMP status",
							Computed:            true,
						},
						"track_omp_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"track_prefix_list": schema.StringAttribute{
							MarkdownDescription: "Track Prefix List",
							Computed:            true,
						},
						"track_prefix_list_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ipv6_addresses": schema.ListNestedAttribute{
							MarkdownDescription: "IPv6 VRRP",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"link_local_address": schema.StringAttribute{
										MarkdownDescription: "Use link-local IPv6 Address",
										Computed:            true,
									},
									"link_local_address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"prefix": schema.StringAttribute{
										MarkdownDescription: "Assign Global IPv6 Prefix",
										Computed:            true,
									},
									"prefix_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Computed:            true,
									},
								},
							},
						},
						"ipv6_secondary_addresses": schema.ListNestedAttribute{
							MarkdownDescription: "IPv6 Secondary IP address",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"prefix": schema.StringAttribute{
										MarkdownDescription: "IPv6 Secondary IP address",
										Computed:            true,
									},
									"prefix_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Computed:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *VPNInterfaceSVIFeatureTemplateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *VPNInterfaceSVIFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *VPNInterfaceSVIFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VPNInterfaceSVI

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	if config.Id.IsNull() && !config.Name.IsNull() {
		res, err := d.client.Get("/template/feature")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve existing templates, got error: %s", err))
			return
		}
		if value := res.Get("data"); len(value.Array()) > 0 {
			value.ForEach(func(k, v gjson.Result) bool {
				if config.Name.ValueString() == v.Get("templateName").String() {
					config.Id = types.StringValue(v.Get("templateId").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found feature template with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
					return false
				}
				return true
			})
		}
		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find feature template with name: %s", config.Name.ValueString()))
			return
		}
	}

	res, err := d.client.Get("/template/feature/object/" + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
