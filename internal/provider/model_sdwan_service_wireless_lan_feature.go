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
type ServiceWirelessLAN struct {
	Id                       types.String              `tfsdk:"id"`
	Version                  types.Int64               `tfsdk:"version"`
	Name                     types.String              `tfsdk:"name"`
	Description              types.String              `tfsdk:"description"`
	FeatureProfileId         types.String              `tfsdk:"feature_profile_id"`
	Enable24g                types.Bool                `tfsdk:"enable_24g"`
	Enable24gVariable        types.String              `tfsdk:"enable_24g_variable"`
	Enable5g                 types.Bool                `tfsdk:"enable_5g"`
	Enable5gVariable         types.String              `tfsdk:"enable_5g_variable"`
	Ssids                    []ServiceWirelessLANSsids `tfsdk:"ssids"`
	Country                  types.String              `tfsdk:"country"`
	CountryVariable          types.String              `tfsdk:"country_variable"`
	Username                 types.String              `tfsdk:"username"`
	UsernameVariable         types.String              `tfsdk:"username_variable"`
	Password                 types.String              `tfsdk:"password"`
	PasswordVariable         types.String              `tfsdk:"password_variable"`
	MeDynamicIpEnabled       types.Bool                `tfsdk:"me_dynamic_ip_enabled"`
	MeIpv4Address            types.String              `tfsdk:"me_ipv4_address"`
	MeIpv4AddressVariable    types.String              `tfsdk:"me_ipv4_address_variable"`
	MeSubnetMask             types.String              `tfsdk:"me_subnet_mask"`
	MeSubnetMaskVariable     types.String              `tfsdk:"me_subnet_mask_variable"`
	MeDefaultGateway         types.String              `tfsdk:"me_default_gateway"`
	MeDefaultGatewayVariable types.String              `tfsdk:"me_default_gateway_variable"`
}

