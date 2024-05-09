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
type CiscoBGP struct {
	Id                         types.String               `tfsdk:"id"`
	Version                    types.Int64                `tfsdk:"version"`
	TemplateType               types.String               `tfsdk:"template_type"`
	Name                       types.String               `tfsdk:"name"`
	Description                types.String               `tfsdk:"description"`
	DeviceTypes                types.Set                  `tfsdk:"device_types"`
	AsNumber                   types.String               `tfsdk:"as_number"`
	AsNumberVariable           types.String               `tfsdk:"as_number_variable"`
	Shutdown                   types.Bool                 `tfsdk:"shutdown"`
	ShutdownVariable           types.String               `tfsdk:"shutdown_variable"`
	RouterId                   types.String               `tfsdk:"router_id"`
	RouterIdVariable           types.String               `tfsdk:"router_id_variable"`
	PropagateAspath            types.Bool                 `tfsdk:"propagate_aspath"`
	PropagateAspathVariable    types.String               `tfsdk:"propagate_aspath_variable"`
	PropagateCommunity         types.Bool                 `tfsdk:"propagate_community"`
	PropagateCommunityVariable types.String               `tfsdk:"propagate_community_variable"`
	Ipv4RouteTargets           []CiscoBGPIpv4RouteTargets `tfsdk:"ipv4_route_targets"`
	Ipv6RouteTargets           []CiscoBGPIpv6RouteTargets `tfsdk:"ipv6_route_targets"`
	MplsInterfaces             []CiscoBGPMplsInterfaces   `tfsdk:"mpls_interfaces"`
	DistanceExternal           types.Int64                `tfsdk:"distance_external"`
	DistanceExternalVariable   types.String               `tfsdk:"distance_external_variable"`
	DistanceInternal           types.Int64                `tfsdk:"distance_internal"`
	DistanceInternalVariable   types.String               `tfsdk:"distance_internal_variable"`
	DistanceLocal              types.Int64                `tfsdk:"distance_local"`
	DistanceLocalVariable      types.String               `tfsdk:"distance_local_variable"`
	Keepalive                  types.Int64                `tfsdk:"keepalive"`
	KeepaliveVariable          types.String               `tfsdk:"keepalive_variable"`
	Holdtime                   types.Int64                `tfsdk:"holdtime"`
	HoldtimeVariable           types.String               `tfsdk:"holdtime_variable"`
	AlwaysCompareMed           types.Bool                 `tfsdk:"always_compare_med"`
	AlwaysCompareMedVariable   types.String               `tfsdk:"always_compare_med_variable"`
	DeterministicMed           types.Bool                 `tfsdk:"deterministic_med"`
	DeterministicMedVariable   types.String               `tfsdk:"deterministic_med_variable"`
	MissingMedWorst            types.Bool                 `tfsdk:"missing_med_worst"`
	MissingMedWorstVariable    types.String               `tfsdk:"missing_med_worst_variable"`
	CompareRouterId            types.Bool                 `tfsdk:"compare_router_id"`
	CompareRouterIdVariable    types.String               `tfsdk:"compare_router_id_variable"`
	MultipathRelax             types.Bool                 `tfsdk:"multipath_relax"`
	MultipathRelaxVariable     types.String               `tfsdk:"multipath_relax_variable"`
	AddressFamilies            []CiscoBGPAddressFamilies  `tfsdk:"address_families"`
	Ipv4Neighbors              []CiscoBGPIpv4Neighbors    `tfsdk:"ipv4_neighbors"`
	Ipv6Neighbors              []CiscoBGPIpv6Neighbors    `tfsdk:"ipv6_neighbors"`
}

type CiscoBGPIpv4RouteTargets struct {
	Optional      types.Bool                       `tfsdk:"optional"`
	VpnId         types.Int64                      `tfsdk:"vpn_id"`
	VpnIdVariable types.String                     `tfsdk:"vpn_id_variable"`
	Export        []CiscoBGPIpv4RouteTargetsExport `tfsdk:"export"`
	Import        []CiscoBGPIpv4RouteTargetsImport `tfsdk:"import"`
}

type CiscoBGPIpv6RouteTargets struct {
	Optional      types.Bool                       `tfsdk:"optional"`
	VpnId         types.Int64                      `tfsdk:"vpn_id"`
	VpnIdVariable types.String                     `tfsdk:"vpn_id_variable"`
	Export        []CiscoBGPIpv6RouteTargetsExport `tfsdk:"export"`
	Import        []CiscoBGPIpv6RouteTargetsImport `tfsdk:"import"`
}

type CiscoBGPMplsInterfaces struct {
	Optional              types.Bool   `tfsdk:"optional"`
	InterfaceName         types.String `tfsdk:"interface_name"`
	InterfaceNameVariable types.String `tfsdk:"interface_name_variable"`
}

type CiscoBGPAddressFamilies struct {
	Optional                            types.Bool                                      `tfsdk:"optional"`
	FamilyType                          types.String                                    `tfsdk:"family_type"`
	Ipv4AggregateAddresses              []CiscoBGPAddressFamiliesIpv4AggregateAddresses `tfsdk:"ipv4_aggregate_addresses"`
	Ipv6AggregateAddresses              []CiscoBGPAddressFamiliesIpv6AggregateAddresses `tfsdk:"ipv6_aggregate_addresses"`
	Ipv4Networks                        []CiscoBGPAddressFamiliesIpv4Networks           `tfsdk:"ipv4_networks"`
	Ipv6Networks                        []CiscoBGPAddressFamiliesIpv6Networks           `tfsdk:"ipv6_networks"`
	MaximumPaths                        types.Int64                                     `tfsdk:"maximum_paths"`
	MaximumPathsVariable                types.String                                    `tfsdk:"maximum_paths_variable"`
	DefaultInformationOriginate         types.Bool                                      `tfsdk:"default_information_originate"`
	DefaultInformationOriginateVariable types.String                                    `tfsdk:"default_information_originate_variable"`
	TableMapPolicy                      types.String                                    `tfsdk:"table_map_policy"`
	TableMapPolicyVariable              types.String                                    `tfsdk:"table_map_policy_variable"`
	TableMapFilter                      types.Bool                                      `tfsdk:"table_map_filter"`
	TableMapFilterVariable              types.String                                    `tfsdk:"table_map_filter_variable"`
	RedistributeRoutes                  []CiscoBGPAddressFamiliesRedistributeRoutes     `tfsdk:"redistribute_routes"`
}

type CiscoBGPIpv4Neighbors struct {
	Optional                  types.Bool                             `tfsdk:"optional"`
	Address                   types.String                           `tfsdk:"address"`
	AddressVariable           types.String                           `tfsdk:"address_variable"`
	Description               types.String                           `tfsdk:"description"`
	DescriptionVariable       types.String                           `tfsdk:"description_variable"`
	Shutdown                  types.Bool                             `tfsdk:"shutdown"`
	ShutdownVariable          types.String                           `tfsdk:"shutdown_variable"`
	RemoteAs                  types.String                           `tfsdk:"remote_as"`
	RemoteAsVariable          types.String                           `tfsdk:"remote_as_variable"`
	Keepalive                 types.Int64                            `tfsdk:"keepalive"`
	KeepaliveVariable         types.String                           `tfsdk:"keepalive_variable"`
	Holdtime                  types.Int64                            `tfsdk:"holdtime"`
	HoldtimeVariable          types.String                           `tfsdk:"holdtime_variable"`
	SourceInterface           types.String                           `tfsdk:"source_interface"`
	SourceInterfaceVariable   types.String                           `tfsdk:"source_interface_variable"`
	NextHopSelf               types.Bool                             `tfsdk:"next_hop_self"`
	NextHopSelfVariable       types.String                           `tfsdk:"next_hop_self_variable"`
	SendCommunity             types.Bool                             `tfsdk:"send_community"`
	SendCommunityVariable     types.String                           `tfsdk:"send_community_variable"`
	SendExtCommunity          types.Bool                             `tfsdk:"send_ext_community"`
	SendExtCommunityVariable  types.String                           `tfsdk:"send_ext_community_variable"`
	EbgpMultihop              types.Int64                            `tfsdk:"ebgp_multihop"`
	EbgpMultihopVariable      types.String                           `tfsdk:"ebgp_multihop_variable"`
	Password                  types.String                           `tfsdk:"password"`
	PasswordVariable          types.String                           `tfsdk:"password_variable"`
	SendLabel                 types.Bool                             `tfsdk:"send_label"`
	SendLabelVariable         types.String                           `tfsdk:"send_label_variable"`
	SendLabelExplicit         types.Bool                             `tfsdk:"send_label_explicit"`
	SendLabelExplicitVariable types.String                           `tfsdk:"send_label_explicit_variable"`
	AsOverride                types.Bool                             `tfsdk:"as_override"`
	AsOverrideVariable        types.String                           `tfsdk:"as_override_variable"`
	AllowAsIn                 types.Int64                            `tfsdk:"allow_as_in"`
	AllowAsInVariable         types.String                           `tfsdk:"allow_as_in_variable"`
	AddressFamilies           []CiscoBGPIpv4NeighborsAddressFamilies `tfsdk:"address_families"`
}

type CiscoBGPIpv6Neighbors struct {
	Optional                  types.Bool                             `tfsdk:"optional"`
	Address                   types.String                           `tfsdk:"address"`
	AddressVariable           types.String                           `tfsdk:"address_variable"`
	Description               types.String                           `tfsdk:"description"`
	DescriptionVariable       types.String                           `tfsdk:"description_variable"`
	Shutdown                  types.Bool                             `tfsdk:"shutdown"`
	ShutdownVariable          types.String                           `tfsdk:"shutdown_variable"`
	RemoteAs                  types.String                           `tfsdk:"remote_as"`
	RemoteAsVariable          types.String                           `tfsdk:"remote_as_variable"`
	Keepalive                 types.Int64                            `tfsdk:"keepalive"`
	KeepaliveVariable         types.String                           `tfsdk:"keepalive_variable"`
	Holdtime                  types.Int64                            `tfsdk:"holdtime"`
	HoldtimeVariable          types.String                           `tfsdk:"holdtime_variable"`
	SourceInterface           types.String                           `tfsdk:"source_interface"`
	SourceInterfaceVariable   types.String                           `tfsdk:"source_interface_variable"`
	NextHopSelf               types.Bool                             `tfsdk:"next_hop_self"`
	NextHopSelfVariable       types.String                           `tfsdk:"next_hop_self_variable"`
	SendCommunity             types.Bool                             `tfsdk:"send_community"`
	SendCommunityVariable     types.String                           `tfsdk:"send_community_variable"`
	SendExtCommunity          types.Bool                             `tfsdk:"send_ext_community"`
	SendExtCommunityVariable  types.String                           `tfsdk:"send_ext_community_variable"`
	EbgpMultihop              types.Int64                            `tfsdk:"ebgp_multihop"`
	EbgpMultihopVariable      types.String                           `tfsdk:"ebgp_multihop_variable"`
	Password                  types.String                           `tfsdk:"password"`
	PasswordVariable          types.String                           `tfsdk:"password_variable"`
	SendLabel                 types.Bool                             `tfsdk:"send_label"`
	SendLabelVariable         types.String                           `tfsdk:"send_label_variable"`
	SendLabelExplicit         types.Bool                             `tfsdk:"send_label_explicit"`
	SendLabelExplicitVariable types.String                           `tfsdk:"send_label_explicit_variable"`
	AsOverride                types.Bool                             `tfsdk:"as_override"`
	AsOverrideVariable        types.String                           `tfsdk:"as_override_variable"`
	AllowAsIn                 types.Int64                            `tfsdk:"allow_as_in"`
	AllowAsInVariable         types.String                           `tfsdk:"allow_as_in_variable"`
	AddressFamilies           []CiscoBGPIpv6NeighborsAddressFamilies `tfsdk:"address_families"`
}

type CiscoBGPIpv4RouteTargetsExport struct {
	Optional      types.Bool   `tfsdk:"optional"`
	AsnIp         types.String `tfsdk:"asn_ip"`
	AsnIpVariable types.String `tfsdk:"asn_ip_variable"`
}
type CiscoBGPIpv4RouteTargetsImport struct {
	Optional      types.Bool   `tfsdk:"optional"`
	AsnIp         types.String `tfsdk:"asn_ip"`
	AsnIpVariable types.String `tfsdk:"asn_ip_variable"`
}

type CiscoBGPIpv6RouteTargetsExport struct {
	Optional      types.Bool   `tfsdk:"optional"`
	AsnIp         types.String `tfsdk:"asn_ip"`
	AsnIpVariable types.String `tfsdk:"asn_ip_variable"`
}
type CiscoBGPIpv6RouteTargetsImport struct {
	Optional      types.Bool   `tfsdk:"optional"`
	AsnIp         types.String `tfsdk:"asn_ip"`
	AsnIpVariable types.String `tfsdk:"asn_ip_variable"`
}

type CiscoBGPAddressFamiliesIpv4AggregateAddresses struct {
	Optional            types.Bool   `tfsdk:"optional"`
	Prefix              types.String `tfsdk:"prefix"`
	PrefixVariable      types.String `tfsdk:"prefix_variable"`
	AsSetPath           types.Bool   `tfsdk:"as_set_path"`
	AsSetPathVariable   types.String `tfsdk:"as_set_path_variable"`
	SummaryOnly         types.Bool   `tfsdk:"summary_only"`
	SummaryOnlyVariable types.String `tfsdk:"summary_only_variable"`
}
type CiscoBGPAddressFamiliesIpv6AggregateAddresses struct {
	Optional            types.Bool   `tfsdk:"optional"`
	Prefix              types.String `tfsdk:"prefix"`
	PrefixVariable      types.String `tfsdk:"prefix_variable"`
	AsSetPath           types.Bool   `tfsdk:"as_set_path"`
	AsSetPathVariable   types.String `tfsdk:"as_set_path_variable"`
	SummaryOnly         types.Bool   `tfsdk:"summary_only"`
	SummaryOnlyVariable types.String `tfsdk:"summary_only_variable"`
}
type CiscoBGPAddressFamiliesIpv4Networks struct {
	Optional       types.Bool   `tfsdk:"optional"`
	Prefix         types.String `tfsdk:"prefix"`
	PrefixVariable types.String `tfsdk:"prefix_variable"`
}
type CiscoBGPAddressFamiliesIpv6Networks struct {
	Optional       types.Bool   `tfsdk:"optional"`
	Prefix         types.String `tfsdk:"prefix"`
	PrefixVariable types.String `tfsdk:"prefix_variable"`
}
type CiscoBGPAddressFamiliesRedistributeRoutes struct {
	Optional            types.Bool   `tfsdk:"optional"`
	Protocol            types.String `tfsdk:"protocol"`
	ProtocolVariable    types.String `tfsdk:"protocol_variable"`
	RoutePolicy         types.String `tfsdk:"route_policy"`
	RoutePolicyVariable types.String `tfsdk:"route_policy_variable"`
}

type CiscoBGPIpv4NeighborsAddressFamilies struct {
	Optional                           types.Bool                                          `tfsdk:"optional"`
	FamilyType                         types.String                                        `tfsdk:"family_type"`
	MaximumPrefixes                    types.Int64                                         `tfsdk:"maximum_prefixes"`
	MaximumPrefixesVariable            types.String                                        `tfsdk:"maximum_prefixes_variable"`
	MaximumPrefixesThreshold           types.Int64                                         `tfsdk:"maximum_prefixes_threshold"`
	MaximumPrefixesThresholdVariable   types.String                                        `tfsdk:"maximum_prefixes_threshold_variable"`
	MaximumPrefixesRestart             types.Int64                                         `tfsdk:"maximum_prefixes_restart"`
	MaximumPrefixesRestartVariable     types.String                                        `tfsdk:"maximum_prefixes_restart_variable"`
	MaximumPrefixesWarningOnly         types.Bool                                          `tfsdk:"maximum_prefixes_warning_only"`
	MaximumPrefixesWarningOnlyVariable types.String                                        `tfsdk:"maximum_prefixes_warning_only_variable"`
	RoutePolicies                      []CiscoBGPIpv4NeighborsAddressFamiliesRoutePolicies `tfsdk:"route_policies"`
}

