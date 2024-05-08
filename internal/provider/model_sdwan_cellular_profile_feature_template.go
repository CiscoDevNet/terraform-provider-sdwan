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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type CellularProfile struct {
	Id                            types.String `tfsdk:"id"`
	Version                       types.Int64  `tfsdk:"version"`
	TemplateType                  types.String `tfsdk:"template_type"`
	Name                          types.String `tfsdk:"name"`
	Description                   types.String `tfsdk:"description"`
	DeviceTypes                   types.Set    `tfsdk:"device_types"`
	IfName                        types.String `tfsdk:"if_name"`
	IfNameVariable                types.String `tfsdk:"if_name_variable"`
	ProfileId                     types.Int64  `tfsdk:"profile_id"`
	ProfileIdVariable             types.String `tfsdk:"profile_id_variable"`
	AccessPointName               types.String `tfsdk:"access_point_name"`
	AccessPointNameVariable       types.String `tfsdk:"access_point_name_variable"`
	AuthenticationType            types.String `tfsdk:"authentication_type"`
	AuthenticationTypeVariable    types.String `tfsdk:"authentication_type_variable"`
	IpAddress                     types.String `tfsdk:"ip_address"`
	IpAddressVariable             types.String `tfsdk:"ip_address_variable"`
	ProfileName                   types.String `tfsdk:"profile_name"`
	ProfileNameVariable           types.String `tfsdk:"profile_name_variable"`
	PacketDataNetworkType         types.String `tfsdk:"packet_data_network_type"`
	PacketDataNetworkTypeVariable types.String `tfsdk:"packet_data_network_type_variable"`
	ProfileUsername               types.String `tfsdk:"profile_username"`
	ProfileUsernameVariable       types.String `tfsdk:"profile_username_variable"`
	ProfilePassword               types.String `tfsdk:"profile_password"`
	ProfilePasswordVariable       types.String `tfsdk:"profile_password_variable"`
	PrimaryDnsAddress             types.String `tfsdk:"primary_dns_address"`
	PrimaryDnsAddressVariable     types.String `tfsdk:"primary_dns_address_variable"`
	SecondaryDnsAddress           types.String `tfsdk:"secondary_dns_address"`
	SecondaryDnsAddressVariable   types.String `tfsdk:"secondary_dns_address_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CellularProfile) getModel() string {
	return "cellular-profile"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CellularProfile) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cellular-profile")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.IfNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"if-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"if-name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"if-name."+"vipVariableName", data.IfNameVariable.ValueString())
	} else if data.IfName.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"if-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"if-name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"if-name."+"vipValue", data.IfName.ValueString())
	}

	if !data.ProfileIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"profile.profile-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.profile-id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"profile.profile-id."+"vipVariableName", data.ProfileIdVariable.ValueString())
	} else if data.ProfileId.IsNull() {
		if !gjson.Get(body, path+"profile").Exists() {
			body, _ = sjson.Set(body, path+"profile", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"profile.profile-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.profile-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"profile.profile-id."+"vipValue", data.ProfileId.ValueInt64())
	}

	if !data.AccessPointNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"profile.apn."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.apn."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"profile.apn."+"vipVariableName", data.AccessPointNameVariable.ValueString())
	} else if data.AccessPointName.IsNull() {
		body, _ = sjson.Set(body, path+"profile.apn."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.apn."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"profile.apn."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.apn."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"profile.apn."+"vipValue", data.AccessPointName.ValueString())
	}

	if !data.AuthenticationTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"profile.auth."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.auth."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"profile.auth."+"vipVariableName", data.AuthenticationTypeVariable.ValueString())
	} else if data.AuthenticationType.IsNull() {
		body, _ = sjson.Set(body, path+"profile.auth."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.auth."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"profile.auth."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.auth."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"profile.auth."+"vipValue", data.AuthenticationType.ValueString())
	}

	if !data.IpAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"profile.ip-addr."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.ip-addr."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"profile.ip-addr."+"vipVariableName", data.IpAddressVariable.ValueString())
	} else if data.IpAddress.IsNull() {
		body, _ = sjson.Set(body, path+"profile.ip-addr."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.ip-addr."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"profile.ip-addr."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.ip-addr."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"profile.ip-addr."+"vipValue", data.IpAddress.ValueString())
	}

	if !data.ProfileNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"profile.name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"profile.name."+"vipVariableName", data.ProfileNameVariable.ValueString())
	} else if data.ProfileName.IsNull() {
		body, _ = sjson.Set(body, path+"profile.name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.name."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"profile.name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"profile.name."+"vipValue", data.ProfileName.ValueString())
	}

	if !data.PacketDataNetworkTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"profile.pdn-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.pdn-type."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"profile.pdn-type."+"vipVariableName", data.PacketDataNetworkTypeVariable.ValueString())
	} else if data.PacketDataNetworkType.IsNull() {
		body, _ = sjson.Set(body, path+"profile.pdn-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.pdn-type."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"profile.pdn-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.pdn-type."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"profile.pdn-type."+"vipValue", data.PacketDataNetworkType.ValueString())
	}

	if !data.ProfileUsernameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"profile.user-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.user-name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"profile.user-name."+"vipVariableName", data.ProfileUsernameVariable.ValueString())
	} else if data.ProfileUsername.IsNull() {
		body, _ = sjson.Set(body, path+"profile.user-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.user-name."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"profile.user-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.user-name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"profile.user-name."+"vipValue", data.ProfileUsername.ValueString())
	}

	if !data.ProfilePasswordVariable.IsNull() {
		body, _ = sjson.Set(body, path+"profile.user-pass."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.user-pass."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"profile.user-pass."+"vipVariableName", data.ProfilePasswordVariable.ValueString())
	} else if data.ProfilePassword.IsNull() {
		body, _ = sjson.Set(body, path+"profile.user-pass."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.user-pass."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"profile.user-pass."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.user-pass."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"profile.user-pass."+"vipValue", data.ProfilePassword.ValueString())
	}

	if !data.PrimaryDnsAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"profile.primary-dns."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.primary-dns."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"profile.primary-dns."+"vipVariableName", data.PrimaryDnsAddressVariable.ValueString())
	} else if data.PrimaryDnsAddress.IsNull() {
		body, _ = sjson.Set(body, path+"profile.primary-dns."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.primary-dns."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"profile.primary-dns."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.primary-dns."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"profile.primary-dns."+"vipValue", data.PrimaryDnsAddress.ValueString())
	}

	if !data.SecondaryDnsAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"profile.secondary-dns."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.secondary-dns."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"profile.secondary-dns."+"vipVariableName", data.SecondaryDnsAddressVariable.ValueString())
	} else if data.SecondaryDnsAddress.IsNull() {
		body, _ = sjson.Set(body, path+"profile.secondary-dns."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.secondary-dns."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"profile.secondary-dns."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"profile.secondary-dns."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"profile.secondary-dns."+"vipValue", data.SecondaryDnsAddress.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CellularProfile) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "if-name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IfName = types.StringNull()

			v := res.Get(path + "if-name.vipVariableName")
			data.IfNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IfName = types.StringNull()
			data.IfNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "if-name.vipValue")
			data.IfName = types.StringValue(v.String())
			data.IfNameVariable = types.StringNull()
		}
	} else {
		data.IfName = types.StringNull()
		data.IfNameVariable = types.StringNull()
	}
	if value := res.Get(path + "profile.profile-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ProfileId = types.Int64Null()

			v := res.Get(path + "profile.profile-id.vipVariableName")
			data.ProfileIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ProfileId = types.Int64Null()
			data.ProfileIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "profile.profile-id.vipValue")
			data.ProfileId = types.Int64Value(v.Int())
			data.ProfileIdVariable = types.StringNull()
		}
	} else {
		data.ProfileId = types.Int64Null()
		data.ProfileIdVariable = types.StringNull()
	}
	if value := res.Get(path + "profile.apn.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AccessPointName = types.StringNull()

			v := res.Get(path + "profile.apn.vipVariableName")
			data.AccessPointNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AccessPointName = types.StringNull()
			data.AccessPointNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "profile.apn.vipValue")
			data.AccessPointName = types.StringValue(v.String())
			data.AccessPointNameVariable = types.StringNull()
		}
	} else {
		data.AccessPointName = types.StringNull()
		data.AccessPointNameVariable = types.StringNull()
	}
	if value := res.Get(path + "profile.auth.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AuthenticationType = types.StringNull()

			v := res.Get(path + "profile.auth.vipVariableName")
			data.AuthenticationTypeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AuthenticationType = types.StringNull()
			data.AuthenticationTypeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "profile.auth.vipValue")
			data.AuthenticationType = types.StringValue(v.String())
			data.AuthenticationTypeVariable = types.StringNull()
		}
	} else {
		data.AuthenticationType = types.StringNull()
		data.AuthenticationTypeVariable = types.StringNull()
	}
	if value := res.Get(path + "profile.ip-addr.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IpAddress = types.StringNull()

			v := res.Get(path + "profile.ip-addr.vipVariableName")
			data.IpAddressVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IpAddress = types.StringNull()
			data.IpAddressVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "profile.ip-addr.vipValue")
			data.IpAddress = types.StringValue(v.String())
			data.IpAddressVariable = types.StringNull()
		}
	} else {
		data.IpAddress = types.StringNull()
		data.IpAddressVariable = types.StringNull()
	}
	if value := res.Get(path + "profile.name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ProfileName = types.StringNull()

			v := res.Get(path + "profile.name.vipVariableName")
			data.ProfileNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ProfileName = types.StringNull()
			data.ProfileNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "profile.name.vipValue")
			data.ProfileName = types.StringValue(v.String())
			data.ProfileNameVariable = types.StringNull()
		}
	} else {
		data.ProfileName = types.StringNull()
		data.ProfileNameVariable = types.StringNull()
	}
	if value := res.Get(path + "profile.pdn-type.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PacketDataNetworkType = types.StringNull()

			v := res.Get(path + "profile.pdn-type.vipVariableName")
			data.PacketDataNetworkTypeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.PacketDataNetworkType = types.StringNull()
			data.PacketDataNetworkTypeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "profile.pdn-type.vipValue")
			data.PacketDataNetworkType = types.StringValue(v.String())
			data.PacketDataNetworkTypeVariable = types.StringNull()
		}
	} else {
		data.PacketDataNetworkType = types.StringNull()
		data.PacketDataNetworkTypeVariable = types.StringNull()
	}
	if value := res.Get(path + "profile.user-name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ProfileUsername = types.StringNull()

			v := res.Get(path + "profile.user-name.vipVariableName")
			data.ProfileUsernameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ProfileUsername = types.StringNull()
			data.ProfileUsernameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "profile.user-name.vipValue")
			data.ProfileUsername = types.StringValue(v.String())
			data.ProfileUsernameVariable = types.StringNull()
		}
	} else {
		data.ProfileUsername = types.StringNull()
		data.ProfileUsernameVariable = types.StringNull()
	}
	if value := res.Get(path + "profile.user-pass.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ProfilePassword = types.StringNull()

			v := res.Get(path + "profile.user-pass.vipVariableName")
			data.ProfilePasswordVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ProfilePassword = types.StringNull()
			data.ProfilePasswordVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "profile.user-pass.vipValue")
			data.ProfilePassword = types.StringValue(v.String())
			data.ProfilePasswordVariable = types.StringNull()
		}
	} else {
		data.ProfilePassword = types.StringNull()
		data.ProfilePasswordVariable = types.StringNull()
	}
	if value := res.Get(path + "profile.primary-dns.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PrimaryDnsAddress = types.StringNull()

			v := res.Get(path + "profile.primary-dns.vipVariableName")
			data.PrimaryDnsAddressVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.PrimaryDnsAddress = types.StringNull()
			data.PrimaryDnsAddressVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "profile.primary-dns.vipValue")
			data.PrimaryDnsAddress = types.StringValue(v.String())
			data.PrimaryDnsAddressVariable = types.StringNull()
		}
	} else {
		data.PrimaryDnsAddress = types.StringNull()
		data.PrimaryDnsAddressVariable = types.StringNull()
	}
	if value := res.Get(path + "profile.secondary-dns.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SecondaryDnsAddress = types.StringNull()

			v := res.Get(path + "profile.secondary-dns.vipVariableName")
			data.SecondaryDnsAddressVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SecondaryDnsAddress = types.StringNull()
			data.SecondaryDnsAddressVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "profile.secondary-dns.vipValue")
			data.SecondaryDnsAddress = types.StringValue(v.String())
			data.SecondaryDnsAddressVariable = types.StringNull()
		}
	} else {
		data.SecondaryDnsAddress = types.StringNull()
		data.SecondaryDnsAddressVariable = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CellularProfile) hasChanges(ctx context.Context, state *CellularProfile) bool {
	hasChanges := false
	if !data.IfName.Equal(state.IfName) {
		hasChanges = true
	}
	if !data.ProfileId.Equal(state.ProfileId) {
		hasChanges = true
	}
	if !data.AccessPointName.Equal(state.AccessPointName) {
		hasChanges = true
	}
	if !data.AuthenticationType.Equal(state.AuthenticationType) {
		hasChanges = true
	}
	if !data.IpAddress.Equal(state.IpAddress) {
		hasChanges = true
	}
	if !data.ProfileName.Equal(state.ProfileName) {
		hasChanges = true
	}
	if !data.PacketDataNetworkType.Equal(state.PacketDataNetworkType) {
		hasChanges = true
	}
	if !data.ProfileUsername.Equal(state.ProfileUsername) {
		hasChanges = true
	}
	if !data.ProfilePassword.Equal(state.ProfilePassword) {
		hasChanges = true
	}
	if !data.PrimaryDnsAddress.Equal(state.PrimaryDnsAddress) {
		hasChanges = true
	}
	if !data.SecondaryDnsAddress.Equal(state.SecondaryDnsAddress) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
