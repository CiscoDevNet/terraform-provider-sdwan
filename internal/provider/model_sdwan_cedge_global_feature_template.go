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
type CEdgeGlobal struct {
	Id                         types.String `tfsdk:"id"`
	Version                    types.Int64  `tfsdk:"version"`
	TemplateType               types.String `tfsdk:"template_type"`
	Name                       types.String `tfsdk:"name"`
	Description                types.String `tfsdk:"description"`
	DeviceTypes                types.Set    `tfsdk:"device_types"`
	Nat64UdpTimeout            types.Int64  `tfsdk:"nat64_udp_timeout"`
	Nat64UdpTimeoutVariable    types.String `tfsdk:"nat64_udp_timeout_variable"`
	Nat64TcpTimeout            types.Int64  `tfsdk:"nat64_tcp_timeout"`
	Nat64TcpTimeoutVariable    types.String `tfsdk:"nat64_tcp_timeout_variable"`
	HttpAuthentication         types.String `tfsdk:"http_authentication"`
	HttpAuthenticationVariable types.String `tfsdk:"http_authentication_variable"`
	SshVersion                 types.Int64  `tfsdk:"ssh_version"`
	SshVersionVariable         types.String `tfsdk:"ssh_version_variable"`
	HttpServer                 types.Bool   `tfsdk:"http_server"`
	HttpServerVariable         types.String `tfsdk:"http_server_variable"`
	HttpsServer                types.Bool   `tfsdk:"https_server"`
	HttpsServerVariable        types.String `tfsdk:"https_server_variable"`
	SourceInterface            types.String `tfsdk:"source_interface"`
	SourceInterfaceVariable    types.String `tfsdk:"source_interface_variable"`
	IpSourceRouting            types.Bool   `tfsdk:"ip_source_routing"`
	IpSourceRoutingVariable    types.String `tfsdk:"ip_source_routing_variable"`
	ArpProxy                   types.Bool   `tfsdk:"arp_proxy"`
	ArpProxyVariable           types.String `tfsdk:"arp_proxy_variable"`
	FtpPassive                 types.Bool   `tfsdk:"ftp_passive"`
	FtpPassiveVariable         types.String `tfsdk:"ftp_passive_variable"`
	RshRcp                     types.Bool   `tfsdk:"rsh_rcp"`
	RshRcpVariable             types.String `tfsdk:"rsh_rcp_variable"`
	Bootp                      types.Bool   `tfsdk:"bootp"`
	BootpVariable              types.String `tfsdk:"bootp_variable"`
	DomainLookup               types.Bool   `tfsdk:"domain_lookup"`
	DomainLookupVariable       types.String `tfsdk:"domain_lookup_variable"`
	TcpKeepalivesOut           types.Bool   `tfsdk:"tcp_keepalives_out"`
	TcpKeepalivesOutVariable   types.String `tfsdk:"tcp_keepalives_out_variable"`
	TcpKeepalivesIn            types.Bool   `tfsdk:"tcp_keepalives_in"`
	TcpKeepalivesInVariable    types.String `tfsdk:"tcp_keepalives_in_variable"`
	TcpSmallServers            types.Bool   `tfsdk:"tcp_small_servers"`
	TcpSmallServersVariable    types.String `tfsdk:"tcp_small_servers_variable"`
	UdpSmallServers            types.Bool   `tfsdk:"udp_small_servers"`
	UdpSmallServersVariable    types.String `tfsdk:"udp_small_servers_variable"`
	Lldp                       types.Bool   `tfsdk:"lldp"`
	LldpVariable               types.String `tfsdk:"lldp_variable"`
	Cdp                        types.Bool   `tfsdk:"cdp"`
	CdpVariable                types.String `tfsdk:"cdp_variable"`
	SnmpIfindexPersist         types.Bool   `tfsdk:"snmp_ifindex_persist"`
	SnmpIfindexPersistVariable types.String `tfsdk:"snmp_ifindex_persist_variable"`
	ConsoleLogging             types.Bool   `tfsdk:"console_logging"`
	ConsoleLoggingVariable     types.String `tfsdk:"console_logging_variable"`
	VtyLogging                 types.Bool   `tfsdk:"vty_logging"`
	VtyLoggingVariable         types.String `tfsdk:"vty_logging_variable"`
	LineVty                    types.Bool   `tfsdk:"line_vty"`
	LineVtyVariable            types.String `tfsdk:"line_vty_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CEdgeGlobal) getModel() string {
	return "cedge_global"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CEdgeGlobal) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cedge_global")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.Nat64UdpTimeoutVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nat64-global.nat64-timeout.udp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat64-global.nat64-timeout.udp-timeout."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"nat64-global.nat64-timeout.udp-timeout."+"vipVariableName", data.Nat64UdpTimeoutVariable.ValueString())
	} else if data.Nat64UdpTimeout.IsNull() {
		body, _ = sjson.Set(body, path+"nat64-global.nat64-timeout.udp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat64-global.nat64-timeout.udp-timeout."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"nat64-global.nat64-timeout.udp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat64-global.nat64-timeout.udp-timeout."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat64-global.nat64-timeout.udp-timeout."+"vipValue", data.Nat64UdpTimeout.ValueInt64())
	}

	if !data.Nat64TcpTimeoutVariable.IsNull() {
		body, _ = sjson.Set(body, path+"nat64-global.nat64-timeout.tcp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat64-global.nat64-timeout.tcp-timeout."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"nat64-global.nat64-timeout.tcp-timeout."+"vipVariableName", data.Nat64TcpTimeoutVariable.ValueString())
	} else if data.Nat64TcpTimeout.IsNull() {
		body, _ = sjson.Set(body, path+"nat64-global.nat64-timeout.tcp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat64-global.nat64-timeout.tcp-timeout."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"nat64-global.nat64-timeout.tcp-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"nat64-global.nat64-timeout.tcp-timeout."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"nat64-global.nat64-timeout.tcp-timeout."+"vipValue", data.Nat64TcpTimeout.ValueInt64())
	}

	if !data.HttpAuthenticationVariable.IsNull() {
		body, _ = sjson.Set(body, path+"http-global.http-authentication."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"http-global.http-authentication."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"http-global.http-authentication."+"vipVariableName", data.HttpAuthenticationVariable.ValueString())
	} else if data.HttpAuthentication.IsNull() {
		body, _ = sjson.Set(body, path+"http-global.http-authentication."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"http-global.http-authentication."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"http-global.http-authentication."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"http-global.http-authentication."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"http-global.http-authentication."+"vipValue", data.HttpAuthentication.ValueString())
	}

	if !data.SshVersionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ssh.version."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ssh.version."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ssh.version."+"vipVariableName", data.SshVersionVariable.ValueString())
	} else if data.SshVersion.IsNull() {
		body, _ = sjson.Set(body, path+"ssh.version."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ssh.version."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ssh.version."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"ssh.version."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ssh.version."+"vipValue", data.SshVersion.ValueInt64())
	}

	if !data.HttpServerVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.http-server."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.http-server."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-ip.http-server."+"vipVariableName", data.HttpServerVariable.ValueString())
	} else if data.HttpServer.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.http-server."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.http-server."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-ip.http-server."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.http-server."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-ip.http-server."+"vipValue", strconv.FormatBool(data.HttpServer.ValueBool()))
	}

	if !data.HttpsServerVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.https-server."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.https-server."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-ip.https-server."+"vipVariableName", data.HttpsServerVariable.ValueString())
	} else if data.HttpsServer.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.https-server."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.https-server."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-ip.https-server."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.https-server."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-ip.https-server."+"vipValue", strconv.FormatBool(data.HttpsServer.ValueBool()))
	}

	if !data.SourceInterfaceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.source-intrf."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.source-intrf."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-ip.source-intrf."+"vipVariableName", data.SourceInterfaceVariable.ValueString())
	} else if data.SourceInterface.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.source-intrf."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.source-intrf."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-ip.source-intrf."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.source-intrf."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-ip.source-intrf."+"vipValue", data.SourceInterface.ValueString())
	}

	if !data.IpSourceRoutingVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.source-route."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.source-route."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-other.source-route."+"vipVariableName", data.IpSourceRoutingVariable.ValueString())
	} else if data.IpSourceRouting.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.source-route."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.source-route."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-other.source-route."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.source-route."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-other.source-route."+"vipValue", strconv.FormatBool(data.IpSourceRouting.ValueBool()))
	}

	if !data.ArpProxyVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.arp-proxy."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.arp-proxy."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-ip.arp-proxy."+"vipVariableName", data.ArpProxyVariable.ValueString())
	} else if data.ArpProxy.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.arp-proxy."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.arp-proxy."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-ip.arp-proxy."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.arp-proxy."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-ip.arp-proxy."+"vipValue", strconv.FormatBool(data.ArpProxy.ValueBool()))
	}

	if !data.FtpPassiveVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.ftp-passive."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.ftp-passive."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-ip.ftp-passive."+"vipVariableName", data.FtpPassiveVariable.ValueString())
	} else if data.FtpPassive.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.ftp-passive."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.ftp-passive."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-ip.ftp-passive."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.ftp-passive."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-ip.ftp-passive."+"vipValue", strconv.FormatBool(data.FtpPassive.ValueBool()))
	}

	if !data.RshRcpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.rcmd."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.rcmd."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-ip.rcmd."+"vipVariableName", data.RshRcpVariable.ValueString())
	} else if data.RshRcp.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.rcmd."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.rcmd."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-ip.rcmd."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.rcmd."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-ip.rcmd."+"vipValue", strconv.FormatBool(data.RshRcp.ValueBool()))
	}

	if !data.BootpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.bootp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.bootp."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-other.bootp."+"vipVariableName", data.BootpVariable.ValueString())
	} else if data.Bootp.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.bootp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.bootp."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-other.bootp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.bootp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-other.bootp."+"vipValue", strconv.FormatBool(data.Bootp.ValueBool()))
	}

	if !data.DomainLookupVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.domain-lookup."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.domain-lookup."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-ip.domain-lookup."+"vipVariableName", data.DomainLookupVariable.ValueString())
	} else if data.DomainLookup.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.domain-lookup."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.domain-lookup."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-ip.domain-lookup."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.domain-lookup."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-ip.domain-lookup."+"vipValue", strconv.FormatBool(data.DomainLookup.ValueBool()))
	}

	if !data.TcpKeepalivesOutVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-keepalives-out."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-keepalives-out."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-keepalives-out."+"vipVariableName", data.TcpKeepalivesOutVariable.ValueString())
	} else if data.TcpKeepalivesOut.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-keepalives-out."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-keepalives-out."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-keepalives-out."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-keepalives-out."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-keepalives-out."+"vipValue", strconv.FormatBool(data.TcpKeepalivesOut.ValueBool()))
	}

	if !data.TcpKeepalivesInVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-keepalives-in."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-keepalives-in."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-keepalives-in."+"vipVariableName", data.TcpKeepalivesInVariable.ValueString())
	} else if data.TcpKeepalivesIn.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-keepalives-in."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-keepalives-in."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-keepalives-in."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-keepalives-in."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-keepalives-in."+"vipValue", strconv.FormatBool(data.TcpKeepalivesIn.ValueBool()))
	}

	if !data.TcpSmallServersVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-small-servers."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-small-servers."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-small-servers."+"vipVariableName", data.TcpSmallServersVariable.ValueString())
	} else if data.TcpSmallServers.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-small-servers."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-small-servers."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-small-servers."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-small-servers."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-other.tcp-small-servers."+"vipValue", strconv.FormatBool(data.TcpSmallServers.ValueBool()))
	}

	if !data.UdpSmallServersVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.udp-small-servers."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.udp-small-servers."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-other.udp-small-servers."+"vipVariableName", data.UdpSmallServersVariable.ValueString())
	} else if data.UdpSmallServers.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.udp-small-servers."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.udp-small-servers."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-other.udp-small-servers."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.udp-small-servers."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-other.udp-small-servers."+"vipValue", strconv.FormatBool(data.UdpSmallServers.ValueBool()))
	}

	if !data.LldpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.lldp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.lldp."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-ip.lldp."+"vipVariableName", data.LldpVariable.ValueString())
	} else if data.Lldp.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.lldp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.lldp."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-ip.lldp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.lldp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-ip.lldp."+"vipValue", strconv.FormatBool(data.Lldp.ValueBool()))
	}

	if !data.CdpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.cdp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.cdp."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-ip.cdp."+"vipVariableName", data.CdpVariable.ValueString())
	} else if data.Cdp.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.cdp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.cdp."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-ip.cdp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.cdp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-ip.cdp."+"vipValue", strconv.FormatBool(data.Cdp.ValueBool()))
	}

	if !data.SnmpIfindexPersistVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.snmp-ifindex-persist."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.snmp-ifindex-persist."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-other.snmp-ifindex-persist."+"vipVariableName", data.SnmpIfindexPersistVariable.ValueString())
	} else if data.SnmpIfindexPersist.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.snmp-ifindex-persist."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.snmp-ifindex-persist."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-other.snmp-ifindex-persist."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.snmp-ifindex-persist."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-other.snmp-ifindex-persist."+"vipValue", strconv.FormatBool(data.SnmpIfindexPersist.ValueBool()))
	}

	if !data.ConsoleLoggingVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.console-logging."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.console-logging."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-other.console-logging."+"vipVariableName", data.ConsoleLoggingVariable.ValueString())
	} else if data.ConsoleLogging.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.console-logging."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.console-logging."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-other.console-logging."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.console-logging."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-other.console-logging."+"vipValue", strconv.FormatBool(data.ConsoleLogging.ValueBool()))
	}

	if !data.VtyLoggingVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.vty-logging."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.vty-logging."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-other.vty-logging."+"vipVariableName", data.VtyLoggingVariable.ValueString())
	} else if data.VtyLogging.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-other.vty-logging."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.vty-logging."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-other.vty-logging."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-other.vty-logging."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-other.vty-logging."+"vipValue", strconv.FormatBool(data.VtyLogging.ValueBool()))
	}

	if !data.LineVtyVariable.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.line-vty."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.line-vty."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"services-global.services-ip.line-vty."+"vipVariableName", data.LineVtyVariable.ValueString())
	} else if data.LineVty.IsNull() {
		body, _ = sjson.Set(body, path+"services-global.services-ip.line-vty."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.line-vty."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"services-global.services-ip.line-vty."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"services-global.services-ip.line-vty."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"services-global.services-ip.line-vty."+"vipValue", strconv.FormatBool(data.LineVty.ValueBool()))
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CEdgeGlobal) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "nat64-global.nat64-timeout.udp-timeout.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Nat64UdpTimeout = types.Int64Null()

			v := res.Get(path + "nat64-global.nat64-timeout.udp-timeout.vipVariableName")
			data.Nat64UdpTimeoutVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Nat64UdpTimeout = types.Int64Null()
			data.Nat64UdpTimeoutVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "nat64-global.nat64-timeout.udp-timeout.vipValue")
			data.Nat64UdpTimeout = types.Int64Value(v.Int())
			data.Nat64UdpTimeoutVariable = types.StringNull()
		}
	} else {
		data.Nat64UdpTimeout = types.Int64Null()
		data.Nat64UdpTimeoutVariable = types.StringNull()
	}
	if value := res.Get(path + "nat64-global.nat64-timeout.tcp-timeout.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Nat64TcpTimeout = types.Int64Null()

			v := res.Get(path + "nat64-global.nat64-timeout.tcp-timeout.vipVariableName")
			data.Nat64TcpTimeoutVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Nat64TcpTimeout = types.Int64Null()
			data.Nat64TcpTimeoutVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "nat64-global.nat64-timeout.tcp-timeout.vipValue")
			data.Nat64TcpTimeout = types.Int64Value(v.Int())
			data.Nat64TcpTimeoutVariable = types.StringNull()
		}
	} else {
		data.Nat64TcpTimeout = types.Int64Null()
		data.Nat64TcpTimeoutVariable = types.StringNull()
	}
	if value := res.Get(path + "http-global.http-authentication.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.HttpAuthentication = types.StringNull()

			v := res.Get(path + "http-global.http-authentication.vipVariableName")
			data.HttpAuthenticationVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.HttpAuthentication = types.StringNull()
			data.HttpAuthenticationVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "http-global.http-authentication.vipValue")
			data.HttpAuthentication = types.StringValue(v.String())
			data.HttpAuthenticationVariable = types.StringNull()
		}
	} else {
		data.HttpAuthentication = types.StringNull()
		data.HttpAuthenticationVariable = types.StringNull()
	}
	if value := res.Get(path + "ssh.version.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SshVersion = types.Int64Null()

			v := res.Get(path + "ssh.version.vipVariableName")
			data.SshVersionVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SshVersion = types.Int64Null()
			data.SshVersionVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ssh.version.vipValue")
			data.SshVersion = types.Int64Value(v.Int())
			data.SshVersionVariable = types.StringNull()
		}
	} else {
		data.SshVersion = types.Int64Null()
		data.SshVersionVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-ip.http-server.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.HttpServer = types.BoolNull()

			v := res.Get(path + "services-global.services-ip.http-server.vipVariableName")
			data.HttpServerVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.HttpServer = types.BoolNull()
			data.HttpServerVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-ip.http-server.vipValue")
			data.HttpServer = types.BoolValue(v.Bool())
			data.HttpServerVariable = types.StringNull()
		}
	} else {
		data.HttpServer = types.BoolNull()
		data.HttpServerVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-ip.https-server.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.HttpsServer = types.BoolNull()

			v := res.Get(path + "services-global.services-ip.https-server.vipVariableName")
			data.HttpsServerVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.HttpsServer = types.BoolNull()
			data.HttpsServerVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-ip.https-server.vipValue")
			data.HttpsServer = types.BoolValue(v.Bool())
			data.HttpsServerVariable = types.StringNull()
		}
	} else {
		data.HttpsServer = types.BoolNull()
		data.HttpsServerVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-ip.source-intrf.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SourceInterface = types.StringNull()

			v := res.Get(path + "services-global.services-ip.source-intrf.vipVariableName")
			data.SourceInterfaceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SourceInterface = types.StringNull()
			data.SourceInterfaceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-ip.source-intrf.vipValue")
			data.SourceInterface = types.StringValue(v.String())
			data.SourceInterfaceVariable = types.StringNull()
		}
	} else {
		data.SourceInterface = types.StringNull()
		data.SourceInterfaceVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-other.source-route.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IpSourceRouting = types.BoolNull()

			v := res.Get(path + "services-global.services-other.source-route.vipVariableName")
			data.IpSourceRoutingVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IpSourceRouting = types.BoolNull()
			data.IpSourceRoutingVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-other.source-route.vipValue")
			data.IpSourceRouting = types.BoolValue(v.Bool())
			data.IpSourceRoutingVariable = types.StringNull()
		}
	} else {
		data.IpSourceRouting = types.BoolNull()
		data.IpSourceRoutingVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-ip.arp-proxy.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ArpProxy = types.BoolNull()

			v := res.Get(path + "services-global.services-ip.arp-proxy.vipVariableName")
			data.ArpProxyVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ArpProxy = types.BoolNull()
			data.ArpProxyVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-ip.arp-proxy.vipValue")
			data.ArpProxy = types.BoolValue(v.Bool())
			data.ArpProxyVariable = types.StringNull()
		}
	} else {
		data.ArpProxy = types.BoolNull()
		data.ArpProxyVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-ip.ftp-passive.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.FtpPassive = types.BoolNull()

			v := res.Get(path + "services-global.services-ip.ftp-passive.vipVariableName")
			data.FtpPassiveVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.FtpPassive = types.BoolNull()
			data.FtpPassiveVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-ip.ftp-passive.vipValue")
			data.FtpPassive = types.BoolValue(v.Bool())
			data.FtpPassiveVariable = types.StringNull()
		}
	} else {
		data.FtpPassive = types.BoolNull()
		data.FtpPassiveVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-ip.rcmd.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.RshRcp = types.BoolNull()

			v := res.Get(path + "services-global.services-ip.rcmd.vipVariableName")
			data.RshRcpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.RshRcp = types.BoolNull()
			data.RshRcpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-ip.rcmd.vipValue")
			data.RshRcp = types.BoolValue(v.Bool())
			data.RshRcpVariable = types.StringNull()
		}
	} else {
		data.RshRcp = types.BoolNull()
		data.RshRcpVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-other.bootp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Bootp = types.BoolNull()

			v := res.Get(path + "services-global.services-other.bootp.vipVariableName")
			data.BootpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Bootp = types.BoolNull()
			data.BootpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-other.bootp.vipValue")
			data.Bootp = types.BoolValue(v.Bool())
			data.BootpVariable = types.StringNull()
		}
	} else {
		data.Bootp = types.BoolNull()
		data.BootpVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-ip.domain-lookup.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DomainLookup = types.BoolNull()

			v := res.Get(path + "services-global.services-ip.domain-lookup.vipVariableName")
			data.DomainLookupVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DomainLookup = types.BoolNull()
			data.DomainLookupVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-ip.domain-lookup.vipValue")
			data.DomainLookup = types.BoolValue(v.Bool())
			data.DomainLookupVariable = types.StringNull()
		}
	} else {
		data.DomainLookup = types.BoolNull()
		data.DomainLookupVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-other.tcp-keepalives-out.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TcpKeepalivesOut = types.BoolNull()

			v := res.Get(path + "services-global.services-other.tcp-keepalives-out.vipVariableName")
			data.TcpKeepalivesOutVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TcpKeepalivesOut = types.BoolNull()
			data.TcpKeepalivesOutVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-other.tcp-keepalives-out.vipValue")
			data.TcpKeepalivesOut = types.BoolValue(v.Bool())
			data.TcpKeepalivesOutVariable = types.StringNull()
		}
	} else {
		data.TcpKeepalivesOut = types.BoolNull()
		data.TcpKeepalivesOutVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-other.tcp-keepalives-in.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TcpKeepalivesIn = types.BoolNull()

			v := res.Get(path + "services-global.services-other.tcp-keepalives-in.vipVariableName")
			data.TcpKeepalivesInVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TcpKeepalivesIn = types.BoolNull()
			data.TcpKeepalivesInVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-other.tcp-keepalives-in.vipValue")
			data.TcpKeepalivesIn = types.BoolValue(v.Bool())
			data.TcpKeepalivesInVariable = types.StringNull()
		}
	} else {
		data.TcpKeepalivesIn = types.BoolNull()
		data.TcpKeepalivesInVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-other.tcp-small-servers.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TcpSmallServers = types.BoolNull()

			v := res.Get(path + "services-global.services-other.tcp-small-servers.vipVariableName")
			data.TcpSmallServersVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TcpSmallServers = types.BoolNull()
			data.TcpSmallServersVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-other.tcp-small-servers.vipValue")
			data.TcpSmallServers = types.BoolValue(v.Bool())
			data.TcpSmallServersVariable = types.StringNull()
		}
	} else {
		data.TcpSmallServers = types.BoolNull()
		data.TcpSmallServersVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-other.udp-small-servers.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.UdpSmallServers = types.BoolNull()

			v := res.Get(path + "services-global.services-other.udp-small-servers.vipVariableName")
			data.UdpSmallServersVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.UdpSmallServers = types.BoolNull()
			data.UdpSmallServersVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-other.udp-small-servers.vipValue")
			data.UdpSmallServers = types.BoolValue(v.Bool())
			data.UdpSmallServersVariable = types.StringNull()
		}
	} else {
		data.UdpSmallServers = types.BoolNull()
		data.UdpSmallServersVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-ip.lldp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Lldp = types.BoolNull()

			v := res.Get(path + "services-global.services-ip.lldp.vipVariableName")
			data.LldpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Lldp = types.BoolNull()
			data.LldpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-ip.lldp.vipValue")
			data.Lldp = types.BoolValue(v.Bool())
			data.LldpVariable = types.StringNull()
		}
	} else {
		data.Lldp = types.BoolNull()
		data.LldpVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-ip.cdp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Cdp = types.BoolNull()

			v := res.Get(path + "services-global.services-ip.cdp.vipVariableName")
			data.CdpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Cdp = types.BoolNull()
			data.CdpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-ip.cdp.vipValue")
			data.Cdp = types.BoolValue(v.Bool())
			data.CdpVariable = types.StringNull()
		}
	} else {
		data.Cdp = types.BoolNull()
		data.CdpVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-other.snmp-ifindex-persist.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SnmpIfindexPersist = types.BoolNull()

			v := res.Get(path + "services-global.services-other.snmp-ifindex-persist.vipVariableName")
			data.SnmpIfindexPersistVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SnmpIfindexPersist = types.BoolNull()
			data.SnmpIfindexPersistVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-other.snmp-ifindex-persist.vipValue")
			data.SnmpIfindexPersist = types.BoolValue(v.Bool())
			data.SnmpIfindexPersistVariable = types.StringNull()
		}
	} else {
		data.SnmpIfindexPersist = types.BoolNull()
		data.SnmpIfindexPersistVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-other.console-logging.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ConsoleLogging = types.BoolNull()

			v := res.Get(path + "services-global.services-other.console-logging.vipVariableName")
			data.ConsoleLoggingVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ConsoleLogging = types.BoolNull()
			data.ConsoleLoggingVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-other.console-logging.vipValue")
			data.ConsoleLogging = types.BoolValue(v.Bool())
			data.ConsoleLoggingVariable = types.StringNull()
		}
	} else {
		data.ConsoleLogging = types.BoolNull()
		data.ConsoleLoggingVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-other.vty-logging.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.VtyLogging = types.BoolNull()

			v := res.Get(path + "services-global.services-other.vty-logging.vipVariableName")
			data.VtyLoggingVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.VtyLogging = types.BoolNull()
			data.VtyLoggingVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-other.vty-logging.vipValue")
			data.VtyLogging = types.BoolValue(v.Bool())
			data.VtyLoggingVariable = types.StringNull()
		}
	} else {
		data.VtyLogging = types.BoolNull()
		data.VtyLoggingVariable = types.StringNull()
	}
	if value := res.Get(path + "services-global.services-ip.line-vty.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.LineVty = types.BoolNull()

			v := res.Get(path + "services-global.services-ip.line-vty.vipVariableName")
			data.LineVtyVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.LineVty = types.BoolNull()
			data.LineVtyVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "services-global.services-ip.line-vty.vipValue")
			data.LineVty = types.BoolValue(v.Bool())
			data.LineVtyVariable = types.StringNull()
		}
	} else {
		data.LineVty = types.BoolNull()
		data.LineVtyVariable = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CEdgeGlobal) hasChanges(ctx context.Context, state *CEdgeGlobal) bool {
	hasChanges := false
	if !data.Nat64UdpTimeout.Equal(state.Nat64UdpTimeout) {
		hasChanges = true
	}
	if !data.Nat64TcpTimeout.Equal(state.Nat64TcpTimeout) {
		hasChanges = true
	}
	if !data.HttpAuthentication.Equal(state.HttpAuthentication) {
		hasChanges = true
	}
	if !data.SshVersion.Equal(state.SshVersion) {
		hasChanges = true
	}
	if !data.HttpServer.Equal(state.HttpServer) {
		hasChanges = true
	}
	if !data.HttpsServer.Equal(state.HttpsServer) {
		hasChanges = true
	}
	if !data.SourceInterface.Equal(state.SourceInterface) {
		hasChanges = true
	}
	if !data.IpSourceRouting.Equal(state.IpSourceRouting) {
		hasChanges = true
	}
	if !data.ArpProxy.Equal(state.ArpProxy) {
		hasChanges = true
	}
	if !data.FtpPassive.Equal(state.FtpPassive) {
		hasChanges = true
	}
	if !data.RshRcp.Equal(state.RshRcp) {
		hasChanges = true
	}
	if !data.Bootp.Equal(state.Bootp) {
		hasChanges = true
	}
	if !data.DomainLookup.Equal(state.DomainLookup) {
		hasChanges = true
	}
	if !data.TcpKeepalivesOut.Equal(state.TcpKeepalivesOut) {
		hasChanges = true
	}
	if !data.TcpKeepalivesIn.Equal(state.TcpKeepalivesIn) {
		hasChanges = true
	}
	if !data.TcpSmallServers.Equal(state.TcpSmallServers) {
		hasChanges = true
	}
	if !data.UdpSmallServers.Equal(state.UdpSmallServers) {
		hasChanges = true
	}
	if !data.Lldp.Equal(state.Lldp) {
		hasChanges = true
	}
	if !data.Cdp.Equal(state.Cdp) {
		hasChanges = true
	}
	if !data.SnmpIfindexPersist.Equal(state.SnmpIfindexPersist) {
		hasChanges = true
	}
	if !data.ConsoleLogging.Equal(state.ConsoleLogging) {
		hasChanges = true
	}
	if !data.VtyLogging.Equal(state.VtyLogging) {
		hasChanges = true
	}
	if !data.LineVty.Equal(state.LineVty) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
