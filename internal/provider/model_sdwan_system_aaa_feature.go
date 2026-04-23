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
	ServerAuthOrder                     types.List                    `tfsdk:"server_auth_order"`
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
	Group             types.List   `tfsdk:"group"`
}

type SystemAAAAuthorizationRules struct {
	RuleId          types.String `tfsdk:"rule_id"`
	Method          types.String `tfsdk:"method"`
	Level           types.String `tfsdk:"level"`
	Group           types.List   `tfsdk:"group"`
	IfAuthenticated types.Bool   `tfsdk:"if_authenticated"`
}

type SystemAAAUsersPublicKeys struct {
	KeyString         types.String `tfsdk:"key_string"`
	KeyStringVariable types.String `tfsdk:"key_string_variable"`
	KeyType           types.String `tfsdk:"key_type"`
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
	KeyVariable        types.String `tfsdk:"key_variable"`
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
	KeyVariable       types.String `tfsdk:"key_variable"`
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
		if true {
			body, _ = sjson.Set(body, path+"authenticationGroup.optionType", "global")
			body, _ = sjson.Set(body, path+"authenticationGroup.value", data.AuthenticationGroup.ValueBool())
		}
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
		if true {
			body, _ = sjson.Set(body, path+"accountingGroup.optionType", "global")
			body, _ = sjson.Set(body, path+"accountingGroup.value", data.AccountingGroup.ValueBool())
		}
	}
	if !data.ServerAuthOrder.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"serverAuthOrder.optionType", "global")
			var values []string
			data.ServerAuthOrder.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"serverAuthOrder.value", values)
		}
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
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
				}
			}

			if !item.PasswordVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "password.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "password.value", item.PasswordVariable.ValueString())
				}
			} else if !item.Password.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "password.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "password.value", item.Password.ValueString())
				}
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
				if true {
					itemBody, _ = sjson.Set(itemBody, "privilege.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "privilege.value", item.Privilege.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "pubkeyChain", []interface{}{})
				for _, childItem := range item.PublicKeys {
					itemChildBody := ""

					if !childItem.KeyStringVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keyString.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "keyString.value", childItem.KeyStringVariable.ValueString())
						}
					} else if !childItem.KeyString.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keyString.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "keyString.value", childItem.KeyString.ValueString())
						}
					}
					if !childItem.KeyType.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keyType.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "keyType.value", childItem.KeyType.ValueString())
						}
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
				if true {
					itemBody, _ = sjson.Set(itemBody, "groupName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "groupName.value", item.GroupName.ValueString())
				}
			}
			if item.Vpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", 0)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
				}
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
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.value", item.SourceInterface.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "server", []interface{}{})
				for _, childItem := range item.Servers {
					itemChildBody := ""
					if !childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
						}
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
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "authPort.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "authPort.value", childItem.AuthPort.ValueInt64())
						}
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
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "acctPort.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "acctPort.value", childItem.AcctPort.ValueInt64())
						}
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
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "timeout.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "timeout.value", childItem.Timeout.ValueInt64())
						}
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
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "retransmit.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "retransmit.value", childItem.Retransmit.ValueInt64())
						}
					}

					if !childItem.KeyVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "key.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "key.value", childItem.KeyVariable.ValueString())
						}
					} else if !childItem.Key.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "key.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "key.value", childItem.Key.ValueString())
						}
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
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "secretKey.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "secretKey.value", childItem.SecretKey.ValueString())
						}
					}
					if childItem.KeyEnum.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keyEnum.optionType", "default")

						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keyEnum.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "keyEnum.value", childItem.KeyEnum.ValueString())
						}
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
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keyType.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "keyType.value", childItem.KeyType.ValueString())
						}
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
				if true {
					itemBody, _ = sjson.Set(itemBody, "groupName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "groupName.value", item.GroupName.ValueString())
				}
			}
			if item.Vpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", 0)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
				}
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
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sourceInterface.value", item.SourceInterface.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "server", []interface{}{})
				for _, childItem := range item.Servers {
					itemChildBody := ""
					if !childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
						}
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
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "port.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "port.value", childItem.Port.ValueInt64())
						}
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
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "timeout.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "timeout.value", childItem.Timeout.ValueInt64())
						}
					}

					if !childItem.KeyVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "key.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "key.value", childItem.KeyVariable.ValueString())
						}
					} else if !childItem.Key.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "key.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "key.value", childItem.Key.ValueString())
						}
					}

					if !childItem.SecretKeyVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "secretKey.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "secretKey.value", childItem.SecretKeyVariable.ValueString())
						}
					} else if !childItem.SecretKey.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "secretKey.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "secretKey.value", childItem.SecretKey.ValueString())
						}
					}
					if childItem.KeyEnum.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keyEnum.optionType", "default")

						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "keyEnum.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "keyEnum.value", childItem.KeyEnum.ValueString())
						}
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
				if true {
					itemBody, _ = sjson.Set(itemBody, "ruleId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ruleId.value", item.RuleId.ValueString())
				}
			}
			if !item.Method.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "method.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "method.value", item.Method.ValueString())
				}
			}
			if item.Level.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "level.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "level.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "level.value", item.Level.ValueString())
				}
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
				if true {
					itemBody, _ = sjson.Set(itemBody, "startStop.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "startStop.value", item.StartStop.ValueBool())
				}
			}
			if !item.Group.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "group.optionType", "global")
					var values []string
					item.Group.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "group.value", values)
				}
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
		if true {
			body, _ = sjson.Set(body, path+"authorizationConsole.optionType", "global")
			body, _ = sjson.Set(body, path+"authorizationConsole.value", data.AuthorizationConsole.ValueBool())
		}
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
		if true {
			body, _ = sjson.Set(body, path+"authorizationConfigCommands.optionType", "global")
			body, _ = sjson.Set(body, path+"authorizationConfigCommands.value", data.AuthorizationConfigCommands.ValueBool())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"authorizationRule", []interface{}{})
		for _, item := range data.AuthorizationRules {
			itemBody := ""
			if !item.RuleId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ruleId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ruleId.value", item.RuleId.ValueString())
				}
			}
			if !item.Method.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "method.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "method.value", item.Method.ValueString())
				}
			}
			if !item.Level.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "level.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "level.value", item.Level.ValueString())
				}
			}
			if !item.Group.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "group.optionType", "global")
					var values []string
					item.Group.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "group.value", values)
				}
			}
			if item.IfAuthenticated.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifAuthenticated.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "ifAuthenticated.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifAuthenticated.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ifAuthenticated.value", item.IfAuthenticated.ValueBool())
				}
			}
			body, _ = sjson.SetRaw(body, path+"authorizationRule.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemAAA) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
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
	data.ServerAuthOrder = types.ListNull(types.StringType)

	if t := res.Get(path + "serverAuthOrder.optionType"); t.Exists() {
		va := res.Get(path + "serverAuthOrder.value")
		if t.String() == "global" {
			data.ServerAuthOrder = helpers.GetStringList(va.Array())
		}
	}
	oldUsers := data.Users
	if value := res.Get(path + "user"); value.Exists() && len(value.Array()) > 0 {
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
			item.Password = types.StringNull()
			item.PasswordVariable = types.StringNull()
			if t := v.Get("password.optionType"); t.Exists() {
				va := v.Get("password.value")
				if t.String() == "variable" {
					item.PasswordVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Password = types.StringValue(va.String())
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
			if cValue := v.Get("pubkeyChain"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.PublicKeys = make([]SystemAAAUsersPublicKeys, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := SystemAAAUsersPublicKeys{}
					cItem.KeyString = types.StringNull()
					cItem.KeyStringVariable = types.StringNull()
					if t := cv.Get("keyString.optionType"); t.Exists() {
						va := cv.Get("keyString.value")
						if t.String() == "variable" {
							cItem.KeyStringVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.KeyString = types.StringValue(va.String())
						}
					}
					cItem.KeyType = types.StringNull()

					if t := cv.Get("keyType.optionType"); t.Exists() {
						va := cv.Get("keyType.value")
						if t.String() == "global" {
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
	} else {
		data.Users = nil
	}
	if !fullRead {
		resultUsers := make([]SystemAAAUsers, 0, len(data.Users))
		matchedUsers := make([]bool, len(data.Users))
		for _, oldItem := range oldUsers {
			for ni := range data.Users {
				if matchedUsers[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.NameVariable.ValueString() != "" || data.Users[ni].NameVariable.ValueString() != "") {
					if oldItem.NameVariable.ValueString() != data.Users[ni].NameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.Name.ValueString() != data.Users[ni].Name.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedUsers[ni] = true
					data.Users[ni].Password = oldItem.Password
					data.Users[ni].PasswordVariable = oldItem.PasswordVariable
					{
						resultC := make([]SystemAAAUsersPublicKeys, 0, len(data.Users[ni].PublicKeys))
						matchedC := make([]bool, len(data.Users[ni].PublicKeys))
						for _, oldCItem := range oldItem.PublicKeys {
							for nci := range data.Users[ni].PublicKeys {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC && (oldCItem.KeyStringVariable.ValueString() != "" || data.Users[ni].PublicKeys[nci].KeyStringVariable.ValueString() != "") {
									if oldCItem.KeyStringVariable.ValueString() != data.Users[ni].PublicKeys[nci].KeyStringVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.KeyString.ValueString() != data.Users[ni].PublicKeys[nci].KeyString.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.Users[ni].PublicKeys[nci])
									break
								}
							}
						}
						for nci := range data.Users[ni].PublicKeys {
							if !matchedC[nci] {
								resultC = append(resultC, data.Users[ni].PublicKeys[nci])
							}
						}
						data.Users[ni].PublicKeys = resultC
					}
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
	oldRadiusGroups := data.RadiusGroups
	if value := res.Get(path + "radius"); value.Exists() && len(value.Array()) > 0 {
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
			if cValue := v.Get("server"); cValue.Exists() && len(cValue.Array()) > 0 {
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
					cItem.Key = types.StringNull()
					cItem.KeyVariable = types.StringNull()
					if t := cv.Get("key.optionType"); t.Exists() {
						va := cv.Get("key.value")
						if t.String() == "variable" {
							cItem.KeyVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Key = types.StringValue(va.String())
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
	} else {
		data.RadiusGroups = nil
	}
	if !fullRead {
		resultRadiusGroups := make([]SystemAAARadiusGroups, 0, len(data.RadiusGroups))
		matchedRadiusGroups := make([]bool, len(data.RadiusGroups))
		for _, oldItem := range oldRadiusGroups {
			for ni := range data.RadiusGroups {
				if matchedRadiusGroups[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.GroupName.ValueString() != data.RadiusGroups[ni].GroupName.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedRadiusGroups[ni] = true
					{
						resultC := make([]SystemAAARadiusGroupsServers, 0, len(data.RadiusGroups[ni].Servers))
						matchedC := make([]bool, len(data.RadiusGroups[ni].Servers))
						for _, oldCItem := range oldItem.Servers {
							for nci := range data.RadiusGroups[ni].Servers {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if oldCItem.Address.ValueString() != data.RadiusGroups[ni].Servers[nci].Address.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									data.RadiusGroups[ni].Servers[nci].Key = oldCItem.Key
									data.RadiusGroups[ni].Servers[nci].KeyVariable = oldCItem.KeyVariable
									resultC = append(resultC, data.RadiusGroups[ni].Servers[nci])
									break
								}
							}
						}
						for nci := range data.RadiusGroups[ni].Servers {
							if !matchedC[nci] {
								resultC = append(resultC, data.RadiusGroups[ni].Servers[nci])
							}
						}
						data.RadiusGroups[ni].Servers = resultC
					}
					resultRadiusGroups = append(resultRadiusGroups, data.RadiusGroups[ni])
					break
				}
			}
		}
		for ni := range data.RadiusGroups {
			if !matchedRadiusGroups[ni] {
				resultRadiusGroups = append(resultRadiusGroups, data.RadiusGroups[ni])
			}
		}
		data.RadiusGroups = resultRadiusGroups
	}
	oldTacacsGroups := data.TacacsGroups
	if value := res.Get(path + "tacacs"); value.Exists() && len(value.Array()) > 0 {
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
			if cValue := v.Get("server"); cValue.Exists() && len(cValue.Array()) > 0 {
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
					cItem.Key = types.StringNull()
					cItem.KeyVariable = types.StringNull()
					if t := cv.Get("key.optionType"); t.Exists() {
						va := cv.Get("key.value")
						if t.String() == "variable" {
							cItem.KeyVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Key = types.StringValue(va.String())
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
	} else {
		data.TacacsGroups = nil
	}
	if !fullRead {
		resultTacacsGroups := make([]SystemAAATacacsGroups, 0, len(data.TacacsGroups))
		matchedTacacsGroups := make([]bool, len(data.TacacsGroups))
		for _, oldItem := range oldTacacsGroups {
			for ni := range data.TacacsGroups {
				if matchedTacacsGroups[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.GroupName.ValueString() != data.TacacsGroups[ni].GroupName.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedTacacsGroups[ni] = true
					{
						resultC := make([]SystemAAATacacsGroupsServers, 0, len(data.TacacsGroups[ni].Servers))
						matchedC := make([]bool, len(data.TacacsGroups[ni].Servers))
						for _, oldCItem := range oldItem.Servers {
							for nci := range data.TacacsGroups[ni].Servers {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC {
									if oldCItem.Address.ValueString() != data.TacacsGroups[ni].Servers[nci].Address.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									data.TacacsGroups[ni].Servers[nci].Key = oldCItem.Key
									data.TacacsGroups[ni].Servers[nci].KeyVariable = oldCItem.KeyVariable
									resultC = append(resultC, data.TacacsGroups[ni].Servers[nci])
									break
								}
							}
						}
						for nci := range data.TacacsGroups[ni].Servers {
							if !matchedC[nci] {
								resultC = append(resultC, data.TacacsGroups[ni].Servers[nci])
							}
						}
						data.TacacsGroups[ni].Servers = resultC
					}
					resultTacacsGroups = append(resultTacacsGroups, data.TacacsGroups[ni])
					break
				}
			}
		}
		for ni := range data.TacacsGroups {
			if !matchedTacacsGroups[ni] {
				resultTacacsGroups = append(resultTacacsGroups, data.TacacsGroups[ni])
			}
		}
		data.TacacsGroups = resultTacacsGroups
	}
	oldAccountingRules := data.AccountingRules
	if value := res.Get(path + "accountingRule"); value.Exists() && len(value.Array()) > 0 {
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
			item.Group = types.ListNull(types.StringType)

			if t := v.Get("group.optionType"); t.Exists() {
				va := v.Get("group.value")
				if t.String() == "global" {
					item.Group = helpers.GetStringList(va.Array())
				}
			}
			data.AccountingRules = append(data.AccountingRules, item)
			return true
		})
	} else {
		data.AccountingRules = nil
	}
	if !fullRead {
		resultAccountingRules := make([]SystemAAAAccountingRules, 0, len(data.AccountingRules))
		matchedAccountingRules := make([]bool, len(data.AccountingRules))
		for _, oldItem := range oldAccountingRules {
			for ni := range data.AccountingRules {
				if matchedAccountingRules[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.RuleId.ValueString() != data.AccountingRules[ni].RuleId.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedAccountingRules[ni] = true
					resultAccountingRules = append(resultAccountingRules, data.AccountingRules[ni])
					break
				}
			}
		}
		for ni := range data.AccountingRules {
			if !matchedAccountingRules[ni] {
				resultAccountingRules = append(resultAccountingRules, data.AccountingRules[ni])
			}
		}
		data.AccountingRules = resultAccountingRules
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
	oldAuthorizationRules := data.AuthorizationRules
	if value := res.Get(path + "authorizationRule"); value.Exists() && len(value.Array()) > 0 {
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
			item.Group = types.ListNull(types.StringType)

			if t := v.Get("group.optionType"); t.Exists() {
				va := v.Get("group.value")
				if t.String() == "global" {
					item.Group = helpers.GetStringList(va.Array())
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
	} else {
		data.AuthorizationRules = nil
	}
	if !fullRead {
		resultAuthorizationRules := make([]SystemAAAAuthorizationRules, 0, len(data.AuthorizationRules))
		matchedAuthorizationRules := make([]bool, len(data.AuthorizationRules))
		for _, oldItem := range oldAuthorizationRules {
			for ni := range data.AuthorizationRules {
				if matchedAuthorizationRules[ni] {
					continue
				}
				keyMatch := true
				if keyMatch {
					if oldItem.RuleId.ValueString() != data.AuthorizationRules[ni].RuleId.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedAuthorizationRules[ni] = true
					resultAuthorizationRules = append(resultAuthorizationRules, data.AuthorizationRules[ni])
					break
				}
			}
		}
		for ni := range data.AuthorizationRules {
			if !matchedAuthorizationRules[ni] {
				resultAuthorizationRules = append(resultAuthorizationRules, data.AuthorizationRules[ni])
			}
		}
		data.AuthorizationRules = resultAuthorizationRules
	}
}

// End of section. //template:end fromBody
