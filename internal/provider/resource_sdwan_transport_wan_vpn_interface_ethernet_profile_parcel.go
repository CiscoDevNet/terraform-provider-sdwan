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
var _ resource.Resource = &TransportWANVPNInterfaceEthernetProfileParcelResource{}
var _ resource.ResourceWithImportState = &TransportWANVPNInterfaceEthernetProfileParcelResource{}

func NewTransportWANVPNInterfaceEthernetProfileParcelResource() resource.Resource {
	return &TransportWANVPNInterfaceEthernetProfileParcelResource{}
}

type TransportWANVPNInterfaceEthernetProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *TransportWANVPNInterfaceEthernetProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transport_wan_vpn_interface_ethernet_profile_parcel"
}

func (r *TransportWANVPNInterfaceEthernetProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Transport WAN VPN Interface Ethernet profile parcel.").AddMinimumVersionDescription("20.12.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the profile parcel",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the profile parcel",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the profile parcel",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the profile parcel",
				Optional:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Feature Profile ID").String,
				Required:            true,
			},
			"transport_wan_vpn_profile_parcel_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Transport WAN VPN Profile Parcel ID").String,
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
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 Configuration Type").AddStringEnumDescription("dynamic", "static").String,
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
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP Address, Attribute conditional on `ipv4_configuration_type` being equal to `static`").String,
				Optional:            true,
			},
			"ipv4_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
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
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
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
			"ipv6_configuration_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 Configuration Type").AddStringEnumDescription("dynamic", "static", "none").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dynamic", "static", "none"),
				},
			},
			"enable_dhcpv6": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable DHCPv6, Attribute conditional on `ipv6_configuration_type` being equal to `dynamic`").String,
				Optional:            true,
			},
			"ipv6_dhcp_secondary_address": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("secondary IPv6 addresses, Attribute conditional on `ipv6_configuration_type` being equal to `dynamic`").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 Address Secondary").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`((^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*(/)(\b([0-9]{1,2}|1[01][0-9]|12[0-8])\b)$))`), ""),
							},
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"ipv6_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 Address Secondary, Attribute conditional on `ipv6_configuration_type` being equal to `static`").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`((^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*(/)(\b([0-9]{1,2}|1[01][0-9]|12[0-8])\b)$))`), ""),
				},
			},
			"ipv6_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ipv6_secondary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static secondary IPv6 addresses, Attribute conditional on `ipv6_configuration_type` being equal to `static`").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv6 Address Secondary").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`((^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*(/)(\b([0-9]{1,2}|1[01][0-9]|12[0-8])\b)$))`), ""),
							},
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"iperf_server": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Iperf server for auto bandwidth detect").String,
				Optional:            true,
			},
			"iperf_server_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"block_non_source_ip": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Block packets originating from IP address that is not from this source").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"block_non_source_ip_variable": schema.StringAttribute{
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
			"auto_detect_bandwidth": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface auto detect bandwidth").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"auto_detect_bandwidth_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Tunnel Interface on/off").AddDefaultValueDescription("false").String,
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
				MarkdownDescription: helpers.NewAttributeDescription("Set color for TLOC").AddStringEnumDescription("default", "mpls", "metro-ethernet", "biz-internet", "public-internet", "lte", "3g", "red", "green", "blue", "gold", "silver", "bronze", "custom1", "custom2", "custom3", "private1", "private2", "private3", "private4", "private5", "private6").AddDefaultValueDescription("mpls").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "mpls", "metro-ethernet", "biz-internet", "public-internet", "lte", "3g", "red", "green", "blue", "gold", "silver", "bronze", "custom1", "custom2", "custom3", "private1", "private2", "private3", "private4", "private5", "private6"),
				},
			},
			"tunnel_interface_color_variable": schema.StringAttribute{
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
			"tunnel_interface_gre_tunnel_destination_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("GRE tunnel destination IP").String,
				Optional:            true,
			},
			"tunnel_interface_gre_tunnel_destination_ip_variable": schema.StringAttribute{
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
			"tunnel_interface_groups": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of groups").AddIntegerRangeDescription(1, 4294967295).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
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
				MarkdownDescription: helpers.NewAttributeDescription("Maximum Control Connections").AddIntegerRangeDescription(0, 100).String,
				Optional:            true,
			},
			"tunnel_interface_max_control_connections_variable": schema.StringAttribute{
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
			"tunnel_interface_vbond_as_stun_server": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Put this wan interface in STUN mode only").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_vbond_as_stun_server_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_exclude_controller_group_list": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Exclude the following controller groups defined in this list.").String,
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
			"tunnel_interface_cts_sgt_propagation": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("CTS SGT Propagation configuration").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_cts_sgt_propagation_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_network_broadcast": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Accept and respond to network-prefix-directed broadcasts").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_network_broadcast_variable": schema.StringAttribute{
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
			"tunnel_interface_allow_ntp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny NTP").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"tunnel_interface_allow_ntp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_allow_ssh": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny SSH").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"tunnel_interface_allow_ssh_variable": schema.StringAttribute{
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
			"tunnel_interface_allow_https": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny HTTPS").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"tunnel_interface_allow_https_variable": schema.StringAttribute{
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
			"tunnel_interface_allow_netconf": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny NETCONF").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_allow_netconf_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tunnel_interface_allow_bfd": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allow/Deny BFD").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"tunnel_interface_allow_bfd_variable": schema.StringAttribute{
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
								int64validator.AtMost(4294967295),
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
					},
				},
			},
			"nat_ipv4": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("enable Network Address Translation on this interface").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"nat_ipv4_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"nat_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("NAT Type").AddStringEnumDescription("interface", "pool", "loopback").AddDefaultValueDescription("interface").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("interface", "pool", "loopback"),
				},
			},
			"nat_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"nat_range_start": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("NAT Pool Range Start").String,
				Optional:            true,
			},
			"nat_range_start_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"nat_range_end": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("NAT Pool Range End").String,
				Optional:            true,
			},
			"nat_range_end_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"nat_prefix_length": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("NAT Pool Prefix Length").AddIntegerRangeDescription(1, 32).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 32),
				},
			},
			"nat_prefix_length_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"nat_overload": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("NAT Overload").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"nat_overload_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"nat_loopback": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("NAT Inside Source Loopback Interface").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"nat_loopback_variable": schema.StringAttribute{
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
			"new_static_nats": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("static NAT").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source IP address to be translated").String,
							Optional:            true,
						},
						"source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"translated_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Statically translated source IP address").String,
							Optional:            true,
						},
						"translated_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Direction of static NAT translation").AddStringEnumDescription("inside", "outside").AddDefaultValueDescription("inside").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("inside", "outside"),
							},
						},
						"source_vpn": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source VPN ID").AddIntegerRangeDescription(0, 65530).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(65530),
							},
						},
						"source_vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"nat_ipv6": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("enable Network Address Translation ipv6 on this interface").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"nat_ipv6_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"nat64": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("NAT64 on this interface").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"nat66": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("NAT66 on this interface").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"static_nat66": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("static NAT66").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source Prefix").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`((^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*(/)(\b([0-9]{1,2}|1[01][0-9]|12[0-8])\b)$))`), ""),
							},
						},
						"source_prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"translated_source_prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Translated Source Prefix").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`((^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*(/)(\b([0-9]{1,2}|1[01][0-9]|12[0-8])\b)$))`), ""),
							},
						},
						"translated_source_prefix_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source VPN ID").AddIntegerRangeDescription(0, 65530).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(65530),
							},
						},
						"source_vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"qos_adaptive": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Adaptive QoS").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"qos_adaptive_period": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Adapt Period(Minutes)").AddIntegerRangeDescription(1, 720).AddDefaultValueDescription("15").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 720),
				},
			},
			"qos_adaptive_period_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"qos_adaptive_bandwidth_upstream": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Shaping Rate Upstream").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"qos_adaptive_min_upstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Upstream min bandwidth limit (kbps)").AddIntegerRangeDescription(8, 100000000).String,
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
				MarkdownDescription: helpers.NewAttributeDescription("Upstream max bandwidth limit (kbps)").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"qos_adaptive_max_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"qos_adaptive_default_upstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Adaptive QoS default upstream bandwidth (kbps)").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"qos_adaptive_default_upstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"qos_adaptive_bandwidth_downstream": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Shaping Rate Downstream").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"qos_adaptive_min_downstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Downstream min bandwidth limit (kbps)").AddIntegerRangeDescription(8, 100000000).String,
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
				MarkdownDescription: helpers.NewAttributeDescription("Downstream max bandwidth limit (kbps)").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"qos_adaptive_max_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"qos_adaptive_default_downstream": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Adaptive QoS default downstream bandwidth (kbps)").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"qos_adaptive_default_downstream_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"qos_shaping_rate": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Shaping Rate (Kbps)").AddIntegerRangeDescription(8, 100000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 100000000),
				},
			},
			"qos_shaping_rate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"arps": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure ARP entries").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP V4 Address").String,
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
			"icmp_redirect_disable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ICMP/ICMPv6 Redirect Disable").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"icmp_redirect_disable_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
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
				MarkdownDescription: helpers.NewAttributeDescription("Interface MTU GigabitEthernet0 <1500..1518>, Other GigabitEthernet <1500..9216> in bytes").AddIntegerRangeDescription(1500, 9216).AddDefaultValueDescription("1500").String,
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
			"tloc_extension": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Extends a local TLOC to a remote node only for vpn 0").String,
				Optional:            true,
			},
			"tloc_extension_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"gre_tunnel_source_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("GRE tunnel source IP").String,
				Optional:            true,
			},
			"gre_tunnel_source_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"xconnect": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Extend remote TLOC over a GRE tunnel to a local WAN interface").String,
				Optional:            true,
			},
			"xconnect_variable": schema.StringAttribute{
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
			"tracker": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable tracker for this interface").String,
				Optional:            true,
			},
			"tracker_variable": schema.StringAttribute{
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

func (r *TransportWANVPNInterfaceEthernetProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *TransportWANVPNInterfaceEthernetProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TransportWANVPNInterfaceEthernet

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
func (r *TransportWANVPNInterfaceEthernetProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state TransportWANVPNInterfaceEthernet

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *TransportWANVPNInterfaceEthernetProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state TransportWANVPNInterfaceEthernet

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
func (r *TransportWANVPNInterfaceEthernetProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TransportWANVPNInterfaceEthernet

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
func (r *TransportWANVPNInterfaceEthernetProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
