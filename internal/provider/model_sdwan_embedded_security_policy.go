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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type EmbeddedSecurity struct {
	Id                                  types.String               `tfsdk:"id"`
	Version                             types.Int64                `tfsdk:"version"`
	Name                                types.String               `tfsdk:"name"`
	Description                         types.String               `tfsdk:"description"`
	FeatureProfileId                    types.String               `tfsdk:"feature_profile_id"`
	Assembly                            []EmbeddedSecurityAssembly `tfsdk:"assembly"`
	TcpSynFloodLimit                    types.String               `tfsdk:"tcp_syn_flood_limit"`
	MaxIncompleteTcpLimit               types.String               `tfsdk:"max_incomplete_tcp_limit"`
	MaxIncompleteUdpLimit               types.String               `tfsdk:"max_incomplete_udp_limit"`
	MaxIncompleteIcmpLimit              types.String               `tfsdk:"max_incomplete_icmp_limit"`
	AuditTrail                          types.String               `tfsdk:"audit_trail"`
	UnifiedLogging                      types.String               `tfsdk:"unified_logging"`
	SessionReclassifyAllow              types.String               `tfsdk:"session_reclassify_allow"`
	ImcpUnreachableAllow                types.String               `tfsdk:"imcp_unreachable_allow"`
	FailureMode                         types.String               `tfsdk:"failure_mode"`
	SecurityLogging                     types.String               `tfsdk:"security_logging"`
	Nat                                 types.Bool                 `tfsdk:"nat"`
	NatVariable                         types.String               `tfsdk:"nat_variable"`
	DownloadUrlDatabaseOnDevice         types.Bool                 `tfsdk:"download_url_database_on_device"`
	DownloadUrlDatabaseOnDeviceVariable types.String               `tfsdk:"download_url_database_on_device_variable"`
	ResourceProfile                     types.String               `tfsdk:"resource_profile"`
	ResourceProfileVariable             types.String               `tfsdk:"resource_profile_variable"`
}

type EmbeddedSecurityAssembly struct {
	AdvancedInspectionProfilePolicyId types.String                      `tfsdk:"advanced_inspection_profile_policy_id"`
	SslDecryptionProfileId            types.String                      `tfsdk:"ssl_decryption_profile_id"`
	NgfwPolicyId                      types.String                      `tfsdk:"ngfw_policy_id"`
	Entries                           []EmbeddedSecurityAssemblyEntries `tfsdk:"entries"`
}

