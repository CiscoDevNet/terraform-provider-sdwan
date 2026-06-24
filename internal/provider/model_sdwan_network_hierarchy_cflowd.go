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

// NetworkHierarchyCflowd represents cflowd settings for a network hierarchy node
type NetworkHierarchyCflowd struct {
	Id                   types.String                      `tfsdk:"id"`
	NodeId               types.String                      `tfsdk:"node_id"`
	FlowActiveTimeout    types.Int64                       `tfsdk:"flow_active_timeout"`
	FlowInactiveTimeout  types.Int64                       `tfsdk:"flow_inactive_timeout"`
	FlowRefreshTime      types.Int64                       `tfsdk:"flow_refresh_time"`
	FlowSamplingInterval types.Int64                       `tfsdk:"flow_sampling_interval"`
	CollectTlocLoopback  types.Bool                        `tfsdk:"collect_tloc_loopback"`
	Protocol             types.String                      `tfsdk:"protocol"`
	CollectTos           types.Bool                        `tfsdk:"collect_tos"`
	CollectDscpOutput    types.Bool                        `tfsdk:"collect_dscp_output"`
	Collectors           []NetworkHierarchyCflowdCollector `tfsdk:"collectors"`
}

// NetworkHierarchyCflowdCollector represents a cflowd collector
type NetworkHierarchyCflowdCollector struct {
	VpnId            types.Int64  `tfsdk:"vpn_id"`
	Address          types.String `tfsdk:"address"`
	UdpPort          types.Int64  `tfsdk:"udp_port"`
	ExportSpread     types.Bool   `tfsdk:"export_spread"`
	BfdMetricsExport types.Bool   `tfsdk:"bfd_metrics_export"`
	ExportInterval   types.Int64  `tfsdk:"export_interval"`
}

func (data NetworkHierarchyCflowd) getPath() string {
	return fmt.Sprintf("/v1/network-hierarchy/%s/network-settings/cflowd", data.NodeId.ValueString())
}

func (data NetworkHierarchyCflowd) getNetworkHierarchyListPath() string {
	return "/v1/network-hierarchy?offset=0&limit=1500"
}

// getGlobalNodeId retrieves the Global node ID from the network hierarchy.
// The Global node is identified by having data.label = "GLOBAL".
func (data NetworkHierarchyCflowd) getGlobalNodeId(hierarchyRes gjson.Result) (string, error) {
	var globalNodeId string

	hierarchyRes.ForEach(func(key, value gjson.Result) bool {
		if value.Get("data.label").String() == "GLOBAL" {
			globalNodeId = value.Get("id").String()
			return false // stop iteration
		}
		return true
	})

	if globalNodeId == "" {
		return "", fmt.Errorf("Global node not found in network hierarchy")
	}

	return globalNodeId, nil
}

func (data NetworkHierarchyCflowd) getPathWithId() string {
	return fmt.Sprintf("/v1/network-hierarchy/%s/network-settings/cflowd/%s", data.NodeId.ValueString(), data.Id.ValueString())
}

