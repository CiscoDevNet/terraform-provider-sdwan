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
type ConfigurationGroupDeploy struct {
	Id                   types.String                      `tfsdk:"id"`
	ConfigurationGroupId types.String                      `tfsdk:"configuration_group_id"`
	Devices              []ConfigurationGroupDeployDevices `tfsdk:"devices"`
}

type ConfigurationGroupDeployDevices struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ConfigurationGroupDeploy) getPath() string {
	return fmt.Sprintf("/v1/config-group/%v/device/deploy/", url.QueryEscape(data.ConfigurationGroupId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ConfigurationGroupDeploy) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "devices", []interface{}{})
		for _, item := range data.Devices {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "devices.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ConfigurationGroupDeploy) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("devices"); value.Exists() && len(value.Array()) > 0 {
		data.Devices = make([]ConfigurationGroupDeployDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ConfigurationGroupDeployDevices{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			data.Devices = append(data.Devices, item)
			return true
		})
	} else {
		if len(data.Devices) > 0 {
			data.Devices = []ConfigurationGroupDeployDevices{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *ConfigurationGroupDeploy) hasChanges(ctx context.Context, state *ConfigurationGroupDeploy) bool {
	hasChanges := false
	if !data.ConfigurationGroupId.Equal(state.ConfigurationGroupId) {
		hasChanges = true
	}
	if len(data.Devices) != len(state.Devices) {
		hasChanges = true
	} else {
		for i := range data.Devices {
			if !data.Devices[i].Id.Equal(state.Devices[i].Id) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

// End of section. //template:end updateVersions
