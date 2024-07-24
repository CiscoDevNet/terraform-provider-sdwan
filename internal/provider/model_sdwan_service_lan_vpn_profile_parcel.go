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
	NetworkAddress         types.String                                       `tfsdk:"network_address"`
	NetworkAddressVariable types.String                                       `tfsdk:"network_address_variable"`
	SubnetMask             types.String                                       `tfsdk:"subnet_mask"`
	SubnetMaskVariable     types.String                                       `tfsdk:"subnet_mask_variable"`
	NextHops               []ServiceLANVPNIpv4StaticRoutesNextHops            `tfsdk:"next_hops"`
	NextHopWithTrackers    []ServiceLANVPNIpv4StaticRoutesNextHopWithTrackers `tfsdk:"next_hop_with_trackers"`
	Null0                  types.Bool                                         `tfsdk:"null0"`
	GatewayDhcp            types.Bool                                         `tfsdk:"gateway_dhcp"`
	Vpn                    types.Bool                                         `tfsdk:"vpn"`
}

type ServiceLANVPNIpv6StaticRoutes struct {
	Prefix         types.String                            `tfsdk:"prefix"`
	PrefixVariable types.String                            `tfsdk:"prefix_variable"`
	NextHops       []ServiceLANVPNIpv6StaticRoutesNextHops `tfsdk:"next_hops"`
	Null0          types.Bool                              `tfsdk:"null0"`
	Nat            types.String                            `tfsdk:"nat"`
	NatVariable    types.String                            `tfsdk:"nat_variable"`
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

type ServiceLANVPNIpv6StaticRoutesNextHops struct {
	Address                        types.String `tfsdk:"address"`
	AddressVariable                types.String `tfsdk:"address_variable"`
	AdministrativeDistance         types.Int64  `tfsdk:"administrative_distance"`
	AdministrativeDistanceVariable types.String `tfsdk:"administrative_distance_variable"`
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
		body, _ = sjson.Set(body, path+"vpnId.optionType", "global")
		body, _ = sjson.Set(body, path+"vpnId.value", data.Vpn.ValueInt64())
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
		body, _ = sjson.Set(body, path+"name.optionType", "global")
		body, _ = sjson.Set(body, path+"name.value", data.ConfigDescription.ValueString())
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
		body, _ = sjson.Set(body, path+"ompAdminDistance.optionType", "global")
		body, _ = sjson.Set(body, path+"ompAdminDistance.value", data.OmpAdminDistanceIpv4.ValueInt64())
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
		body, _ = sjson.Set(body, path+"ompAdminDistanceIpv6.optionType", "global")
		body, _ = sjson.Set(body, path+"ompAdminDistanceIpv6.value", data.OmpAdminDistanceIpv6.ValueInt64())
	}
	if data.EnableSdwanRemoteAccess.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"enableSdra.optionType", "default")
			body, _ = sjson.Set(body, path+"enableSdra.value", false)
		}
	} else {
		body, _ = sjson.Set(body, path+"enableSdra.optionType", "global")
		body, _ = sjson.Set(body, path+"enableSdra.value", data.EnableSdwanRemoteAccess.ValueBool())
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
		body, _ = sjson.Set(body, path+"dnsIpv4.primaryDnsAddressIpv4.optionType", "global")
		body, _ = sjson.Set(body, path+"dnsIpv4.primaryDnsAddressIpv4.value", data.PrimaryDnsAddressIpv4.ValueString())
	}

	if !data.SecondaryDnsAddressIpv4Variable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dnsIpv4.secondaryDnsAddressIpv4.optionType", "variable")
			body, _ = sjson.Set(body, path+"dnsIpv4.secondaryDnsAddressIpv4.value", data.SecondaryDnsAddressIpv4Variable.ValueString())
		}
	} else if !data.SecondaryDnsAddressIpv4.IsNull() {
		body, _ = sjson.Set(body, path+"dnsIpv4.secondaryDnsAddressIpv4.optionType", "global")
		body, _ = sjson.Set(body, path+"dnsIpv4.secondaryDnsAddressIpv4.value", data.SecondaryDnsAddressIpv4.ValueString())
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
		body, _ = sjson.Set(body, path+"dnsIpv6.primaryDnsAddressIpv6.optionType", "global")
		body, _ = sjson.Set(body, path+"dnsIpv6.primaryDnsAddressIpv6.value", data.PrimaryDnsAddressIpv6.ValueString())
	}

	if !data.SecondaryDnsAddressIpv6Variable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"dnsIpv6.secondaryDnsAddressIpv6.optionType", "variable")
			body, _ = sjson.Set(body, path+"dnsIpv6.secondaryDnsAddressIpv6.value", data.SecondaryDnsAddressIpv6Variable.ValueString())
		}
	} else if !data.SecondaryDnsAddressIpv6.IsNull() {
		body, _ = sjson.Set(body, path+"dnsIpv6.secondaryDnsAddressIpv6.optionType", "global")
		body, _ = sjson.Set(body, path+"dnsIpv6.secondaryDnsAddressIpv6.value", data.SecondaryDnsAddressIpv6.ValueString())
	}
	body, _ = sjson.Set(body, path+"newHostMapping", []interface{}{})
	for _, item := range data.HostMappings {
		itemBody := ""

		if !item.HostNameVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "hostName.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "hostName.value", item.HostNameVariable.ValueString())
			}
		} else if !item.HostName.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "hostName.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "hostName.value", item.HostName.ValueString())
		}

		if !item.ListOfIpsVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "listOfIp.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "listOfIp.value", item.ListOfIpsVariable.ValueString())
			}
		} else if !item.ListOfIps.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "listOfIp.optionType", "global")
			var values []string
			item.ListOfIps.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "listOfIp.value", values)
		}
		body, _ = sjson.SetRaw(body, path+"newHostMapping.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"ompAdvertiseIp4", []interface{}{})
	for _, item := range data.AdvertiseOmpIpv4s {
		itemBody := ""

		if !item.ProtocolVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "ompProtocol.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "ompProtocol.value", item.ProtocolVariable.ValueString())
			}
		} else if !item.Protocol.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ompProtocol.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "ompProtocol.value", item.Protocol.ValueString())
		}
		if !item.RoutePolicyId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.value", item.RoutePolicyId.ValueString())
		}

		for _, childItem := range item.Prefixes {
			itemChildBody := ""

			if !childItem.NetworkAddressVariable.IsNull() {
				if true {
					itemChildBody, _ = sjson.Set(itemChildBody, "prefix.address.optionType", "variable")
					itemChildBody, _ = sjson.Set(itemChildBody, "prefix.address.value", childItem.NetworkAddressVariable.ValueString())
				}
			} else if !childItem.NetworkAddress.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix.address.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix.address.value", childItem.NetworkAddress.ValueString())
			}

			if !childItem.SubnetMaskVariable.IsNull() {
				if true {
					itemChildBody, _ = sjson.Set(itemChildBody, "prefix.mask.optionType", "variable")
					itemChildBody, _ = sjson.Set(itemChildBody, "prefix.mask.value", childItem.SubnetMaskVariable.ValueString())
				}
			} else if !childItem.SubnetMask.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix.mask.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix.mask.value", childItem.SubnetMask.ValueString())
			}
			if childItem.AggregateOnly.IsNull() {
				if true {
					itemChildBody, _ = sjson.Set(itemChildBody, "aggregateOnly.optionType", "default")
					itemChildBody, _ = sjson.Set(itemChildBody, "aggregateOnly.value", false)
				}
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregateOnly.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregateOnly.value", childItem.AggregateOnly.ValueBool())
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
				itemChildBody, _ = sjson.Set(itemChildBody, "region.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "region.value", childItem.Region.ValueString())
			}
			itemBody, _ = sjson.SetRaw(itemBody, "prefixList.-1", itemChildBody)
		}
		body, _ = sjson.SetRaw(body, path+"ompAdvertiseIp4.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"ompAdvertiseIpv6", []interface{}{})
	for _, item := range data.AdvertiseOmpIpv6s {
		itemBody := ""

		if !item.ProtocolVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "ompProtocol.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "ompProtocol.value", item.ProtocolVariable.ValueString())
			}
		} else if !item.Protocol.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ompProtocol.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "ompProtocol.value", item.Protocol.ValueString())
		}
		if !item.RoutePolicyId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.value", item.RoutePolicyId.ValueString())
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
			itemBody, _ = sjson.Set(itemBody, "protocolSubType.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "protocolSubType.value", item.ProtocolSubType.ValueString())
		}

		for _, childItem := range item.Prefixes {
			itemChildBody := ""

			if !childItem.PrefixVariable.IsNull() {
				if true {
					itemChildBody, _ = sjson.Set(itemChildBody, "prefix.optionType", "variable")
					itemChildBody, _ = sjson.Set(itemChildBody, "prefix.value", childItem.PrefixVariable.ValueString())
				}
			} else if !childItem.Prefix.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix.value", childItem.Prefix.ValueString())
			}
			if childItem.AggregateOnly.IsNull() {
				if true {
					itemChildBody, _ = sjson.Set(itemChildBody, "aggregateOnly.optionType", "default")
					itemChildBody, _ = sjson.Set(itemChildBody, "aggregateOnly.value", false)
				}
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregateOnly.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "aggregateOnly.value", childItem.AggregateOnly.ValueBool())
			}
			itemBody, _ = sjson.SetRaw(itemBody, "prefixList.-1", itemChildBody)
		}
		body, _ = sjson.SetRaw(body, path+"ompAdvertiseIpv6.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"ipv4Route", []interface{}{})
	for _, item := range data.Ipv4StaticRoutes {
		itemBody := ""

		if !item.NetworkAddressVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.value", item.NetworkAddressVariable.ValueString())
			}
		} else if !item.NetworkAddress.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.value", item.NetworkAddress.ValueString())
		}

		if !item.SubnetMaskVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.value", item.SubnetMaskVariable.ValueString())
			}
		} else if !item.SubnetMask.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.value", item.SubnetMask.ValueString())
		}

		for _, childItem := range item.NextHops {
			itemChildBody := ""

			if !childItem.AddressVariable.IsNull() {
				if true {
					itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "variable")
					itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.AddressVariable.ValueString())
				}
			} else if !childItem.Address.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
			}

			if !childItem.AdministrativeDistanceVariable.IsNull() {
				if true {
					itemChildBody, _ = sjson.Set(itemChildBody, "distance.optionType", "variable")
					itemChildBody, _ = sjson.Set(itemChildBody, "distance.value", childItem.AdministrativeDistanceVariable.ValueString())
				}
			} else if !childItem.AdministrativeDistance.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "distance.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance.value", childItem.AdministrativeDistance.ValueInt64())
			}
			itemBody, _ = sjson.SetRaw(itemBody, "oneOfIpRoute.nextHopContainer.nextHop.-1", itemChildBody)
		}

		for _, childItem := range item.NextHopWithTrackers {
			itemChildBody := ""

			if !childItem.AddressVariable.IsNull() {
				if true {
					itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "variable")
					itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.AddressVariable.ValueString())
				}
			} else if !childItem.Address.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
			}

			if !childItem.AdministrativeDistanceVariable.IsNull() {
				if true {
					itemChildBody, _ = sjson.Set(itemChildBody, "distance.optionType", "variable")
					itemChildBody, _ = sjson.Set(itemChildBody, "distance.value", childItem.AdministrativeDistanceVariable.ValueString())
				}
			} else if !childItem.AdministrativeDistance.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "distance.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance.value", childItem.AdministrativeDistance.ValueInt64())
			}
			if !childItem.TrackerId.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "tracker.refId.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "tracker.refId.value", childItem.TrackerId.ValueString())
			}
			itemBody, _ = sjson.SetRaw(itemBody, "oneOfIpRoute.nextHopContainer.nextHopWithTracker.-1", itemChildBody)
		}
		if !item.Null0.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.null0.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.null0.value", item.Null0.ValueBool())
		}
		if !item.GatewayDhcp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.dhcp.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.dhcp.value", item.GatewayDhcp.ValueBool())
		}
		if !item.Vpn.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.vpn.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.vpn.value", item.Vpn.ValueBool())
		}
		body, _ = sjson.SetRaw(body, path+"ipv4Route.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"ipv6Route", []interface{}{})
	for _, item := range data.Ipv6StaticRoutes {
		itemBody := ""

		if !item.PrefixVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "prefix.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "prefix.value", item.PrefixVariable.ValueString())
			}
		} else if !item.Prefix.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefix.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "prefix.value", item.Prefix.ValueString())
		}

		for _, childItem := range item.NextHops {
			itemChildBody := ""

			if !childItem.AddressVariable.IsNull() {
				if true {
					itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "variable")
					itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.AddressVariable.ValueString())
				}
			} else if !childItem.Address.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "address.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "address.value", childItem.Address.ValueString())
			}

			if !childItem.AdministrativeDistanceVariable.IsNull() {
				if true {
					itemChildBody, _ = sjson.Set(itemChildBody, "distance.optionType", "variable")
					itemChildBody, _ = sjson.Set(itemChildBody, "distance.value", childItem.AdministrativeDistanceVariable.ValueString())
				}
			} else if !childItem.AdministrativeDistance.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "distance.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "distance.value", childItem.AdministrativeDistance.ValueInt64())
			}
			itemBody, _ = sjson.SetRaw(itemBody, "oneOfIpRoute.nextHopContainer.nextHop.-1", itemChildBody)
		}
		if !item.Null0.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.null0.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.null0.value", item.Null0.ValueBool())
		}

		if !item.NatVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.nat.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.nat.value", item.NatVariable.ValueString())
			}
		} else if !item.Nat.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.nat.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "oneOfIpRoute.nat.value", item.Nat.ValueString())
		}
		body, _ = sjson.SetRaw(body, path+"ipv6Route.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"service", []interface{}{})
	for _, item := range data.Services {
		itemBody := ""

		if !item.ServiceTypeVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "serviceType.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "serviceType.value", item.ServiceTypeVariable.ValueString())
			}
		} else if !item.ServiceType.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "serviceType.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "serviceType.value", item.ServiceType.ValueString())
		}

		if !item.Ipv4AddressesVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "ipv4Addresses.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "ipv4Addresses.value", item.Ipv4AddressesVariable.ValueString())
			}
		} else if !item.Ipv4Addresses.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ipv4Addresses.optionType", "global")
			var values []string
			item.Ipv4Addresses.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "ipv4Addresses.value", values)
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
			itemBody, _ = sjson.Set(itemBody, "tracking.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "tracking.value", item.Tracking.ValueBool())
		}
		body, _ = sjson.SetRaw(body, path+"service.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"serviceRoute", []interface{}{})
	for _, item := range data.ServiceRoutes {
		itemBody := ""

		if !item.NetworkAddressVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.value", item.NetworkAddressVariable.ValueString())
			}
		} else if !item.NetworkAddress.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.value", item.NetworkAddress.ValueString())
		}

		if !item.SubnetMaskVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.value", item.SubnetMaskVariable.ValueString())
			}
		} else if !item.SubnetMask.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.value", item.SubnetMask.ValueString())
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
			itemBody, _ = sjson.Set(itemBody, "service.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "service.value", item.Service.ValueString())
		}
		if !item.Vpn.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
		}
		body, _ = sjson.SetRaw(body, path+"serviceRoute.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"greRoute", []interface{}{})
	for _, item := range data.GreRoutes {
		itemBody := ""

		if !item.NetworkAddressVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.value", item.NetworkAddressVariable.ValueString())
			}
		} else if !item.NetworkAddress.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.value", item.NetworkAddress.ValueString())
		}

		if !item.SubnetMaskVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.value", item.SubnetMaskVariable.ValueString())
			}
		} else if !item.SubnetMask.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.value", item.SubnetMask.ValueString())
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
			itemBody, _ = sjson.Set(itemBody, "interface.optionType", "global")
			var values []string
			item.Interface.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "interface.value", values)
		}
		if !item.Vpn.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueInt64())
		}
		body, _ = sjson.SetRaw(body, path+"greRoute.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"ipsecRoute", []interface{}{})
	for _, item := range data.IpsecRoutes {
		itemBody := ""

		if !item.NetworkAddressVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.value", item.NetworkAddressVariable.ValueString())
			}
		} else if !item.NetworkAddress.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "prefix.ipAddress.value", item.NetworkAddress.ValueString())
		}

		if !item.SubnetMaskVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.value", item.SubnetMaskVariable.ValueString())
			}
		} else if !item.SubnetMask.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "prefix.subnetMask.value", item.SubnetMask.ValueString())
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
			itemBody, _ = sjson.Set(itemBody, "interface.optionType", "global")
			var values []string
			item.Interface.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "interface.value", values)
		}
		body, _ = sjson.SetRaw(body, path+"ipsecRoute.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"natPool", []interface{}{})
	for _, item := range data.NatPools {
		itemBody := ""

		if !item.NatPoolNameVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "natPoolName.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "natPoolName.value", item.NatPoolNameVariable.ValueString())
			}
		} else if !item.NatPoolName.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "natPoolName.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "natPoolName.value", item.NatPoolName.ValueInt64())
		}

		if !item.PrefixLengthVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "prefixLength.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "prefixLength.value", item.PrefixLengthVariable.ValueString())
			}
		} else if !item.PrefixLength.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefixLength.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "prefixLength.value", item.PrefixLength.ValueInt64())
		}

		if !item.RangeStartVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "rangeStart.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "rangeStart.value", item.RangeStartVariable.ValueString())
			}
		} else if !item.RangeStart.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "rangeStart.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "rangeStart.value", item.RangeStart.ValueString())
		}

		if !item.RangeEndVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "rangeEnd.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "rangeEnd.value", item.RangeEndVariable.ValueString())
			}
		} else if !item.RangeEnd.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "rangeEnd.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "rangeEnd.value", item.RangeEnd.ValueString())
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
			itemBody, _ = sjson.Set(itemBody, "overload.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "overload.value", item.Overload.ValueBool())
		}

		if !item.DirectionVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "direction.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "direction.value", item.DirectionVariable.ValueString())
			}
		} else if !item.Direction.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "direction.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "direction.value", item.Direction.ValueString())
		}
		if !item.TrackerObjectId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "trackingObject.trackerId.refId.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "trackingObject.trackerId.refId.value", item.TrackerObjectId.ValueString())
		}
		body, _ = sjson.SetRaw(body, path+"natPool.-1", itemBody)
	}
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
			itemBody, _ = sjson.Set(itemBody, "natPoolName.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "natPoolName.value", item.NatPoolName.ValueInt64())
		}

		if !item.SourcePortVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "sourcePort.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "sourcePort.value", item.SourcePortVariable.ValueString())
			}
		} else if !item.SourcePort.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sourcePort.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sourcePort.value", item.SourcePort.ValueInt64())
		}

		if !item.TranslatePortVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "translatePort.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "translatePort.value", item.TranslatePortVariable.ValueString())
			}
		} else if !item.TranslatePort.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "translatePort.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "translatePort.value", item.TranslatePort.ValueInt64())
		}

		if !item.SourceIpVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "sourceIp.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "sourceIp.value", item.SourceIpVariable.ValueString())
			}
		} else if !item.SourceIp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sourceIp.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sourceIp.value", item.SourceIp.ValueString())
		}

		if !item.TranslatedSourceIpVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "TranslatedSourceIp.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "TranslatedSourceIp.value", item.TranslatedSourceIpVariable.ValueString())
			}
		} else if !item.TranslatedSourceIp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "TranslatedSourceIp.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "TranslatedSourceIp.value", item.TranslatedSourceIp.ValueString())
		}

		if !item.ProtocolVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "protocol.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "protocol.value", item.ProtocolVariable.ValueString())
			}
		} else if !item.Protocol.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "protocol.value", item.Protocol.ValueString())
		}
		body, _ = sjson.SetRaw(body, path+"natPortForward.-1", itemBody)
	}
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
			itemBody, _ = sjson.Set(itemBody, "natPoolName.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "natPoolName.value", item.NatPoolName.ValueInt64())
		}

		if !item.SourceIpVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "sourceIp.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "sourceIp.value", item.SourceIpVariable.ValueString())
			}
		} else if !item.SourceIp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sourceIp.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sourceIp.value", item.SourceIp.ValueString())
		}

		if !item.TranslatedSourceIpVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "TranslatedSourceIp.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "TranslatedSourceIp.value", item.TranslatedSourceIpVariable.ValueString())
			}
		} else if !item.TranslatedSourceIp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "TranslatedSourceIp.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "TranslatedSourceIp.value", item.TranslatedSourceIp.ValueString())
		}

		if !item.StaticNatDirectionVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "staticNatDirection.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "staticNatDirection.value", item.StaticNatDirectionVariable.ValueString())
			}
		} else if !item.StaticNatDirection.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "staticNatDirection.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "staticNatDirection.value", item.StaticNatDirection.ValueString())
		}
		if !item.TrackerObjectId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "trackingObject.trackerId.refId.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "trackingObject.trackerId.refId.value", item.TrackerObjectId.ValueString())
		}
		body, _ = sjson.SetRaw(body, path+"staticNat.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"nat64V4Pool", []interface{}{})
	for _, item := range data.Nat64V4Pools {
		itemBody := ""

		if !item.NameVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "nat64V4PoolName.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "nat64V4PoolName.value", item.NameVariable.ValueString())
			}
		} else if !item.Name.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "nat64V4PoolName.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "nat64V4PoolName.value", item.Name.ValueString())
		}

		if !item.RangeStartVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeStart.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeStart.value", item.RangeStartVariable.ValueString())
			}
		} else if !item.RangeStart.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeStart.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeStart.value", item.RangeStart.ValueString())
		}

		if !item.RangeEndVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeEnd.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeEnd.value", item.RangeEndVariable.ValueString())
			}
		} else if !item.RangeEnd.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeEnd.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "nat64V4PoolRangeEnd.value", item.RangeEnd.ValueString())
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
			itemBody, _ = sjson.Set(itemBody, "nat64V4PoolOverload.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "nat64V4PoolOverload.value", item.Overload.ValueBool())
		}
		body, _ = sjson.SetRaw(body, path+"nat64V4Pool.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"routeLeakFromGlobal", []interface{}{})
	for _, item := range data.RouteLeakFromGlobalVpns {
		itemBody := ""

		if !item.RouteProtocolVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "routeProtocol.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "routeProtocol.value", item.RouteProtocolVariable.ValueString())
			}
		} else if !item.RouteProtocol.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "routeProtocol.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "routeProtocol.value", item.RouteProtocol.ValueString())
		}
		if !item.RoutePolicyId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.value", item.RoutePolicyId.ValueString())
		}
		itemBody, _ = sjson.Set(itemBody, "redistributeToProtocol", []interface{}{})
		for _, childItem := range item.Redistributions {
			itemChildBody := ""

			if !childItem.ProtocolVariable.IsNull() {
				if true {
					itemChildBody, _ = sjson.Set(itemChildBody, "protocol.optionType", "variable")
					itemChildBody, _ = sjson.Set(itemChildBody, "protocol.value", childItem.ProtocolVariable.ValueString())
				}
			} else if !childItem.Protocol.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol.value", childItem.Protocol.ValueString())
			}
			if !childItem.RedistributionPolicyId.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "policy.refId.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "policy.refId.value", childItem.RedistributionPolicyId.ValueString())
			}
			itemBody, _ = sjson.SetRaw(itemBody, "redistributeToProtocol.-1", itemChildBody)
		}
		body, _ = sjson.SetRaw(body, path+"routeLeakFromGlobal.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"routeLeakFromService", []interface{}{})
	for _, item := range data.RouteLeakToGlobalVpns {
		itemBody := ""

		if !item.RouteProtocolVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "routeProtocol.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "routeProtocol.value", item.RouteProtocolVariable.ValueString())
			}
		} else if !item.RouteProtocol.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "routeProtocol.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "routeProtocol.value", item.RouteProtocol.ValueString())
		}
		if !item.RoutePolicyId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.value", item.RoutePolicyId.ValueString())
		}
		itemBody, _ = sjson.Set(itemBody, "redistributeToProtocol", []interface{}{})
		for _, childItem := range item.Redistributions {
			itemChildBody := ""

			if !childItem.ProtocolVariable.IsNull() {
				if true {
					itemChildBody, _ = sjson.Set(itemChildBody, "protocol.optionType", "variable")
					itemChildBody, _ = sjson.Set(itemChildBody, "protocol.value", childItem.ProtocolVariable.ValueString())
				}
			} else if !childItem.Protocol.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol.value", childItem.Protocol.ValueString())
			}
			if !childItem.RedistributionPolicyId.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "policy.refId.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "policy.refId.value", childItem.RedistributionPolicyId.ValueString())
			}
			itemBody, _ = sjson.SetRaw(itemBody, "redistributeToProtocol.-1", itemChildBody)
		}
		body, _ = sjson.SetRaw(body, path+"routeLeakFromService.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"routeLeakBetweenServices", []interface{}{})
	for _, item := range data.RouteLeakFromOtherServices {
		itemBody := ""

		if !item.SourceVpnVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "sourceVpn.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "sourceVpn.value", item.SourceVpnVariable.ValueString())
			}
		} else if !item.SourceVpn.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sourceVpn.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "sourceVpn.value", item.SourceVpn.ValueInt64())
		}

		if !item.RouteProtocolVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "routeProtocol.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "routeProtocol.value", item.RouteProtocolVariable.ValueString())
			}
		} else if !item.RouteProtocol.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "routeProtocol.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "routeProtocol.value", item.RouteProtocol.ValueString())
		}
		if !item.RoutePolicyId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.value", item.RoutePolicyId.ValueString())
		}
		itemBody, _ = sjson.Set(itemBody, "redistributeToProtocol", []interface{}{})
		for _, childItem := range item.Redistributions {
			itemChildBody := ""

			if !childItem.ProtocolVariable.IsNull() {
				if true {
					itemChildBody, _ = sjson.Set(itemChildBody, "protocol.optionType", "variable")
					itemChildBody, _ = sjson.Set(itemChildBody, "protocol.value", childItem.ProtocolVariable.ValueString())
				}
			} else if !childItem.Protocol.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol.value", childItem.Protocol.ValueString())
			}
			if !childItem.RedistributionPolicyId.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "policy.refId.optionType", "global")
				itemChildBody, _ = sjson.Set(itemChildBody, "policy.refId.value", childItem.RedistributionPolicyId.ValueString())
			}
			itemBody, _ = sjson.SetRaw(itemBody, "redistributeToProtocol.-1", itemChildBody)
		}
		body, _ = sjson.SetRaw(body, path+"routeLeakBetweenServices.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"mplsVpnIpv4RouteTarget.importRtList", []interface{}{})
	for _, item := range data.Ipv4ImportRouteTargets {
		itemBody := ""

		if !item.RouteTargetVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "rt.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "rt.value", item.RouteTargetVariable.ValueString())
			}
		} else if !item.RouteTarget.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "rt.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "rt.value", item.RouteTarget.ValueString())
		}
		body, _ = sjson.SetRaw(body, path+"mplsVpnIpv4RouteTarget.importRtList.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"mplsVpnIpv4RouteTarget.exportRtList", []interface{}{})
	for _, item := range data.Ipv4ExportRouteTargets {
		itemBody := ""

		if !item.RouteTargetVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "rt.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "rt.value", item.RouteTargetVariable.ValueString())
			}
		} else if !item.RouteTarget.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "rt.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "rt.value", item.RouteTarget.ValueString())
		}
		body, _ = sjson.SetRaw(body, path+"mplsVpnIpv4RouteTarget.exportRtList.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"mplsVpnIpv6RouteTarget.importRtList", []interface{}{})
	for _, item := range data.Ipv6ImportRouteTargets {
		itemBody := ""

		if !item.RouteTargetVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "rt.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "rt.value", item.RouteTargetVariable.ValueString())
			}
		} else if !item.RouteTarget.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "rt.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "rt.value", item.RouteTarget.ValueString())
		}
		body, _ = sjson.SetRaw(body, path+"mplsVpnIpv6RouteTarget.importRtList.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"mplsVpnIpv6RouteTarget.exportRtList", []interface{}{})
	for _, item := range data.Ipv6ExportRouteTargets {
		itemBody := ""

		if !item.RouteTargetVariable.IsNull() {
			if true {
				itemBody, _ = sjson.Set(itemBody, "rt.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "rt.value", item.RouteTargetVariable.ValueString())
			}
		} else if !item.RouteTarget.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "rt.optionType", "global")
			itemBody, _ = sjson.Set(itemBody, "rt.value", item.RouteTarget.ValueString())
		}
		body, _ = sjson.SetRaw(body, path+"mplsVpnIpv6RouteTarget.exportRtList.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *ServiceLANVPN) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "newHostMapping"); value.Exists() {
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
	}
	if value := res.Get(path + "ompAdvertiseIp4"); value.Exists() {
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
			if cValue := v.Get("prefixList"); cValue.Exists() {
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
	}
	if value := res.Get(path + "ompAdvertiseIpv6"); value.Exists() {
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
			if cValue := v.Get("prefixList"); cValue.Exists() {
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
					item.Prefixes = append(item.Prefixes, cItem)
					return true
				})
			}
			data.AdvertiseOmpIpv6s = append(data.AdvertiseOmpIpv6s, item)
			return true
		})
	}
	if value := res.Get(path + "ipv4Route"); value.Exists() {
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
			if cValue := v.Get("oneOfIpRoute.nextHopContainer.nextHop"); cValue.Exists() {
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
			}
			if cValue := v.Get("oneOfIpRoute.nextHopContainer.nextHopWithTracker"); cValue.Exists() {
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
			}
			item.Null0 = types.BoolNull()

			if t := v.Get("oneOfIpRoute.null0.optionType"); t.Exists() {
				va := v.Get("oneOfIpRoute.null0.value")
				if t.String() == "global" {
					item.Null0 = types.BoolValue(va.Bool())
				}
			}
			item.GatewayDhcp = types.BoolNull()

			if t := v.Get("oneOfIpRoute.dhcp.optionType"); t.Exists() {
				va := v.Get("oneOfIpRoute.dhcp.value")
				if t.String() == "global" {
					item.GatewayDhcp = types.BoolValue(va.Bool())
				}
			}
			item.Vpn = types.BoolNull()

			if t := v.Get("oneOfIpRoute.vpn.optionType"); t.Exists() {
				va := v.Get("oneOfIpRoute.vpn.value")
				if t.String() == "global" {
					item.Vpn = types.BoolValue(va.Bool())
				}
			}
			data.Ipv4StaticRoutes = append(data.Ipv4StaticRoutes, item)
			return true
		})
	}
	if value := res.Get(path + "ipv6Route"); value.Exists() {
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
			if cValue := v.Get("oneOfIpRoute.nextHopContainer.nextHop"); cValue.Exists() {
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
			}
			item.Null0 = types.BoolNull()

			if t := v.Get("oneOfIpRoute.null0.optionType"); t.Exists() {
				va := v.Get("oneOfIpRoute.null0.value")
				if t.String() == "global" {
					item.Null0 = types.BoolValue(va.Bool())
				}
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
			}
			data.Ipv6StaticRoutes = append(data.Ipv6StaticRoutes, item)
			return true
		})
	}
	if value := res.Get(path + "service"); value.Exists() {
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
	}
	if value := res.Get(path + "serviceRoute"); value.Exists() {
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
			data.ServiceRoutes = append(data.ServiceRoutes, item)
			return true
		})
	}
	if value := res.Get(path + "greRoute"); value.Exists() {
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
	}
	if value := res.Get(path + "ipsecRoute"); value.Exists() {
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
	}
	if value := res.Get(path + "natPool"); value.Exists() {
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
	}
	if value := res.Get(path + "natPortForward"); value.Exists() {
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
			if t := v.Get("TranslatedSourceIp.optionType"); t.Exists() {
				va := v.Get("TranslatedSourceIp.value")
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
	}
	if value := res.Get(path + "staticNat"); value.Exists() {
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
			if t := v.Get("TranslatedSourceIp.optionType"); t.Exists() {
				va := v.Get("TranslatedSourceIp.value")
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
	}
	if value := res.Get(path + "nat64V4Pool"); value.Exists() {
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
	}
	if value := res.Get(path + "routeLeakFromGlobal"); value.Exists() {
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
			if cValue := v.Get("redistributeToProtocol"); cValue.Exists() {
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
	}
	if value := res.Get(path + "routeLeakFromService"); value.Exists() {
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
			if cValue := v.Get("redistributeToProtocol"); cValue.Exists() {
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
	}
	if value := res.Get(path + "routeLeakBetweenServices"); value.Exists() {
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
			if cValue := v.Get("redistributeToProtocol"); cValue.Exists() {
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
	}
	if value := res.Get(path + "mplsVpnIpv4RouteTarget.importRtList"); value.Exists() {
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
	}
	if value := res.Get(path + "mplsVpnIpv4RouteTarget.exportRtList"); value.Exists() {
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
	}
	if value := res.Get(path + "mplsVpnIpv6RouteTarget.importRtList"); value.Exists() {
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
	}
	if value := res.Get(path + "mplsVpnIpv6RouteTarget.exportRtList"); value.Exists() {
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
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *ServiceLANVPN) updateFromBody(ctx context.Context, res gjson.Result) {
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
	for i := range data.HostMappings {
		keys := [...]string{"hostName"}
		keyValues := [...]string{data.HostMappings[i].HostName.ValueString()}
		keyValuesVariables := [...]string{data.HostMappings[i].HostNameVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "newHostMapping").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.HostMappings[i].HostName = types.StringNull()
		data.HostMappings[i].HostNameVariable = types.StringNull()
		if t := r.Get("hostName.optionType"); t.Exists() {
			va := r.Get("hostName.value")
			if t.String() == "variable" {
				data.HostMappings[i].HostNameVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.HostMappings[i].HostName = types.StringValue(va.String())
			}
		}
		data.HostMappings[i].ListOfIps = types.SetNull(types.StringType)
		data.HostMappings[i].ListOfIpsVariable = types.StringNull()
		if t := r.Get("listOfIp.optionType"); t.Exists() {
			va := r.Get("listOfIp.value")
			if t.String() == "variable" {
				data.HostMappings[i].ListOfIpsVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.HostMappings[i].ListOfIps = helpers.GetStringSet(va.Array())
			}
		}
	}
	for i := range data.AdvertiseOmpIpv4s {
		keys := [...]string{"routePolicy.refId"}
		keyValues := [...]string{data.AdvertiseOmpIpv4s[i].RoutePolicyId.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "ompAdvertiseIp4").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.AdvertiseOmpIpv4s[i].Protocol = types.StringNull()
		data.AdvertiseOmpIpv4s[i].ProtocolVariable = types.StringNull()
		if t := r.Get("ompProtocol.optionType"); t.Exists() {
			va := r.Get("ompProtocol.value")
			if t.String() == "variable" {
				data.AdvertiseOmpIpv4s[i].ProtocolVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.AdvertiseOmpIpv4s[i].Protocol = types.StringValue(va.String())
			}
		}
		data.AdvertiseOmpIpv4s[i].RoutePolicyId = types.StringNull()

		if t := r.Get("routePolicy.refId.optionType"); t.Exists() {
			va := r.Get("routePolicy.refId.value")
			if t.String() == "global" {
				data.AdvertiseOmpIpv4s[i].RoutePolicyId = types.StringValue(va.String())
			}
		}
		for ci := range data.AdvertiseOmpIpv4s[i].Prefixes {
			keys := [...]string{"prefix.address"}
			keyValues := [...]string{data.AdvertiseOmpIpv4s[i].Prefixes[ci].NetworkAddress.ValueString()}
			keyValuesVariables := [...]string{data.AdvertiseOmpIpv4s[i].Prefixes[ci].NetworkAddressVariable.ValueString()}

			var cr gjson.Result
			r.Get("prefixList").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType").String()
						vv := v.Get(keys[ik] + ".value").String()
						if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			data.AdvertiseOmpIpv4s[i].Prefixes[ci].NetworkAddress = types.StringNull()
			data.AdvertiseOmpIpv4s[i].Prefixes[ci].NetworkAddressVariable = types.StringNull()
			if t := cr.Get("prefix.address.optionType"); t.Exists() {
				va := cr.Get("prefix.address.value")
				if t.String() == "variable" {
					data.AdvertiseOmpIpv4s[i].Prefixes[ci].NetworkAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.AdvertiseOmpIpv4s[i].Prefixes[ci].NetworkAddress = types.StringValue(va.String())
				}
			}
			data.AdvertiseOmpIpv4s[i].Prefixes[ci].SubnetMask = types.StringNull()
			data.AdvertiseOmpIpv4s[i].Prefixes[ci].SubnetMaskVariable = types.StringNull()
			if t := cr.Get("prefix.mask.optionType"); t.Exists() {
				va := cr.Get("prefix.mask.value")
				if t.String() == "variable" {
					data.AdvertiseOmpIpv4s[i].Prefixes[ci].SubnetMaskVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.AdvertiseOmpIpv4s[i].Prefixes[ci].SubnetMask = types.StringValue(va.String())
				}
			}
			data.AdvertiseOmpIpv4s[i].Prefixes[ci].AggregateOnly = types.BoolNull()

			if t := cr.Get("aggregateOnly.optionType"); t.Exists() {
				va := cr.Get("aggregateOnly.value")
				if t.String() == "global" {
					data.AdvertiseOmpIpv4s[i].Prefixes[ci].AggregateOnly = types.BoolValue(va.Bool())
				}
			}
			data.AdvertiseOmpIpv4s[i].Prefixes[ci].Region = types.StringNull()
			data.AdvertiseOmpIpv4s[i].Prefixes[ci].RegionVariable = types.StringNull()
			if t := cr.Get("region.optionType"); t.Exists() {
				va := cr.Get("region.value")
				if t.String() == "variable" {
					data.AdvertiseOmpIpv4s[i].Prefixes[ci].RegionVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.AdvertiseOmpIpv4s[i].Prefixes[ci].Region = types.StringValue(va.String())
				}
			}
		}
	}
	for i := range data.AdvertiseOmpIpv6s {
		keys := [...]string{"routePolicy.refId"}
		keyValues := [...]string{data.AdvertiseOmpIpv6s[i].RoutePolicyId.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "ompAdvertiseIpv6").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.AdvertiseOmpIpv6s[i].Protocol = types.StringNull()
		data.AdvertiseOmpIpv6s[i].ProtocolVariable = types.StringNull()
		if t := r.Get("ompProtocol.optionType"); t.Exists() {
			va := r.Get("ompProtocol.value")
			if t.String() == "variable" {
				data.AdvertiseOmpIpv6s[i].ProtocolVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.AdvertiseOmpIpv6s[i].Protocol = types.StringValue(va.String())
			}
		}
		data.AdvertiseOmpIpv6s[i].RoutePolicyId = types.StringNull()

		if t := r.Get("routePolicy.refId.optionType"); t.Exists() {
			va := r.Get("routePolicy.refId.value")
			if t.String() == "global" {
				data.AdvertiseOmpIpv6s[i].RoutePolicyId = types.StringValue(va.String())
			}
		}
		data.AdvertiseOmpIpv6s[i].ProtocolSubType = types.StringNull()
		data.AdvertiseOmpIpv6s[i].ProtocolSubTypeVariable = types.StringNull()
		if t := r.Get("protocolSubType.optionType"); t.Exists() {
			va := r.Get("protocolSubType.value")
			if t.String() == "variable" {
				data.AdvertiseOmpIpv6s[i].ProtocolSubTypeVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.AdvertiseOmpIpv6s[i].ProtocolSubType = types.StringValue(va.String())
			}
		}
		for ci := range data.AdvertiseOmpIpv6s[i].Prefixes {
			keys := [...]string{"prefix"}
			keyValues := [...]string{data.AdvertiseOmpIpv6s[i].Prefixes[ci].Prefix.ValueString()}
			keyValuesVariables := [...]string{data.AdvertiseOmpIpv6s[i].Prefixes[ci].PrefixVariable.ValueString()}

			var cr gjson.Result
			r.Get("prefixList").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType").String()
						vv := v.Get(keys[ik] + ".value").String()
						if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			data.AdvertiseOmpIpv6s[i].Prefixes[ci].Prefix = types.StringNull()
			data.AdvertiseOmpIpv6s[i].Prefixes[ci].PrefixVariable = types.StringNull()
			if t := cr.Get("prefix.optionType"); t.Exists() {
				va := cr.Get("prefix.value")
				if t.String() == "variable" {
					data.AdvertiseOmpIpv6s[i].Prefixes[ci].PrefixVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.AdvertiseOmpIpv6s[i].Prefixes[ci].Prefix = types.StringValue(va.String())
				}
			}
			data.AdvertiseOmpIpv6s[i].Prefixes[ci].AggregateOnly = types.BoolNull()

			if t := cr.Get("aggregateOnly.optionType"); t.Exists() {
				va := cr.Get("aggregateOnly.value")
				if t.String() == "global" {
					data.AdvertiseOmpIpv6s[i].Prefixes[ci].AggregateOnly = types.BoolValue(va.Bool())
				}
			}
		}
	}
	for i := range data.Ipv4StaticRoutes {
		keys := [...]string{"prefix.ipAddress"}
		keyValues := [...]string{data.Ipv4StaticRoutes[i].NetworkAddress.ValueString()}
		keyValuesVariables := [...]string{data.Ipv4StaticRoutes[i].NetworkAddressVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "ipv4Route").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Ipv4StaticRoutes[i].NetworkAddress = types.StringNull()
		data.Ipv4StaticRoutes[i].NetworkAddressVariable = types.StringNull()
		if t := r.Get("prefix.ipAddress.optionType"); t.Exists() {
			va := r.Get("prefix.ipAddress.value")
			if t.String() == "variable" {
				data.Ipv4StaticRoutes[i].NetworkAddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4StaticRoutes[i].NetworkAddress = types.StringValue(va.String())
			}
		}
		data.Ipv4StaticRoutes[i].SubnetMask = types.StringNull()
		data.Ipv4StaticRoutes[i].SubnetMaskVariable = types.StringNull()
		if t := r.Get("prefix.subnetMask.optionType"); t.Exists() {
			va := r.Get("prefix.subnetMask.value")
			if t.String() == "variable" {
				data.Ipv4StaticRoutes[i].SubnetMaskVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4StaticRoutes[i].SubnetMask = types.StringValue(va.String())
			}
		}
		for ci := range data.Ipv4StaticRoutes[i].NextHops {
			keys := [...]string{"address"}
			keyValues := [...]string{data.Ipv4StaticRoutes[i].NextHops[ci].Address.ValueString()}
			keyValuesVariables := [...]string{data.Ipv4StaticRoutes[i].NextHops[ci].AddressVariable.ValueString()}

			var cr gjson.Result
			r.Get("oneOfIpRoute.nextHopContainer.nextHop").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType").String()
						vv := v.Get(keys[ik] + ".value").String()
						if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			data.Ipv4StaticRoutes[i].NextHops[ci].Address = types.StringNull()
			data.Ipv4StaticRoutes[i].NextHops[ci].AddressVariable = types.StringNull()
			if t := cr.Get("address.optionType"); t.Exists() {
				va := cr.Get("address.value")
				if t.String() == "variable" {
					data.Ipv4StaticRoutes[i].NextHops[ci].AddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv4StaticRoutes[i].NextHops[ci].Address = types.StringValue(va.String())
				}
			}
			data.Ipv4StaticRoutes[i].NextHops[ci].AdministrativeDistance = types.Int64Null()
			data.Ipv4StaticRoutes[i].NextHops[ci].AdministrativeDistanceVariable = types.StringNull()
			if t := cr.Get("distance.optionType"); t.Exists() {
				va := cr.Get("distance.value")
				if t.String() == "variable" {
					data.Ipv4StaticRoutes[i].NextHops[ci].AdministrativeDistanceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv4StaticRoutes[i].NextHops[ci].AdministrativeDistance = types.Int64Value(va.Int())
				}
			}
		}
		for ci := range data.Ipv4StaticRoutes[i].NextHopWithTrackers {
			keys := [...]string{}
			keyValues := [...]string{}
			keyValuesVariables := [...]string{}

			var cr gjson.Result
			r.Get("oneOfIpRoute.nextHopContainer.nextHopWithTracker").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType").String()
						vv := v.Get(keys[ik] + ".value").String()
						if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			data.Ipv4StaticRoutes[i].NextHopWithTrackers[ci].Address = types.StringNull()
			data.Ipv4StaticRoutes[i].NextHopWithTrackers[ci].AddressVariable = types.StringNull()
			if t := cr.Get("address.optionType"); t.Exists() {
				va := cr.Get("address.value")
				if t.String() == "variable" {
					data.Ipv4StaticRoutes[i].NextHopWithTrackers[ci].AddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv4StaticRoutes[i].NextHopWithTrackers[ci].Address = types.StringValue(va.String())
				}
			}
			data.Ipv4StaticRoutes[i].NextHopWithTrackers[ci].AdministrativeDistance = types.Int64Null()
			data.Ipv4StaticRoutes[i].NextHopWithTrackers[ci].AdministrativeDistanceVariable = types.StringNull()
			if t := cr.Get("distance.optionType"); t.Exists() {
				va := cr.Get("distance.value")
				if t.String() == "variable" {
					data.Ipv4StaticRoutes[i].NextHopWithTrackers[ci].AdministrativeDistanceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv4StaticRoutes[i].NextHopWithTrackers[ci].AdministrativeDistance = types.Int64Value(va.Int())
				}
			}
			data.Ipv4StaticRoutes[i].NextHopWithTrackers[ci].TrackerId = types.StringNull()

			if t := cr.Get("tracker.refId.optionType"); t.Exists() {
				va := cr.Get("tracker.refId.value")
				if t.String() == "global" {
					data.Ipv4StaticRoutes[i].NextHopWithTrackers[ci].TrackerId = types.StringValue(va.String())
				}
			}
		}
		data.Ipv4StaticRoutes[i].Null0 = types.BoolNull()

		if t := r.Get("oneOfIpRoute.null0.optionType"); t.Exists() {
			va := r.Get("oneOfIpRoute.null0.value")
			if t.String() == "global" {
				data.Ipv4StaticRoutes[i].Null0 = types.BoolValue(va.Bool())
			}
		}
		data.Ipv4StaticRoutes[i].GatewayDhcp = types.BoolNull()

		if t := r.Get("oneOfIpRoute.dhcp.optionType"); t.Exists() {
			va := r.Get("oneOfIpRoute.dhcp.value")
			if t.String() == "global" {
				data.Ipv4StaticRoutes[i].GatewayDhcp = types.BoolValue(va.Bool())
			}
		}
		data.Ipv4StaticRoutes[i].Vpn = types.BoolNull()

		if t := r.Get("oneOfIpRoute.vpn.optionType"); t.Exists() {
			va := r.Get("oneOfIpRoute.vpn.value")
			if t.String() == "global" {
				data.Ipv4StaticRoutes[i].Vpn = types.BoolValue(va.Bool())
			}
		}
	}
	for i := range data.Ipv6StaticRoutes {
		keys := [...]string{"prefix"}
		keyValues := [...]string{data.Ipv6StaticRoutes[i].Prefix.ValueString()}
		keyValuesVariables := [...]string{data.Ipv6StaticRoutes[i].PrefixVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "ipv6Route").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Ipv6StaticRoutes[i].Prefix = types.StringNull()
		data.Ipv6StaticRoutes[i].PrefixVariable = types.StringNull()
		if t := r.Get("prefix.optionType"); t.Exists() {
			va := r.Get("prefix.value")
			if t.String() == "variable" {
				data.Ipv6StaticRoutes[i].PrefixVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6StaticRoutes[i].Prefix = types.StringValue(va.String())
			}
		}
		for ci := range data.Ipv6StaticRoutes[i].NextHops {
			keys := [...]string{"address"}
			keyValues := [...]string{data.Ipv6StaticRoutes[i].NextHops[ci].Address.ValueString()}
			keyValuesVariables := [...]string{data.Ipv6StaticRoutes[i].NextHops[ci].AddressVariable.ValueString()}

			var cr gjson.Result
			r.Get("oneOfIpRoute.nextHopContainer.nextHop").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType").String()
						vv := v.Get(keys[ik] + ".value").String()
						if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			data.Ipv6StaticRoutes[i].NextHops[ci].Address = types.StringNull()
			data.Ipv6StaticRoutes[i].NextHops[ci].AddressVariable = types.StringNull()
			if t := cr.Get("address.optionType"); t.Exists() {
				va := cr.Get("address.value")
				if t.String() == "variable" {
					data.Ipv6StaticRoutes[i].NextHops[ci].AddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv6StaticRoutes[i].NextHops[ci].Address = types.StringValue(va.String())
				}
			}
			data.Ipv6StaticRoutes[i].NextHops[ci].AdministrativeDistance = types.Int64Null()
			data.Ipv6StaticRoutes[i].NextHops[ci].AdministrativeDistanceVariable = types.StringNull()
			if t := cr.Get("distance.optionType"); t.Exists() {
				va := cr.Get("distance.value")
				if t.String() == "variable" {
					data.Ipv6StaticRoutes[i].NextHops[ci].AdministrativeDistanceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv6StaticRoutes[i].NextHops[ci].AdministrativeDistance = types.Int64Value(va.Int())
				}
			}
		}
		data.Ipv6StaticRoutes[i].Null0 = types.BoolNull()

		if t := r.Get("oneOfIpRoute.null0.optionType"); t.Exists() {
			va := r.Get("oneOfIpRoute.null0.value")
			if t.String() == "global" {
				data.Ipv6StaticRoutes[i].Null0 = types.BoolValue(va.Bool())
			}
		}
		data.Ipv6StaticRoutes[i].Nat = types.StringNull()
		data.Ipv6StaticRoutes[i].NatVariable = types.StringNull()
		if t := r.Get("oneOfIpRoute.nat.optionType"); t.Exists() {
			va := r.Get("oneOfIpRoute.nat.value")
			if t.String() == "variable" {
				data.Ipv6StaticRoutes[i].NatVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6StaticRoutes[i].Nat = types.StringValue(va.String())
			}
		}
	}
	for i := range data.Services {
		keys := [...]string{"serviceType"}
		keyValues := [...]string{data.Services[i].ServiceType.ValueString()}
		keyValuesVariables := [...]string{data.Services[i].ServiceTypeVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "service").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Services[i].ServiceType = types.StringNull()
		data.Services[i].ServiceTypeVariable = types.StringNull()
		if t := r.Get("serviceType.optionType"); t.Exists() {
			va := r.Get("serviceType.value")
			if t.String() == "variable" {
				data.Services[i].ServiceTypeVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Services[i].ServiceType = types.StringValue(va.String())
			}
		}
		data.Services[i].Ipv4Addresses = types.SetNull(types.StringType)
		data.Services[i].Ipv4AddressesVariable = types.StringNull()
		if t := r.Get("ipv4Addresses.optionType"); t.Exists() {
			va := r.Get("ipv4Addresses.value")
			if t.String() == "variable" {
				data.Services[i].Ipv4AddressesVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Services[i].Ipv4Addresses = helpers.GetStringSet(va.Array())
			}
		}
		data.Services[i].Tracking = types.BoolNull()
		data.Services[i].TrackingVariable = types.StringNull()
		if t := r.Get("tracking.optionType"); t.Exists() {
			va := r.Get("tracking.value")
			if t.String() == "variable" {
				data.Services[i].TrackingVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Services[i].Tracking = types.BoolValue(va.Bool())
			}
		}
	}
	for i := range data.ServiceRoutes {
		keys := [...]string{"prefix.ipAddress"}
		keyValues := [...]string{data.ServiceRoutes[i].NetworkAddress.ValueString()}
		keyValuesVariables := [...]string{data.ServiceRoutes[i].NetworkAddressVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "serviceRoute").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.ServiceRoutes[i].NetworkAddress = types.StringNull()
		data.ServiceRoutes[i].NetworkAddressVariable = types.StringNull()
		if t := r.Get("prefix.ipAddress.optionType"); t.Exists() {
			va := r.Get("prefix.ipAddress.value")
			if t.String() == "variable" {
				data.ServiceRoutes[i].NetworkAddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.ServiceRoutes[i].NetworkAddress = types.StringValue(va.String())
			}
		}
		data.ServiceRoutes[i].SubnetMask = types.StringNull()
		data.ServiceRoutes[i].SubnetMaskVariable = types.StringNull()
		if t := r.Get("prefix.subnetMask.optionType"); t.Exists() {
			va := r.Get("prefix.subnetMask.value")
			if t.String() == "variable" {
				data.ServiceRoutes[i].SubnetMaskVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.ServiceRoutes[i].SubnetMask = types.StringValue(va.String())
			}
		}
		data.ServiceRoutes[i].Service = types.StringNull()
		data.ServiceRoutes[i].ServiceVariable = types.StringNull()
		if t := r.Get("service.optionType"); t.Exists() {
			va := r.Get("service.value")
			if t.String() == "variable" {
				data.ServiceRoutes[i].ServiceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.ServiceRoutes[i].Service = types.StringValue(va.String())
			}
		}
		data.ServiceRoutes[i].Vpn = types.Int64Null()

		if t := r.Get("vpn.optionType"); t.Exists() {
			va := r.Get("vpn.value")
			if t.String() == "global" {
				data.ServiceRoutes[i].Vpn = types.Int64Value(va.Int())
			}
		}
	}
	for i := range data.GreRoutes {
		keys := [...]string{"prefix.ipAddress"}
		keyValues := [...]string{data.GreRoutes[i].NetworkAddress.ValueString()}
		keyValuesVariables := [...]string{data.GreRoutes[i].NetworkAddressVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "greRoute").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.GreRoutes[i].NetworkAddress = types.StringNull()
		data.GreRoutes[i].NetworkAddressVariable = types.StringNull()
		if t := r.Get("prefix.ipAddress.optionType"); t.Exists() {
			va := r.Get("prefix.ipAddress.value")
			if t.String() == "variable" {
				data.GreRoutes[i].NetworkAddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.GreRoutes[i].NetworkAddress = types.StringValue(va.String())
			}
		}
		data.GreRoutes[i].SubnetMask = types.StringNull()
		data.GreRoutes[i].SubnetMaskVariable = types.StringNull()
		if t := r.Get("prefix.subnetMask.optionType"); t.Exists() {
			va := r.Get("prefix.subnetMask.value")
			if t.String() == "variable" {
				data.GreRoutes[i].SubnetMaskVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.GreRoutes[i].SubnetMask = types.StringValue(va.String())
			}
		}
		data.GreRoutes[i].Interface = types.SetNull(types.StringType)
		data.GreRoutes[i].InterfaceVariable = types.StringNull()
		if t := r.Get("interface.optionType"); t.Exists() {
			va := r.Get("interface.value")
			if t.String() == "variable" {
				data.GreRoutes[i].InterfaceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.GreRoutes[i].Interface = helpers.GetStringSet(va.Array())
			}
		}
		data.GreRoutes[i].Vpn = types.Int64Null()

		if t := r.Get("vpn.optionType"); t.Exists() {
			va := r.Get("vpn.value")
			if t.String() == "global" {
				data.GreRoutes[i].Vpn = types.Int64Value(va.Int())
			}
		}
	}
	for i := range data.IpsecRoutes {
		keys := [...]string{"prefix.ipAddress"}
		keyValues := [...]string{data.IpsecRoutes[i].NetworkAddress.ValueString()}
		keyValuesVariables := [...]string{data.IpsecRoutes[i].NetworkAddressVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "ipsecRoute").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.IpsecRoutes[i].NetworkAddress = types.StringNull()
		data.IpsecRoutes[i].NetworkAddressVariable = types.StringNull()
		if t := r.Get("prefix.ipAddress.optionType"); t.Exists() {
			va := r.Get("prefix.ipAddress.value")
			if t.String() == "variable" {
				data.IpsecRoutes[i].NetworkAddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.IpsecRoutes[i].NetworkAddress = types.StringValue(va.String())
			}
		}
		data.IpsecRoutes[i].SubnetMask = types.StringNull()
		data.IpsecRoutes[i].SubnetMaskVariable = types.StringNull()
		if t := r.Get("prefix.subnetMask.optionType"); t.Exists() {
			va := r.Get("prefix.subnetMask.value")
			if t.String() == "variable" {
				data.IpsecRoutes[i].SubnetMaskVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.IpsecRoutes[i].SubnetMask = types.StringValue(va.String())
			}
		}
		data.IpsecRoutes[i].Interface = types.SetNull(types.StringType)
		data.IpsecRoutes[i].InterfaceVariable = types.StringNull()
		if t := r.Get("interface.optionType"); t.Exists() {
			va := r.Get("interface.value")
			if t.String() == "variable" {
				data.IpsecRoutes[i].InterfaceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.IpsecRoutes[i].Interface = helpers.GetStringSet(va.Array())
			}
		}
	}
	for i := range data.NatPools {
		keys := [...]string{"natPoolName"}
		keyValues := [...]string{strconv.FormatInt(data.NatPools[i].NatPoolName.ValueInt64(), 10)}
		keyValuesVariables := [...]string{data.NatPools[i].NatPoolNameVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "natPool").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.NatPools[i].NatPoolName = types.Int64Null()
		data.NatPools[i].NatPoolNameVariable = types.StringNull()
		if t := r.Get("natPoolName.optionType"); t.Exists() {
			va := r.Get("natPoolName.value")
			if t.String() == "variable" {
				data.NatPools[i].NatPoolNameVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NatPools[i].NatPoolName = types.Int64Value(va.Int())
			}
		}
		data.NatPools[i].PrefixLength = types.Int64Null()
		data.NatPools[i].PrefixLengthVariable = types.StringNull()
		if t := r.Get("prefixLength.optionType"); t.Exists() {
			va := r.Get("prefixLength.value")
			if t.String() == "variable" {
				data.NatPools[i].PrefixLengthVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NatPools[i].PrefixLength = types.Int64Value(va.Int())
			}
		}
		data.NatPools[i].RangeStart = types.StringNull()
		data.NatPools[i].RangeStartVariable = types.StringNull()
		if t := r.Get("rangeStart.optionType"); t.Exists() {
			va := r.Get("rangeStart.value")
			if t.String() == "variable" {
				data.NatPools[i].RangeStartVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NatPools[i].RangeStart = types.StringValue(va.String())
			}
		}
		data.NatPools[i].RangeEnd = types.StringNull()
		data.NatPools[i].RangeEndVariable = types.StringNull()
		if t := r.Get("rangeEnd.optionType"); t.Exists() {
			va := r.Get("rangeEnd.value")
			if t.String() == "variable" {
				data.NatPools[i].RangeEndVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NatPools[i].RangeEnd = types.StringValue(va.String())
			}
		}
		data.NatPools[i].Overload = types.BoolNull()
		data.NatPools[i].OverloadVariable = types.StringNull()
		if t := r.Get("overload.optionType"); t.Exists() {
			va := r.Get("overload.value")
			if t.String() == "variable" {
				data.NatPools[i].OverloadVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NatPools[i].Overload = types.BoolValue(va.Bool())
			}
		}
		data.NatPools[i].Direction = types.StringNull()
		data.NatPools[i].DirectionVariable = types.StringNull()
		if t := r.Get("direction.optionType"); t.Exists() {
			va := r.Get("direction.value")
			if t.String() == "variable" {
				data.NatPools[i].DirectionVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NatPools[i].Direction = types.StringValue(va.String())
			}
		}
		data.NatPools[i].TrackerObjectId = types.StringNull()

		if t := r.Get("trackingObject.trackerId.refId.optionType"); t.Exists() {
			va := r.Get("trackingObject.trackerId.refId.value")
			if t.String() == "global" {
				data.NatPools[i].TrackerObjectId = types.StringValue(va.String())
			}
		}
	}
	for i := range data.NatPortForwards {
		keys := [...]string{"natPoolName"}
		keyValues := [...]string{strconv.FormatInt(data.NatPortForwards[i].NatPoolName.ValueInt64(), 10)}
		keyValuesVariables := [...]string{data.NatPortForwards[i].NatPoolNameVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "natPortForward").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.NatPortForwards[i].NatPoolName = types.Int64Null()
		data.NatPortForwards[i].NatPoolNameVariable = types.StringNull()
		if t := r.Get("natPoolName.optionType"); t.Exists() {
			va := r.Get("natPoolName.value")
			if t.String() == "variable" {
				data.NatPortForwards[i].NatPoolNameVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NatPortForwards[i].NatPoolName = types.Int64Value(va.Int())
			}
		}
		data.NatPortForwards[i].SourcePort = types.Int64Null()
		data.NatPortForwards[i].SourcePortVariable = types.StringNull()
		if t := r.Get("sourcePort.optionType"); t.Exists() {
			va := r.Get("sourcePort.value")
			if t.String() == "variable" {
				data.NatPortForwards[i].SourcePortVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NatPortForwards[i].SourcePort = types.Int64Value(va.Int())
			}
		}
		data.NatPortForwards[i].TranslatePort = types.Int64Null()
		data.NatPortForwards[i].TranslatePortVariable = types.StringNull()
		if t := r.Get("translatePort.optionType"); t.Exists() {
			va := r.Get("translatePort.value")
			if t.String() == "variable" {
				data.NatPortForwards[i].TranslatePortVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NatPortForwards[i].TranslatePort = types.Int64Value(va.Int())
			}
		}
		data.NatPortForwards[i].SourceIp = types.StringNull()
		data.NatPortForwards[i].SourceIpVariable = types.StringNull()
		if t := r.Get("sourceIp.optionType"); t.Exists() {
			va := r.Get("sourceIp.value")
			if t.String() == "variable" {
				data.NatPortForwards[i].SourceIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NatPortForwards[i].SourceIp = types.StringValue(va.String())
			}
		}
		data.NatPortForwards[i].TranslatedSourceIp = types.StringNull()
		data.NatPortForwards[i].TranslatedSourceIpVariable = types.StringNull()
		if t := r.Get("TranslatedSourceIp.optionType"); t.Exists() {
			va := r.Get("TranslatedSourceIp.value")
			if t.String() == "variable" {
				data.NatPortForwards[i].TranslatedSourceIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NatPortForwards[i].TranslatedSourceIp = types.StringValue(va.String())
			}
		}
		data.NatPortForwards[i].Protocol = types.StringNull()
		data.NatPortForwards[i].ProtocolVariable = types.StringNull()
		if t := r.Get("protocol.optionType"); t.Exists() {
			va := r.Get("protocol.value")
			if t.String() == "variable" {
				data.NatPortForwards[i].ProtocolVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.NatPortForwards[i].Protocol = types.StringValue(va.String())
			}
		}
	}
	for i := range data.StaticNats {
		keys := [...]string{"natPoolName"}
		keyValues := [...]string{strconv.FormatInt(data.StaticNats[i].NatPoolName.ValueInt64(), 10)}
		keyValuesVariables := [...]string{data.StaticNats[i].NatPoolNameVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "staticNat").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.StaticNats[i].NatPoolName = types.Int64Null()
		data.StaticNats[i].NatPoolNameVariable = types.StringNull()
		if t := r.Get("natPoolName.optionType"); t.Exists() {
			va := r.Get("natPoolName.value")
			if t.String() == "variable" {
				data.StaticNats[i].NatPoolNameVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.StaticNats[i].NatPoolName = types.Int64Value(va.Int())
			}
		}
		data.StaticNats[i].SourceIp = types.StringNull()
		data.StaticNats[i].SourceIpVariable = types.StringNull()
		if t := r.Get("sourceIp.optionType"); t.Exists() {
			va := r.Get("sourceIp.value")
			if t.String() == "variable" {
				data.StaticNats[i].SourceIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.StaticNats[i].SourceIp = types.StringValue(va.String())
			}
		}
		data.StaticNats[i].TranslatedSourceIp = types.StringNull()
		data.StaticNats[i].TranslatedSourceIpVariable = types.StringNull()
		if t := r.Get("TranslatedSourceIp.optionType"); t.Exists() {
			va := r.Get("TranslatedSourceIp.value")
			if t.String() == "variable" {
				data.StaticNats[i].TranslatedSourceIpVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.StaticNats[i].TranslatedSourceIp = types.StringValue(va.String())
			}
		}
		data.StaticNats[i].StaticNatDirection = types.StringNull()
		data.StaticNats[i].StaticNatDirectionVariable = types.StringNull()
		if t := r.Get("staticNatDirection.optionType"); t.Exists() {
			va := r.Get("staticNatDirection.value")
			if t.String() == "variable" {
				data.StaticNats[i].StaticNatDirectionVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.StaticNats[i].StaticNatDirection = types.StringValue(va.String())
			}
		}
		data.StaticNats[i].TrackerObjectId = types.StringNull()

		if t := r.Get("trackingObject.trackerId.refId.optionType"); t.Exists() {
			va := r.Get("trackingObject.trackerId.refId.value")
			if t.String() == "global" {
				data.StaticNats[i].TrackerObjectId = types.StringValue(va.String())
			}
		}
	}
	for i := range data.Nat64V4Pools {
		keys := [...]string{"nat64V4PoolName"}
		keyValues := [...]string{data.Nat64V4Pools[i].Name.ValueString()}
		keyValuesVariables := [...]string{data.Nat64V4Pools[i].NameVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "nat64V4Pool").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Nat64V4Pools[i].Name = types.StringNull()
		data.Nat64V4Pools[i].NameVariable = types.StringNull()
		if t := r.Get("nat64V4PoolName.optionType"); t.Exists() {
			va := r.Get("nat64V4PoolName.value")
			if t.String() == "variable" {
				data.Nat64V4Pools[i].NameVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Nat64V4Pools[i].Name = types.StringValue(va.String())
			}
		}
		data.Nat64V4Pools[i].RangeStart = types.StringNull()
		data.Nat64V4Pools[i].RangeStartVariable = types.StringNull()
		if t := r.Get("nat64V4PoolRangeStart.optionType"); t.Exists() {
			va := r.Get("nat64V4PoolRangeStart.value")
			if t.String() == "variable" {
				data.Nat64V4Pools[i].RangeStartVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Nat64V4Pools[i].RangeStart = types.StringValue(va.String())
			}
		}
		data.Nat64V4Pools[i].RangeEnd = types.StringNull()
		data.Nat64V4Pools[i].RangeEndVariable = types.StringNull()
		if t := r.Get("nat64V4PoolRangeEnd.optionType"); t.Exists() {
			va := r.Get("nat64V4PoolRangeEnd.value")
			if t.String() == "variable" {
				data.Nat64V4Pools[i].RangeEndVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Nat64V4Pools[i].RangeEnd = types.StringValue(va.String())
			}
		}
		data.Nat64V4Pools[i].Overload = types.BoolNull()
		data.Nat64V4Pools[i].OverloadVariable = types.StringNull()
		if t := r.Get("nat64V4PoolOverload.optionType"); t.Exists() {
			va := r.Get("nat64V4PoolOverload.value")
			if t.String() == "variable" {
				data.Nat64V4Pools[i].OverloadVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Nat64V4Pools[i].Overload = types.BoolValue(va.Bool())
			}
		}
	}
	for i := range data.RouteLeakFromGlobalVpns {
		keys := [...]string{"routePolicy.refId"}
		keyValues := [...]string{data.RouteLeakFromGlobalVpns[i].RoutePolicyId.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "routeLeakFromGlobal").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.RouteLeakFromGlobalVpns[i].RouteProtocol = types.StringNull()
		data.RouteLeakFromGlobalVpns[i].RouteProtocolVariable = types.StringNull()
		if t := r.Get("routeProtocol.optionType"); t.Exists() {
			va := r.Get("routeProtocol.value")
			if t.String() == "variable" {
				data.RouteLeakFromGlobalVpns[i].RouteProtocolVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.RouteLeakFromGlobalVpns[i].RouteProtocol = types.StringValue(va.String())
			}
		}
		data.RouteLeakFromGlobalVpns[i].RoutePolicyId = types.StringNull()

		if t := r.Get("routePolicy.refId.optionType"); t.Exists() {
			va := r.Get("routePolicy.refId.value")
			if t.String() == "global" {
				data.RouteLeakFromGlobalVpns[i].RoutePolicyId = types.StringValue(va.String())
			}
		}
		for ci := range data.RouteLeakFromGlobalVpns[i].Redistributions {
			keys := [...]string{"policy.refId"}
			keyValues := [...]string{data.RouteLeakFromGlobalVpns[i].Redistributions[ci].RedistributionPolicyId.ValueString()}
			keyValuesVariables := [...]string{""}

			var cr gjson.Result
			r.Get("redistributeToProtocol").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType").String()
						vv := v.Get(keys[ik] + ".value").String()
						if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			data.RouteLeakFromGlobalVpns[i].Redistributions[ci].Protocol = types.StringNull()
			data.RouteLeakFromGlobalVpns[i].Redistributions[ci].ProtocolVariable = types.StringNull()
			if t := cr.Get("protocol.optionType"); t.Exists() {
				va := cr.Get("protocol.value")
				if t.String() == "variable" {
					data.RouteLeakFromGlobalVpns[i].Redistributions[ci].ProtocolVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.RouteLeakFromGlobalVpns[i].Redistributions[ci].Protocol = types.StringValue(va.String())
				}
			}
			data.RouteLeakFromGlobalVpns[i].Redistributions[ci].RedistributionPolicyId = types.StringNull()

			if t := cr.Get("policy.refId.optionType"); t.Exists() {
				va := cr.Get("policy.refId.value")
				if t.String() == "global" {
					data.RouteLeakFromGlobalVpns[i].Redistributions[ci].RedistributionPolicyId = types.StringValue(va.String())
				}
			}
		}
	}
	for i := range data.RouteLeakToGlobalVpns {
		keys := [...]string{"routePolicy.refId"}
		keyValues := [...]string{data.RouteLeakToGlobalVpns[i].RoutePolicyId.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "routeLeakFromService").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.RouteLeakToGlobalVpns[i].RouteProtocol = types.StringNull()
		data.RouteLeakToGlobalVpns[i].RouteProtocolVariable = types.StringNull()
		if t := r.Get("routeProtocol.optionType"); t.Exists() {
			va := r.Get("routeProtocol.value")
			if t.String() == "variable" {
				data.RouteLeakToGlobalVpns[i].RouteProtocolVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.RouteLeakToGlobalVpns[i].RouteProtocol = types.StringValue(va.String())
			}
		}
		data.RouteLeakToGlobalVpns[i].RoutePolicyId = types.StringNull()

		if t := r.Get("routePolicy.refId.optionType"); t.Exists() {
			va := r.Get("routePolicy.refId.value")
			if t.String() == "global" {
				data.RouteLeakToGlobalVpns[i].RoutePolicyId = types.StringValue(va.String())
			}
		}
		for ci := range data.RouteLeakToGlobalVpns[i].Redistributions {
			keys := [...]string{"policy.refId"}
			keyValues := [...]string{data.RouteLeakToGlobalVpns[i].Redistributions[ci].RedistributionPolicyId.ValueString()}
			keyValuesVariables := [...]string{""}

			var cr gjson.Result
			r.Get("redistributeToProtocol").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType").String()
						vv := v.Get(keys[ik] + ".value").String()
						if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			data.RouteLeakToGlobalVpns[i].Redistributions[ci].Protocol = types.StringNull()
			data.RouteLeakToGlobalVpns[i].Redistributions[ci].ProtocolVariable = types.StringNull()
			if t := cr.Get("protocol.optionType"); t.Exists() {
				va := cr.Get("protocol.value")
				if t.String() == "variable" {
					data.RouteLeakToGlobalVpns[i].Redistributions[ci].ProtocolVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.RouteLeakToGlobalVpns[i].Redistributions[ci].Protocol = types.StringValue(va.String())
				}
			}
			data.RouteLeakToGlobalVpns[i].Redistributions[ci].RedistributionPolicyId = types.StringNull()

			if t := cr.Get("policy.refId.optionType"); t.Exists() {
				va := cr.Get("policy.refId.value")
				if t.String() == "global" {
					data.RouteLeakToGlobalVpns[i].Redistributions[ci].RedistributionPolicyId = types.StringValue(va.String())
				}
			}
		}
	}
	for i := range data.RouteLeakFromOtherServices {
		keys := [...]string{"routePolicy.refId"}
		keyValues := [...]string{data.RouteLeakFromOtherServices[i].RoutePolicyId.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "routeLeakBetweenServices").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.RouteLeakFromOtherServices[i].SourceVpn = types.Int64Null()
		data.RouteLeakFromOtherServices[i].SourceVpnVariable = types.StringNull()
		if t := r.Get("sourceVpn.optionType"); t.Exists() {
			va := r.Get("sourceVpn.value")
			if t.String() == "variable" {
				data.RouteLeakFromOtherServices[i].SourceVpnVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.RouteLeakFromOtherServices[i].SourceVpn = types.Int64Value(va.Int())
			}
		}
		data.RouteLeakFromOtherServices[i].RouteProtocol = types.StringNull()
		data.RouteLeakFromOtherServices[i].RouteProtocolVariable = types.StringNull()
		if t := r.Get("routeProtocol.optionType"); t.Exists() {
			va := r.Get("routeProtocol.value")
			if t.String() == "variable" {
				data.RouteLeakFromOtherServices[i].RouteProtocolVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.RouteLeakFromOtherServices[i].RouteProtocol = types.StringValue(va.String())
			}
		}
		data.RouteLeakFromOtherServices[i].RoutePolicyId = types.StringNull()

		if t := r.Get("routePolicy.refId.optionType"); t.Exists() {
			va := r.Get("routePolicy.refId.value")
			if t.String() == "global" {
				data.RouteLeakFromOtherServices[i].RoutePolicyId = types.StringValue(va.String())
			}
		}
		for ci := range data.RouteLeakFromOtherServices[i].Redistributions {
			keys := [...]string{"policy.refId"}
			keyValues := [...]string{data.RouteLeakFromOtherServices[i].Redistributions[ci].RedistributionPolicyId.ValueString()}
			keyValuesVariables := [...]string{""}

			var cr gjson.Result
			r.Get("redistributeToProtocol").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType").String()
						vv := v.Get(keys[ik] + ".value").String()
						if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			data.RouteLeakFromOtherServices[i].Redistributions[ci].Protocol = types.StringNull()
			data.RouteLeakFromOtherServices[i].Redistributions[ci].ProtocolVariable = types.StringNull()
			if t := cr.Get("protocol.optionType"); t.Exists() {
				va := cr.Get("protocol.value")
				if t.String() == "variable" {
					data.RouteLeakFromOtherServices[i].Redistributions[ci].ProtocolVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.RouteLeakFromOtherServices[i].Redistributions[ci].Protocol = types.StringValue(va.String())
				}
			}
			data.RouteLeakFromOtherServices[i].Redistributions[ci].RedistributionPolicyId = types.StringNull()

			if t := cr.Get("policy.refId.optionType"); t.Exists() {
				va := cr.Get("policy.refId.value")
				if t.String() == "global" {
					data.RouteLeakFromOtherServices[i].Redistributions[ci].RedistributionPolicyId = types.StringValue(va.String())
				}
			}
		}
	}
	for i := range data.Ipv4ImportRouteTargets {
		keys := [...]string{"rt"}
		keyValues := [...]string{data.Ipv4ImportRouteTargets[i].RouteTarget.ValueString()}
		keyValuesVariables := [...]string{data.Ipv4ImportRouteTargets[i].RouteTargetVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "mplsVpnIpv4RouteTarget.importRtList").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Ipv4ImportRouteTargets[i].RouteTarget = types.StringNull()
		data.Ipv4ImportRouteTargets[i].RouteTargetVariable = types.StringNull()
		if t := r.Get("rt.optionType"); t.Exists() {
			va := r.Get("rt.value")
			if t.String() == "variable" {
				data.Ipv4ImportRouteTargets[i].RouteTargetVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4ImportRouteTargets[i].RouteTarget = types.StringValue(va.String())
			}
		}
	}
	for i := range data.Ipv4ExportRouteTargets {
		keys := [...]string{"rt"}
		keyValues := [...]string{data.Ipv4ExportRouteTargets[i].RouteTarget.ValueString()}
		keyValuesVariables := [...]string{data.Ipv4ExportRouteTargets[i].RouteTargetVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "mplsVpnIpv4RouteTarget.exportRtList").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Ipv4ExportRouteTargets[i].RouteTarget = types.StringNull()
		data.Ipv4ExportRouteTargets[i].RouteTargetVariable = types.StringNull()
		if t := r.Get("rt.optionType"); t.Exists() {
			va := r.Get("rt.value")
			if t.String() == "variable" {
				data.Ipv4ExportRouteTargets[i].RouteTargetVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4ExportRouteTargets[i].RouteTarget = types.StringValue(va.String())
			}
		}
	}
	for i := range data.Ipv6ImportRouteTargets {
		keys := [...]string{"rt"}
		keyValues := [...]string{data.Ipv6ImportRouteTargets[i].RouteTarget.ValueString()}
		keyValuesVariables := [...]string{data.Ipv6ImportRouteTargets[i].RouteTargetVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "mplsVpnIpv6RouteTarget.importRtList").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Ipv6ImportRouteTargets[i].RouteTarget = types.StringNull()
		data.Ipv6ImportRouteTargets[i].RouteTargetVariable = types.StringNull()
		if t := r.Get("rt.optionType"); t.Exists() {
			va := r.Get("rt.value")
			if t.String() == "variable" {
				data.Ipv6ImportRouteTargets[i].RouteTargetVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6ImportRouteTargets[i].RouteTarget = types.StringValue(va.String())
			}
		}
	}
	for i := range data.Ipv6ExportRouteTargets {
		keys := [...]string{"rt"}
		keyValues := [...]string{data.Ipv6ExportRouteTargets[i].RouteTarget.ValueString()}
		keyValuesVariables := [...]string{data.Ipv6ExportRouteTargets[i].RouteTargetVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "mplsVpnIpv6RouteTarget.exportRtList").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType").String()
					vv := v.Get(keys[ik] + ".value").String()
					if (tt == "variable" && vv == keyValuesVariables[ik]) || (tt == "global" && vv == keyValues[ik]) {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Ipv6ExportRouteTargets[i].RouteTarget = types.StringNull()
		data.Ipv6ExportRouteTargets[i].RouteTargetVariable = types.StringNull()
		if t := r.Get("rt.optionType"); t.Exists() {
			va := r.Get("rt.value")
			if t.String() == "variable" {
				data.Ipv6ExportRouteTargets[i].RouteTargetVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6ExportRouteTargets[i].RouteTarget = types.StringValue(va.String())
			}
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *ServiceLANVPN) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.FeatureProfileId.IsNull() {
		return false
	}
	if !data.Vpn.IsNull() {
		return false
	}
	if !data.VpnVariable.IsNull() {
		return false
	}
	if !data.ConfigDescription.IsNull() {
		return false
	}
	if !data.ConfigDescriptionVariable.IsNull() {
		return false
	}
	if !data.OmpAdminDistanceIpv4.IsNull() {
		return false
	}
	if !data.OmpAdminDistanceIpv4Variable.IsNull() {
		return false
	}
	if !data.OmpAdminDistanceIpv6.IsNull() {
		return false
	}
	if !data.OmpAdminDistanceIpv6Variable.IsNull() {
		return false
	}
	if !data.EnableSdwanRemoteAccess.IsNull() {
		return false
	}
	if !data.PrimaryDnsAddressIpv4.IsNull() {
		return false
	}
	if !data.PrimaryDnsAddressIpv4Variable.IsNull() {
		return false
	}
	if !data.SecondaryDnsAddressIpv4.IsNull() {
		return false
	}
	if !data.SecondaryDnsAddressIpv4Variable.IsNull() {
		return false
	}
	if !data.PrimaryDnsAddressIpv6.IsNull() {
		return false
	}
	if !data.PrimaryDnsAddressIpv6Variable.IsNull() {
		return false
	}
	if !data.SecondaryDnsAddressIpv6.IsNull() {
		return false
	}
	if !data.SecondaryDnsAddressIpv6Variable.IsNull() {
		return false
	}
	if len(data.HostMappings) > 0 {
		return false
	}
	if len(data.AdvertiseOmpIpv4s) > 0 {
		return false
	}
	if len(data.AdvertiseOmpIpv6s) > 0 {
		return false
	}
	if len(data.Ipv4StaticRoutes) > 0 {
		return false
	}
	if len(data.Ipv6StaticRoutes) > 0 {
		return false
	}
	if len(data.Services) > 0 {
		return false
	}
	if len(data.ServiceRoutes) > 0 {
		return false
	}
	if len(data.GreRoutes) > 0 {
		return false
	}
	if len(data.IpsecRoutes) > 0 {
		return false
	}
	if len(data.NatPools) > 0 {
		return false
	}
	if len(data.NatPortForwards) > 0 {
		return false
	}
	if len(data.StaticNats) > 0 {
		return false
	}
	if len(data.Nat64V4Pools) > 0 {
		return false
	}
	if len(data.RouteLeakFromGlobalVpns) > 0 {
		return false
	}
	if len(data.RouteLeakToGlobalVpns) > 0 {
		return false
	}
	if len(data.RouteLeakFromOtherServices) > 0 {
		return false
	}
	if len(data.Ipv4ImportRouteTargets) > 0 {
		return false
	}
	if len(data.Ipv4ExportRouteTargets) > 0 {
		return false
	}
	if len(data.Ipv6ImportRouteTargets) > 0 {
		return false
	}
	if len(data.Ipv6ExportRouteTargets) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
