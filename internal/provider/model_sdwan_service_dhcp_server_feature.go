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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ServiceDHCPServer struct {
	Id                     types.String                    `tfsdk:"id"`
	Version                types.Int64                     `tfsdk:"version"`
	Name                   types.String                    `tfsdk:"name"`
	Description            types.String                    `tfsdk:"description"`
	FeatureProfileId       types.String                    `tfsdk:"feature_profile_id"`
	NetworkAddress         types.String                    `tfsdk:"network_address"`
	NetworkAddressVariable types.String                    `tfsdk:"network_address_variable"`
	SubnetMask             types.String                    `tfsdk:"subnet_mask"`
	SubnetMaskVariable     types.String                    `tfsdk:"subnet_mask_variable"`
	Exclude                types.Set                       `tfsdk:"exclude"`
	ExcludeVariable        types.String                    `tfsdk:"exclude_variable"`
	LeaseTime              types.Int64                     `tfsdk:"lease_time"`
	LeaseTimeVariable      types.String                    `tfsdk:"lease_time_variable"`
	InterfaceMtu           types.Int64                     `tfsdk:"interface_mtu"`
	InterfaceMtuVariable   types.String                    `tfsdk:"interface_mtu_variable"`
	DomainName             types.String                    `tfsdk:"domain_name"`
	DomainNameVariable     types.String                    `tfsdk:"domain_name_variable"`
	DefaultGateway         types.String                    `tfsdk:"default_gateway"`
	DefaultGatewayVariable types.String                    `tfsdk:"default_gateway_variable"`
	DnsServers             types.Set                       `tfsdk:"dns_servers"`
	DnsServersVariable     types.String                    `tfsdk:"dns_servers_variable"`
	TftpServers            types.Set                       `tfsdk:"tftp_servers"`
	TftpServersVariable    types.String                    `tfsdk:"tftp_servers_variable"`
	StaticLeases           []ServiceDHCPServerStaticLeases `tfsdk:"static_leases"`
	OptionCodes            []ServiceDHCPServerOptionCodes  `tfsdk:"option_codes"`
}

type ServiceDHCPServerStaticLeases struct {
	MacAddress         types.String `tfsdk:"mac_address"`
	MacAddressVariable types.String `tfsdk:"mac_address_variable"`
	IpAddress          types.String `tfsdk:"ip_address"`
	IpAddressVariable  types.String `tfsdk:"ip_address_variable"`
}

