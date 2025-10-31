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
type SLAClassPolicyObject struct {
	Id                         types.String `tfsdk:"id"`
	Version                    types.Int64  `tfsdk:"version"`
	Name                       types.String `tfsdk:"name"`
	AppProbeClassId            types.String `tfsdk:"app_probe_class_id"`
	AppProbeClassVersion       types.Int64  `tfsdk:"app_probe_class_version"`
	Jitter                     types.Int64  `tfsdk:"jitter"`
	Latency                    types.Int64  `tfsdk:"latency"`
	Loss                       types.Int64  `tfsdk:"loss"`
	FallbackBestTunnelCriteria types.String `tfsdk:"fallback_best_tunnel_criteria"`
	FallbackBestTunnelJitter   types.Int64  `tfsdk:"fallback_best_tunnel_jitter"`
	FallbackBestTunnelLatency  types.Int64  `tfsdk:"fallback_best_tunnel_latency"`
	FallbackBestTunnelLoss     types.Int64  `tfsdk:"fallback_best_tunnel_loss"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SLAClassPolicyObject) getPath() string {
	return "/template/policy/list/sla/"
}

// End of section. //template:end getPath

func (data SLAClassPolicyObject) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "sla")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.AppProbeClassId.IsNull() {
		body, _ = sjson.Set(body, "entries.0.appProbeClass", data.AppProbeClassId.ValueString())
	}
	if !data.Jitter.IsNull() {
		body, _ = sjson.Set(body, "entries.0.jitter", fmt.Sprint(data.Jitter.ValueInt64()))
	}
	if !data.Latency.IsNull() {
		body, _ = sjson.Set(body, "entries.0.latency", fmt.Sprint(data.Latency.ValueInt64()))
	}
	if !data.Loss.IsNull() {
		body, _ = sjson.Set(body, "entries.0.loss", fmt.Sprint(data.Loss.ValueInt64()))
	}
	if !data.FallbackBestTunnelCriteria.IsNull() {
		body, _ = sjson.Set(body, "entries.0.fallbackBestTunnel.criteria", data.FallbackBestTunnelCriteria.ValueString())
	}
	if !data.FallbackBestTunnelJitter.IsNull() {
		body, _ = sjson.Set(body, "entries.0.fallbackBestTunnel.jitterVariance", fmt.Sprint(data.FallbackBestTunnelJitter.ValueInt64()))
	} else if !data.FallbackBestTunnelCriteria.IsNull() {
		body, _ = sjson.Set(body, "entries.0.fallbackBestTunnel.jitterVariance", nil)
	}
	if !data.FallbackBestTunnelLatency.IsNull() {
		body, _ = sjson.Set(body, "entries.0.fallbackBestTunnel.latencyVariance", fmt.Sprint(data.FallbackBestTunnelLatency.ValueInt64()))
	} else if !data.FallbackBestTunnelCriteria.IsNull() {
		body, _ = sjson.Set(body, "entries.0.fallbackBestTunnel.latencyVariance", nil)
	}
	if !data.FallbackBestTunnelLoss.IsNull() {
		body, _ = sjson.Set(body, "entries.0.fallbackBestTunnel.lossVariance", fmt.Sprint(data.FallbackBestTunnelLoss.ValueInt64()))
	} else if !data.FallbackBestTunnelCriteria.IsNull() {
		body, _ = sjson.Set(body, "entries.0.fallbackBestTunnel.lossVariance", nil)
	}
	return body
}

func (data *SLAClassPolicyObject) fromBody(ctx context.Context, res gjson.Result) {
	state := *data
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("entries.0.appProbeClass"); value.Exists() {
		data.AppProbeClassId = types.StringValue(value.String())
	} else {
		data.AppProbeClassId = types.StringNull()
	}
	if value := res.Get("entries.0.jitter"); value.Exists() {
		data.Jitter = types.Int64Value(value.Int())
	} else {
		data.Jitter = types.Int64Null()
	}
	if value := res.Get("entries.0.latency"); value.Exists() {
		data.Latency = types.Int64Value(value.Int())
	} else {
		data.Latency = types.Int64Null()
	}
	if value := res.Get("entries.0.loss"); value.Exists() {
		data.Loss = types.Int64Value(value.Int())
	} else {
		data.Loss = types.Int64Null()
	}
	if value := res.Get("entries.0.fallbackBestTunnel.criteria"); value.Exists() && value.String() != "" {
		data.FallbackBestTunnelCriteria = types.StringValue(value.String())
	} else {
		data.FallbackBestTunnelCriteria = types.StringNull()
	}
	if value := res.Get("entries.0.fallbackBestTunnel.jitterVariance"); value.Exists() && value.String() != "" {
		data.FallbackBestTunnelJitter = types.Int64Value(value.Int())
	} else {
		data.FallbackBestTunnelJitter = types.Int64Null()
	}
	if value := res.Get("entries.0.fallbackBestTunnel.latencyVariance"); value.Exists() && value.String() != "" {
		data.FallbackBestTunnelLatency = types.Int64Value(value.Int())
	} else {
		data.FallbackBestTunnelLatency = types.Int64Null()
	}
	if value := res.Get("entries.0.fallbackBestTunnel.lossVariance"); value.Exists() && value.String() != "" {
		data.FallbackBestTunnelLoss = types.Int64Value(value.Int())
	} else {
		data.FallbackBestTunnelLoss = types.Int64Null()
	}
	data.updateVersions(ctx, &state)
}

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *SLAClassPolicyObject) hasChanges(ctx context.Context, state *SLAClassPolicyObject) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.AppProbeClassId.Equal(state.AppProbeClassId) {
		hasChanges = true
	}
	if !data.Jitter.Equal(state.Jitter) {
		hasChanges = true
	}
	if !data.Latency.Equal(state.Latency) {
		hasChanges = true
	}
	if !data.Loss.Equal(state.Loss) {
		hasChanges = true
	}
	if !data.FallbackBestTunnelCriteria.Equal(state.FallbackBestTunnelCriteria) {
		hasChanges = true
	}
	if !data.FallbackBestTunnelJitter.Equal(state.FallbackBestTunnelJitter) {
		hasChanges = true
	}
	if !data.FallbackBestTunnelLatency.Equal(state.FallbackBestTunnelLatency) {
		hasChanges = true
	}
	if !data.FallbackBestTunnelLoss.Equal(state.FallbackBestTunnelLoss) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *SLAClassPolicyObject) updateVersions(ctx context.Context, state *SLAClassPolicyObject) {
	data.AppProbeClassVersion = state.AppProbeClassVersion
}

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *SLAClassPolicyObject) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
	if data.AppProbeClassId != types.StringNull() {
		data.AppProbeClassVersion = types.Int64Value(0)
	}
}

// End of section. //template:end processImport
