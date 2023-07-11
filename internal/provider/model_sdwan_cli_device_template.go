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

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CLIDeviceTemplate struct {
	Id               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Description      types.String `tfsdk:"description"`
	DeviceType       types.String `tfsdk:"device_type"`
	CLIType          types.String `tfsdk:"cli_type"`
	CLIConfiguration types.String `tfsdk:"cli_configuration"`
}

func (data CLIDeviceTemplate) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "deviceType", data.DeviceType.ValueString())
	body, _ = sjson.Set(body, "templateConfiguration", data.CLIConfiguration.ValueString())
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "configType", "file")
	body, _ = sjson.Set(body, "cliType", data.CLIType.ValueString())
	body, _ = sjson.Set(body, "draftMode", false)
	return body
}

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
	if value := res.Get("templateConfiguration"); value.Exists() {
		data.CLIConfiguration = types.StringValue(value.String())
	} else {
		data.CLIConfiguration = types.StringNull()
	}
	if value := res.Get("cliType"); value.Exists() {
		data.CLIType = types.StringValue(value.String())
	} else {
		data.CLIType = types.StringNull()
	}
}
