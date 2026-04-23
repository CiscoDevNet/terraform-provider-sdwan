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
type ServiceLANVPN struct {
	Id                              types.String                              `tfsdk:"id"`
	Version                         types.Int64                               `tfsdk:"version"`
	Name                            types.String                              `tfsdk:"name"`
	Description                     types.String                              `tfsdk:"description"`
	FeatureProfileId                types.String                              `tfsdk:"feature_profile_id"`
	Vpn                             types.Int64                               `tfsdk:"vpn"`
	VpnVariable                     types.String                              `tfsdk:"vpn_variable"`
	ConfigDescription               types.String                              `tfsdk:"config_description"`
	ConfigDescriptionVariable       types.String                              `tfsdk:"config_description_variable"`
	OmpAdminDistanceIpv4            types.Int64                               `tfsdk:"omp_admin_distance_ipv4"`
	OmpAdminDistanceIpv4Variable    types.String                              `tfsdk:"omp_admin_distance_ipv4_variable"`
	OmpAdminDistanceIpv6            types.Int64                               `tfsdk:"omp_admin_distance_ipv6"`
	OmpAdminDistanceIpv6Variable    types.String                              `tfsdk:"omp_admin_distance_ipv6_variable"`
	EnableSdwanRemoteAccess         types.Bool                                `tfsdk:"enable_sdwan_remote_access"`
	PrimaryDnsAddressIpv4           types.String                              `tfsdk:"primary_dns_address_ipv4"`
	PrimaryDnsAddressIpv4Variable   types.String                              `tfsdk:"primary_dns_address_ipv4_variable"`
	SecondaryDnsAddressIpv4         types.String                              `tfsdk:"secondary_dns_address_ipv4"`
	SecondaryDnsAddressIpv4Variable types.String                              `tfsdk:"secondary_dns_address_ipv4_variable"`
	PrimaryDnsAddressIpv6           types.String                              `tfsdk:"primary_dns_address_ipv6"`
	PrimaryDnsAddressIpv6Variable   types.String                              `tfsdk:"primary_dns_address_ipv6_variable"`
	SecondaryDnsAddressIpv6         types.String                              `tfsdk:"secondary_dns_address_ipv6"`
	SecondaryDnsAddressIpv6Variable types.String                              `tfsdk:"secondary_dns_address_ipv6_variable"`
	HostMappings                    []ServiceLANVPNHostMappings               `tfsdk:"host_mappings"`
	AdvertiseOmpIpv4s               []ServiceLANVPNAdvertiseOmpIpv4s          `tfsdk:"advertise_omp_ipv4s"`
	AdvertiseOmpIpv6s               []ServiceLANVPNAdvertiseOmpIpv6s          `tfsdk:"advertise_omp_ipv6s"`
	Ipv4StaticRoutes                []ServiceLANVPNIpv4StaticRoutes           `tfsdk:"ipv4_static_routes"`
	Ipv6StaticRoutes                []ServiceLANVPNIpv6StaticRoutes           `tfsdk:"ipv6_static_routes"`
	Services                        []ServiceLANVPNServices                   `tfsdk:"services"`
	ServiceRoutes                   []ServiceLANVPNServiceRoutes              `tfsdk:"service_routes"`
	GreRoutes                       []ServiceLANVPNGreRoutes                  `tfsdk:"gre_routes"`
	IpsecRoutes                     []ServiceLANVPNIpsecRoutes                `tfsdk:"ipsec_routes"`
	NatPools                        []ServiceLANVPNNatPools                   `tfsdk:"nat_pools"`
	NatPortForwards                 []ServiceLANVPNNatPortForwards            `tfsdk:"nat_port_forwards"`
	StaticNats                      []ServiceLANVPNStaticNats                 `tfsdk:"static_nats"`
	StaticNatSubnets                []ServiceLANVPNStaticNatSubnets           `tfsdk:"static_nat_subnets"`
	Nat64V4Pools                    []ServiceLANVPNNat64V4Pools               `tfsdk:"nat_64_v4_pools"`
	RouteLeakFromGlobalVpns         []ServiceLANVPNRouteLeakFromGlobalVpns    `tfsdk:"route_leak_from_global_vpns"`
	RouteLeakToGlobalVpns           []ServiceLANVPNRouteLeakToGlobalVpns      `tfsdk:"route_leak_to_global_vpns"`
	RouteLeakFromOtherServices      []ServiceLANVPNRouteLeakFromOtherServices `tfsdk:"route_leak_from_other_services"`
	Ipv4ImportRouteTargets          []ServiceLANVPNIpv4ImportRouteTargets     `tfsdk:"ipv4_import_route_targets"`
	Ipv4ExportRouteTargets          []ServiceLANVPNIpv4ExportRouteTargets     `tfsdk:"ipv4_export_route_targets"`
	Ipv6ImportRouteTargets          []ServiceLANVPNIpv6ImportRouteTargets     `tfsdk:"ipv6_import_route_targets"`
	Ipv6ExportRouteTargets          []ServiceLANVPNIpv6ExportRouteTargets     `tfsdk:"ipv6_export_route_targets"`
}

type ServiceLANVPNHostMappings struct {
	HostName          types.String `tfsdk:"host_name"`
	HostNameVariable  types.String `tfsdk:"host_name_variable"`
	ListOfIps         types.Set    `tfsdk:"list_of_ips"`
	ListOfIpsVariable types.String `tfsdk:"list_of_ips_variable"`
}

type ServiceLANVPNAdvertiseOmpIpv4s struct {
	Protocol         types.String                             `tfsdk:"protocol"`
	ProtocolVariable types.String                             `tfsdk:"protocol_variable"`
	RoutePolicyId    types.String                             `tfsdk:"route_policy_id"`
	Prefixes         []ServiceLANVPNAdvertiseOmpIpv4sPrefixes `tfsdk:"prefixes"`
}

type ServiceLANVPNAdvertiseOmpIpv6s struct {
	Protocol                types.String                             `tfsdk:"protocol"`
	ProtocolVariable        types.String                             `tfsdk:"protocol_variable"`
	RoutePolicyId           types.String                             `tfsdk:"route_policy_id"`
	ProtocolSubType         types.String                             `tfsdk:"protocol_sub_type"`
	ProtocolSubTypeVariable types.String                             `tfsdk:"protocol_sub_type_variable"`
	Prefixes                []ServiceLANVPNAdvertiseOmpIpv6sPrefixes `tfsdk:"prefixes"`
}

type ServiceLANVPNIpv4StaticRoutes struct {
	NetworkAddress                 types.String                                          `tfsdk:"network_address"`
	NetworkAddressVariable         types.String                                          `tfsdk:"network_address_variable"`
	SubnetMask                     types.String                                          `tfsdk:"subnet_mask"`
	SubnetMaskVariable             types.String                                          `tfsdk:"subnet_mask_variable"`
	Gateway                        types.String                                          `tfsdk:"gateway"`
	NextHops                       []ServiceLANVPNIpv4StaticRoutesNextHops               `tfsdk:"next_hops"`
	NextHopWithTrackers            []ServiceLANVPNIpv4StaticRoutesNextHopWithTrackers    `tfsdk:"next_hop_with_trackers"`
	Null0                          types.Bool                                            `tfsdk:"null0"`
	AdministrativeDistance         types.Int64                                           `tfsdk:"administrative_distance"`
	AdministrativeDistanceVariable types.String                                          `tfsdk:"administrative_distance_variable"`
	Dhcp                           types.Bool                                            `tfsdk:"dhcp"`
	Vpn                            types.Bool                                            `tfsdk:"vpn"`
	IpStaticRouteInterface         []ServiceLANVPNIpv4StaticRoutesIpStaticRouteInterface `tfsdk:"ip_static_route_interface"`
}

type ServiceLANVPNIpv6StaticRoutes struct {
	Prefix                   types.String                                            `tfsdk:"prefix"`
	PrefixVariable           types.String                                            `tfsdk:"prefix_variable"`
	Gateway                  types.String                                            `tfsdk:"gateway"`
	NextHops                 []ServiceLANVPNIpv6StaticRoutesNextHops                 `tfsdk:"next_hops"`
	Null0                    types.Bool                                              `tfsdk:"null0"`
	Nat                      types.String                                            `tfsdk:"nat"`
	NatVariable              types.String                                            `tfsdk:"nat_variable"`
	Ipv6StaticRouteInterface []ServiceLANVPNIpv6StaticRoutesIpv6StaticRouteInterface `tfsdk:"ipv6_static_route_interface"`
}

type ServiceLANVPNServices struct {
	ServiceType           types.String `tfsdk:"service_type"`
	ServiceTypeVariable   types.String `tfsdk:"service_type_variable"`
	Ipv4Addresses         types.Set    `tfsdk:"ipv4_addresses"`
	Ipv4AddressesVariable types.String `tfsdk:"ipv4_addresses_variable"`
	Tracking              types.Bool   `tfsdk:"tracking"`
	TrackingVariable      types.String `tfsdk:"tracking_variable"`
}

type ServiceLANVPNServiceRoutes struct {
	NetworkAddress         types.String `tfsdk:"network_address"`
	NetworkAddressVariable types.String `tfsdk:"network_address_variable"`
	SubnetMask             types.String `tfsdk:"subnet_mask"`
	SubnetMaskVariable     types.String `tfsdk:"subnet_mask_variable"`
	Service                types.String `tfsdk:"service"`
	ServiceVariable        types.String `tfsdk:"service_variable"`
	Vpn                    types.Int64  `tfsdk:"vpn"`
	SseInstance            types.String `tfsdk:"sse_instance"`
	SseInstanceVariable    types.String `tfsdk:"sse_instance_variable"`
}

type ServiceLANVPNGreRoutes struct {
	NetworkAddress         types.String `tfsdk:"network_address"`
	NetworkAddressVariable types.String `tfsdk:"network_address_variable"`
	SubnetMask             types.String `tfsdk:"subnet_mask"`
	SubnetMaskVariable     types.String `tfsdk:"subnet_mask_variable"`
	Interface              types.Set    `tfsdk:"interface"`
	InterfaceVariable      types.String `tfsdk:"interface_variable"`
	Vpn                    types.Int64  `tfsdk:"vpn"`
}

type ServiceLANVPNIpsecRoutes struct {
	NetworkAddress         types.String `tfsdk:"network_address"`
	NetworkAddressVariable types.String `tfsdk:"network_address_variable"`
	SubnetMask             types.String `tfsdk:"subnet_mask"`
	SubnetMaskVariable     types.String `tfsdk:"subnet_mask_variable"`
	Interface              types.Set    `tfsdk:"interface"`
	InterfaceVariable      types.String `tfsdk:"interface_variable"`
}

type ServiceLANVPNNatPools struct {
	NatPoolName          types.Int64  `tfsdk:"nat_pool_name"`
	NatPoolNameVariable  types.String `tfsdk:"nat_pool_name_variable"`
	PrefixLength         types.Int64  `tfsdk:"prefix_length"`
	PrefixLengthVariable types.String `tfsdk:"prefix_length_variable"`
	RangeStart           types.String `tfsdk:"range_start"`
	RangeStartVariable   types.String `tfsdk:"range_start_variable"`
	RangeEnd             types.String `tfsdk:"range_end"`
	RangeEndVariable     types.String `tfsdk:"range_end_variable"`
	Overload             types.Bool   `tfsdk:"overload"`
	OverloadVariable     types.String `tfsdk:"overload_variable"`
	Direction            types.String `tfsdk:"direction"`
	DirectionVariable    types.String `tfsdk:"direction_variable"`
	TrackerObjectId      types.String `tfsdk:"tracker_object_id"`
}

type ServiceLANVPNNatPortForwards struct {
	NatPoolName                types.Int64  `tfsdk:"nat_pool_name"`
	NatPoolNameVariable        types.String `tfsdk:"nat_pool_name_variable"`
	SourcePort                 types.Int64  `tfsdk:"source_port"`
	SourcePortVariable         types.String `tfsdk:"source_port_variable"`
	TranslatePort              types.Int64  `tfsdk:"translate_port"`
	TranslatePortVariable      types.String `tfsdk:"translate_port_variable"`
	SourceIp                   types.String `tfsdk:"source_ip"`
	SourceIpVariable           types.String `tfsdk:"source_ip_variable"`
	TranslatedSourceIp         types.String `tfsdk:"translated_source_ip"`
	TranslatedSourceIpVariable types.String `tfsdk:"translated_source_ip_variable"`
	Protocol                   types.String `tfsdk:"protocol"`
	ProtocolVariable           types.String `tfsdk:"protocol_variable"`
}

type ServiceLANVPNStaticNats struct {
	NatPoolName                types.Int64  `tfsdk:"nat_pool_name"`
	NatPoolNameVariable        types.String `tfsdk:"nat_pool_name_variable"`
	SourceIp                   types.String `tfsdk:"source_ip"`
	SourceIpVariable           types.String `tfsdk:"source_ip_variable"`
	TranslatedSourceIp         types.String `tfsdk:"translated_source_ip"`
	TranslatedSourceIpVariable types.String `tfsdk:"translated_source_ip_variable"`
	StaticNatDirection         types.String `tfsdk:"static_nat_direction"`
	StaticNatDirectionVariable types.String `tfsdk:"static_nat_direction_variable"`
	TrackerObjectId            types.String `tfsdk:"tracker_object_id"`
}

type ServiceLANVPNStaticNatSubnets struct {
	SourceIpSubnet                   types.String `tfsdk:"source_ip_subnet"`
	SourceIpSubnetVariable           types.String `tfsdk:"source_ip_subnet_variable"`
	TranslatedSourceIpSubnet         types.String `tfsdk:"translated_source_ip_subnet"`
	TranslatedSourceIpSubnetVariable types.String `tfsdk:"translated_source_ip_subnet_variable"`
	PrefixLength                     types.Int64  `tfsdk:"prefix_length"`
	PrefixLengthVariable             types.String `tfsdk:"prefix_length_variable"`
	StaticNatDirection               types.String `tfsdk:"static_nat_direction"`
	StaticNatDirectionVariable       types.String `tfsdk:"static_nat_direction_variable"`
	TrackerObjectId                  types.String `tfsdk:"tracker_object_id"`
}

type ServiceLANVPNNat64V4Pools struct {
	Name               types.String `tfsdk:"name"`
	NameVariable       types.String `tfsdk:"name_variable"`
	RangeStart         types.String `tfsdk:"range_start"`
	RangeStartVariable types.String `tfsdk:"range_start_variable"`
	RangeEnd           types.String `tfsdk:"range_end"`
	RangeEndVariable   types.String `tfsdk:"range_end_variable"`
	Overload           types.Bool   `tfsdk:"overload"`
	OverloadVariable   types.String `tfsdk:"overload_variable"`
}

type ServiceLANVPNRouteLeakFromGlobalVpns struct {
	RouteProtocol         types.String                                          `tfsdk:"route_protocol"`
	RouteProtocolVariable types.String                                          `tfsdk:"route_protocol_variable"`
	RoutePolicyId         types.String                                          `tfsdk:"route_policy_id"`
	Redistributions       []ServiceLANVPNRouteLeakFromGlobalVpnsRedistributions `tfsdk:"redistributions"`
}

type ServiceLANVPNRouteLeakToGlobalVpns struct {
	RouteProtocol         types.String                                        `tfsdk:"route_protocol"`
	RouteProtocolVariable types.String                                        `tfsdk:"route_protocol_variable"`
	RoutePolicyId         types.String                                        `tfsdk:"route_policy_id"`
	Redistributions       []ServiceLANVPNRouteLeakToGlobalVpnsRedistributions `tfsdk:"redistributions"`
}

type ServiceLANVPNRouteLeakFromOtherServices struct {
	SourceVpn             types.Int64                                              `tfsdk:"source_vpn"`
	SourceVpnVariable     types.String                                             `tfsdk:"source_vpn_variable"`
	RouteProtocol         types.String                                             `tfsdk:"route_protocol"`
	RouteProtocolVariable types.String                                             `tfsdk:"route_protocol_variable"`
	RoutePolicyId         types.String                                             `tfsdk:"route_policy_id"`
	Redistributions       []ServiceLANVPNRouteLeakFromOtherServicesRedistributions `tfsdk:"redistributions"`
}

type ServiceLANVPNIpv4ImportRouteTargets struct {
	RouteTarget         types.String `tfsdk:"route_target"`
	RouteTargetVariable types.String `tfsdk:"route_target_variable"`
}

type ServiceLANVPNIpv4ExportRouteTargets struct {
	RouteTarget         types.String `tfsdk:"route_target"`
	RouteTargetVariable types.String `tfsdk:"route_target_variable"`
}

type ServiceLANVPNIpv6ImportRouteTargets struct {
	RouteTarget         types.String `tfsdk:"route_target"`
	RouteTargetVariable types.String `tfsdk:"route_target_variable"`
}

type ServiceLANVPNIpv6ExportRouteTargets struct {
	RouteTarget         types.String `tfsdk:"route_target"`
	RouteTargetVariable types.String `tfsdk:"route_target_variable"`
}

type ServiceLANVPNAdvertiseOmpIpv4sPrefixes struct {
	NetworkAddress         types.String `tfsdk:"network_address"`
	NetworkAddressVariable types.String `tfsdk:"network_address_variable"`
	SubnetMask             types.String `tfsdk:"subnet_mask"`
	SubnetMaskVariable     types.String `tfsdk:"subnet_mask_variable"`
	AggregateOnly          types.Bool   `tfsdk:"aggregate_only"`
	Region                 types.String `tfsdk:"region"`
	RegionVariable         types.String `tfsdk:"region_variable"`
}

