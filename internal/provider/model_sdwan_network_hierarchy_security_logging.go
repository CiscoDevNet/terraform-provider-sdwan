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

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// NetworkHierarchySecurityLogging represents security logging settings for a network hierarchy node
type NetworkHierarchySecurityLogging struct {
	Id               types.String                                    `tfsdk:"id"`
	NodeId           types.String                                    `tfsdk:"node_id"`
	HighSpeedLogging []NetworkHierarchySecurityLoggingHighSpeedEntry `tfsdk:"high_speed_logging"`
	UtdSyslog        []NetworkHierarchySecurityLoggingUtdSyslogEntry `tfsdk:"utd_syslog"`
}

// NetworkHierarchySecurityLoggingHighSpeedEntry represents a high speed logging entry
type NetworkHierarchySecurityLoggingHighSpeedEntry struct {
	Vrf      types.String `tfsdk:"vrf"`
	ServerIp types.String `tfsdk:"server_ip"`
	Port     types.Int64  `tfsdk:"port"`
}

// NetworkHierarchySecurityLoggingUtdSyslogEntry represents a UTD syslog entry
type NetworkHierarchySecurityLoggingUtdSyslogEntry struct {
	Vpn      types.String `tfsdk:"vpn"`
	ServerIp types.String `tfsdk:"server_ip"`
}

func (data NetworkHierarchySecurityLogging) getPath() string {
	return fmt.Sprintf("/v1/network-hierarchy/%s/network-settings/security-logging", data.NodeId.ValueString())
}

func (data NetworkHierarchySecurityLogging) getNetworkHierarchyListPath() string {
	return "/v1/network-hierarchy?offset=0&limit=1500"
}

// getGlobalNodeId retrieves the Global node ID from the network hierarchy.
// The Global node is identified by having data.label = "GLOBAL".
func (data NetworkHierarchySecurityLogging) getGlobalNodeId(hierarchyRes gjson.Result) (string, error) {
	var globalNodeId string

	hierarchyRes.ForEach(func(key, value gjson.Result) bool {
		if value.Get("data.label").String() == "GLOBAL" {
			globalNodeId = value.Get("id").String()
			return false
		}
		return true
	})

	if globalNodeId == "" {
		return "", fmt.Errorf("Global node not found in network hierarchy")
	}

	return globalNodeId, nil
}

func (data NetworkHierarchySecurityLogging) getPathWithId() string {
	return fmt.Sprintf("/v1/network-hierarchy/%s/network-settings/security-logging/%s", data.NodeId.ValueString(), data.Id.ValueString())
}

func (data NetworkHierarchySecurityLogging) toBody(ctx context.Context) string {
	body := ""

	if len(data.HighSpeedLogging) > 0 {
		body, _ = sjson.Set(body, "data.highSpeedLogging", []interface{}{})
		for _, entry := range data.HighSpeedLogging {
			itemBody := ""
			if !entry.Vrf.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vrf.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "vrf.value", entry.Vrf.ValueString())
			}
			if !entry.ServerIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "serverIp.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "serverIp.value", entry.ServerIp.ValueString())
			}
			if !entry.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "port.value", entry.Port.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "data.highSpeedLogging.-1", itemBody)
		}
	}

	if len(data.UtdSyslog) > 0 {
		body, _ = sjson.Set(body, "data.utdSyslog", []interface{}{})
		for _, entry := range data.UtdSyslog {
			itemBody := ""
			if !entry.Vpn.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "vpn.value", entry.Vpn.ValueString())
			}
			if !entry.ServerIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "serverIp.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "serverIp.value", entry.ServerIp.ValueString())
			}
			body, _ = sjson.SetRaw(body, "data.utdSyslog.-1", itemBody)
		}
	}

	return body
}

func (data NetworkHierarchySecurityLogging) toEmptyBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "data", map[string]interface{}{})
	return body
}

func (data *NetworkHierarchySecurityLogging) fromBody(ctx context.Context, res gjson.Result) {
	// API response wraps data under "payload.data", not just "data"
	if value := res.Get("payload.data.highSpeedLogging"); value.Exists() && len(value.Array()) > 0 {
		data.HighSpeedLogging = make([]NetworkHierarchySecurityLoggingHighSpeedEntry, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := NetworkHierarchySecurityLoggingHighSpeedEntry{}
			if cValue := v.Get("vrf.value"); cValue.Exists() {
				item.Vrf = types.StringValue(cValue.String())
			} else {
				item.Vrf = types.StringNull()
			}
			if cValue := v.Get("serverIp.value"); cValue.Exists() {
				item.ServerIp = types.StringValue(cValue.String())
			} else {
				item.ServerIp = types.StringNull()
			}
			if cValue := v.Get("port.value"); cValue.Exists() {
				item.Port = types.Int64Value(cValue.Int())
			} else {
				item.Port = types.Int64Null()
			}
			data.HighSpeedLogging = append(data.HighSpeedLogging, item)
			return true
		})
	} else {
		if len(data.HighSpeedLogging) > 0 {
			data.HighSpeedLogging = []NetworkHierarchySecurityLoggingHighSpeedEntry{}
		}
	}

	if value := res.Get("payload.data.utdSyslog"); value.Exists() && len(value.Array()) > 0 {
		data.UtdSyslog = make([]NetworkHierarchySecurityLoggingUtdSyslogEntry, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := NetworkHierarchySecurityLoggingUtdSyslogEntry{}
			if cValue := v.Get("vpn.value"); cValue.Exists() {
				item.Vpn = types.StringValue(cValue.String())
			} else {
				item.Vpn = types.StringNull()
			}
			if cValue := v.Get("serverIp.value"); cValue.Exists() {
				item.ServerIp = types.StringValue(cValue.String())
			} else {
				item.ServerIp = types.StringNull()
			}
			data.UtdSyslog = append(data.UtdSyslog, item)
			return true
		})
	} else {
		if len(data.UtdSyslog) > 0 {
			data.UtdSyslog = []NetworkHierarchySecurityLoggingUtdSyslogEntry{}
		}
	}
}

func (data *NetworkHierarchySecurityLogging) hasChanges(ctx context.Context, state *NetworkHierarchySecurityLogging) bool {
	hasChanges := false

	if len(data.HighSpeedLogging) != len(state.HighSpeedLogging) {
		hasChanges = true
	} else {
		for i := range data.HighSpeedLogging {
			if !data.HighSpeedLogging[i].Vrf.Equal(state.HighSpeedLogging[i].Vrf) {
				hasChanges = true
			}
			if !data.HighSpeedLogging[i].ServerIp.Equal(state.HighSpeedLogging[i].ServerIp) {
				hasChanges = true
			}
			if !data.HighSpeedLogging[i].Port.Equal(state.HighSpeedLogging[i].Port) {
				hasChanges = true
			}
		}
	}

	if len(data.UtdSyslog) != len(state.UtdSyslog) {
		hasChanges = true
	} else {
		for i := range data.UtdSyslog {
			if !data.UtdSyslog[i].Vpn.Equal(state.UtdSyslog[i].Vpn) {
				hasChanges = true
			}
			if !data.UtdSyslog[i].ServerIp.Equal(state.UtdSyslog[i].ServerIp) {
				hasChanges = true
			}
		}
	}

	return hasChanges
}
