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
type SystemBanner struct {
	Id               types.String `tfsdk:"id"`
	Version          types.Int64  `tfsdk:"version"`
	Name             types.String `tfsdk:"name"`
	Description      types.String `tfsdk:"description"`
	FeatureProfileId types.String `tfsdk:"feature_profile_id"`
	Login            types.String `tfsdk:"login"`
	LoginVariable    types.String `tfsdk:"login_variable"`
	Motd             types.String `tfsdk:"motd"`
	MotdVariable     types.String `tfsdk:"motd_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SystemBanner) getModel() string {
	return "system_banner"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SystemBanner) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/banner", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SystemBanner) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.LoginVariable.IsNull() {
		body, _ = sjson.Set(body, path+"login.optionType", "variable")
		body, _ = sjson.Set(body, path+"login.value", data.LoginVariable.ValueString())
	} else if data.Login.IsNull() {
		body, _ = sjson.Set(body, path+"login.optionType", "default")
		body, _ = sjson.Set(body, path+"login.value", "")
	} else {
		body, _ = sjson.Set(body, path+"login.optionType", "global")
		body, _ = sjson.Set(body, path+"login.value", data.Login.ValueString())
	}

	if !data.MotdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"motd.optionType", "variable")
		body, _ = sjson.Set(body, path+"motd.value", data.MotdVariable.ValueString())
	} else if data.Motd.IsNull() {
		body, _ = sjson.Set(body, path+"motd.optionType", "default")
		body, _ = sjson.Set(body, path+"motd.value", "")
	} else {
		body, _ = sjson.Set(body, path+"motd.optionType", "global")
		body, _ = sjson.Set(body, path+"motd.value", data.Motd.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemBanner) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Login = types.StringNull()
	data.LoginVariable = types.StringNull()
	if t := res.Get(path + "login.optionType"); t.Exists() {
		va := res.Get(path + "login.value")
		if t.String() == "variable" {
			data.LoginVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Login = types.StringValue(va.String())
		}
	}
	data.Motd = types.StringNull()
	data.MotdVariable = types.StringNull()
	if t := res.Get(path + "motd.optionType"); t.Exists() {
		va := res.Get(path + "motd.value")
		if t.String() == "variable" {
			data.MotdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Motd = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SystemBanner) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Login = types.StringNull()
	data.LoginVariable = types.StringNull()
	if t := res.Get(path + "login.optionType"); t.Exists() {
		va := res.Get(path + "login.value")
		if t.String() == "variable" {
			data.LoginVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Login = types.StringValue(va.String())
		}
	}
	data.Motd = types.StringNull()
	data.MotdVariable = types.StringNull()
	if t := res.Get(path + "motd.optionType"); t.Exists() {
		va := res.Get(path + "motd.value")
		if t.String() == "variable" {
			data.MotdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Motd = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *SystemBanner) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.Login.IsNull() {
		return false
	}
	if !data.LoginVariable.IsNull() {
		return false
	}
	if !data.Motd.IsNull() {
		return false
	}
	if !data.MotdVariable.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
