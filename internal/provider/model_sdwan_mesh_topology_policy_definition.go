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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type MeshTopologyPolicyDefinition struct {
	Id             types.String                          `tfsdk:"id"`
	Version        types.Int64                           `tfsdk:"version"`
	Name           types.String                          `tfsdk:"name"`
	Description    types.String                          `tfsdk:"description"`
	VpnListId      types.String                          `tfsdk:"vpn_list_id"`
	VpnListVersion types.Int64                           `tfsdk:"vpn_list_version"`
	Regions        []MeshTopologyPolicyDefinitionRegions `tfsdk:"regions"`
}

type MeshTopologyPolicyDefinitionRegions struct {
	Name             types.String `tfsdk:"name"`
	SiteListIds      types.List   `tfsdk:"site_list_ids"`
	SiteListVersions types.List   `tfsdk:"site_list_versions"`
}

func (data MeshTopologyPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "type", "mesh")
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.VpnListId.IsNull() {
		body, _ = sjson.Set(body, "definition.vpnList", data.VpnListId.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "definition.regions", []interface{}{})
		for _, item := range data.Regions {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.SiteListIds.IsNull() {
				var values []string
				item.SiteListIds.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "siteLists", values)
			}
			body, _ = sjson.SetRaw(body, "definition.regions.-1", itemBody)
		}
	}
	return body
}

func (data *MeshTopologyPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
	state := *data
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
	if value := res.Get("definition.vpnList"); value.Exists() {
		data.VpnListId = types.StringValue(value.String())
	} else {
		data.VpnListId = types.StringNull()
	}
	if value := res.Get("definition.regions"); value.Exists() && len(value.Array()) > 0 {
		data.Regions = make([]MeshTopologyPolicyDefinitionRegions, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := MeshTopologyPolicyDefinitionRegions{}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("siteLists"); cValue.Exists() {
				item.SiteListIds = helpers.GetStringList(cValue.Array())
			} else {
				item.SiteListIds = types.ListNull(types.StringType)
			}
			data.Regions = append(data.Regions, item)
			return true
		})
	}
	data.updateVersions(ctx, &state)
}

func (data *MeshTopologyPolicyDefinition) hasChanges(ctx context.Context, state *MeshTopologyPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.VpnListId.Equal(state.VpnListId) {
		hasChanges = true
	}
	if len(data.Regions) != len(state.Regions) {
		hasChanges = true
	} else {
		for i := range data.Regions {
			if !data.Regions[i].Name.Equal(state.Regions[i].Name) {
				hasChanges = true
			}
			if !data.Regions[i].SiteListIds.Equal(state.Regions[i].SiteListIds) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

func (data *MeshTopologyPolicyDefinition) updateVersions(ctx context.Context, state *MeshTopologyPolicyDefinition) {
	data.VpnListVersion = state.VpnListVersion
	for i := range data.Regions {
		dataKeys := [...]string{fmt.Sprintf("%v", data.Regions[i].Name.ValueString())}
		stateIndex := -1
		for j := range state.Regions {
			stateKeys := [...]string{fmt.Sprintf("%v", state.Regions[j].Name.ValueString())}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		if stateIndex > -1 && !state.Regions[stateIndex].SiteListVersions.IsNull() {
			data.Regions[i].SiteListVersions = state.Regions[stateIndex].SiteListVersions
		} else {
			data.Regions[i].SiteListVersions = types.ListNull(types.StringType)
		}
	}
}
