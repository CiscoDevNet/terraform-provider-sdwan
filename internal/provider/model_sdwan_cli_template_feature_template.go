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
type CLITemplate struct {
	Id           types.String `tfsdk:"id"`
	Version      types.Int64  `tfsdk:"version"`
	TemplateType types.String `tfsdk:"template_type"`
	Name         types.String `tfsdk:"name"`
	Description  types.String `tfsdk:"description"`
	DeviceTypes  types.Set    `tfsdk:"device_types"`
	CliConfig    types.String `tfsdk:"cli_config"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CLITemplate) getModel() string {
	return "cli-template"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CLITemplate) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cli-template")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."
	if data.CliConfig.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"config."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"config."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"config."+"vipValue", data.CliConfig.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CLITemplate) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "config.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.CliConfig = types.StringNull()

		} else if value.String() == "ignore" {
			data.CliConfig = types.StringNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "config.vipValue")
			data.CliConfig = types.StringValue(v.String())

		}
	} else {
		data.CliConfig = types.StringNull()

	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CLITemplate) hasChanges(ctx context.Context, state *CLITemplate) bool {
	hasChanges := false
	if !data.CliConfig.Equal(state.CliConfig) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
