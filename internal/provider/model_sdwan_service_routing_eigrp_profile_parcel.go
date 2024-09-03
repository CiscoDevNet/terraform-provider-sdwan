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
type ServiceRoutingEIGRP struct {
	Id                            types.String                       `tfsdk:"id"`
	Version                       types.Int64                        `tfsdk:"version"`
	Name                          types.String                       `tfsdk:"name"`
	Description                   types.String                       `tfsdk:"description"`
	FeatureProfileId              types.String                       `tfsdk:"feature_profile_id"`
	AutonomousSystemId            types.Int64                        `tfsdk:"autonomous_system_id"`
	AutonomousSystemIdVariable    types.String                       `tfsdk:"autonomous_system_id_variable"`
	Redistributes                 []ServiceRoutingEIGRPRedistributes `tfsdk:"redistributes"`
	Networks                      []ServiceRoutingEIGRPNetworks      `tfsdk:"networks"`
	HelloInterval                 types.Int64                        `tfsdk:"hello_interval"`
	HelloIntervalVariable         types.String                       `tfsdk:"hello_interval_variable"`
	HoldTime                      types.Int64                        `tfsdk:"hold_time"`
	HoldTimeVariable              types.String                       `tfsdk:"hold_time_variable"`
	Type                          types.String                       `tfsdk:"type"`
	TypeVariable                  types.String                       `tfsdk:"type_variable"`
	HmacAuthenticationKey         types.String                       `tfsdk:"hmac_authentication_key"`
	HmacAuthenticationKeyVariable types.String                       `tfsdk:"hmac_authentication_key_variable"`
	Md5Keys                       []ServiceRoutingEIGRPMd5Keys       `tfsdk:"md5_keys"`
	AfInterfaces                  []ServiceRoutingEIGRPAfInterfaces  `tfsdk:"af_interfaces"`
	RoutePolicyId                 types.String                       `tfsdk:"route_policy_id"`
	Filter                        types.Bool                         `tfsdk:"filter"`
	FilterVariable                types.String                       `tfsdk:"filter_variable"`
}

type ServiceRoutingEIGRPRedistributes struct {
	Protocol         types.String `tfsdk:"protocol"`
	ProtocolVariable types.String `tfsdk:"protocol_variable"`
	RoutePolicyId    types.String `tfsdk:"route_policy_id"`
}

type ServiceRoutingEIGRPNetworks struct {
	IpAddress         types.String `tfsdk:"ip_address"`
	IpAddressVariable types.String `tfsdk:"ip_address_variable"`
	Mask              types.String `tfsdk:"mask"`
	MaskVariable      types.String `tfsdk:"mask_variable"`
}

type ServiceRoutingEIGRPMd5Keys struct {
	KeyId                     types.Int64  `tfsdk:"key_id"`
	KeyIdVariable             types.String `tfsdk:"key_id_variable"`
	AuthenticationKey         types.String `tfsdk:"authentication_key"`
	AuthenticationKeyVariable types.String `tfsdk:"authentication_key_variable"`
}

type ServiceRoutingEIGRPAfInterfaces struct {
	Name             types.String                                      `tfsdk:"name"`
	NameVariable     types.String                                      `tfsdk:"name_variable"`
	Shutdown         types.Bool                                        `tfsdk:"shutdown"`
	ShutdownVariable types.String                                      `tfsdk:"shutdown_variable"`
	SummaryAddresses []ServiceRoutingEIGRPAfInterfacesSummaryAddresses `tfsdk:"summary_addresses"`
}

