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
	Id                         types.String `tfsdk:"id"`
	Version                    types.Int64  `tfsdk:"version"`
	Name                       types.String `tfsdk:"name"`
	Description                types.String `tfsdk:"description"`
	FeatureProfileId           types.String `tfsdk:"feature_profile_id"`
	Enable                     types.Bool   `tfsdk:"enable"`
	EnableVariable             types.String `tfsdk:"enable_variable"`
	Mode                       types.String `tfsdk:"mode"`
	ModeVariable               types.String `tfsdk:"mode_variable"`
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

	if !data.EnableVariable.IsNull() {
		body, _ = sjson.Set(body, path+"enable.optionType", "variable")
		body, _ = sjson.Set(body, path+"enable.value", data.EnableVariable.ValueString())
	} else if data.Enable.IsNull() {
		body, _ = sjson.Set(body, path+"enable.optionType", "default")
		body, _ = sjson.Set(body, path+"enable.value", false)
	} else {
		body, _ = sjson.Set(body, path+"enable.optionType", "global")
		body, _ = sjson.Set(body, path+"enable.value", data.Enable.ValueBool())
	}

	if !data.ModeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"mode.optionType", "variable")
		body, _ = sjson.Set(body, path+"mode.value", data.ModeVariable.ValueString())
	} else if data.Mode.IsNull() {
		body, _ = sjson.Set(body, path+"mode.optionType", "default")
		body, _ = sjson.Set(body, path+"mode.value", "ms-based")
	} else {
		body, _ = sjson.Set(body, path+"mode.optionType", "global")
		body, _ = sjson.Set(body, path+"mode.value", data.Mode.ValueString())
	}

	if !data.NmeaVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nmea.optionType", "variable")
		body, _ = sjson.Set(body, path+"nmea.value", data.NmeaVariable.ValueString())
	} else if data.Nmea.IsNull() {
		body, _ = sjson.Set(body, path+"nmea.optionType", "default")
		body, _ = sjson.Set(body, path+"nmea.value", false)
	} else {
		body, _ = sjson.Set(body, path+"nmea.optionType", "global")
		body, _ = sjson.Set(body, path+"nmea.value", data.Nmea.ValueBool())
	}

	if !data.SourceAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"sourceAddress.optionType", "variable")
		body, _ = sjson.Set(body, path+"sourceAddress.value", data.SourceAddressVariable.ValueString())
	} else if data.SourceAddress.IsNull() {
		body, _ = sjson.Set(body, path+"sourceAddress.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"sourceAddress.optionType", "global")
		body, _ = sjson.Set(body, path+"sourceAddress.value", data.SourceAddress.ValueString())
	}

	if !data.DestinationAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"destinationAddress.optionType", "variable")
		body, _ = sjson.Set(body, path+"destinationAddress.value", data.DestinationAddressVariable.ValueString())
	} else if data.DestinationAddress.IsNull() {
		body, _ = sjson.Set(body, path+"destinationAddress.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"destinationAddress.optionType", "global")
		body, _ = sjson.Set(body, path+"destinationAddress.value", data.DestinationAddress.ValueString())
	}

	if !data.DestinationPortVariable.IsNull() {
		body, _ = sjson.Set(body, path+"destinationPort.optionType", "variable")
		body, _ = sjson.Set(body, path+"destinationPort.value", data.DestinationPortVariable.ValueString())
	} else if data.DestinationPort.IsNull() {
		body, _ = sjson.Set(body, path+"destinationPort.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"destinationPort.optionType", "global")
		body, _ = sjson.Set(body, path+"destinationPort.value", data.DestinationPort.ValueInt64())
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
	data.Enable = types.BoolNull()
	data.EnableVariable = types.StringNull()
	if t := res.Get(path + "enable.optionType"); t.Exists() {
		va := res.Get(path + "enable.value")
		if t.String() == "variable" {
			data.EnableVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Enable = types.BoolValue(va.Bool())
		}
	}
	data.Mode = types.StringNull()
	data.ModeVariable = types.StringNull()
	if t := res.Get(path + "mode.optionType"); t.Exists() {
		va := res.Get(path + "mode.value")
		if t.String() == "variable" {
			data.ModeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Mode = types.StringValue(va.String())
		}
	}
	data.Nmea = types.BoolNull()
	data.NmeaVariable = types.StringNull()
	if t := res.Get(path + "nmea.optionType"); t.Exists() {
		va := res.Get(path + "nmea.value")
		if t.String() == "variable" {
			data.NmeaVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Nmea = types.BoolValue(va.Bool())
		}
	}
	data.SourceAddress = types.StringNull()
	data.SourceAddressVariable = types.StringNull()
	if t := res.Get(path + "sourceAddress.optionType"); t.Exists() {
		va := res.Get(path + "sourceAddress.value")
		if t.String() == "variable" {
			data.SourceAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SourceAddress = types.StringValue(va.String())
		}
	}
	data.DestinationAddress = types.StringNull()
	data.DestinationAddressVariable = types.StringNull()
	if t := res.Get(path + "destinationAddress.optionType"); t.Exists() {
		va := res.Get(path + "destinationAddress.value")
		if t.String() == "variable" {
			data.DestinationAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DestinationAddress = types.StringValue(va.String())
		}
	}
	data.DestinationPort = types.Int64Null()
	data.DestinationPortVariable = types.StringNull()
	if t := res.Get(path + "destinationPort.optionType"); t.Exists() {
		va := res.Get(path + "destinationPort.value")
		if t.String() == "variable" {
			data.DestinationPortVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DestinationPort = types.Int64Value(va.Int())
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
	data.Enable = types.BoolNull()
	data.EnableVariable = types.StringNull()
	if t := res.Get(path + "enable.optionType"); t.Exists() {
		va := res.Get(path + "enable.value")
		if t.String() == "variable" {
			data.EnableVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Enable = types.BoolValue(va.Bool())
		}
	}
	data.Mode = types.StringNull()
	data.ModeVariable = types.StringNull()
	if t := res.Get(path + "mode.optionType"); t.Exists() {
		va := res.Get(path + "mode.value")
		if t.String() == "variable" {
			data.ModeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Mode = types.StringValue(va.String())
		}
	}
	data.Nmea = types.BoolNull()
	data.NmeaVariable = types.StringNull()
	if t := res.Get(path + "nmea.optionType"); t.Exists() {
		va := res.Get(path + "nmea.value")
		if t.String() == "variable" {
			data.NmeaVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Nmea = types.BoolValue(va.Bool())
		}
	}
	data.SourceAddress = types.StringNull()
	data.SourceAddressVariable = types.StringNull()
	if t := res.Get(path + "sourceAddress.optionType"); t.Exists() {
		va := res.Get(path + "sourceAddress.value")
		if t.String() == "variable" {
			data.SourceAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SourceAddress = types.StringValue(va.String())
		}
	}
	data.DestinationAddress = types.StringNull()
	data.DestinationAddressVariable = types.StringNull()
	if t := res.Get(path + "destinationAddress.optionType"); t.Exists() {
		va := res.Get(path + "destinationAddress.value")
		if t.String() == "variable" {
			data.DestinationAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DestinationAddress = types.StringValue(va.String())
		}
	}
	data.DestinationPort = types.Int64Null()
	data.DestinationPortVariable = types.StringNull()
	if t := res.Get(path + "destinationPort.optionType"); t.Exists() {
		va := res.Get(path + "destinationPort.value")
		if t.String() == "variable" {
			data.DestinationPortVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DestinationPort = types.Int64Value(va.Int())
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *TransportGPS) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.Enable.IsNull() {
		return false
	}
	if !data.EnableVariable.IsNull() {
		return false
	}
	if !data.Mode.IsNull() {
		return false
	}
	if !data.ModeVariable.IsNull() {
		return false
	}
	if !data.Nmea.IsNull() {
		return false
	}
	if !data.NmeaVariable.IsNull() {
		return false
	}
	if !data.SourceAddress.IsNull() {
		return false
	}
	if !data.SourceAddressVariable.IsNull() {
		return false
	}
	if !data.DestinationAddress.IsNull() {
		return false
	}
	if !data.DestinationAddressVariable.IsNull() {
		return false
	}
	if !data.DestinationPort.IsNull() {
		return false
	}
	if !data.DestinationPortVariable.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
