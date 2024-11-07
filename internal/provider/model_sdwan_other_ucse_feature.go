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
type OtherUCSE struct {
	Id                           types.String          `tfsdk:"id"`
	Version                      types.Int64           `tfsdk:"version"`
	Name                         types.String          `tfsdk:"name"`
	Description                  types.String          `tfsdk:"description"`
	FeatureProfileId             types.String          `tfsdk:"feature_profile_id"`
	Bay                          types.Int64           `tfsdk:"bay"`
	Slot                         types.Int64           `tfsdk:"slot"`
	AccessPortDedicated          types.Bool            `tfsdk:"access_port_dedicated"`
	AccessPortSharedType         types.String          `tfsdk:"access_port_shared_type"`
	AccessPortSharedFailoverType types.String          `tfsdk:"access_port_shared_failover_type"`
	Ipv4Address                  types.String          `tfsdk:"ipv4_address"`
	Ipv4AddressVariable          types.String          `tfsdk:"ipv4_address_variable"`
	DefaultGateway               types.String          `tfsdk:"default_gateway"`
	DefaultGatewayVariable       types.String          `tfsdk:"default_gateway_variable"`
	VlanId                       types.Int64           `tfsdk:"vlan_id"`
	VlanIdVariable               types.String          `tfsdk:"vlan_id_variable"`
	AssignPriority               types.Int64           `tfsdk:"assign_priority"`
	AssignPriorityVariable       types.String          `tfsdk:"assign_priority_variable"`
	Interfaces                   []OtherUCSEInterfaces `tfsdk:"interfaces"`
}

