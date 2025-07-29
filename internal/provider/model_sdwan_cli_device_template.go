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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type CLIDeviceTemplate struct {
	Id               types.String `tfsdk:"id"`
	Version          types.Int64  `tfsdk:"version"`
	ConfigType       types.String `tfsdk:"config_type"`
	FactoryDefault   types.Bool   `tfsdk:"factory_default"`
	Name             types.String `tfsdk:"name"`
	Description      types.String `tfsdk:"description"`
	DeviceType       types.String `tfsdk:"device_type"`
	CliType          types.String `tfsdk:"cli_type"`
	CliConfiguration types.String `tfsdk:"cli_configuration"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data CLIDeviceTemplate) getPath() string {
	return "/template/device/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CLIDeviceTemplate) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "configType", "file")
	}
	if true {
		body, _ = sjson.Set(body, "factoryDefault", false)
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	}
	if !data.DeviceType.IsNull() {
		body, _ = sjson.Set(body, "deviceType", data.DeviceType.ValueString())
	}
	if !data.CliType.IsNull() {
		body, _ = sjson.Set(body, "cliType", data.CliType.ValueString())
	}
	if !data.CliConfiguration.IsNull() {
		body, _ = sjson.Set(body, "templateConfiguration", data.CliConfiguration.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CLIDeviceTemplate) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("templateName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("templateDescription"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("deviceType"); value.Exists() {
		data.DeviceType = types.StringValue(value.String())
	} else {
		data.DeviceType = types.StringNull()
	}
	if value := res.Get("cliType"); value.Exists() {
		data.CliType = types.StringValue(value.String())
	} else {
		data.CliType = types.StringNull()
	}
	if value := res.Get("templateConfiguration"); value.Exists() {
		data.CliConfiguration = types.StringValue(value.String())
	} else {
		data.CliConfiguration = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CLIDeviceTemplate) hasChanges(ctx context.Context, state *CLIDeviceTemplate) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.DeviceType.Equal(state.DeviceType) {
		hasChanges = true
	}
	if !data.CliType.Equal(state.CliType) {
		hasChanges = true
	}
	if !data.CliConfiguration.Equal(state.CliConfiguration) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *CLIDeviceTemplate) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
	data.ConfigType = types.StringValue("file")
	data.FactoryDefault = types.BoolValue(false)
}

// End of section. //template:end processImport
