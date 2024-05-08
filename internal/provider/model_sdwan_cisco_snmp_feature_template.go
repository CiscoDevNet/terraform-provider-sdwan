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
type CiscoSNMP struct {
	Id               types.String           `tfsdk:"id"`
	Version          types.Int64            `tfsdk:"version"`
	TemplateType     types.String           `tfsdk:"template_type"`
	Name             types.String           `tfsdk:"name"`
	Description      types.String           `tfsdk:"description"`
	DeviceTypes      types.Set              `tfsdk:"device_types"`
	Shutdown         types.Bool             `tfsdk:"shutdown"`
	ShutdownVariable types.String           `tfsdk:"shutdown_variable"`
	Contact          types.String           `tfsdk:"contact"`
	ContactVariable  types.String           `tfsdk:"contact_variable"`
	Location         types.String           `tfsdk:"location"`
	LocationVariable types.String           `tfsdk:"location_variable"`
	Views            []CiscoSNMPViews       `tfsdk:"views"`
	Communities      []CiscoSNMPCommunities `tfsdk:"communities"`
	Groups           []CiscoSNMPGroups      `tfsdk:"groups"`
	Users            []CiscoSNMPUsers       `tfsdk:"users"`
	TrapTargets      []CiscoSNMPTrapTargets `tfsdk:"trap_targets"`
}

type CiscoSNMPViews struct {
	Optional          types.Bool                        `tfsdk:"optional"`
	Name              types.String                      `tfsdk:"name"`
	ObjectIdentifiers []CiscoSNMPViewsObjectIdentifiers `tfsdk:"object_identifiers"`
}

type CiscoSNMPCommunities struct {
	Optional              types.Bool   `tfsdk:"optional"`
	Name                  types.String `tfsdk:"name"`
	View                  types.String `tfsdk:"view"`
	ViewVariable          types.String `tfsdk:"view_variable"`
	Authorization         types.String `tfsdk:"authorization"`
	AuthorizationVariable types.String `tfsdk:"authorization_variable"`
}

type CiscoSNMPGroups struct {
	Optional      types.Bool   `tfsdk:"optional"`
	Name          types.String `tfsdk:"name"`
	SecurityLevel types.String `tfsdk:"security_level"`
	View          types.String `tfsdk:"view"`
	ViewVariable  types.String `tfsdk:"view_variable"`
}

