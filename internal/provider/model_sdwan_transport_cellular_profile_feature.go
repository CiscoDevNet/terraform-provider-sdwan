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
type TransportCellularProfile struct {
	Id                            types.String `tfsdk:"id"`
	Version                       types.Int64  `tfsdk:"version"`
	Name                          types.String `tfsdk:"name"`
	Description                   types.String `tfsdk:"description"`
	FeatureProfileId              types.String `tfsdk:"feature_profile_id"`
	ProfileId                     types.Int64  `tfsdk:"profile_id"`
	ProfileIdVariable             types.String `tfsdk:"profile_id_variable"`
	AccessPointName               types.String `tfsdk:"access_point_name"`
	AccessPointNameVariable       types.String `tfsdk:"access_point_name_variable"`
	NoAuthentication              types.String `tfsdk:"no_authentication"`
	NeedAuthentication            types.String `tfsdk:"need_authentication"`
	PacketDataNetworkType         types.String `tfsdk:"packet_data_network_type"`
	PacketDataNetworkTypeVariable types.String `tfsdk:"packet_data_network_type_variable"`
	NoOverwrite                   types.Bool   `tfsdk:"no_overwrite"`
	NoOverwriteVariable           types.String `tfsdk:"no_overwrite_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TransportCellularProfile) getModel() string {
	return "transport_cellular_profile"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportCellularProfile) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/cellular", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportCellularProfile) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {
		body, _ = sjson.Set(body, path+"configType.optionType", "default")
		body, _ = sjson.Set(body, path+"configType.value", "non-eSim")
	}

	if !data.ProfileIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.id.optionType", "variable")
			body, _ = sjson.Set(body, path+"profileConfig.id.value", data.ProfileIdVariable.ValueString())
		}
	} else if !data.ProfileId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.id.optionType", "global")
			body, _ = sjson.Set(body, path+"profileConfig.id.value", data.ProfileId.ValueInt64())
		}
	}

	if !data.AccessPointNameVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.apn.optionType", "variable")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.apn.value", data.AccessPointNameVariable.ValueString())
		}
	} else if !data.AccessPointName.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.apn.optionType", "global")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.apn.value", data.AccessPointName.ValueString())
		}
	}
	if !data.NoAuthentication.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.noAuthentication.optionType", "global")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.noAuthentication.value", data.NoAuthentication.ValueString())
		}
	}
	if !data.NeedAuthentication.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.needAuthentication.optionType", "global")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.authentication.needAuthentication.value", data.NeedAuthentication.ValueString())
		}
	}

	if !data.PacketDataNetworkTypeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.pdnType.optionType", "variable")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.pdnType.value", data.PacketDataNetworkTypeVariable.ValueString())
		}
	} else if data.PacketDataNetworkType.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.pdnType.optionType", "default")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.pdnType.value", "ipv4")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.pdnType.optionType", "global")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.pdnType.value", data.PacketDataNetworkType.ValueString())
		}
	}

	if !data.NoOverwriteVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.noOverwrite.optionType", "variable")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.noOverwrite.value", data.NoOverwriteVariable.ValueString())
		}
	} else if data.NoOverwrite.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.noOverwrite.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.noOverwrite.optionType", "global")
			body, _ = sjson.Set(body, path+"profileConfig.profileInfo.noOverwrite.value", data.NoOverwrite.ValueBool())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportCellularProfile) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.ProfileId = types.Int64Null()
	data.ProfileIdVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.id.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.id.value")
		if t.String() == "variable" {
			data.ProfileIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ProfileId = types.Int64Value(va.Int())
		}
	}
	data.AccessPointName = types.StringNull()
	data.AccessPointNameVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.profileInfo.apn.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.apn.value")
		if t.String() == "variable" {
			data.AccessPointNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AccessPointName = types.StringValue(va.String())
		}
	}
	data.NoAuthentication = types.StringNull()

	if t := res.Get(path + "profileConfig.profileInfo.authentication.noAuthentication.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.authentication.noAuthentication.value")
		if t.String() == "global" {
			data.NoAuthentication = types.StringValue(va.String())
		}
	}
	data.NeedAuthentication = types.StringNull()

	if t := res.Get(path + "profileConfig.profileInfo.authentication.needAuthentication.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.authentication.needAuthentication.value")
		if t.String() == "global" {
			data.NeedAuthentication = types.StringValue(va.String())
		}
	}
	data.PacketDataNetworkType = types.StringNull()
	data.PacketDataNetworkTypeVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.profileInfo.pdnType.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.pdnType.value")
		if t.String() == "variable" {
			data.PacketDataNetworkTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PacketDataNetworkType = types.StringValue(va.String())
		}
	}
	data.NoOverwrite = types.BoolNull()
	data.NoOverwriteVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.profileInfo.noOverwrite.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.noOverwrite.value")
		if t.String() == "variable" {
			data.NoOverwriteVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NoOverwrite = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TransportCellularProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.ProfileId = types.Int64Null()
	data.ProfileIdVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.id.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.id.value")
		if t.String() == "variable" {
			data.ProfileIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ProfileId = types.Int64Value(va.Int())
		}
	}
	data.AccessPointName = types.StringNull()
	data.AccessPointNameVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.profileInfo.apn.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.apn.value")
		if t.String() == "variable" {
			data.AccessPointNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AccessPointName = types.StringValue(va.String())
		}
	}
	data.NoAuthentication = types.StringNull()

	if t := res.Get(path + "profileConfig.profileInfo.authentication.noAuthentication.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.authentication.noAuthentication.value")
		if t.String() == "global" {
			data.NoAuthentication = types.StringValue(va.String())
		}
	}
	data.NeedAuthentication = types.StringNull()

	if t := res.Get(path + "profileConfig.profileInfo.authentication.needAuthentication.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.authentication.needAuthentication.value")
		if t.String() == "global" {
			data.NeedAuthentication = types.StringValue(va.String())
		}
	}
	data.PacketDataNetworkType = types.StringNull()
	data.PacketDataNetworkTypeVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.profileInfo.pdnType.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.pdnType.value")
		if t.String() == "variable" {
			data.PacketDataNetworkTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PacketDataNetworkType = types.StringValue(va.String())
		}
	}
	data.NoOverwrite = types.BoolNull()
	data.NoOverwriteVariable = types.StringNull()
	if t := res.Get(path + "profileConfig.profileInfo.noOverwrite.optionType"); t.Exists() {
		va := res.Get(path + "profileConfig.profileInfo.noOverwrite.value")
		if t.String() == "variable" {
			data.NoOverwriteVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NoOverwrite = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *TransportCellularProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.ProfileId.IsNull() {
		return false
	}
	if !data.ProfileIdVariable.IsNull() {
		return false
	}
	if !data.AccessPointName.IsNull() {
		return false
	}
	if !data.AccessPointNameVariable.IsNull() {
		return false
	}
	if !data.NoAuthentication.IsNull() {
		return false
	}
	if !data.NeedAuthentication.IsNull() {
		return false
	}
	if !data.PacketDataNetworkType.IsNull() {
		return false
	}
	if !data.PacketDataNetworkTypeVariable.IsNull() {
		return false
	}
	if !data.NoOverwrite.IsNull() {
		return false
	}
	if !data.NoOverwriteVariable.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
