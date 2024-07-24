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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type SystemAAA struct {
	Id                                  types.String                  `tfsdk:"id"`
	Version                             types.Int64                   `tfsdk:"version"`
	Name                                types.String                  `tfsdk:"name"`
	Description                         types.String                  `tfsdk:"description"`
	FeatureProfileId                    types.String                  `tfsdk:"feature_profile_id"`
	AuthenticationGroup                 types.Bool                    `tfsdk:"authentication_group"`
	AuthenticationGroupVariable         types.String                  `tfsdk:"authentication_group_variable"`
	AccountingGroup                     types.Bool                    `tfsdk:"accounting_group"`
	AccountingGroupVariable             types.String                  `tfsdk:"accounting_group_variable"`
	ServerAuthOrder                     types.Set                     `tfsdk:"server_auth_order"`
	Users                               []SystemAAAUsers              `tfsdk:"users"`
	RadiusGroups                        []SystemAAARadiusGroups       `tfsdk:"radius_groups"`
	TacacsGroups                        []SystemAAATacacsGroups       `tfsdk:"tacacs_groups"`
	AccountingRules                     []SystemAAAAccountingRules    `tfsdk:"accounting_rules"`
	AuthorizationConsole                types.Bool                    `tfsdk:"authorization_console"`
	AuthorizationConsoleVariable        types.String                  `tfsdk:"authorization_console_variable"`
	AuthorizationConfigCommands         types.Bool                    `tfsdk:"authorization_config_commands"`
	AuthorizationConfigCommandsVariable types.String                  `tfsdk:"authorization_config_commands_variable"`
	AuthorizationRules                  []SystemAAAAuthorizationRules `tfsdk:"authorization_rules"`
}

type SystemAAAUsers struct {
	Name              types.String               `tfsdk:"name"`
	NameVariable      types.String               `tfsdk:"name_variable"`
	Password          types.String               `tfsdk:"password"`
	PasswordVariable  types.String               `tfsdk:"password_variable"`
	Privilege         types.String               `tfsdk:"privilege"`
	PrivilegeVariable types.String               `tfsdk:"privilege_variable"`
	PublicKeys        []SystemAAAUsersPublicKeys `tfsdk:"public_keys"`
}

type SystemAAARadiusGroups struct {
	GroupName               types.String                   `tfsdk:"group_name"`
	Vpn                     types.Int64                    `tfsdk:"vpn"`
	SourceInterface         types.String                   `tfsdk:"source_interface"`
	SourceInterfaceVariable types.String                   `tfsdk:"source_interface_variable"`
	Servers                 []SystemAAARadiusGroupsServers `tfsdk:"servers"`
}

type SystemAAATacacsGroups struct {
	GroupName               types.String                   `tfsdk:"group_name"`
	Vpn                     types.Int64                    `tfsdk:"vpn"`
	SourceInterface         types.String                   `tfsdk:"source_interface"`
	SourceInterfaceVariable types.String                   `tfsdk:"source_interface_variable"`
	Servers                 []SystemAAATacacsGroupsServers `tfsdk:"servers"`
}

type SystemAAAAccountingRules struct {
	RuleId            types.String `tfsdk:"rule_id"`
	Method            types.String `tfsdk:"method"`
	Level             types.String `tfsdk:"level"`
	StartStop         types.Bool   `tfsdk:"start_stop"`
	StartStopVariable types.String `tfsdk:"start_stop_variable"`
	Group             types.Set    `tfsdk:"group"`
}

type SystemAAAAuthorizationRules struct {
	RuleId          types.String `tfsdk:"rule_id"`
	Method          types.String `tfsdk:"method"`
	Level           types.String `tfsdk:"level"`
	Group           types.Set    `tfsdk:"group"`
	IfAuthenticated types.Bool   `tfsdk:"if_authenticated"`
}

type SystemAAAUsersPublicKeys struct {
	KeyString       types.String `tfsdk:"key_string"`
	KeyType         types.String `tfsdk:"key_type"`
	KeyTypeVariable types.String `tfsdk:"key_type_variable"`
}

type SystemAAARadiusGroupsServers struct {
	Address            types.String `tfsdk:"address"`
	AuthPort           types.Int64  `tfsdk:"auth_port"`
	AuthPortVariable   types.String `tfsdk:"auth_port_variable"`
	AcctPort           types.Int64  `tfsdk:"acct_port"`
	AcctPortVariable   types.String `tfsdk:"acct_port_variable"`
	Timeout            types.Int64  `tfsdk:"timeout"`
	TimeoutVariable    types.String `tfsdk:"timeout_variable"`
	Retransmit         types.Int64  `tfsdk:"retransmit"`
	RetransmitVariable types.String `tfsdk:"retransmit_variable"`
	Key                types.String `tfsdk:"key"`
	SecretKey          types.String `tfsdk:"secret_key"`
	SecretKeyVariable  types.String `tfsdk:"secret_key_variable"`
	KeyEnum            types.String `tfsdk:"key_enum"`
	KeyType            types.String `tfsdk:"key_type"`
	KeyTypeVariable    types.String `tfsdk:"key_type_variable"`
}

