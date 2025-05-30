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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type SecurityPolicy struct {
	Id                                    types.String                `tfsdk:"id"`
	Version                               types.Int64                 `tfsdk:"version"`
	Name                                  types.String                `tfsdk:"name"`
	Description                           types.String                `tfsdk:"description"`
	Mode                                  types.String                `tfsdk:"mode"`
	UseCase                               types.String                `tfsdk:"use_case"`
	Definitions                           []SecurityPolicyDefinitions `tfsdk:"definitions"`
	DirectInternetApplications            types.String                `tfsdk:"direct_internet_applications"`
	TcpSynFloodLimit                      types.String                `tfsdk:"tcp_syn_flood_limit"`
	AuditTrail                            types.String                `tfsdk:"audit_trail"`
	MatchStatisticsPerFilter              types.String                `tfsdk:"match_statistics_per_filter"`
	FailureMode                           types.String                `tfsdk:"failure_mode"`
	HighSpeedLoggingServerIp              types.String                `tfsdk:"high_speed_logging_server_ip"`
	HighSpeedLoggingVpn                   types.String                `tfsdk:"high_speed_logging_vpn"`
	HighSpeedLoggingServerPort            types.String                `tfsdk:"high_speed_logging_server_port"`
	HighSpeedLoggingServerSourceInterface types.String                `tfsdk:"high_speed_logging_server_source_interface"`
	MaxIncompleteIcmpLimit                types.Int64                 `tfsdk:"max_incomplete_icmp_limit"`
	MaxIncompleteTcpLimit                 types.Int64                 `tfsdk:"max_incomplete_tcp_limit"`
	MaxIncompleteUdpLimit                 types.Int64                 `tfsdk:"max_incomplete_udp_limit"`
	SessionReclassifyAllow                types.Bool                  `tfsdk:"session_reclassify_allow"`
	ImcpUnreachableAllow                  types.Bool                  `tfsdk:"imcp_unreachable_allow"`
	UnifiedLogging                        types.Bool                  `tfsdk:"unified_logging"`
	Logging                               []SecurityPolicyLogging     `tfsdk:"logging"`
}

type SecurityPolicyDefinitions struct {
	Id              types.String `tfsdk:"id"`
	Version         types.Int64  `tfsdk:"version"`
	Type            types.String `tfsdk:"type"`
	SourceZone      types.String `tfsdk:"source_zone"`
	DestinationZone types.String `tfsdk:"destination_zone"`
}