type CiscoSNMPUsers struct {
	Optional                       types.Bool   `tfsdk:"optional"`
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

type CiscoSNMPTrapTargets struct {
	Optional                types.Bool   `tfsdk:"optional"`
	VpnId                   types.Int64  `tfsdk:"vpn_id"`
	VpnIdVariable           types.String `tfsdk:"vpn_id_variable"`
	Ip                      types.String `tfsdk:"ip"`
	IpVariable              types.String `tfsdk:"ip_variable"`
	UdpPort                 types.Int64  `tfsdk:"udp_port"`
	UdpPortVariable         types.String `tfsdk:"udp_port_variable"`
	CommunityName           types.String `tfsdk:"community_name"`
	CommunityNameVariable   types.String `tfsdk:"community_name_variable"`
	User                    types.String `tfsdk:"user"`
	UserVariable            types.String `tfsdk:"user_variable"`
	SourceInterface         types.String `tfsdk:"source_interface"`
	SourceInterfaceVariable types.String `tfsdk:"source_interface_variable"`
}

type CiscoSNMPViewsObjectIdentifiers struct {
	Optional        types.Bool   `tfsdk:"optional"`
	Id              types.String `tfsdk:"id"`
	IdVariable      types.String `tfsdk:"id_variable"`
	Exclude         types.Bool   `tfsdk:"exclude"`
	ExcludeVariable types.String `tfsdk:"exclude_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoSNMP) getModel() string {
	return "cisco_snmp"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CiscoSNMP) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_snmp")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.ShutdownVariable.IsNull() {
		body, _ = sjson.Set(body, path+"shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"shutdown."+"vipVariableName", data.ShutdownVariable.ValueString())
	} else if data.Shutdown.IsNull() {
		body, _ = sjson.Set(body, path+"shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"shutdown."+"vipValue", strconv.FormatBool(data.Shutdown.ValueBool()))
	}

	if !data.ContactVariable.IsNull() {
		body, _ = sjson.Set(body, path+"contact."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"contact."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"contact."+"vipVariableName", data.ContactVariable.ValueString())
	} else if data.Contact.IsNull() {
		body, _ = sjson.Set(body, path+"contact."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"contact."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"contact."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"contact."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"contact."+"vipValue", data.Contact.ValueString())
	}

	if !data.LocationVariable.IsNull() {
		body, _ = sjson.Set(body, path+"location."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"location."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"location."+"vipVariableName", data.LocationVariable.ValueString())
	} else if data.Location.IsNull() {
		body, _ = sjson.Set(body, path+"location."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"location."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"location."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"location."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"location."+"vipValue", data.Location.ValueString())
	}
	if len(data.Views) > 0 {
		body, _ = sjson.Set(body, path+"view."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"view."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"view."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"view."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"view."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"view."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"view."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"view."+"vipValue", []interface{}{})
	}
	for _, item := range data.Views {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")
		if item.Name.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.Name.ValueString())
		}
		itemAttributes = append(itemAttributes, "oid")
		if len(item.ObjectIdentifiers) > 0 {
			itemBody, _ = sjson.Set(itemBody, "oid."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "oid."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "oid."+"vipPrimaryKey", []string{"id"})
			itemBody, _ = sjson.Set(itemBody, "oid."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "oid."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "oid."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "oid."+"vipPrimaryKey", []string{"id"})
			itemBody, _ = sjson.Set(itemBody, "oid."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.ObjectIdentifiers {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "id")

			if !childItem.IdVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "id."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "id."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "id."+"vipVariableName", childItem.IdVariable.ValueString())
			} else if childItem.Id.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "id."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "id."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "id."+"vipValue", childItem.Id.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "exclude")

			if !childItem.ExcludeVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "exclude."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "exclude."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "exclude."+"vipVariableName", childItem.ExcludeVariable.ValueString())
			} else if childItem.Exclude.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "exclude."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "exclude."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "exclude."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "exclude."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "exclude."+"vipValue", strconv.FormatBool(childItem.Exclude.ValueBool()))
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "oid."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"view."+"vipValue.-1", itemBody)
	}
	if len(data.Communities) > 0 {
		body, _ = sjson.Set(body, path+"community."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"community."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"community."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"community."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"community."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"community."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"community."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"community."+"vipValue", []interface{}{})
	}
	for _, item := range data.Communities {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")
		if item.Name.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.Name.ValueString())
		}
		itemAttributes = append(itemAttributes, "view")

		if !item.ViewVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "view."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "view."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "view."+"vipVariableName", item.ViewVariable.ValueString())
		} else if item.View.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "view."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "view."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "view."+"vipValue", item.View.ValueString())
		}
		itemAttributes = append(itemAttributes, "authorization")

		if !item.AuthorizationVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "authorization."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "authorization."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "authorization."+"vipVariableName", item.AuthorizationVariable.ValueString())
		} else if item.Authorization.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "authorization."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "authorization."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "authorization."+"vipValue", item.Authorization.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"community."+"vipValue.-1", itemBody)
	}
	if len(data.Groups) > 0 {
		body, _ = sjson.Set(body, path+"group."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"group."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"group."+"vipPrimaryKey", []string{"name", "security-level"})
		body, _ = sjson.Set(body, path+"group."+"vipValue", []interface{}{})
	} else {
	}
	for _, item := range data.Groups {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")
		if item.Name.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.Name.ValueString())
		}
		itemAttributes = append(itemAttributes, "security-level")
		if item.SecurityLevel.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "security-level."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "security-level."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "security-level."+"vipValue", item.SecurityLevel.ValueString())
		}
		itemAttributes = append(itemAttributes, "view")

		if !item.ViewVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "view."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "view."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "view."+"vipVariableName", item.ViewVariable.ValueString())
		} else if item.View.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "view."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "view."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "view."+"vipValue", item.View.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"group."+"vipValue.-1", itemBody)
	}
	if len(data.Users) > 0 {
		body, _ = sjson.Set(body, path+"user."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"user."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"user."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"user."+"vipValue", []interface{}{})
	} else {
	}
	for _, item := range data.Users {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")
		if item.Name.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.Name.ValueString())
		}
		itemAttributes = append(itemAttributes, "auth")

		if !item.AuthenticationProtocolVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "auth."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "auth."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "auth."+"vipVariableName", item.AuthenticationProtocolVariable.ValueString())
		} else if item.AuthenticationProtocol.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "auth."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "auth."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "auth."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "auth."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "auth."+"vipValue", item.AuthenticationProtocol.ValueString())
		}
		itemAttributes = append(itemAttributes, "auth-password")

		if !item.AuthenticationPasswordVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "auth-password."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "auth-password."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "auth-password."+"vipVariableName", item.AuthenticationPasswordVariable.ValueString())
		} else if item.AuthenticationPassword.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "auth-password."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "auth-password."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "auth-password."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "auth-password."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "auth-password."+"vipValue", item.AuthenticationPassword.ValueString())
		}
		itemAttributes = append(itemAttributes, "priv")

		if !item.PrivacyProtocolVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priv."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priv."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "priv."+"vipVariableName", item.PrivacyProtocolVariable.ValueString())
		} else if item.PrivacyProtocol.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priv."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priv."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "priv."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priv."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "priv."+"vipValue", item.PrivacyProtocol.ValueString())
		}
		itemAttributes = append(itemAttributes, "priv-password")

		if !item.PrivacyPasswordVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priv-password."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priv-password."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "priv-password."+"vipVariableName", item.PrivacyPasswordVariable.ValueString())
		} else if item.PrivacyPassword.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "priv-password."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priv-password."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "priv-password."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "priv-password."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "priv-password."+"vipValue", item.PrivacyPassword.ValueString())
		}
		itemAttributes = append(itemAttributes, "group")

		if !item.GroupVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "group."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "group."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "group."+"vipVariableName", item.GroupVariable.ValueString())
		} else if item.Group.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "group."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "group."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "group."+"vipValue", item.Group.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"user."+"vipValue.-1", itemBody)
	}
	if len(data.TrapTargets) > 0 {
		body, _ = sjson.Set(body, path+"trap.target."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"trap.target."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"trap.target."+"vipPrimaryKey", []string{"vpn-id", "ip", "port"})
		body, _ = sjson.Set(body, path+"trap.target."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"trap.target."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"trap.target."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"trap.target."+"vipPrimaryKey", []string{"vpn-id", "ip", "port"})
		body, _ = sjson.Set(body, path+"trap.target."+"vipValue", []interface{}{})
	}
	for _, item := range data.TrapTargets {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "vpn-id")

		if !item.VpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipVariableName", item.VpnIdVariable.ValueString())
		} else if item.VpnId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipValue", item.VpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "ip")

		if !item.IpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipVariableName", item.IpVariable.ValueString())
		} else if item.Ip.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipValue", item.Ip.ValueString())
		}
		itemAttributes = append(itemAttributes, "port")

		if !item.UdpPortVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "port."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "port."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "port."+"vipVariableName", item.UdpPortVariable.ValueString())
		} else if item.UdpPort.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "port."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "port."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "port."+"vipValue", item.UdpPort.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "community-name")

		if !item.CommunityNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "community-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "community-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "community-name."+"vipVariableName", item.CommunityNameVariable.ValueString())
		} else if item.CommunityName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "community-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "community-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "community-name."+"vipValue", item.CommunityName.ValueString())
		}
		itemAttributes = append(itemAttributes, "user")

		if !item.UserVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "user."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "user."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "user."+"vipVariableName", item.UserVariable.ValueString())
		} else if item.User.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "user."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "user."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "user."+"vipValue", item.User.ValueString())
		}
		itemAttributes = append(itemAttributes, "source-interface")

		if !item.SourceInterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipVariableName", item.SourceInterfaceVariable.ValueString())
		} else if item.SourceInterface.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipValue", item.SourceInterface.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"trap.target."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoSNMP) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "shutdown.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Shutdown = types.BoolNull()

			v := res.Get(path + "shutdown.vipVariableName")
			data.ShutdownVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Shutdown = types.BoolNull()
			data.ShutdownVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "shutdown.vipValue")
			data.Shutdown = types.BoolValue(v.Bool())
			data.ShutdownVariable = types.StringNull()
		}
	} else {
		data.Shutdown = types.BoolNull()
		data.ShutdownVariable = types.StringNull()
	}
	if value := res.Get(path + "contact.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Contact = types.StringNull()

			v := res.Get(path + "contact.vipVariableName")
			data.ContactVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Contact = types.StringNull()
			data.ContactVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "contact.vipValue")
			data.Contact = types.StringValue(v.String())
			data.ContactVariable = types.StringNull()
		}
	} else {
		data.Contact = types.StringNull()
		data.ContactVariable = types.StringNull()
	}
	if value := res.Get(path + "location.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Location = types.StringNull()

			v := res.Get(path + "location.vipVariableName")
			data.LocationVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Location = types.StringNull()
			data.LocationVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "location.vipValue")
			data.Location = types.StringValue(v.String())
			data.LocationVariable = types.StringNull()
		}
	} else {
		data.Location = types.StringNull()
		data.LocationVariable = types.StringNull()
	}
	if value := res.Get(path + "view.vipValue"); len(value.Array()) > 0 {
		data.Views = make([]CiscoSNMPViews, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoSNMPViews{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Name = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Name = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.Name = types.StringValue(cv.String())

				}
			} else {
				item.Name = types.StringNull()

			}
			if cValue := v.Get("oid.vipValue"); len(cValue.Array()) > 0 {
				item.ObjectIdentifiers = make([]CiscoSNMPViewsObjectIdentifiers, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoSNMPViewsObjectIdentifiers{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("id.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Id = types.StringNull()

							ccv := cv.Get("id.vipVariableName")
							cItem.IdVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Id = types.StringNull()
							cItem.IdVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("id.vipValue")
							cItem.Id = types.StringValue(ccv.String())
							cItem.IdVariable = types.StringNull()
						}
					} else {
						cItem.Id = types.StringNull()
						cItem.IdVariable = types.StringNull()
					}
					if ccValue := cv.Get("exclude.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Exclude = types.BoolNull()

							ccv := cv.Get("exclude.vipVariableName")
							cItem.ExcludeVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Exclude = types.BoolNull()
							cItem.ExcludeVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("exclude.vipValue")
							cItem.Exclude = types.BoolValue(ccv.Bool())
							cItem.ExcludeVariable = types.StringNull()
						}
					} else {
						cItem.Exclude = types.BoolNull()
						cItem.ExcludeVariable = types.StringNull()
					}
					item.ObjectIdentifiers = append(item.ObjectIdentifiers, cItem)
					return true
				})
			} else {
				if len(item.ObjectIdentifiers) > 0 {
					item.ObjectIdentifiers = []CiscoSNMPViewsObjectIdentifiers{}
				}
			}
			data.Views = append(data.Views, item)
			return true
		})
	} else {
		if len(data.Views) > 0 {
			data.Views = []CiscoSNMPViews{}
		}
	}
	if value := res.Get(path + "community.vipValue"); len(value.Array()) > 0 {
		data.Communities = make([]CiscoSNMPCommunities, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoSNMPCommunities{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Name = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Name = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.Name = types.StringValue(cv.String())

				}
			} else {
				item.Name = types.StringNull()

			}
			if cValue := v.Get("view.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.View = types.StringNull()

					cv := v.Get("view.vipVariableName")
					item.ViewVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.View = types.StringNull()
					item.ViewVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("view.vipValue")
					item.View = types.StringValue(cv.String())
					item.ViewVariable = types.StringNull()
				}
			} else {
				item.View = types.StringNull()
				item.ViewVariable = types.StringNull()
			}
			if cValue := v.Get("authorization.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Authorization = types.StringNull()

					cv := v.Get("authorization.vipVariableName")
					item.AuthorizationVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Authorization = types.StringNull()
					item.AuthorizationVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("authorization.vipValue")
					item.Authorization = types.StringValue(cv.String())
					item.AuthorizationVariable = types.StringNull()
				}
			} else {
				item.Authorization = types.StringNull()
				item.AuthorizationVariable = types.StringNull()
			}
			data.Communities = append(data.Communities, item)
			return true
		})
	} else {
		if len(data.Communities) > 0 {
			data.Communities = []CiscoSNMPCommunities{}
		}
	}
	if value := res.Get(path + "group.vipValue"); len(value.Array()) > 0 {
		data.Groups = make([]CiscoSNMPGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoSNMPGroups{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Name = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Name = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.Name = types.StringValue(cv.String())

				}
			} else {
				item.Name = types.StringNull()

			}
			if cValue := v.Get("security-level.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SecurityLevel = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.SecurityLevel = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("security-level.vipValue")
					item.SecurityLevel = types.StringValue(cv.String())

				}
			} else {
				item.SecurityLevel = types.StringNull()

			}
			if cValue := v.Get("view.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.View = types.StringNull()

					cv := v.Get("view.vipVariableName")
					item.ViewVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.View = types.StringNull()
					item.ViewVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("view.vipValue")
					item.View = types.StringValue(cv.String())
					item.ViewVariable = types.StringNull()
				}
			} else {
				item.View = types.StringNull()
				item.ViewVariable = types.StringNull()
			}
			data.Groups = append(data.Groups, item)
			return true
		})
	} else {
		if len(data.Groups) > 0 {
			data.Groups = []CiscoSNMPGroups{}
		}
	}
	if value := res.Get(path + "user.vipValue"); len(value.Array()) > 0 {
		data.Users = make([]CiscoSNMPUsers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoSNMPUsers{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Name = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Name = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.Name = types.StringValue(cv.String())

				}
			} else {
				item.Name = types.StringNull()

			}
			if cValue := v.Get("auth.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AuthenticationProtocol = types.StringNull()

					cv := v.Get("auth.vipVariableName")
					item.AuthenticationProtocolVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AuthenticationProtocol = types.StringNull()
					item.AuthenticationProtocolVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("auth.vipValue")
					item.AuthenticationProtocol = types.StringValue(cv.String())
					item.AuthenticationProtocolVariable = types.StringNull()
				}
			} else {
				item.AuthenticationProtocol = types.StringNull()
				item.AuthenticationProtocolVariable = types.StringNull()
			}
			if cValue := v.Get("auth-password.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AuthenticationPassword = types.StringNull()

					cv := v.Get("auth-password.vipVariableName")
					item.AuthenticationPasswordVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AuthenticationPassword = types.StringNull()
					item.AuthenticationPasswordVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("auth-password.vipValue")
					item.AuthenticationPassword = types.StringValue(cv.String())
					item.AuthenticationPasswordVariable = types.StringNull()
				}
			} else {
				item.AuthenticationPassword = types.StringNull()
				item.AuthenticationPasswordVariable = types.StringNull()
			}
			if cValue := v.Get("priv.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PrivacyProtocol = types.StringNull()

					cv := v.Get("priv.vipVariableName")
					item.PrivacyProtocolVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.PrivacyProtocol = types.StringNull()
					item.PrivacyProtocolVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("priv.vipValue")
					item.PrivacyProtocol = types.StringValue(cv.String())
					item.PrivacyProtocolVariable = types.StringNull()
				}
			} else {
				item.PrivacyProtocol = types.StringNull()
				item.PrivacyProtocolVariable = types.StringNull()
			}
			if cValue := v.Get("priv-password.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PrivacyPassword = types.StringNull()

					cv := v.Get("priv-password.vipVariableName")
					item.PrivacyPasswordVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.PrivacyPassword = types.StringNull()
					item.PrivacyPasswordVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("priv-password.vipValue")
					item.PrivacyPassword = types.StringValue(cv.String())
					item.PrivacyPasswordVariable = types.StringNull()
				}
			} else {
				item.PrivacyPassword = types.StringNull()
				item.PrivacyPasswordVariable = types.StringNull()
			}
			if cValue := v.Get("group.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Group = types.StringNull()

					cv := v.Get("group.vipVariableName")
					item.GroupVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Group = types.StringNull()
					item.GroupVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("group.vipValue")
					item.Group = types.StringValue(cv.String())
					item.GroupVariable = types.StringNull()
				}
			} else {
				item.Group = types.StringNull()
				item.GroupVariable = types.StringNull()
			}
			data.Users = append(data.Users, item)
			return true
		})
	} else {
		if len(data.Users) > 0 {
			data.Users = []CiscoSNMPUsers{}
		}
	}
	if value := res.Get(path + "trap.target.vipValue"); len(value.Array()) > 0 {
		data.TrapTargets = make([]CiscoSNMPTrapTargets, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoSNMPTrapTargets{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("vpn-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

					cv := v.Get("vpn-id.vipVariableName")
					item.VpnIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()
					item.VpnIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("vpn-id.vipValue")
					item.VpnId = types.Int64Value(cv.Int())
					item.VpnIdVariable = types.StringNull()
				}
			} else {
				item.VpnId = types.Int64Null()
				item.VpnIdVariable = types.StringNull()
			}
			if cValue := v.Get("ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Ip = types.StringNull()

					cv := v.Get("ip.vipVariableName")
					item.IpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Ip = types.StringNull()
					item.IpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ip.vipValue")
					item.Ip = types.StringValue(cv.String())
					item.IpVariable = types.StringNull()
				}
			} else {
				item.Ip = types.StringNull()
				item.IpVariable = types.StringNull()
			}
			if cValue := v.Get("port.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.UdpPort = types.Int64Null()

					cv := v.Get("port.vipVariableName")
					item.UdpPortVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.UdpPort = types.Int64Null()
					item.UdpPortVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("port.vipValue")
					item.UdpPort = types.Int64Value(cv.Int())
					item.UdpPortVariable = types.StringNull()
				}
			} else {
				item.UdpPort = types.Int64Null()
				item.UdpPortVariable = types.StringNull()
			}
			if cValue := v.Get("community-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.CommunityName = types.StringNull()

					cv := v.Get("community-name.vipVariableName")
					item.CommunityNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.CommunityName = types.StringNull()
					item.CommunityNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("community-name.vipValue")
					item.CommunityName = types.StringValue(cv.String())
					item.CommunityNameVariable = types.StringNull()
				}
			} else {
				item.CommunityName = types.StringNull()
				item.CommunityNameVariable = types.StringNull()
			}
			if cValue := v.Get("user.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.User = types.StringNull()

					cv := v.Get("user.vipVariableName")
					item.UserVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.User = types.StringNull()
					item.UserVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("user.vipValue")
					item.User = types.StringValue(cv.String())
					item.UserVariable = types.StringNull()
				}
			} else {
				item.User = types.StringNull()
				item.UserVariable = types.StringNull()
			}
			if cValue := v.Get("source-interface.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceInterface = types.StringNull()

					cv := v.Get("source-interface.vipVariableName")
					item.SourceInterfaceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceInterface = types.StringNull()
					item.SourceInterfaceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-interface.vipValue")
					item.SourceInterface = types.StringValue(cv.String())
					item.SourceInterfaceVariable = types.StringNull()
				}
			} else {
				item.SourceInterface = types.StringNull()
				item.SourceInterfaceVariable = types.StringNull()
			}
			data.TrapTargets = append(data.TrapTargets, item)
			return true
		})
	} else {
		if len(data.TrapTargets) > 0 {
			data.TrapTargets = []CiscoSNMPTrapTargets{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoSNMP) hasChanges(ctx context.Context, state *CiscoSNMP) bool {
	hasChanges := false
	if !data.Shutdown.Equal(state.Shutdown) {
		hasChanges = true
	}
	if !data.Contact.Equal(state.Contact) {
		hasChanges = true
	}
	if !data.Location.Equal(state.Location) {
		hasChanges = true
	}
	if len(data.Views) != len(state.Views) {
		hasChanges = true
	} else {
		for i := range data.Views {
			if !data.Views[i].Name.Equal(state.Views[i].Name) {
				hasChanges = true
			}
			if len(data.Views[i].ObjectIdentifiers) != len(state.Views[i].ObjectIdentifiers) {
				hasChanges = true
			} else {
				for ii := range data.Views[i].ObjectIdentifiers {
					if !data.Views[i].ObjectIdentifiers[ii].Id.Equal(state.Views[i].ObjectIdentifiers[ii].Id) {
						hasChanges = true
					}
					if !data.Views[i].ObjectIdentifiers[ii].Exclude.Equal(state.Views[i].ObjectIdentifiers[ii].Exclude) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.Communities) != len(state.Communities) {
		hasChanges = true
	} else {
		for i := range data.Communities {
			if !data.Communities[i].Name.Equal(state.Communities[i].Name) {
				hasChanges = true
			}
			if !data.Communities[i].View.Equal(state.Communities[i].View) {
				hasChanges = true
			}
			if !data.Communities[i].Authorization.Equal(state.Communities[i].Authorization) {
				hasChanges = true
			}
		}
	}
	if len(data.Groups) != len(state.Groups) {
		hasChanges = true
	} else {
		for i := range data.Groups {
			if !data.Groups[i].Name.Equal(state.Groups[i].Name) {
				hasChanges = true
			}
			if !data.Groups[i].SecurityLevel.Equal(state.Groups[i].SecurityLevel) {
				hasChanges = true
			}
			if !data.Groups[i].View.Equal(state.Groups[i].View) {
				hasChanges = true
			}
		}
	}
	if len(data.Users) != len(state.Users) {
		hasChanges = true
	} else {
		for i := range data.Users {
			if !data.Users[i].Name.Equal(state.Users[i].Name) {
				hasChanges = true
			}
			if !data.Users[i].AuthenticationProtocol.Equal(state.Users[i].AuthenticationProtocol) {
				hasChanges = true
			}
			if !data.Users[i].AuthenticationPassword.Equal(state.Users[i].AuthenticationPassword) {
				hasChanges = true
			}
			if !data.Users[i].PrivacyProtocol.Equal(state.Users[i].PrivacyProtocol) {
				hasChanges = true
			}
			if !data.Users[i].PrivacyPassword.Equal(state.Users[i].PrivacyPassword) {
				hasChanges = true
			}
			if !data.Users[i].Group.Equal(state.Users[i].Group) {
				hasChanges = true
			}
		}
	}
	if len(data.TrapTargets) != len(state.TrapTargets) {
		hasChanges = true
	} else {
		for i := range data.TrapTargets {
			if !data.TrapTargets[i].VpnId.Equal(state.TrapTargets[i].VpnId) {
				hasChanges = true
			}
			if !data.TrapTargets[i].Ip.Equal(state.TrapTargets[i].Ip) {
				hasChanges = true
			}
			if !data.TrapTargets[i].UdpPort.Equal(state.TrapTargets[i].UdpPort) {
				hasChanges = true
			}
			if !data.TrapTargets[i].CommunityName.Equal(state.TrapTargets[i].CommunityName) {
				hasChanges = true
			}
			if !data.TrapTargets[i].User.Equal(state.TrapTargets[i].User) {
				hasChanges = true
			}
			if !data.TrapTargets[i].SourceInterface.Equal(state.TrapTargets[i].SourceInterface) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
