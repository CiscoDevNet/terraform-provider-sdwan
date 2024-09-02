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
type ServiceObjectTrackerGroup struct {
	Id                      types.String                               `tfsdk:"id"`
	Version                 types.Int64                                `tfsdk:"version"`
	Name                    types.String                               `tfsdk:"name"`
	Description             types.String                               `tfsdk:"description"`
	FeatureProfileId        types.String                               `tfsdk:"feature_profile_id"`
	ObjectTrackerId         types.Int64                                `tfsdk:"object_tracker_id"`
	ObjectTrackerIdVariable types.String                               `tfsdk:"object_tracker_id_variable"`
	TrackerElements         []ServiceObjectTrackerGroupTrackerElements `tfsdk:"tracker_elements"`
	Reachable               types.String                               `tfsdk:"reachable"`
	ReachableVariable       types.String                               `tfsdk:"reachable_variable"`
}

type ServiceObjectTrackerGroupTrackerElements struct {
	ObjectTrackerId types.String `tfsdk:"object_tracker_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceObjectTrackerGroup) getModel() string {
	return "service_object_tracker_group"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceObjectTrackerGroup) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/objecttrackergroup", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceObjectTrackerGroup) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.ObjectTrackerIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"objectId.optionType", "variable")
			body, _ = sjson.Set(body, path+"objectId.value", data.ObjectTrackerIdVariable.ValueString())
		}
	} else if !data.ObjectTrackerId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"objectId.optionType", "global")
			body, _ = sjson.Set(body, path+"objectId.value", data.ObjectTrackerId.ValueInt64())
		}
	}
	if true {

		for _, item := range data.TrackerElements {
			itemBody := ""
			if !item.ObjectTrackerId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackerRef.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "trackerRef.refId.value", item.ObjectTrackerId.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"trackerRefs.-1", itemBody)
		}
	}

	if !data.ReachableVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"criteria.optionType", "variable")
			body, _ = sjson.Set(body, path+"criteria.value", data.ReachableVariable.ValueString())
		}
	} else if data.Reachable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"criteria.optionType", "default")
			body, _ = sjson.Set(body, path+"criteria.value", "or")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"criteria.optionType", "global")
			body, _ = sjson.Set(body, path+"criteria.value", data.Reachable.ValueString())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceObjectTrackerGroup) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.ObjectTrackerId = types.Int64Null()
	data.ObjectTrackerIdVariable = types.StringNull()
	if t := res.Get(path + "objectId.optionType"); t.Exists() {
		va := res.Get(path + "objectId.value")
		if t.String() == "variable" {
			data.ObjectTrackerIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ObjectTrackerId = types.Int64Value(va.Int())
		}
	}
	if value := res.Get(path + "trackerRefs"); value.Exists() {
		data.TrackerElements = make([]ServiceObjectTrackerGroupTrackerElements, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceObjectTrackerGroupTrackerElements{}
			item.ObjectTrackerId = types.StringNull()

			if t := v.Get("trackerRef.refId.optionType"); t.Exists() {
				va := v.Get("trackerRef.refId.value")
				if t.String() == "global" {
					item.ObjectTrackerId = types.StringValue(va.String())
				}
			}
			data.TrackerElements = append(data.TrackerElements, item)
			return true
		})
	}
	data.Reachable = types.StringNull()
	data.ReachableVariable = types.StringNull()
	if t := res.Get(path + "criteria.optionType"); t.Exists() {
		va := res.Get(path + "criteria.value")
		if t.String() == "variable" {
			data.ReachableVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Reachable = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ServiceObjectTrackerGroup) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.ObjectTrackerId = types.Int64Null()
	data.ObjectTrackerIdVariable = types.StringNull()
	if t := res.Get(path + "objectId.optionType"); t.Exists() {
		va := res.Get(path + "objectId.value")
		if t.String() == "variable" {
			data.ObjectTrackerIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ObjectTrackerId = types.Int64Value(va.Int())
		}
	}
	for i := range data.TrackerElements {
		keys := [...]string{"trackerRef.refId"}
		keyValues := [...]string{data.TrackerElements[i].ObjectTrackerId.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "trackerRefs").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
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
		data.TrackerElements[i].ObjectTrackerId = types.StringNull()

		if t := r.Get("trackerRef.refId.optionType"); t.Exists() {
			va := r.Get("trackerRef.refId.value")
			if t.String() == "global" {
				data.TrackerElements[i].ObjectTrackerId = types.StringValue(va.String())
			}
		}
	}
	data.Reachable = types.StringNull()
	data.ReachableVariable = types.StringNull()
	if t := res.Get(path + "criteria.optionType"); t.Exists() {
		va := res.Get(path + "criteria.value")
		if t.String() == "variable" {
			data.ReachableVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Reachable = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *ServiceObjectTrackerGroup) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.ObjectTrackerId.IsNull() {
		return false
	}
	if !data.ObjectTrackerIdVariable.IsNull() {
		return false
	}
	if len(data.TrackerElements) > 0 {
		return false
	}
	if !data.Reachable.IsNull() {
		return false
	}
	if !data.ReachableVariable.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
