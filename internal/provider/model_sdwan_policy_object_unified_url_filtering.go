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
type PolicyObjectUnifiedURLFiltering struct {
	Id                  types.String `tfsdk:"id"`
	Version             types.Int64  `tfsdk:"version"`
	Name                types.String `tfsdk:"name"`
	Description         types.String `tfsdk:"description"`
	FeatureProfileId    types.String `tfsdk:"feature_profile_id"`
	WebCategoriesAction types.String `tfsdk:"web_categories_action"`
	WebCategories       types.Set    `tfsdk:"web_categories"`
	WebReputation       types.String `tfsdk:"web_reputation"`
	UrlAllowListId      types.String `tfsdk:"url_allow_list_id"`
	UrlBlockListId      types.String `tfsdk:"url_block_list_id"`
	BlockPageAction     types.String `tfsdk:"block_page_action"`
	BlockPageContents   types.String `tfsdk:"block_page_contents"`
	RedirectUrl         types.String `tfsdk:"redirect_url"`
	EnableAlerts        types.Bool   `tfsdk:"enable_alerts"`
	Alerts              types.Set    `tfsdk:"alerts"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectUnifiedURLFiltering) getModel() string {
	return "policy_object_unified_url_filtering"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectUnifiedURLFiltering) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/unified/url-filtering", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectUnifiedURLFiltering) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.WebCategoriesAction.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"webCategoriesAction.optionType", "global")
			body, _ = sjson.Set(body, path+"webCategoriesAction.value", data.WebCategoriesAction.ValueString())
		}
	}
	if !data.WebCategories.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"webCategories.optionType", "global")
			var values []string
			data.WebCategories.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"webCategories.value", values)
		}
	}
	if !data.WebReputation.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"webReputation.optionType", "global")
			body, _ = sjson.Set(body, path+"webReputation.value", data.WebReputation.ValueString())
		}
	}
	if !data.UrlAllowListId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"urlAllowedList.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"urlAllowedList.refId.value", data.UrlAllowListId.ValueString())
		}
	}
	if !data.UrlBlockListId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"urlBlockedList.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"urlBlockedList.refId.value", data.UrlBlockListId.ValueString())
		}
	}
	if !data.BlockPageAction.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"blockPageAction.optionType", "global")
			body, _ = sjson.Set(body, path+"blockPageAction.value", data.BlockPageAction.ValueString())
		}
	}
	if !data.BlockPageContents.IsNull() {
		if true && data.BlockPageAction.ValueString() == "text" {
			body, _ = sjson.Set(body, path+"blockPageContents.optionType", "global")
			body, _ = sjson.Set(body, path+"blockPageContents.value", data.BlockPageContents.ValueString())
		}
	}
	if !data.RedirectUrl.IsNull() {
		if true && data.BlockPageAction.ValueString() == "redirect-url" {
			body, _ = sjson.Set(body, path+"redirectUrl.optionType", "global")
			body, _ = sjson.Set(body, path+"redirectUrl.value", data.RedirectUrl.ValueString())
		}
	}
	if !data.EnableAlerts.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"enableAlerts.optionType", "global")
			body, _ = sjson.Set(body, path+"enableAlerts.value", data.EnableAlerts.ValueBool())
		}
	}
	if !data.Alerts.IsNull() {
		if true && data.EnableAlerts.ValueBool() == true {
			body, _ = sjson.Set(body, path+"alerts.optionType", "global")
			var values []string
			data.Alerts.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"alerts.value", values)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectUnifiedURLFiltering) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.WebCategoriesAction = types.StringNull()

	if t := res.Get(path + "webCategoriesAction.optionType"); t.Exists() {
		va := res.Get(path + "webCategoriesAction.value")
		if t.String() == "global" {
			data.WebCategoriesAction = types.StringValue(va.String())
		}
	}
	data.WebCategories = types.SetNull(types.StringType)

	if t := res.Get(path + "webCategories.optionType"); t.Exists() {
		va := res.Get(path + "webCategories.value")
		if t.String() == "global" {
			data.WebCategories = helpers.GetStringSet(va.Array())
		}
	}
	data.WebReputation = types.StringNull()

	if t := res.Get(path + "webReputation.optionType"); t.Exists() {
		va := res.Get(path + "webReputation.value")
		if t.String() == "global" {
			data.WebReputation = types.StringValue(va.String())
		}
	}
	data.UrlAllowListId = types.StringNull()

	if t := res.Get(path + "urlAllowedList.refId.optionType"); t.Exists() {
		va := res.Get(path + "urlAllowedList.refId.value")
		if t.String() == "global" {
			data.UrlAllowListId = types.StringValue(va.String())
		}
	}
	data.UrlBlockListId = types.StringNull()

	if t := res.Get(path + "urlBlockedList.refId.optionType"); t.Exists() {
		va := res.Get(path + "urlBlockedList.refId.value")
		if t.String() == "global" {
			data.UrlBlockListId = types.StringValue(va.String())
		}
	}
	data.BlockPageAction = types.StringNull()

	if t := res.Get(path + "blockPageAction.optionType"); t.Exists() {
		va := res.Get(path + "blockPageAction.value")
		if t.String() == "global" {
			data.BlockPageAction = types.StringValue(va.String())
		}
	}
	data.BlockPageContents = types.StringNull()

	if t := res.Get(path + "blockPageContents.optionType"); t.Exists() {
		va := res.Get(path + "blockPageContents.value")
		if t.String() == "global" {
			data.BlockPageContents = types.StringValue(va.String())
		}
		data.BlockPageAction = types.StringValue("text")
	}
	data.RedirectUrl = types.StringNull()

	if t := res.Get(path + "redirectUrl.optionType"); t.Exists() {
		va := res.Get(path + "redirectUrl.value")
		if t.String() == "global" {
			data.RedirectUrl = types.StringValue(va.String())
		}
		data.BlockPageAction = types.StringValue("redirect-url")
	}
	data.EnableAlerts = types.BoolNull()

	if t := res.Get(path + "enableAlerts.optionType"); t.Exists() {
		va := res.Get(path + "enableAlerts.value")
		if t.String() == "global" {
			data.EnableAlerts = types.BoolValue(va.Bool())
		}
	}
	data.Alerts = types.SetNull(types.StringType)

	if t := res.Get(path + "alerts.optionType"); t.Exists() {
		va := res.Get(path + "alerts.value")
		if t.String() == "global" {
			data.Alerts = helpers.GetStringSet(va.Array())
		}
		data.EnableAlerts = types.BoolValue(true)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PolicyObjectUnifiedURLFiltering) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.WebCategoriesAction = types.StringNull()

	if t := res.Get(path + "webCategoriesAction.optionType"); t.Exists() {
		va := res.Get(path + "webCategoriesAction.value")
		if t.String() == "global" {
			data.WebCategoriesAction = types.StringValue(va.String())
		}
	}
	data.WebCategories = types.SetNull(types.StringType)

	if t := res.Get(path + "webCategories.optionType"); t.Exists() {
		va := res.Get(path + "webCategories.value")
		if t.String() == "global" {
			data.WebCategories = helpers.GetStringSet(va.Array())
		}
	}
	data.WebReputation = types.StringNull()

	if t := res.Get(path + "webReputation.optionType"); t.Exists() {
		va := res.Get(path + "webReputation.value")
		if t.String() == "global" {
			data.WebReputation = types.StringValue(va.String())
		}
	}
	data.UrlAllowListId = types.StringNull()

	if t := res.Get(path + "urlAllowedList.refId.optionType"); t.Exists() {
		va := res.Get(path + "urlAllowedList.refId.value")
		if t.String() == "global" {
			data.UrlAllowListId = types.StringValue(va.String())
		}
	}
	data.UrlBlockListId = types.StringNull()

	if t := res.Get(path + "urlBlockedList.refId.optionType"); t.Exists() {
		va := res.Get(path + "urlBlockedList.refId.value")
		if t.String() == "global" {
			data.UrlBlockListId = types.StringValue(va.String())
		}
	}
	data.BlockPageAction = types.StringNull()

	if t := res.Get(path + "blockPageAction.optionType"); t.Exists() {
		va := res.Get(path + "blockPageAction.value")
		if t.String() == "global" {
			data.BlockPageAction = types.StringValue(va.String())
		}
	}
	data.BlockPageContents = types.StringNull()

	if t := res.Get(path + "blockPageContents.optionType"); t.Exists() {
		va := res.Get(path + "blockPageContents.value")
		if t.String() == "global" {
			data.BlockPageContents = types.StringValue(va.String())
		}
	}
	data.RedirectUrl = types.StringNull()

	if t := res.Get(path + "redirectUrl.optionType"); t.Exists() {
		va := res.Get(path + "redirectUrl.value")
		if t.String() == "global" {
			data.RedirectUrl = types.StringValue(va.String())
		}
	}
	data.EnableAlerts = types.BoolNull()

	if t := res.Get(path + "enableAlerts.optionType"); t.Exists() {
		va := res.Get(path + "enableAlerts.value")
		if t.String() == "global" {
			data.EnableAlerts = types.BoolValue(va.Bool())
		}
	}
	data.Alerts = types.SetNull(types.StringType)

	if t := res.Get(path + "alerts.optionType"); t.Exists() {
		va := res.Get(path + "alerts.value")
		if t.String() == "global" {
			data.Alerts = helpers.GetStringSet(va.Array())
		}
	}
}

// End of section. //template:end updateFromBody
