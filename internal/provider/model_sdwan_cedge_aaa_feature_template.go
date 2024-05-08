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
type CEdgeAAA struct {
	Id                                            types.String                 `tfsdk:"id"`
	Version                                       types.Int64                  `tfsdk:"version"`
	TemplateType                                  types.String                 `tfsdk:"template_type"`
	Name                                          types.String                 `tfsdk:"name"`
	Description                                   types.String                 `tfsdk:"description"`
	DeviceTypes                                   types.Set                    `tfsdk:"device_types"`
	Dot1xAuthentication                           types.Bool                   `tfsdk:"dot1x_authentication"`
	Dot1xAuthenticationVariable                   types.String                 `tfsdk:"dot1x_authentication_variable"`
	Dot1xAccounting                               types.Bool                   `tfsdk:"dot1x_accounting"`
	Dot1xAccountingVariable                       types.String                 `tfsdk:"dot1x_accounting_variable"`
	ServerGroupsPriorityOrder                     types.String                 `tfsdk:"server_groups_priority_order"`
	Users                                         []CEdgeAAAUsers              `tfsdk:"users"`
	RadiusServerGroups                            []CEdgeAAARadiusServerGroups `tfsdk:"radius_server_groups"`
	RadiusClients                                 []CEdgeAAARadiusClients      `tfsdk:"radius_clients"`
	RadiusDynamicAuthorServerKey                  types.String                 `tfsdk:"radius_dynamic_author_server_key"`
	RadiusDynamicAuthorServerKeyVariable          types.String                 `tfsdk:"radius_dynamic_author_server_key_variable"`
	RadiusDynamicAuthorDomainStripping            types.String                 `tfsdk:"radius_dynamic_author_domain_stripping"`
	RadiusDynamicAuthorDomainStrippingVariable    types.String                 `tfsdk:"radius_dynamic_author_domain_stripping_variable"`
	RadiusDynamicAuthorAuthenticationType         types.String                 `tfsdk:"radius_dynamic_author_authentication_type"`
	RadiusDynamicAuthorAuthenticationTypeVariable types.String                 `tfsdk:"radius_dynamic_author_authentication_type_variable"`
	RadiusDynamicAuthorPort                       types.Int64                  `tfsdk:"radius_dynamic_author_port"`
	RadiusDynamicAuthorPortVariable               types.String                 `tfsdk:"radius_dynamic_author_port_variable"`
	RadiusTrustsecCtsAuthorizationList            types.String                 `tfsdk:"radius_trustsec_cts_authorization_list"`
	RadiusTrustsecCtsAuthorizationListVariable    types.String                 `tfsdk:"radius_trustsec_cts_authorization_list_variable"`
	RadiusTrustsecGroup                           types.String                 `tfsdk:"radius_trustsec_group"`
	TacacsServerGroups                            []CEdgeAAATacacsServerGroups `tfsdk:"tacacs_server_groups"`
	AccountingRules                               []CEdgeAAAAccountingRules    `tfsdk:"accounting_rules"`
	AuthorizationConsole                          types.Bool                   `tfsdk:"authorization_console"`
	AuthorizationConsoleVariable                  types.String                 `tfsdk:"authorization_console_variable"`
	AuthorizationConfigCommands                   types.Bool                   `tfsdk:"authorization_config_commands"`
	AuthorizationConfigCommandsVariable           types.String                 `tfsdk:"authorization_config_commands_variable"`
	AuthorizationRules                            []CEdgeAAAAuthorizationRules `tfsdk:"authorization_rules"`
}

type CEdgeAAAUsers struct {
	Optional               types.Bool                `tfsdk:"optional"`
	Name                   types.String              `tfsdk:"name"`
	NameVariable           types.String              `tfsdk:"name_variable"`
	Password               types.String              `tfsdk:"password"`
	Secret                 types.String              `tfsdk:"secret"`
	PrivilegeLevel         types.String              `tfsdk:"privilege_level"`
	PrivilegeLevelVariable types.String              `tfsdk:"privilege_level_variable"`
	SshPubkeys             []CEdgeAAAUsersSshPubkeys `tfsdk:"ssh_pubkeys"`
}

type CEdgeAAARadiusServerGroups struct {
	Optional                types.Bool                          `tfsdk:"optional"`
	GroupName               types.String                        `tfsdk:"group_name"`
	VpnId                   types.Int64                         `tfsdk:"vpn_id"`
	SourceInterface         types.String                        `tfsdk:"source_interface"`
	SourceInterfaceVariable types.String                        `tfsdk:"source_interface_variable"`
	Servers                 []CEdgeAAARadiusServerGroupsServers `tfsdk:"servers"`
}

type CEdgeAAARadiusClients struct {
	Optional          types.Bool                               `tfsdk:"optional"`
	ClientIp          types.String                             `tfsdk:"client_ip"`
	ClientIpVariable  types.String                             `tfsdk:"client_ip_variable"`
	VpnConfigurations []CEdgeAAARadiusClientsVpnConfigurations `tfsdk:"vpn_configurations"`
}

type CEdgeAAATacacsServerGroups struct {
	Optional                types.Bool                          `tfsdk:"optional"`
	GroupName               types.String                        `tfsdk:"group_name"`
	VpnId                   types.Int64                         `tfsdk:"vpn_id"`
	SourceInterface         types.String                        `tfsdk:"source_interface"`
	SourceInterfaceVariable types.String                        `tfsdk:"source_interface_variable"`
	Servers                 []CEdgeAAATacacsServerGroupsServers `tfsdk:"servers"`
}

type CEdgeAAAAccountingRules struct {
	Optional          types.Bool   `tfsdk:"optional"`
	Name              types.String `tfsdk:"name"`
	Method            types.String `tfsdk:"method"`
	PrivilegeLevel    types.String `tfsdk:"privilege_level"`
	StartStop         types.Bool   `tfsdk:"start_stop"`
	StartStopVariable types.String `tfsdk:"start_stop_variable"`
	Groups            types.String `tfsdk:"groups"`
}

type CEdgeAAAAuthorizationRules struct {
	Optional       types.Bool   `tfsdk:"optional"`
	Name           types.String `tfsdk:"name"`
	Method         types.String `tfsdk:"method"`
	PrivilegeLevel types.String `tfsdk:"privilege_level"`
	Groups         types.String `tfsdk:"groups"`
	Authenticated  types.Bool   `tfsdk:"authenticated"`
}

type CEdgeAAAUsersSshPubkeys struct {
	Optional        types.Bool   `tfsdk:"optional"`
	KeyString       types.String `tfsdk:"key_string"`
	KeyType         types.String `tfsdk:"key_type"`
	KeyTypeVariable types.String `tfsdk:"key_type_variable"`
}

type CEdgeAAARadiusServerGroupsServers struct {
	Optional                   types.Bool   `tfsdk:"optional"`
	Address                    types.String `tfsdk:"address"`
	AuthenticationPort         types.Int64  `tfsdk:"authentication_port"`
	AuthenticationPortVariable types.String `tfsdk:"authentication_port_variable"`
	AccountingPort             types.Int64  `tfsdk:"accounting_port"`
	AccountingPortVariable     types.String `tfsdk:"accounting_port_variable"`
	Timeout                    types.Int64  `tfsdk:"timeout"`
	TimeoutVariable            types.String `tfsdk:"timeout_variable"`
	Retransmit                 types.Int64  `tfsdk:"retransmit"`
	RetransmitVariable         types.String `tfsdk:"retransmit_variable"`
	Key                        types.String `tfsdk:"key"`
	SecretKey                  types.String `tfsdk:"secret_key"`
	SecretKeyVariable          types.String `tfsdk:"secret_key_variable"`
	EncryptionType             types.String `tfsdk:"encryption_type"`
	KeyType                    types.String `tfsdk:"key_type"`
	KeyTypeVariable            types.String `tfsdk:"key_type_variable"`
}

