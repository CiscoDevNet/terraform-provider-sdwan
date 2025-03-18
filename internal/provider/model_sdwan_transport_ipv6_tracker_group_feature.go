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
type TransportIPv6TrackerGroup struct {
	Id                     types.String                               `tfsdk:"id"`
	Version                types.Int64                                `tfsdk:"version"`
	Name                   types.String                               `tfsdk:"name"`
	Description            types.String                               `tfsdk:"description"`
	FeatureProfileId       types.String                               `tfsdk:"feature_profile_id"`
	TrackerName            types.String                               `tfsdk:"tracker_name"`
	TrackerNameVariable    types.String                               `tfsdk:"tracker_name_variable"`
	TrackerElements        []TransportIPv6TrackerGroupTrackerElements `tfsdk:"tracker_elements"`
	TrackerBoolean         types.String                               `tfsdk:"tracker_boolean"`
	TrackerBooleanVariable types.String                               `tfsdk:"tracker_boolean_variable"`
}

type TransportIPv6TrackerGroupTrackerElements struct {
	TrackerId types.String `tfsdk:"tracker_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TransportIPv6TrackerGroup) getModel() string {
	return "transport_ipv6_tracker_group"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportIPv6TrackerGroup) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/ipv6-trackergroup", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportIPv6TrackerGroup) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.TrackerNameVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trackerGroupName.optionType", "variable")
			body, _ = sjson.Set(body, path+"trackerGroupName.value", data.TrackerNameVariable.ValueString())
		}
	} else if !data.TrackerName.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trackerGroupName.optionType", "global")
			body, _ = sjson.Set(body, path+"trackerGroupName.value", data.TrackerName.ValueString())
		}
	}
	if true {

		for _, item := range data.TrackerElements {
			itemBody := ""
			if !item.TrackerId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackerRef.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "trackerRef.refId.value", item.TrackerId.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"trackerRefs.-1", itemBody)
		}
	}

	if !data.TrackerBooleanVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"combineBoolean.optionType", "variable")
			body, _ = sjson.Set(body, path+"combineBoolean.value", data.TrackerBooleanVariable.ValueString())
		}
	} else if data.TrackerBoolean.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"combineBoolean.optionType", "default")
			body, _ = sjson.Set(body, path+"combineBoolean.value", "or")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"combineBoolean.optionType", "global")
			body, _ = sjson.Set(body, path+"combineBoolean.value", data.TrackerBoolean.ValueString())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportIPv6TrackerGroup) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.TrackerName = types.StringNull()
	data.TrackerNameVariable = types.StringNull()
	if t := res.Get(path + "trackerGroupName.optionType"); t.Exists() {
		va := res.Get(path + "trackerGroupName.value")
		if t.String() == "variable" {
			data.TrackerNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackerName = types.StringValue(va.String())
		}
	}
	if value := res.Get(path + "trackerRefs"); value.Exists() {
		data.TrackerElements = make([]TransportIPv6TrackerGroupTrackerElements, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TransportIPv6TrackerGroupTrackerElements{}
			item.TrackerId = types.StringNull()

			if t := v.Get("trackerRef.refId.optionType"); t.Exists() {
				va := v.Get("trackerRef.refId.value")
				if t.String() == "global" {
					item.TrackerId = types.StringValue(va.String())
				}
			}
			data.TrackerElements = append(data.TrackerElements, item)
			return true
		})
	}
	data.TrackerBoolean = types.StringNull()
	data.TrackerBooleanVariable = types.StringNull()
	if t := res.Get(path + "combineBoolean.optionType"); t.Exists() {
		va := res.Get(path + "combineBoolean.value")
		if t.String() == "variable" {
			data.TrackerBooleanVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackerBoolean = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TransportIPv6TrackerGroup) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.TrackerName = types.StringNull()
	data.TrackerNameVariable = types.StringNull()
	if t := res.Get(path + "trackerGroupName.optionType"); t.Exists() {
		va := res.Get(path + "trackerGroupName.value")
		if t.String() == "variable" {
			data.TrackerNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackerName = types.StringValue(va.String())
		}
	}
	for i := range data.TrackerElements {
		keys := [...]string{"trackerRef.refId"}
		keyValues := [...]string{data.TrackerElements[i].TrackerId.ValueString()}
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
		data.TrackerElements[i].TrackerId = types.StringNull()

		if t := r.Get("trackerRef.refId.optionType"); t.Exists() {
			va := r.Get("trackerRef.refId.value")
			if t.String() == "global" {
				data.TrackerElements[i].TrackerId = types.StringValue(va.String())
			}
		}
	}
	data.TrackerBoolean = types.StringNull()
	data.TrackerBooleanVariable = types.StringNull()
	if t := res.Get(path + "combineBoolean.optionType"); t.Exists() {
		va := res.Get(path + "combineBoolean.value")
		if t.String() == "variable" {
			data.TrackerBooleanVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackerBoolean = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *TransportIPv6TrackerGroup) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.TrackerName.IsNull() {
		return false
	}
	if !data.TrackerNameVariable.IsNull() {
		return false
	}
	if len(data.TrackerElements) > 0 {
		return false
	}
	if !data.TrackerBoolean.IsNull() {
		return false
	}
	if !data.TrackerBooleanVariable.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
