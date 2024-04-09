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

type VPNMembershipPolicyDefinition struct {
	Id          types.String                         `tfsdk:"id"`
	Version     types.Int64                          `tfsdk:"version"`
	Name        types.String                         `tfsdk:"name"`
	Description types.String                         `tfsdk:"description"`
	Sites       []VPNMembershipPolicyDefinitionSites `tfsdk:"sites"`
}

type VPNMembershipPolicyDefinitionSites struct {
	SiteListId      types.String `tfsdk:"site_list_id"`
	SiteListVersion types.Int64  `tfsdk:"site_list_version"`
	VpnListIds      types.Set    `tfsdk:"vpn_list_ids"`
	VpnListVersions types.List   `tfsdk:"vpn_list_versions"`
}

func (data VPNMembershipPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "vpnMembershipGroup")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "definition.sites", []interface{}{})
		for _, item := range data.Sites {
			itemBody := ""
			if !item.SiteListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "siteList", item.SiteListId.ValueString())
			}
			if !item.VpnListIds.IsNull() {
				var values []string
				item.VpnListIds.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "vpnList", values)
			}
			body, _ = sjson.SetRaw(body, "definition.sites.-1", itemBody)
		}
	}
	return body
}

func (data *VPNMembershipPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("definition.sites"); value.Exists() && len(value.Array()) > 0 {
		data.Sites = make([]VPNMembershipPolicyDefinitionSites, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNMembershipPolicyDefinitionSites{}
			if cValue := v.Get("siteList"); cValue.Exists() {
				item.SiteListId = types.StringValue(cValue.String())
			} else {
				item.SiteListId = types.StringNull()
			}
			if cValue := v.Get("vpnList"); cValue.Exists() {
				item.VpnListIds = helpers.GetStringSet(cValue.Array())
			} else {
				item.VpnListIds = types.SetNull(types.StringType)
			}
			data.Sites = append(data.Sites, item)
			return true
		})
	} else {
		if len(data.Sites) > 0 {
			data.Sites = []VPNMembershipPolicyDefinitionSites{}
		}
	}
	data.updateVersions(ctx, &state)
}

func (data *VPNMembershipPolicyDefinition) hasChanges(ctx context.Context, state *VPNMembershipPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if len(data.Sites) != len(state.Sites) {
		hasChanges = true
	} else {
		for i := range data.Sites {
			if !data.Sites[i].SiteListId.Equal(state.Sites[i].SiteListId) {
				hasChanges = true
			}
			if !data.Sites[i].VpnListIds.Equal(state.Sites[i].VpnListIds) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

func (data *VPNMembershipPolicyDefinition) updateVersions(ctx context.Context, state *VPNMembershipPolicyDefinition) {
	for i := range data.Sites {
		dataKeys := [...]string{fmt.Sprintf("%v", data.Sites[i].SiteListId.ValueString())}
		stateIndex := -1
		for j := range state.Sites {
			stateKeys := [...]string{fmt.Sprintf("%v", state.Sites[j].SiteListId.ValueString())}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		if stateIndex > -1 {
			data.Sites[i].SiteListVersion = state.Sites[stateIndex].SiteListVersion
		} else {
			data.Sites[i].SiteListVersion = types.Int64Null()
		}
		if stateIndex > -1 && !state.Sites[stateIndex].VpnListVersions.IsNull() {
			data.Sites[i].VpnListVersions = state.Sites[stateIndex].VpnListVersions
		} else {
			data.Sites[i].VpnListVersions = types.ListNull(types.StringType)
		}
	}
}
