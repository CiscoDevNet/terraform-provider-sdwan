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

type VEdgeInventory struct {
	Id      types.String            `tfsdk:"id"`
	Devices []VEdgeInventoryDevices `tfsdk:"devices"`
}

type VEdgeInventoryDevices struct {
	ChassisNumber types.String `tfsdk:"chassis_number"`
	SiteId        types.String `tfsdk:"site_id"`
	SerialNumber  types.String `tfsdk:"serial_number"`
	Hostname      types.String `tfsdk:"hostname"`
	Validity      types.String `tfsdk:"validity"`
	DeviceType    types.String `tfsdk:"device_type"`
}

func (data VEdgeInventory) getPath() string {
	return "/device/vedgeinventory/detail"
}

func (data VEdgeInventory) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "data", []interface{}{})
		for _, item := range data.Devices {
			itemBody := ""
			if !item.ChassisNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "chasisNumber", item.ChassisNumber.ValueString())
			}
			if !item.SiteId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "site-id", item.SiteId.ValueString())
			}
			if !item.SerialNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "serialNumber", item.SerialNumber.ValueString())
			}
			if !item.Hostname.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "host-name", item.Hostname.ValueString())
			}
			if !item.Validity.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "validity", item.Validity.ValueString())
			}
			if !item.DeviceType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deviceType", item.DeviceType.ValueString())
			}
			body, _ = sjson.SetRaw(body, "data.-1", itemBody)
		}
	}
	return body
}

func (data *VEdgeInventory) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("data"); value.Exists() && len(value.Array()) > 0 {
		data.Devices = make([]VEdgeInventoryDevices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VEdgeInventoryDevices{}
			if cValue := v.Get("chasisNumber"); cValue.Exists() {
				item.ChassisNumber = types.StringValue(cValue.String())
			} else {
				item.ChassisNumber = types.StringNull()
			}
			if cValue := v.Get("site-id"); cValue.Exists() {
				item.SiteId = types.StringValue(cValue.String())
			} else {
				item.SiteId = types.StringNull()
			}
			if cValue := v.Get("serialNumber"); cValue.Exists() {
				item.SerialNumber = types.StringValue(cValue.String())
			} else {
				item.SerialNumber = types.StringNull()
			}
			if cValue := v.Get("host-name"); cValue.Exists() {
				item.Hostname = types.StringValue(cValue.String())
			} else {
				item.Hostname = types.StringNull()
			}
			if cValue := v.Get("validity"); cValue.Exists() {
				item.Validity = types.StringValue(cValue.String())
			} else {
				item.Validity = types.StringNull()
			}
			if cValue := v.Get("deviceType"); cValue.Exists() {
				item.DeviceType = types.StringValue(cValue.String())
			} else {
				item.DeviceType = types.StringNull()
			}
			data.Devices = append(data.Devices, item)
			return true
		})
	} else {
		if len(data.Devices) > 0 {
			data.Devices = []VEdgeInventoryDevices{}
		}
	}
}

func (data *VEdgeInventory) hasChanges(ctx context.Context, state *VEdgeInventory) bool {
	hasChanges := false
	if len(data.Devices) != len(state.Devices) {
		hasChanges = true
	} else {
		for i := range data.Devices {
			if !data.Devices[i].ChassisNumber.Equal(state.Devices[i].ChassisNumber) {
				hasChanges = true
			}
			if !data.Devices[i].SiteId.Equal(state.Devices[i].SiteId) {
				hasChanges = true
			}
			if !data.Devices[i].SerialNumber.Equal(state.Devices[i].SerialNumber) {
				hasChanges = true
			}
			if !data.Devices[i].Hostname.Equal(state.Devices[i].Hostname) {
				hasChanges = true
			}
			if !data.Devices[i].Validity.Equal(state.Devices[i].Validity) {
				hasChanges = true
			}
			if !data.Devices[i].DeviceType.Equal(state.Devices[i].DeviceType) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}
