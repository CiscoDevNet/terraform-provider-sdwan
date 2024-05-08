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
type CiscoWirelessLAN struct {
	Id                               types.String            `tfsdk:"id"`
	Version                          types.Int64             `tfsdk:"version"`
	TemplateType                     types.String            `tfsdk:"template_type"`
	Name                             types.String            `tfsdk:"name"`
	Description                      types.String            `tfsdk:"description"`
	DeviceTypes                      types.Set               `tfsdk:"device_types"`
	Shutdown24ghz                    types.Bool              `tfsdk:"shutdown_2_4ghz"`
	Shutdown24ghzVariable            types.String            `tfsdk:"shutdown_2_4ghz_variable"`
	Shutdown5ghz                     types.Bool              `tfsdk:"shutdown_5ghz"`
	Shutdown5ghzVariable             types.String            `tfsdk:"shutdown_5ghz_variable"`
	Ssids                            []CiscoWirelessLANSsids `tfsdk:"ssids"`
	Country                          types.String            `tfsdk:"country"`
	CountryVariable                  types.String            `tfsdk:"country_variable"`
	Username                         types.String            `tfsdk:"username"`
	UsernameVariable                 types.String            `tfsdk:"username_variable"`
	Password                         types.String            `tfsdk:"password"`
	PasswordVariable                 types.String            `tfsdk:"password_variable"`
	ControllerIpAddress              types.String            `tfsdk:"controller_ip_address"`
	ControllerIpAddressVariable      types.String            `tfsdk:"controller_ip_address_variable"`
	ControllerSubnetMask             types.String            `tfsdk:"controller_subnet_mask"`
	ControllerSubnetMaskVariable     types.String            `tfsdk:"controller_subnet_mask_variable"`
	ControllerDefaultGateway         types.String            `tfsdk:"controller_default_gateway"`
	ControllerDefaultGatewayVariable types.String            `tfsdk:"controller_default_gateway_variable"`
}

