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
type PolicyGroup struct {
	Id                types.String `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	Description       types.String `tfsdk:"description"`
	Solution          types.String `tfsdk:"solution"`
	FeatureProfileIds types.Set    `tfsdk:"feature_profile_ids"`
	PolicyVersions    types.List   `tfsdk:"policy_versions"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicyGroup) getPath() string {
	return "/v1/policy-group/"
}

// End of section. //template:end getPath

func (data PolicyGroup) toBodyPolicyGroup(ctx context.Context) string {
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
	return body
}

func (data *PolicyGroup) fromBodyPolicyGroup(ctx context.Context, res gjson.Result) {
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
}
