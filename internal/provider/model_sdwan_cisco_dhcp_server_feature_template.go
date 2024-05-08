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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type CiscoDHCPServer struct {
	Id                       types.String                  `tfsdk:"id"`
	Version                  types.Int64                   `tfsdk:"version"`
	TemplateType             types.String                  `tfsdk:"template_type"`
	Name                     types.String                  `tfsdk:"name"`
	Description              types.String                  `tfsdk:"description"`
	DeviceTypes              types.Set                     `tfsdk:"device_types"`
	AddressPool              types.String                  `tfsdk:"address_pool"`
	AddressPoolVariable      types.String                  `tfsdk:"address_pool_variable"`
	ExcludeAddresses         types.Set                     `tfsdk:"exclude_addresses"`
	ExcludeAddressesVariable types.String                  `tfsdk:"exclude_addresses_variable"`
	LeaseTime                types.Int64                   `tfsdk:"lease_time"`
	LeaseTimeVariable        types.String                  `tfsdk:"lease_time_variable"`
	InterfaceMtu             types.Int64                   `tfsdk:"interface_mtu"`
	InterfaceMtuVariable     types.String                  `tfsdk:"interface_mtu_variable"`
	DomainName               types.String                  `tfsdk:"domain_name"`
	DomainNameVariable       types.String                  `tfsdk:"domain_name_variable"`
	DefaultGateway           types.String                  `tfsdk:"default_gateway"`
	DefaultGatewayVariable   types.String                  `tfsdk:"default_gateway_variable"`
	DnsServers               types.Set                     `tfsdk:"dns_servers"`
	DnsServersVariable       types.String                  `tfsdk:"dns_servers_variable"`
	TftpServers              types.Set                     `tfsdk:"tftp_servers"`
	TftpServersVariable      types.String                  `tfsdk:"tftp_servers_variable"`
	StaticLeases             []CiscoDHCPServerStaticLeases `tfsdk:"static_leases"`
	Options                  []CiscoDHCPServerOptions      `tfsdk:"options"`
}

type CiscoDHCPServerStaticLeases struct {
	Optional           types.Bool   `tfsdk:"optional"`
	MacAddress         types.String `tfsdk:"mac_address"`
	MacAddressVariable types.String `tfsdk:"mac_address_variable"`
	IpAddress          types.String `tfsdk:"ip_address"`
	IpAddressVariable  types.String `tfsdk:"ip_address_variable"`
	Hostname           types.String `tfsdk:"hostname"`
	HostnameVariable   types.String `tfsdk:"hostname_variable"`
}