type CEdgeAAARadiusClientsVpnConfigurations struct {
	Optional          types.Bool   `tfsdk:"optional"`
	VpnId             types.Int64  `tfsdk:"vpn_id"`
	VpnIdVariable     types.String `tfsdk:"vpn_id_variable"`
	ServerKey         types.String `tfsdk:"server_key"`
	ServerKeyVariable types.String `tfsdk:"server_key_variable"`
}

type CEdgeAAATacacsServerGroupsServers struct {
	Optional          types.Bool   `tfsdk:"optional"`
	Address           types.String `tfsdk:"address"`
	Port              types.Int64  `tfsdk:"port"`
	PortVariable      types.String `tfsdk:"port_variable"`
	Timeout           types.Int64  `tfsdk:"timeout"`
	TimeoutVariable   types.String `tfsdk:"timeout_variable"`
	Key               types.String `tfsdk:"key"`
	SecretKey         types.String `tfsdk:"secret_key"`
	SecretKeyVariable types.String `tfsdk:"secret_key_variable"`
	EncryptionType    types.String `tfsdk:"encryption_type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CEdgeAAA) getModel() string {
	return "cedge_aaa"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CEdgeAAA) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cedge_aaa")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.Dot1xAuthenticationVariable.IsNull() {
		body, _ = sjson.Set(body, path+"authentication.dot1x.default.authentication_group."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"authentication.dot1x.default.authentication_group."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"authentication.dot1x.default.authentication_group."+"vipVariableName", data.Dot1xAuthenticationVariable.ValueString())
	} else if data.Dot1xAuthentication.IsNull() {
		body, _ = sjson.Set(body, path+"authentication.dot1x.default.authentication_group."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"authentication.dot1x.default.authentication_group."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"authentication.dot1x.default.authentication_group."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"authentication.dot1x.default.authentication_group."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"authentication.dot1x.default.authentication_group."+"vipValue", strconv.FormatBool(data.Dot1xAuthentication.ValueBool()))
	}

	if !data.Dot1xAccountingVariable.IsNull() {
		body, _ = sjson.Set(body, path+"accounting.dot1x.default.start-stop.accounting_group."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"accounting.dot1x.default.start-stop.accounting_group."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"accounting.dot1x.default.start-stop.accounting_group."+"vipVariableName", data.Dot1xAccountingVariable.ValueString())
	} else if data.Dot1xAccounting.IsNull() {
		body, _ = sjson.Set(body, path+"accounting.dot1x.default.start-stop.accounting_group."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"accounting.dot1x.default.start-stop.accounting_group."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"accounting.dot1x.default.start-stop.accounting_group."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"accounting.dot1x.default.start-stop.accounting_group."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"accounting.dot1x.default.start-stop.accounting_group."+"vipValue", strconv.FormatBool(data.Dot1xAccounting.ValueBool()))
	}
	if data.ServerGroupsPriorityOrder.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"server-auth-order."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"server-auth-order."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"server-auth-order."+"vipValue", data.ServerGroupsPriorityOrder.ValueString())
	}
	if len(data.Users) > 0 {
		body, _ = sjson.Set(body, path+"user."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"user."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"user."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"user."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"user."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"user."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"user."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"user."+"vipValue", []interface{}{})
	}
	for _, item := range data.Users {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")

		if !item.NameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipVariableName", item.NameVariable.ValueString())
		} else if item.Name.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.Name.ValueString())
		}
		itemAttributes = append(itemAttributes, "password")
		if item.Password.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "password."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "password."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "password."+"vipValue", item.Password.ValueString())
		}
		itemAttributes = append(itemAttributes, "secret")
		if item.Secret.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "secret."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "secret."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "secret."+"vipValue", item.Secret.ValueString())
		}
		itemAttributes = append(itemAttributes, "privilege")

		if !item.PrivilegeLevelVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "privilege."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "privilege."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "privilege."+"vipVariableName", item.PrivilegeLevelVariable.ValueString())
		} else if item.PrivilegeLevel.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "privilege."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "privilege."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "privilege."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "privilege."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "privilege."+"vipValue", item.PrivilegeLevel.ValueString())
		}
		itemAttributes = append(itemAttributes, "pubkey-chain")
		if len(item.SshPubkeys) > 0 {
			itemBody, _ = sjson.Set(itemBody, "pubkey-chain."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "pubkey-chain."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "pubkey-chain."+"vipPrimaryKey", []string{"key-string"})
			itemBody, _ = sjson.Set(itemBody, "pubkey-chain."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "pubkey-chain."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "pubkey-chain."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "pubkey-chain."+"vipPrimaryKey", []string{"key-string"})
			itemBody, _ = sjson.Set(itemBody, "pubkey-chain."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.SshPubkeys {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "key-string")
			if childItem.KeyString.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "key-string."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-string."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-string."+"vipValue", childItem.KeyString.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "key-type")

			if !childItem.KeyTypeVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "key-type."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-type."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-type."+"vipVariableName", childItem.KeyTypeVariable.ValueString())
			} else if childItem.KeyType.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "key-type."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-type."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "key-type."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-type."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-type."+"vipValue", childItem.KeyType.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "pubkey-chain."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"user."+"vipValue.-1", itemBody)
	}
	if len(data.RadiusServerGroups) > 0 {
		body, _ = sjson.Set(body, path+"radius."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"radius."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"radius."+"vipPrimaryKey", []string{"group-name"})
		body, _ = sjson.Set(body, path+"radius."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"radius."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"radius."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"radius."+"vipPrimaryKey", []string{"group-name"})
		body, _ = sjson.Set(body, path+"radius."+"vipValue", []interface{}{})
	}
	for _, item := range data.RadiusServerGroups {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "group-name")
		if item.GroupName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "group-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "group-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "group-name."+"vipValue", item.GroupName.ValueString())
		}
		itemAttributes = append(itemAttributes, "vpn")
		if item.VpnId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", item.VpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "source-interface")

		if !item.SourceInterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipVariableName", item.SourceInterfaceVariable.ValueString())
		} else if item.SourceInterface.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipValue", item.SourceInterface.ValueString())
		}
		itemAttributes = append(itemAttributes, "server")
		if len(item.Servers) > 0 {
			itemBody, _ = sjson.Set(itemBody, "server."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "server."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "server."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "server."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "server."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "server."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "server."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "server."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Servers {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "address")
			if childItem.Address.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipValue", childItem.Address.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "auth-port")

			if !childItem.AuthenticationPortVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "auth-port."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "auth-port."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "auth-port."+"vipVariableName", childItem.AuthenticationPortVariable.ValueString())
			} else if childItem.AuthenticationPort.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "auth-port."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "auth-port."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "auth-port."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "auth-port."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "auth-port."+"vipValue", childItem.AuthenticationPort.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "acct-port")

			if !childItem.AccountingPortVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "acct-port."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "acct-port."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "acct-port."+"vipVariableName", childItem.AccountingPortVariable.ValueString())
			} else if childItem.AccountingPort.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "acct-port."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "acct-port."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "acct-port."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "acct-port."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "acct-port."+"vipValue", childItem.AccountingPort.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "timeout")

			if !childItem.TimeoutVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "timeout."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "timeout."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "timeout."+"vipVariableName", childItem.TimeoutVariable.ValueString())
			} else if childItem.Timeout.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "timeout."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "timeout."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "timeout."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "timeout."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "timeout."+"vipValue", childItem.Timeout.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "retransmit")

			if !childItem.RetransmitVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit."+"vipVariableName", childItem.RetransmitVariable.ValueString())
			} else if childItem.Retransmit.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "retransmit."+"vipValue", childItem.Retransmit.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "key")
			if childItem.Key.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "key."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "key."+"vipValue", childItem.Key.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "secret-key")

			if !childItem.SecretKeyVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "secret-key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "secret-key."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "secret-key."+"vipVariableName", childItem.SecretKeyVariable.ValueString())
			} else if childItem.SecretKey.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "secret-key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "secret-key."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "secret-key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "secret-key."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "secret-key."+"vipValue", childItem.SecretKey.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "key-enum")
			if childItem.EncryptionType.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "key-enum."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-enum."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "key-enum."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-enum."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-enum."+"vipValue", childItem.EncryptionType.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "key-type")

			if !childItem.KeyTypeVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "key-type."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-type."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-type."+"vipVariableName", childItem.KeyTypeVariable.ValueString())
			} else if childItem.KeyType.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "key-type."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-type."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "key-type."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-type."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-type."+"vipValue", childItem.KeyType.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "server."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"radius."+"vipValue.-1", itemBody)
	}
	if len(data.RadiusClients) > 0 {
		body, _ = sjson.Set(body, path+"radius-dynamic-author.radius-client."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.radius-client."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.radius-client."+"vipPrimaryKey", []string{"ip"})
		body, _ = sjson.Set(body, path+"radius-dynamic-author.radius-client."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"radius-dynamic-author.radius-client."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.radius-client."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.radius-client."+"vipPrimaryKey", []string{"ip"})
		body, _ = sjson.Set(body, path+"radius-dynamic-author.radius-client."+"vipValue", []interface{}{})
	}
	for _, item := range data.RadiusClients {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "ip")

		if !item.ClientIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipVariableName", item.ClientIpVariable.ValueString())
		} else if item.ClientIp.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipValue", item.ClientIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "vpn")
		if len(item.VpnConfigurations) > 0 {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipPrimaryKey", []string{"name"})
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipPrimaryKey", []string{"name"})
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.VpnConfigurations {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "name")

			if !childItem.VpnIdVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipVariableName", childItem.VpnIdVariable.ValueString())
			} else if childItem.VpnId.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "name."+"vipValue", childItem.VpnId.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "server-key")

			if !childItem.ServerKeyVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "server-key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "server-key."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "server-key."+"vipVariableName", childItem.ServerKeyVariable.ValueString())
			} else if childItem.ServerKey.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "server-key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "server-key."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "server-key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "server-key."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "server-key."+"vipValue", childItem.ServerKey.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "vpn."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"radius-dynamic-author.radius-client."+"vipValue.-1", itemBody)
	}

	if !data.RadiusDynamicAuthorServerKeyVariable.IsNull() {
		body, _ = sjson.Set(body, path+"radius-dynamic-author.rda-server-key."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.rda-server-key."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.rda-server-key."+"vipVariableName", data.RadiusDynamicAuthorServerKeyVariable.ValueString())
	} else if data.RadiusDynamicAuthorServerKey.IsNull() {
		body, _ = sjson.Set(body, path+"radius-dynamic-author.rda-server-key."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.rda-server-key."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"radius-dynamic-author.rda-server-key."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.rda-server-key."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.rda-server-key."+"vipValue", data.RadiusDynamicAuthorServerKey.ValueString())
	}

	if !data.RadiusDynamicAuthorDomainStrippingVariable.IsNull() {
		body, _ = sjson.Set(body, path+"radius-dynamic-author.domain-stripping."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.domain-stripping."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.domain-stripping."+"vipVariableName", data.RadiusDynamicAuthorDomainStrippingVariable.ValueString())
	} else if data.RadiusDynamicAuthorDomainStripping.IsNull() {
		body, _ = sjson.Set(body, path+"radius-dynamic-author.domain-stripping."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.domain-stripping."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"radius-dynamic-author.domain-stripping."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.domain-stripping."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.domain-stripping."+"vipValue", data.RadiusDynamicAuthorDomainStripping.ValueString())
	}

	if !data.RadiusDynamicAuthorAuthenticationTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"radius-dynamic-author.auth-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.auth-type."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.auth-type."+"vipVariableName", data.RadiusDynamicAuthorAuthenticationTypeVariable.ValueString())
	} else if data.RadiusDynamicAuthorAuthenticationType.IsNull() {
		body, _ = sjson.Set(body, path+"radius-dynamic-author.auth-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.auth-type."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"radius-dynamic-author.auth-type."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.auth-type."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.auth-type."+"vipValue", data.RadiusDynamicAuthorAuthenticationType.ValueString())
	}

	if !data.RadiusDynamicAuthorPortVariable.IsNull() {
		body, _ = sjson.Set(body, path+"radius-dynamic-author.port."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.port."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.port."+"vipVariableName", data.RadiusDynamicAuthorPortVariable.ValueString())
	} else if data.RadiusDynamicAuthorPort.IsNull() {
		body, _ = sjson.Set(body, path+"radius-dynamic-author.port."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.port."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"radius-dynamic-author.port."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.port."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"radius-dynamic-author.port."+"vipValue", data.RadiusDynamicAuthorPort.ValueInt64())
	}

	if !data.RadiusTrustsecCtsAuthorizationListVariable.IsNull() {
		body, _ = sjson.Set(body, path+"radius-trustsec.cts-auth-list."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-trustsec.cts-auth-list."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"radius-trustsec.cts-auth-list."+"vipVariableName", data.RadiusTrustsecCtsAuthorizationListVariable.ValueString())
	} else if data.RadiusTrustsecCtsAuthorizationList.IsNull() {
		body, _ = sjson.Set(body, path+"radius-trustsec.cts-auth-list."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-trustsec.cts-auth-list."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"radius-trustsec.cts-auth-list."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-trustsec.cts-auth-list."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"radius-trustsec.cts-auth-list."+"vipValue", data.RadiusTrustsecCtsAuthorizationList.ValueString())
	}
	if data.RadiusTrustsecGroup.IsNull() {
		body, _ = sjson.Set(body, path+"radius-trustsec.radius-trustsec-group."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-trustsec.radius-trustsec-group."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"radius-trustsec.radius-trustsec-group."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"radius-trustsec.radius-trustsec-group."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"radius-trustsec.radius-trustsec-group."+"vipValue", data.RadiusTrustsecGroup.ValueString())
	}
	if len(data.TacacsServerGroups) > 0 {
		body, _ = sjson.Set(body, path+"tacacs."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"tacacs."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tacacs."+"vipPrimaryKey", []string{"group-name"})
		body, _ = sjson.Set(body, path+"tacacs."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"tacacs."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"tacacs."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"tacacs."+"vipPrimaryKey", []string{"group-name"})
		body, _ = sjson.Set(body, path+"tacacs."+"vipValue", []interface{}{})
	}
	for _, item := range data.TacacsServerGroups {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "group-name")
		if item.GroupName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "group-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "group-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "group-name."+"vipValue", item.GroupName.ValueString())
		}
		itemAttributes = append(itemAttributes, "vpn")
		if item.VpnId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", item.VpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "source-interface")

		if !item.SourceInterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipVariableName", item.SourceInterfaceVariable.ValueString())
		} else if item.SourceInterface.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipValue", item.SourceInterface.ValueString())
		}
		itemAttributes = append(itemAttributes, "server")
		if len(item.Servers) > 0 {
			itemBody, _ = sjson.Set(itemBody, "server."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "server."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "server."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "server."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "server."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "server."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "server."+"vipPrimaryKey", []string{"address"})
			itemBody, _ = sjson.Set(itemBody, "server."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Servers {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "address")
			if childItem.Address.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "address."+"vipValue", childItem.Address.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "port")

			if !childItem.PortVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "port."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "port."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "port."+"vipVariableName", childItem.PortVariable.ValueString())
			} else if childItem.Port.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "port."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "port."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "port."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "port."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "port."+"vipValue", childItem.Port.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "timeout")

			if !childItem.TimeoutVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "timeout."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "timeout."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "timeout."+"vipVariableName", childItem.TimeoutVariable.ValueString())
			} else if childItem.Timeout.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "timeout."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "timeout."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "timeout."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "timeout."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "timeout."+"vipValue", childItem.Timeout.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "key")
			if childItem.Key.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "key."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "key."+"vipValue", childItem.Key.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "secret-key")

			if !childItem.SecretKeyVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "secret-key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "secret-key."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "secret-key."+"vipVariableName", childItem.SecretKeyVariable.ValueString())
			} else if childItem.SecretKey.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "secret-key."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "secret-key."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "secret-key."+"vipValue", childItem.SecretKey.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "key-enum")
			if childItem.EncryptionType.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "key-enum."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-enum."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "key-enum."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-enum."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "key-enum."+"vipValue", childItem.EncryptionType.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "server."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"tacacs."+"vipValue.-1", itemBody)
	}
	if len(data.AccountingRules) > 0 {
		body, _ = sjson.Set(body, path+"accounting.accounting-rule."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"accounting.accounting-rule."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"accounting.accounting-rule."+"vipPrimaryKey", []string{"rule-id"})
		body, _ = sjson.Set(body, path+"accounting.accounting-rule."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"accounting.accounting-rule."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"accounting.accounting-rule."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"accounting.accounting-rule."+"vipPrimaryKey", []string{"rule-id"})
		body, _ = sjson.Set(body, path+"accounting.accounting-rule."+"vipValue", []interface{}{})
	}
	for _, item := range data.AccountingRules {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "rule-id")
		if item.Name.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "rule-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "rule-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "rule-id."+"vipValue", item.Name.ValueString())
		}
		itemAttributes = append(itemAttributes, "method")
		if item.Method.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "method."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "method."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "method."+"vipValue", item.Method.ValueString())
		}
		itemAttributes = append(itemAttributes, "level")
		if item.PrivilegeLevel.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "level."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "level."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "level."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "level."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "level."+"vipValue", item.PrivilegeLevel.ValueString())
		}
		itemAttributes = append(itemAttributes, "start-stop")

		if !item.StartStopVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "start-stop."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "start-stop."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "start-stop."+"vipVariableName", item.StartStopVariable.ValueString())
		} else if item.StartStop.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "start-stop."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "start-stop."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "start-stop."+"vipValue", strconv.FormatBool(item.StartStop.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "group")
		if item.Groups.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "group."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "group."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "group."+"vipValue", item.Groups.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"accounting.accounting-rule."+"vipValue.-1", itemBody)
	}

	if !data.AuthorizationConsoleVariable.IsNull() {
		body, _ = sjson.Set(body, path+"authorization.authorization-console."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"authorization.authorization-console."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"authorization.authorization-console."+"vipVariableName", data.AuthorizationConsoleVariable.ValueString())
	} else if data.AuthorizationConsole.IsNull() {
		body, _ = sjson.Set(body, path+"authorization.authorization-console."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"authorization.authorization-console."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"authorization.authorization-console."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"authorization.authorization-console."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"authorization.authorization-console."+"vipValue", strconv.FormatBool(data.AuthorizationConsole.ValueBool()))
	}

	if !data.AuthorizationConfigCommandsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"authorization.authorization-config-commands."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"authorization.authorization-config-commands."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"authorization.authorization-config-commands."+"vipVariableName", data.AuthorizationConfigCommandsVariable.ValueString())
	} else if data.AuthorizationConfigCommands.IsNull() {
		body, _ = sjson.Set(body, path+"authorization.authorization-config-commands."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"authorization.authorization-config-commands."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"authorization.authorization-config-commands."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"authorization.authorization-config-commands."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"authorization.authorization-config-commands."+"vipValue", strconv.FormatBool(data.AuthorizationConfigCommands.ValueBool()))
	}
	if len(data.AuthorizationRules) > 0 {
		body, _ = sjson.Set(body, path+"authorization.authorization-rule."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"authorization.authorization-rule."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"authorization.authorization-rule."+"vipPrimaryKey", []string{"rule-id"})
		body, _ = sjson.Set(body, path+"authorization.authorization-rule."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"authorization.authorization-rule."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"authorization.authorization-rule."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"authorization.authorization-rule."+"vipPrimaryKey", []string{"rule-id"})
		body, _ = sjson.Set(body, path+"authorization.authorization-rule."+"vipValue", []interface{}{})
	}
	for _, item := range data.AuthorizationRules {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "rule-id")
		if item.Name.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "rule-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "rule-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "rule-id."+"vipValue", item.Name.ValueString())
		}
		itemAttributes = append(itemAttributes, "method")
		if item.Method.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "method."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "method."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "method."+"vipValue", item.Method.ValueString())
		}
		itemAttributes = append(itemAttributes, "level")
		if item.PrivilegeLevel.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "level."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "level."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "level."+"vipValue", item.PrivilegeLevel.ValueString())
		}
		itemAttributes = append(itemAttributes, "group")
		if item.Groups.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "group."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "group."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "group."+"vipValue", item.Groups.ValueString())
		}
		itemAttributes = append(itemAttributes, "if-authenticated")
		if item.Authenticated.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "if-authenticated."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "if-authenticated."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "if-authenticated."+"vipValue", strconv.FormatBool(item.Authenticated.ValueBool()))
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"authorization.authorization-rule."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CEdgeAAA) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "authentication.dot1x.default.authentication_group.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Dot1xAuthentication = types.BoolNull()

			v := res.Get(path + "authentication.dot1x.default.authentication_group.vipVariableName")
			data.Dot1xAuthenticationVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Dot1xAuthentication = types.BoolNull()
			data.Dot1xAuthenticationVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "authentication.dot1x.default.authentication_group.vipValue")
			data.Dot1xAuthentication = types.BoolValue(v.Bool())
			data.Dot1xAuthenticationVariable = types.StringNull()
		}
	} else {
		data.Dot1xAuthentication = types.BoolNull()
		data.Dot1xAuthenticationVariable = types.StringNull()
	}
	if value := res.Get(path + "accounting.dot1x.default.start-stop.accounting_group.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Dot1xAccounting = types.BoolNull()

			v := res.Get(path + "accounting.dot1x.default.start-stop.accounting_group.vipVariableName")
			data.Dot1xAccountingVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Dot1xAccounting = types.BoolNull()
			data.Dot1xAccountingVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "accounting.dot1x.default.start-stop.accounting_group.vipValue")
			data.Dot1xAccounting = types.BoolValue(v.Bool())
			data.Dot1xAccountingVariable = types.StringNull()
		}
	} else {
		data.Dot1xAccounting = types.BoolNull()
		data.Dot1xAccountingVariable = types.StringNull()
	}
	if value := res.Get(path + "server-auth-order.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ServerGroupsPriorityOrder = types.StringNull()

		} else if value.String() == "ignore" {
			data.ServerGroupsPriorityOrder = types.StringNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "server-auth-order.vipValue")
			data.ServerGroupsPriorityOrder = types.StringValue(v.String())

		}
	} else {
		data.ServerGroupsPriorityOrder = types.StringNull()

	}
	if value := res.Get(path + "user.vipValue"); len(value.Array()) > 0 {
		data.Users = make([]CEdgeAAAUsers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CEdgeAAAUsers{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Name = types.StringNull()

					cv := v.Get("name.vipVariableName")
					item.NameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Name = types.StringNull()
					item.NameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.Name = types.StringValue(cv.String())
					item.NameVariable = types.StringNull()
				}
			} else {
				item.Name = types.StringNull()
				item.NameVariable = types.StringNull()
			}
			if cValue := v.Get("password.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Password = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Password = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("password.vipValue")
					item.Password = types.StringValue(cv.String())

				}
			} else {
				item.Password = types.StringNull()

			}
			if cValue := v.Get("secret.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Secret = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Secret = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("secret.vipValue")
					item.Secret = types.StringValue(cv.String())

				}
			} else {
				item.Secret = types.StringNull()

			}
			if cValue := v.Get("privilege.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PrivilegeLevel = types.StringNull()

					cv := v.Get("privilege.vipVariableName")
					item.PrivilegeLevelVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.PrivilegeLevel = types.StringNull()
					item.PrivilegeLevelVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("privilege.vipValue")
					item.PrivilegeLevel = types.StringValue(cv.String())
					item.PrivilegeLevelVariable = types.StringNull()
				}
			} else {
				item.PrivilegeLevel = types.StringNull()
				item.PrivilegeLevelVariable = types.StringNull()
			}
			if cValue := v.Get("pubkey-chain.vipValue"); len(cValue.Array()) > 0 {
				item.SshPubkeys = make([]CEdgeAAAUsersSshPubkeys, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CEdgeAAAUsersSshPubkeys{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("key-string.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.KeyString = types.StringNull()

						} else if ccValue.String() == "ignore" {
							cItem.KeyString = types.StringNull()

						} else if ccValue.String() == "constant" {
							ccv := cv.Get("key-string.vipValue")
							cItem.KeyString = types.StringValue(ccv.String())

						}
					} else {
						cItem.KeyString = types.StringNull()

					}
					if ccValue := cv.Get("key-type.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.KeyType = types.StringNull()

							ccv := cv.Get("key-type.vipVariableName")
							cItem.KeyTypeVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.KeyType = types.StringNull()
							cItem.KeyTypeVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("key-type.vipValue")
							cItem.KeyType = types.StringValue(ccv.String())
							cItem.KeyTypeVariable = types.StringNull()
						}
					} else {
						cItem.KeyType = types.StringNull()
						cItem.KeyTypeVariable = types.StringNull()
					}
					item.SshPubkeys = append(item.SshPubkeys, cItem)
					return true
				})
			} else {
				if len(item.SshPubkeys) > 0 {
					item.SshPubkeys = []CEdgeAAAUsersSshPubkeys{}
				}
			}
			data.Users = append(data.Users, item)
			return true
		})
	} else {
		if len(data.Users) > 0 {
			data.Users = []CEdgeAAAUsers{}
		}
	}
	if value := res.Get(path + "radius.vipValue"); len(value.Array()) > 0 {
		data.RadiusServerGroups = make([]CEdgeAAARadiusServerGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CEdgeAAARadiusServerGroups{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("group-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.GroupName = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.GroupName = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("group-name.vipValue")
					item.GroupName = types.StringValue(cv.String())

				}
			} else {
				item.GroupName = types.StringNull()

			}
			if cValue := v.Get("vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()

				} else if cValue.String() == "constant" {
					cv := v.Get("vpn.vipValue")
					item.VpnId = types.Int64Value(cv.Int())

				}
			} else {
				item.VpnId = types.Int64Null()

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
			if cValue := v.Get("server.vipValue"); len(cValue.Array()) > 0 {
				item.Servers = make([]CEdgeAAARadiusServerGroupsServers, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CEdgeAAARadiusServerGroupsServers{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("address.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Address = types.StringNull()

						} else if ccValue.String() == "ignore" {
							cItem.Address = types.StringNull()

						} else if ccValue.String() == "constant" {
							ccv := cv.Get("address.vipValue")
							cItem.Address = types.StringValue(ccv.String())

						}
					} else {
						cItem.Address = types.StringNull()

					}
					if ccValue := cv.Get("auth-port.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AuthenticationPort = types.Int64Null()

							ccv := cv.Get("auth-port.vipVariableName")
							cItem.AuthenticationPortVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AuthenticationPort = types.Int64Null()
							cItem.AuthenticationPortVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("auth-port.vipValue")
							cItem.AuthenticationPort = types.Int64Value(ccv.Int())
							cItem.AuthenticationPortVariable = types.StringNull()
						}
					} else {
						cItem.AuthenticationPort = types.Int64Null()
						cItem.AuthenticationPortVariable = types.StringNull()
					}
					if ccValue := cv.Get("acct-port.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AccountingPort = types.Int64Null()

							ccv := cv.Get("acct-port.vipVariableName")
							cItem.AccountingPortVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AccountingPort = types.Int64Null()
							cItem.AccountingPortVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("acct-port.vipValue")
							cItem.AccountingPort = types.Int64Value(ccv.Int())
							cItem.AccountingPortVariable = types.StringNull()
						}
					} else {
						cItem.AccountingPort = types.Int64Null()
						cItem.AccountingPortVariable = types.StringNull()
					}
					if ccValue := cv.Get("timeout.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Timeout = types.Int64Null()

							ccv := cv.Get("timeout.vipVariableName")
							cItem.TimeoutVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Timeout = types.Int64Null()
							cItem.TimeoutVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("timeout.vipValue")
							cItem.Timeout = types.Int64Value(ccv.Int())
							cItem.TimeoutVariable = types.StringNull()
						}
					} else {
						cItem.Timeout = types.Int64Null()
						cItem.TimeoutVariable = types.StringNull()
					}
					if ccValue := cv.Get("retransmit.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Retransmit = types.Int64Null()

							ccv := cv.Get("retransmit.vipVariableName")
							cItem.RetransmitVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Retransmit = types.Int64Null()
							cItem.RetransmitVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("retransmit.vipValue")
							cItem.Retransmit = types.Int64Value(ccv.Int())
							cItem.RetransmitVariable = types.StringNull()
						}
					} else {
						cItem.Retransmit = types.Int64Null()
						cItem.RetransmitVariable = types.StringNull()
					}
					if ccValue := cv.Get("key.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Key = types.StringNull()

						} else if ccValue.String() == "ignore" {
							cItem.Key = types.StringNull()

						} else if ccValue.String() == "constant" {
							ccv := cv.Get("key.vipValue")
							cItem.Key = types.StringValue(ccv.String())

						}
					} else {
						cItem.Key = types.StringNull()

					}
					if ccValue := cv.Get("secret-key.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.SecretKey = types.StringNull()

							ccv := cv.Get("secret-key.vipVariableName")
							cItem.SecretKeyVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.SecretKey = types.StringNull()
							cItem.SecretKeyVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("secret-key.vipValue")
							cItem.SecretKey = types.StringValue(ccv.String())
							cItem.SecretKeyVariable = types.StringNull()
						}
					} else {
						cItem.SecretKey = types.StringNull()
						cItem.SecretKeyVariable = types.StringNull()
					}
					if ccValue := cv.Get("key-enum.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.EncryptionType = types.StringNull()

						} else if ccValue.String() == "ignore" {
							cItem.EncryptionType = types.StringNull()

						} else if ccValue.String() == "constant" {
							ccv := cv.Get("key-enum.vipValue")
							cItem.EncryptionType = types.StringValue(ccv.String())

						}
					} else {
						cItem.EncryptionType = types.StringNull()

					}
					if ccValue := cv.Get("key-type.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.KeyType = types.StringNull()

							ccv := cv.Get("key-type.vipVariableName")
							cItem.KeyTypeVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.KeyType = types.StringNull()
							cItem.KeyTypeVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("key-type.vipValue")
							cItem.KeyType = types.StringValue(ccv.String())
							cItem.KeyTypeVariable = types.StringNull()
						}
					} else {
						cItem.KeyType = types.StringNull()
						cItem.KeyTypeVariable = types.StringNull()
					}
					item.Servers = append(item.Servers, cItem)
					return true
				})
			} else {
				if len(item.Servers) > 0 {
					item.Servers = []CEdgeAAARadiusServerGroupsServers{}
				}
			}
			data.RadiusServerGroups = append(data.RadiusServerGroups, item)
			return true
		})
	} else {
		if len(data.RadiusServerGroups) > 0 {
			data.RadiusServerGroups = []CEdgeAAARadiusServerGroups{}
		}
	}
	if value := res.Get(path + "radius-dynamic-author.radius-client.vipValue"); len(value.Array()) > 0 {
		data.RadiusClients = make([]CEdgeAAARadiusClients, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CEdgeAAARadiusClients{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ClientIp = types.StringNull()

					cv := v.Get("ip.vipVariableName")
					item.ClientIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.ClientIp = types.StringNull()
					item.ClientIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ip.vipValue")
					item.ClientIp = types.StringValue(cv.String())
					item.ClientIpVariable = types.StringNull()
				}
			} else {
				item.ClientIp = types.StringNull()
				item.ClientIpVariable = types.StringNull()
			}
			if cValue := v.Get("vpn.vipValue"); len(cValue.Array()) > 0 {
				item.VpnConfigurations = make([]CEdgeAAARadiusClientsVpnConfigurations, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CEdgeAAARadiusClientsVpnConfigurations{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("name.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.VpnId = types.Int64Null()

							ccv := cv.Get("name.vipVariableName")
							cItem.VpnIdVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.VpnId = types.Int64Null()
							cItem.VpnIdVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("name.vipValue")
							cItem.VpnId = types.Int64Value(ccv.Int())
							cItem.VpnIdVariable = types.StringNull()
						}
					} else {
						cItem.VpnId = types.Int64Null()
						cItem.VpnIdVariable = types.StringNull()
					}
					if ccValue := cv.Get("server-key.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.ServerKey = types.StringNull()

							ccv := cv.Get("server-key.vipVariableName")
							cItem.ServerKeyVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.ServerKey = types.StringNull()
							cItem.ServerKeyVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("server-key.vipValue")
							cItem.ServerKey = types.StringValue(ccv.String())
							cItem.ServerKeyVariable = types.StringNull()
						}
					} else {
						cItem.ServerKey = types.StringNull()
						cItem.ServerKeyVariable = types.StringNull()
					}
					item.VpnConfigurations = append(item.VpnConfigurations, cItem)
					return true
				})
			} else {
				if len(item.VpnConfigurations) > 0 {
					item.VpnConfigurations = []CEdgeAAARadiusClientsVpnConfigurations{}
				}
			}
			data.RadiusClients = append(data.RadiusClients, item)
			return true
		})
	} else {
		if len(data.RadiusClients) > 0 {
			data.RadiusClients = []CEdgeAAARadiusClients{}
		}
	}
	if value := res.Get(path + "radius-dynamic-author.rda-server-key.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.RadiusDynamicAuthorServerKey = types.StringNull()

			v := res.Get(path + "radius-dynamic-author.rda-server-key.vipVariableName")
			data.RadiusDynamicAuthorServerKeyVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.RadiusDynamicAuthorServerKey = types.StringNull()
			data.RadiusDynamicAuthorServerKeyVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "radius-dynamic-author.rda-server-key.vipValue")
			data.RadiusDynamicAuthorServerKey = types.StringValue(v.String())
			data.RadiusDynamicAuthorServerKeyVariable = types.StringNull()
		}
	} else {
		data.RadiusDynamicAuthorServerKey = types.StringNull()
		data.RadiusDynamicAuthorServerKeyVariable = types.StringNull()
	}
	if value := res.Get(path + "radius-dynamic-author.domain-stripping.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.RadiusDynamicAuthorDomainStripping = types.StringNull()

			v := res.Get(path + "radius-dynamic-author.domain-stripping.vipVariableName")
			data.RadiusDynamicAuthorDomainStrippingVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.RadiusDynamicAuthorDomainStripping = types.StringNull()
			data.RadiusDynamicAuthorDomainStrippingVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "radius-dynamic-author.domain-stripping.vipValue")
			data.RadiusDynamicAuthorDomainStripping = types.StringValue(v.String())
			data.RadiusDynamicAuthorDomainStrippingVariable = types.StringNull()
		}
	} else {
		data.RadiusDynamicAuthorDomainStripping = types.StringNull()
		data.RadiusDynamicAuthorDomainStrippingVariable = types.StringNull()
	}
	if value := res.Get(path + "radius-dynamic-author.auth-type.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.RadiusDynamicAuthorAuthenticationType = types.StringNull()

			v := res.Get(path + "radius-dynamic-author.auth-type.vipVariableName")
			data.RadiusDynamicAuthorAuthenticationTypeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.RadiusDynamicAuthorAuthenticationType = types.StringNull()
			data.RadiusDynamicAuthorAuthenticationTypeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "radius-dynamic-author.auth-type.vipValue")
			data.RadiusDynamicAuthorAuthenticationType = types.StringValue(v.String())
			data.RadiusDynamicAuthorAuthenticationTypeVariable = types.StringNull()
		}
	} else {
		data.RadiusDynamicAuthorAuthenticationType = types.StringNull()
		data.RadiusDynamicAuthorAuthenticationTypeVariable = types.StringNull()
	}
	if value := res.Get(path + "radius-dynamic-author.port.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.RadiusDynamicAuthorPort = types.Int64Null()

			v := res.Get(path + "radius-dynamic-author.port.vipVariableName")
			data.RadiusDynamicAuthorPortVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.RadiusDynamicAuthorPort = types.Int64Null()
			data.RadiusDynamicAuthorPortVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "radius-dynamic-author.port.vipValue")
			data.RadiusDynamicAuthorPort = types.Int64Value(v.Int())
			data.RadiusDynamicAuthorPortVariable = types.StringNull()
		}
	} else {
		data.RadiusDynamicAuthorPort = types.Int64Null()
		data.RadiusDynamicAuthorPortVariable = types.StringNull()
	}
	if value := res.Get(path + "radius-trustsec.cts-auth-list.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.RadiusTrustsecCtsAuthorizationList = types.StringNull()

			v := res.Get(path + "radius-trustsec.cts-auth-list.vipVariableName")
			data.RadiusTrustsecCtsAuthorizationListVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.RadiusTrustsecCtsAuthorizationList = types.StringNull()
			data.RadiusTrustsecCtsAuthorizationListVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "radius-trustsec.cts-auth-list.vipValue")
			data.RadiusTrustsecCtsAuthorizationList = types.StringValue(v.String())
			data.RadiusTrustsecCtsAuthorizationListVariable = types.StringNull()
		}
	} else {
		data.RadiusTrustsecCtsAuthorizationList = types.StringNull()
		data.RadiusTrustsecCtsAuthorizationListVariable = types.StringNull()
	}
	if value := res.Get(path + "radius-trustsec.radius-trustsec-group.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.RadiusTrustsecGroup = types.StringNull()

		} else if value.String() == "ignore" {
			data.RadiusTrustsecGroup = types.StringNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "radius-trustsec.radius-trustsec-group.vipValue")
			data.RadiusTrustsecGroup = types.StringValue(v.String())

		}
	} else {
		data.RadiusTrustsecGroup = types.StringNull()

	}
	if value := res.Get(path + "tacacs.vipValue"); len(value.Array()) > 0 {
		data.TacacsServerGroups = make([]CEdgeAAATacacsServerGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CEdgeAAATacacsServerGroups{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("group-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.GroupName = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.GroupName = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("group-name.vipValue")
					item.GroupName = types.StringValue(cv.String())

				}
			} else {
				item.GroupName = types.StringNull()

			}
			if cValue := v.Get("vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()

				} else if cValue.String() == "constant" {
					cv := v.Get("vpn.vipValue")
					item.VpnId = types.Int64Value(cv.Int())

				}
			} else {
				item.VpnId = types.Int64Null()

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
			if cValue := v.Get("server.vipValue"); len(cValue.Array()) > 0 {
				item.Servers = make([]CEdgeAAATacacsServerGroupsServers, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CEdgeAAATacacsServerGroupsServers{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("address.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Address = types.StringNull()

						} else if ccValue.String() == "ignore" {
							cItem.Address = types.StringNull()

						} else if ccValue.String() == "constant" {
							ccv := cv.Get("address.vipValue")
							cItem.Address = types.StringValue(ccv.String())

						}
					} else {
						cItem.Address = types.StringNull()

					}
					if ccValue := cv.Get("port.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Port = types.Int64Null()

							ccv := cv.Get("port.vipVariableName")
							cItem.PortVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Port = types.Int64Null()
							cItem.PortVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("port.vipValue")
							cItem.Port = types.Int64Value(ccv.Int())
							cItem.PortVariable = types.StringNull()
						}
					} else {
						cItem.Port = types.Int64Null()
						cItem.PortVariable = types.StringNull()
					}
					if ccValue := cv.Get("timeout.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Timeout = types.Int64Null()

							ccv := cv.Get("timeout.vipVariableName")
							cItem.TimeoutVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Timeout = types.Int64Null()
							cItem.TimeoutVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("timeout.vipValue")
							cItem.Timeout = types.Int64Value(ccv.Int())
							cItem.TimeoutVariable = types.StringNull()
						}
					} else {
						cItem.Timeout = types.Int64Null()
						cItem.TimeoutVariable = types.StringNull()
					}
					if ccValue := cv.Get("key.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Key = types.StringNull()

						} else if ccValue.String() == "ignore" {
							cItem.Key = types.StringNull()

						} else if ccValue.String() == "constant" {
							ccv := cv.Get("key.vipValue")
							cItem.Key = types.StringValue(ccv.String())

						}
					} else {
						cItem.Key = types.StringNull()

					}
					if ccValue := cv.Get("secret-key.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.SecretKey = types.StringNull()

							ccv := cv.Get("secret-key.vipVariableName")
							cItem.SecretKeyVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.SecretKey = types.StringNull()
							cItem.SecretKeyVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("secret-key.vipValue")
							cItem.SecretKey = types.StringValue(ccv.String())
							cItem.SecretKeyVariable = types.StringNull()
						}
					} else {
						cItem.SecretKey = types.StringNull()
						cItem.SecretKeyVariable = types.StringNull()
					}
					if ccValue := cv.Get("key-enum.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.EncryptionType = types.StringNull()

						} else if ccValue.String() == "ignore" {
							cItem.EncryptionType = types.StringNull()

						} else if ccValue.String() == "constant" {
							ccv := cv.Get("key-enum.vipValue")
							cItem.EncryptionType = types.StringValue(ccv.String())

						}
					} else {
						cItem.EncryptionType = types.StringNull()

					}
					item.Servers = append(item.Servers, cItem)
					return true
				})
			} else {
				if len(item.Servers) > 0 {
					item.Servers = []CEdgeAAATacacsServerGroupsServers{}
				}
			}
			data.TacacsServerGroups = append(data.TacacsServerGroups, item)
			return true
		})
	} else {
		if len(data.TacacsServerGroups) > 0 {
			data.TacacsServerGroups = []CEdgeAAATacacsServerGroups{}
		}
	}
	if value := res.Get(path + "accounting.accounting-rule.vipValue"); len(value.Array()) > 0 {
		data.AccountingRules = make([]CEdgeAAAAccountingRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CEdgeAAAAccountingRules{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("rule-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Name = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Name = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("rule-id.vipValue")
					item.Name = types.StringValue(cv.String())

				}
			} else {
				item.Name = types.StringNull()

			}
			if cValue := v.Get("method.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Method = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Method = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("method.vipValue")
					item.Method = types.StringValue(cv.String())

				}
			} else {
				item.Method = types.StringNull()

			}
			if cValue := v.Get("level.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PrivilegeLevel = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.PrivilegeLevel = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("level.vipValue")
					item.PrivilegeLevel = types.StringValue(cv.String())

				}
			} else {
				item.PrivilegeLevel = types.StringNull()

			}
			if cValue := v.Get("start-stop.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.StartStop = types.BoolNull()

					cv := v.Get("start-stop.vipVariableName")
					item.StartStopVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.StartStop = types.BoolNull()
					item.StartStopVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("start-stop.vipValue")
					item.StartStop = types.BoolValue(cv.Bool())
					item.StartStopVariable = types.StringNull()
				}
			} else {
				item.StartStop = types.BoolNull()
				item.StartStopVariable = types.StringNull()
			}
			if cValue := v.Get("group.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Groups = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Groups = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("group.vipValue")
					item.Groups = types.StringValue(cv.String())

				}
			} else {
				item.Groups = types.StringNull()

			}
			data.AccountingRules = append(data.AccountingRules, item)
			return true
		})
	} else {
		if len(data.AccountingRules) > 0 {
			data.AccountingRules = []CEdgeAAAAccountingRules{}
		}
	}
	if value := res.Get(path + "authorization.authorization-console.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AuthorizationConsole = types.BoolNull()

			v := res.Get(path + "authorization.authorization-console.vipVariableName")
			data.AuthorizationConsoleVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AuthorizationConsole = types.BoolNull()
			data.AuthorizationConsoleVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "authorization.authorization-console.vipValue")
			data.AuthorizationConsole = types.BoolValue(v.Bool())
			data.AuthorizationConsoleVariable = types.StringNull()
		}
	} else {
		data.AuthorizationConsole = types.BoolNull()
		data.AuthorizationConsoleVariable = types.StringNull()
	}
	if value := res.Get(path + "authorization.authorization-config-commands.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AuthorizationConfigCommands = types.BoolNull()

			v := res.Get(path + "authorization.authorization-config-commands.vipVariableName")
			data.AuthorizationConfigCommandsVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AuthorizationConfigCommands = types.BoolNull()
			data.AuthorizationConfigCommandsVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "authorization.authorization-config-commands.vipValue")
			data.AuthorizationConfigCommands = types.BoolValue(v.Bool())
			data.AuthorizationConfigCommandsVariable = types.StringNull()
		}
	} else {
		data.AuthorizationConfigCommands = types.BoolNull()
		data.AuthorizationConfigCommandsVariable = types.StringNull()
	}
	if value := res.Get(path + "authorization.authorization-rule.vipValue"); len(value.Array()) > 0 {
		data.AuthorizationRules = make([]CEdgeAAAAuthorizationRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CEdgeAAAAuthorizationRules{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("rule-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Name = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Name = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("rule-id.vipValue")
					item.Name = types.StringValue(cv.String())

				}
			} else {
				item.Name = types.StringNull()

			}
			if cValue := v.Get("method.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Method = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Method = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("method.vipValue")
					item.Method = types.StringValue(cv.String())

				}
			} else {
				item.Method = types.StringNull()

			}
			if cValue := v.Get("level.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PrivilegeLevel = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.PrivilegeLevel = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("level.vipValue")
					item.PrivilegeLevel = types.StringValue(cv.String())

				}
			} else {
				item.PrivilegeLevel = types.StringNull()

			}
			if cValue := v.Get("group.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Groups = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Groups = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("group.vipValue")
					item.Groups = types.StringValue(cv.String())

				}
			} else {
				item.Groups = types.StringNull()

			}
			if cValue := v.Get("if-authenticated.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Authenticated = types.BoolNull()

				} else if cValue.String() == "ignore" {
					item.Authenticated = types.BoolNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("if-authenticated.vipValue")
					item.Authenticated = types.BoolValue(cv.Bool())

				}
			} else {
				item.Authenticated = types.BoolNull()

			}
			data.AuthorizationRules = append(data.AuthorizationRules, item)
			return true
		})
	} else {
		if len(data.AuthorizationRules) > 0 {
			data.AuthorizationRules = []CEdgeAAAAuthorizationRules{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CEdgeAAA) hasChanges(ctx context.Context, state *CEdgeAAA) bool {
	hasChanges := false
	if !data.Dot1xAuthentication.Equal(state.Dot1xAuthentication) {
		hasChanges = true
	}
	if !data.Dot1xAccounting.Equal(state.Dot1xAccounting) {
		hasChanges = true
	}
	if !data.ServerGroupsPriorityOrder.Equal(state.ServerGroupsPriorityOrder) {
		hasChanges = true
	}
	if len(data.Users) != len(state.Users) {
		hasChanges = true
	} else {
		for i := range data.Users {
			if !data.Users[i].Name.Equal(state.Users[i].Name) {
				hasChanges = true
			}
			if !data.Users[i].Password.Equal(state.Users[i].Password) {
				hasChanges = true
			}
			if !data.Users[i].Secret.Equal(state.Users[i].Secret) {
				hasChanges = true
			}
			if !data.Users[i].PrivilegeLevel.Equal(state.Users[i].PrivilegeLevel) {
				hasChanges = true
			}
			if len(data.Users[i].SshPubkeys) != len(state.Users[i].SshPubkeys) {
				hasChanges = true
			} else {
				for ii := range data.Users[i].SshPubkeys {
					if !data.Users[i].SshPubkeys[ii].KeyString.Equal(state.Users[i].SshPubkeys[ii].KeyString) {
						hasChanges = true
					}
					if !data.Users[i].SshPubkeys[ii].KeyType.Equal(state.Users[i].SshPubkeys[ii].KeyType) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.RadiusServerGroups) != len(state.RadiusServerGroups) {
		hasChanges = true
	} else {
		for i := range data.RadiusServerGroups {
			if !data.RadiusServerGroups[i].GroupName.Equal(state.RadiusServerGroups[i].GroupName) {
				hasChanges = true
			}
			if !data.RadiusServerGroups[i].VpnId.Equal(state.RadiusServerGroups[i].VpnId) {
				hasChanges = true
			}
			if !data.RadiusServerGroups[i].SourceInterface.Equal(state.RadiusServerGroups[i].SourceInterface) {
				hasChanges = true
			}
			if len(data.RadiusServerGroups[i].Servers) != len(state.RadiusServerGroups[i].Servers) {
				hasChanges = true
			} else {
				for ii := range data.RadiusServerGroups[i].Servers {
					if !data.RadiusServerGroups[i].Servers[ii].Address.Equal(state.RadiusServerGroups[i].Servers[ii].Address) {
						hasChanges = true
					}
					if !data.RadiusServerGroups[i].Servers[ii].AuthenticationPort.Equal(state.RadiusServerGroups[i].Servers[ii].AuthenticationPort) {
						hasChanges = true
					}
					if !data.RadiusServerGroups[i].Servers[ii].AccountingPort.Equal(state.RadiusServerGroups[i].Servers[ii].AccountingPort) {
						hasChanges = true
					}
					if !data.RadiusServerGroups[i].Servers[ii].Timeout.Equal(state.RadiusServerGroups[i].Servers[ii].Timeout) {
						hasChanges = true
					}
					if !data.RadiusServerGroups[i].Servers[ii].Retransmit.Equal(state.RadiusServerGroups[i].Servers[ii].Retransmit) {
						hasChanges = true
					}
					if !data.RadiusServerGroups[i].Servers[ii].Key.Equal(state.RadiusServerGroups[i].Servers[ii].Key) {
						hasChanges = true
					}
					if !data.RadiusServerGroups[i].Servers[ii].SecretKey.Equal(state.RadiusServerGroups[i].Servers[ii].SecretKey) {
						hasChanges = true
					}
					if !data.RadiusServerGroups[i].Servers[ii].EncryptionType.Equal(state.RadiusServerGroups[i].Servers[ii].EncryptionType) {
						hasChanges = true
					}
					if !data.RadiusServerGroups[i].Servers[ii].KeyType.Equal(state.RadiusServerGroups[i].Servers[ii].KeyType) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.RadiusClients) != len(state.RadiusClients) {
		hasChanges = true
	} else {
		for i := range data.RadiusClients {
			if !data.RadiusClients[i].ClientIp.Equal(state.RadiusClients[i].ClientIp) {
				hasChanges = true
			}
			if len(data.RadiusClients[i].VpnConfigurations) != len(state.RadiusClients[i].VpnConfigurations) {
				hasChanges = true
			} else {
				for ii := range data.RadiusClients[i].VpnConfigurations {
					if !data.RadiusClients[i].VpnConfigurations[ii].VpnId.Equal(state.RadiusClients[i].VpnConfigurations[ii].VpnId) {
						hasChanges = true
					}
					if !data.RadiusClients[i].VpnConfigurations[ii].ServerKey.Equal(state.RadiusClients[i].VpnConfigurations[ii].ServerKey) {
						hasChanges = true
					}
				}
			}
		}
	}
	if !data.RadiusDynamicAuthorServerKey.Equal(state.RadiusDynamicAuthorServerKey) {
		hasChanges = true
	}
	if !data.RadiusDynamicAuthorDomainStripping.Equal(state.RadiusDynamicAuthorDomainStripping) {
		hasChanges = true
	}
	if !data.RadiusDynamicAuthorAuthenticationType.Equal(state.RadiusDynamicAuthorAuthenticationType) {
		hasChanges = true
	}
	if !data.RadiusDynamicAuthorPort.Equal(state.RadiusDynamicAuthorPort) {
		hasChanges = true
	}
	if !data.RadiusTrustsecCtsAuthorizationList.Equal(state.RadiusTrustsecCtsAuthorizationList) {
		hasChanges = true
	}
	if !data.RadiusTrustsecGroup.Equal(state.RadiusTrustsecGroup) {
		hasChanges = true
	}
	if len(data.TacacsServerGroups) != len(state.TacacsServerGroups) {
		hasChanges = true
	} else {
		for i := range data.TacacsServerGroups {
			if !data.TacacsServerGroups[i].GroupName.Equal(state.TacacsServerGroups[i].GroupName) {
				hasChanges = true
			}
			if !data.TacacsServerGroups[i].VpnId.Equal(state.TacacsServerGroups[i].VpnId) {
				hasChanges = true
			}
			if !data.TacacsServerGroups[i].SourceInterface.Equal(state.TacacsServerGroups[i].SourceInterface) {
				hasChanges = true
			}
			if len(data.TacacsServerGroups[i].Servers) != len(state.TacacsServerGroups[i].Servers) {
				hasChanges = true
			} else {
				for ii := range data.TacacsServerGroups[i].Servers {
					if !data.TacacsServerGroups[i].Servers[ii].Address.Equal(state.TacacsServerGroups[i].Servers[ii].Address) {
						hasChanges = true
					}
					if !data.TacacsServerGroups[i].Servers[ii].Port.Equal(state.TacacsServerGroups[i].Servers[ii].Port) {
						hasChanges = true
					}
					if !data.TacacsServerGroups[i].Servers[ii].Timeout.Equal(state.TacacsServerGroups[i].Servers[ii].Timeout) {
						hasChanges = true
					}
					if !data.TacacsServerGroups[i].Servers[ii].Key.Equal(state.TacacsServerGroups[i].Servers[ii].Key) {
						hasChanges = true
					}
					if !data.TacacsServerGroups[i].Servers[ii].SecretKey.Equal(state.TacacsServerGroups[i].Servers[ii].SecretKey) {
						hasChanges = true
					}
					if !data.TacacsServerGroups[i].Servers[ii].EncryptionType.Equal(state.TacacsServerGroups[i].Servers[ii].EncryptionType) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.AccountingRules) != len(state.AccountingRules) {
		hasChanges = true
	} else {
		for i := range data.AccountingRules {
			if !data.AccountingRules[i].Name.Equal(state.AccountingRules[i].Name) {
				hasChanges = true
			}
			if !data.AccountingRules[i].Method.Equal(state.AccountingRules[i].Method) {
				hasChanges = true
			}
			if !data.AccountingRules[i].PrivilegeLevel.Equal(state.AccountingRules[i].PrivilegeLevel) {
				hasChanges = true
			}
			if !data.AccountingRules[i].StartStop.Equal(state.AccountingRules[i].StartStop) {
				hasChanges = true
			}
			if !data.AccountingRules[i].Groups.Equal(state.AccountingRules[i].Groups) {
				hasChanges = true
			}
		}
	}
	if !data.AuthorizationConsole.Equal(state.AuthorizationConsole) {
		hasChanges = true
	}
	if !data.AuthorizationConfigCommands.Equal(state.AuthorizationConfigCommands) {
		hasChanges = true
	}
	if len(data.AuthorizationRules) != len(state.AuthorizationRules) {
		hasChanges = true
	} else {
		for i := range data.AuthorizationRules {
			if !data.AuthorizationRules[i].Name.Equal(state.AuthorizationRules[i].Name) {
				hasChanges = true
			}
			if !data.AuthorizationRules[i].Method.Equal(state.AuthorizationRules[i].Method) {
				hasChanges = true
			}
			if !data.AuthorizationRules[i].PrivilegeLevel.Equal(state.AuthorizationRules[i].PrivilegeLevel) {
				hasChanges = true
			}
			if !data.AuthorizationRules[i].Groups.Equal(state.AuthorizationRules[i].Groups) {
				hasChanges = true
			}
			if !data.AuthorizationRules[i].Authenticated.Equal(state.AuthorizationRules[i].Authenticated) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