type CiscoBGPIpv6NeighborsAddressFamilies struct {
	Optional                           types.Bool                                          `tfsdk:"optional"`
	FamilyType                         types.String                                        `tfsdk:"family_type"`
	MaximumPrefixes                    types.Int64                                         `tfsdk:"maximum_prefixes"`
	MaximumPrefixesVariable            types.String                                        `tfsdk:"maximum_prefixes_variable"`
	MaximumPrefixesThreshold           types.Int64                                         `tfsdk:"maximum_prefixes_threshold"`
	MaximumPrefixesThresholdVariable   types.String                                        `tfsdk:"maximum_prefixes_threshold_variable"`
	MaximumPrefixesRestart             types.Int64                                         `tfsdk:"maximum_prefixes_restart"`
	MaximumPrefixesRestartVariable     types.String                                        `tfsdk:"maximum_prefixes_restart_variable"`
	MaximumPrefixesWarningOnly         types.Bool                                          `tfsdk:"maximum_prefixes_warning_only"`
	MaximumPrefixesWarningOnlyVariable types.String                                        `tfsdk:"maximum_prefixes_warning_only_variable"`
	RoutePolicies                      []CiscoBGPIpv6NeighborsAddressFamiliesRoutePolicies `tfsdk:"route_policies"`
}

type CiscoBGPIpv4NeighborsAddressFamiliesRoutePolicies struct {
	Optional           types.Bool   `tfsdk:"optional"`
	Direction          types.String `tfsdk:"direction"`
	PolicyName         types.String `tfsdk:"policy_name"`
	PolicyNameVariable types.String `tfsdk:"policy_name_variable"`
}

