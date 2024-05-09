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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type Gps struct {
	Id                         types.String `tfsdk:"id"`
	Version                    types.Int64  `tfsdk:"version"`
	TemplateType               types.String `tfsdk:"template_type"`
	Name                       types.String `tfsdk:"name"`
	Description                types.String `tfsdk:"description"`
	DeviceTypes                types.Set    `tfsdk:"device_types"`
	Enable                     types.Bool   `tfsdk:"enable"`
	EnableVariable             types.String `tfsdk:"enable_variable"`
	GpsMode                    types.String `tfsdk:"gps_mode"`
	GpsModeVariable            types.String `tfsdk:"gps_mode_variable"`
	Nmea                       types.Bool   `tfsdk:"nmea"`
	NmeaVariable               types.String `tfsdk:"nmea_variable"`
	SourceAddress              types.String `tfsdk:"source_address"`
	SourceAddressVariable      types.String `tfsdk:"source_address_variable"`
	DestinationAddress         types.String `tfsdk:"destination_address"`
	DestinationAddressVariable types.String `tfsdk:"destination_address_variable"`
	DestinationPort            types.Int64  `tfsdk:"destination_port"`
	DestinationPortVariable    types.String `tfsdk:"destination_port_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data Gps) getModel() string {
	return "cellular-cedge-gps-controller"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data Gps) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cellular-cedge-gps-controller")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.EnableVariable.IsNull() {
		body, _ = sjson.Set(body, path+"gps.enable."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"gps.enable."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"gps.enable."+"vipVariableName", data.EnableVariable.ValueString())
	} else if data.Enable.IsNull() {
		body, _ = sjson.Set(body, path+"gps.enable."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"gps.enable."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"gps.enable."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"gps.enable."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"gps.enable."+"vipValue", strconv.FormatBool(data.Enable.ValueBool()))
	}

	if !data.GpsModeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"gps.mode."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps.mode."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"gps.mode."+"vipVariableName", data.GpsModeVariable.ValueString())
	} else if data.GpsMode.IsNull() {
		body, _ = sjson.Set(body, path+"gps.mode."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps.mode."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"gps.mode."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps.mode."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"gps.mode."+"vipValue", data.GpsMode.ValueString())
	}

	if !data.NmeaVariable.IsNull() {
		body, _ = sjson.Set(body, path+"gps.nmea-conf.nmea."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"gps.nmea-conf.nmea."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"gps.nmea-conf.nmea."+"vipVariableName", data.NmeaVariable.ValueString())
	} else if data.Nmea.IsNull() {
		body, _ = sjson.Set(body, path+"gps.nmea-conf.nmea."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"gps.nmea-conf.nmea."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"gps.nmea-conf.nmea."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"gps.nmea-conf.nmea."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"gps.nmea-conf.nmea."+"vipValue", strconv.FormatBool(data.Nmea.ValueBool()))
	}

	if !data.SourceAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.source-address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.source-address."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.source-address."+"vipVariableName", data.SourceAddressVariable.ValueString())
	} else if data.SourceAddress.IsNull() {
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.source-address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.source-address."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.source-address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.source-address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.source-address."+"vipValue", data.SourceAddress.ValueString())
	}

	if !data.DestinationAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.destination-address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.destination-address."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.destination-address."+"vipVariableName", data.DestinationAddressVariable.ValueString())
	} else if data.DestinationAddress.IsNull() {
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.destination-address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.destination-address."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.destination-address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.destination-address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.destination-address."+"vipValue", data.DestinationAddress.ValueString())
	}

	if !data.DestinationPortVariable.IsNull() {
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.destination-port."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.destination-port."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.destination-port."+"vipVariableName", data.DestinationPortVariable.ValueString())
	} else if data.DestinationPort.IsNull() {
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.destination-port."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.destination-port."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.destination-port."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.destination-port."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"gps.nmea.ip.udp.destination-port."+"vipValue", data.DestinationPort.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *Gps) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "gps.enable.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Enable = types.BoolNull()

			v := res.Get(path + "gps.enable.vipVariableName")
			data.EnableVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Enable = types.BoolNull()
			data.EnableVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "gps.enable.vipValue")
			data.Enable = types.BoolValue(v.Bool())
			data.EnableVariable = types.StringNull()
		}
	} else {
		data.Enable = types.BoolNull()
		data.EnableVariable = types.StringNull()
	}
	if value := res.Get(path + "gps.mode.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.GpsMode = types.StringNull()

			v := res.Get(path + "gps.mode.vipVariableName")
			data.GpsModeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.GpsMode = types.StringNull()
			data.GpsModeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "gps.mode.vipValue")
			data.GpsMode = types.StringValue(v.String())
			data.GpsModeVariable = types.StringNull()
		}
	} else {
		data.GpsMode = types.StringNull()
		data.GpsModeVariable = types.StringNull()
	}
	if value := res.Get(path + "gps.nmea-conf.nmea.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Nmea = types.BoolNull()

			v := res.Get(path + "gps.nmea-conf.nmea.vipVariableName")
			data.NmeaVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Nmea = types.BoolNull()
			data.NmeaVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "gps.nmea-conf.nmea.vipValue")
			data.Nmea = types.BoolValue(v.Bool())
			data.NmeaVariable = types.StringNull()
		}
	} else {
		data.Nmea = types.BoolNull()
		data.NmeaVariable = types.StringNull()
	}
	if value := res.Get(path + "gps.nmea.ip.udp.source-address.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SourceAddress = types.StringNull()

			v := res.Get(path + "gps.nmea.ip.udp.source-address.vipVariableName")
			data.SourceAddressVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SourceAddress = types.StringNull()
			data.SourceAddressVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "gps.nmea.ip.udp.source-address.vipValue")
			data.SourceAddress = types.StringValue(v.String())
			data.SourceAddressVariable = types.StringNull()
		}
	} else {
		data.SourceAddress = types.StringNull()
		data.SourceAddressVariable = types.StringNull()
	}
	if value := res.Get(path + "gps.nmea.ip.udp.destination-address.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DestinationAddress = types.StringNull()

			v := res.Get(path + "gps.nmea.ip.udp.destination-address.vipVariableName")
			data.DestinationAddressVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DestinationAddress = types.StringNull()
			data.DestinationAddressVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "gps.nmea.ip.udp.destination-address.vipValue")
			data.DestinationAddress = types.StringValue(v.String())
			data.DestinationAddressVariable = types.StringNull()
		}
	} else {
		data.DestinationAddress = types.StringNull()
		data.DestinationAddressVariable = types.StringNull()
	}
	if value := res.Get(path + "gps.nmea.ip.udp.destination-port.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DestinationPort = types.Int64Null()

			v := res.Get(path + "gps.nmea.ip.udp.destination-port.vipVariableName")
			data.DestinationPortVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DestinationPort = types.Int64Null()
			data.DestinationPortVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "gps.nmea.ip.udp.destination-port.vipValue")
			data.DestinationPort = types.Int64Value(v.Int())
			data.DestinationPortVariable = types.StringNull()
		}
	} else {
		data.DestinationPort = types.Int64Null()
		data.DestinationPortVariable = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *Gps) hasChanges(ctx context.Context, state *Gps) bool {
	hasChanges := false
	if !data.Enable.Equal(state.Enable) {
		hasChanges = true
	}
	if !data.GpsMode.Equal(state.GpsMode) {
		hasChanges = true
	}
	if !data.Nmea.Equal(state.Nmea) {
		hasChanges = true
	}
	if !data.SourceAddress.Equal(state.SourceAddress) {
		hasChanges = true
	}
	if !data.DestinationAddress.Equal(state.DestinationAddress) {
		hasChanges = true
	}
	if !data.DestinationPort.Equal(state.DestinationPort) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
