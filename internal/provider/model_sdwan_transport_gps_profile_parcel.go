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
type TransportGPS struct {
	Id                             types.String `tfsdk:"id"`
	Version                        types.Int64  `tfsdk:"version"`
	Name                           types.String `tfsdk:"name"`
	Description                    types.String `tfsdk:"description"`
	FeatureProfileId               types.String `tfsdk:"feature_profile_id"`
	GpsEnable                      types.Bool   `tfsdk:"gps_enable"`
	GpsEnableVariable              types.String `tfsdk:"gps_enable_variable"`
	GpsMode                        types.String `tfsdk:"gps_mode"`
	GpsModeVariable                types.String `tfsdk:"gps_mode_variable"`
	NmeaEnable                     types.Bool   `tfsdk:"nmea_enable"`
	NmeaEnableVariable             types.String `tfsdk:"nmea_enable_variable"`
	NmeaSourceAddress              types.String `tfsdk:"nmea_source_address"`
	NmeaSourceAddressVariable      types.String `tfsdk:"nmea_source_address_variable"`
	NmeaDestinationAddress         types.String `tfsdk:"nmea_destination_address"`
	NmeaDestinationAddressVariable types.String `tfsdk:"nmea_destination_address_variable"`
	NmeaDestinationPort            types.Int64  `tfsdk:"nmea_destination_port"`
	NmeaDestinationPortVariable    types.String `tfsdk:"nmea_destination_port_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TransportGPS) getModel() string {
	return "transport_gps"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportGPS) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/gps", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportGPS) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.GpsEnableVariable.IsNull() {
		body, _ = sjson.Set(body, path+"enable.optionType", "variable")
		body, _ = sjson.Set(body, path+"enable.value", data.GpsEnableVariable.ValueString())
	} else if data.GpsEnable.IsNull() {
		body, _ = sjson.Set(body, path+"enable.optionType", "default")
		body, _ = sjson.Set(body, path+"enable.value", false)
	} else {
		body, _ = sjson.Set(body, path+"enable.optionType", "global")
		body, _ = sjson.Set(body, path+"enable.value", data.GpsEnable.ValueBool())
	}

	if !data.GpsModeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"mode.optionType", "variable")
		body, _ = sjson.Set(body, path+"mode.value", data.GpsModeVariable.ValueString())
	} else if data.GpsMode.IsNull() {
		body, _ = sjson.Set(body, path+"mode.optionType", "default")
		body, _ = sjson.Set(body, path+"mode.value", "ms-based")
	} else {
		body, _ = sjson.Set(body, path+"mode.optionType", "global")
		body, _ = sjson.Set(body, path+"mode.value", data.GpsMode.ValueString())
	}

	if !data.NmeaEnableVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nmea.optionType", "variable")
		body, _ = sjson.Set(body, path+"nmea.value", data.NmeaEnableVariable.ValueString())
	} else if data.NmeaEnable.IsNull() {
		body, _ = sjson.Set(body, path+"nmea.optionType", "default")
		body, _ = sjson.Set(body, path+"nmea.value", false)
	} else {
		body, _ = sjson.Set(body, path+"nmea.optionType", "global")
		body, _ = sjson.Set(body, path+"nmea.value", data.NmeaEnable.ValueBool())
	}

	if !data.NmeaSourceAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"sourceAddress.optionType", "variable")
		body, _ = sjson.Set(body, path+"sourceAddress.value", data.NmeaSourceAddressVariable.ValueString())
	} else if data.NmeaSourceAddress.IsNull() {
		body, _ = sjson.Set(body, path+"sourceAddress.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"sourceAddress.optionType", "global")
		body, _ = sjson.Set(body, path+"sourceAddress.value", data.NmeaSourceAddress.ValueString())
	}

	if !data.NmeaDestinationAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"destinationAddress.optionType", "variable")
		body, _ = sjson.Set(body, path+"destinationAddress.value", data.NmeaDestinationAddressVariable.ValueString())
	} else if data.NmeaDestinationAddress.IsNull() {
		body, _ = sjson.Set(body, path+"destinationAddress.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"destinationAddress.optionType", "global")
		body, _ = sjson.Set(body, path+"destinationAddress.value", data.NmeaDestinationAddress.ValueString())
	}

	if !data.NmeaDestinationPortVariable.IsNull() {
		body, _ = sjson.Set(body, path+"destinationPort.optionType", "variable")
		body, _ = sjson.Set(body, path+"destinationPort.value", data.NmeaDestinationPortVariable.ValueString())
	} else if data.NmeaDestinationPort.IsNull() {
		body, _ = sjson.Set(body, path+"destinationPort.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"destinationPort.optionType", "global")
		body, _ = sjson.Set(body, path+"destinationPort.value", data.NmeaDestinationPort.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportGPS) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.GpsEnable = types.BoolNull()
	data.GpsEnableVariable = types.StringNull()
	if t := res.Get(path + "enable.optionType"); t.Exists() {
		va := res.Get(path + "enable.value")
		if t.String() == "variable" {
			data.GpsEnableVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GpsEnable = types.BoolValue(va.Bool())
		}
	}
	data.GpsMode = types.StringNull()
	data.GpsModeVariable = types.StringNull()
	if t := res.Get(path + "mode.optionType"); t.Exists() {
		va := res.Get(path + "mode.value")
		if t.String() == "variable" {
			data.GpsModeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GpsMode = types.StringValue(va.String())
		}
	}
	data.NmeaEnable = types.BoolNull()
	data.NmeaEnableVariable = types.StringNull()
	if t := res.Get(path + "nmea.optionType"); t.Exists() {
		va := res.Get(path + "nmea.value")
		if t.String() == "variable" {
			data.NmeaEnableVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NmeaEnable = types.BoolValue(va.Bool())
		}
	}
	data.NmeaSourceAddress = types.StringNull()
	data.NmeaSourceAddressVariable = types.StringNull()
	if t := res.Get(path + "sourceAddress.optionType"); t.Exists() {
		va := res.Get(path + "sourceAddress.value")
		if t.String() == "variable" {
			data.NmeaSourceAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NmeaSourceAddress = types.StringValue(va.String())
		}
	}
	data.NmeaDestinationAddress = types.StringNull()
	data.NmeaDestinationAddressVariable = types.StringNull()
	if t := res.Get(path + "destinationAddress.optionType"); t.Exists() {
		va := res.Get(path + "destinationAddress.value")
		if t.String() == "variable" {
			data.NmeaDestinationAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NmeaDestinationAddress = types.StringValue(va.String())
		}
	}
	data.NmeaDestinationPort = types.Int64Null()
	data.NmeaDestinationPortVariable = types.StringNull()
	if t := res.Get(path + "destinationPort.optionType"); t.Exists() {
		va := res.Get(path + "destinationPort.value")
		if t.String() == "variable" {
			data.NmeaDestinationPortVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NmeaDestinationPort = types.Int64Value(va.Int())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TransportGPS) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.GpsEnable = types.BoolNull()
	data.GpsEnableVariable = types.StringNull()
	if t := res.Get(path + "enable.optionType"); t.Exists() {
		va := res.Get(path + "enable.value")
		if t.String() == "variable" {
			data.GpsEnableVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GpsEnable = types.BoolValue(va.Bool())
		}
	}
	data.GpsMode = types.StringNull()
	data.GpsModeVariable = types.StringNull()
	if t := res.Get(path + "mode.optionType"); t.Exists() {
		va := res.Get(path + "mode.value")
		if t.String() == "variable" {
			data.GpsModeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GpsMode = types.StringValue(va.String())
		}
	}
	data.NmeaEnable = types.BoolNull()
	data.NmeaEnableVariable = types.StringNull()
	if t := res.Get(path + "nmea.optionType"); t.Exists() {
		va := res.Get(path + "nmea.value")
		if t.String() == "variable" {
			data.NmeaEnableVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NmeaEnable = types.BoolValue(va.Bool())
		}
	}
	data.NmeaSourceAddress = types.StringNull()
	data.NmeaSourceAddressVariable = types.StringNull()
	if t := res.Get(path + "sourceAddress.optionType"); t.Exists() {
		va := res.Get(path + "sourceAddress.value")
		if t.String() == "variable" {
			data.NmeaSourceAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NmeaSourceAddress = types.StringValue(va.String())
		}
	}
	data.NmeaDestinationAddress = types.StringNull()
	data.NmeaDestinationAddressVariable = types.StringNull()
	if t := res.Get(path + "destinationAddress.optionType"); t.Exists() {
		va := res.Get(path + "destinationAddress.value")
		if t.String() == "variable" {
			data.NmeaDestinationAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NmeaDestinationAddress = types.StringValue(va.String())
		}
	}
	data.NmeaDestinationPort = types.Int64Null()
	data.NmeaDestinationPortVariable = types.StringNull()
	if t := res.Get(path + "destinationPort.optionType"); t.Exists() {
		va := res.Get(path + "destinationPort.value")
		if t.String() == "variable" {
			data.NmeaDestinationPortVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NmeaDestinationPort = types.Int64Value(va.Int())
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *TransportGPS) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.GpsEnable.IsNull() {
		return false
	}
	if !data.GpsEnableVariable.IsNull() {
		return false
	}
	if !data.GpsMode.IsNull() {
		return false
	}
	if !data.GpsModeVariable.IsNull() {
		return false
	}
	if !data.NmeaEnable.IsNull() {
		return false
	}
	if !data.NmeaEnableVariable.IsNull() {
		return false
	}
	if !data.NmeaSourceAddress.IsNull() {
		return false
	}
	if !data.NmeaSourceAddressVariable.IsNull() {
		return false
	}
	if !data.NmeaDestinationAddress.IsNull() {
		return false
	}
	if !data.NmeaDestinationAddressVariable.IsNull() {
		return false
	}
	if !data.NmeaDestinationPort.IsNull() {
		return false
	}
	if !data.NmeaDestinationPortVariable.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
