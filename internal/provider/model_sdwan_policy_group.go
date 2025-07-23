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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
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

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicyGroup) toBody(ctx context.Context) string {
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
	if !data.FeatureProfileIds.IsNull() {
		var values []string
		data.FeatureProfileIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "profiles", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicyGroup) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("solution"); value.Exists() {
		data.Solution = types.StringValue(value.String())
	} else {
		data.Solution = types.StringNull()
	}
	if value := res.Get("profiles"); value.Exists() {
		data.FeatureProfileIds = helpers.GetStringSet(value.Array())
	} else {
		data.FeatureProfileIds = types.SetNull(types.StringType)
	}
	data.updateVersions(ctx, &state)
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *PolicyGroup) hasChanges(ctx context.Context, state *PolicyGroup) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.Solution.Equal(state.Solution) {
		hasChanges = true
	}
	if !data.FeatureProfileIds.Equal(state.FeatureProfileIds) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

func (data *PolicyGroup) updateVersions(ctx context.Context, state *PolicyGroup) {
	data.PolicyVersions = state.PolicyVersions
}

// End of section. //template:end updateVersions
