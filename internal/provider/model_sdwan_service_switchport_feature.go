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
type ServiceSwitchport struct {
	Id                 types.String                          `tfsdk:"id"`
	Version            types.Int64                           `tfsdk:"version"`
	Name               types.String                          `tfsdk:"name"`
	Description        types.String                          `tfsdk:"description"`
	FeatureProfileId   types.String                          `tfsdk:"feature_profile_id"`
	Interfaces         []ServiceSwitchportInterfaces         `tfsdk:"interfaces"`
	AgeOutTime         types.Int64                           `tfsdk:"age_out_time"`
	AgeOutTimeVariable types.String                          `tfsdk:"age_out_time_variable"`
	StaticMacAddresses []ServiceSwitchportStaticMacAddresses `tfsdk:"static_mac_addresses"`
}

type ServiceSwitchportInterfaces struct {
	InterfaceName                       types.String `tfsdk:"interface_name"`
	InterfaceNameVariable               types.String `tfsdk:"interface_name_variable"`
	Mode                                types.String `tfsdk:"mode"`
	Shutdown                            types.Bool   `tfsdk:"shutdown"`
	ShutdownVariable                    types.String `tfsdk:"shutdown_variable"`
	Speed                               types.String `tfsdk:"speed"`
	SpeedVariable                       types.String `tfsdk:"speed_variable"`
	Duplex                              types.String `tfsdk:"duplex"`
	DuplexVariable                      types.String `tfsdk:"duplex_variable"`
	SwitchportAccessVlan                types.Int64  `tfsdk:"switchport_access_vlan"`
	SwitchportAccessVlanVariable        types.String `tfsdk:"switchport_access_vlan_variable"`
	SwitchportTrunkAllowedVlans         types.String `tfsdk:"switchport_trunk_allowed_vlans"`
	SwitchportTrunkAllowedVlansVariable types.String `tfsdk:"switchport_trunk_allowed_vlans_variable"`
	SwitchportTrunkNativeVlan           types.Int64  `tfsdk:"switchport_trunk_native_vlan"`
	SwitchportTrunkNativeVlanVariable   types.String `tfsdk:"switchport_trunk_native_vlan_variable"`
	EnableDot1x                         types.Bool   `tfsdk:"enable_dot1x"`
	PortControl                         types.String `tfsdk:"port_control"`
	PortControlVariable                 types.String `tfsdk:"port_control_variable"`
	VoiceVlan                           types.Int64  `tfsdk:"voice_vlan"`
	VoiceVlanVariable                   types.String `tfsdk:"voice_vlan_variable"`
	PaeEnable                           types.Bool   `tfsdk:"pae_enable"`
	PaeEnableVariable                   types.String `tfsdk:"pae_enable_variable"`
	MacAuthenticationBypass             types.Bool   `tfsdk:"mac_authentication_bypass"`
	MacAuthenticationBypassVariable     types.String `tfsdk:"mac_authentication_bypass_variable"`
	HostMode                            types.String `tfsdk:"host_mode"`
	HostModeVariable                    types.String `tfsdk:"host_mode_variable"`
	EnablePeriodicReauth                types.Bool   `tfsdk:"enable_periodic_reauth"`
	EnablePeriodicReauthVariable        types.String `tfsdk:"enable_periodic_reauth_variable"`
	Inactivity                          types.Int64  `tfsdk:"inactivity"`
	InactivityVariable                  types.String `tfsdk:"inactivity_variable"`
	Reauthentication                    types.Int64  `tfsdk:"reauthentication"`
	ReauthenticationVariable            types.String `tfsdk:"reauthentication_variable"`
	ControlDirection                    types.String `tfsdk:"control_direction"`
	ControlDirectionVariable            types.String `tfsdk:"control_direction_variable"`
	RestrictedVlan                      types.Int64  `tfsdk:"restricted_vlan"`
	RestrictedVlanVariable              types.String `tfsdk:"restricted_vlan_variable"`
	GuestVlan                           types.Int64  `tfsdk:"guest_vlan"`
	GuestVlanVariable                   types.String `tfsdk:"guest_vlan_variable"`
	CriticalVlan                        types.Int64  `tfsdk:"critical_vlan"`
	CriticalVlanVariable                types.String `tfsdk:"critical_vlan_variable"`
	EnableVoice                         types.Bool   `tfsdk:"enable_voice"`
	EnableVoiceVariable                 types.String `tfsdk:"enable_voice_variable"`
}

