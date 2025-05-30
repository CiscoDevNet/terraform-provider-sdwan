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
type PolicyObjectExpandedCommunityList struct {
	Id                             types.String `tfsdk:"id"`
	Version                        types.Int64  `tfsdk:"version"`
	Name                           types.String `tfsdk:"name"`
	Description                    types.String `tfsdk:"description"`
	FeatureProfileId               types.String `tfsdk:"feature_profile_id"`
	ExpandedCommunityLists         types.Set    `tfsdk:"expanded_community_lists"`
	ExpandedCommunityListsVariable types.String `tfsdk:"expanded_community_lists_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data PolicyObjectExpandedCommunityList) getModel() string {
	return "policy_object_expanded_community_list"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyObjectExpandedCommunityList) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/policy-object/%v/expanded-community", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyObjectExpandedCommunityList) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.ExpandedCommunityListsVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"expandedCommunityList.optionType", "variable")
			body, _ = sjson.Set(body, path+"expandedCommunityList.value", data.ExpandedCommunityListsVariable.ValueString())
		}
	} else if !data.ExpandedCommunityLists.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"expandedCommunityList.optionType", "global")
			var values []string
			data.ExpandedCommunityLists.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"expandedCommunityList.value", values)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyObjectExpandedCommunityList) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.ExpandedCommunityLists = types.SetNull(types.StringType)
	data.ExpandedCommunityListsVariable = types.StringNull()
	if t := res.Get(path + "expandedCommunityList.optionType"); t.Exists() {
		va := res.Get(path + "expandedCommunityList.value")
		if t.String() == "variable" {
			data.ExpandedCommunityListsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ExpandedCommunityLists = helpers.GetStringSet(va.Array())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PolicyObjectExpandedCommunityList) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.ExpandedCommunityLists = types.SetNull(types.StringType)
	data.ExpandedCommunityListsVariable = types.StringNull()
	if t := res.Get(path + "expandedCommunityList.optionType"); t.Exists() {
		va := res.Get(path + "expandedCommunityList.value")
		if t.String() == "variable" {
			data.ExpandedCommunityListsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ExpandedCommunityLists = helpers.GetStringSet(va.Array())
		}
	}
}

// End of section. //template:end updateFromBody