type ServiceWirelessLANSsids struct {
	SsidName                   types.String `tfsdk:"ssid_name"`
	AdminState                 types.Bool   `tfsdk:"admin_state"`
	AdminStateVariable         types.String `tfsdk:"admin_state_variable"`
	BroadcastSsid              types.Bool   `tfsdk:"broadcast_ssid"`
	BroadcastSsidVariable      types.String `tfsdk:"broadcast_ssid_variable"`
	VlanId                     types.Int64  `tfsdk:"vlan_id"`
	VlanIdVariable             types.String `tfsdk:"vlan_id_variable"`
	RadioType                  types.String `tfsdk:"radio_type"`
	RadioTypeVariable          types.String `tfsdk:"radio_type_variable"`
	SecurityType               types.String `tfsdk:"security_type"`
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
func (data ServiceWirelessLAN) getModel() string {
	return "service_wireless_lan"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceWirelessLAN) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/wirelesslan", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceWirelessLAN) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.Enable24gVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"enable24G.optionType", "variable")
			body, _ = sjson.Set(body, path+"enable24G.value", data.Enable24gVariable.ValueString())
		}
	} else if data.Enable24g.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"enable24G.optionType", "default")
			body, _ = sjson.Set(body, path+"enable24G.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"enable24G.optionType", "global")
			body, _ = sjson.Set(body, path+"enable24G.value", data.Enable24g.ValueBool())
		}
	}

	if !data.Enable5gVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"enable5G.optionType", "variable")
			body, _ = sjson.Set(body, path+"enable5G.value", data.Enable5gVariable.ValueString())
		}
	} else if data.Enable5g.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"enable5G.optionType", "default")
			body, _ = sjson.Set(body, path+"enable5G.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"enable5G.optionType", "global")
			body, _ = sjson.Set(body, path+"enable5G.value", data.Enable5g.ValueBool())
		}
	}
	if true {

		for _, item := range data.Ssids {
			itemBody := ""
			if !item.SsidName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.SsidName.ValueString())
				}
			}

			if !item.AdminStateVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "adminState.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "adminState.value", item.AdminStateVariable.ValueString())
				}
			} else if item.AdminState.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "adminState.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "adminState.value", true)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "adminState.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "adminState.value", item.AdminState.ValueBool())
				}
			}

			if !item.BroadcastSsidVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "broadcastSsid.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "broadcastSsid.value", item.BroadcastSsidVariable.ValueString())
				}
			} else if item.BroadcastSsid.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "broadcastSsid.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "broadcastSsid.value", true)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "broadcastSsid.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "broadcastSsid.value", item.BroadcastSsid.ValueBool())
				}
			}

			if !item.VlanIdVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vlanId.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "vlanId.value", item.VlanIdVariable.ValueString())
				}
			} else if !item.VlanId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vlanId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vlanId.value", item.VlanId.ValueInt64())
				}
			}

			if !item.RadioTypeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "radioType.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "radioType.value", item.RadioTypeVariable.ValueString())
				}
			} else if item.RadioType.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "radioType.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "radioType.value", "all")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "radioType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "radioType.value", item.RadioType.ValueString())
				}
			}
			if !item.SecurityType.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "securityConfig.securityType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "securityConfig.securityType.value", item.SecurityType.ValueString())
				}
			}

			if !item.RadiusServerIpVariable.IsNull() {
				if true && item.SecurityType.ValueString() == "enterprise" {
					itemBody, _ = sjson.Set(itemBody, "securityConfig.radiusServerIp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "securityConfig.radiusServerIp.value", item.RadiusServerIpVariable.ValueString())
				}
			} else if !item.RadiusServerIp.IsNull() {
				if true && item.SecurityType.ValueString() == "enterprise" {
					itemBody, _ = sjson.Set(itemBody, "securityConfig.radiusServerIp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "securityConfig.radiusServerIp.value", item.RadiusServerIp.ValueString())
				}
			}

			if !item.RadiusServerPortVariable.IsNull() {
				if true && item.SecurityType.ValueString() == "enterprise" {
					itemBody, _ = sjson.Set(itemBody, "securityConfig.radiusServerPort.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "securityConfig.radiusServerPort.value", item.RadiusServerPortVariable.ValueString())
				}
			} else if item.RadiusServerPort.IsNull() {
				if true && item.SecurityType.ValueString() == "enterprise" {
					itemBody, _ = sjson.Set(itemBody, "securityConfig.radiusServerPort.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "securityConfig.radiusServerPort.value", 1812)
				}
			} else {
				if true && item.SecurityType.ValueString() == "enterprise" {
					itemBody, _ = sjson.Set(itemBody, "securityConfig.radiusServerPort.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "securityConfig.radiusServerPort.value", item.RadiusServerPort.ValueInt64())
				}
			}

			if !item.RadiusServerSecretVariable.IsNull() {
				if true && item.SecurityType.ValueString() == "enterprise" {
					itemBody, _ = sjson.Set(itemBody, "securityConfig.radiusServerSecret.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "securityConfig.radiusServerSecret.value", item.RadiusServerSecretVariable.ValueString())
				}
			} else if !item.RadiusServerSecret.IsNull() {
				if true && item.SecurityType.ValueString() == "enterprise" {
					itemBody, _ = sjson.Set(itemBody, "securityConfig.radiusServerSecret.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "securityConfig.radiusServerSecret.value", item.RadiusServerSecret.ValueString())
				}
			}

			if !item.PassphraseVariable.IsNull() {
				if true && item.SecurityType.ValueString() == "personal" {
					itemBody, _ = sjson.Set(itemBody, "securityConfig.passphrase.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "securityConfig.passphrase.value", item.PassphraseVariable.ValueString())
				}
			} else if !item.Passphrase.IsNull() {
				if true && item.SecurityType.ValueString() == "personal" {
					itemBody, _ = sjson.Set(itemBody, "securityConfig.passphrase.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "securityConfig.passphrase.value", item.Passphrase.ValueString())
				}
			}

			if !item.QosProfileVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "qosProfile.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "qosProfile.value", item.QosProfileVariable.ValueString())
				}
			} else if item.QosProfile.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "qosProfile.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "qosProfile.value", "silver")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "qosProfile.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "qosProfile.value", item.QosProfile.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"ssid.-1", itemBody)
		}
	}

	if !data.CountryVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"country.optionType", "variable")
			body, _ = sjson.Set(body, path+"country.value", data.CountryVariable.ValueString())
		}
	} else if !data.Country.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"country.optionType", "global")
			body, _ = sjson.Set(body, path+"country.value", data.Country.ValueString())
		}
	}

	if !data.UsernameVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"username.optionType", "variable")
			body, _ = sjson.Set(body, path+"username.value", data.UsernameVariable.ValueString())
		}
	} else if !data.Username.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"username.optionType", "global")
			body, _ = sjson.Set(body, path+"username.value", data.Username.ValueString())
		}
	}

	if !data.PasswordVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"password.optionType", "variable")
			body, _ = sjson.Set(body, path+"password.value", data.PasswordVariable.ValueString())
		}
	} else if !data.Password.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"password.optionType", "global")
			body, _ = sjson.Set(body, path+"password.value", data.Password.ValueString())
		}
	}
	if !data.MeDynamicIpEnabled.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"meIpConfig.meDynamicIpEnabled.optionType", "global")
			body, _ = sjson.Set(body, path+"meIpConfig.meDynamicIpEnabled.value", data.MeDynamicIpEnabled.ValueBool())
		}
	}

	if !data.MeIpv4AddressVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"meIpConfig.meStaticIpCfg.meIpv4Address.optionType", "variable")
			body, _ = sjson.Set(body, path+"meIpConfig.meStaticIpCfg.meIpv4Address.value", data.MeIpv4AddressVariable.ValueString())
		}
	} else if !data.MeIpv4Address.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"meIpConfig.meStaticIpCfg.meIpv4Address.optionType", "global")
			body, _ = sjson.Set(body, path+"meIpConfig.meStaticIpCfg.meIpv4Address.value", data.MeIpv4Address.ValueString())
		}
	}

	if !data.MeSubnetMaskVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"meIpConfig.meStaticIpCfg.netmask.optionType", "variable")
			body, _ = sjson.Set(body, path+"meIpConfig.meStaticIpCfg.netmask.value", data.MeSubnetMaskVariable.ValueString())
		}
	} else if !data.MeSubnetMask.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"meIpConfig.meStaticIpCfg.netmask.optionType", "global")
			body, _ = sjson.Set(body, path+"meIpConfig.meStaticIpCfg.netmask.value", data.MeSubnetMask.ValueString())
		}
	}

	if !data.MeDefaultGatewayVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"meIpConfig.meStaticIpCfg.defaultGateway.optionType", "variable")
			body, _ = sjson.Set(body, path+"meIpConfig.meStaticIpCfg.defaultGateway.value", data.MeDefaultGatewayVariable.ValueString())
		}
	} else if !data.MeDefaultGateway.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"meIpConfig.meStaticIpCfg.defaultGateway.optionType", "global")
			body, _ = sjson.Set(body, path+"meIpConfig.meStaticIpCfg.defaultGateway.value", data.MeDefaultGateway.ValueString())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceWirelessLAN) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Enable24g = types.BoolNull()
	data.Enable24gVariable = types.StringNull()
	if t := res.Get(path + "enable24G.optionType"); t.Exists() {
		va := res.Get(path + "enable24G.value")
		if t.String() == "variable" {
			data.Enable24gVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Enable24g = types.BoolValue(va.Bool())
		}
	}
	data.Enable5g = types.BoolNull()
	data.Enable5gVariable = types.StringNull()
	if t := res.Get(path + "enable5G.optionType"); t.Exists() {
		va := res.Get(path + "enable5G.value")
		if t.String() == "variable" {
			data.Enable5gVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Enable5g = types.BoolValue(va.Bool())
		}
	}
	if value := res.Get(path + "ssid"); value.Exists() {
		data.Ssids = make([]ServiceWirelessLANSsids, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceWirelessLANSsids{}
			item.SsidName = types.StringNull()

			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "global" {
					item.SsidName = types.StringValue(va.String())
				}
			}
			item.AdminState = types.BoolNull()
			item.AdminStateVariable = types.StringNull()
			if t := v.Get("adminState.optionType"); t.Exists() {
				va := v.Get("adminState.value")
				if t.String() == "variable" {
					item.AdminStateVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AdminState = types.BoolValue(va.Bool())
				}
			}
			item.BroadcastSsid = types.BoolNull()
			item.BroadcastSsidVariable = types.StringNull()
			if t := v.Get("broadcastSsid.optionType"); t.Exists() {
				va := v.Get("broadcastSsid.value")
				if t.String() == "variable" {
					item.BroadcastSsidVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.BroadcastSsid = types.BoolValue(va.Bool())
				}
			}
			item.VlanId = types.Int64Null()
			item.VlanIdVariable = types.StringNull()
			if t := v.Get("vlanId.optionType"); t.Exists() {
				va := v.Get("vlanId.value")
				if t.String() == "variable" {
					item.VlanIdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.VlanId = types.Int64Value(va.Int())
				}
			}
			item.RadioType = types.StringNull()
			item.RadioTypeVariable = types.StringNull()
			if t := v.Get("radioType.optionType"); t.Exists() {
				va := v.Get("radioType.value")
				if t.String() == "variable" {
					item.RadioTypeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RadioType = types.StringValue(va.String())
				}
			}
			item.SecurityType = types.StringNull()

			if t := v.Get("securityConfig.securityType.optionType"); t.Exists() {
				va := v.Get("securityConfig.securityType.value")
				if t.String() == "global" {
					item.SecurityType = types.StringValue(va.String())
				}
			}
			item.RadiusServerIp = types.StringNull()
			item.RadiusServerIpVariable = types.StringNull()
			if t := v.Get("securityConfig.radiusServerIp.optionType"); t.Exists() {
				va := v.Get("securityConfig.radiusServerIp.value")
				if t.String() == "variable" {
					item.RadiusServerIpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RadiusServerIp = types.StringValue(va.String())
				}
			}
			item.RadiusServerPort = types.Int64Null()
			item.RadiusServerPortVariable = types.StringNull()
			if t := v.Get("securityConfig.radiusServerPort.optionType"); t.Exists() {
				va := v.Get("securityConfig.radiusServerPort.value")
				if t.String() == "variable" {
					item.RadiusServerPortVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RadiusServerPort = types.Int64Value(va.Int())
				}
			}
			item.RadiusServerSecret = types.StringNull()
			item.RadiusServerSecretVariable = types.StringNull()
			if t := v.Get("securityConfig.radiusServerSecret.optionType"); t.Exists() {
				va := v.Get("securityConfig.radiusServerSecret.value")
				if t.String() == "variable" {
					item.RadiusServerSecretVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RadiusServerSecret = types.StringValue(va.String())
				}
			}
			item.Passphrase = types.StringNull()
			item.PassphraseVariable = types.StringNull()
			if t := v.Get("securityConfig.passphrase.optionType"); t.Exists() {
				va := v.Get("securityConfig.passphrase.value")
				if t.String() == "variable" {
					item.PassphraseVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Passphrase = types.StringValue(va.String())
				}
			}
			item.QosProfile = types.StringNull()
			item.QosProfileVariable = types.StringNull()
			if t := v.Get("qosProfile.optionType"); t.Exists() {
				va := v.Get("qosProfile.value")
				if t.String() == "variable" {
					item.QosProfileVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.QosProfile = types.StringValue(va.String())
				}
			}
			data.Ssids = append(data.Ssids, item)
			return true
		})
	}
	data.Country = types.StringNull()
	data.CountryVariable = types.StringNull()
	if t := res.Get(path + "country.optionType"); t.Exists() {
		va := res.Get(path + "country.value")
		if t.String() == "variable" {
			data.CountryVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Country = types.StringValue(va.String())
		}
	}
	data.Username = types.StringNull()
	data.UsernameVariable = types.StringNull()
	if t := res.Get(path + "username.optionType"); t.Exists() {
		va := res.Get(path + "username.value")
		if t.String() == "variable" {
			data.UsernameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Username = types.StringValue(va.String())
		}
	}
	data.Password = types.StringNull()
	data.PasswordVariable = types.StringNull()
	if t := res.Get(path + "password.optionType"); t.Exists() {
		va := res.Get(path + "password.value")
		if t.String() == "variable" {
			data.PasswordVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Password = types.StringValue(va.String())
		}
	}
	data.MeDynamicIpEnabled = types.BoolNull()

	if t := res.Get(path + "meIpConfig.meDynamicIpEnabled.optionType"); t.Exists() {
		va := res.Get(path + "meIpConfig.meDynamicIpEnabled.value")
		if t.String() == "global" {
			data.MeDynamicIpEnabled = types.BoolValue(va.Bool())
		}
	}
	data.MeIpv4Address = types.StringNull()
	data.MeIpv4AddressVariable = types.StringNull()
	if t := res.Get(path + "meIpConfig.meStaticIpCfg.meIpv4Address.optionType"); t.Exists() {
		va := res.Get(path + "meIpConfig.meStaticIpCfg.meIpv4Address.value")
		if t.String() == "variable" {
			data.MeIpv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MeIpv4Address = types.StringValue(va.String())
		}
	}
	data.MeSubnetMask = types.StringNull()
	data.MeSubnetMaskVariable = types.StringNull()
	if t := res.Get(path + "meIpConfig.meStaticIpCfg.netmask.optionType"); t.Exists() {
		va := res.Get(path + "meIpConfig.meStaticIpCfg.netmask.value")
		if t.String() == "variable" {
			data.MeSubnetMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MeSubnetMask = types.StringValue(va.String())
		}
	}
	data.MeDefaultGateway = types.StringNull()
	data.MeDefaultGatewayVariable = types.StringNull()
	if t := res.Get(path + "meIpConfig.meStaticIpCfg.defaultGateway.optionType"); t.Exists() {
		va := res.Get(path + "meIpConfig.meStaticIpCfg.defaultGateway.value")
		if t.String() == "variable" {
			data.MeDefaultGatewayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MeDefaultGateway = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ServiceWirelessLAN) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Enable24g = types.BoolNull()
	data.Enable24gVariable = types.StringNull()
	if t := res.Get(path + "enable24G.optionType"); t.Exists() {
		va := res.Get(path + "enable24G.value")
		if t.String() == "variable" {
			data.Enable24gVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Enable24g = types.BoolValue(va.Bool())
		}
	}
	data.Enable5g = types.BoolNull()
	data.Enable5gVariable = types.StringNull()
	if t := res.Get(path + "enable5G.optionType"); t.Exists() {
		va := res.Get(path + "enable5G.value")
		if t.String() == "variable" {
			data.Enable5gVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Enable5g = types.BoolValue(va.Bool())
		}
	}
	for i := range data.Ssids {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Ssids[i].SsidName.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "ssid").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Ssids[i].SsidName = types.StringNull()

		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "global" {
				data.Ssids[i].SsidName = types.StringValue(va.String())
			}
		}
		data.Ssids[i].AdminState = types.BoolNull()
		data.Ssids[i].AdminStateVariable = types.StringNull()
		if t := r.Get("adminState.optionType"); t.Exists() {
			va := r.Get("adminState.value")
			if t.String() == "variable" {
				data.Ssids[i].AdminStateVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ssids[i].AdminState = types.BoolValue(va.Bool())
			}
		}
		data.Ssids[i].BroadcastSsid = types.BoolNull()
		data.Ssids[i].BroadcastSsidVariable = types.StringNull()
		if t := r.Get("broadcastSsid.optionType"); t.Exists() {
			va := r.Get("broadcastSsid.value")
			if t.String() == "variable" {
				data.Ssids[i].BroadcastSsidVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ssids[i].BroadcastSsid = types.BoolValue(va.Bool())
			}
		}
		data.Ssids[i].VlanId = types.Int64Null()
		data.Ssids[i].VlanIdVariable = types.StringNull()
		if t := r.Get("vlanId.optionType"); t.Exists() {
			va := r.Get("vlanId.value")
			if t.String() == "variable" {
				data.Ssids[i].VlanIdVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ssids[i].VlanId = types.Int64Value(va.Int())
			}
		}
		data.Ssids[i].RadioType = types.StringNull()
		data.Ssids[i].RadioTypeVariable = types.StringNull()
		if t := r.Get("radioType.optionType"); t.Exists() {
			va := r.Get("radioType.value")
			if t.String() == "variable" {
				data.Ssids[i].RadioTypeVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ssids[i].RadioType = types.StringValue(va.String())
			}
		}
		data.Ssids[i].SecurityType = types.StringNull()

		if t := r.Get("securityConfig.securityType.optionType"); t.Exists() {
			va := r.Get("securityConfig.securityType.value")
			if t.String() == "global" {
				data.Ssids[i].SecurityType = types.StringValue(va.String())
			}
		}
		data.Ssids[i].RadiusServerIp = types.StringNull()
		data.Ssids[i].RadiusServerIpVariable = types.StringNull()
		if t := r.Get("securityConfig.radiusServerIp.optionType"); t.Exists() {
			va := r.Get("securityConfig.radiusServerIp.value")
			if t.String() == "variable" {
				data.Ssids[i].RadiusServerIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ssids[i].RadiusServerIp = types.StringValue(va.String())
			}
		}
		data.Ssids[i].RadiusServerPort = types.Int64Null()
		data.Ssids[i].RadiusServerPortVariable = types.StringNull()
		if t := r.Get("securityConfig.radiusServerPort.optionType"); t.Exists() {
			va := r.Get("securityConfig.radiusServerPort.value")
			if t.String() == "variable" {
				data.Ssids[i].RadiusServerPortVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ssids[i].RadiusServerPort = types.Int64Value(va.Int())
			}
		}
		data.Ssids[i].RadiusServerSecret = types.StringNull()
		data.Ssids[i].RadiusServerSecretVariable = types.StringNull()
		if t := r.Get("securityConfig.radiusServerSecret.optionType"); t.Exists() {
			va := r.Get("securityConfig.radiusServerSecret.value")
			if t.String() == "variable" {
				data.Ssids[i].RadiusServerSecretVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ssids[i].RadiusServerSecret = types.StringValue(va.String())
			}
		}
		data.Ssids[i].Passphrase = types.StringNull()
		data.Ssids[i].PassphraseVariable = types.StringNull()
		if t := r.Get("securityConfig.passphrase.optionType"); t.Exists() {
			va := r.Get("securityConfig.passphrase.value")
			if t.String() == "variable" {
				data.Ssids[i].PassphraseVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ssids[i].Passphrase = types.StringValue(va.String())
			}
		}
		data.Ssids[i].QosProfile = types.StringNull()
		data.Ssids[i].QosProfileVariable = types.StringNull()
		if t := r.Get("qosProfile.optionType"); t.Exists() {
			va := r.Get("qosProfile.value")
			if t.String() == "variable" {
				data.Ssids[i].QosProfileVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ssids[i].QosProfile = types.StringValue(va.String())
			}
		}
	}
	data.Country = types.StringNull()
	data.CountryVariable = types.StringNull()
	if t := res.Get(path + "country.optionType"); t.Exists() {
		va := res.Get(path + "country.value")
		if t.String() == "variable" {
			data.CountryVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Country = types.StringValue(va.String())
		}
	}
	data.Username = types.StringNull()
	data.UsernameVariable = types.StringNull()
	if t := res.Get(path + "username.optionType"); t.Exists() {
		va := res.Get(path + "username.value")
		if t.String() == "variable" {
			data.UsernameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Username = types.StringValue(va.String())
		}
	}
	data.Password = types.StringNull()
	data.PasswordVariable = types.StringNull()
	if t := res.Get(path + "password.optionType"); t.Exists() {
		va := res.Get(path + "password.value")
		if t.String() == "variable" {
			data.PasswordVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Password = types.StringValue(va.String())
		}
	}
	data.MeDynamicIpEnabled = types.BoolNull()

	if t := res.Get(path + "meIpConfig.meDynamicIpEnabled.optionType"); t.Exists() {
		va := res.Get(path + "meIpConfig.meDynamicIpEnabled.value")
		if t.String() == "global" {
			data.MeDynamicIpEnabled = types.BoolValue(va.Bool())
		}
	}
	data.MeIpv4Address = types.StringNull()
	data.MeIpv4AddressVariable = types.StringNull()
	if t := res.Get(path + "meIpConfig.meStaticIpCfg.meIpv4Address.optionType"); t.Exists() {
		va := res.Get(path + "meIpConfig.meStaticIpCfg.meIpv4Address.value")
		if t.String() == "variable" {
			data.MeIpv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MeIpv4Address = types.StringValue(va.String())
		}
	}
	data.MeSubnetMask = types.StringNull()
	data.MeSubnetMaskVariable = types.StringNull()
	if t := res.Get(path + "meIpConfig.meStaticIpCfg.netmask.optionType"); t.Exists() {
		va := res.Get(path + "meIpConfig.meStaticIpCfg.netmask.value")
		if t.String() == "variable" {
			data.MeSubnetMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MeSubnetMask = types.StringValue(va.String())
		}
	}
	data.MeDefaultGateway = types.StringNull()
	data.MeDefaultGatewayVariable = types.StringNull()
	if t := res.Get(path + "meIpConfig.meStaticIpCfg.defaultGateway.optionType"); t.Exists() {
		va := res.Get(path + "meIpConfig.meStaticIpCfg.defaultGateway.value")
		if t.String() == "variable" {
			data.MeDefaultGatewayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MeDefaultGateway = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end updateFromBody
