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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type VPNMembership struct {
	Id          types.String         `tfsdk:"id"`
	Version     types.Int64          `tfsdk:"version"`
	Type        types.String         `tfsdk:"type"`
	Name        types.String         `tfsdk:"name"`
	Description types.String         `tfsdk:"description"`
	Sites       []VPNMembershipSites `tfsdk:"sites"`
}

type VPNMembershipSites struct {
	SiteListId      types.String `tfsdk:"site_list_id"`
	SiteListVersion types.Int64  `tfsdk:"site_list_version"`
	VpnListIds      types.List   `tfsdk:"vpn_list_ids"`
	VpnListVersions types.List   `tfsdk:"vpn_list_versions"`
}

func (data VPNMembership) getType() string {
	return "vpnMembershipGroup"
}

func (data VPNMembership) toBody(ctx context.Context) string {
	body, _ := sjson.Set("", "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	body, _ = sjson.Set(body, "type", "vpnMembershipGroup")
	path := "definition."
	if len(data.Sites) > 0 {
		body, _ = sjson.Set(body, path+"sites", []interface{}{})
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
			body, _ = sjson.SetRaw(body, path+"sites.-1", itemBody)
		}
	}
	return body
}

func (data *VPNMembership) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	path := "definition."
	if value := res.Get(path + "sites"); value.Exists() {
		data.Sites = make([]VPNMembershipSites, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNMembershipSites{}
			if cValue := v.Get("siteList"); cValue.Exists() {
				item.SiteListId = types.StringValue(cValue.String())
			} else {
				item.SiteListId = types.StringNull()
			}
			if cValue := v.Get("vpnList"); cValue.Exists() {
				item.VpnListIds = helpers.GetStringList(cValue.Array())
			} else {
				item.VpnListIds = types.ListNull(types.StringType)
			}
			data.Sites = append(data.Sites, item)
			return true
		})
	}
}

func (data *VPNMembership) hasChanges(ctx context.Context, state *VPNMembership) bool {
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

func (data *VPNMembership) getSiteListVersion(ctx context.Context, id string) types.Int64 {
	for _, item := range data.Sites {
		if item.SiteListId.ValueString() == id {
			return item.SiteListVersion
		}
	}
	return types.Int64Null()
}

func (data *VPNMembership) getVpnListVersions(ctx context.Context, id string) types.List {
	for _, item := range data.Sites {
		if item.SiteListId.ValueString() == id && !item.VpnListVersions.IsNull() {
			return item.VpnListVersions
		}
	}
	return types.ListNull(types.StringType)
}

func (data *VPNMembership) updateVersions(ctx context.Context, state VPNMembership) {
	for s := range data.Sites {
		id := data.Sites[s].SiteListId.ValueString()
		data.Sites[s].SiteListVersion = state.getSiteListVersion(ctx, id)
		data.Sites[s].VpnListVersions = state.getVpnListVersions(ctx, id)
	}
}