type OtherUCSEInterfaces struct {
	InterfaceName            types.String `tfsdk:"interface_name"`
	InterfaceNameVariable    types.String `tfsdk:"interface_name_variable"`
	UcseInterfaceVpn         types.Int64  `tfsdk:"ucse_interface_vpn"`
	UcseInterfaceVpnVariable types.String `tfsdk:"ucse_interface_vpn_variable"`
	Ipv4Address              types.String `tfsdk:"ipv4_address"`
	Ipv4AddressVariable      types.String `tfsdk:"ipv4_address_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data OtherUCSE) getModel() string {
	return "other_ucse"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data OtherUCSE) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/other/%v/ucse", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data OtherUCSE) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.Bay.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"bay.optionType", "global")
			body, _ = sjson.Set(body, path+"bay.value", data.Bay.ValueInt64())
		}
	}
	if !data.Slot.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"slot.optionType", "global")
			body, _ = sjson.Set(body, path+"slot.value", data.Slot.ValueInt64())
		}
	}
	if data.AccessPortDedicated.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"imc.access-port.dedicated.optionType", "default")
			body, _ = sjson.Set(body, path+"imc.access-port.dedicated.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"imc.access-port.dedicated.optionType", "global")
			body, _ = sjson.Set(body, path+"imc.access-port.dedicated.value", data.AccessPortDedicated.ValueBool())
		}
	}
	if !data.AccessPortSharedType.IsNull() {
		if true && data.AccessPortDedicated.ValueBool() == false {
			body, _ = sjson.Set(body, path+"imc.access-port.sharedLom.lomType.optionType", "global")
			body, _ = sjson.Set(body, path+"imc.access-port.sharedLom.lomType.value", data.AccessPortSharedType.ValueString())
		}
	}
	if !data.AccessPortSharedFailoverType.IsNull() {
		if true && data.AccessPortDedicated.ValueBool() == false {
			body, _ = sjson.Set(body, path+"imc.access-port.sharedLom.failOverType.optionType", "global")
			body, _ = sjson.Set(body, path+"imc.access-port.sharedLom.failOverType.value", data.AccessPortSharedFailoverType.ValueString())
		}
	}

	if !data.Ipv4AddressVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"imc.ip.address.optionType", "variable")
			body, _ = sjson.Set(body, path+"imc.ip.address.value", data.Ipv4AddressVariable.ValueString())
		}
	} else if !data.Ipv4Address.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"imc.ip.address.optionType", "global")
			body, _ = sjson.Set(body, path+"imc.ip.address.value", data.Ipv4Address.ValueString())
		}
	}

	if !data.DefaultGatewayVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"imc.ip.defaultGateway.optionType", "variable")
			body, _ = sjson.Set(body, path+"imc.ip.defaultGateway.value", data.DefaultGatewayVariable.ValueString())
		}
	} else if data.DefaultGateway.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"imc.ip.defaultGateway.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"imc.ip.defaultGateway.optionType", "global")
			body, _ = sjson.Set(body, path+"imc.ip.defaultGateway.value", data.DefaultGateway.ValueString())
		}
	}

	if !data.VlanIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"imc.vlan.vlanId.optionType", "variable")
			body, _ = sjson.Set(body, path+"imc.vlan.vlanId.value", data.VlanIdVariable.ValueString())
		}
	} else if data.VlanId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"imc.vlan.vlanId.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"imc.vlan.vlanId.optionType", "global")
			body, _ = sjson.Set(body, path+"imc.vlan.vlanId.value", data.VlanId.ValueInt64())
		}
	}

	if !data.AssignPriorityVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"imc.vlan.priority.optionType", "variable")
			body, _ = sjson.Set(body, path+"imc.vlan.priority.value", data.AssignPriorityVariable.ValueString())
		}
	} else if data.AssignPriority.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"imc.vlan.priority.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"imc.vlan.priority.optionType", "global")
			body, _ = sjson.Set(body, path+"imc.vlan.priority.value", data.AssignPriority.ValueInt64())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"interface", []interface{}{})
		for _, item := range data.Interfaces {
			itemBody := ""

			if !item.InterfaceNameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ifName.value", item.InterfaceNameVariable.ValueString())
				}
			} else if item.InterfaceName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifName.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ifName.value", item.InterfaceName.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "l3.optionType", "default")
				itemBody, _ = sjson.Set(itemBody, "l3.value", true)
			}

			if !item.UcseInterfaceVpnVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ucseInterfaceVpn.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ucseInterfaceVpn.value", item.UcseInterfaceVpnVariable.ValueString())
				}
			} else if item.UcseInterfaceVpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ucseInterfaceVpn.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ucseInterfaceVpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ucseInterfaceVpn.value", item.UcseInterfaceVpn.ValueInt64())
				}
			}

			if !item.Ipv4AddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.Ipv4AddressVariable.ValueString())
				}
			} else if !item.Ipv4Address.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.Ipv4Address.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"interface.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *OtherUCSE) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Bay = types.Int64Null()

	if t := res.Get(path + "bay.optionType"); t.Exists() {
		va := res.Get(path + "bay.value")
		if t.String() == "global" {
			data.Bay = types.Int64Value(va.Int())
		}
	}
	data.Slot = types.Int64Null()

	if t := res.Get(path + "slot.optionType"); t.Exists() {
		va := res.Get(path + "slot.value")
		if t.String() == "global" {
			data.Slot = types.Int64Value(va.Int())
		}
	}
	data.AccessPortDedicated = types.BoolNull()

	if t := res.Get(path + "imc.access-port.dedicated.optionType"); t.Exists() {
		va := res.Get(path + "imc.access-port.dedicated.value")
		if t.String() == "global" {
			data.AccessPortDedicated = types.BoolValue(va.Bool())
		}
	}
	data.AccessPortSharedType = types.StringNull()

	if t := res.Get(path + "imc.access-port.sharedLom.lomType.optionType"); t.Exists() {
		va := res.Get(path + "imc.access-port.sharedLom.lomType.value")
		if t.String() == "global" {
			data.AccessPortSharedType = types.StringValue(va.String())
		}
	}
	data.AccessPortSharedFailoverType = types.StringNull()

	if t := res.Get(path + "imc.access-port.sharedLom.failOverType.optionType"); t.Exists() {
		va := res.Get(path + "imc.access-port.sharedLom.failOverType.value")
		if t.String() == "global" {
			data.AccessPortSharedFailoverType = types.StringValue(va.String())
		}
	}
	data.Ipv4Address = types.StringNull()
	data.Ipv4AddressVariable = types.StringNull()
	if t := res.Get(path + "imc.ip.address.optionType"); t.Exists() {
		va := res.Get(path + "imc.ip.address.value")
		if t.String() == "variable" {
			data.Ipv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Address = types.StringValue(va.String())
		}
	}
	data.DefaultGateway = types.StringNull()
	data.DefaultGatewayVariable = types.StringNull()
	if t := res.Get(path + "imc.ip.defaultGateway.optionType"); t.Exists() {
		va := res.Get(path + "imc.ip.defaultGateway.value")
		if t.String() == "variable" {
			data.DefaultGatewayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultGateway = types.StringValue(va.String())
		}
	}
	data.VlanId = types.Int64Null()
	data.VlanIdVariable = types.StringNull()
	if t := res.Get(path + "imc.vlan.vlanId.optionType"); t.Exists() {
		va := res.Get(path + "imc.vlan.vlanId.value")
		if t.String() == "variable" {
			data.VlanIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.VlanId = types.Int64Value(va.Int())
		}
	}
	data.AssignPriority = types.Int64Null()
	data.AssignPriorityVariable = types.StringNull()
	if t := res.Get(path + "imc.vlan.priority.optionType"); t.Exists() {
		va := res.Get(path + "imc.vlan.priority.value")
		if t.String() == "variable" {
			data.AssignPriorityVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AssignPriority = types.Int64Value(va.Int())
		}
	}
	if value := res.Get(path + "interface"); value.Exists() {
		data.Interfaces = make([]OtherUCSEInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := OtherUCSEInterfaces{}
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
			item.UcseInterfaceVpn = types.Int64Null()
			item.UcseInterfaceVpnVariable = types.StringNull()
			if t := v.Get("ucseInterfaceVpn.optionType"); t.Exists() {
				va := v.Get("ucseInterfaceVpn.value")
				if t.String() == "variable" {
					item.UcseInterfaceVpnVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.UcseInterfaceVpn = types.Int64Value(va.Int())
				}
			}
			item.Ipv4Address = types.StringNull()
			item.Ipv4AddressVariable = types.StringNull()
			if t := v.Get("address.optionType"); t.Exists() {
				va := v.Get("address.value")
				if t.String() == "variable" {
					item.Ipv4AddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Ipv4Address = types.StringValue(va.String())
				}
			}
			data.Interfaces = append(data.Interfaces, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *OtherUCSE) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Bay = types.Int64Null()

	if t := res.Get(path + "bay.optionType"); t.Exists() {
		va := res.Get(path + "bay.value")
		if t.String() == "global" {
			data.Bay = types.Int64Value(va.Int())
		}
	}
	data.Slot = types.Int64Null()

	if t := res.Get(path + "slot.optionType"); t.Exists() {
		va := res.Get(path + "slot.value")
		if t.String() == "global" {
			data.Slot = types.Int64Value(va.Int())
		}
	}
	data.AccessPortDedicated = types.BoolNull()

	if t := res.Get(path + "imc.access-port.dedicated.optionType"); t.Exists() {
		va := res.Get(path + "imc.access-port.dedicated.value")
		if t.String() == "global" {
			data.AccessPortDedicated = types.BoolValue(va.Bool())
		}
	}
	data.AccessPortSharedType = types.StringNull()

	if t := res.Get(path + "imc.access-port.sharedLom.lomType.optionType"); t.Exists() {
		va := res.Get(path + "imc.access-port.sharedLom.lomType.value")
		if t.String() == "global" {
			data.AccessPortSharedType = types.StringValue(va.String())
		}
	}
	data.AccessPortSharedFailoverType = types.StringNull()

	if t := res.Get(path + "imc.access-port.sharedLom.failOverType.optionType"); t.Exists() {
		va := res.Get(path + "imc.access-port.sharedLom.failOverType.value")
		if t.String() == "global" {
			data.AccessPortSharedFailoverType = types.StringValue(va.String())
		}
	}
	data.Ipv4Address = types.StringNull()
	data.Ipv4AddressVariable = types.StringNull()
	if t := res.Get(path + "imc.ip.address.optionType"); t.Exists() {
		va := res.Get(path + "imc.ip.address.value")
		if t.String() == "variable" {
			data.Ipv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Address = types.StringValue(va.String())
		}
	}
	data.DefaultGateway = types.StringNull()
	data.DefaultGatewayVariable = types.StringNull()
	if t := res.Get(path + "imc.ip.defaultGateway.optionType"); t.Exists() {
		va := res.Get(path + "imc.ip.defaultGateway.value")
		if t.String() == "variable" {
			data.DefaultGatewayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultGateway = types.StringValue(va.String())
		}
	}
	data.VlanId = types.Int64Null()
	data.VlanIdVariable = types.StringNull()
	if t := res.Get(path + "imc.vlan.vlanId.optionType"); t.Exists() {
		va := res.Get(path + "imc.vlan.vlanId.value")
		if t.String() == "variable" {
			data.VlanIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.VlanId = types.Int64Value(va.Int())
		}
	}
	data.AssignPriority = types.Int64Null()
	data.AssignPriorityVariable = types.StringNull()
	if t := res.Get(path + "imc.vlan.priority.optionType"); t.Exists() {
		va := res.Get(path + "imc.vlan.priority.value")
		if t.String() == "variable" {
			data.AssignPriorityVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AssignPriority = types.Int64Value(va.Int())
		}
	}
	for i := range data.Interfaces {
		keys := [...]string{"ifName"}
		keyValues := [...]string{data.Interfaces[i].InterfaceName.ValueString()}
		keyValuesVariables := [...]string{data.Interfaces[i].InterfaceNameVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "interface").ForEach(
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
		data.Interfaces[i].InterfaceName = types.StringNull()
		data.Interfaces[i].InterfaceNameVariable = types.StringNull()
		if t := r.Get("ifName.optionType"); t.Exists() {
			va := r.Get("ifName.value")
			if t.String() == "variable" {
				data.Interfaces[i].InterfaceNameVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].InterfaceName = types.StringValue(va.String())
			}
		}
		data.Interfaces[i].UcseInterfaceVpn = types.Int64Null()
		data.Interfaces[i].UcseInterfaceVpnVariable = types.StringNull()
		if t := r.Get("ucseInterfaceVpn.optionType"); t.Exists() {
			va := r.Get("ucseInterfaceVpn.value")
			if t.String() == "variable" {
				data.Interfaces[i].UcseInterfaceVpnVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].UcseInterfaceVpn = types.Int64Value(va.Int())
			}
		}
		data.Interfaces[i].Ipv4Address = types.StringNull()
		data.Interfaces[i].Ipv4AddressVariable = types.StringNull()
		if t := r.Get("address.optionType"); t.Exists() {
			va := r.Get("address.value")
			if t.String() == "variable" {
				data.Interfaces[i].Ipv4AddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Interfaces[i].Ipv4Address = types.StringValue(va.String())
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *OtherUCSE) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.Bay.IsNull() {
		return false
	}
	if !data.Slot.IsNull() {
		return false
	}
	if !data.AccessPortDedicated.IsNull() {
		return false
	}
	if !data.AccessPortSharedType.IsNull() {
		return false
	}
	if !data.AccessPortSharedFailoverType.IsNull() {
		return false
	}
	if !data.Ipv4Address.IsNull() {
		return false
	}
	if !data.Ipv4AddressVariable.IsNull() {
		return false
	}
	if !data.DefaultGateway.IsNull() {
		return false
	}
	if !data.DefaultGatewayVariable.IsNull() {
		return false
	}
	if !data.VlanId.IsNull() {
		return false
	}
	if !data.VlanIdVariable.IsNull() {
		return false
	}
	if !data.AssignPriority.IsNull() {
		return false
	}
	if !data.AssignPriorityVariable.IsNull() {
		return false
	}
	if len(data.Interfaces) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