type EmbeddedSecurityAssemblyEntries struct {
	SourceZoneListId      types.String `tfsdk:"source_zone_list_id"`
	DestinationZoneListId types.String `tfsdk:"destination_zone_list_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data EmbeddedSecurity) getModel() string {
	return "embedded_security"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data EmbeddedSecurity) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/embedded-security/%v/policy", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data EmbeddedSecurity) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {

		for _, item := range data.Assembly {
			itemBody := ""
			if !item.AdvancedInspectionProfilePolicyId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "advancedInspectionProfile.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "advancedInspectionProfile.refId.value", item.AdvancedInspectionProfilePolicyId.ValueString())
				}
			}
			if !item.SslDecryptionProfileId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sslDecryption.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sslDecryption.refId.value", item.SslDecryptionProfileId.ValueString())
				}
			}
			if !item.NgfwPolicyId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ngfirewall.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ngfirewall.refId.value", item.NgfwPolicyId.ValueString())
				}
			}
			if true {

				for _, childItem := range item.Entries {
					itemChildBody := ""
					if !childItem.SourceZoneListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "srcZone.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "srcZone.refId.value", childItem.SourceZoneListId.ValueString())
						}
					}
					if !childItem.DestinationZoneListId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "dstZone.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "dstZone.refId.value", childItem.DestinationZoneListId.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ngfirewall.entries.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"assembly.-1", itemBody)
		}
	}
	if !data.TcpSynFloodLimit.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"settings.tcpSynFloodLimit.optionType", "global")
			body, _ = sjson.Set(body, path+"settings.tcpSynFloodLimit.value", data.TcpSynFloodLimit.ValueString())
		}
	}
	if !data.MaxIncompleteTcpLimit.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"settings.maxIncompleteTcpLimit.optionType", "global")
			body, _ = sjson.Set(body, path+"settings.maxIncompleteTcpLimit.value", data.MaxIncompleteTcpLimit.ValueString())
		}
	}
	if !data.MaxIncompleteUdpLimit.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"settings.maxIncompleteUdpLimit.optionType", "global")
			body, _ = sjson.Set(body, path+"settings.maxIncompleteUdpLimit.value", data.MaxIncompleteUdpLimit.ValueString())
		}
	}
	if !data.MaxIncompleteIcmpLimit.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"settings.maxIncompleteIcmpLimit.optionType", "global")
			body, _ = sjson.Set(body, path+"settings.maxIncompleteIcmpLimit.value", data.MaxIncompleteIcmpLimit.ValueString())
		}
	}
	if !data.AuditTrail.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"settings.auditTrail.optionType", "global")
			body, _ = sjson.Set(body, path+"settings.auditTrail.value", data.AuditTrail.ValueString())
		}
	}
	if !data.UnifiedLogging.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"settings.unifiedLogging.optionType", "global")
			body, _ = sjson.Set(body, path+"settings.unifiedLogging.value", data.UnifiedLogging.ValueString())
		}
	}
	if !data.SessionReclassifyAllow.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"settings.sessionReclassifyAllow.optionType", "global")
			body, _ = sjson.Set(body, path+"settings.sessionReclassifyAllow.value", data.SessionReclassifyAllow.ValueString())
		}
	}
	if !data.ImcpUnreachableAllow.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"settings.icmpUnreachableAllow.optionType", "global")
			body, _ = sjson.Set(body, path+"settings.icmpUnreachableAllow.value", data.ImcpUnreachableAllow.ValueString())
		}
	}
	if !data.FailureMode.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"settings.failureMode.optionType", "global")
			body, _ = sjson.Set(body, path+"settings.failureMode.value", data.FailureMode.ValueString())
		}
	}
	if !data.SecurityLogging.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"settings.securityLogging.optionType", "global")
			body, _ = sjson.Set(body, path+"settings.securityLogging.value", data.SecurityLogging.ValueString())
		}
	}

	if !data.NatVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"appHosting.nat.optionType", "variable")
			body, _ = sjson.Set(body, path+"appHosting.nat.value", data.NatVariable.ValueString())
		}
	} else if !data.Nat.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"appHosting.nat.optionType", "global")
			body, _ = sjson.Set(body, path+"appHosting.nat.value", data.Nat.ValueBool())
		}
	}

	if !data.DownloadUrlDatabaseOnDeviceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"appHosting.databaseUrl.optionType", "variable")
			body, _ = sjson.Set(body, path+"appHosting.databaseUrl.value", data.DownloadUrlDatabaseOnDeviceVariable.ValueString())
		}
	} else if !data.DownloadUrlDatabaseOnDevice.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"appHosting.databaseUrl.optionType", "global")
			body, _ = sjson.Set(body, path+"appHosting.databaseUrl.value", data.DownloadUrlDatabaseOnDevice.ValueBool())
		}
	}

	if !data.ResourceProfileVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"appHosting.resourceProfile.optionType", "variable")
			body, _ = sjson.Set(body, path+"appHosting.resourceProfile.value", data.ResourceProfileVariable.ValueString())
		}
	} else if !data.ResourceProfile.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"appHosting.resourceProfile.optionType", "global")
			body, _ = sjson.Set(body, path+"appHosting.resourceProfile.value", data.ResourceProfile.ValueString())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *EmbeddedSecurity) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	if value := res.Get(path + "assembly"); value.Exists() && len(value.Array()) > 0 {
		data.Assembly = make([]EmbeddedSecurityAssembly, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := EmbeddedSecurityAssembly{}
			item.AdvancedInspectionProfilePolicyId = types.StringNull()

			if t := v.Get("advancedInspectionProfile.refId.optionType"); t.Exists() {
				va := v.Get("advancedInspectionProfile.refId.value")
				if t.String() == "global" {
					item.AdvancedInspectionProfilePolicyId = types.StringValue(va.String())
				}
			}
			item.SslDecryptionProfileId = types.StringNull()

			if t := v.Get("sslDecryption.refId.optionType"); t.Exists() {
				va := v.Get("sslDecryption.refId.value")
				if t.String() == "global" {
					item.SslDecryptionProfileId = types.StringValue(va.String())
				}
			}
			item.NgfwPolicyId = types.StringNull()

			if t := v.Get("ngfirewall.refId.optionType"); t.Exists() {
				va := v.Get("ngfirewall.refId.value")
				if t.String() == "global" {
					item.NgfwPolicyId = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("ngfirewall.entries"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Entries = make([]EmbeddedSecurityAssemblyEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := EmbeddedSecurityAssemblyEntries{}
					cItem.SourceZoneListId = types.StringNull()

					if t := cv.Get("srcZone.refId.optionType"); t.Exists() {
						va := cv.Get("srcZone.refId.value")
						if t.String() == "global" {
							cItem.SourceZoneListId = types.StringValue(va.String())
						}
					}
					cItem.DestinationZoneListId = types.StringNull()

					if t := cv.Get("dstZone.refId.optionType"); t.Exists() {
						va := cv.Get("dstZone.refId.value")
						if t.String() == "global" {
							cItem.DestinationZoneListId = types.StringValue(va.String())
						}
					}
					item.Entries = append(item.Entries, cItem)
					return true
				})
			}
			data.Assembly = append(data.Assembly, item)
			return true
		})
	}
	data.TcpSynFloodLimit = types.StringNull()

	if t := res.Get(path + "settings.tcpSynFloodLimit.optionType"); t.Exists() {
		va := res.Get(path + "settings.tcpSynFloodLimit.value")
		if t.String() == "global" {
			data.TcpSynFloodLimit = types.StringValue(va.String())
		}
	}
	data.MaxIncompleteTcpLimit = types.StringNull()

	if t := res.Get(path + "settings.maxIncompleteTcpLimit.optionType"); t.Exists() {
		va := res.Get(path + "settings.maxIncompleteTcpLimit.value")
		if t.String() == "global" {
			data.MaxIncompleteTcpLimit = types.StringValue(va.String())
		}
	}
	data.MaxIncompleteUdpLimit = types.StringNull()

	if t := res.Get(path + "settings.maxIncompleteUdpLimit.optionType"); t.Exists() {
		va := res.Get(path + "settings.maxIncompleteUdpLimit.value")
		if t.String() == "global" {
			data.MaxIncompleteUdpLimit = types.StringValue(va.String())
		}
	}
	data.MaxIncompleteIcmpLimit = types.StringNull()

	if t := res.Get(path + "settings.maxIncompleteIcmpLimit.optionType"); t.Exists() {
		va := res.Get(path + "settings.maxIncompleteIcmpLimit.value")
		if t.String() == "global" {
			data.MaxIncompleteIcmpLimit = types.StringValue(va.String())
		}
	}
	data.AuditTrail = types.StringNull()

	if t := res.Get(path + "settings.auditTrail.optionType"); t.Exists() {
		va := res.Get(path + "settings.auditTrail.value")
		if t.String() == "global" {
			data.AuditTrail = types.StringValue(va.String())
		}
	}
	data.UnifiedLogging = types.StringNull()

	if t := res.Get(path + "settings.unifiedLogging.optionType"); t.Exists() {
		va := res.Get(path + "settings.unifiedLogging.value")
		if t.String() == "global" {
			data.UnifiedLogging = types.StringValue(va.String())
		}
	}
	data.SessionReclassifyAllow = types.StringNull()

	if t := res.Get(path + "settings.sessionReclassifyAllow.optionType"); t.Exists() {
		va := res.Get(path + "settings.sessionReclassifyAllow.value")
		if t.String() == "global" {
			data.SessionReclassifyAllow = types.StringValue(va.String())
		}
	}
	data.ImcpUnreachableAllow = types.StringNull()

	if t := res.Get(path + "settings.icmpUnreachableAllow.optionType"); t.Exists() {
		va := res.Get(path + "settings.icmpUnreachableAllow.value")
		if t.String() == "global" {
			data.ImcpUnreachableAllow = types.StringValue(va.String())
		}
	}
	data.FailureMode = types.StringNull()

	if t := res.Get(path + "settings.failureMode.optionType"); t.Exists() {
		va := res.Get(path + "settings.failureMode.value")
		if t.String() == "global" {
			data.FailureMode = types.StringValue(va.String())
		}
	}
	data.SecurityLogging = types.StringNull()

	if t := res.Get(path + "settings.securityLogging.optionType"); t.Exists() {
		va := res.Get(path + "settings.securityLogging.value")
		if t.String() == "global" {
			data.SecurityLogging = types.StringValue(va.String())
		}
	}
	data.Nat = types.BoolNull()
	data.NatVariable = types.StringNull()
	if t := res.Get(path + "appHosting.nat.optionType"); t.Exists() {
		va := res.Get(path + "appHosting.nat.value")
		if t.String() == "variable" {
			data.NatVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Nat = types.BoolValue(va.Bool())
		}
	}
	data.DownloadUrlDatabaseOnDevice = types.BoolNull()
	data.DownloadUrlDatabaseOnDeviceVariable = types.StringNull()
	if t := res.Get(path + "appHosting.databaseUrl.optionType"); t.Exists() {
		va := res.Get(path + "appHosting.databaseUrl.value")
		if t.String() == "variable" {
			data.DownloadUrlDatabaseOnDeviceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DownloadUrlDatabaseOnDevice = types.BoolValue(va.Bool())
		}
	}
	data.ResourceProfile = types.StringNull()
	data.ResourceProfileVariable = types.StringNull()
	if t := res.Get(path + "appHosting.resourceProfile.optionType"); t.Exists() {
		va := res.Get(path + "appHosting.resourceProfile.value")
		if t.String() == "variable" {
			data.ResourceProfileVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ResourceProfile = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *EmbeddedSecurity) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	for i := range data.Assembly {
		keys := [...]string{"advancedInspectionProfile.refId", "sslDecryption.refId", "ngfirewall.refId"}
		keyValues := [...]string{data.Assembly[i].AdvancedInspectionProfilePolicyId.ValueString(), data.Assembly[i].SslDecryptionProfileId.ValueString(), data.Assembly[i].NgfwPolicyId.ValueString()}
		keyValuesVariables := [...]string{"", "", ""}

		var r gjson.Result
		res.Get(path + "assembly").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Assembly[i].AdvancedInspectionProfilePolicyId = types.StringNull()

		if t := r.Get("advancedInspectionProfile.refId.optionType"); t.Exists() {
			va := r.Get("advancedInspectionProfile.refId.value")
			if t.String() == "global" {
				data.Assembly[i].AdvancedInspectionProfilePolicyId = types.StringValue(va.String())
			}
		}
		data.Assembly[i].SslDecryptionProfileId = types.StringNull()

		if t := r.Get("sslDecryption.refId.optionType"); t.Exists() {
			va := r.Get("sslDecryption.refId.value")
			if t.String() == "global" {
				data.Assembly[i].SslDecryptionProfileId = types.StringValue(va.String())
			}
		}
		data.Assembly[i].NgfwPolicyId = types.StringNull()

		if t := r.Get("ngfirewall.refId.optionType"); t.Exists() {
			va := r.Get("ngfirewall.refId.value")
			if t.String() == "global" {
				data.Assembly[i].NgfwPolicyId = types.StringValue(va.String())
			}
		}
		for ci := range data.Assembly[i].Entries {
			keys := [...]string{"srcZone.refId", "dstZone.refId"}
			keyValues := [...]string{data.Assembly[i].Entries[ci].SourceZoneListId.ValueString(), data.Assembly[i].Entries[ci].DestinationZoneListId.ValueString()}
			keyValuesVariables := [...]string{"", ""}

			var cr gjson.Result
			r.Get("ngfirewall.entries").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType")
						vv := v.Get(keys[ik] + ".value")
						if tt.Exists() && vv.Exists() {
							if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
								found = true
								continue
							} else if tt.String() == "default" {
								continue
							}
							found = false
							break
						}
						continue
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			data.Assembly[i].Entries[ci].SourceZoneListId = types.StringNull()

			if t := cr.Get("srcZone.refId.optionType"); t.Exists() {
				va := cr.Get("srcZone.refId.value")
				if t.String() == "global" {
					data.Assembly[i].Entries[ci].SourceZoneListId = types.StringValue(va.String())
				}
			}
			data.Assembly[i].Entries[ci].DestinationZoneListId = types.StringNull()

			if t := cr.Get("dstZone.refId.optionType"); t.Exists() {
				va := cr.Get("dstZone.refId.value")
				if t.String() == "global" {
					data.Assembly[i].Entries[ci].DestinationZoneListId = types.StringValue(va.String())
				}
			}
		}
	}
	data.TcpSynFloodLimit = types.StringNull()

	if t := res.Get(path + "settings.tcpSynFloodLimit.optionType"); t.Exists() {
		va := res.Get(path + "settings.tcpSynFloodLimit.value")
		if t.String() == "global" {
			data.TcpSynFloodLimit = types.StringValue(va.String())
		}
	}
	data.MaxIncompleteTcpLimit = types.StringNull()

	if t := res.Get(path + "settings.maxIncompleteTcpLimit.optionType"); t.Exists() {
		va := res.Get(path + "settings.maxIncompleteTcpLimit.value")
		if t.String() == "global" {
			data.MaxIncompleteTcpLimit = types.StringValue(va.String())
		}
	}
	data.MaxIncompleteUdpLimit = types.StringNull()

	if t := res.Get(path + "settings.maxIncompleteUdpLimit.optionType"); t.Exists() {
		va := res.Get(path + "settings.maxIncompleteUdpLimit.value")
		if t.String() == "global" {
			data.MaxIncompleteUdpLimit = types.StringValue(va.String())
		}
	}
	data.MaxIncompleteIcmpLimit = types.StringNull()

	if t := res.Get(path + "settings.maxIncompleteIcmpLimit.optionType"); t.Exists() {
		va := res.Get(path + "settings.maxIncompleteIcmpLimit.value")
		if t.String() == "global" {
			data.MaxIncompleteIcmpLimit = types.StringValue(va.String())
		}
	}
	data.AuditTrail = types.StringNull()

	if t := res.Get(path + "settings.auditTrail.optionType"); t.Exists() {
		va := res.Get(path + "settings.auditTrail.value")
		if t.String() == "global" {
			data.AuditTrail = types.StringValue(va.String())
		}
	}
	data.UnifiedLogging = types.StringNull()

	if t := res.Get(path + "settings.unifiedLogging.optionType"); t.Exists() {
		va := res.Get(path + "settings.unifiedLogging.value")
		if t.String() == "global" {
			data.UnifiedLogging = types.StringValue(va.String())
		}
	}
	data.SessionReclassifyAllow = types.StringNull()

	if t := res.Get(path + "settings.sessionReclassifyAllow.optionType"); t.Exists() {
		va := res.Get(path + "settings.sessionReclassifyAllow.value")
		if t.String() == "global" {
			data.SessionReclassifyAllow = types.StringValue(va.String())
		}
	}
	data.ImcpUnreachableAllow = types.StringNull()

	if t := res.Get(path + "settings.icmpUnreachableAllow.optionType"); t.Exists() {
		va := res.Get(path + "settings.icmpUnreachableAllow.value")
		if t.String() == "global" {
			data.ImcpUnreachableAllow = types.StringValue(va.String())
		}
	}
	data.FailureMode = types.StringNull()

	if t := res.Get(path + "settings.failureMode.optionType"); t.Exists() {
		va := res.Get(path + "settings.failureMode.value")
		if t.String() == "global" {
			data.FailureMode = types.StringValue(va.String())
		}
	}
	data.SecurityLogging = types.StringNull()

	if t := res.Get(path + "settings.securityLogging.optionType"); t.Exists() {
		va := res.Get(path + "settings.securityLogging.value")
		if t.String() == "global" {
			data.SecurityLogging = types.StringValue(va.String())
		}
	}
	data.Nat = types.BoolNull()
	data.NatVariable = types.StringNull()
	if t := res.Get(path + "appHosting.nat.optionType"); t.Exists() {
		va := res.Get(path + "appHosting.nat.value")
		if t.String() == "variable" {
			data.NatVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Nat = types.BoolValue(va.Bool())
		}
	}
	data.DownloadUrlDatabaseOnDevice = types.BoolNull()
	data.DownloadUrlDatabaseOnDeviceVariable = types.StringNull()
	if t := res.Get(path + "appHosting.databaseUrl.optionType"); t.Exists() {
		va := res.Get(path + "appHosting.databaseUrl.value")
		if t.String() == "variable" {
			data.DownloadUrlDatabaseOnDeviceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DownloadUrlDatabaseOnDevice = types.BoolValue(va.Bool())
		}
	}
	data.ResourceProfile = types.StringNull()
	data.ResourceProfileVariable = types.StringNull()
	if t := res.Get(path + "appHosting.resourceProfile.optionType"); t.Exists() {
		va := res.Get(path + "appHosting.resourceProfile.value")
		if t.String() == "variable" {
			data.ResourceProfileVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ResourceProfile = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end updateFromBody
