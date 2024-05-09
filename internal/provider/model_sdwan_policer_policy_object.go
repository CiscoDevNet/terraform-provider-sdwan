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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type PolicerPolicyObject struct {
	Id           types.String `tfsdk:"id"`
	Version      types.Int64  `tfsdk:"version"`
	Name         types.String `tfsdk:"name"`
	Burst        types.Int64  `tfsdk:"burst"`
	ExceedAction types.String `tfsdk:"exceed_action"`
	Rate         types.Int64  `tfsdk:"rate"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PolicerPolicyObject) getPath() string {
	return "/template/policy/list/policer/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PolicerPolicyObject) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "type", "policer")
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Burst.IsNull() {
		body, _ = sjson.Set(body, "entries.0.burst", fmt.Sprint(data.Burst.ValueInt64()))
	}
	if !data.ExceedAction.IsNull() {
		body, _ = sjson.Set(body, "entries.0.exceed", data.ExceedAction.ValueString())
	}
	if !data.Rate.IsNull() {
		body, _ = sjson.Set(body, "entries.0.rate", fmt.Sprint(data.Rate.ValueInt64()))
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PolicerPolicyObject) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("entries.0.burst"); value.Exists() {
		data.Burst = types.Int64Value(value.Int())
	} else {
		data.Burst = types.Int64Null()
	}
	if value := res.Get("entries.0.exceed"); value.Exists() {
		data.ExceedAction = types.StringValue(value.String())
	} else {
		data.ExceedAction = types.StringNull()
	}
	if value := res.Get("entries.0.rate"); value.Exists() {
		data.Rate = types.Int64Value(value.Int())
	} else {
		data.Rate = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *PolicerPolicyObject) hasChanges(ctx context.Context, state *PolicerPolicyObject) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Burst.Equal(state.Burst) {
		hasChanges = true
	}
	if !data.ExceedAction.Equal(state.ExceedAction) {
		hasChanges = true
	}
	if !data.Rate.Equal(state.Rate) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

// End of section. //template:end updateVersions
