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
type CiscoBanner struct {
	Id            types.String `tfsdk:"id"`
	Version       types.Int64  `tfsdk:"version"`
	TemplateType  types.String `tfsdk:"template_type"`
	Name          types.String `tfsdk:"name"`
	Description   types.String `tfsdk:"description"`
	DeviceTypes   types.Set    `tfsdk:"device_types"`
	Login         types.String `tfsdk:"login"`
	LoginVariable types.String `tfsdk:"login_variable"`
	Motd          types.String `tfsdk:"motd"`
	MotdVariable  types.String `tfsdk:"motd_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoBanner) getModel() string {
	return "cisco_banner"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CiscoBanner) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_banner")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.LoginVariable.IsNull() {
		body, _ = sjson.Set(body, path+"login."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"login."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"login."+"vipVariableName", data.LoginVariable.ValueString())
	} else if data.Login.IsNull() {
		body, _ = sjson.Set(body, path+"login."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"login."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"login."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"login."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"login."+"vipValue", data.Login.ValueString())
	}

	if !data.MotdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"motd."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"motd."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"motd."+"vipVariableName", data.MotdVariable.ValueString())
	} else if data.Motd.IsNull() {
		body, _ = sjson.Set(body, path+"motd."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"motd."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"motd."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"motd."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"motd."+"vipValue", data.Motd.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoBanner) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceType"); value.Exists() {
		data.DeviceTypes = helpers.GetStringSet(value.Array())
	} else {
		data.DeviceTypes = types.SetNull(types.StringType)
	}
	if value := res.Get("templateDescription"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("templateName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("templateType"); value.Exists() {
		data.TemplateType = types.StringValue(value.String())
	} else {
		data.TemplateType = types.StringNull()
	}

	path := "templateDefinition."
	if value := res.Get(path + "login.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Login = types.StringNull()

			v := res.Get(path + "login.vipVariableName")
			data.LoginVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Login = types.StringNull()
			data.LoginVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "login.vipValue")
			data.Login = types.StringValue(v.String())
			data.LoginVariable = types.StringNull()
		}
	} else {
		data.Login = types.StringNull()
		data.LoginVariable = types.StringNull()
	}
	if value := res.Get(path + "motd.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Motd = types.StringNull()

			v := res.Get(path + "motd.vipVariableName")
			data.MotdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Motd = types.StringNull()
			data.MotdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "motd.vipValue")
			data.Motd = types.StringValue(v.String())
			data.MotdVariable = types.StringNull()
		}
	} else {
		data.Motd = types.StringNull()
		data.MotdVariable = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoBanner) hasChanges(ctx context.Context, state *CiscoBanner) bool {
	hasChanges := false
	if !data.Login.Equal(state.Login) {
		hasChanges = true
	}
	if !data.Motd.Equal(state.Motd) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