type ServiceSwitchportStaticMacAddresses struct {
	MacAddress            types.String `tfsdk:"mac_address"`
	MacAddressVariable    types.String `tfsdk:"mac_address_variable"`
	VlanId                types.Int64  `tfsdk:"vlan_id"`
	VlanIdVariable        types.String `tfsdk:"vlan_id_variable"`
	InterfaceName         types.String `tfsdk:"interface_name"`
	InterfaceNameVariable types.String `tfsdk:"interface_name_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceSwitchport) getModel() string {
	return "service_switchport"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceSwitchport) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/switchport", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceSwitchport) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {
		body, _ = sjson.Set(body, path+"interface", []interface{}{})
		for _, item := range data.Interfaces {
			itemBody := ""

			if !item.InterfaceNameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ifName.value", item.InterfaceNameVariable.ValueString())
				}
			} else if !item.InterfaceName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ifName.value", item.InterfaceName.ValueString())
				}
			}
			if !item.Mode.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "mode.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "mode.value", item.Mode.ValueString())
				}
			}

			if !item.ShutdownVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "shutdown.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "shutdown.value", item.ShutdownVariable.ValueString())
				}
			} else if item.Shutdown.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "shutdown.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "shutdown.value", true)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "shutdown.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "shutdown.value", item.Shutdown.ValueBool())
				}
			}

			if !item.SpeedVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "speed.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "speed.value", item.SpeedVariable.ValueString())
				}
			} else if item.Speed.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "speed.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "speed.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "speed.value", item.Speed.ValueString())
				}
			}

			if !item.DuplexVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "duplex.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "duplex.value", item.DuplexVariable.ValueString())
				}
			} else if item.Duplex.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "duplex.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "duplex.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "duplex.value", item.Duplex.ValueString())
				}
			}

			if !item.SwitchportAccessVlanVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "switchportAccessVlan.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "switchportAccessVlan.value", item.SwitchportAccessVlanVariable.ValueString())
				}
			} else if item.SwitchportAccessVlan.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "switchportAccessVlan.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "switchportAccessVlan.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "switchportAccessVlan.value", item.SwitchportAccessVlan.ValueInt64())
				}
			}

			if !item.SwitchportTrunkAllowedVlansVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "switchportTrunkAllowedVlans.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "switchportTrunkAllowedVlans.value", item.SwitchportTrunkAllowedVlansVariable.ValueString())
				}
			} else if item.SwitchportTrunkAllowedVlans.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "switchportTrunkAllowedVlans.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "switchportTrunkAllowedVlans.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "switchportTrunkAllowedVlans.value", item.SwitchportTrunkAllowedVlans.ValueString())
				}
			}

			if !item.SwitchportTrunkNativeVlanVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "switchportTrunkNativeVlan.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "switchportTrunkNativeVlan.value", item.SwitchportTrunkNativeVlanVariable.ValueString())
				}
			} else if item.SwitchportTrunkNativeVlan.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "switchportTrunkNativeVlan.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "switchportTrunkNativeVlan.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "switchportTrunkNativeVlan.value", item.SwitchportTrunkNativeVlan.ValueInt64())
				}
			}
			if item.EnableDot1x.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "enableDot1x.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "enableDot1x.value", true)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "enableDot1x.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "enableDot1x.value", item.EnableDot1x.ValueBool())
				}
			}

			if !item.PortControlVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "portControl.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "portControl.value", item.PortControlVariable.ValueString())
				}
			} else if item.PortControl.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "portControl.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "portControl.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "portControl.value", item.PortControl.ValueString())
				}
			}

			if !item.VoiceVlanVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "voiceVlan.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "voiceVlan.value", item.VoiceVlanVariable.ValueString())
				}
			} else if item.VoiceVlan.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "voiceVlan.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "voiceVlan.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "voiceVlan.value", item.VoiceVlan.ValueInt64())
				}
			}

			if !item.PaeEnableVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "paeEnable.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "paeEnable.value", item.PaeEnableVariable.ValueString())
				}
			} else if item.PaeEnable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "paeEnable.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "paeEnable.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "paeEnable.value", item.PaeEnable.ValueBool())
				}
			}

			if !item.MacAuthenticationBypassVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "macAuthenticationBypass.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "macAuthenticationBypass.value", item.MacAuthenticationBypassVariable.ValueString())
				}
			} else if item.MacAuthenticationBypass.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "macAuthenticationBypass.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "macAuthenticationBypass.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "macAuthenticationBypass.value", item.MacAuthenticationBypass.ValueBool())
				}
			}

			if !item.HostModeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "hostMode.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "hostMode.value", item.HostModeVariable.ValueString())
				}
			} else if item.HostMode.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "hostMode.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "hostMode.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "hostMode.value", item.HostMode.ValueString())
				}
			}

			if !item.EnablePeriodicReauthVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "enablePeriodicReauth.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "enablePeriodicReauth.value", item.EnablePeriodicReauthVariable.ValueString())
				}
			} else if item.EnablePeriodicReauth.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "enablePeriodicReauth.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "enablePeriodicReauth.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "enablePeriodicReauth.value", item.EnablePeriodicReauth.ValueBool())
				}
			}

			if !item.InactivityVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "inactivity.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "inactivity.value", item.InactivityVariable.ValueString())
				}
			} else if item.Inactivity.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "inactivity.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "inactivity.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "inactivity.value", item.Inactivity.ValueInt64())
				}
			}

			if !item.ReauthenticationVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "reauthentication.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "reauthentication.value", item.ReauthenticationVariable.ValueString())
				}
			} else if item.Reauthentication.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "reauthentication.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "reauthentication.value", 3600)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "reauthentication.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "reauthentication.value", item.Reauthentication.ValueInt64())
				}
			}

			if !item.ControlDirectionVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "controlDirection.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "controlDirection.value", item.ControlDirectionVariable.ValueString())
				}
			} else if item.ControlDirection.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "controlDirection.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "controlDirection.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "controlDirection.value", item.ControlDirection.ValueString())
				}
			}

			if !item.RestrictedVlanVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "restrictedVlan.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "restrictedVlan.value", item.RestrictedVlanVariable.ValueString())
				}
			} else if item.RestrictedVlan.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "restrictedVlan.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "restrictedVlan.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "restrictedVlan.value", item.RestrictedVlan.ValueInt64())
				}
			}

			if !item.GuestVlanVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "guestVlan.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "guestVlan.value", item.GuestVlanVariable.ValueString())
				}
			} else if item.GuestVlan.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "guestVlan.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "guestVlan.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "guestVlan.value", item.GuestVlan.ValueInt64())
				}
			}

			if !item.CriticalVlanVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "criticalVlan.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "criticalVlan.value", item.CriticalVlanVariable.ValueString())
				}
			} else if item.CriticalVlan.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "criticalVlan.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "criticalVlan.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "criticalVlan.value", item.CriticalVlan.ValueInt64())
				}
			}

			if !item.EnableVoiceVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "enableVoice.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "enableVoice.value", item.EnableVoiceVariable.ValueString())
				}
			} else if item.EnableVoice.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "enableVoice.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "enableVoice.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "enableVoice.value", item.EnableVoice.ValueBool())
				}
			}
			body, _ = sjson.SetRaw(body, path+"interface.-1", itemBody)
		}
	}

	if !data.AgeOutTimeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ageTime.optionType", "variable")
			body, _ = sjson.Set(body, path+"ageTime.value", data.AgeOutTimeVariable.ValueString())
		}
	} else if data.AgeOutTime.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ageTime.optionType", "default")
			body, _ = sjson.Set(body, path+"ageTime.value", 300)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ageTime.optionType", "global")
			body, _ = sjson.Set(body, path+"ageTime.value", data.AgeOutTime.ValueInt64())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"staticMacAddress", []interface{}{})
		for _, item := range data.StaticMacAddresses {
			itemBody := ""

			if !item.MacAddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "macaddr.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "macaddr.value", item.MacAddressVariable.ValueString())
				}
			} else if !item.MacAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "macaddr.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "macaddr.value", item.MacAddress.ValueString())
				}
			}

			if !item.VlanIdVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vlan.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "vlan.value", item.VlanIdVariable.ValueString())
				}
			} else if !item.VlanId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vlan.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vlan.value", item.VlanId.ValueInt64())
				}
			}

			if !item.InterfaceNameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ifName.value", item.InterfaceNameVariable.ValueString())
				}
			} else if !item.InterfaceName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ifName.value", item.InterfaceName.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"staticMacAddress.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceSwitchport) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	oldInterfaces := data.Interfaces
	if value := res.Get(path + "interface"); value.Exists() && len(value.Array()) > 0 {
		data.Interfaces = make([]ServiceSwitchportInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceSwitchportInterfaces{}
			item.InterfaceName = types.StringNull()
			item.InterfaceNameVariable = types.StringNull()
			if t := v.Get("ifName.optionType"); t.Exists() {
				va := v.Get("ifName.value")
				if t.String() == "variable" {
					item.InterfaceNameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.InterfaceName = types.StringValue(va.String())
				}
			}
			item.Mode = types.StringNull()

			if t := v.Get("mode.optionType"); t.Exists() {
				va := v.Get("mode.value")
				if t.String() == "global" {
					item.Mode = types.StringValue(va.String())
				}
			}
			item.Shutdown = types.BoolNull()
			item.ShutdownVariable = types.StringNull()
			if t := v.Get("shutdown.optionType"); t.Exists() {
				va := v.Get("shutdown.value")
				if t.String() == "variable" {
					item.ShutdownVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Shutdown = types.BoolValue(va.Bool())
				}
			}
			item.Speed = types.StringNull()
			item.SpeedVariable = types.StringNull()
			if t := v.Get("speed.optionType"); t.Exists() {
				va := v.Get("speed.value")
				if t.String() == "variable" {
					item.SpeedVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Speed = types.StringValue(va.String())
				}
			}
			item.Duplex = types.StringNull()
			item.DuplexVariable = types.StringNull()
			if t := v.Get("duplex.optionType"); t.Exists() {
				va := v.Get("duplex.value")
				if t.String() == "variable" {
					item.DuplexVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Duplex = types.StringValue(va.String())
				}
			}
			item.SwitchportAccessVlan = types.Int64Null()
			item.SwitchportAccessVlanVariable = types.StringNull()
			if t := v.Get("switchportAccessVlan.optionType"); t.Exists() {
				va := v.Get("switchportAccessVlan.value")
				if t.String() == "variable" {
					item.SwitchportAccessVlanVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SwitchportAccessVlan = types.Int64Value(va.Int())
				}
			}
			item.SwitchportTrunkAllowedVlans = types.StringNull()
			item.SwitchportTrunkAllowedVlansVariable = types.StringNull()
			if t := v.Get("switchportTrunkAllowedVlans.optionType"); t.Exists() {
				va := v.Get("switchportTrunkAllowedVlans.value")
				if t.String() == "variable" {
					item.SwitchportTrunkAllowedVlansVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SwitchportTrunkAllowedVlans = types.StringValue(va.String())
				}
			}
			item.SwitchportTrunkNativeVlan = types.Int64Null()
			item.SwitchportTrunkNativeVlanVariable = types.StringNull()
			if t := v.Get("switchportTrunkNativeVlan.optionType"); t.Exists() {
				va := v.Get("switchportTrunkNativeVlan.value")
				if t.String() == "variable" {
					item.SwitchportTrunkNativeVlanVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SwitchportTrunkNativeVlan = types.Int64Value(va.Int())
				}
			}
			item.EnableDot1x = types.BoolNull()

			if t := v.Get("enableDot1x.optionType"); t.Exists() {
				va := v.Get("enableDot1x.value")
				if t.String() == "global" {
					item.EnableDot1x = types.BoolValue(va.Bool())
				}
			}
			item.PortControl = types.StringNull()
			item.PortControlVariable = types.StringNull()
			if t := v.Get("portControl.optionType"); t.Exists() {
				va := v.Get("portControl.value")
				if t.String() == "variable" {
					item.PortControlVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.PortControl = types.StringValue(va.String())
				}
			}
			item.VoiceVlan = types.Int64Null()
			item.VoiceVlanVariable = types.StringNull()
			if t := v.Get("voiceVlan.optionType"); t.Exists() {
				va := v.Get("voiceVlan.value")
				if t.String() == "variable" {
					item.VoiceVlanVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.VoiceVlan = types.Int64Value(va.Int())
				}
			}
			item.PaeEnable = types.BoolNull()
			item.PaeEnableVariable = types.StringNull()
			if t := v.Get("paeEnable.optionType"); t.Exists() {
				va := v.Get("paeEnable.value")
				if t.String() == "variable" {
					item.PaeEnableVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.PaeEnable = types.BoolValue(va.Bool())
				}
			}
			item.MacAuthenticationBypass = types.BoolNull()
			item.MacAuthenticationBypassVariable = types.StringNull()
			if t := v.Get("macAuthenticationBypass.optionType"); t.Exists() {
				va := v.Get("macAuthenticationBypass.value")
				if t.String() == "variable" {
					item.MacAuthenticationBypassVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.MacAuthenticationBypass = types.BoolValue(va.Bool())
				}
			}
			item.HostMode = types.StringNull()
			item.HostModeVariable = types.StringNull()
			if t := v.Get("hostMode.optionType"); t.Exists() {
				va := v.Get("hostMode.value")
				if t.String() == "variable" {
					item.HostModeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.HostMode = types.StringValue(va.String())
				}
			}
			item.EnablePeriodicReauth = types.BoolNull()
			item.EnablePeriodicReauthVariable = types.StringNull()
			if t := v.Get("enablePeriodicReauth.optionType"); t.Exists() {
				va := v.Get("enablePeriodicReauth.value")
				if t.String() == "variable" {
					item.EnablePeriodicReauthVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.EnablePeriodicReauth = types.BoolValue(va.Bool())
				}
			}
			item.Inactivity = types.Int64Null()
			item.InactivityVariable = types.StringNull()
			if t := v.Get("inactivity.optionType"); t.Exists() {
				va := v.Get("inactivity.value")
				if t.String() == "variable" {
					item.InactivityVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Inactivity = types.Int64Value(va.Int())
				}
			}
			item.Reauthentication = types.Int64Null()
			item.ReauthenticationVariable = types.StringNull()
			if t := v.Get("reauthentication.optionType"); t.Exists() {
				va := v.Get("reauthentication.value")
				if t.String() == "variable" {
					item.ReauthenticationVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Reauthentication = types.Int64Value(va.Int())
				}
			}
			item.ControlDirection = types.StringNull()
			item.ControlDirectionVariable = types.StringNull()
			if t := v.Get("controlDirection.optionType"); t.Exists() {
				va := v.Get("controlDirection.value")
				if t.String() == "variable" {
					item.ControlDirectionVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.ControlDirection = types.StringValue(va.String())
				}
			}
			item.RestrictedVlan = types.Int64Null()
			item.RestrictedVlanVariable = types.StringNull()
			if t := v.Get("restrictedVlan.optionType"); t.Exists() {
				va := v.Get("restrictedVlan.value")
				if t.String() == "variable" {
					item.RestrictedVlanVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RestrictedVlan = types.Int64Value(va.Int())
				}
			}
			item.GuestVlan = types.Int64Null()
			item.GuestVlanVariable = types.StringNull()
			if t := v.Get("guestVlan.optionType"); t.Exists() {
				va := v.Get("guestVlan.value")
				if t.String() == "variable" {
					item.GuestVlanVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.GuestVlan = types.Int64Value(va.Int())
				}
			}
			item.CriticalVlan = types.Int64Null()
			item.CriticalVlanVariable = types.StringNull()
			if t := v.Get("criticalVlan.optionType"); t.Exists() {
				va := v.Get("criticalVlan.value")
				if t.String() == "variable" {
					item.CriticalVlanVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.CriticalVlan = types.Int64Value(va.Int())
				}
			}
			item.EnableVoice = types.BoolNull()
			item.EnableVoiceVariable = types.StringNull()
			if t := v.Get("enableVoice.optionType"); t.Exists() {
				va := v.Get("enableVoice.value")
				if t.String() == "variable" {
					item.EnableVoiceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.EnableVoice = types.BoolValue(va.Bool())
				}
			}
			data.Interfaces = append(data.Interfaces, item)
			return true
		})
	} else {
		data.Interfaces = nil
	}
	if !fullRead {
		resultInterfaces := make([]ServiceSwitchportInterfaces, 0, len(data.Interfaces))
		matchedInterfaces := make([]bool, len(data.Interfaces))
		for _, oldItem := range oldInterfaces {
			for ni := range data.Interfaces {
				if matchedInterfaces[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.InterfaceNameVariable.ValueString() != "" || data.Interfaces[ni].InterfaceNameVariable.ValueString() != "") {
					if oldItem.InterfaceNameVariable.ValueString() != data.Interfaces[ni].InterfaceNameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.InterfaceName.ValueString() != data.Interfaces[ni].InterfaceName.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedInterfaces[ni] = true
					resultInterfaces = append(resultInterfaces, data.Interfaces[ni])
					break
				}
			}
		}
		for ni := range data.Interfaces {
			if !matchedInterfaces[ni] {
				resultInterfaces = append(resultInterfaces, data.Interfaces[ni])
			}
		}
		data.Interfaces = resultInterfaces
	}
	data.AgeOutTime = types.Int64Null()
	data.AgeOutTimeVariable = types.StringNull()
	if t := res.Get(path + "ageTime.optionType"); t.Exists() {
		va := res.Get(path + "ageTime.value")
		if t.String() == "variable" {
			data.AgeOutTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AgeOutTime = types.Int64Value(va.Int())
		}
	}
	oldStaticMacAddresses := data.StaticMacAddresses
	if value := res.Get(path + "staticMacAddress"); value.Exists() && len(value.Array()) > 0 {
		data.StaticMacAddresses = make([]ServiceSwitchportStaticMacAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceSwitchportStaticMacAddresses{}
			item.MacAddress = types.StringNull()
			item.MacAddressVariable = types.StringNull()
			if t := v.Get("macaddr.optionType"); t.Exists() {
				va := v.Get("macaddr.value")
				if t.String() == "variable" {
					item.MacAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.MacAddress = types.StringValue(va.String())
				}
			}
			item.VlanId = types.Int64Null()
			item.VlanIdVariable = types.StringNull()
			if t := v.Get("vlan.optionType"); t.Exists() {
				va := v.Get("vlan.value")
				if t.String() == "variable" {
					item.VlanIdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.VlanId = types.Int64Value(va.Int())
				}
			}
			item.InterfaceName = types.StringNull()
			item.InterfaceNameVariable = types.StringNull()
			if t := v.Get("ifName.optionType"); t.Exists() {
				va := v.Get("ifName.value")
				if t.String() == "variable" {
					item.InterfaceNameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.InterfaceName = types.StringValue(va.String())
				}
			}
			data.StaticMacAddresses = append(data.StaticMacAddresses, item)
			return true
		})
	} else {
		data.StaticMacAddresses = nil
	}
	if !fullRead {
		resultStaticMacAddresses := make([]ServiceSwitchportStaticMacAddresses, 0, len(data.StaticMacAddresses))
		matchedStaticMacAddresses := make([]bool, len(data.StaticMacAddresses))
		for _, oldItem := range oldStaticMacAddresses {
			for ni := range data.StaticMacAddresses {
				if matchedStaticMacAddresses[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.VlanIdVariable.ValueString() != "" || data.StaticMacAddresses[ni].VlanIdVariable.ValueString() != "") {
					if oldItem.VlanIdVariable.ValueString() != data.StaticMacAddresses[ni].VlanIdVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.VlanId.ValueInt64() != data.StaticMacAddresses[ni].VlanId.ValueInt64() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedStaticMacAddresses[ni] = true
					resultStaticMacAddresses = append(resultStaticMacAddresses, data.StaticMacAddresses[ni])
					break
				}
			}
		}
		for ni := range data.StaticMacAddresses {
			if !matchedStaticMacAddresses[ni] {
				resultStaticMacAddresses = append(resultStaticMacAddresses, data.StaticMacAddresses[ni])
			}
		}
		data.StaticMacAddresses = resultStaticMacAddresses
	}
}

// End of section. //template:end fromBody
