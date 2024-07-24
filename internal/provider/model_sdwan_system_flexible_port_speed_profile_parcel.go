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
type SystemFlexiblePortSpeed struct {
	Id               types.String `tfsdk:"id"`
	Version          types.Int64  `tfsdk:"version"`
	Name             types.String `tfsdk:"name"`
	Description      types.String `tfsdk:"description"`
	FeatureProfileId types.String `tfsdk:"feature_profile_id"`
	PortType         types.String `tfsdk:"port_type"`
	PortTypeVariable types.String `tfsdk:"port_type_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SystemFlexiblePortSpeed) getModel() string {
	return "system_flexible_port_speed"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SystemFlexiblePortSpeed) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/flexible-port-speed", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SystemFlexiblePortSpeed) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.PortTypeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"portType.optionType", "variable")
			body, _ = sjson.Set(body, path+"portType.value", data.PortTypeVariable.ValueString())
		}
	} else if data.PortType.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"portType.optionType", "default")
			body, _ = sjson.Set(body, path+"portType.value", "12 ports of 1/10GE + 3 ports 40GE")
		}
	} else {
		body, _ = sjson.Set(body, path+"portType.optionType", "global")
		body, _ = sjson.Set(body, path+"portType.value", data.PortType.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemFlexiblePortSpeed) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.PortType = types.StringNull()
	data.PortTypeVariable = types.StringNull()
	if t := res.Get(path + "portType.optionType"); t.Exists() {
		va := res.Get(path + "portType.value")
		if t.String() == "variable" {
			data.PortTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PortType = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SystemFlexiblePortSpeed) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.PortType = types.StringNull()
	data.PortTypeVariable = types.StringNull()
	if t := res.Get(path + "portType.optionType"); t.Exists() {
		va := res.Get(path + "portType.value")
		if t.String() == "variable" {
			data.PortTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PortType = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *SystemFlexiblePortSpeed) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.PortType.IsNull() {
		return false
	}
	if !data.PortTypeVariable.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
