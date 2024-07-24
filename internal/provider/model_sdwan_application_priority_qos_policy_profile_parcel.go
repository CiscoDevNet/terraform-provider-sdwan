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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ApplicationPriorityQoSPolicy struct {
	Id                      types.String                                `tfsdk:"id"`
	Version                 types.Int64                                 `tfsdk:"version"`
	Name                    types.String                                `tfsdk:"name"`
	Description             types.String                                `tfsdk:"description"`
	FeatureProfileId        types.String                                `tfsdk:"feature_profile_id"`
	TargetInterface         types.Set                                   `tfsdk:"target_interface"`
	TargetInterfaceVariable types.String                                `tfsdk:"target_interface_variable"`
	QosSchedulers           []ApplicationPriorityQoSPolicyQosSchedulers `tfsdk:"qos_schedulers"`
}

type ApplicationPriorityQoSPolicyQosSchedulers struct {
	ForwardingClassId types.String `tfsdk:"forwarding_class_id"`
	Drops             types.String `tfsdk:"drops"`
	Queue             types.String `tfsdk:"queue"`
	Bandwidth         types.String `tfsdk:"bandwidth"`
	SchedulingType    types.String `tfsdk:"scheduling_type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ApplicationPriorityQoSPolicy) getModel() string {
	return "application_priority_qos_policy"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ApplicationPriorityQoSPolicy) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/application-priority/%v/qos-policy", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ApplicationPriorityQoSPolicy) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.TargetInterfaceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"target.interfaces.optionType", "variable")
			body, _ = sjson.Set(body, path+"target.interfaces.value", data.TargetInterfaceVariable.ValueString())
		}
	} else if !data.TargetInterface.IsNull() {
		body, _ = sjson.Set(body, path+"target.interfaces.optionType", "global")
		var values []string
		data.TargetInterface.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"target.interfaces.value", values)
	}
	if true {
		body, _ = sjson.Set(body, path+"qosMap.qosSchedulers", []interface{}{})
		for _, item := range data.QosSchedulers {
			itemBody := ""
			if !item.ForwardingClassId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "classMapRef.refId.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "classMapRef.refId.value", item.ForwardingClassId.ValueString())
			}
			if !item.Drops.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "drops.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "drops.value", item.Drops.ValueString())
			}
			if !item.Queue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "queue.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "queue.value", item.Queue.ValueString())
			}
			if !item.Bandwidth.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "bandwidthPercent.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "bandwidthPercent.value", item.Bandwidth.ValueString())
			}
			if !item.SchedulingType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "scheduling.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "scheduling.value", item.SchedulingType.ValueString())
			}
			body, _ = sjson.SetRaw(body, path+"qosMap.qosSchedulers.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ApplicationPriorityQoSPolicy) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.TargetInterface = types.SetNull(types.StringType)
	data.TargetInterfaceVariable = types.StringNull()
	if t := res.Get(path + "target.interfaces.optionType"); t.Exists() {
		va := res.Get(path + "target.interfaces.value")
		if t.String() == "variable" {
			data.TargetInterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TargetInterface = helpers.GetStringSet(va.Array())
		}
	}
	if value := res.Get(path + "qosMap.qosSchedulers"); value.Exists() {
		data.QosSchedulers = make([]ApplicationPriorityQoSPolicyQosSchedulers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ApplicationPriorityQoSPolicyQosSchedulers{}
			item.ForwardingClassId = types.StringNull()

			if t := v.Get("classMapRef.refId.optionType"); t.Exists() {
				va := v.Get("classMapRef.refId.value")
				if t.String() == "global" {
					item.ForwardingClassId = types.StringValue(va.String())
				}
			}
			item.Drops = types.StringNull()

			if t := v.Get("drops.optionType"); t.Exists() {
				va := v.Get("drops.value")
				if t.String() == "global" {
					item.Drops = types.StringValue(va.String())
				}
			}
			item.Queue = types.StringNull()

			if t := v.Get("queue.optionType"); t.Exists() {
				va := v.Get("queue.value")
				if t.String() == "global" {
					item.Queue = types.StringValue(va.String())
				}
			}
			item.Bandwidth = types.StringNull()

			if t := v.Get("bandwidthPercent.optionType"); t.Exists() {
				va := v.Get("bandwidthPercent.value")
				if t.String() == "global" {
					item.Bandwidth = types.StringValue(va.String())
				}
			}
			item.SchedulingType = types.StringNull()

			if t := v.Get("scheduling.optionType"); t.Exists() {
				va := v.Get("scheduling.value")
				if t.String() == "global" {
					item.SchedulingType = types.StringValue(va.String())
				}
			}
			data.QosSchedulers = append(data.QosSchedulers, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ApplicationPriorityQoSPolicy) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.TargetInterface = types.SetNull(types.StringType)
	data.TargetInterfaceVariable = types.StringNull()
	if t := res.Get(path + "target.interfaces.optionType"); t.Exists() {
		va := res.Get(path + "target.interfaces.value")
		if t.String() == "variable" {
			data.TargetInterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TargetInterface = helpers.GetStringSet(va.Array())
		}
	}
	for i := range data.QosSchedulers {
		keys := [...]string{"queue"}
		keyValues := [...]string{data.QosSchedulers[i].Queue.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "qosMap.qosSchedulers").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.QosSchedulers[i].ForwardingClassId = types.StringNull()

		if t := r.Get("classMapRef.refId.optionType"); t.Exists() {
			va := r.Get("classMapRef.refId.value")
			if t.String() == "global" {
				data.QosSchedulers[i].ForwardingClassId = types.StringValue(va.String())
			}
		}
		data.QosSchedulers[i].Drops = types.StringNull()

		if t := r.Get("drops.optionType"); t.Exists() {
			va := r.Get("drops.value")
			if t.String() == "global" {
				data.QosSchedulers[i].Drops = types.StringValue(va.String())
			}
		}
		data.QosSchedulers[i].Queue = types.StringNull()

		if t := r.Get("queue.optionType"); t.Exists() {
			va := r.Get("queue.value")
			if t.String() == "global" {
				data.QosSchedulers[i].Queue = types.StringValue(va.String())
			}
		}
		data.QosSchedulers[i].Bandwidth = types.StringNull()

		if t := r.Get("bandwidthPercent.optionType"); t.Exists() {
			va := r.Get("bandwidthPercent.value")
			if t.String() == "global" {
				data.QosSchedulers[i].Bandwidth = types.StringValue(va.String())
			}
		}
		data.QosSchedulers[i].SchedulingType = types.StringNull()

		if t := r.Get("scheduling.optionType"); t.Exists() {
			va := r.Get("scheduling.value")
			if t.String() == "global" {
				data.QosSchedulers[i].SchedulingType = types.StringValue(va.String())
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *ApplicationPriorityQoSPolicy) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.TargetInterface.IsNull() {
		return false
	}
	if !data.TargetInterfaceVariable.IsNull() {
		return false
	}
	if len(data.QosSchedulers) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
