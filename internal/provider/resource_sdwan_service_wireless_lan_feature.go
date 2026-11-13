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
var _ resource.Resource = &ServiceWirelessLANProfileParcelResource{}
var _ resource.ResourceWithImportState = &ServiceWirelessLANProfileParcelResource{}

func NewServiceWirelessLANProfileParcelResource() resource.Resource {
	return &ServiceWirelessLANProfileParcelResource{}
}

type ServiceWirelessLANProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *ServiceWirelessLANProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_wireless_lan_feature"
}

func (r *ServiceWirelessLANProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Service Wireless LAN Feature.").AddMinimumVersionDescription("20.15.0").String,

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
			"enable_24g": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("2.4GHz Enabled").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"enable_24g_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"enable_5g": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("5GHz Enabled").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"enable_5g_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"ssids": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Wi-Fi SSID profile").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ssid_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure wlan SSID").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 64),
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-zA-Z!->@^_]*`), ""),
							},
						},
						"admin_state": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set admin state").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"admin_state_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"broadcast_ssid": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable broadcast SSID").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"broadcast_ssid_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vlan_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set VLAN ID").AddIntegerRangeDescription(1, 4094).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4094),
							},
						},
						"vlan_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"radio_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select radio type").AddStringEnumDescription("24ghz", "5ghz", "all").AddDefaultValueDescription("all").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("24ghz", "5ghz", "all"),
							},
						},
						"radio_type_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"security_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select security type").AddStringEnumDescription("enterprise", "personal", "open").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("enterprise", "personal", "open"),
							},
						},
						"radius_server_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set RADIUS server IP, Attribute conditional on `security_type` equal to `enterprise`").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)\.){3}(25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)$`), ""),
							},
						},
						"radius_server_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `security_type` equal to `enterprise`").String,
							Optional:            true,
						},
						"radius_server_port": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set RADIUS server authentication port, Attribute conditional on `security_type` equal to `enterprise`").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("1812").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 65535),
							},
						},
						"radius_server_port_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `security_type` equal to `enterprise`").String,
							Optional:            true,
						},
						"radius_server_secret": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set RADIUS server shared secret, Attribute conditional on `security_type` equal to `enterprise`").AddMutualExclusivityDescription("Only one of `radius_server_secret`, `radius_server_secret_variable` and `radius_server_secret_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `radius_server_secret_wo` together with `radius_server_secret_wo_version`, which keeps it out of state.").String,
							Sensitive:           true,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(8, 32),
							},
						},
						"radius_server_secret_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `security_type` equal to `enterprise`").String,
							Optional:            true,
						},
						"radius_server_secret_wo": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set RADIUS server shared secret, Attribute conditional on `security_type` equal to `enterprise`").AddMutualExclusivityDescription("Only one of `radius_server_secret`, `radius_server_secret_variable` and `radius_server_secret_wo` can be set.").String,
							Sensitive:           true,
							WriteOnly:           true,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(8, 32),
							},
						},
						"radius_server_secret_wo_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `radius_server_secret_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
							Optional:            true,
						},
						"passphrase": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set passphrase, Attribute conditional on `security_type` equal to `personal`").AddMutualExclusivityDescription("Only one of `passphrase`, `passphrase_variable` and `passphrase_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `passphrase_wo` together with `passphrase_wo_version`, which keeps it out of state.").String,
							Sensitive:           true,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(8, 63),
							},
						},
						"passphrase_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `security_type` equal to `personal`").String,
							Optional:            true,
						},
						"passphrase_wo": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set passphrase, Attribute conditional on `security_type` equal to `personal`").AddMutualExclusivityDescription("Only one of `passphrase`, `passphrase_variable` and `passphrase_wo` can be set.").String,
							Sensitive:           true,
							WriteOnly:           true,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(8, 63),
							},
						},
						"passphrase_wo_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `passphrase_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
							Optional:            true,
						},
						"qos_profile": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Select QoS profile").AddStringEnumDescription("platinum", "gold", "silver", "bronze").AddDefaultValueDescription("silver").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("platinum", "gold", "silver", "bronze"),
							},
						},
						"qos_profile_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"country": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Select country").AddStringEnumDescription("AE", "AR", "AT", "AU", "BA", "BB", "BE", "BG", "BH", "BN", "BO", "BR", "BY", "CA", "CA2", "CH", "CL", "CM", "CN", "CO", "CR", "CY", "CZ", "DE", "DK", "DO", "DZ", "EC", "EE", "EG", "ES", "FI", "FJ", "FR", "GB", "GH", "GI", "GR", "HK", "HR", "HU", "ID", "IE", "IL", "IO", "IN", "IQ", "IS", "IT", "J2", "J4", "JM", "JO", "KE", "KN", "KW", "KZ", "LB", "LI", "LK", "LT", "LU", "LV", "LY", "MA", "MC", "ME", "MK", "MN", "MO", "MT", "MX", "MY", "NL", "NO", "NZ", "OM", "PA", "PE", "PH", "PH2", "PK", "PL", "PR", "PT", "PY", "QA", "RO", "RS", "RU", "SA", "SE", "SG", "SI", "SK", "TH", "TN", "TR", "TW", "UA", "US", "UY", "VE", "VN", "ZA").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("AE", "AR", "AT", "AU", "BA", "BB", "BE", "BG", "BH", "BN", "BO", "BR", "BY", "CA", "CA2", "CH", "CL", "CM", "CN", "CO", "CR", "CY", "CZ", "DE", "DK", "DO", "DZ", "EC", "EE", "EG", "ES", "FI", "FJ", "FR", "GB", "GH", "GI", "GR", "HK", "HR", "HU", "ID", "IE", "IL", "IO", "IN", "IQ", "IS", "IT", "J2", "J4", "JM", "JO", "KE", "KN", "KW", "KZ", "LB", "LI", "LK", "LT", "LU", "LV", "LY", "MA", "MC", "ME", "MK", "MN", "MO", "MT", "MX", "MY", "NL", "NO", "NZ", "OM", "PA", "PE", "PH", "PH2", "PK", "PL", "PR", "PT", "PY", "QA", "RO", "RS", "RU", "SA", "SE", "SG", "SI", "SK", "TH", "TN", "TR", "TW", "UA", "US", "UY", "VE", "VN", "ZA"),
				},
			},
			"country_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set management username").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-zA-Z!->@^_]*`), ""),
				},
			},
			"username_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set management password,the password must contains characters from all of the following classes,lowercase letters,uppercase letters,digits,and special characters. No character in the password can be repeated more than three times consecutively. The password must not be the same as the associated username or the username reversed. The password must not be cisco,ocsic,or any variant obtained by changing the capitalization of the letters in word cisco. In addition,you can't substitute 1,l,or ! for i,0 for o,$ for s.").AddMutualExclusivityDescription("Only one of `password`, `password_variable` and `password_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `password_wo` together with `password_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
			},
			"password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"password_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set management password,the password must contains characters from all of the following classes,lowercase letters,uppercase letters,digits,and special characters. No character in the password can be repeated more than three times consecutively. The password must not be the same as the associated username or the username reversed. The password must not be cisco,ocsic,or any variant obtained by changing the capitalization of the letters in word cisco. In addition,you can't substitute 1,l,or ! for i,0 for o,$ for s.").AddMutualExclusivityDescription("Only one of `password`, `password_variable` and `password_wo` can be set.").String,
				Sensitive:           true,
				WriteOnly:           true,
				Optional:            true,
			},
			"password_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `password_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"me_dynamic_ip_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ME management IP dynamic allocated by DHCP").String,
				Required:            true,
			},
			"me_ipv4_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set mobile express controller address").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)\.){3}(25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)$`), ""),
				},
			},
			"me_ipv4_address_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"me_subnet_mask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set mobile express controller subnet mask").AddStringEnumDescription("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.248.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.248.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0"),
				},
			},
			"me_subnet_mask_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"me_default_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set mobile express default gateway").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)\.){3}(25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)$`), ""),
				},
			},
			"me_default_gateway_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
		},
	}
}