type SystemAAATacacsGroupsServers struct {
	Address           types.String `tfsdk:"address"`
	Port              types.Int64  `tfsdk:"port"`
	PortVariable      types.String `tfsdk:"port_variable"`
	Timeout           types.Int64  `tfsdk:"timeout"`
	TimeoutVariable   types.String `tfsdk:"timeout_variable"`
	Key               types.String `tfsdk:"key"`
	SecretKey         types.String `tfsdk:"secret_key"`
	SecretKeyVariable types.String `tfsdk:"secret_key_variable"`
	KeyEnum           types.String `tfsdk:"key_enum"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SystemAAA) getModel() string {
	return "system_aaa"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SystemAAA) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/aaa", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SystemAAA) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.AuthenticationGroupVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"authenticationGroup.optionType", "variable")
			body, _ = sjson.Set(body, path+"authenticationGroup.value", data.AuthenticationGroupVariable.ValueString())
		}
	} else if data.AuthenticationGroup.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"authenticationGroup.optionType", "default")
			body, _ = sjson.Set(body, path+"authenticationGroup.value", false)
		}
	} else {
		body, _ = sjson.Set(body, path+"authenticationGroup.optionType", "global")
		body, _ = sjson.Set(body, path+"authenticationGroup.value", data.AuthenticationGroup.ValueBool())
	}

	if !data.AccountingGroupVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"accountingGroup.optionType", "variable")
			body, _ = sjson.Set(body, path+"accountingGroup.value", data.AccountingGroupVariable.ValueString())
		}
	} else if data.AccountingGroup.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"accountingGroup.optionType", "default")
			body, _ = sjson.Set(body, path+"accountingGroup.value", false)
		}
	} else {
		body, _ = sjson.Set(body, path+"accountingGroup.optionType", "global")
		body, _ = sjson.Set(body, path+"accountingGroup.value", data.AccountingGroup.ValueBool())
	}
	if !data.ServerAuthOrder.IsNull() {
		body, _ = sjson.Set(body, path+"serverAuthOrder.optionType", "global")
		var values []string
		data.ServerAuthOrder.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"serverAuthOrder.value", values)
	}
	if true {

		for _, item := range data.Users {
			itemBody := ""

			if !item.NameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.NameVariable.ValueString())
				}
			} else if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
			}

			if !item.PasswordVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "password.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "password.value", item.PasswordVariable.ValueString())
				}
			} else if !item.Password.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "password.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "password.value", item.Password.ValueString())
			}

			if !item.PrivilegeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "privilege.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "privilege.value", item.PrivilegeVariable.ValueString())
				}
			} else if item.Privilege.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "privilege.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "privilege.value", "15")
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "privilege.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "privilege.value", item.Privilege.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "pubkeyChain", []interface{}{})
				for _, childItem := range item.PublicKeys {
					itemChildBody := ""
					if !childItem.KeyString.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "keyString.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "keyString.value", childItem.KeyString.ValueString())
					}

					if !childItem.KeyTypeVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keyType.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "keyType.value", childItem.KeyTypeVariable.ValueString())
						}
					} else if childItem.KeyType.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keyType.optionType", "default")

						}
					} else {
						itemChildBody, _ = sjson.Set(itemChildBody, "keyType.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "keyType.value", childItem.KeyType.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "pubkeyChain.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"user.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"radius", []interface{}{})
		for _, item := range data.RadiusGroups {
			itemBody := ""
			if !item.GroupName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "groupName.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "groupName.value", item.GroupName.ValueString())
			}
			if item.Vpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", 0)
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
			}

			if !item.SourceInterfaceVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.value", item.SourceInterfaceVariable.ValueString())
				}
			} else if item.SourceInterface.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "default")

				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "sourceInterface.value", item.SourceInterface.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "server", []interface{}{})
				for _, childItem := range item.Servers {
					itemChildBody := ""
					if !childItem.Address.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
					}

					if !childItem.AuthPortVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "authPort.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "authPort.value", childItem.AuthPortVariable.ValueString())
						}
					} else if childItem.AuthPort.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "authPort.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "authPort.value", 1812)
						}
					} else {
						itemChildBody, _ = sjson.Set(itemChildBody, "authPort.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "authPort.value", childItem.AuthPort.ValueInt64())
					}

					if !childItem.AcctPortVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "acctPort.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "acctPort.value", childItem.AcctPortVariable.ValueString())
						}
					} else if childItem.AcctPort.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "acctPort.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "acctPort.value", 1813)
						}
					} else {
						itemChildBody, _ = sjson.Set(itemChildBody, "acctPort.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "acctPort.value", childItem.AcctPort.ValueInt64())
					}

					if !childItem.TimeoutVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "timeout.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "timeout.value", childItem.TimeoutVariable.ValueString())
						}
					} else if childItem.Timeout.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "timeout.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "timeout.value", 5)
						}
					} else {
						itemChildBody, _ = sjson.Set(itemChildBody, "timeout.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "timeout.value", childItem.Timeout.ValueInt64())
					}

					if !childItem.RetransmitVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "retransmit.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "retransmit.value", childItem.RetransmitVariable.ValueString())
						}
					} else if childItem.Retransmit.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "retransmit.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "retransmit.value", 3)
						}
					} else {
						itemChildBody, _ = sjson.Set(itemChildBody, "retransmit.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "retransmit.value", childItem.Retransmit.ValueInt64())
					}
					if !childItem.Key.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "key.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "key.value", childItem.Key.ValueString())
					}

					if !childItem.SecretKeyVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "secretKey.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "secretKey.value", childItem.SecretKeyVariable.ValueString())
						}
					} else if childItem.SecretKey.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "secretKey.optionType", "default")

						}
					} else {
						itemChildBody, _ = sjson.Set(itemChildBody, "secretKey.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "secretKey.value", childItem.SecretKey.ValueString())
					}
					if childItem.KeyEnum.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keyEnum.optionType", "default")

						}
					} else {
						itemChildBody, _ = sjson.Set(itemChildBody, "keyEnum.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "keyEnum.value", childItem.KeyEnum.ValueString())
					}

					if !childItem.KeyTypeVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keyType.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "keyType.value", childItem.KeyTypeVariable.ValueString())
						}
					} else if childItem.KeyType.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keyType.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "keyType.value", "key")
						}
					} else {
						itemChildBody, _ = sjson.Set(itemChildBody, "keyType.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "keyType.value", childItem.KeyType.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "server.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"radius.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"tacacs", []interface{}{})
		for _, item := range data.TacacsGroups {
			itemBody := ""
			if !item.GroupName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "groupName.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "groupName.value", item.GroupName.ValueString())
			}
			if item.Vpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", 0)
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
			}

			if !item.SourceInterfaceVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.value", item.SourceInterfaceVariable.ValueString())
				}
			} else if item.SourceInterface.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "default")

				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "sourceInterface.value", item.SourceInterface.ValueString())
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "server", []interface{}{})
				for _, childItem := range item.Servers {
					itemChildBody := ""
					if !childItem.Address.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
					}

					if !childItem.PortVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "port.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "port.value", childItem.PortVariable.ValueString())
						}
					} else if childItem.Port.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "port.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "port.value", 49)
						}
					} else {
						itemChildBody, _ = sjson.Set(itemChildBody, "port.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "port.value", childItem.Port.ValueInt64())
					}

					if !childItem.TimeoutVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "timeout.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "timeout.value", childItem.TimeoutVariable.ValueString())
						}
					} else if childItem.Timeout.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "timeout.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "timeout.value", 5)
						}
					} else {
						itemChildBody, _ = sjson.Set(itemChildBody, "timeout.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "timeout.value", childItem.Timeout.ValueInt64())
					}
					if !childItem.Key.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "key.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "key.value", childItem.Key.ValueString())
					}

					if !childItem.SecretKeyVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "secretKey.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "secretKey.value", childItem.SecretKeyVariable.ValueString())
						}
					} else if !childItem.SecretKey.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "secretKey.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "secretKey.value", childItem.SecretKey.ValueString())
					}
					if childItem.KeyEnum.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keyEnum.optionType", "default")

						}
					} else {
						itemChildBody, _ = sjson.Set(itemChildBody, "keyEnum.optionType", "global")
						itemChildBody, _ = sjson.Set(itemChildBody, "keyEnum.value", childItem.KeyEnum.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "server.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"tacacs.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"accountingRule", []interface{}{})
		for _, item := range data.AccountingRules {
			itemBody := ""
			if !item.RuleId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ruleId.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "ruleId.value", item.RuleId.ValueString())
			}
			if !item.Method.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "method.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "method.value", item.Method.ValueString())
			}
			if item.Level.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "level.optionType", "default")

				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "level.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "level.value", item.Level.ValueString())
			}

			if !item.StartStopVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "startStop.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "startStop.value", item.StartStopVariable.ValueString())
				}
			} else if item.StartStop.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "startStop.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "startStop.value", true)
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "startStop.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "startStop.value", item.StartStop.ValueBool())
			}
			if !item.Group.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "group.optionType", "global")
				var values []string
				item.Group.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "group.value", values)
			}
			body, _ = sjson.SetRaw(body, path+"accountingRule.-1", itemBody)
		}
	}

	if !data.AuthorizationConsoleVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"authorizationConsole.optionType", "variable")
			body, _ = sjson.Set(body, path+"authorizationConsole.value", data.AuthorizationConsoleVariable.ValueString())
		}
	} else if data.AuthorizationConsole.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"authorizationConsole.optionType", "default")
			body, _ = sjson.Set(body, path+"authorizationConsole.value", false)
		}
	} else {
		body, _ = sjson.Set(body, path+"authorizationConsole.optionType", "global")
		body, _ = sjson.Set(body, path+"authorizationConsole.value", data.AuthorizationConsole.ValueBool())
	}

	if !data.AuthorizationConfigCommandsVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"authorizationConfigCommands.optionType", "variable")
			body, _ = sjson.Set(body, path+"authorizationConfigCommands.value", data.AuthorizationConfigCommandsVariable.ValueString())
		}
	} else if data.AuthorizationConfigCommands.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"authorizationConfigCommands.optionType", "default")
			body, _ = sjson.Set(body, path+"authorizationConfigCommands.value", false)
		}
	} else {
		body, _ = sjson.Set(body, path+"authorizationConfigCommands.optionType", "global")
		body, _ = sjson.Set(body, path+"authorizationConfigCommands.value", data.AuthorizationConfigCommands.ValueBool())
	}
	if true {
		body, _ = sjson.Set(body, path+"authorizationRule", []interface{}{})
		for _, item := range data.AuthorizationRules {
			itemBody := ""
			if !item.RuleId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ruleId.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "ruleId.value", item.RuleId.ValueString())
			}
			if !item.Method.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "method.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "method.value", item.Method.ValueString())
			}
			if !item.Level.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "level.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "level.value", item.Level.ValueString())
			}
			if !item.Group.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "group.optionType", "global")
				var values []string
				item.Group.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "group.value", values)
			}
			if item.IfAuthenticated.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifAuthenticated.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "ifAuthenticated.value", false)
				}
			} else {
				itemBody, _ = sjson.Set(itemBody, "ifAuthenticated.optionType", "global")
				itemBody, _ = sjson.Set(itemBody, "ifAuthenticated.value", item.IfAuthenticated.ValueBool())
			}
			body, _ = sjson.SetRaw(body, path+"authorizationRule.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemAAA) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.AuthenticationGroup = types.BoolNull()
	data.AuthenticationGroupVariable = types.StringNull()
	if t := res.Get(path + "authenticationGroup.optionType"); t.Exists() {
		va := res.Get(path + "authenticationGroup.value")
		if t.String() == "variable" {
			data.AuthenticationGroupVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AuthenticationGroup = types.BoolValue(va.Bool())
		}
	}
	data.AccountingGroup = types.BoolNull()
	data.AccountingGroupVariable = types.StringNull()
	if t := res.Get(path + "accountingGroup.optionType"); t.Exists() {
		va := res.Get(path + "accountingGroup.value")
		if t.String() == "variable" {
			data.AccountingGroupVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AccountingGroup = types.BoolValue(va.Bool())
		}
	}
	data.ServerAuthOrder = types.SetNull(types.StringType)

	if t := res.Get(path + "serverAuthOrder.optionType"); t.Exists() {
		va := res.Get(path + "serverAuthOrder.value")
		if t.String() == "global" {
			data.ServerAuthOrder = helpers.GetStringSet(va.Array())
		}
	}
	if value := res.Get(path + "user"); value.Exists() {
		data.Users = make([]SystemAAAUsers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemAAAUsers{}
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
			item.Privilege = types.StringNull()
			item.PrivilegeVariable = types.StringNull()
			if t := v.Get("privilege.optionType"); t.Exists() {
				va := v.Get("privilege.value")
				if t.String() == "variable" {
					item.PrivilegeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Privilege = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("pubkeyChain"); cValue.Exists() {
				item.PublicKeys = make([]SystemAAAUsersPublicKeys, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := SystemAAAUsersPublicKeys{}
					cItem.KeyString = types.StringNull()

					if t := cv.Get("keyString.optionType"); t.Exists() {
						va := cv.Get("keyString.value")
						if t.String() == "global" {
							cItem.KeyString = types.StringValue(va.String())
						}
					}
					cItem.KeyType = types.StringNull()
					cItem.KeyTypeVariable = types.StringNull()
					if t := cv.Get("keyType.optionType"); t.Exists() {
						va := cv.Get("keyType.value")
						if t.String() == "variable" {
							cItem.KeyTypeVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.KeyType = types.StringValue(va.String())
						}
					}
					item.PublicKeys = append(item.PublicKeys, cItem)
					return true
				})
			}
			data.Users = append(data.Users, item)
			return true
		})
	}
	if value := res.Get(path + "radius"); value.Exists() {
		data.RadiusGroups = make([]SystemAAARadiusGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemAAARadiusGroups{}
			item.GroupName = types.StringNull()

			if t := v.Get("groupName.optionType"); t.Exists() {
				va := v.Get("groupName.value")
				if t.String() == "global" {
					item.GroupName = types.StringValue(va.String())
				}
			}
			item.Vpn = types.Int64Null()

			if t := v.Get("vpn.optionType"); t.Exists() {
				va := v.Get("vpn.value")
				if t.String() == "global" {
					item.Vpn = types.Int64Value(va.Int())
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
			if cValue := v.Get("server"); cValue.Exists() {
				item.Servers = make([]SystemAAARadiusGroupsServers, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := SystemAAARadiusGroupsServers{}
					cItem.Address = types.StringNull()

					if t := cv.Get("address.optionType"); t.Exists() {
						va := cv.Get("address.value")
						if t.String() == "global" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					cItem.AuthPort = types.Int64Null()
					cItem.AuthPortVariable = types.StringNull()
					if t := cv.Get("authPort.optionType"); t.Exists() {
						va := cv.Get("authPort.value")
						if t.String() == "variable" {
							cItem.AuthPortVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.AuthPort = types.Int64Value(va.Int())
						}
					}
					cItem.AcctPort = types.Int64Null()
					cItem.AcctPortVariable = types.StringNull()
					if t := cv.Get("acctPort.optionType"); t.Exists() {
						va := cv.Get("acctPort.value")
						if t.String() == "variable" {
							cItem.AcctPortVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.AcctPort = types.Int64Value(va.Int())
						}
					}
					cItem.Timeout = types.Int64Null()
					cItem.TimeoutVariable = types.StringNull()
					if t := cv.Get("timeout.optionType"); t.Exists() {
						va := cv.Get("timeout.value")
						if t.String() == "variable" {
							cItem.TimeoutVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Timeout = types.Int64Value(va.Int())
						}
					}
					cItem.Retransmit = types.Int64Null()
					cItem.RetransmitVariable = types.StringNull()
					if t := cv.Get("retransmit.optionType"); t.Exists() {
						va := cv.Get("retransmit.value")
						if t.String() == "variable" {
							cItem.RetransmitVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Retransmit = types.Int64Value(va.Int())
						}
					}
					cItem.SecretKey = types.StringNull()
					cItem.SecretKeyVariable = types.StringNull()
					if t := cv.Get("secretKey.optionType"); t.Exists() {
						va := cv.Get("secretKey.value")
						if t.String() == "variable" {
							cItem.SecretKeyVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.SecretKey = types.StringValue(va.String())
						}
					}
					cItem.KeyEnum = types.StringNull()

					if t := cv.Get("keyEnum.optionType"); t.Exists() {
						va := cv.Get("keyEnum.value")
						if t.String() == "global" {
							cItem.KeyEnum = types.StringValue(va.String())
						}
					}
					cItem.KeyType = types.StringNull()
					cItem.KeyTypeVariable = types.StringNull()
					if t := cv.Get("keyType.optionType"); t.Exists() {
						va := cv.Get("keyType.value")
						if t.String() == "variable" {
							cItem.KeyTypeVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.KeyType = types.StringValue(va.String())
						}
					}
					item.Servers = append(item.Servers, cItem)
					return true
				})
			}
			data.RadiusGroups = append(data.RadiusGroups, item)
			return true
		})
	}
	if value := res.Get(path + "tacacs"); value.Exists() {
		data.TacacsGroups = make([]SystemAAATacacsGroups, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemAAATacacsGroups{}
			item.GroupName = types.StringNull()

			if t := v.Get("groupName.optionType"); t.Exists() {
				va := v.Get("groupName.value")
				if t.String() == "global" {
					item.GroupName = types.StringValue(va.String())
				}
			}
			item.Vpn = types.Int64Null()

			if t := v.Get("vpn.optionType"); t.Exists() {
				va := v.Get("vpn.value")
				if t.String() == "global" {
					item.Vpn = types.Int64Value(va.Int())
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
			if cValue := v.Get("server"); cValue.Exists() {
				item.Servers = make([]SystemAAATacacsGroupsServers, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := SystemAAATacacsGroupsServers{}
					cItem.Address = types.StringNull()

					if t := cv.Get("address.optionType"); t.Exists() {
						va := cv.Get("address.value")
						if t.String() == "global" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					cItem.Port = types.Int64Null()
					cItem.PortVariable = types.StringNull()
					if t := cv.Get("port.optionType"); t.Exists() {
						va := cv.Get("port.value")
						if t.String() == "variable" {
							cItem.PortVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Port = types.Int64Value(va.Int())
						}
					}
					cItem.Timeout = types.Int64Null()
					cItem.TimeoutVariable = types.StringNull()
					if t := cv.Get("timeout.optionType"); t.Exists() {
						va := cv.Get("timeout.value")
						if t.String() == "variable" {
							cItem.TimeoutVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Timeout = types.Int64Value(va.Int())
						}
					}
					cItem.SecretKey = types.StringNull()
					cItem.SecretKeyVariable = types.StringNull()
					if t := cv.Get("secretKey.optionType"); t.Exists() {
						va := cv.Get("secretKey.value")
						if t.String() == "variable" {
							cItem.SecretKeyVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.SecretKey = types.StringValue(va.String())
						}
					}
					cItem.KeyEnum = types.StringNull()

					if t := cv.Get("keyEnum.optionType"); t.Exists() {
						va := cv.Get("keyEnum.value")
						if t.String() == "global" {
							cItem.KeyEnum = types.StringValue(va.String())
						}
					}
					item.Servers = append(item.Servers, cItem)
					return true
				})
			}
			data.TacacsGroups = append(data.TacacsGroups, item)
			return true
		})
	}
	if value := res.Get(path + "accountingRule"); value.Exists() {
		data.AccountingRules = make([]SystemAAAAccountingRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemAAAAccountingRules{}
			item.RuleId = types.StringNull()

			if t := v.Get("ruleId.optionType"); t.Exists() {
				va := v.Get("ruleId.value")
				if t.String() == "global" {
					item.RuleId = types.StringValue(va.String())
				}
			}
			item.Method = types.StringNull()

			if t := v.Get("method.optionType"); t.Exists() {
				va := v.Get("method.value")
				if t.String() == "global" {
					item.Method = types.StringValue(va.String())
				}
			}
			item.Level = types.StringNull()

			if t := v.Get("level.optionType"); t.Exists() {
				va := v.Get("level.value")
				if t.String() == "global" {
					item.Level = types.StringValue(va.String())
				}
			}
			item.StartStop = types.BoolNull()
			item.StartStopVariable = types.StringNull()
			if t := v.Get("startStop.optionType"); t.Exists() {
				va := v.Get("startStop.value")
				if t.String() == "variable" {
					item.StartStopVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.StartStop = types.BoolValue(va.Bool())
				}
			}
			item.Group = types.SetNull(types.StringType)

			if t := v.Get("group.optionType"); t.Exists() {
				va := v.Get("group.value")
				if t.String() == "global" {
					item.Group = helpers.GetStringSet(va.Array())
				}
			}
			data.AccountingRules = append(data.AccountingRules, item)
			return true
		})
	}
	data.AuthorizationConsole = types.BoolNull()
	data.AuthorizationConsoleVariable = types.StringNull()
	if t := res.Get(path + "authorizationConsole.optionType"); t.Exists() {
		va := res.Get(path + "authorizationConsole.value")
		if t.String() == "variable" {
			data.AuthorizationConsoleVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AuthorizationConsole = types.BoolValue(va.Bool())
		}
	}
	data.AuthorizationConfigCommands = types.BoolNull()
	data.AuthorizationConfigCommandsVariable = types.StringNull()
	if t := res.Get(path + "authorizationConfigCommands.optionType"); t.Exists() {
		va := res.Get(path + "authorizationConfigCommands.value")
		if t.String() == "variable" {
			data.AuthorizationConfigCommandsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AuthorizationConfigCommands = types.BoolValue(va.Bool())
		}
	}
	if value := res.Get(path + "authorizationRule"); value.Exists() {
		data.AuthorizationRules = make([]SystemAAAAuthorizationRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemAAAAuthorizationRules{}
			item.RuleId = types.StringNull()

			if t := v.Get("ruleId.optionType"); t.Exists() {
				va := v.Get("ruleId.value")
				if t.String() == "global" {
					item.RuleId = types.StringValue(va.String())
				}
			}
			item.Method = types.StringNull()

			if t := v.Get("method.optionType"); t.Exists() {
				va := v.Get("method.value")
				if t.String() == "global" {
					item.Method = types.StringValue(va.String())
				}
			}
			item.Level = types.StringNull()

			if t := v.Get("level.optionType"); t.Exists() {
				va := v.Get("level.value")
				if t.String() == "global" {
					item.Level = types.StringValue(va.String())
				}
			}
			item.Group = types.SetNull(types.StringType)

			if t := v.Get("group.optionType"); t.Exists() {
				va := v.Get("group.value")
				if t.String() == "global" {
					item.Group = helpers.GetStringSet(va.Array())
				}
			}
			item.IfAuthenticated = types.BoolNull()

			if t := v.Get("ifAuthenticated.optionType"); t.Exists() {
				va := v.Get("ifAuthenticated.value")
				if t.String() == "global" {
					item.IfAuthenticated = types.BoolValue(va.Bool())
				}
			}
			data.AuthorizationRules = append(data.AuthorizationRules, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SystemAAA) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.AuthenticationGroup = types.BoolNull()
	data.AuthenticationGroupVariable = types.StringNull()
	if t := res.Get(path + "authenticationGroup.optionType"); t.Exists() {
		va := res.Get(path + "authenticationGroup.value")
		if t.String() == "variable" {
			data.AuthenticationGroupVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AuthenticationGroup = types.BoolValue(va.Bool())
		}
	}
	data.AccountingGroup = types.BoolNull()
	data.AccountingGroupVariable = types.StringNull()
	if t := res.Get(path + "accountingGroup.optionType"); t.Exists() {
		va := res.Get(path + "accountingGroup.value")
		if t.String() == "variable" {
			data.AccountingGroupVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AccountingGroup = types.BoolValue(va.Bool())
		}
	}
	data.ServerAuthOrder = types.SetNull(types.StringType)

	if t := res.Get(path + "serverAuthOrder.optionType"); t.Exists() {
		va := res.Get(path + "serverAuthOrder.value")
		if t.String() == "global" {
			data.ServerAuthOrder = helpers.GetStringSet(va.Array())
		}
	}
	for i := range data.Users {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Users[i].Name.ValueString()}
		keyValuesVariables := [...]string{data.Users[i].NameVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "user").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Users[i].Name = types.StringNull()
		data.Users[i].NameVariable = types.StringNull()
		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "variable" {
				data.Users[i].NameVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Users[i].Name = types.StringValue(va.String())
			}
		}
		data.Users[i].Privilege = types.StringNull()
		data.Users[i].PrivilegeVariable = types.StringNull()
		if t := r.Get("privilege.optionType"); t.Exists() {
			va := r.Get("privilege.value")
			if t.String() == "variable" {
				data.Users[i].PrivilegeVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Users[i].Privilege = types.StringValue(va.String())
			}
		}
		for ci := range data.Users[i].PublicKeys {
			keys := [...]string{"keyString"}
			keyValues := [...]string{data.Users[i].PublicKeys[ci].KeyString.ValueString()}
			keyValuesVariables := [...]string{""}

			var cr gjson.Result
			r.Get("pubkeyChain").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType").String()
						vv := v.Get(keys[ik] + ".value").String()
						if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			data.Users[i].PublicKeys[ci].KeyString = types.StringNull()

			if t := cr.Get("keyString.optionType"); t.Exists() {
				va := cr.Get("keyString.value")
				if t.String() == "global" {
					data.Users[i].PublicKeys[ci].KeyString = types.StringValue(va.String())
				}
			}
			data.Users[i].PublicKeys[ci].KeyType = types.StringNull()
			data.Users[i].PublicKeys[ci].KeyTypeVariable = types.StringNull()
			if t := cr.Get("keyType.optionType"); t.Exists() {
				va := cr.Get("keyType.value")
				if t.String() == "variable" {
					data.Users[i].PublicKeys[ci].KeyTypeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Users[i].PublicKeys[ci].KeyType = types.StringValue(va.String())
				}
			}
		}
	}
	for i := range data.RadiusGroups {
		keys := [...]string{"groupName"}
		keyValues := [...]string{data.RadiusGroups[i].GroupName.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "radius").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.RadiusGroups[i].GroupName = types.StringNull()

		if t := r.Get("groupName.optionType"); t.Exists() {
			va := r.Get("groupName.value")
			if t.String() == "global" {
				data.RadiusGroups[i].GroupName = types.StringValue(va.String())
			}
		}
		data.RadiusGroups[i].Vpn = types.Int64Null()

		if t := r.Get("vpn.optionType"); t.Exists() {
			va := r.Get("vpn.value")
			if t.String() == "global" {
				data.RadiusGroups[i].Vpn = types.Int64Value(va.Int())
			}
		}
		data.RadiusGroups[i].SourceInterface = types.StringNull()
		data.RadiusGroups[i].SourceInterfaceVariable = types.StringNull()
		if t := r.Get("sourceInterface.optionType"); t.Exists() {
			va := r.Get("sourceInterface.value")
			if t.String() == "variable" {
				data.RadiusGroups[i].SourceInterfaceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.RadiusGroups[i].SourceInterface = types.StringValue(va.String())
			}
		}
		for ci := range data.RadiusGroups[i].Servers {
			keys := [...]string{"address"}
			keyValues := [...]string{data.RadiusGroups[i].Servers[ci].Address.ValueString()}
			keyValuesVariables := [...]string{""}

			var cr gjson.Result
			r.Get("server").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType").String()
						vv := v.Get(keys[ik] + ".value").String()
						if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			data.RadiusGroups[i].Servers[ci].Address = types.StringNull()

			if t := cr.Get("address.optionType"); t.Exists() {
				va := cr.Get("address.value")
				if t.String() == "global" {
					data.RadiusGroups[i].Servers[ci].Address = types.StringValue(va.String())
				}
			}
			data.RadiusGroups[i].Servers[ci].AuthPort = types.Int64Null()
			data.RadiusGroups[i].Servers[ci].AuthPortVariable = types.StringNull()
			if t := cr.Get("authPort.optionType"); t.Exists() {
				va := cr.Get("authPort.value")
				if t.String() == "variable" {
					data.RadiusGroups[i].Servers[ci].AuthPortVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.RadiusGroups[i].Servers[ci].AuthPort = types.Int64Value(va.Int())
				}
			}
			data.RadiusGroups[i].Servers[ci].AcctPort = types.Int64Null()
			data.RadiusGroups[i].Servers[ci].AcctPortVariable = types.StringNull()
			if t := cr.Get("acctPort.optionType"); t.Exists() {
				va := cr.Get("acctPort.value")
				if t.String() == "variable" {
					data.RadiusGroups[i].Servers[ci].AcctPortVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.RadiusGroups[i].Servers[ci].AcctPort = types.Int64Value(va.Int())
				}
			}
			data.RadiusGroups[i].Servers[ci].Timeout = types.Int64Null()
			data.RadiusGroups[i].Servers[ci].TimeoutVariable = types.StringNull()
			if t := cr.Get("timeout.optionType"); t.Exists() {
				va := cr.Get("timeout.value")
				if t.String() == "variable" {
					data.RadiusGroups[i].Servers[ci].TimeoutVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.RadiusGroups[i].Servers[ci].Timeout = types.Int64Value(va.Int())
				}
			}
			data.RadiusGroups[i].Servers[ci].Retransmit = types.Int64Null()
			data.RadiusGroups[i].Servers[ci].RetransmitVariable = types.StringNull()
			if t := cr.Get("retransmit.optionType"); t.Exists() {
				va := cr.Get("retransmit.value")
				if t.String() == "variable" {
					data.RadiusGroups[i].Servers[ci].RetransmitVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.RadiusGroups[i].Servers[ci].Retransmit = types.Int64Value(va.Int())
				}
			}
			data.RadiusGroups[i].Servers[ci].SecretKey = types.StringNull()
			data.RadiusGroups[i].Servers[ci].SecretKeyVariable = types.StringNull()
			if t := cr.Get("secretKey.optionType"); t.Exists() {
				va := cr.Get("secretKey.value")
				if t.String() == "variable" {
					data.RadiusGroups[i].Servers[ci].SecretKeyVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.RadiusGroups[i].Servers[ci].SecretKey = types.StringValue(va.String())
				}
			}
			data.RadiusGroups[i].Servers[ci].KeyEnum = types.StringNull()

			if t := cr.Get("keyEnum.optionType"); t.Exists() {
				va := cr.Get("keyEnum.value")
				if t.String() == "global" {
					data.RadiusGroups[i].Servers[ci].KeyEnum = types.StringValue(va.String())
				}
			}
			data.RadiusGroups[i].Servers[ci].KeyType = types.StringNull()
			data.RadiusGroups[i].Servers[ci].KeyTypeVariable = types.StringNull()
			if t := cr.Get("keyType.optionType"); t.Exists() {
				va := cr.Get("keyType.value")
				if t.String() == "variable" {
					data.RadiusGroups[i].Servers[ci].KeyTypeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.RadiusGroups[i].Servers[ci].KeyType = types.StringValue(va.String())
				}
			}
		}
	}
	for i := range data.TacacsGroups {
		keys := [...]string{"groupName"}
		keyValues := [...]string{data.TacacsGroups[i].GroupName.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "tacacs").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.TacacsGroups[i].GroupName = types.StringNull()

		if t := r.Get("groupName.optionType"); t.Exists() {
			va := r.Get("groupName.value")
			if t.String() == "global" {
				data.TacacsGroups[i].GroupName = types.StringValue(va.String())
			}
		}
		data.TacacsGroups[i].Vpn = types.Int64Null()

		if t := r.Get("vpn.optionType"); t.Exists() {
			va := r.Get("vpn.value")
			if t.String() == "global" {
				data.TacacsGroups[i].Vpn = types.Int64Value(va.Int())
			}
		}
		data.TacacsGroups[i].SourceInterface = types.StringNull()
		data.TacacsGroups[i].SourceInterfaceVariable = types.StringNull()
		if t := r.Get("sourceInterface.optionType"); t.Exists() {
			va := r.Get("sourceInterface.value")
			if t.String() == "variable" {
				data.TacacsGroups[i].SourceInterfaceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.TacacsGroups[i].SourceInterface = types.StringValue(va.String())
			}
		}
		for ci := range data.TacacsGroups[i].Servers {
			keys := [...]string{"address"}
			keyValues := [...]string{data.TacacsGroups[i].Servers[ci].Address.ValueString()}
			keyValuesVariables := [...]string{""}

			var cr gjson.Result
			r.Get("server").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType").String()
						vv := v.Get(keys[ik] + ".value").String()
						if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			data.TacacsGroups[i].Servers[ci].Address = types.StringNull()

			if t := cr.Get("address.optionType"); t.Exists() {
				va := cr.Get("address.value")
				if t.String() == "global" {
					data.TacacsGroups[i].Servers[ci].Address = types.StringValue(va.String())
				}
			}
			data.TacacsGroups[i].Servers[ci].Port = types.Int64Null()
			data.TacacsGroups[i].Servers[ci].PortVariable = types.StringNull()
			if t := cr.Get("port.optionType"); t.Exists() {
				va := cr.Get("port.value")
				if t.String() == "variable" {
					data.TacacsGroups[i].Servers[ci].PortVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.TacacsGroups[i].Servers[ci].Port = types.Int64Value(va.Int())
				}
			}
			data.TacacsGroups[i].Servers[ci].Timeout = types.Int64Null()
			data.TacacsGroups[i].Servers[ci].TimeoutVariable = types.StringNull()
			if t := cr.Get("timeout.optionType"); t.Exists() {
				va := cr.Get("timeout.value")
				if t.String() == "variable" {
					data.TacacsGroups[i].Servers[ci].TimeoutVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.TacacsGroups[i].Servers[ci].Timeout = types.Int64Value(va.Int())
				}
			}
			data.TacacsGroups[i].Servers[ci].SecretKey = types.StringNull()
			data.TacacsGroups[i].Servers[ci].SecretKeyVariable = types.StringNull()
			if t := cr.Get("secretKey.optionType"); t.Exists() {
				va := cr.Get("secretKey.value")
				if t.String() == "variable" {
					data.TacacsGroups[i].Servers[ci].SecretKeyVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.TacacsGroups[i].Servers[ci].SecretKey = types.StringValue(va.String())
				}
			}
			data.TacacsGroups[i].Servers[ci].KeyEnum = types.StringNull()

			if t := cr.Get("keyEnum.optionType"); t.Exists() {
				va := cr.Get("keyEnum.value")
				if t.String() == "global" {
					data.TacacsGroups[i].Servers[ci].KeyEnum = types.StringValue(va.String())
				}
			}
		}
	}
	for i := range data.AccountingRules {
		keys := [...]string{"ruleId"}
		keyValues := [...]string{data.AccountingRules[i].RuleId.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "accountingRule").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.AccountingRules[i].RuleId = types.StringNull()

		if t := r.Get("ruleId.optionType"); t.Exists() {
			va := r.Get("ruleId.value")
			if t.String() == "global" {
				data.AccountingRules[i].RuleId = types.StringValue(va.String())
			}
		}
		data.AccountingRules[i].Method = types.StringNull()

		if t := r.Get("method.optionType"); t.Exists() {
			va := r.Get("method.value")
			if t.String() == "global" {
				data.AccountingRules[i].Method = types.StringValue(va.String())
			}
		}
		data.AccountingRules[i].Level = types.StringNull()

		if t := r.Get("level.optionType"); t.Exists() {
			va := r.Get("level.value")
			if t.String() == "global" {
				data.AccountingRules[i].Level = types.StringValue(va.String())
			}
		}
		data.AccountingRules[i].StartStop = types.BoolNull()
		data.AccountingRules[i].StartStopVariable = types.StringNull()
		if t := r.Get("startStop.optionType"); t.Exists() {
			va := r.Get("startStop.value")
			if t.String() == "variable" {
				data.AccountingRules[i].StartStopVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.AccountingRules[i].StartStop = types.BoolValue(va.Bool())
			}
		}
		data.AccountingRules[i].Group = types.SetNull(types.StringType)

		if t := r.Get("group.optionType"); t.Exists() {
			va := r.Get("group.value")
			if t.String() == "global" {
				data.AccountingRules[i].Group = helpers.GetStringSet(va.Array())
			}
		}
	}
	data.AuthorizationConsole = types.BoolNull()
	data.AuthorizationConsoleVariable = types.StringNull()
	if t := res.Get(path + "authorizationConsole.optionType"); t.Exists() {
		va := res.Get(path + "authorizationConsole.value")
		if t.String() == "variable" {
			data.AuthorizationConsoleVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AuthorizationConsole = types.BoolValue(va.Bool())
		}
	}
	data.AuthorizationConfigCommands = types.BoolNull()
	data.AuthorizationConfigCommandsVariable = types.StringNull()
	if t := res.Get(path + "authorizationConfigCommands.optionType"); t.Exists() {
		va := res.Get(path + "authorizationConfigCommands.value")
		if t.String() == "variable" {
			data.AuthorizationConfigCommandsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AuthorizationConfigCommands = types.BoolValue(va.Bool())
		}
	}
	for i := range data.AuthorizationRules {
		keys := [...]string{"ruleId"}
		keyValues := [...]string{data.AuthorizationRules[i].RuleId.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "authorizationRule").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.AuthorizationRules[i].RuleId = types.StringNull()

		if t := r.Get("ruleId.optionType"); t.Exists() {
			va := r.Get("ruleId.value")
			if t.String() == "global" {
				data.AuthorizationRules[i].RuleId = types.StringValue(va.String())
			}
		}
		data.AuthorizationRules[i].Method = types.StringNull()

		if t := r.Get("method.optionType"); t.Exists() {
			va := r.Get("method.value")
			if t.String() == "global" {
				data.AuthorizationRules[i].Method = types.StringValue(va.String())
			}
		}
		data.AuthorizationRules[i].Level = types.StringNull()

		if t := r.Get("level.optionType"); t.Exists() {
			va := r.Get("level.value")
			if t.String() == "global" {
				data.AuthorizationRules[i].Level = types.StringValue(va.String())
			}
		}
		data.AuthorizationRules[i].Group = types.SetNull(types.StringType)

		if t := r.Get("group.optionType"); t.Exists() {
			va := r.Get("group.value")
			if t.String() == "global" {
				data.AuthorizationRules[i].Group = helpers.GetStringSet(va.Array())
			}
		}
		data.AuthorizationRules[i].IfAuthenticated = types.BoolNull()

		if t := r.Get("ifAuthenticated.optionType"); t.Exists() {
			va := r.Get("ifAuthenticated.value")
			if t.String() == "global" {
				data.AuthorizationRules[i].IfAuthenticated = types.BoolValue(va.Bool())
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *SystemAAA) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.AuthenticationGroup.IsNull() {
		return false
	}
	if !data.AuthenticationGroupVariable.IsNull() {
		return false
	}
	if !data.AccountingGroup.IsNull() {
		return false
	}
	if !data.AccountingGroupVariable.IsNull() {
		return false
	}
	if !data.ServerAuthOrder.IsNull() {
		return false
	}
	if len(data.Users) > 0 {
		return false
	}
	if len(data.RadiusGroups) > 0 {
		return false
	}
	if len(data.TacacsGroups) > 0 {
		return false
	}
	if len(data.AccountingRules) > 0 {
		return false
	}
	if !data.AuthorizationConsole.IsNull() {
		return false
	}
	if !data.AuthorizationConsoleVariable.IsNull() {
		return false
	}
	if !data.AuthorizationConfigCommands.IsNull() {
		return false
	}
	if !data.AuthorizationConfigCommandsVariable.IsNull() {
		return false
	}
	if len(data.AuthorizationRules) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
