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
var _ resource.Resource = &OtherTrustSecProfileParcelResource{}
var _ resource.ResourceWithImportState = &OtherTrustSecProfileParcelResource{}

func NewOtherTrustSecProfileParcelResource() resource.Resource {
	return &OtherTrustSecProfileParcelResource{}
}

type OtherTrustSecProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *OtherTrustSecProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_other_trustsec_feature"
}

func (r *OtherTrustSecProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Other TrustSec Feature.").AddMinimumVersionDescription("20.18.0").String,

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
			"device_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Specify the TrustSec Network Access Device ID, should be same as mentioned in the Identity Services Engine (upto 32 char)").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"device_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"device_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the password for the device").AddMutualExclusivityDescription("Only one of `device_password`, `device_password_variable` and `device_password_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `device_password_wo` together with `device_password_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 24),
				},
			},
			"device_password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"device_password_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the password for the device").AddMutualExclusivityDescription("Only one of `device_password`, `device_password_variable` and `device_password_wo` can be set.").String,
				Sensitive:           true,
				WriteOnly:           true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 24),
				},
			},
			"device_password_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `device_password_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"device_sgt": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Local device security group <2..65519>").AddIntegerRangeDescription(2, 65519).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(2, 65519),
				},
			},
			"device_sgt_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"enable_enforcement": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Role-based Access Control enforcement").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"enable_enforcement_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"enable_sxp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable CTS SXP support").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"listener_hold_time_min": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Minimum allowed hold-time for listener in seconds <1..65534>").AddIntegerRangeDescription(1, 65534).AddDefaultValueDescription("90").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65534),
				},
			},
			"listener_hold_time_min_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"listener_hold_time_max": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Maximum allowed hold-time for listener in seconds <1..65534>").AddIntegerRangeDescription(1, 65534).AddDefaultValueDescription("180").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65534),
				},
			},
			"listener_hold_time_max_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"speaker_hold_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Speaker hold-time in seconds <1..65534>").AddIntegerRangeDescription(1, 65534).AddDefaultValueDescription("120").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65534),
				},
			},
			"speaker_hold_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_default_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure SXP default password").AddMutualExclusivityDescription("Only one of `sxp_default_password`, `sxp_default_password_variable` and `sxp_default_password_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `sxp_default_password_wo` together with `sxp_default_password_wo_version`, which keeps it out of state.").String,
				Sensitive:           true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 24),
				},
			},
			"sxp_default_password_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_default_password_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure SXP default password").AddMutualExclusivityDescription("Only one of `sxp_default_password`, `sxp_default_password_variable` and `sxp_default_password_wo` can be set.").String,
				Sensitive:           true,
				WriteOnly:           true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 24),
				},
			},
			"sxp_default_password_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `sxp_default_password_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
				Optional:            true,
			},
			"sxp_key_chain": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure SXP key-chain").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(256),
				},
			},
			"sxp_key_chain_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_log_binding_changes": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enables logging for IP-to-SGT binding changes").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"sxp_log_binding_changes_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_reconciliation_period": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the SXP reconciliation period in seconds <0..64000>").AddIntegerRangeDescription(0, 64000).AddDefaultValueDescription("120").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtMost(64000),
				},
			},
			"sxp_reconciliation_period_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_retry_period": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure Retry period for SXP connection in seconds <0..64000>").AddIntegerRangeDescription(0, 64000).AddDefaultValueDescription("120").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtMost(64000),
				},
			},
			"sxp_retry_period_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_source_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SXP Source IP").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)\.){3}(25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)$`), ""),
				},
			},
			"sxp_source_ip_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"sxp_connections": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure SXP Connections").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"peer_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure SXP Peer IP address (IPv4)").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)\.){3}(25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)$`), ""),
							},
						},
						"peer_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure SXP Source IP address (IPv4)").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^((25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)\.){3}(25[0-5]|2[0-4]\d|1\d\d|[1-9]\d|\d)$`), ""),
							},
						},
						"source_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"preshared_key": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Define Preshared Key type").AddStringEnumDescription("password", "none", "key chain").AddDefaultValueDescription("none").AddMutualExclusivityDescription("Only one of `preshared_key` and `preshared_key_wo` can be set.").AddCoexistenceNote("This attribute stores the secret in Terraform state. Prefer `preshared_key_wo` together with `preshared_key_wo_version`, which keeps it out of state.").String,
							Sensitive:           true,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("password", "none", "key chain"),
							},
						},
						"preshared_key_wo": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Define Preshared Key type").AddStringEnumDescription("password", "none", "key chain").AddDefaultValueDescription("none").AddMutualExclusivityDescription("Only one of `preshared_key` and `preshared_key_wo` can be set.").String,
							Sensitive:           true,
							WriteOnly:           true,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("password", "none", "key chain"),
							},
						},
						"preshared_key_wo_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Rotation trigger for `preshared_key_wo`. Increment this integer whenever the write-only value changes so that Terraform sends the new secret. The value is stored in state; the secret is not.").String,
							Optional:            true,
						},
						"mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Define Mode of connection").AddStringEnumDescription("local", "peer").AddDefaultValueDescription("local").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("local", "peer"),
							},
						},
						"mode_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Define Role of a device <speaker/listener/both>").AddStringEnumDescription("speaker", "listener", "both").AddDefaultValueDescription("speaker").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("speaker", "listener", "both"),
							},
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Connection VPN (VRF) ID").AddIntegerRangeDescription(0, 65527).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(65527),
							},
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"min_hold_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Connection Minimum hold time <0..65535>, Attribute conditional on `mode_type` not equal to `both`").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(65535),
							},
						},
						"min_hold_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on `mode_type` not equal to `both`").String,
							Optional:            true,
						},
						"max_hold_time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure Connection Maximum hold time <0..65535>, Attribute conditional on (`mode` equal to `peer` and `mode_type` equal to `speaker`) or (`mode` equal to `local` and `mode_type` equal to `listener`)").AddIntegerRangeDescription(0, 65535).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(65535),
							},
						},
						"max_hold_time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name, Attribute conditional on (`mode` equal to `peer` and `mode_type` equal to `speaker`) or (`mode` equal to `local` and `mode_type` equal to `listener`)").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *OtherTrustSecProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

var _ resource.ResourceWithValidateConfig = &OtherTrustSecProfileParcelResource{}

// ValidateConfig enforces the rules between a secret attribute, its "_variable"
// sibling and its write-only twin.
//
// All diagnostics here are resource-scoped (AddError/AddWarning, never
// AddAttributeError) and the schema deliberately carries no
// ConflictsWith/ExactlyOneOf/AlsoRequires validators. Terraform renders an
// attribute-scoped diagnostic together with the offending raw configuration line,
// which would print the secret into plan output and CI logs. Sensitive: true does
// not suppress that.
func (r *OtherTrustSecProfileParcelResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {

	var vDevicePassword types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("device_password"), &vDevicePassword)...)
	var vDevicePasswordVariable types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("device_password_variable"), &vDevicePasswordVariable)...)
	var vDevicePasswordWo types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("device_password_wo"), &vDevicePasswordWo)...)
	var vDevicePasswordWoVersion types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("device_password_wo_version"), &vDevicePasswordWoVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An unknown value is one that is not resolved yet at validate time, for example an
	// ephemeral resource attribute or a random_password reference. Treat it as "not
	// decided" rather than "set", otherwise every such configuration trips a false error.
	if !vDevicePassword.IsUnknown() && !vDevicePasswordWo.IsUnknown() && !vDevicePassword.IsNull() && !vDevicePasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `device_password` and `device_password_wo` can be set.",
		)
	}
	if !vDevicePasswordVariable.IsUnknown() && !vDevicePasswordWo.IsUnknown() && !vDevicePasswordVariable.IsNull() && !vDevicePasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"`device_password_variable` and `device_password_wo` cannot both be set. `device_password_variable` defers the value to a device variable, so the write-only value would never be sent.",
		)
	}
	if !vDevicePassword.IsUnknown() && !vDevicePasswordVariable.IsUnknown() && !vDevicePassword.IsNull() && !vDevicePasswordVariable.IsNull() {
		resp.Diagnostics.AddWarning(
			"Conflicting Attribute Combination",
			"Both `device_password` and `device_password_variable` are set. `device_password_variable` takes precedence, so the value of `device_password` is not sent to the device - but it is still stored in Terraform state. Set only one of them.",
		)
	}
	if !vDevicePasswordWo.IsUnknown() && !vDevicePasswordWoVersion.IsUnknown() && !vDevicePasswordWo.IsNull() && vDevicePasswordWoVersion.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`device_password_wo_version` must be set when `device_password_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !vDevicePasswordWoVersion.IsUnknown() && !vDevicePasswordWo.IsUnknown() && !vDevicePasswordWoVersion.IsNull() && vDevicePasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`device_password_wo` must be set when `device_password_wo_version` is used.",
		)
	}

	var vSxpDefaultPassword types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("sxp_default_password"), &vSxpDefaultPassword)...)
	var vSxpDefaultPasswordVariable types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("sxp_default_password_variable"), &vSxpDefaultPasswordVariable)...)
	var vSxpDefaultPasswordWo types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("sxp_default_password_wo"), &vSxpDefaultPasswordWo)...)
	var vSxpDefaultPasswordWoVersion types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("sxp_default_password_wo_version"), &vSxpDefaultPasswordWoVersion)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An unknown value is one that is not resolved yet at validate time, for example an
	// ephemeral resource attribute or a random_password reference. Treat it as "not
	// decided" rather than "set", otherwise every such configuration trips a false error.
	if !vSxpDefaultPassword.IsUnknown() && !vSxpDefaultPasswordWo.IsUnknown() && !vSxpDefaultPassword.IsNull() && !vSxpDefaultPasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Only one of `sxp_default_password` and `sxp_default_password_wo` can be set.",
		)
	}
	if !vSxpDefaultPasswordVariable.IsUnknown() && !vSxpDefaultPasswordWo.IsUnknown() && !vSxpDefaultPasswordVariable.IsNull() && !vSxpDefaultPasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"`sxp_default_password_variable` and `sxp_default_password_wo` cannot both be set. `sxp_default_password_variable` defers the value to a device variable, so the write-only value would never be sent.",
		)
	}
	if !vSxpDefaultPassword.IsUnknown() && !vSxpDefaultPasswordVariable.IsUnknown() && !vSxpDefaultPassword.IsNull() && !vSxpDefaultPasswordVariable.IsNull() {
		resp.Diagnostics.AddWarning(
			"Conflicting Attribute Combination",
			"Both `sxp_default_password` and `sxp_default_password_variable` are set. `sxp_default_password_variable` takes precedence, so the value of `sxp_default_password` is not sent to the device - but it is still stored in Terraform state. Set only one of them.",
		)
	}
	if !vSxpDefaultPasswordWo.IsUnknown() && !vSxpDefaultPasswordWoVersion.IsUnknown() && !vSxpDefaultPasswordWo.IsNull() && vSxpDefaultPasswordWoVersion.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`sxp_default_password_wo_version` must be set when `sxp_default_password_wo` is used. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.",
		)
	}
	if !vSxpDefaultPasswordWoVersion.IsUnknown() && !vSxpDefaultPasswordWo.IsUnknown() && !vSxpDefaultPasswordWoVersion.IsNull() && vSxpDefaultPasswordWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"`sxp_default_password_wo` must be set when `sxp_default_password_wo_version` is used.",
		)
	}

	{
		// Secrets nested in "sxp_connections" are validated per element. A list that cannot be
		// read as a whole - because it is still unknown at validate time - is skipped rather
		// than reported, since there is nothing to check yet.
		var cfgSxpConnections []OtherTrustSecSxpConnections
		if diags := req.Config.GetAttribute(ctx, path.Root("sxp_connections"), &cfgSxpConnections); !diags.HasError() {
			for i := range cfgSxpConnections {
				el := cfgSxpConnections[i]

				// An unknown value is one that is not resolved yet at validate time, for example an
				// ephemeral resource attribute or a random_password reference. Treat it as "not
				// decided" rather than "set", otherwise every such configuration trips a false error.
				if !el.PresharedKey.IsUnknown() && !el.PresharedKeyWo.IsUnknown() && !el.PresharedKey.IsNull() && !el.PresharedKeyWo.IsNull() {
					resp.Diagnostics.AddError(
						"Invalid Attribute Combination",
						fmt.Sprintf("Only one of `preshared_key` and `preshared_key_wo` can be set in `sxp_connections` element %d.", i),
					)
				}
				if !el.PresharedKeyWo.IsUnknown() && !el.PresharedKeyWoVersion.IsUnknown() && !el.PresharedKeyWo.IsNull() && el.PresharedKeyWoVersion.IsNull() {
					resp.Diagnostics.AddError(
						"Missing Required Attribute",
						fmt.Sprintf("`preshared_key_wo_version` must be set when `preshared_key_wo` is used in `sxp_connections` element %d. The write-only value is not stored in state, so Terraform can only detect a change to it through the version.", i),
					)
				}
				if !el.PresharedKeyWoVersion.IsUnknown() && !el.PresharedKeyWo.IsUnknown() && !el.PresharedKeyWoVersion.IsNull() && el.PresharedKeyWo.IsNull() {
					resp.Diagnostics.AddError(
						"Missing Required Attribute",
						fmt.Sprintf("`preshared_key_wo` must be set when `preshared_key_wo_version` is used in `sxp_connections` element %d.", i),
					)
				}
			}
		}
	}
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *OtherTrustSecProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan OtherTrustSec

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "device_password_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because a parcel PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("device_password_wo"), &plan.DevicePasswordWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "sxp_default_password_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because a parcel PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("sxp_default_password_wo"), &plan.SxpDefaultPasswordWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only values in list "sxp_connections" are not stored in plan or state; read the
	// parent list from the configuration and copy them into the plan element by element.
	// The positional join is sound because the parent is an ordered List - write_only_tf
	// is rejected at generate time anywhere under a Set.
	{
		var cfgSxpConnections []OtherTrustSecSxpConnections
		resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("sxp_connections"), &cfgSxpConnections)...)
		if resp.Diagnostics.HasError() {
			return
		}
		for i := range plan.SxpConnections {
			if i < len(cfgSxpConnections) {
				plan.SxpConnections[i].PresharedKeyWo = cfgSxpConnections[i].PresharedKeyWo
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
func (r *OtherTrustSecProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state OtherTrustSec

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
func (r *OtherTrustSecProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state OtherTrustSec

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "device_password_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because a parcel PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("device_password_wo"), &plan.DevicePasswordWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only value "sxp_default_password_wo" is not stored in plan or state; read it from the
	// configuration so it can be sent to the API. Read unconditionally on every
	// Update, because a parcel PUT replaces the whole object.
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("sxp_default_password_wo"), &plan.SxpDefaultPasswordWo)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write-only values in list "sxp_connections" are not stored in plan or state; read the
	// parent list from the configuration and copy them into the plan element by element.
	// The positional join is sound because the parent is an ordered List - write_only_tf
	// is rejected at generate time anywhere under a Set.
	{
		var cfgSxpConnections []OtherTrustSecSxpConnections
		resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("sxp_connections"), &cfgSxpConnections)...)
		if resp.Diagnostics.HasError() {
			return
		}
		for i := range plan.SxpConnections {
			if i < len(cfgSxpConnections) {
				plan.SxpConnections[i].PresharedKeyWo = cfgSxpConnections[i].PresharedKeyWo
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
func (r *OtherTrustSecProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OtherTrustSec

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
func (r *OtherTrustSecProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "other_trustsec_feature_id" + ",feature_profile_id"
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
