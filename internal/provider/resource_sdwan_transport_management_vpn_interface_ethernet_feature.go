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
var _ resource.Resource = &TransportManagementVPNInterfaceEthernetProfileParcelResource{}
var _ resource.ResourceWithImportState = &TransportManagementVPNInterfaceEthernetProfileParcelResource{}

func NewTransportManagementVPNInterfaceEthernetProfileParcelResource() resource.Resource {
	return &TransportManagementVPNInterfaceEthernetProfileParcelResource{}
}

type TransportManagementVPNInterfaceEthernetProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *TransportManagementVPNInterfaceEthernetProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transport_management_vpn_interface_ethernet_feature"
}

func (r *TransportManagementVPNInterfaceEthernetProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Transport Management VPN Interface Ethernet Feature.").AddMinimumVersionDescription("20.12.0").String,

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
			"transport_management_vpn_feature_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Transport Management VPN Feature ID").String,
				Optional:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"shutdown_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(3, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`(ATM|ATM-ACR|AppGigabitEthernet|AppNav-Compress|AppNav-UnCompress|Async|BD-VIF|BDI|CEM|CEM-ACR|Cellular|Dialer|Embedded-Service-Engine|Ethernet|Ethernet-Internal|FastEthernet|FiftyGigabitEthernet|FiveGigabitEthernet|FortyGigabitEthernet|FourHundredGigE|GMPLS|GigabitEthernet|Group-Async|HundredGigE|L2LISP|LISP|Loopback|MFR|Multilink|Port-channel|SM|Serial|Service-Engine|TenGigabitEthernet|Tunnel|TwentyFiveGigE|TwentyFiveGigabitEthernet|TwoGigabitEthernet|TwoHundredGigE|Vif|Virtual-PPP|Virtual-Template|VirtualPortGroup|Vlan|Wlan-GigabitEthernet|nat64|nat66|ntp|nve|ospfv3|overlay|pseudowire|ucse|vasileft|vasiright|vmi)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
				},
			},
			"interface_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"interface_description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 200),
				},
			},
			"interface_description_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv4_configuration_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 Configuration Type").AddStringEnumDescription("dynamic", "static").AddDefaultValueDescription("dynamic").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dynamic", "static"),
				},
			},
			"ipv4_dhcp_distance": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("DHCP Distance, Attribute conditional on `ipv4_configuration_type` being equal to `dynamic`").AddIntegerRangeDescription(1, 65536).AddDefaultValueDescription("1").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65536),
				},
			},
			"ipv4_dhcp_distance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `ipv4_configuration_type` being equal to `dynamic`").String,
				Optional:            true,
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP Address, Attribute conditional on `ipv4_configuration_type` being equal to `static`").String,
				Optional:            true,
			},
			"ipv4_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `ipv4_configuration_type` being equal to `static`").String,
				Optional:            true,
			},
			"ipv4_subnet_mask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Subnet Mask, Attribute conditional on `ipv4_configuration_type` being equal to `static`").AddStringEnumDescription("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0"),
				},
			},
			"ipv4_subnet_mask_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `ipv4_configuration_type` being equal to `static`").String,
				Optional:            true,
			},
			"ipv4_secondary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Secondary IpV4 Addresses, Attribute conditional on `ipv4_configuration_type` being equal to `static`").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IpV4 Address").String,
							Optional:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Subnet Mask").AddStringEnumDescription("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0"),
							},
						},
						"subnet_mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv4_dhcp_helper": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of DHCP IPv4 helper addresses (min 1, max 8)").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ipv4_dhcp_helper_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv4_iperf_server": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Iperf server for auto bandwidth detect").String,
				Optional:            true,
			},
			"ipv4_iperf_server_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv4_auto_detect_bandwidth": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface auto detect bandwidth").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ipv4_auto_detect_bandwidth_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv6_configuration_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 Configuration Type").AddStringEnumDescription("dynamic", "static", "none").AddDefaultValueDescription("none").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dynamic", "static", "none"),
				},
			},
			"enable_dhcpv6": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable DHCPv6, Attribute conditional on `ipv6_configuration_type` being equal to `dynamic`").String,
				Optional:            true,
			},
			"ipv6_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 Address Secondary, Attribute conditional on `ipv6_configuration_type` being equal to `static`").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`((^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*(/)(\b([0-9]{1,2}|1[01][0-9]|12[0-8])\b)$))`), ""),
				},
			},
			"ipv6_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `ipv6_configuration_type` being equal to `static`").String,
				Optional:            true,
			},
			"arp_entries": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure ARP entries").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPV4 Address").String,
							Optional:            true,
						},
						"ip_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"mac_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("MAC Address").String,
							Optional:            true,
						},
						"mac_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"duplex": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Duplex mode").AddStringEnumDescription("full", "half", "auto").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("full", "half", "auto"),
				},
			},
			"duplex_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"mac_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MAC Address").String,
				Optional:            true,
			},
			"mac_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ip_mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP MTU for GigabitEthernet main <576..Interface MTU>, GigabitEthernet subinterface <576..9216>, Other Interfaces <576..2000> in bytes").AddIntegerRangeDescription(576, 9216).AddDefaultValueDescription("1500").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(576, 9216),
				},
			},
			"ip_mtu_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"interface_mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface MTU").AddIntegerRangeDescription(1500, 9216).AddDefaultValueDescription("1500").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1500, 9216),
				},
			},
			"interface_mtu_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tcp_mss": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("TCP MSS on SYN packets, in bytes").AddIntegerRangeDescription(500, 1460).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(500, 1460),
				},
			},
			"tcp_mss_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"speed": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set interface speed").AddStringEnumDescription("10", "100", "1000", "2500", "10000").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("10", "100", "1000", "2500", "10000"),
				},
			},
			"speed_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"arp_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Timeout value for dynamically learned ARP entries, <0..2678400> seconds").AddIntegerRangeDescription(0, 2147483).AddDefaultValueDescription("1200").String,
				Optional:            true,
			},
			"arp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"autonegotiate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Link autonegotiation").String,
				Optional:            true,
			},
			"autonegotiate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"media_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Media type").AddStringEnumDescription("auto-select", "rj45", "sfp").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("auto-select", "rj45", "sfp"),
				},
			},
			"media_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"load_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interval for interface load calculation").AddIntegerRangeDescription(30, 600).AddDefaultValueDescription("30").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(30, 600),
				},
			},
			"load_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"icmp_redirect_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ICMP/ICMPv6 Redirect Disable").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"icmp_redirect_disable_variable": schema.StringAttribute{
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
		},
	}
}

func (r *TransportManagementVPNInterfaceEthernetProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *TransportManagementVPNInterfaceEthernetProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TransportManagementVPNInterfaceEthernet

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
func (r *TransportManagementVPNInterfaceEthernetProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state TransportManagementVPNInterfaceEthernet

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
func (r *TransportManagementVPNInterfaceEthernetProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state TransportManagementVPNInterfaceEthernet

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
func (r *TransportManagementVPNInterfaceEthernetProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TransportManagementVPNInterfaceEthernet

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
func (r *TransportManagementVPNInterfaceEthernetProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 2
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "transport_management_vpn_interface_ethernet_feature_id" + ",feature_profile_id" + ",transport_management_vpn_feature_id"
	if len(parts) != (count + 1) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: %s. Got: %q", pattern, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("feature_profile_id"), parts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("transport_management_vpn_feature_id"), parts[2])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