type CiscoWirelessLANSsids struct {
	Optional                   types.Bool   `tfsdk:"optional"`
	WirelessNetworkName        types.String `tfsdk:"wireless_network_name"`
	AdminState                 types.Bool   `tfsdk:"admin_state"`
	AdminStateVariable         types.String `tfsdk:"admin_state_variable"`
	BroadcastSsid              types.Bool   `tfsdk:"broadcast_ssid"`
	VlanId                     types.Int64  `tfsdk:"vlan_id"`
	VlanIdVariable             types.String `tfsdk:"vlan_id_variable"`
	RadioType                  types.String `tfsdk:"radio_type"`
	RadioTypeVariable          types.String `tfsdk:"radio_type_variable"`
	SecurityType               types.String `tfsdk:"security_type"`
	SecurityTypeVariable       types.String `tfsdk:"security_type_variable"`
	RadiusServerIp             types.String `tfsdk:"radius_server_ip"`
	RadiusServerIpVariable     types.String `tfsdk:"radius_server_ip_variable"`
	RadiusServerPort           types.Int64  `tfsdk:"radius_server_port"`
	RadiusServerPortVariable   types.String `tfsdk:"radius_server_port_variable"`
	RadiusServerSecret         types.String `tfsdk:"radius_server_secret"`
	RadiusServerSecretVariable types.String `tfsdk:"radius_server_secret_variable"`
	Passphrase                 types.String `tfsdk:"passphrase"`
	PassphraseVariable         types.String `tfsdk:"passphrase_variable"`
	QosProfile                 types.String `tfsdk:"qos_profile"`
	QosProfileVariable         types.String `tfsdk:"qos_profile_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoWirelessLAN) getModel() string {
	return "cisco_wireless_lan"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CiscoWirelessLAN) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_wireless_lan")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.Shutdown24ghzVariable.IsNull() {
		body, _ = sjson.Set(body, path+"radio.shutdown-2.4ghz."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radio.shutdown-2.4ghz."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"radio.shutdown-2.4ghz."+"vipVariableName", data.Shutdown24ghzVariable.ValueString())
	} else if data.Shutdown24ghz.IsNull() {
		body, _ = sjson.Set(body, path+"radio.shutdown-2.4ghz."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radio.shutdown-2.4ghz."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"radio.shutdown-2.4ghz."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radio.shutdown-2.4ghz."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"radio.shutdown-2.4ghz."+"vipValue", strconv.FormatBool(data.Shutdown24ghz.ValueBool()))
	}

	if !data.Shutdown5ghzVariable.IsNull() {
		body, _ = sjson.Set(body, path+"radio.shutdown-5ghz."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radio.shutdown-5ghz."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"radio.shutdown-5ghz."+"vipVariableName", data.Shutdown5ghzVariable.ValueString())
	} else if data.Shutdown5ghz.IsNull() {
		body, _ = sjson.Set(body, path+"radio.shutdown-5ghz."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radio.shutdown-5ghz."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"radio.shutdown-5ghz."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radio.shutdown-5ghz."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"radio.shutdown-5ghz."+"vipValue", strconv.FormatBool(data.Shutdown5ghz.ValueBool()))
	}
	if len(data.Ssids) > 0 {
		body, _ = sjson.Set(body, path+"ssid."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ssid."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ssid."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"ssid."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"ssid."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"ssid."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"ssid."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"ssid."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ssids {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")
		if item.WirelessNetworkName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.WirelessNetworkName.ValueString())
		}
		itemAttributes = append(itemAttributes, "admin-state")

		if !item.AdminStateVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "admin-state."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "admin-state."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "admin-state."+"vipVariableName", item.AdminStateVariable.ValueString())
		} else if item.AdminState.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "admin-state."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "admin-state."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "admin-state."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "admin-state."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "admin-state."+"vipValue", strconv.FormatBool(item.AdminState.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "broadcast-ssid")
		if item.BroadcastSsid.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "broadcast-ssid."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "broadcast-ssid."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "broadcast-ssid."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "broadcast-ssid."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "broadcast-ssid."+"vipValue", strconv.FormatBool(item.BroadcastSsid.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "vlan-id")

		if !item.VlanIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vlan-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vlan-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "vlan-id."+"vipVariableName", item.VlanIdVariable.ValueString())
		} else if item.VlanId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "vlan-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vlan-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vlan-id."+"vipValue", item.VlanId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "radio-type")

		if !item.RadioTypeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "radio-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "radio-type."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "radio-type."+"vipVariableName", item.RadioTypeVariable.ValueString())
		} else if item.RadioType.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "radio-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "radio-type."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "radio-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "radio-type."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "radio-type."+"vipValue", item.RadioType.ValueString())
		}
		itemAttributes = append(itemAttributes, "security-type")

		if !item.SecurityTypeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "security-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "security-type."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "security-type."+"vipVariableName", item.SecurityTypeVariable.ValueString())
		} else if item.SecurityType.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "security-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "security-type."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "security-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "security-type."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "security-type."+"vipValue", item.SecurityType.ValueString())
		}
		itemAttributes = append(itemAttributes, "radius-server-ip")

		if !item.RadiusServerIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "radius-server-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "radius-server-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "radius-server-ip."+"vipVariableName", item.RadiusServerIpVariable.ValueString())
		} else if item.RadiusServerIp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "radius-server-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "radius-server-ip."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "radius-server-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "radius-server-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "radius-server-ip."+"vipValue", item.RadiusServerIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "radius-server-port")

		if !item.RadiusServerPortVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "radius-server-port."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "radius-server-port."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "radius-server-port."+"vipVariableName", item.RadiusServerPortVariable.ValueString())
		} else if item.RadiusServerPort.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "radius-server-port."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "radius-server-port."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "radius-server-port."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "radius-server-port."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "radius-server-port."+"vipValue", item.RadiusServerPort.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "radius-server-secret")

		if !item.RadiusServerSecretVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "radius-server-secret."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "radius-server-secret."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "radius-server-secret."+"vipVariableName", item.RadiusServerSecretVariable.ValueString())
		} else if item.RadiusServerSecret.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "radius-server-secret."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "radius-server-secret."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "radius-server-secret."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "radius-server-secret."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "radius-server-secret."+"vipValue", item.RadiusServerSecret.ValueString())
		}
		itemAttributes = append(itemAttributes, "passphrase")

		if !item.PassphraseVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "passphrase."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "passphrase."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "passphrase."+"vipVariableName", item.PassphraseVariable.ValueString())
		} else if item.Passphrase.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "passphrase."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "passphrase."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "passphrase."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "passphrase."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "passphrase."+"vipValue", item.Passphrase.ValueString())
		}
		itemAttributes = append(itemAttributes, "qos-profile")

		if !item.QosProfileVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "qos-profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "qos-profile."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "qos-profile."+"vipVariableName", item.QosProfileVariable.ValueString())
		} else if item.QosProfile.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "qos-profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "qos-profile."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "qos-profile."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "qos-profile."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "qos-profile."+"vipValue", item.QosProfile.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"ssid."+"vipValue.-1", itemBody)
	}

	if !data.CountryVariable.IsNull() {
		body, _ = sjson.Set(body, path+"country."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"country."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"country."+"vipVariableName", data.CountryVariable.ValueString())
	} else if data.Country.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"country."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"country."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"country."+"vipValue", data.Country.ValueString())
	}

	if !data.UsernameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"mgmt.username."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mgmt.username."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"mgmt.username."+"vipVariableName", data.UsernameVariable.ValueString())
	} else if data.Username.IsNull() {
		if !gjson.Get(body, path+"mgmt").Exists() {
			body, _ = sjson.Set(body, path+"mgmt", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"mgmt.username."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mgmt.username."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"mgmt.username."+"vipValue", data.Username.ValueString())
	}

	if !data.PasswordVariable.IsNull() {
		body, _ = sjson.Set(body, path+"mgmt.password."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mgmt.password."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"mgmt.password."+"vipVariableName", data.PasswordVariable.ValueString())
	} else if data.Password.IsNull() {
		if !gjson.Get(body, path+"mgmt").Exists() {
			body, _ = sjson.Set(body, path+"mgmt", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"mgmt.password."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mgmt.password."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"mgmt.password."+"vipValue", data.Password.ValueString())
	}

	if !data.ControllerIpAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"mgmt.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mgmt.address."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"mgmt.address."+"vipVariableName", data.ControllerIpAddressVariable.ValueString())
	} else if data.ControllerIpAddress.IsNull() {
		body, _ = sjson.Set(body, path+"mgmt.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mgmt.address."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"mgmt.address."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mgmt.address."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"mgmt.address."+"vipValue", data.ControllerIpAddress.ValueString())
	}

	if !data.ControllerSubnetMaskVariable.IsNull() {
		body, _ = sjson.Set(body, path+"mgmt.netmask."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mgmt.netmask."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"mgmt.netmask."+"vipVariableName", data.ControllerSubnetMaskVariable.ValueString())
	} else if data.ControllerSubnetMask.IsNull() {
		body, _ = sjson.Set(body, path+"mgmt.netmask."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mgmt.netmask."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"mgmt.netmask."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mgmt.netmask."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"mgmt.netmask."+"vipValue", data.ControllerSubnetMask.ValueString())
	}

	if !data.ControllerDefaultGatewayVariable.IsNull() {
		body, _ = sjson.Set(body, path+"mgmt.default-gateway."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mgmt.default-gateway."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"mgmt.default-gateway."+"vipVariableName", data.ControllerDefaultGatewayVariable.ValueString())
	} else if data.ControllerDefaultGateway.IsNull() {
		body, _ = sjson.Set(body, path+"mgmt.default-gateway."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mgmt.default-gateway."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"mgmt.default-gateway."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"mgmt.default-gateway."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"mgmt.default-gateway."+"vipValue", data.ControllerDefaultGateway.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoWirelessLAN) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "radio.shutdown-2.4ghz.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Shutdown24ghz = types.BoolNull()

			v := res.Get(path + "radio.shutdown-2.4ghz.vipVariableName")
			data.Shutdown24ghzVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Shutdown24ghz = types.BoolNull()
			data.Shutdown24ghzVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "radio.shutdown-2.4ghz.vipValue")
			data.Shutdown24ghz = types.BoolValue(v.Bool())
			data.Shutdown24ghzVariable = types.StringNull()
		}
	} else {
		data.Shutdown24ghz = types.BoolNull()
		data.Shutdown24ghzVariable = types.StringNull()
	}
	if value := res.Get(path + "radio.shutdown-5ghz.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Shutdown5ghz = types.BoolNull()

			v := res.Get(path + "radio.shutdown-5ghz.vipVariableName")
			data.Shutdown5ghzVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Shutdown5ghz = types.BoolNull()
			data.Shutdown5ghzVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "radio.shutdown-5ghz.vipValue")
			data.Shutdown5ghz = types.BoolValue(v.Bool())
			data.Shutdown5ghzVariable = types.StringNull()
		}
	} else {
		data.Shutdown5ghz = types.BoolNull()
		data.Shutdown5ghzVariable = types.StringNull()
	}
	if value := res.Get(path + "ssid.vipValue"); len(value.Array()) > 0 {
		data.Ssids = make([]CiscoWirelessLANSsids, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoWirelessLANSsids{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.WirelessNetworkName = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.WirelessNetworkName = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.WirelessNetworkName = types.StringValue(cv.String())

				}
			} else {
				item.WirelessNetworkName = types.StringNull()

			}
			if cValue := v.Get("admin-state.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AdminState = types.BoolNull()

					cv := v.Get("admin-state.vipVariableName")
					item.AdminStateVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AdminState = types.BoolNull()
					item.AdminStateVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("admin-state.vipValue")
					item.AdminState = types.BoolValue(cv.Bool())
					item.AdminStateVariable = types.StringNull()
				}
			} else {
				item.AdminState = types.BoolNull()
				item.AdminStateVariable = types.StringNull()
			}
			if cValue := v.Get("broadcast-ssid.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.BroadcastSsid = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.BroadcastSsid = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("broadcast-ssid.vipValue")
					item.BroadcastSsid = types.BoolValue(cv.Bool())

				}
			} else {
				item.BroadcastSsid = types.BoolNull()

			}
			if cValue := v.Get("vlan-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VlanId = types.Int64Null()

					cv := v.Get("vlan-id.vipVariableName")
					item.VlanIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.VlanId = types.Int64Null()
					item.VlanIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("vlan-id.vipValue")
					item.VlanId = types.Int64Value(cv.Int())
					item.VlanIdVariable = types.StringNull()
				}
			} else {
				item.VlanId = types.Int64Null()
				item.VlanIdVariable = types.StringNull()
			}
			if cValue := v.Get("radio-type.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.RadioType = types.StringNull()

					cv := v.Get("radio-type.vipVariableName")
					item.RadioTypeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.RadioType = types.StringNull()
					item.RadioTypeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("radio-type.vipValue")
					item.RadioType = types.StringValue(cv.String())
					item.RadioTypeVariable = types.StringNull()
				}
			} else {
				item.RadioType = types.StringNull()
				item.RadioTypeVariable = types.StringNull()
			}
			if cValue := v.Get("security-type.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SecurityType = types.StringNull()

					cv := v.Get("security-type.vipVariableName")
					item.SecurityTypeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SecurityType = types.StringNull()
					item.SecurityTypeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("security-type.vipValue")
					item.SecurityType = types.StringValue(cv.String())
					item.SecurityTypeVariable = types.StringNull()
				}
			} else {
				item.SecurityType = types.StringNull()
				item.SecurityTypeVariable = types.StringNull()
			}
			if cValue := v.Get("radius-server-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.RadiusServerIp = types.StringNull()

					cv := v.Get("radius-server-ip.vipVariableName")
					item.RadiusServerIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.RadiusServerIp = types.StringNull()
					item.RadiusServerIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("radius-server-ip.vipValue")
					item.RadiusServerIp = types.StringValue(cv.String())
					item.RadiusServerIpVariable = types.StringNull()
				}
			} else {
				item.RadiusServerIp = types.StringNull()
				item.RadiusServerIpVariable = types.StringNull()
			}
			if cValue := v.Get("radius-server-port.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.RadiusServerPort = types.Int64Null()

					cv := v.Get("radius-server-port.vipVariableName")
					item.RadiusServerPortVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.RadiusServerPort = types.Int64Null()
					item.RadiusServerPortVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("radius-server-port.vipValue")
					item.RadiusServerPort = types.Int64Value(cv.Int())
					item.RadiusServerPortVariable = types.StringNull()
				}
			} else {
				item.RadiusServerPort = types.Int64Null()
				item.RadiusServerPortVariable = types.StringNull()
			}
			if cValue := v.Get("radius-server-secret.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.RadiusServerSecret = types.StringNull()

					cv := v.Get("radius-server-secret.vipVariableName")
					item.RadiusServerSecretVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.RadiusServerSecret = types.StringNull()
					item.RadiusServerSecretVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("radius-server-secret.vipValue")
					item.RadiusServerSecret = types.StringValue(cv.String())
					item.RadiusServerSecretVariable = types.StringNull()
				}
			} else {
				item.RadiusServerSecret = types.StringNull()
				item.RadiusServerSecretVariable = types.StringNull()
			}
			if cValue := v.Get("passphrase.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Passphrase = types.StringNull()

					cv := v.Get("passphrase.vipVariableName")
					item.PassphraseVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Passphrase = types.StringNull()
					item.PassphraseVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("passphrase.vipValue")
					item.Passphrase = types.StringValue(cv.String())
					item.PassphraseVariable = types.StringNull()
				}
			} else {
				item.Passphrase = types.StringNull()
				item.PassphraseVariable = types.StringNull()
			}
			if cValue := v.Get("qos-profile.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.QosProfile = types.StringNull()

					cv := v.Get("qos-profile.vipVariableName")
					item.QosProfileVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.QosProfile = types.StringNull()
					item.QosProfileVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("qos-profile.vipValue")
					item.QosProfile = types.StringValue(cv.String())
					item.QosProfileVariable = types.StringNull()
				}
			} else {
				item.QosProfile = types.StringNull()
				item.QosProfileVariable = types.StringNull()
			}
			data.Ssids = append(data.Ssids, item)
			return true
		})
	} else {
		if len(data.Ssids) > 0 {
			data.Ssids = []CiscoWirelessLANSsids{}
		}
	}
	if value := res.Get(path + "country.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Country = types.StringNull()

			v := res.Get(path + "country.vipVariableName")
			data.CountryVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Country = types.StringNull()
			data.CountryVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "country.vipValue")
			data.Country = types.StringValue(v.String())
			data.CountryVariable = types.StringNull()
		}
	} else {
		data.Country = types.StringNull()
		data.CountryVariable = types.StringNull()
	}
	if value := res.Get(path + "mgmt.username.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Username = types.StringNull()

			v := res.Get(path + "mgmt.username.vipVariableName")
			data.UsernameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Username = types.StringNull()
			data.UsernameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "mgmt.username.vipValue")
			data.Username = types.StringValue(v.String())
			data.UsernameVariable = types.StringNull()
		}
	} else {
		data.Username = types.StringNull()
		data.UsernameVariable = types.StringNull()
	}
	if value := res.Get(path + "mgmt.password.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Password = types.StringNull()

			v := res.Get(path + "mgmt.password.vipVariableName")
			data.PasswordVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Password = types.StringNull()
			data.PasswordVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "mgmt.password.vipValue")
			data.Password = types.StringValue(v.String())
			data.PasswordVariable = types.StringNull()
		}
	} else {
		data.Password = types.StringNull()
		data.PasswordVariable = types.StringNull()
	}
	if value := res.Get(path + "mgmt.address.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ControllerIpAddress = types.StringNull()

			v := res.Get(path + "mgmt.address.vipVariableName")
			data.ControllerIpAddressVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ControllerIpAddress = types.StringNull()
			data.ControllerIpAddressVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "mgmt.address.vipValue")
			data.ControllerIpAddress = types.StringValue(v.String())
			data.ControllerIpAddressVariable = types.StringNull()
		}
	} else {
		data.ControllerIpAddress = types.StringNull()
		data.ControllerIpAddressVariable = types.StringNull()
	}
	if value := res.Get(path + "mgmt.netmask.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ControllerSubnetMask = types.StringNull()

			v := res.Get(path + "mgmt.netmask.vipVariableName")
			data.ControllerSubnetMaskVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ControllerSubnetMask = types.StringNull()
			data.ControllerSubnetMaskVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "mgmt.netmask.vipValue")
			data.ControllerSubnetMask = types.StringValue(v.String())
			data.ControllerSubnetMaskVariable = types.StringNull()
		}
	} else {
		data.ControllerSubnetMask = types.StringNull()
		data.ControllerSubnetMaskVariable = types.StringNull()
	}
	if value := res.Get(path + "mgmt.default-gateway.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ControllerDefaultGateway = types.StringNull()

			v := res.Get(path + "mgmt.default-gateway.vipVariableName")
			data.ControllerDefaultGatewayVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ControllerDefaultGateway = types.StringNull()
			data.ControllerDefaultGatewayVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "mgmt.default-gateway.vipValue")
			data.ControllerDefaultGateway = types.StringValue(v.String())
			data.ControllerDefaultGatewayVariable = types.StringNull()
		}
	} else {
		data.ControllerDefaultGateway = types.StringNull()
		data.ControllerDefaultGatewayVariable = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoWirelessLAN) hasChanges(ctx context.Context, state *CiscoWirelessLAN) bool {
	hasChanges := false
	if !data.Shutdown24ghz.Equal(state.Shutdown24ghz) {
		hasChanges = true
	}
	if !data.Shutdown5ghz.Equal(state.Shutdown5ghz) {
		hasChanges = true
	}
	if len(data.Ssids) != len(state.Ssids) {
		hasChanges = true
	} else {
		for i := range data.Ssids {
			if !data.Ssids[i].WirelessNetworkName.Equal(state.Ssids[i].WirelessNetworkName) {
				hasChanges = true
			}
			if !data.Ssids[i].AdminState.Equal(state.Ssids[i].AdminState) {
				hasChanges = true
			}
			if !data.Ssids[i].BroadcastSsid.Equal(state.Ssids[i].BroadcastSsid) {
				hasChanges = true
			}
			if !data.Ssids[i].VlanId.Equal(state.Ssids[i].VlanId) {
				hasChanges = true
			}
			if !data.Ssids[i].RadioType.Equal(state.Ssids[i].RadioType) {
				hasChanges = true
			}
			if !data.Ssids[i].SecurityType.Equal(state.Ssids[i].SecurityType) {
				hasChanges = true
			}
			if !data.Ssids[i].RadiusServerIp.Equal(state.Ssids[i].RadiusServerIp) {
				hasChanges = true
			}
			if !data.Ssids[i].RadiusServerPort.Equal(state.Ssids[i].RadiusServerPort) {
				hasChanges = true
			}
			if !data.Ssids[i].RadiusServerSecret.Equal(state.Ssids[i].RadiusServerSecret) {
				hasChanges = true
			}
			if !data.Ssids[i].Passphrase.Equal(state.Ssids[i].Passphrase) {
				hasChanges = true
			}
			if !data.Ssids[i].QosProfile.Equal(state.Ssids[i].QosProfile) {
				hasChanges = true
			}
		}
	}
	if !data.Country.Equal(state.Country) {
		hasChanges = true
	}
	if !data.Username.Equal(state.Username) {
		hasChanges = true
	}
	if !data.Password.Equal(state.Password) {
		hasChanges = true
	}
	if !data.ControllerIpAddress.Equal(state.ControllerIpAddress) {
		hasChanges = true
	}
	if !data.ControllerSubnetMask.Equal(state.ControllerSubnetMask) {
		hasChanges = true
	}
	if !data.ControllerDefaultGateway.Equal(state.ControllerDefaultGateway) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
