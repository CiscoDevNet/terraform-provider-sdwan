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
type ServiceDualRouterHA struct {
	Id                  types.String                          `tfsdk:"id"`
	Version             types.Int64                           `tfsdk:"version"`
	Name                types.String                          `tfsdk:"name"`
	Description         types.String                          `tfsdk:"description"`
	FeatureProfileId    types.String                          `tfsdk:"feature_profile_id"`
	RedundancyGroups    []ServiceDualRouterHARedundancyGroups `tfsdk:"redundancy_groups"`
	EnableOptimizePaths types.Bool                            `tfsdk:"enable_optimize_paths"`
}

type ServiceDualRouterHARedundancyGroups struct {
	GroupId types.Int64                                 `tfsdk:"group_id"`
	VpnIds  []ServiceDualRouterHARedundancyGroupsVpnIds `tfsdk:"vpn_ids"`
	TagName types.String                                `tfsdk:"tag_name"`
}

type ServiceDualRouterHARedundancyGroupsVpnIds struct {
	VpnId types.String `tfsdk:"vpn_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceDualRouterHA) getModel() string {
	return "service_dual_router_ha"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceDualRouterHA) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/dual-router-ha", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceDualRouterHA) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {

		for _, item := range data.RedundancyGroups {
			itemBody := ""
			if !item.GroupId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "groupId", item.GroupId.ValueInt64())
				}
			}
			if true {

				for _, childItem := range item.VpnIds {
					itemChildBody := ""
					if !childItem.VpnId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "refId.value", childItem.VpnId.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "vpnIds.-1", itemChildBody)
				}
			}
			if !item.TagName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tagName", item.TagName.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"redundancyGroups.-1", itemBody)
		}
	}
	if data.EnableOptimizePaths.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"enableOptimizePaths.optionType", "default")
			body, _ = sjson.Set(body, path+"enableOptimizePaths.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"enableOptimizePaths.optionType", "global")
			body, _ = sjson.Set(body, path+"enableOptimizePaths.value", data.EnableOptimizePaths.ValueBool())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceDualRouterHA) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	oldRedundancyGroups := data.RedundancyGroups
	if value := res.Get(path + "redundancyGroups"); value.Exists() && len(value.Array()) > 0 {
		data.RedundancyGroups = make([]ServiceDualRouterHARedundancyGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceDualRouterHARedundancyGroups{}
			item.GroupId = types.Int64Null()

			if va := v.Get("groupId"); va.Exists() {
				item.GroupId = types.Int64Value(va.Int())
			}
			if cValue := v.Get("vpnIds"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.VpnIds = make([]ServiceDualRouterHARedundancyGroupsVpnIds, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceDualRouterHARedundancyGroupsVpnIds{}
					cItem.VpnId = types.StringNull()

					if t := cv.Get("refId.optionType"); t.Exists() {
						va := cv.Get("refId.value")
						if t.String() == "global" {
							cItem.VpnId = types.StringValue(va.String())
						}
					}
					item.VpnIds = append(item.VpnIds, cItem)
					return true
				})
			}
			item.TagName = types.StringNull()

			if va := v.Get("tagName"); va.Exists() {
				item.TagName = types.StringValue(va.String())
			}
			data.RedundancyGroups = append(data.RedundancyGroups, item)
			return true
		})
	} else {
		data.RedundancyGroups = nil
	}
	if !fullRead && data.RedundancyGroups != nil {
		resultRedundancyGroups := make([]ServiceDualRouterHARedundancyGroups, 0, len(data.RedundancyGroups))
		matchedRedundancyGroups := make([]bool, len(data.RedundancyGroups))
		for _, oldItem := range oldRedundancyGroups {
			for ni := range data.RedundancyGroups {
				if matchedRedundancyGroups[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.GroupId.ValueInt64() != data.RedundancyGroups[ni].GroupId.ValueInt64() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedRedundancyGroups[ni] = true
					if data.RedundancyGroups[ni].VpnIds != nil {
						resultC := make([]ServiceDualRouterHARedundancyGroupsVpnIds, 0, len(data.RedundancyGroups[ni].VpnIds))
						matchedC := make([]bool, len(data.RedundancyGroups[ni].VpnIds))
						for _, oldCItem := range oldItem.VpnIds {
							for nci := range data.RedundancyGroups[ni].VpnIds {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if oldCItem.VpnId.ValueString() != data.RedundancyGroups[ni].VpnIds[nci].VpnId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.RedundancyGroups[ni].VpnIds[nci])
									break
								}
							}
						}
						for nci := range data.RedundancyGroups[ni].VpnIds {
							if !matchedC[nci] {
								resultC = append(resultC, data.RedundancyGroups[ni].VpnIds[nci])
							}
						}
						data.RedundancyGroups[ni].VpnIds = resultC
					}
					resultRedundancyGroups = append(resultRedundancyGroups, data.RedundancyGroups[ni])
					break
				}
			}
		}
		for ni := range data.RedundancyGroups {
			if !matchedRedundancyGroups[ni] {
				resultRedundancyGroups = append(resultRedundancyGroups, data.RedundancyGroups[ni])
			}
		}
		data.RedundancyGroups = resultRedundancyGroups
	}
	data.EnableOptimizePaths = types.BoolNull()

	if t := res.Get(path + "enableOptimizePaths.optionType"); t.Exists() {
		va := res.Get(path + "enableOptimizePaths.value")
		if t.String() == "global" {
			data.EnableOptimizePaths = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end fromBody
