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
var _ resource.Resource = &VPNInterfaceCellularFeatureTemplateResource{}
var _ resource.ResourceWithImportState = &VPNInterfaceCellularFeatureTemplateResource{}

func NewVPNInterfaceCellularFeatureTemplateResource() resource.Resource {
	return &VPNInterfaceCellularFeatureTemplateResource{}
}

type VPNInterfaceCellularFeatureTemplateResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *VPNInterfaceCellularFeatureTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_interface_cellular_feature_template"
}

func (r *VPNInterfaceCellularFeatureTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a VPN Interface Cellular feature template.").AddMinimumVersionDescription("15.0.0").String,

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
			"cellular_interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Cellular Interface Name <0-1>").String,
				Optional:            true,
			},
			"cellular_interface_name_variable": schema.StringAttribute{
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
			"ipv6_access_lists": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Apply IPv6 access list").String,
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
			"ipv4_dhcp_helper": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of DHCP server addresses").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ipv4_dhcp_helper_variable": schema.StringAttribute{
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
			"nat": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network Address Translation on this interface").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"nat_refresh_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set NAT refresh mode").AddStringEnumDescription("outbound", "bi-directional").AddDefaultValueDescription("outbound").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("outbound", "bi-directional"),
				},
			},
			"nat_refresh_mode_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"nat_udp_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set NAT UDP session timeout, in minutes").AddIntegerRangeDescription(1, 8947).AddDefaultValueDescription("1").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 8947),
				},
			},
			"nat_udp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"nat_tcp_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set NAT TCP session timeout, in minutes").AddIntegerRangeDescription(1, 8947).AddDefaultValueDescription("60").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 8947),
				},
			},
			"nat_tcp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"nat_block_icmp_error": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Block inbound ICMP error messages").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"nat_block_icmp_error_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"nat_response_to_ping": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Respond to ping requests to NAT interface ip address from the public side").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"nat_response_to_ping_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"nat_port_forwards": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set port-forwarding rules for NAT on this interface").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port_start_range": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Starting port of port range").AddIntegerRangeDescription(0, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"port_end_range": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Ending port of port range").AddIntegerRangeDescription(0, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Layer 4 protocol to apply port forwarding to").AddStringEnumDescription("tcp", "udp").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("tcp", "udp"),
							},
						},
						"private_vpn": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("VPN in which private IP Address resides").AddIntegerRangeDescription(0, 65535).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"private_vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"private_ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Private IP Address to translate to").String,
							Optional:            true,
						},
						"private_ip_address_variable": schema.StringAttribute{
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
			"enable_core_region": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable core region").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"enable_core_region_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"core_region": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable core region").AddStringEnumDescription("core", "core-shared").AddDefaultValueDescription("core").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("core", "core-shared"),
				},
			},
			"core_region_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"secondary_region": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable secondary region").AddStringEnumDescription("off", "secondary-only", "secondary-shared").AddDefaultValueDescription("off").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("off", "secondary-only", "secondary-shared"),
				},
			},
			"secondary_region_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_encapsulations": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Encapsulation for TLOC").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"encapsulation": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Encapsulation").AddStringEnumDescription("gre", "ipsec").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("gre", "ipsec"),
							},
						},
						"preference": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set preference for TLOC").AddIntegerRangeDescription(0, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"preference_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"weight": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set weight for TLOC").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("1").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
						"weight_variable": schema.StringAttribute{
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
			"tunnel_interface_groups": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of groups").String,
				ElementType:         types.Int64Type,
				Optional:            true,
			},
			"tunnel_interface_groups_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_border": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set TLOC as border TLOC").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_border_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"per_tunnel_qos": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Per-tunnel Qos").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"per_tunnel_qos_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"per_tunnel_qos_aggregator": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Per-tunnel QoS Aggregator").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"per_tunnel_qos_aggregator_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_qos_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set tunnel QoS mode").AddStringEnumDescription("spoke").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("spoke"),
				},
			},
			"tunnel_qos_mode_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_color": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set color for TLOC").AddStringEnumDescription("default", "mpls", "metro-ethernet", "biz-internet", "public-internet", "lte", "3g", "red", "green", "blue", "gold", "silver", "bronze", "custom1", "custom2", "custom3", "private1", "private2", "private3", "private4", "private5", "private6").AddDefaultValueDescription("default").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "mpls", "metro-ethernet", "biz-internet", "public-internet", "lte", "3g", "red", "green", "blue", "gold", "silver", "bronze", "custom1", "custom2", "custom3", "private1", "private2", "private3", "private4", "private5", "private6"),
				},
			},
			"tunnel_interface_color_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_last_resort_circuit": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set TLOC as last resort").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_last_resort_circuit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_low_bandwidth_link": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the interface as a low-bandwidth circuit").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_low_bandwidth_link_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_tunnel_tcp_mss": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Tunnel TCP MSS on SYN packets, in bytes").AddIntegerRangeDescription(500, 1460).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(500, 1460),
				},
			},
			"tunnel_interface_tunnel_tcp_mss_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_clear_dont_fragment": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable clear dont fragment (Currently Only SDWAN Tunnel Interface)").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_clear_dont_fragment_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_network_broadcast": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Accept and respond to network-prefix-directed broadcasts)").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_network_broadcast_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_max_control_connections": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the maximum number of control connections for this TLOC").AddIntegerRangeDescription(0, 8).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 8),
				},
			},
			"tunnel_interface_max_control_connections_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_control_connections": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow Control Connection").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"tunnel_interface_control_connections_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_vbond_as_stun_server": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Put this wan interface in STUN mode only").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_vbond_as_stun_server_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_exclude_controller_group_list": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Exclude the following controller groups defined in this list").String,
				ElementType:         types.Int64Type,
				Optional:            true,
			},
			"tunnel_interface_exclude_controller_group_list_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_vmanage_connection_preference": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set interface preference for control connection to vManage <0..8>").AddIntegerRangeDescription(0, 8).AddDefaultValueDescription("5").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 8),
				},
			},
			"tunnel_interface_vmanage_connection_preference_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_port_hop": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Disallow port hopping on the tunnel interface").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"tunnel_interface_port_hop_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_color_restrict": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Restrict this TLOC behavior").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_color_restrict_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_carrier": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set carrier for TLOC").AddStringEnumDescription("default", "carrier1", "carrier2", "carrier3", "carrier4", "carrier5", "carrier6", "carrier7", "carrier8").AddDefaultValueDescription("default").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "carrier1", "carrier2", "carrier3", "carrier4", "carrier5", "carrier6", "carrier7", "carrier8"),
				},
			},
			"tunnel_interface_carrier_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_nat_refresh_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set time period of nat refresh packets <1...60> seconds").AddIntegerRangeDescription(1, 60).AddDefaultValueDescription("5").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 60),
				},
			},
			"tunnel_interface_nat_refresh_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_hello_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set time period of control hello packets <100..600000> milli seconds").AddIntegerRangeDescription(100, 600000).AddDefaultValueDescription("1000").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(100, 600000),
				},
			},
			"tunnel_interface_hello_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_hello_tolerance": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set tolerance of control hello packets <12..6000> seconds").AddIntegerRangeDescription(12, 6000).AddDefaultValueDescription("12").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(12, 6000),
				},
			},
			"tunnel_interface_hello_tolerance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_bind_loopback_tunnel": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Bind loopback tunnel interface to a physical interface").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"tunnel_interface_bind_loopback_tunnel_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_allow_all": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow all traffic. Overrides all other allow-service options if allow-service all is set").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_allow_all_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_allow_bgp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/deny BGP").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_allow_bgp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_allow_dhcp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny DHCP").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"tunnel_interface_allow_dhcp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_allow_dns": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny DNS").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"tunnel_interface_allow_dns_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_allow_icmp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny ICMP").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"tunnel_interface_allow_icmp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_allow_ssh": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny SSH").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_allow_ssh_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_allow_ntp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny NTP").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_allow_ntp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_allow_netconf": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny NETCONF").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_allow_netconf_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_allow_ospf": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny OSPF").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_allow_ospf_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_allow_stun": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny STUN").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_allow_stun_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_allow_snmp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny SNMP").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_allow_snmp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_allow_https": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny Https").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"tunnel_interface_allow_https_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"clear_dont_fragment_bit": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Clear don't fragment bit").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"clear_dont_fragment_bit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"pmtu_discovery": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Path MTU Discovery").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"pmtu_discovery_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ip_mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface MTU <68...1500>, in bytes").AddIntegerRangeDescription(68, 1500).AddDefaultValueDescription("1428").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(68, 1500),
				},
			},
			"ip_mtu_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"static_ingress_qos": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static ingress QoS for the port").AddIntegerRangeDescription(0, 7).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 7),
				},
			},
			"static_ingress_qos_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tcp_mss": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("TCP MSS on SYN packets, in bytes").AddIntegerRangeDescription(552, 1960).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(552, 1960),
				},
			},
			"tcp_mss_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tloc_extension": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Extends a local TLOC to a remote node only for vpn 0").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"tloc_extension_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ip_directed_broadcast": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP Directed-Broadcast").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ip_directed_broadcast_variable": schema.StringAttribute{
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
			"autonegotiate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Link autonegotiation").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"autonegotiate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"qos_adaptive_period": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Periodic timer for adaptive QoS in minutes").AddIntegerRangeDescription(1, 720).AddDefaultValueDescription("15").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 720),
				},
			},
			"qos_adaptive_period_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"qos_adaptive_bandwidth_downstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Adaptive QoS default downstream bandwidth").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"qos_adaptive_bandwidth_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"qos_adaptive_min_downstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Downstream min bandwidth limit").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"qos_adaptive_min_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"qos_adaptive_max_downstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Downstream max bandwidth limit").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"qos_adaptive_max_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"qos_adaptive_bandwidth_upstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Adaptive QoS default upstream bandwidth").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"qos_adaptive_bandwidth_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"qos_adaptive_min_upstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Upstream min bandwidth limit").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"qos_adaptive_min_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"qos_adaptive_max_upstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Upstream max bandwidth limit").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"qos_adaptive_max_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"shaping_rate": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("1ge  interfaces: [0..1000000]kbps; 10ge interfaces: [0..10000000]kbps").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"shaping_rate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"qos_map": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of QoS map").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
				},
			},
			"qos_map_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"qos_map_vpn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of VPN QoS map").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
				},
			},
			"qos_map_vpn_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"bandwidth_upstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface upstream bandwidth capacity, in kbps").AddIntegerRangeDescription(1, 2147483647).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 2147483647),
				},
			},
			"bandwidth_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"bandwidth_downstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface downstream bandwidth capacity, in kbps").AddIntegerRangeDescription(1, 2147483647).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 2147483647),
				},
			},
			"bandwidth_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"write_rule": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of rewrite rule").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
				},
			},
			"write_rule_variable": schema.StringAttribute{
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
			"static_arps": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure static ARP entries").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP Address").String,
							Optional:            true,
						},
						"ip_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("MAC address").String,
							Optional:            true,
						},
						"mac_variable": schema.StringAttribute{
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
		},
	}
}

func (r *VPNInterfaceCellularFeatureTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *VPNInterfaceCellularFeatureTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VPNInterfaceCellular

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
func (r *VPNInterfaceCellularFeatureTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VPNInterfaceCellular

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
func (r *VPNInterfaceCellularFeatureTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state VPNInterfaceCellular

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
func (r *VPNInterfaceCellularFeatureTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VPNInterfaceCellular

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
func (r *VPNInterfaceCellularFeatureTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
