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
	"strconv"

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
			if !item.Name.IsNull() {
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
func (data *SystemSNMP) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "view"); value.Exists() {
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
			if cValue := v.Get("oid"); cValue.Exists() {
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
	}
	if value := res.Get(path + "community"); value.Exists() {
		data.Communities = make([]SystemSNMPCommunities, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemSNMPCommunities{}
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
	}
	if value := res.Get(path + "group"); value.Exists() {
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
	}
	if value := res.Get(path + "user"); value.Exists() {
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
	}
	if value := res.Get(path + "target"); value.Exists() {
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
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SystemSNMP) updateFromBody(ctx context.Context, res gjson.Result) {
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
	for i := range data.Views {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Views[i].Name.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "view").ForEach(
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
		data.Views[i].Name = types.StringNull()

		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "global" {
				data.Views[i].Name = types.StringValue(va.String())
			}
		}
		for ci := range data.Views[i].Oids {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Views[i].Oids[ci].Id.ValueString()}
			keyValuesVariables := [...]string{data.Views[i].Oids[ci].IdVariable.ValueString()}

			var cr gjson.Result
			r.Get("oid").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			data.Views[i].Oids[ci].Id = types.StringNull()
			data.Views[i].Oids[ci].IdVariable = types.StringNull()
			if t := cr.Get("id.optionType"); t.Exists() {
				va := cr.Get("id.value")
				if t.String() == "variable" {
					data.Views[i].Oids[ci].IdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Views[i].Oids[ci].Id = types.StringValue(va.String())
				}
			}
			data.Views[i].Oids[ci].Exclude = types.BoolNull()
			data.Views[i].Oids[ci].ExcludeVariable = types.StringNull()
			if t := cr.Get("exclude.optionType"); t.Exists() {
				va := cr.Get("exclude.value")
				if t.String() == "variable" {
					data.Views[i].Oids[ci].ExcludeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Views[i].Oids[ci].Exclude = types.BoolValue(va.Bool())
				}
			}
		}
	}
	for i := range data.Communities {
		keys := [...]string{"userLabel"}
		keyValues := [...]string{data.Communities[i].UserLabel.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "community").ForEach(
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
		data.Communities[i].UserLabel = types.StringNull()

		if t := r.Get("userLabel.optionType"); t.Exists() {
			va := r.Get("userLabel.value")
			if t.String() == "global" {
				data.Communities[i].UserLabel = types.StringValue(va.String())
			}
		}
		data.Communities[i].View = types.StringNull()
		data.Communities[i].ViewVariable = types.StringNull()
		if t := r.Get("view.optionType"); t.Exists() {
			va := r.Get("view.value")
			if t.String() == "variable" {
				data.Communities[i].ViewVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Communities[i].View = types.StringValue(va.String())
			}
		}
		data.Communities[i].Authorization = types.StringNull()
		data.Communities[i].AuthorizationVariable = types.StringNull()
		if t := r.Get("authorization.optionType"); t.Exists() {
			va := r.Get("authorization.value")
			if t.String() == "variable" {
				data.Communities[i].AuthorizationVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Communities[i].Authorization = types.StringValue(va.String())
			}
		}
	}
	for i := range data.Groups {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Groups[i].Name.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "group").ForEach(
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
		data.Groups[i].Name = types.StringNull()

		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "global" {
				data.Groups[i].Name = types.StringValue(va.String())
			}
		}
		data.Groups[i].SecurityLevel = types.StringNull()

		if t := r.Get("securityLevel.optionType"); t.Exists() {
			va := r.Get("securityLevel.value")
			if t.String() == "global" {
				data.Groups[i].SecurityLevel = types.StringValue(va.String())
			}
		}
		data.Groups[i].View = types.StringNull()
		data.Groups[i].ViewVariable = types.StringNull()
		if t := r.Get("view.optionType"); t.Exists() {
			va := r.Get("view.value")
			if t.String() == "variable" {
				data.Groups[i].ViewVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Groups[i].View = types.StringValue(va.String())
			}
		}
	}
	for i := range data.Users {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Users[i].Name.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "user").ForEach(
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
		data.Users[i].Name = types.StringNull()

		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "global" {
				data.Users[i].Name = types.StringValue(va.String())
			}
		}
		data.Users[i].AuthenticationProtocol = types.StringNull()
		data.Users[i].AuthenticationProtocolVariable = types.StringNull()
		if t := r.Get("auth.optionType"); t.Exists() {
			va := r.Get("auth.value")
			if t.String() == "variable" {
				data.Users[i].AuthenticationProtocolVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Users[i].AuthenticationProtocol = types.StringValue(va.String())
			}
		}
		data.Users[i].AuthenticationPassword = types.StringNull()
		data.Users[i].AuthenticationPasswordVariable = types.StringNull()
		if t := r.Get("authPassword.optionType"); t.Exists() {
			va := r.Get("authPassword.value")
			if t.String() == "variable" {
				data.Users[i].AuthenticationPasswordVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Users[i].AuthenticationPassword = types.StringValue(va.String())
			}
		}
		data.Users[i].PrivacyProtocol = types.StringNull()
		data.Users[i].PrivacyProtocolVariable = types.StringNull()
		if t := r.Get("priv.optionType"); t.Exists() {
			va := r.Get("priv.value")
			if t.String() == "variable" {
				data.Users[i].PrivacyProtocolVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Users[i].PrivacyProtocol = types.StringValue(va.String())
			}
		}
		data.Users[i].PrivacyPassword = types.StringNull()
		data.Users[i].PrivacyPasswordVariable = types.StringNull()
		if t := r.Get("privPassword.optionType"); t.Exists() {
			va := r.Get("privPassword.value")
			if t.String() == "variable" {
				data.Users[i].PrivacyPasswordVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Users[i].PrivacyPassword = types.StringValue(va.String())
			}
		}
		data.Users[i].Group = types.StringNull()
		data.Users[i].GroupVariable = types.StringNull()
		if t := r.Get("group.optionType"); t.Exists() {
			va := r.Get("group.value")
			if t.String() == "variable" {
				data.Users[i].GroupVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Users[i].Group = types.StringValue(va.String())
			}
		}
	}
	for i := range data.TrapTargetServers {
		keys := [...]string{"vpnId", "ip", "port"}
		keyValues := [...]string{strconv.FormatInt(data.TrapTargetServers[i].VpnId.ValueInt64(), 10), data.TrapTargetServers[i].Ip.ValueString(), strconv.FormatInt(data.TrapTargetServers[i].Port.ValueInt64(), 10)}
		keyValuesVariables := [...]string{data.TrapTargetServers[i].VpnIdVariable.ValueString(), data.TrapTargetServers[i].IpVariable.ValueString(), data.TrapTargetServers[i].PortVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "target").ForEach(
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
		data.TrapTargetServers[i].VpnId = types.Int64Null()
		data.TrapTargetServers[i].VpnIdVariable = types.StringNull()
		if t := r.Get("vpnId.optionType"); t.Exists() {
			va := r.Get("vpnId.value")
			if t.String() == "variable" {
				data.TrapTargetServers[i].VpnIdVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.TrapTargetServers[i].VpnId = types.Int64Value(va.Int())
			}
		}
		data.TrapTargetServers[i].Ip = types.StringNull()
		data.TrapTargetServers[i].IpVariable = types.StringNull()
		if t := r.Get("ip.optionType"); t.Exists() {
			va := r.Get("ip.value")
			if t.String() == "variable" {
				data.TrapTargetServers[i].IpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.TrapTargetServers[i].Ip = types.StringValue(va.String())
			}
		}
		data.TrapTargetServers[i].Port = types.Int64Null()
		data.TrapTargetServers[i].PortVariable = types.StringNull()
		if t := r.Get("port.optionType"); t.Exists() {
			va := r.Get("port.value")
			if t.String() == "variable" {
				data.TrapTargetServers[i].PortVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.TrapTargetServers[i].Port = types.Int64Value(va.Int())
			}
		}
		data.TrapTargetServers[i].UserLabel = types.StringNull()

		if t := r.Get("userLabel.optionType"); t.Exists() {
			va := r.Get("userLabel.value")
			if t.String() == "global" {
				data.TrapTargetServers[i].UserLabel = types.StringValue(va.String())
			}
		}
		data.TrapTargetServers[i].User = types.StringNull()
		data.TrapTargetServers[i].UserVariable = types.StringNull()
		if t := r.Get("user.optionType"); t.Exists() {
			va := r.Get("user.value")
			if t.String() == "variable" {
				data.TrapTargetServers[i].UserVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.TrapTargetServers[i].User = types.StringValue(va.String())
			}
		}
		data.TrapTargetServers[i].SourceInterface = types.StringNull()
		data.TrapTargetServers[i].SourceInterfaceVariable = types.StringNull()
		if t := r.Get("sourceInterface.optionType"); t.Exists() {
			va := r.Get("sourceInterface.value")
			if t.String() == "variable" {
				data.TrapTargetServers[i].SourceInterfaceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.TrapTargetServers[i].SourceInterface = types.StringValue(va.String())
			}
		}
	}
}

// End of section. //template:end updateFromBody