type SecurityPolicyLogging struct {
	ExternalSyslogServerIp              types.String `tfsdk:"external_syslog_server_ip"`
	ExternalSyslogServerVpn             types.String `tfsdk:"external_syslog_server_vpn"`
	ExternalSyslogServerSourceInterface types.String `tfsdk:"external_syslog_server_source_interface"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SecurityPolicy) getPath() string {
	return "/template/policy/security/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SecurityPolicy) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "policyType", "feature")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "policyName", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "policyDescription", data.Description.ValueString())
	}
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, "policyMode", data.Mode.ValueString())
	}
	if !data.UseCase.IsNull() {
		body, _ = sjson.Set(body, "policyUseCase", data.UseCase.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "policyDefinition.assembly", []interface{}{})
		for _, item := range data.Definitions {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "definitionId", item.Id.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if !item.SourceZone.IsNull() && item.Type.ValueString() == "zoneBasedFW" {
				itemBody, _ = sjson.Set(itemBody, "entries.srcZoneListId", item.SourceZone.ValueString())
			}
			if !item.DestinationZone.IsNull() && item.Type.ValueString() == "zoneBasedFW" {
				itemBody, _ = sjson.Set(itemBody, "entries.dstZoneListId", item.DestinationZone.ValueString())
			}
			body, _ = sjson.SetRaw(body, "policyDefinition.assembly.-1", itemBody)
		}
	}
	if !data.DirectInternetApplications.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.zoneToNozoneInternet", data.DirectInternetApplications.ValueString())
	}
	if !data.TcpSynFloodLimit.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.tcpSynFloodLimit", data.TcpSynFloodLimit.ValueString())
	}
	if !data.AuditTrail.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.auditTrail", data.AuditTrail.ValueString())
	}
	if !data.MatchStatisticsPerFilter.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.platformMatch", data.MatchStatisticsPerFilter.ValueString())
	}
	if !data.FailureMode.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.failureMode", data.FailureMode.ValueString())
	}
	if !data.HighSpeedLoggingServerIp.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.highSpeedLogging.serverIp", data.HighSpeedLoggingServerIp.ValueString())
	}
	if !data.HighSpeedLoggingVpn.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.highSpeedLogging.vrf", data.HighSpeedLoggingVpn.ValueString())
	}
	if !data.HighSpeedLoggingServerPort.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.highSpeedLogging.port", data.HighSpeedLoggingServerPort.ValueString())
	}
	if !data.HighSpeedLoggingServerSourceInterface.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.highSpeedLogging.sourceInterface", data.HighSpeedLoggingServerSourceInterface.ValueString())
	}
	if !data.MaxIncompleteIcmpLimit.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.maxIncompleteIcmpLimit", data.MaxIncompleteIcmpLimit.ValueInt64())
	}
	if !data.MaxIncompleteTcpLimit.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.maxIncompleteTcpLimit", data.MaxIncompleteTcpLimit.ValueInt64())
	}
	if !data.MaxIncompleteUdpLimit.IsNull() {
		body, _ = sjson.Set(body, "policyDefinition.settings.maxIncompleteUdpLimit", data.MaxIncompleteUdpLimit.ValueInt64())
	}
	if !data.SessionReclassifyAllow.IsNull() {
		if false && data.SessionReclassifyAllow.ValueBool() {
			body, _ = sjson.Set(body, "policyDefinition.settings.sessionReclassifyAllow", "")
		} else {
			body, _ = sjson.Set(body, "policyDefinition.settings.sessionReclassifyAllow", data.SessionReclassifyAllow.ValueBool())
		}
	}
	if !data.ImcpUnreachableAllow.IsNull() {
		if false && data.ImcpUnreachableAllow.ValueBool() {
			body, _ = sjson.Set(body, "policyDefinition.settings.icmpUnreachableAllow", "")
		} else {
			body, _ = sjson.Set(body, "policyDefinition.settings.icmpUnreachableAllow", data.ImcpUnreachableAllow.ValueBool())
		}
	}
	if !data.UnifiedLogging.IsNull() {
		if false && data.UnifiedLogging.ValueBool() {
			body, _ = sjson.Set(body, "policyDefinition.settings.unifiedLogging", "")
		} else {
			body, _ = sjson.Set(body, "policyDefinition.settings.unifiedLogging", data.UnifiedLogging.ValueBool())
		}
	}
	if true {
		body, _ = sjson.Set(body, "policyDefinition.settings.logging", []interface{}{})
		for _, item := range data.Logging {
			itemBody := ""
			if !item.ExternalSyslogServerIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "serverIP", item.ExternalSyslogServerIp.ValueString())
			}
			if !item.ExternalSyslogServerVpn.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vpn", item.ExternalSyslogServerVpn.ValueString())
			}
			if !item.ExternalSyslogServerSourceInterface.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourceInterface", item.ExternalSyslogServerSourceInterface.ValueString())
			}
			body, _ = sjson.SetRaw(body, "policyDefinition.settings.logging.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SecurityPolicy) fromBody(ctx context.Context, res gjson.Result) {
	state := *data
	if value := res.Get("policyName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("policyDescription"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("policyMode"); value.Exists() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("policyUseCase"); value.Exists() {
		data.UseCase = types.StringValue(value.String())
	} else {
		data.UseCase = types.StringNull()
	}
	if value := res.Get("policyDefinition.assembly"); value.Exists() && len(value.Array()) > 0 {
		data.Definitions = make([]SecurityPolicyDefinitions, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SecurityPolicyDefinitions{}
			if cValue := v.Get("definitionId"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("type"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			} else {
				item.Type = types.StringNull()
			}
			if cValue := v.Get("entries.srcZoneListId"); cValue.Exists() && item.Type.ValueString() == "zoneBasedFW" {
				item.SourceZone = types.StringValue(cValue.String())
			} else {
				item.SourceZone = types.StringNull()
			}
			if cValue := v.Get("entries.dstZoneListId"); cValue.Exists() && item.Type.ValueString() == "zoneBasedFW" {
				item.DestinationZone = types.StringValue(cValue.String())
			} else {
				item.DestinationZone = types.StringNull()
			}
			data.Definitions = append(data.Definitions, item)
			return true
		})
	} else {
		if len(data.Definitions) > 0 {
			data.Definitions = []SecurityPolicyDefinitions{}
		}
	}
	if value := res.Get("policyDefinition.settings.zoneToNozoneInternet"); value.Exists() {
		data.DirectInternetApplications = types.StringValue(value.String())
	} else {
		data.DirectInternetApplications = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.tcpSynFloodLimit"); value.Exists() {
		data.TcpSynFloodLimit = types.StringValue(value.String())
	} else {
		data.TcpSynFloodLimit = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.auditTrail"); value.Exists() {
		data.AuditTrail = types.StringValue(value.String())
	} else {
		data.AuditTrail = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.platformMatch"); value.Exists() {
		data.MatchStatisticsPerFilter = types.StringValue(value.String())
	} else {
		data.MatchStatisticsPerFilter = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.failureMode"); value.Exists() {
		data.FailureMode = types.StringValue(value.String())
	} else {
		data.FailureMode = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.highSpeedLogging.serverIp"); value.Exists() {
		data.HighSpeedLoggingServerIp = types.StringValue(value.String())
	} else {
		data.HighSpeedLoggingServerIp = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.highSpeedLogging.vrf"); value.Exists() {
		data.HighSpeedLoggingVpn = types.StringValue(value.String())
	} else {
		data.HighSpeedLoggingVpn = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.highSpeedLogging.port"); value.Exists() {
		data.HighSpeedLoggingServerPort = types.StringValue(value.String())
	} else {
		data.HighSpeedLoggingServerPort = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.highSpeedLogging.sourceInterface"); value.Exists() {
		data.HighSpeedLoggingServerSourceInterface = types.StringValue(value.String())
	} else {
		data.HighSpeedLoggingServerSourceInterface = types.StringNull()
	}
	if value := res.Get("policyDefinition.settings.maxIncompleteIcmpLimit"); value.Exists() {
		data.MaxIncompleteIcmpLimit = types.Int64Value(value.Int())
	} else {
		data.MaxIncompleteIcmpLimit = types.Int64Null()
	}
	if value := res.Get("policyDefinition.settings.maxIncompleteTcpLimit"); value.Exists() {
		data.MaxIncompleteTcpLimit = types.Int64Value(value.Int())
	} else {
		data.MaxIncompleteTcpLimit = types.Int64Null()
	}
	if value := res.Get("policyDefinition.settings.maxIncompleteUdpLimit"); value.Exists() {
		data.MaxIncompleteUdpLimit = types.Int64Value(value.Int())
	} else {
		data.MaxIncompleteUdpLimit = types.Int64Null()
	}
	if value := res.Get("policyDefinition.settings.sessionReclassifyAllow"); value.Exists() {
		if false && value.String() == "" {
			data.SessionReclassifyAllow = types.BoolValue(true)
		} else {
			data.SessionReclassifyAllow = types.BoolValue(value.Bool())
		}
	} else {
		data.SessionReclassifyAllow = types.BoolNull()
	}
	if value := res.Get("policyDefinition.settings.icmpUnreachableAllow"); value.Exists() {
		if false && value.String() == "" {
			data.ImcpUnreachableAllow = types.BoolValue(true)
		} else {
			data.ImcpUnreachableAllow = types.BoolValue(value.Bool())
		}
	} else {
		data.ImcpUnreachableAllow = types.BoolNull()
	}
	if value := res.Get("policyDefinition.settings.unifiedLogging"); value.Exists() {
		if false && value.String() == "" {
			data.UnifiedLogging = types.BoolValue(true)
		} else {
			data.UnifiedLogging = types.BoolValue(value.Bool())
		}
	} else {
		data.UnifiedLogging = types.BoolNull()
	}
	if value := res.Get("policyDefinition.settings.logging"); value.Exists() && len(value.Array()) > 0 {
		data.Logging = make([]SecurityPolicyLogging, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SecurityPolicyLogging{}
			if cValue := v.Get("serverIP"); cValue.Exists() {
				item.ExternalSyslogServerIp = types.StringValue(cValue.String())
			} else {
				item.ExternalSyslogServerIp = types.StringNull()
			}
			if cValue := v.Get("vpn"); cValue.Exists() {
				item.ExternalSyslogServerVpn = types.StringValue(cValue.String())
			} else {
				item.ExternalSyslogServerVpn = types.StringNull()
			}
			if cValue := v.Get("sourceInterface"); cValue.Exists() {
				item.ExternalSyslogServerSourceInterface = types.StringValue(cValue.String())
			} else {
				item.ExternalSyslogServerSourceInterface = types.StringNull()
			}
			data.Logging = append(data.Logging, item)
			return true
		})
	} else {
		if len(data.Logging) > 0 {
			data.Logging = []SecurityPolicyLogging{}
		}
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *SecurityPolicy) hasChanges(ctx context.Context, state *SecurityPolicy) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.Mode.Equal(state.Mode) {
		hasChanges = true
	}
	if !data.UseCase.Equal(state.UseCase) {
		hasChanges = true
	}
	if len(data.Definitions) != len(state.Definitions) {
		hasChanges = true
	} else {
		for i := range data.Definitions {
			if !data.Definitions[i].Id.Equal(state.Definitions[i].Id) {
				hasChanges = true
			}
			if !data.Definitions[i].Type.Equal(state.Definitions[i].Type) {
				hasChanges = true
			}
			if !data.Definitions[i].SourceZone.Equal(state.Definitions[i].SourceZone) {
				hasChanges = true
			}
			if !data.Definitions[i].DestinationZone.Equal(state.Definitions[i].DestinationZone) {
				hasChanges = true
			}
		}
	}
	if !data.DirectInternetApplications.Equal(state.DirectInternetApplications) {
		hasChanges = true
	}
	if !data.TcpSynFloodLimit.Equal(state.TcpSynFloodLimit) {
		hasChanges = true
	}
	if !data.AuditTrail.Equal(state.AuditTrail) {
		hasChanges = true
	}
	if !data.MatchStatisticsPerFilter.Equal(state.MatchStatisticsPerFilter) {
		hasChanges = true
	}
	if !data.FailureMode.Equal(state.FailureMode) {
		hasChanges = true
	}
	if !data.HighSpeedLoggingServerIp.Equal(state.HighSpeedLoggingServerIp) {
		hasChanges = true
	}
	if !data.HighSpeedLoggingVpn.Equal(state.HighSpeedLoggingVpn) {
		hasChanges = true
	}
	if !data.HighSpeedLoggingServerPort.Equal(state.HighSpeedLoggingServerPort) {
		hasChanges = true
	}
	if !data.HighSpeedLoggingServerSourceInterface.Equal(state.HighSpeedLoggingServerSourceInterface) {
		hasChanges = true
	}
	if !data.MaxIncompleteIcmpLimit.Equal(state.MaxIncompleteIcmpLimit) {
		hasChanges = true
	}
	if !data.MaxIncompleteTcpLimit.Equal(state.MaxIncompleteTcpLimit) {
		hasChanges = true
	}
	if !data.MaxIncompleteUdpLimit.Equal(state.MaxIncompleteUdpLimit) {
		hasChanges = true
	}
	if !data.SessionReclassifyAllow.Equal(state.SessionReclassifyAllow) {
		hasChanges = true
	}
	if !data.ImcpUnreachableAllow.Equal(state.ImcpUnreachableAllow) {
		hasChanges = true
	}
	if !data.UnifiedLogging.Equal(state.UnifiedLogging) {
		hasChanges = true
	}
	if len(data.Logging) != len(state.Logging) {
		hasChanges = true
	} else {
		for i := range data.Logging {
			if !data.Logging[i].ExternalSyslogServerIp.Equal(state.Logging[i].ExternalSyslogServerIp) {
				hasChanges = true
			}
			if !data.Logging[i].ExternalSyslogServerVpn.Equal(state.Logging[i].ExternalSyslogServerVpn) {
				hasChanges = true
			}
			if !data.Logging[i].ExternalSyslogServerSourceInterface.Equal(state.Logging[i].ExternalSyslogServerSourceInterface) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *SecurityPolicy) updateVersions(ctx context.Context, state *SecurityPolicy) {
	for i := range data.Definitions {
		dataKeys := [...]string{fmt.Sprintf("%v", data.Definitions[i].Id.ValueString())}
		stateIndex := -1
		for j := range state.Definitions {
			stateKeys := [...]string{fmt.Sprintf("%v", state.Definitions[j].Id.ValueString())}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		if stateIndex > -1 {
			data.Definitions[i].Version = state.Definitions[stateIndex].Version
		} else {
			data.Definitions[i].Version = types.Int64Null()
		}
	}
}

// End of section. //template:end updateVersions