func (data NetworkHierarchyCflowd) toBody(ctx context.Context) string {
	body := ""

	if !data.FlowActiveTimeout.IsNull() {
		body, _ = sjson.Set(body, "data.flowActiveTimeout.optionType", "global")
		body, _ = sjson.Set(body, "data.flowActiveTimeout.value", data.FlowActiveTimeout.ValueInt64())
	}

	if !data.FlowInactiveTimeout.IsNull() {
		body, _ = sjson.Set(body, "data.flowInactiveTimeout.optionType", "global")
		body, _ = sjson.Set(body, "data.flowInactiveTimeout.value", data.FlowInactiveTimeout.ValueInt64())
	}

	if !data.FlowRefreshTime.IsNull() {
		body, _ = sjson.Set(body, "data.flowRefreshTime.optionType", "global")
		body, _ = sjson.Set(body, "data.flowRefreshTime.value", data.FlowRefreshTime.ValueInt64())
	}

	if !data.FlowSamplingInterval.IsNull() {
		body, _ = sjson.Set(body, "data.flowSamplingInterval.optionType", "global")
		body, _ = sjson.Set(body, "data.flowSamplingInterval.value", data.FlowSamplingInterval.ValueInt64())
	}

	if !data.CollectTlocLoopback.IsNull() {
		body, _ = sjson.Set(body, "data.collectTlocLoopback.optionType", "global")
		body, _ = sjson.Set(body, "data.collectTlocLoopback.value", data.CollectTlocLoopback.ValueBool())
	}

	if !data.Protocol.IsNull() {
		body, _ = sjson.Set(body, "data.protocol.optionType", "global")
		body, _ = sjson.Set(body, "data.protocol.value", data.Protocol.ValueString())
	}

	if !data.CollectTos.IsNull() {
		body, _ = sjson.Set(body, "data.customizedIpv4RecordFields.collectTos.optionType", "global")
		body, _ = sjson.Set(body, "data.customizedIpv4RecordFields.collectTos.value", data.CollectTos.ValueBool())
	}
	if !data.CollectDscpOutput.IsNull() {
		body, _ = sjson.Set(body, "data.customizedIpv4RecordFields.collectDscpOutput.optionType", "global")
		body, _ = sjson.Set(body, "data.customizedIpv4RecordFields.collectDscpOutput.value", data.CollectDscpOutput.ValueBool())
	}

	if len(data.Collectors) > 0 {
		body, _ = sjson.Set(body, "data.collectors", []interface{}{})
		for _, collector := range data.Collectors {
			itemBody := ""
			if !collector.VpnId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vpnId.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "vpnId.value", collector.VpnId.ValueInt64())
			}
			if !collector.Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "address.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "address.value", collector.Address.ValueString())
			}
			if !collector.UdpPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "udpPort.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "udpPort.value", collector.UdpPort.ValueInt64())
			}
			if !collector.ExportSpread.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "exportSpread.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "exportSpread.value", collector.ExportSpread.ValueBool())
			}
			if !collector.BfdMetricsExport.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "bfdMetricsExport.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "bfdMetricsExport.value", collector.BfdMetricsExport.ValueBool())
			}
			if !collector.ExportInterval.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "exportInterval.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "exportInterval.value", collector.ExportInterval.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "data.collectors.-1", itemBody)
		}
	}

	return body
}

