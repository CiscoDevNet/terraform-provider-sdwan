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
var _ resource.Resource = &CiscoSecureInternetGatewayFeatureTemplateResource{}
var _ resource.ResourceWithImportState = &CiscoSecureInternetGatewayFeatureTemplateResource{}

func NewCiscoSecureInternetGatewayFeatureTemplateResource() resource.Resource {
	return &CiscoSecureInternetGatewayFeatureTemplateResource{}
}

type CiscoSecureInternetGatewayFeatureTemplateResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *CiscoSecureInternetGatewayFeatureTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_secure_internet_gateway_feature_template"
}

func (r *CiscoSecureInternetGatewayFeatureTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Cisco Secure Internet Gateway feature template.").AddMinimumVersionDescription("15.0.0").String,

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
			"interfaces": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface name: IPsec when present").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface name: IPsec when present").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(4, 8),
							},
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"auto_tunnel_mode": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Auto Tunnel Mode").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"shutdown": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Administrative state").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface description").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 128),
							},
						},
						"description_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ip_unnumbered": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Unnumbered interface").AddDefaultValueDescription("true").String,
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
								stringvalidator.LengthBetween(1, 32),
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
						"application": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable Application Tunnel Type").AddStringEnumDescription("sig").AddDefaultValueDescription("sig").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("sig"),
							},
						},
						"sig_provider": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("SIG Tunnel Provider").AddStringEnumDescription("secure-internet-gateway-umbrella", "secure-internet-gateway-zscaler", "secure-internet-gateway-other").AddDefaultValueDescription("secure-internet-gateway-umbrella").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("secure-internet-gateway-umbrella", "secure-internet-gateway-zscaler", "secure-internet-gateway-other"),
							},
						},
						"tunnel_dc_preference": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("SIG Tunnel Data Center").AddStringEnumDescription("primary-dc", "secondary-dc").AddDefaultValueDescription("primary-dc").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("primary-dc", "secondary-dc"),
							},
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
						"dead_peer_detection_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE keepalive interval (seconds)").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("10").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65535),
							},
						},
						"dead_peer_detection_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dead_peer_detection_retries": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE keepalive retries").AddIntegerRangeDescription(0, 255).AddDefaultValueDescription("3").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 255),
							},
						},
						"dead_peer_detection_retries_variable": schema.StringAttribute{
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
						"ike_pre_shared_key": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use preshared key to authenticate IKE peer").String,
							Optional:            true,
						},
						"ike_pre_shared_key_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ike_rekey_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE rekey interval <300..1209600> seconds").AddIntegerRangeDescription(300, 1209600).AddDefaultValueDescription("14400").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(300, 1209600),
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
							MarkdownDescription: helpers.NewAttributeDescription("IKE Diffie Hellman Groups").AddStringEnumDescription("2", "14", "15", "16").AddDefaultValueDescription("14").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2", "14", "15", "16"),
							},
						},
						"ike_group_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ike_pre_shared_key_dynamic": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Use preshared key to authenticate IKE peer").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"ike_pre_shared_key_local_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE ID for the local endpoint. Input IPv4 address, domain name, or email address").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 63),
							},
						},
						"ike_pre_shared_key_local_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ike_pre_shared_key_remote_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IKE ID for the remote endpoint. Input IPv4 address, domain name, or email address").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 63),
							},
						},
						"ike_pre_shared_key_remote_id_variable": schema.StringAttribute{
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
							MarkdownDescription: helpers.NewAttributeDescription("Replay window size 32..8192 (must be a power of 2)").AddIntegerRangeDescription(64, 4096).AddDefaultValueDescription("512").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(64, 4096),
							},
						},
						"ipsec_replay_window_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ipsec_ciphersuite": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPsec(ESP) encryption and integrity protocol").AddStringEnumDescription("aes256-cbc-sha1", "aes256-cbc-sha384", "aes256-cbc-sha256", "aes256-cbc-sha512", "aes256-gcm", "null-sha1", "null-sha384", "null-sha256", "null-sha512").AddDefaultValueDescription("aes256-gcm").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("aes256-cbc-sha1", "aes256-cbc-sha384", "aes256-cbc-sha256", "aes256-cbc-sha512", "aes256-gcm", "null-sha1", "null-sha384", "null-sha256", "null-sha512"),
							},
						},
						"ipsec_ciphersuite_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ipsec_perfect_forward_secrecy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPsec perfect forward secrecy settings").AddStringEnumDescription("group-2", "group-14", "group-15", "group-16", "none").AddDefaultValueDescription("none").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("group-2", "group-14", "group-15", "group-16", "none"),
							},
						},
						"ipsec_perfect_forward_secrecy_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"tracker": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable tracker for this interface").String,
							Optional:            true,
						},
						"track_enable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable/disable SIG tracking").AddDefaultValueDescription("true").String,
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
						"service_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Service Type").AddStringEnumDescription("sig").AddDefaultValueDescription("sig").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("sig"),
							},
						},
						"interface_pairs": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface Pair for active and backup").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"active_interface": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Active Tunnel Interface for SIG").String,
										Optional:            true,
									},
									"active_interface_weight": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Active Tunnel Interface Weight").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("1").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 255),
										},
									},
									"backup_interface": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Backup Tunnel Interface for SIG").String,
										Optional:            true,
									},
									"backup_interface_weight": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Backup Tunnel Interface Weight").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("1").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 255),
										},
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Optional:            true,
									},
								},
							},
						},
						"zscaler_authentication_required": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enforce Authentication").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"zscaler_xff_forward": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("XFF forwarding enabled").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"zscaler_firewall_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Firewall enabled").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"zscaler_ips_control_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable IPS Control").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"zscaler_caution_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable Caution").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"zscaler_primary_data_center": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Custom Primary Datacenter").AddDefaultValueDescription("Auto").String,
							Optional:            true,
						},
						"zscaler_primary_data_center_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"zscaler_secondary_data_center": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Custom Secondary Datacenter").AddDefaultValueDescription("Auto").String,
							Optional:            true,
						},
						"zscaler_secondary_data_center_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"zscaler_surrogate_ip": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable Surrogate IP").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"zscaler_surrogate_idle_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Idle time to disassociation").AddDefaultValueDescription("0").String,
							Optional:            true,
						},
						"zscaler_surrogate_display_time_unit": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Display time unit").AddStringEnumDescription("MINUTE", "HOUR", "DAY").AddDefaultValueDescription("MINUTE").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("MINUTE", "HOUR", "DAY"),
							},
						},
						"zscaler_surrogate_ip_enforce_for_known_browsers": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enforce Surrogate IP for known browsers").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"zscaler_surrogate_refresh_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Refresh time for re-validation of surrogacy in minutes").AddDefaultValueDescription("0").String,
							Optional:            true,
						},
						"zscaler_surrogate_refresh_time_unit": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Refresh Time unit").AddStringEnumDescription("MINUTE", "HOUR", "DAY").AddDefaultValueDescription("MINUTE").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("MINUTE", "HOUR", "DAY"),
							},
						},
						"zscaler_aup_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable Acceptable User Policy").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"zscaler_aup_block_internet_until_accepted": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("For first-time Acceptable User Policy behavior, block Internet access").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"zscaler_aup_force_ssl_inspection": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("For first-time Acceptable User Policy behavior, force SSL inspection").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"zscaler_aup_timeout": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Custom Acceptable User Policy frequency in days").AddDefaultValueDescription("0").String,
							Optional:            true,
						},
						"zscaler_location_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Zscaler location name (optional)").AddDefaultValueDescription("Auto").String,
							Optional:            true,
						},
						"zscaler_location_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"umbrella_primary_data_center": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Umbrella Primary Datacenter").AddDefaultValueDescription("Auto").String,
							Optional:            true,
						},
						"umbrella_primary_data_center_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"umbrella_secondary_data_center": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Umbrella Secondary Datacenter").AddDefaultValueDescription("Auto").String,
							Optional:            true,
						},
						"umbrella_secondary_data_center_variable": schema.StringAttribute{
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
			"tracker_source_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Source IP address for Tracker").String,
				Optional:            true,
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
							},
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"endpoint_api_url": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("API url of endpoint").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(0, 512),
							},
						},
						"endpoint_api_url_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"threshold": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Probe Timeout threshold <100..1000> milliseconds").AddIntegerRangeDescription(100, 1000).AddDefaultValueDescription("300").String,
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
							MarkdownDescription: helpers.NewAttributeDescription("Probe interval <10..600> seconds").AddIntegerRangeDescription(20, 600).AddDefaultValueDescription("60").String,
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
							MarkdownDescription: helpers.NewAttributeDescription("Probe failure multiplier <1..10> failed attempts").AddIntegerRangeDescription(1, 10).AddDefaultValueDescription("3").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 10),
							},
						},
						"multiplier_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"tracker_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("SIG").AddDefaultValueDescription(" SIG").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("SIG"),
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

func (r *CiscoSecureInternetGatewayFeatureTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *CiscoSecureInternetGatewayFeatureTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CiscoSecureInternetGateway

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
func (r *CiscoSecureInternetGatewayFeatureTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CiscoSecureInternetGateway

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
func (r *CiscoSecureInternetGatewayFeatureTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CiscoSecureInternetGateway

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
func (r *CiscoSecureInternetGatewayFeatureTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CiscoSecureInternetGateway

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
func (r *CiscoSecureInternetGatewayFeatureTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import
