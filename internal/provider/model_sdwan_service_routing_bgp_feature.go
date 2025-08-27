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

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

var MinServiceRoutingBGPUpdateVersion = version.Must(version.NewVersion("20.15.0"))

// Section below is generated&owned by "gen/generator.go". //template:begin types
type ServiceRoutingBGP struct {
	Id                             types.String                              `tfsdk:"id"`
	Version                        types.Int64                               `tfsdk:"version"`
	Name                           types.String                              `tfsdk:"name"`
	Description                    types.String                              `tfsdk:"description"`
	FeatureProfileId               types.String                              `tfsdk:"feature_profile_id"`
	AsNumber                       types.Int64                               `tfsdk:"as_number"`
	AsNumberVariable               types.String                              `tfsdk:"as_number_variable"`
	RouterId                       types.String                              `tfsdk:"router_id"`
	RouterIdVariable               types.String                              `tfsdk:"router_id_variable"`
	PropagateAsPath                types.Bool                                `tfsdk:"propagate_as_path"`
	PropagateAsPathVariable        types.String                              `tfsdk:"propagate_as_path_variable"`
	PropagateCommunity             types.Bool                                `tfsdk:"propagate_community"`
	PropagateCommunityVariable     types.String                              `tfsdk:"propagate_community_variable"`
	ExternalRoutesDistance         types.Int64                               `tfsdk:"external_routes_distance"`
	ExternalRoutesDistanceVariable types.String                              `tfsdk:"external_routes_distance_variable"`
	InternalRoutesDistance         types.Int64                               `tfsdk:"internal_routes_distance"`
	InternalRoutesDistanceVariable types.String                              `tfsdk:"internal_routes_distance_variable"`
	LocalRoutesDistance            types.Int64                               `tfsdk:"local_routes_distance"`
	LocalRoutesDistanceVariable    types.String                              `tfsdk:"local_routes_distance_variable"`
	KeepaliveTime                  types.Int64                               `tfsdk:"keepalive_time"`
	KeepaliveTimeVariable          types.String                              `tfsdk:"keepalive_time_variable"`
	HoldTime                       types.Int64                               `tfsdk:"hold_time"`
	HoldTimeVariable               types.String                              `tfsdk:"hold_time_variable"`
	AlwaysCompareMed               types.Bool                                `tfsdk:"always_compare_med"`
	AlwaysCompareMedVariable       types.String                              `tfsdk:"always_compare_med_variable"`
	DeterministicMed               types.Bool                                `tfsdk:"deterministic_med"`
	DeterministicMedVariable       types.String                              `tfsdk:"deterministic_med_variable"`
	MissingMedAsWorst              types.Bool                                `tfsdk:"missing_med_as_worst"`
	MissingMedAsWorstVariable      types.String                              `tfsdk:"missing_med_as_worst_variable"`
	CompareRouterId                types.Bool                                `tfsdk:"compare_router_id"`
	CompareRouterIdVariable        types.String                              `tfsdk:"compare_router_id_variable"`
	MultipathRelax                 types.Bool                                `tfsdk:"multipath_relax"`
	MultipathRelaxVariable         types.String                              `tfsdk:"multipath_relax_variable"`
	Ipv4Neighbors                  []ServiceRoutingBGPIpv4Neighbors          `tfsdk:"ipv4_neighbors"`
	Ipv6Neighbors                  []ServiceRoutingBGPIpv6Neighbors          `tfsdk:"ipv6_neighbors"`
	Ipv4AggregateAddresses         []ServiceRoutingBGPIpv4AggregateAddresses `tfsdk:"ipv4_aggregate_addresses"`
	Ipv4Networks                   []ServiceRoutingBGPIpv4Networks           `tfsdk:"ipv4_networks"`
	Ipv4EibgpMaximumPaths          types.Int64                               `tfsdk:"ipv4_eibgp_maximum_paths"`
	Ipv4EibgpMaximumPathsVariable  types.String                              `tfsdk:"ipv4_eibgp_maximum_paths_variable"`
	Ipv4Originate                  types.Bool                                `tfsdk:"ipv4_originate"`
	Ipv4OriginateVariable          types.String                              `tfsdk:"ipv4_originate_variable"`
	Ipv4TableMapRoutePolicyId      types.String                              `tfsdk:"ipv4_table_map_route_policy_id"`
	Ipv4TableMapFilter             types.Bool                                `tfsdk:"ipv4_table_map_filter"`
	Ipv4TableMapFilterVariable     types.String                              `tfsdk:"ipv4_table_map_filter_variable"`
	Ipv4Redistributes              []ServiceRoutingBGPIpv4Redistributes      `tfsdk:"ipv4_redistributes"`
	Ipv6AggregateAddresses         []ServiceRoutingBGPIpv6AggregateAddresses `tfsdk:"ipv6_aggregate_addresses"`
	Ipv6Networks                   []ServiceRoutingBGPIpv6Networks           `tfsdk:"ipv6_networks"`
	Ipv6EibgpMaximumPaths          types.Int64                               `tfsdk:"ipv6_eibgp_maximum_paths"`
	Ipv6EibgpMaximumPathsVariable  types.String                              `tfsdk:"ipv6_eibgp_maximum_paths_variable"`
	Ipv6Originate                  types.Bool                                `tfsdk:"ipv6_originate"`
	Ipv6OriginateVariable          types.String                              `tfsdk:"ipv6_originate_variable"`
	Ipv6TableMapRoutePolicyId      types.String                              `tfsdk:"ipv6_table_map_route_policy_id"`
	Ipv6TableMapFilter             types.Bool                                `tfsdk:"ipv6_table_map_filter"`
	Ipv6TableMapFilterVariable     types.String                              `tfsdk:"ipv6_table_map_filter_variable"`
	Ipv6Redistributes              []ServiceRoutingBGPIpv6Redistributes      `tfsdk:"ipv6_redistributes"`
}

type ServiceRoutingBGPIpv4Neighbors struct {
	Address                       types.String                                    `tfsdk:"address"`
	AddressVariable               types.String                                    `tfsdk:"address_variable"`
	Description                   types.String                                    `tfsdk:"description"`
	DescriptionVariable           types.String                                    `tfsdk:"description_variable"`
	Shutdown                      types.Bool                                      `tfsdk:"shutdown"`
	ShutdownVariable              types.String                                    `tfsdk:"shutdown_variable"`
	RemoteAs                      types.Int64                                     `tfsdk:"remote_as"`
	RemoteAsVariable              types.String                                    `tfsdk:"remote_as_variable"`
	LocalAs                       types.Int64                                     `tfsdk:"local_as"`
	LocalAsVariable               types.String                                    `tfsdk:"local_as_variable"`
	KeepaliveTime                 types.Int64                                     `tfsdk:"keepalive_time"`
	KeepaliveTimeVariable         types.String                                    `tfsdk:"keepalive_time_variable"`
	HoldTime                      types.Int64                                     `tfsdk:"hold_time"`
	HoldTimeVariable              types.String                                    `tfsdk:"hold_time_variable"`
	UpdateSourceInterface         types.String                                    `tfsdk:"update_source_interface"`
	UpdateSourceInterfaceVariable types.String                                    `tfsdk:"update_source_interface_variable"`
	NextHopSelf                   types.Bool                                      `tfsdk:"next_hop_self"`
	NextHopSelfVariable           types.String                                    `tfsdk:"next_hop_self_variable"`
	SendCommunity                 types.Bool                                      `tfsdk:"send_community"`
	SendCommunityVariable         types.String                                    `tfsdk:"send_community_variable"`
	SendExtendedCommunity         types.Bool                                      `tfsdk:"send_extended_community"`
	SendExtendedCommunityVariable types.String                                    `tfsdk:"send_extended_community_variable"`
	EbgpMultihop                  types.Int64                                     `tfsdk:"ebgp_multihop"`
	EbgpMultihopVariable          types.String                                    `tfsdk:"ebgp_multihop_variable"`
	Password                      types.String                                    `tfsdk:"password"`
	PasswordVariable              types.String                                    `tfsdk:"password_variable"`
	SendLabel                     types.Bool                                      `tfsdk:"send_label"`
	SendLabelVariable             types.String                                    `tfsdk:"send_label_variable"`
	AsOverride                    types.Bool                                      `tfsdk:"as_override"`
	AsOverrideVariable            types.String                                    `tfsdk:"as_override_variable"`
	AllowasInNumber               types.Int64                                     `tfsdk:"allowas_in_number"`
	AllowasInNumberVariable       types.String                                    `tfsdk:"allowas_in_number_variable"`
	AddressFamilies               []ServiceRoutingBGPIpv4NeighborsAddressFamilies `tfsdk:"address_families"`
}

type ServiceRoutingBGPIpv6Neighbors struct {
	Address                       types.String                                    `tfsdk:"address"`
	AddressVariable               types.String                                    `tfsdk:"address_variable"`
	Description                   types.String                                    `tfsdk:"description"`
	DescriptionVariable           types.String                                    `tfsdk:"description_variable"`
	Shutdown                      types.Bool                                      `tfsdk:"shutdown"`
	ShutdownVariable              types.String                                    `tfsdk:"shutdown_variable"`
	RemoteAs                      types.Int64                                     `tfsdk:"remote_as"`
	RemoteAsVariable              types.String                                    `tfsdk:"remote_as_variable"`
	LocalAs                       types.Int64                                     `tfsdk:"local_as"`
	LocalAsVariable               types.String                                    `tfsdk:"local_as_variable"`
	KeepaliveTime                 types.Int64                                     `tfsdk:"keepalive_time"`
	KeepaliveTimeVariable         types.String                                    `tfsdk:"keepalive_time_variable"`
	HoldTime                      types.Int64                                     `tfsdk:"hold_time"`
	HoldTimeVariable              types.String                                    `tfsdk:"hold_time_variable"`
	UpdateSourceInterface         types.String                                    `tfsdk:"update_source_interface"`
	UpdateSourceInterfaceVariable types.String                                    `tfsdk:"update_source_interface_variable"`
	NextHopSelf                   types.Bool                                      `tfsdk:"next_hop_self"`
	NextHopSelfVariable           types.String                                    `tfsdk:"next_hop_self_variable"`
	SendCommunity                 types.Bool                                      `tfsdk:"send_community"`
	SendCommunityVariable         types.String                                    `tfsdk:"send_community_variable"`
	SendExtendedCommunity         types.Bool                                      `tfsdk:"send_extended_community"`
	SendExtendedCommunityVariable types.String                                    `tfsdk:"send_extended_community_variable"`
	EbgpMultihop                  types.Int64                                     `tfsdk:"ebgp_multihop"`
	EbgpMultihopVariable          types.String                                    `tfsdk:"ebgp_multihop_variable"`
	Password                      types.String                                    `tfsdk:"password"`
	PasswordVariable              types.String                                    `tfsdk:"password_variable"`
	AsOverride                    types.Bool                                      `tfsdk:"as_override"`
	AsOverrideVariable            types.String                                    `tfsdk:"as_override_variable"`
	AllowasInNumber               types.Int64                                     `tfsdk:"allowas_in_number"`
	AllowasInNumberVariable       types.String                                    `tfsdk:"allowas_in_number_variable"`
	AddressFamilies               []ServiceRoutingBGPIpv6NeighborsAddressFamilies `tfsdk:"address_families"`
}

type ServiceRoutingBGPIpv4AggregateAddresses struct {
	NetworkAddress         types.String `tfsdk:"network_address"`
	NetworkAddressVariable types.String `tfsdk:"network_address_variable"`
	SubnetMask             types.String `tfsdk:"subnet_mask"`
	SubnetMaskVariable     types.String `tfsdk:"subnet_mask_variable"`
	AsSetPath              types.Bool   `tfsdk:"as_set_path"`
	AsSetPathVariable      types.String `tfsdk:"as_set_path_variable"`
	SummaryOnly            types.Bool   `tfsdk:"summary_only"`
	SummaryOnlyVariable    types.String `tfsdk:"summary_only_variable"`
}

type ServiceRoutingBGPIpv4Networks struct {
	NetworkAddress         types.String `tfsdk:"network_address"`
	NetworkAddressVariable types.String `tfsdk:"network_address_variable"`
	SubnetMask             types.String `tfsdk:"subnet_mask"`
	SubnetMaskVariable     types.String `tfsdk:"subnet_mask_variable"`
}

type ServiceRoutingBGPIpv4Redistributes struct {
	Protocol           types.String `tfsdk:"protocol"`
	ProtocolVariable   types.String `tfsdk:"protocol_variable"`
	RoutePolicyId      types.String `tfsdk:"route_policy_id"`
	TranslateRibMetric types.Bool   `tfsdk:"translate_rib_metric"`
}

type ServiceRoutingBGPIpv6AggregateAddresses struct {
	AggregatePrefix         types.String `tfsdk:"aggregate_prefix"`
	AggregatePrefixVariable types.String `tfsdk:"aggregate_prefix_variable"`
	AsSetPath               types.Bool   `tfsdk:"as_set_path"`
	AsSetPathVariable       types.String `tfsdk:"as_set_path_variable"`
	SummaryOnly             types.Bool   `tfsdk:"summary_only"`
	SummaryOnlyVariable     types.String `tfsdk:"summary_only_variable"`
}

type ServiceRoutingBGPIpv6Networks struct {
	NetworkPrefix         types.String `tfsdk:"network_prefix"`
	NetworkPrefixVariable types.String `tfsdk:"network_prefix_variable"`
}

type ServiceRoutingBGPIpv6Redistributes struct {
	Protocol           types.String `tfsdk:"protocol"`
	ProtocolVariable   types.String `tfsdk:"protocol_variable"`
	RoutePolicyId      types.String `tfsdk:"route_policy_id"`
	TranslateRibMetric types.Bool   `tfsdk:"translate_rib_metric"`
}

type ServiceRoutingBGPIpv4NeighborsAddressFamilies struct {
	FamilyType                                types.String `tfsdk:"family_type"`
	PolicyType                                types.String `tfsdk:"policy_type"`
	RestartMaxNumberOfPrefixes                types.Int64  `tfsdk:"restart_max_number_of_prefixes"`
	RestartMaxNumberOfPrefixesVariable        types.String `tfsdk:"restart_max_number_of_prefixes_variable"`
	RestartThreshold                          types.Int64  `tfsdk:"restart_threshold"`
	RestartThresholdVariable                  types.String `tfsdk:"restart_threshold_variable"`
	RestartInterval                           types.Int64  `tfsdk:"restart_interval"`
	RestartIntervalVariable                   types.String `tfsdk:"restart_interval_variable"`
	WarningMessageMaxNumberOfPrefixes         types.Int64  `tfsdk:"warning_message_max_number_of_prefixes"`
	WarningMessageMaxNumberOfPrefixesVariable types.String `tfsdk:"warning_message_max_number_of_prefixes_variable"`
	WarningMessageThreshold                   types.Int64  `tfsdk:"warning_message_threshold"`
	WarningMessageThresholdVariable           types.String `tfsdk:"warning_message_threshold_variable"`
	DisablePeerMaxNumberOfPrefixes            types.Int64  `tfsdk:"disable_peer_max_number_of_prefixes"`
	DisablePeerMaxNumberOfPrefixesVariable    types.String `tfsdk:"disable_peer_max_number_of_prefixes_variable"`
	DisablePeerThreshold                      types.Int64  `tfsdk:"disable_peer_threshold"`
	DisablePeerThresholdVariable              types.String `tfsdk:"disable_peer_threshold_variable"`
	InRoutePolicyId                           types.String `tfsdk:"in_route_policy_id"`
	OutRoutePolicyId                          types.String `tfsdk:"out_route_policy_id"`
}