func (data *NetworkHierarchyCflowd) fromBody(ctx context.Context, res gjson.Result) {
	// API response wraps data under "payload.data", not just "data"
	if value := res.Get("payload.data.flowActiveTimeout.value"); value.Exists() {
		data.FlowActiveTimeout = types.Int64Value(value.Int())
	} else {
		data.FlowActiveTimeout = types.Int64Null()
	}

	if value := res.Get("payload.data.flowInactiveTimeout.value"); value.Exists() {
		data.FlowInactiveTimeout = types.Int64Value(value.Int())
	} else {
		data.FlowInactiveTimeout = types.Int64Null()
	}

	if value := res.Get("payload.data.flowRefreshTime.value"); value.Exists() {
		data.FlowRefreshTime = types.Int64Value(value.Int())
	} else {
		data.FlowRefreshTime = types.Int64Null()
	}

	if value := res.Get("payload.data.flowSamplingInterval.value"); value.Exists() {
		data.FlowSamplingInterval = types.Int64Value(value.Int())
	} else {
		data.FlowSamplingInterval = types.Int64Null()
	}

	if value := res.Get("payload.data.collectTlocLoopback.value"); value.Exists() {
		data.CollectTlocLoopback = types.BoolValue(value.Bool())
	} else {
		data.CollectTlocLoopback = types.BoolNull()
	}

	if value := res.Get("payload.data.protocol.value"); value.Exists() {
		data.Protocol = types.StringValue(value.String())
	} else {
		data.Protocol = types.StringNull()
	}

	if value := res.Get("payload.data.customizedIpv4RecordFields.collectTos.value"); value.Exists() {
		data.CollectTos = types.BoolValue(value.Bool())
	} else {
		data.CollectTos = types.BoolNull()
	}
	if value := res.Get("payload.data.customizedIpv4RecordFields.collectDscpOutput.value"); value.Exists() {
		data.CollectDscpOutput = types.BoolValue(value.Bool())
	} else {
		data.CollectDscpOutput = types.BoolNull()
	}

	if value := res.Get("payload.data.collectors"); value.Exists() && len(value.Array()) > 0 {
		data.Collectors = make([]NetworkHierarchyCflowdCollector, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := NetworkHierarchyCflowdCollector{}
			if cValue := v.Get("vpnId.value"); cValue.Exists() {
				item.VpnId = types.Int64Value(cValue.Int())
			} else {
				item.VpnId = types.Int64Null()
			}
			if cValue := v.Get("address.value"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			} else {
				item.Address = types.StringNull()
			}
			if cValue := v.Get("udpPort.value"); cValue.Exists() {
				item.UdpPort = types.Int64Value(cValue.Int())
			} else {
				item.UdpPort = types.Int64Null()
			}
			if cValue := v.Get("exportSpread.value"); cValue.Exists() {
				item.ExportSpread = types.BoolValue(cValue.Bool())
			} else {
				item.ExportSpread = types.BoolNull()
			}
			if cValue := v.Get("bfdMetricsExport.value"); cValue.Exists() {
				item.BfdMetricsExport = types.BoolValue(cValue.Bool())
			} else {
				item.BfdMetricsExport = types.BoolNull()
			}
			if cValue := v.Get("exportInterval.value"); cValue.Exists() {
				item.ExportInterval = types.Int64Value(cValue.Int())
			} else {
				item.ExportInterval = types.Int64Null()
			}
			data.Collectors = append(data.Collectors, item)
			return true
		})
	} else {
		if len(data.Collectors) > 0 {
			data.Collectors = []NetworkHierarchyCflowdCollector{}
		}
	}
}

func (data *NetworkHierarchyCflowd) hasChanges(ctx context.Context, state *NetworkHierarchyCflowd) bool {
	hasChanges := false
	if !data.FlowActiveTimeout.Equal(state.FlowActiveTimeout) {
		hasChanges = true
	}
	if !data.FlowInactiveTimeout.Equal(state.FlowInactiveTimeout) {
		hasChanges = true
	}
	if !data.FlowRefreshTime.Equal(state.FlowRefreshTime) {
		hasChanges = true
	}
	if !data.FlowSamplingInterval.Equal(state.FlowSamplingInterval) {
		hasChanges = true
	}
	if !data.CollectTlocLoopback.Equal(state.CollectTlocLoopback) {
		hasChanges = true
	}
	if !data.Protocol.Equal(state.Protocol) {
		hasChanges = true
	}
	if !data.CollectTos.Equal(state.CollectTos) {
		hasChanges = true
	}
	if !data.CollectDscpOutput.Equal(state.CollectDscpOutput) {
		hasChanges = true
	}
	if len(data.Collectors) != len(state.Collectors) {
		hasChanges = true
	} else {
		for i := range data.Collectors {
			if !data.Collectors[i].VpnId.Equal(state.Collectors[i].VpnId) {
				hasChanges = true
			}
			if !data.Collectors[i].Address.Equal(state.Collectors[i].Address) {
				hasChanges = true
			}
			if !data.Collectors[i].UdpPort.Equal(state.Collectors[i].UdpPort) {
				hasChanges = true
			}
			if !data.Collectors[i].ExportSpread.Equal(state.Collectors[i].ExportSpread) {
				hasChanges = true
			}
			if !data.Collectors[i].BfdMetricsExport.Equal(state.Collectors[i].BfdMetricsExport) {
				hasChanges = true
			}
			if !data.Collectors[i].ExportInterval.Equal(state.Collectors[i].ExportInterval) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}