func (r *ServiceWirelessLANProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

var _ resource.ResourceWithValidateConfig = &ServiceWirelessLANProfileParcelResource{}

// ValidateConfig enforces the rules between a secret attribute, its "_variable"
// sibling and its write-only twin.
//
// All diagnostics here are resource-scoped (AddError/AddWarning, never
// AddAttributeError) and the schema deliberately carries no
// ConflictsWith/ExactlyOneOf/AlsoRequires validators. Terraform renders an
// attribute-scoped diagnostic together with the offending raw configuration line,
// which would print the secret into plan output and CI logs. Sensitive: true does
// not suppress that.
func (r *ServiceWirelessLANProfileParcelResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {

	var vPassword types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password"), &vPassword)...)
	var vPasswordVariable types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password_variable"), &vPasswordVariable)...)
	var vPasswordWo types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password_wo"), &vPasswordWo)...)
	var vPasswordWoVersion types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password_wo_version"), &vPasswordWoVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An unknown value is one that is not resolved yet at validate time, for example an
	// ephemeral resource attribute or a random_password reference. Treat it as "not
	// decided" rather than "set", otherwise every such configuration trips a false error.
	if !vPassword.IsUnknown() && !vPasswordWo.IsUnknown() && !vPassword.IsNull() && !vPasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `password` and `password_wo` can be set.",
		)
	}
	if !vPasswordVariable.IsUnknown() && !vPasswordWo.IsUnknown() && !vPasswordVariable.IsNull() && !vPasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"`password_variable` and `password_wo` cannot both be set. `password_variable` defers the value to a device variable, so the write-only value would never be sent.",
		)
	}
	if !vPassword.IsUnknown() && !vPasswordVariable.IsUnknown() && !vPassword.IsNull() && !vPasswordVariable.IsNull() {
		resp.Diagnostics.AddWarning(
			"Conflicting Attribute Combination",
			"Both `password` and `password_variable` are set. `password_variable` takes precedence, so the value of `password` is not sent to the device - but it is still stored in Terraform state. Set only one of them.",
		)
	}
	if !vPasswordWo.IsUnknown() && !vPasswordWoVersion.IsUnknown() && !vPasswordWo.IsNull() && vPasswordWoVersion.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`password_wo_version` must be set when `password_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !vPasswordWoVersion.IsUnknown() && !vPasswordWo.IsUnknown() && !vPasswordWoVersion.IsNull() && vPasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`password_wo` must be set when `password_wo_version` is used.",
		)
	}

	{
		// Secrets nested in "ssids" are validated per element. A list that cannot be
		// read as a whole - because it is still unknown at validate time - is skipped rather
		// than reported, since there is nothing to check yet.
		var cfgSsids []ServiceWirelessLANSsids
		if diags := req.Config.GetAttribute(ctx, path.Root("ssids"), &cfgSsids); !diags.HasError() {
			for i := range cfgSsids {
				el := cfgSsids[i]

				// An unknown value is one that is not resolved yet at validate time, for example an
				// ephemeral resource attribute or a random_password reference. Treat it as "not
				// decided" rather than "set", otherwise every such configuration trips a false error.
				if !el.RadiusServerSecret.IsUnknown() && !el.RadiusServerSecretWo.IsUnknown() && !el.RadiusServerSecret.IsNull() && !el.RadiusServerSecretWo.IsNull() {
					resp.Diagnostics.AddError(
						"Invalid Attribute Combination",
						fmt.Sprintf("Only one of `radius_server_secret` and `radius_server_secret_wo` can be set in `ssids` element %d.", i),
					)
				}
				if !el.RadiusServerSecretVariable.IsUnknown() && !el.RadiusServerSecretWo.IsUnknown() && !el.RadiusServerSecretVariable.IsNull() && !el.RadiusServerSecretWo.IsNull() {
					resp.Diagnostics.AddError(
						"Invalid Attribute Combination",
						fmt.Sprintf("`radius_server_secret_variable` and `radius_server_secret_wo` cannot both be set in `ssids` element %d. `radius_server_secret_variable` defers the value to a device variable, so the write-only value would never be sent.", i),
					)
				}
				if !el.RadiusServerSecret.IsUnknown() && !el.RadiusServerSecretVariable.IsUnknown() && !el.RadiusServerSecret.IsNull() && !el.RadiusServerSecretVariable.IsNull() {
					resp.Diagnostics.AddWarning(
						"Conflicting Attribute Combination",
						fmt.Sprintf("Both `radius_server_secret` and `radius_server_secret_variable` are set in `ssids` element %d. `radius_server_secret_variable` takes precedence, so the value of `radius_server_secret` is not sent to the device - but it is still stored in Terraform state. Set only one of them.", i),
					)
				}
				if !el.RadiusServerSecretWo.IsUnknown() && !el.RadiusServerSecretWoVersion.IsUnknown() && !el.RadiusServerSecretWo.IsNull() && el.RadiusServerSecretWoVersion.IsNull() {
					resp.Diagnostics.AddError(
						"Missing Required Attribute",
						fmt.Sprintf("`radius_server_secret_wo_version` must be set when `radius_server_secret_wo` is used in `ssids` element %d. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.", i),
					)
				}
				if !el.RadiusServerSecretWoVersion.IsUnknown() && !el.RadiusServerSecretWo.IsUnknown() && !el.RadiusServerSecretWoVersion.IsNull() && el.RadiusServerSecretWo.IsNull() {
					resp.Diagnostics.AddError(
						"Missing Required Attribute",
						fmt.Sprintf("`radius_server_secret_wo` must be set when `radius_server_secret_wo_version` is used in `ssids` element %d.", i),
					)
				}

				// An unknown value is one that is not resolved yet at validate time, for example an
				// ephemeral resource attribute or a random_password reference. Treat it as "not
				// decided" rather than "set", otherwise every such configuration trips a false error.
				if !el.Passphrase.IsUnknown() && !el.PassphraseWo.IsUnknown() && !el.Passphrase.IsNull() && !el.PassphraseWo.IsNull() {
					resp.Diagnostics.AddError(
						"Invalid Attribute Combination",
						fmt.Sprintf("Only one of `passphrase` and `passphrase_wo` can be set in `ssids` element %d.", i),
					)
				}
				if !el.PassphraseVariable.IsUnknown() && !el.PassphraseWo.IsUnknown() && !el.PassphraseVariable.IsNull() && !el.PassphraseWo.IsNull() {
					resp.Diagnostics.AddError(
						"Invalid Attribute Combination",
						fmt.Sprintf("`passphrase_variable` and `passphrase_wo` cannot both be set in `ssids` element %d. `passphrase_variable` defers the value to a device variable, so the write-only value would never be sent.", i),
					)
				}
				if !el.Passphrase.IsUnknown() && !el.PassphraseVariable.IsUnknown() && !el.Passphrase.IsNull() && !el.PassphraseVariable.IsNull() {
					resp.Diagnostics.AddWarning(
						"Conflicting Attribute Combination",
						fmt.Sprintf("Both `passphrase` and `passphrase_variable` are set in `ssids` element %d. `passphrase_variable` takes precedence, so the value of `passphrase` is not sent to the device - but it is still stored in Terraform state. Set only one of them.", i),
					)
				}
				if !el.PassphraseWo.IsUnknown() && !el.PassphraseWoVersion.IsUnknown() && !el.PassphraseWo.IsNull() && el.PassphraseWoVersion.IsNull() {
					resp.Diagnostics.AddError(
						"Missing Required Attribute",
						fmt.Sprintf("`passphrase_wo_version` must be set when `passphrase_wo` is used in `ssids` element %d. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.", i),
					)
				}
				if !el.PassphraseWoVersion.IsUnknown() && !el.PassphraseWo.IsUnknown() && !el.PassphraseWoVersion.IsNull() && el.PassphraseWo.IsNull() {
					resp.Diagnostics.AddError(
						"Missing Required Attribute",
						fmt.Sprintf("`passphrase_wo` must be set when `passphrase_wo_version` is used in `ssids` element %d.", i),
					)
				}
			}
		}
	}
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *ServiceWirelessLANProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ServiceWirelessLAN

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "password_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because a parcel PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password_wo"), &plan.PasswordWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only values in list "ssids" are not stored in plan or state; read the
	// parent list from the configuration and copy them into the plan element by element.
	// The positional join is sound because the parent is an ordered List - write_only_tf
	// is rejected at generate time anywhere under a Set.
	{
		var cfgSsids []ServiceWirelessLANSsids
		resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("ssids"), &cfgSsids)...)
		if resp.Diagnostics.HasError() {
			return
		}
		for i := range plan.Ssids {
			if i < len(cfgSsids) {
				plan.Ssids[i].RadiusServerSecretWo = cfgSsids[i].RadiusServerSecretWo
				plan.Ssids[i].PassphraseWo = cfgSsids[i].PassphraseWo
			}
		}
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
func (r *ServiceWirelessLANProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ServiceWirelessLAN

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
func (r *ServiceWirelessLANProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ServiceWirelessLAN

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "password_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because a parcel PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("password_wo"), &plan.PasswordWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only values in list "ssids" are not stored in plan or state; read the
	// parent list from the configuration and copy them into the plan element by element.
	// The positional join is sound because the parent is an ordered List - write_only_tf
	// is rejected at generate time anywhere under a Set.
	{
		var cfgSsids []ServiceWirelessLANSsids
		resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("ssids"), &cfgSsids)...)
		if resp.Diagnostics.HasError() {
			return
		}
		for i := range plan.Ssids {
			if i < len(cfgSsids) {
				plan.Ssids[i].RadiusServerSecretWo = cfgSsids[i].RadiusServerSecretWo
				plan.Ssids[i].PassphraseWo = cfgSsids[i].PassphraseWo
			}
		}
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
func (r *ServiceWirelessLANProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ServiceWirelessLAN

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
func (r *ServiceWirelessLANProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "service_wireless_lan_feature_id" + ",feature_profile_id"
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
