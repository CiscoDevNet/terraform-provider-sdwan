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
type SystemGlobal struct {
	Id                         types.String `tfsdk:"id"`
	Version                    types.Int64  `tfsdk:"version"`
	Name                       types.String `tfsdk:"name"`
	Description                types.String `tfsdk:"description"`
	FeatureProfileId           types.String `tfsdk:"feature_profile_id"`
	HttpServer                 types.Bool   `tfsdk:"http_server"`
	HttpServerVariable         types.String `tfsdk:"http_server_variable"`
	HttpsServer                types.Bool   `tfsdk:"https_server"`
	HttpsServerVariable        types.String `tfsdk:"https_server_variable"`
	FtpPassive                 types.Bool   `tfsdk:"ftp_passive"`
	FtpPassiveVariable         types.String `tfsdk:"ftp_passive_variable"`
	DomainLookup               types.Bool   `tfsdk:"domain_lookup"`
	DomainLookupVariable       types.String `tfsdk:"domain_lookup_variable"`
	ArpProxy                   types.Bool   `tfsdk:"arp_proxy"`
	ArpProxyVariable           types.String `tfsdk:"arp_proxy_variable"`
	RshRcp                     types.Bool   `tfsdk:"rsh_rcp"`
	RshRcpVariable             types.String `tfsdk:"rsh_rcp_variable"`
	LineVty                    types.Bool   `tfsdk:"line_vty"`
	LineVtyVariable            types.String `tfsdk:"line_vty_variable"`
	Cdp                        types.Bool   `tfsdk:"cdp"`
	CdpVariable                types.String `tfsdk:"cdp_variable"`
	Lldp                       types.Bool   `tfsdk:"lldp"`
	LldpVariable               types.String `tfsdk:"lldp_variable"`
	SourceInterface            types.String `tfsdk:"source_interface"`
	SourceInterfaceVariable    types.String `tfsdk:"source_interface_variable"`
	TcpKeepalivesIn            types.Bool   `tfsdk:"tcp_keepalives_in"`
	TcpKeepalivesInVariable    types.String `tfsdk:"tcp_keepalives_in_variable"`
	TcpKeepalivesOut           types.Bool   `tfsdk:"tcp_keepalives_out"`
	TcpKeepalivesOutVariable   types.String `tfsdk:"tcp_keepalives_out_variable"`
	TcpSmallServers            types.Bool   `tfsdk:"tcp_small_servers"`
	TcpSmallServersVariable    types.String `tfsdk:"tcp_small_servers_variable"`
	UdpSmallServers            types.Bool   `tfsdk:"udp_small_servers"`
	UdpSmallServersVariable    types.String `tfsdk:"udp_small_servers_variable"`
	ConsoleLogging             types.Bool   `tfsdk:"console_logging"`
	ConsoleLoggingVariable     types.String `tfsdk:"console_logging_variable"`
	IpSourceRouting            types.Bool   `tfsdk:"ip_source_routing"`
	IpSourceRoutingVariable    types.String `tfsdk:"ip_source_routing_variable"`
	VtyLineLogging             types.Bool   `tfsdk:"vty_line_logging"`
	VtyLineLoggingVariable     types.String `tfsdk:"vty_line_logging_variable"`
	SnmpIfindexPersist         types.Bool   `tfsdk:"snmp_ifindex_persist"`
	SnmpIfindexPersistVariable types.String `tfsdk:"snmp_ifindex_persist_variable"`
	IgnoreBootp                types.Bool   `tfsdk:"ignore_bootp"`
	IgnoreBootpVariable        types.String `tfsdk:"ignore_bootp_variable"`
	Nat64UdpTimeout            types.Int64  `tfsdk:"nat64_udp_timeout"`
	Nat64UdpTimeoutVariable    types.String `tfsdk:"nat64_udp_timeout_variable"`
	Nat64TcpTimeout            types.Int64  `tfsdk:"nat64_tcp_timeout"`
	Nat64TcpTimeoutVariable    types.String `tfsdk:"nat64_tcp_timeout_variable"`
	HttpAuthentication         types.String `tfsdk:"http_authentication"`
	HttpAuthenticationVariable types.String `tfsdk:"http_authentication_variable"`
	SshVersion                 types.String `tfsdk:"ssh_version"`
	SshVersionVariable         types.String `tfsdk:"ssh_version_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SystemGlobal) getModel() string {
	return "system_global"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SystemGlobal) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/global", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SystemGlobal) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.HttpServerVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpHttpServer.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpHttpServer.value", data.HttpServerVariable.ValueString())
		}
	} else if data.HttpServer.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpHttpServer.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpHttpServer.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpHttpServer.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpHttpServer.value", data.HttpServer.ValueBool())
		}
	}

	if !data.HttpsServerVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpHttpsServer.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpHttpsServer.value", data.HttpsServerVariable.ValueString())
		}
	} else if data.HttpsServer.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpHttpsServer.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpHttpsServer.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpHttpsServer.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpHttpsServer.value", data.HttpsServer.ValueBool())
		}
	}

	if !data.FtpPassiveVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpFtpPassive.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpFtpPassive.value", data.FtpPassiveVariable.ValueString())
		}
	} else if data.FtpPassive.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpFtpPassive.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpFtpPassive.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpFtpPassive.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpFtpPassive.value", data.FtpPassive.ValueBool())
		}
	}

	if !data.DomainLookupVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpDomainLookup.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpDomainLookup.value", data.DomainLookupVariable.ValueString())
		}
	} else if data.DomainLookup.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpDomainLookup.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpDomainLookup.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpDomainLookup.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpDomainLookup.value", data.DomainLookup.ValueBool())
		}
	}

	if !data.ArpProxyVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpArpProxy.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpArpProxy.value", data.ArpProxyVariable.ValueString())
		}
	} else if data.ArpProxy.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpArpProxy.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpArpProxy.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpArpProxy.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpArpProxy.value", data.ArpProxy.ValueBool())
		}
	}

	if !data.RshRcpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpRcmd.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpRcmd.value", data.RshRcpVariable.ValueString())
		}
	} else if data.RshRcp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpRcmd.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpRcmd.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpRcmd.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpRcmd.value", data.RshRcp.ValueBool())
		}
	}

	if !data.LineVtyVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpLineVty.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpLineVty.value", data.LineVtyVariable.ValueString())
		}
	} else if data.LineVty.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpLineVty.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpLineVty.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpLineVty.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpLineVty.value", data.LineVty.ValueBool())
		}
	}

	if !data.CdpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpCdp.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpCdp.value", data.CdpVariable.ValueString())
		}
	} else if data.Cdp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpCdp.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpCdp.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpCdp.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpCdp.value", data.Cdp.ValueBool())
		}
	}

	if !data.LldpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpLldp.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpLldp.value", data.LldpVariable.ValueString())
		}
	} else if data.Lldp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpLldp.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpLldp.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpLldp.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpLldp.value", data.Lldp.ValueBool())
		}
	}

	if !data.SourceInterfaceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpSourceIntrf.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpSourceIntrf.value", data.SourceInterfaceVariable.ValueString())
		}
	} else if data.SourceInterface.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpSourceIntrf.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpSourceIntrf.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.servicesGlobalServicesIpSourceIntrf.value", data.SourceInterface.ValueString())
		}
	}

	if !data.TcpKeepalivesInVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpKeepalivesIn.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpKeepalivesIn.value", data.TcpKeepalivesInVariable.ValueString())
		}
	} else if data.TcpKeepalivesIn.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpKeepalivesIn.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpKeepalivesIn.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpKeepalivesIn.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpKeepalivesIn.value", data.TcpKeepalivesIn.ValueBool())
		}
	}

	if !data.TcpKeepalivesOutVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpKeepalivesOut.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpKeepalivesOut.value", data.TcpKeepalivesOutVariable.ValueString())
		}
	} else if data.TcpKeepalivesOut.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpKeepalivesOut.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpKeepalivesOut.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpKeepalivesOut.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpKeepalivesOut.value", data.TcpKeepalivesOut.ValueBool())
		}
	}

	if !data.TcpSmallServersVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpSmallServers.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpSmallServers.value", data.TcpSmallServersVariable.ValueString())
		}
	} else if data.TcpSmallServers.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpSmallServers.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpSmallServers.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpSmallServers.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsTcpSmallServers.value", data.TcpSmallServers.ValueBool())
		}
	}

	if !data.UdpSmallServersVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsUdpSmallServers.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsUdpSmallServers.value", data.UdpSmallServersVariable.ValueString())
		}
	} else if data.UdpSmallServers.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsUdpSmallServers.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsUdpSmallServers.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsUdpSmallServers.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsUdpSmallServers.value", data.UdpSmallServers.ValueBool())
		}
	}

	if !data.ConsoleLoggingVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsConsoleLogging.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsConsoleLogging.value", data.ConsoleLoggingVariable.ValueString())
		}
	} else if data.ConsoleLogging.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsConsoleLogging.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsConsoleLogging.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsConsoleLogging.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsConsoleLogging.value", data.ConsoleLogging.ValueBool())
		}
	}

	if !data.IpSourceRoutingVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsIPSourceRoute.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsIPSourceRoute.value", data.IpSourceRoutingVariable.ValueString())
		}
	} else if data.IpSourceRouting.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsIPSourceRoute.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsIPSourceRoute.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsIPSourceRoute.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsIPSourceRoute.value", data.IpSourceRouting.ValueBool())
		}
	}

	if !data.VtyLineLoggingVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsVtyLineLogging.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsVtyLineLogging.value", data.VtyLineLoggingVariable.ValueString())
		}
	} else if data.VtyLineLogging.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsVtyLineLogging.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsVtyLineLogging.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsVtyLineLogging.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsVtyLineLogging.value", data.VtyLineLogging.ValueBool())
		}
	}

	if !data.SnmpIfindexPersistVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsSnmpIfindexPersist.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsSnmpIfindexPersist.value", data.SnmpIfindexPersistVariable.ValueString())
		}
	} else if data.SnmpIfindexPersist.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsSnmpIfindexPersist.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsSnmpIfindexPersist.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsSnmpIfindexPersist.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsSnmpIfindexPersist.value", data.SnmpIfindexPersist.ValueBool())
		}
	}

	if !data.IgnoreBootpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsIgnoreBootp.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsIgnoreBootp.value", data.IgnoreBootpVariable.ValueString())
		}
	} else if data.IgnoreBootp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsIgnoreBootp.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsIgnoreBootp.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsIgnoreBootp.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalOtherSettingsIgnoreBootp.value", data.IgnoreBootp.ValueBool())
		}
	}

	if !data.Nat64UdpTimeoutVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsNat64UdpTimeout.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsNat64UdpTimeout.value", data.Nat64UdpTimeoutVariable.ValueString())
		}
	} else if data.Nat64UdpTimeout.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsNat64UdpTimeout.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsNat64UdpTimeout.value", 300)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsNat64UdpTimeout.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsNat64UdpTimeout.value", data.Nat64UdpTimeout.ValueInt64())
		}
	}

	if !data.Nat64TcpTimeoutVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsNat64TcpTimeout.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsNat64TcpTimeout.value", data.Nat64TcpTimeoutVariable.ValueString())
		}
	} else if data.Nat64TcpTimeout.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsNat64TcpTimeout.optionType", "default")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsNat64TcpTimeout.value", 3600)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsNat64TcpTimeout.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsNat64TcpTimeout.value", data.Nat64TcpTimeout.ValueInt64())
		}
	}

	if !data.HttpAuthenticationVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsHttpAuthentication.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsHttpAuthentication.value", data.HttpAuthenticationVariable.ValueString())
		}
	} else if data.HttpAuthentication.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsHttpAuthentication.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsHttpAuthentication.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsHttpAuthentication.value", data.HttpAuthentication.ValueString())
		}
	}

	if !data.SshVersionVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsSSHVersion.optionType", "variable")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsSSHVersion.value", data.SshVersionVariable.ValueString())
		}
	} else if data.SshVersion.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsSSHVersion.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsSSHVersion.optionType", "global")
			body, _ = sjson.Set(body, path+"services_global.services_ip.globalSettingsSSHVersion.value", data.SshVersion.ValueString())
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemGlobal) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.HttpServer = types.BoolNull()
	data.HttpServerVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpHttpServer.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpHttpServer.value")
		if t.String() == "variable" {
			data.HttpServerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.HttpServer = types.BoolValue(va.Bool())
		}
	}
	data.HttpsServer = types.BoolNull()
	data.HttpsServerVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpHttpsServer.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpHttpsServer.value")
		if t.String() == "variable" {
			data.HttpsServerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.HttpsServer = types.BoolValue(va.Bool())
		}
	}
	data.FtpPassive = types.BoolNull()
	data.FtpPassiveVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpFtpPassive.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpFtpPassive.value")
		if t.String() == "variable" {
			data.FtpPassiveVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.FtpPassive = types.BoolValue(va.Bool())
		}
	}
	data.DomainLookup = types.BoolNull()
	data.DomainLookupVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpDomainLookup.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpDomainLookup.value")
		if t.String() == "variable" {
			data.DomainLookupVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DomainLookup = types.BoolValue(va.Bool())
		}
	}
	data.ArpProxy = types.BoolNull()
	data.ArpProxyVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpArpProxy.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpArpProxy.value")
		if t.String() == "variable" {
			data.ArpProxyVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ArpProxy = types.BoolValue(va.Bool())
		}
	}
	data.RshRcp = types.BoolNull()
	data.RshRcpVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpRcmd.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpRcmd.value")
		if t.String() == "variable" {
			data.RshRcpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RshRcp = types.BoolValue(va.Bool())
		}
	}
	data.LineVty = types.BoolNull()
	data.LineVtyVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpLineVty.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpLineVty.value")
		if t.String() == "variable" {
			data.LineVtyVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.LineVty = types.BoolValue(va.Bool())
		}
	}
	data.Cdp = types.BoolNull()
	data.CdpVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpCdp.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpCdp.value")
		if t.String() == "variable" {
			data.CdpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Cdp = types.BoolValue(va.Bool())
		}
	}
	data.Lldp = types.BoolNull()
	data.LldpVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpLldp.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpLldp.value")
		if t.String() == "variable" {
			data.LldpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Lldp = types.BoolValue(va.Bool())
		}
	}
	data.SourceInterface = types.StringNull()
	data.SourceInterfaceVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpSourceIntrf.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpSourceIntrf.value")
		if t.String() == "variable" {
			data.SourceInterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SourceInterface = types.StringValue(va.String())
		}
	}
	data.TcpKeepalivesIn = types.BoolNull()
	data.TcpKeepalivesInVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsTcpKeepalivesIn.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsTcpKeepalivesIn.value")
		if t.String() == "variable" {
			data.TcpKeepalivesInVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TcpKeepalivesIn = types.BoolValue(va.Bool())
		}
	}
	data.TcpKeepalivesOut = types.BoolNull()
	data.TcpKeepalivesOutVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsTcpKeepalivesOut.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsTcpKeepalivesOut.value")
		if t.String() == "variable" {
			data.TcpKeepalivesOutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TcpKeepalivesOut = types.BoolValue(va.Bool())
		}
	}
	data.TcpSmallServers = types.BoolNull()
	data.TcpSmallServersVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsTcpSmallServers.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsTcpSmallServers.value")
		if t.String() == "variable" {
			data.TcpSmallServersVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TcpSmallServers = types.BoolValue(va.Bool())
		}
	}
	data.UdpSmallServers = types.BoolNull()
	data.UdpSmallServersVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsUdpSmallServers.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsUdpSmallServers.value")
		if t.String() == "variable" {
			data.UdpSmallServersVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.UdpSmallServers = types.BoolValue(va.Bool())
		}
	}
	data.ConsoleLogging = types.BoolNull()
	data.ConsoleLoggingVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsConsoleLogging.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsConsoleLogging.value")
		if t.String() == "variable" {
			data.ConsoleLoggingVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ConsoleLogging = types.BoolValue(va.Bool())
		}
	}
	data.IpSourceRouting = types.BoolNull()
	data.IpSourceRoutingVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsIPSourceRoute.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsIPSourceRoute.value")
		if t.String() == "variable" {
			data.IpSourceRoutingVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpSourceRouting = types.BoolValue(va.Bool())
		}
	}
	data.VtyLineLogging = types.BoolNull()
	data.VtyLineLoggingVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsVtyLineLogging.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsVtyLineLogging.value")
		if t.String() == "variable" {
			data.VtyLineLoggingVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.VtyLineLogging = types.BoolValue(va.Bool())
		}
	}
	data.SnmpIfindexPersist = types.BoolNull()
	data.SnmpIfindexPersistVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsSnmpIfindexPersist.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsSnmpIfindexPersist.value")
		if t.String() == "variable" {
			data.SnmpIfindexPersistVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SnmpIfindexPersist = types.BoolValue(va.Bool())
		}
	}
	data.IgnoreBootp = types.BoolNull()
	data.IgnoreBootpVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsIgnoreBootp.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsIgnoreBootp.value")
		if t.String() == "variable" {
			data.IgnoreBootpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IgnoreBootp = types.BoolValue(va.Bool())
		}
	}
	data.Nat64UdpTimeout = types.Int64Null()
	data.Nat64UdpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalSettingsNat64UdpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalSettingsNat64UdpTimeout.value")
		if t.String() == "variable" {
			data.Nat64UdpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Nat64UdpTimeout = types.Int64Value(va.Int())
		}
	}
	data.Nat64TcpTimeout = types.Int64Null()
	data.Nat64TcpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalSettingsNat64TcpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalSettingsNat64TcpTimeout.value")
		if t.String() == "variable" {
			data.Nat64TcpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Nat64TcpTimeout = types.Int64Value(va.Int())
		}
	}
	data.HttpAuthentication = types.StringNull()
	data.HttpAuthenticationVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalSettingsHttpAuthentication.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalSettingsHttpAuthentication.value")
		if t.String() == "variable" {
			data.HttpAuthenticationVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.HttpAuthentication = types.StringValue(va.String())
		}
	}
	data.SshVersion = types.StringNull()
	data.SshVersionVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalSettingsSSHVersion.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalSettingsSSHVersion.value")
		if t.String() == "variable" {
			data.SshVersionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SshVersion = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SystemGlobal) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.HttpServer = types.BoolNull()
	data.HttpServerVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpHttpServer.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpHttpServer.value")
		if t.String() == "variable" {
			data.HttpServerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.HttpServer = types.BoolValue(va.Bool())
		}
	}
	data.HttpsServer = types.BoolNull()
	data.HttpsServerVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpHttpsServer.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpHttpsServer.value")
		if t.String() == "variable" {
			data.HttpsServerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.HttpsServer = types.BoolValue(va.Bool())
		}
	}
	data.FtpPassive = types.BoolNull()
	data.FtpPassiveVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpFtpPassive.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpFtpPassive.value")
		if t.String() == "variable" {
			data.FtpPassiveVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.FtpPassive = types.BoolValue(va.Bool())
		}
	}
	data.DomainLookup = types.BoolNull()
	data.DomainLookupVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpDomainLookup.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpDomainLookup.value")
		if t.String() == "variable" {
			data.DomainLookupVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DomainLookup = types.BoolValue(va.Bool())
		}
	}
	data.ArpProxy = types.BoolNull()
	data.ArpProxyVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpArpProxy.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpArpProxy.value")
		if t.String() == "variable" {
			data.ArpProxyVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ArpProxy = types.BoolValue(va.Bool())
		}
	}
	data.RshRcp = types.BoolNull()
	data.RshRcpVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpRcmd.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpRcmd.value")
		if t.String() == "variable" {
			data.RshRcpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RshRcp = types.BoolValue(va.Bool())
		}
	}
	data.LineVty = types.BoolNull()
	data.LineVtyVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpLineVty.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpLineVty.value")
		if t.String() == "variable" {
			data.LineVtyVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.LineVty = types.BoolValue(va.Bool())
		}
	}
	data.Cdp = types.BoolNull()
	data.CdpVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpCdp.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpCdp.value")
		if t.String() == "variable" {
			data.CdpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Cdp = types.BoolValue(va.Bool())
		}
	}
	data.Lldp = types.BoolNull()
	data.LldpVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpLldp.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpLldp.value")
		if t.String() == "variable" {
			data.LldpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Lldp = types.BoolValue(va.Bool())
		}
	}
	data.SourceInterface = types.StringNull()
	data.SourceInterfaceVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpSourceIntrf.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.servicesGlobalServicesIpSourceIntrf.value")
		if t.String() == "variable" {
			data.SourceInterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SourceInterface = types.StringValue(va.String())
		}
	}
	data.TcpKeepalivesIn = types.BoolNull()
	data.TcpKeepalivesInVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsTcpKeepalivesIn.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsTcpKeepalivesIn.value")
		if t.String() == "variable" {
			data.TcpKeepalivesInVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TcpKeepalivesIn = types.BoolValue(va.Bool())
		}
	}
	data.TcpKeepalivesOut = types.BoolNull()
	data.TcpKeepalivesOutVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsTcpKeepalivesOut.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsTcpKeepalivesOut.value")
		if t.String() == "variable" {
			data.TcpKeepalivesOutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TcpKeepalivesOut = types.BoolValue(va.Bool())
		}
	}
	data.TcpSmallServers = types.BoolNull()
	data.TcpSmallServersVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsTcpSmallServers.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsTcpSmallServers.value")
		if t.String() == "variable" {
			data.TcpSmallServersVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TcpSmallServers = types.BoolValue(va.Bool())
		}
	}
	data.UdpSmallServers = types.BoolNull()
	data.UdpSmallServersVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsUdpSmallServers.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsUdpSmallServers.value")
		if t.String() == "variable" {
			data.UdpSmallServersVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.UdpSmallServers = types.BoolValue(va.Bool())
		}
	}
	data.ConsoleLogging = types.BoolNull()
	data.ConsoleLoggingVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsConsoleLogging.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsConsoleLogging.value")
		if t.String() == "variable" {
			data.ConsoleLoggingVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ConsoleLogging = types.BoolValue(va.Bool())
		}
	}
	data.IpSourceRouting = types.BoolNull()
	data.IpSourceRoutingVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsIPSourceRoute.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsIPSourceRoute.value")
		if t.String() == "variable" {
			data.IpSourceRoutingVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpSourceRouting = types.BoolValue(va.Bool())
		}
	}
	data.VtyLineLogging = types.BoolNull()
	data.VtyLineLoggingVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsVtyLineLogging.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsVtyLineLogging.value")
		if t.String() == "variable" {
			data.VtyLineLoggingVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.VtyLineLogging = types.BoolValue(va.Bool())
		}
	}
	data.SnmpIfindexPersist = types.BoolNull()
	data.SnmpIfindexPersistVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsSnmpIfindexPersist.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsSnmpIfindexPersist.value")
		if t.String() == "variable" {
			data.SnmpIfindexPersistVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SnmpIfindexPersist = types.BoolValue(va.Bool())
		}
	}
	data.IgnoreBootp = types.BoolNull()
	data.IgnoreBootpVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalOtherSettingsIgnoreBootp.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalOtherSettingsIgnoreBootp.value")
		if t.String() == "variable" {
			data.IgnoreBootpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IgnoreBootp = types.BoolValue(va.Bool())
		}
	}
	data.Nat64UdpTimeout = types.Int64Null()
	data.Nat64UdpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalSettingsNat64UdpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalSettingsNat64UdpTimeout.value")
		if t.String() == "variable" {
			data.Nat64UdpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Nat64UdpTimeout = types.Int64Value(va.Int())
		}
	}
	data.Nat64TcpTimeout = types.Int64Null()
	data.Nat64TcpTimeoutVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalSettingsNat64TcpTimeout.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalSettingsNat64TcpTimeout.value")
		if t.String() == "variable" {
			data.Nat64TcpTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Nat64TcpTimeout = types.Int64Value(va.Int())
		}
	}
	data.HttpAuthentication = types.StringNull()
	data.HttpAuthenticationVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalSettingsHttpAuthentication.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalSettingsHttpAuthentication.value")
		if t.String() == "variable" {
			data.HttpAuthenticationVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.HttpAuthentication = types.StringValue(va.String())
		}
	}
	data.SshVersion = types.StringNull()
	data.SshVersionVariable = types.StringNull()
	if t := res.Get(path + "services_global.services_ip.globalSettingsSSHVersion.optionType"); t.Exists() {
		va := res.Get(path + "services_global.services_ip.globalSettingsSSHVersion.value")
		if t.String() == "variable" {
			data.SshVersionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SshVersion = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end updateFromBody