type ServiceRoutingBGPIpv6NeighborsAddressFamilies struct {
	FamilyType                  types.String `tfsdk:"family_type"`
	MaxNumberOfPrefixes         types.Int64  `tfsdk:"max_number_of_prefixes"`
	MaxNumberOfPrefixesVariable types.String `tfsdk:"max_number_of_prefixes_variable"`
	Threshold                   types.Int64  `tfsdk:"threshold"`
	ThresholdVariable           types.String `tfsdk:"threshold_variable"`
	PolicyType                  types.String `tfsdk:"policy_type"`
	RestartInterval             types.Int64  `tfsdk:"restart_interval"`
	RestartIntervalVariable     types.String `tfsdk:"restart_interval_variable"`
	InRoutePolicyId             types.String `tfsdk:"in_route_policy_id"`
	OutRoutePolicyId            types.String `tfsdk:"out_route_policy_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data ServiceRoutingBGP) getModel() string {
	return "service_routing_bgp"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data ServiceRoutingBGP) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/service/%v/routing/bgp", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

func (data ServiceRoutingBGP) toBody(ctx context.Context, version *version.Version) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.AsNumberVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"asNum.optionType", "variable")
			body, _ = sjson.Set(body, path+"asNum.value", data.AsNumberVariable.ValueString())
		}
	} else if !data.AsNumber.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"asNum.optionType", "global")
			body, _ = sjson.Set(body, path+"asNum.value", data.AsNumber.ValueInt64())
		}
	}

	if !data.RouterIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"routerId.optionType", "variable")
			body, _ = sjson.Set(body, path+"routerId.value", data.RouterIdVariable.ValueString())
		}
	} else if data.RouterId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"routerId.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"routerId.optionType", "global")
			body, _ = sjson.Set(body, path+"routerId.value", data.RouterId.ValueString())
		}
	}

	if !data.PropagateAsPathVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"propagateAspath.optionType", "variable")
			body, _ = sjson.Set(body, path+"propagateAspath.value", data.PropagateAsPathVariable.ValueString())
		}
	} else if data.PropagateAsPath.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"propagateAspath.optionType", "default")
			body, _ = sjson.Set(body, path+"propagateAspath.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"propagateAspath.optionType", "global")
			body, _ = sjson.Set(body, path+"propagateAspath.value", data.PropagateAsPath.ValueBool())
		}
	}

	if !data.PropagateCommunityVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"propagateCommunity.optionType", "variable")
			body, _ = sjson.Set(body, path+"propagateCommunity.value", data.PropagateCommunityVariable.ValueString())
		}
	} else if data.PropagateCommunity.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"propagateCommunity.optionType", "default")
			body, _ = sjson.Set(body, path+"propagateCommunity.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"propagateCommunity.optionType", "global")
			body, _ = sjson.Set(body, path+"propagateCommunity.value", data.PropagateCommunity.ValueBool())
		}
	}

	if !data.ExternalRoutesDistanceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"external.optionType", "variable")
			body, _ = sjson.Set(body, path+"external.value", data.ExternalRoutesDistanceVariable.ValueString())
		}
	} else if data.ExternalRoutesDistance.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"external.optionType", "default")
			body, _ = sjson.Set(body, path+"external.value", 20)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"external.optionType", "global")
			body, _ = sjson.Set(body, path+"external.value", data.ExternalRoutesDistance.ValueInt64())
		}
	}

	if !data.InternalRoutesDistanceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"internal.optionType", "variable")
			body, _ = sjson.Set(body, path+"internal.value", data.InternalRoutesDistanceVariable.ValueString())
		}
	} else if data.InternalRoutesDistance.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"internal.optionType", "default")
			body, _ = sjson.Set(body, path+"internal.value", 200)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"internal.optionType", "global")
			body, _ = sjson.Set(body, path+"internal.value", data.InternalRoutesDistance.ValueInt64())
		}
	}

	if !data.LocalRoutesDistanceVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"local.optionType", "variable")
			body, _ = sjson.Set(body, path+"local.value", data.LocalRoutesDistanceVariable.ValueString())
		}
	} else if data.LocalRoutesDistance.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"local.optionType", "default")
			body, _ = sjson.Set(body, path+"local.value", 20)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"local.optionType", "global")
			body, _ = sjson.Set(body, path+"local.value", data.LocalRoutesDistance.ValueInt64())
		}
	}

	if !data.KeepaliveTimeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"keepalive.optionType", "variable")
			body, _ = sjson.Set(body, path+"keepalive.value", data.KeepaliveTimeVariable.ValueString())
		}
	} else if data.KeepaliveTime.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"keepalive.optionType", "default")
			body, _ = sjson.Set(body, path+"keepalive.value", 60)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"keepalive.optionType", "global")
			body, _ = sjson.Set(body, path+"keepalive.value", data.KeepaliveTime.ValueInt64())
		}
	}

	if !data.HoldTimeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"holdtime.optionType", "variable")
			body, _ = sjson.Set(body, path+"holdtime.value", data.HoldTimeVariable.ValueString())
		}
	} else if data.HoldTime.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"holdtime.optionType", "default")
			body, _ = sjson.Set(body, path+"holdtime.value", 180)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"holdtime.optionType", "global")
			body, _ = sjson.Set(body, path+"holdtime.value", data.HoldTime.ValueInt64())
		}
	}

	if !data.AlwaysCompareMedVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"alwaysCompare.optionType", "variable")
			body, _ = sjson.Set(body, path+"alwaysCompare.value", data.AlwaysCompareMedVariable.ValueString())
		}
	} else if data.AlwaysCompareMed.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"alwaysCompare.optionType", "default")
			body, _ = sjson.Set(body, path+"alwaysCompare.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"alwaysCompare.optionType", "global")
			body, _ = sjson.Set(body, path+"alwaysCompare.value", data.AlwaysCompareMed.ValueBool())
		}
	}

	if !data.DeterministicMedVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"deterministic.optionType", "variable")
			body, _ = sjson.Set(body, path+"deterministic.value", data.DeterministicMedVariable.ValueString())
		}
	} else if data.DeterministicMed.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"deterministic.optionType", "default")
			body, _ = sjson.Set(body, path+"deterministic.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"deterministic.optionType", "global")
			body, _ = sjson.Set(body, path+"deterministic.value", data.DeterministicMed.ValueBool())
		}
	}

	if !data.MissingMedAsWorstVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"missingAsWorst.optionType", "variable")
			body, _ = sjson.Set(body, path+"missingAsWorst.value", data.MissingMedAsWorstVariable.ValueString())
		}
	} else if data.MissingMedAsWorst.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"missingAsWorst.optionType", "default")
			body, _ = sjson.Set(body, path+"missingAsWorst.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"missingAsWorst.optionType", "global")
			body, _ = sjson.Set(body, path+"missingAsWorst.value", data.MissingMedAsWorst.ValueBool())
		}
	}

	if !data.CompareRouterIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"compareRouterId.optionType", "variable")
			body, _ = sjson.Set(body, path+"compareRouterId.value", data.CompareRouterIdVariable.ValueString())
		}
	} else if data.CompareRouterId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"compareRouterId.optionType", "default")
			body, _ = sjson.Set(body, path+"compareRouterId.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"compareRouterId.optionType", "global")
			body, _ = sjson.Set(body, path+"compareRouterId.value", data.CompareRouterId.ValueBool())
		}
	}

	if !data.MultipathRelaxVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"multipathRelax.optionType", "variable")
			body, _ = sjson.Set(body, path+"multipathRelax.value", data.MultipathRelaxVariable.ValueString())
		}
	} else if data.MultipathRelax.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"multipathRelax.optionType", "default")
			body, _ = sjson.Set(body, path+"multipathRelax.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"multipathRelax.optionType", "global")
			body, _ = sjson.Set(body, path+"multipathRelax.value", data.MultipathRelax.ValueBool())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"neighbor", []interface{}{})
		for _, item := range data.Ipv4Neighbors {
			itemBody := ""

			if !item.AddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.AddressVariable.ValueString())
				}
			} else if !item.Address.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.Address.ValueString())
				}
			}

			if !item.DescriptionVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "description.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "description.value", item.DescriptionVariable.ValueString())
				}
			} else if item.Description.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "description.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "description.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "description.value", item.Description.ValueString())
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

			if !item.RemoteAsVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "remoteAs.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "remoteAs.value", item.RemoteAsVariable.ValueString())
				}
			} else if !item.RemoteAs.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "remoteAs.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "remoteAs.value", item.RemoteAs.ValueInt64())
				}
			}

			if !item.LocalAsVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "localAs.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "localAs.value", item.LocalAsVariable.ValueString())
				}
			} else if item.LocalAs.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "localAs.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "localAs.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "localAs.value", item.LocalAs.ValueInt64())
				}
			}

			if !item.KeepaliveTimeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "keepalive.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "keepalive.value", item.KeepaliveTimeVariable.ValueString())
				}
			} else if item.KeepaliveTime.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "keepalive.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "keepalive.value", 60)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "keepalive.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "keepalive.value", item.KeepaliveTime.ValueInt64())
				}
			}

			if !item.HoldTimeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "holdtime.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "holdtime.value", item.HoldTimeVariable.ValueString())
				}
			} else if item.HoldTime.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "holdtime.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "holdtime.value", 180)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "holdtime.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "holdtime.value", item.HoldTime.ValueInt64())
				}
			}

			if !item.UpdateSourceInterfaceVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ifName.value", item.UpdateSourceInterfaceVariable.ValueString())
				}
			} else if item.UpdateSourceInterface.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifName.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ifName.value", item.UpdateSourceInterface.ValueString())
				}
			}

			if !item.NextHopSelfVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nextHopSelf.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "nextHopSelf.value", item.NextHopSelfVariable.ValueString())
				}
			} else if item.NextHopSelf.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nextHopSelf.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "nextHopSelf.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nextHopSelf.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "nextHopSelf.value", item.NextHopSelf.ValueBool())
				}
			}

			if !item.SendCommunityVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sendCommunity.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sendCommunity.value", item.SendCommunityVariable.ValueString())
				}
			} else if item.SendCommunity.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sendCommunity.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "sendCommunity.value", true)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sendCommunity.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sendCommunity.value", item.SendCommunity.ValueBool())
				}
			}

			if !item.SendExtendedCommunityVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sendExtCommunity.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sendExtCommunity.value", item.SendExtendedCommunityVariable.ValueString())
				}
			} else if item.SendExtendedCommunity.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sendExtCommunity.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "sendExtCommunity.value", true)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sendExtCommunity.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sendExtCommunity.value", item.SendExtendedCommunity.ValueBool())
				}
			}

			if !item.EbgpMultihopVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ebgpMultihop.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ebgpMultihop.value", item.EbgpMultihopVariable.ValueString())
				}
			} else if item.EbgpMultihop.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ebgpMultihop.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "ebgpMultihop.value", 1)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ebgpMultihop.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ebgpMultihop.value", item.EbgpMultihop.ValueInt64())
				}
			}

			if !item.PasswordVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "password.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "password.value", item.PasswordVariable.ValueString())
				}
			} else if item.Password.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "password.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "password.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "password.value", item.Password.ValueString())
				}
			}

			if !item.SendLabelVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sendLabel.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sendLabel.value", item.SendLabelVariable.ValueString())
				}
			} else if item.SendLabel.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sendLabel.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "sendLabel.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sendLabel.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sendLabel.value", item.SendLabel.ValueBool())
				}
			}

			if !item.AsOverrideVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asOverride.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "asOverride.value", item.AsOverrideVariable.ValueString())
				}
			} else if item.AsOverride.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asOverride.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "asOverride.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asOverride.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "asOverride.value", item.AsOverride.ValueBool())
				}
			}

			if !item.AllowasInNumberVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asNumber.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "asNumber.value", item.AllowasInNumberVariable.ValueString())
				}
			} else if item.AllowasInNumber.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asNumber.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asNumber.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "asNumber.value", item.AllowasInNumber.ValueInt64())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "addressFamily", []interface{}{})
				for _, childItem := range item.AddressFamilies {
					itemChildBody := ""
					if !childItem.FamilyType.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "familyType.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "familyType.value", childItem.FamilyType.ValueString())
						}
					}
					if !childItem.PolicyType.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.policyType.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.policyType.value", childItem.PolicyType.ValueString())
						}
					}

					if !childItem.RestartMaxNumberOfPrefixesVariable.IsNull() {
						if true && childItem.PolicyType.ValueString() == "restart" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.prefixNum.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.prefixNum.value", childItem.RestartMaxNumberOfPrefixesVariable.ValueString())
						}
					} else if !childItem.RestartMaxNumberOfPrefixes.IsNull() {
						if true && childItem.PolicyType.ValueString() == "restart" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.prefixNum.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.prefixNum.value", childItem.RestartMaxNumberOfPrefixes.ValueInt64())
						}
					}

					if !childItem.RestartThresholdVariable.IsNull() {
						if true && childItem.PolicyType.ValueString() == "restart" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.value", childItem.RestartThresholdVariable.ValueString())
						}
					} else if childItem.RestartThreshold.IsNull() {
						if true && childItem.PolicyType.ValueString() == "restart" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.value", 75)
						}
					} else {
						if true && childItem.PolicyType.ValueString() == "restart" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.value", childItem.RestartThreshold.ValueInt64())
						}
					}

					if !childItem.RestartIntervalVariable.IsNull() {
						if true && childItem.PolicyType.ValueString() == "restart" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.restartInterval.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.restartInterval.value", childItem.RestartIntervalVariable.ValueString())
						}
					} else if !childItem.RestartInterval.IsNull() {
						if true && childItem.PolicyType.ValueString() == "restart" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.restartInterval.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.restartInterval.value", childItem.RestartInterval.ValueInt64())
						}
					}

					if !childItem.WarningMessageMaxNumberOfPrefixesVariable.IsNull() {
						if true && childItem.PolicyType.ValueString() == "warning-only" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.prefixNum.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.prefixNum.value", childItem.WarningMessageMaxNumberOfPrefixesVariable.ValueString())
						}
					} else if !childItem.WarningMessageMaxNumberOfPrefixes.IsNull() {
						if true && childItem.PolicyType.ValueString() == "warning-only" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.prefixNum.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.prefixNum.value", childItem.WarningMessageMaxNumberOfPrefixes.ValueInt64())
						}
					}

					if !childItem.WarningMessageThresholdVariable.IsNull() {
						if true && childItem.PolicyType.ValueString() == "warning-only" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.value", childItem.WarningMessageThresholdVariable.ValueString())
						}
					} else if childItem.WarningMessageThreshold.IsNull() {
						if true && childItem.PolicyType.ValueString() == "warning-only" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.value", 75)
						}
					} else {
						if true && childItem.PolicyType.ValueString() == "warning-only" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.value", childItem.WarningMessageThreshold.ValueInt64())
						}
					}

					if !childItem.DisablePeerMaxNumberOfPrefixesVariable.IsNull() {
						if true && childItem.PolicyType.ValueString() == "disable-peer" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.prefixNum.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.prefixNum.value", childItem.DisablePeerMaxNumberOfPrefixesVariable.ValueString())
						}
					} else if !childItem.DisablePeerMaxNumberOfPrefixes.IsNull() {
						if true && childItem.PolicyType.ValueString() == "disable-peer" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.prefixNum.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.prefixNum.value", childItem.DisablePeerMaxNumberOfPrefixes.ValueInt64())
						}
					}

					if !childItem.DisablePeerThresholdVariable.IsNull() {
						if true && childItem.PolicyType.ValueString() == "disable-peer" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.value", childItem.DisablePeerThresholdVariable.ValueString())
						}
					} else if childItem.DisablePeerThreshold.IsNull() {
						if true && childItem.PolicyType.ValueString() == "disable-peer" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.value", 75)
						}
					} else {
						if true && childItem.PolicyType.ValueString() == "disable-peer" {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.value", childItem.DisablePeerThreshold.ValueInt64())
						}
					}
					if !childItem.InRoutePolicyId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "inRoutePolicy.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "inRoutePolicy.refId.value", childItem.InRoutePolicyId.ValueString())
						}
					}
					if !childItem.OutRoutePolicyId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "outRoutePolicy.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "outRoutePolicy.refId.value", childItem.OutRoutePolicyId.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "addressFamily.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"neighbor.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"ipv6Neighbor", []interface{}{})
		for _, item := range data.Ipv6Neighbors {
			itemBody := ""

			if !item.AddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.AddressVariable.ValueString())
				}
			} else if !item.Address.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.Address.ValueString())
				}
			}

			if !item.DescriptionVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "description.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "description.value", item.DescriptionVariable.ValueString())
				}
			} else if item.Description.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "description.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "description.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "description.value", item.Description.ValueString())
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

			if !item.RemoteAsVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "remoteAs.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "remoteAs.value", item.RemoteAsVariable.ValueString())
				}
			} else if !item.RemoteAs.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "remoteAs.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "remoteAs.value", item.RemoteAs.ValueInt64())
				}
			}

			if !item.LocalAsVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "localAs.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "localAs.value", item.LocalAsVariable.ValueString())
				}
			} else if item.LocalAs.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "localAs.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "localAs.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "localAs.value", item.LocalAs.ValueInt64())
				}
			}

			if !item.KeepaliveTimeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "keepalive.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "keepalive.value", item.KeepaliveTimeVariable.ValueString())
				}
			} else if item.KeepaliveTime.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "keepalive.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "keepalive.value", 60)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "keepalive.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "keepalive.value", item.KeepaliveTime.ValueInt64())
				}
			}

			if !item.HoldTimeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "holdtime.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "holdtime.value", item.HoldTimeVariable.ValueString())
				}
			} else if item.HoldTime.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "holdtime.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "holdtime.value", 180)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "holdtime.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "holdtime.value", item.HoldTime.ValueInt64())
				}
			}

			if !item.UpdateSourceInterfaceVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifName.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ifName.value", item.UpdateSourceInterfaceVariable.ValueString())
				}
			} else if item.UpdateSourceInterface.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifName.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ifName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ifName.value", item.UpdateSourceInterface.ValueString())
				}
			}

			if !item.NextHopSelfVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nextHopSelf.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "nextHopSelf.value", item.NextHopSelfVariable.ValueString())
				}
			} else if item.NextHopSelf.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nextHopSelf.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "nextHopSelf.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "nextHopSelf.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "nextHopSelf.value", item.NextHopSelf.ValueBool())
				}
			}

			if !item.SendCommunityVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sendCommunity.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sendCommunity.value", item.SendCommunityVariable.ValueString())
				}
			} else if item.SendCommunity.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sendCommunity.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "sendCommunity.value", true)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sendCommunity.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sendCommunity.value", item.SendCommunity.ValueBool())
				}
			}

			if !item.SendExtendedCommunityVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sendExtCommunity.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "sendExtCommunity.value", item.SendExtendedCommunityVariable.ValueString())
				}
			} else if item.SendExtendedCommunity.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sendExtCommunity.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "sendExtCommunity.value", true)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sendExtCommunity.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sendExtCommunity.value", item.SendExtendedCommunity.ValueBool())
				}
			}

			if !item.EbgpMultihopVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ebgpMultihop.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "ebgpMultihop.value", item.EbgpMultihopVariable.ValueString())
				}
			} else if item.EbgpMultihop.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ebgpMultihop.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "ebgpMultihop.value", 1)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "ebgpMultihop.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "ebgpMultihop.value", item.EbgpMultihop.ValueInt64())
				}
			}

			if !item.PasswordVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "password.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "password.value", item.PasswordVariable.ValueString())
				}
			} else if item.Password.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "password.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "password.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "password.value", item.Password.ValueString())
				}
			}

			if !item.AsOverrideVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asOverride.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "asOverride.value", item.AsOverrideVariable.ValueString())
				}
			} else if item.AsOverride.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asOverride.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "asOverride.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asOverride.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "asOverride.value", item.AsOverride.ValueBool())
				}
			}

			if !item.AllowasInNumberVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asNumber.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "asNumber.value", item.AllowasInNumberVariable.ValueString())
				}
			} else if item.AllowasInNumber.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asNumber.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asNumber.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "asNumber.value", item.AllowasInNumber.ValueInt64())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "addressFamily", []interface{}{})
				for _, childItem := range item.AddressFamilies {
					itemChildBody := ""
					if !childItem.FamilyType.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "familyType.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "familyType.value", childItem.FamilyType.ValueString())
						}
					}

					if !childItem.MaxNumberOfPrefixesVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.prefixNum.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.prefixNum.value", childItem.MaxNumberOfPrefixesVariable.ValueString())
						}
					} else if !childItem.MaxNumberOfPrefixes.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.prefixNum.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.prefixNum.value", childItem.MaxNumberOfPrefixes.ValueInt64())
						}
					}

					if !childItem.ThresholdVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.value", childItem.ThresholdVariable.ValueString())
						}
					} else if childItem.Threshold.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.optionType", "default")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.value", 75)
						}
					} else {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.threshold.value", childItem.Threshold.ValueInt64())
						}
					}
					if !childItem.PolicyType.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.policyType.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.policyType.value", childItem.PolicyType.ValueString())
						}
					}

					if !childItem.RestartIntervalVariable.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.restartInterval.optionType", "variable")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.restartInterval.value", childItem.RestartIntervalVariable.ValueString())
						}
					} else if !childItem.RestartInterval.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.restartInterval.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "maxPrefixConfig.restartInterval.value", childItem.RestartInterval.ValueInt64())
						}
					}
					if !childItem.InRoutePolicyId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "inRoutePolicy.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "inRoutePolicy.refId.value", childItem.InRoutePolicyId.ValueString())
						}
					}
					if !childItem.OutRoutePolicyId.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "outRoutePolicy.refId.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "outRoutePolicy.refId.value", childItem.OutRoutePolicyId.ValueString())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "addressFamily.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"ipv6Neighbor.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"addressFamily.aggregateAddress", []interface{}{})
		for _, item := range data.Ipv4AggregateAddresses {
			itemBody := ""

			if !item.NetworkAddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.address.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.address.value", item.NetworkAddressVariable.ValueString())
				}
			} else if !item.NetworkAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.address.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.address.value", item.NetworkAddress.ValueString())
				}
			}

			if !item.SubnetMaskVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.mask.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.mask.value", item.SubnetMaskVariable.ValueString())
				}
			} else if !item.SubnetMask.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.mask.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.mask.value", item.SubnetMask.ValueString())
				}
			}

			if !item.AsSetPathVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asSet.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "asSet.value", item.AsSetPathVariable.ValueString())
				}
			} else if item.AsSetPath.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asSet.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "asSet.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asSet.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "asSet.value", item.AsSetPath.ValueBool())
				}
			}

			if !item.SummaryOnlyVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "summaryOnly.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "summaryOnly.value", item.SummaryOnlyVariable.ValueString())
				}
			} else if item.SummaryOnly.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "summaryOnly.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "summaryOnly.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "summaryOnly.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "summaryOnly.value", item.SummaryOnly.ValueBool())
				}
			}
			body, _ = sjson.SetRaw(body, path+"addressFamily.aggregateAddress.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"addressFamily.network", []interface{}{})
		for _, item := range data.Ipv4Networks {
			itemBody := ""

			if !item.NetworkAddressVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.address.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.address.value", item.NetworkAddressVariable.ValueString())
				}
			} else if !item.NetworkAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.address.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.address.value", item.NetworkAddress.ValueString())
				}
			}

			if !item.SubnetMaskVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.mask.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.mask.value", item.SubnetMaskVariable.ValueString())
				}
			} else if !item.SubnetMask.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.mask.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.mask.value", item.SubnetMask.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"addressFamily.network.-1", itemBody)
		}
	}

	if !data.Ipv4EibgpMaximumPathsVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"addressFamily.paths.optionType", "variable")
			body, _ = sjson.Set(body, path+"addressFamily.paths.value", data.Ipv4EibgpMaximumPathsVariable.ValueString())
		}
	} else if data.Ipv4EibgpMaximumPaths.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"addressFamily.paths.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"addressFamily.paths.optionType", "global")
			body, _ = sjson.Set(body, path+"addressFamily.paths.value", data.Ipv4EibgpMaximumPaths.ValueInt64())
		}
	}

	if !data.Ipv4OriginateVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"addressFamily.originate.optionType", "variable")
			body, _ = sjson.Set(body, path+"addressFamily.originate.value", data.Ipv4OriginateVariable.ValueString())
		}
	} else if data.Ipv4Originate.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"addressFamily.originate.optionType", "default")
			body, _ = sjson.Set(body, path+"addressFamily.originate.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"addressFamily.originate.optionType", "global")
			body, _ = sjson.Set(body, path+"addressFamily.originate.value", data.Ipv4Originate.ValueBool())
		}
	}
	if data.Ipv4TableMapRoutePolicyId.IsNull() {
		body, _ = sjson.Set(body, path+"addressFamily.name.optionType", "default")
	} else {
		if true {
			body, _ = sjson.Set(body, path+"addressFamily.name.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"addressFamily.name.refId.value", data.Ipv4TableMapRoutePolicyId.ValueString())
		}
	}

	if !data.Ipv4TableMapFilterVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"addressFamily.filter.optionType", "variable")
			body, _ = sjson.Set(body, path+"addressFamily.filter.value", data.Ipv4TableMapFilterVariable.ValueString())
		}
	} else if data.Ipv4TableMapFilter.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"addressFamily.filter.optionType", "default")
			body, _ = sjson.Set(body, path+"addressFamily.filter.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"addressFamily.filter.optionType", "global")
			body, _ = sjson.Set(body, path+"addressFamily.filter.value", data.Ipv4TableMapFilter.ValueBool())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"addressFamily.redistribute", []interface{}{})
		for _, item := range data.Ipv4Redistributes {
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
			if !item.RoutePolicyId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.value", item.RoutePolicyId.ValueString())
				}
			}
			if version.LessThan(MinServiceRoutingBGPUpdateVersion) {
			} else {
				if item.TranslateRibMetric.IsNull() {
					if true && item.Protocol.ValueString() == "omp" {
						itemBody, _ = sjson.Set(itemBody, "translateRibMetric.optionType", "default")
						itemBody, _ = sjson.Set(itemBody, "translateRibMetric.value", false)
					}
				} else {
					if true && item.Protocol.ValueString() == "omp" {
						itemBody, _ = sjson.Set(itemBody, "translateRibMetric.optionType", "global")
						itemBody, _ = sjson.Set(itemBody, "translateRibMetric.value", item.TranslateRibMetric.ValueBool())
					}
				}
			}
			body, _ = sjson.SetRaw(body, path+"addressFamily.redistribute.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"ipv6AddressFamily.ipv6AggregateAddress", []interface{}{})
		for _, item := range data.Ipv6AggregateAddresses {
			itemBody := ""

			if !item.AggregatePrefixVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.value", item.AggregatePrefixVariable.ValueString())
				}
			} else if !item.AggregatePrefix.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.value", item.AggregatePrefix.ValueString())
				}
			}

			if !item.AsSetPathVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asSet.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "asSet.value", item.AsSetPathVariable.ValueString())
				}
			} else if item.AsSetPath.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asSet.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "asSet.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "asSet.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "asSet.value", item.AsSetPath.ValueBool())
				}
			}

			if !item.SummaryOnlyVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "summaryOnly.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "summaryOnly.value", item.SummaryOnlyVariable.ValueString())
				}
			} else if item.SummaryOnly.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "summaryOnly.optionType", "default")
					itemBody, _ = sjson.Set(itemBody, "summaryOnly.value", false)
				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "summaryOnly.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "summaryOnly.value", item.SummaryOnly.ValueBool())
				}
			}
			body, _ = sjson.SetRaw(body, path+"ipv6AddressFamily.ipv6AggregateAddress.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"ipv6AddressFamily.ipv6Network", []interface{}{})
		for _, item := range data.Ipv6Networks {
			itemBody := ""

			if !item.NetworkPrefixVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "prefix.value", item.NetworkPrefixVariable.ValueString())
				}
			} else if !item.NetworkPrefix.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "prefix.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "prefix.value", item.NetworkPrefix.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"ipv6AddressFamily.ipv6Network.-1", itemBody)
		}
	}

	if !data.Ipv6EibgpMaximumPathsVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.paths.optionType", "variable")
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.paths.value", data.Ipv6EibgpMaximumPathsVariable.ValueString())
		}
	} else if data.Ipv6EibgpMaximumPaths.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.paths.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.paths.optionType", "global")
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.paths.value", data.Ipv6EibgpMaximumPaths.ValueInt64())
		}
	}

	if !data.Ipv6OriginateVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.originate.optionType", "variable")
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.originate.value", data.Ipv6OriginateVariable.ValueString())
		}
	} else if data.Ipv6Originate.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.originate.optionType", "default")
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.originate.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.originate.optionType", "global")
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.originate.value", data.Ipv6Originate.ValueBool())
		}
	}
	if data.Ipv6TableMapRoutePolicyId.IsNull() {
		body, _ = sjson.Set(body, path+"ipv6AddressFamily.name.optionType", "default")
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.name.refId.optionType", "global")
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.name.refId.value", data.Ipv6TableMapRoutePolicyId.ValueString())
		}
	}

	if !data.Ipv6TableMapFilterVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.filter.optionType", "variable")
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.filter.value", data.Ipv6TableMapFilterVariable.ValueString())
		}
	} else if data.Ipv6TableMapFilter.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.filter.optionType", "default")
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.filter.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.filter.optionType", "global")
			body, _ = sjson.Set(body, path+"ipv6AddressFamily.filter.value", data.Ipv6TableMapFilter.ValueBool())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"ipv6AddressFamily.redistribute", []interface{}{})
		for _, item := range data.Ipv6Redistributes {
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
			if !item.RoutePolicyId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "routePolicy.refId.value", item.RoutePolicyId.ValueString())
				}
			}
			if version.LessThan(MinServiceRoutingBGPUpdateVersion) {
			} else {
				if item.TranslateRibMetric.IsNull() {
					if true && item.Protocol.ValueString() == "omp" {
						itemBody, _ = sjson.Set(itemBody, "translateRibMetric.optionType", "default")
						itemBody, _ = sjson.Set(itemBody, "translateRibMetric.value", false)
					}
				} else {
					if true && item.Protocol.ValueString() == "omp" {
						itemBody, _ = sjson.Set(itemBody, "translateRibMetric.optionType", "global")
						itemBody, _ = sjson.Set(itemBody, "translateRibMetric.value", item.TranslateRibMetric.ValueBool())
					}
				}
			}
			body, _ = sjson.SetRaw(body, path+"ipv6AddressFamily.redistribute.-1", itemBody)
		}
	}
	return body
}

func (data *ServiceRoutingBGP) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.AsNumber = types.Int64Null()
	data.AsNumberVariable = types.StringNull()
	if t := res.Get(path + "asNum.optionType"); t.Exists() {
		va := res.Get(path + "asNum.value")
		if t.String() == "variable" {
			data.AsNumberVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AsNumber = types.Int64Value(va.Int())
		}
	}
	data.RouterId = types.StringNull()
	data.RouterIdVariable = types.StringNull()
	if t := res.Get(path + "routerId.optionType"); t.Exists() {
		va := res.Get(path + "routerId.value")
		if t.String() == "variable" {
			data.RouterIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RouterId = types.StringValue(va.String())
		}
	}
	data.PropagateAsPath = types.BoolNull()
	data.PropagateAsPathVariable = types.StringNull()
	if t := res.Get(path + "propagateAspath.optionType"); t.Exists() {
		va := res.Get(path + "propagateAspath.value")
		if t.String() == "variable" {
			data.PropagateAsPathVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PropagateAsPath = types.BoolValue(va.Bool())
		}
	}
	data.PropagateCommunity = types.BoolNull()
	data.PropagateCommunityVariable = types.StringNull()
	if t := res.Get(path + "propagateCommunity.optionType"); t.Exists() {
		va := res.Get(path + "propagateCommunity.value")
		if t.String() == "variable" {
			data.PropagateCommunityVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PropagateCommunity = types.BoolValue(va.Bool())
		}
	}
	data.ExternalRoutesDistance = types.Int64Null()
	data.ExternalRoutesDistanceVariable = types.StringNull()
	if t := res.Get(path + "external.optionType"); t.Exists() {
		va := res.Get(path + "external.value")
		if t.String() == "variable" {
			data.ExternalRoutesDistanceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ExternalRoutesDistance = types.Int64Value(va.Int())
		}
	}
	data.InternalRoutesDistance = types.Int64Null()
	data.InternalRoutesDistanceVariable = types.StringNull()
	if t := res.Get(path + "internal.optionType"); t.Exists() {
		va := res.Get(path + "internal.value")
		if t.String() == "variable" {
			data.InternalRoutesDistanceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InternalRoutesDistance = types.Int64Value(va.Int())
		}
	}
	data.LocalRoutesDistance = types.Int64Null()
	data.LocalRoutesDistanceVariable = types.StringNull()
	if t := res.Get(path + "local.optionType"); t.Exists() {
		va := res.Get(path + "local.value")
		if t.String() == "variable" {
			data.LocalRoutesDistanceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.LocalRoutesDistance = types.Int64Value(va.Int())
		}
	}
	data.KeepaliveTime = types.Int64Null()
	data.KeepaliveTimeVariable = types.StringNull()
	if t := res.Get(path + "keepalive.optionType"); t.Exists() {
		va := res.Get(path + "keepalive.value")
		if t.String() == "variable" {
			data.KeepaliveTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.KeepaliveTime = types.Int64Value(va.Int())
		}
	}
	data.HoldTime = types.Int64Null()
	data.HoldTimeVariable = types.StringNull()
	if t := res.Get(path + "holdtime.optionType"); t.Exists() {
		va := res.Get(path + "holdtime.value")
		if t.String() == "variable" {
			data.HoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.HoldTime = types.Int64Value(va.Int())
		}
	}
	data.AlwaysCompareMed = types.BoolNull()
	data.AlwaysCompareMedVariable = types.StringNull()
	if t := res.Get(path + "alwaysCompare.optionType"); t.Exists() {
		va := res.Get(path + "alwaysCompare.value")
		if t.String() == "variable" {
			data.AlwaysCompareMedVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AlwaysCompareMed = types.BoolValue(va.Bool())
		}
	}
	data.DeterministicMed = types.BoolNull()
	data.DeterministicMedVariable = types.StringNull()
	if t := res.Get(path + "deterministic.optionType"); t.Exists() {
		va := res.Get(path + "deterministic.value")
		if t.String() == "variable" {
			data.DeterministicMedVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DeterministicMed = types.BoolValue(va.Bool())
		}
	}
	data.MissingMedAsWorst = types.BoolNull()
	data.MissingMedAsWorstVariable = types.StringNull()
	if t := res.Get(path + "missingAsWorst.optionType"); t.Exists() {
		va := res.Get(path + "missingAsWorst.value")
		if t.String() == "variable" {
			data.MissingMedAsWorstVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MissingMedAsWorst = types.BoolValue(va.Bool())
		}
	}
	data.CompareRouterId = types.BoolNull()
	data.CompareRouterIdVariable = types.StringNull()
	if t := res.Get(path + "compareRouterId.optionType"); t.Exists() {
		va := res.Get(path + "compareRouterId.value")
		if t.String() == "variable" {
			data.CompareRouterIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.CompareRouterId = types.BoolValue(va.Bool())
		}
	}
	data.MultipathRelax = types.BoolNull()
	data.MultipathRelaxVariable = types.StringNull()
	if t := res.Get(path + "multipathRelax.optionType"); t.Exists() {
		va := res.Get(path + "multipathRelax.value")
		if t.String() == "variable" {
			data.MultipathRelaxVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MultipathRelax = types.BoolValue(va.Bool())
		}
	}
	if value := res.Get(path + "neighbor"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv4Neighbors = make([]ServiceRoutingBGPIpv4Neighbors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingBGPIpv4Neighbors{}
			item.Address = types.StringNull()
			item.AddressVariable = types.StringNull()
			if t := v.Get("address.optionType"); t.Exists() {
				va := v.Get("address.value")
				if t.String() == "variable" {
					item.AddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Address = types.StringValue(va.String())
				}
			}
			item.Description = types.StringNull()
			item.DescriptionVariable = types.StringNull()
			if t := v.Get("description.optionType"); t.Exists() {
				va := v.Get("description.value")
				if t.String() == "variable" {
					item.DescriptionVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Description = types.StringValue(va.String())
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
			item.RemoteAs = types.Int64Null()
			item.RemoteAsVariable = types.StringNull()
			if t := v.Get("remoteAs.optionType"); t.Exists() {
				va := v.Get("remoteAs.value")
				if t.String() == "variable" {
					item.RemoteAsVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RemoteAs = types.Int64Value(va.Int())
				}
			}
			item.LocalAs = types.Int64Null()
			item.LocalAsVariable = types.StringNull()
			if t := v.Get("localAs.optionType"); t.Exists() {
				va := v.Get("localAs.value")
				if t.String() == "variable" {
					item.LocalAsVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.LocalAs = types.Int64Value(va.Int())
				}
			}
			item.KeepaliveTime = types.Int64Null()
			item.KeepaliveTimeVariable = types.StringNull()
			if t := v.Get("keepalive.optionType"); t.Exists() {
				va := v.Get("keepalive.value")
				if t.String() == "variable" {
					item.KeepaliveTimeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.KeepaliveTime = types.Int64Value(va.Int())
				}
			}
			item.HoldTime = types.Int64Null()
			item.HoldTimeVariable = types.StringNull()
			if t := v.Get("holdtime.optionType"); t.Exists() {
				va := v.Get("holdtime.value")
				if t.String() == "variable" {
					item.HoldTimeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.HoldTime = types.Int64Value(va.Int())
				}
			}
			item.UpdateSourceInterface = types.StringNull()
			item.UpdateSourceInterfaceVariable = types.StringNull()
			if t := v.Get("ifName.optionType"); t.Exists() {
				va := v.Get("ifName.value")
				if t.String() == "variable" {
					item.UpdateSourceInterfaceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.UpdateSourceInterface = types.StringValue(va.String())
				}
			}
			item.NextHopSelf = types.BoolNull()
			item.NextHopSelfVariable = types.StringNull()
			if t := v.Get("nextHopSelf.optionType"); t.Exists() {
				va := v.Get("nextHopSelf.value")
				if t.String() == "variable" {
					item.NextHopSelfVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NextHopSelf = types.BoolValue(va.Bool())
				}
			}
			item.SendCommunity = types.BoolNull()
			item.SendCommunityVariable = types.StringNull()
			if t := v.Get("sendCommunity.optionType"); t.Exists() {
				va := v.Get("sendCommunity.value")
				if t.String() == "variable" {
					item.SendCommunityVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SendCommunity = types.BoolValue(va.Bool())
				}
			}
			item.SendExtendedCommunity = types.BoolNull()
			item.SendExtendedCommunityVariable = types.StringNull()
			if t := v.Get("sendExtCommunity.optionType"); t.Exists() {
				va := v.Get("sendExtCommunity.value")
				if t.String() == "variable" {
					item.SendExtendedCommunityVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SendExtendedCommunity = types.BoolValue(va.Bool())
				}
			}
			item.EbgpMultihop = types.Int64Null()
			item.EbgpMultihopVariable = types.StringNull()
			if t := v.Get("ebgpMultihop.optionType"); t.Exists() {
				va := v.Get("ebgpMultihop.value")
				if t.String() == "variable" {
					item.EbgpMultihopVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.EbgpMultihop = types.Int64Value(va.Int())
				}
			}
			item.SendLabel = types.BoolNull()
			item.SendLabelVariable = types.StringNull()
			if t := v.Get("sendLabel.optionType"); t.Exists() {
				va := v.Get("sendLabel.value")
				if t.String() == "variable" {
					item.SendLabelVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SendLabel = types.BoolValue(va.Bool())
				}
			}
			item.AsOverride = types.BoolNull()
			item.AsOverrideVariable = types.StringNull()
			if t := v.Get("asOverride.optionType"); t.Exists() {
				va := v.Get("asOverride.value")
				if t.String() == "variable" {
					item.AsOverrideVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AsOverride = types.BoolValue(va.Bool())
				}
			}
			item.AllowasInNumber = types.Int64Null()
			item.AllowasInNumberVariable = types.StringNull()
			if t := v.Get("asNumber.optionType"); t.Exists() {
				va := v.Get("asNumber.value")
				if t.String() == "variable" {
					item.AllowasInNumberVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AllowasInNumber = types.Int64Value(va.Int())
				}
			}
			if cValue := v.Get("addressFamily"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.AddressFamilies = make([]ServiceRoutingBGPIpv4NeighborsAddressFamilies, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceRoutingBGPIpv4NeighborsAddressFamilies{}
					cItem.FamilyType = types.StringNull()

					if t := cv.Get("familyType.optionType"); t.Exists() {
						va := cv.Get("familyType.value")
						if t.String() == "global" {
							cItem.FamilyType = types.StringValue(va.String())
						}
					}
					cItem.PolicyType = types.StringNull()

					if t := cv.Get("maxPrefixConfig.policyType.optionType"); t.Exists() {
						va := cv.Get("maxPrefixConfig.policyType.value")
						if t.String() == "global" {
							cItem.PolicyType = types.StringValue(va.String())
						}
					}
					cItem.RestartMaxNumberOfPrefixes = types.Int64Null()
					cItem.RestartMaxNumberOfPrefixesVariable = types.StringNull()
					if true && cItem.PolicyType.ValueString() == "restart" {
						if t := cv.Get("maxPrefixConfig.prefixNum.optionType"); t.Exists() {
							va := cv.Get("maxPrefixConfig.prefixNum.value")
							if t.String() == "variable" {
								cItem.RestartMaxNumberOfPrefixesVariable = types.StringValue(va.String())
							} else if t.String() == "global" {
								cItem.RestartMaxNumberOfPrefixes = types.Int64Value(va.Int())
							}
							cItem.PolicyType = types.StringValue("restart")
						}
					}
					cItem.RestartThreshold = types.Int64Null()
					cItem.RestartThresholdVariable = types.StringNull()
					if true && cItem.PolicyType.ValueString() == "restart" {
						if t := cv.Get("maxPrefixConfig.threshold.optionType"); t.Exists() {
							va := cv.Get("maxPrefixConfig.threshold.value")
							if t.String() == "variable" {
								cItem.RestartThresholdVariable = types.StringValue(va.String())
							} else if t.String() == "global" {
								cItem.RestartThreshold = types.Int64Value(va.Int())
							}
							cItem.PolicyType = types.StringValue("restart")
						}
					}
					cItem.RestartInterval = types.Int64Null()
					cItem.RestartIntervalVariable = types.StringNull()
					if true && cItem.PolicyType.ValueString() == "restart" {
						if t := cv.Get("maxPrefixConfig.restartInterval.optionType"); t.Exists() {
							va := cv.Get("maxPrefixConfig.restartInterval.value")
							if t.String() == "variable" {
								cItem.RestartIntervalVariable = types.StringValue(va.String())
							} else if t.String() == "global" {
								cItem.RestartInterval = types.Int64Value(va.Int())
							}
							cItem.PolicyType = types.StringValue("restart")
						}
					}
					cItem.WarningMessageMaxNumberOfPrefixes = types.Int64Null()
					cItem.WarningMessageMaxNumberOfPrefixesVariable = types.StringNull()
					if true && cItem.PolicyType.ValueString() == "warning-only" {
						if t := cv.Get("maxPrefixConfig.prefixNum.optionType"); t.Exists() {
							va := cv.Get("maxPrefixConfig.prefixNum.value")
							if t.String() == "variable" {
								cItem.WarningMessageMaxNumberOfPrefixesVariable = types.StringValue(va.String())
							} else if t.String() == "global" {
								cItem.WarningMessageMaxNumberOfPrefixes = types.Int64Value(va.Int())
							}
							cItem.PolicyType = types.StringValue("warning-only")
						}
					}
					cItem.WarningMessageThreshold = types.Int64Null()
					cItem.WarningMessageThresholdVariable = types.StringNull()
					if true && cItem.PolicyType.ValueString() == "warning-only" {
						if t := cv.Get("maxPrefixConfig.threshold.optionType"); t.Exists() {
							va := cv.Get("maxPrefixConfig.threshold.value")
							if t.String() == "variable" {
								cItem.WarningMessageThresholdVariable = types.StringValue(va.String())
							} else if t.String() == "global" {
								cItem.WarningMessageThreshold = types.Int64Value(va.Int())
							}
							cItem.PolicyType = types.StringValue("warning-only")
						}
					}
					cItem.DisablePeerMaxNumberOfPrefixes = types.Int64Null()
					cItem.DisablePeerMaxNumberOfPrefixesVariable = types.StringNull()
					if true && cItem.PolicyType.ValueString() == "disable-peer" {
						if t := cv.Get("maxPrefixConfig.prefixNum.optionType"); t.Exists() {
							va := cv.Get("maxPrefixConfig.prefixNum.value")
							if t.String() == "variable" {
								cItem.DisablePeerMaxNumberOfPrefixesVariable = types.StringValue(va.String())
							} else if t.String() == "global" {
								cItem.DisablePeerMaxNumberOfPrefixes = types.Int64Value(va.Int())
							}
							cItem.PolicyType = types.StringValue("disable-peer")
						}
					}
					cItem.DisablePeerThreshold = types.Int64Null()
					cItem.DisablePeerThresholdVariable = types.StringNull()
					if true && cItem.PolicyType.ValueString() == "disable-peer" {
						if t := cv.Get("maxPrefixConfig.threshold.optionType"); t.Exists() {
							va := cv.Get("maxPrefixConfig.threshold.value")
							if t.String() == "variable" {
								cItem.DisablePeerThresholdVariable = types.StringValue(va.String())
							} else if t.String() == "global" {
								cItem.DisablePeerThreshold = types.Int64Value(va.Int())
							}
							cItem.PolicyType = types.StringValue("disable-peer")
						}
					}
					cItem.InRoutePolicyId = types.StringNull()

					if t := cv.Get("inRoutePolicy.refId.optionType"); t.Exists() {
						va := cv.Get("inRoutePolicy.refId.value")
						if t.String() == "global" {
							cItem.InRoutePolicyId = types.StringValue(va.String())
						}
					}
					cItem.OutRoutePolicyId = types.StringNull()

					if t := cv.Get("outRoutePolicy.refId.optionType"); t.Exists() {
						va := cv.Get("outRoutePolicy.refId.value")
						if t.String() == "global" {
							cItem.OutRoutePolicyId = types.StringValue(va.String())
						}
					}
					item.AddressFamilies = append(item.AddressFamilies, cItem)
					return true
				})
			}
			data.Ipv4Neighbors = append(data.Ipv4Neighbors, item)
			return true
		})
	}
	if value := res.Get(path + "ipv6Neighbor"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv6Neighbors = make([]ServiceRoutingBGPIpv6Neighbors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingBGPIpv6Neighbors{}
			item.Address = types.StringNull()
			item.AddressVariable = types.StringNull()
			if t := v.Get("address.optionType"); t.Exists() {
				va := v.Get("address.value")
				if t.String() == "variable" {
					item.AddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Address = types.StringValue(va.String())
				}
			}
			item.Description = types.StringNull()
			item.DescriptionVariable = types.StringNull()
			if t := v.Get("description.optionType"); t.Exists() {
				va := v.Get("description.value")
				if t.String() == "variable" {
					item.DescriptionVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Description = types.StringValue(va.String())
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
			item.RemoteAs = types.Int64Null()
			item.RemoteAsVariable = types.StringNull()
			if t := v.Get("remoteAs.optionType"); t.Exists() {
				va := v.Get("remoteAs.value")
				if t.String() == "variable" {
					item.RemoteAsVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.RemoteAs = types.Int64Value(va.Int())
				}
			}
			item.LocalAs = types.Int64Null()
			item.LocalAsVariable = types.StringNull()
			if t := v.Get("localAs.optionType"); t.Exists() {
				va := v.Get("localAs.value")
				if t.String() == "variable" {
					item.LocalAsVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.LocalAs = types.Int64Value(va.Int())
				}
			}
			item.KeepaliveTime = types.Int64Null()
			item.KeepaliveTimeVariable = types.StringNull()
			if t := v.Get("keepalive.optionType"); t.Exists() {
				va := v.Get("keepalive.value")
				if t.String() == "variable" {
					item.KeepaliveTimeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.KeepaliveTime = types.Int64Value(va.Int())
				}
			}
			item.HoldTime = types.Int64Null()
			item.HoldTimeVariable = types.StringNull()
			if t := v.Get("holdtime.optionType"); t.Exists() {
				va := v.Get("holdtime.value")
				if t.String() == "variable" {
					item.HoldTimeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.HoldTime = types.Int64Value(va.Int())
				}
			}
			item.UpdateSourceInterface = types.StringNull()
			item.UpdateSourceInterfaceVariable = types.StringNull()
			if t := v.Get("ifName.optionType"); t.Exists() {
				va := v.Get("ifName.value")
				if t.String() == "variable" {
					item.UpdateSourceInterfaceVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.UpdateSourceInterface = types.StringValue(va.String())
				}
			}
			item.NextHopSelf = types.BoolNull()
			item.NextHopSelfVariable = types.StringNull()
			if t := v.Get("nextHopSelf.optionType"); t.Exists() {
				va := v.Get("nextHopSelf.value")
				if t.String() == "variable" {
					item.NextHopSelfVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NextHopSelf = types.BoolValue(va.Bool())
				}
			}
			item.SendCommunity = types.BoolNull()
			item.SendCommunityVariable = types.StringNull()
			if t := v.Get("sendCommunity.optionType"); t.Exists() {
				va := v.Get("sendCommunity.value")
				if t.String() == "variable" {
					item.SendCommunityVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SendCommunity = types.BoolValue(va.Bool())
				}
			}
			item.SendExtendedCommunity = types.BoolNull()
			item.SendExtendedCommunityVariable = types.StringNull()
			if t := v.Get("sendExtCommunity.optionType"); t.Exists() {
				va := v.Get("sendExtCommunity.value")
				if t.String() == "variable" {
					item.SendExtendedCommunityVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SendExtendedCommunity = types.BoolValue(va.Bool())
				}
			}
			item.EbgpMultihop = types.Int64Null()
			item.EbgpMultihopVariable = types.StringNull()
			if t := v.Get("ebgpMultihop.optionType"); t.Exists() {
				va := v.Get("ebgpMultihop.value")
				if t.String() == "variable" {
					item.EbgpMultihopVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.EbgpMultihop = types.Int64Value(va.Int())
				}
			}
			item.AsOverride = types.BoolNull()
			item.AsOverrideVariable = types.StringNull()
			if t := v.Get("asOverride.optionType"); t.Exists() {
				va := v.Get("asOverride.value")
				if t.String() == "variable" {
					item.AsOverrideVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AsOverride = types.BoolValue(va.Bool())
				}
			}
			item.AllowasInNumber = types.Int64Null()
			item.AllowasInNumberVariable = types.StringNull()
			if t := v.Get("asNumber.optionType"); t.Exists() {
				va := v.Get("asNumber.value")
				if t.String() == "variable" {
					item.AllowasInNumberVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AllowasInNumber = types.Int64Value(va.Int())
				}
			}
			if cValue := v.Get("addressFamily"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.AddressFamilies = make([]ServiceRoutingBGPIpv6NeighborsAddressFamilies, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ServiceRoutingBGPIpv6NeighborsAddressFamilies{}
					cItem.FamilyType = types.StringNull()

					if t := cv.Get("familyType.optionType"); t.Exists() {
						va := cv.Get("familyType.value")
						if t.String() == "global" {
							cItem.FamilyType = types.StringValue(va.String())
						}
					}
					cItem.MaxNumberOfPrefixes = types.Int64Null()
					cItem.MaxNumberOfPrefixesVariable = types.StringNull()
					if t := cv.Get("maxPrefixConfig.prefixNum.optionType"); t.Exists() {
						va := cv.Get("maxPrefixConfig.prefixNum.value")
						if t.String() == "variable" {
							cItem.MaxNumberOfPrefixesVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.MaxNumberOfPrefixes = types.Int64Value(va.Int())
						}
					}
					cItem.Threshold = types.Int64Null()
					cItem.ThresholdVariable = types.StringNull()
					if t := cv.Get("maxPrefixConfig.threshold.optionType"); t.Exists() {
						va := cv.Get("maxPrefixConfig.threshold.value")
						if t.String() == "variable" {
							cItem.ThresholdVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.Threshold = types.Int64Value(va.Int())
						}
					}
					cItem.PolicyType = types.StringNull()

					if t := cv.Get("maxPrefixConfig.policyType.optionType"); t.Exists() {
						va := cv.Get("maxPrefixConfig.policyType.value")
						if t.String() == "global" {
							cItem.PolicyType = types.StringValue(va.String())
						}
					}
					cItem.RestartInterval = types.Int64Null()
					cItem.RestartIntervalVariable = types.StringNull()
					if t := cv.Get("maxPrefixConfig.restartInterval.optionType"); t.Exists() {
						va := cv.Get("maxPrefixConfig.restartInterval.value")
						if t.String() == "variable" {
							cItem.RestartIntervalVariable = types.StringValue(va.String())
						} else if t.String() == "global" {
							cItem.RestartInterval = types.Int64Value(va.Int())
						}
					}
					cItem.InRoutePolicyId = types.StringNull()

					if t := cv.Get("inRoutePolicy.refId.optionType"); t.Exists() {
						va := cv.Get("inRoutePolicy.refId.value")
						if t.String() == "global" {
							cItem.InRoutePolicyId = types.StringValue(va.String())
						}
					}
					cItem.OutRoutePolicyId = types.StringNull()

					if t := cv.Get("outRoutePolicy.refId.optionType"); t.Exists() {
						va := cv.Get("outRoutePolicy.refId.value")
						if t.String() == "global" {
							cItem.OutRoutePolicyId = types.StringValue(va.String())
						}
					}
					item.AddressFamilies = append(item.AddressFamilies, cItem)
					return true
				})
			}
			data.Ipv6Neighbors = append(data.Ipv6Neighbors, item)
			return true
		})
	}
	if value := res.Get(path + "addressFamily.aggregateAddress"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv4AggregateAddresses = make([]ServiceRoutingBGPIpv4AggregateAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingBGPIpv4AggregateAddresses{}
			item.NetworkAddress = types.StringNull()
			item.NetworkAddressVariable = types.StringNull()
			if t := v.Get("prefix.address.optionType"); t.Exists() {
				va := v.Get("prefix.address.value")
				if t.String() == "variable" {
					item.NetworkAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NetworkAddress = types.StringValue(va.String())
				}
			}
			item.SubnetMask = types.StringNull()
			item.SubnetMaskVariable = types.StringNull()
			if t := v.Get("prefix.mask.optionType"); t.Exists() {
				va := v.Get("prefix.mask.value")
				if t.String() == "variable" {
					item.SubnetMaskVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SubnetMask = types.StringValue(va.String())
				}
			}
			item.AsSetPath = types.BoolNull()
			item.AsSetPathVariable = types.StringNull()
			if t := v.Get("asSet.optionType"); t.Exists() {
				va := v.Get("asSet.value")
				if t.String() == "variable" {
					item.AsSetPathVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AsSetPath = types.BoolValue(va.Bool())
				}
			}
			item.SummaryOnly = types.BoolNull()
			item.SummaryOnlyVariable = types.StringNull()
			if t := v.Get("summaryOnly.optionType"); t.Exists() {
				va := v.Get("summaryOnly.value")
				if t.String() == "variable" {
					item.SummaryOnlyVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SummaryOnly = types.BoolValue(va.Bool())
				}
			}
			data.Ipv4AggregateAddresses = append(data.Ipv4AggregateAddresses, item)
			return true
		})
	}
	if value := res.Get(path + "addressFamily.network"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv4Networks = make([]ServiceRoutingBGPIpv4Networks, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingBGPIpv4Networks{}
			item.NetworkAddress = types.StringNull()
			item.NetworkAddressVariable = types.StringNull()
			if t := v.Get("prefix.address.optionType"); t.Exists() {
				va := v.Get("prefix.address.value")
				if t.String() == "variable" {
					item.NetworkAddressVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NetworkAddress = types.StringValue(va.String())
				}
			}
			item.SubnetMask = types.StringNull()
			item.SubnetMaskVariable = types.StringNull()
			if t := v.Get("prefix.mask.optionType"); t.Exists() {
				va := v.Get("prefix.mask.value")
				if t.String() == "variable" {
					item.SubnetMaskVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SubnetMask = types.StringValue(va.String())
				}
			}
			data.Ipv4Networks = append(data.Ipv4Networks, item)
			return true
		})
	}
	data.Ipv4EibgpMaximumPaths = types.Int64Null()
	data.Ipv4EibgpMaximumPathsVariable = types.StringNull()
	if t := res.Get(path + "addressFamily.paths.optionType"); t.Exists() {
		va := res.Get(path + "addressFamily.paths.value")
		if t.String() == "variable" {
			data.Ipv4EibgpMaximumPathsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4EibgpMaximumPaths = types.Int64Value(va.Int())
		}
	}
	data.Ipv4Originate = types.BoolNull()
	data.Ipv4OriginateVariable = types.StringNull()
	if t := res.Get(path + "addressFamily.originate.optionType"); t.Exists() {
		va := res.Get(path + "addressFamily.originate.value")
		if t.String() == "variable" {
			data.Ipv4OriginateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Originate = types.BoolValue(va.Bool())
		}
	}
	data.Ipv4TableMapRoutePolicyId = types.StringNull()

	if t := res.Get(path + "addressFamily.name.refId.optionType"); t.Exists() {
		va := res.Get(path + "addressFamily.name.refId.value")
		if t.String() == "global" {
			data.Ipv4TableMapRoutePolicyId = types.StringValue(va.String())
		}
	}
	data.Ipv4TableMapFilter = types.BoolNull()
	data.Ipv4TableMapFilterVariable = types.StringNull()
	if t := res.Get(path + "addressFamily.filter.optionType"); t.Exists() {
		va := res.Get(path + "addressFamily.filter.value")
		if t.String() == "variable" {
			data.Ipv4TableMapFilterVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4TableMapFilter = types.BoolValue(va.Bool())
		}
	}
	if value := res.Get(path + "addressFamily.redistribute"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv4Redistributes = make([]ServiceRoutingBGPIpv4Redistributes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingBGPIpv4Redistributes{}
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
			item.TranslateRibMetric = types.BoolNull()

			if t := v.Get("translateRibMetric.optionType"); t.Exists() {
				va := v.Get("translateRibMetric.value")
				if t.String() == "global" {
					item.TranslateRibMetric = types.BoolValue(va.Bool())
				}
				item.Protocol = types.StringValue("omp")
			}
			data.Ipv4Redistributes = append(data.Ipv4Redistributes, item)
			return true
		})
	}
	if value := res.Get(path + "ipv6AddressFamily.ipv6AggregateAddress"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv6AggregateAddresses = make([]ServiceRoutingBGPIpv6AggregateAddresses, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingBGPIpv6AggregateAddresses{}
			item.AggregatePrefix = types.StringNull()
			item.AggregatePrefixVariable = types.StringNull()
			if t := v.Get("prefix.optionType"); t.Exists() {
				va := v.Get("prefix.value")
				if t.String() == "variable" {
					item.AggregatePrefixVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AggregatePrefix = types.StringValue(va.String())
				}
			}
			item.AsSetPath = types.BoolNull()
			item.AsSetPathVariable = types.StringNull()
			if t := v.Get("asSet.optionType"); t.Exists() {
				va := v.Get("asSet.value")
				if t.String() == "variable" {
					item.AsSetPathVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AsSetPath = types.BoolValue(va.Bool())
				}
			}
			item.SummaryOnly = types.BoolNull()
			item.SummaryOnlyVariable = types.StringNull()
			if t := v.Get("summaryOnly.optionType"); t.Exists() {
				va := v.Get("summaryOnly.value")
				if t.String() == "variable" {
					item.SummaryOnlyVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.SummaryOnly = types.BoolValue(va.Bool())
				}
			}
			data.Ipv6AggregateAddresses = append(data.Ipv6AggregateAddresses, item)
			return true
		})
	}
	if value := res.Get(path + "ipv6AddressFamily.ipv6Network"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv6Networks = make([]ServiceRoutingBGPIpv6Networks, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingBGPIpv6Networks{}
			item.NetworkPrefix = types.StringNull()
			item.NetworkPrefixVariable = types.StringNull()
			if t := v.Get("prefix.optionType"); t.Exists() {
				va := v.Get("prefix.value")
				if t.String() == "variable" {
					item.NetworkPrefixVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.NetworkPrefix = types.StringValue(va.String())
				}
			}
			data.Ipv6Networks = append(data.Ipv6Networks, item)
			return true
		})
	}
	data.Ipv6EibgpMaximumPaths = types.Int64Null()
	data.Ipv6EibgpMaximumPathsVariable = types.StringNull()
	if t := res.Get(path + "ipv6AddressFamily.paths.optionType"); t.Exists() {
		va := res.Get(path + "ipv6AddressFamily.paths.value")
		if t.String() == "variable" {
			data.Ipv6EibgpMaximumPathsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6EibgpMaximumPaths = types.Int64Value(va.Int())
		}
	}
	data.Ipv6Originate = types.BoolNull()
	data.Ipv6OriginateVariable = types.StringNull()
	if t := res.Get(path + "ipv6AddressFamily.originate.optionType"); t.Exists() {
		va := res.Get(path + "ipv6AddressFamily.originate.value")
		if t.String() == "variable" {
			data.Ipv6OriginateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6Originate = types.BoolValue(va.Bool())
		}
	}
	data.Ipv6TableMapRoutePolicyId = types.StringNull()

	if t := res.Get(path + "ipv6AddressFamily.name.refId.optionType"); t.Exists() {
		va := res.Get(path + "ipv6AddressFamily.name.refId.value")
		if t.String() == "global" {
			data.Ipv6TableMapRoutePolicyId = types.StringValue(va.String())
		}
	}
	data.Ipv6TableMapFilter = types.BoolNull()
	data.Ipv6TableMapFilterVariable = types.StringNull()
	if t := res.Get(path + "ipv6AddressFamily.filter.optionType"); t.Exists() {
		va := res.Get(path + "ipv6AddressFamily.filter.value")
		if t.String() == "variable" {
			data.Ipv6TableMapFilterVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6TableMapFilter = types.BoolValue(va.Bool())
		}
	}
	if value := res.Get(path + "ipv6AddressFamily.redistribute"); value.Exists() && len(value.Array()) > 0 {
		data.Ipv6Redistributes = make([]ServiceRoutingBGPIpv6Redistributes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ServiceRoutingBGPIpv6Redistributes{}
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
			item.TranslateRibMetric = types.BoolNull()

			if t := v.Get("translateRibMetric.optionType"); t.Exists() {
				va := v.Get("translateRibMetric.value")
				if t.String() == "global" {
					item.TranslateRibMetric = types.BoolValue(va.Bool())
				}
				item.Protocol = types.StringValue("omp")
			}
			data.Ipv6Redistributes = append(data.Ipv6Redistributes, item)
			return true
		})
	}
}

func (data *ServiceRoutingBGP) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.AsNumber = types.Int64Null()
	data.AsNumberVariable = types.StringNull()
	if t := res.Get(path + "asNum.optionType"); t.Exists() {
		va := res.Get(path + "asNum.value")
		if t.String() == "variable" {
			data.AsNumberVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AsNumber = types.Int64Value(va.Int())
		}
	}
	data.RouterId = types.StringNull()
	data.RouterIdVariable = types.StringNull()
	if t := res.Get(path + "routerId.optionType"); t.Exists() {
		va := res.Get(path + "routerId.value")
		if t.String() == "variable" {
			data.RouterIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.RouterId = types.StringValue(va.String())
		}
	}
	data.PropagateAsPath = types.BoolNull()
	data.PropagateAsPathVariable = types.StringNull()
	if t := res.Get(path + "propagateAspath.optionType"); t.Exists() {
		va := res.Get(path + "propagateAspath.value")
		if t.String() == "variable" {
			data.PropagateAsPathVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PropagateAsPath = types.BoolValue(va.Bool())
		}
	}
	data.PropagateCommunity = types.BoolNull()
	data.PropagateCommunityVariable = types.StringNull()
	if t := res.Get(path + "propagateCommunity.optionType"); t.Exists() {
		va := res.Get(path + "propagateCommunity.value")
		if t.String() == "variable" {
			data.PropagateCommunityVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PropagateCommunity = types.BoolValue(va.Bool())
		}
	}
	data.ExternalRoutesDistance = types.Int64Null()
	data.ExternalRoutesDistanceVariable = types.StringNull()
	if t := res.Get(path + "external.optionType"); t.Exists() {
		va := res.Get(path + "external.value")
		if t.String() == "variable" {
			data.ExternalRoutesDistanceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ExternalRoutesDistance = types.Int64Value(va.Int())
		}
	}
	data.InternalRoutesDistance = types.Int64Null()
	data.InternalRoutesDistanceVariable = types.StringNull()
	if t := res.Get(path + "internal.optionType"); t.Exists() {
		va := res.Get(path + "internal.value")
		if t.String() == "variable" {
			data.InternalRoutesDistanceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.InternalRoutesDistance = types.Int64Value(va.Int())
		}
	}
	data.LocalRoutesDistance = types.Int64Null()
	data.LocalRoutesDistanceVariable = types.StringNull()
	if t := res.Get(path + "local.optionType"); t.Exists() {
		va := res.Get(path + "local.value")
		if t.String() == "variable" {
			data.LocalRoutesDistanceVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.LocalRoutesDistance = types.Int64Value(va.Int())
		}
	}
	data.KeepaliveTime = types.Int64Null()
	data.KeepaliveTimeVariable = types.StringNull()
	if t := res.Get(path + "keepalive.optionType"); t.Exists() {
		va := res.Get(path + "keepalive.value")
		if t.String() == "variable" {
			data.KeepaliveTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.KeepaliveTime = types.Int64Value(va.Int())
		}
	}
	data.HoldTime = types.Int64Null()
	data.HoldTimeVariable = types.StringNull()
	if t := res.Get(path + "holdtime.optionType"); t.Exists() {
		va := res.Get(path + "holdtime.value")
		if t.String() == "variable" {
			data.HoldTimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.HoldTime = types.Int64Value(va.Int())
		}
	}
	data.AlwaysCompareMed = types.BoolNull()
	data.AlwaysCompareMedVariable = types.StringNull()
	if t := res.Get(path + "alwaysCompare.optionType"); t.Exists() {
		va := res.Get(path + "alwaysCompare.value")
		if t.String() == "variable" {
			data.AlwaysCompareMedVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AlwaysCompareMed = types.BoolValue(va.Bool())
		}
	}
	data.DeterministicMed = types.BoolNull()
	data.DeterministicMedVariable = types.StringNull()
	if t := res.Get(path + "deterministic.optionType"); t.Exists() {
		va := res.Get(path + "deterministic.value")
		if t.String() == "variable" {
			data.DeterministicMedVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DeterministicMed = types.BoolValue(va.Bool())
		}
	}
	data.MissingMedAsWorst = types.BoolNull()
	data.MissingMedAsWorstVariable = types.StringNull()
	if t := res.Get(path + "missingAsWorst.optionType"); t.Exists() {
		va := res.Get(path + "missingAsWorst.value")
		if t.String() == "variable" {
			data.MissingMedAsWorstVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MissingMedAsWorst = types.BoolValue(va.Bool())
		}
	}
	data.CompareRouterId = types.BoolNull()
	data.CompareRouterIdVariable = types.StringNull()
	if t := res.Get(path + "compareRouterId.optionType"); t.Exists() {
		va := res.Get(path + "compareRouterId.value")
		if t.String() == "variable" {
			data.CompareRouterIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.CompareRouterId = types.BoolValue(va.Bool())
		}
	}
	data.MultipathRelax = types.BoolNull()
	data.MultipathRelaxVariable = types.StringNull()
	if t := res.Get(path + "multipathRelax.optionType"); t.Exists() {
		va := res.Get(path + "multipathRelax.value")
		if t.String() == "variable" {
			data.MultipathRelaxVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MultipathRelax = types.BoolValue(va.Bool())
		}
	}
	for i := range data.Ipv4Neighbors {
		keys := [...]string{"address"}
		keyValues := [...]string{data.Ipv4Neighbors[i].Address.ValueString()}
		keyValuesVariables := [...]string{data.Ipv4Neighbors[i].AddressVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "neighbor").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
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
		data.Ipv4Neighbors[i].Address = types.StringNull()
		data.Ipv4Neighbors[i].AddressVariable = types.StringNull()
		if t := r.Get("address.optionType"); t.Exists() {
			va := r.Get("address.value")
			if t.String() == "variable" {
				data.Ipv4Neighbors[i].AddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Neighbors[i].Address = types.StringValue(va.String())
			}
		}
		data.Ipv4Neighbors[i].Description = types.StringNull()
		data.Ipv4Neighbors[i].DescriptionVariable = types.StringNull()
		if t := r.Get("description.optionType"); t.Exists() {
			va := r.Get("description.value")
			if t.String() == "variable" {
				data.Ipv4Neighbors[i].DescriptionVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Neighbors[i].Description = types.StringValue(va.String())
			}
		}
		data.Ipv4Neighbors[i].Shutdown = types.BoolNull()
		data.Ipv4Neighbors[i].ShutdownVariable = types.StringNull()
		if t := r.Get("shutdown.optionType"); t.Exists() {
			va := r.Get("shutdown.value")
			if t.String() == "variable" {
				data.Ipv4Neighbors[i].ShutdownVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Neighbors[i].Shutdown = types.BoolValue(va.Bool())
			}
		}
		data.Ipv4Neighbors[i].RemoteAs = types.Int64Null()
		data.Ipv4Neighbors[i].RemoteAsVariable = types.StringNull()
		if t := r.Get("remoteAs.optionType"); t.Exists() {
			va := r.Get("remoteAs.value")
			if t.String() == "variable" {
				data.Ipv4Neighbors[i].RemoteAsVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Neighbors[i].RemoteAs = types.Int64Value(va.Int())
			}
		}
		data.Ipv4Neighbors[i].LocalAs = types.Int64Null()
		data.Ipv4Neighbors[i].LocalAsVariable = types.StringNull()
		if t := r.Get("localAs.optionType"); t.Exists() {
			va := r.Get("localAs.value")
			if t.String() == "variable" {
				data.Ipv4Neighbors[i].LocalAsVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Neighbors[i].LocalAs = types.Int64Value(va.Int())
			}
		}
		data.Ipv4Neighbors[i].KeepaliveTime = types.Int64Null()
		data.Ipv4Neighbors[i].KeepaliveTimeVariable = types.StringNull()
		if t := r.Get("keepalive.optionType"); t.Exists() {
			va := r.Get("keepalive.value")
			if t.String() == "variable" {
				data.Ipv4Neighbors[i].KeepaliveTimeVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Neighbors[i].KeepaliveTime = types.Int64Value(va.Int())
			}
		}
		data.Ipv4Neighbors[i].HoldTime = types.Int64Null()
		data.Ipv4Neighbors[i].HoldTimeVariable = types.StringNull()
		if t := r.Get("holdtime.optionType"); t.Exists() {
			va := r.Get("holdtime.value")
			if t.String() == "variable" {
				data.Ipv4Neighbors[i].HoldTimeVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Neighbors[i].HoldTime = types.Int64Value(va.Int())
			}
		}
		data.Ipv4Neighbors[i].UpdateSourceInterface = types.StringNull()
		data.Ipv4Neighbors[i].UpdateSourceInterfaceVariable = types.StringNull()
		if t := r.Get("ifName.optionType"); t.Exists() {
			va := r.Get("ifName.value")
			if t.String() == "variable" {
				data.Ipv4Neighbors[i].UpdateSourceInterfaceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Neighbors[i].UpdateSourceInterface = types.StringValue(va.String())
			}
		}
		data.Ipv4Neighbors[i].NextHopSelf = types.BoolNull()
		data.Ipv4Neighbors[i].NextHopSelfVariable = types.StringNull()
		if t := r.Get("nextHopSelf.optionType"); t.Exists() {
			va := r.Get("nextHopSelf.value")
			if t.String() == "variable" {
				data.Ipv4Neighbors[i].NextHopSelfVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Neighbors[i].NextHopSelf = types.BoolValue(va.Bool())
			}
		}
		data.Ipv4Neighbors[i].SendCommunity = types.BoolNull()
		data.Ipv4Neighbors[i].SendCommunityVariable = types.StringNull()
		if t := r.Get("sendCommunity.optionType"); t.Exists() {
			va := r.Get("sendCommunity.value")
			if t.String() == "variable" {
				data.Ipv4Neighbors[i].SendCommunityVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Neighbors[i].SendCommunity = types.BoolValue(va.Bool())
			}
		}
		data.Ipv4Neighbors[i].SendExtendedCommunity = types.BoolNull()
		data.Ipv4Neighbors[i].SendExtendedCommunityVariable = types.StringNull()
		if t := r.Get("sendExtCommunity.optionType"); t.Exists() {
			va := r.Get("sendExtCommunity.value")
			if t.String() == "variable" {
				data.Ipv4Neighbors[i].SendExtendedCommunityVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Neighbors[i].SendExtendedCommunity = types.BoolValue(va.Bool())
			}
		}
		data.Ipv4Neighbors[i].EbgpMultihop = types.Int64Null()
		data.Ipv4Neighbors[i].EbgpMultihopVariable = types.StringNull()
		if t := r.Get("ebgpMultihop.optionType"); t.Exists() {
			va := r.Get("ebgpMultihop.value")
			if t.String() == "variable" {
				data.Ipv4Neighbors[i].EbgpMultihopVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Neighbors[i].EbgpMultihop = types.Int64Value(va.Int())
			}
		}
		data.Ipv4Neighbors[i].SendLabel = types.BoolNull()
		data.Ipv4Neighbors[i].SendLabelVariable = types.StringNull()
		if t := r.Get("sendLabel.optionType"); t.Exists() {
			va := r.Get("sendLabel.value")
			if t.String() == "variable" {
				data.Ipv4Neighbors[i].SendLabelVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Neighbors[i].SendLabel = types.BoolValue(va.Bool())
			}
		}
		data.Ipv4Neighbors[i].AsOverride = types.BoolNull()
		data.Ipv4Neighbors[i].AsOverrideVariable = types.StringNull()
		if t := r.Get("asOverride.optionType"); t.Exists() {
			va := r.Get("asOverride.value")
			if t.String() == "variable" {
				data.Ipv4Neighbors[i].AsOverrideVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Neighbors[i].AsOverride = types.BoolValue(va.Bool())
			}
		}
		data.Ipv4Neighbors[i].AllowasInNumber = types.Int64Null()
		data.Ipv4Neighbors[i].AllowasInNumberVariable = types.StringNull()
		if t := r.Get("asNumber.optionType"); t.Exists() {
			va := r.Get("asNumber.value")
			if t.String() == "variable" {
				data.Ipv4Neighbors[i].AllowasInNumberVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Neighbors[i].AllowasInNumber = types.Int64Value(va.Int())
			}
		}
		for ci := range data.Ipv4Neighbors[i].AddressFamilies {
			keys := [...]string{"familyType"}
			keyValues := [...]string{data.Ipv4Neighbors[i].AddressFamilies[ci].FamilyType.ValueString()}
			keyValuesVariables := [...]string{""}

			var cr gjson.Result
			r.Get("addressFamily").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType")
						vv := v.Get(keys[ik] + ".value")
						if tt.Exists() && vv.Exists() {
							if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
								found = true
								continue
							} else if tt.String() == "default" {
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
			data.Ipv4Neighbors[i].AddressFamilies[ci].FamilyType = types.StringNull()

			if t := cr.Get("familyType.optionType"); t.Exists() {
				va := cr.Get("familyType.value")
				if t.String() == "global" {
					data.Ipv4Neighbors[i].AddressFamilies[ci].FamilyType = types.StringValue(va.String())
				}
			}
			data.Ipv4Neighbors[i].AddressFamilies[ci].PolicyType = types.StringNull()

			if t := cr.Get("maxPrefixConfig.policyType.optionType"); t.Exists() {
				va := cr.Get("maxPrefixConfig.policyType.value")
				if t.String() == "global" {
					data.Ipv4Neighbors[i].AddressFamilies[ci].PolicyType = types.StringValue(va.String())
				}
			}
			data.Ipv4Neighbors[i].AddressFamilies[ci].RestartMaxNumberOfPrefixes = types.Int64Null()
			data.Ipv4Neighbors[i].AddressFamilies[ci].RestartMaxNumberOfPrefixesVariable = types.StringNull()
			if true && data.Ipv4Neighbors[i].AddressFamilies[ci].PolicyType.ValueString() == "restart" {
				if t := cr.Get("maxPrefixConfig.prefixNum.optionType"); t.Exists() {
					va := cr.Get("maxPrefixConfig.prefixNum.value")
					if t.String() == "variable" {
						data.Ipv4Neighbors[i].AddressFamilies[ci].RestartMaxNumberOfPrefixesVariable = types.StringValue(va.String())
					} else if t.String() == "global" {
						data.Ipv4Neighbors[i].AddressFamilies[ci].RestartMaxNumberOfPrefixes = types.Int64Value(va.Int())
					}
				}
			}
			data.Ipv4Neighbors[i].AddressFamilies[ci].RestartThreshold = types.Int64Null()
			data.Ipv4Neighbors[i].AddressFamilies[ci].RestartThresholdVariable = types.StringNull()
			if true && data.Ipv4Neighbors[i].AddressFamilies[ci].PolicyType.ValueString() == "restart" {
				if t := cr.Get("maxPrefixConfig.threshold.optionType"); t.Exists() {
					va := cr.Get("maxPrefixConfig.threshold.value")
					if t.String() == "variable" {
						data.Ipv4Neighbors[i].AddressFamilies[ci].RestartThresholdVariable = types.StringValue(va.String())
					} else if t.String() == "global" {
						data.Ipv4Neighbors[i].AddressFamilies[ci].RestartThreshold = types.Int64Value(va.Int())
					}
				}
			}
			data.Ipv4Neighbors[i].AddressFamilies[ci].RestartInterval = types.Int64Null()
			data.Ipv4Neighbors[i].AddressFamilies[ci].RestartIntervalVariable = types.StringNull()
			if true && data.Ipv4Neighbors[i].AddressFamilies[ci].PolicyType.ValueString() == "restart" {
				if t := cr.Get("maxPrefixConfig.restartInterval.optionType"); t.Exists() {
					va := cr.Get("maxPrefixConfig.restartInterval.value")
					if t.String() == "variable" {
						data.Ipv4Neighbors[i].AddressFamilies[ci].RestartIntervalVariable = types.StringValue(va.String())
					} else if t.String() == "global" {
						data.Ipv4Neighbors[i].AddressFamilies[ci].RestartInterval = types.Int64Value(va.Int())
					}
				}
			}
			data.Ipv4Neighbors[i].AddressFamilies[ci].WarningMessageMaxNumberOfPrefixes = types.Int64Null()
			data.Ipv4Neighbors[i].AddressFamilies[ci].WarningMessageMaxNumberOfPrefixesVariable = types.StringNull()
			if true && data.Ipv4Neighbors[i].AddressFamilies[ci].PolicyType.ValueString() == "warning-only" {
				if t := cr.Get("maxPrefixConfig.prefixNum.optionType"); t.Exists() {
					va := cr.Get("maxPrefixConfig.prefixNum.value")
					if t.String() == "variable" {
						data.Ipv4Neighbors[i].AddressFamilies[ci].WarningMessageMaxNumberOfPrefixesVariable = types.StringValue(va.String())
					} else if t.String() == "global" {
						data.Ipv4Neighbors[i].AddressFamilies[ci].WarningMessageMaxNumberOfPrefixes = types.Int64Value(va.Int())
					}
				}
			}
			data.Ipv4Neighbors[i].AddressFamilies[ci].WarningMessageThreshold = types.Int64Null()
			data.Ipv4Neighbors[i].AddressFamilies[ci].WarningMessageThresholdVariable = types.StringNull()
			if true && data.Ipv4Neighbors[i].AddressFamilies[ci].PolicyType.ValueString() == "warning-only" {
				if t := cr.Get("maxPrefixConfig.threshold.optionType"); t.Exists() {
					va := cr.Get("maxPrefixConfig.threshold.value")
					if t.String() == "variable" {
						data.Ipv4Neighbors[i].AddressFamilies[ci].WarningMessageThresholdVariable = types.StringValue(va.String())
					} else if t.String() == "global" {
						data.Ipv4Neighbors[i].AddressFamilies[ci].WarningMessageThreshold = types.Int64Value(va.Int())
					}
				}
			}
			data.Ipv4Neighbors[i].AddressFamilies[ci].DisablePeerMaxNumberOfPrefixes = types.Int64Null()
			data.Ipv4Neighbors[i].AddressFamilies[ci].DisablePeerMaxNumberOfPrefixesVariable = types.StringNull()
			if true && data.Ipv4Neighbors[i].AddressFamilies[ci].PolicyType.ValueString() == "disable-peer" {
				if t := cr.Get("maxPrefixConfig.prefixNum.optionType"); t.Exists() {
					va := cr.Get("maxPrefixConfig.prefixNum.value")
					if t.String() == "variable" {
						data.Ipv4Neighbors[i].AddressFamilies[ci].DisablePeerMaxNumberOfPrefixesVariable = types.StringValue(va.String())
					} else if t.String() == "global" {
						data.Ipv4Neighbors[i].AddressFamilies[ci].DisablePeerMaxNumberOfPrefixes = types.Int64Value(va.Int())
					}
				}
			}
			data.Ipv4Neighbors[i].AddressFamilies[ci].DisablePeerThreshold = types.Int64Null()
			data.Ipv4Neighbors[i].AddressFamilies[ci].DisablePeerThresholdVariable = types.StringNull()
			if true && data.Ipv4Neighbors[i].AddressFamilies[ci].PolicyType.ValueString() == "disable-peer" {
				if t := cr.Get("maxPrefixConfig.threshold.optionType"); t.Exists() {
					va := cr.Get("maxPrefixConfig.threshold.value")
					if t.String() == "variable" {
						data.Ipv4Neighbors[i].AddressFamilies[ci].DisablePeerThresholdVariable = types.StringValue(va.String())
					} else if t.String() == "global" {
						data.Ipv4Neighbors[i].AddressFamilies[ci].DisablePeerThreshold = types.Int64Value(va.Int())
					}
				}
			}
			data.Ipv4Neighbors[i].AddressFamilies[ci].InRoutePolicyId = types.StringNull()

			if t := cr.Get("inRoutePolicy.refId.optionType"); t.Exists() {
				va := cr.Get("inRoutePolicy.refId.value")
				if t.String() == "global" {
					data.Ipv4Neighbors[i].AddressFamilies[ci].InRoutePolicyId = types.StringValue(va.String())
				}
			}
			data.Ipv4Neighbors[i].AddressFamilies[ci].OutRoutePolicyId = types.StringNull()

			if t := cr.Get("outRoutePolicy.refId.optionType"); t.Exists() {
				va := cr.Get("outRoutePolicy.refId.value")
				if t.String() == "global" {
					data.Ipv4Neighbors[i].AddressFamilies[ci].OutRoutePolicyId = types.StringValue(va.String())
				}
			}
		}
	}
	for i := range data.Ipv6Neighbors {
		keys := [...]string{"address"}
		keyValues := [...]string{data.Ipv6Neighbors[i].Address.ValueString()}
		keyValuesVariables := [...]string{data.Ipv6Neighbors[i].AddressVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "ipv6Neighbor").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
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
		data.Ipv6Neighbors[i].Address = types.StringNull()
		data.Ipv6Neighbors[i].AddressVariable = types.StringNull()
		if t := r.Get("address.optionType"); t.Exists() {
			va := r.Get("address.value")
			if t.String() == "variable" {
				data.Ipv6Neighbors[i].AddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Neighbors[i].Address = types.StringValue(va.String())
			}
		}
		data.Ipv6Neighbors[i].Description = types.StringNull()
		data.Ipv6Neighbors[i].DescriptionVariable = types.StringNull()
		if t := r.Get("description.optionType"); t.Exists() {
			va := r.Get("description.value")
			if t.String() == "variable" {
				data.Ipv6Neighbors[i].DescriptionVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Neighbors[i].Description = types.StringValue(va.String())
			}
		}
		data.Ipv6Neighbors[i].Shutdown = types.BoolNull()
		data.Ipv6Neighbors[i].ShutdownVariable = types.StringNull()
		if t := r.Get("shutdown.optionType"); t.Exists() {
			va := r.Get("shutdown.value")
			if t.String() == "variable" {
				data.Ipv6Neighbors[i].ShutdownVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Neighbors[i].Shutdown = types.BoolValue(va.Bool())
			}
		}
		data.Ipv6Neighbors[i].RemoteAs = types.Int64Null()
		data.Ipv6Neighbors[i].RemoteAsVariable = types.StringNull()
		if t := r.Get("remoteAs.optionType"); t.Exists() {
			va := r.Get("remoteAs.value")
			if t.String() == "variable" {
				data.Ipv6Neighbors[i].RemoteAsVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Neighbors[i].RemoteAs = types.Int64Value(va.Int())
			}
		}
		data.Ipv6Neighbors[i].LocalAs = types.Int64Null()
		data.Ipv6Neighbors[i].LocalAsVariable = types.StringNull()
		if t := r.Get("localAs.optionType"); t.Exists() {
			va := r.Get("localAs.value")
			if t.String() == "variable" {
				data.Ipv6Neighbors[i].LocalAsVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Neighbors[i].LocalAs = types.Int64Value(va.Int())
			}
		}
		data.Ipv6Neighbors[i].KeepaliveTime = types.Int64Null()
		data.Ipv6Neighbors[i].KeepaliveTimeVariable = types.StringNull()
		if t := r.Get("keepalive.optionType"); t.Exists() {
			va := r.Get("keepalive.value")
			if t.String() == "variable" {
				data.Ipv6Neighbors[i].KeepaliveTimeVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Neighbors[i].KeepaliveTime = types.Int64Value(va.Int())
			}
		}
		data.Ipv6Neighbors[i].HoldTime = types.Int64Null()
		data.Ipv6Neighbors[i].HoldTimeVariable = types.StringNull()
		if t := r.Get("holdtime.optionType"); t.Exists() {
			va := r.Get("holdtime.value")
			if t.String() == "variable" {
				data.Ipv6Neighbors[i].HoldTimeVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Neighbors[i].HoldTime = types.Int64Value(va.Int())
			}
		}
		data.Ipv6Neighbors[i].UpdateSourceInterface = types.StringNull()
		data.Ipv6Neighbors[i].UpdateSourceInterfaceVariable = types.StringNull()
		if t := r.Get("ifName.optionType"); t.Exists() {
			va := r.Get("ifName.value")
			if t.String() == "variable" {
				data.Ipv6Neighbors[i].UpdateSourceInterfaceVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Neighbors[i].UpdateSourceInterface = types.StringValue(va.String())
			}
		}
		data.Ipv6Neighbors[i].NextHopSelf = types.BoolNull()
		data.Ipv6Neighbors[i].NextHopSelfVariable = types.StringNull()
		if t := r.Get("nextHopSelf.optionType"); t.Exists() {
			va := r.Get("nextHopSelf.value")
			if t.String() == "variable" {
				data.Ipv6Neighbors[i].NextHopSelfVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Neighbors[i].NextHopSelf = types.BoolValue(va.Bool())
			}
		}
		data.Ipv6Neighbors[i].SendCommunity = types.BoolNull()
		data.Ipv6Neighbors[i].SendCommunityVariable = types.StringNull()
		if t := r.Get("sendCommunity.optionType"); t.Exists() {
			va := r.Get("sendCommunity.value")
			if t.String() == "variable" {
				data.Ipv6Neighbors[i].SendCommunityVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Neighbors[i].SendCommunity = types.BoolValue(va.Bool())
			}
		}
		data.Ipv6Neighbors[i].SendExtendedCommunity = types.BoolNull()
		data.Ipv6Neighbors[i].SendExtendedCommunityVariable = types.StringNull()
		if t := r.Get("sendExtCommunity.optionType"); t.Exists() {
			va := r.Get("sendExtCommunity.value")
			if t.String() == "variable" {
				data.Ipv6Neighbors[i].SendExtendedCommunityVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Neighbors[i].SendExtendedCommunity = types.BoolValue(va.Bool())
			}
		}
		data.Ipv6Neighbors[i].EbgpMultihop = types.Int64Null()
		data.Ipv6Neighbors[i].EbgpMultihopVariable = types.StringNull()
		if t := r.Get("ebgpMultihop.optionType"); t.Exists() {
			va := r.Get("ebgpMultihop.value")
			if t.String() == "variable" {
				data.Ipv6Neighbors[i].EbgpMultihopVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Neighbors[i].EbgpMultihop = types.Int64Value(va.Int())
			}
		}
		data.Ipv6Neighbors[i].AsOverride = types.BoolNull()
		data.Ipv6Neighbors[i].AsOverrideVariable = types.StringNull()
		if t := r.Get("asOverride.optionType"); t.Exists() {
			va := r.Get("asOverride.value")
			if t.String() == "variable" {
				data.Ipv6Neighbors[i].AsOverrideVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Neighbors[i].AsOverride = types.BoolValue(va.Bool())
			}
		}
		data.Ipv6Neighbors[i].AllowasInNumber = types.Int64Null()
		data.Ipv6Neighbors[i].AllowasInNumberVariable = types.StringNull()
		if t := r.Get("asNumber.optionType"); t.Exists() {
			va := r.Get("asNumber.value")
			if t.String() == "variable" {
				data.Ipv6Neighbors[i].AllowasInNumberVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Neighbors[i].AllowasInNumber = types.Int64Value(va.Int())
			}
		}
		for ci := range data.Ipv6Neighbors[i].AddressFamilies {
			keys := [...]string{"familyType"}
			keyValues := [...]string{data.Ipv6Neighbors[i].AddressFamilies[ci].FamilyType.ValueString()}
			keyValuesVariables := [...]string{""}

			var cr gjson.Result
			r.Get("addressFamily").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType")
						vv := v.Get(keys[ik] + ".value")
						if tt.Exists() && vv.Exists() {
							if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
								found = true
								continue
							} else if tt.String() == "default" {
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
			data.Ipv6Neighbors[i].AddressFamilies[ci].FamilyType = types.StringNull()

			if t := cr.Get("familyType.optionType"); t.Exists() {
				va := cr.Get("familyType.value")
				if t.String() == "global" {
					data.Ipv6Neighbors[i].AddressFamilies[ci].FamilyType = types.StringValue(va.String())
				}
			}
			data.Ipv6Neighbors[i].AddressFamilies[ci].MaxNumberOfPrefixes = types.Int64Null()
			data.Ipv6Neighbors[i].AddressFamilies[ci].MaxNumberOfPrefixesVariable = types.StringNull()
			if t := cr.Get("maxPrefixConfig.prefixNum.optionType"); t.Exists() {
				va := cr.Get("maxPrefixConfig.prefixNum.value")
				if t.String() == "variable" {
					data.Ipv6Neighbors[i].AddressFamilies[ci].MaxNumberOfPrefixesVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv6Neighbors[i].AddressFamilies[ci].MaxNumberOfPrefixes = types.Int64Value(va.Int())
				}
			}
			data.Ipv6Neighbors[i].AddressFamilies[ci].Threshold = types.Int64Null()
			data.Ipv6Neighbors[i].AddressFamilies[ci].ThresholdVariable = types.StringNull()
			if t := cr.Get("maxPrefixConfig.threshold.optionType"); t.Exists() {
				va := cr.Get("maxPrefixConfig.threshold.value")
				if t.String() == "variable" {
					data.Ipv6Neighbors[i].AddressFamilies[ci].ThresholdVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv6Neighbors[i].AddressFamilies[ci].Threshold = types.Int64Value(va.Int())
				}
			}
			data.Ipv6Neighbors[i].AddressFamilies[ci].PolicyType = types.StringNull()

			if t := cr.Get("maxPrefixConfig.policyType.optionType"); t.Exists() {
				va := cr.Get("maxPrefixConfig.policyType.value")
				if t.String() == "global" {
					data.Ipv6Neighbors[i].AddressFamilies[ci].PolicyType = types.StringValue(va.String())
				}
			}
			data.Ipv6Neighbors[i].AddressFamilies[ci].RestartInterval = types.Int64Null()
			data.Ipv6Neighbors[i].AddressFamilies[ci].RestartIntervalVariable = types.StringNull()
			if t := cr.Get("maxPrefixConfig.restartInterval.optionType"); t.Exists() {
				va := cr.Get("maxPrefixConfig.restartInterval.value")
				if t.String() == "variable" {
					data.Ipv6Neighbors[i].AddressFamilies[ci].RestartIntervalVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					data.Ipv6Neighbors[i].AddressFamilies[ci].RestartInterval = types.Int64Value(va.Int())
				}
			}
			data.Ipv6Neighbors[i].AddressFamilies[ci].InRoutePolicyId = types.StringNull()

			if t := cr.Get("inRoutePolicy.refId.optionType"); t.Exists() {
				va := cr.Get("inRoutePolicy.refId.value")
				if t.String() == "global" {
					data.Ipv6Neighbors[i].AddressFamilies[ci].InRoutePolicyId = types.StringValue(va.String())
				}
			}
			data.Ipv6Neighbors[i].AddressFamilies[ci].OutRoutePolicyId = types.StringNull()

			if t := cr.Get("outRoutePolicy.refId.optionType"); t.Exists() {
				va := cr.Get("outRoutePolicy.refId.value")
				if t.String() == "global" {
					data.Ipv6Neighbors[i].AddressFamilies[ci].OutRoutePolicyId = types.StringValue(va.String())
				}
			}
		}
	}
	for i := range data.Ipv4AggregateAddresses {
		keys := [...]string{"prefix.address"}
		keyValues := [...]string{data.Ipv4AggregateAddresses[i].NetworkAddress.ValueString()}
		keyValuesVariables := [...]string{data.Ipv4AggregateAddresses[i].NetworkAddressVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "addressFamily.aggregateAddress").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
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
		data.Ipv4AggregateAddresses[i].NetworkAddress = types.StringNull()
		data.Ipv4AggregateAddresses[i].NetworkAddressVariable = types.StringNull()
		if t := r.Get("prefix.address.optionType"); t.Exists() {
			va := r.Get("prefix.address.value")
			if t.String() == "variable" {
				data.Ipv4AggregateAddresses[i].NetworkAddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4AggregateAddresses[i].NetworkAddress = types.StringValue(va.String())
			}
		}
		data.Ipv4AggregateAddresses[i].SubnetMask = types.StringNull()
		data.Ipv4AggregateAddresses[i].SubnetMaskVariable = types.StringNull()
		if t := r.Get("prefix.mask.optionType"); t.Exists() {
			va := r.Get("prefix.mask.value")
			if t.String() == "variable" {
				data.Ipv4AggregateAddresses[i].SubnetMaskVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4AggregateAddresses[i].SubnetMask = types.StringValue(va.String())
			}
		}
		data.Ipv4AggregateAddresses[i].AsSetPath = types.BoolNull()
		data.Ipv4AggregateAddresses[i].AsSetPathVariable = types.StringNull()
		if t := r.Get("asSet.optionType"); t.Exists() {
			va := r.Get("asSet.value")
			if t.String() == "variable" {
				data.Ipv4AggregateAddresses[i].AsSetPathVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4AggregateAddresses[i].AsSetPath = types.BoolValue(va.Bool())
			}
		}
		data.Ipv4AggregateAddresses[i].SummaryOnly = types.BoolNull()
		data.Ipv4AggregateAddresses[i].SummaryOnlyVariable = types.StringNull()
		if t := r.Get("summaryOnly.optionType"); t.Exists() {
			va := r.Get("summaryOnly.value")
			if t.String() == "variable" {
				data.Ipv4AggregateAddresses[i].SummaryOnlyVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4AggregateAddresses[i].SummaryOnly = types.BoolValue(va.Bool())
			}
		}
	}
	for i := range data.Ipv4Networks {
		keys := [...]string{"prefix.address"}
		keyValues := [...]string{data.Ipv4Networks[i].NetworkAddress.ValueString()}
		keyValuesVariables := [...]string{data.Ipv4Networks[i].NetworkAddressVariable.ValueString()}

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
						} else if tt.String() == "default" {
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
		data.Ipv4Networks[i].NetworkAddress = types.StringNull()
		data.Ipv4Networks[i].NetworkAddressVariable = types.StringNull()
		if t := r.Get("prefix.address.optionType"); t.Exists() {
			va := r.Get("prefix.address.value")
			if t.String() == "variable" {
				data.Ipv4Networks[i].NetworkAddressVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Networks[i].NetworkAddress = types.StringValue(va.String())
			}
		}
		data.Ipv4Networks[i].SubnetMask = types.StringNull()
		data.Ipv4Networks[i].SubnetMaskVariable = types.StringNull()
		if t := r.Get("prefix.mask.optionType"); t.Exists() {
			va := r.Get("prefix.mask.value")
			if t.String() == "variable" {
				data.Ipv4Networks[i].SubnetMaskVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Networks[i].SubnetMask = types.StringValue(va.String())
			}
		}
	}
	data.Ipv4EibgpMaximumPaths = types.Int64Null()
	data.Ipv4EibgpMaximumPathsVariable = types.StringNull()
	if t := res.Get(path + "addressFamily.paths.optionType"); t.Exists() {
		va := res.Get(path + "addressFamily.paths.value")
		if t.String() == "variable" {
			data.Ipv4EibgpMaximumPathsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4EibgpMaximumPaths = types.Int64Value(va.Int())
		}
	}
	data.Ipv4Originate = types.BoolNull()
	data.Ipv4OriginateVariable = types.StringNull()
	if t := res.Get(path + "addressFamily.originate.optionType"); t.Exists() {
		va := res.Get(path + "addressFamily.originate.value")
		if t.String() == "variable" {
			data.Ipv4OriginateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4Originate = types.BoolValue(va.Bool())
		}
	}
	data.Ipv4TableMapRoutePolicyId = types.StringNull()

	if t := res.Get(path + "addressFamily.name.refId.optionType"); t.Exists() {
		va := res.Get(path + "addressFamily.name.refId.value")
		if t.String() == "global" {
			data.Ipv4TableMapRoutePolicyId = types.StringValue(va.String())
		}
	}
	data.Ipv4TableMapFilter = types.BoolNull()
	data.Ipv4TableMapFilterVariable = types.StringNull()
	if t := res.Get(path + "addressFamily.filter.optionType"); t.Exists() {
		va := res.Get(path + "addressFamily.filter.value")
		if t.String() == "variable" {
			data.Ipv4TableMapFilterVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv4TableMapFilter = types.BoolValue(va.Bool())
		}
	}
	for i := range data.Ipv4Redistributes {
		keys := [...]string{"protocol", "routePolicy.refId", "translateRibMetric"}
		keyValues := [...]string{data.Ipv4Redistributes[i].Protocol.ValueString(), data.Ipv4Redistributes[i].RoutePolicyId.ValueString(), strconv.FormatBool(data.Ipv4Redistributes[i].TranslateRibMetric.ValueBool())}
		keyValuesVariables := [...]string{data.Ipv4Redistributes[i].ProtocolVariable.ValueString(), "", ""}

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
						} else if tt.String() == "default" {
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
		data.Ipv4Redistributes[i].Protocol = types.StringNull()
		data.Ipv4Redistributes[i].ProtocolVariable = types.StringNull()
		if t := r.Get("protocol.optionType"); t.Exists() {
			va := r.Get("protocol.value")
			if t.String() == "variable" {
				data.Ipv4Redistributes[i].ProtocolVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv4Redistributes[i].Protocol = types.StringValue(va.String())
			}
		}
		data.Ipv4Redistributes[i].RoutePolicyId = types.StringNull()

		if t := r.Get("routePolicy.refId.optionType"); t.Exists() {
			va := r.Get("routePolicy.refId.value")
			if t.String() == "global" {
				data.Ipv4Redistributes[i].RoutePolicyId = types.StringValue(va.String())
			}
		}
		data.Ipv4Redistributes[i].TranslateRibMetric = types.BoolNull()

		if t := r.Get("translateRibMetric.optionType"); t.Exists() {
			va := r.Get("translateRibMetric.value")
			if t.String() == "global" {
				data.Ipv4Redistributes[i].TranslateRibMetric = types.BoolValue(va.Bool())
			}
		}
	}
	for i := range data.Ipv6AggregateAddresses {
		keys := [...]string{"prefix"}
		keyValues := [...]string{data.Ipv6AggregateAddresses[i].AggregatePrefix.ValueString()}
		keyValuesVariables := [...]string{data.Ipv6AggregateAddresses[i].AggregatePrefixVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "ipv6AddressFamily.ipv6AggregateAddress").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
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
		data.Ipv6AggregateAddresses[i].AggregatePrefix = types.StringNull()
		data.Ipv6AggregateAddresses[i].AggregatePrefixVariable = types.StringNull()
		if t := r.Get("prefix.optionType"); t.Exists() {
			va := r.Get("prefix.value")
			if t.String() == "variable" {
				data.Ipv6AggregateAddresses[i].AggregatePrefixVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6AggregateAddresses[i].AggregatePrefix = types.StringValue(va.String())
			}
		}
		data.Ipv6AggregateAddresses[i].AsSetPath = types.BoolNull()
		data.Ipv6AggregateAddresses[i].AsSetPathVariable = types.StringNull()
		if t := r.Get("asSet.optionType"); t.Exists() {
			va := r.Get("asSet.value")
			if t.String() == "variable" {
				data.Ipv6AggregateAddresses[i].AsSetPathVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6AggregateAddresses[i].AsSetPath = types.BoolValue(va.Bool())
			}
		}
		data.Ipv6AggregateAddresses[i].SummaryOnly = types.BoolNull()
		data.Ipv6AggregateAddresses[i].SummaryOnlyVariable = types.StringNull()
		if t := r.Get("summaryOnly.optionType"); t.Exists() {
			va := r.Get("summaryOnly.value")
			if t.String() == "variable" {
				data.Ipv6AggregateAddresses[i].SummaryOnlyVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6AggregateAddresses[i].SummaryOnly = types.BoolValue(va.Bool())
			}
		}
	}
	for i := range data.Ipv6Networks {
		keys := [...]string{"prefix"}
		keyValues := [...]string{data.Ipv6Networks[i].NetworkPrefix.ValueString()}
		keyValuesVariables := [...]string{data.Ipv6Networks[i].NetworkPrefixVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "ipv6AddressFamily.ipv6Network").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
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
		data.Ipv6Networks[i].NetworkPrefix = types.StringNull()
		data.Ipv6Networks[i].NetworkPrefixVariable = types.StringNull()
		if t := r.Get("prefix.optionType"); t.Exists() {
			va := r.Get("prefix.value")
			if t.String() == "variable" {
				data.Ipv6Networks[i].NetworkPrefixVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Networks[i].NetworkPrefix = types.StringValue(va.String())
			}
		}
	}
	data.Ipv6EibgpMaximumPaths = types.Int64Null()
	data.Ipv6EibgpMaximumPathsVariable = types.StringNull()
	if t := res.Get(path + "ipv6AddressFamily.paths.optionType"); t.Exists() {
		va := res.Get(path + "ipv6AddressFamily.paths.value")
		if t.String() == "variable" {
			data.Ipv6EibgpMaximumPathsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6EibgpMaximumPaths = types.Int64Value(va.Int())
		}
	}
	data.Ipv6Originate = types.BoolNull()
	data.Ipv6OriginateVariable = types.StringNull()
	if t := res.Get(path + "ipv6AddressFamily.originate.optionType"); t.Exists() {
		va := res.Get(path + "ipv6AddressFamily.originate.value")
		if t.String() == "variable" {
			data.Ipv6OriginateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6Originate = types.BoolValue(va.Bool())
		}
	}
	data.Ipv6TableMapRoutePolicyId = types.StringNull()

	if t := res.Get(path + "ipv6AddressFamily.name.refId.optionType"); t.Exists() {
		va := res.Get(path + "ipv6AddressFamily.name.refId.value")
		if t.String() == "global" {
			data.Ipv6TableMapRoutePolicyId = types.StringValue(va.String())
		}
	}
	data.Ipv6TableMapFilter = types.BoolNull()
	data.Ipv6TableMapFilterVariable = types.StringNull()
	if t := res.Get(path + "ipv6AddressFamily.filter.optionType"); t.Exists() {
		va := res.Get(path + "ipv6AddressFamily.filter.value")
		if t.String() == "variable" {
			data.Ipv6TableMapFilterVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Ipv6TableMapFilter = types.BoolValue(va.Bool())
		}
	}
	for i := range data.Ipv6Redistributes {
		keys := [...]string{"protocol", "routePolicy.refId", "translateRibMetric"}
		keyValues := [...]string{data.Ipv6Redistributes[i].Protocol.ValueString(), data.Ipv6Redistributes[i].RoutePolicyId.ValueString(), strconv.FormatBool(data.Ipv6Redistributes[i].TranslateRibMetric.ValueBool())}
		keyValuesVariables := [...]string{data.Ipv6Redistributes[i].ProtocolVariable.ValueString(), "", ""}

		var r gjson.Result
		res.Get(path + "ipv6AddressFamily.redistribute").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
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
		data.Ipv6Redistributes[i].Protocol = types.StringNull()
		data.Ipv6Redistributes[i].ProtocolVariable = types.StringNull()
		if t := r.Get("protocol.optionType"); t.Exists() {
			va := r.Get("protocol.value")
			if t.String() == "variable" {
				data.Ipv6Redistributes[i].ProtocolVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.Ipv6Redistributes[i].Protocol = types.StringValue(va.String())
			}
		}
		data.Ipv6Redistributes[i].RoutePolicyId = types.StringNull()

		if t := r.Get("routePolicy.refId.optionType"); t.Exists() {
			va := r.Get("routePolicy.refId.value")
			if t.String() == "global" {
				data.Ipv6Redistributes[i].RoutePolicyId = types.StringValue(va.String())
			}
		}
		data.Ipv6Redistributes[i].TranslateRibMetric = types.BoolNull()

		if t := r.Get("translateRibMetric.optionType"); t.Exists() {
			va := r.Get("translateRibMetric.value")
			if t.String() == "global" {
				data.Ipv6Redistributes[i].TranslateRibMetric = types.BoolValue(va.Bool())
			}
		}
	}
}
