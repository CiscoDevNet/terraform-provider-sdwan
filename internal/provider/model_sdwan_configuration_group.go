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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ConfigurationGroup struct {
	Id                  types.String                        `tfsdk:"id"`
	Version             types.Int64                         `tfsdk:"version"`
	Name                types.String                        `tfsdk:"name"`
	Description         types.String                        `tfsdk:"description"`
	Solution            types.String                        `tfsdk:"solution"`
	FeatureProfileIds   types.Set                           `tfsdk:"feature_profile_ids"`
	TopologyDevices     []ConfigurationGroupTopologyDevices `tfsdk:"topology_devices"`
	TopologySiteDevices types.Int64                         `tfsdk:"topology_site_devices"`
	FeatureVersions     types.List                          `tfsdk:"feature_versions"`
}

type ConfigurationGroupTopologyDevices struct {
	CriteriaAttribute   types.String                                           `tfsdk:"criteria_attribute"`
	CriteriaValue       types.String                                           `tfsdk:"criteria_value"`
	UnsupportedFeatures []ConfigurationGroupTopologyDevicesUnsupportedFeatures `tfsdk:"unsupported_features"`
}

type ConfigurationGroupTopologyDevicesUnsupportedFeatures struct {
	ParcelType types.String `tfsdk:"parcel_type"`
	ParcelId   types.String `tfsdk:"parcel_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ConfigurationGroup) getPath() string {
	return "/v1/config-group/"
}

// End of section. //template:end getPath

func (data ConfigurationGroup) toBodyConfigGroup(ctx context.Context) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Solution.IsNull() {
		body, _ = sjson.Set(body, "solution", data.Solution.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "profiles", []interface{}{})
		for _, item := range data.FeatureProfileIds.Elements() {
			itemBody := ""
			if !item.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", strings.Trim(item.String(), "\""))
			}
			body, _ = sjson.SetRaw(body, "profiles.-1", itemBody)
		}
	}
	if true && len(data.TopologyDevices) > 0 {
		body, _ = sjson.Set(body, "topology.devices", []interface{}{})
		for _, item := range data.TopologyDevices {
			itemBody := ""
			if !item.CriteriaAttribute.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "criteria.attribute", item.CriteriaAttribute.ValueString())
			}
			if !item.CriteriaValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "criteria.value", item.CriteriaValue.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "unsupportedFeatures", []interface{}{})
				for _, childItem := range item.UnsupportedFeatures {
					itemChildBody := ""
					if !childItem.ParcelType.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "parcelType", childItem.ParcelType.ValueString())
					}
					if !childItem.ParcelId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "parcelId", childItem.ParcelId.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "unsupportedFeatures.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "topology.devices.-1", itemBody)
		}
	}
	if !data.TopologySiteDevices.IsNull() {
		body, _ = sjson.Set(body, "topology.siteDevices", data.TopologySiteDevices.ValueInt64())
	}
	return body
}

func (data *ConfigurationGroup) fromBodyConfigGroup(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("solution"); value.Exists() {
		data.Solution = types.StringValue(value.String())
	} else {
		data.Solution = types.StringNull()
	}
	if value := res.Get("profiles"); value.Exists() && len(value.Array()) > 0 {
		a := make([]attr.Value, len(value.Array()))
		c := 0
		value.ForEach(func(k, v gjson.Result) bool {
			if cValue := v.Get("id"); cValue.Exists() {
				a[c] = types.StringValue(cValue.String())
			} else {
				a[c] = types.StringNull()
			}
			c += 1
			return true
		})
		data.FeatureProfileIds = types.SetValueMust(types.StringType, a)
	} else {
		data.FeatureProfileIds = types.SetNull(types.StringType)
	}
	if value := res.Get("topology.devices"); value.Exists() && len(value.Array()) > 0 {
		data.TopologyDevices = make([]ConfigurationGroupTopologyDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ConfigurationGroupTopologyDevices{}
			if cValue := v.Get("criteria.attribute"); cValue.Exists() {
				item.CriteriaAttribute = types.StringValue(cValue.String())
			} else {
				item.CriteriaAttribute = types.StringNull()
			}
			if cValue := v.Get("criteria.value"); cValue.Exists() {
				item.CriteriaValue = types.StringValue(cValue.String())
			} else {
				item.CriteriaValue = types.StringNull()
			}
			if cValue := v.Get("unsupportedFeatures"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.UnsupportedFeatures = make([]ConfigurationGroupTopologyDevicesUnsupportedFeatures, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ConfigurationGroupTopologyDevicesUnsupportedFeatures{}
					if ccValue := cv.Get("parcelType"); ccValue.Exists() {
						cItem.ParcelType = types.StringValue(ccValue.String())
					} else {
						cItem.ParcelType = types.StringNull()
					}
					if ccValue := cv.Get("parcelId"); ccValue.Exists() {
						cItem.ParcelId = types.StringValue(ccValue.String())
					} else {
						cItem.ParcelId = types.StringNull()
					}
					item.UnsupportedFeatures = append(item.UnsupportedFeatures, cItem)
					return true
				})
			} else {
				if len(item.UnsupportedFeatures) > 0 {
					item.UnsupportedFeatures = []ConfigurationGroupTopologyDevicesUnsupportedFeatures{}
				}
			}
			data.TopologyDevices = append(data.TopologyDevices, item)
			return true
		})
	} else {
		if len(data.TopologyDevices) > 0 {
			data.TopologyDevices = []ConfigurationGroupTopologyDevices{}
		}
	}
	if value := res.Get("topology.siteDevices"); value.Exists() {
		data.TopologySiteDevices = types.Int64Value(value.Int())
	} else {
		data.TopologySiteDevices = types.Int64Null()
	}
}

func (data ConfigurationGroup) hasFeatureVersionChanges(ctx context.Context, state *ConfigurationGroup) bool {
	var planValues, stateValues []string
	data.FeatureVersions.ElementsAs(ctx, &planValues, false)
	state.FeatureVersions.ElementsAs(ctx, &stateValues, false)
	if len(planValues) != len(stateValues) {
		return true
	}
	for i := range planValues {
		if i >= len(stateValues) || planValues[i] != stateValues[i] {
			return true
		}
	}
	return false
}