type ServiceRoutingEIGRPAfInterfacesSummaryAddresses struct {
	Address         types.String `tfsdk:"address"`
	AddressVariable types.String `tfsdk:"address_variable"`
	Mask            types.String `tfsdk:"mask"`
	MaskVariable    types.String `tfsdk:"mask_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceRoutingEIGRP) getModel() string {
	return "service_routing_eigrp"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceRoutingEIGRP) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/routing/eigrp", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceRoutingEIGRP) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.AutonomousSystemIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"asNum.optionType", "variable")
			body, _ = sjson.Set(body, path+"asNum.value", data.AutonomousSystemIdVariable.ValueString())
		}
	} else if !data.AutonomousSystemId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"asNum.optionType", "global")
			body, _ = sjson.Set(body, path+"asNum.value", data.AutonomousSystemId.ValueInt64())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"addressFamily.redistribute", []interface{}{})
		for _, item := range data.Redistributes {
			itemBody := ""

			if !item.ProtocolVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "protocol.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "protocol.value", item.ProtocolVariable.ValueString())
				}
			} else if !item.Protocol.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "protocol.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "protocol.value", item.Protocol.ValueString())
				}
			}
			if data.RoutePolicyId.IsNull() {
				body, _ = sjson.Set(body, path+"routePolicy.optionType", "default")
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.value", item.RoutePolicyId.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"addressFamily.redistribute.-1", itemBody)
		}
	}
	if true {

		for _, item := range data.Networks {
			itemBody := ""

			if !item.IpAddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.address.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.address.value", item.IpAddressVariable.ValueString())
				}
			} else if !item.IpAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.address.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.address.value", item.IpAddress.ValueString())
				}
			}

			if !item.MaskVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.mask.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.mask.value", item.MaskVariable.ValueString())
				}
			} else if !item.Mask.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.mask.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.mask.value", item.Mask.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"addressFamily.network.-1", itemBody)
		}
	}

	if !data.HelloIntervalVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"helloInterval.optionType", "variable")
			body, _ = sjson.Set(body, path+"helloInterval.value", data.HelloIntervalVariable.ValueString())
		}
	} else if data.HelloInterval.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"helloInterval.optionType", "default")
			body, _ = sjson.Set(body, path+"helloInterval.value", 5)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"helloInterval.optionType", "global")
			body, _ = sjson.Set(body, path+"helloInterval.value", data.HelloInterval.ValueInt64())
		}
	}

	if !data.HoldTimeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"holdTime.optionType", "variable")
			body, _ = sjson.Set(body, path+"holdTime.value", data.HoldTimeVariable.ValueString())
		}
	} else if data.HoldTime.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"holdTime.optionType", "default")
			body, _ = sjson.Set(body, path+"holdTime.value", 15)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"holdTime.optionType", "global")
			body, _ = sjson.Set(body, path+"holdTime.value", data.HoldTime.ValueInt64())
		}
	}

	if !data.TypeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"authentication.type.optionType", "variable")
			body, _ = sjson.Set(body, path+"authentication.type.value", data.TypeVariable.ValueString())
		}
	} else if data.Type.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"authentication.type.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"authentication.type.optionType", "global")
			body, _ = sjson.Set(body, path+"authentication.type.value", data.Type.ValueString())
		}
	}

	if !data.HmacAuthenticationKeyVariable.IsNull() {
		if true && data.Type.ValueString() == "hmac-sha-256" {
			body, _ = sjson.Set(body, path+"authentication.authKey.optionType", "variable")
			body, _ = sjson.Set(body, path+"authentication.authKey.value", data.HmacAuthenticationKeyVariable.ValueString())
		}
	} else if data.HmacAuthenticationKey.IsNull() {
		if true && data.Type.ValueString() == "hmac-sha-256" {
			body, _ = sjson.Set(body, path+"authentication.authKey.optionType", "default")

		}
	} else {
		if true && data.Type.ValueString() == "hmac-sha-256" {
			body, _ = sjson.Set(body, path+"authentication.authKey.optionType", "global")
			body, _ = sjson.Set(body, path+"authentication.authKey.value", data.HmacAuthenticationKey.ValueString())
		}
	}
	if true && data.Type.ValueString() == "md5" {
		body, _ = sjson.Set(body, path+"authentication.key", []interface{}{})
		for _, item := range data.Md5Keys {
			itemBody := ""

			if !item.KeyIdVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "keyId.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "keyId.value", item.KeyIdVariable.ValueString())
				}
			} else if item.KeyId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "keyId.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "keyId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "keyId.value", item.KeyId.ValueInt64())
				}
			}

			if !item.AuthenticationKeyVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "keystring.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "keystring.value", item.AuthenticationKeyVariable.ValueString())
				}
			} else if item.AuthenticationKey.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "keystring.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "keystring.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "keystring.value", item.AuthenticationKey.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"authentication.key.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"afInterface", []interface{}{})
		for _, item := range data.AfInterfaces {
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

			if !item.ShutdownVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "shutdown.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "shutdown.value", item.ShutdownVariable.ValueString())
				}
			} else if item.Shutdown.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "shutdown.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "shutdown.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "shutdown.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "shutdown.value", item.Shutdown.ValueBool())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "summaryAddress", []interface{}{})
				for _, childItem := range item.SummaryAddresses {
					itemChildBody := ""

					if !childItem.AddressVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.address.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.address.value", childItem.AddressVariable.ValueString())
						}
					} else if !childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.address.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.address.value", childItem.Address.ValueString())
						}
					}

					if !childItem.MaskVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.mask.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.mask.value", childItem.MaskVariable.ValueString())
						}
					} else if !childItem.Mask.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.mask.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.mask.value", childItem.Mask.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "summaryAddress.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"afInterface.-1", itemBody)
		}
	}
	if data.RoutePolicyId.IsNull() {
		body, _ = sjson.Set(body, path+"tableMap.name.optionType", "default")
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tableMap.name.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"tableMap.name.refId.value", data.RoutePolicyId.ValueString())
		}
	}

	if !data.FilterVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tableMap.filter.optionType", "variable")
			body, _ = sjson.Set(body, path+"tableMap.filter.value", data.FilterVariable.ValueString())
		}
	} else if data.Filter.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tableMap.filter.optionType", "default")
			body, _ = sjson.Set(body, path+"tableMap.filter.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tableMap.filter.optionType", "global")
			body, _ = sjson.Set(body, path+"tableMap.filter.value", data.Filter.ValueBool())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceRoutingEIGRP) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.AutonomousSystemId = types.Int64Null()
	data.AutonomousSystemIdVariable = types.StringNull()
	if t := res.Get(path + "asNum.optionType"); t.Exists() {
		va := res.Get(path + "asNum.value")
		if t.String() == "variable" {
			data.AutonomousSystemIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AutonomousSystemId = types.Int64Value(va.Int())
		}
	}
	if value := res.Get(path + "addressFamily.redistribute"); value.Exists() {
		data.Redistributes = make([]ServiceRoutingEIGRPRedistributes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingEIGRPRedistributes{}
			item.Protocol = types.StringNull()
			item.ProtocolVariable = types.StringNull()
			if t := v.Get("protocol.optionType"); t.Exists() {
				va := v.Get("protocol.value")
				if t.String() == "variable" {
					item.ProtocolVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Protocol = types.StringValue(va.String())
				}
			}
			item.RoutePolicyId = types.StringNull()

			if t := v.Get("routePolicy.refId.optionType"); t.Exists() {
				va := v.Get("routePolicy.refId.value")
				if t.String() == "global" {
					item.RoutePolicyId = types.StringValue(va.String())
				}
			}
			data.Redistributes = append(data.Redistributes, item)
			return true
		})
	}
	if value := res.Get(path + "addressFamily.network"); value.Exists() {
		data.Networks = make([]ServiceRoutingEIGRPNetworks, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingEIGRPNetworks{}
			item.IpAddress = types.StringNull()
			item.IpAddressVariable = types.StringNull()
			if t := v.Get("prefix.address.optionType"); t.Exists() {
				va := v.Get("prefix.address.value")
				if t.String() == "variable" {
					item.IpAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IpAddress = types.StringValue(va.String())
				}
			}
			item.Mask = types.StringNull()
			item.MaskVariable = types.StringNull()
			if t := v.Get("prefix.mask.optionType"); t.Exists() {
				va := v.Get("prefix.mask.value")
				if t.String() == "variable" {
					item.MaskVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Mask = types.StringValue(va.String())
				}
			}
			data.Networks = append(data.Networks, item)
			return true
		})
	}
	data.HelloInterval = types.Int64Null()
	data.HelloIntervalVariable = types.StringNull()
	if t := res.Get(path + "helloInterval.optionType"); t.Exists() {
		va := res.Get(path + "helloInterval.value")
		if t.String() == "variable" {
			data.HelloIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.HelloInterval = types.Int64Value(va.Int())
		}
	}
	data.HoldTime = types.Int64Null()
	data.HoldTimeVariable = types.StringNull()
	if t := res.Get(path + "holdTime.optionType"); t.Exists() {
		va := res.Get(path + "holdTime.value")
		if t.String() == "variable" {
			data.HoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.HoldTime = types.Int64Value(va.Int())
		}
	}
	data.Type = types.StringNull()
	data.TypeVariable = types.StringNull()
	if t := res.Get(path + "authentication.type.optionType"); t.Exists() {
		va := res.Get(path + "authentication.type.value")
		if t.String() == "variable" {
			data.TypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Type = types.StringValue(va.String())
		}
	}
	if value := res.Get(path + "authentication.key"); value.Exists() {
		data.Md5Keys = make([]ServiceRoutingEIGRPMd5Keys, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingEIGRPMd5Keys{}
			item.KeyId = types.Int64Null()
			item.KeyIdVariable = types.StringNull()
			if t := v.Get("keyId.optionType"); t.Exists() {
				va := v.Get("keyId.value")
				if t.String() == "variable" {
					item.KeyIdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.KeyId = types.Int64Value(va.Int())
				}
			}
			data.Md5Keys = append(data.Md5Keys, item)
			return true
		})
	}
	if value := res.Get(path + "afInterface"); value.Exists() {
		data.AfInterfaces = make([]ServiceRoutingEIGRPAfInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingEIGRPAfInterfaces{}
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
			if cValue := v.Get("summaryAddress"); cValue.Exists() {
				item.SummaryAddresses = make([]ServiceRoutingEIGRPAfInterfacesSummaryAddresses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceRoutingEIGRPAfInterfacesSummaryAddresses{}
					cItem.Address = types.StringNull()
					cItem.AddressVariable = types.StringNull()
					if t := cv.Get("prefix.address.optionType"); t.Exists() {
						va := cv.Get("prefix.address.value")
						if t.String() == "variable" {
							cItem.AddressVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					cItem.Mask = types.StringNull()
					cItem.MaskVariable = types.StringNull()
					if t := cv.Get("prefix.mask.optionType"); t.Exists() {
						va := cv.Get("prefix.mask.value")
						if t.String() == "variable" {
							cItem.MaskVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Mask = types.StringValue(va.String())
						}
					}
					item.SummaryAddresses = append(item.SummaryAddresses, cItem)
					return true
				})
			}
			data.AfInterfaces = append(data.AfInterfaces, item)
			return true
		})
	}
	data.RoutePolicyId = types.StringNull()

	if t := res.Get(path + "tableMap.name.refId.optionType"); t.Exists() {
		va := res.Get(path + "tableMap.name.refId.value")
		if t.String() == "global" {
			data.RoutePolicyId = types.StringValue(va.String())
		}
	}
	data.Filter = types.BoolNull()
	data.FilterVariable = types.StringNull()
	if t := res.Get(path + "tableMap.filter.optionType"); t.Exists() {
		va := res.Get(path + "tableMap.filter.value")
		if t.String() == "variable" {
			data.FilterVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Filter = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ServiceRoutingEIGRP) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.AutonomousSystemId = types.Int64Null()
	data.AutonomousSystemIdVariable = types.StringNull()
	if t := res.Get(path + "asNum.optionType"); t.Exists() {
		va := res.Get(path + "asNum.value")
		if t.String() == "variable" {
			data.AutonomousSystemIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AutonomousSystemId = types.Int64Value(va.Int())
		}
	}
	for i := range data.Redistributes {
		keys := [...]string{"protocol", "routePolicy.refId"}
		keyValues := [...]string{data.Redistributes[i].Protocol.ValueString(), data.Redistributes[i].RoutePolicyId.ValueString()}
		keyValuesVariables := [...]string{data.Redistributes[i].ProtocolVariable.ValueString(), ""}

		var r gjson.Result
		res.Get(path + "addressFamily.redistribute").ForEach(
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
		data.Redistributes[i].Protocol = types.StringNull()
		data.Redistributes[i].ProtocolVariable = types.StringNull()
		if t := r.Get("protocol.optionType"); t.Exists() {
			va := r.Get("protocol.value")
			if t.String() == "variable" {
				data.Redistributes[i].ProtocolVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Redistributes[i].Protocol = types.StringValue(va.String())
			}
		}
		data.Redistributes[i].RoutePolicyId = types.StringNull()

		if t := r.Get("routePolicy.refId.optionType"); t.Exists() {
			va := r.Get("routePolicy.refId.value")
			if t.String() == "global" {
				data.Redistributes[i].RoutePolicyId = types.StringValue(va.String())
			}
		}
	}
	for i := range data.Networks {
		keys := [...]string{"prefix.address", "prefix.mask"}
		keyValues := [...]string{data.Networks[i].IpAddress.ValueString(), data.Networks[i].Mask.ValueString()}
		keyValuesVariables := [...]string{data.Networks[i].IpAddressVariable.ValueString(), data.Networks[i].MaskVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "addressFamily.network").ForEach(
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
		data.Networks[i].IpAddress = types.StringNull()
		data.Networks[i].IpAddressVariable = types.StringNull()
		if t := r.Get("prefix.address.optionType"); t.Exists() {
			va := r.Get("prefix.address.value")
			if t.String() == "variable" {
				data.Networks[i].IpAddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Networks[i].IpAddress = types.StringValue(va.String())
			}
		}
		data.Networks[i].Mask = types.StringNull()
		data.Networks[i].MaskVariable = types.StringNull()
		if t := r.Get("prefix.mask.optionType"); t.Exists() {
			va := r.Get("prefix.mask.value")
			if t.String() == "variable" {
				data.Networks[i].MaskVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Networks[i].Mask = types.StringValue(va.String())
			}
		}
	}
	data.HelloInterval = types.Int64Null()
	data.HelloIntervalVariable = types.StringNull()
	if t := res.Get(path + "helloInterval.optionType"); t.Exists() {
		va := res.Get(path + "helloInterval.value")
		if t.String() == "variable" {
			data.HelloIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.HelloInterval = types.Int64Value(va.Int())
		}
	}
	data.HoldTime = types.Int64Null()
	data.HoldTimeVariable = types.StringNull()
	if t := res.Get(path + "holdTime.optionType"); t.Exists() {
		va := res.Get(path + "holdTime.value")
		if t.String() == "variable" {
			data.HoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.HoldTime = types.Int64Value(va.Int())
		}
	}
	data.Type = types.StringNull()
	data.TypeVariable = types.StringNull()
	if t := res.Get(path + "authentication.type.optionType"); t.Exists() {
		va := res.Get(path + "authentication.type.value")
		if t.String() == "variable" {
			data.TypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Type = types.StringValue(va.String())
		}
	}
	for i := range data.Md5Keys {
		keys := [...]string{"keyId"}
		keyValues := [...]string{strconv.FormatInt(data.Md5Keys[i].KeyId.ValueInt64(), 10)}
		keyValuesVariables := [...]string{data.Md5Keys[i].KeyIdVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "authentication.key").ForEach(
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
		data.Md5Keys[i].KeyId = types.Int64Null()
		data.Md5Keys[i].KeyIdVariable = types.StringNull()
		if t := r.Get("keyId.optionType"); t.Exists() {
			va := r.Get("keyId.value")
			if t.String() == "variable" {
				data.Md5Keys[i].KeyIdVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Md5Keys[i].KeyId = types.Int64Value(va.Int())
			}
		}
	}
	for i := range data.AfInterfaces {
		keys := [...]string{"name", "shutdown"}
		keyValues := [...]string{data.AfInterfaces[i].Name.ValueString(), strconv.FormatBool(data.AfInterfaces[i].Shutdown.ValueBool())}
		keyValuesVariables := [...]string{data.AfInterfaces[i].NameVariable.ValueString(), data.AfInterfaces[i].ShutdownVariable.ValueString(), ""}

		var r gjson.Result
		res.Get(path + "afInterface").ForEach(
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
		data.AfInterfaces[i].Name = types.StringNull()
		data.AfInterfaces[i].NameVariable = types.StringNull()
		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "variable" {
				data.AfInterfaces[i].NameVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.AfInterfaces[i].Name = types.StringValue(va.String())
			}
		}
		data.AfInterfaces[i].Shutdown = types.BoolNull()
		data.AfInterfaces[i].ShutdownVariable = types.StringNull()
		if t := r.Get("shutdown.optionType"); t.Exists() {
			va := r.Get("shutdown.value")
			if t.String() == "variable" {
				data.AfInterfaces[i].ShutdownVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.AfInterfaces[i].Shutdown = types.BoolValue(va.Bool())
			}
		}
		for ci := range data.AfInterfaces[i].SummaryAddresses {
			keys := [...]string{"prefix.address", "prefix.mask"}
			keyValues := [...]string{data.AfInterfaces[i].SummaryAddresses[ci].Address.ValueString(), data.AfInterfaces[i].SummaryAddresses[ci].Mask.ValueString()}
			keyValuesVariables := [...]string{data.AfInterfaces[i].SummaryAddresses[ci].AddressVariable.ValueString(), data.AfInterfaces[i].SummaryAddresses[ci].MaskVariable.ValueString()}

			var cr gjson.Result
			r.Get("summaryAddress").ForEach(
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
			data.AfInterfaces[i].SummaryAddresses[ci].Address = types.StringNull()
			data.AfInterfaces[i].SummaryAddresses[ci].AddressVariable = types.StringNull()
			if t := cr.Get("prefix.address.optionType"); t.Exists() {
				va := cr.Get("prefix.address.value")
				if t.String() == "variable" {
					data.AfInterfaces[i].SummaryAddresses[ci].AddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.AfInterfaces[i].SummaryAddresses[ci].Address = types.StringValue(va.String())
				}
			}
			data.AfInterfaces[i].SummaryAddresses[ci].Mask = types.StringNull()
			data.AfInterfaces[i].SummaryAddresses[ci].MaskVariable = types.StringNull()
			if t := cr.Get("prefix.mask.optionType"); t.Exists() {
				va := cr.Get("prefix.mask.value")
				if t.String() == "variable" {
					data.AfInterfaces[i].SummaryAddresses[ci].MaskVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.AfInterfaces[i].SummaryAddresses[ci].Mask = types.StringValue(va.String())
				}
			}
		}
	}
	data.RoutePolicyId = types.StringNull()

	if t := res.Get(path + "tableMap.name.refId.optionType"); t.Exists() {
		va := res.Get(path + "tableMap.name.refId.value")
		if t.String() == "global" {
			data.RoutePolicyId = types.StringValue(va.String())
		}
	}
	data.Filter = types.BoolNull()
	data.FilterVariable = types.StringNull()
	if t := res.Get(path + "tableMap.filter.optionType"); t.Exists() {
		va := res.Get(path + "tableMap.filter.value")
		if t.String() == "variable" {
			data.FilterVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Filter = types.BoolValue(va.Bool())
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *ServiceRoutingEIGRP) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.AutonomousSystemId.IsNull() {
		return false
	}
	if !data.AutonomousSystemIdVariable.IsNull() {
		return false
	}
	if len(data.Redistributes) > 0 {
		return false
	}
	if len(data.Networks) > 0 {
		return false
	}
	if !data.HelloInterval.IsNull() {
		return false
	}
	if !data.HelloIntervalVariable.IsNull() {
		return false
	}
	if !data.HoldTime.IsNull() {
		return false
	}
	if !data.HoldTimeVariable.IsNull() {
		return false
	}
	if !data.Type.IsNull() {
		return false
	}
	if !data.TypeVariable.IsNull() {
		return false
	}
	if !data.HmacAuthenticationKey.IsNull() {
		return false
	}
	if !data.HmacAuthenticationKeyVariable.IsNull() {
		return false
	}
	if len(data.Md5Keys) > 0 {
		return false
	}
	if len(data.AfInterfaces) > 0 {
		return false
	}
	if !data.RoutePolicyId.IsNull() {
		return false
	}
	if !data.Filter.IsNull() {
		return false
	}
	if !data.FilterVariable.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
