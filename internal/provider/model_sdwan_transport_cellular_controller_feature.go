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
type TransportCellularController struct {
	Id                         types.String `tfsdk:"id"`
	Version                    types.Int64  `tfsdk:"version"`
	Name                       types.String `tfsdk:"name"`
	Description                types.String `tfsdk:"description"`
	FeatureProfileId           types.String `tfsdk:"feature_profile_id"`
	CellularId                 types.String `tfsdk:"cellular_id"`
	CellularIdVariable         types.String `tfsdk:"cellular_id_variable"`
	PrimarySimSlot             types.Int64  `tfsdk:"primary_sim_slot"`
	PrimarySimSlotVariable     types.String `tfsdk:"primary_sim_slot_variable"`
	SimFailoverRetries         types.Int64  `tfsdk:"sim_failover_retries"`
	SimFailoverRetriesVariable types.String `tfsdk:"sim_failover_retries_variable"`
	SimFailoverTimeout         types.Int64  `tfsdk:"sim_failover_timeout"`
	SimFailoverTimeoutVariable types.String `tfsdk:"sim_failover_timeout_variable"`
	FirmwareAutoSim            types.Bool   `tfsdk:"firmware_auto_sim"`
	FirmwareAutoSimVariable    types.String `tfsdk:"firmware_auto_sim_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TransportCellularController) getModel() string {
	return "transport_cellular_controller"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportCellularController) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/cellular-controller", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportCellularController) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {
		body, _ = sjson.Set(body, path+"configType.optionType", "default")
		body, _ = sjson.Set(body, path+"configType.value", "non-eSim")
	}

	if !data.CellularIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"controllerConfig.id.optionType", "variable")
			body, _ = sjson.Set(body, path+"controllerConfig.id.value", data.CellularIdVariable.ValueString())
		}
	} else if !data.CellularId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"controllerConfig.id.optionType", "global")
			body, _ = sjson.Set(body, path+"controllerConfig.id.value", data.CellularId.ValueString())
		}
	}

	if !data.PrimarySimSlotVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"controllerConfig.slot.optionType", "variable")
			body, _ = sjson.Set(body, path+"controllerConfig.slot.value", data.PrimarySimSlotVariable.ValueString())
		}
	} else if data.PrimarySimSlot.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"controllerConfig.slot.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"controllerConfig.slot.optionType", "global")
			body, _ = sjson.Set(body, path+"controllerConfig.slot.value", data.PrimarySimSlot.ValueInt64())
		}
	}

	if !data.SimFailoverRetriesVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"controllerConfig.maxRetry.optionType", "variable")
			body, _ = sjson.Set(body, path+"controllerConfig.maxRetry.value", data.SimFailoverRetriesVariable.ValueString())
		}
	} else if data.SimFailoverRetries.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"controllerConfig.maxRetry.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"controllerConfig.maxRetry.optionType", "global")
			body, _ = sjson.Set(body, path+"controllerConfig.maxRetry.value", data.SimFailoverRetries.ValueInt64())
		}
	}

	if !data.SimFailoverTimeoutVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"controllerConfig.failovertimer.optionType", "variable")
			body, _ = sjson.Set(body, path+"controllerConfig.failovertimer.value", data.SimFailoverTimeoutVariable.ValueString())
		}
	} else if data.SimFailoverTimeout.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"controllerConfig.failovertimer.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"controllerConfig.failovertimer.optionType", "global")
			body, _ = sjson.Set(body, path+"controllerConfig.failovertimer.value", data.SimFailoverTimeout.ValueInt64())
		}
	}

	if !data.FirmwareAutoSimVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"controllerConfig.autoSim.optionType", "variable")
			body, _ = sjson.Set(body, path+"controllerConfig.autoSim.value", data.FirmwareAutoSimVariable.ValueString())
		}
	} else if data.FirmwareAutoSim.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"controllerConfig.autoSim.optionType", "default")
			body, _ = sjson.Set(body, path+"controllerConfig.autoSim.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"controllerConfig.autoSim.optionType", "global")
			body, _ = sjson.Set(body, path+"controllerConfig.autoSim.value", data.FirmwareAutoSim.ValueBool())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportCellularController) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.CellularId = types.StringNull()
	data.CellularIdVariable = types.StringNull()
	if t := res.Get(path + "controllerConfig.id.optionType"); t.Exists() {
		va := res.Get(path + "controllerConfig.id.value")
		if t.String() == "variable" {
			data.CellularIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.CellularId = types.StringValue(va.String())
		}
	}
	data.PrimarySimSlot = types.Int64Null()
	data.PrimarySimSlotVariable = types.StringNull()
	if t := res.Get(path + "controllerConfig.slot.optionType"); t.Exists() {
		va := res.Get(path + "controllerConfig.slot.value")
		if t.String() == "variable" {
			data.PrimarySimSlotVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PrimarySimSlot = types.Int64Value(va.Int())
		}
	}
	data.SimFailoverRetries = types.Int64Null()
	data.SimFailoverRetriesVariable = types.StringNull()
	if t := res.Get(path + "controllerConfig.maxRetry.optionType"); t.Exists() {
		va := res.Get(path + "controllerConfig.maxRetry.value")
		if t.String() == "variable" {
			data.SimFailoverRetriesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SimFailoverRetries = types.Int64Value(va.Int())
		}
	}
	data.SimFailoverTimeout = types.Int64Null()
	data.SimFailoverTimeoutVariable = types.StringNull()
	if t := res.Get(path + "controllerConfig.failovertimer.optionType"); t.Exists() {
		va := res.Get(path + "controllerConfig.failovertimer.value")
		if t.String() == "variable" {
			data.SimFailoverTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SimFailoverTimeout = types.Int64Value(va.Int())
		}
	}
	data.FirmwareAutoSim = types.BoolNull()
	data.FirmwareAutoSimVariable = types.StringNull()
	if t := res.Get(path + "controllerConfig.autoSim.optionType"); t.Exists() {
		va := res.Get(path + "controllerConfig.autoSim.value")
		if t.String() == "variable" {
			data.FirmwareAutoSimVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.FirmwareAutoSim = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TransportCellularController) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.CellularId = types.StringNull()
	data.CellularIdVariable = types.StringNull()
	if t := res.Get(path + "controllerConfig.id.optionType"); t.Exists() {
		va := res.Get(path + "controllerConfig.id.value")
		if t.String() == "variable" {
			data.CellularIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.CellularId = types.StringValue(va.String())
		}
	}
	data.PrimarySimSlot = types.Int64Null()
	data.PrimarySimSlotVariable = types.StringNull()
	if t := res.Get(path + "controllerConfig.slot.optionType"); t.Exists() {
		va := res.Get(path + "controllerConfig.slot.value")
		if t.String() == "variable" {
			data.PrimarySimSlotVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PrimarySimSlot = types.Int64Value(va.Int())
		}
	}
	data.SimFailoverRetries = types.Int64Null()
	data.SimFailoverRetriesVariable = types.StringNull()
	if t := res.Get(path + "controllerConfig.maxRetry.optionType"); t.Exists() {
		va := res.Get(path + "controllerConfig.maxRetry.value")
		if t.String() == "variable" {
			data.SimFailoverRetriesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SimFailoverRetries = types.Int64Value(va.Int())
		}
	}
	data.SimFailoverTimeout = types.Int64Null()
	data.SimFailoverTimeoutVariable = types.StringNull()
	if t := res.Get(path + "controllerConfig.failovertimer.optionType"); t.Exists() {
		va := res.Get(path + "controllerConfig.failovertimer.value")
		if t.String() == "variable" {
			data.SimFailoverTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SimFailoverTimeout = types.Int64Value(va.Int())
		}
	}
	data.FirmwareAutoSim = types.BoolNull()
	data.FirmwareAutoSimVariable = types.StringNull()
	if t := res.Get(path + "controllerConfig.autoSim.optionType"); t.Exists() {
		va := res.Get(path + "controllerConfig.autoSim.value")
		if t.String() == "variable" {
			data.FirmwareAutoSimVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.FirmwareAutoSim = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end updateFromBody