type ServiceLANVPNAdvertiseOmpIpv6sPrefixes struct {
	Prefix         types.String `tfsdk:"prefix"`
	PrefixVariable types.String `tfsdk:"prefix_variable"`
	AggregateOnly  types.Bool   `tfsdk:"aggregate_only"`
	Region         types.String `tfsdk:"region"`
	RegionVariable types.String `tfsdk:"region_variable"`
}

type ServiceLANVPNIpv4StaticRoutesNextHops struct {
	Address                        types.String `tfsdk:"address"`
	AddressVariable                types.String `tfsdk:"address_variable"`
	AdministrativeDistance         types.Int64  `tfsdk:"administrative_distance"`
	AdministrativeDistanceVariable types.String `tfsdk:"administrative_distance_variable"`
}
type ServiceLANVPNIpv4StaticRoutesNextHopWithTrackers struct {
	Address                        types.String `tfsdk:"address"`
	AddressVariable                types.String `tfsdk:"address_variable"`
	AdministrativeDistance         types.Int64  `tfsdk:"administrative_distance"`
	AdministrativeDistanceVariable types.String `tfsdk:"administrative_distance_variable"`
	TrackerId                      types.String `tfsdk:"tracker_id"`
}
type ServiceLANVPNIpv4StaticRoutesIpStaticRouteInterface struct {
	InterfaceName         types.String                                                 `tfsdk:"interface_name"`
	InterfaceNameVariable types.String                                                 `tfsdk:"interface_name_variable"`
	NextHop               []ServiceLANVPNIpv4StaticRoutesIpStaticRouteInterfaceNextHop `tfsdk:"next_hop"`
}

type ServiceLANVPNIpv6StaticRoutesNextHops struct {
	Address                        types.String `tfsdk:"address"`
	AddressVariable                types.String `tfsdk:"address_variable"`
	AdministrativeDistance         types.Int64  `tfsdk:"administrative_distance"`
	AdministrativeDistanceVariable types.String `tfsdk:"administrative_distance_variable"`
}
type ServiceLANVPNIpv6StaticRoutesIpv6StaticRouteInterface struct {
	InterfaceName         types.String                                                   `tfsdk:"interface_name"`
	InterfaceNameVariable types.String                                                   `tfsdk:"interface_name_variable"`
	NextHop               []ServiceLANVPNIpv6StaticRoutesIpv6StaticRouteInterfaceNextHop `tfsdk:"next_hop"`
}

type ServiceLANVPNRouteLeakFromGlobalVpnsRedistributions struct {
	Protocol               types.String `tfsdk:"protocol"`
	ProtocolVariable       types.String `tfsdk:"protocol_variable"`
	RedistributionPolicyId types.String `tfsdk:"redistribution_policy_id"`
}

type ServiceLANVPNRouteLeakToGlobalVpnsRedistributions struct {
	Protocol               types.String `tfsdk:"protocol"`
	ProtocolVariable       types.String `tfsdk:"protocol_variable"`
	RedistributionPolicyId types.String `tfsdk:"redistribution_policy_id"`
}

type ServiceLANVPNRouteLeakFromOtherServicesRedistributions struct {
	Protocol               types.String `tfsdk:"protocol"`
	ProtocolVariable       types.String `tfsdk:"protocol_variable"`
	RedistributionPolicyId types.String `tfsdk:"redistribution_policy_id"`
}

type ServiceLANVPNIpv4StaticRoutesIpStaticRouteInterfaceNextHop struct {
	Address                        types.String `tfsdk:"address"`
	AddressVariable                types.String `tfsdk:"address_variable"`
	AdministrativeDistance         types.Int64  `tfsdk:"administrative_distance"`
	AdministrativeDistanceVariable types.String `tfsdk:"administrative_distance_variable"`
}

