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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ApplicationPriorityPolicySettings struct {
	Id                        types.String `tfsdk:"id"`
	Version                   types.Int64  `tfsdk:"version"`
	Name                      types.String `tfsdk:"name"`
	Description               types.String `tfsdk:"description"`
	FeatureProfileId          types.String `tfsdk:"feature_profile_id"`
	Ipv4ApplicationVisibility types.Bool   `tfsdk:"ipv4_application_visibility"`
	Ipv6ApplicationVisibility types.Bool   `tfsdk:"ipv6_application_visibility"`
	Ipv4FlowVisibility        types.Bool   `tfsdk:"ipv4_flow_visibility"`
	Ipv6FlowVisibility        types.Bool   `tfsdk:"ipv6_flow_visibility"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ApplicationPriorityPolicySettings) getModel() string {
	return "application_priority_policy_settings"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ApplicationPriorityPolicySettings) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/application-priority/%v/policy-settings", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

func (data ApplicationPriorityPolicySettings) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if data.Ipv4ApplicationVisibility.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"appVisibility.optionType", "default")
			body, _ = sjson.Set(body, path+"appVisibility.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"appVisibility.optionType", "global")
			body, _ = sjson.Set(body, path+"appVisibility.value", data.Ipv4ApplicationVisibility.ValueBool())
		}
	}
	if data.Ipv6ApplicationVisibility.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"appVisibilityIPv6.optionType", "default")
			body, _ = sjson.Set(body, path+"appVisibilityIPv6.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"appVisibilityIPv6.optionType", "global")
			body, _ = sjson.Set(body, path+"appVisibilityIPv6.value", data.Ipv6ApplicationVisibility.ValueBool())
		}
	}
	if data.Ipv4FlowVisibility.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"flowVisibility.optionType", "default")
			body, _ = sjson.Set(body, path+"flowVisibility.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"flowVisibility.optionType", "global")
			body, _ = sjson.Set(body, path+"flowVisibility.value", data.Ipv4FlowVisibility.ValueBool())
		}
	}
	if data.Ipv6FlowVisibility.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"flowVisibilityIPv6.optionType", "default")
			body, _ = sjson.Set(body, path+"flowVisibilityIPv6.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"flowVisibilityIPv6.optionType", "global")
			body, _ = sjson.Set(body, path+"flowVisibilityIPv6.value", data.Ipv6FlowVisibility.ValueBool())
		}
	}
	// Section below is manually configured
	body, _ = sjson.Set(body, path+"cflowd.optionType", "network-settings")
	body, _ = sjson.Set(body, path+"cflowd.value", true)
	// Section below is manually configured

	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ApplicationPriorityPolicySettings) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Ipv4ApplicationVisibility = types.BoolNull()

	if t := res.Get(path + "appVisibility.optionType"); t.Exists() {
		va := res.Get(path + "appVisibility.value")
		if t.String() == "global" {
			data.Ipv4ApplicationVisibility = types.BoolValue(va.Bool())
		}
	}
	data.Ipv6ApplicationVisibility = types.BoolNull()

	if t := res.Get(path + "appVisibilityIPv6.optionType"); t.Exists() {
		va := res.Get(path + "appVisibilityIPv6.value")
		if t.String() == "global" {
			data.Ipv6ApplicationVisibility = types.BoolValue(va.Bool())
		}
	}
	data.Ipv4FlowVisibility = types.BoolNull()

	if t := res.Get(path + "flowVisibility.optionType"); t.Exists() {
		va := res.Get(path + "flowVisibility.value")
		if t.String() == "global" {
			data.Ipv4FlowVisibility = types.BoolValue(va.Bool())
		}
	}
	data.Ipv6FlowVisibility = types.BoolNull()

	if t := res.Get(path + "flowVisibilityIPv6.optionType"); t.Exists() {
		va := res.Get(path + "flowVisibilityIPv6.value")
		if t.String() == "global" {
			data.Ipv6FlowVisibility = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end fromBody