type CiscoDHCPServerOptions struct {
	Optional           types.Bool   `tfsdk:"optional"`
	OptionCode         types.Int64  `tfsdk:"option_code"`
	OptionCodeVariable types.String `tfsdk:"option_code_variable"`
	Ascii              types.String `tfsdk:"ascii"`
	AsciiVariable      types.String `tfsdk:"ascii_variable"`
	Hex                types.String `tfsdk:"hex"`
	HexVariable        types.String `tfsdk:"hex_variable"`
	IpAddress          types.Set    `tfsdk:"ip_address"`
	IpAddressVariable  types.String `tfsdk:"ip_address_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoDHCPServer) getModel() string {
	return "cisco_dhcp_server"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CiscoDHCPServer) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_dhcp_server")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.AddressPoolVariable.IsNull() {
		body, _ = sjson.Set(body, path+"address-pool."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"address-pool."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"address-pool."+"vipVariableName", data.AddressPoolVariable.ValueString())
	} else if data.AddressPool.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"address-pool."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"address-pool."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"address-pool."+"vipValue", data.AddressPool.ValueString())
	}

	if !data.ExcludeAddressesVariable.IsNull() {
		body, _ = sjson.Set(body, path+"exclude."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"exclude."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"exclude."+"vipVariableName", data.ExcludeAddressesVariable.ValueString())
	} else if data.ExcludeAddresses.IsNull() {
		body, _ = sjson.Set(body, path+"exclude."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"exclude."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"exclude."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"exclude."+"vipType", "constant")
		var values []string
		data.ExcludeAddresses.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"exclude."+"vipValue", values)
	}

	if !data.LeaseTimeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"lease-time."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"lease-time."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"lease-time."+"vipVariableName", data.LeaseTimeVariable.ValueString())
	} else if data.LeaseTime.IsNull() {
		body, _ = sjson.Set(body, path+"lease-time."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"lease-time."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"lease-time."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"lease-time."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"lease-time."+"vipValue", data.LeaseTime.ValueInt64())
	}

	if !data.InterfaceMtuVariable.IsNull() {
		body, _ = sjson.Set(body, path+"options.interface-mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"options.interface-mtu."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"options.interface-mtu."+"vipVariableName", data.InterfaceMtuVariable.ValueString())
	} else if data.InterfaceMtu.IsNull() {
		body, _ = sjson.Set(body, path+"options.interface-mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"options.interface-mtu."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"options.interface-mtu."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"options.interface-mtu."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"options.interface-mtu."+"vipValue", data.InterfaceMtu.ValueInt64())
	}

	if !data.DomainNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"options.domain-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"options.domain-name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"options.domain-name."+"vipVariableName", data.DomainNameVariable.ValueString())
	} else if data.DomainName.IsNull() {
		body, _ = sjson.Set(body, path+"options.domain-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"options.domain-name."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"options.domain-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"options.domain-name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"options.domain-name."+"vipValue", data.DomainName.ValueString())
	}

	if !data.DefaultGatewayVariable.IsNull() {
		body, _ = sjson.Set(body, path+"options.default-gateway."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"options.default-gateway."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"options.default-gateway."+"vipVariableName", data.DefaultGatewayVariable.ValueString())
	} else if data.DefaultGateway.IsNull() {
		body, _ = sjson.Set(body, path+"options.default-gateway."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"options.default-gateway."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"options.default-gateway."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"options.default-gateway."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"options.default-gateway."+"vipValue", data.DefaultGateway.ValueString())
	}

	if !data.DnsServersVariable.IsNull() {
		body, _ = sjson.Set(body, path+"options.dns-servers."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"options.dns-servers."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"options.dns-servers."+"vipVariableName", data.DnsServersVariable.ValueString())
	} else if data.DnsServers.IsNull() {
		body, _ = sjson.Set(body, path+"options.dns-servers."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"options.dns-servers."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"options.dns-servers."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"options.dns-servers."+"vipType", "constant")
		var values []string
		data.DnsServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"options.dns-servers."+"vipValue", values)
	}

	if !data.TftpServersVariable.IsNull() {
		body, _ = sjson.Set(body, path+"options.tftp-servers."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"options.tftp-servers."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"options.tftp-servers."+"vipVariableName", data.TftpServersVariable.ValueString())
	} else if data.TftpServers.IsNull() {
		body, _ = sjson.Set(body, path+"options.tftp-servers."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"options.tftp-servers."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"options.tftp-servers."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"options.tftp-servers."+"vipType", "constant")
		var values []string
		data.TftpServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"options.tftp-servers."+"vipValue", values)
	}
	if len(data.StaticLeases) > 0 {
		body, _ = sjson.Set(body, path+"static-lease."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"static-lease."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"static-lease."+"vipPrimaryKey", []string{"mac-address"})
		body, _ = sjson.Set(body, path+"static-lease."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"static-lease."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"static-lease."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"static-lease."+"vipPrimaryKey", []string{"mac-address"})
		body, _ = sjson.Set(body, path+"static-lease."+"vipValue", []interface{}{})
	}
	for _, item := range data.StaticLeases {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "mac-address")

		if !item.MacAddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "mac-address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "mac-address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "mac-address."+"vipVariableName", item.MacAddressVariable.ValueString())
		} else if item.MacAddress.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "mac-address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "mac-address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "mac-address."+"vipValue", item.MacAddress.ValueString())
		}
		itemAttributes = append(itemAttributes, "ip")

		if !item.IpAddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipVariableName", item.IpAddressVariable.ValueString())
		} else if item.IpAddress.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipValue", item.IpAddress.ValueString())
		}
		itemAttributes = append(itemAttributes, "host-name")

		if !item.HostnameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "host-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "host-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "host-name."+"vipVariableName", item.HostnameVariable.ValueString())
		} else if item.Hostname.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "host-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "host-name."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "host-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "host-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "host-name."+"vipValue", item.Hostname.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"static-lease."+"vipValue.-1", itemBody)
	}
	if len(data.Options) > 0 {
		body, _ = sjson.Set(body, path+"options.option-code."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"options.option-code."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"options.option-code."+"vipPrimaryKey", []string{"code"})
		body, _ = sjson.Set(body, path+"options.option-code."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"options.option-code."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"options.option-code."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"options.option-code."+"vipPrimaryKey", []string{"code"})
		body, _ = sjson.Set(body, path+"options.option-code."+"vipValue", []interface{}{})
	}
	for _, item := range data.Options {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "code")

		if !item.OptionCodeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "code."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "code."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "code."+"vipVariableName", item.OptionCodeVariable.ValueString())
		} else if item.OptionCode.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "code."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "code."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "code."+"vipValue", item.OptionCode.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "ascii")

		if !item.AsciiVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ascii."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ascii."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ascii."+"vipVariableName", item.AsciiVariable.ValueString())
		} else if item.Ascii.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ascii."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ascii."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ascii."+"vipValue", item.Ascii.ValueString())
		}
		itemAttributes = append(itemAttributes, "hex")

		if !item.HexVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "hex."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "hex."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "hex."+"vipVariableName", item.HexVariable.ValueString())
		} else if item.Hex.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "hex."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "hex."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "hex."+"vipValue", item.Hex.ValueString())
		}
		itemAttributes = append(itemAttributes, "ip")

		if !item.IpAddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipVariableName", item.IpAddressVariable.ValueString())
		} else if item.IpAddress.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipType", "constant")
			var values []string
			item.IpAddress.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipValue", values)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"options.option-code."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoDHCPServer) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "address-pool.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AddressPool = types.StringNull()

			v := res.Get(path + "address-pool.vipVariableName")
			data.AddressPoolVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AddressPool = types.StringNull()
			data.AddressPoolVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "address-pool.vipValue")
			data.AddressPool = types.StringValue(v.String())
			data.AddressPoolVariable = types.StringNull()
		}
	} else {
		data.AddressPool = types.StringNull()
		data.AddressPoolVariable = types.StringNull()
	}
	if value := res.Get(path + "exclude.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.ExcludeAddresses = types.SetNull(types.StringType)

			v := res.Get(path + "exclude.vipVariableName")
			data.ExcludeAddressesVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ExcludeAddresses = types.SetNull(types.StringType)
			data.ExcludeAddressesVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "exclude.vipValue")
			data.ExcludeAddresses = helpers.GetStringSet(v.Array())
			data.ExcludeAddressesVariable = types.StringNull()
		}
	} else {
		data.ExcludeAddresses = types.SetNull(types.StringType)
		data.ExcludeAddressesVariable = types.StringNull()
	}
	if value := res.Get(path + "lease-time.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.LeaseTime = types.Int64Null()

			v := res.Get(path + "lease-time.vipVariableName")
			data.LeaseTimeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.LeaseTime = types.Int64Null()
			data.LeaseTimeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "lease-time.vipValue")
			data.LeaseTime = types.Int64Value(v.Int())
			data.LeaseTimeVariable = types.StringNull()
		}
	} else {
		data.LeaseTime = types.Int64Null()
		data.LeaseTimeVariable = types.StringNull()
	}
	if value := res.Get(path + "options.interface-mtu.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.InterfaceMtu = types.Int64Null()

			v := res.Get(path + "options.interface-mtu.vipVariableName")
			data.InterfaceMtuVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.InterfaceMtu = types.Int64Null()
			data.InterfaceMtuVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "options.interface-mtu.vipValue")
			data.InterfaceMtu = types.Int64Value(v.Int())
			data.InterfaceMtuVariable = types.StringNull()
		}
	} else {
		data.InterfaceMtu = types.Int64Null()
		data.InterfaceMtuVariable = types.StringNull()
	}
	if value := res.Get(path + "options.domain-name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DomainName = types.StringNull()

			v := res.Get(path + "options.domain-name.vipVariableName")
			data.DomainNameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DomainName = types.StringNull()
			data.DomainNameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "options.domain-name.vipValue")
			data.DomainName = types.StringValue(v.String())
			data.DomainNameVariable = types.StringNull()
		}
	} else {
		data.DomainName = types.StringNull()
		data.DomainNameVariable = types.StringNull()
	}
	if value := res.Get(path + "options.default-gateway.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DefaultGateway = types.StringNull()

			v := res.Get(path + "options.default-gateway.vipVariableName")
			data.DefaultGatewayVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DefaultGateway = types.StringNull()
			data.DefaultGatewayVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "options.default-gateway.vipValue")
			data.DefaultGateway = types.StringValue(v.String())
			data.DefaultGatewayVariable = types.StringNull()
		}
	} else {
		data.DefaultGateway = types.StringNull()
		data.DefaultGatewayVariable = types.StringNull()
	}
	if value := res.Get(path + "options.dns-servers.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.DnsServers = types.SetNull(types.StringType)

			v := res.Get(path + "options.dns-servers.vipVariableName")
			data.DnsServersVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DnsServers = types.SetNull(types.StringType)
			data.DnsServersVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "options.dns-servers.vipValue")
			data.DnsServers = helpers.GetStringSet(v.Array())
			data.DnsServersVariable = types.StringNull()
		}
	} else {
		data.DnsServers = types.SetNull(types.StringType)
		data.DnsServersVariable = types.StringNull()
	}
	if value := res.Get(path + "options.tftp-servers.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.TftpServers = types.SetNull(types.StringType)

			v := res.Get(path + "options.tftp-servers.vipVariableName")
			data.TftpServersVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TftpServers = types.SetNull(types.StringType)
			data.TftpServersVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "options.tftp-servers.vipValue")
			data.TftpServers = helpers.GetStringSet(v.Array())
			data.TftpServersVariable = types.StringNull()
		}
	} else {
		data.TftpServers = types.SetNull(types.StringType)
		data.TftpServersVariable = types.StringNull()
	}
	if value := res.Get(path + "static-lease.vipValue"); len(value.Array()) > 0 {
		data.StaticLeases = make([]CiscoDHCPServerStaticLeases, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoDHCPServerStaticLeases{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("mac-address.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.MacAddress = types.StringNull()

					cv := v.Get("mac-address.vipVariableName")
					item.MacAddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.MacAddress = types.StringNull()
					item.MacAddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("mac-address.vipValue")
					item.MacAddress = types.StringValue(cv.String())
					item.MacAddressVariable = types.StringNull()
				}
			} else {
				item.MacAddress = types.StringNull()
				item.MacAddressVariable = types.StringNull()
			}
			if cValue := v.Get("ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.IpAddress = types.StringNull()

					cv := v.Get("ip.vipVariableName")
					item.IpAddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IpAddress = types.StringNull()
					item.IpAddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ip.vipValue")
					item.IpAddress = types.StringValue(cv.String())
					item.IpAddressVariable = types.StringNull()
				}
			} else {
				item.IpAddress = types.StringNull()
				item.IpAddressVariable = types.StringNull()
			}
			if cValue := v.Get("host-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Hostname = types.StringNull()

					cv := v.Get("host-name.vipVariableName")
					item.HostnameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Hostname = types.StringNull()
					item.HostnameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("host-name.vipValue")
					item.Hostname = types.StringValue(cv.String())
					item.HostnameVariable = types.StringNull()
				}
			} else {
				item.Hostname = types.StringNull()
				item.HostnameVariable = types.StringNull()
			}
			data.StaticLeases = append(data.StaticLeases, item)
			return true
		})
	} else {
		if len(data.StaticLeases) > 0 {
			data.StaticLeases = []CiscoDHCPServerStaticLeases{}
		}
	}
	if value := res.Get(path + "options.option-code.vipValue"); len(value.Array()) > 0 {
		data.Options = make([]CiscoDHCPServerOptions, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoDHCPServerOptions{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("code.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.OptionCode = types.Int64Null()

					cv := v.Get("code.vipVariableName")
					item.OptionCodeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.OptionCode = types.Int64Null()
					item.OptionCodeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("code.vipValue")
					item.OptionCode = types.Int64Value(cv.Int())
					item.OptionCodeVariable = types.StringNull()
				}
			} else {
				item.OptionCode = types.Int64Null()
				item.OptionCodeVariable = types.StringNull()
			}
			if cValue := v.Get("ascii.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Ascii = types.StringNull()

					cv := v.Get("ascii.vipVariableName")
					item.AsciiVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Ascii = types.StringNull()
					item.AsciiVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ascii.vipValue")
					item.Ascii = types.StringValue(cv.String())
					item.AsciiVariable = types.StringNull()
				}
			} else {
				item.Ascii = types.StringNull()
				item.AsciiVariable = types.StringNull()
			}
			if cValue := v.Get("hex.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Hex = types.StringNull()

					cv := v.Get("hex.vipVariableName")
					item.HexVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Hex = types.StringNull()
					item.HexVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("hex.vipValue")
					item.Hex = types.StringValue(cv.String())
					item.HexVariable = types.StringNull()
				}
			} else {
				item.Hex = types.StringNull()
				item.HexVariable = types.StringNull()
			}
			if cValue := v.Get("ip.vipType"); len(cValue.Array()) > 0 {
				if cValue.String() == "variableName" {
					item.IpAddress = types.SetNull(types.StringType)

					cv := v.Get("ip.vipVariableName")
					item.IpAddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.IpAddress = types.SetNull(types.StringType)
					item.IpAddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ip.vipValue")
					item.IpAddress = helpers.GetStringSet(cv.Array())
					item.IpAddressVariable = types.StringNull()
				}
			} else {
				item.IpAddress = types.SetNull(types.StringType)
				item.IpAddressVariable = types.StringNull()
			}
			data.Options = append(data.Options, item)
			return true
		})
	} else {
		if len(data.Options) > 0 {
			data.Options = []CiscoDHCPServerOptions{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoDHCPServer) hasChanges(ctx context.Context, state *CiscoDHCPServer) bool {
	hasChanges := false
	if !data.AddressPool.Equal(state.AddressPool) {
		hasChanges = true
	}
	if !data.ExcludeAddresses.Equal(state.ExcludeAddresses) {
		hasChanges = true
	}
	if !data.LeaseTime.Equal(state.LeaseTime) {
		hasChanges = true
	}
	if !data.InterfaceMtu.Equal(state.InterfaceMtu) {
		hasChanges = true
	}
	if !data.DomainName.Equal(state.DomainName) {
		hasChanges = true
	}
	if !data.DefaultGateway.Equal(state.DefaultGateway) {
		hasChanges = true
	}
	if !data.DnsServers.Equal(state.DnsServers) {
		hasChanges = true
	}
	if !data.TftpServers.Equal(state.TftpServers) {
		hasChanges = true
	}
	if len(data.StaticLeases) != len(state.StaticLeases) {
		hasChanges = true
	} else {
		for i := range data.StaticLeases {
			if !data.StaticLeases[i].MacAddress.Equal(state.StaticLeases[i].MacAddress) {
				hasChanges = true
			}
			if !data.StaticLeases[i].IpAddress.Equal(state.StaticLeases[i].IpAddress) {
				hasChanges = true
			}
			if !data.StaticLeases[i].Hostname.Equal(state.StaticLeases[i].Hostname) {
				hasChanges = true
			}
		}
	}
	if len(data.Options) != len(state.Options) {
		hasChanges = true
	} else {
		for i := range data.Options {
			if !data.Options[i].OptionCode.Equal(state.Options[i].OptionCode) {
				hasChanges = true
			}
			if !data.Options[i].Ascii.Equal(state.Options[i].Ascii) {
				hasChanges = true
			}
			if !data.Options[i].Hex.Equal(state.Options[i].Hex) {
				hasChanges = true
			}
			if !data.Options[i].IpAddress.Equal(state.Options[i].IpAddress) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
