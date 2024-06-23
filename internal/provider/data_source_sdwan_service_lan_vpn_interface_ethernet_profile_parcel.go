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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ServiceLANVPNInterfaceEthernetProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &ServiceLANVPNInterfaceEthernetProfileParcelDataSource{}
)

func NewServiceLANVPNInterfaceEthernetProfileParcelDataSource() datasource.DataSource {
	return &ServiceLANVPNInterfaceEthernetProfileParcelDataSource{}
}

type ServiceLANVPNInterfaceEthernetProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *ServiceLANVPNInterfaceEthernetProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_lan_vpn_interface_ethernet_profile_parcel"
}

func (d *ServiceLANVPNInterfaceEthernetProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Service LAN VPN Interface Ethernet profile parcel.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the profile parcel",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the profile parcel",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the profile parcel",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the profile parcel",
				Computed:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: "Feature Profile ID",
				Required:            true,
			},
			"service_lan_vpn_profile_parcel_id": schema.StringAttribute{
				MarkdownDescription: "Service LAN VPN Profile Profile ID",
				Required:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"shutdown_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"interface_name": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"interface_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"interface_description": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"interface_description_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_dhcp_distance": schema.Int64Attribute{
				MarkdownDescription: "DHCP Distance",
				Computed:            true,
			},
			"ipv4_dhcp_distance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_address": schema.StringAttribute{
				MarkdownDescription: "IP Address",
				Computed:            true,
			},
			"ipv4_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_subnet_mask": schema.StringAttribute{
				MarkdownDescription: "Subnet Mask",
				Computed:            true,
			},
			"ipv4_subnet_mask_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_secondary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Secondary IpV4 Addresses",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IpV4 Address",
							Computed:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: "Subnet Mask",
							Computed:            true,
						},
						"subnet_mask_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"ipv4_dhcp_helper": schema.SetAttribute{
				MarkdownDescription: "List of DHCP IPv4 helper addresses (min 1, max 8)",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ipv4_dhcp_helper_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"enable_dhcpv6": schema.BoolAttribute{
				MarkdownDescription: "Enable DHCPv6",
				Computed:            true,
			},
			"ipv6_dhcp_secondary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "secondary IPv6 addresses",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IPv6 Address Secondary",
							Computed:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"ipv6_address": schema.StringAttribute{
				MarkdownDescription: "IPv6 Address Secondary",
				Computed:            true,
			},
			"ipv6_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv6_secondary_addresses": schema.ListNestedAttribute{
				MarkdownDescription: "Static secondary IPv6 addresses",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							MarkdownDescription: "IPv6 Address Secondary",
							Computed:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
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
						"dhcpv6_helper_vpn": schema.Int64Attribute{
							MarkdownDescription: "DHCPv6 Helper VPN",
							Computed:            true,
						},
						"dhcpv6_helper_vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"ipv4_nat": schema.BoolAttribute{
				MarkdownDescription: "enable Network Address Translation on this interface",
				Computed:            true,
			},
			"ipv4_nat_type": schema.StringAttribute{
				MarkdownDescription: "NAT Type",
				Computed:            true,
			},
			"ipv4_nat_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_nat_range_start": schema.StringAttribute{
				MarkdownDescription: "NAT Pool Range Start",
				Computed:            true,
			},
			"ipv4_nat_range_start_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_nat_range_end": schema.StringAttribute{
				MarkdownDescription: "NAT Pool Range End",
				Computed:            true,
			},
			"ipv4_nat_range_end_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_nat_prefix_length": schema.Int64Attribute{
				MarkdownDescription: "NAT Pool Prefix Length",
				Computed:            true,
			},
			"ipv4_nat_prefix_length_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_nat_overload": schema.BoolAttribute{
				MarkdownDescription: "NAT Overload",
				Computed:            true,
			},
			"ipv4_nat_overload_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_nat_loopback": schema.StringAttribute{
				MarkdownDescription: "NAT Inside Source Loopback Interface",
				Computed:            true,
			},
			"ipv4_nat_loopback_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_nat_udp_timeout": schema.Int64Attribute{
				MarkdownDescription: "Set NAT UDP session timeout, in minutes",
				Computed:            true,
			},
			"ipv4_nat_udp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ipv4_nat_tcp_timeout": schema.Int64Attribute{
				MarkdownDescription: "Set NAT TCP session timeout, in minutes",
				Computed:            true,
			},
			"ipv4_nat_tcp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"static_nats": schema.ListNestedAttribute{
				MarkdownDescription: "static NAT",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source_ip": schema.StringAttribute{
							MarkdownDescription: "Source IP address to be translated",
							Computed:            true,
						},
						"source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"translate_ip": schema.StringAttribute{
							MarkdownDescription: "Statically translated source IP address",
							Computed:            true,
						},
						"translate_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"direction": schema.StringAttribute{
							MarkdownDescription: "Direction of static NAT translation",
							Computed:            true,
						},
						"source_vpn": schema.Int64Attribute{
							MarkdownDescription: "Source VPN ID",
							Computed:            true,
						},
						"source_vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"ipv6_nat": schema.BoolAttribute{
				MarkdownDescription: "enable Network Address Translation ipv6 on this interface",
				Computed:            true,
			},
			"nat64": schema.BoolAttribute{
				MarkdownDescription: "NAT64 on this interface",
				Computed:            true,
			},
			"acl_shaping_rate": schema.Int64Attribute{
				MarkdownDescription: "Shaping Rate (Kbps)",
				Computed:            true,
			},
			"acl_shaping_rate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"acl_ipv4_egress_policy_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"acl_ipv4_ingress_policy_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"acl_ipv6_egress_policy_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"acl_ipv6_ingress_policy_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"ipv6_vrrps": schema.ListNestedAttribute{
				MarkdownDescription: "Enable VRRP Ipv6",
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
									"global_address": schema.StringAttribute{
										MarkdownDescription: "Assign Global IPv6 Prefix",
										Computed:            true,
									},
									"global_address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
								},
							},
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
						"address": schema.StringAttribute{
							MarkdownDescription: "VRRP Ip Address",
							Computed:            true,
						},
						"address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"secondary_addresses": schema.ListNestedAttribute{
							MarkdownDescription: "VRRP Secondary Ip Addresses",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										MarkdownDescription: "Ip Address",
										Computed:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"subnet_mask": schema.StringAttribute{
										MarkdownDescription: "Subnet Mask",
										Computed:            true,
									},
									"subnet_mask_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
								},
							},
						},
						"tloc_prefix_change": schema.BoolAttribute{
							MarkdownDescription: "Timer interval for successive advertisements, in milliseconds",
							Computed:            true,
						},
						"tloc_pref_change_value": schema.Int64Attribute{
							MarkdownDescription: "Timer interval for successive advertisements, in milliseconds",
							Computed:            true,
						},
					},
				},
			},
			"arps": schema.ListNestedAttribute{
				MarkdownDescription: "Configure ARP entries",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "IPV4 Address",
							Computed:            true,
						},
						"ip_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"mac_address": schema.StringAttribute{
							MarkdownDescription: "MAC Address",
							Computed:            true,
						},
						"mac_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"trustsec_enable_sgt_propogation": schema.BoolAttribute{
				MarkdownDescription: "Indicates that the interface is trustworthy for CTS",
				Computed:            true,
			},
			"trustsec_propogate": schema.BoolAttribute{
				MarkdownDescription: "Enables the interface for CTS SGT authorization and forwarding",
				Computed:            true,
			},
			"trustsec_security_group_tag": schema.Int64Attribute{
				MarkdownDescription: "SGT value between 2 and 65519",
				Computed:            true,
			},
			"trustsec_security_group_tag_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"trustsec_enable_enforced_propogation": schema.BoolAttribute{
				MarkdownDescription: "Enable/Disable SGT Enforcement on an interface",
				Computed:            true,
			},
			"trustsec_enforced_security_group_tag": schema.Int64Attribute{
				MarkdownDescription: "SGT value between 2 and 65519",
				Computed:            true,
			},
			"trustsec_enforced_security_group_tag_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"duplex": schema.StringAttribute{
				MarkdownDescription: "Duplex mode",
				Computed:            true,
			},
			"duplex_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"mac_address": schema.StringAttribute{
				MarkdownDescription: "MAC Address",
				Computed:            true,
			},
			"mac_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ip_mtu": schema.Int64Attribute{
				MarkdownDescription: "IP MTU for GigabitEthernet main <576..Interface MTU>, GigabitEthernet subinterface <576..9216>, Other Interfaces <576..2000> in bytes",
				Computed:            true,
			},
			"ip_mtu_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"interface_mtu": schema.Int64Attribute{
				MarkdownDescription: "Interface MTU",
				Computed:            true,
			},
			"interface_mtu_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tcp_mss": schema.Int64Attribute{
				MarkdownDescription: "TCP MSS on SYN packets, in bytes",
				Computed:            true,
			},
			"tcp_mss_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"speed": schema.StringAttribute{
				MarkdownDescription: "Set interface speed",
				Computed:            true,
			},
			"speed_variable": schema.StringAttribute{
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
			"autonegotiate": schema.BoolAttribute{
				MarkdownDescription: "Link autonegotiation",
				Computed:            true,
			},
			"autonegotiate_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"media_type": schema.StringAttribute{
				MarkdownDescription: "Media type",
				Computed:            true,
			},
			"media_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"load_interval": schema.Int64Attribute{
				MarkdownDescription: "Interval for interface load calculation",
				Computed:            true,
			},
			"load_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tracker": schema.StringAttribute{
				MarkdownDescription: "Enable tracker for this interface",
				Computed:            true,
			},
			"tracker_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"icmp_redirect_disable": schema.BoolAttribute{
				MarkdownDescription: "ICMP/ICMPv6 Redirect Disable",
				Computed:            true,
			},
			"icmp_redirect_disable_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"xconnect": schema.StringAttribute{
				MarkdownDescription: "Extend remote TLOC over a GRE tunnel to a local LAN interface",
				Computed:            true,
			},
			"xconnect_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ip_directed_broadcast": schema.BoolAttribute{
				MarkdownDescription: "IP Directed-Broadcast",
				Computed:            true,
			},
			"ip_directed_broadcast_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
		},
	}
}

func (d *ServiceLANVPNInterfaceEthernetProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *ServiceLANVPNInterfaceEthernetProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ServiceLANVPNInterfaceEthernet

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
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
