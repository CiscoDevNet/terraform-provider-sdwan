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
type Tag struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Devices     []TagDevices `tfsdk:"devices"`
}

type TagDevices struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data Tag) getPath() string {
	return "/v1/tags"
}

// End of section. //template:end getPath

func (data Tag) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "data", []interface{}{})
		itemBody := ""
		if !data.Name.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name", data.Name.ValueString())
		}
		if !data.Description.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "description", data.Description.ValueString())
		}
		body, _ = sjson.SetRaw(body, "data.-1", itemBody)
	}
	return body
}

func (data Tag) toBodyDeviceAssociation(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "data", []interface{}{})

	if true {
		itemBody := ""
		itemBody, _ = sjson.Set(itemBody, "tagId", data.Id.ValueString())

		if true {
			itemBody, _ = sjson.Set(itemBody, "objects", []interface{}{})
			for _, item := range data.Devices {
				itemChildBody := ""
				if !item.Id.IsNull() {
					itemChildBody, _ = sjson.Set(itemChildBody, "id", item.Id.ValueString())
					itemChildBody, _ = sjson.Set(itemChildBody, "objectType", "DEVICE")
				}
				itemBody, _ = sjson.SetRaw(itemBody, "objects.-1", itemChildBody)
			}
		}
		body, _ = sjson.SetRaw(body, "data.-1", itemBody)
	}
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Tag) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("tagAssociation"); value.Exists() && len(value.Array()) > 0 {
		data.Devices = make([]TagDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TagDevices{}
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
			data.Devices = []TagDevices{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *Tag) hasChanges(ctx context.Context, state *Tag) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
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
