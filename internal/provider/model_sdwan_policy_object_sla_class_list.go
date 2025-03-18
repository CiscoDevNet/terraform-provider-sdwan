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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type PolicyObjectSLAClassList struct {
	Id               types.String                      `tfsdk:"id"`
	Version          types.Int64                       `tfsdk:"version"`
	Name             types.String                      `tfsdk:"name"`
	Description      types.String                      `tfsdk:"description"`
	FeatureProfileId types.String                      `tfsdk:"feature_profile_id"`
	Entries          []PolicyObjectSLAClassListEntries `tfsdk:"entries"`
}

type PolicyObjectSLAClassListEntries struct {
	Latency                           types.Int64  `tfsdk:"latency"`
	Loss                              types.Int64  `tfsdk:"loss"`
	Jitter                            types.Int64  `tfsdk:"jitter"`
	AppProbeClassListId               types.String `tfsdk:"app_probe_class_list_id"`
	FallbackBestTunnelCriteria        types.String `tfsdk:"fallback_best_tunnel_criteria"`
	FallbackBestTunnelLossVariance    types.Int64  `tfsdk:"fallback_best_tunnel_loss_variance"`
	FallbackBestTunnelLatencyVariance types.Int64  `tfsdk:"fallback_best_tunnel_latency_variance"`
	FallbackBestTunnelJitterVariance  types.Int64  `tfsdk:"fallback_best_tunnel_jitter_variance"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectSLAClassList) getModel() string {
	return "policy_object_sla_class_list"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectSLAClassList) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/sla-class", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectSLAClassList) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {

		for _, item := range data.Entries {
			itemBody := ""
			if !item.Latency.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "latency.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "latency.value", item.Latency.ValueInt64())
				}
			}
			if !item.Loss.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "loss.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "loss.value", item.Loss.ValueInt64())
				}
			}
			if !item.Jitter.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "jitter.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "jitter.value", item.Jitter.ValueInt64())
				}
			}
			if !item.AppProbeClassListId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "appProbeClass.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "appProbeClass.refId.value", item.AppProbeClassListId.ValueString())
				}
			}
			if !item.FallbackBestTunnelCriteria.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "fallbackBestTunnel.criteria.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "fallbackBestTunnel.criteria.value", item.FallbackBestTunnelCriteria.ValueString())
				}
			}
			if !item.FallbackBestTunnelLossVariance.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "fallbackBestTunnel.lossVariance.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "fallbackBestTunnel.lossVariance.value", item.FallbackBestTunnelLossVariance.ValueInt64())
				}
			}
			if !item.FallbackBestTunnelLatencyVariance.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "fallbackBestTunnel.latencyVariance.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "fallbackBestTunnel.latencyVariance.value", item.FallbackBestTunnelLatencyVariance.ValueInt64())
				}
			}
			if !item.FallbackBestTunnelJitterVariance.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "fallbackBestTunnel.jitterVariance.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "fallbackBestTunnel.jitterVariance.value", item.FallbackBestTunnelJitterVariance.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectSLAClassList) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	if value := res.Get(path + "entries"); value.Exists() {
		data.Entries = make([]PolicyObjectSLAClassListEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyObjectSLAClassListEntries{}
			item.Latency = types.Int64Null()

			if t := v.Get("latency.optionType"); t.Exists() {
				va := v.Get("latency.value")
				if t.String() == "global" {
					item.Latency = types.Int64Value(va.Int())
				}
			}
			item.Loss = types.Int64Null()

			if t := v.Get("loss.optionType"); t.Exists() {
				va := v.Get("loss.value")
				if t.String() == "global" {
					item.Loss = types.Int64Value(va.Int())
				}
			}
			item.Jitter = types.Int64Null()

			if t := v.Get("jitter.optionType"); t.Exists() {
				va := v.Get("jitter.value")
				if t.String() == "global" {
					item.Jitter = types.Int64Value(va.Int())
				}
			}
			item.AppProbeClassListId = types.StringNull()

			if t := v.Get("appProbeClass.refId.optionType"); t.Exists() {
				va := v.Get("appProbeClass.refId.value")
				if t.String() == "global" {
					item.AppProbeClassListId = types.StringValue(va.String())
				}
			}
			item.FallbackBestTunnelCriteria = types.StringNull()

			if t := v.Get("fallbackBestTunnel.criteria.optionType"); t.Exists() {
				va := v.Get("fallbackBestTunnel.criteria.value")
				if t.String() == "global" {
					item.FallbackBestTunnelCriteria = types.StringValue(va.String())
				}
			}
			item.FallbackBestTunnelLossVariance = types.Int64Null()

			if t := v.Get("fallbackBestTunnel.lossVariance.optionType"); t.Exists() {
				va := v.Get("fallbackBestTunnel.lossVariance.value")
				if t.String() == "global" {
					item.FallbackBestTunnelLossVariance = types.Int64Value(va.Int())
				}
			}
			item.FallbackBestTunnelLatencyVariance = types.Int64Null()

			if t := v.Get("fallbackBestTunnel.latencyVariance.optionType"); t.Exists() {
				va := v.Get("fallbackBestTunnel.latencyVariance.value")
				if t.String() == "global" {
					item.FallbackBestTunnelLatencyVariance = types.Int64Value(va.Int())
				}
			}
			item.FallbackBestTunnelJitterVariance = types.Int64Null()

			if t := v.Get("fallbackBestTunnel.jitterVariance.optionType"); t.Exists() {
				va := v.Get("fallbackBestTunnel.jitterVariance.value")
				if t.String() == "global" {
					item.FallbackBestTunnelJitterVariance = types.Int64Value(va.Int())
				}
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PolicyObjectSLAClassList) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	for i := range data.Entries {
		keys := [...]string{"latency", "loss", "jitter", "appProbeClass.refId", "fallbackBestTunnel.criteria", "fallbackBestTunnel.lossVariance", "fallbackBestTunnel.latencyVariance", "fallbackBestTunnel.jitterVariance"}
		keyValues := [...]string{strconv.FormatInt(data.Entries[i].Latency.ValueInt64(), 10), strconv.FormatInt(data.Entries[i].Loss.ValueInt64(), 10), strconv.FormatInt(data.Entries[i].Jitter.ValueInt64(), 10), data.Entries[i].AppProbeClassListId.ValueString(), data.Entries[i].FallbackBestTunnelCriteria.ValueString(), strconv.FormatInt(data.Entries[i].FallbackBestTunnelLossVariance.ValueInt64(), 10), strconv.FormatInt(data.Entries[i].FallbackBestTunnelLatencyVariance.ValueInt64(), 10), strconv.FormatInt(data.Entries[i].FallbackBestTunnelJitterVariance.ValueInt64(), 10)}
		keyValuesVariables := [...]string{"", "", "", "", "", "", "", ""}

		var r gjson.Result
		res.Get(path + "entries").ForEach(
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
		data.Entries[i].Latency = types.Int64Null()

		if t := r.Get("latency.optionType"); t.Exists() {
			va := r.Get("latency.value")
			if t.String() == "global" {
				data.Entries[i].Latency = types.Int64Value(va.Int())
			}
		}
		data.Entries[i].Loss = types.Int64Null()

		if t := r.Get("loss.optionType"); t.Exists() {
			va := r.Get("loss.value")
			if t.String() == "global" {
				data.Entries[i].Loss = types.Int64Value(va.Int())
			}
		}
		data.Entries[i].Jitter = types.Int64Null()

		if t := r.Get("jitter.optionType"); t.Exists() {
			va := r.Get("jitter.value")
			if t.String() == "global" {
				data.Entries[i].Jitter = types.Int64Value(va.Int())
			}
		}
		data.Entries[i].AppProbeClassListId = types.StringNull()

		if t := r.Get("appProbeClass.refId.optionType"); t.Exists() {
			va := r.Get("appProbeClass.refId.value")
			if t.String() == "global" {
				data.Entries[i].AppProbeClassListId = types.StringValue(va.String())
			}
		}
		data.Entries[i].FallbackBestTunnelCriteria = types.StringNull()

		if t := r.Get("fallbackBestTunnel.criteria.optionType"); t.Exists() {
			va := r.Get("fallbackBestTunnel.criteria.value")
			if t.String() == "global" {
				data.Entries[i].FallbackBestTunnelCriteria = types.StringValue(va.String())
			}
		}
		data.Entries[i].FallbackBestTunnelLossVariance = types.Int64Null()

		if t := r.Get("fallbackBestTunnel.lossVariance.optionType"); t.Exists() {
			va := r.Get("fallbackBestTunnel.lossVariance.value")
			if t.String() == "global" {
				data.Entries[i].FallbackBestTunnelLossVariance = types.Int64Value(va.Int())
			}
		}
		data.Entries[i].FallbackBestTunnelLatencyVariance = types.Int64Null()

		if t := r.Get("fallbackBestTunnel.latencyVariance.optionType"); t.Exists() {
			va := r.Get("fallbackBestTunnel.latencyVariance.value")
			if t.String() == "global" {
				data.Entries[i].FallbackBestTunnelLatencyVariance = types.Int64Value(va.Int())
			}
		}
		data.Entries[i].FallbackBestTunnelJitterVariance = types.Int64Null()

		if t := r.Get("fallbackBestTunnel.jitterVariance.optionType"); t.Exists() {
			va := r.Get("fallbackBestTunnel.jitterVariance.value")
			if t.String() == "global" {
				data.Entries[i].FallbackBestTunnelJitterVariance = types.Int64Value(va.Int())
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *PolicyObjectSLAClassList) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if len(data.Entries) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
