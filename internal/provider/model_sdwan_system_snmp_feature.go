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
type SystemSNMP struct {
	Id                       types.String                  `tfsdk:"id"`
	Version                  types.Int64                   `tfsdk:"version"`
	Name                     types.String                  `tfsdk:"name"`
	Description              types.String                  `tfsdk:"description"`
	FeatureProfileId         types.String                  `tfsdk:"feature_profile_id"`
	Shutdown                 types.Bool                    `tfsdk:"shutdown"`
	ShutdownVariable         types.String                  `tfsdk:"shutdown_variable"`
	ContactPerson            types.String                  `tfsdk:"contact_person"`
	ContactPersonVariable    types.String                  `tfsdk:"contact_person_variable"`
	LocationOfDevice         types.String                  `tfsdk:"location_of_device"`
	LocationOfDeviceVariable types.String                  `tfsdk:"location_of_device_variable"`
	Views                    []SystemSNMPViews             `tfsdk:"views"`
	Communities              []SystemSNMPCommunities       `tfsdk:"communities"`
	Groups                   []SystemSNMPGroups            `tfsdk:"groups"`
	Users                    []SystemSNMPUsers             `tfsdk:"users"`
	TrapTargetServers        []SystemSNMPTrapTargetServers `tfsdk:"trap_target_servers"`
}

type SystemSNMPViews struct {
	Name types.String          `tfsdk:"name"`
	Oids []SystemSNMPViewsOids `tfsdk:"oids"`
}

type SystemSNMPCommunities struct {
	Name                  types.String `tfsdk:"name"`
	NameVariable          types.String `tfsdk:"name_variable"`
	UserLabel             types.String `tfsdk:"user_label"`
	View                  types.String `tfsdk:"view"`
	ViewVariable          types.String `tfsdk:"view_variable"`
	Authorization         types.String `tfsdk:"authorization"`
	AuthorizationVariable types.String `tfsdk:"authorization_variable"`
}

type SystemSNMPGroups struct {
	Name          types.String `tfsdk:"name"`
	SecurityLevel types.String `tfsdk:"security_level"`
	View          types.String `tfsdk:"view"`
	ViewVariable  types.String `tfsdk:"view_variable"`
}

type SystemSNMPUsers struct {
	Name                           types.String `tfsdk:"name"`
	AuthenticationProtocol         types.String `tfsdk:"authentication_protocol"`
	AuthenticationProtocolVariable types.String `tfsdk:"authentication_protocol_variable"`
	AuthenticationPassword         types.String `tfsdk:"authentication_password"`
	AuthenticationPasswordVariable types.String `tfsdk:"authentication_password_variable"`
	PrivacyProtocol                types.String `tfsdk:"privacy_protocol"`
	PrivacyProtocolVariable        types.String `tfsdk:"privacy_protocol_variable"`
	PrivacyPassword                types.String `tfsdk:"privacy_password"`
	PrivacyPasswordVariable        types.String `tfsdk:"privacy_password_variable"`
	Group                          types.String `tfsdk:"group"`
	GroupVariable                  types.String `tfsdk:"group_variable"`
}

type SystemSNMPTrapTargetServers struct {
	VpnId                   types.Int64  `tfsdk:"vpn_id"`
	VpnIdVariable           types.String `tfsdk:"vpn_id_variable"`
	Ip                      types.String `tfsdk:"ip"`
	IpVariable              types.String `tfsdk:"ip_variable"`
	Port                    types.Int64  `tfsdk:"port"`
	PortVariable            types.String `tfsdk:"port_variable"`
	UserLabel               types.String `tfsdk:"user_label"`
	User                    types.String `tfsdk:"user"`
	UserVariable            types.String `tfsdk:"user_variable"`
	SourceInterface         types.String `tfsdk:"source_interface"`
	SourceInterfaceVariable types.String `tfsdk:"source_interface_variable"`
}