type CiscoBGPIpv6NeighborsAddressFamiliesRoutePolicies struct {
	Optional           types.Bool   `tfsdk:"optional"`
	Direction          types.String `tfsdk:"direction"`
	PolicyName         types.String `tfsdk:"policy_name"`
	PolicyNameVariable types.String `tfsdk:"policy_name_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoBGP) getModel() string {
	return "cisco_bgp"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data CiscoBGP) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_bgp")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.AsNumberVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.as-num."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.as-num."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bgp.as-num."+"vipVariableName", data.AsNumberVariable.ValueString())
	} else if data.AsNumber.IsNull() {
		if !gjson.Get(body, path+"bgp").Exists() {
			body, _ = sjson.Set(body, path+"bgp", map[string]interface{}{})
		}
	} else {
		body, _ = sjson.Set(body, path+"bgp.as-num."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.as-num."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.as-num."+"vipValue", data.AsNumber.ValueString())
	}

	if !data.ShutdownVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.shutdown."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bgp.shutdown."+"vipVariableName", data.ShutdownVariable.ValueString())
	} else if data.Shutdown.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.shutdown."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bgp.shutdown."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.shutdown."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.shutdown."+"vipValue", strconv.FormatBool(data.Shutdown.ValueBool()))
	}

	if !data.RouterIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.router-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.router-id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bgp.router-id."+"vipVariableName", data.RouterIdVariable.ValueString())
	} else if data.RouterId.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.router-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.router-id."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bgp.router-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.router-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.router-id."+"vipValue", data.RouterId.ValueString())
	}

	if !data.PropagateAspathVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.propagate-aspath."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.propagate-aspath."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bgp.propagate-aspath."+"vipVariableName", data.PropagateAspathVariable.ValueString())
	} else if data.PropagateAspath.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.propagate-aspath."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.propagate-aspath."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bgp.propagate-aspath."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.propagate-aspath."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.propagate-aspath."+"vipValue", strconv.FormatBool(data.PropagateAspath.ValueBool()))
	}

	if !data.PropagateCommunityVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.propagate-community."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.propagate-community."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bgp.propagate-community."+"vipVariableName", data.PropagateCommunityVariable.ValueString())
	} else if data.PropagateCommunity.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.propagate-community."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.propagate-community."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bgp.propagate-community."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.propagate-community."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.propagate-community."+"vipValue", strconv.FormatBool(data.PropagateCommunity.ValueBool()))
	}
	if len(data.Ipv4RouteTargets) > 0 {
		body, _ = sjson.Set(body, path+"bgp.target.route-target-ipv4."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"bgp.target.route-target-ipv4."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.target.route-target-ipv4."+"vipPrimaryKey", []string{"vpn-id"})
		body, _ = sjson.Set(body, path+"bgp.target.route-target-ipv4."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"bgp.target.route-target-ipv4."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"bgp.target.route-target-ipv4."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"bgp.target.route-target-ipv4."+"vipPrimaryKey", []string{"vpn-id"})
		body, _ = sjson.Set(body, path+"bgp.target.route-target-ipv4."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv4RouteTargets {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "vpn-id")

		if !item.VpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipVariableName", item.VpnIdVariable.ValueString())
		} else if item.VpnId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipValue", item.VpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "export")
		if len(item.Export) > 0 {
			itemBody, _ = sjson.Set(itemBody, "export."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "export."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "export."+"vipPrimaryKey", []string{"asn-ip"})
			itemBody, _ = sjson.Set(itemBody, "export."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "export."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "export."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "export."+"vipPrimaryKey", []string{"asn-ip"})
			itemBody, _ = sjson.Set(itemBody, "export."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Export {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "asn-ip")

			if !childItem.AsnIpVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipVariableName", childItem.AsnIpVariable.ValueString())
			} else if childItem.AsnIp.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipValue", childItem.AsnIp.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "export."+"vipValue.-1", itemChildBody)
		}
		itemAttributes = append(itemAttributes, "import")
		if len(item.Import) > 0 {
			itemBody, _ = sjson.Set(itemBody, "import."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "import."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "import."+"vipPrimaryKey", []string{"asn-ip"})
			itemBody, _ = sjson.Set(itemBody, "import."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "import."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "import."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "import."+"vipPrimaryKey", []string{"asn-ip"})
			itemBody, _ = sjson.Set(itemBody, "import."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Import {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "asn-ip")

			if !childItem.AsnIpVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipVariableName", childItem.AsnIpVariable.ValueString())
			} else if childItem.AsnIp.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipValue", childItem.AsnIp.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "import."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"bgp.target.route-target-ipv4."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv6RouteTargets) > 0 {
		body, _ = sjson.Set(body, path+"bgp.target.route-target-ipv6."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"bgp.target.route-target-ipv6."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.target.route-target-ipv6."+"vipPrimaryKey", []string{"vpn-id"})
		body, _ = sjson.Set(body, path+"bgp.target.route-target-ipv6."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"bgp.target.route-target-ipv6."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"bgp.target.route-target-ipv6."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"bgp.target.route-target-ipv6."+"vipPrimaryKey", []string{"vpn-id"})
		body, _ = sjson.Set(body, path+"bgp.target.route-target-ipv6."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv6RouteTargets {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "vpn-id")

		if !item.VpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipVariableName", item.VpnIdVariable.ValueString())
		} else if item.VpnId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn-id."+"vipValue", item.VpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "export")
		if len(item.Export) > 0 {
			itemBody, _ = sjson.Set(itemBody, "export."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "export."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "export."+"vipPrimaryKey", []string{"asn-ip"})
			itemBody, _ = sjson.Set(itemBody, "export."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "export."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "export."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "export."+"vipPrimaryKey", []string{"asn-ip"})
			itemBody, _ = sjson.Set(itemBody, "export."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Export {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "asn-ip")

			if !childItem.AsnIpVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipVariableName", childItem.AsnIpVariable.ValueString())
			} else if childItem.AsnIp.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipValue", childItem.AsnIp.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "export."+"vipValue.-1", itemChildBody)
		}
		itemAttributes = append(itemAttributes, "import")
		if len(item.Import) > 0 {
			itemBody, _ = sjson.Set(itemBody, "import."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "import."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "import."+"vipPrimaryKey", []string{"asn-ip"})
			itemBody, _ = sjson.Set(itemBody, "import."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "import."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "import."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "import."+"vipPrimaryKey", []string{"asn-ip"})
			itemBody, _ = sjson.Set(itemBody, "import."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.Import {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "asn-ip")

			if !childItem.AsnIpVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipVariableName", childItem.AsnIpVariable.ValueString())
			} else if childItem.AsnIp.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "asn-ip."+"vipValue", childItem.AsnIp.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "import."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"bgp.target.route-target-ipv6."+"vipValue.-1", itemBody)
	}
	if len(data.MplsInterfaces) > 0 {
		body, _ = sjson.Set(body, path+"bgp.mpls-interface."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"bgp.mpls-interface."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.mpls-interface."+"vipPrimaryKey", []string{"if-name"})
		body, _ = sjson.Set(body, path+"bgp.mpls-interface."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"bgp.mpls-interface."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"bgp.mpls-interface."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"bgp.mpls-interface."+"vipPrimaryKey", []string{"if-name"})
		body, _ = sjson.Set(body, path+"bgp.mpls-interface."+"vipValue", []interface{}{})
	}
	for _, item := range data.MplsInterfaces {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "if-name")

		if !item.InterfaceNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipVariableName", item.InterfaceNameVariable.ValueString())
		} else if item.InterfaceName.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "if-name."+"vipValue", item.InterfaceName.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"bgp.mpls-interface."+"vipValue.-1", itemBody)
	}

	if !data.DistanceExternalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.distance.external."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.distance.external."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bgp.distance.external."+"vipVariableName", data.DistanceExternalVariable.ValueString())
	} else if data.DistanceExternal.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.distance.external."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.distance.external."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bgp.distance.external."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.distance.external."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.distance.external."+"vipValue", data.DistanceExternal.ValueInt64())
	}

	if !data.DistanceInternalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.distance.internal."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.distance.internal."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bgp.distance.internal."+"vipVariableName", data.DistanceInternalVariable.ValueString())
	} else if data.DistanceInternal.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.distance.internal."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.distance.internal."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bgp.distance.internal."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.distance.internal."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.distance.internal."+"vipValue", data.DistanceInternal.ValueInt64())
	}

	if !data.DistanceLocalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.distance.local."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.distance.local."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bgp.distance.local."+"vipVariableName", data.DistanceLocalVariable.ValueString())
	} else if data.DistanceLocal.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.distance.local."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.distance.local."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bgp.distance.local."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.distance.local."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.distance.local."+"vipValue", data.DistanceLocal.ValueInt64())
	}

	if !data.KeepaliveVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.timers.keepalive."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.timers.keepalive."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bgp.timers.keepalive."+"vipVariableName", data.KeepaliveVariable.ValueString())
	} else if data.Keepalive.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.timers.keepalive."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.timers.keepalive."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bgp.timers.keepalive."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.timers.keepalive."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.timers.keepalive."+"vipValue", data.Keepalive.ValueInt64())
	}

	if !data.HoldtimeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.timers.holdtime."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.timers.holdtime."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bgp.timers.holdtime."+"vipVariableName", data.HoldtimeVariable.ValueString())
	} else if data.Holdtime.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.timers.holdtime."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.timers.holdtime."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bgp.timers.holdtime."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.timers.holdtime."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.timers.holdtime."+"vipValue", data.Holdtime.ValueInt64())
	}

	if !data.AlwaysCompareMedVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.best-path.med.always-compare."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.best-path.med.always-compare."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bgp.best-path.med.always-compare."+"vipVariableName", data.AlwaysCompareMedVariable.ValueString())
	} else if data.AlwaysCompareMed.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.best-path.med.always-compare."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.best-path.med.always-compare."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bgp.best-path.med.always-compare."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.best-path.med.always-compare."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.best-path.med.always-compare."+"vipValue", strconv.FormatBool(data.AlwaysCompareMed.ValueBool()))
	}

	if !data.DeterministicMedVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.best-path.med.deterministic."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.best-path.med.deterministic."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bgp.best-path.med.deterministic."+"vipVariableName", data.DeterministicMedVariable.ValueString())
	} else if data.DeterministicMed.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.best-path.med.deterministic."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.best-path.med.deterministic."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bgp.best-path.med.deterministic."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.best-path.med.deterministic."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.best-path.med.deterministic."+"vipValue", strconv.FormatBool(data.DeterministicMed.ValueBool()))
	}

	if !data.MissingMedWorstVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.best-path.med.missing-as-worst."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.best-path.med.missing-as-worst."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bgp.best-path.med.missing-as-worst."+"vipVariableName", data.MissingMedWorstVariable.ValueString())
	} else if data.MissingMedWorst.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.best-path.med.missing-as-worst."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.best-path.med.missing-as-worst."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bgp.best-path.med.missing-as-worst."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.best-path.med.missing-as-worst."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.best-path.med.missing-as-worst."+"vipValue", strconv.FormatBool(data.MissingMedWorst.ValueBool()))
	}

	if !data.CompareRouterIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.best-path.compare-router-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.best-path.compare-router-id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bgp.best-path.compare-router-id."+"vipVariableName", data.CompareRouterIdVariable.ValueString())
	} else if data.CompareRouterId.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.best-path.compare-router-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.best-path.compare-router-id."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bgp.best-path.compare-router-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.best-path.compare-router-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.best-path.compare-router-id."+"vipValue", strconv.FormatBool(data.CompareRouterId.ValueBool()))
	}

	if !data.MultipathRelaxVariable.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.best-path.as-path.multipath-relax."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.best-path.as-path.multipath-relax."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"bgp.best-path.as-path.multipath-relax."+"vipVariableName", data.MultipathRelaxVariable.ValueString())
	} else if data.MultipathRelax.IsNull() {
		body, _ = sjson.Set(body, path+"bgp.best-path.as-path.multipath-relax."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.best-path.as-path.multipath-relax."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"bgp.best-path.as-path.multipath-relax."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"bgp.best-path.as-path.multipath-relax."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.best-path.as-path.multipath-relax."+"vipValue", strconv.FormatBool(data.MultipathRelax.ValueBool()))
	}
	if len(data.AddressFamilies) > 0 {
		body, _ = sjson.Set(body, path+"bgp.address-family."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"bgp.address-family."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.address-family."+"vipPrimaryKey", []string{"family-type"})
		body, _ = sjson.Set(body, path+"bgp.address-family."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"bgp.address-family."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"bgp.address-family."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"bgp.address-family."+"vipPrimaryKey", []string{"family-type"})
		body, _ = sjson.Set(body, path+"bgp.address-family."+"vipValue", []interface{}{})
	}
	for _, item := range data.AddressFamilies {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "family-type")
		if item.FamilyType.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "family-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "family-type."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "family-type."+"vipValue", item.FamilyType.ValueString())
		}
		itemAttributes = append(itemAttributes, "aggregate-address")
		if len(item.Ipv4AggregateAddresses) > 0 {
			itemBody, _ = sjson.Set(itemBody, "aggregate-address."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "aggregate-address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "aggregate-address."+"vipPrimaryKey", []string{"prefix"})
			itemBody, _ = sjson.Set(itemBody, "aggregate-address."+"vipValue", []interface{}{})
		} else {
		}
		for _, childItem := range item.Ipv4AggregateAddresses {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "prefix")

			if !childItem.PrefixVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipVariableName", childItem.PrefixVariable.ValueString())
			} else if childItem.Prefix.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipValue", childItem.Prefix.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "as-set")

			if !childItem.AsSetPathVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "as-set."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "as-set."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "as-set."+"vipVariableName", childItem.AsSetPathVariable.ValueString())
			} else if childItem.AsSetPath.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "as-set."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "as-set."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "as-set."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "as-set."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "as-set."+"vipValue", strconv.FormatBool(childItem.AsSetPath.ValueBool()))
			}
			itemChildAttributes = append(itemChildAttributes, "summary-only")

			if !childItem.SummaryOnlyVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "summary-only."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "summary-only."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "summary-only."+"vipVariableName", childItem.SummaryOnlyVariable.ValueString())
			} else if childItem.SummaryOnly.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "summary-only."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "summary-only."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "summary-only."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "summary-only."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "summary-only."+"vipValue", strconv.FormatBool(childItem.SummaryOnly.ValueBool()))
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "aggregate-address."+"vipValue.-1", itemChildBody)
		}
		itemAttributes = append(itemAttributes, "ipv6-aggregate-address")
		if len(item.Ipv6AggregateAddresses) > 0 {
			itemBody, _ = sjson.Set(itemBody, "ipv6-aggregate-address."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "ipv6-aggregate-address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ipv6-aggregate-address."+"vipPrimaryKey", []string{"prefix"})
			itemBody, _ = sjson.Set(itemBody, "ipv6-aggregate-address."+"vipValue", []interface{}{})
		} else {
		}
		for _, childItem := range item.Ipv6AggregateAddresses {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "prefix")

			if !childItem.PrefixVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipVariableName", childItem.PrefixVariable.ValueString())
			} else if childItem.Prefix.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipValue", childItem.Prefix.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "as-set")

			if !childItem.AsSetPathVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "as-set."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "as-set."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "as-set."+"vipVariableName", childItem.AsSetPathVariable.ValueString())
			} else if childItem.AsSetPath.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "as-set."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "as-set."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "as-set."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "as-set."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "as-set."+"vipValue", strconv.FormatBool(childItem.AsSetPath.ValueBool()))
			}
			itemChildAttributes = append(itemChildAttributes, "summary-only")

			if !childItem.SummaryOnlyVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "summary-only."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "summary-only."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "summary-only."+"vipVariableName", childItem.SummaryOnlyVariable.ValueString())
			} else if childItem.SummaryOnly.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "summary-only."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "summary-only."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "summary-only."+"vipObjectType", "node-only")
				itemChildBody, _ = sjson.Set(itemChildBody, "summary-only."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "summary-only."+"vipValue", strconv.FormatBool(childItem.SummaryOnly.ValueBool()))
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "ipv6-aggregate-address."+"vipValue.-1", itemChildBody)
		}
		itemAttributes = append(itemAttributes, "network")
		if len(item.Ipv4Networks) > 0 {
			itemBody, _ = sjson.Set(itemBody, "network."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "network."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "network."+"vipPrimaryKey", []string{"prefix"})
			itemBody, _ = sjson.Set(itemBody, "network."+"vipValue", []interface{}{})
		} else {
		}
		for _, childItem := range item.Ipv4Networks {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "prefix")

			if !childItem.PrefixVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipVariableName", childItem.PrefixVariable.ValueString())
			} else if childItem.Prefix.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipValue", childItem.Prefix.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "network."+"vipValue.-1", itemChildBody)
		}
		itemAttributes = append(itemAttributes, "ipv6-network")
		if len(item.Ipv6Networks) > 0 {
			itemBody, _ = sjson.Set(itemBody, "ipv6-network."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "ipv6-network."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ipv6-network."+"vipPrimaryKey", []string{"prefix"})
			itemBody, _ = sjson.Set(itemBody, "ipv6-network."+"vipValue", []interface{}{})
		} else {
		}
		for _, childItem := range item.Ipv6Networks {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "prefix")

			if !childItem.PrefixVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipVariableName", childItem.PrefixVariable.ValueString())
			} else if childItem.Prefix.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "prefix."+"vipValue", childItem.Prefix.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "ipv6-network."+"vipValue.-1", itemChildBody)
		}
		itemAttributes = append(itemAttributes, "maximum-paths")

		if !item.MaximumPathsVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "maximum-paths.paths."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "maximum-paths.paths."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "maximum-paths.paths."+"vipVariableName", item.MaximumPathsVariable.ValueString())
		} else if item.MaximumPaths.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "maximum-paths.paths."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "maximum-paths.paths."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "maximum-paths.paths."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "maximum-paths.paths."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "maximum-paths.paths."+"vipValue", item.MaximumPaths.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "default-information")

		if !item.DefaultInformationOriginateVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "default-information.originate."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "default-information.originate."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "default-information.originate."+"vipVariableName", item.DefaultInformationOriginateVariable.ValueString())
		} else if item.DefaultInformationOriginate.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "default-information.originate."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "default-information.originate."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "default-information.originate."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "default-information.originate."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "default-information.originate."+"vipValue", strconv.FormatBool(item.DefaultInformationOriginate.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "table-map")

		if !item.TableMapPolicyVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "table-map.name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "table-map.name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "table-map.name."+"vipVariableName", item.TableMapPolicyVariable.ValueString())
		} else if item.TableMapPolicy.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "table-map.name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "table-map.name."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "table-map.name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "table-map.name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "table-map.name."+"vipValue", item.TableMapPolicy.ValueString())
		}
		itemAttributes = append(itemAttributes, "table-map")

		if !item.TableMapFilterVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "table-map.filter."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "table-map.filter."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "table-map.filter."+"vipVariableName", item.TableMapFilterVariable.ValueString())
		} else if item.TableMapFilter.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "table-map.filter."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "table-map.filter."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "table-map.filter."+"vipObjectType", "node-only")
			itemBody, _ = sjson.Set(itemBody, "table-map.filter."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "table-map.filter."+"vipValue", strconv.FormatBool(item.TableMapFilter.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "redistribute")
		if len(item.RedistributeRoutes) > 0 {
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipPrimaryKey", []string{"protocol"})
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipPrimaryKey", []string{"protocol"})
			itemBody, _ = sjson.Set(itemBody, "redistribute."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.RedistributeRoutes {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "protocol")

			if !childItem.ProtocolVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipVariableName", childItem.ProtocolVariable.ValueString())
			} else if childItem.Protocol.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "protocol."+"vipValue", childItem.Protocol.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "route-policy")

			if !childItem.RoutePolicyVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipVariableName", childItem.RoutePolicyVariable.ValueString())
			} else if childItem.RoutePolicy.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipValue", childItem.RoutePolicy.ValueString())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "redistribute."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"bgp.address-family."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv4Neighbors) > 0 {
		body, _ = sjson.Set(body, path+"bgp.neighbor."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"bgp.neighbor."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.neighbor."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"bgp.neighbor."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"bgp.neighbor."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"bgp.neighbor."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"bgp.neighbor."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"bgp.neighbor."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv4Neighbors {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "address")

		if !item.AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipVariableName", item.AddressVariable.ValueString())
		} else if item.Address.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipValue", item.Address.ValueString())
		}
		itemAttributes = append(itemAttributes, "description")

		if !item.DescriptionVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "description."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipVariableName", item.DescriptionVariable.ValueString())
		} else if item.Description.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "description."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "description."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipValue", item.Description.ValueString())
		}
		itemAttributes = append(itemAttributes, "shutdown")

		if !item.ShutdownVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipVariableName", item.ShutdownVariable.ValueString())
		} else if item.Shutdown.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipValue", strconv.FormatBool(item.Shutdown.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "remote-as")

		if !item.RemoteAsVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "remote-as."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "remote-as."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "remote-as."+"vipVariableName", item.RemoteAsVariable.ValueString())
		} else if item.RemoteAs.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "remote-as."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "remote-as."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "remote-as."+"vipValue", item.RemoteAs.ValueString())
		}
		itemAttributes = append(itemAttributes, "timers")

		if !item.KeepaliveVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "timers.keepalive."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timers.keepalive."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "timers.keepalive."+"vipVariableName", item.KeepaliveVariable.ValueString())
		} else if item.Keepalive.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "timers.keepalive."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timers.keepalive."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "timers.keepalive."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timers.keepalive."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "timers.keepalive."+"vipValue", item.Keepalive.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "timers")

		if !item.HoldtimeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "timers.holdtime."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timers.holdtime."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "timers.holdtime."+"vipVariableName", item.HoldtimeVariable.ValueString())
		} else if item.Holdtime.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "timers.holdtime."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timers.holdtime."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "timers.holdtime."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timers.holdtime."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "timers.holdtime."+"vipValue", item.Holdtime.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "update-source")

		if !item.SourceInterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "update-source.if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "update-source.if-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "update-source.if-name."+"vipVariableName", item.SourceInterfaceVariable.ValueString())
		} else if item.SourceInterface.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "update-source.if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "update-source.if-name."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "update-source.if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "update-source.if-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "update-source.if-name."+"vipValue", item.SourceInterface.ValueString())
		}
		itemAttributes = append(itemAttributes, "next-hop-self")

		if !item.NextHopSelfVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "next-hop-self."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "next-hop-self."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "next-hop-self."+"vipVariableName", item.NextHopSelfVariable.ValueString())
		} else if item.NextHopSelf.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "next-hop-self."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "next-hop-self."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "next-hop-self."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "next-hop-self."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "next-hop-self."+"vipValue", strconv.FormatBool(item.NextHopSelf.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "send-community")

		if !item.SendCommunityVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-community."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-community."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "send-community."+"vipVariableName", item.SendCommunityVariable.ValueString())
		} else if item.SendCommunity.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-community."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-community."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "send-community."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-community."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "send-community."+"vipValue", strconv.FormatBool(item.SendCommunity.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "send-ext-community")

		if !item.SendExtCommunityVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-ext-community."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-ext-community."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "send-ext-community."+"vipVariableName", item.SendExtCommunityVariable.ValueString())
		} else if item.SendExtCommunity.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-ext-community."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-ext-community."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "send-ext-community."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-ext-community."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "send-ext-community."+"vipValue", strconv.FormatBool(item.SendExtCommunity.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "ebgp-multihop")

		if !item.EbgpMultihopVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ebgp-multihop."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ebgp-multihop."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ebgp-multihop."+"vipVariableName", item.EbgpMultihopVariable.ValueString())
		} else if item.EbgpMultihop.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ebgp-multihop."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ebgp-multihop."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "ebgp-multihop."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ebgp-multihop."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ebgp-multihop."+"vipValue", item.EbgpMultihop.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "password")

		if !item.PasswordVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "password."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "password."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "password."+"vipVariableName", item.PasswordVariable.ValueString())
		} else if item.Password.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "password."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "password."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "password."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "password."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "password."+"vipValue", item.Password.ValueString())
		}
		itemAttributes = append(itemAttributes, "send-label")

		if !item.SendLabelVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-label."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-label."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "send-label."+"vipVariableName", item.SendLabelVariable.ValueString())
		} else if item.SendLabel.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-label."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-label."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "send-label."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-label."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "send-label."+"vipValue", strconv.FormatBool(item.SendLabel.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "send-label-explicit")

		if !item.SendLabelExplicitVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-label-explicit."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-label-explicit."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "send-label-explicit."+"vipVariableName", item.SendLabelExplicitVariable.ValueString())
		} else if item.SendLabelExplicit.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-label-explicit."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-label-explicit."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "send-label-explicit."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-label-explicit."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "send-label-explicit."+"vipValue", strconv.FormatBool(item.SendLabelExplicit.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "as-override")

		if !item.AsOverrideVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "as-override."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "as-override."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "as-override."+"vipVariableName", item.AsOverrideVariable.ValueString())
		} else if item.AsOverride.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "as-override."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "as-override."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "as-override."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "as-override."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "as-override."+"vipValue", strconv.FormatBool(item.AsOverride.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "allowas-in")

		if !item.AllowAsInVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "allowas-in.as-number."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "allowas-in.as-number."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "allowas-in.as-number."+"vipVariableName", item.AllowAsInVariable.ValueString())
		} else if item.AllowAsIn.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "allowas-in.as-number."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "allowas-in.as-number."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "allowas-in.as-number."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "allowas-in.as-number."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "allowas-in.as-number."+"vipValue", item.AllowAsIn.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "address-family")
		if len(item.AddressFamilies) > 0 {
			itemBody, _ = sjson.Set(itemBody, "address-family."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "address-family."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "address-family."+"vipPrimaryKey", []string{"family-type"})
			itemBody, _ = sjson.Set(itemBody, "address-family."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "address-family."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "address-family."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "address-family."+"vipPrimaryKey", []string{"family-type"})
			itemBody, _ = sjson.Set(itemBody, "address-family."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.AddressFamilies {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "family-type")
			if childItem.FamilyType.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "family-type."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "family-type."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "family-type."+"vipValue", childItem.FamilyType.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "maximum-prefixes")

			if !childItem.MaximumPrefixesVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.prefix-num."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.prefix-num."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.prefix-num."+"vipVariableName", childItem.MaximumPrefixesVariable.ValueString())
			} else if childItem.MaximumPrefixes.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.prefix-num."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.prefix-num."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.prefix-num."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.prefix-num."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.prefix-num."+"vipValue", childItem.MaximumPrefixes.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "maximum-prefixes")

			if !childItem.MaximumPrefixesThresholdVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.threshold."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.threshold."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.threshold."+"vipVariableName", childItem.MaximumPrefixesThresholdVariable.ValueString())
			} else if childItem.MaximumPrefixesThreshold.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.threshold."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.threshold."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.threshold."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.threshold."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.threshold."+"vipValue", childItem.MaximumPrefixesThreshold.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "maximum-prefixes")

			if !childItem.MaximumPrefixesRestartVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.restart."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.restart."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.restart."+"vipVariableName", childItem.MaximumPrefixesRestartVariable.ValueString())
			} else if childItem.MaximumPrefixesRestart.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.restart."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.restart."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.restart."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.restart."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.restart."+"vipValue", childItem.MaximumPrefixesRestart.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "maximum-prefixes")

			if !childItem.MaximumPrefixesWarningOnlyVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.warning-only."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.warning-only."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.warning-only."+"vipVariableName", childItem.MaximumPrefixesWarningOnlyVariable.ValueString())
			} else if childItem.MaximumPrefixesWarningOnly.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.warning-only."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.warning-only."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.warning-only."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.warning-only."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.warning-only."+"vipValue", strconv.FormatBool(childItem.MaximumPrefixesWarningOnly.ValueBool()))
			}
			itemChildAttributes = append(itemChildAttributes, "route-policy")
			if len(childItem.RoutePolicies) > 0 {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipObjectType", "tree")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipPrimaryKey", []string{"direction"})
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipValue", []interface{}{})
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipObjectType", "tree")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipType", "ignore")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipPrimaryKey", []string{"direction"})
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipValue", []interface{}{})
			}
			for _, childChildItem := range childItem.RoutePolicies {
				itemChildChildBody := ""
				itemChildChildAttributes := make([]string, 0)
				itemChildChildAttributes = append(itemChildChildAttributes, "direction")
				if childChildItem.Direction.IsNull() {
				} else {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "direction."+"vipObjectType", "object")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "direction."+"vipType", "constant")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "direction."+"vipValue", childChildItem.Direction.ValueString())
				}
				itemChildChildAttributes = append(itemChildChildAttributes, "pol-name")

				if !childChildItem.PolicyNameVariable.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "pol-name."+"vipObjectType", "object")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "pol-name."+"vipType", "variableName")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "pol-name."+"vipVariableName", childChildItem.PolicyNameVariable.ValueString())
				} else if childChildItem.PolicyName.IsNull() {
				} else {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "pol-name."+"vipObjectType", "object")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "pol-name."+"vipType", "constant")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "pol-name."+"vipValue", childChildItem.PolicyName.ValueString())
				}
				if !childChildItem.Optional.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "vipOptional", childChildItem.Optional.ValueBool())
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "priority-order", itemChildChildAttributes)
				}
				itemChildBody, _ = sjson.SetRaw(itemChildBody, "route-policy."+"vipValue.-1", itemChildChildBody)
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "address-family."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"bgp.neighbor."+"vipValue.-1", itemBody)
	}
	if len(data.Ipv6Neighbors) > 0 {
		body, _ = sjson.Set(body, path+"bgp.ipv6-neighbor."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"bgp.ipv6-neighbor."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"bgp.ipv6-neighbor."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"bgp.ipv6-neighbor."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"bgp.ipv6-neighbor."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"bgp.ipv6-neighbor."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"bgp.ipv6-neighbor."+"vipPrimaryKey", []string{"address"})
		body, _ = sjson.Set(body, path+"bgp.ipv6-neighbor."+"vipValue", []interface{}{})
	}
	for _, item := range data.Ipv6Neighbors {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "address")

		if !item.AddressVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipVariableName", item.AddressVariable.ValueString())
		} else if item.Address.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "address."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "address."+"vipValue", item.Address.ValueString())
		}
		itemAttributes = append(itemAttributes, "description")

		if !item.DescriptionVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "description."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipVariableName", item.DescriptionVariable.ValueString())
		} else if item.Description.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "description."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "description."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "description."+"vipValue", item.Description.ValueString())
		}
		itemAttributes = append(itemAttributes, "shutdown")

		if !item.ShutdownVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipVariableName", item.ShutdownVariable.ValueString())
		} else if item.Shutdown.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "shutdown."+"vipValue", strconv.FormatBool(item.Shutdown.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "remote-as")

		if !item.RemoteAsVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "remote-as."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "remote-as."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "remote-as."+"vipVariableName", item.RemoteAsVariable.ValueString())
		} else if item.RemoteAs.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "remote-as."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "remote-as."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "remote-as."+"vipValue", item.RemoteAs.ValueString())
		}
		itemAttributes = append(itemAttributes, "timers")

		if !item.KeepaliveVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "timers.keepalive."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timers.keepalive."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "timers.keepalive."+"vipVariableName", item.KeepaliveVariable.ValueString())
		} else if item.Keepalive.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "timers.keepalive."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timers.keepalive."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "timers.keepalive."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timers.keepalive."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "timers.keepalive."+"vipValue", item.Keepalive.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "timers")

		if !item.HoldtimeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "timers.holdtime."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timers.holdtime."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "timers.holdtime."+"vipVariableName", item.HoldtimeVariable.ValueString())
		} else if item.Holdtime.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "timers.holdtime."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timers.holdtime."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "timers.holdtime."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "timers.holdtime."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "timers.holdtime."+"vipValue", item.Holdtime.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "update-source")

		if !item.SourceInterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "update-source.if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "update-source.if-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "update-source.if-name."+"vipVariableName", item.SourceInterfaceVariable.ValueString())
		} else if item.SourceInterface.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "update-source.if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "update-source.if-name."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "update-source.if-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "update-source.if-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "update-source.if-name."+"vipValue", item.SourceInterface.ValueString())
		}
		itemAttributes = append(itemAttributes, "next-hop-self")

		if !item.NextHopSelfVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "next-hop-self."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "next-hop-self."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "next-hop-self."+"vipVariableName", item.NextHopSelfVariable.ValueString())
		} else if item.NextHopSelf.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "next-hop-self."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "next-hop-self."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "next-hop-self."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "next-hop-self."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "next-hop-self."+"vipValue", strconv.FormatBool(item.NextHopSelf.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "send-community")

		if !item.SendCommunityVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-community."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-community."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "send-community."+"vipVariableName", item.SendCommunityVariable.ValueString())
		} else if item.SendCommunity.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-community."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-community."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "send-community."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-community."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "send-community."+"vipValue", strconv.FormatBool(item.SendCommunity.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "send-ext-community")

		if !item.SendExtCommunityVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-ext-community."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-ext-community."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "send-ext-community."+"vipVariableName", item.SendExtCommunityVariable.ValueString())
		} else if item.SendExtCommunity.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-ext-community."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-ext-community."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "send-ext-community."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-ext-community."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "send-ext-community."+"vipValue", strconv.FormatBool(item.SendExtCommunity.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "ebgp-multihop")

		if !item.EbgpMultihopVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ebgp-multihop."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ebgp-multihop."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ebgp-multihop."+"vipVariableName", item.EbgpMultihopVariable.ValueString())
		} else if item.EbgpMultihop.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ebgp-multihop."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ebgp-multihop."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "ebgp-multihop."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ebgp-multihop."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ebgp-multihop."+"vipValue", item.EbgpMultihop.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "password")

		if !item.PasswordVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "password."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "password."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "password."+"vipVariableName", item.PasswordVariable.ValueString())
		} else if item.Password.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "password."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "password."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "password."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "password."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "password."+"vipValue", item.Password.ValueString())
		}
		itemAttributes = append(itemAttributes, "send-label")

		if !item.SendLabelVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-label."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-label."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "send-label."+"vipVariableName", item.SendLabelVariable.ValueString())
		} else if item.SendLabel.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-label."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-label."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "send-label."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-label."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "send-label."+"vipValue", strconv.FormatBool(item.SendLabel.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "send-label-explicit")

		if !item.SendLabelExplicitVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-label-explicit."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-label-explicit."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "send-label-explicit."+"vipVariableName", item.SendLabelExplicitVariable.ValueString())
		} else if item.SendLabelExplicit.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "send-label-explicit."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-label-explicit."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "send-label-explicit."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "send-label-explicit."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "send-label-explicit."+"vipValue", strconv.FormatBool(item.SendLabelExplicit.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "as-override")

		if !item.AsOverrideVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "as-override."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "as-override."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "as-override."+"vipVariableName", item.AsOverrideVariable.ValueString())
		} else if item.AsOverride.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "as-override."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "as-override."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "as-override."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "as-override."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "as-override."+"vipValue", strconv.FormatBool(item.AsOverride.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "allowas-in")

		if !item.AllowAsInVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "allowas-in.as-number."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "allowas-in.as-number."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "allowas-in.as-number."+"vipVariableName", item.AllowAsInVariable.ValueString())
		} else if item.AllowAsIn.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "allowas-in.as-number."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "allowas-in.as-number."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "allowas-in.as-number."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "allowas-in.as-number."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "allowas-in.as-number."+"vipValue", item.AllowAsIn.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "address-family")
		if len(item.AddressFamilies) > 0 {
			itemBody, _ = sjson.Set(itemBody, "address-family."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "address-family."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "address-family."+"vipPrimaryKey", []string{"family-type"})
			itemBody, _ = sjson.Set(itemBody, "address-family."+"vipValue", []interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "address-family."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "address-family."+"vipType", "ignore")
			itemBody, _ = sjson.Set(itemBody, "address-family."+"vipPrimaryKey", []string{"family-type"})
			itemBody, _ = sjson.Set(itemBody, "address-family."+"vipValue", []interface{}{})
		}
		for _, childItem := range item.AddressFamilies {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "family-type")
			if childItem.FamilyType.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "family-type."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "family-type."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "family-type."+"vipValue", childItem.FamilyType.ValueString())
			}
			itemChildAttributes = append(itemChildAttributes, "maximum-prefixes")

			if !childItem.MaximumPrefixesVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.prefix-num."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.prefix-num."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.prefix-num."+"vipVariableName", childItem.MaximumPrefixesVariable.ValueString())
			} else if childItem.MaximumPrefixes.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.prefix-num."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.prefix-num."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.prefix-num."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.prefix-num."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.prefix-num."+"vipValue", childItem.MaximumPrefixes.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "maximum-prefixes")

			if !childItem.MaximumPrefixesThresholdVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.threshold."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.threshold."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.threshold."+"vipVariableName", childItem.MaximumPrefixesThresholdVariable.ValueString())
			} else if childItem.MaximumPrefixesThreshold.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.threshold."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.threshold."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.threshold."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.threshold."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.threshold."+"vipValue", childItem.MaximumPrefixesThreshold.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "maximum-prefixes")

			if !childItem.MaximumPrefixesRestartVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.restart."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.restart."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.restart."+"vipVariableName", childItem.MaximumPrefixesRestartVariable.ValueString())
			} else if childItem.MaximumPrefixesRestart.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.restart."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.restart."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.restart."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.restart."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.restart."+"vipValue", childItem.MaximumPrefixesRestart.ValueInt64())
			}
			itemChildAttributes = append(itemChildAttributes, "maximum-prefixes")

			if !childItem.MaximumPrefixesWarningOnlyVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.warning-only."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.warning-only."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.warning-only."+"vipVariableName", childItem.MaximumPrefixesWarningOnlyVariable.ValueString())
			} else if childItem.MaximumPrefixesWarningOnly.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.warning-only."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.warning-only."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.warning-only."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.warning-only."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "maximum-prefixes.warning-only."+"vipValue", strconv.FormatBool(childItem.MaximumPrefixesWarningOnly.ValueBool()))
			}
			itemChildAttributes = append(itemChildAttributes, "route-policy")
			if len(childItem.RoutePolicies) > 0 {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipObjectType", "tree")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipPrimaryKey", []string{"direction"})
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipValue", []interface{}{})
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipObjectType", "tree")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipType", "ignore")
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipPrimaryKey", []string{"direction"})
				itemChildBody, _ = sjson.Set(itemChildBody, "route-policy."+"vipValue", []interface{}{})
			}
			for _, childChildItem := range childItem.RoutePolicies {
				itemChildChildBody := ""
				itemChildChildAttributes := make([]string, 0)
				itemChildChildAttributes = append(itemChildChildAttributes, "direction")
				if childChildItem.Direction.IsNull() {
				} else {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "direction."+"vipObjectType", "object")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "direction."+"vipType", "constant")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "direction."+"vipValue", childChildItem.Direction.ValueString())
				}
				itemChildChildAttributes = append(itemChildChildAttributes, "pol-name")

				if !childChildItem.PolicyNameVariable.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "pol-name."+"vipObjectType", "object")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "pol-name."+"vipType", "variableName")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "pol-name."+"vipVariableName", childChildItem.PolicyNameVariable.ValueString())
				} else if childChildItem.PolicyName.IsNull() {
				} else {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "pol-name."+"vipObjectType", "object")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "pol-name."+"vipType", "constant")
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "pol-name."+"vipValue", childChildItem.PolicyName.ValueString())
				}
				if !childChildItem.Optional.IsNull() {
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "vipOptional", childChildItem.Optional.ValueBool())
					itemChildChildBody, _ = sjson.Set(itemChildChildBody, "priority-order", itemChildChildAttributes)
				}
				itemChildBody, _ = sjson.SetRaw(itemChildBody, "route-policy."+"vipValue.-1", itemChildChildBody)
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "address-family."+"vipValue.-1", itemChildBody)
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"bgp.ipv6-neighbor."+"vipValue.-1", itemBody)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoBGP) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "bgp.as-num.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AsNumber = types.StringNull()

			v := res.Get(path + "bgp.as-num.vipVariableName")
			data.AsNumberVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AsNumber = types.StringNull()
			data.AsNumberVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bgp.as-num.vipValue")
			data.AsNumber = types.StringValue(v.String())
			data.AsNumberVariable = types.StringNull()
		}
	} else {
		data.AsNumber = types.StringNull()
		data.AsNumberVariable = types.StringNull()
	}
	if value := res.Get(path + "bgp.shutdown.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Shutdown = types.BoolNull()

			v := res.Get(path + "bgp.shutdown.vipVariableName")
			data.ShutdownVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Shutdown = types.BoolNull()
			data.ShutdownVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bgp.shutdown.vipValue")
			data.Shutdown = types.BoolValue(v.Bool())
			data.ShutdownVariable = types.StringNull()
		}
	} else {
		data.Shutdown = types.BoolNull()
		data.ShutdownVariable = types.StringNull()
	}
	if value := res.Get(path + "bgp.router-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.RouterId = types.StringNull()

			v := res.Get(path + "bgp.router-id.vipVariableName")
			data.RouterIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.RouterId = types.StringNull()
			data.RouterIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bgp.router-id.vipValue")
			data.RouterId = types.StringValue(v.String())
			data.RouterIdVariable = types.StringNull()
		}
	} else {
		data.RouterId = types.StringNull()
		data.RouterIdVariable = types.StringNull()
	}
	if value := res.Get(path + "bgp.propagate-aspath.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PropagateAspath = types.BoolNull()

			v := res.Get(path + "bgp.propagate-aspath.vipVariableName")
			data.PropagateAspathVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.PropagateAspath = types.BoolNull()
			data.PropagateAspathVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bgp.propagate-aspath.vipValue")
			data.PropagateAspath = types.BoolValue(v.Bool())
			data.PropagateAspathVariable = types.StringNull()
		}
	} else {
		data.PropagateAspath = types.BoolNull()
		data.PropagateAspathVariable = types.StringNull()
	}
	if value := res.Get(path + "bgp.propagate-community.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PropagateCommunity = types.BoolNull()

			v := res.Get(path + "bgp.propagate-community.vipVariableName")
			data.PropagateCommunityVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.PropagateCommunity = types.BoolNull()
			data.PropagateCommunityVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bgp.propagate-community.vipValue")
			data.PropagateCommunity = types.BoolValue(v.Bool())
			data.PropagateCommunityVariable = types.StringNull()
		}
	} else {
		data.PropagateCommunity = types.BoolNull()
		data.PropagateCommunityVariable = types.StringNull()
	}
	if value := res.Get(path + "bgp.target.route-target-ipv4.vipValue"); len(value.Array()) > 0 {
		data.Ipv4RouteTargets = make([]CiscoBGPIpv4RouteTargets, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoBGPIpv4RouteTargets{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("vpn-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

					cv := v.Get("vpn-id.vipVariableName")
					item.VpnIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()
					item.VpnIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("vpn-id.vipValue")
					item.VpnId = types.Int64Value(cv.Int())
					item.VpnIdVariable = types.StringNull()
				}
			} else {
				item.VpnId = types.Int64Null()
				item.VpnIdVariable = types.StringNull()
			}
			if cValue := v.Get("export.vipValue"); len(cValue.Array()) > 0 {
				item.Export = make([]CiscoBGPIpv4RouteTargetsExport, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoBGPIpv4RouteTargetsExport{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("asn-ip.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AsnIp = types.StringNull()

							ccv := cv.Get("asn-ip.vipVariableName")
							cItem.AsnIpVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AsnIp = types.StringNull()
							cItem.AsnIpVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("asn-ip.vipValue")
							cItem.AsnIp = types.StringValue(ccv.String())
							cItem.AsnIpVariable = types.StringNull()
						}
					} else {
						cItem.AsnIp = types.StringNull()
						cItem.AsnIpVariable = types.StringNull()
					}
					item.Export = append(item.Export, cItem)
					return true
				})
			} else {
				if len(item.Export) > 0 {
					item.Export = []CiscoBGPIpv4RouteTargetsExport{}
				}
			}
			if cValue := v.Get("import.vipValue"); len(cValue.Array()) > 0 {
				item.Import = make([]CiscoBGPIpv4RouteTargetsImport, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoBGPIpv4RouteTargetsImport{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("asn-ip.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AsnIp = types.StringNull()

							ccv := cv.Get("asn-ip.vipVariableName")
							cItem.AsnIpVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AsnIp = types.StringNull()
							cItem.AsnIpVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("asn-ip.vipValue")
							cItem.AsnIp = types.StringValue(ccv.String())
							cItem.AsnIpVariable = types.StringNull()
						}
					} else {
						cItem.AsnIp = types.StringNull()
						cItem.AsnIpVariable = types.StringNull()
					}
					item.Import = append(item.Import, cItem)
					return true
				})
			} else {
				if len(item.Import) > 0 {
					item.Import = []CiscoBGPIpv4RouteTargetsImport{}
				}
			}
			data.Ipv4RouteTargets = append(data.Ipv4RouteTargets, item)
			return true
		})
	} else {
		if len(data.Ipv4RouteTargets) > 0 {
			data.Ipv4RouteTargets = []CiscoBGPIpv4RouteTargets{}
		}
	}
	if value := res.Get(path + "bgp.target.route-target-ipv6.vipValue"); len(value.Array()) > 0 {
		data.Ipv6RouteTargets = make([]CiscoBGPIpv6RouteTargets, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoBGPIpv6RouteTargets{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("vpn-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

					cv := v.Get("vpn-id.vipVariableName")
					item.VpnIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()
					item.VpnIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("vpn-id.vipValue")
					item.VpnId = types.Int64Value(cv.Int())
					item.VpnIdVariable = types.StringNull()
				}
			} else {
				item.VpnId = types.Int64Null()
				item.VpnIdVariable = types.StringNull()
			}
			if cValue := v.Get("export.vipValue"); len(cValue.Array()) > 0 {
				item.Export = make([]CiscoBGPIpv6RouteTargetsExport, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoBGPIpv6RouteTargetsExport{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("asn-ip.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AsnIp = types.StringNull()

							ccv := cv.Get("asn-ip.vipVariableName")
							cItem.AsnIpVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AsnIp = types.StringNull()
							cItem.AsnIpVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("asn-ip.vipValue")
							cItem.AsnIp = types.StringValue(ccv.String())
							cItem.AsnIpVariable = types.StringNull()
						}
					} else {
						cItem.AsnIp = types.StringNull()
						cItem.AsnIpVariable = types.StringNull()
					}
					item.Export = append(item.Export, cItem)
					return true
				})
			} else {
				if len(item.Export) > 0 {
					item.Export = []CiscoBGPIpv6RouteTargetsExport{}
				}
			}
			if cValue := v.Get("import.vipValue"); len(cValue.Array()) > 0 {
				item.Import = make([]CiscoBGPIpv6RouteTargetsImport, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoBGPIpv6RouteTargetsImport{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("asn-ip.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AsnIp = types.StringNull()

							ccv := cv.Get("asn-ip.vipVariableName")
							cItem.AsnIpVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AsnIp = types.StringNull()
							cItem.AsnIpVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("asn-ip.vipValue")
							cItem.AsnIp = types.StringValue(ccv.String())
							cItem.AsnIpVariable = types.StringNull()
						}
					} else {
						cItem.AsnIp = types.StringNull()
						cItem.AsnIpVariable = types.StringNull()
					}
					item.Import = append(item.Import, cItem)
					return true
				})
			} else {
				if len(item.Import) > 0 {
					item.Import = []CiscoBGPIpv6RouteTargetsImport{}
				}
			}
			data.Ipv6RouteTargets = append(data.Ipv6RouteTargets, item)
			return true
		})
	} else {
		if len(data.Ipv6RouteTargets) > 0 {
			data.Ipv6RouteTargets = []CiscoBGPIpv6RouteTargets{}
		}
	}
	if value := res.Get(path + "bgp.mpls-interface.vipValue"); len(value.Array()) > 0 {
		data.MplsInterfaces = make([]CiscoBGPMplsInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoBGPMplsInterfaces{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("if-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.InterfaceName = types.StringNull()

					cv := v.Get("if-name.vipVariableName")
					item.InterfaceNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.InterfaceName = types.StringNull()
					item.InterfaceNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("if-name.vipValue")
					item.InterfaceName = types.StringValue(cv.String())
					item.InterfaceNameVariable = types.StringNull()
				}
			} else {
				item.InterfaceName = types.StringNull()
				item.InterfaceNameVariable = types.StringNull()
			}
			data.MplsInterfaces = append(data.MplsInterfaces, item)
			return true
		})
	} else {
		if len(data.MplsInterfaces) > 0 {
			data.MplsInterfaces = []CiscoBGPMplsInterfaces{}
		}
	}
	if value := res.Get(path + "bgp.distance.external.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DistanceExternal = types.Int64Null()

			v := res.Get(path + "bgp.distance.external.vipVariableName")
			data.DistanceExternalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DistanceExternal = types.Int64Null()
			data.DistanceExternalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bgp.distance.external.vipValue")
			data.DistanceExternal = types.Int64Value(v.Int())
			data.DistanceExternalVariable = types.StringNull()
		}
	} else {
		data.DistanceExternal = types.Int64Null()
		data.DistanceExternalVariable = types.StringNull()
	}
	if value := res.Get(path + "bgp.distance.internal.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DistanceInternal = types.Int64Null()

			v := res.Get(path + "bgp.distance.internal.vipVariableName")
			data.DistanceInternalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DistanceInternal = types.Int64Null()
			data.DistanceInternalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bgp.distance.internal.vipValue")
			data.DistanceInternal = types.Int64Value(v.Int())
			data.DistanceInternalVariable = types.StringNull()
		}
	} else {
		data.DistanceInternal = types.Int64Null()
		data.DistanceInternalVariable = types.StringNull()
	}
	if value := res.Get(path + "bgp.distance.local.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DistanceLocal = types.Int64Null()

			v := res.Get(path + "bgp.distance.local.vipVariableName")
			data.DistanceLocalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DistanceLocal = types.Int64Null()
			data.DistanceLocalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bgp.distance.local.vipValue")
			data.DistanceLocal = types.Int64Value(v.Int())
			data.DistanceLocalVariable = types.StringNull()
		}
	} else {
		data.DistanceLocal = types.Int64Null()
		data.DistanceLocalVariable = types.StringNull()
	}
	if value := res.Get(path + "bgp.timers.keepalive.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Keepalive = types.Int64Null()

			v := res.Get(path + "bgp.timers.keepalive.vipVariableName")
			data.KeepaliveVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Keepalive = types.Int64Null()
			data.KeepaliveVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bgp.timers.keepalive.vipValue")
			data.Keepalive = types.Int64Value(v.Int())
			data.KeepaliveVariable = types.StringNull()
		}
	} else {
		data.Keepalive = types.Int64Null()
		data.KeepaliveVariable = types.StringNull()
	}
	if value := res.Get(path + "bgp.timers.holdtime.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Holdtime = types.Int64Null()

			v := res.Get(path + "bgp.timers.holdtime.vipVariableName")
			data.HoldtimeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Holdtime = types.Int64Null()
			data.HoldtimeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bgp.timers.holdtime.vipValue")
			data.Holdtime = types.Int64Value(v.Int())
			data.HoldtimeVariable = types.StringNull()
		}
	} else {
		data.Holdtime = types.Int64Null()
		data.HoldtimeVariable = types.StringNull()
	}
	if value := res.Get(path + "bgp.best-path.med.always-compare.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AlwaysCompareMed = types.BoolNull()

			v := res.Get(path + "bgp.best-path.med.always-compare.vipVariableName")
			data.AlwaysCompareMedVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AlwaysCompareMed = types.BoolNull()
			data.AlwaysCompareMedVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bgp.best-path.med.always-compare.vipValue")
			data.AlwaysCompareMed = types.BoolValue(v.Bool())
			data.AlwaysCompareMedVariable = types.StringNull()
		}
	} else {
		data.AlwaysCompareMed = types.BoolNull()
		data.AlwaysCompareMedVariable = types.StringNull()
	}
	if value := res.Get(path + "bgp.best-path.med.deterministic.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DeterministicMed = types.BoolNull()

			v := res.Get(path + "bgp.best-path.med.deterministic.vipVariableName")
			data.DeterministicMedVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DeterministicMed = types.BoolNull()
			data.DeterministicMedVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bgp.best-path.med.deterministic.vipValue")
			data.DeterministicMed = types.BoolValue(v.Bool())
			data.DeterministicMedVariable = types.StringNull()
		}
	} else {
		data.DeterministicMed = types.BoolNull()
		data.DeterministicMedVariable = types.StringNull()
	}
	if value := res.Get(path + "bgp.best-path.med.missing-as-worst.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.MissingMedWorst = types.BoolNull()

			v := res.Get(path + "bgp.best-path.med.missing-as-worst.vipVariableName")
			data.MissingMedWorstVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.MissingMedWorst = types.BoolNull()
			data.MissingMedWorstVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bgp.best-path.med.missing-as-worst.vipValue")
			data.MissingMedWorst = types.BoolValue(v.Bool())
			data.MissingMedWorstVariable = types.StringNull()
		}
	} else {
		data.MissingMedWorst = types.BoolNull()
		data.MissingMedWorstVariable = types.StringNull()
	}
	if value := res.Get(path + "bgp.best-path.compare-router-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.CompareRouterId = types.BoolNull()

			v := res.Get(path + "bgp.best-path.compare-router-id.vipVariableName")
			data.CompareRouterIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.CompareRouterId = types.BoolNull()
			data.CompareRouterIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bgp.best-path.compare-router-id.vipValue")
			data.CompareRouterId = types.BoolValue(v.Bool())
			data.CompareRouterIdVariable = types.StringNull()
		}
	} else {
		data.CompareRouterId = types.BoolNull()
		data.CompareRouterIdVariable = types.StringNull()
	}
	if value := res.Get(path + "bgp.best-path.as-path.multipath-relax.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.MultipathRelax = types.BoolNull()

			v := res.Get(path + "bgp.best-path.as-path.multipath-relax.vipVariableName")
			data.MultipathRelaxVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.MultipathRelax = types.BoolNull()
			data.MultipathRelaxVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "bgp.best-path.as-path.multipath-relax.vipValue")
			data.MultipathRelax = types.BoolValue(v.Bool())
			data.MultipathRelaxVariable = types.StringNull()
		}
	} else {
		data.MultipathRelax = types.BoolNull()
		data.MultipathRelaxVariable = types.StringNull()
	}
	if value := res.Get(path + "bgp.address-family.vipValue"); len(value.Array()) > 0 {
		data.AddressFamilies = make([]CiscoBGPAddressFamilies, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoBGPAddressFamilies{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("family-type.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.FamilyType = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.FamilyType = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("family-type.vipValue")
					item.FamilyType = types.StringValue(cv.String())

				}
			} else {
				item.FamilyType = types.StringNull()

			}
			if cValue := v.Get("aggregate-address.vipValue"); len(cValue.Array()) > 0 {
				item.Ipv4AggregateAddresses = make([]CiscoBGPAddressFamiliesIpv4AggregateAddresses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoBGPAddressFamiliesIpv4AggregateAddresses{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("prefix.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Prefix = types.StringNull()

							ccv := cv.Get("prefix.vipVariableName")
							cItem.PrefixVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Prefix = types.StringNull()
							cItem.PrefixVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("prefix.vipValue")
							cItem.Prefix = types.StringValue(ccv.String())
							cItem.PrefixVariable = types.StringNull()
						}
					} else {
						cItem.Prefix = types.StringNull()
						cItem.PrefixVariable = types.StringNull()
					}
					if ccValue := cv.Get("as-set.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AsSetPath = types.BoolNull()

							ccv := cv.Get("as-set.vipVariableName")
							cItem.AsSetPathVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AsSetPath = types.BoolNull()
							cItem.AsSetPathVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("as-set.vipValue")
							cItem.AsSetPath = types.BoolValue(ccv.Bool())
							cItem.AsSetPathVariable = types.StringNull()
						}
					} else {
						cItem.AsSetPath = types.BoolNull()
						cItem.AsSetPathVariable = types.StringNull()
					}
					if ccValue := cv.Get("summary-only.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.SummaryOnly = types.BoolNull()

							ccv := cv.Get("summary-only.vipVariableName")
							cItem.SummaryOnlyVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.SummaryOnly = types.BoolNull()
							cItem.SummaryOnlyVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("summary-only.vipValue")
							cItem.SummaryOnly = types.BoolValue(ccv.Bool())
							cItem.SummaryOnlyVariable = types.StringNull()
						}
					} else {
						cItem.SummaryOnly = types.BoolNull()
						cItem.SummaryOnlyVariable = types.StringNull()
					}
					item.Ipv4AggregateAddresses = append(item.Ipv4AggregateAddresses, cItem)
					return true
				})
			} else {
				if len(item.Ipv4AggregateAddresses) > 0 {
					item.Ipv4AggregateAddresses = []CiscoBGPAddressFamiliesIpv4AggregateAddresses{}
				}
			}
			if cValue := v.Get("ipv6-aggregate-address.vipValue"); len(cValue.Array()) > 0 {
				item.Ipv6AggregateAddresses = make([]CiscoBGPAddressFamiliesIpv6AggregateAddresses, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoBGPAddressFamiliesIpv6AggregateAddresses{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("prefix.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Prefix = types.StringNull()

							ccv := cv.Get("prefix.vipVariableName")
							cItem.PrefixVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Prefix = types.StringNull()
							cItem.PrefixVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("prefix.vipValue")
							cItem.Prefix = types.StringValue(ccv.String())
							cItem.PrefixVariable = types.StringNull()
						}
					} else {
						cItem.Prefix = types.StringNull()
						cItem.PrefixVariable = types.StringNull()
					}
					if ccValue := cv.Get("as-set.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.AsSetPath = types.BoolNull()

							ccv := cv.Get("as-set.vipVariableName")
							cItem.AsSetPathVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.AsSetPath = types.BoolNull()
							cItem.AsSetPathVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("as-set.vipValue")
							cItem.AsSetPath = types.BoolValue(ccv.Bool())
							cItem.AsSetPathVariable = types.StringNull()
						}
					} else {
						cItem.AsSetPath = types.BoolNull()
						cItem.AsSetPathVariable = types.StringNull()
					}
					if ccValue := cv.Get("summary-only.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.SummaryOnly = types.BoolNull()

							ccv := cv.Get("summary-only.vipVariableName")
							cItem.SummaryOnlyVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.SummaryOnly = types.BoolNull()
							cItem.SummaryOnlyVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("summary-only.vipValue")
							cItem.SummaryOnly = types.BoolValue(ccv.Bool())
							cItem.SummaryOnlyVariable = types.StringNull()
						}
					} else {
						cItem.SummaryOnly = types.BoolNull()
						cItem.SummaryOnlyVariable = types.StringNull()
					}
					item.Ipv6AggregateAddresses = append(item.Ipv6AggregateAddresses, cItem)
					return true
				})
			} else {
				if len(item.Ipv6AggregateAddresses) > 0 {
					item.Ipv6AggregateAddresses = []CiscoBGPAddressFamiliesIpv6AggregateAddresses{}
				}
			}
			if cValue := v.Get("network.vipValue"); len(cValue.Array()) > 0 {
				item.Ipv4Networks = make([]CiscoBGPAddressFamiliesIpv4Networks, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoBGPAddressFamiliesIpv4Networks{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("prefix.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Prefix = types.StringNull()

							ccv := cv.Get("prefix.vipVariableName")
							cItem.PrefixVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Prefix = types.StringNull()
							cItem.PrefixVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("prefix.vipValue")
							cItem.Prefix = types.StringValue(ccv.String())
							cItem.PrefixVariable = types.StringNull()
						}
					} else {
						cItem.Prefix = types.StringNull()
						cItem.PrefixVariable = types.StringNull()
					}
					item.Ipv4Networks = append(item.Ipv4Networks, cItem)
					return true
				})
			} else {
				if len(item.Ipv4Networks) > 0 {
					item.Ipv4Networks = []CiscoBGPAddressFamiliesIpv4Networks{}
				}
			}
			if cValue := v.Get("ipv6-network.vipValue"); len(cValue.Array()) > 0 {
				item.Ipv6Networks = make([]CiscoBGPAddressFamiliesIpv6Networks, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoBGPAddressFamiliesIpv6Networks{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("prefix.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Prefix = types.StringNull()

							ccv := cv.Get("prefix.vipVariableName")
							cItem.PrefixVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Prefix = types.StringNull()
							cItem.PrefixVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("prefix.vipValue")
							cItem.Prefix = types.StringValue(ccv.String())
							cItem.PrefixVariable = types.StringNull()
						}
					} else {
						cItem.Prefix = types.StringNull()
						cItem.PrefixVariable = types.StringNull()
					}
					item.Ipv6Networks = append(item.Ipv6Networks, cItem)
					return true
				})
			} else {
				if len(item.Ipv6Networks) > 0 {
					item.Ipv6Networks = []CiscoBGPAddressFamiliesIpv6Networks{}
				}
			}
			if cValue := v.Get("maximum-paths.paths.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.MaximumPaths = types.Int64Null()

					cv := v.Get("maximum-paths.paths.vipVariableName")
					item.MaximumPathsVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.MaximumPaths = types.Int64Null()
					item.MaximumPathsVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("maximum-paths.paths.vipValue")
					item.MaximumPaths = types.Int64Value(cv.Int())
					item.MaximumPathsVariable = types.StringNull()
				}
			} else {
				item.MaximumPaths = types.Int64Null()
				item.MaximumPathsVariable = types.StringNull()
			}
			if cValue := v.Get("default-information.originate.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.DefaultInformationOriginate = types.BoolNull()

					cv := v.Get("default-information.originate.vipVariableName")
					item.DefaultInformationOriginateVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.DefaultInformationOriginate = types.BoolNull()
					item.DefaultInformationOriginateVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("default-information.originate.vipValue")
					item.DefaultInformationOriginate = types.BoolValue(cv.Bool())
					item.DefaultInformationOriginateVariable = types.StringNull()
				}
			} else {
				item.DefaultInformationOriginate = types.BoolNull()
				item.DefaultInformationOriginateVariable = types.StringNull()
			}
			if cValue := v.Get("table-map.name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TableMapPolicy = types.StringNull()

					cv := v.Get("table-map.name.vipVariableName")
					item.TableMapPolicyVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TableMapPolicy = types.StringNull()
					item.TableMapPolicyVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("table-map.name.vipValue")
					item.TableMapPolicy = types.StringValue(cv.String())
					item.TableMapPolicyVariable = types.StringNull()
				}
			} else {
				item.TableMapPolicy = types.StringNull()
				item.TableMapPolicyVariable = types.StringNull()
			}
			if cValue := v.Get("table-map.filter.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TableMapFilter = types.BoolNull()

					cv := v.Get("table-map.filter.vipVariableName")
					item.TableMapFilterVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TableMapFilter = types.BoolNull()
					item.TableMapFilterVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("table-map.filter.vipValue")
					item.TableMapFilter = types.BoolValue(cv.Bool())
					item.TableMapFilterVariable = types.StringNull()
				}
			} else {
				item.TableMapFilter = types.BoolNull()
				item.TableMapFilterVariable = types.StringNull()
			}
			if cValue := v.Get("redistribute.vipValue"); len(cValue.Array()) > 0 {
				item.RedistributeRoutes = make([]CiscoBGPAddressFamiliesRedistributeRoutes, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoBGPAddressFamiliesRedistributeRoutes{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("protocol.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.Protocol = types.StringNull()

							ccv := cv.Get("protocol.vipVariableName")
							cItem.ProtocolVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.Protocol = types.StringNull()
							cItem.ProtocolVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("protocol.vipValue")
							cItem.Protocol = types.StringValue(ccv.String())
							cItem.ProtocolVariable = types.StringNull()
						}
					} else {
						cItem.Protocol = types.StringNull()
						cItem.ProtocolVariable = types.StringNull()
					}
					if ccValue := cv.Get("route-policy.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.RoutePolicy = types.StringNull()

							ccv := cv.Get("route-policy.vipVariableName")
							cItem.RoutePolicyVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.RoutePolicy = types.StringNull()
							cItem.RoutePolicyVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("route-policy.vipValue")
							cItem.RoutePolicy = types.StringValue(ccv.String())
							cItem.RoutePolicyVariable = types.StringNull()
						}
					} else {
						cItem.RoutePolicy = types.StringNull()
						cItem.RoutePolicyVariable = types.StringNull()
					}
					item.RedistributeRoutes = append(item.RedistributeRoutes, cItem)
					return true
				})
			} else {
				if len(item.RedistributeRoutes) > 0 {
					item.RedistributeRoutes = []CiscoBGPAddressFamiliesRedistributeRoutes{}
				}
			}
			data.AddressFamilies = append(data.AddressFamilies, item)
			return true
		})
	} else {
		if len(data.AddressFamilies) > 0 {
			data.AddressFamilies = []CiscoBGPAddressFamilies{}
		}
	}
	if value := res.Get(path + "bgp.neighbor.vipValue"); len(value.Array()) > 0 {
		data.Ipv4Neighbors = make([]CiscoBGPIpv4Neighbors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoBGPIpv4Neighbors{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("address.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Address = types.StringNull()

					cv := v.Get("address.vipVariableName")
					item.AddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Address = types.StringNull()
					item.AddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("address.vipValue")
					item.Address = types.StringValue(cv.String())
					item.AddressVariable = types.StringNull()
				}
			} else {
				item.Address = types.StringNull()
				item.AddressVariable = types.StringNull()
			}
			if cValue := v.Get("description.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Description = types.StringNull()

					cv := v.Get("description.vipVariableName")
					item.DescriptionVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Description = types.StringNull()
					item.DescriptionVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("description.vipValue")
					item.Description = types.StringValue(cv.String())
					item.DescriptionVariable = types.StringNull()
				}
			} else {
				item.Description = types.StringNull()
				item.DescriptionVariable = types.StringNull()
			}
			if cValue := v.Get("shutdown.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Shutdown = types.BoolNull()

					cv := v.Get("shutdown.vipVariableName")
					item.ShutdownVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Shutdown = types.BoolNull()
					item.ShutdownVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("shutdown.vipValue")
					item.Shutdown = types.BoolValue(cv.Bool())
					item.ShutdownVariable = types.StringNull()
				}
			} else {
				item.Shutdown = types.BoolNull()
				item.ShutdownVariable = types.StringNull()
			}
			if cValue := v.Get("remote-as.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.RemoteAs = types.StringNull()

					cv := v.Get("remote-as.vipVariableName")
					item.RemoteAsVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.RemoteAs = types.StringNull()
					item.RemoteAsVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("remote-as.vipValue")
					item.RemoteAs = types.StringValue(cv.String())
					item.RemoteAsVariable = types.StringNull()
				}
			} else {
				item.RemoteAs = types.StringNull()
				item.RemoteAsVariable = types.StringNull()
			}
			if cValue := v.Get("timers.keepalive.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Keepalive = types.Int64Null()

					cv := v.Get("timers.keepalive.vipVariableName")
					item.KeepaliveVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Keepalive = types.Int64Null()
					item.KeepaliveVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("timers.keepalive.vipValue")
					item.Keepalive = types.Int64Value(cv.Int())
					item.KeepaliveVariable = types.StringNull()
				}
			} else {
				item.Keepalive = types.Int64Null()
				item.KeepaliveVariable = types.StringNull()
			}
			if cValue := v.Get("timers.holdtime.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Holdtime = types.Int64Null()

					cv := v.Get("timers.holdtime.vipVariableName")
					item.HoldtimeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Holdtime = types.Int64Null()
					item.HoldtimeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("timers.holdtime.vipValue")
					item.Holdtime = types.Int64Value(cv.Int())
					item.HoldtimeVariable = types.StringNull()
				}
			} else {
				item.Holdtime = types.Int64Null()
				item.HoldtimeVariable = types.StringNull()
			}
			if cValue := v.Get("update-source.if-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceInterface = types.StringNull()

					cv := v.Get("update-source.if-name.vipVariableName")
					item.SourceInterfaceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceInterface = types.StringNull()
					item.SourceInterfaceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("update-source.if-name.vipValue")
					item.SourceInterface = types.StringValue(cv.String())
					item.SourceInterfaceVariable = types.StringNull()
				}
			} else {
				item.SourceInterface = types.StringNull()
				item.SourceInterfaceVariable = types.StringNull()
			}
			if cValue := v.Get("next-hop-self.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.NextHopSelf = types.BoolNull()

					cv := v.Get("next-hop-self.vipVariableName")
					item.NextHopSelfVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.NextHopSelf = types.BoolNull()
					item.NextHopSelfVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("next-hop-self.vipValue")
					item.NextHopSelf = types.BoolValue(cv.Bool())
					item.NextHopSelfVariable = types.StringNull()
				}
			} else {
				item.NextHopSelf = types.BoolNull()
				item.NextHopSelfVariable = types.StringNull()
			}
			if cValue := v.Get("send-community.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SendCommunity = types.BoolNull()

					cv := v.Get("send-community.vipVariableName")
					item.SendCommunityVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SendCommunity = types.BoolNull()
					item.SendCommunityVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("send-community.vipValue")
					item.SendCommunity = types.BoolValue(cv.Bool())
					item.SendCommunityVariable = types.StringNull()
				}
			} else {
				item.SendCommunity = types.BoolNull()
				item.SendCommunityVariable = types.StringNull()
			}
			if cValue := v.Get("send-ext-community.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SendExtCommunity = types.BoolNull()

					cv := v.Get("send-ext-community.vipVariableName")
					item.SendExtCommunityVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SendExtCommunity = types.BoolNull()
					item.SendExtCommunityVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("send-ext-community.vipValue")
					item.SendExtCommunity = types.BoolValue(cv.Bool())
					item.SendExtCommunityVariable = types.StringNull()
				}
			} else {
				item.SendExtCommunity = types.BoolNull()
				item.SendExtCommunityVariable = types.StringNull()
			}
			if cValue := v.Get("ebgp-multihop.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.EbgpMultihop = types.Int64Null()

					cv := v.Get("ebgp-multihop.vipVariableName")
					item.EbgpMultihopVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.EbgpMultihop = types.Int64Null()
					item.EbgpMultihopVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ebgp-multihop.vipValue")
					item.EbgpMultihop = types.Int64Value(cv.Int())
					item.EbgpMultihopVariable = types.StringNull()
				}
			} else {
				item.EbgpMultihop = types.Int64Null()
				item.EbgpMultihopVariable = types.StringNull()
			}
			if cValue := v.Get("password.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Password = types.StringNull()

					cv := v.Get("password.vipVariableName")
					item.PasswordVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Password = types.StringNull()
					item.PasswordVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("password.vipValue")
					item.Password = types.StringValue(cv.String())
					item.PasswordVariable = types.StringNull()
				}
			} else {
				item.Password = types.StringNull()
				item.PasswordVariable = types.StringNull()
			}
			if cValue := v.Get("send-label.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SendLabel = types.BoolNull()

					cv := v.Get("send-label.vipVariableName")
					item.SendLabelVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SendLabel = types.BoolNull()
					item.SendLabelVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("send-label.vipValue")
					item.SendLabel = types.BoolValue(cv.Bool())
					item.SendLabelVariable = types.StringNull()
				}
			} else {
				item.SendLabel = types.BoolNull()
				item.SendLabelVariable = types.StringNull()
			}
			if cValue := v.Get("send-label-explicit.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SendLabelExplicit = types.BoolNull()

					cv := v.Get("send-label-explicit.vipVariableName")
					item.SendLabelExplicitVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SendLabelExplicit = types.BoolNull()
					item.SendLabelExplicitVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("send-label-explicit.vipValue")
					item.SendLabelExplicit = types.BoolValue(cv.Bool())
					item.SendLabelExplicitVariable = types.StringNull()
				}
			} else {
				item.SendLabelExplicit = types.BoolNull()
				item.SendLabelExplicitVariable = types.StringNull()
			}
			if cValue := v.Get("as-override.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AsOverride = types.BoolNull()

					cv := v.Get("as-override.vipVariableName")
					item.AsOverrideVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AsOverride = types.BoolNull()
					item.AsOverrideVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("as-override.vipValue")
					item.AsOverride = types.BoolValue(cv.Bool())
					item.AsOverrideVariable = types.StringNull()
				}
			} else {
				item.AsOverride = types.BoolNull()
				item.AsOverrideVariable = types.StringNull()
			}
			if cValue := v.Get("allowas-in.as-number.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AllowAsIn = types.Int64Null()

					cv := v.Get("allowas-in.as-number.vipVariableName")
					item.AllowAsInVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AllowAsIn = types.Int64Null()
					item.AllowAsInVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("allowas-in.as-number.vipValue")
					item.AllowAsIn = types.Int64Value(cv.Int())
					item.AllowAsInVariable = types.StringNull()
				}
			} else {
				item.AllowAsIn = types.Int64Null()
				item.AllowAsInVariable = types.StringNull()
			}
			if cValue := v.Get("address-family.vipValue"); len(cValue.Array()) > 0 {
				item.AddressFamilies = make([]CiscoBGPIpv4NeighborsAddressFamilies, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoBGPIpv4NeighborsAddressFamilies{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("family-type.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.FamilyType = types.StringNull()

						} else if ccValue.String() == "ignore" {
							cItem.FamilyType = types.StringNull()

						} else if ccValue.String() == "constant" {
							ccv := cv.Get("family-type.vipValue")
							cItem.FamilyType = types.StringValue(ccv.String())

						}
					} else {
						cItem.FamilyType = types.StringNull()

					}
					if ccValue := cv.Get("maximum-prefixes.prefix-num.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.MaximumPrefixes = types.Int64Null()

							ccv := cv.Get("maximum-prefixes.prefix-num.vipVariableName")
							cItem.MaximumPrefixesVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.MaximumPrefixes = types.Int64Null()
							cItem.MaximumPrefixesVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("maximum-prefixes.prefix-num.vipValue")
							cItem.MaximumPrefixes = types.Int64Value(ccv.Int())
							cItem.MaximumPrefixesVariable = types.StringNull()
						}
					} else {
						cItem.MaximumPrefixes = types.Int64Null()
						cItem.MaximumPrefixesVariable = types.StringNull()
					}
					if ccValue := cv.Get("maximum-prefixes.threshold.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.MaximumPrefixesThreshold = types.Int64Null()

							ccv := cv.Get("maximum-prefixes.threshold.vipVariableName")
							cItem.MaximumPrefixesThresholdVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.MaximumPrefixesThreshold = types.Int64Null()
							cItem.MaximumPrefixesThresholdVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("maximum-prefixes.threshold.vipValue")
							cItem.MaximumPrefixesThreshold = types.Int64Value(ccv.Int())
							cItem.MaximumPrefixesThresholdVariable = types.StringNull()
						}
					} else {
						cItem.MaximumPrefixesThreshold = types.Int64Null()
						cItem.MaximumPrefixesThresholdVariable = types.StringNull()
					}
					if ccValue := cv.Get("maximum-prefixes.restart.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.MaximumPrefixesRestart = types.Int64Null()

							ccv := cv.Get("maximum-prefixes.restart.vipVariableName")
							cItem.MaximumPrefixesRestartVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.MaximumPrefixesRestart = types.Int64Null()
							cItem.MaximumPrefixesRestartVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("maximum-prefixes.restart.vipValue")
							cItem.MaximumPrefixesRestart = types.Int64Value(ccv.Int())
							cItem.MaximumPrefixesRestartVariable = types.StringNull()
						}
					} else {
						cItem.MaximumPrefixesRestart = types.Int64Null()
						cItem.MaximumPrefixesRestartVariable = types.StringNull()
					}
					if ccValue := cv.Get("maximum-prefixes.warning-only.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.MaximumPrefixesWarningOnly = types.BoolNull()

							ccv := cv.Get("maximum-prefixes.warning-only.vipVariableName")
							cItem.MaximumPrefixesWarningOnlyVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.MaximumPrefixesWarningOnly = types.BoolNull()
							cItem.MaximumPrefixesWarningOnlyVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("maximum-prefixes.warning-only.vipValue")
							cItem.MaximumPrefixesWarningOnly = types.BoolValue(ccv.Bool())
							cItem.MaximumPrefixesWarningOnlyVariable = types.StringNull()
						}
					} else {
						cItem.MaximumPrefixesWarningOnly = types.BoolNull()
						cItem.MaximumPrefixesWarningOnlyVariable = types.StringNull()
					}
					if ccValue := cv.Get("route-policy.vipValue"); len(ccValue.Array()) > 0 {
						cItem.RoutePolicies = make([]CiscoBGPIpv4NeighborsAddressFamiliesRoutePolicies, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := CiscoBGPIpv4NeighborsAddressFamiliesRoutePolicies{}
							if cccValue := ccv.Get("vipOptional"); cccValue.Exists() {
								ccItem.Optional = types.BoolValue(cccValue.Bool())
							} else {
								ccItem.Optional = types.BoolNull()
							}
							if cccValue := ccv.Get("direction.vipType"); cccValue.Exists() {
								if cccValue.String() == "variableName" {
									ccItem.Direction = types.StringNull()

								} else if cccValue.String() == "ignore" {
									ccItem.Direction = types.StringNull()

								} else if cccValue.String() == "constant" {
									cccv := ccv.Get("direction.vipValue")
									ccItem.Direction = types.StringValue(cccv.String())

								}
							} else {
								ccItem.Direction = types.StringNull()

							}
							if cccValue := ccv.Get("pol-name.vipType"); cccValue.Exists() {
								if cccValue.String() == "variableName" {
									ccItem.PolicyName = types.StringNull()

									cccv := ccv.Get("pol-name.vipVariableName")
									ccItem.PolicyNameVariable = types.StringValue(cccv.String())

								} else if cccValue.String() == "ignore" {
									ccItem.PolicyName = types.StringNull()
									ccItem.PolicyNameVariable = types.StringNull()
								} else if cccValue.String() == "constant" {
									cccv := ccv.Get("pol-name.vipValue")
									ccItem.PolicyName = types.StringValue(cccv.String())
									ccItem.PolicyNameVariable = types.StringNull()
								}
							} else {
								ccItem.PolicyName = types.StringNull()
								ccItem.PolicyNameVariable = types.StringNull()
							}
							cItem.RoutePolicies = append(cItem.RoutePolicies, ccItem)
							return true
						})
					} else {
						if len(cItem.RoutePolicies) > 0 {
							cItem.RoutePolicies = []CiscoBGPIpv4NeighborsAddressFamiliesRoutePolicies{}
						}
					}
					item.AddressFamilies = append(item.AddressFamilies, cItem)
					return true
				})
			} else {
				if len(item.AddressFamilies) > 0 {
					item.AddressFamilies = []CiscoBGPIpv4NeighborsAddressFamilies{}
				}
			}
			data.Ipv4Neighbors = append(data.Ipv4Neighbors, item)
			return true
		})
	} else {
		if len(data.Ipv4Neighbors) > 0 {
			data.Ipv4Neighbors = []CiscoBGPIpv4Neighbors{}
		}
	}
	if value := res.Get(path + "bgp.ipv6-neighbor.vipValue"); len(value.Array()) > 0 {
		data.Ipv6Neighbors = make([]CiscoBGPIpv6Neighbors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoBGPIpv6Neighbors{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("address.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Address = types.StringNull()

					cv := v.Get("address.vipVariableName")
					item.AddressVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Address = types.StringNull()
					item.AddressVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("address.vipValue")
					item.Address = types.StringValue(cv.String())
					item.AddressVariable = types.StringNull()
				}
			} else {
				item.Address = types.StringNull()
				item.AddressVariable = types.StringNull()
			}
			if cValue := v.Get("description.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Description = types.StringNull()

					cv := v.Get("description.vipVariableName")
					item.DescriptionVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Description = types.StringNull()
					item.DescriptionVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("description.vipValue")
					item.Description = types.StringValue(cv.String())
					item.DescriptionVariable = types.StringNull()
				}
			} else {
				item.Description = types.StringNull()
				item.DescriptionVariable = types.StringNull()
			}
			if cValue := v.Get("shutdown.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Shutdown = types.BoolNull()

					cv := v.Get("shutdown.vipVariableName")
					item.ShutdownVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Shutdown = types.BoolNull()
					item.ShutdownVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("shutdown.vipValue")
					item.Shutdown = types.BoolValue(cv.Bool())
					item.ShutdownVariable = types.StringNull()
				}
			} else {
				item.Shutdown = types.BoolNull()
				item.ShutdownVariable = types.StringNull()
			}
			if cValue := v.Get("remote-as.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.RemoteAs = types.StringNull()

					cv := v.Get("remote-as.vipVariableName")
					item.RemoteAsVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.RemoteAs = types.StringNull()
					item.RemoteAsVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("remote-as.vipValue")
					item.RemoteAs = types.StringValue(cv.String())
					item.RemoteAsVariable = types.StringNull()
				}
			} else {
				item.RemoteAs = types.StringNull()
				item.RemoteAsVariable = types.StringNull()
			}
			if cValue := v.Get("timers.keepalive.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Keepalive = types.Int64Null()

					cv := v.Get("timers.keepalive.vipVariableName")
					item.KeepaliveVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Keepalive = types.Int64Null()
					item.KeepaliveVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("timers.keepalive.vipValue")
					item.Keepalive = types.Int64Value(cv.Int())
					item.KeepaliveVariable = types.StringNull()
				}
			} else {
				item.Keepalive = types.Int64Null()
				item.KeepaliveVariable = types.StringNull()
			}
			if cValue := v.Get("timers.holdtime.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Holdtime = types.Int64Null()

					cv := v.Get("timers.holdtime.vipVariableName")
					item.HoldtimeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Holdtime = types.Int64Null()
					item.HoldtimeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("timers.holdtime.vipValue")
					item.Holdtime = types.Int64Value(cv.Int())
					item.HoldtimeVariable = types.StringNull()
				}
			} else {
				item.Holdtime = types.Int64Null()
				item.HoldtimeVariable = types.StringNull()
			}
			if cValue := v.Get("update-source.if-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceInterface = types.StringNull()

					cv := v.Get("update-source.if-name.vipVariableName")
					item.SourceInterfaceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceInterface = types.StringNull()
					item.SourceInterfaceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("update-source.if-name.vipValue")
					item.SourceInterface = types.StringValue(cv.String())
					item.SourceInterfaceVariable = types.StringNull()
				}
			} else {
				item.SourceInterface = types.StringNull()
				item.SourceInterfaceVariable = types.StringNull()
			}
			if cValue := v.Get("next-hop-self.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.NextHopSelf = types.BoolNull()

					cv := v.Get("next-hop-self.vipVariableName")
					item.NextHopSelfVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.NextHopSelf = types.BoolNull()
					item.NextHopSelfVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("next-hop-self.vipValue")
					item.NextHopSelf = types.BoolValue(cv.Bool())
					item.NextHopSelfVariable = types.StringNull()
				}
			} else {
				item.NextHopSelf = types.BoolNull()
				item.NextHopSelfVariable = types.StringNull()
			}
			if cValue := v.Get("send-community.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SendCommunity = types.BoolNull()

					cv := v.Get("send-community.vipVariableName")
					item.SendCommunityVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SendCommunity = types.BoolNull()
					item.SendCommunityVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("send-community.vipValue")
					item.SendCommunity = types.BoolValue(cv.Bool())
					item.SendCommunityVariable = types.StringNull()
				}
			} else {
				item.SendCommunity = types.BoolNull()
				item.SendCommunityVariable = types.StringNull()
			}
			if cValue := v.Get("send-ext-community.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SendExtCommunity = types.BoolNull()

					cv := v.Get("send-ext-community.vipVariableName")
					item.SendExtCommunityVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SendExtCommunity = types.BoolNull()
					item.SendExtCommunityVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("send-ext-community.vipValue")
					item.SendExtCommunity = types.BoolValue(cv.Bool())
					item.SendExtCommunityVariable = types.StringNull()
				}
			} else {
				item.SendExtCommunity = types.BoolNull()
				item.SendExtCommunityVariable = types.StringNull()
			}
			if cValue := v.Get("ebgp-multihop.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.EbgpMultihop = types.Int64Null()

					cv := v.Get("ebgp-multihop.vipVariableName")
					item.EbgpMultihopVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.EbgpMultihop = types.Int64Null()
					item.EbgpMultihopVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ebgp-multihop.vipValue")
					item.EbgpMultihop = types.Int64Value(cv.Int())
					item.EbgpMultihopVariable = types.StringNull()
				}
			} else {
				item.EbgpMultihop = types.Int64Null()
				item.EbgpMultihopVariable = types.StringNull()
			}
			if cValue := v.Get("password.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Password = types.StringNull()

					cv := v.Get("password.vipVariableName")
					item.PasswordVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Password = types.StringNull()
					item.PasswordVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("password.vipValue")
					item.Password = types.StringValue(cv.String())
					item.PasswordVariable = types.StringNull()
				}
			} else {
				item.Password = types.StringNull()
				item.PasswordVariable = types.StringNull()
			}
			if cValue := v.Get("send-label.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SendLabel = types.BoolNull()

					cv := v.Get("send-label.vipVariableName")
					item.SendLabelVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SendLabel = types.BoolNull()
					item.SendLabelVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("send-label.vipValue")
					item.SendLabel = types.BoolValue(cv.Bool())
					item.SendLabelVariable = types.StringNull()
				}
			} else {
				item.SendLabel = types.BoolNull()
				item.SendLabelVariable = types.StringNull()
			}
			if cValue := v.Get("send-label-explicit.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SendLabelExplicit = types.BoolNull()

					cv := v.Get("send-label-explicit.vipVariableName")
					item.SendLabelExplicitVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SendLabelExplicit = types.BoolNull()
					item.SendLabelExplicitVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("send-label-explicit.vipValue")
					item.SendLabelExplicit = types.BoolValue(cv.Bool())
					item.SendLabelExplicitVariable = types.StringNull()
				}
			} else {
				item.SendLabelExplicit = types.BoolNull()
				item.SendLabelExplicitVariable = types.StringNull()
			}
			if cValue := v.Get("as-override.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AsOverride = types.BoolNull()

					cv := v.Get("as-override.vipVariableName")
					item.AsOverrideVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AsOverride = types.BoolNull()
					item.AsOverrideVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("as-override.vipValue")
					item.AsOverride = types.BoolValue(cv.Bool())
					item.AsOverrideVariable = types.StringNull()
				}
			} else {
				item.AsOverride = types.BoolNull()
				item.AsOverrideVariable = types.StringNull()
			}
			if cValue := v.Get("allowas-in.as-number.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AllowAsIn = types.Int64Null()

					cv := v.Get("allowas-in.as-number.vipVariableName")
					item.AllowAsInVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AllowAsIn = types.Int64Null()
					item.AllowAsInVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("allowas-in.as-number.vipValue")
					item.AllowAsIn = types.Int64Value(cv.Int())
					item.AllowAsInVariable = types.StringNull()
				}
			} else {
				item.AllowAsIn = types.Int64Null()
				item.AllowAsInVariable = types.StringNull()
			}
			if cValue := v.Get("address-family.vipValue"); len(cValue.Array()) > 0 {
				item.AddressFamilies = make([]CiscoBGPIpv6NeighborsAddressFamilies, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoBGPIpv6NeighborsAddressFamilies{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("family-type.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.FamilyType = types.StringNull()

						} else if ccValue.String() == "ignore" {
							cItem.FamilyType = types.StringNull()

						} else if ccValue.String() == "constant" {
							ccv := cv.Get("family-type.vipValue")
							cItem.FamilyType = types.StringValue(ccv.String())

						}
					} else {
						cItem.FamilyType = types.StringNull()

					}
					if ccValue := cv.Get("maximum-prefixes.prefix-num.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.MaximumPrefixes = types.Int64Null()

							ccv := cv.Get("maximum-prefixes.prefix-num.vipVariableName")
							cItem.MaximumPrefixesVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.MaximumPrefixes = types.Int64Null()
							cItem.MaximumPrefixesVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("maximum-prefixes.prefix-num.vipValue")
							cItem.MaximumPrefixes = types.Int64Value(ccv.Int())
							cItem.MaximumPrefixesVariable = types.StringNull()
						}
					} else {
						cItem.MaximumPrefixes = types.Int64Null()
						cItem.MaximumPrefixesVariable = types.StringNull()
					}
					if ccValue := cv.Get("maximum-prefixes.threshold.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.MaximumPrefixesThreshold = types.Int64Null()

							ccv := cv.Get("maximum-prefixes.threshold.vipVariableName")
							cItem.MaximumPrefixesThresholdVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.MaximumPrefixesThreshold = types.Int64Null()
							cItem.MaximumPrefixesThresholdVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("maximum-prefixes.threshold.vipValue")
							cItem.MaximumPrefixesThreshold = types.Int64Value(ccv.Int())
							cItem.MaximumPrefixesThresholdVariable = types.StringNull()
						}
					} else {
						cItem.MaximumPrefixesThreshold = types.Int64Null()
						cItem.MaximumPrefixesThresholdVariable = types.StringNull()
					}
					if ccValue := cv.Get("maximum-prefixes.restart.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.MaximumPrefixesRestart = types.Int64Null()

							ccv := cv.Get("maximum-prefixes.restart.vipVariableName")
							cItem.MaximumPrefixesRestartVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.MaximumPrefixesRestart = types.Int64Null()
							cItem.MaximumPrefixesRestartVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("maximum-prefixes.restart.vipValue")
							cItem.MaximumPrefixesRestart = types.Int64Value(ccv.Int())
							cItem.MaximumPrefixesRestartVariable = types.StringNull()
						}
					} else {
						cItem.MaximumPrefixesRestart = types.Int64Null()
						cItem.MaximumPrefixesRestartVariable = types.StringNull()
					}
					if ccValue := cv.Get("maximum-prefixes.warning-only.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.MaximumPrefixesWarningOnly = types.BoolNull()

							ccv := cv.Get("maximum-prefixes.warning-only.vipVariableName")
							cItem.MaximumPrefixesWarningOnlyVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.MaximumPrefixesWarningOnly = types.BoolNull()
							cItem.MaximumPrefixesWarningOnlyVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("maximum-prefixes.warning-only.vipValue")
							cItem.MaximumPrefixesWarningOnly = types.BoolValue(ccv.Bool())
							cItem.MaximumPrefixesWarningOnlyVariable = types.StringNull()
						}
					} else {
						cItem.MaximumPrefixesWarningOnly = types.BoolNull()
						cItem.MaximumPrefixesWarningOnlyVariable = types.StringNull()
					}
					if ccValue := cv.Get("route-policy.vipValue"); len(ccValue.Array()) > 0 {
						cItem.RoutePolicies = make([]CiscoBGPIpv6NeighborsAddressFamiliesRoutePolicies, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := CiscoBGPIpv6NeighborsAddressFamiliesRoutePolicies{}
							if cccValue := ccv.Get("vipOptional"); cccValue.Exists() {
								ccItem.Optional = types.BoolValue(cccValue.Bool())
							} else {
								ccItem.Optional = types.BoolNull()
							}
							if cccValue := ccv.Get("direction.vipType"); cccValue.Exists() {
								if cccValue.String() == "variableName" {
									ccItem.Direction = types.StringNull()

								} else if cccValue.String() == "ignore" {
									ccItem.Direction = types.StringNull()

								} else if cccValue.String() == "constant" {
									cccv := ccv.Get("direction.vipValue")
									ccItem.Direction = types.StringValue(cccv.String())

								}
							} else {
								ccItem.Direction = types.StringNull()

							}
							if cccValue := ccv.Get("pol-name.vipType"); cccValue.Exists() {
								if cccValue.String() == "variableName" {
									ccItem.PolicyName = types.StringNull()

									cccv := ccv.Get("pol-name.vipVariableName")
									ccItem.PolicyNameVariable = types.StringValue(cccv.String())

								} else if cccValue.String() == "ignore" {
									ccItem.PolicyName = types.StringNull()
									ccItem.PolicyNameVariable = types.StringNull()
								} else if cccValue.String() == "constant" {
									cccv := ccv.Get("pol-name.vipValue")
									ccItem.PolicyName = types.StringValue(cccv.String())
									ccItem.PolicyNameVariable = types.StringNull()
								}
							} else {
								ccItem.PolicyName = types.StringNull()
								ccItem.PolicyNameVariable = types.StringNull()
							}
							cItem.RoutePolicies = append(cItem.RoutePolicies, ccItem)
							return true
						})
					} else {
						if len(cItem.RoutePolicies) > 0 {
							cItem.RoutePolicies = []CiscoBGPIpv6NeighborsAddressFamiliesRoutePolicies{}
						}
					}
					item.AddressFamilies = append(item.AddressFamilies, cItem)
					return true
				})
			} else {
				if len(item.AddressFamilies) > 0 {
					item.AddressFamilies = []CiscoBGPIpv6NeighborsAddressFamilies{}
				}
			}
			data.Ipv6Neighbors = append(data.Ipv6Neighbors, item)
			return true
		})
	} else {
		if len(data.Ipv6Neighbors) > 0 {
			data.Ipv6Neighbors = []CiscoBGPIpv6Neighbors{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoBGP) hasChanges(ctx context.Context, state *CiscoBGP) bool {
	hasChanges := false
	if !data.AsNumber.Equal(state.AsNumber) {
		hasChanges = true
	}
	if !data.Shutdown.Equal(state.Shutdown) {
		hasChanges = true
	}
	if !data.RouterId.Equal(state.RouterId) {
		hasChanges = true
	}
	if !data.PropagateAspath.Equal(state.PropagateAspath) {
		hasChanges = true
	}
	if !data.PropagateCommunity.Equal(state.PropagateCommunity) {
		hasChanges = true
	}
	if len(data.Ipv4RouteTargets) != len(state.Ipv4RouteTargets) {
		hasChanges = true
	} else {
		for i := range data.Ipv4RouteTargets {
			if !data.Ipv4RouteTargets[i].VpnId.Equal(state.Ipv4RouteTargets[i].VpnId) {
				hasChanges = true
			}
			if len(data.Ipv4RouteTargets[i].Export) != len(state.Ipv4RouteTargets[i].Export) {
				hasChanges = true
			} else {
				for ii := range data.Ipv4RouteTargets[i].Export {
					if !data.Ipv4RouteTargets[i].Export[ii].AsnIp.Equal(state.Ipv4RouteTargets[i].Export[ii].AsnIp) {
						hasChanges = true
					}
				}
			}
			if len(data.Ipv4RouteTargets[i].Import) != len(state.Ipv4RouteTargets[i].Import) {
				hasChanges = true
			} else {
				for ii := range data.Ipv4RouteTargets[i].Import {
					if !data.Ipv4RouteTargets[i].Import[ii].AsnIp.Equal(state.Ipv4RouteTargets[i].Import[ii].AsnIp) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.Ipv6RouteTargets) != len(state.Ipv6RouteTargets) {
		hasChanges = true
	} else {
		for i := range data.Ipv6RouteTargets {
			if !data.Ipv6RouteTargets[i].VpnId.Equal(state.Ipv6RouteTargets[i].VpnId) {
				hasChanges = true
			}
			if len(data.Ipv6RouteTargets[i].Export) != len(state.Ipv6RouteTargets[i].Export) {
				hasChanges = true
			} else {
				for ii := range data.Ipv6RouteTargets[i].Export {
					if !data.Ipv6RouteTargets[i].Export[ii].AsnIp.Equal(state.Ipv6RouteTargets[i].Export[ii].AsnIp) {
						hasChanges = true
					}
				}
			}
			if len(data.Ipv6RouteTargets[i].Import) != len(state.Ipv6RouteTargets[i].Import) {
				hasChanges = true
			} else {
				for ii := range data.Ipv6RouteTargets[i].Import {
					if !data.Ipv6RouteTargets[i].Import[ii].AsnIp.Equal(state.Ipv6RouteTargets[i].Import[ii].AsnIp) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.MplsInterfaces) != len(state.MplsInterfaces) {
		hasChanges = true
	} else {
		for i := range data.MplsInterfaces {
			if !data.MplsInterfaces[i].InterfaceName.Equal(state.MplsInterfaces[i].InterfaceName) {
				hasChanges = true
			}
		}
	}
	if !data.DistanceExternal.Equal(state.DistanceExternal) {
		hasChanges = true
	}
	if !data.DistanceInternal.Equal(state.DistanceInternal) {
		hasChanges = true
	}
	if !data.DistanceLocal.Equal(state.DistanceLocal) {
		hasChanges = true
	}
	if !data.Keepalive.Equal(state.Keepalive) {
		hasChanges = true
	}
	if !data.Holdtime.Equal(state.Holdtime) {
		hasChanges = true
	}
	if !data.AlwaysCompareMed.Equal(state.AlwaysCompareMed) {
		hasChanges = true
	}
	if !data.DeterministicMed.Equal(state.DeterministicMed) {
		hasChanges = true
	}
	if !data.MissingMedWorst.Equal(state.MissingMedWorst) {
		hasChanges = true
	}
	if !data.CompareRouterId.Equal(state.CompareRouterId) {
		hasChanges = true
	}
	if !data.MultipathRelax.Equal(state.MultipathRelax) {
		hasChanges = true
	}
	if len(data.AddressFamilies) != len(state.AddressFamilies) {
		hasChanges = true
	} else {
		for i := range data.AddressFamilies {
			if !data.AddressFamilies[i].FamilyType.Equal(state.AddressFamilies[i].FamilyType) {
				hasChanges = true
			}
			if len(data.AddressFamilies[i].Ipv4AggregateAddresses) != len(state.AddressFamilies[i].Ipv4AggregateAddresses) {
				hasChanges = true
			} else {
				for ii := range data.AddressFamilies[i].Ipv4AggregateAddresses {
					if !data.AddressFamilies[i].Ipv4AggregateAddresses[ii].Prefix.Equal(state.AddressFamilies[i].Ipv4AggregateAddresses[ii].Prefix) {
						hasChanges = true
					}
					if !data.AddressFamilies[i].Ipv4AggregateAddresses[ii].AsSetPath.Equal(state.AddressFamilies[i].Ipv4AggregateAddresses[ii].AsSetPath) {
						hasChanges = true
					}
					if !data.AddressFamilies[i].Ipv4AggregateAddresses[ii].SummaryOnly.Equal(state.AddressFamilies[i].Ipv4AggregateAddresses[ii].SummaryOnly) {
						hasChanges = true
					}
				}
			}
			if len(data.AddressFamilies[i].Ipv6AggregateAddresses) != len(state.AddressFamilies[i].Ipv6AggregateAddresses) {
				hasChanges = true
			} else {
				for ii := range data.AddressFamilies[i].Ipv6AggregateAddresses {
					if !data.AddressFamilies[i].Ipv6AggregateAddresses[ii].Prefix.Equal(state.AddressFamilies[i].Ipv6AggregateAddresses[ii].Prefix) {
						hasChanges = true
					}
					if !data.AddressFamilies[i].Ipv6AggregateAddresses[ii].AsSetPath.Equal(state.AddressFamilies[i].Ipv6AggregateAddresses[ii].AsSetPath) {
						hasChanges = true
					}
					if !data.AddressFamilies[i].Ipv6AggregateAddresses[ii].SummaryOnly.Equal(state.AddressFamilies[i].Ipv6AggregateAddresses[ii].SummaryOnly) {
						hasChanges = true
					}
				}
			}
			if len(data.AddressFamilies[i].Ipv4Networks) != len(state.AddressFamilies[i].Ipv4Networks) {
				hasChanges = true
			} else {
				for ii := range data.AddressFamilies[i].Ipv4Networks {
					if !data.AddressFamilies[i].Ipv4Networks[ii].Prefix.Equal(state.AddressFamilies[i].Ipv4Networks[ii].Prefix) {
						hasChanges = true
					}
				}
			}
			if len(data.AddressFamilies[i].Ipv6Networks) != len(state.AddressFamilies[i].Ipv6Networks) {
				hasChanges = true
			} else {
				for ii := range data.AddressFamilies[i].Ipv6Networks {
					if !data.AddressFamilies[i].Ipv6Networks[ii].Prefix.Equal(state.AddressFamilies[i].Ipv6Networks[ii].Prefix) {
						hasChanges = true
					}
				}
			}
			if !data.AddressFamilies[i].MaximumPaths.Equal(state.AddressFamilies[i].MaximumPaths) {
				hasChanges = true
			}
			if !data.AddressFamilies[i].DefaultInformationOriginate.Equal(state.AddressFamilies[i].DefaultInformationOriginate) {
				hasChanges = true
			}
			if !data.AddressFamilies[i].TableMapPolicy.Equal(state.AddressFamilies[i].TableMapPolicy) {
				hasChanges = true
			}
			if !data.AddressFamilies[i].TableMapFilter.Equal(state.AddressFamilies[i].TableMapFilter) {
				hasChanges = true
			}
			if len(data.AddressFamilies[i].RedistributeRoutes) != len(state.AddressFamilies[i].RedistributeRoutes) {
				hasChanges = true
			} else {
				for ii := range data.AddressFamilies[i].RedistributeRoutes {
					if !data.AddressFamilies[i].RedistributeRoutes[ii].Protocol.Equal(state.AddressFamilies[i].RedistributeRoutes[ii].Protocol) {
						hasChanges = true
					}
					if !data.AddressFamilies[i].RedistributeRoutes[ii].RoutePolicy.Equal(state.AddressFamilies[i].RedistributeRoutes[ii].RoutePolicy) {
						hasChanges = true
					}
				}
			}
		}
	}
	if len(data.Ipv4Neighbors) != len(state.Ipv4Neighbors) {
		hasChanges = true
	} else {
		for i := range data.Ipv4Neighbors {
			if !data.Ipv4Neighbors[i].Address.Equal(state.Ipv4Neighbors[i].Address) {
				hasChanges = true
			}
			if !data.Ipv4Neighbors[i].Description.Equal(state.Ipv4Neighbors[i].Description) {
				hasChanges = true
			}
			if !data.Ipv4Neighbors[i].Shutdown.Equal(state.Ipv4Neighbors[i].Shutdown) {
				hasChanges = true
			}
			if !data.Ipv4Neighbors[i].RemoteAs.Equal(state.Ipv4Neighbors[i].RemoteAs) {
				hasChanges = true
			}
			if !data.Ipv4Neighbors[i].Keepalive.Equal(state.Ipv4Neighbors[i].Keepalive) {
				hasChanges = true
			}
			if !data.Ipv4Neighbors[i].Holdtime.Equal(state.Ipv4Neighbors[i].Holdtime) {
				hasChanges = true
			}
			if !data.Ipv4Neighbors[i].SourceInterface.Equal(state.Ipv4Neighbors[i].SourceInterface) {
				hasChanges = true
			}
			if !data.Ipv4Neighbors[i].NextHopSelf.Equal(state.Ipv4Neighbors[i].NextHopSelf) {
				hasChanges = true
			}
			if !data.Ipv4Neighbors[i].SendCommunity.Equal(state.Ipv4Neighbors[i].SendCommunity) {
				hasChanges = true
			}
			if !data.Ipv4Neighbors[i].SendExtCommunity.Equal(state.Ipv4Neighbors[i].SendExtCommunity) {
				hasChanges = true
			}
			if !data.Ipv4Neighbors[i].EbgpMultihop.Equal(state.Ipv4Neighbors[i].EbgpMultihop) {
				hasChanges = true
			}
			if !data.Ipv4Neighbors[i].Password.Equal(state.Ipv4Neighbors[i].Password) {
				hasChanges = true
			}
			if !data.Ipv4Neighbors[i].SendLabel.Equal(state.Ipv4Neighbors[i].SendLabel) {
				hasChanges = true
			}
			if !data.Ipv4Neighbors[i].SendLabelExplicit.Equal(state.Ipv4Neighbors[i].SendLabelExplicit) {
				hasChanges = true
			}
			if !data.Ipv4Neighbors[i].AsOverride.Equal(state.Ipv4Neighbors[i].AsOverride) {
				hasChanges = true
			}
			if !data.Ipv4Neighbors[i].AllowAsIn.Equal(state.Ipv4Neighbors[i].AllowAsIn) {
				hasChanges = true
			}
			if len(data.Ipv4Neighbors[i].AddressFamilies) != len(state.Ipv4Neighbors[i].AddressFamilies) {
				hasChanges = true
			} else {
				for ii := range data.Ipv4Neighbors[i].AddressFamilies {
					if !data.Ipv4Neighbors[i].AddressFamilies[ii].FamilyType.Equal(state.Ipv4Neighbors[i].AddressFamilies[ii].FamilyType) {
						hasChanges = true
					}
					if !data.Ipv4Neighbors[i].AddressFamilies[ii].MaximumPrefixes.Equal(state.Ipv4Neighbors[i].AddressFamilies[ii].MaximumPrefixes) {
						hasChanges = true
					}
					if !data.Ipv4Neighbors[i].AddressFamilies[ii].MaximumPrefixesThreshold.Equal(state.Ipv4Neighbors[i].AddressFamilies[ii].MaximumPrefixesThreshold) {
						hasChanges = true
					}
					if !data.Ipv4Neighbors[i].AddressFamilies[ii].MaximumPrefixesRestart.Equal(state.Ipv4Neighbors[i].AddressFamilies[ii].MaximumPrefixesRestart) {
						hasChanges = true
					}
					if !data.Ipv4Neighbors[i].AddressFamilies[ii].MaximumPrefixesWarningOnly.Equal(state.Ipv4Neighbors[i].AddressFamilies[ii].MaximumPrefixesWarningOnly) {
						hasChanges = true
					}
					if len(data.Ipv4Neighbors[i].AddressFamilies[ii].RoutePolicies) != len(state.Ipv4Neighbors[i].AddressFamilies[ii].RoutePolicies) {
						hasChanges = true
					} else {
						for iii := range data.Ipv4Neighbors[i].AddressFamilies[ii].RoutePolicies {
							if !data.Ipv4Neighbors[i].AddressFamilies[ii].RoutePolicies[iii].Direction.Equal(state.Ipv4Neighbors[i].AddressFamilies[ii].RoutePolicies[iii].Direction) {
								hasChanges = true
							}
							if !data.Ipv4Neighbors[i].AddressFamilies[ii].RoutePolicies[iii].PolicyName.Equal(state.Ipv4Neighbors[i].AddressFamilies[ii].RoutePolicies[iii].PolicyName) {
								hasChanges = true
							}
						}
					}
				}
			}
		}
	}
	if len(data.Ipv6Neighbors) != len(state.Ipv6Neighbors) {
		hasChanges = true
	} else {
		for i := range data.Ipv6Neighbors {
			if !data.Ipv6Neighbors[i].Address.Equal(state.Ipv6Neighbors[i].Address) {
				hasChanges = true
			}
			if !data.Ipv6Neighbors[i].Description.Equal(state.Ipv6Neighbors[i].Description) {
				hasChanges = true
			}
			if !data.Ipv6Neighbors[i].Shutdown.Equal(state.Ipv6Neighbors[i].Shutdown) {
				hasChanges = true
			}
			if !data.Ipv6Neighbors[i].RemoteAs.Equal(state.Ipv6Neighbors[i].RemoteAs) {
				hasChanges = true
			}
			if !data.Ipv6Neighbors[i].Keepalive.Equal(state.Ipv6Neighbors[i].Keepalive) {
				hasChanges = true
			}
			if !data.Ipv6Neighbors[i].Holdtime.Equal(state.Ipv6Neighbors[i].Holdtime) {
				hasChanges = true
			}
			if !data.Ipv6Neighbors[i].SourceInterface.Equal(state.Ipv6Neighbors[i].SourceInterface) {
				hasChanges = true
			}
			if !data.Ipv6Neighbors[i].NextHopSelf.Equal(state.Ipv6Neighbors[i].NextHopSelf) {
				hasChanges = true
			}
			if !data.Ipv6Neighbors[i].SendCommunity.Equal(state.Ipv6Neighbors[i].SendCommunity) {
				hasChanges = true
			}
			if !data.Ipv6Neighbors[i].SendExtCommunity.Equal(state.Ipv6Neighbors[i].SendExtCommunity) {
				hasChanges = true
			}
			if !data.Ipv6Neighbors[i].EbgpMultihop.Equal(state.Ipv6Neighbors[i].EbgpMultihop) {
				hasChanges = true
			}
			if !data.Ipv6Neighbors[i].Password.Equal(state.Ipv6Neighbors[i].Password) {
				hasChanges = true
			}
			if !data.Ipv6Neighbors[i].SendLabel.Equal(state.Ipv6Neighbors[i].SendLabel) {
				hasChanges = true
			}
			if !data.Ipv6Neighbors[i].SendLabelExplicit.Equal(state.Ipv6Neighbors[i].SendLabelExplicit) {
				hasChanges = true
			}
			if !data.Ipv6Neighbors[i].AsOverride.Equal(state.Ipv6Neighbors[i].AsOverride) {
				hasChanges = true
			}
			if !data.Ipv6Neighbors[i].AllowAsIn.Equal(state.Ipv6Neighbors[i].AllowAsIn) {
				hasChanges = true
			}
			if len(data.Ipv6Neighbors[i].AddressFamilies) != len(state.Ipv6Neighbors[i].AddressFamilies) {
				hasChanges = true
			} else {
				for ii := range data.Ipv6Neighbors[i].AddressFamilies {
					if !data.Ipv6Neighbors[i].AddressFamilies[ii].FamilyType.Equal(state.Ipv6Neighbors[i].AddressFamilies[ii].FamilyType) {
						hasChanges = true
					}
					if !data.Ipv6Neighbors[i].AddressFamilies[ii].MaximumPrefixes.Equal(state.Ipv6Neighbors[i].AddressFamilies[ii].MaximumPrefixes) {
						hasChanges = true
					}
					if !data.Ipv6Neighbors[i].AddressFamilies[ii].MaximumPrefixesThreshold.Equal(state.Ipv6Neighbors[i].AddressFamilies[ii].MaximumPrefixesThreshold) {
						hasChanges = true
					}
					if !data.Ipv6Neighbors[i].AddressFamilies[ii].MaximumPrefixesRestart.Equal(state.Ipv6Neighbors[i].AddressFamilies[ii].MaximumPrefixesRestart) {
						hasChanges = true
					}
					if !data.Ipv6Neighbors[i].AddressFamilies[ii].MaximumPrefixesWarningOnly.Equal(state.Ipv6Neighbors[i].AddressFamilies[ii].MaximumPrefixesWarningOnly) {
						hasChanges = true
					}
					if len(data.Ipv6Neighbors[i].AddressFamilies[ii].RoutePolicies) != len(state.Ipv6Neighbors[i].AddressFamilies[ii].RoutePolicies) {
						hasChanges = true
					} else {
						for iii := range data.Ipv6Neighbors[i].AddressFamilies[ii].RoutePolicies {
							if !data.Ipv6Neighbors[i].AddressFamilies[ii].RoutePolicies[iii].Direction.Equal(state.Ipv6Neighbors[i].AddressFamilies[ii].RoutePolicies[iii].Direction) {
								hasChanges = true
							}
							if !data.Ipv6Neighbors[i].AddressFamilies[ii].RoutePolicies[iii].PolicyName.Equal(state.Ipv6Neighbors[i].AddressFamilies[ii].RoutePolicies[iii].PolicyName) {
								hasChanges = true
							}
						}
					}
				}
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
