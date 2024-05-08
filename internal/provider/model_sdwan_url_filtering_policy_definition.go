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

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type URLFilteringPolicyDefinition struct {
	Id                  types.String `tfsdk:"id"`
	Version             types.Int64  `tfsdk:"version"`
	Name                types.String `tfsdk:"name"`
	Description         types.String `tfsdk:"description"`
	Mode                types.String `tfsdk:"mode"`
	Alerts              types.Set    `tfsdk:"alerts"`
	WebCategories       types.Set    `tfsdk:"web_categories"`
	WebCategoriesAction types.String `tfsdk:"web_categories_action"`
	WebReputation       types.String `tfsdk:"web_reputation"`
	TargetVpns          types.Set    `tfsdk:"target_vpns"`
	AllowUrlListId      types.String `tfsdk:"allow_url_list_id"`
	AllowUrlListVersion types.Int64  `tfsdk:"allow_url_list_version"`
	BlockUrlListId      types.String `tfsdk:"block_url_list_id"`
	BlockUrlListVersion types.Int64  `tfsdk:"block_url_list_version"`
	BlockPageAction     types.String `tfsdk:"block_page_action"`
	BlockPageContents   types.String `tfsdk:"block_page_contents"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data URLFilteringPolicyDefinition) getPath() string {
	return "/template/policy/definition/urlfiltering/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data URLFilteringPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "urlFiltering")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, "mode", data.Mode.ValueString())
	}
	if !data.Alerts.IsNull() {
		var values []string
		data.Alerts.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "definition.alerts", values)
	}
	if !data.WebCategories.IsNull() {
		var values []string
		data.WebCategories.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "definition.webCategories", values)
	}
	if !data.WebCategoriesAction.IsNull() {
		body, _ = sjson.Set(body, "definition.webCategoriesAction", data.WebCategoriesAction.ValueString())
	}
	if !data.WebReputation.IsNull() {
		body, _ = sjson.Set(body, "definition.webReputation", data.WebReputation.ValueString())
	}
	if !data.TargetVpns.IsNull() {
		var values []string
		data.TargetVpns.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "definition.targetVpns", values)
	}
	if !data.AllowUrlListId.IsNull() {
		body, _ = sjson.Set(body, "definition.urlWhiteList.ref", data.AllowUrlListId.ValueString())
	}
	if !data.BlockUrlListId.IsNull() {
		body, _ = sjson.Set(body, "definition.urlBlackList.ref", data.BlockUrlListId.ValueString())
	}
	if !data.BlockPageAction.IsNull() {
		body, _ = sjson.Set(body, "definition.blockPageAction", data.BlockPageAction.ValueString())
	}
	if !data.BlockPageContents.IsNull() {
		body, _ = sjson.Set(body, "definition.blockPageContents", data.BlockPageContents.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *URLFilteringPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("mode"); value.Exists() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("definition.alerts"); value.Exists() {
		data.Alerts = helpers.GetStringSet(value.Array())
	} else {
		data.Alerts = types.SetNull(types.StringType)
	}
	if value := res.Get("definition.webCategories"); value.Exists() {
		data.WebCategories = helpers.GetStringSet(value.Array())
	} else {
		data.WebCategories = types.SetNull(types.StringType)
	}
	if value := res.Get("definition.webCategoriesAction"); value.Exists() {
		data.WebCategoriesAction = types.StringValue(value.String())
	} else {
		data.WebCategoriesAction = types.StringNull()
	}
	if value := res.Get("definition.webReputation"); value.Exists() {
		data.WebReputation = types.StringValue(value.String())
	} else {
		data.WebReputation = types.StringNull()
	}
	if value := res.Get("definition.targetVpns"); value.Exists() {
		data.TargetVpns = helpers.GetStringSet(value.Array())
	} else {
		data.TargetVpns = types.SetNull(types.StringType)
	}
	if value := res.Get("definition.urlWhiteList.ref"); value.Exists() {
		data.AllowUrlListId = types.StringValue(value.String())
	} else {
		data.AllowUrlListId = types.StringNull()
	}
	if value := res.Get("definition.urlBlackList.ref"); value.Exists() {
		data.BlockUrlListId = types.StringValue(value.String())
	} else {
		data.BlockUrlListId = types.StringNull()
	}
	if value := res.Get("definition.blockPageAction"); value.Exists() {
		data.BlockPageAction = types.StringValue(value.String())
	} else {
		data.BlockPageAction = types.StringNull()
	}
	if value := res.Get("definition.blockPageContents"); value.Exists() {
		data.BlockPageContents = types.StringValue(value.String())
	} else {
		data.BlockPageContents = types.StringNull()
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *URLFilteringPolicyDefinition) hasChanges(ctx context.Context, state *URLFilteringPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.Mode.Equal(state.Mode) {
		hasChanges = true
	}
	if !data.Alerts.Equal(state.Alerts) {
		hasChanges = true
	}
	if !data.WebCategories.Equal(state.WebCategories) {
		hasChanges = true
	}
	if !data.WebCategoriesAction.Equal(state.WebCategoriesAction) {
		hasChanges = true
	}
	if !data.WebReputation.Equal(state.WebReputation) {
		hasChanges = true
	}
	if !data.TargetVpns.Equal(state.TargetVpns) {
		hasChanges = true
	}
	if !data.AllowUrlListId.Equal(state.AllowUrlListId) {
		hasChanges = true
	}
	if !data.BlockUrlListId.Equal(state.BlockUrlListId) {
		hasChanges = true
	}
	if !data.BlockPageAction.Equal(state.BlockPageAction) {
		hasChanges = true
	}
	if !data.BlockPageContents.Equal(state.BlockPageContents) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *URLFilteringPolicyDefinition) updateVersions(ctx context.Context, state *URLFilteringPolicyDefinition) {
	data.AllowUrlListVersion = state.AllowUrlListVersion
	data.BlockUrlListVersion = state.BlockUrlListVersion
}

// End of section. //template:end updateVersions