type ServiceDHCPServerOptionCodes struct {
	Code          types.Int64  `tfsdk:"code"`
	CodeVariable  types.String `tfsdk:"code_variable"`
	Ascii         types.String `tfsdk:"ascii"`
	AsciiVariable types.String `tfsdk:"ascii_variable"`
	Hex           types.String `tfsdk:"hex"`
	HexVariable   types.String `tfsdk:"hex_variable"`
	Ip            types.Set    `tfsdk:"ip"`
	IpVariable    types.String `tfsdk:"ip_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceDHCPServer) getModel() string {
	return "service_dhcp_server"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceDHCPServer) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/dhcp-server", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceDHCPServer) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.NetworkAddressVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"addressPool.networkAddress.optionType", "variable")
			body, _ = sjson.Set(body, path+"addressPool.networkAddress.value", data.NetworkAddressVariable.ValueString())
		}
	} else if !data.NetworkAddress.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"addressPool.networkAddress.optionType", "global")
			body, _ = sjson.Set(body, path+"addressPool.networkAddress.value", data.NetworkAddress.ValueString())
		}
	}

	if !data.SubnetMaskVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"addressPool.subnetMask.optionType", "variable")
			body, _ = sjson.Set(body, path+"addressPool.subnetMask.value", data.SubnetMaskVariable.ValueString())
		}
	} else if !data.SubnetMask.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"addressPool.subnetMask.optionType", "global")
			body, _ = sjson.Set(body, path+"addressPool.subnetMask.value", data.SubnetMask.ValueString())
		}
	}

	if !data.ExcludeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"exclude.optionType", "variable")
			body, _ = sjson.Set(body, path+"exclude.value", data.ExcludeVariable.ValueString())
		}
	} else if data.Exclude.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"exclude.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"exclude.optionType", "global")
			var values []string
			data.Exclude.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"exclude.value", values)
		}
	}

	if !data.LeaseTimeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"leaseTime.optionType", "variable")
			body, _ = sjson.Set(body, path+"leaseTime.value", data.LeaseTimeVariable.ValueString())
		}
	} else if data.LeaseTime.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"leaseTime.optionType", "default")
			body, _ = sjson.Set(body, path+"leaseTime.value", 86400)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"leaseTime.optionType", "global")
			body, _ = sjson.Set(body, path+"leaseTime.value", data.LeaseTime.ValueInt64())
		}
	}

	if !data.InterfaceMtuVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"interfaceMtu.optionType", "variable")
			body, _ = sjson.Set(body, path+"interfaceMtu.value", data.InterfaceMtuVariable.ValueString())
		}
	} else if data.InterfaceMtu.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"interfaceMtu.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"interfaceMtu.optionType", "global")
			body, _ = sjson.Set(body, path+"interfaceMtu.value", data.InterfaceMtu.ValueInt64())
		}
	}

	if !data.DomainNameVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"domainName.optionType", "variable")
			body, _ = sjson.Set(body, path+"domainName.value", data.DomainNameVariable.ValueString())
		}
	} else if data.DomainName.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"domainName.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"domainName.optionType", "global")
			body, _ = sjson.Set(body, path+"domainName.value", data.DomainName.ValueString())
		}
	}

	if !data.DefaultGatewayVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"defaultGateway.optionType", "variable")
			body, _ = sjson.Set(body, path+"defaultGateway.value", data.DefaultGatewayVariable.ValueString())
		}
	} else if data.DefaultGateway.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"defaultGateway.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"defaultGateway.optionType", "global")
			body, _ = sjson.Set(body, path+"defaultGateway.value", data.DefaultGateway.ValueString())
		}
	}

	if !data.DnsServersVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dnsServers.optionType", "variable")
			body, _ = sjson.Set(body, path+"dnsServers.value", data.DnsServersVariable.ValueString())
		}
	} else if data.DnsServers.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dnsServers.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"dnsServers.optionType", "global")
			var values []string
			data.DnsServers.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"dnsServers.value", values)
		}
	}

	if !data.TftpServersVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tftpServers.optionType", "variable")
			body, _ = sjson.Set(body, path+"tftpServers.value", data.TftpServersVariable.ValueString())
		}
	} else if data.TftpServers.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"tftpServers.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"tftpServers.optionType", "global")
			var values []string
			data.TftpServers.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"tftpServers.value", values)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"staticLease", []interface{}{})
		for _, item := range data.StaticLeases {
			itemBody := ""

			if !item.MacAddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "macAddress.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "macAddress.value", item.MacAddressVariable.ValueString())
				}
			} else if !item.MacAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "macAddress.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "macAddress.value", item.MacAddress.ValueString())
				}
			}

			if !item.IpAddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ip.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ip.value", item.IpAddressVariable.ValueString())
				}
			} else if !item.IpAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ip.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ip.value", item.IpAddress.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"staticLease.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"optionCode", []interface{}{})
		for _, item := range data.OptionCodes {
			itemBody := ""

			if !item.CodeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "code.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "code.value", item.CodeVariable.ValueString())
				}
			} else if !item.Code.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "code.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "code.value", item.Code.ValueInt64())
				}
			}

			if !item.AsciiVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ascii.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ascii.value", item.AsciiVariable.ValueString())
				}
			} else if !item.Ascii.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ascii.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ascii.value", item.Ascii.ValueString())
				}
			}

			if !item.HexVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "hex.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "hex.value", item.HexVariable.ValueString())
				}
			} else if !item.Hex.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "hex.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "hex.value", item.Hex.ValueString())
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
					var values []string
					item.Ip.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "ip.value", values)
				}
			}
			body, _ = sjson.SetRaw(body, path+"optionCode.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceDHCPServer) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.NetworkAddress = types.StringNull()
	data.NetworkAddressVariable = types.StringNull()
	if t := res.Get(path + "addressPool.networkAddress.optionType"); t.Exists() {
		va := res.Get(path + "addressPool.networkAddress.value")
		if t.String() == "variable" {
			data.NetworkAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NetworkAddress = types.StringValue(va.String())
		}
	}
	data.SubnetMask = types.StringNull()
	data.SubnetMaskVariable = types.StringNull()
	if t := res.Get(path + "addressPool.subnetMask.optionType"); t.Exists() {
		va := res.Get(path + "addressPool.subnetMask.value")
		if t.String() == "variable" {
			data.SubnetMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SubnetMask = types.StringValue(va.String())
		}
	}
	data.Exclude = types.SetNull(types.StringType)
	data.ExcludeVariable = types.StringNull()
	if t := res.Get(path + "exclude.optionType"); t.Exists() {
		va := res.Get(path + "exclude.value")
		if t.String() == "variable" {
			data.ExcludeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Exclude = helpers.GetStringSet(va.Array())
		}
	}
	data.LeaseTime = types.Int64Null()
	data.LeaseTimeVariable = types.StringNull()
	if t := res.Get(path + "leaseTime.optionType"); t.Exists() {
		va := res.Get(path + "leaseTime.value")
		if t.String() == "variable" {
			data.LeaseTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.LeaseTime = types.Int64Value(va.Int())
		}
	}
	data.InterfaceMtu = types.Int64Null()
	data.InterfaceMtuVariable = types.StringNull()
	if t := res.Get(path + "interfaceMtu.optionType"); t.Exists() {
		va := res.Get(path + "interfaceMtu.value")
		if t.String() == "variable" {
			data.InterfaceMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceMtu = types.Int64Value(va.Int())
		}
	}
	data.DomainName = types.StringNull()
	data.DomainNameVariable = types.StringNull()
	if t := res.Get(path + "domainName.optionType"); t.Exists() {
		va := res.Get(path + "domainName.value")
		if t.String() == "variable" {
			data.DomainNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DomainName = types.StringValue(va.String())
		}
	}
	data.DefaultGateway = types.StringNull()
	data.DefaultGatewayVariable = types.StringNull()
	if t := res.Get(path + "defaultGateway.optionType"); t.Exists() {
		va := res.Get(path + "defaultGateway.value")
		if t.String() == "variable" {
			data.DefaultGatewayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultGateway = types.StringValue(va.String())
		}
	}
	data.DnsServers = types.SetNull(types.StringType)
	data.DnsServersVariable = types.StringNull()
	if t := res.Get(path + "dnsServers.optionType"); t.Exists() {
		va := res.Get(path + "dnsServers.value")
		if t.String() == "variable" {
			data.DnsServersVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DnsServers = helpers.GetStringSet(va.Array())
		}
	}
	data.TftpServers = types.SetNull(types.StringType)
	data.TftpServersVariable = types.StringNull()
	if t := res.Get(path + "tftpServers.optionType"); t.Exists() {
		va := res.Get(path + "tftpServers.value")
		if t.String() == "variable" {
			data.TftpServersVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TftpServers = helpers.GetStringSet(va.Array())
		}
	}
	if value := res.Get(path + "staticLease"); value.Exists() {
		data.StaticLeases = make([]ServiceDHCPServerStaticLeases, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceDHCPServerStaticLeases{}
			item.MacAddress = types.StringNull()
			item.MacAddressVariable = types.StringNull()
			if t := v.Get("macAddress.optionType"); t.Exists() {
				va := v.Get("macAddress.value")
				if t.String() == "variable" {
					item.MacAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.MacAddress = types.StringValue(va.String())
				}
			}
			item.IpAddress = types.StringNull()
			item.IpAddressVariable = types.StringNull()
			if t := v.Get("ip.optionType"); t.Exists() {
				va := v.Get("ip.value")
				if t.String() == "variable" {
					item.IpAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.IpAddress = types.StringValue(va.String())
				}
			}
			data.StaticLeases = append(data.StaticLeases, item)
			return true
		})
	}
	if value := res.Get(path + "optionCode"); value.Exists() {
		data.OptionCodes = make([]ServiceDHCPServerOptionCodes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceDHCPServerOptionCodes{}
			item.Code = types.Int64Null()
			item.CodeVariable = types.StringNull()
			if t := v.Get("code.optionType"); t.Exists() {
				va := v.Get("code.value")
				if t.String() == "variable" {
					item.CodeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Code = types.Int64Value(va.Int())
				}
			}
			item.Ascii = types.StringNull()
			item.AsciiVariable = types.StringNull()
			if t := v.Get("ascii.optionType"); t.Exists() {
				va := v.Get("ascii.value")
				if t.String() == "variable" {
					item.AsciiVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Ascii = types.StringValue(va.String())
				}
			}
			item.Hex = types.StringNull()
			item.HexVariable = types.StringNull()
			if t := v.Get("hex.optionType"); t.Exists() {
				va := v.Get("hex.value")
				if t.String() == "variable" {
					item.HexVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Hex = types.StringValue(va.String())
				}
			}
			item.Ip = types.SetNull(types.StringType)
			item.IpVariable = types.StringNull()
			if t := v.Get("ip.optionType"); t.Exists() {
				va := v.Get("ip.value")
				if t.String() == "variable" {
					item.IpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Ip = helpers.GetStringSet(va.Array())
				}
			}
			data.OptionCodes = append(data.OptionCodes, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ServiceDHCPServer) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.NetworkAddress = types.StringNull()
	data.NetworkAddressVariable = types.StringNull()
	if t := res.Get(path + "addressPool.networkAddress.optionType"); t.Exists() {
		va := res.Get(path + "addressPool.networkAddress.value")
		if t.String() == "variable" {
			data.NetworkAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.NetworkAddress = types.StringValue(va.String())
		}
	}
	data.SubnetMask = types.StringNull()
	data.SubnetMaskVariable = types.StringNull()
	if t := res.Get(path + "addressPool.subnetMask.optionType"); t.Exists() {
		va := res.Get(path + "addressPool.subnetMask.value")
		if t.String() == "variable" {
			data.SubnetMaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SubnetMask = types.StringValue(va.String())
		}
	}
	data.Exclude = types.SetNull(types.StringType)
	data.ExcludeVariable = types.StringNull()
	if t := res.Get(path + "exclude.optionType"); t.Exists() {
		va := res.Get(path + "exclude.value")
		if t.String() == "variable" {
			data.ExcludeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Exclude = helpers.GetStringSet(va.Array())
		}
	}
	data.LeaseTime = types.Int64Null()
	data.LeaseTimeVariable = types.StringNull()
	if t := res.Get(path + "leaseTime.optionType"); t.Exists() {
		va := res.Get(path + "leaseTime.value")
		if t.String() == "variable" {
			data.LeaseTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.LeaseTime = types.Int64Value(va.Int())
		}
	}
	data.InterfaceMtu = types.Int64Null()
	data.InterfaceMtuVariable = types.StringNull()
	if t := res.Get(path + "interfaceMtu.optionType"); t.Exists() {
		va := res.Get(path + "interfaceMtu.value")
		if t.String() == "variable" {
			data.InterfaceMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceMtu = types.Int64Value(va.Int())
		}
	}
	data.DomainName = types.StringNull()
	data.DomainNameVariable = types.StringNull()
	if t := res.Get(path + "domainName.optionType"); t.Exists() {
		va := res.Get(path + "domainName.value")
		if t.String() == "variable" {
			data.DomainNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DomainName = types.StringValue(va.String())
		}
	}
	data.DefaultGateway = types.StringNull()
	data.DefaultGatewayVariable = types.StringNull()
	if t := res.Get(path + "defaultGateway.optionType"); t.Exists() {
		va := res.Get(path + "defaultGateway.value")
		if t.String() == "variable" {
			data.DefaultGatewayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DefaultGateway = types.StringValue(va.String())
		}
	}
	data.DnsServers = types.SetNull(types.StringType)
	data.DnsServersVariable = types.StringNull()
	if t := res.Get(path + "dnsServers.optionType"); t.Exists() {
		va := res.Get(path + "dnsServers.value")
		if t.String() == "variable" {
			data.DnsServersVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DnsServers = helpers.GetStringSet(va.Array())
		}
	}
	data.TftpServers = types.SetNull(types.StringType)
	data.TftpServersVariable = types.StringNull()
	if t := res.Get(path + "tftpServers.optionType"); t.Exists() {
		va := res.Get(path + "tftpServers.value")
		if t.String() == "variable" {
			data.TftpServersVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TftpServers = helpers.GetStringSet(va.Array())
		}
	}
	for i := range data.StaticLeases {
		keys := [...]string{"macAddress", "ip"}
		keyValues := [...]string{data.StaticLeases[i].MacAddress.ValueString(), data.StaticLeases[i].IpAddress.ValueString()}
		keyValuesVariables := [...]string{data.StaticLeases[i].MacAddressVariable.ValueString(), data.StaticLeases[i].IpAddressVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "staticLease").ForEach(
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
		data.StaticLeases[i].MacAddress = types.StringNull()
		data.StaticLeases[i].MacAddressVariable = types.StringNull()
		if t := r.Get("macAddress.optionType"); t.Exists() {
			va := r.Get("macAddress.value")
			if t.String() == "variable" {
				data.StaticLeases[i].MacAddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.StaticLeases[i].MacAddress = types.StringValue(va.String())
			}
		}
		data.StaticLeases[i].IpAddress = types.StringNull()
		data.StaticLeases[i].IpAddressVariable = types.StringNull()
		if t := r.Get("ip.optionType"); t.Exists() {
			va := r.Get("ip.value")
			if t.String() == "variable" {
				data.StaticLeases[i].IpAddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.StaticLeases[i].IpAddress = types.StringValue(va.String())
			}
		}
	}
	for i := range data.OptionCodes {
		keys := [...]string{"code", "ascii", "hex"}
		keyValues := [...]string{strconv.FormatInt(data.OptionCodes[i].Code.ValueInt64(), 10), data.OptionCodes[i].Ascii.ValueString(), data.OptionCodes[i].Hex.ValueString()}
		keyValuesVariables := [...]string{data.OptionCodes[i].CodeVariable.ValueString(), data.OptionCodes[i].AsciiVariable.ValueString(), data.OptionCodes[i].HexVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "optionCode").ForEach(
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
		data.OptionCodes[i].Code = types.Int64Null()
		data.OptionCodes[i].CodeVariable = types.StringNull()
		if t := r.Get("code.optionType"); t.Exists() {
			va := r.Get("code.value")
			if t.String() == "variable" {
				data.OptionCodes[i].CodeVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.OptionCodes[i].Code = types.Int64Value(va.Int())
			}
		}
		data.OptionCodes[i].Ascii = types.StringNull()
		data.OptionCodes[i].AsciiVariable = types.StringNull()
		if t := r.Get("ascii.optionType"); t.Exists() {
			va := r.Get("ascii.value")
			if t.String() == "variable" {
				data.OptionCodes[i].AsciiVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.OptionCodes[i].Ascii = types.StringValue(va.String())
			}
		}
		data.OptionCodes[i].Hex = types.StringNull()
		data.OptionCodes[i].HexVariable = types.StringNull()
		if t := r.Get("hex.optionType"); t.Exists() {
			va := r.Get("hex.value")
			if t.String() == "variable" {
				data.OptionCodes[i].HexVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.OptionCodes[i].Hex = types.StringValue(va.String())
			}
		}
		data.OptionCodes[i].Ip = types.SetNull(types.StringType)
		data.OptionCodes[i].IpVariable = types.StringNull()
		if t := r.Get("ip.optionType"); t.Exists() {
			va := r.Get("ip.value")
			if t.String() == "variable" {
				data.OptionCodes[i].IpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.OptionCodes[i].Ip = helpers.GetStringSet(va.Array())
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *ServiceDHCPServer) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.NetworkAddress.IsNull() {
		return false
	}
	if !data.NetworkAddressVariable.IsNull() {
		return false
	}
	if !data.SubnetMask.IsNull() {
		return false
	}
	if !data.SubnetMaskVariable.IsNull() {
		return false
	}
	if !data.Exclude.IsNull() {
		return false
	}
	if !data.ExcludeVariable.IsNull() {
		return false
	}
	if !data.LeaseTime.IsNull() {
		return false
	}
	if !data.LeaseTimeVariable.IsNull() {
		return false
	}
	if !data.InterfaceMtu.IsNull() {
		return false
	}
	if !data.InterfaceMtuVariable.IsNull() {
		return false
	}
	if !data.DomainName.IsNull() {
		return false
	}
	if !data.DomainNameVariable.IsNull() {
		return false
	}
	if !data.DefaultGateway.IsNull() {
		return false
	}
	if !data.DefaultGatewayVariable.IsNull() {
		return false
	}
	if !data.DnsServers.IsNull() {
		return false
	}
	if !data.DnsServersVariable.IsNull() {
		return false
	}
	if !data.TftpServers.IsNull() {
		return false
	}
	if !data.TftpServersVariable.IsNull() {
		return false
	}
	if len(data.StaticLeases) > 0 {
		return false
	}
	if len(data.OptionCodes) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
