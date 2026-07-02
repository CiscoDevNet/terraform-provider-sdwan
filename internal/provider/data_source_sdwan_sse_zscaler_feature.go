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
	_ datasource.DataSource                     = &SSEZscalerProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure        = &SSEZscalerProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigValidators = &SSEZscalerProfileParcelDataSource{}
)

func NewSSEZscalerProfileParcelDataSource() datasource.DataSource {
	return &SSEZscalerProfileParcelDataSource{}
}

type SSEZscalerProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *SSEZscalerProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sse_zscaler_feature"
}

func (d *SSEZscalerProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the SSE Zscaler Feature.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Feature",
				Optional:            true,
				Computed:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Feature",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Feature",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Feature",
				Computed:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: "Feature Profile ID",
				Required:            true,
			},
			"src_vpn": schema.BoolAttribute{
				MarkdownDescription: "Share Source VPN",
				Computed:            true,
			},
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: "Interface name: IPsec when present",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_name": schema.StringAttribute{
							MarkdownDescription: "Interface name: IPsec when present",
							Computed:            true,
						},
						"auto": schema.BoolAttribute{
							MarkdownDescription: "Auto Tunnel Mode",
							Computed:            true,
						},
						"shutdown": schema.BoolAttribute{
							MarkdownDescription: "Administrative state",
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
						"unnumbered": schema.BoolAttribute{
							MarkdownDescription: "Unnumbered interface",
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
						"tunnel_source": schema.StringAttribute{
							MarkdownDescription: "Tunnel source IP Address",
							Computed:            true,
						},
						"tunnel_source_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tunnel_source_interface": schema.StringAttribute{
							MarkdownDescription: "<1..32 characters> Interface name: ge0/<0-..> or ge0/<0-..>.vlanid",
							Computed:            true,
						},
						"tunnel_source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tunnel_route_via": schema.StringAttribute{
							MarkdownDescription: "<1..32 characters> Interface name: ge0/<0-..> or ge0/<0-..>.vlanid",
							Computed:            true,
						},
						"tunnel_route_via_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tunnel_destination": schema.StringAttribute{
							MarkdownDescription: "Tunnel destination IP address",
							Computed:            true,
						},
						"tunnel_destination_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tunnel_set": schema.StringAttribute{
							MarkdownDescription: "Zscaler SSE Tunnel Provider",
							Computed:            true,
						},
						"tunnel_dc_preference": schema.StringAttribute{
							MarkdownDescription: "Zscaler SSE Tunnel Data Center",
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
						"mtu": schema.Int64Attribute{
							MarkdownDescription: "Interface MTU <576..2000>, in bytes",
							Computed:            true,
						},
						"mtu_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dpd_interval": schema.Int64Attribute{
							MarkdownDescription: "IKE keepalive interval (seconds)",
							Computed:            true,
						},
						"dpd_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dpd_retries": schema.Int64Attribute{
							MarkdownDescription: "IKE keepalive retries",
							Computed:            true,
						},
						"dpd_retries_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ike_version": schema.Int64Attribute{
							MarkdownDescription: "IKE Version <1..2>",
							Computed:            true,
						},
						"ike_version_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"pre_shared_secret": schema.StringAttribute{
							MarkdownDescription: "Use preshared key to authenticate IKE peer",
							Computed:            true,
						},
						"pre_shared_secret_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ike_rekey_interval": schema.Int64Attribute{
							MarkdownDescription: "IKE rekey interval <300..1209600> seconds",
							Computed:            true,
						},
						"ike_rekey_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ike_ciphersuite": schema.StringAttribute{
							MarkdownDescription: "IKE identity the IKE preshared secret belongs to",
							Computed:            true,
						},
						"ike_ciphersuite_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ike_group": schema.StringAttribute{
							MarkdownDescription: "IKE Diffie Hellman Groups",
							Computed:            true,
						},
						"ike_group_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"pre_shared_key_dynamic": schema.BoolAttribute{
							MarkdownDescription: "Use preshared key to authenticate IKE peer",
							Computed:            true,
						},
						"ike_local_id": schema.StringAttribute{
							MarkdownDescription: "IKE ID for the local endpoint. Input IPv4 address, domain name, or email address",
							Computed:            true,
						},
						"ike_local_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ike_remote_id": schema.StringAttribute{
							MarkdownDescription: "IKE ID for the remote endpoint. Input IPv4 address, domain name, or email address",
							Computed:            true,
						},
						"ike_remote_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ipsec_rekey_interval": schema.Int64Attribute{
							MarkdownDescription: "IPsec rekey interval <300..1209600> seconds",
							Computed:            true,
						},
						"ipsec_rekey_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ipsec_replay_window": schema.Int64Attribute{
							MarkdownDescription: "Replay window size 32..8192 (must be a power of 2)",
							Computed:            true,
						},
						"ipsec_replay_window_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ipsec_ciphersuite": schema.StringAttribute{
							MarkdownDescription: "IPsec(ESP) encryption and integrity protocol",
							Computed:            true,
						},
						"ipsec_ciphersuite_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"perfect_forward_secrecy": schema.StringAttribute{
							MarkdownDescription: "IPsec perfect forward secrecy settings",
							Computed:            true,
						},
						"perfect_forward_secrecy_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"tracker": schema.StringAttribute{
							MarkdownDescription: "Enable tracker for this interface",
							Computed:            true,
						},
						"track_enable": schema.BoolAttribute{
							MarkdownDescription: "Enable/disable Zscaler SSE tracking",
							Computed:            true,
						},
						"tunnel_public_ip": schema.StringAttribute{
							MarkdownDescription: "Public IP required to setup GRE tunnel to Zscaler",
							Computed:            true,
						},
						"tunnel_public_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"interface_pairs": schema.ListNestedAttribute{
				MarkdownDescription: "Interface Pair for active and backup",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"active_interface": schema.StringAttribute{
							MarkdownDescription: "Active Tunnel Interface for Zscaler SSE",
							Computed:            true,
						},
						"active_interface_weight": schema.Int64Attribute{
							MarkdownDescription: "Active Tunnel Interface Weight",
							Computed:            true,
						},
						"backup_interface": schema.StringAttribute{
							MarkdownDescription: "Backup Tunnel Interface for Zscaler SSE",
							Computed:            true,
						},
						"backup_interface_weight": schema.Int64Attribute{
							MarkdownDescription: "Backup Tunnel Interface Weight",
							Computed:            true,
						},
					},
				},
			},
			"auth_required": schema.BoolAttribute{
				MarkdownDescription: "Enforce Authentication",
				Computed:            true,
			},
			"auth_required_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"xff_forward_enabled": schema.BoolAttribute{
				MarkdownDescription: "XFF forwarding enabled",
				Computed:            true,
			},
			"xff_forward_enabled_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ofw_enabled": schema.BoolAttribute{
				MarkdownDescription: "Firewall enabled",
				Computed:            true,
			},
			"ofw_enabled_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ips_control": schema.BoolAttribute{
				MarkdownDescription: "Enable IPS Control",
				Computed:            true,
			},
			"ips_control_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"caution_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable Caution",
				Computed:            true,
			},
			"caution_enabled_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"primary_data_center": schema.StringAttribute{
				MarkdownDescription: "Custom Primary Datacenter",
				Computed:            true,
			},
			"primary_data_center_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"secondary_data_center": schema.StringAttribute{
				MarkdownDescription: "Custom Secondary Datacenter",
				Computed:            true,
			},
			"secondary_data_center_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"surrogate_ip": schema.BoolAttribute{
				MarkdownDescription: "Enable Surrogate IP",
				Computed:            true,
			},
			"surrogate_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"idle_time": schema.Int64Attribute{
				MarkdownDescription: "Idle time to disassociation",
				Computed:            true,
			},
			"idle_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"display_time_unit": schema.StringAttribute{
				MarkdownDescription: "Display time unit",
				Computed:            true,
			},
			"display_time_unit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ip_enforced_for_known_browsers": schema.BoolAttribute{
				MarkdownDescription: "Enforce Surrogate IP for known browsers",
				Computed:            true,
			},
			"ip_enforced_for_known_browsers_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"refresh_time": schema.Int64Attribute{
				MarkdownDescription: "Refresh time for re-validation of surrogacy in minutes",
				Computed:            true,
			},
			"refresh_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"refresh_time_unit": schema.StringAttribute{
				MarkdownDescription: "Refresh Time unit",
				Computed:            true,
			},
			"refresh_time_unit_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"aup_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable Acceptable User Policy",
				Computed:            true,
			},
			"aup_enabled_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"block_internet_until_accepted": schema.BoolAttribute{
				MarkdownDescription: "For first-time Acceptable User Policy behavior, block Internet access",
				Computed:            true,
			},
			"block_internet_until_accepted_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"force_ssl_inspection": schema.BoolAttribute{
				MarkdownDescription: "For first-time Acceptable User Policy behavior, force SSL inspection",
				Computed:            true,
			},
			"force_ssl_inspection_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"aup_timeout": schema.Int64Attribute{
				MarkdownDescription: "Custom Acceptable User Policy frequency in days",
				Computed:            true,
			},
			"aup_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"location_name": schema.StringAttribute{
				MarkdownDescription: "Zscaler location name (optional)",
				Computed:            true,
			},
			"location_name_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"country": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"country_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"enforce_bandwidth_control": schema.BoolAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"enforce_bandwidth_control_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"dn_bandwidth": schema.Float64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"dn_bandwidth_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"up_bandwidth": schema.Float64Attribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"up_bandwidth_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"sub_locations": schema.ListNestedAttribute{
				MarkdownDescription: "",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"auth_required": schema.BoolAttribute{
							MarkdownDescription: "Enable Enforce Authentication to require users from this location to authenticate to the service.",
							Computed:            true,
						},
						"auth_required_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"surrogate_ip": schema.BoolAttribute{
							MarkdownDescription: "Enable Surrogate IP. Maps users to device IP addresses. This is used to enforce user policies on cookie-compatible traffic.",
							Computed:            true,
						},
						"surrogate_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"idle_time": schema.Int64Attribute{
							MarkdownDescription: "Idle time to disassociation. How long after a completed transaction the service retains the IP address to user mapping.",
							Computed:            true,
						},
						"idle_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"display_time_unit": schema.StringAttribute{
							MarkdownDescription: "Display time unit for Idle Time to Disassociation",
							Computed:            true,
						},
						"display_time_unit_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ip_enforced_for_known_browsers": schema.BoolAttribute{
							MarkdownDescription: "If enabled, and if the IP-user mapping exists, then the Surrogate user identity is used for traffic from known browsers. If disabled, traffic from known browsers will always be challenged using the configured authentication mechanism and Surrogate user identity is ignored",
							Computed:            true,
						},
						"ip_enforced_for_known_browsers_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"refresh_time": schema.Int64Attribute{
							MarkdownDescription: "Length of time that surrogate user identity can be used for traffic from known browsers before it must refresh and re-validate the surrogate user identity",
							Computed:            true,
						},
						"refresh_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"refresh_time_unit": schema.StringAttribute{
							MarkdownDescription: "Refresh Time display unit",
							Computed:            true,
						},
						"refresh_time_unit_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ofw_enabled": schema.BoolAttribute{
							MarkdownDescription: "Enforces firewall at the location",
							Computed:            true,
						},
						"ofw_enabled_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"caution_enabled": schema.BoolAttribute{
							MarkdownDescription: "Enforces a caution policy action and display an end user notification for unauthenticated traffic. If disabled, the action is treated as an allow policy.",
							Computed:            true,
						},
						"caution_enabled_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"service_vpn": schema.SetAttribute{
							MarkdownDescription: "",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"service_vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"internal_ip": schema.ListNestedAttribute{
							MarkdownDescription: "",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"internal_ip_value": schema.StringAttribute{
										MarkdownDescription: "",
										Computed:            true,
									},
									"internal_ip_value_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
								},
							},
						},
						"aup_enabled": schema.BoolAttribute{
							MarkdownDescription: "Displays an Acceptable Use Policy for unauthenticated traffic and require users to accept it",
							Computed:            true,
						},
						"aup_enabled_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"block_internet_until_accepted": schema.BoolAttribute{
							MarkdownDescription: "Disable all access to the internet, including non-HTTP traffic, until the user accepts the Acceptable Use Policy",
							Computed:            true,
						},
						"block_internet_until_accepted_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"force_ssl_inspection": schema.BoolAttribute{
							MarkdownDescription: "Enable to make SSL Interception enforce an Acceptable Use Policy for HTTPS traffic",
							Computed:            true,
						},
						"force_ssl_inspection_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"aup_timeout": schema.Int64Attribute{
							MarkdownDescription: "How frequently in days the Acceptable Use Policy is displayed to users",
							Computed:            true,
						},
						"aup_timeout_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"enforce_bandwidth_control": schema.StringAttribute{
							MarkdownDescription: "Enforce Bandwidth Control for sub location",
							Computed:            true,
						},
						"enforce_bandwidth_control_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"up_bandwidth": schema.Float64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"up_bandwidth_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"dn_bandwidth": schema.Float64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
						"dn_bandwidth_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"tracker_source_ip": schema.StringAttribute{
				MarkdownDescription: "Source IP address for Tracker",
				Computed:            true,
			},
			"tracker_source_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"trackers": schema.ListNestedAttribute{
				MarkdownDescription: "Tracker configuration",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Tracker name",
							Computed:            true,
						},
						"endpoint_api_url": schema.StringAttribute{
							MarkdownDescription: "API url of endpoint",
							Computed:            true,
						},
						"endpoint_api_url_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"threshold": schema.Int64Attribute{
							MarkdownDescription: "Probe Timeout threshold <100..1000> milliseconds",
							Computed:            true,
						},
						"threshold_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"interval": schema.Int64Attribute{
							MarkdownDescription: "Probe interval <10..600> seconds",
							Computed:            true,
						},
						"interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"multiplier": schema.Int64Attribute{
							MarkdownDescription: "Probe failure multiplier <1..10> failed attempts",
							Computed:            true,
						},
						"multiplier_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SSEZscalerProfileParcelDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *SSEZscalerProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *SSEZscalerProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SSEZscaler

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))
	if config.Id.IsNull() && !config.Name.IsNull() {
		// Look up parcel ID by name
		res, err := d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve parcels, got error: %s", err))
			return
		}
		found := false
		res.Get("data").ForEach(func(_, v gjson.Result) bool {
			if v.Get("payload.name").String() == config.Name.ValueString() {
				config.Id = types.StringValue(v.Get("parcelId").String())
				found = true
				return false
			}
			return true
		})
		if !found {
			resp.Diagnostics.AddError("Not Found", fmt.Sprintf("No parcel found with name: %s", config.Name.ValueString()))
			return
		}
	}

	res, err := d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res, true)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
