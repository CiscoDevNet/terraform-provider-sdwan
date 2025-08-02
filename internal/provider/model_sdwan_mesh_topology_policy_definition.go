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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type MeshTopologyPolicyDefinition struct {
	Id             types.String                          `tfsdk:"id"`
	Version        types.Int64                           `tfsdk:"version"`
	Type           types.String                          `tfsdk:"type"`
	Name           types.String                          `tfsdk:"name"`
	Description    types.String                          `tfsdk:"description"`
	VpnListId      types.String                          `tfsdk:"vpn_list_id"`
	VpnListVersion types.Int64                           `tfsdk:"vpn_list_version"`
	Regions        []MeshTopologyPolicyDefinitionRegions `tfsdk:"regions"`
}

type MeshTopologyPolicyDefinitionRegions struct {
	Name             types.String `tfsdk:"name"`
	SiteListIds      types.Set    `tfsdk:"site_list_ids"`
	SiteListVersions types.List   `tfsdk:"site_list_versions"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data MeshTopologyPolicyDefinition) getPath() string {
	return "/template/policy/definition/mesh/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data MeshTopologyPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "mesh")
	}
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

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
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
				item.SiteListIds = helpers.GetStringSet(cValue.Array())
			} else {
				item.SiteListIds = types.SetNull(types.StringType)
			}
			data.Regions = append(data.Regions, item)
			return true
		})
	} else {
		if len(data.Regions) > 0 {
			data.Regions = []MeshTopologyPolicyDefinitionRegions{}
		}
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
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

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

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

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *MeshTopologyPolicyDefinition) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
	if data.VpnListId != types.StringNull() {
		data.VpnListVersion = types.Int64Value(0)
	}
	for i := range data.Regions {
		if !data.Regions[i].SiteListIds.IsNull() {
			data.Regions[i].SiteListVersions = types.ListNull(types.StringType)
		}
	}
}

// End of section. //template:end processImport
