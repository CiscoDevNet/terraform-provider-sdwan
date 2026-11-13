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
	"github.com/hashicorp/go-version"
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
var _ resource.Resource = &SystemNTPProfileParcelResource{}
var _ resource.ResourceWithImportState = &SystemNTPProfileParcelResource{}

func NewSystemNTPProfileParcelResource() resource.Resource {
	return &SystemNTPProfileParcelResource{}
}

type SystemNTPProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *SystemNTPProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_ntp_feature"
}

func (r *SystemNTPProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a System NTP Feature.").AddMinimumVersionDescription("20.15.0").String,

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
			"servers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure NTP servers").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"hostname_ip_address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set hostname or IP address of server").String,
							Optional:            true,
						},
						"hostname_ip_address_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"authentication_key": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set authentication key for the server").AddIntegerRangeDescription(1, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"authentication_key_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vpn": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set VPN in which NTP server is located").AddIntegerRangeDescription(0, 65530).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(65530),
							},
						},
						"vpn_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"ntp_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set NTP version").AddIntegerRangeDescription(1, 4).AddDefaultValueDescription("4").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4),
							},
						},
						"ntp_version_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set interface to use to reach NTP server").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(3, 32),
								stringvalidator.RegexMatches(regexp.MustCompile(`(ATM|ATM-ACR|AppGigabitEthernet|AppNav-Compress|AppNav-UnCompress|Async|BD-VIF|BDI|CEM|CEM-ACR|Cellular|Dialer|Embedded-Service-Engine|Ethernet|Ethernet-Internal|FastEthernet|FiftyGigabitEthernet|FiveGigabitEthernet|FortyGigabitEthernet|FourHundredGigE|GMPLS|GigabitEthernet|Group-Async|HundredGigE|L2LISP|LISP|Loopback|MFR|Multilink|Port-channel|SM|Serial|Service-Engine|TenGigabitEthernet|Tunnel|TwentyFiveGigE|TwentyFiveGigabitEthernet|TwoGigabitEthernet|TwoHundredGigE|Vif|Virtual-PPP|Virtual-Template|VirtualPortGroup|Vlan|Wlan-GigabitEthernet|nat64|nat66|ntp|nve|ospfv3|overlay|pseudowire|ucse|vasileft|vasiright|vmi)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
							},
						},
						"source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"prefer_this_ntp_server": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prefer this NTP server").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"prefer_this_ntp_server_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
			"authentication_keys": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set MD5 authentication key").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Authentication key ID").AddIntegerRangeDescription(1, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 4294967295),
							},
						},
						"key_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"md5_value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enter cleartext or AES-encrypted MD5 authentication key [Note: Catalyst SD-WAN Manager will encrypt this field before saving. Cleartext strings will not be returned back to the user in GET responses for sensitive fields.]").AddMutualExclusivityDescription("Only one of `md5_value`, `md5_value_variable` and `md5_value_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `md5_value_wo` together with `md5_value_wo_version`, which keeps it out of state.").String,
							Sensitive:           true,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
						},
						"md5_value_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"md5_value_wo": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enter cleartext or AES-encrypted MD5 authentication key [Note: Catalyst SD-WAN Manager will encrypt this field before saving. Cleartext strings will not be returned back to the user in GET responses for sensitive fields.]").AddMutualExclusivityDescription("Only one of `md5_value`, `md5_value_variable` and `md5_value_wo` can be set.").String,
							Sensitive:           true,
							WriteOnly:           true,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
						},
						"md5_value_wo_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `md5_value_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
							Optional:            true,
						},
						"hmac_sha2_value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("HMAC-SHA2-256 (digest length = 256 bits, key length = [1-32] bytes), Attribute conditional on SD-WAN Manager version `26.1.1` or higher").AddMutualExclusivityDescription("Only one of `hmac_sha2_value`, `hmac_sha2_value_variable` and `hmac_sha2_value_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `hmac_sha2_value_wo` together with `hmac_sha2_value_wo_version`, which keeps it out of state.").String,
							Sensitive:           true,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^\S{1,32}$`), ""),
							},
						},
						"hmac_sha2_value_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on SD-WAN Manager version `26.1.1` or higher").String,
							Optional:            true,
						},
						"hmac_sha2_value_wo": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("HMAC-SHA2-256 (digest length = 256 bits, key length = [1-32] bytes), Attribute conditional on SD-WAN Manager version `26.1.1` or higher").AddMutualExclusivityDescription("Only one of `hmac_sha2_value`, `hmac_sha2_value_variable` and `hmac_sha2_value_wo` can be set.").String,
							Sensitive:           true,
							WriteOnly:           true,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^\S{1,32}$`), ""),
							},
						},
						"hmac_sha2_value_wo_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `hmac_sha2_value_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
							Optional:            true,
						},
						"cmac_aes128_value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("CMAC-AES-128 (digest length = 128 bits, key length = [16 or 32] bytes), Attribute conditional on SD-WAN Manager version `26.1.1` or higher").AddMutualExclusivityDescription("Only one of `cmac_aes128_value`, `cmac_aes128_value_variable` and `cmac_aes128_value_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `cmac_aes128_value_wo` together with `cmac_aes128_value_wo_version`, which keeps it out of state.").String,
							Sensitive:           true,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(16, 32),
							},
						},
						"cmac_aes128_value_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on SD-WAN Manager version `26.1.1` or higher").String,
							Optional:            true,
						},
						"cmac_aes128_value_wo": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("CMAC-AES-128 (digest length = 128 bits, key length = [16 or 32] bytes), Attribute conditional on SD-WAN Manager version `26.1.1` or higher").AddMutualExclusivityDescription("Only one of `cmac_aes128_value`, `cmac_aes128_value_variable` and `cmac_aes128_value_wo` can be set.").String,
							Sensitive:           true,
							WriteOnly:           true,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(16, 32),
							},
						},
						"cmac_aes128_value_wo_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `cmac_aes128_value_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
							Optional:            true,
						},
					},
				},
			},
			"trusted_keys": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Designate authentication key as trustworthy").String,
				ElementType:         types.Int64Type,
				Optional:            true,
			},
			"trusted_keys_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"authoritative_ntp_server": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable device as NTP Leader").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"authoritative_ntp_server_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"stratum": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable device as NTP Leader").AddIntegerRangeDescription(1, 15).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 15),
				},
			},
			"stratum_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"source_interface": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable device as NTP Leader").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(3, 32),
					stringvalidator.RegexMatches(regexp.MustCompile(`(ATM|ATM-ACR|AppGigabitEthernet|AppNav-Compress|AppNav-UnCompress|Async|BD-VIF|BDI|CEM|CEM-ACR|Cellular|Dialer|Embedded-Service-Engine|Ethernet|Ethernet-Internal|FastEthernet|FiftyGigabitEthernet|FiveGigabitEthernet|FortyGigabitEthernet|FourHundredGigE|GMPLS|GigabitEthernet|Group-Async|HundredGigE|L2LISP|LISP|Loopback|MFR|Multilink|Port-channel|SM|Serial|Service-Engine|TenGigabitEthernet|Tunnel|TwentyFiveGigE|TwentyFiveGigabitEthernet|TwoGigabitEthernet|TwoHundredGigE|Vif|Virtual-PPP|Virtual-Template|VirtualPortGroup|Vlan|Wlan-GigabitEthernet|nat64|nat66|ntp|nve|ospfv3|overlay|pseudowire|ucse|vasileft|vasiright|vmi)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
				},
			},
			"source_interface_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
		},
	}
}

func (r *SystemNTPProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

var _ resource.ResourceWithValidateConfig = &SystemNTPProfileParcelResource{}

// ValidateConfig enforces the rules between a secret attribute, its "_variable"
// sibling and its write-only twin.
//
// All diagnostics here are resource-scoped (AddError/AddWarning, never
// AddAttributeError) and the schema deliberately carries no
// ConflictsWith/ExactlyOneOf/AlsoRequires validators. Terraform renders an
// attribute-scoped diagnostic together with the offending raw configuration line,
// which would print the secret into plan output and CI logs. Sensitive: true does
// not suppress that.
func (r *SystemNTPProfileParcelResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {

	{
		// Secrets nested in "authentication_keys" are validated per element. A list that cannot be
		// read as a whole - because it is still unknown at validate time - is skipped rather
		// than reported, since there is nothing to check yet.
		var cfgAuthenticationKeys []SystemNTPAuthenticationKeys
		if diags := req.Config.GetAttribute(ctx, path.Root("authentication_keys"), &cfgAuthenticationKeys); !diags.HasError() {
			for i := range cfgAuthenticationKeys {
				el := cfgAuthenticationKeys[i]

				// An unknown value is one that is not resolved yet at validate time, for example an
				// ephemeral resource attribute or a random_password reference. Treat it as "not
				// decided" rather than "set", otherwise every such configuration trips a false error.
				if !el.Md5Value.IsUnknown() && !el.Md5ValueWo.IsUnknown() && !el.Md5Value.IsNull() && !el.Md5ValueWo.IsNull() {
					resp.Diagnostics.AddError(
						"Invalid Attribute Combination",
						fmt.Sprintf("Only one of `md5_value` and `md5_value_wo` can be set in `authentication_keys` element %d.", i),
					)
				}
				if !el.Md5ValueVariable.IsUnknown() && !el.Md5ValueWo.IsUnknown() && !el.Md5ValueVariable.IsNull() && !el.Md5ValueWo.IsNull() {
					resp.Diagnostics.AddError(
						"Invalid Attribute Combination",
						fmt.Sprintf("`md5_value_variable` and `md5_value_wo` cannot both be set in `authentication_keys` element %d. `md5_value_variable` defers the value to a device variable, so the write-only value would never be sent.", i),
					)
				}
				if !el.Md5Value.IsUnknown() && !el.Md5ValueVariable.IsUnknown() && !el.Md5Value.IsNull() && !el.Md5ValueVariable.IsNull() {
					resp.Diagnostics.AddWarning(
						"Conflicting Attribute Combination",
						fmt.Sprintf("Both `md5_value` and `md5_value_variable` are set in `authentication_keys` element %d. `md5_value_variable` takes precedence, so the value of `md5_value` is not sent to the device - but it is still stored in Terraform state. Set only one of them.", i),
					)
				}
				if !el.Md5ValueWo.IsUnknown() && !el.Md5ValueWoVersion.IsUnknown() && !el.Md5ValueWo.IsNull() && el.Md5ValueWoVersion.IsNull() {
					resp.Diagnostics.AddError(
						"Missing Required Attribute",
						fmt.Sprintf("`md5_value_wo_version` must be set when `md5_value_wo` is used in `authentication_keys` element %d. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.", i),
					)
				}
				if !el.Md5ValueWoVersion.IsUnknown() && !el.Md5ValueWo.IsUnknown() && !el.Md5ValueWoVersion.IsNull() && el.Md5ValueWo.IsNull() {
					resp.Diagnostics.AddError(
						"Missing Required Attribute",
						fmt.Sprintf("`md5_value_wo` must be set when `md5_value_wo_version` is used in `authentication_keys` element %d.", i),
					)
				}

				// An unknown value is one that is not resolved yet at validate time, for example an
				// ephemeral resource attribute or a random_password reference. Treat it as "not
				// decided" rather than "set", otherwise every such configuration trips a false error.
				if !el.HmacSha2Value.IsUnknown() && !el.HmacSha2ValueWo.IsUnknown() && !el.HmacSha2Value.IsNull() && !el.HmacSha2ValueWo.IsNull() {
					resp.Diagnostics.AddError(
						"Invalid Attribute Combination",
						fmt.Sprintf("Only one of `hmac_sha2_value` and `hmac_sha2_value_wo` can be set in `authentication_keys` element %d.", i),
					)
				}
				if !el.HmacSha2ValueVariable.IsUnknown() && !el.HmacSha2ValueWo.IsUnknown() && !el.HmacSha2ValueVariable.IsNull() && !el.HmacSha2ValueWo.IsNull() {
					resp.Diagnostics.AddError(
						"Invalid Attribute Combination",
						fmt.Sprintf("`hmac_sha2_value_variable` and `hmac_sha2_value_wo` cannot both be set in `authentication_keys` element %d. `hmac_sha2_value_variable` defers the value to a device variable, so the write-only value would never be sent.", i),
					)
				}
				if !el.HmacSha2Value.IsUnknown() && !el.HmacSha2ValueVariable.IsUnknown() && !el.HmacSha2Value.IsNull() && !el.HmacSha2ValueVariable.IsNull() {
					resp.Diagnostics.AddWarning(
						"Conflicting Attribute Combination",
						fmt.Sprintf("Both `hmac_sha2_value` and `hmac_sha2_value_variable` are set in `authentication_keys` element %d. `hmac_sha2_value_variable` takes precedence, so the value of `hmac_sha2_value` is not sent to the device - but it is still stored in Terraform state. Set only one of them.", i),
					)
				}
				if !el.HmacSha2ValueWo.IsUnknown() && !el.HmacSha2ValueWoVersion.IsUnknown() && !el.HmacSha2ValueWo.IsNull() && el.HmacSha2ValueWoVersion.IsNull() {
					resp.Diagnostics.AddError(
						"Missing Required Attribute",
						fmt.Sprintf("`hmac_sha2_value_wo_version` must be set when `hmac_sha2_value_wo` is used in `authentication_keys` element %d. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.", i),
					)
				}
				if !el.HmacSha2ValueWoVersion.IsUnknown() && !el.HmacSha2ValueWo.IsUnknown() && !el.HmacSha2ValueWoVersion.IsNull() && el.HmacSha2ValueWo.IsNull() {
					resp.Diagnostics.AddError(
						"Missing Required Attribute",
						fmt.Sprintf("`hmac_sha2_value_wo` must be set when `hmac_sha2_value_wo_version` is used in `authentication_keys` element %d.", i),
					)
				}

				// An unknown value is one that is not resolved yet at validate time, for example an
				// ephemeral resource attribute or a random_password reference. Treat it as "not
				// decided" rather than "set", otherwise every such configuration trips a false error.
				if !el.CmacAes128Value.IsUnknown() && !el.CmacAes128ValueWo.IsUnknown() && !el.CmacAes128Value.IsNull() && !el.CmacAes128ValueWo.IsNull() {
					resp.Diagnostics.AddError(
						"Invalid Attribute Combination",
						fmt.Sprintf("Only one of `cmac_aes128_value` and `cmac_aes128_value_wo` can be set in `authentication_keys` element %d.", i),
					)
				}
				if !el.CmacAes128ValueVariable.IsUnknown() && !el.CmacAes128ValueWo.IsUnknown() && !el.CmacAes128ValueVariable.IsNull() && !el.CmacAes128ValueWo.IsNull() {
					resp.Diagnostics.AddError(
						"Invalid Attribute Combination",
						fmt.Sprintf("`cmac_aes128_value_variable` and `cmac_aes128_value_wo` cannot both be set in `authentication_keys` element %d. `cmac_aes128_value_variable` defers the value to a device variable, so the write-only value would never be sent.", i),
					)
				}
				if !el.CmacAes128Value.IsUnknown() && !el.CmacAes128ValueVariable.IsUnknown() && !el.CmacAes128Value.IsNull() && !el.CmacAes128ValueVariable.IsNull() {
					resp.Diagnostics.AddWarning(
						"Conflicting Attribute Combination",
						fmt.Sprintf("Both `cmac_aes128_value` and `cmac_aes128_value_variable` are set in `authentication_keys` element %d. `cmac_aes128_value_variable` takes precedence, so the value of `cmac_aes128_value` is not sent to the device - but it is still stored in Terraform state. Set only one of them.", i),
					)
				}
				if !el.CmacAes128ValueWo.IsUnknown() && !el.CmacAes128ValueWoVersion.IsUnknown() && !el.CmacAes128ValueWo.IsNull() && el.CmacAes128ValueWoVersion.IsNull() {
					resp.Diagnostics.AddError(
						"Missing Required Attribute",
						fmt.Sprintf("`cmac_aes128_value_wo_version` must be set when `cmac_aes128_value_wo` is used in `authentication_keys` element %d. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.", i),
					)
				}
				if !el.CmacAes128ValueWoVersion.IsUnknown() && !el.CmacAes128ValueWo.IsUnknown() && !el.CmacAes128ValueWoVersion.IsNull() && el.CmacAes128ValueWo.IsNull() {
					resp.Diagnostics.AddError(
						"Missing Required Attribute",
						fmt.Sprintf("`cmac_aes128_value_wo` must be set when `cmac_aes128_value_wo_version` is used in `authentication_keys` element %d.", i),
					)
				}
			}
		}
	}
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *SystemNTPProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemNTP

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only values in list "authentication_keys" are not stored in plan or state; read the
	// parent list from the configuration and copy them into the plan element by element.
	// The positional join is sound because the parent is an ordered List - write_only_tf
	// is rejected at generate time anywhere under a Set.
	{
		var cfgAuthenticationKeys []SystemNTPAuthenticationKeys
		resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("authentication_keys"), &cfgAuthenticationKeys)...)
		if resp.Diagnostics.HasError() {
			return
		}
		for i := range plan.AuthenticationKeys {
			if i < len(cfgAuthenticationKeys) {
				plan.AuthenticationKeys[i].Md5ValueWo = cfgAuthenticationKeys[i].Md5ValueWo
				plan.AuthenticationKeys[i].HmacSha2ValueWo = cfgAuthenticationKeys[i].HmacSha2ValueWo
				plan.AuthenticationKeys[i].CmacAes128ValueWo = cfgAuthenticationKeys[i].CmacAes128ValueWo
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object

	ver := version.Must(version.NewVersion(r.client.ManagerVersion))
	body := plan.toBody(ctx, ver)

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
func (r *SystemNTPProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemNTP

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
func (r *SystemNTPProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemNTP

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only values in list "authentication_keys" are not stored in plan or state; read the
	// parent list from the configuration and copy them into the plan element by element.
	// The positional join is sound because the parent is an ordered List - write_only_tf
	// is rejected at generate time anywhere under a Set.
	{
		var cfgAuthenticationKeys []SystemNTPAuthenticationKeys
		resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("authentication_keys"), &cfgAuthenticationKeys)...)
		if resp.Diagnostics.HasError() {
			return
		}
		for i := range plan.AuthenticationKeys {
			if i < len(cfgAuthenticationKeys) {
				plan.AuthenticationKeys[i].Md5ValueWo = cfgAuthenticationKeys[i].Md5ValueWo
				plan.AuthenticationKeys[i].HmacSha2ValueWo = cfgAuthenticationKeys[i].HmacSha2ValueWo
				plan.AuthenticationKeys[i].CmacAes128ValueWo = cfgAuthenticationKeys[i].CmacAes128ValueWo
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

	ver := version.Must(version.NewVersion(r.client.ManagerVersion))
	body := plan.toBody(ctx, ver)
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
func (r *SystemNTPProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemNTP

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
func (r *SystemNTPProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "system_ntp_feature_id" + ",feature_profile_id"
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
