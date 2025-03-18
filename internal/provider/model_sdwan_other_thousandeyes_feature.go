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
type OtherThousandEyes struct {
	Id                 types.String                          `tfsdk:"id"`
	Version            types.Int64                           `tfsdk:"version"`
	Name               types.String                          `tfsdk:"name"`
	Description        types.String                          `tfsdk:"description"`
	FeatureProfileId   types.String                          `tfsdk:"feature_profile_id"`
	VirtualApplication []OtherThousandEyesVirtualApplication `tfsdk:"virtual_application"`
}

type OtherThousandEyesVirtualApplication struct {
	AccountGroupToken            types.String `tfsdk:"account_group_token"`
	AccountGroupTokenVariable    types.String `tfsdk:"account_group_token_variable"`
	Vpn                          types.Int64  `tfsdk:"vpn"`
	VpnVariable                  types.String `tfsdk:"vpn_variable"`
	ManagementIp                 types.String `tfsdk:"management_ip"`
	ManagementIpVariable         types.String `tfsdk:"management_ip_variable"`
	ManagementSubnetMask         types.String `tfsdk:"management_subnet_mask"`
	ManagementSubnetMaskVariable types.String `tfsdk:"management_subnet_mask_variable"`
	AgentDefaultGateway          types.String `tfsdk:"agent_default_gateway"`
	AgentDefaultGatewayVariable  types.String `tfsdk:"agent_default_gateway_variable"`
	NameServerIp                 types.String `tfsdk:"name_server_ip"`
	NameServerIpVariable         types.String `tfsdk:"name_server_ip_variable"`
	Hostname                     types.String `tfsdk:"hostname"`
	HostnameVariable             types.String `tfsdk:"hostname_variable"`
	ProxyType                    types.String `tfsdk:"proxy_type"`
	ProxyHost                    types.String `tfsdk:"proxy_host"`
	ProxyHostVariable            types.String `tfsdk:"proxy_host_variable"`
	ProxyPort                    types.Int64  `tfsdk:"proxy_port"`
	ProxyPortVariable            types.String `tfsdk:"proxy_port_variable"`
	PacUrl                       types.String `tfsdk:"pac_url"`
	PacUrlVariable               types.String `tfsdk:"pac_url_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data OtherThousandEyes) getModel() string {
	return "other_thousandeyes"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data OtherThousandEyes) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/other/%v/thousandeyes", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data OtherThousandEyes) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if true {

		for _, item := range data.VirtualApplication {
			itemBody := ""

			if !item.AccountGroupTokenVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "token.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "token.value", item.AccountGroupTokenVariable.ValueString())
				}
			} else if !item.AccountGroupToken.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "token.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "token.value", item.AccountGroupToken.ValueString())
				}
			}

			if !item.VpnVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.VpnVariable.ValueString())
				}
			} else if item.Vpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
				}
			}

			if !item.ManagementIpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "teMgmtIp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "teMgmtIp.value", item.ManagementIpVariable.ValueString())
				}
			} else if !item.ManagementIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "teMgmtIp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "teMgmtIp.value", item.ManagementIp.ValueString())
				}
			}

			if !item.ManagementSubnetMaskVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "teMgmtSubnetMask.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "teMgmtSubnetMask.value", item.ManagementSubnetMaskVariable.ValueString())
				}
			} else if !item.ManagementSubnetMask.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "teMgmtSubnetMask.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "teMgmtSubnetMask.value", item.ManagementSubnetMask.ValueString())
				}
			}

			if !item.AgentDefaultGatewayVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "teVpgIp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "teVpgIp.value", item.AgentDefaultGatewayVariable.ValueString())
				}
			} else if !item.AgentDefaultGateway.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "teVpgIp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "teVpgIp.value", item.AgentDefaultGateway.ValueString())
				}
			}

			if !item.NameServerIpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nameServer.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "nameServer.value", item.NameServerIpVariable.ValueString())
				}
			} else if item.NameServerIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nameServer.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nameServer.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "nameServer.value", item.NameServerIp.ValueString())
				}
			}

			if !item.HostnameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "hostname.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "hostname.value", item.HostnameVariable.ValueString())
				}
			} else if item.Hostname.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "hostname.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "hostname.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "hostname.value", item.Hostname.ValueString())
				}
			}
			if !item.ProxyType.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "proxyConfig.proxyType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "proxyConfig.proxyType.value", item.ProxyType.ValueString())
				}
			}

			if !item.ProxyHostVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "proxyConfig.proxyHost.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "proxyConfig.proxyHost.value", item.ProxyHostVariable.ValueString())
				}
			} else if !item.ProxyHost.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "proxyConfig.proxyHost.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "proxyConfig.proxyHost.value", item.ProxyHost.ValueString())
				}
			}

			if !item.ProxyPortVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "proxyConfig.proxyPort.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "proxyConfig.proxyPort.value", item.ProxyPortVariable.ValueString())
				}
			} else if !item.ProxyPort.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "proxyConfig.proxyPort.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "proxyConfig.proxyPort.value", item.ProxyPort.ValueInt64())
				}
			}

			if !item.PacUrlVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "proxyConfig.pacUrl.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "proxyConfig.pacUrl.value", item.PacUrlVariable.ValueString())
				}
			} else if !item.PacUrl.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "proxyConfig.pacUrl.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "proxyConfig.pacUrl.value", item.PacUrl.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"virtualApplication.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *OtherThousandEyes) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	if value := res.Get(path + "virtualApplication"); value.Exists() {
		data.VirtualApplication = make([]OtherThousandEyesVirtualApplication, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := OtherThousandEyesVirtualApplication{}
			item.AccountGroupToken = types.StringNull()
			item.AccountGroupTokenVariable = types.StringNull()
			if t := v.Get("token.optionType"); t.Exists() {
				va := v.Get("token.value")
				if t.String() == "variable" {
					item.AccountGroupTokenVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AccountGroupToken = types.StringValue(va.String())
				}
			}
			item.Vpn = types.Int64Null()
			item.VpnVariable = types.StringNull()
			if t := v.Get("vpn.optionType"); t.Exists() {
				va := v.Get("vpn.value")
				if t.String() == "variable" {
					item.VpnVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Vpn = types.Int64Value(va.Int())
				}
			}
			item.ManagementIp = types.StringNull()
			item.ManagementIpVariable = types.StringNull()
			if t := v.Get("teMgmtIp.optionType"); t.Exists() {
				va := v.Get("teMgmtIp.value")
				if t.String() == "variable" {
					item.ManagementIpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.ManagementIp = types.StringValue(va.String())
				}
			}
			item.ManagementSubnetMask = types.StringNull()
			item.ManagementSubnetMaskVariable = types.StringNull()
			if t := v.Get("teMgmtSubnetMask.optionType"); t.Exists() {
				va := v.Get("teMgmtSubnetMask.value")
				if t.String() == "variable" {
					item.ManagementSubnetMaskVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.ManagementSubnetMask = types.StringValue(va.String())
				}
			}
			item.AgentDefaultGateway = types.StringNull()
			item.AgentDefaultGatewayVariable = types.StringNull()
			if t := v.Get("teVpgIp.optionType"); t.Exists() {
				va := v.Get("teVpgIp.value")
				if t.String() == "variable" {
					item.AgentDefaultGatewayVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AgentDefaultGateway = types.StringValue(va.String())
				}
			}
			item.NameServerIp = types.StringNull()
			item.NameServerIpVariable = types.StringNull()
			if t := v.Get("nameServer.optionType"); t.Exists() {
				va := v.Get("nameServer.value")
				if t.String() == "variable" {
					item.NameServerIpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NameServerIp = types.StringValue(va.String())
				}
			}
			item.Hostname = types.StringNull()
			item.HostnameVariable = types.StringNull()
			if t := v.Get("hostname.optionType"); t.Exists() {
				va := v.Get("hostname.value")
				if t.String() == "variable" {
					item.HostnameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Hostname = types.StringValue(va.String())
				}
			}
			item.ProxyType = types.StringNull()

			if t := v.Get("proxyConfig.proxyType.optionType"); t.Exists() {
				va := v.Get("proxyConfig.proxyType.value")
				if t.String() == "global" {
					item.ProxyType = types.StringValue(va.String())
				}
			}
			item.ProxyHost = types.StringNull()
			item.ProxyHostVariable = types.StringNull()
			if t := v.Get("proxyConfig.proxyHost.optionType"); t.Exists() {
				va := v.Get("proxyConfig.proxyHost.value")
				if t.String() == "variable" {
					item.ProxyHostVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.ProxyHost = types.StringValue(va.String())
				}
			}
			item.ProxyPort = types.Int64Null()
			item.ProxyPortVariable = types.StringNull()
			if t := v.Get("proxyConfig.proxyPort.optionType"); t.Exists() {
				va := v.Get("proxyConfig.proxyPort.value")
				if t.String() == "variable" {
					item.ProxyPortVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.ProxyPort = types.Int64Value(va.Int())
				}
			}
			item.PacUrl = types.StringNull()
			item.PacUrlVariable = types.StringNull()
			if t := v.Get("proxyConfig.pacUrl.optionType"); t.Exists() {
				va := v.Get("proxyConfig.pacUrl.value")
				if t.String() == "variable" {
					item.PacUrlVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.PacUrl = types.StringValue(va.String())
				}
			}
			data.VirtualApplication = append(data.VirtualApplication, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *OtherThousandEyes) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	for i := range data.VirtualApplication {
		keys := [...]string{"token"}
		keyValues := [...]string{data.VirtualApplication[i].AccountGroupToken.ValueString()}
		keyValuesVariables := [...]string{data.VirtualApplication[i].AccountGroupTokenVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "virtualApplication").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
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
		data.VirtualApplication[i].AccountGroupToken = types.StringNull()
		data.VirtualApplication[i].AccountGroupTokenVariable = types.StringNull()
		if t := r.Get("token.optionType"); t.Exists() {
			va := r.Get("token.value")
			if t.String() == "variable" {
				data.VirtualApplication[i].AccountGroupTokenVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.VirtualApplication[i].AccountGroupToken = types.StringValue(va.String())
			}
		}
		data.VirtualApplication[i].Vpn = types.Int64Null()
		data.VirtualApplication[i].VpnVariable = types.StringNull()
		if t := r.Get("vpn.optionType"); t.Exists() {
			va := r.Get("vpn.value")
			if t.String() == "variable" {
				data.VirtualApplication[i].VpnVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.VirtualApplication[i].Vpn = types.Int64Value(va.Int())
			}
		}
		data.VirtualApplication[i].ManagementIp = types.StringNull()
		data.VirtualApplication[i].ManagementIpVariable = types.StringNull()
		if t := r.Get("teMgmtIp.optionType"); t.Exists() {
			va := r.Get("teMgmtIp.value")
			if t.String() == "variable" {
				data.VirtualApplication[i].ManagementIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.VirtualApplication[i].ManagementIp = types.StringValue(va.String())
			}
		}
		data.VirtualApplication[i].ManagementSubnetMask = types.StringNull()
		data.VirtualApplication[i].ManagementSubnetMaskVariable = types.StringNull()
		if t := r.Get("teMgmtSubnetMask.optionType"); t.Exists() {
			va := r.Get("teMgmtSubnetMask.value")
			if t.String() == "variable" {
				data.VirtualApplication[i].ManagementSubnetMaskVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.VirtualApplication[i].ManagementSubnetMask = types.StringValue(va.String())
			}
		}
		data.VirtualApplication[i].AgentDefaultGateway = types.StringNull()
		data.VirtualApplication[i].AgentDefaultGatewayVariable = types.StringNull()
		if t := r.Get("teVpgIp.optionType"); t.Exists() {
			va := r.Get("teVpgIp.value")
			if t.String() == "variable" {
				data.VirtualApplication[i].AgentDefaultGatewayVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.VirtualApplication[i].AgentDefaultGateway = types.StringValue(va.String())
			}
		}
		data.VirtualApplication[i].NameServerIp = types.StringNull()
		data.VirtualApplication[i].NameServerIpVariable = types.StringNull()
		if t := r.Get("nameServer.optionType"); t.Exists() {
			va := r.Get("nameServer.value")
			if t.String() == "variable" {
				data.VirtualApplication[i].NameServerIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.VirtualApplication[i].NameServerIp = types.StringValue(va.String())
			}
		}
		data.VirtualApplication[i].Hostname = types.StringNull()
		data.VirtualApplication[i].HostnameVariable = types.StringNull()
		if t := r.Get("hostname.optionType"); t.Exists() {
			va := r.Get("hostname.value")
			if t.String() == "variable" {
				data.VirtualApplication[i].HostnameVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.VirtualApplication[i].Hostname = types.StringValue(va.String())
			}
		}
		data.VirtualApplication[i].ProxyType = types.StringNull()

		if t := r.Get("proxyConfig.proxyType.optionType"); t.Exists() {
			va := r.Get("proxyConfig.proxyType.value")
			if t.String() == "global" {
				data.VirtualApplication[i].ProxyType = types.StringValue(va.String())
			}
		}
		data.VirtualApplication[i].ProxyHost = types.StringNull()
		data.VirtualApplication[i].ProxyHostVariable = types.StringNull()
		if t := r.Get("proxyConfig.proxyHost.optionType"); t.Exists() {
			va := r.Get("proxyConfig.proxyHost.value")
			if t.String() == "variable" {
				data.VirtualApplication[i].ProxyHostVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.VirtualApplication[i].ProxyHost = types.StringValue(va.String())
			}
		}
		data.VirtualApplication[i].ProxyPort = types.Int64Null()
		data.VirtualApplication[i].ProxyPortVariable = types.StringNull()
		if t := r.Get("proxyConfig.proxyPort.optionType"); t.Exists() {
			va := r.Get("proxyConfig.proxyPort.value")
			if t.String() == "variable" {
				data.VirtualApplication[i].ProxyPortVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.VirtualApplication[i].ProxyPort = types.Int64Value(va.Int())
			}
		}
		data.VirtualApplication[i].PacUrl = types.StringNull()
		data.VirtualApplication[i].PacUrlVariable = types.StringNull()
		if t := r.Get("proxyConfig.pacUrl.optionType"); t.Exists() {
			va := r.Get("proxyConfig.pacUrl.value")
			if t.String() == "variable" {
				data.VirtualApplication[i].PacUrlVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.VirtualApplication[i].PacUrl = types.StringValue(va.String())
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *OtherThousandEyes) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if len(data.VirtualApplication) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
