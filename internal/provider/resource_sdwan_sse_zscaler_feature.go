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
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
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
var _ resource.Resource = &SSEZscalerProfileParcelResource{}
var _ resource.ResourceWithImportState = &SSEZscalerProfileParcelResource{}

func NewSSEZscalerProfileParcelResource() resource.Resource {
	return &SSEZscalerProfileParcelResource{}
}

type SSEZscalerProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *SSEZscalerProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sse_zscaler_feature"
}

func (r *SSEZscalerProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a SSE Zscaler Feature.").AddMinimumVersionDescription("20.15.0").String,

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
			"src_vpn": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Share Source VPN").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface name: IPsec when present").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface name: IPsec when present").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(4, 8),
								stringvalidator.RegexMatches(regexp.MustCompile(`(^ipsec(?:1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)$)|(^gre(?:1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)$)`), ""),
							},
						},
						"auto": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Auto Tunnel Mode").String,
							Optional:            true,
						},
						"shutdown": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Administrative state").AddDefaultValueDescription("false").String,
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
						"unnumbered": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Unnumbered interface").String,
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
						"tunnel_source": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tunnel source IP Address").String,
							Optional:            true,
						},
						"tunnel_source_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"tunnel_source_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("<1..32 characters> Interface name: ge0/<0-..> or ge0/<0-..>.vlanid").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(3, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`(ATM|ATM-ACR|AppGigabitEthernet|AppNav-Compress|AppNav-UnCompress|Async|BD-VIF|BDI|CEM|CEM-ACR|Cellular|Dialer|Embedded-Service-Engine|Ethernet|Ethernet-Internal|FastEthernet|FiftyGigabitEthernet|FiveGigabitEthernet|FortyGigabitEthernet|FourHundredGigE|GMPLS|GigabitEthernet|Group-Async|HundredGigE|L2LISP|LISP|Loopback|MFR|Multilink|Port-channel|SM|Serial|Service-Engine|TenGigabitEthernet|Tunnel|TwentyFiveGigE|TwentyFiveGigabitEthernet|TwoGigabitEthernet|TwoHundredGigE|Vif|Virtual-PPP|Virtual-Template|VirtualPortGroup|Vlan|Wlan-GigabitEthernet|nat64|nat66|ntp|nve|ospfv3|overlay|pseudowire|ucse|vasileft|vasiright|vmi)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
							},
						},
						"tunnel_source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"tunnel_route_via": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("<1..32 characters> Interface name: ge0/<0-..> or ge0/<0-..>.vlanid").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"tunnel_route_via_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"tunnel_destination": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tunnel destination IP address").String,
							Optional:            true,
						},
						"tunnel_destination_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"tunnel_set": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Zscaler SSE Tunnel Provider").AddStringEnumDescription("secure-internet-gateway-umbrella", "secure-internet-gateway-zscaler", "secure-internet-gateway-other").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("secure-internet-gateway-umbrella", "secure-internet-gateway-zscaler", "secure-internet-gateway-other"),
							},
						},
						"tunnel_dc_preference": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Zscaler SSE Tunnel Data Center").AddStringEnumDescription("primary-dc", "secondary-dc").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("primary-dc", "secondary-dc"),
							},
						},
						"tcp_mss_adjust": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("TCP MSS on SYN packets, in bytes").AddIntegerRangeDescription(500, 1460).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(500, 1460),
							},
						},
						"tcp_mss_adjust_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"mtu": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface MTU <576..2000>, in bytes").AddIntegerRangeDescription(576, 2000).AddDefaultValueDescription("1400").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(576, 2000),
							},
						},
						"mtu_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dpd_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE keepalive interval (seconds)").AddIntegerRangeDescription(10, 3600).AddDefaultValueDescription("10").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(10, 3600),
							},
						},
						"dpd_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dpd_retries": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE keepalive retries").AddIntegerRangeDescription(2, 60).AddDefaultValueDescription("3").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(2, 60),
							},
						},
						"dpd_retries_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ike_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE Version <1..2>").AddIntegerRangeDescription(1, 2).AddDefaultValueDescription("2").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 2),
							},
						},
						"ike_version_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"pre_shared_secret": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use preshared key to authenticate IKE peer").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 127),
							},
						},
						"pre_shared_secret_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ike_rekey_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE rekey interval <300..1209600> seconds").AddIntegerRangeDescription(300, 86400).AddDefaultValueDescription("14400").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(300, 86400),
							},
						},
						"ike_rekey_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ike_ciphersuite": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE identity the IKE preshared secret belongs to").AddStringEnumDescription("aes256-cbc-sha1", "aes256-cbc-sha2", "aes128-cbc-sha1", "aes128-cbc-sha2").AddDefaultValueDescription("aes256-cbc-sha1").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("aes256-cbc-sha1", "aes256-cbc-sha2", "aes128-cbc-sha1", "aes128-cbc-sha2"),
							},
						},
						"ike_ciphersuite_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ike_group": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE Diffie Hellman Groups").AddStringEnumDescription("2", "5", "14", "15", "16", "19", "20", "21").AddDefaultValueDescription("16").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2", "5", "14", "15", "16", "19", "20", "21"),
							},
						},
						"ike_group_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"pre_shared_key_dynamic": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use preshared key to authenticate IKE peer").String,
							Optional:            true,
						},
						"ike_local_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE ID for the local endpoint. Input IPv4 address, domain name, or email address").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 63),
							},
						},
						"ike_local_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ike_remote_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE ID for the remote endpoint. Input IPv4 address, domain name, or email address").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 63),
							},
						},
						"ike_remote_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ipsec_rekey_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPsec rekey interval <300..1209600> seconds").AddIntegerRangeDescription(300, 1209600).AddDefaultValueDescription("3600").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(300, 1209600),
							},
						},
						"ipsec_rekey_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ipsec_replay_window": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Replay window size 32..8192 (must be a power of 2)").AddDefaultValueDescription("512").String,
							Optional:            true,
						},
						"ipsec_replay_window_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ipsec_ciphersuite": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPsec(ESP) encryption and integrity protocol").AddStringEnumDescription("aes256-cbc-sha1", "aes256-cbc-sha384", "aes256-cbc-sha256", "aes256-cbc-sha512", "aes256-gcm").AddDefaultValueDescription("aes256-cbc-sha512").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("aes256-cbc-sha1", "aes256-cbc-sha384", "aes256-cbc-sha256", "aes256-cbc-sha512", "aes256-gcm"),
							},
						},
						"ipsec_ciphersuite_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"perfect_forward_secrecy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPsec perfect forward secrecy settings").AddStringEnumDescription("group-2", "group-5", "group-14", "group-15", "group-16", "group-19", "group-20", "group-21", "none").AddDefaultValueDescription("none").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("group-2", "group-5", "group-14", "group-15", "group-16", "group-19", "group-20", "group-21", "none"),
							},
						},
						"perfect_forward_secrecy_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"tracker": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable tracker for this interface").AddDefaultValueDescription("DefaultTracker").String,
							Optional:            true,
						},
						"track_enable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable/disable Zscaler SSE tracking").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"tunnel_public_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Public IP required to setup GRE tunnel to Zscaler").AddDefaultValueDescription("Auto").String,
							Optional:            true,
						},
						"tunnel_public_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"interface_pairs": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface Pair for active and backup").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"active_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Active Tunnel Interface for Zscaler SSE").String,
							Optional:            true,
						},
						"active_interface_weight": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Active Tunnel Interface Weight").AddIntegerRangeDescription(1, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
						"backup_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Backup Tunnel Interface for Zscaler SSE").String,
							Optional:            true,
						},
						"backup_interface_weight": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Backup Tunnel Interface Weight").AddIntegerRangeDescription(1, 255).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 255),
							},
						},
					},
				},
			},
			"auth_required": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enforce Authentication").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"auth_required_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"xff_forward_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("XFF forwarding enabled").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"xff_forward_enabled_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ofw_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Firewall enabled").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ofw_enabled_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ips_control": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable IPS Control").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ips_control_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"caution_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Caution").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"caution_enabled_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"primary_data_center": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Custom Primary Datacenter").AddDefaultValueDescription("Auto").String,
				Optional:            true,
			},
			"primary_data_center_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"secondary_data_center": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Custom Secondary Datacenter").AddDefaultValueDescription("Auto").String,
				Optional:            true,
			},
			"secondary_data_center_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"surrogate_ip": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Surrogate IP, Attribute conditional on `auth_required` equal to `true`").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"surrogate_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `auth_required` equal to `true`").String,
				Optional:            true,
			},
			"idle_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Idle time to disassociation, Attribute conditional on `surrogate_ip` equal to `true`").AddIntegerAtLeastDescription(1).AddDefaultValueDescription("1").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"idle_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `surrogate_ip` equal to `true`").String,
				Optional:            true,
			},
			"display_time_unit": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Display time unit, Attribute conditional on `surrogate_ip` equal to `true`").AddStringEnumDescription("MINUTE", "HOUR", "DAY").AddDefaultValueDescription("MINUTE").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("MINUTE", "HOUR", "DAY"),
				},
			},
			"display_time_unit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `surrogate_ip` equal to `true`").String,
				Optional:            true,
			},
			"ip_enforced_for_known_browsers": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enforce Surrogate IP for known browsers").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"ip_enforced_for_known_browsers_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"refresh_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Refresh time for re-validation of surrogacy in minutes").AddIntegerAtLeastDescription(1).AddDefaultValueDescription("1").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"refresh_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"refresh_time_unit": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Refresh Time unit").AddStringEnumDescription("MINUTE", "HOUR", "DAY").AddDefaultValueDescription("MINUTE").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("MINUTE", "HOUR", "DAY"),
				},
			},
			"refresh_time_unit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"aup_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Acceptable User Policy").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"aup_enabled_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"block_internet_until_accepted": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("For first-time Acceptable User Policy behavior, block Internet access, Attribute conditional on `aup_enabled` equal to `true`").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"block_internet_until_accepted_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `aup_enabled` equal to `true`").String,
				Optional:            true,
			},
			"force_ssl_inspection": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("For first-time Acceptable User Policy behavior, force SSL inspection, Attribute conditional on `aup_enabled` equal to `true`").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"force_ssl_inspection_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `aup_enabled` equal to `true`").String,
				Optional:            true,
			},
			"aup_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Custom Acceptable User Policy frequency in days, Attribute conditional on `aup_enabled` equal to `true`").AddIntegerRangeDescription(1, 180).AddDefaultValueDescription("1").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 180),
				},
			},
			"aup_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `aup_enabled` equal to `true`").String,
				Optional:            true,
			},
			"location_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Zscaler location name (optional)").AddDefaultValueDescription("Auto").String,
				Optional:            true,
			},
			"location_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"country": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"country_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"enforce_bandwidth_control": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"enforce_bandwidth_control_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"dn_bandwidth": schema.Float64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription(", Attribute conditional on `enforce_bandwidth_control` equal to `true`").AddFloatRangeDescription(0.1, 99999).String,
				Optional:            true,
				Validators: []validator.Float64{
					float64validator.Between(0.1, 99999),
				},
			},
			"dn_bandwidth_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `enforce_bandwidth_control` equal to `true`").String,
				Optional:            true,
			},
			"up_bandwidth": schema.Float64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription(", Attribute conditional on `enforce_bandwidth_control` equal to `true`").AddFloatRangeDescription(0.1, 99999).String,
				Optional:            true,
				Validators: []validator.Float64{
					float64validator.Between(0.1, 99999),
				},
			},
			"up_bandwidth_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `enforce_bandwidth_control` equal to `true`").String,
				Optional:            true,
			},
			"sub_locations": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 128),
								stringvalidator.RegexMatches(regexp.MustCompile(`^[^&<>! "]+$`), ""),
							},
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"auth_required": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable Enforce Authentication to require users from this location to authenticate to the service.").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"auth_required_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"surrogate_ip": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable Surrogate IP. Maps users to device IP addresses. This is used to enforce user policies on cookie-compatible traffic., Attribute conditional on `auth_required` equal to `true`").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"surrogate_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `auth_required` equal to `true`").String,
							Optional:            true,
						},
						"idle_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Idle time to disassociation. How long after a completed transaction the service retains the IP address to user mapping., Attribute conditional on `surrogate_ip` equal to `true`").AddIntegerAtLeastDescription(1).AddDefaultValueDescription("1").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtLeast(1),
							},
						},
						"idle_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `surrogate_ip` equal to `true`").String,
							Optional:            true,
						},
						"display_time_unit": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Display time unit for Idle Time to Disassociation, Attribute conditional on `surrogate_ip` equal to `true`").AddStringEnumDescription("MINUTE", "HOUR", "DAY").AddDefaultValueDescription("MINUTE").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("MINUTE", "HOUR", "DAY"),
							},
						},
						"display_time_unit_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `surrogate_ip` equal to `true`").String,
							Optional:            true,
						},
						"ip_enforced_for_known_browsers": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("If enabled, and if the IP-user mapping exists, then the Surrogate user identity is used for traffic from known browsers. If disabled, traffic from known browsers will always be challenged using the configured authentication mechanism and Surrogate user identity is ignored").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"ip_enforced_for_known_browsers_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"refresh_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Length of time that surrogate user identity can be used for traffic from known browsers before it must refresh and re-validate the surrogate user identity, Attribute conditional on `ip_enforced_for_known_browsers` equal to `true`").AddIntegerAtLeastDescription(1).AddDefaultValueDescription("1").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtLeast(1),
							},
						},
						"refresh_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `ip_enforced_for_known_browsers` equal to `true`").String,
							Optional:            true,
						},
						"refresh_time_unit": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Refresh Time display unit, Attribute conditional on `ip_enforced_for_known_browsers` equal to `true`").AddStringEnumDescription("MINUTE", "HOUR", "DAY").AddDefaultValueDescription("MINUTE").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("MINUTE", "HOUR", "DAY"),
							},
						},
						"refresh_time_unit_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `ip_enforced_for_known_browsers` equal to `true`").String,
							Optional:            true,
						},
						"ofw_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enforces firewall at the location").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"ofw_enabled_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"caution_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enforces a caution policy action and display an end user notification for unauthenticated traffic. If disabled, the action is treated as an allow policy.").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"caution_enabled_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"service_vpn": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"service_vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"internal_ip": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"internal_ip_value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
									},
									"internal_ip_value_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
								},
							},
						},
						"aup_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Displays an Acceptable Use Policy for unauthenticated traffic and require users to accept it").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"aup_enabled_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"block_internet_until_accepted": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Disable all access to the internet, including non-HTTP traffic, until the user accepts the Acceptable Use Policy, Attribute conditional on `aup_enabled` equal to `true`").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"block_internet_until_accepted_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `aup_enabled` equal to `true`").String,
							Optional:            true,
						},
						"force_ssl_inspection": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable to make SSL Interception enforce an Acceptable Use Policy for HTTPS traffic, Attribute conditional on `aup_enabled` equal to `true`").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"force_ssl_inspection_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `aup_enabled` equal to `true`").String,
							Optional:            true,
						},
						"aup_timeout": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("How frequently in days the Acceptable Use Policy is displayed to users, Attribute conditional on `aup_enabled` equal to `true`").AddIntegerRangeDescription(1, 180).AddDefaultValueDescription("1").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 180),
							},
						},
						"aup_timeout_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `aup_enabled` equal to `true`").String,
							Optional:            true,
						},
						"enforce_bandwidth_control": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enforce Bandwidth Control for sub location").AddStringEnumDescription("location-bandwidth", "override", "disabled").AddDefaultValueDescription("location-bandwidth").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("location-bandwidth", "override", "disabled"),
							},
						},
						"enforce_bandwidth_control_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"up_bandwidth": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription(", Attribute conditional on `enforce_bandwidth_control` equal to `override`").AddFloatRangeDescription(0.1, 99999).String,
							Optional:            true,
							Validators: []validator.Float64{
								float64validator.Between(0.1, 99999),
							},
						},
						"up_bandwidth_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `enforce_bandwidth_control` equal to `override`").String,
							Optional:            true,
						},
						"dn_bandwidth": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription(", Attribute conditional on `enforce_bandwidth_control` equal to `override`").AddFloatRangeDescription(0.1, 99999).String,
							Optional:            true,
							Validators: []validator.Float64{
								float64validator.Between(0.1, 99999),
							},
						},
						"dn_bandwidth_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `enforce_bandwidth_control` equal to `override`").String,
							Optional:            true,
						},
					},
				},
			},
			"tracker_source_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Source IP address for Tracker").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)\.){3}(25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)$`), ""),
				},
			},
			"tracker_source_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"trackers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Tracker configuration").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tracker name").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 128),
								stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9\._-]+)*$`), ""),
							},
						},
						"endpoint_api_url": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("API url of endpoint").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthAtMost(512),
								stringvalidator.RegexMatches(regexp.MustCompile(`^http\:\/\/[0-9a-zA-Z]([-.\w]*[0-9a-zA-Z])*(:(0-9)*)*(\/?)([a-zA-Z0-9\-\.\?\,\'\/\\\+&%\$#_]*)?$`), ""),
							},
						},
						"endpoint_api_url_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"threshold": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Probe Timeout threshold <100..1000> milliseconds").AddIntegerRangeDescription(100, 1000).AddDefaultValueDescription("1000").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(100, 1000),
							},
						},
						"threshold_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Probe interval <10..600> seconds").AddIntegerRangeDescription(20, 600).AddDefaultValueDescription("30").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(20, 600),
							},
						},
						"interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"multiplier": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Probe failure multiplier <1..10> failed attempts").AddIntegerRangeDescription(1, 10).AddDefaultValueDescription("2").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 10),
							},
						},
						"multiplier_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *SSEZscalerProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *SSEZscalerProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SSEZscaler

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
func (r *SSEZscalerProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SSEZscaler

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
func (r *SSEZscalerProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SSEZscaler

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
func (r *SSEZscalerProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SSEZscaler

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
func (r *SSEZscalerProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "sse_zscaler_feature_id" + ",feature_profile_id"
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
