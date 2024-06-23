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
type TransportWANVPNInterfaceGRE struct {
	Id                                    types.String `tfsdk:"id"`
	Version                               types.Int64  `tfsdk:"version"`
	Name                                  types.String `tfsdk:"name"`
	Description                           types.String `tfsdk:"description"`
	FeatureProfileId                      types.String `tfsdk:"feature_profile_id"`
	TransportWanVpnProfileParcelId        types.String `tfsdk:"transport_wan_vpn_profile_parcel_id"`
	InterfaceName                         types.String `tfsdk:"interface_name"`
	InterfaceNameVariable                 types.String `tfsdk:"interface_name_variable"`
	InterfaceDescription                  types.String `tfsdk:"interface_description"`
	InterfaceDescriptionVariable          types.String `tfsdk:"interface_description_variable"`
	Ipv4Address                           types.String `tfsdk:"ipv4_address"`
	Ipv4AddressVariable                   types.String `tfsdk:"ipv4_address_variable"`
	Mask                                  types.String `tfsdk:"mask"`
	MaskVariable                          types.String `tfsdk:"mask_variable"`
	Shutdown                              types.Bool   `tfsdk:"shutdown"`
	ShutdownVariable                      types.String `tfsdk:"shutdown_variable"`
	TunnelSourceIpAddress                 types.String `tfsdk:"tunnel_source_ip_address"`
	TunnelSourceIpAddressVariable         types.String `tfsdk:"tunnel_source_ip_address_variable"`
	TunnelSourceInterface                 types.String `tfsdk:"tunnel_source_interface"`
	TunnelSourceInterfaceVariable         types.String `tfsdk:"tunnel_source_interface_variable"`
	LoopbackTunnelSourceInterface         types.String `tfsdk:"loopback_tunnel_source_interface"`
	LoopbackTunnelSourceInterfaceVariable types.String `tfsdk:"loopback_tunnel_source_interface_variable"`
	LoopbackTunnelRouteVia                types.String `tfsdk:"loopback_tunnel_route_via"`
	LoopbackTunnelRouteViaVariable        types.String `tfsdk:"loopback_tunnel_route_via_variable"`
	GreDestinationIpAddress               types.String `tfsdk:"gre_destination_ip_address"`
	GreDestinationIpAddressVariable       types.String `tfsdk:"gre_destination_ip_address_variable"`
	IpMtu                                 types.Int64  `tfsdk:"ip_mtu"`
	IpMtuVariable                         types.String `tfsdk:"ip_mtu_variable"`
	TcpMss                                types.Int64  `tfsdk:"tcp_mss"`
	TcpMssVariable                        types.String `tfsdk:"tcp_mss_variable"`
	ClearDontFragment                     types.Bool   `tfsdk:"clear_dont_fragment"`
	ClearDontFragmentVariable             types.String `tfsdk:"clear_dont_fragment_variable"`
	ApplicationTunnelType                 types.String `tfsdk:"application_tunnel_type"`
	ApplicationTunnelTypeVariable         types.String `tfsdk:"application_tunnel_type_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TransportWANVPNInterfaceGRE) getModel() string {
	return "transport_wan_vpn_interface_gre"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TransportWANVPNInterfaceGRE) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/transport/%v/wan/vpn/%s/interface/gre", url.QueryEscape(data.FeatureProfileId.ValueString()), url.QueryEscape(data.TransportWanVpnProfileParcelId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TransportWANVPNInterfaceGRE) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.InterfaceNameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.ifName.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.ifName.value", data.InterfaceNameVariable.ValueString())
	} else if !data.InterfaceName.IsNull() {
		body, _ = sjson.Set(body, path+"basic.ifName.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.ifName.value", data.InterfaceName.ValueString())
	}

	if !data.InterfaceDescriptionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.description.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.description.value", data.InterfaceDescriptionVariable.ValueString())
	} else if data.InterfaceDescription.IsNull() {
		body, _ = sjson.Set(body, path+"basic.description.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"basic.description.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.description.value", data.InterfaceDescription.ValueString())
	}

	if !data.Ipv4AddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.address.address.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.address.address.value", data.Ipv4AddressVariable.ValueString())
	} else if !data.Ipv4Address.IsNull() {
		body, _ = sjson.Set(body, path+"basic.address.address.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.address.address.value", data.Ipv4Address.ValueString())
	}

	if !data.MaskVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.address.mask.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.address.mask.value", data.MaskVariable.ValueString())
	} else if !data.Mask.IsNull() {
		body, _ = sjson.Set(body, path+"basic.address.mask.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.address.mask.value", data.Mask.ValueString())
	}

	if !data.ShutdownVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.shutdown.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.shutdown.value", data.ShutdownVariable.ValueString())
	} else if data.Shutdown.IsNull() {
		body, _ = sjson.Set(body, path+"basic.shutdown.optionType", "default")
		body, _ = sjson.Set(body, path+"basic.shutdown.value", false)
	} else {
		body, _ = sjson.Set(body, path+"basic.shutdown.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.shutdown.value", data.Shutdown.ValueBool())
	}

	if !data.TunnelSourceIpAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIp.tunnelSource.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIp.tunnelSource.value", data.TunnelSourceIpAddressVariable.ValueString())
	} else if !data.TunnelSourceIpAddress.IsNull() {
		body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIp.tunnelSource.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceIp.tunnelSource.value", data.TunnelSourceIpAddress.ValueString())
	}

	if !data.TunnelSourceInterfaceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceNotLoopback.tunnelSourceInterface.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceNotLoopback.tunnelSourceInterface.value", data.TunnelSourceInterfaceVariable.ValueString())
	} else if !data.TunnelSourceInterface.IsNull() {
		body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceNotLoopback.tunnelSourceInterface.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceNotLoopback.tunnelSourceInterface.value", data.TunnelSourceInterface.ValueString())
	}

	if !data.LoopbackTunnelSourceInterfaceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceLoopback.tunnelSourceInterface.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceLoopback.tunnelSourceInterface.value", data.LoopbackTunnelSourceInterfaceVariable.ValueString())
	} else if !data.LoopbackTunnelSourceInterface.IsNull() {
		body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceLoopback.tunnelSourceInterface.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceLoopback.tunnelSourceInterface.value", data.LoopbackTunnelSourceInterface.ValueString())
	}

	if !data.LoopbackTunnelRouteViaVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceLoopback.tunnelRouteVia.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceLoopback.tunnelRouteVia.value", data.LoopbackTunnelRouteViaVariable.ValueString())
	} else if !data.LoopbackTunnelRouteVia.IsNull() {
		body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceLoopback.tunnelRouteVia.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.tunnelSourceType.sourceLoopback.tunnelRouteVia.value", data.LoopbackTunnelRouteVia.ValueString())
	}

	if !data.GreDestinationIpAddressVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.tunnelDestination.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.tunnelDestination.value", data.GreDestinationIpAddressVariable.ValueString())
	} else if !data.GreDestinationIpAddress.IsNull() {
		body, _ = sjson.Set(body, path+"basic.tunnelDestination.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.tunnelDestination.value", data.GreDestinationIpAddress.ValueString())
	}

	if !data.IpMtuVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.mtu.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.mtu.value", data.IpMtuVariable.ValueString())
	} else if data.IpMtu.IsNull() {
		body, _ = sjson.Set(body, path+"basic.mtu.optionType", "default")
		body, _ = sjson.Set(body, path+"basic.mtu.value", 1500)
	} else {
		body, _ = sjson.Set(body, path+"basic.mtu.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.mtu.value", data.IpMtu.ValueInt64())
	}

	if !data.TcpMssVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.tcpMssAdjust.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.tcpMssAdjust.value", data.TcpMssVariable.ValueString())
	} else if data.TcpMss.IsNull() {
		body, _ = sjson.Set(body, path+"basic.tcpMssAdjust.optionType", "default")

	} else {
		body, _ = sjson.Set(body, path+"basic.tcpMssAdjust.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.tcpMssAdjust.value", data.TcpMss.ValueInt64())
	}

	if !data.ClearDontFragmentVariable.IsNull() {
		body, _ = sjson.Set(body, path+"basic.clearDontFragment.optionType", "variable")
		body, _ = sjson.Set(body, path+"basic.clearDontFragment.value", data.ClearDontFragmentVariable.ValueString())
	} else if data.ClearDontFragment.IsNull() {
		body, _ = sjson.Set(body, path+"basic.clearDontFragment.optionType", "default")
		body, _ = sjson.Set(body, path+"basic.clearDontFragment.value", false)
	} else {
		body, _ = sjson.Set(body, path+"basic.clearDontFragment.optionType", "global")
		body, _ = sjson.Set(body, path+"basic.clearDontFragment.value", data.ClearDontFragment.ValueBool())
	}

	if !data.ApplicationTunnelTypeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.application.optionType", "variable")
		body, _ = sjson.Set(body, path+"advanced.application.value", data.ApplicationTunnelTypeVariable.ValueString())
	} else if !data.ApplicationTunnelType.IsNull() {
		body, _ = sjson.Set(body, path+"advanced.application.optionType", "global")
		body, _ = sjson.Set(body, path+"advanced.application.value", data.ApplicationTunnelType.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TransportWANVPNInterfaceGRE) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.InterfaceName = types.StringNull()
	data.InterfaceNameVariable = types.StringNull()
	if t := res.Get(path + "basic.ifName.optionType"); t.Exists() {
		va := res.Get(path + "basic.ifName.value")
		if t.String() == "variable" {
			data.InterfaceNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceName = types.StringValue(va.String())
		}
	}
	data.InterfaceDescription = types.StringNull()
	data.InterfaceDescriptionVariable = types.StringNull()
	if t := res.Get(path + "basic.description.optionType"); t.Exists() {
		va := res.Get(path + "basic.description.value")
		if t.String() == "variable" {
			data.InterfaceDescriptionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceDescription = types.StringValue(va.String())
		}
	}
	data.Ipv4Address = types.StringNull()
	data.Ipv4AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.address.address.optionType"); t.Exists() {
		va := res.Get(path + "basic.address.address.value")
		if t.String() == "variable" {
			data.Ipv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Address = types.StringValue(va.String())
		}
	}
	data.Mask = types.StringNull()
	data.MaskVariable = types.StringNull()
	if t := res.Get(path + "basic.address.mask.optionType"); t.Exists() {
		va := res.Get(path + "basic.address.mask.value")
		if t.String() == "variable" {
			data.MaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Mask = types.StringValue(va.String())
		}
	}
	data.Shutdown = types.BoolNull()
	data.ShutdownVariable = types.StringNull()
	if t := res.Get(path + "basic.shutdown.optionType"); t.Exists() {
		va := res.Get(path + "basic.shutdown.value")
		if t.String() == "variable" {
			data.ShutdownVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Shutdown = types.BoolValue(va.Bool())
		}
	}
	data.TunnelSourceIpAddress = types.StringNull()
	data.TunnelSourceIpAddressVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceIp.tunnelSource.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceIp.tunnelSource.value")
		if t.String() == "variable" {
			data.TunnelSourceIpAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceIpAddress = types.StringValue(va.String())
		}
	}
	data.TunnelSourceInterface = types.StringNull()
	data.TunnelSourceInterfaceVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceNotLoopback.tunnelSourceInterface.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceNotLoopback.tunnelSourceInterface.value")
		if t.String() == "variable" {
			data.TunnelSourceInterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceInterface = types.StringValue(va.String())
		}
	}
	data.LoopbackTunnelSourceInterface = types.StringNull()
	data.LoopbackTunnelSourceInterfaceVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceLoopback.tunnelSourceInterface.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceLoopback.tunnelSourceInterface.value")
		if t.String() == "variable" {
			data.LoopbackTunnelSourceInterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.LoopbackTunnelSourceInterface = types.StringValue(va.String())
		}
	}
	data.LoopbackTunnelRouteVia = types.StringNull()
	data.LoopbackTunnelRouteViaVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceLoopback.tunnelRouteVia.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceLoopback.tunnelRouteVia.value")
		if t.String() == "variable" {
			data.LoopbackTunnelRouteViaVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.LoopbackTunnelRouteVia = types.StringValue(va.String())
		}
	}
	data.GreDestinationIpAddress = types.StringNull()
	data.GreDestinationIpAddressVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelDestination.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelDestination.value")
		if t.String() == "variable" {
			data.GreDestinationIpAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GreDestinationIpAddress = types.StringValue(va.String())
		}
	}
	data.IpMtu = types.Int64Null()
	data.IpMtuVariable = types.StringNull()
	if t := res.Get(path + "basic.mtu.optionType"); t.Exists() {
		va := res.Get(path + "basic.mtu.value")
		if t.String() == "variable" {
			data.IpMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpMtu = types.Int64Value(va.Int())
		}
	}
	data.TcpMss = types.Int64Null()
	data.TcpMssVariable = types.StringNull()
	if t := res.Get(path + "basic.tcpMssAdjust.optionType"); t.Exists() {
		va := res.Get(path + "basic.tcpMssAdjust.value")
		if t.String() == "variable" {
			data.TcpMssVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TcpMss = types.Int64Value(va.Int())
		}
	}
	data.ClearDontFragment = types.BoolNull()
	data.ClearDontFragmentVariable = types.StringNull()
	if t := res.Get(path + "basic.clearDontFragment.optionType"); t.Exists() {
		va := res.Get(path + "basic.clearDontFragment.value")
		if t.String() == "variable" {
			data.ClearDontFragmentVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ClearDontFragment = types.BoolValue(va.Bool())
		}
	}
	data.ApplicationTunnelType = types.StringNull()
	data.ApplicationTunnelTypeVariable = types.StringNull()
	if t := res.Get(path + "advanced.application.optionType"); t.Exists() {
		va := res.Get(path + "advanced.application.value")
		if t.String() == "variable" {
			data.ApplicationTunnelTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ApplicationTunnelType = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TransportWANVPNInterfaceGRE) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.InterfaceName = types.StringNull()
	data.InterfaceNameVariable = types.StringNull()
	if t := res.Get(path + "basic.ifName.optionType"); t.Exists() {
		va := res.Get(path + "basic.ifName.value")
		if t.String() == "variable" {
			data.InterfaceNameVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceName = types.StringValue(va.String())
		}
	}
	data.InterfaceDescription = types.StringNull()
	data.InterfaceDescriptionVariable = types.StringNull()
	if t := res.Get(path + "basic.description.optionType"); t.Exists() {
		va := res.Get(path + "basic.description.value")
		if t.String() == "variable" {
			data.InterfaceDescriptionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InterfaceDescription = types.StringValue(va.String())
		}
	}
	data.Ipv4Address = types.StringNull()
	data.Ipv4AddressVariable = types.StringNull()
	if t := res.Get(path + "basic.address.address.optionType"); t.Exists() {
		va := res.Get(path + "basic.address.address.value")
		if t.String() == "variable" {
			data.Ipv4AddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Address = types.StringValue(va.String())
		}
	}
	data.Mask = types.StringNull()
	data.MaskVariable = types.StringNull()
	if t := res.Get(path + "basic.address.mask.optionType"); t.Exists() {
		va := res.Get(path + "basic.address.mask.value")
		if t.String() == "variable" {
			data.MaskVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Mask = types.StringValue(va.String())
		}
	}
	data.Shutdown = types.BoolNull()
	data.ShutdownVariable = types.StringNull()
	if t := res.Get(path + "basic.shutdown.optionType"); t.Exists() {
		va := res.Get(path + "basic.shutdown.value")
		if t.String() == "variable" {
			data.ShutdownVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Shutdown = types.BoolValue(va.Bool())
		}
	}
	data.TunnelSourceIpAddress = types.StringNull()
	data.TunnelSourceIpAddressVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceIp.tunnelSource.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceIp.tunnelSource.value")
		if t.String() == "variable" {
			data.TunnelSourceIpAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceIpAddress = types.StringValue(va.String())
		}
	}
	data.TunnelSourceInterface = types.StringNull()
	data.TunnelSourceInterfaceVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceNotLoopback.tunnelSourceInterface.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceNotLoopback.tunnelSourceInterface.value")
		if t.String() == "variable" {
			data.TunnelSourceInterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TunnelSourceInterface = types.StringValue(va.String())
		}
	}
	data.LoopbackTunnelSourceInterface = types.StringNull()
	data.LoopbackTunnelSourceInterfaceVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceLoopback.tunnelSourceInterface.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceLoopback.tunnelSourceInterface.value")
		if t.String() == "variable" {
			data.LoopbackTunnelSourceInterfaceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.LoopbackTunnelSourceInterface = types.StringValue(va.String())
		}
	}
	data.LoopbackTunnelRouteVia = types.StringNull()
	data.LoopbackTunnelRouteViaVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelSourceType.sourceLoopback.tunnelRouteVia.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelSourceType.sourceLoopback.tunnelRouteVia.value")
		if t.String() == "variable" {
			data.LoopbackTunnelRouteViaVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.LoopbackTunnelRouteVia = types.StringValue(va.String())
		}
	}
	data.GreDestinationIpAddress = types.StringNull()
	data.GreDestinationIpAddressVariable = types.StringNull()
	if t := res.Get(path + "basic.tunnelDestination.optionType"); t.Exists() {
		va := res.Get(path + "basic.tunnelDestination.value")
		if t.String() == "variable" {
			data.GreDestinationIpAddressVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GreDestinationIpAddress = types.StringValue(va.String())
		}
	}
	data.IpMtu = types.Int64Null()
	data.IpMtuVariable = types.StringNull()
	if t := res.Get(path + "basic.mtu.optionType"); t.Exists() {
		va := res.Get(path + "basic.mtu.value")
		if t.String() == "variable" {
			data.IpMtuVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IpMtu = types.Int64Value(va.Int())
		}
	}
	data.TcpMss = types.Int64Null()
	data.TcpMssVariable = types.StringNull()
	if t := res.Get(path + "basic.tcpMssAdjust.optionType"); t.Exists() {
		va := res.Get(path + "basic.tcpMssAdjust.value")
		if t.String() == "variable" {
			data.TcpMssVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TcpMss = types.Int64Value(va.Int())
		}
	}
	data.ClearDontFragment = types.BoolNull()
	data.ClearDontFragmentVariable = types.StringNull()
	if t := res.Get(path + "basic.clearDontFragment.optionType"); t.Exists() {
		va := res.Get(path + "basic.clearDontFragment.value")
		if t.String() == "variable" {
			data.ClearDontFragmentVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ClearDontFragment = types.BoolValue(va.Bool())
		}
	}
	data.ApplicationTunnelType = types.StringNull()
	data.ApplicationTunnelTypeVariable = types.StringNull()
	if t := res.Get(path + "advanced.application.optionType"); t.Exists() {
		va := res.Get(path + "advanced.application.value")
		if t.String() == "variable" {
			data.ApplicationTunnelTypeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ApplicationTunnelType = types.StringValue(va.String())
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *TransportWANVPNInterfaceGRE) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.TransportWanVpnProfileParcelId.IsNull() {
		return false
	}
	if !data.InterfaceName.IsNull() {
		return false
	}
	if !data.InterfaceNameVariable.IsNull() {
		return false
	}
	if !data.InterfaceDescription.IsNull() {
		return false
	}
	if !data.InterfaceDescriptionVariable.IsNull() {
		return false
	}
	if !data.Ipv4Address.IsNull() {
		return false
	}
	if !data.Ipv4AddressVariable.IsNull() {
		return false
	}
	if !data.Mask.IsNull() {
		return false
	}
	if !data.MaskVariable.IsNull() {
		return false
	}
	if !data.Shutdown.IsNull() {
		return false
	}
	if !data.ShutdownVariable.IsNull() {
		return false
	}
	if !data.TunnelSourceIpAddress.IsNull() {
		return false
	}
	if !data.TunnelSourceIpAddressVariable.IsNull() {
		return false
	}
	if !data.TunnelSourceInterface.IsNull() {
		return false
	}
	if !data.TunnelSourceInterfaceVariable.IsNull() {
		return false
	}
	if !data.LoopbackTunnelSourceInterface.IsNull() {
		return false
	}
	if !data.LoopbackTunnelSourceInterfaceVariable.IsNull() {
		return false
	}
	if !data.LoopbackTunnelRouteVia.IsNull() {
		return false
	}
	if !data.LoopbackTunnelRouteViaVariable.IsNull() {
		return false
	}
	if !data.GreDestinationIpAddress.IsNull() {
		return false
	}
	if !data.GreDestinationIpAddressVariable.IsNull() {
		return false
	}
	if !data.IpMtu.IsNull() {
		return false
	}
	if !data.IpMtuVariable.IsNull() {
		return false
	}
	if !data.TcpMss.IsNull() {
		return false
	}
	if !data.TcpMssVariable.IsNull() {
		return false
	}
	if !data.ClearDontFragment.IsNull() {
		return false
	}
	if !data.ClearDontFragmentVariable.IsNull() {
		return false
	}
	if !data.ApplicationTunnelType.IsNull() {
		return false
	}
	if !data.ApplicationTunnelTypeVariable.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
