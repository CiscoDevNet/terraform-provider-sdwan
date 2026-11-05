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
var _ resource.Resource = &TransportWANVPNInterfaceDSLIPoEProfileParcelResource{}
var _ resource.ResourceWithImportState = &TransportWANVPNInterfaceDSLIPoEProfileParcelResource{}

func NewTransportWANVPNInterfaceDSLIPoEProfileParcelResource() resource.Resource {
	return &TransportWANVPNInterfaceDSLIPoEProfileParcelResource{}
}

type TransportWANVPNInterfaceDSLIPoEProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *TransportWANVPNInterfaceDSLIPoEProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transport_wan_vpn_interface_dsl_ipoe_feature"
}

func (r *TransportWANVPNInterfaceDSLIPoEProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Transport WAN VPN Interface DSL IPoE Feature.").AddMinimumVersionDescription("20.15.0").String,

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
			"transport_wan_vpn_feature_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Transport WAN VPN Feature ID").String,
				Required:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative state").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"shutdown_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"service_provider": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Service Provider Name").String,
				Optional:            true,
			},
			"service_provider_variable": schema.StringAttribute{
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
			"vdsl_slot": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set module slot/subslot/port number").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(5),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[0-3]/[0-3]/[0-1]$`), ""),
				},
			},
			"vdsl_slot_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"vdsl_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set VDSL operating mode").AddStringEnumDescription("ADSL1", "ADSL2", "ADSL2+", "ANSI", "VDSL2").AddDefaultValueDescription("auto").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ADSL1", "ADSL2", "ADSL2+", "ANSI", "VDSL2"),
				},
			},
			"vdsl_sra": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Seamless rate adaption").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"vdsl_sra_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"nat_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network Address Translation on this interface").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"nat_enabled_variable": schema.StringAttribute{
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
			"interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ethernet Interface/Sub Interface Name including sub interface number").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(3, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`^(ATM|ATM-ACR|AppGigabitEthernet|AppNav-Compress|AppNav-UnCompress|Async|BD-VIF|BDI|CEM|CEM-ACR|Cellular|Dialer|Embedded-Service-Engine|Ethernet|Ethernet-Internal|FastEthernet|FiftyGigabitEthernet|FiveGigabitEthernet|FortyGigabitEthernet|FourHundredGigE|GMPLS|GigabitEthernet|Group-Async|HundredGigE|L2LISP|LISP|Loopback|MFR|Multilink|Port-channel|SM|Serial|Service-Engine|TenGigabitEthernet|Tunnel|TwentyFiveGigE|TwentyFiveGigabitEthernet|TwoGigabitEthernet|TwoHundredGigE|Vif|Virtual-PPP|Virtual-Template|VirtualPortGroup|Vlan|Wlan-GigabitEthernet|nat64|nat66|ntp|nve|ospfv3|overlay|pseudowire|ucse|vasileft|vasiright|vmi)\s*(?:\d+(?:/\d+){0,2}(?::\d+)?(?:\.\d+)?|\d+(?:\.\d+)?)$`), ""),
				},
			},
			"interface_name_variable": schema.StringAttribute{
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
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Encapsulation VLAN id").AddIntegerRangeDescription(1, 4094).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4094),
				},
			},
			"vlan_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv4_address_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 address configuration type").AddStringEnumDescription("dynamic", "static").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dynamic", "static"),
				},
			},
			"ipv4_dhcp_distance": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("DHCP Distance, Attribute conditional on `ipv4_address_type` equal to `dynamic`").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("1").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"ipv4_dhcp_distance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `ipv4_address_type` equal to `dynamic`").String,
				Optional:            true,
			},
			"ipv4_dhcp_helper": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of DHCP IPv4 helper addresses, Attribute conditional on `ipv4_address_type` equal to `dynamic`").String,
				Optional:            true,
			},
			"ipv4_dhcp_helper_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `ipv4_address_type` equal to `dynamic`").String,
				Optional:            true,
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP Address, Attribute conditional on `ipv4_address_type` equal to `static`").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)\.){3}(25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)$`), ""),
				},
			},
			"ipv4_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `ipv4_address_type` equal to `static`").String,
				Optional:            true,
			},
			"ipv4_subnet_mask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Subnet Mask, Attribute conditional on `ipv4_address_type` equal to `static`").AddStringEnumDescription("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.248.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.248.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0"),
				},
			},
			"ipv4_subnet_mask_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `ipv4_address_type` equal to `static`").String,
				Optional:            true,
			},
			"tunnel_interface_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Tunnel Interface on/off").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_per_tunnel_qos": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Per-tunnel Qos").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_per_tunnel_qos_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_qos_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set tunnel QoS mode").AddStringEnumDescription("hub", "spoke").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("hub", "spoke"),
				},
			},
			"tunnel_qos_mode_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_bandwidth_percent": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Tunnels Bandwidth Percent").AddIntegerRangeDescription(1, 100).AddDefaultValueDescription("50").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
				},
			},
			"tunnel_bandwidth_percent_variable": schema.StringAttribute{
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
			"tunnel_interface_color_description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
				},
			},
			"tunnel_interface_color_description_variable": schema.StringAttribute{
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
			"tunnel_interface_last_resort_circuit": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set TLOC as last resort").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_last_resort_circuit_variable": schema.StringAttribute{
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
			"tunnel_interface_groups": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of groups").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.RegexMatches(regexp.MustCompile(`^(([1-9]\d{0,8}?)|([1-3]\d{9}?)|((4[0-1]\d{8})|(42[0-8]\d{7})|(429[0-3]\d{6})|(4294[0-8]\d{5})|(42949[0-5]\d{4})|(429496[0-6]\d{3}|(4294967[0-1]\d{2})|(42949672[0-8]\d)|(429496729[0-5]))))(,(([1-9]\d{0,8}?)|([1-3]\d{9}?)|((4[0-1]\d{8})|(42[0-8]\d{7})|(429[0-3]\d{6})|(4294[0-8]\d{5})|(42949[0-5]\d{4})|(429496[0-6]\d{3}|(4294967[0-1]\d{2})|(42949672[0-8]\d)|(429496729[0-5])))))*$`), ""),
				},
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
			"tunnel_interface_max_control_connections": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the maximum number of control connections for this TLOC").AddIntegerRangeDescription(0, 8).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtMost(8),
				},
			},
			"tunnel_interface_max_control_connections_variable": schema.StringAttribute{
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
			"tunnel_interface_exclude_controller_group_list": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Exclude the following controller groups defined in this list").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
					stringvalidator.RegexMatches(regexp.MustCompile(`^(0|[1-9]\d?|100)(,(0|[1-9]\d?|100))*$`), ""),
				},
			},
			"tunnel_interface_exclude_controller_group_list_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_vmanage_connection_preference": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set interface preference for control connection to vManage <0..8>").AddIntegerRangeDescription(0, 8).AddDefaultValueDescription("5").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtMost(8),
				},
			},
			"tunnel_interface_vmanage_connection_preference_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_full_port_hop": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable port hopping on the tunnel interface").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_full_port_hop_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_port_hop": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The port hop functionality is deprecated for devices 17.18 and higher. Use the full-port-hop field instead").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"tunnel_interface_port_hop_variable": schema.StringAttribute{
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
			"tunnel_interface_network_broadcast": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Accept and respond to network-prefix-directed broadcasts)").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_network_broadcast_variable": schema.StringAttribute{
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
			"tunnel_interface_allow_stun": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny STUN").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_allow_stun_variable": schema.StringAttribute{
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
			"tunnel_interface_allow_snmp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny SNMP").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_allow_snmp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_ipsec_encapsulation": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Encapsulation").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"tunnel_interface_ipsec_preference": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set preference for TLOC").AddIntegerRangeDescription(0, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtMost(4294967295),
				},
			},
			"tunnel_interface_ipsec_preference_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_ipsec_weight": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set weight for TLOC").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("1").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"tunnel_interface_ipsec_weight_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_gre_encapsulation": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Encapsulation").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_gre_preference": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set preference for TLOC").AddIntegerRangeDescription(0, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtMost(4294967295),
				},
			},
			"tunnel_interface_gre_preference_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_gre_weight": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set weight for TLOC").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("1").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"tunnel_interface_gre_weight_variable": schema.StringAttribute{
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
			"multi_region_fabric_enable_core_region": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Core Region").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"multi_region_fabric_core_region": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Core Region").AddStringEnumDescription("core-shared", "core").AddDefaultValueDescription("core-shared").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("core-shared", "core"),
				},
			},
			"multi_region_fabric_enable_secondary_region": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Secondary Region").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"multi_region_fabric_secondary_region": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable secondary region").AddStringEnumDescription("secondary-shared", "secondary-only").AddDefaultValueDescription("secondary-shared").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("secondary-shared", "secondary-only"),
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
			"ip_mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP MTU for GigabitEthernet main <576..Interface MTU>, GigabitEthernet subinterface <576..9216>, Other Interfaces <576..2000> in bytes").AddIntegerRangeDescription(576, 2000).AddDefaultValueDescription("1500").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(576, 2000),
				},
			},
			"ip_mtu_variable": schema.StringAttribute{
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
			},
			"tloc_extension_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"adaptive_qos": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Adaptive QoS").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"adaptive_qos_period": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Adapt Period(Minutes)").AddIntegerRangeDescription(1, 720).AddDefaultValueDescription("15").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 720),
				},
			},
			"adaptive_qos_period_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"shaping_rate_upstream": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Shaping Rate Upstream").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"min_shaping_rate_upstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Upstream min bandwidth limit (kbps)").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"min_shaping_rate_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"max_shaping_rate_upstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Upstream max bandwidth limit (kbps)").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"max_shaping_rate_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"default_shaping_rate_upstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Adaptive QoS default upstream bandwidth (kbps)").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"default_shaping_rate_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"shaping_rate_downstream": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Shaping Rate Downstream").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"min_shaping_rate_downstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Downstream min bandwidth limit (kbps)").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"min_shaping_rate_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"max_shaping_rate_downstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Downstream max bandwidth limit (kbps)").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"max_shaping_rate_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"default_shaping_rate_downstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Adaptive QoS default downstream bandwidth (kbps)").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"default_shaping_rate_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"shaping_rate": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Shaping Rate (Kbps)").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"shaping_rate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"acl_ipv4_egress_feature_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
				},
			},
			"acl_ipv4_egress_feature_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"acl_ipv4_ingress_feature_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
				},
			},
			"acl_ipv4_ingress_feature_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"acl_ipv6_egress_feature_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
				},
			},
			"acl_ipv6_egress_feature_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"acl_ipv6_ingress_feature_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
				},
			},
			"acl_ipv6_ingress_feature_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
		},
	}
}

func (r *TransportWANVPNInterfaceDSLIPoEProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *TransportWANVPNInterfaceDSLIPoEProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TransportWANVPNInterfaceDSLIPoE

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
func (r *TransportWANVPNInterfaceDSLIPoEProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state TransportWANVPNInterfaceDSLIPoE

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
func (r *TransportWANVPNInterfaceDSLIPoEProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state TransportWANVPNInterfaceDSLIPoE

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
func (r *TransportWANVPNInterfaceDSLIPoEProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TransportWANVPNInterfaceDSLIPoE

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
func (r *TransportWANVPNInterfaceDSLIPoEProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 2
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "transport_wan_vpn_interface_dsl_ipoe_feature_id" + ",feature_profile_id" + ",transport_wan_vpn_feature_id"
	if len(parts) != (count + 1) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: %s. Got: %q", pattern, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("feature_profile_id"), parts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("transport_wan_vpn_feature_id"), parts[2])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