type ServiceLANVPNIpv6StaticRoutesIpv6StaticRouteInterfaceNextHop struct {
	Address                        types.String `tfsdk:"address"`
	AddressVariable                types.String `tfsdk:"address_variable"`
	AdministrativeDistance         types.Int64  `tfsdk:"administrative_distance"`
	AdministrativeDistanceVariable types.String `tfsdk:"administrative_distance_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceLANVPN) getModel() string {
	return "service_lan_vpn"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceLANVPN) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/lan/vpn", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data ServiceLANVPN) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.VpnVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"vpnId.optionType", "variable")
			body, _ = sjson.Set(body, path+"vpnId.value", data.VpnVariable.ValueString())
		}
	} else if data.Vpn.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"vpnId.optionType", "default")
			body, _ = sjson.Set(body, path+"vpnId.value", 0)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"vpnId.optionType", "global")
			body, _ = sjson.Set(body, path+"vpnId.value", data.Vpn.ValueInt64())
		}
	}

	if !data.ConfigDescriptionVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"name.optionType", "variable")
			body, _ = sjson.Set(body, path+"name.value", data.ConfigDescriptionVariable.ValueString())
		}
	} else if data.ConfigDescription.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"name.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"name.optionType", "global")
			body, _ = sjson.Set(body, path+"name.value", data.ConfigDescription.ValueString())
		}
	}

	if !data.OmpAdminDistanceIpv4Variable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ompAdminDistance.optionType", "variable")
			body, _ = sjson.Set(body, path+"ompAdminDistance.value", data.OmpAdminDistanceIpv4Variable.ValueString())
		}
	} else if data.OmpAdminDistanceIpv4.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ompAdminDistance.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ompAdminDistance.optionType", "global")
			body, _ = sjson.Set(body, path+"ompAdminDistance.value", data.OmpAdminDistanceIpv4.ValueInt64())
		}
	}

	if !data.OmpAdminDistanceIpv6Variable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv6.optionType", "variable")
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv6.value", data.OmpAdminDistanceIpv6Variable.ValueString())
		}
	} else if data.OmpAdminDistanceIpv6.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv6.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv6.optionType", "global")
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv6.value", data.OmpAdminDistanceIpv6.ValueInt64())
		}
	}
	if data.EnableSdwanRemoteAccess.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"enableSdra.optionType", "default")
			body, _ = sjson.Set(body, path+"enableSdra.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"enableSdra.optionType", "global")
			body, _ = sjson.Set(body, path+"enableSdra.value", data.EnableSdwanRemoteAccess.ValueBool())
		}
	}

	if !data.PrimaryDnsAddressIpv4Variable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dnsIpv4.primaryDnsAddressIpv4.optionType", "variable")
			body, _ = sjson.Set(body, path+"dnsIpv4.primaryDnsAddressIpv4.value", data.PrimaryDnsAddressIpv4Variable.ValueString())
		}
	} else if data.PrimaryDnsAddressIpv4.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dnsIpv4.primaryDnsAddressIpv4.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"dnsIpv4.primaryDnsAddressIpv4.optionType", "global")
			body, _ = sjson.Set(body, path+"dnsIpv4.primaryDnsAddressIpv4.value", data.PrimaryDnsAddressIpv4.ValueString())
		}
	}

	if !data.SecondaryDnsAddressIpv4Variable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dnsIpv4.secondaryDnsAddressIpv4.optionType", "variable")
			body, _ = sjson.Set(body, path+"dnsIpv4.secondaryDnsAddressIpv4.value", data.SecondaryDnsAddressIpv4Variable.ValueString())
		}
	} else if !data.SecondaryDnsAddressIpv4.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dnsIpv4.secondaryDnsAddressIpv4.optionType", "global")
			body, _ = sjson.Set(body, path+"dnsIpv4.secondaryDnsAddressIpv4.value", data.SecondaryDnsAddressIpv4.ValueString())
		}
	}

	if !data.PrimaryDnsAddressIpv6Variable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dnsIpv6.primaryDnsAddressIpv6.optionType", "variable")
			body, _ = sjson.Set(body, path+"dnsIpv6.primaryDnsAddressIpv6.value", data.PrimaryDnsAddressIpv6Variable.ValueString())
		}
	} else if data.PrimaryDnsAddressIpv6.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dnsIpv6.primaryDnsAddressIpv6.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"dnsIpv6.primaryDnsAddressIpv6.optionType", "global")
			body, _ = sjson.Set(body, path+"dnsIpv6.primaryDnsAddressIpv6.value", data.PrimaryDnsAddressIpv6.ValueString())
		}
	}

	if !data.SecondaryDnsAddressIpv6Variable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dnsIpv6.secondaryDnsAddressIpv6.optionType", "variable")
			body, _ = sjson.Set(body, path+"dnsIpv6.secondaryDnsAddressIpv6.value", data.SecondaryDnsAddressIpv6Variable.ValueString())
		}
	} else if !data.SecondaryDnsAddressIpv6.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dnsIpv6.secondaryDnsAddressIpv6.optionType", "global")
			body, _ = sjson.Set(body, path+"dnsIpv6.secondaryDnsAddressIpv6.value", data.SecondaryDnsAddressIpv6.ValueString())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"newHostMapping", []interface{}{})
		for _, item := range data.HostMappings {
			itemBody := ""

			if !item.HostNameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "hostName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "hostName.value", item.HostNameVariable.ValueString())
				}
			} else if !item.HostName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "hostName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "hostName.value", item.HostName.ValueString())
				}
			}

			if !item.ListOfIpsVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "listOfIp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "listOfIp.value", item.ListOfIpsVariable.ValueString())
				}
			} else if !item.ListOfIps.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "listOfIp.optionType", "global")
					var values []string
					item.ListOfIps.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "listOfIp.value", values)
				}
			}
			body, _ = sjson.SetRaw(body, path+"newHostMapping.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"ompAdvertiseIp4", []interface{}{})
		for _, item := range data.AdvertiseOmpIpv4s {
			itemBody := ""

			if !item.ProtocolVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ompProtocol.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ompProtocol.value", item.ProtocolVariable.ValueString())
				}
			} else if !item.Protocol.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ompProtocol.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ompProtocol.value", item.Protocol.ValueString())
				}
			}
			if !item.RoutePolicyId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.value", item.RoutePolicyId.ValueString())
				}
			}
			if true {

				for _, childItem := range item.Prefixes {
					itemChildBody := ""

					if !childItem.NetworkAddressVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.address.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.address.value", childItem.NetworkAddressVariable.ValueString())
						}
					} else if !childItem.NetworkAddress.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.address.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.address.value", childItem.NetworkAddress.ValueString())
						}
					}

					if !childItem.SubnetMaskVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.mask.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.mask.value", childItem.SubnetMaskVariable.ValueString())
						}
					} else if !childItem.SubnetMask.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.mask.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.mask.value", childItem.SubnetMask.ValueString())
						}
					}
					if childItem.AggregateOnly.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "aggregateOnly.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "aggregateOnly.value", false)
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "aggregateOnly.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "aggregateOnly.value", childItem.AggregateOnly.ValueBool())
						}
					}

					if !childItem.RegionVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "region.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "region.value", childItem.RegionVariable.ValueString())
						}
					} else if childItem.Region.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "region.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "region.value", "core-and-access")
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "region.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "region.value", childItem.Region.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "prefixList.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"ompAdvertiseIp4.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"ompAdvertiseIpv6", []interface{}{})
		for _, item := range data.AdvertiseOmpIpv6s {
			itemBody := ""

			if !item.ProtocolVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ompProtocol.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ompProtocol.value", item.ProtocolVariable.ValueString())
				}
			} else if !item.Protocol.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ompProtocol.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ompProtocol.value", item.Protocol.ValueString())
				}
			}
			if !item.RoutePolicyId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.value", item.RoutePolicyId.ValueString())
				}
			}

			if !item.ProtocolSubTypeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "protocolSubType.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "protocolSubType.value", item.ProtocolSubTypeVariable.ValueString())
				}
			} else if item.ProtocolSubType.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "protocolSubType.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "protocolSubType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "protocolSubType.value", item.ProtocolSubType.ValueString())
				}
			}
			if true {

				for _, childItem := range item.Prefixes {
					itemChildBody := ""

					if !childItem.PrefixVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.value", childItem.PrefixVariable.ValueString())
						}
					} else if !childItem.Prefix.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "prefix.value", childItem.Prefix.ValueString())
						}
					}
					if childItem.AggregateOnly.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "aggregateOnly.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "aggregateOnly.value", false)
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "aggregateOnly.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "aggregateOnly.value", childItem.AggregateOnly.ValueBool())
						}
					}

					if !childItem.RegionVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "region.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "region.value", childItem.RegionVariable.ValueString())
						}
					} else if childItem.Region.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "region.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "region.value", "core-and-access")
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "region.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "region.value", childItem.Region.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "prefixList.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"ompAdvertiseIpv6.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"ipv4Route", []interface{}{})
		for _, item := range data.Ipv4StaticRoutes {
			itemBody := ""

			if !item.NetworkAddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.value", item.NetworkAddressVariable.ValueString())
				}
			} else if !item.NetworkAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.value", item.NetworkAddress.ValueString())
				}
			}

			if !item.SubnetMaskVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.value", item.SubnetMaskVariable.ValueString())
				}
			} else if !item.SubnetMask.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.value", item.SubnetMask.ValueString())
				}
			}
			if true && item.Gateway.ValueString() == "nextHop" {

				for _, childItem := range item.NextHops {
					itemChildBody := ""

					if !childItem.AddressVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.AddressVariable.ValueString())
						}
					} else if !childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
						}
					}

					if !childItem.AdministrativeDistanceVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.value", childItem.AdministrativeDistanceVariable.ValueString())
						}
					} else if childItem.AdministrativeDistance.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.value", 1)
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.value", childItem.AdministrativeDistance.ValueInt64())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "oneOfIpRoute.nextHopContainer.nextHop.-1", itemChildBody)
				}
			}
			if true && item.Gateway.ValueString() == "nextHop" {

				for _, childItem := range item.NextHopWithTrackers {
					itemChildBody := ""

					if !childItem.AddressVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.AddressVariable.ValueString())
						}
					} else if !childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
						}
					}

					if !childItem.AdministrativeDistanceVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.value", childItem.AdministrativeDistanceVariable.ValueString())
						}
					} else if childItem.AdministrativeDistance.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.value", 1)
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.value", childItem.AdministrativeDistance.ValueInt64())
						}
					}
					if !childItem.TrackerId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "tracker.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "tracker.refId.value", childItem.TrackerId.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "oneOfIpRoute.nextHopContainer.nextHopWithTracker.-1", itemChildBody)
				}
			}
			if !item.Null0.IsNull() {
				if true && item.Gateway.ValueString() == "null0" {
					itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.null0.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.null0.value", item.Null0.ValueBool())
				}
			}

			if !item.AdministrativeDistanceVariable.IsNull() {
				if true && item.Gateway.ValueString() == "null0" {
					itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.distance.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.distance.value", item.AdministrativeDistanceVariable.ValueString())
				}
			} else if !item.AdministrativeDistance.IsNull() {
				if true && item.Gateway.ValueString() == "null0" {
					itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.distance.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.distance.value", item.AdministrativeDistance.ValueInt64())
				}
			}
			if !item.Dhcp.IsNull() {
				if true && item.Gateway.ValueString() == "dhcp" {
					itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.dhcp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.dhcp.value", item.Dhcp.ValueBool())
				}
			}
			if !item.Vpn.IsNull() {
				if true && item.Gateway.ValueString() == "vpn" {
					itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.vpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.vpn.value", item.Vpn.ValueBool())
				}
			}
			if true && item.Gateway.ValueString() == "staticRouteInterface" {

				for _, childItem := range item.IpStaticRouteInterface {
					itemChildBody := ""

					if !childItem.InterfaceNameVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "interfaceName.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "interfaceName.value", childItem.InterfaceNameVariable.ValueString())
						}
					} else if !childItem.InterfaceName.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "interfaceName.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "interfaceName.value", childItem.InterfaceName.ValueString())
						}
					}
					if true {

						for _, childChildItem := range childItem.NextHop {
							itemChildChildBody := ""

							if !childChildItem.AddressVariable.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "address.optionType", "variable")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "address.value", childChildItem.AddressVariable.ValueString())
								}
							} else if !childChildItem.Address.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "address.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "address.value", childChildItem.Address.ValueString())
								}
							}

							if !childChildItem.AdministrativeDistanceVariable.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "distance.optionType", "variable")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "distance.value", childChildItem.AdministrativeDistanceVariable.ValueString())
								}
							} else if !childChildItem.AdministrativeDistance.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "distance.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "distance.value", childChildItem.AdministrativeDistance.ValueInt64())
								}
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "nextHop.-1", itemChildChildBody)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "oneOfIpRoute.interfaceContainer.ipStaticRouteInterface.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"ipv4Route.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"ipv6Route", []interface{}{})
		for _, item := range data.Ipv6StaticRoutes {
			itemBody := ""

			if !item.PrefixVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.value", item.PrefixVariable.ValueString())
				}
			} else if !item.Prefix.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.value", item.Prefix.ValueString())
				}
			}
			if true && item.Gateway.ValueString() == "nextHop" {

				for _, childItem := range item.NextHops {
					itemChildBody := ""

					if !childItem.AddressVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.AddressVariable.ValueString())
						}
					} else if !childItem.Address.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
						}
					}

					if !childItem.AdministrativeDistanceVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.value", childItem.AdministrativeDistanceVariable.ValueString())
						}
					} else if childItem.AdministrativeDistance.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.value", 1)
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "distance.value", childItem.AdministrativeDistance.ValueInt64())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "oneOfIpRoute.nextHopContainer.nextHop.-1", itemChildBody)
				}
			}
			if !item.Null0.IsNull() {
				if true && item.Gateway.ValueString() == "null0" {
					itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.null0.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.null0.value", item.Null0.ValueBool())
				}
			}

			if !item.NatVariable.IsNull() {
				if true && item.Gateway.ValueString() == "nat" {
					itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.nat.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.nat.value", item.NatVariable.ValueString())
				}
			} else if !item.Nat.IsNull() {
				if true && item.Gateway.ValueString() == "nat" {
					itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.nat.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.nat.value", item.Nat.ValueString())
				}
			}
			if true && item.Gateway.ValueString() == "staticRouteInterface" {

				for _, childItem := range item.Ipv6StaticRouteInterface {
					itemChildBody := ""

					if !childItem.InterfaceNameVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "interfaceName.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "interfaceName.value", childItem.InterfaceNameVariable.ValueString())
						}
					} else if !childItem.InterfaceName.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "interfaceName.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "interfaceName.value", childItem.InterfaceName.ValueString())
						}
					}
					if true {

						for _, childChildItem := range childItem.NextHop {
							itemChildChildBody := ""

							if !childChildItem.AddressVariable.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "address.optionType", "variable")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "address.value", childChildItem.AddressVariable.ValueString())
								}
							} else if !childChildItem.Address.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "address.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "address.value", childChildItem.Address.ValueString())
								}
							}

							if !childChildItem.AdministrativeDistanceVariable.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "distance.optionType", "variable")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "distance.value", childChildItem.AdministrativeDistanceVariable.ValueString())
								}
							} else if !childChildItem.AdministrativeDistance.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "distance.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "distance.value", childChildItem.AdministrativeDistance.ValueInt64())
								}
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "nextHop.-1", itemChildChildBody)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "oneOfIpRoute.interfaceContainer.ipv6StaticRouteInterface.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"ipv6Route.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"service", []interface{}{})
		for _, item := range data.Services {
			itemBody := ""

			if !item.ServiceTypeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "serviceType.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "serviceType.value", item.ServiceTypeVariable.ValueString())
				}
			} else if !item.ServiceType.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "serviceType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "serviceType.value", item.ServiceType.ValueString())
				}
			}

			if !item.Ipv4AddressesVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipv4Addresses.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ipv4Addresses.value", item.Ipv4AddressesVariable.ValueString())
				}
			} else if !item.Ipv4Addresses.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ipv4Addresses.optionType", "global")
					var values []string
					item.Ipv4Addresses.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "ipv4Addresses.value", values)
				}
			}

			if !item.TrackingVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tracking.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "tracking.value", item.TrackingVariable.ValueString())
				}
			} else if item.Tracking.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tracking.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "tracking.value", true)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "tracking.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "tracking.value", item.Tracking.ValueBool())
				}
			}
			body, _ = sjson.SetRaw(body, path+"service.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"serviceRoute", []interface{}{})
		for _, item := range data.ServiceRoutes {
			itemBody := ""

			if !item.NetworkAddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.value", item.NetworkAddressVariable.ValueString())
				}
			} else if !item.NetworkAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.value", item.NetworkAddress.ValueString())
				}
			}

			if !item.SubnetMaskVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.value", item.SubnetMaskVariable.ValueString())
				}
			} else if !item.SubnetMask.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.value", item.SubnetMask.ValueString())
				}
			}

			if !item.ServiceVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "service.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "service.value", item.ServiceVariable.ValueString())
				}
			} else if item.Service.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "service.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "service.value", "SIG")
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "service.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "service.value", item.Service.ValueString())
				}
			}
			if !item.Vpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
				}
			}

			if !item.SseInstanceVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sseInstance.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sseInstance.value", item.SseInstanceVariable.ValueString())
				}
			} else if !item.SseInstance.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sseInstance.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sseInstance.value", item.SseInstance.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"serviceRoute.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"greRoute", []interface{}{})
		for _, item := range data.GreRoutes {
			itemBody := ""

			if !item.NetworkAddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.value", item.NetworkAddressVariable.ValueString())
				}
			} else if !item.NetworkAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.value", item.NetworkAddress.ValueString())
				}
			}

			if !item.SubnetMaskVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.value", item.SubnetMaskVariable.ValueString())
				}
			} else if !item.SubnetMask.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.value", item.SubnetMask.ValueString())
				}
			}

			if !item.InterfaceVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interface.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "interface.value", item.InterfaceVariable.ValueString())
				}
			} else if item.Interface.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interface.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interface.optionType", "global")
					var values []string
					item.Interface.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "interface.value", values)
				}
			}
			if !item.Vpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"greRoute.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"ipsecRoute", []interface{}{})
		for _, item := range data.IpsecRoutes {
			itemBody := ""

			if !item.NetworkAddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.value", item.NetworkAddressVariable.ValueString())
				}
			} else if !item.NetworkAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.value", item.NetworkAddress.ValueString())
				}
			}

			if !item.SubnetMaskVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.value", item.SubnetMaskVariable.ValueString())
				}
			} else if !item.SubnetMask.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.value", item.SubnetMask.ValueString())
				}
			}

			if !item.InterfaceVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interface.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "interface.value", item.InterfaceVariable.ValueString())
				}
			} else if item.Interface.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interface.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "interface.optionType", "global")
					var values []string
					item.Interface.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "interface.value", values)
				}
			}
			body, _ = sjson.SetRaw(body, path+"ipsecRoute.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"natPool", []interface{}{})
		for _, item := range data.NatPools {
			itemBody := ""

			if !item.NatPoolNameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "natPoolName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "natPoolName.value", item.NatPoolNameVariable.ValueString())
				}
			} else if !item.NatPoolName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "natPoolName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "natPoolName.value", item.NatPoolName.ValueInt64())
				}
			}

			if !item.PrefixLengthVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefixLength.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefixLength.value", item.PrefixLengthVariable.ValueString())
				}
			} else if !item.PrefixLength.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefixLength.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefixLength.value", item.PrefixLength.ValueInt64())
				}
			}

			if !item.RangeStartVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "rangeStart.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "rangeStart.value", item.RangeStartVariable.ValueString())
				}
			} else if !item.RangeStart.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "rangeStart.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "rangeStart.value", item.RangeStart.ValueString())
				}
			}

			if !item.RangeEndVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "rangeEnd.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "rangeEnd.value", item.RangeEndVariable.ValueString())
				}
			} else if !item.RangeEnd.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "rangeEnd.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "rangeEnd.value", item.RangeEnd.ValueString())
				}
			}

			if !item.OverloadVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "overload.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "overload.value", item.OverloadVariable.ValueString())
				}
			} else if item.Overload.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "overload.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "overload.value", true)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "overload.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "overload.value", item.Overload.ValueBool())
				}
			}

			if !item.DirectionVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "direction.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "direction.value", item.DirectionVariable.ValueString())
				}
			} else if !item.Direction.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "direction.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "direction.value", item.Direction.ValueString())
				}
			}
			if !item.TrackerObjectId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackingObject.trackerId.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "trackingObject.trackerId.refId.value", item.TrackerObjectId.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"natPool.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"natPortForward", []interface{}{})
		for _, item := range data.NatPortForwards {
			itemBody := ""

			if !item.NatPoolNameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "natPoolName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "natPoolName.value", item.NatPoolNameVariable.ValueString())
				}
			} else if item.NatPoolName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "natPoolName.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "natPoolName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "natPoolName.value", item.NatPoolName.ValueInt64())
				}
			}

			if !item.SourcePortVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourcePort.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourcePort.value", item.SourcePortVariable.ValueString())
				}
			} else if !item.SourcePort.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourcePort.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sourcePort.value", item.SourcePort.ValueInt64())
				}
			}

			if !item.TranslatePortVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "translatePort.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "translatePort.value", item.TranslatePortVariable.ValueString())
				}
			} else if !item.TranslatePort.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "translatePort.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "translatePort.value", item.TranslatePort.ValueInt64())
				}
			}

			if !item.SourceIpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceIp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourceIp.value", item.SourceIpVariable.ValueString())
				}
			} else if !item.SourceIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceIp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sourceIp.value", item.SourceIp.ValueString())
				}
			}

			if !item.TranslatedSourceIpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "translatedSourceIp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "translatedSourceIp.value", item.TranslatedSourceIpVariable.ValueString())
				}
			} else if !item.TranslatedSourceIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "translatedSourceIp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "translatedSourceIp.value", item.TranslatedSourceIp.ValueString())
				}
			}

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
			body, _ = sjson.SetRaw(body, path+"natPortForward.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"staticNat", []interface{}{})
		for _, item := range data.StaticNats {
			itemBody := ""

			if !item.NatPoolNameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "natPoolName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "natPoolName.value", item.NatPoolNameVariable.ValueString())
				}
			} else if item.NatPoolName.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "natPoolName.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "natPoolName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "natPoolName.value", item.NatPoolName.ValueInt64())
				}
			}

			if !item.SourceIpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceIp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourceIp.value", item.SourceIpVariable.ValueString())
				}
			} else if !item.SourceIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceIp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sourceIp.value", item.SourceIp.ValueString())
				}
			}

			if !item.TranslatedSourceIpVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "translatedSourceIp.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "translatedSourceIp.value", item.TranslatedSourceIpVariable.ValueString())
				}
			} else if !item.TranslatedSourceIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "translatedSourceIp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "translatedSourceIp.value", item.TranslatedSourceIp.ValueString())
				}
			}

			if !item.StaticNatDirectionVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "staticNatDirection.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "staticNatDirection.value", item.StaticNatDirectionVariable.ValueString())
				}
			} else if !item.StaticNatDirection.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "staticNatDirection.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "staticNatDirection.value", item.StaticNatDirection.ValueString())
				}
			}
			if !item.TrackerObjectId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackingObject.trackerId.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "trackingObject.trackerId.refId.value", item.TrackerObjectId.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"staticNat.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"staticNatSubnet", []interface{}{})
		for _, item := range data.StaticNatSubnets {
			itemBody := ""

			if !item.SourceIpSubnetVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceIpSubnet.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourceIpSubnet.value", item.SourceIpSubnetVariable.ValueString())
				}
			} else if !item.SourceIpSubnet.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceIpSubnet.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sourceIpSubnet.value", item.SourceIpSubnet.ValueString())
				}
			}

			if !item.TranslatedSourceIpSubnetVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "translatedSourceIpSubnet.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "translatedSourceIpSubnet.value", item.TranslatedSourceIpSubnetVariable.ValueString())
				}
			} else if !item.TranslatedSourceIpSubnet.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "translatedSourceIpSubnet.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "translatedSourceIpSubnet.value", item.TranslatedSourceIpSubnet.ValueString())
				}
			}

			if !item.PrefixLengthVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefixLength.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefixLength.value", item.PrefixLengthVariable.ValueString())
				}
			} else if !item.PrefixLength.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefixLength.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefixLength.value", item.PrefixLength.ValueInt64())
				}
			}

			if !item.StaticNatDirectionVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "staticNatDirection.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "staticNatDirection.value", item.StaticNatDirectionVariable.ValueString())
				}
			} else if !item.StaticNatDirection.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "staticNatDirection.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "staticNatDirection.value", item.StaticNatDirection.ValueString())
				}
			}
			if !item.TrackerObjectId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "trackingObject.trackerId.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "trackingObject.trackerId.refId.value", item.TrackerObjectId.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"staticNatSubnet.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"nat64V4Pool", []interface{}{})
		for _, item := range data.Nat64V4Pools {
			itemBody := ""

			if !item.NameVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolName.value", item.NameVariable.ValueString())
				}
			} else if !item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolName.value", item.Name.ValueString())
				}
			}

			if !item.RangeStartVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeStart.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeStart.value", item.RangeStartVariable.ValueString())
				}
			} else if !item.RangeStart.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeStart.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeStart.value", item.RangeStart.ValueString())
				}
			}

			if !item.RangeEndVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeEnd.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeEnd.value", item.RangeEndVariable.ValueString())
				}
			} else if !item.RangeEnd.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeEnd.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeEnd.value", item.RangeEnd.ValueString())
				}
			}

			if !item.OverloadVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolOverload.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolOverload.value", item.OverloadVariable.ValueString())
				}
			} else if item.Overload.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolOverload.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolOverload.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolOverload.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "nat64V4PoolOverload.value", item.Overload.ValueBool())
				}
			}
			body, _ = sjson.SetRaw(body, path+"nat64V4Pool.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"routeLeakFromGlobal", []interface{}{})
		for _, item := range data.RouteLeakFromGlobalVpns {
			itemBody := ""

			if !item.RouteProtocolVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "routeProtocol.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "routeProtocol.value", item.RouteProtocolVariable.ValueString())
				}
			} else if !item.RouteProtocol.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "routeProtocol.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "routeProtocol.value", item.RouteProtocol.ValueString())
				}
			}
			if !item.RoutePolicyId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.value", item.RoutePolicyId.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "redistributeToProtocol", []interface{}{})
				for _, childItem := range item.Redistributions {
					itemChildBody := ""

					if !childItem.ProtocolVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.value", childItem.ProtocolVariable.ValueString())
						}
					} else if !childItem.Protocol.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.value", childItem.Protocol.ValueString())
						}
					}
					if !childItem.RedistributionPolicyId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "policy.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "policy.refId.value", childItem.RedistributionPolicyId.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "redistributeToProtocol.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"routeLeakFromGlobal.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"routeLeakFromService", []interface{}{})
		for _, item := range data.RouteLeakToGlobalVpns {
			itemBody := ""

			if !item.RouteProtocolVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "routeProtocol.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "routeProtocol.value", item.RouteProtocolVariable.ValueString())
				}
			} else if !item.RouteProtocol.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "routeProtocol.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "routeProtocol.value", item.RouteProtocol.ValueString())
				}
			}
			if !item.RoutePolicyId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.value", item.RoutePolicyId.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "redistributeToProtocol", []interface{}{})
				for _, childItem := range item.Redistributions {
					itemChildBody := ""

					if !childItem.ProtocolVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.value", childItem.ProtocolVariable.ValueString())
						}
					} else if !childItem.Protocol.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.value", childItem.Protocol.ValueString())
						}
					}
					if !childItem.RedistributionPolicyId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "policy.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "policy.refId.value", childItem.RedistributionPolicyId.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "redistributeToProtocol.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"routeLeakFromService.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"routeLeakBetweenServices", []interface{}{})
		for _, item := range data.RouteLeakFromOtherServices {
			itemBody := ""

			if !item.SourceVpnVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceVpn.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sourceVpn.value", item.SourceVpnVariable.ValueString())
				}
			} else if !item.SourceVpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sourceVpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sourceVpn.value", item.SourceVpn.ValueInt64())
				}
			}

			if !item.RouteProtocolVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "routeProtocol.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "routeProtocol.value", item.RouteProtocolVariable.ValueString())
				}
			} else if !item.RouteProtocol.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "routeProtocol.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "routeProtocol.value", item.RouteProtocol.ValueString())
				}
			}
			if !item.RoutePolicyId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.value", item.RoutePolicyId.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "redistributeToProtocol", []interface{}{})
				for _, childItem := range item.Redistributions {
					itemChildBody := ""

					if !childItem.ProtocolVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.value", childItem.ProtocolVariable.ValueString())
						}
					} else if !childItem.Protocol.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "protocol.value", childItem.Protocol.ValueString())
						}
					}
					if !childItem.RedistributionPolicyId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "policy.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "policy.refId.value", childItem.RedistributionPolicyId.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "redistributeToProtocol.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"routeLeakBetweenServices.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"mplsVpnIpv4RouteTarget.importRtList", []interface{}{})
		for _, item := range data.Ipv4ImportRouteTargets {
			itemBody := ""

			if !item.RouteTargetVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "rt.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "rt.value", item.RouteTargetVariable.ValueString())
				}
			} else if !item.RouteTarget.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "rt.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "rt.value", item.RouteTarget.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"mplsVpnIpv4RouteTarget.importRtList.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"mplsVpnIpv4RouteTarget.exportRtList", []interface{}{})
		for _, item := range data.Ipv4ExportRouteTargets {
			itemBody := ""

			if !item.RouteTargetVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "rt.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "rt.value", item.RouteTargetVariable.ValueString())
				}
			} else if !item.RouteTarget.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "rt.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "rt.value", item.RouteTarget.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"mplsVpnIpv4RouteTarget.exportRtList.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"mplsVpnIpv6RouteTarget.importRtList", []interface{}{})
		for _, item := range data.Ipv6ImportRouteTargets {
			itemBody := ""

			if !item.RouteTargetVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "rt.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "rt.value", item.RouteTargetVariable.ValueString())
				}
			} else if !item.RouteTarget.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "rt.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "rt.value", item.RouteTarget.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"mplsVpnIpv6RouteTarget.importRtList.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"mplsVpnIpv6RouteTarget.exportRtList", []interface{}{})
		for _, item := range data.Ipv6ExportRouteTargets {
			itemBody := ""

			if !item.RouteTargetVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "rt.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "rt.value", item.RouteTargetVariable.ValueString())
				}
			} else if !item.RouteTarget.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "rt.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "rt.value", item.RouteTarget.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"mplsVpnIpv6RouteTarget.exportRtList.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceLANVPN) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Vpn = types.Int64Null()
	data.VpnVariable = types.StringNull()
	if t := res.Get(path + "vpnId.optionType"); t.Exists() {
		va := res.Get(path + "vpnId.value")
		if t.String() == "variable" {
			data.VpnVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Vpn = types.Int64Value(va.Int())
		}
	}
	data.ConfigDescription = types.StringNull()
	data.ConfigDescriptionVariable = types.StringNull()
	if t := res.Get(path + "name.optionType"); t.Exists() {
		va := res.Get(path + "name.value")
		if t.String() == "variable" {
			data.ConfigDescriptionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ConfigDescription = types.StringValue(va.String())
		}
	}
	data.OmpAdminDistanceIpv4 = types.Int64Null()
	data.OmpAdminDistanceIpv4Variable = types.StringNull()
	if t := res.Get(path + "ompAdminDistance.optionType"); t.Exists() {
		va := res.Get(path + "ompAdminDistance.value")
		if t.String() == "variable" {
			data.OmpAdminDistanceIpv4Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.OmpAdminDistanceIpv4 = types.Int64Value(va.Int())
		}
	}
	data.OmpAdminDistanceIpv6 = types.Int64Null()
	data.OmpAdminDistanceIpv6Variable = types.StringNull()
	if t := res.Get(path + "ompAdminDistanceIpv6.optionType"); t.Exists() {
		va := res.Get(path + "ompAdminDistanceIpv6.value")
		if t.String() == "variable" {
			data.OmpAdminDistanceIpv6Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.OmpAdminDistanceIpv6 = types.Int64Value(va.Int())
		}
	}
	data.EnableSdwanRemoteAccess = types.BoolNull()

	if t := res.Get(path + "enableSdra.optionType"); t.Exists() {
		va := res.Get(path + "enableSdra.value")
		if t.String() == "global" {
			data.EnableSdwanRemoteAccess = types.BoolValue(va.Bool())
		}
	}
	data.PrimaryDnsAddressIpv4 = types.StringNull()
	data.PrimaryDnsAddressIpv4Variable = types.StringNull()
	if t := res.Get(path + "dnsIpv4.primaryDnsAddressIpv4.optionType"); t.Exists() {
		va := res.Get(path + "dnsIpv4.primaryDnsAddressIpv4.value")
		if t.String() == "variable" {
			data.PrimaryDnsAddressIpv4Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PrimaryDnsAddressIpv4 = types.StringValue(va.String())
		}
	}
	data.SecondaryDnsAddressIpv4 = types.StringNull()
	data.SecondaryDnsAddressIpv4Variable = types.StringNull()
	if t := res.Get(path + "dnsIpv4.secondaryDnsAddressIpv4.optionType"); t.Exists() {
		va := res.Get(path + "dnsIpv4.secondaryDnsAddressIpv4.value")
		if t.String() == "variable" {
			data.SecondaryDnsAddressIpv4Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SecondaryDnsAddressIpv4 = types.StringValue(va.String())
		}
	}
	data.PrimaryDnsAddressIpv6 = types.StringNull()
	data.PrimaryDnsAddressIpv6Variable = types.StringNull()
	if t := res.Get(path + "dnsIpv6.primaryDnsAddressIpv6.optionType"); t.Exists() {
		va := res.Get(path + "dnsIpv6.primaryDnsAddressIpv6.value")
		if t.String() == "variable" {
			data.PrimaryDnsAddressIpv6Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PrimaryDnsAddressIpv6 = types.StringValue(va.String())
		}
	}
	data.SecondaryDnsAddressIpv6 = types.StringNull()
	data.SecondaryDnsAddressIpv6Variable = types.StringNull()
	if t := res.Get(path + "dnsIpv6.secondaryDnsAddressIpv6.optionType"); t.Exists() {
		va := res.Get(path + "dnsIpv6.secondaryDnsAddressIpv6.value")
		if t.String() == "variable" {
			data.SecondaryDnsAddressIpv6Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SecondaryDnsAddressIpv6 = types.StringValue(va.String())
		}
	}
	oldHostMappings := data.HostMappings
	if value := res.Get(path + "newHostMapping"); value.Exists() && len(value.Array()) > 0 {
		data.HostMappings = make([]ServiceLANVPNHostMappings, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNHostMappings{}
			item.HostName = types.StringNull()
			item.HostNameVariable = types.StringNull()
			if t := v.Get("hostName.optionType"); t.Exists() {
				va := v.Get("hostName.value")
				if t.String() == "variable" {
					item.HostNameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.HostName = types.StringValue(va.String())
				}
			}
			item.ListOfIps = types.SetNull(types.StringType)
			item.ListOfIpsVariable = types.StringNull()
			if t := v.Get("listOfIp.optionType"); t.Exists() {
				va := v.Get("listOfIp.value")
				if t.String() == "variable" {
					item.ListOfIpsVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.ListOfIps = helpers.GetStringSet(va.Array())
				}
			}
			data.HostMappings = append(data.HostMappings, item)
			return true
		})
	} else {
		data.HostMappings = nil
	}
	if !fullRead {
		resultHostMappings := make([]ServiceLANVPNHostMappings, 0, len(data.HostMappings))
		matchedHostMappings := make([]bool, len(data.HostMappings))
		for _, oldItem := range oldHostMappings {
			for ni := range data.HostMappings {
				if matchedHostMappings[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.HostNameVariable.ValueString() != "" || data.HostMappings[ni].HostNameVariable.ValueString() != "") {
					if oldItem.HostNameVariable.ValueString() != data.HostMappings[ni].HostNameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.HostName.ValueString() != data.HostMappings[ni].HostName.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedHostMappings[ni] = true
					resultHostMappings = append(resultHostMappings, data.HostMappings[ni])
					break
				}
			}
		}
		for ni := range data.HostMappings {
			if !matchedHostMappings[ni] {
				resultHostMappings = append(resultHostMappings, data.HostMappings[ni])
			}
		}
		data.HostMappings = resultHostMappings
	}
	oldAdvertiseOmpIpv4s := data.AdvertiseOmpIpv4s
	if value := res.Get(path + "ompAdvertiseIp4"); value.Exists() && len(value.Array()) > 0 {
		data.AdvertiseOmpIpv4s = make([]ServiceLANVPNAdvertiseOmpIpv4s, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNAdvertiseOmpIpv4s{}
			item.Protocol = types.StringNull()
			item.ProtocolVariable = types.StringNull()
			if t := v.Get("ompProtocol.optionType"); t.Exists() {
				va := v.Get("ompProtocol.value")
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
			if cValue := v.Get("prefixList"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Prefixes = make([]ServiceLANVPNAdvertiseOmpIpv4sPrefixes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceLANVPNAdvertiseOmpIpv4sPrefixes{}
					cItem.NetworkAddress = types.StringNull()
					cItem.NetworkAddressVariable = types.StringNull()
					if t := cv.Get("prefix.address.optionType"); t.Exists() {
						va := cv.Get("prefix.address.value")
						if t.String() == "variable" {
							cItem.NetworkAddressVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.NetworkAddress = types.StringValue(va.String())
						}
					}
					cItem.SubnetMask = types.StringNull()
					cItem.SubnetMaskVariable = types.StringNull()
					if t := cv.Get("prefix.mask.optionType"); t.Exists() {
						va := cv.Get("prefix.mask.value")
						if t.String() == "variable" {
							cItem.SubnetMaskVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.SubnetMask = types.StringValue(va.String())
						}
					}
					cItem.AggregateOnly = types.BoolNull()

					if t := cv.Get("aggregateOnly.optionType"); t.Exists() {
						va := cv.Get("aggregateOnly.value")
						if t.String() == "global" {
							cItem.AggregateOnly = types.BoolValue(va.Bool())
						}
					}
					cItem.Region = types.StringNull()
					cItem.RegionVariable = types.StringNull()
					if t := cv.Get("region.optionType"); t.Exists() {
						va := cv.Get("region.value")
						if t.String() == "variable" {
							cItem.RegionVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Region = types.StringValue(va.String())
						}
					}
					item.Prefixes = append(item.Prefixes, cItem)
					return true
				})
			}
			data.AdvertiseOmpIpv4s = append(data.AdvertiseOmpIpv4s, item)
			return true
		})
	} else {
		data.AdvertiseOmpIpv4s = nil
	}
	if !fullRead {
		resultAdvertiseOmpIpv4s := make([]ServiceLANVPNAdvertiseOmpIpv4s, 0, len(data.AdvertiseOmpIpv4s))
		matchedAdvertiseOmpIpv4s := make([]bool, len(data.AdvertiseOmpIpv4s))
		for _, oldItem := range oldAdvertiseOmpIpv4s {
			for ni := range data.AdvertiseOmpIpv4s {
				if matchedAdvertiseOmpIpv4s[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.ProtocolVariable.ValueString() != "" || data.AdvertiseOmpIpv4s[ni].ProtocolVariable.ValueString() != "") {
					if oldItem.ProtocolVariable.ValueString() != data.AdvertiseOmpIpv4s[ni].ProtocolVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.Protocol.ValueString() != data.AdvertiseOmpIpv4s[ni].Protocol.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedAdvertiseOmpIpv4s[ni] = true
					{
						resultC := make([]ServiceLANVPNAdvertiseOmpIpv4sPrefixes, 0, len(data.AdvertiseOmpIpv4s[ni].Prefixes))
						matchedC := make([]bool, len(data.AdvertiseOmpIpv4s[ni].Prefixes))
						for _, oldCItem := range oldItem.Prefixes {
							for nci := range data.AdvertiseOmpIpv4s[ni].Prefixes {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC && (oldCItem.NetworkAddressVariable.ValueString() != "" || data.AdvertiseOmpIpv4s[ni].Prefixes[nci].NetworkAddressVariable.ValueString() != "") {
									if oldCItem.NetworkAddressVariable.ValueString() != data.AdvertiseOmpIpv4s[ni].Prefixes[nci].NetworkAddressVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.NetworkAddress.ValueString() != data.AdvertiseOmpIpv4s[ni].Prefixes[nci].NetworkAddress.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.AdvertiseOmpIpv4s[ni].Prefixes[nci])
									break
								}
							}
						}
						for nci := range data.AdvertiseOmpIpv4s[ni].Prefixes {
							if !matchedC[nci] {
								resultC = append(resultC, data.AdvertiseOmpIpv4s[ni].Prefixes[nci])
							}
						}
						data.AdvertiseOmpIpv4s[ni].Prefixes = resultC
					}
					resultAdvertiseOmpIpv4s = append(resultAdvertiseOmpIpv4s, data.AdvertiseOmpIpv4s[ni])
					break
				}
			}
		}
		for ni := range data.AdvertiseOmpIpv4s {
			if !matchedAdvertiseOmpIpv4s[ni] {
				resultAdvertiseOmpIpv4s = append(resultAdvertiseOmpIpv4s, data.AdvertiseOmpIpv4s[ni])
			}
		}
		data.AdvertiseOmpIpv4s = resultAdvertiseOmpIpv4s
	}
	oldAdvertiseOmpIpv6s := data.AdvertiseOmpIpv6s
	if value := res.Get(path + "ompAdvertiseIpv6"); value.Exists() && len(value.Array()) > 0 {
		data.AdvertiseOmpIpv6s = make([]ServiceLANVPNAdvertiseOmpIpv6s, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNAdvertiseOmpIpv6s{}
			item.Protocol = types.StringNull()
			item.ProtocolVariable = types.StringNull()
			if t := v.Get("ompProtocol.optionType"); t.Exists() {
				va := v.Get("ompProtocol.value")
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
			item.ProtocolSubType = types.StringNull()
			item.ProtocolSubTypeVariable = types.StringNull()
			if t := v.Get("protocolSubType.optionType"); t.Exists() {
				va := v.Get("protocolSubType.value")
				if t.String() == "variable" {
					item.ProtocolSubTypeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.ProtocolSubType = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("prefixList"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Prefixes = make([]ServiceLANVPNAdvertiseOmpIpv6sPrefixes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceLANVPNAdvertiseOmpIpv6sPrefixes{}
					cItem.Prefix = types.StringNull()
					cItem.PrefixVariable = types.StringNull()
					if t := cv.Get("prefix.optionType"); t.Exists() {
						va := cv.Get("prefix.value")
						if t.String() == "variable" {
							cItem.PrefixVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Prefix = types.StringValue(va.String())
						}
					}
					cItem.AggregateOnly = types.BoolNull()

					if t := cv.Get("aggregateOnly.optionType"); t.Exists() {
						va := cv.Get("aggregateOnly.value")
						if t.String() == "global" {
							cItem.AggregateOnly = types.BoolValue(va.Bool())
						}
					}
					cItem.Region = types.StringNull()
					cItem.RegionVariable = types.StringNull()
					if t := cv.Get("region.optionType"); t.Exists() {
						va := cv.Get("region.value")
						if t.String() == "variable" {
							cItem.RegionVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Region = types.StringValue(va.String())
						}
					}
					item.Prefixes = append(item.Prefixes, cItem)
					return true
				})
			}
			data.AdvertiseOmpIpv6s = append(data.AdvertiseOmpIpv6s, item)
			return true
		})
	} else {
		data.AdvertiseOmpIpv6s = nil
	}
	if !fullRead {
		resultAdvertiseOmpIpv6s := make([]ServiceLANVPNAdvertiseOmpIpv6s, 0, len(data.AdvertiseOmpIpv6s))
		matchedAdvertiseOmpIpv6s := make([]bool, len(data.AdvertiseOmpIpv6s))
		for _, oldItem := range oldAdvertiseOmpIpv6s {
			for ni := range data.AdvertiseOmpIpv6s {
				if matchedAdvertiseOmpIpv6s[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.ProtocolVariable.ValueString() != "" || data.AdvertiseOmpIpv6s[ni].ProtocolVariable.ValueString() != "") {
					if oldItem.ProtocolVariable.ValueString() != data.AdvertiseOmpIpv6s[ni].ProtocolVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.Protocol.ValueString() != data.AdvertiseOmpIpv6s[ni].Protocol.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedAdvertiseOmpIpv6s[ni] = true
					{
						resultC := make([]ServiceLANVPNAdvertiseOmpIpv6sPrefixes, 0, len(data.AdvertiseOmpIpv6s[ni].Prefixes))
						matchedC := make([]bool, len(data.AdvertiseOmpIpv6s[ni].Prefixes))
						for _, oldCItem := range oldItem.Prefixes {
							for nci := range data.AdvertiseOmpIpv6s[ni].Prefixes {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC && (oldCItem.PrefixVariable.ValueString() != "" || data.AdvertiseOmpIpv6s[ni].Prefixes[nci].PrefixVariable.ValueString() != "") {
									if oldCItem.PrefixVariable.ValueString() != data.AdvertiseOmpIpv6s[ni].Prefixes[nci].PrefixVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.Prefix.ValueString() != data.AdvertiseOmpIpv6s[ni].Prefixes[nci].Prefix.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.AdvertiseOmpIpv6s[ni].Prefixes[nci])
									break
								}
							}
						}
						for nci := range data.AdvertiseOmpIpv6s[ni].Prefixes {
							if !matchedC[nci] {
								resultC = append(resultC, data.AdvertiseOmpIpv6s[ni].Prefixes[nci])
							}
						}
						data.AdvertiseOmpIpv6s[ni].Prefixes = resultC
					}
					resultAdvertiseOmpIpv6s = append(resultAdvertiseOmpIpv6s, data.AdvertiseOmpIpv6s[ni])
					break
				}
			}
		}
		for ni := range data.AdvertiseOmpIpv6s {
			if !matchedAdvertiseOmpIpv6s[ni] {
				resultAdvertiseOmpIpv6s = append(resultAdvertiseOmpIpv6s, data.AdvertiseOmpIpv6s[ni])
			}
		}
		data.AdvertiseOmpIpv6s = resultAdvertiseOmpIpv6s
	}
	oldIpv4StaticRoutes := data.Ipv4StaticRoutes
	if value := res.Get(path + "ipv4Route"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv4StaticRoutes = make([]ServiceLANVPNIpv4StaticRoutes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNIpv4StaticRoutes{}
			item.NetworkAddress = types.StringNull()
			item.NetworkAddressVariable = types.StringNull()
			if t := v.Get("prefix.ipAddress.optionType"); t.Exists() {
				va := v.Get("prefix.ipAddress.value")
				if t.String() == "variable" {
					item.NetworkAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NetworkAddress = types.StringValue(va.String())
				}
			}
			item.SubnetMask = types.StringNull()
			item.SubnetMaskVariable = types.StringNull()
			if t := v.Get("prefix.subnetMask.optionType"); t.Exists() {
				va := v.Get("prefix.subnetMask.value")
				if t.String() == "variable" {
					item.SubnetMaskVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SubnetMask = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("oneOfIpRoute.nextHopContainer.nextHop"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.NextHops = make([]ServiceLANVPNIpv4StaticRoutesNextHops, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceLANVPNIpv4StaticRoutesNextHops{}
					cItem.Address = types.StringNull()
					cItem.AddressVariable = types.StringNull()
					if t := cv.Get("address.optionType"); t.Exists() {
						va := cv.Get("address.value")
						if t.String() == "variable" {
							cItem.AddressVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					cItem.AdministrativeDistance = types.Int64Null()
					cItem.AdministrativeDistanceVariable = types.StringNull()
					if t := cv.Get("distance.optionType"); t.Exists() {
						va := cv.Get("distance.value")
						if t.String() == "variable" {
							cItem.AdministrativeDistanceVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.AdministrativeDistance = types.Int64Value(va.Int())
						}
					}
					item.NextHops = append(item.NextHops, cItem)
					return true
				})
				item.Gateway = types.StringValue("nextHop")
			}
			if cValue := v.Get("oneOfIpRoute.nextHopContainer.nextHopWithTracker"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.NextHopWithTrackers = make([]ServiceLANVPNIpv4StaticRoutesNextHopWithTrackers, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceLANVPNIpv4StaticRoutesNextHopWithTrackers{}
					cItem.Address = types.StringNull()
					cItem.AddressVariable = types.StringNull()
					if t := cv.Get("address.optionType"); t.Exists() {
						va := cv.Get("address.value")
						if t.String() == "variable" {
							cItem.AddressVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					cItem.AdministrativeDistance = types.Int64Null()
					cItem.AdministrativeDistanceVariable = types.StringNull()
					if t := cv.Get("distance.optionType"); t.Exists() {
						va := cv.Get("distance.value")
						if t.String() == "variable" {
							cItem.AdministrativeDistanceVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.AdministrativeDistance = types.Int64Value(va.Int())
						}
					}
					cItem.TrackerId = types.StringNull()

					if t := cv.Get("tracker.refId.optionType"); t.Exists() {
						va := cv.Get("tracker.refId.value")
						if t.String() == "global" {
							cItem.TrackerId = types.StringValue(va.String())
						}
					}
					item.NextHopWithTrackers = append(item.NextHopWithTrackers, cItem)
					return true
				})
				item.Gateway = types.StringValue("nextHop")
			}
			item.Null0 = types.BoolNull()

			if t := v.Get("oneOfIpRoute.null0.optionType"); t.Exists() {
				va := v.Get("oneOfIpRoute.null0.value")
				if t.String() == "global" {
					item.Null0 = types.BoolValue(va.Bool())
				}
				item.Gateway = types.StringValue("null0")
			}
			item.AdministrativeDistance = types.Int64Null()
			item.AdministrativeDistanceVariable = types.StringNull()
			if t := v.Get("oneOfIpRoute.distance.optionType"); t.Exists() {
				va := v.Get("oneOfIpRoute.distance.value")
				if t.String() == "variable" {
					item.AdministrativeDistanceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AdministrativeDistance = types.Int64Value(va.Int())
				}
				item.Gateway = types.StringValue("null0")
			}
			item.Dhcp = types.BoolNull()

			if t := v.Get("oneOfIpRoute.dhcp.optionType"); t.Exists() {
				va := v.Get("oneOfIpRoute.dhcp.value")
				if t.String() == "global" {
					item.Dhcp = types.BoolValue(va.Bool())
				}
				item.Gateway = types.StringValue("dhcp")
			}
			item.Vpn = types.BoolNull()

			if t := v.Get("oneOfIpRoute.vpn.optionType"); t.Exists() {
				va := v.Get("oneOfIpRoute.vpn.value")
				if t.String() == "global" {
					item.Vpn = types.BoolValue(va.Bool())
				}
				item.Gateway = types.StringValue("vpn")
			}
			if cValue := v.Get("oneOfIpRoute.interfaceContainer.ipStaticRouteInterface"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.IpStaticRouteInterface = make([]ServiceLANVPNIpv4StaticRoutesIpStaticRouteInterface, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceLANVPNIpv4StaticRoutesIpStaticRouteInterface{}
					cItem.InterfaceName = types.StringNull()
					cItem.InterfaceNameVariable = types.StringNull()
					if t := cv.Get("interfaceName.optionType"); t.Exists() {
						va := cv.Get("interfaceName.value")
						if t.String() == "variable" {
							cItem.InterfaceNameVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.InterfaceName = types.StringValue(va.String())
						}
					}
					if ccValue := cv.Get("nextHop"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.NextHop = make([]ServiceLANVPNIpv4StaticRoutesIpStaticRouteInterfaceNextHop, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := ServiceLANVPNIpv4StaticRoutesIpStaticRouteInterfaceNextHop{}
							ccItem.Address = types.StringNull()
							ccItem.AddressVariable = types.StringNull()
							if t := ccv.Get("address.optionType"); t.Exists() {
								va := ccv.Get("address.value")
								if t.String() == "variable" {
									ccItem.AddressVariable = types.StringValue(va.String())
								} else if t.String() == "global" {
									ccItem.Address = types.StringValue(va.String())
								}
							}
							ccItem.AdministrativeDistance = types.Int64Null()
							ccItem.AdministrativeDistanceVariable = types.StringNull()
							if t := ccv.Get("distance.optionType"); t.Exists() {
								va := ccv.Get("distance.value")
								if t.String() == "variable" {
									ccItem.AdministrativeDistanceVariable = types.StringValue(va.String())
								} else if t.String() == "global" {
									ccItem.AdministrativeDistance = types.Int64Value(va.Int())
								}
							}
							cItem.NextHop = append(cItem.NextHop, ccItem)
							return true
						})
					}
					item.IpStaticRouteInterface = append(item.IpStaticRouteInterface, cItem)
					return true
				})
				item.Gateway = types.StringValue("staticRouteInterface")
			}
			data.Ipv4StaticRoutes = append(data.Ipv4StaticRoutes, item)
			return true
		})
	} else {
		data.Ipv4StaticRoutes = nil
	}
	if !fullRead {
		resultIpv4StaticRoutes := make([]ServiceLANVPNIpv4StaticRoutes, 0, len(data.Ipv4StaticRoutes))
		matchedIpv4StaticRoutes := make([]bool, len(data.Ipv4StaticRoutes))
		for _, oldItem := range oldIpv4StaticRoutes {
			for ni := range data.Ipv4StaticRoutes {
				if matchedIpv4StaticRoutes[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.NetworkAddressVariable.ValueString() != "" || data.Ipv4StaticRoutes[ni].NetworkAddressVariable.ValueString() != "") {
					if oldItem.NetworkAddressVariable.ValueString() != data.Ipv4StaticRoutes[ni].NetworkAddressVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.NetworkAddress.ValueString() != data.Ipv4StaticRoutes[ni].NetworkAddress.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedIpv4StaticRoutes[ni] = true
					{
						resultC := make([]ServiceLANVPNIpv4StaticRoutesNextHops, 0, len(data.Ipv4StaticRoutes[ni].NextHops))
						matchedC := make([]bool, len(data.Ipv4StaticRoutes[ni].NextHops))
						for _, oldCItem := range oldItem.NextHops {
							for nci := range data.Ipv4StaticRoutes[ni].NextHops {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC && (oldCItem.AddressVariable.ValueString() != "" || data.Ipv4StaticRoutes[ni].NextHops[nci].AddressVariable.ValueString() != "") {
									if oldCItem.AddressVariable.ValueString() != data.Ipv4StaticRoutes[ni].NextHops[nci].AddressVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.Address.ValueString() != data.Ipv4StaticRoutes[ni].NextHops[nci].Address.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.Ipv4StaticRoutes[ni].NextHops[nci])
									break
								}
							}
						}
						for nci := range data.Ipv4StaticRoutes[ni].NextHops {
							if !matchedC[nci] {
								resultC = append(resultC, data.Ipv4StaticRoutes[ni].NextHops[nci])
							}
						}
						data.Ipv4StaticRoutes[ni].NextHops = resultC
					}
					{
						resultC := make([]ServiceLANVPNIpv4StaticRoutesNextHopWithTrackers, 0, len(data.Ipv4StaticRoutes[ni].NextHopWithTrackers))
						matchedC := make([]bool, len(data.Ipv4StaticRoutes[ni].NextHopWithTrackers))
						for _, oldCItem := range oldItem.NextHopWithTrackers {
							for nci := range data.Ipv4StaticRoutes[ni].NextHopWithTrackers {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC && (oldCItem.AddressVariable.ValueString() != "" || data.Ipv4StaticRoutes[ni].NextHopWithTrackers[nci].AddressVariable.ValueString() != "") {
									if oldCItem.AddressVariable.ValueString() != data.Ipv4StaticRoutes[ni].NextHopWithTrackers[nci].AddressVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.Address.ValueString() != data.Ipv4StaticRoutes[ni].NextHopWithTrackers[nci].Address.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC && (oldCItem.AdministrativeDistanceVariable.ValueString() != "" || data.Ipv4StaticRoutes[ni].NextHopWithTrackers[nci].AdministrativeDistanceVariable.ValueString() != "") {
									if oldCItem.AdministrativeDistanceVariable.ValueString() != data.Ipv4StaticRoutes[ni].NextHopWithTrackers[nci].AdministrativeDistanceVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.AdministrativeDistance.ValueInt64() != data.Ipv4StaticRoutes[ni].NextHopWithTrackers[nci].AdministrativeDistance.ValueInt64() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									if oldCItem.TrackerId.ValueString() != data.Ipv4StaticRoutes[ni].NextHopWithTrackers[nci].TrackerId.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.Ipv4StaticRoutes[ni].NextHopWithTrackers[nci])
									break
								}
							}
						}
						for nci := range data.Ipv4StaticRoutes[ni].NextHopWithTrackers {
							if !matchedC[nci] {
								resultC = append(resultC, data.Ipv4StaticRoutes[ni].NextHopWithTrackers[nci])
							}
						}
						data.Ipv4StaticRoutes[ni].NextHopWithTrackers = resultC
					}
					{
						resultC := make([]ServiceLANVPNIpv4StaticRoutesIpStaticRouteInterface, 0, len(data.Ipv4StaticRoutes[ni].IpStaticRouteInterface))
						matchedC := make([]bool, len(data.Ipv4StaticRoutes[ni].IpStaticRouteInterface))
						for _, oldCItem := range oldItem.IpStaticRouteInterface {
							for nci := range data.Ipv4StaticRoutes[ni].IpStaticRouteInterface {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC && (oldCItem.InterfaceNameVariable.ValueString() != "" || data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci].InterfaceNameVariable.ValueString() != "") {
									if oldCItem.InterfaceNameVariable.ValueString() != data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci].InterfaceNameVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.InterfaceName.ValueString() != data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci].InterfaceName.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									{
										resultCC := make([]ServiceLANVPNIpv4StaticRoutesIpStaticRouteInterfaceNextHop, 0, len(data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci].NextHop))
										matchedCC := make([]bool, len(data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci].NextHop))
										for _, oldCCItem := range oldCItem.NextHop {
											for ncci := range data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci].NextHop {
												if matchedCC[ncci] {
													continue
												}
												keyMatchCC := true
												if keyMatchCC && (oldCCItem.AddressVariable.ValueString() != "" || data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci].NextHop[ncci].AddressVariable.ValueString() != "") {
													if oldCCItem.AddressVariable.ValueString() != data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci].NextHop[ncci].AddressVariable.ValueString() {
														keyMatchCC = false
													}
												} else if keyMatchCC {
													if oldCCItem.Address.ValueString() != data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci].NextHop[ncci].Address.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC && (oldCCItem.AdministrativeDistanceVariable.ValueString() != "" || data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci].NextHop[ncci].AdministrativeDistanceVariable.ValueString() != "") {
													if oldCCItem.AdministrativeDistanceVariable.ValueString() != data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci].NextHop[ncci].AdministrativeDistanceVariable.ValueString() {
														keyMatchCC = false
													}
												} else if keyMatchCC {
													if oldCCItem.AdministrativeDistance.ValueInt64() != data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci].NextHop[ncci].AdministrativeDistance.ValueInt64() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													matchedCC[ncci] = true
													resultCC = append(resultCC, data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci].NextHop[ncci])
													break
												}
											}
										}
										for ncci := range data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci].NextHop {
											if !matchedCC[ncci] {
												resultCC = append(resultCC, data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci].NextHop[ncci])
											}
										}
										data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci].NextHop = resultCC
									}
									resultC = append(resultC, data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci])
									break
								}
							}
						}
						for nci := range data.Ipv4StaticRoutes[ni].IpStaticRouteInterface {
							if !matchedC[nci] {
								resultC = append(resultC, data.Ipv4StaticRoutes[ni].IpStaticRouteInterface[nci])
							}
						}
						data.Ipv4StaticRoutes[ni].IpStaticRouteInterface = resultC
					}
					resultIpv4StaticRoutes = append(resultIpv4StaticRoutes, data.Ipv4StaticRoutes[ni])
					break
				}
			}
		}
		for ni := range data.Ipv4StaticRoutes {
			if !matchedIpv4StaticRoutes[ni] {
				resultIpv4StaticRoutes = append(resultIpv4StaticRoutes, data.Ipv4StaticRoutes[ni])
			}
		}
		data.Ipv4StaticRoutes = resultIpv4StaticRoutes
	}
	oldIpv6StaticRoutes := data.Ipv6StaticRoutes
	if value := res.Get(path + "ipv6Route"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv6StaticRoutes = make([]ServiceLANVPNIpv6StaticRoutes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNIpv6StaticRoutes{}
			item.Prefix = types.StringNull()
			item.PrefixVariable = types.StringNull()
			if t := v.Get("prefix.optionType"); t.Exists() {
				va := v.Get("prefix.value")
				if t.String() == "variable" {
					item.PrefixVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Prefix = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("oneOfIpRoute.nextHopContainer.nextHop"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.NextHops = make([]ServiceLANVPNIpv6StaticRoutesNextHops, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceLANVPNIpv6StaticRoutesNextHops{}
					cItem.Address = types.StringNull()
					cItem.AddressVariable = types.StringNull()
					if t := cv.Get("address.optionType"); t.Exists() {
						va := cv.Get("address.value")
						if t.String() == "variable" {
							cItem.AddressVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Address = types.StringValue(va.String())
						}
					}
					cItem.AdministrativeDistance = types.Int64Null()
					cItem.AdministrativeDistanceVariable = types.StringNull()
					if t := cv.Get("distance.optionType"); t.Exists() {
						va := cv.Get("distance.value")
						if t.String() == "variable" {
							cItem.AdministrativeDistanceVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.AdministrativeDistance = types.Int64Value(va.Int())
						}
					}
					item.NextHops = append(item.NextHops, cItem)
					return true
				})
				item.Gateway = types.StringValue("nextHop")
			}
			item.Null0 = types.BoolNull()

			if t := v.Get("oneOfIpRoute.null0.optionType"); t.Exists() {
				va := v.Get("oneOfIpRoute.null0.value")
				if t.String() == "global" {
					item.Null0 = types.BoolValue(va.Bool())
				}
				item.Gateway = types.StringValue("null0")
			}
			item.Nat = types.StringNull()
			item.NatVariable = types.StringNull()
			if t := v.Get("oneOfIpRoute.nat.optionType"); t.Exists() {
				va := v.Get("oneOfIpRoute.nat.value")
				if t.String() == "variable" {
					item.NatVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Nat = types.StringValue(va.String())
				}
				item.Gateway = types.StringValue("nat")
			}
			if cValue := v.Get("oneOfIpRoute.interfaceContainer.ipv6StaticRouteInterface"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Ipv6StaticRouteInterface = make([]ServiceLANVPNIpv6StaticRoutesIpv6StaticRouteInterface, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceLANVPNIpv6StaticRoutesIpv6StaticRouteInterface{}
					cItem.InterfaceName = types.StringNull()
					cItem.InterfaceNameVariable = types.StringNull()
					if t := cv.Get("interfaceName.optionType"); t.Exists() {
						va := cv.Get("interfaceName.value")
						if t.String() == "variable" {
							cItem.InterfaceNameVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.InterfaceName = types.StringValue(va.String())
						}
					}
					if ccValue := cv.Get("nextHop"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.NextHop = make([]ServiceLANVPNIpv6StaticRoutesIpv6StaticRouteInterfaceNextHop, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := ServiceLANVPNIpv6StaticRoutesIpv6StaticRouteInterfaceNextHop{}
							ccItem.Address = types.StringNull()
							ccItem.AddressVariable = types.StringNull()
							if t := ccv.Get("address.optionType"); t.Exists() {
								va := ccv.Get("address.value")
								if t.String() == "variable" {
									ccItem.AddressVariable = types.StringValue(va.String())
								} else if t.String() == "global" {
									ccItem.Address = types.StringValue(va.String())
								}
							}
							ccItem.AdministrativeDistance = types.Int64Null()
							ccItem.AdministrativeDistanceVariable = types.StringNull()
							if t := ccv.Get("distance.optionType"); t.Exists() {
								va := ccv.Get("distance.value")
								if t.String() == "variable" {
									ccItem.AdministrativeDistanceVariable = types.StringValue(va.String())
								} else if t.String() == "global" {
									ccItem.AdministrativeDistance = types.Int64Value(va.Int())
								}
							}
							cItem.NextHop = append(cItem.NextHop, ccItem)
							return true
						})
					}
					item.Ipv6StaticRouteInterface = append(item.Ipv6StaticRouteInterface, cItem)
					return true
				})
				item.Gateway = types.StringValue("staticRouteInterface")
			}
			data.Ipv6StaticRoutes = append(data.Ipv6StaticRoutes, item)
			return true
		})
	} else {
		data.Ipv6StaticRoutes = nil
	}
	if !fullRead {
		resultIpv6StaticRoutes := make([]ServiceLANVPNIpv6StaticRoutes, 0, len(data.Ipv6StaticRoutes))
		matchedIpv6StaticRoutes := make([]bool, len(data.Ipv6StaticRoutes))
		for _, oldItem := range oldIpv6StaticRoutes {
			for ni := range data.Ipv6StaticRoutes {
				if matchedIpv6StaticRoutes[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.PrefixVariable.ValueString() != "" || data.Ipv6StaticRoutes[ni].PrefixVariable.ValueString() != "") {
					if oldItem.PrefixVariable.ValueString() != data.Ipv6StaticRoutes[ni].PrefixVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.Prefix.ValueString() != data.Ipv6StaticRoutes[ni].Prefix.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedIpv6StaticRoutes[ni] = true
					{
						resultC := make([]ServiceLANVPNIpv6StaticRoutesNextHops, 0, len(data.Ipv6StaticRoutes[ni].NextHops))
						matchedC := make([]bool, len(data.Ipv6StaticRoutes[ni].NextHops))
						for _, oldCItem := range oldItem.NextHops {
							for nci := range data.Ipv6StaticRoutes[ni].NextHops {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC && (oldCItem.AddressVariable.ValueString() != "" || data.Ipv6StaticRoutes[ni].NextHops[nci].AddressVariable.ValueString() != "") {
									if oldCItem.AddressVariable.ValueString() != data.Ipv6StaticRoutes[ni].NextHops[nci].AddressVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.Address.ValueString() != data.Ipv6StaticRoutes[ni].NextHops[nci].Address.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.Ipv6StaticRoutes[ni].NextHops[nci])
									break
								}
							}
						}
						for nci := range data.Ipv6StaticRoutes[ni].NextHops {
							if !matchedC[nci] {
								resultC = append(resultC, data.Ipv6StaticRoutes[ni].NextHops[nci])
							}
						}
						data.Ipv6StaticRoutes[ni].NextHops = resultC
					}
					{
						resultC := make([]ServiceLANVPNIpv6StaticRoutesIpv6StaticRouteInterface, 0, len(data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface))
						matchedC := make([]bool, len(data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface))
						for _, oldCItem := range oldItem.Ipv6StaticRouteInterface {
							for nci := range data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC && (oldCItem.InterfaceNameVariable.ValueString() != "" || data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci].InterfaceNameVariable.ValueString() != "") {
									if oldCItem.InterfaceNameVariable.ValueString() != data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci].InterfaceNameVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.InterfaceName.ValueString() != data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci].InterfaceName.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									{
										resultCC := make([]ServiceLANVPNIpv6StaticRoutesIpv6StaticRouteInterfaceNextHop, 0, len(data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci].NextHop))
										matchedCC := make([]bool, len(data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci].NextHop))
										for _, oldCCItem := range oldCItem.NextHop {
											for ncci := range data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci].NextHop {
												if matchedCC[ncci] {
													continue
												}
												keyMatchCC := true
												if keyMatchCC && (oldCCItem.AddressVariable.ValueString() != "" || data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci].NextHop[ncci].AddressVariable.ValueString() != "") {
													if oldCCItem.AddressVariable.ValueString() != data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci].NextHop[ncci].AddressVariable.ValueString() {
														keyMatchCC = false
													}
												} else if keyMatchCC {
													if oldCCItem.Address.ValueString() != data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci].NextHop[ncci].Address.ValueString() {
														keyMatchCC = false
													}
												}
												if keyMatchCC && (oldCCItem.AdministrativeDistanceVariable.ValueString() != "" || data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci].NextHop[ncci].AdministrativeDistanceVariable.ValueString() != "") {
													if oldCCItem.AdministrativeDistanceVariable.ValueString() != data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci].NextHop[ncci].AdministrativeDistanceVariable.ValueString() {
														keyMatchCC = false
													}
												} else if keyMatchCC {
													if oldCCItem.AdministrativeDistance.ValueInt64() != data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci].NextHop[ncci].AdministrativeDistance.ValueInt64() {
														keyMatchCC = false
													}
												}
												if keyMatchCC {
													matchedCC[ncci] = true
													resultCC = append(resultCC, data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci].NextHop[ncci])
													break
												}
											}
										}
										for ncci := range data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci].NextHop {
											if !matchedCC[ncci] {
												resultCC = append(resultCC, data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci].NextHop[ncci])
											}
										}
										data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci].NextHop = resultCC
									}
									resultC = append(resultC, data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci])
									break
								}
							}
						}
						for nci := range data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface {
							if !matchedC[nci] {
								resultC = append(resultC, data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface[nci])
							}
						}
						data.Ipv6StaticRoutes[ni].Ipv6StaticRouteInterface = resultC
					}
					resultIpv6StaticRoutes = append(resultIpv6StaticRoutes, data.Ipv6StaticRoutes[ni])
					break
				}
			}
		}
		for ni := range data.Ipv6StaticRoutes {
			if !matchedIpv6StaticRoutes[ni] {
				resultIpv6StaticRoutes = append(resultIpv6StaticRoutes, data.Ipv6StaticRoutes[ni])
			}
		}
		data.Ipv6StaticRoutes = resultIpv6StaticRoutes
	}
	oldServices := data.Services
	if value := res.Get(path + "service"); value.Exists() && len(value.Array()) > 0 {
		data.Services = make([]ServiceLANVPNServices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNServices{}
			item.ServiceType = types.StringNull()
			item.ServiceTypeVariable = types.StringNull()
			if t := v.Get("serviceType.optionType"); t.Exists() {
				va := v.Get("serviceType.value")
				if t.String() == "variable" {
					item.ServiceTypeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.ServiceType = types.StringValue(va.String())
				}
			}
			item.Ipv4Addresses = types.SetNull(types.StringType)
			item.Ipv4AddressesVariable = types.StringNull()
			if t := v.Get("ipv4Addresses.optionType"); t.Exists() {
				va := v.Get("ipv4Addresses.value")
				if t.String() == "variable" {
					item.Ipv4AddressesVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Ipv4Addresses = helpers.GetStringSet(va.Array())
				}
			}
			item.Tracking = types.BoolNull()
			item.TrackingVariable = types.StringNull()
			if t := v.Get("tracking.optionType"); t.Exists() {
				va := v.Get("tracking.value")
				if t.String() == "variable" {
					item.TrackingVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Tracking = types.BoolValue(va.Bool())
				}
			}
			data.Services = append(data.Services, item)
			return true
		})
	} else {
		data.Services = nil
	}
	if !fullRead {
		resultServices := make([]ServiceLANVPNServices, 0, len(data.Services))
		matchedServices := make([]bool, len(data.Services))
		for _, oldItem := range oldServices {
			for ni := range data.Services {
				if matchedServices[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.ServiceTypeVariable.ValueString() != "" || data.Services[ni].ServiceTypeVariable.ValueString() != "") {
					if oldItem.ServiceTypeVariable.ValueString() != data.Services[ni].ServiceTypeVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.ServiceType.ValueString() != data.Services[ni].ServiceType.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedServices[ni] = true
					resultServices = append(resultServices, data.Services[ni])
					break
				}
			}
		}
		for ni := range data.Services {
			if !matchedServices[ni] {
				resultServices = append(resultServices, data.Services[ni])
			}
		}
		data.Services = resultServices
	}
	oldServiceRoutes := data.ServiceRoutes
	if value := res.Get(path + "serviceRoute"); value.Exists() && len(value.Array()) > 0 {
		data.ServiceRoutes = make([]ServiceLANVPNServiceRoutes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNServiceRoutes{}
			item.NetworkAddress = types.StringNull()
			item.NetworkAddressVariable = types.StringNull()
			if t := v.Get("prefix.ipAddress.optionType"); t.Exists() {
				va := v.Get("prefix.ipAddress.value")
				if t.String() == "variable" {
					item.NetworkAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NetworkAddress = types.StringValue(va.String())
				}
			}
			item.SubnetMask = types.StringNull()
			item.SubnetMaskVariable = types.StringNull()
			if t := v.Get("prefix.subnetMask.optionType"); t.Exists() {
				va := v.Get("prefix.subnetMask.value")
				if t.String() == "variable" {
					item.SubnetMaskVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SubnetMask = types.StringValue(va.String())
				}
			}
			item.Service = types.StringNull()
			item.ServiceVariable = types.StringNull()
			if t := v.Get("service.optionType"); t.Exists() {
				va := v.Get("service.value")
				if t.String() == "variable" {
					item.ServiceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Service = types.StringValue(va.String())
				}
			}
			item.Vpn = types.Int64Null()

			if t := v.Get("vpn.optionType"); t.Exists() {
				va := v.Get("vpn.value")
				if t.String() == "global" {
					item.Vpn = types.Int64Value(va.Int())
				}
			}
			item.SseInstance = types.StringNull()
			item.SseInstanceVariable = types.StringNull()
			if t := v.Get("sseInstance.optionType"); t.Exists() {
				va := v.Get("sseInstance.value")
				if t.String() == "variable" {
					item.SseInstanceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SseInstance = types.StringValue(va.String())
				}
			}
			data.ServiceRoutes = append(data.ServiceRoutes, item)
			return true
		})
	} else {
		data.ServiceRoutes = nil
	}
	if !fullRead {
		resultServiceRoutes := make([]ServiceLANVPNServiceRoutes, 0, len(data.ServiceRoutes))
		matchedServiceRoutes := make([]bool, len(data.ServiceRoutes))
		for _, oldItem := range oldServiceRoutes {
			for ni := range data.ServiceRoutes {
				if matchedServiceRoutes[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.NetworkAddressVariable.ValueString() != "" || data.ServiceRoutes[ni].NetworkAddressVariable.ValueString() != "") {
					if oldItem.NetworkAddressVariable.ValueString() != data.ServiceRoutes[ni].NetworkAddressVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.NetworkAddress.ValueString() != data.ServiceRoutes[ni].NetworkAddress.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedServiceRoutes[ni] = true
					resultServiceRoutes = append(resultServiceRoutes, data.ServiceRoutes[ni])
					break
				}
			}
		}
		for ni := range data.ServiceRoutes {
			if !matchedServiceRoutes[ni] {
				resultServiceRoutes = append(resultServiceRoutes, data.ServiceRoutes[ni])
			}
		}
		data.ServiceRoutes = resultServiceRoutes
	}
	oldGreRoutes := data.GreRoutes
	if value := res.Get(path + "greRoute"); value.Exists() && len(value.Array()) > 0 {
		data.GreRoutes = make([]ServiceLANVPNGreRoutes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNGreRoutes{}
			item.NetworkAddress = types.StringNull()
			item.NetworkAddressVariable = types.StringNull()
			if t := v.Get("prefix.ipAddress.optionType"); t.Exists() {
				va := v.Get("prefix.ipAddress.value")
				if t.String() == "variable" {
					item.NetworkAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NetworkAddress = types.StringValue(va.String())
				}
			}
			item.SubnetMask = types.StringNull()
			item.SubnetMaskVariable = types.StringNull()
			if t := v.Get("prefix.subnetMask.optionType"); t.Exists() {
				va := v.Get("prefix.subnetMask.value")
				if t.String() == "variable" {
					item.SubnetMaskVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SubnetMask = types.StringValue(va.String())
				}
			}
			item.Interface = types.SetNull(types.StringType)
			item.InterfaceVariable = types.StringNull()
			if t := v.Get("interface.optionType"); t.Exists() {
				va := v.Get("interface.value")
				if t.String() == "variable" {
					item.InterfaceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Interface = helpers.GetStringSet(va.Array())
				}
			}
			item.Vpn = types.Int64Null()

			if t := v.Get("vpn.optionType"); t.Exists() {
				va := v.Get("vpn.value")
				if t.String() == "global" {
					item.Vpn = types.Int64Value(va.Int())
				}
			}
			data.GreRoutes = append(data.GreRoutes, item)
			return true
		})
	} else {
		data.GreRoutes = nil
	}
	if !fullRead {
		resultGreRoutes := make([]ServiceLANVPNGreRoutes, 0, len(data.GreRoutes))
		matchedGreRoutes := make([]bool, len(data.GreRoutes))
		for _, oldItem := range oldGreRoutes {
			for ni := range data.GreRoutes {
				if matchedGreRoutes[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.NetworkAddressVariable.ValueString() != "" || data.GreRoutes[ni].NetworkAddressVariable.ValueString() != "") {
					if oldItem.NetworkAddressVariable.ValueString() != data.GreRoutes[ni].NetworkAddressVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.NetworkAddress.ValueString() != data.GreRoutes[ni].NetworkAddress.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedGreRoutes[ni] = true
					resultGreRoutes = append(resultGreRoutes, data.GreRoutes[ni])
					break
				}
			}
		}
		for ni := range data.GreRoutes {
			if !matchedGreRoutes[ni] {
				resultGreRoutes = append(resultGreRoutes, data.GreRoutes[ni])
			}
		}
		data.GreRoutes = resultGreRoutes
	}
	oldIpsecRoutes := data.IpsecRoutes
	if value := res.Get(path + "ipsecRoute"); value.Exists() && len(value.Array()) > 0 {
		data.IpsecRoutes = make([]ServiceLANVPNIpsecRoutes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNIpsecRoutes{}
			item.NetworkAddress = types.StringNull()
			item.NetworkAddressVariable = types.StringNull()
			if t := v.Get("prefix.ipAddress.optionType"); t.Exists() {
				va := v.Get("prefix.ipAddress.value")
				if t.String() == "variable" {
					item.NetworkAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NetworkAddress = types.StringValue(va.String())
				}
			}
			item.SubnetMask = types.StringNull()
			item.SubnetMaskVariable = types.StringNull()
			if t := v.Get("prefix.subnetMask.optionType"); t.Exists() {
				va := v.Get("prefix.subnetMask.value")
				if t.String() == "variable" {
					item.SubnetMaskVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SubnetMask = types.StringValue(va.String())
				}
			}
			item.Interface = types.SetNull(types.StringType)
			item.InterfaceVariable = types.StringNull()
			if t := v.Get("interface.optionType"); t.Exists() {
				va := v.Get("interface.value")
				if t.String() == "variable" {
					item.InterfaceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Interface = helpers.GetStringSet(va.Array())
				}
			}
			data.IpsecRoutes = append(data.IpsecRoutes, item)
			return true
		})
	} else {
		data.IpsecRoutes = nil
	}
	if !fullRead {
		resultIpsecRoutes := make([]ServiceLANVPNIpsecRoutes, 0, len(data.IpsecRoutes))
		matchedIpsecRoutes := make([]bool, len(data.IpsecRoutes))
		for _, oldItem := range oldIpsecRoutes {
			for ni := range data.IpsecRoutes {
				if matchedIpsecRoutes[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.NetworkAddressVariable.ValueString() != "" || data.IpsecRoutes[ni].NetworkAddressVariable.ValueString() != "") {
					if oldItem.NetworkAddressVariable.ValueString() != data.IpsecRoutes[ni].NetworkAddressVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.NetworkAddress.ValueString() != data.IpsecRoutes[ni].NetworkAddress.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedIpsecRoutes[ni] = true
					resultIpsecRoutes = append(resultIpsecRoutes, data.IpsecRoutes[ni])
					break
				}
			}
		}
		for ni := range data.IpsecRoutes {
			if !matchedIpsecRoutes[ni] {
				resultIpsecRoutes = append(resultIpsecRoutes, data.IpsecRoutes[ni])
			}
		}
		data.IpsecRoutes = resultIpsecRoutes
	}
	oldNatPools := data.NatPools
	if value := res.Get(path + "natPool"); value.Exists() && len(value.Array()) > 0 {
		data.NatPools = make([]ServiceLANVPNNatPools, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNNatPools{}
			item.NatPoolName = types.Int64Null()
			item.NatPoolNameVariable = types.StringNull()
			if t := v.Get("natPoolName.optionType"); t.Exists() {
				va := v.Get("natPoolName.value")
				if t.String() == "variable" {
					item.NatPoolNameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NatPoolName = types.Int64Value(va.Int())
				}
			}
			item.PrefixLength = types.Int64Null()
			item.PrefixLengthVariable = types.StringNull()
			if t := v.Get("prefixLength.optionType"); t.Exists() {
				va := v.Get("prefixLength.value")
				if t.String() == "variable" {
					item.PrefixLengthVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.PrefixLength = types.Int64Value(va.Int())
				}
			}
			item.RangeStart = types.StringNull()
			item.RangeStartVariable = types.StringNull()
			if t := v.Get("rangeStart.optionType"); t.Exists() {
				va := v.Get("rangeStart.value")
				if t.String() == "variable" {
					item.RangeStartVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RangeStart = types.StringValue(va.String())
				}
			}
			item.RangeEnd = types.StringNull()
			item.RangeEndVariable = types.StringNull()
			if t := v.Get("rangeEnd.optionType"); t.Exists() {
				va := v.Get("rangeEnd.value")
				if t.String() == "variable" {
					item.RangeEndVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RangeEnd = types.StringValue(va.String())
				}
			}
			item.Overload = types.BoolNull()
			item.OverloadVariable = types.StringNull()
			if t := v.Get("overload.optionType"); t.Exists() {
				va := v.Get("overload.value")
				if t.String() == "variable" {
					item.OverloadVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Overload = types.BoolValue(va.Bool())
				}
			}
			item.Direction = types.StringNull()
			item.DirectionVariable = types.StringNull()
			if t := v.Get("direction.optionType"); t.Exists() {
				va := v.Get("direction.value")
				if t.String() == "variable" {
					item.DirectionVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Direction = types.StringValue(va.String())
				}
			}
			item.TrackerObjectId = types.StringNull()

			if t := v.Get("trackingObject.trackerId.refId.optionType"); t.Exists() {
				va := v.Get("trackingObject.trackerId.refId.value")
				if t.String() == "global" {
					item.TrackerObjectId = types.StringValue(va.String())
				}
			}
			data.NatPools = append(data.NatPools, item)
			return true
		})
	} else {
		data.NatPools = nil
	}
	if !fullRead {
		resultNatPools := make([]ServiceLANVPNNatPools, 0, len(data.NatPools))
		matchedNatPools := make([]bool, len(data.NatPools))
		for _, oldItem := range oldNatPools {
			for ni := range data.NatPools {
				if matchedNatPools[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.NatPoolNameVariable.ValueString() != "" || data.NatPools[ni].NatPoolNameVariable.ValueString() != "") {
					if oldItem.NatPoolNameVariable.ValueString() != data.NatPools[ni].NatPoolNameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.NatPoolName.ValueInt64() != data.NatPools[ni].NatPoolName.ValueInt64() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedNatPools[ni] = true
					resultNatPools = append(resultNatPools, data.NatPools[ni])
					break
				}
			}
		}
		for ni := range data.NatPools {
			if !matchedNatPools[ni] {
				resultNatPools = append(resultNatPools, data.NatPools[ni])
			}
		}
		data.NatPools = resultNatPools
	}
	oldNatPortForwards := data.NatPortForwards
	if value := res.Get(path + "natPortForward"); value.Exists() && len(value.Array()) > 0 {
		data.NatPortForwards = make([]ServiceLANVPNNatPortForwards, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNNatPortForwards{}
			item.NatPoolName = types.Int64Null()
			item.NatPoolNameVariable = types.StringNull()
			if t := v.Get("natPoolName.optionType"); t.Exists() {
				va := v.Get("natPoolName.value")
				if t.String() == "variable" {
					item.NatPoolNameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NatPoolName = types.Int64Value(va.Int())
				}
			}
			item.SourcePort = types.Int64Null()
			item.SourcePortVariable = types.StringNull()
			if t := v.Get("sourcePort.optionType"); t.Exists() {
				va := v.Get("sourcePort.value")
				if t.String() == "variable" {
					item.SourcePortVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourcePort = types.Int64Value(va.Int())
				}
			}
			item.TranslatePort = types.Int64Null()
			item.TranslatePortVariable = types.StringNull()
			if t := v.Get("translatePort.optionType"); t.Exists() {
				va := v.Get("translatePort.value")
				if t.String() == "variable" {
					item.TranslatePortVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TranslatePort = types.Int64Value(va.Int())
				}
			}
			item.SourceIp = types.StringNull()
			item.SourceIpVariable = types.StringNull()
			if t := v.Get("sourceIp.optionType"); t.Exists() {
				va := v.Get("sourceIp.value")
				if t.String() == "variable" {
					item.SourceIpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourceIp = types.StringValue(va.String())
				}
			}
			item.TranslatedSourceIp = types.StringNull()
			item.TranslatedSourceIpVariable = types.StringNull()
			if t := v.Get("translatedSourceIp.optionType"); t.Exists() {
				va := v.Get("translatedSourceIp.value")
				if t.String() == "variable" {
					item.TranslatedSourceIpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TranslatedSourceIp = types.StringValue(va.String())
				}
			}
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
			data.NatPortForwards = append(data.NatPortForwards, item)
			return true
		})
	} else {
		data.NatPortForwards = nil
	}
	if !fullRead {
		resultNatPortForwards := make([]ServiceLANVPNNatPortForwards, 0, len(data.NatPortForwards))
		matchedNatPortForwards := make([]bool, len(data.NatPortForwards))
		for _, oldItem := range oldNatPortForwards {
			for ni := range data.NatPortForwards {
				if matchedNatPortForwards[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.NatPoolNameVariable.ValueString() != "" || data.NatPortForwards[ni].NatPoolNameVariable.ValueString() != "") {
					if oldItem.NatPoolNameVariable.ValueString() != data.NatPortForwards[ni].NatPoolNameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.NatPoolName.ValueInt64() != data.NatPortForwards[ni].NatPoolName.ValueInt64() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedNatPortForwards[ni] = true
					resultNatPortForwards = append(resultNatPortForwards, data.NatPortForwards[ni])
					break
				}
			}
		}
		for ni := range data.NatPortForwards {
			if !matchedNatPortForwards[ni] {
				resultNatPortForwards = append(resultNatPortForwards, data.NatPortForwards[ni])
			}
		}
		data.NatPortForwards = resultNatPortForwards
	}
	oldStaticNats := data.StaticNats
	if value := res.Get(path + "staticNat"); value.Exists() && len(value.Array()) > 0 {
		data.StaticNats = make([]ServiceLANVPNStaticNats, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNStaticNats{}
			item.NatPoolName = types.Int64Null()
			item.NatPoolNameVariable = types.StringNull()
			if t := v.Get("natPoolName.optionType"); t.Exists() {
				va := v.Get("natPoolName.value")
				if t.String() == "variable" {
					item.NatPoolNameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NatPoolName = types.Int64Value(va.Int())
				}
			}
			item.SourceIp = types.StringNull()
			item.SourceIpVariable = types.StringNull()
			if t := v.Get("sourceIp.optionType"); t.Exists() {
				va := v.Get("sourceIp.value")
				if t.String() == "variable" {
					item.SourceIpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourceIp = types.StringValue(va.String())
				}
			}
			item.TranslatedSourceIp = types.StringNull()
			item.TranslatedSourceIpVariable = types.StringNull()
			if t := v.Get("translatedSourceIp.optionType"); t.Exists() {
				va := v.Get("translatedSourceIp.value")
				if t.String() == "variable" {
					item.TranslatedSourceIpVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TranslatedSourceIp = types.StringValue(va.String())
				}
			}
			item.StaticNatDirection = types.StringNull()
			item.StaticNatDirectionVariable = types.StringNull()
			if t := v.Get("staticNatDirection.optionType"); t.Exists() {
				va := v.Get("staticNatDirection.value")
				if t.String() == "variable" {
					item.StaticNatDirectionVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.StaticNatDirection = types.StringValue(va.String())
				}
			}
			item.TrackerObjectId = types.StringNull()

			if t := v.Get("trackingObject.trackerId.refId.optionType"); t.Exists() {
				va := v.Get("trackingObject.trackerId.refId.value")
				if t.String() == "global" {
					item.TrackerObjectId = types.StringValue(va.String())
				}
			}
			data.StaticNats = append(data.StaticNats, item)
			return true
		})
	} else {
		data.StaticNats = nil
	}
	if !fullRead {
		resultStaticNats := make([]ServiceLANVPNStaticNats, 0, len(data.StaticNats))
		matchedStaticNats := make([]bool, len(data.StaticNats))
		for _, oldItem := range oldStaticNats {
			for ni := range data.StaticNats {
				if matchedStaticNats[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.NatPoolNameVariable.ValueString() != "" || data.StaticNats[ni].NatPoolNameVariable.ValueString() != "") {
					if oldItem.NatPoolNameVariable.ValueString() != data.StaticNats[ni].NatPoolNameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.NatPoolName.ValueInt64() != data.StaticNats[ni].NatPoolName.ValueInt64() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedStaticNats[ni] = true
					resultStaticNats = append(resultStaticNats, data.StaticNats[ni])
					break
				}
			}
		}
		for ni := range data.StaticNats {
			if !matchedStaticNats[ni] {
				resultStaticNats = append(resultStaticNats, data.StaticNats[ni])
			}
		}
		data.StaticNats = resultStaticNats
	}
	oldStaticNatSubnets := data.StaticNatSubnets
	if value := res.Get(path + "staticNatSubnet"); value.Exists() && len(value.Array()) > 0 {
		data.StaticNatSubnets = make([]ServiceLANVPNStaticNatSubnets, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNStaticNatSubnets{}
			item.SourceIpSubnet = types.StringNull()
			item.SourceIpSubnetVariable = types.StringNull()
			if t := v.Get("sourceIpSubnet.optionType"); t.Exists() {
				va := v.Get("sourceIpSubnet.value")
				if t.String() == "variable" {
					item.SourceIpSubnetVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourceIpSubnet = types.StringValue(va.String())
				}
			}
			item.TranslatedSourceIpSubnet = types.StringNull()
			item.TranslatedSourceIpSubnetVariable = types.StringNull()
			if t := v.Get("translatedSourceIpSubnet.optionType"); t.Exists() {
				va := v.Get("translatedSourceIpSubnet.value")
				if t.String() == "variable" {
					item.TranslatedSourceIpSubnetVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.TranslatedSourceIpSubnet = types.StringValue(va.String())
				}
			}
			item.PrefixLength = types.Int64Null()
			item.PrefixLengthVariable = types.StringNull()
			if t := v.Get("prefixLength.optionType"); t.Exists() {
				va := v.Get("prefixLength.value")
				if t.String() == "variable" {
					item.PrefixLengthVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.PrefixLength = types.Int64Value(va.Int())
				}
			}
			item.StaticNatDirection = types.StringNull()
			item.StaticNatDirectionVariable = types.StringNull()
			if t := v.Get("staticNatDirection.optionType"); t.Exists() {
				va := v.Get("staticNatDirection.value")
				if t.String() == "variable" {
					item.StaticNatDirectionVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.StaticNatDirection = types.StringValue(va.String())
				}
			}
			item.TrackerObjectId = types.StringNull()

			if t := v.Get("trackingObject.trackerId.refId.optionType"); t.Exists() {
				va := v.Get("trackingObject.trackerId.refId.value")
				if t.String() == "global" {
					item.TrackerObjectId = types.StringValue(va.String())
				}
			}
			data.StaticNatSubnets = append(data.StaticNatSubnets, item)
			return true
		})
	} else {
		data.StaticNatSubnets = nil
	}
	if !fullRead {
		resultStaticNatSubnets := make([]ServiceLANVPNStaticNatSubnets, 0, len(data.StaticNatSubnets))
		matchedStaticNatSubnets := make([]bool, len(data.StaticNatSubnets))
		for _, oldItem := range oldStaticNatSubnets {
			for ni := range data.StaticNatSubnets {
				if matchedStaticNatSubnets[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.SourceIpSubnetVariable.ValueString() != "" || data.StaticNatSubnets[ni].SourceIpSubnetVariable.ValueString() != "") {
					if oldItem.SourceIpSubnetVariable.ValueString() != data.StaticNatSubnets[ni].SourceIpSubnetVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.SourceIpSubnet.ValueString() != data.StaticNatSubnets[ni].SourceIpSubnet.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedStaticNatSubnets[ni] = true
					resultStaticNatSubnets = append(resultStaticNatSubnets, data.StaticNatSubnets[ni])
					break
				}
			}
		}
		for ni := range data.StaticNatSubnets {
			if !matchedStaticNatSubnets[ni] {
				resultStaticNatSubnets = append(resultStaticNatSubnets, data.StaticNatSubnets[ni])
			}
		}
		data.StaticNatSubnets = resultStaticNatSubnets
	}
	oldNat64V4Pools := data.Nat64V4Pools
	if value := res.Get(path + "nat64V4Pool"); value.Exists() && len(value.Array()) > 0 {
		data.Nat64V4Pools = make([]ServiceLANVPNNat64V4Pools, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNNat64V4Pools{}
			item.Name = types.StringNull()
			item.NameVariable = types.StringNull()
			if t := v.Get("nat64V4PoolName.optionType"); t.Exists() {
				va := v.Get("nat64V4PoolName.value")
				if t.String() == "variable" {
					item.NameVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			item.RangeStart = types.StringNull()
			item.RangeStartVariable = types.StringNull()
			if t := v.Get("nat64V4PoolRangeStart.optionType"); t.Exists() {
				va := v.Get("nat64V4PoolRangeStart.value")
				if t.String() == "variable" {
					item.RangeStartVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RangeStart = types.StringValue(va.String())
				}
			}
			item.RangeEnd = types.StringNull()
			item.RangeEndVariable = types.StringNull()
			if t := v.Get("nat64V4PoolRangeEnd.optionType"); t.Exists() {
				va := v.Get("nat64V4PoolRangeEnd.value")
				if t.String() == "variable" {
					item.RangeEndVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RangeEnd = types.StringValue(va.String())
				}
			}
			item.Overload = types.BoolNull()
			item.OverloadVariable = types.StringNull()
			if t := v.Get("nat64V4PoolOverload.optionType"); t.Exists() {
				va := v.Get("nat64V4PoolOverload.value")
				if t.String() == "variable" {
					item.OverloadVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Overload = types.BoolValue(va.Bool())
				}
			}
			data.Nat64V4Pools = append(data.Nat64V4Pools, item)
			return true
		})
	} else {
		data.Nat64V4Pools = nil
	}
	if !fullRead {
		resultNat64V4Pools := make([]ServiceLANVPNNat64V4Pools, 0, len(data.Nat64V4Pools))
		matchedNat64V4Pools := make([]bool, len(data.Nat64V4Pools))
		for _, oldItem := range oldNat64V4Pools {
			for ni := range data.Nat64V4Pools {
				if matchedNat64V4Pools[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.NameVariable.ValueString() != "" || data.Nat64V4Pools[ni].NameVariable.ValueString() != "") {
					if oldItem.NameVariable.ValueString() != data.Nat64V4Pools[ni].NameVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.Name.ValueString() != data.Nat64V4Pools[ni].Name.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedNat64V4Pools[ni] = true
					resultNat64V4Pools = append(resultNat64V4Pools, data.Nat64V4Pools[ni])
					break
				}
			}
		}
		for ni := range data.Nat64V4Pools {
			if !matchedNat64V4Pools[ni] {
				resultNat64V4Pools = append(resultNat64V4Pools, data.Nat64V4Pools[ni])
			}
		}
		data.Nat64V4Pools = resultNat64V4Pools
	}
	oldRouteLeakFromGlobalVpns := data.RouteLeakFromGlobalVpns
	if value := res.Get(path + "routeLeakFromGlobal"); value.Exists() && len(value.Array()) > 0 {
		data.RouteLeakFromGlobalVpns = make([]ServiceLANVPNRouteLeakFromGlobalVpns, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNRouteLeakFromGlobalVpns{}
			item.RouteProtocol = types.StringNull()
			item.RouteProtocolVariable = types.StringNull()
			if t := v.Get("routeProtocol.optionType"); t.Exists() {
				va := v.Get("routeProtocol.value")
				if t.String() == "variable" {
					item.RouteProtocolVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RouteProtocol = types.StringValue(va.String())
				}
			}
			item.RoutePolicyId = types.StringNull()

			if t := v.Get("routePolicy.refId.optionType"); t.Exists() {
				va := v.Get("routePolicy.refId.value")
				if t.String() == "global" {
					item.RoutePolicyId = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("redistributeToProtocol"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Redistributions = make([]ServiceLANVPNRouteLeakFromGlobalVpnsRedistributions, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceLANVPNRouteLeakFromGlobalVpnsRedistributions{}
					cItem.Protocol = types.StringNull()
					cItem.ProtocolVariable = types.StringNull()
					if t := cv.Get("protocol.optionType"); t.Exists() {
						va := cv.Get("protocol.value")
						if t.String() == "variable" {
							cItem.ProtocolVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Protocol = types.StringValue(va.String())
						}
					}
					cItem.RedistributionPolicyId = types.StringNull()

					if t := cv.Get("policy.refId.optionType"); t.Exists() {
						va := cv.Get("policy.refId.value")
						if t.String() == "global" {
							cItem.RedistributionPolicyId = types.StringValue(va.String())
						}
					}
					item.Redistributions = append(item.Redistributions, cItem)
					return true
				})
			}
			data.RouteLeakFromGlobalVpns = append(data.RouteLeakFromGlobalVpns, item)
			return true
		})
	} else {
		data.RouteLeakFromGlobalVpns = nil
	}
	if !fullRead {
		resultRouteLeakFromGlobalVpns := make([]ServiceLANVPNRouteLeakFromGlobalVpns, 0, len(data.RouteLeakFromGlobalVpns))
		matchedRouteLeakFromGlobalVpns := make([]bool, len(data.RouteLeakFromGlobalVpns))
		for _, oldItem := range oldRouteLeakFromGlobalVpns {
			for ni := range data.RouteLeakFromGlobalVpns {
				if matchedRouteLeakFromGlobalVpns[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.RouteProtocolVariable.ValueString() != "" || data.RouteLeakFromGlobalVpns[ni].RouteProtocolVariable.ValueString() != "") {
					if oldItem.RouteProtocolVariable.ValueString() != data.RouteLeakFromGlobalVpns[ni].RouteProtocolVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.RouteProtocol.ValueString() != data.RouteLeakFromGlobalVpns[ni].RouteProtocol.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedRouteLeakFromGlobalVpns[ni] = true
					{
						resultC := make([]ServiceLANVPNRouteLeakFromGlobalVpnsRedistributions, 0, len(data.RouteLeakFromGlobalVpns[ni].Redistributions))
						matchedC := make([]bool, len(data.RouteLeakFromGlobalVpns[ni].Redistributions))
						for _, oldCItem := range oldItem.Redistributions {
							for nci := range data.RouteLeakFromGlobalVpns[ni].Redistributions {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC && (oldCItem.ProtocolVariable.ValueString() != "" || data.RouteLeakFromGlobalVpns[ni].Redistributions[nci].ProtocolVariable.ValueString() != "") {
									if oldCItem.ProtocolVariable.ValueString() != data.RouteLeakFromGlobalVpns[ni].Redistributions[nci].ProtocolVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.Protocol.ValueString() != data.RouteLeakFromGlobalVpns[ni].Redistributions[nci].Protocol.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.RouteLeakFromGlobalVpns[ni].Redistributions[nci])
									break
								}
							}
						}
						for nci := range data.RouteLeakFromGlobalVpns[ni].Redistributions {
							if !matchedC[nci] {
								resultC = append(resultC, data.RouteLeakFromGlobalVpns[ni].Redistributions[nci])
							}
						}
						data.RouteLeakFromGlobalVpns[ni].Redistributions = resultC
					}
					resultRouteLeakFromGlobalVpns = append(resultRouteLeakFromGlobalVpns, data.RouteLeakFromGlobalVpns[ni])
					break
				}
			}
		}
		for ni := range data.RouteLeakFromGlobalVpns {
			if !matchedRouteLeakFromGlobalVpns[ni] {
				resultRouteLeakFromGlobalVpns = append(resultRouteLeakFromGlobalVpns, data.RouteLeakFromGlobalVpns[ni])
			}
		}
		data.RouteLeakFromGlobalVpns = resultRouteLeakFromGlobalVpns
	}
	oldRouteLeakToGlobalVpns := data.RouteLeakToGlobalVpns
	if value := res.Get(path + "routeLeakFromService"); value.Exists() && len(value.Array()) > 0 {
		data.RouteLeakToGlobalVpns = make([]ServiceLANVPNRouteLeakToGlobalVpns, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNRouteLeakToGlobalVpns{}
			item.RouteProtocol = types.StringNull()
			item.RouteProtocolVariable = types.StringNull()
			if t := v.Get("routeProtocol.optionType"); t.Exists() {
				va := v.Get("routeProtocol.value")
				if t.String() == "variable" {
					item.RouteProtocolVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RouteProtocol = types.StringValue(va.String())
				}
			}
			item.RoutePolicyId = types.StringNull()

			if t := v.Get("routePolicy.refId.optionType"); t.Exists() {
				va := v.Get("routePolicy.refId.value")
				if t.String() == "global" {
					item.RoutePolicyId = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("redistributeToProtocol"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Redistributions = make([]ServiceLANVPNRouteLeakToGlobalVpnsRedistributions, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceLANVPNRouteLeakToGlobalVpnsRedistributions{}
					cItem.Protocol = types.StringNull()
					cItem.ProtocolVariable = types.StringNull()
					if t := cv.Get("protocol.optionType"); t.Exists() {
						va := cv.Get("protocol.value")
						if t.String() == "variable" {
							cItem.ProtocolVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Protocol = types.StringValue(va.String())
						}
					}
					cItem.RedistributionPolicyId = types.StringNull()

					if t := cv.Get("policy.refId.optionType"); t.Exists() {
						va := cv.Get("policy.refId.value")
						if t.String() == "global" {
							cItem.RedistributionPolicyId = types.StringValue(va.String())
						}
					}
					item.Redistributions = append(item.Redistributions, cItem)
					return true
				})
			}
			data.RouteLeakToGlobalVpns = append(data.RouteLeakToGlobalVpns, item)
			return true
		})
	} else {
		data.RouteLeakToGlobalVpns = nil
	}
	if !fullRead {
		resultRouteLeakToGlobalVpns := make([]ServiceLANVPNRouteLeakToGlobalVpns, 0, len(data.RouteLeakToGlobalVpns))
		matchedRouteLeakToGlobalVpns := make([]bool, len(data.RouteLeakToGlobalVpns))
		for _, oldItem := range oldRouteLeakToGlobalVpns {
			for ni := range data.RouteLeakToGlobalVpns {
				if matchedRouteLeakToGlobalVpns[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.RouteProtocolVariable.ValueString() != "" || data.RouteLeakToGlobalVpns[ni].RouteProtocolVariable.ValueString() != "") {
					if oldItem.RouteProtocolVariable.ValueString() != data.RouteLeakToGlobalVpns[ni].RouteProtocolVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.RouteProtocol.ValueString() != data.RouteLeakToGlobalVpns[ni].RouteProtocol.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedRouteLeakToGlobalVpns[ni] = true
					{
						resultC := make([]ServiceLANVPNRouteLeakToGlobalVpnsRedistributions, 0, len(data.RouteLeakToGlobalVpns[ni].Redistributions))
						matchedC := make([]bool, len(data.RouteLeakToGlobalVpns[ni].Redistributions))
						for _, oldCItem := range oldItem.Redistributions {
							for nci := range data.RouteLeakToGlobalVpns[ni].Redistributions {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC && (oldCItem.ProtocolVariable.ValueString() != "" || data.RouteLeakToGlobalVpns[ni].Redistributions[nci].ProtocolVariable.ValueString() != "") {
									if oldCItem.ProtocolVariable.ValueString() != data.RouteLeakToGlobalVpns[ni].Redistributions[nci].ProtocolVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.Protocol.ValueString() != data.RouteLeakToGlobalVpns[ni].Redistributions[nci].Protocol.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.RouteLeakToGlobalVpns[ni].Redistributions[nci])
									break
								}
							}
						}
						for nci := range data.RouteLeakToGlobalVpns[ni].Redistributions {
							if !matchedC[nci] {
								resultC = append(resultC, data.RouteLeakToGlobalVpns[ni].Redistributions[nci])
							}
						}
						data.RouteLeakToGlobalVpns[ni].Redistributions = resultC
					}
					resultRouteLeakToGlobalVpns = append(resultRouteLeakToGlobalVpns, data.RouteLeakToGlobalVpns[ni])
					break
				}
			}
		}
		for ni := range data.RouteLeakToGlobalVpns {
			if !matchedRouteLeakToGlobalVpns[ni] {
				resultRouteLeakToGlobalVpns = append(resultRouteLeakToGlobalVpns, data.RouteLeakToGlobalVpns[ni])
			}
		}
		data.RouteLeakToGlobalVpns = resultRouteLeakToGlobalVpns
	}
	oldRouteLeakFromOtherServices := data.RouteLeakFromOtherServices
	if value := res.Get(path + "routeLeakBetweenServices"); value.Exists() && len(value.Array()) > 0 {
		data.RouteLeakFromOtherServices = make([]ServiceLANVPNRouteLeakFromOtherServices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNRouteLeakFromOtherServices{}
			item.SourceVpn = types.Int64Null()
			item.SourceVpnVariable = types.StringNull()
			if t := v.Get("sourceVpn.optionType"); t.Exists() {
				va := v.Get("sourceVpn.value")
				if t.String() == "variable" {
					item.SourceVpnVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SourceVpn = types.Int64Value(va.Int())
				}
			}
			item.RouteProtocol = types.StringNull()
			item.RouteProtocolVariable = types.StringNull()
			if t := v.Get("routeProtocol.optionType"); t.Exists() {
				va := v.Get("routeProtocol.value")
				if t.String() == "variable" {
					item.RouteProtocolVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RouteProtocol = types.StringValue(va.String())
				}
			}
			item.RoutePolicyId = types.StringNull()

			if t := v.Get("routePolicy.refId.optionType"); t.Exists() {
				va := v.Get("routePolicy.refId.value")
				if t.String() == "global" {
					item.RoutePolicyId = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("redistributeToProtocol"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Redistributions = make([]ServiceLANVPNRouteLeakFromOtherServicesRedistributions, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceLANVPNRouteLeakFromOtherServicesRedistributions{}
					cItem.Protocol = types.StringNull()
					cItem.ProtocolVariable = types.StringNull()
					if t := cv.Get("protocol.optionType"); t.Exists() {
						va := cv.Get("protocol.value")
						if t.String() == "variable" {
							cItem.ProtocolVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Protocol = types.StringValue(va.String())
						}
					}
					cItem.RedistributionPolicyId = types.StringNull()

					if t := cv.Get("policy.refId.optionType"); t.Exists() {
						va := cv.Get("policy.refId.value")
						if t.String() == "global" {
							cItem.RedistributionPolicyId = types.StringValue(va.String())
						}
					}
					item.Redistributions = append(item.Redistributions, cItem)
					return true
				})
			}
			data.RouteLeakFromOtherServices = append(data.RouteLeakFromOtherServices, item)
			return true
		})
	} else {
		data.RouteLeakFromOtherServices = nil
	}
	if !fullRead {
		resultRouteLeakFromOtherServices := make([]ServiceLANVPNRouteLeakFromOtherServices, 0, len(data.RouteLeakFromOtherServices))
		matchedRouteLeakFromOtherServices := make([]bool, len(data.RouteLeakFromOtherServices))
		for _, oldItem := range oldRouteLeakFromOtherServices {
			for ni := range data.RouteLeakFromOtherServices {
				if matchedRouteLeakFromOtherServices[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.RouteProtocolVariable.ValueString() != "" || data.RouteLeakFromOtherServices[ni].RouteProtocolVariable.ValueString() != "") {
					if oldItem.RouteProtocolVariable.ValueString() != data.RouteLeakFromOtherServices[ni].RouteProtocolVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.RouteProtocol.ValueString() != data.RouteLeakFromOtherServices[ni].RouteProtocol.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedRouteLeakFromOtherServices[ni] = true
					{
						resultC := make([]ServiceLANVPNRouteLeakFromOtherServicesRedistributions, 0, len(data.RouteLeakFromOtherServices[ni].Redistributions))
						matchedC := make([]bool, len(data.RouteLeakFromOtherServices[ni].Redistributions))
						for _, oldCItem := range oldItem.Redistributions {
							for nci := range data.RouteLeakFromOtherServices[ni].Redistributions {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								if keyMatchC && (oldCItem.ProtocolVariable.ValueString() != "" || data.RouteLeakFromOtherServices[ni].Redistributions[nci].ProtocolVariable.ValueString() != "") {
									if oldCItem.ProtocolVariable.ValueString() != data.RouteLeakFromOtherServices[ni].Redistributions[nci].ProtocolVariable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
									if oldCItem.Protocol.ValueString() != data.RouteLeakFromOtherServices[ni].Redistributions[nci].Protocol.ValueString() {
										keyMatchC = false
									}
								}
								if keyMatchC {
									matchedC[nci] = true
									resultC = append(resultC, data.RouteLeakFromOtherServices[ni].Redistributions[nci])
									break
								}
							}
						}
						for nci := range data.RouteLeakFromOtherServices[ni].Redistributions {
							if !matchedC[nci] {
								resultC = append(resultC, data.RouteLeakFromOtherServices[ni].Redistributions[nci])
							}
						}
						data.RouteLeakFromOtherServices[ni].Redistributions = resultC
					}
					resultRouteLeakFromOtherServices = append(resultRouteLeakFromOtherServices, data.RouteLeakFromOtherServices[ni])
					break
				}
			}
		}
		for ni := range data.RouteLeakFromOtherServices {
			if !matchedRouteLeakFromOtherServices[ni] {
				resultRouteLeakFromOtherServices = append(resultRouteLeakFromOtherServices, data.RouteLeakFromOtherServices[ni])
			}
		}
		data.RouteLeakFromOtherServices = resultRouteLeakFromOtherServices
	}
	oldIpv4ImportRouteTargets := data.Ipv4ImportRouteTargets
	if value := res.Get(path + "mplsVpnIpv4RouteTarget.importRtList"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv4ImportRouteTargets = make([]ServiceLANVPNIpv4ImportRouteTargets, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNIpv4ImportRouteTargets{}
			item.RouteTarget = types.StringNull()
			item.RouteTargetVariable = types.StringNull()
			if t := v.Get("rt.optionType"); t.Exists() {
				va := v.Get("rt.value")
				if t.String() == "variable" {
					item.RouteTargetVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RouteTarget = types.StringValue(va.String())
				}
			}
			data.Ipv4ImportRouteTargets = append(data.Ipv4ImportRouteTargets, item)
			return true
		})
	} else {
		data.Ipv4ImportRouteTargets = nil
	}
	if !fullRead {
		resultIpv4ImportRouteTargets := make([]ServiceLANVPNIpv4ImportRouteTargets, 0, len(data.Ipv4ImportRouteTargets))
		matchedIpv4ImportRouteTargets := make([]bool, len(data.Ipv4ImportRouteTargets))
		for _, oldItem := range oldIpv4ImportRouteTargets {
			for ni := range data.Ipv4ImportRouteTargets {
				if matchedIpv4ImportRouteTargets[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.RouteTargetVariable.ValueString() != "" || data.Ipv4ImportRouteTargets[ni].RouteTargetVariable.ValueString() != "") {
					if oldItem.RouteTargetVariable.ValueString() != data.Ipv4ImportRouteTargets[ni].RouteTargetVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.RouteTarget.ValueString() != data.Ipv4ImportRouteTargets[ni].RouteTarget.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedIpv4ImportRouteTargets[ni] = true
					resultIpv4ImportRouteTargets = append(resultIpv4ImportRouteTargets, data.Ipv4ImportRouteTargets[ni])
					break
				}
			}
		}
		for ni := range data.Ipv4ImportRouteTargets {
			if !matchedIpv4ImportRouteTargets[ni] {
				resultIpv4ImportRouteTargets = append(resultIpv4ImportRouteTargets, data.Ipv4ImportRouteTargets[ni])
			}
		}
		data.Ipv4ImportRouteTargets = resultIpv4ImportRouteTargets
	}
	oldIpv4ExportRouteTargets := data.Ipv4ExportRouteTargets
	if value := res.Get(path + "mplsVpnIpv4RouteTarget.exportRtList"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv4ExportRouteTargets = make([]ServiceLANVPNIpv4ExportRouteTargets, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNIpv4ExportRouteTargets{}
			item.RouteTarget = types.StringNull()
			item.RouteTargetVariable = types.StringNull()
			if t := v.Get("rt.optionType"); t.Exists() {
				va := v.Get("rt.value")
				if t.String() == "variable" {
					item.RouteTargetVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RouteTarget = types.StringValue(va.String())
				}
			}
			data.Ipv4ExportRouteTargets = append(data.Ipv4ExportRouteTargets, item)
			return true
		})
	} else {
		data.Ipv4ExportRouteTargets = nil
	}
	if !fullRead {
		resultIpv4ExportRouteTargets := make([]ServiceLANVPNIpv4ExportRouteTargets, 0, len(data.Ipv4ExportRouteTargets))
		matchedIpv4ExportRouteTargets := make([]bool, len(data.Ipv4ExportRouteTargets))
		for _, oldItem := range oldIpv4ExportRouteTargets {
			for ni := range data.Ipv4ExportRouteTargets {
				if matchedIpv4ExportRouteTargets[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.RouteTargetVariable.ValueString() != "" || data.Ipv4ExportRouteTargets[ni].RouteTargetVariable.ValueString() != "") {
					if oldItem.RouteTargetVariable.ValueString() != data.Ipv4ExportRouteTargets[ni].RouteTargetVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.RouteTarget.ValueString() != data.Ipv4ExportRouteTargets[ni].RouteTarget.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedIpv4ExportRouteTargets[ni] = true
					resultIpv4ExportRouteTargets = append(resultIpv4ExportRouteTargets, data.Ipv4ExportRouteTargets[ni])
					break
				}
			}
		}
		for ni := range data.Ipv4ExportRouteTargets {
			if !matchedIpv4ExportRouteTargets[ni] {
				resultIpv4ExportRouteTargets = append(resultIpv4ExportRouteTargets, data.Ipv4ExportRouteTargets[ni])
			}
		}
		data.Ipv4ExportRouteTargets = resultIpv4ExportRouteTargets
	}
	oldIpv6ImportRouteTargets := data.Ipv6ImportRouteTargets
	if value := res.Get(path + "mplsVpnIpv6RouteTarget.importRtList"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv6ImportRouteTargets = make([]ServiceLANVPNIpv6ImportRouteTargets, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNIpv6ImportRouteTargets{}
			item.RouteTarget = types.StringNull()
			item.RouteTargetVariable = types.StringNull()
			if t := v.Get("rt.optionType"); t.Exists() {
				va := v.Get("rt.value")
				if t.String() == "variable" {
					item.RouteTargetVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RouteTarget = types.StringValue(va.String())
				}
			}
			data.Ipv6ImportRouteTargets = append(data.Ipv6ImportRouteTargets, item)
			return true
		})
	} else {
		data.Ipv6ImportRouteTargets = nil
	}
	if !fullRead {
		resultIpv6ImportRouteTargets := make([]ServiceLANVPNIpv6ImportRouteTargets, 0, len(data.Ipv6ImportRouteTargets))
		matchedIpv6ImportRouteTargets := make([]bool, len(data.Ipv6ImportRouteTargets))
		for _, oldItem := range oldIpv6ImportRouteTargets {
			for ni := range data.Ipv6ImportRouteTargets {
				if matchedIpv6ImportRouteTargets[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.RouteTargetVariable.ValueString() != "" || data.Ipv6ImportRouteTargets[ni].RouteTargetVariable.ValueString() != "") {
					if oldItem.RouteTargetVariable.ValueString() != data.Ipv6ImportRouteTargets[ni].RouteTargetVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.RouteTarget.ValueString() != data.Ipv6ImportRouteTargets[ni].RouteTarget.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedIpv6ImportRouteTargets[ni] = true
					resultIpv6ImportRouteTargets = append(resultIpv6ImportRouteTargets, data.Ipv6ImportRouteTargets[ni])
					break
				}
			}
		}
		for ni := range data.Ipv6ImportRouteTargets {
			if !matchedIpv6ImportRouteTargets[ni] {
				resultIpv6ImportRouteTargets = append(resultIpv6ImportRouteTargets, data.Ipv6ImportRouteTargets[ni])
			}
		}
		data.Ipv6ImportRouteTargets = resultIpv6ImportRouteTargets
	}
	oldIpv6ExportRouteTargets := data.Ipv6ExportRouteTargets
	if value := res.Get(path + "mplsVpnIpv6RouteTarget.exportRtList"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv6ExportRouteTargets = make([]ServiceLANVPNIpv6ExportRouteTargets, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceLANVPNIpv6ExportRouteTargets{}
			item.RouteTarget = types.StringNull()
			item.RouteTargetVariable = types.StringNull()
			if t := v.Get("rt.optionType"); t.Exists() {
				va := v.Get("rt.value")
				if t.String() == "variable" {
					item.RouteTargetVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RouteTarget = types.StringValue(va.String())
				}
			}
			data.Ipv6ExportRouteTargets = append(data.Ipv6ExportRouteTargets, item)
			return true
		})
	} else {
		data.Ipv6ExportRouteTargets = nil
	}
	if !fullRead {
		resultIpv6ExportRouteTargets := make([]ServiceLANVPNIpv6ExportRouteTargets, 0, len(data.Ipv6ExportRouteTargets))
		matchedIpv6ExportRouteTargets := make([]bool, len(data.Ipv6ExportRouteTargets))
		for _, oldItem := range oldIpv6ExportRouteTargets {
			for ni := range data.Ipv6ExportRouteTargets {
				if matchedIpv6ExportRouteTargets[ni] {
					continue
				}
				keyMatch := true
				if keyMatch && (oldItem.RouteTargetVariable.ValueString() != "" || data.Ipv6ExportRouteTargets[ni].RouteTargetVariable.ValueString() != "") {
					if oldItem.RouteTargetVariable.ValueString() != data.Ipv6ExportRouteTargets[ni].RouteTargetVariable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
					if oldItem.RouteTarget.ValueString() != data.Ipv6ExportRouteTargets[ni].RouteTarget.ValueString() {
						keyMatch = false
					}
				}
				if keyMatch {
					matchedIpv6ExportRouteTargets[ni] = true
					resultIpv6ExportRouteTargets = append(resultIpv6ExportRouteTargets, data.Ipv6ExportRouteTargets[ni])
					break
				}
			}
		}
		for ni := range data.Ipv6ExportRouteTargets {
			if !matchedIpv6ExportRouteTargets[ni] {
				resultIpv6ExportRouteTargets = append(resultIpv6ExportRouteTargets, data.Ipv6ExportRouteTargets[ni])
			}
		}
		data.Ipv6ExportRouteTargets = resultIpv6ExportRouteTargets
	}
}

// End of section. //template:end fromBody