type SystemSNMPViewsOids struct {
	Id              types.String `tfsdk:"id"`
	IdVariable      types.String `tfsdk:"id_variable"`
	Exclude         types.Bool   `tfsdk:"exclude"`
	ExcludeVariable types.String `tfsdk:"exclude_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SystemSNMP) getModel() string {
	return "system_snmp"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SystemSNMP) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/snmp", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SystemSNMP) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.ShutdownVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"shutdown.optionType", "variable")
			body, _ = sjson.Set(body, path+"shutdown.value", data.ShutdownVariable.ValueString())
		}
	} else if data.Shutdown.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"shutdown.optionType", "default")
			body, _ = sjson.Set(body, path+"shutdown.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"shutdown.optionType", "global")
			body, _ = sjson.Set(body, path+"shutdown.value", data.Shutdown.ValueBool())
		}
	}

	if !data.ContactPersonVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"contact.optionType", "variable")
			body, _ = sjson.Set(body, path+"contact.value", data.ContactPersonVariable.ValueString())
		}
	} else if data.ContactPerson.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"contact.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"contact.optionType", "global")
			body, _ = sjson.Set(body, path+"contact.value", data.ContactPerson.ValueString())
		}
	}

	if !data.LocationOfDeviceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"location.optionType", "variable")
			body, _ = sjson.Set(body, path+"location.value", data.LocationOfDeviceVariable.ValueString())
		}
	} else if data.LocationOfDevice.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"location.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"location.optionType", "global")
			body, _ = sjson.Set(body, path+"location.value", data.LocationOfDevice.ValueString())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"view", []interface{}{})
		for _, item := range data.Views {
			itemBody := ""
			if !item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "oid", []interface{}{})
				for _, childItem := range item.Oids {
					itemChildBody := ""

					if !childItem.IdVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "id.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "id.value", childItem.IdVariable.ValueString())
						}
					} else if !childItem.Id.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "id.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "id.value", childItem.Id.ValueString())
						}
					}

					if !childItem.ExcludeVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "exclude.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "exclude.value", childItem.ExcludeVariable.ValueString())
						}
					} else if childItem.Exclude.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "exclude.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "exclude.value", false)
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "exclude.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "exclude.value", childItem.Exclude.ValueBool())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "oid.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"view.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"community", []interface{}{})
		for _, item := range data.Communities {
			itemBody := ""

			if !item.NameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.NameVariable.ValueString())
				}
			} else if !item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
				}
			}
			if !item.UserLabel.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "userLabel.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "userLabel.value", item.UserLabel.ValueString())
				}
			}

			if !item.ViewVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "view.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "view.value", item.ViewVariable.ValueString())
				}
			} else if !item.View.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "view.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "view.value", item.View.ValueString())
				}
			}

			if !item.AuthorizationVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "authorization.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "authorization.value", item.AuthorizationVariable.ValueString())
				}
			} else if !item.Authorization.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "authorization.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "authorization.value", item.Authorization.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"community.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"group", []interface{}{})
		for _, item := range data.Groups {
			itemBody := ""
			if !item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
				}
			}
			if !item.SecurityLevel.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "securityLevel.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "securityLevel.value", item.SecurityLevel.ValueString())
				}
			}

			if !item.ViewVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "view.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "view.value", item.ViewVariable.ValueString())
				}
			} else if !item.View.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "view.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "view.value", item.View.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"group.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"user", []interface{}{})
		for _, item := range data.Users {
			itemBody := ""
			if !item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
				}
			}

			if !item.AuthenticationProtocolVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "auth.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "auth.value", item.AuthenticationProtocolVariable.ValueString())
				}
			} else if item.AuthenticationProtocol.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "auth.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "auth.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "auth.value", item.AuthenticationProtocol.ValueString())
				}
			}

			if !item.AuthenticationPasswordVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "authPassword.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "authPassword.value", item.AuthenticationPasswordVariable.ValueString())
				}
			} else if item.AuthenticationPassword.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "authPassword.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "authPassword.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "authPassword.value", item.AuthenticationPassword.ValueString())
				}
			}

			if !item.PrivacyProtocolVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priv.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "priv.value", item.PrivacyProtocolVariable.ValueString())
				}
			} else if item.PrivacyProtocol.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priv.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "priv.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "priv.value", item.PrivacyProtocol.ValueString())
				}
			}

			if !item.PrivacyPasswordVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "privPassword.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "privPassword.value", item.PrivacyPasswordVariable.ValueString())
				}
			} else if item.PrivacyPassword.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "privPassword.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "privPassword.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "privPassword.value", item.PrivacyPassword.ValueString())
				}
			}

			if !item.GroupVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "group.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "group.value", item.GroupVariable.ValueString())
				}
			} else if !item.Group.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "group.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "group.value", item.Group.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"user.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"target", []interface{}{})
		for _, item := range data.TrapTargetServers {
			itemBody := ""

			if !item.VpnIdVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpnId.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "vpnId.value", item.VpnIdVariable.ValueString())
				}
			} else if !item.VpnId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpnId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vpnId.value", item.VpnId.ValueInt64())
				}
			}

			if !item.IpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ip.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ip.value", item.IpVariable.ValueString())
				}
			} else if !item.Ip.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ip.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ip.value", item.Ip.ValueString())
				}
			}

			if !item.PortVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "port.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "port.value", item.PortVariable.ValueString())
				}
			} else if !item.Port.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "port.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "port.value", item.Port.ValueInt64())
				}
			}
			if !item.UserLabel.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "userLabel.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "userLabel.value", item.UserLabel.ValueString())
				}
			}

			if !item.UserVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "user.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "user.value", item.UserVariable.ValueString())
				}
			} else if !item.User.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "user.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "user.value", item.User.ValueString())
				}
			}

			if !item.SourceInterfaceVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.value", item.SourceInterfaceVariable.ValueString())
				}
			} else if !item.SourceInterface.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.value", item.SourceInterface.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"target.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemSNMP) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Shutdown = types.BoolNull()
	data.ShutdownVariable = types.StringNull()
	if t := res.Get(path + "shutdown.optionType"); t.Exists() {
		va := res.Get(path + "shutdown.value")
		if t.String() == "variable" {
			data.ShutdownVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Shutdown = types.BoolValue(va.Bool())
		}
	}
	data.ContactPerson = types.StringNull()
	data.ContactPersonVariable = types.StringNull()
	if t := res.Get(path + "contact.optionType"); t.Exists() {
		va := res.Get(path + "contact.value")
		if t.String() == "variable" {
			data.ContactPersonVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ContactPerson = types.StringValue(va.String())
		}
	}
	data.LocationOfDevice = types.StringNull()
	data.LocationOfDeviceVariable = types.StringNull()
	if t := res.Get(path + "location.optionType"); t.Exists() {
		va := res.Get(path + "location.value")
		if t.String() == "variable" {
			data.LocationOfDeviceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.LocationOfDevice = types.StringValue(va.String())
		}
	}
	oldViews := data.Views
	if value := res.Get(path + "view"); value.Exists() && len(value.Array()) > 0 {
		data.Views = make([]SystemSNMPViews, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemSNMPViews{}
			item.Name = types.StringNull()

			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("oid"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Oids = make([]SystemSNMPViewsOids, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := SystemSNMPViewsOids{}
					cItem.Id = types.StringNull()
					cItem.IdVariable = types.StringNull()
					if t := cv.Get("id.optionType"); t.Exists() {
						va := cv.Get("id.value")
						if t.String() == "variable" {
							cItem.IdVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Id = types.StringValue(va.String())
						}
					}
					cItem.Exclude = types.BoolNull()
					cItem.ExcludeVariable = types.StringNull()
					if t := cv.Get("exclude.optionType"); t.Exists() {
						va := cv.Get("exclude.value")
						if t.String() == "variable" {
							cItem.ExcludeVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Exclude = types.BoolValue(va.Bool())
						}
					}
					item.Oids = append(item.Oids, cItem)
					return true
				})
			}
			data.Views = append(data.Views, item)
			return true
		})
	} else {
		data.Views = nil
	}
	if !fullRead {
		resultViews := make([]SystemSNMPViews, 0, len(data.Views))
		matchedViews := make([]bool, len(data.Views))
		for _, oldItem := range oldViews {
			for ni := range data.Views {
				if matchedViews[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.Name.ValueString() != data.Views[ni].Name.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedViews[ni] = true
					{
						resultC := make([]SystemSNMPViewsOids, 0, len(data.Views[ni].Oids))
						matchedC := make([]bool, len(data.Views[ni].Oids))
						for _, oldCItem := range oldItem.Oids {
							for nci := range data.Views[ni].Oids {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC && (oldCItem.IdVariable.ValueString() != "" || data.Views[ni].Oids[nci].IdVariable.ValueString() != "") {
									if oldCItem.IdVariable.ValueString() != data.Views[ni].Oids[nci].IdVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.Id.ValueString() != data.Views[ni].Oids[nci].Id.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.Views[ni].Oids[nci])
									break
								}
							}
						}
						for nci := range data.Views[ni].Oids {
							if !matchedC[nci] {
								resultC = append(resultC, data.Views[ni].Oids[nci])
							}
						}
						data.Views[ni].Oids = resultC
					}
					resultViews = append(resultViews, data.Views[ni])
					break
				}
			}
		}
		for ni := range data.Views {
			if !matchedViews[ni] {
				resultViews = append(resultViews, data.Views[ni])
			}
		}
		data.Views = resultViews
	}
	oldCommunities := data.Communities
	if value := res.Get(path + "community"); value.Exists() && len(value.Array()) > 0 {
		data.Communities = make([]SystemSNMPCommunities, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemSNMPCommunities{}
			item.Name = types.StringNull()
			item.NameVariable = types.StringNull()
			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "variable" {
					item.NameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			item.UserLabel = types.StringNull()

			if t := v.Get("userLabel.optionType"); t.Exists() {
				va := v.Get("userLabel.value")
				if t.String() == "global" {
					item.UserLabel = types.StringValue(va.String())
				}
			}
			item.View = types.StringNull()
			item.ViewVariable = types.StringNull()
			if t := v.Get("view.optionType"); t.Exists() {
				va := v.Get("view.value")
				if t.String() == "variable" {
					item.ViewVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.View = types.StringValue(va.String())
				}
			}
			item.Authorization = types.StringNull()
			item.AuthorizationVariable = types.StringNull()
			if t := v.Get("authorization.optionType"); t.Exists() {
				va := v.Get("authorization.value")
				if t.String() == "variable" {
					item.AuthorizationVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Authorization = types.StringValue(va.String())
				}
			}
			data.Communities = append(data.Communities, item)
			return true
		})
	} else {
		data.Communities = nil
	}
	if !fullRead {
		resultCommunities := make([]SystemSNMPCommunities, 0, len(data.Communities))
		matchedCommunities := make([]bool, len(data.Communities))
		for _, oldItem := range oldCommunities {
			for ni := range data.Communities {
				if matchedCommunities[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.UserLabel.ValueString() != data.Communities[ni].UserLabel.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedCommunities[ni] = true
					data.Communities[ni].Name = oldItem.Name
					data.Communities[ni].NameVariable = oldItem.NameVariable
					resultCommunities = append(resultCommunities, data.Communities[ni])
					break
				}
			}
		}
		for ni := range data.Communities {
			if !matchedCommunities[ni] {
				resultCommunities = append(resultCommunities, data.Communities[ni])
			}
		}
		data.Communities = resultCommunities
	}
	oldGroups := data.Groups
	if value := res.Get(path + "group"); value.Exists() && len(value.Array()) > 0 {
		data.Groups = make([]SystemSNMPGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemSNMPGroups{}
			item.Name = types.StringNull()

			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			item.SecurityLevel = types.StringNull()

			if t := v.Get("securityLevel.optionType"); t.Exists() {
				va := v.Get("securityLevel.value")
				if t.String() == "global" {
					item.SecurityLevel = types.StringValue(va.String())
				}
			}
			item.View = types.StringNull()
			item.ViewVariable = types.StringNull()
			if t := v.Get("view.optionType"); t.Exists() {
				va := v.Get("view.value")
				if t.String() == "variable" {
					item.ViewVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.View = types.StringValue(va.String())
				}
			}
			data.Groups = append(data.Groups, item)
			return true
		})
	} else {
		data.Groups = nil
	}
	if !fullRead {
		resultGroups := make([]SystemSNMPGroups, 0, len(data.Groups))
		matchedGroups := make([]bool, len(data.Groups))
		for _, oldItem := range oldGroups {
			for ni := range data.Groups {
				if matchedGroups[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.Name.ValueString() != data.Groups[ni].Name.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedGroups[ni] = true
					resultGroups = append(resultGroups, data.Groups[ni])
					break
				}
			}
		}
		for ni := range data.Groups {
			if !matchedGroups[ni] {
				resultGroups = append(resultGroups, data.Groups[ni])
			}
		}
		data.Groups = resultGroups
	}
	oldUsers := data.Users
	if value := res.Get(path + "user"); value.Exists() && len(value.Array()) > 0 {
		data.Users = make([]SystemSNMPUsers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemSNMPUsers{}
			item.Name = types.StringNull()

			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			item.AuthenticationProtocol = types.StringNull()
			item.AuthenticationProtocolVariable = types.StringNull()
			if t := v.Get("auth.optionType"); t.Exists() {
				va := v.Get("auth.value")
				if t.String() == "variable" {
					item.AuthenticationProtocolVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AuthenticationProtocol = types.StringValue(va.String())
				}
			}
			item.AuthenticationPassword = types.StringNull()
			item.AuthenticationPasswordVariable = types.StringNull()
			if t := v.Get("authPassword.optionType"); t.Exists() {
				va := v.Get("authPassword.value")
				if t.String() == "variable" {
					item.AuthenticationPasswordVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AuthenticationPassword = types.StringValue(va.String())
				}
			}
			item.PrivacyProtocol = types.StringNull()
			item.PrivacyProtocolVariable = types.StringNull()
			if t := v.Get("priv.optionType"); t.Exists() {
				va := v.Get("priv.value")
				if t.String() == "variable" {
					item.PrivacyProtocolVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.PrivacyProtocol = types.StringValue(va.String())
				}
			}
			item.PrivacyPassword = types.StringNull()
			item.PrivacyPasswordVariable = types.StringNull()
			if t := v.Get("privPassword.optionType"); t.Exists() {
				va := v.Get("privPassword.value")
				if t.String() == "variable" {
					item.PrivacyPasswordVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.PrivacyPassword = types.StringValue(va.String())
				}
			}
			item.Group = types.StringNull()
			item.GroupVariable = types.StringNull()
			if t := v.Get("group.optionType"); t.Exists() {
				va := v.Get("group.value")
				if t.String() == "variable" {
					item.GroupVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Group = types.StringValue(va.String())
				}
			}
			data.Users = append(data.Users, item)
			return true
		})
	} else {
		data.Users = nil
	}
	if !fullRead {
		resultUsers := make([]SystemSNMPUsers, 0, len(data.Users))
		matchedUsers := make([]bool, len(data.Users))
		for _, oldItem := range oldUsers {
			for ni := range data.Users {
				if matchedUsers[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.Name.ValueString() != data.Users[ni].Name.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedUsers[ni] = true
					resultUsers = append(resultUsers, data.Users[ni])
					break
				}
			}
		}
		for ni := range data.Users {
			if !matchedUsers[ni] {
				resultUsers = append(resultUsers, data.Users[ni])
			}
		}
		data.Users = resultUsers
	}
	oldTrapTargetServers := data.TrapTargetServers
	if value := res.Get(path + "target"); value.Exists() && len(value.Array()) > 0 {
		data.TrapTargetServers = make([]SystemSNMPTrapTargetServers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemSNMPTrapTargetServers{}
			item.VpnId = types.Int64Null()
			item.VpnIdVariable = types.StringNull()
			if t := v.Get("vpnId.optionType"); t.Exists() {
				va := v.Get("vpnId.value")
				if t.String() == "variable" {
					item.VpnIdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.VpnId = types.Int64Value(va.Int())
				}
			}
			item.Ip = types.StringNull()
			item.IpVariable = types.StringNull()
			if t := v.Get("ip.optionType"); t.Exists() {
				va := v.Get("ip.value")
				if t.String() == "variable" {
					item.IpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Ip = types.StringValue(va.String())
				}
			}
			item.Port = types.Int64Null()
			item.PortVariable = types.StringNull()
			if t := v.Get("port.optionType"); t.Exists() {
				va := v.Get("port.value")
				if t.String() == "variable" {
					item.PortVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Port = types.Int64Value(va.Int())
				}
			}
			item.UserLabel = types.StringNull()

			if t := v.Get("userLabel.optionType"); t.Exists() {
				va := v.Get("userLabel.value")
				if t.String() == "global" {
					item.UserLabel = types.StringValue(va.String())
				}
			}
			item.User = types.StringNull()
			item.UserVariable = types.StringNull()
			if t := v.Get("user.optionType"); t.Exists() {
				va := v.Get("user.value")
				if t.String() == "variable" {
					item.UserVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.User = types.StringValue(va.String())
				}
			}
			item.SourceInterface = types.StringNull()
			item.SourceInterfaceVariable = types.StringNull()
			if t := v.Get("sourceInterface.optionType"); t.Exists() {
				va := v.Get("sourceInterface.value")
				if t.String() == "variable" {
					item.SourceInterfaceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourceInterface = types.StringValue(va.String())
				}
			}
			data.TrapTargetServers = append(data.TrapTargetServers, item)
			return true
		})
	} else {
		data.TrapTargetServers = nil
	}
	if !fullRead {
		resultTrapTargetServers := make([]SystemSNMPTrapTargetServers, 0, len(data.TrapTargetServers))
		matchedTrapTargetServers := make([]bool, len(data.TrapTargetServers))
		for _, oldItem := range oldTrapTargetServers {
			for ni := range data.TrapTargetServers {
				if matchedTrapTargetServers[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.VpnIdVariable.ValueString() != "" || data.TrapTargetServers[ni].VpnIdVariable.ValueString() != "") {
					if oldItem.VpnIdVariable.ValueString() != data.TrapTargetServers[ni].VpnIdVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.VpnId.ValueInt64() != data.TrapTargetServers[ni].VpnId.ValueInt64() {
						keyMatch = false
					}
				}
				if keyMatch && (oldItem.IpVariable.ValueString() != "" || data.TrapTargetServers[ni].IpVariable.ValueString() != "") {
					if oldItem.IpVariable.ValueString() != data.TrapTargetServers[ni].IpVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.Ip.ValueString() != data.TrapTargetServers[ni].Ip.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch && (oldItem.PortVariable.ValueString() != "" || data.TrapTargetServers[ni].PortVariable.ValueString() != "") {
					if oldItem.PortVariable.ValueString() != data.TrapTargetServers[ni].PortVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.Port.ValueInt64() != data.TrapTargetServers[ni].Port.ValueInt64() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedTrapTargetServers[ni] = true
					resultTrapTargetServers = append(resultTrapTargetServers, data.TrapTargetServers[ni])
					break
				}
			}
		}
		for ni := range data.TrapTargetServers {
			if !matchedTrapTargetServers[ni] {
				resultTrapTargetServers = append(resultTrapTargetServers, data.TrapTargetServers[ni])
			}
		}
		data.TrapTargetServers = resultTrapTargetServers
	}
}

// End of section. //template:end fromBody
